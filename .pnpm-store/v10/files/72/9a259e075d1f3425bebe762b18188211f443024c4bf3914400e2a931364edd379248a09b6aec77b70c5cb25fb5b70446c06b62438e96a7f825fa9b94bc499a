import type { ScopeManager, Scope } from "eslint-scope";
import type * as ESTree from "estree";
import type { TSESTree } from "@typescript-eslint/types";
import type { Context } from ".";
import type { Comment, SvelteEachBlock, SvelteIfBlock, SvelteName, SvelteNode, SvelteSnippetBlock, Token } from "../ast";
import type { ESLintExtendedProgram } from "../parser";
export type ScriptLetCallback<E extends ESTree.Node> = (es: E, options: ScriptLetCallbackOption) => void;
export type ScriptLetCallbackOption = {
    getScope: (node: ESTree.Node) => Scope;
    registerNodeToScope: (node: any, scope: Scope) => void;
    scopeManager: ScopeManager;
    visitorKeys?: {
        [type: string]: string[];
    };
};
export type ScriptLetRestoreCallback = (node: ESTree.Node, tokens: Token[], comments: Comment[], options: ScriptLetRestoreCallbackOption) => void;
type ScriptLetRestoreCallbackOption = {
    getScope: (node: ESTree.Node) => Scope;
    registerNodeToScope: (node: any, scope: Scope) => void;
    scopeManager: ScopeManager;
    visitorKeys?: {
        [type: string]: string[];
    };
    addPostProcess: (callback: () => void) => void;
};
type StatementNodeType = `${TSESTree.Statement["type"]}`;
type TypeGenHelper = {
    generateUniqueId: (base: string) => string;
};
type ObjectShorthandProperty = ESTree.Property & {
    key: ESTree.Identifier;
    value: ESTree.Identifier;
};
export type ScriptLetBlockParam = {
    node: ESTree.Pattern | SvelteName;
    parent: SvelteNode;
    typing: string;
    callback: (node: ESTree.Pattern, options: ScriptLetCallbackOption) => void;
};
/**
 * A class that handles script fragments.
 * The script fragment AST node remaps and connects to the original directive AST node.
 */
export declare class ScriptLetContext {
    private readonly script;
    private readonly ctx;
    private readonly restoreCallbacks;
    private readonly programRestoreCallbacks;
    private readonly closeScopeCallbacks;
    private readonly unique;
    private currentScriptScopeKind;
    constructor(ctx: Context);
    addExpression<E extends ESTree.Expression>(expression: E | SvelteName, parent: SvelteNode, typing?: string | null, ...callbacks: ScriptLetCallback<E>[]): ScriptLetCallback<E>[];
    addExpressionFromRange<E extends ESTree.Expression>(range: [number, number], parent: SvelteNode, typing?: string | null, ...callbacks: ScriptLetCallback<E>[]): ScriptLetCallback<E>[];
    addObjectShorthandProperty(identifier: SvelteName, parent: SvelteNode, ...callbacks: ScriptLetCallback<ObjectShorthandProperty>[]): void;
    addVariableDeclarator(declarator: ESTree.VariableDeclarator | ESTree.AssignmentExpression, parent: SvelteNode, ...callbacks: ScriptLetCallback<ESTree.VariableDeclarator>[]): ScriptLetCallback<ESTree.VariableDeclarator>[];
    addGenericTypeAliasDeclaration(param: TSESTree.TSTypeParameter, callbackId: (id: TSESTree.Identifier, type: TSESTree.TypeNode) => void, callbackDefault: (type: TSESTree.TypeNode) => void): void;
    nestIfBlock(expression: ESTree.Expression, ifBlock: SvelteIfBlock, callback: ScriptLetCallback<ESTree.Expression>): void;
    nestEachBlock(expression: ESTree.Expression, context: ESTree.Pattern, indexRange: {
        start: number;
        end: number;
    } | null, eachBlock: SvelteEachBlock, callback: (expr: ESTree.Expression, ctx: ESTree.Pattern, index: ESTree.Identifier | null) => void): void;
    nestSnippetBlock(id: ESTree.Identifier, closeParentIndex: number, snippetBlock: SvelteSnippetBlock, kind: "snippet" | null, callback: (id: ESTree.Identifier, params: ESTree.Pattern[]) => void): void;
    nestBlock(block: SvelteNode, params?: ScriptLetBlockParam[] | ((helper: TypeGenHelper | null) => {
        param: ScriptLetBlockParam;
        preparationScript?: {
            script: string;
            nodeType: StatementNodeType;
        }[];
    })): void;
    closeScope(): void;
    addProgramRestore(callback: ScriptLetRestoreCallback): void;
    private appendScript;
    private appendScriptWithoutOffset;
    private pushScope;
    /**
     * Restore AST nodes
     */
    restore(result: ESLintExtendedProgram): void;
    /**
     * Restore AST nodes
     */
    private restoreNodes;
    /**
     * Restore program node
     */
    private restoreProgram;
    private remapNodes;
    private generateUniqueId;
}
export {};
