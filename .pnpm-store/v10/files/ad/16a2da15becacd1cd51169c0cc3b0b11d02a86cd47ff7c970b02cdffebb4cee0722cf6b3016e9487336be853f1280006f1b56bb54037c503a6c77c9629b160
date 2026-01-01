import TransformConfig from './Config';
import type ProgressListenFunction from './ProgressListenFunction';
import ParseProgressReporter from './ParseProgressReporter';
import PdfParser from './PdfParser';
import PdfPipeline from './PdfPipeline';

import AdjustHeight from './transformer/AdjustHeight';
import UnwrapCoordinates from './transformer/UnwrapCoordinates';
import RemoveEmptyItems from './transformer/RemoveEmptyItems';
import CalculateStatistics from './transformer/CacluclateStatistics';
import CompactLines from './transformer/CompactLines';
import SortXWithinLines from './transformer/SortXWithinLines';
import RemoveRepetitiveItems from './transformer/RemoveRepetitiveItems';
import DetectToc from './transformer/DetectToc';
import DetectHeaders from './transformer/DetectHeaders';
import NoOpTransformer from './transformer/NoOpTransformer';
import {type ParseConfig } from './parse';
import DetectListItems from './transformer/DetectListItems';
import DetectBlocks from './transformer/DetectBlocks';
import DetectListLevels from './transformer/DetectListLevels';
import DetectFootnotes from './transformer/DetectFootnotes';
import DetectFontStyles from './transformer/DetectFontStyles';
import DetectLinks from './transformer/DetectLinks';
import DetectCodeQuoteBlocks from './transformer/DetectCodeQuoteBlocks';

export const transformers = [
  new AdjustHeight(),
  new UnwrapCoordinates(),
  new RemoveEmptyItems(),
  new CalculateStatistics(),
  new CompactLines(),
  new SortXWithinLines(),
  new RemoveRepetitiveItems(),
  new DetectFootnotes(),
  new DetectFontStyles(),
  new DetectLinks(),
  new DetectToc(),
  new DetectHeaders(),
  new DetectListItems(),
  new DetectBlocks(),
  new DetectCodeQuoteBlocks(),
  new DetectListLevels(),
  new NoOpTransformer(),
];

export interface Options {
  parseConfig?: ParseConfig;
  transformConfig?: TransformConfig;
}

export function parseReporter(progressListener: ProgressListenFunction) {
  return new ParseProgressReporter(progressListener);
}

export function createPipeline(pdfJs: any, options: Options = {}): PdfPipeline {
  const parser = new PdfParser(pdfJs);
  return new PdfPipeline(parser, options.transformConfig?.transformers || transformers);
}
