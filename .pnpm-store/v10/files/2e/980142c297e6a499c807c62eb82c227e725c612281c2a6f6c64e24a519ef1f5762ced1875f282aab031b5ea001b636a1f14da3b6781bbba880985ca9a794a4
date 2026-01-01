import { DEV } from 'esm-env';
import { json, text } from '@sveltejs/kit';
import { HttpError } from '@sveltejs/kit/internal';
import { with_request_store } from '@sveltejs/kit/internal/server';
import { coalesce_to_error, get_message, get_status } from '../../utils/error.js';
import { negotiate } from '../../utils/http.js';
import { fix_stack_trace } from '../shared-server.js';
import { ENDPOINT_METHODS } from '../../constants.js';
import { escape_html } from '../../utils/escape.js';

/** @param {any} body */
export function is_pojo(body) {
	if (typeof body !== 'object') return false;

	if (body) {
		if (body instanceof Uint8Array) return false;
		if (body instanceof ReadableStream) return false;
	}

	return true;
}

/**
 * @param {Partial<Record<import('types').HttpMethod, any>>} mod
 * @param {import('types').HttpMethod} method
 */
export function method_not_allowed(mod, method) {
	return text(`${method} method not allowed`, {
		status: 405,
		headers: {
			// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
			// "The server must generate an Allow header field in a 405 status code response"
			allow: allowed_methods(mod).join(', ')
		}
	});
}

/** @param {Partial<Record<import('types').HttpMethod, any>>} mod */
export function allowed_methods(mod) {
	const allowed = ENDPOINT_METHODS.filter((method) => method in mod);

	// if there's no HEAD handler, but we have a GET handler, we respond to
	// HEAD requests using the GET handler and omit the response body.
	if ('GET' in mod && !('HEAD' in mod)) {
		allowed.push('HEAD');
	}

	return allowed;
}

/**
 * @param {import('types').SSROptions} options
 */
export function get_global_name(options) {
	return DEV ? '__sveltekit_dev' : `__sveltekit_${options.version_hash}`;
}

/**
 * Return as a response that renders the error.html
 *
 * @param {import('types').SSROptions} options
 * @param {number} status
 * @param {string} message
 */
export function static_error_page(options, status, message) {
	let page = options.templates.error({ status, message: escape_html(message) });

	if (DEV) {
		// inject Vite HMR client, for easier debugging
		page = page.replace('</head>', '<script type="module" src="/@vite/client"></script></head>');
	}

	return text(page, {
		headers: { 'content-type': 'text/html; charset=utf-8' },
		status
	});
}

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').RequestState} state
 * @param {import('types').SSROptions} options
 * @param {unknown} error
 */
export async function handle_fatal_error(event, state, options, error) {
	error = error instanceof HttpError ? error : coalesce_to_error(error);
	const status = get_status(error);
	const body = await handle_error_and_jsonify(event, state, options, error);

	// ideally we'd use sec-fetch-dest instead, but Safari — quelle surprise — doesn't support it
	const type = negotiate(event.request.headers.get('accept') || 'text/html', [
		'application/json',
		'text/html'
	]);

	if (event.isDataRequest || type === 'application/json') {
		return json(body, {
			status
		});
	}

	return static_error_page(options, status, body.message);
}

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {import('types').RequestState} state
 * @param {import('types').SSROptions} options
 * @param {any} error
 * @returns {Promise<App.Error>}
 */
export async function handle_error_and_jsonify(event, state, options, error) {
	if (error instanceof HttpError) {
		// @ts-expect-error custom user errors may not have a message field if App.Error is overwritten
		return { message: 'Unknown Error', ...error.body };
	}

	if (DEV && typeof error == 'object') {
		fix_stack_trace(error);
	}

	const status = get_status(error);
	const message = get_message(error);

	return (
		(await with_request_store({ event, state }, () =>
			options.hooks.handleError({ error, event, status, message })
		)) ?? { message }
	);
}

/**
 * @param {number} status
 * @param {string} location
 */
export function redirect_response(status, location) {
	const response = new Response(undefined, {
		status,
		headers: { location }
	});
	return response;
}

/**
 * @param {import('@sveltejs/kit').RequestEvent} event
 * @param {Error & { path: string }} error
 */
export function clarify_devalue_error(event, error) {
	if (error.path) {
		return (
			`Data returned from \`load\` while rendering ${event.route.id} is not serializable: ${error.message} (${error.path}). ` +
			`If you need to serialize/deserialize custom types, use transport hooks: https://svelte.dev/docs/kit/hooks#Universal-hooks-transport.`
		);
	}

	if (error.path === '') {
		return `Data returned from \`load\` while rendering ${event.route.id} is not a plain object`;
	}

	// belt and braces — this should never happen
	return error.message;
}

/**
 * @param {import('types').ServerDataNode} node
 */
export function serialize_uses(node) {
	const uses = {};

	if (node.uses && node.uses.dependencies.size > 0) {
		uses.dependencies = Array.from(node.uses.dependencies);
	}

	if (node.uses && node.uses.search_params.size > 0) {
		uses.search_params = Array.from(node.uses.search_params);
	}

	if (node.uses && node.uses.params.size > 0) {
		uses.params = Array.from(node.uses.params);
	}

	if (node.uses?.parent) uses.parent = 1;
	if (node.uses?.route) uses.route = 1;
	if (node.uses?.url) uses.url = 1;

	return uses;
}

/**
 * Returns `true` if the given path was prerendered
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @param {string} pathname Should include the base and be decoded
 */
export function has_prerendered_path(manifest, pathname) {
	return (
		manifest._.prerendered_routes.has(pathname) ||
		(pathname.at(-1) === '/' && manifest._.prerendered_routes.has(pathname.slice(0, -1)))
	);
}

/**
 * Formats the error into a nice message with sanitized stack trace
 * @param {number} status
 * @param {Error} error
 * @param {import('@sveltejs/kit').RequestEvent} event
 */
export function format_server_error(status, error, event) {
	const formatted_text = `\n\x1b[1;31m[${status}] ${event.request.method} ${event.url.pathname}\x1b[0m`;

	if (status === 404) {
		return formatted_text;
	}

	return `${formatted_text}\n${DEV ? clean_up_stack_trace(error) : error.stack}`;
}

/**
 * In dev, tidy up stack traces by making paths relative to the current project directory
 * @param {string} file
 */
let relative = (file) => file;

if (DEV) {
	try {
		const path = await import('node:path');
		const process = await import('node:process');

		relative = (file) => path.relative(process.cwd(), file);
	} catch {
		// do nothing
	}
}

/**
 * Provides a refined stack trace by excluding lines following the last occurrence of a line containing +page. +layout. or +server.
 * @param {Error} error
 */
export function clean_up_stack_trace(error) {
	const stack_trace = (error.stack?.split('\n') ?? []).map((line) => {
		return line.replace(/\((.+)(:\d+:\d+)\)$/, (_, file, loc) => `(${relative(file)}${loc})`);
	});

	// progressive enhancement for people who haven't configured kit.files.src to something else
	const last_line_from_src_code = stack_trace.findLastIndex((line) => /\(src[\\/]/.test(line));

	if (last_line_from_src_code === -1) {
		// default to the whole stack trace
		return error.stack;
	}

	return stack_trace.slice(0, last_line_from_src_code + 1).join('\n');
}

/**
 * Returns the filename without the extension. e.g., `+page.server`, `+page`, etc.
 * @param {string | undefined} node_id
 * @returns {string}
 */
export function get_node_type(node_id) {
	const parts = node_id?.split('/');
	const filename = parts?.at(-1);
	if (!filename) return 'unknown';
	const dot_parts = filename.split('.');
	return dot_parts.slice(0, -1).join('.');
}
