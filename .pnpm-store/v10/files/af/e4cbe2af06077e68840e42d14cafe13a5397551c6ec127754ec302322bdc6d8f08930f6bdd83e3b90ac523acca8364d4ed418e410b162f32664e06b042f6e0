import fs from 'node:fs';
import { mkdirp } from '../../../utils/filesystem.js';
import { filter_fonts, find_deps, resolve_symlinks } from './utils.js';
import { s } from '../../../utils/misc.js';
import { normalizePath } from 'vite';
import { basename, join } from 'node:path';
import { create_node_analyser } from '../static_analysis/index.js';


/**
 * @param {string} out
 * @param {import('types').ValidatedKitConfig} kit
 * @param {import('types').ManifestData} manifest_data
 * @param {import('vite').Manifest} server_manifest
 * @param {import('vite').Manifest | null} client_manifest
 * @param {import('vite').Rollup.OutputAsset[] | null} css
 * @param {import('types').RecursiveRequired<import('types').ValidatedConfig['kit']['output']>} output_config
 * @param {Map<string, Record<string, any> | null>} static_exports
 */
export async function build_server_nodes(out, kit, manifest_data, server_manifest, client_manifest, css, output_config, static_exports) {
	mkdirp(`${out}/server/nodes`);
	mkdirp(`${out}/server/stylesheets`);

	/** @type {Map<string, string>} */
	const stylesheet_lookup = new Map();

	if (css) {
		/** @type {Set<string>} */
		const client_stylesheets = new Set();
		for (const key in client_manifest) {
			client_manifest[key].css?.forEach((filename) => {
				client_stylesheets.add(filename);
			});
		}

		/** @type {Map<number, string[]>} */
		const server_stylesheets = new Map();
		manifest_data.nodes.forEach((node, i) => {
			if (!node.component || !server_manifest[node.component]) return;

			const { stylesheets } = find_deps(server_manifest, node.component, false);

			if (stylesheets.length) {
				server_stylesheets.set(i, stylesheets);
			}
		});

		for (const asset of css) {
			// ignore dynamically imported stylesheets since we don't need to inline those
			if (!client_stylesheets.has(asset.fileName) || asset.source.length >= kit.inlineStyleThreshold) {
				continue;
			}

			// We know that the names for entry points are numbers.
			const [index] = basename(asset.fileName).split('.');
			// There can also be other CSS files from shared components
			// for example, which we need to ignore here.
			if (isNaN(+index)) continue;

			const file = `${out}/server/stylesheets/${index}.js`;

			// we need to inline the server stylesheet instead of the client one
			// so that asset paths are correct on document load
			const filenames = server_stylesheets.get(+index);

			if (!filenames) {
				throw new Error('This should never happen, but if it does, it means we failed to find the server stylesheet for a node.');
			}

			const sources = filenames.map((filename) => {
				return fs.readFileSync(`${out}/server/${filename}`, 'utf-8');
			});
			fs.writeFileSync(file, `// ${filenames.join(', ')}\nexport default ${s(sources.join('\n'))};`);

			stylesheet_lookup.set(asset.fileName, index);
		}
	}

	const { get_page_options } = create_node_analyser({
		resolve: (server_node) => {
			// Windows needs the file:// protocol for absolute path dynamic imports
			return import(`file://${join(out, 'server', resolve_symlinks(server_manifest, server_node).chunk.file)}`);
		},
		static_exports
	});

	for (let i = 0; i < manifest_data.nodes.length; i++) {
		const node = manifest_data.nodes[i];

		/** @type {string[]} */
		const imports = [];

		// String representation of
		/** @type {import('types').SSRNode} */
		/** @type {string[]} */
		const exports = [`export const index = ${i};`];

		/** @type {string[]} */
		let imported = [];

		/** @type {string[]} */
		let stylesheets = [];

		/** @type {string[]} */
		let fonts = [];

		if (node.component && client_manifest) {
			exports.push(
				'let component_cache;',
				`export const component = async () => component_cache ??= (await import('../${
					resolve_symlinks(server_manifest, node.component).chunk.file
				}')).default;`
			);
		}

		if (node.universal) {
			const page_options = await get_page_options(node);
			if (!!page_options && page_options.ssr === false) {
				exports.push(`export const universal = ${s(page_options, null, 2)};`)
			} else {
				imports.push(
					`import * as universal from '../${resolve_symlinks(server_manifest, node.universal).chunk.file}';`
				);
				// TODO: when building for analysis, explain why the file was loaded on the server if we fail to load it
				exports.push('export { universal };');
			}
			exports.push(`export const universal_id = ${s(node.universal)};`);
		}

		if (node.server) {
			imports.push(
				`import * as server from '../${resolve_symlinks(server_manifest, node.server).chunk.file}';`
			);
			exports.push('export { server };');
			exports.push(`export const server_id = ${s(node.server)};`);
		}

		if (client_manifest && (node.universal || node.component) && output_config.bundleStrategy === 'split') {
			const entry_path = `${normalizePath(kit.outDir)}/generated/client-optimized/nodes/${i}.js`;
			const entry = find_deps(client_manifest, entry_path, true);

			// eagerly load stylesheets and fonts imported by the SSR-ed page to avoid FOUC.
			// If it is not used during SSR, it can be lazily loaded in the browser.
	
			/** @type {import('types').AssetDependencies | undefined} */
			let component;
			if (node.component) {
				component = find_deps(server_manifest, node.component, true);
			}

			/** @type {import('types').AssetDependencies | undefined} */
			let universal;
			if (node.universal) {
				universal = find_deps(server_manifest, node.universal, true);
			}

			/** @type {Set<string>} */
			const css_used_by_server = new Set();
			/** @type {Set<string>} */
			const assets_used_by_server = new Set();

			entry.stylesheet_map.forEach((value, key) => {
				// pages and layouts are named as node indexes in the client manifest
				// so we need to use the original filename when checking against the server manifest
				if (key === entry_path) {
					key = node.component ?? key;
				}

				if (component?.stylesheet_map.has(key) || universal?.stylesheet_map.has(key)) {
					value.css.forEach(file => css_used_by_server.add(file));
					value.assets.forEach(file => assets_used_by_server.add(file));
				}
			});

			imported = entry.imports;
			stylesheets = Array.from(css_used_by_server);
			fonts = filter_fonts(Array.from(assets_used_by_server));
		}

		exports.push(
			`export const imports = ${s(imported)};`,
			`export const stylesheets = ${s(stylesheets)};`,
			`export const fonts = ${s(fonts)};`
		);

		/** @type {string[]} */
		const styles = [];

		stylesheets.forEach((file) => {
			if (stylesheet_lookup.has(file)) {
				const index = stylesheet_lookup.get(file);
				const name = `stylesheet_${index}`;
				imports.push(`import ${name} from '../stylesheets/${index}.js';`);
				styles.push(`\t${s(file)}: ${name}`);
			}
		});

		if (styles.length > 0) {
			exports.push(`export const inline_styles = () => ({\n${styles.join(',\n')}\n});`);
		}

		fs.writeFileSync(
			`${out}/server/nodes/${i}.js`,
			`${imports.join('\n')}\n\n${exports.join('\n')}\n`
		);
	}
}
