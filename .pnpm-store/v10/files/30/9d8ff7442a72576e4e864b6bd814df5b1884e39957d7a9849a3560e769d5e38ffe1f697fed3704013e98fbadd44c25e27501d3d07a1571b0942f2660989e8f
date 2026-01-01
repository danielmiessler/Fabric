/** @import { RemoteInfo } from 'types' */

/** @type {RemoteInfo['type'][]} */
const types = ['command', 'form', 'prerender', 'query', 'query_batch'];

/**
 * @param {Record<string, any>} module
 * @param {string} file
 * @param {string} hash
 */
export function init_remote_functions(module, file, hash) {
	if (module.default) {
		throw new Error(
			`Cannot export \`default\` from a remote module (${file}) — please use named exports instead`
		);
	}

	for (const [name, fn] of Object.entries(module)) {
		if (!types.includes(fn?.__?.type)) {
			throw new Error(
				`\`${name}\` exported from ${file} is invalid — all exports from this file must be remote functions`
			);
		}

		fn.__.id = `${hash}/${name}`;
		fn.__.name = name;
	}
}
