package cli

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/plugins/ai/dryrun"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
	"github.com/stretchr/testify/require"
)

func TestYouTubePlaylistOutputWithPattern(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/youtube/v3/playlistItems":
			fmt.Fprint(w, `{"items":[{"snippet":{"resourceId":{"videoId":"video1"},"title":"Test video"}}]}`)
		case "/youtube/v3/videos":
			fmt.Fprint(w, `{"items":[{"id":"video1","snippet":{"title":"Test video"},"statistics":{}}]}`)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	transport := server.Client().Transport.(*http.Transport).Clone()
	transport.TLSClientConfig.ServerName = "127.0.0.1"
	transport.DialContext = func(ctx context.Context, _, _ string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "tcp", server.Listener.Addr().String())
	}
	oldTransport := http.DefaultTransport
	http.DefaultTransport = transport
	t.Cleanup(func() { http.DefaultTransport = oldTransport })

	vendor := &stepVendor{Client: dryrun.NewClient(), reply: "Playlist summary"}
	registry := newStepRegistry(t, vendor, &Workflow{Steps: []WorkflowStep{{Pattern: "summarize"}}})
	registry.YouTube = youtube.NewYouTube()
	registry.YouTube.ApiKey.Value = "test-key"
	flags := &Flags{
		YouTube:         "https://www.youtube.com/playlist?list=test-playlist",
		YouTubePlaylist: true,
		Output:          filepath.Join(t.TempDir(), "playlist.csv"),
	}

	input, err := handleToolProcessing(flags, registry)
	require.NoError(t, err)
	require.Empty(t, input)
	csv, err := os.ReadFile(flags.Output)
	require.NoError(t, err)
	require.Equal(t, "VideoID,Title\nvideo1,Test video\n", string(csv))

	flags.Pattern = "summarize"
	flags.Language = "en"
	flags.YouTubeMetadata = true
	flags.Output = filepath.Join(t.TempDir(), "summary.md")
	input, err = handleToolProcessing(flags, registry)
	require.NoError(t, err)
	require.Contains(t, input, "Test video")
	require.NoFileExists(t, flags.Output)

	require.NoError(t, handleChatProcessing(flags, registry, input))
	require.Len(t, vendor.inputs, 1)
	require.Contains(t, vendor.inputs[0], "Test video")
	summary, err := os.ReadFile(flags.Output)
	require.NoError(t, err)
	require.Equal(t, "Playlist summary\n", string(summary))
}
