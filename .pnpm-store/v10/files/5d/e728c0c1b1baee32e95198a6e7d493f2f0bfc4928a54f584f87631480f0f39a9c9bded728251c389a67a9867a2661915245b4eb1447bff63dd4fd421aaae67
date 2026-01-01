import fs from 'node:fs';
import path from 'node:path';
import process from 'node:process';
import MagicString from 'magic-string';
import { posixify, rimraf, walk } from '../../../utils/filesystem.js';
import { compact } from '../../../utils/array.js';
import { ts } from '../ts.js';

/**
 *  @typedef {{
 *   file_name: string;
 *   modified: boolean;
 *   code: string;
 *   exports: any[];
 *  } | null} Proxy
 *
 *  @typedef {{
 *   server: Proxy,
 *   universal: Proxy
 *  }} Proxies
 *
 *  @typedef {Map<import('types').PageNode, {route: import('types').RouteData, proxies: Proxies}>} RoutesMap
 */

const cwd = process.cwd();

/**
 * Creates types for the whole manifest
 * @param {import('types').ValidatedConfig} config
 * @param {import('types').ManifestData} manifest_data
 */
export function write_all_types(config, manifest_data) {
	if (!ts) return;

	const types_dir = `${config.kit.outDir}/types`;

	// empty out files that no longer need to exist
	const routes_dir = posixify(path.relative('.', config.kit.files.routes)).replace(/\.\.\//g, '');
	const expected_directories = new Set(
		manifest_data.routes.map((route) => path.join(routes_dir, route.id))
	);

	if (fs.existsSync(types_dir)) {
		for (const file of walk(types_dir)) {
			const dir = path.dirname(file);
			if (!expected_directories.has(dir)) {
				rimraf(path.join(types_dir, file));
			}
		}
	}

	// Read/write meta data on each invocation, not once per node process,
	// it could be invoked by another process in the meantime.
	const meta_data_file = `${types_dir}/route_meta_data.json`;
	const has_meta_data = fs.existsSync(meta_data_file);
	const meta_data = has_meta_data
		? /** @type {Record<string, string[]>} */ (JSON.parse(fs.readFileSync(meta_data_file, 'utf-8')))
		: {};
	const routes_map = create_routes_map(manifest_data);
	// For each directory, write $types.d.ts
	for (const route of manifest_data.routes) {
		if (!route.leaf && !route.layout && !route.endpoint) continue; // nothing to do

		const outdir = path.join(config.kit.outDir, 'types', routes_dir, route.id);

		// check if the types are out of date
		/** @type {string[]} */
		const input_files = [];

		/** @type {import('types').PageNode | null} */
		let node = route.leaf;
		while (node) {
			if (node.universal) input_files.push(node.universal);
			if (node.server) input_files.push(node.server);
			node = node.parent ?? null;
		}

		/** @type {import('types').PageNode | null} */
		node = route.layout;
		while (node) {
			if (node.universal) input_files.push(node.universal);
			if (node.server) input_files.push(node.server);
			node = node.parent ?? null;
		}

		if (route.endpoint) {
			input_files.push(route.endpoint.file);
		}

		try {
			fs.mkdirSync(outdir, { recursive: true });
		} catch {}

		const output_files = compact(
			fs.readdirSync(outdir).map((name) => {
				const stats = fs.statSync(path.join(outdir, name));
				if (stats.isDirectory()) return;
				return {
					name,
					updated: stats.mtimeMs
				};
			})
		);

		const source_last_updated = Math.max(
			// ctimeMs includes move operations whereas mtimeMs does not
			...input_files.map((file) => fs.statSync(file).ctimeMs)
		);
		const types_last_updated = Math.max(...output_files.map((file) => file.updated));

		const should_generate =
			// source files were generated more recently than the types
			source_last_updated > types_last_updated ||
			// no meta data file exists yet
			!has_meta_data ||
			// some file was deleted
			!meta_data[route.id]?.every((file) => input_files.includes(file));

		if (should_generate) {
			// track which old files end up being surplus to requirements
			const to_delete = new Set(output_files.map((file) => file.name));
			update_types(config, routes_map, route, to_delete);
			meta_data[route.id] = input_files;
		}
	}

	fs.writeFileSync(meta_data_file, JSON.stringify(meta_data, null, '\t'));
}

/**
 * Creates types related to the given file. This should only be called
 * if the file in question was edited, not if it was created/deleted/moved.
 * @param {import('types').ValidatedConfig} config
 * @param {import('types').ManifestData} manifest_data
 * @param {string} file
 */
export function write_types(config, manifest_data, file) {
	if (!ts) return;

	if (!path.basename(file).startsWith('+')) {
		// Not a route file
		return;
	}

	const id = '/' + posixify(path.relative(config.kit.files.routes, path.dirname(file)));

	const route = manifest_data.routes.find((route) => route.id === id);
	if (!route) return;
	if (!route.leaf && !route.layout && !route.endpoint) return; // nothing to do

	update_types(config, create_routes_map(manifest_data), route);
}

/**
 * Collect all leafs into a leaf -> route map
 * @param {import('types').ManifestData} manifest_data
 */
function create_routes_map(manifest_data) {
	/** @type {RoutesMap} */
	const map = new Map();
	for (const route of manifest_data.routes) {
		if (route.leaf) {
			map.set(route.leaf, { route, proxies: { server: null, universal: null } });
		}
	}
	return map;
}

/**
 * Update types for a specific route
 * @param {import('types').ValidatedConfig} config
 * @param {RoutesMap} routes
 * @param {import('types').RouteData} route
 * @param {Set<string>} [to_delete]
 */
function update_types(config, routes, route, to_delete = new Set()) {
	const routes_dir = posixify(path.relative('.', config.kit.files.routes)).replace(/\.\.\//g, '');
	const outdir = path.join(config.kit.outDir, 'types', routes_dir, route.id);

	// now generate new types
	const imports = ["import type * as Kit from '@sveltejs/kit';"];

	/** @type {string[]} */
	const declarations = [];

	/** @type {string[]} */
	const exports = [];

	// add 'Expand' helper
	// Makes sure a type is "repackaged" and therefore more readable
	declarations.push('type Expand<T> = T extends infer O ? { [K in keyof O]: O[K] } : never;');

	// returns the predicate of a matcher's type guard - or string if there is no type guard
	declarations.push(
		// TS complains on infer U, which seems weird, therefore ts-ignore it
		[
			'// @ts-ignore',
			'type MatcherParam<M> = M extends (param : string) => param is infer U ? U extends string ? U : string : string;'
		].join('\n')
	);

	declarations.push(
		'type RouteParams = ' + generate_params_type(route.params, outdir, config) + ';'
	);

	if (route.params.length > 0) {
		exports.push(
			'export type EntryGenerator = () => Promise<Array<RouteParams>> | Array<RouteParams>;'
		);
	}

	declarations.push(`type RouteId = '${route.id}';`);

	// These could also be placed in our public types, but it would bloat them unnecessarily and we may want to change these in the future
	if (route.layout || route.leaf) {
		declarations.push(
			// If T extends the empty object, void is also allowed as a return type
			'type MaybeWithVoid<T> = {} extends T ? T | void : T;',

			// Returns the key of the object whose values are required.
			'export type RequiredKeys<T> = { [K in keyof T]-?: {} extends { [P in K]: T[K] } ? never : K; }[keyof T];',

			// Helper type to get the correct output type for load functions. It should be passed the parent type to check what types from App.PageData are still required.
			// If none, void is also allowed as a return type.
			'type OutputDataShape<T> = MaybeWithVoid<Omit<App.PageData, RequiredKeys<T>> & Partial<Pick<App.PageData, keyof T & keyof App.PageData>> & Record<string, any>>',

			// null & {} == null, we need to prevent that in some situations
			'type EnsureDefined<T> = T extends null | undefined ? {} : T;',

			// Takes a union type and returns a union type where each type also has all properties
			// of all possible types (typed as undefined), making accessing them more ergonomic
			'type OptionalUnion<U extends Record<string, any>, A extends keyof U = U extends U ? keyof U : never> = U extends unknown ? { [P in Exclude<A, keyof U>]?: never } & U : never;',

			// Re-export `Snapshot` from @sveltejs/kit — in future we could use this to infer <T> from the return type of `snapshot.capture`
			'export type Snapshot<T = any> = Kit.Snapshot<T>;'
		);
	}

	if (route.leaf) {
		let route_info = routes.get(route.leaf);
		if (!route_info) {
			// This should be defined, but belts and braces
			route_info = { route, proxies: { server: null, universal: null } };
			routes.set(route.leaf, route_info);
		}

		const {
			declarations: d,
			exports: e,
			proxies
		} = process_node(route.leaf, outdir, true, route_info.proxies);

		exports.push(...e);
		declarations.push(...d);

		if (proxies.server) {
			route_info.proxies.server = proxies.server;
			if (proxies.server?.modified) to_delete.delete(proxies.server.file_name);
		}
		if (proxies.universal) {
			route_info.proxies.universal = proxies.universal;
			if (proxies.universal?.modified) to_delete.delete(proxies.universal.file_name);
		}

		if (route.leaf.server) {
			exports.push(
				'export type Action<OutputData extends Record<string, any> | void = Record<string, any> | void> = Kit.Action<RouteParams, OutputData, RouteId>'
			);
			exports.push(
				'export type Actions<OutputData extends Record<string, any> | void = Record<string, any> | void> = Kit.Actions<RouteParams, OutputData, RouteId>'
			);
		}

		if (route.leaf.server) {
			exports.push('export type PageProps = { data: PageData; form: ActionData }');
		} else {
			exports.push('export type PageProps = { data: PageData }');
		}
	}

	if (route.layout) {
		let all_pages_have_load = true;
		/** @type {import('types').RouteParam[]} */
		const layout_params = [];
		const ids = ['RouteId'];

		route.layout.child_pages?.forEach((page) => {
			const leaf = routes.get(page);
			if (leaf) {
				if (leaf.route.page) ids.push(`"${leaf.route.id}"`);

				for (const param of leaf.route.params) {
					// skip if already added
					if (layout_params.some((p) => p.name === param.name)) continue;
					layout_params.push({ ...param, optional: true });
				}

				ensureProxies(page, leaf.proxies);

				if (
					// Be defensive - if a proxy doesn't exist (because it couldn't be created), assume a load function exists.
					// If we didn't and it's a false negative, the user could wrongfully get a type error on layouts.
					(leaf.proxies.server && !leaf.proxies.server.exports.includes('load')) ||
					(leaf.proxies.universal && !leaf.proxies.universal.exports.includes('load'))
				) {
					all_pages_have_load = false;
				}
			}
			if (!page.server && !page.universal) {
				all_pages_have_load = false;
			}
		});

		if (route.id === '/') {
			// root layout is used for fallback error page, where ID can be null
			ids.push('null');
		}

		declarations.push(`type LayoutRouteId = ${ids.join(' | ')}`);

		declarations.push(
			'type LayoutParams = RouteParams & ' + generate_params_type(layout_params, outdir, config)
		);

		const {
			exports: e,
			declarations: d,
			proxies
		} = process_node(
			route.layout,
			outdir,
			false,
			{ server: null, universal: null },
			all_pages_have_load
		);

		exports.push(...e);
		declarations.push(...d);

		if (proxies.server?.modified) to_delete.delete(proxies.server.file_name);
		if (proxies.universal?.modified) to_delete.delete(proxies.universal.file_name);

		exports.push(
			'export type LayoutProps = { data: LayoutData; children: import("svelte").Snippet }'
		);
	}

	if (route.endpoint) {
		exports.push('export type RequestHandler = Kit.RequestHandler<RouteParams, RouteId>;');
	}

	if (route.leaf?.server || route.layout?.server || route.endpoint) {
		exports.push('export type RequestEvent = Kit.RequestEvent<RouteParams, RouteId>;');
	}

	const output = [imports.join('\n'), declarations.join('\n'), exports.join('\n')]
		.filter(Boolean)
		.join('\n\n');

	fs.writeFileSync(`${outdir}/$types.d.ts`, output);
	to_delete.delete('$types.d.ts');

	for (const file of to_delete) {
		fs.unlinkSync(path.join(outdir, file));
	}
}

/**
 * @param {import('types').PageNode} node
 * @param {string} outdir
 * @param {boolean} is_page
 * @param {Proxies} proxies
 * @param {boolean} [all_pages_have_load]
 */
function process_node(node, outdir, is_page, proxies, all_pages_have_load = true) {
	const params = `${is_page ? 'Route' : 'Layout'}Params`;
	const prefix = is_page ? 'Page' : 'Layout';

	const route_id = is_page ? 'RouteId' : 'LayoutRouteId';

	/** @type {string[]} */
	const declarations = [];
	/** @type {string[]} */
	const exports = [];

	/** @type {string} */
	let server_data;
	/** @type {string} */
	let data;

	ensureProxies(node, proxies);

	if (node.server) {
		const basename = path.basename(node.server);
		const proxy = proxies.server;
		if (proxy?.modified) {
			fs.writeFileSync(`${outdir}/proxy${basename}`, proxy.code);
		}

		server_data = get_data_type(node.server, 'null', proxy, true);

		const parent_type = `${prefix}ServerParentData`;

		declarations.push(`type ${parent_type} = ${get_parent_type(node, 'LayoutServerData')};`);

		// +page.js load present -> server can return all-optional data
		const output_data_shape =
			node.universal || (!is_page && all_pages_have_load)
				? 'Partial<App.PageData> & Record<string, any> | void'
				: `OutputDataShape<${parent_type}>`;
		exports.push(
			`export type ${prefix}ServerLoad<OutputData extends ${output_data_shape} = ${output_data_shape}> = Kit.ServerLoad<${params}, ${parent_type}, OutputData, ${route_id}>;`
		);

		exports.push(`export type ${prefix}ServerLoadEvent = Parameters<${prefix}ServerLoad>[0];`);

		if (is_page) {
			let type = 'unknown';
			if (proxy && proxy.exports.includes('actions')) {
				// If the file wasn't tweaked, we can use the return type of the original file.
				// The advantage is that type updates are reflected without saving.
				const from = proxy.modified
					? `./proxy${replace_ext_with_js(basename)}`
					: path_to_original(outdir, node.server);

				exports.push(
					'type ExcludeActionFailure<T> = T extends Kit.ActionFailure<any> ? never : T extends void ? never : T;',
					'type ActionsSuccess<T extends Record<string, (...args: any) => any>> = { [Key in keyof T]: ExcludeActionFailure<Awaited<ReturnType<T[Key]>>>; }[keyof T];',
					'type ExtractActionFailure<T> = T extends Kit.ActionFailure<infer X>	? X extends void ? never : X : never;',
					'type ActionsFailure<T extends Record<string, (...args: any) => any>> = { [Key in keyof T]: Exclude<ExtractActionFailure<Awaited<ReturnType<T[Key]>>>, void>; }[keyof T];',
					`type ActionsExport = typeof import('${from}').actions`,
					'export type SubmitFunction = Kit.SubmitFunction<Expand<ActionsSuccess<ActionsExport>>, Expand<ActionsFailure<ActionsExport>>>'
				);

				type = 'Expand<Kit.AwaitedActions<ActionsExport>> | null';
			}
			exports.push(`export type ActionData = ${type};`);
		}
	} else {
		server_data = 'null';
	}
	exports.push(`export type ${prefix}ServerData = ${server_data};`);

	const parent_type = `${prefix}ParentData`;
	declarations.push(`type ${parent_type} = ${get_parent_type(node, 'LayoutData')};`);

	if (node.universal) {
		const proxy = proxies.universal;
		if (proxy?.modified) {
			fs.writeFileSync(`${outdir}/proxy${path.basename(node.universal)}`, proxy.code);
		}

		const type = get_data_type(
			node.universal,
			`${parent_type} & EnsureDefined<${prefix}ServerData>`,
			proxy
		);

		data = `Expand<Omit<${parent_type}, keyof ${type}> & OptionalUnion<EnsureDefined<${type}>>>`;

		const output_data_shape =
			!is_page && all_pages_have_load
				? 'Partial<App.PageData> & Record<string, any> | void'
				: `OutputDataShape<${parent_type}>`;
		exports.push(
			`export type ${prefix}Load<OutputData extends ${output_data_shape} = ${output_data_shape}> = Kit.Load<${params}, ${prefix}ServerData, ${parent_type}, OutputData, ${route_id}>;`
		);

		exports.push(`export type ${prefix}LoadEvent = Parameters<${prefix}Load>[0];`);
	} else if (server_data === 'null') {
		data = `Expand<${parent_type}>`;
	} else {
		data = `Expand<Omit<${parent_type}, keyof ${prefix}ServerData> & EnsureDefined<${prefix}ServerData>>`;
	}

	exports.push(`export type ${prefix}Data = ${data};`);

	return { declarations, exports, proxies };

	/**
	 * @param {string} file_path
	 * @param {string} fallback
	 * @param {Proxy} proxy
	 * @param {boolean} expand
	 */
	function get_data_type(file_path, fallback, proxy, expand = false) {
		if (proxy) {
			if (proxy.exports.includes('load')) {
				// If the file wasn't tweaked, we can use the return type of the original file.
				// The advantage is that type updates are reflected without saving.
				const from = proxy.modified
					? `./proxy${replace_ext_with_js(path.basename(file_path))}`
					: path_to_original(outdir, file_path);
				const type = `Kit.LoadProperties<Awaited<ReturnType<typeof import('${from}').load>>>`;
				return expand ? `Expand<OptionalUnion<EnsureDefined<${type}>>>` : type;
			} else {
				return fallback;
			}
		} else {
			return 'unknown';
		}
	}
}

/**
 * This function populates the proxies object, if necessary and not already done.
 * Proxies are used to tweak the code of a file before it's typechecked.
 * They are needed in two places - when generating the types for a page or layout.
 * To not do the same work twice, we generate the proxies once and pass them around.
 *
 * @param {import('types').PageNode} node
 * @param {Proxies} proxies
 */
function ensureProxies(node, proxies) {
	if (node.server && !proxies.server) {
		proxies.server = createProxy(node.server, true);
	}

	if (node.universal && !proxies.universal) {
		proxies.universal = createProxy(node.universal, false);
	}
}

/**
 * @param {string} file_path
 * @param {boolean} is_server
 * @returns {Proxy}
 */
function createProxy(file_path, is_server) {
	const proxy = tweak_types(fs.readFileSync(file_path, 'utf8'), is_server);
	if (proxy) {
		return {
			...proxy,
			file_name: `proxy${path.basename(file_path)}`
		};
	} else {
		return null;
	}
}

/**
 * Get the parent type string by recursively looking up the parent layout and accumulate them to one type.
 * @param {import('types').PageNode} node
 * @param {string} type
 */
function get_parent_type(node, type) {
	const parent_imports = [];

	let parent = node.parent;

	while (parent) {
		const d = node.depth - parent.depth;
		// unshift because we need it the other way round for the import string
		parent_imports.unshift(
			`${d === 0 ? '' : `import('${'../'.repeat(d)}${'$types.js'}').`}${type}`
		);
		parent = parent.parent;
	}

	let parent_str = `EnsureDefined<${parent_imports[0] || '{}'}>`;
	for (let i = 1; i < parent_imports.length; i++) {
		// Omit is necessary because a parent could have a property with the same key which would
		// cause a type conflict. At runtime the child overwrites the parent property in this case,
		// so reflect that in the type definition.
		// EnsureDefined is necessary because {something: string} & null becomes null.
		// Output types of server loads can be null but when passed in through the `parent` parameter they are the empty object instead.
		parent_str = `Omit<${parent_str}, keyof ${parent_imports[i]}> & EnsureDefined<${parent_imports[i]}>`;
	}
	return parent_str;
}

/**
 * @param {string} outdir
 * @param {string} file_path
 */
function path_to_original(outdir, file_path) {
	return posixify(path.relative(outdir, path.join(cwd, replace_ext_with_js(file_path))));
}

/**
 * @param {string} file_path
 */
function replace_ext_with_js(file_path) {
	// Another extension than `.js` (or nothing, but that fails with node16 moduleResolution)
	// will result in TS failing to lookup the file
	const ext = path.extname(file_path);
	return file_path.slice(0, -ext.length) + '.js';
}

/**
 * @param {import('types').RouteParam[]} params
 * @param {string} outdir
 * @param {import('types').ValidatedConfig} config
 */
function generate_params_type(params, outdir, config) {
	/** @param {string} matcher */
	const path_to_matcher = (matcher) =>
		posixify(path.relative(outdir, path.join(config.kit.files.params, matcher + '.js')));

	return `{ ${params
		.map(
			(param) =>
				`${param.name}${param.optional ? '?' : ''}: ${
					param.matcher
						? `MatcherParam<typeof import('${path_to_matcher(param.matcher)}').match>`
						: 'string'
				}`
		)
		.join('; ')} }`;
}

/**
 * @param {string} content
 * @param {boolean} is_server
 * @returns {Omit<NonNullable<Proxy>, 'file_name'> | null}
 */
export function tweak_types(content, is_server) {
	const names = new Set(is_server ? ['load', 'actions'] : ['load']);

	try {
		let modified = false;

		const ast = ts.createSourceFile(
			'filename.ts',
			content,
			ts.ScriptTarget.Latest,
			false,
			ts.ScriptKind.TS
		);

		const code = new MagicString(content);

		const exports = new Map();

		ast.forEachChild((node) => {
			if (
				ts.isExportDeclaration(node) &&
				node.exportClause &&
				ts.isNamedExports(node.exportClause)
			) {
				node.exportClause.elements.forEach((element) => {
					const exported = element.name;
					if (names.has(element.name.text)) {
						const local = element.propertyName || element.name;
						exports.set(exported.text, local.text);
					}
				});
			}

			if (
				ts.canHaveModifiers(node) &&
				ts.getModifiers(node)?.some((modifier) => modifier.kind === ts.SyntaxKind.ExportKeyword)
			) {
				if (ts.isFunctionDeclaration(node) && node.name?.text && names.has(node.name?.text)) {
					exports.set(node.name.text, node.name.text);
				}

				if (ts.isVariableStatement(node)) {
					node.declarationList.declarations.forEach((declaration) => {
						if (ts.isIdentifier(declaration.name) && names.has(declaration.name.text)) {
							exports.set(declaration.name.text, declaration.name.text);
						}
					});
				}
			}
		});

		/**
		 * @param {import('typescript').Node} node
		 * @param {import('typescript').Node} value
		 */
		function replace_jsdoc_type_tags(node, value) {
			let _modified = false;
			// @ts-ignore
			if (node.jsDoc) {
				// @ts-ignore
				for (const comment of node.jsDoc) {
					for (const tag of comment.tags ?? []) {
						if (ts.isJSDocTypeTag(tag)) {
							const is_fn =
								ts.isFunctionDeclaration(value) ||
								ts.isFunctionExpression(value) ||
								ts.isArrowFunction(value);

							if (is_fn && value.parameters?.length > 0) {
								const name = ts.isIdentifier(value.parameters[0].name)
									? value.parameters[0].name.text
									: 'event';
								code.overwrite(tag.tagName.pos, tag.tagName.end, 'param');
								code.prependRight(tag.typeExpression.pos + 1, 'Parameters<');
								code.appendLeft(tag.typeExpression.end - 1, '>[0]');
								code.appendLeft(tag.typeExpression.end, ` ${name}`);
							} else {
								code.overwrite(tag.pos, tag.end, '');
							}
							_modified = true;
						}
					}
				}
			}
			modified = modified || _modified;
			return _modified;
		}

		ast.forEachChild((node) => {
			if (ts.isFunctionDeclaration(node) && node.name?.text && node.name?.text === 'load') {
				// remove JSDoc comment above `export function load ...`
				replace_jsdoc_type_tags(node, node);
			}

			if (ts.isVariableStatement(node)) {
				// remove JSDoc comment above `export const load = ...`
				if (
					ts.isIdentifier(node.declarationList.declarations[0].name) &&
					names.has(node.declarationList.declarations[0].name.text) &&
					node.declarationList.declarations[0].initializer
				) {
					replace_jsdoc_type_tags(node, node.declarationList.declarations[0].initializer);
				}

				for (const declaration of node.declarationList.declarations) {
					if (
						ts.isIdentifier(declaration.name) &&
						declaration.name.text === 'load' &&
						declaration.initializer
					) {
						// edge case — remove JSDoc comment above individual export
						replace_jsdoc_type_tags(declaration, declaration.initializer);

						// remove type from `export const load: Load ...`
						if (declaration.type) {
							let a = declaration.type.pos;
							const b = declaration.type.end;
							while (/\s/.test(content[a])) a += 1;

							const type = content.slice(a, b);
							code.remove(declaration.name.end, declaration.type.end);

							const rhs = declaration.initializer;

							if (
								rhs &&
								(ts.isArrowFunction(rhs) || ts.isFunctionExpression(rhs)) &&
								rhs.parameters.length
							) {
								const arg = rhs.parameters[0];
								const add_parens = content[arg.pos - 1] !== '(';

								if (add_parens) code.prependRight(arg.pos, '(');

								if (arg && !arg.type) {
									code.appendLeft(
										arg.name.end,
										`: Parameters<${type}>[0]` + (add_parens ? ')' : '')
									);
								} else {
									// prevent "type X is imported but not used" (isn't silenced by @ts-nocheck) when svelte-check runs
									code.append(`;null as any as ${type};`);
								}
							} else {
								// prevent "type X is imported but not used" (isn't silenced by @ts-nocheck) when svelte-check runs
								code.append(`;null as any as ${type};`);
							}

							modified = true;
						}
					} else if (
						is_server &&
						ts.isIdentifier(declaration.name) &&
						declaration.name?.text === 'actions' &&
						declaration.initializer
					) {
						// remove JSDoc comment from `export const actions = ..`
						const removed = replace_jsdoc_type_tags(node, declaration.initializer);
						// ... and move type to each individual action
						if (removed) {
							const rhs = declaration.initializer;
							if (ts.isObjectLiteralExpression(rhs)) {
								for (const prop of rhs.properties) {
									if (ts.isPropertyAssignment(prop) && ts.isIdentifier(prop.name)) {
										const rhs = prop.initializer;
										const replaced = replace_jsdoc_type_tags(prop, rhs);
										if (
											!replaced &&
											rhs &&
											(ts.isArrowFunction(rhs) || ts.isFunctionExpression(rhs)) &&
											rhs.parameters?.[0]
										) {
											const name = ts.isIdentifier(rhs.parameters[0].name)
												? rhs.parameters[0].name.text
												: 'event';
											code.prependRight(
												rhs.pos,
												`/** @param {import('./$types').RequestEvent} ${name} */ `
											);
										}
									}
								}
							}
						}

						// remove type from `export const actions: Actions ...`
						if (declaration.type) {
							let a = declaration.type.pos;
							const b = declaration.type.end;
							while (/\s/.test(content[a])) a += 1;

							const type = content.slice(a, b);
							code.remove(declaration.name.end, declaration.type.end);
							code.append(`;null as any as ${type};`);
							modified = true;

							// ... and move type to each individual action
							const rhs = declaration.initializer;
							if (ts.isObjectLiteralExpression(rhs)) {
								for (const prop of rhs.properties) {
									if (ts.isPropertyAssignment(prop) && ts.isIdentifier(prop.name)) {
										const rhs = prop.initializer;

										if (
											rhs &&
											(ts.isArrowFunction(rhs) || ts.isFunctionExpression(rhs)) &&
											rhs.parameters.length
										) {
											const arg = rhs.parameters[0];
											const add_parens = content[arg.pos - 1] !== '(';

											if (add_parens) code.prependRight(arg.pos, '(');

											if (arg && !arg.type) {
												code.appendLeft(
													arg.name.end,
													": import('./$types').RequestEvent" + (add_parens ? ')' : '')
												);
											}
										}
									}
								}
							}
						}
					}
				}
			}
		});

		if (modified) {
			// Ignore all type errors so they don't show up twice when svelte-check runs
			// Account for possible @ts-check which would overwrite @ts-nocheck
			if (code.original.startsWith('// @ts-check')) {
				code.prependLeft('// @ts-check'.length, '\n// @ts-nocheck\n');
			} else {
				code.prepend('// @ts-nocheck\n');
			}
		}

		return {
			modified,
			code: code.toString(),
			exports: Array.from(exports.keys())
		};
	} catch {
		return null;
	}
}
