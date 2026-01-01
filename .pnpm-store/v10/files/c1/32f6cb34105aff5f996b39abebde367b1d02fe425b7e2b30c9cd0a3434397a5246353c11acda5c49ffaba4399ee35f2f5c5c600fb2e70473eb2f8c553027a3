"use strict";
var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __getProtoOf = Object.getPrototypeOf;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __commonJS = (cb, mod) => function __require() {
  return mod || (0, cb[__getOwnPropNames(cb)[0]])((mod = { exports: {} }).exports, mod), mod.exports;
};
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toESM = (mod, isNodeMode, target) => (target = mod != null ? __create(__getProtoOf(mod)) : {}, __copyProps(
  // If the importer is in node compatibility mode or this is not an ESM
  // file that has been converted to a CommonJS file using a Babel-
  // compatible transform (i.e. "__esModule" has not been set), then set
  // "default" to the CommonJS "module.exports" for node compatibility.
  isNodeMode || !mod || !mod.__esModule ? __defProp(target, "default", { value: mod, enumerable: true }) : target,
  mod
));
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

// node_modules/.pnpm/camelcase-css@2.0.1/node_modules/camelcase-css/index.js
var require_camelcase_css = __commonJS({
  "node_modules/.pnpm/camelcase-css@2.0.1/node_modules/camelcase-css/index.js"(exports, module2) {
    "use strict";
    var pattern = /-(\w|$)/g;
    var callback = (dashChar, char) => char.toUpperCase();
    var camelCaseCSS = (property) => {
      property = property.toLowerCase();
      if (property === "float") {
        return "cssFloat";
      } else if (property.startsWith("-ms-")) {
        return property.substr(1).replace(pattern, callback);
      } else {
        return property.replace(pattern, callback);
      }
    };
    module2.exports = camelCaseCSS;
  }
});

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/objectifier.js
var require_objectifier = __commonJS({
  "node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/objectifier.js"(exports, module2) {
    var camelcase = require_camelcase_css();
    var UNITLESS = {
      boxFlex: true,
      boxFlexGroup: true,
      columnCount: true,
      flex: true,
      flexGrow: true,
      flexPositive: true,
      flexShrink: true,
      flexNegative: true,
      fontWeight: true,
      lineClamp: true,
      lineHeight: true,
      opacity: true,
      order: true,
      orphans: true,
      tabSize: true,
      widows: true,
      zIndex: true,
      zoom: true,
      fillOpacity: true,
      strokeDashoffset: true,
      strokeOpacity: true,
      strokeWidth: true
    };
    function atRule(node) {
      if (typeof node.nodes === "undefined") {
        return true;
      } else {
        return process2(node);
      }
    }
    function process2(node) {
      let name;
      let result = {};
      node.each((child) => {
        if (child.type === "atrule") {
          name = "@" + child.name;
          if (child.params)
            name += " " + child.params;
          if (typeof result[name] === "undefined") {
            result[name] = atRule(child);
          } else if (Array.isArray(result[name])) {
            result[name].push(atRule(child));
          } else {
            result[name] = [result[name], atRule(child)];
          }
        } else if (child.type === "rule") {
          let body = process2(child);
          if (result[child.selector]) {
            for (let i in body) {
              result[child.selector][i] = body[i];
            }
          } else {
            result[child.selector] = body;
          }
        } else if (child.type === "decl") {
          if (child.prop[0] === "-" && child.prop[1] === "-") {
            name = child.prop;
          } else if (child.parent && child.parent.selector === ":export") {
            name = child.prop;
          } else {
            name = camelcase(child.prop);
          }
          let value = child.value;
          if (!isNaN(child.value) && UNITLESS[name]) {
            value = parseFloat(child.value);
          }
          if (child.important)
            value += " !important";
          if (typeof result[name] === "undefined") {
            result[name] = value;
          } else if (Array.isArray(result[name])) {
            result[name].push(value);
          } else {
            result[name] = [result[name], value];
          }
        }
      });
      return result;
    }
    module2.exports = process2;
  }
});

// node_modules/.pnpm/picocolors@1.0.0/node_modules/picocolors/picocolors.js
var require_picocolors = __commonJS({
  "node_modules/.pnpm/picocolors@1.0.0/node_modules/picocolors/picocolors.js"(exports, module2) {
    var tty = require("tty");
    var isColorSupported = !("NO_COLOR" in process.env || process.argv.includes("--no-color")) && ("FORCE_COLOR" in process.env || process.argv.includes("--color") || process.platform === "win32" || tty.isatty(1) && process.env.TERM !== "dumb" || "CI" in process.env);
    var formatter = (open, close, replace = open) => (input) => {
      let string = "" + input;
      let index2 = string.indexOf(close, open.length);
      return ~index2 ? open + replaceClose(string, close, replace, index2) + close : open + string + close;
    };
    var replaceClose = (string, close, replace, index2) => {
      let start = string.substring(0, index2) + replace;
      let end = string.substring(index2 + close.length);
      let nextIndex = end.indexOf(close);
      return ~nextIndex ? start + replaceClose(end, close, replace, nextIndex) : start + end;
    };
    var createColors = (enabled = isColorSupported) => ({
      isColorSupported: enabled,
      reset: enabled ? (s) => `\x1B[0m${s}\x1B[0m` : String,
      bold: enabled ? formatter("\x1B[1m", "\x1B[22m", "\x1B[22m\x1B[1m") : String,
      dim: enabled ? formatter("\x1B[2m", "\x1B[22m", "\x1B[22m\x1B[2m") : String,
      italic: enabled ? formatter("\x1B[3m", "\x1B[23m") : String,
      underline: enabled ? formatter("\x1B[4m", "\x1B[24m") : String,
      inverse: enabled ? formatter("\x1B[7m", "\x1B[27m") : String,
      hidden: enabled ? formatter("\x1B[8m", "\x1B[28m") : String,
      strikethrough: enabled ? formatter("\x1B[9m", "\x1B[29m") : String,
      black: enabled ? formatter("\x1B[30m", "\x1B[39m") : String,
      red: enabled ? formatter("\x1B[31m", "\x1B[39m") : String,
      green: enabled ? formatter("\x1B[32m", "\x1B[39m") : String,
      yellow: enabled ? formatter("\x1B[33m", "\x1B[39m") : String,
      blue: enabled ? formatter("\x1B[34m", "\x1B[39m") : String,
      magenta: enabled ? formatter("\x1B[35m", "\x1B[39m") : String,
      cyan: enabled ? formatter("\x1B[36m", "\x1B[39m") : String,
      white: enabled ? formatter("\x1B[37m", "\x1B[39m") : String,
      gray: enabled ? formatter("\x1B[90m", "\x1B[39m") : String,
      bgBlack: enabled ? formatter("\x1B[40m", "\x1B[49m") : String,
      bgRed: enabled ? formatter("\x1B[41m", "\x1B[49m") : String,
      bgGreen: enabled ? formatter("\x1B[42m", "\x1B[49m") : String,
      bgYellow: enabled ? formatter("\x1B[43m", "\x1B[49m") : String,
      bgBlue: enabled ? formatter("\x1B[44m", "\x1B[49m") : String,
      bgMagenta: enabled ? formatter("\x1B[45m", "\x1B[49m") : String,
      bgCyan: enabled ? formatter("\x1B[46m", "\x1B[49m") : String,
      bgWhite: enabled ? formatter("\x1B[47m", "\x1B[49m") : String
    });
    module2.exports = createColors();
    module2.exports.createColors = createColors;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/tokenize.js
var require_tokenize = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/tokenize.js"(exports, module2) {
    "use strict";
    var SINGLE_QUOTE = "'".charCodeAt(0);
    var DOUBLE_QUOTE = '"'.charCodeAt(0);
    var BACKSLASH = "\\".charCodeAt(0);
    var SLASH = "/".charCodeAt(0);
    var NEWLINE = "\n".charCodeAt(0);
    var SPACE = " ".charCodeAt(0);
    var FEED = "\f".charCodeAt(0);
    var TAB = "	".charCodeAt(0);
    var CR = "\r".charCodeAt(0);
    var OPEN_SQUARE = "[".charCodeAt(0);
    var CLOSE_SQUARE = "]".charCodeAt(0);
    var OPEN_PARENTHESES = "(".charCodeAt(0);
    var CLOSE_PARENTHESES = ")".charCodeAt(0);
    var OPEN_CURLY = "{".charCodeAt(0);
    var CLOSE_CURLY = "}".charCodeAt(0);
    var SEMICOLON = ";".charCodeAt(0);
    var ASTERISK = "*".charCodeAt(0);
    var COLON = ":".charCodeAt(0);
    var AT = "@".charCodeAt(0);
    var RE_AT_END = /[\t\n\f\r "#'()/;[\\\]{}]/g;
    var RE_WORD_END = /[\t\n\f\r !"#'():;@[\\\]{}]|\/(?=\*)/g;
    var RE_BAD_BRACKET = /.[\n"'(/\\]/;
    var RE_HEX_ESCAPE = /[\da-f]/i;
    module2.exports = function tokenizer(input, options = {}) {
      let css = input.css.valueOf();
      let ignore = options.ignoreErrors;
      let code, next, quote, content, escape;
      let escaped, escapePos, prev, n, currentToken;
      let length = css.length;
      let pos = 0;
      let buffer = [];
      let returned = [];
      function position() {
        return pos;
      }
      function unclosed(what) {
        throw input.error("Unclosed " + what, pos);
      }
      function endOfFile() {
        return returned.length === 0 && pos >= length;
      }
      function nextToken(opts) {
        if (returned.length)
          return returned.pop();
        if (pos >= length)
          return;
        let ignoreUnclosed = opts ? opts.ignoreUnclosed : false;
        code = css.charCodeAt(pos);
        switch (code) {
          case NEWLINE:
          case SPACE:
          case TAB:
          case CR:
          case FEED: {
            next = pos;
            do {
              next += 1;
              code = css.charCodeAt(next);
            } while (code === SPACE || code === NEWLINE || code === TAB || code === CR || code === FEED);
            currentToken = ["space", css.slice(pos, next)];
            pos = next - 1;
            break;
          }
          case OPEN_SQUARE:
          case CLOSE_SQUARE:
          case OPEN_CURLY:
          case CLOSE_CURLY:
          case COLON:
          case SEMICOLON:
          case CLOSE_PARENTHESES: {
            let controlChar = String.fromCharCode(code);
            currentToken = [controlChar, controlChar, pos];
            break;
          }
          case OPEN_PARENTHESES: {
            prev = buffer.length ? buffer.pop()[1] : "";
            n = css.charCodeAt(pos + 1);
            if (prev === "url" && n !== SINGLE_QUOTE && n !== DOUBLE_QUOTE && n !== SPACE && n !== NEWLINE && n !== TAB && n !== FEED && n !== CR) {
              next = pos;
              do {
                escaped = false;
                next = css.indexOf(")", next + 1);
                if (next === -1) {
                  if (ignore || ignoreUnclosed) {
                    next = pos;
                    break;
                  } else {
                    unclosed("bracket");
                  }
                }
                escapePos = next;
                while (css.charCodeAt(escapePos - 1) === BACKSLASH) {
                  escapePos -= 1;
                  escaped = !escaped;
                }
              } while (escaped);
              currentToken = ["brackets", css.slice(pos, next + 1), pos, next];
              pos = next;
            } else {
              next = css.indexOf(")", pos + 1);
              content = css.slice(pos, next + 1);
              if (next === -1 || RE_BAD_BRACKET.test(content)) {
                currentToken = ["(", "(", pos];
              } else {
                currentToken = ["brackets", content, pos, next];
                pos = next;
              }
            }
            break;
          }
          case SINGLE_QUOTE:
          case DOUBLE_QUOTE: {
            quote = code === SINGLE_QUOTE ? "'" : '"';
            next = pos;
            do {
              escaped = false;
              next = css.indexOf(quote, next + 1);
              if (next === -1) {
                if (ignore || ignoreUnclosed) {
                  next = pos + 1;
                  break;
                } else {
                  unclosed("string");
                }
              }
              escapePos = next;
              while (css.charCodeAt(escapePos - 1) === BACKSLASH) {
                escapePos -= 1;
                escaped = !escaped;
              }
            } while (escaped);
            currentToken = ["string", css.slice(pos, next + 1), pos, next];
            pos = next;
            break;
          }
          case AT: {
            RE_AT_END.lastIndex = pos + 1;
            RE_AT_END.test(css);
            if (RE_AT_END.lastIndex === 0) {
              next = css.length - 1;
            } else {
              next = RE_AT_END.lastIndex - 2;
            }
            currentToken = ["at-word", css.slice(pos, next + 1), pos, next];
            pos = next;
            break;
          }
          case BACKSLASH: {
            next = pos;
            escape = true;
            while (css.charCodeAt(next + 1) === BACKSLASH) {
              next += 1;
              escape = !escape;
            }
            code = css.charCodeAt(next + 1);
            if (escape && code !== SLASH && code !== SPACE && code !== NEWLINE && code !== TAB && code !== CR && code !== FEED) {
              next += 1;
              if (RE_HEX_ESCAPE.test(css.charAt(next))) {
                while (RE_HEX_ESCAPE.test(css.charAt(next + 1))) {
                  next += 1;
                }
                if (css.charCodeAt(next + 1) === SPACE) {
                  next += 1;
                }
              }
            }
            currentToken = ["word", css.slice(pos, next + 1), pos, next];
            pos = next;
            break;
          }
          default: {
            if (code === SLASH && css.charCodeAt(pos + 1) === ASTERISK) {
              next = css.indexOf("*/", pos + 2) + 1;
              if (next === 0) {
                if (ignore || ignoreUnclosed) {
                  next = css.length;
                } else {
                  unclosed("comment");
                }
              }
              currentToken = ["comment", css.slice(pos, next + 1), pos, next];
              pos = next;
            } else {
              RE_WORD_END.lastIndex = pos + 1;
              RE_WORD_END.test(css);
              if (RE_WORD_END.lastIndex === 0) {
                next = css.length - 1;
              } else {
                next = RE_WORD_END.lastIndex - 2;
              }
              currentToken = ["word", css.slice(pos, next + 1), pos, next];
              buffer.push(currentToken);
              pos = next;
            }
            break;
          }
        }
        pos++;
        return currentToken;
      }
      function back(token) {
        returned.push(token);
      }
      return {
        back,
        nextToken,
        endOfFile,
        position
      };
    };
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/terminal-highlight.js
var require_terminal_highlight = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/terminal-highlight.js"(exports, module2) {
    "use strict";
    var pico = require_picocolors();
    var tokenizer = require_tokenize();
    var Input;
    function registerInput(dependant) {
      Input = dependant;
    }
    var HIGHLIGHT_THEME = {
      "brackets": pico.cyan,
      "at-word": pico.cyan,
      "comment": pico.gray,
      "string": pico.green,
      "class": pico.yellow,
      "hash": pico.magenta,
      "call": pico.cyan,
      "(": pico.cyan,
      ")": pico.cyan,
      "{": pico.yellow,
      "}": pico.yellow,
      "[": pico.yellow,
      "]": pico.yellow,
      ":": pico.yellow,
      ";": pico.yellow
    };
    function getTokenType([type, value], processor) {
      if (type === "word") {
        if (value[0] === ".") {
          return "class";
        }
        if (value[0] === "#") {
          return "hash";
        }
      }
      if (!processor.endOfFile()) {
        let next = processor.nextToken();
        processor.back(next);
        if (next[0] === "brackets" || next[0] === "(")
          return "call";
      }
      return type;
    }
    function terminalHighlight(css) {
      let processor = tokenizer(new Input(css), { ignoreErrors: true });
      let result = "";
      while (!processor.endOfFile()) {
        let token = processor.nextToken();
        let color = HIGHLIGHT_THEME[getTokenType(token, processor)];
        if (color) {
          result += token[1].split(/\r?\n/).map((i) => color(i)).join("\n");
        } else {
          result += token[1];
        }
      }
      return result;
    }
    terminalHighlight.registerInput = registerInput;
    module2.exports = terminalHighlight;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/css-syntax-error.js
var require_css_syntax_error = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/css-syntax-error.js"(exports, module2) {
    "use strict";
    var pico = require_picocolors();
    var terminalHighlight = require_terminal_highlight();
    var CssSyntaxError = class extends Error {
      constructor(message, line, column, source, file, plugin3) {
        super(message);
        this.name = "CssSyntaxError";
        this.reason = message;
        if (file) {
          this.file = file;
        }
        if (source) {
          this.source = source;
        }
        if (plugin3) {
          this.plugin = plugin3;
        }
        if (typeof line !== "undefined" && typeof column !== "undefined") {
          if (typeof line === "number") {
            this.line = line;
            this.column = column;
          } else {
            this.line = line.line;
            this.column = line.column;
            this.endLine = column.line;
            this.endColumn = column.column;
          }
        }
        this.setMessage();
        if (Error.captureStackTrace) {
          Error.captureStackTrace(this, CssSyntaxError);
        }
      }
      setMessage() {
        this.message = this.plugin ? this.plugin + ": " : "";
        this.message += this.file ? this.file : "<css input>";
        if (typeof this.line !== "undefined") {
          this.message += ":" + this.line + ":" + this.column;
        }
        this.message += ": " + this.reason;
      }
      showSourceCode(color) {
        if (!this.source)
          return "";
        let css = this.source;
        if (color == null)
          color = pico.isColorSupported;
        if (terminalHighlight) {
          if (color)
            css = terminalHighlight(css);
        }
        let lines = css.split(/\r?\n/);
        let start = Math.max(this.line - 3, 0);
        let end = Math.min(this.line + 2, lines.length);
        let maxWidth = String(end).length;
        let mark, aside;
        if (color) {
          let { bold, red, gray } = pico.createColors(true);
          mark = (text2) => bold(red(text2));
          aside = (text2) => gray(text2);
        } else {
          mark = aside = (str) => str;
        }
        return lines.slice(start, end).map((line, index2) => {
          let number = start + 1 + index2;
          let gutter = " " + (" " + number).slice(-maxWidth) + " | ";
          if (number === this.line) {
            let spacing = aside(gutter.replace(/\d/g, " ")) + line.slice(0, this.column - 1).replace(/[^\t]/g, " ");
            return mark(">") + aside(gutter) + line + "\n " + spacing + mark("^");
          }
          return " " + aside(gutter) + line;
        }).join("\n");
      }
      toString() {
        let code = this.showSourceCode();
        if (code) {
          code = "\n\n" + code + "\n";
        }
        return this.name + ": " + this.message + code;
      }
    };
    module2.exports = CssSyntaxError;
    CssSyntaxError.default = CssSyntaxError;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/symbols.js
var require_symbols = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/symbols.js"(exports, module2) {
    "use strict";
    module2.exports.isClean = Symbol("isClean");
    module2.exports.my = Symbol("my");
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/stringifier.js
var require_stringifier = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/stringifier.js"(exports, module2) {
    "use strict";
    var DEFAULT_RAW = {
      colon: ": ",
      indent: "    ",
      beforeDecl: "\n",
      beforeRule: "\n",
      beforeOpen: " ",
      beforeClose: "\n",
      beforeComment: "\n",
      after: "\n",
      emptyBody: "",
      commentLeft: " ",
      commentRight: " ",
      semicolon: false
    };
    function capitalize(str) {
      return str[0].toUpperCase() + str.slice(1);
    }
    var Stringifier = class {
      constructor(builder) {
        this.builder = builder;
      }
      stringify(node, semicolon) {
        if (!this[node.type]) {
          throw new Error(
            "Unknown AST node type " + node.type + ". Maybe you need to change PostCSS stringifier."
          );
        }
        this[node.type](node, semicolon);
      }
      document(node) {
        this.body(node);
      }
      root(node) {
        this.body(node);
        if (node.raws.after)
          this.builder(node.raws.after);
      }
      comment(node) {
        let left = this.raw(node, "left", "commentLeft");
        let right = this.raw(node, "right", "commentRight");
        this.builder("/*" + left + node.text + right + "*/", node);
      }
      decl(node, semicolon) {
        let between = this.raw(node, "between", "colon");
        let string = node.prop + between + this.rawValue(node, "value");
        if (node.important) {
          string += node.raws.important || " !important";
        }
        if (semicolon)
          string += ";";
        this.builder(string, node);
      }
      rule(node) {
        this.block(node, this.rawValue(node, "selector"));
        if (node.raws.ownSemicolon) {
          this.builder(node.raws.ownSemicolon, node, "end");
        }
      }
      atrule(node, semicolon) {
        let name = "@" + node.name;
        let params = node.params ? this.rawValue(node, "params") : "";
        if (typeof node.raws.afterName !== "undefined") {
          name += node.raws.afterName;
        } else if (params) {
          name += " ";
        }
        if (node.nodes) {
          this.block(node, name + params);
        } else {
          let end = (node.raws.between || "") + (semicolon ? ";" : "");
          this.builder(name + params + end, node);
        }
      }
      body(node) {
        let last = node.nodes.length - 1;
        while (last > 0) {
          if (node.nodes[last].type !== "comment")
            break;
          last -= 1;
        }
        let semicolon = this.raw(node, "semicolon");
        for (let i = 0; i < node.nodes.length; i++) {
          let child = node.nodes[i];
          let before = this.raw(child, "before");
          if (before)
            this.builder(before);
          this.stringify(child, last !== i || semicolon);
        }
      }
      block(node, start) {
        let between = this.raw(node, "between", "beforeOpen");
        this.builder(start + between + "{", node, "start");
        let after;
        if (node.nodes && node.nodes.length) {
          this.body(node);
          after = this.raw(node, "after");
        } else {
          after = this.raw(node, "after", "emptyBody");
        }
        if (after)
          this.builder(after);
        this.builder("}", node, "end");
      }
      raw(node, own, detect) {
        let value;
        if (!detect)
          detect = own;
        if (own) {
          value = node.raws[own];
          if (typeof value !== "undefined")
            return value;
        }
        let parent = node.parent;
        if (detect === "before") {
          if (!parent || parent.type === "root" && parent.first === node) {
            return "";
          }
          if (parent && parent.type === "document") {
            return "";
          }
        }
        if (!parent)
          return DEFAULT_RAW[detect];
        let root = node.root();
        if (!root.rawCache)
          root.rawCache = {};
        if (typeof root.rawCache[detect] !== "undefined") {
          return root.rawCache[detect];
        }
        if (detect === "before" || detect === "after") {
          return this.beforeAfter(node, detect);
        } else {
          let method = "raw" + capitalize(detect);
          if (this[method]) {
            value = this[method](root, node);
          } else {
            root.walk((i) => {
              value = i.raws[own];
              if (typeof value !== "undefined")
                return false;
            });
          }
        }
        if (typeof value === "undefined")
          value = DEFAULT_RAW[detect];
        root.rawCache[detect] = value;
        return value;
      }
      rawSemicolon(root) {
        let value;
        root.walk((i) => {
          if (i.nodes && i.nodes.length && i.last.type === "decl") {
            value = i.raws.semicolon;
            if (typeof value !== "undefined")
              return false;
          }
        });
        return value;
      }
      rawEmptyBody(root) {
        let value;
        root.walk((i) => {
          if (i.nodes && i.nodes.length === 0) {
            value = i.raws.after;
            if (typeof value !== "undefined")
              return false;
          }
        });
        return value;
      }
      rawIndent(root) {
        if (root.raws.indent)
          return root.raws.indent;
        let value;
        root.walk((i) => {
          let p = i.parent;
          if (p && p !== root && p.parent && p.parent === root) {
            if (typeof i.raws.before !== "undefined") {
              let parts = i.raws.before.split("\n");
              value = parts[parts.length - 1];
              value = value.replace(/\S/g, "");
              return false;
            }
          }
        });
        return value;
      }
      rawBeforeComment(root, node) {
        let value;
        root.walkComments((i) => {
          if (typeof i.raws.before !== "undefined") {
            value = i.raws.before;
            if (value.includes("\n")) {
              value = value.replace(/[^\n]+$/, "");
            }
            return false;
          }
        });
        if (typeof value === "undefined") {
          value = this.raw(node, null, "beforeDecl");
        } else if (value) {
          value = value.replace(/\S/g, "");
        }
        return value;
      }
      rawBeforeDecl(root, node) {
        let value;
        root.walkDecls((i) => {
          if (typeof i.raws.before !== "undefined") {
            value = i.raws.before;
            if (value.includes("\n")) {
              value = value.replace(/[^\n]+$/, "");
            }
            return false;
          }
        });
        if (typeof value === "undefined") {
          value = this.raw(node, null, "beforeRule");
        } else if (value) {
          value = value.replace(/\S/g, "");
        }
        return value;
      }
      rawBeforeRule(root) {
        let value;
        root.walk((i) => {
          if (i.nodes && (i.parent !== root || root.first !== i)) {
            if (typeof i.raws.before !== "undefined") {
              value = i.raws.before;
              if (value.includes("\n")) {
                value = value.replace(/[^\n]+$/, "");
              }
              return false;
            }
          }
        });
        if (value)
          value = value.replace(/\S/g, "");
        return value;
      }
      rawBeforeClose(root) {
        let value;
        root.walk((i) => {
          if (i.nodes && i.nodes.length > 0) {
            if (typeof i.raws.after !== "undefined") {
              value = i.raws.after;
              if (value.includes("\n")) {
                value = value.replace(/[^\n]+$/, "");
              }
              return false;
            }
          }
        });
        if (value)
          value = value.replace(/\S/g, "");
        return value;
      }
      rawBeforeOpen(root) {
        let value;
        root.walk((i) => {
          if (i.type !== "decl") {
            value = i.raws.between;
            if (typeof value !== "undefined")
              return false;
          }
        });
        return value;
      }
      rawColon(root) {
        let value;
        root.walkDecls((i) => {
          if (typeof i.raws.between !== "undefined") {
            value = i.raws.between.replace(/[^\s:]/g, "");
            return false;
          }
        });
        return value;
      }
      beforeAfter(node, detect) {
        let value;
        if (node.type === "decl") {
          value = this.raw(node, null, "beforeDecl");
        } else if (node.type === "comment") {
          value = this.raw(node, null, "beforeComment");
        } else if (detect === "before") {
          value = this.raw(node, null, "beforeRule");
        } else {
          value = this.raw(node, null, "beforeClose");
        }
        let buf = node.parent;
        let depth = 0;
        while (buf && buf.type !== "root") {
          depth += 1;
          buf = buf.parent;
        }
        if (value.includes("\n")) {
          let indent = this.raw(node, null, "indent");
          if (indent.length) {
            for (let step = 0; step < depth; step++)
              value += indent;
          }
        }
        return value;
      }
      rawValue(node, prop) {
        let value = node[prop];
        let raw = node.raws[prop];
        if (raw && raw.value === value) {
          return raw.raw;
        }
        return value;
      }
    };
    module2.exports = Stringifier;
    Stringifier.default = Stringifier;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/stringify.js
var require_stringify = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/stringify.js"(exports, module2) {
    "use strict";
    var Stringifier = require_stringifier();
    function stringify(node, builder) {
      let str = new Stringifier(builder);
      str.stringify(node);
    }
    module2.exports = stringify;
    stringify.default = stringify;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/node.js
var require_node = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/node.js"(exports, module2) {
    "use strict";
    var { isClean, my } = require_symbols();
    var CssSyntaxError = require_css_syntax_error();
    var Stringifier = require_stringifier();
    var stringify = require_stringify();
    function cloneNode(obj, parent) {
      let cloned = new obj.constructor();
      for (let i in obj) {
        if (!Object.prototype.hasOwnProperty.call(obj, i)) {
          continue;
        }
        if (i === "proxyCache")
          continue;
        let value = obj[i];
        let type = typeof value;
        if (i === "parent" && type === "object") {
          if (parent)
            cloned[i] = parent;
        } else if (i === "source") {
          cloned[i] = value;
        } else if (Array.isArray(value)) {
          cloned[i] = value.map((j) => cloneNode(j, cloned));
        } else {
          if (type === "object" && value !== null)
            value = cloneNode(value);
          cloned[i] = value;
        }
      }
      return cloned;
    }
    var Node = class {
      constructor(defaults = {}) {
        this.raws = {};
        this[isClean] = false;
        this[my] = true;
        for (let name in defaults) {
          if (name === "nodes") {
            this.nodes = [];
            for (let node of defaults[name]) {
              if (typeof node.clone === "function") {
                this.append(node.clone());
              } else {
                this.append(node);
              }
            }
          } else {
            this[name] = defaults[name];
          }
        }
      }
      error(message, opts = {}) {
        if (this.source) {
          let { start, end } = this.rangeBy(opts);
          return this.source.input.error(
            message,
            { line: start.line, column: start.column },
            { line: end.line, column: end.column },
            opts
          );
        }
        return new CssSyntaxError(message);
      }
      warn(result, text2, opts) {
        let data = { node: this };
        for (let i in opts)
          data[i] = opts[i];
        return result.warn(text2, data);
      }
      remove() {
        if (this.parent) {
          this.parent.removeChild(this);
        }
        this.parent = void 0;
        return this;
      }
      toString(stringifier = stringify) {
        if (stringifier.stringify)
          stringifier = stringifier.stringify;
        let result = "";
        stringifier(this, (i) => {
          result += i;
        });
        return result;
      }
      assign(overrides = {}) {
        for (let name in overrides) {
          this[name] = overrides[name];
        }
        return this;
      }
      clone(overrides = {}) {
        let cloned = cloneNode(this);
        for (let name in overrides) {
          cloned[name] = overrides[name];
        }
        return cloned;
      }
      cloneBefore(overrides = {}) {
        let cloned = this.clone(overrides);
        this.parent.insertBefore(this, cloned);
        return cloned;
      }
      cloneAfter(overrides = {}) {
        let cloned = this.clone(overrides);
        this.parent.insertAfter(this, cloned);
        return cloned;
      }
      replaceWith(...nodes) {
        if (this.parent) {
          let bookmark = this;
          let foundSelf = false;
          for (let node of nodes) {
            if (node === this) {
              foundSelf = true;
            } else if (foundSelf) {
              this.parent.insertAfter(bookmark, node);
              bookmark = node;
            } else {
              this.parent.insertBefore(bookmark, node);
            }
          }
          if (!foundSelf) {
            this.remove();
          }
        }
        return this;
      }
      next() {
        if (!this.parent)
          return void 0;
        let index2 = this.parent.index(this);
        return this.parent.nodes[index2 + 1];
      }
      prev() {
        if (!this.parent)
          return void 0;
        let index2 = this.parent.index(this);
        return this.parent.nodes[index2 - 1];
      }
      before(add) {
        this.parent.insertBefore(this, add);
        return this;
      }
      after(add) {
        this.parent.insertAfter(this, add);
        return this;
      }
      root() {
        let result = this;
        while (result.parent && result.parent.type !== "document") {
          result = result.parent;
        }
        return result;
      }
      raw(prop, defaultType) {
        let str = new Stringifier();
        return str.raw(this, prop, defaultType);
      }
      cleanRaws(keepBetween) {
        delete this.raws.before;
        delete this.raws.after;
        if (!keepBetween)
          delete this.raws.between;
      }
      toJSON(_, inputs) {
        let fixed = {};
        let emitInputs = inputs == null;
        inputs = inputs || /* @__PURE__ */ new Map();
        let inputsNextIndex = 0;
        for (let name in this) {
          if (!Object.prototype.hasOwnProperty.call(this, name)) {
            continue;
          }
          if (name === "parent" || name === "proxyCache")
            continue;
          let value = this[name];
          if (Array.isArray(value)) {
            fixed[name] = value.map((i) => {
              if (typeof i === "object" && i.toJSON) {
                return i.toJSON(null, inputs);
              } else {
                return i;
              }
            });
          } else if (typeof value === "object" && value.toJSON) {
            fixed[name] = value.toJSON(null, inputs);
          } else if (name === "source") {
            let inputId = inputs.get(value.input);
            if (inputId == null) {
              inputId = inputsNextIndex;
              inputs.set(value.input, inputsNextIndex);
              inputsNextIndex++;
            }
            fixed[name] = {
              inputId,
              start: value.start,
              end: value.end
            };
          } else {
            fixed[name] = value;
          }
        }
        if (emitInputs) {
          fixed.inputs = [...inputs.keys()].map((input) => input.toJSON());
        }
        return fixed;
      }
      positionInside(index2) {
        let string = this.toString();
        let column = this.source.start.column;
        let line = this.source.start.line;
        for (let i = 0; i < index2; i++) {
          if (string[i] === "\n") {
            column = 1;
            line += 1;
          } else {
            column += 1;
          }
        }
        return { line, column };
      }
      positionBy(opts) {
        let pos = this.source.start;
        if (opts.index) {
          pos = this.positionInside(opts.index);
        } else if (opts.word) {
          let index2 = this.toString().indexOf(opts.word);
          if (index2 !== -1)
            pos = this.positionInside(index2);
        }
        return pos;
      }
      rangeBy(opts) {
        let start = {
          line: this.source.start.line,
          column: this.source.start.column
        };
        let end = this.source.end ? {
          line: this.source.end.line,
          column: this.source.end.column + 1
        } : {
          line: start.line,
          column: start.column + 1
        };
        if (opts.word) {
          let index2 = this.toString().indexOf(opts.word);
          if (index2 !== -1) {
            start = this.positionInside(index2);
            end = this.positionInside(index2 + opts.word.length);
          }
        } else {
          if (opts.start) {
            start = {
              line: opts.start.line,
              column: opts.start.column
            };
          } else if (opts.index) {
            start = this.positionInside(opts.index);
          }
          if (opts.end) {
            end = {
              line: opts.end.line,
              column: opts.end.column
            };
          } else if (opts.endIndex) {
            end = this.positionInside(opts.endIndex);
          } else if (opts.index) {
            end = this.positionInside(opts.index + 1);
          }
        }
        if (end.line < start.line || end.line === start.line && end.column <= start.column) {
          end = { line: start.line, column: start.column + 1 };
        }
        return { start, end };
      }
      getProxyProcessor() {
        return {
          set(node, prop, value) {
            if (node[prop] === value)
              return true;
            node[prop] = value;
            if (prop === "prop" || prop === "value" || prop === "name" || prop === "params" || prop === "important" || /* c8 ignore next */
            prop === "text") {
              node.markDirty();
            }
            return true;
          },
          get(node, prop) {
            if (prop === "proxyOf") {
              return node;
            } else if (prop === "root") {
              return () => node.root().toProxy();
            } else {
              return node[prop];
            }
          }
        };
      }
      toProxy() {
        if (!this.proxyCache) {
          this.proxyCache = new Proxy(this, this.getProxyProcessor());
        }
        return this.proxyCache;
      }
      addToError(error) {
        error.postcssNode = this;
        if (error.stack && this.source && /\n\s{4}at /.test(error.stack)) {
          let s = this.source;
          error.stack = error.stack.replace(
            /\n\s{4}at /,
            `$&${s.input.from}:${s.start.line}:${s.start.column}$&`
          );
        }
        return error;
      }
      markDirty() {
        if (this[isClean]) {
          this[isClean] = false;
          let next = this;
          while (next = next.parent) {
            next[isClean] = false;
          }
        }
      }
      get proxyOf() {
        return this;
      }
    };
    module2.exports = Node;
    Node.default = Node;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/declaration.js
var require_declaration = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/declaration.js"(exports, module2) {
    "use strict";
    var Node = require_node();
    var Declaration = class extends Node {
      constructor(defaults) {
        if (defaults && typeof defaults.value !== "undefined" && typeof defaults.value !== "string") {
          defaults = { ...defaults, value: String(defaults.value) };
        }
        super(defaults);
        this.type = "decl";
      }
      get variable() {
        return this.prop.startsWith("--") || this.prop[0] === "$";
      }
    };
    module2.exports = Declaration;
    Declaration.default = Declaration;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/base64.js
var require_base64 = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/base64.js"(exports) {
    var intToCharMap = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".split("");
    exports.encode = function(number) {
      if (0 <= number && number < intToCharMap.length) {
        return intToCharMap[number];
      }
      throw new TypeError("Must be between 0 and 63: " + number);
    };
    exports.decode = function(charCode) {
      var bigA = 65;
      var bigZ = 90;
      var littleA = 97;
      var littleZ = 122;
      var zero = 48;
      var nine = 57;
      var plus = 43;
      var slash = 47;
      var littleOffset = 26;
      var numberOffset = 52;
      if (bigA <= charCode && charCode <= bigZ) {
        return charCode - bigA;
      }
      if (littleA <= charCode && charCode <= littleZ) {
        return charCode - littleA + littleOffset;
      }
      if (zero <= charCode && charCode <= nine) {
        return charCode - zero + numberOffset;
      }
      if (charCode == plus) {
        return 62;
      }
      if (charCode == slash) {
        return 63;
      }
      return -1;
    };
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/base64-vlq.js
var require_base64_vlq = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/base64-vlq.js"(exports) {
    var base64 = require_base64();
    var VLQ_BASE_SHIFT = 5;
    var VLQ_BASE = 1 << VLQ_BASE_SHIFT;
    var VLQ_BASE_MASK = VLQ_BASE - 1;
    var VLQ_CONTINUATION_BIT = VLQ_BASE;
    function toVLQSigned(aValue) {
      return aValue < 0 ? (-aValue << 1) + 1 : (aValue << 1) + 0;
    }
    function fromVLQSigned(aValue) {
      var isNegative = (aValue & 1) === 1;
      var shifted = aValue >> 1;
      return isNegative ? -shifted : shifted;
    }
    exports.encode = function base64VLQ_encode(aValue) {
      var encoded = "";
      var digit;
      var vlq = toVLQSigned(aValue);
      do {
        digit = vlq & VLQ_BASE_MASK;
        vlq >>>= VLQ_BASE_SHIFT;
        if (vlq > 0) {
          digit |= VLQ_CONTINUATION_BIT;
        }
        encoded += base64.encode(digit);
      } while (vlq > 0);
      return encoded;
    };
    exports.decode = function base64VLQ_decode(aStr, aIndex, aOutParam) {
      var strLen = aStr.length;
      var result = 0;
      var shift = 0;
      var continuation, digit;
      do {
        if (aIndex >= strLen) {
          throw new Error("Expected more digits in base 64 VLQ value.");
        }
        digit = base64.decode(aStr.charCodeAt(aIndex++));
        if (digit === -1) {
          throw new Error("Invalid base64 digit: " + aStr.charAt(aIndex - 1));
        }
        continuation = !!(digit & VLQ_CONTINUATION_BIT);
        digit &= VLQ_BASE_MASK;
        result = result + (digit << shift);
        shift += VLQ_BASE_SHIFT;
      } while (continuation);
      aOutParam.value = fromVLQSigned(result);
      aOutParam.rest = aIndex;
    };
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/util.js
var require_util = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/util.js"(exports) {
    function getArg(aArgs, aName, aDefaultValue) {
      if (aName in aArgs) {
        return aArgs[aName];
      } else if (arguments.length === 3) {
        return aDefaultValue;
      } else {
        throw new Error('"' + aName + '" is a required argument.');
      }
    }
    exports.getArg = getArg;
    var urlRegexp = /^(?:([\w+\-.]+):)?\/\/(?:(\w+:\w+)@)?([\w.-]*)(?::(\d+))?(.*)$/;
    var dataUrlRegexp = /^data:.+\,.+$/;
    function urlParse(aUrl) {
      var match = aUrl.match(urlRegexp);
      if (!match) {
        return null;
      }
      return {
        scheme: match[1],
        auth: match[2],
        host: match[3],
        port: match[4],
        path: match[5]
      };
    }
    exports.urlParse = urlParse;
    function urlGenerate(aParsedUrl) {
      var url = "";
      if (aParsedUrl.scheme) {
        url += aParsedUrl.scheme + ":";
      }
      url += "//";
      if (aParsedUrl.auth) {
        url += aParsedUrl.auth + "@";
      }
      if (aParsedUrl.host) {
        url += aParsedUrl.host;
      }
      if (aParsedUrl.port) {
        url += ":" + aParsedUrl.port;
      }
      if (aParsedUrl.path) {
        url += aParsedUrl.path;
      }
      return url;
    }
    exports.urlGenerate = urlGenerate;
    var MAX_CACHED_INPUTS = 32;
    function lruMemoize(f) {
      var cache = [];
      return function(input) {
        for (var i = 0; i < cache.length; i++) {
          if (cache[i].input === input) {
            var temp = cache[0];
            cache[0] = cache[i];
            cache[i] = temp;
            return cache[0].result;
          }
        }
        var result = f(input);
        cache.unshift({
          input,
          result
        });
        if (cache.length > MAX_CACHED_INPUTS) {
          cache.pop();
        }
        return result;
      };
    }
    var normalize = lruMemoize(function normalize2(aPath) {
      var path = aPath;
      var url = urlParse(aPath);
      if (url) {
        if (!url.path) {
          return aPath;
        }
        path = url.path;
      }
      var isAbsolute = exports.isAbsolute(path);
      var parts = [];
      var start = 0;
      var i = 0;
      while (true) {
        start = i;
        i = path.indexOf("/", start);
        if (i === -1) {
          parts.push(path.slice(start));
          break;
        } else {
          parts.push(path.slice(start, i));
          while (i < path.length && path[i] === "/") {
            i++;
          }
        }
      }
      for (var part, up = 0, i = parts.length - 1; i >= 0; i--) {
        part = parts[i];
        if (part === ".") {
          parts.splice(i, 1);
        } else if (part === "..") {
          up++;
        } else if (up > 0) {
          if (part === "") {
            parts.splice(i + 1, up);
            up = 0;
          } else {
            parts.splice(i, 2);
            up--;
          }
        }
      }
      path = parts.join("/");
      if (path === "") {
        path = isAbsolute ? "/" : ".";
      }
      if (url) {
        url.path = path;
        return urlGenerate(url);
      }
      return path;
    });
    exports.normalize = normalize;
    function join(aRoot, aPath) {
      if (aRoot === "") {
        aRoot = ".";
      }
      if (aPath === "") {
        aPath = ".";
      }
      var aPathUrl = urlParse(aPath);
      var aRootUrl = urlParse(aRoot);
      if (aRootUrl) {
        aRoot = aRootUrl.path || "/";
      }
      if (aPathUrl && !aPathUrl.scheme) {
        if (aRootUrl) {
          aPathUrl.scheme = aRootUrl.scheme;
        }
        return urlGenerate(aPathUrl);
      }
      if (aPathUrl || aPath.match(dataUrlRegexp)) {
        return aPath;
      }
      if (aRootUrl && !aRootUrl.host && !aRootUrl.path) {
        aRootUrl.host = aPath;
        return urlGenerate(aRootUrl);
      }
      var joined = aPath.charAt(0) === "/" ? aPath : normalize(aRoot.replace(/\/+$/, "") + "/" + aPath);
      if (aRootUrl) {
        aRootUrl.path = joined;
        return urlGenerate(aRootUrl);
      }
      return joined;
    }
    exports.join = join;
    exports.isAbsolute = function(aPath) {
      return aPath.charAt(0) === "/" || urlRegexp.test(aPath);
    };
    function relative(aRoot, aPath) {
      if (aRoot === "") {
        aRoot = ".";
      }
      aRoot = aRoot.replace(/\/$/, "");
      var level = 0;
      while (aPath.indexOf(aRoot + "/") !== 0) {
        var index2 = aRoot.lastIndexOf("/");
        if (index2 < 0) {
          return aPath;
        }
        aRoot = aRoot.slice(0, index2);
        if (aRoot.match(/^([^\/]+:\/)?\/*$/)) {
          return aPath;
        }
        ++level;
      }
      return Array(level + 1).join("../") + aPath.substr(aRoot.length + 1);
    }
    exports.relative = relative;
    var supportsNullProto = function() {
      var obj = /* @__PURE__ */ Object.create(null);
      return !("__proto__" in obj);
    }();
    function identity(s) {
      return s;
    }
    function toSetString(aStr) {
      if (isProtoString(aStr)) {
        return "$" + aStr;
      }
      return aStr;
    }
    exports.toSetString = supportsNullProto ? identity : toSetString;
    function fromSetString(aStr) {
      if (isProtoString(aStr)) {
        return aStr.slice(1);
      }
      return aStr;
    }
    exports.fromSetString = supportsNullProto ? identity : fromSetString;
    function isProtoString(s) {
      if (!s) {
        return false;
      }
      var length = s.length;
      if (length < 9) {
        return false;
      }
      if (s.charCodeAt(length - 1) !== 95 || s.charCodeAt(length - 2) !== 95 || s.charCodeAt(length - 3) !== 111 || s.charCodeAt(length - 4) !== 116 || s.charCodeAt(length - 5) !== 111 || s.charCodeAt(length - 6) !== 114 || s.charCodeAt(length - 7) !== 112 || s.charCodeAt(length - 8) !== 95 || s.charCodeAt(length - 9) !== 95) {
        return false;
      }
      for (var i = length - 10; i >= 0; i--) {
        if (s.charCodeAt(i) !== 36) {
          return false;
        }
      }
      return true;
    }
    function compareByOriginalPositions(mappingA, mappingB, onlyCompareOriginal) {
      var cmp = strcmp(mappingA.source, mappingB.source);
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalLine - mappingB.originalLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalColumn - mappingB.originalColumn;
      if (cmp !== 0 || onlyCompareOriginal) {
        return cmp;
      }
      cmp = mappingA.generatedColumn - mappingB.generatedColumn;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.generatedLine - mappingB.generatedLine;
      if (cmp !== 0) {
        return cmp;
      }
      return strcmp(mappingA.name, mappingB.name);
    }
    exports.compareByOriginalPositions = compareByOriginalPositions;
    function compareByOriginalPositionsNoSource(mappingA, mappingB, onlyCompareOriginal) {
      var cmp;
      cmp = mappingA.originalLine - mappingB.originalLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalColumn - mappingB.originalColumn;
      if (cmp !== 0 || onlyCompareOriginal) {
        return cmp;
      }
      cmp = mappingA.generatedColumn - mappingB.generatedColumn;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.generatedLine - mappingB.generatedLine;
      if (cmp !== 0) {
        return cmp;
      }
      return strcmp(mappingA.name, mappingB.name);
    }
    exports.compareByOriginalPositionsNoSource = compareByOriginalPositionsNoSource;
    function compareByGeneratedPositionsDeflated(mappingA, mappingB, onlyCompareGenerated) {
      var cmp = mappingA.generatedLine - mappingB.generatedLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.generatedColumn - mappingB.generatedColumn;
      if (cmp !== 0 || onlyCompareGenerated) {
        return cmp;
      }
      cmp = strcmp(mappingA.source, mappingB.source);
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalLine - mappingB.originalLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalColumn - mappingB.originalColumn;
      if (cmp !== 0) {
        return cmp;
      }
      return strcmp(mappingA.name, mappingB.name);
    }
    exports.compareByGeneratedPositionsDeflated = compareByGeneratedPositionsDeflated;
    function compareByGeneratedPositionsDeflatedNoLine(mappingA, mappingB, onlyCompareGenerated) {
      var cmp = mappingA.generatedColumn - mappingB.generatedColumn;
      if (cmp !== 0 || onlyCompareGenerated) {
        return cmp;
      }
      cmp = strcmp(mappingA.source, mappingB.source);
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalLine - mappingB.originalLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalColumn - mappingB.originalColumn;
      if (cmp !== 0) {
        return cmp;
      }
      return strcmp(mappingA.name, mappingB.name);
    }
    exports.compareByGeneratedPositionsDeflatedNoLine = compareByGeneratedPositionsDeflatedNoLine;
    function strcmp(aStr1, aStr2) {
      if (aStr1 === aStr2) {
        return 0;
      }
      if (aStr1 === null) {
        return 1;
      }
      if (aStr2 === null) {
        return -1;
      }
      if (aStr1 > aStr2) {
        return 1;
      }
      return -1;
    }
    function compareByGeneratedPositionsInflated(mappingA, mappingB) {
      var cmp = mappingA.generatedLine - mappingB.generatedLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.generatedColumn - mappingB.generatedColumn;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = strcmp(mappingA.source, mappingB.source);
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalLine - mappingB.originalLine;
      if (cmp !== 0) {
        return cmp;
      }
      cmp = mappingA.originalColumn - mappingB.originalColumn;
      if (cmp !== 0) {
        return cmp;
      }
      return strcmp(mappingA.name, mappingB.name);
    }
    exports.compareByGeneratedPositionsInflated = compareByGeneratedPositionsInflated;
    function parseSourceMapInput(str) {
      return JSON.parse(str.replace(/^\)]}'[^\n]*\n/, ""));
    }
    exports.parseSourceMapInput = parseSourceMapInput;
    function computeSourceURL(sourceRoot, sourceURL, sourceMapURL) {
      sourceURL = sourceURL || "";
      if (sourceRoot) {
        if (sourceRoot[sourceRoot.length - 1] !== "/" && sourceURL[0] !== "/") {
          sourceRoot += "/";
        }
        sourceURL = sourceRoot + sourceURL;
      }
      if (sourceMapURL) {
        var parsed = urlParse(sourceMapURL);
        if (!parsed) {
          throw new Error("sourceMapURL could not be parsed");
        }
        if (parsed.path) {
          var index2 = parsed.path.lastIndexOf("/");
          if (index2 >= 0) {
            parsed.path = parsed.path.substring(0, index2 + 1);
          }
        }
        sourceURL = join(urlGenerate(parsed), sourceURL);
      }
      return normalize(sourceURL);
    }
    exports.computeSourceURL = computeSourceURL;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/array-set.js
var require_array_set = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/array-set.js"(exports) {
    var util = require_util();
    var has = Object.prototype.hasOwnProperty;
    var hasNativeMap = typeof Map !== "undefined";
    function ArraySet() {
      this._array = [];
      this._set = hasNativeMap ? /* @__PURE__ */ new Map() : /* @__PURE__ */ Object.create(null);
    }
    ArraySet.fromArray = function ArraySet_fromArray(aArray, aAllowDuplicates) {
      var set = new ArraySet();
      for (var i = 0, len = aArray.length; i < len; i++) {
        set.add(aArray[i], aAllowDuplicates);
      }
      return set;
    };
    ArraySet.prototype.size = function ArraySet_size() {
      return hasNativeMap ? this._set.size : Object.getOwnPropertyNames(this._set).length;
    };
    ArraySet.prototype.add = function ArraySet_add(aStr, aAllowDuplicates) {
      var sStr = hasNativeMap ? aStr : util.toSetString(aStr);
      var isDuplicate = hasNativeMap ? this.has(aStr) : has.call(this._set, sStr);
      var idx = this._array.length;
      if (!isDuplicate || aAllowDuplicates) {
        this._array.push(aStr);
      }
      if (!isDuplicate) {
        if (hasNativeMap) {
          this._set.set(aStr, idx);
        } else {
          this._set[sStr] = idx;
        }
      }
    };
    ArraySet.prototype.has = function ArraySet_has(aStr) {
      if (hasNativeMap) {
        return this._set.has(aStr);
      } else {
        var sStr = util.toSetString(aStr);
        return has.call(this._set, sStr);
      }
    };
    ArraySet.prototype.indexOf = function ArraySet_indexOf(aStr) {
      if (hasNativeMap) {
        var idx = this._set.get(aStr);
        if (idx >= 0) {
          return idx;
        }
      } else {
        var sStr = util.toSetString(aStr);
        if (has.call(this._set, sStr)) {
          return this._set[sStr];
        }
      }
      throw new Error('"' + aStr + '" is not in the set.');
    };
    ArraySet.prototype.at = function ArraySet_at(aIdx) {
      if (aIdx >= 0 && aIdx < this._array.length) {
        return this._array[aIdx];
      }
      throw new Error("No element indexed by " + aIdx);
    };
    ArraySet.prototype.toArray = function ArraySet_toArray() {
      return this._array.slice();
    };
    exports.ArraySet = ArraySet;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/mapping-list.js
var require_mapping_list = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/mapping-list.js"(exports) {
    var util = require_util();
    function generatedPositionAfter(mappingA, mappingB) {
      var lineA = mappingA.generatedLine;
      var lineB = mappingB.generatedLine;
      var columnA = mappingA.generatedColumn;
      var columnB = mappingB.generatedColumn;
      return lineB > lineA || lineB == lineA && columnB >= columnA || util.compareByGeneratedPositionsInflated(mappingA, mappingB) <= 0;
    }
    function MappingList() {
      this._array = [];
      this._sorted = true;
      this._last = { generatedLine: -1, generatedColumn: 0 };
    }
    MappingList.prototype.unsortedForEach = function MappingList_forEach(aCallback, aThisArg) {
      this._array.forEach(aCallback, aThisArg);
    };
    MappingList.prototype.add = function MappingList_add(aMapping) {
      if (generatedPositionAfter(this._last, aMapping)) {
        this._last = aMapping;
        this._array.push(aMapping);
      } else {
        this._sorted = false;
        this._array.push(aMapping);
      }
    };
    MappingList.prototype.toArray = function MappingList_toArray() {
      if (!this._sorted) {
        this._array.sort(util.compareByGeneratedPositionsInflated);
        this._sorted = true;
      }
      return this._array;
    };
    exports.MappingList = MappingList;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/source-map-generator.js
var require_source_map_generator = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/source-map-generator.js"(exports) {
    var base64VLQ = require_base64_vlq();
    var util = require_util();
    var ArraySet = require_array_set().ArraySet;
    var MappingList = require_mapping_list().MappingList;
    function SourceMapGenerator(aArgs) {
      if (!aArgs) {
        aArgs = {};
      }
      this._file = util.getArg(aArgs, "file", null);
      this._sourceRoot = util.getArg(aArgs, "sourceRoot", null);
      this._skipValidation = util.getArg(aArgs, "skipValidation", false);
      this._sources = new ArraySet();
      this._names = new ArraySet();
      this._mappings = new MappingList();
      this._sourcesContents = null;
    }
    SourceMapGenerator.prototype._version = 3;
    SourceMapGenerator.fromSourceMap = function SourceMapGenerator_fromSourceMap(aSourceMapConsumer) {
      var sourceRoot = aSourceMapConsumer.sourceRoot;
      var generator = new SourceMapGenerator({
        file: aSourceMapConsumer.file,
        sourceRoot
      });
      aSourceMapConsumer.eachMapping(function(mapping) {
        var newMapping = {
          generated: {
            line: mapping.generatedLine,
            column: mapping.generatedColumn
          }
        };
        if (mapping.source != null) {
          newMapping.source = mapping.source;
          if (sourceRoot != null) {
            newMapping.source = util.relative(sourceRoot, newMapping.source);
          }
          newMapping.original = {
            line: mapping.originalLine,
            column: mapping.originalColumn
          };
          if (mapping.name != null) {
            newMapping.name = mapping.name;
          }
        }
        generator.addMapping(newMapping);
      });
      aSourceMapConsumer.sources.forEach(function(sourceFile) {
        var sourceRelative = sourceFile;
        if (sourceRoot !== null) {
          sourceRelative = util.relative(sourceRoot, sourceFile);
        }
        if (!generator._sources.has(sourceRelative)) {
          generator._sources.add(sourceRelative);
        }
        var content = aSourceMapConsumer.sourceContentFor(sourceFile);
        if (content != null) {
          generator.setSourceContent(sourceFile, content);
        }
      });
      return generator;
    };
    SourceMapGenerator.prototype.addMapping = function SourceMapGenerator_addMapping(aArgs) {
      var generated = util.getArg(aArgs, "generated");
      var original = util.getArg(aArgs, "original", null);
      var source = util.getArg(aArgs, "source", null);
      var name = util.getArg(aArgs, "name", null);
      if (!this._skipValidation) {
        this._validateMapping(generated, original, source, name);
      }
      if (source != null) {
        source = String(source);
        if (!this._sources.has(source)) {
          this._sources.add(source);
        }
      }
      if (name != null) {
        name = String(name);
        if (!this._names.has(name)) {
          this._names.add(name);
        }
      }
      this._mappings.add({
        generatedLine: generated.line,
        generatedColumn: generated.column,
        originalLine: original != null && original.line,
        originalColumn: original != null && original.column,
        source,
        name
      });
    };
    SourceMapGenerator.prototype.setSourceContent = function SourceMapGenerator_setSourceContent(aSourceFile, aSourceContent) {
      var source = aSourceFile;
      if (this._sourceRoot != null) {
        source = util.relative(this._sourceRoot, source);
      }
      if (aSourceContent != null) {
        if (!this._sourcesContents) {
          this._sourcesContents = /* @__PURE__ */ Object.create(null);
        }
        this._sourcesContents[util.toSetString(source)] = aSourceContent;
      } else if (this._sourcesContents) {
        delete this._sourcesContents[util.toSetString(source)];
        if (Object.keys(this._sourcesContents).length === 0) {
          this._sourcesContents = null;
        }
      }
    };
    SourceMapGenerator.prototype.applySourceMap = function SourceMapGenerator_applySourceMap(aSourceMapConsumer, aSourceFile, aSourceMapPath) {
      var sourceFile = aSourceFile;
      if (aSourceFile == null) {
        if (aSourceMapConsumer.file == null) {
          throw new Error(
            `SourceMapGenerator.prototype.applySourceMap requires either an explicit source file, or the source map's "file" property. Both were omitted.`
          );
        }
        sourceFile = aSourceMapConsumer.file;
      }
      var sourceRoot = this._sourceRoot;
      if (sourceRoot != null) {
        sourceFile = util.relative(sourceRoot, sourceFile);
      }
      var newSources = new ArraySet();
      var newNames = new ArraySet();
      this._mappings.unsortedForEach(function(mapping) {
        if (mapping.source === sourceFile && mapping.originalLine != null) {
          var original = aSourceMapConsumer.originalPositionFor({
            line: mapping.originalLine,
            column: mapping.originalColumn
          });
          if (original.source != null) {
            mapping.source = original.source;
            if (aSourceMapPath != null) {
              mapping.source = util.join(aSourceMapPath, mapping.source);
            }
            if (sourceRoot != null) {
              mapping.source = util.relative(sourceRoot, mapping.source);
            }
            mapping.originalLine = original.line;
            mapping.originalColumn = original.column;
            if (original.name != null) {
              mapping.name = original.name;
            }
          }
        }
        var source = mapping.source;
        if (source != null && !newSources.has(source)) {
          newSources.add(source);
        }
        var name = mapping.name;
        if (name != null && !newNames.has(name)) {
          newNames.add(name);
        }
      }, this);
      this._sources = newSources;
      this._names = newNames;
      aSourceMapConsumer.sources.forEach(function(sourceFile2) {
        var content = aSourceMapConsumer.sourceContentFor(sourceFile2);
        if (content != null) {
          if (aSourceMapPath != null) {
            sourceFile2 = util.join(aSourceMapPath, sourceFile2);
          }
          if (sourceRoot != null) {
            sourceFile2 = util.relative(sourceRoot, sourceFile2);
          }
          this.setSourceContent(sourceFile2, content);
        }
      }, this);
    };
    SourceMapGenerator.prototype._validateMapping = function SourceMapGenerator_validateMapping(aGenerated, aOriginal, aSource, aName) {
      if (aOriginal && typeof aOriginal.line !== "number" && typeof aOriginal.column !== "number") {
        throw new Error(
          "original.line and original.column are not numbers -- you probably meant to omit the original mapping entirely and only map the generated position. If so, pass null for the original mapping instead of an object with empty or null values."
        );
      }
      if (aGenerated && "line" in aGenerated && "column" in aGenerated && aGenerated.line > 0 && aGenerated.column >= 0 && !aOriginal && !aSource && !aName) {
        return;
      } else if (aGenerated && "line" in aGenerated && "column" in aGenerated && aOriginal && "line" in aOriginal && "column" in aOriginal && aGenerated.line > 0 && aGenerated.column >= 0 && aOriginal.line > 0 && aOriginal.column >= 0 && aSource) {
        return;
      } else {
        throw new Error("Invalid mapping: " + JSON.stringify({
          generated: aGenerated,
          source: aSource,
          original: aOriginal,
          name: aName
        }));
      }
    };
    SourceMapGenerator.prototype._serializeMappings = function SourceMapGenerator_serializeMappings() {
      var previousGeneratedColumn = 0;
      var previousGeneratedLine = 1;
      var previousOriginalColumn = 0;
      var previousOriginalLine = 0;
      var previousName = 0;
      var previousSource = 0;
      var result = "";
      var next;
      var mapping;
      var nameIdx;
      var sourceIdx;
      var mappings = this._mappings.toArray();
      for (var i = 0, len = mappings.length; i < len; i++) {
        mapping = mappings[i];
        next = "";
        if (mapping.generatedLine !== previousGeneratedLine) {
          previousGeneratedColumn = 0;
          while (mapping.generatedLine !== previousGeneratedLine) {
            next += ";";
            previousGeneratedLine++;
          }
        } else {
          if (i > 0) {
            if (!util.compareByGeneratedPositionsInflated(mapping, mappings[i - 1])) {
              continue;
            }
            next += ",";
          }
        }
        next += base64VLQ.encode(mapping.generatedColumn - previousGeneratedColumn);
        previousGeneratedColumn = mapping.generatedColumn;
        if (mapping.source != null) {
          sourceIdx = this._sources.indexOf(mapping.source);
          next += base64VLQ.encode(sourceIdx - previousSource);
          previousSource = sourceIdx;
          next += base64VLQ.encode(mapping.originalLine - 1 - previousOriginalLine);
          previousOriginalLine = mapping.originalLine - 1;
          next += base64VLQ.encode(mapping.originalColumn - previousOriginalColumn);
          previousOriginalColumn = mapping.originalColumn;
          if (mapping.name != null) {
            nameIdx = this._names.indexOf(mapping.name);
            next += base64VLQ.encode(nameIdx - previousName);
            previousName = nameIdx;
          }
        }
        result += next;
      }
      return result;
    };
    SourceMapGenerator.prototype._generateSourcesContent = function SourceMapGenerator_generateSourcesContent(aSources, aSourceRoot) {
      return aSources.map(function(source) {
        if (!this._sourcesContents) {
          return null;
        }
        if (aSourceRoot != null) {
          source = util.relative(aSourceRoot, source);
        }
        var key = util.toSetString(source);
        return Object.prototype.hasOwnProperty.call(this._sourcesContents, key) ? this._sourcesContents[key] : null;
      }, this);
    };
    SourceMapGenerator.prototype.toJSON = function SourceMapGenerator_toJSON() {
      var map = {
        version: this._version,
        sources: this._sources.toArray(),
        names: this._names.toArray(),
        mappings: this._serializeMappings()
      };
      if (this._file != null) {
        map.file = this._file;
      }
      if (this._sourceRoot != null) {
        map.sourceRoot = this._sourceRoot;
      }
      if (this._sourcesContents) {
        map.sourcesContent = this._generateSourcesContent(map.sources, map.sourceRoot);
      }
      return map;
    };
    SourceMapGenerator.prototype.toString = function SourceMapGenerator_toString() {
      return JSON.stringify(this.toJSON());
    };
    exports.SourceMapGenerator = SourceMapGenerator;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/binary-search.js
var require_binary_search = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/binary-search.js"(exports) {
    exports.GREATEST_LOWER_BOUND = 1;
    exports.LEAST_UPPER_BOUND = 2;
    function recursiveSearch(aLow, aHigh, aNeedle, aHaystack, aCompare, aBias) {
      var mid = Math.floor((aHigh - aLow) / 2) + aLow;
      var cmp = aCompare(aNeedle, aHaystack[mid], true);
      if (cmp === 0) {
        return mid;
      } else if (cmp > 0) {
        if (aHigh - mid > 1) {
          return recursiveSearch(mid, aHigh, aNeedle, aHaystack, aCompare, aBias);
        }
        if (aBias == exports.LEAST_UPPER_BOUND) {
          return aHigh < aHaystack.length ? aHigh : -1;
        } else {
          return mid;
        }
      } else {
        if (mid - aLow > 1) {
          return recursiveSearch(aLow, mid, aNeedle, aHaystack, aCompare, aBias);
        }
        if (aBias == exports.LEAST_UPPER_BOUND) {
          return mid;
        } else {
          return aLow < 0 ? -1 : aLow;
        }
      }
    }
    exports.search = function search(aNeedle, aHaystack, aCompare, aBias) {
      if (aHaystack.length === 0) {
        return -1;
      }
      var index2 = recursiveSearch(
        -1,
        aHaystack.length,
        aNeedle,
        aHaystack,
        aCompare,
        aBias || exports.GREATEST_LOWER_BOUND
      );
      if (index2 < 0) {
        return -1;
      }
      while (index2 - 1 >= 0) {
        if (aCompare(aHaystack[index2], aHaystack[index2 - 1], true) !== 0) {
          break;
        }
        --index2;
      }
      return index2;
    };
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/quick-sort.js
var require_quick_sort = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/quick-sort.js"(exports) {
    function SortTemplate(comparator) {
      function swap(ary, x, y) {
        var temp = ary[x];
        ary[x] = ary[y];
        ary[y] = temp;
      }
      function randomIntInRange(low, high) {
        return Math.round(low + Math.random() * (high - low));
      }
      function doQuickSort(ary, comparator2, p, r) {
        if (p < r) {
          var pivotIndex = randomIntInRange(p, r);
          var i = p - 1;
          swap(ary, pivotIndex, r);
          var pivot = ary[r];
          for (var j = p; j < r; j++) {
            if (comparator2(ary[j], pivot, false) <= 0) {
              i += 1;
              swap(ary, i, j);
            }
          }
          swap(ary, i + 1, j);
          var q = i + 1;
          doQuickSort(ary, comparator2, p, q - 1);
          doQuickSort(ary, comparator2, q + 1, r);
        }
      }
      return doQuickSort;
    }
    function cloneSort(comparator) {
      let template = SortTemplate.toString();
      let templateFn = new Function(`return ${template}`)();
      return templateFn(comparator);
    }
    var sortCache = /* @__PURE__ */ new WeakMap();
    exports.quickSort = function(ary, comparator, start = 0) {
      let doQuickSort = sortCache.get(comparator);
      if (doQuickSort === void 0) {
        doQuickSort = cloneSort(comparator);
        sortCache.set(comparator, doQuickSort);
      }
      doQuickSort(ary, comparator, start, ary.length - 1);
    };
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/source-map-consumer.js
var require_source_map_consumer = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/source-map-consumer.js"(exports) {
    var util = require_util();
    var binarySearch = require_binary_search();
    var ArraySet = require_array_set().ArraySet;
    var base64VLQ = require_base64_vlq();
    var quickSort = require_quick_sort().quickSort;
    function SourceMapConsumer(aSourceMap, aSourceMapURL) {
      var sourceMap = aSourceMap;
      if (typeof aSourceMap === "string") {
        sourceMap = util.parseSourceMapInput(aSourceMap);
      }
      return sourceMap.sections != null ? new IndexedSourceMapConsumer(sourceMap, aSourceMapURL) : new BasicSourceMapConsumer(sourceMap, aSourceMapURL);
    }
    SourceMapConsumer.fromSourceMap = function(aSourceMap, aSourceMapURL) {
      return BasicSourceMapConsumer.fromSourceMap(aSourceMap, aSourceMapURL);
    };
    SourceMapConsumer.prototype._version = 3;
    SourceMapConsumer.prototype.__generatedMappings = null;
    Object.defineProperty(SourceMapConsumer.prototype, "_generatedMappings", {
      configurable: true,
      enumerable: true,
      get: function() {
        if (!this.__generatedMappings) {
          this._parseMappings(this._mappings, this.sourceRoot);
        }
        return this.__generatedMappings;
      }
    });
    SourceMapConsumer.prototype.__originalMappings = null;
    Object.defineProperty(SourceMapConsumer.prototype, "_originalMappings", {
      configurable: true,
      enumerable: true,
      get: function() {
        if (!this.__originalMappings) {
          this._parseMappings(this._mappings, this.sourceRoot);
        }
        return this.__originalMappings;
      }
    });
    SourceMapConsumer.prototype._charIsMappingSeparator = function SourceMapConsumer_charIsMappingSeparator(aStr, index2) {
      var c = aStr.charAt(index2);
      return c === ";" || c === ",";
    };
    SourceMapConsumer.prototype._parseMappings = function SourceMapConsumer_parseMappings(aStr, aSourceRoot) {
      throw new Error("Subclasses must implement _parseMappings");
    };
    SourceMapConsumer.GENERATED_ORDER = 1;
    SourceMapConsumer.ORIGINAL_ORDER = 2;
    SourceMapConsumer.GREATEST_LOWER_BOUND = 1;
    SourceMapConsumer.LEAST_UPPER_BOUND = 2;
    SourceMapConsumer.prototype.eachMapping = function SourceMapConsumer_eachMapping(aCallback, aContext, aOrder) {
      var context = aContext || null;
      var order = aOrder || SourceMapConsumer.GENERATED_ORDER;
      var mappings;
      switch (order) {
        case SourceMapConsumer.GENERATED_ORDER:
          mappings = this._generatedMappings;
          break;
        case SourceMapConsumer.ORIGINAL_ORDER:
          mappings = this._originalMappings;
          break;
        default:
          throw new Error("Unknown order of iteration.");
      }
      var sourceRoot = this.sourceRoot;
      var boundCallback = aCallback.bind(context);
      var names = this._names;
      var sources = this._sources;
      var sourceMapURL = this._sourceMapURL;
      for (var i = 0, n = mappings.length; i < n; i++) {
        var mapping = mappings[i];
        var source = mapping.source === null ? null : sources.at(mapping.source);
        source = util.computeSourceURL(sourceRoot, source, sourceMapURL);
        boundCallback({
          source,
          generatedLine: mapping.generatedLine,
          generatedColumn: mapping.generatedColumn,
          originalLine: mapping.originalLine,
          originalColumn: mapping.originalColumn,
          name: mapping.name === null ? null : names.at(mapping.name)
        });
      }
    };
    SourceMapConsumer.prototype.allGeneratedPositionsFor = function SourceMapConsumer_allGeneratedPositionsFor(aArgs) {
      var line = util.getArg(aArgs, "line");
      var needle = {
        source: util.getArg(aArgs, "source"),
        originalLine: line,
        originalColumn: util.getArg(aArgs, "column", 0)
      };
      needle.source = this._findSourceIndex(needle.source);
      if (needle.source < 0) {
        return [];
      }
      var mappings = [];
      var index2 = this._findMapping(
        needle,
        this._originalMappings,
        "originalLine",
        "originalColumn",
        util.compareByOriginalPositions,
        binarySearch.LEAST_UPPER_BOUND
      );
      if (index2 >= 0) {
        var mapping = this._originalMappings[index2];
        if (aArgs.column === void 0) {
          var originalLine = mapping.originalLine;
          while (mapping && mapping.originalLine === originalLine) {
            mappings.push({
              line: util.getArg(mapping, "generatedLine", null),
              column: util.getArg(mapping, "generatedColumn", null),
              lastColumn: util.getArg(mapping, "lastGeneratedColumn", null)
            });
            mapping = this._originalMappings[++index2];
          }
        } else {
          var originalColumn = mapping.originalColumn;
          while (mapping && mapping.originalLine === line && mapping.originalColumn == originalColumn) {
            mappings.push({
              line: util.getArg(mapping, "generatedLine", null),
              column: util.getArg(mapping, "generatedColumn", null),
              lastColumn: util.getArg(mapping, "lastGeneratedColumn", null)
            });
            mapping = this._originalMappings[++index2];
          }
        }
      }
      return mappings;
    };
    exports.SourceMapConsumer = SourceMapConsumer;
    function BasicSourceMapConsumer(aSourceMap, aSourceMapURL) {
      var sourceMap = aSourceMap;
      if (typeof aSourceMap === "string") {
        sourceMap = util.parseSourceMapInput(aSourceMap);
      }
      var version = util.getArg(sourceMap, "version");
      var sources = util.getArg(sourceMap, "sources");
      var names = util.getArg(sourceMap, "names", []);
      var sourceRoot = util.getArg(sourceMap, "sourceRoot", null);
      var sourcesContent = util.getArg(sourceMap, "sourcesContent", null);
      var mappings = util.getArg(sourceMap, "mappings");
      var file = util.getArg(sourceMap, "file", null);
      if (version != this._version) {
        throw new Error("Unsupported version: " + version);
      }
      if (sourceRoot) {
        sourceRoot = util.normalize(sourceRoot);
      }
      sources = sources.map(String).map(util.normalize).map(function(source) {
        return sourceRoot && util.isAbsolute(sourceRoot) && util.isAbsolute(source) ? util.relative(sourceRoot, source) : source;
      });
      this._names = ArraySet.fromArray(names.map(String), true);
      this._sources = ArraySet.fromArray(sources, true);
      this._absoluteSources = this._sources.toArray().map(function(s) {
        return util.computeSourceURL(sourceRoot, s, aSourceMapURL);
      });
      this.sourceRoot = sourceRoot;
      this.sourcesContent = sourcesContent;
      this._mappings = mappings;
      this._sourceMapURL = aSourceMapURL;
      this.file = file;
    }
    BasicSourceMapConsumer.prototype = Object.create(SourceMapConsumer.prototype);
    BasicSourceMapConsumer.prototype.consumer = SourceMapConsumer;
    BasicSourceMapConsumer.prototype._findSourceIndex = function(aSource) {
      var relativeSource = aSource;
      if (this.sourceRoot != null) {
        relativeSource = util.relative(this.sourceRoot, relativeSource);
      }
      if (this._sources.has(relativeSource)) {
        return this._sources.indexOf(relativeSource);
      }
      var i;
      for (i = 0; i < this._absoluteSources.length; ++i) {
        if (this._absoluteSources[i] == aSource) {
          return i;
        }
      }
      return -1;
    };
    BasicSourceMapConsumer.fromSourceMap = function SourceMapConsumer_fromSourceMap(aSourceMap, aSourceMapURL) {
      var smc = Object.create(BasicSourceMapConsumer.prototype);
      var names = smc._names = ArraySet.fromArray(aSourceMap._names.toArray(), true);
      var sources = smc._sources = ArraySet.fromArray(aSourceMap._sources.toArray(), true);
      smc.sourceRoot = aSourceMap._sourceRoot;
      smc.sourcesContent = aSourceMap._generateSourcesContent(
        smc._sources.toArray(),
        smc.sourceRoot
      );
      smc.file = aSourceMap._file;
      smc._sourceMapURL = aSourceMapURL;
      smc._absoluteSources = smc._sources.toArray().map(function(s) {
        return util.computeSourceURL(smc.sourceRoot, s, aSourceMapURL);
      });
      var generatedMappings = aSourceMap._mappings.toArray().slice();
      var destGeneratedMappings = smc.__generatedMappings = [];
      var destOriginalMappings = smc.__originalMappings = [];
      for (var i = 0, length = generatedMappings.length; i < length; i++) {
        var srcMapping = generatedMappings[i];
        var destMapping = new Mapping();
        destMapping.generatedLine = srcMapping.generatedLine;
        destMapping.generatedColumn = srcMapping.generatedColumn;
        if (srcMapping.source) {
          destMapping.source = sources.indexOf(srcMapping.source);
          destMapping.originalLine = srcMapping.originalLine;
          destMapping.originalColumn = srcMapping.originalColumn;
          if (srcMapping.name) {
            destMapping.name = names.indexOf(srcMapping.name);
          }
          destOriginalMappings.push(destMapping);
        }
        destGeneratedMappings.push(destMapping);
      }
      quickSort(smc.__originalMappings, util.compareByOriginalPositions);
      return smc;
    };
    BasicSourceMapConsumer.prototype._version = 3;
    Object.defineProperty(BasicSourceMapConsumer.prototype, "sources", {
      get: function() {
        return this._absoluteSources.slice();
      }
    });
    function Mapping() {
      this.generatedLine = 0;
      this.generatedColumn = 0;
      this.source = null;
      this.originalLine = null;
      this.originalColumn = null;
      this.name = null;
    }
    var compareGenerated = util.compareByGeneratedPositionsDeflatedNoLine;
    function sortGenerated(array, start) {
      let l = array.length;
      let n = array.length - start;
      if (n <= 1) {
        return;
      } else if (n == 2) {
        let a = array[start];
        let b = array[start + 1];
        if (compareGenerated(a, b) > 0) {
          array[start] = b;
          array[start + 1] = a;
        }
      } else if (n < 20) {
        for (let i = start; i < l; i++) {
          for (let j = i; j > start; j--) {
            let a = array[j - 1];
            let b = array[j];
            if (compareGenerated(a, b) <= 0) {
              break;
            }
            array[j - 1] = b;
            array[j] = a;
          }
        }
      } else {
        quickSort(array, compareGenerated, start);
      }
    }
    BasicSourceMapConsumer.prototype._parseMappings = function SourceMapConsumer_parseMappings(aStr, aSourceRoot) {
      var generatedLine = 1;
      var previousGeneratedColumn = 0;
      var previousOriginalLine = 0;
      var previousOriginalColumn = 0;
      var previousSource = 0;
      var previousName = 0;
      var length = aStr.length;
      var index2 = 0;
      var cachedSegments = {};
      var temp = {};
      var originalMappings = [];
      var generatedMappings = [];
      var mapping, str, segment, end, value;
      let subarrayStart = 0;
      while (index2 < length) {
        if (aStr.charAt(index2) === ";") {
          generatedLine++;
          index2++;
          previousGeneratedColumn = 0;
          sortGenerated(generatedMappings, subarrayStart);
          subarrayStart = generatedMappings.length;
        } else if (aStr.charAt(index2) === ",") {
          index2++;
        } else {
          mapping = new Mapping();
          mapping.generatedLine = generatedLine;
          for (end = index2; end < length; end++) {
            if (this._charIsMappingSeparator(aStr, end)) {
              break;
            }
          }
          str = aStr.slice(index2, end);
          segment = [];
          while (index2 < end) {
            base64VLQ.decode(aStr, index2, temp);
            value = temp.value;
            index2 = temp.rest;
            segment.push(value);
          }
          if (segment.length === 2) {
            throw new Error("Found a source, but no line and column");
          }
          if (segment.length === 3) {
            throw new Error("Found a source and line, but no column");
          }
          mapping.generatedColumn = previousGeneratedColumn + segment[0];
          previousGeneratedColumn = mapping.generatedColumn;
          if (segment.length > 1) {
            mapping.source = previousSource + segment[1];
            previousSource += segment[1];
            mapping.originalLine = previousOriginalLine + segment[2];
            previousOriginalLine = mapping.originalLine;
            mapping.originalLine += 1;
            mapping.originalColumn = previousOriginalColumn + segment[3];
            previousOriginalColumn = mapping.originalColumn;
            if (segment.length > 4) {
              mapping.name = previousName + segment[4];
              previousName += segment[4];
            }
          }
          generatedMappings.push(mapping);
          if (typeof mapping.originalLine === "number") {
            let currentSource = mapping.source;
            while (originalMappings.length <= currentSource) {
              originalMappings.push(null);
            }
            if (originalMappings[currentSource] === null) {
              originalMappings[currentSource] = [];
            }
            originalMappings[currentSource].push(mapping);
          }
        }
      }
      sortGenerated(generatedMappings, subarrayStart);
      this.__generatedMappings = generatedMappings;
      for (var i = 0; i < originalMappings.length; i++) {
        if (originalMappings[i] != null) {
          quickSort(originalMappings[i], util.compareByOriginalPositionsNoSource);
        }
      }
      this.__originalMappings = [].concat(...originalMappings);
    };
    BasicSourceMapConsumer.prototype._findMapping = function SourceMapConsumer_findMapping(aNeedle, aMappings, aLineName, aColumnName, aComparator, aBias) {
      if (aNeedle[aLineName] <= 0) {
        throw new TypeError("Line must be greater than or equal to 1, got " + aNeedle[aLineName]);
      }
      if (aNeedle[aColumnName] < 0) {
        throw new TypeError("Column must be greater than or equal to 0, got " + aNeedle[aColumnName]);
      }
      return binarySearch.search(aNeedle, aMappings, aComparator, aBias);
    };
    BasicSourceMapConsumer.prototype.computeColumnSpans = function SourceMapConsumer_computeColumnSpans() {
      for (var index2 = 0; index2 < this._generatedMappings.length; ++index2) {
        var mapping = this._generatedMappings[index2];
        if (index2 + 1 < this._generatedMappings.length) {
          var nextMapping = this._generatedMappings[index2 + 1];
          if (mapping.generatedLine === nextMapping.generatedLine) {
            mapping.lastGeneratedColumn = nextMapping.generatedColumn - 1;
            continue;
          }
        }
        mapping.lastGeneratedColumn = Infinity;
      }
    };
    BasicSourceMapConsumer.prototype.originalPositionFor = function SourceMapConsumer_originalPositionFor(aArgs) {
      var needle = {
        generatedLine: util.getArg(aArgs, "line"),
        generatedColumn: util.getArg(aArgs, "column")
      };
      var index2 = this._findMapping(
        needle,
        this._generatedMappings,
        "generatedLine",
        "generatedColumn",
        util.compareByGeneratedPositionsDeflated,
        util.getArg(aArgs, "bias", SourceMapConsumer.GREATEST_LOWER_BOUND)
      );
      if (index2 >= 0) {
        var mapping = this._generatedMappings[index2];
        if (mapping.generatedLine === needle.generatedLine) {
          var source = util.getArg(mapping, "source", null);
          if (source !== null) {
            source = this._sources.at(source);
            source = util.computeSourceURL(this.sourceRoot, source, this._sourceMapURL);
          }
          var name = util.getArg(mapping, "name", null);
          if (name !== null) {
            name = this._names.at(name);
          }
          return {
            source,
            line: util.getArg(mapping, "originalLine", null),
            column: util.getArg(mapping, "originalColumn", null),
            name
          };
        }
      }
      return {
        source: null,
        line: null,
        column: null,
        name: null
      };
    };
    BasicSourceMapConsumer.prototype.hasContentsOfAllSources = function BasicSourceMapConsumer_hasContentsOfAllSources() {
      if (!this.sourcesContent) {
        return false;
      }
      return this.sourcesContent.length >= this._sources.size() && !this.sourcesContent.some(function(sc) {
        return sc == null;
      });
    };
    BasicSourceMapConsumer.prototype.sourceContentFor = function SourceMapConsumer_sourceContentFor(aSource, nullOnMissing) {
      if (!this.sourcesContent) {
        return null;
      }
      var index2 = this._findSourceIndex(aSource);
      if (index2 >= 0) {
        return this.sourcesContent[index2];
      }
      var relativeSource = aSource;
      if (this.sourceRoot != null) {
        relativeSource = util.relative(this.sourceRoot, relativeSource);
      }
      var url;
      if (this.sourceRoot != null && (url = util.urlParse(this.sourceRoot))) {
        var fileUriAbsPath = relativeSource.replace(/^file:\/\//, "");
        if (url.scheme == "file" && this._sources.has(fileUriAbsPath)) {
          return this.sourcesContent[this._sources.indexOf(fileUriAbsPath)];
        }
        if ((!url.path || url.path == "/") && this._sources.has("/" + relativeSource)) {
          return this.sourcesContent[this._sources.indexOf("/" + relativeSource)];
        }
      }
      if (nullOnMissing) {
        return null;
      } else {
        throw new Error('"' + relativeSource + '" is not in the SourceMap.');
      }
    };
    BasicSourceMapConsumer.prototype.generatedPositionFor = function SourceMapConsumer_generatedPositionFor(aArgs) {
      var source = util.getArg(aArgs, "source");
      source = this._findSourceIndex(source);
      if (source < 0) {
        return {
          line: null,
          column: null,
          lastColumn: null
        };
      }
      var needle = {
        source,
        originalLine: util.getArg(aArgs, "line"),
        originalColumn: util.getArg(aArgs, "column")
      };
      var index2 = this._findMapping(
        needle,
        this._originalMappings,
        "originalLine",
        "originalColumn",
        util.compareByOriginalPositions,
        util.getArg(aArgs, "bias", SourceMapConsumer.GREATEST_LOWER_BOUND)
      );
      if (index2 >= 0) {
        var mapping = this._originalMappings[index2];
        if (mapping.source === needle.source) {
          return {
            line: util.getArg(mapping, "generatedLine", null),
            column: util.getArg(mapping, "generatedColumn", null),
            lastColumn: util.getArg(mapping, "lastGeneratedColumn", null)
          };
        }
      }
      return {
        line: null,
        column: null,
        lastColumn: null
      };
    };
    exports.BasicSourceMapConsumer = BasicSourceMapConsumer;
    function IndexedSourceMapConsumer(aSourceMap, aSourceMapURL) {
      var sourceMap = aSourceMap;
      if (typeof aSourceMap === "string") {
        sourceMap = util.parseSourceMapInput(aSourceMap);
      }
      var version = util.getArg(sourceMap, "version");
      var sections = util.getArg(sourceMap, "sections");
      if (version != this._version) {
        throw new Error("Unsupported version: " + version);
      }
      this._sources = new ArraySet();
      this._names = new ArraySet();
      var lastOffset = {
        line: -1,
        column: 0
      };
      this._sections = sections.map(function(s) {
        if (s.url) {
          throw new Error("Support for url field in sections not implemented.");
        }
        var offset = util.getArg(s, "offset");
        var offsetLine = util.getArg(offset, "line");
        var offsetColumn = util.getArg(offset, "column");
        if (offsetLine < lastOffset.line || offsetLine === lastOffset.line && offsetColumn < lastOffset.column) {
          throw new Error("Section offsets must be ordered and non-overlapping.");
        }
        lastOffset = offset;
        return {
          generatedOffset: {
            // The offset fields are 0-based, but we use 1-based indices when
            // encoding/decoding from VLQ.
            generatedLine: offsetLine + 1,
            generatedColumn: offsetColumn + 1
          },
          consumer: new SourceMapConsumer(util.getArg(s, "map"), aSourceMapURL)
        };
      });
    }
    IndexedSourceMapConsumer.prototype = Object.create(SourceMapConsumer.prototype);
    IndexedSourceMapConsumer.prototype.constructor = SourceMapConsumer;
    IndexedSourceMapConsumer.prototype._version = 3;
    Object.defineProperty(IndexedSourceMapConsumer.prototype, "sources", {
      get: function() {
        var sources = [];
        for (var i = 0; i < this._sections.length; i++) {
          for (var j = 0; j < this._sections[i].consumer.sources.length; j++) {
            sources.push(this._sections[i].consumer.sources[j]);
          }
        }
        return sources;
      }
    });
    IndexedSourceMapConsumer.prototype.originalPositionFor = function IndexedSourceMapConsumer_originalPositionFor(aArgs) {
      var needle = {
        generatedLine: util.getArg(aArgs, "line"),
        generatedColumn: util.getArg(aArgs, "column")
      };
      var sectionIndex = binarySearch.search(
        needle,
        this._sections,
        function(needle2, section2) {
          var cmp = needle2.generatedLine - section2.generatedOffset.generatedLine;
          if (cmp) {
            return cmp;
          }
          return needle2.generatedColumn - section2.generatedOffset.generatedColumn;
        }
      );
      var section = this._sections[sectionIndex];
      if (!section) {
        return {
          source: null,
          line: null,
          column: null,
          name: null
        };
      }
      return section.consumer.originalPositionFor({
        line: needle.generatedLine - (section.generatedOffset.generatedLine - 1),
        column: needle.generatedColumn - (section.generatedOffset.generatedLine === needle.generatedLine ? section.generatedOffset.generatedColumn - 1 : 0),
        bias: aArgs.bias
      });
    };
    IndexedSourceMapConsumer.prototype.hasContentsOfAllSources = function IndexedSourceMapConsumer_hasContentsOfAllSources() {
      return this._sections.every(function(s) {
        return s.consumer.hasContentsOfAllSources();
      });
    };
    IndexedSourceMapConsumer.prototype.sourceContentFor = function IndexedSourceMapConsumer_sourceContentFor(aSource, nullOnMissing) {
      for (var i = 0; i < this._sections.length; i++) {
        var section = this._sections[i];
        var content = section.consumer.sourceContentFor(aSource, true);
        if (content) {
          return content;
        }
      }
      if (nullOnMissing) {
        return null;
      } else {
        throw new Error('"' + aSource + '" is not in the SourceMap.');
      }
    };
    IndexedSourceMapConsumer.prototype.generatedPositionFor = function IndexedSourceMapConsumer_generatedPositionFor(aArgs) {
      for (var i = 0; i < this._sections.length; i++) {
        var section = this._sections[i];
        if (section.consumer._findSourceIndex(util.getArg(aArgs, "source")) === -1) {
          continue;
        }
        var generatedPosition = section.consumer.generatedPositionFor(aArgs);
        if (generatedPosition) {
          var ret = {
            line: generatedPosition.line + (section.generatedOffset.generatedLine - 1),
            column: generatedPosition.column + (section.generatedOffset.generatedLine === generatedPosition.line ? section.generatedOffset.generatedColumn - 1 : 0)
          };
          return ret;
        }
      }
      return {
        line: null,
        column: null
      };
    };
    IndexedSourceMapConsumer.prototype._parseMappings = function IndexedSourceMapConsumer_parseMappings(aStr, aSourceRoot) {
      this.__generatedMappings = [];
      this.__originalMappings = [];
      for (var i = 0; i < this._sections.length; i++) {
        var section = this._sections[i];
        var sectionMappings = section.consumer._generatedMappings;
        for (var j = 0; j < sectionMappings.length; j++) {
          var mapping = sectionMappings[j];
          var source = section.consumer._sources.at(mapping.source);
          source = util.computeSourceURL(section.consumer.sourceRoot, source, this._sourceMapURL);
          this._sources.add(source);
          source = this._sources.indexOf(source);
          var name = null;
          if (mapping.name) {
            name = section.consumer._names.at(mapping.name);
            this._names.add(name);
            name = this._names.indexOf(name);
          }
          var adjustedMapping = {
            source,
            generatedLine: mapping.generatedLine + (section.generatedOffset.generatedLine - 1),
            generatedColumn: mapping.generatedColumn + (section.generatedOffset.generatedLine === mapping.generatedLine ? section.generatedOffset.generatedColumn - 1 : 0),
            originalLine: mapping.originalLine,
            originalColumn: mapping.originalColumn,
            name
          };
          this.__generatedMappings.push(adjustedMapping);
          if (typeof adjustedMapping.originalLine === "number") {
            this.__originalMappings.push(adjustedMapping);
          }
        }
      }
      quickSort(this.__generatedMappings, util.compareByGeneratedPositionsDeflated);
      quickSort(this.__originalMappings, util.compareByOriginalPositions);
    };
    exports.IndexedSourceMapConsumer = IndexedSourceMapConsumer;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/source-node.js
var require_source_node = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/lib/source-node.js"(exports) {
    var SourceMapGenerator = require_source_map_generator().SourceMapGenerator;
    var util = require_util();
    var REGEX_NEWLINE = /(\r?\n)/;
    var NEWLINE_CODE = 10;
    var isSourceNode = "$$$isSourceNode$$$";
    function SourceNode(aLine, aColumn, aSource, aChunks, aName) {
      this.children = [];
      this.sourceContents = {};
      this.line = aLine == null ? null : aLine;
      this.column = aColumn == null ? null : aColumn;
      this.source = aSource == null ? null : aSource;
      this.name = aName == null ? null : aName;
      this[isSourceNode] = true;
      if (aChunks != null)
        this.add(aChunks);
    }
    SourceNode.fromStringWithSourceMap = function SourceNode_fromStringWithSourceMap(aGeneratedCode, aSourceMapConsumer, aRelativePath) {
      var node = new SourceNode();
      var remainingLines = aGeneratedCode.split(REGEX_NEWLINE);
      var remainingLinesIndex = 0;
      var shiftNextLine = function() {
        var lineContents = getNextLine();
        var newLine = getNextLine() || "";
        return lineContents + newLine;
        function getNextLine() {
          return remainingLinesIndex < remainingLines.length ? remainingLines[remainingLinesIndex++] : void 0;
        }
      };
      var lastGeneratedLine = 1, lastGeneratedColumn = 0;
      var lastMapping = null;
      aSourceMapConsumer.eachMapping(function(mapping) {
        if (lastMapping !== null) {
          if (lastGeneratedLine < mapping.generatedLine) {
            addMappingWithCode(lastMapping, shiftNextLine());
            lastGeneratedLine++;
            lastGeneratedColumn = 0;
          } else {
            var nextLine = remainingLines[remainingLinesIndex] || "";
            var code = nextLine.substr(0, mapping.generatedColumn - lastGeneratedColumn);
            remainingLines[remainingLinesIndex] = nextLine.substr(mapping.generatedColumn - lastGeneratedColumn);
            lastGeneratedColumn = mapping.generatedColumn;
            addMappingWithCode(lastMapping, code);
            lastMapping = mapping;
            return;
          }
        }
        while (lastGeneratedLine < mapping.generatedLine) {
          node.add(shiftNextLine());
          lastGeneratedLine++;
        }
        if (lastGeneratedColumn < mapping.generatedColumn) {
          var nextLine = remainingLines[remainingLinesIndex] || "";
          node.add(nextLine.substr(0, mapping.generatedColumn));
          remainingLines[remainingLinesIndex] = nextLine.substr(mapping.generatedColumn);
          lastGeneratedColumn = mapping.generatedColumn;
        }
        lastMapping = mapping;
      }, this);
      if (remainingLinesIndex < remainingLines.length) {
        if (lastMapping) {
          addMappingWithCode(lastMapping, shiftNextLine());
        }
        node.add(remainingLines.splice(remainingLinesIndex).join(""));
      }
      aSourceMapConsumer.sources.forEach(function(sourceFile) {
        var content = aSourceMapConsumer.sourceContentFor(sourceFile);
        if (content != null) {
          if (aRelativePath != null) {
            sourceFile = util.join(aRelativePath, sourceFile);
          }
          node.setSourceContent(sourceFile, content);
        }
      });
      return node;
      function addMappingWithCode(mapping, code) {
        if (mapping === null || mapping.source === void 0) {
          node.add(code);
        } else {
          var source = aRelativePath ? util.join(aRelativePath, mapping.source) : mapping.source;
          node.add(new SourceNode(
            mapping.originalLine,
            mapping.originalColumn,
            source,
            code,
            mapping.name
          ));
        }
      }
    };
    SourceNode.prototype.add = function SourceNode_add(aChunk) {
      if (Array.isArray(aChunk)) {
        aChunk.forEach(function(chunk) {
          this.add(chunk);
        }, this);
      } else if (aChunk[isSourceNode] || typeof aChunk === "string") {
        if (aChunk) {
          this.children.push(aChunk);
        }
      } else {
        throw new TypeError(
          "Expected a SourceNode, string, or an array of SourceNodes and strings. Got " + aChunk
        );
      }
      return this;
    };
    SourceNode.prototype.prepend = function SourceNode_prepend(aChunk) {
      if (Array.isArray(aChunk)) {
        for (var i = aChunk.length - 1; i >= 0; i--) {
          this.prepend(aChunk[i]);
        }
      } else if (aChunk[isSourceNode] || typeof aChunk === "string") {
        this.children.unshift(aChunk);
      } else {
        throw new TypeError(
          "Expected a SourceNode, string, or an array of SourceNodes and strings. Got " + aChunk
        );
      }
      return this;
    };
    SourceNode.prototype.walk = function SourceNode_walk(aFn) {
      var chunk;
      for (var i = 0, len = this.children.length; i < len; i++) {
        chunk = this.children[i];
        if (chunk[isSourceNode]) {
          chunk.walk(aFn);
        } else {
          if (chunk !== "") {
            aFn(chunk, {
              source: this.source,
              line: this.line,
              column: this.column,
              name: this.name
            });
          }
        }
      }
    };
    SourceNode.prototype.join = function SourceNode_join(aSep) {
      var newChildren;
      var i;
      var len = this.children.length;
      if (len > 0) {
        newChildren = [];
        for (i = 0; i < len - 1; i++) {
          newChildren.push(this.children[i]);
          newChildren.push(aSep);
        }
        newChildren.push(this.children[i]);
        this.children = newChildren;
      }
      return this;
    };
    SourceNode.prototype.replaceRight = function SourceNode_replaceRight(aPattern, aReplacement) {
      var lastChild = this.children[this.children.length - 1];
      if (lastChild[isSourceNode]) {
        lastChild.replaceRight(aPattern, aReplacement);
      } else if (typeof lastChild === "string") {
        this.children[this.children.length - 1] = lastChild.replace(aPattern, aReplacement);
      } else {
        this.children.push("".replace(aPattern, aReplacement));
      }
      return this;
    };
    SourceNode.prototype.setSourceContent = function SourceNode_setSourceContent(aSourceFile, aSourceContent) {
      this.sourceContents[util.toSetString(aSourceFile)] = aSourceContent;
    };
    SourceNode.prototype.walkSourceContents = function SourceNode_walkSourceContents(aFn) {
      for (var i = 0, len = this.children.length; i < len; i++) {
        if (this.children[i][isSourceNode]) {
          this.children[i].walkSourceContents(aFn);
        }
      }
      var sources = Object.keys(this.sourceContents);
      for (var i = 0, len = sources.length; i < len; i++) {
        aFn(util.fromSetString(sources[i]), this.sourceContents[sources[i]]);
      }
    };
    SourceNode.prototype.toString = function SourceNode_toString() {
      var str = "";
      this.walk(function(chunk) {
        str += chunk;
      });
      return str;
    };
    SourceNode.prototype.toStringWithSourceMap = function SourceNode_toStringWithSourceMap(aArgs) {
      var generated = {
        code: "",
        line: 1,
        column: 0
      };
      var map = new SourceMapGenerator(aArgs);
      var sourceMappingActive = false;
      var lastOriginalSource = null;
      var lastOriginalLine = null;
      var lastOriginalColumn = null;
      var lastOriginalName = null;
      this.walk(function(chunk, original) {
        generated.code += chunk;
        if (original.source !== null && original.line !== null && original.column !== null) {
          if (lastOriginalSource !== original.source || lastOriginalLine !== original.line || lastOriginalColumn !== original.column || lastOriginalName !== original.name) {
            map.addMapping({
              source: original.source,
              original: {
                line: original.line,
                column: original.column
              },
              generated: {
                line: generated.line,
                column: generated.column
              },
              name: original.name
            });
          }
          lastOriginalSource = original.source;
          lastOriginalLine = original.line;
          lastOriginalColumn = original.column;
          lastOriginalName = original.name;
          sourceMappingActive = true;
        } else if (sourceMappingActive) {
          map.addMapping({
            generated: {
              line: generated.line,
              column: generated.column
            }
          });
          lastOriginalSource = null;
          sourceMappingActive = false;
        }
        for (var idx = 0, length = chunk.length; idx < length; idx++) {
          if (chunk.charCodeAt(idx) === NEWLINE_CODE) {
            generated.line++;
            generated.column = 0;
            if (idx + 1 === length) {
              lastOriginalSource = null;
              sourceMappingActive = false;
            } else if (sourceMappingActive) {
              map.addMapping({
                source: original.source,
                original: {
                  line: original.line,
                  column: original.column
                },
                generated: {
                  line: generated.line,
                  column: generated.column
                },
                name: original.name
              });
            }
          } else {
            generated.column++;
          }
        }
      });
      this.walkSourceContents(function(sourceFile, sourceContent) {
        map.setSourceContent(sourceFile, sourceContent);
      });
      return { code: generated.code, map };
    };
    exports.SourceNode = SourceNode;
  }
});

// node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/source-map.js
var require_source_map = __commonJS({
  "node_modules/.pnpm/source-map-js@1.0.2/node_modules/source-map-js/source-map.js"(exports) {
    exports.SourceMapGenerator = require_source_map_generator().SourceMapGenerator;
    exports.SourceMapConsumer = require_source_map_consumer().SourceMapConsumer;
    exports.SourceNode = require_source_node().SourceNode;
  }
});

// node_modules/.pnpm/nanoid@3.3.6/node_modules/nanoid/non-secure/index.cjs
var require_non_secure = __commonJS({
  "node_modules/.pnpm/nanoid@3.3.6/node_modules/nanoid/non-secure/index.cjs"(exports, module2) {
    var urlAlphabet = "useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict";
    var customAlphabet = (alphabet, defaultSize = 21) => {
      return (size = defaultSize) => {
        let id = "";
        let i = size;
        while (i--) {
          id += alphabet[Math.random() * alphabet.length | 0];
        }
        return id;
      };
    };
    var nanoid = (size = 21) => {
      let id = "";
      let i = size;
      while (i--) {
        id += urlAlphabet[Math.random() * 64 | 0];
      }
      return id;
    };
    module2.exports = { nanoid, customAlphabet };
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/previous-map.js
var require_previous_map = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/previous-map.js"(exports, module2) {
    "use strict";
    var { SourceMapConsumer, SourceMapGenerator } = require_source_map();
    var { existsSync, readFileSync } = require("fs");
    var { dirname, join } = require("path");
    function fromBase64(str) {
      if (Buffer) {
        return Buffer.from(str, "base64").toString();
      } else {
        return window.atob(str);
      }
    }
    var PreviousMap = class {
      constructor(css, opts) {
        if (opts.map === false)
          return;
        this.loadAnnotation(css);
        this.inline = this.startWith(this.annotation, "data:");
        let prev = opts.map ? opts.map.prev : void 0;
        let text2 = this.loadMap(opts.from, prev);
        if (!this.mapFile && opts.from) {
          this.mapFile = opts.from;
        }
        if (this.mapFile)
          this.root = dirname(this.mapFile);
        if (text2)
          this.text = text2;
      }
      consumer() {
        if (!this.consumerCache) {
          this.consumerCache = new SourceMapConsumer(this.text);
        }
        return this.consumerCache;
      }
      withContent() {
        return !!(this.consumer().sourcesContent && this.consumer().sourcesContent.length > 0);
      }
      startWith(string, start) {
        if (!string)
          return false;
        return string.substr(0, start.length) === start;
      }
      getAnnotationURL(sourceMapString) {
        return sourceMapString.replace(/^\/\*\s*# sourceMappingURL=/, "").trim();
      }
      loadAnnotation(css) {
        let comments = css.match(/\/\*\s*# sourceMappingURL=/gm);
        if (!comments)
          return;
        let start = css.lastIndexOf(comments.pop());
        let end = css.indexOf("*/", start);
        if (start > -1 && end > -1) {
          this.annotation = this.getAnnotationURL(css.substring(start, end));
        }
      }
      decodeInline(text2) {
        let baseCharsetUri = /^data:application\/json;charset=utf-?8;base64,/;
        let baseUri = /^data:application\/json;base64,/;
        let charsetUri = /^data:application\/json;charset=utf-?8,/;
        let uri = /^data:application\/json,/;
        if (charsetUri.test(text2) || uri.test(text2)) {
          return decodeURIComponent(text2.substr(RegExp.lastMatch.length));
        }
        if (baseCharsetUri.test(text2) || baseUri.test(text2)) {
          return fromBase64(text2.substr(RegExp.lastMatch.length));
        }
        let encoding = text2.match(/data:application\/json;([^,]+),/)[1];
        throw new Error("Unsupported source map encoding " + encoding);
      }
      loadFile(path) {
        this.root = dirname(path);
        if (existsSync(path)) {
          this.mapFile = path;
          return readFileSync(path, "utf-8").toString().trim();
        }
      }
      loadMap(file, prev) {
        if (prev === false)
          return false;
        if (prev) {
          if (typeof prev === "string") {
            return prev;
          } else if (typeof prev === "function") {
            let prevPath = prev(file);
            if (prevPath) {
              let map = this.loadFile(prevPath);
              if (!map) {
                throw new Error(
                  "Unable to load previous source map: " + prevPath.toString()
                );
              }
              return map;
            }
          } else if (prev instanceof SourceMapConsumer) {
            return SourceMapGenerator.fromSourceMap(prev).toString();
          } else if (prev instanceof SourceMapGenerator) {
            return prev.toString();
          } else if (this.isMap(prev)) {
            return JSON.stringify(prev);
          } else {
            throw new Error(
              "Unsupported previous source map format: " + prev.toString()
            );
          }
        } else if (this.inline) {
          return this.decodeInline(this.annotation);
        } else if (this.annotation) {
          let map = this.annotation;
          if (file)
            map = join(dirname(file), map);
          return this.loadFile(map);
        }
      }
      isMap(map) {
        if (typeof map !== "object")
          return false;
        return typeof map.mappings === "string" || typeof map._mappings === "string" || Array.isArray(map.sections);
      }
    };
    module2.exports = PreviousMap;
    PreviousMap.default = PreviousMap;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/input.js
var require_input = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/input.js"(exports, module2) {
    "use strict";
    var { SourceMapConsumer, SourceMapGenerator } = require_source_map();
    var { fileURLToPath, pathToFileURL } = require("url");
    var { resolve, isAbsolute } = require("path");
    var { nanoid } = require_non_secure();
    var terminalHighlight = require_terminal_highlight();
    var CssSyntaxError = require_css_syntax_error();
    var PreviousMap = require_previous_map();
    var fromOffsetCache = Symbol("fromOffsetCache");
    var sourceMapAvailable = Boolean(SourceMapConsumer && SourceMapGenerator);
    var pathAvailable = Boolean(resolve && isAbsolute);
    var Input = class {
      constructor(css, opts = {}) {
        if (css === null || typeof css === "undefined" || typeof css === "object" && !css.toString) {
          throw new Error(`PostCSS received ${css} instead of CSS string`);
        }
        this.css = css.toString();
        if (this.css[0] === "\uFEFF" || this.css[0] === "\uFFFE") {
          this.hasBOM = true;
          this.css = this.css.slice(1);
        } else {
          this.hasBOM = false;
        }
        if (opts.from) {
          if (!pathAvailable || /^\w+:\/\//.test(opts.from) || isAbsolute(opts.from)) {
            this.file = opts.from;
          } else {
            this.file = resolve(opts.from);
          }
        }
        if (pathAvailable && sourceMapAvailable) {
          let map = new PreviousMap(this.css, opts);
          if (map.text) {
            this.map = map;
            let file = map.consumer().file;
            if (!this.file && file)
              this.file = this.mapResolve(file);
          }
        }
        if (!this.file) {
          this.id = "<input css " + nanoid(6) + ">";
        }
        if (this.map)
          this.map.file = this.from;
      }
      fromOffset(offset) {
        let lastLine, lineToIndex;
        if (!this[fromOffsetCache]) {
          let lines = this.css.split("\n");
          lineToIndex = new Array(lines.length);
          let prevIndex = 0;
          for (let i = 0, l = lines.length; i < l; i++) {
            lineToIndex[i] = prevIndex;
            prevIndex += lines[i].length + 1;
          }
          this[fromOffsetCache] = lineToIndex;
        } else {
          lineToIndex = this[fromOffsetCache];
        }
        lastLine = lineToIndex[lineToIndex.length - 1];
        let min = 0;
        if (offset >= lastLine) {
          min = lineToIndex.length - 1;
        } else {
          let max = lineToIndex.length - 2;
          let mid;
          while (min < max) {
            mid = min + (max - min >> 1);
            if (offset < lineToIndex[mid]) {
              max = mid - 1;
            } else if (offset >= lineToIndex[mid + 1]) {
              min = mid + 1;
            } else {
              min = mid;
              break;
            }
          }
        }
        return {
          line: min + 1,
          col: offset - lineToIndex[min] + 1
        };
      }
      error(message, line, column, opts = {}) {
        let result, endLine, endColumn;
        if (line && typeof line === "object") {
          let start = line;
          let end = column;
          if (typeof start.offset === "number") {
            let pos = this.fromOffset(start.offset);
            line = pos.line;
            column = pos.col;
          } else {
            line = start.line;
            column = start.column;
          }
          if (typeof end.offset === "number") {
            let pos = this.fromOffset(end.offset);
            endLine = pos.line;
            endColumn = pos.col;
          } else {
            endLine = end.line;
            endColumn = end.column;
          }
        } else if (!column) {
          let pos = this.fromOffset(line);
          line = pos.line;
          column = pos.col;
        }
        let origin = this.origin(line, column, endLine, endColumn);
        if (origin) {
          result = new CssSyntaxError(
            message,
            origin.endLine === void 0 ? origin.line : { line: origin.line, column: origin.column },
            origin.endLine === void 0 ? origin.column : { line: origin.endLine, column: origin.endColumn },
            origin.source,
            origin.file,
            opts.plugin
          );
        } else {
          result = new CssSyntaxError(
            message,
            endLine === void 0 ? line : { line, column },
            endLine === void 0 ? column : { line: endLine, column: endColumn },
            this.css,
            this.file,
            opts.plugin
          );
        }
        result.input = { line, column, endLine, endColumn, source: this.css };
        if (this.file) {
          if (pathToFileURL) {
            result.input.url = pathToFileURL(this.file).toString();
          }
          result.input.file = this.file;
        }
        return result;
      }
      origin(line, column, endLine, endColumn) {
        if (!this.map)
          return false;
        let consumer = this.map.consumer();
        let from = consumer.originalPositionFor({ line, column });
        if (!from.source)
          return false;
        let to;
        if (typeof endLine === "number") {
          to = consumer.originalPositionFor({ line: endLine, column: endColumn });
        }
        let fromUrl;
        if (isAbsolute(from.source)) {
          fromUrl = pathToFileURL(from.source);
        } else {
          fromUrl = new URL(
            from.source,
            this.map.consumer().sourceRoot || pathToFileURL(this.map.mapFile)
          );
        }
        let result = {
          url: fromUrl.toString(),
          line: from.line,
          column: from.column,
          endLine: to && to.line,
          endColumn: to && to.column
        };
        if (fromUrl.protocol === "file:") {
          if (fileURLToPath) {
            result.file = fileURLToPath(fromUrl);
          } else {
            throw new Error(`file: protocol is not available in this PostCSS build`);
          }
        }
        let source = consumer.sourceContentFor(from.source);
        if (source)
          result.source = source;
        return result;
      }
      mapResolve(file) {
        if (/^\w+:\/\//.test(file)) {
          return file;
        }
        return resolve(this.map.consumer().sourceRoot || this.map.root || ".", file);
      }
      get from() {
        return this.file || this.id;
      }
      toJSON() {
        let json = {};
        for (let name of ["hasBOM", "css", "file", "id"]) {
          if (this[name] != null) {
            json[name] = this[name];
          }
        }
        if (this.map) {
          json.map = { ...this.map };
          if (json.map.consumerCache) {
            json.map.consumerCache = void 0;
          }
        }
        return json;
      }
    };
    module2.exports = Input;
    Input.default = Input;
    if (terminalHighlight && terminalHighlight.registerInput) {
      terminalHighlight.registerInput(Input);
    }
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/map-generator.js
var require_map_generator = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/map-generator.js"(exports, module2) {
    "use strict";
    var { SourceMapConsumer, SourceMapGenerator } = require_source_map();
    var { dirname, resolve, relative, sep } = require("path");
    var { pathToFileURL } = require("url");
    var Input = require_input();
    var sourceMapAvailable = Boolean(SourceMapConsumer && SourceMapGenerator);
    var pathAvailable = Boolean(dirname && resolve && relative && sep);
    var MapGenerator = class {
      constructor(stringify, root, opts, cssString) {
        this.stringify = stringify;
        this.mapOpts = opts.map || {};
        this.root = root;
        this.opts = opts;
        this.css = cssString;
        this.usesFileUrls = !this.mapOpts.from && this.mapOpts.absolute;
      }
      isMap() {
        if (typeof this.opts.map !== "undefined") {
          return !!this.opts.map;
        }
        return this.previous().length > 0;
      }
      previous() {
        if (!this.previousMaps) {
          this.previousMaps = [];
          if (this.root) {
            this.root.walk((node) => {
              if (node.source && node.source.input.map) {
                let map = node.source.input.map;
                if (!this.previousMaps.includes(map)) {
                  this.previousMaps.push(map);
                }
              }
            });
          } else {
            let input = new Input(this.css, this.opts);
            if (input.map)
              this.previousMaps.push(input.map);
          }
        }
        return this.previousMaps;
      }
      isInline() {
        if (typeof this.mapOpts.inline !== "undefined") {
          return this.mapOpts.inline;
        }
        let annotation = this.mapOpts.annotation;
        if (typeof annotation !== "undefined" && annotation !== true) {
          return false;
        }
        if (this.previous().length) {
          return this.previous().some((i) => i.inline);
        }
        return true;
      }
      isSourcesContent() {
        if (typeof this.mapOpts.sourcesContent !== "undefined") {
          return this.mapOpts.sourcesContent;
        }
        if (this.previous().length) {
          return this.previous().some((i) => i.withContent());
        }
        return true;
      }
      clearAnnotation() {
        if (this.mapOpts.annotation === false)
          return;
        if (this.root) {
          let node;
          for (let i = this.root.nodes.length - 1; i >= 0; i--) {
            node = this.root.nodes[i];
            if (node.type !== "comment")
              continue;
            if (node.text.indexOf("# sourceMappingURL=") === 0) {
              this.root.removeChild(i);
            }
          }
        } else if (this.css) {
          this.css = this.css.replace(/(\n)?\/\*#[\S\s]*?\*\/$/gm, "");
        }
      }
      setSourcesContent() {
        let already = {};
        if (this.root) {
          this.root.walk((node) => {
            if (node.source) {
              let from = node.source.input.from;
              if (from && !already[from]) {
                already[from] = true;
                let fromUrl = this.usesFileUrls ? this.toFileUrl(from) : this.toUrl(this.path(from));
                this.map.setSourceContent(fromUrl, node.source.input.css);
              }
            }
          });
        } else if (this.css) {
          let from = this.opts.from ? this.toUrl(this.path(this.opts.from)) : "<no source>";
          this.map.setSourceContent(from, this.css);
        }
      }
      applyPrevMaps() {
        for (let prev of this.previous()) {
          let from = this.toUrl(this.path(prev.file));
          let root = prev.root || dirname(prev.file);
          let map;
          if (this.mapOpts.sourcesContent === false) {
            map = new SourceMapConsumer(prev.text);
            if (map.sourcesContent) {
              map.sourcesContent = map.sourcesContent.map(() => null);
            }
          } else {
            map = prev.consumer();
          }
          this.map.applySourceMap(map, from, this.toUrl(this.path(root)));
        }
      }
      isAnnotation() {
        if (this.isInline()) {
          return true;
        }
        if (typeof this.mapOpts.annotation !== "undefined") {
          return this.mapOpts.annotation;
        }
        if (this.previous().length) {
          return this.previous().some((i) => i.annotation);
        }
        return true;
      }
      toBase64(str) {
        if (Buffer) {
          return Buffer.from(str).toString("base64");
        } else {
          return window.btoa(unescape(encodeURIComponent(str)));
        }
      }
      addAnnotation() {
        let content;
        if (this.isInline()) {
          content = "data:application/json;base64," + this.toBase64(this.map.toString());
        } else if (typeof this.mapOpts.annotation === "string") {
          content = this.mapOpts.annotation;
        } else if (typeof this.mapOpts.annotation === "function") {
          content = this.mapOpts.annotation(this.opts.to, this.root);
        } else {
          content = this.outputFile() + ".map";
        }
        let eol = "\n";
        if (this.css.includes("\r\n"))
          eol = "\r\n";
        this.css += eol + "/*# sourceMappingURL=" + content + " */";
      }
      outputFile() {
        if (this.opts.to) {
          return this.path(this.opts.to);
        } else if (this.opts.from) {
          return this.path(this.opts.from);
        } else {
          return "to.css";
        }
      }
      generateMap() {
        if (this.root) {
          this.generateString();
        } else if (this.previous().length === 1) {
          let prev = this.previous()[0].consumer();
          prev.file = this.outputFile();
          this.map = SourceMapGenerator.fromSourceMap(prev);
        } else {
          this.map = new SourceMapGenerator({ file: this.outputFile() });
          this.map.addMapping({
            source: this.opts.from ? this.toUrl(this.path(this.opts.from)) : "<no source>",
            generated: { line: 1, column: 0 },
            original: { line: 1, column: 0 }
          });
        }
        if (this.isSourcesContent())
          this.setSourcesContent();
        if (this.root && this.previous().length > 0)
          this.applyPrevMaps();
        if (this.isAnnotation())
          this.addAnnotation();
        if (this.isInline()) {
          return [this.css];
        } else {
          return [this.css, this.map];
        }
      }
      path(file) {
        if (file.indexOf("<") === 0)
          return file;
        if (/^\w+:\/\//.test(file))
          return file;
        if (this.mapOpts.absolute)
          return file;
        let from = this.opts.to ? dirname(this.opts.to) : ".";
        if (typeof this.mapOpts.annotation === "string") {
          from = dirname(resolve(from, this.mapOpts.annotation));
        }
        file = relative(from, file);
        return file;
      }
      toUrl(path) {
        if (sep === "\\") {
          path = path.replace(/\\/g, "/");
        }
        return encodeURI(path).replace(/[#?]/g, encodeURIComponent);
      }
      toFileUrl(path) {
        if (pathToFileURL) {
          return pathToFileURL(path).toString();
        } else {
          throw new Error(
            "`map.absolute` option is not available in this PostCSS build"
          );
        }
      }
      sourcePath(node) {
        if (this.mapOpts.from) {
          return this.toUrl(this.mapOpts.from);
        } else if (this.usesFileUrls) {
          return this.toFileUrl(node.source.input.from);
        } else {
          return this.toUrl(this.path(node.source.input.from));
        }
      }
      generateString() {
        this.css = "";
        this.map = new SourceMapGenerator({ file: this.outputFile() });
        let line = 1;
        let column = 1;
        let noSource = "<no source>";
        let mapping = {
          source: "",
          generated: { line: 0, column: 0 },
          original: { line: 0, column: 0 }
        };
        let lines, last;
        this.stringify(this.root, (str, node, type) => {
          this.css += str;
          if (node && type !== "end") {
            mapping.generated.line = line;
            mapping.generated.column = column - 1;
            if (node.source && node.source.start) {
              mapping.source = this.sourcePath(node);
              mapping.original.line = node.source.start.line;
              mapping.original.column = node.source.start.column - 1;
              this.map.addMapping(mapping);
            } else {
              mapping.source = noSource;
              mapping.original.line = 1;
              mapping.original.column = 0;
              this.map.addMapping(mapping);
            }
          }
          lines = str.match(/\n/g);
          if (lines) {
            line += lines.length;
            last = str.lastIndexOf("\n");
            column = str.length - last;
          } else {
            column += str.length;
          }
          if (node && type !== "start") {
            let p = node.parent || { raws: {} };
            let childless = node.type === "decl" || node.type === "atrule" && !node.nodes;
            if (!childless || node !== p.last || p.raws.semicolon) {
              if (node.source && node.source.end) {
                mapping.source = this.sourcePath(node);
                mapping.original.line = node.source.end.line;
                mapping.original.column = node.source.end.column - 1;
                mapping.generated.line = line;
                mapping.generated.column = column - 2;
                this.map.addMapping(mapping);
              } else {
                mapping.source = noSource;
                mapping.original.line = 1;
                mapping.original.column = 0;
                mapping.generated.line = line;
                mapping.generated.column = column - 1;
                this.map.addMapping(mapping);
              }
            }
          }
        });
      }
      generate() {
        this.clearAnnotation();
        if (pathAvailable && sourceMapAvailable && this.isMap()) {
          return this.generateMap();
        } else {
          let result = "";
          this.stringify(this.root, (i) => {
            result += i;
          });
          return [result];
        }
      }
    };
    module2.exports = MapGenerator;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/comment.js
var require_comment = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/comment.js"(exports, module2) {
    "use strict";
    var Node = require_node();
    var Comment = class extends Node {
      constructor(defaults) {
        super(defaults);
        this.type = "comment";
      }
    };
    module2.exports = Comment;
    Comment.default = Comment;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/container.js
var require_container = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/container.js"(exports, module2) {
    "use strict";
    var { isClean, my } = require_symbols();
    var Declaration = require_declaration();
    var Comment = require_comment();
    var Node = require_node();
    var parse2;
    var Rule;
    var AtRule;
    var Root;
    function cleanSource(nodes) {
      return nodes.map((i) => {
        if (i.nodes)
          i.nodes = cleanSource(i.nodes);
        delete i.source;
        return i;
      });
    }
    function markDirtyUp(node) {
      node[isClean] = false;
      if (node.proxyOf.nodes) {
        for (let i of node.proxyOf.nodes) {
          markDirtyUp(i);
        }
      }
    }
    var Container = class extends Node {
      push(child) {
        child.parent = this;
        this.proxyOf.nodes.push(child);
        return this;
      }
      each(callback) {
        if (!this.proxyOf.nodes)
          return void 0;
        let iterator = this.getIterator();
        let index2, result;
        while (this.indexes[iterator] < this.proxyOf.nodes.length) {
          index2 = this.indexes[iterator];
          result = callback(this.proxyOf.nodes[index2], index2);
          if (result === false)
            break;
          this.indexes[iterator] += 1;
        }
        delete this.indexes[iterator];
        return result;
      }
      walk(callback) {
        return this.each((child, i) => {
          let result;
          try {
            result = callback(child, i);
          } catch (e) {
            throw child.addToError(e);
          }
          if (result !== false && child.walk) {
            result = child.walk(callback);
          }
          return result;
        });
      }
      walkDecls(prop, callback) {
        if (!callback) {
          callback = prop;
          return this.walk((child, i) => {
            if (child.type === "decl") {
              return callback(child, i);
            }
          });
        }
        if (prop instanceof RegExp) {
          return this.walk((child, i) => {
            if (child.type === "decl" && prop.test(child.prop)) {
              return callback(child, i);
            }
          });
        }
        return this.walk((child, i) => {
          if (child.type === "decl" && child.prop === prop) {
            return callback(child, i);
          }
        });
      }
      walkRules(selector, callback) {
        if (!callback) {
          callback = selector;
          return this.walk((child, i) => {
            if (child.type === "rule") {
              return callback(child, i);
            }
          });
        }
        if (selector instanceof RegExp) {
          return this.walk((child, i) => {
            if (child.type === "rule" && selector.test(child.selector)) {
              return callback(child, i);
            }
          });
        }
        return this.walk((child, i) => {
          if (child.type === "rule" && child.selector === selector) {
            return callback(child, i);
          }
        });
      }
      walkAtRules(name, callback) {
        if (!callback) {
          callback = name;
          return this.walk((child, i) => {
            if (child.type === "atrule") {
              return callback(child, i);
            }
          });
        }
        if (name instanceof RegExp) {
          return this.walk((child, i) => {
            if (child.type === "atrule" && name.test(child.name)) {
              return callback(child, i);
            }
          });
        }
        return this.walk((child, i) => {
          if (child.type === "atrule" && child.name === name) {
            return callback(child, i);
          }
        });
      }
      walkComments(callback) {
        return this.walk((child, i) => {
          if (child.type === "comment") {
            return callback(child, i);
          }
        });
      }
      append(...children) {
        for (let child of children) {
          let nodes = this.normalize(child, this.last);
          for (let node of nodes)
            this.proxyOf.nodes.push(node);
        }
        this.markDirty();
        return this;
      }
      prepend(...children) {
        children = children.reverse();
        for (let child of children) {
          let nodes = this.normalize(child, this.first, "prepend").reverse();
          for (let node of nodes)
            this.proxyOf.nodes.unshift(node);
          for (let id in this.indexes) {
            this.indexes[id] = this.indexes[id] + nodes.length;
          }
        }
        this.markDirty();
        return this;
      }
      cleanRaws(keepBetween) {
        super.cleanRaws(keepBetween);
        if (this.nodes) {
          for (let node of this.nodes)
            node.cleanRaws(keepBetween);
        }
      }
      insertBefore(exist, add) {
        let existIndex = this.index(exist);
        let type = existIndex === 0 ? "prepend" : false;
        let nodes = this.normalize(add, this.proxyOf.nodes[existIndex], type).reverse();
        existIndex = this.index(exist);
        for (let node of nodes)
          this.proxyOf.nodes.splice(existIndex, 0, node);
        let index2;
        for (let id in this.indexes) {
          index2 = this.indexes[id];
          if (existIndex <= index2) {
            this.indexes[id] = index2 + nodes.length;
          }
        }
        this.markDirty();
        return this;
      }
      insertAfter(exist, add) {
        let existIndex = this.index(exist);
        let nodes = this.normalize(add, this.proxyOf.nodes[existIndex]).reverse();
        existIndex = this.index(exist);
        for (let node of nodes)
          this.proxyOf.nodes.splice(existIndex + 1, 0, node);
        let index2;
        for (let id in this.indexes) {
          index2 = this.indexes[id];
          if (existIndex < index2) {
            this.indexes[id] = index2 + nodes.length;
          }
        }
        this.markDirty();
        return this;
      }
      removeChild(child) {
        child = this.index(child);
        this.proxyOf.nodes[child].parent = void 0;
        this.proxyOf.nodes.splice(child, 1);
        let index2;
        for (let id in this.indexes) {
          index2 = this.indexes[id];
          if (index2 >= child) {
            this.indexes[id] = index2 - 1;
          }
        }
        this.markDirty();
        return this;
      }
      removeAll() {
        for (let node of this.proxyOf.nodes)
          node.parent = void 0;
        this.proxyOf.nodes = [];
        this.markDirty();
        return this;
      }
      replaceValues(pattern, opts, callback) {
        if (!callback) {
          callback = opts;
          opts = {};
        }
        this.walkDecls((decl) => {
          if (opts.props && !opts.props.includes(decl.prop))
            return;
          if (opts.fast && !decl.value.includes(opts.fast))
            return;
          decl.value = decl.value.replace(pattern, callback);
        });
        this.markDirty();
        return this;
      }
      every(condition) {
        return this.nodes.every(condition);
      }
      some(condition) {
        return this.nodes.some(condition);
      }
      index(child) {
        if (typeof child === "number")
          return child;
        if (child.proxyOf)
          child = child.proxyOf;
        return this.proxyOf.nodes.indexOf(child);
      }
      get first() {
        if (!this.proxyOf.nodes)
          return void 0;
        return this.proxyOf.nodes[0];
      }
      get last() {
        if (!this.proxyOf.nodes)
          return void 0;
        return this.proxyOf.nodes[this.proxyOf.nodes.length - 1];
      }
      normalize(nodes, sample) {
        if (typeof nodes === "string") {
          nodes = cleanSource(parse2(nodes).nodes);
        } else if (Array.isArray(nodes)) {
          nodes = nodes.slice(0);
          for (let i of nodes) {
            if (i.parent)
              i.parent.removeChild(i, "ignore");
          }
        } else if (nodes.type === "root" && this.type !== "document") {
          nodes = nodes.nodes.slice(0);
          for (let i of nodes) {
            if (i.parent)
              i.parent.removeChild(i, "ignore");
          }
        } else if (nodes.type) {
          nodes = [nodes];
        } else if (nodes.prop) {
          if (typeof nodes.value === "undefined") {
            throw new Error("Value field is missed in node creation");
          } else if (typeof nodes.value !== "string") {
            nodes.value = String(nodes.value);
          }
          nodes = [new Declaration(nodes)];
        } else if (nodes.selector) {
          nodes = [new Rule(nodes)];
        } else if (nodes.name) {
          nodes = [new AtRule(nodes)];
        } else if (nodes.text) {
          nodes = [new Comment(nodes)];
        } else {
          throw new Error("Unknown node type in node creation");
        }
        let processed = nodes.map((i) => {
          if (!i[my])
            Container.rebuild(i);
          i = i.proxyOf;
          if (i.parent)
            i.parent.removeChild(i);
          if (i[isClean])
            markDirtyUp(i);
          if (typeof i.raws.before === "undefined") {
            if (sample && typeof sample.raws.before !== "undefined") {
              i.raws.before = sample.raws.before.replace(/\S/g, "");
            }
          }
          i.parent = this.proxyOf;
          return i;
        });
        return processed;
      }
      getProxyProcessor() {
        return {
          set(node, prop, value) {
            if (node[prop] === value)
              return true;
            node[prop] = value;
            if (prop === "name" || prop === "params" || prop === "selector") {
              node.markDirty();
            }
            return true;
          },
          get(node, prop) {
            if (prop === "proxyOf") {
              return node;
            } else if (!node[prop]) {
              return node[prop];
            } else if (prop === "each" || typeof prop === "string" && prop.startsWith("walk")) {
              return (...args) => {
                return node[prop](
                  ...args.map((i) => {
                    if (typeof i === "function") {
                      return (child, index2) => i(child.toProxy(), index2);
                    } else {
                      return i;
                    }
                  })
                );
              };
            } else if (prop === "every" || prop === "some") {
              return (cb) => {
                return node[prop](
                  (child, ...other) => cb(child.toProxy(), ...other)
                );
              };
            } else if (prop === "root") {
              return () => node.root().toProxy();
            } else if (prop === "nodes") {
              return node.nodes.map((i) => i.toProxy());
            } else if (prop === "first" || prop === "last") {
              return node[prop].toProxy();
            } else {
              return node[prop];
            }
          }
        };
      }
      getIterator() {
        if (!this.lastEach)
          this.lastEach = 0;
        if (!this.indexes)
          this.indexes = {};
        this.lastEach += 1;
        let iterator = this.lastEach;
        this.indexes[iterator] = 0;
        return iterator;
      }
    };
    Container.registerParse = (dependant) => {
      parse2 = dependant;
    };
    Container.registerRule = (dependant) => {
      Rule = dependant;
    };
    Container.registerAtRule = (dependant) => {
      AtRule = dependant;
    };
    Container.registerRoot = (dependant) => {
      Root = dependant;
    };
    module2.exports = Container;
    Container.default = Container;
    Container.rebuild = (node) => {
      if (node.type === "atrule") {
        Object.setPrototypeOf(node, AtRule.prototype);
      } else if (node.type === "rule") {
        Object.setPrototypeOf(node, Rule.prototype);
      } else if (node.type === "decl") {
        Object.setPrototypeOf(node, Declaration.prototype);
      } else if (node.type === "comment") {
        Object.setPrototypeOf(node, Comment.prototype);
      } else if (node.type === "root") {
        Object.setPrototypeOf(node, Root.prototype);
      }
      node[my] = true;
      if (node.nodes) {
        node.nodes.forEach((child) => {
          Container.rebuild(child);
        });
      }
    };
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/document.js
var require_document = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/document.js"(exports, module2) {
    "use strict";
    var Container = require_container();
    var LazyResult;
    var Processor;
    var Document = class extends Container {
      constructor(defaults) {
        super({ type: "document", ...defaults });
        if (!this.nodes) {
          this.nodes = [];
        }
      }
      toResult(opts = {}) {
        let lazy = new LazyResult(new Processor(), this, opts);
        return lazy.stringify();
      }
    };
    Document.registerLazyResult = (dependant) => {
      LazyResult = dependant;
    };
    Document.registerProcessor = (dependant) => {
      Processor = dependant;
    };
    module2.exports = Document;
    Document.default = Document;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/warn-once.js
var require_warn_once = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/warn-once.js"(exports, module2) {
    "use strict";
    var printed = {};
    module2.exports = function warnOnce(message) {
      if (printed[message])
        return;
      printed[message] = true;
      if (typeof console !== "undefined" && console.warn) {
        console.warn(message);
      }
    };
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/warning.js
var require_warning = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/warning.js"(exports, module2) {
    "use strict";
    var Warning = class {
      constructor(text2, opts = {}) {
        this.type = "warning";
        this.text = text2;
        if (opts.node && opts.node.source) {
          let range = opts.node.rangeBy(opts);
          this.line = range.start.line;
          this.column = range.start.column;
          this.endLine = range.end.line;
          this.endColumn = range.end.column;
        }
        for (let opt in opts)
          this[opt] = opts[opt];
      }
      toString() {
        if (this.node) {
          return this.node.error(this.text, {
            plugin: this.plugin,
            index: this.index,
            word: this.word
          }).message;
        }
        if (this.plugin) {
          return this.plugin + ": " + this.text;
        }
        return this.text;
      }
    };
    module2.exports = Warning;
    Warning.default = Warning;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/result.js
var require_result = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/result.js"(exports, module2) {
    "use strict";
    var Warning = require_warning();
    var Result = class {
      constructor(processor, root, opts) {
        this.processor = processor;
        this.messages = [];
        this.root = root;
        this.opts = opts;
        this.css = void 0;
        this.map = void 0;
      }
      toString() {
        return this.css;
      }
      warn(text2, opts = {}) {
        if (!opts.plugin) {
          if (this.lastPlugin && this.lastPlugin.postcssPlugin) {
            opts.plugin = this.lastPlugin.postcssPlugin;
          }
        }
        let warning = new Warning(text2, opts);
        this.messages.push(warning);
        return warning;
      }
      warnings() {
        return this.messages.filter((i) => i.type === "warning");
      }
      get content() {
        return this.css;
      }
    };
    module2.exports = Result;
    Result.default = Result;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/at-rule.js
var require_at_rule = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/at-rule.js"(exports, module2) {
    "use strict";
    var Container = require_container();
    var AtRule = class extends Container {
      constructor(defaults) {
        super(defaults);
        this.type = "atrule";
      }
      append(...children) {
        if (!this.proxyOf.nodes)
          this.nodes = [];
        return super.append(...children);
      }
      prepend(...children) {
        if (!this.proxyOf.nodes)
          this.nodes = [];
        return super.prepend(...children);
      }
    };
    module2.exports = AtRule;
    AtRule.default = AtRule;
    Container.registerAtRule(AtRule);
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/root.js
var require_root = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/root.js"(exports, module2) {
    "use strict";
    var Container = require_container();
    var LazyResult;
    var Processor;
    var Root = class extends Container {
      constructor(defaults) {
        super(defaults);
        this.type = "root";
        if (!this.nodes)
          this.nodes = [];
      }
      removeChild(child, ignore) {
        let index2 = this.index(child);
        if (!ignore && index2 === 0 && this.nodes.length > 1) {
          this.nodes[1].raws.before = this.nodes[index2].raws.before;
        }
        return super.removeChild(child);
      }
      normalize(child, sample, type) {
        let nodes = super.normalize(child);
        if (sample) {
          if (type === "prepend") {
            if (this.nodes.length > 1) {
              sample.raws.before = this.nodes[1].raws.before;
            } else {
              delete sample.raws.before;
            }
          } else if (this.first !== sample) {
            for (let node of nodes) {
              node.raws.before = sample.raws.before;
            }
          }
        }
        return nodes;
      }
      toResult(opts = {}) {
        let lazy = new LazyResult(new Processor(), this, opts);
        return lazy.stringify();
      }
    };
    Root.registerLazyResult = (dependant) => {
      LazyResult = dependant;
    };
    Root.registerProcessor = (dependant) => {
      Processor = dependant;
    };
    module2.exports = Root;
    Root.default = Root;
    Container.registerRoot(Root);
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/list.js
var require_list = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/list.js"(exports, module2) {
    "use strict";
    var list = {
      split(string, separators, last) {
        let array = [];
        let current = "";
        let split = false;
        let func = 0;
        let inQuote = false;
        let prevQuote = "";
        let escape = false;
        for (let letter of string) {
          if (escape) {
            escape = false;
          } else if (letter === "\\") {
            escape = true;
          } else if (inQuote) {
            if (letter === prevQuote) {
              inQuote = false;
            }
          } else if (letter === '"' || letter === "'") {
            inQuote = true;
            prevQuote = letter;
          } else if (letter === "(") {
            func += 1;
          } else if (letter === ")") {
            if (func > 0)
              func -= 1;
          } else if (func === 0) {
            if (separators.includes(letter))
              split = true;
          }
          if (split) {
            if (current !== "")
              array.push(current.trim());
            current = "";
            split = false;
          } else {
            current += letter;
          }
        }
        if (last || current !== "")
          array.push(current.trim());
        return array;
      },
      space(string) {
        let spaces = [" ", "\n", "	"];
        return list.split(string, spaces);
      },
      comma(string) {
        return list.split(string, [","], true);
      }
    };
    module2.exports = list;
    list.default = list;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/rule.js
var require_rule = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/rule.js"(exports, module2) {
    "use strict";
    var Container = require_container();
    var list = require_list();
    var Rule = class extends Container {
      constructor(defaults) {
        super(defaults);
        this.type = "rule";
        if (!this.nodes)
          this.nodes = [];
      }
      get selectors() {
        return list.comma(this.selector);
      }
      set selectors(values) {
        let match = this.selector ? this.selector.match(/,\s*/) : null;
        let sep = match ? match[0] : "," + this.raw("between", "beforeOpen");
        this.selector = values.join(sep);
      }
    };
    module2.exports = Rule;
    Rule.default = Rule;
    Container.registerRule(Rule);
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/parser.js
var require_parser = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/parser.js"(exports, module2) {
    "use strict";
    var Declaration = require_declaration();
    var tokenizer = require_tokenize();
    var Comment = require_comment();
    var AtRule = require_at_rule();
    var Root = require_root();
    var Rule = require_rule();
    var SAFE_COMMENT_NEIGHBOR = {
      empty: true,
      space: true
    };
    function findLastWithPosition(tokens) {
      for (let i = tokens.length - 1; i >= 0; i--) {
        let token = tokens[i];
        let pos = token[3] || token[2];
        if (pos)
          return pos;
      }
    }
    var Parser = class {
      constructor(input) {
        this.input = input;
        this.root = new Root();
        this.current = this.root;
        this.spaces = "";
        this.semicolon = false;
        this.customProperty = false;
        this.createTokenizer();
        this.root.source = { input, start: { offset: 0, line: 1, column: 1 } };
      }
      createTokenizer() {
        this.tokenizer = tokenizer(this.input);
      }
      parse() {
        let token;
        while (!this.tokenizer.endOfFile()) {
          token = this.tokenizer.nextToken();
          switch (token[0]) {
            case "space":
              this.spaces += token[1];
              break;
            case ";":
              this.freeSemicolon(token);
              break;
            case "}":
              this.end(token);
              break;
            case "comment":
              this.comment(token);
              break;
            case "at-word":
              this.atrule(token);
              break;
            case "{":
              this.emptyRule(token);
              break;
            default:
              this.other(token);
              break;
          }
        }
        this.endFile();
      }
      comment(token) {
        let node = new Comment();
        this.init(node, token[2]);
        node.source.end = this.getPosition(token[3] || token[2]);
        let text2 = token[1].slice(2, -2);
        if (/^\s*$/.test(text2)) {
          node.text = "";
          node.raws.left = text2;
          node.raws.right = "";
        } else {
          let match = text2.match(/^(\s*)([^]*\S)(\s*)$/);
          node.text = match[2];
          node.raws.left = match[1];
          node.raws.right = match[3];
        }
      }
      emptyRule(token) {
        let node = new Rule();
        this.init(node, token[2]);
        node.selector = "";
        node.raws.between = "";
        this.current = node;
      }
      other(start) {
        let end = false;
        let type = null;
        let colon = false;
        let bracket = null;
        let brackets = [];
        let customProperty = start[1].startsWith("--");
        let tokens = [];
        let token = start;
        while (token) {
          type = token[0];
          tokens.push(token);
          if (type === "(" || type === "[") {
            if (!bracket)
              bracket = token;
            brackets.push(type === "(" ? ")" : "]");
          } else if (customProperty && colon && type === "{") {
            if (!bracket)
              bracket = token;
            brackets.push("}");
          } else if (brackets.length === 0) {
            if (type === ";") {
              if (colon) {
                this.decl(tokens, customProperty);
                return;
              } else {
                break;
              }
            } else if (type === "{") {
              this.rule(tokens);
              return;
            } else if (type === "}") {
              this.tokenizer.back(tokens.pop());
              end = true;
              break;
            } else if (type === ":") {
              colon = true;
            }
          } else if (type === brackets[brackets.length - 1]) {
            brackets.pop();
            if (brackets.length === 0)
              bracket = null;
          }
          token = this.tokenizer.nextToken();
        }
        if (this.tokenizer.endOfFile())
          end = true;
        if (brackets.length > 0)
          this.unclosedBracket(bracket);
        if (end && colon) {
          if (!customProperty) {
            while (tokens.length) {
              token = tokens[tokens.length - 1][0];
              if (token !== "space" && token !== "comment")
                break;
              this.tokenizer.back(tokens.pop());
            }
          }
          this.decl(tokens, customProperty);
        } else {
          this.unknownWord(tokens);
        }
      }
      rule(tokens) {
        tokens.pop();
        let node = new Rule();
        this.init(node, tokens[0][2]);
        node.raws.between = this.spacesAndCommentsFromEnd(tokens);
        this.raw(node, "selector", tokens);
        this.current = node;
      }
      decl(tokens, customProperty) {
        let node = new Declaration();
        this.init(node, tokens[0][2]);
        let last = tokens[tokens.length - 1];
        if (last[0] === ";") {
          this.semicolon = true;
          tokens.pop();
        }
        node.source.end = this.getPosition(
          last[3] || last[2] || findLastWithPosition(tokens)
        );
        while (tokens[0][0] !== "word") {
          if (tokens.length === 1)
            this.unknownWord(tokens);
          node.raws.before += tokens.shift()[1];
        }
        node.source.start = this.getPosition(tokens[0][2]);
        node.prop = "";
        while (tokens.length) {
          let type = tokens[0][0];
          if (type === ":" || type === "space" || type === "comment") {
            break;
          }
          node.prop += tokens.shift()[1];
        }
        node.raws.between = "";
        let token;
        while (tokens.length) {
          token = tokens.shift();
          if (token[0] === ":") {
            node.raws.between += token[1];
            break;
          } else {
            if (token[0] === "word" && /\w/.test(token[1])) {
              this.unknownWord([token]);
            }
            node.raws.between += token[1];
          }
        }
        if (node.prop[0] === "_" || node.prop[0] === "*") {
          node.raws.before += node.prop[0];
          node.prop = node.prop.slice(1);
        }
        let firstSpaces = [];
        let next;
        while (tokens.length) {
          next = tokens[0][0];
          if (next !== "space" && next !== "comment")
            break;
          firstSpaces.push(tokens.shift());
        }
        this.precheckMissedSemicolon(tokens);
        for (let i = tokens.length - 1; i >= 0; i--) {
          token = tokens[i];
          if (token[1].toLowerCase() === "!important") {
            node.important = true;
            let string = this.stringFrom(tokens, i);
            string = this.spacesFromEnd(tokens) + string;
            if (string !== " !important")
              node.raws.important = string;
            break;
          } else if (token[1].toLowerCase() === "important") {
            let cache = tokens.slice(0);
            let str = "";
            for (let j = i; j > 0; j--) {
              let type = cache[j][0];
              if (str.trim().indexOf("!") === 0 && type !== "space") {
                break;
              }
              str = cache.pop()[1] + str;
            }
            if (str.trim().indexOf("!") === 0) {
              node.important = true;
              node.raws.important = str;
              tokens = cache;
            }
          }
          if (token[0] !== "space" && token[0] !== "comment") {
            break;
          }
        }
        let hasWord = tokens.some((i) => i[0] !== "space" && i[0] !== "comment");
        if (hasWord) {
          node.raws.between += firstSpaces.map((i) => i[1]).join("");
          firstSpaces = [];
        }
        this.raw(node, "value", firstSpaces.concat(tokens), customProperty);
        if (node.value.includes(":") && !customProperty) {
          this.checkMissedSemicolon(tokens);
        }
      }
      atrule(token) {
        let node = new AtRule();
        node.name = token[1].slice(1);
        if (node.name === "") {
          this.unnamedAtrule(node, token);
        }
        this.init(node, token[2]);
        let type;
        let prev;
        let shift;
        let last = false;
        let open = false;
        let params = [];
        let brackets = [];
        while (!this.tokenizer.endOfFile()) {
          token = this.tokenizer.nextToken();
          type = token[0];
          if (type === "(" || type === "[") {
            brackets.push(type === "(" ? ")" : "]");
          } else if (type === "{" && brackets.length > 0) {
            brackets.push("}");
          } else if (type === brackets[brackets.length - 1]) {
            brackets.pop();
          }
          if (brackets.length === 0) {
            if (type === ";") {
              node.source.end = this.getPosition(token[2]);
              this.semicolon = true;
              break;
            } else if (type === "{") {
              open = true;
              break;
            } else if (type === "}") {
              if (params.length > 0) {
                shift = params.length - 1;
                prev = params[shift];
                while (prev && prev[0] === "space") {
                  prev = params[--shift];
                }
                if (prev) {
                  node.source.end = this.getPosition(prev[3] || prev[2]);
                }
              }
              this.end(token);
              break;
            } else {
              params.push(token);
            }
          } else {
            params.push(token);
          }
          if (this.tokenizer.endOfFile()) {
            last = true;
            break;
          }
        }
        node.raws.between = this.spacesAndCommentsFromEnd(params);
        if (params.length) {
          node.raws.afterName = this.spacesAndCommentsFromStart(params);
          this.raw(node, "params", params);
          if (last) {
            token = params[params.length - 1];
            node.source.end = this.getPosition(token[3] || token[2]);
            this.spaces = node.raws.between;
            node.raws.between = "";
          }
        } else {
          node.raws.afterName = "";
          node.params = "";
        }
        if (open) {
          node.nodes = [];
          this.current = node;
        }
      }
      end(token) {
        if (this.current.nodes && this.current.nodes.length) {
          this.current.raws.semicolon = this.semicolon;
        }
        this.semicolon = false;
        this.current.raws.after = (this.current.raws.after || "") + this.spaces;
        this.spaces = "";
        if (this.current.parent) {
          this.current.source.end = this.getPosition(token[2]);
          this.current = this.current.parent;
        } else {
          this.unexpectedClose(token);
        }
      }
      endFile() {
        if (this.current.parent)
          this.unclosedBlock();
        if (this.current.nodes && this.current.nodes.length) {
          this.current.raws.semicolon = this.semicolon;
        }
        this.current.raws.after = (this.current.raws.after || "") + this.spaces;
      }
      freeSemicolon(token) {
        this.spaces += token[1];
        if (this.current.nodes) {
          let prev = this.current.nodes[this.current.nodes.length - 1];
          if (prev && prev.type === "rule" && !prev.raws.ownSemicolon) {
            prev.raws.ownSemicolon = this.spaces;
            this.spaces = "";
          }
        }
      }
      // Helpers
      getPosition(offset) {
        let pos = this.input.fromOffset(offset);
        return {
          offset,
          line: pos.line,
          column: pos.col
        };
      }
      init(node, offset) {
        this.current.push(node);
        node.source = {
          start: this.getPosition(offset),
          input: this.input
        };
        node.raws.before = this.spaces;
        this.spaces = "";
        if (node.type !== "comment")
          this.semicolon = false;
      }
      raw(node, prop, tokens, customProperty) {
        let token, type;
        let length = tokens.length;
        let value = "";
        let clean = true;
        let next, prev;
        for (let i = 0; i < length; i += 1) {
          token = tokens[i];
          type = token[0];
          if (type === "space" && i === length - 1 && !customProperty) {
            clean = false;
          } else if (type === "comment") {
            prev = tokens[i - 1] ? tokens[i - 1][0] : "empty";
            next = tokens[i + 1] ? tokens[i + 1][0] : "empty";
            if (!SAFE_COMMENT_NEIGHBOR[prev] && !SAFE_COMMENT_NEIGHBOR[next]) {
              if (value.slice(-1) === ",") {
                clean = false;
              } else {
                value += token[1];
              }
            } else {
              clean = false;
            }
          } else {
            value += token[1];
          }
        }
        if (!clean) {
          let raw = tokens.reduce((all, i) => all + i[1], "");
          node.raws[prop] = { value, raw };
        }
        node[prop] = value;
      }
      spacesAndCommentsFromEnd(tokens) {
        let lastTokenType;
        let spaces = "";
        while (tokens.length) {
          lastTokenType = tokens[tokens.length - 1][0];
          if (lastTokenType !== "space" && lastTokenType !== "comment")
            break;
          spaces = tokens.pop()[1] + spaces;
        }
        return spaces;
      }
      spacesAndCommentsFromStart(tokens) {
        let next;
        let spaces = "";
        while (tokens.length) {
          next = tokens[0][0];
          if (next !== "space" && next !== "comment")
            break;
          spaces += tokens.shift()[1];
        }
        return spaces;
      }
      spacesFromEnd(tokens) {
        let lastTokenType;
        let spaces = "";
        while (tokens.length) {
          lastTokenType = tokens[tokens.length - 1][0];
          if (lastTokenType !== "space")
            break;
          spaces = tokens.pop()[1] + spaces;
        }
        return spaces;
      }
      stringFrom(tokens, from) {
        let result = "";
        for (let i = from; i < tokens.length; i++) {
          result += tokens[i][1];
        }
        tokens.splice(from, tokens.length - from);
        return result;
      }
      colon(tokens) {
        let brackets = 0;
        let token, type, prev;
        for (let [i, element] of tokens.entries()) {
          token = element;
          type = token[0];
          if (type === "(") {
            brackets += 1;
          }
          if (type === ")") {
            brackets -= 1;
          }
          if (brackets === 0 && type === ":") {
            if (!prev) {
              this.doubleColon(token);
            } else if (prev[0] === "word" && prev[1] === "progid") {
              continue;
            } else {
              return i;
            }
          }
          prev = token;
        }
        return false;
      }
      // Errors
      unclosedBracket(bracket) {
        throw this.input.error(
          "Unclosed bracket",
          { offset: bracket[2] },
          { offset: bracket[2] + 1 }
        );
      }
      unknownWord(tokens) {
        throw this.input.error(
          "Unknown word",
          { offset: tokens[0][2] },
          { offset: tokens[0][2] + tokens[0][1].length }
        );
      }
      unexpectedClose(token) {
        throw this.input.error(
          "Unexpected }",
          { offset: token[2] },
          { offset: token[2] + 1 }
        );
      }
      unclosedBlock() {
        let pos = this.current.source.start;
        throw this.input.error("Unclosed block", pos.line, pos.column);
      }
      doubleColon(token) {
        throw this.input.error(
          "Double colon",
          { offset: token[2] },
          { offset: token[2] + token[1].length }
        );
      }
      unnamedAtrule(node, token) {
        throw this.input.error(
          "At-rule without name",
          { offset: token[2] },
          { offset: token[2] + token[1].length }
        );
      }
      precheckMissedSemicolon() {
      }
      checkMissedSemicolon(tokens) {
        let colon = this.colon(tokens);
        if (colon === false)
          return;
        let founded = 0;
        let token;
        for (let j = colon - 1; j >= 0; j--) {
          token = tokens[j];
          if (token[0] !== "space") {
            founded += 1;
            if (founded === 2)
              break;
          }
        }
        throw this.input.error(
          "Missed semicolon",
          token[0] === "word" ? token[3] + 1 : token[2]
        );
      }
    };
    module2.exports = Parser;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/parse.js
var require_parse = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/parse.js"(exports, module2) {
    "use strict";
    var Container = require_container();
    var Parser = require_parser();
    var Input = require_input();
    function parse2(css, opts) {
      let input = new Input(css, opts);
      let parser2 = new Parser(input);
      try {
        parser2.parse();
      } catch (e) {
        if (process.env.NODE_ENV !== "production") {
          if (e.name === "CssSyntaxError" && opts && opts.from) {
            if (/\.scss$/i.test(opts.from)) {
              e.message += "\nYou tried to parse SCSS with the standard CSS parser; try again with the postcss-scss parser";
            } else if (/\.sass/i.test(opts.from)) {
              e.message += "\nYou tried to parse Sass with the standard CSS parser; try again with the postcss-sass parser";
            } else if (/\.less$/i.test(opts.from)) {
              e.message += "\nYou tried to parse Less with the standard CSS parser; try again with the postcss-less parser";
            }
          }
        }
        throw e;
      }
      return parser2.root;
    }
    module2.exports = parse2;
    parse2.default = parse2;
    Container.registerParse(parse2);
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/lazy-result.js
var require_lazy_result = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/lazy-result.js"(exports, module2) {
    "use strict";
    var { isClean, my } = require_symbols();
    var MapGenerator = require_map_generator();
    var stringify = require_stringify();
    var Container = require_container();
    var Document = require_document();
    var warnOnce = require_warn_once();
    var Result = require_result();
    var parse2 = require_parse();
    var Root = require_root();
    var TYPE_TO_CLASS_NAME = {
      document: "Document",
      root: "Root",
      atrule: "AtRule",
      rule: "Rule",
      decl: "Declaration",
      comment: "Comment"
    };
    var PLUGIN_PROPS = {
      postcssPlugin: true,
      prepare: true,
      Once: true,
      Document: true,
      Root: true,
      Declaration: true,
      Rule: true,
      AtRule: true,
      Comment: true,
      DeclarationExit: true,
      RuleExit: true,
      AtRuleExit: true,
      CommentExit: true,
      RootExit: true,
      DocumentExit: true,
      OnceExit: true
    };
    var NOT_VISITORS = {
      postcssPlugin: true,
      prepare: true,
      Once: true
    };
    var CHILDREN = 0;
    function isPromise(obj) {
      return typeof obj === "object" && typeof obj.then === "function";
    }
    function getEvents(node) {
      let key = false;
      let type = TYPE_TO_CLASS_NAME[node.type];
      if (node.type === "decl") {
        key = node.prop.toLowerCase();
      } else if (node.type === "atrule") {
        key = node.name.toLowerCase();
      }
      if (key && node.append) {
        return [
          type,
          type + "-" + key,
          CHILDREN,
          type + "Exit",
          type + "Exit-" + key
        ];
      } else if (key) {
        return [type, type + "-" + key, type + "Exit", type + "Exit-" + key];
      } else if (node.append) {
        return [type, CHILDREN, type + "Exit"];
      } else {
        return [type, type + "Exit"];
      }
    }
    function toStack(node) {
      let events;
      if (node.type === "document") {
        events = ["Document", CHILDREN, "DocumentExit"];
      } else if (node.type === "root") {
        events = ["Root", CHILDREN, "RootExit"];
      } else {
        events = getEvents(node);
      }
      return {
        node,
        events,
        eventIndex: 0,
        visitors: [],
        visitorIndex: 0,
        iterator: 0
      };
    }
    function cleanMarks(node) {
      node[isClean] = false;
      if (node.nodes)
        node.nodes.forEach((i) => cleanMarks(i));
      return node;
    }
    var postcss = {};
    var LazyResult = class {
      constructor(processor, css, opts) {
        this.stringified = false;
        this.processed = false;
        let root;
        if (typeof css === "object" && css !== null && (css.type === "root" || css.type === "document")) {
          root = cleanMarks(css);
        } else if (css instanceof LazyResult || css instanceof Result) {
          root = cleanMarks(css.root);
          if (css.map) {
            if (typeof opts.map === "undefined")
              opts.map = {};
            if (!opts.map.inline)
              opts.map.inline = false;
            opts.map.prev = css.map;
          }
        } else {
          let parser2 = parse2;
          if (opts.syntax)
            parser2 = opts.syntax.parse;
          if (opts.parser)
            parser2 = opts.parser;
          if (parser2.parse)
            parser2 = parser2.parse;
          try {
            root = parser2(css, opts);
          } catch (error) {
            this.processed = true;
            this.error = error;
          }
          if (root && !root[my]) {
            Container.rebuild(root);
          }
        }
        this.result = new Result(processor, root, opts);
        this.helpers = { ...postcss, result: this.result, postcss };
        this.plugins = this.processor.plugins.map((plugin3) => {
          if (typeof plugin3 === "object" && plugin3.prepare) {
            return { ...plugin3, ...plugin3.prepare(this.result) };
          } else {
            return plugin3;
          }
        });
      }
      get [Symbol.toStringTag]() {
        return "LazyResult";
      }
      get processor() {
        return this.result.processor;
      }
      get opts() {
        return this.result.opts;
      }
      get css() {
        return this.stringify().css;
      }
      get content() {
        return this.stringify().content;
      }
      get map() {
        return this.stringify().map;
      }
      get root() {
        return this.sync().root;
      }
      get messages() {
        return this.sync().messages;
      }
      warnings() {
        return this.sync().warnings();
      }
      toString() {
        return this.css;
      }
      then(onFulfilled, onRejected) {
        if (process.env.NODE_ENV !== "production") {
          if (!("from" in this.opts)) {
            warnOnce(
              "Without `from` option PostCSS could generate wrong source map and will not find Browserslist config. Set it to CSS file path or to `undefined` to prevent this warning."
            );
          }
        }
        return this.async().then(onFulfilled, onRejected);
      }
      catch(onRejected) {
        return this.async().catch(onRejected);
      }
      finally(onFinally) {
        return this.async().then(onFinally, onFinally);
      }
      async() {
        if (this.error)
          return Promise.reject(this.error);
        if (this.processed)
          return Promise.resolve(this.result);
        if (!this.processing) {
          this.processing = this.runAsync();
        }
        return this.processing;
      }
      sync() {
        if (this.error)
          throw this.error;
        if (this.processed)
          return this.result;
        this.processed = true;
        if (this.processing) {
          throw this.getAsyncError();
        }
        for (let plugin3 of this.plugins) {
          let promise = this.runOnRoot(plugin3);
          if (isPromise(promise)) {
            throw this.getAsyncError();
          }
        }
        this.prepareVisitors();
        if (this.hasListener) {
          let root = this.result.root;
          while (!root[isClean]) {
            root[isClean] = true;
            this.walkSync(root);
          }
          if (this.listeners.OnceExit) {
            if (root.type === "document") {
              for (let subRoot of root.nodes) {
                this.visitSync(this.listeners.OnceExit, subRoot);
              }
            } else {
              this.visitSync(this.listeners.OnceExit, root);
            }
          }
        }
        return this.result;
      }
      stringify() {
        if (this.error)
          throw this.error;
        if (this.stringified)
          return this.result;
        this.stringified = true;
        this.sync();
        let opts = this.result.opts;
        let str = stringify;
        if (opts.syntax)
          str = opts.syntax.stringify;
        if (opts.stringifier)
          str = opts.stringifier;
        if (str.stringify)
          str = str.stringify;
        let map = new MapGenerator(str, this.result.root, this.result.opts);
        let data = map.generate();
        this.result.css = data[0];
        this.result.map = data[1];
        return this.result;
      }
      walkSync(node) {
        node[isClean] = true;
        let events = getEvents(node);
        for (let event of events) {
          if (event === CHILDREN) {
            if (node.nodes) {
              node.each((child) => {
                if (!child[isClean])
                  this.walkSync(child);
              });
            }
          } else {
            let visitors = this.listeners[event];
            if (visitors) {
              if (this.visitSync(visitors, node.toProxy()))
                return;
            }
          }
        }
      }
      visitSync(visitors, node) {
        for (let [plugin3, visitor] of visitors) {
          this.result.lastPlugin = plugin3;
          let promise;
          try {
            promise = visitor(node, this.helpers);
          } catch (e) {
            throw this.handleError(e, node.proxyOf);
          }
          if (node.type !== "root" && node.type !== "document" && !node.parent) {
            return true;
          }
          if (isPromise(promise)) {
            throw this.getAsyncError();
          }
        }
      }
      runOnRoot(plugin3) {
        this.result.lastPlugin = plugin3;
        try {
          if (typeof plugin3 === "object" && plugin3.Once) {
            if (this.result.root.type === "document") {
              let roots = this.result.root.nodes.map(
                (root) => plugin3.Once(root, this.helpers)
              );
              if (isPromise(roots[0])) {
                return Promise.all(roots);
              }
              return roots;
            }
            return plugin3.Once(this.result.root, this.helpers);
          } else if (typeof plugin3 === "function") {
            return plugin3(this.result.root, this.result);
          }
        } catch (error) {
          throw this.handleError(error);
        }
      }
      getAsyncError() {
        throw new Error("Use process(css).then(cb) to work with async plugins");
      }
      handleError(error, node) {
        let plugin3 = this.result.lastPlugin;
        try {
          if (node)
            node.addToError(error);
          this.error = error;
          if (error.name === "CssSyntaxError" && !error.plugin) {
            error.plugin = plugin3.postcssPlugin;
            error.setMessage();
          } else if (plugin3.postcssVersion) {
            if (process.env.NODE_ENV !== "production") {
              let pluginName = plugin3.postcssPlugin;
              let pluginVer = plugin3.postcssVersion;
              let runtimeVer = this.result.processor.version;
              let a = pluginVer.split(".");
              let b = runtimeVer.split(".");
              if (a[0] !== b[0] || parseInt(a[1]) > parseInt(b[1])) {
                console.error(
                  "Unknown error from PostCSS plugin. Your current PostCSS version is " + runtimeVer + ", but " + pluginName + " uses " + pluginVer + ". Perhaps this is the source of the error below."
                );
              }
            }
          }
        } catch (err) {
          if (console && console.error)
            console.error(err);
        }
        return error;
      }
      async runAsync() {
        this.plugin = 0;
        for (let i = 0; i < this.plugins.length; i++) {
          let plugin3 = this.plugins[i];
          let promise = this.runOnRoot(plugin3);
          if (isPromise(promise)) {
            try {
              await promise;
            } catch (error) {
              throw this.handleError(error);
            }
          }
        }
        this.prepareVisitors();
        if (this.hasListener) {
          let root = this.result.root;
          while (!root[isClean]) {
            root[isClean] = true;
            let stack = [toStack(root)];
            while (stack.length > 0) {
              let promise = this.visitTick(stack);
              if (isPromise(promise)) {
                try {
                  await promise;
                } catch (e) {
                  let node = stack[stack.length - 1].node;
                  throw this.handleError(e, node);
                }
              }
            }
          }
          if (this.listeners.OnceExit) {
            for (let [plugin3, visitor] of this.listeners.OnceExit) {
              this.result.lastPlugin = plugin3;
              try {
                if (root.type === "document") {
                  let roots = root.nodes.map(
                    (subRoot) => visitor(subRoot, this.helpers)
                  );
                  await Promise.all(roots);
                } else {
                  await visitor(root, this.helpers);
                }
              } catch (e) {
                throw this.handleError(e);
              }
            }
          }
        }
        this.processed = true;
        return this.stringify();
      }
      prepareVisitors() {
        this.listeners = {};
        let add = (plugin3, type, cb) => {
          if (!this.listeners[type])
            this.listeners[type] = [];
          this.listeners[type].push([plugin3, cb]);
        };
        for (let plugin3 of this.plugins) {
          if (typeof plugin3 === "object") {
            for (let event in plugin3) {
              if (!PLUGIN_PROPS[event] && /^[A-Z]/.test(event)) {
                throw new Error(
                  `Unknown event ${event} in ${plugin3.postcssPlugin}. Try to update PostCSS (${this.processor.version} now).`
                );
              }
              if (!NOT_VISITORS[event]) {
                if (typeof plugin3[event] === "object") {
                  for (let filter in plugin3[event]) {
                    if (filter === "*") {
                      add(plugin3, event, plugin3[event][filter]);
                    } else {
                      add(
                        plugin3,
                        event + "-" + filter.toLowerCase(),
                        plugin3[event][filter]
                      );
                    }
                  }
                } else if (typeof plugin3[event] === "function") {
                  add(plugin3, event, plugin3[event]);
                }
              }
            }
          }
        }
        this.hasListener = Object.keys(this.listeners).length > 0;
      }
      visitTick(stack) {
        let visit = stack[stack.length - 1];
        let { node, visitors } = visit;
        if (node.type !== "root" && node.type !== "document" && !node.parent) {
          stack.pop();
          return;
        }
        if (visitors.length > 0 && visit.visitorIndex < visitors.length) {
          let [plugin3, visitor] = visitors[visit.visitorIndex];
          visit.visitorIndex += 1;
          if (visit.visitorIndex === visitors.length) {
            visit.visitors = [];
            visit.visitorIndex = 0;
          }
          this.result.lastPlugin = plugin3;
          try {
            return visitor(node.toProxy(), this.helpers);
          } catch (e) {
            throw this.handleError(e, node);
          }
        }
        if (visit.iterator !== 0) {
          let iterator = visit.iterator;
          let child;
          while (child = node.nodes[node.indexes[iterator]]) {
            node.indexes[iterator] += 1;
            if (!child[isClean]) {
              child[isClean] = true;
              stack.push(toStack(child));
              return;
            }
          }
          visit.iterator = 0;
          delete node.indexes[iterator];
        }
        let events = visit.events;
        while (visit.eventIndex < events.length) {
          let event = events[visit.eventIndex];
          visit.eventIndex += 1;
          if (event === CHILDREN) {
            if (node.nodes && node.nodes.length) {
              node[isClean] = true;
              visit.iterator = node.getIterator();
            }
            return;
          } else if (this.listeners[event]) {
            visit.visitors = this.listeners[event];
            return;
          }
        }
        stack.pop();
      }
    };
    LazyResult.registerPostcss = (dependant) => {
      postcss = dependant;
    };
    module2.exports = LazyResult;
    LazyResult.default = LazyResult;
    Root.registerLazyResult(LazyResult);
    Document.registerLazyResult(LazyResult);
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/no-work-result.js
var require_no_work_result = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/no-work-result.js"(exports, module2) {
    "use strict";
    var MapGenerator = require_map_generator();
    var stringify = require_stringify();
    var warnOnce = require_warn_once();
    var parse2 = require_parse();
    var Result = require_result();
    var NoWorkResult = class {
      constructor(processor, css, opts) {
        css = css.toString();
        this.stringified = false;
        this._processor = processor;
        this._css = css;
        this._opts = opts;
        this._map = void 0;
        let root;
        let str = stringify;
        this.result = new Result(this._processor, root, this._opts);
        this.result.css = css;
        let self = this;
        Object.defineProperty(this.result, "root", {
          get() {
            return self.root;
          }
        });
        let map = new MapGenerator(str, root, this._opts, css);
        if (map.isMap()) {
          let [generatedCSS, generatedMap] = map.generate();
          if (generatedCSS) {
            this.result.css = generatedCSS;
          }
          if (generatedMap) {
            this.result.map = generatedMap;
          }
        }
      }
      get [Symbol.toStringTag]() {
        return "NoWorkResult";
      }
      get processor() {
        return this.result.processor;
      }
      get opts() {
        return this.result.opts;
      }
      get css() {
        return this.result.css;
      }
      get content() {
        return this.result.css;
      }
      get map() {
        return this.result.map;
      }
      get root() {
        if (this._root) {
          return this._root;
        }
        let root;
        let parser2 = parse2;
        try {
          root = parser2(this._css, this._opts);
        } catch (error) {
          this.error = error;
        }
        if (this.error) {
          throw this.error;
        } else {
          this._root = root;
          return root;
        }
      }
      get messages() {
        return [];
      }
      warnings() {
        return [];
      }
      toString() {
        return this._css;
      }
      then(onFulfilled, onRejected) {
        if (process.env.NODE_ENV !== "production") {
          if (!("from" in this._opts)) {
            warnOnce(
              "Without `from` option PostCSS could generate wrong source map and will not find Browserslist config. Set it to CSS file path or to `undefined` to prevent this warning."
            );
          }
        }
        return this.async().then(onFulfilled, onRejected);
      }
      catch(onRejected) {
        return this.async().catch(onRejected);
      }
      finally(onFinally) {
        return this.async().then(onFinally, onFinally);
      }
      async() {
        if (this.error)
          return Promise.reject(this.error);
        return Promise.resolve(this.result);
      }
      sync() {
        if (this.error)
          throw this.error;
        return this.result;
      }
    };
    module2.exports = NoWorkResult;
    NoWorkResult.default = NoWorkResult;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/processor.js
var require_processor = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/processor.js"(exports, module2) {
    "use strict";
    var NoWorkResult = require_no_work_result();
    var LazyResult = require_lazy_result();
    var Document = require_document();
    var Root = require_root();
    var Processor = class {
      constructor(plugins = []) {
        this.version = "8.4.24";
        this.plugins = this.normalize(plugins);
      }
      use(plugin3) {
        this.plugins = this.plugins.concat(this.normalize([plugin3]));
        return this;
      }
      process(css, opts = {}) {
        if (this.plugins.length === 0 && typeof opts.parser === "undefined" && typeof opts.stringifier === "undefined" && typeof opts.syntax === "undefined") {
          return new NoWorkResult(this, css, opts);
        } else {
          return new LazyResult(this, css, opts);
        }
      }
      normalize(plugins) {
        let normalized = [];
        for (let i of plugins) {
          if (i.postcss === true) {
            i = i();
          } else if (i.postcss) {
            i = i.postcss;
          }
          if (typeof i === "object" && Array.isArray(i.plugins)) {
            normalized = normalized.concat(i.plugins);
          } else if (typeof i === "object" && i.postcssPlugin) {
            normalized.push(i);
          } else if (typeof i === "function") {
            normalized.push(i);
          } else if (typeof i === "object" && (i.parse || i.stringify)) {
            if (process.env.NODE_ENV !== "production") {
              throw new Error(
                "PostCSS syntaxes cannot be used as plugins. Instead, please use one of the syntax/parser/stringifier options as outlined in your PostCSS runner documentation."
              );
            }
          } else {
            throw new Error(i + " is not a PostCSS plugin");
          }
        }
        return normalized;
      }
    };
    module2.exports = Processor;
    Processor.default = Processor;
    Root.registerProcessor(Processor);
    Document.registerProcessor(Processor);
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/fromJSON.js
var require_fromJSON = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/fromJSON.js"(exports, module2) {
    "use strict";
    var Declaration = require_declaration();
    var PreviousMap = require_previous_map();
    var Comment = require_comment();
    var AtRule = require_at_rule();
    var Input = require_input();
    var Root = require_root();
    var Rule = require_rule();
    function fromJSON(json, inputs) {
      if (Array.isArray(json))
        return json.map((n) => fromJSON(n));
      let { inputs: ownInputs, ...defaults } = json;
      if (ownInputs) {
        inputs = [];
        for (let input of ownInputs) {
          let inputHydrated = { ...input, __proto__: Input.prototype };
          if (inputHydrated.map) {
            inputHydrated.map = {
              ...inputHydrated.map,
              __proto__: PreviousMap.prototype
            };
          }
          inputs.push(inputHydrated);
        }
      }
      if (defaults.nodes) {
        defaults.nodes = json.nodes.map((n) => fromJSON(n, inputs));
      }
      if (defaults.source) {
        let { inputId, ...source } = defaults.source;
        defaults.source = source;
        if (inputId != null) {
          defaults.source.input = inputs[inputId];
        }
      }
      if (defaults.type === "root") {
        return new Root(defaults);
      } else if (defaults.type === "decl") {
        return new Declaration(defaults);
      } else if (defaults.type === "rule") {
        return new Rule(defaults);
      } else if (defaults.type === "comment") {
        return new Comment(defaults);
      } else if (defaults.type === "atrule") {
        return new AtRule(defaults);
      } else {
        throw new Error("Unknown node type: " + json.type);
      }
    }
    module2.exports = fromJSON;
    fromJSON.default = fromJSON;
  }
});

// node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/postcss.js
var require_postcss = __commonJS({
  "node_modules/.pnpm/postcss@8.4.24/node_modules/postcss/lib/postcss.js"(exports, module2) {
    "use strict";
    var CssSyntaxError = require_css_syntax_error();
    var Declaration = require_declaration();
    var LazyResult = require_lazy_result();
    var Container = require_container();
    var Processor = require_processor();
    var stringify = require_stringify();
    var fromJSON = require_fromJSON();
    var Document = require_document();
    var Warning = require_warning();
    var Comment = require_comment();
    var AtRule = require_at_rule();
    var Result = require_result();
    var Input = require_input();
    var parse2 = require_parse();
    var list = require_list();
    var Rule = require_rule();
    var Root = require_root();
    var Node = require_node();
    function postcss(...plugins) {
      if (plugins.length === 1 && Array.isArray(plugins[0])) {
        plugins = plugins[0];
      }
      return new Processor(plugins);
    }
    postcss.plugin = function plugin3(name, initializer) {
      let warningPrinted = false;
      function creator(...args) {
        if (console && console.warn && !warningPrinted) {
          warningPrinted = true;
          console.warn(
            name + ": postcss.plugin was deprecated. Migration guide:\nhttps://evilmartians.com/chronicles/postcss-8-plugin-migration"
          );
          if (process.env.LANG && process.env.LANG.startsWith("cn")) {
            console.warn(
              name + ": \u91CC\u9762 postcss.plugin \u88AB\u5F03\u7528. \u8FC1\u79FB\u6307\u5357:\nhttps://www.w3ctech.com/topic/2226"
            );
          }
        }
        let transformer = initializer(...args);
        transformer.postcssPlugin = name;
        transformer.postcssVersion = new Processor().version;
        return transformer;
      }
      let cache;
      Object.defineProperty(creator, "postcss", {
        get() {
          if (!cache)
            cache = creator();
          return cache;
        }
      });
      creator.process = function(css, processOpts, pluginOpts) {
        return postcss([creator(pluginOpts)]).process(css, processOpts);
      };
      return creator;
    };
    postcss.stringify = stringify;
    postcss.parse = parse2;
    postcss.fromJSON = fromJSON;
    postcss.list = list;
    postcss.comment = (defaults) => new Comment(defaults);
    postcss.atRule = (defaults) => new AtRule(defaults);
    postcss.decl = (defaults) => new Declaration(defaults);
    postcss.rule = (defaults) => new Rule(defaults);
    postcss.root = (defaults) => new Root(defaults);
    postcss.document = (defaults) => new Document(defaults);
    postcss.CssSyntaxError = CssSyntaxError;
    postcss.Declaration = Declaration;
    postcss.Container = Container;
    postcss.Processor = Processor;
    postcss.Document = Document;
    postcss.Comment = Comment;
    postcss.Warning = Warning;
    postcss.AtRule = AtRule;
    postcss.Result = Result;
    postcss.Input = Input;
    postcss.Rule = Rule;
    postcss.Root = Root;
    postcss.Node = Node;
    LazyResult.registerPostcss(postcss);
    module2.exports = postcss;
    postcss.default = postcss;
  }
});

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/parser.js
var require_parser2 = __commonJS({
  "node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/parser.js"(exports, module2) {
    var postcss = require_postcss();
    var IMPORTANT = /\s*!important\s*$/i;
    var UNITLESS = {
      "box-flex": true,
      "box-flex-group": true,
      "column-count": true,
      "flex": true,
      "flex-grow": true,
      "flex-positive": true,
      "flex-shrink": true,
      "flex-negative": true,
      "font-weight": true,
      "line-clamp": true,
      "line-height": true,
      "opacity": true,
      "order": true,
      "orphans": true,
      "tab-size": true,
      "widows": true,
      "z-index": true,
      "zoom": true,
      "fill-opacity": true,
      "stroke-dashoffset": true,
      "stroke-opacity": true,
      "stroke-width": true
    };
    function dashify(str) {
      return str.replace(/([A-Z])/g, "-$1").replace(/^ms-/, "-ms-").toLowerCase();
    }
    function decl(parent, name, value) {
      if (value === false || value === null)
        return;
      if (!name.startsWith("--")) {
        name = dashify(name);
      }
      if (typeof value === "number") {
        if (value === 0 || UNITLESS[name]) {
          value = value.toString();
        } else {
          value += "px";
        }
      }
      if (name === "css-float")
        name = "float";
      if (IMPORTANT.test(value)) {
        value = value.replace(IMPORTANT, "");
        parent.push(postcss.decl({ prop: name, value, important: true }));
      } else {
        parent.push(postcss.decl({ prop: name, value }));
      }
    }
    function atRule(parent, parts, value) {
      let node = postcss.atRule({ name: parts[1], params: parts[3] || "" });
      if (typeof value === "object") {
        node.nodes = [];
        parse2(value, node);
      }
      parent.push(node);
    }
    function parse2(obj, parent) {
      let name, value, node;
      for (name in obj) {
        value = obj[name];
        if (value === null || typeof value === "undefined") {
          continue;
        } else if (name[0] === "@") {
          let parts = name.match(/@(\S+)(\s+([\W\w]*)\s*)?/);
          if (Array.isArray(value)) {
            for (let i of value) {
              atRule(parent, parts, i);
            }
          } else {
            atRule(parent, parts, value);
          }
        } else if (Array.isArray(value)) {
          for (let i of value) {
            decl(parent, name, i);
          }
        } else if (typeof value === "object") {
          node = postcss.rule({ selector: name });
          parse2(value, node);
          parent.push(node);
        } else {
          decl(parent, name, value);
        }
      }
    }
    module2.exports = function(obj) {
      let root = postcss.root();
      parse2(obj, root);
      return root;
    };
  }
});

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/process-result.js
var require_process_result = __commonJS({
  "node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/process-result.js"(exports, module2) {
    var objectify2 = require_objectifier();
    module2.exports = function processResult(result) {
      if (console && console.warn) {
        result.warnings().forEach((warn) => {
          let source = warn.plugin || "PostCSS";
          console.warn(source + ": " + warn.text);
        });
      }
      return objectify2(result.root);
    };
  }
});

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/async.js
var require_async = __commonJS({
  "node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/async.js"(exports, module2) {
    var postcss = require_postcss();
    var processResult = require_process_result();
    var parse2 = require_parser2();
    module2.exports = function async2(plugins) {
      let processor = postcss(plugins);
      return async (input) => {
        let result = await processor.process(input, {
          parser: parse2,
          from: void 0
        });
        return processResult(result);
      };
    };
  }
});

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/sync.js
var require_sync = __commonJS({
  "node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/sync.js"(exports, module2) {
    var postcss = require_postcss();
    var processResult = require_process_result();
    var parse2 = require_parser2();
    module2.exports = function(plugins) {
      let processor = postcss(plugins);
      return (input) => {
        let result = processor.process(input, { parser: parse2, from: void 0 });
        return processResult(result);
      };
    };
  }
});

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/index.js
var require_postcss_js = __commonJS({
  "node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/index.js"(exports, module2) {
    var objectify2 = require_objectifier();
    var parse2 = require_parser2();
    var async2 = require_async();
    var sync2 = require_sync();
    module2.exports = {
      objectify: objectify2,
      parse: parse2,
      async: async2,
      sync: sync2
    };
  }
});

// src/tailwind/generated/generated-classes.js
var require_generated_classes = __commonJS({
  "src/tailwind/generated/generated-classes.js"(exports, module2) {
    "use strict";
    module2.exports = { components: { ".hide-scrollbar::-webkit-scrollbar": { "display": "none" }, ".hide-scrollbar": { "msOverflowStyle": "none", "scrollbarWidth": "none" }, ".divider-vertical": { "marginLeft": "auto", "marginRight": "auto", "display": "inline-block", "minHeight": "10px", "borderLeftWidth": "1px", "borderStyle": "solid", "borderColor": "rgb(var(--color-surface-300))" }, ".dark .divider-vertical": { "borderColor": "rgb(var(--color-surface-600))" }, ".h1": { "fontSize": "1.875rem", "lineHeight": "2.25rem", "fontFamily": "var(--theme-font-family-heading)" }, ".h2": { "fontSize": "1.5rem", "lineHeight": "2rem", "fontFamily": "var(--theme-font-family-heading)" }, ".h3": { "fontSize": "1.25rem", "lineHeight": "1.75rem", "fontFamily": "var(--theme-font-family-heading)" }, ".h4": { "fontSize": "1.125rem", "lineHeight": "1.75rem", "fontFamily": "var(--theme-font-family-heading)" }, ".h5": { "fontSize": "1rem", "lineHeight": "1.5rem", "fontFamily": "var(--theme-font-family-heading)" }, ".h6": { "fontSize": "0.875rem", "lineHeight": "1.25rem", "fontFamily": "var(--theme-font-family-heading)" }, ".anchor": { "--tw-text-opacity": "1", "color": "rgb(var(--color-primary-700) / var(--tw-text-opacity))", "textDecorationLine": "underline" }, ".anchor:hover": { "--tw-brightness": "brightness(1.1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ":is(.dark .anchor)": { "--tw-text-opacity": "1", "color": "rgb(var(--color-primary-500) / var(--tw-text-opacity))" }, ".blockquote": { "borderLeftWidth": "8px", "--tw-border-opacity": "1", "borderLeftColor": "rgb(var(--color-secondary-500) / var(--tw-border-opacity))", "paddingRight": "1rem", "paddingLeft": "1rem", "fontSize": "1rem", "lineHeight": "1.5rem", "fontStyle": "italic", "color": "rgba(var(--theme-font-color-base))" }, ".dark .blockquote": { "color": "rgba(var(--theme-font-color-dark))" }, ".kbd": { "fontFamily": 'ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"', "fontSize": "0.875rem", "lineHeight": "1.25rem", "fontWeight": 700, "borderRadius": "0.25rem", "paddingLeft": "0.375rem", "paddingRight": "0.375rem", "paddingTop": "3px", "paddingBottom": "3px", "backgroundColor": "rgb(var(--color-surface-300))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-900) / var(--tw-ring-opacity))", "borderBottomWidth": "2px", "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-surface-900) / var(--tw-border-opacity))" }, ".dark .kbd": { "backgroundColor": "rgb(var(--color-surface-600))" }, ".time": { "fontSize": "0.875rem", "lineHeight": "1.25rem", "--tw-text-opacity": "1", "color": "rgb(var(--color-surface-500) / var(--tw-text-opacity))" }, ":is(.dark .time)": { "--tw-text-opacity": "1", "color": "rgb(var(--color-surface-400) / var(--tw-text-opacity))" }, ".pre": { "overflowX": "auto", "whiteSpace": "pre-wrap", "backgroundColor": "rgb(23 23 23 / 0.9)", "padding": "1rem", "fontFamily": 'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace', "fontSize": "1rem", "lineHeight": "1.5rem", "--tw-text-opacity": "1", "color": "rgb(255 255 255 / var(--tw-text-opacity))", "borderRadius": "var(--theme-rounded-container)" }, ".code": { "whiteSpace": "nowrap", "fontFamily": 'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace', "fontSize": "0.75rem", "lineHeight": "1rem", "--tw-text-opacity": "1", "color": "rgb(var(--color-primary-700) / var(--tw-text-opacity))", "backgroundColor": "rgb(var(--color-primary-500) / 0.3)", "borderRadius": "0.25rem", "paddingTop": "0.125rem", "paddingBottom": "0.125rem", "paddingLeft": "0.25rem", "paddingRight": "0.25rem" }, ":is(.dark .code)": { "--tw-text-opacity": "1", "color": "rgb(var(--color-primary-400) / var(--tw-text-opacity))", "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".ins,\n	.del": { "position": "relative", "display": "block", "padding": "0.125rem", "paddingLeft": "1.25rem", "textDecoration": "none" }, ".ins::before,\n	.del::before": { "position": "absolute", "left": "0.25rem", "fontFamily": 'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace' }, ".ins::before": { "content": "'+'" }, ".del::before": { "content": "'\u2212'" }, ".ins": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-500) / var(--tw-bg-opacity))", "fontFamily": 'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace', "color": "rgb(var(--on-success))" }, ".del": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-500) / var(--tw-bg-opacity))", "fontFamily": 'ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace', "color": "rgb(var(--on-error))" }, ".alert": { "display": "flex", "flexDirection": "column", "alignItems": "flex-start", "padding": "1rem", "color": "rgb(var(--color-surface-900))", "borderRadius": "var(--theme-rounded-container)" }, ".alert > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(1rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(1rem * var(--tw-space-y-reverse))" }, ".dark .alert": { "color": "rgb(var(--color-surface-50))" }, ".alert-message": { "flex": "1 1 auto" }, ".alert-message > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0.5rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0.5rem * var(--tw-space-y-reverse))" }, ".alert-actions": { "display": "flex", "alignItems": "center" }, ".alert-actions > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".badge": { "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "whiteSpace": "nowrap", "fontSize": "0.75rem", "lineHeight": "1rem", "fontWeight": 600, "paddingLeft": "0.5rem", "paddingRight": "0.5rem", "paddingTop": "0.25rem", "paddingBottom": "0.25rem", "borderRadius": "var(--theme-rounded-base)" }, ".badge > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".badge-icon": { "display": "flex", "height": "1.25rem", "width": "1.25rem", "alignItems": "center", "justifyContent": "center", "borderRadius": "9999px", "fontSize": "0.75rem", "lineHeight": "1rem", "fontWeight": 600, "--tw-shadow": "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)", "--tw-shadow-colored": "0 1px 3px 0 var(--tw-shadow-color), 0 1px 2px -1px var(--tw-shadow-color)", "boxShadow": "var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)" }, ".badge-glass": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-color": "rgb(23 23 23 / 0.05)" }, ":is(.dark .badge-glass)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)", "--tw-ring-color": "rgb(250 250 250 / 0.05)" }, ".breadcrumb::-webkit-scrollbar,\n	.breadcrumb-nonresponsive::-webkit-scrollbar": { "display": "none" }, ".breadcrumb,\n	.breadcrumb-nonresponsive": { "msOverflowStyle": "none", "scrollbarWidth": "none", "display": "flex", "width": "100%", "alignItems": "center", "overflowX": "auto" }, ".breadcrumb > :not([hidden]) ~ :not([hidden]),\n	.breadcrumb-nonresponsive > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(1rem * var(--tw-space-x-reverse))", "marginLeft": "calc(1rem * calc(1 - var(--tw-space-x-reverse)))" }, ".crumb": { "display": "flex", "alignItems": "center", "justifyContent": "center" }, ".crumb > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".crumb-separator": { "display": "flex", "opacity": 0.5, "color": "rgb(var(--color-surface-700))" }, ".dark .crumb-separator": { "color": "rgb(var(--color-surface-200))" }, ".breadcrumb li": { "display": "none" }, ".breadcrumb li:nth-last-child(3),\n	.breadcrumb li:nth-last-child(2),\n	.breadcrumb li:nth-last-child(1)": { "display": "block" }, ".button-base-styles": { "fontSize": "1rem", "lineHeight": "1.5rem", "paddingLeft": "1.25rem", "paddingRight": "1.25rem", "paddingTop": "9px", "paddingBottom": "9px", "whiteSpace": "nowrap", "textAlign": "center", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms" }, ".button-base-styles > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".button-base-styles:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn:disabled,\n	.btn-icon:disabled,\n	.btn-group > *:disabled": { "cursor": "not-allowed !important", "opacity": "0.5 !important" }, ".btn:disabled:hover,\n	.btn-icon:disabled:hover,\n	.btn-group > *:disabled:hover": { "--tw-brightness": "brightness(1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn:disabled:active,\n	.btn-icon:disabled:active,\n	.btn-group > *:disabled:active": { "--tw-scale-x": "1", "--tw-scale-y": "1", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))" }, ".btn": { "fontSize": "1rem", "lineHeight": "1.5rem", "paddingLeft": "1.25rem", "paddingRight": "1.25rem", "paddingTop": "9px", "paddingBottom": "9px", "whiteSpace": "nowrap", "textAlign": "center", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms", "borderRadius": "var(--theme-rounded-base)" }, ".btn > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".btn:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn:active": { "--tw-scale-x": "95%", "--tw-scale-y": "95%", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))", "--tw-brightness": "brightness(.9)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn-sm": { "paddingLeft": "0.75rem", "paddingRight": "0.75rem", "paddingTop": "0.375rem", "paddingBottom": "0.375rem", "fontSize": "0.875rem", "lineHeight": "1.25rem" }, ".btn-md,\n	.btn-base": { "paddingLeft": "1.25rem", "paddingRight": "1.25rem", "paddingTop": "9px", "paddingBottom": "9px", "fontSize": "1rem", "lineHeight": "1.5rem" }, ".btn-lg": { "paddingLeft": "1.75rem", "paddingRight": "1.75rem", "paddingTop": "0.75rem", "paddingBottom": "0.75rem", "fontSize": "1.125rem", "lineHeight": "1.75rem" }, ".btn-xl": { "paddingLeft": "2.25rem", "paddingRight": "2.25rem", "paddingTop": "1rem", "paddingBottom": "1rem", "fontSize": "1.25rem", "lineHeight": "1.75rem" }, ".btn-icon": { "fontSize": "1rem", "lineHeight": "1.5rem", "paddingLeft": "1.25rem", "paddingRight": "1.25rem", "paddingTop": "9px", "paddingBottom": "9px", "whiteSpace": "nowrap", "textAlign": "center", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms", "padding": "0px", "aspectRatio": "1 / 1", "width": "43px", "borderRadius": "9999px" }, ".btn-icon > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".btn-icon:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn-icon:active": { "--tw-scale-x": "95%", "--tw-scale-y": "95%", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))", "--tw-brightness": "brightness(.9)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn-icon-sm": { "aspectRatio": "1 / 1", "width": "33px", "fontSize": "0.875rem", "lineHeight": "1.25rem" }, ".btn-icon-base": { "aspectRatio": "1 / 1", "width": "43px", "fontSize": "1rem", "lineHeight": "1.5rem" }, ".btn-icon-lg": { "aspectRatio": "1 / 1", "width": "53px", "fontSize": "1.125rem", "lineHeight": "1.75rem" }, ".btn-icon-xl": { "aspectRatio": "1 / 1", "width": "63px", "fontSize": "1.25rem", "lineHeight": "1.75rem" }, ".btn-group": { "display": "inline-flex", "flexDirection": "row", "overflow": "hidden", "borderRadius": "var(--theme-rounded-base)", "isolation": "isolate" }, ".btn-group > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0px * var(--tw-space-x-reverse))", "marginLeft": "calc(0px * calc(1 - var(--tw-space-x-reverse)))" }, ".btn-group-vertical > *:disabled": { "cursor": "not-allowed !important", "opacity": "0.5 !important" }, ".btn-group-vertical > *:disabled:hover": { "--tw-brightness": "brightness(1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".btn-group-vertical > *:disabled:active": { "--tw-scale-x": "1", "--tw-scale-y": "1", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))" }, ".btn-group-vertical": { "display": "inline-flex", "flexDirection": "column", "overflow": "hidden", "borderRadius": "var(--theme-rounded-container)", "isolation": "isolate" }, ".btn-group-vertical > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0px * var(--tw-space-x-reverse))", "marginLeft": "calc(0px * calc(1 - var(--tw-space-x-reverse)))", "--tw-space-y-reverse": "0", "marginTop": "calc(0px * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0px * var(--tw-space-y-reverse))" }, ".btn-group-vertical button,.btn-group-vertical a": { "fontSize": "1rem", "lineHeight": "1.5rem", "paddingLeft": "1.25rem", "paddingRight": "1.25rem", "paddingTop": "9px", "paddingBottom": "9px", "whiteSpace": "nowrap", "textAlign": "center", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms", "color": "inherit !important", "textDecorationLine": "none !important" }, ".btn-group-vertical button > :not([hidden]) ~ :not([hidden]),.btn-group-vertical a > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".btn-group-vertical button:hover,.btn-group-vertical a:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)", "backgroundColor": "rgb(var(--color-surface-50) / 3%)" }, ".btn-group-vertical button:active,.btn-group-vertical a:active": { "backgroundColor": "rgb(var(--color-surface-900) / 3%)" }, ".btn-group-vertical > * + *": { "borderTopWidth": "1px", "borderLeftWidth": "0px", "borderColor": "rgb(var(--color-surface-500) / 0.2)" }, ".btn-group button,\n	.btn-group a,\n	.btn-group-vertical button,\n	.btn-group-vertical a": { "fontSize": "1rem", "lineHeight": "1.5rem", "paddingLeft": "1.25rem", "paddingRight": "1.25rem", "paddingTop": "9px", "paddingBottom": "9px", "whiteSpace": "nowrap", "textAlign": "center", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms", "color": "inherit !important", "textDecorationLine": "none !important" }, ".btn-group button > :not([hidden]) ~ :not([hidden]),\n	.btn-group a > :not([hidden]) ~ :not([hidden]),\n	.btn-group-vertical button > :not([hidden]) ~ :not([hidden]),\n	.btn-group-vertical a > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, ".btn-group button:hover,\n	.btn-group a:hover,\n	.btn-group-vertical button:hover,\n	.btn-group-vertical a:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)", "backgroundColor": "rgb(var(--color-surface-50) / 3%)" }, ".btn-group button:active,\n	.btn-group a:active,\n	.btn-group-vertical button:active,\n	.btn-group-vertical a:active": { "backgroundColor": "rgb(var(--color-surface-900) / 3%)" }, ".btn-group > * + *": { "borderTopWidth": "0px", "borderLeftWidth": "1px", "borderColor": "rgb(var(--color-surface-500) / 0.2)" }, ".card": { "backgroundColor": "rgb(var(--color-surface-100))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-color": "rgb(23 23 23 / 0.05)", "borderRadius": "var(--theme-rounded-container)" }, ".dark .card": { "backgroundColor": "rgb(var(--color-surface-800))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-color": "rgb(250 250 250 / 0.05)" }, ".card-header": { "padding": "1rem", "paddingBottom": "0px" }, ".card-footer": { "padding": "1rem", "paddingTop": "0px" }, "a.card": { "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms" }, "a.card:hover": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".card-hover": { "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms" }, ".card-hover:hover": { "--tw-scale-x": "101%", "--tw-scale-y": "101%", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))", "--tw-shadow": "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)", "--tw-shadow-colored": "0 20px 25px -5px var(--tw-shadow-color), 0 8px 10px -6px var(--tw-shadow-color)", "boxShadow": "var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)" }, ".codeblock pre": { "counterReset": "line" }, ".codeblock code.lineNumbers > span.line": { "counterIncrement": "line" }, ".codeblock code.lineNumbers > span.line::before": { "content": "counter(line) ''", "opacity": 0.5 }, ".chip": { "cursor": "pointer", "whiteSpace": "nowrap", "paddingLeft": "0.75rem", "paddingRight": "0.75rem", "paddingTop": "0.375rem", "paddingBottom": "0.375rem", "textAlign": "center", "fontSize": "0.75rem", "lineHeight": "1rem", "borderRadius": "0.25rem", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms" }, ".chip > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, "a.chip:hover,\n	button.chip:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".chip-disabled,\n	.chip:disabled": { "cursor": "not-allowed !important", "opacity": "0.5 !important" }, ".chip-disabled:active,\n	.chip:disabled:active": { "--tw-scale-x": "1", "--tw-scale-y": "1", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))" }, ".legend": { "fontSize": "1.25rem", "lineHeight": "1.75rem", "fontFamily": "var(--theme-font-family-heading)" }, ".label > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0.25rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0.25rem * var(--tw-space-y-reverse))" }, ".input,\n	.textarea,\n	.select,\n	.input-group": { "width": "100%", "transitionProperty": "color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "200ms", "backgroundColor": "rgb(var(--color-surface-200))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "borderWidth": "var(--theme-border-base)", "borderColor": "rgb(var(--color-surface-400))" }, ".dark .input,.dark \n	.textarea,.dark \n	.select,.dark \n	.input-group": { "backgroundColor": "rgb(var(--color-surface-700))", "borderColor": "rgb(var(--color-surface-500))" }, ".input:hover,\n	.textarea:hover,\n	.select:hover,\n	.input-group:hover": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".input:focus,\n	.textarea:focus,\n	.select:focus,\n	.input-group:focus": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".input:focus-within,\n	.textarea:focus-within,\n	.select:focus-within,\n	.input-group:focus-within": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-primary-500) / var(--tw-border-opacity))" }, ".input,\n	.input-group": { "borderRadius": "var(--theme-rounded-base)" }, ".textarea,\n	.select": { "borderRadius": "var(--theme-rounded-container)" }, ".select > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0.25rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0.25rem * var(--tw-space-y-reverse))" }, ".select": { "padding": "0.5rem", "paddingRight": "2rem" }, ".select[size]": { "backgroundImage": "none" }, ".select optgroup > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0.25rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0.25rem * var(--tw-space-y-reverse))" }, ".select optgroup": { "fontWeight": 700 }, ".select optgroup option": { "marginLeft": "0px", "paddingLeft": "0px" }, ".select optgroup option:first-of-type": { "marginTop": "0.75rem" }, ".select optgroup option:last-child": { "marginBottom": "0.75rem !important" }, ".select option": { "cursor": "pointer", "paddingLeft": "1rem", "paddingRight": "1rem", "paddingTop": "0.5rem", "paddingBottom": "0.5rem", "backgroundColor": "rgb(var(--color-surface-200))", "borderRadius": "var(--theme-rounded-base)" }, ".dark .select option": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".select option:checked": { "background": "rgb(var(--color-primary-500)) linear-gradient(0deg, rgb(var(--color-primary-500)) 0%, rgb(var(--color-primary-500)) 100%)", "color": "rgb(var(--on-primary))" }, ".checkbox,\n	.radio": { "height": "1.25rem", "width": "1.25rem", "cursor": "pointer", "borderRadius": "0.25rem", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "backgroundColor": "rgb(var(--color-surface-200))", "borderWidth": "var(--theme-border-base)", "borderColor": "rgb(var(--color-surface-400))" }, ".dark .checkbox,.dark \n	.radio": { "backgroundColor": "rgb(var(--color-surface-700))", "borderColor": "rgb(var(--color-surface-500))" }, ".checkbox:hover,\n	.radio:hover": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".checkbox:focus,\n	.radio:focus": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)", "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-primary-500) / var(--tw-border-opacity))" }, ".checkbox:checked,\n	.checkbox:indeterminate,\n	.radio:checked": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-primary-500) / var(--tw-bg-opacity))" }, ".checkbox:checked:hover,\n	.checkbox:indeterminate:hover,\n	.radio:checked:hover": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-primary-500) / var(--tw-bg-opacity))" }, ".checkbox:checked:focus,\n	.checkbox:indeterminate:focus,\n	.radio:checked:focus": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-primary-500) / var(--tw-bg-opacity))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)" }, ".radio": { "borderRadius": "var(--theme-rounded-base)" }, ".input[type='file']": { "padding": "0.25rem" }, ".input[type='color']": { "height": "2.5rem", "width": "2.5rem", "cursor": "pointer", "overflow": "hidden", "borderStyle": "none", "borderRadius": "var(--theme-rounded-base)", "WebkitAppearance": "none" }, ".input[type='color']::-webkit-color-swatch-wrapper": { "padding": "0px" }, ".input[type='color']::-webkit-color-swatch": { "borderStyle": "none" }, ".input[type='color']::-webkit-color-swatch:hover": { "--tw-brightness": "brightness(1.1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".input[type='color']::-moz-color-swatch": { "borderStyle": "none" }, ".input:disabled,\n	.textarea:disabled,\n	.select:disabled,\n	.input-group > input:disabled,\n	.input-group > textarea:disabled,\n	.input-group > select:disabled": { "cursor": "not-allowed !important", "opacity": "0.5 !important" }, ".input:disabled:hover,\n	.textarea:disabled:hover,\n	.select:disabled:hover,\n	.input-group > input:disabled:hover,\n	.input-group > textarea:disabled:hover,\n	.input-group > select:disabled:hover": { "--tw-brightness": "brightness(1) !important", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow) !important" }, ".input[readonly],\n	.textarea[readonly],\n	.select[readonly]": { "cursor": "not-allowed !important", "borderColor": "transparent !important" }, ".input[readonly]:hover,\n	.textarea[readonly]:hover,\n	.select[readonly]:hover": { "--tw-brightness": "brightness(1) !important", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow) !important" }, ".input-group": { "display": "grid", "overflow": "hidden" }, ".input-group input,\n	.input-group select": { "borderWidth": "0px", "backgroundColor": "transparent", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important" }, ".input-group select option": { "backgroundColor": "rgb(var(--color-surface-200))" }, ".dark .input-group select option": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".input-group div,\n	.input-group a,\n	.input-group button": { "display": "flex", "alignItems": "center", "justifyContent": "space-between", "paddingLeft": "1rem", "paddingRight": "1rem" }, ".input-group-divider input,\n	.input-group-divider select,\n	.input-group-divider div,\n	.input-group-divider a": { "borderLeftWidth": "1px", "borderColor": "rgb(var(--color-surface-400))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "minWidth": "fit-content !important" }, ".dark .input-group-divider input,.dark \n	.input-group-divider select,.dark \n	.input-group-divider div,.dark \n	.input-group-divider a": { "borderColor": "rgb(var(--color-surface-500))" }, ".input-group-divider input:focus,\n	.input-group-divider select:focus,\n	.input-group-divider div:focus,\n	.input-group-divider a:focus": { "borderColor": "rgb(var(--color-surface-400))" }, ".dark .input-group-divider input:focus,.dark \n	.input-group-divider select:focus,.dark \n	.input-group-divider div:focus,.dark \n	.input-group-divider a:focus": { "borderColor": "rgb(var(--color-surface-500))" }, ".input-group-divider *:first-child": { "borderLeftWidth": "0px !important" }, ".input-group-shim": { "backgroundColor": "rgb(var(--color-surface-400) / 0.1)", "color": "rgb(var(--color-surface-600))" }, ".dark .input-group-shim": { "color": "rgb(var(--color-surface-300))" }, ".input-success": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-success-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-success-700) / var(--tw-text-opacity))" }, ":is(.dark .input-success)": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-success-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-success-700) / var(--tw-text-opacity))" }, ".input-success::-moz-placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-success-700) / var(--tw-text-opacity))" }, ".input-success:-ms-input-placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-success-700) / var(--tw-text-opacity))" }, ".input-success::placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-success-700) / var(--tw-text-opacity))" }, ".input-warning": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-warning-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-warning-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-warning-700) / var(--tw-text-opacity))" }, ":is(.dark .input-warning)": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-warning-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-warning-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-warning-700) / var(--tw-text-opacity))" }, ".input-warning::-moz-placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-warning-700) / var(--tw-text-opacity))" }, ".input-warning:-ms-input-placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-warning-700) / var(--tw-text-opacity))" }, ".input-warning::placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-warning-700) / var(--tw-text-opacity))" }, ".input-error": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-error-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-error-500) / var(--tw-text-opacity))" }, ":is(.dark .input-error)": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-error-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-error-500) / var(--tw-text-opacity))" }, ".input-error::-moz-placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-error-500) / var(--tw-text-opacity))" }, ".input-error:-ms-input-placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-error-500) / var(--tw-text-opacity))" }, ".input-error::placeholder": { "--tw-text-opacity": "1", "color": "rgb(var(--color-error-500) / var(--tw-text-opacity))" }, ".variant-form-material": { "borderTopLeftRadius": "0.25rem !important", "borderTopRightRadius": "0.25rem !important", "borderBottomLeftRadius": "0px !important", "borderBottomRightRadius": "0px !important", "backgroundColor": "rgb(var(--color-surface-500) / 0.1)", "borderWidth": "0px", "borderBottomWidth": "2px" }, ":is(.dark .variant-form-material)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".variant-form-material[type='file']": { "paddingTop": "0.375rem !important", "paddingBottom": "0.375rem !important" }, ".list,\n	.list-dl,\n	.list-nav ul": { "listStyleType": "none" }, ".list > :not([hidden]) ~ :not([hidden]),\n	.list-dl > :not([hidden]) ~ :not([hidden]),\n	.list-nav ul > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0.25rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0.25rem * var(--tw-space-y-reverse))" }, ".list li": { "display": "flex", "alignItems": "center", "borderRadius": "var(--theme-rounded-base)" }, ".list li > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(1rem * var(--tw-space-x-reverse))", "marginLeft": "calc(1rem * calc(1 - var(--tw-space-x-reverse)))" }, ".list-dl div": { "display": "flex", "alignItems": "center", "padding": "0.5rem", "borderRadius": "var(--theme-rounded-base)" }, ".list-dl div > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(1rem * var(--tw-space-x-reverse))", "marginLeft": "calc(1rem * calc(1 - var(--tw-space-x-reverse)))" }, ".list-nav a,\n	.list-nav button,\n	.list-option": { "display": "flex", "alignItems": "center", "paddingLeft": "1rem", "paddingRight": "1rem", "paddingTop": "0.5rem", "paddingBottom": "0.5rem", "outline": "2px solid transparent", "outlineOffset": "2px", "cursor": "pointer", "borderRadius": "var(--theme-rounded-base)" }, ".list-nav a > :not([hidden]) ~ :not([hidden]),\n	.list-nav button > :not([hidden]) ~ :not([hidden]),\n	.list-option > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(1rem * var(--tw-space-x-reverse))", "marginLeft": "calc(1rem * calc(1 - var(--tw-space-x-reverse)))" }, ".list-nav a:hover,\n	.list-nav button:hover,\n	.list-option:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .list-nav a:hover,.dark \n	.list-nav button:hover,.dark \n	.list-option:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".list-nav a:focus,\n	.list-nav button:focus,\n	.list-option:focus": { "backgroundColor": "rgb(var(--color-primary-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-primary-700))" }, ".dark .list-nav a:focus,.dark \n	.list-nav button:focus,.dark \n	.list-option:focus": { "color": "rgb(var(--color-primary-200))" }, ":is(.dark .list-nav a:focus),:is(.dark \n	.list-nav button:focus),:is(.dark \n	.list-option:focus)": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".logo-cloud": { "display": "grid", "overflow": "hidden", "borderRadius": "var(--theme-rounded-container)" }, ".logo-item": { "flex": "1 1 auto", "textAlign": "center", "--tw-shadow": "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)", "--tw-shadow-colored": "0 1px 3px 0 var(--tw-shadow-color), 0 1px 2px -1px var(--tw-shadow-color)", "boxShadow": "var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)", "display": "flex", "alignItems": "center", "justifyContent": "center", "backgroundColor": "rgb(var(--color-surface-100))", "fontSize": "1rem", "lineHeight": "1.5rem", "fontWeight": 700, "--tw-text-opacity": "1", "color": "rgb(0 0 0 / var(--tw-text-opacity))", "paddingTop": "1rem", "paddingBottom": "1rem" }, ".logo-item > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(1rem * var(--tw-space-x-reverse))", "marginLeft": "calc(1rem * calc(1 - var(--tw-space-x-reverse)))" }, ".dark .logo-item": { "backgroundColor": "rgb(var(--color-surface-800))" }, ":is(.dark .logo-item)": { "--tw-text-opacity": "1", "color": "rgb(255 255 255 / var(--tw-text-opacity))" }, "a.logo-item:hover": { "--tw-brightness": "brightness(1.1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".placeholder": { "height": "1.25rem", "backgroundColor": "rgb(var(--color-surface-300))", "borderRadius": "var(--theme-rounded-base)" }, ".dark .placeholder": { "backgroundColor": "rgb(var(--color-surface-600))" }, ".placeholder-circle": { "aspectRatio": "1 / 1", "borderRadius": "9999px", "backgroundColor": "rgb(var(--color-surface-300))" }, ".dark .placeholder-circle": { "backgroundColor": "rgb(var(--color-surface-600))" }, ".table-container": { "width": "100%", "overflowX": "auto", "borderRadius": "var(--theme-rounded-container)" }, ".table": { "width": "100%", "tableLayout": "auto", "overflow": "hidden", "backgroundColor": "rgb(var(--color-surface-100))", "borderRadius": "var(--theme-rounded-container)", "display": "table" }, ".dark .table": { "backgroundColor": "rgb(var(--color-surface-800))" }, ".table-hover tbody tr:hover": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".table-hover tbody tr:hover:nth-child(even)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".table-interactive tbody tr": { "cursor": "pointer" }, ".table-interactive tbody tr:hover:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .table-interactive tbody tr:hover:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".table-interactive tbody tr:hover:nth-child(even):hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .table-interactive tbody tr:hover:nth-child(even):hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".table-sort-asc::after": { "opacity": 0.5, "--tw-content": "'\u2191' !important", "content": "var(--tw-content) !important" }, ".table-sort-dsc::after": { "opacity": 0.5, "--tw-content": "'\u2193' !important", "content": "var(--tw-content) !important" }, ".table thead": { "borderBottomWidth": "1px", "borderColor": "rgb(var(--color-surface-500) / 0.2)", "backgroundColor": "rgb(var(--color-surface-200))" }, ".dark .table thead": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".table thead tr": { "textAlign": "left", "textTransform": "capitalize" }, ".table thead th": { "padding": "1rem", "fontWeight": 700 }, ".table tbody tr": { "borderBottomWidth": "1px", "borderColor": "rgb(var(--color-surface-500) / 0.2)" }, ".table tbody tr:nth-child(even)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.05)" }, ".table tbody td": { "whiteSpace": "nowrap", "paddingLeft": "0.75rem", "paddingRight": "0.75rem", "paddingTop": "1rem", "paddingBottom": "1rem", "verticalAlign": "top", "fontSize": "0.875rem", "lineHeight": "1.25rem" }, ".table-compact tbody td": { "paddingTop": "0.75rem !important", "paddingBottom": "0.75rem !important" }, ".table-comfortable tbody td": { "paddingTop": "1.25rem !important", "paddingBottom": "1.25rem !important" }, ".table tfoot": { "backgroundColor": "rgb(var(--color-surface-100))" }, ".dark .table tfoot": { "backgroundColor": "rgb(var(--color-surface-800))" }, ".table tfoot tr": { "textAlign": "left", "textTransform": "capitalize" }, ".table tfoot th,\n	.table tfoot td": { "padding": "1rem" }, ".table-row-checked": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.2) !important" }, ".table-cell-fit": { "width": "1%", "whiteSpace": "nowrap" }, ".w-modal-slim": { "width": "100%", "maxWidth": "400px" }, ".w-modal": { "width": "100%", "maxWidth": "640px" }, ".w-modal-wide": { "width": "100%", "maxWidth": "80%" }, ".modal *:focus:not([tabindex='-1']):not(.input):not(.textarea):not(.select):not(.input-group)": { "outlineStyle": "auto", "outlineColor": "-webkit-focus-ring-color" }, "[data-popup] .arrow": { "position": "absolute", "height": "0.5rem", "width": "0.5rem", "--tw-rotate": "45deg", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))" }, ".variant-outline-primary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-primary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))" }, ".variant-outline-secondary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-secondary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))" }, ".variant-outline-tertiary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-tertiary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))" }, ".variant-outline-success": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-success)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))" }, ".variant-outline-warning": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-warning)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))" }, ".variant-outline-error": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-error)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))" }, ".variant-outline,\n	.variant-outline-surface": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline),:is(.dark \n	.variant-outline-surface)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))" }, ".variant-filled": { "backgroundColor": "rgb(var(--color-surface-900))", "color": "rgb(var(--color-surface-50))" }, ".dark .variant-filled": { "backgroundColor": "rgb(var(--color-surface-50))", "color": "rgb(var(--color-surface-900))" }, ".variant-filled-primary": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-primary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-primary))" }, ":is(.dark .variant-filled-primary)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-primary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-primary))" }, ".variant-filled-secondary": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-secondary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-secondary))" }, ":is(.dark .variant-filled-secondary)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-secondary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-secondary))" }, ".variant-filled-tertiary": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-tertiary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-tertiary))" }, ":is(.dark .variant-filled-tertiary)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-tertiary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-tertiary))" }, ".variant-filled-success": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-success))" }, ":is(.dark .variant-filled-success)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-success))" }, ".variant-filled-warning": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-warning-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-warning))" }, ":is(.dark .variant-filled-warning)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-warning-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-warning))" }, ".variant-filled-error": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-error))" }, ":is(.dark .variant-filled-error)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-error))" }, ".variant-filled-surface": { "backgroundColor": "rgb(var(--color-surface-400))", "color": "rgb(var(--on-surface))" }, ".dark .variant-filled-surface": { "backgroundColor": "rgb(var(--color-surface-500))" }, ":is(.dark .variant-filled-surface)": { "color": "rgb(var(--on-surface))" }, ".variant-ringed": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-primary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-primary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-secondary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-secondary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-tertiary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-tertiary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-success": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-success)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-warning": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-warning)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-error": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-error)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ringed-surface": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-surface)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ".variant-ghost-primary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ":is(.dark .variant-ghost-primary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".variant-ghost-secondary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ":is(.dark .variant-ghost-secondary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ".variant-ghost-tertiary": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ":is(.dark .variant-ghost-tertiary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ".variant-ghost-success": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ":is(.dark .variant-ghost-success)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ".variant-ghost-warning": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ":is(.dark .variant-ghost-warning)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ".variant-ghost-error": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ":is(.dark .variant-ghost-error)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ".variant-ghost,\n	.variant-ghost-surface": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ":is(.dark .variant-ghost),:is(.dark \n	.variant-ghost-surface)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".variant-soft-primary": { "backgroundColor": "rgb(var(--color-primary-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-primary-700))" }, ".dark .variant-soft-primary": { "color": "rgb(var(--color-primary-200))" }, ":is(.dark .variant-soft-primary)": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".variant-soft-secondary": { "backgroundColor": "rgb(var(--color-secondary-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-secondary-700))" }, ".dark .variant-soft-secondary": { "color": "rgb(var(--color-secondary-200))" }, ":is(.dark .variant-soft-secondary)": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ".variant-soft-tertiary": { "backgroundColor": "rgb(var(--color-tertiary-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-tertiary-700))" }, ".dark .variant-soft-tertiary": { "color": "rgb(var(--color-tertiary-200))" }, ":is(.dark .variant-soft-tertiary)": { "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ".variant-soft-success": { "backgroundColor": "rgb(var(--color-success-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-success-700))" }, ".dark .variant-soft-success": { "color": "rgb(var(--color-success-200))" }, ":is(.dark .variant-soft-success)": { "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ".variant-soft-warning": { "backgroundColor": "rgb(var(--color-warning-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-warning-700))" }, ".dark .variant-soft-warning": { "color": "rgb(var(--color-warning-200))" }, ":is(.dark .variant-soft-warning)": { "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ".variant-soft-error": { "backgroundColor": "rgb(var(--color-error-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-error-700))" }, ".dark .variant-soft-error": { "color": "rgb(var(--color-error-200))" }, ":is(.dark .variant-soft-error)": { "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ".variant-soft,\n	.variant-soft-surface": { "backgroundColor": "rgb(var(--color-surface-400) / 0.2)", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "color": "rgb(var(--color-surface-700))" }, ".dark .variant-soft,.dark \n	.variant-soft-surface": { "color": "rgb(var(--color-surface-200))" }, ":is(.dark .variant-soft),:is(.dark \n	.variant-soft-surface)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".variant-glass-primary": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-primary)": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".variant-glass-secondary": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-secondary)": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ".variant-glass-tertiary": { "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-tertiary)": { "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ".variant-glass-success": { "backgroundColor": "rgb(var(--color-success-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-success)": { "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ".variant-glass-warning": { "backgroundColor": "rgb(var(--color-warning-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-warning)": { "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ".variant-glass-error": { "backgroundColor": "rgb(var(--color-error-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-error)": { "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ".variant-glass-surface": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass-surface)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".variant-glass": { "backgroundColor": "rgb(var(--color-surface-50) / 0.3)", "--tw-backdrop-blur": "blur(16px)", "backdropFilter": "var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)" }, ":is(.dark .variant-glass)": { "backgroundColor": "rgb(var(--color-surface-900) / 0.3)" }, ".variant-gradient-primary-secondary": { "--tw-gradient-from": "rgb(var(--color-primary-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-primary-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-secondary-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-primary))" }, ".variant-gradient-secondary-tertiary": { "--tw-gradient-from": "rgb(var(--color-secondary-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-secondary-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-tertiary-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-secondary))" }, ".variant-gradient-tertiary-primary": { "--tw-gradient-from": "rgb(var(--color-tertiary-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-tertiary-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-primary-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-tertiary))" }, ".variant-gradient-secondary-primary": { "--tw-gradient-from": "rgb(var(--color-secondary-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-secondary-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-primary-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-secondary))" }, ".variant-gradient-tertiary-secondary": { "--tw-gradient-from": "rgb(var(--color-tertiary-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-tertiary-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-secondary-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-tertiary))" }, ".variant-gradient-primary-tertiary": { "--tw-gradient-from": "rgb(var(--color-primary-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-primary-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-tertiary-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-primary))" }, ".variant-gradient-success-warning": { "--tw-gradient-from": "rgb(var(--color-success-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-success-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-warning-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-success))" }, ".variant-gradient-warning-error": { "--tw-gradient-from": "rgb(var(--color-warning-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-warning-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-error-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-secondary))" }, ".variant-gradient-error-success": { "--tw-gradient-from": "rgb(var(--color-error-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-error-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-success-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-error))" }, ".variant-gradient-warning-success": { "--tw-gradient-from": "rgb(var(--color-warning-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-warning-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-success-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-warning))" }, ".variant-gradient-error-warning": { "--tw-gradient-from": "rgb(var(--color-error-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-error-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-warning-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-secondary))" }, ".variant-gradient-success-error": { "--tw-gradient-from": "rgb(var(--color-success-500) / 1) var(--tw-gradient-from-position)", "--tw-gradient-to": ["rgb(var(--color-success-500) / 0) var(--tw-gradient-to-position)", "rgb(var(--color-error-500) / 1) var(--tw-gradient-to-position)"], "--tw-gradient-stops": "var(--tw-gradient-from), var(--tw-gradient-to)", "color": "rgb(var(--on-success))" }, "@media (min-width: 768px)": [{ ".h1": { "fontSize": "3rem", "lineHeight": 1 } }, { ".h2": { "fontSize": "2.25rem", "lineHeight": "2.5rem" } }, { ".h3": { "fontSize": "1.5rem", "lineHeight": "2rem" } }, { ".h4": { "fontSize": "1.25rem", "lineHeight": "1.75rem" } }, { ".h5": { "fontSize": "1.125rem", "lineHeight": "1.75rem" } }, { ".h6": { "fontSize": "1rem", "lineHeight": "1.5rem" } }, { ".breadcrumb li": { "display": "block" } }, { ".legend": { "fontSize": "1.5rem", "lineHeight": "2rem" } }, { ".logo-item": { "paddingTop": "2rem", "paddingBottom": "2rem" } }, { ".table tbody td": { "whiteSpace": "normal" } }], "@media (min-width: 1024px)": { ".alert": { "flexDirection": "row", "alignItems": "center" }, ".alert > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0px * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0px * var(--tw-space-y-reverse))", "--tw-space-x-reverse": "0", "marginRight": "calc(1rem * var(--tw-space-x-reverse))", "marginLeft": "calc(1rem * calc(1 - var(--tw-space-x-reverse)))" } } }, base: { ".dark .divider-vertical": { "borderColor": "rgb(var(--color-surface-600))" }, ":is(.dark .anchor)": { "--tw-text-opacity": "1", "color": "rgb(var(--color-primary-500) / var(--tw-text-opacity))" }, ".dark .blockquote": { "color": "rgba(var(--theme-font-color-dark))" }, ".dark .kbd": { "backgroundColor": "rgb(var(--color-surface-600))" }, ":is(.dark .time)": { "--tw-text-opacity": "1", "color": "rgb(var(--color-surface-400) / var(--tw-text-opacity))" }, ":is(.dark .code)": { "--tw-text-opacity": "1", "color": "rgb(var(--color-primary-400) / var(--tw-text-opacity))", "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".dark .alert": { "color": "rgb(var(--color-surface-50))" }, ":is(.dark .badge-glass)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)", "--tw-ring-color": "rgb(250 250 250 / 0.05)" }, ".dark .crumb-separator": { "color": "rgb(var(--color-surface-200))" }, ".dark .card": { "backgroundColor": "rgb(var(--color-surface-800))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-color": "rgb(250 250 250 / 0.05)" }, ".legend": { "fontSize": "1.25rem", "lineHeight": "1.75rem", "fontFamily": "var(--theme-font-family-heading)" }, ".label > :not([hidden]) ~ :not([hidden])": { "--tw-space-y-reverse": "0", "marginTop": "calc(0.25rem * calc(1 - var(--tw-space-y-reverse)))", "marginBottom": "calc(0.25rem * var(--tw-space-y-reverse))" }, ".input,\n	.textarea,\n	.select,\n	.input-group": { "width": "100%", "transitionProperty": "color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "200ms", "backgroundColor": "rgb(var(--color-surface-200))", "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color) !important", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(0px + var(--tw-ring-offset-width)) var(--tw-ring-color) !important", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000) !important", "borderWidth": "var(--theme-border-base)", "borderColor": "rgb(var(--color-surface-400))" }, ".dark .input,.dark \n	.textarea,.dark \n	.select,.dark \n	.input-group": { "backgroundColor": "rgb(var(--color-surface-700))", "borderColor": "rgb(var(--color-surface-500))" }, ".input:hover,\n	.textarea:hover,\n	.select:hover,\n	.input-group:hover": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".input:focus,\n	.textarea:focus,\n	.select:focus,\n	.input-group:focus": { "--tw-brightness": "brightness(1.05)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".input:focus-within,\n	.textarea:focus-within,\n	.select:focus-within,\n	.input-group:focus-within": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-primary-500) / var(--tw-border-opacity))" }, ".input,\n	.input-group": { "borderRadius": "var(--theme-rounded-base)" }, ".dark .select option": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".dark .checkbox,.dark \n	.radio": { "backgroundColor": "rgb(var(--color-surface-700))", "borderColor": "rgb(var(--color-surface-500))" }, ".input[type='file']": { "padding": "0.25rem" }, ".input[type='color']": { "height": "2.5rem", "width": "2.5rem", "cursor": "pointer", "overflow": "hidden", "borderStyle": "none", "borderRadius": "var(--theme-rounded-base)", "WebkitAppearance": "none" }, ".input[type='color']::-webkit-color-swatch-wrapper": { "padding": "0px" }, ".input[type='color']::-webkit-color-swatch": { "borderStyle": "none" }, ".input[type='color']::-webkit-color-swatch:hover": { "--tw-brightness": "brightness(1.1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".input[type='color']::-moz-color-swatch": { "borderStyle": "none" }, ".input:disabled,\n	.textarea:disabled,\n	.select:disabled,\n	.input-group > input:disabled,\n	.input-group > textarea:disabled,\n	.input-group > select:disabled": { "cursor": "not-allowed !important", "opacity": "0.5 !important" }, ".input:disabled:hover,\n	.textarea:disabled:hover,\n	.select:disabled:hover,\n	.input-group > input:disabled:hover,\n	.input-group > textarea:disabled:hover,\n	.input-group > select:disabled:hover": { "--tw-brightness": "brightness(1) !important", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow) !important" }, ".input[readonly],\n	.textarea[readonly],\n	.select[readonly]": { "cursor": "not-allowed !important", "borderColor": "transparent !important" }, ".input[readonly]:hover,\n	.textarea[readonly]:hover,\n	.select[readonly]:hover": { "--tw-brightness": "brightness(1) !important", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow) !important" }, ".dark .input-group select option": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".dark .input-group-divider input,.dark \n	.input-group-divider select,.dark \n	.input-group-divider div,.dark \n	.input-group-divider a": { "borderColor": "rgb(var(--color-surface-500))" }, ".dark .input-group-divider input:focus,.dark \n	.input-group-divider select:focus,.dark \n	.input-group-divider div:focus,.dark \n	.input-group-divider a:focus": { "borderColor": "rgb(var(--color-surface-500))" }, ".dark .input-group-shim": { "color": "rgb(var(--color-surface-300))" }, ":is(.dark .input-success)": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-success-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-success-700) / var(--tw-text-opacity))" }, ":is(.dark .input-warning)": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-warning-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-warning-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-warning-700) / var(--tw-text-opacity))" }, ":is(.dark .input-error)": { "--tw-border-opacity": "1", "borderColor": "rgb(var(--color-error-500) / var(--tw-border-opacity))", "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-200) / var(--tw-bg-opacity))", "--tw-text-opacity": "1", "color": "rgb(var(--color-error-500) / var(--tw-text-opacity))" }, ":is(.dark .variant-form-material)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".dark .list-nav a:hover,.dark \n	.list-nav button:hover,.dark \n	.list-option:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .list-nav a:focus,.dark \n	.list-nav button:focus,.dark \n	.list-option:focus": { "color": "rgb(var(--color-primary-200))" }, ":is(.dark .list-nav a:focus),:is(.dark \n	.list-nav button:focus),:is(.dark \n	.list-option:focus)": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".dark .logo-item": { "backgroundColor": "rgb(var(--color-surface-800))" }, ":is(.dark .logo-item)": { "--tw-text-opacity": "1", "color": "rgb(255 255 255 / var(--tw-text-opacity))" }, ".dark .placeholder": { "backgroundColor": "rgb(var(--color-surface-600))" }, ".dark .placeholder-circle": { "backgroundColor": "rgb(var(--color-surface-600))" }, ".dark .table": { "backgroundColor": "rgb(var(--color-surface-800))" }, ".dark .table-interactive tbody tr:hover:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .table-interactive tbody tr:hover:nth-child(even):hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .table thead": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".dark .table tfoot": { "backgroundColor": "rgb(var(--color-surface-800))" }, ":is(.dark .variant-outline-primary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-secondary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-tertiary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-success)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-warning)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline-error)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))" }, ":is(.dark .variant-outline),:is(.dark \n	.variant-outline-surface)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))" }, ".dark .variant-filled": { "backgroundColor": "rgb(var(--color-surface-50))", "color": "rgb(var(--color-surface-900))" }, ":is(.dark .variant-filled-primary)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-primary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-primary))" }, ":is(.dark .variant-filled-secondary)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-secondary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-secondary))" }, ":is(.dark .variant-filled-tertiary)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-tertiary-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-tertiary))" }, ":is(.dark .variant-filled-success)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-success-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-success))" }, ":is(.dark .variant-filled-warning)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-warning-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-warning))" }, ":is(.dark .variant-filled-error)": { "--tw-bg-opacity": "1", "backgroundColor": "rgb(var(--color-error-500) / var(--tw-bg-opacity))", "color": "rgb(var(--on-error))" }, ".dark .variant-filled-surface": { "backgroundColor": "rgb(var(--color-surface-500))" }, ":is(.dark .variant-filled-surface)": { "color": "rgb(var(--on-surface))" }, ":is(.dark .variant-ringed)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-primary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-secondary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-tertiary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-success)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-warning)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-error)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ringed-surface)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "transparent" }, ":is(.dark .variant-ghost-primary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-primary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ":is(.dark .variant-ghost-secondary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-secondary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ":is(.dark .variant-ghost-tertiary)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-tertiary-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ":is(.dark .variant-ghost-success)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-success-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ":is(.dark .variant-ghost-warning)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-warning-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ":is(.dark .variant-ghost-error)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-error-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ":is(.dark .variant-ghost),:is(.dark \n	.variant-ghost-surface)": { "--tw-ring-opacity": "1", "--tw-ring-color": "rgb(var(--color-surface-500) / var(--tw-ring-opacity))", "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ".dark .variant-soft-primary": { "color": "rgb(var(--color-primary-200))" }, ":is(.dark .variant-soft-primary)": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ".dark .variant-soft-secondary": { "color": "rgb(var(--color-secondary-200))" }, ":is(.dark .variant-soft-secondary)": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ".dark .variant-soft-tertiary": { "color": "rgb(var(--color-tertiary-200))" }, ":is(.dark .variant-soft-tertiary)": { "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ".dark .variant-soft-success": { "color": "rgb(var(--color-success-200))" }, ":is(.dark .variant-soft-success)": { "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ".dark .variant-soft-warning": { "color": "rgb(var(--color-warning-200))" }, ":is(.dark .variant-soft-warning)": { "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ".dark .variant-soft-error": { "color": "rgb(var(--color-error-200))" }, ":is(.dark .variant-soft-error)": { "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ".dark .variant-soft,.dark \n	.variant-soft-surface": { "color": "rgb(var(--color-surface-200))" }, ":is(.dark .variant-soft),:is(.dark \n	.variant-soft-surface)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ":is(.dark .variant-glass-primary)": { "backgroundColor": "rgb(var(--color-primary-500) / 0.2)" }, ":is(.dark .variant-glass-secondary)": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.2)" }, ":is(.dark .variant-glass-tertiary)": { "backgroundColor": "rgb(var(--color-tertiary-500) / 0.2)" }, ":is(.dark .variant-glass-success)": { "backgroundColor": "rgb(var(--color-success-500) / 0.2)" }, ":is(.dark .variant-glass-warning)": { "backgroundColor": "rgb(var(--color-warning-500) / 0.2)" }, ":is(.dark .variant-glass-error)": { "backgroundColor": "rgb(var(--color-error-500) / 0.2)" }, ":is(.dark .variant-glass-surface)": { "backgroundColor": "rgb(var(--color-surface-500) / 0.2)" }, ":is(.dark .variant-glass)": { "backgroundColor": "rgb(var(--color-surface-900) / 0.3)" }, "@media (min-width: 768px)": { ".legend": { "fontSize": "1.5rem", "lineHeight": "2rem" } }, ".dark .bg-primary-backdrop-token": { "backgroundColor": "rgb(var(--color-primary-900) / 0.7)" }, ".dark .bg-primary-hover-token:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark .bg-primary-50-900-token": { "backgroundColor": "rgb(var(--color-primary-900))" }, ".dark .bg-primary-100-800-token": { "backgroundColor": "rgb(var(--color-primary-800))" }, ".dark .bg-primary-200-700-token": { "backgroundColor": "rgb(var(--color-primary-700))" }, ".dark .bg-primary-300-600-token": { "backgroundColor": "rgb(var(--color-primary-600))" }, ".dark .bg-primary-400-500-token": { "backgroundColor": "rgb(var(--color-primary-500))" }, ".dark .bg-primary-900-50-token": { "backgroundColor": "rgb(var(--color-primary-50))" }, ".dark .bg-primary-800-100-token": { "backgroundColor": "rgb(var(--color-primary-100))" }, ".dark .bg-primary-700-200-token": { "backgroundColor": "rgb(var(--color-primary-200))" }, ".dark .bg-primary-600-300-token": { "backgroundColor": "rgb(var(--color-primary-300))" }, ".dark .bg-primary-500-400-token": { "backgroundColor": "rgb(var(--color-primary-400))" }, ".dark .bg-secondary-backdrop-token": { "backgroundColor": "rgb(var(--color-secondary-900) / 0.7)" }, ".dark .bg-secondary-hover-token:hover": { "backgroundColor": "rgb(var(--color-secondary-500) / 0.1)" }, ".dark .bg-secondary-50-900-token": { "backgroundColor": "rgb(var(--color-secondary-900))" }, ".dark .bg-secondary-100-800-token": { "backgroundColor": "rgb(var(--color-secondary-800))" }, ".dark .bg-secondary-200-700-token": { "backgroundColor": "rgb(var(--color-secondary-700))" }, ".dark .bg-secondary-300-600-token": { "backgroundColor": "rgb(var(--color-secondary-600))" }, ".dark .bg-secondary-400-500-token": { "backgroundColor": "rgb(var(--color-secondary-500))" }, ".dark .bg-secondary-900-50-token": { "backgroundColor": "rgb(var(--color-secondary-50))" }, ".dark .bg-secondary-800-100-token": { "backgroundColor": "rgb(var(--color-secondary-100))" }, ".dark .bg-secondary-700-200-token": { "backgroundColor": "rgb(var(--color-secondary-200))" }, ".dark .bg-secondary-600-300-token": { "backgroundColor": "rgb(var(--color-secondary-300))" }, ".dark .bg-secondary-500-400-token": { "backgroundColor": "rgb(var(--color-secondary-400))" }, ".dark .bg-tertiary-backdrop-token": { "backgroundColor": "rgb(var(--color-tertiary-900) / 0.7)" }, ".dark .bg-tertiary-hover-token:hover": { "backgroundColor": "rgb(var(--color-tertiary-500) / 0.1)" }, ".dark .bg-tertiary-50-900-token": { "backgroundColor": "rgb(var(--color-tertiary-900))" }, ".dark .bg-tertiary-100-800-token": { "backgroundColor": "rgb(var(--color-tertiary-800))" }, ".dark .bg-tertiary-200-700-token": { "backgroundColor": "rgb(var(--color-tertiary-700))" }, ".dark .bg-tertiary-300-600-token": { "backgroundColor": "rgb(var(--color-tertiary-600))" }, ".dark .bg-tertiary-400-500-token": { "backgroundColor": "rgb(var(--color-tertiary-500))" }, ".dark .bg-tertiary-900-50-token": { "backgroundColor": "rgb(var(--color-tertiary-50))" }, ".dark .bg-tertiary-800-100-token": { "backgroundColor": "rgb(var(--color-tertiary-100))" }, ".dark .bg-tertiary-700-200-token": { "backgroundColor": "rgb(var(--color-tertiary-200))" }, ".dark .bg-tertiary-600-300-token": { "backgroundColor": "rgb(var(--color-tertiary-300))" }, ".dark .bg-tertiary-500-400-token": { "backgroundColor": "rgb(var(--color-tertiary-400))" }, ".dark .bg-success-backdrop-token": { "backgroundColor": "rgb(var(--color-success-900) / 0.7)" }, ".dark .bg-success-hover-token:hover": { "backgroundColor": "rgb(var(--color-success-500) / 0.1)" }, ".dark .bg-success-50-900-token": { "backgroundColor": "rgb(var(--color-success-900))" }, ".dark .bg-success-100-800-token": { "backgroundColor": "rgb(var(--color-success-800))" }, ".dark .bg-success-200-700-token": { "backgroundColor": "rgb(var(--color-success-700))" }, ".dark .bg-success-300-600-token": { "backgroundColor": "rgb(var(--color-success-600))" }, ".dark .bg-success-400-500-token": { "backgroundColor": "rgb(var(--color-success-500))" }, ".dark .bg-success-900-50-token": { "backgroundColor": "rgb(var(--color-success-50))" }, ".dark .bg-success-800-100-token": { "backgroundColor": "rgb(var(--color-success-100))" }, ".dark .bg-success-700-200-token": { "backgroundColor": "rgb(var(--color-success-200))" }, ".dark .bg-success-600-300-token": { "backgroundColor": "rgb(var(--color-success-300))" }, ".dark .bg-success-500-400-token": { "backgroundColor": "rgb(var(--color-success-400))" }, ".dark .bg-warning-backdrop-token": { "backgroundColor": "rgb(var(--color-warning-900) / 0.7)" }, ".dark .bg-warning-hover-token:hover": { "backgroundColor": "rgb(var(--color-warning-500) / 0.1)" }, ".dark .bg-warning-50-900-token": { "backgroundColor": "rgb(var(--color-warning-900))" }, ".dark .bg-warning-100-800-token": { "backgroundColor": "rgb(var(--color-warning-800))" }, ".dark .bg-warning-200-700-token": { "backgroundColor": "rgb(var(--color-warning-700))" }, ".dark .bg-warning-300-600-token": { "backgroundColor": "rgb(var(--color-warning-600))" }, ".dark .bg-warning-400-500-token": { "backgroundColor": "rgb(var(--color-warning-500))" }, ".dark .bg-warning-900-50-token": { "backgroundColor": "rgb(var(--color-warning-50))" }, ".dark .bg-warning-800-100-token": { "backgroundColor": "rgb(var(--color-warning-100))" }, ".dark .bg-warning-700-200-token": { "backgroundColor": "rgb(var(--color-warning-200))" }, ".dark .bg-warning-600-300-token": { "backgroundColor": "rgb(var(--color-warning-300))" }, ".dark .bg-warning-500-400-token": { "backgroundColor": "rgb(var(--color-warning-400))" }, ".dark .bg-error-backdrop-token": { "backgroundColor": "rgb(var(--color-error-900) / 0.7)" }, ".dark .bg-error-hover-token:hover": { "backgroundColor": "rgb(var(--color-error-500) / 0.1)" }, ".dark .bg-error-50-900-token": { "backgroundColor": "rgb(var(--color-error-900))" }, ".dark .bg-error-100-800-token": { "backgroundColor": "rgb(var(--color-error-800))" }, ".dark .bg-error-200-700-token": { "backgroundColor": "rgb(var(--color-error-700))" }, ".dark .bg-error-300-600-token": { "backgroundColor": "rgb(var(--color-error-600))" }, ".dark .bg-error-400-500-token": { "backgroundColor": "rgb(var(--color-error-500))" }, ".dark .bg-error-900-50-token": { "backgroundColor": "rgb(var(--color-error-50))" }, ".dark .bg-error-800-100-token": { "backgroundColor": "rgb(var(--color-error-100))" }, ".dark .bg-error-700-200-token": { "backgroundColor": "rgb(var(--color-error-200))" }, ".dark .bg-error-600-300-token": { "backgroundColor": "rgb(var(--color-error-300))" }, ".dark .bg-error-500-400-token": { "backgroundColor": "rgb(var(--color-error-400))" }, ".dark .bg-surface-backdrop-token": { "backgroundColor": "rgb(var(--color-surface-900) / 0.7)" }, ".dark .bg-surface-hover-token:hover": { "backgroundColor": "rgb(var(--color-surface-500) / 0.1)" }, ".dark .bg-surface-50-900-token": { "backgroundColor": "rgb(var(--color-surface-900))" }, ".dark .bg-surface-100-800-token": { "backgroundColor": "rgb(var(--color-surface-800))" }, ".dark .bg-surface-200-700-token": { "backgroundColor": "rgb(var(--color-surface-700))" }, ".dark .bg-surface-300-600-token": { "backgroundColor": "rgb(var(--color-surface-600))" }, ".dark .bg-surface-400-500-token": { "backgroundColor": "rgb(var(--color-surface-500))" }, ".dark .bg-surface-900-50-token": { "backgroundColor": "rgb(var(--color-surface-50))" }, ".dark .bg-surface-800-100-token": { "backgroundColor": "rgb(var(--color-surface-100))" }, ".dark .bg-surface-700-200-token": { "backgroundColor": "rgb(var(--color-surface-200))" }, ".dark .bg-surface-600-300-token": { "backgroundColor": "rgb(var(--color-surface-300))" }, ".dark .bg-surface-500-400-token": { "backgroundColor": "rgb(var(--color-surface-400))" }, ".dark .border-primary-50-900-token": { "borderColor": "rgb(var(--color-primary-900))" }, ".dark .divide-primary-50-900-token": { "borderColor": "rgb(var(--color-primary-900))" }, ".dark .border-primary-100-800-token": { "borderColor": "rgb(var(--color-primary-800))" }, ".dark .divide-primary-100-800-token": { "borderColor": "rgb(var(--color-primary-800))" }, ".dark .border-primary-200-700-token": { "borderColor": "rgb(var(--color-primary-700))" }, ".dark .divide-primary-200-700-token": { "borderColor": "rgb(var(--color-primary-700))" }, ".dark .border-primary-300-600-token": { "borderColor": "rgb(var(--color-primary-600))" }, ".dark .divide-primary-300-600-token": { "borderColor": "rgb(var(--color-primary-600))" }, ".dark .border-primary-400-500-token": { "borderColor": "rgb(var(--color-primary-500))" }, ".dark .divide-primary-400-500-token": { "borderColor": "rgb(var(--color-primary-500))" }, ".dark .border-primary-900-50-token": { "borderColor": "rgb(var(--color-primary-50))" }, ".dark .divide-primary-900-50-token": { "borderColor": "rgb(var(--color-primary-50))" }, ".dark .border-primary-800-100-token": { "borderColor": "rgb(var(--color-primary-100))" }, ".dark .divide-primary-800-100-token": { "borderColor": "rgb(var(--color-primary-100))" }, ".dark .border-primary-700-200-token": { "borderColor": "rgb(var(--color-primary-200))" }, ".dark .divide-primary-700-200-token": { "borderColor": "rgb(var(--color-primary-200))" }, ".dark .border-primary-600-300-token": { "borderColor": "rgb(var(--color-primary-300))" }, ".dark .divide-primary-600-300-token": { "borderColor": "rgb(var(--color-primary-300))" }, ".dark .border-primary-500-400-token": { "borderColor": "rgb(var(--color-primary-400))" }, ".dark .divide-primary-500-400-token": { "borderColor": "rgb(var(--color-primary-400))" }, ".dark .border-secondary-50-900-token": { "borderColor": "rgb(var(--color-secondary-900))" }, ".dark .divide-secondary-50-900-token": { "borderColor": "rgb(var(--color-secondary-900))" }, ".dark .border-secondary-100-800-token": { "borderColor": "rgb(var(--color-secondary-800))" }, ".dark .divide-secondary-100-800-token": { "borderColor": "rgb(var(--color-secondary-800))" }, ".dark .border-secondary-200-700-token": { "borderColor": "rgb(var(--color-secondary-700))" }, ".dark .divide-secondary-200-700-token": { "borderColor": "rgb(var(--color-secondary-700))" }, ".dark .border-secondary-300-600-token": { "borderColor": "rgb(var(--color-secondary-600))" }, ".dark .divide-secondary-300-600-token": { "borderColor": "rgb(var(--color-secondary-600))" }, ".dark .border-secondary-400-500-token": { "borderColor": "rgb(var(--color-secondary-500))" }, ".dark .divide-secondary-400-500-token": { "borderColor": "rgb(var(--color-secondary-500))" }, ".dark .border-secondary-900-50-token": { "borderColor": "rgb(var(--color-secondary-50))" }, ".dark .divide-secondary-900-50-token": { "borderColor": "rgb(var(--color-secondary-50))" }, ".dark .border-secondary-800-100-token": { "borderColor": "rgb(var(--color-secondary-100))" }, ".dark .divide-secondary-800-100-token": { "borderColor": "rgb(var(--color-secondary-100))" }, ".dark .border-secondary-700-200-token": { "borderColor": "rgb(var(--color-secondary-200))" }, ".dark .divide-secondary-700-200-token": { "borderColor": "rgb(var(--color-secondary-200))" }, ".dark .border-secondary-600-300-token": { "borderColor": "rgb(var(--color-secondary-300))" }, ".dark .divide-secondary-600-300-token": { "borderColor": "rgb(var(--color-secondary-300))" }, ".dark .border-secondary-500-400-token": { "borderColor": "rgb(var(--color-secondary-400))" }, ".dark .divide-secondary-500-400-token": { "borderColor": "rgb(var(--color-secondary-400))" }, ".dark .border-tertiary-50-900-token": { "borderColor": "rgb(var(--color-tertiary-900))" }, ".dark .divide-tertiary-50-900-token": { "borderColor": "rgb(var(--color-tertiary-900))" }, ".dark .border-tertiary-100-800-token": { "borderColor": "rgb(var(--color-tertiary-800))" }, ".dark .divide-tertiary-100-800-token": { "borderColor": "rgb(var(--color-tertiary-800))" }, ".dark .border-tertiary-200-700-token": { "borderColor": "rgb(var(--color-tertiary-700))" }, ".dark .divide-tertiary-200-700-token": { "borderColor": "rgb(var(--color-tertiary-700))" }, ".dark .border-tertiary-300-600-token": { "borderColor": "rgb(var(--color-tertiary-600))" }, ".dark .divide-tertiary-300-600-token": { "borderColor": "rgb(var(--color-tertiary-600))" }, ".dark .border-tertiary-400-500-token": { "borderColor": "rgb(var(--color-tertiary-500))" }, ".dark .divide-tertiary-400-500-token": { "borderColor": "rgb(var(--color-tertiary-500))" }, ".dark .border-tertiary-900-50-token": { "borderColor": "rgb(var(--color-tertiary-50))" }, ".dark .divide-tertiary-900-50-token": { "borderColor": "rgb(var(--color-tertiary-50))" }, ".dark .border-tertiary-800-100-token": { "borderColor": "rgb(var(--color-tertiary-100))" }, ".dark .divide-tertiary-800-100-token": { "borderColor": "rgb(var(--color-tertiary-100))" }, ".dark .border-tertiary-700-200-token": { "borderColor": "rgb(var(--color-tertiary-200))" }, ".dark .divide-tertiary-700-200-token": { "borderColor": "rgb(var(--color-tertiary-200))" }, ".dark .border-tertiary-600-300-token": { "borderColor": "rgb(var(--color-tertiary-300))" }, ".dark .divide-tertiary-600-300-token": { "borderColor": "rgb(var(--color-tertiary-300))" }, ".dark .border-tertiary-500-400-token": { "borderColor": "rgb(var(--color-tertiary-400))" }, ".dark .divide-tertiary-500-400-token": { "borderColor": "rgb(var(--color-tertiary-400))" }, ".dark .border-success-50-900-token": { "borderColor": "rgb(var(--color-success-900))" }, ".dark .divide-success-50-900-token": { "borderColor": "rgb(var(--color-success-900))" }, ".dark .border-success-100-800-token": { "borderColor": "rgb(var(--color-success-800))" }, ".dark .divide-success-100-800-token": { "borderColor": "rgb(var(--color-success-800))" }, ".dark .border-success-200-700-token": { "borderColor": "rgb(var(--color-success-700))" }, ".dark .divide-success-200-700-token": { "borderColor": "rgb(var(--color-success-700))" }, ".dark .border-success-300-600-token": { "borderColor": "rgb(var(--color-success-600))" }, ".dark .divide-success-300-600-token": { "borderColor": "rgb(var(--color-success-600))" }, ".dark .border-success-400-500-token": { "borderColor": "rgb(var(--color-success-500))" }, ".dark .divide-success-400-500-token": { "borderColor": "rgb(var(--color-success-500))" }, ".dark .border-success-900-50-token": { "borderColor": "rgb(var(--color-success-50))" }, ".dark .divide-success-900-50-token": { "borderColor": "rgb(var(--color-success-50))" }, ".dark .border-success-800-100-token": { "borderColor": "rgb(var(--color-success-100))" }, ".dark .divide-success-800-100-token": { "borderColor": "rgb(var(--color-success-100))" }, ".dark .border-success-700-200-token": { "borderColor": "rgb(var(--color-success-200))" }, ".dark .divide-success-700-200-token": { "borderColor": "rgb(var(--color-success-200))" }, ".dark .border-success-600-300-token": { "borderColor": "rgb(var(--color-success-300))" }, ".dark .divide-success-600-300-token": { "borderColor": "rgb(var(--color-success-300))" }, ".dark .border-success-500-400-token": { "borderColor": "rgb(var(--color-success-400))" }, ".dark .divide-success-500-400-token": { "borderColor": "rgb(var(--color-success-400))" }, ".dark .border-warning-50-900-token": { "borderColor": "rgb(var(--color-warning-900))" }, ".dark .divide-warning-50-900-token": { "borderColor": "rgb(var(--color-warning-900))" }, ".dark .border-warning-100-800-token": { "borderColor": "rgb(var(--color-warning-800))" }, ".dark .divide-warning-100-800-token": { "borderColor": "rgb(var(--color-warning-800))" }, ".dark .border-warning-200-700-token": { "borderColor": "rgb(var(--color-warning-700))" }, ".dark .divide-warning-200-700-token": { "borderColor": "rgb(var(--color-warning-700))" }, ".dark .border-warning-300-600-token": { "borderColor": "rgb(var(--color-warning-600))" }, ".dark .divide-warning-300-600-token": { "borderColor": "rgb(var(--color-warning-600))" }, ".dark .border-warning-400-500-token": { "borderColor": "rgb(var(--color-warning-500))" }, ".dark .divide-warning-400-500-token": { "borderColor": "rgb(var(--color-warning-500))" }, ".dark .border-warning-900-50-token": { "borderColor": "rgb(var(--color-warning-50))" }, ".dark .divide-warning-900-50-token": { "borderColor": "rgb(var(--color-warning-50))" }, ".dark .border-warning-800-100-token": { "borderColor": "rgb(var(--color-warning-100))" }, ".dark .divide-warning-800-100-token": { "borderColor": "rgb(var(--color-warning-100))" }, ".dark .border-warning-700-200-token": { "borderColor": "rgb(var(--color-warning-200))" }, ".dark .divide-warning-700-200-token": { "borderColor": "rgb(var(--color-warning-200))" }, ".dark .border-warning-600-300-token": { "borderColor": "rgb(var(--color-warning-300))" }, ".dark .divide-warning-600-300-token": { "borderColor": "rgb(var(--color-warning-300))" }, ".dark .border-warning-500-400-token": { "borderColor": "rgb(var(--color-warning-400))" }, ".dark .divide-warning-500-400-token": { "borderColor": "rgb(var(--color-warning-400))" }, ".dark .border-error-50-900-token": { "borderColor": "rgb(var(--color-error-900))" }, ".dark .divide-error-50-900-token": { "borderColor": "rgb(var(--color-error-900))" }, ".dark .border-error-100-800-token": { "borderColor": "rgb(var(--color-error-800))" }, ".dark .divide-error-100-800-token": { "borderColor": "rgb(var(--color-error-800))" }, ".dark .border-error-200-700-token": { "borderColor": "rgb(var(--color-error-700))" }, ".dark .divide-error-200-700-token": { "borderColor": "rgb(var(--color-error-700))" }, ".dark .border-error-300-600-token": { "borderColor": "rgb(var(--color-error-600))" }, ".dark .divide-error-300-600-token": { "borderColor": "rgb(var(--color-error-600))" }, ".dark .border-error-400-500-token": { "borderColor": "rgb(var(--color-error-500))" }, ".dark .divide-error-400-500-token": { "borderColor": "rgb(var(--color-error-500))" }, ".dark .border-error-900-50-token": { "borderColor": "rgb(var(--color-error-50))" }, ".dark .divide-error-900-50-token": { "borderColor": "rgb(var(--color-error-50))" }, ".dark .border-error-800-100-token": { "borderColor": "rgb(var(--color-error-100))" }, ".dark .divide-error-800-100-token": { "borderColor": "rgb(var(--color-error-100))" }, ".dark .border-error-700-200-token": { "borderColor": "rgb(var(--color-error-200))" }, ".dark .divide-error-700-200-token": { "borderColor": "rgb(var(--color-error-200))" }, ".dark .border-error-600-300-token": { "borderColor": "rgb(var(--color-error-300))" }, ".dark .divide-error-600-300-token": { "borderColor": "rgb(var(--color-error-300))" }, ".dark .border-error-500-400-token": { "borderColor": "rgb(var(--color-error-400))" }, ".dark .divide-error-500-400-token": { "borderColor": "rgb(var(--color-error-400))" }, ".dark .border-surface-50-900-token": { "borderColor": "rgb(var(--color-surface-900))" }, ".dark .divide-surface-50-900-token": { "borderColor": "rgb(var(--color-surface-900))" }, ".dark .border-surface-100-800-token": { "borderColor": "rgb(var(--color-surface-800))" }, ".dark .divide-surface-100-800-token": { "borderColor": "rgb(var(--color-surface-800))" }, ".dark .border-surface-200-700-token": { "borderColor": "rgb(var(--color-surface-700))" }, ".dark .divide-surface-200-700-token": { "borderColor": "rgb(var(--color-surface-700))" }, ".dark .border-surface-300-600-token": { "borderColor": "rgb(var(--color-surface-600))" }, ".dark .divide-surface-300-600-token": { "borderColor": "rgb(var(--color-surface-600))" }, ".dark .border-surface-400-500-token": { "borderColor": "rgb(var(--color-surface-500))" }, ".dark .divide-surface-400-500-token": { "borderColor": "rgb(var(--color-surface-500))" }, ".dark .border-surface-900-50-token": { "borderColor": "rgb(var(--color-surface-50))" }, ".dark .divide-surface-900-50-token": { "borderColor": "rgb(var(--color-surface-50))" }, ".dark .border-surface-800-100-token": { "borderColor": "rgb(var(--color-surface-100))" }, ".dark .divide-surface-800-100-token": { "borderColor": "rgb(var(--color-surface-100))" }, ".dark .border-surface-700-200-token": { "borderColor": "rgb(var(--color-surface-200))" }, ".dark .divide-surface-700-200-token": { "borderColor": "rgb(var(--color-surface-200))" }, ".dark .border-surface-600-300-token": { "borderColor": "rgb(var(--color-surface-300))" }, ".dark .divide-surface-600-300-token": { "borderColor": "rgb(var(--color-surface-300))" }, ".dark .border-surface-500-400-token": { "borderColor": "rgb(var(--color-surface-400))" }, ".dark .divide-surface-500-400-token": { "borderColor": "rgb(var(--color-surface-400))" }, ".dark .fill-token": { "fill": "rgba(var(--theme-font-color-dark))" }, ".dark .text-token": { "color": "rgba(var(--theme-font-color-dark))" }, ".dark .text-primary-50-900-token": { "color": "rgb(var(--color-primary-900))" }, ".dark .decoration-primary-50-900-token": { "textDecorationColor": "rgb(var(--color-primary-900))" }, ".dark .text-primary-100-800-token": { "color": "rgb(var(--color-primary-800))" }, ".dark .decoration-primary-100-800-token": { "textDecorationColor": "rgb(var(--color-primary-800))" }, ".dark .text-primary-200-700-token": { "color": "rgb(var(--color-primary-700))" }, ".dark .decoration-primary-200-700-token": { "textDecorationColor": "rgb(var(--color-primary-700))" }, ".dark .text-primary-300-600-token": { "color": "rgb(var(--color-primary-600))" }, ".dark .decoration-primary-300-600-token": { "textDecorationColor": "rgb(var(--color-primary-600))" }, ".dark .text-primary-400-500-token": { "color": "rgb(var(--color-primary-500))" }, ".dark .decoration-primary-400-500-token": { "textDecorationColor": "rgb(var(--color-primary-500))" }, ".dark .text-primary-900-50-token": { "color": "rgb(var(--color-primary-50))" }, ".dark .decoration-primary-900-50-token": { "textDecorationColor": "rgb(var(--color-primary-50))" }, ".dark .text-primary-800-100-token": { "color": "rgb(var(--color-primary-100))" }, ".dark .decoration-primary-800-100-token": { "textDecorationColor": "rgb(var(--color-primary-100))" }, ".dark .text-primary-700-200-token": { "color": "rgb(var(--color-primary-200))" }, ".dark .decoration-primary-700-200-token": { "textDecorationColor": "rgb(var(--color-primary-200))" }, ".dark .text-primary-600-300-token": { "color": "rgb(var(--color-primary-300))" }, ".dark .decoration-primary-600-300-token": { "textDecorationColor": "rgb(var(--color-primary-300))" }, ".dark .text-primary-500-400-token": { "color": "rgb(var(--color-primary-400))" }, ".dark .decoration-primary-500-400-token": { "textDecorationColor": "rgb(var(--color-primary-400))" }, ".dark .text-secondary-50-900-token": { "color": "rgb(var(--color-secondary-900))" }, ".dark .decoration-secondary-50-900-token": { "textDecorationColor": "rgb(var(--color-secondary-900))" }, ".dark .text-secondary-100-800-token": { "color": "rgb(var(--color-secondary-800))" }, ".dark .decoration-secondary-100-800-token": { "textDecorationColor": "rgb(var(--color-secondary-800))" }, ".dark .text-secondary-200-700-token": { "color": "rgb(var(--color-secondary-700))" }, ".dark .decoration-secondary-200-700-token": { "textDecorationColor": "rgb(var(--color-secondary-700))" }, ".dark .text-secondary-300-600-token": { "color": "rgb(var(--color-secondary-600))" }, ".dark .decoration-secondary-300-600-token": { "textDecorationColor": "rgb(var(--color-secondary-600))" }, ".dark .text-secondary-400-500-token": { "color": "rgb(var(--color-secondary-500))" }, ".dark .decoration-secondary-400-500-token": { "textDecorationColor": "rgb(var(--color-secondary-500))" }, ".dark .text-secondary-900-50-token": { "color": "rgb(var(--color-secondary-50))" }, ".dark .decoration-secondary-900-50-token": { "textDecorationColor": "rgb(var(--color-secondary-50))" }, ".dark .text-secondary-800-100-token": { "color": "rgb(var(--color-secondary-100))" }, ".dark .decoration-secondary-800-100-token": { "textDecorationColor": "rgb(var(--color-secondary-100))" }, ".dark .text-secondary-700-200-token": { "color": "rgb(var(--color-secondary-200))" }, ".dark .decoration-secondary-700-200-token": { "textDecorationColor": "rgb(var(--color-secondary-200))" }, ".dark .text-secondary-600-300-token": { "color": "rgb(var(--color-secondary-300))" }, ".dark .decoration-secondary-600-300-token": { "textDecorationColor": "rgb(var(--color-secondary-300))" }, ".dark .text-secondary-500-400-token": { "color": "rgb(var(--color-secondary-400))" }, ".dark .decoration-secondary-500-400-token": { "textDecorationColor": "rgb(var(--color-secondary-400))" }, ".dark .text-tertiary-50-900-token": { "color": "rgb(var(--color-tertiary-900))" }, ".dark .decoration-tertiary-50-900-token": { "textDecorationColor": "rgb(var(--color-tertiary-900))" }, ".dark .text-tertiary-100-800-token": { "color": "rgb(var(--color-tertiary-800))" }, ".dark .decoration-tertiary-100-800-token": { "textDecorationColor": "rgb(var(--color-tertiary-800))" }, ".dark .text-tertiary-200-700-token": { "color": "rgb(var(--color-tertiary-700))" }, ".dark .decoration-tertiary-200-700-token": { "textDecorationColor": "rgb(var(--color-tertiary-700))" }, ".dark .text-tertiary-300-600-token": { "color": "rgb(var(--color-tertiary-600))" }, ".dark .decoration-tertiary-300-600-token": { "textDecorationColor": "rgb(var(--color-tertiary-600))" }, ".dark .text-tertiary-400-500-token": { "color": "rgb(var(--color-tertiary-500))" }, ".dark .decoration-tertiary-400-500-token": { "textDecorationColor": "rgb(var(--color-tertiary-500))" }, ".dark .text-tertiary-900-50-token": { "color": "rgb(var(--color-tertiary-50))" }, ".dark .decoration-tertiary-900-50-token": { "textDecorationColor": "rgb(var(--color-tertiary-50))" }, ".dark .text-tertiary-800-100-token": { "color": "rgb(var(--color-tertiary-100))" }, ".dark .decoration-tertiary-800-100-token": { "textDecorationColor": "rgb(var(--color-tertiary-100))" }, ".dark .text-tertiary-700-200-token": { "color": "rgb(var(--color-tertiary-200))" }, ".dark .decoration-tertiary-700-200-token": { "textDecorationColor": "rgb(var(--color-tertiary-200))" }, ".dark .text-tertiary-600-300-token": { "color": "rgb(var(--color-tertiary-300))" }, ".dark .decoration-tertiary-600-300-token": { "textDecorationColor": "rgb(var(--color-tertiary-300))" }, ".dark .text-tertiary-500-400-token": { "color": "rgb(var(--color-tertiary-400))" }, ".dark .decoration-tertiary-500-400-token": { "textDecorationColor": "rgb(var(--color-tertiary-400))" }, ".dark .text-success-50-900-token": { "color": "rgb(var(--color-success-900))" }, ".dark .decoration-success-50-900-token": { "textDecorationColor": "rgb(var(--color-success-900))" }, ".dark .text-success-100-800-token": { "color": "rgb(var(--color-success-800))" }, ".dark .decoration-success-100-800-token": { "textDecorationColor": "rgb(var(--color-success-800))" }, ".dark .text-success-200-700-token": { "color": "rgb(var(--color-success-700))" }, ".dark .decoration-success-200-700-token": { "textDecorationColor": "rgb(var(--color-success-700))" }, ".dark .text-success-300-600-token": { "color": "rgb(var(--color-success-600))" }, ".dark .decoration-success-300-600-token": { "textDecorationColor": "rgb(var(--color-success-600))" }, ".dark .text-success-400-500-token": { "color": "rgb(var(--color-success-500))" }, ".dark .decoration-success-400-500-token": { "textDecorationColor": "rgb(var(--color-success-500))" }, ".dark .text-success-900-50-token": { "color": "rgb(var(--color-success-50))" }, ".dark .decoration-success-900-50-token": { "textDecorationColor": "rgb(var(--color-success-50))" }, ".dark .text-success-800-100-token": { "color": "rgb(var(--color-success-100))" }, ".dark .decoration-success-800-100-token": { "textDecorationColor": "rgb(var(--color-success-100))" }, ".dark .text-success-700-200-token": { "color": "rgb(var(--color-success-200))" }, ".dark .decoration-success-700-200-token": { "textDecorationColor": "rgb(var(--color-success-200))" }, ".dark .text-success-600-300-token": { "color": "rgb(var(--color-success-300))" }, ".dark .decoration-success-600-300-token": { "textDecorationColor": "rgb(var(--color-success-300))" }, ".dark .text-success-500-400-token": { "color": "rgb(var(--color-success-400))" }, ".dark .decoration-success-500-400-token": { "textDecorationColor": "rgb(var(--color-success-400))" }, ".dark .text-warning-50-900-token": { "color": "rgb(var(--color-warning-900))" }, ".dark .decoration-warning-50-900-token": { "textDecorationColor": "rgb(var(--color-warning-900))" }, ".dark .text-warning-100-800-token": { "color": "rgb(var(--color-warning-800))" }, ".dark .decoration-warning-100-800-token": { "textDecorationColor": "rgb(var(--color-warning-800))" }, ".dark .text-warning-200-700-token": { "color": "rgb(var(--color-warning-700))" }, ".dark .decoration-warning-200-700-token": { "textDecorationColor": "rgb(var(--color-warning-700))" }, ".dark .text-warning-300-600-token": { "color": "rgb(var(--color-warning-600))" }, ".dark .decoration-warning-300-600-token": { "textDecorationColor": "rgb(var(--color-warning-600))" }, ".dark .text-warning-400-500-token": { "color": "rgb(var(--color-warning-500))" }, ".dark .decoration-warning-400-500-token": { "textDecorationColor": "rgb(var(--color-warning-500))" }, ".dark .text-warning-900-50-token": { "color": "rgb(var(--color-warning-50))" }, ".dark .decoration-warning-900-50-token": { "textDecorationColor": "rgb(var(--color-warning-50))" }, ".dark .text-warning-800-100-token": { "color": "rgb(var(--color-warning-100))" }, ".dark .decoration-warning-800-100-token": { "textDecorationColor": "rgb(var(--color-warning-100))" }, ".dark .text-warning-700-200-token": { "color": "rgb(var(--color-warning-200))" }, ".dark .decoration-warning-700-200-token": { "textDecorationColor": "rgb(var(--color-warning-200))" }, ".dark .text-warning-600-300-token": { "color": "rgb(var(--color-warning-300))" }, ".dark .decoration-warning-600-300-token": { "textDecorationColor": "rgb(var(--color-warning-300))" }, ".dark .text-warning-500-400-token": { "color": "rgb(var(--color-warning-400))" }, ".dark .decoration-warning-500-400-token": { "textDecorationColor": "rgb(var(--color-warning-400))" }, ".dark .text-error-50-900-token": { "color": "rgb(var(--color-error-900))" }, ".dark .decoration-error-50-900-token": { "textDecorationColor": "rgb(var(--color-error-900))" }, ".dark .text-error-100-800-token": { "color": "rgb(var(--color-error-800))" }, ".dark .decoration-error-100-800-token": { "textDecorationColor": "rgb(var(--color-error-800))" }, ".dark .text-error-200-700-token": { "color": "rgb(var(--color-error-700))" }, ".dark .decoration-error-200-700-token": { "textDecorationColor": "rgb(var(--color-error-700))" }, ".dark .text-error-300-600-token": { "color": "rgb(var(--color-error-600))" }, ".dark .decoration-error-300-600-token": { "textDecorationColor": "rgb(var(--color-error-600))" }, ".dark .text-error-400-500-token": { "color": "rgb(var(--color-error-500))" }, ".dark .decoration-error-400-500-token": { "textDecorationColor": "rgb(var(--color-error-500))" }, ".dark .text-error-900-50-token": { "color": "rgb(var(--color-error-50))" }, ".dark .decoration-error-900-50-token": { "textDecorationColor": "rgb(var(--color-error-50))" }, ".dark .text-error-800-100-token": { "color": "rgb(var(--color-error-100))" }, ".dark .decoration-error-800-100-token": { "textDecorationColor": "rgb(var(--color-error-100))" }, ".dark .text-error-700-200-token": { "color": "rgb(var(--color-error-200))" }, ".dark .decoration-error-700-200-token": { "textDecorationColor": "rgb(var(--color-error-200))" }, ".dark .text-error-600-300-token": { "color": "rgb(var(--color-error-300))" }, ".dark .decoration-error-600-300-token": { "textDecorationColor": "rgb(var(--color-error-300))" }, ".dark .text-error-500-400-token": { "color": "rgb(var(--color-error-400))" }, ".dark .decoration-error-500-400-token": { "textDecorationColor": "rgb(var(--color-error-400))" }, ".dark .text-surface-50-900-token": { "color": "rgb(var(--color-surface-900))" }, ".dark .decoration-surface-50-900-token": { "textDecorationColor": "rgb(var(--color-surface-900))" }, ".dark .text-surface-100-800-token": { "color": "rgb(var(--color-surface-800))" }, ".dark .decoration-surface-100-800-token": { "textDecorationColor": "rgb(var(--color-surface-800))" }, ".dark .text-surface-200-700-token": { "color": "rgb(var(--color-surface-700))" }, ".dark .decoration-surface-200-700-token": { "textDecorationColor": "rgb(var(--color-surface-700))" }, ".dark .text-surface-300-600-token": { "color": "rgb(var(--color-surface-600))" }, ".dark .decoration-surface-300-600-token": { "textDecorationColor": "rgb(var(--color-surface-600))" }, ".dark .text-surface-400-500-token": { "color": "rgb(var(--color-surface-500))" }, ".dark .decoration-surface-400-500-token": { "textDecorationColor": "rgb(var(--color-surface-500))" }, ".dark .text-surface-900-50-token": { "color": "rgb(var(--color-surface-50))" }, ".dark .decoration-surface-900-50-token": { "textDecorationColor": "rgb(var(--color-surface-50))" }, ".dark .text-surface-800-100-token": { "color": "rgb(var(--color-surface-100))" }, ".dark .decoration-surface-800-100-token": { "textDecorationColor": "rgb(var(--color-surface-100))" }, ".dark .text-surface-700-200-token": { "color": "rgb(var(--color-surface-200))" }, ".dark .decoration-surface-700-200-token": { "textDecorationColor": "rgb(var(--color-surface-200))" }, ".dark .text-surface-600-300-token": { "color": "rgb(var(--color-surface-300))" }, ".dark .decoration-surface-600-300-token": { "textDecorationColor": "rgb(var(--color-surface-300))" }, ".dark .text-surface-500-400-token": { "color": "rgb(var(--color-surface-400))" }, ".dark .decoration-surface-500-400-token": { "textDecorationColor": "rgb(var(--color-surface-400))" }, ".dark .ring-outline-token": { "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)", "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)", "boxShadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)", "--tw-ring-inset": "inset", "--tw-ring-color": "rgb(250 250 250 / 0.05)" }, ".dark .ring-primary-50-900-token": { "--tw-ring-color": "rgb(var(--color-primary-900) / 1)" }, ".dark .ring-primary-100-800-token": { "--tw-ring-color": "rgb(var(--color-primary-800) / 1)" }, ".dark .ring-primary-200-700-token": { "--tw-ring-color": "rgb(var(--color-primary-700) / 1)" }, ".dark .ring-primary-300-600-token": { "--tw-ring-color": "rgb(var(--color-primary-600) / 1)" }, ".dark .ring-primary-400-500-token": { "--tw-ring-color": "rgb(var(--color-primary-500) / 1)" }, ".dark .ring-primary-900-50-token": { "--tw-ring-color": "rgb(var(--color-primary-50) / 1)" }, ".dark .ring-primary-800-100-token": { "--tw-ring-color": "rgb(var(--color-primary-100) / 1)" }, ".dark .ring-primary-700-200-token": { "--tw-ring-color": "rgb(var(--color-primary-200) / 1)" }, ".dark .ring-primary-600-300-token": { "--tw-ring-color": "rgb(var(--color-primary-300) / 1)" }, ".dark .ring-primary-500-400-token": { "--tw-ring-color": "rgb(var(--color-primary-400) / 1)" }, ".dark .ring-secondary-50-900-token": { "--tw-ring-color": "rgb(var(--color-secondary-900) / 1)" }, ".dark .ring-secondary-100-800-token": { "--tw-ring-color": "rgb(var(--color-secondary-800) / 1)" }, ".dark .ring-secondary-200-700-token": { "--tw-ring-color": "rgb(var(--color-secondary-700) / 1)" }, ".dark .ring-secondary-300-600-token": { "--tw-ring-color": "rgb(var(--color-secondary-600) / 1)" }, ".dark .ring-secondary-400-500-token": { "--tw-ring-color": "rgb(var(--color-secondary-500) / 1)" }, ".dark .ring-secondary-900-50-token": { "--tw-ring-color": "rgb(var(--color-secondary-50) / 1)" }, ".dark .ring-secondary-800-100-token": { "--tw-ring-color": "rgb(var(--color-secondary-100) / 1)" }, ".dark .ring-secondary-700-200-token": { "--tw-ring-color": "rgb(var(--color-secondary-200) / 1)" }, ".dark .ring-secondary-600-300-token": { "--tw-ring-color": "rgb(var(--color-secondary-300) / 1)" }, ".dark .ring-secondary-500-400-token": { "--tw-ring-color": "rgb(var(--color-secondary-400) / 1)" }, ".dark .ring-tertiary-50-900-token": { "--tw-ring-color": "rgb(var(--color-tertiary-900) / 1)" }, ".dark .ring-tertiary-100-800-token": { "--tw-ring-color": "rgb(var(--color-tertiary-800) / 1)" }, ".dark .ring-tertiary-200-700-token": { "--tw-ring-color": "rgb(var(--color-tertiary-700) / 1)" }, ".dark .ring-tertiary-300-600-token": { "--tw-ring-color": "rgb(var(--color-tertiary-600) / 1)" }, ".dark .ring-tertiary-400-500-token": { "--tw-ring-color": "rgb(var(--color-tertiary-500) / 1)" }, ".dark .ring-tertiary-900-50-token": { "--tw-ring-color": "rgb(var(--color-tertiary-50) / 1)" }, ".dark .ring-tertiary-800-100-token": { "--tw-ring-color": "rgb(var(--color-tertiary-100) / 1)" }, ".dark .ring-tertiary-700-200-token": { "--tw-ring-color": "rgb(var(--color-tertiary-200) / 1)" }, ".dark .ring-tertiary-600-300-token": { "--tw-ring-color": "rgb(var(--color-tertiary-300) / 1)" }, ".dark .ring-tertiary-500-400-token": { "--tw-ring-color": "rgb(var(--color-tertiary-400) / 1)" }, ".dark .ring-success-50-900-token": { "--tw-ring-color": "rgb(var(--color-success-900) / 1)" }, ".dark .ring-success-100-800-token": { "--tw-ring-color": "rgb(var(--color-success-800) / 1)" }, ".dark .ring-success-200-700-token": { "--tw-ring-color": "rgb(var(--color-success-700) / 1)" }, ".dark .ring-success-300-600-token": { "--tw-ring-color": "rgb(var(--color-success-600) / 1)" }, ".dark .ring-success-400-500-token": { "--tw-ring-color": "rgb(var(--color-success-500) / 1)" }, ".dark .ring-success-900-50-token": { "--tw-ring-color": "rgb(var(--color-success-50) / 1)" }, ".dark .ring-success-800-100-token": { "--tw-ring-color": "rgb(var(--color-success-100) / 1)" }, ".dark .ring-success-700-200-token": { "--tw-ring-color": "rgb(var(--color-success-200) / 1)" }, ".dark .ring-success-600-300-token": { "--tw-ring-color": "rgb(var(--color-success-300) / 1)" }, ".dark .ring-success-500-400-token": { "--tw-ring-color": "rgb(var(--color-success-400) / 1)" }, ".dark .ring-warning-50-900-token": { "--tw-ring-color": "rgb(var(--color-warning-900) / 1)" }, ".dark .ring-warning-100-800-token": { "--tw-ring-color": "rgb(var(--color-warning-800) / 1)" }, ".dark .ring-warning-200-700-token": { "--tw-ring-color": "rgb(var(--color-warning-700) / 1)" }, ".dark .ring-warning-300-600-token": { "--tw-ring-color": "rgb(var(--color-warning-600) / 1)" }, ".dark .ring-warning-400-500-token": { "--tw-ring-color": "rgb(var(--color-warning-500) / 1)" }, ".dark .ring-warning-900-50-token": { "--tw-ring-color": "rgb(var(--color-warning-50) / 1)" }, ".dark .ring-warning-800-100-token": { "--tw-ring-color": "rgb(var(--color-warning-100) / 1)" }, ".dark .ring-warning-700-200-token": { "--tw-ring-color": "rgb(var(--color-warning-200) / 1)" }, ".dark .ring-warning-600-300-token": { "--tw-ring-color": "rgb(var(--color-warning-300) / 1)" }, ".dark .ring-warning-500-400-token": { "--tw-ring-color": "rgb(var(--color-warning-400) / 1)" }, ".dark .ring-error-50-900-token": { "--tw-ring-color": "rgb(var(--color-error-900) / 1)" }, ".dark .ring-error-100-800-token": { "--tw-ring-color": "rgb(var(--color-error-800) / 1)" }, ".dark .ring-error-200-700-token": { "--tw-ring-color": "rgb(var(--color-error-700) / 1)" }, ".dark .ring-error-300-600-token": { "--tw-ring-color": "rgb(var(--color-error-600) / 1)" }, ".dark .ring-error-400-500-token": { "--tw-ring-color": "rgb(var(--color-error-500) / 1)" }, ".dark .ring-error-900-50-token": { "--tw-ring-color": "rgb(var(--color-error-50) / 1)" }, ".dark .ring-error-800-100-token": { "--tw-ring-color": "rgb(var(--color-error-100) / 1)" }, ".dark .ring-error-700-200-token": { "--tw-ring-color": "rgb(var(--color-error-200) / 1)" }, ".dark .ring-error-600-300-token": { "--tw-ring-color": "rgb(var(--color-error-300) / 1)" }, ".dark .ring-error-500-400-token": { "--tw-ring-color": "rgb(var(--color-error-400) / 1)" }, ".dark .ring-surface-50-900-token": { "--tw-ring-color": "rgb(var(--color-surface-900) / 1)" }, ".dark .ring-surface-100-800-token": { "--tw-ring-color": "rgb(var(--color-surface-800) / 1)" }, ".dark .ring-surface-200-700-token": { "--tw-ring-color": "rgb(var(--color-surface-700) / 1)" }, ".dark .ring-surface-300-600-token": { "--tw-ring-color": "rgb(var(--color-surface-600) / 1)" }, ".dark .ring-surface-400-500-token": { "--tw-ring-color": "rgb(var(--color-surface-500) / 1)" }, ".dark .ring-surface-900-50-token": { "--tw-ring-color": "rgb(var(--color-surface-50) / 1)" }, ".dark .ring-surface-800-100-token": { "--tw-ring-color": "rgb(var(--color-surface-100) / 1)" }, ".dark .ring-surface-700-200-token": { "--tw-ring-color": "rgb(var(--color-surface-200) / 1)" }, ".dark .ring-surface-600-300-token": { "--tw-ring-color": "rgb(var(--color-surface-300) / 1)" }, ".dark .ring-surface-500-400-token": { "--tw-ring-color": "rgb(var(--color-surface-400) / 1)" }, "body": { "backgroundColor": "rgb(var(--color-surface-50))", "fontSize": "1rem", "lineHeight": "1.5rem", "fontFamily": "var(--theme-font-family-base)", "color": "rgba(var(--theme-font-color-base))" }, ".dark body": { "backgroundColor": "rgb(var(--color-surface-900))", "color": "rgba(var(--theme-font-color-dark))" }, "::selection": { "backgroundColor": "rgb(var(--color-primary-500) / 0.3)" }, "html": { "WebkitTapHighlightColor": "rgba(128, 128, 128, 0.5)", "scrollbarColor": "rgba(0, 0, 0, 0.2) rgba(255, 255, 255, 0.05)" }, "::-webkit-scrollbar": { "width": "0.5rem", "height": "0.5rem" }, "::-webkit-scrollbar-track": { "paddingLeft": "1px", "paddingRight": "1px", "backgroundColor": "rgb(var(--color-surface-50)) !important" }, ".dark ::-webkit-scrollbar-track": { "backgroundColor": "rgb(var(--color-surface-900)) !important" }, "::-webkit-scrollbar-thumb": { "backgroundColor": "rgb(var(--color-surface-400))", "borderRadius": "var(--theme-rounded-base)" }, ".dark ::-webkit-scrollbar-thumb": { "backgroundColor": "rgb(var(--color-surface-500))" }, "::-webkit-scrollbar-corner": { "backgroundColor": "rgb(var(--color-surface-50)) !important" }, ".dark ::-webkit-scrollbar-corner": { "backgroundColor": "rgb(var(--color-surface-900)) !important" }, "html.dark": { "scrollbarColor": "rgba(255, 255, 255, 0.1) rgba(0, 0, 0, 0.05)" }, "hr:not(.divider)": { "display": "block", "borderTopWidth": "1px", "borderStyle": "solid", "borderColor": "rgb(var(--color-surface-300))" }, ".dark hr:not(.divider)": { "borderColor": "rgb(var(--color-surface-600))" }, "fieldset,\nlegend,\nlabel": { "display": "block" }, "::-moz-placeholder": { "color": "rgb(var(--color-surface-500))" }, ".dark ::-moz-placeholder": { "color": "rgb(var(--color-surface-400))" }, ":-ms-input-placeholder": { "color": "rgb(var(--color-surface-500))" }, ".dark :-ms-input-placeholder": { "color": "rgb(var(--color-surface-400))" }, "::placeholder": { "color": "rgb(var(--color-surface-500))" }, ".dark ::placeholder": { "color": "rgb(var(--color-surface-400))" }, ":is(.dark input::-webkit-calendar-picker-indicator)": { "--tw-invert": "invert(100%)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, "input[type='search']::-webkit-search-cancel-button": { "WebkitAppearance": "none", "background": `url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 512 512'%3E%3Cpath d='M256 8C119 8 8 119 8 256s111 248 248 248 248-111 248-248S393 8 256 8zm121.6 313.1c4.7 4.7 4.7 12.3 0 17L338 377.6c-4.7 4.7-12.3 4.7-17 0L256 312l-65.1 65.6c-4.7 4.7-12.3 4.7-17 0L134.4 338c-4.7-4.7-4.7-12.3 0-17l65.6-65-65.6-65.1c-4.7-4.7-4.7-12.3 0-17l39.6-39.6c4.7-4.7 12.3-4.7 17 0l65 65.7 65.1-65.6c4.7-4.7 12.3-4.7 17 0l39.6 39.6c4.7 4.7 4.7 12.3 0 17L312 256l65.6 65.1z'/%3E%3C/svg%3E")
		no-repeat 50% 50%`, "pointerEvents": "none", "height": "1rem", "width": "1rem", "borderRadius": "9999px", "backgroundSize": "contain", "opacity": 0 }, "input[type='search']:focus::-webkit-search-cancel-button": { "pointerEvents": "auto", "opacity": 1 }, ":is(.dark input[type='search']::-webkit-search-cancel-button)": { "--tw-invert": "invert(100%)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, "progress": { "webkitAppearance": "none", "MozAppearance": "none", "appearance": "none", "height": "0.5rem", "width": "100%", "overflow": "hidden", "borderRadius": "var(--theme-rounded-base)", "backgroundColor": "rgb(var(--color-surface-400))" }, ".dark progress": { "backgroundColor": "rgb(var(--color-surface-500))" }, "progress::-webkit-progress-bar": { "backgroundColor": "rgb(var(--color-surface-400))" }, ".dark progress::-webkit-progress-bar": { "backgroundColor": "rgb(var(--color-surface-500))" }, "progress::-webkit-progress-value": { "backgroundColor": "rgb(var(--color-surface-900))" }, ".dark progress::-webkit-progress-value": { "backgroundColor": "rgb(var(--color-surface-50))" }, "::-moz-progress-bar": { "backgroundColor": "rgb(var(--color-surface-900))" }, ".dark ::-moz-progress-bar": { "backgroundColor": "rgb(var(--color-surface-50))" }, ":indeterminate::-moz-progress-bar": { "width": "0" }, "input[type='file']:not(.file-dropzone-input)::file-selector-button:disabled": { "cursor": "not-allowed", "opacity": 0.5 }, "input[type='file']:not(.file-dropzone-input)::file-selector-button:disabled:hover": { "--tw-brightness": "brightness(1)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, "input[type='file']:not(.file-dropzone-input)::file-selector-button:disabled:active": { "--tw-scale-x": "1", "--tw-scale-y": "1", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))" }, "input[type='file']:not(.file-dropzone-input)::file-selector-button": { "fontSize": "0.875rem", "lineHeight": "1.25rem", "paddingLeft": "0.75rem", "paddingRight": "0.75rem", "paddingTop": "0.375rem", "paddingBottom": "0.375rem", "whiteSpace": "nowrap", "textAlign": "center", "display": "inline-flex", "alignItems": "center", "justifyContent": "center", "transitionProperty": "all", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms", "borderRadius": "var(--theme-rounded-base)", "backgroundColor": "rgb(var(--color-surface-900))", "color": "rgb(var(--color-surface-50))", "marginRight": "0.5rem", "borderWidth": "0px" }, "input[type='file']:not(.file-dropzone-input)::file-selector-button > :not([hidden]) ~ :not([hidden])": { "--tw-space-x-reverse": "0", "marginRight": "calc(0.5rem * var(--tw-space-x-reverse))", "marginLeft": "calc(0.5rem * calc(1 - var(--tw-space-x-reverse)))" }, "input[type='file']:not(.file-dropzone-input)::file-selector-button:hover": { "--tw-brightness": "brightness(1.15)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, "input[type='file']:not(.file-dropzone-input)::file-selector-button:active": { "--tw-scale-x": "95%", "--tw-scale-y": "95%", "transform": "translate(var(--tw-translate-x), var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))", "--tw-brightness": "brightness(.9)", "filter": "var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)" }, ".dark input[type='file']:not(.file-dropzone-input)::file-selector-button": { "backgroundColor": "rgb(var(--color-surface-50))", "color": "rgb(var(--color-surface-900))" }, "[type='range']": { "width": "100%", "accentColor": "rgb(var(--color-surface-900) / 1)" }, ":is(.dark [type='range'])": { "accentColor": "rgb(var(--color-surface-50) / 1)" }, "[data-sort]": { "cursor": "pointer" }, "[data-sort]:hover:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, ".dark [data-sort]:hover:hover": { "backgroundColor": "rgb(var(--color-primary-500) / 0.1)" }, "[data-sort]::after": { "marginLeft": "0.5rem !important", "opacity": 0, "--tw-content": "'\u2193' !important", "content": "var(--tw-content) !important" }, "[data-popup]": { "position": "absolute", "top": "0px", "left": "0px", "display": "none", "transitionProperty": "opacity", "transitionTimingFunction": "cubic-bezier(0.4, 0, 0.2, 1)", "transitionDuration": "150ms" } } };
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/unesc.js
var require_unesc = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/unesc.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = unesc;
    function gobbleHex(str) {
      var lower = str.toLowerCase();
      var hex = "";
      var spaceTerminated = false;
      for (var i = 0; i < 6 && lower[i] !== void 0; i++) {
        var code = lower.charCodeAt(i);
        var valid = code >= 97 && code <= 102 || code >= 48 && code <= 57;
        spaceTerminated = code === 32;
        if (!valid) {
          break;
        }
        hex += lower[i];
      }
      if (hex.length === 0) {
        return void 0;
      }
      var codePoint = parseInt(hex, 16);
      var isSurrogate = codePoint >= 55296 && codePoint <= 57343;
      if (isSurrogate || codePoint === 0 || codePoint > 1114111) {
        return ["\uFFFD", hex.length + (spaceTerminated ? 1 : 0)];
      }
      return [String.fromCodePoint(codePoint), hex.length + (spaceTerminated ? 1 : 0)];
    }
    var CONTAINS_ESCAPE = /\\/;
    function unesc(str) {
      var needToProcess = CONTAINS_ESCAPE.test(str);
      if (!needToProcess) {
        return str;
      }
      var ret = "";
      for (var i = 0; i < str.length; i++) {
        if (str[i] === "\\") {
          var gobbled = gobbleHex(str.slice(i + 1, i + 7));
          if (gobbled !== void 0) {
            ret += gobbled[0];
            i += gobbled[1];
            continue;
          }
          if (str[i + 1] === "\\") {
            ret += "\\";
            i++;
            continue;
          }
          if (str.length === i + 1) {
            ret += str[i];
          }
          continue;
        }
        ret += str[i];
      }
      return ret;
    }
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/getProp.js
var require_getProp = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/getProp.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = getProp;
    function getProp(obj) {
      for (var _len = arguments.length, props = new Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
        props[_key - 1] = arguments[_key];
      }
      while (props.length > 0) {
        var prop = props.shift();
        if (!obj[prop]) {
          return void 0;
        }
        obj = obj[prop];
      }
      return obj;
    }
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/ensureObject.js
var require_ensureObject = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/ensureObject.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = ensureObject;
    function ensureObject(obj) {
      for (var _len = arguments.length, props = new Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
        props[_key - 1] = arguments[_key];
      }
      while (props.length > 0) {
        var prop = props.shift();
        if (!obj[prop]) {
          obj[prop] = {};
        }
        obj = obj[prop];
      }
    }
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/stripComments.js
var require_stripComments = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/stripComments.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = stripComments;
    function stripComments(str) {
      var s = "";
      var commentStart = str.indexOf("/*");
      var lastEnd = 0;
      while (commentStart >= 0) {
        s = s + str.slice(lastEnd, commentStart);
        var commentEnd = str.indexOf("*/", commentStart + 2);
        if (commentEnd < 0) {
          return s;
        }
        lastEnd = commentEnd + 2;
        commentStart = str.indexOf("/*", lastEnd);
      }
      s = s + str.slice(lastEnd);
      return s;
    }
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/index.js
var require_util2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/util/index.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports.unesc = exports.stripComments = exports.getProp = exports.ensureObject = void 0;
    var _unesc = _interopRequireDefault(require_unesc());
    exports.unesc = _unesc["default"];
    var _getProp = _interopRequireDefault(require_getProp());
    exports.getProp = _getProp["default"];
    var _ensureObject = _interopRequireDefault(require_ensureObject());
    exports.ensureObject = _ensureObject["default"];
    var _stripComments = _interopRequireDefault(require_stripComments());
    exports.stripComments = _stripComments["default"];
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/node.js
var require_node2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/node.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _util = require_util2();
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    var cloneNode = function cloneNode2(obj, parent) {
      if (typeof obj !== "object" || obj === null) {
        return obj;
      }
      var cloned = new obj.constructor();
      for (var i in obj) {
        if (!obj.hasOwnProperty(i)) {
          continue;
        }
        var value = obj[i];
        var type = typeof value;
        if (i === "parent" && type === "object") {
          if (parent) {
            cloned[i] = parent;
          }
        } else if (value instanceof Array) {
          cloned[i] = value.map(function(j) {
            return cloneNode2(j, cloned);
          });
        } else {
          cloned[i] = cloneNode2(value, cloned);
        }
      }
      return cloned;
    };
    var Node = /* @__PURE__ */ function() {
      function Node2(opts) {
        if (opts === void 0) {
          opts = {};
        }
        Object.assign(this, opts);
        this.spaces = this.spaces || {};
        this.spaces.before = this.spaces.before || "";
        this.spaces.after = this.spaces.after || "";
      }
      var _proto = Node2.prototype;
      _proto.remove = function remove() {
        if (this.parent) {
          this.parent.removeChild(this);
        }
        this.parent = void 0;
        return this;
      };
      _proto.replaceWith = function replaceWith() {
        if (this.parent) {
          for (var index2 in arguments) {
            this.parent.insertBefore(this, arguments[index2]);
          }
          this.remove();
        }
        return this;
      };
      _proto.next = function next() {
        return this.parent.at(this.parent.index(this) + 1);
      };
      _proto.prev = function prev() {
        return this.parent.at(this.parent.index(this) - 1);
      };
      _proto.clone = function clone(overrides) {
        if (overrides === void 0) {
          overrides = {};
        }
        var cloned = cloneNode(this);
        for (var name in overrides) {
          cloned[name] = overrides[name];
        }
        return cloned;
      };
      _proto.appendToPropertyAndEscape = function appendToPropertyAndEscape(name, value, valueEscaped) {
        if (!this.raws) {
          this.raws = {};
        }
        var originalValue = this[name];
        var originalEscaped = this.raws[name];
        this[name] = originalValue + value;
        if (originalEscaped || valueEscaped !== value) {
          this.raws[name] = (originalEscaped || originalValue) + valueEscaped;
        } else {
          delete this.raws[name];
        }
      };
      _proto.setPropertyAndEscape = function setPropertyAndEscape(name, value, valueEscaped) {
        if (!this.raws) {
          this.raws = {};
        }
        this[name] = value;
        this.raws[name] = valueEscaped;
      };
      _proto.setPropertyWithoutEscape = function setPropertyWithoutEscape(name, value) {
        this[name] = value;
        if (this.raws) {
          delete this.raws[name];
        }
      };
      _proto.isAtPosition = function isAtPosition(line, column) {
        if (this.source && this.source.start && this.source.end) {
          if (this.source.start.line > line) {
            return false;
          }
          if (this.source.end.line < line) {
            return false;
          }
          if (this.source.start.line === line && this.source.start.column > column) {
            return false;
          }
          if (this.source.end.line === line && this.source.end.column < column) {
            return false;
          }
          return true;
        }
        return void 0;
      };
      _proto.stringifyProperty = function stringifyProperty(name) {
        return this.raws && this.raws[name] || this[name];
      };
      _proto.valueToString = function valueToString() {
        return String(this.stringifyProperty("value"));
      };
      _proto.toString = function toString() {
        return [this.rawSpaceBefore, this.valueToString(), this.rawSpaceAfter].join("");
      };
      _createClass(Node2, [{
        key: "rawSpaceBefore",
        get: function get() {
          var rawSpace = this.raws && this.raws.spaces && this.raws.spaces.before;
          if (rawSpace === void 0) {
            rawSpace = this.spaces && this.spaces.before;
          }
          return rawSpace || "";
        },
        set: function set(raw) {
          (0, _util.ensureObject)(this, "raws", "spaces");
          this.raws.spaces.before = raw;
        }
      }, {
        key: "rawSpaceAfter",
        get: function get() {
          var rawSpace = this.raws && this.raws.spaces && this.raws.spaces.after;
          if (rawSpace === void 0) {
            rawSpace = this.spaces.after;
          }
          return rawSpace || "";
        },
        set: function set(raw) {
          (0, _util.ensureObject)(this, "raws", "spaces");
          this.raws.spaces.after = raw;
        }
      }]);
      return Node2;
    }();
    exports["default"] = Node;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/types.js
var require_types = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/types.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports.UNIVERSAL = exports.TAG = exports.STRING = exports.SELECTOR = exports.ROOT = exports.PSEUDO = exports.NESTING = exports.ID = exports.COMMENT = exports.COMBINATOR = exports.CLASS = exports.ATTRIBUTE = void 0;
    var TAG = "tag";
    exports.TAG = TAG;
    var STRING = "string";
    exports.STRING = STRING;
    var SELECTOR = "selector";
    exports.SELECTOR = SELECTOR;
    var ROOT = "root";
    exports.ROOT = ROOT;
    var PSEUDO = "pseudo";
    exports.PSEUDO = PSEUDO;
    var NESTING = "nesting";
    exports.NESTING = NESTING;
    var ID = "id";
    exports.ID = ID;
    var COMMENT = "comment";
    exports.COMMENT = COMMENT;
    var COMBINATOR = "combinator";
    exports.COMBINATOR = COMBINATOR;
    var CLASS = "class";
    exports.CLASS = CLASS;
    var ATTRIBUTE = "attribute";
    exports.ATTRIBUTE = ATTRIBUTE;
    var UNIVERSAL = "universal";
    exports.UNIVERSAL = UNIVERSAL;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/container.js
var require_container2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/container.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _node = _interopRequireDefault(require_node2());
    var types = _interopRequireWildcard(require_types());
    function _getRequireWildcardCache(nodeInterop) {
      if (typeof WeakMap !== "function")
        return null;
      var cacheBabelInterop = /* @__PURE__ */ new WeakMap();
      var cacheNodeInterop = /* @__PURE__ */ new WeakMap();
      return (_getRequireWildcardCache = function _getRequireWildcardCache2(nodeInterop2) {
        return nodeInterop2 ? cacheNodeInterop : cacheBabelInterop;
      })(nodeInterop);
    }
    function _interopRequireWildcard(obj, nodeInterop) {
      if (!nodeInterop && obj && obj.__esModule) {
        return obj;
      }
      if (obj === null || typeof obj !== "object" && typeof obj !== "function") {
        return { "default": obj };
      }
      var cache = _getRequireWildcardCache(nodeInterop);
      if (cache && cache.has(obj)) {
        return cache.get(obj);
      }
      var newObj = {};
      var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor;
      for (var key in obj) {
        if (key !== "default" && Object.prototype.hasOwnProperty.call(obj, key)) {
          var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null;
          if (desc && (desc.get || desc.set)) {
            Object.defineProperty(newObj, key, desc);
          } else {
            newObj[key] = obj[key];
          }
        }
      }
      newObj["default"] = obj;
      if (cache) {
        cache.set(obj, newObj);
      }
      return newObj;
    }
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _createForOfIteratorHelperLoose(o, allowArrayLike) {
      var it = typeof Symbol !== "undefined" && o[Symbol.iterator] || o["@@iterator"];
      if (it)
        return (it = it.call(o)).next.bind(it);
      if (Array.isArray(o) || (it = _unsupportedIterableToArray(o)) || allowArrayLike && o && typeof o.length === "number") {
        if (it)
          o = it;
        var i = 0;
        return function() {
          if (i >= o.length)
            return { done: true };
          return { done: false, value: o[i++] };
        };
      }
      throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.");
    }
    function _unsupportedIterableToArray(o, minLen) {
      if (!o)
        return;
      if (typeof o === "string")
        return _arrayLikeToArray(o, minLen);
      var n = Object.prototype.toString.call(o).slice(8, -1);
      if (n === "Object" && o.constructor)
        n = o.constructor.name;
      if (n === "Map" || n === "Set")
        return Array.from(o);
      if (n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))
        return _arrayLikeToArray(o, minLen);
    }
    function _arrayLikeToArray(arr, len) {
      if (len == null || len > arr.length)
        len = arr.length;
      for (var i = 0, arr2 = new Array(len); i < len; i++) {
        arr2[i] = arr[i];
      }
      return arr2;
    }
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Container = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(Container2, _Node);
      function Container2(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        if (!_this.nodes) {
          _this.nodes = [];
        }
        return _this;
      }
      var _proto = Container2.prototype;
      _proto.append = function append(selector) {
        selector.parent = this;
        this.nodes.push(selector);
        return this;
      };
      _proto.prepend = function prepend(selector) {
        selector.parent = this;
        this.nodes.unshift(selector);
        return this;
      };
      _proto.at = function at(index2) {
        return this.nodes[index2];
      };
      _proto.index = function index2(child) {
        if (typeof child === "number") {
          return child;
        }
        return this.nodes.indexOf(child);
      };
      _proto.removeChild = function removeChild(child) {
        child = this.index(child);
        this.at(child).parent = void 0;
        this.nodes.splice(child, 1);
        var index2;
        for (var id in this.indexes) {
          index2 = this.indexes[id];
          if (index2 >= child) {
            this.indexes[id] = index2 - 1;
          }
        }
        return this;
      };
      _proto.removeAll = function removeAll() {
        for (var _iterator = _createForOfIteratorHelperLoose(this.nodes), _step; !(_step = _iterator()).done; ) {
          var node = _step.value;
          node.parent = void 0;
        }
        this.nodes = [];
        return this;
      };
      _proto.empty = function empty() {
        return this.removeAll();
      };
      _proto.insertAfter = function insertAfter(oldNode, newNode) {
        newNode.parent = this;
        var oldIndex = this.index(oldNode);
        this.nodes.splice(oldIndex + 1, 0, newNode);
        newNode.parent = this;
        var index2;
        for (var id in this.indexes) {
          index2 = this.indexes[id];
          if (oldIndex <= index2) {
            this.indexes[id] = index2 + 1;
          }
        }
        return this;
      };
      _proto.insertBefore = function insertBefore(oldNode, newNode) {
        newNode.parent = this;
        var oldIndex = this.index(oldNode);
        this.nodes.splice(oldIndex, 0, newNode);
        newNode.parent = this;
        var index2;
        for (var id in this.indexes) {
          index2 = this.indexes[id];
          if (index2 <= oldIndex) {
            this.indexes[id] = index2 + 1;
          }
        }
        return this;
      };
      _proto._findChildAtPosition = function _findChildAtPosition(line, col) {
        var found = void 0;
        this.each(function(node) {
          if (node.atPosition) {
            var foundChild = node.atPosition(line, col);
            if (foundChild) {
              found = foundChild;
              return false;
            }
          } else if (node.isAtPosition(line, col)) {
            found = node;
            return false;
          }
        });
        return found;
      };
      _proto.atPosition = function atPosition(line, col) {
        if (this.isAtPosition(line, col)) {
          return this._findChildAtPosition(line, col) || this;
        } else {
          return void 0;
        }
      };
      _proto._inferEndPosition = function _inferEndPosition() {
        if (this.last && this.last.source && this.last.source.end) {
          this.source = this.source || {};
          this.source.end = this.source.end || {};
          Object.assign(this.source.end, this.last.source.end);
        }
      };
      _proto.each = function each(callback) {
        if (!this.lastEach) {
          this.lastEach = 0;
        }
        if (!this.indexes) {
          this.indexes = {};
        }
        this.lastEach++;
        var id = this.lastEach;
        this.indexes[id] = 0;
        if (!this.length) {
          return void 0;
        }
        var index2, result;
        while (this.indexes[id] < this.length) {
          index2 = this.indexes[id];
          result = callback(this.at(index2), index2);
          if (result === false) {
            break;
          }
          this.indexes[id] += 1;
        }
        delete this.indexes[id];
        if (result === false) {
          return false;
        }
      };
      _proto.walk = function walk(callback) {
        return this.each(function(node, i) {
          var result = callback(node, i);
          if (result !== false && node.length) {
            result = node.walk(callback);
          }
          if (result === false) {
            return false;
          }
        });
      };
      _proto.walkAttributes = function walkAttributes(callback) {
        var _this2 = this;
        return this.walk(function(selector) {
          if (selector.type === types.ATTRIBUTE) {
            return callback.call(_this2, selector);
          }
        });
      };
      _proto.walkClasses = function walkClasses(callback) {
        var _this3 = this;
        return this.walk(function(selector) {
          if (selector.type === types.CLASS) {
            return callback.call(_this3, selector);
          }
        });
      };
      _proto.walkCombinators = function walkCombinators(callback) {
        var _this4 = this;
        return this.walk(function(selector) {
          if (selector.type === types.COMBINATOR) {
            return callback.call(_this4, selector);
          }
        });
      };
      _proto.walkComments = function walkComments(callback) {
        var _this5 = this;
        return this.walk(function(selector) {
          if (selector.type === types.COMMENT) {
            return callback.call(_this5, selector);
          }
        });
      };
      _proto.walkIds = function walkIds(callback) {
        var _this6 = this;
        return this.walk(function(selector) {
          if (selector.type === types.ID) {
            return callback.call(_this6, selector);
          }
        });
      };
      _proto.walkNesting = function walkNesting(callback) {
        var _this7 = this;
        return this.walk(function(selector) {
          if (selector.type === types.NESTING) {
            return callback.call(_this7, selector);
          }
        });
      };
      _proto.walkPseudos = function walkPseudos(callback) {
        var _this8 = this;
        return this.walk(function(selector) {
          if (selector.type === types.PSEUDO) {
            return callback.call(_this8, selector);
          }
        });
      };
      _proto.walkTags = function walkTags(callback) {
        var _this9 = this;
        return this.walk(function(selector) {
          if (selector.type === types.TAG) {
            return callback.call(_this9, selector);
          }
        });
      };
      _proto.walkUniversals = function walkUniversals(callback) {
        var _this10 = this;
        return this.walk(function(selector) {
          if (selector.type === types.UNIVERSAL) {
            return callback.call(_this10, selector);
          }
        });
      };
      _proto.split = function split(callback) {
        var _this11 = this;
        var current = [];
        return this.reduce(function(memo, node, index2) {
          var split2 = callback.call(_this11, node);
          current.push(node);
          if (split2) {
            memo.push(current);
            current = [];
          } else if (index2 === _this11.length - 1) {
            memo.push(current);
          }
          return memo;
        }, []);
      };
      _proto.map = function map(callback) {
        return this.nodes.map(callback);
      };
      _proto.reduce = function reduce(callback, memo) {
        return this.nodes.reduce(callback, memo);
      };
      _proto.every = function every(callback) {
        return this.nodes.every(callback);
      };
      _proto.some = function some(callback) {
        return this.nodes.some(callback);
      };
      _proto.filter = function filter(callback) {
        return this.nodes.filter(callback);
      };
      _proto.sort = function sort(callback) {
        return this.nodes.sort(callback);
      };
      _proto.toString = function toString() {
        return this.map(String).join("");
      };
      _createClass(Container2, [{
        key: "first",
        get: function get() {
          return this.at(0);
        }
      }, {
        key: "last",
        get: function get() {
          return this.at(this.length - 1);
        }
      }, {
        key: "length",
        get: function get() {
          return this.nodes.length;
        }
      }]);
      return Container2;
    }(_node["default"]);
    exports["default"] = Container;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/root.js
var require_root2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/root.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _container = _interopRequireDefault(require_container2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Root = /* @__PURE__ */ function(_Container) {
      _inheritsLoose(Root2, _Container);
      function Root2(opts) {
        var _this;
        _this = _Container.call(this, opts) || this;
        _this.type = _types.ROOT;
        return _this;
      }
      var _proto = Root2.prototype;
      _proto.toString = function toString() {
        var str = this.reduce(function(memo, selector) {
          memo.push(String(selector));
          return memo;
        }, []).join(",");
        return this.trailingComma ? str + "," : str;
      };
      _proto.error = function error(message, options) {
        if (this._error) {
          return this._error(message, options);
        } else {
          return new Error(message);
        }
      };
      _createClass(Root2, [{
        key: "errorGenerator",
        set: function set(handler) {
          this._error = handler;
        }
      }]);
      return Root2;
    }(_container["default"]);
    exports["default"] = Root;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/selector.js
var require_selector = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/selector.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _container = _interopRequireDefault(require_container2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Selector = /* @__PURE__ */ function(_Container) {
      _inheritsLoose(Selector2, _Container);
      function Selector2(opts) {
        var _this;
        _this = _Container.call(this, opts) || this;
        _this.type = _types.SELECTOR;
        return _this;
      }
      return Selector2;
    }(_container["default"]);
    exports["default"] = Selector;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/cssesc@3.0.0/node_modules/cssesc/cssesc.js
var require_cssesc = __commonJS({
  "node_modules/.pnpm/cssesc@3.0.0/node_modules/cssesc/cssesc.js"(exports, module2) {
    "use strict";
    var object = {};
    var hasOwnProperty = object.hasOwnProperty;
    var merge = function merge2(options, defaults) {
      if (!options) {
        return defaults;
      }
      var result = {};
      for (var key in defaults) {
        result[key] = hasOwnProperty.call(options, key) ? options[key] : defaults[key];
      }
      return result;
    };
    var regexAnySingleEscape = /[ -,\.\/:-@\[-\^`\{-~]/;
    var regexSingleEscape = /[ -,\.\/:-@\[\]\^`\{-~]/;
    var regexExcessiveSpaces = /(^|\\+)?(\\[A-F0-9]{1,6})\x20(?![a-fA-F0-9\x20])/g;
    var cssesc = function cssesc2(string, options) {
      options = merge(options, cssesc2.options);
      if (options.quotes != "single" && options.quotes != "double") {
        options.quotes = "single";
      }
      var quote = options.quotes == "double" ? '"' : "'";
      var isIdentifier = options.isIdentifier;
      var firstChar = string.charAt(0);
      var output = "";
      var counter = 0;
      var length = string.length;
      while (counter < length) {
        var character = string.charAt(counter++);
        var codePoint = character.charCodeAt();
        var value = void 0;
        if (codePoint < 32 || codePoint > 126) {
          if (codePoint >= 55296 && codePoint <= 56319 && counter < length) {
            var extra = string.charCodeAt(counter++);
            if ((extra & 64512) == 56320) {
              codePoint = ((codePoint & 1023) << 10) + (extra & 1023) + 65536;
            } else {
              counter--;
            }
          }
          value = "\\" + codePoint.toString(16).toUpperCase() + " ";
        } else {
          if (options.escapeEverything) {
            if (regexAnySingleEscape.test(character)) {
              value = "\\" + character;
            } else {
              value = "\\" + codePoint.toString(16).toUpperCase() + " ";
            }
          } else if (/[\t\n\f\r\x0B]/.test(character)) {
            value = "\\" + codePoint.toString(16).toUpperCase() + " ";
          } else if (character == "\\" || !isIdentifier && (character == '"' && quote == character || character == "'" && quote == character) || isIdentifier && regexSingleEscape.test(character)) {
            value = "\\" + character;
          } else {
            value = character;
          }
        }
        output += value;
      }
      if (isIdentifier) {
        if (/^-[-\d]/.test(output)) {
          output = "\\-" + output.slice(1);
        } else if (/\d/.test(firstChar)) {
          output = "\\3" + firstChar + " " + output.slice(1);
        }
      }
      output = output.replace(regexExcessiveSpaces, function($0, $1, $2) {
        if ($1 && $1.length % 2) {
          return $0;
        }
        return ($1 || "") + $2;
      });
      if (!isIdentifier && options.wrap) {
        return quote + output + quote;
      }
      return output;
    };
    cssesc.options = {
      "escapeEverything": false,
      "isIdentifier": false,
      "quotes": "single",
      "wrap": false
    };
    cssesc.version = "3.0.0";
    module2.exports = cssesc;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/className.js
var require_className = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/className.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _cssesc = _interopRequireDefault(require_cssesc());
    var _util = require_util2();
    var _node = _interopRequireDefault(require_node2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var ClassName = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(ClassName2, _Node);
      function ClassName2(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        _this.type = _types.CLASS;
        _this._constructed = true;
        return _this;
      }
      var _proto = ClassName2.prototype;
      _proto.valueToString = function valueToString() {
        return "." + _Node.prototype.valueToString.call(this);
      };
      _createClass(ClassName2, [{
        key: "value",
        get: function get() {
          return this._value;
        },
        set: function set(v) {
          if (this._constructed) {
            var escaped = (0, _cssesc["default"])(v, {
              isIdentifier: true
            });
            if (escaped !== v) {
              (0, _util.ensureObject)(this, "raws");
              this.raws.value = escaped;
            } else if (this.raws) {
              delete this.raws.value;
            }
          }
          this._value = v;
        }
      }]);
      return ClassName2;
    }(_node["default"]);
    exports["default"] = ClassName;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/comment.js
var require_comment2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/comment.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _node = _interopRequireDefault(require_node2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Comment = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(Comment2, _Node);
      function Comment2(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        _this.type = _types.COMMENT;
        return _this;
      }
      return Comment2;
    }(_node["default"]);
    exports["default"] = Comment;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/id.js
var require_id = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/id.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _node = _interopRequireDefault(require_node2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var ID = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(ID2, _Node);
      function ID2(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        _this.type = _types.ID;
        return _this;
      }
      var _proto = ID2.prototype;
      _proto.valueToString = function valueToString() {
        return "#" + _Node.prototype.valueToString.call(this);
      };
      return ID2;
    }(_node["default"]);
    exports["default"] = ID;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/namespace.js
var require_namespace = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/namespace.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _cssesc = _interopRequireDefault(require_cssesc());
    var _util = require_util2();
    var _node = _interopRequireDefault(require_node2());
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Namespace = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(Namespace2, _Node);
      function Namespace2() {
        return _Node.apply(this, arguments) || this;
      }
      var _proto = Namespace2.prototype;
      _proto.qualifiedName = function qualifiedName(value) {
        if (this.namespace) {
          return this.namespaceString + "|" + value;
        } else {
          return value;
        }
      };
      _proto.valueToString = function valueToString() {
        return this.qualifiedName(_Node.prototype.valueToString.call(this));
      };
      _createClass(Namespace2, [{
        key: "namespace",
        get: function get() {
          return this._namespace;
        },
        set: function set(namespace) {
          if (namespace === true || namespace === "*" || namespace === "&") {
            this._namespace = namespace;
            if (this.raws) {
              delete this.raws.namespace;
            }
            return;
          }
          var escaped = (0, _cssesc["default"])(namespace, {
            isIdentifier: true
          });
          this._namespace = namespace;
          if (escaped !== namespace) {
            (0, _util.ensureObject)(this, "raws");
            this.raws.namespace = escaped;
          } else if (this.raws) {
            delete this.raws.namespace;
          }
        }
      }, {
        key: "ns",
        get: function get() {
          return this._namespace;
        },
        set: function set(namespace) {
          this.namespace = namespace;
        }
      }, {
        key: "namespaceString",
        get: function get() {
          if (this.namespace) {
            var ns = this.stringifyProperty("namespace");
            if (ns === true) {
              return "";
            } else {
              return ns;
            }
          } else {
            return "";
          }
        }
      }]);
      return Namespace2;
    }(_node["default"]);
    exports["default"] = Namespace;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/tag.js
var require_tag = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/tag.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _namespace = _interopRequireDefault(require_namespace());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Tag = /* @__PURE__ */ function(_Namespace) {
      _inheritsLoose(Tag2, _Namespace);
      function Tag2(opts) {
        var _this;
        _this = _Namespace.call(this, opts) || this;
        _this.type = _types.TAG;
        return _this;
      }
      return Tag2;
    }(_namespace["default"]);
    exports["default"] = Tag;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/string.js
var require_string = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/string.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _node = _interopRequireDefault(require_node2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var String2 = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(String3, _Node);
      function String3(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        _this.type = _types.STRING;
        return _this;
      }
      return String3;
    }(_node["default"]);
    exports["default"] = String2;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/pseudo.js
var require_pseudo = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/pseudo.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _container = _interopRequireDefault(require_container2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Pseudo = /* @__PURE__ */ function(_Container) {
      _inheritsLoose(Pseudo2, _Container);
      function Pseudo2(opts) {
        var _this;
        _this = _Container.call(this, opts) || this;
        _this.type = _types.PSEUDO;
        return _this;
      }
      var _proto = Pseudo2.prototype;
      _proto.toString = function toString() {
        var params = this.length ? "(" + this.map(String).join(",") + ")" : "";
        return [this.rawSpaceBefore, this.stringifyProperty("value"), params, this.rawSpaceAfter].join("");
      };
      return Pseudo2;
    }(_container["default"]);
    exports["default"] = Pseudo;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/util-deprecate@1.0.2/node_modules/util-deprecate/node.js
var require_node3 = __commonJS({
  "node_modules/.pnpm/util-deprecate@1.0.2/node_modules/util-deprecate/node.js"(exports, module2) {
    module2.exports = require("util").deprecate;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/attribute.js
var require_attribute = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/attribute.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    exports.unescapeValue = unescapeValue;
    var _cssesc = _interopRequireDefault(require_cssesc());
    var _unesc = _interopRequireDefault(require_unesc());
    var _namespace = _interopRequireDefault(require_namespace());
    var _types = require_types();
    var _CSSESC_QUOTE_OPTIONS;
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var deprecate = require_node3();
    var WRAPPED_IN_QUOTES = /^('|")([^]*)\1$/;
    var warnOfDeprecatedValueAssignment = deprecate(function() {
    }, "Assigning an attribute a value containing characters that might need to be escaped is deprecated. Call attribute.setValue() instead.");
    var warnOfDeprecatedQuotedAssignment = deprecate(function() {
    }, "Assigning attr.quoted is deprecated and has no effect. Assign to attr.quoteMark instead.");
    var warnOfDeprecatedConstructor = deprecate(function() {
    }, "Constructing an Attribute selector with a value without specifying quoteMark is deprecated. Note: The value should be unescaped now.");
    function unescapeValue(value) {
      var deprecatedUsage = false;
      var quoteMark = null;
      var unescaped = value;
      var m = unescaped.match(WRAPPED_IN_QUOTES);
      if (m) {
        quoteMark = m[1];
        unescaped = m[2];
      }
      unescaped = (0, _unesc["default"])(unescaped);
      if (unescaped !== value) {
        deprecatedUsage = true;
      }
      return {
        deprecatedUsage,
        unescaped,
        quoteMark
      };
    }
    function handleDeprecatedContructorOpts(opts) {
      if (opts.quoteMark !== void 0) {
        return opts;
      }
      if (opts.value === void 0) {
        return opts;
      }
      warnOfDeprecatedConstructor();
      var _unescapeValue = unescapeValue(opts.value), quoteMark = _unescapeValue.quoteMark, unescaped = _unescapeValue.unescaped;
      if (!opts.raws) {
        opts.raws = {};
      }
      if (opts.raws.value === void 0) {
        opts.raws.value = opts.value;
      }
      opts.value = unescaped;
      opts.quoteMark = quoteMark;
      return opts;
    }
    var Attribute = /* @__PURE__ */ function(_Namespace) {
      _inheritsLoose(Attribute2, _Namespace);
      function Attribute2(opts) {
        var _this;
        if (opts === void 0) {
          opts = {};
        }
        _this = _Namespace.call(this, handleDeprecatedContructorOpts(opts)) || this;
        _this.type = _types.ATTRIBUTE;
        _this.raws = _this.raws || {};
        Object.defineProperty(_this.raws, "unquoted", {
          get: deprecate(function() {
            return _this.value;
          }, "attr.raws.unquoted is deprecated. Call attr.value instead."),
          set: deprecate(function() {
            return _this.value;
          }, "Setting attr.raws.unquoted is deprecated and has no effect. attr.value is unescaped by default now.")
        });
        _this._constructed = true;
        return _this;
      }
      var _proto = Attribute2.prototype;
      _proto.getQuotedValue = function getQuotedValue(options) {
        if (options === void 0) {
          options = {};
        }
        var quoteMark = this._determineQuoteMark(options);
        var cssescopts = CSSESC_QUOTE_OPTIONS[quoteMark];
        var escaped = (0, _cssesc["default"])(this._value, cssescopts);
        return escaped;
      };
      _proto._determineQuoteMark = function _determineQuoteMark(options) {
        return options.smart ? this.smartQuoteMark(options) : this.preferredQuoteMark(options);
      };
      _proto.setValue = function setValue(value, options) {
        if (options === void 0) {
          options = {};
        }
        this._value = value;
        this._quoteMark = this._determineQuoteMark(options);
        this._syncRawValue();
      };
      _proto.smartQuoteMark = function smartQuoteMark(options) {
        var v = this.value;
        var numSingleQuotes = v.replace(/[^']/g, "").length;
        var numDoubleQuotes = v.replace(/[^"]/g, "").length;
        if (numSingleQuotes + numDoubleQuotes === 0) {
          var escaped = (0, _cssesc["default"])(v, {
            isIdentifier: true
          });
          if (escaped === v) {
            return Attribute2.NO_QUOTE;
          } else {
            var pref = this.preferredQuoteMark(options);
            if (pref === Attribute2.NO_QUOTE) {
              var quote = this.quoteMark || options.quoteMark || Attribute2.DOUBLE_QUOTE;
              var opts = CSSESC_QUOTE_OPTIONS[quote];
              var quoteValue = (0, _cssesc["default"])(v, opts);
              if (quoteValue.length < escaped.length) {
                return quote;
              }
            }
            return pref;
          }
        } else if (numDoubleQuotes === numSingleQuotes) {
          return this.preferredQuoteMark(options);
        } else if (numDoubleQuotes < numSingleQuotes) {
          return Attribute2.DOUBLE_QUOTE;
        } else {
          return Attribute2.SINGLE_QUOTE;
        }
      };
      _proto.preferredQuoteMark = function preferredQuoteMark(options) {
        var quoteMark = options.preferCurrentQuoteMark ? this.quoteMark : options.quoteMark;
        if (quoteMark === void 0) {
          quoteMark = options.preferCurrentQuoteMark ? options.quoteMark : this.quoteMark;
        }
        if (quoteMark === void 0) {
          quoteMark = Attribute2.DOUBLE_QUOTE;
        }
        return quoteMark;
      };
      _proto._syncRawValue = function _syncRawValue() {
        var rawValue = (0, _cssesc["default"])(this._value, CSSESC_QUOTE_OPTIONS[this.quoteMark]);
        if (rawValue === this._value) {
          if (this.raws) {
            delete this.raws.value;
          }
        } else {
          this.raws.value = rawValue;
        }
      };
      _proto._handleEscapes = function _handleEscapes(prop, value) {
        if (this._constructed) {
          var escaped = (0, _cssesc["default"])(value, {
            isIdentifier: true
          });
          if (escaped !== value) {
            this.raws[prop] = escaped;
          } else {
            delete this.raws[prop];
          }
        }
      };
      _proto._spacesFor = function _spacesFor(name) {
        var attrSpaces = {
          before: "",
          after: ""
        };
        var spaces = this.spaces[name] || {};
        var rawSpaces = this.raws.spaces && this.raws.spaces[name] || {};
        return Object.assign(attrSpaces, spaces, rawSpaces);
      };
      _proto._stringFor = function _stringFor(name, spaceName, concat) {
        if (spaceName === void 0) {
          spaceName = name;
        }
        if (concat === void 0) {
          concat = defaultAttrConcat;
        }
        var attrSpaces = this._spacesFor(spaceName);
        return concat(this.stringifyProperty(name), attrSpaces);
      };
      _proto.offsetOf = function offsetOf(name) {
        var count = 1;
        var attributeSpaces = this._spacesFor("attribute");
        count += attributeSpaces.before.length;
        if (name === "namespace" || name === "ns") {
          return this.namespace ? count : -1;
        }
        if (name === "attributeNS") {
          return count;
        }
        count += this.namespaceString.length;
        if (this.namespace) {
          count += 1;
        }
        if (name === "attribute") {
          return count;
        }
        count += this.stringifyProperty("attribute").length;
        count += attributeSpaces.after.length;
        var operatorSpaces = this._spacesFor("operator");
        count += operatorSpaces.before.length;
        var operator = this.stringifyProperty("operator");
        if (name === "operator") {
          return operator ? count : -1;
        }
        count += operator.length;
        count += operatorSpaces.after.length;
        var valueSpaces = this._spacesFor("value");
        count += valueSpaces.before.length;
        var value = this.stringifyProperty("value");
        if (name === "value") {
          return value ? count : -1;
        }
        count += value.length;
        count += valueSpaces.after.length;
        var insensitiveSpaces = this._spacesFor("insensitive");
        count += insensitiveSpaces.before.length;
        if (name === "insensitive") {
          return this.insensitive ? count : -1;
        }
        return -1;
      };
      _proto.toString = function toString() {
        var _this2 = this;
        var selector = [this.rawSpaceBefore, "["];
        selector.push(this._stringFor("qualifiedAttribute", "attribute"));
        if (this.operator && (this.value || this.value === "")) {
          selector.push(this._stringFor("operator"));
          selector.push(this._stringFor("value"));
          selector.push(this._stringFor("insensitiveFlag", "insensitive", function(attrValue, attrSpaces) {
            if (attrValue.length > 0 && !_this2.quoted && attrSpaces.before.length === 0 && !(_this2.spaces.value && _this2.spaces.value.after)) {
              attrSpaces.before = " ";
            }
            return defaultAttrConcat(attrValue, attrSpaces);
          }));
        }
        selector.push("]");
        selector.push(this.rawSpaceAfter);
        return selector.join("");
      };
      _createClass(Attribute2, [{
        key: "quoted",
        get: function get() {
          var qm = this.quoteMark;
          return qm === "'" || qm === '"';
        },
        set: function set(value) {
          warnOfDeprecatedQuotedAssignment();
        }
        /**
         * returns a single (`'`) or double (`"`) quote character if the value is quoted.
         * returns `null` if the value is not quoted.
         * returns `undefined` if the quotation state is unknown (this can happen when
         * the attribute is constructed without specifying a quote mark.)
         */
      }, {
        key: "quoteMark",
        get: function get() {
          return this._quoteMark;
        },
        set: function set(quoteMark) {
          if (!this._constructed) {
            this._quoteMark = quoteMark;
            return;
          }
          if (this._quoteMark !== quoteMark) {
            this._quoteMark = quoteMark;
            this._syncRawValue();
          }
        }
      }, {
        key: "qualifiedAttribute",
        get: function get() {
          return this.qualifiedName(this.raws.attribute || this.attribute);
        }
      }, {
        key: "insensitiveFlag",
        get: function get() {
          return this.insensitive ? "i" : "";
        }
      }, {
        key: "value",
        get: function get() {
          return this._value;
        },
        set: (
          /**
           * Before 3.0, the value had to be set to an escaped value including any wrapped
           * quote marks. In 3.0, the semantics of `Attribute.value` changed so that the value
           * is unescaped during parsing and any quote marks are removed.
           *
           * Because the ambiguity of this semantic change, if you set `attr.value = newValue`,
           * a deprecation warning is raised when the new value contains any characters that would
           * require escaping (including if it contains wrapped quotes).
           *
           * Instead, you should call `attr.setValue(newValue, opts)` and pass options that describe
           * how the new value is quoted.
           */
          function set(v) {
            if (this._constructed) {
              var _unescapeValue2 = unescapeValue(v), deprecatedUsage = _unescapeValue2.deprecatedUsage, unescaped = _unescapeValue2.unescaped, quoteMark = _unescapeValue2.quoteMark;
              if (deprecatedUsage) {
                warnOfDeprecatedValueAssignment();
              }
              if (unescaped === this._value && quoteMark === this._quoteMark) {
                return;
              }
              this._value = unescaped;
              this._quoteMark = quoteMark;
              this._syncRawValue();
            } else {
              this._value = v;
            }
          }
        )
      }, {
        key: "insensitive",
        get: function get() {
          return this._insensitive;
        },
        set: function set(insensitive) {
          if (!insensitive) {
            this._insensitive = false;
            if (this.raws && (this.raws.insensitiveFlag === "I" || this.raws.insensitiveFlag === "i")) {
              this.raws.insensitiveFlag = void 0;
            }
          }
          this._insensitive = insensitive;
        }
      }, {
        key: "attribute",
        get: function get() {
          return this._attribute;
        },
        set: function set(name) {
          this._handleEscapes("attribute", name);
          this._attribute = name;
        }
      }]);
      return Attribute2;
    }(_namespace["default"]);
    exports["default"] = Attribute;
    Attribute.NO_QUOTE = null;
    Attribute.SINGLE_QUOTE = "'";
    Attribute.DOUBLE_QUOTE = '"';
    var CSSESC_QUOTE_OPTIONS = (_CSSESC_QUOTE_OPTIONS = {
      "'": {
        quotes: "single",
        wrap: true
      },
      '"': {
        quotes: "double",
        wrap: true
      }
    }, _CSSESC_QUOTE_OPTIONS[null] = {
      isIdentifier: true
    }, _CSSESC_QUOTE_OPTIONS);
    function defaultAttrConcat(attrValue, attrSpaces) {
      return "" + attrSpaces.before + attrValue + attrSpaces.after;
    }
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/universal.js
var require_universal = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/universal.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _namespace = _interopRequireDefault(require_namespace());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Universal = /* @__PURE__ */ function(_Namespace) {
      _inheritsLoose(Universal2, _Namespace);
      function Universal2(opts) {
        var _this;
        _this = _Namespace.call(this, opts) || this;
        _this.type = _types.UNIVERSAL;
        _this.value = "*";
        return _this;
      }
      return Universal2;
    }(_namespace["default"]);
    exports["default"] = Universal;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/combinator.js
var require_combinator = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/combinator.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _node = _interopRequireDefault(require_node2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Combinator = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(Combinator2, _Node);
      function Combinator2(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        _this.type = _types.COMBINATOR;
        return _this;
      }
      return Combinator2;
    }(_node["default"]);
    exports["default"] = Combinator;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/nesting.js
var require_nesting = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/nesting.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _node = _interopRequireDefault(require_node2());
    var _types = require_types();
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _inheritsLoose(subClass, superClass) {
      subClass.prototype = Object.create(superClass.prototype);
      subClass.prototype.constructor = subClass;
      _setPrototypeOf(subClass, superClass);
    }
    function _setPrototypeOf(o, p) {
      _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function _setPrototypeOf2(o2, p2) {
        o2.__proto__ = p2;
        return o2;
      };
      return _setPrototypeOf(o, p);
    }
    var Nesting = /* @__PURE__ */ function(_Node) {
      _inheritsLoose(Nesting2, _Node);
      function Nesting2(opts) {
        var _this;
        _this = _Node.call(this, opts) || this;
        _this.type = _types.NESTING;
        _this.value = "&";
        return _this;
      }
      return Nesting2;
    }(_node["default"]);
    exports["default"] = Nesting;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/sortAscending.js
var require_sortAscending = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/sortAscending.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = sortAscending;
    function sortAscending(list) {
      return list.sort(function(a, b) {
        return a - b;
      });
    }
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/tokenTypes.js
var require_tokenTypes = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/tokenTypes.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports.word = exports.tilde = exports.tab = exports.str = exports.space = exports.slash = exports.singleQuote = exports.semicolon = exports.plus = exports.pipe = exports.openSquare = exports.openParenthesis = exports.newline = exports.greaterThan = exports.feed = exports.equals = exports.doubleQuote = exports.dollar = exports.cr = exports.comment = exports.comma = exports.combinator = exports.colon = exports.closeSquare = exports.closeParenthesis = exports.caret = exports.bang = exports.backslash = exports.at = exports.asterisk = exports.ampersand = void 0;
    var ampersand = 38;
    exports.ampersand = ampersand;
    var asterisk = 42;
    exports.asterisk = asterisk;
    var at = 64;
    exports.at = at;
    var comma = 44;
    exports.comma = comma;
    var colon = 58;
    exports.colon = colon;
    var semicolon = 59;
    exports.semicolon = semicolon;
    var openParenthesis = 40;
    exports.openParenthesis = openParenthesis;
    var closeParenthesis = 41;
    exports.closeParenthesis = closeParenthesis;
    var openSquare = 91;
    exports.openSquare = openSquare;
    var closeSquare = 93;
    exports.closeSquare = closeSquare;
    var dollar = 36;
    exports.dollar = dollar;
    var tilde = 126;
    exports.tilde = tilde;
    var caret = 94;
    exports.caret = caret;
    var plus = 43;
    exports.plus = plus;
    var equals = 61;
    exports.equals = equals;
    var pipe = 124;
    exports.pipe = pipe;
    var greaterThan = 62;
    exports.greaterThan = greaterThan;
    var space = 32;
    exports.space = space;
    var singleQuote = 39;
    exports.singleQuote = singleQuote;
    var doubleQuote = 34;
    exports.doubleQuote = doubleQuote;
    var slash = 47;
    exports.slash = slash;
    var bang = 33;
    exports.bang = bang;
    var backslash = 92;
    exports.backslash = backslash;
    var cr = 13;
    exports.cr = cr;
    var feed = 12;
    exports.feed = feed;
    var newline = 10;
    exports.newline = newline;
    var tab = 9;
    exports.tab = tab;
    var str = singleQuote;
    exports.str = str;
    var comment = -1;
    exports.comment = comment;
    var word = -2;
    exports.word = word;
    var combinator = -3;
    exports.combinator = combinator;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/tokenize.js
var require_tokenize2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/tokenize.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports.FIELDS = void 0;
    exports["default"] = tokenize;
    var t = _interopRequireWildcard(require_tokenTypes());
    var _unescapable;
    var _wordDelimiters;
    function _getRequireWildcardCache(nodeInterop) {
      if (typeof WeakMap !== "function")
        return null;
      var cacheBabelInterop = /* @__PURE__ */ new WeakMap();
      var cacheNodeInterop = /* @__PURE__ */ new WeakMap();
      return (_getRequireWildcardCache = function _getRequireWildcardCache2(nodeInterop2) {
        return nodeInterop2 ? cacheNodeInterop : cacheBabelInterop;
      })(nodeInterop);
    }
    function _interopRequireWildcard(obj, nodeInterop) {
      if (!nodeInterop && obj && obj.__esModule) {
        return obj;
      }
      if (obj === null || typeof obj !== "object" && typeof obj !== "function") {
        return { "default": obj };
      }
      var cache = _getRequireWildcardCache(nodeInterop);
      if (cache && cache.has(obj)) {
        return cache.get(obj);
      }
      var newObj = {};
      var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor;
      for (var key in obj) {
        if (key !== "default" && Object.prototype.hasOwnProperty.call(obj, key)) {
          var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null;
          if (desc && (desc.get || desc.set)) {
            Object.defineProperty(newObj, key, desc);
          } else {
            newObj[key] = obj[key];
          }
        }
      }
      newObj["default"] = obj;
      if (cache) {
        cache.set(obj, newObj);
      }
      return newObj;
    }
    var unescapable = (_unescapable = {}, _unescapable[t.tab] = true, _unescapable[t.newline] = true, _unescapable[t.cr] = true, _unescapable[t.feed] = true, _unescapable);
    var wordDelimiters = (_wordDelimiters = {}, _wordDelimiters[t.space] = true, _wordDelimiters[t.tab] = true, _wordDelimiters[t.newline] = true, _wordDelimiters[t.cr] = true, _wordDelimiters[t.feed] = true, _wordDelimiters[t.ampersand] = true, _wordDelimiters[t.asterisk] = true, _wordDelimiters[t.bang] = true, _wordDelimiters[t.comma] = true, _wordDelimiters[t.colon] = true, _wordDelimiters[t.semicolon] = true, _wordDelimiters[t.openParenthesis] = true, _wordDelimiters[t.closeParenthesis] = true, _wordDelimiters[t.openSquare] = true, _wordDelimiters[t.closeSquare] = true, _wordDelimiters[t.singleQuote] = true, _wordDelimiters[t.doubleQuote] = true, _wordDelimiters[t.plus] = true, _wordDelimiters[t.pipe] = true, _wordDelimiters[t.tilde] = true, _wordDelimiters[t.greaterThan] = true, _wordDelimiters[t.equals] = true, _wordDelimiters[t.dollar] = true, _wordDelimiters[t.caret] = true, _wordDelimiters[t.slash] = true, _wordDelimiters);
    var hex = {};
    var hexChars = "0123456789abcdefABCDEF";
    for (i = 0; i < hexChars.length; i++) {
      hex[hexChars.charCodeAt(i)] = true;
    }
    var i;
    function consumeWord(css, start) {
      var next = start;
      var code;
      do {
        code = css.charCodeAt(next);
        if (wordDelimiters[code]) {
          return next - 1;
        } else if (code === t.backslash) {
          next = consumeEscape(css, next) + 1;
        } else {
          next++;
        }
      } while (next < css.length);
      return next - 1;
    }
    function consumeEscape(css, start) {
      var next = start;
      var code = css.charCodeAt(next + 1);
      if (unescapable[code]) {
      } else if (hex[code]) {
        var hexDigits = 0;
        do {
          next++;
          hexDigits++;
          code = css.charCodeAt(next + 1);
        } while (hex[code] && hexDigits < 6);
        if (hexDigits < 6 && code === t.space) {
          next++;
        }
      } else {
        next++;
      }
      return next;
    }
    var FIELDS = {
      TYPE: 0,
      START_LINE: 1,
      START_COL: 2,
      END_LINE: 3,
      END_COL: 4,
      START_POS: 5,
      END_POS: 6
    };
    exports.FIELDS = FIELDS;
    function tokenize(input) {
      var tokens = [];
      var css = input.css.valueOf();
      var _css = css, length = _css.length;
      var offset = -1;
      var line = 1;
      var start = 0;
      var end = 0;
      var code, content, endColumn, endLine, escaped, escapePos, last, lines, next, nextLine, nextOffset, quote, tokenType;
      function unclosed(what, fix) {
        if (input.safe) {
          css += fix;
          next = css.length - 1;
        } else {
          throw input.error("Unclosed " + what, line, start - offset, start);
        }
      }
      while (start < length) {
        code = css.charCodeAt(start);
        if (code === t.newline) {
          offset = start;
          line += 1;
        }
        switch (code) {
          case t.space:
          case t.tab:
          case t.newline:
          case t.cr:
          case t.feed:
            next = start;
            do {
              next += 1;
              code = css.charCodeAt(next);
              if (code === t.newline) {
                offset = next;
                line += 1;
              }
            } while (code === t.space || code === t.newline || code === t.tab || code === t.cr || code === t.feed);
            tokenType = t.space;
            endLine = line;
            endColumn = next - offset - 1;
            end = next;
            break;
          case t.plus:
          case t.greaterThan:
          case t.tilde:
          case t.pipe:
            next = start;
            do {
              next += 1;
              code = css.charCodeAt(next);
            } while (code === t.plus || code === t.greaterThan || code === t.tilde || code === t.pipe);
            tokenType = t.combinator;
            endLine = line;
            endColumn = start - offset;
            end = next;
            break;
          case t.asterisk:
          case t.ampersand:
          case t.bang:
          case t.comma:
          case t.equals:
          case t.dollar:
          case t.caret:
          case t.openSquare:
          case t.closeSquare:
          case t.colon:
          case t.semicolon:
          case t.openParenthesis:
          case t.closeParenthesis:
            next = start;
            tokenType = code;
            endLine = line;
            endColumn = start - offset;
            end = next + 1;
            break;
          case t.singleQuote:
          case t.doubleQuote:
            quote = code === t.singleQuote ? "'" : '"';
            next = start;
            do {
              escaped = false;
              next = css.indexOf(quote, next + 1);
              if (next === -1) {
                unclosed("quote", quote);
              }
              escapePos = next;
              while (css.charCodeAt(escapePos - 1) === t.backslash) {
                escapePos -= 1;
                escaped = !escaped;
              }
            } while (escaped);
            tokenType = t.str;
            endLine = line;
            endColumn = start - offset;
            end = next + 1;
            break;
          default:
            if (code === t.slash && css.charCodeAt(start + 1) === t.asterisk) {
              next = css.indexOf("*/", start + 2) + 1;
              if (next === 0) {
                unclosed("comment", "*/");
              }
              content = css.slice(start, next + 1);
              lines = content.split("\n");
              last = lines.length - 1;
              if (last > 0) {
                nextLine = line + last;
                nextOffset = next - lines[last].length;
              } else {
                nextLine = line;
                nextOffset = offset;
              }
              tokenType = t.comment;
              line = nextLine;
              endLine = nextLine;
              endColumn = next - nextOffset;
            } else if (code === t.slash) {
              next = start;
              tokenType = code;
              endLine = line;
              endColumn = start - offset;
              end = next + 1;
            } else {
              next = consumeWord(css, start);
              tokenType = t.word;
              endLine = line;
              endColumn = next - offset;
            }
            end = next + 1;
            break;
        }
        tokens.push([
          tokenType,
          // [0] Token type
          line,
          // [1] Starting line
          start - offset,
          // [2] Starting column
          endLine,
          // [3] Ending line
          endColumn,
          // [4] Ending column
          start,
          // [5] Start position / Source index
          end
          // [6] End position
        ]);
        if (nextOffset) {
          offset = nextOffset;
          nextOffset = null;
        }
        start = end;
      }
      return tokens;
    }
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/parser.js
var require_parser3 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/parser.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _root = _interopRequireDefault(require_root2());
    var _selector = _interopRequireDefault(require_selector());
    var _className = _interopRequireDefault(require_className());
    var _comment = _interopRequireDefault(require_comment2());
    var _id = _interopRequireDefault(require_id());
    var _tag = _interopRequireDefault(require_tag());
    var _string = _interopRequireDefault(require_string());
    var _pseudo = _interopRequireDefault(require_pseudo());
    var _attribute = _interopRequireWildcard(require_attribute());
    var _universal = _interopRequireDefault(require_universal());
    var _combinator = _interopRequireDefault(require_combinator());
    var _nesting = _interopRequireDefault(require_nesting());
    var _sortAscending = _interopRequireDefault(require_sortAscending());
    var _tokenize = _interopRequireWildcard(require_tokenize2());
    var tokens = _interopRequireWildcard(require_tokenTypes());
    var types = _interopRequireWildcard(require_types());
    var _util = require_util2();
    var _WHITESPACE_TOKENS;
    var _Object$assign;
    function _getRequireWildcardCache(nodeInterop) {
      if (typeof WeakMap !== "function")
        return null;
      var cacheBabelInterop = /* @__PURE__ */ new WeakMap();
      var cacheNodeInterop = /* @__PURE__ */ new WeakMap();
      return (_getRequireWildcardCache = function _getRequireWildcardCache2(nodeInterop2) {
        return nodeInterop2 ? cacheNodeInterop : cacheBabelInterop;
      })(nodeInterop);
    }
    function _interopRequireWildcard(obj, nodeInterop) {
      if (!nodeInterop && obj && obj.__esModule) {
        return obj;
      }
      if (obj === null || typeof obj !== "object" && typeof obj !== "function") {
        return { "default": obj };
      }
      var cache = _getRequireWildcardCache(nodeInterop);
      if (cache && cache.has(obj)) {
        return cache.get(obj);
      }
      var newObj = {};
      var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor;
      for (var key in obj) {
        if (key !== "default" && Object.prototype.hasOwnProperty.call(obj, key)) {
          var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null;
          if (desc && (desc.get || desc.set)) {
            Object.defineProperty(newObj, key, desc);
          } else {
            newObj[key] = obj[key];
          }
        }
      }
      newObj["default"] = obj;
      if (cache) {
        cache.set(obj, newObj);
      }
      return newObj;
    }
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    function _defineProperties(target, props) {
      for (var i = 0; i < props.length; i++) {
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor)
          descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
      }
    }
    function _createClass(Constructor, protoProps, staticProps) {
      if (protoProps)
        _defineProperties(Constructor.prototype, protoProps);
      if (staticProps)
        _defineProperties(Constructor, staticProps);
      Object.defineProperty(Constructor, "prototype", { writable: false });
      return Constructor;
    }
    var WHITESPACE_TOKENS = (_WHITESPACE_TOKENS = {}, _WHITESPACE_TOKENS[tokens.space] = true, _WHITESPACE_TOKENS[tokens.cr] = true, _WHITESPACE_TOKENS[tokens.feed] = true, _WHITESPACE_TOKENS[tokens.newline] = true, _WHITESPACE_TOKENS[tokens.tab] = true, _WHITESPACE_TOKENS);
    var WHITESPACE_EQUIV_TOKENS = Object.assign({}, WHITESPACE_TOKENS, (_Object$assign = {}, _Object$assign[tokens.comment] = true, _Object$assign));
    function tokenStart(token) {
      return {
        line: token[_tokenize.FIELDS.START_LINE],
        column: token[_tokenize.FIELDS.START_COL]
      };
    }
    function tokenEnd(token) {
      return {
        line: token[_tokenize.FIELDS.END_LINE],
        column: token[_tokenize.FIELDS.END_COL]
      };
    }
    function getSource(startLine, startColumn, endLine, endColumn) {
      return {
        start: {
          line: startLine,
          column: startColumn
        },
        end: {
          line: endLine,
          column: endColumn
        }
      };
    }
    function getTokenSource(token) {
      return getSource(token[_tokenize.FIELDS.START_LINE], token[_tokenize.FIELDS.START_COL], token[_tokenize.FIELDS.END_LINE], token[_tokenize.FIELDS.END_COL]);
    }
    function getTokenSourceSpan(startToken, endToken) {
      if (!startToken) {
        return void 0;
      }
      return getSource(startToken[_tokenize.FIELDS.START_LINE], startToken[_tokenize.FIELDS.START_COL], endToken[_tokenize.FIELDS.END_LINE], endToken[_tokenize.FIELDS.END_COL]);
    }
    function unescapeProp(node, prop) {
      var value = node[prop];
      if (typeof value !== "string") {
        return;
      }
      if (value.indexOf("\\") !== -1) {
        (0, _util.ensureObject)(node, "raws");
        node[prop] = (0, _util.unesc)(value);
        if (node.raws[prop] === void 0) {
          node.raws[prop] = value;
        }
      }
      return node;
    }
    function indexesOf(array, item) {
      var i = -1;
      var indexes = [];
      while ((i = array.indexOf(item, i + 1)) !== -1) {
        indexes.push(i);
      }
      return indexes;
    }
    function uniqs() {
      var list = Array.prototype.concat.apply([], arguments);
      return list.filter(function(item, i) {
        return i === list.indexOf(item);
      });
    }
    var Parser = /* @__PURE__ */ function() {
      function Parser2(rule, options) {
        if (options === void 0) {
          options = {};
        }
        this.rule = rule;
        this.options = Object.assign({
          lossy: false,
          safe: false
        }, options);
        this.position = 0;
        this.css = typeof this.rule === "string" ? this.rule : this.rule.selector;
        this.tokens = (0, _tokenize["default"])({
          css: this.css,
          error: this._errorGenerator(),
          safe: this.options.safe
        });
        var rootSource = getTokenSourceSpan(this.tokens[0], this.tokens[this.tokens.length - 1]);
        this.root = new _root["default"]({
          source: rootSource
        });
        this.root.errorGenerator = this._errorGenerator();
        var selector = new _selector["default"]({
          source: {
            start: {
              line: 1,
              column: 1
            }
          }
        });
        this.root.append(selector);
        this.current = selector;
        this.loop();
      }
      var _proto = Parser2.prototype;
      _proto._errorGenerator = function _errorGenerator() {
        var _this = this;
        return function(message, errorOptions) {
          if (typeof _this.rule === "string") {
            return new Error(message);
          }
          return _this.rule.error(message, errorOptions);
        };
      };
      _proto.attribute = function attribute() {
        var attr = [];
        var startingToken = this.currToken;
        this.position++;
        while (this.position < this.tokens.length && this.currToken[_tokenize.FIELDS.TYPE] !== tokens.closeSquare) {
          attr.push(this.currToken);
          this.position++;
        }
        if (this.currToken[_tokenize.FIELDS.TYPE] !== tokens.closeSquare) {
          return this.expected("closing square bracket", this.currToken[_tokenize.FIELDS.START_POS]);
        }
        var len = attr.length;
        var node = {
          source: getSource(startingToken[1], startingToken[2], this.currToken[3], this.currToken[4]),
          sourceIndex: startingToken[_tokenize.FIELDS.START_POS]
        };
        if (len === 1 && !~[tokens.word].indexOf(attr[0][_tokenize.FIELDS.TYPE])) {
          return this.expected("attribute", attr[0][_tokenize.FIELDS.START_POS]);
        }
        var pos = 0;
        var spaceBefore = "";
        var commentBefore = "";
        var lastAdded = null;
        var spaceAfterMeaningfulToken = false;
        while (pos < len) {
          var token = attr[pos];
          var content = this.content(token);
          var next = attr[pos + 1];
          switch (token[_tokenize.FIELDS.TYPE]) {
            case tokens.space:
              spaceAfterMeaningfulToken = true;
              if (this.options.lossy) {
                break;
              }
              if (lastAdded) {
                (0, _util.ensureObject)(node, "spaces", lastAdded);
                var prevContent = node.spaces[lastAdded].after || "";
                node.spaces[lastAdded].after = prevContent + content;
                var existingComment = (0, _util.getProp)(node, "raws", "spaces", lastAdded, "after") || null;
                if (existingComment) {
                  node.raws.spaces[lastAdded].after = existingComment + content;
                }
              } else {
                spaceBefore = spaceBefore + content;
                commentBefore = commentBefore + content;
              }
              break;
            case tokens.asterisk:
              if (next[_tokenize.FIELDS.TYPE] === tokens.equals) {
                node.operator = content;
                lastAdded = "operator";
              } else if ((!node.namespace || lastAdded === "namespace" && !spaceAfterMeaningfulToken) && next) {
                if (spaceBefore) {
                  (0, _util.ensureObject)(node, "spaces", "attribute");
                  node.spaces.attribute.before = spaceBefore;
                  spaceBefore = "";
                }
                if (commentBefore) {
                  (0, _util.ensureObject)(node, "raws", "spaces", "attribute");
                  node.raws.spaces.attribute.before = spaceBefore;
                  commentBefore = "";
                }
                node.namespace = (node.namespace || "") + content;
                var rawValue = (0, _util.getProp)(node, "raws", "namespace") || null;
                if (rawValue) {
                  node.raws.namespace += content;
                }
                lastAdded = "namespace";
              }
              spaceAfterMeaningfulToken = false;
              break;
            case tokens.dollar:
              if (lastAdded === "value") {
                var oldRawValue = (0, _util.getProp)(node, "raws", "value");
                node.value += "$";
                if (oldRawValue) {
                  node.raws.value = oldRawValue + "$";
                }
                break;
              }
            case tokens.caret:
              if (next[_tokenize.FIELDS.TYPE] === tokens.equals) {
                node.operator = content;
                lastAdded = "operator";
              }
              spaceAfterMeaningfulToken = false;
              break;
            case tokens.combinator:
              if (content === "~" && next[_tokenize.FIELDS.TYPE] === tokens.equals) {
                node.operator = content;
                lastAdded = "operator";
              }
              if (content !== "|") {
                spaceAfterMeaningfulToken = false;
                break;
              }
              if (next[_tokenize.FIELDS.TYPE] === tokens.equals) {
                node.operator = content;
                lastAdded = "operator";
              } else if (!node.namespace && !node.attribute) {
                node.namespace = true;
              }
              spaceAfterMeaningfulToken = false;
              break;
            case tokens.word:
              if (next && this.content(next) === "|" && attr[pos + 2] && attr[pos + 2][_tokenize.FIELDS.TYPE] !== tokens.equals && // this look-ahead probably fails with comment nodes involved.
              !node.operator && !node.namespace) {
                node.namespace = content;
                lastAdded = "namespace";
              } else if (!node.attribute || lastAdded === "attribute" && !spaceAfterMeaningfulToken) {
                if (spaceBefore) {
                  (0, _util.ensureObject)(node, "spaces", "attribute");
                  node.spaces.attribute.before = spaceBefore;
                  spaceBefore = "";
                }
                if (commentBefore) {
                  (0, _util.ensureObject)(node, "raws", "spaces", "attribute");
                  node.raws.spaces.attribute.before = commentBefore;
                  commentBefore = "";
                }
                node.attribute = (node.attribute || "") + content;
                var _rawValue = (0, _util.getProp)(node, "raws", "attribute") || null;
                if (_rawValue) {
                  node.raws.attribute += content;
                }
                lastAdded = "attribute";
              } else if (!node.value && node.value !== "" || lastAdded === "value" && !(spaceAfterMeaningfulToken || node.quoteMark)) {
                var _unescaped = (0, _util.unesc)(content);
                var _oldRawValue = (0, _util.getProp)(node, "raws", "value") || "";
                var oldValue = node.value || "";
                node.value = oldValue + _unescaped;
                node.quoteMark = null;
                if (_unescaped !== content || _oldRawValue) {
                  (0, _util.ensureObject)(node, "raws");
                  node.raws.value = (_oldRawValue || oldValue) + content;
                }
                lastAdded = "value";
              } else {
                var insensitive = content === "i" || content === "I";
                if ((node.value || node.value === "") && (node.quoteMark || spaceAfterMeaningfulToken)) {
                  node.insensitive = insensitive;
                  if (!insensitive || content === "I") {
                    (0, _util.ensureObject)(node, "raws");
                    node.raws.insensitiveFlag = content;
                  }
                  lastAdded = "insensitive";
                  if (spaceBefore) {
                    (0, _util.ensureObject)(node, "spaces", "insensitive");
                    node.spaces.insensitive.before = spaceBefore;
                    spaceBefore = "";
                  }
                  if (commentBefore) {
                    (0, _util.ensureObject)(node, "raws", "spaces", "insensitive");
                    node.raws.spaces.insensitive.before = commentBefore;
                    commentBefore = "";
                  }
                } else if (node.value || node.value === "") {
                  lastAdded = "value";
                  node.value += content;
                  if (node.raws.value) {
                    node.raws.value += content;
                  }
                }
              }
              spaceAfterMeaningfulToken = false;
              break;
            case tokens.str:
              if (!node.attribute || !node.operator) {
                return this.error("Expected an attribute followed by an operator preceding the string.", {
                  index: token[_tokenize.FIELDS.START_POS]
                });
              }
              var _unescapeValue = (0, _attribute.unescapeValue)(content), unescaped = _unescapeValue.unescaped, quoteMark = _unescapeValue.quoteMark;
              node.value = unescaped;
              node.quoteMark = quoteMark;
              lastAdded = "value";
              (0, _util.ensureObject)(node, "raws");
              node.raws.value = content;
              spaceAfterMeaningfulToken = false;
              break;
            case tokens.equals:
              if (!node.attribute) {
                return this.expected("attribute", token[_tokenize.FIELDS.START_POS], content);
              }
              if (node.value) {
                return this.error('Unexpected "=" found; an operator was already defined.', {
                  index: token[_tokenize.FIELDS.START_POS]
                });
              }
              node.operator = node.operator ? node.operator + content : content;
              lastAdded = "operator";
              spaceAfterMeaningfulToken = false;
              break;
            case tokens.comment:
              if (lastAdded) {
                if (spaceAfterMeaningfulToken || next && next[_tokenize.FIELDS.TYPE] === tokens.space || lastAdded === "insensitive") {
                  var lastComment = (0, _util.getProp)(node, "spaces", lastAdded, "after") || "";
                  var rawLastComment = (0, _util.getProp)(node, "raws", "spaces", lastAdded, "after") || lastComment;
                  (0, _util.ensureObject)(node, "raws", "spaces", lastAdded);
                  node.raws.spaces[lastAdded].after = rawLastComment + content;
                } else {
                  var lastValue = node[lastAdded] || "";
                  var rawLastValue = (0, _util.getProp)(node, "raws", lastAdded) || lastValue;
                  (0, _util.ensureObject)(node, "raws");
                  node.raws[lastAdded] = rawLastValue + content;
                }
              } else {
                commentBefore = commentBefore + content;
              }
              break;
            default:
              return this.error('Unexpected "' + content + '" found.', {
                index: token[_tokenize.FIELDS.START_POS]
              });
          }
          pos++;
        }
        unescapeProp(node, "attribute");
        unescapeProp(node, "namespace");
        this.newNode(new _attribute["default"](node));
        this.position++;
      };
      _proto.parseWhitespaceEquivalentTokens = function parseWhitespaceEquivalentTokens(stopPosition) {
        if (stopPosition < 0) {
          stopPosition = this.tokens.length;
        }
        var startPosition = this.position;
        var nodes = [];
        var space = "";
        var lastComment = void 0;
        do {
          if (WHITESPACE_TOKENS[this.currToken[_tokenize.FIELDS.TYPE]]) {
            if (!this.options.lossy) {
              space += this.content();
            }
          } else if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.comment) {
            var spaces = {};
            if (space) {
              spaces.before = space;
              space = "";
            }
            lastComment = new _comment["default"]({
              value: this.content(),
              source: getTokenSource(this.currToken),
              sourceIndex: this.currToken[_tokenize.FIELDS.START_POS],
              spaces
            });
            nodes.push(lastComment);
          }
        } while (++this.position < stopPosition);
        if (space) {
          if (lastComment) {
            lastComment.spaces.after = space;
          } else if (!this.options.lossy) {
            var firstToken = this.tokens[startPosition];
            var lastToken = this.tokens[this.position - 1];
            nodes.push(new _string["default"]({
              value: "",
              source: getSource(firstToken[_tokenize.FIELDS.START_LINE], firstToken[_tokenize.FIELDS.START_COL], lastToken[_tokenize.FIELDS.END_LINE], lastToken[_tokenize.FIELDS.END_COL]),
              sourceIndex: firstToken[_tokenize.FIELDS.START_POS],
              spaces: {
                before: space,
                after: ""
              }
            }));
          }
        }
        return nodes;
      };
      _proto.convertWhitespaceNodesToSpace = function convertWhitespaceNodesToSpace(nodes, requiredSpace) {
        var _this2 = this;
        if (requiredSpace === void 0) {
          requiredSpace = false;
        }
        var space = "";
        var rawSpace = "";
        nodes.forEach(function(n) {
          var spaceBefore = _this2.lossySpace(n.spaces.before, requiredSpace);
          var rawSpaceBefore = _this2.lossySpace(n.rawSpaceBefore, requiredSpace);
          space += spaceBefore + _this2.lossySpace(n.spaces.after, requiredSpace && spaceBefore.length === 0);
          rawSpace += spaceBefore + n.value + _this2.lossySpace(n.rawSpaceAfter, requiredSpace && rawSpaceBefore.length === 0);
        });
        if (rawSpace === space) {
          rawSpace = void 0;
        }
        var result = {
          space,
          rawSpace
        };
        return result;
      };
      _proto.isNamedCombinator = function isNamedCombinator(position) {
        if (position === void 0) {
          position = this.position;
        }
        return this.tokens[position + 0] && this.tokens[position + 0][_tokenize.FIELDS.TYPE] === tokens.slash && this.tokens[position + 1] && this.tokens[position + 1][_tokenize.FIELDS.TYPE] === tokens.word && this.tokens[position + 2] && this.tokens[position + 2][_tokenize.FIELDS.TYPE] === tokens.slash;
      };
      _proto.namedCombinator = function namedCombinator() {
        if (this.isNamedCombinator()) {
          var nameRaw = this.content(this.tokens[this.position + 1]);
          var name = (0, _util.unesc)(nameRaw).toLowerCase();
          var raws = {};
          if (name !== nameRaw) {
            raws.value = "/" + nameRaw + "/";
          }
          var node = new _combinator["default"]({
            value: "/" + name + "/",
            source: getSource(this.currToken[_tokenize.FIELDS.START_LINE], this.currToken[_tokenize.FIELDS.START_COL], this.tokens[this.position + 2][_tokenize.FIELDS.END_LINE], this.tokens[this.position + 2][_tokenize.FIELDS.END_COL]),
            sourceIndex: this.currToken[_tokenize.FIELDS.START_POS],
            raws
          });
          this.position = this.position + 3;
          return node;
        } else {
          this.unexpected();
        }
      };
      _proto.combinator = function combinator() {
        var _this3 = this;
        if (this.content() === "|") {
          return this.namespace();
        }
        var nextSigTokenPos = this.locateNextMeaningfulToken(this.position);
        if (nextSigTokenPos < 0 || this.tokens[nextSigTokenPos][_tokenize.FIELDS.TYPE] === tokens.comma) {
          var nodes = this.parseWhitespaceEquivalentTokens(nextSigTokenPos);
          if (nodes.length > 0) {
            var last = this.current.last;
            if (last) {
              var _this$convertWhitespa = this.convertWhitespaceNodesToSpace(nodes), space = _this$convertWhitespa.space, rawSpace = _this$convertWhitespa.rawSpace;
              if (rawSpace !== void 0) {
                last.rawSpaceAfter += rawSpace;
              }
              last.spaces.after += space;
            } else {
              nodes.forEach(function(n) {
                return _this3.newNode(n);
              });
            }
          }
          return;
        }
        var firstToken = this.currToken;
        var spaceOrDescendantSelectorNodes = void 0;
        if (nextSigTokenPos > this.position) {
          spaceOrDescendantSelectorNodes = this.parseWhitespaceEquivalentTokens(nextSigTokenPos);
        }
        var node;
        if (this.isNamedCombinator()) {
          node = this.namedCombinator();
        } else if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.combinator) {
          node = new _combinator["default"]({
            value: this.content(),
            source: getTokenSource(this.currToken),
            sourceIndex: this.currToken[_tokenize.FIELDS.START_POS]
          });
          this.position++;
        } else if (WHITESPACE_TOKENS[this.currToken[_tokenize.FIELDS.TYPE]]) {
        } else if (!spaceOrDescendantSelectorNodes) {
          this.unexpected();
        }
        if (node) {
          if (spaceOrDescendantSelectorNodes) {
            var _this$convertWhitespa2 = this.convertWhitespaceNodesToSpace(spaceOrDescendantSelectorNodes), _space = _this$convertWhitespa2.space, _rawSpace = _this$convertWhitespa2.rawSpace;
            node.spaces.before = _space;
            node.rawSpaceBefore = _rawSpace;
          }
        } else {
          var _this$convertWhitespa3 = this.convertWhitespaceNodesToSpace(spaceOrDescendantSelectorNodes, true), _space2 = _this$convertWhitespa3.space, _rawSpace2 = _this$convertWhitespa3.rawSpace;
          if (!_rawSpace2) {
            _rawSpace2 = _space2;
          }
          var spaces = {};
          var raws = {
            spaces: {}
          };
          if (_space2.endsWith(" ") && _rawSpace2.endsWith(" ")) {
            spaces.before = _space2.slice(0, _space2.length - 1);
            raws.spaces.before = _rawSpace2.slice(0, _rawSpace2.length - 1);
          } else if (_space2.startsWith(" ") && _rawSpace2.startsWith(" ")) {
            spaces.after = _space2.slice(1);
            raws.spaces.after = _rawSpace2.slice(1);
          } else {
            raws.value = _rawSpace2;
          }
          node = new _combinator["default"]({
            value: " ",
            source: getTokenSourceSpan(firstToken, this.tokens[this.position - 1]),
            sourceIndex: firstToken[_tokenize.FIELDS.START_POS],
            spaces,
            raws
          });
        }
        if (this.currToken && this.currToken[_tokenize.FIELDS.TYPE] === tokens.space) {
          node.spaces.after = this.optionalSpace(this.content());
          this.position++;
        }
        return this.newNode(node);
      };
      _proto.comma = function comma() {
        if (this.position === this.tokens.length - 1) {
          this.root.trailingComma = true;
          this.position++;
          return;
        }
        this.current._inferEndPosition();
        var selector = new _selector["default"]({
          source: {
            start: tokenStart(this.tokens[this.position + 1])
          }
        });
        this.current.parent.append(selector);
        this.current = selector;
        this.position++;
      };
      _proto.comment = function comment() {
        var current = this.currToken;
        this.newNode(new _comment["default"]({
          value: this.content(),
          source: getTokenSource(current),
          sourceIndex: current[_tokenize.FIELDS.START_POS]
        }));
        this.position++;
      };
      _proto.error = function error(message, opts) {
        throw this.root.error(message, opts);
      };
      _proto.missingBackslash = function missingBackslash() {
        return this.error("Expected a backslash preceding the semicolon.", {
          index: this.currToken[_tokenize.FIELDS.START_POS]
        });
      };
      _proto.missingParenthesis = function missingParenthesis() {
        return this.expected("opening parenthesis", this.currToken[_tokenize.FIELDS.START_POS]);
      };
      _proto.missingSquareBracket = function missingSquareBracket() {
        return this.expected("opening square bracket", this.currToken[_tokenize.FIELDS.START_POS]);
      };
      _proto.unexpected = function unexpected() {
        return this.error("Unexpected '" + this.content() + "'. Escaping special characters with \\ may help.", this.currToken[_tokenize.FIELDS.START_POS]);
      };
      _proto.unexpectedPipe = function unexpectedPipe() {
        return this.error("Unexpected '|'.", this.currToken[_tokenize.FIELDS.START_POS]);
      };
      _proto.namespace = function namespace() {
        var before = this.prevToken && this.content(this.prevToken) || true;
        if (this.nextToken[_tokenize.FIELDS.TYPE] === tokens.word) {
          this.position++;
          return this.word(before);
        } else if (this.nextToken[_tokenize.FIELDS.TYPE] === tokens.asterisk) {
          this.position++;
          return this.universal(before);
        }
        this.unexpectedPipe();
      };
      _proto.nesting = function nesting() {
        if (this.nextToken) {
          var nextContent = this.content(this.nextToken);
          if (nextContent === "|") {
            this.position++;
            return;
          }
        }
        var current = this.currToken;
        this.newNode(new _nesting["default"]({
          value: this.content(),
          source: getTokenSource(current),
          sourceIndex: current[_tokenize.FIELDS.START_POS]
        }));
        this.position++;
      };
      _proto.parentheses = function parentheses() {
        var last = this.current.last;
        var unbalanced = 1;
        this.position++;
        if (last && last.type === types.PSEUDO) {
          var selector = new _selector["default"]({
            source: {
              start: tokenStart(this.tokens[this.position - 1])
            }
          });
          var cache = this.current;
          last.append(selector);
          this.current = selector;
          while (this.position < this.tokens.length && unbalanced) {
            if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.openParenthesis) {
              unbalanced++;
            }
            if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.closeParenthesis) {
              unbalanced--;
            }
            if (unbalanced) {
              this.parse();
            } else {
              this.current.source.end = tokenEnd(this.currToken);
              this.current.parent.source.end = tokenEnd(this.currToken);
              this.position++;
            }
          }
          this.current = cache;
        } else {
          var parenStart = this.currToken;
          var parenValue = "(";
          var parenEnd;
          while (this.position < this.tokens.length && unbalanced) {
            if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.openParenthesis) {
              unbalanced++;
            }
            if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.closeParenthesis) {
              unbalanced--;
            }
            parenEnd = this.currToken;
            parenValue += this.parseParenthesisToken(this.currToken);
            this.position++;
          }
          if (last) {
            last.appendToPropertyAndEscape("value", parenValue, parenValue);
          } else {
            this.newNode(new _string["default"]({
              value: parenValue,
              source: getSource(parenStart[_tokenize.FIELDS.START_LINE], parenStart[_tokenize.FIELDS.START_COL], parenEnd[_tokenize.FIELDS.END_LINE], parenEnd[_tokenize.FIELDS.END_COL]),
              sourceIndex: parenStart[_tokenize.FIELDS.START_POS]
            }));
          }
        }
        if (unbalanced) {
          return this.expected("closing parenthesis", this.currToken[_tokenize.FIELDS.START_POS]);
        }
      };
      _proto.pseudo = function pseudo() {
        var _this4 = this;
        var pseudoStr = "";
        var startingToken = this.currToken;
        while (this.currToken && this.currToken[_tokenize.FIELDS.TYPE] === tokens.colon) {
          pseudoStr += this.content();
          this.position++;
        }
        if (!this.currToken) {
          return this.expected(["pseudo-class", "pseudo-element"], this.position - 1);
        }
        if (this.currToken[_tokenize.FIELDS.TYPE] === tokens.word) {
          this.splitWord(false, function(first, length) {
            pseudoStr += first;
            _this4.newNode(new _pseudo["default"]({
              value: pseudoStr,
              source: getTokenSourceSpan(startingToken, _this4.currToken),
              sourceIndex: startingToken[_tokenize.FIELDS.START_POS]
            }));
            if (length > 1 && _this4.nextToken && _this4.nextToken[_tokenize.FIELDS.TYPE] === tokens.openParenthesis) {
              _this4.error("Misplaced parenthesis.", {
                index: _this4.nextToken[_tokenize.FIELDS.START_POS]
              });
            }
          });
        } else {
          return this.expected(["pseudo-class", "pseudo-element"], this.currToken[_tokenize.FIELDS.START_POS]);
        }
      };
      _proto.space = function space() {
        var content = this.content();
        if (this.position === 0 || this.prevToken[_tokenize.FIELDS.TYPE] === tokens.comma || this.prevToken[_tokenize.FIELDS.TYPE] === tokens.openParenthesis || this.current.nodes.every(function(node) {
          return node.type === "comment";
        })) {
          this.spaces = this.optionalSpace(content);
          this.position++;
        } else if (this.position === this.tokens.length - 1 || this.nextToken[_tokenize.FIELDS.TYPE] === tokens.comma || this.nextToken[_tokenize.FIELDS.TYPE] === tokens.closeParenthesis) {
          this.current.last.spaces.after = this.optionalSpace(content);
          this.position++;
        } else {
          this.combinator();
        }
      };
      _proto.string = function string() {
        var current = this.currToken;
        this.newNode(new _string["default"]({
          value: this.content(),
          source: getTokenSource(current),
          sourceIndex: current[_tokenize.FIELDS.START_POS]
        }));
        this.position++;
      };
      _proto.universal = function universal(namespace) {
        var nextToken = this.nextToken;
        if (nextToken && this.content(nextToken) === "|") {
          this.position++;
          return this.namespace();
        }
        var current = this.currToken;
        this.newNode(new _universal["default"]({
          value: this.content(),
          source: getTokenSource(current),
          sourceIndex: current[_tokenize.FIELDS.START_POS]
        }), namespace);
        this.position++;
      };
      _proto.splitWord = function splitWord(namespace, firstCallback) {
        var _this5 = this;
        var nextToken = this.nextToken;
        var word = this.content();
        while (nextToken && ~[tokens.dollar, tokens.caret, tokens.equals, tokens.word].indexOf(nextToken[_tokenize.FIELDS.TYPE])) {
          this.position++;
          var current = this.content();
          word += current;
          if (current.lastIndexOf("\\") === current.length - 1) {
            var next = this.nextToken;
            if (next && next[_tokenize.FIELDS.TYPE] === tokens.space) {
              word += this.requiredSpace(this.content(next));
              this.position++;
            }
          }
          nextToken = this.nextToken;
        }
        var hasClass = indexesOf(word, ".").filter(function(i) {
          var escapedDot = word[i - 1] === "\\";
          var isKeyframesPercent = /^\d+\.\d+%$/.test(word);
          return !escapedDot && !isKeyframesPercent;
        });
        var hasId = indexesOf(word, "#").filter(function(i) {
          return word[i - 1] !== "\\";
        });
        var interpolations = indexesOf(word, "#{");
        if (interpolations.length) {
          hasId = hasId.filter(function(hashIndex) {
            return !~interpolations.indexOf(hashIndex);
          });
        }
        var indices = (0, _sortAscending["default"])(uniqs([0].concat(hasClass, hasId)));
        indices.forEach(function(ind, i) {
          var index2 = indices[i + 1] || word.length;
          var value = word.slice(ind, index2);
          if (i === 0 && firstCallback) {
            return firstCallback.call(_this5, value, indices.length);
          }
          var node;
          var current2 = _this5.currToken;
          var sourceIndex = current2[_tokenize.FIELDS.START_POS] + indices[i];
          var source = getSource(current2[1], current2[2] + ind, current2[3], current2[2] + (index2 - 1));
          if (~hasClass.indexOf(ind)) {
            var classNameOpts = {
              value: value.slice(1),
              source,
              sourceIndex
            };
            node = new _className["default"](unescapeProp(classNameOpts, "value"));
          } else if (~hasId.indexOf(ind)) {
            var idOpts = {
              value: value.slice(1),
              source,
              sourceIndex
            };
            node = new _id["default"](unescapeProp(idOpts, "value"));
          } else {
            var tagOpts = {
              value,
              source,
              sourceIndex
            };
            unescapeProp(tagOpts, "value");
            node = new _tag["default"](tagOpts);
          }
          _this5.newNode(node, namespace);
          namespace = null;
        });
        this.position++;
      };
      _proto.word = function word(namespace) {
        var nextToken = this.nextToken;
        if (nextToken && this.content(nextToken) === "|") {
          this.position++;
          return this.namespace();
        }
        return this.splitWord(namespace);
      };
      _proto.loop = function loop() {
        while (this.position < this.tokens.length) {
          this.parse(true);
        }
        this.current._inferEndPosition();
        return this.root;
      };
      _proto.parse = function parse2(throwOnParenthesis) {
        switch (this.currToken[_tokenize.FIELDS.TYPE]) {
          case tokens.space:
            this.space();
            break;
          case tokens.comment:
            this.comment();
            break;
          case tokens.openParenthesis:
            this.parentheses();
            break;
          case tokens.closeParenthesis:
            if (throwOnParenthesis) {
              this.missingParenthesis();
            }
            break;
          case tokens.openSquare:
            this.attribute();
            break;
          case tokens.dollar:
          case tokens.caret:
          case tokens.equals:
          case tokens.word:
            this.word();
            break;
          case tokens.colon:
            this.pseudo();
            break;
          case tokens.comma:
            this.comma();
            break;
          case tokens.asterisk:
            this.universal();
            break;
          case tokens.ampersand:
            this.nesting();
            break;
          case tokens.slash:
          case tokens.combinator:
            this.combinator();
            break;
          case tokens.str:
            this.string();
            break;
          case tokens.closeSquare:
            this.missingSquareBracket();
          case tokens.semicolon:
            this.missingBackslash();
          default:
            this.unexpected();
        }
      };
      _proto.expected = function expected(description, index2, found) {
        if (Array.isArray(description)) {
          var last = description.pop();
          description = description.join(", ") + " or " + last;
        }
        var an = /^[aeiou]/.test(description[0]) ? "an" : "a";
        if (!found) {
          return this.error("Expected " + an + " " + description + ".", {
            index: index2
          });
        }
        return this.error("Expected " + an + " " + description + ', found "' + found + '" instead.', {
          index: index2
        });
      };
      _proto.requiredSpace = function requiredSpace(space) {
        return this.options.lossy ? " " : space;
      };
      _proto.optionalSpace = function optionalSpace(space) {
        return this.options.lossy ? "" : space;
      };
      _proto.lossySpace = function lossySpace(space, required) {
        if (this.options.lossy) {
          return required ? " " : "";
        } else {
          return space;
        }
      };
      _proto.parseParenthesisToken = function parseParenthesisToken(token) {
        var content = this.content(token);
        if (token[_tokenize.FIELDS.TYPE] === tokens.space) {
          return this.requiredSpace(content);
        } else {
          return content;
        }
      };
      _proto.newNode = function newNode(node, namespace) {
        if (namespace) {
          if (/^ +$/.test(namespace)) {
            if (!this.options.lossy) {
              this.spaces = (this.spaces || "") + namespace;
            }
            namespace = true;
          }
          node.namespace = namespace;
          unescapeProp(node, "namespace");
        }
        if (this.spaces) {
          node.spaces.before = this.spaces;
          this.spaces = "";
        }
        return this.current.append(node);
      };
      _proto.content = function content(token) {
        if (token === void 0) {
          token = this.currToken;
        }
        return this.css.slice(token[_tokenize.FIELDS.START_POS], token[_tokenize.FIELDS.END_POS]);
      };
      _proto.locateNextMeaningfulToken = function locateNextMeaningfulToken(startPosition) {
        if (startPosition === void 0) {
          startPosition = this.position + 1;
        }
        var searchPosition = startPosition;
        while (searchPosition < this.tokens.length) {
          if (WHITESPACE_EQUIV_TOKENS[this.tokens[searchPosition][_tokenize.FIELDS.TYPE]]) {
            searchPosition++;
            continue;
          } else {
            return searchPosition;
          }
        }
        return -1;
      };
      _createClass(Parser2, [{
        key: "currToken",
        get: function get() {
          return this.tokens[this.position];
        }
      }, {
        key: "nextToken",
        get: function get() {
          return this.tokens[this.position + 1];
        }
      }, {
        key: "prevToken",
        get: function get() {
          return this.tokens[this.position - 1];
        }
      }]);
      return Parser2;
    }();
    exports["default"] = Parser;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/processor.js
var require_processor2 = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/processor.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _parser = _interopRequireDefault(require_parser3());
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    var Processor = /* @__PURE__ */ function() {
      function Processor2(func, options) {
        this.func = func || function noop() {
        };
        this.funcRes = null;
        this.options = options;
      }
      var _proto = Processor2.prototype;
      _proto._shouldUpdateSelector = function _shouldUpdateSelector(rule, options) {
        if (options === void 0) {
          options = {};
        }
        var merged = Object.assign({}, this.options, options);
        if (merged.updateSelector === false) {
          return false;
        } else {
          return typeof rule !== "string";
        }
      };
      _proto._isLossy = function _isLossy(options) {
        if (options === void 0) {
          options = {};
        }
        var merged = Object.assign({}, this.options, options);
        if (merged.lossless === false) {
          return true;
        } else {
          return false;
        }
      };
      _proto._root = function _root(rule, options) {
        if (options === void 0) {
          options = {};
        }
        var parser2 = new _parser["default"](rule, this._parseOptions(options));
        return parser2.root;
      };
      _proto._parseOptions = function _parseOptions(options) {
        return {
          lossy: this._isLossy(options)
        };
      };
      _proto._run = function _run(rule, options) {
        var _this = this;
        if (options === void 0) {
          options = {};
        }
        return new Promise(function(resolve, reject) {
          try {
            var root = _this._root(rule, options);
            Promise.resolve(_this.func(root)).then(function(transform) {
              var string = void 0;
              if (_this._shouldUpdateSelector(rule, options)) {
                string = root.toString();
                rule.selector = string;
              }
              return {
                transform,
                root,
                string
              };
            }).then(resolve, reject);
          } catch (e) {
            reject(e);
            return;
          }
        });
      };
      _proto._runSync = function _runSync(rule, options) {
        if (options === void 0) {
          options = {};
        }
        var root = this._root(rule, options);
        var transform = this.func(root);
        if (transform && typeof transform.then === "function") {
          throw new Error("Selector processor returned a promise to a synchronous call.");
        }
        var string = void 0;
        if (options.updateSelector && typeof rule !== "string") {
          string = root.toString();
          rule.selector = string;
        }
        return {
          transform,
          root,
          string
        };
      };
      _proto.ast = function ast(rule, options) {
        return this._run(rule, options).then(function(result) {
          return result.root;
        });
      };
      _proto.astSync = function astSync(rule, options) {
        return this._runSync(rule, options).root;
      };
      _proto.transform = function transform(rule, options) {
        return this._run(rule, options).then(function(result) {
          return result.transform;
        });
      };
      _proto.transformSync = function transformSync(rule, options) {
        return this._runSync(rule, options).transform;
      };
      _proto.process = function process2(rule, options) {
        return this._run(rule, options).then(function(result) {
          return result.string || result.root.toString();
        });
      };
      _proto.processSync = function processSync(rule, options) {
        var result = this._runSync(rule, options);
        return result.string || result.root.toString();
      };
      return Processor2;
    }();
    exports["default"] = Processor;
    module2.exports = exports.default;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/constructors.js
var require_constructors = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/constructors.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports.universal = exports.tag = exports.string = exports.selector = exports.root = exports.pseudo = exports.nesting = exports.id = exports.comment = exports.combinator = exports.className = exports.attribute = void 0;
    var _attribute = _interopRequireDefault(require_attribute());
    var _className = _interopRequireDefault(require_className());
    var _combinator = _interopRequireDefault(require_combinator());
    var _comment = _interopRequireDefault(require_comment2());
    var _id = _interopRequireDefault(require_id());
    var _nesting = _interopRequireDefault(require_nesting());
    var _pseudo = _interopRequireDefault(require_pseudo());
    var _root = _interopRequireDefault(require_root2());
    var _selector = _interopRequireDefault(require_selector());
    var _string = _interopRequireDefault(require_string());
    var _tag = _interopRequireDefault(require_tag());
    var _universal = _interopRequireDefault(require_universal());
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    var attribute = function attribute2(opts) {
      return new _attribute["default"](opts);
    };
    exports.attribute = attribute;
    var className = function className2(opts) {
      return new _className["default"](opts);
    };
    exports.className = className;
    var combinator = function combinator2(opts) {
      return new _combinator["default"](opts);
    };
    exports.combinator = combinator;
    var comment = function comment2(opts) {
      return new _comment["default"](opts);
    };
    exports.comment = comment;
    var id = function id2(opts) {
      return new _id["default"](opts);
    };
    exports.id = id;
    var nesting = function nesting2(opts) {
      return new _nesting["default"](opts);
    };
    exports.nesting = nesting;
    var pseudo = function pseudo2(opts) {
      return new _pseudo["default"](opts);
    };
    exports.pseudo = pseudo;
    var root = function root2(opts) {
      return new _root["default"](opts);
    };
    exports.root = root;
    var selector = function selector2(opts) {
      return new _selector["default"](opts);
    };
    exports.selector = selector;
    var string = function string2(opts) {
      return new _string["default"](opts);
    };
    exports.string = string;
    var tag = function tag2(opts) {
      return new _tag["default"](opts);
    };
    exports.tag = tag;
    var universal = function universal2(opts) {
      return new _universal["default"](opts);
    };
    exports.universal = universal;
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/guards.js
var require_guards = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/guards.js"(exports) {
    "use strict";
    exports.__esModule = true;
    exports.isComment = exports.isCombinator = exports.isClassName = exports.isAttribute = void 0;
    exports.isContainer = isContainer;
    exports.isIdentifier = void 0;
    exports.isNamespace = isNamespace;
    exports.isNesting = void 0;
    exports.isNode = isNode;
    exports.isPseudo = void 0;
    exports.isPseudoClass = isPseudoClass;
    exports.isPseudoElement = isPseudoElement;
    exports.isUniversal = exports.isTag = exports.isString = exports.isSelector = exports.isRoot = void 0;
    var _types = require_types();
    var _IS_TYPE;
    var IS_TYPE = (_IS_TYPE = {}, _IS_TYPE[_types.ATTRIBUTE] = true, _IS_TYPE[_types.CLASS] = true, _IS_TYPE[_types.COMBINATOR] = true, _IS_TYPE[_types.COMMENT] = true, _IS_TYPE[_types.ID] = true, _IS_TYPE[_types.NESTING] = true, _IS_TYPE[_types.PSEUDO] = true, _IS_TYPE[_types.ROOT] = true, _IS_TYPE[_types.SELECTOR] = true, _IS_TYPE[_types.STRING] = true, _IS_TYPE[_types.TAG] = true, _IS_TYPE[_types.UNIVERSAL] = true, _IS_TYPE);
    function isNode(node) {
      return typeof node === "object" && IS_TYPE[node.type];
    }
    function isNodeType(type, node) {
      return isNode(node) && node.type === type;
    }
    var isAttribute = isNodeType.bind(null, _types.ATTRIBUTE);
    exports.isAttribute = isAttribute;
    var isClassName = isNodeType.bind(null, _types.CLASS);
    exports.isClassName = isClassName;
    var isCombinator = isNodeType.bind(null, _types.COMBINATOR);
    exports.isCombinator = isCombinator;
    var isComment = isNodeType.bind(null, _types.COMMENT);
    exports.isComment = isComment;
    var isIdentifier = isNodeType.bind(null, _types.ID);
    exports.isIdentifier = isIdentifier;
    var isNesting = isNodeType.bind(null, _types.NESTING);
    exports.isNesting = isNesting;
    var isPseudo = isNodeType.bind(null, _types.PSEUDO);
    exports.isPseudo = isPseudo;
    var isRoot = isNodeType.bind(null, _types.ROOT);
    exports.isRoot = isRoot;
    var isSelector = isNodeType.bind(null, _types.SELECTOR);
    exports.isSelector = isSelector;
    var isString = isNodeType.bind(null, _types.STRING);
    exports.isString = isString;
    var isTag = isNodeType.bind(null, _types.TAG);
    exports.isTag = isTag;
    var isUniversal = isNodeType.bind(null, _types.UNIVERSAL);
    exports.isUniversal = isUniversal;
    function isPseudoElement(node) {
      return isPseudo(node) && node.value && (node.value.startsWith("::") || node.value.toLowerCase() === ":before" || node.value.toLowerCase() === ":after" || node.value.toLowerCase() === ":first-letter" || node.value.toLowerCase() === ":first-line");
    }
    function isPseudoClass(node) {
      return isPseudo(node) && !isPseudoElement(node);
    }
    function isContainer(node) {
      return !!(isNode(node) && node.walk);
    }
    function isNamespace(node) {
      return isAttribute(node) || isTag(node);
    }
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/index.js
var require_selectors = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/selectors/index.js"(exports) {
    "use strict";
    exports.__esModule = true;
    var _types = require_types();
    Object.keys(_types).forEach(function(key) {
      if (key === "default" || key === "__esModule")
        return;
      if (key in exports && exports[key] === _types[key])
        return;
      exports[key] = _types[key];
    });
    var _constructors = require_constructors();
    Object.keys(_constructors).forEach(function(key) {
      if (key === "default" || key === "__esModule")
        return;
      if (key in exports && exports[key] === _constructors[key])
        return;
      exports[key] = _constructors[key];
    });
    var _guards = require_guards();
    Object.keys(_guards).forEach(function(key) {
      if (key === "default" || key === "__esModule")
        return;
      if (key in exports && exports[key] === _guards[key])
        return;
      exports[key] = _guards[key];
    });
  }
});

// node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/index.js
var require_dist = __commonJS({
  "node_modules/.pnpm/postcss-selector-parser@6.0.13/node_modules/postcss-selector-parser/dist/index.js"(exports, module2) {
    "use strict";
    exports.__esModule = true;
    exports["default"] = void 0;
    var _processor = _interopRequireDefault(require_processor2());
    var selectors = _interopRequireWildcard(require_selectors());
    function _getRequireWildcardCache(nodeInterop) {
      if (typeof WeakMap !== "function")
        return null;
      var cacheBabelInterop = /* @__PURE__ */ new WeakMap();
      var cacheNodeInterop = /* @__PURE__ */ new WeakMap();
      return (_getRequireWildcardCache = function _getRequireWildcardCache2(nodeInterop2) {
        return nodeInterop2 ? cacheNodeInterop : cacheBabelInterop;
      })(nodeInterop);
    }
    function _interopRequireWildcard(obj, nodeInterop) {
      if (!nodeInterop && obj && obj.__esModule) {
        return obj;
      }
      if (obj === null || typeof obj !== "object" && typeof obj !== "function") {
        return { "default": obj };
      }
      var cache = _getRequireWildcardCache(nodeInterop);
      if (cache && cache.has(obj)) {
        return cache.get(obj);
      }
      var newObj = {};
      var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor;
      for (var key in obj) {
        if (key !== "default" && Object.prototype.hasOwnProperty.call(obj, key)) {
          var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null;
          if (desc && (desc.get || desc.set)) {
            Object.defineProperty(newObj, key, desc);
          } else {
            newObj[key] = obj[key];
          }
        }
      }
      newObj["default"] = obj;
      if (cache) {
        cache.set(obj, newObj);
      }
      return newObj;
    }
    function _interopRequireDefault(obj) {
      return obj && obj.__esModule ? obj : { "default": obj };
    }
    var parser2 = function parser3(processor) {
      return new _processor["default"](processor);
    };
    Object.assign(parser2, selectors);
    delete parser2.__esModule;
    var _default = parser2;
    exports["default"] = _default;
    module2.exports = exports.default;
  }
});

// src/index.ts
var src_exports = {};
__export(src_exports, {
  getThemeProperties: () => getThemeProperties,
  skeleton: () => skeleton2
});
module.exports = __toCommonJS(src_exports);
var import_plugin2 = __toESM(require("tailwindcss/plugin.js"));

// node_modules/.pnpm/postcss-js@4.0.1_postcss@8.4.24/node_modules/postcss-js/index.mjs
var import_index = __toESM(require_postcss_js(), 1);
var postcss_js_default = import_index.default;
var objectify = import_index.default.objectify;
var parse = import_index.default.parse;
var async = import_index.default.async;
var sync = import_index.default.sync;

// src/tailwind/core.ts
var import_plugin = __toESM(require("tailwindcss/plugin.js"));

// src/tailwind/settings.ts
var settings = {
  colorNames: ["primary", "secondary", "tertiary", "success", "warning", "error", "surface"],
  colorShades: [50, 100, 200, 300, 400, 500, 600, 700, 800, 900],
  colorPairings: [
    // forward:
    { light: 50, dark: 900 },
    { light: 100, dark: 800 },
    { light: 200, dark: 700 },
    { light: 300, dark: 600 },
    { light: 400, dark: 500 },
    // backwards
    { light: 900, dark: 50 },
    { light: 800, dark: 100 },
    { light: 700, dark: 200 },
    { light: 600, dark: 300 },
    { light: 500, dark: 400 }
  ]
};
var settings_default = settings;

// src/tailwind/colors.ts
function generatePaletteShades(colorName) {
  const shadeObj = {};
  settings_default.colorShades.forEach((s) => shadeObj[s] = `rgb(var(--color-${colorName}-${s}) / <alpha-value>)`);
  return shadeObj;
}
var colors = () => {
  const paletteObj = {};
  settings_default.colorNames.forEach((n) => paletteObj[n] = generatePaletteShades(n));
  return paletteObj;
};
var colors_default = colors;

// src/tailwind/tokens/backgrounds.ts
var backdropAlpha = 0.7;
var hoverAlpha = 0.1;
var backgrounds = () => {
  const classes = {};
  settings.colorNames.forEach((n) => {
    classes[`.bg-${n}-backdrop-token`] = { "background-color": `rgb(var(--color-${n}-400) / ${backdropAlpha})` };
    classes[`.dark .bg-${n}-backdrop-token`] = { "background-color": `rgb(var(--color-${n}-900) / ${backdropAlpha})` };
    classes[`.bg-${n}-hover-token:hover`] = { "background-color": `rgb(var(--color-${n}-500) / ${hoverAlpha})` };
    classes[`.dark .bg-${n}-hover-token:hover`] = { "background-color": `rgb(var(--color-${n}-500) / ${hoverAlpha})` };
    classes[`.bg-${n}-active-token`] = {
      "background-color": `rgb(var(--color-${n}-500)) !important`,
      color: `rgb(var(--on-${n}))`,
      fill: `rgb(var(--on-${n}))`
    };
    settings.colorPairings.forEach((p) => {
      classes[`.bg-${n}-${p.light}-${p.dark}-token`] = { "background-color": `rgb(var(--color-${n}-${p.light}))` };
      classes[`.dark .bg-${n}-${p.light}-${p.dark}-token`] = { "background-color": `rgb(var(--color-${n}-${p.dark}))` };
    });
  });
  return classes;
};
var backgrounds_default = backgrounds;

// src/tailwind/tokens/borders.ts
var borders = () => {
  const classes = {
    // Border Width - ex: .border-token
    ".border-token": { "border-width": "var(--theme-border-base)" }
  };
  settings.colorNames.forEach((n) => {
    settings.colorPairings.forEach((p) => {
      classes[`.border-${n}-${p.light}-${p.dark}-token`] = { "border-color": `rgb(var(--color-${n}-${p.light}))` };
      classes[`.dark .border-${n}-${p.light}-${p.dark}-token`] = { "border-color": `rgb(var(--color-${n}-${p.dark}))` };
      classes[`.divide-${n}-${p.light}-${p.dark}-token`] = { "border-color": `rgb(var(--color-${n}-${p.light}))` };
      classes[`.dark .divide-${n}-${p.light}-${p.dark}-token`] = { "border-color": `rgb(var(--color-${n}-${p.dark}))` };
    });
  });
  return classes;
};
var borders_default = borders;

// src/tailwind/tokens/border-radius.ts
var borderRadius = () => {
  return {
    // Base
    ".rounded-token": { "border-radius": "var(--theme-rounded-base)" },
    ".rounded-tl-token": { "border-top-left-radius": "var(--theme-rounded-base)" },
    ".rounded-tr-token": { "border-top-right-radius": "var(--theme-rounded-base)" },
    ".rounded-bl-token": { "border-bottom-left-radius": "var(--theme-rounded-base)" },
    ".rounded-br-token": { "border-bottom-right-radius": "var(--theme-rounded-base)" },
    // Container
    ".rounded-container-token": { "border-radius": "var(--theme-rounded-container)" },
    ".rounded-tl-container-token": { "border-top-left-radius": "var(--theme-rounded-container)" },
    ".rounded-tr-container-token": { "border-top-right-radius": "var(--theme-rounded-container)" },
    ".rounded-bl-container-token": { "border-bottom-left-radius": "var(--theme-rounded-container)" },
    ".rounded-br-container-token": { "border-bottom-right-radius": "var(--theme-rounded-container)" }
  };
};
var border_radius_default = borderRadius;

// src/tailwind/tokens/fills.ts
var fills = () => {
  const classes = {
    ".fill-base-token": { fill: "rgba(var(--theme-font-color-base))" },
    ".fill-dark-token": { fill: "rgba(var(--theme-font-color-dark))" },
    // Fill Token - ex: .fill-token
    ".fill-token": { fill: "rgba(var(--theme-font-color-base))" },
    ".dark .fill-token": { fill: "rgba(var(--theme-font-color-dark))" }
  };
  settings.colorNames.forEach((n) => {
    classes[`.fill-on-${n}-token`] = { fill: `rgb(var(--on-${n}))` };
  });
  return classes;
};
var fills_default = fills;

// src/tailwind/tokens/text.ts
var text = () => {
  const classes = {
    // Font Family
    ".font-heading-token": { "font-family": "var(--theme-font-family-heading)" },
    ".font-token": { "font-family": "var(--theme-font-family-base)" },
    // Default Text Colors
    ".text-base-token": { color: "rgba(var(--theme-font-color-base))" },
    ".text-dark-token": { color: "rgba(var(--theme-font-color-dark))" },
    // Light/Dark Text Color  - ex: .text-token
    ".text-token": { color: "rgba(var(--theme-font-color-base))" },
    ".dark .text-token": { color: "rgba(var(--theme-font-color-dark))" }
  };
  settings.colorNames.forEach((n) => {
    classes[`.text-on-${n}-token`] = { color: `rgb(var(--on-${n}))` };
    settings.colorPairings.forEach((p) => {
      classes[`.text-${n}-${p.light}-${p.dark}-token`] = { color: `rgb(var(--color-${n}-${p.light}))` };
      classes[`.dark .text-${n}-${p.light}-${p.dark}-token`] = { color: `rgb(var(--color-${n}-${p.dark}))` };
      classes[`.decoration-${n}-${p.light}-${p.dark}-token`] = { "text-decoration-color": `rgb(var(--color-${n}-${p.light}))` };
      classes[`.dark .decoration-${n}-${p.light}-${p.dark}-token`] = { "text-decoration-color": `rgb(var(--color-${n}-${p.dark}))` };
    });
  });
  return classes;
};
var text_default = text;

// src/tailwind/tokens/rings.ts
var ringTokenTheme = {
  "--tw-ring-offset-shadow": `var(--tw-ring-inset) 0 0 0 var(--theme-border-base) var(--tw-ring-offset-color)`,
  "--tw-ring-shadow": `var(--tw-ring-inset) 0 0 0 calc(2px + var(--theme-border-base)) var(--tw-ring-color)`,
  "box-shadow": `var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)`
};
var ringOutlineShared = {
  // .ring-[1px]
  "--tw-ring-offset-shadow": "var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color)",
  "--tw-ring-shadow": "var(--tw-ring-inset) 0 0 0 calc(1px + var(--tw-ring-offset-width)) var(--tw-ring-color)",
  "box-shadow": "var(--tw-ring-offset-shadow), var(--tw-ring-shadow), var(--tw-shadow, 0 0 #0000)",
  // .ring-inset
  "--tw-ring-inset": "inset"
};
var rings = () => {
  const classes = {
    ".ring-token": {
      ...ringTokenTheme
    },
    // Ring Outline (for cards)
    // Example: .ring-outline-token
    ".ring-outline-token": {
      ...ringOutlineShared,
      "--tw-ring-color": "rgb(23 23 23 / 0.05)"
      // neutral-900, 5% opacity
    },
    ".dark .ring-outline-token": {
      ...ringOutlineShared,
      "--tw-ring-color": "rgb(250 250 250 / 0.05)"
      // neutral-50, 5% opacity
    }
  };
  settings.colorNames.forEach((n) => {
    settings.colorPairings.forEach((p) => {
      classes[`.ring-${n}-${p.light}-${p.dark}-token`] = {
        "--tw-ring-color": `rgb(var(--color-${n}-${p.light}) / 1)`
      };
      classes[`.dark .ring-${n}-${p.light}-${p.dark}-token`] = {
        "--tw-ring-color": `rgb(var(--color-${n}-${p.dark}) / 1)`
      };
    });
  });
  return classes;
};
var rings_default = rings;

// src/tailwind/core.ts
var coreUtilities = {
  // Implement Skeleton design token classes
  ...backgrounds_default(),
  ...borders_default(),
  ...border_radius_default(),
  ...fills_default(),
  ...text_default(),
  ...rings_default()
};
var coreConfig = {
  theme: {
    extend: {
      // Implement Skeleton theme colors
      colors: colors_default()
    }
  }
};
function getSkeletonClasses() {
  try {
    const { components, base } = require_generated_classes();
    if (typeof components !== "object" || typeof base !== "object") {
      console.error("Failed to load Skeleton classes");
      process.exitCode = 1;
    }
    return { components, base };
  } catch {
    console.error("generated-classes.js hasn't generated yet");
  }
  return { components: void 0, base: void 0 };
}
var corePlugin = (0, import_plugin.default)(({ addUtilities }) => {
  addUtilities(coreUtilities);
}, coreConfig);

// src/tailwind/themes/crimson.ts
var crimson = {
  name: "crimson",
  properties: {
    "--theme-font-family-base": "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,\n		'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'",
    "--theme-font-family-heading": "system-ui",
    "--theme-font-color-base": "var(--color-surface-900)",
    "--theme-font-color-dark": "var(--color-surface-50)",
    "--theme-rounded-base": "24px",
    "--theme-rounded-container": "24px",
    "--theme-border-base": "1px",
    "--on-primary": "255 255 255",
    "--on-secondary": "255 255 255",
    "--on-tertiary": "0 0 0",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "0 0 0",
    "--on-surface": "255 255 255",
    "--color-primary-50": "249 220 226",
    "--color-primary-100": "246 208 216",
    "--color-primary-200": "244 197 206",
    "--color-primary-300": "238 162 177",
    "--color-primary-400": "225 92 119",
    "--color-primary-500": "212 22 60",
    "--color-primary-600": "191 20 54",
    "--color-primary-700": "159 17 45",
    "--color-primary-800": "127 13 36",
    "--color-primary-900": "104 11 29",
    "--color-secondary-50": "227 237 243",
    "--color-secondary-100": "218 231 239",
    "--color-secondary-200": "209 225 235",
    "--color-secondary-300": "181 206 223",
    "--color-secondary-400": "126 170 199",
    "--color-secondary-500": "70 133 175",
    "--color-secondary-600": "63 120 158",
    "--color-secondary-700": "53 100 131",
    "--color-secondary-800": "42 80 105",
    "--color-secondary-900": "34 65 86",
    "--color-tertiary-50": "246 244 244",
    "--color-tertiary-100": "242 240 240",
    "--color-tertiary-200": "239 237 236",
    "--color-tertiary-300": "230 226 225",
    "--color-tertiary-400": "211 204 203",
    "--color-tertiary-500": "192 182 180",
    "--color-tertiary-600": "173 164 162",
    "--color-tertiary-700": "144 137 135",
    "--color-tertiary-800": "115 109 108",
    "--color-tertiary-900": "94 89 88",
    "--color-success-50": "246 250 239",
    "--color-success-100": "243 248 234",
    "--color-success-200": "240 247 229",
    "--color-success-300": "230 241 213",
    "--color-success-400": "212 231 182",
    "--color-success-500": "193 221 151",
    "--color-success-600": "174 199 136",
    "--color-success-700": "145 166 113",
    "--color-success-800": "116 133 91",
    "--color-success-900": "95 108 74",
    "--color-warning-50": "251 246 231",
    "--color-warning-100": "250 243 223",
    "--color-warning-200": "248 240 215",
    "--color-warning-300": "244 231 191",
    "--color-warning-400": "236 212 142",
    "--color-warning-500": "228 194 94",
    "--color-warning-600": "205 175 85",
    "--color-warning-700": "171 146 71",
    "--color-warning-800": "137 116 56",
    "--color-warning-900": "112 95 46",
    "--color-error-50": "248 236 236",
    "--color-error-100": "246 229 230",
    "--color-error-200": "244 223 224",
    "--color-error-300": "237 204 205",
    "--color-error-400": "224 165 167",
    "--color-error-500": "210 127 129",
    "--color-error-600": "189 114 116",
    "--color-error-700": "158 95 97",
    "--color-error-800": "126 76 77",
    "--color-error-900": "103 62 63",
    "--color-surface-50": "223 224 226",
    "--color-surface-100": "213 213 217",
    "--color-surface-200": "202 203 207",
    "--color-surface-300": "170 171 179",
    "--color-surface-400": "107 109 121",
    "--color-surface-500": "43 46 64",
    "--color-surface-600": "39 41 58",
    "--color-surface-700": "32 35 48",
    "--color-surface-800": "26 28 38",
    "--color-surface-900": "21 23 31"
  },
  properties_dark: {},
  enhancements: {}
};
var crimson_default = crimson;

// src/tailwind/themes/gold-nouveau.ts
var goldNouveau = {
  name: "gold-nouveau",
  properties: {
    "--theme-font-family-base": "system-ui, sans-serif",
    "--theme-font-family-heading": "'Quicksand', sans-serif",
    "--theme-font-color-base": "var(--color-surface-900)",
    "--theme-font-color-dark": "var(--color-surface-50)",
    "--theme-rounded-base": "4px",
    "--theme-rounded-container": "4px",
    "--theme-border-base": "1px",
    "--on-primary": "255 255 255",
    "--on-secondary": "255 255 255",
    "--on-tertiary": "255 255 255",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "250 248 252",
    "--color-primary-100": "242 238 247",
    "--color-primary-200": "229 220 239",
    "--color-primary-300": "209 192 226",
    "--color-primary-400": "162 129 197",
    "--color-primary-500": "116 74 161",
    "--color-primary-600": "83 53 115",
    "--color-primary-700": "60 39 84",
    "--color-primary-800": "35 22 49",
    "--color-primary-900": "18 11 24",
    "--color-secondary-50": "218 234 251",
    "--color-secondary-100": "205 227 250",
    "--color-secondary-200": "193 220 249",
    "--color-secondary-300": "155 199 245",
    "--color-secondary-400": "81 156 237",
    "--color-secondary-500": "6 114 229",
    "--color-secondary-600": "5 103 206",
    "--color-secondary-700": "5 86 172",
    "--color-secondary-800": "4 68 137",
    "--color-secondary-900": "3 56 112",
    "--color-tertiary-50": "236 235 250",
    "--color-tertiary-100": "229 228 248",
    "--color-tertiary-200": "223 221 247",
    "--color-tertiary-300": "204 201 241",
    "--color-tertiary-400": "165 161 231",
    "--color-tertiary-500": "127 120 221",
    "--color-tertiary-600": "114 108 199",
    "--color-tertiary-700": "95 90 166",
    "--color-tertiary-800": "76 72 133",
    "--color-tertiary-900": "62 59 108",
    "--color-success-50": "234 246 237",
    "--color-success-100": "227 243 231",
    "--color-success-200": "220 241 225",
    "--color-success-300": "199 232 206",
    "--color-success-400": "156 214 170",
    "--color-success-500": "114 197 133",
    "--color-success-600": "103 177 120",
    "--color-success-700": "86 148 100",
    "--color-success-800": "68 118 80",
    "--color-success-900": "56 97 65",
    "--color-warning-50": "251 236 218",
    "--color-warning-100": "250 229 206",
    "--color-warning-200": "249 223 193",
    "--color-warning-300": "245 204 156",
    "--color-warning-400": "238 165 82",
    "--color-warning-500": "231 127 8",
    "--color-warning-600": "208 114 7",
    "--color-warning-700": "173 95 6",
    "--color-warning-800": "139 76 5",
    "--color-warning-900": "113 62 4",
    "--color-error-50": "238 219 222",
    "--color-error-100": "233 207 211",
    "--color-error-200": "227 195 200",
    "--color-error-300": "210 159 167",
    "--color-error-400": "177 87 100",
    "--color-error-500": "143 15 34",
    "--color-error-600": "129 14 31",
    "--color-error-700": "107 11 26",
    "--color-error-800": "86 9 20",
    "--color-error-900": "70 7 17",
    "--color-surface-50": "250 248 252",
    "--color-surface-100": "242 238 247",
    "--color-surface-200": "229 220 239",
    "--color-surface-300": "209 192 226",
    "--color-surface-400": "162 129 197",
    "--color-surface-500": "116 74 161",
    "--color-surface-600": "83 53 115",
    "--color-surface-700": "60 39 84",
    "--color-surface-800": "35 22 49",
    "--color-surface-900": "18 11 24"
  },
  properties_dark: {
    "--on-primary": "0 0 0",
    "--color-primary-50": "251 247 224",
    "--color-primary-100": "250 244 214",
    "--color-primary-200": "249 241 204",
    "--color-primary-300": "245 233 173",
    "--color-primary-400": "238 217 112",
    "--color-primary-500": "230 200 51",
    "--color-primary-600": "207 180 46",
    "--color-primary-700": "173 150 38",
    "--color-primary-800": "138 120 31",
    "--color-primary-900": "113 98 25"
  },
  enhancements: {
    "[data-theme='gold-nouveau'] h1,\n[data-theme='gold-nouveau'] h2,\n[data-theme='gold-nouveau'] h3,\n[data-theme='gold-nouveau'] h4,\n[data-theme='gold-nouveau'] h5,\n[data-theme='gold-nouveau'] h6": { fontWeight: "bold" },
    "[data-theme='gold-nouveau']": {
      backgroundImage: "radial-gradient(at 0% 100%, rgba(var(--color-secondary-500) / 0.33) 0px, transparent 50%),\n		radial-gradient(at 98% 100%, rgba(var(--color-error-500) / 0.33) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var gold_nouveau_default = goldNouveau;

// src/tailwind/themes/hamlindigo.ts
var hamlindigo = {
  name: "hamlindigo",
  properties: {
    "--theme-font-family-base": "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,\n		'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'",
    "--theme-font-family-heading": "serif",
    "--theme-font-color-base": "0 0 0",
    "--theme-font-color-dark": "255 255 255",
    "--theme-rounded-base": "2px",
    "--theme-rounded-container": "2px",
    "--theme-border-base": "2px",
    "--on-primary": "0 0 0",
    "--on-secondary": "255 255 255",
    "--on-tertiary": "255 255 255",
    "--on-success": "255 255 255",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "242 245 253",
    "--color-primary-100": "238 242 252",
    "--color-primary-200": "233 239 252",
    "--color-primary-300": "220 229 249",
    "--color-primary-400": "194 210 245",
    "--color-primary-500": "168 190 241",
    "--color-primary-600": "151 171 217",
    "--color-primary-700": "126 143 181",
    "--color-primary-800": "101 114 145",
    "--color-primary-900": "82 93 118",
    "--color-secondary-50": "241 238 230",
    "--color-secondary-100": "237 232 222",
    "--color-secondary-200": "232 227 214",
    "--color-secondary-300": "219 210 189",
    "--color-secondary-400": "191 176 140",
    "--color-secondary-500": "164 142 91",
    "--color-secondary-600": "148 128 82",
    "--color-secondary-700": "123 107 68",
    "--color-secondary-800": "98 85 55",
    "--color-secondary-900": "80 70 45",
    "--color-tertiary-50": "231 239 241",
    "--color-tertiary-100": "223 234 237",
    "--color-tertiary-200": "216 229 232",
    "--color-tertiary-300": "192 213 218",
    "--color-tertiary-400": "144 182 191",
    "--color-tertiary-500": "97 151 163",
    "--color-tertiary-600": "87 136 147",
    "--color-tertiary-700": "73 113 122",
    "--color-tertiary-800": "58 91 98",
    "--color-tertiary-900": "48 74 80",
    "--color-success-50": "227 239 236",
    "--color-success-100": "218 234 229",
    "--color-success-200": "209 228 223",
    "--color-success-300": "181 212 203",
    "--color-success-400": "126 180 164",
    "--color-success-500": "71 148 125",
    "--color-success-600": "64 133 113",
    "--color-success-700": "53 111 94",
    "--color-success-800": "43 89 75",
    "--color-success-900": "35 73 61",
    "--color-warning-50": "249 242 226",
    "--color-warning-100": "248 238 216",
    "--color-warning-200": "246 234 207",
    "--color-warning-300": "240 221 178",
    "--color-warning-400": "229 195 120",
    "--color-warning-500": "218 169 62",
    "--color-warning-600": "196 152 56",
    "--color-warning-700": "164 127 47",
    "--color-warning-800": "131 101 37",
    "--color-warning-900": "107 83 30",
    "--color-error-50": "241 231 234",
    "--color-error-100": "236 223 227",
    "--color-error-200": "232 216 221",
    "--color-error-300": "218 192 200",
    "--color-error-400": "190 144 158",
    "--color-error-500": "162 97 117",
    "--color-error-600": "146 87 105",
    "--color-error-700": "122 73 88",
    "--color-error-800": "97 58 70",
    "--color-error-900": "79 48 57",
    "--color-surface-50": "232 234 241",
    "--color-surface-100": "224 228 237",
    "--color-surface-200": "216 221 232",
    "--color-surface-300": "193 200 218",
    "--color-surface-400": "146 159 191",
    "--color-surface-500": "99 118 163",
    "--color-surface-600": "89 106 147",
    "--color-surface-700": "74 89 122",
    "--color-surface-800": "59 71 98",
    "--color-surface-900": "49 58 80"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='hamlindigo']": {
      backgroundImage: `url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='8' height='8' viewBox='0 0 8 8'%3E%3Cg fill='%23e0e4ed' fill-opacity='0.5'%3E%3Cpath fill-rule='evenodd' d='M0 0h4v4H0V0zm4 4h4v4H4V4z'/%3E%3C/g%3E%3C/svg%3E")`
    },
    ".dark [data-theme='hamlindigo']": {
      backgroundImage: `url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='8' height='8' viewBox='0 0 8 8'%3E%3Cg fill='%233b4762' fill-opacity='0.2'%3E%3Cpath fill-rule='evenodd' d='M0 0h4v4H0V0zm4 4h4v4H4V4z'/%3E%3C/g%3E%3C/svg%3E")`
    }
  }
};
var hamlindigo_default = hamlindigo;

// src/tailwind/themes/modern.ts
var modern = {
  name: "modern",
  properties: {
    "--theme-font-family-base": "'Quicksand', sans-serif",
    "--theme-font-family-heading": "'Quicksand', sans-serif",
    "--theme-font-color-base": "var(--color-surface-900)",
    "--theme-font-color-dark": "var(--color-tertiary-50)",
    "--theme-rounded-base": "9999px",
    "--theme-rounded-container": "24px",
    "--theme-border-base": "3px",
    "--on-primary": "255 255 255",
    "--on-secondary": "0 0 0",
    "--on-tertiary": "0 0 0",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "252 228 240",
    "--color-primary-100": "251 218 235",
    "--color-primary-200": "250 209 230",
    "--color-primary-300": "247 182 214",
    "--color-primary-400": "242 127 184",
    "--color-primary-500": "236 72 153",
    "--color-primary-600": "212 65 138",
    "--color-primary-700": "177 54 115",
    "--color-primary-800": "142 43 92",
    "--color-primary-900": "116 35 75",
    "--color-secondary-50": "218 244 249",
    "--color-secondary-100": "205 240 246",
    "--color-secondary-200": "193 237 244",
    "--color-secondary-300": "155 226 238",
    "--color-secondary-400": "81 204 225",
    "--color-secondary-500": "6 182 212",
    "--color-secondary-600": "5 164 191",
    "--color-secondary-700": "5 137 159",
    "--color-secondary-800": "4 109 127",
    "--color-secondary-900": "3 89 104",
    "--color-tertiary-50": "220 244 242",
    "--color-tertiary-100": "208 241 237",
    "--color-tertiary-200": "196 237 233",
    "--color-tertiary-300": "161 227 219",
    "--color-tertiary-400": "91 205 193",
    "--color-tertiary-500": "20 184 166",
    "--color-tertiary-600": "18 166 149",
    "--color-tertiary-700": "15 138 125",
    "--color-tertiary-800": "12 110 100",
    "--color-tertiary-900": "10 90 81",
    "--color-success-50": "237 247 220",
    "--color-success-100": "230 245 208",
    "--color-success-200": "224 242 197",
    "--color-success-300": "206 235 162",
    "--color-success-400": "169 219 92",
    "--color-success-500": "132 204 22",
    "--color-success-600": "119 184 20",
    "--color-success-700": "99 153 17",
    "--color-success-800": "79 122 13",
    "--color-success-900": "65 100 11",
    "--color-warning-50": "252 244 218",
    "--color-warning-100": "251 240 206",
    "--color-warning-200": "250 236 193",
    "--color-warning-300": "247 225 156",
    "--color-warning-400": "240 202 82",
    "--color-warning-500": "234 179 8",
    "--color-warning-600": "211 161 7",
    "--color-warning-700": "176 134 6",
    "--color-warning-800": "140 107 5",
    "--color-warning-900": "115 88 4",
    "--color-error-50": "253 227 227",
    "--color-error-100": "252 218 218",
    "--color-error-200": "251 208 208",
    "--color-error-300": "249 180 180",
    "--color-error-400": "244 124 124",
    "--color-error-500": "239 68 68",
    "--color-error-600": "215 61 61",
    "--color-error-700": "179 51 51",
    "--color-error-800": "143 41 41",
    "--color-error-900": "117 33 33",
    "--color-surface-50": "232 232 253",
    "--color-surface-100": "224 224 252",
    "--color-surface-200": "216 217 252",
    "--color-surface-300": "193 194 249",
    "--color-surface-400": "146 148 245",
    "--color-surface-500": "99 102 241",
    "--color-surface-600": "89 92 217",
    "--color-surface-700": "74 77 181",
    "--color-surface-800": "59 61 145",
    "--color-surface-900": "49 50 118"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='modern'] h1,\n[data-theme='modern'] h2,\n[data-theme='modern'] h3,\n[data-theme='modern'] h4,\n[data-theme='modern'] h5,\n[data-theme='modern'] h6,\n[data-theme='modern'] a,\n[data-theme='modern'] button": { fontWeight: "bold" },
    "[data-theme='modern']": {
      backgroundImage: "radial-gradient(at 76% 0%, hsla(189,100%,56%,0.36) 0px, transparent 50%),\n		radial-gradient(at 1% 0%, hsla(340,100%,76%,0.26) 0px, transparent 50%),\n		radial-gradient(at 20% 100%, hsla(241,100%,70%,0.47) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    },
    ".dark [data-theme='modern']": {
      backgroundImage: "radial-gradient(at 76% 0%, hsla(189,100%,56%,0.20) 0px, transparent 50%),\n		radial-gradient(at 1% 0%, hsla(340,100%,76%,0.15) 0px, transparent 50%),\n		radial-gradient(at 20% 100%, hsla(241,100%,70%,0.30) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var modern_default = modern;

// src/tailwind/themes/rocket.ts
var rocket = {
  name: "rocket",
  properties: {
    "--theme-font-family-base": "system-ui",
    "--theme-font-family-heading": "'Space Grotesk', sans-serif",
    "--theme-font-color-base": "var(--color-primary-900)",
    "--theme-font-color-dark": "var(--color-primary-100)",
    "--theme-rounded-base": "0px",
    "--theme-rounded-container": "0px",
    "--theme-border-base": "0px",
    "--on-primary": "0 0 0",
    "--on-secondary": "255 255 255",
    "--on-tertiary": "255 255 255",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "218 244 249",
    "--color-primary-100": "205 240 246",
    "--color-primary-200": "193 237 244",
    "--color-primary-300": "155 226 238",
    "--color-primary-400": "81 204 225",
    "--color-primary-500": "6 182 212",
    "--color-primary-600": "5 164 191",
    "--color-primary-700": "5 137 159",
    "--color-primary-800": "4 109 127",
    "--color-primary-900": "3 89 104",
    "--color-secondary-50": "226 236 254",
    "--color-secondary-100": "216 230 253",
    "--color-secondary-200": "206 224 253",
    "--color-secondary-300": "177 205 251",
    "--color-secondary-400": "118 168 249",
    "--color-secondary-500": "59 130 246",
    "--color-secondary-600": "53 117 221",
    "--color-secondary-700": "44 98 185",
    "--color-secondary-800": "35 78 148",
    "--color-secondary-900": "29 64 121",
    "--color-tertiary-50": "242 230 254",
    "--color-tertiary-100": "238 221 253",
    "--color-tertiary-200": "233 213 253",
    "--color-tertiary-300": "220 187 252",
    "--color-tertiary-400": "194 136 249",
    "--color-tertiary-500": "168 85 247",
    "--color-tertiary-600": "151 77 222",
    "--color-tertiary-700": "126 64 185",
    "--color-tertiary-800": "101 51 148",
    "--color-tertiary-900": "82 42 121",
    "--color-success-50": "228 247 220",
    "--color-success-100": "219 245 208",
    "--color-success-200": "210 242 197",
    "--color-success-300": "183 234 161",
    "--color-success-400": "130 219 91",
    "--color-success-500": "76 203 21",
    "--color-success-600": "68 183 19",
    "--color-success-700": "57 152 16",
    "--color-success-800": "46 122 13",
    "--color-success-900": "37 99 10",
    "--color-warning-50": "253 246 223",
    "--color-warning-100": "253 243 212",
    "--color-warning-200": "252 240 202",
    "--color-warning-300": "251 230 170",
    "--color-warning-400": "247 212 106",
    "--color-warning-500": "244 193 42",
    "--color-warning-600": "220 174 38",
    "--color-warning-700": "183 145 32",
    "--color-warning-800": "146 116 25",
    "--color-warning-900": "120 95 21",
    "--color-error-50": "244 223 230",
    "--color-error-100": "240 213 221",
    "--color-error-200": "237 202 213",
    "--color-error-300": "225 171 187",
    "--color-error-400": "203 107 136",
    "--color-error-500": "181 44 85",
    "--color-error-600": "163 40 77",
    "--color-error-700": "136 33 64",
    "--color-error-800": "109 26 51",
    "--color-error-900": "89 22 42",
    "--color-surface-50": "232 234 238",
    "--color-surface-100": "224 227 232",
    "--color-surface-200": "216 220 226",
    "--color-surface-300": "193 199 209",
    "--color-surface-400": "147 158 174",
    "--color-surface-500": "100 116 139",
    "--color-surface-600": "90 104 125",
    "--color-surface-700": "75 87 104",
    "--color-surface-800": "60 70 83",
    "--color-surface-900": "49 57 68"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='rocket'] h1,\n[data-theme='rocket'] h2,\n[data-theme='rocket'] h3,\n[data-theme='rocket'] h4,\n[data-theme='rocket'] h5,\n[data-theme='rocket'] h6": { fontWeight: "bold" },
    "[data-theme='rocket']": {
      backgroundImage: "radial-gradient(at 0% 0%, rgba(var(--color-secondary-500) / 0.33) 0px, transparent 50%),\n			radial-gradient(at 98% 1%, rgba(var(--color-error-500) / 0.33) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var rocket_default = rocket;

// src/tailwind/themes/sahara.ts
var sahara = {
  name: "sahara",
  properties: {
    "--theme-font-family-base": "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,\n		'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'",
    "--theme-font-family-heading": "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue',\n		Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'",
    "--theme-font-color-base": "var(--color-secondary-900)",
    "--theme-font-color-dark": "var(--color-primary-100)",
    "--theme-rounded-base": "9999px",
    "--theme-rounded-container": "24px",
    "--theme-border-base": "1px",
    "--on-primary": "0 0 0",
    "--on-secondary": "0 0 0",
    "--on-tertiary": "0 0 0",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "252 242 225",
    "--color-primary-100": "251 238 215",
    "--color-primary-200": "250 234 205",
    "--color-primary-300": "247 221 175",
    "--color-primary-400": "242 196 114",
    "--color-primary-500": "236 170 54",
    "--color-primary-600": "212 153 49",
    "--color-primary-700": "177 128 41",
    "--color-primary-800": "142 102 32",
    "--color-primary-900": "116 83 26",
    "--color-secondary-50": "225 247 245",
    "--color-secondary-100": "216 245 241",
    "--color-secondary-200": "206 242 238",
    "--color-secondary-300": "176 234 227",
    "--color-secondary-400": "117 219 207",
    "--color-secondary-500": "58 203 186",
    "--color-secondary-600": "52 183 167",
    "--color-secondary-700": "44 152 140",
    "--color-secondary-800": "35 122 112",
    "--color-secondary-900": "28 99 91",
    "--color-tertiary-50": "245 250 237",
    "--color-tertiary-100": "241 249 231",
    "--color-tertiary-200": "238 247 225",
    "--color-tertiary-300": "228 242 207",
    "--color-tertiary-400": "207 233 170",
    "--color-tertiary-500": "187 223 134",
    "--color-tertiary-600": "168 201 121",
    "--color-tertiary-700": "140 167 101",
    "--color-tertiary-800": "112 134 80",
    "--color-tertiary-900": "92 109 66",
    "--color-success-50": "237 247 220",
    "--color-success-100": "230 245 208",
    "--color-success-200": "224 242 197",
    "--color-success-300": "206 235 162",
    "--color-success-400": "169 219 92",
    "--color-success-500": "132 204 22",
    "--color-success-600": "119 184 20",
    "--color-success-700": "99 153 17",
    "--color-success-800": "79 122 13",
    "--color-success-900": "65 100 11",
    "--color-warning-50": "251 246 230",
    "--color-warning-100": "250 243 221",
    "--color-warning-200": "249 240 213",
    "--color-warning-300": "245 230 188",
    "--color-warning-400": "237 212 137",
    "--color-warning-500": "229 193 87",
    "--color-warning-600": "206 174 78",
    "--color-warning-700": "172 145 65",
    "--color-warning-800": "137 116 52",
    "--color-warning-900": "112 95 43",
    "--color-error-50": "250 231 240",
    "--color-error-100": "248 222 235",
    "--color-error-200": "246 214 230",
    "--color-error-300": "241 190 215",
    "--color-error-400": "230 141 186",
    "--color-error-500": "219 92 156",
    "--color-error-600": "197 83 140",
    "--color-error-700": "164 69 117",
    "--color-error-800": "131 55 94",
    "--color-error-900": "107 45 76",
    "--color-surface-50": "249 228 232",
    "--color-surface-100": "248 220 224",
    "--color-surface-200": "246 211 217",
    "--color-surface-300": "240 184 193",
    "--color-surface-400": "229 131 147",
    "--color-surface-500": "218 78 101",
    "--color-surface-600": "196 70 91",
    "--color-surface-700": "164 59 76",
    "--color-surface-800": "131 47 61",
    "--color-surface-900": "107 38 49"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='sahara'] h1,\n[data-theme='sahara'] h2,\n[data-theme='sahara'] h3,\n[data-theme='sahara'] h4,\n[data-theme='sahara'] h5,\n[data-theme='sahara'] h6": { fontWeight: "600" },
    "[data-theme='sahara'] p": { fontWeight: "400" },
    "[data-theme='sahara']": {
      backgroundImage: "radial-gradient(at 100% 36%, hsla(37,81%,56%,0.15) 0px, transparent 50%),\n		radial-gradient(at 7% 0%, hsla(37,81%,56%,0.20) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    },
    ".dark [data-theme='sahara']": {
      backgroundImage: "radial-gradient(at 100% 36%, hsla(37,81%,56%,0.15) 0px, transparent 50%),\n		radial-gradient(at 7% 0%, hsla(37,81%,56%,0.20) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var sahara_default = sahara;

// src/tailwind/themes/seafoam.ts
var seafoam = {
  name: "seafoam",
  properties: {
    "--theme-font-family-base": "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,\n		'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'",
    "--theme-font-family-heading": "'Playfair Display', serif",
    "--theme-font-color-base": "var(--color-surface-900)",
    "--theme-font-color-dark": "var(--color-secondary-100)",
    "--theme-rounded-base": "16px",
    "--theme-rounded-container": "16px",
    "--theme-border-base": "3px",
    "--on-primary": "0 0 0",
    "--on-secondary": "255 255 255",
    "--on-tertiary": "255 255 255",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "0 0 0",
    "--color-primary-50": "237 248 247",
    "--color-primary-100": "231 246 245",
    "--color-primary-200": "225 243 242",
    "--color-primary-300": "207 236 234",
    "--color-primary-400": "170 222 219",
    "--color-primary-500": "134 208 203",
    "--color-primary-600": "121 187 183",
    "--color-primary-700": "101 156 152",
    "--color-primary-800": "80 125 122",
    "--color-primary-900": "66 102 99",
    "--color-secondary-50": "222 224 230",
    "--color-secondary-100": "211 214 221",
    "--color-secondary-200": "200 204 213",
    "--color-secondary-300": "166 173 187",
    "--color-secondary-400": "100 112 136",
    "--color-secondary-500": "33 51 85",
    "--color-secondary-600": "30 46 77",
    "--color-secondary-700": "25 38 64",
    "--color-secondary-800": "20 31 51",
    "--color-secondary-900": "16 25 42",
    "--color-tertiary-50": "255 226 217",
    "--color-tertiary-100": "255 216 204",
    "--color-tertiary-200": "255 207 191",
    "--color-tertiary-300": "255 177 153",
    "--color-tertiary-400": "255 119 77",
    "--color-tertiary-500": "255 61 0",
    "--color-tertiary-600": "230 55 0",
    "--color-tertiary-700": "191 46 0",
    "--color-tertiary-800": "153 37 0",
    "--color-tertiary-900": "125 30 0",
    "--color-success-50": "218 251 241",
    "--color-success-100": "205 250 236",
    "--color-success-200": "193 249 232",
    "--color-success-300": "155 245 218",
    "--color-success-400": "81 237 190",
    "--color-success-500": "6 229 162",
    "--color-success-600": "5 206 146",
    "--color-success-700": "5 172 122",
    "--color-success-800": "4 137 97",
    "--color-success-900": "3 112 79",
    "--color-warning-50": "252 251 230",
    "--color-warning-100": "251 250 221",
    "--color-warning-200": "250 249 213",
    "--color-warning-300": "247 245 188",
    "--color-warning-400": "240 237 137",
    "--color-warning-500": "234 229 87",
    "--color-warning-600": "211 206 78",
    "--color-warning-700": "176 172 65",
    "--color-warning-800": "140 137 52",
    "--color-warning-900": "115 112 43",
    "--color-error-50": "248 227 227",
    "--color-error-100": "246 218 218",
    "--color-error-200": "244 209 209",
    "--color-error-300": "237 181 181",
    "--color-error-400": "224 126 126",
    "--color-error-500": "210 70 70",
    "--color-error-600": "189 63 63",
    "--color-error-700": "158 53 53",
    "--color-error-800": "126 42 42",
    "--color-error-900": "103 34 34",
    "--color-surface-50": "222 248 249",
    "--color-surface-100": "211 246 246",
    "--color-surface-200": "201 244 244",
    "--color-surface-300": "168 237 238",
    "--color-surface-400": "102 223 225",
    "--color-surface-500": "37 209 212",
    "--color-surface-600": "33 188 191",
    "--color-surface-700": "28 157 159",
    "--color-surface-800": "22 125 127",
    "--color-surface-900": "18 102 104"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='seafoam'] h1,\n[data-theme='seafoam'] h2,\n[data-theme='seafoam'] h3,\n[data-theme='seafoam'] h4,\n[data-theme='seafoam'] h5,\n[data-theme='seafoam'] h6": { fontWeight: "bold", fontStyle: "italic", letterSpacing: "1px" },
    "[data-theme='seafoam']": {
      background: "linear-gradient(0deg, rgba(203, 221, 254, 0.75) 0%, rgba(163, 209, 206, 0.75) 100%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    },
    ".dark [data-theme='seafoam']": {
      background: "linear-gradient(0deg, rgba(33, 50, 83, 1) 0%, rgba(8, 132, 124, 1) 100%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var seafoam_default = seafoam;

// src/tailwind/themes/skeleton.ts
var skeleton = {
  name: "skeleton",
  properties: {
    "--theme-font-family-base": "system-ui",
    "--theme-font-family-heading": "system-ui",
    "--theme-font-color-base": "0 0 0",
    "--theme-font-color-dark": "255 255 255",
    "--theme-rounded-base": "9999px",
    "--theme-rounded-container": "8px",
    "--theme-border-base": "1px",
    "--on-primary": "0 0 0",
    "--on-secondary": "255 255 255",
    "--on-tertiary": "0 0 0",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "219 245 236",
    "--color-primary-100": "207 241 230",
    "--color-primary-200": "195 238 224",
    "--color-primary-300": "159 227 205",
    "--color-primary-400": "87 207 167",
    "--color-primary-500": "15 186 129",
    "--color-primary-600": "14 167 116",
    "--color-primary-700": "11 140 97",
    "--color-primary-800": "9 112 77",
    "--color-primary-900": "7 91 63",
    "--color-secondary-50": "229 227 251",
    "--color-secondary-100": "220 218 250",
    "--color-secondary-200": "211 209 249",
    "--color-secondary-300": "185 181 245",
    "--color-secondary-400": "132 126 237",
    "--color-secondary-500": "79 70 229",
    "--color-secondary-600": "71 63 206",
    "--color-secondary-700": "59 53 172",
    "--color-secondary-800": "47 42 137",
    "--color-secondary-900": "39 34 112",
    "--color-tertiary-50": "219 242 252",
    "--color-tertiary-100": "207 237 251",
    "--color-tertiary-200": "195 233 250",
    "--color-tertiary-300": "159 219 246",
    "--color-tertiary-400": "86 192 240",
    "--color-tertiary-500": "14 165 233",
    "--color-tertiary-600": "13 149 210",
    "--color-tertiary-700": "11 124 175",
    "--color-tertiary-800": "8 99 140",
    "--color-tertiary-900": "7 81 114",
    "--color-success-50": "237 247 220",
    "--color-success-100": "230 245 208",
    "--color-success-200": "224 242 197",
    "--color-success-300": "206 235 162",
    "--color-success-400": "169 219 92",
    "--color-success-500": "132 204 22",
    "--color-success-600": "119 184 20",
    "--color-success-700": "99 153 17",
    "--color-success-800": "79 122 13",
    "--color-success-900": "65 100 11",
    "--color-warning-50": "252 244 218",
    "--color-warning-100": "251 240 206",
    "--color-warning-200": "250 236 193",
    "--color-warning-300": "247 225 156",
    "--color-warning-400": "240 202 82",
    "--color-warning-500": "234 179 8",
    "--color-warning-600": "211 161 7",
    "--color-warning-700": "176 134 6",
    "--color-warning-800": "140 107 5",
    "--color-warning-900": "115 88 4",
    "--color-error-50": "249 221 234",
    "--color-error-100": "246 209 228",
    "--color-error-200": "244 198 221",
    "--color-error-300": "238 163 200",
    "--color-error-400": "225 94 159",
    "--color-error-500": "212 25 118",
    "--color-error-600": "191 23 106",
    "--color-error-700": "159 19 89",
    "--color-error-800": "127 15 71",
    "--color-error-900": "104 12 58",
    "--color-surface-50": "228 230 238",
    "--color-surface-100": "219 222 233",
    "--color-surface-200": "210 214 227",
    "--color-surface-300": "182 189 210",
    "--color-surface-400": "128 140 177",
    "--color-surface-500": "73 90 143",
    "--color-surface-600": "66 81 129",
    "--color-surface-700": "55 68 107",
    "--color-surface-800": "44 54 86",
    "--color-surface-900": "36 44 70"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='skeleton'] h1,\n[data-theme='skeleton'] h2,\n[data-theme='skeleton'] h3,\n[data-theme='skeleton'] h4,\n[data-theme='skeleton'] h5,\n[data-theme='skeleton'] h6": { fontWeight: "bold" },
    "[data-theme='skeleton']": {
      backgroundImage: "radial-gradient(at 0% 0%, rgba(var(--color-secondary-500) / 0.33) 0px, transparent 50%),\n			radial-gradient(at 98% 1%, rgba(var(--color-error-500) / 0.33) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var skeleton_default = skeleton;

// src/tailwind/themes/vintage.ts
var vintage = {
  name: "vintage",
  properties: {
    "--theme-font-family-base": "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial,\n		'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'",
    "--theme-font-family-heading": "'Abril Fatface', cursive",
    "--theme-font-color-base": "var(--color-primary-900)",
    "--theme-font-color-dark": "var(--color-primary-100)",
    "--theme-rounded-base": "2px",
    "--theme-rounded-container": "4px",
    "--theme-border-base": "1px",
    "--on-primary": "0 0 0",
    "--on-secondary": "0 0 0",
    "--on-tertiary": "0 0 0",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "0 0 0",
    "--on-surface": "255 255 255",
    "--color-primary-50": "252 237 221",
    "--color-primary-100": "251 231 209",
    "--color-primary-200": "250 225 198",
    "--color-primary-300": "247 207 163",
    "--color-primary-400": "240 170 95",
    "--color-primary-500": "234 134 26",
    "--color-primary-600": "211 121 23",
    "--color-primary-700": "176 101 20",
    "--color-primary-800": "140 80 16",
    "--color-primary-900": "115 66 13",
    "--color-secondary-50": "239 248 242",
    "--color-secondary-100": "234 245 237",
    "--color-secondary-200": "229 243 233",
    "--color-secondary-300": "213 235 219",
    "--color-secondary-400": "182 221 192",
    "--color-secondary-500": "151 206 165",
    "--color-secondary-600": "136 185 149",
    "--color-secondary-700": "113 155 124",
    "--color-secondary-800": "91 124 99",
    "--color-secondary-900": "74 101 81",
    "--color-tertiary-50": "218 244 249",
    "--color-tertiary-100": "205 240 246",
    "--color-tertiary-200": "193 237 244",
    "--color-tertiary-300": "155 226 238",
    "--color-tertiary-400": "81 204 225",
    "--color-tertiary-500": "6 182 212",
    "--color-tertiary-600": "5 164 191",
    "--color-tertiary-700": "5 137 159",
    "--color-tertiary-800": "4 109 127",
    "--color-tertiary-900": "3 89 104",
    "--color-success-50": "237 247 231",
    "--color-success-100": "230 245 223",
    "--color-success-200": "224 242 215",
    "--color-success-300": "206 234 190",
    "--color-success-400": "169 219 142",
    "--color-success-500": "132 203 93",
    "--color-success-600": "119 183 84",
    "--color-success-700": "99 152 70",
    "--color-success-800": "79 122 56",
    "--color-success-900": "65 99 46",
    "--color-warning-50": "253 243 222",
    "--color-warning-100": "252 238 211",
    "--color-warning-200": "252 234 200",
    "--color-warning-300": "250 222 167",
    "--color-warning-400": "246 197 101",
    "--color-warning-500": "242 172 35",
    "--color-warning-600": "218 155 32",
    "--color-warning-700": "182 129 26",
    "--color-warning-800": "145 103 21",
    "--color-warning-900": "119 84 17",
    "--color-error-50": "249 236 235",
    "--color-error-100": "247 229 228",
    "--color-error-200": "245 223 221",
    "--color-error-300": "238 203 201",
    "--color-error-400": "226 165 161",
    "--color-error-500": "213 126 120",
    "--color-error-600": "192 113 108",
    "--color-error-700": "160 95 90",
    "--color-error-800": "128 76 72",
    "--color-error-900": "104 62 59",
    "--color-surface-50": "226 225 224",
    "--color-surface-100": "217 215 214",
    "--color-surface-200": "207 205 204",
    "--color-surface-300": "178 175 173",
    "--color-surface-400": "121 115 111",
    "--color-surface-500": "63 55 49",
    "--color-surface-600": "57 50 44",
    "--color-surface-700": "47 41 37",
    "--color-surface-800": "38 33 29",
    "--color-surface-900": "31 27 24"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='vintage'] h1,\n[data-theme='vintage'] h2,\n[data-theme='vintage'] h3,\n[data-theme='vintage'] h4,\n[data-theme='vintage'] h5,\n[data-theme='vintage'] h6": { letterSpacing: "1px" },
    "[data-theme='vintage']": {
      backgroundImage: "radial-gradient(at 100% 0%, hsla(135,34%,70%,0.20) 0px, transparent 50%),\n		radial-gradient(at 85% 100%, hsla(31,83%,50%,0.20) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    },
    ".dark [data-theme='vintage']": {
      backgroundImage: "radial-gradient(at 100% 0%, hsla(135,34%,70%,0.14) 0px, transparent 50%),\n		radial-gradient(at 85% 100%, hsla(31,83%,50%,0.14) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    }
  }
};
var vintage_default = vintage;

// src/tailwind/themes/wintry.ts
var wintry = {
  name: "wintry",
  properties: {
    "--theme-font-family-heading": "Inter, system-ui, sans-serif",
    "--theme-font-family-base": "system-ui",
    "--theme-font-color-base": "23 37 84",
    "--theme-font-color-dark": "255 255 255",
    "--theme-rounded-base": "9999px",
    "--theme-rounded-container": "6px",
    "--theme-border-base": "1px",
    "--on-primary": "0 0 0",
    "--on-secondary": "0 0 0",
    "--on-tertiary": "255 255 255",
    "--on-success": "0 0 0",
    "--on-warning": "0 0 0",
    "--on-error": "255 255 255",
    "--on-surface": "255 255 255",
    "--color-primary-50": "239 246 255",
    "--color-primary-100": "219 234 254",
    "--color-primary-200": "191 219 254",
    "--color-primary-300": "147 197 253",
    "--color-primary-400": "96 165 250",
    "--color-primary-500": "59 130 246",
    "--color-primary-600": "37 99 235",
    "--color-primary-700": "29 78 216",
    "--color-primary-800": "30 64 175",
    "--color-primary-900": "30 58 138",
    "--color-secondary-50": "240 249 255",
    "--color-secondary-100": "224 242 254",
    "--color-secondary-200": "186 230 253",
    "--color-secondary-300": "125 211 252",
    "--color-secondary-400": "56 189 248",
    "--color-secondary-500": "14 165 233",
    "--color-secondary-600": "2 132 199",
    "--color-secondary-700": "3 105 161",
    "--color-secondary-800": "7 89 133",
    "--color-secondary-900": "12 74 110",
    "--color-tertiary-50": "238 242 255",
    "--color-tertiary-100": "224 231 255",
    "--color-tertiary-200": "199 210 254",
    "--color-tertiary-300": "165 180 252",
    "--color-tertiary-400": "129 140 248",
    "--color-tertiary-500": "99 102 241",
    "--color-tertiary-600": "79 70 229",
    "--color-tertiary-700": "67 56 202",
    "--color-tertiary-800": "55 48 163",
    "--color-tertiary-900": "49 46 129",
    "--color-success-50": "237 247 220",
    "--color-success-100": "230 245 208",
    "--color-success-200": "224 242 197",
    "--color-success-300": "206 235 162",
    "--color-success-400": "169 219 92",
    "--color-success-500": "132 204 22",
    "--color-success-600": "119 184 20",
    "--color-success-700": "99 153 17",
    "--color-success-800": "79 122 13",
    "--color-success-900": "65 100 11",
    "--color-warning-50": "252 244 218",
    "--color-warning-100": "251 240 206",
    "--color-warning-200": "250 236 193",
    "--color-warning-300": "247 225 156",
    "--color-warning-400": "240 202 82",
    "--color-warning-500": "234 179 8",
    "--color-warning-600": "211 161 7",
    "--color-warning-700": "176 134 6",
    "--color-warning-800": "140 107 5",
    "--color-warning-900": "115 88 4",
    "--color-error-50": "249 221 234",
    "--color-error-100": "246 209 228",
    "--color-error-200": "244 198 221",
    "--color-error-300": "238 163 200",
    "--color-error-400": "225 94 159",
    "--color-error-500": "212 25 118",
    "--color-error-600": "191 23 106",
    "--color-error-700": "159 19 89",
    "--color-error-800": "127 15 71",
    "--color-error-900": "104 12 58",
    "--color-surface-50": "249 250 251",
    "--color-surface-100": "243 244 246",
    "--color-surface-200": "229 231 235",
    "--color-surface-300": "209 213 219",
    "--color-surface-400": "156 163 175",
    "--color-surface-500": "107 114 128",
    "--color-surface-600": "75 85 99",
    "--color-surface-700": "55 65 81",
    "--color-surface-800": "31 41 55",
    "--color-surface-900": "17 24 39"
  },
  properties_dark: {},
  enhancements: {
    "[data-theme='wintry'] h1,\n[data-theme='wintry'] h2,\n[data-theme='wintry'] h3,\n[data-theme='wintry'] h4,\n[data-theme='wintry'] h5,\n[data-theme='wintry'] h6": { fontWeight: "bold" },
    "[data-theme='wintry']": {
      backgroundImage: "radial-gradient(at 50% 0%, rgba(var(--color-secondary-500) / 0.50) 0px, transparent 75%), radial-gradient(at 100% 0%, rgba(var(--color-tertiary-500) / 0.40) 0px, transparent 50%)",
      backgroundAttachment: "fixed",
      backgroundPosition: "center",
      backgroundRepeat: "no-repeat",
      backgroundSize: "cover"
    },
    ".dark [data-theme='wintry']": {
      backgroundImage: "radial-gradient(at 50% 0%, rgba(var(--color-secondary-500) / 0.18) 0px, transparent 75%), radial-gradient(at 100% 0%, rgba(var(--color-tertiary-500) / 0.18) 0px, transparent 50%)"
    }
  }
};
var wintry_default = wintry;

// src/tailwind/themes/index.ts
var themes = { crimson: crimson_default, "gold-nouveau": gold_nouveau_default, hamlindigo: hamlindigo_default, modern: modern_default, rocket: rocket_default, sahara: sahara_default, seafoam: seafoam_default, skeleton: skeleton_default, vintage: vintage_default, wintry: wintry_default };
function getThemeProperties(themeName) {
  return themes[themeName].properties;
}

// src/tailwind/prefixSelector.ts
var import_postcss_selector_parser = __toESM(require_dist());
function prefixSelector(prefix, selector, prependNegative = false) {
  if (prefix === "") {
    return selector;
  }
  const ast = (0, import_postcss_selector_parser.default)().astSync(selector);
  ast.walkClasses((classSelector) => {
    const baseClass = classSelector.value;
    if (baseClass === "dark") {
      return;
    }
    const shouldPlaceNegativeBeforePrefix = prependNegative && baseClass.startsWith("-");
    classSelector.value = shouldPlaceNegativeBeforePrefix ? `-${prefix}${baseClass.slice(1)}` : `${prefix}${baseClass}`;
  });
  return ast.toString();
}

// src/index.ts
var skeleton2 = import_plugin2.default.withOptions(
  // Plugin Creator
  (options) => {
    return ({ addBase, addComponents, addUtilities }) => {
      const { components, base } = getSkeletonClasses();
      let baseStyles = {};
      let componentStyles = components;
      if (options?.base !== false) {
        addBase(base);
      }
      options?.themes?.custom?.forEach((theme) => {
        baseStyles[`:root [data-theme='${theme.name}']`] = theme.properties;
        if (theme.properties_dark) {
          baseStyles[`.dark [data-theme='${theme.name}']`] = theme.properties_dark;
        }
      });
      options?.themes?.preset?.forEach((theme) => {
        if (typeof theme === "string") {
          const themeName = theme;
          baseStyles[`:root [data-theme='${themeName}']`] = themes[themeName].properties;
          baseStyles[`.dark [data-theme='${themeName}']`] = themes[themeName].properties_dark;
          return;
        }
        if (!("properties" in theme)) {
          baseStyles[`:root [data-theme='${theme.name}']`] = themes[theme.name].properties;
          baseStyles[`.dark [data-theme='${theme.name}']`] = themes[theme.name].properties_dark;
          if (theme.enhancements === true) {
            baseStyles = { ...baseStyles, ...themes[theme.name].enhancements };
          }
          return;
        }
      });
      if (options?.prefix) {
        const prefix = options?.prefix;
        const root = postcss_js_default.parse(components);
        root.walkRules((rule) => {
          rule.selector = prefixSelector(prefix, rule.selector);
        });
        componentStyles = postcss_js_default.objectify(root);
      }
      addBase(baseStyles);
      addUtilities(coreUtilities);
      addComponents(componentStyles, { respectPrefix: false });
    };
  },
  // Config
  () => {
    return { ...coreConfig };
  }
);
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  getThemeProperties,
  skeleton
});
/*! Bundled license information:

cssesc/cssesc.js:
  (*! https://mths.be/cssesc v3.0.0 by @mathias *)
*/
//# sourceMappingURL=index.js.map