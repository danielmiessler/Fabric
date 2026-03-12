package cli

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/stretchr/testify/require"
)

func TestHasPipelineRuntimeOnlyFlags(t *testing.T) {
	t.Parallel()

	require.False(t, hasPipelineRuntimeOnlyFlags(&Flags{}))
	require.True(t, hasPipelineRuntimeOnlyFlags(&Flags{FromStage: "prepare"}))
	require.True(t, hasPipelineRuntimeOnlyFlags(&Flags{ToStage: "publish"}))
	require.True(t, hasPipelineRuntimeOnlyFlags(&Flags{OnlyStage: "render"}))
	require.True(t, hasPipelineRuntimeOnlyFlags(&Flags{PipelineEventsJSON: true}))
}

func TestResolvePipelineSource(t *testing.T) {
	t.Run("accepts stdin source", func(t *testing.T) {
		source, err := resolvePipelineSource(&Flags{
			stdinProvided: true,
			stdinMessage:  "stdin payload",
		}, nil)
		require.NoError(t, err)
		require.Equal(t, "stdin", string(source.Mode))
		require.Equal(t, "stdin payload", source.Payload)
	})

	t.Run("accepts file source", func(t *testing.T) {
		tempDir := t.TempDir()
		sourcePath := filepath.Join(tempDir, "input.txt")
		require.NoError(t, os.WriteFile(sourcePath, []byte("file payload"), 0o644))

		source, err := resolvePipelineSource(&Flags{Source: sourcePath}, nil)
		require.NoError(t, err)
		require.Equal(t, "source", string(source.Mode))
		require.Equal(t, "file payload", source.Payload)
		require.Equal(t, sourcePath, source.Reference)
	})

	t.Run("accepts directory source", func(t *testing.T) {
		sourceDir := t.TempDir()

		source, err := resolvePipelineSource(&Flags{Source: sourceDir}, nil)
		require.NoError(t, err)
		require.Equal(t, "source", string(source.Mode))
		require.Equal(t, sourceDir, source.Payload)
		require.Equal(t, sourceDir, source.Reference)
	})

	t.Run("rejects ambiguous source selection", func(t *testing.T) {
		_, err := resolvePipelineSource(&Flags{
			stdinProvided: true,
			stdinMessage:  "stdin payload",
			Source:        "input.txt",
		}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "requires exactly one source")
	})

	t.Run("rejects missing source selection", func(t *testing.T) {
		_, err := resolvePipelineSource(&Flags{}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "requires exactly one source")
	})

	t.Run("rejects scrape without configured registry", func(t *testing.T) {
		_, err := resolvePipelineSource(&Flags{ScrapeURL: "https://example.com"}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "scraping functionality is not configured")
	})

	t.Run("accepts scrape url source", func(t *testing.T) {
		registry := newToolProcessingRegistry(t)
		installHTTPTransport(t, func(req *http.Request) (*http.Response, error) {
			require.Equal(t, "https://r.jina.ai/https://example.com/article", req.URL.String())
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("scraped article")),
				Request:    req,
			}, nil
		})

		source, err := resolvePipelineSource(&Flags{ScrapeURL: "https://example.com/article"}, registry)
		require.NoError(t, err)
		require.Equal(t, "scrape_url", string(source.Mode))
		require.Equal(t, "https://example.com/article", source.Reference)
		require.Equal(t, "scraped article", source.Payload)
	})

	t.Run("returns stat error for missing source file", func(t *testing.T) {
		_, err := resolvePipelineSource(&Flags{Source: filepath.Join(t.TempDir(), "missing.txt")}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "stat source path")
	})
}

func TestAssignWithConversion(t *testing.T) {
	t.Parallel()

	var intValue int
	require.NoError(t, assignWithConversion(reflect.ValueOf(&intValue).Elem(), reflect.ValueOf("42.9")))
	require.Equal(t, 42, intValue)

	var floatValue float64
	require.NoError(t, assignWithConversion(reflect.ValueOf(&floatValue).Elem(), reflect.ValueOf("3.14")))
	require.Equal(t, 3.14, floatValue)

	var boolValue bool
	require.NoError(t, assignWithConversion(reflect.ValueOf(&boolValue).Elem(), reflect.ValueOf("true")))
	require.True(t, boolValue)

	var stringValue string
	err := assignWithConversion(reflect.ValueOf(&stringValue).Elem(), reflect.ValueOf(123))
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported conversion")
}

func TestReadStdinUnexported(t *testing.T) {
	original := os.Stdin
	reader, writer, err := os.Pipe()
	require.NoError(t, err)
	os.Stdin = reader
	t.Cleanup(func() {
		os.Stdin = original
		_ = reader.Close()
	})

	_, err = writer.WriteString("stdin content without trailing newline")
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	content, err := readStdin()
	require.NoError(t, err)
	require.Equal(t, "stdin content without trailing newline", content)
}

func TestBuildChatRequest(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	attachmentPath := filepath.Join(tempDir, "pixel.png")
	require.NoError(t, os.WriteFile(attachmentPath, onePixelPNG(t), 0o644))

	flags := &Flags{
		Context:          "example-context",
		Session:          "example-session",
		Pattern:          "example-pattern",
		Strategy:         "strategy-a",
		PatternVariables: map[string]string{"topic": "notes"},
		Attachments:      []string{attachmentPath},
		Message:          "  hello world  ",
		Language:         "en-us",
		InputHasVars:     true,
	}

	request, err := flags.BuildChatRequest("meta-info")
	require.NoError(t, err)
	require.Equal(t, "example-context", request.ContextName)
	require.Equal(t, "example-session", request.SessionName)
	require.Equal(t, "example-pattern", request.PatternName)
	require.Equal(t, "strategy-a", request.StrategyName)
	require.Equal(t, map[string]string{"topic": "notes"}, request.PatternVariables)
	require.True(t, request.InputHasVars)
	require.Equal(t, "meta-info", request.Meta)
	require.Equal(t, "en-US", request.Language)

	require.NotNil(t, request.Message)
	require.Len(t, request.Message.MultiContent, 2)
	require.Equal(t, chat.ChatMessagePartTypeText, request.Message.MultiContent[0].Type)
	require.Equal(t, "hello world", request.Message.MultiContent[0].Text)
	require.Equal(t, chat.ChatMessagePartTypeImageURL, request.Message.MultiContent[1].Type)
	require.NotNil(t, request.Message.MultiContent[1].ImageURL)
	require.True(t, strings.HasPrefix(request.Message.MultiContent[1].ImageURL.URL, "data:image/png;base64,"))
}

func TestAppendMessage(t *testing.T) {
	t.Parallel()

	require.Equal(t, "next", AppendMessage("", "next"))
	require.Equal(t, "first\nsecond", AppendMessage("first", "second"))
}

func TestIsChatRequest(t *testing.T) {
	t.Parallel()

	require.False(t, (&Flags{}).IsChatRequest())
	require.True(t, (&Flags{Message: "hello"}).IsChatRequest())
	require.True(t, (&Flags{Attachments: []string{"image.png"}}).IsChatRequest())
	require.True(t, (&Flags{Context: "ctx"}).IsChatRequest())
	require.True(t, (&Flags{Session: "session"}).IsChatRequest())
	require.True(t, (&Flags{Pattern: "pattern"}).IsChatRequest())
}

func TestFlagsWriteOutput(t *testing.T) {
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "flags-output.txt")

	flags := &Flags{Output: outputPath}
	stdout, err := captureStdout(func() error {
		return flags.WriteOutput("hello from flags")
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "hello from flags")

	content, readErr := os.ReadFile(outputPath)
	require.NoError(t, readErr)
	require.Equal(t, "hello from flags\n", string(content))
}

func TestPackageWriteOutput(t *testing.T) {
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "package-output.txt")

	stdout, err := captureStdout(func() error {
		return WriteOutput("hello from package", outputPath)
	})
	require.NoError(t, err)
	require.Contains(t, stdout, "hello from package")

	content, readErr := os.ReadFile(outputPath)
	require.NoError(t, readErr)
	require.Equal(t, "hello from package\n", string(content))
}

func onePixelPNG(t *testing.T) []byte {
	t.Helper()

	data, err := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO7Z0ioAAAAASUVORK5CYII=")
	require.NoError(t, err)
	return data
}
