import path from 'node:path';
import { GENERATED_COMMENT } from '../../constants.js';
import { write_if_changed } from './utils.js';

// `declare module "svelte/elements"` needs to happen in a non-ambient module, and dts-buddy generates one big ambient module,
// so we can't add it there - therefore generate the typings ourselves here.
// We're not using the `declare namespace svelteHTML` variant because that one doesn't augment the HTMLAttributes interface
// people could use to type their own components.
// The T generic is needed or else there's a "all declarations must have identical type parameters" error.
const template = `
${GENERATED_COMMENT}

declare module "svelte/elements" {
	export interface HTMLAttributes<T> {
		'data-sveltekit-keepfocus'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-noscroll'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-preload-code'?:
			| true
			| ''
			| 'eager'
			| 'viewport'
			| 'hover'
			| 'tap'
			| 'off'
			| undefined
			| null;
		'data-sveltekit-preload-data'?: true | '' | 'hover' | 'tap' | 'off' | undefined | null;
		'data-sveltekit-reload'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-replacestate'?: true | '' | 'off' | undefined | null;
	}
}

export {};
`;

/**
 * Writes non-ambient declarations to the output directory
 * @param {import('types').ValidatedKitConfig} config
 */
export function write_non_ambient(config) {
	write_if_changed(path.join(config.outDir, 'non-ambient.d.ts'), template);
}
