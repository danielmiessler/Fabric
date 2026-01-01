/** @import { RemoteCommand } from '@sveltejs/kit' */
/** @import { RemoteInfo, MaybePromise } from 'types' */
/** @import { StandardSchemaV1 } from '@standard-schema/spec' */
import { get_request_store } from '@sveltejs/kit/internal/server';
import { create_validator, run_remote_function } from './shared.js';

/**
 * Creates a remote command. When called from the browser, the function will be invoked on the server via a `fetch` call.
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#command) for full documentation.
 *
 * @template Output
 * @overload
 * @param {() => Output} fn
 * @returns {RemoteCommand<void, Output>}
 * @since 2.27
 */
/**
 * Creates a remote command. When called from the browser, the function will be invoked on the server via a `fetch` call.
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#command) for full documentation.
 *
 * @template Input
 * @template Output
 * @overload
 * @param {'unchecked'} validate
 * @param {(arg: Input) => Output} fn
 * @returns {RemoteCommand<Input, Output>}
 * @since 2.27
 */
/**
 * Creates a remote command. When called from the browser, the function will be invoked on the server via a `fetch` call.
 *
 * See [Remote functions](https://svelte.dev/docs/kit/remote-functions#command) for full documentation.
 *
 * @template {StandardSchemaV1} Schema
 * @template Output
 * @overload
 * @param {Schema} validate
 * @param {(arg: StandardSchemaV1.InferOutput<Schema>) => Output} fn
 * @returns {RemoteCommand<StandardSchemaV1.InferInput<Schema>, Output>}
 * @since 2.27
 */
/**
 * @template Input
 * @template Output
 * @param {any} validate_or_fn
 * @param {(arg?: Input) => Output} [maybe_fn]
 * @returns {RemoteCommand<Input, Output>}
 * @since 2.27
 */
/*@__NO_SIDE_EFFECTS__*/
export function command(validate_or_fn, maybe_fn) {
	/** @type {(arg?: Input) => Output} */
	const fn = maybe_fn ?? validate_or_fn;

	/** @type {(arg?: any) => MaybePromise<Input>} */
	const validate = create_validator(validate_or_fn, maybe_fn);

	/** @type {RemoteInfo} */
	const __ = { type: 'command', id: '', name: '' };

	/** @type {RemoteCommand<Input, Output> & { __: RemoteInfo }} */
	const wrapper = (arg) => {
		const { event, state } = get_request_store();

		if (state.is_endpoint_request) {
			if (!['POST', 'PUT', 'PATCH', 'DELETE'].includes(event.request.method)) {
				throw new Error(
					`Cannot call a command (\`${__.name}(${maybe_fn ? '...' : ''})\`) from a ${event.request.method} handler`
				);
			}
		} else if (!event.isRemoteRequest) {
			throw new Error(
				`Cannot call a command (\`${__.name}(${maybe_fn ? '...' : ''})\`) during server-side rendering`
			);
		}

		state.refreshes ??= {};

		const promise = Promise.resolve(run_remote_function(event, state, true, arg, validate, fn));

		// @ts-expect-error
		promise.updates = () => {
			throw new Error(`Cannot call '${__.name}(...).updates(...)' on the server`);
		};

		return /** @type {ReturnType<RemoteCommand<Input, Output>>} */ (promise);
	};

	Object.defineProperty(wrapper, '__', { value: __ });

	// On the server, pending is always 0
	Object.defineProperty(wrapper, 'pending', {
		get: () => 0
	});

	return wrapper;
}
