import { log } from './log.js';
import { performance } from 'node:perf_hooks';
import { normalizePath } from 'vite';

/** @type {import('../types/vite-plugin-svelte-stats.d.ts').CollectionOptions} */
const defaultCollectionOptions = {
	// log after 500ms and more than one file processed
	logInProgress: (c, now) => now - c.collectionStart > 500 && c.stats.length > 1,
	// always log results
	logResult: () => true
};

/**
 * @param {number} n
 * @returns
 */
function humanDuration(n) {
	// 99.9ms  0.10s
	return n < 100 ? `${n.toFixed(1)}ms` : `${(n / 1000).toFixed(2)}s`;
}

/**
 * @param {import('../types/vite-plugin-svelte-stats.d.ts').PackageStats[]} pkgStats
 * @returns {string}
 */
function formatPackageStats(pkgStats) {
	const statLines = pkgStats.map((pkgStat) => {
		const duration = pkgStat.duration;
		const avg = duration / pkgStat.files;
		return [pkgStat.pkg, `${pkgStat.files}`, humanDuration(duration), humanDuration(avg)];
	});
	statLines.unshift(['package', 'files', 'time', 'avg']);
	const columnWidths = statLines.reduce(
		(widths, row) => {
			for (let i = 0; i < row.length; i++) {
				const cell = row[i];
				if (widths[i] < cell.length) {
					widths[i] = cell.length;
				}
			}
			return widths;
		},
		statLines[0].map(() => 0)
	);

	const table = statLines
		.map((row) =>
			row
				.map((cell, i) => {
					if (i === 0) {
						return cell.padEnd(columnWidths[i], ' ');
					} else {
						return cell.padStart(columnWidths[i], ' ');
					}
				})
				.join('\t')
		)
		.join('\n');
	return table;
}

/**
 * @class
 */
export class VitePluginSvelteStats {
	// package directory -> package name
	/** @type {import('./vite-plugin-svelte-cache.js').VitePluginSvelteCache} */
	#cache;
	/** @type {import('../types/vite-plugin-svelte-stats.d.ts').StatCollection[]} */
	#collections = [];

	/**
	 * @param {import('./vite-plugin-svelte-cache.js').VitePluginSvelteCache} cache
	 */
	constructor(cache) {
		this.#cache = cache;
	}

	/**
	 * @param {string} name
	 * @param {Partial<import('../types/vite-plugin-svelte-stats.d.ts').CollectionOptions>} [opts]
	 * @returns {import('../types/vite-plugin-svelte-stats.d.ts').StatCollection}
	 */
	startCollection(name, opts) {
		const options = {
			...defaultCollectionOptions,
			...opts
		};
		/** @type {import('../types/vite-plugin-svelte-stats.d.ts').Stat[]} */
		const stats = [];
		const collectionStart = performance.now();
		const _this = this;
		let hasLoggedProgress = false;
		/** @type {import('../types/vite-plugin-svelte-stats.d.ts').StatCollection} */
		const collection = {
			name,
			options,
			stats,
			collectionStart,
			finished: false,
			start(file) {
				if (collection.finished) {
					throw new Error('called after finish() has been used');
				}
				file = normalizePath(file);
				const start = performance.now();
				/** @type {import('../types/vite-plugin-svelte-stats.d.ts').Stat} */
				const stat = { file, start, end: start };
				return () => {
					const now = performance.now();
					stat.end = now;
					stats.push(stat);
					if (!hasLoggedProgress && options.logInProgress(collection, now)) {
						hasLoggedProgress = true;
						log.debug(`${name} in progress ...`, undefined, 'stats');
					}
				};
			},
			async finish() {
				await _this.#finish(collection);
			}
		};
		_this.#collections.push(collection);
		return collection;
	}

	async finishAll() {
		await Promise.all(this.#collections.map((c) => c.finish()));
	}

	/**
	 * @param {import('../types/vite-plugin-svelte-stats.d.ts').StatCollection} collection
	 */
	async #finish(collection) {
		try {
			collection.finished = true;
			const now = performance.now();
			collection.duration = now - collection.collectionStart;
			const logResult = collection.options.logResult(collection);
			if (logResult) {
				await this.#aggregateStatsResult(collection);
				log.debug(
					`${collection.name} done.\n${formatPackageStats(
						/** @type {import('../types/vite-plugin-svelte-stats.d.ts').PackageStats[]}*/ (
							collection.packageStats
						)
					)}`,
					undefined,
					'stats'
				);
			}
			// cut some ties to free it for garbage collection
			const index = this.#collections.indexOf(collection);
			this.#collections.splice(index, 1);
			collection.stats.length = 0;
			collection.stats = [];
			if (collection.packageStats) {
				collection.packageStats.length = 0;
				collection.packageStats = [];
			}
			collection.start = () => () => {};
			collection.finish = () => {};
		} catch (e) {
			// this should not happen, but stats taking also should not break the process
			log.debug.once(`failed to finish stats for ${collection.name}\n`, e, 'stats');
		}
	}

	/**
	 * @param {import('../types/vite-plugin-svelte-stats.d.ts').StatCollection} collection
	 */
	async #aggregateStatsResult(collection) {
		const stats = collection.stats;
		for (const stat of stats) {
			stat.pkg = (await this.#cache.getPackageInfo(stat.file)).name;
		}

		// group stats
		/** @type {Record<string, import('../types/vite-plugin-svelte-stats.d.ts').PackageStats>} */
		const grouped = {};
		stats.forEach((stat) => {
			const pkg = /** @type {string} */ (stat.pkg);
			let group = grouped[pkg];
			if (!group) {
				group = grouped[pkg] = {
					files: 0,
					duration: 0,
					pkg
				};
			}
			group.files += 1;
			group.duration += stat.end - stat.start;
		});

		const groups = Object.values(grouped);
		groups.sort((a, b) => b.duration - a.duration);
		collection.packageStats = groups;
	}
}
