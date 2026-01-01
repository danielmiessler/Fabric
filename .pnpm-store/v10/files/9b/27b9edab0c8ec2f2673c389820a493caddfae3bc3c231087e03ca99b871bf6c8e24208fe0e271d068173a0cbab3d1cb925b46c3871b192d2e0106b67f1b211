/**
 * `$env/dynamic/private`
 * @type {Record<string, string>}
 */
export let private_env = {};

/**
 * `$env/dynamic/public`. When prerendering, this will be a proxy that forbids reads
 * @type {Record<string, string>}
 */
export let public_env = {};

/**
 * The same as `public_env`, but without the proxy. Use for `%sveltekit.env.PUBLIC_FOO%`
 * @type {Record<string, string>}
 */
export let safe_public_env = {};

/** @param {any} error */
export let fix_stack_trace = (error) => error?.stack;

/** @type {(environment: Record<string, string>) => void} */
export function set_private_env(environment) {
	private_env = environment;
}

/** @type {(environment: Record<string, string>) => void} */
export function set_public_env(environment) {
	public_env = environment;
}

/** @type {(environment: Record<string, string>) => void} */
export function set_safe_public_env(environment) {
	safe_public_env = environment;
}

/** @param {(error: Error) => string} value */
export function set_fix_stack_trace(value) {
	fix_stack_trace = value;
}
