package githelper

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestCopyFile_RefusesSymlink checks that copyFile does not follow a symlink.
func TestCopyFile_RefusesSymlink(t *testing.T) {
	dir := t.TempDir()

	target := filepath.Join(dir, "target.txt")
	if err := os.WriteFile(target, []byte("outside"), 0o600); err != nil {
		t.Fatalf("write target: %v", err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("symlink unsupported: %v", err)
	}
	dst := filepath.Join(dir, "copied")

	if err := copyFile(link, dst); err == nil {
		t.Fatalf("copyFile followed a symlink; expected refusal")
	}
	if _, err := os.Stat(dst); err == nil {
		t.Fatalf("symlinked content was copied to destination")
	}
}

// TestCopyFile_CopiesRegular checks that copyFile still copies a regular file.
func TestCopyFile_CopiesRegular(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "good.txt")
	if err := os.WriteFile(src, []byte("hello"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	dst := filepath.Join(dir, "out.txt")
	if err := copyFile(src, dst); err != nil {
		t.Fatalf("copyFile(regular) error: %v", err)
	}
	got, err := os.ReadFile(dst)
	if err != nil || string(got) != "hello" {
		t.Fatalf("copy mismatch: %q err=%v", got, err)
	}
}

// makeRepo makes a git repository with patterns/a.md, a symlink
// patterns/link.md and patterns/sub/b.md, and returns its path.
func makeRepo(t *testing.T) string {
	t.Helper()
	repo := t.TempDir()
	if err := os.MkdirAll(filepath.Join(repo, "patterns", "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repo, "patterns", "a.md"), []byte("a"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repo, "patterns", "sub", "b.md"), []byte("b"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("a.md", filepath.Join(repo, "patterns", "link.md")); err != nil {
		t.Skipf("symlink unsupported: %v", err)
	}
	for _, args := range [][]string{
		{"init", "-q"},
		{"add", "."},
		{"-c", "user.name=t", "-c", "user.email=t@t", "commit", "-q", "-m", "x"},
	} {
		cmd := exec.Command("git", args...)
		cmd.Dir = repo
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	return repo
}

// TestFetchFilesFromRepo checks the go-git path. It also checks that
// FetchFilesFromRepo adds the "/" to PathPrefix and obeys SingleDirectory.
func TestFetchFilesFromRepo(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}
	repo := makeRepo(t)

	dest := t.TempDir()
	if err := FetchFilesFromRepo(FetchOptions{RepoURL: "file://" + repo, PathPrefix: "patterns", DestDir: dest, SingleDirectory: true}); err != nil {
		t.Fatal(err)
	}
	if got, err := os.ReadFile(filepath.Join(dest, "a.md")); err != nil || string(got) != "a" {
		t.Fatalf("a.md = %q, %v", got, err)
	}
	if _, err := os.Stat(filepath.Join(dest, "sub")); err == nil {
		t.Fatal("SingleDirectory copied a subfolder")
	}
	if fi, err := os.Lstat(filepath.Join(dest, "link.md")); err == nil && fi.Mode()&os.ModeSymlink != 0 {
		t.Fatal("the clone made a symlink")
	}

	dest = t.TempDir()
	if err := fetchFilesViaGoGit(FetchOptions{RepoURL: "file://" + repo, PathPrefix: "patterns/", DestDir: dest}); err != nil {
		t.Fatal(err)
	}
	if got, err := os.ReadFile(filepath.Join(dest, "sub", "b.md")); err != nil || string(got) != "b" {
		t.Fatalf("sub/b.md = %q, %v", got, err)
	}
}

// TestFetchFilesViaGitCLI checks that a URL that looks like a git option
// does not run as an option. It also checks that a good clone copies the
// regular files and skips a symlink.
func TestFetchFilesViaGitCLI(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}
	repo := makeRepo(t)

	// With "--", git reads the URL as a repository name and names it in
	// the error. With no "--", git reads it as an option.
	marker := filepath.Join(t.TempDir(), "marker")
	badURL := "--upload-pack=touch " + marker
	err := fetchFilesViaGitCLI(FetchOptions{RepoURL: badURL, DestDir: t.TempDir()})
	if err == nil || !strings.Contains(err.Error(), "'"+badURL+"'") {
		t.Fatalf("clone of an option-like URL: got %v, want an error that names the URL as the repository", err)
	}
	if _, statErr := os.Stat(marker); statErr == nil {
		t.Fatal("git ran the --upload-pack value")
	}

	dest := t.TempDir()
	if err := fetchFilesViaGitCLI(FetchOptions{RepoURL: "file://" + repo, PathPrefix: "patterns/", DestDir: dest}); err != nil {
		t.Fatal(err)
	}
	if got, err := os.ReadFile(filepath.Join(dest, "a.md")); err != nil || string(got) != "a" {
		t.Fatalf("a.md = %q, %v", got, err)
	}
	if _, err := os.Lstat(filepath.Join(dest, "link.md")); err == nil {
		t.Fatal("the clone copied a symlink")
	}
}
