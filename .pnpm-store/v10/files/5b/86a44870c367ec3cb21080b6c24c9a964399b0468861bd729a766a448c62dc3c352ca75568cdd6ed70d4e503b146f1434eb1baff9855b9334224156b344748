<script context="module">import { fly, fade } from "svelte/transition";
import { prefersReducedMotionStore } from "../../index.js";
import { dynamicTransition } from "../../internal/transitions.js";
</script>

<script generics="TransitionIn extends Transition = FlyTransition, TransitionOut extends Transition = FlyTransition">import { createEventDispatcher } from "svelte";
const dispatch = createEventDispatcher();
import { focusTrap } from "../../actions/FocusTrap/focusTrap.js";
import { getModalStore } from "./stores.js";
export let components = {};
export let position = "items-center";
export let background = "bg-surface-100-800-token";
export let width = "w-modal";
export let height = "h-auto";
export let padding = "p-4";
export let spacing = "space-y-4";
export let rounded = "rounded-container-token";
export let shadow = "shadow-xl";
export let zIndex = "z-[999]";
export let buttonNeutral = "variant-ghost-surface";
export let buttonPositive = "variant-filled";
export let buttonTextCancel = "Cancel";
export let buttonTextConfirm = "Confirm";
export let buttonTextSubmit = "Submit";
export let regionBackdrop = "";
export let regionHeader = "text-2xl font-bold";
export let regionBody = "max-h-[200px] overflow-hidden";
export let regionFooter = "flex justify-end space-x-2";
export let transitions = !$prefersReducedMotionStore;
export let transitionIn = fly;
export let transitionInParams = { duration: 150, opacity: 0, x: 0, y: 100 };
export let transitionOut = fly;
export let transitionOutParams = { duration: 150, opacity: 0, x: 0, y: 100 };
const cBackdrop = "fixed top-0 left-0 right-0 bottom-0 bg-surface-backdrop-token p-4";
const cTransitionLayer = "w-full h-fit min-h-full overflow-y-auto flex justify-center";
const cModal = "block overflow-y-auto";
const cModalImage = "w-full h-auto";
let promptValue;
const buttonTextDefaults = {
  buttonTextCancel,
  buttonTextConfirm,
  buttonTextSubmit
};
let currentComponent;
let registeredInteractionWithBackdrop = false;
let modalElement;
let windowHeight;
let backdropOverflow = "overflow-y-hidden";
const modalStore = getModalStore();
$:
  if ($modalStore.length)
    handleModals($modalStore);
function handleModals(modals) {
  if (modals[0].type === "prompt")
    promptValue = modals[0].value;
  buttonTextCancel = modals[0].buttonTextCancel || buttonTextDefaults.buttonTextCancel;
  buttonTextConfirm = modals[0].buttonTextConfirm || buttonTextDefaults.buttonTextConfirm;
  buttonTextSubmit = modals[0].buttonTextSubmit || buttonTextDefaults.buttonTextSubmit;
  currentComponent = typeof modals[0].component === "string" ? components[modals[0].component] : modals[0].component;
}
function onModalHeightChange(modal) {
  let modalHeight = modal?.clientHeight;
  if (!modalHeight)
    modalHeight = modal?.firstChild?.clientHeight;
  if (!modalHeight)
    return;
  if (modalHeight > windowHeight) {
    backdropOverflow = "overflow-y-auto";
  } else {
    backdropOverflow = "overflow-y-hidden";
  }
}
$:
  onModalHeightChange(modalElement);
function onBackdropInteractionBegin(event) {
  if (!(event.target instanceof Element))
    return;
  const classList = event.target.classList;
  if (classList.contains("modal-backdrop") || classList.contains("modal-transition")) {
    registeredInteractionWithBackdrop = true;
  }
}
function onBackdropInteractionEnd(event) {
  if (!(event.target instanceof Element))
    return;
  const classList = event.target.classList;
  if ((classList.contains("modal-backdrop") || classList.contains("modal-transition")) && registeredInteractionWithBackdrop) {
    if ($modalStore[0].response)
      $modalStore[0].response(void 0);
    modalStore.close();
    dispatch("backdrop", event);
  }
  registeredInteractionWithBackdrop = false;
}
function onClose() {
  if ($modalStore[0].response)
    $modalStore[0].response(false);
  modalStore.close();
}
function onConfirm() {
  if ($modalStore[0].response)
    $modalStore[0].response(true);
  modalStore.close();
}
function onPromptSubmit(event) {
  event.preventDefault();
  if ($modalStore[0].response) {
    if ($modalStore[0].valueAttr !== void 0 && "type" in $modalStore[0].valueAttr && $modalStore[0].valueAttr.type === "number")
      $modalStore[0].response(parseInt(promptValue));
    else
      $modalStore[0].response(promptValue);
  }
  modalStore.close();
}
function onKeyDown(event) {
  if (!$modalStore.length)
    return;
  if (event.code === "Escape")
    onClose();
}
$:
  cPosition = $modalStore[0]?.position ?? position;
$:
  classesBackdrop = `${cBackdrop} ${regionBackdrop} ${zIndex} ${$$props.class ?? ""} ${$modalStore[0]?.backdropClasses ?? ""}`;
$:
  classesTransitionLayer = `${cTransitionLayer} ${cPosition ?? ""}`;
$:
  classesModal = `${cModal} ${background} ${width} ${height} ${padding} ${spacing} ${rounded} ${shadow} ${$modalStore[0]?.modalClasses ?? ""}`;
$:
  parent = {
    position,
    // ---
    background,
    width,
    height,
    padding,
    spacing,
    rounded,
    shadow,
    // ---
    buttonNeutral,
    buttonPositive,
    buttonTextCancel,
    buttonTextConfirm,
    buttonTextSubmit,
    // ---
    regionBackdrop,
    regionHeader,
    regionBody,
    regionFooter,
    // ---
    onClose
  };
</script>

<svelte:window bind:innerHeight={windowHeight} on:keydown={onKeyDown} />

{#if $modalStore.length > 0}
	{#key $modalStore}
		<!-- Backdrop -->
		<!-- FIXME: resolve a11y warnings -->
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		<div
			class="modal-backdrop {classesBackdrop} {backdropOverflow}"
			data-testid="modal-backdrop"
			on:mousedown={onBackdropInteractionBegin}
			on:mouseup={onBackdropInteractionEnd}
			on:touchstart|passive
			on:touchend|passive
			transition:dynamicTransition|global={{ transition: fade, params: { duration: 150 }, enabled: transitions }}
			use:focusTrap={true}
		>
			<!-- Transition Layer -->
			<div
				class="modal-transition {classesTransitionLayer}"
				in:dynamicTransition|global={{ transition: transitionIn, params: transitionInParams, enabled: transitions }}
				out:dynamicTransition|global={{ transition: transitionOut, params: transitionOutParams, enabled: transitions }}
			>
				{#if $modalStore[0].type !== 'component'}
					<!-- Modal: Presets -->
					<div
						class="modal {classesModal}"
						bind:this={modalElement}
						data-testid="modal"
						role="dialog"
						aria-modal="true"
						aria-label={$modalStore[0].title ?? ''}
					>
						<!-- Header -->
						{#if $modalStore[0]?.title}
							<header class="modal-header {regionHeader}">{@html $modalStore[0].title}</header>
						{/if}
						<!-- Body -->
						{#if $modalStore[0]?.body}
							<article class="modal-body {regionBody}">{@html $modalStore[0].body}</article>
						{/if}
						<!-- Image -->
						{#if $modalStore[0]?.image && typeof $modalStore[0]?.image === 'string'}
							<img class="modal-image {cModalImage}" src={$modalStore[0]?.image} alt="Modal" />
						{/if}
						<!-- Type -->
						{#if $modalStore[0].type === 'alert'}
							<!-- Template: Alert -->
							<footer class="modal-footer {regionFooter}">
								<button type="button" class="btn {buttonNeutral}" on:click={onClose}>{buttonTextCancel}</button>
							</footer>
						{:else if $modalStore[0].type === 'confirm'}
							<!-- Template: Confirm -->
							<footer class="modal-footer {regionFooter}">
								<button type="button" class="btn {buttonNeutral}" on:click={onClose}>{buttonTextCancel}</button>
								<button type="button" class="btn {buttonPositive}" on:click={onConfirm}>{buttonTextConfirm}</button>
							</footer>
						{:else if $modalStore[0].type === 'prompt'}
							<!-- Template: Prompt -->
							<form class="space-y-4" on:submit={onPromptSubmit}>
								<input class="modal-prompt-input input" name="prompt" type="text" bind:value={promptValue} {...$modalStore[0].valueAttr} />
								<footer class="modal-footer {regionFooter}">
									<button type="button" class="btn {buttonNeutral}" on:click={onClose}>{buttonTextCancel}</button>
									<button type="submit" class="btn {buttonPositive}">{buttonTextSubmit}</button>
								</footer>
							</form>
						{/if}
					</div>
				{:else}
					<!-- Modal: Components -->
					<!-- Note: keep `contents` class to allow widths from children -->
					<div
						bind:this={modalElement}
						class="modal contents {$modalStore[0]?.modalClasses ?? ''}"
						data-testid="modal-component"
						role="dialog"
						aria-modal="true"
						aria-label={$modalStore[0].title ?? ''}
					>
						{#if currentComponent?.slot}
							<svelte:component this={currentComponent?.ref} {...currentComponent?.props} {parent}>
								{@html currentComponent?.slot}
							</svelte:component>
						{:else}
							<svelte:component this={currentComponent?.ref} {...currentComponent?.props} {parent} />
						{/if}
					</div>
				{/if}
			</div>
		</div>
	{/key}
{/if}
