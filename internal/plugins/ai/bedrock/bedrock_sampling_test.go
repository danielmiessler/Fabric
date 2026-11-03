package bedrock

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const converseResponse = `{"output":{"message":{"role":"assistant","content":[{"text":"ok"}]}},"stopReason":"end_turn","usage":{"inputTokens":1,"outputTokens":1,"totalTokens":2}}`

// recordingClient answers every request with response and stores the last request body.
func recordingClient(t *testing.T, contentType string, response []byte) (*BedrockClient, *[]byte) {
	t.Helper()
	var captured []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(response)
	}))
	t.Cleanup(server.Close)

	client := NewClient()
	client.runtimeClient = bedrockruntime.New(bedrockruntime.Options{
		Region:       "us-east-1",
		BaseEndpoint: aws.String(server.URL),
		Credentials:  credentials.NewStaticCredentialsProvider("id", "secret", ""),
	})
	return client, &captured
}

func inferenceConfigOf(t *testing.T, body []byte) map[string]any {
	t.Helper()
	var payload map[string]any
	require.NoError(t, json.Unmarshal(body, &payload))
	config, _ := payload["inferenceConfig"].(map[string]any)
	return config
}

func sendOnce(t *testing.T, stream bool, model string) []byte {
	t.Helper()
	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}
	opts := &domain.ChatOptions{Model: model, Temperature: 0.7, TopP: domain.DefaultTopP}

	if stream {
		var events bytes.Buffer
		streamMessage(t, &events, "event", "messageStop", `{"stopReason":"end_turn"}`)
		client, captured := recordingClient(t, "application/vnd.amazon.eventstream", events.Bytes())
		channel := make(chan domain.StreamUpdate, 16)
		require.NoError(t, client.SendStream(context.Background(), msgs, opts, channel))
		return *captured
	}

	client, captured := recordingClient(t, "application/json", []byte(converseResponse))
	_, err := client.Send(context.Background(), msgs, opts)
	require.NoError(t, err)
	return *captured
}

// Claude models that reject temperature fail the whole request with a 400 if it carries one.
func TestSampling_OmittedForModelsThatRejectIt(t *testing.T) {
	for _, stream := range []bool{true, false} {
		for _, model := range []string{"us.anthropic.claude-opus-5-5-v1:0", "anthropic.claude-opus-4-7"} {
			config := inferenceConfigOf(t, sendOnce(t, stream, model))
			assert.NotContains(t, config, "temperature", "stream=%v model=%s", stream, model)
		}
	}
}

func TestSampling_KeptForOtherModels(t *testing.T) {
	for _, stream := range []bool{true, false} {
		config := inferenceConfigOf(t, sendOnce(t, stream, "anthropic.claude-sonnet-4-5-20250929-v1:0"))
		assert.InDelta(t, 0.7, config["temperature"], 0.001, "stream=%v", stream)
	}
}
