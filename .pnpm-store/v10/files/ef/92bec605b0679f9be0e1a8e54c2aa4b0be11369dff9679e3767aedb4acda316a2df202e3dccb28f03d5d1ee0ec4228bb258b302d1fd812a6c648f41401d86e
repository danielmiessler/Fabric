/**
 * When inside a double-quoted attribute value, only `&` and `"` hold special meaning.
 * @see https://html.spec.whatwg.org/multipage/parsing.html#attribute-value-(double-quoted)-state
 * @type {Record<string, string>}
 */
const escape_html_attr_dict = {
	'&': '&amp;',
	'"': '&quot;'
	// Svelte also escapes < because the escape function could be called inside a `noscript` there
	// https://github.com/sveltejs/svelte/security/advisories/GHSA-8266-84wp-wv5c
	// However, that doesn't apply in SvelteKit
};

/**
 * @type {Record<string, string>}
 */
const escape_html_dict = {
	'&': '&amp;',
	'<': '&lt;'
};

const surrogates = // high surrogate without paired low surrogate
	'[\\ud800-\\udbff](?![\\udc00-\\udfff])|' +
	// a valid surrogate pair, the only match with 2 code units
	// we match it so that we can match unpaired low surrogates in the same pass
	// TODO: use lookbehind assertions once they are widely supported: (?<![\ud800-udbff])[\udc00-\udfff]
	'[\\ud800-\\udbff][\\udc00-\\udfff]|' +
	// unpaired low surrogate (see previous match)
	'[\\udc00-\\udfff]';

const escape_html_attr_regex = new RegExp(
	`[${Object.keys(escape_html_attr_dict).join('')}]|` + surrogates,
	'g'
);

const escape_html_regex = new RegExp(
	`[${Object.keys(escape_html_dict).join('')}]|` + surrogates,
	'g'
);

/**
 * Escapes unpaired surrogates (which are allowed in js strings but invalid in HTML) and
 * escapes characters that are special.
 *
 * @param {string} str
 * @param {boolean} [is_attr]
 * @returns {string} escaped string
 * @example const html = `<tag data-value="${escape_html('value', true)}">...</tag>`;
 */
export function escape_html(str, is_attr) {
	const dict = is_attr ? escape_html_attr_dict : escape_html_dict;
	const escaped_str = str.replace(is_attr ? escape_html_attr_regex : escape_html_regex, (match) => {
		if (match.length === 2) {
			// valid surrogate pair
			return match;
		}

		return dict[match] ?? `&#${match.charCodeAt(0)};`;
	});

	return escaped_str;
}
