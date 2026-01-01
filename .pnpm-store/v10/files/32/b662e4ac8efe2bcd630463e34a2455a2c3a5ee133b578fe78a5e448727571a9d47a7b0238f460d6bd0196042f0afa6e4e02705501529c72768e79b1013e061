/**
 * @typedef ErrorInfo
 *   Info on a `parse5` error.
 * @property {string} reason
 *   Reason of error.
 * @property {string} description
 *   More info on error.
 * @property {false} [url]
 *   Turn off if this is not documented in the html5 spec (optional).
 */

export const errors = {
  /** @type {ErrorInfo} */
  abandonedHeadElementChild: {
    reason: 'Unexpected metadata element after head',
    description:
      'Unexpected element after head. Expected the element before `</head>`',
    url: false
  },
  /** @type {ErrorInfo} */
  abruptClosingOfEmptyComment: {
    reason: 'Unexpected abruptly closed empty comment',
    description: 'Unexpected `>` or `->`. Expected `-->` to close comments'
  },
  /** @type {ErrorInfo} */
  abruptDoctypePublicIdentifier: {
    reason: 'Unexpected abruptly closed public identifier',
    description:
      'Unexpected `>`. Expected a closing `"` or `\'` after the public identifier'
  },
  /** @type {ErrorInfo} */
  abruptDoctypeSystemIdentifier: {
    reason: 'Unexpected abruptly closed system identifier',
    description:
      'Unexpected `>`. Expected a closing `"` or `\'` after the identifier identifier'
  },
  /** @type {ErrorInfo} */
  absenceOfDigitsInNumericCharacterReference: {
    reason: 'Unexpected non-digit at start of numeric character reference',
    description:
      'Unexpected `%c`. Expected `[0-9]` for decimal references or `[0-9a-fA-F]` for hexadecimal references'
  },
  /** @type {ErrorInfo} */
  cdataInHtmlContent: {
    reason: 'Unexpected CDATA section in HTML',
    description:
      'Unexpected `<![CDATA[` in HTML. Remove it, use a comment, or encode special characters instead'
  },
  /** @type {ErrorInfo} */
  characterReferenceOutsideUnicodeRange: {
    reason: 'Unexpected too big numeric character reference',
    description:
      'Unexpectedly high character reference. Expected character references to be at most hexadecimal 10ffff (or decimal 1114111)'
  },
  /** @type {ErrorInfo} */
  closingOfElementWithOpenChildElements: {
    reason: 'Unexpected closing tag with open child elements',
    description:
      'Unexpectedly closing tag. Expected other tags to be closed first',
    url: false
  },
  /** @type {ErrorInfo} */
  controlCharacterInInputStream: {
    reason: 'Unexpected control character',
    description:
      'Unexpected control character `%x`. Expected a non-control code point, 0x00, or ASCII whitespace'
  },
  /** @type {ErrorInfo} */
  controlCharacterReference: {
    reason: 'Unexpected control character reference',
    description:
      'Unexpectedly control character in reference. Expected a non-control code point, 0x00, or ASCII whitespace'
  },
  /** @type {ErrorInfo} */
  disallowedContentInNoscriptInHead: {
    reason: 'Disallowed content inside `<noscript>` in `<head>`',
    description:
      'Unexpected text character `%c`. Only use text in `<noscript>`s in `<body>`',
    url: false
  },
  /** @type {ErrorInfo} */
  duplicateAttribute: {
    reason: 'Unexpected duplicate attribute',
    description:
      'Unexpectedly double attribute. Expected attributes to occur only once'
  },
  /** @type {ErrorInfo} */
  endTagWithAttributes: {
    reason: 'Unexpected attribute on closing tag',
    description: 'Unexpected attribute. Expected `>` instead'
  },
  /** @type {ErrorInfo} */
  endTagWithTrailingSolidus: {
    reason: 'Unexpected slash at end of closing tag',
    description: 'Unexpected `%c-1`. Expected `>` instead'
  },
  /** @type {ErrorInfo} */
  endTagWithoutMatchingOpenElement: {
    reason: 'Unexpected unopened end tag',
    description: 'Unexpected end tag. Expected no end tag or another end tag',
    url: false
  },
  /** @type {ErrorInfo} */
  eofBeforeTagName: {
    reason: 'Unexpected end of file',
    description: 'Unexpected end of file. Expected tag name instead'
  },
  /** @type {ErrorInfo} */
  eofInCdata: {
    reason: 'Unexpected end of file in CDATA',
    description: 'Unexpected end of file. Expected `]]>` to close the CDATA'
  },
  /** @type {ErrorInfo} */
  eofInComment: {
    reason: 'Unexpected end of file in comment',
    description: 'Unexpected end of file. Expected `-->` to close the comment'
  },
  /** @type {ErrorInfo} */
  eofInDoctype: {
    reason: 'Unexpected end of file in doctype',
    description:
      'Unexpected end of file. Expected a valid doctype (such as `<!doctype html>`)'
  },
  /** @type {ErrorInfo} */
  eofInElementThatCanContainOnlyText: {
    reason: 'Unexpected end of file in element that can only contain text',
    description: 'Unexpected end of file. Expected text or a closing tag',
    url: false
  },
  /** @type {ErrorInfo} */
  eofInScriptHtmlCommentLikeText: {
    reason: 'Unexpected end of file in comment inside script',
    description: 'Unexpected end of file. Expected `-->` to close the comment'
  },
  /** @type {ErrorInfo} */
  eofInTag: {
    reason: 'Unexpected end of file in tag',
    description: 'Unexpected end of file. Expected `>` to close the tag'
  },
  /** @type {ErrorInfo} */
  incorrectlyClosedComment: {
    reason: 'Incorrectly closed comment',
    description: 'Unexpected `%c-1`. Expected `-->` to close the comment'
  },
  /** @type {ErrorInfo} */
  incorrectlyOpenedComment: {
    reason: 'Incorrectly opened comment',
    description: 'Unexpected `%c`. Expected `<!--` to open the comment'
  },
  /** @type {ErrorInfo} */
  invalidCharacterSequenceAfterDoctypeName: {
    reason: 'Invalid sequence after doctype name',
    description: 'Unexpected sequence at `%c`. Expected `public` or `system`'
  },
  /** @type {ErrorInfo} */
  invalidFirstCharacterOfTagName: {
    reason: 'Invalid first character in tag name',
    description: 'Unexpected `%c`. Expected an ASCII letter instead'
  },
  /** @type {ErrorInfo} */
  misplacedDoctype: {
    reason: 'Misplaced doctype',
    description: 'Unexpected doctype. Expected doctype before head',
    url: false
  },
  /** @type {ErrorInfo} */
  misplacedStartTagForHeadElement: {
    reason: 'Misplaced `<head>` start tag',
    description:
      'Unexpected start tag `<head>`. Expected `<head>` directly after doctype',
    url: false
  },
  /** @type {ErrorInfo} */
  missingAttributeValue: {
    reason: 'Missing attribute value',
    description:
      'Unexpected `%c-1`. Expected an attribute value or no `%c-1` instead'
  },
  /** @type {ErrorInfo} */
  missingDoctype: {
    reason: 'Missing doctype before other content',
    description: 'Expected a `<!doctype html>` before anything else',
    url: false
  },
  /** @type {ErrorInfo} */
  missingDoctypeName: {
    reason: 'Missing doctype name',
    description: 'Unexpected doctype end at `%c`. Expected `html` instead'
  },
  /** @type {ErrorInfo} */
  missingDoctypePublicIdentifier: {
    reason: 'Missing public identifier in doctype',
    description: 'Unexpected `%c`. Expected identifier for `public` instead'
  },
  /** @type {ErrorInfo} */
  missingDoctypeSystemIdentifier: {
    reason: 'Missing system identifier in doctype',
    description:
      'Unexpected `%c`. Expected identifier for `system` instead (suggested: `"about:legacy-compat"`)'
  },
  /** @type {ErrorInfo} */
  missingEndTagName: {
    reason: 'Missing name in end tag',
    description: 'Unexpected `%c`. Expected an ASCII letter instead'
  },
  /** @type {ErrorInfo} */
  missingQuoteBeforeDoctypePublicIdentifier: {
    reason: 'Missing quote before public identifier in doctype',
    description: 'Unexpected `%c`. Expected `"` or `\'` instead'
  },
  /** @type {ErrorInfo} */
  missingQuoteBeforeDoctypeSystemIdentifier: {
    reason: 'Missing quote before system identifier in doctype',
    description: 'Unexpected `%c`. Expected `"` or `\'` instead'
  },
  /** @type {ErrorInfo} */
  missingSemicolonAfterCharacterReference: {
    reason: 'Missing semicolon after character reference',
    description: 'Unexpected `%c`. Expected `;` instead'
  },
  /** @type {ErrorInfo} */
  missingWhitespaceAfterDoctypePublicKeyword: {
    reason: 'Missing whitespace after public identifier in doctype',
    description: 'Unexpected `%c`. Expected ASCII whitespace instead'
  },
  /** @type {ErrorInfo} */
  missingWhitespaceAfterDoctypeSystemKeyword: {
    reason: 'Missing whitespace after system identifier in doctype',
    description: 'Unexpected `%c`. Expected ASCII whitespace instead'
  },
  /** @type {ErrorInfo} */
  missingWhitespaceBeforeDoctypeName: {
    reason: 'Missing whitespace before doctype name',
    description: 'Unexpected `%c`. Expected ASCII whitespace instead'
  },
  /** @type {ErrorInfo} */
  missingWhitespaceBetweenAttributes: {
    reason: 'Missing whitespace between attributes',
    description: 'Unexpected `%c`. Expected ASCII whitespace instead'
  },
  /** @type {ErrorInfo} */
  missingWhitespaceBetweenDoctypePublicAndSystemIdentifiers: {
    reason:
      'Missing whitespace between public and system identifiers in doctype',
    description: 'Unexpected `%c`. Expected ASCII whitespace instead'
  },
  /** @type {ErrorInfo} */
  nestedComment: {
    reason: 'Unexpected nested comment',
    description: 'Unexpected `<!--`. Expected `-->`'
  },
  /** @type {ErrorInfo} */
  nestedNoscriptInHead: {
    reason: 'Unexpected nested `<noscript>` in `<head>`',
    description:
      'Unexpected `<noscript>`. Expected a closing tag or a meta element',
    url: false
  },
  /** @type {ErrorInfo} */
  nonConformingDoctype: {
    reason: 'Unexpected non-conforming doctype declaration',
    description:
      'Expected `<!doctype html>` or `<!doctype html system "about:legacy-compat">`',
    url: false
  },
  /** @type {ErrorInfo} */
  nonVoidHtmlElementStartTagWithTrailingSolidus: {
    reason: 'Unexpected trailing slash on start tag of non-void element',
    description: 'Unexpected `/`. Expected `>` instead'
  },
  /** @type {ErrorInfo} */
  noncharacterCharacterReference: {
    reason:
      'Unexpected noncharacter code point referenced by character reference',
    description: 'Unexpected code point. Do not use noncharacters in HTML'
  },
  /** @type {ErrorInfo} */
  noncharacterInInputStream: {
    reason: 'Unexpected noncharacter character',
    description: 'Unexpected code point `%x`. Do not use noncharacters in HTML'
  },
  /** @type {ErrorInfo} */
  nullCharacterReference: {
    reason: 'Unexpected NULL character referenced by character reference',
    description: 'Unexpected code point. Do not use NULL characters in HTML'
  },
  /** @type {ErrorInfo} */
  openElementsLeftAfterEof: {
    reason: 'Unexpected end of file',
    description: 'Unexpected end of file. Expected closing tag instead',
    url: false
  },
  /** @type {ErrorInfo} */
  surrogateCharacterReference: {
    reason: 'Unexpected surrogate character referenced by character reference',
    description:
      'Unexpected code point. Do not use lone surrogate characters in HTML'
  },
  /** @type {ErrorInfo} */
  surrogateInInputStream: {
    reason: 'Unexpected surrogate character',
    description:
      'Unexpected code point `%x`. Do not use lone surrogate characters in HTML'
  },
  /** @type {ErrorInfo} */
  unexpectedCharacterAfterDoctypeSystemIdentifier: {
    reason: 'Invalid character after system identifier in doctype',
    description: 'Unexpected character at `%c`. Expected `>`'
  },
  /** @type {ErrorInfo} */
  unexpectedCharacterInAttributeName: {
    reason: 'Unexpected character in attribute name',
    description:
      'Unexpected `%c`. Expected whitespace, `/`, `>`, `=`, or probably an ASCII letter'
  },
  /** @type {ErrorInfo} */
  unexpectedCharacterInUnquotedAttributeValue: {
    reason: 'Unexpected character in unquoted attribute value',
    description: 'Unexpected `%c`. Quote the attribute value to include it'
  },
  /** @type {ErrorInfo} */
  unexpectedEqualsSignBeforeAttributeName: {
    reason: 'Unexpected equals sign before attribute name',
    description: 'Unexpected `%c`. Add an attribute name before it'
  },
  /** @type {ErrorInfo} */
  unexpectedNullCharacter: {
    reason: 'Unexpected NULL character',
    description:
      'Unexpected code point `%x`. Do not use NULL characters in HTML'
  },
  /** @type {ErrorInfo} */
  unexpectedQuestionMarkInsteadOfTagName: {
    reason: 'Unexpected question mark instead of tag name',
    description: 'Unexpected `%c`. Expected an ASCII letter instead'
  },
  /** @type {ErrorInfo} */
  unexpectedSolidusInTag: {
    reason: 'Unexpected slash in tag',
    description:
      'Unexpected `%c-1`. Expected it followed by `>` or in a quoted attribute value'
  },
  /** @type {ErrorInfo} */
  unknownNamedCharacterReference: {
    reason: 'Unexpected unknown named character reference',
    description:
      'Unexpected character reference. Expected known named character references'
  }
}
