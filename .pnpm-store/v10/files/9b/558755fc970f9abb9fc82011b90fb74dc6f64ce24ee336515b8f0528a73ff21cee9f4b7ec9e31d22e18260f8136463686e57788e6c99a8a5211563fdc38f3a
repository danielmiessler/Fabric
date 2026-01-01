/* eslint-disable no-unused-vars */
import { normalizePath } from 'vite';
import { isDebugNamespaceEnabled, log } from './log.js';
import { loadSvelteConfig } from './load-svelte-config.js';
import {
	FAQ_LINK_MISSING_EXPORTS_CONDITION,
	SVELTE_EXPORT_CONDITIONS,
	SVELTE_HMR_IMPORTS,
	SVELTE_IMPORTS,
	SVELTE_RESOLVE_MAIN_FIELDS,
	VITE_RESOLVE_MAIN_FIELDS
} from './constants.js';

import path from 'node:path';
import { esbuildSveltePlugin, facadeEsbuildSveltePluginName } from './esbuild.js';
import { addExtraPreprocessors } from './preprocess.js';
import deepmerge from 'deepmerge';
import {
	crawlFrameworkPkgs,
	isDepExcluded,
	isDepExternaled,
	isDepIncluded,
	isDepNoExternaled
} from 'vitefu';

import { isCommonDepWithoutSvelteField } from './dependencies.js';
import { VitePluginSvelteStats } from './vite-plugin-svelte-stats.js';
import { VitePluginSvelteCache } from './vite-plugin-svelte-cache.js';
import { isSvelte5, isSvelte5WithHMRSupport } from './svelte-version.js';

const allowedPluginOptions = new Set([
	'include',
	'exclude',
	'emitCss',
	'hot',
	'ignorePluginPreprocessors',
	'disableDependencyReinclusion',
	'prebundleSvelteLibraries',
	'inspector',
	'dynamicCompileOptions',
	'experimental'
]);

const knownRootOptions = new Set(['extensions', 'compilerOptions', 'preprocess', 'onwarn']);

const allowedInlineOptions = new Set(['configFile', ...allowedPluginOptions, ...knownRootOptions]);

/**
 * @param {Partial<import('../public.d.ts').Options>} [inlineOptions]
 */
export function validateInlineOptions(inlineOptions) {
	const invalidKeys = Object.keys(inlineOptions || {}).filter(
		(key) => !allowedInlineOptions.has(key)
	);
	if (invalidKeys.length) {
		log.warn(`invalid plugin options "${invalidKeys.join(', ')}" in inline config`, inlineOptions);
	}
}

/**
 * @param {Partial<import('../public.d.ts').SvelteConfig>} [config]
 * @returns {Partial<import('../public.d.ts').Options> | undefined}
 */
function convertPluginOptions(config) {
	if (!config) {
		return;
	}
	const invalidRootOptions = Object.keys(config).filter((key) => allowedPluginOptions.has(key));
	if (invalidRootOptions.length > 0) {
		throw new Error(
			`Invalid options in svelte config. Move the following options into 'vitePlugin:{...}': ${invalidRootOptions.join(
				', '
			)}`
		);
	}
	if (!config.vitePlugin) {
		return config;
	}
	const pluginOptions = config.vitePlugin;
	const pluginOptionKeys = Object.keys(pluginOptions);

	const rootOptionsInPluginOptions = pluginOptionKeys.filter((key) => knownRootOptions.has(key));
	if (rootOptionsInPluginOptions.length > 0) {
		throw new Error(
			`Invalid options in svelte config under vitePlugin:{...}', move them to the config root : ${rootOptionsInPluginOptions.join(
				', '
			)}`
		);
	}
	const duplicateOptions = pluginOptionKeys.filter((key) =>
		Object.prototype.hasOwnProperty.call(config, key)
	);
	if (duplicateOptions.length > 0) {
		throw new Error(
			`Invalid duplicate options in svelte config under vitePlugin:{...}', they are defined in root too and must only exist once: ${duplicateOptions.join(
				', '
			)}`
		);
	}
	const unknownPluginOptions = pluginOptionKeys.filter((key) => !allowedPluginOptions.has(key));
	if (unknownPluginOptions.length > 0) {
		log.warn(
			`ignoring unknown plugin options in svelte config under vitePlugin:{...}: ${unknownPluginOptions.join(
				', '
			)}`
		);
		unknownPluginOptions.forEach((unkownOption) => {
			// @ts-ignore
			delete pluginOptions[unkownOption];
		});
	}
	/** @type {import('../public.d.ts').Options} */
	const result = {
		...config,
		...pluginOptions
	};
	// @ts-expect-error it exists
	delete result.vitePlugin;

	return result;
}

/**
 * used in config phase, merges the default options, svelte config, and inline options
 * @param {Partial<import('../public.d.ts').Options> | undefined} inlineOptions
 * @param {import('vite').UserConfig} viteUserConfig
 * @param {import('vite').ConfigEnv} viteEnv
 * @returns {Promise<import('../types/options.d.ts').PreResolvedOptions>}
 */
export async function preResolveOptions(inlineOptions, viteUserConfig, viteEnv) {
	if (!inlineOptions) {
		inlineOptions = {};
	}
	/** @type {import('vite').UserConfig} */
	const viteConfigWithResolvedRoot = {
		...viteUserConfig,
		root: resolveViteRoot(viteUserConfig)
	};
	const isBuild = viteEnv.command === 'build';
	/** @type {Partial<import('../types/options.d.ts').PreResolvedOptions>} */
	const defaultOptions = {
		extensions: ['.svelte'],
		emitCss: true,
		prebundleSvelteLibraries: !isBuild
	};
	const svelteConfig = convertPluginOptions(
		await loadSvelteConfig(viteConfigWithResolvedRoot, inlineOptions)
	);
	/** @type {Partial<import('../types/options.d.ts').PreResolvedOptions>} */
	const extraOptions = {
		root: viteConfigWithResolvedRoot.root,
		isBuild,
		isServe: viteEnv.command === 'serve',
		isDebug: process.env.DEBUG != null
	};

	const merged = /** @type {import('../types/options.d.ts').PreResolvedOptions} */ (
		mergeConfigs(defaultOptions, svelteConfig, inlineOptions, extraOptions)
	);
	// configFile of svelteConfig contains the absolute path it was loaded from,
	// prefer it over the possibly relative inline path
	if (svelteConfig?.configFile) {
		merged.configFile = svelteConfig.configFile;
	}
	return merged;
}

/**
 * @template T
 * @param  {(Partial<T> | undefined)[]} configs
 * @returns T
 */
function mergeConfigs(...configs) {
	/** @type {Partial<T>} */
	let result = {};
	for (const config of configs.filter((x) => x != null)) {
		result = deepmerge(result, /** @type {Partial<T>} */ (config), {
			// replace arrays
			arrayMerge: (target, source) => source ?? target
		});
	}
	return /** @type {T} */ result;
}

/**
 * used in configResolved phase, merges a contextual default config, pre-resolved options, and some preprocessors. also validates the final config.
 *
 * @param {import('../types/options.d.ts').PreResolvedOptions} preResolveOptions
 * @param {import('vite').ResolvedConfig} viteConfig
 * @param {VitePluginSvelteCache} cache
 * @returns {import('../types/options.d.ts').ResolvedOptions}
 */
export function resolveOptions(preResolveOptions, viteConfig, cache) {
	const css = preResolveOptions.emitCss ? 'external' : 'injected';
	/** @type {Partial<import('../public.d.ts').Options>} */
	const defaultOptions = {
		compilerOptions: {
			css,
			dev: !viteConfig.isProduction
		}
	};
	const hot =
		!viteConfig.isProduction && !preResolveOptions.isBuild && viteConfig.server.hmr !== false;
	if (isSvelte5) {
		if (isSvelte5WithHMRSupport) {
			// @ts-expect-error svelte4 does not have hmr option
			defaultOptions.compilerOptions.hmr = hot;
		}
	} else {
		defaultOptions.hot = !hot
			? false
			: {
					injectCss: css === 'injected',
					partialAccept: !!viteConfig.experimental?.hmrPartialAccept
				};
	}
	/** @type {Partial<import('../types/options.d.ts').ResolvedOptions>} */
	const extraOptions = {
		root: viteConfig.root,
		isProduction: viteConfig.isProduction
	};
	const merged = /** @type {import('../types/options.d.ts').ResolvedOptions}*/ (
		mergeConfigs(defaultOptions, preResolveOptions, extraOptions)
	);

	removeIgnoredOptions(merged);
	handleDeprecatedOptions(merged);
	addExtraPreprocessors(merged, viteConfig);
	enforceOptionsForHmr(merged, viteConfig);
	enforceOptionsForProduction(merged);
	// mergeConfigs would mangle functions on the stats class, so do this afterwards
	if (log.debug.enabled && isDebugNamespaceEnabled('stats')) {
		merged.stats = new VitePluginSvelteStats(cache);
	}
	return merged;
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @param {import('vite').ResolvedConfig} viteConfig
 */
function enforceOptionsForHmr(options, viteConfig) {
	if (options.hot && viteConfig.server.hmr === false) {
		log.warn(
			'vite config server.hmr is false but hot is true. Forcing hot to false as it would not work.'
		);
		options.hot = false;
	}
	if (isSvelte5) {
		if (options.hot && isSvelte5WithHMRSupport) {
			log.warn(
				'svelte 5 has hmr integrated in core. Please remove the hot option and use compilerOptions.hmr instead'
			);
			delete options.hot;
			// @ts-expect-error hmr option doesn't exist in svelte4
			options.compilerOptions.hmr = true;
		}
	} else {
		if (options.hot) {
			if (!options.compilerOptions.dev) {
				log.warn('hmr is enabled but compilerOptions.dev is false, forcing it to true');
				options.compilerOptions.dev = true;
			}
			if (options.emitCss) {
				if (options.hot !== true && options.hot.injectCss) {
					log.warn('hmr and emitCss are enabled but hot.injectCss is true, forcing it to false');
					options.hot.injectCss = false;
				}
				const css = options.compilerOptions.css;
				if (css === true || css === 'injected') {
					const forcedCss = 'external';
					log.warn(
						`hmr and emitCss are enabled but compilerOptions.css is ${css}, forcing it to ${forcedCss}`
					);
					options.compilerOptions.css = forcedCss;
				}
			} else {
				if (options.hot === true || !options.hot.injectCss) {
					log.warn(
						'hmr with emitCss disabled requires option hot.injectCss to be enabled, forcing it to true'
					);
					if (options.hot === true) {
						options.hot = { injectCss: true };
					} else {
						options.hot.injectCss = true;
					}
				}
				const css = options.compilerOptions.css;
				if (!(css === true || css === 'injected')) {
					const forcedCss = 'injected';
					log.warn(
						`hmr with emitCss disabled requires compilerOptions.css to be enabled, forcing it to ${forcedCss}`
					);
					options.compilerOptions.css = forcedCss;
				}
			}
		}
	}
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 */
function enforceOptionsForProduction(options) {
	if (options.isProduction) {
		if (options.hot) {
			log.warn('options.hot is enabled but does not work on production build, forcing it to false');
			options.hot = false;
		}
		if (options.compilerOptions.dev) {
			log.warn(
				'you are building for production but compilerOptions.dev is true, forcing it to false'
			);
			options.compilerOptions.dev = false;
		}
	}
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 */
function removeIgnoredOptions(options) {
	const ignoredCompilerOptions = ['generate', 'format', 'filename'];
	if (options.hot && options.emitCss) {
		ignoredCompilerOptions.push('cssHash');
	}
	const passedCompilerOptions = Object.keys(options.compilerOptions || {});
	const passedIgnored = passedCompilerOptions.filter((o) => ignoredCompilerOptions.includes(o));
	if (passedIgnored.length) {
		log.warn(
			`The following Svelte compilerOptions are controlled by vite-plugin-svelte and essential to its functionality. User-specified values are ignored. Please remove them from your configuration: ${passedIgnored.join(
				', '
			)}`
		);
		passedIgnored.forEach((ignored) => {
			// @ts-expect-error string access
			delete options.compilerOptions[ignored];
		});
	}
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 */
function handleDeprecatedOptions(options) {
	const experimental = /** @type {Record<string, any>} */ (options.experimental);
	if (experimental) {
		for (const promoted of ['prebundleSvelteLibraries', 'inspector', 'dynamicCompileOptions']) {
			if (experimental[promoted]) {
				//@ts-expect-error untyped assign
				options[promoted] = experimental[promoted];
				delete experimental[promoted];
				log.warn(
					`Option "experimental.${promoted}" is no longer experimental and has moved to "${promoted}". Please update your Svelte or Vite config.`
				);
			}
		}
		if (experimental.generateMissingPreprocessorSourcemaps) {
			log.warn('experimental.generateMissingPreprocessorSourcemaps has been removed.');
		}
	}
}

/**
 * vite passes unresolved `root`option to config hook but we need the resolved value, so do it here
 *
 * @see https://github.com/sveltejs/vite-plugin-svelte/issues/113
 * @see https://github.com/vitejs/vite/blob/43c957de8a99bb326afd732c962f42127b0a4d1e/packages/vite/src/node/config.ts#L293
 *
 * @param {import('vite').UserConfig} viteConfig
 * @returns {string | undefined}
 */
function resolveViteRoot(viteConfig) {
	return normalizePath(viteConfig.root ? path.resolve(viteConfig.root) : process.cwd());
}

/**
 * @param {import('../types/options.d.ts').PreResolvedOptions} options
 * @param {import('vite').UserConfig} config
 * @returns {Promise<Partial<import('vite').UserConfig>>}
 */
export async function buildExtraViteConfig(options, config) {
	// make sure we only readd vite default mainFields when no other plugin has changed the config already
	// see https://github.com/sveltejs/vite-plugin-svelte/issues/581
	if (!config.resolve) {
		config.resolve = {};
	}
	config.resolve.mainFields = [
		...SVELTE_RESOLVE_MAIN_FIELDS,
		...(config.resolve.mainFields ?? VITE_RESOLVE_MAIN_FIELDS)
	];

	/** @type {Partial<import('vite').UserConfig>} */
	const extraViteConfig = {
		resolve: {
			dedupe: [...SVELTE_IMPORTS, ...SVELTE_HMR_IMPORTS],
			conditions: [...SVELTE_EXPORT_CONDITIONS]
		}
		// this option is still awaiting a PR in vite to be supported
		// see https://github.com/sveltejs/vite-plugin-svelte/issues/60
		// @ts-ignore
		// knownJsSrcExtensions: options.extensions
	};

	const extraSvelteConfig = buildExtraConfigForSvelte(config);
	const extraDepsConfig = await buildExtraConfigForDependencies(options, config);
	// merge extra svelte and deps config, but make sure dep values are not contradicting svelte
	extraViteConfig.optimizeDeps = {
		include: [
			...extraSvelteConfig.optimizeDeps.include,
			...extraDepsConfig.optimizeDeps.include.filter(
				(dep) => !isDepExcluded(dep, extraSvelteConfig.optimizeDeps.exclude)
			)
		],
		exclude: [
			...extraSvelteConfig.optimizeDeps.exclude,
			...extraDepsConfig.optimizeDeps.exclude.filter(
				(dep) => !isDepIncluded(dep, extraSvelteConfig.optimizeDeps.include)
			)
		]
	};

	extraViteConfig.ssr = {
		external: [
			...extraSvelteConfig.ssr.external,
			...extraDepsConfig.ssr.external.filter(
				(dep) => !isDepNoExternaled(dep, extraSvelteConfig.ssr.noExternal)
			)
		],
		noExternal: [
			...extraSvelteConfig.ssr.noExternal,
			...extraDepsConfig.ssr.noExternal.filter(
				(dep) => !isDepExternaled(dep, extraSvelteConfig.ssr.external)
			)
		]
	};

	// handle prebundling for svelte files
	if (options.prebundleSvelteLibraries) {
		extraViteConfig.optimizeDeps = {
			...extraViteConfig.optimizeDeps,
			// Experimental Vite API to allow these extensions to be scanned and prebundled
			// @ts-ignore
			extensions: options.extensions ?? ['.svelte'],
			// Add esbuild plugin to prebundle Svelte files.
			// Currently a placeholder as more information is needed after Vite config is resolved,
			// the real Svelte plugin is added in `patchResolvedViteConfig()`
			esbuildOptions: {
				plugins: [{ name: facadeEsbuildSveltePluginName, setup: () => {} }]
			}
		};
	}

	// enable hmrPartialAccept if not explicitly disabled
	if (
		(options.hot == null ||
			options.hot === true ||
			(options.hot && options.hot.partialAccept !== false)) && // deviate from svelte-hmr, default to true
		config.experimental?.hmrPartialAccept !== false
	) {
		log.debug('enabling "experimental.hmrPartialAccept" in vite config', undefined, 'config');
		extraViteConfig.experimental = { hmrPartialAccept: true };
	}
	validateViteConfig(extraViteConfig, config, options);
	return extraViteConfig;
}

/**
 * @param {Partial<import('vite').UserConfig>} extraViteConfig
 * @param {import('vite').UserConfig} config
 * @param {import('../types/options.d.ts').PreResolvedOptions} options
 */
function validateViteConfig(extraViteConfig, config, options) {
	const { prebundleSvelteLibraries, isBuild } = options;
	if (prebundleSvelteLibraries) {
		/** @type {(option: 'dev' | 'build' | boolean)=> boolean} */
		const isEnabled = (option) => option !== true && option !== (isBuild ? 'build' : 'dev');
		/** @type {(name: string, value: 'dev' | 'build' | boolean, recommendation: string)=> void} */
		const logWarning = (name, value, recommendation) =>
			log.warn.once(
				`Incompatible options: \`prebundleSvelteLibraries: true\` and vite \`${name}: ${JSON.stringify(
					value
				)}\` ${isBuild ? 'during build.' : '.'} ${recommendation}`
			);
		const viteOptimizeDepsDisabled = config.optimizeDeps?.disabled ?? 'build'; // fall back to vite default
		const isOptimizeDepsEnabled = isEnabled(viteOptimizeDepsDisabled);
		if (!isBuild && !isOptimizeDepsEnabled) {
			logWarning(
				'optimizeDeps.disabled',
				viteOptimizeDepsDisabled,
				'Forcing `optimizeDeps.disabled: "build"`. Disable prebundleSvelteLibraries or update your vite config to enable optimizeDeps during dev.'
			);
			if (!extraViteConfig.optimizeDeps) {
				extraViteConfig.optimizeDeps = {};
			}
			extraViteConfig.optimizeDeps.disabled = 'build';
		} else if (isBuild && isOptimizeDepsEnabled) {
			logWarning(
				'optimizeDeps.disabled',
				viteOptimizeDepsDisabled,
				'Disable optimizeDeps or prebundleSvelteLibraries for build if you experience errors.'
			);
		}
	}
}

/**
 * @param {import('../types/options.d.ts').PreResolvedOptions} options
 * @param {import('vite').UserConfig} config
 * @returns {Promise<import('vitefu').CrawlFrameworkPkgsResult>}
 */
async function buildExtraConfigForDependencies(options, config) {
	// extra handling for svelte dependencies in the project
	const packagesWithoutSvelteExportsCondition = new Set();
	const depsConfig = await crawlFrameworkPkgs({
		root: options.root,
		isBuild: options.isBuild,
		viteUserConfig: config,
		isFrameworkPkgByJson(pkgJson) {
			let hasSvelteCondition = false;
			if (typeof pkgJson.exports === 'object') {
				// use replacer as a simple way to iterate over nested keys
				JSON.stringify(pkgJson.exports, (key, value) => {
					if (SVELTE_EXPORT_CONDITIONS.includes(key)) {
						hasSvelteCondition = true;
					}
					return value;
				});
			}
			const hasSvelteField = !!pkgJson.svelte;
			if (hasSvelteField && !hasSvelteCondition) {
				packagesWithoutSvelteExportsCondition.add(`${pkgJson.name}@${pkgJson.version}`);
			}
			return hasSvelteCondition || hasSvelteField;
		},
		isSemiFrameworkPkgByJson(pkgJson) {
			return !!pkgJson.dependencies?.svelte || !!pkgJson.peerDependencies?.svelte;
		},
		isFrameworkPkgByName(pkgName) {
			const isNotSveltePackage = isCommonDepWithoutSvelteField(pkgName);
			if (isNotSveltePackage) {
				return false;
			} else {
				return undefined;
			}
		}
	});
	if (
		!options.experimental?.disableSvelteResolveWarnings &&
		packagesWithoutSvelteExportsCondition?.size > 0
	) {
		log.warn(
			`WARNING: The following packages have a svelte field in their package.json but no exports condition for svelte.\n\n${[
				...packagesWithoutSvelteExportsCondition
			].join('\n')}\n\nPlease see ${FAQ_LINK_MISSING_EXPORTS_CONDITION} for details.`
		);
	}
	log.debug('extra config for dependencies generated by vitefu', depsConfig, 'config');

	if (options.prebundleSvelteLibraries) {
		// prebundling enabled, so we don't need extra dependency excludes
		depsConfig.optimizeDeps.exclude = [];
		// but keep dependency reinclusions of explicit user excludes
		const userExclude = config.optimizeDeps?.exclude;
		depsConfig.optimizeDeps.include = !userExclude
			? []
			: depsConfig.optimizeDeps.include.filter((dep) => {
					// reincludes look like this: foo > bar > baz
					// in case foo or bar are excluded, we have to retain the reinclude even with prebundling
					return (
						dep.includes('>') &&
						dep
							.split('>')
							.slice(0, -1)
							.some((d) => isDepExcluded(d.trim(), userExclude))
					);
				});
	}
	if (options.disableDependencyReinclusion === true) {
		depsConfig.optimizeDeps.include = depsConfig.optimizeDeps.include.filter(
			(dep) => !dep.includes('>')
		);
	} else if (Array.isArray(options.disableDependencyReinclusion)) {
		const disabledDeps = options.disableDependencyReinclusion;
		depsConfig.optimizeDeps.include = depsConfig.optimizeDeps.include.filter((dep) => {
			if (!dep.includes('>')) return true;
			const trimDep = dep.replace(/\s+/g, '');
			return disabledDeps.some((disabled) => trimDep.includes(`${disabled}>`));
		});
	}

	log.debug('post-processed extra config for dependencies', depsConfig, 'config');

	return depsConfig;
}

/**
 * @param {import('vite').UserConfig} config
 * @returns {import('vite').UserConfig & { optimizeDeps: { include: string[], exclude:string[] }, ssr: { noExternal:(string|RegExp)[], external: string[] } } }
 */
function buildExtraConfigForSvelte(config) {
	// include svelte imports for optimization unless explicitly excluded
	/** @type {string[]} */
	const include = [];
	const exclude = ['svelte-hmr'];
	if (!isDepExcluded('svelte', config.optimizeDeps?.exclude ?? [])) {
		const svelteImportsToInclude = SVELTE_IMPORTS.filter((x) => x !== 'svelte/ssr'); // not used on clientside
		log.debug(
			`adding bare svelte packages to optimizeDeps.include: ${svelteImportsToInclude.join(', ')} `,
			undefined,
			'config'
		);
		include.push(...svelteImportsToInclude);
	} else {
		log.debug(
			'"svelte" is excluded in optimizeDeps.exclude, skipped adding it to include.',
			undefined,
			'config'
		);
	}
	/** @type {(string | RegExp)[]} */
	const noExternal = [];
	/** @type {string[]} */
	const external = [];
	// add svelte to ssr.noExternal unless it is present in ssr.external
	// so we can resolve it with svelte/ssr
	if (!isDepExternaled('svelte', config.ssr?.external ?? [])) {
		noExternal.push('svelte', /^svelte\//);
	}
	return { optimizeDeps: { include, exclude }, ssr: { noExternal, external } };
}

/**
 * @param {import('vite').ResolvedConfig} viteConfig
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 */
export function patchResolvedViteConfig(viteConfig, options) {
	if (options.preprocess) {
		for (const preprocessor of arraify(options.preprocess)) {
			if (preprocessor.style && '__resolvedConfig' in preprocessor.style) {
				preprocessor.style.__resolvedConfig = viteConfig;
			}
		}
	}

	// replace facade esbuild plugin with a real one
	const facadeEsbuildSveltePlugin = viteConfig.optimizeDeps.esbuildOptions?.plugins?.find(
		(plugin) => plugin.name === facadeEsbuildSveltePluginName
	);
	if (facadeEsbuildSveltePlugin) {
		Object.assign(facadeEsbuildSveltePlugin, esbuildSveltePlugin(options));
	}
}

/**
 * @template T
 * @param {T | T[]} value
 * @returns {T[]}
 */
function arraify(value) {
	return Array.isArray(value) ? value : [value];
}
