"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-svelte-internal', {
    meta: {
        docs: {
            description: 'svelte/internal will be removed in Svelte 6.',
            category: 'Best Practices',
            // TODO Switch to recommended in the major version.
            // recommended: true,
            recommended: false
        },
        schema: [],
        messages: {
            unexpected: 'Using svelte/internal is prohibited. This will be removed in Svelte 6.'
        },
        type: 'problem'
    },
    create(context) {
        function report(node) {
            context.report({
                node,
                messageId: 'unexpected'
            });
        }
        function isSvelteInternal(value) {
            return value === 'svelte/internal' || value.startsWith('svelte/internal/');
        }
        return {
            ImportDeclaration(node) {
                if (node.source && isSvelteInternal(node.source.value)) {
                    report(node);
                }
            },
            ImportExpression(node) {
                if (node.source &&
                    node.source.type === 'Literal' &&
                    typeof node.source.value === 'string' &&
                    isSvelteInternal(node.source.value)) {
                    report(node);
                }
            },
            ExportNamedDeclaration(node) {
                if (node.source && isSvelteInternal(node.source.value)) {
                    report(node);
                }
            },
            ExportAllDeclaration(node) {
                if (node.source && isSvelteInternal(node.source.value)) {
                    report(node);
                }
            }
        };
    }
});
