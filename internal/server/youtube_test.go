package restapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
	"github.com/gin-gonic/gin"
)

// fakeYtDlp is a yt-dlp stand-in. It writes one VTT file into the folder of
// the -o argument.
const fakeYtDlp = `#!/bin/sh
while [ $# -gt 0 ]; do
	[ "$1" = "-o" ] && out="$2"
	shift
done
cat > "$(dirname "$out")/video.en.vtt" <<'EOF'
WEBVTT

00:00:01.000 --> 00:00:02.000
hello world
EOF
`

func TestYouTubeTranscript(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	// With no API key, GrabMetadata fails before it uses the network.
	NewYouTubeHandler(r, &core.PluginRegistry{YouTube: youtube.NewYouTube()})

	fakeBin := t.TempDir()
	fakePath := fakeBin + string(os.PathListSeparator) + os.Getenv("PATH")
	if err := os.WriteFile(filepath.Join(fakeBin, "yt-dlp"), []byte(fakeYtDlp), 0o755); err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name     string
		path     string
		body     string
		wantCode int
		wantText string
	}{
		{"missing url", fakePath, `{}`, http.StatusBadRequest, "invalid request"},
		{"blank url", fakePath, `{"url":""}`, http.StatusBadRequest, "invalid request"},
		{"invalid url", fakePath, `{"url":"https://example.com/"}`, http.StatusBadRequest, "error"},
		{"playlist url", fakePath, `{"url":"https://www.youtube.com/playlist?list=PL1234567890"}`, http.StatusBadRequest, "playlist"},
		{"yt-dlp missing", t.TempDir(), `{"url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ"}`, http.StatusInternalServerError, "error"},
		{"valid url", fakePath, `{"url":"https://www.youtube.com/watch?v=dQw4w9WgXcQ"}`, http.StatusOK, "hello world"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("PATH", tc.path)
			req := httptest.NewRequest(http.MethodPost, "/youtube/transcript", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.wantCode {
				t.Fatalf("got %d, want %d: %s", w.Code, tc.wantCode, w.Body)
			}
			if !strings.Contains(w.Body.String(), tc.wantText) {
				t.Errorf("body %s does not contain %q", w.Body, tc.wantText)
			}
			if tc.wantCode != http.StatusOK {
				return
			}
			var resp YouTubeResponse
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatal(err)
			}
			if resp.VideoId != "dQw4w9WgXcQ" || resp.Title != "dQw4w9WgXcQ" {
				t.Errorf("got videoId %q, title %q, want the video ID in both", resp.VideoId, resp.Title)
			}
		})
	}
}
