<script context="module">import { slide } from "svelte/transition";
import { prefersReducedMotionStore } from "../../index.js";
</script>

<script generics="TransitionIn extends Transition = SlideTransition, TransitionOut extends Transition = SlideTransition">import { writable } from "svelte/store";
import { setContext } from "svelte";
export let autocollapse = false;
export let width = "w-full";
export let spacing = "space-y-1";
export let disabled = false;
export let padding = "py-2 px-4";
export let hover = "hover:bg-primary-hover-token";
export let rounded = "rounded-container-token";
export let caretOpen = "rotate-180";
export let caretClosed = "";
export let regionControl = "";
export let regionPanel = "space-y-4";
export let regionCaret = "";
export let transitions = !$prefersReducedMotionStore;
export let transitionIn = slide;
export let transitionInParams = { duration: 200 };
export let transitionOut = slide;
export let transitionOutParams = { duration: 200 };
const active = writable(null);
setContext("active", active);
setContext("autocollapse", autocollapse);
setContext("disabled", disabled);
setContext("padding", padding);
setContext("hover", hover);
setContext("rounded", rounded);
setContext("caretOpen", caretOpen);
setContext("caretClosed", caretClosed);
setContext("regionControl", regionControl);
setContext("regionPanel", regionPanel);
setContext("regionCaret", regionCaret);
setContext("transitions", transitions);
setContext("transitionIn", transitionIn);
setContext("transitionInParams", transitionInParams);
setContext("transitionOut", transitionOut);
setContext("transitionOutParams", transitionOutParams);
$:
  classesBase = `${width} ${spacing} ${$$props.class ?? ""}`;
</script>

<!-- @component The Accordion parent element. -->

<div class="accordion {classesBase}" data-testid="accordion">
	<slot />
</div>
