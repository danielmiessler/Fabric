"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformer = exports.loadTsconfig = void 0;
const path_1 = require("path");
const typescript_1 = __importDefault(require("typescript"));
const compiler_1 = require("svelte/compiler");
const magic_string_1 = __importDefault(require("magic-string"));
const sorcery_1 = require("sorcery");
const errors_1 = require("../modules/errors");
const markup_1 = require("../modules/markup");
const utils_1 = require("../modules/utils");
const package_json_1 = __importDefault(require("svelte/package.json"));
const injectedCodeSeparator = 'const $$$$$$$$ = null;';
/**
 * Map of valid tsconfigs (no errors). Key is the path.
 */
const tsconfigMap = new Map();
const importsNotUsedAsValuesDeprecated = parseInt(typescript_1.default.version.split('.')[0], 10) >= 5;
function createFormatDiagnosticsHost(cwd) {
    return {
        getCanonicalFileName: (fileName) => fileName.replace('.injected.ts', ''),
        getCurrentDirectory: () => cwd,
        getNewLine: () => typescript_1.default.sys.newLine,
    };
}
function formatDiagnostics(diagnostics, basePath) {
    if (Array.isArray(diagnostics)) {
        return typescript_1.default.formatDiagnosticsWithColorAndContext(diagnostics, createFormatDiagnosticsHost(basePath));
    }
    return typescript_1.default.formatDiagnostic(diagnostics, createFormatDiagnosticsHost(basePath));
}
const importTransformer = (context) => {
    const visit = (node) => {
        var _a, _b, _c;
        if (typescript_1.default.isImportDeclaration(node)) {
            if ((_a = node.importClause) === null || _a === void 0 ? void 0 : _a.isTypeOnly) {
                if (!((_b = typescript_1.default.factory) === null || _b === void 0 ? void 0 : _b.createEmptyStatement)) {
                    // @ts-expect-error removed in TS 5.0
                    return typescript_1.default.createEmptyStatement();
                }
                return typescript_1.default.factory.createEmptyStatement();
            }
            if (!((_c = typescript_1.default.factory) === null || _c === void 0 ? void 0 : _c.createImportDeclaration)) {
                // @ts-expect-error removed in TS 5.0
                return typescript_1.default.createImportDeclaration(
                // @ts-expect-error removed in TS 5.0
                node.decorators, node.modifiers, node.importClause, node.moduleSpecifier);
            }
            return typescript_1.default.factory.createImportDeclaration(node.modifiers, node.importClause, node.moduleSpecifier);
        }
        return typescript_1.default.visitEachChild(node, (child) => visit(child), context);
    };
    return (node) => {
        typescript_1.default.visitNode(node, visit);
        return node;
    };
};
function getScriptContent(markup, module) {
    const regex = (0, markup_1.createTagRegex)('script', 'gi');
    let match;
    while ((match = regex.exec(markup)) !== null) {
        const { context } = (0, markup_1.parseAttributes)(match[1] || '');
        if ((context !== 'module' && !module) || (context === 'module' && module)) {
            return match[2];
        }
    }
    return '';
}
function createSourceMapChain({ filename, content, compilerOptions, }) {
    if (compilerOptions.sourceMap) {
        return {
            content: {
                [filename]: content,
            },
            sourcemaps: {},
        };
    }
}
function injectVarsToCode({ content, markup, filename, attributes, sourceMapChain, }) {
    if (!markup)
        return content;
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore different in Svelte 5
    const { vars } = (0, compiler_1.compile)((0, markup_1.stripTags)(markup), {
        generate: false,
        varsReport: 'full',
        errorMode: 'warn',
        filename,
    });
    const sep = `\n${injectedCodeSeparator}\n`;
    const varnames = vars.map((v) => v.name.startsWith('$') && !v.name.startsWith('$$')
        ? `${v.name},${v.name.slice(1)}`
        : v.name);
    const contentForCodestores = content +
        // Append instance script content because it's valid
        // to import a store in module script and autosubscribe to it in instance script
        ((attributes === null || attributes === void 0 ? void 0 : attributes.context) === 'module' ? getScriptContent(markup, false) : '');
    // This regex extracts all possible store variables
    // TODO investigate if it's possible to achieve this with a
    // TS transformer (previous attemps have failed)
    const codestores = Array.from(contentForCodestores.match(/\$[^\s();:,[\]{}.?!+\-=*/\\~|&%<>^`"'°§#0-9][^\s();:,[\]{}.?!+\-=*/\\~|&%<>^`"'°§#]*/g) || [], (name) => name.slice(1)).filter((name) => !utils_1.JAVASCRIPT_RESERVED_KEYWORD_SET.has(name));
    const varsString = [...codestores, ...varnames].join(',');
    const injectedVars = `const $$vars$$ = [${varsString}];`;
    // Append instance/markup script content because it's valid
    // to import things in one and reference it in the other.
    const injectedCode = (attributes === null || attributes === void 0 ? void 0 : attributes.context) === 'module'
        ? `${sep}${getScriptContent(markup, false)}\n${injectedVars}`
        : `${sep}${getScriptContent(markup, true)}\n${injectedVars}`;
    if (sourceMapChain) {
        const ms = new magic_string_1.default(content).append(injectedCode);
        const fname = `${filename}.injected.ts`;
        const code = ms.toString();
        const map = ms.generateMap({
            source: filename,
            file: fname,
            hires: true,
        });
        sourceMapChain.content[fname] = code;
        sourceMapChain.sourcemaps[fname] = map;
        return code;
    }
    return `${content}${injectedCode}`;
}
function stripInjectedCode({ transpiledCode, markup, filename, sourceMapChain, }) {
    if (!markup)
        return transpiledCode;
    const injectedCodeStart = transpiledCode.indexOf(injectedCodeSeparator);
    if (sourceMapChain) {
        const ms = new magic_string_1.default(transpiledCode).snip(0, injectedCodeStart);
        const source = `${filename}.transpiled.js`;
        const file = `${filename}.js`;
        const code = ms.toString();
        const map = ms.generateMap({
            source,
            file,
            hires: true,
        });
        sourceMapChain.content[file] = code;
        sourceMapChain.sourcemaps[file] = map;
        return code;
    }
    return transpiledCode.slice(0, injectedCodeStart);
}
async function concatSourceMaps({ filename, markup, sourceMapChain, }) {
    if (!sourceMapChain)
        return;
    if (!markup) {
        return sourceMapChain.sourcemaps[`${filename}.js`];
    }
    const chain = await (0, sorcery_1.load)(`${filename}.js`, sourceMapChain);
    return chain.apply();
}
let warned = false;
function getCompilerOptions({ filename, options, basePath, }) {
    var _a, _b;
    const inputOptions = (_a = options.compilerOptions) !== null && _a !== void 0 ? _a : {};
    const { errors, options: convertedCompilerOptions } = options.tsconfigFile !== false || options.tsconfigDirectory
        ? loadTsconfig(inputOptions, filename, options)
        : typescript_1.default.convertCompilerOptionsFromJson(inputOptions, basePath);
    if (errors.length) {
        throw new Error(formatDiagnostics(errors, basePath));
    }
    const compilerOptions = {
        target: typescript_1.default.ScriptTarget.ES2015,
        ...convertedCompilerOptions,
        // force module(resolution) to esnext and a compatible moduleResolution. Reason:
        // transpileModule treats NodeNext as CommonJS because it doesn't read the package.json.
        // Also see https://github.com/microsoft/TypeScript/issues/53022 (the filename workaround doesn't work).
        module: typescript_1.default.ModuleKind.ESNext,
        moduleResolution: (_b = typescript_1.default.ModuleResolutionKind.NodeJs) !== null && _b !== void 0 ? _b : typescript_1.default.ModuleResolutionKind.Node10,
        allowNonTsExtensions: true,
        // Clear outDir since it causes source map issues when the files aren't actually written to disk.
        outDir: undefined,
    };
    if (!importsNotUsedAsValuesDeprecated ||
        (compilerOptions.ignoreDeprecations === '5.0' &&
            !compilerOptions.verbatimModuleSyntax)) {
        compilerOptions.importsNotUsedAsValues = typescript_1.default.ImportsNotUsedAsValues.Error;
    }
    // Ease TS 5 migration pains a little and add the deprecation flag automatically if needed
    if (importsNotUsedAsValuesDeprecated &&
        !compilerOptions.ignoreDeprecations &&
        (compilerOptions.importsNotUsedAsValues ||
            compilerOptions.preserveValueImports)) {
        if (!warned) {
            warned = true;
            console.warn('tsconfig options "importsNotUsedAsValues" and "preserveValueImports" are deprecated. ' +
                'Either set "ignoreDeprecations" to "5.0" in your tsconfig.json to silence this warning, ' +
                'or replace them in favor of the new "verbatimModuleSyntax" flag.');
        }
        compilerOptions.ignoreDeprecations = '5.0';
    }
    if (compilerOptions.target === typescript_1.default.ScriptTarget.ES3 ||
        compilerOptions.target === typescript_1.default.ScriptTarget.ES5) {
        throw new Error(`Svelte only supports es6+ syntax. Set your 'compilerOptions.target' to 'es6' or higher.`);
    }
    return compilerOptions;
}
function transpileTs({ code, fileName, basePath, options, compilerOptions, transformers, }) {
    const { outputText: transpiledCode, sourceMapText, diagnostics, } = typescript_1.default.transpileModule(code, {
        fileName,
        compilerOptions,
        reportDiagnostics: options.reportDiagnostics !== false,
        transformers,
    });
    if (diagnostics && diagnostics.length > 0) {
        // could this be handled elsewhere?
        const hasError = diagnostics.some((d) => d.category === typescript_1.default.DiagnosticCategory.Error);
        const formattedDiagnostics = formatDiagnostics(diagnostics, basePath);
        console.log(formattedDiagnostics);
        if (hasError) {
            (0, errors_1.throwTypescriptError)();
        }
    }
    return { transpiledCode, sourceMapText, diagnostics };
}
function loadTsconfig(compilerOptionsJSON, filename, tsOptions) {
    if (typeof tsOptions.tsconfigFile === 'boolean') {
        return { errors: [], options: compilerOptionsJSON };
    }
    let basePath = process.cwd();
    const fileDirectory = (tsOptions.tsconfigDirectory ||
        (0, path_1.dirname)(filename));
    let tsconfigFile = tsOptions.tsconfigFile ||
        typescript_1.default.findConfigFile(fileDirectory, typescript_1.default.sys.fileExists);
    if (!tsconfigFile) {
        return { errors: [], options: compilerOptionsJSON };
    }
    tsconfigFile = (0, path_1.isAbsolute)(tsconfigFile)
        ? tsconfigFile
        : (0, path_1.join)(basePath, tsconfigFile);
    basePath = (0, path_1.dirname)(tsconfigFile);
    if (tsconfigMap.has(tsconfigFile)) {
        return {
            errors: [],
            options: tsconfigMap.get(tsconfigFile),
        };
    }
    const { error, config } = typescript_1.default.readConfigFile(tsconfigFile, typescript_1.default.sys.readFile);
    if (error) {
        throw new Error(formatDiagnostics(error, basePath));
    }
    // Do this so TS will not search for initial files which might take a while
    config.include = [];
    let { errors, options } = typescript_1.default.parseJsonConfigFileContent(config, typescript_1.default.sys, basePath, compilerOptionsJSON, tsconfigFile);
    // Filter out "no files found error"
    errors = errors.filter((d) => d.code !== 18003);
    if (errors.length === 0) {
        tsconfigMap.set(tsconfigFile, options);
    }
    return { errors, options };
}
exports.loadTsconfig = loadTsconfig;
async function mixedImportsTranspiler({ content, filename = 'source.svelte', markup, options = {}, attributes, compilerOptions, basePath, }) {
    const sourceMapChain = createSourceMapChain({
        filename,
        content,
        compilerOptions,
    });
    const injectedCode = injectVarsToCode({
        content,
        markup,
        filename,
        attributes,
        sourceMapChain,
    });
    const { transpiledCode, sourceMapText, diagnostics } = transpileTs({
        code: injectedCode,
        fileName: `${filename}.injected.ts`,
        basePath,
        options,
        compilerOptions,
    });
    if (sourceMapChain && sourceMapText) {
        const fname = `${filename}.transpiled.js`;
        sourceMapChain.content[fname] = transpiledCode;
        sourceMapChain.sourcemaps[fname] = JSON.parse(sourceMapText);
    }
    const code = stripInjectedCode({
        transpiledCode,
        markup,
        filename,
        sourceMapChain,
    });
    // Sorcery tries to load the code/map from disk if it's empty,
    // prevent that because it would try to load inexistent files
    // https://github.com/Rich-Harris/sorcery/issues/167
    if (!code) {
        return { code, diagnostics };
    }
    const map = await concatSourceMaps({
        filename,
        markup,
        sourceMapChain,
    });
    return {
        code,
        map,
        diagnostics,
    };
}
async function simpleTranspiler({ content, filename = 'source.svelte', options = {}, compilerOptions, basePath, }) {
    const { transpiledCode, sourceMapText, diagnostics } = transpileTs({
        code: content,
        // `preserveValueImports` essentially does the same as our import transformer,
        // keeping all imports that are not type imports
        transformers: compilerOptions.preserveValueImports ||
            compilerOptions.verbatimModuleSyntax
            ? undefined
            : { before: [importTransformer] },
        fileName: filename,
        basePath,
        options,
        compilerOptions,
    });
    return {
        code: transpiledCode,
        map: sourceMapText,
        diagnostics,
    };
}
const transformer = async ({ content, filename, markup, options = {}, attributes, }) => {
    const basePath = process.cwd();
    if (filename == null)
        return { code: content };
    filename = (0, path_1.isAbsolute)(filename) ? filename : (0, path_1.resolve)(basePath, filename);
    const compilerOptions = getCompilerOptions({ filename, options, basePath });
    const versionParts = package_json_1.default.version.split('.');
    const canUseMixedImportsTranspiler = +versionParts[0] === 4 ||
        (+versionParts[0] === 3 && +versionParts[1] >= 39);
    if (!canUseMixedImportsTranspiler && options.handleMixedImports) {
        throw new Error('You need at least Svelte 3.39 and at most Svelte 4.x to use the handleMixedImports option. ' +
            'The option is no longer available for Svelte 5 and beyond. Use the verbatimModuleSyntax TypeScript option instead.');
    }
    const handleMixedImports = !compilerOptions.preserveValueImports &&
        !compilerOptions.verbatimModuleSyntax &&
        (options.handleMixedImports === false
            ? false
            : options.handleMixedImports || canUseMixedImportsTranspiler);
    return handleMixedImports
        ? mixedImportsTranspiler({
            content,
            filename,
            markup,
            options,
            attributes,
            compilerOptions,
            basePath,
        })
        : simpleTranspiler({
            content,
            filename,
            markup,
            options,
            attributes,
            compilerOptions,
            basePath,
        });
};
exports.transformer = transformer;
