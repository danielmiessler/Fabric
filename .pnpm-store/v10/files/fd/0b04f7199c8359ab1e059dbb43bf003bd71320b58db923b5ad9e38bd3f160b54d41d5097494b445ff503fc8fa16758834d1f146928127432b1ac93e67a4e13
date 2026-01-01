import type { RuleContext, SourceCode } from '../types';
/**
 * Returns an extended instance of `context.sourceCode` or the result of `context.getSourceCode()`.
 * Extended instances can use new APIs such as `getScope(node)` even with old ESLint.
 */
export declare function getSourceCode(context: RuleContext): SourceCode;
/**
 * Gets the value of `context.filename`, but for older ESLint it returns the result of `context.getFilename()`.
 */
export declare function getFilename(context: RuleContext): string;
/**
 * Gets the value of `context.physicalFilename`,
 * but for older ESLint it returns the result of `context.getPhysicalFilename()`.
 * Versions older than v7.28.0 return a value guessed from the result of `context.getFilename()`,
 * but it may be incorrect.
 */
export declare function getPhysicalFilename(context: RuleContext): string;
/**
 * Gets the value of `context.cwd`, but for older ESLint it returns the result of `context.getCwd()`.
 * Versions older than v6.6.0 return a value from the result of `process.cwd()`.
 */
export declare function getCwd(context: RuleContext): string;
