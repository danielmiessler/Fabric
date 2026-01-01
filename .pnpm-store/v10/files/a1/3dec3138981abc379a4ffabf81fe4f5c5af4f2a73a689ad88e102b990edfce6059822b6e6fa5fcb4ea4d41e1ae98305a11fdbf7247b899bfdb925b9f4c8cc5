import * as Options from './options';
import type { Processed as SvelteProcessed, Preprocessor as SveltePreprocessor, PreprocessorGroup } from 'svelte/types/compiler/preprocess';
export { Options };
export { PreprocessorGroup } from 'svelte/types/compiler/preprocess';
export type PreprocessorArgs = Preprocessor extends (options: infer T) => any ? T : never;
export type TransformerArgs<T> = {
    content: string;
    filename?: string;
    attributes?: Record<string, any>;
    map?: string | object;
    markup?: string;
    diagnostics?: unknown[];
    options?: T;
};
/**
 * Small extension to the official SvelteProcessed type
 * to include possible diagnostics.
 * Used for the typescript transformer.
 */
export type Processed = SvelteProcessed & {
    diagnostics?: any[];
};
/**
 * Svelte preprocessor type with guaranteed Processed results
 *
 * The official type also considers `void`
 * */
export type Preprocessor = (options: Parameters<SveltePreprocessor>[0]) => Processed | Promise<Processed>;
export type Transformer<T> = (args: TransformerArgs<T>) => Processed | Promise<Processed>;
export type TransformerOptions<T = any> = boolean | T | Transformer<T>;
export interface Transformers {
    babel?: TransformerOptions<Options.Babel>;
    typescript?: TransformerOptions<Options.Typescript>;
    scss?: TransformerOptions<Options.Sass>;
    sass?: TransformerOptions<Options.Sass>;
    less?: TransformerOptions<Options.Less>;
    stylus?: TransformerOptions<Options.Stylus>;
    postcss?: TransformerOptions<Options.Postcss>;
    coffeescript?: TransformerOptions<Options.Coffeescript>;
    pug?: TransformerOptions<Options.Pug>;
    globalStyle?: Options.GlobalStyle;
    replace?: Options.Replace;
    [language: string]: TransformerOptions;
}
export type AutoPreprocessGroup = PreprocessorGroup;
export type AutoPreprocessOptions = {
    markupTagName?: string;
    aliases?: Array<[string, string]>;
    preserve?: string[];
    sourceMap?: boolean;
    babel?: TransformerOptions<Options.Babel>;
    typescript?: TransformerOptions<Options.Typescript>;
    scss?: TransformerOptions<Options.Sass>;
    sass?: TransformerOptions<Options.Sass>;
    less?: TransformerOptions<Options.Less>;
    stylus?: TransformerOptions<Options.Stylus>;
    postcss?: TransformerOptions<Options.Postcss>;
    coffeescript?: TransformerOptions<Options.Coffeescript>;
    pug?: TransformerOptions<Options.Pug>;
    globalStyle?: Options.GlobalStyle | boolean;
    replace?: Options.Replace;
    [languageName: string]: TransformerOptions;
};
