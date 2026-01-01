import { VirtualTypeScriptContext } from "../context";
import type { SvelteHTMLElement } from "../../../ast";
import type { NormalizedParserOptions } from "../../parser-options";
import type { SvelteParseContext } from "../../svelte-parse-context";
export type AnalyzeTypeScriptContext = {
    slots: Set<SvelteHTMLElement>;
    svelteParseContext: SvelteParseContext;
};
/**
 * Analyze TypeScript source code in <script>.
 * Generate virtual code to provide correct type information for Svelte store reference names, scopes, and runes.
 * See https://github.com/sveltejs/svelte-eslint-parser/blob/main/docs/internal-mechanism.md#scope-types
 */
export declare function analyzeTypeScriptInSvelte(code: {
    script: string;
    rootScope: string;
    render: string;
}, attrs: Record<string, string | undefined>, parserOptions: NormalizedParserOptions, context: AnalyzeTypeScriptContext): VirtualTypeScriptContext;
/**
 * Analyze TypeScript source code.
 * Generate virtual code to provide correct type information for Svelte runes.
 * See https://github.com/sveltejs/svelte-eslint-parser/blob/main/docs/internal-mechanism.md#scope-types
 */
export declare function analyzeTypeScript(code: string, attrs: Record<string, string | undefined>, parserOptions: NormalizedParserOptions, svelteParseContext: SvelteParseContext): VirtualTypeScriptContext;
