<script>import { getContext } from "svelte";
export let group;
export let name;
export let value;
export let title = "";
export let label = "";
export let rounded = getContext("rounded");
export let padding = getContext("padding");
export let active = getContext("active");
export let hover = getContext("hover");
export let color = getContext("color");
export let fill = getContext("fill");
export let regionLabel = getContext("regionLabel");
const cBase = "flex-auto";
const cWrapper = "text-base text-center cursor-pointer";
const cDisabled = "opacity-50 cursor-not-allowed";
let elemInput;
function onKeyDown(event) {
  if (["Enter", "Space"].includes(event.code)) {
    event.preventDefault();
    elemInput.click();
  }
}
$:
  checked = value === group;
$:
  classesActive = checked ? `${active} ${color} ${fill}` : hover;
$:
  classesDisabled = $$props.disabled ? cDisabled : "";
$:
  classsBase = `${cBase}`;
$:
  classesWrapper = `${cWrapper} ${padding} ${rounded} ${classesActive} ${classesDisabled} ${$$props.class ?? ""}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<label class="radio-label {classsBase} {regionLabel}">
	<!-- A11y attributes are not allowed on <label> -->
	<div
		class="radio-item {classesWrapper}"
		data-testid="radio-item"
		role="radio"
		aria-checked={checked}
		aria-label={label}
		tabindex="0"
		{title}
		on:keydown={onKeyDown}
		on:keydown
		on:keyup
		on:keypress
	>
		<!-- NOTE: Don't use `hidden` as it prevents `required` from operating -->
		<div class="h-0 w-0 overflow-hidden">
			<input bind:this={elemInput} type="radio" bind:group {name} {value} {...prunedRestProps()} tabindex="-1" on:click on:change />
		</div>
		<slot />
	</div>
</label>
