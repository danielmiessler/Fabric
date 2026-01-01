import fs from 'node:fs';
import path from 'node:path';
import process from 'node:process';
import { URL } from 'node:url';
import { AsyncLocalStorage } from 'node:async_hooks';
import colors from 'kleur';
import sirv from 'sirv';
import { isCSSRequest, loadEnv, buildErrorMessage } from 'vite';
import { createReadableStream, getRequest, setResponse } from '../../../exports/node/index.js';
import { installPolyfills } from '../../../exports/node/polyfills.js';
import { coalesce_to_error } from '../../../utils/error.js';
import { from_fs, posixify, resolve_entry, to_fs } from '../../../utils/filesystem.js';
import { load_error_page } from '../../../core/config/index.js';
import { SVELTE_KIT_ASSETS } from '../../../constants.js';
import * as sync from '../../../core/sync/sync.js';
import { get_mime_lookup, runtime_base } from '../../../core/utils.js';
import { compact } from '../../../utils/array.js';
import { not_found } from '../utils.js';
import { SCHEME } from '../../../utils/url.js';
import { check_feature } from '../../../utils/features.js';
import { escape_html } from '../../../utils/escape.js';
import { create_node_analyser } from '../static_analysis/index.js';

const cwd = process.cwd();
// vite-specifc queries that we should skip handling for css urls
const vite_css_query_regex = /(?:\?|&)(?:raw|url|inline)(?:&|$)/;

/**
 * @param {import('vite').ViteDevServer} vite
 * @param {import('vite').ResolvedConfig} vite_config
 * @param {import('types').ValidatedConfig} svelte_config
 * @param {() => Array<{ hash: string, file: string }>} get_remotes
 * @return {Promise<Promise<() => void>>}
 */
export async function dev(vite, vite_config, svelte_config, get_remotes) {
	installPolyfills();

	const async_local_storage = new AsyncLocalStorage();

	globalThis.__SVELTEKIT_TRACK__ = (label) => {
		const context = async_local_storage.getStore();
		if (!context || context.prerender === true) return;

		check_feature(context.event.route.id, context.config, label, svelte_config.kit.adapter);
	};

	const fetch = globalThis.fetch;
	globalThis.fetch = (info, init) => {
		if (typeof info === 'string' && !SCHEME.test(info)) {
			throw new Error(
				`Cannot use relative URL (${info}) with global fetch — use \`event.fetch\` instead: https://svelte.dev/docs/kit/web-standards#fetch-apis`
			);
		}

		return fetch(info, init);
	};

	sync.init(svelte_config, vite_config.mode);

	/** @type {import('types').ManifestData} */
	let manifest_data;
	/** @type {import('@sveltejs/kit').SSRManifest} */
	let manifest;

	/** @type {Error | null} */
	let manifest_error = null;

	/** @param {string} url */
	async function loud_ssr_load_module(url) {
		try {
			return await vite.ssrLoadModule(url, { fixStacktrace: true });
		} catch (/** @type {any} */ err) {
			const msg = buildErrorMessage(err, [colors.red(`Internal server error: ${err.message}`)]);

			if (!vite.config.logger.hasErrorLogged(err)) {
				vite.config.logger.error(msg, { error: err });
			}

			vite.ws.send({
				type: 'error',
				err: {
					...err,
					// these properties are non-enumerable and will
					// not be serialized unless we explicitly include them
					message: err.message,
					stack: err.stack
				}
			});

			throw err;
		}
	}

	/** @param {string} id */
	async function resolve(id) {
		const url = id.startsWith('..') ? to_fs(path.posix.resolve(id)) : `/${id}`;

		const module = await loud_ssr_load_module(url);

		const module_node = await vite.moduleGraph.getModuleByUrl(url);
		if (!module_node) throw new Error(`Could not find node for ${url}`);

		return { module, module_node, url };
	}

	/** @type {(file: string) => void} */
	let invalidate_page_options;

	function update_manifest() {
		try {
			({ manifest_data } = sync.create(svelte_config));

			if (manifest_error) {
				manifest_error = null;
				vite.ws.send({ type: 'full-reload' });
			}
		} catch (error) {
			manifest_error = /** @type {Error} */ (error);

			console.error(colors.bold().red(manifest_error.message));
			vite.ws.send({
				type: 'error',
				err: {
					message: manifest_error.message ?? 'Invalid routes',
					stack: ''
				}
			});

			return;
		}

		const node_analyser = create_node_analyser({
			resolve: async (server_node) => {
				const { module } = await resolve(server_node);
				return module;
			}
		});
		invalidate_page_options = node_analyser.invalidate_page_options;

		manifest = {
			appDir: svelte_config.kit.appDir,
			appPath: svelte_config.kit.appDir,
			assets: new Set(manifest_data.assets.map((asset) => asset.file)),
			mimeTypes: get_mime_lookup(manifest_data),
			_: {
				client: {
					start: `${runtime_base}/client/entry.js`,
					app: `${to_fs(svelte_config.kit.outDir)}/generated/client/app.js`,
					imports: [],
					stylesheets: [],
					fonts: [],
					uses_env_dynamic_public: true,
					nodes:
						svelte_config.kit.router.resolution === 'client'
							? undefined
							: manifest_data.nodes.map((node, i) => {
									if (node.component || node.universal) {
										return `${svelte_config.kit.paths.base}${to_fs(svelte_config.kit.outDir)}/generated/client/nodes/${i}.js`;
									}
								}),
					// `css` is not necessary in dev, as the JS file from `nodes` will reference the CSS file
					routes:
						svelte_config.kit.router.resolution === 'client'
							? undefined
							: compact(
									manifest_data.routes.map((route) => {
										if (!route.page) return;

										return {
											id: route.id,
											pattern: route.pattern,
											params: route.params,
											layouts: route.page.layouts.map((l) =>
												l !== undefined ? [!!manifest_data.nodes[l].server, l] : undefined
											),
											errors: route.page.errors,
											leaf: [!!manifest_data.nodes[route.page.leaf].server, route.page.leaf]
										};
									})
								)
				},
				server_assets: new Proxy(
					{},
					{
						has: (_, /** @type {string} */ file) => fs.existsSync(from_fs(file)),
						get: (_, /** @type {string} */ file) => fs.statSync(from_fs(file)).size
					}
				),
				nodes: manifest_data.nodes.map((node, index) => {
					return async () => {
						/** @type {import('types').SSRNode} */
						const result = {};
						result.index = index;
						result.universal_id = node.universal;
						result.server_id = node.server;

						// these are unused in dev, but it's easier to include them
						result.imports = [];
						result.stylesheets = [];
						result.fonts = [];

						/** @type {import('vite').ModuleNode[]} */
						const module_nodes = [];

						if (node.component) {
							result.component = async () => {
								const { module_node, module } = await resolve(
									/** @type {string} */ (node.component)
								);

								module_nodes.push(module_node);

								return module.default;
							};
						}

						if (node.universal) {
							const page_options = await node_analyser.get_page_options(node);
							if (page_options?.ssr === false) {
								result.universal = page_options;
							} else {
								// TODO: explain why the file was loaded on the server if we fail to load it
								const { module, module_node } = await resolve(node.universal);
								module_nodes.push(module_node);
								result.universal = module;
							}
						}

						if (node.server) {
							const { module } = await resolve(node.server);
							result.server = module;
						}

						// in dev we inline all styles to avoid FOUC. this gets populated lazily so that
						// components/stylesheets loaded via import() during `load` are included
						result.inline_styles = async () => {
							/** @type {Set<import('vite').ModuleNode | import('vite').EnvironmentModuleNode>} */
							const deps = new Set();

							for (const module_node of module_nodes) {
								await find_deps(vite, module_node, deps);
							}

							/** @type {Record<string, string>} */
							const styles = {};

							for (const dep of deps) {
								if (isCSSRequest(dep.url) && !vite_css_query_regex.test(dep.url)) {
									const inlineCssUrl = dep.url.includes('?')
										? dep.url.replace('?', '?inline&')
										: dep.url + '?inline';
									try {
										const mod = await vite.ssrLoadModule(inlineCssUrl);
										styles[dep.url] = mod.default;
									} catch {
										// this can happen with dynamically imported modules, I think
										// because the Vite module graph doesn't distinguish between
										// static and dynamic imports? TODO investigate, submit fix
									}
								}
							}

							return styles;
						};

						return result;
					};
				}),
				prerendered_routes: new Set(),
				get remotes() {
					return Object.fromEntries(
						get_remotes().map((remote) => [
							remote.hash,
							() => vite.ssrLoadModule(remote.file).then((module) => ({ default: module }))
						])
					);
				},
				routes: compact(
					manifest_data.routes.map((route) => {
						if (!route.page && !route.endpoint) return null;

						const endpoint = route.endpoint;

						return {
							id: route.id,
							pattern: route.pattern,
							params: route.params,
							page: route.page,
							endpoint: endpoint
								? async () => {
										const url = path.resolve(cwd, endpoint.file);
										return await loud_ssr_load_module(url);
									}
								: null,
							endpoint_id: endpoint?.file
						};
					})
				),
				matchers: async () => {
					/** @type {Record<string, import('@sveltejs/kit').ParamMatcher>} */
					const matchers = {};

					for (const key in manifest_data.matchers) {
						const file = manifest_data.matchers[key];
						const url = path.resolve(cwd, file);
						const module = await vite.ssrLoadModule(url, { fixStacktrace: true });

						if (module.match) {
							matchers[key] = module.match;
						} else {
							throw new Error(`${file} does not export a \`match\` function`);
						}
					}

					return matchers;
				}
			}
		};
	}

	/** @param {Error} error */
	function fix_stack_trace(error) {
		try {
			vite.ssrFixStacktrace(error);
		} catch {
			// ssrFixStacktrace can fail on StackBlitz web containers and we don't know why
			// by ignoring it the line numbers are wrong, but at least we can show the error
		}
		return error.stack;
	}

	update_manifest();

	/**
	 * @param {string} event
	 * @param {(file: string) => void} cb
	 */
	const watch = (event, cb) => {
		vite.watcher.on(event, (file) => {
			if (
				file.startsWith(svelte_config.kit.files.routes + path.sep) ||
				file.startsWith(svelte_config.kit.files.params + path.sep) ||
				svelte_config.kit.moduleExtensions.some((ext) => file.endsWith(`.remote${ext}`)) ||
				// in contrast to server hooks, client hooks are written to the client manifest
				// and therefore need rebuilding when they are added/removed
				file.startsWith(svelte_config.kit.files.hooks.client)
			) {
				cb(file);
			}
		});
	};
	/** @type {NodeJS.Timeout | null } */
	let timeout = null;
	/** @param {() => void} to_run */
	const debounce = (to_run) => {
		timeout && clearTimeout(timeout);
		timeout = setTimeout(() => {
			timeout = null;
			to_run();
		}, 100);
	};

	// flag to skip watchers if server is already restarting
	let restarting = false;

	// Debounce add/unlink events because in case of folder deletion or moves
	// they fire in rapid succession, causing needless invocations.
	// These watchers only run for routes, param matchers, and client hooks.
	watch('add', () => debounce(update_manifest));
	watch('unlink', () => debounce(update_manifest));
	watch('change', (file) => {
		// Don't run for a single file if the whole manifest is about to get updated
		if (timeout || restarting) return;

		if (/\+(page|layout).*$/.test(file)) {
			invalidate_page_options(path.relative(cwd, file));
		}

		sync.update(svelte_config, manifest_data, file);
	});

	const { appTemplate, errorTemplate, serviceWorker, hooks } = svelte_config.kit.files;

	// vite client only executes a full reload if the triggering html file path is index.html
	// kit defaults to src/app.html, so unless user changed that to index.html
	// send the vite client a full-reload event without path being set
	if (appTemplate !== 'index.html') {
		vite.watcher.on('change', (file) => {
			if (file === appTemplate && !restarting) {
				vite.ws.send({ type: 'full-reload' });
			}
		});
	}

	vite.watcher.on('all', (_, file) => {
		if (
			file === appTemplate ||
			file === errorTemplate ||
			file.startsWith(serviceWorker) ||
			file.startsWith(hooks.server)
		) {
			sync.server(svelte_config);
		}
	});

	vite.watcher.on('change', async (file) => {
		// changing the svelte config requires restarting the dev server
		// the config is only read on start and passed on to vite-plugin-svelte
		// which needs up-to-date values to operate correctly
		if (file.match(/[/\\]svelte\.config\.[jt]s$/)) {
			console.log(`svelte config changed, restarting vite dev-server. changed file: ${file}`);
			restarting = true;
			await vite.restart();
		}
	});

	const assets = svelte_config.kit.paths.assets ? SVELTE_KIT_ASSETS : svelte_config.kit.paths.base;
	const asset_server = sirv(svelte_config.kit.files.assets, {
		dev: true,
		etag: true,
		maxAge: 0,
		extensions: [],
		setHeaders: (res) => {
			res.setHeader('access-control-allow-origin', '*');
		}
	});

	vite.middlewares.use((req, res, next) => {
		const base = `${vite.config.server.https ? 'https' : 'http'}://${
			req.headers[':authority'] || req.headers.host
		}`;

		const decoded = decodeURI(new URL(base + req.url).pathname);

		if (decoded.startsWith(assets)) {
			const pathname = decoded.slice(assets.length);
			const file = svelte_config.kit.files.assets + pathname;

			if (fs.existsSync(file) && !fs.statSync(file).isDirectory()) {
				if (has_correct_case(file, svelte_config.kit.files.assets)) {
					req.url = encodeURI(pathname); // don't need query/hash
					asset_server(req, res);
					return;
				}
			}
		}

		next();
	});

	const env = loadEnv(vite_config.mode, svelte_config.kit.env.dir, '');
	const emulator = await svelte_config.kit.adapter?.emulate?.();

	return () => {
		const serve_static_middleware = vite.middlewares.stack.find(
			(middleware) =>
				/** @type {function} */ (middleware.handle).name === 'viteServeStaticMiddleware'
		);

		// Vite will give a 403 on URLs like /test, /static, and /package.json preventing us from
		// serving routes with those names. See https://github.com/vitejs/vite/issues/7363
		remove_static_middlewares(vite.middlewares);

		vite.middlewares.use(async (req, res) => {
			// Vite's base middleware strips out the base path. Restore it
			const original_url = req.url;
			req.url = req.originalUrl;
			try {
				const base = `${vite.config.server.https ? 'https' : 'http'}://${
					req.headers[':authority'] || req.headers.host
				}`;

				const decoded = decodeURI(new URL(base + req.url).pathname);
				const file = posixify(path.resolve(decoded.slice(svelte_config.kit.paths.base.length + 1)));
				const is_file = fs.existsSync(file) && !fs.statSync(file).isDirectory();
				const allowed =
					!vite_config.server.fs.strict ||
					vite_config.server.fs.allow.some((dir) => file.startsWith(dir));

				if (is_file && allowed) {
					req.url = original_url;
					// @ts-expect-error
					serve_static_middleware.handle(req, res);
					return;
				}

				if (!decoded.startsWith(svelte_config.kit.paths.base)) {
					return not_found(req, res, svelte_config.kit.paths.base);
				}

				if (decoded === svelte_config.kit.paths.base + '/service-worker.js') {
					const resolved = resolve_entry(svelte_config.kit.files.serviceWorker);

					if (resolved) {
						res.writeHead(200, {
							'content-type': 'application/javascript'
						});
						res.end(`import '${svelte_config.kit.paths.base}${to_fs(resolved)}';`);
					} else {
						res.writeHead(404);
						res.end('not found');
					}

					return;
				}

				const resolved_instrumentation = resolve_entry(
					path.join(svelte_config.kit.files.src, 'instrumentation.server')
				);

				if (resolved_instrumentation) {
					await vite.ssrLoadModule(resolved_instrumentation);
				}

				// we have to import `Server` before calling `set_assets`
				const { Server } = /** @type {import('types').ServerModule} */ (
					await vite.ssrLoadModule(`${runtime_base}/server/index.js`, { fixStacktrace: true })
				);

				const { set_fix_stack_trace } = await vite.ssrLoadModule(
					`${runtime_base}/shared-server.js`
				);
				set_fix_stack_trace(fix_stack_trace);

				const { set_assets } = await vite.ssrLoadModule('$app/paths/internal/server');
				set_assets(assets);

				const server = new Server(manifest);

				await server.init({
					env,
					read: (file) => createReadableStream(from_fs(file))
				});

				const request = await getRequest({
					base,
					request: req
				});

				if (manifest_error) {
					console.error(colors.bold().red(manifest_error.message));

					const error_page = load_error_page(svelte_config);

					/** @param {{ status: number; message: string }} opts */
					const error_template = ({ status, message }) => {
						return error_page
							.replace(/%sveltekit\.status%/g, String(status))
							.replace(/%sveltekit\.error\.message%/g, escape_html(message));
					};

					res.writeHead(500, {
						'Content-Type': 'text/html; charset=utf-8'
					});
					res.end(
						error_template({ status: 500, message: manifest_error.message ?? 'Invalid routes' })
					);

					return;
				}

				const rendered = await server.respond(request, {
					getClientAddress: () => {
						const { remoteAddress } = req.socket;
						if (remoteAddress) return remoteAddress;
						throw new Error('Could not determine clientAddress');
					},
					read: (file) => {
						if (file in manifest._.server_assets) {
							return fs.readFileSync(from_fs(file));
						}

						return fs.readFileSync(path.join(svelte_config.kit.files.assets, file));
					},
					before_handle: (event, config, prerender) => {
						async_local_storage.enterWith({ event, config, prerender });
					},
					emulator
				});

				if (rendered.status === 404) {
					// @ts-expect-error
					serve_static_middleware.handle(req, res, () => {
						void setResponse(res, rendered);
					});
				} else {
					void setResponse(res, rendered);
				}
			} catch (e) {
				const error = coalesce_to_error(e);
				res.statusCode = 500;
				res.end(fix_stack_trace(error));
			}
		});
	};
}

/**
 * @param {import('connect').Server} server
 */
function remove_static_middlewares(server) {
	const static_middlewares = ['viteServeStaticMiddleware', 'viteServePublicMiddleware'];
	for (let i = server.stack.length - 1; i > 0; i--) {
		// @ts-expect-error using internals
		if (static_middlewares.includes(server.stack[i].handle.name)) {
			server.stack.splice(i, 1);
		}
	}
}

/**
 * @param {import('vite').ViteDevServer} vite
 * @param {import('vite').ModuleNode | import('vite').EnvironmentModuleNode} node
 * @param {Set<import('vite').ModuleNode | import('vite').EnvironmentModuleNode>} deps
 */
async function find_deps(vite, node, deps) {
	// since `ssrTransformResult.deps` contains URLs instead of `ModuleNode`s, this process is asynchronous.
	// instead of using `await`, we resolve all branches in parallel.
	/** @type {Promise<void>[]} */
	const branches = [];

	/** @param {import('vite').ModuleNode | import('vite').EnvironmentModuleNode} node */
	async function add(node) {
		if (!deps.has(node)) {
			deps.add(node);
			await find_deps(vite, node, deps);
		}
	}

	/** @param {string} url */
	async function add_by_url(url) {
		const node = await get_server_module_by_url(vite, url);

		if (node) {
			await add(node);
		}
	}

	const transform_result =
		/** @type {import('vite').ModuleNode} */ (node).ssrTransformResult || node.transformResult;

	if (transform_result) {
		if (transform_result.deps) {
			transform_result.deps.forEach((url) => branches.push(add_by_url(url)));
		}

		if (transform_result.dynamicDeps) {
			transform_result.dynamicDeps.forEach((url) => branches.push(add_by_url(url)));
		}
	} else {
		node.importedModules.forEach((node) => branches.push(add(node)));
	}

	await Promise.all(branches);
}

/**
 * @param {import('vite').ViteDevServer} vite
 * @param {string} url
 */
function get_server_module_by_url(vite, url) {
	return vite.environments
		? vite.environments.ssr.moduleGraph.getModuleByUrl(url)
		: vite.moduleGraph.getModuleByUrl(url, true);
}

/**
 * Determine if a file is being requested with the correct case,
 * to ensure consistent behaviour between dev and prod and across
 * operating systems. Note that we can't use realpath here,
 * because we don't want to follow symlinks
 * @param {string} file
 * @param {string} assets
 * @returns {boolean}
 */
function has_correct_case(file, assets) {
	if (file === assets) return true;

	const parent = path.dirname(file);

	if (fs.readdirSync(parent).includes(path.basename(file))) {
		return has_correct_case(parent, assets);
	}

	return false;
}
