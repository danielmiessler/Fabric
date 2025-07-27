<script lang="ts">
  import PostMeta from './PostMeta.svelte';
  import type { Post } from './post-interface'
  import Spinner from '$lib/components/ui/spinner/spinner.svelte';
  import Toc from '$lib/components/ui/toc/Toc.svelte';

  export let post: Post; 
</script>

<article class="py-6">
  {#if !post?.content || !post?.metadata}
    <div class="flex min-h-[400px] items-center justify-center">
      <div class="flex items-center gap-2">
        <Spinner class="h-6 w-6" />
        <span class="text-sm text-muted-foreground">Loading post...</span>
      </div>
    </div>
  {:else}
    <div class="space-y-4 pl-8 ml-8">
      <h1 class="inline-block text-4xl font-bold inherit-colors lg:text-5xl">{post.metadata.title}</h1>
      <PostMeta data={post.metadata} />
    </div> 
    <div class="items-center py-8 mx-auto gap-8 max-w-7xl relative prose prose-slate dark:prose-invert">
      {#if typeof post.content === 'function'}
        <Toc />
        <svelte:component this={post.content} />
      {:else if typeof post.content === 'string'}
        {post.content}
      {:else}
        <div class="flex gap-2">
          <Spinner class="h-8 w-8" />
          <span class="text-sm text-muted-foreground">Loading content...</span>
        </div>
      {/if}
    </div>
  {/if}
</article>
