"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseStyleContext = parseStyleContext;
exports.styleNodeLoc = styleNodeLoc;
exports.styleNodeRange = styleNodeRange;
const postcss_1 = __importDefault(require("postcss"));
const postcss_scss_1 = require("postcss-scss");
/**
 * Extracts style source from a SvelteStyleElement and parses it into a PostCSS AST.
 */
function parseStyleContext(styleElement, ctx) {
    if (!styleElement || !styleElement.endTag) {
        return { status: "no-style-element" };
    }
    let sourceLang = "css";
    for (const attribute of styleElement.startTag.attributes) {
        if (attribute.type === "SvelteAttribute" &&
            attribute.key.name === "lang" &&
            attribute.value.length > 0 &&
            attribute.value[0].type === "SvelteLiteral") {
            sourceLang = attribute.value[0].value;
        }
    }
    let parseFn, sourceAst;
    switch (sourceLang) {
        case "css":
            parseFn = postcss_1.default.parse;
            break;
        case "scss":
            parseFn = postcss_scss_1.parse;
            break;
        default:
            return { status: "unknown-lang", sourceLang };
    }
    const styleCode = ctx.code.slice(styleElement.startTag.range[1], styleElement.endTag.range[0]);
    try {
        sourceAst = parseFn(styleCode, {
            from: ctx.parserOptions.filePath,
        });
    }
    catch (error) {
        return { status: "parse-error", sourceLang, error };
    }
    fixPostCSSNodeLocation(sourceAst, styleElement);
    sourceAst.walk((node) => fixPostCSSNodeLocation(node, styleElement));
    return { status: "success", sourceLang, sourceAst };
}
/**
 * Extracts a node location (like that of any ESLint node) from a parsed svelte style node.
 */
function styleNodeLoc(node) {
    if (node.source === undefined) {
        return {};
    }
    return {
        start: node.source.start !== undefined
            ? {
                line: node.source.start.line,
                column: node.source.start.column - 1,
            }
            : undefined,
        end: node.source.end !== undefined
            ? {
                line: node.source.end.line,
                column: node.source.end.column,
            }
            : undefined,
    };
}
/**
 * Extracts a node range (like that of any ESLint node) from a parsed svelte style node.
 */
function styleNodeRange(node) {
    if (node.source === undefined) {
        return [undefined, undefined];
    }
    return [
        node.source.start !== undefined ? node.source.start.offset : undefined,
        node.source.end !== undefined ? node.source.end.offset + 1 : undefined,
    ];
}
/**
 * Fixes PostCSS AST locations to be relative to the whole file instead of relative to the <style> element.
 */
function fixPostCSSNodeLocation(node, styleElement) {
    var _a, _b, _c, _d, _e, _f, _g, _h, _j, _k, _l, _m;
    if (((_b = (_a = node.source) === null || _a === void 0 ? void 0 : _a.start) === null || _b === void 0 ? void 0 : _b.offset) !== undefined) {
        node.source.start.offset += styleElement.startTag.range[1];
    }
    if (((_d = (_c = node.source) === null || _c === void 0 ? void 0 : _c.start) === null || _d === void 0 ? void 0 : _d.line) !== undefined) {
        node.source.start.line += styleElement.loc.start.line - 1;
    }
    if (((_f = (_e = node.source) === null || _e === void 0 ? void 0 : _e.end) === null || _f === void 0 ? void 0 : _f.offset) !== undefined) {
        node.source.end.offset += styleElement.startTag.range[1];
    }
    if (((_h = (_g = node.source) === null || _g === void 0 ? void 0 : _g.end) === null || _h === void 0 ? void 0 : _h.line) !== undefined) {
        node.source.end.line += styleElement.loc.start.line - 1;
    }
    if (((_k = (_j = node.source) === null || _j === void 0 ? void 0 : _j.start) === null || _k === void 0 ? void 0 : _k.line) === styleElement.loc.start.line) {
        node.source.start.column += styleElement.startTag.loc.end.column;
    }
    if (((_m = (_l = node.source) === null || _l === void 0 ? void 0 : _l.end) === null || _m === void 0 ? void 0 : _m.line) === styleElement.loc.start.line) {
        node.source.end.column += styleElement.startTag.loc.end.column;
    }
}
