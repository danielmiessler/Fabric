import { api } from './base';
import type { VendorModel, ModelsResponse } from '$lib/interfaces/model-interface';

export const modelsApi = {
  async getAvailable(): Promise<VendorModel[]> {
    try {
      const response = await api.fetch<ModelsResponse>('/models/names');
      
      if (!response.data?.vendors) {
        throw new Error('Invalid response format: missing vendors data');
      }
      
      // The server sends null for the model list of a vendor with no models.
      // A nil slice in Go becomes null in JSON. Ollama does this when it is in
      // the configuration but serves no models. Skip such a vendor so that it
      // does not hide the models of the other vendors.
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
