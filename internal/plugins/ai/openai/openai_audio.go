package openai

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	debuglog "github.com/danielmiessler/fabric/internal/log"

	openai "github.com/openai/openai-go"
)

// MaxAudioFileSize defines the maximum allowed size for audio uploads (25MB).
const MaxAudioFileSize int64 = 25 * 1024 * 1024

// AllowedTranscriptionModels lists the models supported for transcription.
var AllowedTranscriptionModels = []string{
	string(openai.AudioModelWhisper1),
	string(openai.AudioModelGPT4oMiniTranscribe),
	string(openai.AudioModelGPT4oTranscribe),
}

// allowedAudioExtensions defines the supported input file extensions.
var allowedAudioExtensions = map[string]struct{}{
	".mp3":  {},
	".mp4":  {},
	".mpeg": {},
	".mpga": {},
	".m4a":  {},
	".wav":  {},
	".webm": {},
}

// TranscribeFile transcribes the given audio file using the specified model. If the file
// exceeds the size limit, it can optionally be split into chunks using ffmpeg.
func (o *Client) TranscribeFile(ctx context.Context, filePath, model string, split bool) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if !slices.Contains(AllowedTranscriptionModels, model) {
		return "", fmt.Errorf("model '%s' is not supported for transcription", model)
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	if _, ok := allowedAudioExtensions[ext]; !ok {
		return "", fmt.Errorf("unsupported audio format '%s'", ext)
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	var files []string
	var cleanup func()
	if info.Size() > MaxAudioFileSize {
		if !split {
			return "", fmt.Errorf("file %s exceeds 25MB limit; use --split-media-file to enable automatic splitting", filePath)
		}
		debuglog.Log("File %s is larger than the size limit... breaking it up into chunks...\n", filePath)
		if files, cleanup, err = splitAudioFile(filePath, ext, MaxAudioFileSize); err != nil {
			return "", err
		}
		defer cleanup()
	} else {
		files = []string{filePath}
	}

	var builder strings.Builder
	for i, f := range files {
		debuglog.Log("Using model %s to transcribe part %d (file name: %s)...\n", model, i+1, f)
		var chunk *os.File
		if chunk, err = os.Open(f); err != nil {
			return "", err
		}
		params := openai.AudioTranscriptionNewParams{
			File:  chunk,
			Model: openai.AudioModel(model),
		}
		var resp *openai.Transcription
		resp, err = o.ApiClient.Audio.Transcriptions.New(ctx, params)
		chunk.Close()
		if err != nil {
			return "", err
		}
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(resp.Text)
	}

	return builder.String(), nil
}

// splitAudioFile splits the source file into chunks smaller than maxSize using ffmpeg.
// It returns the list of chunk file paths and a cleanup function.
func splitAudioFile(src, ext string, maxSize int64) (files []string, cleanup func(), err error) {
	if _, err = exec.LookPath("ffmpeg"); err != nil {
		return nil, nil, fmt.Errorf("ffmpeg not found: please install it")
	}

	var dir string
	if dir, err = os.MkdirTemp("", "fabric-audio-*"); err != nil {
		return nil, nil, err
	}
	cleanup = func() { os.RemoveAll(dir) }

	segmentTime := 600 // start with 10 minutes
	for {
		pattern := filepath.Join(dir, "chunk-%03d"+ext)
		debuglog.Log("Running ffmpeg to split audio into %d-second chunks...\n", segmentTime)
		cmd := exec.Command("ffmpeg", "-y", "-i", src, "-f", "segment", "-segment_time", fmt.Sprintf("%d", segmentTime), "-c", "copy", pattern)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err = cmd.Run(); err != nil {
			return nil, cleanup, fmt.Errorf("ffmpeg failed: %v: %s", err, stderr.String())
		}

		if files, err = filepath.Glob(filepath.Join(dir, "chunk-*"+ext)); err != nil {
			return nil, cleanup, err
		}
		sort.Strings(files)

		tooBig := false
		for _, f := range files {
			var info os.FileInfo
			if info, err = os.Stat(f); err != nil {
				return nil, cleanup, err
			}
			if info.Size() > maxSize {
				tooBig = true
				break
			}
		}
		if !tooBig {
			return files, cleanup, nil
		}
		for _, f := range files {
			_ = os.Remove(f)
		}
		if segmentTime <= 1 {
			return nil, cleanup, fmt.Errorf("unable to split file into acceptable size chunks")
		}
		segmentTime /= 2
	}
}
