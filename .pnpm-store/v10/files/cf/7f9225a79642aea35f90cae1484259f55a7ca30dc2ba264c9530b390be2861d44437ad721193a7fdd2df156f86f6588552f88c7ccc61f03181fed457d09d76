import fs from 'node:fs';
import { log } from './log.js';
import { knownSvelteConfigNames } from './load-svelte-config.js';
import path from 'node:path';

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @param {import('./vite-plugin-svelte-cache.js').VitePluginSvelteCache} cache
 * @param {import('../types/id.d.ts').IdParser} requestParser
 * @returns {void}
 */
export function setupWatchers(options, cache, requestParser) {
	const { server, configFile: svelteConfigFile } = options;
	if (!server) {
		return;
	}
	const { watcher, ws } = server;
	const { root, server: serverConfig } = server.config;
	/** @type {(filename: string) => void} */
	const emitChangeEventOnDependants = (filename) => {
		const dependants = cache.getDependants(filename);
		dependants.forEach((dependant) => {
			if (fs.existsSync(dependant)) {
				log.debug(
					`emitting virtual change event for "${dependant}" because depdendency "${filename}" changed`,
					undefined,
					'hmr'
				);
				watcher.emit('change', dependant);
			}
		});
	};
	/** @type {(filename: string) => void} */
	const removeUnlinkedFromCache = (filename) => {
		const svelteRequest = requestParser(filename, false);
		if (svelteRequest) {
			const removedFromCache = cache.remove(svelteRequest);
			if (removedFromCache) {
				log.debug(`cleared VitePluginSvelteCache for deleted file ${filename}`, undefined, 'hmr');
			}
		}
	};
	/** @type {(filename: string) => void} */
	const triggerViteRestart = (filename) => {
		if (serverConfig.middlewareMode) {
			// in middlewareMode we can't restart the server automatically
			// show the user an overlay instead
			const message =
				'Svelte config change detected, restart your dev process to apply the changes.';
			log.info(message, filename);
			ws.send({
				type: 'error',
				err: { message, stack: '', plugin: 'vite-plugin-svelte', id: filename }
			});
		} else {
			log.info(`svelte config changed: restarting vite server. - file: ${filename}`);
			server.restart();
		}
	};

	// collection of watcher listeners by event
	/** @type {Record<string, Function[]>} */
	const listenerCollection = {
		add: [],
		change: [emitChangeEventOnDependants],
		unlink: [removeUnlinkedFromCache, emitChangeEventOnDependants]
	};

	if (svelteConfigFile !== false) {
		// configFile false means we ignore the file and external process is responsible
		const possibleSvelteConfigs = knownSvelteConfigNames.map((cfg) => path.join(root, cfg));
		/** @type {(filename: string) => void} */
		const restartOnConfigAdd = (filename) => {
			if (possibleSvelteConfigs.includes(filename)) {
				triggerViteRestart(filename);
			}
		};

		/** @type {(filename: string) => void} */
		const restartOnConfigChange = (filename) => {
			if (filename === svelteConfigFile) {
				triggerViteRestart(filename);
			}
		};

		if (svelteConfigFile) {
			listenerCollection.change.push(restartOnConfigChange);
			listenerCollection.unlink.push(restartOnConfigChange);
		} else {
			listenerCollection.add.push(restartOnConfigAdd);
		}
	}

	Object.entries(listenerCollection).forEach(([evt, listeners]) => {
		if (listeners.length > 0) {
			watcher.on(evt, (filename) => listeners.forEach((listener) => listener(filename)));
		}
	});
}

/**
 * taken from vite utils
 * @param {import('vite').FSWatcher} watcher
 * @param {string | null} file
 * @param {string} root
 * @returns {void}
 */
export function ensureWatchedFile(watcher, file, root) {
	if (
		file &&
		// only need to watch if out of root
		!file.startsWith(root + '/') &&
		// some rollup plugins use null bytes for private resolved Ids
		!file.includes('\0') &&
		fs.existsSync(file)
	) {
		// resolve file to normalized system path
		watcher.add(path.resolve(file));
	}
}
