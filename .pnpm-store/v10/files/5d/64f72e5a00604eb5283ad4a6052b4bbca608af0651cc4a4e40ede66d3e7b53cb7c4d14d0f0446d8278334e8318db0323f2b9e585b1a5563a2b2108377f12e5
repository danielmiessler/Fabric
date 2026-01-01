/** @import { Transport } from '@sveltejs/kit' */
import * as devalue from 'devalue';
import { base64_decode, base64_encode, text_decoder } from './utils.js';

/**
 * @param {string} route_id
 * @param {string} dep
 */
export function validate_depends(route_id, dep) {
	const match = /^(moz-icon|view-source|jar):/.exec(dep);
	if (match) {
		console.warn(
			`${route_id}: Calling \`depends('${dep}')\` will throw an error in Firefox because \`${match[1]}\` is a special URI scheme`
		);
	}
}

export const INVALIDATED_PARAM = 'x-sveltekit-invalidated';

export const TRAILING_SLASH_PARAM = 'x-sveltekit-trailing-slash';

/**
 * @param {any} data
 * @param {string} [location_description]
 */
export function validate_load_response(data, location_description) {
	if (data != null && Object.getPrototypeOf(data) !== Object.prototype) {
		throw new Error(
			`a load function ${location_description} returned ${
				typeof data !== 'object'
					? `a ${typeof data}`
					: data instanceof Response
						? 'a Response object'
						: Array.isArray(data)
							? 'an array'
							: 'a non-plain object'
			}, but must return a plain object at the top level (i.e. \`return {...}\`)`
		);
	}
}

/**
 * Try to `devalue.stringify` the data object using the provided transport encoders.
 * @param {any} data
 * @param {Transport} transport
 */
export function stringify(data, transport) {
	const encoders = Object.fromEntries(Object.entries(transport).map(([k, v]) => [k, v.encode]));

	return devalue.stringify(data, encoders);
}

/**
 * Stringifies the argument (if any) for a remote function in such a way that
 * it is both a valid URL and a valid file name (necessary for prerendering).
 * @param {any} value
 * @param {Transport} transport
 */
export function stringify_remote_arg(value, transport) {
	if (value === undefined) return '';

	// If people hit file/url size limits, we can look into using something like compress_and_encode_text from svelte.dev beyond a certain size
	const json_string = stringify(value, transport);

	const bytes = new TextEncoder().encode(json_string);
	return base64_encode(bytes).replaceAll('=', '').replaceAll('+', '-').replaceAll('/', '_');
}

/**
 * Parses the argument (if any) for a remote function
 * @param {string} string
 * @param {Transport} transport
 */
export function parse_remote_arg(string, transport) {
	if (!string) return undefined;

	const json_string = text_decoder.decode(
		// no need to add back `=` characters, atob can handle it
		base64_decode(string.replaceAll('-', '+').replaceAll('_', '/'))
	);

	const decoders = Object.fromEntries(Object.entries(transport).map(([k, v]) => [k, v.decode]));

	return devalue.parse(json_string, decoders);
}

/**
 * @param {string} id
 * @param {string} payload
 */
export function create_remote_key(id, payload) {
	return id + '/' + payload;
}
