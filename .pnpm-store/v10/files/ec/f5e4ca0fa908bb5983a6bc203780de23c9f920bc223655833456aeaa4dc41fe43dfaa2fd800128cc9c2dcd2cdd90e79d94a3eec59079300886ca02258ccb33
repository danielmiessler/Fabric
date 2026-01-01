import { BROWSER, DEV } from 'esm-env';
import { hash } from '../hash.js';
import { b64_decode } from '../utils.js';

let loading = 0;

/** @type {typeof fetch} */
const native_fetch = BROWSER ? window.fetch : /** @type {any} */ (() => {});

export function lock_fetch() {
	loading += 1;
}

export function unlock_fetch() {
	loading -= 1;
}

if (DEV && BROWSER) {
	let can_inspect_stack_trace = false;

	// detect whether async stack traces work
	// eslint-disable-next-line @typescript-eslint/require-await
	const check_stack_trace = async () => {
		const stack = /** @type {string} */ (new Error().stack);
		can_inspect_stack_trace = stack.includes('check_stack_trace');
	};

	void check_stack_trace();

	/**
	 * @param {RequestInfo | URL} input
	 * @param {RequestInit & Record<string, any> | undefined} init
	 */
	window.fetch = (input, init) => {
		// Check if fetch was called via load_node. the lock method only checks if it was called at the
		// same time, but not necessarily if it was called from `load`.
		// We use just the filename as the method name sometimes does not appear on the CI.
		const url = input instanceof Request ? input.url : input.toString();
		const stack_array = /** @type {string} */ (new Error().stack).split('\n');
		// We need to do a cutoff because Safari and Firefox maintain the stack
		// across events and for example traces a `fetch` call triggered from a button
		// back to the creation of the event listener and the element creation itself,
		// where at some point client.js will show up, leading to false positives.
		const cutoff = stack_array.findIndex((a) => a.includes('load@') || a.includes('at load'));
		const stack = stack_array.slice(0, cutoff + 2).join('\n');

		const in_load_heuristic = can_inspect_stack_trace
			? stack.includes('src/runtime/client/client.js')
			: loading;

		// This flag is set in initial_fetch and subsequent_fetch
		const used_kit_fetch = init?.__sveltekit_fetch__;

		if (in_load_heuristic && !used_kit_fetch) {
			console.warn(
				`Loading ${url} using \`window.fetch\`. For best results, use the \`fetch\` that is passed to your \`load\` function: https://svelte.dev/docs/kit/load#making-fetch-requests`
			);
		}

		const method = input instanceof Request ? input.method : init?.method || 'GET';

		if (method !== 'GET') {
			cache.delete(build_selector(input));
		}

		return native_fetch(input, init);
	};
} else if (BROWSER) {
	window.fetch = (input, init) => {
		const method = input instanceof Request ? input.method : init?.method || 'GET';

		if (method !== 'GET') {
			cache.delete(build_selector(input));
		}

		return native_fetch(input, init);
	};
}

const cache = new Map();

/**
 * Should be called on the initial run of load functions that hydrate the page.
 * Saves any requests with cache-control max-age to the cache.
 * @param {URL | string} resource
 * @param {RequestInit} [opts]
 */
export function initial_fetch(resource, opts) {
	const selector = build_selector(resource, opts);

	const script = document.querySelector(selector);
	if (script?.textContent) {
		let { body, ...init } = JSON.parse(script.textContent);

		const ttl = script.getAttribute('data-ttl');
		if (ttl) cache.set(selector, { body, init, ttl: 1000 * Number(ttl) });
		const b64 = script.getAttribute('data-b64');
		if (b64 !== null) {
			// Can't use native_fetch('data:...;base64,${body}')
			// csp can block the request
			body = b64_decode(body);
		}

		return Promise.resolve(new Response(body, init));
	}

	return DEV ? dev_fetch(resource, opts) : window.fetch(resource, opts);
}

/**
 * Tries to get the response from the cache, if max-age allows it, else does a fetch.
 * @param {URL | string} resource
 * @param {string} resolved
 * @param {RequestInit} [opts]
 */
export function subsequent_fetch(resource, resolved, opts) {
	if (cache.size > 0) {
		const selector = build_selector(resource, opts);
		const cached = cache.get(selector);
		if (cached) {
			// https://developer.mozilla.org/en-US/docs/Web/API/Request/cache#value
			if (
				performance.now() < cached.ttl &&
				['default', 'force-cache', 'only-if-cached', undefined].includes(opts?.cache)
			) {
				return new Response(cached.body, cached.init);
			}

			cache.delete(selector);
		}
	}

	return DEV ? dev_fetch(resolved, opts) : window.fetch(resolved, opts);
}

/**
 * @param {RequestInfo | URL} resource
 * @param {RequestInit & Record<string, any> | undefined} opts
 */
export function dev_fetch(resource, opts) {
	const patched_opts = { ...opts };
	// This assigns the __sveltekit_fetch__ flag and makes it non-enumerable
	Object.defineProperty(patched_opts, '__sveltekit_fetch__', {
		value: true,
		writable: true,
		configurable: true
	});
	return window.fetch(resource, patched_opts);
}

/**
 * Build the cache key for a given request
 * @param {URL | RequestInfo} resource
 * @param {RequestInit} [opts]
 */
function build_selector(resource, opts) {
	const url = JSON.stringify(resource instanceof Request ? resource.url : resource);

	let selector = `script[data-sveltekit-fetched][data-url=${url}]`;

	if (opts?.headers || opts?.body) {
		/** @type {import('types').StrictBody[]} */
		const values = [];

		if (opts.headers) {
			values.push([...new Headers(opts.headers)].join(','));
		}

		if (opts.body && (typeof opts.body === 'string' || ArrayBuffer.isView(opts.body))) {
			values.push(opts.body);
		}

		selector += `[data-hash="${hash(...values)}"]`;
	}

	return selector;
}
