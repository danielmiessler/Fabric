"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getSvelteIgnoreItems = getSvelteIgnoreItems;
const compat_1 = require("../../utils/compat");
const SVELTE_IGNORE_PATTERN = /^\s*svelte-ignore/m;
/**
 * Map of legacy code -> new code
 * See https://github.com/sveltejs/svelte/blob/c9202a889612df3c2fcb369096a5573668be99d6/packages/svelte/src/compiler/utils/extract_svelte_ignore.js#L6
 */
const V5_REPLACEMENTS = {
    'non-top-level-reactive-declaration': 'reactive_declaration_invalid_placement',
    'module-script-reactive-declaration': 'reactive_declaration_module_script',
    'empty-block': 'block_empty',
    'avoid-is': 'attribute_avoid_is',
    'invalid-html-attribute': 'attribute_invalid_property_name',
    'a11y-structure': 'a11y_figcaption_parent',
    'illegal-attribute-character': 'attribute_illegal_colon',
    'invalid-rest-eachblock-binding': 'bind_invalid_each_rest',
    'unused-export-let': 'export_let_unused'
};
/** Extract all svelte-ignore comment items */
function getSvelteIgnoreItems(context) {
    const sourceCode = (0, compat_1.getSourceCode)(context);
    const ignoreComments = [];
    for (const comment of sourceCode.getAllComments()) {
        const ignores = extractSvelteIgnore(comment.value, comment.range[0] + 2, comment);
        if (ignores) {
            ignoreComments.push(...ignores);
        }
        else if (hasMissingCodeIgnore(comment.value)) {
            ignoreComments.push({
                range: comment.range,
                code: null,
                token: comment
            });
        }
    }
    for (const token of sourceCode.ast.tokens) {
        if (token.type === 'HTMLComment') {
            const text = token.value.slice(4, -3);
            const ignores = extractSvelteIgnore(text, token.range[0] + 4, token);
            if (ignores) {
                ignoreComments.push(...ignores);
            }
            else if (hasMissingCodeIgnore(text)) {
                ignoreComments.push({
                    range: token.range,
                    code: null,
                    token
                });
            }
        }
    }
    ignoreComments.sort((a, b) => b.range[0] - a.range[0]);
    return ignoreComments;
}
/** Extract svelte-ignore rule names */
function extractSvelteIgnore(text, startIndex, token) {
    const m1 = SVELTE_IGNORE_PATTERN.exec(text);
    if (!m1) {
        return null;
    }
    const ignoreStart = m1.index + m1[0].length;
    const beforeText = text.slice(ignoreStart);
    if (!/^\s/.test(beforeText) || !beforeText.trim()) {
        return null;
    }
    let start = startIndex + ignoreStart;
    const results = [];
    for (const code of beforeText.split(/\s/)) {
        const end = start + code.length;
        const trimmed = code.trim();
        if (trimmed) {
            results.push({
                code: trimmed,
                codeForV5: V5_REPLACEMENTS[trimmed] || trimmed.replace(/-/gu, '_'),
                range: [start, end],
                token
            });
        }
        start = end + 1; /* space */
    }
    return results;
}
/** Checks whether given comment has missing code svelte-ignore */
function hasMissingCodeIgnore(text) {
    const m1 = SVELTE_IGNORE_PATTERN.exec(text);
    if (!m1) {
        return false;
    }
    const ignoreStart = m1.index + m1[0].length;
    const beforeText = text.slice(ignoreStart);
    return !beforeText.trim();
}
