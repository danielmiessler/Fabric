"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseTemplate = parseTemplate;
const compiler_1 = require("svelte/compiler");
const index_1 = require("./converts/index");
const __1 = require("..");
const svelte_version_1 = require("./svelte-version");
/**
 * Parse for template
 */
function parseTemplate(code, ctx, parserOptions) {
    try {
        const options = Object.assign({ filename: parserOptions.filePath }, (svelte_version_1.svelteVersion.gte(5) ? { modern: true } : {}));
        const svelteAst = (0, compiler_1.parse)(code, options);
        const ast = (0, index_1.convertSvelteRoot)(svelteAst, ctx);
        return {
            ast,
            svelteAst,
        };
    }
    catch (e) {
        if (typeof e.pos === "number") {
            const err = new __1.ParseError(e.message, e.pos, ctx);
            err.svelteCompilerError = e;
            throw err;
        }
        throw e;
    }
}
