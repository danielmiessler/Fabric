"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_store_1 = require("./reference-helpers/svelte-store");
exports.default = (0, utils_1.createRule)('no-store-async', {
    meta: {
        docs: {
            description: 'disallow using async/await inside svelte stores because it causes issues with the auto-unsubscribing features',
            category: 'Possible Errors',
            // TODO Switch to recommended in the major version.
            // recommended: true,
            recommended: false,
            default: 'error'
        },
        schema: [],
        messages: {
            unexpected: 'Do not pass async functions to svelte stores.'
        },
        type: 'problem'
    },
    create(context) {
        return {
            Program() {
                for (const { node } of (0, svelte_store_1.extractStoreReferences)(context)) {
                    const [, fn] = node.arguments;
                    if (!fn ||
                        (fn.type !== 'ArrowFunctionExpression' && fn.type !== 'FunctionExpression') ||
                        !fn.async) {
                        continue;
                    }
                    const start = fn.loc.start;
                    context.report({
                        node: fn,
                        loc: {
                            start,
                            end: {
                                line: start.line,
                                column: start.column + 5
                            }
                        },
                        messageId: 'unexpected'
                    });
                }
            }
        };
    }
});
