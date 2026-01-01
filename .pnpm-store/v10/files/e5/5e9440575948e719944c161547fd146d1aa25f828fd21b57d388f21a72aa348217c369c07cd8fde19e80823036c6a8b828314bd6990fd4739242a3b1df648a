/** @import { Node } from 'estree' */

/**
 * @param {Node} node
 * @param {Node} parent
 * @returns {boolean}
 */
export default function is_reference(node, parent) {
	if (node.type === 'MemberExpression') {
		return !node.computed && is_reference(node.object, node);
	}

	if (node.type !== 'Identifier') return false;

	switch (parent?.type) {
		// disregard `bar` in `foo.bar`
		case 'MemberExpression':
			return parent.computed || node === parent.object;

		// disregard the `foo` in `class {foo(){}}` but keep it in `class {[foo](){}}`
		case 'MethodDefinition':
			return parent.computed;

		// disregard the `meta` in `import.meta`
		case 'MetaProperty':
			return parent.meta === node;

		// disregard the `foo` in `class {foo=bar}` but keep it in `class {[foo]=bar}` and `class {bar=foo}`
		case 'PropertyDefinition':
			return parent.computed || node === parent.value;

		// disregard the `bar` in `{ bar: foo }`, but keep it in `{ [bar]: foo }`
		case 'Property':
			return parent.computed || node === parent.value;

		// disregard the `bar` in `export { foo as bar }` or
		// the foo in `import { foo as bar }`
		case 'ExportSpecifier':
		case 'ImportSpecifier':
			return node === parent.local;

		// disregard the `foo` in `foo: while (...) { ... break foo; ... continue foo;}`
		case 'LabeledStatement':
		case 'BreakStatement':
		case 'ContinueStatement':
			return false;

		default:
			return true;
	}
}
