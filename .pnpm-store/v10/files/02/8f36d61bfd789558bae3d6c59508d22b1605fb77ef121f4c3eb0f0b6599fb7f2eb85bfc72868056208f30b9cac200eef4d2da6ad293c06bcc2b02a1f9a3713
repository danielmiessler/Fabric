/* eslint-disable no-console */
import { cyan, red, yellow } from 'kleur/colors';
import debug from 'debug';
import { VERSION } from 'svelte/compiler';

/** @type {import('../types/log.d.ts').LogLevel[]} */
const levels = ['debug', 'info', 'warn', 'error', 'silent'];
const prefix = 'vite-plugin-svelte';
/** @type {Record<import('../types/log.d.ts').LogLevel, any>} */
const loggers = {
	debug: {
		log: debug(`${prefix}`),
		enabled: false,
		isDebug: true
	},
	info: {
		color: cyan,
		log: console.log,
		enabled: true
	},
	warn: {
		color: yellow,
		log: console.warn,
		enabled: true
	},
	error: {
		color: red,
		log: console.error,
		enabled: true
	},
	silent: {
		enabled: false
	}
};

/** @type {import('../types/log.d.ts').LogLevel} */
let _level = 'info';
/**
 * @param {import('../types/log.d.ts').LogLevel} level
 * @returns {void}
 */
function setLevel(level) {
	if (level === _level) {
		return;
	}
	const levelIndex = levels.indexOf(level);
	if (levelIndex > -1) {
		_level = level;
		for (let i = 0; i < levels.length; i++) {
			loggers[levels[i]].enabled = i >= levelIndex;
		}
	} else {
		_log(loggers.error, `invalid log level: ${level} `);
	}
}

/**
 * @param {any} logger
 * @param {string} message
 * @param {any} [payload]
 * @param {string} [namespace]
 * @returns
 */
function _log(logger, message, payload, namespace) {
	if (!logger.enabled) {
		return;
	}
	if (logger.isDebug) {
		let log = logger.log;
		if (namespace) {
			if (!isDebugNamespaceEnabled(namespace)) {
				return;
			}
			log = logger.log.extend(namespace);
		}
		payload !== undefined ? log(message, payload) : log(message);
	} else {
		logger.log(
			logger.color(
				`${new Date().toLocaleTimeString()} [${prefix}${
					namespace ? `:${namespace}` : ''
				}] ${message}`
			)
		);
		if (payload) {
			logger.log(payload);
		}
	}
}

/**
 * @param {import('../types/log.d.ts').LogLevel} level
 * @returns {import('../types/log.d.ts').LogFn}
 */
function createLogger(level) {
	const logger = loggers[level];
	const logFn = /** @type {import('../types/log.d.ts').LogFn} */ (_log.bind(null, logger));
	/** @type {Set<string>} */
	const logged = new Set();
	/** @type {import('../types/log.d.ts').SimpleLogFn} */
	const once = function (message, payload, namespace) {
		if (!logger.enabled || logged.has(message)) {
			return;
		}
		logged.add(message);
		logFn.apply(null, [message, payload, namespace]);
	};
	Object.defineProperty(logFn, 'enabled', {
		get() {
			return logger.enabled;
		}
	});
	Object.defineProperty(logFn, 'once', {
		get() {
			return once;
		}
	});
	return logFn;
}

export const log = {
	debug: createLogger('debug'),
	info: createLogger('info'),
	warn: createLogger('warn'),
	error: createLogger('error'),
	setLevel
};

/**
 * @param {import('../types/id.d.ts').SvelteRequest | import('../types/id.d.ts').SvelteModuleRequest} svelteRequest
 * @param {import('svelte/types/compiler/interfaces').Warning[]} warnings
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 */
export function logCompilerWarnings(svelteRequest, warnings, options) {
	const { emitCss, onwarn, isBuild } = options;
	const sendViaWS = !isBuild && options.experimental?.sendWarningsToBrowser;
	let warn = isBuild ? warnBuild : warnDev;
	/** @type {import('svelte/types/compiler/interfaces').Warning[]} */
	const handledByDefaultWarn = [];
	const notIgnored = warnings?.filter((w) => !ignoreCompilerWarning(w, isBuild, emitCss));
	const extra = buildExtraWarnings(warnings, isBuild);
	const allWarnings = [...notIgnored, ...extra];
	if (sendViaWS) {
		const _warn = warn;
		/** @type {(w: import('svelte/types/compiler/interfaces').Warning) => void} */
		warn = (w) => {
			handledByDefaultWarn.push(w);
			_warn(w);
		};
	}
	allWarnings.forEach((warning) => {
		if (onwarn) {
			onwarn(warning, warn);
		} else {
			warn(warning);
		}
	});
	if (sendViaWS) {
		/** @type {import('../types/log.d.ts').SvelteWarningsMessage} */
		const message = {
			id: svelteRequest.id,
			filename: svelteRequest.filename,
			normalizedFilename: svelteRequest.normalizedFilename,
			timestamp: svelteRequest.timestamp,
			warnings: handledByDefaultWarn, // allWarnings filtered by warnings where onwarn did not call the default handler
			allWarnings, // includes warnings filtered by onwarn and our extra vite plugin svelte warnings
			rawWarnings: warnings // raw compiler output
		};
		log.debug(`sending svelte:warnings message for ${svelteRequest.normalizedFilename}`);
		options.server?.ws?.send('svelte:warnings', message);
	}
}

/**
 * @param {import('svelte/types/compiler/interfaces').Warning} warning
 * @param {boolean} isBuild
 * @param {boolean} [emitCss]
 * @returns {boolean}
 */
function ignoreCompilerWarning(warning, isBuild, emitCss) {
	return (
		(!emitCss && warning.code === 'css-unused-selector') || // same as rollup-plugin-svelte
		(!isBuild && isNoScopableElementWarning(warning))
	);
}

/**
 *
 * @param {import('svelte/types/compiler/interfaces').Warning} warning
 * @returns {boolean}
 */
function isNoScopableElementWarning(warning) {
	// see https://github.com/sveltejs/vite-plugin-svelte/issues/153
	return warning.code === 'css-unused-selector' && warning.message.includes('"*"');
}

/**
 *
 * @param {import('svelte/types/compiler/interfaces').Warning[]} warnings
 * @param {boolean} isBuild
 * @returns {import('svelte/types/compiler/interfaces').Warning[]}
 */
function buildExtraWarnings(warnings, isBuild) {
	const extraWarnings = [];
	if (!isBuild) {
		const noScopableElementWarnings = warnings.filter((w) => isNoScopableElementWarning(w));
		if (noScopableElementWarnings.length > 0) {
			// in case there are multiple, use last one as that is the one caused by our *{} rule
			const noScopableElementWarning =
				noScopableElementWarnings[noScopableElementWarnings.length - 1];
			extraWarnings.push({
				...noScopableElementWarning,
				code: 'vite-plugin-svelte-css-no-scopable-elements',
				message:
					"No scopable elements found in template. If you're using global styles in the style tag, you should move it into an external stylesheet file and import it in JS. See https://github.com/sveltejs/vite-plugin-svelte/blob/main/docs/faq.md#where-should-i-put-my-global-styles."
			});
		}
	}
	return extraWarnings;
}

/**
 * @param {import('svelte/types/compiler/interfaces').Warning} w
 */
function warnDev(w) {
	log.info.enabled && log.info(buildExtendedLogMessage(w));
}

/**
 * @param {import('svelte/types/compiler/interfaces').Warning} w
 */
function warnBuild(w) {
	log.warn.enabled && log.warn(buildExtendedLogMessage(w), w.frame);
}

/**
 * @param {import('svelte/types/compiler/interfaces').Warning} w
 */
export function buildExtendedLogMessage(w) {
	const parts = [];
	if (w.filename) {
		parts.push(w.filename);
	}
	if (w.start) {
		parts.push(':', w.start.line, ':', w.start.column);
	}
	if (w.message) {
		if (parts.length > 0) {
			parts.push(' ');
		}
		parts.push(w.message);
	}
	return parts.join('');
}

/**
 * @param {string} namespace
 * @returns {boolean}
 */
export function isDebugNamespaceEnabled(namespace) {
	return debug.enabled(`${prefix}:${namespace}`);
}

export function logSvelte5Warning() {
	log.warn(
		`You are using Svelte ${VERSION} with vite-plugin-svelte@3. Active Svelte 5 support has moved to vite-plugin-svelte@4.
			To receive bug fixes and new features update your devDependencies to "@sveltejs/vite-plugin-svelte": "^4.0.0-next.6" and install.
			For framework integrations that depend on it, you might have to add an override:
			"overrides": {"@sveltejs/vite-plugin-svelte": "^4.0.0-next.6"}`.replace(/\n\s*/gm, '\n\t')
	);
}
