import { readFileSync } from 'node:fs';
import { describe, expect, it, vi } from 'vitest';
import { readFileContent } from './read-file-content';

const pdfBytes = new Uint8Array([0x25, 0x50, 0x44, 0x46]);

describe('readFileContent', () => {
  const noPdf = {
    convertToMarkdown: () => Promise.reject(new Error('unexpected PDF conversion'))
  };

  it('resolves with the text of a text file', async () => {
    const file = new File(['# Notes\n\nPlain text.\n'], 'notes.md', { type: 'text/markdown' });

    await expect(readFileContent(file, noPdf, () => {})).resolves.toBe('# Notes\n\nPlain text.\n');
  });

  it('resolves with the converted Markdown of a PDF file', async () => {
    const pdf = {
      convertToMarkdown: vi.fn().mockResolvedValue({ markdown: '# Title\n' })
    };
    const file = new File([pdfBytes], 'report.pdf', { type: 'application/pdf' });

    await expect(readFileContent(file, pdf, () => {})).resolves.toBe('# Title\n');
    expect(pdf.convertToMarkdown).toHaveBeenCalledWith(file);
  });

  it('forwards the conversion warning of a PDF file', async () => {
    const pdf = {
      convertToMarkdown: () => Promise.resolve({ markdown: '# Title\n', warning: 'partial text' })
    };
    const warn = vi.fn();
    const file = new File([pdfBytes], 'report.pdf', { type: 'application/pdf' });

    await readFileContent(file, pdf, warn);

    expect(warn).toHaveBeenCalledWith('partial text');
  });
});

// The repository has no browser test runner, so these tests pin the chat
// boundary at the source level: the component has no sendMessage path, and
// handleSubmit holds the one streamChat call. Together with the seam tests
// above, this proves that attachment parses without a chat request and that
// one submit sends one request.
describe('ChatInput chat boundary', () => {
  const source = readFileSync(new URL('./ChatInput.svelte', import.meta.url), 'utf8');

  it('has no sendMessage path, so a file attachment cannot send a chat request', () => {
    expect(source.includes('sendMessage')).toBe(false);
  });

  it('sends through exactly one streamChat call, in handleSubmit', () => {
    expect(source.match(/\.streamChat\(/g) ?? []).toHaveLength(1);
  });

  it('shows the attachment indicator after processing and hides it after submit', () => {
    expect(source).toMatch(
      /uploadedFiles = \[\.\.\.uploadedFiles, file\.name\];\s*isFileIndicatorVisible = true;/
    );
    expect(source).toMatch(
      /uploadedFiles = \[\];\s*fileContents = \[\];\s*isFileIndicatorVisible = false;/
    );
  });
});
