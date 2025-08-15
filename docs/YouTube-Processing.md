# YouTube Processing with Fabric

Fabric provides powerful YouTube video processing capabilities that allow you to extract transcripts, comments, and metadata from YouTube videos and playlists. This guide covers all the available options and common use cases.

## Prerequisites

- **yt-dlp**: Required for transcript extraction. Install on MacOS with:

  ```bash
  brew install yt-dlp
  ```

  Or use the package manager of your choice for your operating system.

  See the [yt-dlp wiki page](https://github.com/yt-dlp/yt-dlp/wiki/Installation) for your specific installation instructions.

- **YouTube API Key** (optional): Only needed for comments and metadata extraction. Configure with:

  ```bash
  fabric --setup
  ```

## Basic Usage

### Extract Transcript

Extract a video transcript and process it with a pattern:

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern summarize
```

### Extract Transcript with Timestamps

Get transcript with timestamps preserved:

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --transcript-with-timestamps --pattern extract_wisdom
```

### Extract Comments

Get video comments (requires YouTube API key):

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --comments --pattern analyze_claims
```

### Extract Metadata

Get video metadata as JSON:

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --metadata
```

## Advanced Options

### Custom yt-dlp Arguments

Pass additional arguments to yt-dlp for advanced functionality. **User-provided arguments take precedence** over built-in fabric arguments, giving you full control:

```bash
# Use browser cookies for age-restricted or private videos
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--cookies-from-browser brave"

# Override language selection (takes precedence over -g flag)
fabric -g en -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--sub-langs es,fr"

# Use specific format
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--format best"

# Handle rate limiting (slow down requests)
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--sleep-requests 1"

# Multiple arguments (use quotes)
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--cookies-from-browser firefox --write-info-json"

# Combine rate limiting with authentication
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--cookies-from-browser brave --sleep-requests 1"

# Override subtitle format (takes precedence over built-in --sub-format vtt)
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --yt-dlp-args="--sub-format srt"
```

#### Argument Precedence

Fabric constructs the yt-dlp command in this order:

1. **Built-in base arguments** (`--write-auto-subs`, `--skip-download`, etc.)
2. **Language selection** (from `-g` flag): `--sub-langs LANGUAGE`
3. **User arguments** (from `--yt-dlp-args`): **These override any conflicting built-in arguments**
4. **Video URL**

This means you can override any built-in behavior by specifying it in `--yt-dlp-args`.

### Playlist Processing

Process entire playlists:

```bash
# Process all videos in a playlist
fabric -y "https://www.youtube.com/playlist?list=PLAYLIST_ID" --playlist --pattern summarize

# Save playlist videos to CSV
fabric -y "https://www.youtube.com/playlist?list=PLAYLIST_ID" --playlist -o playlist.csv
```

### Language Support

Specify transcript language:

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" -g es --pattern translate
```

## Combining Options

You can combine multiple YouTube processing options:

```bash
# Get transcript, comments, and metadata
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" \
  --transcript \
  --comments \
  --metadata \
  --pattern comprehensive_analysis
```

## Output Options

### Save to File

```bash
# Save output to file
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern summarize -o summary.md

# Save entire session including input
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern summarize --output-session -o full_session.md
```

### Stream Output

Get real-time streaming output:

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern summarize --stream
```

## Common Use Cases

### Content Analysis

```bash
# Analyze video content for key insights
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern extract_wisdom

# Check claims made in the video
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern analyze_claims
```

### Educational Content

```bash
# Create study notes from educational videos
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern create_study_notes

# Extract key concepts and definitions
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern extract_concepts
```

### Meeting/Conference Processing

```bash
# Summarize conference talks with timestamps
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" \
  --transcript-with-timestamps \
  --pattern meeting_summary

# Extract action items from recorded meetings
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern extract_action_items
```

### Content Creation

```bash
# Create social media posts from video content
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern create_social_posts

# Generate blog post from video transcript
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" --pattern write_blog_post
```

## Troubleshooting

### Common Issues

1. **"yt-dlp not found"**: Install yt-dlp using pip or your package manager
2. **Age-restricted videos**: Use `--yt-dlp-args="--cookies-from-browser BROWSER"`
3. **No subtitles available**: Some videos don't have auto-generated subtitles
4. **API rate limits**: YouTube API has daily quotas for comments/metadata
5. **HTTP 429 errors**: YouTube is rate limiting subtitle requests

### Error Messages

- **"YouTube is not configured"**: Run `fabric --setup` to configure YouTube API
- **"yt-dlp failed"**: Check video URL and try with `--yt-dlp-args` for authentication
- **"No transcript content found"**: Video may not have subtitles available
- **"HTTP Error 429: Too Many Requests"**: YouTube rate limit exceeded. This is increasingly common. Solutions:
  - **Wait 10-30 minutes and try again** (most effective)
  - Use longer sleep: `--yt-dlp-args="--sleep-requests 5"`
  - Try with browser cookies: `--yt-dlp-args="--cookies-from-browser brave --sleep-requests 5"`
  - **Try a different video** - some videos are less restricted
  - **Use a VPN** - different IP address may help
  - **Try without language specification** - let yt-dlp choose any available language
  - **Try English instead** - `fabric -g en` (English subtitles may be less rate-limited)

### Language Fallback Behavior

When you specify a language (e.g., `-g es` for Spanish) but that language isn't available or fails to download:

1. **Automatic fallback**: Fabric automatically retries without language specification
2. **Smart file detection**: If the fallback downloads a different language (e.g., English), Fabric will automatically detect and use it
3. **No manual intervention needed**: The process is transparent to the user

```bash
# Even if Spanish isn't available, this will work with whatever language yt-dlp finds
fabric -g es -y "https://youtube.com/watch?v=VIDEO_ID" --pattern summarize
```

## Configuration

### YAML Configuration

You can set default yt-dlp arguments in your config file (`~/.config/fabric/config.yaml`):

```yaml
ytDlpArgs: "--cookies-from-browser brave --write-info-json"
```

### Environment Variables

Set up your YouTube API key:

```bash
export FABRIC_YOUTUBE_API_KEY="your_api_key_here"
```

## Tips and Best Practices

1. **Use specific patterns**: Choose patterns that match your use case for better results
2. **Combine with other tools**: Pipe output to other commands or save to files for further processing
3. **Batch processing**: Use playlists to process multiple videos efficiently
4. **Authentication**: Use browser cookies for accessing private or age-restricted content
5. **Language support**: Specify language codes for better transcript accuracy
6. **Rate limiting**: If you encounter 429 errors, use `--sleep-requests 1` to slow down requests
7. **Persistent settings**: Set common yt-dlp args in your config file to avoid repeating them
8. **Argument precedence**: Use `--yt-dlp-args` to override any built-in behavior when needed
9. **Testing**: Use `yt-dlp --list-subs URL` to see available subtitle languages before processing

## Examples

### Quick Video Summary

```bash
fabric -y "https://www.youtube.com/watch?v=dQw4w9WgXcQ" --pattern summarize --stream
```

### Detailed Analysis with Authentication

```bash
fabric -y "https://www.youtube.com/watch?v=VIDEO_ID" \
  --yt-dlp-args="--cookies-from-browser chrome" \
  --transcript-with-timestamps \
  --comments \
  --pattern comprehensive_analysis \
  -o analysis.md
```

### Playlist Processing

```bash
fabric -y "https://www.youtube.com/playlist?list=PLrAXtmRdnEQy6nuLvVUxpDnx4C0823vBN" \
  --playlist \
  --pattern extract_wisdom \
  -o playlist_wisdom.md
```

### Override Built-in Language Selection

```bash
# Built-in language selection (-g es) is overridden by user args
fabric -g es -y "https://www.youtube.com/watch?v=VIDEO_ID" \
  --yt-dlp-args="--sub-langs fr,de,en" \
  --pattern translate
```

For more patterns and advanced usage, see the main [Fabric documentation](../README.md).
