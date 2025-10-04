<script>
  import '../app.postcss';
  import ToastContainer from '$lib/components/ui/toast/ToastContainer.svelte';
  import AppHeader from '$lib/components/layout/AppHeader.svelte';
  import SessionSidebar from '$lib/components/layout/SessionSidebar.svelte';
  import ContextPanel from '$lib/components/layout/ContextPanel.svelte';
  import Footer from '$lib/components/home/Footer.svelte';
  import { onMount } from 'svelte';
  import { toastStore } from '$lib/store/toast-store';
  import { initializeStores } from '@skeletonlabs/skeleton';

  initializeStores();

  onMount(() => {
    toastStore.info('Welcome to Fabric — connect your workflows and share insights.');
  });
</script>

<div class="relative min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 text-foreground">
  <div class="pointer-events-none absolute inset-0 bg-[radial-gradient(circle_at_top,_rgba(59,130,246,0.12),_transparent_55%)]"></div>
  <div class="relative z-10 flex min-h-screen flex-col">
    <ToastContainer />
    <AppHeader />
    <div class="flex flex-1 overflow-hidden">
      <SessionSidebar />
      <main class="flex-1 overflow-hidden">
        <div class="h-full overflow-y-auto px-6 py-8">
          <slot />
        </div>
      </main>
      <ContextPanel />
    </div>
    <Footer />
  </div>
</div>
