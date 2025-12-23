import { json, error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { exec } from 'child_process';
import { promisify } from 'util';
import path from 'path';

const execAsync = promisify(exec);

interface ObsidianRequest {
  pattern: string;
  noteName: string;
  content: string;
}

function escapeShellArg(arg: string): string {
  // Replace single quotes with '\'' and wrap in single quotes
  return `'${arg.replace(/'/g, "'\\''")}'`;
}

// Helper function to get and validate user's vault path
function getUserVaultPath(userId: string): string {
  // Validate userId to prevent path traversal
  if (!userId || /[\/\\]|\.\./.test(userId)) {
    throw new Error('Invalid user identifier');
  }
  
  // Construct user-specific vault directory
  const baseDir = path.join(process.cwd(), 'user_vaults');
  const userVaultPath = path.join(baseDir, userId);
  
  // Ensure the resolved path is within baseDir to prevent escape attempts
  const resolvedPath = path.resolve(userVaultPath);
  const resolvedBaseDir = path.resolve(baseDir);
  
  if (!resolvedPath.startsWith(resolvedBaseDir)) {
    throw new Error('Access denied: Invalid vault path');
  }
  
  return userVaultPath;
}

export const POST: RequestHandler = async ({ request, locals }) => {
  let tempFile: string | undefined;

  try {
    // Check if user is authenticated
    const userId = locals?.user?.id;
    if (!userId) {
      return json(
        { error: 'Unauthorized: User must be authenticated' },
        { status: 401 }
      );
    }

    // Parse and validate request
    const body = await request.json() as ObsidianRequest;
    if (!body.pattern || !body.noteName || !body.content) {
      return json(
        { error: 'Missing required fields: pattern, noteName, or content' },
        { status: 400 }
      );
    }

    // Get user's vault path with authorization checks
    const userVaultDir = getUserVaultPath(userId);

    console.log('\n=== Obsidian Request ===');
    console.log('1. User ID:', userId);
    console.log('2. Pattern:', body.pattern);
    console.log('3. Note name:', body.noteName);
    console.log('4. Content length:', body.content.length);

  


    

    // Format content with markdown code blocks
    const formattedContent = `\`\`\`markdown\n${body.content}\n\`\`\``;
    const escapedFormattedContent = escapeShellArg(formattedContent);

    // Generate file name and path using user's vault directory
    const fileName = `${new Date().toISOString().split('T')[0]}-${body.noteName}.md`;
    const filePath = path.join(userVaultDir, fileName);
    
    await execAsync(`mkdir -p "${userVaultDir}"`);
    console.log('5. Ensured user vault directory exists');


    // Create temp file
    tempFile = `/tmp/fabric-${Date.now()}.txt`;

    // Write formatted content to temp file
    await execAsync(`echo ${escapedFormattedContent} > "${tempFile}"`);
    console.log('6. Wrote formatted content to temp file');

    // Copy from temp file to final location (safer than direct write)
    await execAsync(`cp "${tempFile}" "${filePath}"`);
    console.log('7. Copied content to final location:', filePath);

    // Verify file was created and has content
    const { stdout: lsOutput } = await execAsync(`ls -l "${filePath}"`);
    const { stdout: wcOutput } = await execAsync(`wc -l "${filePath}"`);
    console.log('8. File verification:', lsOutput);
    console.log('9. Line count:', wcOutput);

    // Return success response with file details
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

  } finally {
    // Clean up temp file if it exists
    if (tempFile) {
      try {
        await execAsync(`rm -f "${tempFile}"`);
        console.log('10. Cleaned up temp file');
      } catch (cleanupError) {
        console.error('Failed to clean up temp file:', cleanupError);
      }
    }
  }
};