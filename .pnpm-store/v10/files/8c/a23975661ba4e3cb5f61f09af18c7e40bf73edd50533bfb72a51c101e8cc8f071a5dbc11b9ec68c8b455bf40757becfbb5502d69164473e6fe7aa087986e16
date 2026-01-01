"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('require-event-dispatcher-types', {
    meta: {
        docs: {
            description: 'require type parameters for `createEventDispatcher`',
            category: 'Best Practices',
            recommended: false
        },
        schema: [],
        messages: {
            missingTypeParameter: `Type parameters missing for the \`createEventDispatcher\` function call.`
        },
        type: 'suggestion'
    },
    create(context) {
        let isTs = false;
        return {
            SvelteScriptElement(node) {
                const lang = (0, ast_utils_1.getLangValue)(node)?.toLowerCase();
                if (lang === 'ts' || lang === 'typescript') {
                    isTs = true;
                }
            },
            'Program:exit'() {
                if (!isTs) {
                    return;
                }
                const referenceTracker = new eslint_utils_1.ReferenceTracker((0, compat_1.getSourceCode)(context).scopeManager.globalScope);
                for (const { node: n } of referenceTracker.iterateEsmReferences({
                    svelte: {
                        [eslint_utils_1.ReferenceTracker.ESM]: true,
                        createEventDispatcher: {
                            [eslint_utils_1.ReferenceTracker.CALL]: true
                        }
                    }
                })) {
                    const node = n;
                    if ((node.typeArguments ?? node.typeParameters) === undefined) {
                        context.report({ node, messageId: 'missingTypeParameter' });
                    }
                }
            }
        };
    }
});
