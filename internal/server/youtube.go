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

type YouTubeRequest struct {
	URL        string `json:"url"`
	Language   string `json:"language"`
	Timestamps bool   `json:"timestamps"`
}

type YouTubeResponse struct {
	Transcript string `json:"transcript"`
	Title      string `json:"title"`
}

func NewYouTubeHandler(r *gin.Engine, registry *core.PluginRegistry) *YouTubeHandler {
	handler := &YouTubeHandler{yt: registry.YouTube}
	r.POST("/youtube/transcript", handler.Transcript)
	return handler
}

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

	c.JSON(http.StatusOK, YouTubeResponse{Transcript: transcript, Title: videoID})
}
