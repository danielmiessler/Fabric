import type { ESLintExtendedProgram } from "..";
import type { NormalizedParserOptions } from "../parser-options";
import type { SvelteParseContext } from "../svelte-parse-context";
import type { AnalyzeTypeScriptContext } from "./analyze";
/**
 * Parse for TypeScript in <script>
 */
export declare function parseTypeScriptInSvelte(code: {
    script: string;
    render: string;
    rootScope: string;
}, attrs: Record<string, string | undefined>, parserOptions: NormalizedParserOptions, context: AnalyzeTypeScriptContext): ESLintExtendedProgram;
/**
 * Parse for TypeScript
 */
export declare function parseTypeScript(code: string, attrs: Record<string, string | undefined>, parserOptions: NormalizedParserOptions, svelteParseContext: SvelteParseContext): ESLintExtendedProgram;
