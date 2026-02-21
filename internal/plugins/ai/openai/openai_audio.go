package openai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"

	openai "github.com/openai/openai-go"
)

// transcriptionResult holds the result of a single chunk transcription.
type transcriptionResult struct {
	index int
	text  string
	err   error
}

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
		return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_audio_model_not_supported_for_transcription"), model))
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	if _, ok := allowedAudioExtensions[ext]; !ok {
		return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_audio_unsupported_audio_format"), ext))
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	var files []string
	var cleanup func()
	if info.Size() > MaxAudioFileSize {
		if !split {
			return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_audio_file_exceeds_limit_enable_split"), filePath))
		}
		debuglog.Log("%s\n", fmt.Sprintf(i18n.T("openai_audio_file_exceeds_limit_splitting"), filePath))
		if files, cleanup, err = splitAudioFile(filePath, ext, MaxAudioFileSize); err != nil {
			return "", err
		}
		defer cleanup()
	} else {
		files = []string{filePath}
	}

	resultsChan := make(chan transcriptionResult, len(files))
	var wg sync.WaitGroup

	for i, f := range files {
		wg.Add(1)
		go func(index int, filePath string) {
			defer wg.Done()
			debuglog.Log("%s\n", fmt.Sprintf(i18n.T("openai_audio_using_model_to_transcribe_part"), model, index+1, filePath))

			chunk, openErr := os.Open(filePath)
			if openErr != nil {
				resultsChan <- transcriptionResult{index: index, err: openErr}
				return
			}
			defer chunk.Close()

			params := openai.AudioTranscriptionNewParams{
				File:  chunk,
				Model: openai.AudioModel(model),
			}
			resp, transcribeErr := o.ApiClient.Audio.Transcriptions.New(ctx, params)
			if transcribeErr != nil {
				resultsChan <- transcriptionResult{index: index, err: transcribeErr}
				return
			}
			resultsChan <- transcriptionResult{index: index, text: resp.Text}
		}(i, f)
	}

	wg.Wait()
	close(resultsChan)

	results := make([]transcriptionResult, 0, len(files))
	for result := range resultsChan {
		if result.err != nil {
			return "", result.err
		}
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].index < results[j].index
	})

	var builder strings.Builder
	for i, result := range results {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(result.text)
	}

	return builder.String(), nil
}

// splitAudioFile splits the source file into chunks smaller than maxSize using ffmpeg.
// It returns the list of chunk file paths and a cleanup function.
func splitAudioFile(src, ext string, maxSize int64) (files []string, cleanup func(), err error) {
	if _, err = exec.LookPath("ffmpeg"); err != nil {
		return nil, nil, errors.New(i18n.T("openai_audio_ffmpeg_not_found_install"))
	}

	var dir string
	if dir, err = os.MkdirTemp("", "fabric-audio-*"); err != nil {
		return nil, nil, err
	}
	cleanup = func() { os.RemoveAll(dir) }

	segmentTime := 600 // start with 10 minutes
	for {
		pattern := filepath.Join(dir, "chunk-%03d"+ext)
		debuglog.Log("%s\n", fmt.Sprintf(i18n.T("openai_audio_running_ffmpeg_split_chunks"), segmentTime))
		cmd := exec.Command("ffmpeg", "-y", "-i", src, "-f", "segment", "-segment_time", fmt.Sprintf("%d", segmentTime), "-c", "copy", pattern)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err = cmd.Run(); err != nil {
			return nil, cleanup, fmt.Errorf("%s", fmt.Sprintf(i18n.T("openai_audio_ffmpeg_failed"), err, stderr.String()))
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
			return nil, cleanup, errors.New(i18n.T("openai_audio_unable_to_split_acceptable_size_chunks"))
		}
		segmentTime /= 2
	}
}
