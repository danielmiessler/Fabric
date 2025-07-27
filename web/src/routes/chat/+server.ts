import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { YoutubeTranscript } from 'youtube-transcript';

export const POST: RequestHandler = async ({ request }) => {
  try {
    const body = await request.json();
    console.log('\n=== Request Analysis ===');
    console.log('1. Raw request body:', JSON.stringify(body, null, 2));

    // Handle YouTube URL request
    if (body.url) {
      console.log('2. Processing YouTube URL:', {
        url: body.url,
        language: body.language,
        hasLanguageParam: true
      });

      // Extract video ID
      const match = body.url.match(/(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/);
      const videoId = match ? match[1] : null;

      if (!videoId) {
        return json({ error: 'Invalid YouTube URL' }, { status: 400 });
      }

      console.log('3. Video ID:', {
        id: videoId,
        language: body.language
      });

      const transcriptItems = await YoutubeTranscript.fetchTranscript(videoId);
      const transcript = transcriptItems
        .map(item => item.text)
        .join(' ');

      // Create response with transcript and language
      const response = {
        transcript,
        title: videoId,
        language: body.language
      };

      console.log('4. Transcript processed:', {
        length: transcript.length,
        language: body.language,
        firstChars: transcript.substring(0, 50),
        responseSize: JSON.stringify(response).length
      });

      return json(response);
    }

    // Handle pattern execution request
    console.log('\n=== Server Request Analysis ===');
    console.log('1. Request overview:', {
      pattern: body.prompts?.[0]?.patternName,
      hasPrompts: !!body.prompts?.length,
      messageCount: body.messages?.length,
      isYouTube: body.url ? 'Yes' : 'No',
      language: body.language
    });

    // Removed redundant language instruction logic; Go backend handles this
    // if (body.prompts?.[0] && body.language && body.language !== 'en') {
    //   const languageInstruction = `. Please use the language '${body.language}' for the output.`;
    //   if (!body.prompts[0].userInput?.includes(languageInstruction)) {
    //     body.prompts[0].userInput = (body.prompts[0].userInput || '') + languageInstruction;
    //   }
    // }

    console.log('2. Language analysis:', {
      input: body.prompts?.[0]?.userInput?.substring(0, 100), // Note: This input no longer has the instruction appended here
      hasLanguageInstruction: body.prompts?.[0]?.userInput?.includes('language'),
      containsFr: body.prompts?.[0]?.userInput?.includes('fr'),
      containsEn: body.prompts?.[0]?.userInput?.includes('en'),
      requestLanguage: body.language
    });

    // Log full request for debugging
    console.log('3. Full request:', JSON.stringify(body, null, 2));

    // Log important fields
    console.log('4. Key fields:', {
      patternName: body.prompts?.[0]?.patternName,
      inputLength: body.prompts?.[0]?.userInput?.length,
      systemPromptLength: body.prompts?.[0]?.systemPrompt?.length,
      messageCount: body.messages?.length
    });

    console.log('5. Sending to Fabric backend...');
    const fabricResponse = await fetch('http://localhost:8080/api/chat', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    });

    console.log('6. Fabric response:', {
      status: fabricResponse.status,
      ok: fabricResponse.ok,
      statusText: fabricResponse.statusText
    });

    if (!fabricResponse.ok) {
      console.error('Error from Fabric API:', {
        status: fabricResponse.status,
        statusText: fabricResponse.statusText
      });
      throw new Error(`Fabric API error: ${fabricResponse.statusText}`);
    }

    const stream = fabricResponse.body;
    if (!stream) {
      throw new Error('No response from fabric backend');
    }

    // Create a TransformStream to inspect the data without modifying it
    const transformStream = new TransformStream({
      transform(chunk, controller) {
        const text = new TextDecoder().decode(chunk);
        if (text.startsWith('data: ')) {
          try {
            const data = JSON.parse(text.slice(6));
            console.log('Stream chunk format:', {
              type: data.type,
              format: data.format,
              contentLength: data.content?.length
            });
          } catch (e) {
            console.log('Failed to parse stream chunk:', text);
          }
        }
        controller.enqueue(chunk);
      }
    });

    // Pipe through the transform stream
    const transformedStream = stream.pipeThrough(transformStream);

    // Return the transformed stream
    const response = new Response(transformedStream, {
      headers: {
        'Content-Type': 'text/event-stream',
        'Cache-Control': 'no-cache',
        'Connection': 'keep-alive'
      }
    });

    return response;

  } catch (error) {
    console.error('\n=== Error ===');
    console.error('Type:', error?.constructor?.name);
    console.error('Message:', error instanceof Error ? error.message : String(error));
    console.error('Stack:', error instanceof Error ? error.stack : 'No stack trace');
    return json(
      { error: error instanceof Error ? error.message : 'Failed to process request' },
      { status: 500 }
    );
  }
};
