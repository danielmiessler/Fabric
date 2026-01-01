import type { ESLintExtendedProgram, ESLintProgram } from ".";
import type * as tsESLintParser from "@typescript-eslint/parser";
type TSESLintParser = typeof tsESLintParser;
/**
 * The type of basic ESLint custom parser.
 * e.g. espree
 */
export type BasicParserObject = {
    parse(code: string, options: any): ESLintProgram;
    parseForESLint: undefined;
};
/**
 * The type of ESLint custom parser enhanced for ESLint.
 * e.g. @babel/eslint-parser, @typescript-eslint/parser
 */
export type EnhancedParserObject = {
    parseForESLint(code: string, options: any): ESLintExtendedProgram;
    parse: undefined;
};
/**
 * The type of ESLint (custom) parsers.
 */
export type ParserObject = EnhancedParserObject | BasicParserObject;
/** Checks whether given object is ParserObject */
export declare function isParserObject(value: unknown): value is ParserObject;
/** Checks whether given object is EnhancedParserObject */
export declare function isEnhancedParserObject(value: unknown): value is EnhancedParserObject;
/** Checks whether given object is BasicParserObject */
export declare function isBasicParserObject(value: unknown): value is BasicParserObject;
/** Checks whether given object maybe "@typescript-eslint/parser" */
export declare function maybeTSESLintParserObject(value: unknown): value is TSESLintParser;
/** Checks whether given object is "@typescript-eslint/parser" */
export declare function isTSESLintParserObject(value: unknown): value is TSESLintParser;
export {};
