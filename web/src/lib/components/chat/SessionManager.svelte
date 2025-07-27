<script lang="ts">
  import { onMount } from 'svelte';
  import { RotateCcw, Trash2, Save, Copy, File as FileIcon } from 'lucide-svelte';
  import { sessions, sessionAPI } from '$lib/store/session-store';
  import { chatState, clearMessages, revertLastMessage, currentSession, messageStore } from '$lib/store/chat-store';
  import { Button } from '$lib/components/ui/button';
  import { toastService } from '$lib/services/toast-service';
  
  let sessionsList: string[] = [];
  $: sessionName = $currentSession;
  $: if ($sessions) {
    sessionsList = $sessions.map(s => s.Name);
  }
  
  onMount(async () => {
    try {
      await sessionAPI.loadSessions();
    } catch (error) {
      console.error('Failed to load sessions:', error);
    }
  });
  
  async function saveSession() {
    try {
      await sessionAPI.exportToFile($chatState.messages);
    } catch (error) {
      console.error('Failed to save session:', error);
    }
  }

  async function loadSession() {
    try {
      const messages = await sessionAPI.importFromFile();
      messageStore.set(messages);
    } catch (error) {
      console.error('Failed to load session:', error);
    }
  }

  async function copyToClipboard() {
    try {
      await navigator.clipboard.writeText($chatState.messages.map(m => m.content).join('\n'));
      toastService.success('Chat copied to clipboard!');
    } catch (err) {
      toastService.error('Failed to copy transcript');
    }
  }
</script>

<div class="p-1 m-1 mr-2">
  <div class="flex gap-2">
    <Button variant="outline" size="icon" aria-label="Revert Last Message" on:click={revertLastMessage}>
        <RotateCcw class="h-4 w-4" />
    </Button>
    <Button variant="outline" size="icon" aria-label="Clear Chat" on:click={clearMessages}>
        <Trash2 class="h-4 w-4" />
    </Button>
    <Button variant="outline" size="icon" aria-label="Copy Chat" on:click={copyToClipboard}>
        <Copy class="h-4 w-4" />
    </Button>
    <Button variant="outline" size="icon" aria-label="Load Session" on:click={loadSession}>
        <FileIcon class="h-4 w-4" />
    </Button>
    <Button variant="outline" size="icon" aria-label="Save Session" on:click={saveSession}>
        <Save class="h-4 w-4" />
    </Button>
  </div>
</div>
