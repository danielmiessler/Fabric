/** @import { RequestEvent } from '@sveltejs/kit' */
/** @import { ServerHooks, MaybePromise, RequestState, RemoteInfo, RequestStore } from 'types' */
import { parse } from 'devalue';
import { error } from '@sveltejs/kit';
import { with_request_store, get_request_store } from '@sveltejs/kit/internal/server';
import { stringify_remote_arg } from '../../../shared.js';

/**
 * @param {any} validate_or_fn
 * @param {(arg?: any) => any} [maybe_fn]
 * @returns {(arg?: any) => MaybePromise<any>}
 */
export function create_validator(validate_or_fn, maybe_fn) {
	// prevent functions without validators being called with arguments
	if (!maybe_fn) {
		return (arg) => {
			if (arg !== undefined) {
				error(400, 'Bad Request');
			}
		};
	}

	// if 'unchecked', pass input through without validating
	if (validate_or_fn === 'unchecked') {
		return (arg) => arg;
	}

	// use https://standardschema.dev validator if provided
	if ('~standard' in validate_or_fn) {
		return async (arg) => {
			// Get event before async validation to ensure it's available in server environments without AsyncLocalStorage, too
			const { event, state } = get_request_store();
			// access property and call method in one go to preserve potential this context
			const result = await validate_or_fn['~standard'].validate(arg);

			// if the `issues` field exists, the validation failed
			if (result.issues) {
				error(
					400,
					await state.handleValidationError({
						issues: result.issues,
						event
					})
				);
			}

			return result.value;
		};
	}

	throw new Error(
		'Invalid validator passed to remote function. Expected "unchecked" or a Standard Schema (https://standardschema.dev)'
	);
}

/**
 * In case of a single remote function call, just returns the result.
 *
 * In case of a full page reload, returns the response for a remote function call,
 * either from the cache or by invoking the function.
 * Also saves an uneval'ed version of the result for later HTML inlining for hydration.
 *
 * @template {MaybePromise<any>} T
 * @param {RemoteInfo} info
 * @param {any} arg
 * @param {RequestState} state
 * @param {() => Promise<T>} get_result
 * @returns {Promise<T>}
 */
export async function get_response(info, arg, state, get_result) {
	// wait a beat, in case `myQuery().set(...)` or `myQuery().refresh()` is immediately called
	// eslint-disable-next-line @typescript-eslint/await-thenable
	await 0;

	const cache = get_cache(info, state);

	return (cache[stringify_remote_arg(arg, state.transport)] ??= get_result());
}

/**
 * @param {any} data
 * @param {ServerHooks['transport']} transport
 */
export function parse_remote_response(data, transport) {
	/** @type {Record<string, any>} */
	const revivers = {};
	for (const key in transport) {
		revivers[key] = transport[key].decode;
	}

	return parse(data, revivers);
}

/**
 * Like `with_event` but removes things from `event` you cannot see/call in remote functions, such as `setHeaders`.
 * @template T
 * @param {RequestEvent} event
 * @param {RequestState} state
 * @param {boolean} allow_cookies
 * @param {any} arg
 * @param {(arg: any) => any} validate
 * @param {(arg?: any) => T} fn
 */
export async function run_remote_function(event, state, allow_cookies, arg, validate, fn) {
	/** @type {RequestStore} */
	const store = {
		event: {
			...event,
			setHeaders: () => {
				throw new Error('setHeaders is not allowed in remote functions');
			},
			cookies: {
				...event.cookies,
				set: (name, value, opts) => {
					if (!allow_cookies) {
						throw new Error('Cannot set cookies in `query` or `prerender` functions');
					}

					if (opts.path && !opts.path.startsWith('/')) {
						throw new Error('Cookies set in remote functions must have an absolute path');
					}

					return event.cookies.set(name, value, opts);
				},
				delete: (name, opts) => {
					if (!allow_cookies) {
						throw new Error('Cannot delete cookies in `query` or `prerender` functions');
					}

					if (opts.path && !opts.path.startsWith('/')) {
						throw new Error('Cookies deleted in remote functions must have an absolute path');
					}

					return event.cookies.delete(name, opts);
				}
			}
		},
		state: {
			...state,
			is_in_remote_function: true
		}
	};

	// In two parts, each with_event, so that runtimes without async local storage can still get the event at the start of the function
	const validated = await with_request_store(store, () => validate(arg));
	return with_request_store(store, () => fn(validated));
}

/**
 * @param {RemoteInfo} info
 * @param {RequestState} state
 */
export function get_cache(info, state = get_request_store().state) {
	let cache = state.remote_data?.get(info);

	if (cache === undefined) {
		cache = {};
		(state.remote_data ??= new Map()).set(info, cache);
	}

	return cache;
}
