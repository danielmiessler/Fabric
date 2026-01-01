import { loadEnv } from 'vite';
import { debug } from './debug.js';

/** @type {import('./public.d.ts').Options} */
export const defaultInspectorOptions = {
	toggleKeyCombo: process.platform === 'darwin' ? 'meta-shift' : 'control-shift',
	navKeys: { parent: 'ArrowUp', child: 'ArrowDown', next: 'ArrowRight', prev: 'ArrowLeft' },
	escapeKeys: ['Backspace', 'Escape'],
	openKey: 'Enter',
	holdMode: true,
	showToggleButton: 'active',
	toggleButtonPos: 'top-right',
	customStyles: true
};

/**
 * @param {import('vite').ResolvedConfig} config
 * @returns {Partial<import('./public.d.ts').Options> | boolean | void}
 */
export function parseEnvironmentOptions(config) {
	const env = loadEnv(config.mode, config.envDir ?? process.cwd(), 'SVELTE_INSPECTOR');
	const options = env?.SVELTE_INSPECTOR_OPTIONS;
	const toggle = env?.SVELTE_INSPECTOR_TOGGLE;
	if (options) {
		try {
			const parsed = JSON.parse(options);
			const parsedType = typeof parsed;
			if (parsedType === 'boolean') {
				return parsed;
			} else if (parsedType === 'object') {
				if (Array.isArray(parsed)) {
					throw new Error('invalid type, expected object map but got array');
				}
				const parsedKeys = Object.keys(parsed);
				const defaultKeys = Object.keys(defaultInspectorOptions);
				const unknownKeys = parsedKeys.filter((k) => !defaultKeys.includes(k));
				if (unknownKeys.length) {
					config.logger.warn(
						`[vite-plugin-svelte-inspector] ignoring unknown options in environment SVELTE_INSPECTOR_OPTIONS: ${unknownKeys.join(
							', '
						)}`
					);
					for (const key of unknownKeys) {
						delete parsed[key];
					}
				}
				debug('loaded environment config', parsed);
				return parsed;
			}
		} catch (e) {
			config.logger.error(
				`[vite-plugin-svelte-inspector] failed to parse inspector options from environment SVELTE_INSPECTOR_OPTIONS="${options}"\n${e}`
			);
		}
	} else if (toggle) {
		const keyConfig = {
			toggleKeyCombo: toggle
		};
		debug('loaded environment config', keyConfig);
		return keyConfig;
	}
}
