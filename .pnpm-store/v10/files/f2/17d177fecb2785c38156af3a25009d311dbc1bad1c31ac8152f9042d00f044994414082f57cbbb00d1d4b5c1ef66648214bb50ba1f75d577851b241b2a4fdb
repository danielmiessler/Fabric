import fs from 'node:fs';
import { toRollupError } from './error.js';
import { log } from './log.js';
import { isSvelte4, isSvelte5 } from './svelte-version.js';
/**
 * utility function to compile ?raw and ?direct requests in load hook
 *
 * @param {import('../types/id.d.ts').SvelteRequest} svelteRequest
 * @param {import('../types/compile.d.ts').CompileSvelte} compileSvelte
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {Promise<string>}
 */
export async function loadRaw(svelteRequest, compileSvelte, options) {
	const { id, filename, query } = svelteRequest;

	// raw svelte subrequest, compile on the fly and return requested subpart
	let compileData;
	const source = fs.readFileSync(filename, 'utf-8');
	try {
		//avoid compileSvelte doing extra ssr stuff unless requested
		//@ts-ignore //@ts-expect-error generate value differs between svelte4 and 5
		svelteRequest.ssr = query.compilerOptions?.generate === (isSvelte4 ? 'ssr' : 'server');
		const type = query.type;
		compileData = await compileSvelte(svelteRequest, source, {
			...options,
			// don't use dynamic vite-plugin-svelte defaults here to ensure stable result between ssr,dev and build
			compilerOptions: {
				dev: false,
				css: 'external',
				enableSourcemap: isSvelte5
					? undefined
					: query.sourcemap
						? {
								js: type === 'script' || type === 'all',
								css: type === 'style' || type === 'all'
							}
						: false,
				...svelteRequest.query.compilerOptions
			},
			hot: false,
			emitCss: true
		});
	} catch (e) {
		throw toRollupError(e, options);
	}
	let result;
	if (query.type === 'style') {
		result = compileData.compiled.css;
	} else if (query.type === 'script') {
		result = compileData.compiled.js;
	} else if (query.type === 'preprocessed') {
		result = compileData.preprocessed;
	} else if (query.type === 'all' && query.raw) {
		return allToRawExports(compileData, source);
	} else {
		throw new Error(
			`invalid "type=${query.type}" in ${id}. supported are script, style, preprocessed, all`
		);
	}
	if (query.direct) {
		const supportedDirectTypes = ['script', 'style'];
		if (!supportedDirectTypes.includes(query.type)) {
			throw new Error(
				`invalid "type=${
					query.type
				}" combined with direct in ${id}. supported are: ${supportedDirectTypes.join(', ')}`
			);
		}
		log.debug(`load returns direct result for ${id}`, undefined, 'load');
		let directOutput = result.code;
		if (query.sourcemap && result.map?.toUrl) {
			const map = `sourceMappingURL=${result.map.toUrl()}`;
			if (query.type === 'style') {
				directOutput += `\n\n/*# ${map} */\n`;
			} else if (query.type === 'script') {
				directOutput += `\n\n//# ${map}\n`;
			}
		}
		return directOutput;
	} else if (query.raw) {
		log.debug(`load returns raw result for ${id}`, undefined, 'load');
		return toRawExports(result);
	} else {
		throw new Error(`invalid raw mode in ${id}, supported are raw, direct`);
	}
}

/**
 * turn compileData and source into a flat list of raw exports
 *
 * @param {import('../types/compile.d.ts').CompileData} compileData
 * @param {string} source
 */
function allToRawExports(compileData, source) {
	// flatten CompileData
	/** @type {Partial<import('../types/compile.d.ts').CompileData & { source: string }>} */
	const exports = {
		...compileData,
		...compileData.compiled,
		source
	};
	delete exports.compiled;
	delete exports.filename; // absolute path, remove to avoid it in output
	return toRawExports(exports);
}

/**
 * turn object into raw exports.
 *
 * every prop is returned as a const export, and if prop 'code' exists it is additionally added as default export
 *
 * eg {'foo':'bar','code':'baz'} results in
 *
 *  ```js
 *  export const code='baz'
 *  export const foo='bar'
 *  export default code
 *  ```
 * @param {object} object
 * @returns {string}
 */
function toRawExports(object) {
	let exports =
		Object.entries(object)
			//eslint-disable-next-line no-unused-vars
			.filter(([key, value]) => typeof value !== 'function') // preprocess output has a toString function that's enumerable
			.sort(([a], [b]) => (a < b ? -1 : a === b ? 0 : 1))
			.map(([key, value]) => `export const ${key}=${JSON.stringify(value)}`)
			.join('\n') + '\n';
	if (Object.prototype.hasOwnProperty.call(object, 'code')) {
		exports += 'export default code\n';
	}
	return exports;
}
