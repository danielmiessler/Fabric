/**
 * @import {Root} from 'hast'
 * @import {Options as FromHtmlOptions} from 'hast-util-from-html'
 * @import {Parser, Processor} from 'unified'
 */

/**
 * @typedef {Omit<FromHtmlOptions, 'onerror'> & RehypeParseFields} Options
 *   Configuration.
 *
 * @typedef RehypeParseFields
 *   Extra fields.
 * @property {boolean | null | undefined} [emitParseErrors=false]
 *   Whether to emit parse errors while parsing (default: `false`).
 *
 *   > 👉 **Note**: parse errors are currently being added to HTML.
 *   > Not all errors emitted by parse5 (or us) are specced yet.
 *   > Some documentation may still be missing.
 */

import {fromHtml} from 'hast-util-from-html'

/**
 * Plugin to add support for parsing from HTML.
 *
 * > 👉 **Note**: this is not an XML parser.
 * > It supports SVG as embedded in HTML.
 * > It does not support the features available in XML.
 * > Passing SVG files might break but fragments of modern SVG should be fine.
 * > Use [`xast-util-from-xml`][xast-util-from-xml] to parse XML.
 *
 * @param {Options | null | undefined} [options]
 *   Configuration (optional).
 * @returns {undefined}
 *   Nothing.
 */
export default function rehypeParse(options) {
  /** @type {Processor<Root>} */
  // @ts-expect-error: TS in JSDoc generates wrong types if `this` is typed regularly.
  const self = this
  const {emitParseErrors, ...settings} = {...self.data('settings'), ...options}

  self.parser = parser

  /**
   * @type {Parser<Root>}
   */
  function parser(document, file) {
    return fromHtml(document, {
      ...settings,
      onerror: emitParseErrors
        ? function (message) {
            if (file.path) {
              message.name = file.path + ':' + message.name
              message.file = file.path
            }

            file.messages.push(message)
          }
        : undefined
    })
  }
}
