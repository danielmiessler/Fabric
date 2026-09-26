import { api } from './base';
import type { VendorModel, ModelsResponse } from '$lib/interfaces/model-interface';

export const modelsApi = {
  async getAvailable(): Promise<VendorModel[]> {
    try {
      const response = await api.fetch<ModelsResponse>('/models/names');
      
      if (!response.data?.vendors) {
        throw new Error('Invalid response format: missing vendors data');
      }
      
      // The server sends null for the model list of a vendor that it can find
      // no models for, because an empty slice in Go becomes null in JSON.
      // Ollama does this when it is in the configuration but serves no models.
      // Skip such a vendor: one of them must not hide the models of the others.
      return Object.entries(response.data.vendors).flatMap(([vendor, models]) =>
        Array.isArray(models)
          ? models.map(model => ({
              name: model,
              vendor
            }))
          : []
      );
    } catch (error) {
      console.error("Failed to fetch models:", error);
      throw error;
    }
  },
};
