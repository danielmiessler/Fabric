import { api } from './base';
import type { VendorModel, ModelsResponse } from '$lib/interfaces/model-interface';

export const modelsApi = {
  async getAvailable(): Promise<VendorModel[]> {
    try {
      const response = await api.fetch<ModelsResponse>('/models/names');
      
      if (!response.data?.vendors) {
        throw new Error('Invalid response format: missing vendors data');
      }
      
      return Object.entries(response.data.vendors).flatMap(([vendor, models]) =>
        models.map(model => ({
          name: model,
          vendor
        }))
      );
    } catch (error) {
      console.error("Failed to fetch models:", error);
      throw error;
    }
  },
};
