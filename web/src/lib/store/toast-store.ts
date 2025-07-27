import { writable } from 'svelte/store';

export interface ToastMessage {
  message: string;
  type: 'success' | 'error' | 'info';
  id: number;
}

function createToastStore() {
  const { subscribe, update } = writable<ToastMessage[]>([]);
  let nextId = 1;

  return {
    subscribe,
    success: (message: string) => {
      update(toasts => [...toasts, { message, type: 'success', id: nextId++ }]);
    },
    error: (message: string) => {
      update(toasts => [...toasts, { message, type: 'error', id: nextId++ }]);
    },
    info: (message: string) => {
      update(toasts => [...toasts, { message, type: 'info', id: nextId++ }]);
    },
    remove: (id: number) => {
      update(toasts => toasts.filter(t => t.id !== id));
    }
  };
}

export const toastStore = createToastStore();
