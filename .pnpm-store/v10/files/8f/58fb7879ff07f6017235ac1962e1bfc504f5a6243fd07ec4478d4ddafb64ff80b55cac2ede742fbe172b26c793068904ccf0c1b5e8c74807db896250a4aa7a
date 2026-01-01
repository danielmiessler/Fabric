<!--
  @license lucide-svelte v0.309.0 - ISC

This source code is licensed under the ISC license.
See the LICENSE file in the root directory of this source tree.
-->

<script>import defaultAttributes from './defaultAttributes';
export let name;
export let color = 'currentColor';
export let size = 24;
export let strokeWidth = 2;
export let absoluteStrokeWidth = false;
export let iconNode;
</script>

<svg
  {...defaultAttributes}
  {...$$restProps}
  width={size}
  height={size}
  stroke={color}
  stroke-width={
    absoluteStrokeWidth
      ? Number(strokeWidth) * 24 / Number(size)
      : strokeWidth
  }
  class={`lucide-icon lucide lucide-${name} ${$$props.class ?? ''}`}
>
  {#each iconNode as [tag, attrs]}
    <svelte:element this={tag} {...attrs}/>
  {/each}
  <slot />
</svg>
