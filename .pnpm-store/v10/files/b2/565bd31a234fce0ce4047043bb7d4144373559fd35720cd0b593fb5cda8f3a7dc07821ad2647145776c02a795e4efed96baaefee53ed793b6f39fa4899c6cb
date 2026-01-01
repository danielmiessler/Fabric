import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { pathToFileURL } from 'node:url';
import { installPolyfills } from '../../exports/node/polyfills.js';
import { load_config } from '../config/index.js';
import { forked } from '../../utils/fork.js';

export default forked(import.meta.url, generate_fallback);

/**
 * @param {{
 *   manifest_path: string;
 *   env: Record<string, string>
 * }} opts
 */
async function generate_fallback({ manifest_path, env }) {
	/** @type {import('types').ValidatedKitConfig} */
	const config = (await load_config()).kit;

	installPolyfills();

	const server_root = join(config.outDir, 'output');

	/** @type {import('types').ServerInternalModule} */
	const { set_building } = await import(pathToFileURL(`${server_root}/server/internal.js`).href);

	/** @type {import('types').ServerModule} */
	const { Server } = await import(pathToFileURL(`${server_root}/server/index.js`).href);

	/** @type {import('@sveltejs/kit').SSRManifest} */
	const manifest = (await import(pathToFileURL(manifest_path).href)).manifest;

	set_building();

	const server = new Server(manifest);
	await server.init({ env });

	const response = await server.respond(new Request(config.prerender.origin + '/[fallback]'), {
		getClientAddress: () => {
			throw new Error('Cannot read clientAddress during prerendering');
		},
		prerendering: {
			fallback: true,
			dependencies: new Map(),
			remote_responses: new Map()
		},
		read: (file) => readFileSync(join(config.files.assets, file))
	});

	if (response.ok) {
		return await response.text();
	}

	throw new Error(`Could not create a fallback page — failed with status ${response.status}`);
}
