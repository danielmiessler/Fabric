package geminicommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genai"
)

func TestExtractTextSkipsThoughtParts(t *testing.T) {
	response := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{{
			Content: &genai.Content{Parts: []*genai.Part{
				{Text: "**Planning the answer**", Thought: true},
				{Text: "The answer is 2."},
			}},
		}},
	}

	assert.Equal(t, "The answer is 2.", ExtractText(response))
	assert.Equal(t, "The answer is 2.", ExtractTextWithCitations(response))
}

func TestOutputTokenCountIncludesThoughts(t *testing.T) {
	usage := &genai.GenerateContentResponseUsageMetadata{
		PromptTokenCount:     10,
		CandidatesTokenCount: 20,
		ThoughtsTokenCount:   70,
		TotalTokenCount:      100,
	}

	assert.Equal(t, 90, OutputTokenCount(usage))
}
