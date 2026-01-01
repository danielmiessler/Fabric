import type * as Compiler from "./svelte-ast-types-for-v5";
import type * as SvAST from "./svelte-ast-types";
import type { NormalizedParserOptions } from "./parser-options";
import type { SvelteConfig } from "../svelte-config";
/** The context for parsing. */
export type SvelteParseContext = {
    /**
     * Whether to use Runes mode.
     * May be `true` if the user is using Svelte v5.
     * Resolved from `svelte.config.js` or `parserOptions`, but may be overridden by `<svelte:options>`.
     */
    runes: boolean;
    /** The version of "svelte/compiler". */
    compilerVersion: string;
    /** The result of static analysis of `svelte.config.js`. */
    svelteConfig: SvelteConfig | null;
};
export declare function isEnableRunes(svelteConfig: SvelteConfig | null, parserOptions: NormalizedParserOptions): boolean;
export declare function resolveSvelteParseContextForSvelte(svelteConfig: SvelteConfig | null, parserOptions: NormalizedParserOptions, svelteAst: Compiler.Root | SvAST.AstLegacy): SvelteParseContext;
export declare function resolveSvelteParseContextForSvelteScript(svelteConfig: SvelteConfig | null, parserOptions: NormalizedParserOptions): SvelteParseContext;
