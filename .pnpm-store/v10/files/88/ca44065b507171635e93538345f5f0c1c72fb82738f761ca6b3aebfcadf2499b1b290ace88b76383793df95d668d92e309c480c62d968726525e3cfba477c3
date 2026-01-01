import { RestoreContext } from "./restore";
import type { TSESParseForESLintResult } from "./types";
/**
 * Context for virtual TypeScript code.
 * See https://github.com/sveltejs/svelte-eslint-parser/blob/main/docs/internal-mechanism.md#scope-types
 */
export declare class VirtualTypeScriptContext {
    private readonly originalCode;
    readonly restoreContext: RestoreContext;
    script: string;
    private consumedIndex;
    private readonly unique;
    _beforeResult: TSESParseForESLintResult | null;
    constructor(code: string);
    skipOriginalOffset(offset: number): void;
    skipUntilOriginalOffset(offset: number): void;
    appendOriginalToEnd(): void;
    appendOriginal(index: number): void;
    appendVirtualScript(virtualFragment: string): void;
    generateUniqueId(base: string): string;
}
