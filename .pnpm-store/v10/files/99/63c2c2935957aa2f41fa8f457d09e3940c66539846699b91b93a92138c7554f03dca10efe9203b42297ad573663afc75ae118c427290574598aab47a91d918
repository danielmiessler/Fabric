import fs from 'node:fs';
import path from 'node:path';
import { normalizePath } from 'vite';

/**
 * Adds transitive JS and CSS dependencies to the js and css inputs.
 * @param {import('vite').Manifest} manifest
 * @param {string} entry
 * @param {boolean} add_dynamic_css
 * @returns {import('types').AssetDependencies}
 */
export function find_deps(manifest, entry, add_dynamic_css) {
	/** @type {Set<string>} */
	const seen = new Set();

	/** @type {Set<string>} */
	const imports = new Set();

	/** @type {Set<string>} */
	const stylesheets = new Set();

	/** @type {Set<string>} */
	const imported_assets = new Set();

	/** @type {Map<string, { css: Set<string>; assets: Set<string> }>} */
	const stylesheet_map = new Map();

	/**
	 * @param {string} current
	 * @param {boolean} add_js
	 * @param {string} initial_importer
	 * @param {number} dynamic_import_depth
	 */
	function traverse(current, add_js, initial_importer, dynamic_import_depth) {
		if (seen.has(current)) return;
		seen.add(current);

		const { chunk } = resolve_symlinks(manifest, current);

		if (add_js) imports.add(chunk.file);

		if (chunk.assets) {
			chunk.assets.forEach(asset => imported_assets.add(asset));
		}

		if (chunk.css) {
			chunk.css.forEach((file) => stylesheets.add(file));
		}

		if (chunk.imports) {
			chunk.imports.forEach((file) => traverse(file, add_js, initial_importer, dynamic_import_depth));
		}

		if (!add_dynamic_css) return;

		if ((chunk.css || chunk.assets) && dynamic_import_depth <= 1) {
			// group files based on the initial importer because if a file is only ever
			// a transitive dependency, it doesn't have a suitable name we can map back to
			// the server manifest
			if (stylesheet_map.has(initial_importer)) {
				const { css, assets } = /** @type {{ css: Set<string>; assets: Set<string> }} */ (stylesheet_map.get(initial_importer));
				if (chunk.css) chunk.css.forEach((file) => css.add(file));
				if (chunk.assets) chunk.assets.forEach((file) => assets.add(file));
			} else {
				stylesheet_map.set(initial_importer, {
					css: new Set(chunk.css),
					assets: new Set(chunk.assets)
				});
			}
		}

		if (chunk.dynamicImports) {
			dynamic_import_depth++;
			chunk.dynamicImports.forEach((file) => {
				traverse(file, false, file, dynamic_import_depth);
			});
		}
	}

	const { chunk, file } = resolve_symlinks(manifest, entry);

	traverse(file, true, entry, 0);

	const assets = Array.from(imported_assets);

	return {
		assets,
		file: chunk.file,
		imports: Array.from(imports),
		stylesheets: Array.from(stylesheets),
		// TODO do we need this separately?
		fonts: filter_fonts(assets),
		stylesheet_map
	};
}

/**
 * @param {import('vite').Manifest} manifest
 * @param {string} file
 */
export function resolve_symlinks(manifest, file) {
	while (!manifest[file]) {
		const next = normalizePath(path.relative('.', fs.realpathSync(file)));
		if (next === file) throw new Error(`Could not find file "${file}" in Vite manifest`);
		file = next;
	}

	const chunk = manifest[file];

	return { chunk, file };
}

/**
 * @param {string[]} assets 
 * @returns {string[]}
 */
export function filter_fonts(assets) {
	return assets.filter((asset) => /\.(woff2?|ttf|otf)$/.test(asset));
}

const method_names = new Set((['GET', 'HEAD', 'PUT', 'POST', 'DELETE', 'PATCH', 'OPTIONS']));

// If we'd written this in TypeScript, it could be easy...
/**
 * @param {string} str
 * @returns {str is import('types').HttpMethod}
 */
export function is_http_method(str) {
	return method_names.has(str);
}

/**
 * @param {import('types').ValidatedKitConfig} config
 * @returns {string}
 */
export function assets_base(config) {
	return (config.paths.assets || config.paths.base || '.') + '/';
}
