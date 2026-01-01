"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.globalifySelector = void 0;
/* eslint-disable line-comment-position */
/*
 * Split a selector string (ex: div > foo ~ .potato) by
 * separators: space, >, +, ~ and comma (maybe not needed)
 * We use a negative lookbehind assertion to prevent matching
 * escaped combinators like `\~`.
 */
// TODO: maybe replace this ugly pattern with an actual selector parser? (https://github.com/leaverou/parsel, 2kb)
const combinatorPattern = /(?<!\\)(?:\\\\)*([ >+~,]\s*)(?![^[]+\]|\d)/g;
function globalifySelector(selector) {
    const parts = selector.trim().split(combinatorPattern);
    const newSelector = [];
    for (let i = 0; i < parts.length; i++) {
        const part = parts[i];
        // if this is a separator or a :global
        if (i % 2 !== 0 || part === '' || part.startsWith(':global')) {
            newSelector.push(part);
            continue;
        }
        // :local() with scope
        if (part.startsWith(':local(')) {
            newSelector.push(part.replace(/:local\((.+?)\)/g, '$1'));
            continue;
        }
        // :local inlined in a selector
        if (part.startsWith(':local')) {
            // + 2 to ignore the :local and space combinator
            const startIndex = i + 2;
            let endIndex = parts.findIndex((p, idx) => idx > startIndex && p.startsWith(':global'));
            endIndex = endIndex === -1 ? parts.length - 1 : endIndex;
            newSelector.push(...parts.slice(startIndex, endIndex + 1));
            i = endIndex;
            continue;
        }
        newSelector.push(`:global(${part})`);
    }
    return newSelector.join('');
}
exports.globalifySelector = globalifySelector;
