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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getSvelteCompileWarnings = getSvelteCompileWarnings;
const compiler = __importStar(require("svelte/compiler"));
const sourcemap_codec_1 = require("@jridgewell/sourcemap-codec");
const lines_and_columns_1 = require("../../utils/lines-and-columns");
const typescript_1 = require("./transform/typescript");
const babel_1 = require("./transform/babel");
const postcss_1 = require("./transform/postcss");
const sass_1 = require("./transform/sass");
const less_1 = require("./transform/less");
const stylus_1 = require("./transform/stylus");
const ignore_comment_1 = require("./ignore-comment");
const extract_leading_comments_1 = require("./extract-leading-comments");
const ast_utils_1 = require("../../utils/ast-utils");
const path_1 = __importDefault(require("path"));
const fs_1 = __importDefault(require("fs"));
const semver_1 = __importDefault(require("semver"));
const compat_1 = require("../../utils/compat");
const STYLE_TRANSFORMS = {
    postcss: postcss_1.transform,
    pcss: postcss_1.transform,
    scss: (node, text, context) => (0, sass_1.transform)(node, text, context, 'scss'),
    sass: (node, text, context) => (0, sass_1.transform)(node, text, context, 'sass'),
    less: less_1.transform,
    stylus: stylus_1.transform,
    styl: stylus_1.transform
};
const CSS_WARN_CODES = new Set([
    'css-unused-selector',
    'css_unused_selector',
    'css-invalid-global',
    'css-invalid-global-selector'
]);
const cacheAll = new WeakMap();
/**
 * Get svelte compile warnings
 */
function getSvelteCompileWarnings(context) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const cache = cacheAll.get(sourceCode.ast);
    if (cache) {
        return cache;
    }
    const result = getSvelteCompileWarningsWithoutCache(context);
    cacheAll.set(sourceCode.ast, result);
    return result;
}
/**
 * Get svelte compile warnings
 */
function getSvelteCompileWarningsWithoutCache(context) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    // Process for styles
    const styleElementsWithNotCSS = [...extractStyleElementsWithLangOtherThanCSS(context)];
    const stripStyleElements = [];
    const transformResults = [];
    for (const style of styleElementsWithNotCSS) {
        const transform = STYLE_TRANSFORMS[style.lang];
        if (transform) {
            const result = transform(style.node, (0, compat_1.getSourceCode)(context).text, context);
            if (result) {
                transformResults.push(result);
                continue;
            }
        }
        stripStyleElements.push(style.node);
    }
    const stripStyleTokens = stripStyleElements.flatMap((e) => e.children);
    const ignoreComments = (0, ignore_comment_1.getSvelteIgnoreItems)(context).filter((item) => item.code != null);
    const text = buildStrippedText(context, ignoreComments, stripStyleTokens);
    transformResults.push(...transformScripts(context, text));
    if (!transformResults.length) {
        const warnings = getWarningsFromCode(text, context);
        return {
            ...processIgnore(warnings.warnings, warnings.kind, stripStyleElements, ignoreComments, context),
            kind: warnings.kind,
            stripStyleElements
        };
    }
    class RemapContext {
        constructor() {
            this.originalStart = 0;
            this.code = '';
            this.locs = null;
            this.mapIndexes = [];
        }
        appendOriginal(endIndex) {
            const codeStart = this.code.length;
            const start = this.originalStart;
            this.code += text.slice(start, endIndex);
            this.originalStart = endIndex;
            const offset = start - codeStart;
            this.mapIndexes.push({
                range: [codeStart, this.code.length],
                remap(index) {
                    return index + offset;
                }
            });
        }
        postprocess() {
            this.appendOriginal(text.length);
            return this.code;
        }
        appendTranspile(output) {
            const endIndex = output.inputRange[1];
            const codeStart = this.code.length;
            const start = this.originalStart;
            const inputText = text.slice(start, endIndex);
            const outputText = `${output.output}\n`;
            this.code += outputText;
            this.originalStart = endIndex;
            let outputLocs = null;
            let inputLocs = null;
            let decoded = null;
            this.mapIndexes.push({
                range: [codeStart, this.code.length],
                remap: (index) => {
                    outputLocs = outputLocs ?? new lines_and_columns_1.LinesAndColumns(outputText);
                    inputLocs = inputLocs ?? new lines_and_columns_1.LinesAndColumns(inputText);
                    const outputCodePos = outputLocs.getLocFromIndex(index - codeStart);
                    const inputCodePos = remapPosition(outputCodePos);
                    return inputLocs.getIndexFromLoc(inputCodePos) + start;
                }
            });
            /** Remapping source position */
            function remapPosition(pos) {
                decoded = decoded ?? (0, sourcemap_codec_1.decode)(output.mappings);
                const lineMaps = decoded[pos.line - 1];
                if (!lineMaps?.length) {
                    for (let line = pos.line - 1; line >= 0; line--) {
                        const prevLineMaps = decoded[line];
                        if (prevLineMaps?.length) {
                            const [, , sourceCodeLine, sourceCodeColumn] = prevLineMaps[prevLineMaps.length - 1];
                            return {
                                line: sourceCodeLine + 1,
                                column: sourceCodeColumn
                            };
                        }
                    }
                    return { line: -1, column: -1 };
                }
                for (let index = 0; index < lineMaps.length - 1; index++) {
                    const [generateCodeColumn, , sourceCodeLine, sourceCodeColumn] = lineMaps[index];
                    if (generateCodeColumn <= pos.column && pos.column < lineMaps[index + 1][0]) {
                        return {
                            line: sourceCodeLine + 1,
                            column: sourceCodeColumn + (pos.column - generateCodeColumn)
                        };
                    }
                }
                const [generateCodeColumn, , sourceCodeLine, sourceCodeColumn] = lineMaps[lineMaps.length - 1];
                return {
                    line: sourceCodeLine + 1,
                    column: sourceCodeColumn + (pos.column - generateCodeColumn)
                };
            }
        }
        remapLocs(points) {
            const mapIndexes = this.mapIndexes;
            const locs = (this.locs = this.locs ?? new lines_and_columns_1.LinesAndColumns(this.code));
            let start = undefined;
            let end = undefined;
            if (points.start) {
                const index = locs.getIndexFromLoc(points.start);
                const remapped = remapIndex(index);
                if (remapped) {
                    start = { ...sourceCode.getLocFromIndex(remapped), character: remapped };
                }
            }
            if (points.end) {
                const index = locs.getIndexFromLoc(points.end);
                const remapped = remapIndex(index - 1 /* include index */);
                if (remapped) {
                    const character = remapped + 1; /* restore */
                    end = { ...sourceCode.getLocFromIndex(character), character };
                }
            }
            return { start, end };
            /** remap index */
            function remapIndex(index) {
                for (const mapIndex of mapIndexes) {
                    if (mapIndex.range[0] <= index && index < mapIndex.range[1]) {
                        return mapIndex.remap(index);
                    }
                }
                return null;
            }
        }
    }
    const remapContext = new RemapContext();
    for (const result of transformResults.sort((a, b) => a.inputRange[0] - b.inputRange[0])) {
        remapContext.appendOriginal(result.inputRange[0]);
        remapContext.appendTranspile(result);
    }
    const code = remapContext.postprocess();
    const baseWarnings = getWarningsFromCode(code, context);
    const warnings = [];
    for (const warn of baseWarnings.warnings) {
        let loc = null;
        /** Get re-mapped location */
        // eslint-disable-next-line func-style -- ignore
        const getLoc = function getLoc() {
            if (loc) {
                return loc;
            }
            return (loc = remapContext.remapLocs(warn));
        };
        warnings.push({
            code: warn.code,
            message: warn.message,
            get start() {
                return getLoc().start;
            },
            get end() {
                return getLoc().end;
            }
        });
    }
    return {
        ...processIgnore(warnings, baseWarnings.kind, stripStyleElements, ignoreComments, context),
        kind: baseWarnings.kind,
        stripStyleElements
    };
}
/**
 * Extracts the style with the lang attribute other than CSS.
 */
function* extractStyleElementsWithLangOtherThanCSS(context) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const root = sourceCode.ast;
    for (const node of root.body) {
        if (node.type === 'SvelteStyleElement') {
            const lang = (0, ast_utils_1.getLangValue)(node);
            if (lang != null && lang.toLowerCase() !== 'css') {
                yield { node, lang: lang.toLowerCase() };
            }
        }
    }
}
/**
 * Build the text stripped of tokens that are not needed for compilation.
 */
function buildStrippedText(context, ignoreComments, stripStyleTokens) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const baseText = sourceCode.text;
    const stripTokens = new Set([...ignoreComments.map((item) => item.token), ...stripStyleTokens]);
    if (!stripTokens.size) {
        return baseText;
    }
    let code = '';
    let start = 0;
    for (const token of [...stripTokens].sort((a, b) => a.range[0] - b.range[0])) {
        code +=
            baseText.slice(start, token.range[0]) +
                baseText.slice(...token.range).replace(/[^\t\n\r ]/g, ' ');
        start = token.range[1];
    }
    code += baseText.slice(start);
    return code;
}
/** Returns the result of transforming the required script for the transform. */
function* transformScripts(context, text) {
    const transform = isUseTypeScript(context)
        ? (0, typescript_1.hasTypeScript)(context)
            ? typescript_1.transform
            : babel_1.transform
        : isUseBabel(context)
            ? babel_1.transform
            : null;
    const sourceCode = (0, compat_1.getSourceCode)(context);
    if (transform) {
        const root = sourceCode.ast;
        for (const node of root.body) {
            if (node.type === 'SvelteScriptElement') {
                const result = transform(node, text, context);
                if (result) {
                    yield result;
                }
            }
        }
    }
}
function hasTagOption(program) {
    return program.body.some((body) => {
        if (body.type !== 'SvelteElement' || body.kind !== 'special') {
            return false;
        }
        if (body.name.name !== 'svelte:options') {
            return false;
        }
        return Boolean((0, ast_utils_1.findAttribute)(body, 'tag'));
    });
}
/**
 * Get compile warnings
 */
function getWarningsFromCode(code, context) {
    try {
        const result = compiler.compile(code, {
            generate: false,
            ...(semver_1.default.satisfies(compiler.VERSION, '>=4.0.0-0')
                ? { customElement: true }
                : hasTagOption((0, compat_1.getSourceCode)(context).ast)
                    ? { customElement: true }
                    : {})
        });
        return { warnings: result.warnings, kind: 'warn' };
        // eslint-disable-next-line @typescript-eslint/no-explicit-any -- ignore
    }
    catch (e) {
        return {
            warnings: [
                {
                    code: e.code,
                    message: e.message,
                    start: e.start,
                    end: e.end
                }
            ],
            kind: 'error'
        };
    }
}
/** Ignore process */
function processIgnore(warnings, kind, stripStyleElements, ignoreComments, context) {
    if (kind === 'error') {
        return {
            warnings,
            unusedIgnores: ignoreComments
        };
    }
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const unusedIgnores = new Set(ignoreComments);
    const remainingWarning = new Set(warnings);
    for (const warning of warnings) {
        if (!warning.code) {
            continue;
        }
        let node = getWarningNode(warning);
        while (node) {
            for (const comment of (0, extract_leading_comments_1.extractLeadingComments)(context, node).reverse()) {
                const ignoreItem = ignoreComments.find((item) => item.token === comment &&
                    (item.code === warning.code || item.codeForV5 === warning.code));
                if (ignoreItem) {
                    unusedIgnores.delete(ignoreItem);
                    remainingWarning.delete(warning);
                    break;
                }
            }
            node = getIgnoreParent(node);
        }
    }
    // Stripped styles are ignored from compilation and cannot determine css errors.
    for (const node of stripStyleElements) {
        for (const comment of (0, extract_leading_comments_1.extractLeadingComments)(context, node).reverse()) {
            const ignoreItem = ignoreComments.find((item) => item.token === comment &&
                (CSS_WARN_CODES.has(item.code) || CSS_WARN_CODES.has(item.codeForV5)));
            if (ignoreItem) {
                unusedIgnores.delete(ignoreItem);
                break;
            }
        }
    }
    return {
        warnings: [...remainingWarning],
        unusedIgnores: [...unusedIgnores]
    };
    /** Get ignore target parent node */
    function getIgnoreParent(node) {
        if (node.type !== 'SvelteElement' &&
            node.type !== 'SvelteIfBlock' &&
            node.type !== 'SvelteKeyBlock' &&
            node.type !== 'SvelteEachBlock' &&
            node.type !== 'SvelteAwaitBlock') {
            return null;
        }
        const parent = node.parent;
        if (parent.type === 'SvelteElseBlock') {
            return parent.parent; // SvelteIfBlock or SvelteEachBlock
        }
        if (parent.type === 'SvelteAwaitPendingBlock' ||
            parent.type === 'SvelteAwaitThenBlock' ||
            parent.type === 'SvelteAwaitCatchBlock') {
            return parent.parent; // SvelteAwaitBlock
        }
        if (parent.type !== 'SvelteElement' &&
            parent.type !== 'SvelteIfBlock' &&
            parent.type !== 'SvelteKeyBlock' &&
            parent.type !== 'SvelteEachBlock'
        // && parent.type !== "SvelteAwaitBlock"
        ) {
            return null;
        }
        return parent;
    }
    /** Get warning node */
    function getWarningNode(warning) {
        const indexes = getWarningIndexes(warning);
        if (indexes.start != null) {
            const node = getWarningTargetNodeFromIndex(indexes.start);
            if (node) {
                return node;
            }
            if (indexes.end != null) {
                const center = Math.floor(indexes.start + (indexes.end - indexes.start) / 2);
                return getWarningTargetNodeFromIndex(center);
            }
        }
        if (indexes.end != null) {
            return getWarningTargetNodeFromIndex(indexes.end);
        }
        return null;
    }
    /**
     * Get warning target node from the given index
     */
    function getWarningTargetNodeFromIndex(index) {
        let targetNode = sourceCode.getNodeByRangeIndex(index);
        while (targetNode) {
            if (targetNode.type === 'SvelteElement' || targetNode.type === 'SvelteStyleElement') {
                return targetNode;
            }
            if (targetNode.parent) {
                if (targetNode.parent.type === 'Program' ||
                    targetNode.parent.type === 'SvelteScriptElement') {
                    return targetNode;
                }
            }
            else {
                return null;
            }
            targetNode = targetNode.parent || null;
        }
        return null;
    }
    /** Get warning index */
    function getWarningIndexes(warning) {
        const start = warning.start && sourceCode.getIndexFromLoc(warning.start);
        const end = warning.end && sourceCode.getIndexFromLoc(warning.end);
        return { start, end };
    }
}
/**
 * Check if using TypeScript.
 */
function isUseTypeScript(context) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    if (sourceCode.parserServices.esTreeNodeToTSNodeMap)
        return true;
    const root = sourceCode.ast;
    for (const node of root.body) {
        if (node.type === 'SvelteScriptElement') {
            const lang = (0, ast_utils_1.getLangValue)(node)?.toLowerCase();
            if (lang === 'ts' || lang === 'typescript') {
                return true;
            }
        }
    }
    return false;
}
/**
 * Check if using Babel.
 */
function isUseBabel(context) {
    const parser = context.parserOptions?.parser;
    if (!parser) {
        return false;
    }
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const root = sourceCode.ast;
    let scriptLang = 'js';
    for (const node of root.body) {
        if (node.type === 'SvelteScriptElement') {
            const lang = (0, ast_utils_1.getLangValue)(node)?.toLowerCase();
            if (lang === 'ts' || lang === 'typescript') {
                scriptLang = lang;
                break;
            }
        }
    }
    const parserName = getParserName(scriptLang, parser);
    if (!parserName) {
        return false;
    }
    if (parserName === '@babel/eslint-parser') {
        return true;
    }
    if (parserName.includes('@babel/eslint-parser')) {
        let targetPath = parserName;
        while (targetPath) {
            const pkgPath = path_1.default.join(targetPath, 'package.json');
            if (fs_1.default.existsSync(pkgPath)) {
                try {
                    return JSON.parse(fs_1.default.readFileSync(pkgPath, 'utf-8'))?.name === '@babel/eslint-parser';
                }
                catch {
                    return false;
                }
            }
            const parent = path_1.default.dirname(targetPath);
            if (targetPath === parent) {
                break;
            }
            targetPath = parent;
        }
    }
    return false;
    /** Get script parser name */
    function getParserName(lang, parser) {
        if (typeof parser === 'string') {
            return parser;
        }
        else if (typeof parser === 'object') {
            const name = parser[lang];
            if (typeof name === 'string') {
                return name;
            }
        }
        return null;
    }
}
