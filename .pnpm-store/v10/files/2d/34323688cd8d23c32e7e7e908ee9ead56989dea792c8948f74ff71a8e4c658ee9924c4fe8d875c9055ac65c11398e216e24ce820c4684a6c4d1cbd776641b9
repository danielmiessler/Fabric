import { resolve } from '../../utils/url.js';
import { decode } from './entities.js';

const DOCTYPE = 'DOCTYPE';
const CDATA_OPEN = '[CDATA[';
const CDATA_CLOSE = ']]>';
const COMMENT_OPEN = '--';
const COMMENT_CLOSE = '-->';

const TAG_OPEN = /[a-zA-Z]/;
const TAG_CHAR = /[a-zA-Z0-9]/;
const ATTRIBUTE_NAME = /[^\t\n\f />"'=]/;

const WHITESPACE = /[\s\n\r]/;

const CRAWLABLE_META_NAME_ATTRS = new Set([
	'og:url',
	'og:image',
	'og:image:url',
	'og:image:secure_url',
	'og:video',
	'og:video:url',
	'og:video:secure_url',
	'og:audio',
	'og:audio:url',
	'og:audio:secure_url',
	'twitter:image'
]);

/**
 * @param {string} html
 * @param {string} base
 */
export function crawl(html, base) {
	/** @type {string[]} */
	const ids = [];

	/** @type {string[]} */
	const hrefs = [];

	let i = 0;
	main: while (i < html.length) {
		const char = html[i];

		if (char === '<') {
			if (html[i + 1] === '!') {
				i += 2;

				if (html.slice(i, i + DOCTYPE.length).toUpperCase() === DOCTYPE) {
					i += DOCTYPE.length;
					while (i < html.length) {
						if (html[i++] === '>') {
							continue main;
						}
					}
				}

				// skip cdata
				if (html.slice(i, i + CDATA_OPEN.length) === CDATA_OPEN) {
					i += CDATA_OPEN.length;
					while (i < html.length) {
						if (html.slice(i, i + CDATA_CLOSE.length) === CDATA_CLOSE) {
							i += CDATA_CLOSE.length;
							continue main;
						}

						i += 1;
					}
				}

				// skip comments
				if (html.slice(i, i + COMMENT_OPEN.length) === COMMENT_OPEN) {
					i += COMMENT_OPEN.length;
					while (i < html.length) {
						if (html.slice(i, i + COMMENT_CLOSE.length) === COMMENT_CLOSE) {
							i += COMMENT_CLOSE.length;
							continue main;
						}

						i += 1;
					}
				}
			}

			// parse opening tags
			const start = ++i;
			if (TAG_OPEN.test(html[start])) {
				while (i < html.length) {
					if (!TAG_CHAR.test(html[i])) {
						break;
					}

					i += 1;
				}

				const tag = html.slice(start, i).toUpperCase();

				/** @type {Record<string, string>} */
				const attributes = {};

				if (tag === 'SCRIPT' || tag === 'STYLE') {
					while (i < html.length) {
						if (
							html[i] === '<' &&
							html[i + 1] === '/' &&
							html.slice(i + 2, i + 2 + tag.length).toUpperCase() === tag
						) {
							continue main;
						}

						i += 1;
					}
				}

				while (i < html.length) {
					const start = i;

					const char = html[start];
					if (char === '>') break;

					if (ATTRIBUTE_NAME.test(char)) {
						i += 1;

						while (i < html.length) {
							if (!ATTRIBUTE_NAME.test(html[i])) {
								break;
							}

							i += 1;
						}

						const name = html.slice(start, i).toLowerCase();

						while (WHITESPACE.test(html[i])) i += 1;

						if (html[i] === '=') {
							i += 1;
							while (WHITESPACE.test(html[i])) i += 1;

							let value;

							if (html[i] === "'" || html[i] === '"') {
								const quote = html[i++];

								const start = i;
								let escaped = false;

								while (i < html.length) {
									if (escaped) {
										escaped = false;
									} else {
										const char = html[i];

										if (html[i] === quote) {
											break;
										}

										if (char === '\\') {
											escaped = true;
										}
									}

									i += 1;
								}

								value = html.slice(start, i);
							} else {
								const start = i;
								while (html[i] !== '>' && !WHITESPACE.test(html[i])) i += 1;
								value = html.slice(start, i);

								i -= 1;
							}

							value = decode(value);
							attributes[name] = value;
						} else {
							i -= 1;
						}
					}

					i += 1;
				}

				const { href, id, name, property, rel, src, srcset, content } = attributes;

				if (href) {
					if (tag === 'BASE') {
						base = resolve(base, href);
					} else if (!rel || !/\bexternal\b/i.test(rel)) {
						hrefs.push(resolve(base, href));
					}
				}

				if (id) {
					ids.push(id);
				}

				if (name && tag === 'A') {
					ids.push(name);
				}

				if (src) {
					hrefs.push(resolve(base, src));
				}

				if (srcset) {
					let value = srcset;
					const candidates = [];
					let insideURL = true;
					value = value.trim();
					for (let i = 0; i < value.length; i++) {
						if (value[i] === ',' && (!insideURL || (insideURL && WHITESPACE.test(value[i + 1])))) {
							candidates.push(value.slice(0, i));
							value = value.substring(i + 1).trim();
							i = 0;
							insideURL = true;
						} else if (WHITESPACE.test(value[i])) {
							insideURL = false;
						}
					}
					candidates.push(value);
					for (const candidate of candidates) {
						const src = candidate.split(WHITESPACE)[0];
						if (src) hrefs.push(resolve(base, src));
					}
				}

				if (tag === 'META' && content) {
					const attr = name ?? property;

					if (attr && CRAWLABLE_META_NAME_ATTRS.has(attr)) {
						hrefs.push(resolve(base, content));
					}
				}
			}
		}

		i += 1;
	}

	return { ids, hrefs };
}
