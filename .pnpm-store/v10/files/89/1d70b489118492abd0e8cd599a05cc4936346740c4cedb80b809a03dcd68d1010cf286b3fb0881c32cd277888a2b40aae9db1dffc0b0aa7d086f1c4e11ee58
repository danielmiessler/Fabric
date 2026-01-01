import { createReadStream } from 'node:fs';
import { Readable } from 'node:stream';
import * as set_cookie_parser from 'set-cookie-parser';
import { SvelteKitError } from '../internal/index.js';

/**
 * @param {import('http').IncomingMessage} req
 * @param {number} [body_size_limit]
 */
function get_raw_body(req, body_size_limit) {
	const h = req.headers;

	if (!h['content-type']) {
		return null;
	}

	const content_length = Number(h['content-length']);

	// check if no request body
	if (
		(req.httpVersionMajor === 1 && isNaN(content_length) && h['transfer-encoding'] == null) ||
		content_length === 0
	) {
		return null;
	}

	if (req.destroyed) {
		const readable = new ReadableStream();
		void readable.cancel();
		return readable;
	}

	let size = 0;
	let cancelled = false;

	return new ReadableStream({
		start(controller) {
			if (body_size_limit !== undefined && content_length > body_size_limit) {
				let message = `Content-length of ${content_length} exceeds limit of ${body_size_limit} bytes.`;

				if (body_size_limit === 0) {
					// https://github.com/sveltejs/kit/pull/11589
					// TODO this exists to aid migration — remove in a future version
					message += ' To disable body size limits, specify Infinity rather than 0.';
				}

				const error = new SvelteKitError(413, 'Payload Too Large', message);

				controller.error(error);
				return;
			}

			req.on('error', (error) => {
				cancelled = true;
				controller.error(error);
			});

			req.on('end', () => {
				if (cancelled) return;
				controller.close();
			});

			req.on('data', (chunk) => {
				if (cancelled) return;

				size += chunk.length;
				if (size > content_length) {
					cancelled = true;

					const constraint = content_length ? 'content-length' : 'BODY_SIZE_LIMIT';
					const message = `request body size exceeded ${constraint} of ${content_length}`;

					const error = new SvelteKitError(413, 'Payload Too Large', message);
					controller.error(error);

					return;
				}

				controller.enqueue(chunk);

				if (controller.desiredSize === null || controller.desiredSize <= 0) {
					req.pause();
				}
			});
		},

		pull() {
			req.resume();
		},

		cancel(reason) {
			cancelled = true;
			req.destroy(reason);
		}
	});
}

/**
 * @param {{
 *   request: import('http').IncomingMessage;
 *   base: string;
 *   bodySizeLimit?: number;
 * }} options
 * @returns {Promise<Request>}
 */
// TODO 3.0 make the signature synchronous?
// eslint-disable-next-line @typescript-eslint/require-await
export async function getRequest({ request, base, bodySizeLimit }) {
	let headers = /** @type {Record<string, string>} */ (request.headers);
	if (request.httpVersionMajor >= 2) {
		// the Request constructor rejects headers with ':' in the name
		headers = Object.assign({}, headers);
		// https://www.rfc-editor.org/rfc/rfc9113.html#section-8.3.1-2.3.5
		if (headers[':authority']) {
			headers.host = headers[':authority'];
		}
		delete headers[':authority'];
		delete headers[':method'];
		delete headers[':path'];
		delete headers[':scheme'];
	}

	// TODO: Whenever Node >=22 is minimum supported version, we can use `request.readableAborted`
	// @see https://github.com/nodejs/node/blob/5cf3c3e24c7257a0c6192ed8ef71efec8ddac22b/lib/internal/streams/readable.js#L1443-L1453
	const controller = new AbortController();
	let errored = false;
	let end_emitted = false;
	request.once('error', () => (errored = true));
	request.once('end', () => (end_emitted = true));
	request.once('close', () => {
		if ((errored || request.destroyed) && !end_emitted) {
			controller.abort();
		}
	});

	return new Request(base + request.url, {
		// @ts-expect-error
		duplex: 'half',
		method: request.method,
		headers: Object.entries(headers),
		signal: controller.signal,
		body:
			request.method === 'GET' || request.method === 'HEAD'
				? undefined
				: get_raw_body(request, bodySizeLimit)
	});
}

/**
 * @param {import('http').ServerResponse} res
 * @param {Response} response
 * @returns {Promise<void>}
 */
// TODO 3.0 make the signature synchronous?
// eslint-disable-next-line @typescript-eslint/require-await
export async function setResponse(res, response) {
	for (const [key, value] of response.headers) {
		try {
			res.setHeader(
				key,
				key === 'set-cookie'
					? set_cookie_parser.splitCookiesString(
							// This is absurd but necessary, TODO: investigate why
							/** @type {string}*/ (response.headers.get(key))
						)
					: value
			);
		} catch (error) {
			res.getHeaderNames().forEach((name) => res.removeHeader(name));
			res.writeHead(500).end(String(error));
			return;
		}
	}

	res.writeHead(response.status);

	if (!response.body) {
		res.end();
		return;
	}

	if (response.body.locked) {
		res.end(
			'Fatal error: Response body is locked. ' +
				"This can happen when the response was already read (for example through 'response.json()' or 'response.text()')."
		);
		return;
	}

	const reader = response.body.getReader();

	if (res.destroyed) {
		void reader.cancel();
		return;
	}

	const cancel = (/** @type {Error|undefined} */ error) => {
		res.off('close', cancel);
		res.off('error', cancel);

		// If the reader has already been interrupted with an error earlier,
		// then it will appear here, it is useless, but it needs to be catch.
		reader.cancel(error).catch(() => {});
		if (error) res.destroy(error);
	};

	res.on('close', cancel);
	res.on('error', cancel);

	void next();
	async function next() {
		try {
			for (;;) {
				const { done, value } = await reader.read();

				if (done) break;

				if (!res.write(value)) {
					res.once('drain', next);
					return;
				}
			}
			res.end();
		} catch (error) {
			cancel(error instanceof Error ? error : new Error(String(error)));
		}
	}
}

/**
 * Converts a file on disk to a readable stream
 * @param {string} file
 * @returns {ReadableStream}
 * @since 2.4.0
 */
export function createReadableStream(file) {
	return /** @type {ReadableStream} */ (Readable.toWeb(createReadStream(file)));
}
