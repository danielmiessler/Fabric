package openai

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// modelsCacheTTL is the age limit for a fresh model list. Model catalogs change
// rarely, so a long TTL avoids discovery endpoints that rate-limit with HTTP 429.
const modelsCacheTTL = 24 * time.Hour

// modelsCacheDir is a variable so that tests can point the cache at a temporary directory.
var modelsCacheDir = defaultModelsCacheDir

func defaultModelsCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "fabric", "cache", "models"), nil
}

type modelsCacheEntry struct {
	URL       string    `json:"url"`
	FetchedAt time.Time `json:"fetched_at"`
	Models    []string  `json:"models"`
}

var cacheNameSanitizer = regexp.MustCompile(`[^A-Za-z0-9_-]+`)

// modelsCacheFile builds the cache path from the provider name and a hash of the
// URL. A provider that changes its endpoint then reads a new file, not a stale one.
func modelsCacheFile(dir, providerName, fullURL string) string {
	sum := sha256.Sum256([]byte(fullURL))
	slug := cacheNameSanitizer.ReplaceAllString(providerName, "_")
	name := fmt.Sprintf("%s-%s.json", slug, hex.EncodeToString(sum[:])[:12])
	return filepath.Join(dir, name)
}

// readModelsCache returns the cached list for the provider and URL. A maxAge of
// 0 or less accepts an entry of any age. An empty cached list is a miss.
func readModelsCache(providerName, fullURL string, maxAge time.Duration) ([]string, bool) {
	dir, err := modelsCacheDir()
	if err != nil {
		return nil, false
	}
	data, err := os.ReadFile(modelsCacheFile(dir, providerName, fullURL))
	if err != nil {
		return nil, false
	}
	var entry modelsCacheEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, false
	}
	if entry.URL != fullURL || len(entry.Models) == 0 {
		return nil, false
	}
	if maxAge > 0 && time.Since(entry.FetchedAt) > maxAge {
		return nil, false
	}
	return entry.Models, true
}

// writeModelsCache skips an empty list, so a transient empty response does not
// stay cached for the full TTL.
func writeModelsCache(providerName, fullURL string, models []string) error {
	if len(models) == 0 {
		return nil
	}
	dir, err := modelsCacheDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	entry := modelsCacheEntry{URL: fullURL, FetchedAt: time.Now(), Models: models}
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return os.WriteFile(modelsCacheFile(dir, providerName, fullURL), data, 0o600)
}
