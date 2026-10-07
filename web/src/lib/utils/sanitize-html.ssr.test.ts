// @vitest-environment node
import { describe, expect, it } from 'vitest';
import { renderMarkdown } from './sanitize-html';

// Server rendering has no DOM, thus DOMPurify cannot clean HTML.
describe('renderMarkdown with no DOM', () => {
  it('gives escaped text', () => {
    expect(renderMarkdown('**a** <b>')).toBe('**a** &#60;b&#62;');
  });
});
