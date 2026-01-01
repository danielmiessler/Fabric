# rehype-autolink-headings

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

**[rehype][]** plugin to add links from headings back to themselves.

## Contents

*   [What is this?](#what-is-this)
*   [When should I use this?](#when-should-i-use-this)
*   [Install](#install)
*   [Use](#use)
*   [API](#api)
    *   [`unified().use(rehypeAutolinkHeadings[, options])`](#unifieduserehypeautolinkheadings-options)
    *   [`Behavior`](#behavior)
    *   [`Build`](#build)
    *   [`BuildProperties`](#buildproperties)
    *   [`Options`](#options)
*   [Examples](#examples)
    *   [Example: different behaviors](#example-different-behaviors)
    *   [Example: building content with `hastscript`](#example-building-content-with-hastscript)
    *   [Example: passing content from a string of HTML](#example-passing-content-from-a-string-of-html)
    *   [Example: group](#example-group)
*   [Types](#types)
*   [Compatibility](#compatibility)
*   [Security](#security)
*   [Related](#related)
*   [Contribute](#contribute)
*   [License](#license)

## What is this?

This package is a [unified][] ([rehype][]) plugin to add links from headings
back to themselves.
It looks for headings (so `<h1>` through `<h6>`) that have `id` properties,
and injects a link to themselves.
Similar functionality is applied by many places that render markdown.
For example, when browsing this readme on GitHub or npm, an anchor is added
to headings, which you can share to point people to a particular place in a
document.

**unified** is a project that transforms content with abstract syntax trees
(ASTs).
**rehype** adds support for HTML to unified.
**hast** is the HTML AST that rehype uses.
This is a rehype plugin that adds links to headings in the AST.

## When should I use this?

This plugin is useful when you have relatively long documents, where you want
users to be able to link to particular sections, and you already have `id`
properties set on all (or certain?) headings.

A different plugin, [`rehype-slug`][rehype-slug], adds `id`s to headings.
When a heading doesn’t already have an `id` property, it creates a slug from
it, and adds that as the `id` property.
When using both plugins together, all headings (whether explicitly with a
certain `id` or automatically with a generate one) will get a link back to
themselves.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install rehype-autolink-headings
```

In Deno with [`esm.sh`][esmsh]:

```js
import rehypeAutolinkHeadings from 'https://esm.sh/rehype-autolink-headings@7'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import rehypeAutolinkHeadings from 'https://esm.sh/rehype-autolink-headings@7?bundle'
</script>
```

## Use

Say we have the following file `example.html`:

```html
<h1>Solar System</h1>
<h2>Formation and evolution</h2>
<h2>Structure and composition</h2>
<h3>Orbits</h3>
<h3>Composition</h3>
<h3>Distances and scales</h3>
<h3>Interplanetary environment</h3>
<p>…</p>
```

…and our module `example.js` contains:

```js
import {rehype} from 'rehype'
import rehypeAutolinkHeadings from 'rehype-autolink-headings'
import rehypeSlug from 'rehype-slug'
import {read} from 'to-vfile'

const file = await rehype()
  .data('settings', {fragment: true})
  .use(rehypeSlug)
  .use(rehypeAutolinkHeadings)
  .process(await read('example.html'))

console.log(String(file))
```

…then running `node example.js` yields:

```html
<h1 id="solar-system"><a aria-hidden="true" tabindex="-1" href="#solar-system"><span class="icon icon-link"></span></a>Solar System</h1>
<h2 id="formation-and-evolution"><a aria-hidden="true" tabindex="-1" href="#formation-and-evolution"><span class="icon icon-link"></span></a>Formation and evolution</h2>
<h2 id="structure-and-composition"><a aria-hidden="true" tabindex="-1" href="#structure-and-composition"><span class="icon icon-link"></span></a>Structure and composition</h2>
<h3 id="orbits"><a aria-hidden="true" tabindex="-1" href="#orbits"><span class="icon icon-link"></span></a>Orbits</h3>
<h3 id="composition"><a aria-hidden="true" tabindex="-1" href="#composition"><span class="icon icon-link"></span></a>Composition</h3>
<h3 id="distances-and-scales"><a aria-hidden="true" tabindex="-1" href="#distances-and-scales"><span class="icon icon-link"></span></a>Distances and scales</h3>
<h3 id="interplanetary-environment"><a aria-hidden="true" tabindex="-1" href="#interplanetary-environment"><span class="icon icon-link"></span></a>Interplanetary environment</h3>
<p>…</p>
```

## API

This package exports no identifiers.
The default export is [`rehypeAutolinkHeadings`][api-rehype-autolink-headings].

### `unified().use(rehypeAutolinkHeadings[, options])`

Add links from headings back to themselves.

###### Parameters

*   `options` ([`Options`][api-options], optional)
    — configuration

###### Returns

Transform ([`Transformer`][unified-transformer]).

###### Notes

This plugin only applies to headings with `id`s.
Use `rehype-slug` to generate `id`s for headings that don’t have them.

Several behaviors are supported:

*   `'prepend'` (default) — inject link before the heading text
*   `'append'` — inject link after the heading text
*   `'wrap'` — wrap the whole heading text with the link
*   `'before'` — insert link before the heading
*   `'after'` — insert link after the heading

### `Behavior`

Behavior (TypeScript type).

###### Type

```ts
type Behavior = 'after' | 'append' | 'before' | 'prepend' | 'wrap'
```

### `Build`

Generate content (TypeScript type).

###### Parameters

*   `element` ([`Element`][hast-element])
    — current heading

###### Returns

Content ([`Array<Node>`][hast-node] or `Node`).

### `BuildProperties`

Generate properties (TypeScript type).

###### Parameters

*   `element` ([`Element`][hast-element])
    — current heading

###### Returns

Properties ([`Properties`][hast-properties]).

### `Options`

Configuration (TypeScript type).

###### Fields

*   `behavior` ([`Behavior`][api-behavior], default: `'prepend'`)
    — how to create links
*   `content` ([`Array<Node>`][hast-node], `Node`, or [`Build`][api-build],
    default: if `'wrap'` then `undefined`, otherwise equivalent of
    `<span class="icon icon-link"></span>`)
    — content to insert in the link;
    if `behavior` is `'wrap'` and `Build` is passed, its result replaces the
    existing content, otherwise the content is added after existing content
*   `group` ([`Array<Node>`][hast-node], `Node`, or [`Build`][api-build],
    optional)
    — content to wrap the heading and link with, if `behavior` is `'after'` or
    `'before'`
*   `headingProperties` ([`BuildProperties`][api-build-properties] or
    [`Properties`][hast-properties], optional)
    — extra properties to set on the heading
*   `properties` ([`BuildProperties`][api-build-properties] or
    [`Properties`][hast-properties], default:
    `{ariaHidden: true, tabIndex: -1}` if `'append'` or `'prepend'`, otherwise
    `undefined`)
    — extra properties to set on the link when injecting
*   `test` ([`Test`][hast-util-is-element-test], optional)
    — extra test for which headings are linked

## Examples

### Example: different behaviors

This example shows what each behavior generates by default.

```js
import {rehype} from 'rehype'
import rehypeAutolinkHeadings from 'rehype-autolink-headings'

const behaviors = ['after', 'append', 'before', 'prepend', 'wrap']
let index = -1
while (++index < behaviors.length) {
  const behavior = behaviors[index]
  console.log(
    String(
      await rehype()
        .data('settings', {fragment: true})
        .use(rehypeAutolinkHeadings, {behavior})
        .process('<h1 id="' + behavior + '">' + behavior + '</h1>')
    )
  )
}
```

Yields:

```html
<h1 id="after">after</h1><a href="#after"><span class="icon icon-link"></span></a>
<h1 id="append">append<a aria-hidden="true" tabindex="-1" href="#append"><span class="icon icon-link"></span></a></h1>
<a href="#before"><span class="icon icon-link"></span></a><h1 id="before">before</h1>
<h1 id="prepend"><a aria-hidden="true" tabindex="-1" href="#prepend"><span class="icon icon-link"></span></a>prepend</h1>
<h1 id="wrap"><a href="#wrap">wrap</a></h1>
```

### Example: building content with `hastscript`

The following example passes `options.content` as a function, to generate an
accessible description specific to each link.
It uses [`hastscript`][hastscript] to build nodes.

```js
import {h} from 'hastscript'
import {toString} from 'hast-util-to-string'
import {rehype} from 'rehype'
import rehypeAutolinkHeadings from 'rehype-autolink-headings'

const file = await rehype()
  .data('settings', {fragment: true})
  .use(rehypeAutolinkHeadings, {
    content(node) {
      return [
        h('span.visually-hidden', 'Read the “', toString(node), '” section'),
        h('span.icon.icon-link', {ariaHidden: 'true'})
      ]
    }
  })
  .process('<h1 id="pluto">Pluto</h1>')

console.log(String(file))
```

Yields:

```html
<h1 id="pluto"><a aria-hidden="true" tabindex="-1" href="#pluto"><span class="visually-hidden">Read the “Pluto” section</span><span class="icon icon-link" aria-hidden="true"></span></a>Pluto</h1>
```

### Example: passing content from a string of HTML

The following example passes `content` as nodes.
It uses [`hast-util-from-html-isomorphic`][hast-util-from-html-isomorphic] to
build nodes from a string of HTML.

```js
/**
 * @typedef {import('hast').ElementContent} ElementContent
 */

import {fromHtmlIsomorphic} from 'hast-util-from-html-isomorphic'
import {rehype} from 'rehype'
import rehypeAutolinkHeadings from 'rehype-autolink-headings'

const file = await rehype()
  .data('settings', {fragment: true})
  .use(rehypeAutolinkHeadings, {
    content: /** @type {Array<ElementContent>} */ (
      fromHtmlIsomorphic(
        '<svg height="10" width="10"><circle cx="5" cy="5" r="5" fill="black" /></svg>',
        {fragment: true}
      ).children
    )
  })
  .process('<h1 id="makemake">Makemake</h1>')

console.log(String(file))
```

Yields:

```html
<h1 id="makemake"><a aria-hidden="true" tabindex="-1" href="#makemake"><svg height="10" width="10"><circle cx="5" cy="5" r="5" fill="black"></circle></svg></a>Makemake</h1>
```

### Example: group

The following example passes `group` as a function, to dynamically generate a
differing element that wraps the heading.
It uses [`hastscript`][hastscript] to build nodes.

```js
import {h} from 'hastscript'
import {rehype} from 'rehype'
import rehypeAutolinkHeadings from 'rehype-autolink-headings'

const file = await rehype()
  .data('settings', {fragment: true})
  .use(rehypeAutolinkHeadings, {
    behavior: 'before',
    group(node) {
      return h('.heading-' + node.tagName.charAt(1) + '-group')
    }
  })
  .process('<h1 id="ceres">Ceres</h1>')

console.log(String(file))
```

Yields:

```html
<div class="heading-1-group"><a href="#ceres"><span class="icon icon-link"></span></a><h1 id="ceres">Ceres</h1></div>
```

## Types

This package is fully typed with [TypeScript][].
It exports the additional types
[`Behavior`][api-behavior],
[`Build`][api-build],
[`BuildProperties`][api-build-properties], and
[`Options`][api-options].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line,
`rehype-autolink-headings@^7`, compatible with Node.js 16.

This plugin works with `rehype-parse` version 1+, `rehype-stringify` version 1+,
`rehype` version 1+, and `unified` version 4+.

## Security

Use of `rehype-autolink-headings` can open you up to a
[cross-site scripting (XSS)][xss] attack if you pass user provided content in
`content`, `group`, or `properties`.

Always be wary of user input and use [`rehype-sanitize`][rehype-sanitize].

## Related

*   [`rehype-slug`][rehype-slug]
    — add `id`s to headings
*   [`rehype-highlight`](https://github.com/rehypejs/rehype-highlight)
    — apply syntax highlighting to code blocks
*   [`rehype-toc`](https://github.com/JS-DevTools/rehype-toc)
    — add a table of contents (TOC)

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

[build-badge]: https://github.com/rehypejs/rehype-autolink-headings/workflows/main/badge.svg

[build]: https://github.com/rehypejs/rehype-autolink-headings/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/rehypejs/rehype-autolink-headings.svg

[coverage]: https://codecov.io/github/rehypejs/rehype-autolink-headings

[downloads-badge]: https://img.shields.io/npm/dm/rehype-autolink-headings.svg

[downloads]: https://www.npmjs.com/package/rehype-autolink-headings

[size-badge]: https://img.shields.io/bundlejs/size/rehype-autolink-headings

[size]: https://bundlejs.com/?q=rehype-autolink-headings

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

[hast-element]: https://github.com/syntax-tree/hast#element

[hast-node]: https://github.com/syntax-tree/hast#nodes

[hast-util-is-element-test]: https://github.com/syntax-tree/hast-util-is-element#test

[hast-properties]: https://github.com/syntax-tree/hast#properties

[hastscript]: https://github.com/syntax-tree/hastscript

[hast-util-from-html-isomorphic]: https://github.com/syntax-tree/hast-util-from-html-isomorphic

[rehype]: https://github.com/rehypejs/rehype

[rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[typescript]: https://www.typescriptlang.org

[unified]: https://github.com/unifiedjs/unified

[unified-transformer]: https://github.com/unifiedjs/unified#transformer

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[rehype-slug]: https://github.com/rehypejs/rehype-slug

[api-behavior]: #behavior

[api-build]: #build

[api-build-properties]: #buildproperties

[api-options]: #options

[api-rehype-autolink-headings]: #unifieduserehypeautolinkheadings-options
