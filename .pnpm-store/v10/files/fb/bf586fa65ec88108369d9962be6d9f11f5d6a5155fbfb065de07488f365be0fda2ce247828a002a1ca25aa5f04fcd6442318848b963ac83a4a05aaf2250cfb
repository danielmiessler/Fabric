import { CookieSerializeOptions } from 'cookie';
import {
	CspDirectives,
	ServerDataNode,
	SSRNode,
	ServerDataSkippedNode,
	ServerErrorNode
} from 'types';
import { Csp } from './csp.js';

export interface Fetched {
	url: string;
	method: string;
	request_body?: string | ArrayBufferView | null;
	request_headers?: HeadersInit | undefined;
	response_body: string | undefined;
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

export type ServerDataSerializer = {
	add_node(i: number, node: ServerDataNode | null): void;
	get_data(csp: Csp): { data: string; chunks: AsyncIterable<string> | null };
	set_max_nodes(i: number): void;
};

export type ServerDataSerializerJson = {
	add_node(
		i: number,
		node: ServerDataSkippedNode | ServerDataNode | ServerErrorNode | null | undefined
	): void;
	get_data(): { data: string; chunks: AsyncIterable<string> | null };
};
