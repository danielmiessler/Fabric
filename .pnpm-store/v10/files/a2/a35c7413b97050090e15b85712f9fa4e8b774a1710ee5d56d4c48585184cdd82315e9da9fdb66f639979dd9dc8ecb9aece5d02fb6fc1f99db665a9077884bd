/* if `bundleStrategy` is 'single' or 'inline', this file is used as the entry point */

import * as kit from './entry.js';

// @ts-expect-error
import * as app from '__sveltekit/manifest';

/**
 *
 * @param {HTMLElement} element
 * @param {import('./types.js').HydrateOptions} options
 */
export function start(element, options) {
	void kit.start(app, element, options);
}
