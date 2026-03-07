package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/pipeline"
)

var pipelineRunOptions = func(invocationDir string) pipeline.RunOptions {
	return pipeline.RunOptions{InvocationDir: invocationDir}
}

func handlePipelineCommands(currentFlags *Flags, registry *core.PluginRegistry) (handled bool, err error) {
	if currentFlags.ValidatePipeline == "" && currentFlags.Pipeline == "" {
		return false, nil
	}

	if currentFlags.ValidatePipeline != "" && currentFlags.Pipeline != "" {
		return true, fmt.Errorf("--validate-pipeline and --pipeline cannot be used together")
	}

	loader, err := pipeline.NewDefaultLoader()
	if err != nil {
		return true, err
	}

	if currentFlags.ValidatePipeline != "" {
		var pipe *pipeline.Pipeline
		if pipe, err = loader.LoadFilePath(currentFlags.ValidatePipeline); err != nil {
			return true, err
		}
		if err = pipeline.Preflight(context.Background(), pipe, pipeline.PreflightOptions{Registry: registry}); err != nil {
			return true, err
		}
		fmt.Printf("pipeline valid: %s\n", pipe.Name)
		return true, nil
	}

	if currentFlags.Pattern != "" || currentFlags.Context != "" || currentFlags.Session != "" || len(currentFlags.Attachments) > 0 || currentFlags.argumentMessage != "" {
		return true, fmt.Errorf("--pipeline cannot be combined with pattern/chat-style inputs")
	}

	var pipe *pipeline.Pipeline
	if pipe, err = loader.LoadNamed(currentFlags.Pipeline); err != nil {
		return true, err
	}
	if err = pipeline.Preflight(context.Background(), pipe, pipeline.PreflightOptions{Registry: registry}); err != nil {
		return true, err
	}

	if currentFlags.ValidateOnly {
		fmt.Printf("pipeline valid: %s\n", pipe.Name)
		return true, nil
	}

	var source pipeline.RunSource
	if source, err = resolvePipelineSource(currentFlags, registry); err != nil {
		return true, err
	}

	invocationDir, err := os.Getwd()
	if err != nil {
		return true, fmt.Errorf("determine current working directory: %w", err)
	}

	runner := pipeline.NewRunner(os.Stdout, os.Stderr, registry)
	result, err := runner.Run(context.Background(), pipe, source, pipelineRunOptions(invocationDir))
	if currentFlags.Output != "" && result != nil && result.FinalOutput != "" {
		if outputErr := CreateOutputFile(result.FinalOutput, currentFlags.Output); outputErr != nil {
			if err != nil {
				return true, fmt.Errorf("%w; output file error: %w", err, outputErr)
			}
			return true, outputErr
		}
	}
	if err != nil {
		return true, err
	}

	return true, nil
}

func resolvePipelineSource(currentFlags *Flags, registry *core.PluginRegistry) (pipeline.RunSource, error) {
	sourceCount := 0
	if currentFlags.stdinProvided {
		sourceCount++
	}
	if currentFlags.Source != "" {
		sourceCount++
	}
	if currentFlags.ScrapeURL != "" {
		sourceCount++
	}

	if sourceCount != 1 {
		return pipeline.RunSource{}, fmt.Errorf("--pipeline requires exactly one source: stdin, --source, or --scrape_url")
	}

	if currentFlags.stdinProvided {
		return pipeline.RunSource{
			Mode:    pipeline.SourceModeStdin,
			Payload: currentFlags.stdinMessage,
		}, nil
	}

	if currentFlags.ScrapeURL != "" {
		if registry == nil || registry.Jina == nil || !registry.Jina.IsConfigured() {
			return pipeline.RunSource{}, fmt.Errorf("%s", i18n.T("scraping_not_configured"))
		}
		scrapedContent, err := registry.Jina.ScrapeURL(currentFlags.ScrapeURL)
		if err != nil {
			return pipeline.RunSource{}, err
		}
		return pipeline.RunSource{
			Mode:      pipeline.SourceModeScrapeURL,
			Reference: currentFlags.ScrapeURL,
			Payload:   scrapedContent,
		}, nil
	}

	absPath, err := filepath.Abs(currentFlags.Source)
	if err != nil {
		return pipeline.RunSource{}, fmt.Errorf("resolve source path %s: %w", currentFlags.Source, err)
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return pipeline.RunSource{}, fmt.Errorf("stat source path %s: %w", absPath, err)
	}

	if info.IsDir() {
		return pipeline.RunSource{
			Mode:      pipeline.SourceModeSource,
			Reference: absPath,
			Payload:   absPath,
		}, nil
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return pipeline.RunSource{}, fmt.Errorf("read source file %s: %w", absPath, err)
	}

	return pipeline.RunSource{
		Mode:      pipeline.SourceModeSource,
		Reference: absPath,
		Payload:   string(content),
	}, nil
}
