export type OnigurumaToEsOptions = {
    accuracy?: "default" | "strict";
    avoidSubclass?: boolean;
    flags?: string;
    global?: boolean;
    hasIndices?: boolean;
    rules?: {
        allowOrphanBackrefs?: boolean;
        asciiWordBoundaries?: boolean;
        captureGroup?: boolean;
        recursionLimit?: number;
        singleline?: boolean;
    };
    target?: "auto" | "ES2025" | "ES2024" | "ES2018";
    verbose?: boolean;
};
import { EmulatedRegExp } from './subclass.js';
/**
@typedef {{
  accuracy?: keyof Accuracy;
  avoidSubclass?: boolean;
  flags?: string;
  global?: boolean;
  hasIndices?: boolean;
  rules?: {
    allowOrphanBackrefs?: boolean;
    asciiWordBoundaries?: boolean;
    captureGroup?: boolean;
    recursionLimit?: number;
    singleline?: boolean;
  };
  target?: keyof Target;
  verbose?: boolean;
}} OnigurumaToEsOptions
*/
/**
Accepts an Oniguruma pattern and returns the details needed to construct an equivalent JavaScript `RegExp`.
@param {string} pattern Oniguruma regex pattern.
@param {OnigurumaToEsOptions} [options]
@returns {{
  pattern: string;
  flags: string;
  options?: import('./subclass.js').EmulatedRegExpOptions;
}}
*/
export function toDetails(pattern: string, options?: OnigurumaToEsOptions): {
    pattern: string;
    flags: string;
    options?: import("./subclass.js").EmulatedRegExpOptions;
};
/**
Returns an Oniguruma AST generated from an Oniguruma pattern.
@param {string} pattern Oniguruma regex pattern.
@param {{
  flags?: string;
  rules?: {
    captureGroup?: boolean;
  };
}} [options]
@returns {import('./parse.js').OnigurumaAst}
*/
export function toOnigurumaAst(pattern: string, options?: {
    flags?: string;
    rules?: {
        captureGroup?: boolean;
    };
}): import("./parse.js").OnigurumaAst;
/**
Accepts an Oniguruma pattern and returns an equivalent JavaScript `RegExp`.
@param {string} pattern Oniguruma regex pattern.
@param {OnigurumaToEsOptions} [options]
@returns {RegExp | EmulatedRegExp}
*/
export function toRegExp(pattern: string, options?: OnigurumaToEsOptions): RegExp | EmulatedRegExp;
export { EmulatedRegExp };
