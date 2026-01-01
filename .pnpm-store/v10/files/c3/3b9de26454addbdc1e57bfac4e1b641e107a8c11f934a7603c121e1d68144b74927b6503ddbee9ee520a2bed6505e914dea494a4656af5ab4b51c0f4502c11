/**
 * @param {string} text
 * @returns {ArrayBufferLike}
 */
export function b64_decode(text) {
	const d = atob(text);

	const u8 = new Uint8Array(d.length);

	for (let i = 0; i < d.length; i++) {
		u8[i] = d.charCodeAt(i);
	}

	return u8.buffer;
}

/**
 * @param {ArrayBuffer} buffer
 * @returns {string}
 */
export function b64_encode(buffer) {
	if (globalThis.Buffer) {
		return Buffer.from(buffer).toString('base64');
	}

	const little_endian = new Uint8Array(new Uint16Array([1]).buffer)[0] > 0;

	// The Uint16Array(Uint8Array(...)) ensures the code points are padded with 0's
	return btoa(
		new TextDecoder(little_endian ? 'utf-16le' : 'utf-16be').decode(
			new Uint16Array(new Uint8Array(buffer))
		)
	);
}

/**
 * Like node's path.relative, but without using node
 * @param {string} from
 * @param {string} to
 */
export function get_relative_path(from, to) {
	const from_parts = from.split(/[/\\]/);
	const to_parts = to.split(/[/\\]/);
	from_parts.pop(); // get dirname

	while (from_parts[0] === to_parts[0]) {
		from_parts.shift();
		to_parts.shift();
	}

	let i = from_parts.length;
	while (i--) from_parts[i] = '..';

	return from_parts.concat(to_parts).join('/');
}
