/** @import { RemoteQuery, RemoteQueryFunction } from '@sveltejs/kit' */
/** @import { RemoteInfo, MaybePromise } from 'types' */
/** @import { StandardSchemaV1 } from '@standard-schema/spec' */
import { get_request_store } from '@sveltejs/kit/internal/server';
import { create_remote_key, stringify_remote_arg } from '../../../shared.js';
import { prerendering } from '__sveltekit/environment';
import { create_validator, get_cache, get_response, run_remote_function } from './shared.js';

/**
 * Creates a remote query. When called from the browser, the function will be invoked on the server via a `fetch` call.
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#query) for full documentation.
 *
 * @template Output
 * @overload
 * @param {() => MaybePromise<Output>} fn
 * @returns {RemoteQueryFunction<void, Output>}
 * @since 2.27
 */
/**
 * Creates a remote query. When called from the browser, the function will be invoked on the server via a `fetch` call.
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#query) for full documentation.
 *
 * @template Input
 * @template Output
 * @overload
 * @param {'unchecked'} validate
 * @param {(arg: Input) => MaybePromise<Output>} fn
 * @returns {RemoteQueryFunction<Input, Output>}
 * @since 2.27
 */
/**
 * Creates a remote query. When called from the browser, the function will be invoked on the server via a `fetch` call.
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#query) for full documentation.
 *
 * @template {StandardSchemaV1} Schema
 * @template Output
 * @overload
 * @param {Schema} schema
 * @param {(arg: StandardSchemaV1.InferOutput<Schema>) => MaybePromise<Output>} fn
 * @returns {RemoteQueryFunction<StandardSchemaV1.InferInput<Schema>, Output>}
 * @since 2.27
 */
/**
 * @template Input
 * @template Output
 * @param {any} validate_or_fn
 * @param {(args?: Input) => MaybePromise<Output>} [maybe_fn]
 * @returns {RemoteQueryFunction<Input, Output>}
 * @since 2.27
 */
/*@__NO_SIDE_EFFECTS__*/
export function query(validate_or_fn, maybe_fn) {
	/** @type {(arg?: Input) => Output} */
	const fn = maybe_fn ?? validate_or_fn;

	/** @type {(arg?: any) => MaybePromise<Input>} */
	const validate = create_validator(validate_or_fn, maybe_fn);

	/** @type {RemoteInfo} */
	const __ = { type: 'query', id: '', name: '' };

	/** @type {RemoteQueryFunction<Input, Output> & { __: RemoteInfo }} */
	const wrapper = (arg) => {
		if (prerendering) {
			throw new Error(
				`Cannot call query '${__.name}' while prerendering, as prerendered pages need static data. Use 'prerender' from $app/server instead`
			);
		}

		const { event, state } = get_request_store();

		const get_remote_function_result = () =>
			run_remote_function(event, state, false, arg, validate, fn);

		/** @type {Promise<any> & Partial<RemoteQuery<any>>} */
		const promise = get_response(__, arg, state, get_remote_function_result);

		promise.catch(() => {});

		promise.set = (value) => update_refresh_value(get_refresh_context(__, 'set', arg), value);

		promise.refresh = () => {
			const refresh_context = get_refresh_context(__, 'refresh', arg);
			const is_immediate_refresh = !refresh_context.cache[refresh_context.cache_key];
			const value = is_immediate_refresh ? promise : get_remote_function_result();
			return update_refresh_value(refresh_context, value, is_immediate_refresh);
		};

		promise.withOverride = () => {
			throw new Error(`Cannot call '${__.name}.withOverride()' on the server`);
		};

		return /** @type {RemoteQuery<Output>} */ (promise);
	};

	Object.defineProperty(wrapper, '__', { value: __ });

	return wrapper;
}

/**
 * Creates a batch query function that collects multiple calls and executes them in a single request
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#query.batch) for full documentation.
 *
 * @template Input
 * @template Output
 * @overload
 * @param {'unchecked'} validate
 * @param {(args: Input[]) => MaybePromise<(arg: Input, idx: number) => Output>} fn
 * @returns {RemoteQueryFunction<Input, Output>}
 * @since 2.35
 */
/**
 * Creates a batch query function that collects multiple calls and executes them in a single request
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#query.batch) for full documentation.
 *
 * @template {StandardSchemaV1} Schema
 * @template Output
 * @overload
 * @param {Schema} schema
 * @param {(args: StandardSchemaV1.InferOutput<Schema>[]) => MaybePromise<(arg: StandardSchemaV1.InferOutput<Schema>, idx: number) => Output>} fn
 * @returns {RemoteQueryFunction<StandardSchemaV1.InferInput<Schema>, Output>}
 * @since 2.35
 */
/**
 * @template Input
 * @template Output
 * @param {any} validate_or_fn
 * @param {(args?: Input[]) => MaybePromise<(arg: Input, idx: number) => Output>} [maybe_fn]
 * @returns {RemoteQueryFunction<Input, Output>}
 * @since 2.35
 */
/*@__NO_SIDE_EFFECTS__*/
function batch(validate_or_fn, maybe_fn) {
	/** @type {(args?: Input[]) => (arg: Input, idx: number) => Output} */
	const fn = maybe_fn ?? validate_or_fn;

	/** @type {(arg?: any) => MaybePromise<Input>} */
	const validate = create_validator(validate_or_fn, maybe_fn);

	/** @type {RemoteInfo & { type: 'query_batch' }} */
	const __ = {
		type: 'query_batch',
		id: '',
		name: '',
		run: (args) => {
			const { event, state } = get_request_store();

			return run_remote_function(
				event,
				state,
				false,
				args,
				(array) => Promise.all(array.map(validate)),
				fn
			);
		}
	};

	/** @type {{ args: any[], resolvers: Array<{resolve: (value: any) => void, reject: (error: any) => void}> }} */
	let batching = { args: [], resolvers: [] };

	/** @type {RemoteQueryFunction<Input, Output> & { __: RemoteInfo }} */
	const wrapper = (arg) => {
		if (prerendering) {
			throw new Error(
				`Cannot call query.batch '${__.name}' while prerendering, as prerendered pages need static data. Use 'prerender' from $app/server instead`
			);
		}

		const { event, state } = get_request_store();

		const get_remote_function_result = () => {
			// Collect all the calls to the same query in the same macrotask,
			// then execute them as one backend request.
			return new Promise((resolve, reject) => {
				// We don't need to deduplicate args here, because get_response already caches/reuses identical calls
				batching.args.push(arg);
				batching.resolvers.push({ resolve, reject });

				if (batching.args.length > 1) return;

				setTimeout(async () => {
					const batched = batching;
					batching = { args: [], resolvers: [] };

					try {
						const get_result = await run_remote_function(
							event,
							state,
							false,
							batched.args,
							(array) => Promise.all(array.map(validate)),
							fn
						);

						for (let i = 0; i < batched.resolvers.length; i++) {
							try {
								batched.resolvers[i].resolve(get_result(batched.args[i], i));
							} catch (error) {
								batched.resolvers[i].reject(error);
							}
						}
					} catch (error) {
						for (const resolver of batched.resolvers) {
							resolver.reject(error);
						}
					}
				}, 0);
			});
		};

		/** @type {Promise<any> & Partial<RemoteQuery<any>>} */
		const promise = get_response(__, arg, state, get_remote_function_result);

		promise.catch(() => {});

		promise.set = (value) => update_refresh_value(get_refresh_context(__, 'set', arg), value);

		promise.refresh = () => {
			const refresh_context = get_refresh_context(__, 'refresh', arg);
			const is_immediate_refresh = !refresh_context.cache[refresh_context.cache_key];
			const value = is_immediate_refresh ? promise : get_remote_function_result();
			return update_refresh_value(refresh_context, value, is_immediate_refresh);
		};

		promise.withOverride = () => {
			throw new Error(`Cannot call '${__.name}.withOverride()' on the server`);
		};

		return /** @type {RemoteQuery<Output>} */ (promise);
	};

	Object.defineProperty(wrapper, '__', { value: __ });

	return wrapper;
}

// Add batch as a property to the query function
Object.defineProperty(query, 'batch', { value: batch, enumerable: true });

/**
 * @param {RemoteInfo} __
 * @param {'set' | 'refresh'} action
 * @param {any} [arg]
 * @returns {{ __: RemoteInfo; state: any; refreshes: Record<string, Promise<any>>; cache: Record<string, Promise<any>>; refreshes_key: string; cache_key: string }}
 */
function get_refresh_context(__, action, arg) {
	const { state } = get_request_store();
	const { refreshes } = state;

	if (!refreshes) {
		const name = __.type === 'query_batch' ? `query.batch '${__.name}'` : `query '${__.name}'`;
		throw new Error(
			`Cannot call ${action} on ${name} because it is not executed in the context of a command/form remote function`
		);
	}

	const cache = get_cache(__, state);
	const cache_key = stringify_remote_arg(arg, state.transport);
	const refreshes_key = create_remote_key(__.id, cache_key);

	return { __, state, refreshes, refreshes_key, cache, cache_key };
}

/**
 * @param {{ __: RemoteInfo; refreshes: Record<string, Promise<any>>; cache: Record<string, Promise<any>>; refreshes_key: string; cache_key: string }} context
 * @param {any} value
 * @param {boolean} [is_immediate_refresh=false]
 * @returns {Promise<void>}
 */
function update_refresh_value(
	{ __, refreshes, refreshes_key, cache, cache_key },
	value,
	is_immediate_refresh = false
) {
	const promise = Promise.resolve(value);

	if (!is_immediate_refresh) {
		cache[cache_key] = promise;
	}

	if (__.id) {
		refreshes[refreshes_key] = promise;
	}

	return promise.then(() => {});
}
