import { app_dir, base } from '$app/paths/internal/client';
import { version } from '__sveltekit/environment';
import * as devalue from 'devalue';
import { DEV } from 'esm-env';
import { app, remote_responses } from '../client.js';
import { create_remote_function, remote_request } from './shared.svelte.js';

// Initialize Cache API for prerender functions
const CACHE_NAME = DEV ? `sveltekit:${Date.now()}` : `sveltekit:${version}`;
/** @type {Cache | undefined} */
let prerender_cache;

const prerender_cache_ready = (async () => {
	if (typeof caches !== 'undefined') {
		try {
			prerender_cache = await caches.open(CACHE_NAME);

			// Clean up old cache versions
			const cache_names = await caches.keys();
			for (const cache_name of cache_names) {
				if (cache_name.startsWith('sveltekit:') && cache_name !== CACHE_NAME) {
					await caches.delete(cache_name);
				}
			}
		} catch (error) {
			console.warn('Failed to initialize SvelteKit cache:', error);
		}
	}
})();

/**
 * @template T
 * @implements {Partial<Promise<T>>}
 */
class Prerender {
	/** @type {Promise<T>} */
	#promise;

	#loading = $state(true);
	#ready = $state(false);

	/** @type {T | undefined} */
	#current = $state.raw();

	#error = $state.raw(undefined);

	/**
	 * @param {() => Promise<T>} fn
	 */
	constructor(fn) {
		this.#promise = fn().then(
			(value) => {
				this.#loading = false;
				this.#ready = true;
				this.#current = value;
				return value;
			},
			(error) => {
				this.#loading = false;
				this.#error = error;
				throw error;
			}
		);
	}

	/**
	 *
	 * @param {((value: any) => any) | null | undefined} onfulfilled
	 * @param {((reason: any) => any) | null | undefined} [onrejected]
	 * @returns
	 */
	then(onfulfilled, onrejected) {
		return this.#promise.then(onfulfilled, onrejected);
	}

	/**
	 * @param {((reason: any) => any) | null | undefined} onrejected
	 */
	catch(onrejected) {
		return this.#promise.catch(onrejected);
	}

	/**
	 * @param {(() => any) | null | undefined} onfinally
	 */
	finally(onfinally) {
		return this.#promise.finally(onfinally);
	}

	get current() {
		return this.#current;
	}

	get error() {
		return this.#error;
	}

	/**
	 * Returns true if the resource is loading.
	 */
	get loading() {
		return this.#loading;
	}

	/**
	 * Returns true once the resource has been loaded.
	 */
	get ready() {
		return this.#ready;
	}
}

/**
 * @param {string} url
 * @param {string} encoded
 */
function put(url, encoded) {
	return /** @type {Cache} */ (prerender_cache)
		.put(
			url,
			// We need to create a new response because the original response is already consumed
			new Response(encoded, {
				headers: {
					'Content-Type': 'application/json'
				}
			})
		)
		.catch(() => {
			// Nothing we can do here
		});
}

/**
 * @param {string} id
 */
export function prerender(id) {
	return create_remote_function(id, (cache_key, payload) => {
		return new Prerender(async () => {
			await prerender_cache_ready;

			const url = `${base}/${app_dir}/remote/${id}${payload ? `/${payload}` : ''}`;

			if (Object.hasOwn(remote_responses, cache_key)) {
				const data = remote_responses[cache_key];

				if (prerender_cache) {
					void put(url, devalue.stringify(data, app.encoders));
				}

				return data;
			}

			// Check the Cache API first
			if (prerender_cache) {
				try {
					const cached_response = await prerender_cache.match(url);

					if (cached_response) {
						const cached_result = await cached_response.text();
						return devalue.parse(cached_result, app.decoders);
					}
				} catch {
					// Nothing we can do here
				}
			}

			const encoded = await remote_request(url);

			// For successful prerender requests, save to cache
			if (prerender_cache) {
				void put(url, encoded);
			}

			return devalue.parse(encoded, app.decoders);
		});
	});
}
