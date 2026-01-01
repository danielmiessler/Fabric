import * as acorn from 'acorn';
import { walk } from 'estree-walker';
import { id, re } from './utils/id.js';
import { get_comment_handlers } from './utils/comments.js';

/** @typedef {import('estree').Expression} Expression */
/** @typedef {import('estree').Node} Node */
/** @typedef {import('estree').ObjectExpression} ObjectExpression */
/** @typedef {import('estree').Property} Property */
/** @typedef {import('estree').SpreadElement} SpreadElement */

/** @typedef {import('./utils/comments').CommentWithLocation} CommentWithLocation */

/** @type {Record<string, string>} */
const sigils = {
	'@': 'AT',
	'#': 'HASH'
};

/** @param {TemplateStringsArray} strings */
const join = (strings) => {
	let str = strings[0];
	for (let i = 1; i < strings.length; i += 1) {
		str += `_${id}_${i - 1}_${strings[i]}`;
	}
	return str.replace(
		/([@#])(\w+)/g,
		(_m, sigil, name) => `_${id}_${sigils[sigil]}_${name}`
	);
};

/**
 * @param {any[]} array
 * @param {any[]} target
 */
const flatten_body = (array, target) => {
	for (let i = 0; i < array.length; i += 1) {
		const statement = array[i];
		if (Array.isArray(statement)) {
			flatten_body(statement, target);
			continue;
		}

		if (statement.type === 'ExpressionStatement') {
			if (statement.expression === EMPTY) continue;

			if (Array.isArray(statement.expression)) {
				// TODO this is hacktacular
				let node = statement.expression[0];
				while (Array.isArray(node)) node = node[0];
				if (node) node.leadingComments = statement.leadingComments;

				flatten_body(statement.expression, target);
				continue;
			}

			if (/(Expression|Literal)$/.test(statement.expression.type)) {
				target.push(statement);
				continue;
			}

			if (statement.leadingComments)
				statement.expression.leadingComments = statement.leadingComments;
			if (statement.trailingComments)
				statement.expression.trailingComments = statement.trailingComments;

			target.push(statement.expression);
			continue;
		}

		target.push(statement);
	}

	return target;
};

/**
 * @param {any[]} array
 * @param {any[]} target
 */
const flatten_properties = (array, target) => {
	for (let i = 0; i < array.length; i += 1) {
		const property = array[i];

		if (property.value === EMPTY) continue;

		if (property.key === property.value && Array.isArray(property.key)) {
			flatten_properties(property.key, target);
			continue;
		}

		target.push(property);
	}

	return target;
};

/**
 * @param {any[]} nodes
 * @param {any[]} target
 */
const flatten = (nodes, target) => {
	for (let i = 0; i < nodes.length; i += 1) {
		const node = nodes[i];

		if (node === EMPTY) continue;

		if (Array.isArray(node)) {
			flatten(node, target);
			continue;
		}

		target.push(node);
	}

	return target;
};

const EMPTY = { type: 'Empty' };

/**
 *
 * @param {CommentWithLocation[]} comments
 * @param {string} raw
 * @returns {any}
 */
const acorn_opts = (comments, raw) => {
	const { onComment } = get_comment_handlers(comments, raw);
	return {
		ecmaVersion: 2022,
		sourceType: 'module',
		allowAwaitOutsideFunction: true,
		allowImportExportEverywhere: true,
		allowReturnOutsideFunction: true,
		onComment
	};
};

/**
 * @param {string} raw
 * @param {Node} node
 * @param {any[]} values
 * @param {CommentWithLocation[]} comments
 */
const inject = (raw, node, values, comments) => {
	comments.forEach((comment) => {
		comment.value = comment.value.replace(re, (m, i) =>
			+i in values ? values[+i] : m
		);
	});

	const { enter, leave } = get_comment_handlers(comments, raw);

	return walk(node, {
		enter,

		/** @param {any} node */
		leave(node) {
			if (node.type === 'Identifier') {
				re.lastIndex = 0;
				const match = re.exec(node.name);

				if (match) {
					if (match[1]) {
						if (+match[1] in values) {
							let value = values[+match[1]];

							if (typeof value === 'string') {
								value = {
									type: 'Identifier',
									name: value,
									leadingComments: node.leadingComments,
									trailingComments: node.trailingComments
								};
							} else if (typeof value === 'number') {
								value = {
									type: 'Literal',
									value,
									leadingComments: node.leadingComments,
									trailingComments: node.trailingComments
								};
							}

							this.replace(value || EMPTY);
						}
					} else {
						node.name = `${match[2] ? `@` : `#`}${match[4]}`;
					}
				}
			}

			if (node.type === 'Literal') {
				if (typeof node.value === 'string') {
					re.lastIndex = 0;
					const new_value = /** @type {string} */ (node.value).replace(
						re,
						(m, i) => (+i in values ? values[+i] : m)
					);
					const has_changed = new_value !== node.value;
					node.value = new_value;
					if (has_changed && node.raw) {
						// preserve the quotes
						node.raw = `${node.raw[0]}${JSON.stringify(node.value).slice(
							1,
							-1
						)}${node.raw[node.raw.length - 1]}`;
					}
				}
			}

			if (node.type === 'TemplateElement') {
				re.lastIndex = 0;
				node.value.raw = /** @type {string} */ (node.value.raw).replace(
					re,
					(m, i) => (+i in values ? values[+i] : m)
				);
			}

			if (node.type === 'Program' || node.type === 'BlockStatement') {
				node.body = flatten_body(node.body, []);
			}

			if (node.type === 'ObjectExpression' || node.type === 'ObjectPattern') {
				node.properties = flatten_properties(node.properties, []);
			}

			if (node.type === 'ArrayExpression' || node.type === 'ArrayPattern') {
				node.elements = flatten(node.elements, []);
			}

			if (
				node.type === 'FunctionExpression' ||
				node.type === 'FunctionDeclaration' ||
				node.type === 'ArrowFunctionExpression'
			) {
				node.params = flatten(node.params, []);
			}

			if (node.type === 'CallExpression' || node.type === 'NewExpression') {
				node.arguments = flatten(node.arguments, []);
			}

			if (
				node.type === 'ImportDeclaration' ||
				node.type === 'ExportNamedDeclaration'
			) {
				node.specifiers = flatten(node.specifiers, []);
			}

			if (node.type === 'ForStatement') {
				node.init = node.init === EMPTY ? null : node.init;
				node.test = node.test === EMPTY ? null : node.test;
				node.update = node.update === EMPTY ? null : node.update;
			}

			leave(node);
		}
	});
};

/**
 *
 * @param {TemplateStringsArray} strings
 * @param  {any[]} values
 * @returns {Node[]}
 */
export function b(strings, ...values) {
	const str = join(strings);

	/** @type {CommentWithLocation[]} */
	const comments = [];

	try {
		let ast = /** @type {any} */ (acorn.parse(str, acorn_opts(comments, str)));

		ast = inject(str, ast, values, comments);

		return ast.body;
	} catch (err) {
		handle_error(str, err);
	}
}

/**
 *
 * @param {TemplateStringsArray} strings
 * @param  {any[]} values
 * @returns {Expression & { start: Number, end: number }}
 */
export function x(strings, ...values) {
	const str = join(strings);

	/** @type {CommentWithLocation[]} */
	const comments = [];

	try {
		let expression =
			/** @type {Expression & { start: Number, end: number }} */ (
				acorn.parseExpressionAt(str, 0, acorn_opts(comments, str))
			);
		const match = /\S+/.exec(str.slice(expression.end));
		if (match) {
			throw new Error(`Unexpected token '${match[0]}'`);
		}

		expression = /** @type {Expression & { start: Number, end: number }} */ (
			inject(str, expression, values, comments)
		);

		return expression;
	} catch (err) {
		handle_error(str, err);
	}
}

/**
 *
 * @param {TemplateStringsArray} strings
 * @param  {any[]} values
 * @returns {(Property | SpreadElement) & { start: Number, end: number }}
 */
export function p(strings, ...values) {
	const str = `{${join(strings)}}`;

	/** @type {CommentWithLocation[]} */
	const comments = [];

	try {
		let expression = /** @type {any} */ (
			acorn.parseExpressionAt(str, 0, acorn_opts(comments, str))
		);

		expression = inject(str, expression, values, comments);

		return expression.properties[0];
	} catch (err) {
		handle_error(str, err);
	}
}

/**
 * @param {string} str
 * @param {Error} err
 */
function handle_error(str, err) {
	// TODO location/code frame

	re.lastIndex = 0;

	str = str.replace(re, (m, i, at, hash, name) => {
		if (at) return `@${name}`;
		if (hash) return `#${name}`;

		return '${...}';
	});

	console.log(`failed to parse:\n${str}`);
	throw err;
}

export { print } from './print/index.js';

/**
 * @param {string} source
 * @param {any} opts
 */
export const parse = (source, opts) => {
	/** @type {CommentWithLocation[]} */
	const comments = [];
	const { onComment, enter, leave } = get_comment_handlers(comments, source);
	const ast = /** @type {any} */ (acorn.parse(source, { onComment, ...opts }));
	walk(ast, { enter, leave });
	return ast;
};

/**
 * @param {string} source
 * @param {number} index
 * @param {any} opts
 */
export const parseExpressionAt = (source, index, opts) => {
	/** @type {CommentWithLocation[]} */
	const comments = [];
	const { onComment, enter, leave } = get_comment_handlers(comments, source);
	const ast = /** @type {any} */ (
		acorn.parseExpressionAt(source, index, { onComment, ...opts })
	);
	walk(ast, { enter, leave });
	return ast;
};
