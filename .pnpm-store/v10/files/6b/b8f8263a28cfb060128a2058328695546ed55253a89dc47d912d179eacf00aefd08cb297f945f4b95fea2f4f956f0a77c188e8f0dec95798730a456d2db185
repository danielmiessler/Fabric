import {
	page as client_page,
	navigating as client_navigating,
	updated as client_updated
} from './client.js';
import {
	page as server_page,
	navigating as server_navigating,
	updated as server_updated
} from './server.js';
import { BROWSER } from 'esm-env';

/**
 * A read-only reactive object with information about the current page, serving several use cases:
 * - retrieving the combined `data` of all pages/layouts anywhere in your component tree (also see [loading data](https://svelte.dev/docs/kit/load))
 * - retrieving the current value of the `form` prop anywhere in your component tree (also see [form actions](https://svelte.dev/docs/kit/form-actions))
 * - retrieving the page state that was set through `goto`, `pushState` or `replaceState` (also see [goto](https://svelte.dev/docs/kit/$app-navigation#goto) and [shallow routing](https://svelte.dev/docs/kit/shallow-routing))
 * - retrieving metadata such as the URL you're on, the current route and its parameters, and whether or not there was an error
 *
 * ```svelte
 * <!--- file: +layout.svelte --->
 * <script>
 * 	import { page } from '$app/state';
 * </script>
 *
 * <p>Currently at {page.url.pathname}</p>
 *
 * {#if page.error}
 * 	<span class="red">Problem detected</span>
 * {:else}
 * 	<span class="small">All systems operational</span>
 * {/if}
 * ```
 *
 * Changes to `page` are available exclusively with runes. (The legacy reactivity syntax will not reflect any changes)
 *
 * ```svelte
 * <!--- file: +page.svelte --->
 * <script>
 * 	import { page } from '$app/state';
 * 	const id = $derived(page.params.id); // This will correctly update id for usage on this page
 * 	$: badId = page.params.id; // Do not use; will never update after initial load
 * </script>
 * ```
 *
 * On the server, values can only be read during rendering (in other words _not_ in e.g. `load` functions). In the browser, the values can be read at any time.
 *
 * @type {import('@sveltejs/kit').Page}
 */
export const page = BROWSER ? client_page : server_page;

/**
 * A read-only object representing an in-progress navigation, with `from`, `to`, `type` and (if `type === 'popstate'`) `delta` properties.
 * Values are `null` when no navigation is occurring, or during server rendering.
 * @type {import('@sveltejs/kit').Navigation | { from: null, to: null, type: null, willUnload: null, delta: null, complete: null }}
 */
// @ts-expect-error
export const navigating = BROWSER ? client_navigating : server_navigating;

/**
 * A read-only reactive value that's initially `false`. If [`version.pollInterval`](https://svelte.dev/docs/kit/configuration#version) is a non-zero value, SvelteKit will poll for new versions of the app and update `current` to `true` when it detects one. `updated.check()` will force an immediate check, regardless of polling.
 * @type {{ get current(): boolean; check(): Promise<boolean>; }}
 */
export const updated = BROWSER ? client_updated : server_updated;
