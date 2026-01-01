"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const coffeescript_1 = __importDefault(require("coffeescript"));
const transformer = ({ content, filename, options, }) => {
    const coffeeOptions = {
        filename,
        /*
         * Since `coffeescript` transpiles variables to `var` definitions, it uses a safety mechanism to prevent variables from bleeding to outside contexts. This mechanism consists of wrapping your `coffeescript` code inside an IIFE which, unfortunately, prevents `svelte` from finding your variables. To bypass this behavior, `svelte-preprocess` sets the [`bare` coffeescript compiler option](https://coffeescript.org/#lexical-scope) to `true`.
         */
        bare: true,
        ...options,
    };
    if (coffeeOptions.sourceMap) {
        const { js: code, v3SourceMap } = coffeescript_1.default.compile(content, coffeeOptions);
        const map = JSON.parse(v3SourceMap);
        return { code, map };
    }
    return { code: coffeescript_1.default.compile(content, coffeeOptions) };
};
exports.transformer = transformer;
