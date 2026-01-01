<script>import { createEventDispatcher } from "svelte";
import { BROWSER } from "esm-env";
const dispatch = createEventDispatcher();
import { prefersReducedMotionStore } from "../../index.js";
import { focusTrap } from "../../actions/FocusTrap/focusTrap.js";
import { getDrawerStore } from "./stores.js";
import { fade, fly } from "svelte/transition";
import { dynamicTransition } from "../../internal/transitions.js";
import { cubicIn } from "svelte/easing";
export let position = "left";
export let bgDrawer = "bg-surface-100-800-token";
export let border = "";
export let rounded = "";
export let shadow = "shadow-xl";
export let width = "";
export let height = "";
export let bgBackdrop = "bg-surface-backdrop-token";
export let blur = "";
export let padding = "";
export let zIndex = "z-40";
export let regionBackdrop = "";
export let regionDrawer = "";
export let labelledby = "";
export let describedby = "";
export let duration = 200;
export let transitions = !$prefersReducedMotionStore;
export let opacityTransition = true;
const presets = {
  top: { alignment: "items-start", width: "w-full", height: "h-[50%]", rounded: "rounded-bl-container-token rounded-br-container-token" },
  bottom: { alignment: "items-end", width: "w-full", height: " h-[50%]", rounded: "rounded-tl-container-token rounded-tr-container-token" },
  left: { alignment: "justify-start", width: "w-[90%]", height: "h-full", rounded: "rounded-tr-container-token rounded-br-container-token" },
  right: { alignment: "justify-end", width: "w-[90%]", height: "h-full", rounded: "rounded-tl-container-token rounded-bl-container-token" }
};
let elemBackdrop;
let elemDrawer;
let anim = { x: 0, y: 0 };
const drawerStore = getDrawerStore();
const cBackdrop = "fixed top-0 left-0 right-0 bottom-0 flex";
const cDrawer = "overflow-y-auto transition-transform";
const propDefaults = {
  position,
  bgBackdrop,
  blur,
  padding,
  bgDrawer,
  border,
  rounded,
  shadow,
  width,
  height,
  opacityTransition,
  regionBackdrop,
  regionDrawer,
  labelledby,
  describedby,
  duration
};
function applyPropSettings(settings) {
  position = settings.position || propDefaults.position;
  bgBackdrop = settings.bgBackdrop || propDefaults.bgBackdrop;
  blur = settings.blur || propDefaults.blur;
  padding = settings.padding || propDefaults.padding;
  bgDrawer = settings.bgDrawer || propDefaults.bgDrawer;
  border = settings.border || propDefaults.border;
  rounded = settings.rounded || propDefaults.rounded;
  shadow = settings.shadow || propDefaults.shadow;
  width = settings.width || propDefaults.width;
  height = settings.height || propDefaults.height;
  regionBackdrop = settings.regionBackdrop || propDefaults.regionBackdrop;
  regionDrawer = settings.regionDrawer || propDefaults.regionDrawer;
  labelledby = settings.labelledby || propDefaults.labelledby;
  describedby = settings.describedby || propDefaults.describedby;
  opacityTransition = settings.opacityTransition || propDefaults.opacityTransition;
  duration = settings.duration || propDefaults.duration;
}
function applyAnimationSettings() {
  if (!BROWSER)
    return;
  switch (position) {
    case "top":
      anim = { x: 0, y: -window.innerWidth };
      break;
    case "bottom":
      anim = { x: 0, y: window.innerWidth };
      break;
    case "left":
      anim = { x: -window.innerHeight, y: 0 };
      break;
    case "right":
      anim = { x: window.innerHeight, y: 0 };
      break;
    default:
      console.error("Error: unknown position property value.");
      break;
  }
}
function onDrawerInteraction(event) {
  if (event.target === elemBackdrop) {
    drawerStore.close();
    dispatch("backdrop", event);
  } else {
    dispatch("drawer", event);
  }
}
function onKeydownWindow(event) {
  if (!$drawerStore)
    return;
  if (event.code === "Escape")
    drawerStore.close();
}
drawerStore.subscribe((settings) => {
  if (settings.open !== true)
    return;
  applyPropSettings(settings);
  applyAnimationSettings();
});
$:
  classesPosition = presets[position].alignment;
$:
  classesWidth = width ? width : presets[position].width;
$:
  classesHeight = height ? height : presets[position].height;
$:
  classesRounded = rounded ? rounded : presets[position].rounded;
$:
  classesBackdrop = `${cBackdrop} ${bgBackdrop} ${padding} ${blur} ${classesPosition} ${regionBackdrop} ${zIndex} ${$$props.class ?? ""}`;
$:
  classesDrawer = `${cDrawer} ${bgDrawer} ${border} ${rounded} ${shadow} ${classesWidth} ${classesHeight} ${classesRounded} ${regionDrawer}`;
</script>

<svelte:window on:keydown={onKeydownWindow} />

{#if $drawerStore.open === true}
	<!-- Backdrop -->
	<!-- FIXME: resolve a11y warnings -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div
		bind:this={elemBackdrop}
		class="drawer-backdrop {classesBackdrop}"
		data-testid="drawer-backdrop"
		on:mousedown={onDrawerInteraction}
		on:touchstart|passive
		on:touchend|passive
		on:keypress
		in:dynamicTransition|local={{
			transition: fade,
			params: { duration },
			enabled: transitions && opacityTransition
		}}
		out:dynamicTransition|local={{
			transition: fade,
			params: { duration },
			enabled: transitions && opacityTransition
		}}
		use:focusTrap={true}
	>
		<!-- Drawer -->
		<!-- separate In/Out so anim values update -->
		<div
			bind:this={elemDrawer}
			class="drawer {classesDrawer}"
			data-testid="drawer"
			role="dialog"
			aria-modal="true"
			aria-labelledby={labelledby}
			aria-describedby={describedby}
			in:dynamicTransition|local={{
				transition: fly,
				params: { x: anim.x, y: anim.y, duration, opacity: opacityTransition ? undefined : 1 },
				enabled: transitions
			}}
			out:dynamicTransition|local={{
				transition: fly,
				params: { x: anim.x, y: anim.y, duration, opacity: opacityTransition ? undefined : 1, easing: cubicIn },
				enabled: transitions
			}}
		>
			<!-- Slot: Default -->
			<slot />
		</div>
	</div>
{/if}
