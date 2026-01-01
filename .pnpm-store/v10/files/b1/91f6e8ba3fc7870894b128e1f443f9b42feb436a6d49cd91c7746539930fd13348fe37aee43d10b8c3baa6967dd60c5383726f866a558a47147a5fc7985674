/** @import { StandardSchemaV1 } from '@standard-schema/spec' */

import { HttpError, Redirect, ActionFailure, ValidationError } from './internal/index.js';
import { BROWSER, DEV } from 'esm-env';
import {
	add_data_suffix,
	add_resolution_suffix,
	has_data_suffix,
	has_resolution_suffix,
	strip_data_suffix,
	strip_resolution_suffix
} from '../runtime/pathname.js';
import { text_encoder } from '../runtime/utils.js';

export { VERSION } from '../version.js';

// TODO 3.0: remove these types as they are not used anymore (we can't remove them yet because that would be a breaking change)
/**
 * @template {number} TNumber
 * @template {any[]} [TArray=[]]
 * @typedef {TNumber extends TArray['length'] ? TArray[number] : LessThan<TNumber, [...TArray, TArray['length']]>} LessThan
 */

/**
 * @template {number} TStart
 * @template {number} TEnd
 * @typedef {Exclude<TEnd | LessThan<TEnd>, LessThan<TStart>>} NumericRange
 */

// Keep the status codes as `number` because restricting to certain numbers makes it unnecessarily hard to use compared to the benefits
// (we have runtime errors already to check for invalid codes). Also see https://github.com/sveltejs/kit/issues/11780

// we have to repeat the JSDoc because the display for function overloads is broken
// see https://github.com/microsoft/TypeScript/issues/55056

/**
 * Throws an error with a HTTP status code and an optional message.
 * When called during request handling, this will cause SvelteKit to
 * return an error response without invoking `handleError`.
 * Make sure you're not catching the thrown error, which would prevent SvelteKit from handling it.
 * @param {number} status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
 * @param {App.Error} body An object that conforms to the App.Error type. If a string is passed, it will be used as the message property.
 * @overload
 * @param {number} status
 * @param {App.Error} body
 * @return {never}
 * @throws {HttpError} This error instructs SvelteKit to initiate HTTP error handling.
 * @throws {Error} If the provided status is invalid (not between 400 and 599).
 */
/**
 * Throws an error with a HTTP status code and an optional message.
 * When called during request handling, this will cause SvelteKit to
 * return an error response without invoking `handleError`.
 * Make sure you're not catching the thrown error, which would prevent SvelteKit from handling it.
 * @param {number} status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
 * @param {{ message: string } extends App.Error ? App.Error | string | undefined : never} [body] An object that conforms to the App.Error type. If a string is passed, it will be used as the message property.
 * @overload
 * @param {number} status
 * @param {{ message: string } extends App.Error ? App.Error | string | undefined : never} [body]
 * @return {never}
 * @throws {HttpError} This error instructs SvelteKit to initiate HTTP error handling.
 * @throws {Error} If the provided status is invalid (not between 400 and 599).
 */
/**
 * Throws an error with a HTTP status code and an optional message.
 * When called during request handling, this will cause SvelteKit to
 * return an error response without invoking `handleError`.
 * Make sure you're not catching the thrown error, which would prevent SvelteKit from handling it.
 * @param {number} status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
 * @param {{ message: string } extends App.Error ? App.Error | string | undefined : never} body An object that conforms to the App.Error type. If a string is passed, it will be used as the message property.
 * @return {never}
 * @throws {HttpError} This error instructs SvelteKit to initiate HTTP error handling.
 * @throws {Error} If the provided status is invalid (not between 400 and 599).
 */
export function error(status, body) {
	if ((!BROWSER || DEV) && (isNaN(status) || status < 400 || status > 599)) {
		throw new Error(`HTTP error status codes must be between 400 and 599 — ${status} is invalid`);
	}

	throw new HttpError(status, body);
}

/**
 * Checks whether this is an error thrown by {@link error}.
 * @template {number} T
 * @param {unknown} e
 * @param {T} [status] The status to filter for.
 * @return {e is (HttpError & { status: T extends undefined ? never : T })}
 */
export function isHttpError(e, status) {
	if (!(e instanceof HttpError)) return false;
	return !status || e.status === status;
}

/**
 * Redirect a request. When called during request handling, SvelteKit will return a redirect response.
 * Make sure you're not catching the thrown redirect, which would prevent SvelteKit from handling it.
 *
 * Most common status codes:
 *  * `303 See Other`: redirect as a GET request (often used after a form POST request)
 *  * `307 Temporary Redirect`: redirect will keep the request method
 *  * `308 Permanent Redirect`: redirect will keep the request method, SEO will be transferred to the new page
 *
 * [See all redirect status codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages)
 *
 * @param {300 | 301 | 302 | 303 | 304 | 305 | 306 | 307 | 308 | ({} & number)} status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages). Must be in the range 300-308.
 * @param {string | URL} location The location to redirect to.
 * @throws {Redirect} This error instructs SvelteKit to redirect to the specified location.
 * @throws {Error} If the provided status is invalid.
 * @return {never}
 */
export function redirect(status, location) {
	if ((!BROWSER || DEV) && (isNaN(status) || status < 300 || status > 308)) {
		throw new Error('Invalid status code');
	}

	throw new Redirect(
		// @ts-ignore
		status,
		location.toString()
	);
}

/**
 * Checks whether this is a redirect thrown by {@link redirect}.
 * @param {unknown} e The object to check.
 * @return {e is Redirect}
 */
export function isRedirect(e) {
	return e instanceof Redirect;
}

/**
 * Create a JSON `Response` object from the supplied data.
 * @param {any} data The value that will be serialized as JSON.
 * @param {ResponseInit} [init] Options such as `status` and `headers` that will be added to the response. `Content-Type: application/json` and `Content-Length` headers will be added automatically.
 */
export function json(data, init) {
	// TODO deprecate this in favour of `Response.json` when it's
	// more widely supported
	const body = JSON.stringify(data);

	// we can't just do `text(JSON.stringify(data), init)` because
	// it will set a default `content-type` header. duplicated code
	// means less duplicated work
	const headers = new Headers(init?.headers);
	if (!headers.has('content-length')) {
		headers.set('content-length', text_encoder.encode(body).byteLength.toString());
	}

	if (!headers.has('content-type')) {
		headers.set('content-type', 'application/json');
	}

	return new Response(body, {
		...init,
		headers
	});
}

/**
 * Create a `Response` object from the supplied body.
 * @param {string} body The value that will be used as-is.
 * @param {ResponseInit} [init] Options such as `status` and `headers` that will be added to the response. A `Content-Length` header will be added automatically.
 */
export function text(body, init) {
	const headers = new Headers(init?.headers);
	if (!headers.has('content-length')) {
		const encoded = text_encoder.encode(body);
		headers.set('content-length', encoded.byteLength.toString());
		return new Response(encoded, {
			...init,
			headers
		});
	}

	return new Response(body, {
		...init,
		headers
	});
}

/**
 * Create an `ActionFailure` object. Call when form submission fails.
 * @param {number} status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
 * @overload
 * @param {number} status
 * @returns {import('./public.js').ActionFailure<undefined>}
 */
/**
 * Create an `ActionFailure` object. Call when form submission fails.
 * @template [T=undefined]
 * @param {number} status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
 * @param {T} data Data associated with the failure (e.g. validation errors)
 * @overload
 * @param {number} status
 * @param {T} data
 * @returns {import('./public.js').ActionFailure<T>}
 */
/**
 * Create an `ActionFailure` object. Call when form submission fails.
 * @param {number} status
 * @param {any} [data]
 * @returns {import('./public.js').ActionFailure<any>}
 */
export function fail(status, data) {
	// @ts-expect-error unique symbol missing
	return new ActionFailure(status, data);
}

/**
 * Checks whether this is an action failure thrown by {@link fail}.
 * @param {unknown} e The object to check.
 * @return {e is import('./public.js').ActionFailure}
 */
export function isActionFailure(e) {
	return e instanceof ActionFailure;
}

/**
 * Use this to throw a validation error to imperatively fail form validation.
 * Can be used in combination with `issue` passed to form actions to create field-specific issues.
 *
 * @example
 * ```ts
 * import { invalid } from '@sveltejs/kit';
 * import { form } from '$app/server';
 * import { tryLogin } from '$lib/server/auth';
 * import * as v from 'valibot';
 *
 * export const login = form(
 *   v.object({ name: v.string(), _password: v.string() }),
 *   async ({ name, _password }) => {
 *     const success = tryLogin(name, _password);
 *     if (!success) {
 *       invalid('Incorrect username or password');
 *     }
 *
 *     // ...
 *   }
 * );
 * ```
 * @param  {...(StandardSchemaV1.Issue | string)} issues
 * @returns {never}
 * @since 2.47.3
 */
export function invalid(...issues) {
	throw new ValidationError(
		issues.map((issue) => (typeof issue === 'string' ? { message: issue } : issue))
	);
}

/**
 * Checks whether this is an validation error thrown by {@link invalid}.
 * @param {unknown} e The object to check.
 * @return {e is import('./public.js').ActionFailure}
 * @since 2.47.3
 */
export function isValidationError(e) {
	return e instanceof ValidationError;
}

/**
 * Strips possible SvelteKit-internal suffixes and trailing slashes from the URL pathname.
 * Returns the normalized URL as well as a method for adding the potential suffix back
 * based on a new pathname (possibly including search) or URL.
 * ```js
 * import { normalizeUrl } from '@sveltejs/kit';
 *
 * const { url, denormalize } = normalizeUrl('/blog/post/__data.json');
 * console.log(url.pathname); // /blog/post
 * console.log(denormalize('/blog/post/a')); // /blog/post/a/__data.json
 * ```
 * @param {URL | string} url
 * @returns {{ url: URL, wasNormalized: boolean, denormalize: (url?: string | URL) => URL }}
 * @since 2.18.0
 */
export function normalizeUrl(url) {
	url = new URL(url, 'http://internal');

	const is_route_resolution = has_resolution_suffix(url.pathname);
	const is_data_request = has_data_suffix(url.pathname);
	const has_trailing_slash = url.pathname !== '/' && url.pathname.endsWith('/');

	if (is_route_resolution) {
		url.pathname = strip_resolution_suffix(url.pathname);
	} else if (is_data_request) {
		url.pathname = strip_data_suffix(url.pathname);
	} else if (has_trailing_slash) {
		url.pathname = url.pathname.slice(0, -1);
	}

	return {
		url,
		wasNormalized: is_data_request || is_route_resolution || has_trailing_slash,
		denormalize: (new_url = url) => {
			new_url = new URL(new_url, url);
			if (is_route_resolution) {
				new_url.pathname = add_resolution_suffix(new_url.pathname);
			} else if (is_data_request) {
				new_url.pathname = add_data_suffix(new_url.pathname);
			} else if (has_trailing_slash && !new_url.pathname.endsWith('/')) {
				new_url.pathname += '/';
			}
			return new_url;
		}
	};
}
