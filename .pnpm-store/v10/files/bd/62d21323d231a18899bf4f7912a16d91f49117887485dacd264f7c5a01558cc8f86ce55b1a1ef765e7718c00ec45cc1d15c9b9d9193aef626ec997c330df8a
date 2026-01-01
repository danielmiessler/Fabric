<script>import { createEventDispatcher, getContext } from "svelte";
import { dynamicTransition } from "../../internal/transitions.js";
const dispatch = createEventDispatcher();
export let open = false;
export let id = String(Math.random());
const cBase = "";
const cControl = "text-start w-full flex items-center space-x-4";
const cControlIcons = "fill-current w-3 transition-transform duration-[200ms]";
const cPanel = "";
export let autocollapse = getContext("autocollapse");
export let active = getContext("active");
export let disabled = getContext("disabled");
export let padding = getContext("padding");
export let hover = getContext("hover");
export let rounded = getContext("rounded");
export let caretOpen = getContext("caretOpen");
export let caretClosed = getContext("caretClosed");
export let regionControl = getContext("regionControl");
export let regionPanel = getContext("regionPanel");
export let regionCaret = getContext("regionCaret");
export let transitions = getContext("transitions");
export let transitionIn = getContext("transitionIn");
export let transitionInParams = getContext("transitionInParams");
export let transitionOut = getContext("transitionOut");
export let transitionOutParams = getContext("transitionOutParams");
const svgCaretIcon = `
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
			<path d="M201.4 374.6c12.5 12.5 32.8 12.5 45.3 0l160-160c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L224 306.7 86.6 169.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l160 160z" />
		</svg>`;
function setActive(event) {
  if (autocollapse === true) {
    active.set(id);
  } else {
    open = !open;
  }
  onToggle(event);
}
function onToggle(event) {
  const currentOpenState = autocollapse ? $active === id : open;
  dispatch("toggle", {
    event,
    id,
    panelId: `accordion-panel-${id}`,
    open: currentOpenState,
    autocollapse
  });
}
if (autocollapse && open)
  setActive();
$:
  if (open && autocollapse)
    setActive();
$:
  openState = autocollapse ? $active === id : open;
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesControl = `${cControl} ${padding} ${hover} ${rounded} ${regionControl}`;
$:
  classesCaretState = openState ? caretOpen : caretClosed;
$:
  classesControlCaret = `${cControlIcons} ${regionCaret} ${classesCaretState}`;
$:
  classesControlIcons = `${cControlIcons} ${regionCaret}`;
$:
  classesPanel = `${cPanel} ${padding} ${rounded} ${regionPanel}`;
</script>

<!-- @component The Accordion child element. -->

<div class="accordion-item {classesBase}" data-testid="accordion-item">
	<!-- Control -->
	<button
		type="button"
		class="accordion-control {classesControl}"
		{id}
		on:click={setActive}
		on:click
		on:keydown
		on:keyup
		on:keypress
		aria-expanded={openState}
		aria-controls="accordion-panel-{id}"
		{disabled}
	>
		<!-- Lead -->
		{#if $$slots.lead}
			<div class="accordion-lead">
				<slot name="lead" />
			</div>
		{/if}
		<!-- Summary -->
		<div class="accordion-summary flex-1">
			<slot name="summary">(summary)</slot>
		</div>
		<!-- Icons -->
		{#if $$slots.iconClosed || $$slots.iconOpen}
			<!-- Custom -->
			<!-- If a custom icon is provided, do not use rotation -->
			<div class="accordion-summary-icons {classesControlIcons}">
				{#if openState}
					<slot name="iconClosed">{@html svgCaretIcon}</slot>
				{:else}
					<slot name="iconOpen">{@html svgCaretIcon}</slot>
				{/if}
			</div>
		{:else}
			<!-- SVG Caret -->
			<div class="accordion-summary-caret {classesControlCaret}">{@html svgCaretIcon}</div>
		{/if}
	</button>
	<!-- Panel -->
	{#if openState}
		<div
			class="accordion-panel {classesPanel}"
			id="accordion-panel-{id}"
			in:dynamicTransition|local={{ transition: transitionIn, params: transitionInParams, enabled: transitions }}
			out:dynamicTransition|local={{ transition: transitionOut, params: transitionOutParams, enabled: transitions }}
			role="region"
			aria-hidden={!openState}
			aria-labelledby={id}
		>
			<slot name="content">(content)</slot>
		</div>
	{/if}
</div>
