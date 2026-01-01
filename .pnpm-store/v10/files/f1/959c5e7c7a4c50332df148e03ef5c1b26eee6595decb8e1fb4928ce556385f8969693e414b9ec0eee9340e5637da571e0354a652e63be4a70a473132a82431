"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const utils_1 = require("../utils");
const ast_utils_1 = require("../utils/ast-utils");
const compat_1 = require("../utils/compat");
const VALUE_SCHEMA = { enum: ['never', 'always'] };
/**
 * Normalize options.
 */
function parseOptions(options) {
    return {
        textExpressions: options?.textExpressions || 'never',
        attributesAndProps: options?.attributesAndProps || 'never',
        directiveExpressions: options?.directiveExpressions || 'never',
        tags: {
            openingBrace: options?.tags?.openingBrace || 'never',
            closingBrace: options?.tags?.closingBrace || 'never'
        }
    };
}
exports.default = (0, utils_1.createRule)('mustache-spacing', {
    meta: {
        docs: {
            description: 'enforce unified spacing in mustache',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: true
        },
        fixable: 'code',
        schema: [
            {
                type: 'object',
                properties: {
                    textExpressions: VALUE_SCHEMA,
                    attributesAndProps: VALUE_SCHEMA,
                    directiveExpressions: VALUE_SCHEMA,
                    tags: {
                        type: 'object',
                        properties: {
                            openingBrace: VALUE_SCHEMA,
                            closingBrace: {
                                enum: ['never', 'always', 'always-after-expression']
                            }
                        },
                        additionalProperties: false
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            expectedOpening: "Expected 1 space after '{', but not found.",
            expectedClosing: "Expected 1 space before '}', but not found.",
            unexpectedOpening: "Expected no space after '{', but found.",
            unexpectedClosing: "Expected no space before '}', but found."
        },
        type: 'layout'
    },
    create(context) {
        const options = parseOptions(context.options[0]);
        const sourceCode = (0, compat_1.getSourceCode)(context);
        /** Verify */
        function verifyBraces(openingBrace, closingBrace, openingOption, closingOption, hasExpression) {
            const firstToken = sourceCode.getTokenAfter(openingBrace, {
                includeComments: true
            });
            if (openingOption === 'always') {
                if (openingBrace.range[1] === firstToken.range[0]) {
                    context.report({
                        node: openingBrace,
                        messageId: 'expectedOpening',
                        fix: (fixer) => fixer.insertTextAfter(openingBrace, ' ')
                    });
                }
            }
            else if (openingBrace.range[1] !== firstToken.range[0]) {
                context.report({
                    loc: {
                        start: openingBrace.loc.start,
                        end: firstToken.loc.start
                    },
                    messageId: 'unexpectedOpening',
                    fix: (fixer) => fixer.removeRange([openingBrace.range[1], firstToken.range[0]])
                });
            }
            if (!closingBrace) {
                return;
            }
            const lastToken = sourceCode.getTokenBefore(closingBrace, {
                includeComments: true
            });
            if (closingOption === 'always' ||
                (closingOption === 'always-after-expression' && hasExpression)) {
                if (closingBrace.range[0] === lastToken.range[1]) {
                    context.report({
                        node: closingBrace,
                        messageId: 'expectedClosing',
                        fix: (fixer) => fixer.insertTextBefore(closingBrace, ' ')
                    });
                }
            }
            else if (closingBrace.range[0] !== lastToken.range[1]) {
                context.report({
                    loc: {
                        start: lastToken.loc.end,
                        end: closingBrace.loc.end
                    },
                    messageId: 'unexpectedClosing',
                    fix: (fixer) => fixer.removeRange([lastToken.range[1], closingBrace.range[0]])
                });
            }
        }
        /** Verify for {expr} */
        function verifyExpression(node, option) {
            const mustacheTokens = (0, ast_utils_1.getMustacheTokens)(node, sourceCode);
            if (!mustacheTokens) {
                return;
            }
            verifyBraces(mustacheTokens.openToken, mustacheTokens.closeToken, option, option);
        }
        return {
            SvelteMustacheTag(node) {
                if (node.kind === 'raw') {
                    const mustacheTokens = (0, ast_utils_1.getMustacheTokens)(node, sourceCode);
                    verifyBraces(mustacheTokens.openToken, mustacheTokens.closeToken, options.tags.openingBrace, options.tags.closingBrace, true);
                    return;
                }
                let option;
                if (node.parent.type === 'SvelteAttribute') {
                    option = options.attributesAndProps;
                }
                else if (node.parent.type === 'SvelteStyleDirective') {
                    option = options.directiveExpressions;
                }
                else {
                    option = options.textExpressions;
                }
                verifyExpression(node, option);
            },
            SvelteShorthandAttribute(node) {
                verifyExpression(node, options.attributesAndProps);
            },
            SvelteSpreadAttribute(node) {
                verifyExpression(node, options.attributesAndProps);
            },
            SvelteDirective(node) {
                verifyExpression(node, options.directiveExpressions);
            },
            SvelteSpecialDirective(node) {
                verifyExpression(node, options.directiveExpressions);
            },
            SvelteDebugTag(node) {
                const mustacheTokens = (0, ast_utils_1.getMustacheTokens)(node, sourceCode);
                verifyBraces(mustacheTokens.openToken, mustacheTokens.closeToken, options.tags.openingBrace, options.tags.closingBrace, true);
            },
            SvelteRenderTag(node) {
                const mustacheTokens = (0, ast_utils_1.getMustacheTokens)(node, sourceCode);
                verifyBraces(mustacheTokens.openToken, mustacheTokens.closeToken, options.tags.openingBrace, options.tags.closingBrace, true);
            },
            SvelteIfBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockClosingToken = sourceCode.getTokenAfter(node.expression, {
                    includeComments: false,
                    filter: eslint_utils_1.isClosingBraceToken
                });
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, true);
                if (node.elseif) {
                    return;
                }
                const closeBlockClosingToken = sourceCode.getLastToken(node);
                const closeBlockOpeningToken = sourceCode.getTokenBefore(closeBlockClosingToken, {
                    includeComments: false,
                    filter: eslint_utils_1.isOpeningBraceToken
                });
                verifyBraces(closeBlockOpeningToken, closeBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, false);
            },
            SvelteElseBlock(node) {
                if (node.elseif) {
                    return;
                }
                const openToken = sourceCode.getFirstToken(node);
                const closeToken = sourceCode.getTokenAfter(openToken, {
                    includeComments: false,
                    filter: eslint_utils_1.isClosingBraceToken
                });
                verifyBraces(openToken, closeToken, options.tags.openingBrace, options.tags.closingBrace, false);
            },
            SvelteEachBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockClosingToken = sourceCode.getTokenAfter(node.key || node.index || node.context, {
                    includeComments: false,
                    filter: eslint_utils_1.isClosingBraceToken
                });
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, true);
                const closeBlockClosingToken = sourceCode.getLastToken(node);
                const closeBlockOpeningToken = sourceCode.getTokenBefore(closeBlockClosingToken, {
                    includeComments: false,
                    filter: eslint_utils_1.isOpeningBraceToken
                });
                verifyBraces(closeBlockOpeningToken, closeBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, false);
            },
            SvelteKeyBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockClosingToken = sourceCode.getTokenAfter(node.expression, {
                    includeComments: false,
                    filter: eslint_utils_1.isClosingBraceToken
                });
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, true);
                const closeBlockClosingToken = sourceCode.getLastToken(node);
                const closeBlockOpeningToken = sourceCode.getTokenBefore(closeBlockClosingToken, {
                    includeComments: false,
                    filter: eslint_utils_1.isOpeningBraceToken
                });
                verifyBraces(closeBlockOpeningToken, closeBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, false);
            },
            SvelteAwaitBlock(node) {
                const closeBlockClosingToken = sourceCode.getLastToken(node);
                const closeBlockOpeningToken = sourceCode.getTokenBefore(closeBlockClosingToken, {
                    includeComments: false,
                    filter: eslint_utils_1.isOpeningBraceToken
                });
                verifyBraces(closeBlockOpeningToken, closeBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, false);
            },
            SvelteAwaitPendingBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockClosingToken = sourceCode.getTokenAfter(node.parent.expression, {
                    includeComments: false,
                    filter: eslint_utils_1.isClosingBraceToken
                });
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, true);
            },
            SvelteAwaitThenBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockLast = node.value || (node.awaitThen ? node.parent.expression : null);
                const openBlockClosingToken = openBlockLast
                    ? sourceCode.getTokenAfter(openBlockLast, {
                        includeComments: false,
                        filter: eslint_utils_1.isClosingBraceToken
                    })
                    : null;
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, Boolean(openBlockClosingToken &&
                    openBlockLast &&
                    openBlockClosingToken === sourceCode.getTokenAfter(openBlockLast)));
            },
            SvelteAwaitCatchBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockLast = node.error || (node.awaitCatch ? node.parent.expression : null);
                const openBlockClosingToken = openBlockLast
                    ? sourceCode.getTokenAfter(openBlockLast, {
                        includeComments: false,
                        filter: eslint_utils_1.isClosingBraceToken
                    })
                    : null;
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, Boolean(openBlockClosingToken &&
                    openBlockLast &&
                    openBlockClosingToken === sourceCode.getTokenAfter(openBlockLast)));
            },
            SvelteSnippetBlock(node) {
                const openBlockOpeningToken = sourceCode.getFirstToken(node);
                const openBlockClosingToken = sourceCode.getTokenAfter(node.params[node.params.length - 1] || node.id, {
                    includeComments: false,
                    filter: eslint_utils_1.isClosingBraceToken
                });
                verifyBraces(openBlockOpeningToken, openBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, true);
                const closeBlockClosingToken = sourceCode.getLastToken(node);
                const closeBlockOpeningToken = sourceCode.getTokenBefore(closeBlockClosingToken, {
                    includeComments: false,
                    filter: eslint_utils_1.isOpeningBraceToken
                });
                verifyBraces(closeBlockOpeningToken, closeBlockClosingToken, options.tags.openingBrace, options.tags.closingBrace, false);
            }
        };
    }
});
