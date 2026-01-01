import { base, assets, relative } from '$app/paths/internal/server';
import { text } from '@sveltejs/kit';
import { s } from '../../../utils/misc.js';
import { exec } from '../../../utils/routing.js';
import { decode_params } from '../../../utils/url.js';
import { get_relative_path } from '../../utils.js';

/**
 * @param {import('types').SSRClientRoute} route
 * @param {URL} url
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @returns {string}
 */
export function generate_route_object(route, url, manifest) {
	const { errors, layouts, leaf } = route;

	const nodes = [...errors, ...layouts.map((l) => l?.[1]), leaf[1]]
		.filter((n) => typeof n === 'number')
		.map((n) => `'${n}': () => ${create_client_import(manifest._.client.nodes?.[n], url)}`)
		.join(',\n\t\t');

	// stringified version of
	/** @type {import('types').CSRRouteServer} */
	return [
		`{\n\tid: ${s(route.id)}`,
		`errors: ${s(route.errors)}`,
		`layouts: ${s(route.layouts)}`,
		`leaf: ${s(route.leaf)}`,
		`nodes: {\n\t\t${nodes}\n\t}\n}`
	].join(',\n\t');
}

/**
 * @param {string | undefined} import_path
 * @param {URL} url
 */
function create_client_import(import_path, url) {
	if (!import_path) return 'Promise.resolve({})';

	// During DEV, Vite will make the paths absolute (e.g. /@fs/...)
	if (import_path[0] === '/') {
		return `import('${import_path}')`;
	}

	// During PROD, they're root-relative
	if (assets !== '') {
		return `import('${assets}/${import_path}')`;
	}

	if (!relative) {
		return `import('${base}/${import_path}')`;
	}

	// Else we make them relative to the server-side route resolution request
	// to support IPFS, the internet archive, etc.
	let path = get_relative_path(url.pathname, `${base}/${import_path}`);
	if (path[0] !== '.') path = `./${path}`;
	return `import('${path}')`;
}

/**
 * @param {string} resolved_path
 * @param {URL} url
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @returns {Promise<Response>}
 */
export async function resolve_route(resolved_path, url, manifest) {
	if (!manifest._.client.routes) {
		return text('Server-side route resolution disabled', { status: 400 });
	}

	/** @type {import('types').SSRClientRoute | null} */
	let route = null;
	/** @type {Record<string, string>} */
	let params = {};

	const matchers = await manifest._.matchers();

	for (const candidate of manifest._.client.routes) {
		const match = candidate.pattern.exec(resolved_path);
		if (!match) continue;

		const matched = exec(match, candidate.params, matchers);
		if (matched) {
			route = candidate;
			params = decode_params(matched);
			break;
		}
	}

	return create_server_routing_response(route, params, url, manifest).response;
}

/**
 * @param {import('types').SSRClientRoute | null} route
 * @param {Partial<Record<string, string>>} params
 * @param {URL} url
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @returns {{response: Response, body: string}}
 */
export function create_server_routing_response(route, params, url, manifest) {
	const headers = new Headers({
		'content-type': 'application/javascript; charset=utf-8'
	});

	if (route) {
		const csr_route = generate_route_object(route, url, manifest);
		const body = `${create_css_import(route, url, manifest)}\nexport const route = ${csr_route}; export const params = ${JSON.stringify(params)};`;

		return { response: text(body, { headers }), body };
	} else {
		return { response: text('', { headers }), body: '' };
	}
}

/**
 * This function generates the client-side import for the CSS files that are
 * associated with the current route. Vite takes care of that when using
 * client-side route resolution, but for server-side resolution it does
 * not know about the CSS files automatically.
 *
 * @param {import('types').SSRClientRoute} route
 * @param {URL} url
 * @param {import('@sveltejs/kit').SSRManifest} manifest
 * @returns {string}
 */
function create_css_import(route, url, manifest) {
	const { errors, layouts, leaf } = route;

	let css = '';

	for (const node of [...errors, ...layouts.map((l) => l?.[1]), leaf[1]]) {
		if (typeof node !== 'number') continue;
		const node_css = manifest._.client.css?.[node];
		for (const css_path of node_css ?? []) {
			css += `'${assets || base}/${css_path}',`;
		}
	}

	if (!css) return '';

	return `${create_client_import(/** @type {string} */ (manifest._.client.start), url)}.then(x => x.load_css([${css}]));`;
}
