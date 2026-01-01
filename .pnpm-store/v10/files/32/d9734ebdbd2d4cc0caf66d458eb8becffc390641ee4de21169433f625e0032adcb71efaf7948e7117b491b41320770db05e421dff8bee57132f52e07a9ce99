<!-- Reference: https://dribbble.com/shots/16221169-Figma-Material-Ui-components-Steppers-and-sliders -->
<script>import { getContext, onDestroy } from "svelte";
import { dynamicTransition } from "../../internal/transitions.js";
export let locked = false;
export let regionHeader = "";
export let regionContent = "";
export let regionNavigation = "";
export let state = getContext("state");
export let stepTerm = getContext("stepTerm");
export let gap = getContext("gap");
export let justify = getContext("justify");
export let onNext = getContext("onNext");
export let onBack = getContext("onBack");
export let onComplete = getContext("onComplete");
export let buttonBack = getContext("buttonBack");
export let buttonBackType = getContext("buttonBackType");
export let buttonBackLabel = getContext("buttonBackLabel");
export let buttonNext = getContext("buttonNext");
export let buttonNextType = getContext("buttonNextType");
export let buttonNextLabel = getContext("buttonNextLabel");
export let buttonComplete = getContext("buttonComplete");
export let buttonCompleteType = getContext("buttonCompleteType");
export let buttonCompleteLabel = getContext("buttonCompleteLabel");
export let transitions = getContext("transitions");
export let transitionIn = getContext("transitionIn");
export let transitionInParams = getContext("transitionInParams");
export let transitionOut = getContext("transitionOut");
export let transitionOutParams = getContext("transitionOutParams");
const stepIndex = $state.total;
$state.total++;
const cBase = "space-y-4";
const cHeader = "text-2xl font-bold";
const cContent = "space-y-4";
const cNavigation = "flex";
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesHeader = `${cHeader} ${regionHeader}`;
$:
  classesContent = `${cContent} ${regionContent}`;
$:
  classesNavigation = `${cNavigation} ${justify} ${gap} ${regionNavigation}`;
onDestroy(() => {
  $state.total--;
});
</script>

{#if stepIndex === $state.current}
	<div class="step {classesBase}" data-testid="step">
		<!-- Slot: Header -->
		<header class="step-header {classesHeader}">
			<slot name="header">{stepTerm} {stepIndex + 1}</slot>
		</header>
		<!-- Slot: Default -->
		<div class="step-content {classesContent}">
			<slot>({stepTerm} {stepIndex + 1} Content)</slot>
		</div>
		<!-- Navigation -->
		{#if $state.total > 1}
			<div
				class="step-navigation {classesNavigation}"
				in:dynamicTransition|local={{ transition: transitionIn, params: transitionInParams, enabled: transitions }}
				out:dynamicTransition|local={{ transition: transitionOut, params: transitionOutParams, enabled: transitions }}
			>
				{#if stepIndex === 0 && $$slots.navigation}
					<!-- Slot: Navigation -->
					<div class="step-navigation-slot">
						<slot name="navigation" />
					</div>
				{:else}
					<!-- Button: Back -->
					<button type={buttonBackType} class="btn {buttonBack}" on:click={() => onBack(stepIndex)} disabled={$state.current === 0}>
						{@html buttonBackLabel}
					</button>
				{/if}
				{#if stepIndex < $state.total - 1}
					<!-- Button: Next -->
					<button type={buttonNextType} class="btn {buttonNext}" on:click={() => onNext(locked, stepIndex)} disabled={locked}>
						{#if locked}
							<svg class="w-3 aspect-square fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
								<path
									d="M144 144v48H304V144c0-44.2-35.8-80-80-80s-80 35.8-80 80zM80 192V144C80 64.5 144.5 0 224 0s144 64.5 144 144v48h16c35.3 0 64 28.7 64 64V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V256c0-35.3 28.7-64 64-64H80z"
								/>
							</svg>
						{/if}
						<span>{@html buttonNextLabel}</span>
					</button>
				{:else}
					<!-- Button: Complete -->
					<button type={buttonCompleteType} class="btn {buttonComplete}" on:click={() => onComplete(stepIndex)} disabled={locked}>
						{@html buttonCompleteLabel}
					</button>
				{/if}
			</div>
		{/if}
	</div>
{/if}
