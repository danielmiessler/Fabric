<script lang="ts">
  import ChatInput from "./ChatInput.svelte";
  import ChatMessages from "./ChatMessages.svelte";
  import ModelConfig from "./ModelConfig.svelte";
  import DropdownGroup from "./DropdownGroup.svelte";
  import NoteDrawer from "$lib/components/ui/noteDrawer/NoteDrawer.svelte";
  import { Button } from "$lib/components/ui/button";
  import { Input } from "$lib/components/ui/input";
  import { Label } from "$lib/components/ui/label";
  import { Checkbox } from "$lib/components/ui/checkbox";
  import Tooltip from "$lib/components/ui/tooltip/Tooltip.svelte";
  import { Textarea } from "$lib/components/ui/textarea";
  import { obsidianSettings } from "$lib/store/obsidian-store";
  import { featureFlags } from "$lib/config/features";
  import { drawerStore } from '$lib/store/drawer-store';
  import { systemPrompt, selectedPatternName } from "$lib/store/pattern-store";
  import { onMount } from "svelte";

  function openDrawer() {
    drawerStore.open({});
  }

  // Column widths, in percent
  let leftColumnWidth = 50;
  let rightColumnWidth = 50;
  let isDragging = false;
  
  // Message input heights, in percent
  const DEFAULT_INPUT_HEIGHT = 30; // Percent of the left column height
  const MAX_INPUT_HEIGHT = DEFAULT_INPUT_HEIGHT * 2;
  const MIN_SYSTEM_INSTRUCTIONS_HEIGHT = 20;
  let messageInputHeight = DEFAULT_INPUT_HEIGHT;
  let systemInstructionsHeight = 100 - DEFAULT_INPUT_HEIGHT;
  let isVerticalDragging = false;
  let initialMouseY = 0;
  let initialInputHeight = 0;
  
  function startResize(e: MouseEvent | KeyboardEvent) {
    isDragging = true;
    e.preventDefault();
    
    window.addEventListener('mousemove', handleResize);
    window.addEventListener('mouseup', stopResize);
  }
  
  function handleKeyDown(e: KeyboardEvent) {
    if (e.key === 'Enter' || e.key === ' ') {
      startResize(e);
    }
  }
  
  function handleResize(e: MouseEvent) {
    if (!isDragging) return;
    
    const container = document.querySelector('.chat-container');
    if (!container) return;
    
    const containerRect = container.getBoundingClientRect();
    const containerWidth = containerRect.width;
    
    const percentage = ((e.clientX - containerRect.left) / containerWidth) * 100;
    
    // Keep the left column at 40-80% and the right column at 20-60%
    leftColumnWidth = Math.min(Math.max(percentage, 40), 80);
    rightColumnWidth = 100 - leftColumnWidth;
  }
  
  function startVerticalResize(e: MouseEvent | KeyboardEvent) {
    isVerticalDragging = true;
    e.preventDefault();
    
    if (e instanceof MouseEvent) {
      initialMouseY = e.clientY;
      initialInputHeight = messageInputHeight;
    }
    
    window.addEventListener('mousemove', handleVerticalResize);
    window.addEventListener('mouseup', stopVerticalResize);
  }
  
  function handleVerticalKeyDown(e: KeyboardEvent) {
    if (e.key === 'Enter' || e.key === ' ') {
      startVerticalResize(e);
    }
  }
  
  function handleVerticalResize(e: MouseEvent) {
    if (!isVerticalDragging) return;
    
    const leftColumn = document.querySelector('.left-column');
    if (!leftColumn) return;
    
    const sysInstructions = leftColumn.querySelector('.system-instructions');
    if (!sysInstructions) return;
    
    const columnRect = leftColumn.getBoundingClientRect();
    const columnHeight = columnRect.height;
    
    const mouseDelta = e.clientY - initialMouseY;
    const deltaPercentage = (mouseDelta / columnHeight) * 100;
    const newHeight = initialInputHeight + deltaPercentage;
    
    // Limit the height so the system instructions stay visible
    const minHeight = DEFAULT_INPUT_HEIGHT * 0.25;
    const maxHeight = Math.min(MAX_INPUT_HEIGHT, 100 - MIN_SYSTEM_INSTRUCTIONS_HEIGHT);
    
    const constrainedHeight = Math.min(Math.max(newHeight, minHeight), maxHeight);
    const newSysInstructionsHeight = 100 - constrainedHeight;
    
    // The system instructions have min-h-[100px]. Do not resize below that.
    const sysInstructionsPixelHeight = (columnHeight * newSysInstructionsHeight) / 100;
    if (sysInstructionsPixelHeight < 100) return;
    
    messageInputHeight = constrainedHeight;
    systemInstructionsHeight = newSysInstructionsHeight;
  }
  
  function stopVerticalResize() {
    isVerticalDragging = false;
    window.removeEventListener('mousemove', handleVerticalResize);
    window.removeEventListener('mouseup', stopVerticalResize);
  }
  
  function stopResize() {
    isDragging = false;
    window.removeEventListener('mousemove', handleResize);
    window.removeEventListener('mouseup', stopResize);
  }

  onMount(() => {
    return () => {
      window.removeEventListener('mousemove', handleResize);
      window.removeEventListener('mouseup', stopResize);
      window.removeEventListener('mousemove', handleVerticalResize);
      window.removeEventListener('mouseup', stopVerticalResize);
    };
  });

  $: showObsidian = $featureFlags.enableObsidianIntegration;
</script>

<div class="chat-container flex h-full min-h-0 w-full gap-0 p-2">
  <aside class="flex flex-col gap-2 pr-2 left-column" style="width: {leftColumnWidth}%">
    <div class="bg-background/5 p-2 rounded-lg">
      <div class="rounded-lg bg-background/10">
        <DropdownGroup />
      </div>
    </div>

    <div class="bg-background/5 rounded-lg overflow-hidden" style="height: {messageInputHeight}%; max-height: {MAX_INPUT_HEIGHT}%">
      <ChatInput />
    </div>

    <button 
      class="vertical-resize-handle" 
      on:mousedown={startVerticalResize}
      on:keydown={handleVerticalKeyDown}
      type="button"
      aria-label="Resize message input and system instructions"
    ></button>

    <div class="flex-1 min-h-[100px] bg-background/5 p-2 rounded-lg system-instructions">
      <div class="h-full flex flex-col">
        <Textarea
          bind:value={$systemPrompt}
          readonly={true}
          placeholder="System instructions will appear here when you select a pattern..."
          class="w-full flex-1 bg-primary-800/30 rounded-lg border-none whitespace-pre-wrap overflow-y-auto resize-none text-sm scrollbar-thin scrollbar-thumb-white/10 scrollbar-track-transparent hover:scrollbar-thumb-white/20"
        />
      </div>
    </div>
  </aside>

  <button 
    class="resize-handle" 
    on:mousedown={startResize}
    on:keydown={handleKeyDown}
    type="button"
    aria-label="Resize chat panels"
  ></button>

  <div class="flex flex-col gap-2" style="width: {rightColumnWidth}%">
    <div class="flex items-center justify-between px-2 py-1">
      <div class="flex items-center gap-2">
        {#if showObsidian}
          <div class="flex items-center gap-2">
            <div class="flex items-center gap-1">
              <Checkbox
                bind:checked={$obsidianSettings.saveToObsidian}
                id="save-to-obsidian"
                class="h-3 w-3"
              />
              <Label for="save-to-obsidian" class="text-xs text-white/70">Save to Obsidian</Label>
            </div>
            {#if $obsidianSettings.saveToObsidian}
              <Input
                id="note-name"
                bind:value={$obsidianSettings.noteName}
                placeholder="Note name..."
                class="h-6 text-xs w-48 bg-white/5 border-none focus:ring-1 ring-white/20"
              />
            {/if}
          </div>
        {/if}
      </div>
      <Button variant="ghost" size="sm" class="h-6 px-2 text-xs opacity-70 hover:opacity-100" on:click={openDrawer}>
        <Tooltip text="Take Notes" position="left">
          <span>Take Notes</span>
        </Tooltip>
      </Button>
    </div>

    <div class="flex-1 flex flex-col min-h-0">
      <div class="flex-1 min-h-0 bg-background/5 rounded-lg overflow-y-scroll scrollbar-thin scrollbar-thumb-white/10 scrollbar-track-transparent hover:scrollbar-thumb-white/20">
        <ChatMessages />
        <div class="h-32"></div> <!-- Spacer so the container can scroll past the last message -->
      </div>
    </div>
  </div>
</div>

<NoteDrawer />

<style>
  .resize-handle {
    width: 6px;
    margin: 0 -3px;
    height: 100%;
    cursor: col-resize;
    position: relative;
    z-index: 10;
    transition: background-color 0.2s;
  }

  .resize-handle::after {
    content: "";
    position: absolute;
    top: 0;
    left: 50%;
    transform: translateX(-50%);
    height: 100%;
    width: 2px;
    background-color: rgba(255, 255, 255, 0.1);
    transition: background-color 0.2s, width 0.2s;
  }

  .resize-handle:hover::after,
  .resize-handle:focus::after {
    background-color: rgba(255, 255, 255, 0.3);
    width: 4px;
  }

  .resize-handle:focus {
    outline: none;
  }

  .resize-handle:focus-visible::after {
    background-color: rgba(255, 255, 255, 0.5);
    width: 4px;
  }

  .vertical-resize-handle {
    height: 6px;
    margin: -3px 0;
    width: 100%;
    cursor: row-resize;
    position: relative;
    z-index: 10;
    transition: background-color 0.2s;
  }

  .vertical-resize-handle::after {
    content: "";
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    width: 100%;
    height: 2px;
    background-color: rgba(255, 255, 255, 0.1);
    transition: background-color 0.2s, height 0.2s;
  }

  .vertical-resize-handle:hover::after,
  .vertical-resize-handle:focus::after {
    background-color: rgba(255, 255, 255, 0.3);
    height: 4px;
  }

  .vertical-resize-handle:focus {
    outline: none;
  }

  .vertical-resize-handle:focus-visible::after {
    background-color: rgba(255, 255, 255, 0.5);
    height: 4px;
  }

  @keyframes flash {
    0% { opacity: 1; }
    50% { opacity: 0.5; }
    100% { opacity: 1; }
  }
</style>
