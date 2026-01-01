/// <reference types="svelte" />
/// <reference types="vite/client" />

declare module '@sveltejs/kit' {
	import type { CompileOptions } from 'svelte/compiler';
	import type { PluginOptions } from '@sveltejs/vite-plugin-svelte';
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

	const uniqueSymbol: unique symbol;

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
		 * @param directory The directory containing the files to be compressed
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
	interface AdapterEntry {
		/**
		 * A string that uniquely identifies an HTTP service (e.g. serverless function) and is used for deduplication.
		 * For example, `/foo/a-[b]` and `/foo/[c]` are different routes, but would both
		 * be represented in a Netlify _redirects file as `/foo/:param`, so they share an ID
		 */
		id: string;

		/**
		 * A function that compares the candidate route with the current route to determine
		 * if it should be grouped with the current route.
		 *
		 * Use cases:
		 * - Fallback pages: `/foo/[c]` is a fallback for `/foo/a-[b]`, and `/[...catchall]` is a fallback for all routes
		 * - Grouping routes that share a common `config`: `/foo` should be deployed to the edge, `/bar` and `/baz` should be deployed to a serverless function
		 */
		filter(route: RouteDefinition): boolean;

		/**
		 * A function that is invoked once the entry has been created. This is where you
		 * should write the function to the filesystem and generate redirect manifests.
		 */
		complete(entry: { generateManifest(opts: { relativePath: string }): string }): MaybePromise<void>;
	}

	// Based on https://github.com/josh-hemphill/csp-typed-directives/blob/latest/src/csp.types.ts
	//
	// MIT License
	//
	// Copyright (c) 2021-present, Joshua Hemphill
	// Copyright (c) 2021, Tecnico Corporation
	//
	// Permission is hereby granted, free of charge, to any person obtaining a copy
	// of this software and associated documentation files (the "Software"), to deal
	// in the Software without restriction, including without limitation the rights
	// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	// copies of the Software, and to permit persons to whom the Software is
	// furnished to do so, subject to the following conditions:
	//
	// The above copyright notice and this permission notice shall be included in all
	// copies or substantial portions of the Software.
	//
	// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	// SOFTWARE.

	namespace Csp {
		type ActionSource = 'strict-dynamic' | 'report-sample';
		type BaseSource =
			| 'self'
			| 'unsafe-eval'
			| 'unsafe-hashes'
			| 'unsafe-inline'
			| 'wasm-unsafe-eval'
			| 'none';
		type CryptoSource = `${'nonce' | 'sha256' | 'sha384' | 'sha512'}-${string}`;
		type FrameSource = HostSource | SchemeSource | 'self' | 'none';
		type HostNameScheme = `${string}.${string}` | 'localhost';
		type HostSource = `${HostProtocolSchemes}${HostNameScheme}${PortScheme}`;
		type HostProtocolSchemes = `${string}://` | '';
		type HttpDelineator = '/' | '?' | '#' | '\\';
		type PortScheme = `:${number}` | '' | ':*';
		type SchemeSource = 'http:' | 'https:' | 'data:' | 'mediastream:' | 'blob:' | 'filesystem:';
		type Source = HostSource | SchemeSource | CryptoSource | BaseSource;
		type Sources = Source[];
	}

	interface CspDirectives {
		'child-src'?: Csp.Sources;
		'default-src'?: Array<Csp.Source | Csp.ActionSource>;
		'frame-src'?: Csp.Sources;
		'worker-src'?: Csp.Sources;
		'connect-src'?: Csp.Sources;
		'font-src'?: Csp.Sources;
		'img-src'?: Csp.Sources;
		'manifest-src'?: Csp.Sources;
		'media-src'?: Csp.Sources;
		'object-src'?: Csp.Sources;
		'prefetch-src'?: Csp.Sources;
		'script-src'?: Array<Csp.Source | Csp.ActionSource>;
		'script-src-elem'?: Csp.Sources;
		'script-src-attr'?: Csp.Sources;
		'style-src'?: Array<Csp.Source | Csp.ActionSource>;
		'style-src-elem'?: Csp.Sources;
		'style-src-attr'?: Csp.Sources;
		'base-uri'?: Array<Csp.Source | Csp.ActionSource>;
		sandbox?: Array<
			| 'allow-downloads-without-user-activation'
			| 'allow-forms'
			| 'allow-modals'
			| 'allow-orientation-lock'
			| 'allow-pointer-lock'
			| 'allow-popups'
			| 'allow-popups-to-escape-sandbox'
			| 'allow-presentation'
			| 'allow-same-origin'
			| 'allow-scripts'
			| 'allow-storage-access-by-user-activation'
			| 'allow-top-navigation'
			| 'allow-top-navigation-by-user-activation'
		>;
		'form-action'?: Array<Csp.Source | Csp.ActionSource>;
		'frame-ancestors'?: Array<Csp.HostSource | Csp.SchemeSource | Csp.FrameSource>;
		'navigate-to'?: Array<Csp.Source | Csp.ActionSource>;
		'report-uri'?: string[];
		'report-to'?: string[];

		'require-trusted-types-for'?: Array<'script'>;
		'trusted-types'?: Array<'none' | 'allow-duplicates' | '*' | string>;
		'upgrade-insecure-requests'?: boolean;

		/** @deprecated */
		'require-sri-for'?: Array<'script' | 'style' | 'script style'>;

		/** @deprecated */
		'block-all-mixed-content'?: boolean;

		/** @deprecated */
		'plugin-types'?: Array<`${string}/${string}` | 'none'>;

		/** @deprecated */
		referrer?: Array<
			| 'no-referrer'
			| 'no-referrer-when-downgrade'
			| 'origin'
			| 'origin-when-cross-origin'
			| 'same-origin'
			| 'strict-origin'
			| 'strict-origin-when-cross-origin'
			| 'unsafe-url'
			| 'none'
		>;
	}

	type HttpMethod = 'GET' | 'HEAD' | 'POST' | 'PUT' | 'DELETE' | 'PATCH' | 'OPTIONS';

	interface Logger {
		(msg: string): void;
		success(msg: string): void;
		error(msg: string): void;
		warn(msg: string): void;
		minor(msg: string): void;
		info(msg: string): void;
	}

	type MaybePromise<T> = T | Promise<T>;

	interface Prerendered {
		/**
		 * A map of `path` to `{ file }` objects, where a path like `/foo` corresponds to `foo.html` and a path like `/bar/` corresponds to `bar/index.html`.
		 */
		pages: Map<
			string,
			{
				/** The location of the .html file relative to the output directory */
				file: string;
			}
		>;
		/**
		 * A map of `path` to `{ type }` objects.
		 */
		assets: Map<
			string,
			{
				/** The MIME type of the asset */
				type: string;
			}
		>;
		/**
		 * A map of redirects encountered during prerendering.
		 */
		redirects: Map<
			string,
			{
				status: number;
				location: string;
			}
		>;
		/** An array of prerendered paths (without trailing slashes, regardless of the trailingSlash config) */
		paths: string[];
	}

	interface PrerenderHttpErrorHandler {
		(details: {
			status: number;
			path: string;
			referrer: string | null;
			referenceType: 'linked' | 'fetched';
			message: string;
		}): void;
	}

	interface PrerenderMissingIdHandler {
		(details: { path: string; id: string; referrers: string[]; message: string }): void;
	}

	interface PrerenderEntryGeneratorMismatchHandler {
		(details: { generatedFromId: string; entry: string; matchedId: string; message: string }): void;
	}

	type PrerenderHttpErrorHandlerValue = 'fail' | 'warn' | 'ignore' | PrerenderHttpErrorHandler;
	type PrerenderMissingIdHandlerValue = 'fail' | 'warn' | 'ignore' | PrerenderMissingIdHandler;
	type PrerenderEntryGeneratorMismatchHandlerValue =
		| 'fail'
		| 'warn'
		| 'ignore'
		| PrerenderEntryGeneratorMismatchHandler;

	export type PrerenderOption = boolean | 'auto';

	interface RequestOptions {
		getClientAddress(): string;
		platform?: App.Platform;
	}

	interface RouteSegment {
		content: string;
		dynamic: boolean;
		rest: boolean;
	}

	type TrailingSlash = 'never' | 'always' | 'ignore';
	interface Asset {
		file: string;
		size: number;
		type: string | null;
	}

	interface BuildData {
		app_dir: string;
		app_path: string;
		manifest_data: ManifestData;
		out_dir: string;
		service_worker: string | null;
		client: {
			/** Path to the client entry point. */
			start: string;
			/** Path to the generated `app.js` file that contains the client manifest. Only set in case of `bundleStrategy === 'split'`. */
			app?: string;
			/** JS files that the client entry point relies on. */
			imports: string[];
			/**
			 * JS files that represent the entry points of the layouts/pages.
			 * An entry is undefined if the layout/page has no component or universal file (i.e. only has a `.server.js` file).
			 * Only set in case of `router.resolution === 'server'`.
			 */
			nodes?: Array<string | undefined>;
			/**
			 * CSS files referenced in the entry points of the layouts/pages.
			 * An entry is undefined if the layout/page has no component or universal file (i.e. only has a `.server.js` file) or if has no CSS.
			 * Only set in case of `router.resolution === 'server'`.
			 */
			css?: Array<string[] | undefined>;
			/**
			 * Contains the client route manifest in a form suitable for the server which is used for server side route resolution.
			 * Notably, it contains all routes, regardless of whether they are prerendered or not (those are missing in the optimized server route manifest).
			 * Only set in case of `router.resolution === 'server'`.
			 */
			routes?: SSRClientRoute[];
			stylesheets: string[];
			fonts: string[];
			uses_env_dynamic_public: boolean;
			/** Only set in case of `bundleStrategy === 'inline'`. */
			inline?: {
				script: string;
				style: string | undefined;
			};
		} | null;
		server_manifest: import('vite').Manifest;
	}

	interface ManifestData {
		/** Static files from `kit.config.files.assets`. */
		assets: Asset[];
		hooks: {
			client: string | null;
			server: string | null;
			universal: string | null;
		};
		nodes: PageNode[];
		routes: RouteData[];
		matchers: Record<string, string>;
	}

	interface PageNode {
		depth: number;
		/** The `+page/layout.svelte`. */
		component?: string; // TODO supply default component if it's missing (bit of an edge case)
		/** The `+page/layout.js/.ts`. */
		universal?: string;
		/** The `+page/layout.server.js/ts`. */
		server?: string;
		parent_id?: string;
		parent?: PageNode;
		/** Filled with the pages that reference this layout (if this is a layout). */
		child_pages?: PageNode[];
	}

	type RecursiveRequired<T> = {
		// Recursive implementation of TypeScript's Required utility type.
		// Will recursively continue until it reaches a primitive or Function
		// eslint-disable-next-line @typescript-eslint/no-unsafe-function-type
		[K in keyof T]-?: Extract<T[K], Function> extends never // If it does not have a Function type
			? RecursiveRequired<T[K]> // recursively continue through.
			: T[K]; // Use the exact type for everything else
	};

	interface RouteParam {
		name: string;
		matcher: string;
		optional: boolean;
		rest: boolean;
		chained: boolean;
	}

	/**
	 * Represents a route segment in the app. It can either be an intermediate node
	 * with only layout/error pages, or a leaf, at which point either `page` and `leaf`
	 * or `endpoint` is set.
	 */
	interface RouteData {
		id: string;
		parent: RouteData | null;

		segment: string;
		pattern: RegExp;
		params: RouteParam[];

		layout: PageNode | null;
		error: PageNode | null;
		leaf: PageNode | null;

		page: {
			layouts: Array<number | undefined>;
			errors: Array<number | undefined>;
			leaf: number;
		} | null;

		endpoint: {
			file: string;
		} | null;
	}

	interface SSRComponent {
		default: {
			render(
				props: Record<string, any>,
				opts: { context: Map<any, any> }
			): {
				html: string;
				head: string;
				css: {
					code: string;
					map: any; // TODO
				};
			};
		};
	}

	type SSRComponentLoader = () => Promise<SSRComponent>;

	interface UniversalNode {
		load?: Load;
		prerender?: PrerenderOption;
		ssr?: boolean;
		csr?: boolean;
		trailingSlash?: TrailingSlash;
		config?: any;
		entries?: PrerenderEntryGenerator;
	}

	interface ServerNode {
		load?: ServerLoad;
		prerender?: PrerenderOption;
		ssr?: boolean;
		csr?: boolean;
		trailingSlash?: TrailingSlash;
		actions?: Actions;
		config?: any;
		entries?: PrerenderEntryGenerator;
	}

	interface SSRNode {
		/** index into the `nodes` array in the generated `client/app.js`. */
		index: number;
		/** external JS files that are loaded on the client. `imports[0]` is the entry point (e.g. `client/nodes/0.js`) */
		imports: string[];
		/** external CSS files that are loaded on the client */
		stylesheets: string[];
		/** external font files that are loaded on the client */
		fonts: string[];

		universal_id?: string;
		server_id?: string;

		/** inlined styles */
		inline_styles?(): MaybePromise<Record<string, string>>;
		/** Svelte component */
		component?: SSRComponentLoader;
		/** +page.js or +layout.js */
		universal?: UniversalNode;
		/** +page.server.js, +layout.server.js, or +server.js */
		server?: ServerNode;
	}

	type SSRNodeLoader = () => Promise<SSRNode>;

	interface PageNodeIndexes {
		errors: Array<number | undefined>;
		layouts: Array<number | undefined>;
		leaf: number;
	}

	type PrerenderEntryGenerator = () => MaybePromise<Array<Record<string, string>>>;

	type SSREndpoint = Partial<Record<HttpMethod, RequestHandler>> & {
		prerender?: PrerenderOption;
		trailingSlash?: TrailingSlash;
		config?: any;
		entries?: PrerenderEntryGenerator;
		fallback?: RequestHandler;
	};

	interface SSRRoute {
		id: string;
		pattern: RegExp;
		params: RouteParam[];
		page: PageNodeIndexes | null;
		endpoint: (() => Promise<SSREndpoint>) | null;
		endpoint_id?: string;
	}

	interface SSRClientRoute {
		id: string;
		pattern: RegExp;
		params: RouteParam[];
		errors: Array<number | undefined>;
		layouts: Array<[has_server_load: boolean, node_id: number] | undefined>;
		leaf: [has_server_load: boolean, node_id: number];
	}

	type ValidatedConfig = Config & {
		kit: ValidatedKitConfig;
		extensions: string[];
	};

	type ValidatedKitConfig = Omit<RecursiveRequired<KitConfig>, 'adapter'> & {
		adapter?: Adapter;
	};
	/**
	 * Throws an error with a HTTP status code and an optional message.
	 * When called during request handling, this will cause SvelteKit to
	 * return an error response without invoking `handleError`.
	 * Make sure you're not catching the thrown error, which would prevent SvelteKit from handling it.
	 * @param status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
	 * @param body An object that conforms to the App.Error type. If a string is passed, it will be used as the message property.
	 * @throws {HttpError} This error instructs SvelteKit to initiate HTTP error handling.
	 * @throws {Error} If the provided status is invalid (not between 400 and 599).
	 */
	export function error(status: number, body: App.Error): never;
	/**
	 * Throws an error with a HTTP status code and an optional message.
	 * When called during request handling, this will cause SvelteKit to
	 * return an error response without invoking `handleError`.
	 * Make sure you're not catching the thrown error, which would prevent SvelteKit from handling it.
	 * @param status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
	 * @param body An object that conforms to the App.Error type. If a string is passed, it will be used as the message property.
	 * @throws {HttpError} This error instructs SvelteKit to initiate HTTP error handling.
	 * @throws {Error} If the provided status is invalid (not between 400 and 599).
	 */
	export function error(status: number, body?: {
		message: string;
	} extends App.Error ? App.Error | string | undefined : never): never;
	/**
	 * Checks whether this is an error thrown by {@link error}.
	 * @param status The status to filter for.
	 * */
	export function isHttpError<T extends number>(e: unknown, status?: T | undefined): e is (HttpError_1 & {
		status: T extends undefined ? never : T;
	});
	/**
	 * Redirect a request. When called during request handling, SvelteKit will return a redirect response.
	 * Make sure you're not catching the thrown redirect, which would prevent SvelteKit from handling it.
	 *
	 * Most common status codes:
	 *  * `303 See Other`: redirect as a GET request (often used after a form POST request)
	 *  * `307 Temporary Redirect`: redirect will keep the request method
	 *  * `308 Permanent Redirect`: redirect will keep the request method, SEO will be transferred to the new page
	 *
	 * [See all redirect status codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages)
	 *
	 * @param status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages). Must be in the range 300-308.
	 * @param location The location to redirect to.
	 * @throws {Redirect} This error instructs SvelteKit to redirect to the specified location.
	 * @throws {Error} If the provided status is invalid.
	 * */
	export function redirect(status: 300 | 301 | 302 | 303 | 304 | 305 | 306 | 307 | 308 | ({} & number), location: string | URL): never;
	/**
	 * Checks whether this is a redirect thrown by {@link redirect}.
	 * @param e The object to check.
	 * */
	export function isRedirect(e: unknown): e is Redirect_1;
	/**
	 * Create a JSON `Response` object from the supplied data.
	 * @param data The value that will be serialized as JSON.
	 * @param init Options such as `status` and `headers` that will be added to the response. `Content-Type: application/json` and `Content-Length` headers will be added automatically.
	 */
	export function json(data: any, init?: ResponseInit | undefined): Response;
	/**
	 * Create a `Response` object from the supplied body.
	 * @param body The value that will be used as-is.
	 * @param init Options such as `status` and `headers` that will be added to the response. A `Content-Length` header will be added automatically.
	 */
	export function text(body: string, init?: ResponseInit | undefined): Response;
	/**
	 * Create an `ActionFailure` object. Call when form submission fails.
	 * @param status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
	 * */
	export function fail(status: number): ActionFailure<undefined>;
	/**
	 * Create an `ActionFailure` object. Call when form submission fails.
	 * @param status The [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses). Must be in the range 400-599.
	 * @param data Data associated with the failure (e.g. validation errors)
	 * */
	export function fail<T extends Record<string, unknown> | undefined = undefined>(status: number, data: T): ActionFailure<T>;
	/**
	 * Checks whether this is an action failure thrown by {@link fail}.
	 * @param e The object to check.
	 * */
	export function isActionFailure(e: unknown): e is ActionFailure;
	/**
	 * Strips possible SvelteKit-internal suffixes and trailing slashes from the URL pathname.
	 * Returns the normalized URL as well as a method for adding the potential suffix back
	 * based on a new pathname (possibly including search) or URL.
	 * ```js
	 * import { normalizeUrl } from '@sveltejs/kit';
	 *
	 * const { url, denormalize } = normalizeUrl('/blog/post/__data.json');
	 * console.log(url.pathname); // /blog/post
	 * console.log(denormalize('/blog/post/a')); // /blog/post/a/__data.json
	 * ```
	 * @since 2.18.0
	 */
	export function normalizeUrl(url: URL | string): {
		url: URL;
		wasNormalized: boolean;
		denormalize: (url?: string | URL) => URL;
	};
	export type LessThan<TNumber extends number, TArray extends any[] = []> = TNumber extends TArray["length"] ? TArray[number] : LessThan<TNumber, [...TArray, TArray["length"]]>;
	export type NumericRange<TStart extends number, TEnd extends number> = Exclude<TEnd | LessThan<TEnd>, LessThan<TStart>>;
	export const VERSION: string;
	class HttpError_1 {
		
		constructor(status: number, body: {
			message: string;
		} extends App.Error ? (App.Error | string | undefined) : App.Error);
		status: number;
		body: App.Error;
		toString(): string;
	}
	class Redirect_1 {
		
		constructor(status: 300 | 301 | 302 | 303 | 304 | 305 | 306 | 307 | 308, location: string);
		status: 301 | 302 | 303 | 307 | 308 | 300 | 304 | 305 | 306;
		location: string;
	}

	export {};
}

declare module '@sveltejs/kit/hooks' {
	/**
	 * A helper function for sequencing multiple `handle` calls in a middleware-like manner.
	 * The behavior for the `handle` options is as follows:
	 * - `transformPageChunk` is applied in reverse order and merged
	 * - `preload` is applied in forward order, the first option "wins" and no `preload` options after it are called
	 * - `filterSerializedResponseHeaders` behaves the same as `preload`
	 *
	 * ```js
	 * /// file: src/hooks.server.js
	 * import { sequence } from '@sveltejs/kit/hooks';
	 *
	 * /// type: import('@sveltejs/kit').Handle
	 * async function first({ event, resolve }) {
	 * 	console.log('first pre-processing');
	 * 	const result = await resolve(event, {
	 * 		transformPageChunk: ({ html }) => {
	 * 			// transforms are applied in reverse order
	 * 			console.log('first transform');
	 * 			return html;
	 * 		},
	 * 		preload: () => {
	 * 			// this one wins as it's the first defined in the chain
	 * 			console.log('first preload');
	 * 			return true;
	 * 		}
	 * 	});
	 * 	console.log('first post-processing');
	 * 	return result;
	 * }
	 *
	 * /// type: import('@sveltejs/kit').Handle
	 * async function second({ event, resolve }) {
	 * 	console.log('second pre-processing');
	 * 	const result = await resolve(event, {
	 * 		transformPageChunk: ({ html }) => {
	 * 			console.log('second transform');
	 * 			return html;
	 * 		},
	 * 		preload: () => {
	 * 			console.log('second preload');
	 * 			return true;
	 * 		},
	 * 		filterSerializedResponseHeaders: () => {
	 * 			// this one wins as it's the first defined in the chain
	 * 			console.log('second filterSerializedResponseHeaders');
	 * 			return true;
	 * 		}
	 * 	});
	 * 	console.log('second post-processing');
	 * 	return result;
	 * }
	 *
	 * export const handle = sequence(first, second);
	 * ```
	 *
	 * The example above would print:
	 *
	 * ```
	 * first pre-processing
	 * first preload
	 * second pre-processing
	 * second filterSerializedResponseHeaders
	 * second transform
	 * first transform
	 * second post-processing
	 * first post-processing
	 * ```
	 *
	 * @param handlers The chain of `handle` functions
	 * */
	export function sequence(...handlers: import("@sveltejs/kit").Handle[]): import("@sveltejs/kit").Handle;

	export {};
}

declare module '@sveltejs/kit/node' {
	export function getRequest({ request, base, bodySizeLimit }: {
		request: import("http").IncomingMessage;
		base: string;
		bodySizeLimit?: number;
	}): Promise<Request>;

	export function setResponse(res: import("http").ServerResponse, response: Response): Promise<void>;
	/**
	 * Converts a file on disk to a readable stream
	 * @since 2.4.0
	 */
	export function createReadableStream(file: string): ReadableStream;

	export {};
}

declare module '@sveltejs/kit/node/polyfills' {
	/**
	 * Make various web APIs available as globals:
	 * - `crypto`
	 * - `File`
	 */
	export function installPolyfills(): void;

	export {};
}

declare module '@sveltejs/kit/vite' {
	/**
	 * Returns the SvelteKit Vite plugins.
	 * */
	export function sveltekit(): Promise<import("vite").Plugin[]>;

	export {};
}

declare module '$app/environment' {
	/**
	 * `true` if the app is running in the browser.
	 */
	export const browser: boolean;

	/**
	 * Whether the dev server is running. This is not guaranteed to correspond to `NODE_ENV` or `MODE`.
	 */
	export const dev: boolean;

	/**
	 * SvelteKit analyses your app during the `build` step by running it. During this process, `building` is `true`. This also applies during prerendering.
	 */
	export const building: boolean;

	/**
	 * The value of `config.kit.version.name`.
	 */
	export const version: string;

	export {};
}

declare module '$app/forms' {
	/**
	 * Use this function to deserialize the response from a form submission.
	 * Usage:
	 *
	 * ```js
	 * import { deserialize } from '$app/forms';
	 *
	 * async function handleSubmit(event) {
	 *   const response = await fetch('/form?/action', {
	 *     method: 'POST',
	 *     body: new FormData(event.target)
	 *   });
	 *
	 *   const result = deserialize(await response.text());
	 *   // ...
	 * }
	 * ```
	 * */
	export function deserialize<Success extends Record<string, unknown> | undefined, Failure extends Record<string, unknown> | undefined>(result: string): import("@sveltejs/kit").ActionResult<Success, Failure>;
	/**
	 * This action enhances a `<form>` element that otherwise would work without JavaScript.
	 *
	 * The `submit` function is called upon submission with the given FormData and the `action` that should be triggered.
	 * If `cancel` is called, the form will not be submitted.
	 * You can use the abort `controller` to cancel the submission in case another one starts.
	 * If a function is returned, that function is called with the response from the server.
	 * If nothing is returned, the fallback will be used.
	 *
	 * If this function or its return value isn't set, it
	 * - falls back to updating the `form` prop with the returned data if the action is on the same page as the form
	 * - updates `page.status`
	 * - resets the `<form>` element and invalidates all data in case of successful submission with no redirect response
	 * - redirects in case of a redirect response
	 * - redirects to the nearest error page in case of an unexpected error
	 *
	 * If you provide a custom function with a callback and want to use the default behavior, invoke `update` in your callback.
	 * It accepts an options object
	 * - `reset: false` if you don't want the `<form>` values to be reset after a successful submission
	 * - `invalidateAll: false` if you don't want the action to call `invalidateAll` after submission
	 * @param form_element The form element
	 * @param submit Submit callback
	 */
	export function enhance<Success extends Record<string, unknown> | undefined, Failure extends Record<string, unknown> | undefined>(form_element: HTMLFormElement, submit?: import("@sveltejs/kit").SubmitFunction<Success, Failure>): {
		destroy(): void;
	};
	/**
	 * This action updates the `form` property of the current page with the given data and updates `page.status`.
	 * In case of an error, it redirects to the nearest error page.
	 * */
	export function applyAction<Success extends Record<string, unknown> | undefined, Failure extends Record<string, unknown> | undefined>(result: import("@sveltejs/kit").ActionResult<Success, Failure>): Promise<void>;

	export {};
}

declare module '$app/navigation' {
	/**
	 * A lifecycle function that runs the supplied `callback` when the current component mounts, and also whenever we navigate to a URL.
	 *
	 * `afterNavigate` must be called during a component initialization. It remains active as long as the component is mounted.
	 * */
	export function afterNavigate(callback: (navigation: import("@sveltejs/kit").AfterNavigate) => void): void;
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
	 * */
	export function beforeNavigate(callback: (navigation: import("@sveltejs/kit").BeforeNavigate) => void): void;
	/**
	 * A lifecycle function that runs the supplied `callback` immediately before we navigate to a new URL except during full-page navigations.
	 *
	 * If you return a `Promise`, SvelteKit will wait for it to resolve before completing the navigation. This allows you to — for example — use `document.startViewTransition`. Avoid promises that are slow to resolve, since navigation will appear stalled to the user.
	 *
	 * If a function (or a `Promise` that resolves to a function) is returned from the callback, it will be called once the DOM has updated.
	 *
	 * `onNavigate` must be called during a component initialization. It remains active as long as the component is mounted.
	 * */
	export function onNavigate(callback: (navigation: import("@sveltejs/kit").OnNavigate) => MaybePromise<(() => void) | void>): void;
	/**
	 * If called when the page is being updated following a navigation (in `onMount` or `afterNavigate` or an action, for example), this disables SvelteKit's built-in scroll handling.
	 * This is generally discouraged, since it breaks user expectations.
	 * */
	export function disableScrollHandling(): void;
	/**
	 * Allows you to navigate programmatically to a given route, with options such as keeping the current element focused.
	 * Returns a Promise that resolves when SvelteKit navigates (or fails to navigate, in which case the promise rejects) to the specified `url`.
	 *
	 * For external URLs, use `window.location = url` instead of calling `goto(url)`.
	 *
	 * @param url Where to navigate to. Note that if you've set [`config.kit.paths.base`](https://svelte.dev/docs/kit/configuration#paths) and the URL is root-relative, you need to prepend the base path if you want to navigate within the app.
	 * @param {Object} opts Options related to the navigation
	 * */
	export function goto(url: string | URL, opts?: {
		replaceState?: boolean | undefined;
		noScroll?: boolean | undefined;
		keepFocus?: boolean | undefined;
		invalidateAll?: boolean | undefined;
		invalidate?: (string | URL | ((url: URL) => boolean))[] | undefined;
		state?: App.PageState | undefined;
	} | undefined): Promise<void>;
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
	 * @param resource The invalidated URL
	 * */
	export function invalidate(resource: string | URL | ((url: URL) => boolean)): Promise<void>;
	/**
	 * Causes all `load` functions belonging to the currently active page to re-run. Returns a `Promise` that resolves when the page is subsequently updated.
	 * */
	export function invalidateAll(): Promise<void>;
	/**
	 * Programmatically preloads the given page, which means
	 *  1. ensuring that the code for the page is loaded, and
	 *  2. calling the page's load function with the appropriate options.
	 *
	 * This is the same behaviour that SvelteKit triggers when the user taps or mouses over an `<a>` element with `data-sveltekit-preload-data`.
	 * If the next navigation is to `href`, the values returned from load will be used, making navigation instantaneous.
	 * Returns a Promise that resolves with the result of running the new route's `load` functions once the preload is complete.
	 *
	 * @param href Page to preload
	 * */
	export function preloadData(href: string): Promise<{
		type: "loaded";
		status: number;
		data: Record<string, any>;
	} | {
		type: "redirect";
		location: string;
	}>;
	/**
	 * Programmatically imports the code for routes that haven't yet been fetched.
	 * Typically, you might call this to speed up subsequent navigation.
	 *
	 * You can specify routes by any matching pathname such as `/about` (to match `src/routes/about/+page.svelte`) or `/blog/*` (to match `src/routes/blog/[slug]/+page.svelte`).
	 *
	 * Unlike `preloadData`, this won't call `load` functions.
	 * Returns a Promise that resolves when the modules have been imported.
	 *
	 * */
	export function preloadCode(pathname: string): Promise<void>;
	/**
	 * Programmatically create a new history entry with the given `page.state`. To use the current URL, you can pass `''` as the first argument. Used for [shallow routing](https://svelte.dev/docs/kit/shallow-routing).
	 *
	 * */
	export function pushState(url: string | URL, state: App.PageState): void;
	/**
	 * Programmatically replace the current history entry with the given `page.state`. To use the current URL, you can pass `''` as the first argument. Used for [shallow routing](https://svelte.dev/docs/kit/shallow-routing).
	 *
	 * */
	export function replaceState(url: string | URL, state: App.PageState): void;
	type MaybePromise<T> = T | Promise<T>;

	export {};
}

declare module '$app/paths' {
	/**
	 * A string that matches [`config.kit.paths.base`](https://svelte.dev/docs/kit/configuration#paths).
	 *
	 * Example usage: `<a href="{base}/your-page">Link</a>`
	 */
	export let base: '' | `/${string}`;

	/**
	 * An absolute path that matches [`config.kit.paths.assets`](https://svelte.dev/docs/kit/configuration#paths).
	 *
	 * > [!NOTE] If a value for `config.kit.paths.assets` is specified, it will be replaced with `'/_svelte_kit_assets'` during `vite dev` or `vite preview`, since the assets don't yet live at their eventual URL.
	 */
	export let assets: '' | `https://${string}` | `http://${string}` | '/_svelte_kit_assets';

	/**
	 * Populate a route ID with params to resolve a pathname.
	 * @example
	 * ```js
	 * import { resolveRoute } from '$app/paths';
	 *
	 * resolveRoute(
	 *   `/blog/[slug]/[...somethingElse]`,
	 *   {
	 *     slug: 'hello-world',
	 *     somethingElse: 'something/else'
	 *   }
	 * ); // `/blog/hello-world/something/else`
	 * ```
	 */
	export function resolveRoute(id: string, params: Record<string, string | undefined>): string;

	export {};
}

declare module '$app/server' {
	import type { RequestEvent } from '@sveltejs/kit';
	/**
	 * Read the contents of an imported asset from the filesystem
	 * @example
	 * ```js
	 * import { read } from '$app/server';
	 * import somefile from './somefile.txt';
	 *
	 * const asset = read(somefile);
	 * const text = await asset.text();
	 * ```
	 * @since 2.4.0
	 */
	export function read(asset: string): Response;
	/**
	 * Returns the current `RequestEvent`. Can be used inside server hooks, server `load` functions, actions, and endpoints (and functions called by them).
	 *
	 * In environments without [`AsyncLocalStorage`](https://nodejs.org/api/async_context.html#class-asynclocalstorage), this must be called synchronously (i.e. not after an `await`).
	 * @since 2.20.0
	 */
	export function getRequestEvent(): RequestEvent<Partial<Record<string, string>>, string | null>;

	export {};
}

declare module '$app/state' {
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
	 * */
	export const page: import("@sveltejs/kit").Page;
	/**
	 * A read-only object representing an in-progress navigation, with `from`, `to`, `type` and (if `type === 'popstate'`) `delta` properties.
	 * Values are `null` when no navigation is occurring, or during server rendering.
	 * */
	export const navigating: import("@sveltejs/kit").Navigation | {
		from: null;
		to: null;
		type: null;
		willUnload: null;
		delta: null;
		complete: null;
	};
	/**
	 * A read-only reactive value that's initially `false`. If [`version.pollInterval`](https://svelte.dev/docs/kit/configuration#version) is a non-zero value, SvelteKit will poll for new versions of the app and update `current` to `true` when it detects one. `updated.check()` will force an immediate check, regardless of polling.
	 * */
	export const updated: {
		get current(): boolean;
		check(): Promise<boolean>;
	};

	export {};
}

declare module '$app/stores' {
	export function getStores(): {
		
		page: typeof page;
		
		navigating: typeof navigating;
		
		updated: typeof updated;
	};
	/**
	 * A readable store whose value contains page data.
	 *
	 * On the server, this store can only be subscribed to during component initialization. In the browser, it can be subscribed to at any time.
	 *
	 * @deprecated Use `page` from `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
	 * */
	export const page: import("svelte/store").Readable<import("@sveltejs/kit").Page>;
	/**
	 * A readable store.
	 * When navigating starts, its value is a `Navigation` object with `from`, `to`, `type` and (if `type === 'popstate'`) `delta` properties.
	 * When navigating finishes, its value reverts to `null`.
	 *
	 * On the server, this store can only be subscribed to during component initialization. In the browser, it can be subscribed to at any time.
	 *
	 * @deprecated Use `navigating` from `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
	 * */
	export const navigating: import("svelte/store").Readable<import("@sveltejs/kit").Navigation | null>;
	/**
	 * A readable store whose initial value is `false`. If [`version.pollInterval`](https://svelte.dev/docs/kit/configuration#version) is a non-zero value, SvelteKit will poll for new versions of the app and update the store value to `true` when it detects one. `updated.check()` will force an immediate check, regardless of polling.
	 *
	 * On the server, this store can only be subscribed to during component initialization. In the browser, it can be subscribed to at any time.
	 *
	 * @deprecated Use `updated` from `$app/state` instead (requires Svelte 5, [see docs for more info](https://svelte.dev/docs/kit/migrating-to-sveltekit-2#SvelteKit-2.12:-$app-stores-deprecated))
	 * */
	export const updated: import("svelte/store").Readable<boolean> & {
		check(): Promise<boolean>;
	};

	export {};
}/**
 * It's possible to tell SvelteKit how to type objects inside your app by declaring the `App` namespace. By default, a new project will have a file called `src/app.d.ts` containing the following:
 *
 * ```ts
 * declare global {
 * 	namespace App {
 * 		// interface Error {}
 * 		// interface Locals {}
 * 		// interface PageData {}
 * 		// interface PageState {}
 * 		// interface Platform {}
 * 	}
 * }
 *
 * export {};
 * ```
 *
 * The `export {}` line exists because without it, the file would be treated as an _ambient module_ which prevents you from adding `import` declarations.
 * If you need to add ambient `declare module` declarations, do so in a separate file like `src/ambient.d.ts`.
 *
 * By populating these interfaces, you will gain type safety when using `event.locals`, `event.platform`, and `data` from `load` functions.
 */
declare namespace App {
	/**
	 * Defines the common shape of expected and unexpected errors. Expected errors are thrown using the `error` function. Unexpected errors are handled by the `handleError` hooks which should return this shape.
	 */
	export interface Error {
		message: string;
	}

	/**
	 * The interface that defines `event.locals`, which can be accessed in server [hooks](https://svelte.dev/docs/kit/hooks) (`handle`, and `handleError`), server-only `load` functions, and `+server.js` files.
	 */
	export interface Locals {}

	/**
	 * Defines the common shape of the [page.data state](https://svelte.dev/docs/kit/$app-state#page) and [$page.data store](https://svelte.dev/docs/kit/$app-stores#page) - that is, the data that is shared between all pages.
	 * The `Load` and `ServerLoad` functions in `./$types` will be narrowed accordingly.
	 * Use optional properties for data that is only present on specific pages. Do not add an index signature (`[key: string]: any`).
	 */
	export interface PageData {}

	/**
	 * The shape of the `page.state` object, which can be manipulated using the [`pushState`](https://svelte.dev/docs/kit/$app-navigation#pushState) and [`replaceState`](https://svelte.dev/docs/kit/$app-navigation#replaceState) functions from `$app/navigation`.
	 */
	export interface PageState {}

	/**
	 * If your adapter provides [platform-specific context](https://svelte.dev/docs/kit/adapters#Platform-specific-context) via `event.platform`, you can specify it here.
	 */
	export interface Platform {}
}

/**
 * This module is only available to [service workers](https://svelte.dev/docs/kit/service-workers).
 */
declare module '$service-worker' {
	/**
	 * The `base` path of the deployment. Typically this is equivalent to `config.kit.paths.base`, but it is calculated from `location.pathname` meaning that it will continue to work correctly if the site is deployed to a subdirectory.
	 * Note that there is a `base` but no `assets`, since service workers cannot be used if `config.kit.paths.assets` is specified.
	 */
	export const base: string;
	/**
	 * An array of URL strings representing the files generated by Vite, suitable for caching with `cache.addAll(build)`.
	 * During development, this is an empty array.
	 */
	export const build: string[];
	/**
	 * An array of URL strings representing the files in your static directory, or whatever directory is specified by `config.kit.files.assets`. You can customize which files are included from `static` directory using [`config.kit.serviceWorker.files`](https://svelte.dev/docs/kit/configuration)
	 */
	export const files: string[];
	/**
	 * An array of pathnames corresponding to prerendered pages and endpoints.
	 * During development, this is an empty array.
	 */
	export const prerendered: string[];
	/**
	 * See [`config.kit.version`](https://svelte.dev/docs/kit/configuration#version). It's useful for generating unique cache names inside your service worker, so that a later deployment of your app can invalidate old caches.
	 */
	export const version: string;
}

//# sourceMappingURL=index.d.ts.map