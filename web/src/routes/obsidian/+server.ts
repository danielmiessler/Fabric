import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { exec } from 'child_process';
import { promisify } from 'util';

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

export const POST: RequestHandler = async ({ request }) => {
  let tempFile: string | undefined;

  try {
    // Parse and validate request
    const body = await request.json() as ObsidianRequest;
    if (!body.pattern || !body.noteName || !body.content) {
      return json(
        { error: 'Missing required fields: pattern, noteName, or content' },
        { status: 400 }
      );
    }

    console.log('\n=== Obsidian Request ===');
    console.log('1. Pattern:', body.pattern);
    console.log('2. Note name:', body.noteName);
    console.log('3. Content length:', body.content.length);

  


    

    // Format content with markdown code blocks
    const formattedContent = `\`\`\`markdown\n${body.content}\n\`\`\``;
    const escapedFormattedContent = escapeShellArg(formattedContent);

    // Generate file name and path
    const fileName = `${new Date().toISOString().split('T')[0]}-${body.noteName}.md`;
   
    const obsidianDir = 'myfiles/Fabric_obsidian';
    const filePath = `${obsidianDir}/${fileName}`;
    await execAsync(`mkdir -p "${obsidianDir}"`);
    console.log('4. Ensured Obsidian directory exists');


    // Create temp file
    tempFile = `/tmp/fabric-${Date.now()}.txt`;

    // Write formatted content to temp file
    await execAsync(`echo ${escapedFormattedContent} > "${tempFile}"`);
    console.log('5. Wrote formatted content to temp file');

    // Copy from temp file to final location (safer than direct write)
    await execAsync(`cp "${tempFile}" "${filePath}"`);
    console.log('6. Copied content to final location:', filePath);

    // Verify file was created and has content
    const { stdout: lsOutput } = await execAsync(`ls -l "${filePath}"`);
    const { stdout: wcOutput } = await execAsync(`wc -l "${filePath}"`);
    console.log('7. File verification:', lsOutput);
    console.log('8. Line count:', wcOutput);

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
        console.log('9. Cleaned up temp file');
      } catch (cleanupError) {
        console.error('Failed to clean up temp file:', cleanupError);
      }
    }
  }
};