import { Pathname, RouteId, RouteParams } from '$app/types';

export type ResolveArgs<T extends RouteId | Pathname> = T extends RouteId
	? RouteParams<T> extends Record<string, never>
		? [route: T]
		: [route: T, params: RouteParams<T>]
	: [route: T];
