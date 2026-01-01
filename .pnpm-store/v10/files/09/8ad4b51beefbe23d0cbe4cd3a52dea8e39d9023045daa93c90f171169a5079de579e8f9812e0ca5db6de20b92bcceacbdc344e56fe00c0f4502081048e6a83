/**
 * Add links from headings back to themselves.
 *
 * ###### Notes
 *
 * This plugin only applies to headings with `id`s.
 * Use `rehype-slug` to generate `id`s for headings that don’t have them.
 *
 * Several behaviors are supported:
 *
 * *   `'prepend'` (default) — inject link before the heading text
 * *   `'append'` — inject link after the heading text
 * *   `'wrap'` — wrap the whole heading text with the link
 * *   `'before'` — insert link before the heading
 * *   `'after'` — insert link after the heading
 *
 * @param {Readonly<Options> | null | undefined} [options]
 *   Configuration (optional).
 * @returns
 *   Transform.
 */
export default function rehypeAutolinkHeadings(options?: Readonly<Options> | null | undefined): (tree: Root) => undefined;
export type Element = import('hast').Element;
export type ElementContent = import('hast').ElementContent;
export type Properties = import('hast').Properties;
export type Root = import('hast').Root;
export type Test = import('hast-util-is-element').Test;
/**
 * Behavior.
 */
export type Behavior = 'after' | 'append' | 'before' | 'prepend' | 'wrap';
/**
 * Generate content.
 */
export type Build = (element: Readonly<Element>) => Array<ElementContent> | ElementContent;
/**
 * Generate properties.
 */
export type BuildProperties = (element: Readonly<Element>) => Properties;
/**
 * Configuration.
 */
export type Options = {
    /**
     * How to create links (default: `'prepend'`).
     */
    behavior?: Behavior | null | undefined;
    /**
     * Content to insert in the link (default: if `'wrap'` then `undefined`,
     * otherwise `<span class="icon icon-link"></span>`);
     * if `behavior` is `'wrap'` and `Build` is passed, its result replaces the
     * existing content, otherwise the content is added after existing content.
     */
    content?: Readonly<ElementContent> | ReadonlyArray<ElementContent> | Build | null | undefined;
    /**
     * Content to wrap the heading and link with, if `behavior` is `'after'` or
     * `'before'` (optional).
     */
    group?: Readonly<ElementContent> | ReadonlyArray<ElementContent> | Build | null | undefined;
    /**
     * Extra properties to set on the heading (optional).
     */
    headingProperties?: Readonly<Properties> | BuildProperties | null | undefined;
    /**
     * Extra properties to set on the link when injecting (default:
     * `{ariaHidden: true, tabIndex: -1}` if `'append'` or `'prepend'`, otherwise
     * `undefined`).
     */
    properties?: Readonly<Properties> | BuildProperties | null | undefined;
    /**
     * Extra test for which headings are linked (optional).
     */
    test?: Test | null | undefined;
};
/**
 * Deep clone.
 *
 * See: <https://github.com/microsoft/TypeScript-DOM-lib-generator/issues/1237#issuecomment-1345515448>
 */
export type Cloneable<T> = T extends Record<any, any> ? { -readonly [k in keyof T]: Cloneable<T[k]>; } : T;
