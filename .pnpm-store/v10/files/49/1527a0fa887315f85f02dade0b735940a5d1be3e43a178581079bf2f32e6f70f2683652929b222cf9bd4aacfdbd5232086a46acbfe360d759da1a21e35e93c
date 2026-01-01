import { getContext } from 'svelte';
import { BROWSER } from 'esm-env';
import { stores as browser_stores } from '../client/client.js';

/**
 * A function that returns all of the contextual stores. On the server, this must be called during component initialization.
 * Only use this if you need to defer store subscription until after the component has mounted, for some reason.
 *
 * @deprecated Use `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
 */
export const getStores = () => {
	const stores = BROWSER ? browser_stores : getContext('__svelte__');

	return {
		/** @type {typeof page} */
		page: {
			subscribe: stores.page.subscribe
		},
		/** @type {typeof navigating} */
		navigating: {
			subscribe: stores.navigating.subscribe
		},
		/** @type {typeof updated} */
		updated: stores.updated
	};
};

/**
 * A readable store whose value contains page data.
 *
 * On the server, this store can only be subscribed to during component initialization. In the browser, it can be subscribed to at any time.
 *
 * @deprecated Use `page` from `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
 * @type {import('svelte/store').Readable<import('@sveltejs/kit').Page>}
 */
export const page = {
	subscribe(fn) {
		const store = __SVELTEKIT_DEV__ ? get_store('page') : getStores().page;
		return store.subscribe(fn);
	}
};

/**
 * A readable store.
 * When navigating starts, its value is a `Navigation` object with `from`, `to`, `type` and (if `type === 'popstate'`) `delta` properties.
 * When navigating finishes, its value reverts to `null`.
 *
 * On the server, this store can only be subscribed to during component initialization. In the browser, it can be subscribed to at any time.
 *
 * @deprecated Use `navigating` from `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
 * @type {import('svelte/store').Readable<import('@sveltejs/kit').Navigation | null>}
 */
export const navigating = {
	subscribe(fn) {
		const store = __SVELTEKIT_DEV__ ? get_store('navigating') : getStores().navigating;
		return store.subscribe(fn);
	}
};

/**
 * A readable store whose initial value is `false`. If [`version.pollInterval`](https://svelte.dev/docs/kit/configuration#version) is a non-zero value, SvelteKit will poll for new versions of the app and update the store value to `true` when it detects one. `updated.check()` will force an immediate check, regardless of polling.
 *
 * On the server, this store can only be subscribed to during component initialization. In the browser, it can be subscribed to at any time.
 *
 * @deprecated Use `updated` from `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
 * @type {import('svelte/store').Readable<boolean> & { check(): Promise<boolean> }}
 */
export const updated = {
	subscribe(fn) {
		const store = __SVELTEKIT_DEV__ ? get_store('updated') : getStores().updated;

		if (BROWSER) {
			updated.check = store.check;
		}

		return store.subscribe(fn);
	},
	check: () => {
		throw new Error(
			BROWSER
				? 'Cannot check updated store before subscribing'
				: 'Can only check updated store in browser'
		);
	}
};

/**
 * @template {keyof ReturnType<typeof getStores>} Name
 * @param {Name} name
 * @returns {ReturnType<typeof getStores>[Name]}
 */
function get_store(name) {
	try {
		return getStores()[name];
	} catch {
		throw new Error(
			`Cannot subscribe to '${name}' store on the server outside of a Svelte component, as it is bound to the current request via component context. This prevents state from leaking between users.` +
				'For more information, see https://svelte.dev/docs/kit/state-management#avoid-shared-state-on-the-server'
		);
	}
}
