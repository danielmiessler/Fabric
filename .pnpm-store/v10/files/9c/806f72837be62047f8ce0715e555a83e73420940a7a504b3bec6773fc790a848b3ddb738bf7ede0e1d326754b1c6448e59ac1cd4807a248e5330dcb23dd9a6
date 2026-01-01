"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-immutable-reactive-statements', {
    meta: {
        docs: {
            description: "disallow reactive statements that don't reference reactive values.",
            category: 'Best Practices',
            // TODO Switch to recommended in the major version.
            recommended: false
        },
        schema: [],
        messages: {
            immutable: 'This statement is not reactive because all variables referenced in the reactive statement are immutable.'
        },
        type: 'suggestion'
    },
    create(context) {
        const scopeManager = (0, compat_1.getSourceCode)(context).scopeManager;
        const globalScope = scopeManager.globalScope;
        const toplevelScope = globalScope?.childScopes.find((scope) => scope.type === 'module') || globalScope;
        if (!globalScope || !toplevelScope) {
            return {};
        }
        const cacheMutableVariable = new WeakMap();
        /**
         * Checks whether the given reference is a mutable variable or not.
         */
        function isMutableVariableReference(reference) {
            if (reference.identifier.name.startsWith('$')) {
                // It is reactive store reference.
                return true;
            }
            if (!reference.resolved) {
                // Unknown variable
                return true;
            }
            return isMutableVariable(reference.resolved);
        }
        /**
         * Checks whether the given variable is a mutable variable or not.
         */
        function isMutableVariable(variable) {
            const cache = cacheMutableVariable.get(variable);
            if (cache != null) {
                return cache;
            }
            if (variable.defs.length === 0) {
                // Global variables are assumed to be immutable.
                return true;
            }
            const isMutableDefine = variable.defs.some((def) => {
                if (def.type === 'ImportBinding') {
                    return false;
                }
                if (def.node.type === 'AssignmentExpression') {
                    // Reactive values
                    return true;
                }
                if (def.type === 'Variable') {
                    const parent = def.parent;
                    if (parent.kind === 'const') {
                        if (def.node.init &&
                            (def.node.init.type === 'FunctionExpression' ||
                                def.node.init.type === 'ArrowFunctionExpression' ||
                                def.node.init.type === 'Literal')) {
                            return false;
                        }
                    }
                    else {
                        const pp = parent.parent;
                        if (pp && pp.type === 'ExportNamedDeclaration' && pp.declaration === parent) {
                            // Props
                            return true;
                        }
                    }
                    return hasWrite(variable);
                }
                return false;
            });
            cacheMutableVariable.set(variable, isMutableDefine);
            return isMutableDefine;
        }
        /** Checks whether the given variable has a write or reactive store reference or not. */
        function hasWrite(variable) {
            const defIds = variable.defs.map((def) => def.name);
            for (const reference of variable.references) {
                if (reference.isWrite() &&
                    !defIds.some((defId) => defId.range[0] <= reference.identifier.range[0] &&
                        reference.identifier.range[1] <= defId.range[1])) {
                    return true;
                }
                if (hasWriteMember(reference.identifier)) {
                    return true;
                }
            }
            return false;
        }
        /** Checks whether the given expression has writing to a member or not. */
        function hasWriteMember(expr) {
            if (expr.type === 'JSXIdentifier')
                return false;
            const parent = expr.parent;
            if (parent.type === 'AssignmentExpression') {
                return parent.left === expr;
            }
            if (parent.type === 'UpdateExpression') {
                return parent.argument === expr;
            }
            if (parent.type === 'UnaryExpression') {
                return parent.operator === 'delete' && parent.argument === expr;
            }
            if (parent.type === 'MemberExpression') {
                return parent.object === expr && hasWriteMember(parent);
            }
            if (parent.type === 'SvelteDirective') {
                return parent.kind === 'Binding' && parent.expression === expr;
            }
            if (parent.type === 'SvelteEachBlock') {
                return parent.expression === expr && hasWriteReference(parent.context);
            }
            return false;
        }
        /** Checks whether the given pattern has writing or not. */
        function hasWriteReference(pattern) {
            for (const id of (0, ast_utils_1.iterateIdentifiers)(pattern)) {
                const variable = (0, ast_utils_1.findVariable)(context, id);
                if (variable && hasWrite(variable))
                    return true;
            }
            return false;
        }
        /**
         * Iterates through references to top-level variables in the given range.
         */
        function* iterateRangeReferences(scope, range) {
            for (const variable of scope.variables) {
                for (const reference of variable.references) {
                    if (range[0] <= reference.identifier.range[0] &&
                        reference.identifier.range[1] <= range[1]) {
                        yield reference;
                    }
                }
            }
        }
        return {
            SvelteReactiveStatement(node) {
                for (const reference of iterateRangeReferences(toplevelScope, node.range)) {
                    if (reference.isWriteOnly()) {
                        continue;
                    }
                    if (isMutableVariableReference(reference)) {
                        return;
                    }
                }
                for (const through of toplevelScope.through.filter((reference) => node.range[0] <= reference.identifier.range[0] &&
                    reference.identifier.range[1] <= node.range[1])) {
                    if (through.identifier.name.startsWith('$$')) {
                        // Builtin `$$` vars
                        return;
                    }
                    if (through.resolved == null) {
                        // Do not report if there are missing references.
                        return;
                    }
                }
                context.report({
                    node: node.body.type === 'ExpressionStatement' &&
                        node.body.expression.type === 'AssignmentExpression' &&
                        node.body.expression.operator === '='
                        ? node.body.expression.right
                        : node.body,
                    messageId: 'immutable'
                });
            }
        };
    }
});
