package dryrun

import (
	"reflect"
	"strings"
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

func TestHandleSchema_WithSchemaContent(t *testing.T) {
	client := NewClient()
	opts := &domain.ChatOptions{SchemaContent: "{}"}
	err := client.HandleSchema(opts)
	if err != nil {
		t.Errorf("Expected no error when SchemaContent is set, but got %v", err)
	}
}

func TestSendStream_WithSchemaContent(t *testing.T) {
	client := NewClient()
	msgs := []*chat.ChatCompletionMessage{
		{Role: "user", Content: "Test message"},
	}
	schemaContent := `{"type": "object"}`
	opts := &domain.ChatOptions{
		Model:         "dry-run-model",
		SchemaContent: schemaContent,
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

	expectedSchemaLine := "SchemaContent: " + schemaContent

	foundSchemaContent := false
	for _, msg := range receivedMessages {
		if strings.Contains(msg, expectedSchemaLine) {
			foundSchemaContent = true
			break
		}
	}

	if !foundSchemaContent {
		t.Errorf("Expected dry run output to contain SchemaContent, but it didn't. Received messages: %v", receivedMessages)
	}
}
