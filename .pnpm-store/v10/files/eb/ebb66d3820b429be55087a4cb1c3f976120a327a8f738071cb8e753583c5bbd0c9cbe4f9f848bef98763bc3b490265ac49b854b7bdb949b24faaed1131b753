import ToLineItemTransformation from '../ToLineItemTransformation.jsx';
import ParseResult from '../../ParseResult.jsx';
import LineItem from '../../LineItem.jsx';
import Word from '../../Word.jsx';
import HeadlineFinder from '../../HeadlineFinder.jsx';
import { REMOVED_ANNOTATION, ADDED_ANNOTATION } from '../../Annotation.jsx';
import BlockType from '../../markdown/BlockType.jsx';
import { headlineByLevel } from '../../markdown/BlockType.jsx';
import { isDigit, isNumber, wordMatch, hasOnly } from '../../../stringFunctions.jsx'

//Detect table of contents pages plus linked headlines
export default class DetectTOC extends ToLineItemTransformation {

    constructor() {
        super("Detect TOC");
    }

    transform(parseResult:ParseResult) {
        const tocPages = [];
        const maxPagesToEvaluate = Math.min(20, parseResult.pages.length);
        const linkLeveler = new LinkLeveler();


        var tocLinks = [];
        var lastTocPage;
        var headlineItem;
        parseResult.pages.slice(0, maxPagesToEvaluate).forEach(page => {
            var lineItemsWithDigits = 0;
            const unknownLines = new Set();
            const pageTocLinks = [];
            var lastWordsWithoutNumber;
            var lastLine;
            //find lines ending with a number per page
            page.items.forEach(line => {
                var words = line.words.filter(word => !hasOnly(word.string, '.'));
                const digits = [];
                while (words.length > 0 && isNumber(words[words.length - 1].string)) {
                    const lastWord = words.pop();
                    digits.unshift(lastWord.string);
                }

                if (digits.length == 0 && words.length > 0) {
                    const lastWord = words[words.length - 1];
                    while (isDigit(lastWord.string.charCodeAt(lastWord.string.length - 1))) {
                        digits.unshift(lastWord.string.charAt(lastWord.string.length - 1))
                        lastWord.string = lastWord.string.substring(0, lastWord.string.length - 1);
                    }
                }
                var endsWithDigit = digits.length > 0;
                if (endsWithDigit) {
                    endsWithDigit = true;
                    if (lastWordsWithoutNumber) { // 2-line item ?
                        words.push(...lastWordsWithoutNumber);
                        lastWordsWithoutNumber = null;
                    }
                    pageTocLinks.push(new TocLink({
                        pageNumber: parseInt(digits.join('')),
                        lineItem: new LineItem({
                            ...line,
                            words: words
                        })
                    }));
                    lineItemsWithDigits++;
                } else {
                    if (!headlineItem) {
                        headlineItem = line;
                    } else {
                        if (lastWordsWithoutNumber) {
                            unknownLines.add(lastLine);
                        }
                        lastWordsWithoutNumber = words;
                        lastLine = line;
                    }
                }
            });

            // page has been processed
            if (lineItemsWithDigits * 100 / page.items.length > 75) {
                tocPages.push(page.index + 1);
                lastTocPage = page;
                linkLeveler.levelPageItems(pageTocLinks);
                tocLinks.push(...pageTocLinks);

                const newBlocks = [];
                page.items.forEach((line) => {
                    if (!unknownLines.has(line)) {
                        line.annotation = REMOVED_ANNOTATION;
                    }
                    newBlocks.push(line);
                    if (line === headlineItem) {
                        newBlocks.push(new LineItem({
                            ...line,
                            type: BlockType.H2,
                            annotation: ADDED_ANNOTATION
                        }));
                    }
                });
                page.items = newBlocks;
            } else {
                headlineItem = null;
            }
        });

        //all  pages have been processed
        var foundHeadlines = tocLinks.length;
        const notFoundHeadlines = [];
        const foundBySize = [];
        const headlineTypeToHeightRange = {}; //H1={min:23, max:25}

        if (tocPages.length > 0) {
            // Add TOC items
            tocLinks.forEach(tocLink => {
                lastTocPage.items.push(new LineItem({
                    words: [new Word({
                        string: ' '.repeat(tocLink.level * 3) + '-'
                    })].concat(tocLink.lineItem.words),
                    type: BlockType.TOC,
                    annotation: ADDED_ANNOTATION
                }));
            });

            // Add linked headers
            const pageMapping = detectPageMappingNumber(parseResult.pages.filter(page => page.index > lastTocPage.index), tocLinks);
            tocLinks.forEach(tocLink => {
                var linkedPage = parseResult.pages[tocLink.pageNumber + pageMapping];
                var foundHealineItems;
                if (linkedPage) {
                    foundHealineItems = findHeadlineItems(linkedPage, tocLink.lineItem.text());
                    if (!foundHealineItems) { // pages are off by 1 ?
                        linkedPage = parseResult.pages[tocLink.pageNumber + pageMapping + 1];
                        if (linkedPage) {
                            foundHealineItems = findHeadlineItems(linkedPage, tocLink.lineItem.text());
                        }
                    }
                }
                if (foundHealineItems) {
                    addHeadlineItems(linkedPage, tocLink, foundHealineItems, headlineTypeToHeightRange)
                } else {
                    notFoundHeadlines.push(tocLink);
                }
            });

            // Try to find linked headers by height
            var fromPage = lastTocPage.index + 2;
            var lastNotFound = [];
            const rollupLastNotFound = (currentPageNumber) => {
                if (lastNotFound.length > 0) {
                    lastNotFound.forEach(notFoundTocLink => {
                        const headlineType = headlineByLevel(notFoundTocLink.level + 2);
                        const heightRange = headlineTypeToHeightRange[headlineType.name];
                        if (heightRange) {
                            const [pageIndex, lineIndex] = findPageAndLineFromHeadline(parseResult.pages, notFoundTocLink, heightRange, fromPage, currentPageNumber);
                            if (lineIndex > -1) {
                                const page = parseResult.pages[pageIndex];
                                page.items[lineIndex].annotation = REMOVED_ANNOTATION;
                                page.items.splice(lineIndex + 1, 0, new LineItem({
                                    ...notFoundTocLink.lineItem,
                                    type: headlineType,
                                    annotation: ADDED_ANNOTATION,
                                }));
                                foundBySize.push(notFoundTocLink);
                            }
                        }
                    });
                    lastNotFound = [];
                }
            }
            if (notFoundHeadlines.length > 0) {
                tocLinks.forEach(tocLink => {
                    if (notFoundHeadlines.includes(tocLink)) {
                        lastNotFound.push(tocLink);
                    } else {
                        rollupLastNotFound(tocLink.pageNumber);
                        fromPage = tocLink.pageNumber;
                    }
                });
                if (lastNotFound.length > 0) {
                    rollupLastNotFound(parseResult.pages.length);
                }
            }
        }



        const messages = [];
        messages.push('Detected ' + tocPages.length + ' table of content pages');
        if (tocPages.length > 0) {
            messages.push('TOC headline heights: ' + JSON.stringify(headlineTypeToHeightRange));
            messages.push('Found TOC headlines: ' + (foundHeadlines - notFoundHeadlines.length + foundBySize.length) + '/' + foundHeadlines);
        }
        if (notFoundHeadlines.length > 0) {
            messages.push('Found TOC headlines (by size): ' + foundBySize.map(tocLink => tocLink.lineItem.text()));
            messages.push('Missing TOC headlines: ' + notFoundHeadlines.filter(fTocLink => !foundBySize.includes(fTocLink)).map(tocLink => tocLink.lineItem.text() + '=>' + tocLink.pageNumber));
        }
        return new ParseResult({
            ...parseResult,
            globals: {
                ...parseResult.globals,
                tocPages: tocPages,
                headlineTypeToHeightRange: headlineTypeToHeightRange
            },
            messages: messages
        });
    }

}

//Find out how the TOC page link actualy translates to the page.index
function detectPageMappingNumber(pages, tocLinks) {
    for ( var tocLink of tocLinks ) {
        const page = findPageWithHeadline(pages, tocLink.lineItem.text());
        if (page) {
            return page.index - tocLink.pageNumber;
        }
    }
    return null;
}

function findPageWithHeadline(pages, headline) {
    for ( var page of pages ) {
        if (findHeadlineItems(page, headline)) {
            return page;
        }
    }
    return null;
}

function findHeadlineItems(page, headline) {
    const headlineFinder = new HeadlineFinder({
        headline: headline
    });
    var lineIndex = 0;
    for ( var line of page.items ) {
        const headlineItems = headlineFinder.consume(line);
        if (headlineItems) {
            return {
                lineIndex: lineIndex,
                headlineItems: headlineItems
            };
        }
        lineIndex++;
    }
    return null;
}

function addHeadlineItems(page, tocLink, foundItems, headlineTypeToHeightRange) {
    foundItems.headlineItems.forEach(item => item.annotation = REMOVED_ANNOTATION);
    const headlineType = headlineByLevel(tocLink.level + 2);
    const headlineHeight = foundItems.headlineItems.reduce((max, item) => Math.max(max, item.height), 0);
    page.items.splice(foundItems.lineIndex + 1, 0, new LineItem({
        ...foundItems.headlineItems[0],
        words: tocLink.lineItem.words,
        height: headlineHeight,
        type: headlineType,
        annotation: ADDED_ANNOTATION
    }));
    var range = headlineTypeToHeightRange[headlineType.name];
    if (range) {
        range.min = Math.min(range.min, headlineHeight);
        range.max = Math.max(range.max, headlineHeight);
    } else {
        range = {
            min: headlineHeight,
            max: headlineHeight
        };
        headlineTypeToHeightRange[headlineType.name] = range;
    }
}

function findPageAndLineFromHeadline(pages, tocLink, heightRange, fromPage, toPage) {
    const linkText = tocLink.lineItem.text().toUpperCase();
    for (var i = fromPage; i <= toPage; i++) {
        const page = pages[i - 1];
        const lineIndex = page.items.findIndex(line => {
            if (!line.type && !line.annotation && line.height >= heightRange.min && line.height <= heightRange.max) {
                const match = wordMatch(linkText, line.text());
                return match >= 0.5;
            }
            return false;
        });
        if (lineIndex > -1) return [i - 1, lineIndex];
    }
    return [-1, -1];
}

class LinkLeveler {
    constructor() {
        this.levelByMethod = null;
        this.uniqueFonts = [];
    }

    levelPageItems(tocLinks:TocLink[]) {
        if (!this.levelByMethod) {
            const uniqueX = this.calculateUniqueX(tocLinks);
            if (uniqueX.length > 1) {
                this.levelByMethod = this.levelByXDiff;
            } else {
                const uniqueFonts = this.calculateUniqueFonts(tocLinks);
                if (uniqueFonts.length > 1) {
                    this.uniqueFonts = uniqueFonts;
                    this.levelByMethod = this.levelByFont;
                } else {
                    this.levelByMethod = this.levelToZero;
                }
            }
        }
        this.levelByMethod(tocLinks);
    }

    levelByXDiff(tocLinks) {
        const uniqueX = this.calculateUniqueX(tocLinks);
        tocLinks.forEach(link => {
            link.level = uniqueX.indexOf(link.lineItem.x);
        });
    }

    levelByFont(tocLinks) {
        tocLinks.forEach(link => {
            link.level = this.uniqueFonts.indexOf(link.lineItem.font);
        });
    }

    levelToZero(tocLinks) {
        tocLinks.forEach(link => {
            link.level = 0;
        });
    }

    calculateUniqueX(tocLinks) {
        var uniqueX = tocLinks.reduce(function(uniquesArray, link) {
            if (uniquesArray.indexOf(link.lineItem.x) < 0) uniquesArray.push(link.lineItem.x);
            return uniquesArray;
        }, []);

        uniqueX.sort((a, b) => {
            return a - b
        });

        return uniqueX;
    }

    calculateUniqueFonts(tocLinks) {
        var uniqueFont = tocLinks.reduce(function(uniquesArray, link) {
            if (uniquesArray.indexOf(link.lineItem.font) < 0) uniquesArray.push(link.lineItem.font);
            return uniquesArray;
        }, []);

        return uniqueFont;
    }

}

class TocLink {
    constructor(options) {
        this.lineItem = options.lineItem;
        this.pageNumber = options.pageNumber;
        this.level = 0;
    }
}
