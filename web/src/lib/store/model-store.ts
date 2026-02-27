import { derived, writable } from 'svelte/store';
import { modelsApi } from '$lib/api/models';
import { configApi } from '$lib/api/config';
import type { VendorModel, ModelConfig } from '$lib/interfaces/model-interface';

export const modelConfig = writable<ModelConfig>({
  model: '',
  temperature: 0.7,
  maxLength: 2000,
  top_p: 0.9,
  frequency: 0.5,
  presence: 0
});

export const availableModels = writable<VendorModel[]>([]);
export const selectedVendor = writable<string>('');

// Sorted unique vendor names
export const vendorNames = derived(availableModels, ($models) =>
  [...new Set($models.map(m => m.vendor))].sort((a, b) =>
    a.toLowerCase().localeCompare(b.toLowerCase())
  )
);

// Models filtered by selected vendor
export const filteredModels = derived(
  [availableModels, selectedVendor],
  ([$models, $vendor]) =>
    $vendor ? $models.filter(m => m.vendor === $vendor) : $models
);

// Initialize available models
export async function loadAvailableModels() {
  try {
    const models = await modelsApi.getAvailable();
    console.log('Load models:', models);
    const uniqueModels = [...new Map(models.map(model => [`${model.vendor}:${model.name}`, model])).values()];
    // Sort like CLI: vendor name (case-insensitive), then model name (case-insensitive)
    uniqueModels.sort((a, b) => {
      const vendorCmp = a.vendor.toLowerCase().localeCompare(b.vendor.toLowerCase());
      if (vendorCmp !== 0) return vendorCmp;
      return a.name.toLowerCase().localeCompare(b.name.toLowerCase());
    });
    availableModels.set(uniqueModels);
  } catch (error) {
    console.error('Client failed to load available models:', error);
    availableModels.set([]);
  }
}

// Initialize config
export async function initializeConfig() {
  try {
    const config = await configApi.get();
    modelConfig.set(config);
  } catch (error) {
    console.error('Failed to load config:', error);
  }
}
