"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertRenderTag = convertRenderTag;
const common_1 = require("./common");
/** Convert for RenderTag */
function convertRenderTag(node, parent, ctx) {
    const mustache = Object.assign({ type: "SvelteRenderTag", expression: null, parent }, ctx.getConvertLocation(node));
    const callRange = (0, common_1.getWithLoc)(node.expression);
    ctx.scriptLet.addExpressionFromRange([callRange.start, callRange.end], mustache, null, (expression) => {
        mustache.expression = expression;
        mustache.expression.parent = mustache;
    });
    const atRenderStart = ctx.code.indexOf("@render", mustache.range[0]);
    ctx.addToken("MustacheKeyword", {
        start: atRenderStart,
        end: atRenderStart + 7,
    });
    return mustache;
}
