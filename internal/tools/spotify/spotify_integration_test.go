//go:build integration

// Integration tests for the Spotify API. They need real credentials in
// SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET and run only with
// go test -tags=integration ./internal/tools/spotify/...

package spotify

import (
	"os"
	"testing"
)

// The Joe Rogan Experience show and one of its episodes.
const (
	// cspell:disable-next-line
	testShowID = "4rOoJ6Egrf8K2IrywzwOMk"
	// If Spotify removes this episode, use the ID of any other episode of the show.
	testEpisodeID = "512ojhOuo1ktJprKbVcKyQ"
)

func setupIntegrationClient(t *testing.T) *Spotify {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		t.Skip("Skipping integration test: SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET must be set")
	}

	s := NewSpotify()
	s.ClientId.Value = clientID
	s.ClientSecret.Value = clientSecret

	return s
}

func TestIntegration_GetShowMetadata(t *testing.T) {
	s := setupIntegrationClient(t)

	metadata, err := s.GetShowMetadata(testShowID)
	if err != nil {
		t.Fatalf("GetShowMetadata failed: %v", err)
	}

	if metadata == nil {
		t.Fatal("GetShowMetadata returned nil metadata")
	}

	if metadata.Id != testShowID {
		t.Errorf("Expected show ID %s, got %s", testShowID, metadata.Id)
	}

	if metadata.Name == "" {
		t.Error("Show name should not be empty")
	}

	if metadata.Publisher == "" {
		t.Error("Show publisher should not be empty")
	}

	t.Logf("Show: %s by %s (%d episodes)", metadata.Name, metadata.Publisher, metadata.TotalEpisodes)
}

func TestIntegration_GetEpisodeMetadata(t *testing.T) {
	s := setupIntegrationClient(t)

	metadata, err := s.GetEpisodeMetadata(testEpisodeID)
	if err != nil {
		t.Fatalf("GetEpisodeMetadata failed: %v", err)
	}

	if metadata == nil {
		t.Fatal("GetEpisodeMetadata returned nil metadata")
	}

	if metadata.Id != testEpisodeID {
		t.Errorf("Expected episode ID %s, got %s", testEpisodeID, metadata.Id)
	}

	if metadata.Name == "" {
		t.Error("Episode name should not be empty")
	}

	if metadata.DurationMinutes <= 0 {
		t.Error("Episode duration should be positive")
	}

	t.Logf("Episode: %s (%d minutes)", metadata.Name, metadata.DurationMinutes)
}

func TestIntegration_SearchShows(t *testing.T) {
	s := setupIntegrationClient(t)

	result, err := s.SearchShows("technology podcast", 5)
	if err != nil {
		t.Fatalf("SearchShows failed: %v", err)
	}

	if result == nil {
		t.Fatal("SearchShows returned nil result")
	}

	if len(result.Shows) == 0 {
		t.Error("SearchShows should return at least one result for 'technology podcast'")
	}

	for i, show := range result.Shows {
		t.Logf("Result %d: %s by %s", i+1, show.Name, show.Publisher)
	}
}

func TestIntegration_GetShowEpisodes(t *testing.T) {
	s := setupIntegrationClient(t)

	episodes, err := s.GetShowEpisodes(testShowID, 5)
	if err != nil {
		t.Fatalf("GetShowEpisodes failed: %v", err)
	}

	if len(episodes) == 0 {
		t.Error("GetShowEpisodes should return at least one episode")
	}

	for i, ep := range episodes {
		t.Logf("Episode %d: %s (%d min)", i+1, ep.Name, ep.DurationMinutes)
	}
}

func TestIntegration_GrabMetadataForURL_Show(t *testing.T) {
	s := setupIntegrationClient(t)

	url := "https://open.spotify.com/show/" + testShowID

	metadata, err := s.GrabMetadataForURL(url)
	if err != nil {
		t.Fatalf("GrabMetadataForURL failed: %v", err)
	}

	show, ok := metadata.(*ShowMetadata)
	if !ok {
		t.Fatalf("Expected ShowMetadata, got %T", metadata)
	}

	if show.Id != testShowID {
		t.Errorf("Expected show ID %s, got %s", testShowID, show.Id)
	}
}

func TestIntegration_GrabMetadataForURL_Episode(t *testing.T) {
	s := setupIntegrationClient(t)

	url := "https://open.spotify.com/episode/" + testEpisodeID

	metadata, err := s.GrabMetadataForURL(url)
	if err != nil {
		t.Fatalf("GrabMetadataForURL failed: %v", err)
	}

	episode, ok := metadata.(*EpisodeMetadata)
	if !ok {
		t.Fatalf("Expected EpisodeMetadata, got %T", metadata)
	}

	if episode.Id != testEpisodeID {
		t.Errorf("Expected episode ID %s, got %s", testEpisodeID, episode.Id)
	}
}

func TestIntegration_FormatMetadataAsText(t *testing.T) {
	s := setupIntegrationClient(t)

	metadata, err := s.GrabMetadataForURL("https://open.spotify.com/show/" + testShowID)
	if err != nil {
		t.Fatalf("GrabMetadataForURL failed: %v", err)
	}

	text := s.FormatMetadataAsText(metadata)

	if text == "" {
		t.Error("FormatMetadataAsText returned empty string")
	}

	t.Logf("Formatted metadata:\n%s", text)
}

func TestIntegration_GetShowMetadata_InvalidID(t *testing.T) {
	s := setupIntegrationClient(t)

	_, err := s.GetShowMetadata("invalid_show_id_12345")
	if err == nil {
		t.Error("GetShowMetadata with invalid ID should return an error")
	}
	t.Logf("Expected error for invalid show ID: %v", err)
}

func TestIntegration_GetEpisodeMetadata_InvalidID(t *testing.T) {
	s := setupIntegrationClient(t)

	_, err := s.GetEpisodeMetadata("invalid_episode_id_12345")
	if err == nil {
		t.Error("GetEpisodeMetadata with invalid ID should return an error")
	}
	t.Logf("Expected error for invalid episode ID: %v", err)
}

func TestIntegration_SearchShows_NoResults(t *testing.T) {
	s := setupIntegrationClient(t)

	// cspell:disable-next-line
	result, err := s.SearchShows("xyzzy_nonexistent_podcast_12345_zyxwv", 5)
	if err != nil {
		t.Fatalf("SearchShows failed: %v", err)
	}

	if result == nil {
		t.Fatal("SearchShows returned nil result")
	}

	if len(result.Shows) > 0 {
		t.Logf("WARNING: Unexpectedly found %d results for nonsense query (test may need updating)", len(result.Shows))
	} else {
		t.Log("Search correctly returned 0 results for nonsense query")
	}
}
