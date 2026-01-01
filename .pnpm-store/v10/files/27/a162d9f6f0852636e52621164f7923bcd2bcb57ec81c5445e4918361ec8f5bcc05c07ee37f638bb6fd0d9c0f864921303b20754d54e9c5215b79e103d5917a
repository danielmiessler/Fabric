<script>import { createEventDispatcher } from "svelte";
export let value = 0;
export let max = 5;
export let interactive = false;
export let text = "text-token";
export let fill = "fill-token";
export let justify = "justify-center";
export let spacing = "space-x-2";
export let regionIcon = "";
const dispatch = createEventDispatcher();
function iconClick(index) {
  dispatch("icon", {
    index: index + 1
  });
}
function isFull(value2, index) {
  return Math.floor(value2) >= index + 1;
}
function isHalf(value2, index) {
  return value2 === index + 0.5;
}
const cBase = "w-full flex";
$:
  classesBase = `${cBase} ${text} ${fill} ${justify} ${spacing} ${$$props.class ?? ""}`;
</script>

<div class="ratings {classesBase}" data-testid="rating-bar">
	<!-- eslint-disable-next-line @typescript-eslint/no-unused-vars -->
	{#each { length: max } as _, i}
		{#if interactive}
			<button class="rating-icon {regionIcon}" type="button" on:click={() => iconClick(i)}>
				{#if isFull(value, i)}
					<slot name="full" />
				{:else if isHalf(value, i)}
					<slot name="half" />
				{:else}
					<slot name="empty" />
				{/if}
			</button>
		{:else}
			<span class="rating-icon {regionIcon}">
				{#if isFull(value, i)}
					<slot name="full" />
				{:else if isHalf(value, i)}
					<slot name="half" />
				{:else}
					<slot name="empty" />
				{/if}
			</span>
		{/if}
	{/each}
</div>
