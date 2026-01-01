"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getTagInfo = void 0;
/* eslint-disable node/prefer-promises/fs */
const fs_1 = require("fs");
const path_1 = require("path");
const language_1 = require("./language");
const utils_1 = require("./utils");
const resolveSrc = (importerFile, srcPath) => (0, path_1.resolve)((0, path_1.dirname)(importerFile), srcPath);
const getSrcContent = (file) => {
    return new Promise((resolve, reject) => {
        (0, fs_1.readFile)(file, (error, data) => {
            // istanbul ignore if
            if (error)
                reject(error);
            else
                resolve(data.toString());
        });
    });
};
async function doesFileExist(file) {
    return new Promise((resolve) => (0, fs_1.access)(file, 0, (err) => resolve(!err)));
}
const getTagInfo = async ({ attributes, filename, content, markup, }) => {
    const dependencies = [];
    // catches empty content and self-closing tags
    const isEmptyContent = content == null || content.trim().length === 0;
    /** only include src file if content of tag is empty */
    if (attributes.src && isEmptyContent) {
        // istanbul ignore if
        if (typeof attributes.src !== 'string') {
            throw new Error('src attribute must be string');
        }
        let path = attributes.src;
        /** Only try to get local files (path starts with ./ or ../) */
        if ((0, utils_1.isValidLocalPath)(path) && filename) {
            path = resolveSrc(filename, path);
            if (await doesFileExist(path)) {
                content = await getSrcContent(path);
                dependencies.push(path);
            }
            else {
                console.warn(`[svelte-preprocess] The file  "${path}" was not found.`);
            }
        }
    }
    const { lang, alias } = (0, language_1.getLanguage)(attributes);
    return {
        filename,
        attributes,
        content,
        lang,
        alias,
        dependencies,
        markup,
    };
};
exports.getTagInfo = getTagInfo;
