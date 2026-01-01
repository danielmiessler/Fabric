import * as devalue from 'devalue';
import { readable, writable } from 'svelte/store';
import { DEV } from 'esm-env';
import { text } from '@sveltejs/kit';
import * as paths from '$app/paths/internal/server';
import { hash } from '../../../utils/hash.js';
import { serialize_data } from './serialize_data.js';
import { s } from '../../../utils/misc.js';
import { Csp } from './csp.js';
import { uneval_action_response } from './actions.js';
import { public_env } from '../../shared-server.js';
import { SVELTE_KIT_ASSETS } from '../../../constants.js';
import { SCHEME } from '../../../utils/url.js';
import { create_server_routing_response, generate_route_object } from './server_routing.js';
import { add_resolution_suffix } from '../../pathname.js';
import { try_get_request_store, with_request_store } from '@sveltejs/kit/internal/server';
import { text_encoder } from '../../utils.js';
import { get_global_name } from '../utils.js';
import { create_remote_key } from '../../shared.js';

// TODO rename this function/module

const updated = {
	...readable(false),
	check: () => false
};

/**
 * Creates the HTML response.
 * @param {{
 *   branch: Array<import('./types.js').Loaded>;
 *   fetched: Array<import('./types.js').Fetched>;
 *   options: import('types').SSROptions;
 *   manifest: import('@sveltejs/kit').SSRManifest;
 *   state: import('types').SSRState;
 *   page_config: { ssr: boolean; csr: boolean };
 *   status: number;
 *   error: App.Error | null;
 *   event: import('@sveltejs/kit').RequestEvent;
 *   event_state: import('types').RequestState;
 *   resolve_opts: import('types').RequiredResolveOptions;
 *   action_result?: import('@sveltejs/kit').ActionResult;
 *   data_serializer: import('./types.js').ServerDataSerializer
 * }} opts
 */
export async function render_response({
	branch,
	fetched,
	options,
	manifest,
	state,
	page_config,
	status,
	error = null,
	event,
	event_state,
	resolve_opts,
	action_result,
	data_serializer
}) {
	if (state.prerendering) {
		if (options.csp.mode === 'nonce') {
			throw new Error('Cannot use prerendering if config.kit.csp.mode === "nonce"');
		}

		if (options.app_template_contains_nonce) {
			throw new Error('Cannot use prerendering if page template contains %sveltekit.nonce%');
		}
	}

	const { client } = manifest._;

	const modulepreloads = new Set(client.imports);
	const stylesheets = new Set(client.stylesheets);
	const fonts = new Set(client.fonts);

	/**
	 * The value of the Link header that is added to the response when not prerendering
	 * @type {Set<string>}
	 */
	const link_headers = new Set();

	/**
	 * `<link>` tags that are added to prerendered responses
	 * (note that stylesheets are always added, prerendered or not)
	 * @type {Set<string>}
	 */
	const link_tags = new Set();

	/** @type {Map<string, string>} */
	// TODO if we add a client entry point one day, we will need to include inline_styles with the entry, otherwise stylesheets will be linked even if they are below inlineStyleThreshold
	const inline_styles = new Map();

	let rendered;

	const form_value =
		action_result?.type === 'success' || action_result?.type === 'failure'
			? (action_result.data ?? null)
			: null;

	/** @type {string} */
	let base = paths.base;

	/** @type {string} */
	let assets = paths.assets;

	/**
	 * An expression that will evaluate in the client to determine the resolved base path.
	 * We use a relative path when possible to support IPFS, the internet archive, etc.
	 */
	let base_expression = s(paths.base);

	// if appropriate, use relative paths for greater portability
	if (paths.relative) {
		if (!state.prerendering?.fallback) {
			const segments = event.url.pathname.slice(paths.base.length).split('/').slice(2);

			base = segments.map(() => '..').join('/') || '.';

			// resolve e.g. '../..' against current location, then remove trailing slash
			base_expression = `new URL(${s(base)}, location).pathname.slice(0, -1)`;

			if (!paths.assets || (paths.assets[0] === '/' && paths.assets !== SVELTE_KIT_ASSETS)) {
				assets = base;
			}
		} else if (options.hash_routing) {
			// we have to assume that we're in the right place
			base_expression = "new URL('.', location).pathname.slice(0, -1)";
		}
	}

	if (page_config.ssr) {
		/** @type {Record<string, any>} */
		const props = {
			stores: {
				page: writable(null),
				navigating: writable(null),
				updated
			},
			constructors: await Promise.all(
				branch.map(({ node }) => {
					if (!node.component) {
						// Can only be the leaf, layouts have a fallback component generated
						throw new Error(`Missing +page.svelte component for route ${event.route.id}`);
					}
					return node.component();
				})
			),
			form: form_value
		};

		let data = {};

		// props_n (instead of props[n]) makes it easy to avoid
		// unnecessary updates for layout components
		for (let i = 0; i < branch.length; i += 1) {
			data = { ...data, ...branch[i].data };
			props[`data_${i}`] = data;
		}

		props.page = {
			error,
			params: /** @type {Record<string, any>} */ (event.params),
			route: event.route,
			status,
			url: event.url,
			data,
			form: form_value,
			state: {}
		};

		const render_opts = {
			context: new Map([
				[
					'__request__',
					{
						page: props.page
					}
				]
			])
		};

		const fetch = globalThis.fetch;

		try {
			if (DEV) {
				let warned = false;
				globalThis.fetch = (info, init) => {
					if (typeof info === 'string' && !SCHEME.test(info)) {
						throw new Error(
							`Cannot call \`fetch\` eagerly during server-side rendering with relative URL (${info}) — put your \`fetch\` calls inside \`onMount\` or a \`load\` function instead`
						);
					} else if (!warned && !try_get_request_store()?.state.is_in_remote_function) {
						console.warn(
							'Avoid calling `fetch` eagerly during server-side rendering — put your `fetch` calls inside `onMount` or a `load` function instead'
						);
						warned = true;
					}

					return fetch(info, init);
				};
			}

			rendered = await with_request_store({ event, state: event_state }, async () => {
				// use relative paths during rendering, so that the resulting HTML is as
				// portable as possible, but reset afterwards
				if (paths.relative) paths.override({ base, assets });

				const maybe_promise = options.root.render(props, render_opts);
				// We have to invoke .then eagerly here in order to kick off rendering: it's only starting on access,
				// and `await maybe_promise` would eagerly access the .then property but call its function only after a tick, which is too late
				// for the paths.reset() below and for any eager getRequestEvent() calls during rendering without AsyncLocalStorage available.
				const rendered =
					options.async && 'then' in maybe_promise
						? /** @type {ReturnType<typeof options.root.render> & Promise<any>} */ (
								maybe_promise
							).then((r) => r)
						: maybe_promise;

				// TODO 3.0 remove options.async
				if (options.async) {
					// we reset this synchronously, rather than after async rendering is complete,
					// to avoid cross-talk between requests. This is a breaking change for
					// anyone who opts into async SSR, since `base` and `assets` will no
					// longer be relative to the current pathname.
					// TODO 3.0 remove `base` and `assets` in favour of `resolve(...)` and `asset(...)`
					paths.reset();
				}

				const { head, html, css } = options.async ? await rendered : rendered;

				return { head, html, css };
			});
		} finally {
			if (DEV) {
				globalThis.fetch = fetch;
			}

			paths.reset(); // just in case `options.root.render(...)` failed
		}

		for (const { node } of branch) {
			for (const url of node.imports) modulepreloads.add(url);
			for (const url of node.stylesheets) stylesheets.add(url);
			for (const url of node.fonts) fonts.add(url);

			if (node.inline_styles && !client.inline) {
				Object.entries(await node.inline_styles()).forEach(([k, v]) => inline_styles.set(k, v));
			}
		}
	} else {
		rendered = { head: '', html: '', css: { code: '', map: null } };
	}

	let head = '';
	let body = rendered.html;

	const csp = new Csp(options.csp, {
		prerender: !!state.prerendering
	});

	/** @param {string} path */
	const prefixed = (path) => {
		if (path.startsWith('/')) {
			// Vite makes the start script available through the base path and without it.
			// We load it via the base path in order to support remote IDE environments which proxy
			// all URLs under the base path during development.
			return paths.base + path;
		}
		return `${assets}/${path}`;
	};

	// inline styles can come from `bundleStrategy: 'inline'` or `inlineStyleThreshold`
	const style = client.inline
		? client.inline?.style
		: Array.from(inline_styles.values()).join('\n');

	if (style) {
		const attributes = DEV ? [' data-sveltekit'] : [];
		if (csp.style_needs_nonce) attributes.push(` nonce="${csp.nonce}"`);

		csp.add_style(style);

		head += `\n\t<style${attributes.join('')}>${style}</style>`;
	}

	for (const dep of stylesheets) {
		const path = prefixed(dep);

		const attributes = ['rel="stylesheet"'];

		if (inline_styles.has(dep)) {
			// don't load stylesheets that are already inlined
			// include them in disabled state so that Vite can detect them and doesn't try to add them
			attributes.push('disabled', 'media="(max-width: 0)"');
		} else {
			if (resolve_opts.preload({ type: 'css', path })) {
				link_headers.add(`<${encodeURI(path)}>; rel="preload"; as="style"; nopush`);
			}
		}

		head += `\n\t\t<link href="${path}" ${attributes.join(' ')}>`;
	}

	for (const dep of fonts) {
		const path = prefixed(dep);

		if (resolve_opts.preload({ type: 'font', path })) {
			const ext = dep.slice(dep.lastIndexOf('.') + 1);

			link_tags.add(`<link rel="preload" as="font" type="font/${ext}" href="${path}" crossorigin>`);

			link_headers.add(
				`<${encodeURI(path)}>; rel="preload"; as="font"; type="font/${ext}"; crossorigin; nopush`
			);
		}
	}

	const global = get_global_name(options);
	const { data, chunks } = data_serializer.get_data(csp);

	if (page_config.ssr && page_config.csr) {
		body += `\n\t\t\t${fetched
			.map((item) =>
				serialize_data(item, resolve_opts.filterSerializedResponseHeaders, !!state.prerendering)
			)
			.join('\n\t\t\t')}`;
	}

	if (page_config.csr) {
		const route = manifest._.client.routes?.find((r) => r.id === event.route.id) ?? null;

		if (client.uses_env_dynamic_public && state.prerendering) {
			modulepreloads.add(`${paths.app_dir}/env.js`);
		}

		if (!client.inline) {
			const included_modulepreloads = Array.from(modulepreloads, (dep) => prefixed(dep)).filter(
				(path) => resolve_opts.preload({ type: 'js', path })
			);

			for (const path of included_modulepreloads) {
				// see the kit.output.preloadStrategy option for details on why we have multiple options here
				link_headers.add(`<${encodeURI(path)}>; rel="modulepreload"; nopush`);

				if (options.preload_strategy !== 'modulepreload') {
					head += `\n\t\t<link rel="preload" as="script" crossorigin="anonymous" href="${path}">`;
				} else {
					link_tags.add(`<link rel="modulepreload" href="${path}">`);
				}
			}
		}

		if (state.prerendering && link_tags.size > 0) {
			head += Array.from(link_tags)
				.map((tag) => `\n\t\t${tag}`)
				.join('');
		}

		// prerender a `/path/to/page/__route.js` module
		if (manifest._.client.routes && state.prerendering && !state.prerendering.fallback) {
			const pathname = add_resolution_suffix(event.url.pathname);

			state.prerendering.dependencies.set(
				pathname,
				create_server_routing_response(route, event.params, new URL(pathname, event.url), manifest)
			);
		}

		const blocks = [];

		// when serving a prerendered page in an app that uses $env/dynamic/public, we must
		// import the env.js module so that it evaluates before any user code can evaluate.
		// TODO revert to using top-level await once https://bugs.webkit.org/show_bug.cgi?id=242740 is fixed
		// https://github.com/sveltejs/kit/pull/11601
		const load_env_eagerly = client.uses_env_dynamic_public && state.prerendering;

		const properties = [`base: ${base_expression}`];

		if (paths.assets) {
			properties.push(`assets: ${s(paths.assets)}`);
		}

		if (client.uses_env_dynamic_public) {
			properties.push(`env: ${load_env_eagerly ? 'null' : s(public_env)}`);
		}

		if (chunks) {
			blocks.push('const deferred = new Map();');

			properties.push(`defer: (id) => new Promise((fulfil, reject) => {
							deferred.set(id, { fulfil, reject });
						})`);

			let app_declaration = '';

			if (Object.keys(options.hooks.transport).length > 0) {
				if (client.inline) {
					app_declaration = `const app = __sveltekit_${options.version_hash}.app.app;`;
				} else if (client.app) {
					app_declaration = `const app = await import(${s(prefixed(client.app))});`;
				} else {
					app_declaration = `const { app } = await import(${s(prefixed(client.start))});`;
				}
			}

			const prelude = app_declaration
				? `${app_declaration}
							const [data, error] = fn(app);`
				: `const [data, error] = fn();`;

			// When resolving, the id might not yet be available due to the data
			// be evaluated upon init of kit, so we use a timeout to retry
			properties.push(`resolve: async (id, fn) => {
							${prelude}

							const try_to_resolve = () => {
								if (!deferred.has(id)) {
									setTimeout(try_to_resolve, 0);
									return;
								}
								const { fulfil, reject } = deferred.get(id);
								deferred.delete(id);
								if (error) reject(error);
								else fulfil(data);
							}
							try_to_resolve();
						}`);
		}

		// create this before declaring `data`, which may contain references to `${global}`
		blocks.push(`${global} = {
						${properties.join(',\n\t\t\t\t\t\t')}
					};`);

		const args = ['element'];

		blocks.push('const element = document.currentScript.parentElement;');

		if (page_config.ssr) {
			const serialized = { form: 'null', error: 'null' };

			if (form_value) {
				serialized.form = uneval_action_response(
					form_value,
					/** @type {string} */ (event.route.id),
					options.hooks.transport
				);
			}

			if (error) {
				serialized.error = devalue.uneval(error);
			}

			const hydrate = [
				`node_ids: [${branch.map(({ node }) => node.index).join(', ')}]`,
				`data: ${data}`,
				`form: ${serialized.form}`,
				`error: ${serialized.error}`
			];

			if (status !== 200) {
				hydrate.push(`status: ${status}`);
			}

			if (manifest._.client.routes) {
				if (route) {
					const stringified = generate_route_object(route, event.url, manifest).replaceAll(
						'\n',
						'\n\t\t\t\t\t\t\t'
					); // make output after it's put together with the rest more readable
					hydrate.push(`params: ${devalue.uneval(event.params)}`, `server_route: ${stringified}`);
				}
			} else if (options.embedded) {
				hydrate.push(`params: ${devalue.uneval(event.params)}`, `route: ${s(event.route)}`);
			}

			const indent = '\t'.repeat(load_env_eagerly ? 7 : 6);
			args.push(`{\n${indent}\t${hydrate.join(`,\n${indent}\t`)}\n${indent}}`);
		}

		const { remote_data: remote_cache } = event_state;

		let serialized_remote_data = '';

		if (remote_cache) {
			/** @type {Record<string, any>} */
			const remote = {};

			for (const [info, cache] of remote_cache) {
				// remote functions without an `id` aren't exported, and thus
				// cannot be called from the client
				if (!info.id) continue;

				for (const key in cache) {
					remote[create_remote_key(info.id, key)] = await cache[key];
				}
			}

			// TODO this is repeated in a few places — dedupe it
			const replacer = (/** @type {any} */ thing) => {
				for (const key in options.hooks.transport) {
					const encoded = options.hooks.transport[key].encode(thing);
					if (encoded) {
						return `app.decode('${key}', ${devalue.uneval(encoded, replacer)})`;
					}
				}
			};

			serialized_remote_data = `${global}.data = ${devalue.uneval(remote, replacer)};\n\n\t\t\t\t\t\t`;
		}

		// `client.app` is a proxy for `bundleStrategy === 'split'`
		const boot = client.inline
			? `${client.inline.script}

					${serialized_remote_data}${global}.app.start(${args.join(', ')});`
			: client.app
				? `Promise.all([
						import(${s(prefixed(client.start))}),
						import(${s(prefixed(client.app))})
					]).then(([kit, app]) => {
						${serialized_remote_data}kit.start(app, ${args.join(', ')});
					});`
				: `import(${s(prefixed(client.start))}).then((app) => {
						${serialized_remote_data}app.start(${args.join(', ')})
					});`;

		if (load_env_eagerly) {
			blocks.push(`import(${s(`${base}/${paths.app_dir}/env.js`)}).then(({ env }) => {
						${global}.env = env;

						${boot.replace(/\n/g, '\n\t')}
					});`);
		} else {
			blocks.push(boot);
		}

		if (options.service_worker) {
			let opts = DEV ? ", { type: 'module' }" : '';
			if (options.service_worker_options != null) {
				const service_worker_options = { ...options.service_worker_options };
				if (DEV) {
					service_worker_options.type = 'module';
				}
				opts = `, ${s(service_worker_options)}`;
			}

			// we use an anonymous function instead of an arrow function to support
			// older browsers (https://github.com/sveltejs/kit/pull/5417)
			blocks.push(`if ('serviceWorker' in navigator) {
						addEventListener('load', function () {
							navigator.serviceWorker.register('${prefixed('service-worker.js')}'${opts});
						});
					}`);
		}

		const init_app = `
				{
					${blocks.join('\n\n\t\t\t\t\t')}
				}
			`;
		csp.add_script(init_app);

		body += `\n\t\t\t<script${
			csp.script_needs_nonce ? ` nonce="${csp.nonce}"` : ''
		}>${init_app}</script>\n\t\t`;
	}

	const headers = new Headers({
		'x-sveltekit-page': 'true',
		'content-type': 'text/html'
	});

	if (state.prerendering) {
		// TODO read headers set with setHeaders and convert into http-equiv where possible
		const http_equiv = [];

		const csp_headers = csp.csp_provider.get_meta();
		if (csp_headers) {
			http_equiv.push(csp_headers);
		}

		if (state.prerendering.cache) {
			http_equiv.push(`<meta http-equiv="cache-control" content="${state.prerendering.cache}">`);
		}

		if (http_equiv.length > 0) {
			head = http_equiv.join('\n') + head;
		}
	} else {
		const csp_header = csp.csp_provider.get_header();
		if (csp_header) {
			headers.set('content-security-policy', csp_header);
		}
		const report_only_header = csp.report_only_provider.get_header();
		if (report_only_header) {
			headers.set('content-security-policy-report-only', report_only_header);
		}

		if (link_headers.size) {
			headers.set('link', Array.from(link_headers).join(', '));
		}
	}

	// add the content after the script/css links so the link elements are parsed first
	head += rendered.head;

	const html = options.templates.app({
		head,
		body,
		assets,
		nonce: /** @type {string} */ (csp.nonce),
		env: public_env
	});

	// TODO flush chunks as early as we can
	const transformed =
		(await resolve_opts.transformPageChunk({
			html,
			done: true
		})) || '';

	if (!chunks) {
		headers.set('etag', `"${hash(transformed)}"`);
	}

	if (DEV) {
		if (page_config.csr) {
			if (transformed.split('<!--').length < html.split('<!--').length) {
				// the \u001B stuff is ANSI codes, so that we don't need to add a library to the runtime
				// https://svelte.dev/playground/1b3f49696f0c44c881c34587f2537aa2?version=4.2.19
				console.warn(
					"\u001B[1m\u001B[31mRemoving comments in transformPageChunk can break Svelte's hydration\u001B[39m\u001B[22m"
				);
			}
		} else {
			if (chunks) {
				console.warn(
					'\u001B[1m\u001B[31mReturning promises from server `load` functions will only work if `csr === true`\u001B[39m\u001B[22m'
				);
			}
		}
	}

	return !chunks
		? text(transformed, {
				status,
				headers
			})
		: new Response(
				new ReadableStream({
					async start(controller) {
						controller.enqueue(text_encoder.encode(transformed + '\n'));
						for await (const chunk of chunks) {
							if (chunk.length) controller.enqueue(text_encoder.encode(chunk));
						}
						controller.close();
					},

					type: 'bytes'
				}),
				{
					headers
				}
			);
}
