import type { Transformer, Preprocessor } from '../types';
/** Create a tag matching regexp. */
export declare function createTagRegex(tagName: string, flags?: string): RegExp;
/** Strip script and style tags from markup. */
export declare function stripTags(markup: string): string;
/** Transform an attribute string into a key-value object */
export declare function parseAttributes(attributesStr: string): Record<string, any>;
export declare function transformMarkup({ content, filename }: {
    content: string;
    filename?: string;
}, transformer: Preprocessor | Transformer<unknown>, options?: Record<string, any>): Promise<import("../types").Processed>;
