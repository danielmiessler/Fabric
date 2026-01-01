import type { Comment, SvelteProgram, Token } from "../ast";
import type { Program } from "estree";
import type { ScopeManager } from "eslint-scope";
import type * as SvAST from "./svelte-ast-types";
import type * as Compiler from "./svelte-ast-types-for-v5";
import { type StyleContext, type StyleContextNoStyleElement, type StyleContextParseError, type StyleContextSuccess, type StyleContextUnknownLang } from "./style-context";
import { type SvelteParseContext } from "./svelte-parse-context";
export { StyleContext, StyleContextNoStyleElement, StyleContextParseError, StyleContextSuccess, StyleContextUnknownLang, };
export interface ESLintProgram extends Program {
    comments: Comment[];
    tokens: Token[];
}
/**
 * The parsing result of ESLint custom parsers.
 */
export interface ESLintExtendedProgram {
    ast: ESLintProgram;
    services?: Record<string, any>;
    visitorKeys?: {
        [type: string]: string[];
    };
    scopeManager?: ScopeManager;
    _virtualScriptCode?: string;
}
type ParseResult = {
    ast: SvelteProgram;
    services: Record<string, any> & ({
        isSvelte: true;
        isSvelteScript: false;
        getSvelteHtmlAst: () => SvAST.Fragment | Compiler.Fragment;
        getStyleContext: () => StyleContext;
        svelteParseContext: SvelteParseContext;
    } | {
        isSvelte: false;
        isSvelteScript: true;
        svelteParseContext: SvelteParseContext;
    });
    visitorKeys: {
        [type: string]: string[];
    };
    scopeManager: ScopeManager;
};
/**
 * Parse source code
 */
export declare function parseForESLint(code: string, options?: any): ParseResult;
