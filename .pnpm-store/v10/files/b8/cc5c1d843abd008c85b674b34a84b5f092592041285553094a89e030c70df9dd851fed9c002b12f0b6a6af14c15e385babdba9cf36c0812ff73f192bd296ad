import type { AST } from 'svelte-eslint-parser';
import type { RuleContext } from '../../types';
import type { TSESTree } from '@typescript-eslint/types';
/**
 * Parse style attribute value
 */
export declare function parseStyleAttributeValue(node: AST.SvelteAttribute, context: RuleContext): SvelteStyleRoot<AST.SvelteMustacheTagText> | null;
export type SvelteStyleInterpolation = AST.SvelteMustacheTagText | TSESTree.Expression;
export interface SvelteStyleNode<E extends SvelteStyleInterpolation> {
    nodes?: SvelteStyleChildNode<E>[];
    range: AST.Range;
    loc: AST.SourceLocation;
}
export interface SvelteStyleRoot<E extends SvelteStyleInterpolation = SvelteStyleInterpolation> {
    type: 'root';
    nodes: (SvelteStyleChildNode<E> | SvelteStyleInline<E>)[];
}
export interface SvelteStyleInline<E extends SvelteStyleInterpolation = SvelteStyleInterpolation> extends SvelteStyleNode<E> {
    type: 'inline';
    node: E;
    getInlineStyle(node: TSESTree.Expression): SvelteStyleRoot<TSESTree.Expression> | null;
    getAllInlineStyles(): Map<TSESTree.Expression, SvelteStyleRoot<TSESTree.Expression>>;
}
export interface SvelteStyleDeclaration<E extends SvelteStyleInterpolation = SvelteStyleInterpolation> extends SvelteStyleNode<E> {
    type: 'decl';
    prop: {
        name: string;
        range: AST.Range;
        loc: AST.SourceLocation;
        interpolations: E[];
    };
    value: {
        value: string;
        range: AST.Range;
        loc: AST.SourceLocation;
        interpolations: E[];
    };
    important: boolean;
    addInterpolation: (tagOrExpr: E) => void;
    unknownInterpolations: E[];
}
export interface SvelteStyleComment extends SvelteStyleNode<never> {
    type: 'comment';
    addInterpolation: (tagOrExpr: SvelteStyleInterpolation) => void;
}
export type SvelteStyleChildNode<E extends SvelteStyleInterpolation = SvelteStyleInterpolation> = SvelteStyleDeclaration<E> | SvelteStyleComment;
