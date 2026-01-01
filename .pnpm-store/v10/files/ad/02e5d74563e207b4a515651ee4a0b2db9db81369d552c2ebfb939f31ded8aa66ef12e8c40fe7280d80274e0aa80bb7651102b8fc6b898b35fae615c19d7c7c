import fs from 'node:fs';
import path from 'node:path';

/** @param {string} dir */
export function mkdirp(dir) {
	try {
		fs.mkdirSync(dir, { recursive: true });
	} catch (/** @type {any} */ e) {
		if (e.code === 'EEXIST') {
			if (!fs.statSync(dir).isDirectory()) {
				throw new Error(`Cannot create directory ${dir}, a file already exists at this position`);
			}
			return;
		}
		throw e;
	}
}

/** @param {string} path */
export function rimraf(path) {
	fs.rmSync(path, { force: true, recursive: true });
}

/**
 * @param {string} source
 * @param {string} target
 * @param {{
 *   filter?: (basename: string) => boolean;
 *   replace?: Record<string, string>;
 * }} opts
 */
export function copy(source, target, opts = {}) {
	if (!fs.existsSync(source)) return [];

	/** @type {string[]} */
	const files = [];

	const prefix = posixify(target) + '/';

	const regex = opts.replace
		? new RegExp(`\\b(${Object.keys(opts.replace).join('|')})\\b`, 'g')
		: null;

	/**
	 * @param {string} from
	 * @param {string} to
	 */
	function go(from, to) {
		if (opts.filter && !opts.filter(path.basename(from))) return;

		const stats = fs.statSync(from);

		if (stats.isDirectory()) {
			fs.readdirSync(from).forEach((file) => {
				go(path.join(from, file), path.join(to, file));
			});
		} else {
			mkdirp(path.dirname(to));

			if (opts.replace) {
				const data = fs.readFileSync(from, 'utf-8');
				fs.writeFileSync(
					to,
					data.replace(
						/** @type {RegExp} */ (regex),
						(_match, key) => /** @type {Record<string, string>} */ (opts.replace)[key]
					)
				);
			} else {
				fs.copyFileSync(from, to);
			}

			files.push(to === target ? posixify(path.basename(to)) : posixify(to).replace(prefix, ''));
		}
	}

	go(source, target);

	return files;
}

/**
 * Get a list of all files in a directory
 * @param {string} cwd - the directory to walk
 * @param {boolean} [dirs] - whether to include directories in the result
 * @returns {string[]} a list of all found files (and possibly directories) relative to `cwd`
 */
export function walk(cwd, dirs = false) {
	/** @type {string[]} */
	const all_files = [];

	/** @param {string} dir */
	function walk_dir(dir) {
		const files = fs.readdirSync(path.join(cwd, dir));

		for (const file of files) {
			const joined = path.join(dir, file);
			const stats = fs.statSync(path.join(cwd, joined));
			if (stats.isDirectory()) {
				if (dirs) all_files.push(joined);
				walk_dir(joined);
			} else {
				all_files.push(joined);
			}
		}
	}

	return (walk_dir(''), all_files);
}

/** @param {string} str */
export function posixify(str) {
	return str.replace(/\\/g, '/');
}

/**
 * Like `path.join`, but posixified and with a leading `./` if necessary
 * @param {string[]} str
 */
export function join_relative(...str) {
	let result = posixify(path.join(...str));
	if (!result.startsWith('.')) {
		result = `./${result}`;
	}
	return result;
}

/**
 * Like `path.relative`, but always posixified and with a leading `./` if necessary.
 * Useful for JS imports so the path can safely reside inside of `node_modules`.
 * Otherwise paths could be falsely interpreted as package paths.
 * @param {string} from
 * @param {string} to
 */
export function relative_path(from, to) {
	return join_relative(path.relative(from, to));
}

/**
 * Prepend given path with `/@fs` prefix
 * @param {string} str
 */
export function to_fs(str) {
	str = posixify(str);
	return `/@fs${
		// Windows/Linux separation - Windows starts with a drive letter, we need a / in front there
		str.startsWith('/') ? '' : '/'
	}${str}`;
}

/**
 * Removes `/@fs` prefix from given path and posixifies it
 * @param {string} str
 */
export function from_fs(str) {
	str = posixify(str);
	if (!str.startsWith('/@fs')) return str;

	str = str.slice(4);
	// Windows/Linux separation - Windows starts with a drive letter, we need to strip the additional / here
	return str[2] === ':' && /[A-Z]/.test(str[1]) ? str.slice(1) : str;
}

/**
 * Given an entry point like [cwd]/src/hooks, returns a filename like [cwd]/src/hooks.js or [cwd]/src/hooks/index.js
 * @param {string} entry
 * @returns {string|null}
 */
export function resolve_entry(entry) {
	if (fs.existsSync(entry)) {
		const stats = fs.statSync(entry);
		if (stats.isFile()) {
			return entry;
		}

		const index = path.join(entry, 'index');
		if (fs.existsSync(index + '.js') || fs.existsSync(index + '.ts')) {
			return resolve_entry(index);
		}
	}

	const dir = path.dirname(entry);

	if (fs.existsSync(dir)) {
		const base = path.basename(entry);
		const files = fs.readdirSync(dir);
		const found = files.find((file) => {
			return file.replace(/\.(js|ts)$/, '') === base && fs.statSync(path.join(dir, file)).isFile();
		});

		if (found) return path.join(dir, found);
	}

	return null;
}

/** @param {string} file */
export function read(file) {
	return fs.readFileSync(file, 'utf-8');
}
