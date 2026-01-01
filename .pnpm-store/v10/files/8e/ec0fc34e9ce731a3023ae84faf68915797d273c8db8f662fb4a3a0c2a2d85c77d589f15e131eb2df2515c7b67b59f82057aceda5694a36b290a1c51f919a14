import fs from 'node:fs';
import path from 'node:path';
import { mkdirp } from '../../utils/filesystem.js';
import { import_peer } from '../../utils/import.js';

/** @type {{ VERSION: string }} */
const { VERSION } = await import_peer('svelte/compiler');

/** @type {Map<string, string>} */
const previous_contents = new Map();

/**
 * @param {string} file
 * @param {string} code
 */
export function write_if_changed(file, code) {
	if (code !== previous_contents.get(file)) {
		write(file, code);
	}
}

/**
 * @param {string} file
 * @param {string} code
 */
export function write(file, code) {
	previous_contents.set(file, code);
	mkdirp(path.dirname(file));
	fs.writeFileSync(file, code);
}

/** @type {WeakMap<TemplateStringsArray, { strings: string[], indents: string[] }>} */
const dedent_map = new WeakMap();

/**
 * Allows indenting template strings without the extra indentation ending up in the result.
 * Still allows indentation of lines relative to one another in the template string.
 * @param {TemplateStringsArray} strings
 * @param {any[]} values
 */
export function dedent(strings, ...values) {
	let dedented = dedent_map.get(strings);

	if (!dedented) {
		const indentation = /** @type {RegExpExecArray} */ (/\n?([ \t]*)/.exec(strings[0]))[1];
		const pattern = new RegExp(`^${indentation}`, 'gm');

		dedented = {
			strings: strings.map((str) => str.replace(pattern, '')),
			indents: []
		};

		let current = '\n';

		for (let i = 0; i < values.length; i += 1) {
			const string = dedented.strings[i];
			const match = /\n([ \t]*)$/.exec(string);

			if (match) current = match[0];
			dedented.indents[i] = current;
		}

		dedent_map.set(strings, dedented);
	}

	let str = dedented.strings[0];
	for (let i = 0; i < values.length; i += 1) {
		str += String(values[i]).replace(/\n/g, dedented.indents[i]) + dedented.strings[i + 1];
	}

	str = str.trim();

	return str;
}

export function isSvelte5Plus() {
	return Number(VERSION[0]) >= 5;
}
