import fs from 'node:fs';
import path from 'node:path';
import process from 'node:process';
import colors from 'kleur';
import { posixify } from '../../utils/filesystem.js';
import { write_if_changed } from './utils.js';

/**
 * @param {string} cwd
 * @param {string} file
 */
function maybe_file(cwd, file) {
	const resolved = path.resolve(cwd, file);
	if (fs.existsSync(resolved)) {
		return resolved;
	}
}

/**
 * @param {string} file
 */
function project_relative(file) {
	return posixify(path.relative('.', file));
}

/**
 * @param {string} file
 */
function remove_trailing_slashstar(file) {
	if (file.endsWith('/*')) {
		return file.slice(0, -2);
	} else {
		return file;
	}
}

/**
 * Generates the tsconfig that the user's tsconfig inherits from.
 * @param {import('types').ValidatedKitConfig} kit
 */
export function write_tsconfig(kit, cwd = process.cwd()) {
	const out = path.join(kit.outDir, 'tsconfig.json');

	const user_config = load_user_tsconfig(cwd);
	if (user_config) validate_user_config(cwd, out, user_config);

	write_if_changed(out, JSON.stringify(get_tsconfig(kit), null, '\t'));
}

/**
 * Generates the tsconfig that the user's tsconfig inherits from.
 * @param {import('types').ValidatedKitConfig} kit
 */
export function get_tsconfig(kit) {
	/** @param {string} file */
	const config_relative = (file) => posixify(path.relative(kit.outDir, file));

	const include = new Set([
		'ambient.d.ts', // careful: changing this name would be a breaking change, because it's referenced in the service-workers documentation
		'non-ambient.d.ts',
		'./types/**/$types.d.ts',
		config_relative('vite.config.js'),
		config_relative('vite.config.ts')
	]);
	// TODO(v2): find a better way to include all src files. We can't just use routes/lib only because
	// people might have other folders/files in src that they want included.
	const src_includes = [kit.files.routes, kit.files.lib, path.resolve('src')].filter((dir) => {
		const relative = path.relative(path.resolve('src'), dir);
		return !relative || relative.startsWith('..');
	});
	for (const dir of src_includes) {
		include.add(config_relative(`${dir}/**/*.js`));
		include.add(config_relative(`${dir}/**/*.ts`));
		include.add(config_relative(`${dir}/**/*.svelte`));
	}

	// Test folder is a special case - we advocate putting tests in a top-level test folder
	// and it's not configurable (should we make it?)
	const test_folder = project_relative('tests');
	include.add(config_relative(`${test_folder}/**/*.js`));
	include.add(config_relative(`${test_folder}/**/*.ts`));
	include.add(config_relative(`${test_folder}/**/*.svelte`));

	const exclude = [config_relative('node_modules/**')];
	// Add service worker to exclude list so that worker types references in it don't spill over into the rest of the app
	// (i.e. suddenly ServiceWorkerGlobalScope would be available throughout the app, and some types might even clash)
	if (path.extname(kit.files.serviceWorker)) {
		exclude.push(config_relative(kit.files.serviceWorker));
	} else {
		exclude.push(config_relative(`${kit.files.serviceWorker}.js`));
		exclude.push(config_relative(`${kit.files.serviceWorker}/**/*.js`));
		exclude.push(config_relative(`${kit.files.serviceWorker}.ts`));
		exclude.push(config_relative(`${kit.files.serviceWorker}/**/*.ts`));
		exclude.push(config_relative(`${kit.files.serviceWorker}.d.ts`));
		exclude.push(config_relative(`${kit.files.serviceWorker}/**/*.d.ts`));
	}

	const config = {
		compilerOptions: {
			// generated options
			paths: get_tsconfig_paths(kit),
			rootDirs: [config_relative('.'), './types'],

			// essential options
			// svelte-preprocess cannot figure out whether you have a value or a type, so tell TypeScript
			// to enforce using \`import type\` instead of \`import\` for Types.
			// Also, TypeScript doesn't know about import usages in the template because it only sees the
			// script of a Svelte file. Therefore preserve all value imports.
			verbatimModuleSyntax: true,
			// Vite compiles modules one at a time
			isolatedModules: true,

			// This is required for svelte-package to work as expected
			// Can be overwritten
			lib: ['esnext', 'DOM', 'DOM.Iterable'],
			moduleResolution: 'bundler',
			module: 'esnext',
			noEmit: true, // prevent tsconfig error "overwriting input files" - Vite handles the build and ignores this
			target: 'esnext'
		},
		include: [...include],
		exclude
	};

	return kit.typescript.config(config) ?? config;
}

/** @param {string} cwd */
function load_user_tsconfig(cwd) {
	const file = maybe_file(cwd, 'tsconfig.json') || maybe_file(cwd, 'jsconfig.json');

	if (!file) return;

	// we have to eval the file, since it's not parseable as JSON (contains comments)
	const json = fs.readFileSync(file, 'utf-8');

	return {
		kind: path.basename(file),
		options: (0, eval)(`(${json})`)
	};
}

/**
 * @param {string} cwd
 * @param {string} out
 * @param {{ kind: string, options: any }} config
 */
function validate_user_config(cwd, out, config) {
	// we need to check that the user's tsconfig extends the framework config
	const extend = config.options.extends;
	const extends_framework_config =
		typeof extend === 'string'
			? path.resolve(cwd, extend) === out
			: Array.isArray(extend)
				? extend.some((e) => path.resolve(cwd, e) === out)
				: false;

	const options = config.options.compilerOptions || {};

	if (extends_framework_config) {
		const { paths, baseUrl } = options;

		if (baseUrl || paths) {
			console.warn(
				colors
					.bold()
					.yellow(
						`You have specified a baseUrl and/or paths in your ${config.kind} which interferes with SvelteKit's auto-generated tsconfig.json. ` +
							'Remove it to avoid problems with intellisense. For path aliases, use `kit.alias` instead: https://svelte.dev/docs/kit/configuration#alias'
					)
			);
		}
	} else {
		let relative = posixify(path.relative('.', out));
		if (!relative.startsWith('./')) relative = './' + relative;

		console.warn(
			colors
				.bold()
				.yellow(`Your ${config.kind} should extend the configuration generated by SvelteKit:`)
		);
		console.warn(`{\n  "extends": "${relative}"\n}`);
	}
}

// <something><optional /*>
const alias_regex = /^(.+?)(\/\*)?$/;
// <path><optional /* or .fileending>
const value_regex = /^(.*?)((\/\*)|(\.\w+))?$/;

/**
 * Generates tsconfig path aliases from kit's aliases.
 * Related to vite alias creation.
 *
 * @param {import('types').ValidatedKitConfig} config
 */
function get_tsconfig_paths(config) {
	/** @param {string} file */
	const config_relative = (file) => {
		let relative_path = path.relative(config.outDir, file);
		if (!relative_path.startsWith('..')) {
			relative_path = './' + relative_path;
		}
		return posixify(relative_path);
	};

	const alias = { ...config.alias };
	if (fs.existsSync(project_relative(config.files.lib))) {
		alias['$lib'] = project_relative(config.files.lib);
	}

	/** @type {Record<string, string[]>} */
	const paths = {};

	for (const [key, value] of Object.entries(alias)) {
		const key_match = alias_regex.exec(key);
		if (!key_match) throw new Error(`Invalid alias key: ${key}`);

		const value_match = value_regex.exec(value);
		if (!value_match) throw new Error(`Invalid alias value: ${value}`);

		const rel_path = config_relative(remove_trailing_slashstar(value));
		const slashstar = key_match[2];

		if (slashstar) {
			paths[key] = [rel_path + '/*'];
		} else {
			paths[key] = [rel_path];
			const fileending = value_match[4];

			if (!fileending && !(key + '/*' in alias)) {
				paths[key + '/*'] = [rel_path + '/*'];
			}
		}
	}

	return paths;
}
