/**
 * @param {Set<string>} expected
 */
function validator(expected) {
	/**
	 * @param {any} module
	 * @param {string} [file]
	 */
	function validate(module, file) {
		if (!module) return;

		for (const key in module) {
			if (key[0] === '_' || expected.has(key)) continue; // key is valid in this module

			const values = [...expected.values()];

			const hint =
				hint_for_supported_files(key, file?.slice(file.lastIndexOf('.'))) ??
				`valid exports are ${values.join(', ')}, or anything with a '_' prefix`;

			throw new Error(`Invalid export '${key}'${file ? ` in ${file}` : ''} (${hint})`);
		}
	}

	return validate;
}

/**
 * @param {string} key
 * @param {string} ext
 * @returns {string | void}
 */
function hint_for_supported_files(key, ext = '.js') {
	const supported_files = [];

	if (valid_layout_exports.has(key)) {
		supported_files.push(`+layout${ext}`);
	}

	if (valid_page_exports.has(key)) {
		supported_files.push(`+page${ext}`);
	}

	if (valid_layout_server_exports.has(key)) {
		supported_files.push(`+layout.server${ext}`);
	}

	if (valid_page_server_exports.has(key)) {
		supported_files.push(`+page.server${ext}`);
	}

	if (valid_server_exports.has(key)) {
		supported_files.push(`+server${ext}`);
	}

	if (supported_files.length > 0) {
		return `'${key}' is a valid export in ${supported_files.slice(0, -1).join(', ')}${
			supported_files.length > 1 ? ' or ' : ''
		}${supported_files.at(-1)}`;
	}
}

const valid_layout_exports = new Set([
	'load',
	'prerender',
	'csr',
	'ssr',
	'trailingSlash',
	'config'
]);
const valid_page_exports = new Set([...valid_layout_exports, 'entries']);
const valid_layout_server_exports = new Set([...valid_layout_exports]);
const valid_page_server_exports = new Set([...valid_layout_server_exports, 'actions', 'entries']);
const valid_server_exports = new Set([
	'GET',
	'POST',
	'PATCH',
	'PUT',
	'DELETE',
	'OPTIONS',
	'HEAD',
	'fallback',
	'prerender',
	'trailingSlash',
	'config',
	'entries'
]);

export const validate_layout_exports = validator(valid_layout_exports);
export const validate_page_exports = validator(valid_page_exports);
export const validate_layout_server_exports = validator(valid_layout_server_exports);
export const validate_page_server_exports = validator(valid_page_server_exports);
export const validate_server_exports = validator(valid_server_exports);
