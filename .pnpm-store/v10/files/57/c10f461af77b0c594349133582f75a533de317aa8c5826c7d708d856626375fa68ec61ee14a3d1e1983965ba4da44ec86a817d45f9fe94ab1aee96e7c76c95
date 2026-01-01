# rehype-slug

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

**[rehype][]** plugin to add `id`s to headings.

## Contents

*   [What is this?](#what-is-this)
*   [When should I use this?](#when-should-i-use-this)
*   [Install](#install)
*   [Use](#use)
*   [API](#api)
    *   [`unified().use(rehypeSlug[, options])`](#unifieduserehypeslug-options)
    *   [`Options`](#options)
*   [Types](#types)
*   [Compatibility](#compatibility)
*   [Security](#security)
*   [Related](#related)
*   [Contribute](#contribute)
*   [License](#license)

## What is this?

This package is a [unified][] ([rehype][]) plugin to add `id`s to headings.
It looks for headings (so `<h1>` through `<h6>`) that do not yet have `id`s
and adds `id` attributes to them based on the text they contain.
The algorithm that does this is [`github-slugger`][github-slugger], which
matches how GitHub works.

**unified** is a project that transforms content with abstract syntax trees
(ASTs).
**rehype** adds support for HTML to unified.
**hast** is the HTML AST that rehype uses.
This is a rehype plugin that adds `id`s to headings in the AST.

## When should I use this?

This plugin is useful when you have relatively long documents and you want to be
able to link to particular sections.

A different plugin, [`rehype-autolink-headings`][rehype-autolink-headings], adds
links to these headings back to themselves, which is useful as it lets users
more easily link to particular sections.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install rehype-slug
```

In Deno with [`esm.sh`][esmsh]:

```js
import rehypeSlug from 'https://esm.sh/rehype-slug@6'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import rehypeSlug from 'https://esm.sh/rehype-slug@6?bundle'
</script>
```

## Use

Say we have the following file `example.html`:

```html
<h1 id=some-id>Lorem ipsum</h1>
<h2>Dolor sit amet 😪</h2>
<h3>consectetur &amp; adipisicing</h3>
<h4>elit</h4>
<h5>elit</h5>
```

…and our module `example.js` looks as follows:

```js
import {read} from 'to-vfile'
import {rehype} from 'rehype'
import rehypeSlug from 'rehype-slug'

const file = await rehype()
  .data('settings', {fragment: true})
  .use(rehypeSlug)
  .process(await read('example.html'))

console.log(String(file))
```

…then running `node example.js` yields:

```html
<h1 id="some-id">Lorem ipsum</h1>
<h2 id="dolor-sit-amet-">Dolor sit amet 😪</h2>
<h3 id="consectetur--adipisicing">consectetur &#x26; adipisicing</h3>
<h4 id="elit">elit</h4>
<h5 id="elit-1">elit</h5>
```

## API

This package exports no identifiers.
The default export is [`rehypeSlug`][api-rehype-slug].

### `unified().use(rehypeSlug[, options])`

Add `id`s to headings.

###### Parameters

*   `options` ([`Options`][api-options], optional)
    — configuration

###### Returns

Transform ([`Transformer`][unified-transformer]).

### `Options`

Configuration (TypeScript type).

###### Fields

*   `prefix` (`string`, default: `''`)
    — prefix to add in front of `id`s

## Types

This package is fully typed with [TypeScript][].
It exports the additional type [`Options`][api-options].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line, `rehype-slug@^6`,
compatible with Node.js 16.

This plugin works with `rehype-parse` version 1+, `rehype-stringify` version 1+,
`rehype` version 1+, and `unified` version 4+.

## Security

Use of `rehype-slug` can open you up to a [cross-site scripting (XSS)][xss]
attack as it sets `id` attributes on headings, which causes what is known
as “DOM clobbering”.
Please use [`rehype-sanitize`][rehype-sanitize] and see its
[Example: headings (DOM clobbering)][rehype-sanitize-example] for information on
how to properly solve it.

## Related

*   [`rehype-autolink-headings`][rehype-autolink-headings]
    — add links to headings with IDs back to themselves

## Contribute

See [`contributing.md`][contributing] in [`rehypejs/.github`][health] for ways
to get started.
See [`support.md`][support] for ways to get help.

This project has a [code of conduct][coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## License

[MIT][license] © [Titus Wormer][author]

<!-- Definitions -->

[build-badge]: https://github.com/rehypejs/rehype-slug/workflows/main/badge.svg

[build]: https://github.com/rehypejs/rehype-slug/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/rehypejs/rehype-slug.svg

[coverage]: https://codecov.io/github/rehypejs/rehype-slug

[downloads-badge]: https://img.shields.io/npm/dm/rehype-slug.svg

[downloads]: https://www.npmjs.com/package/rehype-slug

[size-badge]: https://img.shields.io/bundlejs/size/rehype-slug

[size]: https://bundlejs.com/?q=rehype-slug

[sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[backers-badge]: https://opencollective.com/unified/backers/badge.svg

[collective]: https://opencollective.com/unified

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[chat]: https://github.com/rehypejs/rehype/discussions

[npm]: https://docs.npmjs.com/cli/install

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[esmsh]: https://esm.sh

[health]: https://github.com/rehypejs/.github

[contributing]: https://github.com/rehypejs/.github/blob/main/contributing.md

[support]: https://github.com/rehypejs/.github/blob/main/support.md

[coc]: https://github.com/rehypejs/.github/blob/main/code-of-conduct.md

[license]: license

[author]: https://wooorm.com

[github-slugger]: https://github.com/Flet/github-slugger

[rehype]: https://github.com/rehypejs/rehype

[rehype-autolink-headings]: https://github.com/rehypejs/rehype-autolink-headings

[rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[rehype-sanitize-example]: https://github.com/rehypejs/rehype-sanitize#example-headings-dom-clobbering

[typescript]: https://www.typescriptlang.org

[unified]: https://github.com/unifiedjs/unified

[unified-transformer]: https://github.com/unifiedjs/unified#transformer

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[api-options]: #options

[api-rehype-slug]: #unifieduserehypeslug-options
