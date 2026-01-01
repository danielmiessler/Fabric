"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.LinesAndColumns = exports.Context = exports.ScriptsSourceCode = void 0;
const script_let_1 = require("./script-let");
const let_directive_collection_1 = require("./let-directive-collection");
const html_1 = require("../parser/html");
const utils_1 = require("../utils");
const parser_options_1 = require("../parser/parser-options");
class ScriptsSourceCode {
    constructor(script, attrs) {
        this._appendScriptLets = null;
        this.separateIndexes = [];
        this.raw = script;
        this.trimmedRaw = script.trimEnd();
        this.attrs = attrs;
        this.separateIndexes = [script.length];
    }
    getCurrentVirtualCode() {
        if (this._appendScriptLets == null) {
            return this.raw;
        }
        return (this.trimmedRaw +
            this._appendScriptLets.separate +
            this._appendScriptLets.beforeSpaces +
            this._appendScriptLets.render +
            this._appendScriptLets.snippet +
            this._appendScriptLets.generics);
    }
    getCurrentVirtualCodeInfo() {
        if (this._appendScriptLets == null) {
            return { script: this.raw, render: "", rootScope: "" };
        }
        return {
            script: this.trimmedRaw + this._appendScriptLets.separate,
            render: this._appendScriptLets.beforeSpaces + this._appendScriptLets.render,
            rootScope: this._appendScriptLets.snippet + this._appendScriptLets.generics,
        };
    }
    getCurrentVirtualCodeLength() {
        if (this._appendScriptLets == null) {
            return this.raw.length;
        }
        return (this.trimmedRaw.length +
            this._appendScriptLets.separate.length +
            this._appendScriptLets.beforeSpaces.length +
            this._appendScriptLets.render.length +
            this._appendScriptLets.snippet.length +
            this._appendScriptLets.generics.length);
    }
    addLet(letCode, kind) {
        if (this._appendScriptLets == null) {
            const currentLength = this.trimmedRaw.length;
            this.separateIndexes = [currentLength, currentLength + 1];
            const after = this.raw.slice(currentLength + 2);
            this._appendScriptLets = {
                separate: "\n;",
                beforeSpaces: after,
                render: "",
                snippet: "",
                generics: "",
            };
        }
        const start = this.getCurrentVirtualCodeLength();
        this._appendScriptLets[kind] += letCode;
        return {
            start,
            end: this.getCurrentVirtualCodeLength(),
        };
    }
    stripCode(start, end) {
        this.raw =
            this.raw.slice(0, start) +
                this.raw.slice(start, end).replace(/[^\n\r ]/g, " ") +
                this.raw.slice(end);
        this.trimmedRaw =
            this.trimmedRaw.slice(0, start) +
                this.trimmedRaw.slice(start, end).replace(/[^\n\r ]/g, " ") +
                this.trimmedRaw.slice(end);
    }
}
exports.ScriptsSourceCode = ScriptsSourceCode;
class Context {
    constructor(code, parserOptions) {
        this.tokens = [];
        this.comments = [];
        this.locsMap = new Map();
        this.letDirCollections = new let_directive_collection_1.LetDirectiveCollections();
        this.slots = new Set();
        this.elements = new Map();
        this.snippets = [];
        // ----- States ------
        this.state = {};
        this.blocks = [];
        this.code = code;
        this.parserOptions = parserOptions;
        this.locs = new LinesAndColumns(code);
        const spaces = code.replace(/[^\n\r ]/g, " ");
        let templateCode = "";
        let scriptCode = "";
        const scriptAttrs = {};
        let start = 0;
        for (const block of extractBlocks(code)) {
            if (block.tag === "template") {
                if (block.selfClosing) {
                    continue;
                }
                const lang = block.attrs.find((attr) => attr.name === "lang");
                if (!lang || !Array.isArray(lang.value)) {
                    continue;
                }
                const langValue = lang.value[0];
                if (!langValue ||
                    langValue.type !== "Text" ||
                    langValue.data === "html") {
                    continue;
                }
            }
            this.blocks.push(block);
            if (block.selfClosing) {
                // Self-closing blocks are temporarily replaced with `<s---->` or `<t---->` tag
                // because the svelte compiler cannot parse self-closing block(script, style) tags.
                // It will be restored later in `convertHTMLElement()` processing.
                templateCode += `${code.slice(start, block.startTagRange[0] + 2 /* `<` and first letter */)}${"-".repeat(block.tag.length - 1 /* skip first letter */)}${code.slice(block.startTagRange[0] + 1 /* skip `<` */ + block.tag.length, block.startTagRange[1])}`;
                scriptCode += spaces.slice(start, block.startTagRange[1]);
                start = block.startTagRange[1];
            }
            else {
                templateCode +=
                    code.slice(start, block.contentRange[0]) +
                        spaces.slice(block.contentRange[0], block.contentRange[1]);
                if (block.tag === "script") {
                    scriptCode +=
                        spaces.slice(start, block.contentRange[0]) +
                            code.slice(...block.contentRange);
                    for (const attr of block.attrs) {
                        if (Array.isArray(attr.value)) {
                            const attrValue = attr.value[0];
                            scriptAttrs[attr.name] =
                                attrValue && attrValue.type === "Text"
                                    ? attrValue.data
                                    : undefined;
                        }
                    }
                }
                else {
                    scriptCode += spaces.slice(start, block.contentRange[1]);
                }
                start = block.contentRange[1];
            }
        }
        templateCode += code.slice(start);
        scriptCode += spaces.slice(start);
        this.sourceCode = {
            template: templateCode,
            scripts: new ScriptsSourceCode(scriptCode, scriptAttrs),
        };
        this.scriptLet = new script_let_1.ScriptLetContext(this);
    }
    getLocFromIndex(index) {
        let loc = this.locsMap.get(index);
        if (!loc) {
            loc = this.locs.getLocFromIndex(index);
            this.locsMap.set(index, loc);
        }
        return {
            line: loc.line,
            column: loc.column,
        };
    }
    getIndexFromLoc(loc) {
        return this.locs.getIndexFromLoc(loc);
    }
    /**
     * Get the location information of the given node.
     * @param node The node.
     */
    getConvertLocation(node) {
        const { start, end } = node;
        return {
            range: [start, end],
            loc: {
                start: this.getLocFromIndex(start),
                end: this.getLocFromIndex(end),
            },
        };
    }
    addComment(comment) {
        this.comments.push(comment);
    }
    /**
     * Add token to tokens
     */
    addToken(type, range) {
        const token = Object.assign({ type, value: this.getText(range) }, this.getConvertLocation(range));
        this.tokens.push(token);
        return token;
    }
    /**
     * get text
     */
    getText(range) {
        return this.code.slice(range.start, range.end);
    }
    isTypeScript() {
        if (this.state.isTypeScript != null) {
            return this.state.isTypeScript;
        }
        const lang = this.sourceCode.scripts.attrs.lang;
        return (this.state.isTypeScript = (0, parser_options_1.isTypeScript)(this.parserOptions, lang));
    }
    stripScriptCode(start, end) {
        this.sourceCode.scripts.stripCode(start, end);
    }
    findBlock(element) {
        const tag = element.type === "SvelteScriptElement"
            ? "script"
            : element.type === "SvelteStyleElement"
                ? "style"
                : element.name.name.toLowerCase();
        return this.blocks.find((block) => block.tag === tag &&
            !block.selfClosing &&
            element.range[0] <= block.contentRange[0] &&
            block.contentRange[1] <= element.range[1]);
    }
    findSelfClosingBlock(element) {
        return this.blocks.find((block) => Boolean(block.selfClosing &&
            element.startTag.range[0] <= block.startTagRange[0] &&
            block.startTagRange[1] <= element.startTag.range[1]));
    }
}
exports.Context = Context;
/** Extract <script> blocks */
function* extractBlocks(code) {
    const startTagOpenRe = /<!--[\s\S]*?-->|<(script|style|template)([\s>])/giu;
    const endScriptTagRe = /<\/script>/giu;
    const endStyleTagRe = /<\/style>/giu;
    const endTemplateTagRe = /<\/template>/giu;
    let startTagOpenMatch;
    while ((startTagOpenMatch = startTagOpenRe.exec(code))) {
        const [, tag, nextChar] = startTagOpenMatch;
        if (!tag) {
            continue;
        }
        const startTagStart = startTagOpenMatch.index;
        let startTagEnd = startTagOpenRe.lastIndex;
        const lowerTag = tag.toLowerCase();
        let attrs = [];
        if (!nextChar.trim()) {
            const attrsData = (0, html_1.parseAttributes)(code, startTagOpenRe.lastIndex);
            attrs = attrsData.attributes;
            startTagEnd = attrsData.index;
            if (code[startTagEnd] === "/" && code[startTagEnd + 1] === ">") {
                yield {
                    tag: lowerTag,
                    originalTag: tag,
                    attrs,
                    selfClosing: true,
                    startTagRange: [startTagStart, startTagEnd + 2],
                };
                continue;
            }
            if (code[startTagEnd] === ">") {
                startTagEnd++;
            }
            else {
                continue;
            }
        }
        const endTagRe = lowerTag === "script"
            ? endScriptTagRe
            : lowerTag === "style"
                ? endStyleTagRe
                : endTemplateTagRe;
        endTagRe.lastIndex = startTagEnd;
        const endTagMatch = endTagRe.exec(code);
        if (endTagMatch) {
            const endTagStart = endTagMatch.index;
            const endTagEnd = endTagRe.lastIndex;
            yield {
                tag: lowerTag,
                originalTag: tag,
                attrs,
                startTagRange: [startTagStart, startTagEnd],
                contentRange: [startTagEnd, endTagStart],
                endTagRange: [endTagStart, endTagEnd],
            };
            startTagOpenRe.lastIndex = endTagEnd;
        }
    }
}
class LinesAndColumns {
    constructor(code) {
        const len = code.length;
        const lineStartIndices = [0];
        for (let index = 0; index < len; index++) {
            const c = code[index];
            if (c === "\r") {
                const next = code[index + 1] || "";
                if (next === "\n") {
                    index++;
                }
                lineStartIndices.push(index + 1);
            }
            else if (c === "\n") {
                lineStartIndices.push(index + 1);
            }
        }
        this.lineStartIndices = lineStartIndices;
    }
    getLocFromIndex(index) {
        const lineNumber = (0, utils_1.sortedLastIndex)(this.lineStartIndices, (target) => target - index);
        return {
            line: lineNumber,
            column: index - this.lineStartIndices[lineNumber - 1],
        };
    }
    getIndexFromLoc(loc) {
        const lineStartIndex = this.lineStartIndices[loc.line - 1];
        const positionIndex = lineStartIndex + loc.column;
        return positionIndex;
    }
    /**
     * Get the location information of the given indexes.
     */
    getLocations(start, end) {
        return {
            range: [start, end],
            loc: {
                start: this.getLocFromIndex(start),
                end: this.getLocFromIndex(end),
            },
        };
    }
}
exports.LinesAndColumns = LinesAndColumns;
