import type * as SvAST from "../svelte-ast-types";
import type * as Compiler from "../svelte-ast-types-for-v5";
import type { SvelteAwaitBlock, SvelteEachBlock, SvelteIfBlock, SvelteIfBlockAlone, SvelteIfBlockElseIf, SvelteKeyBlock, SvelteSnippetBlock } from "../../ast";
import type { Context } from "../../context";
export declare function convertIfBlock(node: SvAST.IfBlock | Compiler.IfBlock, parent: SvelteIfBlock["parent"], ctx: Context): SvelteIfBlockAlone;
export declare function convertIfBlock(node: SvAST.IfBlock | Compiler.IfBlock, parent: SvelteIfBlock["parent"], ctx: Context, elseifContext?: {
    start: number;
}): SvelteIfBlockElseIf;
/** Convert for EachBlock */
export declare function convertEachBlock(node: SvAST.EachBlock | Compiler.EachBlock, parent: SvelteEachBlock["parent"], ctx: Context): SvelteEachBlock;
/** Convert for AwaitBlock */
export declare function convertAwaitBlock(node: SvAST.AwaitBlock | Compiler.AwaitBlock, parent: SvelteAwaitBlock["parent"], ctx: Context): SvelteAwaitBlock;
/** Convert for KeyBlock */
export declare function convertKeyBlock(node: SvAST.KeyBlock | Compiler.KeyBlock, parent: SvelteKeyBlock["parent"], ctx: Context): SvelteKeyBlock;
/** Convert for SnippetBlock */
export declare function convertSnippetBlock(node: Compiler.SnippetBlock, parent: SvelteSnippetBlock["parent"], ctx: Context): SvelteSnippetBlock;
