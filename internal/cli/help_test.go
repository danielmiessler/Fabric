package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	openaivendor "github.com/danielmiessler/fabric/internal/plugins/ai/openai"
	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTranslatedHelpWriterGetOriginalDescription(t *testing.T) {
	t.Parallel()

	writer := NewTranslatedHelpWriter(flags.NewParser(&Flags{}, flags.None), &bytes.Buffer{})

	assert.Equal(t, "Run a named pipeline", writer.getOriginalDescription("pipeline"))
	assert.Empty(t, writer.getOriginalDescription("missing-flag"))
}

func TestTranslatedHelpWriterGetTranslatedDescription(t *testing.T) {
	t.Parallel()

	_, err := i18n.Init("en")
	require.NoError(t, err)

	writer := NewTranslatedHelpWriter(flags.NewParser(&Flags{}, flags.None), &bytes.Buffer{})

	assert.Equal(t, i18n.T("choose_pattern_from_available"), writer.getTranslatedDescription("pattern"))
	assert.Equal(t, "Run a named pipeline", writer.getTranslatedDescription("pipeline"))
	assert.Equal(t, i18n.T("no_description_available"), writer.getTranslatedDescription("missing-flag"))
}

func TestTranslatedHelpWriterWriteHelp(t *testing.T) {
	t.Parallel()

	_, err := i18n.Init("en")
	require.NoError(t, err)

	var buf bytes.Buffer
	parser := flags.NewParser(&Flags{}, flags.None)
	parser.Name = "fabric"

	writer := NewTranslatedHelpWriter(parser, &buf)
	writer.WriteHelp()

	output := buf.String()
	assert.Contains(t, output, i18n.T("usage_header"))
	assert.Contains(t, output, "fabric "+i18n.T("options_placeholder"))
	assert.Contains(t, output, i18n.T("application_options_header"))
	assert.Contains(t, output, "--pattern=")
	assert.Contains(t, output, "--stream")
	assert.NotContains(t, output, "--stream=")
	assert.Contains(t, output, "--temperature=")
	assert.Contains(t, output, "(default: 0.7)")
	assert.Contains(t, output, i18n.T("help_options_header"))
	assert.Contains(t, output, "--help")
}

func TestCustomHelpHandler(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	oldArgs := os.Args
	t.Cleanup(func() { os.Args = oldArgs })
	os.Args = []string{"fabric", "--language", "en"}

	var buf bytes.Buffer
	parser := flags.NewParser(&Flags{}, flags.None)
	parser.Name = "fabric"

	CustomHelpHandler(parser, &buf)

	output := buf.String()
	assert.NotEmpty(t, output)
	assert.True(t, strings.Contains(output, "--pattern="))
	assert.True(t, strings.Contains(output, "--help"))
}

func TestConfigureOpenAIResponsesAPI(t *testing.T) {
	openAIClient := openaivendor.NewClient()
	openAIClient.ImplementsResponses = false

	registry := &core.PluginRegistry{
		VendorsAll: ai.NewVendorsManager(),
	}
	registry.VendorsAll.AddVendors(openAIClient, &cliStubVendor{name: "Anthropic"})

	configureOpenAIResponsesAPI(registry, false)
	assert.True(t, openAIClient.ImplementsResponses)

	configureOpenAIResponsesAPI(registry, true)
	assert.False(t, openAIClient.ImplementsResponses)

	configureOpenAIResponsesAPI(nil, false)
	configureOpenAIResponsesAPI(&core.PluginRegistry{}, false)
}
