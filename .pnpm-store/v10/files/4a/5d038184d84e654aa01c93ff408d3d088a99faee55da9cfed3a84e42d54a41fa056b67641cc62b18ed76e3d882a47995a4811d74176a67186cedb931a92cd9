import { existsSync, readFileSync, statSync, writeFileSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { pathToFileURL } from 'node:url';
import { installPolyfills } from '../../exports/node/polyfills.js';
import { mkdirp, posixify, walk } from '../../utils/filesystem.js';
import { decode_uri, is_root_relative, resolve } from '../../utils/url.js';
import { escape_html } from '../../utils/escape.js';
import { logger } from '../utils.js';
import { load_config } from '../config/index.js';
import { get_route_segments } from '../../utils/routing.js';
import { queue } from './queue.js';
import { crawl } from './crawl.js';
import { forked } from '../../utils/fork.js';
import * as devalue from 'devalue';
import { createReadableStream } from '@sveltejs/kit/node';
import generate_fallback from './fallback.js';

export default forked(import.meta.url, prerender);

// https://html.spec.whatwg.org/multipage/browsing-the-web.html#scrolling-to-a-fragment
// "If fragment is the empty string, then return the special value top of the document."
// ...and
// "If decodedFragment is an ASCII case-insensitive match for the string 'top', then return the top of the document."
const SPECIAL_HASHLINKS = new Set(['', 'top']);

/**
 * @param {{
 *   hash: boolean;
 *   out: string;
 *   manifest_path: string;
 *   metadata: import('types').ServerMetadata;
 *   verbose: boolean;
 *   env: Record<string, string>
 * }} opts
 */
async function prerender({ hash, out, manifest_path, metadata, verbose, env }) {
	/** @type {import('@sveltejs/kit').SSRManifest} */
	const manifest = (await import(pathToFileURL(manifest_path).href)).manifest;

	/** @type {import('types').ServerInternalModule} */
	const internal = await import(pathToFileURL(`${out}/server/internal.js`).href);

	/** @type {import('types').ServerModule} */
	const { Server } = await import(pathToFileURL(`${out}/server/index.js`).href);

	// configure `import { building } from '$app/environment'` —
	// essential we do this before analysing the code
	internal.set_building();
	internal.set_prerendering();

	/**
	 * @template {{message: string}} T
	 * @template {Omit<T, 'message'>} K
	 * @param {import('types').Logger} log
	 * @param {'fail' | 'warn' | 'ignore' | ((details: T) => void)} input
	 * @param {(details: K) => string} format
	 * @returns {(details: K) => void}
	 */
	function normalise_error_handler(log, input, format) {
		switch (input) {
			case 'fail':
				return (details) => {
					throw new Error(format(details));
				};
			case 'warn':
				return (details) => {
					log.error(format(details));
				};
			case 'ignore':
				return () => {};
			default:
				// @ts-expect-error TS thinks T might be of a different kind, but it's not
				return (details) => input({ ...details, message: format(details) });
		}
	}

	const OK = 2;
	const REDIRECT = 3;

	/** @type {import('types').Prerendered} */
	const prerendered = {
		pages: new Map(),
		assets: new Map(),
		redirects: new Map(),
		paths: []
	};

	/** @type {import('types').PrerenderMap} */
	const prerender_map = new Map();

	for (const [id, { prerender }] of metadata.routes) {
		if (prerender !== undefined) {
			prerender_map.set(id, prerender);
		}
	}

	/** @type {Set<string>} */
	const prerendered_routes = new Set();

	/** @type {import('types').ValidatedKitConfig} */
	const config = (await load_config()).kit;

	if (hash) {
		const fallback = await generate_fallback({
			manifest_path,
			env
		});

		const file = output_filename('/', true);
		const dest = `${config.outDir}/output/prerendered/pages/${file}`;

		mkdirp(dirname(dest));
		writeFileSync(dest, fallback);

		prerendered.pages.set('/', { file });

		return { prerendered, prerender_map };
	}

	const emulator = await config.adapter?.emulate?.();

	/** @type {import('types').Logger} */
	const log = logger({ verbose });

	installPolyfills();

	/** @type {Map<string, string>} */
	const saved = new Map();

	const handle_http_error = normalise_error_handler(
		log,
		config.prerender.handleHttpError,
		({ status, path, referrer, referenceType }) => {
			const message =
				status === 404 && !path.startsWith(config.paths.base)
					? `${path} does not begin with \`base\`, which is configured in \`paths.base\` and can be imported from \`$app/paths\` - see https://svelte.dev/docs/kit/configuration#paths for more info`
					: path;

			return `${status} ${message}${referrer ? ` (${referenceType} from ${referrer})` : ''}`;
		}
	);

	const handle_missing_id = normalise_error_handler(
		log,
		config.prerender.handleMissingId,
		({ path, id, referrers }) => {
			return (
				`The following pages contain links to ${path}#${id}, but no element with id="${id}" exists on ${path} - see the \`handleMissingId\` option in https://svelte.dev/docs/kit/configuration#prerender for more info:` +
				referrers.map((l) => `\n  - ${l}`).join('')
			);
		}
	);

	const handle_entry_generator_mismatch = normalise_error_handler(
		log,
		config.prerender.handleEntryGeneratorMismatch,
		({ generatedFromId, entry, matchedId }) => {
			return `The entries export from ${generatedFromId} generated entry ${entry}, which was matched by ${matchedId} - see the \`handleEntryGeneratorMismatch\` option in https://svelte.dev/docs/kit/configuration#prerender for more info.`;
		}
	);

	const q = queue(config.prerender.concurrency);

	/**
	 * @param {string} path
	 * @param {boolean} is_html
	 */
	function output_filename(path, is_html) {
		const file = path.slice(config.paths.base.length + 1) || 'index.html';

		if (is_html && !file.endsWith('.html')) {
			return file + (file.endsWith('/') ? 'index.html' : '.html');
		}

		return file;
	}

	const files = new Set(walk(`${out}/client`).map(posixify));
	files.add(`${config.appDir}/env.js`);

	const immutable = `${config.appDir}/immutable`;
	if (existsSync(`${out}/server/${immutable}`)) {
		for (const file of walk(`${out}/server/${immutable}`)) {
			files.add(posixify(`${config.appDir}/immutable/${file}`));
		}
	}
	const seen = new Set();
	const written = new Set();

	/** @type {Map<string, Set<string>>} */
	const expected_hashlinks = new Map();

	/** @type {Map<string, string[]>} */
	const actual_hashlinks = new Map();

	/**
	 * @param {string | null} referrer
	 * @param {string} decoded
	 * @param {string} [encoded]
	 * @param {string} [generated_from_id]
	 */
	function enqueue(referrer, decoded, encoded, generated_from_id) {
		if (seen.has(decoded)) return;
		seen.add(decoded);

		const file = decoded.slice(config.paths.base.length + 1);
		if (files.has(file)) return;

		return q.add(() => visit(decoded, encoded || encodeURI(decoded), referrer, generated_from_id));
	}

	/**
	 * @param {string} decoded
	 * @param {string} encoded
	 * @param {string?} referrer
	 * @param {string} [generated_from_id]
	 */
	async function visit(decoded, encoded, referrer, generated_from_id) {
		if (!decoded.startsWith(config.paths.base)) {
			handle_http_error({ status: 404, path: decoded, referrer, referenceType: 'linked' });
			return;
		}

		/** @type {Map<string, import('types').PrerenderDependency>} */
		const dependencies = new Map();

		const response = await server.respond(new Request(config.prerender.origin + encoded), {
			getClientAddress() {
				throw new Error('Cannot read clientAddress during prerendering');
			},
			prerendering: {
				dependencies
			},
			read: (file) => {
				// stuff we just wrote
				const filepath = saved.get(file);
				if (filepath) return readFileSync(filepath);

				// stuff in `static`
				return readFileSync(join(config.files.assets, file));
			},
			emulator
		});

		const encoded_id = response.headers.get('x-sveltekit-routeid');
		const decoded_id = encoded_id && decode_uri(encoded_id);
		if (
			decoded_id !== null &&
			generated_from_id !== undefined &&
			decoded_id !== generated_from_id
		) {
			handle_entry_generator_mismatch({
				generatedFromId: generated_from_id,
				entry: decoded,
				matchedId: decoded_id
			});
		}

		const body = Buffer.from(await response.arrayBuffer());

		save('pages', response, body, decoded, encoded, referrer, 'linked');

		for (const [dependency_path, result] of dependencies) {
			// this seems circuitous, but using new URL allows us to not care
			// whether dependency_path is encoded or not
			const encoded_dependency_path = new URL(dependency_path, 'http://localhost').pathname;
			const decoded_dependency_path = decode_uri(encoded_dependency_path);

			const headers = Object.fromEntries(result.response.headers);

			const prerender = headers['x-sveltekit-prerender'];
			if (prerender) {
				const encoded_route_id = headers['x-sveltekit-routeid'];
				if (encoded_route_id != null) {
					const route_id = decode_uri(encoded_route_id);
					const existing_value = prerender_map.get(route_id);
					if (existing_value !== 'auto') {
						prerender_map.set(route_id, prerender === 'true' ? true : 'auto');
					}
				}
			}

			const body = result.body ?? new Uint8Array(await result.response.arrayBuffer());

			save(
				'dependencies',
				result.response,
				body,
				decoded_dependency_path,
				encoded_dependency_path,
				decoded,
				'fetched'
			);
		}

		// avoid triggering `filterSerializeResponseHeaders` guard
		const headers = Object.fromEntries(response.headers);

		if (config.prerender.crawl && headers['content-type'] === 'text/html') {
			const { ids, hrefs } = crawl(body.toString(), decoded);

			actual_hashlinks.set(decoded, ids);

			/** @param {string} href */
			const removePrerenderOrigin = (href) => {
				if (href.startsWith(config.prerender.origin)) {
					if (href === config.prerender.origin) return '/';
					if (href.at(config.prerender.origin.length) !== '/') return href;
					return href.slice(config.prerender.origin.length);
				}
				return href;
			};

			for (const href of hrefs.map(removePrerenderOrigin)) {
				if (!is_root_relative(href)) continue;

				const { pathname, search, hash } = new URL(href, 'http://localhost');

				if (search) {
					// TODO warn that query strings have no effect on statically-exported pages
				}

				if (hash) {
					const key = decode_uri(pathname + hash);

					if (!expected_hashlinks.has(key)) {
						expected_hashlinks.set(key, new Set());
					}

					/** @type {Set<string>} */ (expected_hashlinks.get(key)).add(decoded);
				}

				void enqueue(decoded, decode_uri(pathname), pathname);
			}
		}
	}

	/**
	 * @param {'pages' | 'dependencies'} category
	 * @param {Response} response
	 * @param {string | Uint8Array} body
	 * @param {string} decoded
	 * @param {string} encoded
	 * @param {string | null} referrer
	 * @param {'linked' | 'fetched'} referenceType
	 */
	function save(category, response, body, decoded, encoded, referrer, referenceType) {
		const response_type = Math.floor(response.status / 100);
		const headers = Object.fromEntries(response.headers);

		const type = headers['content-type'];
		const is_html = response_type === REDIRECT || type === 'text/html';

		const file = output_filename(decoded, is_html);
		const dest = `${config.outDir}/output/prerendered/${category}/${file}`;

		if (written.has(file)) return;

		const encoded_route_id = response.headers.get('x-sveltekit-routeid');
		const route_id = encoded_route_id != null ? decode_uri(encoded_route_id) : null;
		if (route_id !== null) prerendered_routes.add(route_id);

		if (response_type === REDIRECT) {
			const location = headers['location'];

			if (location) {
				const resolved = resolve(encoded, location);
				if (is_root_relative(resolved)) {
					void enqueue(decoded, decode_uri(resolved), resolved);
				}

				if (!headers['x-sveltekit-normalize']) {
					mkdirp(dirname(dest));

					log.warn(`${response.status} ${decoded} -> ${location}`);

					writeFileSync(
						dest,
						`<script>location.href=${devalue.uneval(
							location
						)};</script><meta http-equiv="refresh" content="${escape_html(
							`0;url=${location}`,
							true
						)}">`
					);

					written.add(file);

					if (!prerendered.redirects.has(decoded)) {
						prerendered.redirects.set(decoded, {
							status: response.status,
							location: resolved
						});

						prerendered.paths.push(decoded);
					}
				}
			} else {
				log.warn(`location header missing on redirect received from ${decoded}`);
			}

			return;
		}

		if (response.status === 200) {
			if (existsSync(dest) && statSync(dest).isDirectory()) {
				throw new Error(
					`Cannot save ${decoded} as it is already a directory. See https://svelte.dev/docs/kit/page-options#prerender-route-conflicts for more information`
				);
			}

			const dir = dirname(dest);

			if (existsSync(dir) && !statSync(dir).isDirectory()) {
				const parent = decoded.split('/').slice(0, -1).join('/');
				throw new Error(
					`Cannot save ${decoded} as ${parent} is already a file. See https://svelte.dev/docs/kit/page-options#prerender-route-conflicts for more information`
				);
			}

			mkdirp(dir);

			log.info(`${response.status} ${decoded}`);
			writeFileSync(dest, body);
			written.add(file);

			if (is_html) {
				prerendered.pages.set(decoded, {
					file
				});
			} else {
				prerendered.assets.set(decoded, {
					type
				});
			}

			prerendered.paths.push(decoded);
		} else if (response_type !== OK) {
			handle_http_error({ status: response.status, path: decoded, referrer, referenceType });
		}

		manifest.assets.add(file);
		saved.set(file, dest);
	}

	/** @type {Array<{ id: string, entries: Array<string>}>} */
	const route_level_entries = [];
	for (const [id, { entries }] of metadata.routes.entries()) {
		if (entries) {
			route_level_entries.push({ id, entries });
		}
	}

	let has_prerenderable_routes = false;

	for (const value of prerender_map.values()) {
		if (value) {
			has_prerenderable_routes = true;
			break;
		}
	}

	if (
		(config.prerender.entries.length === 0 && route_level_entries.length === 0) ||
		!has_prerenderable_routes
	) {
		return { prerendered, prerender_map };
	}

	log.info('Prerendering');

	const server = new Server(manifest);
	await server.init({
		env,
		read: (file) => createReadableStream(`${config.outDir}/output/server/${file}`)
	});

	for (const entry of config.prerender.entries) {
		if (entry === '*') {
			for (const [id, prerender] of prerender_map) {
				if (prerender) {
					// remove optional parameters from the route
					const segments = get_route_segments(id).filter((segment) => !segment.startsWith('[['));
					const processed_id = '/' + segments.join('/');

					if (processed_id.includes('[')) continue;
					const path = `/${get_route_segments(processed_id).join('/')}`;
					void enqueue(null, config.paths.base + path);
				}
			}
		} else {
			void enqueue(null, config.paths.base + entry);
		}
	}

	for (const { id, entries } of route_level_entries) {
		for (const entry of entries) {
			void enqueue(null, config.paths.base + entry, undefined, id);
		}
	}

	await q.done();

	// handle invalid fragment links
	for (const [key, referrers] of expected_hashlinks) {
		const index = key.indexOf('#');
		const path = key.slice(0, index);
		const id = key.slice(index + 1);

		const hashlinks = actual_hashlinks.get(path);
		// ignore fragment links to pages that were not prerendered
		if (!hashlinks) continue;

		if (!hashlinks.includes(id) && !SPECIAL_HASHLINKS.has(id)) {
			handle_missing_id({ id, path, referrers: Array.from(referrers) });
		}
	}

	/** @type {string[]} */
	const not_prerendered = [];

	for (const [route_id, prerender] of prerender_map) {
		if (prerender === true && !prerendered_routes.has(route_id)) {
			not_prerendered.push(route_id);
		}
	}

	if (not_prerendered.length > 0) {
		const list = not_prerendered.map((id) => `  - ${id}`).join('\n');

		throw new Error(
			`The following routes were marked as prerenderable, but were not prerendered because they were not found while crawling your app:\n${list}\n\nSee https://svelte.dev/docs/kit/page-options#prerender-troubleshooting for info on how to solve this`
		);
	}

	return { prerendered, prerender_map };
}
