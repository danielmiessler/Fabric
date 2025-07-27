
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

<div class="bg-primary-800 rounded-lg flex flex-col h-[85vh] w-[600px] shadow-lg relative">
  <div class="flex flex-col border-b border-primary-700/30">
    <div class="flex justify-between items-center p-4">
      <b class="text-lg text-muted-foreground font-bold">Pattern Descriptions</b>
      <button
        on:click={closeModal}
        class="text-muted-foreground hover:text-primary-300 transition-colors"
      >
        ✕
      </button>
    </div>
    
    <div class="px-4 pb-4 flex items-center justify-between">
      <div class="flex-1 flex items-center">
        <div class="flex-1 mr-2">
          <Input
            bind:value={searchQuery}
            placeholder="Search patterns..."
            class="text-emerald-900"
          />
        </div>
        
        <!-- Favorites button similar to PatternTilesModal -->
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
    </div>

    <!-- Selected tags display -->
    <div class="px-4 pb-2">
      <div class="text-sm text-white/70 bg-primary-700/30 rounded-md p-2 flex justify-between items-center">
        <div class="flex flex-wrap gap-1 items-center">
          <span class="mr-1">Tags:</span>
          {#if selectedTags.length > 0}
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
          {:else}
            <span class="text-primary-300/50">none</span>
          {/if}
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
  </div>

  <div class="patterns-container p-4 flex-1 overflow-y-auto">
    {#if filteredPatterns.length === 0}
      <div class="flex justify-center items-center h-full">
        <p class="text-primary-300">
          {showOnlyFavorites
            ? "No favorite patterns found. Add some favorites first!"
            : "No patterns found matching your search."}
        </p>
      </div>
    {:else}
      <div class="patterns-list space-y-2">
        {#each filteredPatterns as pattern}
          <div class="pattern-item bg-primary/10 rounded-lg p-3">
            <div class="flex justify-between items-start gap-4 mb-2">
              <button
                class="text-xl font-bold text-primary-300 hover:text-primary-100 cursor-pointer transition-colors text-left w-full"
                on:click={() => selectPattern(pattern.Name)}
              >
                {pattern.Name}
              </button>
              <button
                class="text-muted-foreground hover:text-primary-300 transition-colors"
                on:click|stopPropagation={() => toggleFavorite(pattern.Name)}
              >
                {#if $favorites.includes(pattern.Name)}
                  <span class="text-yellow-400">★</span>
                {:else}
                  <span class="text-primary-400 hover:text-yellow-300">☆</span>
                {/if}
              </button>
            </div>
            <p class="text-sm text-muted-foreground break-words leading-relaxed">{pattern.Description}</p>
          </div>
        {/each}
      </div>
    {/if}
  </div>

  <TagFilterPanel 
    patterns={$patterns} 
    on:tagsChanged={handleTagFilter}
    bind:this={tagFilterRef}
    hideToggleButton={false}
  />
</div>

<style>
/* Custom scrollbar styling */
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

/* h3.pattern-name {
  word-break: break-all;
  hyphens: auto;
  overflow-wrap: break-word;
} */

.custom-scrollbar {
  scrollbar-width: thin;
  scrollbar-color: rgba(156, 163, 175, 0.3) rgba(31, 41, 55, 0.2);
}

.patterns-container {
  flex: 1;
  overflow-y: auto;
  scrollbar-width: thin;
  -ms-overflow-style: thin;
}

.patterns-list {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 560px;
  margin: 0 auto;
}

.pattern-item {
  display: flex;
  flex-direction: column;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.pattern-item:last-child {
  border-bottom: none;
}
</style>
