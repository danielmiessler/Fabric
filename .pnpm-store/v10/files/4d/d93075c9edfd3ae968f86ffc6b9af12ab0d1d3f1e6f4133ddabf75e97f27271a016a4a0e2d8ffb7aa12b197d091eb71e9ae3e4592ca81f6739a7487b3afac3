import type { SvelteParseContext } from "./svelte-parse-context";
declare const globalsForSvelte: readonly ["$$slots", "$$props", "$$restProps"];
export declare const globalsForRunes: readonly ["$state", "$derived", "$effect", "$props", "$bindable", "$inspect", "$host"];
type Global = (typeof globalsForSvelte)[number] | (typeof globalsForRunes)[number];
export declare function getGlobalsForSvelte(svelteParseContext: SvelteParseContext): readonly Global[];
export declare function getGlobalsForSvelteScript(svelteParseContext: SvelteParseContext): readonly Global[];
export {};
