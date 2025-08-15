# Voice Samples

This directory contains sample audio files demonstrating different Gemini TTS voices.

## Sample Files

Each voice sample says "The quick brown fox jumped over the lazy dog" to demonstrate the voice characteristics:

- **Kore.wav** - Firm and confident (default voice)
- **Charon.wav** - Informative and clear
- **Vega.wav** - Smooth and pleasant
- **Capella.wav** - Warm and welcoming
- **Achird.wav** - Friendly and approachable
- **Lyra.wav** - Melodic and expressive

## Generating Samples

To generate these samples, use the following commands:

```bash
# Generate each voice sample
echo "The quick brown fox jumped over the lazy dog" | fabric -m gemini-2.5-flash-preview-tts --voice Kore -o docs/voices/Kore.wav
echo "The quick brown fox jumped over the lazy dog" | fabric -m gemini-2.5-flash-preview-tts --voice Charon -o docs/voices/Charon.wav
echo "The quick brown fox jumped over the lazy dog" | fabric -m gemini-2.5-flash-preview-tts --voice Vega -o docs/voices/Vega.wav
echo "The quick brown fox jumped over the lazy dog" | fabric -m gemini-2.5-flash-preview-tts --voice Capella -o docs/voices/Capella.wav
echo "The quick brown fox jumped over the lazy dog" | fabric -m gemini-2.5-flash-preview-tts --voice Achird -o docs/voices/Achird.wav
echo "The quick brown fox jumped over the lazy dog" | fabric -m gemini-2.5-flash-preview-tts --voice Lyra -o docs/voices/Lyra.wav
```

## Audio Format

- **Format**: WAV (uncompressed)
- **Sample Rate**: 24kHz
- **Bit Depth**: 16-bit
- **Channels**: Mono
- **Approximate Size**: ~500KB per sample
