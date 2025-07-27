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

	defer os.Remove(fileName)
}
