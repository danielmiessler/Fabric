import path from 'node:path';
import { relative_path, resolve_entry } from '../../utils/filesystem.js';
import { s } from '../../utils/misc.js';
import { dedent, isSvelte5Plus, write_if_changed } from './utils.js';
import colors from 'kleur';

/**
 * Writes the client manifest to disk. The manifest is used to power the router. It contains the
 * list of routes and corresponding Svelte components (i.e. pages and layouts).
 * @param {import('types').ValidatedKitConfig} kit
 * @param {import('types').ManifestData} manifest_data
 * @param {string} output
 * @param {import('types').ServerMetadata['nodes']} [metadata] If this is omitted, we have to assume that all routes with a `+layout/page.server.js` file have a server load function
 */
export function write_client_manifest(kit, manifest_data, output, metadata) {
	const client_routing = kit.router.resolution === 'client';

	/**
	 * Creates a module that exports a `CSRPageNode`
	 * @param {import('types').PageNode} node
	 */
	function generate_node(node) {
		const declarations = [];

		if (node.universal) {
			declarations.push(
				`import * as universal from ${s(relative_path(`${output}/nodes`, node.universal))};`,
				'export { universal };'
			);
		}

		if (node.component) {
			declarations.push(
				`export { default as component } from ${s(
					relative_path(`${output}/nodes`, node.component)
				)};`
			);
		}

		return declarations.join('\n');
	}

	/** @type {Map<import('types').PageNode, number>} */
	const indices = new Map();
	const nodes = manifest_data.nodes
		.map((node, i) => {
			indices.set(node, i);

			write_if_changed(`${output}/nodes/${i}.js`, generate_node(node));
			return `() => import('./nodes/${i}')`;
		})
		// If route resolution happens on the server, we only need the root layout and root error page
		// upfront, the rest is loaded on demand as the user navigates the app
		.slice(0, client_routing ? manifest_data.nodes.length : 2)
		.join(',\n');

	const layouts_with_server_load = new Set();

	let dictionary = dedent`
		{
			${manifest_data.routes
				.map((route) => {
					if (route.page) {
						const errors = route.page.errors.slice(1).map((n) => n ?? '');
						const layouts = route.page.layouts.slice(1).map((n) => n ?? '');

						while (layouts.at(-1) === '') layouts.pop();
						while (errors.at(-1) === '') errors.pop();

						let leaf_has_server_load = false;
						if (route.leaf) {
							if (metadata) {
								const i = /** @type {number} */ (indices.get(route.leaf));

								leaf_has_server_load = metadata[i].has_server_load;
							} else if (route.leaf.server) {
								leaf_has_server_load = true;
							}
						}

						// Encode whether or not the route uses server data
						// using the ones' complement, to save space
						const array = [`${leaf_has_server_load ? '~' : ''}${route.page.leaf}`];

						// Encode whether or not the layout uses server data.
						// It's a different method compared to pages because layouts
						// are reused across pages, so we save space by doing it this way.
						route.page.layouts.forEach((layout) => {
							if (layout == undefined) return;

							let layout_has_server_load = false;

							if (metadata) {
								layout_has_server_load = metadata[layout].has_server_load;
							} else if (manifest_data.nodes[layout].server) {
								layout_has_server_load = true;
							}

							if (layout_has_server_load) {
								layouts_with_server_load.add(layout);
							}
						});

						// only include non-root layout/error nodes if they exist
						if (layouts.length > 0 || errors.length > 0) array.push(`[${layouts.join(',')}]`);
						if (errors.length > 0) array.push(`[${errors.join(',')}]`);

						return `${s(route.id)}: [${array.join(',')}]`;
					}
				})
				.filter(Boolean)
				.join(',\n')}
		}
	`;

	if (!client_routing) {
		dictionary = '{}';
		const root_layout = layouts_with_server_load.has(0);
		layouts_with_server_load.clear();
		if (root_layout) layouts_with_server_load.add(0);
	}

	const client_hooks_file = resolve_entry(kit.files.hooks.client);
	const universal_hooks_file = resolve_entry(kit.files.hooks.universal);

	const typo = resolve_entry('src/+hooks.client');
	if (typo) {
		console.log(
			colors
				.bold()
				.yellow(
					`Unexpected + prefix. Did you mean ${typo.split('/').at(-1)?.slice(1)}?` +
						` at ${path.resolve(typo)}`
				)
		);
	}

	// Stringified version of
	/** @type {import('../../runtime/client/types.js').SvelteKitApp} */
	write_if_changed(
		`${output}/app.js`,
		dedent`
			${
				client_hooks_file
					? `import * as client_hooks from '${relative_path(output, client_hooks_file)}';`
					: ''
			}
			${
				universal_hooks_file
					? `import * as universal_hooks from '${relative_path(output, universal_hooks_file)}';`
					: ''
			}

			${client_routing ? "export { matchers } from './matchers.js';" : 'export const matchers = {};'}

			export const nodes = [
				${nodes}
			];

			export const server_loads = [${[...layouts_with_server_load].join(',')}];

			export const dictionary = ${dictionary};

			export const hooks = {
				handleError: ${
					client_hooks_file ? 'client_hooks.handleError || ' : ''
				}(({ error }) => { console.error(error) }),
				${client_hooks_file ? 'init: client_hooks.init,' : ''}
				reroute: ${universal_hooks_file ? 'universal_hooks.reroute || ' : ''}(() => {}),
				transport: ${universal_hooks_file ? 'universal_hooks.transport || ' : ''}{}
			};

			export const decoders = Object.fromEntries(Object.entries(hooks.transport).map(([k, v]) => [k, v.decode]));

			export const hash = ${s(kit.router.type === 'hash')};

			export const decode = (type, value) => decoders[type](value);

			export { default as root } from '../root.${isSvelte5Plus() ? 'js' : 'svelte'}';
		`
	);

	if (client_routing) {
		// write matchers to a separate module so that we don't
		// need to worry about name conflicts
		const imports = [];
		const matchers = [];

		for (const key in manifest_data.matchers) {
			const src = manifest_data.matchers[key];

			imports.push(`import { match as ${key} } from ${s(relative_path(output, src))};`);
			matchers.push(key);
		}

		const module = imports.length
			? `${imports.join('\n')}\n\nexport const matchers = { ${matchers.join(', ')} };`
			: 'export const matchers = {};';

		write_if_changed(`${output}/matchers.js`, module);
	}
}
