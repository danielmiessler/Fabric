/**
 * @typedef {import('hast').Element} Element
 * @typedef {import('hast').Nodes} Nodes
 */

const own = {}.hasOwnProperty

/**
 * Check if `node` is an element and has a `name` property.
 *
 * @template {string} Key
 *   Type of key.
 * @param {Nodes} node
 *   Node to check (typically `Element`).
 * @param {Key} name
 *   Property name to check.
 * @returns {node is Element & {properties: Record<Key, Array<number | string> | number | string | true>}}}
 *   Whether `node` is an element that has a `name` property.
 *
 *   Note: see <https://github.com/DefinitelyTyped/DefinitelyTyped/blob/27c9274/types/hast/index.d.ts#L37C29-L37C98>.
 */
export function hasProperty(node, name) {
  const value =
    node.type === 'element' &&
    own.call(node.properties, name) &&
    node.properties[name]

  return value !== null && value !== undefined && value !== false
}
