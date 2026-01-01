# hast-util-parse-selector

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

[hast][] utility to create an element from a simple CSS selector.

## Contents

*   [What is this?](#what-is-this)
*   [When should I use this?](#when-should-i-use-this)
*   [Install](#install)
*   [Use](#use)
*   [API](#api)
    *   [`parseSelector(selector?[, defaultTagName])`](#parseselectorselector-defaulttagname)
*   [Types](#types)
*   [Compatibility](#compatibility)
*   [Security](#security)
*   [Related](#related)
*   [Contribute](#contribute)
*   [License](#license)

## What is this?

This package is a tiny utility that helps create elements.

## When should I use this?

This utility is super niche.
You probably want the more powerful [`hastscript`][hastscript] or
[`hast-util-from-selector`][hast-util-from-selector]

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install hast-util-parse-selector
```

In Deno with [`esm.sh`][esmsh]:

```js
import {parseSelector} from 'https://esm.sh/hast-util-parse-selector@4'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import {parseSelector} from 'https://esm.sh/hast-util-parse-selector@4?bundle'
</script>
```

## Use

```js
import {parseSelector} from 'hast-util-parse-selector'

console.log(parseSelector('.quux#bar.baz.qux'))
```

Yields:

```js
{ type: 'element',
  tagName: 'div',
  properties: { id: 'bar', className: [ 'quux', 'baz', 'qux' ] },
  children: [] }
```

## API

This package exports the identifier [`parseSelector`][api-parse-selector].
There is no default export.

### `parseSelector(selector?[, defaultTagName])`

Create a hast element from a simple CSS selector.

###### Parameters

*   `selector` (`string`, optional)
    —  simple CSS selector, can contain a tag name (`foo`), classes (`.bar`),
    and an ID (`#baz`), multiple classes are allowed, uses the last ID if
    multiple IDs are found
*   `defaultTagName` (`string`, default: `'div'`)
    — tag name to use if `selector` does not specify one

###### Returns

Built element ([`Element`][element]).

## Types

This package is fully typed with [TypeScript][].
It exports no additional types.

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line,
`hast-util-parse-selector@^4`, compatible with Node.js 16.

## Security

Improper use of the `selector` or `defaultTagName` can open you up to a
[cross-site scripting (XSS)][xss] attack as the value of `tagName`, when
resolving to `script`, injects a `script` element into the syntax tree.

Do not use user input in `selector` or use
[`hast-util-santize`][hast-util-sanitize].

## Related

*   [`hast-util-from-selector`](https://github.com/syntax-tree/hast-util-from-selector)
    — parse complex CSS selectors to nodes

## Contribute

See [`contributing.md`][contributing] in [`syntax-tree/.github`][health] for
ways to get started.
See [`support.md`][support] for ways to get help.

This project has a [code of conduct][coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## License

[MIT][license] © [Titus Wormer][author]

<!-- Definitions -->

[build-badge]: https://github.com/syntax-tree/hast-util-parse-selector/workflows/main/badge.svg

[build]: https://github.com/syntax-tree/hast-util-parse-selector/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/syntax-tree/hast-util-parse-selector.svg

[coverage]: https://codecov.io/github/syntax-tree/hast-util-parse-selector

[downloads-badge]: https://img.shields.io/npm/dm/hast-util-parse-selector.svg

[downloads]: https://www.npmjs.com/package/hast-util-parse-selector

[size-badge]: https://img.shields.io/badge/dynamic/json?label=minzipped%20size&query=$.size.compressedSize&url=https://deno.bundlejs.com/?q=hast-util-parse-selector

[size]: https://bundlejs.com/?q=hast-util-parse-selector

[sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[backers-badge]: https://opencollective.com/unified/backers/badge.svg

[collective]: https://opencollective.com/unified

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[chat]: https://github.com/syntax-tree/unist/discussions

[npm]: https://docs.npmjs.com/cli/install

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[esmsh]: https://esm.sh

[typescript]: https://www.typescriptlang.org

[license]: license

[author]: https://wooorm.com

[health]: https://github.com/syntax-tree/.github

[contributing]: https://github.com/syntax-tree/.github/blob/main/contributing.md

[support]: https://github.com/syntax-tree/.github/blob/main/support.md

[coc]: https://github.com/syntax-tree/.github/blob/main/code-of-conduct.md

[hast]: https://github.com/syntax-tree/hast

[hast-util-sanitize]: https://github.com/syntax-tree/hast-util-sanitize

[hastscript]: https://github.com/syntax-tree/hastscript

[hast-util-from-selector]: https://github.com/syntax-tree/hast-util-from-selector

[element]: https://github.com/syntax-tree/hast#element

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[api-parse-selector]: #parseselectorselector-defaulttagname
