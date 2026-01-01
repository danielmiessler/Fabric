/**
 * Read a value from `sessionStorage`
 * @param {string} key
 * @param {(value: string) => any} parse
 */
/*@__NO_SIDE_EFFECTS__*/
export function get(key, parse = JSON.parse) {
	try {
		return parse(sessionStorage[key]);
	} catch {
		// do nothing
	}
}

/**
 * Write a value to `sessionStorage`
 * @param {string} key
 * @param {any} value
 * @param {(value: any) => string} stringify
 */
export function set(key, value, stringify = JSON.stringify) {
	const data = stringify(value);
	try {
		sessionStorage[key] = data;
	} catch {
		// do nothing
	}
}
