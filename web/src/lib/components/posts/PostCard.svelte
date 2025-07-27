<script lang="ts">
  import { formatDistance } from 'date-fns';
  import type { Post } from './post-interface';
  import PostMeta from './PostMeta.svelte';
  import Card from '$lib/components/ui/cards/card.svelte';
  import { cn } from '$lib/utils/utils';

  export let post: Post;
  export let className: string = '';

  function parseDate(dateStr: string): Date {
      // Handle both ISO strings and YYYY-MM-DD formats
      return new Date(dateStr);
  }
</script>

<article class="card card-hover group relative rounded-lg border p-6 hover:bg-primary-500/50 {className}">
  <a 
    href="/posts/{post.slug}"
    class="absolute inset-0" 
    data-sveltekit-preload-data="off"
  >
    <span class="sr-only">View {post.metadata?.title}</span>
  </a>
  <div class="flex flex-col justify-between space-y-4">
    <div class="space-y-2">
      <!-- <img src={post.metadata?.images?.[0]} alt="Posts Cards" class="rounded-lg" /> -->
      <h2 class="text-xl font-semibold tracking-tight">{post.metadata?.title}</h2>
      <p class="text-muted-foreground">{post.metadata?.description}</p>
    </div>
    <div class="flex items-center space-x-4 text-sm text-muted-foreground">
      <time datetime={post.metadata?.date}>
        {#if post.metadata?.date}
          {formatDistance(parseDate(post.metadata.date), new Date(), { addSuffix: false })}
        {/if}
      </time>
      {#if post.metadata?.tags?.length > 0}
        <span class="text-xs">•</span>
        <div class="flex flex-wrap gap-2">
          {#each post.metadata?.tags as tag}
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
