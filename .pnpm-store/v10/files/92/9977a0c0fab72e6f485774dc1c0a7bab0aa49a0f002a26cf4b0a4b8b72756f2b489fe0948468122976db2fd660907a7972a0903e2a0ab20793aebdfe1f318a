import MagicString from 'magic-string';
import { log } from './log.js';
import path from 'node:path';
import { normalizePath } from 'vite';

/**
 * this appends a *{} rule to component styles to force the svelte compiler to add style classes to all nodes
 * That means adding/removing class rules from <style> node won't trigger js updates as the scope classes are not changed
 *
 * only used during dev with enabled css hmr
 *
 * @returns {import('svelte/compiler').PreprocessorGroup}
 */
export function createInjectScopeEverythingRulePreprocessorGroup() {
	return {
		name: 'inject-scope-everything-rule',
		style({ content, filename }) {
			const s = new MagicString(content);
			s.append(' *{}');
			return {
				code: s.toString(),
				map: s.generateDecodedMap({
					source: filename ? path.basename(filename) : undefined,
					hires: true
				})
			};
		}
	};
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @param {import('vite').ResolvedConfig} config
 * @returns {{
 * 	prependPreprocessors: import('svelte/compiler').PreprocessorGroup[],
 * 	appendPreprocessors: import('svelte/compiler').PreprocessorGroup[]
 * }}
 */
function buildExtraPreprocessors(options, config) {
	/** @type {import('svelte/compiler').PreprocessorGroup[]} */
	const prependPreprocessors = [];
	/** @type {import('svelte/compiler').PreprocessorGroup[]} */
	const appendPreprocessors = [];

	// @ts-ignore
	const pluginsWithPreprocessorsDeprecated = config.plugins.filter((p) => p?.sveltePreprocess);
	if (pluginsWithPreprocessorsDeprecated.length > 0) {
		log.warn(
			`The following plugins use the deprecated 'plugin.sveltePreprocess' field. Please contact their maintainers and ask them to move it to 'plugin.api.sveltePreprocess': ${pluginsWithPreprocessorsDeprecated
				.map((p) => p.name)
				.join(', ')}`
		);
		// patch plugin to avoid breaking
		pluginsWithPreprocessorsDeprecated.forEach((p) => {
			if (!p.api) {
				p.api = {};
			}
			if (p.api.sveltePreprocess === undefined) {
				// @ts-ignore
				p.api.sveltePreprocess = p.sveltePreprocess;
			} else {
				log.error(
					`ignoring plugin.sveltePreprocess of ${p.name} because it already defined plugin.api.sveltePreprocess.`
				);
			}
		});
	}
	/** @type {import('vite').Plugin[]} */
	const pluginsWithPreprocessors = config.plugins.filter((p) => p?.api?.sveltePreprocess);
	/** @type {import('vite').Plugin[]} */
	const ignored = [];
	/** @type {import('vite').Plugin[]} */
	const included = [];
	for (const p of pluginsWithPreprocessors) {
		if (
			options.ignorePluginPreprocessors === true ||
			(Array.isArray(options.ignorePluginPreprocessors) &&
				options.ignorePluginPreprocessors?.includes(p.name))
		) {
			ignored.push(p);
		} else {
			included.push(p);
		}
	}
	if (ignored.length > 0) {
		log.debug(
			`Ignoring svelte preprocessors defined by these vite plugins: ${ignored
				.map((p) => p.name)
				.join(', ')}`,
			undefined,
			'preprocess'
		);
	}
	if (included.length > 0) {
		log.debug(
			`Adding svelte preprocessors defined by these vite plugins: ${included
				.map((p) => p.name)
				.join(', ')}`,
			undefined,
			'preprocess'
		);
		appendPreprocessors.push(...pluginsWithPreprocessors.map((p) => p.api.sveltePreprocess));
	}

	return { prependPreprocessors, appendPreprocessors };
}

/**
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @param {import('vite').ResolvedConfig} config
 */
export function addExtraPreprocessors(options, config) {
	const { prependPreprocessors, appendPreprocessors } = buildExtraPreprocessors(options, config);
	if (prependPreprocessors.length > 0 || appendPreprocessors.length > 0) {
		if (!options.preprocess) {
			options.preprocess = [...prependPreprocessors, ...appendPreprocessors];
		} else if (Array.isArray(options.preprocess)) {
			options.preprocess.unshift(...prependPreprocessors);
			options.preprocess.push(...appendPreprocessors);
		} else {
			options.preprocess = [...prependPreprocessors, options.preprocess, ...appendPreprocessors];
		}
	}
}

/**
 *
 * @param filename {string}
 * @param dependencies  {string[]}
 * @returns {({dependencies: string[], warnings:import('svelte/types/compiler/interfaces').Warning[] })}
 */
export function checkPreprocessDependencies(filename, dependencies) {
	/** @type {import('svelte/types/compiler/interfaces').Warning[]} */
	const warnings = [];

	// to find self, we have to compare normalized filenames, but must keep the original values in `dependencies`
	// because otherwise file watching on windows doesn't work
	// so we track idx and filter by that in the end
	/** @type {number[]} */
	const selfIdx = [];
	const normalizedFullFilename = normalizePath(filename);
	const normalizedDeps = dependencies.map(normalizePath);
	for (let i = 0; i < normalizedDeps.length; i++) {
		if (normalizedDeps[i] === normalizedFullFilename) {
			selfIdx.push(i);
		}
	}
	const hasSelfDependency = selfIdx.length > 0;
	if (hasSelfDependency) {
		warnings.push({
			code: 'vite-plugin-svelte-preprocess-depends-on-self',
			message:
				'svelte.preprocess returned this file as a dependency of itself. This can be caused by an invalid configuration or importing generated code that depends on .svelte files (eg. tailwind base css)',
			filename
		});
	}

	if (dependencies.length > 10) {
		warnings.push({
			code: 'vite-plugin-svelte-preprocess-many-dependencies',
			message: `svelte.preprocess depends on more than 10 external files which can cause slow builds and poor DX, try to reduce them. Found: ${dependencies.join(
				', '
			)}`,
			filename
		});
	}
	return {
		dependencies: hasSelfDependency
			? dependencies.filter((_, i) => !selfIdx.includes(i)) // remove self dependency
			: dependencies,
		warnings
	};
}
