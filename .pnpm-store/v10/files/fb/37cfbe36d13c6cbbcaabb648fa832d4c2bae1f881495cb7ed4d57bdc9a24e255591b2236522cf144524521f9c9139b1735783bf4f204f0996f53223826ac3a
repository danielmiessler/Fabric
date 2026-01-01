<!--This file is generated-->

# hast-util-to-string

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][funding-sponsors-badge]][funding]
[![Backers][funding-backers-badge]][funding]
[![Chat][chat-badge]][chat]

[`hast`][hast] utility to get the plain-text value of a node.

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
  * [`toString(node)`](#tostringnode)
* [Syntax](#syntax)
* [Syntax tree](#syntax-tree)
* [Types](#types)
* [Compatibility](#compatibility)
* [Security](#security)
* [Contribute](#contribute)
* [License](#license)

## What is this?

This package is a utility to get the plain-text value of a node.

## When should I use this?

You can use this package when you want to get the plain text value of a node.
The algorithm used by this package is like the DOMs `Node#textContent`
getter.

To use the DOMs `Node#innerText` algorithm instead, use
[`hast-util-to-text`](https://github.com/syntax-tree/hast-util-to-text).
`innerText` is aware of how things are displayed, for example turning hard
breaks (`<br>`) into line endings.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install hast-util-to-string
```

In Deno with [`esm.sh`][esm-sh]:

```js
import {toString} from 'https://esm.sh/hast-util-to-string@3'
```

In browsers with [`esm.sh`][esm-sh]:

```html
<script type="module">
  import {toString} from 'https://esm.sh/hast-util-to-string@3?bundle'
</script>
```

## Use

```js
import {h} from 'hastscript'
import {toString} from 'hast-util-to-string'

toString(h('p', 'Alpha'))
//=> 'Alpha'
toString(h('div', [h('b', 'Bold'), ' and ', h('i', 'italic'), '.']))
//=> 'Bold and italic.'
```

## API

This package exports the identifier
`toString`.
There is no default export.

### `toString(node)`

Get the plain-text value of a node.

###### Parameters

* `node` (`Node`) — node to serialize

###### Returns

Serialized node (`string`).

## Syntax

HTML is parsed according to WHATWG HTML (the living standard), which is also
followed by all browsers.

## Syntax tree

The syntax tree used is [hast][].

## Types

This package is fully typed with [TypeScript][].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line,
`hast-util-to-string@^3`,
compatible with Node.js 16.

## Security

As **rehype** works on HTML and improper use of HTML can open you up to a
[cross-site scripting (XSS)][xss] attack, use of rehype can also be unsafe.
Use [`rehype-sanitize`][rehype-sanitize] to make the tree safe.

## Contribute

See [`contributing.md`][contributing] in [`rehypejs/.github`][health] for ways
to get started.
See [`support.md`][support] for ways to get help.

This project has a [code of conduct][coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## License

[MIT][license] © [Titus Wormer][author]

[author]: https://wooorm.com

[build]: https://github.com/rehypejs/rehype-minify/actions

[build-badge]: https://github.com/rehypejs/rehype-minify/workflows/main/badge.svg

[chat]: https://github.com/rehypejs/rehype/discussions

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[coc]: https://github.com/rehypejs/.github/blob/main/code-of-conduct.md

[contributing]: https://github.com/rehypejs/.github/blob/main/contributing.md

[coverage]: https://codecov.io/github/rehypejs/rehype-minify

[coverage-badge]: https://img.shields.io/codecov/c/github/rehypejs/rehype-minify.svg

[downloads]: https://www.npmjs.com/package/hast-util-to-string

[downloads-badge]: https://img.shields.io/npm/dm/hast-util-to-string.svg

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[esm-sh]: https://esm.sh

[funding]: https://opencollective.com/unified

[funding-backers-badge]: https://opencollective.com/unified/backers/badge.svg

[funding-sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[hast]: https://github.com/syntax-tree/hast

[health]: https://github.com/rehypejs/.github

[license]: https://github.com/rehypejs/rehype-minify/blob/main/license

[npm]: https://docs.npmjs.com/cli/install

[rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[size]: https://bundlejs.com/?q=hast-util-to-string

[size-badge]: https://img.shields.io/bundlejs/size/hast-util-to-string

[support]: https://github.com/rehypejs/.github/blob/main/support.md

[typescript]: https://www.typescriptlang.org

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting
