import YAML from 'yaml';

export type NoteFrontmatter = Record<string, unknown>;

function cleanContentPreview(content: string): string {
  const cleaned = content
    .replace(/[#*`_]/g, '')
    .replace(/\s+/g, ' ')
    .trim();

  return cleaned.slice(0, 150) + (cleaned.length > 150 ? '...' : '');
}

function defaultTitle(now: Date, noteName?: string): string {
  const trimmedName = noteName?.trim();
  if (trimmedName) {
    return trimmedName;
  }

  return `Note ${now.toLocaleString()}`;
}

function cloneValue<T>(value: T): T {
  if (Array.isArray(value)) {
    return value.map((item) => cloneValue(item)) as T;
  }

  if (value && typeof value === 'object') {
    return Object.fromEntries(
      Object.entries(value as Record<string, unknown>).map(([key, nestedValue]) => [key, cloneValue(nestedValue)])
    ) as T;
  }

  return value;
}

function hasValue(value: unknown): boolean {
  if (typeof value === 'string') {
    return value.trim().length > 0;
  }

  return value !== null && value !== undefined;
}

export function createFallbackNoteFrontmatter(
  content: string,
  options: { noteName?: string; now?: Date } = {}
): NoteFrontmatter {
  const now = options.now ?? new Date();
  const timestamp = now.toISOString();

  return {
    title: defaultTitle(now, options.noteName),
    aliases: [''],
    description: cleanContentPreview(content),
    date: timestamp,
    tags: ['inbox', 'note'],
    updated: timestamp,
    author: 'User'
  };
}

export function mergePatternFrontmatter(
  patternFrontmatter: NoteFrontmatter | null | undefined,
  noteName: string,
  now: Date = new Date()
): NoteFrontmatter | null {
  if (!patternFrontmatter || Object.keys(patternFrontmatter).length === 0) {
    return null;
  }

  const merged = cloneValue(patternFrontmatter);

  if (!hasValue(merged.title)) {
    merged.title = defaultTitle(now, noteName);
  }

  if (!hasValue(merged.created) && !hasValue(merged.date)) {
    merged.created = now.toISOString().split('T')[0];
  }

  return merged;
}

export function buildNoteFrontmatter(
  content: string,
  options: { noteName?: string; patternFrontmatter?: NoteFrontmatter | null; now?: Date } = {}
): NoteFrontmatter {
  const now = options.now ?? new Date();
  const mergedPatternFrontmatter = options.noteName
    ? mergePatternFrontmatter(options.patternFrontmatter, options.noteName, now)
    : null;

  if (mergedPatternFrontmatter) {
    return mergedPatternFrontmatter;
  }

  return createFallbackNoteFrontmatter(content, {
    noteName: options.noteName,
    now
  });
}

export function serializeMarkdownNote(frontmatter: NoteFrontmatter, body: string): string {
  const yamlText = YAML.stringify(frontmatter).trimEnd();
  return `---\n${yamlText}\n---\n\n${body}`;
}
