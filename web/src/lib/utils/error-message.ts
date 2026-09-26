/**
 * Makes the text that tells the person what went wrong.
 *
 * The server writes its stream errors as "Error: <what happened>", and the
 * client showed that text after a second "Error: " of its own, which gave
 * "Error: Error: could not get pattern ...". Add the word only when the text
 * does not start with it.
 */
export function formatErrorMessage(error: unknown): string {
  const message =
    error instanceof Error ? error.message : typeof error === 'string' ? error : String(error);

  const trimmed = message.trim();
  if (trimmed === '') {
    return 'Error: unknown error';
  }
  if (/^error\b/i.test(trimmed)) {
    return trimmed;
  }
  return `Error: ${trimmed}`;
}
