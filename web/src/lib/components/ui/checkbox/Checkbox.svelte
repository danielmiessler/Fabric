<script lang="ts">
  import { cn } from "$lib/utils/utils";
  
  export let checked: boolean = false;
  export let id: string | undefined = undefined;
  export let disabled: boolean = false;
  let className: string | undefined = undefined;
  export { className as class };
</script>

<div class="flex items-center">
  <input
    type="checkbox"
    {id}
    bind:checked
    {disabled}
    class={cn(
      "h-4 w-4 rounded border-gray-300 text-primary-600 focus:ring-primary-500 disabled:cursor-not-allowed disabled:opacity-50",
      className
    )}
  />
</div>
