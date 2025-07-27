package anthropic

// OAuth Testing Strategy:
//
// This test suite covers OAuth functionality while avoiding real external calls.
// Key principles:
// 1. Never trigger real OAuth flows that would open browsers or call external APIs
// 2. Use temporary directories and mock tokens for isolated testing
// 3. Skip integration tests that would require real OAuth servers
// 4. Test error paths and edge cases safely
//
// Tests are categorized as:
// - Unit tests: Test individual functions with mocked data (SAFE)
// - Integration tests: Would require real OAuth servers (SKIPPED)
// - Error path tests: Test failure scenarios safely (SAFE)

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/danielmiessler/fabric/internal/util"
)

// createTestToken creates a test OAuth token
func createTestToken(accessToken, refreshToken string, expiresIn int64) *util.OAuthToken {
	return &util.OAuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Unix() + expiresIn,
		TokenType:    "Bearer",
		Scope:        "org:create_api_key user:profile user:inference",
	}
}

// createExpiredToken creates an expired test token
func createExpiredToken(accessToken, refreshToken string) *util.OAuthToken {
	return &util.OAuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Unix() - 3600, // Expired 1 hour ago
		TokenType:    "Bearer",
		Scope:        "org:create_api_key user:profile user:inference",
	}
}

// mockTokenServer creates a mock OAuth token server for testing
func mockTokenServer(_ *testing.T, responses map[string]interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/oauth/token" {
			http.NotFound(w, r)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}

		var req map[string]string
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		grantType := req["grant_type"]
		response, exists := responses[grantType]
		if !exists {
			http.Error(w, "Unsupported grant type", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if errorResp, ok := response.(map[string]interface{}); ok && errorResp["error"] != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(response)
	}))
}

func TestGeneratePKCE(t *testing.T) {
	verifier, challenge, err := generatePKCE()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if verifier == "" {
		t.Error("Expected non-empty verifier")
	}

	if challenge == "" {
		t.Error("Expected non-empty challenge")
	}

	if len(verifier) < 43 { // Base64 encoded 32 bytes should be at least 43 chars
		t.Errorf("Verifier too short: %d chars", len(verifier))
	}

	if len(challenge) < 43 { // SHA256 hash should be at least 43 chars when base64 encoded
		t.Errorf("Challenge too short: %d chars", len(challenge))
	}
}

func TestExchangeToken_Success(t *testing.T) {
	// Create mock server
	server := mockTokenServer(t, map[string]interface{}{
		"authorization_code": map[string]interface{}{
			"access_token":  "test_access_token",
			"refresh_token": "test_refresh_token",
			"expires_in":    3600,
			"token_type":    "Bearer",
			"scope":         "org:create_api_key user:profile user:inference",
		},
	})
	defer server.Close()

	// Create a temporary directory for token storage
	tempDir := t.TempDir()

	// Mock the storage creation to use our temp directory
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// Set up a fake home directory
	fakeHome := filepath.Join(tempDir, "home")
	os.MkdirAll(filepath.Join(fakeHome, ".config", "fabric"), 0755)
	os.Setenv("HOME", fakeHome)

	// This test would need the actual exchangeToken function to be modified to accept a custom URL
	// For now, we'll test the logic without the actual HTTP call
	t.Skip("Skipping integration test - would need URL injection for proper testing")
}
func TestRefreshToken_Success(t *testing.T) {
	// Create temporary directory and set up fake home
	tempDir := t.TempDir()
	fakeHome := filepath.Join(tempDir, "home")
	configDir := filepath.Join(fakeHome, ".config", "fabric")
	os.MkdirAll(configDir, 0755)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", fakeHome)

	// Create an expired token
	expiredToken := createExpiredToken("old_access_token", "valid_refresh_token")

	// Save the expired token
	tokenPath := filepath.Join(configDir, ".test_oauth")
	data, _ := json.MarshalIndent(expiredToken, "", "  ")
	os.WriteFile(tokenPath, data, 0600)

	// Create mock server for refresh
	server := mockTokenServer(t, map[string]interface{}{
		"refresh_token": map[string]interface{}{
			"access_token":  "new_access_token",
			"refresh_token": "new_refresh_token",
			"expires_in":    3600,
			"token_type":    "Bearer",
			"scope":         "org:create_api_key user:profile user:inference",
		},
	})
	defer server.Close()

	// This test would need the RefreshToken function to accept a custom URL
	t.Skip("Skipping integration test - would need URL injection for proper testing")
}

func TestRefreshToken_NoRefreshToken(t *testing.T) {
	// Create temporary directory and set up fake home
	tempDir := t.TempDir()
	fakeHome := filepath.Join(tempDir, "home")
	configDir := filepath.Join(fakeHome, ".config", "fabric")
	os.MkdirAll(configDir, 0755)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", fakeHome)

	// Create a token without refresh token
	tokenWithoutRefresh := &util.OAuthToken{
		AccessToken:  "access_token",
		RefreshToken: "", // No refresh token
		ExpiresAt:    time.Now().Unix() - 3600,
		TokenType:    "Bearer",
		Scope:        "org:create_api_key user:profile user:inference",
	}

	// Save the token
	tokenPath := filepath.Join(configDir, ".test_oauth")
	data, _ := json.MarshalIndent(tokenWithoutRefresh, "", "  ")
	os.WriteFile(tokenPath, data, 0600)

	// Test RefreshToken
	_, err := RefreshToken("test")

	if err == nil {
		t.Error("Expected error when no refresh token available")
	}

	if !strings.Contains(err.Error(), "no refresh token available") {
		t.Errorf("Expected 'no refresh token available' error, got: %v", err)
	}
}

func TestRefreshToken_NoStoredToken(t *testing.T) {
	// Create temporary directory and set up fake home
	tempDir := t.TempDir()
	fakeHome := filepath.Join(tempDir, "home")
	configDir := filepath.Join(fakeHome, ".config", "fabric")
	os.MkdirAll(configDir, 0755)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", fakeHome)

	// Don't create any token file

	// Test RefreshToken
	_, err := RefreshToken("nonexistent")

	if err == nil {
		t.Error("Expected error when no token stored")
	}
}

func TestOAuthTransport_RoundTrip(t *testing.T) {
	// Create a mock client
	client := &Client{}

	// Create the transport
	transport := NewOAuthTransport(client, http.DefaultTransport)

	// Create a test request
	req := httptest.NewRequest("GET", "https://api.anthropic.com/v1/messages", nil)
	req.Header.Set("x-api-key", "should-be-removed")

	// Create temporary directory and set up fake home with valid token
	tempDir := t.TempDir()
	fakeHome := filepath.Join(tempDir, "home")
	configDir := filepath.Join(fakeHome, ".config", "fabric")
	os.MkdirAll(configDir, 0755)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", fakeHome)

	// Create a valid token
	validToken := createTestToken("valid_access_token", "refresh_token", 3600)
	tokenPath := filepath.Join(configDir, fmt.Sprintf(".%s_oauth", authTokenIdentifier))
	data, _ := json.MarshalIndent(validToken, "", "  ")
	os.WriteFile(tokenPath, data, 0600)

	// Create a mock server to handle the request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that OAuth headers are set correctly
		auth := r.Header.Get("Authorization")
		if auth != "Bearer valid_access_token" {
			t.Errorf("Expected 'Bearer valid_access_token', got '%s'", auth)
		}

		beta := r.Header.Get("anthropic-beta")
		if beta != "oauth-2025-04-20" {
			t.Errorf("Expected 'oauth-2025-04-20', got '%s'", beta)
		}

		userAgent := r.Header.Get("User-Agent")
		if userAgent != "ai-sdk/anthropic" {
			t.Errorf("Expected 'ai-sdk/anthropic', got '%s'", userAgent)
		}

		// Check that x-api-key header is removed
		if r.Header.Get("x-api-key") != "" {
			t.Error("Expected x-api-key header to be removed")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer server.Close()

	// Update the request URL to point to our mock server
	req.URL.Host = strings.TrimPrefix(server.URL, "http://")
	req.URL.Scheme = "http"

	// Execute the request
	resp, err := transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestRunOAuthFlow_ExistingValidToken(t *testing.T) {
	// Create temporary directory and set up fake home
	tempDir := t.TempDir()
	fakeHome := filepath.Join(tempDir, "home")
	configDir := filepath.Join(fakeHome, ".config", "fabric")
	os.MkdirAll(configDir, 0755)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", fakeHome)

	// Create a valid token
	validToken := createTestToken("existing_valid_token", "refresh_token", 3600)
	tokenPath := filepath.Join(configDir, ".test_oauth")
	data, _ := json.MarshalIndent(validToken, "", "  ")
	os.WriteFile(tokenPath, data, 0600)

	// Test RunOAuthFlow - should return existing token without starting OAuth flow
	token, err := RunOAuthFlow("test")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token != "existing_valid_token" {
		t.Errorf("Expected 'existing_valid_token', got '%s'", token)
	}
}

// Test helper functions
func TestCreateTestToken(t *testing.T) {
	token := createTestToken("access", "refresh", 3600)

	if token.AccessToken != "access" {
		t.Errorf("Expected access token 'access', got '%s'", token.AccessToken)
	}

	if token.RefreshToken != "refresh" {
		t.Errorf("Expected refresh token 'refresh', got '%s'", token.RefreshToken)
	}

	if token.IsExpired(5) {
		t.Error("Expected token to not be expired")
	}
}

func TestCreateExpiredToken(t *testing.T) {
	token := createExpiredToken("access", "refresh")

	if !token.IsExpired(5) {
		t.Error("Expected token to be expired")
	}
}

// TestTokenExpirationLogic tests the token expiration detection without OAuth flows
func TestTokenExpirationLogic(t *testing.T) {
	// Test valid token
	validToken := createTestToken("access", "refresh", 3600)
	if validToken.IsExpired(5) {
		t.Error("Valid token should not be expired")
	}

	// Test expired token
	expiredToken := createExpiredToken("access", "refresh")
	if !expiredToken.IsExpired(5) {
		t.Error("Expired token should be expired")
	}

	// Test token expiring soon (within buffer)
	soonExpiredToken := createTestToken("access", "refresh", 240) // 4 minutes
	if !soonExpiredToken.IsExpired(5) {                           // 5 minute buffer
		t.Error("Token expiring within buffer should be considered expired")
	}
}

// TestGetValidTokenWithValidToken tests the getValidToken method with a valid token
func TestGetValidTokenWithValidToken(t *testing.T) {
	// Create temporary directory and set up fake home
	tempDir := t.TempDir()
	fakeHome := filepath.Join(tempDir, "home")
	configDir := filepath.Join(fakeHome, ".config", "fabric")
	os.MkdirAll(configDir, 0755)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", fakeHome)

	// Create a valid token
	validToken := createTestToken("valid_access_token", "refresh_token", 3600)
	tokenPath := filepath.Join(configDir, ".test_oauth")
	data, _ := json.MarshalIndent(validToken, "", "  ")
	os.WriteFile(tokenPath, data, 0600)

	// Create transport
	client := &Client{}
	transport := NewOAuthTransport(client, http.DefaultTransport)

	// Test getValidToken - this should return the valid token without any OAuth flow
	token, err := transport.getValidToken("test")

	if err != nil {
		t.Fatalf("Expected no error with valid token, got: %v", err)
	}

	if token != "valid_access_token" {
		t.Errorf("Expected 'valid_access_token', got '%s'", token)
	}
}

// Benchmark tests
func BenchmarkGeneratePKCE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := generatePKCE()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTokenIsExpired(b *testing.B) {
	token := createTestToken("access", "refresh", 3600)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token.IsExpired(5)
	}
}
