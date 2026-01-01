"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.analyzeScope = analyzeScope;
exports.analyzeReactiveScope = analyzeReactiveScope;
exports.analyzeStoreScope = analyzeStoreScope;
exports.analyzePropsScope = analyzePropsScope;
exports.analyzeSnippetsScope = analyzeSnippetsScope;
const eslint_scope_1 = require("eslint-scope");
const traverse_1 = require("../traverse");
const scope_1 = require("../scope");
const utils_1 = require("../utils");
/**
 * Analyze scope
 */
function analyzeScope(node, parserOptions) {
    const ecmaVersion = parserOptions.ecmaVersion || 2020;
    const ecmaFeatures = parserOptions.ecmaFeatures || {};
    const sourceType = parserOptions.sourceType || "module";
    const root = node.type === "Program"
        ? node
        : {
            type: "Program",
            body: [node],
            sourceType,
        };
    return (0, eslint_scope_1.analyze)(root, {
        ignoreEval: true,
        nodejsScope: false,
        impliedStrict: ecmaFeatures.impliedStrict,
        ecmaVersion: typeof ecmaVersion === "number" ? ecmaVersion : 2022,
        sourceType,
        fallback: traverse_1.getFallbackKeys,
    });
}
/** Analyze reactive scope */
function analyzeReactiveScope(scopeManager) {
    for (const reference of [...scopeManager.globalScope.through]) {
        const parent = reference.writeExpr && getParent(reference.writeExpr);
        if ((parent === null || parent === void 0 ? void 0 : parent.type) === "AssignmentExpression") {
            const pp = getParent(parent);
            if ((pp === null || pp === void 0 ? void 0 : pp.type) === "ExpressionStatement") {
                const ppp = getParent(pp);
                if ((ppp === null || ppp === void 0 ? void 0 : ppp.type) === "SvelteReactiveStatement" && ppp.label.name === "$") {
                    const referenceScope = reference.from;
                    if (referenceScope.type === "module") {
                        // It is computed
                        transformComputedVariable(parent, ppp, reference);
                        continue;
                    }
                }
            }
        }
    }
    /** Transform ref to ComputedVariable */
    function transformComputedVariable(node, parent, reference) {
        const referenceScope = reference.from;
        const name = reference.identifier.name;
        let variable = referenceScope.set.get(name);
        if (!variable) {
            variable = new eslint_scope_1.Variable();
            variable.scope = referenceScope;
            variable.name = name;
            (0, utils_1.addElementToSortedArray)(variable.defs, {
                type: "ComputedVariable",
                node: node,
                parent: parent,
                name: reference.identifier,
            }, (a, b) => a.node.range[0] - b.node.range[0]);
            (0, scope_1.addVariable)(referenceScope.variables, variable);
            referenceScope.set.set(name, variable);
        }
        (0, utils_1.addElementToSortedArray)(variable.identifiers, reference.identifier, (a, b) => a.range[0] - b.range[0]);
        reference.resolved = variable;
        removeReferenceFromThrough(reference, referenceScope);
    }
}
/**
 * Analyze store scope. e.g. $count
 */
function analyzeStoreScope(scopeManager) {
    const moduleScope = scopeManager.scopes.find((scope) => scope.type === "module");
    if (!moduleScope) {
        return;
    }
    const toBeMarkAsUsedReferences = [];
    for (const reference of [...scopeManager.globalScope.through]) {
        if (reference.identifier.name.startsWith("$")) {
            const realName = reference.identifier.name.slice(1);
            const variable = moduleScope.set.get(realName);
            if (variable) {
                if (reference.isWriteOnly()) {
                    // Need mark as used
                    toBeMarkAsUsedReferences.push(reference);
                }
                // It does not write directly to the original variable.
                // Therefore, this variable is always a reference.
                reference.isWrite = () => false;
                reference.isWriteOnly = () => false;
                reference.isReadWrite = () => false;
                reference.isReadOnly = () => true;
                reference.isRead = () => true;
                (0, scope_1.addReference)(variable.references, reference);
                reference.resolved = variable;
                removeReferenceFromThrough(reference, moduleScope);
            }
        }
    }
    for (const variable of new Set(toBeMarkAsUsedReferences.map((ref) => ref.resolved))) {
        if (variable.references.some((ref) => !toBeMarkAsUsedReferences.includes(ref) &&
            ref.identifier !== variable.identifiers[0])) {
            // It is already used.
            continue;
        }
        // Add the virtual reference for reading.
        addVirtualReference(variable.identifiers[0], variable, moduleScope, {
            read: true,
        }).svelteMarkAsUsed = true;
    }
}
/** Transform props exports */
function analyzePropsScope(body, scopeManager, svelteParseContext) {
    var _a;
    const moduleScope = scopeManager.scopes.find((scope) => scope.type === "module");
    if (!moduleScope) {
        return;
    }
    for (const node of body.body) {
        if (node.type === "ExportNamedDeclaration") {
            // Process for Svelte v4 style props. e.g. `export let x`;
            if (node.declaration) {
                if (node.declaration.type === "VariableDeclaration") {
                    for (const decl of node.declaration.declarations) {
                        for (const pattern of extractPattern(decl.id)) {
                            if (pattern.type === "Identifier") {
                                addPropReference(pattern, moduleScope);
                            }
                        }
                    }
                }
            }
            else {
                for (const spec of node.specifiers) {
                    if (spec.local.type !== "Literal") {
                        addPropReference(spec.local, moduleScope);
                    }
                }
            }
        }
        else if (node.type === "VariableDeclaration") {
            if (svelteParseContext.runes) {
                // Process for Svelte v5 Runes props. e.g. `let { x = $bindable() } = $props()`;
                for (const decl of node.declarations) {
                    if (((_a = decl.init) === null || _a === void 0 ? void 0 : _a.type) === "CallExpression" &&
                        decl.init.callee.type === "Identifier" &&
                        decl.init.callee.name === "$props" &&
                        decl.id.type === "ObjectPattern") {
                        for (const pattern of extractPattern(decl.id)) {
                            if (pattern.type === "AssignmentPattern" &&
                                pattern.left.type === "Identifier" &&
                                pattern.right.type === "CallExpression" &&
                                pattern.right.callee.type === "Identifier" &&
                                pattern.right.callee.name === "$bindable") {
                                addPropReference(pattern.left, moduleScope);
                            }
                        }
                    }
                }
            }
        }
    }
    function* extractPattern(node) {
        yield node;
        if (node.type === "Identifier") {
            return;
        }
        if (node.type === "ObjectPattern") {
            for (const prop of node.properties) {
                if (prop.type === "Property") {
                    yield* extractPattern(prop.value);
                }
                else {
                    yield* extractPattern(prop);
                }
            }
            return;
        }
        if (node.type === "ArrayPattern") {
            for (const elem of node.elements) {
                if (elem) {
                    yield* extractPattern(elem);
                }
            }
            return;
        }
        if (node.type === "AssignmentPattern") {
            yield* extractPattern(node.left);
            return;
        }
        if (node.type === "RestElement") {
            yield* extractPattern(node.argument);
        }
    }
    /** Add virtual prop reference */
    function addPropReference(node, scope) {
        for (const variable of scope.variables) {
            if (variable.name !== node.name) {
                continue;
            }
            if (variable.references.some((ref) => ref.sveltePropReference)) {
                continue;
            }
            // Add the virtual reference for writing.
            const reference = addVirtualReference(Object.assign(Object.assign({}, node), { 
                // @ts-expect-error -- ignore
                parent: body, loc: {
                    start: Object.assign({}, node.loc.start),
                    end: Object.assign({}, node.loc.end),
                }, range: [...node.range] }), variable, scope, {
                write: true,
                read: true,
            });
            reference.sveltePropReference = true;
        }
    }
}
/** Analyze snippets in component scope */
function analyzeSnippetsScope(snippets, scopeManager) {
    for (const snippet of snippets) {
        const parent = snippet.parent;
        if (parent.type === "SvelteElement" &&
            (parent.kind === "component" ||
                (parent.kind === "special" && parent.name.name === "svelte:component"))) {
            const scope = (0, scope_1.getScopeFromNode)(scopeManager, snippet.id);
            const upperScope = scope.upper;
            if (!upperScope)
                continue;
            const variable = upperScope.set.get(snippet.id.name);
            if (!variable)
                continue;
            // Add the virtual reference for reading.
            const reference = addVirtualReference(snippet.id, variable, upperScope, {
                read: true,
            });
            reference.svelteSnippetReference = true;
        }
    }
}
/** Remove reference from through */
function removeReferenceFromThrough(reference, baseScope) {
    const variable = reference.resolved;
    const name = reference.identifier.name;
    let scope = baseScope;
    while (scope) {
        scope.through = scope.through.filter((ref) => {
            if (reference === ref) {
                return false;
            }
            else if (ref.identifier.name === name) {
                ref.resolved = variable;
                if (!variable.references.includes(ref)) {
                    (0, scope_1.addReference)(variable.references, ref);
                }
                return false;
            }
            return true;
        });
        scope = scope.upper;
    }
}
/**
 * Add the virtual reference.
 */
function addVirtualReference(node, variable, scope, readWrite) {
    const reference = new eslint_scope_1.Reference();
    reference.svelteVirtualReference = true;
    reference.from = scope;
    reference.identifier = node;
    reference.isWrite = () => Boolean(readWrite.write);
    reference.isWriteOnly = () => Boolean(readWrite.write) && !readWrite.read;
    reference.isRead = () => Boolean(readWrite.read);
    reference.isReadOnly = () => Boolean(readWrite.read) && !readWrite.write;
    reference.isReadWrite = () => Boolean(readWrite.read && readWrite.write);
    (0, scope_1.addReference)(variable.references, reference);
    reference.resolved = variable;
    return reference;
}
/** Get parent node */
function getParent(node) {
    return node.parent;
}
