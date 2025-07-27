<script lang="ts">
  import { fade, scale } from 'svelte/transition';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher<{
    close: void;
  }>();

  export let show = false;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions a11y-no-static-element-interactions -->
{#if show}
<div
class="fixed inset-0 bg-black/50 flex items-center justify-center z-50 mt-2"
on:click={() => dispatch('close')}
on:keydown={(e) => e.key === 'Escape' && dispatch('close')}
role="dialog"
aria-modal="true"
aria-label="Modal dialog"
tabindex="-1"
transition:fade={{ duration: 200 }}
>
<div
class="relative"
on:click|stopPropagation
role="document"
aria-label="Modal content"
transition:scale={{ duration: 200 }}
>
<slot />
</div>
</div>
{/if}



<style>
  .fixed {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
</style>
