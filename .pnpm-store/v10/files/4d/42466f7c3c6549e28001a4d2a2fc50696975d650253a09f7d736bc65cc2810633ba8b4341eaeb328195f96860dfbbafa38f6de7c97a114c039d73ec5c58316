"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-extra-reactive-curlies', {
    meta: {
        docs: {
            description: 'disallow wrapping single reactive statements in curly braces',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: false
        },
        hasSuggestions: true,
        schema: [],
        messages: {
            extraCurlies: `Do not wrap reactive statements in curly braces unless necessary.`,
            removeExtraCurlies: `Remove the unnecessary curly braces.`
        },
        type: 'suggestion'
    },
    create(context) {
        return {
            // $: { foo = "bar"; }
            [`SvelteReactiveStatement > BlockStatement[body.length=1]`]: (node) => {
                const source = (0, compat_1.getSourceCode)(context);
                return context.report({
                    node,
                    loc: node.loc,
                    messageId: 'extraCurlies',
                    suggest: [
                        {
                            messageId: 'removeExtraCurlies',
                            fix(fixer) {
                                const tokens = source.getTokens(node, { includeComments: true });
                                // Remove everything up to the second token, and the entire last token since
                                // those are known to be "{" and "}"
                                return [
                                    fixer.removeRange([tokens[0].range[0], tokens[1].range[0]]),
                                    fixer.removeRange([
                                        tokens[tokens.length - 2].range[1],
                                        tokens[tokens.length - 1].range[1]
                                    ])
                                ];
                            }
                        }
                    ]
                });
            }
        };
    }
});
