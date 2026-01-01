<script>
  export let id = "";
  export let title = "";
  export let altThumb = "";
  export let play = false;
</script>

{#key play}
  <img
    src="https://i.ytimg.com/vi/{id}/{altThumb
      ? 'hqdefault'
      : 'maxresdefault'}.jpg"
    {title}
    alt="Youtube video: {title}"
    referrerpolicy="no-referrer"
  />
{/key}

<style>
  img {
    height: auto;
    aspect-ratio: var(--aspect-ratio);
    width: 100%;
  }
</style>
