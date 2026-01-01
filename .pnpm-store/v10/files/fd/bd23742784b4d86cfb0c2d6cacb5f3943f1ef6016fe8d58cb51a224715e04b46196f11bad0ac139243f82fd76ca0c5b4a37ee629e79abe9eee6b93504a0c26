"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertMustacheTag = convertMustacheTag;
exports.convertRawMustacheTag = convertRawMustacheTag;
exports.convertDebugTag = convertDebugTag;
const utils_1 = require("../../utils");
/** Convert for MustacheTag */
function convertMustacheTag(node, parent, typing, ctx) {
    return convertMustacheTag0(node, "text", parent, typing, ctx);
}
/** Convert for MustacheTag */
function convertRawMustacheTag(node, parent, ctx) {
    const mustache = convertMustacheTag0(node, "raw", parent, null, ctx);
    const atHtmlStart = ctx.code.indexOf("@html", mustache.range[0]);
    ctx.addToken("MustacheKeyword", {
        start: atHtmlStart,
        end: atHtmlStart + 5,
    });
    return mustache;
}
/** Convert for DebugTag */
function convertDebugTag(node, parent, ctx) {
    const mustache = Object.assign({ type: "SvelteDebugTag", identifiers: [], parent }, ctx.getConvertLocation(node));
    for (const id of node.identifiers) {
        ctx.scriptLet.addExpression(id, mustache, null, (es) => {
            mustache.identifiers.push(es);
        });
    }
    const atDebugStart = ctx.code.indexOf("@debug", mustache.range[0]);
    ctx.addToken("MustacheKeyword", {
        start: atDebugStart,
        end: atDebugStart + 6,
    });
    return mustache;
}
/** Convert to MustacheTag */
function convertMustacheTag0(node, kind, parent, typing, ctx) {
    const mustache = Object.assign({ type: "SvelteMustacheTag", kind, expression: null, parent }, ctx.getConvertLocation(node));
    ctx.scriptLet.addExpression(node.expression, mustache, (0, utils_1.hasTypeInfo)(node.expression) ? null : typing, (es) => {
        mustache.expression = es;
    });
    return mustache;
}
