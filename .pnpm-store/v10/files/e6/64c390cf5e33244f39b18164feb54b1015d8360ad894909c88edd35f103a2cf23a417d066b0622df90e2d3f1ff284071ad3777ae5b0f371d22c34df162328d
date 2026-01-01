import type ParseReporter from './ParseReporter';
import type ProgressListenFunction from './ProgressListenFunction';
import Progress from './Progress';
export default class ParseProgressReporter implements ParseReporter {
    progress: Progress;
    pagesToParse: number;
    progressListenFunction: ProgressListenFunction;
    constructor(progressListenFunction: ProgressListenFunction);
    parsedDocumentHeader(numberOfPages: number): void;
    parsedMetadata(): void;
    parsedPage(index: number): void;
    parsedFonts(): void;
}
