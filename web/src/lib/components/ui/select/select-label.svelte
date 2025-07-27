<script>
  import { Select as SelectPrimitive } from "bits-ui";
  import { cn } from "$lib/utils/utils.ts";
  let className = undefined;
  export { className as class };
</script>

<SelectPrimitive.Label class={cn("px-2 py-1.5 text-sm variant-filled-secondary font-semibold", className)} {...$$restProps}>
	<slot />
</SelectPrimitive.Label>
