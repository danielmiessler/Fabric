import Item from './Item';
import Metadata from './Metadata';
import type ParseReporter from './ParseReporter';
import ParseResult from './ParseResult';

export const PARSE_SCHEMA = ['transform', 'width', 'height', 'str', 'fontName', 'dir'];

/**
 * Parses a PDF via PDFJS and returns a ParseResult which contains more or less the original data from PDFJS.
 */
export default class PdfParser {
  pdfjs: any;
  defaultParams: object;
  schema = PARSE_SCHEMA;

  constructor(pdfjs: any, defaultParams = {}) {
    this.pdfjs = pdfjs;
    this.defaultParams = defaultParams;
  }

  async parse(src: string | Uint8Array | object, reporter: ParseReporter): Promise<ParseResult> {
    const documentInitParameters = { ...this.defaultParams, ...this.documentInitParameters(src) };
    return this.pdfjs
      .getDocument(documentInitParameters)
      .promise.then((pdfjsDocument: any) => {
        reporter.parsedDocumentHeader(pdfjsDocument.numPages);
        return Promise.all([
          pdfjsDocument,
          pdfjsDocument.getMetadata().then((pdfjsMetadata: any) => {
            reporter.parsedMetadata();
            return new Metadata(pdfjsMetadata);
          }),
          this.extractPagesSequentially(pdfjsDocument, reporter),
        ]);
      })
      .then(([pdfjsDocument, metadata, pages]: [any, Metadata, ParsedPage[]]) => {
        return Promise.all([
          pdfjsDocument,
          metadata,
          pages,
          this.gatherFontObjects(pages).finally(() => reporter.parsedFonts()),
        ]);
      })
      .then(([pdfjsDocument, metadata, pages, fontMap]: [any, Metadata, ParsedPage[], Map<string, object>]) => {
        const pdfjsPages = pages.map((page: any) => page.pdfjsPage);
        const items = pages.reduce((allItems: any[], page: any) => allItems.concat(page.items), []);
        const pageViewports = pdfjsPages.map((page: any) => {
          const viewPort = page.getViewport({ scale: 1.0 });
          return {
            transformFunction: (itemTransform: number[]) =>
              this.pdfjs.Util.transform(viewPort.transform, itemTransform),
          };
        });
        return new ParseResult(
          fontMap,
          pdfjsDocument.numPages,
          pdfjsPages,
          pageViewports,
          metadata,
          this.schema,
          items,
        );
      });
  }

  private extractPagesSequentially(pdfjsDocument: any, reporter: ParseReporter): Promise<ParsedPage[]> {
    return [...Array(pdfjsDocument.numPages)].reduce((accumulatorPromise, _, index) => {
      return accumulatorPromise.then((accumulatedResults: ParsedPage[]) => {
        return pdfjsDocument.getPage(index + 1).then((pdfjsPage: any) => {
          return pdfjsPage
            .getTextContent({
              normalizeWhitespace: false,
              disableCombineTextItems: true,
            })
            .then((textContent: any) => {
              const items = textContent.items.map((pdfjsItem: any) => new Item(index, pdfjsItem));
              reporter.parsedPage(index);
              return [...accumulatedResults, { index, pdfjsPage, items }];
            });
        });
      });
    }, Promise.resolve([]));
  }

  private gatherFontObjects(pages: ParsedPage[]): Promise<Map<string, object>> {
    const uniqueFontIds = new Set<string>();
    return pages.reduce((promise, page) => {
      const unknownPageFonts = page.items.reduce((unknowns: string[], item) => {
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
    }, Promise.resolve(new Map<string, object>()));
  }

  private documentInitParameters(src: string | Uint8Array | object): object {
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

  private isArrayBuffer(object: any) {
    return typeof object === 'object' && object !== null && object.byteLength !== undefined;
  }
}

interface ParsedPage {
  index: number;
  pdfjsPage: any;
  items: Item[];
}
