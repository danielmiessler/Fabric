import fs from 'node:fs';
import * as vite from 'vite';
import { dedent } from '../../../core/sync/utils.js';
import { s } from '../../../utils/misc.js';
import { get_config_aliases, strip_virtual_prefix, get_env, normalize_id } from '../utils.js';
import { create_static_module } from '../../../core/env.js';
import { env_static_public, service_worker } from '../module_ids.js';

// @ts-ignore `vite.rolldownVersion` only exists in `rolldown-vite`
const isRolldown = !!vite.rolldownVersion;

/**
 * @param {string} out
 * @param {import('types').ValidatedKitConfig} kit
 * @param {import('vite').ResolvedConfig} vite_config
 * @param {import('types').ManifestData} manifest_data
 * @param {string} service_worker_entry_file
 * @param {import('types').Prerendered} prerendered
 * @param {import('vite').Manifest} client_manifest
 */
export async function build_service_worker(
	out,
	kit,
	vite_config,
	manifest_data,
	service_worker_entry_file,
	prerendered,
	client_manifest
) {
	const build = new Set();
	for (const key in client_manifest) {
		const { file, css = [], assets = [] } = client_manifest[key];
		build.add(file);
		css.forEach((file) => build.add(file));
		assets.forEach((file) => build.add(file));
	}

	// in a service worker, `location` is the location of the service worker itself,
	// which is guaranteed to be `<base>/service-worker.js`
	const base = "location.pathname.split('/').slice(0, -1).join('/')";

	const service_worker_code = dedent`
		export const base = /*@__PURE__*/ ${base};

		export const build = [
			${Array.from(build)
				.map((file) => `base + ${s(`/${file}`)}`)
				.join(',\n')}
		];

		export const files = [
			${manifest_data.assets
				.filter((asset) => kit.serviceWorker.files(asset.file))
				.map((asset) => `base + ${s(`/${asset.file}`)}`)
				.join(',\n')}
		];

		export const prerendered = [
			${prerendered.paths.map((path) => `base + ${s(path.replace(kit.paths.base, ''))}`).join(',\n')}
		];

		export const version = ${s(kit.version.name)};
	`;

	const env = get_env(kit.env, vite_config.mode);

	/**
	 * @type {import('vite').Plugin}
	 */
	const sw_virtual_modules = {
		name: 'service-worker-build-virtual-modules',
		resolveId(id) {
			if (id.startsWith('$env/') || id.startsWith('$app/') || id === '$service-worker') {
				// ids with :$ don't work with reverse proxies like nginx
				return `\0virtual:${id.substring(1)}`;
			}
		},

		load(id) {
			if (!id.startsWith('\0virtual:')) return;

			if (id === service_worker) {
				return service_worker_code;
			}

			if (id === env_static_public) {
				return create_static_module('$env/static/public', env.public);
			}

			const normalized_cwd = vite.normalizePath(process.cwd());
			const normalized_lib = vite.normalizePath(kit.files.lib);
			const relative = normalize_id(id, normalized_lib, normalized_cwd);
			const stripped = strip_virtual_prefix(relative);
			throw new Error(
				`Cannot import ${stripped} into service-worker code. Only the modules $service-worker and $env/static/public are available in service workers.`
			);
		}
	};

	await vite.build({
		build: {
			modulePreload: false,
			rollupOptions: {
				input: {
					'service-worker': service_worker_entry_file
				},
				output: {
					// .mjs so that esbuild doesn't incorrectly inject `export` https://github.com/vitejs/vite/issues/15379
					entryFileNames: `service-worker.${isRolldown ? 'js' : 'mjs'}`,
					assetFileNames: `${kit.appDir}/immutable/assets/[name].[hash][extname]`,
					inlineDynamicImports: true
				}
			},
			outDir: `${out}/client`,
			emptyOutDir: false,
			minify: vite_config.build.minify
		},
		configFile: false,
		define: vite_config.define,
		publicDir: false,
		plugins: [sw_virtual_modules],
		resolve: {
			alias: [...get_config_aliases(kit)]
		},
		experimental: {
			renderBuiltUrl(filename) {
				return {
					runtime: `new URL(${JSON.stringify(filename)}, location.href).pathname`
				};
			}
		}
	});

	// rename .mjs to .js to avoid incorrect MIME types with ancient webservers
	if (!isRolldown) {
		fs.renameSync(`${out}/client/service-worker.mjs`, `${out}/client/service-worker.js`);
	}
}
