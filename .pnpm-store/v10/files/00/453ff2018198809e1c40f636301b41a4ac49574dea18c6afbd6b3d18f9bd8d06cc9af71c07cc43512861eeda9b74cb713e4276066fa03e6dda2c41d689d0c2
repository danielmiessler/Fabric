<script>export let background = "bg-surface-100-800-token";
export let border = "";
export let padding = "p-4";
export let shadow = "";
export let spacing = "space-y-4";
export let gridColumns = "grid-cols-[auto_1fr_auto]";
export let gap = "gap-4";
export let regionRowMain = "";
export let regionRowHeadline = "";
export let slotLead = "";
export let slotDefault = "";
export let slotTrail = "";
export let label = "";
export let labelledby = "";
const cBase = "flex flex-col";
const cRowMain = "grid items-center";
const cRowHeadline = "";
const cSlotLead = "flex-none flex justify-between items-center";
const cSlotDefault = "flex-auto";
const cSlotTrail = "flex-none flex items-center space-x-4";
$:
  classesBase = `${cBase} ${background} ${border} ${spacing} ${padding} ${shadow} ${$$props.class ?? ""}`;
$:
  classesRowMain = `${cRowMain} ${gridColumns} ${gap} ${regionRowMain}`;
$:
  classesRowHeadline = `${cRowHeadline} ${regionRowHeadline}`;
$:
  classesSlotLead = `${cSlotLead} ${slotLead}`;
$:
  classesSlotDefault = `${cSlotDefault} ${slotDefault}`;
$:
  classesSlotTrail = `${cSlotTrail} ${slotTrail}`;
</script>

<div class="app-bar {classesBase}" data-testid="app-bar" role="toolbar" aria-label={label} aria-labelledby={labelledby}>
	<!-- Row: Main -->
	<div class="app-bar-row-main {classesRowMain}">
		<!-- Slot: lead -->
		{#if $$slots.lead}
			<div class="app-bar-slot-lead {classesSlotLead}"><slot name="lead" /></div>
		{/if}
		<!-- Slot: default -->
		<div class="app-bar-slot-default {classesSlotDefault}"><slot /></div>
		<!-- Slot: trail -->
		{#if $$slots.trail}
			<div class="app-bar-slot-trail {classesSlotTrail}"><slot name="trail" /></div>
		{/if}
	</div>
	<!-- Row: Headline -->
	{#if $$slots.headline}
		<div class="app-bar-row-headline {classesRowHeadline}"><slot name="headline" /></div>
	{/if}
</div>
