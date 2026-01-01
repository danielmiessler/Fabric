"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_kit_1 = require("../utils/svelte-kit");
const EXPECTED_PROP_NAMES = ['data', 'errors', 'form', 'snapshot'];
exports.default = (0, utils_1.createRule)('valid-prop-names-in-kit-pages', {
    meta: {
        docs: {
            description: 'disallow props other than data or errors in SvelteKit page components.',
            category: 'Possible Errors',
            // TODO Switch to recommended in the major version.
            recommended: false
        },
        schema: [],
        messages: {
            unexpected: 'disallow props other than data or errors in SvelteKit page components.'
        },
        type: 'problem'
    },
    create(context) {
        if (!(0, svelte_kit_1.isKitPageComponent)(context))
            return {};
        let isScript = false;
        return {
            // <script>
            'Program > SvelteScriptElement > SvelteStartTag': (node) => {
                // except for <script context="module">
                isScript = !node.attributes.some((a) => a.type === 'SvelteAttribute' &&
                    a.key.name === 'context' &&
                    a.value.some((v) => v.type === 'SvelteLiteral' && v.value === 'module'));
            },
            // </script>
            'Program > SvelteScriptElement:exit': () => {
                isScript = false;
            },
            'ExportNamedDeclaration > VariableDeclaration > VariableDeclarator': (node) => {
                if (!isScript)
                    return;
                // export let foo
                if (node.id.type === 'Identifier') {
                    if (!EXPECTED_PROP_NAMES.includes(node.id.name)) {
                        context.report({
                            node,
                            loc: node.loc,
                            messageId: 'unexpected'
                        });
                    }
                    return;
                }
                // export let { xxx, yyy } = zzz
                if (node.id.type !== 'ObjectPattern')
                    return;
                for (const p of node.id.properties) {
                    if (p.type === 'Property' &&
                        p.value.type === 'Identifier' &&
                        !EXPECTED_PROP_NAMES.includes(p.value.name)) {
                        context.report({
                            node: p.value,
                            loc: p.value.loc,
                            messageId: 'unexpected'
                        });
                    }
                }
            }
        };
    }
});
