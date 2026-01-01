"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.transformMarkup = exports.parseAttributes = exports.stripTags = exports.createTagRegex = void 0;
/** Create a tag matching regexp. */
function createTagRegex(tagName, flags) {
    return new RegExp(`/<!--[^]*?-->|<${tagName}(\\s[^]*?)?(?:>([^]*?)<\\/${tagName}>|\\/>)`, flags);
}
exports.createTagRegex = createTagRegex;
/** Strip script and style tags from markup. */
function stripTags(markup) {
    return markup
        .replace(createTagRegex('style', 'gi'), '')
        .replace(createTagRegex('script', 'gi'), '');
}
exports.stripTags = stripTags;
/** Transform an attribute string into a key-value object */
function parseAttributes(attributesStr) {
    return attributesStr
        .split(/\s+/)
        .filter(Boolean)
        .reduce((acc, attr) => {
        const [name, value] = attr.split('=');
        // istanbul ignore next
        acc[name] = value ? value.replace(/['"]/g, '') : true;
        return acc;
    }, {});
}
exports.parseAttributes = parseAttributes;
async function transformMarkup({ content, filename }, transformer, options = {}) {
    let { markupTagName = 'template' } = options;
    markupTagName = markupTagName.toLocaleLowerCase();
    const markupPattern = createTagRegex(markupTagName);
    const templateMatch = content.match(markupPattern);
    /** If no <template> was found, run the transformer over the whole thing */
    if (!templateMatch || templateMatch.index == null) {
        return transformer({
            content,
            markup: content,
            attributes: {},
            filename,
            options,
        });
    }
    const [fullMatch, attributesStr = '', templateCode] = templateMatch;
    const attributes = parseAttributes(attributesStr);
    /** Transform the found template code */
    let { code, map, dependencies } = await transformer({
        content: templateCode,
        markup: templateCode,
        attributes,
        filename,
        options,
    });
    code =
        content.slice(0, templateMatch.index) +
            code +
            content.slice(templateMatch.index + fullMatch.length);
    return { code, map, dependencies };
}
exports.transformMarkup = transformMarkup;
