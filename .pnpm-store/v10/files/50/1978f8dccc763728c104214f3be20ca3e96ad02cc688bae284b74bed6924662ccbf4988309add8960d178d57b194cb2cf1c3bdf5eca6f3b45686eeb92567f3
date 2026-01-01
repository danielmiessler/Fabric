import path from 'node:path';

const IS_WINDOWS = process.platform === 'win32';

/**
 * @typedef {{
 *  file?: string;
 *  sources?: string[];
 *  sourceRoot?: string;
 * }} SourceMapFileRefs
 */

/**
 * convert absolute paths in sourcemap file refs to their relative equivalents to avoid leaking fs info
 *
 * map is modified in place.
 *
 * @param {SourceMapFileRefs | undefined} map sourcemap
 * @param {string} filename absolute path to file the sourcemap is for
 */
export function mapToRelative(map, filename) {
	if (!map) {
		return;
	}
	const sourceRoot = map.sourceRoot;
	const dirname = path.dirname(filename);

	/** @type {(s: string) => string} */
	const toRelative = (s) => {
		if (!s) {
			return s;
		}
		/** @type {string} */
		let sourcePath;
		if (s.startsWith('file:///')) {
			// windows has file:///C:/foo and posix has file:///foo, so we have to remove one extra on windows
			sourcePath = s.slice(IS_WINDOWS ? 8 : 7);
		} else if (sourceRoot) {
			const sep = sourceRoot[sourceRoot.length - 1] === '/' || s[0] === '/' ? '' : '/';
			sourcePath = `${sourceRoot}${sep}${s}`;
		} else {
			sourcePath = s;
		}
		return path.isAbsolute(sourcePath) ? path.relative(dirname, sourcePath) : sourcePath;
	};

	if (map.file) {
		map.file = path.basename(filename);
	}
	if (map.sources) {
		map.sources = map.sources.map(toRelative);
	}
	if (map.sourceRoot) {
		// we have prepended sourceRoot and computed relative paths from it
		// remove it here to avoid downstream processing prepending it again
		delete map.sourceRoot;
	}
}

/**
 * vitePreprocess uses an extra lang extension to tell vite about the type of preprocessor to use
 * This function removes it afterwards to get back working file refs
 *
 * map is modified in place.
 *
 * @param {SourceMapFileRefs | undefined} map the output sourcemap
 * @param {string} suffix the suffix to remove
 */
export function removeLangSuffix(map, suffix) {
	if (!map) {
		return;
	}
	/** @type {(s:string)=> string} */
	const removeSuffix = (s) => (s?.endsWith(suffix) ? s.slice(0, -1 * suffix.length) : s);
	if (map.file) {
		map.file = removeSuffix(map.file);
	}
	if (map.sources) {
		map.sources = map.sources.map(removeSuffix);
	}
}
