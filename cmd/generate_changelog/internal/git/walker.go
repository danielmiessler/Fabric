package git

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

var (
	// The versionPattern matches version commit messages with or without the optional "chore(release): " prefix.
	// Examples of matching commit messages:
	//   - "chore(release): Update version to v1.2.3"
	//   - "Update version to v1.2.3"
	// Examples of non-matching commit messages:
	//   - "fix: Update version to v1.2.3" (missing "chore(release): " or "Update version to")
	//   - "chore(release): Update version to 1.2.3" (missing "v" prefix in version)
	//   - "Update version to v1.2" (incomplete version number)
	versionPattern = regexp.MustCompile(`(?:chore\(release\): )?Update version to (v\d+\.\d+\.\d+)`)
	prPattern      = regexp.MustCompile(`Merge pull request #(\d+)`)
)

type Walker struct {
	repo *git.Repository
}

func NewWalker(repoPath string) (*Walker, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	return &Walker{repo: repo}, nil
}

// GetLatestTag returns the name of the most recent tag by committer date
func (w *Walker) GetLatestTag() (string, error) {
	tagRefs, err := w.repo.Tags()
	if err != nil {
		return "", err
	}

	var latestTagCommit *object.Commit
	var latestTagName string

	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		revision := plumbing.Revision(tagRef.Name().String())
		tagCommitHash, err := w.repo.ResolveRevision(revision)
		if err != nil {
			return err
		}

		commit, err := w.repo.CommitObject(*tagCommitHash)
		if err != nil {
			return err
		}

		if latestTagCommit == nil {
			latestTagCommit = commit
			latestTagName = tagRef.Name().Short() // Get short name like "v1.4.245"
		}

		if commit.Committer.When.After(latestTagCommit.Committer.When) {
			latestTagCommit = commit
			latestTagName = tagRef.Name().Short()
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return latestTagName, nil
}

// WalkCommitsSinceTag walks commits from the specified tag to HEAD and returns only "Unreleased" version
func (w *Walker) WalkCommitsSinceTag(tagName string) (*Version, error) {
	// Get the tag reference
	tagRef, err := w.repo.Tag(tagName)
	if err != nil {
		return nil, fmt.Errorf("failed to find tag %s: %w", tagName, err)
	}

	// Get the commit that the tag points to
	tagCommit, err := w.repo.CommitObject(tagRef.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get tag commit: %w", err)
	}

	// Get HEAD
	headRef, err := w.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	// Walk from HEAD back to the tag commit (exclusive)
	commitIter, err := w.repo.Log(&git.LogOptions{
		From:  headRef.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get commit log: %w", err)
	}

	version := &Version{
		Name:    "Unreleased",
		Commits: []*Commit{},
	}

	prNumbers := []int{}

	err = commitIter.ForEach(func(c *object.Commit) error {
		// Stop when we reach the tag commit (don't include it)
		if c.Hash == tagCommit.Hash {
			return fmt.Errorf("reached tag commit") // Use error to break out of iteration
		}

		commit := &Commit{
			SHA:     c.Hash.String(),
			Message: strings.TrimSpace(c.Message),
			Date:    c.Committer.When,
		}

		// Check for version patterns
		if versionMatch := versionPattern.FindStringSubmatch(commit.Message); versionMatch != nil {
			commit.IsVersion = true
		}

		// Check for PR merge patterns
		if prMatch := prPattern.FindStringSubmatch(commit.Message); prMatch != nil {
			if prNumber, err := strconv.Atoi(prMatch[1]); err == nil {
				commit.PRNumber = prNumber
				prNumbers = append(prNumbers, prNumber)
			}
		}

		version.Commits = append(version.Commits, commit)
		return nil
	})

	// Ignore the "reached tag commit" error - it's expected
	if err != nil && !strings.Contains(err.Error(), "reached tag commit") {
		return nil, fmt.Errorf("failed to walk commits: %w", err)
	}

	// Remove duplicates from prNumbers and set them
	prNumbersMap := make(map[int]bool)
	for _, prNum := range prNumbers {
		prNumbersMap[prNum] = true
	}

	version.PRNumbers = make([]int, 0, len(prNumbersMap))
	for prNum := range prNumbersMap {
		version.PRNumbers = append(version.PRNumbers, prNum)
	}

	return version, nil
}

func (w *Walker) WalkHistory() (map[string]*Version, error) {
	ref, err := w.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	commitIter, err := w.repo.Log(&git.LogOptions{
		From:  ref.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get commit log: %w", err)
	}

	versions := make(map[string]*Version)
	currentVersion := "Unreleased"
	versions[currentVersion] = &Version{
		Name:    currentVersion,
		Commits: []*Commit{},
	}

	prNumbers := make(map[string][]int)

	err = commitIter.ForEach(func(c *object.Commit) error {
		// c.Message = Summarize(c.Message)
		commit := &Commit{
			SHA:     c.Hash.String(),
			Message: strings.TrimSpace(c.Message),
			Author:  c.Author.Name,
			Email:   c.Author.Email,
			Date:    c.Author.When,
			IsMerge: len(c.ParentHashes) > 1,
		}

		if matches := versionPattern.FindStringSubmatch(commit.Message); len(matches) > 1 {
			commit.IsVersion = true
			commit.Version = matches[1]
			currentVersion = commit.Version

			if _, exists := versions[currentVersion]; !exists {
				versions[currentVersion] = &Version{
					Name:      currentVersion,
					Date:      commit.Date,
					CommitSHA: commit.SHA,
					Commits:   []*Commit{},
				}
			}
			return nil
		}

		if matches := prPattern.FindStringSubmatch(commit.Message); len(matches) > 1 {
			prNumber := 0
			fmt.Sscanf(matches[1], "%d", &prNumber)
			commit.PRNumber = prNumber

			prNumbers[currentVersion] = append(prNumbers[currentVersion], prNumber)
		}

		versions[currentVersion].Commits = append(versions[currentVersion].Commits, commit)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk commits: %w", err)
	}

	for version, prs := range prNumbers {
		versions[version].PRNumbers = dedupInts(prs)
	}

	return versions, nil
}

func (w *Walker) GetRepoInfo() (owner string, name string, err error) {
	remotes, err := w.repo.Remotes()
	if err != nil {
		return "", "", fmt.Errorf("failed to get remotes: %w", err)
	}

	// First try upstream (preferred for forks)
	for _, remote := range remotes {
		if remote.Config().Name == "upstream" {
			urls := remote.Config().URLs
			if len(urls) > 0 {
				owner, name = parseGitHubURL(urls[0])
				if owner != "" && name != "" {
					return owner, name, nil
				}
			}
		}
	}

	// Then try origin
	for _, remote := range remotes {
		if remote.Config().Name == "origin" {
			urls := remote.Config().URLs
			if len(urls) > 0 {
				owner, name = parseGitHubURL(urls[0])
				if owner != "" && name != "" {
					return owner, name, nil
				}
			}
		}
	}

	return "danielmiessler", "fabric", nil
}

func parseGitHubURL(url string) (owner, repo string) {
	patterns := []string{
		`github\.com[:/]([^/]+)/([^/.]+)`,
		`github\.com[:/]([^/]+)/([^/]+)\.git$`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(url)
		if len(matches) > 2 {
			return matches[1], matches[2]
		}
	}

	return "", ""
}

// WalkHistorySinceTag walks git history from HEAD down to (but not including) the specified tag
// and returns any version commits found along the way
func (w *Walker) WalkHistorySinceTag(sinceTag string) (map[string]*Version, error) {
	// Get the commit SHA for the sinceTag
	tagRef, err := w.repo.Tag(sinceTag)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag %s: %w", sinceTag, err)
	}

	tagCommit, err := w.repo.CommitObject(tagRef.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get commit for tag %s: %w", sinceTag, err)
	}

	// Get HEAD reference
	ref, err := w.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	// Walk from HEAD down to the tag commit (excluding it)
	commitIter, err := w.repo.Log(&git.LogOptions{
		From:  ref.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create commit iterator: %w", err)
	}
	defer commitIter.Close()

	versions := make(map[string]*Version)
	currentVersion := "Unreleased"
	prNumbers := make(map[string][]int)

	err = commitIter.ForEach(func(c *object.Commit) error {
		// Stop iteration when the hash of the current commit matches the hash of the specified sinceTag commit
		if c.Hash == tagCommit.Hash {
			return storer.ErrStop
		}

		commit := &Commit{
			SHA:     c.Hash.String(),
			Message: strings.TrimSpace(c.Message),
			Author:  c.Author.Name,
			Email:   c.Author.Email,
			Date:    c.Author.When,
			IsMerge: len(c.ParentHashes) > 1,
		}

		// Check for version pattern
		if matches := versionPattern.FindStringSubmatch(commit.Message); len(matches) > 1 {
			commit.IsVersion = true
			commit.Version = matches[1]
			currentVersion = commit.Version

			if _, exists := versions[currentVersion]; !exists {
				versions[currentVersion] = &Version{
					Name:      currentVersion,
					Date:      commit.Date,
					CommitSHA: commit.SHA,
					Commits:   []*Commit{},
				}
			}
			return nil
		}

		// Check for PR merge pattern
		if matches := prPattern.FindStringSubmatch(commit.Message); len(matches) > 1 {
			prNumber, err := strconv.Atoi(matches[1])
			if err != nil {
				// Handle parsing error (e.g., log it or skip processing)
				return fmt.Errorf("failed to parse PR number: %v", err)
			}
			commit.PRNumber = prNumber

			prNumbers[currentVersion] = append(prNumbers[currentVersion], prNumber)
		}

		// Add commit to current version
		if _, exists := versions[currentVersion]; !exists {
			versions[currentVersion] = &Version{
				Name:      currentVersion,
				Date:      time.Time{}, // Zero value, will be set by version commit
				CommitSHA: "",
				Commits:   []*Commit{},
			}
		}

		versions[currentVersion].Commits = append(versions[currentVersion].Commits, commit)
		return nil
	})

	// Handle the stop condition - storer.ErrStop is expected
	if err == storer.ErrStop {
		err = nil
	}

	// Assign collected PR numbers to each version
	for version, prs := range prNumbers {
		versions[version].PRNumbers = dedupInts(prs)
	}

	return versions, err
}

func dedupInts(ints []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, i := range ints {
		if !seen[i] {
			seen[i] = true
			result = append(result, i)
		}
	}

	return result
}

// Worktree returns the git worktree for performing git operations
func (w *Walker) Worktree() (*git.Worktree, error) {
	return w.repo.Worktree()
}

// Repository returns the underlying git repository
func (w *Walker) Repository() *git.Repository {
	return w.repo
}

// IsWorkingDirectoryClean checks if the working directory has any uncommitted changes
func (w *Walker) IsWorkingDirectoryClean() (bool, error) {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return false, fmt.Errorf("failed to get git status: %w", err)
	}

	return status.IsClean(), nil
}

// GetStatusDetails returns a detailed status of the working directory
func (w *Walker) GetStatusDetails() (string, error) {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return "", fmt.Errorf("failed to get git status: %w", err)
	}

	if status.IsClean() {
		return "", nil
	}

	var details strings.Builder
	for file, fileStatus := range status {
		details.WriteString(fmt.Sprintf("  %c%c %s\n", fileStatus.Staging, fileStatus.Worktree, file))
	}

	return details.String(), nil
}

// AddFile adds a file to the git index
func (w *Walker) AddFile(filename string) error {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = worktree.Add(filename)
	if err != nil {
		return fmt.Errorf("failed to add file %s: %w", filename, err)
	}

	return nil
}

// CommitChanges creates a commit with the given message
func (w *Walker) CommitChanges(message string) (plumbing.Hash, error) {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return plumbing.ZeroHash, fmt.Errorf("failed to get worktree: %w", err)
	}

	// Get git config for author information
	cfg, err := w.repo.Config()
	if err != nil {
		return plumbing.ZeroHash, fmt.Errorf("failed to get git config: %w", err)
	}

	var authorName, authorEmail string
	if cfg.User.Name != "" {
		authorName = cfg.User.Name
	} else {
		authorName = "Changelog Bot"
	}
	if cfg.User.Email != "" {
		authorEmail = cfg.User.Email
	} else {
		authorEmail = "bot@changelog.local"
	}

	commit, err := worktree.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  authorName,
			Email: authorEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		return plumbing.ZeroHash, fmt.Errorf("failed to commit: %w", err)
	}

	return commit, nil
}

// PushToRemote pushes the current branch to the remote repository
// It automatically detects GitHub repositories and uses token authentication when available
func (w *Walker) PushToRemote() error {
	pushOptions := &git.PushOptions{}

	// Check if we have a GitHub token for authentication
	if githubToken := util.GetTokenFromEnv(""); githubToken != "" {
		// Get remote URL to check if it's a GitHub repository
		remotes, err := w.repo.Remotes()
		if err == nil && len(remotes) > 0 {
			// Get the origin remote (or first remote if origin doesn't exist)
			var remote *git.Remote
			for _, r := range remotes {
				if r.Config().Name == "origin" {
					remote = r
					break
				}
			}
			if remote == nil {
				remote = remotes[0]
			}

			// Check if this is a GitHub repository
			urls := remote.Config().URLs
			if len(urls) > 0 {
				url := urls[0]
				if strings.Contains(url, "github.com") {
					// Use token authentication for GitHub repositories
					pushOptions.Auth = &http.BasicAuth{
						Username: "token", // GitHub expects "token" as username
						Password: githubToken,
					}
				}
			}
		}
	}

	err := w.repo.Push(pushOptions)
	if err != nil {
		return fmt.Errorf("failed to push: %w", err)
	}
	return nil
}

// RemoveFile removes a file from both the working directory and git index
func (w *Walker) RemoveFile(filename string) error {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = worktree.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to remove file %s: %w", filename, err)
	}

	return nil
}
