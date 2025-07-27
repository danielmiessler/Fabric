<script>
  import { cn } from "$lib/utils/utils";
  let className = undefined;
  export let value = undefined;
  export { className as class };
  export let readonly = undefined;
</script>

<textarea
	class={cn(
		"border-input placeholder:text-muted-foreground focus-visible:ring-ring flex min-h-[60px] w-full rounded-lg border bg-transparent px-3 py-2 text-sm shadow-lg focus-visible:outline-none focus-visible:ring-1 disabled:cursor-not-allowed disabled:opacity-50",
		className
	)}
	bind:value
	{readonly}
	on:blur
	on:change
	on:click
	on:focus
	on:keydown
	on:keypress
	on:keyup
	on:mouseover
	on:mouseenter
	on:mouseleave
	on:paste
	on:input
	{...$$restProps}
></textarea>
