package template

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// ComputeHash computes SHA-256 hash of a file at given path.
// Returns the hex-encoded hash string or an error if the operation fails.
func ComputeHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// ComputeStringHash returns hex-encoded SHA-256 hash of the given string
func ComputeStringHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
