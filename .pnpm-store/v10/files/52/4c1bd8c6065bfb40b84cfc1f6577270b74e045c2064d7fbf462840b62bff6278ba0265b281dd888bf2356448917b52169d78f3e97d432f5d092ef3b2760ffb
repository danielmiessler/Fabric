import { text } from '@sveltejs/kit';
import { HttpError, SvelteKitError, Redirect } from '@sveltejs/kit/internal';
import { normalize_error } from '../../../utils/error.js';
import { once } from '../../../utils/functions.js';
import { server_data_serializer_json } from '../page/data_serializer.js';
import { load_server_data } from '../page/load_data.js';
import { handle_error_and_jsonify } from '../utils.js';
import { normalize_path } from '../../../utils/url.js';
import { text_encoder } from '../../utils.js';

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {import('types').SSRRoute} route
 * @param {import('types').SSROptions} options
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @param {import('types').SSRState} state
 * @param {boolean[] | undefined} invalidated_data_nodes
 * @param {import('types').TrailingSlash} trailing_slash
 * @returns {Promise<Response>}
 */
export async function render_data(
	event,
	event_state,
	route,
	options,
	manifest,
	state,
	invalidated_data_nodes,
	trailing_slash
) {
	if (!route.page) {
		// requesting /__data.json should fail for a +server.js
		return new Response(undefined, {
			status: 404
		});
	}

	try {
		const node_ids = [...route.page.layouts, route.page.leaf];
		const invalidated = invalidated_data_nodes ?? node_ids.map(() => true);

		let aborted = false;

		const url = new URL(event.url);
		url.pathname = normalize_path(url.pathname, trailing_slash);

		const new_event = { ...event, url };

		const functions = node_ids.map((n, i) => {
			return once(async () => {
				try {
					if (aborted) {
						return /** @type {import('types').ServerDataSkippedNode} */ ({
							type: 'skip'
						});
					}

					// == because it could be undefined (in dev) or null (in build, because of JSON.stringify)
					const node = n == undefined ? n : await manifest._.nodes[n]();
					// load this. for the child, return as is. for the final result, stream things
					return load_server_data({
						event: new_event,
						event_state,
						state,
						node,
						parent: async () => {
							/** @type {Record<string, any>} */
							const data = {};
							for (let j = 0; j < i; j += 1) {
								const parent = /** @type {import('types').ServerDataNode | null} */ (
									await functions[j]()
								);

								if (parent) {
									Object.assign(data, parent.data);
								}
							}
							return data;
						}
					});
				} catch (e) {
					aborted = true;
					throw e;
				}
			});
		});

		const promises = functions.map(async (fn, i) => {
			if (!invalidated[i]) {
				return /** @type {import('types').ServerDataSkippedNode} */ ({
					type: 'skip'
				});
			}

			return fn();
		});

		let length = promises.length;
		const nodes = await Promise.all(
			promises.map((p, i) =>
				p.catch(async (error) => {
					if (error instanceof Redirect) {
						throw error;
					}

					// Math.min because array isn't guaranteed to resolve in order
					length = Math.min(length, i + 1);

					return /** @type {import('types').ServerErrorNode} */ ({
						type: 'error',
						error: await handle_error_and_jsonify(event, event_state, options, error),
						status:
							error instanceof HttpError || error instanceof SvelteKitError
								? error.status
								: undefined
					});
				})
			)
		);

		const data_serializer = server_data_serializer_json(event, event_state, options);
		for (let i = 0; i < nodes.length; i++) data_serializer.add_node(i, nodes[i]);
		const { data, chunks } = data_serializer.get_data();

		if (!chunks) {
			// use a normal JSON response where possible, so we get `content-length`
			// and can use browser JSON devtools for easier inspecting
			return json_response(data);
		}

		return new Response(
			new ReadableStream({
				async start(controller) {
					controller.enqueue(text_encoder.encode(data));
					for await (const chunk of chunks) {
						controller.enqueue(text_encoder.encode(chunk));
					}
					controller.close();
				},

				type: 'bytes'
			}),
			{
				headers: {
					// we use a proprietary content type to prevent buffering.
					// the `text` prefix makes it inspectable
					'content-type': 'text/sveltekit-data',
					'cache-control': 'private, no-store'
				}
			}
		);
	} catch (e) {
		const error = normalize_error(e);

		if (error instanceof Redirect) {
			return redirect_json_response(error);
		} else {
			return json_response(await handle_error_and_jsonify(event, event_state, options, error), 500);
		}
	}
}

/**
 * @param {Record<string, any> | string} json
 * @param {number} [status]
 */
function json_response(json, status = 200) {
	return text(typeof json === 'string' ? json : JSON.stringify(json), {
		status,
		headers: {
			'content-type': 'application/json',
			'cache-control': 'private, no-store'
		}
	});
}

/**
 * @param {Redirect} redirect
 */
export function redirect_json_response(redirect) {
	return json_response(
		/** @type {import('types').ServerRedirectNode} */ ({
			type: 'redirect',
			location: redirect.location
		})
	);
}
