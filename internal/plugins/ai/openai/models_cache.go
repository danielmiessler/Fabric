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

// modelsCacheTTL is how long a successfully fetched model list is considered
// fresh. Model catalogs change rarely, so a long TTL avoids re-hitting
// discovery endpoints that aggressively rate-limit (e.g. GitHub Models'
// catalog, which returns HTTP 429 with a Retry-After header).
const modelsCacheTTL = 24 * time.Hour

// modelsCacheDir returns the directory used to cache provider model lists. It
// is a package variable so tests can redirect it to a temporary location.
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

// modelsCacheFile derives a stable, collision-resistant cache path from the
// provider name and the full request URL. The URL is hashed so a provider that
// changes its endpoint does not read a stale entry from the old one.
func modelsCacheFile(dir, providerName, fullURL string) string {
	sum := sha256.Sum256([]byte(fullURL))
	slug := cacheNameSanitizer.ReplaceAllString(providerName, "_")
	name := fmt.Sprintf("%s-%s.json", slug, hex.EncodeToString(sum[:])[:12])
	return filepath.Join(dir, name)
}

// readModelsCache returns the cached model list for (providerName, fullURL).
// When maxAge > 0 the entry must be younger than maxAge; maxAge <= 0 accepts an
// entry of any age (used to fall back to a stale list when a fetch fails).
// Empty cached lists are never returned.
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

// writeModelsCache persists a successfully fetched model list. Empty lists are
// not cached so a transient empty response does not stick for the whole TTL.
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
