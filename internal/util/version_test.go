package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeVersion(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "plain version", input: "v1.4.478", expected: "v1.4.478"},
		{name: "strips dirty marker", input: "v1.4.478+dirty", expected: "v1.4.478"},
		{name: "strips build metadata", input: "v1.4.478+incompatible", expected: "v1.4.478"},
		{name: "devel is empty", input: "(devel)", expected: ""},
		{name: "blank is empty", input: "   ", expected: ""},
		{name: "trims whitespace", input: " v1.4.478 ", expected: "v1.4.478"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, normalizeVersion(tc.input))
		})
	}
}

func TestFabricVersion_NonEmpty(t *testing.T) {
	assert.NotEmpty(t, FabricVersion())
}
