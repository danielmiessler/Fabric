import Globals from './Globals';
import Item from './Item';
import ParseProgressReporter from './ParseProgressReporter';
import ParseResult from './ParseResult';
import PdfParser from './PdfParser';
import type ProgressListenFunction from './ProgressListenFunction';
import { assert } from './assert';
import ItemTransformer from './transformer/ItemTransformer';
import TransformContext from './transformer/TransformContext';

export type PdfSource = string | Uint8Array | object;

export interface Converter {
  convert: (items: Item[]) => string;
}

interface Options {
  debug: boolean;
  progressListener: ProgressListenFunction;
}

const defaultOptions: Options = {
  debug: false,
  progressListener: () => {},
};

export interface TransformationResult {
  convert(converter: Converter): string;
}

export async function parseAndTransform(
  src: PdfSource,
  parser: PdfParser,
  transformers: ItemTransformer[],
  progressListener: ProgressListenFunction,
): Promise<TransformationResult> {
  const parseResult = await parseAndVerifyTransformers(src, parser, transformers, progressListener);
  const transformedItems = transform(parseResult, transformers);
  return Promise.resolve({
    convert: (converter: Converter) => converter.convert(transformedItems),
  });
}

export async function convert(
  src: PdfSource,
  parser: PdfParser,
  transformers: ItemTransformer[],
  converter: Converter,
  options: Options = defaultOptions,
): Promise<string> {
  // parse
  const parseResult = await parser.parse(src, new ParseProgressReporter(options.progressListener));
  verifyRequiredColumns(parseResult.schema, transformers);

  // transform
  let items = parseResult.items;
  let globals = new Globals();
  const context = new TransformContext(parseResult.fontMap, parseResult.pageViewports, globals);
  transformers.forEach((transformer) => {
    const result = transformer.transform(context, items);
    globals = globals.withValues(result.globals);
    items = result.items;
  });

  // convert
  return converter.convert(items);
}

async function parseAndVerifyTransformers(
  src: PdfSource,
  parser: PdfParser,
  transformers: ItemTransformer[],
  progressListener: ProgressListenFunction,
): Promise<ParseResult> {
  return parser.parse(src, new ParseProgressReporter(progressListener)).then((parseResult) => {
    verifyRequiredColumns(parseResult.schema, transformers);
    return parseResult;
  });
}

/**
 * Goes through all transformer and makes sure each required column is available in its predecessor schema.
 *
 * @param inputSchema
 * @param transformers
 */
export function verifyRequiredColumns(inputSchema: string[], transformers: ItemTransformer[]) {
  const schemas: string[][] = [inputSchema];
  for (let idx = 0; idx < transformers.length; idx++) {
    const transformer = transformers[idx];
    const predecessorSchema = schemas[idx];
    transformer.descriptor.requireColumns?.forEach((column) => {
      assert(
        predecessorSchema.includes(column),
        `Input schema [${predecessorSchema.join(', ')}] for transformer '${
          transformer.name
        }' does not contain the required column '${column}'`,
      );
    });
    const outputSchema = transformer.schemaTransformer(predecessorSchema);
    schemas.push(outputSchema);
  }
}

export function transform(parseResult: ParseResult, transformers: ItemTransformer[]) {
  let items = parseResult.items;
  let globals = new Globals();
  const context = new TransformContext(parseResult.fontMap, parseResult.pageViewports, globals);
  transformers.forEach((transformer) => {
    const result = transformer.transform(context, items);
    globals = globals.withValues(result.globals);
    items = result.items;
  });
  return items;
}
