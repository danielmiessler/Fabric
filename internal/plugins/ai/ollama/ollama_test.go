package ollama

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/util"
	ollamaapi "github.com/ollama/ollama/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadImageBytes_DataURLValidationErrorsAreLocalized(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	client := &Client{}

	_, err = client.loadImageBytes(context.Background(), "data:image/png;base64")
	require.Error(t, err)
	assert.Equal(t, i18n.T("ollama_invalid_data_url_format"), err.Error())

	_, err = client.loadImageBytes(context.Background(), "data:image/png;base64,%%%%")
	require.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), strings.Split(i18n.T("ollama_failed_decode_data_url"), "%v")[0]))
}

// allowTestServer lets loadImageBytes connect to the loopback address of
// server. The check stays on for all other addresses.
func allowTestServer(t *testing.T, server *httptest.Server) {
	old := imageDialControl
	imageDialControl = func(network, address string, c syscall.RawConn) error {
		if address == server.Listener.Addr().String() {
			return nil
		}
		return util.DenyNonPublicAddress(network, address, c)
	}
	t.Cleanup(func() { imageDialControl = old })
}

func TestLoadImageBytes_HTTPFetchErrorIsLocalized(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)
	allowTestServer(t, server)

	client := &Client{httpClient: server.Client()}

	_, err = client.loadImageBytes(context.Background(), server.URL+"/image.png")
	require.Error(t, err)
	assert.Equal(t,
		fmt.Sprintf(i18n.T("ollama_failed_fetch_image"), server.URL+"/image.png", "500 Internal Server Error"),
		err.Error(),
	)
}

func TestLoadImageBytes_RefusesLoopbackAddress(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	var called bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called = true
		_, _ = w.Write([]byte("img"))
	}))
	t.Cleanup(server.Close)

	client := &Client{}
	_, err = client.loadImageBytes(context.Background(), server.URL+"/image.png")
	require.Error(t, err)
	assert.Contains(t, err.Error(), i18n.T("util_error_non_public_address"))
	assert.False(t, called)
}

func TestLoadImageBytes_RefusesOtherSchemes(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)
	client := &Client{}
	for _, u := range []string{"file:///etc/hosts", "gopher://127.0.0.1"} {
		_, err := client.loadImageBytes(context.Background(), u)
		assert.ErrorContains(t, err, "unsupported protocol scheme", u)
	}
}

func TestLoadImageBytes_SendsNoAPIKey(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	gotAuth := "not called"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte("img"))
	}))
	t.Cleanup(server.Close)
	allowTestServer(t, server)

	// configure gives the Ollama API client a transport that adds the key.
	client := NewClient()
	client.ApiUrl.Value = server.URL
	client.ApiKey.Value = "test-key"
	require.NoError(t, client.configure())

	img, err := client.loadImageBytes(context.Background(), server.URL+"/image.png")
	require.NoError(t, err)
	assert.Equal(t, []byte("img"), img)
	assert.Empty(t, gotAuth)
}

func TestLoadImageBytes_RefusesLargeImage(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(make([]byte, maxImageSize+1))
	}))
	t.Cleanup(server.Close)
	allowTestServer(t, server)

	img, err := (&Client{}).loadImageBytes(context.Background(), server.URL+"/image.png")
	assert.Nil(t, img)
	assert.EqualError(t, err, fmt.Sprintf(i18n.T("ollama_image_too_large"), server.URL+"/image.png", maxImageSize))
}

func TestLoadImageBytes_DataURLSuccess(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	client := &Client{}
	expected := []byte("hello world")
	dataURL := "data:image/png;base64," + base64.StdEncoding.EncodeToString(expected)

	got, err := client.loadImageBytes(context.Background(), dataURL)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

// TestSendStreamHonorsContextCancellation verifies that cancelling the caller's
// context aborts an in-flight Ollama generation, instead of running the stream to
// server completion. Regression test for issue #2196.
func TestSendStreamHonorsContextCancellation(t *testing.T) {
	const totalChunks = 20

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		for i := range totalChunks {
			if r.Context().Err() != nil {
				return
			}
			_ = enc.Encode(ollamaapi.ChatResponse{Message: ollamaapi.Message{Content: "chunk "}, Done: i == totalChunks-1})
			w.(http.Flusher).Flush()
			time.Sleep(25 * time.Millisecond)
		}
	}))
	t.Cleanup(server.Close)

	baseURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	client := &Client{client: ollamaapi.NewClient(baseURL, server.Client())}
	channel := make(chan domain.StreamUpdate)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	errCh := make(chan error, 1)
	go func() {
		errCh <- client.SendStream(ctx, []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hello"}}, &domain.ChatOptions{Model: "test-model"}, channel)
	}()

	// Consume the first chunk, then cancel mid-stream.
	_, ok := <-channel
	require.True(t, ok, "expected at least one stream update before cancellation")
	cancel()

	// Drain any remaining updates so SendStream can return.
	for range channel {
	}

	require.ErrorIs(t, <-errCh, context.Canceled)
}

func TestSendStreamClosesChannelOnChatError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"ollama failed"}` + "\n"))
	}))
	t.Cleanup(server.Close)

	baseURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	client := &Client{client: ollamaapi.NewClient(baseURL, server.Client())}
	channel := make(chan domain.StreamUpdate)

	err = client.SendStream(
		context.Background(),
		[]*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hello"}},
		&domain.ChatOptions{Model: "missing-model"},
		channel,
	)

	require.Error(t, err)
	_, ok := <-channel
	assert.False(t, ok, "stream channel should be closed when Ollama chat returns an error")
}
