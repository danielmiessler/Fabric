"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('html-closing-bracket-spacing', {
    meta: {
        docs: {
            description: "require or disallow a space before tag's closing brackets",
            category: 'Stylistic Issues',
            conflictWithPrettier: true,
            recommended: false
        },
        schema: [
            {
                type: 'object',
                properties: {
                    startTag: {
                        enum: ['always', 'never', 'ignore']
                    },
                    endTag: {
                        enum: ['always', 'never', 'ignore']
                    },
                    selfClosingTag: {
                        enum: ['always', 'never', 'ignore']
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            expectedSpace: "Expected space before '>', but not found.",
            unexpectedSpace: "Expected no space before '>', but found."
        },
        fixable: 'whitespace',
        type: 'layout'
    },
    create(ctx) {
        const options = {
            startTag: 'never',
            endTag: 'never',
            selfClosingTag: 'always',
            ...ctx.options[0]
        };
        const src = ctx.getSourceCode();
        /**
         * Returns true if string contains newline characters
         */
        function containsNewline(string) {
            return string.includes('\n');
        }
        /**
         * Report
         */
        function report(node, shouldHave) {
            const tagSrc = src.getText(node);
            const match = /(\s*)\/?>$/.exec(tagSrc);
            const end = node.range[1];
            const start = node.range[1] - match[0].length;
            const loc = {
                start: src.getLocFromIndex(start),
                end: src.getLocFromIndex(end)
            };
            ctx.report({
                loc,
                messageId: shouldHave ? 'expectedSpace' : 'unexpectedSpace',
                *fix(fixer) {
                    if (shouldHave) {
                        yield fixer.insertTextBeforeRange([start, end], ' ');
                    }
                    else {
                        const spaces = match[1];
                        yield fixer.removeRange([start, start + spaces.length]);
                    }
                }
            });
        }
        return {
            'SvelteStartTag, SvelteEndTag'(node) {
                const tagType = node.type === 'SvelteEndTag'
                    ? 'endTag'
                    : node.selfClosing
                        ? 'selfClosingTag'
                        : 'startTag';
                if (options[tagType] === 'ignore')
                    return;
                const tagSrc = src.getText(node);
                const match = /(\s*)\/?>$/.exec(tagSrc);
                if (containsNewline(match[1]))
                    return;
                if (options[tagType] === 'always' && !match[1]) {
                    report(node, true);
                }
                else if (options[tagType] === 'never' && match[1]) {
                    report(node, false);
                }
            }
        };
    }
});
