"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('shorthand-attribute', {
    meta: {
        docs: {
            description: 'enforce use of shorthand syntax in attribute',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: true
        },
        fixable: 'code',
        schema: [
            {
                type: 'object',
                properties: {
                    prefer: { enum: ['always', 'never'] }
                },
                additionalProperties: false
            }
        ],
        messages: {
            expectedShorthand: 'Expected shorthand attribute.',
            expectedRegular: 'Expected regular attribute syntax.'
        },
        type: 'layout' // "problem", or "layout",
    },
    create(context) {
        const sourceCode = (0, compat_1.getSourceCode)(context);
        const always = context.options[0]?.prefer !== 'never';
        return always
            ? {
                SvelteAttribute(node) {
                    if (node.value.length !== 1) {
                        return;
                    }
                    const value = node.value[0];
                    if (value.type !== 'SvelteMustacheTag' || value.expression.type !== 'Identifier') {
                        return;
                    }
                    if (node.key.name === value.expression.name) {
                        context.report({
                            node,
                            messageId: 'expectedShorthand',
                            *fix(fixer) {
                                const quoteAndRange = (0, ast_utils_1.getAttributeValueQuoteAndRange)(node, sourceCode);
                                if (!quoteAndRange) {
                                    return;
                                }
                                if (quoteAndRange.quote === 'double' || quoteAndRange.quote === 'single') {
                                    yield fixer.removeRange([node.range[0], quoteAndRange.firstToken.range[1]]);
                                    yield fixer.remove(quoteAndRange.lastToken);
                                }
                                else {
                                    yield fixer.removeRange([node.range[0], quoteAndRange.range[0]]);
                                }
                            }
                        });
                    }
                }
            }
            : {
                SvelteShorthandAttribute(node) {
                    context.report({
                        node,
                        messageId: 'expectedRegular',
                        fix(fixer) {
                            return fixer.insertTextBefore(node, `${node.key.name}=`);
                        }
                    });
                }
            };
    }
});
