package githelper

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
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

// FetchFilesFromRepo clones a git repo and extracts files from a specific folder.
// It tries go-git first, and falls back to the git CLI if available.
func FetchFilesFromRepo(opts FetchOptions) error {
	// Ensure path prefix ends with slash
	if !strings.HasSuffix(opts.PathPrefix, "/") {
		opts.PathPrefix = opts.PathPrefix + "/"
	}

	// Try go-git first (in-memory clone)
	goGitErr := fetchFilesViaGoGit(opts)
	if goGitErr == nil {
		return nil
	}

	// go-git failed; try git CLI fallback if available
	if _, lookErr := exec.LookPath("git"); lookErr != nil {
		return goGitErr
	}

	cliErr := fetchFilesViaGitCLI(opts)
	if cliErr == nil {
		return nil
	}

	return fmt.Errorf("%w; git CLI fallback also failed: %v", goGitErr, cliErr)
}

// fetchFilesViaGoGit clones a repo in memory using go-git and extracts files.
func fetchFilesViaGoGit(opts FetchOptions) error {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:   opts.RepoURL,
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf(i18n.T("githelper_failed_clone_repository"), err)
	}

	ref, err := r.Head()
	if err != nil {
		return fmt.Errorf(i18n.T("githelper_failed_get_head"), err)
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return fmt.Errorf(i18n.T("githelper_failed_get_commit"), err)
	}

	tree, err := commit.Tree()
	if err != nil {
		return fmt.Errorf(i18n.T("githelper_failed_get_tree"), err)
	}

	if err := os.MkdirAll(opts.DestDir, 0755); err != nil {
		return fmt.Errorf(i18n.T("githelper_failed_create_dest_directory"), err)
	}

	return tree.Files().ForEach(func(f *object.File) error {
		if !strings.HasPrefix(f.Name, opts.PathPrefix) {
			return nil
		}

		if opts.SingleDirectory {
			remainingPath := strings.TrimPrefix(f.Name, opts.PathPrefix)
			if strings.Contains(remainingPath, "/") {
				return nil
			}
		}

		relativePath := strings.TrimPrefix(f.Name, opts.PathPrefix)
		localPath := filepath.Join(opts.DestDir, relativePath)

		if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
			return err
		}

		reader, err := f.Reader()
		if err != nil {
			return err
		}
		defer reader.Close()

		file, err := os.Create(localPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, reader)
		return err
	})
}

// fetchFilesViaGitCLI clones a repo using the git CLI binary and extracts files.
// This serves as a fallback when go-git fails (e.g., DNS resolution issues on Termux).
func fetchFilesViaGitCLI(opts FetchOptions) error {
	tmpDir, err := os.MkdirTemp("", "fabric-git-clone-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	cmd := exec.Command("git", "clone", "--depth", "1", opts.RepoURL, tmpDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git clone failed: %w: %s", err, string(output))
	}

	// Source directory within the clone (trim trailing slash for filepath.Join)
	srcDir := filepath.Join(tmpDir, strings.TrimSuffix(opts.PathPrefix, "/"))

	if err := os.MkdirAll(opts.DestDir, 0755); err != nil {
		return fmt.Errorf(i18n.T("githelper_failed_create_dest_directory"), err)
	}

	return filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		if opts.SingleDirectory {
			if strings.Contains(relativePath, string(filepath.Separator)) {
				return nil
			}
		}

		destPath := filepath.Join(opts.DestDir, relativePath)

		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}

		return copyFile(path, destPath)
	})
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
