import type { SvelteAwaitBlock, SvelteAwaitCatchBlock, SvelteAwaitPendingBlock, SvelteAwaitThenBlock, SvelteConstTag, SvelteDebugTag, SvelteEachBlock, SvelteElement, SvelteElseBlockAlone, SvelteHTMLComment, SvelteIfBlock, SvelteIfBlockAlone, SvelteKeyBlock, SvelteMustacheTag, SvelteProgram, SvelteRenderTag, SvelteScriptElement, SvelteSnippetBlock, SvelteStyleElement, SvelteText } from "../../ast";
import type { Context } from "../../context";
import type * as SvAST from "../svelte-ast-types";
import type * as Compiler from "../svelte-ast-types-for-v5";
import type { Child } from "../compat";
/** Convert for Fragment or Element or ... */
export declare function convertChildren(fragment: {
    children?: SvAST.TemplateNode[];
} | Compiler.Fragment | {
    nodes: (Child | SvAST.TemplateNode)[];
}, parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock, ctx: Context): IterableIterator<SvelteText | SvelteElement | SvelteMustacheTag | SvelteDebugTag | SvelteConstTag | SvelteRenderTag | SvelteIfBlockAlone | SvelteEachBlock | SvelteAwaitBlock | SvelteKeyBlock | SvelteSnippetBlock | SvelteHTMLComment>;
/** Extract element tag and tokens */
export declare function extractElementTags<E extends SvelteScriptElement | SvelteElement | SvelteStyleElement>(element: E, ctx: Context, options: {
    buildNameNode: (openTokenRange: {
        start: number;
        end: number;
    }) => E["name"];
    extractAttribute?: boolean;
}): void;
