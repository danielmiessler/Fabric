import { api } from './base';
import type { Context } from '$lib/interfaces/context-interface';

export const contextAPI = {
  async getAvailable(): Promise<Context[]> {
    const response = await api.fetch<Context[]>('/contexts/names');
    return response.data || [];
  }
}

// TODO: add context element somewhere in the UI
// Should the file upload functionality be used as the context element? 
// Or should there be another area to upload a context file?
