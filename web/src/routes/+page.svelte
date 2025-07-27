<script lang="ts">
  import Terminal from '$lib/components/terminal/Terminal.svelte';
  import Connections from '$lib/components/ui/connections/Connections.svelte';
  import { slide } from 'svelte/transition';
  import { quintOut } from 'svelte/easing';
  
  let augemented = false;
  let showTerminal = false;

  setTimeout(() => {
    augemented = true;
    showTerminal = true;
  }, 1000);
</script>

<div class="absolute inset-0 -z-10 overflow-hidden h-96">
  <Connections  particleCount={100} particleSize={3} particleSpeed={0.1} connectionDistance={100}/>
</div>

<div class="flex flex-col justify-between items-center">
  <h1 class="h1 text-8xl font-bold font-sans mt-8">
    <span class="bg-gradient-to-br to-blue-500 from-cyan-300 bg-clip-text text-transparent box-decoration-clone">fabric</span>
  </h1>
  {#if augemented}
    <div class="py-2 mb-4" transition:slide|local="{{delay: 250, duration: 3000, easing: quintOut }}">
      <h2 class="h2 text-2xl text-center font-extrabold bg-gradient-to-br to-blue-500 from-cyan-300 bg-clip-text text-transparent pb-2">Human Flourishing via AI Augmentation</h2>
      <div class="text-2xl">
        <p class="mt-2 font-bold">Fill in the blanks... </p>
        <i class="mt-2 font-medium justify-end">I believe one of the biggest problems in my world is <code class="code variant-filled-secondary">___________,</code> which is why I'm building/creating/doing <code class="code variant-filled-secondary">___________.</code>
      </div>
    </div>
  {/if}
</div>

{#if showTerminal}
  <div >
    <Terminal />
  </div>
{/if}

