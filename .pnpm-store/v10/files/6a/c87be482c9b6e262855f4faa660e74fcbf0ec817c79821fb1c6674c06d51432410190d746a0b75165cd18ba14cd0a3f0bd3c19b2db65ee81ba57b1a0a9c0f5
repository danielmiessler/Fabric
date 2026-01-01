/**
 * @import {Root} from 'hast'
 * @import {Options} from 'hast-util-to-html'
 * @import {Compiler, Processor} from 'unified'
 */

import {toHtml} from 'hast-util-to-html'

/**
 * Plugin to add support for serializing as HTML.
 *
 * @param {Options | null | undefined} [options]
 *   Configuration (optional).
 * @returns {undefined}
 *   Nothing.
 */
export default function rehypeStringify(options) {
  /** @type {Processor<undefined, undefined, undefined, Root, string>} */
  // @ts-expect-error: TS in JSDoc generates wrong types if `this` is typed regularly.
  const self = this
  const settings = {...self.data('settings'), ...options}

  self.compiler = compiler

  /**
   * @type {Compiler<Root, string>}
   */
  function compiler(tree) {
    return toHtml(tree, settings)
  }
}
