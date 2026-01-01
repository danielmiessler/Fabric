import { RouteId, Pathname, ResolvedPathname } from '$app/types';
import { ResolveArgs } from './types.js';

export { resolve, asset } from './client.js';

/**
 * A string that matches [`config.kit.paths.base`](https://svelte.dev/docs/kit/configuration#paths).
 *
 * Example usage: `<a href="{base}/your-page">Link</a>`
 *
 * @deprecated Use [`resolve(...)`](https://svelte.dev/docs/kit/$app-paths#resolve) instead
 */
export let base: '' | `/${string}`;

/**
 * An absolute path that matches [`config.kit.paths.assets`](https://svelte.dev/docs/kit/configuration#paths).
 *
 * > [!NOTE] If a value for `config.kit.paths.assets` is specified, it will be replaced with `'/_svelte_kit_assets'` during `vite dev` or `vite preview`, since the assets don't yet live at their eventual URL.
 *
 * @deprecated Use [`asset(...)`](https://svelte.dev/docs/kit/$app-paths#asset) instead
 */
export let assets: '' | `https://${string}` | `http://${string}` | '/_svelte_kit_assets';

/**
 * @deprecated Use [`resolve(...)`](https://svelte.dev/docs/kit/$app-paths#resolve) instead
 */
export function resolveRoute<T extends RouteId | Pathname>(
	...args: ResolveArgs<T>
): ResolvedPathname;
