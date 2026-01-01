import { DEV } from 'esm-env';
import { disable_search, make_trackable } from '../../../utils/url.js';
import { validate_depends, validate_load_response } from '../../shared.js';
import { with_request_store, merge_tracing } from '@sveltejs/kit/internal/server';
import { record_span } from '../../telemetry/record_span.js';
import { base64_encode, text_decoder } from '../../utils.js';
import { NULL_BODY_STATUS } from '../constants.js';
import { get_node_type } from '../utils.js';

/**
 * Calls the user's server `load` function.
 * @param {{
 *   event: import('@sveltejs/kit').RequestEvent;
 *   event_state: import('types').RequestState;
 *   state: import('types').SSRState;
 *   node: import('types').SSRNode | undefined;
 *   parent: () => Promise<Record<string, any>>;
 * }} opts
 * @returns {Promise<import('types').ServerDataNode | null>}
 */
export async function load_server_data({ event, event_state, state, node, parent }) {
	if (!node?.server) return null;

	let is_tracking = true;

	const uses = {
		dependencies: new Set(),
		params: new Set(),
		parent: false,
		route: false,
		url: false,
		search_params: new Set()
	};

	const load = node.server.load;
	// TODO: shouldn't this be calculated using PageNodes? there could be a trailingSlash option on a layout
	const slash = node.server.trailingSlash;

	if (!load) {
		return { type: 'data', data: null, uses, slash };
	}

	const url = make_trackable(
		event.url,
		() => {
			if (DEV && done && !uses.url) {
				console.warn(
					`${node.server_id}: Accessing URL properties in a promise handler after \`load(...)\` has returned will not cause the function to re-run when the URL changes`
				);
			}

			if (is_tracking) {
				uses.url = true;
			}
		},
		(param) => {
			if (DEV && done && !uses.search_params.has(param)) {
				console.warn(
					`${node.server_id}: Accessing URL properties in a promise handler after \`load(...)\` has returned will not cause the function to re-run when the URL changes`
				);
			}

			if (is_tracking) {
				uses.search_params.add(param);
			}
		}
	);

	if (state.prerendering) {
		disable_search(url);
	}

	let done = false;

	const result = await record_span({
		name: 'sveltekit.load',
		attributes: {
			'sveltekit.load.node_id': node.server_id || 'unknown',
			'sveltekit.load.node_type': get_node_type(node.server_id),
			'sveltekit.load.environment': 'server',
			'http.route': event.route.id || 'unknown'
		},
		fn: async (current) => {
			const traced_event = merge_tracing(event, current);
			const result = await with_request_store({ event: traced_event, state: event_state }, () =>
				load.call(null, {
					...traced_event,
					fetch: (info, init) => {
						const url = new URL(info instanceof Request ? info.url : info, event.url);

						if (DEV && done && !uses.dependencies.has(url.href)) {
							console.warn(
								`${node.server_id}: Calling \`event.fetch(...)\` in a promise handler after \`load(...)\` has returned will not cause the function to re-run when the dependency is invalidated`
							);
						}

						// Note: server fetches are not added to uses.depends due to security concerns
						return event.fetch(info, init);
					},
					/** @param {string[]} deps */
					depends: (...deps) => {
						for (const dep of deps) {
							const { href } = new URL(dep, event.url);

							if (DEV) {
								validate_depends(node.server_id || 'missing route ID', dep);

								if (done && !uses.dependencies.has(href)) {
									console.warn(
										`${node.server_id}: Calling \`depends(...)\` in a promise handler after \`load(...)\` has returned will not cause the function to re-run when the dependency is invalidated`
									);
								}
							}

							uses.dependencies.add(href);
						}
					},
					params: new Proxy(event.params, {
						get: (target, key) => {
							if (DEV && done && typeof key === 'string' && !uses.params.has(key)) {
								console.warn(
									`${node.server_id}: Accessing \`params.${String(
										key
									)}\` in a promise handler after \`load(...)\` has returned will not cause the function to re-run when the param changes`
								);
							}

							if (is_tracking) {
								uses.params.add(key);
							}
							return target[/** @type {string} */ (key)];
						}
					}),
					parent: async () => {
						if (DEV && done && !uses.parent) {
							console.warn(
								`${node.server_id}: Calling \`parent(...)\` in a promise handler after \`load(...)\` has returned will not cause the function to re-run when parent data changes`
							);
						}

						if (is_tracking) {
							uses.parent = true;
						}
						return parent();
					},
					route: new Proxy(event.route, {
						get: (target, key) => {
							if (DEV && done && typeof key === 'string' && !uses.route) {
								console.warn(
									`${node.server_id}: Accessing \`route.${String(
										key
									)}\` in a promise handler after \`load(...)\` has returned will not cause the function to re-run when the route changes`
								);
							}

							if (is_tracking) {
								uses.route = true;
							}
							return target[/** @type {'id'} */ (key)];
						}
					}),
					url,
					untrack(fn) {
						is_tracking = false;
						try {
							return fn();
						} finally {
							is_tracking = true;
						}
					}
				})
			);

			return result;
		}
	});

	if (DEV) {
		validate_load_response(result, `in ${node.server_id}`);
	}

	done = true;

	return {
		type: 'data',
		data: result ?? null,
		uses,
		slash
	};
}

/**
 * Calls the user's `load` function.
 * @param {{
 *   event: import('@sveltejs/kit').RequestEvent;
 *   event_state: import('types').RequestState;
 *   fetched: import('./types.js').Fetched[];
 *   node: import('types').SSRNode | undefined;
 *   parent: () => Promise<Record<string, any>>;
 *   resolve_opts: import('types').RequiredResolveOptions;
 *   server_data_promise: Promise<import('types').ServerDataNode | null>;
 *   state: import('types').SSRState;
 *   csr: boolean;
 * }} opts
 * @returns {Promise<Record<string, any | Promise<any>> | null>}
 */
export async function load_data({
	event,
	event_state,
	fetched,
	node,
	parent,
	server_data_promise,
	state,
	resolve_opts,
	csr
}) {
	const server_data_node = await server_data_promise;

	const load = node?.universal?.load;

	if (!load) {
		return server_data_node?.data ?? null;
	}

	const result = await record_span({
		name: 'sveltekit.load',
		attributes: {
			'sveltekit.load.node_id': node.universal_id || 'unknown',
			'sveltekit.load.node_type': get_node_type(node.universal_id),
			'sveltekit.load.environment': 'server',
			'http.route': event.route.id || 'unknown'
		},
		fn: async (current) => {
			const traced_event = merge_tracing(event, current);
			return await with_request_store({ event: traced_event, state: event_state }, () =>
				load.call(null, {
					url: event.url,
					params: event.params,
					data: server_data_node?.data ?? null,
					route: event.route,
					fetch: create_universal_fetch(event, state, fetched, csr, resolve_opts),
					setHeaders: event.setHeaders,
					depends: () => {},
					parent,
					untrack: (fn) => fn(),
					tracing: traced_event.tracing
				})
			);
		}
	});

	if (DEV) {
		validate_load_response(result, `in ${node.universal_id}`);
	}

	return result ?? null;
}

/**
 * @param {Pick<import('@sveltejs/kit').RequestEvent, 'fetch' | 'url' | 'request' | 'route'>} event
 * @param {import('types').SSRState} state
 * @param {import('./types.js').Fetched[]} fetched
 * @param {boolean} csr
 * @param {Pick<Required<import('@sveltejs/kit').ResolveOptions>, 'filterSerializedResponseHeaders'>} resolve_opts
 * @returns {typeof fetch}
 */
export function create_universal_fetch(event, state, fetched, csr, resolve_opts) {
	/**
	 * @param {URL | RequestInfo} input
	 * @param {RequestInit} [init]
	 */
	const universal_fetch = async (input, init) => {
		const cloned_body = input instanceof Request && input.body ? input.clone().body : null;

		const cloned_headers =
			input instanceof Request && [...input.headers].length
				? new Headers(input.headers)
				: init?.headers;

		let response = await event.fetch(input, init);

		const url = new URL(input instanceof Request ? input.url : input, event.url);
		const same_origin = url.origin === event.url.origin;

		/** @type {import('types').PrerenderDependency} */
		let dependency;

		if (same_origin) {
			if (state.prerendering) {
				dependency = { response, body: null };
				state.prerendering.dependencies.set(url.pathname, dependency);
			}
		} else if (url.protocol === 'https:' || url.protocol === 'http:') {
			// simulate CORS errors and "no access to body in no-cors mode" server-side for consistency with client-side behaviour
			const mode = input instanceof Request ? input.mode : (init?.mode ?? 'cors');
			if (mode === 'no-cors') {
				response = new Response('', {
					status: response.status,
					statusText: response.statusText,
					headers: response.headers
				});
			} else {
				const acao = response.headers.get('access-control-allow-origin');
				if (!acao || (acao !== event.url.origin && acao !== '*')) {
					throw new Error(
						`CORS error: ${
							acao ? 'Incorrect' : 'No'
						} 'Access-Control-Allow-Origin' header is present on the requested resource`
					);
				}
			}
		}

		/** @type {ReadableStream<Uint8Array>} */
		let teed_body;

		const proxy = new Proxy(response, {
			get(response, key, receiver) {
				/**
				 * @param {string | undefined} body
				 * @param {boolean} is_b64
				 */
				async function push_fetched(body, is_b64) {
					const status_number = Number(response.status);
					if (isNaN(status_number)) {
						throw new Error(
							`response.status is not a number. value: "${
								response.status
							}" type: ${typeof response.status}`
						);
					}

					fetched.push({
						url: same_origin ? url.href.slice(event.url.origin.length) : url.href,
						method: event.request.method,
						request_body: /** @type {string | ArrayBufferView | undefined} */ (
							input instanceof Request && cloned_body
								? await stream_to_string(cloned_body)
								: init?.body
						),
						request_headers: cloned_headers,
						response_body: body,
						response,
						is_b64
					});
				}

				if (key === 'body') {
					if (response.body === null) {
						return null;
					}

					if (teed_body) {
						return teed_body;
					}

					const [a, b] = response.body.tee();

					void (async () => {
						let result = new Uint8Array();

						for await (const chunk of a) {
							const combined = new Uint8Array(result.length + chunk.length);

							combined.set(result, 0);
							combined.set(chunk, result.length);

							result = combined;
						}

						if (dependency) {
							dependency.body = new Uint8Array(result);
						}

						void push_fetched(base64_encode(result), true);
					})();

					return (teed_body = b);
				}

				if (key === 'arrayBuffer') {
					return async () => {
						const buffer = await response.arrayBuffer();

						const bytes = new Uint8Array(buffer);

						if (dependency) {
							dependency.body = bytes;
						}

						if (buffer instanceof ArrayBuffer) {
							await push_fetched(base64_encode(bytes), true);
						}

						return buffer;
					};
				}

				async function text() {
					const body = await response.text();

					if (body === '' && NULL_BODY_STATUS.includes(response.status)) {
						await push_fetched(undefined, false);
						return undefined;
					}

					if (!body || typeof body === 'string') {
						await push_fetched(body, false);
					}

					if (dependency) {
						dependency.body = body;
					}

					return body;
				}

				if (key === 'text') {
					return text;
				}

				if (key === 'json') {
					return async () => {
						const body = await text();
						return body ? JSON.parse(body) : undefined;
					};
				}

				const value = Reflect.get(response, key, response);

				if (value instanceof Function) {
					// On Node v24+, the Response object has a private element #state – we
					// need to bind this function to the response in order to allow it to
					// access this private element. Defining the name and length ensure it
					// is identical to the original function when introspected.
					return Object.defineProperties(
						/**
						 * @this {any}
						 */
						function () {
							return Reflect.apply(value, this === receiver ? response : this, arguments);
						},
						{
							name: { value: value.name },
							length: { value: value.length }
						}
					);
				}

				return value;
			}
		});

		if (csr) {
			// ensure that excluded headers can't be read
			const get = response.headers.get;
			response.headers.get = (key) => {
				const lower = key.toLowerCase();
				const value = get.call(response.headers, lower);
				if (value && !lower.startsWith('x-sveltekit-')) {
					const included = resolve_opts.filterSerializedResponseHeaders(lower, value);
					if (!included) {
						throw new Error(
							`Failed to get response header "${lower}" — it must be included by the \`filterSerializedResponseHeaders\` option: https://svelte.dev/docs/kit/hooks#Server-hooks-handle (at ${event.route.id})`
						);
					}
				}

				return value;
			};
		}

		return proxy;
	};

	// Don't make this function `async`! Otherwise, the user has to `catch` promises they use for streaming responses or else
	// it will be an unhandled rejection. Instead, we add a `.catch(() => {})` ourselves below to this from happening.
	return (input, init) => {
		// See docs in fetch.js for why we need to do this
		const response = universal_fetch(input, init);
		response.catch(() => {});
		return response;
	};
}

/**
 * @param {ReadableStream<Uint8Array>} stream
 */
async function stream_to_string(stream) {
	let result = '';
	const reader = stream.getReader();
	while (true) {
		const { done, value } = await reader.read();
		if (done) {
			break;
		}
		result += text_decoder.decode(value);
	}
	return result;
}
