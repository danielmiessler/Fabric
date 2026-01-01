import fs from 'node:fs';
import { join } from 'node:path';
import { pathToFileURL } from 'node:url';
import { lookup } from 'mrmime';
import sirv from 'sirv';
import { loadEnv, normalizePath } from 'vite';
import { createReadableStream, getRequest, setResponse } from '../../../exports/node/index.js';
import { installPolyfills } from '../../../exports/node/polyfills.js';
import { SVELTE_KIT_ASSETS } from '../../../constants.js';
import { not_found } from '../utils.js';

/** @typedef {import('http').IncomingMessage} Req */
/** @typedef {import('http').ServerResponse} Res */
/** @typedef {(req: Req, res: Res, next: () => void) => void} Handler */

/**
 * @param {import('vite').PreviewServer} vite
 * @param {import('vite').ResolvedConfig} vite_config
 * @param {import('types').ValidatedConfig} svelte_config
 */
export async function preview(vite, vite_config, svelte_config) {
	installPolyfills();

	const { paths } = svelte_config.kit;
	const base = paths.base;
	const assets = paths.assets ? SVELTE_KIT_ASSETS : paths.base;

	const protocol = vite_config.preview.https ? 'https' : 'http';

	const etag = `"${Date.now()}"`;

	const dir = join(svelte_config.kit.outDir, 'output/server');

	if (!fs.existsSync(dir)) {
		throw new Error(`Server files not found at ${dir}, did you run \`build\` first?`);
	}

	/** @type {import('types').ServerInternalModule} */
	const { set_assets } = await import(pathToFileURL(join(dir, 'internal.js')).href);

	/** @type {import('types').ServerModule} */
	const { Server } = await import(pathToFileURL(join(dir, 'index.js')).href);

	const { manifest } = await import(pathToFileURL(join(dir, 'manifest.js')).href);

	set_assets(assets);

	const server = new Server(manifest);
	await server.init({
		env: loadEnv(vite_config.mode, svelte_config.kit.env.dir, ''),
		read: (file) => createReadableStream(`${dir}/${file}`)
	});

	const emulator = await svelte_config.kit.adapter?.emulate?.();

	return () => {
		// Remove the base middleware. It screws with the URL.
		// It also only lets through requests beginning with the base path, so that requests beginning
		// with the assets URL never reach us. We could serve assets separately before the base
		// middleware, but we'd need that to occur after the compression and cors middlewares, so would
		// need to insert it manually into the stack, which would be at least as bad as doing this.
		for (let i = vite.middlewares.stack.length - 1; i > 0; i--) {
			// @ts-expect-error using internals
			if (vite.middlewares.stack[i].handle.name === 'viteBaseMiddleware') {
				vite.middlewares.stack.splice(i, 1);
			}
		}

		// generated client assets and the contents of `static`
		vite.middlewares.use(
			scoped(
				assets,
				sirv(join(svelte_config.kit.outDir, 'output/client'), {
					setHeaders: (res, pathname) => {
						// only apply to immutable directory, not e.g. version.json
						if (pathname.startsWith(`/${svelte_config.kit.appDir}/immutable`)) {
							res.setHeader('cache-control', 'public,max-age=31536000,immutable');
						}
					}
				})
			)
		);

		vite.middlewares.use((req, res, next) => {
			const original_url = /** @type {string} */ (req.url);
			const { pathname, search } = new URL(original_url, 'http://dummy');

			// if `paths.base === '/a/b/c`, then the root route is `/a/b/c/`,
			// regardless of the `trailingSlash` route option
			if (base.length > 1 && pathname === base) {
				let location = base + '/';
				if (search) location += search;
				res.writeHead(307, {
					location
				});
				res.end();
				return;
			}

			if (pathname.startsWith(base)) {
				next();
			} else {
				res.statusCode = 404;
				not_found(req, res, base);
			}
		});

		// prerendered dependencies
		vite.middlewares.use(
			scoped(base, mutable(join(svelte_config.kit.outDir, 'output/prerendered/dependencies')))
		);

		// prerendered pages (we can't just use sirv because we need to
		// preserve the correct trailingSlash behaviour)
		vite.middlewares.use(
			scoped(base, (req, res, next) => {
				let if_none_match_value = req.headers['if-none-match'];

				if (if_none_match_value?.startsWith('W/"')) {
					if_none_match_value = if_none_match_value.substring(2);
				}

				if (if_none_match_value === etag) {
					res.statusCode = 304;
					res.end();
					return;
				}

				const { pathname, search } = new URL(/** @type {string} */ (req.url), 'http://dummy');

				const dir = pathname.startsWith(`/${svelte_config.kit.appDir}/remote/`) ? 'data' : 'pages';

				let filename = normalizePath(
					join(svelte_config.kit.outDir, `output/prerendered/${dir}` + pathname)
				);

				try {
					filename = decodeURI(filename);
				} catch {
					// malformed URI
				}

				let prerendered = is_file(filename);

				if (!prerendered) {
					const has_trailing_slash = pathname.endsWith('/');
					const html_filename = `${filename}${has_trailing_slash ? 'index.html' : '.html'}`;

					/** @type {string | undefined} */
					let redirect;

					if (is_file(html_filename)) {
						filename = html_filename;
						prerendered = true;
					} else if (has_trailing_slash) {
						if (is_file(filename.slice(0, -1) + '.html')) {
							redirect = pathname.slice(0, -1);
						}
					} else if (is_file(filename + '/index.html')) {
						redirect = pathname + '/';
					}

					if (redirect) {
						if (search) redirect += search;
						res.writeHead(307, {
							location: redirect
						});

						res.end();

						return;
					}
				}

				if (prerendered) {
					res.writeHead(200, {
						'content-type': lookup(pathname) || 'text/html',
						etag
					});

					fs.createReadStream(filename).pipe(res);
				} else {
					next();
				}
			})
		);

		// SSR
		vite.middlewares.use(async (req, res) => {
			const host = req.headers[':authority'] || req.headers.host;

			const request = await getRequest({
				base: `${protocol}://${host}`,
				request: req
			});

			await setResponse(
				res,
				await server.respond(request, {
					getClientAddress: () => {
						const { remoteAddress } = req.socket;
						if (remoteAddress) return remoteAddress;
						throw new Error('Could not determine clientAddress');
					},
					read: (file) => {
						if (file in manifest._.server_assets) {
							return fs.readFileSync(join(dir, file));
						}

						return fs.readFileSync(join(svelte_config.kit.files.assets, file));
					},
					emulator
				})
			);
		});
	};
}

/**
 * @param {string} dir
 * @returns {Handler}
 */
const mutable = (dir) =>
	fs.existsSync(dir)
		? sirv(dir, {
				etag: true,
				maxAge: 0
			})
		: (_req, _res, next) => next();

/**
 * @param {string} scope
 * @param {Handler} handler
 * @returns {Handler}
 */
function scoped(scope, handler) {
	if (scope === '') return handler;

	return (req, res, next) => {
		if (req.url?.startsWith(scope)) {
			const original_url = req.url;
			req.url = req.url.slice(scope.length);
			handler(req, res, () => {
				req.url = original_url;
				next();
			});
		} else {
			next();
		}
	};
}

/** @param {string} path */
function is_file(path) {
	return fs.existsSync(path) && !fs.statSync(path).isDirectory();
}
