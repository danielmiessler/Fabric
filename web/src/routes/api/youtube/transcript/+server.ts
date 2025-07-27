import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { YoutubeTranscript } from 'youtube-transcript';

export const POST: RequestHandler = async ({ request }) => {
  try {
    const body = await request.json();
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
      .join(' ');

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