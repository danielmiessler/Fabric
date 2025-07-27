import type { PageLoad } from './$types';
import type { Frontmatter } from '$lib/utils/markdown';

export const load: PageLoad = async () => {
	const postFiles = import.meta.glob('/src/lib/content/posts/*.{md,svx}', { eager: true });
	
	const posts = Object.entries(postFiles).map(([path, post]: [string, any]) => {
		const slug = path.split('/').pop()?.replace(/\.(md|svx)$/, '');
		return {
			slug,
			meta: {
				title: post.metadata.title,
				date: post.metadata.date,
				description: post.metadata.description,
				tags: post.metadata.tags || []
			}
		};
	});

	const tags = posts.reduce((acc, post) => {
		post.meta.tags.forEach((tag: string) => {
			if (!acc[tag]) {
				acc[tag] = [];
			}
			acc[tag].push(post);
		});
		return acc;
	}, {} as Record<string, Frontmatter[]>);

	return {
		tags,
		postsCount: posts.length
	};
};
