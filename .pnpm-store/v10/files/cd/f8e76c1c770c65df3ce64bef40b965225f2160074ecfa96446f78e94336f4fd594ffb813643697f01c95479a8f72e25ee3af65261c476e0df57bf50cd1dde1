/**
 * @import {Element, Root} from 'hast'
 * @import {Child, Properties, PropertyValue, Result, Style, createH as CreateH} from './create-h.js'
 */

/**
 * @typedef {Record<string, Child | PropertyValue | Style>} JSXProps
 */

// Make VS code see references to above symbols.
''

/**
 * Create an automatic runtime.
 *
 * @param {ReturnType<CreateH>} f
 *   `h` function.
 * @returns
 *   Automatic JSX runtime.
 */
export function createAutomaticRuntime(f) {
  /**
   * @overload
   * @param {null} type
   * @param {{children?: Child}} properties
   * @param {string | null | undefined} [key]
   * @returns {Root}
   *
   * @overload
   * @param {string} type
   * @param {JSXProps} properties
   * @param {string | null | undefined} [key]
   * @returns {Element}
   *
   * @param {string | null} type
   *   Element name or `null` to get a root.
   * @param {Properties & {children?: Child}} properties
   *   Properties.
   * @returns {Result}
   *   Result.
   */
  function jsx(type, properties) {
    const {children, ...properties_} = properties
    const result =
      // @ts-ignore: `children` is fine: TS has a recursion problem which
      // sometimes generates broken types.
      type === null ? f(null, children) : f(type, properties_, children)
    return result
  }

  return {Fragment: null, jsxDEV: jsx, jsxs: jsx, jsx}
}
