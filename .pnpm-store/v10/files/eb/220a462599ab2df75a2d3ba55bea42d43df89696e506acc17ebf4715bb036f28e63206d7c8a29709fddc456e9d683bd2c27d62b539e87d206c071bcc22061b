"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.removeAllScopeAndVariableAndReference = removeAllScopeAndVariableAndReference;
exports.getScopeFromNode = getScopeFromNode;
exports.findVariable = findVariable;
exports.getProgramScope = getProgramScope;
exports.removeIdentifierVariable = removeIdentifierVariable;
exports.getAllReferences = getAllReferences;
exports.removeIdentifierReference = removeIdentifierReference;
exports.removeReference = removeReference;
exports.removeScope = removeScope;
exports.replaceScope = replaceScope;
exports.addVariable = addVariable;
exports.addReference = addReference;
exports.addAllReferences = addAllReferences;
exports.simplifyScope = simplifyScope;
const traverse_1 = require("../traverse");
const utils_1 = require("../utils");
/** Remove all scope, variable, and reference */
function removeAllScopeAndVariableAndReference(target, info) {
    const targetScopes = new Set();
    (0, traverse_1.traverseNodes)(target, {
        visitorKeys: info.visitorKeys,
        enterNode(node) {
            const scope = info.scopeManager.acquire(node);
            if (scope) {
                targetScopes.add(scope);
                return;
            }
            if (node.type === "Identifier") {
                let scope = getScopeFromNode(info.scopeManager, node);
                while (scope &&
                    scope.block.type !== "Program" &&
                    target.range[0] <= scope.block.range[0] &&
                    scope.block.range[1] <= target.range[1]) {
                    scope = scope.upper;
                }
                if (targetScopes.has(scope)) {
                    return;
                }
                removeIdentifierVariable(node, scope);
                removeIdentifierReference(node, scope);
            }
        },
        leaveNode() {
            // noop
        },
    });
    for (const scope of targetScopes) {
        removeScope(info.scopeManager, scope);
    }
}
/**
 * Gets the scope for the current node
 */
function getScopeFromNode(scopeManager, currentNode) {
    let node = currentNode;
    for (; node; node = node.parent || null) {
        const scope = scopeManager.acquire(node, false);
        if (scope) {
            if (scope.type === "function-expression-name") {
                return scope.childScopes[0];
            }
            if (scope.type === "global" &&
                node.type === "Program" &&
                node.sourceType === "module") {
                return scope.childScopes.find((s) => s.type === "module") || scope;
            }
            return scope;
        }
    }
    const global = scopeManager.globalScope;
    return global;
}
/**
 * Find the variable of a given identifier.
 */
function findVariable(scopeManager, node) {
    let scope = getScopeFromNode(scopeManager, node);
    while (scope != null) {
        const variable = scope.set.get(node.name);
        if (variable != null) {
            return variable;
        }
        scope = scope.upper;
    }
    return null;
}
/**
 * Gets the scope for the Program node
 */
function getProgramScope(scopeManager) {
    const globalScope = scopeManager.globalScope;
    return (globalScope.childScopes.find((s) => s.type === "module") || globalScope);
}
/** Remove variable */
function removeIdentifierVariable(node, scope) {
    if (node.type === "ObjectPattern") {
        for (const prop of node.properties) {
            if (prop.type === "Property") {
                removeIdentifierVariable(prop.value, scope);
            }
            else if (prop.type === "RestElement") {
                removeIdentifierVariable(prop, scope);
            }
        }
        return;
    }
    if (node.type === "ArrayPattern") {
        for (const element of node.elements) {
            if (!element)
                continue;
            removeIdentifierVariable(element, scope);
        }
        return;
    }
    if (node.type === "AssignmentPattern") {
        removeIdentifierVariable(node.left, scope);
        return;
    }
    if (node.type === "RestElement") {
        removeIdentifierVariable(node.argument, scope);
        return;
    }
    if (node.type === "MemberExpression") {
        return;
    }
    if (node.type !== "Identifier") {
        return;
    }
    for (let varIndex = 0; varIndex < scope.variables.length; varIndex++) {
        const variable = scope.variables[varIndex];
        const defIndex = variable.defs.findIndex((def) => def.name === node);
        if (defIndex < 0) {
            continue;
        }
        variable.defs.splice(defIndex, 1);
        if (variable.defs.length === 0) {
            // Remove variable
            referencesToThrough(variable.references, scope);
            variable.references.forEach((r) => {
                if (r.init)
                    r.init = false;
                r.resolved = null;
            });
            scope.variables.splice(varIndex, 1);
            const name = node.name;
            if (variable === scope.set.get(name)) {
                scope.set.delete(name);
            }
        }
        else {
            const idIndex = variable.identifiers.indexOf(node);
            if (idIndex >= 0) {
                variable.identifiers.splice(idIndex, 1);
            }
        }
        return;
    }
}
/** Get all references */
function* getAllReferences(node, scope) {
    if (node.type === "ObjectPattern") {
        for (const prop of node.properties) {
            if (prop.type === "Property") {
                yield* getAllReferences(prop.value, scope);
            }
            else if (prop.type === "RestElement") {
                yield* getAllReferences(prop, scope);
            }
        }
        return;
    }
    if (node.type === "ArrayPattern") {
        for (const element of node.elements) {
            if (!element)
                continue;
            yield* getAllReferences(element, scope);
        }
        return;
    }
    if (node.type === "AssignmentPattern") {
        yield* getAllReferences(node.left, scope);
        return;
    }
    if (node.type === "RestElement") {
        yield* getAllReferences(node.argument, scope);
        return;
    }
    if (node.type === "MemberExpression") {
        return;
    }
    if (node.type !== "Identifier") {
        return;
    }
    const ref = scope.references.find((ref) => ref.identifier === node);
    if (ref)
        yield ref;
}
/** Remove reference */
function removeIdentifierReference(node, scope) {
    const reference = scope.references.find((ref) => ref.identifier === node);
    if (reference) {
        removeReference(reference, scope);
        return true;
    }
    const location = node.range[0];
    const pendingScopes = [];
    for (const childScope of scope.childScopes) {
        const range = childScope.block.range;
        if (range[0] <= location && location < range[1]) {
            if (removeIdentifierReference(node, childScope)) {
                return true;
            }
        }
        else {
            pendingScopes.push(childScope);
        }
    }
    for (const childScope of pendingScopes) {
        if (removeIdentifierReference(node, childScope)) {
            return true;
        }
    }
    return false;
}
/** Remove reference */
function removeReference(reference, baseScope) {
    if (reference.resolved) {
        if (reference.resolved.defs.some((d) => d.name === reference.identifier)) {
            // remove var
            const varIndex = baseScope.variables.indexOf(reference.resolved);
            if (varIndex >= 0) {
                baseScope.variables.splice(varIndex, 1);
            }
            const name = reference.identifier.name;
            if (reference.resolved === baseScope.set.get(name)) {
                baseScope.set.delete(name);
            }
        }
        else {
            const refIndex = reference.resolved.references.indexOf(reference);
            if (refIndex >= 0) {
                reference.resolved.references.splice(refIndex, 1);
            }
        }
    }
    let scope = baseScope;
    while (scope) {
        const refIndex = scope.references.indexOf(reference);
        if (refIndex >= 0) {
            scope.references.splice(refIndex, 1);
        }
        const throughIndex = scope.through.indexOf(reference);
        if (throughIndex >= 0) {
            scope.through.splice(throughIndex, 1);
        }
        scope = scope.upper;
    }
}
/** Move reference to through */
function referencesToThrough(references, baseScope) {
    let scope = baseScope;
    while (scope) {
        addAllReferences(scope.through, references);
        scope = scope.upper;
    }
}
/** Remove scope */
function removeScope(scopeManager, scope) {
    while (scope.childScopes[0]) {
        removeScope(scopeManager, scope.childScopes[0]);
    }
    while (scope.references[0]) {
        removeReference(scope.references[0], scope);
    }
    const upper = scope.upper;
    if (upper) {
        const index = upper.childScopes.indexOf(scope);
        if (index >= 0) {
            upper.childScopes.splice(index, 1);
        }
    }
    const index = scopeManager.scopes.indexOf(scope);
    if (index >= 0) {
        scopeManager.scopes.splice(index, 1);
    }
}
/** Replace scope */
function replaceScope(scopeManager, scope, newChildScopes = []) {
    // remove scope from scopeManager
    scopeManager.scopes = scopeManager.scopes.filter((s) => s !== scope);
    const upper = scope.upper;
    if (upper) {
        // remove scope from upper and marge childScopes
        upper.childScopes.splice(upper.childScopes.indexOf(scope), 1, ...newChildScopes);
        for (const child of newChildScopes) {
            child.upper = upper;
            replaceVariableScope(child, scope);
        }
    }
    /** Replace variableScope  */
    function replaceVariableScope(child, replaceTarget) {
        if (child.variableScope === replaceTarget) {
            child.variableScope = child.upper.variableScope;
            for (const c of child.childScopes) {
                replaceVariableScope(c, replaceTarget);
            }
        }
    }
}
/**
 * Add variable to array
 */
function addVariable(list, variable) {
    (0, utils_1.addElementToSortedArray)(list, variable, (a, b) => {
        const idA = getFirstId(a);
        const idB = getFirstId(b);
        return idA.range[0] - idB.range[0];
    });
    /** Get first id from give variable */
    function getFirstId(v) {
        var _a, _b;
        return v.identifiers[0] || ((_a = v.defs[0]) === null || _a === void 0 ? void 0 : _a.name) || ((_b = v.references[0]) === null || _b === void 0 ? void 0 : _b.identifier);
    }
}
/**
 * Add reference to array
 */
function addReference(list, reference) {
    (0, utils_1.addElementToSortedArray)(list, reference, (a, b) => a.identifier.range[0] - b.identifier.range[0]);
}
/**
 * Add all references to array
 */
function addAllReferences(list, elements) {
    (0, utils_1.addElementsToSortedArray)(list, elements, (a, b) => a.identifier.range[0] - b.identifier.range[0]);
}
/**
 * Simplify scope data.
 * @deprecated For Debug
 */
function simplifyScope(scope) {
    return {
        type: scope.type,
        childScopes: scope.childScopes.map(simplifyScope),
        block: {
            type: scope.block.type,
            loc: JSON.stringify(scope.block.loc),
        },
        variables: scope.type === "global" ? null : simplifyVariables(scope.variables),
        references: scope.references.map(simplifyReference),
        through: scope.through.map(simplifyReference),
        get original() {
            return scope;
        },
    };
}
/**
 * Simplify variables data.
 * @deprecated For Debug
 */
function simplifyVariables(variables) {
    return Object.fromEntries(variables.map((v) => {
        var _a;
        return [
            v.name,
            {
                loc: JSON.stringify((_a = v.defs[0]) === null || _a === void 0 ? void 0 : _a.node.loc),
            },
        ];
    }));
}
/**
 * Simplify reference data.
 * @deprecated For Debug
 */
function simplifyReference(reference) {
    return {
        name: reference.identifier.name,
        loc: JSON.stringify(reference.identifier.loc),
    };
}
