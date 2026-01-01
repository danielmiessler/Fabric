"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-at-debug-tags', {
    meta: {
        docs: {
            description: 'disallow the use of `{@debug}`',
            category: 'Best Practices',
            recommended: true,
            default: 'warn'
        },
        schema: [],
        messages: {
            unexpected: 'Unexpected `{@debug}`.'
        },
        type: 'problem'
    },
    create(context) {
        return {
            SvelteDebugTag(node) {
                context.report({
                    node,
                    messageId: 'unexpected'
                });
            }
        };
    }
});
