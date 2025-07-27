<script lang="ts">
  export {};
  import { Label } from "$lib/components/ui/label";
  import { Slider } from "$lib/components/ui/slider";
  import { modelConfig } from "$lib/store/model-store";
  import { slide } from 'svelte/transition';
  import { cubicOut } from 'svelte/easing';
  import { browser } from '$app/environment';
  import { clickOutside } from '$lib/actions/clickOutside';
  import Tooltip from "$lib/components/ui/tooltip/Tooltip.svelte";

  // Load expanded state from localStorage
  const STORAGE_KEY = 'modelConfigExpanded';
  let isExpanded = false;
  if (browser) {
    const stored = localStorage.getItem(STORAGE_KEY);
    isExpanded = stored ? JSON.parse(stored) : false;
  }

  // Save expanded state
  function toggleExpanded() {
    isExpanded = !isExpanded;
    saveState();
  }

  function saveState() {
    if (browser) {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(isExpanded));
    }
  }

  function handleClickOutside() {
    if (isExpanded) {
      isExpanded = false;
      saveState();
    }
  }

  const settings = [
    { key: 'maxLength', label: 'Maximum Length', min: 1, max: 4000, step: 1, tooltip: "Maximum number of tokens in the response" },
    { key: 'temperature', label: 'Temperature', min: 0, max: 2, step: 0.1, tooltip: "Higher values make output more random, lower values more focused" },
    { key: 'top_p', label: 'Top P', min: 0, max: 1, step: 0.01, tooltip: "Controls diversity via nucleus sampling" },
    { key: 'frequency', label: 'Frequency Penalty', min: 0, max: 1, step: 0.01, tooltip: "Reduces repetition of the same words" },
    { key: 'presence', label: 'Presence Penalty', min: 0, max: 1, step: 0.01, tooltip: "Reduces repetition of similar topics" }
  ] as const;
</script>

<div class="w-full" use:clickOutside={handleClickOutside}>
  <button 
    class="w-full flex items-center py-2 px-2 hover:text-white/90 transition-colors rounded-t"
    on:click={toggleExpanded}
  >
    <span class="text-sm font-semibold">Model Configuration</span>
    <span class="transform transition-transform duration-200 opacity-70 ml-1 text-xs" class:rotate-180={isExpanded}>
      ▼
    </span>
  </button>

  {#if isExpanded}
    <div 
      class="pt-2 px-2 space-y-3"
      transition:slide={{ 
        duration: 200,
        easing: cubicOut,
      }}
    >
      {#each settings as setting}
      <div class="group">
        <div class="flex justify-between items-center mb-0.5">
          <Tooltip text={setting.tooltip} position="right">
            <Label class="text-[10px] text-white/70 cursor-help group-hover:text-white/90 transition-colors">{setting.label}</Label>
          </Tooltip>
          <span class="text-[10px] font-mono text-white/50 group-hover:text-white/70 transition-colors">
            {typeof $modelConfig[setting.key] === 'number' ? $modelConfig[setting.key].toFixed(2) : $modelConfig[setting.key]}
          </span>
        </div>
        <Slider
          bind:value={$modelConfig[setting.key]}
          min={setting.min}
          max={setting.max}
          step={setting.step}
          class="h-3 group-hover:opacity-90 transition-opacity"
        />
      </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  :global(.slider) {
    height: 0.75rem !important;
  }
</style>
