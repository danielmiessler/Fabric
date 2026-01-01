import type {Root} from 'hast'
import type {Plugin} from 'unified'
import type {Options} from 'hast-util-to-html'

export type {CharacterReferences, Options} from 'hast-util-to-html'

/**
 * Plugin to add support for serializing as HTML.
 *
 * @this
 *   Unified processor.
 * @param
 *   Configuration (optional).
 * @returns
 *   Nothing.
 */
declare const rehypeStringify: Plugin<
  [(Options | null | undefined)?],
  Root,
  string
>
export default rehypeStringify

// Add custom settings supported when `rehype-stringify` is added.
declare module 'unified' {
  interface Settings extends Options {}
}
