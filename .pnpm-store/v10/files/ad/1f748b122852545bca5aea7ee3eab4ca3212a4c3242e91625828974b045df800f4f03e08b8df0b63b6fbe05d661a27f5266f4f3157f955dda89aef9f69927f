import { escape_html } from '../../../utils/escape.js';
import { hash } from '../../hash.js';

/**
 * Inside a script element, only `</script` and `<!--` hold special meaning to the HTML parser.
 *
 * The first closes the script element, so everything after is treated as raw HTML.
 * The second disables further parsing until `-->`, so the script element might be unexpectedly
 * kept open until until an unrelated HTML comment in the page.
 *
 * U+2028 LINE SEPARATOR and U+2029 PARAGRAPH SEPARATOR are escaped for the sake of pre-2018
 * browsers.
 *
 * @see tests for unsafe parsing examples.
 * @see https://html.spec.whatwg.org/multipage/scripting.html#restrictions-for-contents-of-script-elements
 * @see https://html.spec.whatwg.org/multipage/syntax.html#cdata-rcdata-restrictions
 * @see https://html.spec.whatwg.org/multipage/parsing.html#script-data-state
 * @see https://html.spec.whatwg.org/multipage/parsing.html#script-data-double-escaped-state
 * @see https://github.com/tc39/proposal-json-superset
 * @type {Record<string, string>}
 */
const replacements = {
	'<': '\\u003C',
	'\u2028': '\\u2028',
	'\u2029': '\\u2029'
};

const pattern = new RegExp(`[${Object.keys(replacements).join('')}]`, 'g');

/**
 * Generates a raw HTML string containing a safe script element carrying data and associated attributes.
 *
 * It escapes all the special characters needed to guarantee the element is unbroken, but care must
 * be taken to ensure it is inserted in the document at an acceptable position for a script element,
 * and that the resulting string isn't further modified.
 *
 * @param {import('./types.js').Fetched} fetched
 * @param {(name: string, value: string) => boolean} filter
 * @param {boolean} [prerendering]
 * @returns {string} The raw HTML of a script element carrying the JSON payload.
 * @example const html = serialize_data('/data.json', null, { foo: 'bar' });
 */
export function serialize_data(fetched, filter, prerendering = false) {
	/** @type {Record<string, string>} */
	const headers = {};

	let cache_control = null;
	let age = null;
	let varyAny = false;

	for (const [key, value] of fetched.response.headers) {
		if (filter(key, value)) {
			headers[key] = value;
		}

		if (key === 'cache-control') cache_control = value;
		else if (key === 'age') age = value;
		else if (key === 'vary' && value.trim() === '*') varyAny = true;
	}

	const payload = {
		status: fetched.response.status,
		statusText: fetched.response.statusText,
		headers,
		body: fetched.response_body
	};

	const safe_payload = JSON.stringify(payload).replace(pattern, (match) => replacements[match]);

	const attrs = [
		'type="application/json"',
		'data-sveltekit-fetched',
		`data-url="${escape_html(fetched.url, true)}"`
	];

	if (fetched.is_b64) {
		attrs.push('data-b64');
	}

	if (fetched.request_headers || fetched.request_body) {
		/** @type {import('types').StrictBody[]} */
		const values = [];

		if (fetched.request_headers) {
			values.push([...new Headers(fetched.request_headers)].join(','));
		}

		if (fetched.request_body) {
			values.push(fetched.request_body);
		}

		attrs.push(`data-hash="${hash(...values)}"`);
	}

	// Compute the time the response should be cached, taking into account max-age and age.
	// Do not cache at all if a `Vary: *` header is present, as this indicates that the
	// cache is likely to get busted.
	if (!prerendering && fetched.method === 'GET' && cache_control && !varyAny) {
		const match = /s-maxage=(\d+)/g.exec(cache_control) ?? /max-age=(\d+)/g.exec(cache_control);
		if (match) {
			const ttl = +match[1] - +(age ?? '0');
			attrs.push(`data-ttl="${ttl}"`);
		}
	}

	return `<script ${attrs.join(' ')}>${safe_payload}</script>`;
}
