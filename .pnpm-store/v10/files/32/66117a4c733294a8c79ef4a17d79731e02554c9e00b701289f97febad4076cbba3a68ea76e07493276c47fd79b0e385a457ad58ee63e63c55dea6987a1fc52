import type { ScopeManager, Scope, Reference, Variable } from "eslint-scope";
import type * as ESTree from "estree";
import type { TSESTree } from "@typescript-eslint/types";
/** Remove all scope, variable, and reference */
export declare function removeAllScopeAndVariableAndReference(target: ESTree.Node | TSESTree.Node, info: {
    visitorKeys?: {
        [type: string]: string[];
    } | {
        readonly [type: string]: readonly string[] | undefined;
    };
    scopeManager: ScopeManager;
}): void;
/**
 * Gets the scope for the current node
 */
export declare function getScopeFromNode(scopeManager: ScopeManager, currentNode: ESTree.Node): Scope;
/**
 * Find the variable of a given identifier.
 */
export declare function findVariable(scopeManager: ScopeManager, node: ESTree.Identifier): Variable | null;
/**
 * Gets the scope for the Program node
 */
export declare function getProgramScope(scopeManager: ScopeManager): Scope;
/** Remove variable */
export declare function removeIdentifierVariable(node: ESTree.Pattern | TSESTree.BindingName | TSESTree.RestElement | TSESTree.DestructuringPattern, scope: Scope): void;
/** Get all references */
export declare function getAllReferences(node: ESTree.Pattern | TSESTree.BindingName | TSESTree.RestElement | TSESTree.DestructuringPattern, scope: Scope): Iterable<Reference>;
/** Remove reference */
export declare function removeIdentifierReference(node: ESTree.Identifier, scope: Scope): boolean;
/** Remove reference */
export declare function removeReference(reference: Reference, baseScope: Scope): void;
/** Remove scope */
export declare function removeScope(scopeManager: ScopeManager, scope: Scope): void;
/** Replace scope */
export declare function replaceScope(scopeManager: ScopeManager, scope: Scope, newChildScopes?: Scope[]): void;
/**
 * Add variable to array
 */
export declare function addVariable(list: Variable[], variable: Variable): void;
/**
 * Add reference to array
 */
export declare function addReference(list: Reference[], reference: Reference): void;
/**
 * Add all references to array
 */
export declare function addAllReferences(list: Reference[], elements: Reference[]): void;
/**
 * Simplify scope data.
 * @deprecated For Debug
 */
export declare function simplifyScope(scope: Scope): unknown;
