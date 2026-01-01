/** @import { Builder } from '@sveltejs/kit' */
/** @import { ResolvedConfig } from 'vite' */
/** @import { RouteDefinition } from '@sveltejs/kit' */
/** @import { RouteData, ValidatedConfig, BuildData, ServerMetadata, ServerMetadataRoute, Prerendered, PrerenderMap, Logger, RemoteChunk } from 'types' */
import colors from 'kleur';
import { createReadStream, createWriteStream, existsSync, statSync } from 'node:fs';
import { extname, resolve, join, dirname, relative } from 'node:path';
import { pipeline } from 'node:stream';
import { promisify } from 'node:util';
import zlib from 'node:zlib';
import { copy, rimraf, mkdirp, posixify } from '../../utils/filesystem.js';
import { generate_manifest } from '../generate_manifest/index.js';
import { get_route_segments } from '../../utils/routing.js';
import { get_env } from '../../exports/vite/utils.js';
import generate_fallback from '../postbuild/fallback.js';
import { write } from '../sync/utils.js';
import { list_files } from '../utils.js';
import { find_server_assets } from '../generate_manifest/find_server_assets.js';
import { reserved } from '../env.js';

const pipe = promisify(pipeline);
const extensions = ['.html', '.js', '.mjs', '.json', '.css', '.svg', '.xml', '.wasm'];

/**
 * Creates the Builder which is passed to adapters for building the application.
 * @param {{
 *   config: ValidatedConfig;
 *   build_data: BuildData;
 *   server_metadata: ServerMetadata;
 *   route_data: RouteData[];
 *   prerendered: Prerendered;
 *   prerender_map: PrerenderMap;
 *   log: Logger;
 *   vite_config: ResolvedConfig;
 *   remotes: RemoteChunk[]
 * }} opts
 * @returns {Builder}
 */
export function create_builder({
	config,
	build_data,
	server_metadata,
	route_data,
	prerendered,
	prerender_map,
	log,
	vite_config,
	remotes
}) {
	/** @type {Map<RouteDefinition, RouteData>} */
	const lookup = new Map();

	/**
	 * Rather than exposing the internal `RouteData` type, which is subject to change,
	 * we expose a stable type that adapters can use to group/filter routes
	 */
	const routes = route_data.map((route) => {
		const { config, methods, page, api } = /** @type {ServerMetadataRoute} */ (
			server_metadata.routes.get(route.id)
		);

		/** @type {RouteDefinition} */
		const facade = {
			id: route.id,
			api,
			page,
			segments: get_route_segments(route.id).map((segment) => ({
				dynamic: segment.includes('['),
				rest: segment.includes('[...'),
				content: segment
			})),
			pattern: route.pattern,
			prerender: prerender_map.get(route.id) ?? false,
			methods,
			config
		};

		lookup.set(facade, route);

		return facade;
	});

	return {
		log,
		rimraf,
		mkdirp,
		copy,

		config,
		prerendered,
		routes,

		async compress(directory) {
			if (!existsSync(directory)) {
				return;
			}

			const files = list_files(directory, (file) => extensions.includes(extname(file))).map(
				(file) => resolve(directory, file)
			);

			await Promise.all(
				files.flatMap((file) => [compress_file(file, 'gz'), compress_file(file, 'br')])
			);
		},

		async createEntries(fn) {
			const seen = new Set();

			for (let i = 0; i < route_data.length; i += 1) {
				const route = route_data[i];
				if (prerender_map.get(route.id) === true) continue;
				const { id, filter, complete } = fn(routes[i]);

				if (seen.has(id)) continue;
				seen.add(id);

				const group = [route];

				// figure out which lower priority routes should be considered fallbacks
				for (let j = i + 1; j < route_data.length; j += 1) {
					if (prerender_map.get(routes[j].id) === true) continue;
					if (filter(routes[j])) {
						group.push(route_data[j]);
					}
				}

				const filtered = new Set(group);

				// heuristic: if /foo/[bar] is included, /foo/[bar].json should
				// also be included, since the page likely needs the endpoint
				// TODO is this still necessary, given the new way of doing things?
				filtered.forEach((route) => {
					if (route.page) {
						const endpoint = route_data.find((candidate) => candidate.id === route.id + '.json');

						if (endpoint) {
							filtered.add(endpoint);
						}
					}
				});

				if (filtered.size > 0) {
					await complete({
						generateManifest: ({ relativePath }) =>
							generate_manifest({
								build_data,
								prerendered: [],
								relative_path: relativePath,
								routes: Array.from(filtered),
								remotes
							})
					});
				}
			}
		},

		findServerAssets(route_data) {
			return find_server_assets(
				build_data,
				route_data.map((route) => /** @type {import('types').RouteData} */ (lookup.get(route)))
			);
		},

		async generateFallback(dest) {
			const manifest_path = `${config.kit.outDir}/output/server/manifest-full.js`;
			const env = get_env(config.kit.env, vite_config.mode);

			const fallback = await generate_fallback({
				manifest_path,
				env: { ...env.private, ...env.public }
			});

			if (existsSync(dest)) {
				console.log(
					colors
						.bold()
						.yellow(
							`Overwriting ${dest} with fallback page. Consider using a different name for the fallback.`
						)
				);
			}

			write(dest, fallback);
		},

		generateEnvModule() {
			const dest = `${config.kit.outDir}/output/prerendered/dependencies/${config.kit.appDir}/env.js`;
			const env = get_env(config.kit.env, vite_config.mode);

			write(dest, `export const env=${JSON.stringify(env.public)}`);
		},

		generateManifest({ relativePath, routes: subset }) {
			return generate_manifest({
				build_data,
				prerendered: prerendered.paths,
				relative_path: relativePath,
				routes: subset
					? subset.map((route) => /** @type {import('types').RouteData} */ (lookup.get(route)))
					: route_data.filter((route) => prerender_map.get(route.id) !== true),
				remotes
			});
		},

		getBuildDirectory(name) {
			return `${config.kit.outDir}/${name}`;
		},

		getClientDirectory() {
			return `${config.kit.outDir}/output/client`;
		},

		getServerDirectory() {
			return `${config.kit.outDir}/output/server`;
		},

		getAppPath() {
			return build_data.app_path;
		},

		writeClient(dest) {
			return copy(`${config.kit.outDir}/output/client`, dest, {
				// avoid making vite build artefacts public
				filter: (basename) => basename !== '.vite'
			});
		},

		writePrerendered(dest) {
			const source = `${config.kit.outDir}/output/prerendered`;

			return [
				...copy(`${source}/pages`, dest),
				...copy(`${source}/dependencies`, dest),
				...copy(`${source}/data`, dest)
			];
		},

		writeServer(dest) {
			return copy(`${config.kit.outDir}/output/server`, dest);
		},

		hasServerInstrumentationFile() {
			return existsSync(`${config.kit.outDir}/output/server/instrumentation.server.js`);
		},

		instrument({
			entrypoint,
			instrumentation,
			start = join(dirname(entrypoint), 'start.js'),
			module = {
				exports: ['default']
			}
		}) {
			if (!existsSync(instrumentation)) {
				throw new Error(
					`Instrumentation file ${instrumentation} not found. This is probably a bug in your adapter.`
				);
			}
			if (!existsSync(entrypoint)) {
				throw new Error(
					`Entrypoint file ${entrypoint} not found. This is probably a bug in your adapter.`
				);
			}

			copy(entrypoint, start);
			if (existsSync(`${entrypoint}.map`)) {
				copy(`${entrypoint}.map`, `${start}.map`);
			}

			const relative_instrumentation = posixify(relative(dirname(entrypoint), instrumentation));
			const relative_start = posixify(relative(dirname(entrypoint), start));

			const facade =
				'generateText' in module
					? module.generateText({
							instrumentation: relative_instrumentation,
							start: relative_start
						})
					: create_instrumentation_facade({
							instrumentation: relative_instrumentation,
							start: relative_start,
							exports: module.exports
						});

			rimraf(entrypoint);
			write(entrypoint, facade);
		}
	};
}

/**
 * @param {string} file
 * @param {'gz' | 'br'} format
 */
async function compress_file(file, format = 'gz') {
	const compress =
		format == 'br'
			? zlib.createBrotliCompress({
					params: {
						[zlib.constants.BROTLI_PARAM_MODE]: zlib.constants.BROTLI_MODE_TEXT,
						[zlib.constants.BROTLI_PARAM_QUALITY]: zlib.constants.BROTLI_MAX_QUALITY,
						[zlib.constants.BROTLI_PARAM_SIZE_HINT]: statSync(file).size
					}
				})
			: zlib.createGzip({ level: zlib.constants.Z_BEST_COMPRESSION });

	const source = createReadStream(file);
	const destination = createWriteStream(`${file}.${format}`);

	await pipe(source, compress, destination);
}

/**
 * Given a list of exports, generate a facade that:
 * - Imports the instrumentation file
 * - Imports `exports` from the entrypoint (dynamically, if `tla` is true)
 * - Re-exports `exports` from the entrypoint
 *
 * `default` receives special treatment: It will be imported as `default` and exported with `export default`.
 *
 * @param {{ instrumentation: string; start: string; exports: string[] }} opts
 * @returns {string}
 */
function create_instrumentation_facade({ instrumentation, start, exports }) {
	const import_instrumentation = `import './${instrumentation}';`;

	let alias_index = 0;
	const aliases = new Map();

	for (const name of exports.filter((name) => reserved.has(name))) {
		/*
		 * you can do evil things like `export { c as class }`.
		 * in order to import these, you need to alias them, and then un-alias them when re-exporting
		 * this map will allow us to generate the following:
		 * import { class as _1 } from 'entrypoint';
		 * export { _1 as class };
		 */
		let alias = `_${alias_index++}`;
		while (exports.includes(alias)) {
			alias = `_${alias_index++}`;
		}

		aliases.set(name, alias);
	}

	const import_statements = [];
	const export_statements = [];

	for (const name of exports) {
		const alias = aliases.get(name);
		if (alias) {
			import_statements.push(`${name}: ${alias}`);
			export_statements.push(`${alias} as ${name}`);
		} else {
			import_statements.push(`${name}`);
			export_statements.push(`${name}`);
		}
	}

	const entrypoint_facade = [
		`const { ${import_statements.join(', ')} } = await import('./${start}');`,
		export_statements.length > 0 ? `export { ${export_statements.join(', ')} };` : ''
	]
		.filter(Boolean)
		.join('\n');

	return `${import_instrumentation}\n${entrypoint_facade}`;
}
