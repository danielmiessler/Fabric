"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transform = transform;
const load_module_1 = require("../../../utils/load-module");
const compat_1 = require("../../../utils/compat");
/**
 * Transpile with stylus
 */
function transform(node, text, context) {
    const stylus = loadStylus(context);
    if (!stylus) {
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
    const filename = `${(0, compat_1.getFilename)(context)}.stylus`;
    try {
        let output;
        const style = stylus(code, {
            filename
        }).set('sourcemap', {});
        style.render((_error, code) => {
            output = code;
        });
        if (output == null) {
            return null;
        }
        return {
            inputRange,
            output,
            mappings: style.sourcemap.mappings
        };
    }
    catch (_e) {
        return null;
    }
}
/**
 * Load stylus
 */
function loadStylus(context) {
    return (0, load_module_1.loadModule)(context, 'stylus');
}
