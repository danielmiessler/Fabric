// Package spotify provides Spotify Web API integration for podcast metadata retrieval.
//
// Requirements:
// - Spotify Developer Account: Required to obtain Client ID and Client Secret
// - Client Credentials: Stored in .env file via fabric --setup
//
// The implementation uses OAuth2 Client Credentials flow for authentication.
// Note: The Spotify Web API does NOT provide access to podcast transcripts.
// For transcript functionality, users should use fabric's --transcribe-file feature
// with audio obtained from other sources.
package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/plugins"
)

const (
	// Spotify API endpoints
	tokenURL   = "https://accounts.spotify.com/api/token"
	apiBaseURL = "https://api.spotify.com/v1"
)

// URL pattern regexes for parsing Spotify URLs
var (
	showPatternRegex    = regexp.MustCompile(`spotify\.com/show/([a-zA-Z0-9]+)`)
	episodePatternRegex = regexp.MustCompile(`spotify\.com/episode/([a-zA-Z0-9]+)`)
)

// NewSpotify creates a new Spotify client instance.
func NewSpotify() *Spotify {
	label := "Spotify"

	ret := &Spotify{}

	ret.PluginBase = &plugins.PluginBase{
		Name:             i18n.T("spotify_label"),
		SetupDescription: i18n.T("spotify_setup_description") + " " + i18n.T("optional_marker"),
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
	}

	ret.ClientId = ret.AddSetupQuestionWithEnvName("Client ID", false, i18n.T("spotify_client_id_question"))
	ret.ClientSecret = ret.AddSetupQuestionWithEnvName("Client Secret", false, i18n.T("spotify_client_secret_question"))

	return ret
}

// Spotify represents a Spotify API client.
type Spotify struct {
	*plugins.PluginBase
	ClientId     *plugins.SetupQuestion
	ClientSecret *plugins.SetupQuestion

	// OAuth2 token management
	accessToken string
	tokenExpiry time.Time
	tokenMutex  sync.RWMutex
	httpClient  *http.Client
}

// initClient ensures the HTTP client and access token are initialized.
func (s *Spotify) initClient() error {
	if s.httpClient == nil {
		s.httpClient = &http.Client{Timeout: 30 * time.Second}
	}

	// Check if we need to refresh the token
	s.tokenMutex.RLock()
	needsRefresh := s.accessToken == "" || time.Now().After(s.tokenExpiry)
	s.tokenMutex.RUnlock()

	if needsRefresh {
		return s.refreshAccessToken()
	}
	return nil
}

// refreshAccessToken obtains a new access token using Client Credentials flow.
func (s *Spotify) refreshAccessToken() error {
	if s.ClientId.Value == "" || s.ClientSecret.Value == "" {
		return errors.New(i18n.T("spotify_not_configured"))
	}

	// Prepare the token request
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf(i18n.T("spotify_failed_create_token_request"), err)
	}

	// Set Basic Auth header with Client ID and Secret
	auth := base64.StdEncoding.EncodeToString([]byte(s.ClientId.Value + ":" + s.ClientSecret.Value))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf(i18n.T("spotify_failed_request_access_token"), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(i18n.T("spotify_failed_get_access_token"), resp.StatusCode, string(body))
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return fmt.Errorf(i18n.T("spotify_failed_decode_token_response"), err)
	}

	s.tokenMutex.Lock()
	s.accessToken = tokenResp.AccessToken
	// Set expiry slightly before actual expiry to avoid edge cases
	s.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-60) * time.Second)
	s.tokenMutex.Unlock()

	return nil
}

// doRequest performs an authenticated request to the Spotify API.
func (s *Spotify) doRequest(method, endpoint string) ([]byte, error) {
	if err := s.initClient(); err != nil {
		return nil, err
	}

	reqURL := apiBaseURL + endpoint
	req, err := http.NewRequest(method, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_create_request"), err)
	}

	s.tokenMutex.RLock()
	req.Header.Set("Authorization", "Bearer "+s.accessToken)
	s.tokenMutex.RUnlock()

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_execute_request"), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_read_response_body"), err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(i18n.T("spotify_api_request_failed"), resp.StatusCode, string(body))
	}

	return body, nil
}

// GetShowOrEpisodeId extracts show or episode ID from a Spotify URL.
func (s *Spotify) GetShowOrEpisodeId(urlStr string) (showId string, episodeId string, err error) {
	// Extract show ID
	showMatch := showPatternRegex.FindStringSubmatch(urlStr)
	if len(showMatch) > 1 {
		showId = showMatch[1]
	}

	// Extract episode ID
	episodeMatch := episodePatternRegex.FindStringSubmatch(urlStr)
	if len(episodeMatch) > 1 {
		episodeId = episodeMatch[1]
	}

	if showId == "" && episodeId == "" {
		err = fmt.Errorf(i18n.T("spotify_invalid_url"), urlStr)
	}
	return
}

// ShowMetadata represents metadata for a Spotify show (podcast).
type ShowMetadata struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Publisher     string   `json:"publisher"`
	TotalEpisodes int      `json:"total_episodes"`
	Languages     []string `json:"languages"`
	MediaType     string   `json:"media_type"`
	ExternalURL   string   `json:"external_url"`
	ImageURL      string   `json:"image_url,omitempty"`
}

// EpisodeMetadata represents metadata for a Spotify episode.
type EpisodeMetadata struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ReleaseDate     string `json:"release_date"`
	DurationMs      int    `json:"duration_ms"`
	DurationMinutes int    `json:"duration_minutes"`
	Language        string `json:"language"`
	Explicit        bool   `json:"explicit"`
	ExternalURL     string `json:"external_url"`
	AudioPreviewURL string `json:"audio_preview_url,omitempty"`
	ImageURL        string `json:"image_url,omitempty"`
	ShowId          string `json:"show_id"`
	ShowName        string `json:"show_name"`
}

// SearchResult represents a search result item.
type SearchResult struct {
	Shows []ShowMetadata `json:"shows"`
}

// GetShowMetadata retrieves metadata for a Spotify show (podcast).
func (s *Spotify) GetShowMetadata(showId string) (*ShowMetadata, error) {
	body, err := s.doRequest("GET", "/shows/"+showId)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_error_getting_metadata"), err)
	}

	var resp struct {
		Id            string   `json:"id"`
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		Publisher     string   `json:"publisher"`
		TotalEpisodes int      `json:"total_episodes"`
		Languages     []string `json:"languages"`
		MediaType     string   `json:"media_type"`
		ExternalUrls  struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_parse_show_metadata"), err)
	}

	if resp.Id == "" {
		return nil, fmt.Errorf(i18n.T("spotify_no_show_found"), showId)
	}

	metadata := &ShowMetadata{
		Id:            resp.Id,
		Name:          resp.Name,
		Description:   resp.Description,
		Publisher:     resp.Publisher,
		TotalEpisodes: resp.TotalEpisodes,
		Languages:     resp.Languages,
		MediaType:     resp.MediaType,
		ExternalURL:   resp.ExternalUrls.Spotify,
	}

	if len(resp.Images) > 0 {
		metadata.ImageURL = resp.Images[0].URL
	}

	return metadata, nil
}

// GetEpisodeMetadata retrieves metadata for a Spotify episode.
func (s *Spotify) GetEpisodeMetadata(episodeId string) (*EpisodeMetadata, error) {
	body, err := s.doRequest("GET", "/episodes/"+episodeId)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_error_getting_metadata"), err)
	}

	var resp struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		ReleaseDate  string `json:"release_date"`
		DurationMs   int    `json:"duration_ms"`
		Language     string `json:"language"`
		Explicit     bool   `json:"explicit"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		AudioPreviewUrl string `json:"audio_preview_url"`
		Images          []struct {
			URL string `json:"url"`
		} `json:"images"`
		Show struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"show"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_parse_episode_metadata"), err)
	}

	if resp.Id == "" {
		return nil, fmt.Errorf(i18n.T("spotify_no_episode_found"), episodeId)
	}

	metadata := &EpisodeMetadata{
		Id:              resp.Id,
		Name:            resp.Name,
		Description:     resp.Description,
		ReleaseDate:     resp.ReleaseDate,
		DurationMs:      resp.DurationMs,
		DurationMinutes: resp.DurationMs / 60000,
		Language:        resp.Language,
		Explicit:        resp.Explicit,
		ExternalURL:     resp.ExternalUrls.Spotify,
		AudioPreviewURL: resp.AudioPreviewUrl,
		ShowId:          resp.Show.Id,
		ShowName:        resp.Show.Name,
	}

	if len(resp.Images) > 0 {
		metadata.ImageURL = resp.Images[0].URL
	}

	return metadata, nil
}

// SearchShows searches for podcasts/shows matching the query.
func (s *Spotify) SearchShows(query string, limit int) (*SearchResult, error) {
	if limit <= 0 || limit > 50 {
		limit = 20 // Default limit
	}

	endpoint := fmt.Sprintf("/search?q=%s&type=show&limit=%d", url.QueryEscape(query), limit)
	body, err := s.doRequest("GET", endpoint)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_search_failed"), err)
	}

	var resp struct {
		Shows struct {
			Items []struct {
				Id            string   `json:"id"`
				Name          string   `json:"name"`
				Description   string   `json:"description"`
				Publisher     string   `json:"publisher"`
				TotalEpisodes int      `json:"total_episodes"`
				Languages     []string `json:"languages"`
				MediaType     string   `json:"media_type"`
				ExternalUrls  struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Images []struct {
					URL string `json:"url"`
				} `json:"images"`
			} `json:"items"`
		} `json:"shows"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_parse_search_results"), err)
	}

	result := &SearchResult{
		Shows: make([]ShowMetadata, 0, len(resp.Shows.Items)),
	}

	for _, item := range resp.Shows.Items {
		show := ShowMetadata{
			Id:            item.Id,
			Name:          item.Name,
			Description:   item.Description,
			Publisher:     item.Publisher,
			TotalEpisodes: item.TotalEpisodes,
			Languages:     item.Languages,
			MediaType:     item.MediaType,
			ExternalURL:   item.ExternalUrls.Spotify,
		}
		if len(item.Images) > 0 {
			show.ImageURL = item.Images[0].URL
		}
		result.Shows = append(result.Shows, show)
	}

	return result, nil
}

// GetShowEpisodes retrieves episodes for a given show.
func (s *Spotify) GetShowEpisodes(showId string, limit int) ([]EpisodeMetadata, error) {
	if limit <= 0 || limit > 50 {
		limit = 20 // Default limit
	}

	endpoint := fmt.Sprintf("/shows/%s/episodes?limit=%d", showId, limit)
	body, err := s.doRequest("GET", endpoint)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_get_show_episodes"), err)
	}

	var resp struct {
		Items []struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			ReleaseDate  string `json:"release_date"`
			DurationMs   int    `json:"duration_ms"`
			Language     string `json:"language"`
			Explicit     bool   `json:"explicit"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			AudioPreviewUrl string `json:"audio_preview_url"`
			Images          []struct {
				URL string `json:"url"`
			} `json:"images"`
		} `json:"items"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf(i18n.T("spotify_failed_parse_episodes"), err)
	}

	episodes := make([]EpisodeMetadata, 0, len(resp.Items))
	for _, item := range resp.Items {
		ep := EpisodeMetadata{
			Id:              item.Id,
			Name:            item.Name,
			Description:     item.Description,
			ReleaseDate:     item.ReleaseDate,
			DurationMs:      item.DurationMs,
			DurationMinutes: item.DurationMs / 60000,
			Language:        item.Language,
			Explicit:        item.Explicit,
			ExternalURL:     item.ExternalUrls.Spotify,
			AudioPreviewURL: item.AudioPreviewUrl,
			ShowId:          showId,
		}
		if len(item.Images) > 0 {
			ep.ImageURL = item.Images[0].URL
		}
		episodes = append(episodes, ep)
	}

	return episodes, nil
}

// GrabMetadataForURL retrieves metadata for a Spotify URL (show or episode).
func (s *Spotify) GrabMetadataForURL(urlStr string) (any, error) {
	showId, episodeId, err := s.GetShowOrEpisodeId(urlStr)
	if err != nil {
		return nil, err
	}

	if episodeId != "" {
		return s.GetEpisodeMetadata(episodeId)
	}

	if showId != "" {
		return s.GetShowMetadata(showId)
	}

	return nil, fmt.Errorf(i18n.T("spotify_invalid_url"), urlStr)
}

// FormatMetadataAsText formats metadata as human-readable text suitable for LLM processing.
func (s *Spotify) FormatMetadataAsText(metadata any) string {
	var sb strings.Builder

	switch m := metadata.(type) {
	case *ShowMetadata:
		sb.WriteString(i18n.T("spotify_show_header") + "\n\n")
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_title_label")+"\n", m.Name))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_publisher_label")+"\n", m.Publisher))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_total_episodes_label")+"\n", m.TotalEpisodes))
		if len(m.Languages) > 0 {
			sb.WriteString(fmt.Sprintf(i18n.T("spotify_languages_label")+"\n", strings.Join(m.Languages, ", ")))
		}
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_media_type_label")+"\n", m.MediaType))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_url_label")+"\n\n", m.ExternalURL))
		sb.WriteString(i18n.T("spotify_description_header") + "\n\n")
		sb.WriteString(m.Description)
		sb.WriteString("\n")

	case *EpisodeMetadata:
		sb.WriteString(i18n.T("spotify_episode_header") + "\n\n")
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_title_label")+"\n", m.Name))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_show_name_label")+"\n", m.ShowName))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_release_date_label")+"\n", m.ReleaseDate))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_duration_label")+"\n", m.DurationMinutes))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_language_field_label")+"\n", m.Language))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_explicit_label")+"\n", m.Explicit))
		sb.WriteString(fmt.Sprintf(i18n.T("spotify_url_label")+"\n", m.ExternalURL))
		if m.AudioPreviewURL != "" {
			sb.WriteString(fmt.Sprintf(i18n.T("spotify_audio_preview_label")+"\n", m.AudioPreviewURL))
		}
		sb.WriteString("\n" + i18n.T("spotify_description_header") + "\n\n")
		sb.WriteString(m.Description)
		sb.WriteString("\n")

	case *SearchResult:
		sb.WriteString(i18n.T("spotify_search_results_header") + "\n\n")
		for i, show := range m.Shows {
			sb.WriteString(fmt.Sprintf("## %d. %s\n", i+1, show.Name))
			sb.WriteString(fmt.Sprintf(i18n.T("spotify_search_publisher_label")+"\n", show.Publisher))
			sb.WriteString(fmt.Sprintf(i18n.T("spotify_search_episodes_label")+"\n", show.TotalEpisodes))
			sb.WriteString(fmt.Sprintf(i18n.T("spotify_search_url_label")+"\n", show.ExternalURL))
			// Truncate description for search results
			desc := show.Description
			if len(desc) > 200 {
				desc = desc[:200] + "..."
			}
			sb.WriteString(fmt.Sprintf(i18n.T("spotify_search_description_label")+"\n\n", desc))
		}
	}

	return sb.String()
}
