<script lang="ts">
  import { formatDistance } from 'date-fns';
  import type { PageData } from './$types';

  export let data: PageData;

  $: ({ tag, posts } = data);
</script>

<div class="container py-12">
	<div class="mb-8 flex items-center justify-between">
		<h1 class="text-3xl font-bold">Posts tagged with "{tag}"</h1>
		<a href="/tags" class="text-sm text-muted-foreground hover:text-foreground">← Back to tags</a>
	</div>

	<div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
		{#each posts as post}
			<article class="group relative rounded-lg border p-6 hover:bg-muted/50">
				<a href="/posts/{post.slug}" class="absolute inset-0">
					<span class="sr-only">View {post.meta.title}</span>
				</a>
				<div class="flex flex-col justify-between space-y-4">
					<div class="space-y-2">
						<h2 class="text-xl font-semibold tracking-tight">{post.meta.title}</h2>
						<p class="text-muted-foreground">{post.meta.description}</p>
					</div>
					<div class="flex items-center space-x-4 text-sm text-muted-foreground">
						<time datetime={post.meta.date}>
							{formatDistance(new Date(post.meta.date), new Date(), { addSuffix: true })}
						</time>
					</div>
				</div>
			</article>
		{/each}
	</div>
</div>
