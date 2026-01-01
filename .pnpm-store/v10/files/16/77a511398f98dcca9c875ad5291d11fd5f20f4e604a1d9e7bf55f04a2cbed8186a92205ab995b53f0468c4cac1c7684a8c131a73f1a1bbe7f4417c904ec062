/**
 * Create a hast element from a simple CSS selector.
 *
 * @template {string} Selector
 *   Type of selector.
 * @template {string} [DefaultTagName='div']
 *   Type of default tag name (default: `'div'`).
 * @param {Selector | null | undefined} [selector]
 *   Simple CSS selector (optional).
 *
 *   Can contain a tag name (`foo`), classes (`.bar`), and an ID (`#baz`).
 *   Multiple classes are allowed.
 *   Uses the last ID if multiple IDs are found.
 * @param {DefaultTagName | null | undefined} [defaultTagName='div']
 *   Tag name to use if `selector` does not specify one (default: `'div'`).
 * @returns {Element & {tagName: ExtractTagName<Selector, DefaultTagName>}}
 *   Built element.
 */
export function parseSelector<Selector extends string, DefaultTagName extends string = "div">(selector?: Selector | null | undefined, defaultTagName?: DefaultTagName | null | undefined): import("hast").Element & {
    tagName: ExtractTagName<Selector, DefaultTagName>;
};
export type Element = import('hast').Element;
export type Properties = import('hast').Properties;
/**
 * Extract tag name from a simple selector.
 */
export type ExtractTagName<SimpleSelector extends string, DefaultTagName extends string> = SimpleSelector extends "" ? DefaultTagName : SimpleSelector extends `${infer TagName}.${infer Rest}` ? ExtractTagName<TagName, DefaultTagName> : SimpleSelector extends `${infer TagName_1}#${infer Rest_1}` ? ExtractTagName<TagName_1, DefaultTagName> : SimpleSelector extends string ? SimpleSelector : DefaultTagName;
