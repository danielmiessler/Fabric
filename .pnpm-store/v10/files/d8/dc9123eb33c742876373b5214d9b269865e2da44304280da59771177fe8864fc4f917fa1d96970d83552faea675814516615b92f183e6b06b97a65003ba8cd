/**
 * Holds the information which (zero based) page index maps to a page number.
 */
export default class PageMapping {
    pageFactor: number;
    detectedOnPage: boolean;
    constructor(pageFactor?: number, detectedOnPage?: boolean);
    /**
     * Translates a given page index to a page number label as printed on the page. E.g [0,1,2,3,4] could become [I, II, 1, 2].
     * @param pageIndex
     */
    pageLabel(pageIndex: number): string;
    shifted(): boolean;
}
