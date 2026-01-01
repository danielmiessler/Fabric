import type {Position} from 'unist'
import type {VFile} from 'vfile'

export {fromParse5} from './lib/index.js'

/**
 * Configuration.
 */
export interface Options {
  /**
   * File used to add positional info to nodes (optional).
   *
   * If given, the file should represent the original HTML source.
   */
  file?: VFile | null | undefined

  /**
   * Which space the document is in (default: `'html'`).
   *
   * When an `<svg>` element is found in the HTML space, this package already
   * automatically switches to and from the SVG space when entering and exiting
   * it.
   */
  space?: Space | null | undefined

  /**
   * Whether to add extra positional info about starting tags, closing tags,
   * and attributes to elements (default: `false`).
   *
   * > 👉 **Note**: only used when `file` is given.
   */
  verbose?: boolean | null | undefined
}

/**
 * Namespace.
 */
export type Space = 'html' | 'svg'

// Register data on hast.
declare module 'hast' {
  interface ElementData {
    position: {
      /**
       * Positional info of the start tag of an element.
       *
       * Field added by `hast-util-from-parse5` (a utility used inside
       * `rehype-parse` responsible for parsing HTML), when passing
       * `verbose: true`.
       */
      opening?: Position | undefined

      /**
       * Positional info of the end tag of an element.
       *
       * Field added by `hast-util-from-parse5` (a utility used inside
       * `rehype-parse` responsible for parsing HTML), when passing
       * `verbose: true`.
       */
      closing?: Position | undefined

      /**
       * Positional info of the properties of an element.
       *
       * Field added by `hast-util-from-parse5` (a utility used inside
       * `rehype-parse` responsible for parsing HTML), when passing
       * `verbose: true`.
       */
      properties?: Record<string, Position | undefined> | undefined
    }
  }

  interface RootData {
    /**
     * Whether the document was using quirksmode.
     *
     * Field added by `hast-util-from-parse5` (a utility used inside
     * `rehype-parse` responsible for parsing HTML).
     */
    quirksMode?: boolean | undefined
  }
}
