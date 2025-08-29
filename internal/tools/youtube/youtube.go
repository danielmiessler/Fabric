// Package youtube provides YouTube video transcript and comment extraction functionality.
//
// Requirements:
// - yt-dlp: Required for transcript extraction (must be installed separately)
// - YouTube API key: Optional, only needed for comments and metadata extraction
//
// The implementation uses yt-dlp for reliable transcript extraction and the YouTube API
// for comments/metadata. Old YouTube scraping methods have been removed due to
// frequent changes and rate limiting.
package youtube

import (
	"bytes"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/plugins"
	"github.com/kballard/go-shellquote"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var timestampRegex *regexp.Regexp
var languageFileRegex *regexp.Regexp
var videoPatternRegex *regexp.Regexp
var playlistPatternRegex *regexp.Regexp
var vttTagRegex *regexp.Regexp
var durationRegex *regexp.Regexp

const TimeGapForRepeats = 10 // seconds

func init() {
	// Match timestamps like "00:00:01.234" or just numbers or sequence numbers
	timestampRegex = regexp.MustCompile(`^\d+$|^\d{1,2}:\d{2}(:\d{2})?(\.\d{3})?$`)
	// Match language-specific VTT files like .en.vtt, .es.vtt, .en-US.vtt, .pt-BR.vtt
	languageFileRegex = regexp.MustCompile(`\.[a-z]{2}(-[A-Z]{2})?\.vtt$`)
	// YouTube video ID pattern
	videoPatternRegex = regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:live\/|[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|(?:s(?:horts)\/)|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]*)`)
	// YouTube playlist ID pattern
	playlistPatternRegex = regexp.MustCompile(`[?&]list=([a-zA-Z0-9_-]+)`)
	// VTT formatting tags like <c.colorE5E5E5>, </c>, etc.
	vttTagRegex = regexp.MustCompile(`<[^>]*>`)
	// YouTube duration format PT1H2M3S
	durationRegex = regexp.MustCompile(`(?i)PT(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)
}

func NewYouTube() (ret *YouTube) {

	label := "YouTube"
	ret = &YouTube{}

	ret.PluginBase = &plugins.PluginBase{
		Name:             label,
		SetupDescription: label + " - to grab video transcripts (via yt-dlp) and comments/metadata (via YouTube API)",
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
	}

	ret.ApiKey = ret.AddSetupQuestion("API key", true)

	return
}

type YouTube struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion

	normalizeRegex *regexp.Regexp
	service        *youtube.Service
}

func (o *YouTube) initService() (err error) {
	if o.service == nil {
		if o.ApiKey.Value == "" {
			err = fmt.Errorf("YouTube API key required for comments and metadata. Run 'fabric --setup' to configure")
			return
		}
		o.normalizeRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
		ctx := context.Background()
		o.service, err = youtube.NewService(ctx, option.WithAPIKey(o.ApiKey.Value))
	}
	return
}

func (o *YouTube) GetVideoOrPlaylistId(url string) (videoId string, playlistId string, err error) {
	// Extract video ID using pre-compiled regex
	videoMatch := videoPatternRegex.FindStringSubmatch(url)
	if len(videoMatch) > 1 {
		videoId = videoMatch[1]
	}

	// Extract playlist ID using pre-compiled regex
	playlistMatch := playlistPatternRegex.FindStringSubmatch(url)
	if len(playlistMatch) > 1 {
		playlistId = playlistMatch[1]
	}

	if videoId == "" && playlistId == "" {
		err = fmt.Errorf("invalid YouTube URL, can't get video or playlist ID: '%s'", url)
	}
	return
}

func (o *YouTube) GrabTranscriptForUrl(url string, language string) (ret string, err error) {
	var videoId string
	var playlistId string
	if videoId, playlistId, err = o.GetVideoOrPlaylistId(url); err != nil {
		return
	} else if videoId == "" && playlistId != "" {
		err = fmt.Errorf("URL is a playlist, not a video")
		return
	}

	return o.GrabTranscript(videoId, language)
}

func (o *YouTube) GrabTranscript(videoId string, language string) (ret string, err error) {
	// Use yt-dlp for reliable transcript extraction
	return o.GrabTranscriptWithArgs(videoId, language, "")
}

func (o *YouTube) GrabTranscriptWithArgs(videoId string, language string, additionalArgs string) (ret string, err error) {
	// Use yt-dlp for reliable transcript extraction
	return o.tryMethodYtDlp(videoId, language, additionalArgs)
}

func (o *YouTube) GrabTranscriptWithTimestamps(videoId string, language string) (ret string, err error) {
	// Use yt-dlp for reliable transcript extraction with timestamps
	return o.GrabTranscriptWithTimestampsWithArgs(videoId, language, "")
}

func (o *YouTube) GrabTranscriptWithTimestampsWithArgs(videoId string, language string, additionalArgs string) (ret string, err error) {
	// Use yt-dlp for reliable transcript extraction with timestamps
	return o.tryMethodYtDlpWithTimestamps(videoId, language, additionalArgs)
}

// tryMethodYtDlpInternal is a helper function to reduce duplication between
// tryMethodYtDlp and tryMethodYtDlpWithTimestamps.
func (o *YouTube) tryMethodYtDlpInternal(videoId string, language string, additionalArgs string, processVTTFileFunc func(filename string) (string, error)) (ret string, err error) {
	// Check if yt-dlp is available
	if _, err = exec.LookPath("yt-dlp"); err != nil {
		err = fmt.Errorf("yt-dlp not found in PATH. Please install yt-dlp to use YouTube transcript functionality")
		return
	}

	// Create a temporary directory for yt-dlp output (cross-platform)
	tempDir := filepath.Join(os.TempDir(), "fabric-youtube-"+videoId)
	if err = os.MkdirAll(tempDir, 0755); err != nil {
		err = fmt.Errorf("failed to create temp directory: %v", err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Use yt-dlp to get transcript
	videoURL := "https://www.youtube.com/watch?v=" + videoId
	outputPath := filepath.Join(tempDir, "%(title)s.%(ext)s")

	baseArgs := []string{
		"--write-auto-subs",
		"--skip-download",
		"--sub-format", "vtt",
		"--quiet",
		"--no-warnings",
		"-o", outputPath,
	}

	args := append([]string{}, baseArgs...)

	// Add built-in language selection first
	if language != "" {
		langMatch := language
		if len(langMatch) > 2 {
			langMatch = langMatch[:2]
		}
		langOpts := language + "," + langMatch + ".*," + langMatch
		args = append(args, "--sub-langs", langOpts)
	}

	// Add user-provided arguments last so they take precedence
	if additionalArgs != "" {
		additionalArgsList, err := shellquote.Split(additionalArgs)
		if err != nil {
			return "", fmt.Errorf("invalid yt-dlp arguments: %v", err)
		}
		args = append(args, additionalArgsList...)
	}

	args = append(args, videoURL)

	cmd := exec.Command("yt-dlp", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		stderrStr := stderr.String()

		// Check for specific YouTube errors
		if strings.Contains(stderrStr, "429") || strings.Contains(stderrStr, "Too Many Requests") {
			err = fmt.Errorf("YouTube rate limit exceeded. Try again later or use different yt-dlp arguments like '--sleep-requests 1' to slow down requests. Error: %v", err)
			return
		}

		if strings.Contains(stderrStr, "Sign in to confirm you're not a bot") || strings.Contains(stderrStr, "Use --cookies-from-browser") {
			err = fmt.Errorf("YouTube requires authentication (bot detection). Use --yt-dlp-args='--cookies-from-browser BROWSER' where BROWSER is chrome, firefox, brave, etc. Error: %v", err)
			return
		}

		if language != "" {
			// Fallback: try without specifying language (let yt-dlp choose best available)
			stderr.Reset()
			fallbackArgs := append([]string{}, baseArgs...)

			// Add additional arguments if provided
			if additionalArgs != "" {
				additionalArgsList, parseErr := shellquote.Split(additionalArgs)
				if parseErr != nil {
					return "", fmt.Errorf("invalid yt-dlp arguments: %v", parseErr)
				}
				fallbackArgs = append(fallbackArgs, additionalArgsList...)
			}

			// Don't specify language, let yt-dlp choose
			fallbackArgs = append(fallbackArgs, videoURL)
			cmd = exec.Command("yt-dlp", fallbackArgs...)
			cmd.Stderr = &stderr
			if err = cmd.Run(); err != nil {
				stderrStr2 := stderr.String()
				if strings.Contains(stderrStr2, "429") || strings.Contains(stderrStr2, "Too Many Requests") {
					err = fmt.Errorf("YouTube rate limit exceeded. Try again later or use different yt-dlp arguments like '--sleep-requests 1'. Error: %v", err)
				} else {
					err = fmt.Errorf("yt-dlp failed with language '%s' and fallback. Original error: %s. Fallback error: %s", language, stderrStr, stderrStr2)
				}
				return
			}
		} else {
			err = fmt.Errorf("yt-dlp failed: %v, stderr: %s", err, stderrStr)
			return
		}
	}

	// Find VTT files using cross-platform approach
	// Try to find files with the requested language first, but fall back to any VTT file
	vttFiles, err := o.findVTTFilesWithFallback(tempDir, language)
	if err != nil {
		return "", err
	}

	return processVTTFileFunc(vttFiles[0])
}

func (o *YouTube) tryMethodYtDlp(videoId string, language string, additionalArgs string) (ret string, err error) {
	return o.tryMethodYtDlpInternal(videoId, language, additionalArgs, o.readAndCleanVTTFile)
}

func (o *YouTube) tryMethodYtDlpWithTimestamps(videoId string, language string, additionalArgs string) (ret string, err error) {
	return o.tryMethodYtDlpInternal(videoId, language, additionalArgs, o.readAndFormatVTTWithTimestamps)
}

func (o *YouTube) readAndCleanVTTFile(filename string) (ret string, err error) {
	var content []byte
	if content, err = os.ReadFile(filename); err != nil {
		return
	}

	// Convert VTT to plain text
	lines := strings.Split(string(content), "\n")
	var textBuilder strings.Builder
	seenSegments := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Skip WEBVTT header, timestamps, and empty lines
		if line == "" || line == "WEBVTT" || strings.Contains(line, "-->") ||
			strings.HasPrefix(line, "NOTE") || strings.HasPrefix(line, "STYLE") ||
			strings.HasPrefix(line, "Kind:") || strings.HasPrefix(line, "Language:") ||
			isTimeStamp(line) {
			continue
		}
		// Remove VTT formatting tags
		line = removeVTTTags(line)
		if line != "" {
			if _, exists := seenSegments[line]; !exists {
				textBuilder.WriteString(line)
				textBuilder.WriteString(" ")
				seenSegments[line] = struct{}{}
			}
		}
	}

	ret = strings.TrimSpace(textBuilder.String())
	if ret == "" {
		err = fmt.Errorf("no transcript content found in VTT file")
	}
	return
}

func (o *YouTube) readAndFormatVTTWithTimestamps(filename string) (ret string, err error) {
	var content []byte
	if content, err = os.ReadFile(filename); err != nil {
		return
	}

	// Parse VTT and preserve timestamps
	lines := strings.Split(string(content), "\n")
	var textBuilder strings.Builder
	var currentTimestamp string
	// Track content with timestamps to allow repeats after significant time gaps
	// This preserves legitimate repeated content (choruses, recurring phrases, etc.)
	// while still filtering out immediate duplicates from VTT formatting issues
	seenSegments := make(map[string]string) // text -> last timestamp seen

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip WEBVTT header and empty lines
		if line == "" || line == "WEBVTT" || strings.HasPrefix(line, "NOTE") ||
			strings.HasPrefix(line, "STYLE") || strings.HasPrefix(line, "Kind:") ||
			strings.HasPrefix(line, "Language:") {
			continue
		}

		// Check if this line is a timestamp
		if strings.Contains(line, "-->") {
			// Extract start time for this segment
			parts := strings.Split(line, " --> ")
			if len(parts) >= 1 {
				currentTimestamp = formatVTTTimestamp(parts[0])
			}
			continue
		}

		// Skip numeric sequence identifiers
		if isTimeStamp(line) && !strings.Contains(line, ":") {
			continue
		}

		// This should be transcript text
		if line != "" {
			// Remove VTT formatting tags
			cleanText := removeVTTTags(line)
			if cleanText != "" && currentTimestamp != "" {
				// Check if we should include this segment
				shouldInclude := true
				if lastTimestamp, exists := seenSegments[cleanText]; exists {
					// Calculate time difference to determine if this is a legitimate repeat
					if !shouldIncludeRepeat(lastTimestamp, currentTimestamp) {
						shouldInclude = false
					}
				}

				if shouldInclude {
					timestampedLine := fmt.Sprintf("[%s] %s", currentTimestamp, cleanText)
					textBuilder.WriteString(timestampedLine + "\n")
					seenSegments[cleanText] = currentTimestamp
				}
			}
		}
	}

	ret = strings.TrimSpace(textBuilder.String())
	if ret == "" {
		err = fmt.Errorf("no transcript content found in VTT file")
	}
	return
}

func formatVTTTimestamp(vttTime string) string {
	// VTT timestamps are in format "00:00:01.234" - convert to "00:00:01"
	parts := strings.Split(vttTime, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return vttTime
}

func isTimeStamp(s string) bool {
	return timestampRegex.MatchString(s)
}

func removeVTTTags(s string) string {
	// Remove VTT tags like <c.colorE5E5E5>, </c>, etc.
	return vttTagRegex.ReplaceAllString(s, "")
}

// shouldIncludeRepeat determines if repeated content should be included based on time gap
func shouldIncludeRepeat(lastTimestamp, currentTimestamp string) bool {
	// Parse timestamps to calculate time difference
	lastSeconds, err1 := parseTimestampToSeconds(lastTimestamp)
	currentSeconds, err2 := parseTimestampToSeconds(currentTimestamp)

	if err1 != nil || err2 != nil {
		// If we can't parse timestamps, err on the side of inclusion
		return true
	}

	// Allow repeats if there's at least a TimeGapForRepeats gap
	// This threshold can be adjusted based on use case:
	// - 10 seconds works well for most content
	// - Could be made configurable in the future
	timeDiffSeconds := currentSeconds - lastSeconds
	return timeDiffSeconds >= TimeGapForRepeats
}

// parseTimestampToSeconds converts timestamp string (HH:MM:SS or MM:SS) to total seconds
func parseTimestampToSeconds(timestamp string) (int, error) {
	parts := strings.Split(timestamp, ":")
	if len(parts) < 2 || len(parts) > 3 {
		return 0, fmt.Errorf("invalid timestamp format: %s", timestamp)
	}

	var hours, minutes, seconds int
	var err error

	if len(parts) == 3 {
		// HH:MM:SS format
		if hours, err = strconv.Atoi(parts[0]); err != nil {
			return 0, err
		}
		if minutes, err = strconv.Atoi(parts[1]); err != nil {
			return 0, err
		}
		if seconds, err = parseSeconds(parts[2]); err != nil {
			return 0, err
		}
	} else {
		// MM:SS format
		if minutes, err = strconv.Atoi(parts[0]); err != nil {
			return 0, err
		}
		if seconds, err = parseSeconds(parts[1]); err != nil {
			return 0, err
		}
	}

	return hours*3600 + minutes*60 + seconds, nil
}

func parseSeconds(seconds_str string) (int, error) {
	var seconds int
	var err error
	if strings.Contains(seconds_str, ".") {
		// Handle fractional seconds
		second_parts := strings.Split(seconds_str, ".")
		if seconds, err = strconv.Atoi(second_parts[0]); err != nil {
			return 0, err
		}
	} else {
		if seconds, err = strconv.Atoi(seconds_str); err != nil {
			return 0, err
		}
	}
	return seconds, nil
}

func (o *YouTube) GrabComments(videoId string) (ret []string, err error) {
	if err = o.initService(); err != nil {
		return
	}

	call := o.service.CommentThreads.List([]string{"snippet", "replies"}).VideoId(videoId).TextFormat("plainText").MaxResults(100)
	var response *youtube.CommentThreadListResponse
	if response, err = call.Do(); err != nil {
		log.Printf("Failed to fetch comments: %v", err)
		return
	}

	for _, item := range response.Items {
		topLevelComment := item.Snippet.TopLevelComment.Snippet.TextDisplay
		ret = append(ret, topLevelComment)

		if item.Replies != nil {
			for _, reply := range item.Replies.Comments {
				replyText := reply.Snippet.TextDisplay
				ret = append(ret, "    - "+replyText)
			}
		}
	}
	return
}

func (o *YouTube) GrabDurationForUrl(url string) (ret int, err error) {
	if err = o.initService(); err != nil {
		return
	}

	var videoId string
	var playlistId string
	if videoId, playlistId, err = o.GetVideoOrPlaylistId(url); err != nil {
		return
	} else if videoId == "" && playlistId != "" {
		err = fmt.Errorf("URL is a playlist, not a video")
		return
	}
	return o.GrabDuration(videoId)
}

func (o *YouTube) GrabDuration(videoId string) (ret int, err error) {
	var videoResponse *youtube.VideoListResponse
	if videoResponse, err = o.service.Videos.List([]string{"contentDetails"}).Id(videoId).Do(); err != nil {
		err = fmt.Errorf("error getting video details: %v", err)
		return
	}

	durationStr := videoResponse.Items[0].ContentDetails.Duration

	matches := durationRegex.FindStringSubmatch(durationStr)
	if len(matches) == 0 {
		return 0, fmt.Errorf("invalid duration string: %s", durationStr)
	}

	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])
	seconds, _ := strconv.Atoi(matches[3])

	ret = hours*60 + minutes + seconds/60

	return
}

func (o *YouTube) Grab(url string, options *Options) (ret *VideoInfo, err error) {
	var videoId string
	var playlistId string
	if videoId, playlistId, err = o.GetVideoOrPlaylistId(url); err != nil {
		return
	} else if videoId == "" && playlistId != "" {
		err = fmt.Errorf("URL is a playlist, not a video")
		return
	}

	ret = &VideoInfo{}

	if options.Metadata {
		if ret.Metadata, err = o.GrabMetadata(videoId); err != nil {
			err = fmt.Errorf("error getting video metadata: %v", err)
			return
		}
	}

	if options.Duration {
		if ret.Duration, err = o.GrabDuration(videoId); err != nil {
			err = fmt.Errorf("error parsing video duration: %v", err)
			return
		}

	}

	if options.Comments {
		if ret.Comments, err = o.GrabComments(videoId); err != nil {
			err = fmt.Errorf("error getting comments: %v", err)
			return
		}
	}

	if options.Transcript {
		if ret.Transcript, err = o.GrabTranscript(videoId, "en"); err != nil {
			return
		}
	}

	if options.TranscriptWithTimestamps {
		if ret.Transcript, err = o.GrabTranscriptWithTimestamps(videoId, "en"); err != nil {
			return
		}
	}

	return
}

// FetchPlaylistVideos fetches all videos from a YouTube playlist.
func (o *YouTube) FetchPlaylistVideos(playlistID string) (ret []*VideoMeta, err error) {
	if err = o.initService(); err != nil {
		return
	}

	nextPageToken := ""
	for {
		call := o.service.PlaylistItems.List([]string{"snippet"}).PlaylistId(playlistID).MaxResults(50)
		if nextPageToken != "" {
			call = call.PageToken(nextPageToken)
		}

		var response *youtube.PlaylistItemListResponse
		if response, err = call.Do(); err != nil {
			return
		}

		for _, item := range response.Items {
			videoID := item.Snippet.ResourceId.VideoId
			title := item.Snippet.Title
			ret = append(ret, &VideoMeta{videoID, title, o.normalizeFileName(title)})
		}

		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}

		time.Sleep(1 * time.Second) // Pause to respect API rate limit
	}
	return
}

// SaveVideosToCSV saves the list of videos to a CSV file.
func (o *YouTube) SaveVideosToCSV(filename string, videos []*VideoMeta) (err error) {
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err = writer.Write([]string{"VideoID", "Title"}); err != nil {
		return
	}

	// Write video data
	for _, record := range videos {
		if err = writer.Write([]string{record.Id, record.Title}); err != nil {
			return
		}
	}

	return
}

// FetchAndSavePlaylist fetches all videos in a playlist and saves them to a CSV file.
func (o *YouTube) FetchAndSavePlaylist(playlistID, filename string) (err error) {
	var videos []*VideoMeta
	if videos, err = o.FetchPlaylistVideos(playlistID); err != nil {
		err = fmt.Errorf("error fetching playlist videos: %v", err)
		return
	}

	if err = o.SaveVideosToCSV(filename, videos); err != nil {
		err = fmt.Errorf("error saving videos to CSV: %v", err)
		return
	}

	fmt.Println("Playlist saved to", filename)
	return
}

func (o *YouTube) FetchAndPrintPlaylist(playlistID string) (err error) {
	var videos []*VideoMeta
	if videos, err = o.FetchPlaylistVideos(playlistID); err != nil {
		err = fmt.Errorf("error fetching playlist videos: %v", err)
		return
	}

	fmt.Printf("Playlist: %s\n", playlistID)
	fmt.Printf("VideoId: Title\n")
	for _, video := range videos {
		fmt.Printf("%s: %s\n", video.Id, video.Title)
	}
	return
}

func (o *YouTube) normalizeFileName(name string) string {
	return o.normalizeRegex.ReplaceAllString(name, "_")

}

// findVTTFilesWithFallback searches for VTT files, handling fallback scenarios
// where the requested language might not be available
func (o *YouTube) findVTTFilesWithFallback(dir, requestedLanguage string) ([]string, error) {
	var vttFiles []string

	// Walk through the directory to find VTT files
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".vtt") {
			vttFiles = append(vttFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %v", err)
	}

	if len(vttFiles) == 0 {
		return nil, fmt.Errorf("no VTT files found in directory")
	}

	// If no specific language requested, return the first file
	if requestedLanguage == "" {
		return []string{vttFiles[0]}, nil
	}

	// First, try to find files with the requested language
	for _, file := range vttFiles {
		if strings.Contains(file, "."+requestedLanguage+".vtt") {
			return []string{file}, nil
		}
	}

	// If requested language not found, check if we have any language-specific files
	// This handles the fallback case where yt-dlp downloaded a different language
	for _, file := range vttFiles {
		// Look for any language pattern (e.g., .en.vtt, .es.vtt, etc.)
		if languageFileRegex.MatchString(file) {
			return []string{file}, nil
		}
	}

	// If no language-specific files found, return the first VTT file
	return []string{vttFiles[0]}, nil
}

type VideoMeta struct {
	Id              string
	Title           string
	TitleNormalized string
}

type Options struct {
	Duration                 bool
	Transcript               bool
	TranscriptWithTimestamps bool
	Comments                 bool
	Lang                     string
	Metadata                 bool
}

type VideoInfo struct {
	Transcript string         `json:"transcript"`
	Duration   int            `json:"duration"`
	Comments   []string       `json:"comments"`
	Metadata   *VideoMetadata `json:"metadata,omitempty"`
}

type VideoMetadata struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	PublishedAt  string   `json:"publishedAt"`
	ChannelId    string   `json:"channelId"`
	ChannelTitle string   `json:"channelTitle"`
	CategoryId   string   `json:"categoryId"`
	Tags         []string `json:"tags"`
	ViewCount    uint64   `json:"viewCount"`
	LikeCount    uint64   `json:"likeCount"`
}

func (o *YouTube) GrabMetadata(videoId string) (metadata *VideoMetadata, err error) {
	if err = o.initService(); err != nil {
		return
	}

	call := o.service.Videos.List([]string{"snippet", "statistics"}).Id(videoId)
	var response *youtube.VideoListResponse
	if response, err = call.Do(); err != nil {
		return nil, fmt.Errorf("error getting video metadata: %v", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("no video found with ID: %s", videoId)
	}

	video := response.Items[0]
	viewCount := video.Statistics.ViewCount
	likeCount := video.Statistics.LikeCount

	metadata = &VideoMetadata{
		Id:           video.Id,
		Title:        video.Snippet.Title,
		Description:  video.Snippet.Description,
		PublishedAt:  video.Snippet.PublishedAt,
		ChannelId:    video.Snippet.ChannelId,
		ChannelTitle: video.Snippet.ChannelTitle,
		CategoryId:   video.Snippet.CategoryId,
		Tags:         video.Snippet.Tags,
		ViewCount:    viewCount,
		LikeCount:    likeCount,
	}
	return
}

func (o *YouTube) GrabByFlags() (ret *VideoInfo, err error) {
	options := &Options{}
	flag.BoolVar(&options.Duration, "duration", false, "Output only the duration")
	flag.BoolVar(&options.Transcript, "transcript", false, "Output only the transcript")
	flag.BoolVar(&options.TranscriptWithTimestamps, "transcriptWithTimestamps", false, "Output only the transcript with timestamps")
	flag.BoolVar(&options.Comments, "comments", false, "Output the comments on the video")
	flag.StringVar(&options.Lang, "lang", "en", "Language for the transcript (default: English)")
	flag.BoolVar(&options.Metadata, "metadata", false, "Output video metadata")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("Error: No URL provided.")
	}

	url := flag.Arg(0)
	ret, err = o.Grab(url, options)
	return
}
