package youtube

import (
	"strings"
	"testing"
)

func TestParseSeconds(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			name:    "integer seconds",
			input:   "42",
			want:    42,
			wantErr: false,
		},
		{
			name:    "fractional seconds",
			input:   "42.567",
			want:    42,
			wantErr: false,
		},
		{
			name:    "zero",
			input:   "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "zero with fraction",
			input:   "0.999",
			want:    0,
			wantErr: false,
		},
		{
			name:    "decimal point at start",
			input:   ".5",
			want:    0,
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSeconds(tt.input)

			// Check error condition
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseSeconds(%q) expected error but got none", tt.input)
				}
				return
			}

			// Check success condition
			if err != nil {
				t.Fatalf("parseSeconds(%q) unexpected error: %v", tt.input, err)
			}

			if got != tt.want {
				t.Errorf("parseSeconds(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestExtractAndValidateVideoId(t *testing.T) {
	yt := NewYouTube()

	tests := []struct {
		name      string
		url       string
		wantId    string
		wantError bool
		errorMsg  string
	}{
		{
			name:      "valid video URL",
			url:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			wantId:    "dQw4w9WgXcQ",
			wantError: false,
		},
		{
			name:      "valid short URL",
			url:       "https://youtu.be/dQw4w9WgXcQ",
			wantId:    "dQw4w9WgXcQ",
			wantError: false,
		},
		{
			name:      "video with playlist URL - should extract video",
			url:       "https://www.youtube.com/watch?v=dQw4w9WgXcQ&list=PLrAXtmErZgOeiKm4sgNOknGvNjby9efdf",
			wantId:    "dQw4w9WgXcQ",
			wantError: false,
		},
		{
			name:      "playlist-only URL",
			url:       "https://www.youtube.com/playlist?list=PLrAXtmErZgOeiKm4sgNOknGvNjby9efdf",
			wantId:    "",
			wantError: true,
			errorMsg:  "URL is a playlist, not a video",
		},
		{
			name:      "invalid URL",
			url:       "https://example.com",
			wantId:    "",
			wantError: true,
			errorMsg:  "invalid YouTube URL",
		},
		{
			name:      "empty URL",
			url:       "",
			wantId:    "",
			wantError: true,
		},
		{
			name:      "malformed URL",
			url:       "not-a-url",
			wantId:    "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yt.extractAndValidateVideoId(tt.url)

			if tt.wantError {
				if err == nil {
					t.Errorf("extractAndValidateVideoId(%q) expected error but got none", tt.url)
					return
				}
				if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("extractAndValidateVideoId(%q) error = %v, want error containing %q", tt.url, err, tt.errorMsg)
				}
				// Verify empty videoId is returned on error
				if got != "" {
					t.Errorf("extractAndValidateVideoId(%q) returned videoId %q on error, want empty string", tt.url, got)
				}
				return
			}

			if err != nil {
				t.Errorf("extractAndValidateVideoId(%q) unexpected error = %v", tt.url, err)
				return
			}

			if got != tt.wantId {
				t.Errorf("extractAndValidateVideoId(%q) = %q, want %q", tt.url, got, tt.wantId)
			}
		})
	}
}
