package dryrun

import (
	"reflect"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/domain"
)

// Test generated using Keploy
func TestListModels_ReturnsExpectedModel(t *testing.T) {
	client := NewClient()
	models, err := client.ListModels()
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
	channel := make(chan string)
	go func() {
		err := client.SendStream(msgs, opts, channel)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}()
	var receivedMessages []string
	for msg := range channel {
		receivedMessages = append(receivedMessages, msg)
	}
	if len(receivedMessages) == 0 {
		t.Errorf("Expected to receive messages, but got none")
	}
}
