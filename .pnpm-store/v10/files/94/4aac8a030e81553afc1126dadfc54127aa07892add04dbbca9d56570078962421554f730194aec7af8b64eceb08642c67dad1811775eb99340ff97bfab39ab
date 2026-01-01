import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { get_env } from '../../exports/vite/utils.js';
import { GENERATED_COMMENT } from '../../constants.js';
import { create_dynamic_types, create_static_types } from '../env.js';
import { write_if_changed } from './utils.js';

// TODO these types should be described in a neutral place, rather than
// inside either `packages/kit` or `svelte.dev/docs/kit`
const descriptions_dir = fileURLToPath(new URL('../../../src/types/synthetic', import.meta.url));

/** @param {string} filename */
function read_description(filename) {
	const content = fs.readFileSync(`${descriptions_dir}/${filename}`, 'utf8');
	return `/**\n${content
		.trim()
		.split('\n')
		.map((line) => ` * ${line}`)
		.join('\n')}\n */`;
}

/**
 * @param {import('types').Env} env
 * @param {{
 * 	public_prefix: string;
 * 	private_prefix: string;
 * }} prefixes
 */
const template = (env, prefixes) => `
${GENERATED_COMMENT}

/// <reference types="@sveltejs/kit" />

${read_description('$env+static+private.md')}
${create_static_types('private', env)}

${read_description('$env+static+public.md')}
${create_static_types('public', env)}

${read_description('$env+dynamic+private.md')}
${create_dynamic_types('private', env, prefixes)}

${read_description('$env+dynamic+public.md')}
${create_dynamic_types('public', env, prefixes)}
`;

/**
 * Writes ambient declarations including types reference to @sveltejs/kit,
 * and the existing environment variables in process.env to
 * $env/static/private and $env/static/public
 * @param {import('types').ValidatedKitConfig} config
 * @param {string} mode The Vite mode
 */
export function write_ambient(config, mode) {
	const env = get_env(config.env, mode);
	const { publicPrefix: public_prefix, privatePrefix: private_prefix } = config.env;

	write_if_changed(
		path.join(config.outDir, 'ambient.d.ts'),
		template(env, { public_prefix, private_prefix })
	);
}
