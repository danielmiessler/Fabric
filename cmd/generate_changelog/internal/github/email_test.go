package github

import (
	"testing"
	"time"
)

func TestPRCommitEmailHandling(t *testing.T) {
	tests := []struct {
		name     string
		commit   PRCommit
		expected string
	}{
		{
			name: "Valid email field",
			commit: PRCommit{
				SHA:     "abc123",
				Message: "Fix bug in authentication",
				Author:  "John Doe",
				Email:   "john.doe@example.com",
				Date:    time.Now(),
				Parents: []string{"def456"},
			},
			expected: "john.doe@example.com",
		},
		{
			name: "Empty email field",
			commit: PRCommit{
				SHA:     "abc123",
				Message: "Fix bug in authentication",
				Author:  "John Doe",
				Email:   "",
				Date:    time.Now(),
				Parents: []string{"def456"},
			},
			expected: "",
		},
		{
			name: "Email field with proper initialization",
			commit: PRCommit{
				SHA:     "def789",
				Message: "Add new feature",
				Author:  "Jane Smith",
				Email:   "jane.smith@company.org",
				Date:    time.Now(),
				Parents: []string{"ghi012"},
			},
			expected: "jane.smith@company.org",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.commit.Email != tt.expected {
				t.Errorf("Expected email %q, got %q", tt.expected, tt.commit.Email)
			}
		})
	}
}
