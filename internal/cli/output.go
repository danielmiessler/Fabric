package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

func CopyToClipboard(message string) (err error) {
	if err = clipboard.WriteAll(message); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("could_not_copy_to_clipboard"), err))
	}
	return
}

func CreateOutputFile(message string, fileName string) (err error) {
	if _, err = os.Stat(fileName); err == nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("file_already_exists_not_overwriting"), fileName))
		return
	}
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_creating_file"), err))
		return
	}
	defer file.Close()
	if _, err = file.WriteString(message); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_writing_to_file"), err))
	} else {
		debuglog.Log("\n\n[Output also written to %s]\n", fileName)
	}
	return
}

// CreateAudioOutputFile creates a binary file for audio data
func CreateAudioOutputFile(audioData []byte, fileName string) (err error) {
	// If no extension is provided, default to .wav
	if filepath.Ext(fileName) == "" {
		fileName += ".wav"
	}

	// File existence check is now done in the CLI layer before TTS generation
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_creating_audio_file"), err))
		return
	}
	defer file.Close()

	if _, err = file.Write(audioData); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_writing_audio_data"), err))
	}
	// No redundant output message here - the CLI layer handles success messaging
	return
}

// IsAudioFormat checks if the filename suggests an audio format
func IsAudioFormat(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	audioExts := []string{".wav", ".mp3", ".m4a", ".aac", ".ogg", ".flac"}
	for _, audioExt := range audioExts {
		if ext == audioExt {
			return true
		}
	}
	return false
}
