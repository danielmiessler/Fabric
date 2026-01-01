<!-- To access props and events using reference -->
<svelte:options accessors />

<script>import { getContext, createEventDispatcher } from "svelte";
export let group = void 0;
export let name = void 0;
export let value = void 0;
export let checked = false;
export let children = [];
export let spacing = "space-x-4";
export let open = getContext("open");
export let selection = getContext("selection");
export let multiple = getContext("multiple");
export let disabled = getContext("disabled");
export let indeterminate = false;
export let padding = getContext("padding");
export let indent = getContext("indent");
export let hover = getContext("hover");
export let rounded = getContext("rounded");
export let caretOpen = getContext("caretOpen");
export let caretClosed = getContext("caretClosed");
export let hyphenOpacity = getContext("hyphenOpacity");
export let regionSummary = getContext("regionSummary");
export let regionSymbol = getContext("regionSymbol");
export let regionChildren = getContext("regionChildren");
export let hideLead = false;
export let hideChildren = false;
let treeItem;
let childrenDiv;
function onSummaryClick(event) {
  if (disabled)
    event.preventDefault();
}
$:
  if (multiple)
    updateCheckbox(group, indeterminate);
$:
  if (multiple)
    updateGroup(checked, indeterminate);
function updateCheckbox(group2, indeterminate2) {
  if (!Array.isArray(group2))
    return;
  checked = group2.indexOf(value) >= 0;
  dispatch("groupChange", { checked, indeterminate: indeterminate2 });
  dispatch("childChange");
}
function updateGroup(checked2, indeterminate2) {
  if (!Array.isArray(group))
    return;
  const index = group.indexOf(value);
  if (checked2) {
    if (index < 0) {
      group.push(value);
      group = group;
    }
  } else {
    if (index >= 0) {
      group.splice(index, 1);
      group = group;
    }
  }
  if (!indeterminate2) {
    onParentChange();
  }
}
$:
  if (!multiple)
    updateRadio(group);
$:
  if (!multiple)
    updateRadioGroup(checked);
function updateRadio(group2) {
  checked = group2 === value;
  dispatch("groupChange", { checked, indeterminate: false });
  if (group2)
    dispatch("childChange");
}
function updateRadioGroup(checked2) {
  if (checked2 && group !== value)
    group = value;
  else if (!checked2 && group === value)
    group = "";
}
function onChildValueChange() {
  if (multiple) {
    if (!Array.isArray(group))
      return;
    const childrenValues = children.map((c) => c.value);
    const childrenGroup = children[0].group;
    const index = group.indexOf(value);
    if (children.some((c) => c.indeterminate)) {
      indeterminate = true;
      if (index >= 0) {
        group.splice(index, 1);
        group = group;
      }
    } else if (childrenValues.every((c) => Array.isArray(childrenGroup) && childrenGroup.includes(c))) {
      indeterminate = false;
      if (index < 0) {
        group.push(value);
        group = group;
      }
    } else if (childrenValues.some((c) => Array.isArray(childrenGroup) && childrenGroup.includes(c))) {
      indeterminate = true;
      if (index >= 0) {
        group.splice(index, 1);
        group = group;
      }
    } else {
      indeterminate = false;
      if (index >= 0) {
        group.splice(index, 1);
        group = group;
      }
    }
  } else {
    if (group !== value && children.some((c) => c.checked)) {
      group = value;
    } else if (group === value && !children.some((c) => c.checked)) {
      group = "";
    }
  }
  dispatch("childChange");
}
export function onParentChange() {
  if (!multiple || !children || children.length === 0)
    return;
  if (!Array.isArray(group))
    return;
  const index = group.indexOf(value);
  const checkChild = (child) => {
    if (!child || !Array.isArray(child.group))
      return;
    child.indeterminate = false;
    if (child.group.indexOf(child.value) < 0) {
      child.group.push(child.value);
      child.group = child.group;
    }
  };
  const uncheckChild = (child) => {
    if (!child || !Array.isArray(child.group))
      return;
    child.indeterminate = false;
    const childIndex = child.group.indexOf(child.value);
    if (childIndex >= 0) {
      child.group.splice(childIndex, 1);
      child.group = child.group;
    }
  };
  children.forEach((child) => {
    if (!child)
      return;
    index >= 0 ? checkChild(child) : uncheckChild(child);
    child.onParentChange();
  });
}
$:
  if (!multiple && group !== void 0) {
    if (group !== value) {
      children.forEach((child) => {
        if (child)
          child.group = "";
      });
    }
  }
const dispatch = createEventDispatcher();
$:
  dispatch("toggle", { open });
$:
  children.forEach((child) => {
    if (child)
      child.$on("childChange", onChildValueChange);
  });
function onKeyDown(event) {
  function getRootTree() {
    let currentElement = treeItem;
    while (currentElement !== null) {
      if (currentElement.classList.contains("tree"))
        return currentElement;
      currentElement = currentElement.parentElement;
    }
    return void 0;
  }
  let rootTree = void 0;
  let lastVisibleElement = null;
  switch (event.code) {
    case "ArrowRight":
      if (!open)
        open = true;
      else if ($$slots.children && !hideChildren) {
        const child = childrenDiv.querySelector("details>summary");
        if (child)
          child.focus();
      }
      break;
    case "ArrowLeft":
      if (open)
        open = false;
      else {
        const parent = treeItem.parentElement?.parentElement;
        if (parent && parent.tagName === "DETAILS")
          parent.querySelector("summary")?.focus();
      }
      break;
    case "Home":
      event.preventDefault();
      rootTree = getRootTree();
      if (rootTree)
        rootTree?.querySelector("summary")?.focus();
      break;
    case "End":
      event.preventDefault();
      rootTree = getRootTree();
      if (rootTree) {
        const detailsElements = rootTree?.querySelectorAll("details");
        if (!detailsElements)
          return;
        for (let i = detailsElements.length - 1; i >= 0; i--) {
          const details = detailsElements[i];
          if (details.parentElement?.classList?.contains("tree") || details.parentElement?.parentElement?.getAttribute("open") !== null) {
            lastVisibleElement = details;
            break;
          } else if (details.parentElement?.parentElement?.tagName !== "details") {
            lastVisibleElement = details.parentElement.parentElement;
            break;
          }
        }
        if (lastVisibleElement) {
          const summary = lastVisibleElement.querySelector("summary");
          if (summary)
            summary.focus();
        }
      }
      break;
  }
}
const cBase = "";
const cSummary = "list-none [&::-webkit-details-marker]:hidden flex items-center cursor-pointer";
const cSymbol = "fill-current w-3 text-center transition-transform duration-[200ms]";
const cChildren = "";
const cDisabled = "opacity-50 !cursor-not-allowed";
$:
  classesCaretState = open && $$slots.children && !hideChildren ? caretOpen : caretClosed;
$:
  classesDisabled = disabled ? cDisabled : "";
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesSummary = `${cSummary} ${classesDisabled} ${spacing} ${rounded} ${padding} ${hover} ${regionSummary}`;
$:
  classesSymbol = `${cSymbol} ${classesCaret} ${regionSymbol}`;
$:
  classesCaret = `${classesCaretState}`;
$:
  classesHyphen = `${hyphenOpacity}`;
$:
  classesChildren = `${cChildren} ${indent} ${regionChildren}`;
</script>

<details bind:this={treeItem} bind:open class="tree-item {classesBase}" data-testid="tree-item" aria-disabled={disabled}>
	<summary
		class="tree-item-summary {classesSummary}"
		role="treeitem"
		aria-selected={selection ? checked : undefined}
		aria-expanded={$$slots.children ? open : undefined}
		on:click={onSummaryClick}
		on:click
		on:keydown={onKeyDown}
		on:keydown
		on:keyup
	>
		<!-- Symbol -->
		<div class="tree-summary-symbol {classesSymbol}">
			{#if $$slots.children && !hideChildren}
				<!-- SVG Caret -->
				<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
					<path
						d="M201.4 374.6c12.5 12.5 32.8 12.5 45.3 0l160-160c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L224 306.7 86.6 169.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l160 160z"
					/>
				</svg>
			{:else}
				<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="w-3 {classesHyphen}">
					<path d="M432 256c0 17.7-14.3 32-32 32L48 288c-17.7 0-32-14.3-32-32s14.3-32 32-32l352 0c17.7 0 32 14.3 32 32z" />
				</svg>
			{/if}
		</div>

		<!-- Selection -->
		{#if selection && name && group !== undefined}
			{#if multiple}
				<input
					class="checkbox tree-item-checkbox"
					type="checkbox"
					{name}
					{value}
					bind:checked
					bind:indeterminate
					on:change={onParentChange}
				/>
			{:else}
				<input class="radio tree-item-radio" type="radio" bind:group {name} {value} />
			{/if}
		{/if}

		<!-- Slot: Lead -->
		{#if $$slots.lead && !hideLead}
			<div class="tree-item-lead">
				<slot name="lead" />
			</div>
		{/if}
		<!-- Slot: Content -->
		<div class="tree-item-content">
			<slot />
		</div>
	</summary>
	<div bind:this={childrenDiv} class="tree-item-children {classesChildren}" role="group">
		<slot name="children" />
	</div>
</details>
