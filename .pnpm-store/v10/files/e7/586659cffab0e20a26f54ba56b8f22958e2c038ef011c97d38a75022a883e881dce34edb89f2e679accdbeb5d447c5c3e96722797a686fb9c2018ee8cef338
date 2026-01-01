import type { ASTNode, SourceCode } from '../../types';
import type { AST } from 'svelte-eslint-parser';
import type { OffsetContext } from './offset-context';
export type AnyToken = AST.Token | AST.Comment;
export type MaybeNode = {
    type: string;
    range: [number, number];
    loc: AST.SourceLocation;
};
export type IndentOptions = {
    indentChar: ' ' | '\t';
    indentScript: boolean;
    indentSize: number;
    switchCase: number;
    alignAttributesVertically: boolean;
    ignoredNodes: string[];
};
export type IndentContext = {
    sourceCode: SourceCode;
    options: IndentOptions;
    offsets: OffsetContext;
};
/**
 * Get the first and last tokens of the given node.
 * If the node is parenthesized, this gets the outermost parentheses.
 * If the node have whitespace at the start and the end, they will be skipped.
 */
export declare function getFirstAndLastTokens(sourceCode: SourceCode, node: ASTNode | AnyToken | MaybeNode, borderOffset?: number): {
    firstToken: AST.Token;
    lastToken: AST.Token;
};
/**
 * Check whether the given node or token is the beginning of a line.
 */
export declare function isBeginningOfLine(sourceCode: SourceCode, node: ASTNode | AnyToken | MaybeNode): boolean;
/**
 * Check whether the given node is the beginning of element.
 */
export declare function isBeginningOfElement(node: AST.SvelteText): boolean;
