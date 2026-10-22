package bedrock

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// streamMessage encodes one event of a ConverseStream response.
func streamMessage(t *testing.T, buf *bytes.Buffer, messageType, name, payload string) {
	t.Helper()
	kind := ":event-type"
	if messageType == "exception" {
		kind = ":exception-type"
	}
	msg := eventstream.Message{
		Headers: eventstream.Headers{
			{Name: ":message-type", Value: eventstream.StringValue(messageType)},
			{Name: kind, Value: eventstream.StringValue(name)},
			{Name: ":content-type", Value: eventstream.StringValue("application/json")},
		},
		Payload: []byte(payload),
	}
	require.NoError(t, eventstream.NewEncoder().Encode(buf, msg))
}

// streamClient serves body as the ConverseStream response.
func streamClient(t *testing.T, body []byte) *BedrockClient {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.amazon.eventstream")
		_, _ = w.Write(body)
	}))
	t.Cleanup(server.Close)

	client := NewClient()
	client.runtimeClient = bedrockruntime.New(bedrockruntime.Options{
		Region:       "us-east-1",
		BaseEndpoint: aws.String(server.URL),
		Credentials:  credentials.NewStaticCredentialsProvider("id", "secret", ""),
	})
	return client
}

func sendStream(t *testing.T, client *BedrockClient) (updates []domain.StreamUpdate, err error) {
	t.Helper()
	channel := make(chan domain.StreamUpdate, 16)
	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}
	err = client.SendStream(context.Background(), msgs, &domain.ChatOptions{Model: "test-model"}, channel)
	for u := range channel {
		updates = append(updates, u)
	}
	return updates, err
}

// The usage is in the Metadata event, which follows MessageStop, so it must still be read
// after the stop event.
func TestSendStream_EmitsUsageAfterMessageStop(t *testing.T) {
	var body bytes.Buffer
	streamMessage(t, &body, "event", "contentBlockDelta", `{"contentBlockIndex":0,"delta":{"text":"Hello"}}`)
	streamMessage(t, &body, "event", "messageStop", `{"stopReason":"end_turn"}`)
	streamMessage(t, &body, "event", "metadata", `{"usage":{"inputTokens":7,"outputTokens":3,"totalTokens":10},"metrics":{"latencyMs":1}}`)

	updates, err := sendStream(t, streamClient(t, body.Bytes()))
	require.NoError(t, err)

	var usage *domain.UsageMetadata
	for _, u := range updates {
		if u.Type == domain.StreamTypeUsage {
			usage = u.Usage
		}
	}
	require.NotNil(t, usage, "the usage in the Metadata event was never emitted")
	assert.Equal(t, 7, usage.InputTokens)
	assert.Equal(t, 3, usage.OutputTokens)
	assert.Equal(t, 10, usage.TotalTokens)
}

// An error in the middle of the stream must be returned, not reported as a finished answer.
func TestSendStream_ReturnsMidStreamError(t *testing.T) {
	var body bytes.Buffer
	streamMessage(t, &body, "event", "contentBlockDelta", `{"contentBlockIndex":0,"delta":{"text":"Hel"}}`)
	streamMessage(t, &body, "exception", "modelStreamErrorException", `{"message":"model stream failed"}`)

	updates, err := sendStream(t, streamClient(t, body.Bytes()))
	require.Error(t, err, "a stream that ended in an error was reported as a success")
	assert.Contains(t, err.Error(), "model stream failed")

	require.NotEmpty(t, updates)
	assert.Equal(t, "Hel", updates[0].Content)
}

// hangingServer answers nothing until the test ends.
func hangingServer(t *testing.T) *httptest.Server {
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
	return server
}

// Cancelling the caller's context must stop SendStream instead of letting the request run on.
func TestSendStream_HonorsContextCancellation(t *testing.T) {
	server := hangingServer(t)
	client := NewClient()
	client.runtimeClient = bedrockruntime.New(bedrockruntime.Options{
		Region:           "us-east-1",
		BaseEndpoint:     aws.String(server.URL),
		Credentials:      credentials.NewStaticCredentialsProvider("id", "secret", ""),
		RetryMaxAttempts: 1,
	})

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	channel := make(chan domain.StreamUpdate, 16)
	msgs := []*chat.ChatCompletionMessage{{Role: chat.ChatMessageRoleUser, Content: "hi"}}
	done := make(chan error, 1)
	go func() { done <- client.SendStream(ctx, msgs, &domain.ChatOptions{Model: "test-model"}, channel) }()

	select {
	case err := <-done:
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	case <-time.After(3 * time.Second):
		t.Fatal("SendStream ignored the cancelled context")
	}
}

// Cancelling the caller's context must stop ListModels as well.
func TestListModels_HonorsContextCancellation(t *testing.T) {
	server := hangingServer(t)
	client := NewClient()
	client.controlPlaneClient = bedrock.New(bedrock.Options{
		Region:           "us-east-1",
		BaseEndpoint:     aws.String(server.URL),
		Credentials:      credentials.NewStaticCredentialsProvider("id", "secret", ""),
		RetryMaxAttempts: 1,
	})

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	done := make(chan error, 1)
	go func() {
		_, err := client.ListModels(ctx)
		done <- err
	}()

	select {
	case err := <-done:
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	case <-time.After(3 * time.Second):
		t.Fatal("ListModels ignored the cancelled context")
	}
}
