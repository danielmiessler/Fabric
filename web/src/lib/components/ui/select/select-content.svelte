<script>
  import { Select as SelectPrimitive } from "bits-ui";
  import { cn } from "$lib/utils";
  import { getContext } from "svelte";

  let className = undefined;
  export let sideOffset = 4;
  export let position = "popper";
</script>

<div class={cn("relative z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-lg data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
		position === "popper" &&
			"data-[side=bottom]:translate-y-1 data-[side=left]:-translate-x-1 data-[side=right]:translate-x-1 data-[side=top]:-translate-y-1",
		className
	)}
>
	<SelectPrimitive.Content {position} {sideOffset} {...$$restProps}>
		<div
			class={cn(
				"p-1",
				position === "popper" &&
					"h-[var(--radix-select-trigger-height)] w-full min-w-[var(--radix-select-trigger-width)]"
			)}
		>
			<slot />
		</div>
	</SelectPrimitive.Content>
</div>
