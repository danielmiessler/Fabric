package youtube

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFindVTTFilesWithFallback_SkipsSymlink checks that a .vtt symlink is not
// returned and that a regular .vtt file is returned.
func TestFindVTTFilesWithFallback_SkipsSymlink(t *testing.T) {
	dir := t.TempDir()
	outside := filepath.Join(t.TempDir(), "outside.vtt")
	if err := os.WriteFile(outside, []byte("WEBVTT\n\noutside"), 0o600); err != nil {
		t.Fatalf("write outside: %v", err)
	}
	// A .vtt symlink to a file outside the folder.
	if err := os.Symlink(outside, filepath.Join(dir, "link.vtt")); err != nil {
		t.Skipf("symlink unsupported: %v", err)
	}
	// A regular .vtt file.
	good := filepath.Join(dir, "good.en.vtt")
	if err := os.WriteFile(good, []byte("WEBVTT\n\nok"), 0o600); err != nil {
		t.Fatalf("write good: %v", err)
	}

	o := &YouTube{}
	got, err := o.findVTTFilesWithFallback(dir, "")
	if err != nil {
		t.Fatalf("findVTTFilesWithFallback error: %v", err)
	}
	for _, f := range got {
		if filepath.Base(f) == "link.vtt" {
			t.Fatalf(".vtt symlink was returned: %v", got)
		}
	}
	if len(got) != 1 || filepath.Base(got[0]) != "good.en.vtt" {
		t.Fatalf("expected only the regular .vtt, got %v", got)
	}
}
