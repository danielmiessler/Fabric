"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.indexOf = indexOf;
exports.lastIndexOf = lastIndexOf;
exports.getWithLoc = getWithLoc;
/** indexOf */
function indexOf(str, search, start, end) {
    const endIndex = end !== null && end !== void 0 ? end : str.length;
    for (let index = start; index < endIndex; index++) {
        const c = str[index];
        if (search(c, index)) {
            return index;
        }
    }
    return -1;
}
/** lastIndexOf */
function lastIndexOf(str, search, end) {
    for (let index = end; index >= 0; index--) {
        const c = str[index];
        if (search(c, index)) {
            return index;
        }
    }
    return -1;
}
/** Get node with location */
function getWithLoc(node) {
    return node;
}
