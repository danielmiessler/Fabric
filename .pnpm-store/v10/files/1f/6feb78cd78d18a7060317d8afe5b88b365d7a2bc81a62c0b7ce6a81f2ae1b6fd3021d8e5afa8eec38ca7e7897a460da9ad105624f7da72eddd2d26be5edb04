"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-inspect', {
    meta: {
        docs: {
            description: 'Warns against the use of `$inspect` directive',
            category: 'Best Practices',
            // TODO: Enable recommended in major version
            recommended: false,
            default: 'warn'
        },
        schema: [],
        messages: {
            unexpected: 'Do not use $inspect directive'
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            Identifier(node) {
                if (node.name !== '$inspect') {
                    return;
                }
                context.report({ messageId: 'unexpected', node });
            }
        };
    }
});
