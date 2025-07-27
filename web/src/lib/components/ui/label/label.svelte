<script lang="ts">
  let className: string = '';
  export { className as class };
</script>

<label
  class="block text-sm font-medium text-gray-700 dark:text-gray-200 {className}"
  {...$$restProps}
>
  <slot />
</label>