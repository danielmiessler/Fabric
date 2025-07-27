<script lang="ts">
  import { cn } from "$lib/utils/utils";
  import { buttonVariants } from "./index.js";

  let className: string | undefined = undefined;
  export let variant: string = "";
  export let size: string = "default";
  export { className as class };

  $: classes = cn(
    buttonVariants.base,
    buttonVariants.variants.variant[variant as keyof typeof buttonVariants.variants.variant],
    buttonVariants.variants.size[size as keyof typeof buttonVariants.variants.size],
    className
  );
</script>

<button
	class={classes}
	type="button"
	on:click
	on:keydown
	{...$$restProps}
>
	<slot />
</button>
