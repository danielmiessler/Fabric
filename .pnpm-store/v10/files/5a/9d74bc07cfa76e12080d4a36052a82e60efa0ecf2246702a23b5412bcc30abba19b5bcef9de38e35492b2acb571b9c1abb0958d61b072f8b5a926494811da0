"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_store_1 = require("./reference-helpers/svelte-store");
exports.default = (0, utils_1.createRule)('require-stores-init', {
    meta: {
        docs: {
            description: 'require initial value in store',
            category: 'Best Practices',
            recommended: false
        },
        schema: [],
        messages: {
            storeDefaultValue: `Always set a default value for svelte stores.`
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            Program() {
                for (const { node, name } of (0, svelte_store_1.extractStoreReferences)(context)) {
                    const minArgs = name === 'writable' || name === 'readable' ? 1 : name === 'derived' ? 3 : 0;
                    if (node.arguments.length >= minArgs ||
                        node.arguments.some((arg) => arg.type === 'SpreadElement')) {
                        continue;
                    }
                    context.report({
                        node,
                        messageId: 'storeDefaultValue'
                    });
                }
            }
        };
    }
});
