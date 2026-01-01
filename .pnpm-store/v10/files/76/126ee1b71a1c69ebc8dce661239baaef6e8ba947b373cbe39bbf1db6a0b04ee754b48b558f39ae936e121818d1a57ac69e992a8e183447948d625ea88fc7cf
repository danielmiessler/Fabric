"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseAttributes = parseAttributes;
const espree_1 = require("./espree");
const RE_IS_SPACE = /^\s$/u;
class State {
    constructor(code, index) {
        this.code = code;
        this.index = index;
    }
    getCurr() {
        return this.code[this.index];
    }
    skipSpaces() {
        while (this.currIsSpace()) {
            this.advance();
            if (this.eof())
                break;
        }
    }
    currIsSpace() {
        return RE_IS_SPACE.test(this.getCurr() || "");
    }
    currIs(expect) {
        return this.code.startsWith(expect, this.index);
    }
    eof() {
        return this.index >= this.code.length;
    }
    eat(expect) {
        if (!this.currIs(expect)) {
            return null;
        }
        this.index += expect.length;
        return expect;
    }
    advance() {
        this.index++;
        return this.getCurr();
    }
}
/** Parse HTML attributes */
function parseAttributes(code, startIndex) {
    const attributes = [];
    const state = new State(code, startIndex);
    while (!state.eof()) {
        state.skipSpaces();
        if (state.currIs(">") || state.currIs("/>") || state.eof())
            break;
        attributes.push(parseAttribute(state));
    }
    return { attributes, index: state.index };
}
/** Parse HTML attribute */
function parseAttribute(state) {
    const start = state.index;
    // parse key
    const key = parseAttributeKey(state);
    const keyEnd = state.index;
    state.skipSpaces();
    if (!state.eat("=")) {
        return {
            type: "Attribute",
            name: key,
            value: true,
            start,
            end: keyEnd,
        };
    }
    state.skipSpaces();
    if (state.eof()) {
        return {
            type: "Attribute",
            name: key,
            value: true,
            start,
            end: keyEnd,
        };
    }
    // parse value
    const value = parseAttributeValue(state);
    return {
        type: "Attribute",
        name: key,
        value: [value],
        start,
        end: state.index,
    };
}
/** Parse HTML attribute key */
function parseAttributeKey(state) {
    const start = state.index;
    while (state.advance()) {
        if (state.currIs("=") ||
            state.currIs(">") ||
            state.currIs("/>") ||
            state.currIsSpace()) {
            break;
        }
    }
    const end = state.index;
    return state.code.slice(start, end);
}
/** Parse HTML attribute value */
function parseAttributeValue(state) {
    const start = state.index;
    const quote = state.eat('"') || state.eat("'");
    const startBk = state.index;
    const expression = parseAttributeMustache(state);
    if (expression) {
        if (!quote || state.eat(quote)) {
            const end = state.index;
            return {
                type: "ExpressionTag",
                expression,
                start,
                end,
            };
        }
    }
    state.index = startBk;
    if (quote) {
        if (state.eof()) {
            return {
                type: "Text",
                data: quote,
                raw: quote,
                start,
                end: state.index,
            };
        }
        let c;
        while ((c = state.getCurr())) {
            state.advance();
            if (c === quote) {
                const end = state.index;
                const data = state.code.slice(start + 1, end - 1);
                return {
                    type: "Text",
                    data,
                    raw: data,
                    start: start + 1,
                    end: end - 1,
                };
            }
        }
    }
    else {
        while (state.advance()) {
            if (state.currIsSpace() || state.currIs(">") || state.currIs("/>")) {
                break;
            }
        }
    }
    const end = state.index;
    const data = state.code.slice(start, end);
    return {
        type: "Text",
        data,
        raw: data,
        start,
        end,
    };
}
/** Parse mustache */
function parseAttributeMustache(state) {
    if (!state.eat("{")) {
        return null;
    }
    // parse simple expression
    const leadingComments = [];
    const startBk = state.index;
    state.skipSpaces();
    let start = state.index;
    while (!state.eof()) {
        if (state.eat("//")) {
            leadingComments.push(parseInlineComment(state.index - 2));
            state.skipSpaces();
            start = state.index;
            continue;
        }
        if (state.eat("/*")) {
            leadingComments.push(parseBlockComment(state.index - 2));
            state.skipSpaces();
            start = state.index;
            continue;
        }
        const stringQuote = state.eat('"') || state.eat("'");
        if (stringQuote) {
            skipString(stringQuote);
            state.skipSpaces();
            continue;
        }
        const endCandidate = state.index;
        state.skipSpaces();
        if (state.eat("}")) {
            const end = endCandidate;
            try {
                const espree = (0, espree_1.getEspree)();
                const expression = espree.parse(state.code.slice(start, end), {
                    ecmaVersion: espree.latestEcmaVersion,
                }).body[0].expression;
                delete expression.range;
                return Object.assign(Object.assign({}, expression), { leadingComments,
                    start,
                    end });
            }
            catch (_a) {
                break;
            }
        }
        state.advance();
    }
    state.index = startBk;
    return null;
    function parseInlineComment(tokenStart) {
        const valueStart = state.index;
        let valueEnd = null;
        while (!state.eof()) {
            if (state.eat("\n")) {
                valueEnd = state.index - 1;
                break;
            }
            state.advance();
        }
        if (valueEnd == null) {
            valueEnd = state.index;
        }
        return {
            type: "Line",
            value: state.code.slice(valueStart, valueEnd),
            start: tokenStart,
            end: state.index,
        };
    }
    function parseBlockComment(tokenStart) {
        const valueStart = state.index;
        let valueEnd = null;
        while (!state.eof()) {
            if (state.eat("*/")) {
                valueEnd = state.index - 2;
                break;
            }
            state.advance();
        }
        if (valueEnd == null) {
            valueEnd = state.index;
        }
        return {
            type: "Block",
            value: state.code.slice(valueStart, valueEnd),
            start: tokenStart,
            end: state.index,
        };
    }
    function skipString(stringQuote) {
        while (!state.eof()) {
            if (state.eat(stringQuote)) {
                break;
            }
            if (state.eat("\\")) {
                // escape
                state.advance();
            }
            state.advance();
        }
    }
}
