import { base, assets, relative, initial_base } from './internal/server.js';
import { resolve_route } from '../../../utils/routing.js';
import { try_get_request_store } from '@sveltejs/kit/internal/server';

/** @type {import('./client.js').asset} */
export function asset(file) {
	// @ts-expect-error we use the `resolve` mechanism, but with the 'wrong' input
	return assets ? assets + file : resolve(file);
}

/** @type {import('./client.js').resolve} */
export function resolve(id, params) {
	const resolved = resolve_route(id, /** @type {Record<string, string>} */ (params));

	if (relative) {
		const store = try_get_request_store();

		if (store && !store.state.prerendering?.fallback) {
			const after_base = store.event.url.pathname.slice(initial_base.length);
			const segments = after_base.split('/').slice(2);
			const prefix = segments.map(() => '..').join('/') || '.';

			return prefix + resolved;
		}
	}

	return base + resolved;
}

export { base, assets, resolve as resolveRoute };
