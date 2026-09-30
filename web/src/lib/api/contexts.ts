import { api } from './base';
import type { Context } from '$lib/interfaces/context-interface';

export const contextAPI = {
  async getAvailable(): Promise<Context[]> {
    const response = await api.fetch<Context[]>('/contexts/names');
    return response.data || [];
  }
}

// TODO: add a context element to the UI. It could be the file upload
// control or a separate area to upload a context.
