"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-restricted-html-elements', {
    meta: {
        docs: {
            description: 'disallow specific HTML elements',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: false
        },
        schema: {
            type: 'array',
            items: {
                oneOf: [
                    { type: 'string' },
                    {
                        type: 'object',
                        properties: {
                            elements: {
                                type: 'array',
                                items: {
                                    type: ['string']
                                },
                                uniqueItems: true,
                                minItems: 1
                            },
                            message: { type: 'string', minLength: 1 }
                        },
                        additionalProperties: false,
                        minItems: 1
                    }
                ]
            },
            uniqueItems: true,
            minItems: 1
        },
        messages: {},
        type: 'suggestion'
    },
    create(context) {
        return {
            SvelteElement(node) {
                if (node.kind !== 'html')
                    return;
                const { name } = node;
                if (name.type !== 'SvelteName')
                    return;
                for (const option of context.options) {
                    const message = option.message || `Unexpected use of forbidden HTML element ${name.name}.`;
                    const elements = option.elements || [option];
                    for (const element of elements) {
                        if (element === name.name) {
                            context.report({
                                message,
                                node: node.startTag
                            });
                        }
                    }
                }
            }
        };
    }
});
