package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleConfigurationCommandsNoop(t *testing.T) {
	handled, err := handleConfigurationCommands(&Flags{}, nil)
	require.False(t, handled)
	require.NoError(t, err)
}
