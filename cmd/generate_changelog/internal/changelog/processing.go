package changelog

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/git"
	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/github"
)

var (
	mergePatterns     []*regexp.Regexp
	mergePatternsOnce sync.Once
)

func getMergePatterns() []*regexp.Regexp {
	mergePatternsOnce.Do(func() {
		mergePatterns = []*regexp.Regexp{
			regexp.MustCompile(`^Merge pull request #\d+`),      // "Merge pull request #123 from..."
			regexp.MustCompile(`^Merge branch '.*' into .*`),    // "Merge branch 'feature' into main"
			regexp.MustCompile(`^Merge remote-tracking branch`), // "Merge remote-tracking branch..."
			regexp.MustCompile(`^Merge '.*' into .*`),           // "Merge 'feature' into main"
		}
	})
	return mergePatterns
}

// isMergeCommit reports whether a commit has more than one parent or a merge-style message.
func isMergeCommit(commit github.PRCommit) bool {
	if len(commit.Parents) > 1 {
		return true
	}

	// Commits fetched with GraphQL have no parent list, so also match the message.
	mergePatterns := getMergePatterns()
	for _, pattern := range mergePatterns {
		if pattern.MatchString(commit.Message) {
			return true
		}
	}

	return false
}

// calculateVersionDate returns the newest commit date across the PRs, or the current time when none exists.
func calculateVersionDate(fetchedPRs []*github.PR) time.Time {
	versionDate := time.Now()
	if len(fetchedPRs) > 0 {
		var mostRecentCommitDate time.Time
		for _, pr := range fetchedPRs {
			for _, commit := range pr.Commits {
				if commit.Date.After(mostRecentCommitDate) {
					mostRecentCommitDate = commit.Date
				}
			}
		}
		if !mostRecentCommitDate.IsZero() {
			versionDate = mostRecentCommitDate
		}
	}
	return versionDate
}

// ProcessIncomingPR processes a single PR for changelog entry creation
func (g *Generator) ProcessIncomingPR(prNumber int) error {
	if err := g.validatePRState(prNumber); err != nil {
		return fmt.Errorf("PR validation failed: %w", err)
	}

	if err := g.validateGitStatus(); err != nil {
		return fmt.Errorf("git status validation failed: %w", err)
	}

	// The validation call above did not fetch commits.
	pr, err := g.ghClient.GetPRWithCommits(prNumber)
	if err != nil {
		return fmt.Errorf("failed to fetch PR %d: %w", prNumber, err)
	}

	content := g.formatPR(pr)

	if g.cfg.EnableAISummary {
		aiContent, err := SummarizeVersionContent(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: AI summarization failed: %v\n", err)
		} else if !checkForAIError(aiContent) {
			content = strings.TrimSpace(aiContent)
		}
	}

	if err := g.ensureIncomingDir(); err != nil {
		return fmt.Errorf("failed to create incoming directory: %w", err)
	}

	filename := filepath.Join(g.cfg.IncomingDir, fmt.Sprintf("%d.txt", prNumber))

	content = strings.TrimSpace(content) + "\n"

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write incoming file: %w", err)
	}

	if err := g.commitAndPushIncoming(prNumber, filename); err != nil {
		return fmt.Errorf("failed to commit and push: %w", err)
	}

	fmt.Printf("Successfully created incoming changelog entry: %s\n", filename)
	return nil
}

// CreateNewChangelogEntry aggregates all incoming PR files for release and includes direct commits
func (g *Generator) CreateNewChangelogEntry(version string) error {
	files, err := filepath.Glob(filepath.Join(g.cfg.IncomingDir, "*.txt"))
	if err != nil {
		return fmt.Errorf("failed to scan incoming directory: %w", err)
	}

	var content strings.Builder
	var processingErrors []string

	for i, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			processingErrors = append(processingErrors, fmt.Sprintf("failed to read %s: %v", file, err))
			continue
		}
		content.WriteString(string(data))
		if i < len(files)-1 {
			content.WriteString("\n")
		}
	}

	if len(processingErrors) > 0 {
		return fmt.Errorf("encountered errors while processing incoming files: %s", strings.Join(processingErrors, "; "))
	}

	// Collect the PR numbers and commit SHAs of the incoming files. The direct-commit scan below excludes them.
	processedPRs := make(map[int]bool)
	processedCommitSHAs := make(map[string]bool)
	var fetchedPRs []*github.PR
	var prNumbers []int

	for _, file := range files {
		filename := filepath.Base(file)
		if prNumStr, ok := strings.CutSuffix(filename, ".txt"); ok {
			if prNum, err := strconv.Atoi(prNumStr); err == nil {
				processedPRs[prNum] = true
				prNumbers = append(prNumbers, prNum)

				if pr, err := g.ghClient.GetPRWithCommits(prNum); err == nil {
					fetchedPRs = append(fetchedPRs, pr)
					for _, commit := range pr.Commits {
						processedCommitSHAs[commit.SHA] = true
					}
				}
			}
		}
	}

	directCommitsContent, err := g.getDirectCommitsSinceLastRelease(processedPRs, processedCommitSHAs)
	if err != nil {
		return fmt.Errorf("failed to get direct commits since last release: %w", err)
	}
	if directCommitsContent != "" {
		if content.Len() > 0 {
			content.WriteString("\n")
		}
		content.WriteString(directCommitsContent)
	}

	if content.Len() == 0 {
		if len(files) == 0 {
			fmt.Fprintf(os.Stderr, "No incoming PR files found in %s and no direct commits since last release\n", g.cfg.IncomingDir)
		} else {
			fmt.Fprintf(os.Stderr, "No content found in incoming files and no direct commits since last release\n")
		}
		return nil
	}

	versionDate := calculateVersionDate(fetchedPRs)

	entry := fmt.Sprintf("## %s (%s)\n\n%s",
		version, versionDate.Format("2006-01-02"), strings.TrimLeft(content.String(), "\n"))

	if err := g.insertVersionAtTop(entry); err != nil {
		return fmt.Errorf("failed to update CHANGELOG.md: %w", err)
	}

	if g.cache != nil {
		if len(fetchedPRs) > 0 {
			if err := g.cache.SavePRBatch(fetchedPRs); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to save PR batch to cache: %v\n", err)
			}

			if err := g.cache.SaveCommitPRMappings(fetchedPRs); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to cache commit mappings: %v\n", err)
			}

			for _, pr := range fetchedPRs {
				for _, commit := range pr.Commits {
					commitDate := commit.Date
					if commitDate.IsZero() {
						commitDate = time.Now()
						fmt.Fprintf(os.Stderr, "Warning: Commit %s has invalid timestamp, using current time as fallback\n", commit.SHA)
					}

					gitCommit := &git.Commit{
						SHA:      commit.SHA,
						Message:  commit.Message,
						Author:   commit.Author,
						Email:    commit.Email,
						Date:     commitDate,
						IsMerge:  isMergeCommit(commit),
						PRNumber: pr.Number,
					}
					if err := g.cache.SaveCommit(gitCommit, version); err != nil {
						fmt.Fprintf(os.Stderr, "Warning: Failed to save commit %s to cache: %v\n", commit.SHA, err)
					}
				}
			}
		}

		newVersionEntry := &git.Version{
			Name:      version,
			Date:      versionDate,
			CommitSHA: "", // the release commit does not exist yet
			PRNumbers: prNumbers,
			AISummary: content.String(),
		}

		if err := g.cache.SaveVersion(newVersionEntry); err != nil {
			return fmt.Errorf("failed to save new version entry to database: %w", err)
		}
	}

	for _, file := range files {
		relativeFile, err := filepath.Rel(g.cfg.RepoPath, file)
		if err != nil {
			relativeFile = file
		}

		// Git removal updates both the filesystem and the index.
		if err := g.gitWalker.RemoveFile(relativeFile); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to remove %s from git index: %v\n", relativeFile, err)
			if err := os.Remove(file); err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to remove %s from the filesystem after failing to remove it from the git index.\n", relativeFile)
				fmt.Fprintf(os.Stderr, "Filesystem error: %v\n", err)
				fmt.Fprintf(os.Stderr, "Manual intervention required:\n")
				fmt.Fprintf(os.Stderr, "  1. Remove the file %s manually (using the OS-specific command)\n", file)
				fmt.Fprintf(os.Stderr, "  2. Remove from git index: git rm --cached %s\n", relativeFile)
				fmt.Fprintf(os.Stderr, "  3. Or reset git index: git reset HEAD %s\n", relativeFile)
			}
		}
	}

	// Update the cache metadata before staging, so the release commit includes it.
	if g.cache != nil {
		if err := g.cache.SetLastProcessedTag(version); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to update last_processed_tag: %v\n", err)
		}

		// Set last_pr_sync to the version date, not the current time.
		// Later runs then fetch the PRs merged after this version.
		if err := g.cache.SetLastPRSync(versionDate); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to update last_pr_sync: %v\n", err)
		}
	}

	if err := g.stageChangesForRelease(); err != nil {
		return fmt.Errorf("critical: failed to stage changes for release: %w", err)
	}

	fmt.Printf("Successfully processed %d incoming PR files for version %s\n", len(files), version)
	return nil
}

// getDirectCommitsSinceLastRelease formats the commits since the latest tag that belong to no PR.
func (g *Generator) getDirectCommitsSinceLastRelease(processedPRs map[int]bool, processedCommitSHAs map[string]bool) (string, error) {
	latestTag, err := g.gitWalker.GetLatestTag()
	if err != nil {
		return "", fmt.Errorf("failed to get latest tag: %w", err)
	}

	unreleasedVersion, err := g.gitWalker.WalkCommitsSinceTag(latestTag)
	if err != nil {
		return "", fmt.Errorf("failed to walk commits since tag %s: %w", latestTag, err)
	}

	if unreleasedVersion == nil || len(unreleasedVersion.Commits) == 0 {
		return "", nil
	}

	var directCommits []*git.Commit
	for _, commit := range unreleasedVersion.Commits {
		if commit.IsVersion {
			continue
		}

		if commit.PRNumber > 0 && processedPRs[commit.PRNumber] {
			continue
		}

		// Only merge commits carry a PR number. The SHA set catches the other commits of a PR.
		if processedCommitSHAs[commit.SHA] {
			continue
		}

		if commit.PRNumber == 0 {
			directCommits = append(directCommits, commit)
		}
	}

	if len(directCommits) == 0 {
		return "", nil
	}

	var sb strings.Builder
	sb.WriteString("### Direct commits\n\n")

	sort.Slice(directCommits, func(i, j int) bool {
		return directCommits[i].Date.After(directCommits[j].Date)
	})

	for _, commit := range directCommits {
		message := g.formatCommitMessage(strings.TrimSpace(commit.Message))
		if message != "" && !g.isDuplicateMessage(message, directCommits) {
			sb.WriteString(fmt.Sprintf("- %s\n", message))
		}
	}

	return sb.String(), nil
}

// validatePRState rejects a PR that is not open unless --closed-ok is set, and an open PR that is not mergeable.
func (g *Generator) validatePRState(prNumber int) error {
	details, err := g.ghClient.GetPRValidationDetails(prNumber)
	if err != nil {
		return fmt.Errorf("failed to fetch PR %d: %w", prNumber, err)
	}

	if details.State != "open" && !g.cfg.ClosedOK {
		return fmt.Errorf("PR %d is not open (current state: %s); use --closed-ok to process it anyway", prNumber, details.State)
	}

	// GitHub returns a nil mergeable flag for closed and merged PRs. Check it only for open PRs.
	if details.State == "open" && !details.Mergeable {
		return fmt.Errorf("PR %d is not mergeable - please resolve conflicts first", prNumber)
	}

	return nil
}

func (g *Generator) validateGitStatus() error {
	isClean, err := g.gitWalker.IsWorkingDirectoryClean()
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if !isClean {
		statusDetails, statusErr := g.gitWalker.GetStatusDetails()
		if statusErr == nil && statusDetails != "" {
			return fmt.Errorf("working directory is not clean - please commit or stash changes before proceeding:\n%s", statusDetails)
		}
		return fmt.Errorf("working directory is not clean - please commit or stash changes before proceeding")
	}

	return nil
}

func (g *Generator) ensureIncomingDir() error {
	if err := os.MkdirAll(g.cfg.IncomingDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", g.cfg.IncomingDir, err)
	}
	return nil
}

// commitAndPushIncoming commits the incoming file and pushes when --push is set.
func (g *Generator) commitAndPushIncoming(prNumber int, filename string) error {
	relativeFilename, err := filepath.Rel(g.cfg.RepoPath, filename)
	if err != nil {
		relativeFilename = filename
	}

	if err := g.gitWalker.AddFile(relativeFilename); err != nil {
		return fmt.Errorf("failed to add file %s: %w", relativeFilename, err)
	}

	commitMessage := fmt.Sprintf("chore: incoming %d changelog entry", prNumber)
	_, err = g.gitWalker.CommitChanges(commitMessage)
	if err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	if g.cfg.Push {
		if err := g.gitWalker.PushToRemote(); err != nil {
			return fmt.Errorf("failed to push to remote: %w", err)
		}
	} else {
		fmt.Println("Commit created successfully. Please review and push manually.")
	}

	return nil
}

// detectVersion reads version.nix, then falls back to the latest tag, then to "v1.0.0".
func (g *Generator) detectVersion() (string, error) {
	versionNixPath := filepath.Join(g.cfg.RepoPath, "version.nix")
	if _, err := os.Stat(versionNixPath); err == nil {
		data, err := os.ReadFile(versionNixPath)
		if err != nil {
			return "", fmt.Errorf("failed to read version.nix: %w", err)
		}

		versionRegex := regexp.MustCompile(`"([^"]+)"`)
		matches := versionRegex.FindStringSubmatch(string(data))
		if len(matches) > 1 {
			return matches[1], nil
		}
	}

	latestTag, err := g.gitWalker.GetLatestTag()
	if err != nil {
		return "", fmt.Errorf("failed to get latest tag: %w", err)
	}

	if latestTag == "" {
		return "v1.0.0", nil
	}

	return latestTag, nil
}

// insertVersionAtTop writes the entry below the "# Changelog" header and creates the file when absent.
func (g *Generator) insertVersionAtTop(entry string) error {
	changelogPath := filepath.Join(g.cfg.RepoPath, "CHANGELOG.md")
	header := "# Changelog"
	headerRegex := regexp.MustCompile(`(?m)^# Changelog\s*`)

	existingContent, err := os.ReadFile(changelogPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to read existing CHANGELOG.md: %w", err)
		}
		newContent := fmt.Sprintf("%s\n\n%s\n", header, entry)
		return os.WriteFile(changelogPath, []byte(newContent), 0644)
	}

	contentStr := string(existingContent)
	var newContent string

	if loc := headerRegex.FindStringIndex(contentStr); loc != nil {
		insertionPoint := loc[1]
		// Skip blank lines after the header to avoid double spacing.
		for insertionPoint < len(contentStr) && (contentStr[insertionPoint] == '\n' || contentStr[insertionPoint] == '\r') {
			insertionPoint++
		}
		newContent = contentStr[:loc[1]] + entry + "\n" + contentStr[insertionPoint:]
	} else {
		newContent = fmt.Sprintf("%s\n\n%s\n\n%s", header, entry, contentStr)
	}

	return os.WriteFile(changelogPath, []byte(newContent), 0644)
}

// stageChangesForRelease adds CHANGELOG.md and the cache database to the git index.
func (g *Generator) stageChangesForRelease() error {
	changelogPath := filepath.Join(g.cfg.RepoPath, "CHANGELOG.md")
	relativeChangelog, err := filepath.Rel(g.cfg.RepoPath, changelogPath)
	if err != nil {
		relativeChangelog = "CHANGELOG.md"
	}

	relativeCacheFile, err := filepath.Rel(g.cfg.RepoPath, g.cfg.CacheFile)
	if err != nil {
		relativeCacheFile = g.cfg.CacheFile
	}

	if err := g.gitWalker.AddFile(relativeChangelog); err != nil {
		return fmt.Errorf("failed to add %s: %w", relativeChangelog, err)
	}

	if err := g.gitWalker.AddFile(relativeCacheFile); err != nil {
		return fmt.Errorf("failed to add %s: %w", relativeCacheFile, err)
	}

	return nil
}
