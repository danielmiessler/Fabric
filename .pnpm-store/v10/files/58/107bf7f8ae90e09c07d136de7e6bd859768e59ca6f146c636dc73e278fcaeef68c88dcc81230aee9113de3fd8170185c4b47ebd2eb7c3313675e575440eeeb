"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-reactive-functions', {
    meta: {
        docs: {
            description: "it's not necessary to define functions in reactive statements",
            category: 'Best Practices',
            recommended: false
        },
        hasSuggestions: true,
        schema: [],
        messages: {
            noReactiveFns: `Do not create functions inside reactive statements unless absolutely necessary.`,
            fixReactiveFns: `Move the function out of the reactive statement`
        },
        type: 'suggestion' // "problem", or "layout",
    },
    create(context) {
        return {
            // $: foo = () => { ... }
            [`SvelteReactiveStatement > ExpressionStatement > AssignmentExpression > :function`](node) {
                // Move upwards to include the entire label
                const parent = node.parent?.parent?.parent;
                if (!parent) {
                    return false;
                }
                const source = (0, compat_1.getSourceCode)(context);
                return context.report({
                    node: parent,
                    loc: parent.loc,
                    messageId: 'noReactiveFns',
                    suggest: [
                        {
                            messageId: 'fixReactiveFns',
                            fix(fixer) {
                                const tokens = source.getFirstTokens(parent, {
                                    includeComments: false,
                                    count: 3
                                });
                                const noExtraSpace = source.isSpaceBetweenTokens(tokens[1], tokens[2]);
                                // Replace the entire reactive label with "const"
                                return fixer.replaceTextRange([tokens[0].range[0], tokens[1].range[1]], noExtraSpace ? 'const' : 'const ');
                            }
                        }
                    ]
                });
            }
        };
    }
});
