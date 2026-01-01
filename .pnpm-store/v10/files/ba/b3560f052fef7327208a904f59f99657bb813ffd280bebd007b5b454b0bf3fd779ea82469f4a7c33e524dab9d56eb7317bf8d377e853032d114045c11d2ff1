"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.convertConstTag = convertConstTag;
const compat_1 = require("../compat");
/** Convert for ConstTag */
function convertConstTag(node, parent, ctx) {
    const mustache = Object.assign({ type: "SvelteConstTag", declaration: null, parent }, ctx.getConvertLocation(node));
    ctx.scriptLet.addVariableDeclarator((0, compat_1.getDeclaratorFromConstTag)(node), mustache, (declaration) => {
        mustache.declaration = declaration;
    });
    const atConstStart = ctx.code.indexOf("@const", mustache.range[0]);
    ctx.addToken("MustacheKeyword", {
        start: atConstStart,
        end: atConstStart + 6,
    });
    return mustache;
}
