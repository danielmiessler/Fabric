"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transform = transform;
exports.hasBabel = hasBabel;
const load_module_1 = require("../../../utils/load-module");
const compat_1 = require("../../../utils/compat");
/**
 * Transpile with babel
 */
function transform(node, text, context) {
    const babel = loadBabel(context);
    if (!babel) {
        return null;
    }
    let inputRange;
    if (node.endTag) {
        inputRange = [node.startTag.range[1], node.endTag.range[0]];
    }
    else {
        inputRange = [node.startTag.range[1], node.range[1]];
    }
    const code = text.slice(...inputRange);
    try {
        const output = babel.transformSync(code, {
            sourceType: 'module',
            sourceMaps: true,
            minified: false,
            ast: false,
            code: true,
            cwd: (0, compat_1.getCwd)(context)
        });
        if (!output) {
            return null;
        }
        return {
            inputRange,
            output: output.code,
            mappings: output.map.mappings
        };
    }
    catch (_e) {
        return null;
    }
}
/** Check if project has Babel. */
function hasBabel(context) {
    return Boolean(loadBabel(context));
}
/**
 * Load babel
 */
function loadBabel(context) {
    return (0, load_module_1.loadModule)(context, '@babel/core');
}
