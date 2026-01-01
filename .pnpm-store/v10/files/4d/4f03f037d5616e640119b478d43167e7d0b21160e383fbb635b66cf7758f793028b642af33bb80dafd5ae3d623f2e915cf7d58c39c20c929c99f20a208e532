import { HttpError, SvelteKitError, Redirect } from '../../control.js';
import { normalize_error } from '../../../utils/error.js';
import { once } from '../../../utils/functions.js';
import { load_server_data } from '../page/load_data.js';
import { clarify_devalue_error, handle_error_and_jsonify, serialize_uses } from '../utils.js';
import { normalize_path } from '../../../utils/url.js';
import { text } from '../../../exports/index.js';
import * as devalue from 'devalue';
import { create_async_iterator } from '../../../utils/streaming.js';

const encoder = new TextEncoder();

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
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
						error: await handle_error_and_jsonify(event, options, error),
						status:
							error instanceof HttpError || error instanceof SvelteKitError
								? error.status
								: undefined
					});
				})
			)
		);

		const { data, chunks } = get_data_json(event, options, nodes);

		if (!chunks) {
			// use a normal JSON response where possible, so we get `content-length`
			// and can use browser JSON devtools for easier inspecting
			return json_response(data);
		}

		return new Response(
			new ReadableStream({
				async start(controller) {
					controller.enqueue(encoder.encode(data));
					for await (const chunk of chunks) {
						controller.enqueue(encoder.encode(chunk));
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
			return json_response(await handle_error_and_jsonify(event, options, error), 500);
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
	return json_response({
		type: 'redirect',
		location: redirect.location
	});
}

/**
 * If the serialized data contains promises, `chunks` will be an
 * async iterable containing their resolutions
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').SSROptions} options
 * @param {Array<import('types').ServerDataSkippedNode | import('types').ServerDataNode | import('types').ServerErrorNode | null | undefined>} nodes
 *  @returns {{ data: string, chunks: AsyncIterable<string> | null }}
 */
export function get_data_json(event, options, nodes) {
	let promise_id = 1;
	let count = 0;

	const { iterator, push, done } = create_async_iterator();

	const reducers = {
		...Object.fromEntries(
			Object.entries(options.hooks.transport).map(([key, value]) => [key, value.encode])
		),
		/** @param {any} thing */
		Promise: (thing) => {
			if (typeof thing?.then === 'function') {
				const id = promise_id++;
				count += 1;

				/** @type {'data' | 'error'} */
				let key = 'data';

				thing
					.catch(
						/** @param {any} e */ async (e) => {
							key = 'error';
							return handle_error_and_jsonify(event, options, /** @type {any} */ (e));
						}
					)
					.then(
						/** @param {any} value */
						async (value) => {
							let str;
							try {
								str = devalue.stringify(value, reducers);
							} catch {
								const error = await handle_error_and_jsonify(
									event,
									options,
									new Error(`Failed to serialize promise while rendering ${event.route.id}`)
								);

								key = 'error';
								str = devalue.stringify(error, reducers);
							}

							count -= 1;

							push(`{"type":"chunk","id":${id},"${key}":${str}}\n`);
							if (count === 0) done();
						}
					);

				return id;
			}
		}
	};

	try {
		const strings = nodes.map((node) => {
			if (!node) return 'null';

			if (node.type === 'error' || node.type === 'skip') {
				return JSON.stringify(node);
			}

			return `{"type":"data","data":${devalue.stringify(node.data, reducers)},"uses":${JSON.stringify(
				serialize_uses(node)
			)}${node.slash ? `,"slash":${JSON.stringify(node.slash)}` : ''}}`;
		});

		return {
			data: `{"type":"data","nodes":[${strings.join(',')}]}\n`,
			chunks: count > 0 ? iterator : null
		};
	} catch (e) {
		// @ts-expect-error
		e.path = 'data' + e.path;
		throw new Error(clarify_devalue_error(event, /** @type {any} */ (e)));
	}
}
