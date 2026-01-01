# rehype-stringify

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

**[rehype][]** plugin to add support for serializing to HTML.

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
  * [`unified().use(rehypeStringify[, options])`](#unifieduserehypestringify-options)
  * [`CharacterReferences`](#characterreferences)
  * [`Options`](#options)
* [Syntax](#syntax)
* [Syntax tree](#syntax-tree)
* [Types](#types)
* [Compatibility](#compatibility)
* [Security](#security)
* [Contribute](#contribute)
* [Sponsor](#sponsor)
* [License](#license)

## What is this?

This package is a [unified][] ([rehype][]) plugin that defines how to take a
syntax tree as input and turn it into serialized HTML.
When it’s used, HTML is serialized as the final result.

See [the monorepo readme][rehype] for info on what the rehype ecosystem is.

## When should I use this?

This plugin adds support to unified for serializing HTML.
If you also need to parse HTML, you can alternatively use
[`rehype`][rehype-core], which combines unified,
[`rehype-parse`][rehype-parse], and this plugin.

When you are in a browser, trust your content, don’t need formatting options,
and value a smaller bundle size, you can use
[`rehype-dom-stringify`][rehype-dom-stringify] instead.

If you don’t use plugins and have access to a syntax tree, you can directly use
[`hast-util-to-html`][hast-util-to-html], which is used inside this plugin.
rehype focusses on making it easier to transform content by abstracting such
internals away.

A different plugin, [`rehype-format`][rehype-format], improves the readability
of HTML source code as it adds insignificant but pretty whitespace between
elements.
There is also the preset [`rehype-minify`][rehype-minify] for when you want the
inverse: minified and mangled HTML.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install rehype-stringify
```

In Deno with [`esm.sh`][esmsh]:

```js
import rehypeStringify from 'https://esm.sh/rehype-stringify@10'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import rehypeStringify from 'https://esm.sh/rehype-stringify@10?bundle'
</script>
```

## Use

Say we have the following module `example.js`:

```js
import remarkRehype from 'remark-rehype'
import rehypeStringify from 'rehype-stringify'
import remarkGfm from 'remark-gfm'
import remarkParse from 'remark-parse'
import {unified} from 'unified'

const file = await unified()
  .use(remarkParse)
  .use(remarkGfm)
  .use(remarkRehype)
  .use(rehypeStringify)
  .process('# Hi\n\n*Hello*, world!')

console.log(String(file))
```

…running that with `node example.js` yields:

```html
<h1>Hi</h1>
<p><em>Hello</em>, world!</p>
```

## API

This package exports no identifiers.
The default export is [`rehypeStringify`][api-rehype-stringify].

### `unified().use(rehypeStringify[, options])`

Plugin to add support for serializing to HTML.

###### Parameters

* `options` ([`Options`][api-options], optional)
  — configuration

###### Returns

Nothing (`undefined`).

### `CharacterReferences`

How to serialize character references (TypeScript type).

> ⚠️ **Note**: `omitOptionalSemicolons` creates what HTML calls “parse errors”
> but is otherwise still valid HTML — don’t use this except when building a
> minifier.
> Omitting semicolons is possible for certain named and numeric references in
> some cases.

> ⚠️ **Note**: `useNamedReferences` can be omitted when using
> `useShortestReferences`.

###### Fields

* `useNamedReferences` (`boolean`, default: `false`)
  — prefer named character references (`&amp;`) where possible
* `omitOptionalSemicolons` (`boolean`, default: `false`)
  — whether to omit semicolons when possible
* `useShortestReferences` (`boolean`, default: `false`)
  — prefer the shortest possible reference, if that results in less bytes

### `Options`

Configuration (TypeScript type).

> ⚠️ **Danger**: only set `allowDangerousCharacters` and `allowDangerousHtml` if
> you completely trust the content.

> 👉 **Note**: `allowParseErrors`, `bogusComments`, `tightAttributes`, and
> `tightDoctype`
> intentionally create parse errors in markup (how parse errors are handled is
> well defined, so this works but isn’t pretty).

> 👉 **Note**: this is not an XML serializer.
> It supports SVG as embedded in HTML.
> It does not support the features available in XML.
> Use [`xast-util-to-xml`][xast-util-to-xml] to serialize XML.

###### Fields

* `allowDangerousCharacters` (`boolean`, default: `false`)
  — do not encode some characters which cause XSS vulnerabilities in older
  browsers
* `allowDangerousHtml` (`boolean`, default: `false`)
  — allow [`Raw`][raw] nodes and insert them as raw HTML; when `false`, `Raw`
  nodes are encoded
* `allowParseErrors` (`boolean`, default: `false`)
  — do not encode characters which cause parse errors (even though they
  work), to save bytes; not used in the SVG space.
* `bogusComments` (`boolean`, default: `false`)
  — use “bogus comments” instead of comments to save byes: `<?charlie>`
  instead of `<!--charlie-->`
* `characterReferences` ([`CharacterReferences`][api-character-references],
  optional)
  — configure how to serialize character references
* `closeEmptyElements` (`boolean`, default: `false`)
  — close SVG elements without any content with slash (`/`) on the opening
  tag instead of an end tag: `<circle />` instead of `<circle></circle>`;
  see `tightSelfClosing` to control whether a space is used before the slash;
  not used in the HTML space
* `closeSelfClosing` (`boolean`, default: `false`)
  — close self-closing nodes with an extra slash (`/`): `<img />` instead of
  `<img>`; see `tightSelfClosing` to control whether a space is used before
  the slash; not used in the SVG space.
* `collapseEmptyAttributes` (`boolean`, default: `false`)
  — collapse empty attributes: get `class` instead of `class=""`; not used in
  the SVG space; boolean attributes (such as `hidden`) are always collapsed
* `omitOptionalTags` (`boolean`, default: `false`)
  — omit optional opening and closing tags; to illustrate, in
  `<ol><li>one</li><li>two</li></ol>`, both `</li>` closing tags can be
  omitted, the first because it’s followed by another `li`, the last because
  it’s followed by nothing; not used in the SVG space
* `preferUnquoted` (`boolean`, default: `false`)
  — leave attributes unquoted if that results in less bytes; not used in the
  SVG space
* `quote` (`'"'` or `"'"`, default: `'"'`)
  — preferred quote to use
* `quoteSmart` (`boolean`, default: `false`)
  — use the other quote if that results in less bytes
* `space` (`'html'` or `'svg'`, default: `'html'`)
  — which space the document is in; when an `<svg>` element is found in the
  HTML space, this package already automatically switches to and from the SVG
* `tightAttributes` (`boolean`, default: `false`)
  — join attributes together, without whitespace, if possible: get
  `class="a b"title="c d"` instead of `class="a b" title="c d"` to save
  bytes; not used in the SVG space
* `tightCommaSeparatedLists` (`boolean`, default: `false`)
  — join known comma-separated attribute values with just a comma (`,`),
  instead of padding them on the right as well (`,␠`, where `␠` represents a
  space)
* `tightDoctype` (`boolean`, default: `false`)
  — drop unneeded spaces in doctypes: `<!doctypehtml>` instead of
  `<!doctype html>` to save bytes
* `tightSelfClosing` (`boolean`, default: `false`).
  — do not use an extra space when closing self-closing elements: `<img/>`
  instead of `<img />`; only used if `closeSelfClosing: true` or
  `closeEmptyElements: true`
* `upperDoctype` (`boolean`, default: `false`).
  — use a `<!DOCTYPE…` instead of `<!doctype…`; useless except for XHTML
* `voids` (`Array<string>`, default:
  [`html-void-elements`][html-void-elements])
  — tag names of elements to serialize without closing tag; not used in the
  SVG space

## Syntax

HTML is serialized according to WHATWG HTML (the living standard), which is also
followed by all browsers.

## Syntax tree

The syntax tree format used in rehype is [hast][].

## Types

This package is fully typed with [TypeScript][].
It exports the additional types
[`CharacterReferences`][api-character-references] and
[`Options`][api-options].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line, `rehype-stringify@^10`,
compatible with Node.js 16.

## Security

As **rehype** works on HTML, and improper use of HTML can open you up to a
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

[downloads-badge]: https://img.shields.io/npm/dm/rehype-stringify.svg

[downloads]: https://www.npmjs.com/package/rehype-stringify

[size-badge]: https://img.shields.io/bundlejs/size/rehype-stringify

[size]: https://bundlejs.com/?q=rehype-stringify

[sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[backers-badge]: https://opencollective.com/unified/backers/badge.svg

[collective]: https://opencollective.com/unified

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[chat]: https://github.com/rehypejs/rehype/discussions

[health]: https://github.com/rehypejs/.github

[security]: https://github.com/rehypejs/.github/blob/main/security.md

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

[rehype-parse]: ../rehype-parse/

[rehype-core]: ../rehype/

[rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[rehype-format]: https://github.com/rehypejs/rehype-format

[rehype-minify]: https://github.com/rehypejs/rehype-minify

[rehype-dom-stringify]: https://github.com/rehypejs/rehype-dom/tree/main/packages/rehype-dom-stringify

[hast-util-to-html]: https://github.com/syntax-tree/hast-util-to-html

[xast-util-to-xml]: https://github.com/syntax-tree/xast-util-to-xml

[html-void-elements]: https://github.com/wooorm/html-void-elements

<!-- To do: use `remark-rehype` link if that’s released? -->

[raw]: https://github.com/syntax-tree/mdast-util-to-hast?tab=readme-ov-file#raw

[api-character-references]: #characterreferences

[api-options]: #options

[api-rehype-stringify]: #unifieduserehypestringify-options
