<script lang="ts">
  import { cn } from "$lib/utils/utils";

  export let value: any = undefined;
  export let disabled = false;
  let className: string | undefined = undefined;
  export { className as class };
</script>

<div class="relative">
	<select
		{disabled}
		bind:value
		class={cn(
			"select flex h-8 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-xs font-medium text-white/90 shadow-lg ring-offset-background placeholder:text-muted-primary focus:outline-none focus:ring-1 focus:ring-ring hover:text-white disabled:cursor-not-allowed disabled:opacity-50 appearance-none transition-colors",
			className
		)}
		{...$$restProps}
	>
		<slot />
	</select>
</div>
