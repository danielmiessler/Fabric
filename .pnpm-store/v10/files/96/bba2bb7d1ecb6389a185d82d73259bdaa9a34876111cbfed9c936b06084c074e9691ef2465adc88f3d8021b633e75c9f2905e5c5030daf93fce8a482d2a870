"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
function getPhrase(lineBreaks) {
    switch (lineBreaks) {
        case 0: {
            return 'no line breaks';
        }
        case 1: {
            return '1 line break';
        }
        default: {
            return `${lineBreaks} line breaks`;
        }
    }
}
function getExpectedLineBreaks(node, options, type) {
    const isSelfClosingTag = node.type === 'SvelteStartTag' && node.selfClosing;
    if (isSelfClosingTag && options.selfClosingTag && options.selfClosingTag[type]) {
        return options.selfClosingTag[type] === 'always' ? 1 : 0;
    }
    return options[type] === 'always' ? 1 : 0;
}
function getSelfClosingData(sourceCode, node, options) {
    const tokens = sourceCode.getTokens(node);
    const closingToken = tokens[tokens.length - 2];
    if (closingToken.value !== '/') {
        return null;
    }
    const prevToken = sourceCode.getTokenBefore(closingToken);
    const type = node.loc.start.line === prevToken.loc.end.line ? 'singleline' : 'multiline';
    const expectedLineBreaks = getExpectedLineBreaks(node, options, type);
    const actualLineBreaks = closingToken.loc.start.line - prevToken.loc.end.line;
    return { actualLineBreaks, expectedLineBreaks, startToken: prevToken, endToken: closingToken };
}
function getNodeData(sourceCode, node, options) {
    const closingToken = sourceCode.getLastToken(node);
    if (closingToken.value !== '>') {
        return null;
    }
    const prevToken = sourceCode.getTokenBefore(closingToken);
    const type = node.loc.start.line === prevToken.loc.end.line ? 'singleline' : 'multiline';
    const expectedLineBreaks = getExpectedLineBreaks(node, options, type);
    const actualLineBreaks = closingToken.loc.start.line - prevToken.loc.end.line;
    return { actualLineBreaks, expectedLineBreaks, startToken: prevToken, endToken: closingToken };
}
exports.default = (0, utils_1.createRule)('html-closing-bracket-new-line', {
    meta: {
        docs: {
            description: "Require or disallow a line break before tag's closing brackets",
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: true
        },
        schema: [
            {
                type: 'object',
                properties: {
                    singleline: { enum: ['always', 'never'] },
                    multiline: { enum: ['always', 'never'] },
                    selfClosingTag: {
                        type: 'object',
                        properties: {
                            singleline: { enum: ['always', 'never'] },
                            multiline: { enum: ['always', 'never'] }
                        },
                        additionalProperties: false,
                        minProperties: 1
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            expectedBeforeClosingBracket: 'Expected {{expected}} before closing bracket, but {{actual}} found.'
        },
        fixable: 'code',
        type: 'suggestion'
    },
    create(context) {
        const options = context.options[0] ?? {};
        options.singleline ?? (options.singleline = 'never');
        options.multiline ?? (options.multiline = 'always');
        const sourceCode = (0, compat_1.getSourceCode)(context);
        return {
            'SvelteStartTag, SvelteEndTag'(node) {
                const data = node.type === 'SvelteStartTag' && node.selfClosing
                    ? getSelfClosingData(sourceCode, node, options)
                    : getNodeData(sourceCode, node, options);
                if (!data) {
                    return;
                }
                const { actualLineBreaks, expectedLineBreaks, startToken, endToken } = data;
                if (actualLineBreaks !== expectedLineBreaks) {
                    // For SvelteEndTag, does not make sense to add a line break, so we only fix if there are extra line breaks
                    if (node.type === 'SvelteEndTag' && expectedLineBreaks !== 0) {
                        return;
                    }
                    context.report({
                        node,
                        loc: { start: startToken.loc.end, end: endToken.loc.start },
                        messageId: 'expectedBeforeClosingBracket',
                        data: {
                            expected: getPhrase(expectedLineBreaks),
                            actual: getPhrase(actualLineBreaks)
                        },
                        fix(fixer) {
                            const range = [startToken.range[1], endToken.range[0]];
                            const text = '\n'.repeat(expectedLineBreaks);
                            return fixer.replaceTextRange(range, text);
                        }
                    });
                }
            }
        };
    }
});
