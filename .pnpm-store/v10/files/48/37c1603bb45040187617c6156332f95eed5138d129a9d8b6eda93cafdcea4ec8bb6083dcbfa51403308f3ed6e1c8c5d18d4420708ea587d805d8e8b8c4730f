import type { VisitorKeys } from "eslint-visitor-keys";
import type ESTree from "estree";
import type { SvelteNode } from "./ast";
import type { TSESTree } from "@typescript-eslint/types";
import type { VisitorKeys as TSESVisitorKeys } from "@typescript-eslint/visitor-keys";
/**
 * Get the keys of the given node to traverse it.
 * @param node The node to get.
 * @returns The keys to traverse.
 */
export declare function getFallbackKeys(node: any): string[];
/**
 * Get the keys of the given node to traverse it.
 * @param node The node to get.
 * @returns The keys to traverse.
 */
export declare function getKeys(node: any, visitorKeys?: VisitorKeys | TSESVisitorKeys): string[];
/**
 * Get the nodes of the given node.
 * @param node The node to get.
 */
export declare function getNodes(node: any, key: string): IterableIterator<SvelteNode | ESTree.Node | TSESTree.Node>;
export interface Visitor<N> {
    visitorKeys?: VisitorKeys | TSESVisitorKeys;
    enterNode(node: N, parent: N | null): void;
    leaveNode(node: N, parent: N | null): void;
}
export declare function traverseNodes(node: SvelteNode, visitor: Visitor<SvelteNode | ESTree.Node>): void;
export declare function traverseNodes(node: ESTree.Node, visitor: Visitor<ESTree.Node>): void;
export declare function traverseNodes(node: TSESTree.Node, visitor: Visitor<TSESTree.Node>): void;
export declare function traverseNodes(node: ESTree.Node | TSESTree.Node, visitor: Visitor<ESTree.Node | TSESTree.Node>): void;
