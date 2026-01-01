import { join } from 'node:path';
import { pathToFileURL } from 'node:url';
import { validate_server_exports } from '../../utils/exports.js';
import { load_config } from '../config/index.js';
import { forked } from '../../utils/fork.js';
import { installPolyfills } from '../../exports/node/polyfills.js';
import { ENDPOINT_METHODS } from '../../constants.js';
import { filter_private_env, filter_public_env } from '../../utils/env.js';
import { has_server_load, resolve_route } from '../../utils/routing.js';
import { check_feature } from '../../utils/features.js';
import { createReadableStream } from '@sveltejs/kit/node';
import { PageNodes } from '../../utils/page_nodes.js';
import { build_server_nodes } from '../../exports/vite/build/build_server.js';

export default forked(import.meta.url, analyse);

/**
 * @param {{
 *   hash: boolean;
 *   manifest_path: string;
 *   manifest_data: import('types').ManifestData;
 *   server_manifest: import('vite').Manifest;
 *   tracked_features: Record<string, string[]>;
 *   env: Record<string, string>;
 *   out: string;
 *   output_config: import('types').RecursiveRequired<import('types').ValidatedConfig['kit']['output']>;
 * }} opts
 */
async function analyse({
	hash,
	manifest_path,
	manifest_data,
	server_manifest,
	tracked_features,
	env,
	out,
	output_config
}) {
	/** @type {import('@sveltejs/kit').SSRManifest} */
	const manifest = (await import(pathToFileURL(manifest_path).href)).manifest;

	/** @type {import('types').ValidatedKitConfig} */
	const config = (await load_config()).kit;

	const server_root = join(config.outDir, 'output');

	/** @type {import('types').ServerInternalModule} */
	const internal = await import(pathToFileURL(`${server_root}/server/internal.js`).href);

	installPolyfills();

	// configure `import { building } from '$app/environment'` —
	// essential we do this before analysing the code
	internal.set_building();

	// set env, in case it's used in initialisation
	const { publicPrefix: public_prefix, privatePrefix: private_prefix } = config.env;
	const private_env = filter_private_env(env, { public_prefix, private_prefix });
	const public_env = filter_public_env(env, { public_prefix, private_prefix });
	internal.set_private_env(private_env);
	internal.set_public_env(public_env);
	internal.set_safe_public_env(public_env);
	internal.set_manifest(manifest);
	internal.set_read_implementation((file) => createReadableStream(`${server_root}/server/${file}`));

	const static_exports = new Map();

	// first, build server nodes without the client manifest so we can analyse it
	await build_server_nodes(
		out,
		config,
		manifest_data,
		server_manifest,
		null,
		null,
		output_config,
		static_exports
	);

	/** @type {import('types').ServerMetadata} */
	const metadata = {
		nodes: [],
		routes: new Map()
	};

	const nodes = await Promise.all(manifest._.nodes.map((loader) => loader()));

	// analyse nodes
	for (const node of nodes) {
		if (hash && node.universal) {
			const options = Object.keys(node.universal).filter((o) => o !== 'load');
			if (options.length > 0) {
				throw new Error(
					`Page options are ignored when \`router.type === 'hash'\` (${node.universal_id} has ${options
						.filter((o) => o !== 'load')
						.map((o) => `'${o}'`)
						.join(', ')})`
				);
			}
		}

		metadata.nodes[node.index] = {
			has_server_load: has_server_load(node)
		};
	}

	// analyse routes
	for (const route of manifest._.routes) {
		const page =
			route.page &&
			analyse_page(
				route.page.layouts.map((n) => (n === undefined ? n : nodes[n])),
				nodes[route.page.leaf]
			);

		const endpoint = route.endpoint && analyse_endpoint(route, await route.endpoint());

		if (page?.prerender && endpoint?.prerender) {
			throw new Error(`Cannot prerender a route with both +page and +server files (${route.id})`);
		}

		if (page?.config && endpoint?.config) {
			for (const key in { ...page.config, ...endpoint.config }) {
				if (JSON.stringify(page.config[key]) !== JSON.stringify(endpoint.config[key])) {
					throw new Error(
						`Mismatched route config for ${route.id} — the +page and +server files must export the same config, if any`
					);
				}
			}
		}

		const route_config = page?.config ?? endpoint?.config ?? {};
		const prerender = page?.prerender ?? endpoint?.prerender;

		if (prerender !== true) {
			for (const feature of list_features(
				route,
				manifest_data,
				server_manifest,
				tracked_features
			)) {
				check_feature(route.id, route_config, feature, config.adapter);
			}
		}

		const page_methods = page?.methods ?? [];
		const api_methods = endpoint?.methods ?? [];
		const entries = page?.entries ?? endpoint?.entries;

		metadata.routes.set(route.id, {
			config: route_config,
			methods: Array.from(new Set([...page_methods, ...api_methods])),
			page: {
				methods: page_methods
			},
			api: {
				methods: api_methods
			},
			prerender,
			entries:
				entries && (await entries()).map((entry_object) => resolve_route(route.id, entry_object))
		});
	}

	return { metadata, static_exports };
}

/**
 * @param {import('types').SSRRoute} route
 * @param {import('types').SSREndpoint} mod
 */
function analyse_endpoint(route, mod) {
	validate_server_exports(mod, route.id);

	if (mod.prerender && (mod.POST || mod.PATCH || mod.PUT || mod.DELETE)) {
		throw new Error(
			`Cannot prerender a +server file with POST, PATCH, PUT, or DELETE (${route.id})`
		);
	}

	/** @type {Array<import('types').HttpMethod | '*'>} */
	const methods = [];

	for (const method of /** @type {import('types').HttpMethod[]} */ (ENDPOINT_METHODS)) {
		if (mod[method]) methods.push(method);
	}

	if (mod.fallback) {
		methods.push('*');
	}

	return {
		config: mod.config,
		entries: mod.entries,
		methods,
		prerender: mod.prerender ?? false
	};
}

/**
 * @param {Array<import('types').SSRNode | undefined>} layouts
 * @param {import('types').SSRNode} leaf
 */
function analyse_page(layouts, leaf) {
	/** @type {Array<'GET' | 'POST'>} */
	const methods = ['GET'];
	if (leaf.server?.actions) methods.push('POST');

	const nodes = new PageNodes([...layouts, leaf]);
	nodes.validate();

	return {
		config: nodes.get_config(),
		entries: leaf.universal?.entries ?? leaf.server?.entries,
		methods,
		prerender: nodes.prerender()
	};
}

/**
 * @param {import('types').SSRRoute} route
 * @param {import('types').ManifestData} manifest_data
 * @param {import('vite').Manifest} server_manifest
 * @param {Record<string, string[]>} tracked_features
 */
function list_features(route, manifest_data, server_manifest, tracked_features) {
	const features = new Set();

	const route_data = /** @type {import('types').RouteData} */ (
		manifest_data.routes.find((r) => r.id === route.id)
	);

	/** @param {string} id */
	function visit(id) {
		const chunk = server_manifest[id];
		if (!chunk) return;

		if (chunk.file in tracked_features) {
			for (const feature of tracked_features[chunk.file]) {
				features.add(feature);
			}
		}

		if (chunk.imports) {
			for (const id of chunk.imports) {
				visit(id);
			}
		}
	}

	let page_node = route_data?.leaf;
	while (page_node) {
		if (page_node.server) visit(page_node.server);
		page_node = page_node.parent ?? null;
	}

	if (route_data.endpoint) {
		visit(route_data.endpoint.file);
	}

	if (manifest_data.hooks.server) {
		// TODO if hooks.server.js imports `read`, it will be in the entry chunk
		// we don't currently account for that case
		visit(manifest_data.hooks.server);
	}

	return Array.from(features);
}
