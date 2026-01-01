/**
 * @param {Record<string, string>} env
 * @param {string} allowed
 * @param {string} disallowed
 * @returns {Record<string, string>}
 */
export function filter_env(env, allowed, disallowed) {
	return Object.fromEntries(
		Object.entries(env).filter(
			([k]) => k.startsWith(allowed) && (disallowed === '' || !k.startsWith(disallowed))
		)
	);
}
