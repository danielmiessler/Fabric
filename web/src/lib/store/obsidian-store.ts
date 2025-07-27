import { writable, get } from 'svelte/store';
import { featureFlags } from '../config/features';

export interface ObsidianSettings {
  saveToObsidian: boolean;
  noteName: string;
}

// Keep existing defaultSettings
const defaultSettings: ObsidianSettings = {
  saveToObsidian: false,
  noteName: ''
};

// Keep existing store initialization
export const obsidianSettings = writable<ObsidianSettings>(defaultSettings);

// Add notification store
export const saveNotification = writable<string>('');

// Keep existing update function with notification enhancement
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
    
    // Add notification after successful save
    if (settings.saveToObsidian === false && current.noteName) {
      saveNotification.set('Note saved to Obsidian!');
      setTimeout(() => saveNotification.set(''), 3000);
    }
    
    console.log('Updated Obsidian settings:', updated);
    return updated;
  });
}

// Reset settings to default
export function resetObsidianSettings() {
  const enabled = get(featureFlags).enableObsidianIntegration;
  if (!enabled) return;
  
  obsidianSettings.set(defaultSettings);
}

// Helper to get file path
export function getObsidianFilePath(noteName: string): string | undefined {
  const enabled = get(featureFlags).enableObsidianIntegration;
  if (!enabled || !noteName) return undefined;

  return `myfiles/Fabric_obsidian/${
    new Date().toISOString().split('T')[0]
  }-${noteName.trim()}.md`;
  
  
  
}

