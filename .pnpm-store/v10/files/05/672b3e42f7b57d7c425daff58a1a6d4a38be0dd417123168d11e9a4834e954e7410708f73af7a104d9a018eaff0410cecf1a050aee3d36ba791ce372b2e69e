import type ESTree from "estree";
interface BaseNode {
    start: number;
    end: number;
}
export interface AstLegacy {
    html?: Fragment;
    css?: Style;
    instance?: Script;
    module?: Script;
}
export declare type TemplateNode = Text | MustacheTag | RawMustacheTag | DebugTag | ConstTag | Directive | StyleDirective | Element | InlineComponent | Window | Document | Body | Head | Title | Options | SlotTemplate | Slot | Comment | IfBlock | EachBlock | AwaitBlock | KeyBlock;
export interface Fragment extends BaseNode {
    type: "Fragment";
    children: TemplateNode[];
}
export interface Text extends BaseNode {
    type: "Text";
    data: string;
}
export interface MustacheTag extends BaseNode {
    type: "MustacheTag";
    expression: ESTree.Expression;
}
export interface RawMustacheTag extends BaseNode {
    type: "RawMustacheTag";
    expression: ESTree.Expression;
}
export interface DebugTag extends BaseNode {
    type: "DebugTag";
    identifiers: ESTree.Identifier[];
}
export interface ConstTag extends BaseNode {
    type: "ConstTag";
    expression: ESTree.AssignmentExpression;
}
export interface IfBlock extends BaseNode {
    type: "IfBlock";
    expression: ESTree.Expression;
    children: TemplateNode[];
    else: ElseBlock | undefined;
    elseif: true | undefined;
}
export interface ElseBlock extends BaseNode {
    type: "ElseBlock";
    children: TemplateNode[] | [IfBlock & {
        elseif: true;
    }];
}
export interface EachBlock extends BaseNode {
    type: "EachBlock";
    expression: ESTree.Expression;
    context: ESTree.Pattern;
    index?: string;
    key: ESTree.Expression | undefined;
    children: TemplateNode[];
    else: ElseBlock | undefined;
}
export interface AwaitBlock extends BaseNode {
    type: "AwaitBlock";
    expression: ESTree.Expression;
    pending: PendingBlock;
    value: ESTree.Pattern | null;
    then: ThenBlock;
    error: ESTree.Pattern | null;
    catch: CatchBlock;
}
export interface PendingBlock extends BaseNode {
    type: "PendingBlock";
    skip: boolean;
    children: TemplateNode[];
}
export interface ThenBlock extends BaseNode {
    type: "ThenBlock";
    skip: boolean;
    children: TemplateNode[];
}
export interface CatchBlock extends BaseNode {
    type: "CatchBlock";
    skip: boolean;
    children: TemplateNode[];
}
export interface KeyBlock extends BaseNode {
    type: "KeyBlock";
    expression: ESTree.Expression;
    children: TemplateNode[];
}
export interface BaseElement extends BaseNode {
    type: "Element";
    name: string;
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface BasicElement extends BaseElement {
    tag?: undefined;
}
export interface SvelteElement extends BaseElement {
    name: "svelte:element";
    tag: ESTree.Expression | string;
}
export type Element = BasicElement | SvelteElement;
export interface BaseInlineComponent extends BaseNode {
    type: "InlineComponent";
    name: string;
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Window extends BaseNode {
    type: "Window";
    name: "svelte:window";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Document extends BaseNode {
    type: "Document";
    name: "svelte:document";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Body extends BaseNode {
    type: "Body";
    name: "svelte:body";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Head extends BaseNode {
    type: "Head";
    name: "svelte:head";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Title extends BaseNode {
    type: "Title";
    name: "title";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Options extends BaseNode {
    type: "Options";
    name: "svelte:options";
    children?: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface SlotTemplate extends BaseNode {
    type: "SlotTemplate";
    name: "svelte:fragment";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface BasicInlineComponent extends BaseInlineComponent {
    expression?: undefined;
}
export interface InlineSvelteComponent extends BaseInlineComponent {
    name: "svelte:component";
    expression: ESTree.Expression;
}
export type InlineComponent = BasicInlineComponent | InlineSvelteComponent;
export interface Slot extends BaseNode {
    type: "Slot";
    name: "slot";
    children: TemplateNode[];
    attributes: AttributeOrDirective[];
}
export interface Comment extends BaseNode {
    type: "Comment";
    data: string;
}
export interface Attribute extends BaseNode {
    type: "Attribute";
    name: string;
    value: (Text | MustacheTag)[] | [AttributeShorthand] | true;
}
export interface Spread extends BaseNode {
    type: "Spread";
    expression: ESTree.Expression;
}
export interface AttributeShorthand extends BaseNode {
    type: "AttributeShorthand";
    expression: ESTree.Identifier;
}
export type AttributeOrDirective = Attribute | Spread | Directive | StyleDirective;
interface BaseDirective extends BaseNode {
    name: string;
    modifiers: string[];
}
export interface DirectiveForExpression extends BaseDirective {
    type: "Action" | "Animation" | "Binding" | "Class" | "EventHandler" | "Ref";
    expression: null | ESTree.Expression;
}
export interface LetDirective extends BaseDirective {
    type: "Let";
    expression: null | ESTree.Pattern;
}
export interface TransitionDirective extends BaseDirective {
    type: "Transition";
    intro: boolean;
    outro: boolean;
    expression: null | ESTree.Expression;
}
export interface StyleDirective extends BaseDirective {
    type: "StyleDirective";
    value: (Text | MustacheTag)[] | true;
}
export type Directive = DirectiveForExpression | TransitionDirective | LetDirective;
export interface Script extends BaseNode {
    type: "Script";
    context: string;
    children: TemplateNode[];
}
export interface Style extends BaseNode {
    type: "Style";
}
export {};
