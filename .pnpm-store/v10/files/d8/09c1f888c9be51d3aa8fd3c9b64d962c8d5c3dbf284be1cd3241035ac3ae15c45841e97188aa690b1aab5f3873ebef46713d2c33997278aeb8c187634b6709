import type {Child, Properties, Result} from './create-h.js'

export namespace JSX {
  /**
   * Define the return value of JSX syntax.
   */
  type Element = Result

  /**
   * Key of this interface defines as what prop children are passed.
   */
  interface ElementChildrenAttribute {
    /**
     * Only the key matters, not the value.
     */
    children?: never
  }

  /**
   * Disallow the use of functional components.
   */
  type IntrinsicAttributes = never

  /**
   * Define the prop types for known elements.
   *
   * For `hastscript` this defines any string may be used in combination with
   * `hast` `Properties`.
   *
   * This **must** be an interface.
   */
  type IntrinsicElements = Record<
    string,
    | Properties
    | {
        /**
         * The prop that matches `ElementChildrenAttribute` key defines the
         * type of JSX children, defines the children type.
         */
        children?: Child
      }
  >
}
