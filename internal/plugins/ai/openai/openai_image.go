package openai

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/danielmiessler/fabric/internal/domain"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/responses"
)

// ImageGenerationResponseType is the type used for image generation calls in responses
const ImageGenerationResponseType = "image_generation_call"
const ImageGenerationToolType = "image_generation"

// ImageGenerationSupportedModels lists all models that support image generation
var ImageGenerationSupportedModels = []string{
	"gpt-4.1-nano",
	"o3",
	"gpt-5",
	"gpt-5-nano",
	"gpt-5.2",
}

func supportsImageGeneration(model string) bool {
	return slices.Contains(ImageGenerationSupportedModels, model)
}

func getOutputFormatFromExtension(imagePath string) string {
	if imagePath == "" {
		return "png"
	}

	ext := strings.ToLower(filepath.Ext(imagePath))
	switch ext {
	case ".png":
		return "png"
	case ".webp":
		return "webp"
	case ".jpg":
		return "jpeg"
	case ".jpeg":
		return "jpeg"
	default:
		return "png"
	}
}

func (o *Client) addImageGenerationTool(opts *domain.ChatOptions, tools []responses.ToolUnionParam) []responses.ToolUnionParam {
	if o.shouldUseImageGeneration(opts) {
		outputFormat := getOutputFormatFromExtension(opts.ImageFile)

		imageGenTool := responses.ToolUnionParam{
			OfImageGeneration: &responses.ToolImageGenerationParam{
				Type:         ImageGenerationToolType,
				Model:        "gpt-image-1",
				OutputFormat: outputFormat,
			},
		}

		if opts.ImageQuality != "" {
			imageGenTool.OfImageGeneration.Quality = opts.ImageQuality
		}

		if opts.ImageSize != "" {
			imageGenTool.OfImageGeneration.Size = opts.ImageSize
		}

		if opts.ImageBackground != "" {
			imageGenTool.OfImageGeneration.Background = opts.ImageBackground
		}

		// The API applies output_compression only to jpeg and webp.
		if opts.ImageCompression != 0 {
			imageGenTool.OfImageGeneration.OutputCompression = param.NewOpt(int64(opts.ImageCompression))
		}

		tools = append(tools, imageGenTool)
	}
	return tools
}

func (o *Client) shouldUseImageGeneration(opts *domain.ChatOptions) bool {
	return opts.ImageFile != ""
}

func (o *Client) extractAndSaveImages(resp *responses.Response, opts *domain.ChatOptions) error {
	if opts.ImageFile == "" {
		return nil
	}

	for _, item := range resp.Output {
		if item.Type == ImageGenerationResponseType {
			imageCall := item.AsImageGenerationCall()
			if imageCall.Status == "completed" && imageCall.Result != "" {
				imageData, err := base64.StdEncoding.DecodeString(imageCall.Result)
				if err != nil {
					return fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_image_failed_to_decode_image_data"), err))
				}

				dir := filepath.Dir(opts.ImageFile)
				if dir != "." {
					if err := os.MkdirAll(dir, 0755); err != nil {
						return fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_image_failed_to_create_directory"), dir, err))
					}
				}

				if err := os.WriteFile(opts.ImageFile, imageData, 0644); err != nil {
					return fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_image_failed_to_save_image"), opts.ImageFile, err))
				}

				fmt.Printf("%s\n", fmt.Sprintf(i18n.T("openai_image_saved_to"), opts.ImageFile))
				return nil
			}
		}
	}

	return nil
}
