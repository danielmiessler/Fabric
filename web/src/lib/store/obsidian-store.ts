import { writable, get } from 'svelte/store';
import { featureFlags } from '../config/features';

export interface ObsidianSettings {
  saveToObsidian: boolean;
  noteName: string;
}

const defaultSettings: ObsidianSettings = {
  saveToObsidian: false,
  noteName: ''
};

export const obsidianSettings = writable<ObsidianSettings>(defaultSettings);

export const saveNotification = writable<string>('');

export function updateObsidianSettings(settings: Partial<ObsidianSettings>) {
  const enabled = get(featureFlags).enableObsidianIntegration;
  console.log('Updating Obsidian settings:', settings, 'Integration enabled:', enabled);
  
  if (!enabled) {
    console.log('Obsidian integration disabled, not updating settings');
    return;
  }
  
  obsidianSettings.update(current => {
    const updated = {
      ...current,
      ...settings
    };
    
    // ChatInput sets saveToObsidian to false after a successful save.
    if (settings.saveToObsidian === false && current.noteName) {
      saveNotification.set('Note saved to Obsidian!');
      setTimeout(() => saveNotification.set(''), 3000);
    }
    
    console.log('Updated Obsidian settings:', updated);
    return updated;
  });
}

export function resetObsidianSettings() {
  const enabled = get(featureFlags).enableObsidianIntegration;
  if (!enabled) return;
  
  obsidianSettings.set(defaultSettings);
}

export function getObsidianFilePath(noteName: string): string | undefined {
  const enabled = get(featureFlags).enableObsidianIntegration;
  if (!enabled || !noteName) return undefined;

  return `myfiles/Fabric_obsidian/${
    new Date().toISOString().split('T')[0]
  }-${noteName.trim()}.md`;
  
  
  
}

