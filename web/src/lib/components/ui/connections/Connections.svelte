<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { ParticleSystem } from './ParticleSystem';
  import { createParticleGradient } from '$lib/components/ui/connections/canvas';

  export let particleCount = 100;
  export let particleSize = 3;
  export let particleSpeed = 0.5;
  export let connectionDistance = 100;

  let canvas: HTMLCanvasElement;
  let ctx: CanvasRenderingContext2D;
  let animationFrame: number;
  let particleSystem: ParticleSystem;
  let isMouseOver = false;
  let browser = false;

  function handleMouseMove(event: MouseEvent) {
    if (!isMouseOver) return;
    const rect = canvas.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;
    particleSystem?.updateMousePosition(x, y);
  }

  function handleMouseEnter() {
    isMouseOver = true;
  }

  function handleMouseLeave() {
    isMouseOver = false;
    const centerX = canvas.width / 2;
    const centerY = canvas.height / 2;
    particleSystem?.updateMousePosition(centerX, centerY);
  }

  function handleResize() {
    if (!browser) return;
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    ctx.scale(window.devicePixelRatio, window.devicePixelRatio);
    particleSystem?.updateDimensions(canvas.width, canvas.height);
  }

  function drawConnections() {
    const particles = particleSystem.getParticles();
    ctx.lineWidth = 0.5;

    for (let i = 0; i < particles.length; i++) {
      for (let j = i + 1; j < particles.length; j++) {
        const dx = particles[i].x - particles[j].x;
        const dy = particles[i].y - particles[j].y;
        const distance = Math.sqrt(dx * dx + dy * dy);

        if (distance < connectionDistance) {
          const alpha = 1 - (distance / connectionDistance);
          ctx.strokeStyle = `rgba(255, 255, 255, ${alpha * 0.15})`; // Slightly reduced opacity
          ctx.beginPath();
          ctx.moveTo(particles[i].x, particles[i].y);
          ctx.lineTo(particles[j].x, particles[j].y);
          ctx.stroke();
        }
      }
    }
  }

  function drawParticles() {
    const particles = particleSystem.getParticles();

    particles.forEach(particle => {
      const gradient = createParticleGradient(
        ctx,
        particle.x,
        particle.y,
        particle.size,
        particle.color
      );

      ctx.beginPath();
      ctx.fillStyle = gradient;
      ctx.arc(particle.x, particle.y, particle.size, 0, Math.PI * 2);
      ctx.fill();
    });
  }

  function animate() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    particleSystem.update();
    drawConnections();
    drawParticles();

    animationFrame = requestAnimationFrame(animate);
  }

  onMount(() => {
    browser = true;
    if (!browser) return;

    ctx = canvas.getContext('2d')!;
    handleResize();

    particleSystem = new ParticleSystem(
      particleCount,
      particleSize,
      particleSpeed,
      canvas.width,
      canvas.height
    );

    window.addEventListener('resize', handleResize);
    animationFrame = requestAnimationFrame(animate);
  });

  onDestroy(() => {
    if (!browser) return;
    window.removeEventListener('resize', handleResize);
    cancelAnimationFrame(animationFrame);
  });
</script>
  
<canvas
  bind:this={canvas}
  on:mousemove={handleMouseMove}
  on:mouseenter={handleMouseEnter}
  on:mouseleave={handleMouseLeave}
  class="particle-wave"
/>
  
<style>
.particle-wave {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: transparent;
}
</style>
