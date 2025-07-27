import { createPipeline, transformers } from 'pdf-to-markdown-core/lib/src';
import { PARSE_SCHEMA } from 'pdf-to-markdown-core/lib/src/PdfParser';
import * as pdfjs from 'pdfjs-dist';
import pdfConfig from './pdf-config';

export class PdfConversionService {
  constructor() {
    if (typeof window !== 'undefined') {
      console.log('PDF.js version:', pdfjs.version);
      // Initialize PDF.js configuration from the shared config
      pdfConfig.initialize();
      console.log('Worker configuration complete');
    }
  }

  async convertToMarkdown(file: File): Promise<string> {
    console.log('Starting PDF conversion:', {
      fileName: file.name,
      fileSize: file.size
    });

    const buffer = await file.arrayBuffer();
    console.log('Buffer created:', buffer.byteLength);

    const pipeline = createPipeline(pdfjs, {
      transformConfig: { 
        transformers 
      }
    });
    console.log('Pipeline created');

    const result = await pipeline.parse(
      buffer,
      (progress) => console.log('Processing:', {
        stage: progress.stages,
        details: progress.stageDetails,
        progress: progress.stageProgress
      })
    );
    console.log('Parse complete, validating result');

    const transformed = result.transform();
    console.log('Transform applied:', transformed);

    const markdown = transformed.convert({
        convert: (items) => {
          console.log('PDF Structure:', {
            itemCount: items.length,
            firstItem: items[0],
            schema: PARSE_SCHEMA  // ['transform', 'width', 'height', 'str', 'fontName', 'dir']
          });
      
          const text = items
            .map(item => item.value('str'))  // Using 'str' instead of 'text' based on PARSE_SCHEMA
            .filter(Boolean)
            .join('\n');
      
          console.log('Converted text:', {
            length: text.length,
            preview: text.substring(0, 100)
          });
      
          return text;
        }
      });
      

    return markdown;
  }
}

    






