import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { YoutubeTranscript } from 'youtube-transcript';

const MAX_BODY_SIZE = 2048; // 2KB limit for a YouTube URL request

export const POST: RequestHandler = async ({ request }) => {
  try {
    const contentLength = request.headers.get('content-length');
    if (contentLength && parseInt(contentLength, 10) > MAX_BODY_SIZE) {
      return json({ error: 'Request body too large' }, { status: 413 });
    }
    const bodyText = await request.text();
    if (bodyText.length > MAX_BODY_SIZE) {
      return json({ error: 'Request body too large' }, { status: 413 });
    }
    const body = JSON.parse(bodyText);
    console.log('Received request body:', body);

    const { url } = body;
    if (!url) {
      return json({ error: 'URL is required' }, { status: 400 });
    }

    console.log('Fetching transcript for URL:', url);
    
    // Extract video ID
    const match = url.match(/(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/);
    const videoId = match ? match[1] : null;
    
    if (!videoId) {
      return json({ error: 'Invalid YouTube URL' }, { status: 400 });
    }

    const transcriptItems = await YoutubeTranscript.fetchTranscript(videoId);
    const transcript = transcriptItems
      .map(item => item.text)
      .join('\n');

    const response = {
      transcript,
      title: videoId
    };

    console.log('Successfully fetched transcript, preparing response');
    console.log('Response (first 200 chars):', transcript.slice(0, 200) + '...');

    return json(response);
  } catch (error) {
    console.error('Server error:', error);
    return json(
      { error: error instanceof Error ? error.message : 'Failed to fetch transcript' },
      { status: 500 }
    );
  }
};