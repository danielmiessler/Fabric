"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-reactive-literals', {
    meta: {
        docs: {
            description: "don't assign literal values in reactive statements",
            category: 'Best Practices',
            recommended: false
        },
        hasSuggestions: true,
        schema: [],
        messages: {
            noReactiveLiterals: `Do not assign literal values inside reactive statements unless absolutely necessary.`,
            fixReactiveLiteral: `Move the literal out of the reactive statement into an assignment`
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            [`SvelteReactiveStatement > ExpressionStatement > AssignmentExpression:matches(${[
                // $: foo = "foo";
                // $: foo = 1;
                `[right.type="Literal"]`,
                // $: foo = [];
                `[right.type="ArrayExpression"][right.elements.length=0]`,
                // $: foo = {};
                `[right.type="ObjectExpression"][right.properties.length=0]`
            ].join(',')})`](node) {
                // Move upwards to include the entire reactive statement
                const parent = node.parent?.parent;
                if (!parent) {
                    return false;
                }
                const source = (0, compat_1.getSourceCode)(context);
                return context.report({
                    node: parent,
                    loc: parent.loc,
                    messageId: 'noReactiveLiterals',
                    suggest: [
                        {
                            messageId: 'fixReactiveLiteral',
                            fix(fixer) {
                                return [
                                    // Insert "let" + whatever was in there
                                    fixer.insertTextBefore(parent, `let ${source.getText(node)}`),
                                    // Remove the original reactive statement
                                    fixer.remove(parent)
                                ];
                            }
                        }
                    ]
                });
            }
        };
    }
});
