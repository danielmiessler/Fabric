"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.extractLeadingComments = extractLeadingComments;
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const compat_1 = require("../../utils/compat");
/** Extract comments */
function extractLeadingComments(context, node) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const beforeToken = sourceCode.getTokenBefore(node, {
        includeComments: false,
        filter(token) {
            if ((0, eslint_utils_1.isOpeningParenToken)(token)) {
                return false;
            }
            const astToken = token;
            if (astToken.type === 'HTMLText') {
                return false;
            }
            return astToken.type !== 'HTMLComment';
        }
    });
    if (beforeToken) {
        return sourceCode
            .getTokensBetween(beforeToken, node, { includeComments: true })
            .filter(isComment);
    }
    return sourceCode.getTokensBefore(node, { includeComments: true }).filter(isComment);
}
/** Checks whether given token is comment token */
function isComment(token) {
    return token.type === 'HTMLComment' || token.type === 'Block' || token.type === 'Line';
}
