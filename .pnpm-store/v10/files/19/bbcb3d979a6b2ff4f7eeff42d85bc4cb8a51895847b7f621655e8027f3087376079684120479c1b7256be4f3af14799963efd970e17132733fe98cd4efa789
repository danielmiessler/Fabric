"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('require-each-key', {
    meta: {
        docs: {
            description: 'require keyed `{#each}` block',
            category: 'Best Practices',
            recommended: false
        },
        schema: [],
        messages: { expectedKey: 'Each block should have a key' },
        type: 'suggestion'
    },
    create(context) {
        return {
            SvelteEachBlock(node) {
                if (node.key == null) {
                    context.report({
                        node,
                        messageId: 'expectedKey'
                    });
                }
            }
        };
    }
});
