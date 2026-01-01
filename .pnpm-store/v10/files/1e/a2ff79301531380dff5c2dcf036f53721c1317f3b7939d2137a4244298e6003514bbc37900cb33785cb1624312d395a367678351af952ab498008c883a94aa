import 'svelte'; // pick up `declare module "*.svelte"`
import 'vite/client'; // pick up `declare module "*.jpg"`, etc.
import '../types/ambient.js';

import { CompileOptions } from 'svelte/compiler';
import {
	AdapterEntry,
	CspDirectives,
	HttpMethod,
	Logger,
	MaybePromise,
	Prerendered,
	PrerenderEntryGeneratorMismatchHandlerValue,
	PrerenderHttpErrorHandlerValue,
	PrerenderMissingIdHandlerValue,
	PrerenderOption,
	RequestOptions,
	RouteSegment
} from '../types/private.js';
import { BuildData, SSRNodeLoader, SSRRoute, ValidatedConfig } from 'types';
import type { PluginOptions } from '@sveltejs/vite-plugin-svelte';

export { PrerenderOption } from '../types/private.js';

/**
 * [Adapters](https://svelte.dev/docs/kit/adapters) are responsible for taking the production build and turning it into something that can be deployed to a platform of your choosing.
 */
export interface Adapter {
	/**
	 * The name of the adapter, using for logging. Will typically correspond to the package name.
	 */
	name: string;
	/**
	 * This function is called after SvelteKit has built your app.
	 * @param builder An object provided by SvelteKit that contains methods for adapting the app
	 */
	adapt: (builder: Builder) => MaybePromise<void>;
	/**
	 * Checks called during dev and build to determine whether specific features will work in production with this adapter.
	 */
	supports?: {
		/**
		 * Test support for `read` from `$app/server`.
		 * @param details.config The merged route config
		 */
		read?: (details: { config: any; route: { id: string } }) => boolean;
	};
	/**
	 * Creates an `Emulator`, which allows the adapter to influence the environment
	 * during dev, build and prerendering.
	 */
	emulate?: () => MaybePromise<Emulator>;
}

export type LoadProperties<input extends Record<string, any> | void> = input extends void
	? undefined // needs to be undefined, because void will break intellisense
	: input extends Record<string, any>
		? input
		: unknown;

export type AwaitedActions<T extends Record<string, (...args: any) => any>> = OptionalUnion<
	{
		[Key in keyof T]: UnpackValidationError<Awaited<ReturnType<T[Key]>>>;
	}[keyof T]
>;

// Takes a union type and returns a union type where each type also has all properties
// of all possible types (typed as undefined), making accessing them more ergonomic
type OptionalUnion<
	U extends Record<string, any>, // not unknown, else interfaces don't satisfy this constraint
	A extends keyof U = U extends U ? keyof U : never
> = U extends unknown ? { [P in Exclude<A, keyof U>]?: never } & U : never;

declare const uniqueSymbol: unique symbol;

export interface ActionFailure<T extends Record<string, unknown> | undefined = undefined> {
	status: number;
	data: T;
	[uniqueSymbol]: true; // necessary or else UnpackValidationError could wrongly unpack objects with the same shape as ActionFailure
}

type UnpackValidationError<T> =
	T extends ActionFailure<infer X>
		? X
		: T extends void
			? undefined // needs to be undefined, because void will corrupt union type
			: T;

/**
 * This object is passed to the `adapt` function of adapters.
 * It contains various methods and properties that are useful for adapting the app.
 */
export interface Builder {
	/** Print messages to the console. `log.info` and `log.minor` are silent unless Vite's `logLevel` is `info`. */
	log: Logger;
	/** Remove `dir` and all its contents. */
	rimraf: (dir: string) => void;
	/** Create `dir` and any required parent directories. */
	mkdirp: (dir: string) => void;

	/** The fully resolved `svelte.config.js`. */
	config: ValidatedConfig;
	/** Information about prerendered pages and assets, if any. */
	prerendered: Prerendered;
	/** An array of all routes (including prerendered) */
	routes: RouteDefinition[];

	// TODO 3.0 remove this method
	/**
	 * Create separate functions that map to one or more routes of your app.
	 * @param fn A function that groups a set of routes into an entry point
	 * @deprecated Use `builder.routes` instead
	 */
	createEntries: (fn: (route: RouteDefinition) => AdapterEntry) => Promise<void>;

	/**
	 * Find all the assets imported by server files belonging to `routes`
	 */
	findServerAssets: (routes: RouteDefinition[]) => string[];

	/**
	 * Generate a fallback page for a static webserver to use when no route is matched. Useful for single-page apps.
	 */
	generateFallback: (dest: string) => Promise<void>;

	/**
	 * Generate a module exposing build-time environment variables as `$env/dynamic/public`.
	 */
	generateEnvModule: () => void;

	/**
	 * Generate a server-side manifest to initialise the SvelteKit [server](https://svelte.dev/docs/kit/@sveltejs-kit#Server) with.
	 * @param opts a relative path to the base directory of the app and optionally in which format (esm or cjs) the manifest should be generated
	 */
	generateManifest: (opts: { relativePath: string; routes?: RouteDefinition[] }) => string;

	/**
	 * Resolve a path to the `name` directory inside `outDir`, e.g. `/path/to/.svelte-kit/my-adapter`.
	 * @param name path to the file, relative to the build directory
	 */
	getBuildDirectory: (name: string) => string;
	/** Get the fully resolved path to the directory containing client-side assets, including the contents of your `static` directory. */
	getClientDirectory: () => string;
	/** Get the fully resolved path to the directory containing server-side code. */
	getServerDirectory: () => string;
	/** Get the application path including any configured `base` path, e.g. `my-base-path/_app`. */
	getAppPath: () => string;

	/**
	 * Write client assets to `dest`.
	 * @param dest the destination folder
	 * @returns an array of files written to `dest`
	 */
	writeClient: (dest: string) => string[];
	/**
	 * Write prerendered files to `dest`.
	 * @param dest the destination folder
	 * @returns an array of files written to `dest`
	 */
	writePrerendered: (dest: string) => string[];
	/**
	 * Write server-side code to `dest`.
	 * @param dest the destination folder
	 * @returns an array of files written to `dest`
	 */
	writeServer: (dest: string) => string[];
	/**
	 * Copy a file or directory.
	 * @param from the source file or directory
	 * @param to the destination file or directory
	 * @param opts.filter a function to determine whether a file or directory should be copied
	 * @param opts.replace a map of strings to replace
	 * @returns an array of files that were copied
	 */
	copy: (
		from: string,
		to: string,
		opts?: {
			filter?(basename: string): boolean;
			replace?: Record<string, string>;
		}
	) => string[];

	/**
	 * Compress files in `directory` with gzip and brotli, where appropriate. Generates `.gz` and `.br` files alongside the originals.
	 * @param {string} directory The directory containing the files to be compressed
	 */
	compress: (directory: string) => Promise<void>;
}

export interface Config {
	/**
	 * Options passed to [`svelte.compile`](https://svelte.dev/docs/svelte/svelte-compiler#CompileOptions).
	 * @default {}
	 */
	compilerOptions?: CompileOptions;
	/**
	 * List of file extensions that should be treated as Svelte files.
	 * @default [".svelte"]
	 */
	extensions?: string[];
	/** SvelteKit options */
	kit?: KitConfig;
	/** Preprocessor options, if any. Preprocessing can alternatively also be done through Vite's preprocessor capabilities. */
	preprocess?: any;
	/** `vite-plugin-svelte` plugin options. */
	vitePlugin?: PluginOptions;
	/** Any additional options required by tooling that integrates with Svelte. */
	[key: string]: any;
}

export interface Cookies {
	/**
	 * Gets a cookie that was previously set with `cookies.set`, or from the request headers.
	 * @param name the name of the cookie
	 * @param opts the options, passed directly to `cookie.parse`. See documentation [here](https://github.com/jshttp/cookie#cookieparsestr-options)
	 */
	get: (name: string, opts?: import('cookie').CookieParseOptions) => string | undefined;

	/**
	 * Gets all cookies that were previously set with `cookies.set`, or from the request headers.
	 * @param opts the options, passed directly to `cookie.parse`. See documentation [here](https://github.com/jshttp/cookie#cookieparsestr-options)
	 */
	getAll: (opts?: import('cookie').CookieParseOptions) => Array<{ name: string; value: string }>;

	/**
	 * Sets a cookie. This will add a `set-cookie` header to the response, but also make the cookie available via `cookies.get` or `cookies.getAll` during the current request.
	 *
	 * The `httpOnly` and `secure` options are `true` by default (except on http://localhost, where `secure` is `false`), and must be explicitly disabled if you want cookies to be readable by client-side JavaScript and/or transmitted over HTTP. The `sameSite` option defaults to `lax`.
	 *
	 * You must specify a `path` for the cookie. In most cases you should explicitly set `path: '/'` to make the cookie available throughout your app. You can use relative paths, or set `path: ''` to make the cookie only available on the current path and its children
	 * @param name the name of the cookie
	 * @param value the cookie value
	 * @param opts the options, passed directly to `cookie.serialize`. See documentation [here](https://github.com/jshttp/cookie#cookieserializename-value-options)
	 */
	set: (
		name: string,
		value: string,
		opts: import('cookie').CookieSerializeOptions & { path: string }
	) => void;

	/**
	 * Deletes a cookie by setting its value to an empty string and setting the expiry date in the past.
	 *
	 * You must specify a `path` for the cookie. In most cases you should explicitly set `path: '/'` to make the cookie available throughout your app. You can use relative paths, or set `path: ''` to make the cookie only available on the current path and its children
	 * @param name the name of the cookie
	 * @param opts the options, passed directly to `cookie.serialize`. The `path` must match the path of the cookie you want to delete. See documentation [here](https://github.com/jshttp/cookie#cookieserializename-value-options)
	 */
	delete: (name: string, opts: import('cookie').CookieSerializeOptions & { path: string }) => void;

	/**
	 * Serialize a cookie name-value pair into a `Set-Cookie` header string, but don't apply it to the response.
	 *
	 * The `httpOnly` and `secure` options are `true` by default (except on http://localhost, where `secure` is `false`), and must be explicitly disabled if you want cookies to be readable by client-side JavaScript and/or transmitted over HTTP. The `sameSite` option defaults to `lax`.
	 *
	 * You must specify a `path` for the cookie. In most cases you should explicitly set `path: '/'` to make the cookie available throughout your app. You can use relative paths, or set `path: ''` to make the cookie only available on the current path and its children
	 *
	 * @param name the name of the cookie
	 * @param value the cookie value
	 * @param opts the options, passed directly to `cookie.serialize`. See documentation [here](https://github.com/jshttp/cookie#cookieserializename-value-options)
	 */
	serialize: (
		name: string,
		value: string,
		opts: import('cookie').CookieSerializeOptions & { path: string }
	) => string;
}

/**
 * A collection of functions that influence the environment during dev, build and prerendering
 */
export interface Emulator {
	/**
	 * A function that is called with the current route `config` and `prerender` option
	 * and returns an `App.Platform` object
	 */
	platform?(details: { config: any; prerender: PrerenderOption }): MaybePromise<App.Platform>;
}

export interface KitConfig {
	/**
	 * Your [adapter](https://svelte.dev/docs/kit/adapters) is run when executing `vite build`. It determines how the output is converted for different platforms.
	 * @default undefined
	 */
	adapter?: Adapter;
	/**
	 * An object containing zero or more aliases used to replace values in `import` statements. These aliases are automatically passed to Vite and TypeScript.
	 *
	 * ```js
	 * /// file: svelte.config.js
	 * /// type: import('@sveltejs/kit').Config
	 * const config = {
	 *   kit: {
	 *     alias: {
	 *       // this will match a file
	 *       'my-file': 'path/to/my-file.js',
	 *
	 *       // this will match a directory and its contents
	 *       // (`my-directory/x` resolves to `path/to/my-directory/x`)
	 *       'my-directory': 'path/to/my-directory',
	 *
	 *       // an alias ending /* will only match
	 *       // the contents of a directory, not the directory itself
	 *       'my-directory/*': 'path/to/my-directory/*'
	 *     }
	 *   }
	 * };
	 * ```
	 *
	 * > [!NOTE] The built-in `$lib` alias is controlled by `config.kit.files.lib` as it is used for packaging.
	 *
	 * > [!NOTE] You will need to run `npm run dev` to have SvelteKit automatically generate the required alias configuration in `jsconfig.json` or `tsconfig.json`.
	 * @default {}
	 */
	alias?: Record<string, string>;
	/**
	 * The directory where SvelteKit keeps its stuff, including static assets (such as JS and CSS) and internally-used routes.
	 *
	 * If `paths.assets` is specified, there will be two app directories — `${paths.assets}/${appDir}` and `${paths.base}/${appDir}`.
	 * @default "_app"
	 */
	appDir?: string;
	/**
	 * [Content Security Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) configuration. CSP helps to protect your users against cross-site scripting (XSS) attacks, by limiting the places resources can be loaded from. For example, a configuration like this...
	 *
	 * ```js
	 * /// file: svelte.config.js
	 * /// type: import('@sveltejs/kit').Config
	 * const config = {
	 *   kit: {
	 *     csp: {
	 *       directives: {
	 *         'script-src': ['self']
	 *       },
	 *       // must be specified with either the `report-uri` or `report-to` directives, or both
	 *       reportOnly: {
	 *         'script-src': ['self'],
	 *         'report-uri': ['/']
	 *       }
	 *     }
	 *   }
	 * };
	 *
	 * export default config;
	 * ```
	 *
	 * ...would prevent scripts loading from external sites. SvelteKit will augment the specified directives with nonces or hashes (depending on `mode`) for any inline styles and scripts it generates.
	 *
	 * To add a nonce for scripts and links manually included in `src/app.html`, you may use the placeholder `%sveltekit.nonce%` (for example `<script nonce="%sveltekit.nonce%">`).
	 *
	 * When pages are prerendered, the CSP header is added via a `<meta http-equiv>` tag (note that in this case, `frame-ancestors`, `report-uri` and `sandbox` directives will be ignored).
	 *
	 * > [!NOTE] When `mode` is `'auto'`, SvelteKit will use nonces for dynamically rendered pages and hashes for prerendered pages. Using nonces with prerendered pages is insecure and therefore forbidden.
	 *
	 * > [!NOTE] Note that most [Svelte transitions](https://svelte.dev/tutorial/svelte/transition) work by creating an inline `<style>` element. If you use these in your app, you must either leave the `style-src` directive unspecified or add `unsafe-inline`.
	 *
	 * If this level of configuration is insufficient and you have more dynamic requirements, you can use the [`handle` hook](https://svelte.dev/docs/kit/hooks#Server-hooks-handle) to roll your own CSP.
	 */
	csp?: {
		/**
		 * Whether to use hashes or nonces to restrict `<script>` and `<style>` elements. `'auto'` will use hashes for prerendered pages, and nonces for dynamically rendered pages.
		 */
		mode?: 'hash' | 'nonce' | 'auto';
		/**
		 * Directives that will be added to `Content-Security-Policy` headers.
		 */
		directives?: CspDirectives;
		/**
		 * Directives that will be added to `Content-Security-Policy-Report-Only` headers.
		 */
		reportOnly?: CspDirectives;
	};
	/**
	 * Protection against [cross-site request forgery (CSRF)](https://owasp.org/www-community/attacks/csrf) attacks.
	 */
	csrf?: {
		/**
		 * Whether to check the incoming `origin` header for `POST`, `PUT`, `PATCH`, or `DELETE` form submissions and verify that it matches the server's origin.
		 *
		 * To allow people to make `POST`, `PUT`, `PATCH`, or `DELETE` requests with a `Content-Type` of `application/x-www-form-urlencoded`, `multipart/form-data`, or `text/plain` to your app from other origins, you will need to disable this option. Be careful!
		 * @default true
		 */
		checkOrigin?: boolean;
	};
	/**
	 * Whether or not the app is embedded inside a larger app. If `true`, SvelteKit will add its event listeners related to navigation etc on the parent of `%sveltekit.body%` instead of `window`, and will pass `params` from the server rather than inferring them from `location.pathname`.
	 * Note that it is generally not supported to embed multiple SvelteKit apps on the same page and use client-side SvelteKit features within them (things such as pushing to the history state assume a single instance).
	 * @default false
	 */
	embedded?: boolean;
	/**
	 * Environment variable configuration
	 */
	env?: {
		/**
		 * The directory to search for `.env` files.
		 * @default "."
		 */
		dir?: string;
		/**
		 * A prefix that signals that an environment variable is safe to expose to client-side code. See [`$env/static/public`](https://svelte.dev/docs/kit/$env-static-public) and [`$env/dynamic/public`](https://svelte.dev/docs/kit/$env-dynamic-public). Note that Vite's [`envPrefix`](https://vitejs.dev/config/shared-options.html#envprefix) must be set separately if you are using Vite's environment variable handling - though use of that feature should generally be unnecessary.
		 * @default "PUBLIC_"
		 */
		publicPrefix?: string;
		/**
		 * A prefix that signals that an environment variable is unsafe to expose to client-side code. Environment variables matching neither the public nor the private prefix will be discarded completely. See [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private) and [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private).
		 * @default ""
		 * @since 1.21.0
		 */
		privatePrefix?: string;
	};
	/**
	 * Where to find various files within your project.
	 */
	files?: {
		/**
		 * a place to put static files that should have stable URLs and undergo no processing, such as `favicon.ico` or `manifest.json`
		 * @default "static"
		 */
		assets?: string;
		hooks?: {
			/**
			 * The location of your client [hooks](https://svelte.dev/docs/kit/hooks).
			 * @default "src/hooks.client"
			 */
			client?: string;
			/**
			 * The location of your server [hooks](https://svelte.dev/docs/kit/hooks).
			 * @default "src/hooks.server"
			 */
			server?: string;
			/**
			 * The location of your universal [hooks](https://svelte.dev/docs/kit/hooks).
			 * @default "src/hooks"
			 * @since 2.3.0
			 */
			universal?: string;
		};
		/**
		 * your app's internal library, accessible throughout the codebase as `$lib`
		 * @default "src/lib"
		 */
		lib?: string;
		/**
		 * a directory containing [parameter matchers](https://svelte.dev/docs/kit/advanced-routing#Matching)
		 * @default "src/params"
		 */
		params?: string;
		/**
		 * the files that define the structure of your app (see [Routing](https://svelte.dev/docs/kit/routing))
		 * @default "src/routes"
		 */
		routes?: string;
		/**
		 * the location of your service worker's entry point (see [Service workers](https://svelte.dev/docs/kit/service-workers))
		 * @default "src/service-worker"
		 */
		serviceWorker?: string;
		/**
		 * the location of the template for HTML responses
		 * @default "src/app.html"
		 */
		appTemplate?: string;
		/**
		 * the location of the template for fallback error responses
		 * @default "src/error.html"
		 */
		errorTemplate?: string;
	};
	/**
	 * Inline CSS inside a `<style>` block at the head of the HTML. This option is a number that specifies the maximum length of a CSS file in UTF-16 code units, as specified by the [String.length](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/length) property, to be inlined. All CSS files needed for the page and smaller than this value are merged and inlined in a `<style>` block.
	 *
	 * > [!NOTE] This results in fewer initial requests and can improve your [First Contentful Paint](https://web.dev/first-contentful-paint) score. However, it generates larger HTML output and reduces the effectiveness of browser caches. Use it advisedly.
	 * @default 0
	 */
	inlineStyleThreshold?: number;
	/**
	 * An array of file extensions that SvelteKit will treat as modules. Files with extensions that match neither `config.extensions` nor `config.kit.moduleExtensions` will be ignored by the router.
	 * @default [".js", ".ts"]
	 */
	moduleExtensions?: string[];
	/**
	 * The directory that SvelteKit writes files to during `dev` and `build`. You should exclude this directory from version control.
	 * @default ".svelte-kit"
	 */
	outDir?: string;
	/**
	 * Options related to the build output format
	 */
	output?: {
		/**
		 * SvelteKit will preload the JavaScript modules needed for the initial page to avoid import 'waterfalls', resulting in faster application startup. There
		 * are three strategies with different trade-offs:
		 * - `modulepreload` - uses `<link rel="modulepreload">`. This delivers the best results in Chromium-based browsers, in Firefox 115+, and Safari 17+. It is ignored in older browsers.
		 * - `preload-js` - uses `<link rel="preload">`. Prevents waterfalls in Chromium and Safari, but Chromium will parse each module twice (once as a script, once as a module). Causes modules to be requested twice in Firefox. This is a good setting if you want to maximise performance for users on iOS devices at the cost of a very slight degradation for Chromium users.
		 * - `preload-mjs` - uses `<link rel="preload">` but with the `.mjs` extension which prevents double-parsing in Chromium. Some static webservers will fail to serve .mjs files with a `Content-Type: application/javascript` header, which will cause your application to break. If that doesn't apply to you, this is the option that will deliver the best performance for the largest number of users, until `modulepreload` is more widely supported.
		 * @default "modulepreload"
		 * @since 1.8.4
		 */
		preloadStrategy?: 'modulepreload' | 'preload-js' | 'preload-mjs';
		/**
		 * The bundle strategy option affects how your app's JavaScript and CSS files are loaded.
		 * - If `'split'`, splits the app up into multiple .js/.css files so that they are loaded lazily as the user navigates around the app. This is the default, and is recommended for most scenarios.
		 * - If `'single'`, creates just one .js bundle and one .css file containing code for the entire app.
		 * - If `'inline'`, inlines all JavaScript and CSS of the entire app into the HTML. The result is usable without a server (i.e. you can just open the file in your browser).
		 *
		 * When using `'split'`, you can also adjust the bundling behaviour by setting [`output.experimentalMinChunkSize`](https://rollupjs.org/configuration-options/#output-experimentalminchunksize) and [`output.manualChunks`](https://rollupjs.org/configuration-options/#output-manualchunks) inside your Vite config's [`build.rollupOptions`](https://vite.dev/config/build-options.html#build-rollupoptions).
		 *
		 * If you want to inline your assets, you'll need to set Vite's [`build.assetsInlineLimit`](https://vite.dev/config/build-options.html#build-assetsinlinelimit) option to an appropriate size then import your assets through Vite.
		 *
		 * ```js
		 * /// file: vite.config.js
		 * import { sveltekit } from '@sveltejs/kit/vite';
		 * import { defineConfig } from 'vite';
		 *
		 * export default defineConfig({
		 *   plugins: [sveltekit()],
		 *   build: {
		 *     // inline all imported assets
		 *     assetsInlineLimit: Infinity
		 *   }
		 * });
		 * ```
		 *
		 * ```svelte
		 * /// file: src/routes/+layout.svelte
		 * <script>
		 *   // import the asset through Vite
		 *   import favicon from './favicon.png';
		 * </script>
		 *
		 * <svelte:head>
		 *   <!-- this asset will be inlined as a base64 URL -->
		 *   <link rel="icon" href={favicon} />
		 * </svelte:head>
		 * ```
		 * @default 'split'
		 * @since 2.13.0
		 */
		bundleStrategy?: 'split' | 'single' | 'inline';
	};
	paths?: {
		/**
		 * An absolute path that your app's files are served from. This is useful if your files are served from a storage bucket of some kind.
		 * @default ""
		 */
		assets?: '' | `http://${string}` | `https://${string}`;
		/**
		 * A root-relative path that must start, but not end with `/` (e.g. `/base-path`), unless it is the empty string. This specifies where your app is served from and allows the app to live on a non-root path. Note that you need to prepend all your root-relative links with the base value or they will point to the root of your domain, not your `base` (this is how the browser works). You can use [`base` from `$app/paths`](https://svelte.dev/docs/kit/$app-paths#base) for that: `<a href="{base}/your-page">Link</a>`. If you find yourself writing this often, it may make sense to extract this into a reusable component.
		 * @default ""
		 */
		base?: '' | `/${string}`;
		/**
		 * Whether to use relative asset paths.
		 *
		 * If `true`, `base` and `assets` imported from `$app/paths` will be replaced with relative asset paths during server-side rendering, resulting in more portable HTML.
		 * If `false`, `%sveltekit.assets%` and references to build artifacts will always be root-relative paths, unless `paths.assets` is an external URL
		 *
		 * [Single-page app](https://svelte.dev/docs/kit/single-page-apps) fallback pages will always use absolute paths, regardless of this setting.
		 *
		 * If your app uses a `<base>` element, you should set this to `false`, otherwise asset URLs will incorrectly be resolved against the `<base>` URL rather than the current page.
		 *
		 * In 1.0, `undefined` was a valid value, which was set by default. In that case, if `paths.assets` was not external, SvelteKit would replace `%sveltekit.assets%` with a relative path and use relative paths to reference build artifacts, but `base` and `assets` imported from `$app/paths` would be as specified in your config.
		 *
		 * @default true
		 * @since 1.9.0
		 */
		relative?: boolean;
	};
	/**
	 * See [Prerendering](https://svelte.dev/docs/kit/page-options#prerender).
	 */
	prerender?: {
		/**
		 * How many pages can be prerendered simultaneously. JS is single-threaded, but in cases where prerendering performance is network-bound (for example loading content from a remote CMS) this can speed things up by processing other tasks while waiting on the network response.
		 * @default 1
		 */
		concurrency?: number;
		/**
		 * Whether SvelteKit should find pages to prerender by following links from `entries`.
		 * @default true
		 */
		crawl?: boolean;
		/**
		 * An array of pages to prerender, or start crawling from (if `crawl: true`). The `*` string includes all routes containing no required `[parameters]`  with optional parameters included as being empty (since SvelteKit doesn't know what value any parameters should have).
		 * @default ["*"]
		 */
		entries?: Array<'*' | `/${string}`>;
		/**
		 * How to respond to HTTP errors encountered while prerendering the app.
		 *
		 * - `'fail'` — fail the build
		 * - `'ignore'` - silently ignore the failure and continue
		 * - `'warn'` — continue, but print a warning
		 * - `(details) => void` — a custom error handler that takes a `details` object with `status`, `path`, `referrer`, `referenceType` and `message` properties. If you `throw` from this function, the build will fail
		 *
		 * ```js
		 * /// file: svelte.config.js
		 * /// type: import('@sveltejs/kit').Config
		 * const config = {
		 *   kit: {
		 *     prerender: {
		 *       handleHttpError: ({ path, referrer, message }) => {
		 *         // ignore deliberate link to shiny 404 page
		 *         if (path === '/not-found' && referrer === '/blog/how-we-built-our-404-page') {
		 *           return;
		 *         }
		 *
		 *         // otherwise fail the build
		 *         throw new Error(message);
		 *       }
		 *     }
		 *   }
		 * };
		 * ```
		 *
		 * @default "fail"
		 * @since 1.15.7
		 */
		handleHttpError?: PrerenderHttpErrorHandlerValue;
		/**
		 * How to respond when hash links from one prerendered page to another don't correspond to an `id` on the destination page.
		 *
		 * - `'fail'` — fail the build
		 * - `'ignore'` - silently ignore the failure and continue
		 * - `'warn'` — continue, but print a warning
		 * - `(details) => void` — a custom error handler that takes a `details` object with `path`, `id`, `referrers` and `message` properties. If you `throw` from this function, the build will fail
		 *
		 * @default "fail"
		 * @since 1.15.7
		 */
		handleMissingId?: PrerenderMissingIdHandlerValue;
		/**
		 * How to respond when an entry generated by the `entries` export doesn't match the route it was generated from.
		 *
		 * - `'fail'` — fail the build
		 * - `'ignore'` - silently ignore the failure and continue
		 * - `'warn'` — continue, but print a warning
		 * - `(details) => void` — a custom error handler that takes a `details` object with `generatedFromId`, `entry`, `matchedId` and `message` properties. If you `throw` from this function, the build will fail
		 *
		 * @default "fail"
		 * @since 1.16.0
		 */
		handleEntryGeneratorMismatch?: PrerenderEntryGeneratorMismatchHandlerValue;
		/**
		 * The value of `url.origin` during prerendering; useful if it is included in rendered content.
		 * @default "http://sveltekit-prerender"
		 */
		origin?: string;
	};
	router?: {
		/**
		 * What type of client-side router to use.
		 * - `'pathname'` is the default and means the current URL pathname determines the route
		 * - `'hash'` means the route is determined by `location.hash`. In this case, SSR and prerendering are disabled. This is only recommended if `pathname` is not an option, for example because you don't control the webserver where your app is deployed.
		 *   It comes with some caveats: you can't use server-side rendering (or indeed any server logic), and you have to make sure that the links in your app all start with #/, or they won't work. Beyond that, everything works exactly like a normal SvelteKit app.
		 *
		 * @default "pathname"
		 * @since 2.14.0
		 */
		type?: 'pathname' | 'hash';
		/**
		 * How to determine which route to load when navigating to a new page.
		 *
		 * By default, SvelteKit will serve a route manifest to the browser.
		 * When navigating, this manifest is used (along with the `reroute` hook, if it exists) to determine which components to load and which `load` functions to run.
		 * Because everything happens on the client, this decision can be made immediately. The drawback is that the manifest needs to be
		 * loaded and parsed before the first navigation can happen, which may have an impact if your app contains many routes.
		 *
		 * Alternatively, SvelteKit can determine the route on the server. This means that for every navigation to a path that has not yet been visited, the server will be asked to determine the route.
		 * This has several advantages:
		 * - The client does not need to load the routing manifest upfront, which can lead to faster initial page loads
		 * - The list of routes is hidden from public view
		 * - The server has an opportunity to intercept each navigation (for example through a middleware), enabling (for example) A/B testing opaque to SvelteKit

		 * The drawback is that for unvisited paths, resolution will take slightly longer (though this is mitigated by [preloading](https://svelte.dev/docs/kit/link-options#data-sveltekit-preload-data)).
		 *
		 * > [!NOTE] When using server-side route resolution and prerendering, the resolution is prerendered along with the route itself.
		 *
		 * @default "client"
		 * @since 2.17.0
		 */
		resolution?: 'client' | 'server';
	};
	serviceWorker?: {
		/**
		 * Whether to automatically register the service worker, if it exists.
		 * @default true
		 */
		register?: boolean;
		/**
		 * Determine which files in your `static` directory will be available in `$service-worker.files`.
		 * @default (filename) => !/\.DS_Store/.test(filename)
		 */
		files?(filepath: string): boolean;
	};
	typescript?: {
		/**
		 * A function that allows you to edit the generated `tsconfig.json`. You can mutate the config (recommended) or return a new one.
		 * This is useful for extending a shared `tsconfig.json` in a monorepo root, for example.
		 * @default (config) => config
		 * @since 1.3.0
		 */
		config?: (config: Record<string, any>) => Record<string, any> | void;
	};
	/**
	 * Client-side navigation can be buggy if you deploy a new version of your app while people are using it. If the code for the new page is already loaded, it may have stale content; if it isn't, the app's route manifest may point to a JavaScript file that no longer exists.
	 * SvelteKit helps you solve this problem through version management.
	 * If SvelteKit encounters an error while loading the page and detects that a new version has been deployed (using the `name` specified here, which defaults to a timestamp of the build) it will fall back to traditional full-page navigation.
	 * Not all navigations will result in an error though, for example if the JavaScript for the next page is already loaded. If you still want to force a full-page navigation in these cases, use techniques such as setting the `pollInterval` and then using `beforeNavigate`:
	 * ```html
	 * /// file: +layout.svelte
	 * <script>
	 *   import { beforeNavigate } from '$app/navigation';
	 *   import { updated } from '$app/state';
	 *
	 *   beforeNavigate(({ willUnload, to }) => {
	 *     if (updated.current && !willUnload && to?.url) {
	 *       location.href = to.url.href;
	 *     }
	 *   });
	 * </script>
	 * ```
	 *
	 * If you set `pollInterval` to a non-zero value, SvelteKit will poll for new versions in the background and set the value of [`updated.current`](https://svelte.dev/docs/kit/$app-state#updated) `true` when it detects one.
	 */
	version?: {
		/**
		 * The current app version string. If specified, this must be deterministic (e.g. a commit ref rather than `Math.random()` or `Date.now().toString()`), otherwise defaults to a timestamp of the build.
		 *
		 * For example, to use the current commit hash, you could do use `git rev-parse HEAD`:
		 *
		 * ```js
		 * /// file: svelte.config.js
		 * import * as child_process from 'node:child_process';
		 *
		 * export default {
		 *   kit: {
		 *     version: {
		 *       name: child_process.execSync('git rev-parse HEAD').toString().trim()
		 *     }
		 *   }
		 * };
		 * ```
		 */
		name?: string;
		/**
		 * The interval in milliseconds to poll for version changes. If this is `0`, no polling occurs.
		 * @default 0
		 */
		pollInterval?: number;
	};
}

/**
 * The [`handle`](https://svelte.dev/docs/kit/hooks#Server-hooks-handle) hook runs every time the SvelteKit server receives a [request](https://svelte.dev/docs/kit/web-standards#Fetch-APIs-Request) and
 * determines the [response](https://svelte.dev/docs/kit/web-standards#Fetch-APIs-Response).
 * It receives an `event` object representing the request and a function called `resolve`, which renders the route and generates a `Response`.
 * This allows you to modify response headers or bodies, or bypass SvelteKit entirely (for implementing routes programmatically, for example).
 */
export type Handle = (input: {
	event: RequestEvent;
	resolve: (event: RequestEvent, opts?: ResolveOptions) => MaybePromise<Response>;
}) => MaybePromise<Response>;

/**
 * The server-side [`handleError`](https://svelte.dev/docs/kit/hooks#Shared-hooks-handleError) hook runs when an unexpected error is thrown while responding to a request.
 *
 * If an unexpected error is thrown during loading or rendering, this function will be called with the error and the event.
 * Make sure that this function _never_ throws an error.
 */
export type HandleServerError = (input: {
	error: unknown;
	event: RequestEvent;
	status: number;
	message: string;
}) => MaybePromise<void | App.Error>;

/**
 * The client-side [`handleError`](https://svelte.dev/docs/kit/hooks#Shared-hooks-handleError) hook runs when an unexpected error is thrown while navigating.
 *
 * If an unexpected error is thrown during loading or the following render, this function will be called with the error and the event.
 * Make sure that this function _never_ throws an error.
 */
export type HandleClientError = (input: {
	error: unknown;
	event: NavigationEvent;
	status: number;
	message: string;
}) => MaybePromise<void | App.Error>;

/**
 * The [`handleFetch`](https://svelte.dev/docs/kit/hooks#Server-hooks-handleFetch) hook allows you to modify (or replace) the result of an [`event.fetch`](https://svelte.dev/docs/kit/load#Making-fetch-requests) call that runs on the server (or during prerendering) inside an endpoint, `load`, `action`, `handle`, `handleError` or `reroute`.
 */
export type HandleFetch = (input: {
	event: RequestEvent;
	request: Request;
	fetch: typeof fetch;
}) => MaybePromise<Response>;

/**
 * The [`init`](https://svelte.dev/docs/kit/hooks#Shared-hooks-init) will be invoked before the server responds to its first request
 * @since 2.10.0
 */
export type ServerInit = () => MaybePromise<void>;

/**
 * The [`init`](https://svelte.dev/docs/kit/hooks#Shared-hooks-init) will be invoked once the app starts in the browser
 * @since 2.10.0
 */
export type ClientInit = () => MaybePromise<void>;

/**
 * The [`reroute`](https://svelte.dev/docs/kit/hooks#Universal-hooks-reroute) hook allows you to modify the URL before it is used to determine which route to render.
 * @since 2.3.0
 */
export type Reroute = (event: { url: URL; fetch: typeof fetch }) => MaybePromise<void | string>;

/**
 * The [`transport`](https://svelte.dev/docs/kit/hooks#Universal-hooks-transport) hook allows you to transport custom types across the server/client boundary.
 *
 * Each transporter has a pair of `encode` and `decode` functions. On the server, `encode` determines whether a value is an instance of the custom type and, if so, returns a non-falsy encoding of the value which can be an object or an array (or `false` otherwise).
 *
 * In the browser, `decode` turns the encoding back into an instance of the custom type.
 *
 * ```ts
 * import type { Transport } from '@sveltejs/kit';
 *
 * declare class MyCustomType {
 * 	data: any
 * }
 *
 * // hooks.js
 * export const transport: Transport = {
 * 	MyCustomType: {
 * 		encode: (value) => value instanceof MyCustomType && [value.data],
 * 		decode: ([data]) => new MyCustomType(data)
 * 	}
 * };
 * ```
 * @since 2.11.0
 */
export type Transport = Record<string, Transporter>;

/**
 * A member of the [`transport`](https://svelte.dev/docs/kit/hooks#Universal-hooks-transport) hook.
 */
export interface Transporter<
	T = any,
	U = Exclude<any, false | 0 | '' | null | undefined | typeof NaN>
> {
	encode: (value: T) => false | U;
	decode: (data: U) => T;
}

/**
 * The generic form of `PageLoad` and `LayoutLoad`. You should import those from `./$types` (see [generated types](https://svelte.dev/docs/kit/types#Generated-types))
 * rather than using `Load` directly.
 */
export type Load<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	InputData extends Record<string, unknown> | null = Record<string, any> | null,
	ParentData extends Record<string, unknown> = Record<string, any>,
	OutputData extends Record<string, unknown> | void = Record<string, any> | void,
	RouteId extends string | null = string | null
> = (event: LoadEvent<Params, InputData, ParentData, RouteId>) => MaybePromise<OutputData>;

/**
 * The generic form of `PageLoadEvent` and `LayoutLoadEvent`. You should import those from `./$types` (see [generated types](https://svelte.dev/docs/kit/types#Generated-types))
 * rather than using `LoadEvent` directly.
 */
export interface LoadEvent<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	Data extends Record<string, unknown> | null = Record<string, any> | null,
	ParentData extends Record<string, unknown> = Record<string, any>,
	RouteId extends string | null = string | null
> extends NavigationEvent<Params, RouteId> {
	/**
	 * `fetch` is equivalent to the [native `fetch` web API](https://developer.mozilla.org/en-US/docs/Web/API/fetch), with a few additional features:
	 *
	 * - It can be used to make credentialed requests on the server, as it inherits the `cookie` and `authorization` headers for the page request.
	 * - It can make relative requests on the server (ordinarily, `fetch` requires a URL with an origin when used in a server context).
	 * - Internal requests (e.g. for `+server.js` routes) go directly to the handler function when running on the server, without the overhead of an HTTP call.
	 * - During server-side rendering, the response will be captured and inlined into the rendered HTML by hooking into the `text` and `json` methods of the `Response` object. Note that headers will _not_ be serialized, unless explicitly included via [`filterSerializedResponseHeaders`](https://svelte.dev/docs/kit/hooks#Server-hooks-handle)
	 * - During hydration, the response will be read from the HTML, guaranteeing consistency and preventing an additional network request.
	 *
	 * You can learn more about making credentialed requests with cookies [here](https://svelte.dev/docs/kit/load#Cookies)
	 */
	fetch: typeof fetch;
	/**
	 * Contains the data returned by the route's server `load` function (in `+layout.server.js` or `+page.server.js`), if any.
	 */
	data: Data;
	/**
	 * If you need to set headers for the response, you can do so using the this method. This is useful if you want the page to be cached, for example:
	 *
	 *	```js
	 *	/// file: src/routes/blog/+page.js
	 *	export async function load({ fetch, setHeaders }) {
	 *		const url = `https://cms.example.com/articles.json`;
	 *		const response = await fetch(url);
	 *
	 *		setHeaders({
	 *			age: response.headers.get('age'),
	 *			'cache-control': response.headers.get('cache-control')
	 *		});
	 *
	 *		return response.json();
	 *	}
	 *	```
	 *
	 * Setting the same header multiple times (even in separate `load` functions) is an error — you can only set a given header once.
	 *
	 * You cannot add a `set-cookie` header with `setHeaders` — use the [`cookies`](https://svelte.dev/docs/kit/@sveltejs-kit#Cookies) API in a server-only `load` function instead.
	 *
	 * `setHeaders` has no effect when a `load` function runs in the browser.
	 */
	setHeaders: (headers: Record<string, string>) => void;
	/**
	 * `await parent()` returns data from parent `+layout.js` `load` functions.
	 * Implicitly, a missing `+layout.js` is treated as a `({ data }) => data` function, meaning that it will return and forward data from parent `+layout.server.js` files.
	 *
	 * Be careful not to introduce accidental waterfalls when using `await parent()`. If for example you only want to merge parent data into the returned output, call it _after_ fetching your other data.
	 */
	parent: () => Promise<ParentData>;
	/**
	 * This function declares that the `load` function has a _dependency_ on one or more URLs or custom identifiers, which can subsequently be used with [`invalidate()`](https://svelte.dev/docs/kit/$app-navigation#invalidate) to cause `load` to rerun.
	 *
	 * Most of the time you won't need this, as `fetch` calls `depends` on your behalf — it's only necessary if you're using a custom API client that bypasses `fetch`.
	 *
	 * URLs can be absolute or relative to the page being loaded, and must be [encoded](https://developer.mozilla.org/en-US/docs/Glossary/percent-encoding).
	 *
	 * Custom identifiers have to be prefixed with one or more lowercase letters followed by a colon to conform to the [URI specification](https://www.rfc-editor.org/rfc/rfc3986.html).
	 *
	 * The following example shows how to use `depends` to register a dependency on a custom identifier, which is `invalidate`d after a button click, making the `load` function rerun.
	 *
	 * ```js
	 * /// file: src/routes/+page.js
	 * let count = 0;
	 * export async function load({ depends }) {
	 * 	depends('increase:count');
	 *
	 * 	return { count: count++ };
	 * }
	 * ```
	 *
	 * ```html
	 * /// file: src/routes/+page.svelte
	 * <script>
	 * 	import { invalidate } from '$app/navigation';
	 *
	 * 	let { data } = $props();
	 *
	 * 	const increase = async () => {
	 * 		await invalidate('increase:count');
	 * 	}
	 * </script>
	 *
	 * <p>{data.count}<p>
	 * <button on:click={increase}>Increase Count</button>
	 * ```
	 */
	depends: (...deps: Array<`${string}:${string}`>) => void;
	/**
	 * Use this function to opt out of dependency tracking for everything that is synchronously called within the callback. Example:
	 *
	 * ```js
	 * /// file: src/routes/+page.server.js
	 * export async function load({ untrack, url }) {
	 * 	// Untrack url.pathname so that path changes don't trigger a rerun
	 * 	if (untrack(() => url.pathname === '/')) {
	 * 		return { message: 'Welcome!' };
	 * 	}
	 * }
	 * ```
	 */
	untrack: <T>(fn: () => T) => T;
}

export interface NavigationEvent<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	RouteId extends string | null = string | null
> {
	/**
	 * The parameters of the current page - e.g. for a route like `/blog/[slug]`, a `{ slug: string }` object
	 */
	params: Params;
	/**
	 * Info about the current route
	 */
	route: {
		/**
		 * The ID of the current route - e.g. for `src/routes/blog/[slug]`, it would be `/blog/[slug]`. It is `null` when no route is matched.
		 */
		id: RouteId;
	};
	/**
	 * The URL of the current page
	 */
	url: URL;
}

/**
 * Information about the target of a specific navigation.
 */
export interface NavigationTarget {
	/**
	 * Parameters of the target page - e.g. for a route like `/blog/[slug]`, a `{ slug: string }` object.
	 * Is `null` if the target is not part of the SvelteKit app (could not be resolved to a route).
	 */
	params: Record<string, string> | null;
	/**
	 * Info about the target route
	 */
	route: {
		/**
		 * The ID of the current route - e.g. for `src/routes/blog/[slug]`, it would be `/blog/[slug]`. It is `null` when no route is matched.
		 */
		id: string | null;
	};
	/**
	 * The URL that is navigated to
	 */
	url: URL;
}

/**
 * - `enter`: The app has hydrated/started
 * - `form`: The user submitted a `<form>` with a GET method
 * - `leave`: The user is leaving the app by closing the tab or using the back/forward buttons to go to a different document
 * - `link`: Navigation was triggered by a link click
 * - `goto`: Navigation was triggered by a `goto(...)` call or a redirect
 * - `popstate`: Navigation was triggered by back/forward navigation
 */
export type NavigationType = 'enter' | 'form' | 'leave' | 'link' | 'goto' | 'popstate';

export interface Navigation {
	/**
	 * Where navigation was triggered from
	 */
	from: NavigationTarget | null;
	/**
	 * Where navigation is going to/has gone to
	 */
	to: NavigationTarget | null;
	/**
	 * The type of navigation:
	 * - `form`: The user submitted a `<form>`
	 * - `leave`: The app is being left either because the tab is being closed or a navigation to a different document is occurring
	 * - `link`: Navigation was triggered by a link click
	 * - `goto`: Navigation was triggered by a `goto(...)` call or a redirect
	 * - `popstate`: Navigation was triggered by back/forward navigation
	 */
	type: Exclude<NavigationType, 'enter'>;
	/**
	 * Whether or not the navigation will result in the page being unloaded (i.e. not a client-side navigation)
	 */
	willUnload: boolean;
	/**
	 * In case of a history back/forward navigation, the number of steps to go back/forward
	 */
	delta?: number;
	/**
	 * A promise that resolves once the navigation is complete, and rejects if the navigation
	 * fails or is aborted. In the case of a `willUnload` navigation, the promise will never resolve
	 */
	complete: Promise<void>;
}

/**
 * The argument passed to [`beforeNavigate`](https://svelte.dev/docs/kit/$app-navigation#beforeNavigate) callbacks.
 */
export interface BeforeNavigate extends Navigation {
	/**
	 * Call this to prevent the navigation from starting.
	 */
	cancel: () => void;
}

/**
 * The argument passed to [`onNavigate`](https://svelte.dev/docs/kit/$app-navigation#onNavigate) callbacks.
 */
export interface OnNavigate extends Navigation {
	/**
	 * The type of navigation:
	 * - `form`: The user submitted a `<form>`
	 * - `link`: Navigation was triggered by a link click
	 * - `goto`: Navigation was triggered by a `goto(...)` call or a redirect
	 * - `popstate`: Navigation was triggered by back/forward navigation
	 */
	type: Exclude<NavigationType, 'enter' | 'leave'>;
	/**
	 * Since `onNavigate` callbacks are called immediately before a client-side navigation, they will never be called with a navigation that unloads the page.
	 */
	willUnload: false;
}

/**
 * The argument passed to [`afterNavigate`](https://svelte.dev/docs/kit/$app-navigation#afterNavigate) callbacks.
 */
export interface AfterNavigate extends Omit<Navigation, 'type'> {
	/**
	 * The type of navigation:
	 * - `enter`: The app has hydrated/started
	 * - `form`: The user submitted a `<form>`
	 * - `link`: Navigation was triggered by a link click
	 * - `goto`: Navigation was triggered by a `goto(...)` call or a redirect
	 * - `popstate`: Navigation was triggered by back/forward navigation
	 */
	type: Exclude<NavigationType, 'leave'>;
	/**
	 * Since `afterNavigate` callbacks are called after a navigation completes, they will never be called with a navigation that unloads the page.
	 */
	willUnload: false;
}

/**
 * The shape of the [`page`](https://svelte.dev/docs/kit/$app-state#page) reactive object and the [`$page`](https://svelte.dev/docs/kit/$app-stores) store.
 */
export interface Page<
	Params extends Record<string, string> = Record<string, string>,
	RouteId extends string | null = string | null
> {
	/**
	 * The URL of the current page.
	 */
	url: URL;
	/**
	 * The parameters of the current page - e.g. for a route like `/blog/[slug]`, a `{ slug: string }` object.
	 */
	params: Params;
	/**
	 * Info about the current route.
	 */
	route: {
		/**
		 * The ID of the current route - e.g. for `src/routes/blog/[slug]`, it would be `/blog/[slug]`. It is `null` when no route is matched.
		 */
		id: RouteId;
	};
	/**
	 * HTTP status code of the current page.
	 */
	status: number;
	/**
	 * The error object of the current page, if any. Filled from the `handleError` hooks.
	 */
	error: App.Error | null;
	/**
	 * The merged result of all data from all `load` functions on the current page. You can type a common denominator through `App.PageData`.
	 */
	data: App.PageData & Record<string, any>;
	/**
	 * The page state, which can be manipulated using the [`pushState`](https://svelte.dev/docs/kit/$app-navigation#pushState) and [`replaceState`](https://svelte.dev/docs/kit/$app-navigation#replaceState) functions from `$app/navigation`.
	 */
	state: App.PageState;
	/**
	 * Filled only after a form submission. See [form actions](https://svelte.dev/docs/kit/form-actions) for more info.
	 */
	form: any;
}

/**
 * The shape of a param matcher. See [matching](https://svelte.dev/docs/kit/advanced-routing#Matching) for more info.
 */
export type ParamMatcher = (param: string) => boolean;

export interface RequestEvent<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	RouteId extends string | null = string | null
> {
	/**
	 * Get or set cookies related to the current request
	 */
	cookies: Cookies;
	/**
	 * `fetch` is equivalent to the [native `fetch` web API](https://developer.mozilla.org/en-US/docs/Web/API/fetch), with a few additional features:
	 *
	 * - It can be used to make credentialed requests on the server, as it inherits the `cookie` and `authorization` headers for the page request.
	 * - It can make relative requests on the server (ordinarily, `fetch` requires a URL with an origin when used in a server context).
	 * - Internal requests (e.g. for `+server.js` routes) go directly to the handler function when running on the server, without the overhead of an HTTP call.
	 * - During server-side rendering, the response will be captured and inlined into the rendered HTML by hooking into the `text` and `json` methods of the `Response` object. Note that headers will _not_ be serialized, unless explicitly included via [`filterSerializedResponseHeaders`](https://svelte.dev/docs/kit/hooks#Server-hooks-handle)
	 * - During hydration, the response will be read from the HTML, guaranteeing consistency and preventing an additional network request.
	 *
	 * You can learn more about making credentialed requests with cookies [here](https://svelte.dev/docs/kit/load#Cookies).
	 */
	fetch: typeof fetch;
	/**
	 * The client's IP address, set by the adapter.
	 */
	getClientAddress: () => string;
	/**
	 * Contains custom data that was added to the request within the [`server handle hook`](https://svelte.dev/docs/kit/hooks#Server-hooks-handle).
	 */
	locals: App.Locals;
	/**
	 * The parameters of the current route - e.g. for a route like `/blog/[slug]`, a `{ slug: string }` object.
	 */
	params: Params;
	/**
	 * Additional data made available through the adapter.
	 */
	platform: Readonly<App.Platform> | undefined;
	/**
	 * The original request object.
	 */
	request: Request;
	/**
	 * Info about the current route.
	 */
	route: {
		/**
		 * The ID of the current route - e.g. for `src/routes/blog/[slug]`, it would be `/blog/[slug]`. It is `null` when no route is matched.
		 */
		id: RouteId;
	};
	/**
	 * If you need to set headers for the response, you can do so using the this method. This is useful if you want the page to be cached, for example:
	 *
	 *	```js
	 *	/// file: src/routes/blog/+page.js
	 *	export async function load({ fetch, setHeaders }) {
	 *		const url = `https://cms.example.com/articles.json`;
	 *		const response = await fetch(url);
	 *
	 *		setHeaders({
	 *			age: response.headers.get('age'),
	 *			'cache-control': response.headers.get('cache-control')
	 *		});
	 *
	 *		return response.json();
	 *	}
	 *	```
	 *
	 * Setting the same header multiple times (even in separate `load` functions) is an error — you can only set a given header once.
	 *
	 * You cannot add a `set-cookie` header with `setHeaders` — use the [`cookies`](https://svelte.dev/docs/kit/@sveltejs-kit#Cookies) API instead.
	 */
	setHeaders: (headers: Record<string, string>) => void;
	/**
	 * The requested URL.
	 */
	url: URL;
	/**
	 * `true` if the request comes from the client asking for `+page/layout.server.js` data. The `url` property will be stripped of the internal information
	 * related to the data request in this case. Use this property instead if the distinction is important to you.
	 */
	isDataRequest: boolean;
	/**
	 * `true` for `+server.js` calls coming from SvelteKit without the overhead of actually making an HTTP request. This happens when you make same-origin `fetch` requests on the server.
	 */
	isSubRequest: boolean;
}

/**
 * A `(event: RequestEvent) => Response` function exported from a `+server.js` file that corresponds to an HTTP verb (`GET`, `PUT`, `PATCH`, etc) and handles requests with that method.
 *
 * It receives `Params` as the first generic argument, which you can skip by using [generated types](https://svelte.dev/docs/kit/types#Generated-types) instead.
 */
export type RequestHandler<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	RouteId extends string | null = string | null
> = (event: RequestEvent<Params, RouteId>) => MaybePromise<Response>;

export interface ResolveOptions {
	/**
	 * Applies custom transforms to HTML. If `done` is true, it's the final chunk. Chunks are not guaranteed to be well-formed HTML
	 * (they could include an element's opening tag but not its closing tag, for example)
	 * but they will always be split at sensible boundaries such as `%sveltekit.head%` or layout/page components.
	 * @param input the html chunk and the info if this is the last chunk
	 */
	transformPageChunk?: (input: { html: string; done: boolean }) => MaybePromise<string | undefined>;
	/**
	 * Determines which headers should be included in serialized responses when a `load` function loads a resource with `fetch`.
	 * By default, none will be included.
	 * @param name header name
	 * @param value header value
	 */
	filterSerializedResponseHeaders?: (name: string, value: string) => boolean;
	/**
	 * Determines what should be added to the `<head>` tag to preload it.
	 * By default, `js` and `css` files will be preloaded.
	 * @param input the type of the file and its path
	 */
	preload?: (input: { type: 'font' | 'css' | 'js' | 'asset'; path: string }) => boolean;
}

export interface RouteDefinition<Config = any> {
	id: string;
	api: {
		methods: Array<HttpMethod | '*'>;
	};
	page: {
		methods: Array<Extract<HttpMethod, 'GET' | 'POST'>>;
	};
	pattern: RegExp;
	prerender: PrerenderOption;
	segments: RouteSegment[];
	methods: Array<HttpMethod | '*'>;
	config: Config;
}

export class Server {
	constructor(manifest: SSRManifest);
	init(options: ServerInitOptions): Promise<void>;
	respond(request: Request, options: RequestOptions): Promise<Response>;
}

export interface ServerInitOptions {
	/** A map of environment variables. */
	env: Record<string, string>;
	/** A function that turns an asset filename into a `ReadableStream`. Required for the `read` export from `$app/server` to work. */
	read?: (file: string) => ReadableStream;
}

export interface SSRManifest {
	appDir: string;
	appPath: string;
	/** Static files from `kit.config.files.assets` and the service worker (if any). */
	assets: Set<string>;
	mimeTypes: Record<string, string>;

	/** private fields */
	_: {
		client: NonNullable<BuildData['client']>;
		nodes: SSRNodeLoader[];
		routes: SSRRoute[];
		prerendered_routes: Set<string>;
		matchers: () => Promise<Record<string, ParamMatcher>>;
		/** A `[file]: size` map of all assets imported by server code. */
		server_assets: Record<string, number>;
	};
}

/**
 * The generic form of `PageServerLoad` and `LayoutServerLoad`. You should import those from `./$types` (see [generated types](https://svelte.dev/docs/kit/types#Generated-types))
 * rather than using `ServerLoad` directly.
 */
export type ServerLoad<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	ParentData extends Record<string, any> = Record<string, any>,
	OutputData extends Record<string, any> | void = Record<string, any> | void,
	RouteId extends string | null = string | null
> = (event: ServerLoadEvent<Params, ParentData, RouteId>) => MaybePromise<OutputData>;

export interface ServerLoadEvent<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	ParentData extends Record<string, any> = Record<string, any>,
	RouteId extends string | null = string | null
> extends RequestEvent<Params, RouteId> {
	/**
	 * `await parent()` returns data from parent `+layout.server.js` `load` functions.
	 *
	 * Be careful not to introduce accidental waterfalls when using `await parent()`. If for example you only want to merge parent data into the returned output, call it _after_ fetching your other data.
	 */
	parent: () => Promise<ParentData>;
	/**
	 * This function declares that the `load` function has a _dependency_ on one or more URLs or custom identifiers, which can subsequently be used with [`invalidate()`](https://svelte.dev/docs/kit/$app-navigation#invalidate) to cause `load` to rerun.
	 *
	 * Most of the time you won't need this, as `fetch` calls `depends` on your behalf — it's only necessary if you're using a custom API client that bypasses `fetch`.
	 *
	 * URLs can be absolute or relative to the page being loaded, and must be [encoded](https://developer.mozilla.org/en-US/docs/Glossary/percent-encoding).
	 *
	 * Custom identifiers have to be prefixed with one or more lowercase letters followed by a colon to conform to the [URI specification](https://www.rfc-editor.org/rfc/rfc3986.html).
	 *
	 * The following example shows how to use `depends` to register a dependency on a custom identifier, which is `invalidate`d after a button click, making the `load` function rerun.
	 *
	 * ```js
	 * /// file: src/routes/+page.js
	 * let count = 0;
	 * export async function load({ depends }) {
	 * 	depends('increase:count');
	 *
	 * 	return { count: count++ };
	 * }
	 * ```
	 *
	 * ```html
	 * /// file: src/routes/+page.svelte
	 * <script>
	 * 	import { invalidate } from '$app/navigation';
	 *
	 * 	let { data } = $props();
	 *
	 * 	const increase = async () => {
	 * 		await invalidate('increase:count');
	 * 	}
	 * </script>
	 *
	 * <p>{data.count}<p>
	 * <button on:click={increase}>Increase Count</button>
	 * ```
	 */
	depends: (...deps: string[]) => void;
	/**
	 * Use this function to opt out of dependency tracking for everything that is synchronously called within the callback. Example:
	 *
	 * ```js
	 * /// file: src/routes/+page.js
	 * export async function load({ untrack, url }) {
	 * 	// Untrack url.pathname so that path changes don't trigger a rerun
	 * 	if (untrack(() => url.pathname === '/')) {
	 * 		return { message: 'Welcome!' };
	 * 	}
	 * }
	 * ```
	 */
	untrack: <T>(fn: () => T) => T;
}

/**
 * Shape of a form action method that is part of `export const actions = {...}` in `+page.server.js`.
 * See [form actions](https://svelte.dev/docs/kit/form-actions) for more information.
 */
export type Action<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	OutputData extends Record<string, any> | void = Record<string, any> | void,
	RouteId extends string | null = string | null
> = (event: RequestEvent<Params, RouteId>) => MaybePromise<OutputData>;

/**
 * Shape of the `export const actions = {...}` object in `+page.server.js`.
 * See [form actions](https://svelte.dev/docs/kit/form-actions) for more information.
 */
export type Actions<
	Params extends Partial<Record<string, string>> = Partial<Record<string, string>>,
	OutputData extends Record<string, any> | void = Record<string, any> | void,
	RouteId extends string | null = string | null
> = Record<string, Action<Params, OutputData, RouteId>>;

/**
 * When calling a form action via fetch, the response will be one of these shapes.
 * ```svelte
 * <form method="post" use:enhance={() => {
 *   return ({ result }) => {
 * 		// result is of type ActionResult
 *   };
 * }}
 * ```
 */
export type ActionResult<
	Success extends Record<string, unknown> | undefined = Record<string, any>,
	Failure extends Record<string, unknown> | undefined = Record<string, any>
> =
	| { type: 'success'; status: number; data?: Success }
	| { type: 'failure'; status: number; data?: Failure }
	| { type: 'redirect'; status: number; location: string }
	| { type: 'error'; status?: number; error: any };

/**
 * The object returned by the [`error`](https://svelte.dev/docs/kit/@sveltejs-kit#error) function.
 */
export interface HttpError {
	/** The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses), in the range 400-599. */
	status: number;
	/** The content of the error. */
	body: App.Error;
}

/**
 * The object returned by the [`redirect`](https://svelte.dev/docs/kit/@sveltejs-kit#redirect) function.
 */
export interface Redirect {
	/** The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages), in the range 300-308. */
	status: 300 | 301 | 302 | 303 | 304 | 305 | 306 | 307 | 308;
	/** The location to redirect to. */
	location: string;
}

export type SubmitFunction<
	Success extends Record<string, unknown> | undefined = Record<string, any>,
	Failure extends Record<string, unknown> | undefined = Record<string, any>
> = (input: {
	action: URL;
	formData: FormData;
	formElement: HTMLFormElement;
	controller: AbortController;
	submitter: HTMLElement | null;
	cancel: () => void;
}) => MaybePromise<
	| void
	| ((opts: {
			formData: FormData;
			formElement: HTMLFormElement;
			action: URL;
			result: ActionResult<Success, Failure>;
			/**
			 * Call this to get the default behavior of a form submission response.
			 * @param options Set `reset: false` if you don't want the `<form>` values to be reset after a successful submission.
			 * @param invalidateAll Set `invalidateAll: false` if you don't want the action to call `invalidateAll` after submission.
			 */
			update: (options?: { reset?: boolean; invalidateAll?: boolean }) => Promise<void>;
	  }) => MaybePromise<void>)
>;

/**
 * The type of `export const snapshot` exported from a page or layout component.
 */
export interface Snapshot<T = any> {
	capture: () => T;
	restore: (snapshot: T) => void;
}

export * from './index.js';
