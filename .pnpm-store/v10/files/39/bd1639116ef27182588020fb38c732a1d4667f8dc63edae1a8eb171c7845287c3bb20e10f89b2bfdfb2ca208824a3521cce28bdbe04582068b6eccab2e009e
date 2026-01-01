/**
 * @import {Element, Nodes, RootContent, Root} from 'hast'
 * @import {Info, Schema} from 'property-information'
 */

/**
 * @typedef {Array<Nodes | PrimitiveChild>} ArrayChildNested
 *   List of children (deep).
 */

/**
 * @typedef {Array<ArrayChildNested | Nodes | PrimitiveChild>} ArrayChild
 *   List of children.
 */

/**
 * @typedef {Array<number | string>} ArrayValue
 *   List of property values for space- or comma separated values (such as `className`).
 */

/**
 * @typedef {ArrayChild | Nodes | PrimitiveChild} Child
 *   Acceptable child value.
 */

/**
 * @typedef {number | string | null | undefined} PrimitiveChild
 *   Primitive children, either ignored (nullish), or turned into text nodes.
 */

/**
 * @typedef {boolean | number | string | null | undefined} PrimitiveValue
 *   Primitive property value.
 */

/**
 * @typedef {Record<string, PropertyValue | Style>} Properties
 *   Acceptable value for element properties.
 */

/**
 * @typedef {ArrayValue | PrimitiveValue} PropertyValue
 *   Primitive value or list value.
 */

/**
 * @typedef {Element | Root} Result
 *   Result from a `h` (or `s`) call.
 */

/**
 * @typedef {number | string} StyleValue
 *   Value for a CSS style field.
 */

/**
 * @typedef {Record<string, StyleValue>} Style
 *   Supported value of a `style` prop.
 */

import {parse as parseCommas} from 'comma-separated-tokens'
import {parseSelector} from 'hast-util-parse-selector'
import {find, normalize} from 'property-information'
import {parse as parseSpaces} from 'space-separated-tokens'

/**
 * @param {Schema} schema
 *   Schema to use.
 * @param {string} defaultTagName
 *   Default tag name.
 * @param {ReadonlyArray<string> | undefined} [caseSensitive]
 *   Case-sensitive tag names (default: `undefined`).
 * @returns
 *   `h`.
 */
export function createH(schema, defaultTagName, caseSensitive) {
  const adjust = caseSensitive ? createAdjustMap(caseSensitive) : undefined

  /**
   * Hyperscript compatible DSL for creating virtual hast trees.
   *
   * @overload
   * @param {null | undefined} [selector]
   * @param {...Child} children
   * @returns {Root}
   *
   * @overload
   * @param {string} selector
   * @param {Properties} properties
   * @param {...Child} children
   * @returns {Element}
   *
   * @overload
   * @param {string} selector
   * @param {...Child} children
   * @returns {Element}
   *
   * @param {string | null | undefined} [selector]
   *   Selector.
   * @param {Child | Properties | null | undefined} [properties]
   *   Properties (or first child) (default: `undefined`).
   * @param {...Child} children
   *   Children.
   * @returns {Result}
   *   Result.
   */
  function h(selector, properties, ...children) {
    /** @type {Result} */
    let node

    if (selector === null || selector === undefined) {
      node = {type: 'root', children: []}
      // Properties are not supported for roots.
      const child = /** @type {Child} */ (properties)
      children.unshift(child)
    } else {
      node = parseSelector(selector, defaultTagName)
      // Normalize the name.
      const lower = node.tagName.toLowerCase()
      const adjusted = adjust ? adjust.get(lower) : undefined
      node.tagName = adjusted || lower

      // Handle properties.
      if (isChild(properties)) {
        children.unshift(properties)
      } else {
        for (const [key, value] of Object.entries(properties)) {
          addProperty(schema, node.properties, key, value)
        }
      }
    }

    // Handle children.
    for (const child of children) {
      addChild(node.children, child)
    }

    if (node.type === 'element' && node.tagName === 'template') {
      node.content = {type: 'root', children: node.children}
      node.children = []
    }

    return node
  }

  return h
}

/**
 * Check if something is properties or a child.
 *
 * @param {Child | Properties} value
 *   Value to check.
 * @returns {value is Child}
 *   Whether `value` is definitely a child.
 */
function isChild(value) {
  // Never properties if not an object.
  if (value === null || typeof value !== 'object' || Array.isArray(value)) {
    return true
  }

  // Never node without `type`; that’s the main discriminator.
  if (typeof value.type !== 'string') return false

  // Slower check: never property value if object or array with
  // non-number/strings.
  const record = /** @type {Record<string, unknown>} */ (value)
  const keys = Object.keys(value)

  for (const key of keys) {
    const value = record[key]

    if (value && typeof value === 'object') {
      if (!Array.isArray(value)) return true

      const list = /** @type {ReadonlyArray<unknown>} */ (value)

      for (const item of list) {
        if (typeof item !== 'number' && typeof item !== 'string') {
          return true
        }
      }
    }
  }

  // Also see empty `children` as a node.
  if ('children' in value && Array.isArray(value.children)) {
    return true
  }

  // Default to properties, someone can always pass an empty object,
  // put `data: {}` in a node,
  // or wrap it in an array.
  return false
}

/**
 * @param {Schema} schema
 *   Schema.
 * @param {Properties} properties
 *   Properties object.
 * @param {string} key
 *   Property name.
 * @param {PropertyValue | Style} value
 *   Property value.
 * @returns {undefined}
 *   Nothing.
 */
function addProperty(schema, properties, key, value) {
  const info = find(schema, key)
  /** @type {PropertyValue} */
  let result

  // Ignore nullish and NaN values.
  if (value === null || value === undefined) return

  if (typeof value === 'number') {
    // Ignore NaN.
    if (Number.isNaN(value)) return

    result = value
  }
  // Booleans.
  else if (typeof value === 'boolean') {
    result = value
  }
  // Handle list values.
  else if (typeof value === 'string') {
    if (info.spaceSeparated) {
      result = parseSpaces(value)
    } else if (info.commaSeparated) {
      result = parseCommas(value)
    } else if (info.commaOrSpaceSeparated) {
      result = parseSpaces(parseCommas(value).join(' '))
    } else {
      result = parsePrimitive(info, info.property, value)
    }
  } else if (Array.isArray(value)) {
    result = [...value]
  } else {
    result = info.property === 'style' ? style(value) : String(value)
  }

  if (Array.isArray(result)) {
    /** @type {Array<number | string>} */
    const finalResult = []

    for (const item of result) {
      // Assume no booleans in array.
      finalResult.push(
        /** @type {number | string} */ (
          parsePrimitive(info, info.property, item)
        )
      )
    }

    result = finalResult
  }

  // Class names (which can be added both on the `selector` and here).
  if (info.property === 'className' && Array.isArray(properties.className)) {
    // Assume no booleans in `className`.
    result = properties.className.concat(
      /** @type {Array<number | string> | number | string} */ (result)
    )
  }

  properties[info.property] = result
}

/**
 * @param {Array<RootContent>} nodes
 *   Children.
 * @param {Child} value
 *   Child.
 * @returns {undefined}
 *   Nothing.
 */
function addChild(nodes, value) {
  if (value === null || value === undefined) {
    // Empty.
  } else if (typeof value === 'number' || typeof value === 'string') {
    nodes.push({type: 'text', value: String(value)})
  } else if (Array.isArray(value)) {
    for (const child of value) {
      addChild(nodes, child)
    }
  } else if (typeof value === 'object' && 'type' in value) {
    if (value.type === 'root') {
      addChild(nodes, value.children)
    } else {
      nodes.push(value)
    }
  } else {
    throw new Error('Expected node, nodes, or string, got `' + value + '`')
  }
}

/**
 * Parse a single primitives.
 *
 * @param {Info} info
 *   Property information.
 * @param {string} name
 *   Property name.
 * @param {PrimitiveValue} value
 *   Property value.
 * @returns {PrimitiveValue}
 *   Property value.
 */
function parsePrimitive(info, name, value) {
  if (typeof value === 'string') {
    if (info.number && value && !Number.isNaN(Number(value))) {
      return Number(value)
    }

    if (
      (info.boolean || info.overloadedBoolean) &&
      (value === '' || normalize(value) === normalize(name))
    ) {
      return true
    }
  }

  return value
}

/**
 * Serialize a `style` object as a string.
 *
 * @param {Style} styles
 *   Style object.
 * @returns {string}
 *   CSS string.
 */
function style(styles) {
  /** @type {Array<string>} */
  const result = []

  for (const [key, value] of Object.entries(styles)) {
    result.push([key, value].join(': '))
  }

  return result.join('; ')
}

/**
 * Create a map to adjust casing.
 *
 * @param {ReadonlyArray<string>} values
 *   List of properly cased keys.
 * @returns {Map<string, string>}
 *   Map of lowercase keys to uppercase keys.
 */
function createAdjustMap(values) {
  /** @type {Map<string, string>} */
  const result = new Map()

  for (const value of values) {
    result.set(value.toLowerCase(), value)
  }

  return result
}
