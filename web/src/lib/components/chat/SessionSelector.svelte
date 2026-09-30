<script lang="ts">
  import { Select } from "$lib/components/ui/select";
  import { Label } from "$lib/components/ui/label";
  import { currentSession, setSession, messageStore } from '$lib/store/chat-store';
  import { sessionAPI, sessions } from '$lib/store/session-store';
  import { onMount } from 'svelte';

  let sessionInput = '';

  $: sessionsList = $sessions?.map(s => s.Name) ?? [];

  function handleSessionInput() {
    const trimmed = sessionInput.trim();
    if (trimmed) {
      setSession(trimmed);
    } else {
      sessionInput = '';
      setSession(null);
    }
  }

  let previousSessionInput = '';

  async function handleSessionSelect() {
    // Selecting the placeholder option restores the previous value.
    if (!sessionInput) {
      sessionInput = previousSessionInput || $currentSession || '';
      return;
    }

    if (sessionInput === $currentSession) {
      return;
    }

    previousSessionInput = sessionInput;
    setSession(sessionInput);

    // Load the session history so the chat shows the earlier messages.
    try {
      const messages = await sessionAPI.loadSessionMessages(sessionInput);
      messageStore.set(messages);
    } catch (error) {
      console.error('Failed to load session messages:', error);
    }
  }

  onMount(async () => {
    try {
      await sessionAPI.loadSessions();
    } catch (error) {
      console.error('Failed to load sessions:', error);
    }
    sessionInput = $currentSession ?? '';
  });
</script>

<div>
  <Label for="session-input" class="text-xs text-white/70 mb-1 block">Session Name</Label>
  <input
    id="session-input"
    type="text"
    bind:value={sessionInput}
    on:blur={handleSessionInput}
    on:keydown={(e) => e.key === 'Enter' && handleSessionInput()}
    placeholder="Enter session name..."
    class="w-full px-3 py-2 text-sm bg-primary-800/30 border-none rounded-md hover:bg-primary-800/40 transition-colors text-white placeholder-white/50 focus:ring-1 focus:ring-white/20 focus:outline-none"
  />
  {#if sessionsList.length > 0}
    <Select
      bind:value={sessionInput}
      on:change={handleSessionSelect}
      class="mt-2 bg-primary-800/30 border-none hover:bg-primary-800/40 transition-colors"
    >
      <option value="">Load existing session...</option>
      {#each sessionsList as session}
        <option value={session}>{session}</option>
      {/each}
    </Select>
  {/if}
</div>
