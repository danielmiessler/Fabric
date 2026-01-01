<script>import { initializeToastStore, getToastStore } from "./stores.js";
import Toast from "./Toast.svelte";
export let toastSettings = [];
export let max = void 0;
initializeToastStore();
const toastStore = getToastStore();
toastSettings.forEach(({ triggerDelay, ...settings }) => {
  if (triggerDelay) {
    setTimeout(() => toastStore.trigger(settings), triggerDelay);
  } else {
    toastStore.trigger(settings);
  }
});
</script>

<Toast {max} />
