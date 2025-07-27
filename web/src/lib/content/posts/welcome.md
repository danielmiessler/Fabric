---
title: Welcome to Your Blog
description: First post of your new SvelteKit blog
date: 2024-01-17
tags: 
  - welcome
  - blog
updated: 2024-01-17
author: Your Name Here
aliases: 
  - Welcome!
---
<script>
  import { Button } from '$lib/components/ui/button';
  import NoteDrawer from '$lib/components/ui/noteDrawer/NoteDrawer.svelte';
  import { getDrawerStore } from '@skeletonlabs/skeleton';
  import { page } from '$app/stores';
  import { beforeNavigate } from '$app/navigation';
  
  const drawerStore = getDrawerStore();
  function openDrawer() {
    drawerStore.open({});
  }

  beforeNavigate(() => {
    drawerStore.close();
  });

  $: isVisible = $page.url.pathname.startsWith('/welcome');
</script>

This is the first post of your new blog, powered by [SvelteKit](/posts/getting-started), [Obsidian](/obsidian), and [Fabric](/about). We are excited to share this project with you and we hope you find it useful for your own writing and experiences.

**Get started:**
<div class="flex text-inherit justify-start mt-2">
    <Button
        variant="primary"
        class="btn border variant-filled-primary text-align-center"
        on:click={openDrawer}
    >Open Drawer
    </Button>
</div>
<NoteDrawer />

This part of the application is edited in <a href="http://localhost:5173/posts/obsidian" name="Obsidian">Obsidian</a>.

## What to Expect

- Updates on Incorporating Fabric into your workflow
- How to use Obsidian to manage you notes and workflows
- How to use Fabric and Obsidian to write and publish
- More ways to use Obsidian and Fabric together!

Stay tuned for more content! 


 
