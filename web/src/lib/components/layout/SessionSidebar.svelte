<script lang="ts">
  import { supabaseSessions, supabaseSessionsError, supabaseSessionsLoading, loadSupabaseSessions } from '$lib/store/supabase-session-store';
  import { onMount } from 'svelte';
  import { derived } from 'svelte/store';
  import { formatRelative } from 'date-fns';
  import { Loader2, Plus, History } from 'lucide-svelte';
  import { goto } from '$app/navigation';

  const sessions = supabaseSessions;
  const loading = supabaseSessionsLoading;
  const error = supabaseSessionsError;

  const groupedSessions = derived(sessions, ($sessions) => {
    return $sessions.map((session) => {
      let relative: string | null = null;
      if (session.updated_at) {
        try {
          relative = formatRelative(new Date(session.updated_at), new Date());
        } catch (err) {
          relative = null;
        }
      }
      return { ...session, relative };
    });
  });

  onMount(() => {
    loadSupabaseSessions().catch(() => {
      // error handled via store
    });
  });

  function handleCreateSession() {
    goto('/chat');
  }
</script>

<aside class="hidden w-72 shrink-0 border-r border-border/60 bg-background/40 backdrop-blur lg:flex lg:flex-col">
  <div class="flex items-center justify-between border-b border-border/60 px-4 py-3">
    <div class="flex items-center gap-2 text-sm font-semibold uppercase tracking-wide text-muted-foreground">
      <History class="h-4 w-4" />
      Sessions
    </div>
    <button
      class="inline-flex h-8 items-center gap-1 rounded-md border border-border bg-background px-2 text-xs font-medium transition-colors hover:bg-accent hover:text-accent-foreground"
      on:click={handleCreateSession}
      type="button"
    >
      <Plus class="h-4 w-4" />
      New
    </button>
  </div>

  {#if $loading}
    <div class="flex flex-1 items-center justify-center text-muted-foreground">
      <Loader2 class="mr-2 h-4 w-4 animate-spin" /> Loading sessions
    </div>
  {:else if $error}
    <div class="flex flex-1 flex-col items-center justify-center gap-2 px-4 text-center text-sm text-destructive">
      <p>{$error}</p>
      <button
        class="inline-flex h-8 items-center rounded-md border border-border bg-background px-3 text-xs font-medium transition-colors hover:bg-accent hover:text-accent-foreground"
        on:click={() => loadSupabaseSessions()}
        type="button"
      >
        Try again
      </button>
    </div>
  {:else}
    <div class="flex-1 overflow-y-auto">
      <ul class="grid gap-1 px-3 py-4 text-sm">
        {#each $groupedSessions as session}
          <li>
            <button
              class="group flex w-full flex-col gap-0.5 rounded-md border border-transparent px-3 py-2 text-left transition hover:border-border hover:bg-accent/40 hover:text-foreground"
              type="button"
              on:click={() => goto(`/chat?session=${session.id}`)}
            >
              <span class="line-clamp-1 font-medium group-hover:text-foreground">{session.title ?? 'Untitled session'}</span>
              {#if session.description}
                <span class="line-clamp-2 text-xs text-muted-foreground group-hover:text-foreground/80">{session.description}</span>
              {/if}
              {#if session.relative}
                <span class="text-[0.7rem] uppercase text-muted-foreground/80">Updated {session.relative}</span>
              {/if}
            </button>
          </li>
        {/each}
        {#if !$groupedSessions.length}
          <li class="rounded-md border border-dashed border-border/70 px-3 py-6 text-center text-xs text-muted-foreground">
            No sessions synced from Supabase yet.
          </li>
        {/if}
      </ul>
    </div>
  {/if}
</aside>
