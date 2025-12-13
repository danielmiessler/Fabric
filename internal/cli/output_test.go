package cli

import (
	"os"
	"testing"
)

func TestCopyToClipboard(t *testing.T) {
	t.Skip("skipping test, because of docker env. in ci.")

	message := "test message"
	err := CopyToClipboard(message)
	if err != nil {
		t.Fatalf("CopyToClipboard() error = %v", err)
	}
}

func TestCreateOutputFile(t *testing.T) {

	fileName := "test_output.txt"
	message := "test message"
	err := CreateOutputFile(message, fileName)
	if err != nil {
		t.Fatalf("CreateOutputFile() error = %v", err)
	}

	t.Cleanup(func() { os.Remove(fileName) })

	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	expected := message + "\n"
	if string(data) != expected {
		t.Fatalf("expected file contents %q, got %q", expected, data)
	}
}

func TestCreateOutputFileMessageWithTrailingNewline(t *testing.T) {
	fileName := "test_output_with_newline.txt"
	message := "test message with newline\n"

	if err := CreateOutputFile(message, fileName); err != nil {
		t.Fatalf("CreateOutputFile() error = %v", err)
	}
	t.Cleanup(func() { os.Remove(fileName) })

	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	if string(data) != message {
		t.Fatalf("expected file contents %q, got %q", message, data)
	}
}
