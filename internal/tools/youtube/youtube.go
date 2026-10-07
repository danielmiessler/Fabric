// Package youtube provides YouTube video transcript and comment extraction functionality.
//
// Requirements:
// - yt-dlp: Required for transcript extraction (must be installed separately)
// - ffmpeg and tesseract: Required, with yt-dlp, for visual extraction
// - YouTube API key: Optional, only needed for comments, metadata, duration, and playlists
//
// The implementation uses yt-dlp for transcript extraction and the YouTube API
// for comments/metadata.
package youtube

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
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
	// A cue sequence number, or a timestamp such as "01:02" or "00:00:01.234".
	timestampRegex = regexp.MustCompile(`^\d+$|^\d{1,2}:\d{2}(:\d{2})?(\.\d{3})?$`)
	// Language-tagged VTT file names such as .en.vtt or .pt-BR.vtt.
	languageFileRegex = regexp.MustCompile(`\.[a-z]{2}(-[A-Z]{2})?\.vtt$`)
	// Captures the video ID from watch, live, shorts, embed, v, and youtu.be URLs.
	videoPatternRegex = regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:live\/|[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|(?:s(?:horts)\/)|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]*)`)
	playlistPatternRegex = regexp.MustCompile(`[?&]list=([a-zA-Z0-9_-]+)`)
	// VTT formatting tags such as <c.colorE5E5E5> and </c>.
	vttTagRegex = regexp.MustCompile(`<[^>]*>`)
	// ISO 8601 duration from the Data API, such as PT1H2M3S.
	durationRegex = regexp.MustCompile(`(?i)PT(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)
}

func NewYouTube() (ret *YouTube) {

	label := "YouTube"
	ret = &YouTube{}

	ret.PluginBase = &plugins.PluginBase{
		Name:             i18n.T("youtube_label"),
		SetupDescription: i18n.T("youtube_setup_description") + " " + i18n.T("optional_marker"),
		EnvNamePrefix:    plugins.BuildEnvVariablePrefix(label),
	}

	ret.ApiKey = ret.AddSetupQuestion("API key", false)

	return
}

type YouTube struct {
	*plugins.PluginBase
	ApiKey *plugins.SetupQuestion
	// YtDlpArgs holds the yt-dlp arguments from the --yt-dlp-args flag or the
	// config file. GrabTranscript and GrabTranscriptWithTimestamps use them.
	YtDlpArgs string

	normalizeRegex *regexp.Regexp
	service        *youtube.Service
}

func (o *YouTube) initService() (err error) {
	if o.service == nil {
		if o.ApiKey.Value == "" {
			err = errors.New(i18n.T("youtube_api_key_required"))
			return
		}
		o.normalizeRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
		ctx := context.Background()
		o.service, err = youtube.NewService(ctx, option.WithAPIKey(o.ApiKey.Value))
	}
	return
}

func (o *YouTube) GetVideoOrPlaylistId(url string) (videoId string, playlistId string, err error) {
	videoMatch := videoPatternRegex.FindStringSubmatch(url)
	if len(videoMatch) > 1 {
		videoId = videoMatch[1]
	}

	playlistMatch := playlistPatternRegex.FindStringSubmatch(url)
	if len(playlistMatch) > 1 {
		playlistId = playlistMatch[1]
	}

	if videoId == "" && playlistId == "" {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_url"), url))
	}
	return
}

// extractAndValidateVideoId returns the video ID in url. It returns an error
// when url is invalid or names only a playlist.
func (o *YouTube) extractAndValidateVideoId(url string) (videoId string, err error) {
	var playlistId string
	if videoId, playlistId, err = o.GetVideoOrPlaylistId(url); err != nil {
		return "", err
	}
	if videoId == "" && playlistId != "" {
		return "", errors.New(i18n.T("youtube_url_is_playlist_not_video"))
	}
	if videoId == "" {
		return "", errors.New(i18n.T("youtube_no_video_id_found"))
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

// GrabTranscript retrieves the transcript for the specified video ID using yt-dlp
// and the arguments in the YtDlpArgs field.
// The language parameter specifies the preferred subtitle language code (e.g., "en", "es").
// It returns the transcript text or an error if the transcript cannot be retrieved.
func (o *YouTube) GrabTranscript(videoId string, language string) (ret string, err error) {
	return o.GrabTranscriptWithArgs(videoId, language, o.YtDlpArgs)
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
// video ID using yt-dlp and the arguments in the YtDlpArgs field. The language parameter
// specifies the preferred subtitle language code.
// Each line in the returned transcript is prefixed with a timestamp in [HH:MM:SS] format.
// It returns the timestamped transcript text or an error if the transcript cannot be retrieved.
func (o *YouTube) GrabTranscriptWithTimestamps(videoId string, language string) (ret string, err error) {
	return o.GrabTranscriptWithTimestampsWithArgs(videoId, language, o.YtDlpArgs)
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
		return errors.New(i18n.T("youtube_ytdlp_stderr_error"))
	}
	return nil
}

// ytDlpLangArgs returns the built-in --sub-langs flag for language. yt-dlp
// does not replace a repeated --sub-langs flag. It adds the values together.
// So the function returns nil when the user arguments set the languages.
func ytDlpLangArgs(language string, userArgs []string) []string {
	userSetsLangs := slices.ContainsFunc(userArgs, func(a string) bool {
		return strings.HasPrefix(a, "--sub-lang")
	})
	if language == "" || userSetsLangs {
		return nil
	}
	langMatch := language
	if len(language) > 2 {
		langMatch = language[:2]
	}
	langOpts := language + "," + langMatch + ".*"
	if langMatch != language {
		langOpts += "," + langMatch
	}
	return []string{"--sub-langs", langOpts}
}

// tryMethodYtDlpInternal downloads the subtitles for videoId with yt-dlp, then
// applies processVTTFileFunc to the VTT file.
func (o *YouTube) tryMethodYtDlpInternal(videoId string, language string, additionalArgs string, processVTTFileFunc func(filename string) (string, error)) (ret string, err error) {
	if _, err = exec.LookPath("yt-dlp"); err != nil {
		err = errors.New(i18n.T("youtube_ytdlp_not_found"))
		return
	}

	// Make a new temporary folder with a random name and mode 0700. Another
	// local user cannot know the name, so cannot put a file in it first.
	tempDir, err := os.MkdirTemp("", "fabric-youtube-")
	if err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_failed_create_temp_dir"), err))
		return
	}
	defer os.RemoveAll(tempDir)

	videoURL := "https://www.youtube.com/watch?v=" + videoId
	outputPath := filepath.Join(tempDir, "%(title)s.%(ext)s")

	baseArgs := []string{
		"--write-auto-subs",
		"--skip-download",
		"--sub-format", "vtt",
		"-o", outputPath,
	}

	var userArgs []string
	if additionalArgs != "" {
		if userArgs, err = shellquote.Split(additionalArgs); err != nil {
			return "", fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_ytdlp_arguments"), err))
		}
	}
	langArgs := ytDlpLangArgs(language, userArgs)

	for retry := 1; retry >= 0; retry-- {
		var ytOutput []byte
		args := slices.Concat(baseArgs, langArgs, userArgs, []string{videoURL})
		cmd := exec.Command("yt-dlp", args...)
		debuglog.Debug(debuglog.Trace, "yt-dlp %+v\n", cmd.Args)
		ytOutput, err = cmd.CombinedOutput()
		ytReader := bytes.NewReader(ytOutput)
		if err = detectError(ytReader); err == nil {
			break
		}
		// Retry without the built-in language filter.
		langArgs = nil
	}
	if err != nil {
		return
	}
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

	lines := strings.Split(string(content), "\n")
	var textBuilder strings.Builder
	seenSegments := make(map[string]struct{})

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line == "WEBVTT" || strings.Contains(line, "-->") ||
			strings.HasPrefix(line, "NOTE") || strings.HasPrefix(line, "STYLE") ||
			strings.HasPrefix(line, "Kind:") || strings.HasPrefix(line, "Language:") ||
			isTimeStamp(line) {
			continue
		}
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
		err = errors.New(i18n.T("youtube_no_transcript_content"))
	}
	return
}

func (o *YouTube) readAndFormatVTTWithTimestamps(filename string) (ret string, err error) {
	var content []byte
	if content, err = os.ReadFile(filename); err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	var textBuilder strings.Builder
	var currentTimestamp string
	// Map each line to the timestamp where it last appeared. A repeat within
	// TimeGapForRepeats seconds is a VTT duplicate. A later repeat is real
	// content, such as a chorus.
	seenSegments := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || line == "WEBVTT" || strings.HasPrefix(line, "NOTE") ||
			strings.HasPrefix(line, "STYLE") || strings.HasPrefix(line, "Kind:") ||
			strings.HasPrefix(line, "Language:") {
			continue
		}

		if strings.Contains(line, "-->") {
			parts := strings.Split(line, " --> ")
			if len(parts) >= 1 {
				currentTimestamp = formatVTTTimestamp(parts[0])
			}
			continue
		}

		// Skip cue sequence numbers.
		if isTimeStamp(line) && !strings.Contains(line, ":") {
			continue
		}

		if line != "" {
			cleanText := removeVTTTags(line)
			if cleanText != "" && currentTimestamp != "" {
				shouldInclude := true
				if lastTimestamp, exists := seenSegments[cleanText]; exists {
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
		err = errors.New(i18n.T("youtube_no_transcript_content"))
	}
	return
}

func formatVTTTimestamp(vttTime string) string {
	// Drops the milliseconds, so "00:00:01.234" becomes "00:00:01".
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
	return vttTagRegex.ReplaceAllString(s, "")
}

// shouldIncludeRepeat reports whether at least TimeGapForRepeats seconds
// separate the two timestamps.
func shouldIncludeRepeat(lastTimestamp, currentTimestamp string) bool {
	lastSeconds, err1 := parseTimestampToSeconds(lastTimestamp)
	currentSeconds, err2 := parseTimestampToSeconds(currentTimestamp)

	if err1 != nil || err2 != nil {
		// Keep the line when a timestamp does not parse.
		return true
	}

	timeDiffSeconds := currentSeconds - lastSeconds
	return timeDiffSeconds >= TimeGapForRepeats
}

// parseTimestampToSeconds converts an HH:MM:SS or MM:SS timestamp to seconds.
func parseTimestampToSeconds(timestamp string) (int, error) {
	parts := strings.Split(timestamp, ":")
	if len(parts) < 2 || len(parts) > 3 {
		return 0, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_invalid_timestamp_format"), timestamp))
	}

	var hours, minutes, seconds int
	var err error

	if len(parts) == 3 {
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
		return 0, errors.New(i18n.T("youtube_empty_seconds_string"))
	}

	intPart := secondsStr
	if idx := strings.Index(secondsStr, "."); idx != -1 {
		if idx == 0 {
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
		log.Printf(i18n.T("youtube_failed_fetch_comments"), err)
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

	if options.Visual {
		if ret.VisualText, err = o.GrabVisual(videoId, options.Lang, options.YtDlpArgs, options.VisualSensitivity, options.VisualFps); err != nil {
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

	if err = writer.Write([]string{"VideoID", "Title"}); err != nil {
		return
	}

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

	fmt.Printf("%s\n", fmt.Sprintf(i18n.T("youtube_playlist_saved_to"), filename))
	return
}

func (o *YouTube) FetchAndPrintPlaylist(playlistID string) (err error) {
	var videos []*VideoMeta
	if videos, err = o.FetchPlaylistVideos(playlistID); err != nil {
		err = fmt.Errorf("%s", fmt.Sprintf(i18n.T("error_fetching_playlist_videos"), err))
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf(i18n.T("youtube_playlist_header"), playlistID))
	fmt.Printf("%s\n", i18n.T("youtube_video_id_title_header"))
	for _, video := range videos {
		fmt.Printf("%s: %s\n", video.Id, video.Title)
	}
	return
}

func (o *YouTube) normalizeFileName(name string) string {
	return o.normalizeRegex.ReplaceAllString(name, "_")

}

// findVTTFilesWithFallback returns one VTT file from dir. When requestedLanguage
// is set, it prefers a file tagged with that language, then any language-tagged
// file. Otherwise it returns the first file found.
func (o *YouTube) findVTTFilesWithFallback(dir, requestedLanguage string) ([]string, error) {
	var vttFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Use only regular files. Do not follow a .vtt symlink, because it can
		// point to a file outside the folder.
		if info.Mode().IsRegular() && strings.HasSuffix(strings.ToLower(path), ".vtt") {
			vttFiles = append(vttFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("%s", fmt.Sprintf(i18n.T("youtube_failed_walk_directory"), err))
	}

	if len(vttFiles) == 0 {
		return nil, errors.New(i18n.T("youtube_no_vtt_files_found"))
	}

	if requestedLanguage == "" {
		return []string{vttFiles[0]}, nil
	}

	for _, file := range vttFiles {
		if strings.Contains(file, "."+requestedLanguage+".vtt") {
			return []string{file}, nil
		}
	}

	for _, file := range vttFiles {
		if languageFileRegex.MatchString(file) {
			return []string{file}, nil
		}
	}

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
	Visual                   bool
	VisualSensitivity        float64
	VisualFps                int
	Comments                 bool
	Lang                     string
	Metadata                 bool
	YtDlpArgs                string
}

type VideoInfo struct {
	Transcript string         `json:"transcript"`
	VisualText string         `json:"visualText,omitempty"`
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
	flag.BoolVar(&options.Visual, "visual", false, i18n.T("youtube_extract_visual_data_help"))
	flag.Float64Var(&options.VisualSensitivity, "visual-sensitivity", 0.4, i18n.T("youtube_visual_sensitivity_help"))
	flag.IntVar(&options.VisualFps, "visual-fps", 0, i18n.T("youtube_visual_fps_help"))
	flag.BoolVar(&options.Comments, "comments", false, "Output the comments on the video")
	flag.StringVar(&options.Lang, "lang", "en", "Language for the transcript (default: English)")
	flag.BoolVar(&options.Metadata, "metadata", false, "Output video metadata")
	flag.StringVar(&options.YtDlpArgs, "yt-dlp-args", "", i18n.T("additional_yt_dlp_args"))
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalf("%s", i18n.T("youtube_no_url_provided"))
	}

	url := flag.Arg(0)
	ret, err = o.Grab(url, options)
	return
}

// GrabVisual retrieves visual data from the video by extracting frames via FFmpeg and OCR parsing them via Tesseract.
func (o *YouTube) GrabVisual(videoId string, language string, additionalArgs string, sensitivity float64, fps int) (string, error) {
	if _, err := exec.LookPath("yt-dlp"); err != nil {
		return "", errors.New(i18n.T("youtube_ytdlp_required_visual_extraction"))
	}
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return "", errors.New(i18n.T("youtube_ffmpeg_required_visual_extraction"))
	}
	if _, err := exec.LookPath("tesseract"); err != nil {
		return "", errors.New(i18n.T("youtube_tesseract_required_visual_extraction"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	tempDir, err := os.MkdirTemp("", "fabric-vfabric-"+videoId+"-*")
	if err != nil {
		return "", fmt.Errorf(i18n.T("youtube_failed_create_temp_dir"), err)
	}
	defer os.RemoveAll(tempDir)

	videoURL := "https://www.youtube.com/watch?v=" + videoId

	ytArgs := []string{"-f", "bv", "--get-url"}
	if additionalArgs != "" {
		parsed, parseErr := shellquote.Split(additionalArgs)
		if parseErr != nil {
			return "", fmt.Errorf(i18n.T("youtube_invalid_ytdlp_arguments"), parseErr)
		}
		ytArgs = append(ytArgs, parsed...)
	}
	ytArgs = append(ytArgs, "--", videoURL)

	cmdUrl := exec.CommandContext(ctx, "yt-dlp", ytArgs...)
	urlBytes, err := cmdUrl.Output()
	if err != nil {
		return "", fmt.Errorf(i18n.T("youtube_failed_get_stream_url"), err)
	}

	streamUrls := strings.Split(strings.TrimSpace(string(urlBytes)), "\n")
	var streamUrl string
	for _, u := range streamUrls {
		if strings.HasPrefix(u, "http") {
			streamUrl = strings.TrimSpace(u)
			break
		}
	}
	if streamUrl == "" {
		return "", errors.New(i18n.T("youtube_failed_parse_http_stream_url"))
	}

	var filter string
	if fps > 0 {
		filter = fmt.Sprintf("fps=%d", fps)
	} else {
		filter = fmt.Sprintf("select='gt(scene,%f)'", sensitivity)
	}

	framePattern := filepath.Join(tempDir, "frame_%04d.jpg")
	cmdFfmpeg := exec.CommandContext(ctx, "ffmpeg", "-i", streamUrl, "-vf", filter, "-fps_mode", "vfr", framePattern)
	if out, err := cmdFfmpeg.CombinedOutput(); err != nil {
		return "", fmt.Errorf(i18n.T("youtube_ffmpeg_frame_extraction_failed"), err, string(out))
	}

	files, err := filepath.Glob(filepath.Join(tempDir, "frame_*.jpg"))
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	results := make([]string, len(files))
	var errs []error
	var errMut sync.Mutex
	sem := make(chan struct{}, runtime.NumCPU())

	for i, file := range files {
		wg.Add(1)
		sem <- struct{}{}
		go func(idx int, f string) {
			defer wg.Done()
			defer func() { <-sem }()

			cmdOcr := exec.CommandContext(ctx, "tesseract", f, "-", "stdout")
			var stdoutBuf, stderrBuf bytes.Buffer
			cmdOcr.Stdout = &stdoutBuf
			cmdOcr.Stderr = &stderrBuf
			ocrErr := cmdOcr.Run()
			if ocrErr != nil {
				errMut.Lock()
				errs = append(errs, fmt.Errorf(i18n.T("youtube_tesseract_frame_failed"), idx, ocrErr, stderrBuf.String()))
				errMut.Unlock()
				return
			}

			text := strings.TrimSpace(stdoutBuf.String())
			if text != "" && len(text) > 10 {
				results[idx] = text
			}
		}(i, file)
	}
	wg.Wait()

	if len(errs) > 0 {
		return "", errs[0]
	}

	var sb strings.Builder
	for i, text := range results {
		if text != "" {
			secs := i
			hours := secs / 3600
			mins := (secs % 3600) / 60
			sec := secs % 60
			sb.WriteString(fmt.Sprintf("\n%02d:%02d:%02d.000 --> %02d:%02d:%02d.999\n", hours, mins, sec, hours, mins, sec))
			sb.WriteString(i18n.T("youtube_visual_frame_cue"))
			sb.WriteString("\n")
			sb.WriteString(text)
			sb.WriteString("\n")
		}
	}

	ret := sb.String()
	if ret == "" {
		return "", errors.New(i18n.T("youtube_no_clear_text_visual_frames"))
	}
	return ret, nil
}
