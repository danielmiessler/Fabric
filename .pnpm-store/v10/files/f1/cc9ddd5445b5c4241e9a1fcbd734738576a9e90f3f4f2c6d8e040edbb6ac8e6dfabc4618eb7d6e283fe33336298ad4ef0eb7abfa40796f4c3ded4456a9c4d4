import { isCSSRequest, preprocessCSS, resolveConfig, transformWithEsbuild } from 'vite';
import { mapToRelative, removeLangSuffix } from './utils/sourcemaps.js';

/**
 * @typedef {(code: string, filename: string) => Promise<{ code: string; map?: any; deps?: Set<string> }>} CssTransform
 */

const supportedScriptLangs = ['ts'];

export const lang_sep = '.vite-preprocess';

/**
 * @param {import('./public.d.ts').VitePreprocessOptions} [opts]
 * @returns {import('svelte/compiler').PreprocessorGroup}
 */
export function vitePreprocess(opts) {
	/** @type {import('svelte/compiler').PreprocessorGroup} */
	const preprocessor = { name: 'vite-preprocess' };
	if (opts?.script !== false) {
		preprocessor.script = viteScript().script;
	}
	if (opts?.style !== false) {
		const styleOpts = typeof opts?.style == 'object' ? opts?.style : undefined;
		preprocessor.style = viteStyle(styleOpts).style;
	}
	return preprocessor;
}

/**
 * @returns {{ script: import('svelte/compiler').Preprocessor }}
 */
function viteScript() {
	return {
		async script({ attributes, content, filename = '' }) {
			const lang = /** @type {string} */ (attributes.lang);
			if (!supportedScriptLangs.includes(lang)) return;
			const { code, map } = await transformWithEsbuild(content, filename, {
				loader: /** @type {import('vite').ESBuildOptions['loader']} */ (lang),
				target: 'esnext',
				tsconfigRaw: {
					compilerOptions: {
						// svelte typescript needs this flag to work with type imports
						importsNotUsedAsValues: 'preserve',
						preserveValueImports: true
					}
				}
			});

			mapToRelative(map, filename);

			return {
				code,
				map
			};
		}
	};
}

/**
 * @param {import('vite').ResolvedConfig | import('vite').InlineConfig} config
 * @returns {{ style: import('svelte/compiler').Preprocessor }}
 */
function viteStyle(config = {}) {
	/** @type {Promise<CssTransform> | CssTransform} */
	let cssTransform;
	/** @type {import('svelte/compiler').Preprocessor} */
	const style = async ({ attributes, content, filename = '' }) => {
		const ext = attributes.lang ? `.${attributes.lang}` : '.css';
		if (attributes.lang && !isCSSRequest(ext)) return;
		if (!cssTransform) {
			cssTransform = createCssTransform(style, config).then((t) => (cssTransform = t));
		}
		const transform = await cssTransform;
		const suffix = `${lang_sep}${ext}`;
		const moduleId = `${filename}${suffix}`;
		const { code, map, deps } = await transform(content, moduleId);
		removeLangSuffix(map, suffix);
		mapToRelative(map, filename);
		const dependencies = deps ? Array.from(deps).filter((d) => !d.endsWith(suffix)) : undefined;
		return {
			code,
			map: map ?? undefined,
			dependencies
		};
	};
	// @ts-expect-error tag so can be found by v-p-s
	style.__resolvedConfig = null;
	return { style };
}

/**
 * @param {import('svelte/compiler').Preprocessor} style
 * @param {import('vite').ResolvedConfig | import('vite').InlineConfig} config
 * @returns {Promise<CssTransform>}
 */
async function createCssTransform(style, config) {
	/** @type {import('vite').ResolvedConfig} */
	let resolvedConfig;
	// @ts-expect-error special prop added if running in v-p-s
	if (style.__resolvedConfig) {
		// @ts-expect-error
		resolvedConfig = style.__resolvedConfig;
	} else if (isResolvedConfig(config)) {
		resolvedConfig = config;
	} else {
		resolvedConfig = await resolveConfig(
			config,
			process.env.NODE_ENV === 'production' ? 'build' : 'serve'
		);
	}
	return async (code, filename) => {
		return preprocessCSS(code, filename, resolvedConfig);
	};
}

/**
 * @param {any} config
 * @returns {config is import('vite').ResolvedConfig}
 */
function isResolvedConfig(config) {
	return !!config.inlineConfig;
}
