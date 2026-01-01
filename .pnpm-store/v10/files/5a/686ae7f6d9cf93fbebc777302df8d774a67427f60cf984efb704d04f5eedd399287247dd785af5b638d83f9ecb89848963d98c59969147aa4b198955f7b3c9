import * as svelte from 'svelte/compiler';
// @ts-ignore
import { createMakeHot } from 'svelte-hmr';
import { safeBase64Hash } from './hash.js';
import { log } from './log.js';

import {
	checkPreprocessDependencies,
	createInjectScopeEverythingRulePreprocessorGroup
} from './preprocess.js';
import { mapToRelative } from './sourcemaps.js';
import { enhanceCompileError } from './error.js';
import { isSvelte5 } from './svelte-version.js';

// TODO this is a patched version of https://github.com/sveltejs/vite-plugin-svelte/pull/796/files#diff-3bce0b33034aad4b35ca094893671f7e7ddf4d27254ae7b9b0f912027a001b15R10
// which is closer to the other regexes in at least not falling into commented script
// but ideally would be shared exactly with svelte and other tools that use it
const scriptLangRE =
	/<!--[^]*?-->|<script (?:[^>]*|(?:[^=>'"/]+=(?:"[^"]*"|'[^']*'|[^>\s]+)\s+)*)lang=["']?([^"' >]+)["']?[^>]*>/g;

/**
 * @param {Function} [makeHot]
 * @returns {import('../types/compile.d.ts').CompileSvelte}
 */
export const _createCompileSvelte = (makeHot) => {
	/** @type {import('../types/vite-plugin-svelte-stats.d.ts').StatCollection | undefined} */
	let stats;
	const devStylePreprocessor = createInjectScopeEverythingRulePreprocessorGroup();
	/** @type {import('../types/compile.d.ts').CompileSvelte} */
	return async function compileSvelte(svelteRequest, code, options) {
		const { filename, normalizedFilename, cssId, ssr, raw } = svelteRequest;
		const { emitCss = true } = options;
		/** @type {string[]} */
		const dependencies = [];
		/** @type {import('svelte/types/compiler/interfaces').Warning[]} */
		const warnings = [];

		if (options.stats) {
			if (options.isBuild) {
				if (!stats) {
					// build is either completely ssr or csr, create stats collector on first compile
					// it is then finished in the buildEnd hook.
					stats = options.stats.startCollection(`${ssr ? 'ssr' : 'dom'} compile`, {
						logInProgress: () => false
					});
				}
			} else {
				// dev time ssr, it's a ssr request and there are no stats, assume new page load and start collecting
				if (ssr && !stats) {
					stats = options.stats.startCollection('ssr compile');
				}
				// stats are being collected but this isn't an ssr request, assume page loaded and stop collecting
				if (!ssr && stats) {
					stats.finish();
					stats = undefined;
				}
				// TODO find a way to trace dom compile during dev
				// problem: we need to call finish at some point but have no way to tell if page load finished
				// also they for hmr updates too
			}
		}
		/** @type {import('svelte/compiler').CompileOptions} */
		const compileOptions = {
			...options.compilerOptions,
			filename,
			// @ts-expect-error svelte5 uses server/client, svelte4 uses ssr/dom
			generate: isSvelte5 ? (ssr ? 'server' : 'client') : ssr ? 'ssr' : 'dom'
		};

		if (options.hot && options.emitCss) {
			const hash = `s-${safeBase64Hash(normalizedFilename)}`;
			compileOptions.cssHash = () => hash;
		}

		let preprocessed;
		let preprocessors = options.preprocess;
		if (!options.isBuild && options.emitCss && options.hot) {
			// inject preprocessor that ensures css hmr works better
			if (!Array.isArray(preprocessors)) {
				preprocessors = preprocessors
					? [preprocessors, devStylePreprocessor]
					: [devStylePreprocessor];
			} else {
				preprocessors = preprocessors.concat(devStylePreprocessor);
			}
		}
		if (preprocessors) {
			try {
				preprocessed = await svelte.preprocess(code, preprocessors, { filename }); // full filename here so postcss works
			} catch (e) {
				e.message = `Error while preprocessing ${filename}${e.message ? ` - ${e.message}` : ''}`;
				throw e;
			}

			if (preprocessed.dependencies?.length) {
				const checked = checkPreprocessDependencies(filename, preprocessed.dependencies);
				if (checked.warnings.length) {
					warnings.push(...checked.warnings);
				}
				if (checked.dependencies.length) {
					dependencies.push(...checked.dependencies);
				}
			}

			if (preprocessed.map) compileOptions.sourcemap = preprocessed.map;
		}
		if (typeof preprocessed?.map === 'object') {
			mapToRelative(preprocessed?.map, filename);
		}
		if (raw && svelteRequest.query.type === 'preprocessed') {
			// @ts-expect-error shortcut
			return /** @type {import('../types/compile.d.ts').CompileData} */ {
				preprocessed: preprocessed ?? { code }
			};
		}
		const finalCode = preprocessed ? preprocessed.code : code;
		const dynamicCompileOptions = await options?.dynamicCompileOptions?.({
			filename,
			code: finalCode,
			compileOptions
		});
		if (dynamicCompileOptions && log.debug.enabled) {
			log.debug(
				`dynamic compile options for  ${filename}: ${JSON.stringify(dynamicCompileOptions)}`,
				undefined,
				'compile'
			);
		}
		const finalCompileOptions = dynamicCompileOptions
			? {
					...compileOptions,
					...dynamicCompileOptions
				}
			: compileOptions;

		const endStat = stats?.start(filename);
		/** @type {import('svelte/types/compiler/interfaces').CompileResult} */
		let compiled;
		try {
			compiled = svelte.compile(finalCode, finalCompileOptions);
		} catch (e) {
			enhanceCompileError(e, code, preprocessors);
			throw e;
		}

		if (endStat) {
			endStat();
		}
		mapToRelative(compiled.js?.map, filename);
		mapToRelative(compiled.css?.map, filename);
		if (warnings.length) {
			if (!compiled.warnings) {
				compiled.warnings = [];
			}
			compiled.warnings.push(...warnings);
		}
		if (!raw) {
			// wire css import and code for hmr
			const hasCss = compiled.css?.code?.trim().length > 0;
			// compiler might not emit css with mode none or it may be empty
			if (emitCss && hasCss) {
				// TODO properly update sourcemap?
				compiled.js.code += `\nimport ${JSON.stringify(cssId)};\n`;
			}

			// only apply hmr when not in ssr context and hot options are set
			if (!ssr && makeHot) {
				compiled.js.code = makeHot({
					id: filename,
					compiledCode: compiled.js.code,
					//@ts-expect-error hot isn't a boolean at this point
					hotOptions: { ...options.hot, injectCss: options.hot?.injectCss === true && hasCss },
					compiled,
					originalCode: code,
					compileOptions: finalCompileOptions
				});
			}
		}

		let lang = 'js';
		for (const match of code.matchAll(scriptLangRE)) {
			if (match[1]) {
				lang = match[1];
				break;
			}
		}

		return {
			filename,
			normalizedFilename,
			lang,
			// @ts-ignore
			compiled,
			ssr,
			dependencies,
			preprocessed: preprocessed ?? { code }
		};
	};
};

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {Function | undefined}
 */
function buildMakeHot(options) {
	const needsMakeHot =
		!isSvelte5 && options.hot !== false && options.isServe && !options.isProduction;
	if (needsMakeHot) {
		// @ts-ignore
		const hotApi = options?.hot?.hotApi;
		// @ts-ignore
		const adapter = options?.hot?.adapter;
		return createMakeHot({
			walk: svelte.walk,
			hotApi,
			adapter,
			hotOptions: { noOverlay: true, .../** @type {object} */ (options.hot) }
		});
	}
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {import('../types/compile.d.ts').CompileSvelte}
 */
export function createCompileSvelte(options) {
	const makeHot = buildMakeHot(options);
	return _createCompileSvelte(makeHot);
}
