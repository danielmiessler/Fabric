import { exec, parse_route_id } from '../../utils/routing.js';

/**
 * @param {import('./types.js').SvelteKitApp} app
 * @returns {import('types').CSRRoute[]}
 */
export function parse({ nodes, server_loads, dictionary, matchers }) {
	const layouts_with_server_load = new Set(server_loads);

	return Object.entries(dictionary).map(([id, [leaf, layouts, errors]]) => {
		const { pattern, params } = parse_route_id(id);

		/** @type {import('types').CSRRoute} */
		const route = {
			id,
			/** @param {string} path */
			exec: (path) => {
				const match = pattern.exec(path);
				if (match) return exec(match, params, matchers);
			},
			errors: [1, ...(errors || [])].map((n) => nodes[n]),
			layouts: [0, ...(layouts || [])].map(create_layout_loader),
			leaf: create_leaf_loader(leaf)
		};

		// bit of a hack, but ensures that layout/error node lists are the same
		// length, without which the wrong data will be applied if the route
		// manifest looks like `[[a, b], [c,], d]`
		route.errors.length = route.layouts.length = Math.max(
			route.errors.length,
			route.layouts.length
		);

		return route;
	});

	/**
	 * @param {number} id
	 * @returns {[boolean, import('types').CSRPageNodeLoader]}
	 */
	function create_leaf_loader(id) {
		// whether or not the route uses the server data is
		// encoded using the ones' complement, to save space
		const uses_server_data = id < 0;
		if (uses_server_data) id = ~id;
		return [uses_server_data, nodes[id]];
	}

	/**
	 * @param {number | undefined} id
	 * @returns {[boolean, import('types').CSRPageNodeLoader] | undefined}
	 */
	function create_layout_loader(id) {
		// whether or not the layout uses the server data is
		// encoded in the layouts array, to save space
		return id === undefined ? id : [layouts_with_server_load.has(id), nodes[id]];
	}
}

/**
 * @param {import('types').CSRRouteServer} input
 * @param {import('types').CSRPageNodeLoader[]} app_nodes Will be modified if a new node is loaded that's not already in the array
 * @returns {import('types').CSRRoute}
 */
export function parse_server_route({ nodes, id, leaf, layouts, errors }, app_nodes) {
	return {
		id,
		exec: () => ({}), // dummy function; exec already happened on the server
		// By writing to app_nodes only when a loader at that index is not already defined,
		// we ensure that loaders have referential equality when they load the same node.
		// Code elsewhere in client.js relies on this referential equality to determine
		// if a loader is different and should therefore (re-)run.
		errors: errors.map((n) => (n ? (app_nodes[n] ||= nodes[n]) : undefined)),
		layouts: layouts.map((n) => (n ? [n[0], (app_nodes[n[1]] ||= nodes[n[1]])] : undefined)),
		leaf: [leaf[0], (app_nodes[leaf[1]] ||= nodes[leaf[1]])]
	};
}
