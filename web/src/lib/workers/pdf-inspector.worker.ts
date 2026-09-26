import init, { processPdf, type PdfProcessResult } from '@firecrawl/pdf-inspector-wasm';

// Request, service → worker. The buffer moves through the transfer list.
export interface PdfRequest {
  id: number;
  buffer: ArrayBuffer;
}

// Response, worker → service.
export type PdfResponse =
  { id: number; ok: true; result: PdfProcessResult } | { id: number; ok: false; error: string };

// init() runs one time, at module evaluation. Each message handler awaits
// the same promise. An initialization failure becomes an error response.
const ready = init();

self.onmessage = async (event: MessageEvent<PdfRequest>) => {
  const { id, buffer } = event.data;
  try {
    await ready;
    const result = processPdf(new Uint8Array(buffer), {
      profile: 'compact',
      includePageMarkers: true
    });
    self.postMessage({ id, ok: true, result } satisfies PdfResponse);
  } catch (err) {
    self.postMessage({ id, ok: false, error: String(err) } satisfies PdfResponse);
  }
};
