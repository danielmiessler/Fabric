import { describe, expect, it, vi, beforeEach } from 'vitest';
import { api } from './base';
import { modelsApi } from './models';

describe('modelsApi.getAvailable', () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it('returns one entry for each model of each vendor', async () => {
    vi.spyOn(api, 'fetch').mockResolvedValue({
      data: { models: ['a', 'b'], vendors: { Anthropic: ['a'], OpenAI: ['b'] } }
    });

    const models = await modelsApi.getAvailable();

    expect(models).toEqual([
      { name: 'a', vendor: 'Anthropic' },
      { name: 'b', vendor: 'OpenAI' }
    ]);
  });

  // The server sends null for the model list of a vendor that it can reach no
  // models for, because an empty slice in Go becomes null in JSON. Ollama does
  // this when it is in the configuration but has no models. One such vendor
  // must not stop the models of the other vendors.
  it('skips a vendor whose model list is null', async () => {
    vi.spyOn(api, 'fetch').mockResolvedValue({
      data: {
        models: ['a'],
        vendors: { Anthropic: ['a'], Ollama: null } as unknown as Record<string, string[]>
      }
    });

    const models = await modelsApi.getAvailable();

    expect(models).toEqual([{ name: 'a', vendor: 'Anthropic' }]);
  });

  it('skips a vendor whose model list is not an array', async () => {
    vi.spyOn(api, 'fetch').mockResolvedValue({
      data: {
        models: [],
        vendors: { Broken: 'not-a-list' } as unknown as Record<string, string[]>
      }
    });

    await expect(modelsApi.getAvailable()).resolves.toEqual([]);
  });

  it('reports an error when the response holds no vendors', async () => {
    vi.spyOn(api, 'fetch').mockResolvedValue({ data: { models: [] } as never });

    await expect(modelsApi.getAvailable()).rejects.toThrow('missing vendors data');
  });
});
