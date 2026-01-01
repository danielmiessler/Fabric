import type { Node, Root } from "postcss";
import type { Context } from "../context";
import type { SourceLocation, SvelteStyleElement } from "../ast";
export type StyleContext = StyleContextNoStyleElement | StyleContextParseError | StyleContextSuccess | StyleContextUnknownLang;
export interface StyleContextNoStyleElement {
    status: "no-style-element";
}
export interface StyleContextParseError {
    status: "parse-error";
    sourceLang: string;
    error: any;
}
export interface StyleContextSuccess {
    status: "success";
    sourceLang: string;
    sourceAst: Root;
}
export interface StyleContextUnknownLang {
    status: "unknown-lang";
    sourceLang: string;
}
/**
 * Extracts style source from a SvelteStyleElement and parses it into a PostCSS AST.
 */
export declare function parseStyleContext(styleElement: SvelteStyleElement | undefined, ctx: Context): StyleContext;
/**
 * Extracts a node location (like that of any ESLint node) from a parsed svelte style node.
 */
export declare function styleNodeLoc(node: Node): Partial<SourceLocation>;
/**
 * Extracts a node range (like that of any ESLint node) from a parsed svelte style node.
 */
export declare function styleNodeRange(node: Node): [number | undefined, number | undefined];
