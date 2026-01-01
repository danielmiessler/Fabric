"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
exports.defineVisitor = defineVisitor;
const SV = __importStar(require("./svelte"));
const ES = __importStar(require("./es"));
const TS = __importStar(require("./ts"));
const ast_1 = require("./ast");
const eslint_utils_1 = require("@eslint-community/eslint-utils");
const offset_context_1 = require("./offset-context");
const compat_1 = require("../../utils/compat");
/**
 * Normalize options.
 * @param type The type of indentation.
 * @param options Other options.
 * @param defaultOptions The default value of options.
 * @returns Normalized options.
 */
function parseOptions(options, defaultOptions) {
    const ret = {
        indentChar: ' ',
        indentScript: true,
        indentSize: 2,
        switchCase: 1,
        alignAttributesVertically: false,
        ignoredNodes: [],
        ...defaultOptions
    };
    if (Number.isSafeInteger(options.indent)) {
        ret.indentSize = Number(options.indent);
    }
    else if (options.indent === 'tab') {
        ret.indentChar = '\t';
        ret.indentSize = 1;
    }
    if (typeof options.indentScript === 'boolean') {
        ret.indentScript = options.indentScript;
    }
    if (options.switchCase != null && Number.isSafeInteger(options.switchCase)) {
        ret.switchCase = options.switchCase;
    }
    if (options.ignoredNodes != null) {
        ret.ignoredNodes = options.ignoredNodes;
    }
    if (options.alignAttributesVertically && ret.indentChar === ' ') {
        ret.alignAttributesVertically = true;
    }
    else if (ret.indentChar !== ' ') {
        ret.alignAttributesVertically = false;
    }
    return ret;
}
/**
 * Creates AST event handlers for html-indent.
 *
 * @param context The rule context.
 * @param defaultOptions The default value of options.
 * @returns AST event handlers.
 */
function defineVisitor(context, defaultOptions) {
    if (!(0, compat_1.getFilename)(context).endsWith('.svelte'))
        return {};
    const options = parseOptions(context.options[0] || {}, defaultOptions);
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const offsets = new offset_context_1.OffsetContext({ sourceCode, options });
    /**
     * Get the text of the indentation part of the given location.
     */
    function getIndentText({ line, column }) {
        return sourceCode.lines[line - 1].slice(0, column);
    }
    /**
     * Validate the given token with the pre-calculated expected indentation.
     */
    function validateToken(token, expectedIndent) {
        const line = token.loc.start.line;
        const indentText = getIndentText(token.loc.start);
        // `indentText` contains non-whitespace characters.
        if (indentText.trim() !== '') {
            return;
        }
        const actualIndent = token.loc.start.column;
        const mismatchCharIndexes = [];
        for (let i = 0; i < indentText.length; ++i) {
            if (indentText[i] !== options.indentChar) {
                mismatchCharIndexes.push(i);
            }
        }
        if (actualIndent !== expectedIndent) {
            const loc = {
                start: { line, column: 0 },
                end: { line, column: actualIndent }
            };
            context.report({
                loc,
                messageId: 'unexpectedIndentation',
                data: {
                    expectedIndent: String(expectedIndent),
                    actualIndent: String(actualIndent),
                    expectedUnit: options.indentChar === '\t' ? 'tab' : 'space',
                    actualUnit: mismatchCharIndexes.length
                        ? 'whitespace'
                        : options.indentChar === '\t'
                            ? 'tab'
                            : 'space',
                    expectedIndentPlural: expectedIndent === 1 ? '' : 's',
                    actualIndentPlural: actualIndent === 1 ? '' : 's'
                },
                fix(fixer) {
                    return fixer.replaceTextRange([sourceCode.getIndexFromLoc(loc.start), sourceCode.getIndexFromLoc(loc.end)], options.indentChar.repeat(expectedIndent));
                }
            });
            return;
        }
        for (const i of mismatchCharIndexes) {
            const loc = {
                start: { line, column: i },
                end: { line, column: i + 1 }
            };
            context.report({
                loc,
                messageId: 'unexpectedChar',
                data: {
                    expected: JSON.stringify(options.indentChar),
                    actual: JSON.stringify(indentText[i])
                },
                fix(fixer) {
                    return fixer.replaceTextRange([sourceCode.getIndexFromLoc(loc.start), sourceCode.getIndexFromLoc(loc.end)], options.indentChar);
                }
            });
        }
    }
    /** Process line tokens */
    function processLine(tokens, prevComments, prevToken, calculator) {
        const firstToken = tokens[0];
        const actualIndent = firstToken.loc.start.column;
        const expectedIndent = calculator.getExpectedIndentFromTokens(tokens);
        if (expectedIndent == null) {
            calculator.saveExpectedIndent(tokens, actualIndent);
            return;
        }
        calculator.saveExpectedIndent(tokens, Math.min(...tokens
            .map((t) => calculator.getExpectedIndentFromToken(t))
            .filter((i) => i != null)));
        let prev = prevToken;
        if (prevComments.length) {
            if (prev && prev.loc.end.line < prevComments[0].loc.start.line) {
                validateToken(prevComments[0], expectedIndent);
            }
            prev = prevComments[prevComments.length - 1];
        }
        if (prev && prev.loc.end.line < tokens[0].loc.start.line) {
            validateToken(tokens[0], expectedIndent);
        }
    }
    const indentContext = {
        sourceCode,
        options,
        offsets
    };
    const nodesVisitor = {
        ...ES.defineVisitor(indentContext),
        ...SV.defineVisitor(indentContext),
        ...TS.defineVisitor(indentContext)
    };
    const knownNodes = new Set(Object.keys(nodesVisitor));
    /**
     * Build a visitor combined with a visitor to handle the given ignore selector.
     */
    function compositingIgnoresVisitor(visitor) {
        for (const ignoreSelector of options.ignoredNodes) {
            const key = `${ignoreSelector}:exit`;
            if (visitor[key]) {
                // eslint-disable-next-line @typescript-eslint/no-explicit-any -- ignore
                const handler = visitor[key];
                visitor[key] = function (node, ...args) {
                    const ret = handler.call(this, node, ...args);
                    offsets.ignore(node);
                    return ret;
                };
            }
            else {
                visitor[key] = (node) => offsets.ignore(node);
            }
        }
        return visitor;
    }
    return compositingIgnoresVisitor({
        ...nodesVisitor,
        '*:exit'(node) {
            // Ignore tokens of unknown nodes.
            if (!knownNodes.has(node.type)) {
                // debugger
                // console.log(node.type, node.loc!.start.line)
                offsets.ignore(node);
            }
        },
        'Program:exit'(node) {
            const calculator = offsets.getOffsetCalculator();
            let prevToken = null;
            for (const { prevComments, tokens } of iterateLineTokens()) {
                processLine(tokens, prevComments, prevToken, calculator);
                prevToken = tokens[tokens.length - 1];
            }
            /** Iterate line tokens */
            function* iterateLineTokens() {
                let line = 0;
                let prevComments = [];
                let bufferTokens = [];
                for (const token of sourceCode.getTokens(node, {
                    includeComments: true,
                    filter: ast_1.isNotWhitespace
                })) {
                    const thisLine = token.loc.start.line;
                    if (line === thisLine || bufferTokens.length === 0) {
                        bufferTokens.push(token);
                    }
                    else {
                        if ((0, eslint_utils_1.isCommentToken)(bufferTokens[0]) && bufferTokens.every(eslint_utils_1.isCommentToken)) {
                            prevComments.push(bufferTokens[0]);
                        }
                        else {
                            yield {
                                prevComments,
                                tokens: bufferTokens
                            };
                            prevComments = [];
                        }
                        bufferTokens = [token];
                    }
                    line = thisLine;
                }
                if (bufferTokens.length && !bufferTokens.every(eslint_utils_1.isCommentToken)) {
                    yield {
                        prevComments,
                        tokens: bufferTokens
                    };
                }
            }
        }
    });
}
