<script>import { getContext } from "svelte";
export let selected = false;
export let regionLead = "flex justify-center items-center";
export let regionLabel = "";
export let hover = getContext("hover");
export let active = getContext("active");
export let spacing = getContext("spacing");
export let aspectRatio = getContext("aspectRatio");
const cBase = "unstyled";
const cWrapper = "w-full flex flex-col justify-center items-stretch text-center space-y-1";
const cLabel = "font-bold text-xs";
$:
  classActive = selected ? active : "";
$:
  classesBase = `${cBase} ${$$props.class || ""}`;
$:
  classesWrapper = `${cWrapper} ${aspectRatio} ${hover} ${spacing} ${classActive}`;
$:
  classesLead = `${regionLead}`;
$:
  classesLabel = `${cLabel} ${regionLabel}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<a class="app-rail-anchor {classesBase}" href={$$props.href} {...prunedRestProps()} data-testid="app-rail-anchor" on:click>
	<div class="app-rail-wrapper {classesWrapper}">
		{#if $$slots.lead}<div class="app-rail-lead {classesLead}"><slot name="lead" /></div>{/if}
		<div class="app-rail-label {classesLabel}"><slot /></div>
	</div>
</a>
