package githelper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

// FetchOptions defines options for fetching files from a git repo
type FetchOptions struct {
	// RepoURL is the URL of the git repository
	RepoURL string

	// PathPrefix is the folder within the repo to extract (e.g. "patterns/")
	PathPrefix string

	// DestDir is where the files will be saved locally
	DestDir string

	// SingleDirectory if true, only fetch files directly in the specified directory
	// without recursing into subdirectories
	SingleDirectory bool
}

// FetchFilesFromRepo clones a git repo and extracts files from a specific folder
func FetchFilesFromRepo(opts FetchOptions) error {
	// Ensure path prefix ends with slash
	if !strings.HasSuffix(opts.PathPrefix, "/") {
		opts.PathPrefix = opts.PathPrefix + "/"
	}

	// Clone the repository in memory
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:   opts.RepoURL,
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	// Get HEAD reference
	ref, err := r.Head()
	if err != nil {
		return fmt.Errorf("failed to get repository HEAD: %w", err)
	}

	// Get commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return fmt.Errorf("failed to get commit: %w", err)
	}

	// Get the file tree
	tree, err := commit.Tree()
	if err != nil {
		return fmt.Errorf("failed to get tree: %w", err)
	}

	// Ensure destination directory exists
	if err := os.MkdirAll(opts.DestDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Extract files from the tree
	return tree.Files().ForEach(func(f *object.File) error {
		// Only process files in the specified path
		if !strings.HasPrefix(f.Name, opts.PathPrefix) {
			return nil
		}

		// For SingleDirectory mode, skip files in subdirectories
		if opts.SingleDirectory {
			remainingPath := strings.TrimPrefix(f.Name, opts.PathPrefix)
			if strings.Contains(remainingPath, "/") {
				return nil
			}
		}

		// Create local path for the file, removing the prefix
		relativePath := strings.TrimPrefix(f.Name, opts.PathPrefix)
		localPath := filepath.Join(opts.DestDir, relativePath)

		// Ensure directory structure exists
		if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
			return err
		}

		// Get file contents
		reader, err := f.Reader()
		if err != nil {
			return err
		}
		defer reader.Close()

		// Create and write to local file
		file, err := os.Create(localPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, reader)
		return err
	})
}
