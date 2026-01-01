declare global {
	const __SVELTEKIT_ADAPTER_NAME__: string;
	const __SVELTEKIT_APP_DIR__: string;
	const __SVELTEKIT_APP_VERSION_FILE__: string;
	const __SVELTEKIT_APP_VERSION_POLL_INTERVAL__: number;
	const __SVELTEKIT_EMBEDDED__: boolean;
	const __SVELTEKIT_PATHS_ASSETS__: string;
	const __SVELTEKIT_PATHS_BASE__: string;
	const __SVELTEKIT_PATHS_RELATIVE__: boolean;
	/** True if `config.kit.experimental.instrumentation.server` is `true` */
	const __SVELTEKIT_SERVER_TRACING_ENABLED__: boolean;
	/** true if corresponding config option is set to true */
	const __SVELTEKIT_EXPERIMENTAL__REMOTE_FUNCTIONS__: boolean;
	/** True if `config.kit.router.resolution === 'client'` */
	const __SVELTEKIT_CLIENT_ROUTING__: boolean;
	/** True if `config.kit.router.type === 'hash'` */
	const __SVELTEKIT_HASH_ROUTING__: boolean;
	/**
	 * True if any node in the manifest has a server load function.
	 * Used for treeshaking server load code from client bundles when no server loads exist.
	 */
	const __SVELTEKIT_HAS_SERVER_LOAD__: boolean;
	/**
	 * True if any node in the manifest has a universal load function.
	 * Used for treeshaking universal load code from client bundles when no universal loads exist.
	 */
	const __SVELTEKIT_HAS_UNIVERSAL_LOAD__: boolean;
	/** The `__sveltekit_abc123` object in the init `<script>` */
	const __SVELTEKIT_PAYLOAD__: {
		/** The basepath, usually relative to the current page */
		base: string;
		/** Path to externally-hosted assets */
		assets?: string;
		/** Public environment variables */
		env?: Record<string, string>;
		/** Serialized data from remote functions */
		data?: Record<string, any>;
		/** Create a placeholder promise */
		defer?: (id: number) => Promise<any>;
		/** Resolve a placeholder promise */
		resolve?: (data: { id: number; data: any; error: any }) => void;
	};
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
