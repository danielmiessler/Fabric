"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-at-html-tags', {
    meta: {
        docs: {
            description: 'disallow use of `{@html}` to prevent XSS attack',
            category: 'Security Vulnerability',
            recommended: true
        },
        schema: [],
        messages: {
            unexpected: '`{@html}` can lead to XSS attack.'
        },
        type: 'suggestion' // "problem",
    },
    create(context) {
        return {
            'SvelteMustacheTag[kind=raw]'(node) {
                context.report({
                    node,
                    messageId: 'unexpected'
                });
            }
        };
    }
});
