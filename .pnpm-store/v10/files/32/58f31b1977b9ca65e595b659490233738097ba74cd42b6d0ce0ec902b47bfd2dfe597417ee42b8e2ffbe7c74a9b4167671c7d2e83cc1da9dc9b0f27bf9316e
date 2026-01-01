"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const path_1 = require("path");
const less_1 = __importDefault(require("less"));
const utils_1 = require("../modules/utils");
const transformer = async ({ content, filename, options = {}, }) => {
    options = {
        paths: (0, utils_1.getIncludePaths)(filename, options.paths),
        ...options,
    };
    const { css, map, imports } = await less_1.default.render(content, {
        sourceMap: {},
        filename,
        ...options,
    });
    const dependencies = imports.map((path) => (0, path_1.isAbsolute)(path) ? path : (0, path_1.join)(process.cwd(), path));
    return {
        code: css,
        map,
        dependencies,
    };
};
exports.transformer = transformer;
