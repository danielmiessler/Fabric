"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = void 0;
const core_1 = require("@babel/core");
const transformer = async ({ content, filename, options, map = undefined, }) => {
    const babelOptions = {
        ...options,
        inputSourceMap: typeof map === 'string' ? JSON.parse(map) : map !== null && map !== void 0 ? map : undefined,
        sourceType: 'module',
        // istanbul ignore next
        sourceMaps: !!(options === null || options === void 0 ? void 0 : options.sourceMaps),
        filename,
        minified: false,
        ast: false,
        code: true,
        caller: {
            name: 'svelte-preprocess',
            supportsStaticESM: true,
            supportsDynamicImport: true,
            // this isn't supported by Svelte but let it error with a good error on this syntax untouched
            supportsTopLevelAwait: true,
            // todo: this can be enabled once all "peer deps" understand this
            // this syntax is supported since rollup@1.26.0 and webpack@5.0.0-beta.21
            // supportsExportNamespaceFrom: true,
            ...options === null || options === void 0 ? void 0 : options.caller,
        },
    };
    const result = await (0, core_1.transformAsync)(content, babelOptions);
    if (result == null) {
        return { code: content };
    }
    const { code, map: sourcemap } = result;
    return {
        code: code,
        map: sourcemap !== null && sourcemap !== void 0 ? sourcemap : undefined,
    };
};
exports.transformer = transformer;
