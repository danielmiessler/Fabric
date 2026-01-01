<script>import { setContext } from "svelte";
export let selection = false;
export let multiple = false;
export let width = "w-full";
export let spacing = "space-y-1";
export let open = false;
export let disabled = false;
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
export function expandAll() {
  const detailsElements = tree.querySelectorAll("details.tree-item");
  detailsElements.forEach((details) => {
    if (!details.open) {
      const summary = details.querySelector("summary.tree-item-summary");
      if (summary)
        summary.click();
    }
  });
}
export function collapseAll() {
  const detailsElements = tree.querySelectorAll("details.tree-item");
  detailsElements.forEach((details) => {
    if (details.open) {
      const summary = details.querySelector("summary.tree-item-summary");
      if (summary)
        summary.click();
    }
  });
}
setContext("open", open);
setContext("selection", selection);
setContext("multiple", multiple);
setContext("disabled", disabled);
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
$:
  classesBase = `${width} ${spacing} ${$$props.class ?? ""}`;
let tree;
</script>

<div
	bind:this={tree}
	class="tree {classesBase}"
	data-testid="tree"
	role="tree"
	aria-multiselectable={multiple}
	aria-label={labelledby}
	aria-disabled={disabled}
>
	<slot />
</div>
