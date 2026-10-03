// Saves a note from NoteDrawer.svelte. See note-store.ts.
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { writeFile } from 'fs/promises';
import { join, resolve, basename } from 'path';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

export const POST: RequestHandler = async ({ request }) => {
  try {
    const { filename, content } = await request.json();

    if (!filename || !content) {
      return json({ error: 'Filename and content are required' }, { status: 400 });
    }

    const __filename = fileURLToPath(import.meta.url);
    const __dirname = dirname(__filename);
    const inboxDir = resolve(__dirname, '..', '..', '..', 'myfiles', 'inbox');

    // basename strips path traversal sequences (CWE-22).
    const safeFilename = basename(filename);
    if (!safeFilename) {
      return json({ error: 'Invalid filename' }, { status: 400 });
    }

    const inboxPath = join(inboxDir, safeFilename);

    // Defense in depth: the resolved path must stay inside inboxDir (CWE-22).
    if (!inboxPath.startsWith(inboxDir + '/') && inboxPath !== inboxDir) {
      return json({ error: 'Invalid filename' }, { status: 400 });
    }

    await writeFile(inboxPath, content, 'utf-8');

    return json({ success: true, filename: safeFilename });
  } catch (error) {
    console.error('Server error:', error);
    return json(
      { error: error instanceof Error ? error.message : 'Failed to save note' },
      { status: 500 }
    );
  }
};
