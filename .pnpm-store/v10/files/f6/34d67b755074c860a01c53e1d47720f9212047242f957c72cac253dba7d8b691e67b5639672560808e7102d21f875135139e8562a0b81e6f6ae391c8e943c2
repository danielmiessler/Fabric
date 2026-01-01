/**
 * @import {Element, Text} from 'hast'
 * @import {InlineCode} from 'mdast'
 * @import {State} from '../state.js'
 */

/**
 * Turn an mdast `inlineCode` node into hast.
 *
 * @param {State} state
 *   Info passed around.
 * @param {InlineCode} node
 *   mdast node.
 * @returns {Element}
 *   hast node.
 */
export function inlineCode(state, node) {
  /** @type {Text} */
  const text = {type: 'text', value: node.value.replace(/\r?\n|\r/g, ' ')}
  state.patch(node, text)

  /** @type {Element} */
  const result = {
    type: 'element',
    tagName: 'code',
    properties: {},
    children: [text]
  }
  state.patch(node, result)
  return state.applyData(node, result)
}
