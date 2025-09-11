package cli

import (
	"fmt"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
)

// handleToolProcessing handles YouTube and web scraping tool processing
func handleToolProcessing(currentFlags *Flags, registry *core.PluginRegistry) (messageTools string, err error) {
	if currentFlags.YouTube != "" {
		if !registry.YouTube.IsConfigured() {
			err = fmt.Errorf("%s", i18n.T("youtube_not_configured"))
			return
		}

		var videoId string
		var playlistId string
		if videoId, playlistId, err = registry.YouTube.GetVideoOrPlaylistId(currentFlags.YouTube); err != nil {
			return
		} else if (videoId == "" || currentFlags.YouTubePlaylist) && playlistId != "" {
			if currentFlags.Output != "" {
				err = registry.YouTube.FetchAndSavePlaylist(playlistId, currentFlags.Output)
			} else {
				var videos []*youtube.VideoMeta
				if videos, err = registry.YouTube.FetchPlaylistVideos(playlistId); err != nil {
					err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_fetching_playlist_videos"), err))
					return
				}

				for _, video := range videos {
					var message string
					if message, err = processYoutubeVideo(currentFlags, registry, video.Id); err != nil {
						return
					}

					if !currentFlags.IsChatRequest() {
						if err = WriteOutput(message, fmt.Sprintf("%v.md", video.TitleNormalized)); err != nil {
							return
						}
					} else {
						messageTools = AppendMessage(messageTools, message)
					}
				}
			}
			return
		}

		if messageTools, err = processYoutubeVideo(currentFlags, registry, videoId); err != nil {
			return
		}
		if !currentFlags.IsChatRequest() {
			err = currentFlags.WriteOutput(messageTools)
			return
		}
	}

	if currentFlags.ScrapeURL != "" || currentFlags.ScrapeQuestion != "" {
		if !registry.Jina.IsConfigured() {
			err = fmt.Errorf("%s", i18n.T("scraping_not_configured"))
			return
		}
		// Check if the scrape_url flag is set and call ScrapeURL
		if currentFlags.ScrapeURL != "" {
			var website string
			if website, err = registry.Jina.ScrapeURL(currentFlags.ScrapeURL); err != nil {
				return
			}
			messageTools = AppendMessage(messageTools, website)
		}

		// Check if the scrape_question flag is set and call ScrapeQuestion
		if currentFlags.ScrapeQuestion != "" {
			var website string
			if website, err = registry.Jina.ScrapeQuestion(currentFlags.ScrapeQuestion); err != nil {
				return
			}

			messageTools = AppendMessage(messageTools, website)
		}

		if !currentFlags.IsChatRequest() {
			err = currentFlags.WriteOutput(messageTools)
			return
		}
	}

	return
}
