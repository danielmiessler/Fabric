import buffer from 'node:buffer';
import { webcrypto as crypto } from 'node:crypto';

// `buffer.File` was added in Node 18.13.0 while the `File` global was added in Node 20.0.0
const File = /** @type {import('node:buffer') & { File?: File}} */ (buffer).File;

/** @type {Record<string, any>} */
const globals = {
	crypto,
	File
};

// exported for dev/preview and node environments
/**
 * Make various web APIs available as globals:
 * - `crypto`
 * - `File`
 */
export function installPolyfills() {
	for (const name in globals) {
		if (name in globalThis) continue;

		Object.defineProperty(globalThis, name, {
			enumerable: true,
			configurable: true,
			writable: true,
			value: globals[name]
		});
	}
}
