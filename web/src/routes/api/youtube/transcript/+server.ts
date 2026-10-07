import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { YoutubeTranscript } from 'youtube-transcript';
import { youtubeVideoId } from '$lib/utils/validators';

export const POST: RequestHandler = async ({ request }) => {
  try {
    const body = await request.json();
    const { url } = body;
    if (!url) {
      return json({ error: 'URL is required' }, { status: 400 });
    }

    console.log('Fetching transcript for URL:', url);
    
    const videoId = youtubeVideoId(url);
    
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

    return json(response);
  } catch (error) {
    console.error('Server error:', error);
    return json(
      { error: 'Failed to fetch transcript' },
      { status: 500 }
    );
  }
};