<script>export let initials = "";
export let fill = "fill-token";
export let fontSize = 150;
export let src = "";
export let fallback = "";
export let action = () => {
};
export let actionParams = "";
export let background = "bg-surface-400-500-token";
export let width = "w-16";
export let border = "";
export let rounded = "rounded-full";
export let shadow = "";
export let cursor = "";
let cBase = "flex aspect-square text-surface-50 font-semibold justify-center items-center overflow-hidden isolate";
let cImage = "w-full object-cover";
$:
  classesBase = `${cBase} ${background} ${width} ${border} ${rounded} ${shadow} ${cursor} ${$$props.class ?? ""}`;
function prunedRestProps() {
  delete $$restProps.class;
  return $$restProps;
}
</script>

<figure class="avatar {classesBase}" data-testid="avatar" {...prunedRestProps()}>
	{#if src || fallback}
		<img
			class="avatar-image {cImage}"
			style={$$props.style ?? ''}
			{src}
			alt={$$props.alt || ''}
			use:action={actionParams}
			on:error={() => (src = fallback)}
		/>
	{:else if initials}
		<svg class="avatar-initials w-full h-full" viewBox="0 0 512 512">
			<text
				x="50%"
				y="50%"
				dominant-baseline="central"
				text-anchor="middle"
				font-weight="bold"
				font-size={fontSize}
				class="avatar-text {fill}"
			>
				{String(initials).substring(0, 2).toUpperCase()}
			</text>
		</svg>
	{:else}
		<slot />
	{/if}
</figure>
