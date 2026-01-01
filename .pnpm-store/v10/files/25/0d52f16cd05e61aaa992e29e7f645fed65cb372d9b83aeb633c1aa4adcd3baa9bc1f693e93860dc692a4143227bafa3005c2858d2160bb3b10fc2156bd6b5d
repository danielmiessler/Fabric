const DATA_SUFFIX = '/__data.json';
const HTML_DATA_SUFFIX = '.html__data.json';

/** @param {string} pathname */
export function has_data_suffix(pathname) {
	return pathname.endsWith(DATA_SUFFIX) || pathname.endsWith(HTML_DATA_SUFFIX);
}

/** @param {string} pathname */
export function add_data_suffix(pathname) {
	if (pathname.endsWith('.html')) return pathname.replace(/\.html$/, HTML_DATA_SUFFIX);
	return pathname.replace(/\/$/, '') + DATA_SUFFIX;
}

/** @param {string} pathname */
export function strip_data_suffix(pathname) {
	if (pathname.endsWith(HTML_DATA_SUFFIX)) {
		return pathname.slice(0, -HTML_DATA_SUFFIX.length) + '.html';
	}

	return pathname.slice(0, -DATA_SUFFIX.length);
}

const ROUTE_SUFFIX = '/__route.js';

/**
 * @param {string} pathname
 * @returns {boolean}
 */
export function has_resolution_suffix(pathname) {
	return pathname.endsWith(ROUTE_SUFFIX);
}

/**
 * Convert a regular URL to a route to send to SvelteKit's server-side route resolution endpoint
 * @param {string} pathname
 * @returns {string}
 */
export function add_resolution_suffix(pathname) {
	return pathname.replace(/\/$/, '') + ROUTE_SUFFIX;
}

/**
 * @param {string} pathname
 * @returns {string}
 */
export function strip_resolution_suffix(pathname) {
	return pathname.slice(0, -ROUTE_SUFFIX.length);
}
