<script context="module">import { fly } from "svelte/transition";
import { prefersReducedMotionStore } from "../../index.js";
import { dynamicTransition } from "../../internal/transitions.js";
</script>

<script generics="TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition">import { flip } from "svelte/animate";
import { getToastStore } from "./stores.js";
const toastStore = getToastStore();
export let position = "b";
export let max = 3;
export let background = "variant-filled-secondary";
export let width = "max-w-[640px]";
export let color = "";
export let padding = "p-4";
export let spacing = "space-x-4";
export let rounded = "rounded-container-token";
export let shadow = "shadow-lg";
export let zIndex = "z-[888]";
export let buttonAction = "btn variant-filled";
export let buttonDismiss = "btn-icon btn-icon-sm variant-filled";
export let buttonDismissLabel = "\u2715";
export let transitions = !$prefersReducedMotionStore;
export let transitionIn = fly;
export let transitionInParams = { duration: 250 };
export let transitionOut = fly;
export let transitionOutParams = { duration: 250 };
const cWrapper = "flex fixed top-0 left-0 right-0 bottom-0 pointer-events-none";
const cSnackbar = "flex flex-col gap-y-2";
const cToast = "flex justify-between items-center pointer-events-auto";
const cToastActions = "flex items-center space-x-2";
let cPosition;
let cAlign;
let animAxis = { x: 0, y: 0 };
switch (position) {
  case "t":
    cPosition = "justify-center items-start";
    cAlign = "items-center";
    animAxis = { x: 0, y: -100 };
    break;
  case "b":
    cPosition = "justify-center items-end";
    cAlign = "items-center";
    animAxis = { x: 0, y: 100 };
    break;
  case "l":
    cPosition = "justify-start items-center";
    cAlign = "items-start";
    animAxis = { x: -100, y: 0 };
    break;
  case "r":
    cPosition = "justify-end items-center";
    cAlign = "items-end";
    animAxis = { x: 100, y: 0 };
    break;
  case "tl":
    cPosition = "justify-start items-start";
    cAlign = "items-start";
    animAxis = { x: -100, y: 0 };
    break;
  case "tr":
    cPosition = "justify-end items-start";
    cAlign = "items-end";
    animAxis = { x: 100, y: 0 };
    break;
  case "bl":
    cPosition = "justify-start items-end";
    cAlign = "items-start";
    animAxis = { x: -100, y: 0 };
    break;
  case "br":
    cPosition = "justify-end items-end";
    cAlign = "items-end";
    animAxis = { x: 100, y: 0 };
    break;
}
function onAction(index) {
  $toastStore[index]?.action?.response();
  toastStore.close($toastStore[index].id);
}
function onMouseEnter(index) {
  if ($toastStore[index]?.hoverable) {
    toastStore.freeze(index);
    classesSnackbar += " scale-[105%]";
  }
}
function onMouseLeave(index) {
  if ($toastStore[index]?.hoverable) {
    toastStore.unfreeze(index);
    classesSnackbar = classesSnackbar.replace(" scale-[105%]", "");
  }
}
let wrapperVisible = false;
$:
  classesWrapper = `${cWrapper} ${cPosition} ${zIndex} ${$$props.class || ""}`;
$:
  classesSnackbar = `${cSnackbar} ${cAlign} ${padding}`;
$:
  classesToast = `${cToast} ${width} ${color} ${padding} ${spacing} ${rounded} ${shadow}`;
$:
  filteredToasts = Array.from($toastStore).slice(0, max);
$:
  if (filteredToasts.length) {
    wrapperVisible = true;
  }
</script>

{#if filteredToasts.length > 0 || wrapperVisible}
	<!-- Wrapper -->
	<div class="snackbar-wrapper {classesWrapper}" data-testid="snackbar-wrapper">
		<!-- List -->
		<div class="snackbar {classesSnackbar}">
			{#each filteredToasts as t, i (t)}
				<div
					animate:flip={{ duration: transitions ? 250 : 0 }}
					in:dynamicTransition|global={{
						transition: transitionIn,
						params: { x: animAxis.x, y: animAxis.y, ...transitionInParams },
						enabled: transitions
					}}
					out:dynamicTransition|global={{
						transition: transitionOut,
						params: { x: animAxis.x, y: animAxis.y, ...transitionOutParams },
						enabled: transitions
					}}
					on:outroend={() => {
						const outroFinishedForLastToastOnQueue = filteredToasts.length === 0;
						if (outroFinishedForLastToastOnQueue) wrapperVisible = false;
					}}
					on:mouseenter={() => onMouseEnter(i)}
					on:mouseleave={() => onMouseLeave(i)}
					role={t.hideDismiss ? 'alert' : 'alertdialog'}
					aria-live="polite"
				>
					<!-- Toast -->
					<div class="toast {classesToast} {t.background ?? background} {t.classes ?? ''}" data-testid="toast">
						<div class="text-base">{@html t.message}</div>
						{#if t.action || !t.hideDismiss}
							<div class="toast-actions {cToastActions}">
								{#if t.action}<button class={buttonAction} on:click={() => onAction(i)}>{@html t.action.label}</button>{/if}
								{#if !t.hideDismiss}<button class={buttonDismiss} aria-label="Dismiss toast" on:click={() => toastStore.close(t.id)}
										>{buttonDismissLabel}</button
									>{/if}
							</div>
						{/if}
					</div>
				</div>
			{/each}
		</div>
	</div>
{/if}
