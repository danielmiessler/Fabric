package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

func CopyToClipboard(message string) (err error) {
	if err = clipboard.WriteAll(message); err != nil {
		err = fmt.Errorf(i18n.T("could_not_copy_to_clipboard"), err)
	}
	return
}

func CreateOutputFile(message string, fileName string) (err error) {
	if _, err = os.Stat(fileName); err == nil {
		err = fmt.Errorf(i18n.T("file_already_exists_not_overwriting"), fileName)
		return
	}
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		err = fmt.Errorf(i18n.T("error_creating_file"), err)
		return
	}
	defer file.Close()
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	if _, err = file.WriteString(message); err != nil {
		err = fmt.Errorf(i18n.T("error_writing_to_file"), err)
	} else {
		debuglog.Log("\n\n[Output also written to %s]\n", fileName)
	}
	return
}

// CreateAudioOutputFile creates a binary file for audio data
func CreateAudioOutputFile(audioData []byte, fileName string) (err error) {
	if filepath.Ext(fileName) == "" {
		fileName += ".wav"
	}

	// Unlike CreateOutputFile, this function overwrites and prints nothing.
	// handleChatProcessing checks for an existing file before the TTS call and prints the success message.
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		err = fmt.Errorf(i18n.T("error_creating_audio_file"), err)
		return
	}
	defer file.Close()

	if _, err = file.Write(audioData); err != nil {
		err = fmt.Errorf(i18n.T("error_writing_audio_data"), err)
	}
	return
}

// IsAudioFormat checks if the filename suggests an audio format
func IsAudioFormat(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	audioExts := []string{".wav", ".mp3", ".m4a", ".aac", ".ogg", ".flac"}
	return slices.Contains(audioExts, ext)
}
