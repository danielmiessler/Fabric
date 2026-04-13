package fsdb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestPatternsEntity(t *testing.T) (*PatternsEntity, func()) {
	// Create a temporary directory for test patterns
	tmpDir, err := os.MkdirTemp("", "test-patterns-*")
	require.NoError(t, err)

	entity := &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       tmpDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}

	// Return cleanup function
	cleanup := func() {
		os.RemoveAll(tmpDir)
	}

	return entity, cleanup
}

// Helper to create a test pattern file
func createTestPattern(t *testing.T, entity *PatternsEntity, name, content string) {
	patternDir := filepath.Join(entity.Dir, name)
	err := os.MkdirAll(patternDir, 0755)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(patternDir, entity.SystemPatternFile), []byte(content), 0644)
	require.NoError(t, err)
}

func TestApplyVariables(t *testing.T) {
	entity := &PatternsEntity{}

	tests := []struct {
		name      string
		pattern   *Pattern
		variables map[string]string
		input     string
		want      string
		wantErr   bool
	}{
		{
			name: "pattern with explicit input placement",
			pattern: &Pattern{
				Pattern: "You are a {{role}}.\n{{input}}\nPlease analyze.",
			},
			variables: map[string]string{
				"role": "security expert",
			},
			input: "Check this code",
			want:  "You are a security expert.\nCheck this code\nPlease analyze.",
		},
		{
			name: "pattern without input variable gets input appended",
			pattern: &Pattern{
				Pattern: "You are a {{role}}.\nPlease analyze.",
			},
			variables: map[string]string{
				"role": "code reviewer",
			},
			input: "Review this PR",
			want:  "You are a code reviewer.\nPlease analyze.\nReview this PR",
		},
		// ... previous test cases ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := entity.applyVariables(tt.pattern, tt.variables, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, tt.pattern.Pattern)
		})
	}
}

func TestGetApplyVariables(t *testing.T) {
	entity, cleanup := setupTestPatternsEntity(t)
	defer cleanup()

	// Create a test pattern
	createTestPattern(t, entity, "test-pattern", "You are a {{role}}.\n{{input}}")
	createTestPattern(t, entity, "frontmatter-pattern", `---
title: "{{role}} review"
tags:
  - "{{role}}"
  - inbox
summary: "{{input}}"
---
You are a {{role}}.
{{input}}`)

	tests := []struct {
		name            string
		source          string
		variables       map[string]string
		input           string
		want            string
		wantFrontmatter map[string]any
		wantErr         bool
	}{
		{
			name:   "basic pattern with variables and input",
			source: "test-pattern",
			variables: map[string]string{
				"role": "reviewer",
			},
			input: "check this code",
			want:  "You are a reviewer.\ncheck this code",
		},
		{
			name:   "pattern with frontmatter resolves body and metadata",
			source: "frontmatter-pattern",
			variables: map[string]string{
				"role": "reviewer",
			},
			input: "check this code",
			want:  "You are a reviewer.\ncheck this code",
			wantFrontmatter: map[string]any{
				"title":   "reviewer review",
				"tags":    []any{"reviewer", "inbox"},
				"summary": "check this code",
			},
		},
		{
			name:      "pattern with missing variable",
			source:    "test-pattern",
			variables: map[string]string{},
			input:     "test input",
			wantErr:   true,
		},
		{
			name:    "non-existent pattern",
			source:  "non-existent",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := entity.GetApplyVariables(tt.source, tt.variables, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, result.Pattern)
			assert.Equal(t, tt.wantFrontmatter, result.Frontmatter)
		})
	}
}

func TestGetWithoutVariables(t *testing.T) {
	entity, cleanup := setupTestPatternsEntity(t)
	defer cleanup()

	createTestPattern(t, entity, "test-pattern", `---
summary: "{{input}}"
---
Prefix {{input}} {{roam}}`)

	result, err := entity.GetWithoutVariables("test-pattern", "hello")
	require.NoError(t, err)
	assert.Equal(t, "Prefix hello {{roam}}", result.Pattern)
	assert.Equal(t, map[string]any{"summary": "hello"}, result.Frontmatter)

	createTestPattern(t, entity, "no-input", "Static content")
	result, err = entity.GetWithoutVariables("no-input", "hi")
	require.NoError(t, err)
	assert.Equal(t, "Static content\nhi", result.Pattern)
}

func TestPatternsEntity_Save(t *testing.T) {
	entity, cleanup := setupTestPatternsEntity(t)
	defer cleanup()

	name := "new-pattern"
	content := []byte("test pattern content")
	require.NoError(t, entity.Save(name, content))

	patternDir := filepath.Join(entity.Dir, name)
	info, err := os.Stat(patternDir)
	require.NoError(t, err)
	assert.True(t, info.IsDir())

	data, err := os.ReadFile(filepath.Join(patternDir, entity.SystemPatternFile))
	require.NoError(t, err)
	assert.Equal(t, content, data)
}

func TestPatternsEntity_CustomPatterns(t *testing.T) {
	// Create main patterns directory
	mainDir, err := os.MkdirTemp("", "test-main-patterns-*")
	require.NoError(t, err)
	defer os.RemoveAll(mainDir)

	// Create custom patterns directory
	customDir, err := os.MkdirTemp("", "test-custom-patterns-*")
	require.NoError(t, err)
	defer os.RemoveAll(customDir)

	entity := &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       mainDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
		CustomPatternsDir: customDir,
	}

	// Create a pattern in main directory
	createTestPattern(t, &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       mainDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}, "main-pattern", "Main pattern content")

	// Create a pattern in custom directory
	createTestPattern(t, &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       customDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}, "custom-pattern", "Custom pattern content")

	// Create a pattern with same name in both directories (custom should override)
	createTestPattern(t, &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       mainDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}, "shared-pattern", "Main shared pattern")

	createTestPattern(t, &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       customDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}, "shared-pattern", `---
title: "Custom"
---
Custom shared pattern`)

	// Test GetNames includes both directories
	names, err := entity.GetNames()
	require.NoError(t, err)
	assert.Contains(t, names, "main-pattern")
	assert.Contains(t, names, "custom-pattern")
	assert.Contains(t, names, "shared-pattern")

	// Test that custom pattern overrides main pattern
	pattern, err := entity.getFromDB("shared-pattern")
	require.NoError(t, err)
	assert.Equal(t, "Custom shared pattern", pattern.Pattern)
	assert.Equal(t, map[string]any{"title": "Custom"}, pattern.Frontmatter)

	// Test that main pattern is accessible when not overridden
	pattern, err = entity.getFromDB("main-pattern")
	require.NoError(t, err)
	assert.Equal(t, "Main pattern content", pattern.Pattern)

	// Test GetRaw also respects custom patterns directory
	rawPattern, err := entity.GetRaw("shared-pattern")
	require.NoError(t, err)
	assert.Equal(t, "Custom shared pattern", rawPattern.Pattern)
	assert.Equal(t, map[string]any{"title": "Custom"}, rawPattern.Frontmatter)

	// Test that custom pattern is accessible
	pattern, err = entity.getFromDB("custom-pattern")
	require.NoError(t, err)
	assert.Equal(t, "Custom pattern content", pattern.Pattern)
}

func TestGetRawFrontmatterUnresolvedExact(t *testing.T) {
	entity, cleanup := setupTestPatternsEntity(t)
	defer cleanup()

	createTestPattern(t, entity, "frontmatter-pattern", `---
title: "{{role}} review"
summary: "{{input}}"
---
Review {{input}}`)

	result, err := entity.GetRaw("frontmatter-pattern")
	require.NoError(t, err)
	assert.Equal(t, "Review {{input}}", result.Pattern)
	assert.Equal(t, map[string]any{
		"title":   "{{role}} review",
		"summary": "{{input}}",
	}, result.Frontmatter)
}

func TestInvalidFrontmatterReturnsError(t *testing.T) {
	entity, cleanup := setupTestPatternsEntity(t)
	defer cleanup()

	createTestPattern(t, entity, "invalid-frontmatter", `---
title: [unterminated
---
Body`)

	_, err := entity.GetRaw("invalid-frontmatter")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid YAML frontmatter")
}

func TestEmptyFrontmatterBlock(t *testing.T) {
	entity, cleanup := setupTestPatternsEntity(t)
	defer cleanup()

	createTestPattern(t, entity, "empty-frontmatter", `---
---
Body`)

	result, err := entity.GetRaw("empty-frontmatter")
	require.NoError(t, err)
	assert.Equal(t, "Body", result.Pattern)
	assert.Nil(t, result.Frontmatter)
}

func TestFilePatternFrontmatter(t *testing.T) {
	entity := &PatternsEntity{}
	tmpDir, err := os.MkdirTemp("", "test-pattern-file-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	patternPath := filepath.Join(tmpDir, "pattern.md")
	err = os.WriteFile(patternPath, []byte(`---
title: "{{role}}"
---
Body {{input}}`), 0644)
	require.NoError(t, err)

	result, err := entity.GetApplyVariables(patternPath, map[string]string{"role": "reviewer"}, "hello")
	require.NoError(t, err)
	assert.Equal(t, "Body hello", result.Pattern)
	assert.Equal(t, map[string]any{"title": "reviewer"}, result.Frontmatter)
}

func TestPatternsEntity_CustomPatternsEmpty(t *testing.T) {
	// Test behavior when custom patterns directory is empty or doesn't exist
	mainDir, err := os.MkdirTemp("", "test-main-patterns-*")
	require.NoError(t, err)
	defer os.RemoveAll(mainDir)

	entity := &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       mainDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
		CustomPatternsDir: "/nonexistent/directory",
	}

	// Create a pattern in main directory
	createTestPattern(t, &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       mainDir,
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}, "main-pattern", "Main pattern content")

	// Test GetNames works even with nonexistent custom directory
	names, err := entity.GetNames()
	require.NoError(t, err)
	assert.Contains(t, names, "main-pattern")

	// Test that main pattern is accessible
	pattern, err := entity.getFromDB("main-pattern")
	require.NoError(t, err)
	assert.Equal(t, "Main pattern content", pattern.Pattern)
}
