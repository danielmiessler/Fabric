package changelog

import (
	"testing"
	"time"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/github"
)

func TestIsMergeCommit(t *testing.T) {
	tests := []struct {
		name     string
		commit   github.PRCommit
		expected bool
	}{
		{
			name: "Regular commit with single parent",
			commit: github.PRCommit{
				SHA:     "abc123",
				Message: "Fix bug in user authentication",
				Author:  "John Doe",
				Date:    time.Now(),
				Parents: []string{"def456"},
			},
			expected: false,
		},
		{
			name: "Merge commit with multiple parents",
			commit: github.PRCommit{
				SHA:     "abc123",
				Message: "Merge pull request #42 from feature/auth",
				Author:  "GitHub",
				Date:    time.Now(),
				Parents: []string{"def456", "ghi789"},
			},
			expected: true,
		},
		{
			name: "Merge commit detected by message pattern only",
			commit: github.PRCommit{
				SHA:     "abc123",
				Message: "Merge pull request #123 from user/feature-branch",
				Author:  "GitHub",
				Date:    time.Now(),
				Parents: []string{}, // Empty parents - fallback to message detection
			},
			expected: true,
		},
		{
			name: "Merge branch commit pattern",
			commit: github.PRCommit{
				SHA:     "abc123",
				Message: "Merge branch 'feature' into main",
				Author:  "Developer",
				Date:    time.Now(),
				Parents: []string{"def456"}, // Single parent but merge pattern
			},
			expected: true,
		},
		{
			name: "Regular commit with no merge patterns",
			commit: github.PRCommit{
				SHA:     "abc123",
				Message: "Add new feature for user management",
				Author:  "Jane Doe",
				Date:    time.Now(),
				Parents: []string{"def456"},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMergeCommit(tt.commit)
			if result != tt.expected {
				t.Errorf("isMergeCommit() = %v, expected %v for commit: %s",
					result, tt.expected, tt.commit.Message)
			}
		})
	}
}
