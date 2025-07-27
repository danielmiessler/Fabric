<script lang='ts'>
  import { getToastStore } from '@skeletonlabs/skeleton';
  import { Button } from "$lib/components/ui/button";
  import Input from '$lib/components/ui/input/Input.svelte';
  import { Toast } from '@skeletonlabs/skeleton';

  let url = '';
  let transcript = '';
  let loading = false;
  let error = '';
  let title = '';

  const toastStore = getToastStore();

  async function fetchTranscript() {
    function isValidYouTubeUrl(url: string) {
      const pattern = /^(https?:\/\/)?(www\.)?(youtube\.com|youtu\.be)\/.+$/;
      return pattern.test(url);
    }

    if (!isValidYouTubeUrl(url)) {
      error = 'Please enter a valid YouTube URL';
      toastStore.trigger({
        message: error,
        background: 'variant-filled-error'
      });
      return;
    }

    loading = true;
    error = '';

    try {
      const response = await fetch('/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify({ url })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch transcript');
      }

      const data = await response.json();
      console.log('Parsed response data:', data);

      transcript = data.transcript;
      title = data.title;

    } finally {
      loading = false;
    }
  }

  async function copyToClipboard() {
    try {
      await navigator.clipboard.writeText(transcript);
      toastStore.trigger({
        message: 'Transcript copied to clipboard!',
        background: 'variant-filled-success'
      });
    } catch (err) {
      toastStore.trigger({
        message: 'Failed to copy transcript',
        background: 'variant-filled-error'
      });
    }
  }
</script>


<div class="flex gap-2">
  <Input
    bind:value={url}
    placeholder="YouTube URL"
    class="flex-1 rounded-full border bg-background px-4"
    disabled={loading}
  />
  <Button
    variant="secondary"
    on:click={fetchTranscript}
    disabled={loading || !url}
  >
    {#if loading}
      <div class="spinner-border" />
    {:else}
      Get 
    {/if}
  </Button>
</div>

{#if error}
  <div class="bg-destructive/15 text-destructive rounded-lg p-2">{error}</div>
{/if}

{#if transcript}
  <Toast position="l" />
  <div class="space-y-4 border rounded-lg p-4 bg-muted/50 h-96">
    <div class="flex justify-between items-center">
      <h3 class="text-xs font-semibold">{title || 'Transcript'}</h3>
      <Button
        variant="outline"
        size="sm"
        on:click={copyToClipboard}
      >
        Copy to Clipboard
      </Button>
    </div>
    <textarea
      class="w-full text-xs rounded-md border bg-background px-3 py-2 resize-none h-72"
      readonly
      value={transcript}
    ></textarea>
  </div>
{/if}

<style>
.spinner-border {
  width: 1rem;
  height: 1rem;
  border: 2px solid currentColor;
  border-right-color: transparent;
  border-radius: 50%;
  animation: spin 0.75s linear infinite;
}

@keyframes spin {
from { transform: rotate(0deg); }
to { transform: rotate(360deg); }
}
</style>
