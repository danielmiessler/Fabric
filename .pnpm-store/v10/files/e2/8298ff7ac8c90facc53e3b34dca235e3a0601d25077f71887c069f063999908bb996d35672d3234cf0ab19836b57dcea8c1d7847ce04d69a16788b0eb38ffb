import colors from 'kleur';
import { create_builder } from './builder.js';

/**
 * @param {import('types').ValidatedConfig} config
 * @param {import('types').BuildData} build_data
 * @param {import('types').ServerMetadata} server_metadata
 * @param {import('types').Prerendered} prerendered
 * @param {import('types').PrerenderMap} prerender_map
 * @param {import('types').Logger} log
 * @param {import('types').RemoteChunk[]} remotes
 * @param {import('vite').ResolvedConfig} vite_config
 */
export async function adapt(
	config,
	build_data,
	server_metadata,
	prerendered,
	prerender_map,
	log,
	remotes,
	vite_config
) {
	// This is only called when adapter is truthy, so the cast is safe
	const { name, adapt } = /** @type {import('@sveltejs/kit').Adapter} */ (config.kit.adapter);

	console.log(colors.bold().cyan(`\n> Using ${name}`));

	const builder = create_builder({
		config,
		build_data,
		server_metadata,
		route_data: build_data.manifest_data.routes.filter((route) => route.page || route.endpoint),
		prerendered,
		prerender_map,
		log,
		remotes,
		vite_config
	});

	await adapt(builder);

	log.success('done');
}
