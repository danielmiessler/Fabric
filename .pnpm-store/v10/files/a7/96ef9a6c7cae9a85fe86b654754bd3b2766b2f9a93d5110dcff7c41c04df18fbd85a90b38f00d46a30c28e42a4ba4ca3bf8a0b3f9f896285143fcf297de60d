"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const transformer = async ({ content, options, }) => {
    let newContent = content;
    if (options == null) {
        return { code: content };
    }
    for (const [regex, replacer] of options) {
        newContent = newContent.replace(regex, replacer);
    }
    return {
        code: newContent,
    };
};
exports.transformer = transformer;
