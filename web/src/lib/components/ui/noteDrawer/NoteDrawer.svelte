<script lang="ts">
  import { onMount } from 'svelte';
  import { fly } from 'svelte/transition';
  import { noteStore } from '$lib/store/note-store';
  import { afterNavigate, beforeNavigate } from '$app/navigation';
  import { clickOutside } from '$lib/actions/clickOutside';
  import { drawerStore } from '$lib/store/drawer-store';
  import { toastStore } from '$lib/store/toast-store';

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
      toastStore.warning('Cannot save empty note');
      return;
    }

    try {
      saving = true;
      await noteStore.save();

      toastStore.success('Note saved successfully!');
    } catch (error) {
      console.error('Failed to save note:', error);
      toastStore.error(error instanceof Error ? error.message : 'Failed to save notes');
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

<!-- Skeleton 5 removed the Drawer utility. The backdrop and the panel below take
  its place. The panel keeps the width, height, padding and margin that the
  Drawer received through its props. The height was `calc(100vh-theme(spacing.32))`,
  and spacing.32 is 8rem. -->
{#if $drawerStore.open}
  <div class="fixed inset-0 z-40 bg-surface-950/50" aria-hidden="true"></div>
  <aside
    class="fixed left-0 top-0 z-50 mt-16 flex h-[calc(100vh-8rem)] w-[40%] flex-col bg-surface-100-900 p-4 shadow-xl"
    transition:fly={{ x: -320, duration: 200 }}
  >
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
        ></textarea>
      </div>

      <footer class="flex-none flex justify-between items-center p-2 border-t border-white/10">
        <span class="text-xs opacity-70">
          {#if $noteStore.isDirty}
            Unsaved changes
          {/if}
        </span>
        <div class="flex gap-2">
          <button
            class="btn btn-sm preset-filled-surface-500"
            on:click={noteStore.reset}
          >
            Reset
          </button>
          <button
            class="btn btn-sm preset-filled-primary-500"
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
  </aside>
{/if}
