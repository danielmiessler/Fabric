/**
 * @typedef {import('hast').Element} Element
 * @typedef {import('hast').Parents} Parents
 */

/**
 * @template Fn
 * @template Fallback
 * @typedef {Fn extends (value: any) => value is infer Thing ? Thing : Fallback} Predicate
 */

/**
 * @callback Check
 *   Check that an arbitrary value is an element.
 * @param {unknown} this
 *   Context object (`this`) to call `test` with
 * @param {unknown} [element]
 *   Anything (typically a node).
 * @param {number | null | undefined} [index]
 *   Position of `element` in its parent.
 * @param {Parents | null | undefined} [parent]
 *   Parent of `element`.
 * @returns {boolean}
 *   Whether this is an element and passes a test.
 *
 * @typedef {Array<TestFunction | string> | TestFunction | string | null | undefined} Test
 *   Check for an arbitrary element.
 *
 *   * when `string`, checks that the element has that tag name
 *   * when `function`, see `TestFunction`
 *   * when `Array`, checks if one of the subtests pass
 *
 * @callback TestFunction
 *   Check if an element passes a test.
 * @param {unknown} this
 *   The given context.
 * @param {Element} element
 *   An element.
 * @param {number | undefined} [index]
 *   Position of `element` in its parent.
 * @param {Parents | undefined} [parent]
 *   Parent of `element`.
 * @returns {boolean | undefined | void}
 *   Whether this element passes the test.
 *
 *   Note: `void` is included until TS sees no return as `undefined`.
 */

/**
 * Check if `element` is an `Element` and whether it passes the given test.
 *
 * @param element
 *   Thing to check, typically `element`.
 * @param test
 *   Check for a specific element.
 * @param index
 *   Position of `element` in its parent.
 * @param parent
 *   Parent of `element`.
 * @param context
 *   Context object (`this`) to call `test` with.
 * @returns
 *   Whether `element` is an `Element` and passes a test.
 * @throws
 *   When an incorrect `test`, `index`, or `parent` is given; there is no error
 *   thrown when `element` is not a node or not an element.
 */
export const isElement =
  // Note: overloads in JSDoc can’t yet use different `@template`s.
  /**
   * @type {(
   *   (<Condition extends TestFunction>(element: unknown, test: Condition, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is Element & Predicate<Condition, Element>) &
   *   (<Condition extends string>(element: unknown, test: Condition, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is Element & {tagName: Condition}) &
   *   ((element?: null | undefined) => false) &
   *   ((element: unknown, test?: null | undefined, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is Element) &
   *   ((element: unknown, test?: Test, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => boolean)
   * )}
   */
  (
    /**
     * @param {unknown} [element]
     * @param {Test | undefined} [test]
     * @param {number | null | undefined} [index]
     * @param {Parents | null | undefined} [parent]
     * @param {unknown} [context]
     * @returns {boolean}
     */
    // eslint-disable-next-line max-params
    function (element, test, index, parent, context) {
      const check = convertElement(test)

      if (
        index !== null &&
        index !== undefined &&
        (typeof index !== 'number' ||
          index < 0 ||
          index === Number.POSITIVE_INFINITY)
      ) {
        throw new Error('Expected positive finite `index`')
      }

      if (
        parent !== null &&
        parent !== undefined &&
        (!parent.type || !parent.children)
      ) {
        throw new Error('Expected valid `parent`')
      }

      if (
        (index === null || index === undefined) !==
        (parent === null || parent === undefined)
      ) {
        throw new Error('Expected both `index` and `parent`')
      }

      return looksLikeAnElement(element)
        ? check.call(context, element, index, parent)
        : false
    }
  )

/**
 * Generate a check from a test.
 *
 * Useful if you’re going to test many nodes, for example when creating a
 * utility where something else passes a compatible test.
 *
 * The created function is a bit faster because it expects valid input only:
 * an `element`, `index`, and `parent`.
 *
 * @param test
 *   A test for a specific element.
 * @returns
 *   A check.
 */
export const convertElement =
  // Note: overloads in JSDoc can’t yet use different `@template`s.
  /**
   * @type {(
   *   (<Condition extends TestFunction>(test: Condition) => (element: unknown, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is Element & Predicate<Condition, Element>) &
   *   (<Condition extends string>(test: Condition) => (element: unknown, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is Element & {tagName: Condition}) &
   *   ((test?: null | undefined) => (element?: unknown, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is Element) &
   *   ((test?: Test) => Check)
   * )}
   */
  (
    /**
     * @param {Test | null | undefined} [test]
     * @returns {Check}
     */
    function (test) {
      if (test === null || test === undefined) {
        return element
      }

      if (typeof test === 'string') {
        return tagNameFactory(test)
      }

      // Assume array.
      if (typeof test === 'object') {
        return anyFactory(test)
      }

      if (typeof test === 'function') {
        return castFactory(test)
      }

      throw new Error('Expected function, string, or array as `test`')
    }
  )

/**
 * Handle multiple tests.
 *
 * @param {Array<TestFunction | string>} tests
 * @returns {Check}
 */
function anyFactory(tests) {
  /** @type {Array<Check>} */
  const checks = []
  let index = -1

  while (++index < tests.length) {
    checks[index] = convertElement(tests[index])
  }

  return castFactory(any)

  /**
   * @this {unknown}
   * @type {TestFunction}
   */
  function any(...parameters) {
    let index = -1

    while (++index < checks.length) {
      if (checks[index].apply(this, parameters)) return true
    }

    return false
  }
}

/**
 * Turn a string into a test for an element with a certain type.
 *
 * @param {string} check
 * @returns {Check}
 */
function tagNameFactory(check) {
  return castFactory(tagName)

  /**
   * @param {Element} element
   * @returns {boolean}
   */
  function tagName(element) {
    return element.tagName === check
  }
}

/**
 * Turn a custom test into a test for an element that passes that test.
 *
 * @param {TestFunction} testFunction
 * @returns {Check}
 */
function castFactory(testFunction) {
  return check

  /**
   * @this {unknown}
   * @type {Check}
   */
  function check(value, index, parent) {
    return Boolean(
      looksLikeAnElement(value) &&
        testFunction.call(
          this,
          value,
          typeof index === 'number' ? index : undefined,
          parent || undefined
        )
    )
  }
}

/**
 * Make sure something is an element.
 *
 * @param {unknown} element
 * @returns {element is Element}
 */
function element(element) {
  return Boolean(
    element &&
      typeof element === 'object' &&
      'type' in element &&
      element.type === 'element' &&
      'tagName' in element &&
      typeof element.tagName === 'string'
  )
}

/**
 * @param {unknown} value
 * @returns {value is Element}
 */
function looksLikeAnElement(value) {
  return (
    value !== null &&
    typeof value === 'object' &&
    'type' in value &&
    'tagName' in value
  )
}
