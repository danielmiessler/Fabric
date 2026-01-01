"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('spaced-html-comment', {
    meta: {
        docs: {
            description: 'enforce consistent spacing after the `<!--` and before the `-->` in a HTML comment',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: false
        },
        fixable: 'whitespace',
        schema: [
            {
                enum: ['always', 'never']
            }
        ],
        messages: {
            expectedSpaceBefore: "Expected space or tab before '-->' in comment.",
            expectedSpaceAfter: "Expected space or tab after '<!--' in comment.",
            unexpectedSpaceBefore: "Unexpected space or tab before '-->' in comment.",
            unexpectedSpaceAfter: "Unexpected space or tab after '<!--' in comment."
        },
        type: 'layout'
    },
    create(context) {
        // Unless the first option is never, require a space
        const requireSpace = context.options[0] !== 'never';
        return {
            SvelteHTMLComment(node) {
                if (!node.value.trim()) {
                    return;
                }
                if (requireSpace) {
                    if (/^\S/u.test(node.value)) {
                        context.report({
                            node,
                            messageId: 'expectedSpaceAfter',
                            fix(fixer) {
                                return fixer.insertTextBeforeRange([node.range[0] + 4, node.range[1]], ' ');
                            }
                        });
                    }
                    if (/\S$/u.test(node.value)) {
                        context.report({
                            node,
                            messageId: 'expectedSpaceBefore',
                            fix(fixer) {
                                return fixer.insertTextAfterRange([node.range[0], node.range[1] - 3], ' ');
                            }
                        });
                    }
                }
                else {
                    const beginSpaces = /^[^\S\n\r]/u.exec(node.value)?.[0];
                    if (beginSpaces) {
                        context.report({
                            node,
                            messageId: 'unexpectedSpaceAfter',
                            fix(fixer) {
                                return fixer.removeRange([
                                    node.range[0] + 4,
                                    node.range[0] + 4 + beginSpaces.length
                                ]);
                            }
                        });
                    }
                    const endSpaces = /(?<=\S)[^\S\n\r]$/u.exec(node.value)?.[0];
                    if (endSpaces) {
                        context.report({
                            node,
                            messageId: 'unexpectedSpaceBefore',
                            fix(fixer) {
                                return fixer.removeRange([node.range[1] - 3 - endSpaces.length, node.range[1] - 3]);
                            }
                        });
                    }
                }
            }
        };
    }
});
