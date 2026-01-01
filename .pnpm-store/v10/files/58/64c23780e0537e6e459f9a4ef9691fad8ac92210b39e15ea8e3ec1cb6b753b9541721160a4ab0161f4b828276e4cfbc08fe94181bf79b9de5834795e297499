"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getSourceCode = getSourceCode;
exports.getFilename = getFilename;
exports.getPhysicalFilename = getPhysicalFilename;
exports.getCwd = getCwd;
const eslint_compat_utils_1 = require("eslint-compat-utils");
// export function getSourceCode(context: RuleContext): SourceCode;
// export function getSourceCode(context: Rule.RuleContext): ESLintSourceCode;
/**
 * Returns an extended instance of `context.sourceCode` or the result of `context.getSourceCode()`.
 * Extended instances can use new APIs such as `getScope(node)` even with old ESLint.
 */
function getSourceCode(context) {
    return (0, eslint_compat_utils_1.getSourceCode)(context);
}
/**
 * Gets the value of `context.filename`, but for older ESLint it returns the result of `context.getFilename()`.
 */
function getFilename(context) {
    return (0, eslint_compat_utils_1.getFilename)(context);
}
/**
 * Gets the value of `context.physicalFilename`,
 * but for older ESLint it returns the result of `context.getPhysicalFilename()`.
 * Versions older than v7.28.0 return a value guessed from the result of `context.getFilename()`,
 * but it may be incorrect.
 */
function getPhysicalFilename(context) {
    return (0, eslint_compat_utils_1.getPhysicalFilename)(context);
}
/**
 * Gets the value of `context.cwd`, but for older ESLint it returns the result of `context.getCwd()`.
 * Versions older than v6.6.0 return a value from the result of `process.cwd()`.
 */
function getCwd(context) {
    return (0, eslint_compat_utils_1.getCwd)(context);
}
