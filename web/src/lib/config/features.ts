import { writable } from 'svelte/store';

interface FeatureFlags {
  enableObsidianIntegration: boolean;
}

export const featureFlags = writable<FeatureFlags>({
  enableObsidianIntegration: true  // Set to true for development
});

export function toggleObsidianIntegration(enabled: boolean) {
  featureFlags.update(flags => ({
    ...flags,
    enableObsidianIntegration: enabled
  }));
}