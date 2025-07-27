package cli

import (
	"os"
	"testing"

	"github.com/danielmiessler/fabric/internal/core"

	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	t.Skip("Skipping test for now, collision with flag -t")
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{os.Args[0]}
	err := Cli("test")
	assert.Error(t, err)
	assert.Equal(t, core.NoSessionPatternUserMessages, err.Error())
}
