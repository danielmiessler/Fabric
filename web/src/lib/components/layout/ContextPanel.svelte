<script lang="ts">
  import { supabaseSessions } from '$lib/store/supabase-session-store';
  import { Lightbulb, Database, Cog } from 'lucide-svelte';
</script>

<aside class="hidden w-80 shrink-0 border-l border-border/60 bg-background/30 backdrop-blur xl:flex xl:flex-col">
  <div class="border-b border-border/60 px-4 py-3">
    <div class="text-sm font-semibold uppercase tracking-wide text-muted-foreground">Context</div>
  </div>

  <div class="flex-1 overflow-y-auto px-4 py-6">
    <section class="mb-6">
      <header class="mb-2 flex items-center gap-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
        <Lightbulb class="h-3.5 w-3.5" /> Active Session
      </header>
      <div class="rounded-lg border border-border/60 bg-background/80 p-3 text-sm text-muted-foreground">
        <p>Select a conversation to inspect its notes and contexts.</p>
        <p class="mt-2 text-xs uppercase text-muted-foreground/80">Sessions available: {$supabaseSessions.length}</p>
      </div>
    </section>

    <section class="mb-6">
      <header class="mb-2 flex items-center gap-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
        <Database class="h-3.5 w-3.5" /> Supabase
      </header>
      <div class="space-y-2 text-sm text-muted-foreground">
        <p>Patterns, sessions, and notes can be synchronized from Supabase for collaborative workflows.</p>
        <p class="text-xs text-muted-foreground/70">Connection status updates will appear here.</p>
      </div>
    </section>

    <section>
      <header class="mb-2 flex items-center gap-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
        <Cog class="h-3.5 w-3.5" /> Quick Actions
      </header>
      <div class="grid gap-2 text-sm">
        <button type="button" class="rounded-md border border-border/60 bg-background px-3 py-2 text-left transition hover:bg-accent hover:text-accent-foreground">
          Configure models
        </button>
        <button type="button" class="rounded-md border border-border/60 bg-background px-3 py-2 text-left transition hover:bg-accent hover:text-accent-foreground">
          Manage strategies
        </button>
        <button type="button" class="rounded-md border border-border/60 bg-background px-3 py-2 text-left transition hover:bg-accent hover:text-accent-foreground">
          Review insights
        </button>
      </div>
    </section>
  </div>
</aside>
