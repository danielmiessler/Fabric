package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// OAuthToken represents stored OAuth token information
type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

// IsExpired checks if the token is expired or will expire within the buffer time
func (t *OAuthToken) IsExpired(bufferMinutes int) bool {
	if t.ExpiresAt == 0 {
		return true
	}
	bufferTime := time.Duration(bufferMinutes) * time.Minute
	return time.Now().Add(bufferTime).Unix() >= t.ExpiresAt
}

// OAuthStorage handles persistent storage of OAuth tokens
type OAuthStorage struct {
	configDir string
}

// NewOAuthStorage creates a new OAuth storage instance
func NewOAuthStorage() (*OAuthStorage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "fabric")

	// Ensure config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	return &OAuthStorage{configDir: configDir}, nil
}

// GetTokenPath returns the file path for a provider's OAuth token
func (s *OAuthStorage) GetTokenPath(provider string) string {
	return filepath.Join(s.configDir, fmt.Sprintf(".%s_oauth", provider))
}

// SaveToken saves an OAuth token to disk with proper permissions
func (s *OAuthStorage) SaveToken(provider string, token *OAuthToken) error {
	tokenPath := s.GetTokenPath(provider)

	// Marshal token to JSON
	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	// Write to temporary file first for atomic operation
	tempPath := tokenPath + ".tmp"
	if err := os.WriteFile(tempPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write token file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tempPath, tokenPath); err != nil {
		os.Remove(tempPath) // Clean up temp file
		return fmt.Errorf("failed to save token file: %w", err)
	}

	return nil
}

// LoadToken loads an OAuth token from disk
func (s *OAuthStorage) LoadToken(provider string) (*OAuthToken, error) {
	tokenPath := s.GetTokenPath(provider)

	// Check if file exists
	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		return nil, nil // No token stored
	}

	// Read token file
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read token file: %w", err)
	}

	// Unmarshal token
	var token OAuthToken
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("failed to parse token file: %w", err)
	}

	return &token, nil
}

// DeleteToken removes a stored OAuth token
func (s *OAuthStorage) DeleteToken(provider string) error {
	tokenPath := s.GetTokenPath(provider)

	if err := os.Remove(tokenPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete token file: %w", err)
	}

	return nil
}

// HasValidToken checks if a valid (non-expired) token exists for a provider
func (s *OAuthStorage) HasValidToken(provider string, bufferMinutes int) bool {
	token, err := s.LoadToken(provider)
	if err != nil || token == nil {
		return false
	}

	return !token.IsExpired(bufferMinutes)
}
