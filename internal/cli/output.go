package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	debuglog "github.com/danielmiessler/fabric/internal/log"
)

func CopyToClipboard(message string) (err error) {
	if err = clipboard.WriteAll(message); err != nil {
		err = fmt.Errorf("could not copy to clipboard: %v", err)
	}
	return
}

func CreateOutputFile(message string, fileName string) (err error) {
	if _, err = os.Stat(fileName); err == nil {
		err = fmt.Errorf("file %s already exists, not overwriting. Rename the existing file or choose a different name", fileName)
		return
	}
	var file *os.File
	if file, err = os.Create(fileName); err != nil {
		err = fmt.Errorf("error creating file: %v", err)
		return
	}
	defer file.Close()
	if _, err = file.WriteString(message); err != nil {
		err = fmt.Errorf("error writing to file: %v", err)
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
		err = fmt.Errorf("error creating audio file: %v", err)
		return
	}
	defer file.Close()

	if _, err = file.Write(audioData); err != nil {
		err = fmt.Errorf("error writing audio data to file: %v", err)
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
