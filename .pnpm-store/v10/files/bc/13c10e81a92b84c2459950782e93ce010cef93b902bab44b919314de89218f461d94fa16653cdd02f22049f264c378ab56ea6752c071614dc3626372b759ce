<script>import { setContext } from "svelte";
export let justify = "justify-start";
export let border = "border-b border-surface-400-500-token";
export let active = "border-b-2 border-surface-900-50-token";
export let hover = "hover:variant-soft";
export let flex = "flex-none";
export let padding = "px-4 py-2";
export let rounded = "rounded-tl-container-token rounded-tr-container-token";
export let spacing = "space-y-1";
export let regionList = "";
export let regionPanel = "";
export let labelledby = "";
export let panel = "";
setContext("active", active);
setContext("hover", hover);
setContext("flex", flex);
setContext("padding", padding);
setContext("rounded", rounded);
setContext("spacing", spacing);
const cBase = "space-y-4";
const cList = "flex overflow-x-auto hide-scrollbar";
const cPanel = "";
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesList = `${cList} ${justify} ${border} ${regionList}`;
$:
  classesPanel = `${cPanel} ${regionPanel}`;
</script>

<div class="tab-group {classesBase}" data-testid="tab-group">
	<!-- Tab List -->
	<div class="tab-list {classesList}" role="tablist" aria-labelledby={labelledby}>
		<slot />
	</div>
	<!-- Tab Panel -->
	{#if $$slots.panel}
		<div class="tab-panel {classesPanel}" role="tabpanel" aria-labelledby={panel} tabindex="0">
			<slot name="panel" />
		</div>
	{/if}
</div>
