<script>
  import { disableScrollHandling } from "$app/navigation";
  import { onMount } from "svelte";

  onMount(() => {
    disableScrollHandling();
  });
</script>

<div id="page" class="page-wrapper">
  <div class="viewport-container flex h-[calc(100vh-8rem)]">
    <slot />
  </div>
</div>

<style>
  .viewport-container {
    width: 100vw;
    overflow: hidden;
    position: fixed;
    left: 0;
  }

  .page-wrapper :global(#page) {
    display: block;
    flex: none;
    overflow: hidden;
  }

  :global(#page-content) {
    flex: none;
    overflow: hidden;
  }

  :global(.viewport-container *) {
    overflow: hidden;
  }
</style>

