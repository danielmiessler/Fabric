// Drawer Stores
import { writable } from 'svelte/store';
import { getContext, setContext } from 'svelte';
const DRAWER_STORE_KEY = 'drawerStore';
/**
 * Retrieves the `drawerStore`.
 *
 * This can *ONLY* be called from the **top level** of components!
 *
 * @example
 * ```svelte
 * <script>
 * 	import { getDrawerStore } from "@skeletonlabs/skeleton";
 *
 * 	const drawerStore = getDrawerStore();
 *
 * 	drawerStore.open();
 * </script>
 * ```
 */
export function getDrawerStore() {
    const drawerStore = getContext(DRAWER_STORE_KEY);
    if (!drawerStore)
        throw new Error('drawerStore is not initialized. Please ensure that `initializeStores()` is invoked in the root layout file of this app!');
    return drawerStore;
}
/**
 * Initializes the `drawerStore`.
 */
export function initializeDrawerStore() {
    const drawerStore = drawerService();
    return setContext(DRAWER_STORE_KEY, drawerStore);
}
function drawerService() {
    const { subscribe, set, update } = writable({});
    return {
        subscribe,
        set,
        update,
        /** Open the drawer. */
        open: (newSettings) => update(() => {
            return { open: true, ...newSettings };
        }),
        /** Close the drawer. */
        close: () => update((d) => {
            d.open = false;
            return d;
        })
    };
}
