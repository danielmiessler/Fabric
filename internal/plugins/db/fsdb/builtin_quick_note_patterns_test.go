package fsdb

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func builtinPatternsEntity(t *testing.T) *PatternsEntity {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok)

	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(file), "../../../../"))
	return &PatternsEntity{
		StorageEntity: &StorageEntity{
			Dir:       filepath.Join(repoRoot, "data", "patterns"),
			Label:     "patterns",
			ItemIsDir: true,
		},
		SystemPatternFile: "system.md",
	}
}

func TestBuiltinQuickNotePatternsAreDiscoverable(t *testing.T) {
	entity := builtinPatternsEntity(t)

	names, err := entity.GetNames()
	require.NoError(t, err)

	assert.Contains(t, names, "techNote")
	assert.Contains(t, names, "nontechNote")
}

func TestBuiltinTechNotePatternContract(t *testing.T) {
	entity := builtinPatternsEntity(t)

	pattern, err := entity.GetRaw("techNote")
	require.NoError(t, err)

	assert.Contains(t, pattern.Pattern, "{{input}}")
	assert.Contains(t, pattern.Pattern, "## 🧠 Session Focus")
	assert.Contains(t, pattern.Pattern, "## ✅ Learning Outcomes")
	assert.Contains(t, pattern.Pattern, "## 🧭 Topic Index")
	assert.Contains(t, pattern.Pattern, "## 🗺️ Conceptual Roadmap")
	assert.Contains(t, pattern.Pattern, "## 💻 Coding Walkthroughs")
	assert.Contains(t, pattern.Pattern, "## 🧩 HOTS (High-Order Thinking)")
	assert.Contains(t, pattern.Pattern, "output one learner-facing Markdown note only")
	assert.Contains(t, pattern.Pattern, "do not emit pipeline manifests")
	assert.Contains(t, pattern.Pattern, "do not describe stages or your process")
}

func TestBuiltinNonTechNotePatternContract(t *testing.T) {
	entity := builtinPatternsEntity(t)

	pattern, err := entity.GetRaw("nontechNote")
	require.NoError(t, err)

	assert.Contains(t, pattern.Pattern, "{{input}}")
	assert.Contains(t, pattern.Pattern, "## 🧠 Session Focus")
	assert.Contains(t, pattern.Pattern, "## ✅ What To Understand")
	assert.Contains(t, pattern.Pattern, "## 📚 Core Ideas")
	assert.Contains(t, pattern.Pattern, "## 🌱 Reflection Prompts")
	assert.Contains(t, pattern.Pattern, "do not assume code, formulas, or equations")
	assert.Contains(t, pattern.Pattern, "output one learner-facing Markdown note only")
	assert.Contains(t, pattern.Pattern, "do not emit pipeline manifests")
	assert.Contains(t, pattern.Pattern, "do not describe stages or your process")
}

func TestBuiltinQuickNotePatternsApplyInput(t *testing.T) {
	entity := builtinPatternsEntity(t)

	pattern, err := entity.GetWithoutVariables("techNote", "sample transcript")
	require.NoError(t, err)

	assert.Contains(t, pattern.Pattern, "sample transcript")
	assert.NotContains(t, pattern.Pattern, "{{input}}")
}
