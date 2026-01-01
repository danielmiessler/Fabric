/**
 * Core package of PurgeCSS
 *
 * Contains the core methods to analyze the files, remove unused CSS.
 *
 * @packageDocumentation
 */

import * as postcss from 'postcss';

/* Excluded from this release type: AtRules */

/**
 * @public
 */
export declare type ComplexSafelist = {
    standard?: StringRegExpArray;
    /**
     * You can safelist selectors and their children based on a regular
     * expression with `safelist.deep`
     *
     * @example
     *
     * ```ts
     * const purgecss = await new PurgeCSS().purge({
     *   content: [],
     *   css: [],
     *   safelist: {
     *     deep: [/red$/]
     *   }
     * })
     * ```
     *
     * In this example, selectors such as `.bg-red .child-of-bg` will be left
     * in the final CSS, even if `child-of-bg` is not found.
     *
     */
    deep?: RegExp[];
    greedy?: RegExp[];
    variables?: StringRegExpArray;
    keyframes?: StringRegExpArray;
};

/**
 * @public
 */
export declare const defaultOptions: Options;

/**
 * @public
 */
export declare type ExtractorFunction<T = string> = (content: T) => ExtractorResult;

/**
 * @public
 */
export declare type ExtractorResult = ExtractorResultDetailed | string[];

/**
 * @public
 */
export declare interface ExtractorResultDetailed {
    attributes: {
        names: string[];
        values: string[];
    };
    classes: string[];
    ids: string[];
    tags: string[];
    undetermined: string[];
}

/**
 * @public
 */
export declare class ExtractorResultSets {
    private undetermined;
    private attrNames;
    private attrValues;
    private classes;
    private ids;
    private tags;
    constructor(er: ExtractorResult);
    merge(that: ExtractorResult | ExtractorResultSets): this;
    hasAttrName(name: string): boolean;
    private someAttrValue;
    hasAttrPrefix(prefix: string): boolean;
    hasAttrSuffix(suffix: string): boolean;
    hasAttrSubstr(substr: string): boolean;
    hasAttrValue(value: string): boolean;
    hasClass(name: string): boolean;
    hasId(id: string): boolean;
    hasTag(tag: string): boolean;
}

/**
 * @public
 */
export declare interface Extractors {
    extensions: string[];
    extractor: ExtractorFunction;
}

/* Excluded from this release type: IgnoreType */

/**
 * Merge two extractor selectors
 *
 * @param extractorSelectorsA - extractor selectors A
 * @param extractorSelectorsB - extractor selectors B
 * @returns  the merged extractor result sets
 *
 * @public
 */
export declare function mergeExtractorSelectors(...extractors: (ExtractorResultDetailed | ExtractorResultSets)[]): ExtractorResultSets;

/**
 * Options used by PurgeCSS to remove unused CSS
 * Those options are used internally
 * @see {@link UserDefinedOptions} for the options defined by the user
 *
 * @public
 */
export declare interface Options {
    /**
     * You can specify content that should be analyzed by PurgeCSS with an
     * array of filenames or globs. The files can be HTML, Pug, Blade, etc.
     *
     * @example
     *
     * ```ts
     * await new PurgeCSS().purge({
     *   content: ['index.html', '*.js', '*.html', '*.vue'],
     *   css: ['css/app.css']
     * })
     * ```
     *
     * @example
     * PurgeCSS also works with raw content. To do this, you need to pass an
     * object with the `raw` property instead of a filename. To work properly
     * with custom extractors you need to pass the `extension` property along
     * with the raw content.
     *
     * ```ts
     * await new PurgeCSS().purge({
     *   content: [
     *     {
     *       raw: '<html><body><div class="app"></div></body></html>',
     *       extension: 'html'
     *     },
     *     '*.js',
     *     '*.html',
     *     '*.vue'
     *   ],
     *   css: [
     *     {
     *       raw: 'body { margin: 0 }'
     *     },
     *     'css/app.css'
     *   ]
     * })
     * ```
     */
    content: Array<string | RawContent>;
    /**
     * Similar to content, you can specify css that should be processed by
     * PurgeCSS with an array of filenames or globs
     */
    css: Array<string | RawCSS>;
    defaultExtractor: ExtractorFunction;
    extractors: Array<Extractors>;
    /**
     * If there are any unused \@font-face rules in your css, you can remove
     * them by setting the `fontFace` option to `true`.
     *
     * @defaultValue `false`
     *
     * @example
     * ```ts
     * await new PurgeCSS().purge({
     *   content: ['index.html', '*.js', '*.html', '*.vue'],
     *   css: ['css/app.css'],
     *   fontFace: true
     * })
     * ```
     */
    fontFace: boolean;
    keyframes: boolean;
    output?: string;
    rejected: boolean;
    rejectedCss: boolean;
    /** {@inheritDoc postcss#SourceMapOptions} */
    sourceMap: boolean | (postcss.SourceMapOptions & {
        to?: string;
    });
    stdin: boolean;
    stdout: boolean;
    variables: boolean;
    /**
     * You can indicate which selectors are safe to leave in the final CSS.
     * This can be accomplished with the option safelist.
     */
    safelist: Required<ComplexSafelist>;
    /**
     * Blocklist will block the CSS selectors from appearing in the final
     * output CSS. The selectors will be removed even when they are seen
     * as used by PurgeCSS.
     */
    blocklist: StringRegExpArray;
    /**
     * If you provide globs for the content parameter, you can use this option
     * to exclude certain files or folders that would otherwise be scanned.
     * Pass an array of globs matching items that should be excluded.
     * (Note: this option has no effect if content is not globs.)
     */
    skippedContentGlobs: Array<string>;
    /**
     * Option to add custom CSS attribute selectors like "aria-selected",
     * "data-selected", ...etc.
     */
    dynamicAttributes: string[];
}

/**
 * @public
 */
export declare type PostCSSRoot = postcss.Root;

/**
 * Class used to instantiate PurgeCSS and can then be used
 * to purge CSS files.
 *
 * @example
 * ```ts
 * await new PurgeCSS().purge({
 *    content: ['index.html'],
 *    css: ['css/app.css']
 * })
 * ```
 *
 * @public
 */
export declare class PurgeCSS {
    private ignore;
    private atRules;
    private usedAnimations;
    private usedFontFaces;
    selectorsRemoved: Set<string>;
    removedNodes: postcss.Node[];
    variablesStructure: VariablesStructure;
    options: Options;
    private collectDeclarationsData;
    /**
     * Get the extractor corresponding to the extension file
     * @param filename - Name of the file
     * @param extractors - Array of extractors definition
     */
    private getFileExtractor;
    /**
     * Extract the selectors present in the files using a PurgeCSS extractor
     *
     * @param files - Array of files path or glob pattern
     * @param extractors - Array of extractors
     */
    extractSelectorsFromFiles(files: string[], extractors: Extractors[]): Promise<ExtractorResultSets>;
    /**
     * Extract the selectors present in the passed string using a PurgeCSS extractor
     *
     * @param content - Array of content
     * @param extractors - Array of extractors
     */
    extractSelectorsFromString(content: RawContent[], extractors: Extractors[]): Promise<ExtractorResultSets>;
    /**
     * Evaluate at-rule and register it for future reference
     * @param node - node of postcss AST
     */
    private evaluateAtRule;
    /**
     * Evaluate css selector and decide if it should be removed or not
     *
     * @param node - node of postcss AST
     * @param selectors - selectors used in content files
     */
    private evaluateRule;
    /**
     * Get the purged version of the css based on the files
     *
     * @param cssOptions - css options, files or raw strings
     * @param selectors - set of extracted css selectors
     */
    getPurgedCSS(cssOptions: Array<string | RawCSS>, selectors: ExtractorResultSets): Promise<ResultPurge[]>;
    /**
     * Check if the keyframe is safelisted with the option safelist keyframes
     *
     * @param keyframesName - name of the keyframe animation
     */
    private isKeyframesSafelisted;
    /**
     * Check if the selector is blocklisted with the option blocklist
     *
     * @param selector - css selector
     */
    private isSelectorBlocklisted;
    /**
     * Check if the selector is safelisted with the option safelist standard
     *
     * @param selector - css selector
     */
    private isSelectorSafelisted;
    /**
     * Check if the selector is safelisted with the option safelist deep
     *
     * @param selector - selector
     */
    private isSelectorSafelistedDeep;
    /**
     * Check if the selector is safelisted with the option safelist greedy
     *
     * @param selector - selector
     */
    private isSelectorSafelistedGreedy;
    /**
     * Remove unused CSS
     *
     * @param userOptions - PurgeCSS options or path to the configuration file
     * @returns an array of object containing the filename and the associated CSS
     *
     * @example Using a configuration file named purgecss.config.js
     * ```ts
     * const purgeCSSResults = await new PurgeCSS().purge()
     * ```
     *
     * @example Using a custom path to the configuration file
     * ```ts
     * const purgeCSSResults = await new PurgeCSS().purge('./purgecss.config.js')
     * ```
     *
     * @example Using the PurgeCSS options
     * ```ts
     * const purgeCSSResults = await new PurgeCSS().purge({
     *    content: ['index.html', '**\/*.js', '**\/*.html', '**\/*.vue'],
     *    css: ['css/app.css']
     * })
     * ```
     */
    purge(userOptions: UserDefinedOptions | string | undefined): Promise<ResultPurge[]>;
    /**
     * Remove unused CSS variables
     */
    removeUnusedCSSVariables(): void;
    /**
     * Remove unused font-faces
     */
    removeUnusedFontFaces(): void;
    /**
     * Remove unused keyframes
     */
    removeUnusedKeyframes(): void;
    /**
     * Transform a selector node into a string
     */
    private getSelectorValue;
    /**
     * Determine if the selector should be kept, based on the selectors found in the files
     *
     * @param selector - set of css selectors found in the content files or string
     * @param selectorsFromExtractor - selectors in the css rule
     *
     * @returns true if the selector should be kept in the processed CSS
     */
    private shouldKeepSelector;
    /**
     * Walk through the CSS AST and remove unused CSS
     *
     * @param root - root node of the postcss AST
     * @param selectors - selectors used in content files
     */
    walkThroughCSS(root: PostCSSRoot, selectors: ExtractorResultSets): void;
}

/**
 * @public
 */
export declare interface RawContent<T = string> {
    extension: string;
    raw: T;
}

/**
 * @public
 */
export declare interface RawCSS {
    raw: string;
    name?: string;
}

/**
 * @public
 */
export declare interface ResultPurge {
    css: string;
    /**
     * sourceMap property will be empty if
     * {@link UserDefinedOptions.sourceMap} inline is not set to false, as the
     * source map will be contained within the text of ResultPurge.css
     */
    sourceMap?: string;
    rejectedCss?: string;
    file?: string;
    rejected?: string[];
}

/**
 * Load the configuration file from the path
 *
 * @param configFile - Path of the config file
 * @returns The options from the configuration file
 *
 * @throws Error
 * This exception is thrown if the configuration file was not imported
 *
 * @public
 */
export declare function setOptions(configFile?: string): Promise<Options>;

/**
 * Format the user defined safelist into a standardized safelist object
 *
 * @param userDefinedSafelist - the user defined safelist
 * @returns the formatted safelist object that can be used in the PurgeCSS options
 *
 * @public
 */
export declare function standardizeSafelist(userDefinedSafelist?: UserDefinedSafelist): Required<ComplexSafelist>;

/**
 * @public
 */
export declare type StringRegExpArray = Array<RegExp | string>;

/**
 * Options used by PurgeCSS to remove unused CSS
 *
 * @public
 */
export declare interface UserDefinedOptions {
    /** {@inheritDoc purgecss#Options.content} */
    content: Array<string | RawContent>;
    /** {@inheritDoc purgecss#Options.css} */
    css: Array<string | RawCSS>;
    /** {@inheritDoc purgecss#Options.defaultExtractor} */
    defaultExtractor?: ExtractorFunction;
    /** {@inheritDoc purgecss#Options.extractors} */
    extractors?: Array<Extractors>;
    /** {@inheritDoc purgecss#Options.fontFace} */
    fontFace?: boolean;
    /** {@inheritDoc purgecss#Options.keyframes} */
    keyframes?: boolean;
    /** {@inheritDoc purgecss#Options.output} */
    output?: string;
    /** {@inheritDoc purgecss#Options.rejected} */
    rejected?: boolean;
    /** {@inheritDoc purgecss#Options.rejectedCss} */
    rejectedCss?: boolean;
    /** {@inheritDoc purgecss#Options.sourceMap } */
    sourceMap?: boolean | (postcss.SourceMapOptions & {
        to?: string;
    });
    /** {@inheritDoc purgecss#Options.stdin} */
    stdin?: boolean;
    /** {@inheritDoc purgecss#Options.stdout} */
    stdout?: boolean;
    /** {@inheritDoc purgecss#Options.variables} */
    variables?: boolean;
    /** {@inheritDoc purgecss#Options.safelist} */
    safelist?: UserDefinedSafelist;
    /** {@inheritDoc purgecss#Options.blocklist} */
    blocklist?: StringRegExpArray;
    /** {@inheritDoc purgecss#Options.skippedContentGlobs} */
    skippedContentGlobs?: Array<string>;
    /** {@inheritDoc purgecss#Options.dynamicAttributes} */
    dynamicAttributes?: string[];
}

/**
 * @public
 */
export declare type UserDefinedSafelist = StringRegExpArray | ComplexSafelist;

/**
 * @public
 */
export declare class VariableNode {
    nodes: VariableNode[];
    value: postcss.Declaration;
    isUsed: boolean;
    constructor(declaration: postcss.Declaration);
}

/**
 * @public
 */
export declare class VariablesStructure {
    nodes: Map<string, VariableNode[]>;
    usedVariables: Set<string>;
    safelist: StringRegExpArray;
    addVariable(declaration: postcss.Declaration): void;
    addVariableUsage(declaration: postcss.Declaration, matchedVariables: IterableIterator<RegExpMatchArray>): void;
    addVariableUsageInProperties(matchedVariables: IterableIterator<RegExpMatchArray>): void;
    setAsUsed(variableName: string): void;
    removeUnused(): void;
    isVariablesSafelisted(variable: string): boolean;
}

export { }
