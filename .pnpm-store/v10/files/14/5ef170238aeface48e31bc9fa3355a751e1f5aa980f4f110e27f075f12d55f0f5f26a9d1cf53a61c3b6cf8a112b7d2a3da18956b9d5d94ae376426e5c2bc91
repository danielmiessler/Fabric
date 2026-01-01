import type { CompileOptions } from 'svelte/types/compiler/interfaces';

export type SvelteQueryTypes = 'style' | 'script' | 'preprocessed' | 'all';

export interface RequestQuery {
	// our own
	svelte?: boolean;
	type?: SvelteQueryTypes;
	sourcemap?: boolean;
	compilerOptions?: Pick<
		CompileOptions,
		'generate' | 'dev' | 'css' | 'hydratable' | 'customElement' | 'immutable' | 'enableSourcemap'
	>;
	// vite specific
	url?: boolean;
	raw?: boolean;
	direct?: boolean;
}

export interface SvelteRequest {
	id: string;
	cssId: string;
	filename: string;
	normalizedFilename: string;
	query: RequestQuery;
	timestamp: number;
	ssr: boolean;
	raw: boolean;
}

export interface SvelteModuleRequest {
	id: string;
	filename: string;
	normalizedFilename: string;
	query: RequestQuery;
	timestamp: number;
	ssr: boolean;
}

export type IdParser = (id: string, ssr: boolean, timestamp?: number) => SvelteRequest | undefined;
export type ModuleIdParser = (
	id: string,
	ssr: boolean,
	timestamp?: number
) => SvelteModuleRequest | undefined;
