# rehype-parse

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

**[rehype][]** plugin to add support for parsing from HTML.

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
  * [`unified().use(rehypeParse[, options])`](#unifieduserehypeparse-options)
  * [`ErrorCode`](#errorcode)
  * [`ErrorSeverity`](#errorseverity)
  * [`Options`](#options)
* [Examples](#examples)
  * [Example: fragment versus document](#example-fragment-versus-document)
  * [Example: whitespace around and inside `<html>`](#example-whitespace-around-and-inside-html)
  * [Example: parse errors](#example-parse-errors)
* [Syntax](#syntax)
* [Syntax tree](#syntax-tree)
* [Types](#types)
* [Compatibility](#compatibility)
* [Security](#security)
* [Contribute](#contribute)
* [Sponsor](#sponsor)
* [License](#license)

## What is this?

This package is a [unified][] ([rehype][]) plugin that defines how to take HTML
as input and turn it into a syntax tree.
When it’s used, HTML can be parsed and other rehype plugins can be used after
it.

See [the monorepo readme][rehype] for info on what the rehype ecosystem is.

## When should I use this?

This plugin adds support to unified for parsing HTML.
If you also need to serialize HTML, you can alternatively use
[`rehype`][rehype-core], which combines unified, this plugin, and
[`rehype-stringify`][rehype-stringify].

When you are in a browser, trust your content, don’t need positional info, and
value a smaller bundle size, you can use [`rehype-dom-parse`][rehype-dom-parse]
instead.

If you don’t use plugins and want to access the syntax tree, you can directly
use [`hast-util-from-html`][hast-util-from-html], which is used inside this
plugin.
rehype focusses on making it easier to transform content by abstracting such
internals away.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install rehype-parse
```

In Deno with [`esm.sh`][esmsh]:

```js
import rehypeParse from 'https://esm.sh/rehype-parse@9'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import rehypeParse from 'https://esm.sh/rehype-parse@9?bundle'
</script>
```

## Use

Say we have the following module `example.js`:

```js
import rehypeParse from 'rehype-parse'
import rehypeRemark from 'rehype-remark'
import remarkStringify from 'remark-stringify'
import {unified} from 'unified'

const file = await unified()
  .use(rehypeParse)
  .use(rehypeRemark)
  .use(remarkStringify)
  .process('<h1>Hello, world!</h1>')

console.log(String(file))
```

…running that with `node example.js` yields:

```markdown
# Hello, world!
```

## API

This package exports no identifiers.
The default export is [`rehypeParse`][api-rehype-parse].

### `unified().use(rehypeParse[, options])`

Plugin to add support for parsing from HTML.

###### Parameters

* `options` ([`Options`][api-options], optional)
  — configuration

###### Returns

Nothing (`undefined`).

### `ErrorCode`

Known names of [parse errors][parse-errors] (TypeScript type).

For a bit more info on each error, see
[`hast-util-from-html`][hast-util-from-html-errors].

###### Type

```ts
type ErrorCode =
  | 'abandonedHeadElementChild'
  | 'abruptClosingOfEmptyComment'
  | 'abruptDoctypePublicIdentifier'
  | 'abruptDoctypeSystemIdentifier'
  | 'absenceOfDigitsInNumericCharacterReference'
  | 'cdataInHtmlContent'
  | 'characterReferenceOutsideUnicodeRange'
  | 'closingOfElementWithOpenChildElements'
  | 'controlCharacterInInputStream'
  | 'controlCharacterReference'
  | 'disallowedContentInNoscriptInHead'
  | 'duplicateAttribute'
  | 'endTagWithAttributes'
  | 'endTagWithTrailingSolidus'
  | 'endTagWithoutMatchingOpenElement'
  | 'eofBeforeTagName'
  | 'eofInCdata'
  | 'eofInComment'
  | 'eofInDoctype'
  | 'eofInElementThatCanContainOnlyText'
  | 'eofInScriptHtmlCommentLikeText'
  | 'eofInTag'
  | 'incorrectlyClosedComment'
  | 'incorrectlyOpenedComment'
  | 'invalidCharacterSequenceAfterDoctypeName'
  | 'invalidFirstCharacterOfTagName'
  | 'misplacedDoctype'
  | 'misplacedStartTagForHeadElement'
  | 'missingAttributeValue'
  | 'missingDoctype'
  | 'missingDoctypeName'
  | 'missingDoctypePublicIdentifier'
  | 'missingDoctypeSystemIdentifier'
  | 'missingEndTagName'
  | 'missingQuoteBeforeDoctypePublicIdentifier'
  | 'missingQuoteBeforeDoctypeSystemIdentifier'
  | 'missingSemicolonAfterCharacterReference'
  | 'missingWhitespaceAfterDoctypePublicKeyword'
  | 'missingWhitespaceAfterDoctypeSystemKeyword'
  | 'missingWhitespaceBeforeDoctypeName'
  | 'missingWhitespaceBetweenAttributes'
  | 'missingWhitespaceBetweenDoctypePublicAndSystemIdentifiers'
  | 'nestedComment'
  | 'nestedNoscriptInHead'
  | 'nonConformingDoctype'
  | 'nonVoidHtmlElementStartTagWithTrailingSolidus'
  | 'noncharacterCharacterReference'
  | 'noncharacterInInputStream'
  | 'nullCharacterReference'
  | 'openElementsLeftAfterEof'
  | 'surrogateCharacterReference'
  | 'surrogateInInputStream'
  | 'unexpectedCharacterAfterDoctypeSystemIdentifier'
  | 'unexpectedCharacterInAttributeName'
  | 'unexpectedCharacterInUnquotedAttributeValue'
  | 'unexpectedEqualsSignBeforeAttributeName'
  | 'unexpectedNullCharacter'
  | 'unexpectedQuestionMarkInsteadOfTagName'
  | 'unexpectedSolidusInTag'
  | 'unknownNamedCharacterReference'
```

### `ErrorSeverity`

Error severity (TypeScript type).

* `0` or `false`
  — turn the parse error off
* `1` or `true`
  — turn the parse error into a warning
* `2`
  — turn the parse error into an actual error: processing stops

###### Type

```ts
type ErrorSeverity = boolean | 0 | 1 | 2
```

### `Options`

Configuration (TypeScript type).

> 👉 **Note**: this is not an XML parser.
> It supports SVG as embedded in HTML.
> It does not support the features available in XML.
> Passing SVG files might break but fragments of modern SVG should be fine.
> Use [`xast-util-from-xml`][xast-util-from-xml] to parse XML.

###### Fields

* `fragment` (`boolean`, default: `false`)
  — whether to parse as a fragment; by default unopened `html`, `head`, and
  `body` elements are opened
* `emitParseErrors` (`boolean`, default: `false`)
  — whether to emit [parse errors][parse-errors] while parsing
* `space` (`'html'` or `'svg'`, default: `'html'`)
  — which space the document is in
* `verbose` (`boolean`, default: `false`)
  — add extra positional info about attributes, start tags, and end tags
* [`[key in ErrorCode]`][api-error-code]
  ([`ErrorSeverity`][api-error-severity], default: `1` if
  `options.emitParseErrors`, otherwise `0`)
  — configure specific [parse errors][parse-errors]

## Examples

### Example: fragment versus document

The following example shows the difference between parsing as a document and
parsing as a fragment:

```js
import rehypeParse from 'rehype-parse'
import rehypeStringify from 'rehype-stringify'
import {unified} from 'unified'

const doc = '<title>Hi!</title><h1>Hello!</h1>'

console.log(
  String(
    await unified()
      .use(rehypeParse, {fragment: true})
      .use(rehypeStringify)
      .process(doc)
  )
)

console.log(
  String(
    await unified()
      .use(rehypeParse, {fragment: false})
      .use(rehypeStringify)
      .process(doc)
  )
)
```

…yields:

```html
<title>Hi!</title><h1>Hello!</h1>
```

```html
<html><head><title>Hi!</title></head><body><h1>Hello!</h1></body></html>
```

> 👉 **Note**: observe that when a whole document is expected (second example),
> missing elements are opened and closed.

### Example: whitespace around and inside `<html>`

The following example shows how whitespace is handled when around and directly
inside the `<html>` element:

```js
import rehypeParse from 'rehype-parse'
import rehypeStringify from 'rehype-stringify'
import {unified} from 'unified'

const doc = `<!doctype html>
<html lang=en>
  <head>
    <title>Hi!</title>
  </head>
  <body>
    <h1>Hello!</h1>
  </body>
</html>`

console.log(
  String(await unified().use(rehypeParse).use(rehypeStringify).process(doc))
)
```

…yields (where `␠` represents a space character):

```html
<!doctype html><html lang="en"><head>
    <title>Hi!</title>
  </head>
  <body>
    <h1>Hello!</h1>
␠␠
</body></html>
```

> 👉 **Note**: observe that the line ending before `<html>` is ignored, the line
> ending and two spaces before `<head>` is moved inside it, and the line ending
> after `</body>` is moved before it.

This behavior is described by the HTML standard (see the section 13.2.6.4.1
“The ‘initial’ insertion mode” and adjacent states) which rehype follows.

The changes to this meaningless whitespace should not matter, except when
formatting markup, in which case [`rehype-format`][rehype-format] can be used to
improve the source code.

### Example: parse errors

The following example shows how HTML parse errors can be enabled and configured:

```js
import rehypeParse from 'rehype-parse'
import rehypeStringify from 'rehype-stringify'
import {unified} from 'unified'
import {reporter} from 'vfile-reporter'

const file = await unified()
  .use(rehypeParse, {
    emitParseErrors: true, // Emit all.
    missingWhitespaceBeforeDoctypeName: 2, // Mark one as a fatal error.
    nonVoidHtmlElementStartTagWithTrailingSolidus: false // Ignore one.
  })
  .use(rehypeStringify).process(`<!doctypehtml>
<title class="a" class="b">Hello…</title>
<h1/>World!</h1>`)

console.log(reporter(file))
```

…yields:

```html
1:10-1:10 error   Missing whitespace before doctype name missing-whitespace-before-doctype-name hast-util-from-html
2:23-2:23 warning Unexpected duplicate attribute         duplicate-attribute                    hast-util-from-html

2 messages (✖ 1 error, ⚠ 1 warning)
```

> 🧑‍🏫 **Info**: messages in unified are warnings instead of errors.
> Other linters (such as ESLint) almost always use errors.
> Why?
> Those tools *only* check code style.
> They don’t generate, transform, and format code, which is what rehype and
> unified focus on, too.
> Errors in unified mean the same as an exception in your JavaScript code: a
> crash.
> That’s why we use warnings instead, because we continue checking more HTML and
> continue running more plugins.

## Syntax

HTML is parsed according to WHATWG HTML (the living standard), which is also
followed by all browsers.

## Syntax tree

The syntax tree format used in rehype is [hast][].

## Types

This package is fully typed with [TypeScript][].
It exports the additional types [`ErrorCode`][api-error-code],
[`ErrorSeverity`][api-error-severity], and
[`Options`][api-options].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line, `rehype-parse@^9`,
compatible with Node.js 16.

## Security

As **rehype** works on HTML and improper use of HTML can open you up to a
[cross-site scripting (XSS)][xss] attack, use of rehype can also be unsafe.
Use [`rehype-sanitize`][rehype-sanitize] to make the tree safe.

Use of rehype plugins could also open you up to other attacks.
Carefully assess each plugin and the risks involved in using them.

For info on how to submit a report, see our [security policy][security].

## Contribute

See [`contributing.md`][contributing] in [`rehypejs/.github`][health] for ways
to get started.
See [`support.md`][support] for ways to get help.

This project has a [code of conduct][coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## Sponsor

Support this effort and give back by sponsoring on [OpenCollective][collective]!

<table>
<tr valign="middle">
<td width="20%" align="center" rowspan="2" colspan="2">
  <a href="https://vercel.com">Vercel</a><br><br>
  <a href="https://vercel.com"><img src="https://avatars1.githubusercontent.com/u/14985020?s=256&v=4" width="128"></a>
</td>
<td width="20%" align="center" rowspan="2" colspan="2">
  <a href="https://motif.land">Motif</a><br><br>
  <a href="https://motif.land"><img src="https://avatars1.githubusercontent.com/u/74457950?s=256&v=4" width="128"></a>
</td>
<td width="20%" align="center" rowspan="2" colspan="2">
  <a href="https://www.hashicorp.com">HashiCorp</a><br><br>
  <a href="https://www.hashicorp.com"><img src="https://avatars1.githubusercontent.com/u/761456?s=256&v=4" width="128"></a>
</td>
<td width="20%" align="center" rowspan="2" colspan="2">
  <a href="https://www.gitbook.com">GitBook</a><br><br>
  <a href="https://www.gitbook.com"><img src="https://avatars1.githubusercontent.com/u/7111340?s=256&v=4" width="128"></a>
</td>
<td width="20%" align="center" rowspan="2" colspan="2">
  <a href="https://www.gatsbyjs.org">Gatsby</a><br><br>
  <a href="https://www.gatsbyjs.org"><img src="https://avatars1.githubusercontent.com/u/12551863?s=256&v=4" width="128"></a>
</td>
</tr>
<tr valign="middle">
</tr>
<tr valign="middle">
<td width="20%" align="center" rowspan="2" colspan="2">
  <a href="https://www.netlify.com">Netlify</a><br><br>
  <!--OC has a sharper image-->
  <a href="https://www.netlify.com"><img src="https://images.opencollective.com/netlify/4087de2/logo/256.png" width="128"></a>
</td>
<td width="10%" align="center">
  <a href="https://www.coinbase.com">Coinbase</a><br><br>
  <a href="https://www.coinbase.com"><img src="https://avatars1.githubusercontent.com/u/1885080?s=256&v=4" width="64"></a>
</td>
<td width="10%" align="center">
  <a href="https://themeisle.com">ThemeIsle</a><br><br>
  <a href="https://themeisle.com"><img src="https://avatars1.githubusercontent.com/u/58979018?s=128&v=4" width="64"></a>
</td>
<td width="10%" align="center">
  <a href="https://expo.io">Expo</a><br><br>
  <a href="https://expo.io"><img src="https://avatars1.githubusercontent.com/u/12504344?s=128&v=4" width="64"></a>
</td>
<td width="10%" align="center">
  <a href="https://boostnote.io">Boost Note</a><br><br>
  <a href="https://boostnote.io"><img src="https://images.opencollective.com/boosthub/6318083/logo/128.png" width="64"></a>
</td>
<td width="10%" align="center">
  <a href="https://markdown.space">Markdown Space</a><br><br>
  <a href="https://markdown.space"><img src="https://images.opencollective.com/markdown-space/e1038ed/logo/128.png" width="64"></a>
</td>
<td width="10%" align="center">
  <a href="https://www.holloway.com">Holloway</a><br><br>
  <a href="https://www.holloway.com"><img src="https://avatars1.githubusercontent.com/u/35904294?s=128&v=4" width="64"></a>
</td>
<td width="10%"></td>
<td width="10%"></td>
</tr>
<tr valign="middle">
<td width="100%" align="center" colspan="8">
  <br>
  <a href="https://opencollective.com/unified"><strong>You?</strong></a>
  <br><br>
</td>
</tr>
</table>

## License

[MIT][license] © [Titus Wormer][author]

<!-- Definitions -->

[build-badge]: https://github.com/rehypejs/rehype/workflows/main/badge.svg

[build]: https://github.com/rehypejs/rehype/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/rehypejs/rehype.svg

[coverage]: https://codecov.io/github/rehypejs/rehype

[downloads-badge]: https://img.shields.io/npm/dm/rehype-parse.svg

[downloads]: https://www.npmjs.com/package/rehype-parse

[size-badge]: https://img.shields.io/bundlejs/size/rehype-parse

[size]: https://bundlejs.com/?q=rehype-parse

[sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[backers-badge]: https://opencollective.com/unified/backers/badge.svg

[collective]: https://opencollective.com/unified

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[chat]: https://github.com/rehypejs/rehype/discussions

[security]: https://github.com/rehypejs/.github/blob/main/security.md

[health]: https://github.com/rehypejs/.github

[contributing]: https://github.com/rehypejs/.github/blob/main/contributing.md

[support]: https://github.com/rehypejs/.github/blob/main/support.md

[coc]: https://github.com/rehypejs/.github/blob/main/code-of-conduct.md

[license]: https://github.com/rehypejs/rehype/blob/main/license

[author]: https://wooorm.com

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[npm]: https://docs.npmjs.com/cli/install

[esmsh]: https://esm.sh

[unified]: https://github.com/unifiedjs/unified

[rehype]: https://github.com/rehypejs/rehype

[hast]: https://github.com/syntax-tree/hast

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[typescript]: https://www.typescriptlang.org

[hast-util-from-html]: https://github.com/syntax-tree/hast-util-from-html

[hast-util-from-html-errors]: https://github.com/syntax-tree/hast-util-from-html#optionskey-in-errorcode

[xast-util-from-xml]: https://github.com/syntax-tree/xast-util-from-xml

[rehype-dom-parse]: https://github.com/rehypejs/rehype-dom/tree/main/packages/rehype-dom-parse

[rehype-format]: https://github.com/rehypejs/rehype-format

[rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[parse-errors]: https://html.spec.whatwg.org/multipage/parsing.html#parse-errors

[rehype-core]: ../rehype/

[rehype-stringify]: ../rehype-stringify/

[api-error-code]: #errorcode

[api-error-severity]: #errorseverity

[api-options]: #options

[api-rehype-parse]: #unifieduserehypeparse-options
