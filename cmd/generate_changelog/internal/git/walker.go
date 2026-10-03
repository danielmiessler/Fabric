package git

import (
	"fmt"
	"os/exec"
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
	// versionPattern matches version commit messages, with or without the "chore(release): " prefix.
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
			latestTagName = tagRef.Name().Short()
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
	tagRef, err := w.repo.Tag(tagName)
	if err != nil {
		return nil, fmt.Errorf("failed to find tag %s: %w", tagName, err)
	}

	tagCommit, err := w.repo.CommitObject(tagRef.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get tag commit: %w", err)
	}

	headRef, err := w.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

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
		if c.Hash == tagCommit.Hash {
			return fmt.Errorf("reached tag commit") // sentinel that stops the iteration
		}

		commit := &Commit{
			SHA:     c.Hash.String(),
			Message: strings.TrimSpace(c.Message),
			Date:    c.Committer.When,
		}

		if versionMatch := versionPattern.FindStringSubmatch(commit.Message); versionMatch != nil {
			commit.IsVersion = true
		}

		if prMatch := prPattern.FindStringSubmatch(commit.Message); prMatch != nil {
			if prNumber, err := strconv.Atoi(prMatch[1]); err == nil {
				commit.PRNumber = prNumber
				prNumbers = append(prNumbers, prNumber)
			}
		}

		version.Commits = append(version.Commits, commit)
		return nil
	})

	// The sentinel error is not a failure.
	if err != nil && !strings.Contains(err.Error(), "reached tag commit") {
		return nil, fmt.Errorf("failed to walk commits: %w", err)
	}

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

	// Prefer upstream so a fork reports its parent repository.
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
	tagRef, err := w.repo.Tag(sinceTag)
	if err != nil {
		return nil, fmt.Errorf("failed to get tag %s: %w", sinceTag, err)
	}

	tagCommit, err := w.repo.CommitObject(tagRef.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get commit for tag %s: %w", sinceTag, err)
	}

	ref, err := w.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

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
			prNumber, err := strconv.Atoi(matches[1])
			if err != nil {
				return fmt.Errorf("failed to parse PR number: %v", err)
			}
			commit.PRNumber = prNumber

			prNumbers[currentVersion] = append(prNumbers[currentVersion], prNumber)
		}

		if _, exists := versions[currentVersion]; !exists {
			versions[currentVersion] = &Version{
				Name:      currentVersion,
				Date:      time.Time{},
				CommitSHA: "",
				Commits:   []*Commit{},
			}
		}

		versions[currentVersion].Commits = append(versions[currentVersion].Commits, commit)
		return nil
	})

	// storer.ErrStop is not a failure.
	if err == storer.ErrStop {
		err = nil
	}

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
// Uses native git CLI instead of go-git to properly handle worktree scenarios
func (w *Walker) IsWorkingDirectoryClean() (bool, error) {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("failed to get worktree: %w", err)
	}

	worktreePath := worktree.Filesystem.Root()

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = worktreePath

	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to get git status: %w", err)
	}

	return len(strings.TrimSpace(string(output))) == 0, nil
}

// GetStatusDetails returns a detailed status of the working directory
// Uses native git CLI instead of go-git to properly handle worktree scenarios
func (w *Walker) GetStatusDetails() (string, error) {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("failed to get worktree: %w", err)
	}

	worktreePath := worktree.Filesystem.Root()

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = worktreePath

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git status: %w", err)
	}

	return string(output), nil
}

// AddFile adds a file to the git index
// Uses native git CLI instead of go-git to properly handle worktree scenarios
func (w *Walker) AddFile(filename string) error {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	worktreePath := worktree.Filesystem.Root()

	cmd := exec.Command("git", "add", filename)
	cmd.Dir = worktreePath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add file %s: %w (output: %s)", filename, err, string(output))
	}

	return nil
}

// CommitChanges creates a commit with the given message
// Uses native git CLI instead of go-git to properly handle worktree scenarios
func (w *Walker) CommitChanges(message string) (plumbing.Hash, error) {
	worktree, err := w.repo.Worktree()
	if err != nil {
		return plumbing.ZeroHash, fmt.Errorf("failed to get worktree: %w", err)
	}

	worktreePath := worktree.Filesystem.Root()

	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Dir = worktreePath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return plumbing.ZeroHash, fmt.Errorf("failed to commit: %w (output: %s)", err, string(output))
	}

	hashCmd := exec.Command("git", "rev-parse", "HEAD")
	hashCmd.Dir = worktreePath

	hashOutput, err := hashCmd.Output()
	if err != nil {
		return plumbing.ZeroHash, fmt.Errorf("failed to get HEAD after commit: %w", err)
	}

	hashStr := strings.TrimSpace(string(hashOutput))
	return plumbing.NewHash(hashStr), nil
}

// PushToRemote pushes the current branch to the remote repository
// It automatically detects GitHub repositories and uses token authentication when available
func (w *Walker) PushToRemote() error {
	pushOptions := &git.PushOptions{}

	if githubToken := util.GetTokenFromEnv(""); githubToken != "" {
		remotes, err := w.repo.Remotes()
		if err == nil && len(remotes) > 0 {
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

			urls := remote.Config().URLs
			if len(urls) > 0 {
				url := urls[0]
				if strings.Contains(url, "github.com") {
					pushOptions.Auth = &http.BasicAuth{
						Username: "token", // any non-empty username works
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
