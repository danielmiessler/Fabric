<script>import { getContext } from "svelte";
export let group;
export let name;
export let value;
export let title = "";
export let regionLead = "";
export let regionLabel = "";
export let hover = getContext("hover");
export let active = getContext("active");
export let spacing = getContext("spacing");
export let width = getContext("width");
export let aspectRatio = getContext("aspectRatio");
const cBase = "cursor-pointer";
const cWrapper = "flex flex-col justify-center items-stretch w-full";
const cInterface = "text-center";
const cLabel = "font-bold text-xs";
let elemInput;
$:
  classActive = group === value ? active : "";
$:
  classesBase = `${cBase} ${$$props.class || ""}`;
$:
  classesWrapper = `${cWrapper} ${aspectRatio} ${width} ${hover} ${classActive}`;
$:
  classesInterface = `${cInterface} ${spacing}`;
$:
  classesLead = `${regionLead}`;
$:
  classesLabel = `${cLabel} ${regionLabel}`;
function onKeyDown(event) {
  if (["Enter", "Space"].includes(event.code)) {
    event.preventDefault();
    selectElemInput();
  }
}
function selectElemInput() {
  elemInput.click();
}
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<label class="app-rail-tile {classesBase}" data-testid="app-rail-tile" {title} on:mouseover on:mouseleave on:focus on:blur>
	<!-- A11y attributes are not allowed on <label> -->
	<!-- FIXME: resolve a11y warnings -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<button class="app-rail-wrapper {classesWrapper}" tabindex="0" on:click={selectElemInput} on:click on:keydown={onKeyDown} on:keyup on:keypress>
		<!-- NOTE: Don't use `hidden` as it prevents `required` from operating -->
		<div class="h-0 w-0 overflow-hidden">
			<input bind:this={elemInput} type="radio" bind:group {name} {value} {...prunedRestProps()} tabindex="-1" on:click on:change />
		</div>
		<!-- Interface -->
		<div class="app-rail-interface {classesInterface}">
			{#if $$slots.lead}<div class="app-rail-lead {classesLead}"><slot name="lead" /></div>{/if}
			<div class="app-rail-label {classesLabel}"><slot /></div>
		</div>
	</button>
</label>
