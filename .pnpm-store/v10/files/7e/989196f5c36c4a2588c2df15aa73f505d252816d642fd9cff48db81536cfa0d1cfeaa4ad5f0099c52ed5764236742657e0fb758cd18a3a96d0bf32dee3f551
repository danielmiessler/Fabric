import type { ComponentType } from 'svelte';
export interface TreeViewNode {
    /** Nodes Unique ID */
    id: string;
    /** Main content. accepts HTML or svelte component. */
    content: string | ComponentType;
    /** Main content props. only used when the Content is a svelte component. */
    contentProps?: object;
    /** Lead content. accepts HTML or svelte component. */
    lead?: string | ComponentType;
    /** lead props. only used when the Lead is a svelte component. */
    leadProps?: object;
    /** children nodes. */
    children?: TreeViewNode[];
    /** Set the input's value. */
    value?: unknown;
}
