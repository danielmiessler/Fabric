"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseStyleAttributeValue = parseStyleAttributeValue;
const template_safe_parser_1 = __importDefault(require("./template-safe-parser"));
const postcss_1 = require("postcss");
const compat_1 = require("../compat");
/** Parse for CSS */
function safeParseCss(css) {
    try {
        const input = new postcss_1.Input(css);
        const parser = new template_safe_parser_1.default(input);
        parser.parse();
        return parser.root;
    }
    catch {
        return null;
    }
}
const cache = new WeakMap();
/**
 * Parse style attribute value
 */
function parseStyleAttributeValue(node, context) {
    if (cache.has(node)) {
        return cache.get(node) || null;
    }
    cache.set(node, null);
    if (!node.value?.length) {
        return null;
    }
    const startOffset = node.value[0].range[0];
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const cssCode = node.value.map((value) => sourceCode.getText(value)).join('');
    const root = safeParseCss(cssCode);
    if (!root) {
        return root;
    }
    const ctx = {
        startOffset,
        value: node.value,
        context
    };
    const mustacheTags = node.value.filter((v) => v.type === 'SvelteMustacheTag');
    const converted = convertRoot(root, mustacheTags, (e) => e.range, ctx);
    cache.set(node, converted);
    return converted;
}
class IgnoreError extends Error {
}
/** Checks wether the given node is string literal or not  */
function isStringLiteral(node) {
    return node.type === 'Literal' && typeof node.value === 'string';
}
/** convert root node */
function convertRoot(root, interpolations, getRange, ctx) {
    const nodes = [];
    for (const child of root.nodes) {
        const converted = convertChild(child, ctx);
        if (!converted) {
            return null;
        }
        while (interpolations[0]) {
            const tagOrExpr = interpolations[0];
            if (tagOrExpr.range[1] <= converted.range[0]) {
                nodes.push(buildSvelteStyleInline(tagOrExpr));
                interpolations.shift();
                continue;
            }
            if (tagOrExpr.range[0] < converted.range[1]) {
                try {
                    converted.addInterpolation(tagOrExpr);
                }
                catch (e) {
                    if (e instanceof IgnoreError)
                        return null;
                    throw e;
                }
                interpolations.shift();
                continue;
            }
            break;
        }
        nodes.push(converted);
    }
    nodes.push(...interpolations.map(buildSvelteStyleInline));
    return {
        type: 'root',
        nodes
    };
    /** Build SvelteStyleInline */
    function buildSvelteStyleInline(tagOrExpr) {
        const inlineStyles = new Map();
        let range = null;
        /** Get range */
        function getRangeForInline() {
            if (range) {
                return range;
            }
            return range ?? (range = getRange(tagOrExpr));
        }
        return {
            type: 'inline',
            node: tagOrExpr,
            get range() {
                return getRangeForInline();
            },
            get loc() {
                return toLoc(getRangeForInline(), ctx);
            },
            getInlineStyle(node) {
                return getInlineStyle(node);
            },
            getAllInlineStyles() {
                const allInlineStyles = new Map();
                for (const node of extractExpressions(tagOrExpr)) {
                    const style = getInlineStyle(node);
                    if (style) {
                        allInlineStyles.set(node, style);
                    }
                }
                return allInlineStyles;
            }
        };
        /** Get inline style node */
        function getInlineStyle(node) {
            if (node.type === 'SvelteMustacheTag') {
                return getInlineStyle(node.expression);
            }
            if (inlineStyles.has(node)) {
                return inlineStyles.get(node) || null;
            }
            const sourceCode = (0, compat_1.getSourceCode)(ctx.context);
            inlineStyles.set(node, null);
            let converted;
            if (isStringLiteral(node)) {
                const root = safeParseCss(sourceCode.getText(node).slice(1, -1));
                if (!root) {
                    return null;
                }
                converted = convertRoot(root, [], () => [0, 0], {
                    ...ctx,
                    startOffset: node.range[0] + 1
                });
            }
            else if (node.type === 'TemplateLiteral') {
                const root = safeParseCss(sourceCode.getText(node).slice(1, -1));
                if (!root) {
                    return null;
                }
                converted = convertRoot(root, [...node.expressions], (e) => {
                    const index = node.expressions.indexOf(e);
                    return [node.quasis[index].range[1] - 2, node.quasis[index + 1].range[0] + 1];
                }, {
                    ...ctx,
                    startOffset: node.range[0] + 1
                });
            }
            else {
                return null;
            }
            inlineStyles.set(node, converted);
            return converted;
        }
        /** Extract all expressions */
        function* extractExpressions(node) {
            if (node.type === 'SvelteMustacheTag') {
                yield* extractExpressions(node.expression);
            }
            else if (isStringLiteral(node)) {
                yield node;
            }
            else if (node.type === 'TemplateLiteral') {
                yield node;
            }
            else if (node.type === 'ConditionalExpression') {
                yield* extractExpressions(node.consequent);
                yield* extractExpressions(node.alternate);
            }
            else if (node.type === 'LogicalExpression') {
                yield* extractExpressions(node.left);
                yield* extractExpressions(node.right);
            }
        }
    }
}
/** convert child node */
function convertChild(node, ctx) {
    const range = convertRange(node, ctx);
    if (node.type === 'decl') {
        const propRange = [range[0], range[0] + node.prop.length];
        const declValueStartIndex = propRange[1] + (node.raws.between || '').length;
        const valueRange = [
            declValueStartIndex,
            declValueStartIndex + (node.raws.value?.value || node.value).length
        ];
        const prop = {
            name: node.prop,
            range: propRange,
            get loc() {
                return toLoc(propRange, ctx);
            },
            interpolations: []
        };
        const value = {
            value: node.value,
            range: valueRange,
            get loc() {
                return toLoc(valueRange, ctx);
            },
            interpolations: []
        };
        const unknownInterpolations = [];
        return {
            type: 'decl',
            prop,
            value,
            important: node.important,
            range,
            get loc() {
                return toLoc(range, ctx);
            },
            addInterpolation(tagOrExpr) {
                const index = tagOrExpr.range[0];
                if (prop.range[0] <= index && index < prop.range[1]) {
                    prop.interpolations.push(tagOrExpr);
                    return;
                }
                if (value.range[0] <= index && index < value.range[1]) {
                    value.interpolations.push(tagOrExpr);
                    return;
                }
                unknownInterpolations.push(tagOrExpr);
            },
            unknownInterpolations
        };
    }
    if (node.type === 'comment') {
        return {
            type: 'comment',
            range,
            get loc() {
                return toLoc(range, ctx);
            },
            addInterpolation: () => {
                throw new IgnoreError();
            }
        };
    }
    if (node.type === 'atrule') {
        // unexpected node
        return null;
    }
    if (node.type === 'rule') {
        // unexpected node
        return null;
    }
    // unknown node
    return null;
}
/** convert range */
function convertRange(node, ctx) {
    return [ctx.startOffset + node.source.start.offset, ctx.startOffset + node.source.end.offset];
}
/** convert range */
function toLoc(range, ctx) {
    return {
        start: (0, compat_1.getSourceCode)(ctx.context).getLocFromIndex(range[0]),
        end: (0, compat_1.getSourceCode)(ctx.context).getLocFromIndex(range[1])
    };
}
