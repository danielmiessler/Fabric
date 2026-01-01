"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.globalsForRunes = void 0;
exports.getGlobalsForSvelte = getGlobalsForSvelte;
exports.getGlobalsForSvelteScript = getGlobalsForSvelteScript;
const globalsForSvelte = ["$$slots", "$$props", "$$restProps"];
exports.globalsForRunes = [
    "$state",
    "$derived",
    "$effect",
    "$props",
    "$bindable",
    "$inspect",
    "$host",
];
function getGlobalsForSvelte(svelteParseContext) {
    if (svelteParseContext.runes) {
        return [...globalsForSvelte, ...exports.globalsForRunes];
    }
    return globalsForSvelte;
}
function getGlobalsForSvelteScript(svelteParseContext) {
    if (svelteParseContext.runes) {
        return exports.globalsForRunes;
    }
    return [];
}
