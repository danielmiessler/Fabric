# rehype-external-links

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

**[rehype][]** plugin to add `rel` (and `target`) to external links.

## Contents

*   [What is this?](#what-is-this)
*   [When should I use this?](#when-should-i-use-this)
*   [Install](#install)
*   [Use](#use)
*   [API](#api)
    *   [`unified().use(rehypeExternalLinks[, options])`](#unifieduserehypeexternallinks-options)
    *   [`CreateContent`](#createcontent)
    *   [`CreateProperties`](#createproperties)
    *   [`CreateRel`](#createrel)
    *   [`CreateTarget`](#createtarget)
    *   [`Options`](#options)
    *   [`Target`](#target)
*   [Types](#types)
*   [Compatibility](#compatibility)
*   [Security](#security)
*   [Contribute](#contribute)
*   [License](#license)

## What is this?

This package is a [unified][] ([rehype][]) plugin to add `rel` (and `target`)
attributes to external links.
It is particularly useful when displaying user content on your reputable site,
because users could link to disreputable sources (spam, scams, etc), as search
engines and other bots will discredit your site for linking to them (or
legitimize their sites).
In short: linking to something signals trust, but you can’t trust users.
This plugin adds certain `rel` attributes to prevent that from happening.

**unified** is a project that transforms content with abstract syntax trees
(ASTs).
**rehype** adds support for HTML to unified.
**hast** is the HTML AST that rehype uses.
This is a rehype plugin that adds `rel` (and `target`) to `<a>`s in the AST.

## When should I use this?

This project is useful when you want to display user content from authors you
don’t trust (such as comments), as they might include links you don’t endorse,
on your website.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install rehype-external-links
```

In Deno with [`esm.sh`][esmsh]:

```js
import rehypeExternalLinks from 'https://esm.sh/rehype-external-links@3'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import rehypeExternalLinks from 'https://esm.sh/rehype-external-links@3?bundle'
</script>
```

## Use

Say our module `example.js` contains:

```js
import rehypeExternalLinks from 'rehype-external-links'
import remarkParse from 'remark-parse'
import remarkRehype from 'remark-rehype'
import rehypeStringify from 'rehype-stringify'
import {unified} from 'unified'

const file = await unified()
  .use(remarkParse)
  .use(remarkRehype)
  .use(rehypeExternalLinks, {rel: ['nofollow']})
  .use(rehypeStringify)
  .process('[rehype](https://github.com/rehypejs/rehype)')

console.log(String(file))
```

…then running `node example.js` yields:

```html
<p><a href="https://github.com/rehypejs/rehype" rel="nofollow">rehype</a></p>
```

## API

This package exports no identifiers.
The default export is [`rehypeExternalLinks`][api-rehype-external-links].

### `unified().use(rehypeExternalLinks[, options])`

Automatically add `rel` (and `target`?) to external links.

###### Parameters

*   `options` ([`Options`][api-options], optional)
    — configuration

###### Returns

Transform ([`Transformer`][unified-transformer]).

###### Notes

You should [likely not configure `target`][css-tricks].

You should at least set `rel` to `['nofollow']`.
When using a `target`, add `noopener` and `noreferrer` to avoid exploitation
of the `window.opener` API.

When using a `target`, you should set `content` to adhere to accessibility
guidelines by [giving users advanced warning when opening a new window][g201].

### `CreateContent`

Create a target for the element (TypeScript type).

###### Parameters

*   `element` ([`Element`][hast-element])
    — element to check

###### Returns

Content to add (`Array<Node>` or `Node`, optional).

### `CreateProperties`

Create properties for an element (TypeScript type).

###### Parameters

*   `element` ([`Element`][hast-element])
    — element to check

###### Returns

Properties to add ([`Properties`][hast-properties], optional).

### `CreateRel`

Create a `rel` for the element (TypeScript type).

###### Parameters

*   `element` ([`Element`][hast-element])
    — element to check

###### Returns

`rel` to use (`Array<string>`, optional).

### `CreateTarget`

Create a `target` for the element (TypeScript type).

###### Parameters

*   `element` ([`Element`][hast-element])
    — element to check

###### Returns

`target` to use ([`Target`][api-target], optional).

### `Options`

Configuration (TypeScript type).

###### Fields

*   `content` (`Array<Node>`, [`CreateContent`][api-create-content], or `Node`,
    optional)
    — content to insert at the end of external links; will be inserted in a
    `<span>` element; useful for improving accessibility by giving users
    advanced warning when opening a new window
*   `contentProperties` ([`CreateProperties`][api-create-properties] or
    [`Properties`][hast-properties], optional)
    — properties to add to the `span` wrapping `content`
*   `properties` ([`CreateProperties`][api-create-properties] or
    [`Properties`][hast-properties], optional)
    — properties to add to the link itself
*   `protocols` (`Array<string>`, default: `['http', 'https']`)
    — protocols to see as external, such as `mailto` or `tel`
*   `rel` (`Array<string>`, [`CreateRel`][api-create-rel], or `string`,
    default: `['nofollow']`)
    — [link types][mdn-rel] to hint about the referenced documents; pass an
    empty array (`[]`) to not set `rel`s on links; when using a `target`, add `noopener`
    and `noreferrer` to avoid exploitation of the `window.opener` API
*   `target` ([`CreateTarget`][api-create-target] or [`Target`][api-target],
    optional)
    — how to display referenced documents; the default (nothing) is to not set
    `target`s on links
*   `test` ([`Test`][is-test], optional)
    — extra test to define which external link elements are modified; any test
    that can be given to `hast-util-is-element` is supported

### `Target`

Target (TypeScript type).

###### Type

```ts
type Target = '_blank' | '_parent' | '_self' | '_top'
```

## Types

This package is fully typed with [TypeScript][].
It exports the additional types
[`CreateContent`][api-create-content],
[`CreateProperties`][api-create-properties],
[`CreateRel`][api-create-rel],
[`CreateTarget`][api-create-target],
[`Options`][api-options], and
[`Target`][api-target].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line, `rehype-external-links@^3`,
compatible with Node.js 16.

This plugin works with `rehype-parse` version 3+, `rehype-stringify` version 3+,
`rehype` version 4+, and `unified` version 6+.

## Security

Improper use of `rehype-external-links` can open you up to a
[cross-site scripting (XSS)][xss] attack.

Either do not combine this plugin with user content or use
[`rehype-sanitize`][rehype-sanitize].

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

[build-badge]: https://github.com/rehypejs/rehype-external-links/workflows/main/badge.svg

[build]: https://github.com/rehypejs/rehype-external-links/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/rehypejs/rehype-external-links.svg

[coverage]: https://codecov.io/github/rehypejs/rehype-external-links

[downloads-badge]: https://img.shields.io/npm/dm/rehype-external-links.svg

[downloads]: https://www.npmjs.com/package/rehype-external-links

[size-badge]: https://img.shields.io/bundlejs/size/rehype-external-links

[size]: https://bundlejs.com/?q=rehype-external-links

[sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[backers-badge]: https://opencollective.com/unified/backers/badge.svg

[collective]: https://opencollective.com/unified

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[chat]: https://github.com/rehypejs/rehype/discussions

[npm]: https://docs.npmjs.com/cli/install

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[esmsh]: https://esm.sh

[health]: https://github.com/rehypejs/.github

[contributing]: https://github.com/rehypejs/.github/blob/HEAD/contributing.md

[support]: https://github.com/rehypejs/.github/blob/HEAD/support.md

[coc]: https://github.com/rehypejs/.github/blob/HEAD/code-of-conduct.md

[license]: license

[author]: https://wooorm.com

[hast-properties]: https://github.com/syntax-tree/hast#properties

[is-test]: https://github.com/syntax-tree/hast-util-is-element#test

[mdn-rel]: https://developer.mozilla.org/en-US/docs/Web/HTML/Link_types

[rehype]: https://github.com/rehypejs/rehype

[rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[typescript]: https://www.typescriptlang.org

[unified]: https://github.com/unifiedjs/unified

[unified-transformer]: https://github.com/unifiedjs/unified#transformer

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[hast-element]: https://github.com/syntax-tree/hast#element

[g201]: https://www.w3.org/WAI/WCAG21/Techniques/general/G201

[css-tricks]: https://css-tricks.com/use-target_blank/

[api-create-content]: #createcontent

[api-create-properties]: #createproperties

[api-create-rel]: #createrel

[api-create-target]: #createtarget

[api-options]: #options

[api-target]: #target

[api-rehype-external-links]: #unifieduserehypeexternallinks-options
