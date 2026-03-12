package cli

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCliVersion(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{originalArgs[0], "--version"}

	stdout, err := captureStdout(func() error {
		return Cli("test-version")
	})
	require.NoError(t, err)
	require.Equal(t, "test-version\n", stdout)
}
