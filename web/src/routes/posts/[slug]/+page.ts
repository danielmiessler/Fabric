import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import type { Frontmatter } from '$lib/utils/markdown';

const posts = import.meta.glob<{ metadata: Frontmatter, default: unknown }>('/src/lib/content/posts/*.{md,svx}', { eager: true });

export const load: PageLoad = async ({ params }) => {
    const post = Object.entries(posts).find(([path]) => 
        path.endsWith(`${params.slug}.md`) || path.endsWith(`${params.slug}.svx`)
    );

    if (!post) {
        throw error(404, `Post ${params.slug} not found`);
    }

    function formatDateOnly(dateStr: string): string {
        const date = new Date(dateStr);
        return date.toISOString().split('T')[0];
    }

    return {
        content: post[1].default,
        metadata: {
            ...post[1].metadata,
            // Only keep the date portion YYYY-MM-DD
            date: formatDateOnly(post[1].metadata.date),
            updated: post[1].metadata.updated 
                ? formatDateOnly(post[1].metadata.updated)
                : formatDateOnly(post[1].metadata.date)
        },
        slug: params.slug
    };
};
