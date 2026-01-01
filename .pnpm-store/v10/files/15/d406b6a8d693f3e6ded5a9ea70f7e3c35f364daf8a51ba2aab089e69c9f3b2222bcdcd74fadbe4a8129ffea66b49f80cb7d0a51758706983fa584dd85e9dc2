import type {VFileMessage} from 'vfile-message'
import type {Options as FromParse5Options} from 'hast-util-from-parse5'
import type {errors} from './errors.js'

/**
 * Known names of parse errors.
 */
export type ErrorCode = keyof typeof errors

/**
 * Error severity:
 *
 * * `0` or `false`
 * — turn the parse error off
 * * `1` or `true`
 * — turn the parse error into a warning
 * * `2`
 * — turn the parse error into an actual error: processing stops.
 */
export type ErrorSeverity = boolean | 0 | 1 | 2

/**
 * Handle parse errors.
 */
export type OnError = (error: VFileMessage) => undefined | void

/**
 * Options that define the severity of errors.
 */
export type ErrorOptions = Partial<
  Record<ErrorCode, ErrorSeverity | null | undefined>
>

/**
 * Configuration.
 */
export interface Options extends ErrorOptions, FromParse5Options {
  /**
   * The `file` field from `hast-util-from-parse5` is not supported.
   */
  file?: never
  /**
   * Specify whether to parse a fragment, instead of a complete document
   * (default: `false`).
   *
   * In document mode, unopened `html`, `head`, and `body` elements are opened
   * in just the right places.
   */
  fragment?: boolean | null | undefined
  /**
   * Call `onerror` with parse errors while parsing (optional).
   *
   * > 👉 **Note**: parse errors are currently being added to HTML.
   * > Not all errors emitted by parse5 (or us) are specced yet.
   * > Some documentation may still be missing.
   *
   * Specific rules can be turned off by setting them to `false` (or `0`).
   * The default, when `emitParseErrors: true`, is `true` (or `1`), and means
   * that rules emit as warnings.
   * Rules can also be configured with `2`, to turn them into fatal errors.
   */
  onerror?: OnError | null | undefined
}
