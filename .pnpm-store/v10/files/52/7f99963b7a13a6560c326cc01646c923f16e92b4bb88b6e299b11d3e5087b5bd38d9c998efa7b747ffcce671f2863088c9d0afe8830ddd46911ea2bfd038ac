"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const PHRASES = {
    ObjectExpression: 'object',
    ArrayExpression: 'array',
    ArrowFunctionExpression: 'function',
    FunctionExpression: 'function',
    ClassExpression: 'class'
};
exports.default = (0, utils_1.createRule)('no-object-in-text-mustaches', {
    meta: {
        docs: {
            description: 'disallow objects in text mustache interpolation',
            category: 'Possible Errors',
            recommended: true
        },
        schema: [],
        messages: {
            unexpected: 'Unexpected {{phrase}} in text mustache interpolation.'
        },
        type: 'problem' // "problem", or "layout",
    },
    create(context) {
        return {
            SvelteMustacheTag(node) {
                const { expression } = node;
                if (expression.type !== 'ObjectExpression' &&
                    expression.type !== 'ArrayExpression' &&
                    expression.type !== 'ArrowFunctionExpression' &&
                    expression.type !== 'FunctionExpression' &&
                    expression.type !== 'ClassExpression') {
                    return;
                }
                if (node.parent.type === 'SvelteAttribute') {
                    if (node.parent.value.length === 1) {
                        // Maybe props
                        return;
                    }
                }
                context.report({
                    node,
                    messageId: 'unexpected',
                    data: {
                        phrase: PHRASES[expression.type]
                    }
                });
            }
        };
    }
});
