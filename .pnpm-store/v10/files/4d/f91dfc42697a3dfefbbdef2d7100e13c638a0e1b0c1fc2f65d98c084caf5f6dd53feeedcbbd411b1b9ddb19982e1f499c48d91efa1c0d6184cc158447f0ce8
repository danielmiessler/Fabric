import * as perisopic from 'periscopic';
import { handle } from './handlers.js';
import { encode } from '@jridgewell/sourcemap-codec';

/** @type {(str?: string) => string} str */
let btoa = () => {
	throw new Error(
		'Unsupported environment: `window.btoa` or `Buffer` should be supported.'
	);
};

if (typeof window !== 'undefined' && typeof window.btoa === 'function') {
	btoa = (str) => window.btoa(unescape(encodeURIComponent(str)));
} else if (typeof Buffer === 'function') {
	btoa = (str) => Buffer.from(str, 'utf-8').toString('base64');
}

/** @typedef {import('estree').Node} Node */

/**
 * @typedef {{
 *   file?: string;
 *   sourceMapSource?: string;
 *   sourceMapContent?: string;
 *   sourceMapEncodeMappings?: boolean; // default true
 *   getName?: (name: string) => string;
 * }} PrintOptions
 */

/**
 * @param {Node} node
 * @param {PrintOptions} opts
 * @returns {{ code: string, map: any }} // TODO
 */
export function print(node, opts = {}) {
	if (Array.isArray(node)) {
		return print(
			{
				type: 'Program',
				body: node,
				sourceType: 'module'
			},
			opts
		);
	}

	const {
		getName = /** @param {string} x */ (x) => {
			throw new Error(`Unhandled sigil @${x}`);
		}
	} = opts;

	let { map: scope_map, scope } = perisopic.analyze(node);
	const deconflicted = new WeakMap();

	const chunks = handle(node, {
		indent: '',
		getName,
		scope,
		scope_map,
		deconflicted,
		comments: []
	});

	/** @typedef {[number, number, number, number]} Segment */

	let code = '';
	let current_column = 0;

	/** @type {Segment[][]} */
	let mappings = [];

	/** @type {Segment[]} */
	let current_line = [];

	for (let i = 0; i < chunks.length; i += 1) {
		const chunk = chunks[i];

		code += chunk.content;

		if (chunk.loc) {
			current_line.push([
				current_column,
				0, // source index is always zero
				chunk.loc.start.line - 1,
				chunk.loc.start.column
			]);
		}

		for (let i = 0; i < chunk.content.length; i += 1) {
			if (chunk.content[i] === '\n') {
				mappings.push(current_line);
				current_line = [];
				current_column = 0;
			} else {
				current_column += 1;
			}
		}

		if (chunk.loc) {
			current_line.push([
				current_column,
				0, // source index is always zero
				chunk.loc.end.line - 1,
				chunk.loc.end.column
			]);
		}
	}

	mappings.push(current_line);

	const map = {
		version: 3,
		/** @type {string[]} */
		names: [],
		sources: [opts.sourceMapSource || null],
		sourcesContent: [opts.sourceMapContent || null],
		mappings:
			opts.sourceMapEncodeMappings == undefined || opts.sourceMapEncodeMappings
				? encode(mappings)
				: mappings
	};

	Object.defineProperties(map, {
		toString: {
			enumerable: false,
			value: function toString() {
				return JSON.stringify(this);
			}
		},
		toUrl: {
			enumerable: false,
			value: function toUrl() {
				return (
					'data:application/json;charset=utf-8;base64,' + btoa(this.toString())
				);
			}
		}
	});

	return {
		code,
		map
	};
}
