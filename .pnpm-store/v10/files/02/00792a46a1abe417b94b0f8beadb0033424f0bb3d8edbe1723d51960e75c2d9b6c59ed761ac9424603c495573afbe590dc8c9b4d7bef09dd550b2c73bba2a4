import type {Root} from 'hast'
import type {Plugin} from 'unified'
import type {Options} from './lib/index.js'

export type {ErrorCode, ErrorSeverity} from 'hast-util-from-html'
export type {Options} from './lib/index.js'

/**
 * Plugin to add support for parsing from HTML.
 *
 * @this
 *   Unified processor.
 * @param
 *   Configuration (optional).
 * @returns
 *   Nothing.
 */
declare const rehypeParse: Plugin<[(Options | null | undefined)?], string, Root>
export default rehypeParse

// Add custom settings supported when `rehype-parse` is added.
declare module 'unified' {
  interface Settings extends Options {}
}
