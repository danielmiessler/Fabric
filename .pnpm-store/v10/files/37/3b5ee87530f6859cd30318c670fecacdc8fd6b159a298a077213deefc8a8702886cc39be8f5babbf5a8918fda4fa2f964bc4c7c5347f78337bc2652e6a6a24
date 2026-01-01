"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transform = transform;
const postcss_1 = __importDefault(require("postcss"));
const postcss_load_config_1 = __importDefault(require("postcss-load-config"));
const compat_1 = require("../../../utils/compat");
/**
 * Transform with postcss
 */
function transform(node, text, context) {
    const postcssConfig = context.settings?.svelte?.compileOptions?.postcss;
    if (postcssConfig === false) {
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
    const filename = `${(0, compat_1.getFilename)(context)}.css`;
    try {
        const configFilePath = postcssConfig?.configFilePath;
        const config = postcss_load_config_1.default.sync({
            cwd: (0, compat_1.getCwd)(context),
            from: filename
        }, typeof configFilePath === 'string' ? configFilePath : undefined);
        const result = (0, postcss_1.default)(config.plugins).process(code, {
            ...config.options,
            map: {
                inline: false
            }
        });
        return {
            inputRange,
            output: result.content,
            mappings: result.map.toJSON().mappings
        };
    }
    catch (_e) {
        // console.log(e)
        return null;
    }
}
