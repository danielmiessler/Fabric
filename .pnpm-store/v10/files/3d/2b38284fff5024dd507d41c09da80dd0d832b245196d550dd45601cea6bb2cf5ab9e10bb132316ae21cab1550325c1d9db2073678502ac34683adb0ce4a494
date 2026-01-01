<script context="module">import { fly, scale } from "svelte/transition";
import { prefersReducedMotionStore } from "../../index.js";
import { dynamicTransition } from "../../internal/transitions.js";
</script>

<script

	generics="ListTransitionIn extends Transition = FlyTransition, ListTransitionOut extends Transition = FlyTransition, ChipTransitionIn extends Transition = ScaleTransition, ChipTransitionOut extends Transition = ScaleTransition"
>import { createEventDispatcher, onMount } from "svelte";
import { flip } from "svelte/animate";
const dispatch = createEventDispatcher();
export let input = "";
export let name;
export let value = [];
export let whitelist = [];
export let max = -1;
export let minlength = -1;
export let maxlength = -1;
export let allowUpperCase = false;
export let allowDuplicates = false;
export let validation = () => true;
export let duration = 150;
export let required = false;
export let chips = "variant-filled";
export let invalid = "input-error";
export let padding = "p-2";
export let rounded = "rounded-container-token";
export let regionChipWrapper = "";
export let regionChipList = "";
export let regionInput = "";
export let label = "Chips select";
export let transitions = !$prefersReducedMotionStore;
export let listTransitionIn = fly;
export let listTransitionInParams = { duration: 150, opacity: 0, y: -20 };
export let listTransitionOut = fly;
export let listTransitionOutParams = { duration: 150, opacity: 0, y: -20 };
export let chipTransitionIn = scale;
export let chipTransitionInParams = { duration: 150, opacity: 0 };
export let chipTransitionOut = scale;
export let chipTransitionOutParams = { duration: 150, opacity: 0 };
const cBase = "textarea cursor-pointer";
const cChipWrapper = "space-y-4";
const cChipList = "flex flex-wrap gap-2";
const cInputField = "unstyled bg-transparent border-0 !ring-0 p-0 w-full";
let inputValid = true;
let chipValues = value?.map((val) => {
  return { val, id: Math.random() };
}) || [];
function resetFormHandler() {
  value = [];
}
let selectElement;
onMount(() => {
  if (!selectElement.form)
    return;
  const externalForm = selectElement.form;
  externalForm.addEventListener("reset", resetFormHandler);
  return () => {
    externalForm.removeEventListener("reset", resetFormHandler);
  };
});
function validateCustom(chip) {
  return validation === void 0 || validation(chip);
}
function validateCount() {
  return max === -1 || value.length < max;
}
function validateLength(chip) {
  return (minlength === -1 || chip.length >= minlength) && (maxlength === -1 || chip.length <= maxlength);
}
function validateWhiteList(chip) {
  return whitelist.length === 0 || whitelist.includes(chip);
}
function validateDuplicates(chip) {
  return allowDuplicates || !value.includes(chip);
}
function validate(chip = "") {
  if (!chip && !input)
    return false;
  chip = chip !== "" ? chip.trim() : input.trim();
  return validateCustom(chip) && validateCount() && validateLength(chip) && validateWhiteList(chip) && validateDuplicates(chip);
}
function addChipCommon(chip) {
  chip = allowUpperCase ? chip : chip.toLowerCase();
  value.push(chip);
  value = value;
  chipValues.push({ val: chip, id: Math.random() });
  chipValues = chipValues;
}
function removeChipCommon(chip) {
  let chipIndex = value.indexOf(chip);
  value.splice(chipIndex, 1);
  value = value;
  chipValues.splice(chipIndex, 1);
  chipValues = chipValues;
}
function onKeyHandler(event) {
  if (event.key !== "Enter")
    return;
  event.preventDefault();
  inputValid = validate();
  if (inputValid === false) {
    dispatch("invalid", { event, input });
    return;
  }
  addChipCommon(input);
  dispatch("add", { event, chipIndex: value.length - 1, chipValue: input });
  input = "";
}
function removeChipInternally(event, chipIndex, chipValue) {
  if ($$restProps.disabled)
    return;
  removeChipCommon(chipValue);
  dispatch("remove", { event, chipIndex, chipValue });
}
export function addChip(chip) {
  inputValid = validate(chip);
  if (inputValid === false) {
    dispatch("invalidManually", { input: chip });
    return;
  }
  addChipCommon(chip);
  dispatch("addManually", { chipIndex: value.length - 1, chipValue: chip });
}
export function removeChip(chip) {
  if ($$restProps.disabled)
    return;
  removeChipCommon(chip);
  dispatch("removeManually", { chipValue: chip });
}
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
$:
  classesInvalid = inputValid === false ? invalid : "";
$:
  classesBase = `${cBase} ${padding} ${rounded} ${$$props.class ?? ""} ${classesInvalid}`;
$:
  classesChipWrapper = `${cChipWrapper} ${regionChipWrapper}`;
$:
  classesChipList = `${cChipList} ${regionChipList}`;
$:
  classesInput = `${cInputField} ${regionInput}`;
$:
  chipValues = value?.map((val, i) => {
    if (chipValues[i]?.val === val)
      return chipValues[i];
    return { id: Math.random(), val };
  }) || [];
</script>

<div class="input-chip {classesBase}" class:opacity-50={$$restProps.disabled}>
	<!-- NOTE: Don't use `hidden` as it prevents `required` from operating -->
	<div class="h-0 overflow-hidden">
		<select bind:this={selectElement} bind:value {name} multiple {required} aria-label={label} tabindex="-1">
			<!-- NOTE: options are required! -->
			{#each value as option}<option value={option}>{option}</option>{/each}
		</select>
	</div>
	<!-- Chip Wrapper -->
	<div class="input-chip-wrapper {classesChipWrapper}">
		<!-- Input Field -->
		<input
			type="text"
			bind:value={input}
			placeholder={$$restProps.placeholder ?? 'Enter values...'}
			class="input-chip-field {classesInput}"
			on:keydown={onKeyHandler}
			on:input
			on:focus
			on:blur
			disabled={$$restProps.disabled}
			{...prunedRestProps()}
		/>
		<!-- Chip List -->
		{#if chipValues.length}
			<div
				class="input-chip-list {classesChipList}"
				in:dynamicTransition|local={{ transition: listTransitionIn, params: listTransitionInParams, enabled: transitions }}
				out:dynamicTransition|local={{ transition: listTransitionOut, params: listTransitionOutParams, enabled: transitions }}
			>
				{#each chipValues as { id, val }, i (id)}
					<!-- Wrapping div required for FLIP animation -->
					<div animate:flip={{ duration }}>
						<button
							type="button"
							class="chip {chips}"
							on:click={(e) => removeChipInternally(e, i, val)}
							on:click
							on:keypress
							on:keydown
							on:keyup
							in:dynamicTransition|local={{ transition: chipTransitionIn, params: chipTransitionInParams, enabled: transitions }}
							out:dynamicTransition|local={{ transition: chipTransitionOut, params: chipTransitionOutParams, enabled: transitions }}
						>
							<span>{val}</span>
							<span>✕</span>
						</button>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
