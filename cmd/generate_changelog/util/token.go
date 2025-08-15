package util

import (
	"os"
)

// GetTokenFromEnv returns a GitHub token based on the following precedence order:
//  1. If tokenValue is non-empty, it is returned.
//  2. Otherwise, if the GITHUB_TOKEN environment variable is set, its value is returned.
//  3. Otherwise, if the GH_TOKEN environment variable is set, its value is returned.
//  4. If none of the above are set, an empty string is returned.
//
// Example:
//
//	os.Setenv("GITHUB_TOKEN", "abc")
//	os.Setenv("GH_TOKEN", "def")
//	GetTokenFromEnv("xyz") // returns "xyz"
//	GetTokenFromEnv("")    // returns "abc"
//	os.Unsetenv("GITHUB_TOKEN")
//	GetTokenFromEnv("")    // returns "def"
//	os.Unsetenv("GH_TOKEN")
//	GetTokenFromEnv("")    // returns ""
func GetTokenFromEnv(tokenValue string) string {
	if tokenValue == "" {
		tokenValue = os.Getenv("GITHUB_TOKEN")
		if tokenValue == "" {
			tokenValue = os.Getenv("GH_TOKEN")
		}
	}
	return tokenValue
}
