"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.PARSE_SCHEMA = void 0;
const Item_1 = require("./Item");
const Metadata_1 = require("./Metadata");
const ParseResult_1 = require("./ParseResult");
exports.PARSE_SCHEMA = ['transform', 'width', 'height', 'str', 'fontName', 'dir'];
/**
 * Parses a PDF via PDFJS and returns a ParseResult which contains more or less the original data from PDFJS.
 */
class PdfParser {
    constructor(pdfjs, defaultParams = {}) {
        this.schema = exports.PARSE_SCHEMA;
        this.pdfjs = pdfjs;
        this.defaultParams = defaultParams;
    }
    parse(src, reporter) {
        return __awaiter(this, void 0, void 0, function* () {
            const documentInitParameters = Object.assign(Object.assign({}, this.defaultParams), this.documentInitParameters(src));
            return this.pdfjs
                .getDocument(documentInitParameters)
                .promise.then((pdfjsDocument) => {
                reporter.parsedDocumentHeader(pdfjsDocument.numPages);
                return Promise.all([
                    pdfjsDocument,
                    pdfjsDocument.getMetadata().then((pdfjsMetadata) => {
                        reporter.parsedMetadata();
                        return new Metadata_1.default(pdfjsMetadata);
                    }),
                    this.extractPagesSequentially(pdfjsDocument, reporter),
                ]);
            })
                .then(([pdfjsDocument, metadata, pages]) => {
                return Promise.all([
                    pdfjsDocument,
                    metadata,
                    pages,
                    this.gatherFontObjects(pages).finally(() => reporter.parsedFonts()),
                ]);
            })
                .then(([pdfjsDocument, metadata, pages, fontMap]) => {
                const pdfjsPages = pages.map((page) => page.pdfjsPage);
                const items = pages.reduce((allItems, page) => allItems.concat(page.items), []);
                const pageViewports = pdfjsPages.map((page) => {
                    const viewPort = page.getViewport({ scale: 1.0 });
                    return {
                        transformFunction: (itemTransform) => this.pdfjs.Util.transform(viewPort.transform, itemTransform),
                    };
                });
                return new ParseResult_1.default(fontMap, pdfjsDocument.numPages, pdfjsPages, pageViewports, metadata, this.schema, items);
            });
        });
    }
    extractPagesSequentially(pdfjsDocument, reporter) {
        return [...Array(pdfjsDocument.numPages)].reduce((accumulatorPromise, _, index) => {
            return accumulatorPromise.then((accumulatedResults) => {
                return pdfjsDocument.getPage(index + 1).then((pdfjsPage) => {
                    return pdfjsPage
                        .getTextContent({
                        normalizeWhitespace: false,
                        disableCombineTextItems: true,
                    })
                        .then((textContent) => {
                        const items = textContent.items.map((pdfjsItem) => new Item_1.default(index, pdfjsItem));
                        reporter.parsedPage(index);
                        return [...accumulatedResults, { index, pdfjsPage, items }];
                    });
                });
            });
        }, Promise.resolve([]));
    }
    gatherFontObjects(pages) {
        const uniqueFontIds = new Set();
        return pages.reduce((promise, page) => {
            const unknownPageFonts = page.items.reduce((unknowns, item) => {
                const fontId = item.data['fontName'];
                if (!uniqueFontIds.has(fontId) && fontId.startsWith('g_d')) {
                    uniqueFontIds.add(fontId);
                    unknowns.push(fontId);
                }
                return unknowns;
            }, []);
            if (unknownPageFonts.length > 0) {
                // console.log(`Fetch fonts ${unknownPageFonts} for page ${page.index}`);
                promise = promise.then((fontMap) => {
                    return page.pdfjsPage.getOperatorList().then(() => {
                        unknownPageFonts.forEach((fontId) => {
                            const fontObject = page.pdfjsPage.commonObjs.get(fontId);
                            fontMap.set(fontId, fontObject);
                        });
                        return fontMap;
                    });
                });
            }
            return promise;
        }, Promise.resolve(new Map()));
    }
    documentInitParameters(src) {
        if (typeof src === 'string') {
            return { url: src };
        }
        if (this.isArrayBuffer(src)) {
            return { data: src };
        }
        if (typeof src === 'object') {
            return src;
        }
        throw new Error('Invalid PDFjs parameter for getDocument. Need either Uint8Array, string or a parameter object');
    }
    isArrayBuffer(object) {
        return typeof object === 'object' && object !== null && object.byteLength !== undefined;
    }
}
exports.default = PdfParser;
