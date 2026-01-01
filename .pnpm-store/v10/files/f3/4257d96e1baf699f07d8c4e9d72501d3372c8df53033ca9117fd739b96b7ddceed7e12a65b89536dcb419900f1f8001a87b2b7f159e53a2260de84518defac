<script>import { createEventDispatcher } from "svelte";
const dispatch = createEventDispatcher();
export let name;
export let checked = false;
export let size = "md";
export let background = "bg-surface-400 dark:bg-surface-700";
export let active = "bg-surface-900 dark:bg-surface-300";
export let border = "";
export let rounded = "rounded-full";
export let label = "";
const cBase = "inline-block";
const cLabel = "unstyled flex items-center";
const cTrack = "flex transition-all duration-[200ms] cursor-pointer";
const cThumb = "w-[50%] h-full scale-[0.8] transition-all duration-[200ms] shadow";
let trackSize;
switch (size) {
  case "sm":
    trackSize = "w-12 h-6";
    break;
  case "lg":
    trackSize = "w-20 h-10";
    break;
  default:
    trackSize = "w-16 h-8";
}
function onKeyDown(event) {
  if (["Enter", "Space"].includes(event.code)) {
    event.preventDefault();
    dispatch("keyup", event);
    const inputElem = event.currentTarget.firstChild;
    inputElem.click();
  }
}
$:
  cTrackActive = checked ? active : `${background} cursor-pointer`;
$:
  cThumbBackground = checked ? "bg-white/75" : "bg-white";
$:
  cThumbPos = checked ? "translate-x-full" : "";
$:
  classesDisabled = $$props.disabled === true ? "opacity-50" : "hover:brightness-[105%] dark:hover:brightness-110 cursor-pointer";
$:
  classesBase = `${cBase} ${rounded} ${classesDisabled} ${$$props.class ?? ""}`;
$:
  classesLabel = `${cLabel}`;
$:
  classesTrack = `${cTrack} ${border} ${rounded} ${trackSize} ${cTrackActive}`;
$:
  classesThumb = `${cThumb} ${rounded} ${cThumbBackground} ${cThumbPos}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<div
	id={label}
	class="slide-toggle {classesBase}"
	data-testid="slide-toggle"
	on:keydown={onKeyDown}
	role="switch"
	aria-label={label}
	aria-checked={checked}
	tabindex="0"
>
	<label class="slide-toggle-label {classesLabel}">
		<!-- Hidden Input -->
		<input
			type="checkbox"
			class="slide-toggle-input hidden"
			bind:checked
			{name}
			on:click
			on:keydown
			on:keyup
			on:keypress
			on:mouseover
			on:change
			on:focus
			on:blur
			{...prunedRestProps()}
			disabled={$$props.disabled}
		/>
		<!-- Slider Track/Thumb -->
		<div class="slide-toggle-track {classesTrack}" class:cursor-not-allowed={$$props.disabled}>
			<div class="slide-toggle-thumb {classesThumb}" class:cursor-not-allowed={$$props.disabled}></div>
		</div>
		<!-- Label -->
		{#if $$slots.default}<div class="slide-toggle-text ml-3"><slot /></div>{/if}
	</label>
</div>
