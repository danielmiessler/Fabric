import { tsPlugin } from '@sveltejs/acorn-typescript';
import { Parser } from 'acorn';
import { read } from '../../../utils/filesystem.js';

const inheritable_page_options = new Set(['ssr', 'prerender', 'csr', 'trailingSlash', 'config']);

const valid_page_options = new Set([...inheritable_page_options, 'entries', 'load']);

const skip_parsing_regex = new RegExp(
	`${Array.from(valid_page_options).join('|')}|(?:export[\\s\\n]+\\*[\\s\\n]+from)`
);

const parser = Parser.extend(tsPlugin());

/**
 * Collects page options from a +page.js/+layout.js file, ignoring reassignments
 * and using the declared value (except for load functions, for which the value is `true`).
 * Returns `null` if any export is too difficult to analyse.
 * @param {string} filename The name of the file to report when an error occurs
 * @param {string} input
 * @returns {Record<string, any> | null}
 */
export function statically_analyse_page_options(filename, input) {
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
		const page_options = new Map();

		for (const statement of source.body) {
			// ignore export all declarations with aliases that are not page options
			if (
				statement.type === 'ExportAllDeclaration' &&
				statement.exported &&
				!valid_page_options.has(get_name(statement.exported))
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
					if (!valid_page_options.has(exported_name)) {
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
									page_options.set(
										/** @type {string} */ (export_specifiers.get(variable_declarator.id.name)),
										variable_declarator.init.value
									);
									export_specifiers.delete(variable_declarator.id.name);
									continue;
								}

								// Special case: We only want to know that 'load' is exported (in a way that doesn't cause truthy checks in other places to trigger)
								if (variable_declarator.id.name === 'load') {
									page_options.set('load', null);
									export_specifiers.delete('load');
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
				if (valid_page_options.has(statement.declaration.id.name)) {
					// Special case: We only want to know that 'load' is exported (in a way that doesn't cause truthy checks in other places to trigger)
					if (statement.declaration.id.name === 'load') {
						page_options.set('load', null);
					} else {
						return null;
					}
				}
				continue;
			}

			for (const declaration of statement.declaration.declarations) {
				if (declaration.id.type !== 'Identifier') {
					return null;
				}

				if (!valid_page_options.has(declaration.id.name)) {
					continue;
				}

				if (declaration.init?.type === 'Literal') {
					page_options.set(declaration.id.name, declaration.init.value);
					continue;
				}

				// Special case: We only want to know that 'load' is exported (in a way that doesn't cause truthy checks in other places to trigger)
				if (declaration.id.name === 'load') {
					page_options.set('load', null);
					continue;
				}

				// references a declaration we can't easily evaluate statically
				return null;
			}
		}

		return Object.fromEntries(page_options);
	} catch (error) {
		if (error instanceof Error) {
			error.message = `Failed to statically analyse page options for ${filename}. ${error.message}`;
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
 *   static_exports?: Map<string, { page_options: Record<string, any> | null, children: string[] }>;
 * }} opts
 */
export function create_node_analyser({ resolve, static_exports = new Map() }) {
	/**
	 * Computes the final page options (may include load function as `load: null`; special case) for a node (if possible). Otherwise, returns `null`.
	 * @param {import('types').PageNode} node
	 * @returns {Promise<Record<string, any> | null>}
	 */
	const get_page_options = async (node) => {
		const key = node.universal || node.server;
		if (key && static_exports.has(key)) {
			return { ...static_exports.get(key)?.page_options };
		}

		/** @type {Record<string, any>} */
		let page_options = {};

		if (node.parent) {
			const parent_options = await get_page_options(node.parent);

			const parent_key = node.parent.universal || node.parent.server;
			if (key && parent_key) {
				static_exports.get(parent_key)?.children.push(key);
			}

			if (parent_options === null) {
				// if the parent cannot be analysed, we can't know what page options
				// the child node inherits, so we also mark it as unanalysable
				if (key) {
					static_exports.set(key, { page_options: null, children: [] });
				}
				return null;
			}

			page_options = { ...parent_options };
		}

		if (node.server) {
			const module = await resolve(node.server);
			for (const page_option in inheritable_page_options) {
				if (page_option in module) {
					page_options[page_option] = module[page_option];
				}
			}
		}

		if (node.universal) {
			const input = read(node.universal);
			const universal_page_options = statically_analyse_page_options(node.universal, input);

			if (universal_page_options === null) {
				static_exports.set(node.universal, { page_options: null, children: [] });
				return null;
			}

			page_options = { ...page_options, ...universal_page_options };
		}

		if (key) {
			static_exports.set(key, { page_options, children: [] });
		}

		return page_options;
	};

	/**
	 * @param {string} file
	 */
	const invalidate_page_options = (file) => {
		static_exports.get(file)?.children.forEach((child) => static_exports.delete(child));
		static_exports.delete(file);
	};

	return {
		get_page_options,
		invalidate_page_options
	};
}
