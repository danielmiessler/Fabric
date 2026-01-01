import { parse, serialize } from 'cookie';
import { normalize_path, resolve } from '../../utils/url.js';
import { add_data_suffix } from '../pathname.js';

// eslint-disable-next-line no-control-regex -- control characters are invalid in cookie names
const INVALID_COOKIE_CHARACTER_REGEX = /[\x00-\x1F\x7F()<>@,;:"/[\]?={} \t]/;

/**
 * Tracks all cookies set during dev mode so we can emit warnings
 * when we detect that there's likely cookie misusage due to wrong paths
 *
 * @type {Record<string, Set<string>>} */
const cookie_paths = {};

/**
 * Cookies that are larger than this size (including the name and other
 * attributes) are discarded by browsers
 */
const MAX_COOKIE_SIZE = 4129;

// TODO 3.0 remove this check
/** @param {import('./page/types.js').Cookie['options']} options */
function validate_options(options) {
	if (options?.path === undefined) {
		throw new Error('You must specify a `path` when setting, deleting or serializing cookies');
	}
}

/**
 * @param {Request} request
 * @param {URL} url
 */
export function get_cookies(request, url) {
	const header = request.headers.get('cookie') ?? '';
	const initial_cookies = parse(header, { decode: (value) => value });

	/** @type {string | undefined} */
	let normalized_url;

	/** @type {Record<string, import('./page/types.js').Cookie>} */
	const new_cookies = {};

	/** @type {import('cookie').CookieSerializeOptions} */
	const defaults = {
		httpOnly: true,
		sameSite: 'lax',
		secure: url.hostname === 'localhost' && url.protocol === 'http:' ? false : true
	};

	/** @type {import('@sveltejs/kit').Cookies} */
	const cookies = {
		// The JSDoc param annotations appearing below for get, set and delete
		// are necessary to expose the `cookie` library types to
		// typescript users. `@type {import('@sveltejs/kit').Cookies}` above is not
		// sufficient to do so.

		/**
		 * @param {string} name
		 * @param {import('cookie').CookieParseOptions} [opts]
		 */
		get(name, opts) {
			const c = new_cookies[name];
			if (
				c &&
				domain_matches(url.hostname, c.options.domain) &&
				path_matches(url.pathname, c.options.path)
			) {
				return c.value;
			}

			const req_cookies = parse(header, { decode: opts?.decode });
			const cookie = req_cookies[name]; // the decoded string or undefined

			// in development, if the cookie was set during this session with `cookies.set`,
			// but at a different path, warn the user. (ignore cookies from request headers,
			// since we don't know which path they were set at)
			if (__SVELTEKIT_DEV__ && !cookie) {
				const paths = Array.from(cookie_paths[name] ?? []).filter((path) => {
					// we only care about paths that are _more_ specific than the current path
					return path_matches(path, url.pathname) && path !== url.pathname;
				});

				if (paths.length > 0) {
					console.warn(
						// prettier-ignore
						`'${name}' cookie does not exist for ${url.pathname}, but was previously set at ${conjoin([...paths])}. Did you mean to set its 'path' to '/' instead?`
					);
				}
			}

			return cookie;
		},

		/**
		 * @param {import('cookie').CookieParseOptions} [opts]
		 */
		getAll(opts) {
			const cookies = parse(header, { decode: opts?.decode });

			for (const c of Object.values(new_cookies)) {
				if (
					domain_matches(url.hostname, c.options.domain) &&
					path_matches(url.pathname, c.options.path)
				) {
					cookies[c.name] = c.value;
				}
			}

			return Object.entries(cookies).map(([name, value]) => ({ name, value }));
		},

		/**
		 * @param {string} name
		 * @param {string} value
		 * @param {import('./page/types.js').Cookie['options']} options
		 */
		set(name, value, options) {
			// TODO: remove this check in 3.0
			const illegal_characters = name.match(INVALID_COOKIE_CHARACTER_REGEX);
			if (illegal_characters) {
				console.warn(
					`The cookie name "${name}" will be invalid in SvelteKit 3.0 as it contains ${illegal_characters.join(
						' and '
					)}. See RFC 2616 for more details https://datatracker.ietf.org/doc/html/rfc2616#section-2.2`
				);
			}

			validate_options(options);
			set_internal(name, value, { ...defaults, ...options });
		},

		/**
		 * @param {string} name
		 *  @param {import('./page/types.js').Cookie['options']} options
		 */
		delete(name, options) {
			validate_options(options);
			cookies.set(name, '', { ...options, maxAge: 0 });
		},

		/**
		 * @param {string} name
		 * @param {string} value
		 *  @param {import('./page/types.js').Cookie['options']} options
		 */
		serialize(name, value, options) {
			validate_options(options);

			let path = options.path;

			if (!options.domain || options.domain === url.hostname) {
				if (!normalized_url) {
					throw new Error('Cannot serialize cookies until after the route is determined');
				}
				path = resolve(normalized_url, path);
			}

			return serialize(name, value, { ...defaults, ...options, path });
		}
	};

	/**
	 * @param {URL} destination
	 * @param {string | null} header
	 */
	function get_cookie_header(destination, header) {
		/** @type {Record<string, string>} */
		const combined_cookies = {
			// cookies sent by the user agent have lowest precedence
			...initial_cookies
		};

		// cookies previous set during this event with cookies.set have higher precedence
		for (const key in new_cookies) {
			const cookie = new_cookies[key];
			if (!domain_matches(destination.hostname, cookie.options.domain)) continue;
			if (!path_matches(destination.pathname, cookie.options.path)) continue;

			const encoder = cookie.options.encode || encodeURIComponent;
			combined_cookies[cookie.name] = encoder(cookie.value);
		}

		// explicit header has highest precedence
		if (header) {
			const parsed = parse(header, { decode: (value) => value });
			for (const name in parsed) {
				combined_cookies[name] = parsed[name];
			}
		}

		return Object.entries(combined_cookies)
			.map(([name, value]) => `${name}=${value}`)
			.join('; ');
	}

	/** @type {Array<() => void>} */
	const internal_queue = [];

	/**
	 * @param {string} name
	 * @param {string} value
	 * @param {import('./page/types.js').Cookie['options']} options
	 */
	function set_internal(name, value, options) {
		if (!normalized_url) {
			internal_queue.push(() => set_internal(name, value, options));
			return;
		}

		let path = options.path;

		if (!options.domain || options.domain === url.hostname) {
			path = resolve(normalized_url, path);
		}

		new_cookies[name] = { name, value, options: { ...options, path } };

		if (__SVELTEKIT_DEV__) {
			const serialized = serialize(name, value, new_cookies[name].options);
			if (new TextEncoder().encode(serialized).byteLength > MAX_COOKIE_SIZE) {
				throw new Error(`Cookie "${name}" is too large, and will be discarded by the browser`);
			}

			cookie_paths[name] ??= new Set();

			if (!value) {
				cookie_paths[name].delete(path);
			} else {
				cookie_paths[name].add(path);
			}
		}
	}

	/**
	 * @param {import('types').TrailingSlash} trailing_slash
	 */
	function set_trailing_slash(trailing_slash) {
		normalized_url = normalize_path(url.pathname, trailing_slash);
		internal_queue.forEach((fn) => fn());
	}

	return { cookies, new_cookies, get_cookie_header, set_internal, set_trailing_slash };
}

/**
 * @param {string} hostname
 * @param {string} [constraint]
 */
export function domain_matches(hostname, constraint) {
	if (!constraint) return true;

	const normalized = constraint[0] === '.' ? constraint.slice(1) : constraint;

	if (hostname === normalized) return true;
	return hostname.endsWith('.' + normalized);
}

/**
 * @param {string} path
 * @param {string} [constraint]
 */
export function path_matches(path, constraint) {
	if (!constraint) return true;

	const normalized = constraint.endsWith('/') ? constraint.slice(0, -1) : constraint;

	if (path === normalized) return true;
	return path.startsWith(normalized + '/');
}

/**
 * @param {Headers} headers
 * @param {import('./page/types.js').Cookie[]} cookies
 */
export function add_cookies_to_headers(headers, cookies) {
	for (const new_cookie of cookies) {
		const { name, value, options } = new_cookie;
		headers.append('set-cookie', serialize(name, value, options));

		// special case — for routes ending with .html, the route data lives in a sibling
		// `.html__data.json` file rather than a child `/__data.json` file, which means
		// we need to duplicate the cookie
		if (options.path.endsWith('.html')) {
			const path = add_data_suffix(options.path);
			headers.append('set-cookie', serialize(name, value, { ...options, path }));
		}
	}
}

/**
 * @param {string[]} array
 */
function conjoin(array) {
	if (array.length <= 2) return array.join(' and ');
	return `${array.slice(0, -1).join(', ')} and ${array.at(-1)}`;
}
