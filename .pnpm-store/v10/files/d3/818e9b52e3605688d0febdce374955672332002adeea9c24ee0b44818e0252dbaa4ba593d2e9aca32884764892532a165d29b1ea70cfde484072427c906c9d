import ParseResult from '../ParseResult.jsx';

// A transformation from an PdfPage to an PdfPage
export default class Transformation {

    constructor(name, itemType) {
        if (this.constructor === Transformation) {
            throw new TypeError("Can not construct abstract class.");
        }
        if (this.transform === Transformation.prototype.transform) {
            throw new TypeError("Please implement abstract method 'transform()'.");
        }
        this.name = name;
        this.itemType = itemType;
    }

    showModificationCheckbox() {
        return false;
    }

    createPageView(page, modificationsOnly) { // eslint-disable-line no-unused-vars
        throw new TypeError("Do not call abstract method foo from child.");
    }

    // Transform an incoming ParseResult into an outgoing ParseResult
    transform(parseResult: ParseResult) { // eslint-disable-line no-unused-vars
        throw new TypeError("Do not call abstract method foo from child.");
    }

    // Sometimes the transform() does only visualize a change. This methods then does the actual change.
    completeTransform(parseResult: ParseResult) { // eslint-disable-line no-unused-vars
        parseResult.messages = [];
        return parseResult;
    }


}