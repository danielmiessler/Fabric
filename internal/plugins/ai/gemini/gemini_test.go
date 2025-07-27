package gemini

import (
	"testing"

	"google.golang.org/genai"
)

// Test buildModelNameFull method
func TestBuildModelNameFull(t *testing.T) {
	client := &Client{}

	tests := []struct {
		input    string
		expected string
	}{
		{"chat-bison-001", "models/chat-bison-001"},
		{"models/chat-bison-001", "models/chat-bison-001"},
		{"gemini-2.5-flash-preview-tts", "models/gemini-2.5-flash-preview-tts"},
	}

	for _, test := range tests {
		result := client.buildModelNameFull(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v, got %v", test.input, test.expected, result)
		}
	}
}

// Test extractTextFromResponse method
func TestExtractTextFromResponse(t *testing.T) {
	client := &Client{}
	response := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []*genai.Part{
						{Text: "Hello, "},
						{Text: "world!"},
					},
				},
			},
		},
	}
	expected := "Hello, world!"

	result := client.extractTextFromResponse(response)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Test isTTSModel method
func TestIsTTSModel(t *testing.T) {
	client := &Client{}

	tests := []struct {
		modelName string
		expected  bool
	}{
		{"gemini-2.5-flash-preview-tts", true},
		{"text-to-speech-model", true},
		{"TTS-MODEL", true},
		{"gemini-pro", false},
		{"chat-bison", false},
		{"", false},
	}

	for _, test := range tests {
		result := client.isTTSModel(test.modelName)
		if result != test.expected {
			t.Errorf("For model %v, expected %v, got %v", test.modelName, test.expected, result)
		}
	}
}

// Test generateWAVFile method (basic test)
func TestGenerateWAVFile(t *testing.T) {
	client := &Client{}

	// Test with minimal PCM data
	pcmData := []byte{0x00, 0x01, 0x02, 0x03}

	result, err := client.generateWAVFile(pcmData)
	if err != nil {
		t.Errorf("generateWAVFile failed: %v", err)
	}

	// Check that we got some data back
	if len(result) == 0 {
		t.Error("generateWAVFile returned empty data")
	}

	// Check that it starts with RIFF header
	if len(result) >= 4 && string(result[0:4]) != "RIFF" {
		t.Error("Generated WAV data doesn't start with RIFF header")
	}
}
