// heavily based on https://github.com/davidbonnet/astring
// released under MIT license https://github.com/davidbonnet/astring/blob/master/LICENSE

import { re } from '../utils/id.js';
import { push_array } from '../utils/push_array.js';

/** @typedef {import('estree').ArrowFunctionExpression} ArrowFunctionExpression */
/** @typedef {import('estree').BinaryExpression} BinaryExpression */
/** @typedef {import('estree').CallExpression} CallExpression */
/** @typedef {import('estree').Comment} Comment */
/** @typedef {import('estree').ExportSpecifier} ExportSpecifier */
/** @typedef {import('estree').Expression} Expression */
/** @typedef {import('estree').FunctionDeclaration} FunctionDeclaration */
/** @typedef {import('estree').ImportDeclaration} ImportDeclaration */
/** @typedef {import('estree').ImportSpecifier} ImportSpecifier */
/** @typedef {import('estree').Literal} Literal */
/** @typedef {import('estree').LogicalExpression} LogicalExpression */
/** @typedef {import('estree').NewExpression} NewExpression */
/** @typedef {import('estree').Node} Node */
/** @typedef {import('estree').ObjectExpression} ObjectExpression */
/** @typedef {import('estree').Pattern} Pattern */
/** @typedef {import('estree').Property} Property */
/** @typedef {import('estree').PropertyDefinition} PropertyDefinition */
/** @typedef {import('estree').SequenceExpression} SequenceExpression */
/** @typedef {import('estree').SimpleCallExpression} SimpleCallExpression */
/** @typedef {import('estree').SwitchStatement} SwitchStatement */
/** @typedef {import('estree').VariableDeclaration} VariableDeclaration */
/** @typedef {import('estree').StaticBlock} StaticBlock */
/** @typedef {import('estree').PrivateIdentifier} PrivateIdenifier*/

/**
 * @typedef {{
 *   content: string;
 *   loc?: {
 *     start: { line: number; column: number; };
 *     end: { line: number; column: number; };
 *   };
 *   has_newline: boolean;
 * }} Chunk
 */

/**
 * @typedef {(node: any, state: State) => Chunk[]} Handler
 */

/**
 * @typedef {{
 *   indent: string;
 *   scope: any; // TODO import from periscopic
 *   scope_map: WeakMap<Node, any>;
 *   getName: (name: string) => string;
 *   deconflicted: WeakMap<Node, Map<string, string>>;
 *   comments: Comment[];
 * }} State
 */

/**
 * @param {Node} node
 * @param {State} state
 * @returns {Chunk[]}
 */
export function handle(node, state) {
	const handler = handlers[node.type];

	if (!handler) {
		throw new Error(`Not implemented ${node.type}`);
	}

	const result = handler(node, state);

	if (node.leadingComments) {
		result.unshift(
			c(
				node.leadingComments
					.map((comment) =>
						comment.type === 'Block'
							? `/*${comment.value}*/${
									/** @type {any} */ (comment).has_trailing_newline
										? `\n${state.indent}`
										: ` `
							  }`
							: `//${comment.value}${
									/** @type {any} */ (comment).has_trailing_newline
										? `\n${state.indent}`
										: ` `
							  }`
					)
					.join(``)
			)
		);
	}

	if (node.trailingComments) {
		state.comments.push(node.trailingComments[0]); // there is only ever one
	}

	return result;
}

/**
 * @param {string} content
 * @param {Node} [node]
 * @returns {Chunk}
 */
function c(content, node) {
	return {
		content,
		loc: node && node.loc,
		has_newline: /\n/.test(content)
	};
}

const OPERATOR_PRECEDENCE = {
	'||': 2,
	'&&': 3,
	'??': 4,
	'|': 5,
	'^': 6,
	'&': 7,
	'==': 8,
	'!=': 8,
	'===': 8,
	'!==': 8,
	'<': 9,
	'>': 9,
	'<=': 9,
	'>=': 9,
	in: 9,
	instanceof: 9,
	'<<': 10,
	'>>': 10,
	'>>>': 10,
	'+': 11,
	'-': 11,
	'*': 12,
	'%': 12,
	'/': 12,
	'**': 13
};

/** @type {Record<string, number>} */
const EXPRESSIONS_PRECEDENCE = {
	ArrayExpression: 20,
	TaggedTemplateExpression: 20,
	ThisExpression: 20,
	Identifier: 20,
	Literal: 18,
	TemplateLiteral: 20,
	Super: 20,
	SequenceExpression: 20,
	MemberExpression: 19,
	CallExpression: 19,
	NewExpression: 19,
	AwaitExpression: 17,
	ClassExpression: 17,
	FunctionExpression: 17,
	ObjectExpression: 17,
	UpdateExpression: 16,
	UnaryExpression: 15,
	BinaryExpression: 14,
	LogicalExpression: 13,
	ConditionalExpression: 4,
	ArrowFunctionExpression: 3,
	AssignmentExpression: 3,
	YieldExpression: 2,
	RestElement: 1
};

/**
 *
 * @param {Expression} node
 * @param {BinaryExpression | LogicalExpression} parent
 * @param {boolean} is_right
 * @returns
 */
function needs_parens(node, parent, is_right) {
	// special case where logical expressions and coalesce expressions cannot be mixed,
	// either of them need to be wrapped with parentheses
	if (
		node.type === 'LogicalExpression' &&
		parent.type === 'LogicalExpression' &&
		((parent.operator === '??' && node.operator !== '??') ||
			(parent.operator !== '??' && node.operator === '??'))
	) {
		return true;
	}

	const precedence = EXPRESSIONS_PRECEDENCE[node.type];
	const parent_precedence = EXPRESSIONS_PRECEDENCE[parent.type];

	if (precedence !== parent_precedence) {
		// Different node types
		return (
			(!is_right &&
				precedence === 15 &&
				parent_precedence === 14 &&
				parent.operator === '**') ||
			precedence < parent_precedence
		);
	}

	if (precedence !== 13 && precedence !== 14) {
		// Not a `LogicalExpression` or `BinaryExpression`
		return false;
	}

	if (
		/** @type {BinaryExpression} */ (node).operator === '**' &&
		parent.operator === '**'
	) {
		// Exponentiation operator has right-to-left associativity
		return !is_right;
	}

	if (is_right) {
		// Parenthesis are used if both operators have the same precedence
		return (
			OPERATOR_PRECEDENCE[/** @type {BinaryExpression} */ (node).operator] <=
			OPERATOR_PRECEDENCE[parent.operator]
		);
	}

	return (
		OPERATOR_PRECEDENCE[/** @type {BinaryExpression} */ (node).operator] <
		OPERATOR_PRECEDENCE[parent.operator]
	);
}

/** @param {Node} node */
function has_call_expression(node) {
	while (node) {
		if (node.type[0] === 'CallExpression') {
			return true;
		} else if (node.type === 'MemberExpression') {
			node = node.object;
		} else {
			return false;
		}
	}
}

/** @param {Chunk[]} chunks */
const has_newline = (chunks) => {
	for (let i = 0; i < chunks.length; i += 1) {
		if (chunks[i].has_newline) return true;
	}
	return false;
};

/** @param {Chunk[]} chunks */
const get_length = (chunks) => {
	let total = 0;
	for (let i = 0; i < chunks.length; i += 1) {
		total += chunks[i].content.length;
	}
	return total;
};

/**
 * @param {number} a
 * @param {number} b
 */
const sum = (a, b) => a + b;

/**
 * @param {Chunk[][]} nodes
 * @param {Chunk} separator
 * @returns {Chunk[]}
 */
const join = (nodes, separator) => {
	if (nodes.length === 0) return [];

	const joined = [...nodes[0]];
	for (let i = 1; i < nodes.length; i += 1) {
		joined.push(separator);
		push_array(joined, nodes[i]);
	}
	return joined;
};

/**
 * @param {(node: any, state: State) => Chunk[]} fn
 */
const scoped = (fn) => {
	/**
	 * @param {any} node
	 * @param {State} state
	 */
	const scoped_fn = (node, state) => {
		return fn(node, {
			...state,
			scope: state.scope_map.get(node)
		});
	};

	return scoped_fn;
};

/**
 * @param {string} name
 * @param {Set<string>} names
 */
const deconflict = (name, names) => {
	const original = name;
	let i = 1;

	while (names.has(name)) {
		name = `${original}$${i++}`;
	}

	return name;
};

/**
 * @param {Node[]} nodes
 * @param {State} state
 */
const handle_body = (nodes, state) => {
	const chunks = [];

	const body = nodes.map((statement) => {
		const chunks = handle(statement, {
			...state,
			indent: state.indent
		});

		let add_newline = false;

		while (state.comments.length) {
			const comment = state.comments.shift();
			const prefix = add_newline ? `\n${state.indent}` : ` `;

			chunks.push(
				c(
					comment.type === 'Block'
						? `${prefix}/*${comment.value}*/`
						: `${prefix}//${comment.value}`
				)
			);

			add_newline = comment.type === 'Line';
		}

		return chunks;
	});

	let needed_padding = false;

	for (let i = 0; i < body.length; i += 1) {
		const needs_padding = has_newline(body[i]);

		if (i > 0) {
			chunks.push(
				c(
					needs_padding || needed_padding
						? `\n\n${state.indent}`
						: `\n${state.indent}`
				)
			);
		}

		push_array(chunks, body[i]);

		needed_padding = needs_padding;
	}

	return chunks;
};

/**
 * @param {VariableDeclaration} node
 * @param {State} state
 */
const handle_var_declaration = (node, state) => {
	const chunks = [c(`${node.kind} `)];

	const declarators = node.declarations.map((d) =>
		handle(d, {
			...state,
			indent: state.indent + (node.declarations.length === 1 ? '' : '\t')
		})
	);

	const multiple_lines =
		declarators.some(has_newline) ||
		declarators.map(get_length).reduce(sum, 0) +
			(state.indent.length + declarators.length - 1) * 2 >
			80;

	const separator = c(multiple_lines ? `,\n${state.indent}\t` : ', ');

	push_array(chunks, join(declarators, separator));

	return chunks;
};

/** @type {Record<string, Handler>} */
const handlers = {
	Program(node, state) {
		return handle_body(node.body, state);
	},

	BlockStatement: scoped((node, state) => {
		return [
			c(`{\n${state.indent}\t`),
			...handle_body(node.body, { ...state, indent: state.indent + '\t' }),
			c(`\n${state.indent}}`)
		];
	}),

	EmptyStatement(node, state) {
		return [c(';')];
	},

	ParenthesizedExpression(node, state) {
		return handle(node.expression, state);
	},

	ExpressionStatement(node, state) {
		if (
			node.expression.type === 'AssignmentExpression' &&
			node.expression.left.type === 'ObjectPattern'
		) {
			// is an AssignmentExpression to an ObjectPattern
			return [c('('), ...handle(node.expression, state), c(');')];
		}

		return [...handle(node.expression, state), c(';')];
	},

	IfStatement(node, state) {
		const chunks = [
			c('if ('),
			...handle(node.test, state),
			c(') '),
			...handle(node.consequent, state)
		];

		if (node.alternate) {
			chunks.push(c(' else '));
			push_array(chunks, handle(node.alternate, state));
		}

		return chunks;
	},

	LabeledStatement(node, state) {
		return [...handle(node.label, state), c(': '), ...handle(node.body, state)];
	},

	BreakStatement(node, state) {
		return node.label
			? [c('break '), ...handle(node.label, state), c(';')]
			: [c('break;')];
	},

	ContinueStatement(node, state) {
		return node.label
			? [c('continue '), ...handle(node.label, state), c(';')]
			: [c('continue;')];
	},

	WithStatement(node, state) {
		return [
			c('with ('),
			...handle(node.object, state),
			c(') '),
			...handle(node.body, state)
		];
	},

	SwitchStatement(/** @type {SwitchStatement} */ node, state) {
		const chunks = [
			c('switch ('),
			...handle(node.discriminant, state),
			c(') {')
		];

		node.cases.forEach((block) => {
			if (block.test) {
				chunks.push(c(`\n${state.indent}\tcase `));
				push_array(
					chunks,
					handle(block.test, { ...state, indent: `${state.indent}\t` })
				);
				chunks.push(c(':'));
			} else {
				chunks.push(c(`\n${state.indent}\tdefault:`));
			}

			block.consequent.forEach((statement) => {
				chunks.push(c(`\n${state.indent}\t\t`));
				push_array(
					chunks,
					handle(statement, { ...state, indent: `${state.indent}\t\t` })
				);
			});
		});

		chunks.push(c(`\n${state.indent}}`));

		return chunks;
	},

	ReturnStatement(node, state) {
		if (node.argument) {
			const contains_comment =
				node.argument.leadingComments &&
				node.argument.leadingComments.some(
					(
						/** @type import('../utils/comments.js').CommentWithLocation */ comment
					) => comment.has_trailing_newline
				);
			return [
				c(contains_comment ? 'return (' : 'return '),
				...handle(node.argument, state),
				c(contains_comment ? ');' : ';')
			];
		} else {
			return [c('return;')];
		}
	},

	ThrowStatement(node, state) {
		return [c('throw '), ...handle(node.argument, state), c(';')];
	},

	TryStatement(node, state) {
		const chunks = [c('try '), ...handle(node.block, state)];

		if (node.handler) {
			if (node.handler.param) {
				chunks.push(c(' catch('));
				push_array(chunks, handle(node.handler.param, state));
				chunks.push(c(') '));
			} else {
				chunks.push(c(' catch '));
			}

			push_array(chunks, handle(node.handler.body, state));
		}

		if (node.finalizer) {
			chunks.push(c(' finally '));
			push_array(chunks, handle(node.finalizer, state));
		}

		return chunks;
	},

	WhileStatement(node, state) {
		return [
			c('while ('),
			...handle(node.test, state),
			c(') '),
			...handle(node.body, state)
		];
	},

	DoWhileStatement(node, state) {
		return [
			c('do '),
			...handle(node.body, state),
			c(' while ('),
			...handle(node.test, state),
			c(');')
		];
	},

	ForStatement: scoped((node, state) => {
		const chunks = [c('for (')];

		if (node.init) {
			if (node.init.type === 'VariableDeclaration') {
				push_array(chunks, handle_var_declaration(node.init, state));
			} else {
				push_array(chunks, handle(node.init, state));
			}
		}

		chunks.push(c('; '));
		if (node.test) push_array(chunks, handle(node.test, state));
		chunks.push(c('; '));
		if (node.update) push_array(chunks, handle(node.update, state));

		chunks.push(c(') '));
		push_array(chunks, handle(node.body, state));

		return chunks;
	}),

	ForInStatement: scoped((node, state) => {
		const chunks = [c(`for ${node.await ? 'await ' : ''}(`)];

		if (node.left.type === 'VariableDeclaration') {
			push_array(chunks, handle_var_declaration(node.left, state));
		} else {
			push_array(chunks, handle(node.left, state));
		}

		chunks.push(c(node.type === 'ForInStatement' ? ` in ` : ` of `));
		push_array(chunks, handle(node.right, state));
		chunks.push(c(') '));
		push_array(chunks, handle(node.body, state));

		return chunks;
	}),

	DebuggerStatement(node, state) {
		return [c('debugger', node), c(';')];
	},

	FunctionDeclaration: scoped(
		(/** @type {FunctionDeclaration} */ node, state) => {
			const chunks = [];

			if (node.async) chunks.push(c('async '));
			chunks.push(c(node.generator ? 'function* ' : 'function '));
			if (node.id) push_array(chunks, handle(node.id, state));
			chunks.push(c('('));

			const params = node.params.map((p) =>
				handle(p, {
					...state,
					indent: state.indent + '\t'
				})
			);

			const multiple_lines =
				params.some(has_newline) ||
				params.map(get_length).reduce(sum, 0) +
					(state.indent.length + params.length - 1) * 2 >
					80;

			const separator = c(multiple_lines ? `,\n${state.indent}` : ', ');

			if (multiple_lines) {
				chunks.push(c(`\n${state.indent}\t`));
				push_array(chunks, join(params, separator));
				chunks.push(c(`\n${state.indent}`));
			} else {
				push_array(chunks, join(params, separator));
			}

			chunks.push(c(') '));
			push_array(chunks, handle(node.body, state));

			return chunks;
		}
	),

	VariableDeclaration(node, state) {
		return handle_var_declaration(node, state).concat(c(';'));
	},

	VariableDeclarator(node, state) {
		if (node.init) {
			return [...handle(node.id, state), c(' = '), ...handle(node.init, state)];
		} else {
			return handle(node.id, state);
		}
	},

	ClassDeclaration(node, state) {
		const chunks = [c('class ')];

		if (node.id) {
			push_array(chunks, handle(node.id, state));
			chunks.push(c(' '));
		}

		if (node.superClass) {
			chunks.push(c('extends '));
			push_array(chunks, handle(node.superClass, state));
			chunks.push(c(' '));
		}

		push_array(chunks, handle(node.body, state));

		return chunks;
	},

	ImportDeclaration(/** @type {ImportDeclaration} */ node, state) {
		const chunks = [c('import ')];

		const { length } = node.specifiers;
		const source = handle(node.source, state);

		if (length > 0) {
			let i = 0;

			while (i < length) {
				if (i > 0) {
					chunks.push(c(', '));
				}

				const specifier = node.specifiers[i];

				if (specifier.type === 'ImportDefaultSpecifier') {
					chunks.push(c(specifier.local.name, specifier));
					i += 1;
				} else if (specifier.type === 'ImportNamespaceSpecifier') {
					chunks.push(c('* as ' + specifier.local.name, specifier));
					i += 1;
				} else {
					break;
				}
			}

			if (i < length) {
				// we have named specifiers
				const specifiers = node.specifiers
					.slice(i)
					.map((/** @type {ImportSpecifier} */ specifier) => {
						const name = handle(specifier.imported, state)[0];
						const as = handle(specifier.local, state)[0];

						if (name.content === as.content) {
							return [as];
						}

						return [name, c(' as '), as];
					});

				const width =
					get_length(chunks) +
					specifiers.map(get_length).reduce(sum, 0) +
					2 * specifiers.length +
					6 +
					get_length(source);

				if (width > 80) {
					chunks.push(c(`{\n\t`));
					push_array(chunks, join(specifiers, c(',\n\t')));
					chunks.push(c('\n}'));
				} else {
					chunks.push(c(`{ `));
					push_array(chunks, join(specifiers, c(', ')));
					chunks.push(c(' }'));
				}
			}

			chunks.push(c(' from '));
		}

		push_array(chunks, source);
		chunks.push(c(';'));

		return chunks;
	},

	ImportExpression(node, state) {
		return [c('import('), ...handle(node.source, state), c(')')];
	},

	ExportDefaultDeclaration(node, state) {
		const chunks = [c(`export default `), ...handle(node.declaration, state)];

		if (node.declaration.type !== 'FunctionDeclaration') {
			chunks.push(c(';'));
		}

		return chunks;
	},

	ExportNamedDeclaration(node, state) {
		const chunks = [c('export ')];

		if (node.declaration) {
			push_array(chunks, handle(node.declaration, state));
		} else {
			const specifiers = node.specifiers.map(
				(/** @type {ExportSpecifier} */ specifier) => {
					const name = handle(specifier.local, state)[0];
					const as = handle(specifier.exported, state)[0];

					if (name.content === as.content) {
						return [name];
					}

					return [name, c(' as '), as];
				}
			);

			const width =
				7 + specifiers.map(get_length).reduce(sum, 0) + 2 * specifiers.length;

			if (width > 80) {
				chunks.push(c('{\n\t'));
				push_array(chunks, join(specifiers, c(',\n\t')));
				chunks.push(c('\n}'));
			} else {
				chunks.push(c('{ '));
				push_array(chunks, join(specifiers, c(', ')));
				chunks.push(c(' }'));
			}

			if (node.source) {
				chunks.push(c(' from '));
				push_array(chunks, handle(node.source, state));
			}
		}

		chunks.push(c(';'));

		return chunks;
	},

	ExportAllDeclaration(node, state) {
		return [c(`export * from `), ...handle(node.source, state), c(`;`)];
	},

	MethodDefinition(node, state) {
		const chunks = [];

		if (node.static) {
			chunks.push(c('static '));
		}

		if (node.kind === 'get' || node.kind === 'set') {
			// Getter or setter
			chunks.push(c(node.kind + ' '));
		}

		if (node.value.async) {
			chunks.push(c('async '));
		}

		if (node.value.generator) {
			chunks.push(c('*'));
		}

		if (node.computed) {
			chunks.push(c('['));
			push_array(chunks, handle(node.key, state));
			chunks.push(c(']'));
		} else {
			push_array(chunks, handle(node.key, state));
		}

		chunks.push(c('('));

		const { params } = node.value;
		for (let i = 0; i < params.length; i += 1) {
			push_array(chunks, handle(params[i], state));
			if (i < params.length - 1) chunks.push(c(', '));
		}

		chunks.push(c(') '));
		push_array(chunks, handle(node.value.body, state));

		return chunks;
	},

	ArrowFunctionExpression: scoped(
		(/** @type {ArrowFunctionExpression} */ node, state) => {
			const chunks = [];

			if (node.async) chunks.push(c('async '));

			if (node.params.length === 1 && node.params[0].type === 'Identifier') {
				push_array(chunks, handle(node.params[0], state));
			} else {
				const params = node.params.map((param) =>
					handle(param, {
						...state,
						indent: state.indent + '\t'
					})
				);

				chunks.push(c('('));
				push_array(chunks, join(params, c(', ')));
				chunks.push(c(')'));
			}

			chunks.push(c(' => '));

			if (
				node.body.type === 'ObjectExpression' ||
				(node.body.type === 'AssignmentExpression' &&
					node.body.left.type === 'ObjectPattern')
			) {
				chunks.push(c('('));
				push_array(chunks, handle(node.body, state));
				chunks.push(c(')'));
			} else {
				push_array(chunks, handle(node.body, state));
			}

			return chunks;
		}
	),

	ThisExpression(node, state) {
		return [c('this', node)];
	},

	Super(node, state) {
		return [c('super', node)];
	},

	RestElement(node, state) {
		return [c('...'), ...handle(node.argument, state)];
	},

	YieldExpression(node, state) {
		if (node.argument) {
			return [
				c(node.delegate ? `yield* ` : `yield `),
				...handle(node.argument, state)
			];
		}

		return [c(node.delegate ? `yield*` : `yield`)];
	},

	AwaitExpression(node, state) {
		if (node.argument) {
			const precedence = EXPRESSIONS_PRECEDENCE[node.argument.type];

			if (precedence && precedence < EXPRESSIONS_PRECEDENCE.AwaitExpression) {
				return [c('await ('), ...handle(node.argument, state), c(')')];
			} else {
				return [c('await '), ...handle(node.argument, state)];
			}
		}

		return [c('await')];
	},

	TemplateLiteral(node, state) {
		const chunks = [c('`')];

		const { quasis, expressions } = node;

		for (let i = 0; i < expressions.length; i++) {
			chunks.push(c(quasis[i].value.raw), c('${'));
			push_array(chunks, handle(expressions[i], state));
			chunks.push(c('}'));
		}

		chunks.push(c(quasis[quasis.length - 1].value.raw), c('`'));

		return chunks;
	},

	TaggedTemplateExpression(node, state) {
		return handle(node.tag, state).concat(handle(node.quasi, state));
	},

	ArrayExpression(node, state) {
		const chunks = [c('[')];

		/** @type {Chunk[][]} */
		const elements = [];

		/** @type {Chunk[]} */
		let sparse_commas = [];

		for (let i = 0; i < node.elements.length; i += 1) {
			// can't use map/forEach because of sparse arrays
			const element = node.elements[i];
			if (element) {
				elements.push([
					...sparse_commas,
					...handle(element, {
						...state,
						indent: state.indent + '\t'
					})
				]);
				sparse_commas = [];
			} else {
				sparse_commas.push(c(','));
			}
		}

		const multiple_lines =
			elements.some(has_newline) ||
			elements.map(get_length).reduce(sum, 0) +
				(state.indent.length + elements.length - 1) * 2 >
				80;

		if (multiple_lines) {
			chunks.push(c(`\n${state.indent}\t`));
			push_array(chunks, join(elements, c(`,\n${state.indent}\t`)));
			chunks.push(c(`\n${state.indent}`));
			push_array(chunks, sparse_commas);
		} else {
			push_array(chunks, join(elements, c(', ')));
			push_array(chunks, sparse_commas);
		}

		chunks.push(c(']'));

		return chunks;
	},

	ObjectExpression(/** @type {ObjectExpression} */ node, state) {
		if (node.properties.length === 0) {
			return [c('{}')];
		}

		let has_inline_comment = false;

		/** @type {Chunk[]} */
		const chunks = [];
		const separator = c(', ');

		node.properties.forEach((p, i) => {
			push_array(
				chunks,
				handle(p, {
					...state,
					indent: state.indent + '\t'
				})
			);

			if (state.comments.length) {
				// TODO generalise this, so it works with ArrayExpressions and other things.
				// At present, stuff will just get appended to the closest statement/declaration
				chunks.push(c(', '));

				while (state.comments.length) {
					const comment = state.comments.shift();

					chunks.push(
						c(
							comment.type === 'Block'
								? `/*${comment.value}*/\n${state.indent}\t`
								: `//${comment.value}\n${state.indent}\t`
						)
					);

					if (comment.type === 'Line') {
						has_inline_comment = true;
					}
				}
			} else {
				if (i < node.properties.length - 1) {
					chunks.push(separator);
				}
			}
		});

		const multiple_lines =
			has_inline_comment || has_newline(chunks) || get_length(chunks) > 40;

		if (multiple_lines) {
			separator.content = `,\n${state.indent}\t`;
		}

		return [
			c(multiple_lines ? `{\n${state.indent}\t` : `{ `),
			...chunks,
			c(multiple_lines ? `\n${state.indent}}` : ` }`)
		];
	},

	Property(node, state) {
		const value = handle(node.value, state);

		if (node.key === node.value) {
			return value;
		}

		// special case
		if (
			!node.computed &&
			node.value.type === 'AssignmentPattern' &&
			node.value.left.type === 'Identifier' &&
			node.value.left.name === node.key.name
		) {
			return value;
		}

		if (
			!node.computed &&
			node.value.type === 'Identifier' &&
			((node.key.type === 'Identifier' && node.key.name === value[0].content) ||
				(node.key.type === 'Literal' && node.key.value === value[0].content))
		) {
			return value;
		}

		const key = handle(node.key, state);

		if (node.value.type === 'FunctionExpression' && !node.value.id) {
			state = {
				...state,
				scope: state.scope_map.get(node.value)
			};

			const chunks = node.kind !== 'init' ? [c(`${node.kind} `)] : [];

			if (node.value.async) {
				chunks.push(c('async '));
			}
			if (node.value.generator) {
				chunks.push(c('*'));
			}

			push_array(chunks, node.computed ? [c('['), ...key, c(']')] : key);
			chunks.push(c('('));
			push_array(
				chunks,
				join(
					node.value.params.map((/** @type {Pattern} */ param) =>
						handle(param, state)
					),
					c(', ')
				)
			);
			chunks.push(c(') '));
			push_array(chunks, handle(node.value.body, state));

			return chunks;
		}

		if (node.computed) {
			return [c('['), ...key, c(']: '), ...value];
		}

		return [...key, c(': '), ...value];
	},

	ObjectPattern(node, state) {
		const chunks = [c('{ ')];

		for (let i = 0; i < node.properties.length; i += 1) {
			push_array(chunks, handle(node.properties[i], state));
			if (i < node.properties.length - 1) chunks.push(c(', '));
		}

		chunks.push(c(' }'));

		return chunks;
	},

	SequenceExpression(/** @type {SequenceExpression} */ node, state) {
		const expressions = node.expressions.map((e) => handle(e, state));

		return [c('('), ...join(expressions, c(', ')), c(')')];
	},

	UnaryExpression(node, state) {
		const chunks = [c(node.operator)];

		if (node.operator.length > 1) {
			chunks.push(c(' '));
		}

		if (
			EXPRESSIONS_PRECEDENCE[node.argument.type] <
			EXPRESSIONS_PRECEDENCE.UnaryExpression
		) {
			chunks.push(c('('));
			push_array(chunks, handle(node.argument, state));
			chunks.push(c(')'));
		} else {
			push_array(chunks, handle(node.argument, state));
		}

		return chunks;
	},

	UpdateExpression(node, state) {
		return node.prefix
			? [c(node.operator), ...handle(node.argument, state)]
			: [...handle(node.argument, state), c(node.operator)];
	},

	AssignmentExpression(node, state) {
		return [
			...handle(node.left, state),
			c(` ${node.operator || '='} `),
			...handle(node.right, state)
		];
	},

	BinaryExpression(node, state) {
		/**
		 * @type any[]
		 */
		const chunks = [];

		// TODO
		// const is_in = node.operator === 'in';
		// if (is_in) {
		// 	// Avoids confusion in `for` loops initializers
		// 	chunks.push(c('('));
		// }

		if (needs_parens(node.left, node, false)) {
			chunks.push(c('('));
			push_array(chunks, handle(node.left, state));
			chunks.push(c(')'));
		} else {
			push_array(chunks, handle(node.left, state));
		}

		chunks.push(c(` ${node.operator} `));

		if (needs_parens(node.right, node, true)) {
			chunks.push(c('('));
			push_array(chunks, handle(node.right, state));
			chunks.push(c(')'));
		} else {
			push_array(chunks, handle(node.right, state));
		}

		return chunks;
	},

	ConditionalExpression(node, state) {
		/**
		 * @type any[]
		 */
		const chunks = [];

		if (
			EXPRESSIONS_PRECEDENCE[node.test.type] >
			EXPRESSIONS_PRECEDENCE.ConditionalExpression
		) {
			push_array(chunks, handle(node.test, state));
		} else {
			chunks.push(c('('));
			push_array(chunks, handle(node.test, state));
			chunks.push(c(')'));
		}

		const child_state = { ...state, indent: state.indent + '\t' };

		const consequent = handle(node.consequent, child_state);
		const alternate = handle(node.alternate, child_state);

		const multiple_lines =
			has_newline(consequent) ||
			has_newline(alternate) ||
			get_length(chunks) + get_length(consequent) + get_length(alternate) > 50;

		if (multiple_lines) {
			chunks.push(c(`\n${state.indent}? `));
			push_array(chunks, consequent);
			chunks.push(c(`\n${state.indent}: `));
			push_array(chunks, alternate);
		} else {
			chunks.push(c(` ? `));
			push_array(chunks, consequent);
			chunks.push(c(` : `));
			push_array(chunks, alternate);
		}

		return chunks;
	},

	NewExpression(/** @type {NewExpression} */ node, state) {
		const chunks = [c('new ')];

		if (
			EXPRESSIONS_PRECEDENCE[node.callee.type] <
				EXPRESSIONS_PRECEDENCE.CallExpression ||
			has_call_expression(node.callee)
		) {
			chunks.push(c('('));
			push_array(chunks, handle(node.callee, state));
			chunks.push(c(')'));
		} else {
			push_array(chunks, handle(node.callee, state));
		}

		// TODO this is copied from CallExpression — DRY it out
		const args = node.arguments.map((arg) =>
			handle(arg, {
				...state,
				indent: state.indent + '\t'
			})
		);

		const separator = args.some(has_newline) // TODO or length exceeds 80
			? c(',\n' + state.indent)
			: c(', ');

		chunks.push(c('('));
		push_array(chunks, join(args, separator));
		chunks.push(c(')'));

		return chunks;
	},

	ChainExpression(node, state) {
		return handle(node.expression, state);
	},

	CallExpression(/** @type {CallExpression} */ node, state) {
		/**
		 * @type any[]
		 */
		const chunks = [];

		if (
			EXPRESSIONS_PRECEDENCE[node.callee.type] <
			EXPRESSIONS_PRECEDENCE.CallExpression
		) {
			chunks.push(c('('));
			push_array(chunks, handle(node.callee, state));
			chunks.push(c(')'));
		} else {
			push_array(chunks, handle(node.callee, state));
		}

		if (/** @type {SimpleCallExpression} */ (node).optional) {
			chunks.push(c('?.'));
		}

		let has_inline_comment = false;
		let arg_chunks = [];
		outer: for (const arg of node.arguments) {
			const chunks = [];
			while (state.comments.length) {
				const comment = state.comments.shift();
				if (comment.type === 'Line') {
					has_inline_comment = true;
					break outer;
				}
				chunks.push(
					c(
						comment.type === 'Block'
							? `/*${comment.value}*/ `
							: `//${comment.value}`
					)
				);
			}
			push_array(chunks, handle(arg, state));
			arg_chunks.push(chunks);
		}

		const multiple_lines =
			has_inline_comment || arg_chunks.slice(0, -1).some(has_newline); // TODO or length exceeds 80
		if (multiple_lines) {
			// need to handle args again. TODO find alternative approach?
			const args = node.arguments.map((arg, i) => {
				const chunks = handle(arg, {
					...state,
					indent: `${state.indent}\t`
				});
				if (i < node.arguments.length - 1) chunks.push(c(','));
				while (state.comments.length) {
					const comment = state.comments.shift();
					chunks.push(
						c(
							comment.type === 'Block'
								? ` /*${comment.value}*/ `
								: ` //${comment.value}`
						)
					);
				}
				return chunks;
			});

			chunks.push(c(`(\n${state.indent}\t`));
			push_array(chunks, join(args, c(`\n${state.indent}\t`)));
			chunks.push(c(`\n${state.indent})`));
		} else {
			chunks.push(c('('));
			push_array(chunks, join(arg_chunks, c(', ')));
			chunks.push(c(')'));
		}

		return chunks;
	},

	MemberExpression(node, state) {
		/**
		 * @type any[]
		 */
		const chunks = [];

		if (
			EXPRESSIONS_PRECEDENCE[node.object.type] <
			EXPRESSIONS_PRECEDENCE.MemberExpression
		) {
			chunks.push(c('('));
			push_array(chunks, handle(node.object, state));
			chunks.push(c(')'));
		} else {
			push_array(chunks, handle(node.object, state));
		}

		if (node.computed) {
			if (node.optional) {
				chunks.push(c('?.'));
			}
			chunks.push(c('['));
			push_array(chunks, handle(node.property, state));
			chunks.push(c(']'));
		} else {
			chunks.push(c(node.optional ? '?.' : '.'));
			push_array(chunks, handle(node.property, state));
		}

		return chunks;
	},

	MetaProperty(node, state) {
		return [
			...handle(node.meta, state),
			c('.'),
			...handle(node.property, state)
		];
	},

	Identifier(node, state) {
		let name = node.name;

		if (name[0] === '@') {
			name = state.getName(name.slice(1));
		} else if (node.name[0] === '#') {
			const owner = state.scope.find_owner(node.name);

			if (!owner) {
				throw new Error(`Could not find owner for node`);
			}

			if (!state.deconflicted.has(owner)) {
				state.deconflicted.set(owner, new Map());
			}

			const deconflict_map = state.deconflicted.get(owner);

			if (!deconflict_map.has(node.name)) {
				deconflict_map.set(
					node.name,
					deconflict(node.name.slice(1), owner.references)
				);
			}

			name = deconflict_map.get(node.name);
		}

		return [c(name, node)];
	},

	Literal(/** @type {Literal} */ node, state) {
		if (typeof node.value === 'string') {
			return [
				// TODO do we need to handle weird unicode characters somehow?
				// str.replace(/\\u(\d{4})/g, (m, n) => String.fromCharCode(+n))
				c(
					(node.raw || JSON.stringify(node.value)).replace(
						re,
						(_m, _i, at, hash, name) => {
							if (at) return '@' + name;
							if (hash) return '#' + name;
							throw new Error(`this shouldn't happen`);
						}
					),
					node
				)
			];
		}

		return [c(node.raw || String(node.value), node)];
	},

	PropertyDefinition(/** @type {PropertyDefinition} */ node, state) {
		const chunks = [];

		if (node.static) {
			chunks.push(c('static '));
		}

		if (node.computed) {
			chunks.push(c('['), ...handle(node.key, state), c(']'));
		} else {
			chunks.push(...handle(node.key, state));
		}

		if (node.value) {
			chunks.push(c(' = '));

			chunks.push(...handle(node.value, state));
		}

		chunks.push(c(';'));

		return chunks;
	},

	StaticBlock(/** @type {StaticBlock} */ node, state) {
		const chunks = [c('static ')];

		push_array(chunks, handlers.BlockStatement(node, state));

		return chunks;
	},

	PrivateIdentifier(/** @type {PrivateIdenifier} */ node, state) {
		const chunks = [c('#')];

		push_array(chunks, [c(node.name, node)]);

		return chunks;
	}
};

handlers.ForOfStatement = handlers.ForInStatement;
handlers.FunctionExpression = handlers.FunctionDeclaration;
handlers.ClassExpression = handlers.ClassDeclaration;
handlers.ClassBody = handlers.BlockStatement;
handlers.SpreadElement = handlers.RestElement;
handlers.ArrayPattern = handlers.ArrayExpression;
handlers.LogicalExpression = handlers.BinaryExpression;
handlers.AssignmentPattern = handlers.AssignmentExpression;
