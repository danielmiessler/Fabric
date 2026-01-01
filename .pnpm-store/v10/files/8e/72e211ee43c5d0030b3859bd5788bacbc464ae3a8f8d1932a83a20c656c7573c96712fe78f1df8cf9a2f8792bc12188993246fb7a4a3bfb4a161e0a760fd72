/** @import { RemoteChunk } from 'types' */
import fs from 'node:fs';
import path from 'node:path';
import * as mime from 'mrmime';
import { s } from '../../utils/misc.js';
import { get_mime_lookup } from '../utils.js';
import { resolve_symlinks } from '../../exports/vite/build/utils.js';
import { compact } from '../../utils/array.js';
import { join_relative } from '../../utils/filesystem.js';
import { dedent } from '../sync/utils.js';
import { find_server_assets } from './find_server_assets.js';
import { uneval } from 'devalue';

/**
 * Generates the data used to write the server-side manifest.js file. This data is used in the Vite
 * build process, to power routing, etc.
 * @param {{
 *   build_data: import('types').BuildData;
 *   prerendered: string[];
 *   relative_path: string;
 *   routes: import('types').RouteData[];
 *   remotes: RemoteChunk[];
 * }} opts
 */
export function generate_manifest({ build_data, prerendered, relative_path, routes, remotes }) {
	/**
	 * @type {Map<any, number>} The new index of each node in the filtered nodes array
	 */
	const reindexed = new Map();
	/**
	 * All nodes actually used in the routes definition (prerendered routes are omitted).
	 * Root layout/error is always included as they are needed for 404 and root errors.
	 * @type {Set<any>}
	 */
	const used_nodes = new Set([0, 1]);

	const server_assets = find_server_assets(build_data, routes);

	for (const route of routes) {
		if (route.page) {
			for (const i of route.page.layouts) used_nodes.add(i);
			for (const i of route.page.errors) used_nodes.add(i);
			used_nodes.add(route.page.leaf);
		}
	}

	const node_paths = compact(
		build_data.manifest_data.nodes.map((_, i) => {
			if (used_nodes.has(i)) {
				reindexed.set(i, reindexed.size);
				return join_relative(relative_path, `/nodes/${i}.js`);
			}
		})
	);

	/** @type {(path: string) => string} */
	const loader = (path) => `__memo(() => import('${path}'))`;

	const assets = build_data.manifest_data.assets.map((asset) => asset.file);
	if (build_data.service_worker) {
		assets.push(build_data.service_worker);
	}

	// In case of server-side route resolution, we need to include all matchers. Prerendered routes are not part
	// of the server manifest, and they could reference matchers that then would not be included.
	const matchers = new Set(
		build_data.client?.nodes ? Object.keys(build_data.manifest_data.matchers) : undefined
	);

	/** @param {Array<number | undefined>} indexes */
	function get_nodes(indexes) {
		const string = indexes.map((n) => reindexed.get(n) ?? '').join(',');

		// since JavaScript ignores trailing commas, we need to insert a dummy
		// comma so that the array has the correct length if the last item
		// is undefined
		return `[${string},]`;
	}

	const mime_types = get_mime_lookup(build_data.manifest_data);

	/** @type {Record<string, number>} */
	const files = {};
	for (const file of server_assets) {
		files[file] = fs.statSync(path.resolve(build_data.out_dir, 'server', file)).size;

		const ext = path.extname(file);
		mime_types[ext] ??= mime.lookup(ext) || '';
	}

	// prettier-ignore
	// String representation of
	/** @template {import('@sveltejs/kit').SSRManifest} T */
	const manifest_expr = dedent`
		{
			appDir: ${s(build_data.app_dir)},
			appPath: ${s(build_data.app_path)},
			assets: new Set(${s(assets)}),
			mimeTypes: ${s(mime_types)},
			_: {
				client: ${uneval(build_data.client)},
				nodes: [
					${(node_paths).map(loader).join(',\n')}
				],
				remotes: {
					${remotes.map((remote) => `'${remote.hash}': ${loader(join_relative(relative_path, `chunks/remote-${remote.hash}.js`))}`).join(',\n')}
				},
				routes: [
					${routes.map(route => {
						if (!route.page && !route.endpoint) return;

						route.params.forEach(param => {
							if (param.matcher) matchers.add(param.matcher);
						});

						return dedent`
							{
								id: ${s(route.id)},
								pattern: ${route.pattern},
								params: ${s(route.params)},
								page: ${route.page ? `{ layouts: ${get_nodes(route.page.layouts)}, errors: ${get_nodes(route.page.errors)}, leaf: ${reindexed.get(route.page.leaf)} }` : 'null'},
								endpoint: ${route.endpoint ? loader(join_relative(relative_path, resolve_symlinks(build_data.server_manifest, route.endpoint.file).chunk.file)) : 'null'}
							}
						`;
					}).filter(Boolean).join(',\n')}
				],
				prerendered_routes: new Set(${s(prerendered)}),
				matchers: async () => {
					${Array.from(
						matchers,
						type => `const { match: ${type} } = await import ('${(join_relative(relative_path, `/entries/matchers/${type}.js`))}')`
					).join('\n')}
					return { ${Array.from(matchers).join(', ')} };
				},
				server_assets: ${s(files)}
			}
		}
	`;

	// Memoize the loaders to prevent Node from doing unnecessary work
	// on every dynamic import call
	return dedent`
		(() => {
		function __memo(fn) {
			let value;
			return () => value ??= (value = fn());
		}

		return ${manifest_expr}
		})()
	`;
}
