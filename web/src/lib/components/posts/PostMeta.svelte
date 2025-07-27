<script lang="ts">
  import { formatDistance } from 'date-fns';
  import type { PostMetadata } from './post-interface';

  export let data: PostMetadata;
  export let showUpdated = true;
</script>

<div class="flex flex-col gap-2 text-sm text-muted-foreground">
  <div class="flex items-center gap-2">
    <span><b>Published on: </b>{data.date}</span>
    {#if showUpdated && data.updated !== data.date}
      <span>· Updated on {data.updated}</span>
    {/if}
  </div>
  {#if data.author}
    <div class="text-xs font-bold">By {data.author}</div>
  {/if}
  {#if data.description}
    <p class="text-base font-bold">{data.description}</p>
  {/if}
  {#if data.tags && data.tags.length > 0}
    <div class="flex flex-wrap gap-2">
      {#each data.tags as tag}
        <a
          href="/tags/{tag}"
          class="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors hover:bg-muted/50"
        >
          {tag}
        </a>
      {/each}
    </div>
  {/if}
</div>
