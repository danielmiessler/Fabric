/// <reference types="svelte" />
import type { DrawerSettings } from './types.js';
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
export declare function getDrawerStore(): DrawerStore;
/**
 * Initializes the `drawerStore`.
 */
export declare function initializeDrawerStore(): DrawerStore;
export type DrawerStore = ReturnType<typeof drawerService>;
declare function drawerService(): {
    subscribe: (this: void, run: import("svelte/store").Subscriber<DrawerSettings>, invalidate?: import("svelte/store").Invalidator<DrawerSettings> | undefined) => import("svelte/store").Unsubscriber;
    set: (this: void, value: DrawerSettings) => void;
    update: (this: void, updater: import("svelte/store").Updater<DrawerSettings>) => void;
    /** Open the drawer. */
    open: (newSettings?: DrawerSettings) => void;
    /** Close the drawer. */
    close: () => void;
};
export {};
