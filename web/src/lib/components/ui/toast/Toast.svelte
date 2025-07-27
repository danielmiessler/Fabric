<script lang="ts">
  import { toastStore } from '$lib/store/toast-store';
  import { fly } from 'svelte/transition';
  import { onMount } from 'svelte';
  import type { ToastMessage } from '$lib/store/toast-store';

  export let toast: ToastMessage;
  const TOAST_TIMEOUT = 5000;

  onMount(() => {
      const timer = setTimeout(() => {
          toastStore.remove(toast.id);
      }, TOAST_TIMEOUT);

      return () => clearTimeout(timer);
  });
</script>

<div
  class="fixed bottom-4 right-4 p-4 rounded-lg shadow-lg"
  class:bg-green-100={toast.type === 'success'}
  class:bg-red-100={toast.type === 'error'}
  class:bg-blue-100={toast.type === 'info'}
  transition:fly={{ y: 200, duration: 300 }}
>
  <p
    class:text-green-800={toast.type === 'success'}
    class:text-red-800={toast.type === 'error'}
    class:text-blue-800={toast.type === 'info'}
  >
    {toast.message}
  </p>
</div>
