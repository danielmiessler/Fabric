import PdfParser from './PdfParser';

export interface ParseConfig {
  // See DocumentInitParameters from https://mozilla.github.io/pdf.js/api/draft/module-pdfjsLib.html#DocumentInitParameters
  pdfjsParams?: object;
  // TODO keep pdfPages ?
}

const defaultConfig: ParseConfig = {
  pdfjsParams: {
    // TODO check if that cmap thing makes sense since we don't bundle them
    cMapUrl: 'cmaps/',
    cMapPacked: true,
  },
};

export function parser(pdfJs: any, options: ParseConfig = defaultConfig) {
  return new PdfParser(pdfJs, options.pdfjsParams);
}
