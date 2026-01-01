import { createRequire } from 'node:module';
import path from 'node:path';
import fs from 'node:fs';
import { pathToFileURL } from 'node:url';
import { log } from './log.js';

// used to require cjs config in esm.
// NOTE dynamic import() cjs technically works, but timestamp query cache bust
// have no effect, likely because it has another internal cache?
/** @type {NodeRequire}*/
let esmRequire;

export const knownSvelteConfigNames = [
	'svelte.config.js',
	'svelte.config.cjs',
	'svelte.config.mjs'
];

/**
 * @param {string} filePath
 * @param {number} timestamp
 */
async function dynamicImportDefault(filePath, timestamp) {
	return await import(filePath + '?t=' + timestamp).then((m) => m.default);
}

/**
 * @param {import('vite').UserConfig} [viteConfig]
 * @param {Partial<import('../public.d.ts').Options>} [inlineOptions]
 * @returns {Promise<Partial<import('../public.d.ts').SvelteConfig> | undefined>}
 */
export async function loadSvelteConfig(viteConfig, inlineOptions) {
	if (inlineOptions?.configFile === false) {
		return;
	}
	const configFile = findConfigToLoad(viteConfig, inlineOptions);
	if (configFile) {
		let err;
		// try to use dynamic import for svelte.config.js first
		if (configFile.endsWith('.js') || configFile.endsWith('.mjs')) {
			try {
				const result = await dynamicImportDefault(
					pathToFileURL(configFile).href,
					fs.statSync(configFile).mtimeMs
				);
				if (result != null) {
					return {
						...result,
						configFile
					};
				} else {
					throw new Error(`invalid export in ${configFile}`);
				}
			} catch (e) {
				log.error(`failed to import config ${configFile}`, e);
				err = e;
			}
		}
		// cjs or error with dynamic import
		if (!configFile.endsWith('.mjs')) {
			try {
				// identify which require function to use (esm and cjs mode)
				const _require = import.meta.url
					? esmRequire ?? (esmRequire = createRequire(import.meta.url))
					: // eslint-disable-next-line no-undef
						require;

				// avoid loading cached version on reload
				delete _require.cache[_require.resolve(configFile)];
				const result = _require(configFile);
				if (result != null) {
					return {
						...result,
						configFile
					};
				} else {
					throw new Error(`invalid export in ${configFile}`);
				}
			} catch (e) {
				log.error(`failed to require config ${configFile}`, e);
				if (!err) {
					err = e;
				}
			}
		}
		// failed to load existing config file
		throw err;
	}
}

/**
 * @param {import('vite').UserConfig | undefined} viteConfig
 * @param {Partial<import('../public.d.ts').Options> | undefined} inlineOptions
 * @returns {string | undefined}
 */
function findConfigToLoad(viteConfig, inlineOptions) {
	const root = viteConfig?.root || process.cwd();
	if (inlineOptions?.configFile) {
		const abolutePath = path.isAbsolute(inlineOptions.configFile)
			? inlineOptions.configFile
			: path.resolve(root, inlineOptions.configFile);
		if (!fs.existsSync(abolutePath)) {
			throw new Error(`failed to find svelte config file ${abolutePath}.`);
		}
		return abolutePath;
	} else {
		const existingKnownConfigFiles = knownSvelteConfigNames
			.map((candidate) => path.resolve(root, candidate))
			.filter((file) => fs.existsSync(file));
		if (existingKnownConfigFiles.length === 0) {
			log.debug(`no svelte config found at ${root}`, undefined, 'config');
			return;
		} else if (existingKnownConfigFiles.length > 1) {
			log.warn(
				`found more than one svelte config file, using ${existingKnownConfigFiles[0]}. you should only have one!`,
				existingKnownConfigFiles
			);
		}
		return existingKnownConfigFiles[0];
	}
}
