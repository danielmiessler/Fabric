import {
	validate_layout_exports,
	validate_layout_server_exports,
	validate_page_exports,
	validate_page_server_exports
} from './exports.js';

export class PageNodes {
	data;

	/**
	 * @param {Array<import('types').SSRNode | undefined>} nodes
	 */
	constructor(nodes) {
		this.data = nodes;
	}

	layouts() {
		return this.data.slice(0, -1);
	}

	page() {
		return this.data.at(-1);
	}

	validate() {
		for (const layout of this.layouts()) {
			if (layout) {
				validate_layout_server_exports(layout.server, /** @type {string} */ (layout.server_id));
				validate_layout_exports(layout.universal, /** @type {string} */ (layout.universal_id));
			}
		}

		const page = this.page();
		if (page) {
			validate_page_server_exports(page.server, /** @type {string} */ (page.server_id));
			validate_page_exports(page.universal, /** @type {string} */ (page.universal_id));
		}
	}

	/**
	 * @template {'prerender' | 'ssr' | 'csr' | 'trailingSlash'} Option
	 * @param {Option} option
	 * @returns {Value | undefined}
	 */
	#get_option(option) {
		/** @typedef {(import('types').UniversalNode | import('types').ServerNode)[Option]} Value */

		return this.data.reduce((value, node) => {
			return node?.universal?.[option] ?? node?.server?.[option] ?? value;
		}, /** @type {Value | undefined} */ (undefined));
	}

	csr() {
		return this.#get_option('csr') ?? true;
	}

	ssr() {
		return this.#get_option('ssr') ?? true;
	}

	prerender() {
		return this.#get_option('prerender') ?? false;
	}

	trailing_slash() {
		return this.#get_option('trailingSlash') ?? 'never';
	}

	get_config() {
		/** @type {any} */
		let current = {};

		for (const node of this.data) {
			if (!node?.universal?.config && !node?.server?.config) continue;

			current = {
				...current,
				// TODO: should we override the server config value with the universal value similar to other page options?
				...node?.universal?.config,
				...node?.server?.config
			};
		}

		// TODO 3.0 always return `current`? then we can get rid of `?? {}` in other places
		return Object.keys(current).length ? current : undefined;
	}

	should_prerender_data() {
		return this.data.some(
			// prerender in case of trailingSlash because the client retrieves that value from the server
			(node) => node?.server?.load || node?.server?.trailingSlash !== undefined
		);
	}
}
