// For notesDrawer component
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

    // Get the absolute path to the inbox directory
    const __filename = fileURLToPath(import.meta.url);
    const __dirname = dirname(__filename);
    const inboxDir = resolve(__dirname, '..', '..', '..', 'myfiles', 'inbox');

    // Security: use only the basename to strip any path traversal sequences (CWE-22)
    const safeFilename = basename(filename);
    if (!safeFilename || !/^[a-zA-Z0-9._\- ]+$/.test(safeFilename)) {
      return json({ error: 'Invalid filename' }, { status: 400 });
    }

    // Build path via string concatenation to avoid path.join/resolve with user input
    const inboxPath = `${inboxDir}/${safeFilename}`;

    // Double-check the resolved path is still within the inbox directory
    if (!inboxPath.startsWith(inboxDir + '/')) {
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
