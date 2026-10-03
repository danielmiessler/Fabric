import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { mkdir, writeFile, stat } from 'fs/promises';
import { resolve, basename, join } from 'path';

interface ObsidianRequest {
  pattern: string;
  noteName: string;
  content: string;
}

// Allowlist for note names. It blocks path traversal (CWE-22).
const SAFE_NOTE_NAME = /^[a-zA-Z0-9 _.-]+$/;

export const POST: RequestHandler = async ({ request }) => {
  try {
    const body = await request.json() as ObsidianRequest;
    if (!body.pattern || !body.noteName || !body.content) {
      return json(
        { error: 'Missing required fields: pattern, noteName, or content' },
        { status: 400 }
      );
    }

    const safeNoteName = basename(body.noteName);
    if (!safeNoteName || !SAFE_NOTE_NAME.test(safeNoteName)) {
      return json({ error: 'Invalid note name' }, { status: 400 });
    }

    console.log('\n=== Obsidian Request ===');
    console.log('1. Pattern:', body.pattern);
    console.log('2. Note name:', safeNoteName);
    console.log('3. Content length:', body.content.length);

    const formattedContent = `\`\`\`markdown\n${body.content}\n\`\`\``;

    const fileName = `${new Date().toISOString().split('T')[0]}-${safeNoteName}.md`;
    const obsidianDir = resolve('myfiles/Fabric_obsidian');
    const filePath = join(obsidianDir, fileName);

    // Defense in depth: the resolved path must stay inside obsidianDir (CWE-22).
    if (!filePath.startsWith(obsidianDir + '/') && filePath !== obsidianDir) {
      return json({ error: 'Invalid note name' }, { status: 400 });
    }

    await mkdir(obsidianDir, { recursive: true });
    console.log('4. Ensured Obsidian directory exists');

    await writeFile(filePath, formattedContent, 'utf-8');
    console.log('5. Wrote content to final location:', filePath);

    const fileStats = await stat(filePath);
    const lineCount = formattedContent.split('\n').length;
    console.log('6. File verification: size =', fileStats.size, 'bytes,', lineCount, 'lines');

    return json({
      success: true,
      fileName,
      filePath,
      message: `Successfully saved to ${fileName}`
    });

  } catch (error) {
    console.error('\n=== Error ===');
    console.error('Type:', error?.constructor?.name);
    console.error('Message:', error instanceof Error ? error.message : String(error));
    console.error('Stack:', error instanceof Error ? error.stack : 'No stack trace');

    return json(
      {
        error: error instanceof Error ? error.message : 'Failed to process request',
        details: error instanceof Error ? error.stack : undefined
      },
      { status: 500 }
    );
  }
};
