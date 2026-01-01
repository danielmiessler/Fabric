"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const eslint_core_1 = require("../utils/eslint-core");
const coreRule = (0, eslint_core_1.getCoreRule)('no-inner-declarations');
exports.default = (0, utils_1.createRule)('no-inner-declarations', {
    meta: {
        ...coreRule.meta,
        docs: {
            description: 'disallow variable or `function` declarations in nested blocks',
            category: 'Extension Rules',
            recommended: true,
            extensionRule: 'no-inner-declarations'
        },
        fixable: coreRule.meta.fixable,
        schema: coreRule.meta.schema,
        messages: coreRule.meta.messages,
        type: coreRule.meta.type
    },
    create(context) {
        return (0, eslint_core_1.defineWrapperListener)(coreRule, context, {
            createListenerProxy(coreListener) {
                return (0, eslint_core_1.buildProxyListener)(coreListener, (node) => {
                    return (0, eslint_core_1.getProxyNode)(node, {
                        get parent() {
                            if (node.parent?.type === 'SvelteScriptElement') {
                                return node.parent.parent;
                            }
                            return node.parent;
                        }
                    });
                });
            }
        });
    }
});
