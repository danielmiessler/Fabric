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
const postcss_1 = __importDefault(require("postcss"));
async function process({ options: { plugins = [], parser, syntax } = {}, content, filename, sourceMap, }) {
    const { css, map, messages } = await (0, postcss_1.default)(plugins).process(content, {
        from: filename,
        to: filename,
        map: { prev: sourceMap, inline: false },
        parser,
        syntax,
    });
    const dependencies = messages.reduce((acc, msg) => {
        // istanbul ignore if
        if (msg.type !== 'dependency')
            return acc;
        acc.push(msg.file);
        return acc;
    }, []);
    return { code: css, map, dependencies };
}
async function getConfigFromFile(options) {
    try {
        /** If not, look for a postcss config file */
        const { default: postcssLoadConfig } = await Promise.resolve().then(() => __importStar(require(`postcss-load-config`)));
        const loadedConfig = await postcssLoadConfig(options, options === null || options === void 0 ? void 0 : options.configFilePath);
        return {
            error: null,
            config: {
                plugins: loadedConfig.plugins,
                // `postcss-load-config` puts all other props in a `options` object
                ...loadedConfig.options,
            },
        };
    }
    catch (e) {
        return {
            config: null,
            error: e,
        };
    }
}
/** Adapted from https://github.com/TehShrike/svelte-preprocess-postcss */
const transformer = async ({ content, filename, options = {}, map, }) => {
    let fileConfig = null;
    if (!options.plugins) {
        fileConfig = await getConfigFromFile(options);
        options = { ...options, ...fileConfig.config };
    }
    if (options.plugins || options.syntax || options.parser) {
        return process({ options, content, filename, sourceMap: map });
    }
    if ((fileConfig === null || fileConfig === void 0 ? void 0 : fileConfig.error) != null) {
        console.error(`[svelte-preprocess] PostCSS configuration was not passed or is invalid. If you expect to load it from a file make sure to install "postcss-load-config" and try again.\n\n${fileConfig.error}`);
    }
    return { code: content, map, dependencies: [] };
};
exports.transformer = transformer;
