/**
 * @returns {import('types').Deferred & { promise: Promise<any> }}}
 */
function defer() {
	let fulfil;
	let reject;

	const promise = new Promise((f, r) => {
		fulfil = f;
		reject = r;
	});

	// @ts-expect-error
	return { promise, fulfil, reject };
}

/**
 * Create an async iterator and a function to push values into it
 * @returns {{
 *   iterator: AsyncIterable<any>;
 *   push: (value: any) => void;
 *   done: () => void;
 * }}
 */
export function create_async_iterator() {
	const deferred = [defer()];

	return {
		iterator: {
			[Symbol.asyncIterator]() {
				return {
					next: async () => {
						const next = await deferred[0].promise;
						if (!next.done) deferred.shift();
						return next;
					}
				};
			}
		},
		push: (value) => {
			deferred[deferred.length - 1].fulfil({
				value,
				done: false
			});
			deferred.push(defer());
		},
		done: () => {
			deferred[deferred.length - 1].fulfil({ done: true });
		}
	};
}
