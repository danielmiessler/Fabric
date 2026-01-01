/// <reference types="svelte" />
import type { ToastSettings, Toast } from './types.js';
/**
 * Retrieves the `toastStore`.
 *
 * This can *ONLY* be called from the **top level** of components!
 *
 * @example
 * ```svelte
 * <script>
 * 	import { getToastStore } from "@skeletonlabs/skeleton";
 *
 * 	const toastStore = getToastStore();
 *
 * 	toastStore.open({ message: "Welcome!" });
 * </script>
 * ```
 */
export declare function getToastStore(): ToastStore;
/**
 * Initializes the `toastStore`.
 */
export declare function initializeToastStore(): ToastStore;
export type ToastStore = ReturnType<typeof toastService>;
declare function toastService(): {
    subscribe: (this: void, run: import("svelte/store").Subscriber<Toast[]>, invalidate?: import("svelte/store").Invalidator<Toast[]> | undefined) => import("svelte/store").Unsubscriber;
    close: (id: string) => void;
    /** Add a new toast to the queue. */
    trigger: (toast: ToastSettings) => string;
    /** Remain visible on hover */
    freeze: (index: number) => void;
    /** Cancel remain visible on leave */
    unfreeze: (index: number) => void;
    /** Remove all toasts from queue */
    clear: () => void;
};
export {};
