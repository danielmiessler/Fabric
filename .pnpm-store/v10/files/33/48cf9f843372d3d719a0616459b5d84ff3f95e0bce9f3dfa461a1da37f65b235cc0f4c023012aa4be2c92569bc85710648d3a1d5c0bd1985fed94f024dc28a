import type { PreprocessorArgs } from '../types';
export declare const ALIAS_MAP: Map<string, string>;
export declare const SOURCE_MAP_PROP_MAP: Record<string, [string[], any]>;
export declare function getLanguageDefaults(lang?: string | null): null | Record<string, any>;
export declare function addLanguageAlias(entries: Array<[string, string]>): void;
export declare function getLanguageFromAlias(alias?: string | null): string | null | undefined;
export declare function isAliasOf(alias?: string | null, lang?: string | null): boolean;
export declare const getLanguage: (attributes: PreprocessorArgs['attributes']) => {
    lang: string | null | undefined;
    alias: string | null;
};
