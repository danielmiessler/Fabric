import { BROWSER, DEV } from 'esm-env';
import { onMount, tick } from 'svelte';
import {
	decode_params,
	decode_pathname,
	strip_hash,
	make_trackable,
	normalize_path
} from '../../utils/url.js';
import { dev_fetch, initial_fetch, lock_fetch, subsequent_fetch, unlock_fetch } from './fetcher.js';
import { parse, parse_server_route } from './parse.js';
import * as storage from './session-storage.js';
import {
	find_anchor,
	resolve_url,
	get_link_info,
	get_router_options,
	is_external_url,
	origin,
	scroll_state,
	notifiable_store,
	create_updated_store,
	load_css
} from './utils.js';
import { base } from '__sveltekit/paths';
import * as devalue from 'devalue';
import {
	HISTORY_INDEX,
	NAVIGATION_INDEX,
	PRELOAD_PRIORITIES,
	SCROLL_KEY,
	STATES_KEY,
	SNAPSHOT_KEY,
	PAGE_URL_KEY
} from './constants.js';
import { validate_page_exports } from '../../utils/exports.js';
import { compact } from '../../utils/array.js';
import { HttpError, Redirect, SvelteKitError } from '../control.js';
import { INVALIDATED_PARAM, TRAILING_SLASH_PARAM, validate_depends } from '../shared.js';
import { get_message, get_status } from '../../utils/error.js';
import { writable } from 'svelte/store';
import { page, update, navigating } from './state.svelte.js';
import { add_data_suffix, add_resolution_suffix } from '../pathname.js';

export { load_css };

const ICON_REL_ATTRIBUTES = new Set(['icon', 'shortcut icon', 'apple-touch-icon']);

let errored = false;

// We track the scroll position associated with each history entry in sessionStorage,
// rather than on history.state itself, because when navigation is driven by
// popstate it's too late to update the scroll position associated with the
// state we're navigating from
/**
 * history index -> { x, y }
 * @type {Record<number, { x: number; y: number }>}
 */
const scroll_positions = storage.get(SCROLL_KEY) ?? {};

/**
 * navigation index -> any
 * @type {Record<string, any[]>}
 */
const snapshots = storage.get(SNAPSHOT_KEY) ?? {};

if (DEV && BROWSER) {
	let warned = false;

	const current_module_url = import.meta.url.split('?')[0]; // remove query params that vite adds to the URL when it is loaded from node_modules

	const warn = () => {
		if (warned) return;

		// Rather than saving a pointer to the original history methods, which would prevent monkeypatching by other libs,
		// inspect the stack trace to see if we're being called from within SvelteKit.
		let stack = new Error().stack?.split('\n');
		if (!stack) return;
		if (!stack[0].includes('https:') && !stack[0].includes('http:')) stack = stack.slice(1); // Chrome includes the error message in the stack
		stack = stack.slice(2); // remove `warn` and the place where `warn` was called
		// Can be falsy if was called directly from an anonymous function
		if (stack[0]?.includes(current_module_url)) return;

		warned = true;

		console.warn(
			"Avoid using `history.pushState(...)` and `history.replaceState(...)` as these will conflict with SvelteKit's router. Use the `pushState` and `replaceState` imports from `$app/navigation` instead."
		);
	};

	const push_state = history.pushState;
	history.pushState = (...args) => {
		warn();
		return push_state.apply(history, args);
	};

	const replace_state = history.replaceState;
	history.replaceState = (...args) => {
		warn();
		return replace_state.apply(history, args);
	};
}

export const stores = {
	url: /* @__PURE__ */ notifiable_store({}),
	page: /* @__PURE__ */ notifiable_store({}),
	navigating: /* @__PURE__ */ writable(
		/** @type {import('@sveltejs/kit').Navigation | null} */ (null)
	),
	updated: /* @__PURE__ */ create_updated_store()
};

/** @param {number} index */
function update_scroll_positions(index) {
	scroll_positions[index] = scroll_state();
}

/**
 * @param {number} current_history_index
 * @param {number} current_navigation_index
 */
function clear_onward_history(current_history_index, current_navigation_index) {
	// if we navigated back, then pushed a new state, we can
	// release memory by pruning the scroll/snapshot lookup
	let i = current_history_index + 1;
	while (scroll_positions[i]) {
		delete scroll_positions[i];
		i += 1;
	}

	i = current_navigation_index + 1;
	while (snapshots[i]) {
		delete snapshots[i];
		i += 1;
	}
}

/**
 * Loads `href` the old-fashioned way, with a full page reload.
 * Returns a `Promise` that never resolves (to prevent any
 * subsequent work, e.g. history manipulation, from happening)
 * @param {URL} url
 */
function native_navigation(url) {
	location.href = url.href;
	return new Promise(() => {});
}

/**
 * Checks whether a service worker is registered, and if it is,
 * tries to update it.
 */
async function update_service_worker() {
	if ('serviceWorker' in navigator) {
		const registration = await navigator.serviceWorker.getRegistration(base || '/');
		if (registration) {
			await registration.update();
		}
	}
}

function noop() {}

/** @type {import('types').CSRRoute[]} All routes of the app. Only available when kit.router.resolution=client */
let routes;
/** @type {import('types').CSRPageNodeLoader} */
let default_layout_loader;
/** @type {import('types').CSRPageNodeLoader} */
let default_error_loader;
/** @type {HTMLElement} */
let container;
/** @type {HTMLElement} */
let target;
/** @type {import('./types.js').SvelteKitApp} */
export let app;

/** @type {Array<((url: URL) => boolean)>} */
const invalidated = [];

/**
 * An array of the `+layout.svelte` and `+page.svelte` component instances
 * that currently live on the page — used for capturing and restoring snapshots.
 * It's updated/manipulated through `bind:this` in `Root.svelte`.
 * @type {import('svelte').SvelteComponent[]}
 */
const components = [];

/** @type {{id: string, token: {}, promise: Promise<import('./types.js').NavigationResult>} | null} */
let load_cache = null;

/**
 * @type {Map<string, Promise<URL>>}
 * Cache for client-side rerouting, since it could contain async calls which we want to
 * avoid running multiple times which would slow down navigations (e.g. else preloading
 * wouldn't help because on navigation it would be called again). Since `reroute` should be
 * a pure function (i.e. always return the same) value it's safe to cache across navigations.
 * The server reroute calls don't need to be cached because they are called using `import(...)`
 * which is cached per the JS spec.
 */
const reroute_cache = new Map();

/**
 * Note on before_navigate_callbacks, on_navigate_callbacks and after_navigate_callbacks:
 * do not re-assign as some closures keep references to these Sets
 */
/** @type {Set<(navigation: import('@sveltejs/kit').BeforeNavigate) => void>} */
const before_navigate_callbacks = new Set();

/** @type {Set<(navigation: import('@sveltejs/kit').OnNavigate) => import('types').MaybePromise<(() => void) | void>>} */
const on_navigate_callbacks = new Set();

/** @type {Set<(navigation: import('@sveltejs/kit').AfterNavigate) => void>} */
const after_navigate_callbacks = new Set();

/** @type {import('./types.js').NavigationState} */
let current = {
	branch: [],
	error: null,
	// @ts-ignore - we need the initial value to be null
	url: null
};

/** this being true means we SSR'd */
let hydrated = false;
let started = false;
let autoscroll = true;
let updating = false;
let is_navigating = false;
let hash_navigating = false;
/** True as soon as there happened one client-side navigation (excluding the SvelteKit-initialized initial one when in SPA mode) */
let has_navigated = false;

let force_invalidation = false;

/** @type {import('svelte').SvelteComponent} */
let root;

/** @type {number} keeping track of the history index in order to prevent popstate navigation events if needed */
let current_history_index;

/** @type {number} */
let current_navigation_index;

/** @type {{}} */
let token;

/**
 * A set of tokens which are associated to current preloads.
 * If a preload becomes a real navigation, it's removed from the set.
 * If a preload token is in the set and the preload errors, the error
 * handling logic (for example reloading) is skipped.
 */
const preload_tokens = new Set();

/** @type {Promise<void> | null} */
let pending_invalidate;

/**
 * @param {import('./types.js').SvelteKitApp} _app
 * @param {HTMLElement} _target
 * @param {Parameters<typeof _hydrate>[1]} [hydrate]
 */
export async function start(_app, _target, hydrate) {
	if (DEV && _target === document.body) {
		console.warn(
			'Placing %sveltekit.body% directly inside <body> is not recommended, as your app may break for users who have certain browser extensions installed.\n\nConsider wrapping it in an element:\n\n<div style="display: contents">\n  %sveltekit.body%\n</div>'
		);
	}

	// detect basic auth credentials in the current URL
	// https://github.com/sveltejs/kit/pull/11179
	// if so, refresh the page without credentials
	if (document.URL !== location.href) {
		// eslint-disable-next-line no-self-assign
		location.href = location.href;
	}

	app = _app;

	await _app.hooks.init?.();

	routes = __SVELTEKIT_CLIENT_ROUTING__ ? parse(_app) : [];
	container = __SVELTEKIT_EMBEDDED__ ? _target : document.documentElement;
	target = _target;

	// we import the root layout/error nodes eagerly, so that
	// connectivity errors after initialisation don't nuke the app
	default_layout_loader = _app.nodes[0];
	default_error_loader = _app.nodes[1];
	void default_layout_loader();
	void default_error_loader();

	current_history_index = history.state?.[HISTORY_INDEX];
	current_navigation_index = history.state?.[NAVIGATION_INDEX];

	if (!current_history_index) {
		// we use Date.now() as an offset so that cross-document navigations
		// within the app don't result in data loss
		current_history_index = current_navigation_index = Date.now();

		// create initial history entry, so we can return here
		history.replaceState(
			{
				...history.state,
				[HISTORY_INDEX]: current_history_index,
				[NAVIGATION_INDEX]: current_navigation_index
			},
			''
		);
	}

	// if we reload the page, or Cmd-Shift-T back to it,
	// recover scroll position
	const scroll = scroll_positions[current_history_index];
	if (scroll) {
		history.scrollRestoration = 'manual';
		scrollTo(scroll.x, scroll.y);
	}

	if (hydrate) {
		await _hydrate(target, hydrate);
	} else {
		await navigate({
			type: 'enter',
			url: resolve_url(app.hash ? decode_hash(new URL(location.href)) : location.href),
			replace_state: true
		});
	}

	_start_router();
}

async function _invalidate() {
	// Accept all invalidations as they come, don't swallow any while another invalidation
	// is running because subsequent invalidations may make earlier ones outdated,
	// but batch multiple synchronous invalidations.
	await (pending_invalidate ||= Promise.resolve());
	if (!pending_invalidate) return;
	pending_invalidate = null;

	const nav_token = (token = {});
	const intent = await get_navigation_intent(current.url, true);

	// Clear preload, it might be affected by the invalidation.
	// Also solves an edge case where a preload is triggered, the navigation for it
	// was then triggered and is still running while the invalidation kicks in,
	// at which point the invalidation should take over and "win".
	load_cache = null;

	const navigation_result = intent && (await load_route(intent));
	if (!navigation_result || nav_token !== token) return;

	if (navigation_result.type === 'redirect') {
		return _goto(new URL(navigation_result.location, current.url).href, {}, 1, nav_token);
	}

	if (navigation_result.props.page) {
		Object.assign(page, navigation_result.props.page);
	}
	current = navigation_result.state;
	reset_invalidation();
	root.$set(navigation_result.props);
	update(navigation_result.props.page);
}

function reset_invalidation() {
	invalidated.length = 0;
	force_invalidation = false;
}

/** @param {number} index */
function capture_snapshot(index) {
	if (components.some((c) => c?.snapshot)) {
		snapshots[index] = components.map((c) => c?.snapshot?.capture());
	}
}

/** @param {number} index */
function restore_snapshot(index) {
	snapshots[index]?.forEach((value, i) => {
		components[i]?.snapshot?.restore(value);
	});
}

function persist_state() {
	update_scroll_positions(current_history_index);
	storage.set(SCROLL_KEY, scroll_positions);

	capture_snapshot(current_navigation_index);
	storage.set(SNAPSHOT_KEY, snapshots);
}

/**
 * @param {string | URL} url
 * @param {{ replaceState?: boolean; noScroll?: boolean; keepFocus?: boolean; invalidateAll?: boolean; invalidate?: Array<string | URL | ((url: URL) => boolean)>; state?: Record<string, any> }} options
 * @param {number} redirect_count
 * @param {{}} [nav_token]
 */
async function _goto(url, options, redirect_count, nav_token) {
	return navigate({
		type: 'goto',
		url: resolve_url(url),
		keepfocus: options.keepFocus,
		noscroll: options.noScroll,
		replace_state: options.replaceState,
		state: options.state,
		redirect_count,
		nav_token,
		accept: () => {
			if (options.invalidateAll) {
				force_invalidation = true;
			}

			if (options.invalidate) {
				options.invalidate.forEach(push_invalidated);
			}
		}
	});
}

/** @param {import('./types.js').NavigationIntent} intent */
async function _preload_data(intent) {
	// Reuse the existing pending preload if it's for the same navigation.
	// Prevents an edge case where same preload is triggered multiple times,
	// then a later one is becoming the real navigation and the preload tokens
	// get out of sync.
	if (intent.id !== load_cache?.id) {
		const preload = {};
		preload_tokens.add(preload);
		load_cache = {
			id: intent.id,
			token: preload,
			promise: load_route({ ...intent, preload }).then((result) => {
				preload_tokens.delete(preload);
				if (result.type === 'loaded' && result.state.error) {
					// Don't cache errors, because they might be transient
					load_cache = null;
				}
				return result;
			})
		};
	}

	return load_cache.promise;
}

/**
 * @param {URL} url
 * @returns {Promise<void>}
 */
async function _preload_code(url) {
	const route = (await get_navigation_intent(url, false))?.route;

	if (route) {
		await Promise.all([...route.layouts, route.leaf].map((load) => load?.[1]()));
	}
}

/**
 * @param {import('./types.js').NavigationFinished} result
 * @param {HTMLElement} target
 * @param {boolean} hydrate
 */
function initialize(result, target, hydrate) {
	if (DEV && result.state.error && document.querySelector('vite-error-overlay')) return;

	current = result.state;

	const style = document.querySelector('style[data-sveltekit]');
	if (style) style.remove();

	Object.assign(page, /** @type {import('@sveltejs/kit').Page} */ (result.props.page));

	root = new app.root({
		target,
		props: { ...result.props, stores, components },
		hydrate,
		// @ts-ignore Svelte 5 specific: asynchronously instantiate the component, i.e. don't call flushSync
		sync: false
	});

	restore_snapshot(current_navigation_index);

	if (hydrate) {
		/** @type {import('@sveltejs/kit').AfterNavigate} */
		const navigation = {
			from: null,
			to: {
				params: current.params,
				route: { id: current.route?.id ?? null },
				url: new URL(location.href)
			},
			willUnload: false,
			type: 'enter',
			complete: Promise.resolve()
		};

		after_navigate_callbacks.forEach((fn) => fn(navigation));
	}

	started = true;
}

/**
 *
 * @param {{
 *   url: URL;
 *   params: Record<string, string>;
 *   branch: Array<import('./types.js').BranchNode | undefined>;
 *   status: number;
 *   error: App.Error | null;
 *   route: import('types').CSRRoute | null;
 *   form?: Record<string, any> | null;
 * }} opts
 */
function get_navigation_result_from_branch({ url, params, branch, status, error, route, form }) {
	/** @type {import('types').TrailingSlash} */
	let slash = 'never';

	// if `paths.base === '/a/b/c`, then the root route is always `/a/b/c/`, regardless of
	// the `trailingSlash` route option, so that relative paths to JS and CSS work
	if (base && (url.pathname === base || url.pathname === base + '/')) {
		slash = 'always';
	} else {
		for (const node of branch) {
			if (node?.slash !== undefined) slash = node.slash;
		}
	}

	url.pathname = normalize_path(url.pathname, slash);

	// eslint-disable-next-line
	url.search = url.search; // turn `/?` into `/`

	/** @type {import('./types.js').NavigationFinished} */
	const result = {
		type: 'loaded',
		state: {
			url,
			params,
			branch,
			error,
			route
		},
		props: {
			// @ts-ignore Somehow it's getting SvelteComponent and SvelteComponentDev mixed up
			constructors: compact(branch).map((branch_node) => branch_node.node.component),
			page: clone_page(page)
		}
	};

	if (form !== undefined) {
		result.props.form = form;
	}

	let data = {};
	let data_changed = !page;

	let p = 0;

	for (let i = 0; i < Math.max(branch.length, current.branch.length); i += 1) {
		const node = branch[i];
		const prev = current.branch[i];

		if (node?.data !== prev?.data) data_changed = true;
		if (!node) continue;

		data = { ...data, ...node.data };

		// Only set props if the node actually updated. This prevents needless rerenders.
		if (data_changed) {
			result.props[`data_${p}`] = data;
		}

		p += 1;
	}

	const page_changed =
		!current.url ||
		url.href !== current.url.href ||
		current.error !== error ||
		(form !== undefined && form !== page.form) ||
		data_changed;

	if (page_changed) {
		result.props.page = {
			error,
			params,
			route: {
				id: route?.id ?? null
			},
			state: {},
			status,
			url: new URL(url),
			form: form ?? null,
			// The whole page store is updated, but this way the object reference stays the same
			data: data_changed ? data : page.data
		};
	}

	return result;
}

/**
 * Call the universal load function of the given node, if it exists.
 *
 * @param {{
 *   loader: import('types').CSRPageNodeLoader;
 * 	 parent: () => Promise<Record<string, any>>;
 *   url: URL;
 *   params: Record<string, string>;
 *   route: { id: string | null };
 * 	 server_data_node: import('./types.js').DataNode | null;
 * }} options
 * @returns {Promise<import('./types.js').BranchNode>}
 */
async function load_node({ loader, parent, url, params, route, server_data_node }) {
	/** @type {Record<string, any> | null} */
	let data = null;

	let is_tracking = true;

	/** @type {import('types').Uses} */
	const uses = {
		dependencies: new Set(),
		params: new Set(),
		parent: false,
		route: false,
		url: false,
		search_params: new Set()
	};

	const node = await loader();

	if (DEV) {
		validate_page_exports(node.universal);

		if (node.universal && app.hash) {
			const options = Object.keys(node.universal).filter((o) => o !== 'load');

			if (options.length > 0) {
				throw new Error(
					`Page options are ignored when \`router.type === 'hash'\` (${route.id} has ${options
						.filter((o) => o !== 'load')
						.map((o) => `'${o}'`)
						.join(', ')})`
				);
			}
		}
	}

	if (node.universal?.load) {
		/** @param {string[]} deps */
		function depends(...deps) {
			for (const dep of deps) {
				if (DEV) validate_depends(/** @type {string} */ (route.id), dep);

				const { href } = new URL(dep, url);
				uses.dependencies.add(href);
			}
		}

		/** @type {import('@sveltejs/kit').LoadEvent} */
		const load_input = {
			route: new Proxy(route, {
				get: (target, key) => {
					if (is_tracking) {
						uses.route = true;
					}
					return target[/** @type {'id'} */ (key)];
				}
			}),
			params: new Proxy(params, {
				get: (target, key) => {
					if (is_tracking) {
						uses.params.add(/** @type {string} */ (key));
					}
					return target[/** @type {string} */ (key)];
				}
			}),
			data: server_data_node?.data ?? null,
			url: make_trackable(
				url,
				() => {
					if (is_tracking) {
						uses.url = true;
					}
				},
				(param) => {
					if (is_tracking) {
						uses.search_params.add(param);
					}
				},
				app.hash
			),
			async fetch(resource, init) {
				if (resource instanceof Request) {
					// we're not allowed to modify the received `Request` object, so in order
					// to fixup relative urls we create a new equivalent `init` object instead
					init = {
						// the request body must be consumed in memory until browsers
						// implement streaming request bodies and/or the body getter
						body:
							resource.method === 'GET' || resource.method === 'HEAD'
								? undefined
								: await resource.blob(),
						cache: resource.cache,
						credentials: resource.credentials,
						// the headers are undefined on the server if the Headers object is empty
						// so we need to make sure they are also undefined here if there are no headers
						headers: [...resource.headers].length ? resource.headers : undefined,
						integrity: resource.integrity,
						keepalive: resource.keepalive,
						method: resource.method,
						mode: resource.mode,
						redirect: resource.redirect,
						referrer: resource.referrer,
						referrerPolicy: resource.referrerPolicy,
						signal: resource.signal,
						...init
					};
				}

				const { resolved, promise } = resolve_fetch_url(resource, init, url);

				if (is_tracking) {
					depends(resolved.href);
				}

				return promise;
			},
			setHeaders: () => {}, // noop
			depends,
			parent() {
				if (is_tracking) {
					uses.parent = true;
				}
				return parent();
			},
			untrack(fn) {
				is_tracking = false;
				try {
					return fn();
				} finally {
					is_tracking = true;
				}
			}
		};

		if (DEV) {
			try {
				lock_fetch();
				data = (await node.universal.load.call(null, load_input)) ?? null;
				if (data != null && Object.getPrototypeOf(data) !== Object.prototype) {
					throw new Error(
						`a load function related to route '${route.id}' returned ${
							typeof data !== 'object'
								? `a ${typeof data}`
								: data instanceof Response
									? 'a Response object'
									: Array.isArray(data)
										? 'an array'
										: 'a non-plain object'
						}, but must return a plain object at the top level (i.e. \`return {...}\`)`
					);
				}
			} finally {
				unlock_fetch();
			}
		} else {
			data = (await node.universal.load.call(null, load_input)) ?? null;
		}
	}

	return {
		node,
		loader,
		server: server_data_node,
		universal: node.universal?.load ? { type: 'data', data, uses } : null,
		data: data ?? server_data_node?.data ?? null,
		slash: node.universal?.trailingSlash ?? server_data_node?.slash
	};
}

/**
 * @param {Request | string | URL} input
 * @param {RequestInit | undefined} init
 * @param {URL} url
 */
function resolve_fetch_url(input, init, url) {
	let requested = input instanceof Request ? input.url : input;

	// we must fixup relative urls so they are resolved from the target page
	const resolved = new URL(requested, url);

	// match ssr serialized data url, which is important to find cached responses
	if (resolved.origin === url.origin) {
		requested = resolved.href.slice(url.origin.length);
	}

	// prerendered pages may be served from any origin, so `initial_fetch` urls shouldn't be resolved
	const promise = started
		? subsequent_fetch(requested, resolved.href, init)
		: initial_fetch(requested, init);

	return { resolved, promise };
}

/**
 * @param {boolean} parent_changed
 * @param {boolean} route_changed
 * @param {boolean} url_changed
 * @param {Set<string>} search_params_changed
 * @param {import('types').Uses | undefined} uses
 * @param {Record<string, string>} params
 */
function has_changed(
	parent_changed,
	route_changed,
	url_changed,
	search_params_changed,
	uses,
	params
) {
	if (force_invalidation) return true;

	if (!uses) return false;

	if (uses.parent && parent_changed) return true;
	if (uses.route && route_changed) return true;
	if (uses.url && url_changed) return true;

	for (const tracked_params of uses.search_params) {
		if (search_params_changed.has(tracked_params)) return true;
	}

	for (const param of uses.params) {
		if (params[param] !== current.params[param]) return true;
	}

	for (const href of uses.dependencies) {
		if (invalidated.some((fn) => fn(new URL(href)))) return true;
	}

	return false;
}

/**
 * @param {import('types').ServerDataNode | import('types').ServerDataSkippedNode | null} node
 * @param {import('./types.js').DataNode | null} [previous]
 * @returns {import('./types.js').DataNode | null}
 */
function create_data_node(node, previous) {
	if (node?.type === 'data') return node;
	if (node?.type === 'skip') return previous ?? null;
	return null;
}

/**
 * @param {URL | null} old_url
 * @param {URL} new_url
 */
function diff_search_params(old_url, new_url) {
	if (!old_url) return new Set(new_url.searchParams.keys());

	const changed = new Set([...old_url.searchParams.keys(), ...new_url.searchParams.keys()]);

	for (const key of changed) {
		const old_values = old_url.searchParams.getAll(key);
		const new_values = new_url.searchParams.getAll(key);

		if (
			old_values.every((value) => new_values.includes(value)) &&
			new_values.every((value) => old_values.includes(value))
		) {
			changed.delete(key);
		}
	}

	return changed;
}

/**
 * @param {Omit<import('./types.js').NavigationFinished['state'], 'branch'> & { error: App.Error }} opts
 * @returns {import('./types.js').NavigationFinished}
 */
function preload_error({ error, url, route, params }) {
	return {
		type: 'loaded',
		state: {
			error,
			url,
			route,
			params,
			branch: []
		},
		props: {
			page: clone_page(page),
			constructors: []
		}
	};
}

/**
 * @param {import('./types.js').NavigationIntent & { preload?: {} }} intent
 * @returns {Promise<import('./types.js').NavigationResult>}
 */
async function load_route({ id, invalidating, url, params, route, preload }) {
	if (load_cache?.id === id) {
		// the preload becomes the real navigation
		preload_tokens.delete(load_cache.token);
		return load_cache.promise;
	}

	const { errors, layouts, leaf } = route;

	const loaders = [...layouts, leaf];

	// preload modules to avoid waterfall, but handle rejections
	// so they don't get reported to Sentry et al (we don't need
	// to act on the failures at this point)
	errors.forEach((loader) => loader?.().catch(() => {}));
	loaders.forEach((loader) => loader?.[1]().catch(() => {}));

	/** @type {import('types').ServerNodesResponse | import('types').ServerRedirectNode | null} */
	let server_data = null;
	const url_changed = current.url ? id !== get_page_key(current.url) : false;
	const route_changed = current.route ? route.id !== current.route.id : false;
	const search_params_changed = diff_search_params(current.url, url);

	let parent_invalid = false;
	const invalid_server_nodes = loaders.map((loader, i) => {
		const previous = current.branch[i];

		const invalid =
			!!loader?.[0] &&
			(previous?.loader !== loader[1] ||
				has_changed(
					parent_invalid,
					route_changed,
					url_changed,
					search_params_changed,
					previous.server?.uses,
					params
				));

		if (invalid) {
			// For the next one
			parent_invalid = true;
		}

		return invalid;
	});

	if (invalid_server_nodes.some(Boolean)) {
		try {
			server_data = await load_data(url, invalid_server_nodes);
		} catch (error) {
			const handled_error = await handle_error(error, { url, params, route: { id } });

			if (preload_tokens.has(preload)) {
				return preload_error({ error: handled_error, url, params, route });
			}

			return load_root_error_page({
				status: get_status(error),
				error: handled_error,
				url,
				route
			});
		}

		if (server_data.type === 'redirect') {
			return server_data;
		}
	}

	const server_data_nodes = server_data?.nodes;

	let parent_changed = false;

	const branch_promises = loaders.map(async (loader, i) => {
		if (!loader) return;

		/** @type {import('./types.js').BranchNode | undefined} */
		const previous = current.branch[i];

		const server_data_node = server_data_nodes?.[i];

		// re-use data from previous load if it's still valid
		const valid =
			(!server_data_node || server_data_node.type === 'skip') &&
			loader[1] === previous?.loader &&
			!has_changed(
				parent_changed,
				route_changed,
				url_changed,
				search_params_changed,
				previous.universal?.uses,
				params
			);
		if (valid) return previous;

		parent_changed = true;

		if (server_data_node?.type === 'error') {
			// rethrow and catch below
			throw server_data_node;
		}

		return load_node({
			loader: loader[1],
			url,
			params,
			route,
			parent: async () => {
				const data = {};
				for (let j = 0; j < i; j += 1) {
					Object.assign(data, (await branch_promises[j])?.data);
				}
				return data;
			},
			server_data_node: create_data_node(
				// server_data_node is undefined if it wasn't reloaded from the server;
				// and if current loader uses server data, we want to reuse previous data.
				server_data_node === undefined && loader[0] ? { type: 'skip' } : (server_data_node ?? null),
				loader[0] ? previous?.server : undefined
			)
		});
	});

	// if we don't do this, rejections will be unhandled
	for (const p of branch_promises) p.catch(() => {});

	/** @type {Array<import('./types.js').BranchNode | undefined>} */
	const branch = [];

	for (let i = 0; i < loaders.length; i += 1) {
		if (loaders[i]) {
			try {
				branch.push(await branch_promises[i]);
			} catch (err) {
				if (err instanceof Redirect) {
					return {
						type: 'redirect',
						location: err.location
					};
				}

				if (preload_tokens.has(preload)) {
					return preload_error({
						error: await handle_error(err, { params, url, route: { id: route.id } }),
						url,
						params,
						route
					});
				}

				let status = get_status(err);
				/** @type {App.Error} */
				let error;

				if (server_data_nodes?.includes(/** @type {import('types').ServerErrorNode} */ (err))) {
					// this is the server error rethrown above, reconstruct but don't invoke
					// the client error handler; it should've already been handled on the server
					status = /** @type {import('types').ServerErrorNode} */ (err).status ?? status;
					error = /** @type {import('types').ServerErrorNode} */ (err).error;
				} else if (err instanceof HttpError) {
					error = err.body;
				} else {
					// Referenced node could have been removed due to redeploy, check
					const updated = await stores.updated.check();
					if (updated) {
						// Before reloading, try to update the service worker if it exists
						await update_service_worker();
						return await native_navigation(url);
					}

					error = await handle_error(err, { params, url, route: { id: route.id } });
				}

				const error_load = await load_nearest_error_page(i, branch, errors);
				if (error_load) {
					return get_navigation_result_from_branch({
						url,
						params,
						branch: branch.slice(0, error_load.idx).concat(error_load.node),
						status,
						error,
						route
					});
				} else {
					return await server_fallback(url, { id: route.id }, error, status);
				}
			}
		} else {
			// push an empty slot so we can rewind past gaps to the
			// layout that corresponds with an +error.svelte page
			branch.push(undefined);
		}
	}

	return get_navigation_result_from_branch({
		url,
		params,
		branch,
		status: 200,
		error: null,
		route,
		// Reset `form` on navigation, but not invalidation
		form: invalidating ? undefined : null
	});
}

/**
 * @param {number} i Start index to backtrack from
 * @param {Array<import('./types.js').BranchNode | undefined>} branch Branch to backtrack
 * @param {Array<import('types').CSRPageNodeLoader | undefined>} errors All error pages for this branch
 * @returns {Promise<{idx: number; node: import('./types.js').BranchNode} | undefined>}
 */
async function load_nearest_error_page(i, branch, errors) {
	while (i--) {
		if (errors[i]) {
			let j = i;
			while (!branch[j]) j -= 1;
			try {
				return {
					idx: j + 1,
					node: {
						node: await /** @type {import('types').CSRPageNodeLoader } */ (errors[i])(),
						loader: /** @type {import('types').CSRPageNodeLoader } */ (errors[i]),
						data: {},
						server: null,
						universal: null
					}
				};
			} catch {
				continue;
			}
		}
	}
}

/**
 * @param {{
 *   status: number;
 *   error: App.Error;
 *   url: URL;
 *   route: { id: string | null }
 * }} opts
 * @returns {Promise<import('./types.js').NavigationFinished>}
 */
async function load_root_error_page({ status, error, url, route }) {
	/** @type {Record<string, string>} */
	const params = {}; // error page does not have params

	/** @type {import('types').ServerDataNode | null} */
	let server_data_node = null;

	const default_layout_has_server_load = app.server_loads[0] === 0;

	if (default_layout_has_server_load) {
		// TODO post-https://github.com/sveltejs/kit/discussions/6124 we can use
		// existing root layout data
		try {
			const server_data = await load_data(url, [true]);

			if (
				server_data.type !== 'data' ||
				(server_data.nodes[0] && server_data.nodes[0].type !== 'data')
			) {
				throw 0;
			}

			server_data_node = server_data.nodes[0] ?? null;
		} catch {
			// at this point we have no choice but to fall back to the server, if it wouldn't
			// bring us right back here, turning this into an endless loop
			if (url.origin !== origin || url.pathname !== location.pathname || hydrated) {
				await native_navigation(url);
			}
		}
	}

	try {
		const root_layout = await load_node({
			loader: default_layout_loader,
			url,
			params,
			route,
			parent: () => Promise.resolve({}),
			server_data_node: create_data_node(server_data_node)
		});

		/** @type {import('./types.js').BranchNode} */
		const root_error = {
			node: await default_error_loader(),
			loader: default_error_loader,
			universal: null,
			server: null,
			data: null
		};

		return get_navigation_result_from_branch({
			url,
			params,
			branch: [root_layout, root_error],
			status,
			error,
			route: null
		});
	} catch (error) {
		if (error instanceof Redirect) {
			return _goto(new URL(error.location, location.href), {}, 0);
		}

		// TODO: this falls back to the server when a server exists, but what about SPA mode?
		throw error;
	}
}

/**
 * Resolve the relative rerouted URL for a client-side navigation
 * @param {URL} url
 * @returns {Promise<URL | undefined>}
 */
async function get_rerouted_url(url) {
	const href = url.href;

	if (reroute_cache.has(href)) {
		return reroute_cache.get(href);
	}

	let rerouted;

	try {
		const promise = (async () => {
			// reroute could alter the given URL, so we pass a copy
			let rerouted =
				(await app.hooks.reroute({
					url: new URL(url),
					fetch: async (input, init) => {
						return resolve_fetch_url(input, init, url).promise;
					}
				})) ?? url;

			if (typeof rerouted === 'string') {
				const tmp = new URL(url); // do not mutate the incoming URL

				if (app.hash) {
					tmp.hash = rerouted;
				} else {
					tmp.pathname = rerouted;
				}

				rerouted = tmp;
			}

			return rerouted;
		})();

		reroute_cache.set(href, promise);
		rerouted = await promise;
	} catch (e) {
		reroute_cache.delete(href);
		if (DEV) {
			// in development, print the error...
			console.error(e);

			// ...and pause execution, since otherwise we will immediately reload the page
			debugger; // eslint-disable-line
		}

		// fall back to native navigation
		return;
	}

	return rerouted;
}

/**
 * Resolve the full info (which route, params, etc.) for a client-side navigation from the URL,
 * taking the reroute hook into account. If this isn't a client-side-navigation (or the URL is undefined),
 * returns undefined.
 * @param {URL | undefined} url
 * @param {boolean} invalidating
 * @returns {Promise<import('./types.js').NavigationIntent | undefined>}
 */
async function get_navigation_intent(url, invalidating) {
	if (!url) return;
	if (is_external_url(url, base, app.hash)) return;

	if (__SVELTEKIT_CLIENT_ROUTING__) {
		const rerouted = await get_rerouted_url(url);
		if (!rerouted) return;

		const path = get_url_path(rerouted);

		for (const route of routes) {
			const params = route.exec(path);

			if (params) {
				return {
					id: get_page_key(url),
					invalidating,
					route,
					params: decode_params(params),
					url
				};
			}
		}
	} else {
		/** @type {{ route?: import('types').CSRRouteServer, params: Record<string, string>}} */
		const { route, params } = await import(
			/* @vite-ignore */
			add_resolution_suffix(url.pathname)
		);

		if (!route) return;

		return {
			id: get_page_key(url),
			invalidating,
			route: parse_server_route(route, app.nodes),
			params,
			url
		};
	}
}

/** @param {URL} url */
function get_url_path(url) {
	return (
		decode_pathname(
			app.hash ? url.hash.replace(/^#/, '').replace(/[?#].+/, '') : url.pathname.slice(base.length)
		) || '/'
	);
}

/** @param {URL} url */
function get_page_key(url) {
	return (app.hash ? url.hash.replace(/^#/, '') : url.pathname) + url.search;
}

/**
 * @param {{
 *   url: URL;
 *   type: import('@sveltejs/kit').Navigation["type"];
 *   intent?: import('./types.js').NavigationIntent;
 *   delta?: number;
 * }} opts
 */
function _before_navigate({ url, type, intent, delta }) {
	let should_block = false;

	const nav = create_navigation(current, intent, url, type);

	if (delta !== undefined) {
		nav.navigation.delta = delta;
	}

	const cancellable = {
		...nav.navigation,
		cancel: () => {
			should_block = true;
			nav.reject(new Error('navigation cancelled'));
		}
	};

	if (!is_navigating) {
		// Don't run the event during redirects
		before_navigate_callbacks.forEach((fn) => fn(cancellable));
	}

	return should_block ? null : nav;
}

/**
 * @param {{
 *   type: import('@sveltejs/kit').NavigationType;
 *   url: URL;
 *   popped?: {
 *     state: Record<string, any>;
 *     scroll: { x: number, y: number };
 *     delta: number;
 *   };
 *   keepfocus?: boolean;
 *   noscroll?: boolean;
 *   replace_state?: boolean;
 *   state?: Record<string, any>;
 *   redirect_count?: number;
 *   nav_token?: {};
 *   accept?: () => void;
 *   block?: () => void;
 * }} opts
 */
async function navigate({
	type,
	url,
	popped,
	keepfocus,
	noscroll,
	replace_state,
	state = {},
	redirect_count = 0,
	nav_token = {},
	accept = noop,
	block = noop
}) {
	const prev_token = token;
	token = nav_token;

	const intent = await get_navigation_intent(url, false);
	const nav =
		type === 'enter'
			? create_navigation(current, intent, url, type)
			: _before_navigate({ url, type, delta: popped?.delta, intent });

	if (!nav) {
		block();
		if (token === nav_token) token = prev_token;
		return;
	}

	// store this before calling `accept()`, which may change the index
	const previous_history_index = current_history_index;
	const previous_navigation_index = current_navigation_index;

	accept();

	is_navigating = true;

	if (started && nav.navigation.type !== 'enter') {
		stores.navigating.set((navigating.current = nav.navigation));
	}

	let navigation_result = intent && (await load_route(intent));

	if (!navigation_result) {
		if (is_external_url(url, base, app.hash)) {
			if (DEV && app.hash) {
				// Special case for hash mode during DEV: If someone accidentally forgets to use a hash for the link,
				// they would end up here in an endless loop. Fall back to error page in that case
				navigation_result = await server_fallback(
					url,
					{ id: null },
					await handle_error(
						new SvelteKitError(
							404,
							'Not Found',
							`Not found: ${url.pathname} (did you forget the hash?)`
						),
						{
							url,
							params: {},
							route: { id: null }
						}
					),
					404
				);
			} else {
				return await native_navigation(url);
			}
		} else {
			navigation_result = await server_fallback(
				url,
				{ id: null },
				await handle_error(new SvelteKitError(404, 'Not Found', `Not found: ${url.pathname}`), {
					url,
					params: {},
					route: { id: null }
				}),
				404
			);
		}
	}

	// if this is an internal navigation intent, use the normalized
	// URL for the rest of the function
	url = intent?.url || url;

	// abort if user navigated during update
	if (token !== nav_token) {
		nav.reject(new Error('navigation aborted'));
		return false;
	}

	if (navigation_result.type === 'redirect') {
		// whatwg fetch spec https://fetch.spec.whatwg.org/#http-redirect-fetch says to error after 20 redirects
		if (redirect_count >= 20) {
			navigation_result = await load_root_error_page({
				status: 500,
				error: await handle_error(new Error('Redirect loop'), {
					url,
					params: {},
					route: { id: null }
				}),
				url,
				route: { id: null }
			});
		} else {
			await _goto(new URL(navigation_result.location, url).href, {}, redirect_count + 1, nav_token);
			return false;
		}
	} else if (/** @type {number} */ (navigation_result.props.page.status) >= 400) {
		const updated = await stores.updated.check();
		if (updated) {
			// Before reloading, try to update the service worker if it exists
			await update_service_worker();
			await native_navigation(url);
		}
	}

	// reset invalidation only after a finished navigation. If there are redirects or
	// additional invalidations, they should get the same invalidation treatment
	reset_invalidation();

	updating = true;

	update_scroll_positions(previous_history_index);
	capture_snapshot(previous_navigation_index);

	// ensure the url pathname matches the page's trailing slash option
	if (navigation_result.props.page.url.pathname !== url.pathname) {
		url.pathname = navigation_result.props.page.url.pathname;
	}

	state = popped ? popped.state : state;

	if (!popped) {
		// this is a new navigation, rather than a popstate
		const change = replace_state ? 0 : 1;

		const entry = {
			[HISTORY_INDEX]: (current_history_index += change),
			[NAVIGATION_INDEX]: (current_navigation_index += change),
			[STATES_KEY]: state
		};

		const fn = replace_state ? history.replaceState : history.pushState;
		fn.call(history, entry, '', url);

		if (!replace_state) {
			clear_onward_history(current_history_index, current_navigation_index);
		}
	}

	// reset preload synchronously after the history state has been set to avoid race conditions
	load_cache = null;

	navigation_result.props.page.state = state;

	if (started) {
		current = navigation_result.state;

		// reset url before updating page store
		if (navigation_result.props.page) {
			navigation_result.props.page.url = url;
		}

		const after_navigate = (
			await Promise.all(
				Array.from(on_navigate_callbacks, (fn) =>
					fn(/** @type {import('@sveltejs/kit').OnNavigate} */ (nav.navigation))
				)
			)
		).filter(/** @returns {value is () => void} */ (value) => typeof value === 'function');

		if (after_navigate.length > 0) {
			function cleanup() {
				after_navigate.forEach((fn) => {
					after_navigate_callbacks.delete(fn);
				});
			}

			after_navigate.push(cleanup);

			after_navigate.forEach((fn) => {
				after_navigate_callbacks.add(fn);
			});
		}

		root.$set(navigation_result.props);
		update(navigation_result.props.page);
		has_navigated = true;
	} else {
		initialize(navigation_result, target, false);
	}

	const { activeElement } = document;

	// need to render the DOM before we can scroll to the rendered elements and do focus management
	await tick();

	// we reset scroll before dealing with focus, to avoid a flash of unscrolled content
	const scroll = popped ? popped.scroll : noscroll ? scroll_state() : null;

	if (autoscroll) {
		const deep_linked =
			url.hash &&
			document.getElementById(
				decodeURIComponent(app.hash ? (url.hash.split('#')[2] ?? '') : url.hash.slice(1))
			);
		if (scroll) {
			scrollTo(scroll.x, scroll.y);
		} else if (deep_linked) {
			// Here we use `scrollIntoView` on the element instead of `scrollTo`
			// because it natively supports the `scroll-margin` and `scroll-behavior`
			// CSS properties.
			deep_linked.scrollIntoView();
		} else {
			scrollTo(0, 0);
		}
	}

	const changed_focus =
		// reset focus only if any manual focus management didn't override it
		document.activeElement !== activeElement &&
		// also refocus when activeElement is body already because the
		// focus event might not have been fired on it yet
		document.activeElement !== document.body;

	if (!keepfocus && !changed_focus) {
		reset_focus();
	}

	autoscroll = true;

	if (navigation_result.props.page) {
		Object.assign(page, navigation_result.props.page);
	}

	is_navigating = false;

	if (type === 'popstate') {
		restore_snapshot(current_navigation_index);
	}

	nav.fulfil(undefined);

	after_navigate_callbacks.forEach((fn) =>
		fn(/** @type {import('@sveltejs/kit').AfterNavigate} */ (nav.navigation))
	);

	stores.navigating.set((navigating.current = null));

	updating = false;
}

/**
 * Does a full page reload if it wouldn't result in an endless loop in the SPA case
 * @param {URL} url
 * @param {{ id: string | null }} route
 * @param {App.Error} error
 * @param {number} status
 * @returns {Promise<import('./types.js').NavigationFinished>}
 */
async function server_fallback(url, route, error, status) {
	if (url.origin === origin && url.pathname === location.pathname && !hydrated) {
		// We would reload the same page we're currently on, which isn't hydrated,
		// which means no SSR, which means we would end up in an endless loop
		return await load_root_error_page({
			status,
			error,
			url,
			route
		});
	}

	if (DEV && status !== 404) {
		console.error(
			'An error occurred while loading the page. This will cause a full page reload. (This message will only appear during development.)'
		);

		debugger; // eslint-disable-line
	}

	return await native_navigation(url);
}

if (import.meta.hot) {
	import.meta.hot.on('vite:beforeUpdate', () => {
		if (current.error) location.reload();
	});
}

/** @typedef {(typeof PRELOAD_PRIORITIES)['hover'] | (typeof PRELOAD_PRIORITIES)['tap']} PreloadDataPriority */

function setup_preload() {
	/** @type {NodeJS.Timeout} */
	let mousemove_timeout;
	/** @type {Element} */
	let current_a;
	/** @type {PreloadDataPriority} */
	let current_priority;

	container.addEventListener('mousemove', (event) => {
		const target = /** @type {Element} */ (event.target);

		clearTimeout(mousemove_timeout);
		mousemove_timeout = setTimeout(() => {
			void preload(target, PRELOAD_PRIORITIES.hover);
		}, 20);
	});

	/** @param {Event} event */
	function tap(event) {
		if (event.defaultPrevented) return;
		void preload(/** @type {Element} */ (event.composedPath()[0]), PRELOAD_PRIORITIES.tap);
	}

	container.addEventListener('mousedown', tap);
	container.addEventListener('touchstart', tap, { passive: true });

	const observer = new IntersectionObserver(
		(entries) => {
			for (const entry of entries) {
				if (entry.isIntersecting) {
					void _preload_code(new URL(/** @type {HTMLAnchorElement} */ (entry.target).href));
					observer.unobserve(entry.target);
				}
			}
		},
		{ threshold: 0 }
	);

	/**
	 * @param {Element} element
	 * @param {PreloadDataPriority} priority
	 */
	async function preload(element, priority) {
		const a = find_anchor(element, container);

		// we don't want to preload data again if the user has already hovered/tapped
		const interacted = a === current_a && priority >= current_priority;
		if (!a || interacted) return;

		const { url, external, download } = get_link_info(a, base, app.hash);
		if (external || download) return;

		const options = get_router_options(a);

		// we don't want to preload data for a page we're already on
		const same_url = url && get_page_key(current.url) === get_page_key(url);
		if (options.reload || same_url) return;

		if (priority <= options.preload_data) {
			current_a = a;
			// we don't want to preload data again on tap if we've already preloaded it on hover
			current_priority = PRELOAD_PRIORITIES.tap;

			const intent = await get_navigation_intent(url, false);
			if (!intent) return;

			if (DEV) {
				void _preload_data(intent).then((result) => {
					if (result.type === 'loaded' && result.state.error) {
						console.warn(
							`Preloading data for ${intent.url.pathname} failed with the following error: ${result.state.error.message}\n` +
								'If this error is transient, you can ignore it. Otherwise, consider disabling preloading for this route. ' +
								'This route was preloaded due to a data-sveltekit-preload-data attribute. ' +
								'See https://svelte.dev/docs/kit/link-options for more info'
						);
					}
				});
			} else {
				void _preload_data(intent);
			}
		} else if (priority <= options.preload_code) {
			current_a = a;
			current_priority = priority;
			void _preload_code(/** @type {URL} */ (url));
		}
	}

	function after_navigate() {
		observer.disconnect();

		for (const a of container.querySelectorAll('a')) {
			const { url, external, download } = get_link_info(a, base, app.hash);
			if (external || download) continue;

			const options = get_router_options(a);
			if (options.reload) continue;

			if (options.preload_code === PRELOAD_PRIORITIES.viewport) {
				observer.observe(a);
			}

			if (options.preload_code === PRELOAD_PRIORITIES.eager) {
				void _preload_code(/** @type {URL} */ (url));
			}
		}
	}

	after_navigate_callbacks.add(after_navigate);
	after_navigate();
}

/**
 * @param {unknown} error
 * @param {import('@sveltejs/kit').NavigationEvent} event
 * @returns {import('types').MaybePromise<App.Error>}
 */
function handle_error(error, event) {
	if (error instanceof HttpError) {
		return error.body;
	}

	if (DEV) {
		errored = true;
		console.warn('The next HMR update will cause the page to reload');
	}

	const status = get_status(error);
	const message = get_message(error);

	return (
		app.hooks.handleError({ error, event, status, message }) ?? /** @type {any} */ ({ message })
	);
}

/**
 * @template {Function} T
 * @param {Set<T>} callbacks
 * @param {T} callback
 */
function add_navigation_callback(callbacks, callback) {
	onMount(() => {
		callbacks.add(callback);

		return () => {
			callbacks.delete(callback);
		};
	});
}

/**
 * A lifecycle function that runs the supplied `callback` when the current component mounts, and also whenever we navigate to a URL.
 *
 * `afterNavigate` must be called during a component initialization. It remains active as long as the component is mounted.
 * @param {(navigation: import('@sveltejs/kit').AfterNavigate) => void} callback
 * @returns {void}
 */
export function afterNavigate(callback) {
	add_navigation_callback(after_navigate_callbacks, callback);
}

/**
 * A navigation interceptor that triggers before we navigate to a URL, whether by clicking a link, calling `goto(...)`, or using the browser back/forward controls.
 *
 * Calling `cancel()` will prevent the navigation from completing. If `navigation.type === 'leave'` — meaning the user is navigating away from the app (or closing the tab) — calling `cancel` will trigger the native browser unload confirmation dialog. In this case, the navigation may or may not be cancelled depending on the user's response.
 *
 * When a navigation isn't to a SvelteKit-owned route (and therefore controlled by SvelteKit's client-side router), `navigation.to.route.id` will be `null`.
 *
 * If the navigation will (if not cancelled) cause the document to unload — in other words `'leave'` navigations and `'link'` navigations where `navigation.to.route === null` — `navigation.willUnload` is `true`.
 *
 * `beforeNavigate` must be called during a component initialization. It remains active as long as the component is mounted.
 * @param {(navigation: import('@sveltejs/kit').BeforeNavigate) => void} callback
 * @returns {void}
 */
export function beforeNavigate(callback) {
	add_navigation_callback(before_navigate_callbacks, callback);
}

/**
 * A lifecycle function that runs the supplied `callback` immediately before we navigate to a new URL except during full-page navigations.
 *
 * If you return a `Promise`, SvelteKit will wait for it to resolve before completing the navigation. This allows you to — for example — use `document.startViewTransition`. Avoid promises that are slow to resolve, since navigation will appear stalled to the user.
 *
 * If a function (or a `Promise` that resolves to a function) is returned from the callback, it will be called once the DOM has updated.
 *
 * `onNavigate` must be called during a component initialization. It remains active as long as the component is mounted.
 * @param {(navigation: import('@sveltejs/kit').OnNavigate) => import('types').MaybePromise<(() => void) | void>} callback
 * @returns {void}
 */
export function onNavigate(callback) {
	add_navigation_callback(on_navigate_callbacks, callback);
}

/**
 * If called when the page is being updated following a navigation (in `onMount` or `afterNavigate` or an action, for example), this disables SvelteKit's built-in scroll handling.
 * This is generally discouraged, since it breaks user expectations.
 * @returns {void}
 */
export function disableScrollHandling() {
	if (!BROWSER) {
		throw new Error('Cannot call disableScrollHandling() on the server');
	}

	if (DEV && started && !updating) {
		throw new Error('Can only disable scroll handling during navigation');
	}

	if (updating || !started) {
		autoscroll = false;
	}
}

/**
 * Allows you to navigate programmatically to a given route, with options such as keeping the current element focused.
 * Returns a Promise that resolves when SvelteKit navigates (or fails to navigate, in which case the promise rejects) to the specified `url`.
 *
 * For external URLs, use `window.location = url` instead of calling `goto(url)`.
 *
 * @param {string | URL} url Where to navigate to. Note that if you've set [`config.kit.paths.base`](https://svelte.dev/docs/kit/configuration#paths) and the URL is root-relative, you need to prepend the base path if you want to navigate within the app.
 * @param {Object} [opts] Options related to the navigation
 * @param {boolean} [opts.replaceState] If `true`, will replace the current `history` entry rather than creating a new one with `pushState`
 * @param {boolean} [opts.noScroll] If `true`, the browser will maintain its scroll position rather than scrolling to the top of the page after navigation
 * @param {boolean} [opts.keepFocus] If `true`, the currently focused element will retain focus after navigation. Otherwise, focus will be reset to the body
 * @param {boolean} [opts.invalidateAll] If `true`, all `load` functions of the page will be rerun. See https://svelte.dev/docs/kit/load#rerunning-load-functions for more info on invalidation.
 * @param {Array<string | URL | ((url: URL) => boolean)>} [opts.invalidate] Causes any load functions to re-run if they depend on one of the urls
 * @param {App.PageState} [opts.state] An optional object that will be available as `page.state`
 * @returns {Promise<void>}
 */
export function goto(url, opts = {}) {
	if (!BROWSER) {
		throw new Error('Cannot call goto(...) on the server');
	}

	url = new URL(resolve_url(url));

	if (url.origin !== origin) {
		return Promise.reject(
			new Error(
				DEV
					? `Cannot use \`goto\` with an external URL. Use \`window.location = "${url}"\` instead`
					: 'goto: invalid URL'
			)
		);
	}

	return _goto(url, opts, 0);
}

/**
 * Causes any `load` functions belonging to the currently active page to re-run if they depend on the `url` in question, via `fetch` or `depends`. Returns a `Promise` that resolves when the page is subsequently updated.
 *
 * If the argument is given as a `string` or `URL`, it must resolve to the same URL that was passed to `fetch` or `depends` (including query parameters).
 * To create a custom identifier, use a string beginning with `[a-z]+:` (e.g. `custom:state`) — this is a valid URL.
 *
 * The `function` argument can be used define a custom predicate. It receives the full `URL` and causes `load` to rerun if `true` is returned.
 * This can be useful if you want to invalidate based on a pattern instead of a exact match.
 *
 * ```ts
 * // Example: Match '/path' regardless of the query parameters
 * import { invalidate } from '$app/navigation';
 *
 * invalidate((url) => url.pathname === '/path');
 * ```
 * @param {string | URL | ((url: URL) => boolean)} resource The invalidated URL
 * @returns {Promise<void>}
 */
export function invalidate(resource) {
	if (!BROWSER) {
		throw new Error('Cannot call invalidate(...) on the server');
	}

	push_invalidated(resource);

	return _invalidate();
}

/**
 * @param {string | URL | ((url: URL) => boolean)} resource The invalidated URL
 */
function push_invalidated(resource) {
	if (typeof resource === 'function') {
		invalidated.push(resource);
	} else {
		const { href } = new URL(resource, location.href);
		invalidated.push((url) => url.href === href);
	}
}

/**
 * Causes all `load` functions belonging to the currently active page to re-run. Returns a `Promise` that resolves when the page is subsequently updated.
 * @returns {Promise<void>}
 */
export function invalidateAll() {
	if (!BROWSER) {
		throw new Error('Cannot call invalidateAll() on the server');
	}

	force_invalidation = true;
	return _invalidate();
}

/**
 * Programmatically preloads the given page, which means
 *  1. ensuring that the code for the page is loaded, and
 *  2. calling the page's load function with the appropriate options.
 *
 * This is the same behaviour that SvelteKit triggers when the user taps or mouses over an `<a>` element with `data-sveltekit-preload-data`.
 * If the next navigation is to `href`, the values returned from load will be used, making navigation instantaneous.
 * Returns a Promise that resolves with the result of running the new route's `load` functions once the preload is complete.
 *
 * @param {string} href Page to preload
 * @returns {Promise<{ type: 'loaded'; status: number; data: Record<string, any> } | { type: 'redirect'; location: string }>}
 */
export async function preloadData(href) {
	if (!BROWSER) {
		throw new Error('Cannot call preloadData(...) on the server');
	}

	const url = resolve_url(href);
	const intent = await get_navigation_intent(url, false);

	if (!intent) {
		throw new Error(`Attempted to preload a URL that does not belong to this app: ${url}`);
	}

	const result = await _preload_data(intent);
	if (result.type === 'redirect') {
		return {
			type: result.type,
			location: result.location
		};
	}

	const { status, data } = result.props.page ?? page;
	return { type: result.type, status, data };
}

/**
 * Programmatically imports the code for routes that haven't yet been fetched.
 * Typically, you might call this to speed up subsequent navigation.
 *
 * You can specify routes by any matching pathname such as `/about` (to match `src/routes/about/+page.svelte`) or `/blog/*` (to match `src/routes/blog/[slug]/+page.svelte`).
 *
 * Unlike `preloadData`, this won't call `load` functions.
 * Returns a Promise that resolves when the modules have been imported.
 *
 * @param {string} pathname
 * @returns {Promise<void>}
 */
export async function preloadCode(pathname) {
	if (!BROWSER) {
		throw new Error('Cannot call preloadCode(...) on the server');
	}

	const url = new URL(pathname, current.url);

	if (DEV) {
		if (!pathname.startsWith('/')) {
			throw new Error(
				'argument passed to preloadCode must be a pathname (i.e. "/about" rather than "http://example.com/about"'
			);
		}

		if (!pathname.startsWith(base)) {
			throw new Error(
				`pathname passed to preloadCode must start with \`paths.base\` (i.e. "${base}${pathname}" rather than "${pathname}")`
			);
		}

		if (__SVELTEKIT_CLIENT_ROUTING__) {
			const rerouted = await get_rerouted_url(url);
			if (!rerouted || !routes.find((route) => route.exec(get_url_path(rerouted)))) {
				throw new Error(`'${pathname}' did not match any routes`);
			}
		}
	}

	return _preload_code(url);
}

/**
 * Programmatically create a new history entry with the given `page.state`. To use the current URL, you can pass `''` as the first argument. Used for [shallow routing](https://svelte.dev/docs/kit/shallow-routing).
 *
 * @param {string | URL} url
 * @param {App.PageState} state
 * @returns {void}
 */
export function pushState(url, state) {
	if (!BROWSER) {
		throw new Error('Cannot call pushState(...) on the server');
	}

	if (DEV) {
		if (!started) {
			throw new Error('Cannot call pushState(...) before router is initialized');
		}

		try {
			// use `devalue.stringify` as a convenient way to ensure we exclude values that can't be properly rehydrated, such as custom class instances
			devalue.stringify(state);
		} catch (error) {
			// @ts-expect-error
			throw new Error(`Could not serialize state${error.path}`);
		}
	}

	update_scroll_positions(current_history_index);

	const opts = {
		[HISTORY_INDEX]: (current_history_index += 1),
		[NAVIGATION_INDEX]: current_navigation_index,
		[PAGE_URL_KEY]: page.url.href,
		[STATES_KEY]: state
	};

	history.pushState(opts, '', resolve_url(url));
	has_navigated = true;

	page.state = state;
	root.$set({
		// we need to assign a new page object so that subscribers are correctly notified
		page: clone_page(page)
	});

	clear_onward_history(current_history_index, current_navigation_index);
}

/**
 * Programmatically replace the current history entry with the given `page.state`. To use the current URL, you can pass `''` as the first argument. Used for [shallow routing](https://svelte.dev/docs/kit/shallow-routing).
 *
 * @param {string | URL} url
 * @param {App.PageState} state
 * @returns {void}
 */
export function replaceState(url, state) {
	if (!BROWSER) {
		throw new Error('Cannot call replaceState(...) on the server');
	}

	if (DEV) {
		if (!started) {
			throw new Error('Cannot call replaceState(...) before router is initialized');
		}

		try {
			// use `devalue.stringify` as a convenient way to ensure we exclude values that can't be properly rehydrated, such as custom class instances
			devalue.stringify(state);
		} catch (error) {
			// @ts-expect-error
			throw new Error(`Could not serialize state${error.path}`);
		}
	}

	const opts = {
		[HISTORY_INDEX]: current_history_index,
		[NAVIGATION_INDEX]: current_navigation_index,
		[PAGE_URL_KEY]: page.url.href,
		[STATES_KEY]: state
	};

	history.replaceState(opts, '', resolve_url(url));

	page.state = state;
	root.$set({
		page: clone_page(page)
	});
}

/**
 * This action updates the `form` property of the current page with the given data and updates `page.status`.
 * In case of an error, it redirects to the nearest error page.
 * @template {Record<string, unknown> | undefined} Success
 * @template {Record<string, unknown> | undefined} Failure
 * @param {import('@sveltejs/kit').ActionResult<Success, Failure>} result
 * @returns {Promise<void>}
 */
export async function applyAction(result) {
	if (!BROWSER) {
		throw new Error('Cannot call applyAction(...) on the server');
	}

	if (result.type === 'error') {
		const url = new URL(location.href);

		const { branch, route } = current;
		if (!route) return;

		const error_load = await load_nearest_error_page(current.branch.length, branch, route.errors);
		if (error_load) {
			const navigation_result = get_navigation_result_from_branch({
				url,
				params: current.params,
				branch: branch.slice(0, error_load.idx).concat(error_load.node),
				status: result.status ?? 500,
				error: result.error,
				route
			});

			current = navigation_result.state;

			root.$set(navigation_result.props);
			update(navigation_result.props.page);

			void tick().then(reset_focus);
		}
	} else if (result.type === 'redirect') {
		await _goto(result.location, { invalidateAll: true }, 0);
	} else {
		page.form = result.data;
		page.status = result.status;

		/** @type {Record<string, any>} */
		root.$set({
			// this brings Svelte's view of the world in line with SvelteKit's
			// after use:enhance reset the form....
			form: null,
			page: clone_page(page)
		});

		// ...so that setting the `form` prop takes effect and isn't ignored
		await tick();
		root.$set({ form: result.data });

		if (result.type === 'success') {
			reset_focus();
		}
	}
}

function _start_router() {
	history.scrollRestoration = 'manual';

	// Adopted from Nuxt.js
	// Reset scrollRestoration to auto when leaving page, allowing page reload
	// and back-navigation from other pages to use the browser to restore the
	// scrolling position.
	addEventListener('beforeunload', (e) => {
		let should_block = false;

		persist_state();

		if (!is_navigating) {
			const nav = create_navigation(current, undefined, null, 'leave');

			// If we're navigating, beforeNavigate was already called. If we end up in here during navigation,
			// it's due to an external or full-page-reload link, for which we don't want to call the hook again.
			/** @type {import('@sveltejs/kit').BeforeNavigate} */
			const navigation = {
				...nav.navigation,
				cancel: () => {
					should_block = true;
					nav.reject(new Error('navigation cancelled'));
				}
			};

			before_navigate_callbacks.forEach((fn) => fn(navigation));
		}

		if (should_block) {
			e.preventDefault();
			e.returnValue = '';
		} else {
			history.scrollRestoration = 'auto';
		}
	});

	addEventListener('visibilitychange', () => {
		if (document.visibilityState === 'hidden') {
			persist_state();
		}
	});

	// @ts-expect-error this isn't supported everywhere yet
	if (!navigator.connection?.saveData) {
		setup_preload();
	}

	/** @param {MouseEvent} event */
	container.addEventListener('click', async (event) => {
		// Adapted from https://github.com/visionmedia/page.js
		// MIT license https://github.com/visionmedia/page.js#license
		if (event.button || event.which !== 1) return;
		if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
		if (event.defaultPrevented) return;

		const a = find_anchor(/** @type {Element} */ (event.composedPath()[0]), container);
		if (!a) return;

		const { url, external, target, download } = get_link_info(a, base, app.hash);
		if (!url) return;

		// bail out before `beforeNavigate` if link opens in a different tab
		if (target === '_parent' || target === '_top') {
			if (window.parent !== window) return;
		} else if (target && target !== '_self') {
			return;
		}

		const options = get_router_options(a);
		const is_svg_a_element = a instanceof SVGAElement;

		// Ignore URL protocols that differ to the current one and are not http(s) (e.g. `mailto:`, `tel:`, `myapp:`, etc.)
		// This may be wrong when the protocol is x: and the link goes to y:.. which should be treated as an external
		// navigation, but it's not clear how to handle that case and it's not likely to come up in practice.
		// MEMO: Without this condition, firefox will open mailer twice.
		// See:
		// - https://github.com/sveltejs/kit/issues/4045
		// - https://github.com/sveltejs/kit/issues/5725
		// - https://github.com/sveltejs/kit/issues/6496
		if (
			!is_svg_a_element &&
			url.protocol !== location.protocol &&
			!(url.protocol === 'https:' || url.protocol === 'http:')
		)
			return;

		if (download) return;

		const [nonhash, hash] = (app.hash ? url.hash.replace(/^#/, '') : url.href).split('#');
		const same_pathname = nonhash === strip_hash(location);

		// Ignore the following but fire beforeNavigate
		if (external || (options.reload && (!same_pathname || !hash))) {
			if (_before_navigate({ url, type: 'link' })) {
				// set `navigating` to `true` to prevent `beforeNavigate` callbacks
				// being called when the page unloads
				is_navigating = true;
			} else {
				event.preventDefault();
			}

			return;
		}

		// Check if new url only differs by hash and use the browser default behavior in that case
		// This will ensure the `hashchange` event is fired
		// Removing the hash does a full page navigation in the browser, so make sure a hash is present
		if (hash !== undefined && same_pathname) {
			// If we are trying to navigate to the same hash, we should only
			// attempt to scroll to that element and avoid any history changes.
			// Otherwise, this can cause Firefox to incorrectly assign a null
			// history state value without any signal that we can detect.
			const [, current_hash] = current.url.href.split('#');
			if (current_hash === hash) {
				event.preventDefault();

				// We're already on /# and click on a link that goes to /#, or we're on
				// /#top and click on a link that goes to /#top. In those cases just go to
				// the top of the page, and avoid a history change.
				if (hash === '' || (hash === 'top' && a.ownerDocument.getElementById('top') === null)) {
					window.scrollTo({ top: 0 });
				} else {
					const element = a.ownerDocument.getElementById(decodeURIComponent(hash));
					if (element) {
						element.scrollIntoView();
						element.focus();
					}
				}

				return;
			}
			// set this flag to distinguish between navigations triggered by
			// clicking a hash link and those triggered by popstate
			hash_navigating = true;

			update_scroll_positions(current_history_index);

			update_url(url);

			if (!options.replace_state) return;

			// hashchange event shouldn't occur if the router is replacing state.
			hash_navigating = false;
		}

		event.preventDefault();

		// allow the browser to repaint before navigating —
		// this prevents INP scores being penalised
		await new Promise((fulfil) => {
			requestAnimationFrame(() => {
				setTimeout(fulfil, 0);
			});

			setTimeout(fulfil, 100); // fallback for edge case where rAF doesn't fire because e.g. tab was backgrounded
		});

		await navigate({
			type: 'link',
			url,
			keepfocus: options.keepfocus,
			noscroll: options.noscroll,
			replace_state: options.replace_state ?? url.href === location.href
		});
	});

	container.addEventListener('submit', (event) => {
		if (event.defaultPrevented) return;

		const form = /** @type {HTMLFormElement} */ (
			HTMLFormElement.prototype.cloneNode.call(event.target)
		);

		const submitter = /** @type {HTMLButtonElement | HTMLInputElement | null} */ (event.submitter);

		const target = submitter?.formTarget || form.target;

		if (target === '_blank') return;

		const method = submitter?.formMethod || form.method;

		if (method !== 'get') return;

		// It is impossible to use form actions with hash router, so we just ignore handling them here
		const url = new URL(
			(submitter?.hasAttribute('formaction') && submitter?.formAction) || form.action
		);

		if (is_external_url(url, base, false)) return;

		const event_form = /** @type {HTMLFormElement} */ (event.target);

		const options = get_router_options(event_form);
		if (options.reload) return;

		event.preventDefault();
		event.stopPropagation();

		const data = new FormData(event_form);

		const submitter_name = submitter?.getAttribute('name');
		if (submitter_name) {
			data.append(submitter_name, submitter?.getAttribute('value') ?? '');
		}

		// @ts-expect-error `URLSearchParams(fd)` is kosher, but typescript doesn't know that
		url.search = new URLSearchParams(data).toString();

		void navigate({
			type: 'form',
			url,
			keepfocus: options.keepfocus,
			noscroll: options.noscroll,
			replace_state: options.replace_state ?? url.href === location.href
		});
	});

	addEventListener('popstate', async (event) => {
		if (event.state?.[HISTORY_INDEX]) {
			const history_index = event.state[HISTORY_INDEX];
			token = {};

			// if a popstate-driven navigation is cancelled, we need to counteract it
			// with history.go, which means we end up back here, hence this check
			if (history_index === current_history_index) return;

			const scroll = scroll_positions[history_index];
			const state = event.state[STATES_KEY] ?? {};
			const url = new URL(event.state[PAGE_URL_KEY] ?? location.href);
			const navigation_index = event.state[NAVIGATION_INDEX];
			const is_hash_change = current.url ? strip_hash(location) === strip_hash(current.url) : false;
			const shallow =
				navigation_index === current_navigation_index && (has_navigated || is_hash_change);

			if (shallow) {
				// We don't need to navigate, we just need to update scroll and/or state.
				// This happens with hash links and `pushState`/`replaceState`. The
				// exception is if we haven't navigated yet, since we could have
				// got here after a modal navigation then a reload
				if (state !== page.state) {
					page.state = state;
				}

				update_url(url);

				scroll_positions[current_history_index] = scroll_state();
				if (scroll) scrollTo(scroll.x, scroll.y);

				current_history_index = history_index;
				return;
			}

			const delta = history_index - current_history_index;

			await navigate({
				type: 'popstate',
				url,
				popped: {
					state,
					scroll,
					delta
				},
				accept: () => {
					current_history_index = history_index;
					current_navigation_index = navigation_index;
				},
				block: () => {
					history.go(-delta);
				},
				nav_token: token
			});
		} else {
			// since popstate event is also emitted when an anchor referencing the same
			// document is clicked, we have to check that the router isn't already handling
			// the navigation. otherwise we would be updating the page store twice.
			if (!hash_navigating) {
				const url = new URL(location.href);
				update_url(url);

				// if the user edits the hash via the browser URL bar, trigger a full-page
				// reload to align with pathname router behavior
				if (app.hash) {
					location.reload();
				}
			}
		}
	});

	addEventListener('hashchange', () => {
		// if the hashchange happened as a result of clicking on a link,
		// we need to update history, otherwise we have to leave it alone
		if (hash_navigating) {
			hash_navigating = false;
			history.replaceState(
				{
					...history.state,
					[HISTORY_INDEX]: ++current_history_index,
					[NAVIGATION_INDEX]: current_navigation_index
				},
				'',
				location.href
			);
		}
	});

	// fix link[rel=icon], because browsers will occasionally try to load relative
	// URLs after a pushState/replaceState, resulting in a 404 — see
	// https://github.com/sveltejs/kit/issues/3748#issuecomment-1125980897
	for (const link of document.querySelectorAll('link')) {
		if (ICON_REL_ATTRIBUTES.has(link.rel)) {
			link.href = link.href; // eslint-disable-line
		}
	}

	addEventListener('pageshow', (event) => {
		// If the user navigates to another site and then uses the back button and
		// bfcache hits, we need to set navigating to null, the site doesn't know
		// the navigation away from it was successful.
		// Info about bfcache here: https://web.dev/bfcache
		if (event.persisted) {
			stores.navigating.set((navigating.current = null));
		}
	});

	/**
	 * @param {URL} url
	 */
	function update_url(url) {
		current.url = page.url = url;
		stores.page.set(clone_page(page));
		stores.page.notify();
	}
}

/**
 * @param {HTMLElement} target
 * @param {import('./types.js').HydrateOptions} opts
 */
async function _hydrate(
	target,
	{ status = 200, error, node_ids, params, route, server_route, data: server_data_nodes, form }
) {
	hydrated = true;

	const url = new URL(location.href);

	/** @type {import('types').CSRRoute | undefined} */
	let parsed_route;

	if (__SVELTEKIT_CLIENT_ROUTING__) {
		if (!__SVELTEKIT_EMBEDDED__) {
			// See https://github.com/sveltejs/kit/pull/4935#issuecomment-1328093358 for one motivation
			// of determining the params on the client side.
			({ params = {}, route = { id: null } } = (await get_navigation_intent(url, false)) || {});
		}

		parsed_route = routes.find(({ id }) => id === route.id);
	} else {
		// undefined in case of 404
		if (server_route) {
			parsed_route = route = parse_server_route(server_route, app.nodes);
		} else {
			route = { id: null };
			params = {};
		}
	}

	/** @type {import('./types.js').NavigationFinished | undefined} */
	let result;
	let hydrate = true;

	try {
		const branch_promises = node_ids.map(async (n, i) => {
			const server_data_node = server_data_nodes[i];
			// Type isn't completely accurate, we still need to deserialize uses
			if (server_data_node?.uses) {
				server_data_node.uses = deserialize_uses(server_data_node.uses);
			}

			return load_node({
				loader: app.nodes[n],
				url,
				params,
				route,
				parent: async () => {
					const data = {};
					for (let j = 0; j < i; j += 1) {
						Object.assign(data, (await branch_promises[j]).data);
					}
					return data;
				},
				server_data_node: create_data_node(server_data_node)
			});
		});

		/** @type {Array<import('./types.js').BranchNode | undefined>} */
		const branch = await Promise.all(branch_promises);

		// server-side will have compacted the branch, reinstate empty slots
		// so that error boundaries can be lined up correctly
		if (parsed_route) {
			const layouts = parsed_route.layouts;
			for (let i = 0; i < layouts.length; i++) {
				if (!layouts[i]) {
					branch.splice(i, 0, undefined);
				}
			}
		}

		result = get_navigation_result_from_branch({
			url,
			params,
			branch,
			status,
			error,
			form,
			route: parsed_route ?? null
		});
	} catch (error) {
		if (error instanceof Redirect) {
			// this is a real edge case — `load` would need to return
			// a redirect but only in the browser
			await native_navigation(new URL(error.location, location.href));
			return;
		}

		result = await load_root_error_page({
			status: get_status(error),
			error: await handle_error(error, { url, params, route }),
			url,
			route
		});

		target.textContent = '';
		hydrate = false;
	}

	if (result.props.page) {
		result.props.page.state = {};
	}

	initialize(result, target, hydrate);
}

/**
 * @param {URL} url
 * @param {boolean[]} invalid
 * @returns {Promise<import('types').ServerNodesResponse | import('types').ServerRedirectNode>}
 */
async function load_data(url, invalid) {
	const data_url = new URL(url);
	data_url.pathname = add_data_suffix(url.pathname);
	if (url.pathname.endsWith('/')) {
		data_url.searchParams.append(TRAILING_SLASH_PARAM, '1');
	}
	if (DEV && url.searchParams.has(INVALIDATED_PARAM)) {
		throw new Error(`Cannot used reserved query parameter "${INVALIDATED_PARAM}"`);
	}
	data_url.searchParams.append(INVALIDATED_PARAM, invalid.map((i) => (i ? '1' : '0')).join(''));

	// use window.fetch directly to allow using a 3rd party-patched fetch implementation
	const fetcher = DEV ? dev_fetch : window.fetch;
	const res = await fetcher(data_url.href, {});

	if (!res.ok) {
		// error message is a JSON-stringified string which devalue can't handle at the top level
		// turn it into a HttpError to not call handleError on the client again (was already handled on the server)
		// if `__data.json` doesn't exist or the server has an internal error,
		// avoid parsing the HTML error page as a JSON
		/** @type {string | undefined} */
		let message;
		if (res.headers.get('content-type')?.includes('application/json')) {
			message = await res.json();
		} else if (res.status === 404) {
			message = 'Not Found';
		} else if (res.status === 500) {
			message = 'Internal Error';
		}
		throw new HttpError(res.status, message);
	}

	// TODO: fix eslint error / figure out if it actually applies to our situation
	// eslint-disable-next-line
	return new Promise(async (resolve) => {
		/**
		 * Map of deferred promises that will be resolved by a subsequent chunk of data
		 * @type {Map<string, import('types').Deferred>}
		 */
		const deferreds = new Map();
		const reader = /** @type {ReadableStream<Uint8Array>} */ (res.body).getReader();
		const decoder = new TextDecoder();

		/**
		 * @param {any} data
		 */
		function deserialize(data) {
			return devalue.unflatten(data, {
				...app.decoders,
				Promise: (id) => {
					return new Promise((fulfil, reject) => {
						deferreds.set(id, { fulfil, reject });
					});
				}
			});
		}

		let text = '';

		while (true) {
			// Format follows ndjson (each line is a JSON object) or regular JSON spec
			const { done, value } = await reader.read();
			if (done && !text) break;

			text += !value && text ? '\n' : decoder.decode(value, { stream: true }); // no value -> final chunk -> add a new line to trigger the last parse

			while (true) {
				const split = text.indexOf('\n');
				if (split === -1) {
					break;
				}

				const node = JSON.parse(text.slice(0, split));
				text = text.slice(split + 1);

				if (node.type === 'redirect') {
					return resolve(node);
				}

				if (node.type === 'data') {
					// This is the first (and possibly only, if no pending promises) chunk
					node.nodes?.forEach((/** @type {any} */ node) => {
						if (node?.type === 'data') {
							node.uses = deserialize_uses(node.uses);
							node.data = deserialize(node.data);
						}
					});

					resolve(node);
				} else if (node.type === 'chunk') {
					// This is a subsequent chunk containing deferred data
					const { id, data, error } = node;
					const deferred = /** @type {import('types').Deferred} */ (deferreds.get(id));
					deferreds.delete(id);

					if (error) {
						deferred.reject(deserialize(error));
					} else {
						deferred.fulfil(deserialize(data));
					}
				}
			}
		}
	});

	// TODO edge case handling necessary? stream() read fails?
}

/**
 * @param {any} uses
 * @return {import('types').Uses}
 */
function deserialize_uses(uses) {
	return {
		dependencies: new Set(uses?.dependencies ?? []),
		params: new Set(uses?.params ?? []),
		parent: !!uses?.parent,
		route: !!uses?.route,
		url: !!uses?.url,
		search_params: new Set(uses?.search_params ?? [])
	};
}

function reset_focus() {
	const autofocus = document.querySelector('[autofocus]');
	if (autofocus) {
		// @ts-ignore
		autofocus.focus();
	} else {
		// Reset page selection and focus
		// We try to mimic browsers' behaviour as closely as possible by targeting the
		// first scrollable region, but unfortunately it's not a perfect match — e.g.
		// shift-tabbing won't immediately cycle up from the end of the page on Chromium
		// See https://html.spec.whatwg.org/multipage/interaction.html#get-the-focusable-area
		const root = document.body;
		const tabindex = root.getAttribute('tabindex');

		root.tabIndex = -1;
		// @ts-expect-error
		root.focus({ preventScroll: true, focusVisible: false });

		// restore `tabindex` as to prevent `root` from stealing input from elements
		if (tabindex !== null) {
			root.setAttribute('tabindex', tabindex);
		} else {
			root.removeAttribute('tabindex');
		}

		// capture current selection, so we can compare the state after
		// snapshot restoration and afterNavigate callbacks have run
		const selection = getSelection();

		if (selection && selection.type !== 'None') {
			/** @type {Range[]} */
			const ranges = [];

			for (let i = 0; i < selection.rangeCount; i += 1) {
				ranges.push(selection.getRangeAt(i));
			}

			setTimeout(() => {
				if (selection.rangeCount !== ranges.length) return;

				for (let i = 0; i < selection.rangeCount; i += 1) {
					const a = ranges[i];
					const b = selection.getRangeAt(i);

					// we need to do a deep comparison rather than just `a !== b` because
					// Safari behaves differently to other browsers
					if (
						a.commonAncestorContainer !== b.commonAncestorContainer ||
						a.startContainer !== b.startContainer ||
						a.endContainer !== b.endContainer ||
						a.startOffset !== b.startOffset ||
						a.endOffset !== b.endOffset
					) {
						return;
					}
				}

				// if the selection hasn't changed (as a result of an element being (auto)focused,
				// or a programmatic selection, we reset everything as part of the navigation)
				// fixes https://github.com/sveltejs/kit/issues/8439
				selection.removeAllRanges();
			});
		}
	}
}

/**
 * @template {import('@sveltejs/kit').NavigationType} T
 * @param {import('./types.js').NavigationState} current
 * @param {import('./types.js').NavigationIntent | undefined} intent
 * @param {URL | null} url
 * @param {T} type
 */
function create_navigation(current, intent, url, type) {
	/** @type {(value: any) => void} */
	let fulfil;

	/** @type {(error: any) => void} */
	let reject;

	const complete = new Promise((f, r) => {
		fulfil = f;
		reject = r;
	});

	// Handle any errors off-chain so that it doesn't show up as an unhandled rejection
	complete.catch(() => {});

	/** @type {Omit<import('@sveltejs/kit').Navigation, 'type'> & { type: T }} */
	const navigation = {
		from: {
			params: current.params,
			route: { id: current.route?.id ?? null },
			url: current.url
		},
		to: url && {
			params: intent?.params ?? null,
			route: { id: intent?.route?.id ?? null },
			url
		},
		willUnload: !intent,
		type,
		complete
	};

	return {
		navigation,
		// @ts-expect-error
		fulfil,
		// @ts-expect-error
		reject
	};
}

/**
 * TODO: remove this in 3.0 when the page store is also removed
 *
 * We need to assign a new page object so that subscribers are correctly notified.
 * However, spreading `{ ...page }` returns an empty object so we manually
 * assign to each property instead.
 *
 * @param {import('@sveltejs/kit').Page} page
 */
function clone_page(page) {
	return {
		data: page.data,
		error: page.error,
		form: page.form,
		params: page.params,
		route: page.route,
		state: page.state,
		status: page.status,
		url: page.url
	};
}

/**
 * @param {URL} url
 * @returns {URL}
 */
function decode_hash(url) {
	const new_url = new URL(url);
	// Safari, for some reason, does change # to %23, when entered through the address bar
	new_url.hash = decodeURIComponent(url.hash);
	return new_url;
}

if (DEV) {
	// Nasty hack to silence harmless warnings the user can do nothing about
	const console_warn = console.warn;
	console.warn = function warn(...args) {
		if (
			args.length === 1 &&
			/<(Layout|Page|Error)(_[\w$]+)?> was created (with unknown|without expected) prop '(data|form)'/.test(
				args[0]
			)
		) {
			return;
		}
		console_warn(...args);
	};

	if (import.meta.hot) {
		import.meta.hot.on('vite:beforeUpdate', () => {
			if (errored) {
				location.reload();
			}
		});
	}
}
