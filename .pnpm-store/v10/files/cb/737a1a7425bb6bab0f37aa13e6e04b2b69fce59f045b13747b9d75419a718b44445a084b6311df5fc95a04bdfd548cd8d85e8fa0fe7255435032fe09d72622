<script>import { getContext } from "svelte";
export let group;
export let name;
export let value;
export let title = "";
export let controls = "";
export let regionTab = "";
export let active = getContext("active");
export let hover = getContext("hover");
export let flex = getContext("flex");
export let padding = getContext("padding");
export let rounded = getContext("rounded");
export let spacing = getContext("spacing");
const cBase = "text-center cursor-pointer transition-colors duration-100";
const cInterface = "";
let elemInput;
function onKeyDown(event) {
  if (["Enter", "Space"].includes(event.code)) {
    event.preventDefault();
    elemInput.click();
  } else if (event.code === "ArrowRight") {
    const tabList = elemInput.closest(".tab-list");
    if (!tabList)
      return;
    const tabs = Array.from(tabList.querySelectorAll(".tab"));
    const currTab = elemInput.closest(".tab");
    if (!currTab)
      return;
    const currIndex = tabs.indexOf(currTab);
    const nextIndex = currIndex + 1 >= tabs.length ? 0 : currIndex + 1;
    const nextTab = tabs[nextIndex];
    const nextTabInput = nextTab?.querySelector("input");
    if (nextTab && nextTabInput) {
      nextTabInput.click();
      nextTab.focus();
    }
  } else if (event.code === "ArrowLeft") {
    const tabList = elemInput.closest(".tab-list");
    if (!tabList)
      return;
    const tabs = Array.from(tabList.querySelectorAll(".tab"));
    const currTab = elemInput.closest(".tab");
    if (!currTab)
      return;
    const currIndex = tabs.indexOf(currTab);
    const nextIndex = currIndex - 1 < 0 ? tabs.length - 1 : currIndex - 1;
    const nextTab = tabs[nextIndex];
    const nextTabInput = nextTab?.querySelector("input");
    if (nextTab && nextTabInput) {
      nextTabInput.click();
      nextTab.focus();
    }
  }
}
$:
  selected = value === group;
$:
  classesActive = selected ? active : hover;
$:
  classesBase = `${cBase} ${flex} ${padding} ${rounded} ${classesActive} ${$$props.class ?? ""}`;
$:
  classesInterface = `${cInterface} ${spacing}`;
$:
  classesTab = `${regionTab}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<label class={classesBase} {title}>
	<!-- A11y attributes are not allowed on <label> -->
	<div
		class="tab {classesTab}"
		data-testid="tab"
		role="tab"
		aria-controls={controls}
		aria-selected={selected}
		tabindex={selected ? 0 : -1}
		on:keydown={onKeyDown}
		on:keydown
		on:keyup
		on:keypress
	>
		<!-- NOTE: Don't use `hidden` as it prevents `required` from operating -->
		<div class="h-0 w-0 overflow-hidden">
			<input bind:this={elemInput} type="radio" bind:group {name} {value} {...prunedRestProps()} tabindex="-1" on:click on:change />
		</div>
		<!-- Interface -->
		<div class="tab-interface {classesInterface}">
			{#if $$slots.lead}<div class="tab-lead"><slot name="lead" /></div>{/if}
			<div class="tab-label"><slot /></div>
		</div>
	</div>
</label>
