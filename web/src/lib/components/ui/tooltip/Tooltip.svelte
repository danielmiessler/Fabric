<script lang="ts">
  import { onMount } from 'svelte';
  import { calculateTooltipPosition, formatPositionStyle, type TooltipPosition } from './positioning';

  export let text: string;
  // biome-ignore lint/style/useConst: Svelte props must use 'let' even when not reassigned
  export let position: TooltipPosition = 'top';

  let tooltipVisible = false;
  // eslint-disable-next-line no-unassigned-vars -- Assigned via bind:this in template
  let triggerElement: HTMLDivElement;
  let isBrowser = false;
  // biome-ignore lint/correctness/noUnusedVariables: Used in template for aria-describedby and id
  const tooltipId = `tooltip-${Math.random().toString(36).substring(2, 9)}`;

  // Reactive tooltip positioning - recalculates when position or element changes
  $: tooltipStyle = triggerElement && tooltipVisible
    ? formatPositionStyle(calculateTooltipPosition(triggerElement.getBoundingClientRect(), position))
    : '';

  function updatePosition() {
    if (triggerElement && tooltipVisible) {
      tooltipStyle = formatPositionStyle(calculateTooltipPosition(triggerElement.getBoundingClientRect(), position));
    }
  }

  // biome-ignore lint/correctness/noUnusedVariables: Used in template event handlers
  function showTooltip() {
    tooltipVisible = true;
  }

  // biome-ignore lint/correctness/noUnusedVariables: Used in template event handlers
  function hideTooltip() {
    tooltipVisible = false;
  }

  // Handle window scroll and resize to keep tooltip positioned correctly
  // Only runs in browser (not during SSR)
  onMount(() => {
    isBrowser = true;
    return () => {
      if (isBrowser) {
        window.removeEventListener('scroll', updatePosition, true);
        window.removeEventListener('resize', updatePosition);
      }
    };
  });

  // Add/remove event listeners reactively when tooltip visibility changes
  $: if (isBrowser && tooltipVisible) {
    window.addEventListener('scroll', updatePosition, true);
    window.addEventListener('resize', updatePosition);
  } else if (isBrowser && !tooltipVisible) {
    window.removeEventListener('scroll', updatePosition, true);
    window.removeEventListener('resize', updatePosition);
  }
</script>

<div class="tooltip-container">
  <div
    bind:this={triggerElement}
    class="tooltip-trigger"
    on:mouseenter={showTooltip}
    on:mouseleave={hideTooltip}
    on:focusin={showTooltip}
    on:focusout={hideTooltip}
    aria-describedby={tooltipVisible ? tooltipId : undefined}
    role="button"
    tabindex="0"
  >
    <slot></slot>
  </div>

  {#if tooltipVisible}
    <div
      id={tooltipId}
      class="tooltip fixed z-[9999] px-2 py-1 text-xs rounded bg-gray-900/90 text-white whitespace-nowrap shadow-lg backdrop-blur-sm"
      class:top="{position === 'top'}"
      class:bottom="{position === 'bottom'}"
      class:left="{position === 'left'}"
      class:right="{position === 'right'}"
      style={tooltipStyle}
      role="tooltip"
    >
      {text}
      <div class="tooltip-arrow" role="presentation"></div>
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
    transition: opacity 150ms ease-in-out;
    opacity: 1;
  }

  .tooltip.top {
    transform: translate(-50%, -100%);
  }

  .tooltip.bottom {
    transform: translate(-50%, 0);
  }

  .tooltip.left {
    transform: translate(-100%, -50%);
  }

  .tooltip.right {
    transform: translate(0, -50%);
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
