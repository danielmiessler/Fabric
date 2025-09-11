package cli

import (
	"context"
	"fmt"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
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
		return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("vendor_not_configured"), vendorName))
	}
	tr, ok := vendor.(transcriber)
	if !ok {
		return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("vendor_no_transcription_support"), vendorName))
	}
	model := flags.TranscribeModel
	if model == "" {
		return "", fmt.Errorf("%s", i18n.T("transcription_model_required"))
	}
	if message, err = tr.TranscribeFile(context.Background(), flags.TranscribeFile, model, flags.SplitMediaFile); err != nil {
		return
	}
	return
}
