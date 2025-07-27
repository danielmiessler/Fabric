<script lang="ts">
  import { onMount } from 'svelte';
  import { Select } from "$lib/components/ui/select";
  import { modelConfig, availableModels, loadAvailableModels } from "$lib/store/model-store";

  onMount(async () => {
    await loadAvailableModels();
  });
</script>

<div class="min-w-0">
  <Select
    bind:value={$modelConfig.model}
    class="bg-primary-800/30 border-none hover:bg-primary-800/40 transition-colors"
  >
    <option value="">Default Model</option>
    {#each $availableModels as model (model.name)}
      <option value={model.name}>{model.vendor} - {model.name}</option>
    {/each}
  </Select>
</div>
