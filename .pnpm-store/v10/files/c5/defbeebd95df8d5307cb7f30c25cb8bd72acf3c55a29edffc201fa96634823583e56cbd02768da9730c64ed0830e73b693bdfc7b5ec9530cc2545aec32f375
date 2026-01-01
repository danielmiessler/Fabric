"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const fs_1 = require("fs");
const path_1 = __importDefault(require("path"));
const utils_1 = require("../modules/utils");
const tildeImporter = (url, prev) => {
    if (!url.startsWith('~')) {
        return null;
    }
    // not sure why this ends up here, but let's remove it
    prev = prev.replace('http://localhost', '');
    // on windows, path comes encoded
    if (process.platform === 'win32') {
        prev = decodeURIComponent(prev);
    }
    const modulePath = path_1.default.join('node_modules', ...url.slice(1).split(/[\\/]/g));
    const foundPath = (0, utils_1.findUp)({ what: modulePath, from: prev });
    // istanbul ignore if
    if (foundPath == null) {
        return null;
    }
    const contents = (0, fs_1.readFileSync)(foundPath).toString();
    return { contents };
};
const transformer = async ({ content, filename, options = {}, }) => {
    var _a;
    const { renderSync } = await Promise.resolve().then(() => __importStar(require('sass')));
    const { prependData, ...restOptions } = options;
    const sassOptions = {
        ...restOptions,
        includePaths: (0, utils_1.getIncludePaths)(filename, options.includePaths),
        sourceMap: true,
        sourceMapEmbed: false,
        omitSourceMapUrl: true,
        outFile: `${filename}.css`,
        outputStyle: 'expanded',
        file: filename,
        data: content,
    };
    if (Array.isArray(sassOptions.importer)) {
        sassOptions.importer = [tildeImporter, ...sassOptions.importer];
    }
    else if (sassOptions.importer == null) {
        sassOptions.importer = [tildeImporter];
    }
    else {
        sassOptions.importer = [sassOptions.importer, tildeImporter];
    }
    // scss errors if passed an empty string
    if (content.length === 0) {
        return { code: '' };
    }
    const compiled = renderSync(sassOptions);
    // We need to normalize the path for windows, because the sass compiler
    // returns a windows path in posix format __just for the entry__ (the dependency list below is fine 🤷)
    // More info: https://github.com/sveltejs/svelte-preprocess/issues/619
    const normalizedEntryPath = process.platform === 'win32'
        ? compiled.stats.entry.split('/').join(path_1.default.win32.sep)
        : compiled.stats.entry;
    // For some reason, scss includes the main 'file' in the array, we don't want that
    // Unfortunately I didn't manage to reproduce this in the test env
    // More info: https://github.com/sveltejs/svelte-preprocess/issues/346
    const absoluteEntryPath = path_1.default.isAbsolute(normalizedEntryPath)
        ? normalizedEntryPath
        : path_1.default.join(process.cwd(), normalizedEntryPath);
    const processed = {
        code: compiled.css.toString(),
        map: (_a = compiled.map) === null || _a === void 0 ? void 0 : _a.toString(),
        dependencies: Array.from(compiled.stats.includedFiles).filter((filepath) => filepath !== absoluteEntryPath),
    };
    return processed;
};
exports.transformer = transformer;
