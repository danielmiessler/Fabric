/** Internal version of $app/environment */
declare module '__sveltekit/environment' {
	export const building: boolean;
	export const prerendering: boolean;
	export const version: string;
	export function set_building(): void;
	export function set_prerendering(): void;
}

/** Internal version of $app/paths */
declare module '__sveltekit/paths' {
	export let base: '' | `/${string}`;
	export let assets: '' | `https://${string}` | `http://${string}` | '/_svelte_kit_assets';
	export let app_dir: string;
	export let relative: boolean;
	export function reset(): void;
	export function override(paths: { base: string; assets: string }): void;
	export function set_assets(path: string): void;
}

/** Internal version of $app/server */
declare module '__sveltekit/server' {
	import { SSRManifest } from '@sveltejs/kit';

	export let manifest: SSRManifest;
	export function read_implementation(path: string): ReadableStream;
	export function set_manifest(manifest: SSRManifest): void;
	export function set_read_implementation(fn: (path: string) => ReadableStream): void;
}
