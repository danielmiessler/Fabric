import { CookieSerializeOptions } from 'cookie';
import { SSRNode, CspDirectives, ServerDataNode } from 'types';

export interface Fetched {
	url: string;
	method: string;
	request_body?: string | ArrayBufferView | null;
	request_headers?: HeadersInit | undefined;
	response_body: string;
	response: Response;
	is_b64?: boolean;
}

export type Loaded = {
	node: SSRNode;
	data: Record<string, any> | null;
	server_data: ServerDataNode | null;
};

type CspMode = 'hash' | 'nonce' | 'auto';

export interface CspConfig {
	mode: CspMode;
	directives: CspDirectives;
	reportOnly: CspDirectives;
}

export interface CspOpts {
	prerender: boolean;
}

export interface Cookie {
	name: string;
	value: string;
	options: CookieSerializeOptions & { path: string };
}
