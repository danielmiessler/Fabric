import { buildExtendedLogMessage } from './log.js';

/**
 * convert an error thrown by svelte.compile to a RollupError so that vite displays it in a user friendly way
 * @param {import('svelte/types/compiler/interfaces').Warning & Error} error a svelte compiler error, which is a mix of Warning and an error
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {import('vite').Rollup.RollupError} the converted error
 */
export function toRollupError(error, options) {
	const { filename, frame, start, code, name, stack } = error;
	/** @type {import('vite').Rollup.RollupError} */
	const rollupError = {
		name, // needed otherwise sveltekit coalesce_to_error turns it into a string
		id: filename,
		message: buildExtendedLogMessage(error), // include filename:line:column so that it's clickable
		frame: formatFrameForVite(frame),
		code,
		stack: options.isBuild || options.isDebug || !frame ? stack : ''
	};
	if (start) {
		rollupError.loc = {
			line: start.line,
			column: start.column,
			file: filename
		};
	}
	return rollupError;
}

/**
 * convert an error thrown by svelte.compile to an esbuild PartialMessage
 * @param {import('svelte/types/compiler/interfaces').Warning & Error} error a svelte compiler error, which is a mix of Warning and an error
 * @param {import('../types/options.d.ts').ResolvedOptions} options
 * @returns {import('esbuild').PartialMessage} the converted error
 */
export function toESBuildError(error, options) {
	const { filename, frame, start, stack } = error;
	/** @type {import('esbuild').PartialMessage} */
	const partialMessage = {
		text: buildExtendedLogMessage(error)
	};
	if (start) {
		partialMessage.location = {
			line: start.line,
			column: start.column,
			file: filename,
			lineText: lineFromFrame(start.line, frame) // needed to get a meaningful error message on cli
		};
	}
	if (options.isBuild || options.isDebug || !frame) {
		partialMessage.detail = stack;
	}
	return partialMessage;
}

/**
 * extract line with number from codeframe
 *
 * @param {number} lineNo
 * @param {string} [frame]
 * @returns {string}
 */
function lineFromFrame(lineNo, frame) {
	if (!frame) {
		return '';
	}
	const lines = frame.split('\n');
	const errorLine = lines.find((line) => line.trimStart().startsWith(`${lineNo}: `));
	return errorLine ? errorLine.substring(errorLine.indexOf(': ') + 3) : '';
}

/**
 * vite error overlay expects a specific format to show frames
 * this reformats svelte frame (colon separated, less whitespace)
 * to one that vite displays on overlay ( pipe separated, more whitespace)
 * e.g.
 * ```
 * 1: foo
 * 2: bar;
 *       ^
 * 3: baz
 * ```
 * to
 * ```
 *  1 | foo
 *  2 | bar;
 *         ^
 *  3 | baz
 * ```
 * @see https://github.com/vitejs/vite/blob/96591bf9989529de839ba89958755eafe4c445ae/packages/vite/src/client/overlay.ts#L116
 * @param {string} [frame]
 * @returns {string}
 */
function formatFrameForVite(frame) {
	if (!frame) {
		return '';
	}
	return frame
		.split('\n')
		.map((line) => (line.match(/^\s+\^/) ? '   ' + line : ' ' + line.replace(':', ' | ')))
		.join('\n');
}

/**
 * @param {import('svelte/types/compiler/interfaces').Warning & Error} err a svelte compiler error, which is a mix of Warning and an error
 * @param {string} originalCode
 * @param {import('../public.d.ts').Options['preprocess']} [preprocessors]
 */
export function enhanceCompileError(err, originalCode, preprocessors) {
	preprocessors = arraify(preprocessors ?? []);

	/** @type {string[]} */
	const additionalMessages = [];

	// Handle incorrect TypeScript usage
	if (err.code === 'parse-error') {
		// Reference from Svelte: https://github.com/sveltejs/svelte/blob/9926347ad9dbdd0f3324d5538e25dcb7f5e442f8/packages/svelte/src/compiler/preprocess/index.js#L259
		const scriptRe =
			/<!--[^]*?-->|<script((?:\s+[^=>'"/]+=(?:"[^"]*"|'[^']*'|[^>\s]+)|\s+[^=>'"/]+)*\s*)(?:\/>|>([\S\s]*?)<\/script>)/g;
		const errIndex = err.pos ?? -1;

		let m;
		while ((m = scriptRe.exec(originalCode))) {
			const matchStart = m.index;
			const matchEnd = matchStart + m[0].length;
			const isErrorInScript = matchStart <= errIndex && errIndex <= matchEnd;
			if (isErrorInScript) {
				// Warn missing lang="ts"
				const hasLangTs = m[1]?.includes('lang="ts"');
				if (!hasLangTs) {
					additionalMessages.push('Did you forget to add lang="ts" to your script tag?');
				}
				// Warn missing script preprocessor
				if (preprocessors.every((p) => p.script == null)) {
					const preprocessorType = hasLangTs ? 'TypeScript' : 'script';
					additionalMessages.push(
						`Did you forget to add a ${preprocessorType} preprocessor? See https://github.com/sveltejs/vite-plugin-svelte/blob/main/docs/preprocess.md for more information.`
					);
				}
			}
		}
	}

	// Handle incorrect CSS preprocessor usage
	if (err.code === 'css-syntax-error') {
		// Reference from Svelte: https://github.com/sveltejs/svelte/blob/9926347ad9dbdd0f3324d5538e25dcb7f5e442f8/packages/svelte/src/compiler/preprocess/index.js#L257
		const styleRe =
			/<!--[^]*?-->|<style((?:\s+[^=>'"/]+=(?:"[^"]*"|'[^']*'|[^>\s]+)|\s+[^=>'"/]+)*\s*)(?:\/>|>([\S\s]*?)<\/style>)/g;

		let m;
		while ((m = styleRe.exec(originalCode))) {
			// Warn missing lang attribute
			if (!m[1]?.includes('lang=')) {
				additionalMessages.push('Did you forget to add a lang attribute to your style tag?');
			}
			// Warn missing style preprocessor
			if (
				preprocessors.every((p) => p.style == null || p.name === 'inject-scope-everything-rule')
			) {
				const preprocessorType = m[1]?.match(/lang="(.+?)"/)?.[1] ?? 'style';
				additionalMessages.push(
					`Did you forget to add a ${preprocessorType} preprocessor? See https://github.com/sveltejs/vite-plugin-svelte/blob/main/docs/preprocess.md for more information.`
				);
			}
		}
	}

	if (additionalMessages.length) {
		err.message += '\n\n- ' + additionalMessages.join('\n- ');
	}

	return err;
}

/**
 * @param {T | T[]} value
 * @template T
 */
function arraify(value) {
	return Array.isArray(value) ? value : [value];
}
