<script lang="ts">
  import { toastStore } from '$lib/store/toast-store';
  import Toast from './Toast.svelte';
</script>

<div class="fixed bottom-0 right-0 p-4 space-y-2 z-50">
  {#each $toastStore as toast (toast.id)}
      <Toast {toast} />
  {/each}
</div>
