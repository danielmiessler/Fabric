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
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
	debuglog "github.com/danielmiessler/fabric/internal/log"
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

	ret.ApiKey = ret.AddSetupQuestion("API key", false)

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
			err = fmt.Errorf("%s", i18n.T("youtube_api_key_required"))
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
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_url"), url))
	}
	return
}

// extractAndValidateVideoId extracts a video ID from the given URL and validates
// that the URL points to a video rather than a playlist-only resource.
// It returns an error if the URL is invalid or contains only playlist information.
func (o *YouTube) extractAndValidateVideoId(url string) (videoId string, err error) {
	var playlistId string
	if videoId, playlistId, err = o.GetVideoOrPlaylistId(url); err != nil {
		return "", err
	}
	if videoId == "" && playlistId != "" {
		return "", fmt.Errorf("%s", i18n.T("youtube_url_is_playlist_not_video"))
	}
	if videoId == "" {
		return "", fmt.Errorf("%s", i18n.T("youtube_no_video_id_found"))
	}
	return videoId, nil
}

func (o *YouTube) GrabTranscriptForUrl(url string, language string) (ret string, err error) {
	var videoId string
	if videoId, err = o.extractAndValidateVideoId(url); err != nil {
		return
	}
	return o.GrabTranscript(videoId, language)
}

// GrabTranscript retrieves the transcript for the specified video ID using yt-dlp.
// The language parameter specifies the preferred subtitle language code (e.g., "en", "es").
// It returns the transcript text or an error if the transcript cannot be retrieved.
func (o *YouTube) GrabTranscript(videoId string, language string) (ret string, err error) {
	return o.GrabTranscriptWithArgs(videoId, language, "")
}

// GrabTranscriptWithArgs retrieves the transcript for the specified video ID using yt-dlp
// with custom command-line arguments. The language parameter specifies the preferred subtitle
// language code. The additionalArgs parameter allows passing extra yt-dlp options like
// "--cookies-from-browser brave" for authentication.
// It returns the transcript text or an error if the transcript cannot be retrieved.
func (o *YouTube) GrabTranscriptWithArgs(videoId string, language string, additionalArgs string) (ret string, err error) {
	return o.tryMethodYtDlp(videoId, language, additionalArgs)
}

// GrabTranscriptWithTimestamps retrieves the transcript with timestamps for the specified
// video ID using yt-dlp. The language parameter specifies the preferred subtitle language code.
// Each line in the returned transcript is prefixed with a timestamp in [HH:MM:SS] format.
// It returns the timestamped transcript text or an error if the transcript cannot be retrieved.
func (o *YouTube) GrabTranscriptWithTimestamps(videoId string, language string) (ret string, err error) {
	return o.GrabTranscriptWithTimestampsWithArgs(videoId, language, "")
}

// GrabTranscriptWithTimestampsWithArgs retrieves the transcript with timestamps for the specified
// video ID using yt-dlp with custom command-line arguments. The language parameter specifies the
// preferred subtitle language code. The additionalArgs parameter allows passing extra yt-dlp options.
// Each line in the returned transcript is prefixed with a timestamp in [HH:MM:SS] format.
// It returns the timestamped transcript text or an error if the transcript cannot be retrieved.
func (o *YouTube) GrabTranscriptWithTimestampsWithArgs(videoId string, language string, additionalArgs string) (ret string, err error) {
	return o.tryMethodYtDlpWithTimestamps(videoId, language, additionalArgs)
}

func detectError(ytOutput io.Reader) error {
	scanner := bufio.NewScanner(ytOutput)
	for scanner.Scan() {
		curLine := scanner.Text()
		debuglog.Debug(debuglog.Trace, "%s\n", curLine)
		errorMessages := map[string]string{
			"429":                                 i18n.T("youtube_rate_limit_exceeded"),
			"Too Many Requests":                   i18n.T("youtube_rate_limit_exceeded"),
			"Sign in to confirm you're not a bot": i18n.T("youtube_auth_required_bot_detection"),
			"Use --cookies-from-browser":          i18n.T("youtube_auth_required_bot_detection"),
		}

		for key, message := range errorMessages {
			if strings.Contains(curLine, key) {
				return fmt.Errorf("%s", message)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%s", i18n.T("youtube_ytdlp_stderr_error"))
	}
	return nil
}

func noLangs(args []string) []string {
	var (
		i int
		v string
	)
	for i, v = range args {
		if strings.Contains(v, "--sub-langs") {
			break
		}
	}
	if i == 0 || i == len(args)-1 {
		return args
	}
	return append(args[0:i], args[i+2:]...)
}

// tryMethodYtDlpInternal is a helper function to reduce duplication between
// tryMethodYtDlp and tryMethodYtDlpWithTimestamps.
func (o *YouTube) tryMethodYtDlpInternal(videoId string, language string, additionalArgs string, processVTTFileFunc func(filename string) (string, error)) (ret string, err error) {
	// Check if yt-dlp is available
	if _, err = exec.LookPath("yt-dlp"); err != nil {
		err = fmt.Errorf("%s", i18n.T("youtube_ytdlp_not_found"))
		return
	}

	// Create a temporary directory for yt-dlp output (cross-platform)
	tempDir := filepath.Join(os.TempDir(), "fabric-youtube-"+videoId)
	if err = os.MkdirAll(tempDir, 0755); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_failed_create_temp_dir"), err))
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
		"-o", outputPath,
	}

	args := append([]string{}, baseArgs...)

	// Add built-in language selection first
	if language != "" {
		langMatch := language[:2]
		langOpts := language + "," + langMatch + ".*"
		if langMatch != language {
			langOpts += "," + langMatch
		}
		args = append(args, "--sub-langs", langOpts)
	}

	// Add user-provided arguments last so they take precedence
	if additionalArgs != "" {
		additionalArgsList, err := shellquote.Split(additionalArgs)
		if err != nil {
			return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_ytdlp_arguments"), err))
		}
		args = append(args, additionalArgsList...)
	}

	args = append(args, videoURL)

	for retry := 1; retry >= 0; retry-- {
		var ytOutput []byte
		cmd := exec.Command("yt-dlp", args...)
		debuglog.Debug(debuglog.Trace, "yt-dlp %+v\n", cmd.Args)
		ytOutput, err = cmd.CombinedOutput()
		ytReader := bytes.NewReader(ytOutput)
		if err = detectError(ytReader); err == nil {
			break
		}
		args = noLangs(args)
	}
	if err != nil {
		return
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
		err = fmt.Errorf("%s", i18n.T("youtube_no_transcript_content"))
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
		err = fmt.Errorf("%s", i18n.T("youtube_no_transcript_content"))
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
		return 0, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_timestamp_format"), timestamp))
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

func parseSeconds(secondsStr string) (int, error) {
	if secondsStr == "" {
		return 0, fmt.Errorf("%s", i18n.T("youtube_empty_seconds_string"))
	}

	// Extract integer part (before decimal point if present)
	intPart := secondsStr
	if idx := strings.Index(secondsStr, "."); idx != -1 {
		if idx == 0 {
			// Handle cases like ".5" -> treat as "0"
			intPart = "0"
		} else {
			intPart = secondsStr[:idx]
		}
	}

	seconds, err := strconv.Atoi(intPart)
	if err != nil {
		return 0, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_seconds_format"), secondsStr, err))
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
	if videoId, err = o.extractAndValidateVideoId(url); err != nil {
		return
	}
	return o.GrabDuration(videoId)
}

func (o *YouTube) GrabDuration(videoId string) (ret int, err error) {
	var videoResponse *youtube.VideoListResponse
	if videoResponse, err = o.service.Videos.List([]string{"contentDetails"}).Id(videoId).Do(); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_error_getting_video_details"), err))
		return
	}

	durationStr := videoResponse.Items[0].ContentDetails.Duration

	matches := durationRegex.FindStringSubmatch(durationStr)
	if len(matches) == 0 {
		return 0, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_duration_string"), durationStr))
	}

	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])
	seconds, _ := strconv.Atoi(matches[3])

	ret = hours*60 + minutes + seconds/60

	return
}

func (o *YouTube) Grab(url string, options *Options) (ret *VideoInfo, err error) {
	var videoId string
	if videoId, err = o.extractAndValidateVideoId(url); err != nil {
		return
	}

	ret = &VideoInfo{}

	if options.Metadata {
		if ret.Metadata, err = o.GrabMetadata(videoId); err != nil {
			err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_error_getting_metadata"), err))
			return
		}
	}

	if options.Duration {
		if ret.Duration, err = o.GrabDuration(videoId); err != nil {
			err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_error_parsing_duration"), err))
			return
		}

	}

	if options.Comments {
		if ret.Comments, err = o.GrabComments(videoId); err != nil {
			err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_error_getting_comments"), err))
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
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_fetching_playlist_videos"), err))
		return
	}

	if err = o.SaveVideosToCSV(filename, videos); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_error_saving_csv"), err))
		return
	}

	fmt.Println("Playlist saved to", filename)
	return
}

func (o *YouTube) FetchAndPrintPlaylist(playlistID string) (err error) {
	var videos []*VideoMeta
	if videos, err = o.FetchPlaylistVideos(playlistID); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_fetching_playlist_videos"), err))
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
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_failed_walk_directory"), err))
	}

	if len(vttFiles) == 0 {
		return nil, fmt.Errorf("%s", i18n.T("youtube_no_vtt_files_found"))
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
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_error_getting_metadata"), err))
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_no_video_found_with_id"), videoId))
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
