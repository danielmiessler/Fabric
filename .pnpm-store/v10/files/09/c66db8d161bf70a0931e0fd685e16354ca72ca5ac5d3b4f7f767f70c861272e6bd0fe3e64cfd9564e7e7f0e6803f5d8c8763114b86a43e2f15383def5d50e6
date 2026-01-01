/**
 * @import {Element as HastElement, Text as HastText} from 'hast'
 * @import {Text as MdastText} from 'mdast'
 * @import {State} from '../state.js'
 */

import {trimLines} from 'trim-lines'

/**
 * Turn an mdast `text` node into hast.
 *
 * @param {State} state
 *   Info passed around.
 * @param {MdastText} node
 *   mdast node.
 * @returns {HastElement | HastText}
 *   hast node.
 */
export function text(state, node) {
  /** @type {HastText} */
  const result = {type: 'text', value: trimLines(String(node.value))}
  state.patch(node, result)
  return state.applyData(node, result)
}
