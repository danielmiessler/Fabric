<script lang="ts">
  import { onMount } from 'svelte';
  import { Select } from "$lib/components/ui/select";
  import { modelConfig, loadAvailableModels, selectedVendor, vendorNames, filteredModels } from "$lib/store/model-store";

  onMount(async () => {
    await loadAvailableModels();
  });
</script>

<div class="min-w-0 flex flex-col gap-2">
  <Select
    bind:value={$selectedVendor}
    class="bg-primary-800/30 border-none hover:bg-primary-800/40 transition-colors"
  >
    <option value="">All Vendors</option>
    {#each $vendorNames as vendor (vendor)}
      <option value={vendor}>{vendor}</option>
    {/each}
  </Select>
  <Select
    bind:value={$modelConfig.model}
    class="bg-primary-800/30 border-none hover:bg-primary-800/40 transition-colors"
  >
    <option value="">Default Model</option>
    {#each $filteredModels as model (`${model.vendor}:${model.name}`)}
      <option value={model.name}>{model.vendor} - {model.name}</option>
    {/each}
  </Select>
</div>
