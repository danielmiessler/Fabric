<script lang="ts">
  export let text: string;
  export let position: 'top' | 'bottom' | 'left' | 'right' = 'top';

  let tooltipVisible = false;
  let tooltipElement: HTMLDivElement;
  let triggerElement: HTMLDivElement;
  let showTimeout: ReturnType<typeof setTimeout> | null = null;

  function showTooltip() {    
    // Clear any existing timeout
    if (showTimeout) {
      clearTimeout(showTimeout);
    }
    
    // Set 500ms 0.5 second delay before showing tooltip (reduced for testing)
    showTimeout = setTimeout(() => {
      tooltipVisible = true;
    }, 500);
  }

  function hideTooltip() {  
    // Clear the show timeout if user leaves before the set delay
    if (showTimeout) {
      clearTimeout(showTimeout);
      showTimeout = null;
    }
    
    tooltipVisible = false;
  }

  function updateTooltipPosition() {
    if (!tooltipVisible || !triggerElement || !tooltipElement) {
      return;
    }
    
    const triggerRect = triggerElement.getBoundingClientRect();
    const tooltipRect = tooltipElement.getBoundingClientRect();
    const viewportWidth = window.innerWidth;
    const viewportHeight = window.innerHeight;
    
    let top = 0;
    let left = 0;
    let actualPosition = position;
    
    // Smart positioning - flip to opposite side if not enough space
    switch (position) {
      case 'top':
        // Check if there's enough space above
        if (triggerRect.top - tooltipRect.height - 8 < 8) {
          actualPosition = 'bottom';
          top = triggerRect.bottom + 8;
        } else {
          top = triggerRect.top - tooltipRect.height - 8;
        }
        left = triggerRect.left + (triggerRect.width / 2) - (tooltipRect.width / 2);
        break;
        
      case 'bottom':
        // Check if there's enough space below
        if (triggerRect.bottom + tooltipRect.height + 8 > viewportHeight - 8) {
          actualPosition = 'top';
          top = triggerRect.top - tooltipRect.height - 8;
        } else {
          top = triggerRect.bottom + 8;
        }
        left = triggerRect.left + (triggerRect.width / 2) - (tooltipRect.width / 2);
        break;
        
      case 'left':
        // Check if there's enough space to the left
        if (triggerRect.left - tooltipRect.width - 8 < 8) {
          actualPosition = 'right';
          left = triggerRect.right + 8;
        } else {
          left = triggerRect.left - tooltipRect.width - 8;
        }
        top = triggerRect.top + (triggerRect.height / 2) - (tooltipRect.height / 2);
        break;
        
      case 'right':
        // Check if there's enough space to the right
        if (triggerRect.right + tooltipRect.width + 8 > viewportWidth - 8) {
          actualPosition = 'left';
          left = triggerRect.left - tooltipRect.width - 8;
        } else {
          left = triggerRect.right + 8;
        }
        top = triggerRect.top + (triggerRect.height / 2) - (tooltipRect.height / 2);
        break;
    }
    
    // Final viewport boundary enforcement (fallback)
    if (left < 8) left = 8;
    if (left + tooltipRect.width > viewportWidth - 8) {
      left = viewportWidth - tooltipRect.width - 8;
    }
    if (top < 8) top = 8;
    if (top + tooltipRect.height > viewportHeight - 8) {
      top = viewportHeight - tooltipRect.height - 8;
    }
    
    tooltipElement.style.top = `${top}px`;
    tooltipElement.style.left = `${left}px`;
  }

  // Update position when tooltip becomes visible
  $: if (tooltipVisible) {
    setTimeout(updateTooltipPosition, 0);
  }
</script>

<!-- Tooltip trigger container - no positioning constraints -->
<div 
  class="tooltip-trigger cursor-pointer" 
  bind:this={triggerElement}
  on:mouseenter={showTooltip}
  on:mouseleave={hideTooltip}
  on:focusin={showTooltip}
  on:focusout={hideTooltip}
  role="tooltip"
  aria-label="Tooltip trigger"
>
  <slot />
</div>

<!-- Tooltip rendered outside normal flow -->
{#if tooltipVisible}
  <div
    bind:this={tooltipElement}
    class="tooltip fixed z-[9999] px-3 py-2 text-sm rounded-lg bg-gray-500 text-white whitespace-nowrap shadow-xl border-2 border-gray-600"
    role="tooltip"
    aria-label={text}
  >
    {text}
  </div>
{/if}
