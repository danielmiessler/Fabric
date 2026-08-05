package youtube

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeExecutable(t *testing.T, dir, name, contents string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(contents), 0o755); err != nil {
		t.Fatalf("failed to write %s: %v", name, err)
	}
	return path
}

func TestGrabVisual_UsesFrameTimesFromFFmpeg(t *testing.T) {
	binDir := t.TempDir()
	writeExecutable(t, binDir, "yt-dlp", "#!/usr/bin/env bash\necho 'http://example.invalid/video'\n")
	writeExecutable(t, binDir, "ffmpeg", "#!/usr/bin/env bash\nout=\"${@: -1}\"\ntouch \"${out//%04d/0001}\"\ntouch \"${out//%04d/0002}\"\necho '[Parsed_showinfo_1 @ 0x1] n:   0 pts:      0 pts_time:0' >&2\necho '[Parsed_showinfo_1 @ 0x1] n:   1 pts:    500 pts_time:0.5' >&2\n")
	writeExecutable(t, binDir, "tesseract", "#!/usr/bin/env bash\ncase \"$1\" in\n  *0001.jpg) echo 'First frame text' ;;\n  *0002.jpg) echo 'Second frame text' ;;\n  *) echo 'Unknown frame text' ;;\nesac\n")

	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

	yt := NewYouTube()
	got, err := yt.GrabVisual("video123", "en", "", 0.4, 2)
	if err != nil {
		t.Fatalf("GrabVisual returned error: %v", err)
	}

	if !strings.Contains(got, "00:00:00.500 --> 00:00:01.499") {
		t.Fatalf("expected second frame to use ffmpeg-derived timing, got %q", got)
	}
	if !strings.Contains(got, "Second frame text") {
		t.Fatalf("expected second frame OCR text in output, got %q", got)
	}
}
