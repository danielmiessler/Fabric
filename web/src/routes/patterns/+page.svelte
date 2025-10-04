<script lang="ts">
  import { onMount } from 'svelte';
  import { formatDistanceToNow } from 'date-fns';
  import Modal from '$lib/components/ui/modal/Modal.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Checkbox } from '$lib/components/ui/checkbox/Checkbox.svelte';
  import {
    supabasePatterns,
    supabasePatternsLoading,
    supabasePatternsError,
    loadSupabasePatterns,
    createSupabasePattern,
    updateSupabasePattern,
    deleteSupabasePattern,
    type SupabasePattern,
    type SupabasePatternPayload
  } from '$lib/store/supabase-patterns-store';
  import { setSystemPrompt, selectedPatternName } from '$lib/store/pattern-store';
  import { toastStore } from '$lib/store/toast-store';

  const emptyForm = () => ({
    name: '',
    description: '',
    body: '',
    tagsText: '',
    isSystem: false
  });

  let form = emptyForm();
  let showCreateModal = false;
  let showEditModal = false;
  let editingPattern: SupabasePattern | null = null;
  let searchQuery = '';
  let tagQuery = '';

  onMount(() => {
    loadSupabasePatterns().catch((err) => {
      console.error('Failed to load Supabase patterns', err);
      toastStore.error('Unable to load Supabase patterns');
    });
  });

  const tagsFromText = (value: string): string[] =>
    value
      .split(',')
      .map((tag) => tag.trim())
      .filter(Boolean);

  const resetForm = () => {
    form = emptyForm();
    editingPattern = null;
  };

  const openCreateModal = () => {
    resetForm();
    showCreateModal = true;
  };

  const openEditModal = (pattern: SupabasePattern) => {
    editingPattern = pattern;
    form = {
      name: pattern.name,
      description: pattern.description ?? '',
      body: pattern.body,
      tagsText: pattern.tags?.join(', ') ?? '',
      isSystem: pattern.is_system
    };
    showEditModal = true;
  };

  const closeModals = () => {
    showCreateModal = false;
    showEditModal = false;
    resetForm();
  };

  const toPayload = (): SupabasePatternPayload => ({
    name: form.name.trim(),
    description: form.description?.trim() || undefined,
    body: form.body,
    tags: tagsFromText(form.tagsText),
    is_system: form.isSystem
  });

  const handleCreate = async () => {
    try {
      const payload = toPayload();
      if (!payload.name || !payload.body) {
        toastStore.error('Name and body are required');
        return;
      }
      await createSupabasePattern(payload);
      toastStore.success('Pattern created');
      closeModals();
    } catch (err) {
      console.error(err);
      toastStore.error(err instanceof Error ? err.message : 'Unable to create pattern');
    }
  };

  const handleUpdate = async () => {
    if (!editingPattern) return;
    try {
      const payload = toPayload();
      if (!payload.name || !payload.body) {
        toastStore.error('Name and body are required');
        return;
      }
      const updated = await updateSupabasePattern(editingPattern.id, payload);
      if (!updated) {
        toastStore.error('Pattern no longer exists');
      } else {
        toastStore.success('Pattern updated');
      }
      closeModals();
    } catch (err) {
      console.error(err);
      toastStore.error(err instanceof Error ? err.message : 'Unable to update pattern');
    }
  };

  const handleDelete = async (pattern: SupabasePattern) => {
    const confirmed = confirm(`Delete pattern \"${pattern.name}\"?`);
    if (!confirmed) return;
    try {
      await deleteSupabasePattern(pattern.id);
      toastStore.info('Pattern deleted');
    } catch (err) {
      console.error(err);
      toastStore.error(err instanceof Error ? err.message : 'Unable to delete pattern');
    }
  };

  const handleUsePattern = (pattern: SupabasePattern) => {
    setSystemPrompt(pattern.body);
    selectedPatternName.set(`supabase:${pattern.name}`);
    toastStore.success('Pattern applied to current session');
  };

  const handleCopy = async (pattern: SupabasePattern) => {
    try {
      await navigator.clipboard.writeText(pattern.body);
      toastStore.success('Pattern copied to clipboard');
    } catch (err) {
      console.error(err);
      toastStore.error('Failed to copy pattern');
    }
  };

  const formatTimestamp = (pattern: SupabasePattern): string => {
    const iso = pattern.updated_at ?? pattern.created_at;
    if (!iso) return 'just now';
    try {
      return formatDistanceToNow(new Date(iso), { addSuffix: true });
    } catch (err) {
      console.error('Failed to format timestamp', err);
      return 'recently';
    }
  };

  $: activeTags = tagQuery
    .split(',')
    .map((tag) => tag.trim().toLowerCase())
    .filter(Boolean);

  $: filteredPatterns = $supabasePatterns.filter((pattern) => {
    const matchesSearch = !searchQuery
      || pattern.name.toLowerCase().includes(searchQuery.toLowerCase())
      || pattern.description?.toLowerCase().includes(searchQuery.toLowerCase())
      || pattern.tags?.some((tag) => tag.toLowerCase().includes(searchQuery.toLowerCase()));

    if (!matchesSearch) return false;

    if (activeTags.length === 0) return true;

    const patternTags = pattern.tags?.map((tag) => tag.toLowerCase()) ?? [];
    return activeTags.every((tag) => patternTags.includes(tag));
  });
</script>

<section class="flex flex-col gap-6">
  <header class="flex flex-col gap-2">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-semibold text-foreground">Shared Pattern Library</h1>
        <p class="text-sm text-muted-foreground">Manage Supabase-backed prompts for your team.</p>
      </div>
      <Button on:click={openCreateModal} variant="default">New pattern</Button>
    </div>
    <div class="grid gap-3 md:grid-cols-3">
      <Input bind:value={searchQuery} placeholder="Search by name, description, or tag" />
      <Input
        bind:value={tagQuery}
        placeholder="Filter by tags (comma separated)"
        class="md:col-span-2"
      />
    </div>
  </header>

  {#if $supabasePatternsLoading}
    <div class="flex h-48 items-center justify-center text-muted-foreground">
      Loading patterns…
    </div>
  {:else if $supabasePatternsError}
    <div class="rounded-md border border-destructive/40 bg-destructive/10 p-4 text-sm text-destructive">
      {$supabasePatternsError}
    </div>
  {:else if filteredPatterns.length === 0}
    <div class="rounded-md border border-border/60 bg-background/40 p-6 text-center text-muted-foreground">
      No patterns found. Try adjusting your search or add a new pattern.
    </div>
  {:else}
    <div class="grid gap-4">
      {#each filteredPatterns as pattern}
        <article class="rounded-lg border border-border/60 bg-background/60 p-4 shadow-sm transition hover:border-primary/50">
          <div class="flex flex-col gap-3 md:flex-row md:items-start md:justify-between">
            <div class="space-y-2">
              <div class="flex items-center gap-3">
                <h2 class="text-xl font-semibold text-foreground">{pattern.name}</h2>
                {#if pattern.is_system}
                  <span class="rounded-full bg-primary/20 px-2 py-0.5 text-xs font-medium text-primary-200">
                    system
                  </span>
                {/if}
              </div>
              {#if pattern.description}
                <p class="text-sm text-muted-foreground">{pattern.description}</p>
              {/if}
              {#if pattern.tags?.length}
                <div class="flex flex-wrap gap-2 text-xs text-muted-foreground">
                  {#each pattern.tags as tag}
                    <span class="rounded-full bg-muted/40 px-2 py-0.5 text-muted-foreground">{tag}</span>
                  {/each}
                </div>
              {/if}
              <small class="text-xs text-muted-foreground">
                Updated {formatTimestamp(pattern)}
              </small>
            </div>
            <div class="flex flex-col gap-2 md:items-end">
              <div class="flex gap-2">
                <Button variant="secondary" on:click={() => handleUsePattern(pattern)}>Use</Button>
                <Button variant="outline" on:click={() => handleCopy(pattern)}>Copy</Button>
                <Button variant="outline" on:click={() => openEditModal(pattern)}>Edit</Button>
                <Button variant="destructive" on:click={() => handleDelete(pattern)}>Delete</Button>
              </div>
            </div>
          </div>
        </article>
      {/each}
    </div>
  {/if}

  <Modal bind:show={showCreateModal} on:close={closeModals}>
    <div class="w-[32rem] max-w-full rounded-xl border border-border/60 bg-background p-6 shadow-xl">
      <header class="mb-4">
        <h2 class="text-xl font-semibold text-foreground">New pattern</h2>
        <p class="text-sm text-muted-foreground">Create a shared prompt stored in Supabase.</p>
      </header>
      <div class="flex flex-col gap-4">
        <Input bind:value={form.name} placeholder="Pattern name" />
        <Input bind:value={form.description} placeholder="Short description" />
        <Textarea
          bind:value={form.body}
          rows={8}
          placeholder="Prompt body"
          class="resize-y"
        />
        <Input bind:value={form.tagsText} placeholder="Tags (comma separated)" />
        <label class="flex items-center gap-2 text-sm text-foreground">
          <Checkbox bind:checked={form.isSystem} />
          Mark as system prompt
        </label>
        <div class="flex justify-end gap-2">
          <Button variant="outline" on:click={closeModals}>Cancel</Button>
          <Button on:click={handleCreate}>Create</Button>
        </div>
      </div>
    </div>
  </Modal>

  <Modal bind:show={showEditModal} on:close={closeModals}>
    <div class="w-[32rem] max-w-full rounded-xl border border-border/60 bg-background p-6 shadow-xl">
      <header class="mb-4">
        <h2 class="text-xl font-semibold text-foreground">Edit pattern</h2>
        <p class="text-sm text-muted-foreground">Update details for this shared prompt.</p>
      </header>
      <div class="flex flex-col gap-4">
        <Input bind:value={form.name} placeholder="Pattern name" />
        <Input bind:value={form.description} placeholder="Short description" />
        <Textarea
          bind:value={form.body}
          rows={8}
          placeholder="Prompt body"
          class="resize-y"
        />
        <Input bind:value={form.tagsText} placeholder="Tags (comma separated)" />
        <label class="flex items-center gap-2 text-sm text-foreground">
          <Checkbox bind:checked={form.isSystem} />
          Mark as system prompt
        </label>
        <div class="flex justify-end gap-2">
          <Button variant="outline" on:click={closeModals}>Cancel</Button>
          <Button on:click={handleUpdate}>Save changes</Button>
        </div>
      </div>
    </div>
  </Modal>
</section>
