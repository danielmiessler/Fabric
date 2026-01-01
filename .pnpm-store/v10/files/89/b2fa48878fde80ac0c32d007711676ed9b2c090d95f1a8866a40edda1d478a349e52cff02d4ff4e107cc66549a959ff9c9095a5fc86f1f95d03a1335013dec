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
exports.sveltePreprocess = exports.transform = void 0;
const utils_1 = require("./modules/utils");
const tagInfo_1 = require("./modules/tagInfo");
const language_1 = require("./modules/language");
const prepareContent_1 = require("./modules/prepareContent");
const markup_1 = require("./modules/markup");
const TARGET_LANGUAGES = Object.freeze({
    markup: 'html',
    style: 'css',
    script: 'javascript',
});
const transform = async (name, options, { content, markup, map, filename, attributes }) => {
    if (name == null || options === false) {
        return { code: content };
    }
    if (typeof options === 'function') {
        return options({ content, map, filename, attributes });
    }
    // todo: maybe add a try-catch here looking for module-not-found errors
    const { transformer } = await Promise.resolve(`${`./transformers/${name}`}`).then(s => __importStar(require(s)));
    return transformer({
        content,
        markup,
        filename,
        map,
        attributes,
        options: typeof options === 'boolean' ? null : options,
    });
};
exports.transform = transform;
function sveltePreprocess(_a) {
    var _b, _c;
    var { aliases, markupTagName = 'template', preserve = [], sourceMap = (_c = ((_b = process === null || process === void 0 ? void 0 : process.env) === null || _b === void 0 ? void 0 : _b.NODE_ENV) === 'development') !== null && _c !== void 0 ? _c : false, ...rest } = _a === void 0 ? {} : _a;
    const transformers = rest;
    if (aliases === null || aliases === void 0 ? void 0 : aliases.length) {
        (0, language_1.addLanguageAlias)(aliases);
    }
    function resolveLanguageArgs(lang, alias) {
        const langOpts = transformers[lang];
        const aliasOpts = alias ? transformers[alias] : undefined;
        const opts = {};
        if (typeof langOpts === 'object') {
            Object.assign(opts, langOpts);
        }
        Object.assign(opts, (0, language_1.getLanguageDefaults)(lang), (0, language_1.getLanguageDefaults)(alias));
        if (lang !== alias && typeof aliasOpts === 'object') {
            Object.assign(opts, aliasOpts);
        }
        if (sourceMap && lang in language_1.SOURCE_MAP_PROP_MAP) {
            const [path, value] = language_1.SOURCE_MAP_PROP_MAP[lang];
            (0, utils_1.setProp)(opts, path, value);
        }
        return opts;
    }
    function getTransformerOptions(lang, alias, { ignoreAliasOverride } = {}) {
        if (lang == null)
            return null;
        const langOpts = transformers[lang];
        const aliasOpts = alias ? transformers[alias] : undefined;
        if (!ignoreAliasOverride && typeof aliasOpts === 'function') {
            return aliasOpts;
        }
        if (typeof langOpts === 'function')
            return langOpts;
        if (aliasOpts === false || langOpts === false)
            return false;
        return resolveLanguageArgs(lang, alias);
    }
    const getTransformerTo = (type, targetLanguage) => async (svelteFile) => {
        let { content, markup, filename, lang, alias, dependencies, attributes } = await (0, tagInfo_1.getTagInfo)(svelteFile);
        if (lang == null || alias == null) {
            alias = TARGET_LANGUAGES[type];
            lang = (0, language_1.getLanguageFromAlias)(alias);
        }
        if ((lang && preserve.includes(lang)) || preserve.includes(alias)) {
            return { code: content };
        }
        const transformerOptions = getTransformerOptions(lang, alias);
        content = (0, prepareContent_1.prepareContent)({
            options: transformerOptions,
            content,
        });
        if (lang === targetLanguage) {
            // has override method for alias
            // example: sugarss override should work apart from postcss
            if (typeof transformerOptions === 'function' && alias !== lang) {
                return transformerOptions({ content, filename, attributes });
            }
            // otherwise, we're done here
            return { code: content, dependencies };
        }
        const transformed = await (0, exports.transform)(lang, transformerOptions, {
            content,
            markup,
            filename,
            attributes,
        });
        return {
            ...transformed,
            dependencies: (0, utils_1.concat)(dependencies, transformed.dependencies),
        };
    };
    const scriptTransformer = getTransformerTo('script', 'javascript');
    const cssTransformer = getTransformerTo('style', 'css');
    const markupTransformer = getTransformerTo('markup', 'html');
    const markup = async ({ content, filename }) => {
        if (transformers.replace) {
            const transformed = await (0, exports.transform)('replace', transformers.replace, {
                content,
                markup: content,
                filename,
            });
            content = transformed.code;
        }
        return (0, markup_1.transformMarkup)({ content, filename }, markupTransformer, {
            // we only pass the markupTagName because the rest of options
            // is fetched internally by the `markupTransformer`
            markupTagName,
        });
    };
    const script = async ({ content, attributes, markup: fullMarkup, filename, }) => {
        const transformResult = await scriptTransformer({
            content,
            attributes,
            markup: fullMarkup,
            filename,
        });
        let { code, map, dependencies, diagnostics } = transformResult;
        if (transformers.babel) {
            const transformed = await (0, exports.transform)('babel', getTransformerOptions('babel'), { content: code, markup: fullMarkup, map, filename, attributes });
            code = transformed.code;
            map = transformed.map;
            dependencies = (0, utils_1.concat)(dependencies, transformed.dependencies);
            diagnostics = (0, utils_1.concat)(diagnostics, transformed.diagnostics);
        }
        return { code, map, dependencies, diagnostics };
    };
    const style = async ({ content, attributes, markup: fullMarkup, filename, }) => {
        const transformResult = await cssTransformer({
            content,
            attributes,
            markup: fullMarkup,
            filename,
        });
        let { code, map, dependencies } = transformResult;
        const hasPostcss = await (0, utils_1.hasDepInstalled)('postcss');
        // istanbul ignore else
        if (hasPostcss) {
            if (transformers.postcss) {
                const { alias, lang } = (0, language_1.getLanguage)(attributes);
                const postcssOptions = getTransformerOptions('postcss', (0, language_1.isAliasOf)(alias, lang) ? alias : null, 
                // todo: this seems wrong and ugly
                { ignoreAliasOverride: true });
                const transformed = await (0, exports.transform)('postcss', postcssOptions, {
                    content: code,
                    markup: fullMarkup,
                    map,
                    filename,
                    attributes,
                });
                code = transformed.code;
                map = transformed.map;
                dependencies = (0, utils_1.concat)(dependencies, transformed.dependencies);
            }
            const transformed = await (0, exports.transform)('globalStyle', getTransformerOptions('globalStyle'), { content: code, markup: fullMarkup, map, filename, attributes });
            code = transformed.code;
            map = transformed.map;
        }
        else if ('global' in attributes) {
            console.warn(`[svelte-preprocess] 'global' attribute found, but 'postcss' is not installed. 'postcss' is used to walk through the CSS and transform any necessary selector.`);
        }
        return { code, map, dependencies };
    };
    return {
        markup,
        script,
        style,
    };
}
exports.sveltePreprocess = sveltePreprocess;
