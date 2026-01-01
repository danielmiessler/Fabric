"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-ignored-unsubscribe', {
    meta: {
        docs: {
            description: 'disallow ignoring the unsubscribe method returned by the `subscribe()` on Svelte stores.',
            category: 'Best Practices',
            recommended: false
        },
        fixable: undefined,
        hasSuggestions: false,
        messages: {
            forbidden: 'Ignoring returned value of the subscribe method is forbidden.'
        },
        schema: [],
        type: 'problem'
    },
    create: (context) => {
        return {
            "ExpressionStatement > CallExpression > MemberExpression.callee[property.name='subscribe']": (node) => {
                context.report({
                    messageId: 'forbidden',
                    node: node.property
                });
            }
        };
    }
});
