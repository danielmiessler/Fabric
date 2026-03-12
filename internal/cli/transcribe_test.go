package cli

import (
	"errors"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/ai"
	"github.com/stretchr/testify/require"
)

func TestHandleTranscriptionVendorNotConfigured(t *testing.T) {
	registry := &core.PluginRegistry{VendorManager: ai.NewVendorsManager()}

	message, err := handleTranscription(&Flags{
		Vendor:          "MissingVendor",
		TranscribeFile:  "audio.mp3",
		TranscribeModel: "gpt-4o-mini-transcribe",
	}, registry)
	require.Error(t, err)
	require.Empty(t, message)
	require.Contains(t, err.Error(), "MissingVendor")
}

func TestHandleTranscriptionVendorWithoutSupport(t *testing.T) {
	manager := ai.NewVendorsManager()
	manager.AddVendors(&cliStubVendor{name: "NoTranscribe"})

	registry := &core.PluginRegistry{VendorManager: manager}

	message, err := handleTranscription(&Flags{
		Vendor:          "NoTranscribe",
		TranscribeFile:  "audio.mp3",
		TranscribeModel: "gpt-4o-mini-transcribe",
	}, registry)
	require.Error(t, err)
	require.Empty(t, message)
	require.Contains(t, err.Error(), "NoTranscribe")
	require.Contains(t, err.Error(), "does not support audio transcription")
}

func TestHandleTranscriptionRequiresModel(t *testing.T) {
	manager := ai.NewVendorsManager()
	manager.AddVendors(&cliStubTranscriberVendor{cliStubVendor: cliStubVendor{name: "OpenAI"}})

	registry := &core.PluginRegistry{VendorManager: manager}

	message, err := handleTranscription(&Flags{
		TranscribeFile: "audio.mp3",
	}, registry)
	require.Error(t, err)
	require.Empty(t, message)
	require.Contains(t, err.Error(), "transcription model is required")
}

func TestHandleTranscriptionUsesDefaultVendorAndPassesArguments(t *testing.T) {
	vendor := &cliStubTranscriberVendor{
		cliStubVendor: cliStubVendor{name: "OpenAI"},
		transcription: "transcribed content",
	}

	manager := ai.NewVendorsManager()
	manager.AddVendors(vendor)

	registry := &core.PluginRegistry{VendorManager: manager}

	message, err := handleTranscription(&Flags{
		TranscribeFile:  "audio.mp3",
		TranscribeModel: "whisper-1",
		SplitMediaFile:  true,
	}, registry)
	require.NoError(t, err)
	require.Equal(t, "transcribed content", message)
	require.Equal(t, "audio.mp3", vendor.lastFile)
	require.Equal(t, "whisper-1", vendor.lastModel)
	require.True(t, vendor.lastSplit)
	require.False(t, vendor.lastNilCtx)
}

func TestHandleTranscriptionPropagatesVendorErrors(t *testing.T) {
	vendor := &cliStubTranscriberVendor{
		cliStubVendor: cliStubVendor{name: "OpenAI"},
		transcribeErr: errors.New("vendor transcription failed"),
	}

	manager := ai.NewVendorsManager()
	manager.AddVendors(vendor)

	registry := &core.PluginRegistry{VendorManager: manager}

	message, err := handleTranscription(&Flags{
		TranscribeFile:  "audio.mp3",
		TranscribeModel: "whisper-1",
	}, registry)
	require.Error(t, err)
	require.Empty(t, message)
	require.Contains(t, err.Error(), "vendor transcription failed")
}
