<script lang="ts">
    import { createEventDispatcher, onMount } from 'svelte';
    import TagFilterPanel from '$lib/components/patterns/TagFilterPanel.svelte';
    let tagFilterRef: TagFilterPanel;
    let selectedTags: string[] = [];
    import { cn } from "$lib/utils/utils";
    import type { Pattern } from '$lib/interfaces/pattern-interface';
    import { patterns, patternAPI, selectedPatternName } from '$lib/store/pattern-store';
    import { favorites } from '$lib/store/favorites-store';
    import { Input } from "$lib/components/ui/input";
    
    const dispatch = createEventDispatcher();
    let isTagPanelOpen = false;
    let searchQuery = '';
    let showOnlyFavorites = false;
    
    onMount(async () => {
      try {
        await patternAPI.loadPatterns();
      } catch (error) {
        console.error('Error loading patterns:', error);
      }
    });
    
function toggleFavorite(patternName: string) {
    favorites.toggleFavorite(patternName);
}

function selectPattern(patternName: string) {
  patternAPI.selectPattern(patternName);
  dispatch('select', patternName);
}

function closeModal() {
  dispatch('close');
}

function toggleTagPanel() {
    isTagPanelOpen = !isTagPanelOpen;
}

function handleTagFilter(event: CustomEvent<string[]>) {
  selectedTags = event.detail;
}

function toggleFavoritesFilter() {
  showOnlyFavorites = !showOnlyFavorites;
}

// Apply filtering based on search query, favorites filter, and tag selection
$: filteredPatterns = $patterns
  .filter(p => {
    // Apply favorites filter if enabled
    if (showOnlyFavorites && !$favorites.includes(p.Name)) {
      return false;
    }
    
    // Apply tag filter if any tags are selected
    if (selectedTags.length > 0) {
      if (!p.tags || !selectedTags.every(tag => p.tags.includes(tag))) {
        return false;
      }
    }
    
    // Apply search filter if query exists
    if (searchQuery.trim()) {
      return (
        p.Name.toLowerCase().includes(searchQuery.toLowerCase()) || 
        p.Description.toLowerCase().includes(searchQuery.toLowerCase()) ||
        (p.tags && p.tags.some(tag => tag.toLowerCase().includes(searchQuery.toLowerCase())))
      );
    }
    
    return true;
  });
</script>

<!-- Main container with flexible layout -->
<div class="flex h-[85vh]">
  <!-- Modal container with responsive positioning -->
  <div class={cn(
      "flex flex-col bg-primary-800 rounded-lg shadow-xl transition-all duration-300",
      isTagPanelOpen
        ? "w-[75vw]" 
        : "w-full max-w-[95vw] mx-auto"
    )}>
    <!-- Header with grid layout -->
    <div class="grid grid-cols-[auto_auto_1fr_auto] items-center p-4 border-b border-primary-700/30 sticky top-0 bg-primary-800 z-10">
      <!-- Left column: Title -->
      <h2 class="text-xl font-semibold text-primary-200 mr-4">Pattern Library</h2>
          
      <!-- Second column: Search -->
      <div class="mr-4">
        <Input 
          bind:value={searchQuery}
          placeholder="Search patterns..." 
          class="w-full min-w-[300px] text-emerald-900"
        />
      </div>
      
      <!-- Third column: Favorites button -->
      <div class="flex items-center">
        <button
          on:click={toggleFavoritesFilter}
          class={cn(
            "px-3 py-1.5 rounded-md text-sm font-medium transition-all",
            showOnlyFavorites 
              ? "bg-yellow-500/20 text-yellow-300 border border-yellow-500/30" 
              : "bg-primary-700/30 text-primary-300 border border-primary-600/20 hover:bg-primary-700/50"
          )}
        >
          <span class="mr-1">{showOnlyFavorites ? "★" : "☆"}</span>
          Favorites
        </button>
      </div>
        
      <!-- Fourth column: Other controls -->
      <div class="flex items-center gap-3 justify-end">
        <!-- Single tag panel toggle button -->
        <button
          on:click={toggleTagPanel}
          class={cn(
            "px-3 py-1.5 rounded-md text-sm font-medium transition-all",
            isTagPanelOpen
              ? "bg-blue-500/20 text-blue-300 border border-blue-500/30"
              : "bg-primary-700/30 text-primary-300 border border-primary-600/20 hover:bg-primary-700/50"
          )}
        >
          {isTagPanelOpen ? "Close Filter Tags ◀" : "Open Filter Tags ▶"}
        </button>

        <!-- Close modal button -->
        <button
          on:click={closeModal}
          class="px-2 py-2 rounded-full bg-primary-700/40 text-primary-200 hover:bg-primary-700/60 hover:text-primary-100"
          aria-label="Close modal"
        >
          <span class="text-xl font-bold">×</span>
        </button>
      </div>
    </div>

    <!-- Selected tags display -->
    {#if selectedTags.length > 0}
      <div class="px-4 pb-2 pt-2 border-b border-primary-700/30">
        <div class="text-sm text-white/70 bg-primary-700/30 rounded-md p-2 flex justify-between items-center">
          <div class="flex flex-wrap gap-1 items-center">
            <span class="mr-1">Tags:</span>
            {#each selectedTags as tag}
              <div class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-primary-600/40 text-primary-200 border border-primary-500/30">
                {tag}
                <button 
                  class="ml-1 text-xs text-primary-300 hover:text-primary-100"
                  on:click={() => {
                    selectedTags = selectedTags.filter(t => t !== tag);
                  }}
                >
                  ×
                </button>
              </div>
            {/each}
          </div>
          <button 
            class="px-2 py-1 text-xs text-white/70 bg-primary-600/30 rounded hover:bg-primary-600/50 transition-colors"
            on:click={() => {
              selectedTags = [];
              if (tagFilterRef && typeof tagFilterRef.reset === 'function') {
                tagFilterRef.reset();
              }
            }}
          >
            reset
          </button>
        </div>
      </div>
    {/if}
      
    <!-- Pattern tiles grid with scrolling -->
    <div class="flex-1 overflow-y-auto p-4 pattern-grid-container">
      {#if filteredPatterns.length === 0}
        <div class="flex justify-center items-center h-full">
          <p class="text-primary-300">
            {showOnlyFavorites
              ? "No favorite patterns found. Add some favorites first!"
              : "No patterns found matching your search."}
          </p>
        </div>
      {:else}
        <div class={cn(
          "grid grid-cols-1 sm:grid-cols-2 gap-4",
          isTagPanelOpen 
            ? "md:grid-cols-2 lg:grid-cols-3" 
            : "md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
        )}>
          {#each filteredPatterns as pattern}
            <button
              class="text-left border-2 border-primary-600/40 rounded-lg shadow-md hover:shadow-lg p-4 flex flex-col h-58 bg-primary-700/30 hover:bg-primary-700/50 transition-all transform hover:-translate-y-1 duration-200"
              on:click={() => selectPattern(pattern.Name)}
            >
              <div class="flex justify-between items-start mb-2">
                <h3 class="pattern-name font-bold text-base text-primary-200 leading-tight break-all overflow-hidden pr-2 w-[85%]">{pattern.Name}</h3>
                <button
                  on:click|stopPropagation={() => toggleFavorite(pattern.Name)}
                  class="focus:outline-none ml-1 mt-0.5"
                  aria-label={$favorites.includes(pattern.Name) ? 'Remove from favorites' : 'Add to favorites'}
                >
                  {#if $favorites.includes(pattern.Name)}
                    <span class="text-yellow-400 text-xl">★</span>
                  {:else}
                    <span class="text-primary-400 text-xl hover:text-yellow-300">☆</span>
                  {/if}
                </button>
              </div>
              
              <!-- Pattern description with scrolling if needed -->
              <div class="flex-grow overflow-y-auto mb-1 pr-1 custom-scrollbar">
                <p class="text-sm text-primary-300/90 leading-relaxed">{pattern.Description}</p>
              </div>
              
              <!-- Tags section -->
              {#if pattern.tags && pattern.tags.length > 0}
                <div class="flex flex-wrap gap-1 mt-2">
                  {#each pattern.tags as tag}
                    <span class="inline-flex items-center px-1 py-0.25 rounded-full text-[8px] font-medium bg-primary-600/40 text-primary-200 border border-primary-500/30">
                      {tag}
                    </span>
                  {/each}
                </div>
              {/if}
            </button>
          {/each}
        </div>
      {/if}
    </div>
  </div>

  <!-- Tag filter panel - positioned on the right when open -->
  {#if isTagPanelOpen}
    <div class="tag-panel-container">
      <div class="tag-panel-header">
        <button class="tag-panel-close" on:click={toggleTagPanel}>
          <span class="text-lg">×</span>
        </button>
        <h3 class="text-sm font-medium text-primary-200">Filter Tags</h3>
      </div>
      <div class="tag-panel-content">
        <TagFilterPanel 
          patterns={$patterns} 
          on:tagsChanged={handleTagFilter}
          bind:this={tagFilterRef}
          hideToggleButton={true}
        />
      </div>
    </div>
  {/if}
</div>

<style>
  /* Custom scrollbar styling remains the same */
  .custom-scrollbar::-webkit-scrollbar {
    width: 4px;
  }
  
  .custom-scrollbar::-webkit-scrollbar-track {
    background: rgba(31, 41, 55, 0.2);
    border-radius: 4px;
  }
  
  .custom-scrollbar::-webkit-scrollbar-thumb {
    background: rgba(156, 163, 175, 0.3);
    border-radius: 4px;
  }
  
  .custom-scrollbar::-webkit-scrollbar-thumb:hover {
    background: rgba(156, 163, 175, 0.5);
  }
  
  /* Add this to your <style> section */
  h3.pattern-name {
    word-break: break-all;      /* Force breaks anywhere if needed */
    hyphens: auto;              /* Enable hyphenation */
    overflow-wrap: break-word;  /* Fallback */
  }
 
  .custom-scrollbar {
    scrollbar-width: thin;
    scrollbar-color: rgba(156, 163, 175, 0.3) rgba(31, 41, 55, 0.2);
  }
  
  .pattern-grid-container {
    overflow-y: auto;
    scrollbar-width: thin;
    scrollbar-color: rgba(156, 163, 175, 0.3) rgba(31, 41, 55, 0.2);
  }

  /* Tag panel styling */
  .tag-panel-container {
    width: 20vw;
    height: 100%;
    background-color: #1e293b; /* Use a solid color instead of var */
    border-left: 1px solid rgba(255, 255, 255, 0.1);
    z-index: 20;
    box-shadow: -2px 0 10px rgba(0, 0, 0, 0.3);
  }

  .tag-panel-header {
    display: flex;
    align-items: center;
    padding: 12px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  .tag-panel-close {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 24px;
    height: 24px;
    border-radius: 50%;
    background: rgba(255, 255, 255, 0.1);
    margin-right: 8px;
    cursor: pointer;
  }

  .tag-panel-close:hover {
    background: rgba(255, 255, 255, 0.2);
  }

  .tag-panel-content {
    height: calc(100% - 49px);
    overflow-y: auto;
  }
</style>
