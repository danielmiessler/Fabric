"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
/**
 * Table of contents usually parsed by  `DetectToc.ts`.
 */
class TOC {
    constructor(tocHeadlineItems, pages, detectedHeadlineLevels) {
        this.tocHeadlineItems = tocHeadlineItems;
        this.pages = pages;
        this.detectedHeadlineLevels = detectedHeadlineLevels;
    }
    startPage() {
        return Math.min(...this.pages);
    }
    endPage() {
        return Math.max(...this.pages);
    }
}
exports.default = TOC;
