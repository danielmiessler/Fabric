package openai

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTranscribeFile_ValidationErrorsAreLocalized(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	client := &Client{}

	audioFile, err := os.CreateTemp("", "transcribe-valid-*.mp3")
	require.NoError(t, err)
	require.NoError(t, audioFile.Close())
	t.Cleanup(func() { _ = os.Remove(audioFile.Name()) })

	_, err = client.TranscribeFile(context.Background(), audioFile.Name(), "not-a-model", false)
	require.Error(t, err)
	assert.Equal(t,
		fmt.Sprintf(i18n.T("openai_audio_model_not_supported_for_transcription"), "not-a-model"),
		err.Error(),
	)

	unsupportedFile, err := os.CreateTemp("", "transcribe-invalid-*.txt")
	require.NoError(t, err)
	require.NoError(t, unsupportedFile.Close())
	t.Cleanup(func() { _ = os.Remove(unsupportedFile.Name()) })

	_, err = client.TranscribeFile(context.Background(), unsupportedFile.Name(), AllowedTranscriptionModels[0], false)
	require.Error(t, err)
	assert.Equal(t,
		fmt.Sprintf(i18n.T("openai_audio_unsupported_audio_format"), filepath.Ext(unsupportedFile.Name())),
		err.Error(),
	)
}

func TestTranscribeFile_FileSizeLimitErrorIsLocalized(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	client := &Client{}
	largeFile, err := os.CreateTemp("", "transcribe-large-*.mp3")
	require.NoError(t, err)
	t.Cleanup(func() { _ = os.Remove(largeFile.Name()) })

	require.NoError(t, largeFile.Truncate(MaxAudioFileSize+1))
	require.NoError(t, largeFile.Close())

	_, err = client.TranscribeFile(context.Background(), largeFile.Name(), AllowedTranscriptionModels[0], false)
	require.Error(t, err)
	assert.Equal(t,
		fmt.Sprintf(i18n.T("openai_audio_file_exceeds_limit_enable_split"), largeFile.Name()),
		err.Error(),
	)
}

func TestTranscribeFile_MissingFileReturnsStatError(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	client := &Client{}
	missingPath := filepath.Join(t.TempDir(), "missing.mp3")

	_, err = client.TranscribeFile(context.Background(), missingPath, AllowedTranscriptionModels[0], false)
	require.Error(t, err)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestSplitAudioFileRequiresFFmpeg(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	t.Setenv("PATH", "")

	files, cleanup, err := splitAudioFile("input.mp3", ".mp3", MaxAudioFileSize)
	require.Error(t, err)
	require.Nil(t, files)
	require.Nil(t, cleanup)
	assert.Equal(t, i18n.T("openai_audio_ffmpeg_not_found_install"), err.Error())
}

func TestTranscribeFile_UsesNilContextAndReturnsServerText(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	var authHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "/audio/transcriptions", r.URL.Path)
		require.NoError(t, r.ParseMultipartForm(1<<20))
		require.Equal(t, AllowedTranscriptionModels[0], r.FormValue("model"))
		w.Header().Set("Content-Type", "application/json")
		_, writeErr := w.Write([]byte(`{"text":"transcribed speech"}`))
		require.NoError(t, writeErr)
	}))
	defer srv.Close()

	client := newConfiguredAudioTestClient(t, srv.URL, "test-key")
	audioFile := filepath.Join(t.TempDir(), "sample.mp3")
	require.NoError(t, os.WriteFile(audioFile, []byte("audio"), 0o644))

	text, err := client.TranscribeFile(nil, audioFile, AllowedTranscriptionModels[0], false)
	require.NoError(t, err)
	require.Equal(t, "transcribed speech", text)
	require.Equal(t, "Bearer test-key", authHeader)
}

func TestTranscribeFile_PropagatesAPIError(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		_, writeErr := w.Write([]byte(`{"error":{"message":"upstream failed"}}`))
		require.NoError(t, writeErr)
	}))
	defer srv.Close()

	client := newConfiguredAudioTestClient(t, srv.URL, "test-key")
	audioFile := filepath.Join(t.TempDir(), "sample.mp3")
	require.NoError(t, os.WriteFile(audioFile, []byte("audio"), 0o644))

	_, err = client.TranscribeFile(context.Background(), audioFile, AllowedTranscriptionModels[0], false)
	require.Error(t, err)
	require.Contains(t, err.Error(), "upstream failed")
}

func TestSplitAudioFileRetriesUntilChunksFit(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	ffmpegDir := t.TempDir()
	ffmpegPath := filepath.Join(ffmpegDir, "ffmpeg")
	script := `#!/bin/sh
segment_time=""
pattern=""
prev=""
for arg in "$@"; do
  if [ "$prev" = "-segment_time" ]; then
    segment_time="$arg"
  fi
  prev="$arg"
  pattern="$arg"
done
prefix=${pattern%\%03d*}
suffix=${pattern#*%03d}
file0="${prefix}000${suffix}"
file1="${prefix}001${suffix}"
if [ "$segment_time" = "600" ]; then
  printf '123456789' > "$file0"
else
  printf '1234' > "$file0"
  printf '5678' > "$file1"
fi
`
	require.NoError(t, os.WriteFile(ffmpegPath, []byte(script), 0o755))
	t.Setenv("PATH", ffmpegDir)

	audioFile := filepath.Join(t.TempDir(), "input.mp3")
	require.NoError(t, os.WriteFile(audioFile, []byte("audio"), 0o644))

	files, cleanup, err := splitAudioFile(audioFile, ".mp3", 8)
	require.NoError(t, err)
	require.Len(t, files, 2)
	for _, file := range files {
		info, statErr := os.Stat(file)
		require.NoError(t, statErr)
		require.LessOrEqual(t, info.Size(), int64(8))
	}

	chunkDir := filepath.Dir(files[0])
	require.DirExists(t, chunkDir)
	cleanup()
	_, statErr := os.Stat(chunkDir)
	require.Error(t, statErr)
	require.True(t, os.IsNotExist(statErr))
}

func TestSplitAudioFileReturnsFFmpegStderrOnFailure(t *testing.T) {
	_, err := i18n.Init("en")
	require.NoError(t, err)

	ffmpegDir := t.TempDir()
	ffmpegPath := filepath.Join(ffmpegDir, "ffmpeg")
	script := "#!/bin/sh\nprintf 'ffmpeg exploded' >&2\nexit 2\n"
	require.NoError(t, os.WriteFile(ffmpegPath, []byte(script), 0o755))
	t.Setenv("PATH", ffmpegDir)

	audioFile := filepath.Join(t.TempDir(), "input.mp3")
	require.NoError(t, os.WriteFile(audioFile, []byte("audio"), 0o644))

	files, cleanup, err := splitAudioFile(audioFile, ".mp3", 8)
	require.Error(t, err)
	require.Nil(t, files)
	require.NotNil(t, cleanup)
	require.Contains(t, err.Error(), "ffmpeg exploded")
}

func newConfiguredAudioTestClient(t *testing.T, baseURL, apiKey string) *Client {
	t.Helper()

	client := NewClientCompatible("OpenAI", baseURL, nil)
	client.ApiKey.Value = apiKey
	client.ApiBaseURL.Value = baseURL
	require.NoError(t, client.configure())
	return client
}
