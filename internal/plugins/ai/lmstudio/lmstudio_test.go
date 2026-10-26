package lmstudio

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func hangingClient(t *testing.T) *Client {
	t.Helper()
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-release:
		}
	}))
	t.Cleanup(func() {
		close(release)
		server.Close()
	})

	client := NewClient()
	client.ApiUrl.Value = server.URL
	client.HttpClient = server.Client()
	return client
}

func TestSendStreamHonorsContextCancellation(t *testing.T) {
	client := hangingClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	channel := make(chan domain.StreamUpdate, 1)
	done := make(chan error, 1)
	go func() {
		done <- client.SendStream(ctx, []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}, &domain.ChatOptions{Model: "m"}, channel)
	}()

	select {
	case err := <-done:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(3 * time.Second):
		t.Fatal("SendStream ignored the cancelled context")
	}
}

func TestListModelsHonorsContextCancellation(t *testing.T) {
	client := hangingClient(t)
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	done := make(chan error, 1)
	go func() {
		_, err := client.ListModels(ctx)
		done <- err
	}()

	select {
	case err := <-done:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(3 * time.Second):
		t.Fatal("ListModels ignored the cancelled context")
	}
}

// chatPayloads sends one Send and one SendStream request with opts and returns the JSON body of each.
func chatPayloads(t *testing.T, opts *domain.ChatOptions) (send, stream map[string]any) {
	t.Helper()

	var bodies []map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		bodies = append(bodies, body)
		if body["stream"] == true {
			w.Header().Set("Content-Type", "text/event-stream")
			_, _ = w.Write([]byte("data: [DONE]\n\n"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"ok"}}]}`))
	}))
	t.Cleanup(server.Close)

	client := NewClient()
	client.ApiUrl.Value = server.URL
	client.HttpClient = server.Client()

	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hello"}}

	_, err := client.Send(context.Background(), msgs, opts)
	require.NoError(t, err)

	channel := make(chan domain.StreamUpdate, 4)
	require.NoError(t, client.SendStream(context.Background(), msgs, opts, channel))

	require.Len(t, bodies, 2)
	return bodies[0], bodies[1]
}

func TestChatPayloadCarriesTheSamplingOptions(t *testing.T) {
	t.Parallel()

	send, stream := chatPayloads(t, &domain.ChatOptions{
		Model:            "m",
		Temperature:      0.3,
		TopP:             0.8,
		MaxTokens:        256,
		PresencePenalty:  0.5,
		FrequencyPenalty: 0.25,
		Seed:             42,
	})

	for name, body := range map[string]map[string]any{"Send": send, "SendStream": stream} {
		require.Equal(t, 0.3, body["temperature"], name)
		require.Equal(t, 0.8, body["top_p"], name)
		require.EqualValues(t, 256, body["max_tokens"], name)
		require.Equal(t, 0.5, body["presence_penalty"], name)
		require.Equal(t, 0.25, body["frequency_penalty"], name)
		require.EqualValues(t, 42, body["seed"], name)
	}
}

func TestChatPayloadOmitsUnsetOptionsAndHonorsRaw(t *testing.T) {
	t.Parallel()

	send, stream := chatPayloads(t, &domain.ChatOptions{Model: "m", Temperature: 0.7})
	for name, body := range map[string]map[string]any{"Send": send, "SendStream": stream} {
		require.Equal(t, 0.7, body["temperature"], name)
		for _, key := range []string{"top_p", "max_tokens", "presence_penalty", "frequency_penalty", "seed"} {
			require.NotContains(t, body, key, name)
		}
	}

	send, stream = chatPayloads(t, &domain.ChatOptions{Model: "m", Raw: true, Temperature: 0.7, Seed: 42})
	for name, body := range map[string]map[string]any{"Send": send, "SendStream": stream} {
		require.NotContains(t, body, "temperature", name)
		require.NotContains(t, body, "seed", name)
	}
}
