import { GENERATED_COMMENT } from '../constants.js';
import { dedent } from './sync/utils.js';
import { runtime_base } from './utils.js';

/**
 * @typedef {'public' | 'private'} EnvType
 */

/**
 * @param {string} id
 * @param {Record<string, string>} env
 * @returns {string}
 */
export function create_static_module(id, env) {
	/** @type {string[]} */
	const declarations = [];

	for (const key in env) {
		if (!valid_identifier.test(key) || reserved.has(key)) {
			continue;
		}

		const comment = `/** @type {import('${id}').${key}} */`;
		const declaration = `export const ${key} = ${JSON.stringify(env[key])};`;

		declarations.push(`${comment}\n${declaration}`);
	}

	return GENERATED_COMMENT + declarations.join('\n\n');
}

/**
 * @param {EnvType} type
 * @param {Record<string, string> | undefined} dev_values If in a development mode, values to pre-populate the module with.
 */
export function create_dynamic_module(type, dev_values) {
	if (dev_values) {
		const keys = Object.entries(dev_values).map(
			([k, v]) => `${JSON.stringify(k)}: ${JSON.stringify(v)}`
		);
		return `export const env = {\n${keys.join(',\n')}\n}`;
	}
	return `export { ${type}_env as env } from '${runtime_base}/shared-server.js';`;
}

/**
 * @param {EnvType} id
 * @param {import('types').Env} env
 * @returns {string}
 */
export function create_static_types(id, env) {
	const declarations = Object.keys(env[id])
		.filter((k) => valid_identifier.test(k))
		.map((k) => `export const ${k}: string;`);

	return dedent`
		declare module '$env/static/${id}' {
			${declarations.join('\n')}
		}
	`;
}

/**
 * @param {EnvType} id
 * @param {import('types').Env} env
 * @param {{
 * 	public_prefix: string;
 * 	private_prefix: string;
 * }} prefixes
 * @returns {string}
 */
export function create_dynamic_types(id, env, { public_prefix, private_prefix }) {
	const properties = Object.keys(env[id])
		.filter((k) => valid_identifier.test(k))
		.map((k) => `${k}: string;`);

	const public_prefixed = `[key: \`${public_prefix}\${string}\`]`;
	const private_prefixed = `[key: \`${private_prefix}\${string}\`]`;

	if (id === 'private') {
		if (public_prefix) {
			properties.push(`${public_prefixed}: undefined;`);
		}
		properties.push(`${private_prefixed}: string | undefined;`);
	} else {
		if (private_prefix) {
			properties.push(`${private_prefixed}: undefined;`);
		}
		properties.push(`${public_prefixed}: string | undefined;`);
	}

	return dedent`
		declare module '$env/dynamic/${id}' {
			export const env: {
				${properties.join('\n')}
			}
		}
	`;
}

export const reserved = new Set([
	'do',
	'if',
	'in',
	'for',
	'let',
	'new',
	'try',
	'var',
	'case',
	'else',
	'enum',
	'eval',
	'null',
	'this',
	'true',
	'void',
	'with',
	'await',
	'break',
	'catch',
	'class',
	'const',
	'false',
	'super',
	'throw',
	'while',
	'yield',
	'delete',
	'export',
	'import',
	'public',
	'return',
	'static',
	'switch',
	'typeof',
	'default',
	'extends',
	'finally',
	'package',
	'private',
	'continue',
	'debugger',
	'function',
	'arguments',
	'interface',
	'protected',
	'implements',
	'instanceof'
]);

export const valid_identifier = /^[a-zA-Z_$][a-zA-Z0-9_$]*$/;
