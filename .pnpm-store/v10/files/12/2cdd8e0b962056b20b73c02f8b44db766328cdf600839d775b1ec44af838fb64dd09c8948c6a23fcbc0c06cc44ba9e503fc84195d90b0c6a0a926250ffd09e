/**
 * @import {ElementContent, Element, Properties} from 'hast'
 * @import {LinkReference} from 'mdast'
 * @import {State} from '../state.js'
 */

import {normalizeUri} from 'micromark-util-sanitize-uri'
import {revert} from '../revert.js'

/**
 * Turn an mdast `linkReference` node into hast.
 *
 * @param {State} state
 *   Info passed around.
 * @param {LinkReference} node
 *   mdast node.
 * @returns {Array<ElementContent> | ElementContent}
 *   hast node.
 */
export function linkReference(state, node) {
  const id = String(node.identifier).toUpperCase()
  const definition = state.definitionById.get(id)

  if (!definition) {
    return revert(state, node)
  }

  /** @type {Properties} */
  const properties = {href: normalizeUri(definition.url || '')}

  if (definition.title !== null && definition.title !== undefined) {
    properties.title = definition.title
  }

  /** @type {Element} */
  const result = {
    type: 'element',
    tagName: 'a',
    properties,
    children: state.all(node)
  }
  state.patch(node, result)
  return state.applyData(node, result)
}
