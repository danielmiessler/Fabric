# Gemini Text-to-Speech (TTS) Guide

Fabric supports Google Gemini's text-to-speech (TTS) capabilities, allowing you to convert text into high-quality audio using various AI-generated voices.

## Overview

The Gemini TTS feature in Fabric allows you to:

- Convert text input into audio using Google's Gemini TTS models
- Choose from 30+ different AI voices with varying characteristics
- Generate high-quality WAV audio files
- Integrate TTS generation into your existing Fabric workflows

## Usage

### Basic TTS Generation

To generate audio from text using TTS:

```bash
# Basic TTS with default voice (Kore)
echo "Hello, this is a test of Gemini TTS" | fabric -m gemini-2.5-flash-preview-tts -o output.wav

# Using a specific voice
echo "Hello, this is a test with the Charon voice" | fabric -m gemini-2.5-flash-preview-tts --voice Charon -o output.wav

# Using TTS with a pattern
fabric -p summarize --voice Puck -m gemini-2.5-flash-preview-tts -o summary.wav < document.txt
```

### Voice Selection

Use the `--voice` flag to specify which voice to use for TTS generation:

```bash
fabric -m gemini-2.5-flash-preview-tts --voice Zephyr -o output.wav "Your text here"
```

If no voice is specified, the default voice "Kore" will be used.

## Available Voices

Gemini TTS supports 30+ different voices, each with unique characteristics:

### Popular Voices

- **Kore** - Firm and confident (default)
- **Charon** - Informative and clear
- **Puck** - Upbeat and energetic
- **Zephyr** - Bright and cheerful
- **Leda** - Youthful and energetic
- **Aoede** - Breezy and natural

### Complete Voice List

- Kore, Charon, Puck, Fenrir, Aoede, Leda, Orus, Zephyr
- Autonoe, Callirhoe, Despina, Erinome, Gacrux, Laomedeia
- Pulcherrima, Sulafat, Vindemiatrix, Achernar, Achird
- Algenib, Algieba, Alnilam, Enceladus, Iapetus, Rasalgethi
- Sadachbia, Zubenelgenubi, Vega, Capella, Lyra

### Listing Available Voices

To see all available voices with descriptions:

```bash
# List all voices with characteristics
fabric --list-gemini-voices

# List voice names only (for shell completion)
fabric --list-gemini-voices --shell-complete-list
```

## Rate Limits

Google Gemini TTS has usage quotas that vary by plan:

### Free Tier

- **15 requests per day** per project per TTS model
- Quota resets daily
- Applies to all TTS models (e.g., `gemini-2.5-flash-preview-tts`)

### Rate Limit Errors

If you exceed your quota, you'll see an error like:

```text
Error 429: You exceeded your current quota, please check your plan and billing details
```

**Solutions:**

- Wait for daily quota reset (typically at midnight UTC)
- Upgrade to a paid plan for higher limits
- Use TTS generation strategically for important content

For current rate limits and pricing, visit: <https://ai.google.dev/gemini-api/docs/rate-limits>

## Configuration

### Command Line Options

- `--voice <voice_name>` - Specify the TTS voice to use
- `-o <filename.wav>` - Output audio file (required for TTS models)
- `-m <tts_model>` - Specify a TTS-capable model (e.g., `gemini-2.5-flash-preview-tts`)

### YAML Configuration

You can also set a default voice in your Fabric configuration file (`~/.config/fabric/config.yaml`):

```yaml
voice: "Charon"  # Set your preferred default voice
```

## Requirements

- Valid Google Gemini API key configured in Fabric
- TTS-capable Gemini model (models containing "tts" in the name)
- Audio output must be specified with `-o filename.wav`

## Troubleshooting

### Common Issues

#### Error: "TTS model requires audio output"

- Solution: Always specify an output file with `-o filename.wav` when using TTS models

#### Error: "Invalid voice 'X'"

- Solution: Check that the voice name is spelled correctly and matches one of the supported voices listed above

#### Error: "TTS generation failed"

- Solution: Verify your Gemini API key is valid and you have sufficient quota

### Getting Help

For additional help with TTS features:

```bash
fabric --help
```

## Technical Details

- **Audio Format**: WAV files with 24kHz sample rate, 16-bit depth, mono channel
- **Language Support**: Automatic language detection for 24+ languages
- **Model Requirements**: Models must contain "tts", "preview-tts", or "text-to-speech" in the name
- **Voice Selection**: Uses Google's PrebuiltVoiceConfig system for consistent voice quality

---

For more information about Fabric, visit the [main documentation](../README.md).
