/** @import { RemoteCommand, RemoteQueryOverride } from '@sveltejs/kit' */
/** @import { RemoteFunctionResponse } from 'types' */
/** @import { Query } from './query.svelte.js' */
import { app_dir, base } from '$app/paths/internal/client';
import * as devalue from 'devalue';
import { HttpError } from '@sveltejs/kit/internal';
import { app } from '../client.js';
import { stringify_remote_arg } from '../../shared.js';
import { refresh_queries, release_overrides } from './shared.svelte.js';

/**
 * Client-version of the `command` function from `$app/server`.
 * @param {string} id
 * @returns {RemoteCommand<any, any>}
 */
export function command(id) {
	/** @type {number} */
	let pending_count = $state(0);

	// Careful: This function MUST be synchronous (can't use the async keyword) because the return type has to be a promise with an updates() method.
	// If we make it async, the return type will be a promise that resolves to a promise with an updates() method, which is not what we want.
	/** @type {RemoteCommand<any, any>} */
	const command_function = (arg) => {
		/** @type {Array<Query<any> | RemoteQueryOverride>} */
		let updates = [];

		// Increment pending count when command starts
		pending_count++;

		/** @type {Promise<any> & { updates: (...args: any[]) => any }} */
		const promise = (async () => {
			try {
				// Wait a tick to give room for the `updates` method to be called
				await Promise.resolve();

				const response = await fetch(`${base}/${app_dir}/remote/${id}`, {
					method: 'POST',
					body: JSON.stringify({
						payload: stringify_remote_arg(arg, app.hooks.transport),
						refreshes: updates.map((u) => u._key)
					}),
					headers: {
						'Content-Type': 'application/json',
						'x-sveltekit-pathname': location.pathname,
						'x-sveltekit-search': location.search
					}
				});

				if (!response.ok) {
					release_overrides(updates);
					// We only end up here in case of a network error or if the server has an internal error
					// (which shouldn't happen because we handle errors on the server and always send a 200 response)
					throw new Error('Failed to execute remote function');
				}

				const result = /** @type {RemoteFunctionResponse} */ (await response.json());
				if (result.type === 'redirect') {
					release_overrides(updates);
					throw new Error(
						'Redirects are not allowed in commands. Return a result instead and use goto on the client'
					);
				} else if (result.type === 'error') {
					release_overrides(updates);
					throw new HttpError(result.status ?? 500, result.error);
				} else {
					if (result.refreshes) {
						refresh_queries(result.refreshes, updates);
					}

					return devalue.parse(result.result, app.decoders);
				}
			} finally {
				// Decrement pending count when command completes
				pending_count--;
			}
		})();

		promise.updates = (/** @type {any} */ ...args) => {
			updates = args;
			// @ts-expect-error Don't allow updates to be called multiple times
			delete promise.updates;
			return promise;
		};

		return promise;
	};

	Object.defineProperty(command_function, 'pending', {
		get: () => pending_count
	});

	return command_function;
}
