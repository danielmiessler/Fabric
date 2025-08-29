package cli

import (
	"context"
	"fmt"

	"github.com/danielmiessler/fabric/internal/core"
)

type transcriber interface {
	TranscribeFile(ctx context.Context, filePath, model string, split bool) (string, error)
}

func handleTranscription(flags *Flags, registry *core.PluginRegistry) (message string, err error) {
	vendorName := flags.Vendor
	if vendorName == "" {
		vendorName = "OpenAI"
	}
	vendor, ok := registry.VendorManager.VendorsByName[vendorName]
	if !ok {
		return "", fmt.Errorf("vendor %s not configured", vendorName)
	}
	tr, ok := vendor.(transcriber)
	if !ok {
		return "", fmt.Errorf("vendor %s does not support audio transcription", vendorName)
	}
	model := flags.TranscribeModel
	if model == "" {
		return "", fmt.Errorf("transcription model is required (use --transcribe-model)")
	}
	if message, err = tr.TranscribeFile(context.Background(), flags.TranscribeFile, model, flags.SplitMediaFile); err != nil {
		return
	}
	return
}
