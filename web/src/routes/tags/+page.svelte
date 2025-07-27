<script lang="ts">
  import type { PageData } from './$types';
  // import TagList from '$components/ui/tag-list/TagList.svelte';

  export let data: PageData;

  $: tags = data.tags;
  $: postsCount = data.postsCount;
</script>

<div class="container py-12">
	<h1 class="mb-8 text-3xl font-bold">{Object.keys(tags).length} Tags in {postsCount} Posts</h1>

	<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
		{#each Object.entries(tags) as [tag, posts]}
			<a
				href="/tags/{tag}"
				class="flex items-center justify-between rounded-lg border p-4 hover:bg-muted/50"
			>
				<span class="text-lg font-semibold">{tag}</span>
				<span class="text-sm text-muted-foreground">{posts.length} posts</span>
				<!-- WIP -->
				<!-- <TagList tags={posts.map((post) => post.meta.tags).flat()} /> -->
			</a>
		{/each}
	</div>
</div>
