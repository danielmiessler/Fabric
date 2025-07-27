<script lang="ts">
  import { Select as SelectPrimitive } from "bits-ui";
  import { cn } from "$lib/utils";

  type $$Props = SelectPrimitive.ValueProps;

  let className: $$Props["class"] = undefined;
  export { className as class };
  export let placeholder: $$Props["placeholder"] = undefined;
</script>

<SelectPrimitive.Value
	class={cn("text-sm", className)}
	{placeholder}
	{...$$restProps}
>
	<slot />
</SelectPrimitive.Value>
