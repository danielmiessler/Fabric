# Using Speech-To-Text (STT) with Fabric

Fabric supports speech-to-text transcription of audio and video files using OpenAI's transcription models. This feature allows you to convert spoken content into text that can then be processed through Fabric's patterns.

## Overview

The STT feature integrates OpenAI's Whisper and GPT-4o transcription models to convert audio/video files into text. The transcribed text is automatically passed as input to your chosen pattern or chat session.

## Requirements

- OpenAI API key configured in Fabric
- For files larger than 25MB: `ffmpeg` installed on your system
- Supported audio/video formats: `.mp3`, `.mp4`, `.mpeg`, `.mpga`, `.m4a`, `.wav`, `.webm`

## Basic Usage

### Simple Transcription

To transcribe an audio file and send the result to a pattern:

```bash
fabric --transcribe-file /path/to/audio.mp3 --transcribe-model whisper-1 --pattern summarize
```

### Transcription Only

To just transcribe a file without applying a pattern:

```bash
fabric --transcribe-file /path/to/audio.mp3 --transcribe-model whisper-1
```

## Command Line Flags

### Required Flags

- `--transcribe-file`: Path to the audio or video file to transcribe
- `--transcribe-model`: Model to use for transcription (required when using transcription)

### Optional Flags

- `--split-media-file`: Automatically split files larger than 25MB into chunks using ffmpeg

## Available Models

You can list all available transcription models with:

```bash
fabric --list-transcription-models
```

Currently supported models:

- `whisper-1`: OpenAI's Whisper model
- `gpt-4o-mini-transcribe`: GPT-4o Mini transcription model
- `gpt-4o-transcribe`: GPT-4o transcription model

## File Size Handling

### Files Under 25MB

Files under the 25MB limit are processed directly without any special handling.

### Files Over 25MB

For files exceeding OpenAI's 25MB limit, you have two options:

1. **Manual handling**: The command will fail with an error message suggesting to use `--split-media-file`
2. **Automatic splitting**: Use the `--split-media-file` flag to automatically split the file into chunks

```bash
fabric --transcribe-file large_recording.mp4 --transcribe-model whisper-1 --split-media-file --pattern summarize
```

When splitting is enabled:

- Fabric uses `ffmpeg` to split the file into 10-minute segments initially
- If segments are still too large, it reduces the segment time by half repeatedly
- All segments are transcribed and the results are concatenated
- Temporary files are automatically cleaned up after processing

## Integration with Patterns

The transcribed text is seamlessly integrated into Fabric's workflow:

1. File is transcribed using the specified model
2. Transcribed text becomes the input message
3. Text is sent to the specified pattern or chat session

### Example Workflows

**Meeting transcription and summarization:**

```bash
fabric --transcribe-file meeting.mp4 --transcribe-model gpt-4o-transcribe --pattern summarize
```

**Interview analysis:**

```bash
fabric --transcribe-file interview.mp3 --transcribe-model whisper-1 --pattern extract_insights
```

**Large video file processing:**

```bash
fabric --transcribe-file presentation.mp4 --transcribe-model gpt-4o-transcribe --split-media-file --pattern create_summary
```

## Error Handling

Common error scenarios:

- **Unsupported format**: Only the listed audio/video formats are supported
- **File too large**: Use `--split-media-file` for files over 25MB
- **Missing ffmpeg**: Install ffmpeg for automatic file splitting
- **Invalid model**: Use `--list-transcription-models` to see available models
- **Missing model**: The `--transcribe-model` flag is required when using `--transcribe-file`

## Technical Details

### Implementation

- Transcription is handled in `internal/cli/transcribe.go:14`
- OpenAI-specific implementation in `internal/plugins/ai/openai/openai_audio.go:41`
- File splitting uses ffmpeg with configurable segment duration
- Supports any vendor that implements the `transcriber` interface

### Processing Pipeline

1. CLI validates file format and size
2. If file > 25MB and splitting enabled, file is split using ffmpeg
3. Each file/segment is sent to OpenAI's transcription API
4. Results are concatenated with spaces between segments
5. Transcribed text is passed as input to the main Fabric pipeline

### Vendor Support

Currently, only OpenAI is supported for transcription, but the interface allows for future expansion to other vendors that provide transcription capabilities.
