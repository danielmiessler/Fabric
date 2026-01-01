/**
 * Automatically add `rel` (and `target`?) to external links.
 *
 * ###### Notes
 *
 * You should [likely not configure `target`][css-tricks].
 *
 * You should at least set `rel` to `['nofollow']`.
 * When using a `target`, add `noopener` and `noreferrer` to avoid exploitation
 * of the `window.opener` API.
 *
 * When using a `target`, you should set `content` to adhere to accessibility
 * guidelines by giving users advanced warning when opening a new window.
 *
 * [css-tricks]: https://css-tricks.com/use-target_blank/
 *
 * @param {Readonly<Options> | null | undefined} [options]
 *   Configuration (optional).
 * @returns
 *   Transform.
 */
export default function rehypeExternalLinks(options?: Readonly<Options> | null | undefined): (tree: Root) => undefined;
export type Element = import('hast').Element;
export type ElementContent = import('hast').ElementContent;
export type Properties = import('hast').Properties;
export type Root = import('hast').Root;
export type Test = import('hast-util-is-element').Test;
/**
 * Create a target for the element.
 */
export type CreateContent = (element: Element) => Array<ElementContent> | ElementContent | null | undefined;
/**
 * Create properties for an element.
 */
export type CreateProperties = (element: Element) => Properties | null | undefined;
/**
 * Create a `rel` for the element.
 */
export type CreateRel = (element: Element) => Array<string> | string | null | undefined;
/**
 * Create a `target` for the element.
 */
export type CreateTarget = (element: Element) => Target | null | undefined;
/**
 * Configuration.
 */
export type Options = {
    /**
     * Content to insert at the end of external links (optional); will be
     * inserted in a `<span>` element; useful for improving accessibility by
     * giving users advanced warning when opening a new window.
     */
    content?: Array<ElementContent> | CreateContent | ElementContent | null | undefined;
    /**
     * Properties to add to the `span` wrapping `content` (optional).
     */
    contentProperties?: CreateProperties | Properties | null | undefined;
    /**
     * Properties to add to the link itself (optional).
     */
    properties?: CreateProperties | Properties | null | undefined;
    /**
     * Protocols to check, such as `mailto` or `tel` (default: `['http',
     * 'https']`).
     */
    protocols?: Array<string> | null | undefined;
    /**
     * Link types to hint about the referenced documents (default:
     * `['nofollow']`); pass an empty array (`[]`) to not set `rel`s on links;
     * when using a `target`, add `noopener` and `noreferrer` to avoid
     * exploitation of the `window.opener` API.
     */
    rel?: Array<string> | CreateRel | string | null | undefined;
    /**
     * How to display referenced documents; the default (nothing) is to not set
     * `target`s on links.
     */
    target?: CreateTarget | Target | null | undefined;
    /**
     * Extra test to define which external link elements are modified (optional);
     * any test that can be given to `hast-util-is-element` is supported.
     */
    test?: Test | null | undefined;
};
/**
 * Target.
 */
export type Target = '_blank' | '_parent' | '_self' | '_top';
