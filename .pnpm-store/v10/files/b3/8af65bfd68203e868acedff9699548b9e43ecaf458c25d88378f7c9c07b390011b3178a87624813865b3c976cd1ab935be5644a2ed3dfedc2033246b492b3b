/**
 * @typedef {import('hast').Nodes} Nodes
 */

import {hasProperty} from 'hast-util-has-property'

// Always interactive nodes.
const alwaysInteractive = new Set([
  'button',
  'details',
  'embed',
  'iframe',
  'keygen',
  'label',
  'select',
  'textarea'
])

/**
 * Check if the given value is *interactive content*.
 *
 * @param {Nodes} node
 *   Node to check.
 * @returns {boolean}
 *   Whether `node` is an `Element` that is classified as *interactive
 *   content*.
 *
 *   The following elements are interactive:
 *
 *   * `a` with `href`
 *   * `audio` or `video` with `controls`
 *   * `img` or `object` with `useMap`
 *   * `input` without `hidden`
 *   * any element with a `tabIndex`
 *   * the elements `button`, `details`, `embed`, `iframe`, `keygen`, `label`,
 *      `select`, and `textarea`
 */
export function interactive(node) {
  if (node.type !== 'element') {
    return false
  }

  const name = node.tagName

  return (
    (name === 'a' && hasProperty(node, 'href')) ||
    (name === 'audio' && hasProperty(node, 'controls')) ||
    (name === 'video' && hasProperty(node, 'controls')) ||
    (name === 'object' && hasProperty(node, 'useMap')) ||
    (name === 'img' && hasProperty(node, 'useMap')) ||
    (name === 'input' && node.properties.type !== 'hidden') ||
    hasProperty(node, 'tabIndex') ||
    alwaysInteractive.has(name)
  )
}
