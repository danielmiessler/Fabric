import * as devalue from 'devalue';
import { compact } from '../../../utils/array.js';
import { create_async_iterator } from '../../../utils/streaming.js';
import {
	clarify_devalue_error,
	get_global_name,
	handle_error_and_jsonify,
	serialize_uses
} from '../utils.js';

/**
 * If the serialized data contains promises, `chunks` will be an
 * async iterable containing their resolutions
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {import('types').SSROptions} options
 * @returns {import('./types.js').ServerDataSerializer}
 */
export function server_data_serializer(event, event_state, options) {
	let promise_id = 1;
	let max_nodes = -1;

	const iterator = create_async_iterator();
	const global = get_global_name(options);

	/** @param {number} index */
	function get_replacer(index) {
		/** @param {any} thing */
		return function replacer(thing) {
			if (typeof thing?.then === 'function') {
				const id = promise_id++;

				const promise = thing
					.then(/** @param {any} data */ (data) => ({ data }))
					.catch(
						/** @param {any} error */ async (error) => ({
							error: await handle_error_and_jsonify(event, event_state, options, error)
						})
					)
					.then(
						/**
						 * @param {{data: any; error: any}} result
						 */
						async ({ data, error }) => {
							let str;
							try {
								str = devalue.uneval(error ? [, error] : [data], replacer);
							} catch {
								error = await handle_error_and_jsonify(
									event,
									event_state,
									options,
									new Error(`Failed to serialize promise while rendering ${event.route.id}`)
								);
								data = undefined;
								str = devalue.uneval([, error], replacer);
							}

							return {
								index,
								str: `${global}.resolve(${id}, ${str.includes('app.decode') ? `(app) => ${str}` : `() => ${str}`})`
							};
						}
					);

				iterator.add(promise);

				return `${global}.defer(${id})`;
			} else {
				for (const key in options.hooks.transport) {
					const encoded = options.hooks.transport[key].encode(thing);
					if (encoded) {
						return `app.decode('${key}', ${devalue.uneval(encoded, replacer)})`;
					}
				}
			}
		};
	}

	const strings = /** @type {string[]} */ ([]);

	return {
		set_max_nodes(i) {
			max_nodes = i;
		},

		add_node(i, node) {
			try {
				if (!node) {
					strings[i] = 'null';
					return;
				}

				/** @type {any} */
				const payload = { type: 'data', data: node.data, uses: serialize_uses(node) };
				if (node.slash) payload.slash = node.slash;

				strings[i] = devalue.uneval(payload, get_replacer(i));
			} catch (e) {
				// @ts-expect-error
				e.path = e.path.slice(1);
				throw new Error(clarify_devalue_error(event, /** @type {any} */ (e)));
			}
		},

		get_data(csp) {
			const open = `<script${csp.script_needs_nonce ? ` nonce="${csp.nonce}"` : ''}>`;
			const close = `</script>\n`;

			return {
				data: `[${compact(max_nodes > -1 ? strings.slice(0, max_nodes) : strings).join(',')}]`,
				chunks:
					promise_id > 1
						? iterator.iterate(({ index, str }) => {
								if (max_nodes > -1 && index >= max_nodes) {
									return '';
								}

								return open + str + close;
							})
						: null
			};
		}
	};
}

/**
 * If the serialized data contains promises, `chunks` will be an
 * async iterable containing their resolutions
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {import('types').SSROptions} options
 * @returns {import('./types.js').ServerDataSerializerJson}
 */
export function server_data_serializer_json(event, event_state, options) {
	let promise_id = 1;

	const iterator = create_async_iterator();

	const reducers = {
		...Object.fromEntries(
			Object.entries(options.hooks.transport).map(([key, value]) => [key, value.encode])
		),
		/** @param {any} thing */
		Promise: (thing) => {
			if (typeof thing?.then !== 'function') {
				return;
			}

			const id = promise_id++;

			/** @type {'data' | 'error'} */
			let key = 'data';

			const promise = thing
				.catch(
					/** @param {any} e */ async (e) => {
						key = 'error';
						return handle_error_and_jsonify(event, event_state, options, /** @type {any} */ (e));
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
								event_state,
								options,
								new Error(`Failed to serialize promise while rendering ${event.route.id}`)
							);

							key = 'error';
							str = devalue.stringify(error, reducers);
						}

						return `{"type":"chunk","id":${id},"${key}":${str}}\n`;
					}
				);

			iterator.add(promise);

			return id;
		}
	};

	const strings = /** @type {string[]} */ ([]);

	return {
		add_node(i, node) {
			try {
				if (!node) {
					strings[i] = 'null';
					return;
				}

				if (node.type === 'error' || node.type === 'skip') {
					strings[i] = JSON.stringify(node);
					return;
				}

				strings[i] =
					`{"type":"data","data":${devalue.stringify(node.data, reducers)},"uses":${JSON.stringify(
						serialize_uses(node)
					)}${node.slash ? `,"slash":${JSON.stringify(node.slash)}` : ''}}`;
			} catch (e) {
				// @ts-expect-error
				e.path = 'data' + e.path;
				throw new Error(clarify_devalue_error(event, /** @type {any} */ (e)));
			}
		},

		get_data() {
			return {
				data: `{"type":"data","nodes":[${strings.join(',')}]}\n`,
				chunks: promise_id > 1 ? iterator.iterate() : null
			};
		}
	};
}
