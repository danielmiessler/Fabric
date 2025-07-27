package util

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestOAuthToken_IsExpired(t *testing.T) {
	tests := []struct {
		name          string
		expiresAt     int64
		bufferMinutes int
		expected      bool
	}{
		{
			name:          "token not expired",
			expiresAt:     time.Now().Unix() + 3600, // 1 hour from now
			bufferMinutes: 5,
			expected:      false,
		},
		{
			name:          "token expired",
			expiresAt:     time.Now().Unix() - 3600, // 1 hour ago
			bufferMinutes: 5,
			expected:      true,
		},
		{
			name:          "token expires within buffer",
			expiresAt:     time.Now().Unix() + 120, // 2 minutes from now
			bufferMinutes: 5,
			expected:      true,
		},
		{
			name:          "zero expiry time",
			expiresAt:     0,
			bufferMinutes: 5,
			expected:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := &OAuthToken{ExpiresAt: tt.expiresAt}
			if got := token.IsExpired(tt.bufferMinutes); got != tt.expected {
				t.Errorf("IsExpired() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestOAuthStorage_SaveAndLoadToken(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "fabric_oauth_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create storage with custom config dir
	storage := &OAuthStorage{configDir: tempDir}

	// Test token
	token := &OAuthToken{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresAt:    time.Now().Unix() + 3600,
		TokenType:    "Bearer",
		Scope:        "test_scope",
	}

	// Test saving token
	err = storage.SaveToken("test_provider", token)
	if err != nil {
		t.Fatalf("Failed to save token: %v", err)
	}

	// Verify file exists and has correct permissions
	tokenPath := storage.GetTokenPath("test_provider")
	info, err := os.Stat(tokenPath)
	if err != nil {
		t.Fatalf("Token file not created: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("Token file has wrong permissions: %v, want 0600", info.Mode().Perm())
	}

	// Test loading token
	loadedToken, err := storage.LoadToken("test_provider")
	if err != nil {
		t.Fatalf("Failed to load token: %v", err)
	}
	if loadedToken == nil {
		t.Fatal("Loaded token is nil")
	}

	// Verify token data
	if loadedToken.AccessToken != token.AccessToken {
		t.Errorf("AccessToken mismatch: got %v, want %v", loadedToken.AccessToken, token.AccessToken)
	}
	if loadedToken.RefreshToken != token.RefreshToken {
		t.Errorf("RefreshToken mismatch: got %v, want %v", loadedToken.RefreshToken, token.RefreshToken)
	}
	if loadedToken.ExpiresAt != token.ExpiresAt {
		t.Errorf("ExpiresAt mismatch: got %v, want %v", loadedToken.ExpiresAt, token.ExpiresAt)
	}
}

func TestOAuthStorage_LoadNonExistentToken(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "fabric_oauth_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	storage := &OAuthStorage{configDir: tempDir}

	// Try to load non-existent token
	token, err := storage.LoadToken("nonexistent")
	if err != nil {
		t.Fatalf("Unexpected error loading non-existent token: %v", err)
	}
	if token != nil {
		t.Error("Expected nil token for non-existent provider")
	}
}

func TestOAuthStorage_DeleteToken(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "fabric_oauth_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	storage := &OAuthStorage{configDir: tempDir}

	// Create and save a token
	token := &OAuthToken{
		AccessToken:  "test_token",
		RefreshToken: "test_refresh",
		ExpiresAt:    time.Now().Unix() + 3600,
	}
	err = storage.SaveToken("test_provider", token)
	if err != nil {
		t.Fatalf("Failed to save token: %v", err)
	}

	// Verify token exists
	tokenPath := storage.GetTokenPath("test_provider")
	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		t.Fatal("Token file should exist before deletion")
	}

	// Delete token
	err = storage.DeleteToken("test_provider")
	if err != nil {
		t.Fatalf("Failed to delete token: %v", err)
	}

	// Verify token is deleted
	if _, err := os.Stat(tokenPath); !os.IsNotExist(err) {
		t.Error("Token file should not exist after deletion")
	}

	// Test deleting non-existent token (should not error)
	err = storage.DeleteToken("nonexistent")
	if err != nil {
		t.Errorf("Deleting non-existent token should not error: %v", err)
	}
}

func TestOAuthStorage_HasValidToken(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "fabric_oauth_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	storage := &OAuthStorage{configDir: tempDir}

	// Test with no token
	if storage.HasValidToken("test_provider", 5) {
		t.Error("Should return false when no token exists")
	}

	// Save valid token
	validToken := &OAuthToken{
		AccessToken:  "valid_token",
		RefreshToken: "refresh_token",
		ExpiresAt:    time.Now().Unix() + 3600, // 1 hour from now
	}
	err = storage.SaveToken("test_provider", validToken)
	if err != nil {
		t.Fatalf("Failed to save valid token: %v", err)
	}

	// Test with valid token
	if !storage.HasValidToken("test_provider", 5) {
		t.Error("Should return true for valid token")
	}

	// Save expired token
	expiredToken := &OAuthToken{
		AccessToken:  "expired_token",
		RefreshToken: "refresh_token",
		ExpiresAt:    time.Now().Unix() - 3600, // 1 hour ago
	}
	err = storage.SaveToken("expired_provider", expiredToken)
	if err != nil {
		t.Fatalf("Failed to save expired token: %v", err)
	}

	// Test with expired token
	if storage.HasValidToken("expired_provider", 5) {
		t.Error("Should return false for expired token")
	}
}

func TestOAuthStorage_GetTokenPath(t *testing.T) {
	storage := &OAuthStorage{configDir: "/test/config"}

	expected := filepath.Join("/test/config", ".test_provider_oauth")
	actual := storage.GetTokenPath("test_provider")

	if actual != expected {
		t.Errorf("GetTokenPath() = %v, want %v", actual, expected)
	}
}
