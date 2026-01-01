"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.OffsetCalculator = exports.OffsetContext = void 0;
const ast_1 = require("./ast");
const commons_1 = require("./commons");
const commons_2 = require("./commons");
class OffsetContext {
    constructor(arg) {
        this.offsets = new Map();
        this.ignoreRanges = new Map();
        this.sourceCode = arg.sourceCode;
        this.options = arg.options;
    }
    /**
     * Set offset to the given index.
     */
    setOffsetIndex(index, offset, base) {
        if (index === base) {
            return;
        }
        this.offsets.set(index, {
            type: 0 /* OffsetDataType.normal */,
            base,
            offset
        });
    }
    /**
     * Set align indent to the given index.
     */
    setAlignIndent(index, alignIndent, base) {
        if (index === base) {
            return;
        }
        this.offsets.set(index, {
            type: 1 /* OffsetDataType.align */,
            base,
            alignIndent
        });
    }
    /**
     * Set offset to the given tokens.
     */
    setOffsetToken(token, offset, baseToken) {
        if (!token) {
            return;
        }
        if (Array.isArray(token)) {
            for (const t of token) {
                this.setOffsetToken(t, offset, baseToken);
            }
            return;
        }
        this.setOffsetIndex(token.range[0], offset, baseToken.range[0]);
    }
    /**
     * Copy offset to the given index from srcIndex.
     */
    copyOffset(index, srcIndex) {
        const offsetData = this.offsets.get(srcIndex);
        if (!offsetData) {
            return;
        }
        if (offsetData.type === 2 /* OffsetDataType.start */) {
            this.setStartOffsetIndex(index, offsetData.offset);
        }
        else if (offsetData.type === 1 /* OffsetDataType.align */) {
            this.setAlignIndent(index, offsetData.alignIndent, offsetData.base);
        }
        else {
            this.setOffsetIndex(index, offsetData.offset, offsetData.base);
        }
    }
    /**
     * Set start offset to the given index.
     */
    setStartOffsetIndex(index, offset) {
        this.offsets.set(index, {
            type: 2 /* OffsetDataType.start */,
            offset
        });
    }
    /**
     * Set start offset to the given token.
     */
    setStartOffsetToken(token, offset) {
        if (!token) {
            return;
        }
        if (Array.isArray(token)) {
            for (const t of token) {
                this.setStartOffsetToken(t, offset);
            }
            return;
        }
        this.setStartOffsetIndex(token.range[0], offset);
    }
    setOffsetElementList(nodes, baseNodeOrToken, lastNodeOrToken, offset, align) {
        let setIndent = (token, baseToken) => this.setOffsetToken(token, offset, baseToken);
        if (align) {
            for (const n of nodes) {
                if (n) {
                    if (!(0, commons_1.isBeginningOfLine)(this.sourceCode, n)) {
                        const startLoc = n.loc.start;
                        const alignIndent = startLoc.column - /^\s*/u.exec(this.sourceCode.lines[startLoc.line - 1])[0].length;
                        setIndent = (token, baseToken) => this.setAlignIndent(token.range[0], alignIndent, baseToken.range[0]);
                    }
                    break;
                }
            }
        }
        this._setOffsetElementList(nodes, baseNodeOrToken, lastNodeOrToken, setIndent);
    }
    _setOffsetElementList(nodes, baseNodeOrToken, lastNodeOrToken, setIndent) {
        const baseToken = this.sourceCode.getFirstToken(baseNodeOrToken);
        let prevToken = this.sourceCode.getLastToken(baseNodeOrToken);
        for (const node of nodes) {
            if (node == null) {
                continue;
            }
            const elementTokens = (0, commons_2.getFirstAndLastTokens)(this.sourceCode, node, prevToken.range[1]);
            let t = prevToken;
            while ((t = this.sourceCode.getTokenAfter(t, {
                includeComments: true,
                filter: ast_1.isNotWhitespace
            })) != null &&
                t.range[1] <= elementTokens.firstToken.range[0]) {
                setIndent(t, baseToken);
            }
            setIndent(elementTokens.firstToken, baseToken);
            prevToken = elementTokens.lastToken;
        }
        if (lastNodeOrToken) {
            const lastToken = this.sourceCode.getFirstToken(lastNodeOrToken);
            let t = prevToken;
            while ((t = this.sourceCode.getTokenAfter(t, {
                includeComments: true,
                filter: ast_1.isNotWhitespace
            })) != null &&
                t.range[1] <= lastToken.range[0]) {
                setIndent(t, baseToken);
            }
            this.setOffsetToken(lastToken, 0, baseToken);
        }
    }
    /**
     * Ignore range of the given node.
     */
    ignore(node) {
        const range = node.range;
        const n = this.ignoreRanges.get(range[0]) ?? 0;
        this.ignoreRanges.set(range[0], Math.max(n, range[1]));
    }
    getOffsetCalculator() {
        // eslint-disable-next-line @typescript-eslint/no-use-before-define -- ignore
        return new OffsetCalculator({
            offsets: this.offsets,
            options: this.options,
            ignoreRanges: [...this.ignoreRanges]
        });
    }
}
exports.OffsetContext = OffsetContext;
class OffsetCalculator {
    constructor(arg) {
        this.offsets = arg.offsets;
        this.options = arg.options;
        this.ignoreRanges = arg.ignoreRanges;
    }
    /**
     * Calculate correct indentation of the given index.
     */
    getExpectedIndentFromIndex(index) {
        const offsetInfo = this.offsets.get(index);
        if (offsetInfo == null) {
            return null;
        }
        if (offsetInfo.expectedIndent != null) {
            return offsetInfo.expectedIndent;
        }
        if (offsetInfo.type === 2 /* OffsetDataType.start */) {
            return offsetInfo.offset * this.options.indentSize;
        }
        const baseIndent = this.getExpectedIndentFromIndex(offsetInfo.base);
        if (baseIndent == null) {
            return null;
        }
        if (offsetInfo.type === 1 /* OffsetDataType.align */) {
            return baseIndent + offsetInfo.alignIndent;
        }
        return baseIndent + offsetInfo.offset * this.options.indentSize;
    }
    /**
     * Calculate correct indentation of the given token.
     */
    getExpectedIndentFromToken(token) {
        return this.getExpectedIndentFromIndex(token.range[0]);
    }
    /**
     * Calculate correct indentation of the line of the given tokens.
     */
    getExpectedIndentFromTokens(tokens) {
        for (const token of tokens) {
            const index = token.range[0];
            if (this.ignoreRanges.some(([f, t]) => f <= index && index < t)) {
                return null;
            }
            const expectedIndent = this.getExpectedIndentFromIndex(index);
            if (expectedIndent != null) {
                return expectedIndent;
            }
        }
        return null;
    }
    /** Save expected indent to give tokens */
    saveExpectedIndent(tokens, expectedIndent) {
        for (const token of tokens) {
            const offsetInfo = this.offsets.get(token.range[0]);
            if (offsetInfo == null) {
                continue;
            }
            offsetInfo.expectedIndent = offsetInfo.expectedIndent ?? expectedIndent;
        }
    }
}
exports.OffsetCalculator = OffsetCalculator;
