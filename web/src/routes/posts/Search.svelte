<script lang="ts">
  import { formatDistance } from 'date-fns';
  import type { PageData } from './$types';
  import Card from '$lib/components/ui/cards/card.svelte';
  import { slide } from 'svelte/transition';
  import { elasticOut, quintOut } from 'svelte/easing';
  import { InputChip } from '@skeletonlabs/skeleton';

  let cards = false;
  let searchQuery = '';
  let selectedTags: string[] = [];
  let allTags: string[] = [];

  export let data: PageData;
  $: posts = data.posts;

  // Extract all unique tags from Posts
  $: {
    const tagSet = new Set<string>();
    posts.forEach(post => {
      post.meta.tags.forEach(tag => tagSet.add(tag));
    });
    allTags = Array.from(tagSet);
  }

  // Filter posts based on selected tags-container
  $: filteredPosts = posts.filter(post => {
    if (selectedTags.length === 0) return true;
    return selectedTags.every(tag =>
      post.meta.tags.some(postTag => postTag.toLowerCase() === tag.toLowerCase())
    );
  });

  function validateTag(value: string): boolean {
    return allTags.some(tag => tag.toLowerCase() === value.toLowerCase());
  }

  let visible: boolean = true;
</script>

<!-- This file can be deleted, It think it has better search functionality but it needs work to ...work
Could this be the new component for the search bar?

<script lang="ts">
	import { formatDistance } from 'date-fns';
	import type { PageData } from './$types';
	import Card from '$lib/components/ui/cards/card.svelte';
  	import { Youtube } from 'svelte-youtube-lite';
	import { slide } from 'svelte/transition';
	import { elasticOut, quintOut } from 'svelte/easing';
	import { InputChip } from '@skeletonlabs/skeleton';

	let cards = false;
	let searchQuery = '';
	let selectedTags: string[] = [];
	let allTags: string[] = [];

	export let data: PageData;
	$: posts = data.posts;
	
	// Extract all unique tags from posts
	$: {
		const tagSet = new Set<string>();
		posts.forEach(post => {
			post.meta.tags.forEach(tag => tagSet.add(tag));
		});
		allTags = Array.from(tagSet);
	}

	// Filter posts based on selected tags
	$: filteredPosts = posts.filter(post => {
		if (selectedTags.length === 0) return true;
		return selectedTags.every(tag => 
			post.meta.tags.some(postTag => postTag.toLowerCase() === tag.toLowerCase())
		);
	});

	function validateTag(value: string): boolean {
		return allTags.some(tag => tag.toLowerCase() === value.toLowerCase());
	}

	let visible: boolean = true;
</script>

<div class="container py-12">
	<h1 class="mb-4 text-3xl font-bold">Blog Posts</h1>
	<p class="text-sm mb-4 font-small">This blog is maintained in an Obsidian Vault</p>

	<div >
	  <div class="container mx-auto ml-auto grid grid-cols-1 md:grid-cols-2 gap-4 justify-end">
	    <div class="container mx-auto justify-left">
	      	<img src="https://img.shields.io/github/languages/top/danielmiessler/fabric" alt="Github top language">
	      	<img src="https://img.shields.io/github/last-commit/danielmiessler/fabric" alt="GitHub last commit">
	      	<img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
	      	<br>
	      	<hr class="!border-t-4" />
	      	<br>
	      	<h4 class="h4"><b>Leverage Proven Patterns</b></h4>
	      	<br>
	      	<Youtube id="UbDyjIIGaxQ" title="Network Chuck Explains fabric" />
			<p>Post your favorite videos.</p>
	      	<br>

	    </div>
		<div>
		<h4 class="h4"><b>Share Your Most Important Thoughts and Ideas</b></h4>
		<br>
	    <Card
	        header="Let Your Voice Be Heard"
	        imageUrl="/brain.png"
	        imageAlt="Blog post header image"
	        title="Blogging, Podcasting, Videos, and More."
	        content="What will you create?"
	        authorName="Your Name Here"
	        authorAvatarUrl=""
	        link="/"
	    />
		</div>
	  </div>

	  <div class="container mx-auto ml-auto grid grid-cols-1 md:grid-cols-2 gap-4 mt-8">
	    <Card
	        header="Curate Your Content"
	        imageUrl="/electric.png"
	        imageAlt="Blog post header image"
	        title="Enter a new title here"
	        content="What will you share"
	        authorName="Your Name Here"
	        authorAvatarUrl=""
	        link="/"
	    />
	    <div class="container mx-auto justify-right">
	      <blockquote class="blockquote">There are countless use cases for AI. What will you use if for?</blockquote>
	    </div>
	  </div>
	  <div class="container mx-auto ml-auto grid grid-cols-1 md:grid-cols-2 gap-4 justify-end mt-8 pb-8">
	    <div class="container mx-auto justify-left">
	      <hr class="!border-t-4" />
	      <br>
	      <h4 class="h4">Showcase your interests. Tell people what you've been working on. Create your community.</h4>
	    </div>

	    <Card
	        header="Explore the Possibilities"
	        imageUrl=""
	        imageAlt="Blog post header image"
	        title="Enter a new title here"
	        content="What will you share?"
	        authorName="Your Name Here"
	        authorAvatarUrl=""
	        link="/"
	    />
	  </div>
	</div>
-->

	<!-- Tag search and filter section -->
<div class="mb-6">
  <div class="flex flex-col gap-4">
    <div class="flex flex-col gap-2">
      <InputChip
        bind:value={selectedTags}
        name="tags"
        placeholder="Search and press Enter to add tags..."
        validation={validateTag}
        allowDuplicates={false}
        class="input"
      />
      <div class="tags-container overflow-x-auto pb-2">
        <div class="flex gap-2">
          {#each allTags.filter(tag => tag.toLowerCase().includes(searchQuery.toLowerCase())) as tag}
            <button
              class="tag-button px-3 py-1 rounded-full text-sm font-medium transition-colors
              {selectedTags.includes(tag.toLowerCase()) 
                ? 'bg-primary text-primary-foreground' 
                : 'bg-secondary hover:bg-secondary/80'}"
              on:click={() => {
                const tagLower = tag.toLowerCase();
                if (!selectedTags.includes(tagLower)) {
                  selectedTags = [...selectedTags, tagLower];
                }
                searchQuery = '';
              }}
            >
              {tag}
            </button>
          {/each}
        </div>
      </div>
    </div>
  </div>
</div>

{#if filteredPosts.length === 0}
  {#if !visible}
    <aside class="alert variant-ghost">
      <div>(icon)</div>
      <slot:fragment href="./+error.svelte" />
      <div class="alert-actions">(buttons)</div>
    </aside>
  {/if}
{:else}
  <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
    {#each filteredPosts as post}
      <article class="card card-hover group relative rounded-lg border p-6 hover:bg-muted/50">
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
            {#if post.meta.tags.length > 0}
              <span class="text-xs">•</span>
              <div class="flex flex-wrap gap-2">
                {#each post.meta.tags as tag}
                  <a
                    href="/tags/{tag}"
                    class="inline-flex items-center rounded-md border px-2 py-0.5 text-xs font-semibold transition-colors hover:bg-secondary"
                  >
                    {tag}
                  </a>
                {/each}
              </div>
            {/if}
          </div>
        </div>

      </article>
    {/each}
    <!-- 	<Paginator records={posts} limit={6} buttonClass="btn" /> -->
  </div>
{/if}



<style>
.tags-container {
  scrollbar-width: thin;
  scrollbar-color: var(--color-primary) transparent;
}

.tags-container::-webkit-scrollbar {
  height: 6px;
}

.tags-container::-webkit-scrollbar-track {
  background: transparent;
}

.tags-container::-webkit-scrollbar-thumb {
  background-color: var(--color-primary);
  border-radius: 6px;
}

.tag-button {
  white-space: nowrap;
}
</style>
