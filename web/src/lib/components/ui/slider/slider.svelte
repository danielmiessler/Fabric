<script lang="ts">
  import { cn } from "$lib/utils/utils";

  let className: string | undefined = undefined;
  export let value = 0;
  export let min = 0;
  export let max = 100;
  export let step = 1;
  export { className as class };

  let sliderEl: HTMLDivElement;
  let isDragging = false;

  $: percentage = ((value - min) / (max - min)) * 100;

  function handleMouseDown(e: MouseEvent) {
    isDragging = true;
    updateValue(e);
    window.addEventListener('mousemove', handleMouseMove);
    window.addEventListener('mouseup', handleMouseUp);
  }

  function handleMouseMove(e: MouseEvent) {
    if (!isDragging) return;
    updateValue(e);
  }

  function handleMouseUp() {
    isDragging = false;
    window.removeEventListener('mousemove', handleMouseMove);
    window.removeEventListener('mouseup', handleMouseUp);
  }

  function updateValue(e: MouseEvent) {
    if (!sliderEl) return;
    const rect = sliderEl.getBoundingClientRect();
    const pos = (e.clientX - rect.left) / rect.width;
    const rawValue = min + (max - min) * pos;
    const steppedValue = Math.round(rawValue / step) * step;
    value = Math.max(min, Math.min(max, steppedValue));
  }

  function handleKeyDown(e: KeyboardEvent) {
    const stepSize = e.shiftKey ? 10 : 1;
    switch (e.key) {
      case 'ArrowLeft':
      case 'ArrowDown':
        e.preventDefault();
        value = Math.max(min, value - step);
        break;
      case 'ArrowRight':
      case 'ArrowUp':
        e.preventDefault();
        value = Math.min(max, value + step);
        break;
      case 'Home':
        e.preventDefault();
        value = min;
        break;
      case 'End':
        e.preventDefault();
        value = max;
        break;
    }
  }
</script>

<div
	bind:this={sliderEl}
	class={cn("relative flex w-full touch-none select-none items-center", className)}
	on:mousedown={handleMouseDown}
	on:keydown={handleKeyDown}
	role="slider"
	tabindex="0"
	aria-valuemin={min}
	aria-valuemax={max}
	aria-valuenow={value}
>
	<div class="relative h-0.5 w-full grow overflow-hidden rounded-full bg-white/10">
		<div 
			class="absolute h-full bg-white/30 transition-all" 
			style="width: {percentage}%" 
		/>
	</div>
	<div
		class="absolute h-2.5 w-2.5 rounded-full bg-white/70 ring-1 ring-white/10 shadow-sm transition-all hover:scale-110 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-white/30 disabled:pointer-events-none disabled:opacity-50"
		style="left: calc({percentage}% - 0.3125rem)"
	/>
</div>
