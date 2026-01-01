import type { RuleListener, RuleContext, RuleModule } from '../types';
import type { TSESTree } from '@typescript-eslint/types';
import type { AST as SvAST } from 'svelte-eslint-parser';
/**
 * Define the wrapped core rule.
 */
export declare function defineWrapperListener(coreRule: RuleModule, context: RuleContext, proxyOptions: {
    createListenerProxy?: (listener: RuleListener) => RuleListener;
}): RuleListener;
/**
 * Get the proxy node
 */
export declare function getProxyNode(node: {
    type: string;
}, properties: any): any;
/**
 * Build the proxy rule listener
 */
export declare function buildProxyListener(base: RuleListener, convertNode: (node: SvAST.SvelteNode | (TSESTree.Node & {
    parent: SvAST.SvelteNode | TSESTree.Node | null;
})) => any): RuleListener;
/**
 * Get the core rule implementation from the rule name
 */
export declare function getCoreRule(ruleName: string): RuleModule;
