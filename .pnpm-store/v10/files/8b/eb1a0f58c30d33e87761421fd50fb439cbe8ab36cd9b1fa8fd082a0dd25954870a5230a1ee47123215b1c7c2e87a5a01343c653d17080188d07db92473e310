"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
/** Checks wether the given attr node is target="_blank"  */
function isTargetBlank(node) {
    return node.key.name === 'target' && (0, ast_utils_1.getStaticAttributeValue)(node) === '_blank';
}
/** Checks wether the given element node has secure rel="..." */
function hasSecureRel(node, allowReferrer) {
    const attr = (0, ast_utils_1.findAttribute)(node, 'rel');
    if (attr) {
        const tags = [];
        for (const value of attr.value) {
            if (value.type === 'SvelteLiteral') {
                tags.push(...value.value.toLowerCase().split(' '));
            }
        }
        return tags && tags.includes('noopener') && (allowReferrer || tags.includes('noreferrer'));
    }
    return false;
}
/** Checks wether the given element node has external link */
function hasExternalLink(node) {
    return node.attributes.some((attr) => attr.type === 'SvelteAttribute' &&
        attr.key.name === 'href' &&
        attr.value.length >= 1 &&
        attr.value[0].type === 'SvelteLiteral' &&
        /^(?:\w+:|\/\/)/.test(attr.value[0].value));
}
/** Checks wether the given element node has dynamic link */
function hasDynamicLink(node) {
    const attr = (0, ast_utils_1.findAttribute)(node, 'href');
    if (attr) {
        return attr.value.some((v) => v.type === 'SvelteMustacheTag');
    }
    return Boolean((0, ast_utils_1.findShorthandAttribute)(node, 'href')) || Boolean((0, ast_utils_1.findBindDirective)(node, 'href'));
}
exports.default = (0, utils_1.createRule)('no-target-blank', {
    meta: {
        docs: {
            description: 'disallow `target="_blank"` attribute without `rel="noopener noreferrer"`',
            category: 'Security Vulnerability',
            recommended: false
        },
        schema: [
            {
                type: 'object',
                properties: {
                    allowReferrer: {
                        type: 'boolean'
                    },
                    enforceDynamicLinks: {
                        enum: ['always', 'never']
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            disallow: 'Using target="_blank" without rel="noopener noreferrer" is a security risk.'
        },
        type: 'problem'
    },
    create(context) {
        const configuration = context.options[0] || {};
        const allowReferrer = Boolean(configuration.allowReferrer) || false;
        const enforceDynamicLinks = configuration.enforceDynamicLinks || 'always';
        return {
            SvelteAttribute(node) {
                if (!isTargetBlank(node) || hasSecureRel(node.parent, allowReferrer)) {
                    return;
                }
                const hasDangerHref = hasExternalLink(node.parent) ||
                    (enforceDynamicLinks === 'always' && hasDynamicLink(node.parent));
                if (hasDangerHref) {
                    context.report({
                        node,
                        message: 'Using target="_blank" without rel="noopener noreferrer" is a security risk.'
                    });
                }
            }
        };
    }
});
