import type ParseReporter from './ParseReporter';
import type ProgressListenFunction from './ProgressListenFunction';
import Progress from './Progress';

export default class ParseProgressReporter implements ParseReporter {
  progress = new Progress(['Document Header', 'Metadata', 'Pages', 'Fonts'], [0.01, 0.01, 0.97, 0.01]);
  pagesToParse = 0;
  progressListenFunction: ProgressListenFunction;

  constructor(progressListenFunction: ProgressListenFunction) {
    this.progressListenFunction = progressListenFunction;
  }

  parsedDocumentHeader(numberOfPages: number): void {
    this.pagesToParse = numberOfPages;
    this.progress.stageProgress[0] = 1;
    this.progress.stageDetails[2] = `0 / ${numberOfPages}`;
    this.progressListenFunction(this.progress);
  }

  parsedMetadata(): void {
    this.progress.stageProgress[1] = 1;
    this.progressListenFunction(this.progress);
  }

  parsedPage(index: number): void {
    const pagesParsed = index + 1;
    this.progress.stageProgress[2] = pagesParsed / this.pagesToParse;
    this.progress.stageDetails[2] = `${pagesParsed} / ${this.pagesToParse}`;
    this.progressListenFunction(this.progress);
  }

  parsedFonts(): void {
    this.progress.stageProgress[3] = 1;
    this.progressListenFunction(this.progress);
  }
}
