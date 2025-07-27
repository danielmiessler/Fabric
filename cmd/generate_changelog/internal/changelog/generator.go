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
			// Get the current latest tag from git
			currentTag, err := g.gitWalker.GetLatestTag()
			if err == nil {
				// Load cached data - we can use it even if there are new tags
				cachedVersions, err := g.cache.GetVersions()
				if err == nil && len(cachedVersions) > 0 {
					g.versions = cachedVersions

					// Load cached PRs
					for _, version := range g.versions {
						for _, prNum := range version.PRNumbers {
							if pr, err := g.cache.GetPR(prNum); err == nil && pr != nil {
								g.prs[prNum] = pr
							}
						}
					}

					// If we have new tags since cache, process the new versions only
					if currentTag != cachedTag {
						fmt.Fprintf(os.Stderr, "Processing new versions since %s...\n", cachedTag)
						newVersions, err := g.gitWalker.WalkHistorySinceTag(cachedTag)
						if err != nil {
							fmt.Fprintf(os.Stderr, "Warning: Failed to walk history since tag %s: %v\n", cachedTag, err)
						} else {
							// Merge new versions into cached versions (only add if not already cached)
							for name, version := range newVersions {
								if name != "Unreleased" { // Handle Unreleased separately
									if existingVersion, exists := g.versions[name]; !exists {
										g.versions[name] = version
									} else {
										// Update existing version with new PR numbers if they're missing
										if len(existingVersion.PRNumbers) == 0 && len(version.PRNumbers) > 0 {
											existingVersion.PRNumbers = version.PRNumbers
										}
									}
								}
							}
						}
					}

					// Always update Unreleased section with latest commits
					unreleasedVersion, err := g.gitWalker.WalkCommitsSinceTag(currentTag)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Warning: Failed to walk commits since tag %s: %v\n", currentTag, err)
					} else if unreleasedVersion != nil {
						// Preserve existing AI summary if available
						if existingUnreleased, exists := g.versions["Unreleased"]; exists {
							unreleasedVersion.AISummary = existingUnreleased.AISummary
						}
						// Replace or add the unreleased version
						g.versions["Unreleased"] = unreleasedVersion
					}

					// Save any new versions to cache (after potential AI processing)
					if currentTag != cachedTag {
						for _, version := range g.versions {
							// Skip versions that were already cached and Unreleased
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

						// Update the last processed tag
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

		// Save the latest tag as our cache anchor point
		if latestTag, err := g.gitWalker.GetLatestTag(); err == nil && latestTag != "" {
			if err := g.cache.SetLastProcessedTag(latestTag); err != nil {
				return fmt.Errorf("failed to save last processed tag: %w", err)
			}
		}
	}

	return nil
}

func (g *Generator) fetchPRs(forcePRSync bool) error {
	// First, load all cached PRs
	if g.cache != nil {
		cachedPRs, err := g.cache.GetAllPRs()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to load cached PRs: %v\n", err)
		} else {
			g.prs = cachedPRs
		}
	}

	// Check if we need to fetch new PRs
	var lastSync time.Time
	if g.cache != nil {
		lastSync, _ = g.cache.GetLastPRSync()
	}

	// Check if we need to sync for missing PRs
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
	// If we have never synced or it's been more than 24 hours, do a full sync
	// Also sync if we have versions with PR numbers that aren't cached
	needsSync := lastSync.IsZero() || time.Since(lastSync) > 24*time.Hour || forcePRSync || missingPRs

	if !needsSync {
		fmt.Fprintf(os.Stderr, "Using cached PR data (last sync: %s)\n", lastSync.Format("2006-01-02 15:04:05"))
		return nil
	}

	fmt.Fprintf(os.Stderr, "Fetching merged PRs from GitHub using GraphQL...\n")

	// Use GraphQL for ultimate performance - gets everything in ~5-10 calls
	prs, err := g.ghClient.FetchAllMergedPRsGraphQL(lastSync)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GraphQL fetch failed, falling back to REST API: %v\n", err)
		// Fall back to REST API
		prs, err = g.ghClient.FetchAllMergedPRs(lastSync)
		if err != nil {
			return fmt.Errorf("both GraphQL and REST API failed: %w", err)
		}
	}

	// Update our PR map with new data
	for _, pr := range prs {
		g.prs[pr.Number] = pr
	}

	// Save all PRs to cache in a batch transaction
	if g.cache != nil && len(prs) > 0 {
		// Save PRs
		if err := g.cache.SavePRBatch(prs); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to cache PRs: %v\n", err)
		}

		// Save SHA→PR mappings for lightning-fast git operations
		if err := g.cache.SaveCommitPRMappings(prs); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to cache commit mappings: %v\n", err)
		}

		// Update last sync timestamp
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

	// Collect all released versions (non-"Unreleased")
	for name, version := range g.versions {
		if name != "Unreleased" {
			releasedVersions = append(releasedVersions, version)
		}
	}

	// Sort released versions by date (newest first)
	sort.Slice(releasedVersions, func(i, j int) bool {
		return releasedVersions[i].Date.After(releasedVersions[j].Date)
	})

	// Add "Unreleased" first if it exists and has commits
	if unreleased, exists := g.versions["Unreleased"]; exists && len(unreleased.Commits) > 0 {
		versions = append(versions, unreleased)
	}

	// Add sorted released versions
	versions = append(versions, releasedVersions...)

	if g.cfg.Limit > 0 && len(versions) > g.cfg.Limit {
		versions = versions[:g.cfg.Limit]
	}

	return versions
}

func (g *Generator) formatVersion(version *git.Version) string {
	var sb strings.Builder

	// Generate raw content
	rawContent := g.generateRawVersionContent(version)
	if rawContent == "" {
		return ""
	}

	header := g.formatVersionHeader(version)
	sb.WriteString(("\n"))
	sb.WriteString(header)

	// If AI summarization is enabled, enhance with AI
	if g.cfg.EnableAISummary {
		// For "Unreleased", check if content has changed since last AI summary
		if version.Name == "Unreleased" && version.AISummary != "" && g.cache != nil {
			// Get cached content hash
			cachedHash, err := g.cache.GetUnreleasedContentHash()
			if err == nil {
				// Calculate current content hash
				currentHash := hashContent(rawContent)
				if cachedHash == currentHash {
					// Content unchanged, use cached summary
					fmt.Fprintf(os.Stderr, "✅ %s content unchanged (skipping AI)\n", version.Name)
					sb.WriteString(version.AISummary)
					return fixMarkdown(sb.String())
				}
			}
		}

		// For released versions, if we have cached AI summary, use it!
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

		// Cache the AI summary and content hash
		version.AISummary = aiSummary
		if g.cache != nil {
			if err := g.cache.UpdateVersionAISummary(version.Name, aiSummary); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to cache AI summary: %v\n", err)
			}
			// Cache content hash for "Unreleased" to detect changes
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
	// Check for common AI error patterns
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

// formatVersionHeader formats just the version header (## ...)
func (g *Generator) formatVersionHeader(version *git.Version) string {
	if version.Name == "Unreleased" {
		return "## Unreleased\n\n"
	}
	return fmt.Sprintf("\n## %s (%s)\n\n", version.Name, version.Date.Format("2006-01-02"))
}

// generateRawVersionContent generates the raw content (PRs + commits) for a version
func (g *Generator) generateRawVersionContent(version *git.Version) string {
	var sb strings.Builder

	// Build a set of commit SHAs that are part of fetched PRs
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
		// Skip version bump commits from output
		if commit.IsVersion {
			continue
		}

		// If this commit is part of a fetched PR, don't include it in direct commits
		if prCommitSHAs[commit.SHA] {
			continue
		}

		if commit.PRNumber > 0 {
			prCommits[commit.PRNumber] = append(prCommits[commit.PRNumber], commit)
		} else {
			directCommits = append(directCommits, commit)
		}
	}

	// There are occasionally no PRs or direct commits other than version bumps, so we handle that gracefully
	if len(prCommits) == 0 && len(directCommits) == 0 {
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
		// Sort direct commits by date (newest first) for consistent ordering
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

	// Fix MD032/blank-around-lists: Lists should be surrounded by blank lines
	lines := strings.Split(content, "\n")
	inList := false
	preListNewline := false
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			if !inList {
				inList = true
				// Ensure there's a blank line before the list starts
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
		lines := strings.Split(pr.Body, "\n")
		for _, line := range lines {
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
	// No hard tabs
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
	// Turn "*"" lists into "-" lists"
	message = strings.ReplaceAll(message, "* ", "- ")
	// Remove extra spaces around dashes
	message = strings.ReplaceAll(message, "-   ", "- ")
	message = strings.ReplaceAll(message, "-  ", "- ")
	// turn bare URL into <URL>
	if strings.Contains(message, "http://") || strings.Contains(message, "https://") {
		// Use regex to wrap bare URLs with angle brackets
		urlRegex := regexp.MustCompile(`\b(https?://[^\s<>]+)`)
		message = urlRegex.ReplaceAllString(message, "<$1>")
	}

	// Replace "## LINKS\n" with "- "
	message = strings.ReplaceAll(message, "## LINKS\n", "- ")
	// Dependabot messages: "- [Commits]" should become "\n- [Commits]"
	message = strings.TrimSpace(message)
	// Turn multiple newlines into a single newline
	message = strings.TrimSpace(strings.ReplaceAll(message, "\n\n", "\n"))
	// Fix inline trailing spaces
	message = strings.ReplaceAll(message, " \n", "\n")
	// Fix weird indent before list,
	message = strings.ReplaceAll(message, "\n - ", "\n- ")

	// blanks-around-lists MD032 fix
	// Use regex to ensure blank line before list items that don't already have one
	listRegex := regexp.MustCompile(`(?m)([^\n-].*[^:\n])\n([-*] .*)`)
	message = listRegex.ReplaceAllString(message, "$1\n\n$2")

	// Change random first-level "#" to 4th level "####"
	// This is a hack to fix spurious first-level headings that are not actual headings
	// but rather just comments or notes in the commit message.
	message = strings.ReplaceAll(message, "# ", "\n#### ")
	message = strings.ReplaceAll(message, "\n\n\n", "\n\n")

	// Wrap any non-wrapped Emails with angle brackets
	emailRegex := regexp.MustCompile(`([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
	message = emailRegex.ReplaceAllString(message, "<$1>")

	// Wrap any non-wrapped URLs with angle brackets
	urlRegex := regexp.MustCompile(`(https?://[^\s<]+)`)
	message = urlRegex.ReplaceAllString(message, "<$1>")

	message = strings.ReplaceAll(message, "<<", "<")
	message = strings.ReplaceAll(message, ">>", ">")

	// Fix some spurious Issue/PR links at the beginning of a commit message line
	prOrIssueLinkRegex := regexp.MustCompile("\n" + `(#\d+)`)
	message = prOrIssueLinkRegex.ReplaceAllString(message, " $1")

	// Remove leading/trailing whitespace
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

// hashContent generates a SHA256 hash of the content for change detection
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

	// Step 1: Force PR sync (pass true explicitly)
	fmt.Fprintf(os.Stderr, "[PR_SYNC] Forcing PR sync from GitHub...\n")
	if err := g.fetchPRs(true); err != nil {
		return fmt.Errorf("failed to sync PRs: %w", err)
	}

	// Step 2: Rebuild git history and verify versions/commits completeness
	fmt.Fprintf(os.Stderr, "[VERIFY] Verifying git history and version completeness...\n")
	if err := g.syncGitHistory(); err != nil {
		return fmt.Errorf("failed to sync git history: %w", err)
	}

	// Step 3: Verify commit-PR mappings
	fmt.Fprintf(os.Stderr, "[MAPPING] Verifying commit-PR mappings...\n")
	if err := g.verifyCommitPRMappings(); err != nil {
		return fmt.Errorf("failed to verify commit-PR mappings: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[SUCCESS] Database synchronization completed successfully!\n")
	return nil
}

// syncGitHistory walks the complete git history and ensures all versions and commits are cached
func (g *Generator) syncGitHistory() error {
	// Walk complete git history (reuse existing logic)
	versions, err := g.gitWalker.WalkHistory()
	if err != nil {
		return fmt.Errorf("failed to walk git history: %w", err)
	}

	// Save only new versions and commits (preserve existing data)
	var newVersions, newCommits int
	for _, version := range versions {
		// Only save version if it doesn't exist
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

		// Only save commits that don't exist
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

	// Update last processed tag
	if latestTag, err := g.gitWalker.GetLatestTag(); err == nil && latestTag != "" {
		if err := g.cache.SetLastProcessedTag(latestTag); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to update last processed tag: %v\n", err)
		}
	}

	fmt.Fprintf(os.Stderr, "   Added %d new versions and %d new commits (preserved existing data)\n", newVersions, newCommits)
	return nil
}

// verifyCommitPRMappings ensures all PR commits have proper mappings
func (g *Generator) verifyCommitPRMappings() error {
	// Get all cached PRs
	allPRs, err := g.cache.GetAllPRs()
	if err != nil {
		return fmt.Errorf("failed to get cached PRs: %w", err)
	}

	// Convert to slice for batch operations (reuse existing logic)
	var prSlice []*github.PR
	for _, pr := range allPRs {
		prSlice = append(prSlice, pr)
	}

	// Save commit-PR mappings (reuse existing logic)
	if err := g.cache.SaveCommitPRMappings(prSlice); err != nil {
		return fmt.Errorf("failed to save commit-PR mappings: %w", err)
	}

	fmt.Fprintf(os.Stderr, "   Verified mappings for %d PRs\n", len(prSlice))
	return nil
}
