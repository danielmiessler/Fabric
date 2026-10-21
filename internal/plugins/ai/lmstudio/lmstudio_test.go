package lmstudio

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestListModelsUsesBearerTokenWhenConfigured(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/models", r.URL.Path)
		require.Equal(t, "Bearer secret", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":[{"id":"model-1"}]}`))
	}))
	defer server.Close()

	client := NewClient()
	client.ApiUrl.Value = server.URL
	client.ApiKey.Value = "secret"
	client.HttpClient = server.Client()

	models, err := client.ListModels(context.Background())
	require.NoError(t, err)
	require.Equal(t, []string{"model-1"}, models)
}

func TestSendEndpointsUseBearerTokenWhenConfigured(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer secret", r.Header.Get("Authorization"))
		switch r.URL.Path {
		case "/chat/completions":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"ok"}}]}`))
		case "/completions":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"choices":[{"text":"ok"}]}`))
		case "/embeddings":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"data":[{"embedding":[1,2]}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := NewClient()
	client.ApiUrl.Value = server.URL
	client.ApiKey.Value = "secret"
	client.HttpClient = server.Client()

	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hello"}}
	opts := &domain.ChatOptions{Model: "test-model"}

	_, err := client.Send(context.Background(), msgs, opts)
	require.NoError(t, err)

	_, err = client.Complete(context.Background(), "hello", opts)
	require.NoError(t, err)

	_, err = client.GetEmbeddings(context.Background(), "hello", opts)
	require.NoError(t, err)
}

func TestListModelsDoesNotSendBearerForWhitespaceOnlyKey(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Empty(t, r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":[{"id":"model-1"}]}`))
	}))
	defer server.Close()

	client := NewClient()
	client.ApiUrl.Value = server.URL
	client.ApiKey.Value = "   "
	client.HttpClient = server.Client()

	models, err := client.ListModels(context.Background())
	require.NoError(t, err)
	require.Equal(t, []string{"model-1"}, models)
}

func TestSendStreamClosesChannelOnError(t *testing.T) {
	t.Parallel()

	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer failing.Close()

	down := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	down.Close()

	for name, url := range map[string]string{"http 500": failing.URL, "server down": down.URL} {
		t.Run(name, func(t *testing.T) {
			client := NewClient()
			client.ApiUrl.Value = url
			client.HttpClient = &http.Client{}

			ch := make(chan domain.StreamUpdate)
			require.Error(t, client.SendStream(context.Background(), []*chat.ChatCompletionMessage{{Role: "user", Content: "hi"}}, &domain.ChatOptions{Model: "m"}, ch))

			// chatter.go ranges over the channel until the vendor closes it.
			select {
			case _, ok := <-ch:
				require.False(t, ok, "unexpected update")
			default:
				t.Fatal("channel was not closed")
			}
		})
	}
}
