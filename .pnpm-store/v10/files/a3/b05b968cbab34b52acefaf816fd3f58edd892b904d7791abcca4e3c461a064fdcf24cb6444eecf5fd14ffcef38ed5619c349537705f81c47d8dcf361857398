import * as set_cookie_parser from 'set-cookie-parser';
import { respond } from './respond.js';
import * as paths from '$app/paths/internal/server';
import { read_implementation } from '__sveltekit/server';
import { has_prerendered_path } from './utils.js';

/**
 * @param {{
 *   event: import('@sveltejs/kit').RequestEvent;
 *   options: import('types').SSROptions;
 *   manifest: import('@sveltejs/kit').SSRManifest;
 *   state: import('types').SSRState;
 *   get_cookie_header: (url: URL, header: string | null) => string;
 *   set_internal: (name: string, value: string, opts: import('./page/types.js').Cookie['options']) => void;
 * }} opts
 * @returns {typeof fetch}
 */
export function create_fetch({ event, options, manifest, state, get_cookie_header, set_internal }) {
	/**
	 * @type {typeof fetch}
	 */
	const server_fetch = async (info, init) => {
		const original_request = normalize_fetch_input(info, init, event.url);

		// some runtimes (e.g. Cloudflare) error if you access `request.mode`,
		// annoyingly, so we need to read the value from the `init` object instead
		let mode = (info instanceof Request ? info.mode : init?.mode) ?? 'cors';
		let credentials =
			(info instanceof Request ? info.credentials : init?.credentials) ?? 'same-origin';

		return options.hooks.handleFetch({
			event,
			request: original_request,
			fetch: async (info, init) => {
				const request = normalize_fetch_input(info, init, event.url);

				const url = new URL(request.url);

				if (!request.headers.has('origin')) {
					request.headers.set('origin', event.url.origin);
				}

				if (info !== original_request) {
					mode = (info instanceof Request ? info.mode : init?.mode) ?? 'cors';
					credentials =
						(info instanceof Request ? info.credentials : init?.credentials) ?? 'same-origin';
				}

				// Remove Origin, according to https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin#description
				if (
					(request.method === 'GET' || request.method === 'HEAD') &&
					((mode === 'no-cors' && url.origin !== event.url.origin) ||
						url.origin === event.url.origin)
				) {
					request.headers.delete('origin');
				}

				if (url.origin !== event.url.origin) {
					// Allow cookie passthrough for "credentials: same-origin" and "credentials: include"
					// if SvelteKit is serving my.domain.com:
					// -        domain.com WILL NOT receive cookies
					// -     my.domain.com WILL receive cookies
					// -    api.domain.dom WILL NOT receive cookies
					// - sub.my.domain.com WILL receive cookies
					// ports do not affect the resolution
					// leading dot prevents mydomain.com matching domain.com
					// Do not forward other cookies for "credentials: include" because we don't know
					// which cookie belongs to which domain (browser does not pass this info)
					if (`.${url.hostname}`.endsWith(`.${event.url.hostname}`) && credentials !== 'omit') {
						const cookie = get_cookie_header(url, request.headers.get('cookie'));
						if (cookie) request.headers.set('cookie', cookie);
					}

					return fetch(request);
				}

				// handle fetch requests for static assets. e.g. prebaked data, etc.
				// we need to support everything the browser's fetch supports
				const prefix = paths.assets || paths.base;
				const decoded = decodeURIComponent(url.pathname);
				const filename = (
					decoded.startsWith(prefix) ? decoded.slice(prefix.length) : decoded
				).slice(1);
				const filename_html = `${filename}/index.html`; // path may also match path/index.html

				const is_asset = manifest.assets.has(filename) || filename in manifest._.server_assets;
				const is_asset_html =
					manifest.assets.has(filename_html) || filename_html in manifest._.server_assets;

				if (is_asset || is_asset_html) {
					const file = is_asset ? filename : filename_html;

					if (state.read) {
						const type = is_asset
							? manifest.mimeTypes[filename.slice(filename.lastIndexOf('.'))]
							: 'text/html';

						return new Response(state.read(file), {
							headers: type ? { 'content-type': type } : {}
						});
					} else if (read_implementation && file in manifest._.server_assets) {
						const length = manifest._.server_assets[file];
						const type = manifest.mimeTypes[file.slice(file.lastIndexOf('.'))];

						return new Response(read_implementation(file), {
							headers: {
								'Content-Length': '' + length,
								'Content-Type': type
							}
						});
					}

					return await fetch(request);
				}

				if (has_prerendered_path(manifest, paths.base + decoded)) {
					// The path of something prerendered could match a different route
					// that is still in the manifest, leading to the wrong route being loaded.
					// We therefore bail early here. The prerendered logic is different for
					// each adapter, (except maybe for prerendered redirects)
					// so we need to make an actual HTTP request.
					return await fetch(request);
				}

				if (credentials !== 'omit') {
					const cookie = get_cookie_header(url, request.headers.get('cookie'));
					if (cookie) {
						request.headers.set('cookie', cookie);
					}

					const authorization = event.request.headers.get('authorization');
					if (authorization && !request.headers.has('authorization')) {
						request.headers.set('authorization', authorization);
					}
				}

				if (!request.headers.has('accept')) {
					request.headers.set('accept', '*/*');
				}

				if (!request.headers.has('accept-language')) {
					request.headers.set(
						'accept-language',
						/** @type {string} */ (event.request.headers.get('accept-language'))
					);
				}

				const response = await internal_fetch(request, options, manifest, state);

				const set_cookie = response.headers.get('set-cookie');
				if (set_cookie) {
					for (const str of set_cookie_parser.splitCookiesString(set_cookie)) {
						const { name, value, ...options } = set_cookie_parser.parseString(str, {
							decodeValues: false
						});

						const path = options.path ?? (url.pathname.split('/').slice(0, -1).join('/') || '/');

						// options.sameSite is string, something more specific is required - type cast is safe
						set_internal(name, value, {
							path,
							encode: (value) => value,
							.../** @type {import('cookie').CookieSerializeOptions} */ (options)
						});
					}
				}

				return response;
			}
		});
	};

	// Don't make this function `async`! Otherwise, the user has to `catch` promises they use for streaming responses or else
	// it will be an unhandled rejection. Instead, we add a `.catch(() => {})` ourselves below to prevent this from happening.
	return (input, init) => {
		// See docs in fetch.js for why we need to do this
		const response = server_fetch(input, init);
		response.catch(() => {});
		return response;
	};
}

/**
 * @param {RequestInfo | URL} info
 * @param {RequestInit | undefined} init
 * @param {URL} url
 */
function normalize_fetch_input(info, init, url) {
	if (info instanceof Request) {
		return info;
	}

	return new Request(typeof info === 'string' ? new URL(info, url) : info, init);
}

/**
 * @param {Request} request
 * @param {import('types').SSROptions} options
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @param {import('types').SSRState} state
 * @returns {Promise<Response>}
 */
async function internal_fetch(request, options, manifest, state) {
	if (request.signal) {
		if (request.signal.aborted) {
			throw new DOMException('The operation was aborted.', 'AbortError');
		}

		let remove_abort_listener = () => {};
		/** @type {Promise<never>} */
		const abort_promise = new Promise((_, reject) => {
			const on_abort = () => {
				reject(new DOMException('The operation was aborted.', 'AbortError'));
			};
			request.signal.addEventListener('abort', on_abort, { once: true });
			remove_abort_listener = () => request.signal.removeEventListener('abort', on_abort);
		});

		const result = await Promise.race([
			respond(request, options, manifest, {
				...state,
				depth: state.depth + 1
			}),
			abort_promise
		]);
		remove_abort_listener();
		return result;
	} else {
		return await respond(request, options, manifest, {
			...state,
			depth: state.depth + 1
		});
	}
}
