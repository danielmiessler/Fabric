import { text_encoder } from '../../utils.js';

/**
 * SHA-256 hashing function adapted from https://bitwiseshiftleft.github.io/sjcl
 * modified and redistributed under BSD license
 * @param {string} data
 */
export function sha256(data) {
	if (!key[0]) precompute();

	const out = init.slice(0);
	const array = encode(data);

	for (let i = 0; i < array.length; i += 16) {
		const w = array.subarray(i, i + 16);

		let tmp;
		let a;
		let b;

		let out0 = out[0];
		let out1 = out[1];
		let out2 = out[2];
		let out3 = out[3];
		let out4 = out[4];
		let out5 = out[5];
		let out6 = out[6];
		let out7 = out[7];

		/* Rationale for placement of |0 :
		 * If a value can overflow is original 32 bits by a factor of more than a few
		 * million (2^23 ish), there is a possibility that it might overflow the
		 * 53-bit mantissa and lose precision.
		 *
		 * To avoid this, we clamp back to 32 bits by |'ing with 0 on any value that
		 * propagates around the loop, and on the hash state out[]. I don't believe
		 * that the clamps on out4 and on out0 are strictly necessary, but it's close
		 * (for out4 anyway), and better safe than sorry.
		 *
		 * The clamps on out[] are necessary for the output to be correct even in the
		 * common case and for short inputs.
		 */

		for (let i = 0; i < 64; i++) {
			// load up the input word for this round

			if (i < 16) {
				tmp = w[i];
			} else {
				a = w[(i + 1) & 15];

				b = w[(i + 14) & 15];

				tmp = w[i & 15] =
					(((a >>> 7) ^ (a >>> 18) ^ (a >>> 3) ^ (a << 25) ^ (a << 14)) +
						((b >>> 17) ^ (b >>> 19) ^ (b >>> 10) ^ (b << 15) ^ (b << 13)) +
						w[i & 15] +
						w[(i + 9) & 15]) |
					0;
			}

			tmp =
				tmp +
				out7 +
				((out4 >>> 6) ^ (out4 >>> 11) ^ (out4 >>> 25) ^ (out4 << 26) ^ (out4 << 21) ^ (out4 << 7)) +
				(out6 ^ (out4 & (out5 ^ out6))) +
				key[i]; // | 0;

			// shift register
			out7 = out6;
			out6 = out5;
			out5 = out4;

			out4 = (out3 + tmp) | 0;

			out3 = out2;
			out2 = out1;
			out1 = out0;

			out0 =
				(tmp +
					((out1 & out2) ^ (out3 & (out1 ^ out2))) +
					((out1 >>> 2) ^
						(out1 >>> 13) ^
						(out1 >>> 22) ^
						(out1 << 30) ^
						(out1 << 19) ^
						(out1 << 10))) |
				0;
		}

		out[0] = (out[0] + out0) | 0;
		out[1] = (out[1] + out1) | 0;
		out[2] = (out[2] + out2) | 0;
		out[3] = (out[3] + out3) | 0;
		out[4] = (out[4] + out4) | 0;
		out[5] = (out[5] + out5) | 0;
		out[6] = (out[6] + out6) | 0;
		out[7] = (out[7] + out7) | 0;
	}

	const bytes = new Uint8Array(out.buffer);
	reverse_endianness(bytes);

	return btoa(String.fromCharCode(...bytes));
}

/** The SHA-256 initialization vector */
const init = new Uint32Array(8);

/** The SHA-256 hash key */
const key = new Uint32Array(64);

/** Function to precompute init and key. */
function precompute() {
	/** @param {number} x */
	function frac(x) {
		return (x - Math.floor(x)) * 0x100000000;
	}

	let prime = 2;

	for (let i = 0; i < 64; prime++) {
		let is_prime = true;

		for (let factor = 2; factor * factor <= prime; factor++) {
			if (prime % factor === 0) {
				is_prime = false;

				break;
			}
		}

		if (is_prime) {
			if (i < 8) {
				init[i] = frac(prime ** (1 / 2));
			}

			key[i] = frac(prime ** (1 / 3));

			i++;
		}
	}
}

/** @param {Uint8Array} bytes */
function reverse_endianness(bytes) {
	for (let i = 0; i < bytes.length; i += 4) {
		const a = bytes[i + 0];
		const b = bytes[i + 1];
		const c = bytes[i + 2];
		const d = bytes[i + 3];

		bytes[i + 0] = d;
		bytes[i + 1] = c;
		bytes[i + 2] = b;
		bytes[i + 3] = a;
	}
}

/** @param {string} str */
function encode(str) {
	const encoded = text_encoder.encode(str);
	const length = encoded.length * 8;

	// result should be a multiple of 512 bits in length,
	// with room for a 1 (after the data) and two 32-bit
	// words containing the original input bit length
	const size = 512 * Math.ceil((length + 65) / 512);
	const bytes = new Uint8Array(size / 8);
	bytes.set(encoded);

	// append a 1
	bytes[encoded.length] = 0b10000000;

	reverse_endianness(bytes);

	// add the input bit length
	const words = new Uint32Array(bytes.buffer);
	words[words.length - 2] = Math.floor(length / 0x100000000); // this will always be zero for us
	words[words.length - 1] = length;

	return words;
}
