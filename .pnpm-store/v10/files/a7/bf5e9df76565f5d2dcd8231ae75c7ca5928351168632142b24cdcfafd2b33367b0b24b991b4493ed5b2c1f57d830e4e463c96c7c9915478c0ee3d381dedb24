/** @see https://github.com/microsoft/TypeScript/blob/904e7dd97dc8da1352c8e05d70829dff17c73214/src/lib/es2024.promise.d.ts */

/**
 * @template T
 * @typedef {{
 *   promise: Promise<T>;
 *   resolve: (value: T | PromiseLike<T>) => void;
 *   reject: (reason?: any) => void;
 * }} PromiseWithResolvers<T>
 */

/**
 * TODO: Whenever Node >21 is minimum supported version, we can use `Promise.withResolvers` to avoid this ceremony
 *
 * @template T
 * @returns {PromiseWithResolvers<T>}
 */
export function with_resolvers() {
	let resolve;
	let reject;

	const promise = new Promise((res, rej) => {
		resolve = res;
		reject = rej;
	});

	// @ts-expect-error `resolve` and `reject` are assigned!
	return { promise, resolve, reject };
}
