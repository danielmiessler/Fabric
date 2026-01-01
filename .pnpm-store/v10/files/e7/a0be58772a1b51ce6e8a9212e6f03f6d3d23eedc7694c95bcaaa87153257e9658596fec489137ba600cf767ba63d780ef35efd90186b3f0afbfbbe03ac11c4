/**
 * A helper function for sequencing multiple `handle` calls in a middleware-like manner.
 * The behavior for the `handle` options is as follows:
 * - `transformPageChunk` is applied in reverse order and merged
 * - `preload` is applied in forward order, the first option "wins" and no `preload` options after it are called
 * - `filterSerializedResponseHeaders` behaves the same as `preload`
 *
 * ```js
 * /// file: src/hooks.server.js
 * import { sequence } from '@sveltejs/kit/hooks';
 *
 * /// type: import('@sveltejs/kit').Handle
 * async function first({ event, resolve }) {
 * 	console.log('first pre-processing');
 * 	const result = await resolve(event, {
 * 		transformPageChunk: ({ html }) => {
 * 			// transforms are applied in reverse order
 * 			console.log('first transform');
 * 			return html;
 * 		},
 * 		preload: () => {
 * 			// this one wins as it's the first defined in the chain
 * 			console.log('first preload');
 * 			return true;
 * 		}
 * 	});
 * 	console.log('first post-processing');
 * 	return result;
 * }
 *
 * /// type: import('@sveltejs/kit').Handle
 * async function second({ event, resolve }) {
 * 	console.log('second pre-processing');
 * 	const result = await resolve(event, {
 * 		transformPageChunk: ({ html }) => {
 * 			console.log('second transform');
 * 			return html;
 * 		},
 * 		preload: () => {
 * 			console.log('second preload');
 * 			return true;
 * 		},
 * 		filterSerializedResponseHeaders: () => {
 * 			// this one wins as it's the first defined in the chain
 * 			console.log('second filterSerializedResponseHeaders');
 * 			return true;
 * 		}
 * 	});
 * 	console.log('second post-processing');
 * 	return result;
 * }
 *
 * export const handle = sequence(first, second);
 * ```
 *
 * The example above would print:
 *
 * ```
 * first pre-processing
 * first preload
 * second pre-processing
 * second filterSerializedResponseHeaders
 * second transform
 * first transform
 * second post-processing
 * first post-processing
 * ```
 *
 * @param {...import('@sveltejs/kit').Handle} handlers The chain of `handle` functions
 * @returns {import('@sveltejs/kit').Handle}
 */
export function sequence(...handlers) {
	const length = handlers.length;
	if (!length) return ({ event, resolve }) => resolve(event);

	return ({ event, resolve }) => {
		return apply_handle(0, event, {});

		/**
		 * @param {number} i
		 * @param {import('@sveltejs/kit').RequestEvent} event
		 * @param {import('@sveltejs/kit').ResolveOptions | undefined} parent_options
		 * @returns {import('types').MaybePromise<Response>}
		 */
		function apply_handle(i, event, parent_options) {
			const handle = handlers[i];

			return handle({
				event,
				resolve: (event, options) => {
					/** @type {import('@sveltejs/kit').ResolveOptions['transformPageChunk']} */
					const transformPageChunk = async ({ html, done }) => {
						if (options?.transformPageChunk) {
							html = (await options.transformPageChunk({ html, done })) ?? '';
						}

						if (parent_options?.transformPageChunk) {
							html = (await parent_options.transformPageChunk({ html, done })) ?? '';
						}

						return html;
					};

					/** @type {import('@sveltejs/kit').ResolveOptions['filterSerializedResponseHeaders']} */
					const filterSerializedResponseHeaders =
						parent_options?.filterSerializedResponseHeaders ??
						options?.filterSerializedResponseHeaders;

					/** @type {import('@sveltejs/kit').ResolveOptions['preload']} */
					const preload = parent_options?.preload ?? options?.preload;

					return i < length - 1
						? apply_handle(i + 1, event, {
								transformPageChunk,
								filterSerializedResponseHeaders,
								preload
							})
						: resolve(event, { transformPageChunk, filterSerializedResponseHeaders, preload });
				}
			});
		}
	};
}
