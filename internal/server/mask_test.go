package restapi

import (
	"strings"
	"testing"
)

// TestMaskAPIKey_HidesLengthAndCharacters makes sure that the mask does not
// show the key length or characters from the key, but shows that a key is set.
func TestMaskAPIKey_HidesLengthAndCharacters(t *testing.T) {
	short := maskAPIKey("sk-ABCDEFGHIJKLMNOP")
	long := maskAPIKey("sk-ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	if short == "" || long == "" {
		t.Fatalf("a set key must produce a non-empty mask")
	}
	if short != long {
		t.Errorf("mask shows key length: %q vs %q", short, long)
	}
	if strings.ContainsAny(short, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") {
		t.Errorf("mask shows key characters: %q", short)
	}
	if !isRedacted(short) {
		t.Errorf("mask must be detectable as redacted on resubmit: %q", short)
	}
	if isRedacted("sk-ab*cd") || isRedacted("http://host/*") {
		t.Errorf("a value that contains '*' but is not the mask must not be redacted")
	}
	if maskAPIKey("") != "" {
		t.Errorf("an unset key must mask to empty string")
	}
}
