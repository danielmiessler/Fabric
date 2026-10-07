package fsdb

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
)

func TestSessions_GetOrCreateSession(t *testing.T) {
	dir := t.TempDir()
	sessions := &SessionsEntity{
		StorageEntity: &StorageEntity{Dir: dir, FileExtension: ".json"},
	}
	sessionName := "testSession"
	session, err := sessions.Get(sessionName)
	if err != nil {
		t.Fatalf("failed to get or create session: %v", err)
	}
	if session.Name != sessionName {
		t.Errorf("expected session name %v, got %v", sessionName, session.Name)
	}
}

// Get must reject an invalid name and must not answer with a new empty
// session. GET /sessions/<bad-name> is then a 400, the same as for the
// other entities.
func TestSessions_GetRejectsInvalidNames(t *testing.T) {
	sessions := &SessionsEntity{
		StorageEntity: &StorageEntity{Dir: t.TempDir(), FileExtension: ".json"},
	}
	for _, name := range invalidStorageNames {
		if _, err := sessions.Get(name); err == nil {
			t.Errorf("Get(%q) succeeded, want error", name)
		}
	}
}

func TestSessions_SaveSession(t *testing.T) {
	dir := t.TempDir()
	sessions := &SessionsEntity{
		StorageEntity: &StorageEntity{Dir: dir, FileExtension: ".json"},
	}
	sessionName := "testSession"
	session := &Session{Name: sessionName, Messages: []*chat.ChatCompletionMessage{{Content: "message1"}}}
	err := sessions.SaveSession(session)
	if err != nil {
		t.Fatalf("failed to save session: %v", err)
	}
	if !sessions.Exists(sessionName) {
		t.Errorf("expected session to be saved")
	}
}

// PrintSession must not write terminal control sequences from a stored reply.
func TestSessions_PrintSessionRemovesControlSequences(t *testing.T) {
	sessions := &SessionsEntity{
		StorageEntity: &StorageEntity{Dir: t.TempDir(), FileExtension: ".json"},
	}
	session := &Session{Name: "s", Messages: []*chat.ChatCompletionMessage{
		{Role: "assistant", Content: "a\x1b]52;c;eA==\x07b"},
	}}
	if err := sessions.SaveSession(session); err != nil {
		t.Fatal(err)
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	oldStdout := os.Stdout
	t.Cleanup(func() { os.Stdout = oldStdout })
	os.Stdout = w
	printErr := sessions.PrintSession("s")
	os.Stdout = oldStdout
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if printErr != nil {
		t.Fatal(printErr)
	}
	if strings.ContainsRune(string(out), '\x1b') || !strings.Contains(string(out), "ab") {
		t.Errorf("PrintSession wrote %q", out)
	}
}

// A null entry in a session file loads as a nil message. The vendor
// messages, Append and String must skip it.
func TestSession_NilMessageIsSkipped(t *testing.T) {
	s := &Session{Messages: []*chat.ChatCompletionMessage{
		nil,
		{Role: "user", Content: "hi"},
		nil,
	}}
	if got := s.GetVendorMessages(); len(got) != 1 {
		t.Fatalf("expected 1 vendor message, got %d", len(got))
	}
	s.Append(nil)
	if got := s.String(); !strings.Contains(got, "hi") {
		t.Fatalf("expected String to contain the message, got %q", got)
	}
}

// String must not stop on an image part with no image_url.
func TestSession_StringSkipsImagePartWithNoURL(t *testing.T) {
	s := &Session{Messages: []*chat.ChatCompletionMessage{{
		Role:         chat.ChatMessageRoleUser,
		MultiContent: []chat.ChatMessagePart{{Type: chat.ChatMessagePartTypeImageURL}},
	}}}
	_ = s.String()
}
