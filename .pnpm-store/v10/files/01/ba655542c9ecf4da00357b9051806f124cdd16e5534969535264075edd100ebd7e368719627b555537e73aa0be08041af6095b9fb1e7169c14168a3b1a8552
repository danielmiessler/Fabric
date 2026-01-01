"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
exports.default = (0, utils_1.createRule)('button-has-type', {
    meta: {
        docs: {
            description: 'disallow usage of button without an explicit type attribute',
            category: 'Best Practices',
            recommended: false
        },
        schema: [
            {
                type: 'object',
                properties: {
                    button: {
                        type: 'boolean'
                    },
                    submit: {
                        type: 'boolean'
                    },
                    reset: {
                        type: 'boolean'
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            missingTypeAttribute: 'Missing an explicit type attribute for button.',
            invalidTypeAttribute: '{{value}} is an invalid value for button type attribute.',
            forbiddenTypeAttribute: '{{value}} is a forbidden value for button type attribute.',
            emptyTypeAttribute: 'A value must be set for button type attribute.'
        },
        type: 'suggestion' // "problem",
    },
    create(context) {
        const configuration = {
            button: true,
            submit: true,
            reset: true,
            ...(context.options[0] ?? {})
        };
        /**
         * Checks whether given text is known button type
         */
        function isButtonType(type) {
            return type === 'button' || type === 'submit' || type === 'reset';
        }
        /**
         * Report
         */
        function report(node, messageId, data = {}) {
            context.report({
                node,
                messageId,
                data
            });
        }
        /**
         * Validate attribute
         */
        function validateAttribute(attribute) {
            if (attribute.value.length === 0) {
                report(attribute, 'emptyTypeAttribute');
                return;
            }
            const strValue = (0, ast_utils_1.getStaticAttributeValue)(attribute);
            if (strValue == null) {
                return;
            }
            if (!isButtonType(strValue)) {
                report(attribute, 'invalidTypeAttribute', { value: strValue });
            }
            else if (!configuration[strValue]) {
                report(attribute, 'forbiddenTypeAttribute', { value: strValue });
            }
        }
        /**
         * @param {VDirective} directive
         */
        function validateDirective(directive) {
            if (!directive.expression) {
                report(directive, 'emptyTypeAttribute');
            }
        }
        return {
            "SvelteElement[name.name='button'] > SvelteStartTag"(node) {
                const typeAttr = (0, ast_utils_1.findAttribute)(node, 'type');
                if (typeAttr) {
                    validateAttribute(typeAttr);
                    return;
                }
                const typeDir = (0, ast_utils_1.findBindDirective)(node, 'type');
                if (typeDir) {
                    validateDirective(typeDir);
                    return;
                }
                const typeShortAttr = (0, ast_utils_1.findShorthandAttribute)(node, 'type');
                if (typeShortAttr) {
                    return;
                }
                for (const attr of node.attributes) {
                    if (attr.type === 'SvelteSpreadAttribute') {
                        return;
                    }
                }
                report(node, 'missingTypeAttribute');
            }
        };
    }
});
