/**
 * @import {Element} from 'hast'
 * @import {Blockquote} from 'mdast'
 * @import {State} from '../state.js'
 */

/**
 * Turn an mdast `blockquote` node into hast.
 *
 * @param {State} state
 *   Info passed around.
 * @param {Blockquote} node
 *   mdast node.
 * @returns {Element}
 *   hast node.
 */
export function blockquote(state, node) {
  /** @type {Element} */
  const result = {
    type: 'element',
    tagName: 'blockquote',
    properties: {},
    children: state.wrap(state.all(node), true)
  }
  state.patch(node, result)
  return state.applyData(node, result)
}
