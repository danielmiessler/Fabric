import { isSvelte5 } from './svelte-version.js';

export const VITE_RESOLVE_MAIN_FIELDS = ['browser', 'module', 'jsnext:main', 'jsnext'];

export const SVELTE_RESOLVE_MAIN_FIELDS = ['svelte'];

export const SVELTE_IMPORTS = [
	'svelte/animate',
	'svelte/easing',
	'svelte/internal',
	'svelte/internal/disclose-version',
	'svelte/motion',
	'svelte/ssr',
	'svelte/store',
	'svelte/transition',
	'svelte'
];

export const SVELTE_HMR_IMPORTS = [
	'svelte-hmr/runtime/hot-api-esm.js',
	'svelte-hmr/runtime/proxy-adapter-dom.js',
	'svelte-hmr'
];

if (isSvelte5) {
	SVELTE_IMPORTS.push('svelte/server', 'svelte/internal/server', 'svelte/legacy');
	SVELTE_HMR_IMPORTS.length = 0; // truncate, svelte-hmr isn't used with svelte5
}

export const SVELTE_EXPORT_CONDITIONS = ['svelte'];

export const FAQ_LINK_MISSING_EXPORTS_CONDITION =
	'https://github.com/sveltejs/vite-plugin-svelte/blob/main/docs/faq.md#missing-exports-condition';
