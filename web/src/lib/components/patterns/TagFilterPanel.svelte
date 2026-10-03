<script lang="ts">
    import type { Pattern } from '$lib/interfaces/pattern-interface';
    import { createEventDispatcher } from 'svelte';
    
    const dispatch = createEventDispatcher<{
        tagsChanged: string[];
    }>();

    export let patterns: Pattern[];
    export let hideToggleButton = false;
    let selectedTags: string[] = [];
    let isExpanded = false;

    function toggleTag(tag: string) {
        selectedTags = selectedTags.includes(tag)
            ? selectedTags.filter(t => t !== tag)
            : [...selectedTags, tag];
        dispatch('tagsChanged', selectedTags);
    }

    function togglePanel() {
        isExpanded = !isExpanded;
    }

    export function reset() {
        selectedTags = [];
        isExpanded = false;
        dispatch('tagsChanged', selectedTags);
    }
</script>

<div class="tag-panel {isExpanded ? 'expanded' : ''} {hideToggleButton ? 'embedded' : ''}" style="z-index: 50">
    {#if !hideToggleButton}
    <div class="panel-header">
        <button class="close-btn" on:click={togglePanel}>
            {isExpanded ? 'Close Filter Tags ◀' : 'Open Filter Tags ▶'}
        </button>
    </div>
    {/if}
    
    <div class="panel-content {hideToggleButton ? 'always-visible' : ''}">
        <div class="reset-container">
            <button 
                class="reset-btn"
                on:click={() => {
                    selectedTags = [];
                    dispatch('tagsChanged', selectedTags);
                }}
            >
                Reset All Tags
            </button>
        </div>
        {#each Array.from(new Set(patterns.flatMap(p => p.tags || []))).sort() as tag}
            <button 
                class="tag-brick {selectedTags.includes(tag) ? 'selected' : ''}"
                on:click={() => toggleTag(tag)}
            >
                {tag}
            </button>
        {/each}
    </div>
</div>
<style>
   .tag-panel {
    position: fixed;
    left: calc(50% + 300px); /* 300px is half the width of the centered PatternList modal */
    top: 50%;
    transform: translateY(-50%);
    width: 300px;
    transition: left 0.3s ease;
}

.tag-panel.embedded {
    position: relative;
    left: auto;
    top: auto;
    transform: none;
    width: 100%;
    height: 100%;
}

.tag-panel.expanded {
    left: calc(50% + 360px); /* 60px right of the modal edge when open */
}

.panel-content {
    display: none;
    padding: 12px;
    flex-wrap: wrap;
    gap: 6px;
    max-height: 80vh;
    overflow-y: auto;
}

.embedded .panel-content {
    max-height: 100%;
}

.panel-content.always-visible {
    display: flex;
}

.tag-brick {
    padding: 4px 8px;
    font-size: 0.8rem;
    border-radius: 12px;
    background: rgba(255,255,255,0.1);
    cursor: pointer;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
}

.reset-container {
    width: 100%;
    padding-bottom: 8px;
    margin-bottom: 8px;
    border-bottom: 1px solid rgba(255,255,255,0.1);
}

.reset-btn {
    width: 100%;
    padding: 6px;
    font-size: 0.8rem;
    color: var(--primary-300);
    background: rgba(255,255,255,0.05);
    border-radius: 4px;
    transition: all 0.2s;
}

.reset-btn:hover {
    background: rgba(255,255,255,0.1);
}

.expanded .panel-content {
    display: flex;
}

.panel-header {
    padding: 8px;
    border-bottom: 1px solid rgba(255,255,255,0.1);
}

.close-btn {
    width: auto;
    padding: 6px;
    position: absolute;
    font-size: 0.8rem;
    color: var(--primary-300);
    background: rgba(255,255,255,0.05);
    border-radius: 4px;
    transition: all 0.2s;
    text-align: left;
}

.tag-panel:not(.expanded) .close-btn {
    top: -290px; /* Aligns the button with the search bar */
    margin-left: 10px;
}

.expanded .close-btn {
    position: relative;
    top: 0;
    margin-left: -50px;
}

.close-btn:hover {
    background: rgba(255,255,255,0.1);
}

.tag-brick.selected {
    background: var(--primary-300);
}
</style>
