declare global {
	const __SVELTEKIT_ADAPTER_NAME__: string;
	const __SVELTEKIT_APP_VERSION_FILE__: string;
	const __SVELTEKIT_APP_VERSION_POLL_INTERVAL__: number;
	const __SVELTEKIT_DEV__: boolean;
	const __SVELTEKIT_EMBEDDED__: boolean;
	/** True if `config.kit.router.resolution === 'client'` */
	const __SVELTEKIT_CLIENT_ROUTING__: boolean;
	/**
	 * This makes the use of specific features visible at both dev and build time, in such a
	 * way that we can error when they are not supported by the target platform.
	 *
	 * During dev, `globalThis.__SVELTEKIT_TRACK__` is a function that grabs the current `event`
	 * and route `config` (from an AsyncLocalStorage instance) and calls the relevant `supports`
	 * function on the adapter (e.g. `adapter.supports.read(...)`).
	 *
	 * At build time, if the function containing the `__SVELTEKIT_TRACK__` call is untreeshaken,
	 * we locate it in the `renderChunk` build hook and a) make a note of the chunk that contains
	 * it and b) replace it with a comment. Later, we can use this information to establish
	 * which routes use which feature, and use the same `adapter.supports.read(...)` function
	 * to throw an error if the feature would fail in production.
	 */
	var __SVELTEKIT_TRACK__: (label: string) => void;
	var Bun: object;
	var Deno: object;
}

export {};
