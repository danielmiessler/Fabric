import { public_env } from '../shared-server.js';

/** @type {string} */
let body;

/** @type {string} */
let etag;

/** @type {Headers} */
let headers;

/**
 * @param {Request} request
 * @returns {Response}
 */
export function get_public_env(request) {
	body ??= `export const env=${JSON.stringify(public_env)}`;
	etag ??= `W/${Date.now()}`;
	headers ??= new Headers({
		'content-type': 'application/javascript; charset=utf-8',
		etag
	});

	if (request.headers.get('if-none-match') === etag) {
		return new Response(undefined, { status: 304, headers });
	}

	return new Response(body, { headers });
}
