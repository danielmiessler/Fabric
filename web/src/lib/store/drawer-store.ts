import { writable } from 'svelte/store';

// Skeleton 5 removed the drawer utility and its store. This store keeps the
// small part of the Skeleton v2 API that this project used, so that the call
// sites stay the same: drawerStore.open(), drawerStore.close(), and a read of
// $drawerStore.open.

export interface DrawerSettings {
  id?: string;
  meta?: unknown;
}

export interface DrawerState extends DrawerSettings {
  open: boolean;
}

function createDrawerStore() {
  const { subscribe, set, update } = writable<DrawerState>({ open: false });

  return {
    subscribe,
    open: (settings: DrawerSettings = {}) => set({ ...settings, open: true }),
    close: () => update((state) => ({ ...state, open: false }))
  };
}

export const drawerStore = createDrawerStore();
