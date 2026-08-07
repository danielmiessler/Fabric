package ollama

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
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

func TestLoadImageBytes_HTTPFetchErrorIsLocalized(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)

	client := &Client{httpClient: server.Client()}

	_, err = client.loadImageBytes(context.Background(), server.URL+"/image.png")
	require.Error(t, err)
	assert.Equal(t,
		fmt.Sprintf(i18n.T("ollama_failed_fetch_image"), server.URL+"/image.png", "500 Internal Server Error"),
		err.Error(),
	)
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
