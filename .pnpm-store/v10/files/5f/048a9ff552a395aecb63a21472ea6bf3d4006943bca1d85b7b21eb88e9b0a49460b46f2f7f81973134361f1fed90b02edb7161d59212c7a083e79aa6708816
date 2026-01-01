"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const indent_helpers_1 = require("./indent-helpers");
exports.default = (0, utils_1.createRule)('indent', {
    meta: {
        docs: {
            description: 'enforce consistent indentation',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: true
        },
        fixable: 'whitespace',
        schema: [
            {
                type: 'object',
                properties: {
                    indent: {
                        anyOf: [{ type: 'integer', minimum: 1 }, { enum: ['tab'] }]
                    },
                    indentScript: { type: 'boolean' },
                    switchCase: { type: 'integer', minimum: 0 },
                    alignAttributesVertically: { type: 'boolean' },
                    ignoredNodes: {
                        type: 'array',
                        items: {
                            allOf: [
                                { type: 'string' },
                                { not: { type: 'string', pattern: ':exit$' } },
                                { not: { type: 'string', pattern: '^\\s*$' } }
                            ]
                        },
                        uniqueItems: true,
                        additionalItems: false
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            unexpectedChar: 'Expected {{expected}} character, but found {{actual}} character.',
            unexpectedIndentation: 'Expected indentation of {{expectedIndent}} {{expectedUnit}}{{expectedIndentPlural}} but found {{actualIndent}} {{actualUnit}}{{actualIndentPlural}}.'
        },
        type: 'layout' // "problem",
    },
    create(context) {
        return (0, indent_helpers_1.defineVisitor)(context, {});
    }
});
