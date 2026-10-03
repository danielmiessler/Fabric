<script>
  import '../app.css';
  import ToastContainer from '$lib/components/ui/toast/ToastContainer.svelte';
  import Footer from '$lib/components/home/Footer.svelte';
  import Header from '$lib/components/home/Header.svelte';
  import { page } from '$app/stores';
  import { fly } from 'svelte/transition';
  import { onMount } from 'svelte';
  import { toastStore } from '$lib/store/toast-store';

  onMount(() => {
    toastStore.info("👋 Welcome to the site! Tell people about yourself and what you do.");
  });
</script>

<ToastContainer />

<!-- Skeleton 5 removed the AppShell component and asks each project to write
  its own layout. The structure below does what AppShell did for this page.
  A column fills the window. The header and the footer keep their height,
  and the middle section scrolls. -->
{#key $page.url.pathname}
  <div class="relative flex h-full w-full flex-col overflow-hidden">
    <div class="fixed inset-0 bg-gradient-to-br from-primary-500/20 via-tertiary-500/20 to-secondary-500/20 -z-10"></div>
    <header class="flex-none">
      <Header />

      <div class="h-2 py-4"></div>
    </header>
    <div
      class="min-h-0 flex-auto overflow-y-auto"
      in:fly={{ duration: 500, delay: 100, y: 100 }}
    >
      <main class="main m-auto">
        <slot />
      </main>
    </div>

    <footer class="flex-none">
      <Footer />
    </footer>
  </div>
{/key}

<style>
main {
  padding: 2rem;
  box-sizing: border-box;
  overflow-y: auto;
}
</style>
