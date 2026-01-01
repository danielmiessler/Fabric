import { readFileSync } from 'node:fs';
import { dirname } from 'node:path';
import { findClosestPkgJsonPath } from 'vitefu';
import { normalizePath } from 'vite';

/**
 * @typedef {{
 * 	name: string;
 * 	version: string;
 * 	svelte?: string;
 * 	path: string;
 * }} PackageInfo
 */

/**
 * @class
 */
export class VitePluginSvelteCache {
	/** @type {Map<string, import('../types/compile.d.ts').Code>} */
	#css = new Map();
	/** @type {Map<string, import('../types/compile.d.ts').Code>} */
	#js = new Map();
	/** @type {Map<string, string[]>} */
	#dependencies = new Map();
	/** @type {Map<string, Set<string>>} */
	#dependants = new Map();
	/** @type {Map<string, any>} */
	#errors = new Map();
	/** @type {PackageInfo[]} */
	#packageInfos = [];

	/**
	 * @param {import('../types/compile.d.ts').CompileData} compileData
	 */
	update(compileData) {
		this.#errors.delete(compileData.normalizedFilename);
		this.#updateCSS(compileData);
		this.#updateJS(compileData);
		this.#updateDependencies(compileData);
	}

	/**
	 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
	 * @returns {boolean}
	 */
	has(svelteRequest) {
		const id = svelteRequest.normalizedFilename;
		return this.#errors.has(id) || this.#js.has(id) || this.#css.has(id);
	}

	/**
	 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
	 * @param {any} error
	 */
	setError(svelteRequest, error) {
		// keep dependency info, otherwise errors in dependants would not trigger an update after fixing
		// because they are no longer watched
		this.remove(svelteRequest, true);
		this.#errors.set(svelteRequest.normalizedFilename, error);
	}

	/**
	 * @param {import('../types/compile.d.ts').CompileData} compileData
	 */
	#updateCSS(compileData) {
		this.#css.set(compileData.normalizedFilename, compileData.compiled.css);
	}

	/**
	 * @param {import('../types/compile.d.ts').CompileData} compileData
	 */
	#updateJS(compileData) {
		if (!compileData.ssr) {
			// do not cache SSR js
			this.#js.set(compileData.normalizedFilename, compileData.compiled.js);
		}
	}

	/**
	 * @param {import('../types/compile.d.ts').CompileData} compileData
	 */
	#updateDependencies(compileData) {
		const id = compileData.normalizedFilename;
		const prevDependencies = this.#dependencies.get(id) || [];
		const dependencies = compileData.dependencies;
		this.#dependencies.set(id, dependencies);
		const removed = prevDependencies.filter((d) => !dependencies.includes(d));
		const added = dependencies.filter((d) => !prevDependencies.includes(d));
		added.forEach((d) => {
			if (!this.#dependants.has(d)) {
				this.#dependants.set(d, new Set());
			}
			/** @type {Set<string>} */ (this.#dependants.get(d)).add(compileData.filename);
		});
		removed.forEach((d) => {
			/** @type {Set<string>} */ (this.#dependants.get(d)).delete(compileData.filename);
		});
	}

	/**
	 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
	 * @param {boolean} [keepDependencies]
	 * @returns {boolean}
	 */
	remove(svelteRequest, keepDependencies = false) {
		const id = svelteRequest.normalizedFilename;
		let removed = false;
		if (this.#errors.delete(id)) {
			removed = true;
		}
		if (this.#js.delete(id)) {
			removed = true;
		}
		if (this.#css.delete(id)) {
			removed = true;
		}
		if (!keepDependencies) {
			const dependencies = this.#dependencies.get(id);
			if (dependencies) {
				removed = true;
				dependencies.forEach((d) => {
					const dependants = this.#dependants.get(d);
					if (dependants && dependants.has(svelteRequest.filename)) {
						dependants.delete(svelteRequest.filename);
					}
				});
				this.#dependencies.delete(id);
			}
		}

		return removed;
	}

	/**
	 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
	 * @returns {import('../types/compile.d.ts').Code | undefined}
	 */
	getCSS(svelteRequest) {
		return this.#css.get(svelteRequest.normalizedFilename);
	}

	/**
	 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
	 * @returns {import('../types/compile.d.ts').Code | undefined}
	 */
	getJS(svelteRequest) {
		if (!svelteRequest.ssr) {
			// SSR js isn't cached
			return this.#js.get(svelteRequest.normalizedFilename);
		}
	}
	/**
	 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
	 * @returns {any}
	 */
	getError(svelteRequest) {
		return this.#errors.get(svelteRequest.normalizedFilename);
	}

	/**
	 * @param {string} path
	 * @returns {string[]}
	 */
	getDependants(path) {
		const dependants = this.#dependants.get(path);
		return dependants ? [...dependants] : [];
	}

	/**
	 * @param {string} file
	 * @returns {Promise<PackageInfo>}
	 */
	async getPackageInfo(file) {
		let info = this.#packageInfos.find((pi) => file.startsWith(pi.path));
		if (!info) {
			info = await findPackageInfo(file);
			this.#packageInfos.push(info);
		}
		return info;
	}
}

/**
 * utility to get some info from the closest package.json with a "name" set
 *
 * @param {string} file to find info for
 * @returns {Promise<PackageInfo>}
 */
async function findPackageInfo(file) {
	/** @type {PackageInfo} */
	const info = {
		name: '$unknown',
		version: '0.0.0-unknown',
		path: '$unknown'
	};
	let path = await findClosestPkgJsonPath(file, (pkgPath) => {
		const pkg = JSON.parse(readFileSync(pkgPath, 'utf-8'));
		if (pkg.name != null) {
			info.name = pkg.name;
			if (pkg.version != null) {
				info.version = pkg.version;
			}
			info.svelte = pkg.svelte;
			return true;
		}
		return false;
	});
	// return normalized path with appended '/' so .startsWith works for future file checks
	path = normalizePath(dirname(path ?? file)) + '/';
	info.path = path;
	return info;
}
