package changelog

import (
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/cache"
	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/config"
	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/git"
	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/github"
)

type Generator struct {
	cfg       *config.Config
	gitWalker *git.Walker
	ghClient  *github.Client
	cache     *cache.Cache
	versions  map[string]*git.Version
	prs       map[int]*github.PR
}

func New(cfg *config.Config) (*Generator, error) {
	gitWalker, err := git.NewWalker(cfg.RepoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create git walker: %w", err)
	}

	owner, repo, err := gitWalker.GetRepoInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get repo info: %w", err)
	}

	ghClient := github.NewClient(cfg.GitHubToken, owner, repo)

	var c *cache.Cache
	if !cfg.NoCache {
		c, err = cache.New(cfg.CacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to create cache: %w", err)
		}

		if cfg.RebuildCache {
			if err := c.Clear(); err != nil {
				return nil, fmt.Errorf("failed to clear cache: %w", err)
			}
		}
	}

	return &Generator{
		cfg:       cfg,
		gitWalker: gitWalker,
		ghClient:  ghClient,
		cache:     c,
		prs:       make(map[int]*github.PR),
	}, nil
}

func (g *Generator) Generate() (string, error) {
	if err := g.collectData(); err != nil {
		return "", fmt.Errorf("failed to collect data: %w", err)
	}

	if err := g.fetchPRs(g.cfg.ForcePRSync); err != nil {
		return "", fmt.Errorf("failed to fetch PRs: %w", err)
	}

	return g.formatChangelog(), nil
}

func (g *Generator) collectData() error {
	if g.cache != nil && !g.cfg.RebuildCache {
		cachedTag, err := g.cache.GetLastProcessedTag()
		if err != nil {
			return fmt.Errorf("failed to get last processed tag: %w", err)
		}

		if cachedTag != "" {
			currentTag, err := g.gitWalker.GetLatestTag()
			if err == nil {
				// Cached versions stay valid when git has newer tags. The new tags merge in below.
				cachedVersions, err := g.cache.GetVersions()
				if err == nil && len(cachedVersions) > 0 {
					g.versions = cachedVersions

					for _, version := range g.versions {
						for _, prNum := range version.PRNumbers {
							if pr, err := g.cache.GetPR(prNum); err == nil && pr != nil {
								g.prs[prNum] = pr
							}
						}
					}

					// Walk only the versions added after the cached tag.
					if currentTag != cachedTag {
						fmt.Fprintf(os.Stderr, "Processing new versions since %s...\n", cachedTag)
						newVersions, err := g.gitWalker.WalkHistorySinceTag(cachedTag)
						if err != nil {
							fmt.Fprintf(os.Stderr, "Warning: Failed to walk history since tag %s: %v\n", cachedTag, err)
						} else {
							// Add new versions to the cached set. Keep cached versions, but fill in missing PR numbers.
							for name, version := range newVersions {
								if name != "Unreleased" { // rebuilt below
									if existingVersion, exists := g.versions[name]; !exists {
										g.versions[name] = version
									} else {
										if len(existingVersion.PRNumbers) == 0 && len(version.PRNumbers) > 0 {
											existingVersion.PRNumbers = version.PRNumbers
										}
									}
								}
							}
						}
					}

					// Rebuild Unreleased from git on every run.
					unreleasedVersion, err := g.gitWalker.WalkCommitsSinceTag(currentTag)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Warning: Failed to walk commits since tag %s: %v\n", currentTag, err)
					} else if unreleasedVersion != nil {
						if existingUnreleased, exists := g.versions["Unreleased"]; exists {
							unreleasedVersion.AISummary = existingUnreleased.AISummary
						}
						g.versions["Unreleased"] = unreleasedVersion
					}

					// Save every version and its commits. Unreleased is never saved.
					if currentTag != cachedTag {
						for _, version := range g.versions {
							if version.Name != "Unreleased" {
								if err := g.cache.SaveVersion(version); err != nil {
									fmt.Fprintf(os.Stderr, "Warning: Failed to save version to cache: %v\n", err)
								}

								for _, commit := range version.Commits {
									if err := g.cache.SaveCommit(commit, version.Name); err != nil {
										fmt.Fprintf(os.Stderr, "Warning: Failed to save commit to cache: %v\n", err)
									}
								}
							}
						}

						if err := g.cache.SetLastProcessedTag(currentTag); err != nil {
							fmt.Fprintf(os.Stderr, "Warning: Failed to update last processed tag: %v\n", err)
						}
					}

					return nil
				}
			}
		}
	}

	versions, err := g.gitWalker.WalkHistory()
	if err != nil {
		return fmt.Errorf("failed to walk history: %w", err)
	}

	g.versions = versions

	if g.cache != nil {
		for _, version := range versions {
			if err := g.cache.SaveVersion(version); err != nil {
				return fmt.Errorf("failed to save version to cache: %w", err)
			}

			for _, commit := range version.Commits {
				if err := g.cache.SaveCommit(commit, version.Name); err != nil {
					return fmt.Errorf("failed to save commit to cache: %w", err)
				}
			}
		}

		// The latest tag is the anchor for later incremental walks.
		if latestTag, err := g.gitWalker.GetLatestTag(); err == nil && latestTag != "" {
			if err := g.cache.SetLastProcessedTag(latestTag); err != nil {
				return fmt.Errorf("failed to save last processed tag: %w", err)
			}
		}
	}

	return nil
}

func (g *Generator) fetchPRs(forcePRSync bool) error {
	if g.cache != nil {
		cachedPRs, err := g.cache.GetAllPRs()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to load cached PRs: %v\n", err)
		} else {
			g.prs = cachedPRs
		}
	}

	var lastSync time.Time
	if g.cache != nil {
		lastSync, _ = g.cache.GetLastPRSync()
	}

	missingPRs := false
	for _, version := range g.versions {
		for _, prNum := range version.PRNumbers {
			if _, exists := g.prs[prNum]; !exists {
				missingPRs = true
				break
			}
		}
		if missingPRs {
			break
		}
	}

	if missingPRs {
		fmt.Fprintf(os.Stderr, "Full sync triggered due to missing PRs in cache.\n")
	}
	needsSync := lastSync.IsZero() || time.Since(lastSync) > 24*time.Hour || forcePRSync || missingPRs

	if !needsSync {
		fmt.Fprintf(os.Stderr, "Using cached PR data (last sync: %s)\n", lastSync.Format("2006-01-02 15:04:05"))
		return nil
	}

	fmt.Fprintf(os.Stderr, "Fetching merged PRs from GitHub using GraphQL...\n")

	// GraphQL returns PRs with their commits in pages of 100. REST needs one call per PR.
	prs, err := g.ghClient.FetchAllMergedPRsGraphQL(lastSync)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GraphQL fetch failed, falling back to REST API: %v\n", err)
		prs, err = g.ghClient.FetchAllMergedPRs(lastSync)
		if err != nil {
			return fmt.Errorf("both GraphQL and REST API failed: %w", err)
		}
	}

	for _, pr := range prs {
		g.prs[pr.Number] = pr
	}

	if g.cache != nil && len(prs) > 0 {
		if err := g.cache.SavePRBatch(prs); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to cache PRs: %v\n", err)
		}

		if err := g.cache.SaveCommitPRMappings(prs); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to cache commit mappings: %v\n", err)
		}

		if err := g.cache.SetLastPRSync(time.Now()); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to update last sync timestamp: %v\n", err)
		}
	}

	if len(prs) > 0 {
		fmt.Fprintf(os.Stderr, "Fetched %d PRs with commits (total cached: %d)\n", len(prs), len(g.prs))
	}

	return nil
}

func (g *Generator) formatChangelog() string {
	var sb strings.Builder
	sb.WriteString("# Changelog\n")

	versionList := g.getSortedVersions()

	for _, version := range versionList {
		if g.cfg.Version != "" && version.Name != g.cfg.Version {
			continue
		}

		versionText := g.formatVersion(version)
		if versionText != "" {
			sb.WriteString("\n")
			sb.WriteString(versionText)
		}
	}

	return sb.String()
}

func (g *Generator) getSortedVersions() []*git.Version {
	var versions []*git.Version
	var releasedVersions []*git.Version

	for name, version := range g.versions {
		if name != "Unreleased" {
			releasedVersions = append(releasedVersions, version)
		}
	}

	sort.Slice(releasedVersions, func(i, j int) bool {
		return releasedVersions[i].Date.After(releasedVersions[j].Date)
	})

	if unreleased, exists := g.versions["Unreleased"]; exists && len(unreleased.Commits) > 0 {
		versions = append(versions, unreleased)
	}

	versions = append(versions, releasedVersions...)

	if g.cfg.Limit > 0 && len(versions) > g.cfg.Limit {
		versions = versions[:g.cfg.Limit]
	}

	return versions
}

func (g *Generator) formatVersion(version *git.Version) string {
	var sb strings.Builder

	rawContent := g.generateRawVersionContent(version)
	if rawContent == "" {
		return ""
	}

	header := g.formatVersionHeader(version)
	sb.WriteString(("\n"))
	sb.WriteString(header)

	if g.cfg.EnableAISummary {
		// Reuse the Unreleased summary when the raw content hash is unchanged.
		if version.Name == "Unreleased" && version.AISummary != "" && g.cache != nil {
			cachedHash, err := g.cache.GetUnreleasedContentHash()
			if err == nil {
				currentHash := hashContent(rawContent)
				if cachedHash == currentHash {
					fmt.Fprintf(os.Stderr, "✅ %s content unchanged (skipping AI)\n", version.Name)
					sb.WriteString(version.AISummary)
					return fixMarkdown(sb.String())
				}
			}
		}

		// Released content does not change, so a cached summary is final.
		if version.Name != "Unreleased" && version.AISummary != "" {
			fmt.Fprintf(os.Stderr, "✅ %s already summarized (skipping)\n", version.Name)
			sb.WriteString(version.AISummary)
			return fixMarkdown(sb.String())
		}

		fmt.Fprintf(os.Stderr, "🤖 AI summarizing %s...", version.Name)

		aiSummary, err := SummarizeVersionContent(rawContent)
		if err != nil {
			fmt.Fprintf(os.Stderr, " Failed: %v\n", err)
			sb.WriteString((rawContent))
			return fixMarkdown(sb.String())
		}
		if checkForAIError(aiSummary) {
			fmt.Fprintf(os.Stderr, " AI error detected, using raw content instead\n")
			sb.WriteString(rawContent)
			fmt.Fprintf(os.Stderr, "Raw Content was: (%d bytes) %s \n", len(rawContent), rawContent)
			fmt.Fprintf(os.Stderr, "AI Summary was: (%d bytes) %s\n", len(aiSummary), aiSummary)
			return fixMarkdown(sb.String())
		}

		fmt.Fprintf(os.Stderr, " Done!\n")
		aiSummary = strings.TrimSpace(aiSummary)

		version.AISummary = aiSummary
		if g.cache != nil {
			if err := g.cache.UpdateVersionAISummary(version.Name, aiSummary); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to cache AI summary: %v\n", err)
			}
			// The hash lets the next run detect Unreleased content changes.
			if version.Name == "Unreleased" {
				if err := g.cache.SetUnreleasedContentHash(hashContent(rawContent)); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: Failed to cache content hash: %v\n", err)
				}
			}
		}

		sb.WriteString(aiSummary)
		return fixMarkdown(sb.String())
	}

	sb.WriteString(rawContent)
	return fixMarkdown(sb.String())
}

func checkForAIError(summary string) bool {
	errorPatterns := []string{
		"I don't see any", "please provide",
		"content you've provided appears to be incomplete",
	}

	for _, pattern := range errorPatterns {
		if strings.Contains(summary, pattern) {
			return true
		}
	}

	return false
}

func (g *Generator) formatVersionHeader(version *git.Version) string {
	if version.Name == "Unreleased" {
		return "## Unreleased\n\n"
	}
	return fmt.Sprintf("\n## %s (%s)\n\n", version.Name, version.Date.Format("2006-01-02"))
}

// generateRawVersionContent renders the PRs and direct commits of a version as markdown.
func (g *Generator) generateRawVersionContent(version *git.Version) string {
	var sb strings.Builder

	prCommitSHAs := make(map[string]bool)
	for _, prNum := range version.PRNumbers {
		if pr, exists := g.prs[prNum]; exists {
			for _, prCommit := range pr.Commits {
				prCommitSHAs[prCommit.SHA] = true
			}
		}
	}

	prCommits := make(map[int][]*git.Commit)
	directCommits := []*git.Commit{}

	for _, commit := range version.Commits {
		if commit.IsVersion {
			continue
		}

		// Commits of a fetched PR appear under that PR, not as direct commits.
		if prCommitSHAs[commit.SHA] {
			continue
		}

		if commit.PRNumber > 0 {
			prCommits[commit.PRNumber] = append(prCommits[commit.PRNumber], commit)
		} else {
			directCommits = append(directCommits, commit)
		}
	}

	// A version with only a version bump commit has no content.
	// PRs listed in version.PRNumbers still count as content.
	if len(prCommits) == 0 && len(directCommits) == 0 && len(version.PRNumbers) == 0 {
		return ""
	}

	prependNewline := ""
	for _, prNum := range version.PRNumbers {
		if pr, exists := g.prs[prNum]; exists {
			sb.WriteString(prependNewline)
			sb.WriteString(g.formatPR(pr))
			prependNewline = "\n"
		}
	}

	if len(directCommits) > 0 {
		sort.Slice(directCommits, func(i, j int) bool {
			return directCommits[i].Date.After(directCommits[j].Date)
		})

		sb.WriteString(prependNewline + "### Direct commits\n\n")
		for _, commit := range directCommits {
			message := g.formatCommitMessage(strings.TrimSpace(commit.Message))
			if message != "" && !g.isDuplicateMessage(message, directCommits) {
				sb.WriteString(fmt.Sprintf("- %s\n", message))
			}
		}
	}

	return fixMarkdown(
		strings.ReplaceAll(sb.String(), "\n-\n", "\n"), // Remove empty list items
	)
}

func fixMarkdown(content string) string {

	// markdownlint MD032 (blanks-around-lists): insert a blank line before each list.
	lines := strings.Split(content, "\n")
	inList := false
	preListNewline := false
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			if !inList {
				inList = true
				if !preListNewline && i > 0 && lines[i-1] != "" {
					line = "\n" + line
					preListNewline = true
				}
			}
		} else {
			if inList {
				inList = false
				preListNewline = false
			}
		}
		lines[i] = strings.TrimRight(line, " \t")
	}

	fixedContent := strings.TrimSpace(strings.Join(lines, "\n"))

	return fixedContent + "\n"
}

func (g *Generator) formatPR(pr *github.PR) string {
	var sb strings.Builder

	pr.Title = strings.TrimRight(strings.TrimSpace(pr.Title), ".")

	// Add type indicator for non-users
	authorName := pr.Author
	switch pr.AuthorType {
	case "bot":
		authorName += "[bot]"
	case "organization":
		authorName += "[org]"
	}

	sb.WriteString(fmt.Sprintf("### PR [#%d](%s) by [%s](%s): %s\n\n",
		pr.Number, pr.URL, authorName, pr.AuthorURL, strings.TrimSpace(pr.Title)))

	changes := g.extractChanges(pr)
	for _, change := range changes {
		if change != "" {
			sb.WriteString(fmt.Sprintf("- %s\n", change))
		}
	}

	return sb.String()
}

func (g *Generator) extractChanges(pr *github.PR) []string {
	var changes []string
	seen := make(map[string]bool)

	for _, commit := range pr.Commits {
		message := g.formatCommitMessage(commit.Message)
		if message != "" && !seen[message] {
			seen[message] = true
			changes = append(changes, message)
		}
	}

	if len(changes) == 0 && pr.Body != "" {
		lines := strings.SplitSeq(pr.Body, "\n")
		for line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
				change := strings.TrimPrefix(strings.TrimPrefix(line, "- "), "* ")
				if change != "" {
					changes = append(changes, change)
				}
			}
		}
	}

	return changes
}

func normalizeLineEndings(content string) string {
	return strings.ReplaceAll(content, "\r\n", "\n")
}

func (g *Generator) formatCommitMessage(message string) string {
	strings_to_remove := []string{
		"### CHANGES\n", "## CHANGES\n", "# CHANGES\n",
		"...\n", "---\n", "## Changes\n", "## Change",
		"Update version to v..1 and commit\n",
		"# What this Pull Request (PR) does\n",
		"# Conflicts:",
	}

	message = normalizeLineEndings(message)
	message = strings.ReplaceAll(message, "\t", " ")

	if len(message) > 0 {
		message = strings.ToUpper(message[:1]) + message[1:]
	}

	for _, str := range strings_to_remove {
		if strings.Contains(message, str) {
			message = strings.ReplaceAll(message, str, "")
		}
	}

	message = fixFormatting(message)

	return message
}

func fixFormatting(message string) string {
	message = strings.ReplaceAll(message, "* ", "- ")
	message = strings.ReplaceAll(message, "-   ", "- ")
	message = strings.ReplaceAll(message, "-  ", "- ")
	if strings.Contains(message, "http://") || strings.Contains(message, "https://") {
		urlRegex := regexp.MustCompile(`\b(https?://[^\s<>]+)`)
		message = urlRegex.ReplaceAllString(message, "<$1>")
	}

	message = strings.ReplaceAll(message, "## LINKS\n", "- ")
	message = strings.TrimSpace(message)
	message = strings.TrimSpace(strings.ReplaceAll(message, "\n\n", "\n"))
	message = strings.ReplaceAll(message, " \n", "\n")
	message = strings.ReplaceAll(message, "\n - ", "\n- ")

	// markdownlint MD032: add a blank line between a text line and the list item after it.
	// A text line that ends in a colon is exempt.
	listRegex := regexp.MustCompile(`(?m)([^\n-].*[^:\n])\n([-*] .*)`)
	message = listRegex.ReplaceAllString(message, "$1\n\n$2")

	// Rewrite "# " headings as level 4. Commit messages use "#" for notes, not document headings.
	message = strings.ReplaceAll(message, "# ", "\n#### ")
	message = strings.ReplaceAll(message, "\n\n\n", "\n\n")

	// Wrap emails and URLs in angle brackets. The "<<" cleanup below undoes double wrapping.
	emailRegex := regexp.MustCompile(`([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
	message = emailRegex.ReplaceAllString(message, "<$1>")

	urlRegex := regexp.MustCompile(`(https?://[^\s<]+)`)
	message = urlRegex.ReplaceAllString(message, "<$1>")

	message = strings.ReplaceAll(message, "<<", "<")
	message = strings.ReplaceAll(message, ">>", ">")

	// A line that starts with "#123" trips markdownlint MD018. Join it to the line before.
	prOrIssueLinkRegex := regexp.MustCompile("\n" + `(#\d+)`)
	message = prOrIssueLinkRegex.ReplaceAllString(message, " $1")

	message = strings.TrimSpace(message)
	return message
}

func (g *Generator) isDuplicateMessage(message string, commits []*git.Commit) bool {
	if message == "." || strings.ToLower(message) == "fix" {
		count := 0
		for _, commit := range commits {
			formatted := g.formatCommitMessage(commit.Message)
			if formatted == message {
				count++
				if count > 1 {
					return true
				}
			}
		}
	}
	return false
}

// hashContent returns the SHA-256 hex digest that detects Unreleased content changes.
func hashContent(content string) string {
	hash := sha256.Sum256([]byte(content))
	return fmt.Sprintf("%x", hash)
}

// SyncDatabase performs a comprehensive database synchronization and validation
func (g *Generator) SyncDatabase() error {
	if g.cache == nil {
		return fmt.Errorf("cache is disabled, cannot sync database")
	}

	fmt.Fprintf(os.Stderr, "[SYNC] Starting database synchronization...\n")

	fmt.Fprintf(os.Stderr, "[PR_SYNC] Forcing PR sync from GitHub...\n")
	if err := g.fetchPRs(true); err != nil {
		return fmt.Errorf("failed to sync PRs: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[VERIFY] Verifying git history and version completeness...\n")
	if err := g.syncGitHistory(); err != nil {
		return fmt.Errorf("failed to sync git history: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[MAPPING] Verifying commit-PR mappings...\n")
	if err := g.verifyCommitPRMappings(); err != nil {
		return fmt.Errorf("failed to verify commit-PR mappings: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[SUCCESS] Database synchronization completed successfully!\n")
	return nil
}

// syncGitHistory walks the full git history and caches the versions and commits that are missing.
func (g *Generator) syncGitHistory() error {
	versions, err := g.gitWalker.WalkHistory()
	if err != nil {
		return fmt.Errorf("failed to walk git history: %w", err)
	}

	var newVersions, newCommits int
	for _, version := range versions {
		exists, err := g.cache.VersionExists(version.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to check existence of version %s: %v. This may affect the completeness of the sync operation.\n", version.Name, err)
			continue
		}
		if !exists {
			if err := g.cache.SaveVersion(version); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to save version %s: %v\n", version.Name, err)
			} else {
				newVersions++
			}
		}

		for _, commit := range version.Commits {
			exists, err := g.cache.CommitExists(commit.SHA)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to check commit %s existence: %v\n", commit.SHA, err)
				continue
			}
			if !exists {
				if err := g.cache.SaveCommit(commit, version.Name); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: Failed to save commit %s: %v\n", commit.SHA, err)
				} else {
					newCommits++
				}
			}
		}
	}

	if latestTag, err := g.gitWalker.GetLatestTag(); err == nil && latestTag != "" {
		if err := g.cache.SetLastProcessedTag(latestTag); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to update last processed tag: %v\n", err)
		}
	}

	fmt.Fprintf(os.Stderr, "   Added %d new versions and %d new commits (preserved existing data)\n", newVersions, newCommits)
	return nil
}

// verifyCommitPRMappings rewrites the commit-to-PR mapping table from all cached PRs.
func (g *Generator) verifyCommitPRMappings() error {
	allPRs, err := g.cache.GetAllPRs()
	if err != nil {
		return fmt.Errorf("failed to get cached PRs: %w", err)
	}

	var prSlice []*github.PR
	for _, pr := range allPRs {
		prSlice = append(prSlice, pr)
	}

	if err := g.cache.SaveCommitPRMappings(prSlice); err != nil {
		return fmt.Errorf("failed to save commit-PR mappings: %w", err)
	}

	fmt.Fprintf(os.Stderr, "   Verified mappings for %d PRs\n", len(prSlice))
	return nil
}
