import { describe, expect, it } from 'vitest';
import YAML from 'yaml';
import {
  buildNoteFrontmatter,
  createFallbackNoteFrontmatter,
  mergePatternFrontmatter,
  serializeMarkdownNote
} from './frontmatter';

describe('frontmatter utilities', () => {
  it('creates fallback frontmatter with legacy-style defaults', () => {
    const now = new Date('2025-01-02T03:04:05.000Z');
    const frontmatter = createFallbackNoteFrontmatter('# Hello world', { noteName: 'My Note', now });

    expect(frontmatter).toEqual({
      title: 'My Note',
      aliases: [''],
      description: 'Hello world',
      date: '2025-01-02T03:04:05.000Z',
      tags: ['inbox', 'note'],
      updated: '2025-01-02T03:04:05.000Z',
      author: 'User'
    });
  });

  it('merges pattern frontmatter and only applies minimal runtime defaults', () => {
    const now = new Date('2025-01-02T03:04:05.000Z');
    const frontmatter = mergePatternFrontmatter(
      {
        status: 'active',
        tags: ['obsidian']
      },
      'Reference Note',
      now
    );

    expect(frontmatter).toEqual({
      status: 'active',
      tags: ['obsidian'],
      title: 'Reference Note',
      created: '2025-01-02'
    });
  });

  it('preserves authored created/date fields when pattern frontmatter exists', () => {
    const now = new Date('2025-01-02T03:04:05.000Z');
    const frontmatter = buildNoteFrontmatter('Body', {
      noteName: 'Reference Note',
      patternFrontmatter: {
        title: 'Authored Title',
        date: '1970-01-01',
        status: 'draft'
      },
      now
    });

    expect(frontmatter).toEqual({
      title: 'Authored Title',
      date: '1970-01-01',
      status: 'draft'
    });
  });

  it('serializes a markdown note with YAML frontmatter and raw body', () => {
    const serialized = serializeMarkdownNote(
      {
        title: 'Reference Note',
        tags: ['obsidian', 'reference']
      },
      '# Heading\n\nBody copy'
    );

    expect(serialized.startsWith('---\n')).toBe(true);
    expect(serialized.includes('\n---\n\n# Heading\n\nBody copy')).toBe(true);

    const [, yamlText] = serialized.split('---\n');
    const parsed = YAML.parse(yamlText.split('\n---')[0]);
    expect(parsed).toEqual({
      title: 'Reference Note',
      tags: ['obsidian', 'reference']
    });
  });
});
