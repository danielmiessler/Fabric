export interface Frontmatter {
  title: string | null;
  aliases?: string[];
  description: string | null;
  date: string;
  tags: string[] | null;
  updated?: string;
  author?: string;
  layout?: string;
  images?: string[];
}

export interface MdsvexCompileData {
  fm: Frontmatter;
  [key: string]: unknown;
}
