import type { PdfProcessResult } from '@firecrawl/pdf-inspector-wasm';
import type { PdfRequest, PdfResponse } from '../workers/pdf-inspector.worker';

export interface PdfConversion {
	markdown: string;
	warning?: string;
}

// Maps the raw worker result to UI behavior. Throws when there is no usable text.
export function interpretPdfResult(result: PdfProcessResult, fileName: string): PdfConversion {
	if (result.pdfType === 'Scanned' || result.pdfType === 'ImageBased') {
		throw new Error(
			`${fileName} contains no machine-readable text. OCR is necessary, and the Fabric web interface does not include OCR.`
		);
	}
	const markdown = result.markdown;
	if (!markdown || markdown.trim().length === 0) {
		throw new Error(`${fileName}: the conversion returned no text.`);
	}
	const warnings: string[] = [];
	if (result.pdfType === 'Mixed') {
		warnings.push(
			`${fileName}: pages ${result.pagesNeedingOcr.join(', ')} contain no machine-readable text. OCR is necessary for those pages, and their content is not included.`
		);
	}
	if (result.hasEncodingIssues) {
		warnings.push(`Some text in ${fileName} did not decode correctly.`);
	}
	return warnings.length > 0 ? { markdown, warning: warnings.join(' ') } : { markdown };
}

export class PdfConversionService {
	private worker: Worker | null = null;
	private nextId = 1;
	private pending = new Map<
		number,
		{ resolve: (result: PdfProcessResult) => void; reject: (error: Error) => void }
	>();

	async convertToMarkdown(file: File): Promise<PdfConversion> {
		const buffer = await file.arrayBuffer();
		const worker = this.getWorker();
		const id = this.nextId++;
		const result = await new Promise<PdfProcessResult>((resolve, reject) => {
			this.pending.set(id, { resolve, reject });
			worker.postMessage({ id, buffer } satisfies PdfRequest, [buffer]);
		});
		return interpretPdfResult(result, file.name);
	}

	// Lazy creation delays the WASM download until the first PDF attachment.
	private getWorker(): Worker {
		if (this.worker) return this.worker;
		const worker = new Worker(new URL('../workers/pdf-inspector.worker.ts', import.meta.url), {
			type: 'module'
		});
		worker.onmessage = (event: MessageEvent<PdfResponse>) => {
			const response = event.data;
			const entry = this.pending.get(response.id);
			if (!entry) return;
			this.pending.delete(response.id);
			if (response.ok) {
				entry.resolve(response.result);
			} else {
				entry.reject(new Error(response.error));
			}
		};
		// Fires only when the worker script itself does not load. The handler
		// terminates its own worker, so a later replacement worker is safe.
		worker.onerror = () => {
			for (const { reject } of this.pending.values()) {
				reject(new Error('The PDF worker failed.'));
			}
			this.pending.clear();
			worker.terminate();
			this.worker = null;
		};
		this.worker = worker;
		return worker;
	}
}
