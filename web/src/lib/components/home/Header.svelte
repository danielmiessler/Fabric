<script lang="ts">
  import { page } from '$app/stores';
  import { Sun, Moon, Menu, X, Github, FileText } from 'lucide-svelte';
  import { Avatar } from '@skeletonlabs/skeleton';
  import { fade } from 'svelte/transition';
  import { theme, cycleTheme, initTheme } from '$lib/store/theme-store';
  import { onMount } from 'svelte';
  import Modal from '$lib/components/ui/modal/Modal.svelte';
  import PatternList from '$lib/components/patterns/PatternList.svelte';
  import PatternTilesModal from '$lib/components/ui/modal/PatternTilesModal.svelte';
  import HelpModal from '$lib/components/ui/help/HelpModal.svelte';
  import { selectedPatternName } from '$lib/store/pattern-store';

  let isMenuOpen = false;
  let showPatternModal = false;
  let showPatternTilesModal = false;
  let showHelpModal = false;

  function goToGithub() {
    window.open('https://github.com/danielmiessler/fabric', '_blank');
  }

  function toggleMenu() {
    isMenuOpen = !isMenuOpen;
  }

  $: currentPath = $page.url.pathname;
  $: isDarkMode = $theme === 'my-custom-theme';

  const navItems = [
    { href: '/', label: 'Home' },
    { href: '/posts', label: 'Posts' },
    // { href: '/tags', label: 'Tags' },
    { href: '/chat', label: 'Chat' },
    //{ href: '/obsidian', label: 'Obsidian' },
    { href: '/contact', label: 'Contact' },
    { href: '/about', label: 'About' },
  ];

  onMount(() => {
    initTheme();
  }); 
</script>

<header class="fixed top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
  <div class="container flex h-16 items-center justify-between px-4">
    <div class="flex items-center gap-4">
      <Avatar 
        src="/fabric-logo.png" 
        width="w-10" 
        rounded="rounded-full" 
        class="border-2 border-primary/20"
      />
      <a href="/" class="flex items-center">
        <span class="text-lg font-semibold">fabric</span>
      </a>
    </div>

    <!-- Desktop Navigation -->
    <nav class="hidden flex-1 px-8 md:flex">
      <ul class="flex items-center space-x-8">
        {#each navItems as { href, label }}
          <li>
            <a
              {href}
              class="text-sm font-medium transition-colors hover:text-primary {currentPath === href ? 'text-primary' : 'text-foreground/60'}"
            >
              {label}
            </a>
          </li>
        {/each}
      </ul>
    </nav>

    <div class="flex items-center gap-4">
      <!-- Pattern Buttons Group -->
      <div class="flex items-center gap-3 mr-4">
        <!-- Pattern Tiles Button -->
        <button name="pattern-tiles"
          on:click={() => showPatternTilesModal = true}
          class="inline-flex h-10 items-center justify-center rounded-full border bg-background px-4 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground gap-2"
          aria-label="Pattern Tiles"
        >
          <FileText class="h-4 w-4" />
          <span>Pattern Tiles</span>
        </button>
        
        <!-- Or text -->
        <span class="text-sm text-foreground/60 mx-1">or</span>
        
        <!-- Pattern List Button -->
        <button name="pattern-list"
          on:click={() => showPatternModal = true}
          class="inline-flex h-10 items-center justify-center rounded-full border bg-background px-4 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground gap-2"
          aria-label="Pattern List"
        >
          <FileText class="h-4 w-4" />
          <span>Pattern List</span>
        </button>
      </div>


      <button name="github"
        on:click={goToGithub}
        class="inline-flex h-9 w-9 items-center justify-center rounded-full border bg-background text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground"
        aria-label="GitHub"
      >
        <Github class="h-4 w-4" />
        <span class="sr-only">GitHub</span>
      </button>

      <button name="toggle-theme"
        on:click={cycleTheme}
        class="inline-flex h-9 w-9 items-center justify-center rounded-full border bg-background text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground"
        aria-label="Toggle theme"
      >
        {#if isDarkMode}
          <Sun class="h-4 w-4" />
        {:else}
          <Moon class="h-4 w-4" />
        {/if}
        <span class="sr-only">Toggle theme</span>
      </button>

      <button name="help"
        on:click={() => showHelpModal = true}
        class="inline-flex h-9 w-9 items-center justify-center rounded-full border bg-background text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground ml-3"
        aria-label="Help"
      >
        <span class="text-xl font-bold text-white/90 hover:text-white">?</span>
        <span class="sr-only">Help</span>
      </button>

      <!-- Mobile Menu Button -->
      <button name="toggle-menu"
        class="inline-flex h-9 w-9 items-center justify-center rounded-lg border bg-background text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground md:hidden"
        on:click={toggleMenu}
        aria-expanded={isMenuOpen}
        aria-label="Toggle menu"
      >
        {#if isMenuOpen}
          <X class="h-4 w-4" />
        {:else}
          <Menu class="h-4 w-4" />
        {/if}
      </button>
    </div>
  </div>

  <!-- Mobile Navigation -->
  {#if isMenuOpen}
    <div class="container md:hidden" transition:fade={{ duration: 200 }}>
      <nav class="flex flex-col space-y-4 p-4">
        {#each navItems as { href, label }}
          <a
            {href}
            class="text-base font-medium transition-colors hover:text-primary {currentPath === href ? 'text-primary' : 'text-foreground/60'}"
            on:click={() => (isMenuOpen = false)}
          >
            {label}
          </a>
        {/each}
      </nav>
    </div>
  {/if}
</header>

<Modal
  show={showPatternModal}
  on:close={() => showPatternModal = false}
>
  <PatternList
    on:close={() => showPatternModal = false}
    on:select={(e) => {
      selectedPatternName.set(e.detail);
      showPatternModal = false;
    }}
  />
</Modal>

<Modal
  show={showHelpModal}
  on:close={() => showHelpModal = false}
>
  <HelpModal
    on:close={() => showHelpModal = false}
  />
</Modal>

<Modal
  show={showPatternTilesModal}
  on:close={() => showPatternTilesModal = false}
>
  <PatternTilesModal
    on:close={() => showPatternTilesModal = false}
    on:select={(e) => {
      selectedPatternName.set(e.detail);
      showPatternTilesModal = false;
    }}
  />
</Modal>
