import { find_deps } from '../../exports/vite/build/utils.js';

/**
 * Finds all the assets that are imported by server files associated with `routes`
 * @param {import('types').BuildData} build_data
 * @param {import('types').RouteData[]} routes
 */
export function find_server_assets(build_data, routes) {
	/**
	 * All nodes actually used in the routes definition (prerendered routes are omitted).
	 * Root layout/error is always included as they are needed for 404 and root errors.
	 * @type {Set<any>}
	 */
	const used_nodes = new Set([0, 1]);

	// TODO add hooks.server.js asset imports
	/** @type {Set<string>} */
	const server_assets = new Set();

	/** @param {string} id */
	function add_assets(id) {
		if (id in build_data.server_manifest) {
			const deps = find_deps(build_data.server_manifest, id, false);
			for (const asset of deps.assets) {
				server_assets.add(asset);
			}
		}
	}

	for (const route of routes) {
		if (route.page) {
			for (const i of route.page.layouts) used_nodes.add(i);
			for (const i of route.page.errors) used_nodes.add(i);
			used_nodes.add(route.page.leaf);
		}

		if (route.endpoint) {
			add_assets(route.endpoint.file);
		}
	}

	for (const n of used_nodes) {
		const node = build_data.manifest_data.nodes[n];
		if (node?.universal) add_assets(node.universal);
		if (node?.server) add_assets(node.server);
	}

	if (build_data.manifest_data.hooks.server) {
		add_assets(build_data.manifest_data.hooks.server);
	}

	return Array.from(server_assets);
}
