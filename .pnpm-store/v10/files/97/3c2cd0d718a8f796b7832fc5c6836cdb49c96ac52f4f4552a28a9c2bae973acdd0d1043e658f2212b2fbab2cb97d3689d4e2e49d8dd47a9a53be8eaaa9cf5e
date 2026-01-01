/**
 * Hash using djb2
 * @param {import('types').StrictBody[]} values
 */
export function hash(...values) {
	let hash = 5381;

	for (const value of values) {
		if (typeof value === 'string') {
			let i = value.length;
			while (i) hash = (hash * 33) ^ value.charCodeAt(--i);
		} else if (ArrayBuffer.isView(value)) {
			const buffer = new Uint8Array(value.buffer, value.byteOffset, value.byteLength);
			let i = buffer.length;
			while (i) hash = (hash * 33) ^ buffer[--i];
		} else {
			throw new TypeError('value must be a string or TypedArray');
		}
	}

	return (hash >>> 0).toString(36);
}
