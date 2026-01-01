"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const svelte_eslint_parser_1 = require("svelte-eslint-parser");
const compat_1 = require("../utils/compat");
/**
 * Get usage of `tick`
 */
function extractTickReferences(context) {
    const referenceTracker = new eslint_utils_1.ReferenceTracker((0, compat_1.getSourceCode)(context).scopeManager.globalScope);
    const a = referenceTracker.iterateEsmReferences({
        svelte: {
            [eslint_utils_1.ReferenceTracker.ESM]: true,
            tick: {
                [eslint_utils_1.ReferenceTracker.CALL]: true
            }
        }
    });
    return Array.from(a).map(({ node, path }) => {
        return {
            node: node,
            name: path[path.length - 1]
        };
    });
}
/**
 * Get usage of `setTimeout`, `setInterval`, `queueMicrotask`
 */
function extractTaskReferences(context) {
    const referenceTracker = new eslint_utils_1.ReferenceTracker((0, compat_1.getSourceCode)(context).scopeManager.globalScope);
    const a = referenceTracker.iterateGlobalReferences({
        setTimeout: { [eslint_utils_1.ReferenceTracker.CALL]: true },
        setInterval: { [eslint_utils_1.ReferenceTracker.CALL]: true },
        queueMicrotask: { [eslint_utils_1.ReferenceTracker.CALL]: true }
    });
    return Array.from(a).map(({ node, path }) => {
        return {
            node: node,
            name: path[path.length - 1]
        };
    });
}
/**
 * If `node` is inside of `maybeAncestorNode`, return true.
 */
function isChildNode(maybeAncestorNode, node) {
    let parent = node.parent;
    while (parent) {
        if (parent === maybeAncestorNode)
            return true;
        parent = parent.parent;
    }
    return false;
}
/**
 * Return true if `node` is a function call.
 */
function isFunctionCall(node) {
    if (node.type !== 'Identifier')
        return false;
    const { parent } = node;
    if (parent?.type !== 'CallExpression')
        return false;
    return parent.callee.type === 'Identifier' && parent.callee.name === node.name;
}
/**
 * Return true if `node` is a reactive variable.
 */
function isReactiveVariableNode(reactiveVariableReferences, node) {
    if (node.type !== 'Identifier')
        return false;
    return reactiveVariableReferences.includes(node);
}
/**
 * e.g. foo.bar = baz + 1
 * If node is `foo`, return true.
 * Otherwise, return false.
 */
function isNodeForAssign(node) {
    const { parent } = node;
    if (parent?.type === 'AssignmentExpression') {
        return parent.left.type === 'Identifier' && parent.left.name === node.name;
    }
    return (parent?.type === 'MemberExpression' &&
        parent.parent?.type === 'AssignmentExpression' &&
        parent.parent.left.type === 'MemberExpression' &&
        parent.parent.left.object.type === 'Identifier' &&
        parent.parent.left.object.name === node.name);
}
/**
 * Return true if `node` is inside of `then` or `catch`.
 */
function isPromiseThenOrCatchBody(node) {
    if (!getDeclarationBody(node))
        return false;
    const { parent } = node;
    if (parent?.type !== 'CallExpression' || parent?.callee?.type !== 'MemberExpression') {
        return false;
    }
    const { property } = parent.callee;
    if (property?.type !== 'Identifier')
        return false;
    return ['then', 'catch'].includes(property.name);
}
/**
 * Get all reactive variable reference.
 */
function getReactiveVariableReferences(context) {
    const scopeManager = (0, compat_1.getSourceCode)(context).scopeManager;
    // Find the top-level (module or global) scope.
    // Any variable defined at the top-level (module scope or global scope) can be made reactive.
    const toplevelScope = scopeManager.globalScope?.childScopes.find((scope) => scope.type === 'module') ||
        scopeManager.globalScope;
    if (!toplevelScope) {
        return [];
    }
    // Extracts all reactive references to variables defined in the top-level scope.
    const reactiveVariableNodes = [];
    for (const variable of toplevelScope.variables) {
        for (const reference of variable.references) {
            if (reference.identifier.type === 'Identifier' && !isFunctionCall(reference.identifier)) {
                reactiveVariableNodes.push(reference.identifier);
            }
        }
    }
    return reactiveVariableNodes;
}
/**
 * Get all tracked reactive variables.
 */
function getTrackedVariableNodes(reactiveVariableReferences, ast) {
    const reactiveVariableNodes = new Set();
    for (const identifier of reactiveVariableReferences) {
        if (
        // If the identifier is within the reactive statement range,
        // it is used within the reactive statement.
        ast.range[0] <= identifier.range[0] &&
            identifier.range[1] <= ast.range[1]) {
            reactiveVariableNodes.add(identifier);
        }
    }
    return reactiveVariableNodes;
}
/**  */
function getDeclarationBody(node, functionName) {
    if (node.type === 'VariableDeclarator' &&
        node.id.type === 'Identifier' &&
        (!functionName || node.id.name === functionName)) {
        if (node.init?.type === 'ArrowFunctionExpression' || node.init?.type === 'FunctionExpression') {
            return node.init.body;
        }
    }
    else if (node.type === 'FunctionDeclaration' &&
        node.id?.type === 'Identifier' &&
        (!functionName || node.id?.name === functionName)) {
        return node.body;
    }
    else if (!functionName && node.type === 'ArrowFunctionExpression') {
        return node.body;
    }
    return null;
}
/**  */
function getFunctionDeclarationNode(context, functionCall) {
    const variable = (0, ast_utils_1.findVariable)(context, functionCall);
    if (!variable) {
        return null;
    }
    for (const def of variable.defs) {
        if (def.type === 'FunctionName') {
            if (def.node.type === 'FunctionDeclaration') {
                return def.node.body;
            }
        }
        if (def.type === 'Variable') {
            if (def.node.init &&
                (def.node.init.type === 'FunctionExpression' ||
                    def.node.init.type === 'ArrowFunctionExpression')) {
                return def.node.init.body;
            }
        }
    }
    return null;
}
/**
 * If the node is inside of a function, return true.
 *
 * e.g. `$: await foo()`
 * if `node` is `foo`, return false because reactive statement is not function.
 *
 * e.g. `const bar = () => foo()`
 * if `node` is `foo`, return true.
 *
 */
function isInsideOfFunction(node) {
    let parent = node;
    while (parent) {
        parent = parent.parent;
        if (!parent)
            break;
        if (parent.type === 'FunctionDeclaration' && parent.async)
            return true;
        if (parent.type === 'VariableDeclarator' &&
            (parent.init?.type === 'FunctionExpression' ||
                parent.init?.type === 'ArrowFunctionExpression') &&
            parent.init?.async) {
            return true;
        }
    }
    return false;
}
/** Let's lint! */
function doLint(context, ast, callFuncIdentifiers, tickCallExpressions, taskReferences, reactiveVariableNames, reactiveVariableReferences, pIsSameTask) {
    const processed = new Set();
    verifyInternal(ast, callFuncIdentifiers, pIsSameTask);
    /** verify for node */
    function verifyInternal(ast, callFuncIdentifiers, pIsSameTask) {
        if (processed.has(ast)) {
            // Avoid infinite recursion with recursive references.
            return;
        }
        processed.add(ast);
        let isSameMicroTask = pIsSameTask;
        const differentMicroTaskEnterNodes = [];
        (0, svelte_eslint_parser_1.traverseNodes)(ast, {
            enterNode(node) {
                // Promise.then() or Promise.catch() is called.
                if (isPromiseThenOrCatchBody(node)) {
                    differentMicroTaskEnterNodes.push(node);
                    isSameMicroTask = false;
                }
                // `tick`, `setTimeout`, `setInterval` , `queueMicrotask` is called
                for (const { node: callExpression } of [...tickCallExpressions, ...taskReferences]) {
                    if (isChildNode(callExpression, node)) {
                        differentMicroTaskEnterNodes.push(node);
                        isSameMicroTask = false;
                    }
                }
                // left side of await block
                if (node.parent?.type === 'AssignmentExpression' &&
                    node.parent?.right.type === 'AwaitExpression' &&
                    node.parent?.left === node) {
                    differentMicroTaskEnterNodes.push(node);
                    isSameMicroTask = false;
                }
                if (node.type === 'Identifier' && isFunctionCall(node)) {
                    // traverse used functions body
                    const functionDeclarationNode = getFunctionDeclarationNode(context, node);
                    if (functionDeclarationNode) {
                        verifyInternal(functionDeclarationNode, [...callFuncIdentifiers, node], isSameMicroTask);
                    }
                }
                if (!isSameMicroTask) {
                    if (isReactiveVariableNode(reactiveVariableReferences, node) &&
                        reactiveVariableNames.includes(node.name) &&
                        isNodeForAssign(node)) {
                        context.report({
                            node,
                            loc: node.loc,
                            messageId: 'unexpected'
                        });
                        callFuncIdentifiers.forEach((callFuncIdentifier) => {
                            context.report({
                                node: callFuncIdentifier,
                                loc: callFuncIdentifier.loc,
                                messageId: 'unexpectedCall',
                                data: {
                                    variableName: node.name
                                }
                            });
                        });
                    }
                }
            },
            leaveNode(node) {
                if (node.type === 'AwaitExpression') {
                    if (ast.parent?.type === 'SvelteReactiveStatement') {
                        // MEMO: It checks that `await` is used in reactive statement directly or not.
                        // If `await` is used in inner function of a reactive statement, result of `isInsideOfFunction` will be `true`.
                        if (!isInsideOfFunction(node)) {
                            isSameMicroTask = false;
                        }
                    }
                    else {
                        isSameMicroTask = false;
                    }
                }
                if (differentMicroTaskEnterNodes.includes(node)) {
                    isSameMicroTask = true;
                }
            }
        });
    }
}
exports.default = (0, utils_1.createRule)('infinite-reactive-loop', {
    meta: {
        docs: {
            description: "Svelte runtime prevents calling the same reactive statement twice in a microtask. But between different microtask, it doesn't prevent.",
            category: 'Possible Errors',
            // TODO Switch to recommended in the major version.
            recommended: false
        },
        schema: [],
        messages: {
            unexpected: 'Possibly it may occur an infinite reactive loop.',
            unexpectedCall: 'Possibly it may occur an infinite reactive loop because this function may update `{{variableName}}`.'
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            ['SvelteReactiveStatement']: (ast) => {
                const tickCallExpressions = extractTickReferences(context);
                const taskReferences = extractTaskReferences(context);
                const reactiveVariableReferences = getReactiveVariableReferences(context);
                const trackedVariableNodes = getTrackedVariableNodes(reactiveVariableReferences, ast);
                doLint(context, ast.body, [], tickCallExpressions, taskReferences, Array.from(trackedVariableNodes).map((node) => node.name), reactiveVariableReferences, true);
            }
        };
    }
});
