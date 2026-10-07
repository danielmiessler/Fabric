package dryrun

import (
	"context"
	"reflect"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

// Test generated using Keploy
func TestListModels_ReturnsExpectedModel(t *testing.T) {
	client := NewClient()
	models, err := client.ListModels(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := []string{"dry-run-model"}
	if !reflect.DeepEqual(models, expected) {
		t.Errorf("Expected %v, got %v", expected, models)
	}
}

// Test generated using Keploy
func TestSetup_ReturnsNil(t *testing.T) {
	client := NewClient()
	err := client.Setup()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// Test generated using Keploy
func TestSendStream_SendsMessages(t *testing.T) {
	client := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Test message"},
	}
	opts := &domain.ChatOptions{
		Model: "dry-run-model",
	}
	channel := make(chan domain.StreamUpdate)
	go func() {
		err := client.SendStream(context.Background(), msgs, opts, channel)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}()
	var receivedMessages []string
	for msg := range channel {
		receivedMessages = append(receivedMessages, msg.Content)
	}
	if len(receivedMessages) == 0 {
		t.Errorf("Expected to receive messages, but got none")
	}
}

func TestConstructRequest_ExactOutput(t *testing.T) {
	msgs := []*chat.ChatCompletionMessage{
		{Role: chat.ChatMessageRoleSystem, Content: "SYSTEM"},
		{Role: chat.ChatMessageRoleUser, Content: "hello"},
	}
	got := NewClient().constructRequest(msgs, &domain.ChatOptions{Model: "m"})
	want := "Dry run: Would send the following request:\n\n" +
		"System:\nSYSTEM\n\nUser:\nhello\n\n" +
		"Options:\nModel: m\nTemperature: 0.000000\nTopP: 0.000000\n" +
		"PresencePenalty: 0.000000\nFrequencyPenalty: 0.000000\n"
	if got != want {
		t.Errorf("constructRequest() = %q, want %q", got, want)
	}
}
