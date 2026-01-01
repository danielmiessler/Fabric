<script>
  export let isCustomPlayButton;
</script>

{#if isCustomPlayButton}
  <button class="custom__play__btn" on:click aria-label="Play YouTube video">
    <slot />
  </button>
{:else}
  <button class="play__btn" on:click aria-label="Play YouTube video"
    ><svg
      xmlns="http://www.w3.org/2000/svg"
      aria-hidden="true"
      class="iconify iconify--mdi"
      viewBox="0 0 24 24"
    >
      <path
        fill="#ff4e45"
        d="m10 15 5.19-3L10 9v6m11.56-7.83c.13.47.22 1.1.28 1.9.07.8.1 1.49.1 2.09L22 12c0 2.19-.16 3.8-.44 4.83-.25.9-.83 1.48-1.73 1.73-.47.13-1.33.22-2.65.28-1.3.07-2.49.1-3.59.1L12 19c-4.19 0-6.8-.16-7.83-.44-.9-.25-1.48-.83-1.73-1.73-.13-.47-.22-1.1-.28-1.9-.07-.8-.1-1.49-.1-2.09L2 12c0-2.19.16-3.8.44-4.83.25-.9.83-1.48 1.73-1.73.47-.13 1.33-.22 2.65-.28 1.3-.07 2.49-.1 3.59-.1L12 5c4.19 0 6.8.16 7.83.44.9.25 1.48.83 1.73 1.73Z"
      />
    </svg>
  </button>
{/if}

<style>
  .play__btn,
  .custom__play__btn {
    all: unset;
  }
  .play__btn {
    display: grid;
    place-items: center;
    height: 80px;
    width: 80px;
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    margin: auto;
    cursor: pointer;
  }
  .custom__play__btn {
    display: grid;
    place-items: center;
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    margin: auto;
    cursor: pointer;
  }
</style>
