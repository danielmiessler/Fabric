package cli

import (
	"github.com/danielmiessler/fabric/internal/plugins/db/fsdb"
)

// handleManagementCommands runs the context and session wipe and print commands.
// It returns handled = true when a command ran and the caller must exit.
func handleManagementCommands(currentFlags *Flags, fabricDb *fsdb.Db) (handled bool, err error) {
	if currentFlags.WipeContext != "" {
		err = fabricDb.Contexts.Delete(currentFlags.WipeContext)
		return true, err
	}

	if currentFlags.WipeSession != "" {
		err = fabricDb.Sessions.Delete(currentFlags.WipeSession)
		return true, err
	}

	if currentFlags.PrintSession != "" {
		err = fabricDb.Sessions.PrintSession(currentFlags.PrintSession)
		return true, err
	}

	if currentFlags.PrintContext != "" {
		err = fabricDb.Contexts.PrintContext(currentFlags.PrintContext)
		return true, err
	}

	return false, nil
}
