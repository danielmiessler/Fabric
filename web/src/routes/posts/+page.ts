import type { PageLoad } from './$types';
import type { Frontmatter } from '$lib/utils/markdown';

// This is duplicated at components/ui/tagSearch/tags.ts
// Consider removing this duplication

const posts = import.meta.glob<{ metadata: Frontmatter }>('/src/lib/content/posts/*.{md,svx}', { eager: true });

export const load: PageLoad = async () => {
    try {
        const allPosts = Object.entries(posts).map(([path, post]) => ({
            slug: path.split('/').pop()?.replace(/\.(md|svx)$/, '') ?? '',
            metadata: post.metadata,
                /* date: post.metadata.date,
                updated: post.metadata.updated || post.metadata.date */
            //}
        }));

        // Sort posts by date, newest first
        allPosts.sort((a, b) => 
            new Date(b.metadata.date).getTime() - new Date(a.metadata.date).getTime()
        );

        return { posts: allPosts };
    } catch (e) {
        console.error('Failed to load posts:', e);
        throw Error();
    }
};
