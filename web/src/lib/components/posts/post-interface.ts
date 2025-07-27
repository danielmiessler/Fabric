import type { SvelteComponent } from 'svelte';
import type { Frontmatter } from '$lib/utils/markdown';

export type PostMetadata = Frontmatter;

export interface Post {
    /** URL-friendly identifier for the post */
    slug: string;
    /** Post metadata from frontmatter */
    metadata: PostMetadata;
    /** Compiled Svelte component or HTML string */
    content: string | typeof SvelteComponent;
}
