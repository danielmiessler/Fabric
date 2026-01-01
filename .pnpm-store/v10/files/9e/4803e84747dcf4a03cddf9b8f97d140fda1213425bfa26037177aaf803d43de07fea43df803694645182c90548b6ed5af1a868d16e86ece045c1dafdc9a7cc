import { render_response } from './render.js';
import { load_data, load_server_data } from './load_data.js';
import { handle_error_and_jsonify, static_error_page, redirect_response } from '../utils.js';
import { Redirect } from '../../control.js';
import { get_status } from '../../../utils/error.js';
import { PageNodes } from '../../../utils/page_nodes.js';

/**
 * @typedef {import('./types.js').Loaded} Loaded
 */

/**
 * @param {{
 *   event: import('@sveltejs/kit').RequestEvent;
 *   options: import('types').SSROptions;
 *   manifest: import('@sveltejs/kit').SSRManifest;
 *   state: import('types').SSRState;
 *   status: number;
 *   error: unknown;
 *   resolve_opts: import('types').RequiredResolveOptions;
 * }} opts
 */
export async function respond_with_error({
	event,
	options,
	manifest,
	state,
	status,
	error,
	resolve_opts
}) {
	// reroute to the fallback page to prevent an infinite chain of requests.
	if (event.request.headers.get('x-sveltekit-error')) {
		return static_error_page(options, status, /** @type {Error} */ (error).message);
	}

	/** @type {import('./types.js').Fetched[]} */
	const fetched = [];

	try {
		const branch = [];
		const default_layout = await manifest._.nodes[0](); // 0 is always the root layout
		const nodes = new PageNodes([default_layout]);
		const ssr = nodes.ssr();
		const csr = nodes.csr();

		if (ssr) {
			state.error = true;

			const server_data_promise = load_server_data({
				event,
				state,
				node: default_layout,
				// eslint-disable-next-line @typescript-eslint/require-await
				parent: async () => ({})
			});

			const server_data = await server_data_promise;

			const data = await load_data({
				event,
				fetched,
				node: default_layout,
				// eslint-disable-next-line @typescript-eslint/require-await
				parent: async () => ({}),
				resolve_opts,
				server_data_promise,
				state,
				csr
			});

			branch.push(
				{
					node: default_layout,
					server_data,
					data
				},
				{
					node: await manifest._.nodes[1](), // 1 is always the root error
					data: null,
					server_data: null
				}
			);
		}

		return await render_response({
			options,
			manifest,
			state,
			page_config: {
				ssr,
				csr
			},
			status,
			error: await handle_error_and_jsonify(event, options, error),
			branch,
			fetched,
			event,
			resolve_opts
		});
	} catch (e) {
		// Edge case: If route is a 404 and the user redirects to somewhere from the root layout,
		// we end up here.
		if (e instanceof Redirect) {
			return redirect_response(e.status, e.location);
		}

		return static_error_page(
			options,
			get_status(e),
			(await handle_error_and_jsonify(event, options, e)).message
		);
	}
}
