<script>
  import Toc from '$lib/components/ui/toc/Toc.svelte';
  import Content from './README.md';
</script>

{#if Content}
  <div class="items-center mx-auto pt-8 grid-cols-[80%_20%] grid gap-8 max-w-7xl relative">
    <svelte:component this={Content} />
    <Toc />
  </div>
{:else}
  <div class="container py-12">
    <h1 class="mb-8 text-3xl font-bold">Sorry</h1>
    <div class="flex min-h-[400px] items-center justify-center text-center">
      <p class="text-lg font-medium">Nothing found</p>
      <p class="mt-2 text-sm text-muted-foreground">Check back later for new content.</p>
    </div>
  </div>
{/if}

