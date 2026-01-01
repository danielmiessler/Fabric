import { Lexer } from './Lexer.js';
import { Parser } from './Parser.js';
import { Tokenizer } from './Tokenizer.js';
import { Renderer } from './Renderer.js';
import { TextRenderer } from './TextRenderer.js';
import { Slugger } from './Slugger.js';
import { Hooks } from './Hooks.js';
import { Marked } from './Instance.js';
import { changeDefaults, getDefaults, defaults } from './defaults.js';

const markedInstance = new Marked(defaults);

/**
 * Marked
 */
export function marked(src, opt, callback) {
  return markedInstance.parse(src, opt, callback);
}

/**
 * Options
 */

marked.options =
marked.setOptions = function(opt) {
  markedInstance.setOptions(opt);
  marked.defaults = markedInstance.defaults;
  changeDefaults(marked.defaults);
  return marked;
};

marked.getDefaults = getDefaults;

marked.defaults = defaults;

/**
 * Use Extension
 */

marked.use = function(...args) {
  markedInstance.use(...args);
  marked.defaults = markedInstance.defaults;
  changeDefaults(marked.defaults);
  return marked;
};

/**
 * Run callback for every token
 */

marked.walkTokens = function(tokens, callback) {
  return markedInstance.walkTokens(tokens, callback);
};

/**
 * Parse Inline
 * @param {string} src
 */
marked.parseInline = markedInstance.parseInline;

/**
 * Expose
 */
marked.Parser = Parser;
marked.parser = Parser.parse;
marked.Renderer = Renderer;
marked.TextRenderer = TextRenderer;
marked.Lexer = Lexer;
marked.lexer = Lexer.lex;
marked.Tokenizer = Tokenizer;
marked.Slugger = Slugger;
marked.Hooks = Hooks;
marked.parse = marked;

export const options = marked.options;
export const setOptions = marked.setOptions;
export const use = marked.use;
export const walkTokens = marked.walkTokens;
export const parseInline = marked.parseInline;
export const parse = marked;
export const parser = Parser.parse;
export const lexer = Lexer.lex;
export { defaults, getDefaults } from './defaults.js';
export { Lexer } from './Lexer.js';
export { Parser } from './Parser.js';
export { Tokenizer } from './Tokenizer.js';
export { Renderer } from './Renderer.js';
export { TextRenderer } from './TextRenderer.js';
export { Slugger } from './Slugger.js';
export { Hooks } from './Hooks.js';
export { Marked } from './Instance.js';
