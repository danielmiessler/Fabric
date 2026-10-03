<script lang="ts">
  import { onMount } from 'svelte';
  import { Select } from "$lib/components/ui/select";
  import { patterns, patternAPI, systemPrompt, selectedPatternName } from "$lib/store/pattern-store";
  import { get } from 'svelte/store';

  let selectedPreset = $selectedPatternName || "";

  selectedPatternName.subscribe(value => {
    if (value && value !== selectedPreset) {
      console.log('Pattern selected from modal:', value);
      selectedPreset = value;
    }
  });

  // selectPattern handles an empty string, so call it on every change.
  $: {
    console.log('Dropdown selection changed to:', selectedPreset);
    try {
      patternAPI.selectPattern(selectedPreset);

      const currentSystemPrompt = get(systemPrompt);
      const currentPattern = get(selectedPatternName);
      console.log('After dropdown selection - Pattern:', currentPattern);
      console.log('After dropdown selection - System Prompt length:', currentSystemPrompt?.length);
    } catch (error) {
      console.error('Error processing pattern selection:', error);
    }
  }

    onMount(async () => {
      await patternAPI.loadPatterns();
    });
</script>

<div class="min-w-0">
  <Select
    bind:value={selectedPreset}
    class="bg-primary-800/30 border-none hover:bg-primary-800/40 transition-colors"
  >
    <option value="">Load a pattern...</option>
    {#each $patterns as pattern}
      <option value={pattern.Name}>{pattern.Name}</option>
    {/each}
  </Select>
</div>
