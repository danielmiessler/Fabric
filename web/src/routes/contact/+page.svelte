<script>
  import Content from './contact.md'
</script>

{#if Content}
  <div class="container max-w-3xl">
    <div class="space-y-4 mx-auto py-8">
      <svelte:component this={Content} />
    </div>
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
