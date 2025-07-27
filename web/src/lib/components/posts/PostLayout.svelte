<script>
  export let aliases
  export let date 
  export let tags
  export let title
  export let description
  export let author
  export let updated
  //export let content
</script>

<article class="prose prose-slate dark:prose-invert max-w-5xl flex-1">
  {#if aliases}
  <h1 class="inline-block text-4xl font-bold inherit-colors lg:text-5xl">{aliases}</h1>
  {/if}
  <slot />
</article>

<style lang="postcss">

:global(h1) {
  @apply h1; 
}

:global(h2) {
  @apply h2; 
}

:global(p) {
  @layer p; 
}

:global(ul) {
  @layer ul; 
}

:global(li) {
  @layer li; 
}

:global(a) {
  @layer a; 
}

:global(img) {
  @layer img; 
}

:global(blockquote) {
  @layer blockquote; 
}

:global(code) {
  @layer code; 
}

</style> 
