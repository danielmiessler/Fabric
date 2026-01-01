import path from 'node:path';
import create_manifest_data from './create_manifest_data/index.js';
import { write_client_manifest } from './write_client_manifest.js';
import { write_root } from './write_root.js';
import { write_tsconfig } from './write_tsconfig.js';
import { write_types, write_all_types } from './write_types/index.js';
import { write_ambient } from './write_ambient.js';
import { write_non_ambient } from './write_non_ambient.js';
import { write_server } from './write_server.js';

/**
 * Initialize SvelteKit's generated files.
 * @param {import('types').ValidatedConfig} config
 * @param {string} mode
 */
export function init(config, mode) {
	write_tsconfig(config.kit);
	write_ambient(config.kit, mode);
	write_non_ambient(config.kit);
}

/**
 * Update SvelteKit's generated files
 * @param {import('types').ValidatedConfig} config
 */
export function create(config) {
	const manifest_data = create_manifest_data({ config });

	const output = path.join(config.kit.outDir, 'generated');

	write_client_manifest(config.kit, manifest_data, `${output}/client`);
	write_server(config, output);
	write_root(manifest_data, output);
	write_all_types(config, manifest_data);

	return { manifest_data };
}

/**
 * Update SvelteKit's generated files in response to a single file content update.
 * Do not call this when the file in question was created/deleted.
 *
 * @param {import('types').ValidatedConfig} config
 * @param {import('types').ManifestData} manifest_data
 * @param {string} file
 */
export function update(config, manifest_data, file) {
	write_types(config, manifest_data, file);
}

/**
 * Run sync.init and sync.create in series, returning the result from sync.create.
 * @param {import('types').ValidatedConfig} config
 * @param {string} mode The Vite mode
 */
export function all(config, mode) {
	init(config, mode);
	return create(config);
}

/**
 * Run sync.init and then generate all type files.
 * @param {import('types').ValidatedConfig} config
 * @param {string} mode The Vite mode
 */
export function all_types(config, mode) {
	init(config, mode);
	const manifest_data = create_manifest_data({ config });
	write_all_types(config, manifest_data);
}

/**
 * Regenerate __SERVER__/internal.js in response to src/{app.html,error.html,service-worker.js} changing
 * @param {import('types').ValidatedConfig} config
 */
export function server(config) {
	write_server(config, path.join(config.kit.outDir, 'generated'));
}
