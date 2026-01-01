import fs from 'node:fs';
import path from 'node:path';
import process from 'node:process';

import colors from 'kleur';

import { copy, mkdirp, posixify, read, resolve_entry, rimraf } from '../../utils/filesystem.js';
import { create_static_module, create_dynamic_module } from '../../core/env.js';
import * as sync from '../../core/sync/sync.js';
import { create_assets } from '../../core/sync/create_manifest_data/index.js';
import { runtime_directory, logger } from '../../core/utils.js';
import { load_config } from '../../core/config/index.js';
import { generate_manifest } from '../../core/generate_manifest/index.js';
import { build_server_nodes } from './build/build_server.js';
import { build_service_worker } from './build/build_service_worker.js';
import { assets_base, find_deps, resolve_symlinks } from './build/utils.js';
import { dev } from './dev/index.js';
import { is_illegal, module_guard } from './graph_analysis/index.js';
import { preview } from './preview/index.js';
import { get_config_aliases, get_env, normalize_id, strip_virtual_prefix } from './utils.js';
import { write_client_manifest } from '../../core/sync/write_client_manifest.js';
import prerender from '../../core/postbuild/prerender.js';
import analyse from '../../core/postbuild/analyse.js';
import { s } from '../../utils/misc.js';
import { hash } from '../../runtime/hash.js';
import { dedent, isSvelte5Plus } from '../../core/sync/utils.js';
import {
	env_dynamic_private,
	env_dynamic_public,
	env_static_private,
	env_static_public,
	service_worker,
	sveltekit_environment,
	sveltekit_paths,
	sveltekit_server
} from './module_ids.js';
import { import_peer } from '../../utils/import.js';
import { compact } from '../../utils/array.js';

const cwd = process.cwd();

/** @type {import('./types.js').EnforcedConfig} */
const enforced_config = {
	appType: true,
	base: true,
	build: {
		cssCodeSplit: true,
		emptyOutDir: true,
		lib: {
			entry: true,
			name: true,
			formats: true
		},
		manifest: true,
		outDir: true,
		rollupOptions: {
			input: true,
			output: {
				format: true,
				entryFileNames: true,
				chunkFileNames: true,
				assetFileNames: true
			},
			preserveEntrySignatures: true
		},
		ssr: true
	},
	publicDir: true,
	resolve: {
		alias: {
			$app: true,
			$lib: true,
			'$service-worker': true
		}
	},
	root: true
};

const options_regex = /(export\s+const\s+(prerender|csr|ssr|trailingSlash))\s*=/s;

/** @type {Set<string>} */
const warned = new Set();

/** @type {import('svelte/compiler').PreprocessorGroup} */
const warning_preprocessor = {
	script: ({ content, filename }) => {
		if (!filename) return;

		const basename = path.basename(filename);
		if (basename.startsWith('+page.') || basename.startsWith('+layout.')) {
			const match = content.match(options_regex);
			if (match) {
				const fixed = basename.replace('.svelte', '(.server).js/ts');

				const message =
					`\n${colors.bold().red(path.relative('.', filename))}\n` +
					`\`${match[1]}\` will be ignored — move it to ${fixed} instead. See https://svelte.dev/docs/kit/page-options for more information.`;

				if (!warned.has(message)) {
					console.log(message);
					warned.add(message);
				}
			}
		}
	},
	markup: ({ content, filename }) => {
		if (!filename) return;

		const basename = path.basename(filename);
		const has_children =
			content.includes('<slot') || (isSvelte5Plus() && content.includes('{@render'));

		if (basename.startsWith('+layout.') && !has_children) {
			const message =
				`\n${colors.bold().red(path.relative('.', filename))}\n` +
				`\`<slot />\`${isSvelte5Plus() ? ' or `{@render ...}` tag' : ''}` +
				' missing — inner content will not be rendered';

			if (!warned.has(message)) {
				console.log(message);
				warned.add(message);
			}
		}
	}
};

/**
 * Returns the SvelteKit Vite plugins.
 * @returns {Promise<import('vite').Plugin[]>}
 */
export async function sveltekit() {
	const svelte_config = await load_config();

	/** @type {import('@sveltejs/vite-plugin-svelte').Options['preprocess']} */
	let preprocess = svelte_config.preprocess;
	if (Array.isArray(preprocess)) {
		preprocess = [...preprocess, warning_preprocessor];
	} else if (preprocess) {
		preprocess = [preprocess, warning_preprocessor];
	} else {
		preprocess = warning_preprocessor;
	}

	/** @type {import('@sveltejs/vite-plugin-svelte').Options} */
	const vite_plugin_svelte_options = {
		configFile: false,
		extensions: svelte_config.extensions,
		preprocess,
		onwarn: svelte_config.onwarn,
		compilerOptions: {
			// @ts-ignore - ignore this property when running `pnpm check` against Svelte 5 in the ecosystem CI
			hydratable: isSvelte5Plus() ? undefined : true,
			...svelte_config.compilerOptions
		},
		...svelte_config.vitePlugin
	};

	const { svelte } = await import_peer('@sveltejs/vite-plugin-svelte');

	return [...svelte(vite_plugin_svelte_options), ...(await kit({ svelte_config }))];
}

// These variables live outside the `kit()` function because it is re-invoked by each Vite build

let secondary_build_started = false;

/** @type {import('types').ManifestData} */
let manifest_data;

/**
 * Returns the SvelteKit Vite plugin. Vite executes Rollup hooks as well as some of its own.
 * Background reading is available at:
 * - https://vitejs.dev/guide/api-plugin.html
 * - https://rollupjs.org/guide/en/#plugin-development
 *
 * You can get an idea of the lifecycle by looking at the flow charts here:
 * - https://rollupjs.org/guide/en/#build-hooks
 * - https://rollupjs.org/guide/en/#output-generation-hooks
 *
 * @param {{ svelte_config: import('types').ValidatedConfig }} options
 * @return {Promise<import('vite').Plugin[]>}
 */
async function kit({ svelte_config }) {
	const vite = await import_peer('vite');

	const { kit } = svelte_config;
	const out = `${kit.outDir}/output`;

	const version_hash = hash(kit.version.name);

	/** @type {import('vite').ResolvedConfig} */
	let vite_config;

	/** @type {import('vite').ConfigEnv} */
	let vite_config_env;

	/** @type {boolean} */
	let is_build;

	/** @type {{ public: Record<string, string>; private: Record<string, string> }} */
	let env;

	/** @type {() => Promise<void>} */
	let finalise;

	/** @type {import('vite').UserConfig} */
	let initial_config;

	const service_worker_entry_file = resolve_entry(kit.files.serviceWorker);
	const parsed_service_worker = path.parse(kit.files.serviceWorker);

	/**
	 * A map showing which features (such as `$app/server:read`) are defined
	 * in which chunks, so that we can later determine which routes use which features
	 * @type {Record<string, string[]>}
	 */
	const tracked_features = {};

	const sourcemapIgnoreList = /** @param {string} relative_path */ (relative_path) =>
		relative_path.includes('node_modules') || relative_path.includes(kit.outDir);

	/** @type {import('vite').Plugin} */
	const plugin_setup = {
		name: 'vite-plugin-sveltekit-setup',

		/**
		 * Build the SvelteKit-provided Vite config to be merged with the user's vite.config.js file.
		 * @see https://vitejs.dev/guide/api-plugin.html#config
		 */
		config(config, config_env) {
			initial_config = config;
			vite_config_env = config_env;
			is_build = config_env.command === 'build';

			env = get_env(kit.env, vite_config_env.mode);

			const allow = new Set([
				kit.files.lib,
				kit.files.routes,
				kit.outDir,
				path.resolve('src'), // TODO this isn't correct if user changed all his files to sth else than src (like in test/options)
				path.resolve('node_modules'),
				path.resolve(vite.searchForWorkspaceRoot(cwd), 'node_modules')
			]);

			// We can only add directories to the allow list, so we find out
			// if there's a client hooks file and pass its directory
			const client_hooks = resolve_entry(kit.files.hooks.client);
			if (client_hooks) allow.add(path.dirname(client_hooks));

			const generated = path.posix.join(kit.outDir, 'generated');

			// dev and preview config can be shared
			/** @type {import('vite').UserConfig} */
			const new_config = {
				resolve: {
					alias: [
						{ find: '__SERVER__', replacement: `${generated}/server` },
						{ find: '$app', replacement: `${runtime_directory}/app` },
						...get_config_aliases(kit)
					]
				},
				root: cwd,
				server: {
					cors: { preflightContinue: true },
					fs: {
						allow: [...allow]
					},
					sourcemapIgnoreList,
					watch: {
						ignored: [
							// Ignore all siblings of config.kit.outDir/generated
							`${posixify(kit.outDir)}/!(generated)`
						]
					}
				},
				preview: {
					cors: { preflightContinue: true }
				},
				optimizeDeps: {
					entries: [
						`${kit.files.routes}/**/+*.{svelte,js,ts}`,
						`!${kit.files.routes}/**/+*server.*`
					],
					exclude: [
						'@sveltejs/kit',
						// exclude kit features so that libraries using them work even when they are prebundled
						// this does not affect app code, just handling of imported libraries that use $app or $env
						'$app',
						'$env'
					]
				},
				ssr: {
					noExternal: [
						// This ensures that esm-env is inlined into the server output with the
						// export conditions resolved correctly through Vite. This prevents adapters
						// that bundle later on from resolving the export conditions incorrectly
						// and for example include browser-only code in the server output
						// because they for example use esbuild.build with `platform: 'browser'`
						'esm-env',
						// We need this for two reasons:
						// 1. Without this, `@sveltejs/kit` imports are kept as-is in the server output,
						//    and that causes modules and therefore classes like `Redirect` to be imported twice
						//    under different IDs, which breaks a bunch of stuff because of failing instanceof checks.
						// 2. Vitest bypasses Vite when loading external modules, so we bundle
						//    when it is detected to keep our virtual modules working.
						//    See https://github.com/sveltejs/kit/pull/9172
						//    and https://vitest.dev/config/#deps-registernodeloader
						'@sveltejs/kit'
					]
				}
			};

			if (is_build) {
				if (!new_config.build) new_config.build = {};
				new_config.build.ssr = !secondary_build_started;

				new_config.define = {
					__SVELTEKIT_ADAPTER_NAME__: s(kit.adapter?.name),
					__SVELTEKIT_APP_VERSION_FILE__: s(`${kit.appDir}/version.json`),
					__SVELTEKIT_APP_VERSION_POLL_INTERVAL__: s(kit.version.pollInterval),
					__SVELTEKIT_DEV__: 'false',
					__SVELTEKIT_EMBEDDED__: kit.embedded ? 'true' : 'false',
					__SVELTEKIT_CLIENT_ROUTING__: kit.router.resolution === 'client' ? 'true' : 'false'
				};

				if (!secondary_build_started) {
					manifest_data = sync.all(svelte_config, config_env.mode).manifest_data;
				}
			} else {
				new_config.define = {
					__SVELTEKIT_APP_VERSION_POLL_INTERVAL__: '0',
					__SVELTEKIT_DEV__: 'true',
					__SVELTEKIT_EMBEDDED__: kit.embedded ? 'true' : 'false',
					__SVELTEKIT_CLIENT_ROUTING__: kit.router.resolution === 'client' ? 'true' : 'false'
				};

				// These Kit dependencies are packaged as CommonJS, which means they must always be externalized.
				// Without this, the tests will still pass but `pnpm dev` will fail in projects that link `@sveltejs/kit`.
				/** @type {NonNullable<import('vite').UserConfig['ssr']>} */ (new_config.ssr).external = [
					'cookie',
					'set-cookie-parser'
				];
			}

			warn_overridden_config(config, new_config);

			return new_config;
		},

		/**
		 * Stores the final config.
		 */
		configResolved(config) {
			vite_config = config;
		}
	};

	/** @type {Map<string, string>} */
	const import_map = new Map();

	/** @type {import('vite').Plugin} */
	const plugin_virtual_modules = {
		name: 'vite-plugin-sveltekit-virtual-modules',

		resolveId(id, importer) {
			if (id === '__sveltekit/manifest') {
				return `${kit.outDir}/generated/client-optimized/app.js`;
			}

			// If importing from a service-worker, only allow $service-worker & $env/static/public, but none of the other virtual modules.
			// This check won't catch transitive imports, but it will warn when the import comes from a service-worker directly.
			// Transitive imports will be caught during the build.
			if (importer) {
				const parsed_importer = path.parse(importer);

				const importer_is_service_worker =
					parsed_importer.dir === parsed_service_worker.dir &&
					parsed_importer.name === parsed_service_worker.name;

				if (importer_is_service_worker && id !== '$service-worker' && id !== '$env/static/public') {
					const normalized_cwd = vite.normalizePath(cwd);
					const normalized_lib = vite.normalizePath(kit.files.lib);
					throw new Error(
						`Cannot import ${normalize_id(
							id,
							normalized_lib,
							normalized_cwd
						)} into service-worker code. Only the modules $service-worker and $env/static/public are available in service workers.`
					);
				}

				import_map.set(id, importer);
			}

			// treat $env/static/[public|private] as virtual
			if (id.startsWith('$env/') || id === '$service-worker') {
				// ids with :$ don't work with reverse proxies like nginx
				return `\0virtual:${id.substring(1)}`;
			}
			if (id.startsWith('__sveltekit/')) {
				return `\0virtual:${id}`;
			}
		},

		load(id, options) {
			const browser = !options?.ssr;

			const global = is_build
				? `globalThis.__sveltekit_${version_hash}`
				: 'globalThis.__sveltekit_dev';

			if (options?.ssr === false && process.env.TEST !== 'true') {
				const normalized_cwd = vite.normalizePath(cwd);
				const normalized_lib = vite.normalizePath(kit.files.lib);
				if (
					is_illegal(id, {
						cwd: normalized_cwd,
						node_modules: vite.normalizePath(path.resolve('node_modules')),
						server: vite.normalizePath(path.join(normalized_lib, 'server'))
					})
				) {
					const relative = normalize_id(id, normalized_lib, normalized_cwd);

					const illegal_module = strip_virtual_prefix(relative);

					const error_prefix = `Cannot import ${illegal_module} into client-side code. This could leak sensitive information.`;
					const error_suffix = `
Tips:
 - To resolve this error, ensure that no exports from ${illegal_module} are used, even transitively, in client-side code.
 - If you're only using the import as a type, change it to \`import type\`.
 - If you're not sure which module is causing this, try building your app -- it will create a more helpful error.`;

					if (import_map.has(illegal_module)) {
						const importer = path.relative(
							cwd,
							/** @type {string} */ (import_map.get(illegal_module))
						);
						throw new Error(`${error_prefix}\nImported by: ${importer}.${error_suffix}`);
					}

					throw new Error(`${error_prefix}${error_suffix}`);
				}
			}

			switch (id) {
				case env_static_private:
					return create_static_module('$env/static/private', env.private);

				case env_static_public:
					return create_static_module('$env/static/public', env.public);

				case env_dynamic_private:
					return create_dynamic_module(
						'private',
						vite_config_env.command === 'serve' ? env.private : undefined
					);

				case env_dynamic_public:
					// populate `$env/dynamic/public` from `window`
					if (browser) {
						return `export const env = ${global}.env;`;
					}

					return create_dynamic_module(
						'public',
						vite_config_env.command === 'serve' ? env.public : undefined
					);

				case service_worker:
					return create_service_worker_module(svelte_config);

				// for internal use only. it's published as $app/paths externally
				// we use this alias so that we won't collide with user aliases
				case sveltekit_paths: {
					const { assets, base } = svelte_config.kit.paths;

					// use the values defined in `global`, but fall back to hard-coded values
					// for the sake of things like Vitest which may import this module
					// outside the context of a page
					if (browser) {
						return dedent`
							export const base = ${global}?.base ?? ${s(base)};
							export const assets = ${global}?.assets ?? ${assets ? s(assets) : 'base'};
							export const app_dir = ${s(kit.appDir)};
						`;
					}

					return dedent`
						export let base = ${s(base)};
						export let assets = ${assets ? s(assets) : 'base'};
						export const app_dir = ${s(kit.appDir)};

						export const relative = ${svelte_config.kit.paths.relative};

						const initial = { base, assets };

						export function override(paths) {
							base = paths.base;
							assets = paths.assets;
						}

						export function reset() {
							base = initial.base;
							assets = initial.assets;
						}

						/** @param {string} path */
						export function set_assets(path) {
							assets = initial.assets = path;
						}
					`;
				}

				case sveltekit_environment: {
					const { version } = svelte_config.kit;

					return dedent`
						export const version = ${s(version.name)};
						export let building = false;
						export let prerendering = false;

						export function set_building() {
							building = true;
						}

						export function set_prerendering() {
							prerendering = true;
						}
					`;
				}

				case sveltekit_server: {
					return dedent`
						export let read_implementation = null;

						export let manifest = null;

						export function set_read_implementation(fn) {
							read_implementation = fn;
						}

						export function set_manifest(_) {
							manifest = _;
						}
					`;
				}
			}
		}
	};

	/**
	 * Ensures that client-side code can't accidentally import server-side code,
	 * whether in `*.server.js` files, `$app/server`, `$lib/server`, or `$env/[static|dynamic]/private`
	 * @type {import('vite').Plugin}
	 */
	const plugin_guard = {
		name: 'vite-plugin-sveltekit-guard',

		writeBundle: {
			sequential: true,
			handler(_options) {
				if (vite_config.build.ssr) return;

				const guard = module_guard(this, {
					cwd: vite.normalizePath(process.cwd()),
					lib: vite.normalizePath(kit.files.lib)
				});

				manifest_data.nodes.forEach((_node, i) => {
					const id = vite.normalizePath(
						path.resolve(kit.outDir, `generated/client-optimized/nodes/${i}.js`)
					);

					guard.check(id);
				});
			}
		}
	};

	/** @type {import('vite').Plugin} */
	const plugin_compile = {
		name: 'vite-plugin-sveltekit-compile',

		/**
		 * Build the SvelteKit-provided Vite config to be merged with the user's vite.config.js file.
		 * @see https://vitejs.dev/guide/api-plugin.html#config
		 */
		config(config) {
			/** @type {import('vite').UserConfig} */
			let new_config;

			if (is_build) {
				const ssr = /** @type {boolean} */ (config.build?.ssr);
				const prefix = `${kit.appDir}/immutable`;

				/** @type {Record<string, string>} */
				const input = {};

				if (ssr) {
					input.index = `${runtime_directory}/server/index.js`;
					input.internal = `${kit.outDir}/generated/server/internal.js`;

					// add entry points for every endpoint...
					manifest_data.routes.forEach((route) => {
						if (route.endpoint) {
							const resolved = path.resolve(route.endpoint.file);
							const relative = decodeURIComponent(path.relative(kit.files.routes, resolved));
							const name = posixify(path.join('entries/endpoints', relative.replace(/\.js$/, '')));
							input[name] = resolved;
						}
					});

					// ...and every component used by pages...
					manifest_data.nodes.forEach((node) => {
						for (const file of [node.component, node.universal, node.server]) {
							if (file) {
								const resolved = path.resolve(file);
								const relative = decodeURIComponent(path.relative(kit.files.routes, resolved));

								const name = relative.startsWith('..')
									? posixify(path.join('entries/fallbacks', path.basename(file)))
									: posixify(path.join('entries/pages', relative.replace(/\.js$/, '')));
								input[name] = resolved;
							}
						}
					});

					// ...and every matcher
					Object.entries(manifest_data.matchers).forEach(([key, file]) => {
						const name = posixify(path.join('entries/matchers', key));
						input[name] = path.resolve(file);
					});
				} else if (svelte_config.kit.output.bundleStrategy !== 'split') {
					input['bundle'] = `${runtime_directory}/client/bundle.js`;
				} else {
					input['entry/start'] = `${runtime_directory}/client/entry.js`;
					input['entry/app'] = `${kit.outDir}/generated/client-optimized/app.js`;
					manifest_data.nodes.forEach((node, i) => {
						if (node.component || node.universal) {
							input[`nodes/${i}`] = `${kit.outDir}/generated/client-optimized/nodes/${i}.js`;
						}
					});
				}

				// see the kit.output.preloadStrategy option for details on why we have multiple options here
				const ext = kit.output.preloadStrategy === 'preload-mjs' ? 'mjs' : 'js';

				// We could always use a relative asset base path here, but it's better for performance not to.
				// E.g. Vite generates `new URL('/asset.png', import.meta).href` for a relative path vs just '/asset.png'.
				// That's larger and takes longer to run and also causes an HTML diff between SSR and client
				// causing us to do a more expensive hydration check.
				const client_base =
					kit.paths.relative !== false || kit.paths.assets ? './' : kit.paths.base || '/';

				const inline = !ssr && svelte_config.kit.output.bundleStrategy === 'inline';
				const split = ssr || svelte_config.kit.output.bundleStrategy === 'split';

				new_config = {
					base: ssr ? assets_base(kit) : client_base,
					build: {
						copyPublicDir: !ssr,
						cssCodeSplit: svelte_config.kit.output.bundleStrategy !== 'inline',
						cssMinify: initial_config.build?.minify == null ? true : !!initial_config.build.minify,
						// don't use the default name to avoid collisions with 'static/manifest.json'
						manifest: '.vite/manifest.json', // TODO: remove this after bumping peer dep to vite 5
						outDir: `${out}/${ssr ? 'server' : 'client'}`,
						rollupOptions: {
							input: inline ? input['bundle'] : input,
							output: {
								format: inline ? 'iife' : 'esm',
								name: `__sveltekit_${version_hash}.app`,
								entryFileNames: ssr ? '[name].js' : `${prefix}/[name].[hash].${ext}`,
								chunkFileNames: ssr ? 'chunks/[name].js' : `${prefix}/chunks/[hash].${ext}`,
								assetFileNames: `${prefix}/assets/[name].[hash][extname]`,
								hoistTransitiveImports: false,
								sourcemapIgnoreList,
								manualChunks: split ? undefined : () => 'bundle',
								inlineDynamicImports: false
							},
							preserveEntrySignatures: 'strict',
							onwarn(warning, handler) {
								if (
									warning.code === 'MISSING_EXPORT' &&
									warning.id === `${kit.outDir}/generated/client-optimized/app.js`
								) {
									// ignore e.g. undefined `handleError` hook when
									// referencing `client_hooks.handleError`
									return;
								}

								handler(warning);
							}
						},
						ssrEmitAssets: true,
						target: ssr ? 'node18.13' : undefined
					},
					publicDir: kit.files.assets,
					worker: {
						rollupOptions: {
							output: {
								entryFileNames: `${prefix}/workers/[name]-[hash].js`,
								chunkFileNames: `${prefix}/workers/chunks/[hash].js`,
								assetFileNames: `${prefix}/workers/assets/[name]-[hash][extname]`,
								hoistTransitiveImports: false
							}
						}
					}
				};
			} else {
				new_config = {
					appType: 'custom',
					base: kit.paths.base,
					build: {
						rollupOptions: {
							// Vite dependency crawler needs an explicit JS entry point
							// eventhough server otherwise works without it
							input: `${runtime_directory}/client/entry.js`
						}
					},
					publicDir: kit.files.assets
				};
			}

			warn_overridden_config(config, new_config);

			return new_config;
		},

		/**
		 * Adds the SvelteKit middleware to do SSR in dev mode.
		 * @see https://vitejs.dev/guide/api-plugin.html#configureserver
		 */
		async configureServer(vite) {
			return await dev(vite, vite_config, svelte_config);
		},

		/**
		 * Adds the SvelteKit middleware to do SSR in preview mode.
		 * @see https://vitejs.dev/guide/api-plugin.html#configurepreviewserver
		 */
		configurePreviewServer(vite) {
			return preview(vite, vite_config, svelte_config);
		},

		/**
		 * Clears the output directories.
		 */
		buildStart() {
			if (secondary_build_started) return;

			if (is_build) {
				if (!vite_config.build.watch) {
					rimraf(out);
				}
				mkdirp(out);
			}
		},

		renderChunk(code, chunk) {
			if (code.includes('__SVELTEKIT_TRACK__')) {
				return {
					code: code.replace(/__SVELTEKIT_TRACK__\('(.+?)'\)/g, (_, label) => {
						(tracked_features[chunk.name + '.js'] ??= []).push(label);
						// put extra whitespace at the end of the comment to preserve the source size and avoid interfering with source maps
						return `/* track ${label}            */`;
					}),
					map: null // TODO we may need to generate a sourcemap in future
				};
			}
		},

		generateBundle() {
			if (vite_config.build.ssr) return;

			this.emitFile({
				type: 'asset',
				fileName: `${kit.appDir}/version.json`,
				source: s({ version: kit.version.name })
			});
		},

		/**
		 * Vite builds a single bundle. We need three bundles: client, server, and service worker.
		 * The user's package.json scripts will invoke the Vite CLI to execute the server build. We
		 * then use this hook to kick off builds for the client and service worker.
		 */
		writeBundle: {
			sequential: true,
			async handler(_options) {
				if (secondary_build_started) return; // only run this once

				const verbose = vite_config.logLevel === 'info';
				const log = logger({ verbose });

				/** @type {import('vite').Manifest} */
				const server_manifest = JSON.parse(read(`${out}/server/${vite_config.build.manifest}`));

				/** @type {import('types').BuildData} */
				const build_data = {
					app_dir: kit.appDir,
					app_path: `${kit.paths.base.slice(1)}${kit.paths.base ? '/' : ''}${kit.appDir}`,
					manifest_data,
					out_dir: out,
					service_worker: service_worker_entry_file ? 'service-worker.js' : null, // TODO make file configurable?
					client: null,
					server_manifest
				};

				const manifest_path = `${out}/server/manifest-full.js`;
				fs.writeFileSync(
					manifest_path,
					`export const manifest = ${generate_manifest({
						build_data,
						prerendered: [],
						relative_path: '.',
						routes: manifest_data.routes
					})};\n`
				);

				log.info('Analysing routes');

				const { metadata, static_exports } = await analyse({
					hash: kit.router.type === 'hash',
					manifest_path,
					manifest_data,
					server_manifest,
					tracked_features,
					env: { ...env.private, ...env.public },
					out,
					output_config: svelte_config.output
				});

				log.info('Building app');

				// create client build
				write_client_manifest(
					kit,
					manifest_data,
					`${kit.outDir}/generated/client-optimized`,
					metadata.nodes
				);

				secondary_build_started = true;

				const { output } = /** @type {import('vite').Rollup.RollupOutput} */ (
					await vite.build({
						configFile: vite_config.configFile,
						// CLI args
						mode: vite_config_env.mode,
						logLevel: vite_config.logLevel,
						clearScreen: vite_config.clearScreen,
						build: {
							minify: initial_config.build?.minify,
							assetsInlineLimit: vite_config.build.assetsInlineLimit,
							sourcemap: vite_config.build.sourcemap
						},
						optimizeDeps: {
							force: vite_config.optimizeDeps.force
						}
					})
				);

				copy(
					`${out}/server/${kit.appDir}/immutable/assets`,
					`${out}/client/${kit.appDir}/immutable/assets`
				);

				/** @type {import('vite').Manifest} */
				const client_manifest = JSON.parse(read(`${out}/client/${vite_config.build.manifest}`));

				/**
				 * @param {string} entry
				 * @param {boolean} [add_dynamic_css]
				 */
				const deps_of = (entry, add_dynamic_css = false) =>
					find_deps(client_manifest, posixify(path.relative('.', entry)), add_dynamic_css);

				if (svelte_config.kit.output.bundleStrategy === 'split') {
					const start = deps_of(`${runtime_directory}/client/entry.js`);
					const app = deps_of(`${kit.outDir}/generated/client-optimized/app.js`);

					build_data.client = {
						start: start.file,
						app: app.file,
						imports: [...start.imports, ...app.imports],
						stylesheets: [...start.stylesheets, ...app.stylesheets],
						fonts: [...start.fonts, ...app.fonts],
						uses_env_dynamic_public: output.some(
							(chunk) => chunk.type === 'chunk' && chunk.modules[env_dynamic_public]
						)
					};

					// In case of server-side route resolution, we create a purpose-built route manifest that is
					// similar to that on the client, with as much information computed upfront so that we
					// don't need to include any code of the actual routes in the server bundle.
					if (svelte_config.kit.router.resolution === 'server') {
						const nodes = manifest_data.nodes.map((node, i) => {
							if (node.component || node.universal) {
								const entry = `${kit.outDir}/generated/client-optimized/nodes/${i}.js`;
								const deps = deps_of(entry, true);
								const file = resolve_symlinks(
									client_manifest,
									`${kit.outDir}/generated/client-optimized/nodes/${i}.js`
								).chunk.file;

								return { file, css: deps.stylesheets };
							}
						});
						build_data.client.nodes = nodes.map((node) => node?.file);
						build_data.client.css = nodes.map((node) => node?.css);

						build_data.client.routes = compact(
							manifest_data.routes.map((route) => {
								if (!route.page) return;

								return {
									id: route.id,
									pattern: route.pattern,
									params: route.params,
									layouts: route.page.layouts.map((l) =>
										l !== undefined ? [metadata.nodes[l].has_server_load, l] : undefined
									),
									errors: route.page.errors,
									leaf: [metadata.nodes[route.page.leaf].has_server_load, route.page.leaf]
								};
							})
						);
					}
				} else {
					const start = deps_of(`${runtime_directory}/client/bundle.js`);

					build_data.client = {
						start: start.file,
						imports: start.imports,
						stylesheets: start.stylesheets,
						fonts: start.fonts,
						uses_env_dynamic_public: output.some(
							(chunk) => chunk.type === 'chunk' && chunk.modules[env_dynamic_public]
						)
					};

					if (svelte_config.kit.output.bundleStrategy === 'inline') {
						const style = /** @type {import('rollup').OutputAsset} */ (
							output.find(
								(chunk) =>
									chunk.type === 'asset' &&
									chunk.names.length === 1 &&
									chunk.names[0] === 'style.css'
							)
						);

						build_data.client.inline = {
							script: read(`${out}/client/${start.file}`),
							style: /** @type {string | undefined} */ (style?.source)
						};
					}
				}

				const css = output.filter(
					/** @type {(value: any) => value is import('vite').Rollup.OutputAsset} */
					(value) => value.type === 'asset' && value.fileName.endsWith('.css')
				);

				// regenerate manifest now that we have client entry...
				fs.writeFileSync(
					manifest_path,
					`export const manifest = ${generate_manifest({
						build_data,
						prerendered: [],
						relative_path: '.',
						routes: manifest_data.routes
					})};\n`
				);

				// regenerate nodes with the client manifest...
				await build_server_nodes(
					out,
					kit,
					manifest_data,
					server_manifest,
					client_manifest,
					css,
					svelte_config.kit.output,
					static_exports
				);

				// ...and prerender
				const { prerendered, prerender_map } = await prerender({
					hash: kit.router.type === 'hash',
					out,
					manifest_path,
					metadata,
					verbose,
					env: { ...env.private, ...env.public }
				});

				// generate a new manifest that doesn't include prerendered pages
				fs.writeFileSync(
					`${out}/server/manifest.js`,
					`export const manifest = ${generate_manifest({
						build_data,
						prerendered: prerendered.paths,
						relative_path: '.',
						routes: manifest_data.routes.filter((route) => prerender_map.get(route.id) !== true)
					})};\n`
				);

				if (service_worker_entry_file) {
					if (kit.paths.assets) {
						throw new Error('Cannot use service worker alongside config.kit.paths.assets');
					}

					log.info('Building service worker');

					await build_service_worker(
						out,
						kit,
						{
							...vite_config,
							build: {
								...vite_config.build,
								minify: initial_config.build?.minify ?? 'esbuild'
							}
						},
						manifest_data,
						service_worker_entry_file,
						prerendered,
						client_manifest
					);
				}

				// we need to defer this to closeBundle, so that adapters copy files
				// created by other Vite plugins
				finalise = async () => {
					console.log(
						`\nRun ${colors
							.bold()
							.cyan('npm run preview')} to preview your production build locally.`
					);

					if (kit.adapter) {
						const { adapt } = await import('../../core/adapt/index.js');
						await adapt(
							svelte_config,
							build_data,
							metadata,
							prerendered,
							prerender_map,
							log,
							vite_config
						);
					} else {
						console.log(colors.bold().yellow('\nNo adapter specified'));

						const link = colors.bold().cyan('https://svelte.dev/docs/kit/adapters');
						console.log(
							`See ${link} to learn how to configure your app to run on the platform of your choosing`
						);
					}

					secondary_build_started = false;
				};
			}
		},

		/**
		 * Runs the adapter.
		 */
		closeBundle: {
			sequential: true,
			async handler() {
				if (!vite_config.build.ssr) return;
				await finalise?.();
			}
		}
	};

	return [plugin_setup, plugin_virtual_modules, plugin_guard, plugin_compile];
}

/**
 * @param {Record<string, any>} config
 * @param {Record<string, any>} resolved_config
 */
function warn_overridden_config(config, resolved_config) {
	const overridden = find_overridden_config(config, resolved_config, enforced_config, '', []);

	if (overridden.length > 0) {
		console.error(
			colors.bold().red('The following Vite config options will be overridden by SvelteKit:') +
				overridden.map((key) => `\n  - ${key}`).join('')
		);
	}
}

/**
 * @param {Record<string, any>} config
 * @param {Record<string, any>} resolved_config
 * @param {import('./types.js').EnforcedConfig} enforced_config
 * @param {string} path
 * @param {string[]} out used locally to compute the return value
 */
function find_overridden_config(config, resolved_config, enforced_config, path, out) {
	if (config == null || resolved_config == null) {
		return out;
	}

	for (const key in enforced_config) {
		if (typeof config === 'object' && key in config && key in resolved_config) {
			const enforced = enforced_config[key];

			if (enforced === true) {
				if (config[key] !== resolved_config[key]) {
					out.push(path + key);
				}
			} else {
				find_overridden_config(config[key], resolved_config[key], enforced, path + key + '.', out);
			}
		}
	}
	return out;
}

/**
 * @param {import('types').ValidatedConfig} config
 */
const create_service_worker_module = (config) => dedent`
	if (typeof self === 'undefined' || self instanceof ServiceWorkerGlobalScope === false) {
		throw new Error('This module can only be imported inside a service worker');
	}

	export const build = [];
	export const files = [
		${create_assets(config)
			.filter((asset) => config.kit.serviceWorker.files(asset.file))
			.map((asset) => `${s(`${config.kit.paths.base}/${asset.file}`)}`)
			.join(',\n')}
	];
	export const prerendered = [];
	export const version = ${s(config.kit.version.name)};
`;
