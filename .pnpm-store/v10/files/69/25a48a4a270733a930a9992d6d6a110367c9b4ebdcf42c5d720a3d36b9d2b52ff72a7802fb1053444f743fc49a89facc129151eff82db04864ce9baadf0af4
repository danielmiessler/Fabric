import { ENDPOINT_METHODS, PAGE_METHODS } from '../../constants.js';
import { negotiate } from '../../utils/http.js';
import { with_event } from '../app/server/event.js';
import { Redirect } from '../control.js';
import { method_not_allowed } from './utils.js';

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').SSREndpoint} mod
 * @param {import('types').SSRState} state
 * @returns {Promise<Response>}
 */
export async function render_endpoint(event, mod, state) {
	const method = /** @type {import('types').HttpMethod} */ (event.request.method);

	let handler = mod[method] || mod.fallback;

	if (method === 'HEAD' && !mod.HEAD && mod.GET) {
		handler = mod.GET;
	}

	if (!handler) {
		return method_not_allowed(mod, method);
	}

	const prerender = mod.prerender ?? state.prerender_default;

	if (prerender && (mod.POST || mod.PATCH || mod.PUT || mod.DELETE)) {
		throw new Error('Cannot prerender endpoints that have mutative methods');
	}

	if (state.prerendering && !state.prerendering.inside_reroute && !prerender) {
		if (state.depth > 0) {
			// if request came from a prerendered page, bail
			throw new Error(`${event.route.id} is not prerenderable`);
		} else {
			// if request came direct from the crawler, signal that
			// this route cannot be prerendered, but don't bail
			return new Response(undefined, { status: 204 });
		}
	}

	try {
		const response = await with_event(event, () =>
			handler(/** @type {import('@sveltejs/kit').RequestEvent<Record<string, any>>} */ (event))
		);

		if (!(response instanceof Response)) {
			throw new Error(
				`Invalid response from route ${event.url.pathname}: handler should return a Response object`
			);
		}

		if (state.prerendering && (!state.prerendering.inside_reroute || prerender)) {
			// The returned Response might have immutable Headers
			// so we should clone them before trying to mutate them.
			// We also need to clone the response body since it may be read twice during prerendering
			const cloned = new Response(response.clone().body, {
				status: response.status,
				statusText: response.statusText,
				headers: new Headers(response.headers)
			});
			cloned.headers.set('x-sveltekit-prerender', String(prerender));

			if (state.prerendering.inside_reroute && prerender) {
				// Without this, the route wouldn't be recorded as prerendered,
				// because there's nothing after reroute that would do that.
				cloned.headers.set(
					'x-sveltekit-routeid',
					encodeURI(/** @type {string} */ (event.route.id))
				);
				state.prerendering.dependencies.set(event.url.pathname, { response: cloned, body: null });
			} else {
				return cloned;
			}
		}

		return response;
	} catch (e) {
		if (e instanceof Redirect) {
			return new Response(undefined, {
				status: e.status,
				headers: { location: e.location }
			});
		}

		throw e;
	}
}

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 */
export function is_endpoint_request(event) {
	const { method, headers } = event.request;

	// These methods exist exclusively for endpoints
	if (ENDPOINT_METHODS.includes(method) && !PAGE_METHODS.includes(method)) {
		return true;
	}

	// use:enhance uses a custom header to disambiguate
	if (method === 'POST' && headers.get('x-sveltekit-action') === 'true') return false;

	// GET/POST requests may be for endpoints or pages. We prefer endpoints if this isn't a text/html request
	const accept = event.request.headers.get('accept') ?? '*/*';
	return negotiate(accept, ['*', 'text/html']) !== 'text/html';
}
