"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
/**
 * Strip quotes string
 */
function stripQuotes(text) {
    if ((text.startsWith('"') || text.startsWith("'") || text.startsWith('`')) &&
        text.endsWith(text[0])) {
        return text.slice(1, -1);
    }
    return null;
}
exports.default = (0, utils_1.createRule)('no-useless-mustaches', {
    meta: {
        docs: {
            description: 'disallow unnecessary mustache interpolations',
            category: 'Best Practices',
            recommended: false
        },
        fixable: 'code',
        schema: [
            {
                type: 'object',
                properties: {
                    ignoreIncludesComment: {
                        type: 'boolean'
                    },
                    ignoreStringEscape: {
                        type: 'boolean'
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            unexpected: 'Unexpected mustache interpolation with a string literal value.'
        },
        type: 'suggestion' // "problem",
    },
    create(context) {
        const opts = context.options[0] || {};
        const ignoreIncludesComment = Boolean(opts.ignoreIncludesComment);
        const ignoreStringEscape = Boolean(opts.ignoreStringEscape);
        const sourceCode = (0, compat_1.getSourceCode)(context);
        /**
         * Report if the value expression is string literals
         * @param node the node to check
         */
        function verify(node) {
            if (node.kind === 'raw') {
                return;
            }
            const { expression } = node;
            let strValue, rawValue;
            if (expression.type === 'Literal') {
                if (typeof expression.value !== 'string') {
                    return;
                }
                strValue = expression.value;
                rawValue = sourceCode.getText(expression).slice(1, -1);
            }
            else if (expression.type === 'TemplateLiteral') {
                if (expression.expressions.length > 0) {
                    return;
                }
                strValue = expression.quasis[0].value.cooked;
                rawValue = expression.quasis[0].value.raw;
            }
            else {
                return;
            }
            const hasComment = sourceCode
                .getTokens(node, { includeComments: true })
                .some((t) => t.type === 'Block' || t.type === 'Line');
            if (ignoreIncludesComment && hasComment) {
                return;
            }
            let hasEscape = false;
            if (rawValue !== strValue) {
                // check escapes
                const chars = [...rawValue];
                let c = chars.shift();
                while (c) {
                    if (c === '\\') {
                        c = chars.shift();
                        if (c == null ||
                            // ignore "\\", '"', "'", "`" and "$"
                            'nrvtbfux'.includes(c)) {
                            // has useful escape.
                            hasEscape = true;
                            break;
                        }
                    }
                    c = chars.shift();
                }
            }
            if (ignoreStringEscape && hasEscape) {
                return;
            }
            context.report({
                node,
                messageId: 'unexpected',
                fix(fixer) {
                    if (hasComment || hasEscape) {
                        // cannot fix
                        return null;
                    }
                    const text = stripQuotes(sourceCode.getText(expression));
                    if (text == null) {
                        // unknowns
                        return null;
                    }
                    if (text.includes('\n') || /^\s|\s$/u.test(text)) {
                        // It doesn't autofix because another rule like indent or eol space might remove spaces.
                        return null;
                    }
                    const unescaped = text.replace(/\\([\s\S])/g, '$1');
                    if (node.parent.type === 'SvelteAttribute' ||
                        node.parent.type === 'SvelteStyleDirective') {
                        const div = sourceCode.text.slice(node.parent.key.range[1], node.parent.value[0].range[0]);
                        if (!div.endsWith('"') && !div.endsWith("'")) {
                            return [
                                fixer.insertTextBefore(node.parent.value[0], '"'),
                                fixer.replaceText(node, unescaped.replace(/"/gu, '&quot;')),
                                fixer.insertTextAfter(node.parent.value[node.parent.value.length - 1], '"')
                            ];
                        }
                        return fixer.replaceText(node, unescaped);
                    }
                    return fixer.replaceText(node, unescaped.replace(/</gu, '&lt;').replace(/>/gu, '&gt;'));
                }
            });
        }
        return {
            SvelteMustacheTag: verify
        };
    }
});
