package cli

import (
	"errors"
	"fmt"

	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/danielmiessler/fabric/internal/tools/youtube"
)

// handleToolProcessing runs the YouTube, web scrape, and Spotify tools and collects their output.
// When the run is not a chat request, it also prints the output.
func handleToolProcessing(currentFlags *Flags, registry *core.PluginRegistry) (messageTools string, err error) {
	if currentFlags.YouTube != "" {
		if !registry.YouTube.IsConfigured() {
			err = errors.New(i18n.T("youtube_not_configured"))
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
			err = errors.New(i18n.T("scraping_not_configured"))
			return
		}
		if currentFlags.ScrapeURL != "" {
			var website string
			if website, err = registry.Jina.ScrapeURL(currentFlags.ScrapeURL); err != nil {
				return
			}
			messageTools = AppendMessage(messageTools, website)
		}

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

	if currentFlags.Spotify != "" {
		if !registry.Spotify.IsConfigured() {
			err = errors.New(i18n.T("spotify_not_configured"))
			return
		}

		var metadata any
		if metadata, err = registry.Spotify.GrabMetadataForURL(currentFlags.Spotify); err != nil {
			return
		}

		formattedMetadata := registry.Spotify.FormatMetadataAsText(metadata)
		messageTools = AppendMessage(messageTools, formattedMetadata)

		if !currentFlags.IsChatRequest() {
			err = currentFlags.WriteOutput(messageTools)
			return
		}
	}

	return
}
