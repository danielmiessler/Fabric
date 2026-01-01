/**
 * hast utility to get the plain-text value of a node.
 *
 * ## What is this?
 *
 * This package is a utility to get the plain-text value of a node.
 *
 * ## When should I use this?
 *
 * You can use this package when you want to get the plain text value of a node.
 * The algorithm used by this package is like the DOMs `Node#textContent`
 * getter.
 *
 * To use the DOMs `Node#innerText` algorithm instead, use
 * [`hast-util-to-text`](https://github.com/syntax-tree/hast-util-to-text).
 * `innerText` is aware of how things are displayed, for example turning hard
 * breaks (`<br>`) into line endings.
 *
 * ## Use
 *
 * ```js
 * import {h} from 'hastscript'
 * import {toString} from 'hast-util-to-string'
 *
 * toString(h('p', 'Alpha'))
 * //=> 'Alpha'
 * toString(h('div', [h('b', 'Bold'), ' and ', h('i', 'italic'), '.']))
 * //=> 'Bold and italic.'
 * ```
 *
 * ## API
 *
 * ### `toString(node)`
 *
 * Get the plain-text value of a node.
 *
 * ###### Parameters
 *
 * *   `node` (`Node`) — node to serialize
 *
 * ###### Returns
 *
 * Serialized node (`string`).
 */

export {toString} from './lib/index.js'
