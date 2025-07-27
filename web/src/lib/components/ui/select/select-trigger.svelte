<script>
  import { Select as SelectPrimitive } from "bits-ui";
  import { ChevronDown } from "lucide-svelte";
  import { cn } from "$lib/utils";
  import { getContext } from "svelte";

  let className = undefined;
  export { className as class };
</script>

<SelectPrimitive.Trigger
	class={cn(
		"flex h-9 w-full items-center justify-between rounded-md border border-input bg-transparent px-3 py-2 text-sm shadow-sm ring-offset-background placeholder:text-muted-foreground focus:outline-none focus:ring-1 focus:ring-ring disabled:cursor-not-allowed disabled:opacity-50",
		className
	)}
	{...$$restProps}
>
	<slot />
	<ChevronDown class="h-4 w-4 opacity-50" />
</SelectPrimitive.Trigger>
