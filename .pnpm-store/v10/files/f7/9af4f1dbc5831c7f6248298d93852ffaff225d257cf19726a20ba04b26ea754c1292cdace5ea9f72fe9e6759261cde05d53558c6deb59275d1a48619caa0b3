<script>import { getContext } from "svelte";
export let group;
export let name;
export let value;
export let disabled = getContext("disabled");
export let multiple = getContext("multiple");
export let rounded = getContext("rounded");
export let active = getContext("active");
export let hover = getContext("hover");
export let padding = getContext("padding");
export let regionLead = getContext("regionLead");
export let regionDefault = getContext("regionDefault");
export let regionTrail = getContext("regionTrail");
const cBase = "cursor-pointer -outline-offset-[3px]";
const cDisabled = "opacity-50 !cursor-default";
const cLabel = "flex items-center space-x-4";
let checked;
let elemInput;
function areDeeplyEqual(param1, param2) {
  if (param1 === param2)
    return true;
  if (!(param1 instanceof Object) || !(param2 instanceof Object))
    return false;
  const keys1 = Object.keys(param1);
  const keys2 = Object.keys(param2);
  if (keys1.length !== keys2.length)
    return false;
  for (const key of keys1) {
    const value1 = param1[key];
    const value2 = param2[key];
    if (!areDeeplyEqual(value1, value2))
      return false;
  }
  return true;
}
$:
  if (multiple)
    updateCheckbox(group);
$:
  if (multiple)
    updateGroup(checked);
function updateCheckbox(group2) {
  checked = group2.indexOf(value) >= 0;
}
function updateGroup(checked2) {
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
}
function onKeyDown(event) {
  if (["Enter", "Space"].includes(event.code)) {
    event.preventDefault();
    elemInput.click();
  }
}
const cRegionLead = "";
const cRegionDefault = "flex-1";
const cRegionTrail = "";
$:
  selected = multiple ? group.some((groupVal) => areDeeplyEqual(value, groupVal)) : areDeeplyEqual(group, value);
$:
  classesActive = selected ? active : !disabled ? hover : "";
$:
  classesDisabled = disabled ? cDisabled : "";
$:
  classesBase = `${cBase} ${classesDisabled} ${rounded} ${padding} ${classesActive} ${$$props.class ?? ""}`;
$:
  classesLabel = `${cLabel}`;
$:
  classesRegionLead = `${cRegionLead} ${regionLead}`;
$:
  classesRegionDefault = `${cRegionDefault} ${regionDefault}`;
$:
  classesRegionTrail = `${cRegionTrail} ${regionTrail}`;
</script>

<label>
	<!-- A11y attributes are not allowed on <label> -->
	<div
		class="listbox-item {classesBase}"
		data-testid="listbox-item"
		role="option"
		aria-selected={selected}
		tabindex="0"
		on:keydown={onKeyDown}
		on:keydown
		on:keyup
		on:keypress
	>
		<!-- NOTE: Don't use `hidden` as it prevents `required` from operating -->
		<div class="h-0 w-0 overflow-hidden">
			{#if multiple}
				<input {disabled} bind:this={elemInput} type="checkbox" {name} {value} bind:checked tabindex="-1" on:click on:change />
			{:else}
				<input {disabled} bind:this={elemInput} type="radio" bind:group {name} {value} tabindex="-1" on:click on:change />
			{/if}
		</div>
		<!-- <slot /> -->
		<div class="listbox-label {classesLabel}">
			<!-- Slot: Lead -->
			{#if $$slots.lead}<div class="listbox-label-lead {classesRegionLead}"><slot name="lead" /></div>{/if}
			<!-- Slot: Default -->
			<div class="listbox-label-content {classesRegionDefault}"><slot /></div>
			<!-- Slot: Trail -->
			{#if $$slots.trail}<div class="listbox-label-trail {classesRegionTrail}"><slot name="trail" /></div>{/if}
		</div>
	</div>
</label>
