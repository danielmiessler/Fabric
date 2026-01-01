/**
 * Removes nullish values from an array.
 *
 * @template T
 * @param {Array<T>} arr
 */
export function compact(arr) {
	return arr.filter(/** @returns {val is NonNullable<T>} */ (val) => val != null);
}
