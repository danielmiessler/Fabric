<script>import { getContext } from "svelte";
export let selected = false;
export let controls = "";
export let active = getContext("active");
export let hover = getContext("hover");
export let flex = getContext("flex");
export let padding = getContext("padding");
export let rounded = getContext("rounded");
export let spacing = getContext("spacing");
const cBase = "text-center cursor-pointer transition-colors duration-100";
const cInterface = "";
$:
  classesActive = selected ? active : hover;
$:
  classesBase = `${cBase} ${flex} ${padding} ${rounded} ${classesActive} ${$$props.class ?? ""}`;
$:
  classesInterface = `${cInterface} ${spacing}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<a
	class="tab-anchor {classesBase}"
	href={$$props.href}
	{...prunedRestProps()}
	aria-controls={controls}
	role="button"
	on:click
	on:keydown
	on:keyup
	on:keypress
	on:mouseover
	on:mouseleave
	on:focus
	on:blur
	data-testid="tab-anchor"
>
	<!-- Interface -->
	<div class="tab-interface {classesInterface}">
		{#if $$slots.lead}<div class="tab-lead"><slot name="lead" /></div>{/if}
		<div class="tab-label"><slot /></div>
	</div>
</a>
