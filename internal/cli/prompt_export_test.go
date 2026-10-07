package cli

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	langtool "github.com/danielmiessler/fabric/internal/tools/lang"
)

func TestRenderPromptExport(t *testing.T) {
	tests := []struct {
		name    string
		flags   Flags
		pattern string
		context string
		history []*chat.ChatCompletionMessage
		tools   string
		want    string
	}{
		{
			name:    "pattern",
			flags:   Flags{Pattern: "p", Message: "user input"},
			pattern: "PATTERN\n{{input}}",
			want:    "User:\nPATTERN\nuser input\n\n",
		},
		{
			name:    "raw merges system into user",
			flags:   Flags{Context: "c", Message: "user input", Raw: true},
			context: "CONTEXT",
			want:    "User:\nCONTEXT\nuser input\n\n",
		},
		{
			name:    "context and user",
			flags:   Flags{Context: "c", Message: "user input"},
			context: "CONTEXT",
			want:    "System:\nCONTEXT\n\nUser:\nuser input\n\n",
		},
		{
			name:  "named session history",
			flags: Flags{Session: "existing", Message: "next"},
			history: []*chat.ChatCompletionMessage{
				{Role: chat.ChatMessageRoleUser, Content: "earlier"},
				{Role: chat.ChatMessageRoleAssistant, Content: "reply"},
			},
			want: "User:\nearlier\n\nAssistant:\nreply\n\nUser:\nnext\n\n",
		},
		{
			name:  "missing session",
			flags: Flags{Session: "missing", Message: "hello"},
			want:  "User:\nhello\n\n",
		},
		{
			name:    "tool input",
			flags:   Flags{Pattern: "p"},
			pattern: "PATTERN\n{{input}}",
			tools:   "tool data",
			want:    "User:\nPATTERN\ntool data\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := fsdb.NewDb(t.TempDir())
			must(t, os.WriteFile(db.EnvFilePath, nil, 0o644))
			must(t, db.Configure())
			if tt.pattern != "" {
				dir := filepath.Join(db.Patterns.Dir, "p")
				must(t, os.MkdirAll(dir, 0o755))
				must(t, os.WriteFile(filepath.Join(dir, db.Patterns.SystemPatternFile), []byte(tt.pattern), 0o644))
			}
			if tt.context != "" {
				must(t, os.WriteFile(filepath.Join(db.Contexts.Dir, "c"), []byte(tt.context), 0o644))
			}
			if tt.history != nil {
				must(t, db.Sessions.SaveSession(&fsdb.Session{Name: tt.flags.Session, Messages: tt.history}))
			}
			registry := &core.PluginRegistry{Db: db, Language: langtool.NewLanguage()}

			// A missing session must not print the new-session notice.
			var got string
			printed := captureStdout(t, func() {
				var err error
				got, err = renderPromptExport(&tt.flags, registry, "meta", tt.tools)
				must(t, err)
			})

			if got != tt.want {
				t.Fatalf("renderPromptExport() = %q, want %q", got, tt.want)
			}
			if len(printed) != 0 {
				t.Fatalf("expected no stdout output, got %q", printed)
			}
			// The export must not write a session file.
			if tt.flags.Session != "" && tt.history == nil {
				if _, err := os.Stat(filepath.Join(db.Sessions.Dir, tt.flags.Session+".json")); !os.IsNotExist(err) {
					t.Fatalf("expected no session file, got err %v", err)
				}
			}
		})
	}
}

func TestHandlePromptExport(t *testing.T) {
	db := fsdb.NewDb(t.TempDir())
	must(t, os.WriteFile(db.EnvFilePath, nil, 0o644))
	must(t, db.Configure())
	registry := &core.PluginRegistry{Db: db, Language: langtool.NewLanguage()}
	const want = "User:\nhello\n\n"

	t.Run("not handled", func(t *testing.T) {
		handled, err := handlePromptExport(&Flags{Message: "hello"}, registry, "")
		if handled || err != nil {
			t.Fatalf("got handled=%v err=%v, want false, nil", handled, err)
		}
	})

	t.Run("stdout", func(t *testing.T) {
		printed := captureStdout(t, func() {
			handled, err := handlePromptExport(&Flags{PrintPrompt: true, Message: "hello"}, registry, "")
			must(t, err)
			if !handled {
				t.Fatal("expected handled")
			}
		})
		if string(printed) != want {
			t.Fatalf("stdout = %q, want %q", printed, want)
		}
	})

	t.Run("output file", func(t *testing.T) {
		out := filepath.Join(t.TempDir(), "prompt.txt")
		printed := captureStdout(t, func() {
			_, err := handlePromptExport(&Flags{PrintPrompt: true, Message: "hello", Output: out}, registry, "")
			must(t, err)
		})
		if len(printed) != 0 {
			t.Fatalf("expected no stdout output, got %q", printed)
		}
		data, err := os.ReadFile(out)
		must(t, err)
		if string(data) != want {
			t.Fatalf("file = %q, want %q", data, want)
		}
	})

	t.Run("output file in missing directory", func(t *testing.T) {
		out := filepath.Join(t.TempDir(), "missing", "prompt.txt")
		handled, err := handlePromptExport(&Flags{PrintPrompt: true, Message: "hello", Output: out}, registry, "")
		if !handled || err == nil {
			t.Fatalf("got handled=%v err=%v, want true and an error", handled, err)
		}
	})
}

func TestValidatePromptExportFlags(t *testing.T) {
	for name, flags := range map[string]*Flags{
		"dry run":        {PrintPrompt: true, DryRun: true},
		"output session": {PrintPrompt: true, OutputSession: true},
		"workflow":       {PrintPrompt: true, Workflow: "wf"},
		"audio output":   {PrintPrompt: true, Output: "x.wav"},
	} {
		if validatePromptExportFlags(flags) == nil {
			t.Errorf("%s: expected validation error", name)
		}
	}
}

// captureStdout runs fn and returns what it writes to stdout.
func captureStdout(t *testing.T, fn func()) []byte {
	t.Helper()
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	must(t, err)
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()
	fn()
	must(t, w.Close())
	printed, err := io.ReadAll(r)
	must(t, err)
	return printed
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
