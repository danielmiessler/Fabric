/** @type {Set<string>} */
const VALID_CACHE_CONTROL_DIRECTIVES = new Set([
	'max-age',
	'public',
	'private',
	'no-cache',
	'no-store',
	'must-revalidate',
	'proxy-revalidate',
	's-maxage',
	'immutable',
	'stale-while-revalidate',
	'stale-if-error',
	'no-transform',
	'only-if-cached',
	'max-stale',
	'min-fresh'
]);

const CONTENT_TYPE_PATTERN =
	/^(application|audio|example|font|haptics|image|message|model|multipart|text|video|x-[a-z]+)\/[-+.\w]+$/i;

/** @type {Record<string, (value: string) => void>} */
const HEADER_VALIDATORS = {
	'cache-control': (value) => {
		const error_suffix = `(While parsing "${value}".)`;
		const parts = value.split(',').map((part) => part.trim());
		if (parts.some((part) => !part)) {
			throw new Error(`\`cache-control\` header contains empty directives. ${error_suffix}`);
		}

		const directives = parts.map((part) => part.split('=')[0].toLowerCase());
		const invalid = directives.find((directive) => !VALID_CACHE_CONTROL_DIRECTIVES.has(directive));
		if (invalid) {
			throw new Error(
				`Invalid cache-control directive "${invalid}". Did you mean one of: ${[...VALID_CACHE_CONTROL_DIRECTIVES].join(', ')}? ${error_suffix}`
			);
		}
	},

	'content-type': (value) => {
		const type = value.split(';')[0].trim();
		const error_suffix = `(While parsing "${value}".)`;
		if (!CONTENT_TYPE_PATTERN.test(type)) {
			throw new Error(`Invalid content-type value "${type}". ${error_suffix}`);
		}
	}
};

/**
 * @param {Record<string, string>} headers
 */
export function validateHeaders(headers) {
	for (const [key, value] of Object.entries(headers)) {
		const validator = HEADER_VALIDATORS[key.toLowerCase()];
		try {
			validator?.(value);
		} catch (error) {
			if (error instanceof Error) {
				console.warn(`[SvelteKit] ${error.message}`);
			}
		}
	}
}
