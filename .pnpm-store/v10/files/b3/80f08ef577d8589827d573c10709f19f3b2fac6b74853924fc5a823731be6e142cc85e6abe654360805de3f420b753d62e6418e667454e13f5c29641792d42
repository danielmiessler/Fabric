import { SvelteComponent } from 'svelte';
import {
	ClientHooks,
	CSRPageNode,
	CSRPageNodeLoader,
	CSRRoute,
	CSRRouteServer,
	ServerDataNode,
	TrailingSlash,
	Uses
} from 'types';
import { Page, ParamMatcher } from '@sveltejs/kit';

export interface SvelteKitApp {
	/**
	 * A list of all the error/layout/page nodes used in the app.
	 * - In case of router.resolution=client, this is filled completely upfront.
	 * - In case of router.resolution=server, this is filled with the root layout and root error page
	 *   at the beginning and then filled up as the user navigates around the app, loading new nodes
	 */
	nodes: CSRPageNodeLoader[];

	/**
	 * A list of all layout node ids that have a server load function.
	 * Pages are not present because it's shorter to encode it on the leaf itself.
	 *
	 * In case of router.resolution=server, this only contains one entry for the root layout.
	 */
	server_loads: number[];

	/**
	 * A map of `[routeId: string]: [leaf, layouts, errors]` tuples, which
	 * is parsed into an array of routes on startup. The numbers refer to the indices in `nodes`.
	 * If the leaf number is negative, it means it does use a server load function and the complement is the node index.
	 * The route layout and error nodes are not referenced, they are always number 0 and 1 and always apply.
	 *
	 * In case of router.resolution=server, this object is empty, as resolution happens on the server.
	 */
	dictionary: Record<string, [leaf: number, layouts: number[], errors?: number[]]>;

	/**
	 * A map of `[matcherName: string]: (..) => boolean`, which is used to match route parameters.
	 *
	 * In case of router.resolution=server, this object is empty, as resolution happens on the server.
	 */
	matchers: Record<string, ParamMatcher>;

	hooks: ClientHooks;

	decode: (type: string, value: any) => any;

	decoders: Record<string, (data: any) => any>;
	encoders: Record<string, (data: any) => any>;

	/**
	 * Whether or not we're using hash-based routing
	 */
	hash: boolean;

	root: typeof SvelteComponent;
}

export type NavigationIntent = {
	/** `url.pathname + url.search`  */
	id: string;
	/** Whether we are invalidating or navigating */
	invalidating: boolean;
	/** The route parameters */
	params: Record<string, string>;
	/** The route that matches `path` */
	route: CSRRoute;
	/** The destination URL */
	url: URL;
};

export type NavigationResult = NavigationRedirect | NavigationFinished;

export type NavigationRedirect = {
	type: 'redirect';
	location: string;
};

export type NavigationFinished = {
	type: 'loaded';
	state: NavigationState;
	props: {
		constructors: Array<typeof SvelteComponent>;
		components?: SvelteComponent[];
		page: Page;
		form?: Record<string, any> | null;
		[key: `data_${number}`]: Record<string, any>;
	};
};

export type BranchNode = {
	node: CSRPageNode;
	loader: CSRPageNodeLoader;
	server: DataNode | null;
	universal: DataNode | null;
	data: Record<string, any> | null;
	slash?: TrailingSlash;
};

export interface DataNode {
	type: 'data';
	data: Record<string, any> | null;
	uses: Uses;
	slash?: TrailingSlash;
}

export interface NavigationState {
	branch: Array<BranchNode | undefined>;
	error: App.Error | null;
	params: Record<string, string>;
	route: CSRRoute | null;
	url: URL;
}

export interface HydrateOptions {
	status: number;
	error: App.Error | null;
	node_ids: number[];
	params: Record<string, string>;
	route: { id: string | null };
	/** Only used when `router.resolution=server`; can then still be undefined in case of 404 */
	server_route?: CSRRouteServer;
	data: Array<ServerDataNode | null>;
	form: Record<string, any> | null;
	/** The results of all remote functions executed during SSR so that they can be reused during hydration */
	remote: Record<string, any>;
}
