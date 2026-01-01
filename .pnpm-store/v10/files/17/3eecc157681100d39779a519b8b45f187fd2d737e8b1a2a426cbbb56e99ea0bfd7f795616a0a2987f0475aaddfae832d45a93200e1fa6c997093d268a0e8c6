import type ESTree from "estree";
import type { TSESTree } from "@typescript-eslint/types";
import type { BaseNode } from "./base";
import type { Token, Comment } from "./common";
export type SvelteHTMLNode = SvelteProgram | SvelteScriptElement | SvelteStyleElement | SvelteElement | SvelteStartTag | SvelteEndTag | SvelteName | SvelteMemberExpressionName | SvelteText | SvelteLiteral | SvelteMustacheTag | SvelteDebugTag | SvelteConstTag | SvelteRenderTag | SvelteIfBlock | SvelteElseBlock | SvelteEachBlock | SvelteAwaitBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock | SvelteAttribute | SvelteShorthandAttribute | SvelteSpreadAttribute | SvelteDirective | SvelteStyleDirective | SvelteSpecialDirective | SvelteGenericsDirective | SvelteDirectiveKey | SvelteSpecialDirectiveKey | SvelteHTMLComment;
/** Node of Svelte program root */
export interface SvelteProgram extends BaseNode {
    type: "Program";
    body: (SvelteScriptElement | SvelteStyleElement | Child)[];
    sourceType: "script" | "module";
    comments: Comment[];
    tokens: Token[];
    parent: null;
}
/** Node of elements like HTML element. */
export type SvelteElement = SvelteHTMLElement | SvelteComponentElement | SvelteSpecialElement;
type BaseSvelteElement = BaseNode;
/** Node of `<script>` element. */
export interface SvelteScriptElement extends BaseSvelteElement {
    type: "SvelteScriptElement";
    name: SvelteName;
    startTag: SvelteStartTag;
    body: ESTree.Program["body"];
    endTag: SvelteEndTag | null;
    parent: SvelteProgram;
}
/** Node of `<style>` element. */
export interface SvelteStyleElement extends BaseSvelteElement {
    type: "SvelteStyleElement";
    name: SvelteName;
    startTag: SvelteStartTag;
    children: [SvelteText];
    endTag: SvelteEndTag | null;
    parent: SvelteProgram;
}
/** Node of HTML element. */
export interface SvelteHTMLElement extends BaseSvelteElement {
    type: "SvelteElement";
    kind: "html";
    name: SvelteName;
    startTag: SvelteStartTag;
    children: Child[];
    endTag: SvelteEndTag | null;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of Svelte component element. */
export interface SvelteComponentElement extends BaseSvelteElement {
    type: "SvelteElement";
    kind: "component";
    name: ESTree.Identifier | SvelteMemberExpressionName;
    startTag: SvelteStartTag;
    children: Child[];
    endTag: SvelteEndTag | null;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of Svelte special component element. e.g. `<svelte:window>` */
export interface SvelteSpecialElement extends BaseSvelteElement {
    type: "SvelteElement";
    kind: "special";
    name: SvelteName;
    startTag: SvelteStartTag;
    children: Child[];
    endTag: SvelteEndTag | null;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of start tag. */
export interface SvelteStartTag extends BaseNode {
    type: "SvelteStartTag";
    attributes: (SvelteAttribute | SvelteShorthandAttribute | SvelteSpreadAttribute | SvelteDirective | SvelteStyleDirective | SvelteSpecialDirective | SvelteGenericsDirective)[];
    selfClosing: boolean;
    parent: SvelteElement | SvelteScriptElement | SvelteStyleElement;
}
/** Node of end tag. */
export interface SvelteEndTag extends BaseNode {
    type: "SvelteEndTag";
    parent: SvelteElement | SvelteScriptElement | SvelteStyleElement;
}
/** Node of names. It is used for element names other than components and normal attribute names. */
export interface SvelteName extends BaseNode {
    type: "SvelteName";
    name: string;
    parent: SvelteElement | SvelteScriptElement | SvelteStyleElement | SvelteAttribute | SvelteMemberExpressionName | SvelteDirectiveKey;
}
/** Nodes that may be used in component names. The component names separated by dots. */
export interface SvelteMemberExpressionName extends BaseNode {
    type: "SvelteMemberExpressionName";
    object: SvelteMemberExpressionName | ESTree.Identifier;
    property: SvelteName;
    parent: SvelteComponentElement;
}
type Child = SvelteElement | SvelteText | SvelteMustacheTag | SvelteDebugTag | SvelteConstTag | SvelteRenderTag | SvelteIfBlockAlone | SvelteEachBlock | SvelteAwaitBlock | SvelteKeyBlock | SvelteSnippetBlock | SvelteHTMLComment;
/** Node of text. like HTML text. */
export interface SvelteText extends BaseNode {
    type: "SvelteText";
    value: string;
    parent: SvelteProgram | SvelteElement | SvelteStyleElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of literal. */
export interface SvelteLiteral extends BaseNode {
    type: "SvelteLiteral";
    value: string;
    parent: SvelteAttribute | SvelteStyleDirective;
}
/** Node of mustache tag. e.g. `{...}`, `{@html ...}`. Like JSXExpressionContainer */
export type SvelteMustacheTag = SvelteMustacheTagText | SvelteMustacheTagRaw;
interface BaseSvelteMustacheTag extends BaseNode {
    type: "SvelteMustacheTag";
    kind: "text" | "raw";
    expression: ESTree.Expression;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock | SvelteAttribute | SvelteStyleDirective;
}
/** Node of mustache tag. e.g. `{...}``. Like JSXExpressionContainer */
export interface SvelteMustacheTagText extends BaseSvelteMustacheTag {
    kind: "text";
}
/** Node of mustache tag. e.g. `{@html ...}`. Like JSXExpressionContainer */
export interface SvelteMustacheTagRaw extends BaseSvelteMustacheTag {
    kind: "raw";
}
/** Node of debug mustache tag. e.g. `{@debug}` */
export interface SvelteDebugTag extends BaseNode {
    type: "SvelteDebugTag";
    identifiers: ESTree.Identifier[];
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock | SvelteAttribute;
}
/** Node of const tag. e.g. `{@const}` */
export interface SvelteConstTag extends BaseNode {
    type: "SvelteConstTag";
    declaration: ESTree.VariableDeclarator;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock | SvelteAttribute;
}
/** Node of render tag. e.g. `{@render}` */
export interface SvelteRenderTag extends BaseNode {
    type: "SvelteRenderTag";
    expression: ESTree.SimpleCallExpression | (ESTree.ChainExpression & {
        expression: ESTree.SimpleCallExpression;
    });
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of if block. e.g. `{#if}` */
export type SvelteIfBlock = SvelteIfBlockAlone | SvelteIfBlockElseIf;
interface BaseSvelteIfBlock extends BaseNode {
    type: "SvelteIfBlock";
    elseif: boolean;
    expression: ESTree.Expression;
    children: Child[];
    else: SvelteElseBlock | null;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlock | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of if block. e.g. `{#if}` */
export interface SvelteIfBlockAlone extends BaseSvelteIfBlock {
    elseif: false;
    parent: Exclude<BaseSvelteIfBlock["parent"], SvelteElseBlockElseIf>;
}
/** Node of if block. e.g. `{:else if}` */
export interface SvelteIfBlockElseIf extends BaseSvelteIfBlock {
    elseif: true;
    parent: SvelteElseBlockElseIf;
}
/** Node of else block. e.g. `{:else}` */
export type SvelteElseBlock = SvelteElseBlockAlone | SvelteElseBlockElseIf;
interface BaseSvelteElseBlock extends BaseNode {
    type: "SvelteElseBlock";
    elseif: boolean;
    children: (Child | SvelteIfBlockElseIf)[];
    parent: SvelteIfBlock | SvelteEachBlock;
}
/** Node of else block. e.g. `{:else}` */
export interface SvelteElseBlockAlone extends BaseSvelteElseBlock {
    elseif: false;
    children: Child[];
}
/** Node of else block. e.g. `{:else if ...}` */
export interface SvelteElseBlockElseIf extends BaseSvelteElseBlock {
    elseif: true;
    children: [SvelteIfBlockElseIf];
}
/** Node of each block. e.g. `{#each}` */
export interface SvelteEachBlock extends BaseNode {
    type: "SvelteEachBlock";
    expression: ESTree.Expression;
    context: ESTree.Pattern;
    index: ESTree.Identifier | null;
    key: ESTree.Expression | null;
    children: Child[];
    else: SvelteElseBlockAlone | null;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of await block. e.g. `{#await}`, `{#await ... then ... }`, `{#await ... catch ... }` */
export type SvelteAwaitBlock = SvelteAwaitBlockAwaitPending | SvelteAwaitBlockAwaitThen | SvelteAwaitBlockAwaitCatch;
interface BaseSvelteAwaitBlock extends BaseNode {
    type: "SvelteAwaitBlock";
    expression: ESTree.Expression;
    kind: "await" | "await-then" | "await-catch";
    pending: SvelteAwaitPendingBlock | null;
    then: SvelteAwaitThenBlock | null;
    catch: SvelteAwaitCatchBlock | null;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of await block. e.g. `{#await}` */
export interface SvelteAwaitBlockAwaitPending extends BaseSvelteAwaitBlock {
    kind: "await";
    pending: SvelteAwaitPendingBlock;
    then: SvelteAwaitThenBlockAlone | null;
    catch: SvelteAwaitCatchBlockAlone | null;
}
/** Node of await block. e.g. `{#await ... then ... }` */
export interface SvelteAwaitBlockAwaitThen extends BaseSvelteAwaitBlock {
    kind: "await-then";
    pending: null;
    then: SvelteAwaitThenBlockAwaitThen;
    catch: SvelteAwaitCatchBlockAlone | null;
}
/** Node of await block. e.g. `{#await ... catch ... }` */
export interface SvelteAwaitBlockAwaitCatch extends BaseSvelteAwaitBlock {
    kind: "await-catch";
    pending: null;
    then: null;
    catch: SvelteAwaitCatchBlockAwaitCatch;
}
/** Node of await pending block. e.g. `{#await expr} ... {:then}` */
export interface SvelteAwaitPendingBlock extends BaseNode {
    type: "SvelteAwaitPendingBlock";
    children: Child[];
    parent: SvelteAwaitBlock;
}
/** Node of await then block. e.g. `{:then}`, `{#await ... then ...}` */
export type SvelteAwaitThenBlock = SvelteAwaitThenBlockAlone | SvelteAwaitThenBlockAwaitThen;
interface BaseSvelteAwaitThenBlock extends BaseNode {
    type: "SvelteAwaitThenBlock";
    awaitThen: boolean;
    value: ESTree.Pattern | null;
    children: Child[];
    parent: SvelteAwaitBlock;
}
/** Node of await then block. e.g. `{:then}` */
export interface SvelteAwaitThenBlockAlone extends BaseSvelteAwaitThenBlock {
    awaitThen: false;
    parent: SvelteAwaitBlockAwaitPending;
}
/** Node of await then block. e.g. `{#await ... then ...}` */
export interface SvelteAwaitThenBlockAwaitThen extends BaseSvelteAwaitThenBlock {
    awaitThen: true;
    parent: SvelteAwaitBlockAwaitThen;
}
/** Node of await catch block. e.g. `{:catch}`, `{#await ... catch ... }` */
export type SvelteAwaitCatchBlock = SvelteAwaitCatchBlockAlone | SvelteAwaitCatchBlockAwaitCatch;
interface BaseSvelteAwaitCatchBlock extends BaseNode {
    type: "SvelteAwaitCatchBlock";
    awaitCatch: boolean;
    error: ESTree.Pattern | null;
    children: Child[];
    parent: SvelteAwaitBlock;
}
/** Node of await catch block. e.g. `{:catch}` */
export interface SvelteAwaitCatchBlockAlone extends BaseSvelteAwaitCatchBlock {
    awaitCatch: false;
    parent: SvelteAwaitBlockAwaitPending | SvelteAwaitBlockAwaitThen;
}
/** Node of await catch block. e.g. `{#await ... catch ... }` */
export interface SvelteAwaitCatchBlockAwaitCatch extends BaseSvelteAwaitCatchBlock {
    awaitCatch: true;
    error: ESTree.Pattern | null;
    children: Child[];
    parent: SvelteAwaitBlockAwaitCatch;
}
/** Node of key block. e.g. `{#key}` */
export interface SvelteKeyBlock extends BaseNode {
    type: "SvelteKeyBlock";
    expression: ESTree.Expression;
    children: Child[];
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of snippet block. e.g. `{#snippet}` */
export interface SvelteSnippetBlock extends BaseNode {
    type: "SvelteSnippetBlock";
    id: ESTree.Identifier;
    params: ESTree.Pattern[];
    children: Child[];
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of HTML comment. */
export interface SvelteHTMLComment extends BaseNode {
    type: "SvelteHTMLComment";
    value: string;
    parent: SvelteProgram | SvelteElement | SvelteIfBlock | SvelteElseBlockAlone | SvelteEachBlock | SvelteAwaitPendingBlock | SvelteAwaitThenBlock | SvelteAwaitCatchBlock | SvelteKeyBlock | SvelteSnippetBlock;
}
/** Node of HTML attribute. */
export interface SvelteAttribute extends BaseNode {
    type: "SvelteAttribute";
    key: SvelteName;
    boolean: boolean;
    value: (SvelteLiteral | SvelteMustacheTagText)[];
    parent: SvelteStartTag;
}
/** Node of shorthand attribute. e.g. `<img {src}>` */
export interface SvelteShorthandAttribute extends BaseNode {
    type: "SvelteShorthandAttribute";
    key: ESTree.Identifier;
    value: ESTree.Identifier;
    parent: SvelteStartTag;
}
/** Node of spread attribute. e.g. `<Info {...pkg}/>`. Like JSXSpreadAttribute */
export interface SvelteSpreadAttribute extends BaseNode {
    type: "SvelteSpreadAttribute";
    argument: ESTree.Expression;
    parent: SvelteStartTag;
}
/** Node of directive. e.g. `<input bind:value />` */
export type SvelteDirective = SvelteActionDirective | SvelteAnimationDirective | SvelteBindingDirective | SvelteClassDirective | SvelteEventHandlerDirective | SvelteLetDirective | SvelteRefDirective | SvelteTransitionDirective;
export type SvelteDirectiveKey = SvelteDirectiveKeyTextName | SvelteDirectiveKeyFunctionName | SvelteDirectiveKeyForEventHandler | SvelteDirectiveKeyForAction | SvelteDirectiveKeyForStyleShorthand;
interface BaseSvelteDirectiveKey<N extends ESTree.Expression | SvelteName> extends BaseNode {
    type: "SvelteDirectiveKey";
    name: N;
    modifiers: string[];
    parent: SvelteDirective | SvelteStyleDirective;
}
export type SvelteDirectiveKeyTextName = BaseSvelteDirectiveKey<SvelteName>;
export type SvelteDirectiveKeyFunctionName = BaseSvelteDirectiveKey<ESTree.Identifier>;
export type SvelteDirectiveKeyForEventHandler = BaseSvelteDirectiveKey<SvelteName>;
export type SvelteDirectiveKeyForAction = BaseSvelteDirectiveKey<ESTree.Identifier | ESTree.MemberExpression | SvelteName>;
export type SvelteDirectiveKeyForStyleShorthand = BaseSvelteDirectiveKey<ESTree.Identifier>;
interface BaseSvelteDirective extends BaseNode {
    type: "SvelteDirective";
    key: SvelteDirectiveKey;
    parent: SvelteStartTag;
}
export interface SvelteActionDirective extends BaseSvelteDirective {
    kind: "Action";
    key: SvelteDirectiveKeyForAction;
    expression: null | ESTree.Expression;
}
export interface SvelteAnimationDirective extends BaseSvelteDirective {
    kind: "Animation";
    key: SvelteDirectiveKeyFunctionName;
    expression: null | ESTree.Expression;
}
export interface SvelteBindingDirective extends BaseSvelteDirective {
    kind: "Binding";
    key: SvelteDirectiveKeyTextName;
    shorthand: boolean;
    expression: null | ESTree.Expression;
}
export interface SvelteClassDirective extends BaseSvelteDirective {
    kind: "Class";
    key: SvelteDirectiveKeyTextName;
    shorthand: boolean;
    expression: null | ESTree.Expression;
}
export interface SvelteEventHandlerDirective extends BaseSvelteDirective {
    kind: "EventHandler";
    key: SvelteDirectiveKeyForEventHandler;
    expression: null | ESTree.Expression;
}
export interface SvelteLetDirective extends BaseSvelteDirective {
    kind: "Let";
    key: SvelteDirectiveKeyTextName;
    expression: null | ESTree.Pattern;
}
export interface SvelteRefDirective extends BaseSvelteDirective {
    kind: "Ref";
    key: SvelteDirectiveKeyTextName;
    expression: null | ESTree.Expression;
}
export interface SvelteTransitionDirective extends BaseSvelteDirective {
    kind: "Transition";
    key: SvelteDirectiveKeyFunctionName;
    intro: boolean;
    outro: boolean;
    expression: null | ESTree.Expression;
}
/** Node of style directive. e.g. `<input style:color />` */
export type SvelteStyleDirective = SvelteStyleDirectiveShorthand | SvelteStyleDirectiveLongform;
interface BaseSvelteStyleDirective extends BaseNode {
    type: "SvelteStyleDirective";
    key: SvelteDirectiveKeyTextName | SvelteDirectiveKeyForStyleShorthand;
    value: (SvelteLiteral | SvelteMustacheTagText)[];
    parent: SvelteStartTag;
}
export interface SvelteStyleDirectiveShorthand extends BaseSvelteStyleDirective {
    key: SvelteDirectiveKeyForStyleShorthand;
    shorthand: true;
    value: [];
}
export interface SvelteStyleDirectiveLongform extends BaseSvelteStyleDirective {
    key: SvelteDirectiveKeyTextName;
    shorthand: false;
    value: (SvelteLiteral | SvelteMustacheTagText)[];
}
export interface SvelteSpecialDirectiveKey extends BaseNode {
    type: "SvelteSpecialDirectiveKey";
    parent: SvelteSpecialDirective;
}
export interface SvelteSpecialDirective extends BaseNode {
    type: "SvelteSpecialDirective";
    kind: "this";
    key: SvelteSpecialDirectiveKey;
    expression: ESTree.Expression;
    parent: SvelteStartTag;
}
export interface SvelteGenericsDirective extends BaseNode {
    type: "SvelteGenericsDirective";
    key: SvelteName & {
        name: "generics";
    };
    params: TSESTree.TSTypeParameterDeclaration["params"];
    parent: SvelteStartTag;
}
export {};
