import { promises as fs } from 'node:fs';
import path from 'node:path';

// List of options that changes the prebundling result
/** @type {(keyof import('../types/options.d.ts').ResolvedOptions)[]} */
const PREBUNDLE_SENSITIVE_OPTIONS = [
	'compilerOptions',
	'configFile',
	'experimental',
	'extensions',
	'ignorePluginPreprocessors',
	'preprocess'
];

/**
 * @param {string} cacheDir
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {Promise<boolean>} Whether the Svelte metadata has changed
 */
export async function saveSvelteMetadata(cacheDir, options) {
	const svelteMetadata = generateSvelteMetadata(options);
	const svelteMetadataPath = path.resolve(cacheDir, '_svelte_metadata.json');

	const currentSvelteMetadata = JSON.stringify(svelteMetadata, (_, value) => {
		// Handle preprocessors
		return typeof value === 'function' ? value.toString() : value;
	});

	/** @type {string | undefined} */
	let existingSvelteMetadata;
	try {
		existingSvelteMetadata = await fs.readFile(svelteMetadataPath, 'utf8');
	} catch {
		// ignore
	}

	await fs.mkdir(cacheDir, { recursive: true });
	await fs.writeFile(svelteMetadataPath, currentSvelteMetadata);
	return currentSvelteMetadata !== existingSvelteMetadata;
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {Partial<import('../types/options.d.ts').ResolvedOptions>}
 */
function generateSvelteMetadata(options) {
	/** @type {Record<string, any>} */
	const metadata = {};
	for (const key of PREBUNDLE_SENSITIVE_OPTIONS) {
		metadata[key] = options[key];
	}
	return metadata;
}
