<script lang="ts">
  import { page } from '$app/stores';
  import { derived } from 'svelte/store';
  import { Menu, Github, HelpCircle, Sun, Moon } from 'lucide-svelte';
  import { cycleTheme, initTheme, theme } from '$lib/store/theme-store';
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  const navItems = [
    { href: '/', label: 'Home' },
    { href: '/chat', label: 'Chat' },
    { href: '/patterns', label: 'Library' },
    { href: '/posts', label: 'Knowledge' },
    { href: '/about', label: 'About' },
    { href: '/contact', label: 'Contact' }
  ];

  const currentPath = derived(page, ($page) => $page.url.pathname);

  onMount(() => {
    initTheme();
  });

  let showMobileNav = false;

  function toggleMobileNav() {
    showMobileNav = !showMobileNav;
  }

  function openGithub() {
    window.open('https://github.com/danielmiessler/fabric', '_blank', 'noreferrer');
  }

  function openHelp() {
    goto('/about');
  }
</script>

<header class="sticky top-0 z-40 border-b border-border/60 bg-background/80 backdrop-blur">
  <div class="mx-auto flex h-16 w-full max-w-screen-2xl items-center gap-4 px-4">
    <button
      class="inline-flex h-9 w-9 items-center justify-center rounded-md border border-border/60 bg-background text-sm font-medium transition hover:bg-accent hover:text-accent-foreground lg:hidden"
      type="button"
      on:click={toggleMobileNav}
      aria-expanded={showMobileNav}
      aria-label="Toggle navigation"
    >
      <Menu class="h-4 w-4" />
    </button>

    <a href="/" class="font-semibold uppercase tracking-[0.3em]">fabric</a>

    <nav class="hidden flex-1 items-center gap-1 text-sm font-medium text-muted-foreground lg:flex">
      {#each navItems as item}
        <a
          href={item.href}
          class={`rounded-md px-3 py-2 transition ${$currentPath === item.href ? 'bg-accent/50 text-foreground' : 'hover:text-foreground/90'}`}
        >
          {item.label}
        </a>
      {/each}
    </nav>

    <div class="ml-auto flex items-center gap-2">
      <button
        type="button"
        class="inline-flex h-9 w-9 items-center justify-center rounded-md border border-border/60 bg-background transition hover:bg-accent hover:text-accent-foreground"
        on:click={openGithub}
        aria-label="Open GitHub"
      >
        <Github class="h-4 w-4" />
      </button>
      <button
        type="button"
        class="inline-flex h-9 w-9 items-center justify-center rounded-md border border-border/60 bg-background transition hover:bg-accent hover:text-accent-foreground"
        on:click={cycleTheme}
        aria-label="Toggle theme"
      >
        {#if $theme === 'my-custom-theme'}
          <Sun class="h-4 w-4" />
        {:else}
          <Moon class="h-4 w-4" />
        {/if}
      </button>
      <button
        type="button"
        class="inline-flex h-9 w-9 items-center justify-center rounded-md border border-border/60 bg-background transition hover:bg-accent hover:text-accent-foreground"
        on:click={openHelp}
        aria-label="Help"
      >
        <HelpCircle class="h-4 w-4" />
      </button>
    </div>
  </div>

  {#if showMobileNav}
    <nav class="border-t border-border/60 bg-background/95 px-4 py-3 text-sm font-medium text-muted-foreground lg:hidden">
      <ul class="grid gap-2">
        {#each navItems as item}
          <li>
            <a
              href={item.href}
              class={`block rounded-md px-3 py-2 transition ${$currentPath === item.href ? 'bg-accent/50 text-foreground' : 'hover:text-foreground/90'}`}
              on:click={() => (showMobileNav = false)}
            >
              {item.label}
            </a>
          </li>
        {/each}
      </ul>
    </nav>
  {/if}
</header>
