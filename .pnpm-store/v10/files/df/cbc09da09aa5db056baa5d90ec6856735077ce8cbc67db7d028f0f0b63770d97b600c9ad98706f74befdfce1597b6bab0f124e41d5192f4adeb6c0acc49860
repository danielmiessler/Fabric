"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.toRegExp = toRegExp;
exports.isRegExp = isRegExp;
const RE_REGEXP_STR = /^\/(.+)\/([A-Za-z]*)$/u;
/**
 * Convert a string to the `RegExp`.
 * Normal strings (e.g. `"foo"`) is converted to `/^foo$/` of `RegExp`.
 * Strings like `"/^foo/i"` are converted to `/^foo/i` of `RegExp`.
 *
 * @param {string} string The string to convert.
 * @returns {RegExp} Returns the `RegExp`.
 */
function toRegExp(string) {
    const parts = RE_REGEXP_STR.exec(string);
    if (parts) {
        return new RegExp(parts[1], parts[2]);
    }
    return { test: (s) => s === string };
}
/**
 * Checks whether given string is regexp string
 * @param {string} string
 * @returns {boolean}
 */
function isRegExp(string) {
    return Boolean(RE_REGEXP_STR.test(string));
}
