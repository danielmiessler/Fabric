<script>import { afterUpdate } from "svelte";
export let name;
export let id = String(Math.random());
export let value = 0;
export let min = 0;
export let max = 100;
export let step = 1;
export let ticked = false;
export let accent = "accent-surface-900 dark:accent-surface-50";
export let label = "";
const cBase = "space-y-2";
const cBaseLabel = "";
const cBaseContent = "flex justify-center py-2";
const cBaseInput = "w-full h-2";
let tickmarks;
function setTicks() {
  if (ticked == false)
    return;
  tickmarks = Array.from({ length: max - min + 1 }, (_, i) => i + min);
}
if (ticked)
  setTicks();
afterUpdate(() => {
  setTicks();
});
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesInput = `${cBaseInput} ${accent}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<div class="range-slider {classesBase}" data-testid="range-slider">
	<!-- Slot: Default -->
	{#if $$slots.default}<label class="range-slider-label {cBaseLabel}" for={id}><slot /></label>{/if}

	<!-- Content -->
	<div class="range-content {cBaseContent}">
		<!-- Input -->
		<input
			type="range"
			{id}
			{name}
			class="range-slider-input {classesInput}"
			list="tickmarks-{id}"
			aria-label={label}
			{min}
			{max}
			{step}
			bind:value
			on:click
			on:change
			on:blur
			{...prunedRestProps()}
		/>

		<!-- Tickmarks -->
		{#if ticked && tickmarks && tickmarks.length}
			<datalist id="tickmarks-{id}" class="range-slider-ticks">
				{#each tickmarks as tm}
					<option value={tm} label={tm}></option>
				{/each}
			</datalist>
		{/if}
	</div>

	<!-- Slot: Trail -->
	{#if $$slots.trail}<div class="range-slider-trail"><slot name="trail" /></div>{/if}
</div>
