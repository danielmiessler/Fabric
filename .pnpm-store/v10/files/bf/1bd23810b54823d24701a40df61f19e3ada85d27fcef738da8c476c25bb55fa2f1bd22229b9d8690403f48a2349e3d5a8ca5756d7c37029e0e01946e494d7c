"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transform = transform;
const load_module_1 = require("../../../utils/load-module");
const compat_1 = require("../../../utils/compat");
/**
 * Transpile with less
 */
function transform(node, text, context) {
    const less = loadLess(context);
    if (!less) {
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
    const filename = `${(0, compat_1.getFilename)(context)}.less`;
    try {
        let output;
        less.render(code, {
            sourceMap: {},
            syncImport: true,
            filename,
            lint: false
        }, (_error, result) => {
            output = result;
        });
        if (!output) {
            return null;
        }
        return {
            inputRange,
            output: output.css,
            mappings: JSON.parse(output.map).mappings
        };
    }
    catch (_e) {
        return null;
    }
}
/**
 * Load less
 */
function loadLess(context) {
    return (0, load_module_1.loadModule)(context, 'less');
}
