import type {Child, Properties, Result} from './create-h.js'

/**
 * This unique symbol is declared to specify the key on which JSX children are
 * passed, without conflicting with the `Attributes` type.
 */
declare const children: unique symbol

/**
 * Define the return value of JSX syntax.
 */
export type Element = Result

/**
 * Key of this interface defines as what prop children are passed.
 */
export interface ElementChildrenAttribute {
  /**
   * Only the key matters, not the value.
   */
  [children]?: never
}

/**
 * Disallow the use of functional components.
 */
export type IntrinsicAttributes = never

/**
 * Define the prop types for known elements.
 *
 * For `hastscript` this defines any string may be used in combination with
 * `hast` `Properties`.
 *
 * This **must** be an interface.
 */
export type IntrinsicElements = Record<
  string,
  | Properties
  | {
      /**
       * The prop that matches `ElementChildrenAttribute` key defines the
       * type of JSX children, defines the children type.
       */
      [children]?: Child
    }
>
