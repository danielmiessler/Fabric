<script lang="ts">
  import { ChevronLeft, ChevronRight } from 'lucide-svelte';
  import { slide } from 'svelte/transition';
  import { cn } from '$lib/utils/utils';

  export let tags: string[] = [];
  export let tagsPerPage = 5;
  export let className: string | undefined = undefined;

  let currentPage = 0;
  let containerWidth: number;

  $: totalPages = Math.ceil(tags.length / tagsPerPage);
  $: startIndex = currentPage * tagsPerPage;
  $: endIndex = Math.min(startIndex + tagsPerPage, tags.length);
  $: visibleTags = tags.slice(startIndex, endIndex);
  $: canGoBack = currentPage > 0;
  $: canGoForward = currentPage < totalPages - 1;

  function nextPage() {
    if (canGoForward) {
      currentPage++;
    }
  }

  function prevPage() {
    if (canGoBack) {
      currentPage--;
    }
  }
</script>

<div class={cn('relative flex items-center gap-2', className)} bind:clientWidth={containerWidth}>
	{#if totalPages > 1 && canGoBack}
		<button
			on:click={prevPage}
			class="flex h-6 w-6 items-center justify-center rounded-md border bg-background hover:bg-muted"
			transition:slide
		>
			<ChevronLeft class="h-4 w-4" />
			<span class="sr-only">Previous page</span>
		</button>
	{/if}

	<div class="flex flex-wrap gap-2">
		{#each visibleTags as tag (tag)}
			<a
				href="/tags/{tag}"
				class="inline-flex items-center rounded-md border px-2 py-0.5 text-xs font-semibold transition-colors hover:bg-muted"
			>
				{tag}
			</a>
		{/each}
	</div>

	{#if totalPages > 1 && canGoForward}
		<button
			on:click={nextPage}
			class="flex h-6 w-6 items-center justify-center rounded-md border bg-background hover:bg-muted"
			transition:slide
		>
			<ChevronRight class="h-4 w-4" />
			<span class="sr-only">Next page</span>
		</button>
	{/if}
</div>
