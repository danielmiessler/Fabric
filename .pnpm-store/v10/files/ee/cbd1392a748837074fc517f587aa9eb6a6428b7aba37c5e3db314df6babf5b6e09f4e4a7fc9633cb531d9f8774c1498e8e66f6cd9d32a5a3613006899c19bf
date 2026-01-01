import ToLineItemTransformation from '../ToLineItemTransformation.jsx';
import ParseResult from '../../ParseResult.jsx';
import { DETECTED_ANNOTATION } from '../../Annotation.jsx';
import BlockType from '../../markdown/BlockType.jsx';
import { headlineByLevel } from '../../markdown/BlockType.jsx';
import { isListItem } from '../../../stringFunctions.jsx';

//Detect headlines based on heights
export default class DetectHeaders extends ToLineItemTransformation {

    constructor() {
        super("Detect Headers");
    }

    transform(parseResult:ParseResult) {
        const {tocPages, headlineTypeToHeightRange, mostUsedHeight, mostUsedDistance, mostUsedFont, maxHeight} = parseResult.globals;
        const hasToc = tocPages.length > 0;
        var detectedHeaders = 0;

        // Handle title pages
        const pagesWithMaxHeight = findPagesWithMaxHeight(parseResult.pages, maxHeight);
        const min2ndLevelHeaderHeigthOnMaxPage = mostUsedHeight + ((maxHeight - mostUsedHeight) / 4);
        pagesWithMaxHeight.forEach(titlePage => {
            titlePage.items.forEach(item => {
                const height = item.height;
                if (!item.type && height > min2ndLevelHeaderHeigthOnMaxPage) {
                    if (height == maxHeight) {
                        item.type = BlockType.H1;
                    } else {
                        item.type = BlockType.H2;
                    }
                    item.annotation = DETECTED_ANNOTATION;
                    detectedHeaders++;
                }
            });
        });

        if (hasToc) { //Use existing headline heights to find additional headlines
            const headlineTypes = Object.keys(headlineTypeToHeightRange);
            headlineTypes.forEach(headlineType => {
                var range = headlineTypeToHeightRange[headlineType];
                if (range.max > mostUsedHeight) { //use only very clear headlines, only use max
                    parseResult.pages.forEach(page => {
                        page.items.forEach(item => {
                            if (!item.type && item.height == range.max) {
                                item.annotation = DETECTED_ANNOTATION;
                                item.type = BlockType.enumValueOf(headlineType);
                                detectedHeaders++
                            }
                        });
                    });
                }

            });
        } else { //Categorize headlines by the text heights
            const heights = [];
            var lastHeight;
            parseResult.pages.forEach(page => {
                page.items.forEach(item => {
                    if (!item.type && item.height > mostUsedHeight && !isListItem(item.text())) {
                        if (!heights.includes(item.height) && (!lastHeight || lastHeight > item.height)) {
                            heights.push(item.height);
                        }
                    }
                });
            });
            heights.sort((a, b) => b - a);

            heights.forEach((height, i) => {
                const headlineLevel = i + 2;
                if (headlineLevel <= 6) {
                    const headlineType = headlineByLevel(2 + i);
                    parseResult.pages.forEach(page => {
                        page.items.forEach(item => {
                            if (!item.type && item.height == height && !isListItem(item.text())) {
                                detectedHeaders++;
                                item.annotation = DETECTED_ANNOTATION;
                                item.type = headlineType;
                            }
                        });
                    });
                }
            });
        }

        //find headlines which have paragraph height
        var smallesHeadlineLevel = 1;
        parseResult.pages.forEach(page => {
            page.items.forEach(item => {
                if (item.type && item.type.headline) {
                    smallesHeadlineLevel = Math.max(smallesHeadlineLevel, item.type.headlineLevel);
                }
            });
        });
        if (smallesHeadlineLevel < 6) {
            const nextHeadlineType = headlineByLevel(smallesHeadlineLevel + 1);
            parseResult.pages.forEach(page => {
                var lastItem;
                page.items.forEach(item => {
                    if (!item.type
                            && item.height == mostUsedHeight
                            && item.font !== mostUsedFont
                            && (!lastItem || lastItem.y < item.y || (lastItem.type && lastItem.type.headline) || (lastItem.y - item.y > mostUsedDistance * 2))
                            && item.text() === item.text().toUpperCase()
                    ) {
                        detectedHeaders++;
                        item.annotation = DETECTED_ANNOTATION;
                        item.type = nextHeadlineType;
                    }
                    lastItem = item;
                });
            });
        }


        return new ParseResult({
            ...parseResult,
            messages: [
                'Detected ' + detectedHeaders + ' headlines.',
            ]
        });

    }

}

function findPagesWithMaxHeight(pages, maxHeight) {
    const maxHeaderPagesSet = new Set();
    pages.forEach(page => {
        page.items.forEach(item => {
            if (!item.type && item.height == maxHeight) {
                maxHeaderPagesSet.add(page);
            }
        });
    });
    return maxHeaderPagesSet;
}

