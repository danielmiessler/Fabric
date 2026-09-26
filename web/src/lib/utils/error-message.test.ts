import { describe, expect, it } from 'vitest';
import { formatErrorMessage } from './error-message';

describe('formatErrorMessage', () => {
  // The server sends 'Error: could not get pattern write_essay: missing
  // required variable: author_name'. The reader must not see the word twice.
  it('keeps one Error word when the text already starts with it', () => {
    const serverText =
      'Error: could not get pattern write_essay: missing required variable: author_name';

    expect(formatErrorMessage(new Error(serverText))).toBe(serverText);
  });

  it('adds the Error word when the text does not start with it', () => {
    expect(formatErrorMessage(new Error('the vendor refused the request'))).toBe(
      'Error: the vendor refused the request'
    );
  });

  it('reads a plain string', () => {
    expect(formatErrorMessage('something broke')).toBe('Error: something broke');
  });

  it('gives a message for an empty error', () => {
    expect(formatErrorMessage(new Error(''))).toBe('Error: unknown error');
    expect(formatErrorMessage('   ')).toBe('Error: unknown error');
  });

  it('reads a value that is not an error and not a string', () => {
    expect(formatErrorMessage({ code: 12 })).toBe('Error: [object Object]');
    expect(formatErrorMessage(null)).toBe('Error: null');
  });
});
