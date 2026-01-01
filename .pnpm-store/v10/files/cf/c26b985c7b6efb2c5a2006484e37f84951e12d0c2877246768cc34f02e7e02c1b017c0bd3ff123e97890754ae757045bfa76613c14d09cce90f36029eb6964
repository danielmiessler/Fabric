"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
/**
 * Holds the information which (zero based) page index maps to a page number.
 */
class PageMapping {
    constructor(pageFactor = 1, detectedOnPage = false) {
        this.pageFactor = pageFactor;
        this.detectedOnPage = detectedOnPage;
    }
    /**
     * Translates a given page index to a page number label as printed on the page. E.g [0,1,2,3,4] could become [I, II, 1, 2].
     * @param pageIndex
     */
    pageLabel(pageIndex) {
        const pageNumber = pageIndex + this.pageFactor;
        if (pageNumber < 1) {
            return romanize(Math.abs(pageNumber - this.pageFactor) + 1);
        }
        return `${pageNumber}`;
    }
    shifted() {
        return this.pageFactor != 1;
    }
}
exports.default = PageMapping;
function romanize(num) {
    var lookup = { M: 1000, CM: 900, D: 500, CD: 400, C: 100, XC: 90, L: 50, XL: 40, X: 10, IX: 9, V: 5, IV: 4, I: 1 }, roman = '', i;
    for (i in lookup) {
        while (num >= lookup[i]) {
            roman += i;
            num -= lookup[i];
        }
    }
    return roman;
}
