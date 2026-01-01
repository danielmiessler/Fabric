import { tsPlugin } from '@sveltejs/acorn-typescript';
import { Parser } from 'acorn';
import { read } from '../../../utils/filesystem.js';

const inheritable_page_options = new Set(['ssr', 'prerender', 'csr', 'trailingSlash', 'config']);

const page_options = new Set([...inheritable_page_options, 'entries']);

const skip_parsing_regex = new RegExp(
	`${Array.from(page_options).join('|')}|(?:export[\\s\\n]+\\*[\\s\\n]+from)`
);

const parser = Parser.extend(tsPlugin());

/**
 * Collects exported page options from a +page.js/+layout.js file.
 * We ignore reassignments and use the declared value.
 * Returns `null` if any export is too difficult to analyse.
 * @param {string} filename
 * @param {string} input
 * @returns {Record<string, any> | null}
 */
export function statically_analyse_exports(filename, input) {
	// if there's a chance there are no page exports or export all declaration,
	// then we can skip the AST parsing which is expensive
	if (!skip_parsing_regex.test(input)) {
		return {};
	}

	try {
		const source = parser.parse(input, {
			sourceType: 'module',
			ecmaVersion: 'latest'
		});

		/** @type {Map<string, import('acorn').Literal['value']>} */
		const static_exports = new Map();

		for (const statement of source.body) {
			// ignore export all declarations with aliases that are not page options
			if (
				statement.type === 'ExportAllDeclaration' &&
				statement.exported &&
				!page_options.has(get_name(statement.exported))
			) {
				continue;
			}

			if (
				statement.type === 'ExportDefaultDeclaration' ||
				statement.type === 'ExportAllDeclaration'
			) {
				return null;
			} else if (statement.type !== 'ExportNamedDeclaration') {
				continue;
			}

			if (statement.specifiers.length) {
				/** @type {Map<string, string>} */
				const export_specifiers = new Map();
				for (const specifier of statement.specifiers) {
					const exported_name = get_name(specifier.exported);
					if (!page_options.has(exported_name)) {
						continue;
					}

					if (statement.source) {
						return null;
					}

					export_specifiers.set(get_name(specifier.local), exported_name);
				}

				for (const statement of source.body) {
					switch (statement.type) {
						case 'ImportDeclaration': {
							for (const import_specifier of statement.specifiers) {
								if (export_specifiers.has(import_specifier.local.name)) {
									return null;
								}
							}
							break;
						}
						case 'ExportNamedDeclaration':
						case 'VariableDeclaration':
						case 'FunctionDeclaration':
						case 'ClassDeclaration': {
							const declaration =
								statement.type === 'ExportNamedDeclaration' ? statement.declaration : statement;

							if (!declaration) {
								break;
							}

							// class and function declarations
							if (declaration.type !== 'VariableDeclaration') {
								if (export_specifiers.has(declaration.id.name)) {
									return null;
								}
								break;
							}

							for (const variable_declarator of declaration.declarations) {
								if (
									variable_declarator.id.type !== 'Identifier' ||
									!export_specifiers.has(variable_declarator.id.name)
								) {
									continue;
								}

								if (variable_declarator.init?.type === 'Literal') {
									static_exports.set(
										/** @type {string} */ (export_specifiers.get(variable_declarator.id.name)),
										variable_declarator.init.value
									);
									export_specifiers.delete(variable_declarator.id.name);
									continue;
								}

								// references a declaration we can't easily evaluate statically
								return null;
							}
							break;
						}
					}
				}

				// there were some export specifiers that we couldn't resolve
				if (export_specifiers.size) {
					return null;
				}
				continue;
			}

			if (!statement.declaration) {
				continue;
			}

			// class and function declarations
			if (statement.declaration.type !== 'VariableDeclaration') {
				if (page_options.has(statement.declaration.id.name)) {
					return null;
				}
				continue;
			}

			for (const declaration of statement.declaration.declarations) {
				if (declaration.id.type !== 'Identifier') {
					return null;
				}

				if (!page_options.has(declaration.id.name)) {
					continue;
				}

				if (declaration.init?.type === 'Literal') {
					static_exports.set(declaration.id.name, declaration.init.value);
					continue;
				}

				// references a declaration we can't easily evaluate statically
				return null;
			}
		}

		return Object.fromEntries(static_exports);
	} catch (error) {
		if (error instanceof Error) {
			error.message = `Failed to statically analyse ${filename}. ${error.message}`;
		}
		throw error;
	}
}

/**
 * @param {import('acorn').Identifier | import('acorn').Literal} node
 * @returns {string}
 */
export function get_name(node) {
	return node.type === 'Identifier' ? node.name : /** @type {string} */ (node.value);
}

/**
 * @param {{
 *   resolve: (file: string) => Promise<Record<string, any>>;
 *   static_exports?: Map<string, Record<string, any> | null>;
 * }} opts
 */
export function create_node_analyser({ resolve, static_exports = new Map() }) {
	/**
	 * Computes the final page options for a node (if possible). Otherwise, returns `null`.
	 * @param {import('types').PageNode} node
	 * @returns {Promise<import('types').UniversalNode | null>}
	 */
	const get_page_options = async (node) => {
		if (node.universal && static_exports.has(node.universal)) {
			return /** @type {import('types').UniversalNode | null} */ (
				static_exports.get(node.universal)
			);
		}

		/** @type {Record<string, any> | null} */
		let page_options = {};

		if (node.server) {
			const module = await resolve(node.server);
			for (const key in inheritable_page_options) {
				if (key in module) {
					page_options[key] = module[key];
				}
			}
		}

		if (node.universal) {
			let universal_exports = static_exports.get(node.universal);
			if (universal_exports === undefined) {
				const input = read(node.universal);
				universal_exports = statically_analyse_exports(node.universal, input);
			}

			if (universal_exports === null) {
				static_exports.set(node.universal, null);
				return null;
			}

			page_options = { ...page_options, ...universal_exports };
		}

		if (node.parent) {
			const parent_options = await get_page_options(node.parent);
			if (parent_options === null) {
				// if the parent cannot be statically analysed, we can't know what
				// page options the current node inherits, so we invalidate it too
				if (node.universal) {
					static_exports.set(node.universal, null);
				}
				return null;
			}

			page_options = { ...parent_options, ...page_options };
		}

		if (node.universal) {
			static_exports.set(node.universal, page_options);
		}

		return page_options;
	};

	/**
	 * @param {string} file
	 * @returns {void}
	 */
	const invalidate_page_options = (file) => {
		static_exports.delete(file);
	};

	return { get_page_options, invalidate_page_options };
}
