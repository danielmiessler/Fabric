import type Item from './Item';
import type Metadata from './Metadata';
import type PageViewport from './parse/PageViewport';

export default class ParseResult {
  constructor(
    public fontMap: Map<string, object>,
    public pageCount: number,
    public pdfjsPages: any[],
    public pageViewports: PageViewport[],
    public metadata: Metadata,
    public schema: string[],
    public items: Item[],
  ) {}
}
