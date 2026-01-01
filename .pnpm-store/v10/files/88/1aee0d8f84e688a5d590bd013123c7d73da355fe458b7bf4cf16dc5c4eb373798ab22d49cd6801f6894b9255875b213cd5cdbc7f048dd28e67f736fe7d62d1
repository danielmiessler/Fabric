"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_store_1 = require("./reference-helpers/svelte-store");
exports.default = (0, utils_1.createRule)('require-store-callbacks-use-set-param', {
    meta: {
        docs: {
            description: 'store callbacks must use `set` param',
            category: 'Possible Errors',
            recommended: false
        },
        schema: [],
        messages: {
            unexpected: 'Store callbacks must use `set` param.'
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            Program() {
                for (const { node } of (0, svelte_store_1.extractStoreReferences)(context, ['readable', 'writable'])) {
                    const [_, fn] = node.arguments;
                    if (!fn || (fn.type !== 'ArrowFunctionExpression' && fn.type !== 'FunctionExpression')) {
                        continue;
                    }
                    const param = fn.params[0];
                    if (!param || (param.type === 'Identifier' && param.name !== 'set')) {
                        context.report({
                            node: fn,
                            loc: fn.loc,
                            messageId: 'unexpected'
                        });
                    }
                }
            }
        };
    }
});
