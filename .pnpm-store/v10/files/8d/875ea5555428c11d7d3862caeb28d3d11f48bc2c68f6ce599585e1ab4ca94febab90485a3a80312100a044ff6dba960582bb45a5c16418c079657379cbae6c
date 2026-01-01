import { text } from '@sveltejs/kit';
import { Redirect } from '@sveltejs/kit/internal';
import { compact } from '../../../utils/array.js';
import { get_status, normalize_error } from '../../../utils/error.js';
import { add_data_suffix } from '../../pathname.js';
import { redirect_response, static_error_page, handle_error_and_jsonify } from '../utils.js';
import {
	handle_action_json_request,
	handle_action_request,
	is_action_json_request,
	is_action_request
} from './actions.js';
import { server_data_serializer, server_data_serializer_json } from './data_serializer.js';
import { load_data, load_server_data } from './load_data.js';
import { render_response } from './render.js';
import { respond_with_error } from './respond_with_error.js';
import { DEV } from 'esm-env';
import { get_remote_action, handle_remote_form_post } from '../remote.js';
import { PageNodes } from '../../../utils/page_nodes.js';

/**
 * The maximum request depth permitted before assuming we're stuck in an infinite loop
 */
const MAX_DEPTH = 10;

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {import('types').PageNodeIndexes} page
 * @param {import('types').SSROptions} options
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @param {import('types').SSRState} state
 * @param {import('../../../utils/page_nodes.js').PageNodes} nodes
 * @param {import('types').RequiredResolveOptions} resolve_opts
 * @returns {Promise<Response>}
 */
export async function render_page(
	event,
	event_state,
	page,
	options,
	manifest,
	state,
	nodes,
	resolve_opts
) {
	if (state.depth > MAX_DEPTH) {
		// infinite request cycle detected
		return text(`Not found: ${event.url.pathname}`, {
			status: 404 // TODO in some cases this should be 500. not sure how to differentiate
		});
	}

	if (is_action_json_request(event)) {
		const node = await manifest._.nodes[page.leaf]();
		return handle_action_json_request(event, event_state, options, node?.server);
	}

	try {
		const leaf_node = /** @type {import('types').SSRNode} */ (nodes.page());

		let status = 200;

		/** @type {import('@sveltejs/kit').ActionResult | undefined} */
		let action_result = undefined;

		if (is_action_request(event)) {
			const remote_id = get_remote_action(event.url);
			if (remote_id) {
				action_result = await handle_remote_form_post(event, event_state, manifest, remote_id);
			} else {
				// for action requests, first call handler in +page.server.js
				// (this also determines status code)
				action_result = await handle_action_request(event, event_state, leaf_node.server);
			}

			if (action_result?.type === 'redirect') {
				return redirect_response(action_result.status, action_result.location);
			}
			if (action_result?.type === 'error') {
				status = get_status(action_result.error);
			}
			if (action_result?.type === 'failure') {
				status = action_result.status;
			}
		}

		// it's crucial that we do this before returning the non-SSR response, otherwise
		// SvelteKit will erroneously believe that the path has been prerendered,
		// causing functions to be omitted from the manifest generated later
		const should_prerender = nodes.prerender();
		if (should_prerender) {
			const mod = leaf_node.server;
			if (mod?.actions) {
				throw new Error('Cannot prerender pages with actions');
			}
		} else if (state.prerendering) {
			// if the page isn't marked as prerenderable, then bail out at this point
			return new Response(undefined, {
				status: 204
			});
		}

		// if we fetch any endpoints while loading data for this page, they should
		// inherit the prerender option of the page
		state.prerender_default = should_prerender;

		const should_prerender_data = nodes.should_prerender_data();
		const data_pathname = add_data_suffix(event.url.pathname);

		/** @type {import('./types.js').Fetched[]} */
		const fetched = [];

		const ssr = nodes.ssr();
		const csr = nodes.csr();

		// renders an empty 'shell' page if SSR is turned off and if there is
		// no server data to prerender. As a result, the load functions and rendering
		// only occur client-side.
		if (ssr === false && !(state.prerendering && should_prerender_data)) {
			// if the user makes a request through a non-enhanced form, the returned value is lost
			// because there is no SSR or client-side handling of the response
			if (DEV && action_result && !event.request.headers.has('x-sveltekit-action')) {
				if (action_result.type === 'error') {
					console.warn(
						"The form action returned an error, but +error.svelte wasn't rendered because SSR is off. To get the error page with CSR, enhance your form with `use:enhance`. See https://svelte.dev/docs/kit/form-actions#progressive-enhancement-use-enhance"
					);
				} else if (action_result.data) {
					/// case: lost data
					console.warn(
						"The form action returned a value, but it isn't available in `page.form`, because SSR is off. To handle the returned value in CSR, enhance your form with `use:enhance`. See https://svelte.dev/docs/kit/form-actions#progressive-enhancement-use-enhance"
					);
				}
			}

			return await render_response({
				branch: [],
				fetched,
				page_config: {
					ssr: false,
					csr
				},
				status,
				error: null,
				event,
				event_state,
				options,
				manifest,
				state,
				resolve_opts,
				data_serializer: server_data_serializer(event, event_state, options)
			});
		}

		/** @type {Array<import('./types.js').Loaded | null>} */
		const branch = [];

		/** @type {Error | null} */
		let load_error = null;

		const data_serializer = server_data_serializer(event, event_state, options);
		const data_serializer_json =
			state.prerendering && should_prerender_data
				? server_data_serializer_json(event, event_state, options)
				: null;

		/** @type {Array<Promise<import('types').ServerDataNode | null>>} */
		const server_promises = nodes.data.map((node, i) => {
			if (load_error) {
				// if an error happens immediately, don't bother with the rest of the nodes
				throw load_error;
			}

			return Promise.resolve().then(async () => {
				try {
					if (node === leaf_node && action_result?.type === 'error') {
						// we wait until here to throw the error so that we can use
						// any nested +error.svelte components that were defined
						throw action_result.error;
					}

					const server_data = await load_server_data({
						event,
						event_state,
						state,
						node,
						parent: async () => {
							/** @type {Record<string, any>} */
							const data = {};
							for (let j = 0; j < i; j += 1) {
								const parent = await server_promises[j];
								if (parent) Object.assign(data, parent.data);
							}
							return data;
						}
					});

					if (node) {
						data_serializer.add_node(i, server_data);
					}

					data_serializer_json?.add_node(i, server_data);

					return server_data;
				} catch (e) {
					load_error = /** @type {Error} */ (e);
					throw load_error;
				}
			});
		});

		/** @type {Array<Promise<Record<string, any> | null>>} */
		const load_promises = nodes.data.map((node, i) => {
			if (load_error) throw load_error;
			return Promise.resolve().then(async () => {
				try {
					return await load_data({
						event,
						event_state,
						fetched,
						node,
						parent: async () => {
							const data = {};
							for (let j = 0; j < i; j += 1) {
								Object.assign(data, await load_promises[j]);
							}
							return data;
						},
						resolve_opts,
						server_data_promise: server_promises[i],
						state,
						csr
					});
				} catch (e) {
					load_error = /** @type {Error} */ (e);
					throw load_error;
				}
			});
		});

		// if we don't do this, rejections will be unhandled
		for (const p of server_promises) p.catch(() => {});
		for (const p of load_promises) p.catch(() => {});

		for (let i = 0; i < nodes.data.length; i += 1) {
			const node = nodes.data[i];

			if (node) {
				try {
					const server_data = await server_promises[i];
					const data = await load_promises[i];

					branch.push({ node, server_data, data });
				} catch (e) {
					const err = normalize_error(e);

					if (err instanceof Redirect) {
						if (state.prerendering && should_prerender_data) {
							const body = JSON.stringify({
								type: 'redirect',
								location: err.location
							});

							state.prerendering.dependencies.set(data_pathname, {
								response: text(body),
								body
							});
						}

						return redirect_response(err.status, err.location);
					}

					const status = get_status(err);
					const error = await handle_error_and_jsonify(event, event_state, options, err);

					while (i--) {
						if (page.errors[i]) {
							const index = /** @type {number} */ (page.errors[i]);
							const node = await manifest._.nodes[index]();

							let j = i;
							while (!branch[j]) j -= 1;

							data_serializer.set_max_nodes(j + 1);

							const layouts = compact(branch.slice(0, j + 1));
							const nodes = new PageNodes(layouts.map((layout) => layout.node));

							return await render_response({
								event,
								event_state,
								options,
								manifest,
								state,
								resolve_opts,
								page_config: {
									ssr: nodes.ssr(),
									csr: nodes.csr()
								},
								status,
								error,
								branch: layouts.concat({
									node,
									data: null,
									server_data: null
								}),
								fetched,
								data_serializer
							});
						}
					}

					// if we're still here, it means the error happened in the root layout,
					// which means we have to fall back to error.html
					return static_error_page(options, status, error.message);
				}
			} else {
				// push an empty slot so we can rewind past gaps to the
				// layout that corresponds with an +error.svelte page
				branch.push(null);
			}
		}

		if (state.prerendering && data_serializer_json) {
			// ndjson format
			let { data, chunks } = data_serializer_json.get_data();

			if (chunks) {
				for await (const chunk of chunks) {
					data += chunk;
				}
			}

			state.prerendering.dependencies.set(data_pathname, {
				response: text(data),
				body: data
			});
		}

		return await render_response({
			event,
			event_state,
			options,
			manifest,
			state,
			resolve_opts,
			page_config: {
				csr,
				ssr
			},
			status,
			error: null,
			branch: ssr === false ? [] : compact(branch),
			action_result,
			fetched,
			data_serializer:
				ssr === false ? server_data_serializer(event, event_state, options) : data_serializer
		});
	} catch (e) {
		// if we end up here, it means the data loaded successfully
		// but the page failed to render, or that a prerendering error occurred
		return await respond_with_error({
			event,
			event_state,
			options,
			manifest,
			state,
			status: 500,
			error: e,
			resolve_opts
		});
	}
}
