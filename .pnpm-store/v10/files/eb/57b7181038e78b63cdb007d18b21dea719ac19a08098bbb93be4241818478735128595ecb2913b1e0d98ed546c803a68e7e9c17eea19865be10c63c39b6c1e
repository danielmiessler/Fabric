import PdfParser from './PdfParser';
import ProgressListenFunction from './ProgressListenFunction';
import ParseProgressReporter from './ParseProgressReporter';
import ItemTransformer from './transformer/ItemTransformer';
import Debugger from './Debugger';
import { Converter, PdfSource, TransformationResult, transform } from './convert';

export interface ParseContinuation {
  debug(): Debugger;
  transform(): TransformationResult;
}

export default class PdfPipeline {
  parser: PdfParser;
  transformers: ItemTransformer[];

  constructor(parser: PdfParser, transformers: ItemTransformer[]) {
    this.parser = parser;
    this.transformers = transformers;
  }

  async parse(src: PdfSource, progressListener: ProgressListenFunction): Promise<ParseContinuation> {
    return this.parser.parse(src, new ParseProgressReporter(progressListener)).then((parseResult) => {
      return {
        debug: () =>
          new Debugger(
            parseResult.fontMap,
            parseResult.pageViewports,
            parseResult.pageCount,
            parseResult.schema,
            parseResult.items,
            this.transformers,
          ),
        transform: () => {
          const items = transform(parseResult, this.transformers);
          return {
            convert: (converter: Converter) => converter.convert(items),
          };
        },
      };
    });
  }
}
