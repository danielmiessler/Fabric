import { writable } from 'svelte/store';
import type { ChatConfig } from '$lib/interfaces/chat-interface';

const defaultConfig: ChatConfig = {
  temperature: 0.7,
  top_p: 1,
  frequency_penalty: 0,
  presence_penalty: 0
};

export const chatConfig = writable<ChatConfig>(defaultConfig);

export function updateConfig(newConfig: Partial<ChatConfig>): void {
  chatConfig.update(config => ({
    ...config,
    ...newConfig
  }));
}

export function resetConfig(): void {
  chatConfig.set(defaultConfig);
}

export { type ChatConfig };
