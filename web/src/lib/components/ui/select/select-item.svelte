<script>
  import { cn } from "$lib/utils";

  let className = undefined;
  export let value;
  export let disabled = false;
  export { className as class };
</script>

<option
	{value}
	{disabled}
	class={cn("relative flex w-full cursor-default select-none py-1.5 pl-2 pr-2 text-sm outline-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50", className)}
>
	<slot />
</option>
