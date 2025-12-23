// For notesDrawer component
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { writeFile } from 'fs/promises';
import { join, resolve } from 'path';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

export const POST: RequestHandler = async ({ request, locals }) => {
  try {
    // Check authentication - ensure user is authenticated before allowing file operations
    if (!locals.user) {
      return json({ error: 'Unauthorized' }, { status: 401 });
    }

    const { filename, content } = await request.json();

    if (!filename || !content) {
      return json({ error: 'Filename and content are required' }, { status: 400 });
    }

    // Validate and sanitize the filename to prevent path traversal attacks
    if (typeof filename !== 'string' || filename.includes('..') || filename.includes('/') || filename.includes('\\')) {
      return json({ error: 'Invalid filename' }, { status: 400 });
    }

    // Get the absolute path to the inbox directory
    const __filename = fileURLToPath(import.meta.url);
    const __dirname = dirname(__filename);
    // const inboxPath = join(__dirname, '..', 'myfiles', 'inbox', filename);
    // New version using environment variables:
    // const inboxPath = join(process.env.DATA_DIR || './web/myfiles', 'inbox', filename);
    const baseDir = resolve(__dirname, '..', '..', '..', 'myfiles', 'inbox');
    const inboxPath = resolve(baseDir, filename);

    // Ensure the resolved path is within the intended directory
    if (!inboxPath.startsWith(baseDir)) {
      return json({ error: 'Invalid filename' }, { status: 400 });
    }

    await writeFile(inboxPath, content, 'utf-8');

    return json({ success: true, filename });
  } catch (error) {
    console.error('Server error:', error);
    return json(
      { error: error instanceof Error ? error.message : 'Failed to save note' },
      { status: 500 }
    );
  }
};
