"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const Progress_1 = require("./Progress");
class ParseProgressReporter {
    constructor(progressListenFunction) {
        this.progress = new Progress_1.default(['Document Header', 'Metadata', 'Pages', 'Fonts'], [0.01, 0.01, 0.97, 0.01]);
        this.pagesToParse = 0;
        this.progressListenFunction = progressListenFunction;
    }
    parsedDocumentHeader(numberOfPages) {
        this.pagesToParse = numberOfPages;
        this.progress.stageProgress[0] = 1;
        this.progress.stageDetails[2] = `0 / ${numberOfPages}`;
        this.progressListenFunction(this.progress);
    }
    parsedMetadata() {
        this.progress.stageProgress[1] = 1;
        this.progressListenFunction(this.progress);
    }
    parsedPage(index) {
        const pagesParsed = index + 1;
        this.progress.stageProgress[2] = pagesParsed / this.pagesToParse;
        this.progress.stageDetails[2] = `${pagesParsed} / ${this.pagesToParse}`;
        this.progressListenFunction(this.progress);
    }
    parsedFonts() {
        this.progress.stageProgress[3] = 1;
        this.progressListenFunction(this.progress);
    }
}
exports.default = ParseProgressReporter;
