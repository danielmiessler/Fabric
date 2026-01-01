/** @import { PromiseWithResolvers } from '../../utils/promise.js' */
import { with_resolvers } from '../../utils/promise.js';
import { IN_WEBCONTAINER } from './constants.js';
import { respond } from './respond.js';
import { set_private_env, set_public_env } from '../shared-server.js';
import { options, get_hooks } from '__SERVER__/internal.js';
import { DEV } from 'esm-env';
import { filter_env } from '../../utils/env.js';
import { format_server_error } from './utils.js';
import { set_read_implementation, set_manifest } from '__sveltekit/server';
import { set_app } from './app.js';

/** @type {Promise<any>} */
let init_promise;

/** @type {Promise<void> | null} */
let current = null;

export class Server {
	/** @type {import('types').SSROptions} */
	#options;

	/** @type {import('@sveltejs/kit').SSRManifest} */
	#manifest;

	/** @param {import('@sveltejs/kit').SSRManifest} manifest */
	constructor(manifest) {
		/** @type {import('types').SSROptions} */
		this.#options = options;
		this.#manifest = manifest;

		// Since AsyncLocalStorage is not working in webcontainers, we don't reset `sync_store`
		// in `src/exports/internal/event.js` and handle only one request at a time.
		if (IN_WEBCONTAINER) {
			const respond = this.respond.bind(this);

			/** @type {typeof respond} */
			this.respond = async (...args) => {
				const { promise, resolve } = /** @type {PromiseWithResolvers<void>} */ (with_resolvers());

				const previous = current;
				current = promise;

				await previous;
				return respond(...args).finally(resolve);
			};
		}

		set_manifest(manifest);
	}

	/**
	 * @param {import('@sveltejs/kit').ServerInitOptions} opts
	 */
	async init({ env, read }) {
		// Take care: Some adapters may have to call `Server.init` per-request to set env vars,
		// so anything that shouldn't be rerun should be wrapped in an `if` block to make sure it hasn't
		// been done already.

		// set env, in case it's used in initialisation
		const { env_public_prefix, env_private_prefix } = this.#options;

		set_private_env(filter_env(env, env_private_prefix, env_public_prefix));
		set_public_env(filter_env(env, env_public_prefix, env_private_prefix));

		if (read) {
			// Wrap the read function to handle MaybePromise<ReadableStream>
			// and ensure the public API stays synchronous
			/** @param {string} file */
			const wrapped_read = (file) => {
				const result = read(file);
				if (result instanceof ReadableStream) {
					return result;
				} else {
					return new ReadableStream({
						async start(controller) {
							try {
								const stream = await Promise.resolve(result);
								if (!stream) {
									controller.close();
									return;
								}

								const reader = stream.getReader();

								while (true) {
									const { done, value } = await reader.read();
									if (done) break;
									controller.enqueue(value);
								}

								controller.close();
							} catch (error) {
								controller.error(error);
							}
						}
					});
				}
			};

			set_read_implementation(wrapped_read);
		}

		// During DEV and for some adapters this function might be called in quick succession,
		// so we need to make sure we're not invoking this logic (most notably the init hook) multiple times
		await (init_promise ??= (async () => {
			try {
				const module = await get_hooks();

				this.#options.hooks = {
					handle: module.handle || (({ event, resolve }) => resolve(event)),
					handleError:
						module.handleError ||
						(({ status, error, event }) => {
							const error_message = format_server_error(
								status,
								/** @type {Error} */ (error),
								event
							);
							console.error(error_message);
						}),
					handleFetch: module.handleFetch || (({ request, fetch }) => fetch(request)),
					handleValidationError:
						module.handleValidationError ||
						(({ issues }) => {
							console.error('Remote function schema validation failed:', issues);
							return { message: 'Bad Request' };
						}),
					reroute: module.reroute || (() => {}),
					transport: module.transport || {}
				};

				set_app({
					decoders: module.transport
						? Object.fromEntries(Object.entries(module.transport).map(([k, v]) => [k, v.decode]))
						: {}
				});

				if (module.init) {
					await module.init();
				}
			} catch (e) {
				if (DEV) {
					this.#options.hooks = {
						handle: () => {
							throw e;
						},
						handleError: ({ error }) => console.error(error),
						handleFetch: ({ request, fetch }) => fetch(request),
						handleValidationError: () => {
							return { message: 'Bad Request' };
						},
						reroute: () => {},
						transport: {}
					};

					set_app({
						decoders: {}
					});
				} else {
					throw e;
				}
			}
		})());
	}

	/**
	 * @param {Request} request
	 * @param {import('types').RequestOptions} options
	 */
	async respond(request, options) {
		return respond(request, this.#options, this.#manifest, {
			...options,
			error: false,
			depth: 0
		});
	}
}
