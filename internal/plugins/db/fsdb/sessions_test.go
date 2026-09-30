package fsdb

import (
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
