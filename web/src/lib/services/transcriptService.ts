import { languageStore } from '$lib/store/language-store';
import { get } from 'svelte/store';

export interface TranscriptResponse {
  transcript: string;
  title: string;
}

function decodeHtmlEntities(text: string): string {
  const textarea = document.createElement('textarea');
  textarea.innerHTML = text;
  return textarea.value;
}

export async function getTranscript(url: string): Promise<TranscriptResponse> {
  try {
    const originalLanguage = get(languageStore);
    console.log('\n=== YouTube Transcript Service Start ===');
    console.log('1. Request details:', {
      url,
      endpoint: '/api/youtube/transcript',
      method: 'POST',
      isYouTubeURL: url.includes('youtube.com') || url.includes('youtu.be'),
      originalLanguage
    });

    const response = await fetch('/api/youtube/transcript', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        url,
        language: originalLanguage // Pass original language to server
      })
    });

    console.log('2. Server response:', {
      status: response.status,
      ok: response.ok,
      type: response.type,
      originalLanguage,
      currentLanguage: get(languageStore)
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || `HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    if (data.error) {
      throw new Error(data.error);
    }

    // Decode HTML entities in transcript
    data.transcript = decodeHtmlEntities(data.transcript);

    // Ensure language is preserved
    if (get(languageStore) !== originalLanguage) {
      console.log('3a. Restoring original language:', originalLanguage);
      languageStore.set(originalLanguage);
    }

    console.log('3b. Processed transcript:', {
      status: response.status,
      transcriptLength: data.transcript.length,
      firstChars: data.transcript.substring(0, 100),
      hasError: !!data.error,
      videoId: data.title,
      originalLanguage,
      currentLanguage: get(languageStore)
    });

    return data;
  } catch (error) {
    console.error('Transcript fetch error:', error);
    throw error instanceof Error ? error : new Error('Failed to fetch transcript');
  }
}
