/**
 * @template T
 * @param {() => T} fn
 */
export function once(fn) {
	let done = false;

	/** @type T */
	let result;

	return () => {
		if (done) return result;
		done = true;
		return (result = fn());
	};
}
