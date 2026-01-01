import ToLineItemTransformation from '../ToLineItemTransformation.jsx';
import ParseResult from '../../ParseResult.jsx';
import { REMOVED_ANNOTATION } from '../../Annotation.jsx';

import { isDigit } from '../../../stringFunctions.jsx'


function hashCodeIgnoringSpacesAndNumbers(string) {
    var hash = 0;
    if (string.trim().length === 0) return hash;
    for (var i = 0; i < string.length; i++) {
        const charCode = string.charCodeAt(i);
        if (!isDigit(charCode) && charCode != 32 && charCode != 160) {
            hash = ((hash << 5) - hash) + charCode;
            hash |= 0; // Convert to 32bit integer
        }
    }
    return hash;
}


// Remove elements with similar content on same page positions, like page numbers, licenes information, etc...
export default class RemoveRepetitiveElements extends ToLineItemTransformation {

    constructor() {
        super("Remove Repetitive Elements");
    }

    // The idea is the following:
    // - For each page, collect all items of the first, and all items of the last line
    // - Calculate how often these items occur accros all pages (hash ignoring numbers, whitespace, upper/lowercase)
    // - Delete items occuring on more then 2/3 of all pages
    transform(parseResult:ParseResult) {

        // find first and last lines per page
        const pageStore = [];
        const minLineHashRepetitions = {};
        const maxLineHashRepetitions = {};
        parseResult.pages.forEach(page => {
            const minMaxItems = page.items.reduce((itemStore, item) => {
                if (item.y < itemStore.minY) {
                    itemStore.minElements = [item];
                    itemStore.minY = item.y;
                } else if (item.y == itemStore.minY) {
                    itemStore.minElements.push(item);
                }
                if (item.y > itemStore.maxY) {
                    itemStore.maxElements = [item];
                    itemStore.maxY = item.y;
                } else if (item.y == itemStore.maxY) {
                    itemStore.maxElements.push(item);
                }
                return itemStore;
            }, {
                minY: 999,
                maxY: 0,
                minElements: [],
                maxElements: []
            });

            const minLineHash = hashCodeIgnoringSpacesAndNumbers(minMaxItems.minElements.reduce((combinedString, item) => combinedString + item.text().toUpperCase(), ''));
            const maxLineHash = hashCodeIgnoringSpacesAndNumbers(minMaxItems.maxElements.reduce((combinedString, item) => combinedString + item.text().toUpperCase(), ''));
            pageStore.push({
                minElements: minMaxItems.minElements,
                maxElements: minMaxItems.maxElements,
                minLineHash: minLineHash,
                maxLineHash: maxLineHash
            });
            minLineHashRepetitions[minLineHash] = minLineHashRepetitions[minLineHash] ? minLineHashRepetitions[minLineHash] + 1 : 1;
            maxLineHashRepetitions[maxLineHash] = maxLineHashRepetitions[maxLineHash] ? maxLineHashRepetitions[maxLineHash] + 1 : 1;
        });

        // now annoate all removed items
        var removedHeader = 0;
        var removedFooter = 0;
        parseResult.pages.forEach((page, i) => {
            if (minLineHashRepetitions[pageStore[i].minLineHash] >= Math.max(3, parseResult.pages.length * 2 / 3)) {
                pageStore[i].minElements.forEach(item => {
                    item.annotation = REMOVED_ANNOTATION;
                });
                removedFooter++;
            }
            if (maxLineHashRepetitions[pageStore[i].maxLineHash] >= Math.max(3, parseResult.pages.length * 2 / 3)) {
                pageStore[i].maxElements.forEach(item => {
                    item.annotation = REMOVED_ANNOTATION;
                });
                removedHeader++;
            }
        });

        return new ParseResult({
            ...parseResult,
            messages: [
                'Removed Header: ' + removedHeader,
                'Removed Footers: ' + removedFooter
            ]
        });
    }

}