import type { ModelConfig } from '$lib/interfaces/model-interface';
import { api } from './base';

const DEFAULT_CONFIG: Omit<ModelConfig, 'model'> = {
  temperature: 0.7,
  top_p: 0.9,
  frequency: .5,
  presence: 0,
  maxLength: 2000
};

export const configApi = {
  async get(): Promise<ModelConfig> {
    try {
      const response = await api.fetch<ModelConfig>('/config');
      
      if (!response.data) {
        return { ...DEFAULT_CONFIG, model: '' };
      }
      
      return response.data;
    } catch (error) {
      console.error('Failed to fetch config:', error);
      throw error;
    }
  }
};
