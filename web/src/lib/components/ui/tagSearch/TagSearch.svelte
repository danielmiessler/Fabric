<script lang="ts">
  import { InputChip } from '@skeletonlabs/skeleton';
  import type { PostMetadata } from '$lib/types';
  import type { Post } from '$lib/interfaces/post-interface'
  import PostCard from '$lib/components/posts/PostCard.svelte';

  let searchQuery = '';
  let selectedTags: string[] = [];
  let allTags: string[] = [];

  let data: PageData;
  let posts = data.posts || [];

  // Extract all unique tags from posts
  $: {
    const tagSet = new Set<string>();
    posts?.forEach(post => {
      post.metadata?.tags?.forEach(tag => tagSet.add(tag));
    });
    allTags = Array.from(tagSet);
  }

  // Filter posts based on selected tags
  $: filteredPosts = posts?.filter(post => {
    if (selectedTags.length === 0) return true;
    return selectedTags.every(tag => 
      post.metadata?.tags?.some(postTag => postTag.toLowerCase() === tag.toLowerCase())
    );
  }) || [];

  // Filter posts based on search query
  $: searchResults = filteredPosts.filter(post => {
    if (!searchQuery) return true;
    const query = searchQuery.toLowerCase();
    return (
      post.metadata?.title?.toLowerCase().includes(query) ||
      post.metadata?.description?.toLowerCase().includes(query) ||
      post.metadata?.tags?.some(tag => tag.toLowerCase().includes(query))
    );
  });

  function validateTag(value: string): boolean {
    return allTags.some(tag => tag.toLowerCase() === value.toLowerCase());
  }
</script>


<div class="container py-12">
	<div class="my-4">
		<InputChip
			name="tags"
			placeholder="Filter by tags..."
			validation={validateTag}
			bind:value={selectedTags}
			/>
  </div>
</div>
<div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
	{#each searchResults as post}
		<PostCard {post} /> <!-- TODO: Add images to post metadata --> 
	{/each}
 </div>

