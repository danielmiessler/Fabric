import type ParseReporter from './ParseReporter';
import ParseResult from './ParseResult';
export declare const PARSE_SCHEMA: string[];
/**
 * Parses a PDF via PDFJS and returns a ParseResult which contains more or less the original data from PDFJS.
 */
export default class PdfParser {
    pdfjs: any;
    defaultParams: object;
    schema: string[];
    constructor(pdfjs: any, defaultParams?: {});
    parse(src: string | Uint8Array | object, reporter: ParseReporter): Promise<ParseResult>;
    private extractPagesSequentially;
    private gatherFontObjects;
    private documentInitParameters;
    private isArrayBuffer;
}
