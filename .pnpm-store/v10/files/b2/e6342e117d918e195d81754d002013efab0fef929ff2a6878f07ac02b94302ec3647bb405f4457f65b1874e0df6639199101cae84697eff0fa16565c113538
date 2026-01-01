import type { BaseNode } from "./base";
export interface Position {
    /** >= 1 */
    line: number;
    /** >= 0 */
    column: number;
}
export type Range = [number, number];
export interface SourceLocation {
    start: Position;
    end: Position;
}
export interface Locations {
    loc: SourceLocation;
    range: Range;
}
export interface Token extends BaseNode {
    type: "Boolean" | "Null" | "Identifier" | "Keyword" | "Punctuator" | "JSXIdentifier" | "JSXText" | "Numeric" | "String" | "RegularExpression" | "Template" | "HTMLText" | "HTMLIdentifier" | "MustacheKeyword" | "HTMLComment";
    value: string;
}
export interface Comment extends BaseNode {
    type: "Line" | "Block";
    value: string;
}
