import process from 'node:process';
import fs from 'node:fs';
import path from 'node:path';

/**
 * Resolves a peer dependency relative to the current CWD. Duplicated with `packages/adapter-auto`
 * @param {string} dependency
 */
function resolve_peer(dependency) {
	let [name, ...parts] = dependency.split('/');
	if (name[0] === '@') name += `/${parts.shift()}`;

	let dir = process.cwd();

	while (!fs.existsSync(`${dir}/node_modules/${name}/package.json`)) {
		if (dir === (dir = path.dirname(dir))) {
			throw new Error(
				`Could not resolve peer dependency "${name}" relative to your project — please install it and try again.`
			);
		}
	}

	const pkg_dir = `${dir}/node_modules/${name}`;
	const pkg = JSON.parse(fs.readFileSync(`${pkg_dir}/package.json`, 'utf-8'));

	const subpackage = ['.', ...parts].join('/');

	let exported = pkg.exports[subpackage];

	while (typeof exported !== 'string') {
		if (!exported) {
			throw new Error(`Could not find valid "${subpackage}" export in ${name}/package.json`);
		}

		exported = exported['import'] ?? exported['default'];
	}

	return path.resolve(pkg_dir, exported);
}

/**
 * Resolve a dependency relative to the current working directory,
 * rather than relative to this package (but falls back to trying that, if necessary)
 * @param {string} dependency
 */
export async function import_peer(dependency) {
	try {
		return await import(resolve_peer(dependency));
	} catch {
		return await import(dependency);
	}
}
