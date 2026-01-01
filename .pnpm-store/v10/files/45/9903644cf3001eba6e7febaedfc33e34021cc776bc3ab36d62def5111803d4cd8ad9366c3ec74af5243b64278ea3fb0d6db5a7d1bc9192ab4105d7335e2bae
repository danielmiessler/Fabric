import { log, logCompilerWarnings } from './utils/log.js';
import { toRollupError } from './utils/error.js';

/**
 * Vite-specific HMR handling
 *
 * @param {Function} compileSvelte
 * @param {import('vite').HmrContext} ctx
 * @param {import('./types/id.d.ts').SvelteRequest} svelteRequest
 * @param {import('./utils/vite-plugin-svelte-cache.js').VitePluginSvelteCache} cache
 * @param {import('./types/options.d.ts').ResolvedOptions} options
 * @returns {Promise<import('vite').ModuleNode[] | void>}
 */
export async function handleHotUpdate(compileSvelte, ctx, svelteRequest, cache, options) {
	if (!cache.has(svelteRequest)) {
		// file hasn't been requested yet (e.g. async component)
		log.debug(
			`handleHotUpdate called before initial transform for ${svelteRequest.id}`,
			undefined,
			'hmr'
		);
		return;
	}
	const { read, server, modules } = ctx;

	const cachedJS = cache.getJS(svelteRequest);
	const cachedCss = cache.getCSS(svelteRequest);

	const content = await read();
	/** @type {import('./types/compile.d.ts').CompileData} */
	let compileData;
	try {
		compileData = await compileSvelte(svelteRequest, content, options);
		cache.update(compileData);
	} catch (e) {
		cache.setError(svelteRequest, e);
		throw toRollupError(e, options);
	}

	const affectedModules = [...modules];

	const cssIdx = modules.findIndex((m) => m.id === svelteRequest.cssId);
	if (cssIdx > -1) {
		const cssUpdated = cssChanged(cachedCss, compileData.compiled.css);
		if (!cssUpdated) {
			log.debug(`skipping unchanged css for ${svelteRequest.cssId}`, undefined, 'hmr');
			affectedModules.splice(cssIdx, 1);
		}
	}
	const jsIdx = modules.findIndex((m) => m.id === svelteRequest.id);
	if (jsIdx > -1) {
		const jsUpdated = jsChanged(cachedJS, compileData.compiled.js, svelteRequest.filename);
		if (!jsUpdated) {
			log.debug(`skipping unchanged js for ${svelteRequest.id}`, undefined, 'hmr');
			affectedModules.splice(jsIdx, 1);
			// transform won't be called, log warnings here
			logCompilerWarnings(svelteRequest, compileData.compiled.warnings, options);
		}
	}

	// TODO is this enough? see also: https://github.com/vitejs/vite/issues/2274
	const ssrModulesToInvalidate = affectedModules.filter((m) => !!m.ssrTransformResult);
	if (ssrModulesToInvalidate.length > 0) {
		log.debug(
			`invalidating modules ${ssrModulesToInvalidate.map((m) => m.id).join(', ')}`,
			undefined,
			'hmr'
		);
		ssrModulesToInvalidate.forEach((moduleNode) => server.moduleGraph.invalidateModule(moduleNode));
	}
	if (affectedModules.length > 0) {
		log.debug(
			`handleHotUpdate for ${svelteRequest.id} result: ${affectedModules
				.map((m) => m.id)
				.join(', ')}`,
			undefined,
			'hmr'
		);
	}
	return affectedModules;
}

/**
 * @param {import('./types/compile.d.ts').Code} [prev]
 * @param {import('./types/compile.d.ts').Code} [next]
 * @returns {boolean}
 */
function cssChanged(prev, next) {
	return !isCodeEqual(prev?.code, next?.code);
}

/**
 * @param {import('./types/compile.d.ts').Code} [prev]
 * @param {import('./types/compile.d.ts').Code} [next]
 * @param {string} [filename]
 * @returns {boolean}
 */
function jsChanged(prev, next, filename) {
	const prevJs = prev?.code;
	const nextJs = next?.code;
	const isStrictEqual = isCodeEqual(prevJs, nextJs);
	if (isStrictEqual) {
		return false;
	}
	const isLooseEqual = isCodeEqual(normalizeJsCode(prevJs), normalizeJsCode(nextJs));
	if (!isStrictEqual && isLooseEqual) {
		log.debug(
			`ignoring compiler output js change for ${filename} as it is equal to previous output after normalization`,
			undefined,
			'hmr'
		);
	}
	return !isLooseEqual;
}

/**
 * @param {string} [prev]
 * @param {string} [next]
 * @returns {boolean}
 */
function isCodeEqual(prev, next) {
	if (!prev && !next) {
		return true;
	}
	if ((!prev && next) || (prev && !next)) {
		return false;
	}
	return prev === next;
}

/**
 * remove code that only changes metadata and does not require a js update for the component to keep working
 *
 * 1) add_location() calls. These add location metadata to elements, only used by some dev tools
 * 2) ... maybe more (or less) in the future
 *
 * @param {string} [code]
 * @returns {string | undefined}
 */
function normalizeJsCode(code) {
	if (!code) {
		return code;
	}
	return code.replace(/\s*\badd_location\s*\([^)]*\)\s*;?/g, '');
}
