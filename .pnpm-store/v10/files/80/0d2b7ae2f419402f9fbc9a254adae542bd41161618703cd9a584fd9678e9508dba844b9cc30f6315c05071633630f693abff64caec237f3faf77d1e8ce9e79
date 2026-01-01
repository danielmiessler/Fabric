<script>import { setContext } from "svelte";
export let background = "bg-surface-100-800-token";
export let border = "";
export let width = "w-20";
export let height = "h-full";
export let gap = "gap-0";
export let regionLead = "";
export let regionDefault = "";
export let regionTrail = "";
export let hover = "bg-primary-hover-token";
export let active = "bg-primary-active-token";
export let spacing = "space-y-1";
export let aspectRatio = "aspect-square";
setContext("active", active);
setContext("hover", hover);
setContext("spacing", spacing);
setContext("aspectRatio", aspectRatio);
const cBase = "grid grid-rows-[auto_1fr_auto] overflow-y-auto";
const cRegionLead = "box-border";
const cRegionDefault = "box-border";
const cRegionTrail = "box-border";
$:
  classesBase = `${cBase} ${background} ${border} ${width} ${height} ${gap} ${$$props.class || ""}`;
$:
  classesRegionLead = `${cRegionLead} ${regionLead}`;
$:
  classesRegionDefault = `${cRegionDefault} ${regionDefault}`;
$:
  classesRegionTrail = `${cRegionTrail} ${regionTrail}`;
</script>

<!-- @component A vertical navigation rail component. -->

<div class="app-rail {classesBase}" data-testid="app-rail">
	<!-- Slot: lead -->
	<div class="app-bar-lead {classesRegionLead}"><slot name="lead" /></div>
	<!-- Slot: Default -->
	<div class="app-bar-default {classesRegionDefault}"><slot /></div>
	<!-- Slot: lead -->
	<div class="app-bar-trail {classesRegionTrail}"><slot name="trail" /></div>
</div>
