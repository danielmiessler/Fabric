<script lang="ts">
  import { Drawer, getDrawerStore, getToastStore } from '@skeletonlabs/skeleton';
  import type { DrawerStore } from '@skeletonlabs/skeleton';
  import { onMount } from 'svelte';
  import { noteStore } from '$lib/store/note-store';
  import { afterNavigate, beforeNavigate } from '$app/navigation';
  import { clickOutside } from '$lib/actions/clickOutside';
  
  const drawerStore = getDrawerStore();
  const toastStore = getToastStore();
  
  let textareaEl: HTMLTextAreaElement;
  let saving = false;

  let content = '';
  
  // Auto-resize textarea
  function adjustTextareaHeight() {
    if (textareaEl) {
      textareaEl.style.height = 'auto';
      textareaEl.style.height = textareaEl.scrollHeight + 'px';
    }
  }
  
  async function saveContent() {
    if (!$noteStore.content.trim()) {
      toastStore.trigger({
        message: 'Cannot save empty note',
        background: 'variant-filled-warning'
      });
      return;
    }

    try {
      saving = true;
      await noteStore.save();
      
      toastStore.trigger({
        message: `Note saved successfully!`,
        background: 'variant-filled-success'
      });
    } catch (error) {
      console.error('Failed to save note:', error);
      toastStore.trigger({
        message: error instanceof Error ? error.message : 'Failed to save notes',
        background: 'variant-filled-error'
      });
    } finally {
      saving = false;
    }
  }
  
  // Prompt user if trying to close with unsaved changes
  $: if ($drawerStore.open === false && $noteStore.isDirty) {
    if (confirm('You have unsaved changes. Are you sure you want to close?')) {
      noteStore.reset();
    } else {
      drawerStore.open({});
    }
  }
  
  // Load saved content when drawer opens
  $: if ($drawerStore.open) {
    const savedContent = localStorage.getItem('savedText');
    if (savedContent) {
      noteStore.updateContent(savedContent);
      noteStore.save();
    }
  }
  
  // Keyboard shortcuts
  function handleKeydown(event: KeyboardEvent) {
    if ((event.ctrlKey || event.metaKey) && event.key === 's') {
      event.preventDefault();
      saveContent();
    }
  }
  
  onMount(() => {
    adjustTextareaHeight();
  });
</script>

<Drawer width="w-[40%]" class="flex flex-col h-[calc(100vh-theme(spacing.32))] p-4 mt-16">
  {#if $drawerStore.open}
    <div 
      class="flex flex-col h-full"
      use:clickOutside={() => {
        if ($noteStore.isDirty) {
          if (confirm('You have unsaved changes. Are you sure you want to close?')) {
            noteStore.reset();
            drawerStore.close();
          }
        } else {
          drawerStore.close();
        }
      }}
    >
      <header class="flex-none p-2 border-b border-white/10">
        <div class="flex justify-between items-center">
          <h2 class="text-lg font-semibold">Notes</h2>
          {#if $noteStore.lastSaved}
            <span class="text-xs opacity-70">
              Last saved: {$noteStore.lastSaved.toLocaleTimeString()}
            </span>
          {/if}
        </div>
        <div class="flex gap-4 mt-2 text-xs opacity-70">
          <span>Notes saved to <code>inbox/</code></span>
          <span>Ctrl + S to save</span>
        </div>
      </header>

      <div class="flex-1 p-2">
        <textarea
        bind:this={textareaEl}
        value={$noteStore.content}
        on:input={e => noteStore.updateContent(e.currentTarget.value)}
        on:keydown={handleKeydown}
        class="w-full h-full min-h-[300px] resize-none p-2 rounded-lg bg-primary-800/30 border-none text-sm"
        placeholder="Enter your text here..." 
        />
      </div>

      <footer class="flex-none flex justify-between items-center p-2 border-t border-white/10">
        <span class="text-xs opacity-70">
          {#if $noteStore.isDirty}
            Unsaved changes
          {/if}
        </span>
        <div class="flex gap-2">
          <button
            class="btn btn-sm variant-filled-surface"
            on:click={noteStore.reset}
          >
            Reset
          </button>
          <button
            class="btn btn-sm variant-filled-primary"
            on:click={saveContent}
          >
            {#if saving}
              Saving...
            {:else}
              Save
            {/if}
          </button>
        </div>
      </footer>
    </div>
  {/if}
</Drawer>
