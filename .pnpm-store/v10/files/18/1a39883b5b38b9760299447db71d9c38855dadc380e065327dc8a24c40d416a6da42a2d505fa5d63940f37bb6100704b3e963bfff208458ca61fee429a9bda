import type { JSONSchema4 } from 'json-schema';
import type { Linter, Rule, SourceCode as ESLintSourceCode } from 'eslint';
import type { AST, StyleContext, SvelteConfig } from 'svelte-eslint-parser';
import type { TSESTree } from '@typescript-eslint/types';
import type { ScopeManager, Scope, Variable } from '@typescript-eslint/scope-manager';
import type { ASTNode, ASTNodeWithParent, ASTNodeListener } from './types-for-node';
import type * as TS from 'typescript';
export type { ASTNode, ASTNodeWithParent, ASTNodeListener };
export interface RuleListener extends ASTNodeListener {
    onCodePathStart?(codePath: Rule.CodePath, node: never): void;
    onCodePathEnd?(codePath: Rule.CodePath, node: never): void;
    onCodePathSegmentStart?(segment: Rule.CodePathSegment, node: never): void;
    onCodePathSegmentEnd?(segment: Rule.CodePathSegment, node: never): void;
    onCodePathSegmentLoop?(fromSegment: Rule.CodePathSegment, toSegment: Rule.CodePathSegment, node: never): void;
    [key: string]: ((codePath: Rule.CodePath, node: never) => void) | ((segment: Rule.CodePathSegment, node: never) => void) | ((fromSegment: Rule.CodePathSegment, toSegment: Rule.CodePathSegment, node: never) => void) | ASTNodeListener[keyof ASTNodeListener] | ((node: never) => void) | undefined;
}
export interface RuleModule {
    meta: RuleMetaData;
    create: (context: RuleContext) => RuleListener;
}
export type RuleCategory = 'Possible Errors' | 'Security Vulnerability' | 'Best Practices' | 'Stylistic Issues' | 'Extension Rules' | 'SvelteKit' | 'Experimental' | 'System';
export interface RuleMetaData {
    docs: {
        description: string;
        category: RuleCategory;
        recommended: boolean | 'base';
        extensionRule?: string | {
            plugin: string;
            url: string;
        };
        url: string;
        ruleId: string;
        ruleName: string;
        default?: 'error' | 'warn';
        conflictWithPrettier?: boolean;
    };
    messages: {
        [messageId: string]: string;
    };
    fixable?: 'code' | 'whitespace';
    hasSuggestions?: boolean;
    schema: JSONSchema4 | JSONSchema4[];
    deprecated?: boolean;
    replacedBy?: string[] | {
        note: string;
    };
    type: 'problem' | 'suggestion' | 'layout';
}
export interface PartialRuleModule {
    meta: PartialRuleMetaData;
    create: (context: RuleContext) => RuleListener;
}
export interface PartialRuleMetaData {
    docs: {
        description: string;
        recommended: boolean | 'base';
        extensionRule?: string | {
            plugin: string;
            url: string;
        };
        default?: 'error' | 'warn';
    } & ({
        category: Exclude<RuleCategory, 'Stylistic Issues'>;
        conflictWithPrettier?: boolean;
    } | {
        category: 'Stylistic Issues';
        conflictWithPrettier: boolean;
    });
    messages: {
        [messageId: string]: string;
    };
    fixable?: 'code' | 'whitespace';
    hasSuggestions?: boolean;
    schema: JSONSchema4 | JSONSchema4[];
    deprecated?: boolean;
    replacedBy?: string[] | {
        note: string;
    };
    type: 'problem' | 'suggestion' | 'layout';
}
export type RuleContext = {
    id: string;
    options: any[];
    settings?: {
        svelte?: {
            ignoreWarnings?: unknown;
            compileOptions?: {
                babel?: boolean;
                postcss?: false | {
                    configFilePath?: unknown;
                };
            };
            kit?: {
                files?: {
                    routes?: string;
                };
            };
        };
    };
    parserPath: string;
    parserOptions: Linter.ParserOptions;
    parserServices: ESLintSourceCode.ParserServices;
    getAncestors(): ASTNode[];
    getDeclaredVariables(node: TSESTree.Node): Variable[];
    getFilename(): string;
    getScope(): Scope;
    getSourceCode(): SourceCode;
    markVariableAsUsed(name: string): boolean;
    report(descriptor: ReportDescriptor): void;
    getCwd?: () => string;
    getPhysicalFilename?: () => string;
};
export type NodeOrToken = {
    type: string;
    loc?: AST.SourceLocation | null;
    range?: [number, number];
};
interface ReportDescriptorOptionsBase {
    data?: {
        [key: string]: string;
    };
    fix?: null | ((fixer: RuleFixer) => null | Rule.Fix | IterableIterator<Rule.Fix> | Rule.Fix[]);
}
type SuggestionDescriptorMessage = {
    desc: string;
} | {
    messageId: string;
};
export type SuggestionReportDescriptor = SuggestionDescriptorMessage & ReportDescriptorOptionsBase;
interface ReportDescriptorOptions extends ReportDescriptorOptionsBase {
    suggest?: SuggestionReportDescriptor[] | null;
}
type ReportDescriptor = ReportDescriptorMessage & ReportDescriptorLocation & ReportDescriptorOptions;
type ReportDescriptorMessage = {
    message: string;
} | {
    messageId: string;
};
type ReportDescriptorLocation = {
    node: NodeOrToken;
} | {
    loc: AST.SourceLocation | {
        line: number;
        column: number;
    };
};
export interface RuleFixer {
    insertTextAfter(nodeOrToken: NodeOrToken, text: string): Rule.Fix;
    insertTextAfterRange(range: AST.Range, text: string): Rule.Fix;
    insertTextBefore(nodeOrToken: NodeOrToken, text: string): Rule.Fix;
    insertTextBeforeRange(range: AST.Range, text: string): Rule.Fix;
    remove(nodeOrToken: NodeOrToken): Rule.Fix;
    removeRange(range: AST.Range): Rule.Fix;
    replaceText(nodeOrToken: NodeOrToken, text: string): Rule.Fix;
    replaceTextRange(range: AST.Range, text: string): Rule.Fix;
}
export declare namespace SourceCode {
    function splitLines(text: string): string[];
}
export interface SourceCode {
    text: string;
    ast: AST.SvelteProgram;
    lines: string[];
    hasBOM: boolean;
    parserServices: {
        isSvelte?: boolean;
        isSvelteScript?: boolean;
        getSvelteHtmlAst?: () => unknown;
        getStyleContext?: () => StyleContext;
        svelteParseContext?: {
            /**
             * Whether to use Runes mode.
             * May be `true` if the user is using Svelte v5.
             * Resolved from `svelte.config.js` or `parserOptions`, but may be overridden by `<svelte:options>`.
             */
            runes?: boolean;
            /** The version of "svelte/compiler". */
            compilerVersion?: string;
            /** The result of static analysis of `svelte.config.js`. */
            svelteConfig?: SvelteConfig | null;
        };
        program?: TS.Program;
        esTreeNodeToTSNodeMap?: ReadonlyMap<unknown, TS.Node>;
        tsNodeToESTreeNodeMap?: ReadonlyMap<TS.Node, ASTNode>;
        hasFullTypeInformation?: boolean;
        [key: string]: unknown;
    };
    scopeManager: ScopeManager;
    visitorKeys: ESLintSourceCode.VisitorKeys;
    getText(node?: NodeOrToken, beforeCount?: number, afterCount?: number): string;
    getLines(): string[];
    getAllComments(): AST.Comment[];
    getComments(node: NodeOrToken): {
        leading: AST.Comment[];
        trailing: AST.Comment[];
    };
    getJSDocComment(node: NodeOrToken): AST.Comment | null;
    getNodeByRangeIndex(index: number): ASTNodeWithParent | null;
    isSpaceBetweenTokens(first: AST.Token, second: AST.Token): boolean;
    getLocFromIndex(index: number): AST.Position;
    getIndexFromLoc(location: AST.Position): number;
    getTokenByRangeStart(offset: number, options?: {
        includeComments?: boolean;
    }): AST.Token | AST.Comment | null;
    getFirstToken(node: NodeOrToken): AST.Token;
    getFirstToken(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getFirstToken']>[1]): AST.Token | AST.Comment | null;
    getFirstTokens(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getFirstTokens']>[1]): (AST.Token | AST.Comment)[];
    getLastToken(node: NodeOrToken): AST.Token;
    getLastToken(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getLastToken']>[1]): AST.Token | AST.Comment | null;
    getLastTokens(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getLastTokens']>[1]): (AST.Token | AST.Comment)[];
    getTokenBefore(node: NodeOrToken): AST.Token | null;
    getTokenBefore(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getTokenBefore']>[1]): AST.Token | AST.Comment | null;
    getTokensBefore(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getTokensBefore']>[1]): (AST.Token | AST.Comment)[];
    getTokenAfter(node: NodeOrToken): AST.Token | null;
    getTokenAfter(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getTokenAfter']>[1]): AST.Token | AST.Comment | null;
    getTokensAfter(node: NodeOrToken, options?: Parameters<ESLintSourceCode['getTokensAfter']>[1]): (AST.Token | AST.Comment)[];
    getFirstTokenBetween(left: NodeOrToken, right: NodeOrToken, options?: Parameters<ESLintSourceCode['getFirstTokenBetween']>[2]): AST.Token | AST.Comment | null;
    getFirstTokensBetween(left: NodeOrToken, right: NodeOrToken, options?: Parameters<ESLintSourceCode['getFirstTokensBetween']>[2]): (AST.Token | AST.Comment)[];
    getLastTokenBetween(left: NodeOrToken, right: NodeOrToken, options?: Parameters<ESLintSourceCode['getLastTokenBetween']>[2]): AST.Token | AST.Comment | null;
    getLastTokensBetween(left: NodeOrToken, right: NodeOrToken, options?: Parameters<ESLintSourceCode['getLastTokensBetween']>[2]): (AST.Token | AST.Comment)[];
    getTokensBetween(left: NodeOrToken, right: NodeOrToken, padding?: Parameters<ESLintSourceCode['getTokensBetween']>[2]): (AST.Token | AST.Comment)[];
    getTokens(node: NodeOrToken, beforeCount?: number, afterCount?: number): AST.Token[];
    getTokens(node: NodeOrToken, options: Parameters<ESLintSourceCode['getTokens']>[1]): (AST.Token | AST.Comment)[];
    commentsExistBetween(left: NodeOrToken, right: NodeOrToken): boolean;
    getCommentsBefore(nodeOrToken: NodeOrToken | AST.Token): AST.Comment[];
    getCommentsAfter(nodeOrToken: NodeOrToken | AST.Token): AST.Comment[];
    getCommentsInside(node: NodeOrToken): AST.Comment[];
}
