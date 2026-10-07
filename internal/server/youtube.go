package restapi

import (
	"net/http"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
	"github.com/gin-gonic/gin"
)

type YouTubeHandler struct {
	yt *youtube.YouTube
}

// YouTubeRequest represents a request to get a YouTube video transcript
type YouTubeRequest struct {
	URL        string `json:"url" binding:"required" example:"https://www.youtube.com/watch?v=dQw4w9WgXcQ"` // YouTube video URL (required)
	Language   string `json:"language,omitempty" example:"en"`                                              // Language code for transcript (default: "en")
	Timestamps bool   `json:"timestamps,omitempty" example:"false"`                                         // Include timestamps in the transcript (default: false)
}

// YouTubeResponse represents the response containing video transcript and metadata
type YouTubeResponse struct {
	Transcript  string `json:"transcript" example:"This is the video transcript..."`                // The video transcript text
	VideoId     string `json:"videoId" example:"dQw4w9WgXcQ"`                                       // YouTube video ID
	Title       string `json:"title" example:"Example Video Title"`                                 // Video title from YouTube metadata
	Description string `json:"description" example:"This is the video description from YouTube..."` // Video description from YouTube metadata
}

func NewYouTubeHandler(r *gin.Engine, registry *core.PluginRegistry) *YouTubeHandler {
	handler := &YouTubeHandler{yt: registry.YouTube}
	r.POST("/youtube/transcript", requireJSON, handler.Transcript)
	return handler
}

// Transcript godoc
// @Summary Get YouTube video transcript
// @Description Retrieves the transcript of a YouTube video along with video metadata (title and description)
// @Tags youtube
// @Accept json
// @Produce json
// @Param request body YouTubeRequest true "YouTube transcript request with URL, language, and timestamp options"
// @Success 200 {object} YouTubeResponse "Successful response with transcript and metadata"
// @Failure 400 {object} map[string]string "Bad request - invalid URL or playlist URL provided"
// @Failure 500 {object} map[string]string "Internal server error - failed to retrieve transcript or metadata"
// @Security ApiKeyAuth
// @Router /youtube/transcript [post]
func (h *YouTubeHandler) Transcript(c *gin.Context) {
	var req YouTubeRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}
	language := req.Language
	if language == "" {
		language = "en"
	}

	var videoID, playlistID string
	var err error
	if videoID, playlistID, err = h.yt.GetVideoOrPlaylistId(req.URL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if videoID == "" && playlistID != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is a playlist, not a video"})
		return
	}

	// Try to get metadata (requires valid YouTube API key), but don't fail if unavailable
	// This allows the endpoint to work for transcript extraction even without API key
	var metadata *youtube.VideoMetadata
	var title, description string
	if metadata, err = h.yt.GrabMetadata(videoID); err == nil {
		// Metadata available - use title and description from API
		title = metadata.Title
		description = metadata.Description
	} else {
		// No valid API key or metadata fetch failed - fallback to videoID as title
		title = videoID
		description = ""
	}

	var transcript string
	if req.Timestamps {
		transcript, err = h.yt.GrabTranscriptWithTimestamps(videoID, language)
	} else {
		transcript, err = h.yt.GrabTranscript(videoID, language)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, YouTubeResponse{Transcript: transcript, VideoId: videoID, Title: title, Description: description})
}
