package fsdb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/i18n"
)

func TestStorage_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir}
	name := "test"
	content := []byte("test content")
	if err := storage.Save(name, content); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	loadedContent, err := storage.Load(name)
	if err != nil {
		t.Fatalf("failed to load content: %v", err)
	}
	if string(loadedContent) != string(content) {
		t.Errorf("expected %v, got %v", string(content), string(loadedContent))
	}
}

func TestStorage_Exists(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir}
	name := "test"
	if storage.Exists(name) {
		t.Errorf("expected file to not exist")
	}
	if err := storage.Save(name, []byte("test content")); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	if !storage.Exists(name) {
		t.Errorf("expected file to exist")
	}
}

func TestStorage_Delete(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir}
	name := "test"
	if err := storage.Save(name, []byte("test content")); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	if err := storage.Delete(name); err != nil {
		t.Fatalf("failed to delete content: %v", err)
	}
	if storage.Exists(name) {
		t.Errorf("expected file to be deleted")
	}
}

// invalidStorageNames are names that ValidateStorageName must reject on
// each platform. The storage tests and the pattern traversal tests
// share this list, and one new attack name gets a test at each
// location. The backslash cases guard the `\` half of the separator
// check. That half is the Windows-only escape guard that a "simplify
// to filepath.Base" refactor removes without a test failure. The colon,
// reserved-name, and trailing dot and space cases guard the Windows
// protections: NTFS alternate data streams, DOS device names, and name
// suffixes that Windows removes.
var invalidStorageNames = []string{
	"..", "../keep.txt", "/etc/passwd", "foo/../../keep.txt", ".", "",
	`foo\bar`, `..\x`,
	"foo:bar", "NUL", "con.txt", "CON.tar.gz", "foo.", "foo ",
}

// newTraversalFixture returns a storage entity in a temporary root and
// a marker file out of the entity directory. It also returns a check
// that fails the test if the marker or the entity directory is gone.
func newTraversalFixture(t *testing.T) (storage *StorageEntity, checkSurvived func()) {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, "contexts")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(root, "keep.txt")
	if err := os.WriteFile(marker, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}
	storage = &StorageEntity{Dir: dir, Label: "Contexts"}
	checkSurvived = func() {
		t.Helper()
		if _, err := os.Stat(marker); err != nil {
			t.Fatalf("parent marker was removed: %v", err)
		}
		if _, err := os.Stat(dir); err != nil {
			t.Fatalf("storage dir was removed: %v", err)
		}
	}
	return
}

func TestStorage_RejectsPathTraversal(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}

	storage, checkSurvived := newTraversalFixture(t)
	for _, name := range invalidStorageNames {
		t.Run(name, func(t *testing.T) {
			if err := storage.Delete(name); err == nil {
				t.Fatalf("Delete(%q) succeeded, want error", name)
			}
			if err := storage.Save(name, []byte("pwned")); err == nil {
				t.Fatalf("Save(%q) succeeded, want error", name)
			}
			if _, err := storage.Load(name); err == nil {
				t.Fatalf("Load(%q) succeeded, want error", name)
			}
			if storage.Exists(name) {
				t.Fatalf("Exists(%q) is true, want false", name)
			}
		})
	}

	checkSurvived()
}

func TestInvalidStorageNameError_DefaultMessage(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}
	err := &InvalidStorageNameError{Name: "bad:name"}
	if got, want := err.Error(), `invalid name: "bad:name"`; got != want {
		t.Fatalf("Error() = %q, want %q", got, want)
	}
}

func TestStorage_Rename(t *testing.T) {
	dir := t.TempDir()
	storage := &StorageEntity{Dir: dir, Label: "Contexts"}
	if err := storage.Save("old", []byte("content")); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}
	if err := storage.Rename("old", "new"); err != nil {
		t.Fatalf("failed to rename: %v", err)
	}
	if storage.Exists("old") {
		t.Errorf("expected old name to be gone")
	}
	loaded, err := storage.Load("new")
	if err != nil {
		t.Fatalf("failed to load renamed content: %v", err)
	}
	if string(loaded) != "content" {
		t.Errorf("expected %q, got %q", "content", string(loaded))
	}
}

func TestStorage_RenameRejectsPathTraversal(t *testing.T) {
	if _, err := i18n.Init("en"); err != nil {
		t.Fatalf("i18n.Init() error = %v", err)
	}

	storage, checkSurvived := newTraversalFixture(t)
	if err := storage.Save("ok", []byte("content")); err != nil {
		t.Fatalf("failed to save content: %v", err)
	}

	for _, name := range invalidStorageNames {
		t.Run(name, func(t *testing.T) {
			if err := storage.Rename("ok", name); err == nil {
				t.Fatalf("Rename(%q, %q) succeeded, want error", "ok", name)
			}
			if err := storage.Rename(name, "ok"); err == nil {
				t.Fatalf("Rename(%q, %q) succeeded, want error", name, "ok")
			}
		})
	}

	checkSurvived()
	if !storage.Exists("ok") {
		t.Fatalf("legitimate entry was moved or deleted")
	}
}

// mustSymlink makes a symlink. If symlinks are not available, for
// example on Windows without the privilege, it skips the test.
func mustSymlink(t *testing.T, target, link string) {
	t.Helper()
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("cannot create symlink: %v", err)
	}
}

func TestStorage_RejectsSymlinkEscape(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "store")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(root, "outside.txt")
	if err := os.WriteFile(outside, []byte("keep"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustSymlink(t, outside, filepath.Join(dir, "escape"))
	mustSymlink(t, filepath.Join(root, "missing.txt"), filepath.Join(dir, "dangling"))

	storage := &StorageEntity{Dir: dir}
	for _, name := range []string{"escape", "dangling"} {
		if _, err := storage.Load(name); err == nil {
			t.Fatalf("Load(%q) through an outside symlink did not fail", name)
		}
		if err := storage.Save(name, []byte("pwned")); err == nil {
			t.Fatalf("Save(%q) through an outside symlink did not fail", name)
		}
	}
	if got, _ := os.ReadFile(outside); string(got) != "keep" {
		t.Fatalf("outside file was overwritten: %q", got)
	}
	if _, err := os.Stat(filepath.Join(root, "missing.txt")); err == nil {
		t.Fatal("dangling symlink target was created")
	}
}

// Load and Save operate through a symlink that stays in the storage
// directory, and through a storage directory that is a symlink.
func TestStorage_AllowsInternalAndDirSymlinks(t *testing.T) {
	root := t.TempDir()
	realDir := filepath.Join(root, "real")
	if err := os.MkdirAll(realDir, 0o755); err != nil {
		t.Fatal(err)
	}

	storage := &StorageEntity{Dir: realDir}
	if err := storage.Save("target", []byte("content")); err != nil {
		t.Fatal(err)
	}
	mustSymlink(t, filepath.Join(realDir, "target"), filepath.Join(realDir, "alias"))
	if got, err := storage.Load("alias"); err != nil || string(got) != "content" {
		t.Fatalf("Load through an internal symlink: got %q, err %v", got, err)
	}

	linkDir := filepath.Join(root, "link")
	mustSymlink(t, realDir, linkDir)
	linked := &StorageEntity{Dir: linkDir}
	if got, err := linked.Load("target"); err != nil || string(got) != "content" {
		t.Fatalf("Load via a symlinked storage dir: got %q, err %v", got, err)
	}
	if err := linked.Save("new", []byte("x")); err != nil {
		t.Fatalf("Save via a symlinked storage dir: %v", err)
	}
}
