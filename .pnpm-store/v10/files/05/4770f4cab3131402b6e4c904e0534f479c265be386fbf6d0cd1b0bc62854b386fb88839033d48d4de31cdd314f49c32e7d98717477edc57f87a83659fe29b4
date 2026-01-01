"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addElementToSortedArray = addElementToSortedArray;
exports.addElementsToSortedArray = addElementsToSortedArray;
exports.sortedLastIndex = sortedLastIndex;
exports.hasTypeInfo = hasTypeInfo;
/**
 * Add element to a sorted array
 */
function addElementToSortedArray(array, element, compare) {
    const index = sortedLastIndex(array, (target) => compare(target, element));
    array.splice(index, 0, element);
}
/**
 * Add element to a sorted array
 */
function addElementsToSortedArray(array, elements, compare) {
    if (!elements.length) {
        return;
    }
    let last = elements[0];
    let index = sortedLastIndex(array, (target) => compare(target, last));
    for (const element of elements) {
        if (compare(last, element) > 0) {
            index = sortedLastIndex(array, (target) => compare(target, element));
        }
        let e = array[index];
        while (e && compare(e, element) <= 0) {
            e = array[++index];
        }
        array.splice(index, 0, element);
        last = element;
    }
}
/**
 * Uses a binary search to determine the highest index at which value should be inserted into array in order to maintain its sort order.
 */
function sortedLastIndex(array, compare) {
    let lower = 0;
    let upper = array.length;
    while (lower < upper) {
        const mid = Math.floor(lower + (upper - lower) / 2);
        const target = compare(array[mid]);
        if (target < 0) {
            lower = mid + 1;
        }
        else if (target > 0) {
            upper = mid;
        }
        else {
            return mid + 1;
        }
    }
    return upper;
}
function hasTypeInfo(element) {
    var _a;
    if ((_a = element.type) === null || _a === void 0 ? void 0 : _a.startsWith("TS")) {
        return true;
    }
    for (const key of Object.keys(element)) {
        if (key === "parent")
            continue;
        const value = element[key];
        if (value == null)
            continue;
        if (typeof value === "object") {
            if (hasTypeInfo(value))
                return true;
        }
        else if (Array.isArray(value)) {
            for (const v of value) {
                if (typeof v === "object" && hasTypeInfo(v))
                    return true;
            }
        }
    }
    return false;
}
