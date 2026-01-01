<!-- Reference: https://css-tricks.com/building-progress-ring-quickly/ -->
<script>import { afterUpdate } from "svelte";
export let value = void 0;
export let stroke = 40;
export let font = 56;
export let strokeLinecap = "butt";
export let transition = "transition-[stroke-dashoffset]";
export let width = "w-36";
export let meter = "stroke-surface-900 dark:stroke-surface-50";
export let track = "stroke-surface-500/30";
export let fill = "fill-token";
export let labelledby = "";
const cBase = "progress-radial relative overflow-hidden";
const cBaseTrack = "fill-transparent";
const cBaseMeter = "fill-transparent -rotate-90 origin-[50%_50%]";
const baseSize = 512;
const radius = baseSize / 2 - stroke / 2;
let circumference = radius;
let dashoffset;
function setProgress(percent) {
  circumference = radius * 2 * Math.PI;
  dashoffset = circumference - percent / 100 * circumference;
}
setProgress(0);
afterUpdate(() => {
  setProgress(value === void 0 ? 25 : value);
});
$:
  classesBase = `${cBase} ${width} ${$$props.class ?? ""}`;
</script>

<figure
	class="progress-radial {classesBase}"
	data-testid="progress-radial"
	role="meter"
	aria-labelledby={labelledby}
	aria-valuenow={value || 0}
	aria-valuetext={value ? `${value}%` : 'Indeterminate Spinner'}
	aria-valuemin={0}
	aria-valuemax={100}
>
	<!-- Draw SVG -->
	<svg viewBox="0 0 {baseSize} {baseSize}" class="rounded-full" class:animate-spin={value === undefined}>
		<!-- Track -->
		<circle class="progress-radial-track {cBaseTrack} {track}" stroke-width={stroke} r={radius} cx="50%" cy="50%" />

		<!-- Meter -->
		<circle
			class="progress-radial-meter {cBaseMeter} {meter} {transition}"
			stroke-width={stroke}
			r={radius}
			cx="50%"
			cy="50%"
			style:stroke-dasharray="{circumference}
			{circumference}"
			style:stroke-dashoffset={dashoffset}
			stroke-linecap={strokeLinecap}
		/>

		<!-- Center Text -->
		{#if value != undefined && value >= 0 && $$slots.default}
			<text
				x="50%"
				y="50%"
				text-anchor="middle"
				dominant-baseline="middle"
				font-weight="bold"
				font-size={font}
				class="progress-radial-text {fill}"><slot /></text
			>
		{/if}
	</svg>
</figure>
