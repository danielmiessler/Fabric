import { onMount } from 'svelte';
import { updated_listener } from './utils.js';

/** @type {import('@sveltejs/kit').Page} */
export let page;

/** @type {{ current: import('@sveltejs/kit').Navigation | null }} */
export let navigating;

/** @type {{ current: boolean }} */
export let updated;

// this is a bootleg way to tell if we're in old svelte or new svelte
const is_legacy =
	onMount.toString().includes('$$') || /function \w+\(\) \{\}/.test(onMount.toString());

if (is_legacy) {
	page = {
		data: {},
		form: null,
		error: null,
		params: {},
		route: { id: null },
		state: {},
		status: -1,
		url: new URL('https://example.com')
	};
	navigating = { current: null };
	updated = { current: false };
} else {
	page = new (class Page {
		data = $state.raw({});
		form = $state.raw(null);
		error = $state.raw(null);
		params = $state.raw({});
		route = $state.raw({ id: null });
		state = $state.raw({});
		status = $state.raw(-1);
		url = $state.raw(new URL('https://example.com'));
	})();

	navigating = new (class Navigating {
		current = $state.raw(null);
	})();

	updated = new (class Updated {
		current = $state.raw(false);
	})();
	updated_listener.v = () => (updated.current = true);
}

/**
 * @param {import('@sveltejs/kit').Page} new_page
 */
export function update(new_page) {
	Object.assign(page, new_page);
}
