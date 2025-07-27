<script lang="ts">
  export let text: string;
  export let position: 'top' | 'bottom' | 'left' | 'right' = 'top';

  let tooltipVisible = false;
  let tooltipElement: HTMLDivElement;

  function showTooltip() {
    tooltipVisible = true;
  }

  function hideTooltip() {
    tooltipVisible = false;
  }
</script>

<!-- svelte-ignore a11y-no-noninteractive-element-interactions a11y-mouse-events-have-key-events -->
<div class="tooltip-container">
  <div 
    class="tooltip-trigger"
    on:mouseenter={showTooltip}
    on:mouseleave={hideTooltip}
    on:focusin={showTooltip}
    on:focusout={hideTooltip}
    role="tooltip"
    aria-label="Tooltip trigger"
  >
    <slot />
  </div>
  
  {#if tooltipVisible}
    <div
      bind:this={tooltipElement}
      class="tooltip absolute z-[9999] px-2 py-1 text-xs rounded bg-gray-900/90 text-white whitespace-nowrap shadow-lg backdrop-blur-sm"
      class:top="{position === 'top'}"
      class:bottom="{position === 'bottom'}"
      class:left="{position === 'left'}"
      class:right="{position === 'right'}"
      role="tooltip"
      aria-label={text}
    >
      {text}
      <div class="tooltip-arrow" role="presentation" />
    </div>
  {/if}
</div>

<style>
  .tooltip-container {
    position: relative;
    display: inline-block;
  }

  .tooltip-trigger {
    display: inline-flex;
  }

  .tooltip {
    pointer-events: none;
    transition: all 150ms ease-in-out;
    opacity: 1;
  }

  .tooltip.top {
    bottom: calc(100% + 5px);
    left: 50%;
    transform: translateX(-50%);
  }

  .tooltip.bottom {
    top: calc(100% + 5px);
    left: 50%;
    transform: translateX(-50%);
  }

  .tooltip.left {
    right: calc(100% + 5px);
    top: 50%;
    transform: translateY(-50%);
  }

  .tooltip.right {
    left: calc(100% + 5px);
    top: 50%;
    transform: translateY(-50%);
  }

  .tooltip-arrow {
    position: absolute;
    width: 8px;
    height: 8px;
    background: inherit;
    transform: rotate(45deg);
  }

  .tooltip.top .tooltip-arrow {
    bottom: -4px;
    left: 50%;
    margin-left: -4px;
  }

  .tooltip.bottom .tooltip-arrow {
    top: -4px;
    left: 50%;
    margin-left: -4px;
  }

  .tooltip.left .tooltip-arrow {
    right: -4px;
    top: 50%;
    margin-top: -4px;
  }

  .tooltip.right .tooltip-arrow {
    left: -4px;
    top: 50%;
    margin-top: -4px;
  }
</style>
