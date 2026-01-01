import fs from 'node:fs';
import path from 'node:path';
import process from 'node:process';
import { fileURLToPath } from 'node:url';
import colors from 'kleur';
import { posixify, to_fs } from '../utils/filesystem.js';

/**
 * Resolved path of the `runtime` directory
 *
 * TODO Windows issue:
 * Vite or sth else somehow sets the driver letter inconsistently to lower or upper case depending on the run environment.
 * In playwright debug mode run through VS Code this a root-to-lowercase conversion is needed in order for the tests to run.
 * If we do this conversion in other cases it has the opposite effect though and fails.
 */
export const runtime_directory = posixify(fileURLToPath(new URL('../runtime', import.meta.url)));

/**
 * This allows us to import SvelteKit internals that aren't exposed via `pkg.exports` in a
 * way that works whether `@sveltejs/kit` is installed inside the project's `node_modules`
 * or in a workspace root
 */
export const runtime_base = runtime_directory.startsWith(process.cwd())
	? `/${path.relative('.', runtime_directory)}`
	: to_fs(runtime_directory);

function noop() {}

/** @param {{ verbose: boolean }} opts */
export function logger({ verbose }) {
	/** @type {import('types').Logger} */
	const log = (msg) => console.log(msg.replace(/^/gm, '  '));

	/** @param {string} msg */
	const err = (msg) => console.error(msg.replace(/^/gm, '  '));

	log.success = (msg) => log(colors.green(`✔ ${msg}`));
	log.error = (msg) => err(colors.bold().red(msg));
	log.warn = (msg) => log(colors.bold().yellow(msg));

	log.minor = verbose ? (msg) => log(colors.grey(msg)) : noop;
	log.info = verbose ? log : noop;

	return log;
}

/** @param {import('types').ManifestData} manifest_data */
export function get_mime_lookup(manifest_data) {
	/** @type {Record<string, string>} */
	const mime = {};

	manifest_data.assets.forEach((asset) => {
		if (asset.type) {
			const ext = path.extname(asset.file);
			mime[ext] = asset.type;
		}
	});

	return mime;
}

/**
 * @param {string} dir
 * @param {(file: string) => boolean} [filter]
 */
export function list_files(dir, filter) {
	/** @type {string[]} */
	const files = [];

	/** @param {string} current */
	function walk(current) {
		for (const file of fs.readdirSync(path.resolve(dir, current))) {
			const child = path.posix.join(current, file);
			if (fs.statSync(path.resolve(dir, child)).isDirectory()) {
				walk(child);
			} else {
				if (!filter || filter(child)) {
					files.push(child);
				}
			}
		}
	}

	if (fs.existsSync(dir)) walk('');

	return files;
}
