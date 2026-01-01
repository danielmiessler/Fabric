/** Compatibility for Svelte v4 <-> v5 */
import type ESTree from "estree";
import type * as SvAST from "./svelte-ast-types";
import type * as Compiler from "./svelte-ast-types-for-v5";
export type Child = Compiler.Text | Compiler.Tag | Compiler.ElementLike | Compiler.Block | Compiler.Comment;
type HasChildren = {
    children?: SvAST.TemplateNode[];
};
export declare function getFragmentFromRoot(svelteAst: Compiler.Root | SvAST.AstLegacy): SvAST.Fragment | Compiler.Fragment | undefined;
export declare function getInstanceFromRoot(svelteAst: Compiler.Root | SvAST.AstLegacy): SvAST.Script | Compiler.Script | null | undefined;
export declare function getModuleFromRoot(svelteAst: Compiler.Root | SvAST.AstLegacy): SvAST.Script | Compiler.Script | null | undefined;
export declare function getOptionsFromRoot(svelteAst: Compiler.Root | SvAST.AstLegacy): Compiler.SvelteOptionsRaw | null;
export declare function getChildren(fragment: Required<HasChildren> | {
    nodes: (Child | SvAST.TemplateNode)[];
}): (SvAST.TemplateNode | Child)[];
export declare function getChildren(fragment: HasChildren | {
    nodes: (Child | SvAST.TemplateNode)[];
}): (SvAST.TemplateNode | Child)[] | undefined;
export declare function trimChildren(children: (SvAST.TemplateNode | Child)[]): (SvAST.TemplateNode | Child)[];
export declare function getFragment(element: {
    fragment: Compiler.Fragment;
} | Required<HasChildren>): Compiler.Fragment | Required<HasChildren>;
export declare function getFragment(element: {
    fragment: Compiler.Fragment;
} | HasChildren): Compiler.Fragment | HasChildren;
export declare function getModifiers(node: SvAST.Directive | SvAST.StyleDirective | Compiler.Directive): string[];
export declare function getTestFromIfBlock(block: SvAST.IfBlock | Compiler.IfBlock): ESTree.Expression;
export declare function getConsequentFromIfBlock(block: SvAST.IfBlock | Compiler.IfBlock): Compiler.Fragment | SvAST.IfBlock;
export declare function getAlternateFromIfBlock(block: SvAST.IfBlock | Compiler.IfBlock): Compiler.Fragment | SvAST.ElseBlock | null;
export declare function getBodyFromEachBlock(block: SvAST.EachBlock | Compiler.EachBlock): Compiler.Fragment | SvAST.EachBlock;
export declare function getFallbackFromEachBlock(block: SvAST.EachBlock | Compiler.EachBlock): Compiler.Fragment | SvAST.ElseBlock | null;
export declare function getPendingFromAwaitBlock(block: SvAST.AwaitBlock | Compiler.AwaitBlock): Compiler.Fragment | SvAST.PendingBlock | null;
export declare function getThenFromAwaitBlock(block: SvAST.AwaitBlock | Compiler.AwaitBlock): Compiler.Fragment | SvAST.ThenBlock | null;
export declare function getCatchFromAwaitBlock(block: SvAST.AwaitBlock | Compiler.AwaitBlock): Compiler.Fragment | SvAST.CatchBlock | null;
export declare function getDeclaratorFromConstTag(node: SvAST.ConstTag | Compiler.ConstTag): ESTree.AssignmentExpression | Compiler.ConstTag["declaration"]["declarations"][0];
export {};
