"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-reactive-reassign', {
    meta: {
        docs: {
            description: 'disallow reassigning reactive values',
            category: 'Possible Errors',
            // TODO Switch to recommended in the major version.
            recommended: false
        },
        schema: [
            {
                type: 'object',
                properties: {
                    props: {
                        type: 'boolean'
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            assignmentToReactiveValue: "Assignment to reactive value '{{name}}'.",
            assignmentToReactiveValueProp: "Assignment to property of reactive value '{{name}}'."
        },
        type: 'problem'
    },
    create(context) {
        const props = context.options[0]?.props !== false; // default true
        const sourceCode = (0, compat_1.getSourceCode)(context);
        const scopeManager = sourceCode.scopeManager;
        const globalScope = scopeManager.globalScope;
        const toplevelScope = globalScope?.childScopes.find((scope) => scope.type === 'module') || globalScope;
        if (!globalScope || !toplevelScope) {
            return {};
        }
        const CHECK_REASSIGN = {
            UpdateExpression: 
            // e.g. foo ++, foo --
            ({ parent }) => ({ type: 'reassign', node: parent }),
            UnaryExpression: ({ parent }) => {
                if (parent.operator === 'delete') {
                    // e.g. delete foo.prop
                    return { type: 'reassign', node: parent };
                }
                return null;
            },
            AssignmentExpression: ({ node, parent }) => {
                if (parent.left === node) {
                    // e.g. foo = 42, foo += 42, foo -= 42
                    return { type: 'reassign', node: parent };
                }
                return null;
            },
            ForInStatement: ({ node, parent }) => {
                if (parent.left === node) {
                    // e.g. for (foo in itr)
                    return { type: 'reassign', node: parent };
                }
                return null;
            },
            ForOfStatement: ({ node, parent }) => {
                if (parent.left === node) {
                    // e.g. for (foo of itr)
                    return { type: 'reassign', node: parent };
                }
                return null;
            },
            CallExpression: ({ node, parent, pathNodes }) => {
                if (pathNodes.length > 0 && parent.callee === node) {
                    const mem = pathNodes[pathNodes.length - 1];
                    const callName = (0, eslint_utils_1.getPropertyName)(mem);
                    if (callName &&
                        /^(?:push|pop|shift|unshift|reverse|splice|sort|copyWithin|fill)$/u.test(callName)) {
                        // e.g. foo.push()
                        return {
                            type: 'reassign',
                            node: parent,
                            pathNodes: pathNodes.slice(0, -1)
                        };
                    }
                }
                return null;
            },
            MemberExpression: ({ node, parent, pathNodes }) => {
                if (parent.object === node) {
                    // The context to check next.
                    return {
                        type: 'check',
                        node: parent,
                        pathNodes: [...pathNodes, parent]
                    };
                }
                return null;
            },
            ChainExpression: ({ parent }) => {
                // e.g. `foo?.prop`
                // The context to check next.
                return { type: 'check', node: parent };
            },
            ConditionalExpression: ({ node, parent }) => {
                if (parent.test === node) {
                    return null;
                }
                // The context to check next for `(test ? foo : bar).prop`.
                return { type: 'check', node: parent };
            },
            Property: ({ node, parent }) => {
                if (parent.value === node && parent.parent && parent.parent.type === 'ObjectPattern') {
                    // The context to check next for `({a: foo} = obj)`.
                    return { type: 'check', node: parent.parent };
                }
                return null;
            },
            ArrayPattern: ({ node, parent }) => {
                if (parent.elements.includes(node)) {
                    // The context to check next for `([foo] = obj)`.
                    return { type: 'check', node: parent };
                }
                return null;
            },
            RestElement: ({ node, parent }) => {
                if (parent.argument === node && parent.parent) {
                    // The context to check next for `({...foo} = obj)`.
                    return {
                        type: 'check',
                        node: parent.parent
                    };
                }
                return null;
            },
            SvelteDirective: ({ node, parent }) => {
                if (parent.kind !== 'Binding') {
                    return null;
                }
                if (parent.shorthand || parent.expression === node) {
                    return {
                        type: 'reassign',
                        node: parent
                    };
                }
                return null;
            }
        };
        /**
         * Returns the reassign information for the given expression node if it has a reassign.
         */
        function getReassignData(expr) {
            let pathNodes = [];
            let node = expr;
            let parent;
            while ((parent = node.parent)) {
                const check = CHECK_REASSIGN[parent.type];
                if (!check) {
                    return null;
                }
                const result = check({ node, parent, pathNodes });
                if (!result) {
                    return null;
                }
                pathNodes = result.pathNodes || pathNodes;
                if (result.type === 'reassign') {
                    return {
                        node: result.node,
                        pathNodes
                    };
                }
                node = result.node;
            }
            return null;
        }
        return {
            SvelteReactiveStatement(node) {
                if (node.body.type !== 'ExpressionStatement' ||
                    node.body.expression.type !== 'AssignmentExpression' ||
                    node.body.expression.operator !== '=') {
                    return;
                }
                const assignment = node.body.expression;
                for (const variable of toplevelScope.variables) {
                    if (!variable.defs.some((def) => def.node === assignment)) {
                        continue;
                    }
                    for (const reference of variable.references) {
                        const id = reference.identifier;
                        if ((assignment.left.range[0] <= id.range[0] &&
                            id.range[1] <= assignment.left.range[1]) ||
                            id.type === 'JSXIdentifier') {
                            continue;
                        }
                        const reassign = getReassignData(id);
                        if (!reassign) {
                            continue;
                        }
                        // Suppresses reporting if the props option is set to `false` and reassigned to properties.
                        if (!props && reassign.pathNodes.length > 0)
                            continue;
                        context.report({
                            node: reassign.node,
                            messageId: reassign.pathNodes.length === 0
                                ? 'assignmentToReactiveValue'
                                : 'assignmentToReactiveValueProp',
                            data: {
                                name: id.name
                            }
                        });
                    }
                }
            }
        };
    }
});
