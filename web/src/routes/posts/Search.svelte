<script lang="ts">
  import { formatDistance } from 'date-fns';
  import type { PageData } from './$types';
  import Card from '$lib/components/ui/cards/card.svelte';
  import { slide } from 'svelte/transition';
  import { elasticOut, quintOut } from 'svelte/easing';
  import TagsInputField from '$lib/components/ui/tagsInput/TagsInputField.svelte';

  let cards = false;
  let searchQuery = '';
  let selectedTags: string[] = [];
  let allTags: string[] = [];

  export let data: PageData;
  $: posts = data.posts;

  $: {
    const tagSet = new Set<string>();
    posts.forEach(post => {
      post.meta.tags.forEach(tag => tagSet.add(tag));
    });
    allTags = Array.from(tagSet);
  }

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

<!-- No file imports this component. -->

<div class="mb-6">
  <div class="flex flex-col gap-4">
    <div class="flex flex-col gap-2">
      <TagsInputField
        bind:value={selectedTags}
        name="tags"
        placeholder="Search and press Enter to add tags..."
        validation={validateTag}
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
    <aside class="alert preset-tonal">
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
