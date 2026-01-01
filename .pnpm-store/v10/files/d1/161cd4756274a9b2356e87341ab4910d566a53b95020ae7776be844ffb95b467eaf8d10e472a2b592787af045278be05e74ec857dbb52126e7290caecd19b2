"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const path_1 = __importDefault(require("path"));
const stylus_1 = __importDefault(require("stylus"));
const utils_1 = require("../modules/utils");
const transformer = ({ content, filename, options = {}, }) => {
    options = {
        paths: (0, utils_1.getIncludePaths)(filename, options.paths),
        ...options,
    };
    return new Promise((resolve, reject) => {
        const style = (0, stylus_1.default)(content, {
            filename,
            ...options,
        }).set('sourcemap', options.sourcemap);
        style.render((err, css) => {
            var _a;
            // istanbul ignore next
            if (err)
                reject(err);
            if ((_a = style.sourcemap) === null || _a === void 0 ? void 0 : _a.sources) {
                style.sourcemap.sources = style.sourcemap.sources.map((source) => path_1.default.resolve(source));
            }
            resolve({
                code: css,
                map: style.sourcemap,
                // .map() necessary for windows compatibility
                dependencies: style
                    .deps(filename)
                    .map((filePath) => path_1.default.resolve(filePath)),
            });
        });
    });
};
exports.transformer = transformer;
