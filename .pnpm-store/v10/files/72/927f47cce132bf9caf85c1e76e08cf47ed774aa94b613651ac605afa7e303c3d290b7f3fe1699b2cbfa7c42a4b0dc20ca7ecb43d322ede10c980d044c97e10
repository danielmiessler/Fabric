"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertText = convertText;
exports.convertTextToLiteral = convertTextToLiteral;
/** Convert for Text */
function convertText(node, parent, ctx) {
    const text = Object.assign({ type: "SvelteText", value: node.data, parent }, ctx.getConvertLocation(node));
    extractTextTokens(node, ctx);
    return text;
}
/** Convert for Text to Literal */
function convertTextToLiteral(node, parent, ctx) {
    const text = Object.assign({ type: "SvelteLiteral", value: node.data, parent }, ctx.getConvertLocation(node));
    extractTextTokens(node, ctx);
    return text;
}
/** Extract tokens */
function extractTextTokens(node, ctx) {
    const loc = node;
    let start = loc.start;
    let word = false;
    for (let index = loc.start; index < loc.end; index++) {
        if (word !== Boolean(ctx.code[index].trim())) {
            if (start < index) {
                ctx.addToken("HTMLText", { start, end: index });
            }
            word = !word;
            start = index;
        }
    }
    if (start < loc.end) {
        ctx.addToken("HTMLText", { start, end: loc.end });
    }
}
