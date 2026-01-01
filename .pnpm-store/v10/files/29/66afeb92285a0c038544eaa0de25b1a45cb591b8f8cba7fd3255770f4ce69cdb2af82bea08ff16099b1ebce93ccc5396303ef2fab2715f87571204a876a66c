/**
 * @param {Record<string, string>} env
 * @param {{
 * 		public_prefix: string;
 * 		private_prefix: string;
 * }} prefixes
 * @returns {Record<string, string>}
 */
export function filter_private_env(env, { public_prefix, private_prefix }) {
	return Object.fromEntries(
		Object.entries(env).filter(
			([k]) =>
				k.startsWith(private_prefix) && (public_prefix === '' || !k.startsWith(public_prefix))
		)
	);
}

/**
 * @param {Record<string, string>} env
 * @param {{
 * 		public_prefix: string;
 *    private_prefix: string;
 * }} prefixes
 * @returns {Record<string, string>}
 */
export function filter_public_env(env, { public_prefix, private_prefix }) {
	return Object.fromEntries(
		Object.entries(env).filter(
			([k]) =>
				k.startsWith(public_prefix) && (private_prefix === '' || !k.startsWith(private_prefix))
		)
	);
}
