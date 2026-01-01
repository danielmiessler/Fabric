import { Rule, SourceCode } from 'eslint';

/**
 * Returns an extended instance of `context.sourceCode` or the result of `context.getSourceCode()`.
 * Extended instances can use new APIs such as `getScope(node)` even with old ESLint.
 */
declare function getSourceCode(context: Rule.RuleContext): SourceCode;

/**
 * Gets the value of `context.cwd`, but for older ESLint it returns the result of `context.getCwd()`.
 * Versions older than v6.6.0 return a value from the result of `process.cwd()`.
 */
declare function getCwd(context: Rule.RuleContext): string;

/**
 * Gets the value of `context.filename`, but for older ESLint it returns the result of `context.getFilename()`.
 */
declare function getFilename(context: Rule.RuleContext): string;

/**
 * Gets the value of `context.physicalFilename`,
 * but for older ESLint it returns the result of `context.getPhysicalFilename()`.
 * Versions older than v7.28.0 return a value guessed from the result of `context.getFilename()`,
 * but it may be incorrect.
 */
declare function getPhysicalFilename(context: Rule.RuleContext): string;

export { getCwd, getFilename, getPhysicalFilename, getSourceCode };
