import fs from 'node:fs';
import path from 'node:path';
import process from 'node:process';
import colors from 'kleur';
import { lookup } from 'mrmime';
import { list_files, runtime_directory } from '../../utils.js';
import { posixify, resolve_entry } from '../../../utils/filesystem.js';
import { parse_route_id } from '../../../utils/routing.js';
import { sort_routes } from './sort.js';
import { isSvelte5Plus } from '../utils.js';

/**
 * Generates the manifest data used for the client-side manifest and types generation.
 * @param {{
 *   config: import('types').ValidatedConfig;
 *   fallback?: string;
 *   cwd?: string;
 * }} opts
 * @returns {import('types').ManifestData}
 */
export default function create_manifest_data({
	config,
	fallback = `${runtime_directory}/components/${isSvelte5Plus() ? 'svelte-5' : 'svelte-4'}`,
	cwd = process.cwd()
}) {
	const assets = create_assets(config);
	const hooks = create_hooks(config, cwd);
	const matchers = create_matchers(config, cwd);
	const { nodes, routes } = create_routes_and_nodes(cwd, config, fallback);

	for (const route of routes) {
		for (const param of route.params) {
			if (param.matcher && !matchers[param.matcher]) {
				throw new Error(`No matcher found for parameter '${param.matcher}' in route ${route.id}`);
			}
		}
	}

	return {
		assets,
		hooks,
		matchers,
		nodes,
		routes
	};
}

/**
 * @param {import('types').ValidatedConfig} config
 */
export function create_assets(config) {
	return list_files(config.kit.files.assets).map((file) => ({
		file,
		size: fs.statSync(path.resolve(config.kit.files.assets, file)).size,
		type: lookup(file) || null
	}));
}

/**
 * @param {import('types').ValidatedConfig} config
 * @param {string} cwd
 */
function create_hooks(config, cwd) {
	const client = resolve_entry(config.kit.files.hooks.client);
	const server = resolve_entry(config.kit.files.hooks.server);
	const universal = resolve_entry(config.kit.files.hooks.universal);

	return {
		client: client && posixify(path.relative(cwd, client)),
		server: server && posixify(path.relative(cwd, server)),
		universal: universal && posixify(path.relative(cwd, universal))
	};
}

/**
 * @param {import('types').ValidatedConfig} config
 * @param {string} cwd
 */
function create_matchers(config, cwd) {
	const params_base = path.relative(cwd, config.kit.files.params);

	/** @type {Record<string, string>} */
	const matchers = {};
	if (fs.existsSync(config.kit.files.params)) {
		for (const file of fs.readdirSync(config.kit.files.params)) {
			const ext = path.extname(file);
			if (!config.kit.moduleExtensions.includes(ext)) continue;
			const type = file.slice(0, -ext.length);

			if (/^\w+$/.test(type)) {
				const matcher_file = path.join(params_base, file);

				// Disallow same matcher with different extensions
				if (matchers[type]) {
					throw new Error(`Duplicate matchers: ${matcher_file} and ${matchers[type]}`);
				} else {
					matchers[type] = matcher_file;
				}
			} else {
				// Allow for matcher test collocation
				if (type.endsWith('.test') || type.endsWith('.spec')) continue;

				throw new Error(
					`Matcher names can only have underscores and alphanumeric characters — "${file}" is invalid`
				);
			}
		}
	}

	return matchers;
}

/**
 * @param {import('types').ValidatedConfig} config
 * @param {string} cwd
 * @param {string} fallback
 */
function create_routes_and_nodes(cwd, config, fallback) {
	/** @type {import('types').RouteData[]} */
	const routes = [];

	const routes_base = posixify(path.relative(cwd, config.kit.files.routes));

	const valid_extensions = [...config.extensions, ...config.kit.moduleExtensions];

	/** @type {import('types').PageNode[]} */
	const nodes = [];

	if (fs.existsSync(config.kit.files.routes)) {
		/**
		 * @param {number} depth
		 * @param {string} id
		 * @param {string} segment
		 * @param {import('types').RouteData | null} parent
		 */
		const walk = (depth, id, segment, parent) => {
			const unescaped = id.replace(/\[([ux])\+([^\]]+)\]/gi, (match, type, code) => {
				if (match !== match.toLowerCase()) {
					throw new Error(`Character escape sequence in ${id} must be lowercase`);
				}

				if (!/[0-9a-f]+/.test(code)) {
					throw new Error(`Invalid character escape sequence in ${id}`);
				}

				if (type === 'x') {
					if (code.length !== 2) {
						throw new Error(`Hexadecimal escape sequence in ${id} must be two characters`);
					}

					return String.fromCharCode(parseInt(code, 16));
				} else {
					if (code.length < 4 || code.length > 6) {
						throw new Error(
							`Unicode escape sequence in ${id} must be between four and six characters`
						);
					}

					return String.fromCharCode(parseInt(code, 16));
				}
			});

			if (/\]\[/.test(unescaped)) {
				throw new Error(`Invalid route ${id} — parameters must be separated`);
			}

			if (count_occurrences('[', id) !== count_occurrences(']', id)) {
				throw new Error(`Invalid route ${id} — brackets are unbalanced`);
			}

			if (/#/.test(segment)) {
				// Vite will barf on files with # in them
				throw new Error(`Route ${id} should be renamed to ${id.replace(/#/g, '[x+23]')}`);
			}

			if (/\[\.\.\.\w+\]\/\[\[/.test(id)) {
				throw new Error(
					`Invalid route ${id} — an [[optional]] route segment cannot follow a [...rest] route segment`
				);
			}

			if (/\[\[\.\.\./.test(id)) {
				throw new Error(
					`Invalid route ${id} — a rest route segment is always optional, remove the outer square brackets`
				);
			}

			const { pattern, params } = parse_route_id(id);

			/** @type {import('types').RouteData} */
			const route = {
				id,
				parent,

				segment,
				pattern,
				params,

				layout: null,
				error: null,
				leaf: null,
				page: null,
				endpoint: null
			};

			// important to do this before walking children, so that child
			// routes appear later
			routes.push(route);

			// if we don't do this, the route map becomes unwieldy to console.log
			Object.defineProperty(route, 'parent', { enumerable: false });

			const dir = path.join(cwd, routes_base, id);

			// We can't use withFileTypes because of a NodeJs bug which returns wrong results
			// with isDirectory() in case of symlinks: https://github.com/nodejs/node/issues/30646
			const files = fs.readdirSync(dir).map((name) => ({
				is_dir: fs.statSync(path.join(dir, name)).isDirectory(),
				name
			}));

			// process files first
			for (const file of files) {
				if (file.is_dir) continue;

				const ext = valid_extensions.find((ext) => file.name.endsWith(ext));
				if (!ext) continue;

				if (!file.name.startsWith('+')) {
					const name = file.name.slice(0, -ext.length);
					// check if it is a valid route filename but missing the + prefix
					const typo =
						/^(?:(page(?:@(.*))?)|(layout(?:@(.*))?)|(error))$/.test(name) ||
						/^(?:(server)|(page(?:(@[a-zA-Z0-9_-]*))?(\.server)?)|(layout(?:(@[a-zA-Z0-9_-]*))?(\.server)?))$/.test(
							name
						);
					if (typo) {
						console.log(
							colors
								.bold()
								.yellow(
									`Missing route file prefix. Did you mean +${file.name}?` +
										` at ${path.join(dir, file.name)}`
								)
						);
					}

					continue;
				}

				if (file.name.endsWith('.d.ts')) {
					let name = file.name.slice(0, -5);
					const ext = valid_extensions.find((ext) => name.endsWith(ext));
					if (ext) name = name.slice(0, -ext.length);

					const valid =
						/^\+(?:(page(?:@(.*))?)|(layout(?:@(.*))?)|(error))$/.test(name) ||
						/^\+(?:(server)|(page(?:(@[a-zA-Z0-9_-]*))?(\.server)?)|(layout(?:(@[a-zA-Z0-9_-]*))?(\.server)?))$/.test(
							name
						);

					if (valid) continue;
				}

				const project_relative = posixify(path.relative(cwd, path.join(dir, file.name)));

				const item = analyze(
					project_relative,
					file.name,
					config.extensions,
					config.kit.moduleExtensions
				);

				if (config.kit.router.type === 'hash' && item.kind === 'server') {
					throw new Error(
						`Cannot use server-only files in an app with \`router.type === 'hash': ${project_relative}`
					);
				}

				/**
				 * @param {string} type
				 * @param {string} existing_file
				 */
				function duplicate_files_error(type, existing_file) {
					return new Error(
						`Multiple ${type} files found in ${routes_base}${route.id} : ${path.basename(
							existing_file
						)} and ${file.name}`
					);
				}

				if (item.kind === 'component') {
					if (item.is_error) {
						route.error = {
							depth,
							component: project_relative
						};
					} else if (item.is_layout) {
						if (!route.layout) {
							route.layout = { depth, child_pages: [] };
						} else if (route.layout.component) {
							throw duplicate_files_error('layout component', route.layout.component);
						}

						route.layout.component = project_relative;
						if (item.uses_layout !== undefined) route.layout.parent_id = item.uses_layout;
					} else {
						if (!route.leaf) {
							route.leaf = { depth };
						} else if (route.leaf.component) {
							throw duplicate_files_error('page component', route.leaf.component);
						}

						route.leaf.component = project_relative;
						if (item.uses_layout !== undefined) route.leaf.parent_id = item.uses_layout;
					}
				} else if (item.is_layout) {
					if (!route.layout) {
						route.layout = { depth, child_pages: [] };
					} else if (route.layout[item.kind]) {
						throw duplicate_files_error(
							item.kind + ' layout module',
							/** @type {string} */ (route.layout[item.kind])
						);
					}

					route.layout[item.kind] = project_relative;
				} else if (item.is_page) {
					if (!route.leaf) {
						route.leaf = { depth };
					} else if (route.leaf[item.kind]) {
						throw duplicate_files_error(
							item.kind + ' page module',
							/** @type {string} */ (route.leaf[item.kind])
						);
					}

					route.leaf[item.kind] = project_relative;
				} else {
					if (route.endpoint) {
						throw duplicate_files_error('endpoint', route.endpoint.file);
					}

					route.endpoint = {
						file: project_relative
					};
				}
			}

			// then handle children
			for (const file of files) {
				if (file.is_dir) {
					walk(depth + 1, path.posix.join(id, file.name), file.name, route);
				}
			}
		};

		walk(0, '/', '', null);

		if (routes.length === 1) {
			const root = routes[0];
			if (!root.leaf && !root.error && !root.layout && !root.endpoint) {
				throw new Error(
					'No routes found. If you are using a custom src/routes directory, make sure it is specified in svelte.config.js'
				);
			}
		}
	} else {
		// If there's no routes directory, we'll just create a single empty route. This ensures the root layout and
		// error components are included in the manifest, which is needed for subsequent build/dev commands to work
		routes.push({
			id: '/',
			segment: '',
			pattern: /^$/,
			params: [],
			parent: null,
			layout: null,
			error: null,
			leaf: null,
			page: null,
			endpoint: null
		});
	}

	prevent_conflicts(routes);

	const root = routes[0];

	if (!root.layout?.component) {
		if (!root.layout) root.layout = { depth: 0, child_pages: [] };
		root.layout.component = posixify(path.relative(cwd, `${fallback}/layout.svelte`));
	}

	if (!root.error?.component) {
		if (!root.error) root.error = { depth: 0 };
		root.error.component = posixify(path.relative(cwd, `${fallback}/error.svelte`));
	}

	// we do layouts/errors first as they are more likely to be reused,
	// and smaller indexes take fewer bytes. also, this guarantees that
	// the default error/layout are 0/1
	for (const route of routes) {
		if (route.layout) {
			if (!route.layout?.component) {
				route.layout.component = posixify(path.relative(cwd, `${fallback}/layout.svelte`));
			}
			nodes.push(route.layout);
		}
		if (route.error) nodes.push(route.error);
	}

	for (const route of routes) {
		if (route.leaf) nodes.push(route.leaf);
	}

	const indexes = new Map(nodes.map((node, i) => [node, i]));

	for (const route of routes) {
		if (!route.leaf) continue;

		route.page = {
			layouts: [],
			errors: [],
			leaf: /** @type {number} */ (indexes.get(route.leaf))
		};

		/** @type {import('types').RouteData | null} */
		let current_route = route;
		let current_node = route.leaf;
		let parent_id = route.leaf.parent_id;

		while (current_route) {
			if (parent_id === undefined || current_route.segment === parent_id) {
				if (current_route.layout || current_route.error) {
					route.page.layouts.unshift(
						current_route.layout ? indexes.get(current_route.layout) : undefined
					);
					route.page.errors.unshift(
						current_route.error ? indexes.get(current_route.error) : undefined
					);
				}

				if (current_route.layout) {
					/** @type {import('types').PageNode[]} */ (current_route.layout.child_pages).push(
						route.leaf
					);
					current_node.parent = current_node = current_route.layout;
					parent_id = current_node.parent_id;
				} else {
					parent_id = undefined;
				}
			}

			current_route = current_route.parent;
		}

		if (parent_id !== undefined) {
			throw new Error(`${current_node.component} references missing segment "${parent_id}"`);
		}
	}

	return {
		nodes,
		routes: sort_routes(routes)
	};
}

/**
 * @param {string} project_relative
 * @param {string} file
 * @param {string[]} component_extensions
 * @param {string[]} module_extensions
 * @returns {import('./types.js').RouteFile}
 */
function analyze(project_relative, file, component_extensions, module_extensions) {
	const component_extension = component_extensions.find((ext) => file.endsWith(ext));
	if (component_extension) {
		const name = file.slice(0, -component_extension.length);
		const pattern = /^\+(?:(page(?:@(.*))?)|(layout(?:@(.*))?)|(error))$/;
		const match = pattern.exec(name);
		if (!match) {
			throw new Error(`Files prefixed with + are reserved (saw ${project_relative})`);
		}

		return {
			kind: 'component',
			is_page: !!match[1],
			is_layout: !!match[3],
			is_error: !!match[5],
			uses_layout: match[2] ?? match[4]
		};
	}

	const module_extension = module_extensions.find((ext) => file.endsWith(ext));
	if (module_extension) {
		const name = file.slice(0, -module_extension.length);
		const pattern =
			/^\+(?:(server)|(page(?:(@[a-zA-Z0-9_-]*))?(\.server)?)|(layout(?:(@[a-zA-Z0-9_-]*))?(\.server)?))$/;
		const match = pattern.exec(name);
		if (!match) {
			throw new Error(`Files prefixed with + are reserved (saw ${project_relative})`);
		} else if (match[3] || match[6]) {
			throw new Error(
				// prettier-ignore
				`Only Svelte files can reference named layouts. Remove '${match[3] || match[6]}' from ${file} (at ${project_relative})`
			);
		}

		const kind = match[1] || match[4] || match[7] ? 'server' : 'universal';

		return {
			kind,
			is_page: !!match[2],
			is_layout: !!match[5]
		};
	}

	throw new Error(`Files and directories prefixed with + are reserved (saw ${project_relative})`);
}

/**
 * @param {string} needle
 * @param {string} haystack
 */
function count_occurrences(needle, haystack) {
	let count = 0;
	for (let i = 0; i < haystack.length; i += 1) {
		if (haystack[i] === needle) count += 1;
	}
	return count;
}

/** @param {import('types').RouteData[]} routes */
function prevent_conflicts(routes) {
	/** @type {Map<string, string>} */
	const lookup = new Map();

	for (const route of routes) {
		if (!route.leaf && !route.endpoint) continue;

		const normalized = normalize_route_id(route.id);

		// find all permutations created by optional parameters
		const split = normalized.split(/<\?(.+?)>/g);

		let permutations = [/** @type {string} */ (split[0])];

		// turn `x/[[optional]]/y` into `x/y` and `x/[required]/y`
		for (let i = 1; i < split.length; i += 2) {
			const matcher = split[i];
			const next = split[i + 1];

			permutations = permutations.reduce((a, b) => {
				a.push(b + next);
				if (!(matcher === '*' && b.endsWith('//'))) a.push(b + `<${matcher}>${next}`);
				return a;
			}, /** @type {string[]} */ ([]));
		}

		for (const permutation of permutations) {
			// remove leading/trailing/duplicated slashes caused by prior
			// manipulation of optional parameters and (groups)
			const key = permutation
				.replace(/\/{2,}/, '/')
				.replace(/^\//, '')
				.replace(/\/$/, '');

			if (lookup.has(key)) {
				throw new Error(
					`The "${lookup.get(key)}" and "${route.id}" routes conflict with each other`
				);
			}

			lookup.set(key, route.id);
		}
	}
}

/** @param {string} id */
function normalize_route_id(id) {
	return (
		id
			// remove groups
			.replace(/(?<=^|\/)\(.+?\)(?=$|\/)/g, '')

			.replace(/\[[ux]\+([0-9a-f]+)\]/g, (_, x) =>
				String.fromCharCode(parseInt(x, 16)).replace(/\//g, '%2f')
			)

			// replace `[param]` with `<*>`, `[param=x]` with `<x>`, and `[[param]]` with `<?*>`
			.replace(
				/\[(?:(\[)|(\.\.\.))?.+?(=.+?)?\]\]?/g,
				(_, optional, rest, matcher) => `<${optional ? '?' : ''}${rest ?? ''}${matcher ?? '*'}>`
			)
	);
}
