/** @import { ActionResult, RemoteForm, RequestEvent, SSRManifest } from '@sveltejs/kit' */
/** @import { RemoteFunctionResponse, RemoteInfo, RequestState, SSROptions } from 'types' */

import { json, error } from '@sveltejs/kit';
import { HttpError, Redirect, SvelteKitError } from '@sveltejs/kit/internal';
import { with_request_store, merge_tracing } from '@sveltejs/kit/internal/server';
import { app_dir, base } from '$app/paths/internal/server';
import { is_form_content_type } from '../../utils/http.js';
import { parse_remote_arg, stringify } from '../shared.js';
import { handle_error_and_jsonify } from './utils.js';
import { normalize_error } from '../../utils/error.js';
import { check_incorrect_fail_use } from './page/actions.js';
import { DEV } from 'esm-env';
import { record_span } from '../telemetry/record_span.js';
import { deserialize_binary_form } from '../form-utils.js';

/** @type {typeof handle_remote_call_internal} */
export async function handle_remote_call(event, state, options, manifest, id) {
	return record_span({
		name: 'sveltekit.remote.call',
		attributes: {
			'sveltekit.remote.call.id': id
		},
		fn: (current) => {
			const traced_event = merge_tracing(event, current);
			return with_request_store({ event: traced_event, state }, () =>
				handle_remote_call_internal(traced_event, state, options, manifest, id)
			);
		}
	});
}

/**
 * @param {RequestEvent} event
 * @param {RequestState} state
 * @param {SSROptions} options
 * @param {SSRManifest} manifest
 * @param {string} id
 */
async function handle_remote_call_internal(event, state, options, manifest, id) {
	const [hash, name, additional_args] = id.split('/');
	const remotes = manifest._.remotes;

	if (!remotes[hash]) error(404);

	const module = await remotes[hash]();
	const fn = module.default[name];

	if (!fn) error(404);

	/** @type {RemoteInfo} */
	const info = fn.__;
	const transport = options.hooks.transport;

	event.tracing.current.setAttributes({
		'sveltekit.remote.call.type': info.type,
		'sveltekit.remote.call.name': info.name
	});

	/** @type {string[] | undefined} */
	let form_client_refreshes;

	try {
		if (info.type === 'query_batch') {
			if (event.request.method !== 'POST') {
				throw new SvelteKitError(
					405,
					'Method Not Allowed',
					`\`query.batch\` functions must be invoked via POST request, not ${event.request.method}`
				);
			}

			/** @type {{ payloads: string[] }} */
			const { payloads } = await event.request.json();

			const args = payloads.map((payload) => parse_remote_arg(payload, transport));
			const get_result = await with_request_store({ event, state }, () => info.run(args));
			const results = await Promise.all(
				args.map(async (arg, i) => {
					try {
						return { type: 'result', data: get_result(arg, i) };
					} catch (error) {
						return {
							type: 'error',
							error: await handle_error_and_jsonify(event, state, options, error),
							status:
								error instanceof HttpError || error instanceof SvelteKitError ? error.status : 500
						};
					}
				})
			);

			return json(
				/** @type {RemoteFunctionResponse} */ ({
					type: 'result',
					result: stringify(results, transport)
				})
			);
		}

		if (info.type === 'form') {
			if (event.request.method !== 'POST') {
				throw new SvelteKitError(
					405,
					'Method Not Allowed',
					`\`form\` functions must be invoked via POST request, not ${event.request.method}`
				);
			}

			if (!is_form_content_type(event.request)) {
				throw new SvelteKitError(
					415,
					'Unsupported Media Type',
					`\`form\` functions expect form-encoded data — received ${event.request.headers.get(
						'content-type'
					)}`
				);
			}

			const { data, meta, form_data } = await deserialize_binary_form(event.request);

			// If this is a keyed form instance (created via form.for(key)), add the key to the form data (unless already set)
			// Note that additional_args will only be set if the form is not enhanced, as enhanced forms transfer the key inside `data`.
			if (additional_args && !('id' in data)) {
				data.id = JSON.parse(decodeURIComponent(additional_args));
			}

			const fn = info.fn;
			const result = await with_request_store({ event, state }, () => fn(data, meta, form_data));

			return json(
				/** @type {RemoteFunctionResponse} */ ({
					type: 'result',
					result: stringify(result, transport),
					refreshes: result.issues ? undefined : await serialize_refreshes(meta.remote_refreshes)
				})
			);
		}

		if (info.type === 'command') {
			/** @type {{ payload: string, refreshes: string[] }} */
			const { payload, refreshes } = await event.request.json();
			const arg = parse_remote_arg(payload, transport);
			const data = await with_request_store({ event, state }, () => fn(arg));

			return json(
				/** @type {RemoteFunctionResponse} */ ({
					type: 'result',
					result: stringify(data, transport),
					refreshes: await serialize_refreshes(refreshes)
				})
			);
		}

		const payload =
			info.type === 'prerender'
				? additional_args
				: /** @type {string} */ (
						// new URL(...) necessary because we're hiding the URL from the user in the event object
						new URL(event.request.url).searchParams.get('payload')
					);

		const data = await with_request_store({ event, state }, () =>
			fn(parse_remote_arg(payload, transport))
		);

		return json(
			/** @type {RemoteFunctionResponse} */ ({
				type: 'result',
				result: stringify(data, transport)
			})
		);
	} catch (error) {
		if (error instanceof Redirect) {
			return json(
				/** @type {RemoteFunctionResponse} */ ({
					type: 'redirect',
					location: error.location,
					refreshes: await serialize_refreshes(form_client_refreshes)
				})
			);
		}

		const status =
			error instanceof HttpError || error instanceof SvelteKitError ? error.status : 500;

		return json(
			/** @type {RemoteFunctionResponse} */ ({
				type: 'error',
				error: await handle_error_and_jsonify(event, state, options, error),
				status
			}),
			{
				// By setting a non-200 during prerendering we fail the prerender process (unless handleHttpError handles it).
				// Errors at runtime will be passed to the client and are handled there
				status: state.prerendering ? status : undefined,
				headers: {
					'cache-control': 'private, no-store'
				}
			}
		);
	}

	/**
	 * @param {string[]=} client_refreshes
	 */
	async function serialize_refreshes(client_refreshes) {
		const refreshes = state.refreshes ?? {};

		if (client_refreshes) {
			for (const key of client_refreshes) {
				if (refreshes[key] !== undefined) continue;

				const [hash, name, payload] = key.split('/');

				const loader = manifest._.remotes[hash];
				const fn = (await loader?.())?.default?.[name];

				if (!fn) error(400, 'Bad Request');

				refreshes[key] = with_request_store({ event, state }, () =>
					fn(parse_remote_arg(payload, transport))
				);
			}
		}

		if (Object.keys(refreshes).length === 0) {
			return undefined;
		}

		return stringify(
			Object.fromEntries(
				await Promise.all(
					Object.entries(refreshes).map(async ([key, promise]) => [key, await promise])
				)
			),
			transport
		);
	}
}

/** @type {typeof handle_remote_form_post_internal} */
export async function handle_remote_form_post(event, state, manifest, id) {
	return record_span({
		name: 'sveltekit.remote.form.post',
		attributes: {
			'sveltekit.remote.form.post.id': id
		},
		fn: (current) => {
			const traced_event = merge_tracing(event, current);
			return with_request_store({ event: traced_event, state }, () =>
				handle_remote_form_post_internal(traced_event, state, manifest, id)
			);
		}
	});
}

/**
 * @param {RequestEvent} event
 * @param {RequestState} state
 * @param {SSRManifest} manifest
 * @param {string} id
 * @returns {Promise<ActionResult>}
 */
async function handle_remote_form_post_internal(event, state, manifest, id) {
	const [hash, name, action_id] = id.split('/');
	const remotes = manifest._.remotes;
	const module = await remotes[hash]?.();

	let form = /** @type {RemoteForm<any, any>} */ (module?.default[name]);

	if (!form) {
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

	if (action_id) {
		// @ts-expect-error
		form = with_request_store({ event, state }, () => form.for(JSON.parse(action_id)));
	}

	try {
		const fn = /** @type {RemoteInfo & { type: 'form' }} */ (/** @type {any} */ (form).__).fn;

		const { data, meta, form_data } = await deserialize_binary_form(event.request);
		if (action_id && !('id' in data)) {
			data.id = JSON.parse(decodeURIComponent(action_id));
		}

		await with_request_store({ event, state }, () => fn(data, meta, form_data));

		// We don't want the data to appear on `let { form } = $props()`, which is why we're not returning it.
		// It is instead available on `myForm.result`, setting of which happens within the remote `form` function.
		return {
			type: 'success',
			status: 200
		};
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
 * @param {URL} url
 */
export function get_remote_id(url) {
	return (
		url.pathname.startsWith(`${base}/${app_dir}/remote/`) &&
		url.pathname.replace(`${base}/${app_dir}/remote/`, '')
	);
}

/**
 * @param {URL} url
 */
export function get_remote_action(url) {
	return url.searchParams.get('/remote');
}
