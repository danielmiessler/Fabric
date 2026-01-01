import { createFilter, normalizePath } from 'vite';
import * as fs from 'node:fs';
import { log } from './log.js';

const VITE_FS_PREFIX = '/@fs/';
const IS_WINDOWS = process.platform === 'win32';

const SUPPORTED_COMPILER_OPTIONS = [
	'generate',
	'dev',
	'css',
	'hydratable',
	'customElement',
	'immutable',
	'enableSourcemap'
];
const TYPES_WITH_COMPILER_OPTIONS = ['style', 'script', 'all'];

/**
 * @param {string} id
 * @returns {{ filename: string, rawQuery: string }}
 */
function splitId(id) {
	const parts = id.split('?', 2);
	const filename = parts[0];
	const rawQuery = parts[1];
	return { filename, rawQuery };
}

/**
 * @param {string} id
 * @param {string} filename
 * @param {string} rawQuery
 * @param {string} root
 * @param {number} timestamp
 * @param {boolean} ssr
 * @returns {import('../types/id.d.ts').SvelteRequest | undefined}
 */
function parseToSvelteRequest(id, filename, rawQuery, root, timestamp, ssr) {
	const query = parseRequestQuery(rawQuery);
	const rawOrDirect = !!(query.raw || query.direct);
	if (query.url || (!query.svelte && rawOrDirect)) {
		// skip requests with special vite tags
		return;
	}
	const raw = rawOrDirect;
	const normalizedFilename = normalize(filename, root);
	const cssId = createVirtualImportId(filename, root, 'style');

	return {
		id,
		filename,
		normalizedFilename,
		cssId,
		query,
		timestamp,
		ssr,
		raw
	};
}

/**
 * @param {string} filename
 * @param {string} root
 * @param {import('../types/id.d.ts').SvelteQueryTypes} type
 * @returns {string}
 */
function createVirtualImportId(filename, root, type) {
	const parts = ['svelte', `type=${type}`];
	if (type === 'style') {
		parts.push('lang.css');
	}
	if (existsInRoot(filename, root)) {
		filename = root + filename;
	} else if (filename.startsWith(VITE_FS_PREFIX)) {
		filename = IS_WINDOWS
			? filename.slice(VITE_FS_PREFIX.length) // remove /@fs/ from /@fs/C:/...
			: filename.slice(VITE_FS_PREFIX.length - 1); // remove /@fs from /@fs/home/user
	}
	// return same virtual id format as vite-plugin-vue eg ...App.svelte?svelte&type=style&lang.css
	return `${filename}?${parts.join('&')}`;
}

/**
 * @param {string} rawQuery
 * @returns {import('../types/id.d.ts').RequestQuery}
 */
function parseRequestQuery(rawQuery) {
	const query = Object.fromEntries(new URLSearchParams(rawQuery));
	for (const key in query) {
		if (query[key] === '') {
			// @ts-ignore
			query[key] = true;
		}
	}
	const compilerOptions = query.compilerOptions;
	if (compilerOptions) {
		if (!((query.raw || query.direct) && TYPES_WITH_COMPILER_OPTIONS.includes(query.type))) {
			throw new Error(
				`Invalid compilerOptions in query ${rawQuery}. CompilerOptions are only supported for raw or direct queries with type in "${TYPES_WITH_COMPILER_OPTIONS.join(
					', '
				)}" e.g. '?svelte&raw&type=script&compilerOptions={"generate":"ssr","dev":false}`
			);
		}
		try {
			const parsed = JSON.parse(compilerOptions);
			const invalid = Object.keys(parsed).filter(
				(key) => !SUPPORTED_COMPILER_OPTIONS.includes(key)
			);
			if (invalid.length) {
				throw new Error(
					`Invalid compilerOptions in query ${rawQuery}: ${invalid.join(
						', '
					)}. Supported: ${SUPPORTED_COMPILER_OPTIONS.join(', ')}`
				);
			}
			query.compilerOptions = parsed;
		} catch (e) {
			log.error('failed to parse request query compilerOptions', e);
			throw e;
		}
	}

	return /** @type {import('../types/id.d.ts').RequestQuery}*/ query;
}

/**
 * posixify and remove root at start
 *
 * @param {string} filename
 * @param {string} normalizedRoot
 * @returns {string}
 */
function normalize(filename, normalizedRoot) {
	return stripRoot(normalizePath(filename), normalizedRoot);
}

/**
 * @param {string} filename
 * @param {string} root
 * @returns {boolean}
 */
function existsInRoot(filename, root) {
	if (filename.startsWith(VITE_FS_PREFIX)) {
		return false; // vite already tagged it as out of root
	}
	return fs.existsSync(root + filename);
}

/**
 * @param {string} normalizedFilename
 * @param {string} normalizedRoot
 * @returns {string}
 */
function stripRoot(normalizedFilename, normalizedRoot) {
	return normalizedFilename.startsWith(normalizedRoot + '/')
		? normalizedFilename.slice(normalizedRoot.length)
		: normalizedFilename;
}

/**
 * @param {import('../public.d.ts').Options['include'] | undefined} include
 * @param {import('../public.d.ts').Options['exclude'] | undefined} exclude
 * @param {string[]} extensions
 * @returns {(filename: string) => boolean}
 */
function buildFilter(include, exclude, extensions) {
	const rollupFilter = createFilter(include, exclude);
	return (filename) => rollupFilter(filename) && extensions.some((ext) => filename.endsWith(ext));
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {import('../types/id.d.ts').IdParser}
 */
export function buildIdParser(options) {
	const { include, exclude, extensions, root } = options;
	const normalizedRoot = normalizePath(root);
	const filter = buildFilter(include, exclude, extensions ?? []);
	return (id, ssr, timestamp = Date.now()) => {
		const { filename, rawQuery } = splitId(id);
		if (filter(filename)) {
			return parseToSvelteRequest(id, filename, rawQuery, normalizedRoot, timestamp, ssr);
		}
	};
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {import('../types/id.d.ts').ModuleIdParser}
 */
export function buildModuleIdParser(options) {
	const { include, exclude, extensions } = options?.experimental?.compileModule ?? {};
	const root = options.root;
	const normalizedRoot = normalizePath(root);
	const filter = buildFilter(include, exclude, extensions ?? ['.svelte.js', '.svelte.ts']);
	return (id, ssr, timestamp = Date.now()) => {
		const { filename, rawQuery } = splitId(id);
		if (filter(filename)) {
			return parseToSvelteModuleRequest(id, filename, rawQuery, normalizedRoot, timestamp, ssr);
		}
	};
}

/**
 * @param {string} id
 * @param {string} filename
 * @param {string} rawQuery
 * @param {string} root
 * @param {number} timestamp
 * @param {boolean} ssr
 * @returns {import('../types/id.d.ts').SvelteModuleRequest | undefined}
 */
function parseToSvelteModuleRequest(id, filename, rawQuery, root, timestamp, ssr) {
	const query = parseRequestQuery(rawQuery);

	if (query.url || query.raw || query.direct) {
		// skip requests with special vite tags
		return;
	}

	const normalizedFilename = normalize(filename, root);

	return {
		id,
		filename,
		normalizedFilename,
		query,
		timestamp,
		ssr
	};
}
