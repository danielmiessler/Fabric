import Item from './Item';
import { HeadlineType } from './text-types';

/**
 * Table of contents usually parsed by  `DetectToc.ts`.
 */
export default class TOC {
  constructor(
    public tocHeadlineItems: Item[],
    public pages: number[],
    public detectedHeadlineLevels: Set<HeadlineType>,
  ) {}

  startPage(): number {
    return Math.min(...this.pages);
  }

  endPage(): number {
    return Math.max(...this.pages);
  }
}
