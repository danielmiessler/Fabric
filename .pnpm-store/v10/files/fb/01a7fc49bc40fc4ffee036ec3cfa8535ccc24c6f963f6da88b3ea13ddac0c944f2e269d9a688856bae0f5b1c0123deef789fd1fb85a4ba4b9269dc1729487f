// A abstract PageItem class, can be TextItem, LineItem or LineItemBlock
export default class PageItem {

    constructor(options) {
        if (this.constructor === PageItem) {
            throw new TypeError("Can not construct abstract class.");
        }
        this.type = options.type;
        this.annotation = options.annotation;
        this.parsedElements = options.parsedElements;
    }

}

export class ParsedElements {

    constructor(options) {
        this.footnoteLinks = options.footnoteLinks || [];
        this.footnotes = options.footnotes || [];
        this.containLinks = options.containLinks;
        this.formattedWords = options.formattedWords;
    }

    add(parsedElements) {
        this.footnoteLinks = this.footnoteLinks.concat(parsedElements.footnoteLinks);
        this.footnotes = this.footnotes.concat(parsedElements.footnotes);
        this.containLinks = this.containLinks || parsedElements.containLinks;
        this.formattedWords += parsedElements.formattedWords;
    }

}