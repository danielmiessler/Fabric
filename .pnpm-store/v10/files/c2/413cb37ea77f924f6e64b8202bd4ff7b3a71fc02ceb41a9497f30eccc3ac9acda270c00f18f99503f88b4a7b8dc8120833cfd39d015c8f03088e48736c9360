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
export const isElement: (<Condition extends TestFunction>(element: unknown, test: Condition, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is import("hast").Element & Predicate<Condition, import("hast").Element>) & (<Condition_1 extends string>(element: unknown, test: Condition_1, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is import("hast").Element & {
    tagName: Condition_1;
}) & ((element?: null | undefined) => false) & ((element: unknown, test?: null | undefined, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is import("hast").Element) & ((element: unknown, test?: Test, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => boolean);
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
export const convertElement: (<Condition extends TestFunction>(test: Condition) => (element: unknown, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is import("hast").Element & Predicate<Condition, import("hast").Element>) & (<Condition_1 extends string>(test: Condition_1) => (element: unknown, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is import("hast").Element & {
    tagName: Condition_1;
}) & ((test?: null | undefined) => (element?: unknown, index?: number | null | undefined, parent?: Parents | null | undefined, context?: unknown) => element is import("hast").Element) & ((test?: Test) => Check);
export type Element = import('hast').Element;
export type Parents = import('hast').Parents;
export type Predicate<Fn, Fallback> = Fn extends (value: any) => value is infer Thing ? Thing : Fallback;
/**
 * Check that an arbitrary value is an element.
 */
export type Check = (this: unknown, element?: unknown, index?: number | null | undefined, parent?: Parents | null | undefined) => boolean;
/**
 * Check for an arbitrary element.
 *
 * * when `string`, checks that the element has that tag name
 * * when `function`, see `TestFunction`
 * * when `Array`, checks if one of the subtests pass
 */
export type Test = Array<TestFunction | string> | TestFunction | string | null | undefined;
/**
 * Check if an element passes a test.
 */
export type TestFunction = (this: unknown, element: Element, index?: number | undefined, parent?: Parents | undefined) => boolean | undefined | void;
