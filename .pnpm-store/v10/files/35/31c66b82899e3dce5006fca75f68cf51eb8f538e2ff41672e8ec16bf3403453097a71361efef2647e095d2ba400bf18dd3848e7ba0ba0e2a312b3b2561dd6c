"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getFirstAndLastTokens = getFirstAndLastTokens;
exports.isBeginningOfLine = isBeginningOfLine;
exports.isBeginningOfElement = isBeginningOfElement;
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const ast_1 = require("./ast");
/**
 * Get the first and last tokens of the given node.
 * If the node is parenthesized, this gets the outermost parentheses.
 * If the node have whitespace at the start and the end, they will be skipped.
 */
function getFirstAndLastTokens(sourceCode, node, borderOffset = 0) {
    let firstToken = sourceCode.getFirstToken(node);
    let lastToken = sourceCode.getLastToken(node);
    // Get the outermost left parenthesis if it's parenthesized.
    let left, right;
    while ((left = sourceCode.getTokenBefore(firstToken)) != null &&
        (right = sourceCode.getTokenAfter(lastToken)) != null &&
        (0, eslint_utils_1.isOpeningParenToken)(left) &&
        (0, eslint_utils_1.isClosingParenToken)(right) &&
        borderOffset <= left.range[0]) {
        firstToken = left;
        lastToken = right;
    }
    while ((0, ast_1.isWhitespace)(firstToken) && firstToken.range[0] < lastToken.range[0]) {
        firstToken = sourceCode.getTokenAfter(firstToken);
    }
    while ((0, ast_1.isWhitespace)(lastToken) && firstToken.range[0] < lastToken.range[0]) {
        lastToken = sourceCode.getTokenBefore(lastToken);
    }
    return { firstToken, lastToken };
}
/**
 * Check whether the given node or token is the beginning of a line.
 */
function isBeginningOfLine(sourceCode, node) {
    const prevToken = sourceCode.getTokenBefore(node, {
        includeComments: false,
        filter: ast_1.isNotWhitespace
    });
    return !prevToken || prevToken.loc.end.line < node.loc.start.line;
}
/**
 * Check whether the given node is the beginning of element.
 */
function isBeginningOfElement(node) {
    if (node.parent.type === 'SvelteElement' ||
        node.parent.type === 'SvelteAwaitCatchBlock' ||
        node.parent.type === 'SvelteAwaitPendingBlock' ||
        node.parent.type === 'SvelteAwaitThenBlock' ||
        node.parent.type === 'SvelteEachBlock' ||
        node.parent.type === 'SvelteElseBlock' ||
        node.parent.type === 'SvelteIfBlock' ||
        node.parent.type === 'SvelteKeyBlock' ||
        node.parent.type === 'SvelteSnippetBlock' ||
        node.parent.type === 'SvelteStyleElement') {
        return node.parent.children[0] === node;
    }
    if (node.parent.type === 'Program') {
        return node.parent.body[0] === node;
    }
    return assertNever(node.parent);
}
/**
 * Throws an error when invoked.
 */
function assertNever(value) {
    throw new Error(`This part of the code should never be reached but ${value} made it through.`);
}
