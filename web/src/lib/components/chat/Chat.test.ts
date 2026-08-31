import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';

const source = readFileSync(new URL('./Chat.svelte', import.meta.url), 'utf8');

describe('Chat viewport layout', () => {
  it('fits the route viewport instead of extending to the full screen height', () => {
    expect(source).toMatch(/class="chat-container[^"]*\bh-full\b[^"]*\bmin-h-0\b/);
    expect(source).not.toMatch(/class="chat-container[^"]*\bh-screen\b/);
  });
});
