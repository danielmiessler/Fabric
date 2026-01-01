/** @import { RequestEvent, ActionResult, Actions } from '@sveltejs/kit' */
/** @import { SSROptions, SSRNode, ServerNode, ServerHooks } from 'types' */
import * as devalue from 'devalue';
import { DEV } from 'esm-env';
import { json } from '@sveltejs/kit';
import { HttpError, Redirect, ActionFailure, SvelteKitError } from '@sveltejs/kit/internal';
import { with_request_store, merge_tracing } from '@sveltejs/kit/internal/server';
import { get_status, normalize_error } from '../../../utils/error.js';
import { is_form_content_type, negotiate } from '../../../utils/http.js';
import { handle_error_and_jsonify } from '../utils.js';
import { record_span } from '../../telemetry/record_span.js';

/** @param {RequestEvent} event */
export function is_action_json_request(event) {
	const accept = negotiate(event.request.headers.get('accept') ?? '*/*', [
		'application/json',
		'text/html'
	]);

	return accept === 'application/json' && event.request.method === 'POST';
}

/**
 * @param {RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {SSROptions} options
 * @param {SSRNode['server'] | undefined} server
 */
export async function handle_action_json_request(event, event_state, options, server) {
	const actions = server?.actions;

	if (!actions) {
		const no_actions_error = new SvelteKitError(
			405,
			'Method Not Allowed',
			`POST method not allowed. No form actions exist for ${DEV ? `the page at ${event.route.id}` : 'this page'}`
		);

		return action_json(
			{
				type: 'error',
				error: await handle_error_and_jsonify(event, event_state, options, no_actions_error)
			},
			{
				status: no_actions_error.status,
				headers: {
					// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
					// "The server must generate an Allow header field in a 405 status code response"
					allow: 'GET'
				}
			}
		);
	}

	check_named_default_separate(actions);

	try {
		const data = await call_action(event, event_state, actions);

		if (DEV) {
			validate_action_return(data);
		}

		if (data instanceof ActionFailure) {
			return action_json({
				type: 'failure',
				status: data.status,
				// @ts-expect-error we assign a string to what is supposed to be an object. That's ok
				// because we don't use the object outside, and this way we have better code navigation
				// through knowing where the related interface is used.
				data: stringify_action_response(
					data.data,
					/** @type {string} */ (event.route.id),
					options.hooks.transport
				)
			});
		} else {
			return action_json({
				type: 'success',
				status: data ? 200 : 204,
				// @ts-expect-error see comment above
				data: stringify_action_response(
					data,
					/** @type {string} */ (event.route.id),
					options.hooks.transport
				)
			});
		}
	} catch (e) {
		const err = normalize_error(e);

		if (err instanceof Redirect) {
			return action_json_redirect(err);
		}

		return action_json(
			{
				type: 'error',
				error: await handle_error_and_jsonify(
					event,
					event_state,
					options,
					check_incorrect_fail_use(err)
				)
			},
			{
				status: get_status(err)
			}
		);
	}
}

/**
 * @param {HttpError | Error} error
 */
export function check_incorrect_fail_use(error) {
	return error instanceof ActionFailure
		? new Error('Cannot "throw fail()". Use "return fail()"')
		: error;
}

/**
 * @param {Redirect} redirect
 */
export function action_json_redirect(redirect) {
	return action_json({
		type: 'redirect',
		status: redirect.status,
		location: redirect.location
	});
}

/**
 * @param {ActionResult} data
 * @param {ResponseInit} [init]
 */
function action_json(data, init) {
	return json(data, init);
}

/**
 * @param {RequestEvent} event
 */
export function is_action_request(event) {
	return event.request.method === 'POST';
}

/**
 * @param {RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {SSRNode['server'] | undefined} server
 * @returns {Promise<ActionResult>}
 */
export async function handle_action_request(event, event_state, server) {
	const actions = server?.actions;

	if (!actions) {
		// TODO should this be a different error altogether?
		event.setHeaders({
			// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
			// "The server must generate an Allow header field in a 405 status code response"
			allow: 'GET'
		});
		return {
			type: 'error',
			error: new SvelteKitError(
				405,
				'Method Not Allowed',
				`POST method not allowed. No form actions exist for ${DEV ? `the page at ${event.route.id}` : 'this page'}`
			)
		};
	}

	check_named_default_separate(actions);

	try {
		const data = await call_action(event, event_state, actions);

		if (DEV) {
			validate_action_return(data);
		}

		if (data instanceof ActionFailure) {
			return {
				type: 'failure',
				status: data.status,
				data: data.data
			};
		} else {
			return {
				type: 'success',
				status: 200,
				// @ts-expect-error this will be removed upon serialization, so `undefined` is the same as omission
				data
			};
		}
	} catch (e) {
		const err = normalize_error(e);

		if (err instanceof Redirect) {
			return {
				type: 'redirect',
				status: err.status,
				location: err.location
			};
		}

		return {
			type: 'error',
			error: check_incorrect_fail_use(err)
		};
	}
}

/**
 * @param {Actions} actions
 */
function check_named_default_separate(actions) {
	if (actions.default && Object.keys(actions).length > 1) {
		throw new Error(
			'When using named actions, the default action cannot be used. See the docs for more info: https://svelte.dev/docs/kit/form-actions#named-actions'
		);
	}
}

/**
 * @param {RequestEvent} event
 * @param {import('types').RequestState} event_state
 * @param {NonNullable<ServerNode['actions']>} actions
 * @throws {Redirect | HttpError | SvelteKitError | Error}
 */
async function call_action(event, event_state, actions) {
	const url = new URL(event.request.url);

	let name = 'default';
	for (const param of url.searchParams) {
		if (param[0].startsWith('/')) {
			name = param[0].slice(1);
			if (name === 'default') {
				throw new Error('Cannot use reserved action name "default"');
			}
			break;
		}
	}

	const action = actions[name];
	if (!action) {
		throw new SvelteKitError(404, 'Not Found', `No action with name '${name}' found`);
	}

	if (!is_form_content_type(event.request)) {
		throw new SvelteKitError(
			415,
			'Unsupported Media Type',
			`Form actions expect form-encoded data — received ${event.request.headers.get(
				'content-type'
			)}`
		);
	}

	return record_span({
		name: 'sveltekit.form_action',
		attributes: {
			'sveltekit.form_action.name': name,
			'http.route': event.route.id || 'unknown'
		},
		fn: async (current) => {
			const traced_event = merge_tracing(event, current);
			const result = await with_request_store({ event: traced_event, state: event_state }, () =>
				action(traced_event)
			);
			if (result instanceof ActionFailure) {
				current.setAttributes({
					'sveltekit.form_action.result.type': 'failure',
					'sveltekit.form_action.result.status': result.status
				});
			}

			return result;
		}
	});
}

/** @param {any} data */
function validate_action_return(data) {
	if (data instanceof Redirect) {
		throw new Error('Cannot `return redirect(...)` — use `redirect(...)` instead');
	}

	if (data instanceof HttpError) {
		throw new Error('Cannot `return error(...)` — use `error(...)` or `return fail(...)` instead');
	}
}

/**
 * Try to `devalue.uneval` the data object, and if it fails, return a proper Error with context
 * @param {any} data
 * @param {string} route_id
 * @param {ServerHooks['transport']} transport
 */
export function uneval_action_response(data, route_id, transport) {
	const replacer = (/** @type {any} */ thing) => {
		for (const key in transport) {
			const encoded = transport[key].encode(thing);
			if (encoded) {
				return `app.decode('${key}', ${devalue.uneval(encoded, replacer)})`;
			}
		}
	};

	return try_serialize(data, (value) => devalue.uneval(value, replacer), route_id);
}

/**
 * Try to `devalue.stringify` the data object, and if it fails, return a proper Error with context
 * @param {any} data
 * @param {string} route_id
 * @param {ServerHooks['transport']} transport
 */
function stringify_action_response(data, route_id, transport) {
	const encoders = Object.fromEntries(
		Object.entries(transport).map(([key, value]) => [key, value.encode])
	);

	return try_serialize(data, (value) => devalue.stringify(value, encoders), route_id);
}

/**
 * @param {any} data
 * @param {(data: any) => string} fn
 * @param {string} route_id
 */
function try_serialize(data, fn, route_id) {
	try {
		return fn(data);
	} catch (e) {
		// If we're here, the data could not be serialized with devalue
		const error = /** @type {any} */ (e);

		// if someone tries to use `json()` in their action
		if (data instanceof Response) {
			throw new Error(
				`Data returned from action inside ${route_id} is not serializable. Form actions need to return plain objects or fail(). E.g. return { success: true } or return fail(400, { message: "invalid" });`
			);
		}

		// if devalue could not serialize a property on the object, etc.
		if ('path' in error) {
			let message = `Data returned from action inside ${route_id} is not serializable: ${error.message}`;
			if (error.path !== '') message += ` (data.${error.path})`;
			throw new Error(message);
		}

		throw error;
	}
}
