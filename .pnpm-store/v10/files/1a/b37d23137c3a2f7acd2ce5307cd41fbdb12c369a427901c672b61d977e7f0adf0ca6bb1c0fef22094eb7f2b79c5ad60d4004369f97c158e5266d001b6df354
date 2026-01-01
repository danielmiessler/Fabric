/**
 * @typedef {import('hast').Element} Element
 * @typedef {import('hast').Properties} Properties
 */

/**
 * @template {string} SimpleSelector
 *   Selector type.
 * @template {string} DefaultTagName
 *   Default tag name.
 * @typedef {(
 *   SimpleSelector extends ''
 *     ? DefaultTagName
 *     : SimpleSelector extends `${infer TagName}.${infer Rest}`
 *     ? ExtractTagName<TagName, DefaultTagName>
 *     : SimpleSelector extends `${infer TagName}#${infer Rest}`
 *     ? ExtractTagName<TagName, DefaultTagName>
 *     : SimpleSelector extends string
 *     ? SimpleSelector
 *     : DefaultTagName
 * )} ExtractTagName
 *   Extract tag name from a simple selector.
 */

const search = /[#.]/g

/**
 * Create a hast element from a simple CSS selector.
 *
 * @template {string} Selector
 *   Type of selector.
 * @template {string} [DefaultTagName='div']
 *   Type of default tag name (default: `'div'`).
 * @param {Selector | null | undefined} [selector]
 *   Simple CSS selector (optional).
 *
 *   Can contain a tag name (`foo`), classes (`.bar`), and an ID (`#baz`).
 *   Multiple classes are allowed.
 *   Uses the last ID if multiple IDs are found.
 * @param {DefaultTagName | null | undefined} [defaultTagName='div']
 *   Tag name to use if `selector` does not specify one (default: `'div'`).
 * @returns {Element & {tagName: ExtractTagName<Selector, DefaultTagName>}}
 *   Built element.
 */
export function parseSelector(selector, defaultTagName) {
  const value = selector || ''
  /** @type {Properties} */
  const props = {}
  let start = 0
  /** @type {string | undefined} */
  let previous
  /** @type {string | undefined} */
  let tagName

  while (start < value.length) {
    search.lastIndex = start
    const match = search.exec(value)
    const subvalue = value.slice(start, match ? match.index : value.length)

    if (subvalue) {
      if (!previous) {
        tagName = subvalue
      } else if (previous === '#') {
        props.id = subvalue
      } else if (Array.isArray(props.className)) {
        props.className.push(subvalue)
      } else {
        props.className = [subvalue]
      }

      start += subvalue.length
    }

    if (match) {
      previous = match[0]
      start++
    }
  }

  return {
    type: 'element',
    // @ts-expect-error: tag name is parsed.
    tagName: tagName || defaultTagName || 'div',
    properties: props,
    children: []
  }
}
