package spotify

import (
	"strings"
	"testing"
)

func TestGetShowOrEpisodeId(t *testing.T) {
	s := NewSpotify()

	tests := []struct {
		name          string
		url           string
		wantShowId    string
		wantEpisodeId string
		wantError     bool
		errorMsg      string
	}{
		{
			name: "valid show URL",
			url:  "https://open.spotify.com/show/4rOoJ6Egrf8K2IrywzwOMk",
			// cspell:disable-next-line
			wantShowId:    "4rOoJ6Egrf8K2IrywzwOMk",
			wantEpisodeId: "",
			wantError:     false,
		},
		{
			name:          "valid episode URL",
			url:           "https://open.spotify.com/episode/512ojhOuo1ktJprKbVcKyQ",
			wantShowId:    "",
			wantEpisodeId: "512ojhOuo1ktJprKbVcKyQ",
			wantError:     false,
		},
		{
			name: "show URL with query params",
			url:  "https://open.spotify.com/show/4rOoJ6Egrf8K2IrywzwOMk?si=abc123",
			// cspell:disable-next-line
			wantShowId:    "4rOoJ6Egrf8K2IrywzwOMk",
			wantEpisodeId: "",
			wantError:     false,
		},
		{
			name:          "episode URL with query params",
			url:           "https://open.spotify.com/episode/512ojhOuo1ktJprKbVcKyQ?si=def456",
			wantShowId:    "",
			wantEpisodeId: "512ojhOuo1ktJprKbVcKyQ",
			wantError:     false,
		},
		{
			name:          "invalid URL - no show or episode",
			url:           "https://open.spotify.com/track/4uLU6hMCjMI75M1A2tKUQC",
			wantShowId:    "",
			wantEpisodeId: "",
			wantError:     true,
			errorMsg:      "invalid Spotify URL",
		},
		{
			name:          "invalid URL - not spotify",
			url:           "https://example.com/show/123",
			wantShowId:    "",
			wantEpisodeId: "",
			wantError:     true,
			errorMsg:      "invalid Spotify URL",
		},
		{
			name:          "empty URL",
			url:           "",
			wantShowId:    "",
			wantEpisodeId: "",
			wantError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			showId, episodeId, err := s.GetShowOrEpisodeId(tt.url)

			if tt.wantError {
				if err == nil {
					t.Errorf("GetShowOrEpisodeId(%q) expected error but got none", tt.url)
					return
				}
				if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("GetShowOrEpisodeId(%q) error = %v, want error containing %q", tt.url, err, tt.errorMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("GetShowOrEpisodeId(%q) unexpected error = %v", tt.url, err)
				return
			}

			if showId != tt.wantShowId {
				t.Errorf("GetShowOrEpisodeId(%q) showId = %q, want %q", tt.url, showId, tt.wantShowId)
			}

			if episodeId != tt.wantEpisodeId {
				t.Errorf("GetShowOrEpisodeId(%q) episodeId = %q, want %q", tt.url, episodeId, tt.wantEpisodeId)
			}
		})
	}
}

func TestFormatMetadataAsText_ShowMetadata(t *testing.T) {
	s := NewSpotify()

	show := &ShowMetadata{
		Id:            "test123",
		Name:          "Test Podcast",
		Description:   "A test podcast description",
		Publisher:     "Test Publisher",
		TotalEpisodes: 100,
		Languages:     []string{"en", "es"},
		MediaType:     "audio",
		ExternalURL:   "https://open.spotify.com/show/test123",
	}

	result := s.FormatMetadataAsText(show)

	if !strings.Contains(result, "# Spotify Podcast/Show") {
		t.Error("FormatMetadataAsText missing header for show")
	}
	if !strings.Contains(result, "**Title**: Test Podcast") {
		t.Error("FormatMetadataAsText missing title")
	}
	if !strings.Contains(result, "**Publisher**: Test Publisher") {
		t.Error("FormatMetadataAsText missing publisher")
	}
	if !strings.Contains(result, "**Total Episodes**: 100") {
		t.Error("FormatMetadataAsText missing total episodes")
	}
	if !strings.Contains(result, "en, es") {
		t.Error("FormatMetadataAsText missing languages")
	}
	if !strings.Contains(result, "A test podcast description") {
		t.Error("FormatMetadataAsText missing description")
	}
}

func TestFormatMetadataAsText_EpisodeMetadata(t *testing.T) {
	s := NewSpotify()

	episode := &EpisodeMetadata{
		Id:              "ep123",
		Name:            "Test Episode",
		Description:     "A test episode description",
		ReleaseDate:     "2024-01-15",
		DurationMs:      3600000,
		DurationMinutes: 60,
		Language:        "en",
		Explicit:        false,
		ExternalURL:     "https://open.spotify.com/episode/ep123",
		ShowId:          "show123",
		ShowName:        "Test Show",
	}

	result := s.FormatMetadataAsText(episode)

	if !strings.Contains(result, "# Spotify Episode") {
		t.Error("FormatMetadataAsText missing header for episode")
	}
	if !strings.Contains(result, "**Title**: Test Episode") {
		t.Error("FormatMetadataAsText missing title")
	}
	if !strings.Contains(result, "**Show**: Test Show") {
		t.Error("FormatMetadataAsText missing show name")
	}
	if !strings.Contains(result, "**Release Date**: 2024-01-15") {
		t.Error("FormatMetadataAsText missing release date")
	}
	if !strings.Contains(result, "**Duration**: 60 minutes") {
		t.Error("FormatMetadataAsText missing duration")
	}
	if !strings.Contains(result, "A test episode description") {
		t.Error("FormatMetadataAsText missing description")
	}
}

func TestFormatMetadataAsText_SearchResult(t *testing.T) {
	s := NewSpotify()

	searchResult := &SearchResult{
		Shows: []ShowMetadata{
			{
				Id:            "show1",
				Name:          "First Show",
				Description:   "First show description",
				Publisher:     "Publisher One",
				TotalEpisodes: 50,
				ExternalURL:   "https://open.spotify.com/show/show1",
			},
			{
				Id:            "show2",
				Name:          "Second Show",
				Description:   "Second show description",
				Publisher:     "Publisher Two",
				TotalEpisodes: 25,
				ExternalURL:   "https://open.spotify.com/show/show2",
			},
		},
	}

	result := s.FormatMetadataAsText(searchResult)

	if !strings.Contains(result, "# Spotify Search Results") {
		t.Error("FormatMetadataAsText missing header for search results")
	}
	if !strings.Contains(result, "## 1. First Show") {
		t.Error("FormatMetadataAsText missing first show")
	}
	if !strings.Contains(result, "## 2. Second Show") {
		t.Error("FormatMetadataAsText missing second show")
	}
	if !strings.Contains(result, "**Publisher**: Publisher One") {
		t.Error("FormatMetadataAsText missing publisher for first show")
	}
	if !strings.Contains(result, "**Episodes**: 50") {
		t.Error("FormatMetadataAsText missing episode count")
	}
}

func TestFormatMetadataAsText_NilAndUnknownTypes(t *testing.T) {
	s := NewSpotify()

	result := s.FormatMetadataAsText(nil)
	if result != "" {
		t.Errorf("FormatMetadataAsText(nil) should return empty string, got %q", result)
	}

	result = s.FormatMetadataAsText("unexpected string type")
	if result != "" {
		t.Errorf("FormatMetadataAsText(string) should return empty string, got %q", result)
	}

	result = s.FormatMetadataAsText(12345)
	if result != "" {
		t.Errorf("FormatMetadataAsText(int) should return empty string, got %q", result)
	}
}

func TestNewSpotify(t *testing.T) {
	s := NewSpotify()

	if s == nil {
		t.Fatal("NewSpotify() returned nil")
	}

	if s.PluginBase == nil {
		t.Error("NewSpotify() PluginBase is nil")
	}

	if s.ClientId == nil {
		t.Error("NewSpotify() ClientId is nil")
	}

	if s.ClientSecret == nil {
		t.Error("NewSpotify() ClientSecret is nil")
	}
}

func TestSpotify_IsConfigured(t *testing.T) {
	s := NewSpotify()

	// ClientId and ClientSecret are optional, so IsConfigured returns true when they are empty.
	if !s.IsConfigured() {
		t.Error("NewSpotify() should be configured (optional settings are valid when empty)")
	}

	s.ClientId.Value = "test_client_id"
	s.ClientSecret.Value = "test_client_secret"

	if !s.IsConfigured() {
		t.Error("Spotify should be configured after setting credentials")
	}
}

func TestSpotify_HasCredentials(t *testing.T) {
	s := NewSpotify()

	if s.ClientId.Value != "" || s.ClientSecret.Value != "" {
		t.Error("NewSpotify() should have empty credentials initially")
	}

	s.ClientId.Value = "test_client_id"
	s.ClientSecret.Value = "test_client_secret"

	if s.ClientId.Value != "test_client_id" {
		t.Error("ClientId should be set")
	}
	if s.ClientSecret.Value != "test_client_secret" {
		t.Error("ClientSecret should be set")
	}
}
