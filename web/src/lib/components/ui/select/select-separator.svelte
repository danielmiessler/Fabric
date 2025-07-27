<script>
  import { Select as SelectPrimitive } from "bits-ui";
  import { cn } from "$lib/utils/utils.ts";
  let className = undefined;
  export { className as class };
</script>

<SelectPrimitive.Separator class={cn("bg-muted -mx-1 my-1 h-px", className)} {...$$restProps} />
