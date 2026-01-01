<script context="module">import { fade } from "svelte/transition";
import { prefersReducedMotionStore } from "../../index.js";
</script>

<script generics="TransitionIn extends Transition = FadeTransition, TransitionOut extends Transition = FadeTransition">import { createEventDispatcher, setContext } from "svelte";
import { writable } from "svelte/store";
import { dynamicTransition } from "../../internal/transitions.js";
const dispatch = createEventDispatcher();
export let gap = "gap-4";
export let stepTerm = "Step";
export let badge = "variant-filled-surface";
export let active = "variant-filled";
export let border = "border-surface-400-500-token";
export let start = 0;
export let justify = "justify-between";
export let buttonBack = "variant-ghost";
export let buttonBackType = "button";
export let buttonBackLabel = "&larr; Back";
export let buttonNext = "variant-filled";
export let buttonNextType = "button";
export let buttonNextLabel = "Next &rarr;";
export let buttonComplete = "variant-filled-primary";
export let buttonCompleteType = "button";
export let buttonCompleteLabel = "Complete";
export let regionHeader = "";
export let regionContent = "";
export let transitions = !$prefersReducedMotionStore;
export let transitionIn = fade;
export let transitionInParams = { duration: 100 };
export let transitionOut = fade;
export let transitionOutParams = { duration: 100 };
let state = writable({ current: start, total: 0 });
async function onNext(locked, stepIndex) {
  await new Promise((resolve) => setTimeout(resolve));
  if (locked)
    return;
  $state.current++;
  dispatch("next", { step: stepIndex, state: $state });
  dispatch("step", { step: stepIndex, state: $state });
}
function onBack(stepIndex) {
  $state.current--;
  dispatch("back", { step: stepIndex, state: $state });
  dispatch("step", { step: stepIndex, state: $state });
}
function onComplete(stepIndex) {
  dispatch("complete", { step: stepIndex, state: $state });
}
setContext("state", state);
setContext("stepTerm", stepTerm);
setContext("gap", gap);
setContext("justify", justify);
setContext("onNext", onNext);
setContext("onBack", onBack);
setContext("onComplete", onComplete);
setContext("buttonBack", buttonBack);
setContext("buttonBackType", buttonBackType);
setContext("buttonBackLabel", buttonBackLabel);
setContext("buttonNext", buttonNext);
setContext("buttonNextType", buttonNextType);
setContext("buttonNextLabel", buttonNextLabel);
setContext("buttonComplete", buttonComplete);
setContext("buttonCompleteType", buttonCompleteType);
setContext("buttonCompleteLabel", buttonCompleteLabel);
setContext("transitions", transitions);
setContext("transitionIn", transitionIn);
setContext("transitionInParams", transitionInParams);
setContext("transitionOut", transitionOut);
setContext("transitionOutParams", transitionOutParams);
const cBase = "space-y-4";
const cHeader = "flex items-center border-t mt-[15px]";
const cHeaderStep = "-mt-[15px] transition-all duration-300";
const cContent = "";
$:
  isActive = (step) => step === $state.current;
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesHeader = `${cHeader} ${border} ${gap} ${regionHeader}`;
$:
  classesHeaderStep = `${cHeaderStep}`;
$:
  classesBadge = (step) => isActive(step) ? active : badge;
$:
  classesContent = `${cContent} ${regionContent}`;
</script>

<div class="stepper {classesBase}" data-testid="stepper">
	<!-- Header -->
	{#if $state.total}
		<header
			class="stepper-header {classesHeader}"
			in:dynamicTransition|local={{ transition: transitionIn, params: transitionInParams, enabled: transitions }}
			out:dynamicTransition|local={{ transition: transitionOut, params: transitionOutParams, enabled: transitions }}
		>
			{#each Array.from(Array($state.total).keys()) as step}
				<div class="stepper-header-step {classesHeaderStep}" class:flex-1={isActive(step)}>
					<span class="badge {classesBadge(step)}">{isActive(step) ? `${stepTerm} ${step + 1}` : step + 1}</span>
				</div>
			{/each}
		</header>
	{/if}
	<!-- Content -->
	<div class="stepper-content {classesContent}">
		<slot />
	</div>
</div>
