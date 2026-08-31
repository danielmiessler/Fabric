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

<!-- ToastContainer holds the position for the group. This element must stay in
  the flow of that container, because a fixed position here would put every
  toast in the same corner, one on top of another, and would also stop the
  spacing that the container sets between them. -->
<div
  class="p-4 rounded-lg shadow-lg"
  class:bg-green-100={toast.type === 'success'}
  class:bg-red-100={toast.type === 'error'}
  class:bg-amber-100={toast.type === 'warning'}
  class:bg-blue-100={toast.type === 'info'}
  transition:fly={{ y: 200, duration: 300 }}
>
  <p
    class:text-green-800={toast.type === 'success'}
    class:text-red-800={toast.type === 'error'}
    class:text-amber-800={toast.type === 'warning'}
    class:text-blue-800={toast.type === 'info'}
  >
    {toast.message}
  </p>
</div>
