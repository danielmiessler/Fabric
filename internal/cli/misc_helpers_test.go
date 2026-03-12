package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/atotto/clipboard"
)

func TestCreateOutputFileDoesNotOverwriteExistingFile(t *testing.T) {
	fileName := filepath.Join(t.TempDir(), "existing.txt")
	original := "already here\n"
	if err := os.WriteFile(fileName, []byte(original), 0o644); err != nil {
		t.Fatalf("seed existing file: %v", err)
	}

	err := CreateOutputFile("new content", fileName)
	if err == nil {
		t.Fatal("CreateOutputFile() expected error for existing file, got nil")
	}

	got, readErr := os.ReadFile(fileName)
	if readErr != nil {
		t.Fatalf("read preserved file: %v", readErr)
	}
	if string(got) != original {
		t.Fatalf("existing file content changed: want %q, got %q", original, string(got))
	}
}

func TestCopyToClipboardUsesAvailableBinary(t *testing.T) {
	if clipboard.Unsupported {
		t.Skip("clipboard utility not available on this platform at init time")
	}

	tempDir := t.TempDir()
	capturePath := filepath.Join(tempDir, "clipboard.txt")
	script := "#!/bin/sh\ncat > \"$CLIPBOARD_CAPTURE_PATH\"\n"
	commandNames := []string{"pbcopy", "xclip", "xsel", "wl-copy", "clip.exe", "termux-clipboard-set"}
	for _, name := range commandNames {
		if err := os.WriteFile(filepath.Join(tempDir, name), []byte(script), 0o755); err != nil {
			t.Fatalf("write fake clipboard command %s: %v", name, err)
		}
	}

	t.Setenv("CLIPBOARD_CAPTURE_PATH", capturePath)
	t.Setenv("PATH", tempDir+string(os.PathListSeparator)+os.Getenv("PATH"))

	if err := CopyToClipboard("copied text"); err != nil {
		t.Fatalf("CopyToClipboard() error = %v", err)
	}

	got, err := os.ReadFile(capturePath)
	if err != nil {
		t.Fatalf("read clipboard capture: %v", err)
	}
	if string(got) != "copied text" {
		t.Fatalf("clipboard content = %q, want %q", string(got), "copied text")
	}
}

func TestCopyToClipboardReturnsErrorWhenCommandUnavailable(t *testing.T) {
	t.Setenv("PATH", t.TempDir())

	err := CopyToClipboard("copied text")
	if err == nil {
		t.Fatal("CopyToClipboard() expected error when clipboard command is unavailable, got nil")
	}
}

func TestCreateAudioOutputFileDefaultsToWAVExtension(t *testing.T) {
	baseName := filepath.Join(t.TempDir(), "speech")
	audio := []byte{0x00, 0x01, 0x02, 0x03}

	if err := CreateAudioOutputFile(audio, baseName); err != nil {
		t.Fatalf("CreateAudioOutputFile() error = %v", err)
	}

	got, err := os.ReadFile(baseName + ".wav")
	if err != nil {
		t.Fatalf("read generated audio file: %v", err)
	}
	if string(got) != string(audio) {
		t.Fatalf("audio bytes mismatch: want %v, got %v", audio, got)
	}
}

func TestCreateOutputFileAppendsTrailingNewline(t *testing.T) {
	fileName := filepath.Join(t.TempDir(), "output.txt")

	if err := CreateOutputFile("line without newline", fileName); err != nil {
		t.Fatalf("CreateOutputFile() error = %v", err)
	}

	got, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("read output file: %v", err)
	}
	if string(got) != "line without newline\n" {
		t.Fatalf("CreateOutputFile() content = %q, want trailing newline", string(got))
	}
}

func TestCreateOutputFileReturnsCreateErrorForMissingParent(t *testing.T) {
	fileName := filepath.Join(t.TempDir(), "missing", "output.txt")

	err := CreateOutputFile("hello", fileName)
	if err == nil {
		t.Fatal("CreateOutputFile() expected error for missing parent directory, got nil")
	}
}

func TestCreateAudioOutputFilePreservesProvidedExtension(t *testing.T) {
	fileName := filepath.Join(t.TempDir(), "speech.mp3")
	audio := []byte{0x10, 0x11, 0x12}

	if err := CreateAudioOutputFile(audio, fileName); err != nil {
		t.Fatalf("CreateAudioOutputFile() error = %v", err)
	}

	if _, err := os.Stat(fileName + ".wav"); !os.IsNotExist(err) {
		t.Fatalf("unexpected default wav file created: %v", err)
	}

	got, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("read audio file: %v", err)
	}
	if string(got) != string(audio) {
		t.Fatalf("audio bytes mismatch: want %v, got %v", audio, got)
	}
}

func TestCreateAudioOutputFileReturnsErrorForMissingParent(t *testing.T) {
	fileName := filepath.Join(t.TempDir(), "missing", "speech.wav")

	err := CreateAudioOutputFile([]byte{0x01}, fileName)
	if err == nil {
		t.Fatal("CreateAudioOutputFile() expected error for missing parent directory, got nil")
	}
}

func TestIsAudioFormat(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     bool
	}{
		{name: "wav", fileName: "voice.wav", want: true},
		{name: "uppercase mp3", fileName: "VOICE.MP3", want: true},
		{name: "unsupported extension", fileName: "notes.txt", want: false},
		{name: "no extension", fileName: "voice", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAudioFormat(tt.fileName); got != tt.want {
				t.Fatalf("IsAudioFormat(%q) = %v, want %v", tt.fileName, got, tt.want)
			}
		})
	}
}

func TestParseDebugLevel(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{name: "missing flag", args: []string{"--copy"}, want: 0},
		{name: "space separated", args: []string{"--debug", "3"}, want: 3},
		{name: "equals form", args: []string{"--debug=4"}, want: 4},
		{name: "invalid value", args: []string{"--debug", "oops"}, want: 0},
		{name: "first valid wins", args: []string{"--debug=2", "--debug=4"}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDebugLevel(tt.args); got != tt.want {
				t.Fatalf("parseDebugLevel(%v) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}

func TestDetectLanguageFromArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "long flag separate value", args: []string{"fabric", "--language", "es"}, want: "es"},
		{name: "long flag equals", args: []string{"fabric", "--language=fr"}, want: "fr"},
		{name: "short flag separate value", args: []string{"fabric", "-g", "de"}, want: "de"},
		{name: "short flag equals", args: []string{"fabric", "-g=hi"}, want: "hi"},
		{name: "missing value", args: []string{"fabric", "--language"}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			t.Cleanup(func() { os.Args = oldArgs })
			os.Args = tt.args

			if got := detectLanguageFromArgs(); got != tt.want {
				t.Fatalf("detectLanguageFromArgs() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDetectLanguageFromEnv(t *testing.T) {
	tests := []struct {
		name       string
		lcAll      string
		lcMessages string
		lang       string
		want       string
	}{
		{
			name:       "LC_ALL takes precedence and normalizes locale",
			lcAll:      "es_ES.UTF-8",
			lcMessages: "de_DE.UTF-8",
			lang:       "fr_FR.UTF-8",
			want:       "es",
		},
		{
			name:       "LC_MESSAGES used when LC_ALL unset",
			lcMessages: "de_DE.UTF-8",
			lang:       "fr_FR.UTF-8",
			want:       "de",
		},
		{
			name: "LANG used when more specific vars unset",
			lang: "pt_BR.UTF-8",
			want: "pt",
		},
		{
			name: "C locale ignored",
			lang: "C",
			want: "",
		},
		{
			name: "POSIX locale ignored",
			lang: "POSIX",
			want: "",
		},
		{
			name: "plain language value returned as-is",
			lang: "ja",
			want: "ja",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("LC_ALL", tt.lcAll)
			t.Setenv("LC_MESSAGES", tt.lcMessages)
			t.Setenv("LANG", tt.lang)

			if got := detectLanguageFromEnv(); got != tt.want {
				t.Fatalf("detectLanguageFromEnv() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIsTTSModel(t *testing.T) {
	tests := []struct {
		name  string
		model string
		want  bool
	}{
		{name: "tts suffix", model: "gpt-4o-mini-tts", want: true},
		{name: "preview tts", model: "Preview-TTS", want: true},
		{name: "text to speech phrase", model: "acme-text-to-speech-v1", want: true},
		{name: "plain chat model", model: "gpt-4o-mini", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTTSModel(tt.model); got != tt.want {
				t.Fatalf("isTTSModel(%q) = %v, want %v", tt.model, got, tt.want)
			}
		})
	}
}
