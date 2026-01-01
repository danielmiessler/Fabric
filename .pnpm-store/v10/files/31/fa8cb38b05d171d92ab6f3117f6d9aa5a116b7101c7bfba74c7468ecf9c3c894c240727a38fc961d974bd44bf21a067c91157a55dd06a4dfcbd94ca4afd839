"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const known_css_properties_1 = require("known-css-properties");
const regexp_1 = require("../utils/regexp");
const css_utils_1 = require("../utils/css-utils");
exports.default = (0, utils_1.createRule)('no-unknown-style-directive-property', {
    meta: {
        docs: {
            description: 'disallow unknown `style:property`',
            category: 'Possible Errors',
            recommended: true
        },
        schema: [
            {
                type: 'object',
                properties: {
                    ignoreProperties: {
                        type: 'array',
                        items: {
                            type: 'string'
                        },
                        uniqueItems: true,
                        minItems: 1
                    },
                    ignorePrefixed: { type: 'boolean' }
                },
                additionalProperties: false
            }
        ],
        messages: {
            unknown: "Unexpected unknown style directive property '{{property}}'."
        },
        type: 'problem'
    },
    create(context) {
        const ignoreProperties = [...(context.options[0]?.ignoreProperties ?? [])].map(regexp_1.toRegExp);
        const ignorePrefixed = context.options[0]?.ignorePrefixed ?? true;
        const knownProperties = new Set(known_css_properties_1.all);
        /** Checks whether given name is valid */
        function validName(name) {
            return (name.startsWith('--') ||
                knownProperties.has(name) ||
                ignoreProperties.some((r) => r.test(name)) ||
                (ignorePrefixed && (0, css_utils_1.hasVendorPrefix)(name)));
        }
        return {
            SvelteStyleDirective(node) {
                const prop = node.key.name;
                if (validName(prop.name)) {
                    return;
                }
                context.report({
                    node: prop,
                    messageId: 'unknown',
                    data: {
                        property: prop.name
                    }
                });
            }
        };
    }
});
