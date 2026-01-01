import * as crypto from 'node:crypto';

const hashes = Object.create(null);

//TODO shorter?
const hash_length = 12;

/**
 * replaces +/= in base64 output so they don't interfere
 *
 * @param {string} input
 * @returns {string} base64 hash safe to use in any context
 */
export function safeBase64Hash(input) {
	if (hashes[input]) {
		return hashes[input];
	}
	//TODO if performance really matters, use a faster one like xx-hash etc.
	// should be evenly distributed because short input length and similarities in paths could cause collisions otherwise
	// OR DON'T USE A HASH AT ALL, what about a simple counter?
	const md5 = crypto.createHash('md5');
	md5.update(input);
	const hash = toSafe(md5.digest('base64')).slice(0, hash_length);
	hashes[input] = hash;
	return hash;
}

/** @type {Record<string, string>} */
const replacements = {
	'+': '-',
	'/': '_',
	'=': ''
};

const replaceRE = new RegExp(`[${Object.keys(replacements).join('')}]`, 'g');

/**
 * @param {string} base64
 * @returns {string}
 */
function toSafe(base64) {
	return base64.replace(replaceRE, (x) => replacements[x]);
}
