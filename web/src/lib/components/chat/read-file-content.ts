import type { PdfConversion } from '../../services/PdfConversionService';

// The parser seam for ChatInput. It converts one attached file to text.
// Keep this module free of chat and store imports: that constraint is what
// makes a file attachment unable to send a chat request.
export async function readFileContent(
  file: File,
  pdf: { convertToMarkdown(file: File): Promise<PdfConversion> },
  warn: (message: string) => void
): Promise<string> {
  if (file.type === 'application/pdf') {
    const { markdown, warning } = await pdf.convertToMarkdown(file);
    if (warning) warn(warning);
    return markdown;
  }
  return file.text();
}
