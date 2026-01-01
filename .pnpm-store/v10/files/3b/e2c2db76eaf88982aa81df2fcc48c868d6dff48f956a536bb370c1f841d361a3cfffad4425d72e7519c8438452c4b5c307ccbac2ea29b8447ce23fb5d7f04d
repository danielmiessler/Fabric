/** @import { Asset, RouteId, Pathname, ResolvedPathname } from '$app/types' */
/** @import { ResolveArgs } from './types.js' */
import { base, assets, hash_routing } from './internal/client.js';
import { resolve_route } from '../../../utils/routing.js';

/**
 * Resolve the URL of an asset in your `static` directory, by prefixing it with [`config.kit.paths.assets`](https://svelte.dev/docs/kit/configuration#paths) if configured, or otherwise by prefixing it with the base path.
 *
 * During server rendering, the base path is relative and depends on the page currently being rendered.
 *
 * @example
 * ```svelte
 * <script>
 * 	import { asset } from '$app/paths';
 * </script>
 *
 * <img alt="a potato" src={asset('/potato.jpg')} />
 * ```
 * @since 2.26
 *
 * @param {Asset} file
 * @returns {string}
 */
export function asset(file) {
	return (assets || base) + file;
}

const pathname_prefix = hash_routing ? '#' : '';

/**
 * Resolve a pathname by prefixing it with the base path, if any, or resolve a route ID by populating dynamic segments with parameters.
 *
 * During server rendering, the base path is relative and depends on the page currently being rendered.
 *
 * @example
 * ```js
 * import { resolve } from '$app/paths';
 *
 * // using a pathname
 * const resolved = resolve(`/blog/hello-world`);
 *
 * // using a route ID plus parameters
 * const resolved = resolve('/blog/[slug]', {
 * 	slug: 'hello-world'
 * });
 * ```
 * @since 2.26
 *
 * @template {RouteId | Pathname} T
 * @param {ResolveArgs<T>} args
 * @returns {ResolvedPathname}
 */
export function resolve(...args) {
	// The type error is correct here, and if someone doesn't pass params when they should there's a runtime error,
	// but we don't want to adjust the internal resolve_route function to accept `undefined`, hence the type cast.
	return (
		base + pathname_prefix + resolve_route(args[0], /** @type {Record<string, string>} */ (args[1]))
	);
}

export { base, assets, resolve as resolveRoute };
