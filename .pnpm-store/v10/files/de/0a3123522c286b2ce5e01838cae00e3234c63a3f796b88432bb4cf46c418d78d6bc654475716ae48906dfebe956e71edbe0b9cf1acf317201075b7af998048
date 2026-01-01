import { SvelteComponent } from 'svelte';
import {
	Config,
	ServerLoad,
	Handle,
	HandleServerError,
	KitConfig,
	Load,
	RequestHandler,
	ResolveOptions,
	Server,
	ServerInitOptions,
	HandleFetch,
	Actions,
	HandleClientError,
	Reroute,
	RequestEvent,
	SSRManifest,
	Emulator,
	Adapter,
	ServerInit,
	ClientInit,
	Transporter
} from '@sveltejs/kit';
import {
	HttpMethod,
	MaybePromise,
	PrerenderOption,
	RequestOptions,
	TrailingSlash
} from './private.js';

export interface ServerModule {
	Server: typeof InternalServer;
}

export interface ServerInternalModule {
	set_assets(path: string): void;
	set_building(): void;
	set_manifest(manifest: SSRManifest): void;
	set_prerendering(): void;
	set_private_env(environment: Record<string, string>): void;
	set_public_env(environment: Record<string, string>): void;
	set_read_implementation(implementation: (path: string) => ReadableStream): void;
	set_safe_public_env(environment: Record<string, string>): void;
	set_version(version: string): void;
	set_fix_stack_trace(fix_stack_trace: (error: unknown) => string): void;
}

export interface Asset {
	file: string;
	size: number;
	type: string | null;
}

export interface AssetDependencies {
	assets: string[];
	file: string;
	imports: string[];
	stylesheets: string[];
	fonts: string[];
	stylesheet_map: Map<string, { css: Set<string>; assets: Set<string> }>;
}

export interface BuildData {
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

export interface CSRPageNode {
	component: typeof SvelteComponent;
	universal: {
		load?: Load;
		trailingSlash?: TrailingSlash;
	};
}

export type CSRPageNodeLoader = () => Promise<CSRPageNode>;

/**
 * Definition of a client side route.
 * The boolean in the tuples indicates whether the route has a server load.
 */
export type CSRRoute = {
	id: string;
	exec(path: string): undefined | Record<string, string>;
	errors: Array<CSRPageNodeLoader | undefined>;
	layouts: Array<[has_server_load: boolean, node_loader: CSRPageNodeLoader] | undefined>;
	leaf: [has_server_load: boolean, node_loader: CSRPageNodeLoader];
};

/**
 * Definition of a client side route as transported via `<pathname>/__route.js` when using server-side route resolution.
 */
export type CSRRouteServer = {
	id: string;
	errors: Array<number | undefined>;
	layouts: Array<[has_server_load: boolean, node_id: number] | undefined>;
	leaf: [has_server_load: boolean, node_id: number];
	nodes: Record<string, CSRPageNodeLoader>;
};

export interface Deferred {
	fulfil: (value: any) => void;
	reject: (error: Error) => void;
}

export type GetParams = (match: RegExpExecArray) => Record<string, string>;

export interface ServerHooks {
	handleFetch: HandleFetch;
	handle: Handle;
	handleError: HandleServerError;
	reroute: Reroute;
	transport: Record<string, Transporter>;
	init?: ServerInit;
}

export interface ClientHooks {
	handleError: HandleClientError;
	reroute: Reroute;
	transport: Record<string, Transporter>;
	init?: ClientInit;
}

export interface Env {
	private: Record<string, string>;
	public: Record<string, string>;
}

export class InternalServer extends Server {
	init(options: ServerInitOptions): Promise<void>;
	respond(
		request: Request,
		options: RequestOptions & {
			prerendering?: PrerenderOptions;
			read: (file: string) => Buffer;
			/** A hook called before `handle` during dev, so that `AsyncLocalStorage` can be populated. */
			before_handle?: (event: RequestEvent, config: any, prerender: PrerenderOption) => void;
			emulator?: Emulator;
		}
	): Promise<Response>;
}

export interface ManifestData {
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

export interface PageNode {
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

export interface PrerenderDependency {
	response: Response;
	body: null | string | Uint8Array;
}

export interface PrerenderOptions {
	cache?: string; // including this here is a bit of a hack, but it makes it easy to add <meta http-equiv>
	fallback?: boolean;
	dependencies: Map<string, PrerenderDependency>;
	/** True for the duration of a call to the `reroute` hook */
	inside_reroute?: boolean;
}

export type RecursiveRequired<T> = {
	// Recursive implementation of TypeScript's Required utility type.
	// Will recursively continue until it reaches a primitive or Function
	// eslint-disable-next-line @typescript-eslint/no-unsafe-function-type
	[K in keyof T]-?: Extract<T[K], Function> extends never // If it does not have a Function type
		? RecursiveRequired<T[K]> // recursively continue through.
		: T[K]; // Use the exact type for everything else
};

export type RequiredResolveOptions = Required<ResolveOptions>;

export interface RouteParam {
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
export interface RouteData {
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

export type ServerRedirectNode = {
	type: 'redirect';
	location: string;
};

export type ServerNodesResponse = {
	type: 'data';
	/**
	 * If `null`, then there was no load function <- TODO is this outdated now with the recent changes?
	 */
	nodes: Array<ServerDataNode | ServerDataSkippedNode | ServerErrorNode | null>;
};

export type ServerDataResponse = ServerRedirectNode | ServerNodesResponse;

/**
 * Signals a successful response of the server `load` function.
 * The `uses` property tells the client when it's possible to reuse this data
 * in a subsequent request.
 */
export interface ServerDataNode {
	type: 'data';
	/**
	 * The serialized version of this contains a serialized representation of any deferred promises,
	 * which will be resolved later through chunk nodes.
	 */
	data: Record<string, any> | null;
	uses: Uses;
	slash?: TrailingSlash;
}

/**
 * Resolved data/error of a deferred promise.
 */
export interface ServerDataChunkNode {
	type: 'chunk';
	id: number;
	data?: Record<string, any>;
	error?: any;
}

/**
 * Signals that the server `load` function was not run, and the
 * client should use what it has in memory.
 */
export interface ServerDataSkippedNode {
	type: 'skip';
}

/**
 * Signals that the server `load` function failed.
 */
export interface ServerErrorNode {
	type: 'error';
	error: App.Error;
	/**
	 * Only set for HttpErrors.
	 */
	status?: number;
}

export interface ServerMetadataRoute {
	config: any;
	api: {
		methods: Array<HttpMethod | '*'>;
	};
	page: {
		methods: Array<'GET' | 'POST'>;
	};
	methods: Array<HttpMethod | '*'>;
	prerender: PrerenderOption | undefined;
	entries: string[] | undefined;
}

export interface ServerMetadata {
	nodes: Array<{
		/** Also `true` when using `trailingSlash`, because we need to do a server request in that case to get its value. */
		has_server_load: boolean;
	}>;
	routes: Map<string, ServerMetadataRoute>;
}

export interface SSRComponent {
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

export type SSRComponentLoader = () => Promise<SSRComponent>;

export interface UniversalNode {
	load?: Load;
	prerender?: PrerenderOption;
	ssr?: boolean;
	csr?: boolean;
	trailingSlash?: TrailingSlash;
	config?: any;
	entries?: PrerenderEntryGenerator;
}

export interface ServerNode {
	load?: ServerLoad;
	prerender?: PrerenderOption;
	ssr?: boolean;
	csr?: boolean;
	trailingSlash?: TrailingSlash;
	actions?: Actions;
	config?: any;
	entries?: PrerenderEntryGenerator;
}

export interface SSRNode {
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

export type SSRNodeLoader = () => Promise<SSRNode>;

export interface SSROptions {
	app_template_contains_nonce: boolean;
	csp: ValidatedConfig['kit']['csp'];
	csrf_check_origin: boolean;
	embedded: boolean;
	env_public_prefix: string;
	env_private_prefix: string;
	hash_routing: boolean;
	hooks: ServerHooks;
	preload_strategy: ValidatedConfig['kit']['output']['preloadStrategy'];
	root: SSRComponent['default'];
	service_worker: boolean;
	templates: {
		app(values: {
			head: string;
			body: string;
			assets: string;
			nonce: string;
			env: Record<string, string>;
		}): string;
		error(values: { message: string; status: number }): string;
	};
	version_hash: string;
}

export interface PageNodeIndexes {
	errors: Array<number | undefined>;
	layouts: Array<number | undefined>;
	leaf: number;
}

export type PrerenderEntryGenerator = () => MaybePromise<Array<Record<string, string>>>;

export type SSREndpoint = Partial<Record<HttpMethod, RequestHandler>> & {
	prerender?: PrerenderOption;
	trailingSlash?: TrailingSlash;
	config?: any;
	entries?: PrerenderEntryGenerator;
	fallback?: RequestHandler;
};

export interface SSRRoute {
	id: string;
	pattern: RegExp;
	params: RouteParam[];
	page: PageNodeIndexes | null;
	endpoint: (() => Promise<SSREndpoint>) | null;
	endpoint_id?: string;
}

export interface SSRClientRoute {
	id: string;
	pattern: RegExp;
	params: RouteParam[];
	errors: Array<number | undefined>;
	layouts: Array<[has_server_load: boolean, node_id: number] | undefined>;
	leaf: [has_server_load: boolean, node_id: number];
}

export interface SSRState {
	fallback?: string;
	getClientAddress(): string;
	/**
	 * True if we're currently attempting to render an error page.
	 */
	error: boolean;
	/**
	 * Allows us to prevent `event.fetch` from making infinitely looping internal requests.
	 */
	depth: number;
	platform?: any;
	prerendering?: PrerenderOptions;
	/**
	 * When fetching data from a +server.js endpoint in `load`, the page's
	 * prerender option is inherited by the endpoint, unless overridden.
	 */
	prerender_default?: PrerenderOption;
	read?: (file: string) => Buffer;
	/**
	 * Used to setup `__SVELTEKIT_TRACK__` which checks if a used feature is supported.
	 * E.g. if `read` from `$app/server` is used, it checks whether the route's config is compatible.
	 */
	before_handle?: (event: RequestEvent, config: any, prerender: PrerenderOption) => void;
	emulator?: Emulator;
}

export type StrictBody = string | ArrayBufferView;

export interface Uses {
	dependencies: Set<string>;
	params: Set<string>;
	parent: boolean;
	route: boolean;
	url: boolean;
	search_params: Set<string>;
}

export type ValidatedConfig = Config & {
	kit: ValidatedKitConfig;
	extensions: string[];
};

export type ValidatedKitConfig = Omit<RecursiveRequired<KitConfig>, 'adapter'> & {
	adapter?: Adapter;
};

export * from '../exports/index.js';
export * from './private.js';
