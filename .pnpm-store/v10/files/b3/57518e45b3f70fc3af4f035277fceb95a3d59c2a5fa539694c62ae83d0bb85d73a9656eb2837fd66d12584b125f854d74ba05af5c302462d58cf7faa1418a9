"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const EVENTS_TYPE_NAME = '$$Events';
exports.default = (0, utils_1.createRule)('experimental-require-strict-events', {
    meta: {
        docs: {
            description: 'require the strictEvents attribute on `<script>` tags',
            category: 'Experimental',
            recommended: false
        },
        schema: [],
        messages: {
            missingStrictEvents: `The component must have the strictEvents attribute on its <script> tag or it must define the $$Events interface.`
        },
        type: 'suggestion'
    },
    create(context) {
        let isTs = false;
        let hasAttribute = false;
        let hasDeclaredEvents = false;
        let scriptNode;
        return {
            SvelteScriptElement(node) {
                const lang = (0, ast_utils_1.getLangValue)(node)?.toLowerCase();
                isTs = lang === 'ts' || lang === 'typescript';
                hasAttribute = (0, ast_utils_1.findAttribute)(node, 'strictEvents') !== null;
                scriptNode = node;
            },
            TSInterfaceDeclaration(node) {
                if (node.id.name === EVENTS_TYPE_NAME) {
                    hasDeclaredEvents = true;
                }
            },
            TSTypeAliasDeclaration(node) {
                if (node.id.name === EVENTS_TYPE_NAME) {
                    hasDeclaredEvents = true;
                }
            },
            'Program:exit'() {
                if (isTs && !hasAttribute && !hasDeclaredEvents) {
                    context.report({
                        node: scriptNode,
                        messageId: 'missingStrictEvents'
                    });
                }
            }
        };
    }
});
