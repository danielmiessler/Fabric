package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestFilePluginConfinement checks that each operation rejects a path that
// goes out of the working folder, and that a relative path in the folder
// works. The test uses the default fileReadRoot ".".
func TestFilePluginConfinement(t *testing.T) {
	base := t.TempDir()
	root := filepath.Join(base, "root")
	if err := os.Mkdir(root, 0o755); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(base, "outside.txt")
	if err := os.WriteFile(outside, []byte("outside"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "inside.txt"), []byte("inside"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Two symbolic links go out of the root. One stays in the root.
	links := map[string]string{
		"relative-out.txt": filepath.Join("..", "outside.txt"),
		"absolute-out.txt": outside,
		"relative-in.txt":  "inside.txt",
	}
	for name, target := range links {
		if err := os.Symlink(target, filepath.Join(root, name)); err != nil {
			t.Skipf("cannot make a symbolic link: %v", err)
		}
	}

	t.Chdir(root)

	p := &FilePlugin{}

	for _, path := range []string{"inside.txt", "relative-in.txt"} {
		if got, err := p.Apply("read", path); err != nil || got != "inside" {
			t.Errorf("read %q: got %q, %v; want %q", path, got, err, "inside")
		}
	}

	rejected := []string{
		outside,                           // absolute path out of the root
		filepath.Join(root, "inside.txt"), // absolute path in the root
		"~/outside.txt",
		filepath.Join("..", "outside.txt"),
		"relative-out.txt",
		"absolute-out.txt",
	}
	for _, path := range rejected {
		for _, op := range []string{"read", "tail", "size", "modified"} {
			value := path
			if op == "tail" {
				value += "|1"
			}
			if got, err := p.Apply(op, value); err == nil {
				t.Errorf("%s %q: got %q, want an error", op, path, got)
			}
		}
		if got, _ := p.Apply("exists", path); got == "true" {
			t.Errorf("exists %q: got true, want false or an error", path)
		}
	}
}

// TestFilePluginReadStopsAtLimit checks that read stops at MaxFileSize when
// Stat does not give the correct size. Stat gives the size 0 for /dev/zero.
func TestFilePluginReadStopsAtLimit(t *testing.T) {
	if _, err := os.Stat("/dev/zero"); err != nil {
		t.Skip("no /dev/zero on this system")
	}

	oldRoot := fileReadRoot
	fileReadRoot = "/dev"
	defer func() { fileReadRoot = oldRoot }()

	_, err := (&FilePlugin{}).Apply("read", "zero")
	if err == nil || !strings.Contains(err.Error(), "exceeds limit") {
		t.Fatalf("read zero: got error %v, want a size limit error", err)
	}
}

// TestFilePluginTailStopsAtLimit checks that tail stops at MaxFileSize when
// Stat does not give the correct size. /dev/urandom has the size 0 and it
// does not end. It contains newline bytes, thus the scanner does not stop
// because of a long line.
func TestFilePluginTailStopsAtLimit(t *testing.T) {
	if _, err := os.Stat("/dev/urandom"); err != nil {
		t.Skip("no /dev/urandom on this system")
	}

	oldRoot := fileReadRoot
	fileReadRoot = "/dev"
	defer func() { fileReadRoot = oldRoot }()

	_, err := (&FilePlugin{}).Apply("tail", "urandom|5")
	if err == nil || !strings.Contains(err.Error(), "exceeds limit") {
		t.Fatalf("tail urandom: got error %v, want a size limit error", err)
	}
}

// TestFilePluginTailLargeLineCount checks that a large line count does not
// cause a large allocation, and that tail gives all lines of a small file.
func TestFilePluginTailLargeLineCount(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, "small.txt"), []byte("a\nb\nc\n"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}

	oldRoot := fileReadRoot
	fileReadRoot = tmpDir
	defer func() { fileReadRoot = oldRoot }()

	got, err := (&FilePlugin{}).Apply("tail", "small.txt|2000000000")
	if err != nil {
		t.Fatalf("tail with a large line count: %v", err)
	}
	if got != "a\nb\nc" {
		t.Fatalf("tail returned %q, want all 3 lines", got)
	}
}
