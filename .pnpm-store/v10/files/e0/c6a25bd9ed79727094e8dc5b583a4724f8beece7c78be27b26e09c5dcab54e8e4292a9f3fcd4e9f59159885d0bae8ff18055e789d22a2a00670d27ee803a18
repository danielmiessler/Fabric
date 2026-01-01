<script>import { afterUpdate } from "svelte";
import { tailwindDefaultColors } from "./settings.js";
export let stops = [{ color: ["neutral", 500], start: 0, end: 100 }];
export let legend = false;
export let spin = false;
export let width = "w-24";
export let hover = "bg-primary-hover-token";
export let digits = 0;
export let regionCaption = "";
export let regionCone = "";
export let regionLegend = "";
let cone;
let generatedLegendList;
const cBase = "flex flex-col items-center space-y-4 w-";
const cCaption = "text-center";
const cCone = "block aspect-square rounded-full";
const cLegend = "text-sm w-full";
const cSwatch = "block aspect-square bg-black w-5 rounded-full mr-2";
function setColorValue(color) {
  if (typeof color === "string")
    return color;
  const colorSet = tailwindDefaultColors.find((c) => c.label === color[0]);
  return colorSet?.shades[color[1]].hex;
}
function genConicGradient() {
  let d = stops.map((v) => `${setColorValue(v.color)} ${v.start}% ${v.end}%`);
  cone = `conic-gradient(${d.join(", ")})`;
}
function genLegend() {
  if (!legend)
    return;
  generatedLegendList = stops.map((v) => {
    return {
      label: v.label,
      color: setColorValue(v.color),
      value: (v.end - v.start).toFixed(digits)
    };
  });
}
afterUpdate(() => {
  genConicGradient();
  genLegend();
});
$:
  classesBase = `${cBase} ${$$props.class ?? ""}`;
$:
  classesCaption = `${cCaption} ${regionCaption}`;
$:
  classesCone = `${cCone} ${width} ${regionCone}`;
$:
  classesLegend = `${cLegend} ${regionLegend}`;
</script>

<figure class="conic-gradient {classesBase}" data-testid="conic-gradient">
	<!-- Label -->
	{#if $$slots.default}
		<figcaption class="conic-caption {classesCaption}"><slot /></figcaption>
	{/if}
	<!-- Conic Gradient -->
	{#if cone}
		<div class="conic-cone {classesCone}" class:animate-spin={spin} style:background={cone}></div>
	{/if}
	<!-- Legend -->
	{#if legend && generatedLegendList}
		<ul class="conic-list list {classesLegend}">
			{#each generatedLegendList as { color, label, value }}
				<li class="conic-item {hover}">
					<span class="conic-swatch {cSwatch}" style:background={color}></span>
					<span class="conic-label flex-auto">{label}</span>
					<strong class="conic-value">{value}%</strong>
				</li>
			{/each}
		</ul>
	{/if}
</figure>
