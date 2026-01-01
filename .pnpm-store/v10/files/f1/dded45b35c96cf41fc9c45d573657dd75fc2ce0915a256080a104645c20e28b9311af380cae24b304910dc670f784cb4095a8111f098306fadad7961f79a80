<script>import { createEventDispatcher, setContext } from "svelte";
import RecursiveTreeViewItem from "./RecursiveTreeViewItem.svelte";
export let selection = false;
export let multiple = false;
export let relational = false;
export let nodes = [];
export let expandedNodes = [];
export let disabledNodes = [];
export let checkedNodes = [];
export let indeterminateNodes = [];
export let width = "w-full";
export let spacing = "space-y-1";
export let padding = "py-4 px-4";
export let indent = "ml-4";
export let hover = "hover:variant-soft";
export let rounded = "rounded-container-token";
export let caretOpen = "rotate-180";
export let caretClosed = "";
export let hyphenOpacity = "opacity-10";
export let regionSummary = "";
export let regionSymbol = "";
export let regionChildren = "";
export let labelledby = "";
setContext("selection", selection);
setContext("multiple", multiple);
setContext("relational", relational);
setContext("padding", padding);
setContext("indent", indent);
setContext("hover", hover);
setContext("rounded", rounded);
setContext("caretOpen", caretOpen);
setContext("caretClosed", caretClosed);
setContext("hyphenOpacity", hyphenOpacity);
setContext("regionSummary", regionSummary);
setContext("regionSymbol", regionSymbol);
setContext("regionChildren", regionChildren);
const dispatch = createEventDispatcher();
function onClick(event) {
  dispatch("click", {
    id: event.detail.id
  });
}
function onToggle(event) {
  dispatch("toggle", {
    id: event.detail.id
  });
}
$:
  classesBase = `${width} ${spacing} ${$$props.class ?? ""}`;
</script>

<div class="tree {classesBase}" data-testid="tree" role="tree" aria-multiselectable={multiple} aria-label={labelledby}>
	{#if nodes && nodes.length > 0}
		<RecursiveTreeViewItem
			{nodes}
			bind:expandedNodes
			bind:disabledNodes
			bind:checkedNodes
			bind:indeterminateNodes
			on:click={onClick}
			on:toggle={onToggle}
		/>
	{/if}
</div>
