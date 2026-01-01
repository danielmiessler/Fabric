<script>export let value = void 0;
export let min = 0;
export let max = 100;
export let height = "h-2";
export let rounded = "rounded-token";
export let transition = "transition-[width]";
export let animIndeterminate = "anim-indeterminate";
export let meter = "bg-surface-900-50-token";
export let track = "bg-surface-200-700-token";
export let labelledby = "";
const cTrack = "w-full overflow-hidden";
const cMeter = "h-full";
$:
  fillPercent = value ? 100 * (value - min) / (max - min) : 0;
$:
  indeterminate = value === void 0 || value < 0;
$:
  classesIndeterminate = indeterminate ? animIndeterminate : "";
$:
  classesTrack = `${cTrack} ${track} ${height} ${rounded} ${$$props.class ?? ""}`;
$:
  classesMeter = `${cMeter} ${meter} ${rounded} ${classesIndeterminate} ${transition}`;
</script>

<!-- Track -->
<div
	class="progress-bar {classesTrack}"
	data-testid="progress-bar"
	role="progressbar"
	aria-labelledby={labelledby}
	aria-valuenow={value}
	aria-valuemin={min}
	aria-valuemax={max - min}
>
	<!-- Meter -->
	<div class="progress-bar-meter {classesMeter}" style:width="{indeterminate ? 100 : fillPercent}%"></div>
</div>

<style>
	.anim-indeterminate {
		transform-origin: 0% 50%;
		animation: anim-indeterminate 2s infinite linear;
	}
	/* prettier-ignore */
	@keyframes anim-indeterminate {
		0% { transform: translateX(0) scaleX(0); }
		40% { transform: translateX(0) scaleX(0.4); }
		100% { transform: translateX(100%) scaleX(0.5); }
	}
</style>
