"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.parseTypeScriptInSvelte = parseTypeScriptInSvelte;
exports.parseTypeScript = parseTypeScript;
const script_1 = require("../script");
const analyze_1 = require("./analyze");
const set_parent_1 = require("./set-parent");
/**
 * Parse for TypeScript in <script>
 */
function parseTypeScriptInSvelte(code, attrs, parserOptions, context) {
    const tsCtx = (0, analyze_1.analyzeTypeScriptInSvelte)(code, attrs, parserOptions, context);
    const result = (0, script_1.parseScriptInSvelte)(tsCtx.script, attrs, parserOptions);
    tsCtx.restoreContext.restore(result);
    return result;
}
/**
 * Parse for TypeScript
 */
function parseTypeScript(code, attrs, parserOptions, svelteParseContext) {
    const tsCtx = (0, analyze_1.analyzeTypeScript)(code, attrs, parserOptions, svelteParseContext);
    const result = (0, script_1.parseScript)(tsCtx.script, attrs, parserOptions);
    (0, set_parent_1.setParent)(result);
    tsCtx.restoreContext.restore(result);
    return result;
}
