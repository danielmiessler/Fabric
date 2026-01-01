"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('block-lang', {
    meta: {
        docs: {
            description: 'disallows the use of languages other than those specified in the configuration for the lang attribute of `<script>` and `<style>` blocks.',
            category: 'Best Practices',
            recommended: false
        },
        schema: [
            {
                type: 'object',
                properties: {
                    enforceScriptPresent: {
                        type: 'boolean'
                    },
                    enforceStylePresent: {
                        type: 'boolean'
                    },
                    script: {
                        oneOf: [
                            {
                                type: ['string', 'null']
                            },
                            {
                                type: 'array',
                                items: {
                                    type: ['string', 'null']
                                },
                                minItems: 1
                            }
                        ]
                    },
                    style: {
                        oneOf: [
                            {
                                type: ['string', 'null']
                            },
                            {
                                type: 'array',
                                items: {
                                    type: ['string', 'null']
                                },
                                minItems: 1
                            }
                        ]
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {},
        type: 'suggestion'
    },
    create(context) {
        if (!(0, compat_1.getSourceCode)(context).parserServices.isSvelte) {
            return {};
        }
        const enforceScriptPresent = context.options[0]?.enforceScriptPresent ?? false;
        const enforceStylePresent = context.options[0]?.enforceStylePresent ?? false;
        const scriptOption = context.options[0]?.script ?? null;
        const allowedScriptLangs = Array.isArray(scriptOption)
            ? scriptOption
            : [scriptOption];
        const scriptNodes = [];
        const styleOption = context.options[0]?.style ?? null;
        const allowedStyleLangs = Array.isArray(styleOption)
            ? styleOption
            : [styleOption];
        const styleNodes = [];
        return {
            SvelteScriptElement(node) {
                scriptNodes.push(node);
            },
            SvelteStyleElement(node) {
                styleNodes.push(node);
            },
            'Program:exit'() {
                if (scriptNodes.length === 0 && enforceScriptPresent) {
                    context.report({
                        loc: { line: 1, column: 1 },
                        message: `The <script> block should be present and its lang attribute should be ${prettyPrintLangs(allowedScriptLangs)}.`
                    });
                }
                for (const scriptNode of scriptNodes) {
                    if (!allowedScriptLangs.includes((0, ast_utils_1.getLangValue)(scriptNode)?.toLowerCase() ?? null)) {
                        context.report({
                            node: scriptNode,
                            message: `The lang attribute of the <script> block should be ${prettyPrintLangs(allowedScriptLangs)}.`
                        });
                    }
                }
                if (styleNodes.length === 0 && enforceStylePresent) {
                    context.report({
                        loc: { line: 1, column: 1 },
                        message: `The <style> block should be present and its lang attribute should be ${prettyPrintLangs(allowedStyleLangs)}.`
                    });
                }
                for (const styleNode of styleNodes) {
                    if (!allowedStyleLangs.includes((0, ast_utils_1.getLangValue)(styleNode)?.toLowerCase() ?? null)) {
                        context.report({
                            node: styleNode,
                            message: `The lang attribute of the <style> block should be ${prettyPrintLangs(allowedStyleLangs)}.`
                        });
                    }
                }
            }
        };
    }
});
/**
 * Prints the list of allowed languages, with special handling of the `null` option.
 */
function prettyPrintLangs(langs) {
    const hasNull = langs.includes(null);
    const nonNullLangs = langs.filter((lang) => lang !== null).map((lang) => `"${lang}"`);
    if (nonNullLangs.length === 0) {
        // No special behaviour for `hasNull`, because that can never happen.
        return 'omitted';
    }
    const hasNullText = hasNull ? 'either omitted or ' : '';
    const nonNullText = nonNullLangs.length === 1 ? nonNullLangs[0] : `one of ${nonNullLangs.join(', ')}`;
    return hasNullText + nonNullText;
}
