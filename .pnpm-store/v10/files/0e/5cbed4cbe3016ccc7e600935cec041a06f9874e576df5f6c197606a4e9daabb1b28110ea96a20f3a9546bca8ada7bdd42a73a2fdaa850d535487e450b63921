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
Object.defineProperty(exports, "__esModule", { value: true });
const tagInfo_1 = require("../modules/tagInfo");
const utils_1 = require("../modules/utils");
const prepareContent_1 = require("../modules/prepareContent");
exports.default = (options) => ({
    async script(svelteFile) {
        const { transformer } = await Promise.resolve().then(() => __importStar(require('../transformers/coffeescript')));
        let { content, filename, attributes, lang, dependencies } = await (0, tagInfo_1.getTagInfo)(svelteFile);
        if (lang !== 'coffeescript') {
            return { code: content };
        }
        content = (0, prepareContent_1.prepareContent)({
            options: {
                ...options,
                stripIndent: true,
            },
            content,
        });
        const transformed = await transformer({
            content,
            filename,
            attributes,
            options,
        });
        return {
            ...transformed,
            dependencies: (0, utils_1.concat)(dependencies, transformed.dependencies),
        };
    },
});
