"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isWhitespace = isWhitespace;
exports.isNotWhitespace = isNotWhitespace;
/**
 * Check whether the given token is a whitespace.
 */
function isWhitespace(token) {
    return (token != null &&
        ((token.type === 'HTMLText' && !token.value.trim()) ||
            (token.type === 'JSXText' && !token.value.trim())));
}
/**
 * Check whether the given token is a not whitespace.
 */
function isNotWhitespace(token) {
    return (token != null &&
        (token.type !== 'HTMLText' || Boolean(token.value.trim())) &&
        (token.type !== 'JSXText' || Boolean(token.value.trim())));
}
