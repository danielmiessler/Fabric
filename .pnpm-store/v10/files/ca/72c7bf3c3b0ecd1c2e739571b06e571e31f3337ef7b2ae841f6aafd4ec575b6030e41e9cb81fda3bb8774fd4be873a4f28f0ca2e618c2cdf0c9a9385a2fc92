import { respond } from './respond.js';
import { set_private_env, set_public_env, set_safe_public_env } from '../shared-server.js';
import { options, get_hooks } from '__SERVER__/internal.js';
import { DEV } from 'esm-env';
import { filter_private_env, filter_public_env } from '../../utils/env.js';
import { prerendering } from '__sveltekit/environment';
import { set_read_implementation, set_manifest } from '__sveltekit/server';

/** @type {ProxyHandler<{ type: 'public' | 'private' }>} */
const prerender_env_handler = {
	get({ type }, prop) {
		throw new Error(
			`Cannot read values from $env/dynamic/${type} while prerendering (attempted to read env.${prop.toString()}). Use $env/static/${type} instead`
		);
	}
};

/** @type {Promise<any>} */
let init_promise;

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

		set_manifest(manifest);
	}

	/**
	 * @param {{
	 *   env: Record<string, string>;
	 *   read?: (file: string) => ReadableStream;
	 * }} opts
	 */
	async init({ env, read }) {
		// Take care: Some adapters may have to call `Server.init` per-request to set env vars,
		// so anything that shouldn't be rerun should be wrapped in an `if` block to make sure it hasn't
		// been done already.

		// set env, in case it's used in initialisation
		const prefixes = {
			public_prefix: this.#options.env_public_prefix,
			private_prefix: this.#options.env_private_prefix
		};

		const private_env = filter_private_env(env, prefixes);
		const public_env = filter_public_env(env, prefixes);

		set_private_env(
			prerendering ? new Proxy({ type: 'private' }, prerender_env_handler) : private_env
		);
		set_public_env(
			prerendering ? new Proxy({ type: 'public' }, prerender_env_handler) : public_env
		);
		set_safe_public_env(public_env);

		if (read) {
			set_read_implementation(read);
		}

		// During DEV and for some adapters this function might be called in quick succession,
		// so we need to make sure we're not invoking this logic (most notably the init hook) multiple times
		await (init_promise ??= (async () => {
			try {
				const module = await get_hooks();

				this.#options.hooks = {
					handle: module.handle || (({ event, resolve }) => resolve(event)),
					handleError: module.handleError || (({ error }) => console.error(error)),
					handleFetch: module.handleFetch || (({ request, fetch }) => fetch(request)),
					reroute: module.reroute || (() => {}),
					transport: module.transport || {}
				};

				if (module.init) {
					await module.init();
				}
			} catch (error) {
				if (DEV) {
					this.#options.hooks = {
						handle: () => {
							throw error;
						},
						handleError: ({ error }) => console.error(error),
						handleFetch: ({ request, fetch }) => fetch(request),
						reroute: () => {},
						transport: {}
					};
				} else {
					throw error;
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
