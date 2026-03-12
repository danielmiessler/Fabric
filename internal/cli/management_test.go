package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danielmiessler/fabric/internal/chat"
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
	"github.com/stretchr/testify/require"
)

func TestHandleManagementCommandsNoop(t *testing.T) {
	handled, err := handleManagementCommands(&Flags{}, nil)
	require.False(t, handled)
	require.NoError(t, err)
}

func TestHandleManagementCommandsPrintAndDelete(t *testing.T) {
	db := newConfiguredTestDB(t)

	require.NoError(t, db.Contexts.Save("project", []byte("context body")))
	require.NoError(t, db.Sessions.SaveSession(&fsdb.Session{
		Name: "session-a",
		Messages: []*chat.ChatCompletionMessage{
			{Role: chat.ChatMessageRoleUser, Content: "hello session"},
		},
	}))

	contextStdout, err := captureStdout(func() error {
		handled, runErr := handleManagementCommands(&Flags{PrintContext: "project"}, db)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, contextStdout, "context body")

	sessionStdout, err := captureStdout(func() error {
		handled, runErr := handleManagementCommands(&Flags{PrintSession: "session-a"}, db)
		require.True(t, handled)
		return runErr
	})
	require.NoError(t, err)
	require.Contains(t, sessionStdout, "hello session")

	handled, err := handleManagementCommands(&Flags{WipeContext: "project"}, db)
	require.True(t, handled)
	require.NoError(t, err)
	_, statErr := os.Stat(db.Contexts.BuildFilePathByName("project"))
	require.Error(t, statErr)
	require.True(t, os.IsNotExist(statErr))

	handled, err = handleManagementCommands(&Flags{WipeSession: "session-a"}, db)
	require.True(t, handled)
	require.NoError(t, err)
	_, statErr = os.Stat(db.Sessions.BuildFilePathByName("session-a"))
	require.Error(t, statErr)
	require.True(t, os.IsNotExist(statErr))
}

func newConfiguredTestDB(t *testing.T) *fsdb.Db {
	t.Helper()

	dbDir := filepath.Join(t.TempDir(), "fabric-db")
	require.NoError(t, os.MkdirAll(dbDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(dbDir, ".env"), []byte{}, 0o644))

	db := fsdb.NewDb(dbDir)
	require.NoError(t, db.Configure())
	return db
}
