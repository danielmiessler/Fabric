import type { PdfProcessResult } from '@firecrawl/pdf-inspector-wasm';
import { afterEach, describe, expect, it, vi } from 'vitest';
import { interpretPdfResult, PdfConversionService } from './PdfConversionService';

// Plain result objects replace the worker and the WASM module in these tests.
// The shape copies the PdfProcessResult fields that the spike measured.
function pdfResult(overrides: Partial<PdfProcessResult> = {}): PdfProcessResult {
  return {
    pdfType: 'TextBased',
    markdown: '# Title\n\nBody text.\n',
    pageCount: 1,
    processingTimeMs: 5,
    pagesNeedingOcr: [],
    ocrReasonsByPage: [],
    confidence: 0.5,
    layout: { isComplex: false, pagesWithTables: [], pagesWithColumns: [] },
    hasEncodingIssues: false,
    ...overrides
  };
}

describe('interpretPdfResult', () => {
  it('returns the Markdown with no warning for a TextBased result', () => {
    const conversion = interpretPdfResult(pdfResult(), 'report.pdf');

    expect(conversion).toEqual({ markdown: '# Title\n\nBody text.\n' });
  });

  it('throws the OCR message for a Scanned result', () => {
    const scanned = pdfResult({ pdfType: 'Scanned', markdown: undefined });

    expect(() => interpretPdfResult(scanned, 'report.pdf')).toThrowError(
      'report.pdf contains no machine-readable text. OCR is necessary, and the Fabric web interface does not include OCR.'
    );
  });

  it('throws the OCR message for an ImageBased result', () => {
    const imageBased = pdfResult({ pdfType: 'ImageBased', markdown: undefined });

    expect(() => interpretPdfResult(imageBased, 'report.pdf')).toThrowError(
      'report.pdf contains no machine-readable text. OCR is necessary, and the Fabric web interface does not include OCR.'
    );
  });

  it('keeps the Markdown and warns with the 1-indexed page list for a Mixed result', () => {
    const mixed = pdfResult({ pdfType: 'Mixed', pagesNeedingOcr: [2, 5] });

    const conversion = interpretPdfResult(mixed, 'report.pdf');

    expect(conversion.markdown).toBe('# Title\n\nBody text.\n');
    expect(conversion.warning).toBe(
      'report.pdf: pages 2, 5 contain no machine-readable text. OCR is necessary for those pages, and their content is not included.'
    );
  });

  it('adds the decode warning when hasEncodingIssues is true', () => {
    const conversion = interpretPdfResult(pdfResult({ hasEncodingIssues: true }), 'report.pdf');

    expect(conversion.markdown).toBe('# Title\n\nBody text.\n');
    expect(conversion.warning).toBe('Some text in report.pdf did not decode correctly.');
  });

  it('joins the Mixed warning and the decode warning with one space', () => {
    const mixed = pdfResult({
      pdfType: 'Mixed',
      pagesNeedingOcr: [3],
      hasEncodingIssues: true
    });

    const conversion = interpretPdfResult(mixed, 'report.pdf');

    expect(conversion.warning).toBe(
      'report.pdf: pages 3 contain no machine-readable text. OCR is necessary for those pages, and their content is not included. ' +
        'Some text in report.pdf did not decode correctly.'
    );
  });

  it('throws when the markdown is absent', () => {
    const noMarkdown = pdfResult({ markdown: undefined });

    expect(() => interpretPdfResult(noMarkdown, 'report.pdf')).toThrowError(
      'report.pdf: the conversion returned no text.'
    );
  });

  it('throws when the markdown is blank', () => {
    const blank = pdfResult({ markdown: '   \n' });

    expect(() => interpretPdfResult(blank, 'report.pdf')).toThrowError(
      'report.pdf: the conversion returned no text.'
    );
  });
});

// The fake worker records each posted request and lets a test send a response.
// This specifies the message contract without the WASM module.
class FakeWorker {
  static instances: FakeWorker[] = [];
  onmessage: ((event: { data: unknown }) => void) | null = null;
  onerror: ((event: { message?: string }) => void) | null = null;
  terminated = false;
  posted: Array<{ data: { id: number; buffer: ArrayBuffer }; transfer: Transferable[] }> = [];

  constructor() {
    FakeWorker.instances.push(this);
  }

  postMessage(data: { id: number; buffer: ArrayBuffer }, transfer: Transferable[]) {
    this.posted.push({ data, transfer });
  }

  terminate() {
    this.terminated = true;
  }
}

describe('PdfConversionService worker messages', () => {
  afterEach(() => {
    vi.unstubAllGlobals();
    FakeWorker.instances.length = 0;
  });

  function setup() {
    vi.stubGlobal('Worker', FakeWorker);
    const service = new PdfConversionService();
    const file = new File([new Uint8Array([0x25, 0x50, 0x44, 0x46])], 'report.pdf', {
      type: 'application/pdf'
    });
    return { service, file };
  }

  it('posts the buffer in the transfer list and resolves on a success response', async () => {
    const { service, file } = setup();

    const promise = service.convertToMarkdown(file);
    await vi.waitFor(() => expect(FakeWorker.instances[0]?.posted).toHaveLength(1));

    const worker = FakeWorker.instances[0];
    const { data, transfer } = worker.posted[0];
    expect(transfer).toEqual([data.buffer]);

    worker.onmessage?.({ data: { id: data.id, ok: true, result: pdfResult() } });

    await expect(promise).resolves.toEqual({ markdown: '# Title\n\nBody text.\n' });
  });

  it('rejects with the worker error text on a parser error response', async () => {
    const { service, file } = setup();

    const promise = service.convertToMarkdown(file);
    await vi.waitFor(() => expect(FakeWorker.instances[0]?.posted).toHaveLength(1));

    const worker = FakeWorker.instances[0];
    const { data } = worker.posted[0];
    worker.onmessage?.({
      data: {
        id: data.id,
        ok: false,
        error: 'Error: process PDF: Not a PDF: file appears to be plain text'
      }
    });

    await expect(promise).rejects.toThrowError('Not a PDF: file appears to be plain text');
  });

  it('rejects all pending requests and terminates the worker on a worker error', async () => {
    const { service, file } = setup();

    const first = service.convertToMarkdown(file);
    const second = service.convertToMarkdown(file);
    await vi.waitFor(() => expect(FakeWorker.instances[0]?.posted).toHaveLength(2));

    const worker = FakeWorker.instances[0];
    worker.onerror?.({ message: 'worker script did not load' });

    await expect(first).rejects.toThrowError('The PDF worker failed.');
    await expect(second).rejects.toThrowError('The PDF worker failed.');
    expect(worker.terminated).toBe(true);
  });

  it('creates a new worker for the request that follows a worker error', async () => {
    const { service, file } = setup();

    const failed = service.convertToMarkdown(file);
    await vi.waitFor(() => expect(FakeWorker.instances[0]?.posted).toHaveLength(1));
    FakeWorker.instances[0].onerror?.({ message: 'worker script did not load' });
    await expect(failed).rejects.toThrowError();

    const retry = service.convertToMarkdown(file);
    await vi.waitFor(() => expect(FakeWorker.instances).toHaveLength(2));

    const worker = FakeWorker.instances[1];
    await vi.waitFor(() => expect(worker.posted).toHaveLength(1));
    const { data } = worker.posted[0];
    worker.onmessage?.({ data: { id: data.id, ok: true, result: pdfResult() } });

    await expect(retry).resolves.toEqual({ markdown: '# Title\n\nBody text.\n' });
  });
});
