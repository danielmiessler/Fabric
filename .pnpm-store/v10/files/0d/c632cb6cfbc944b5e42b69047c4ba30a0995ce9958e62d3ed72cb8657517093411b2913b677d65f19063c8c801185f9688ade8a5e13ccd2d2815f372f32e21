# hast-util-heading-rank

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

[hast][] utility to get the rank (also known as depth or level) of headings.

## Contents

*   [What is this?](#what-is-this)
*   [When should I use this?](#when-should-i-use-this)
*   [Install](#install)
*   [Use](#use)
*   [API](#api)
    *   [`headingRank(node)`](#headingranknode)
*   [Types](#types)
*   [Compatibility](#compatibility)
*   [Security](#security)
*   [Related](#related)
*   [Contribute](#contribute)
*   [License](#license)

## What is this?

This package is a utility that lets you get the rank (`1`, `6`) of heading
elements (`h1`, `h6`).

## When should I use this?

This utility is pretty niche, if you’re here you probably know what you’re
looking for!

To change heading ranks, use
[`hast-util-shift-heading`][hast-util-shift-heading].

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install hast-util-heading-rank
```

In Deno with [`esm.sh`][esmsh]:

```js
import {headingRank} from 'https://esm.sh/hast-util-heading-rank@3'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import {headingRank} from 'https://esm.sh/hast-util-heading-rank@3?bundle'
</script>
```

## Use

```js
import {h} from 'hastscript'
import {headingRank} from 'hast-util-heading-rank'

headingRank(h('p', 'Alpha')) //=> undefined
headingRank(h('h5', 'Alpha')) //=> 5
```

## API

This package exports the identifier [`headingRank`][api-heading-rank].
There is no default export.

### `headingRank(node)`

Get the rank (`1` to `6`) of headings (`h1` to `h6`).

###### Parameters

*   `node` ([`Node`][node])
    — node to check

###### Returns

Rank of the heading (`number`) or `undefined` if not a heading.

## Types

This package is fully typed with [TypeScript][].
It exports no additional types.

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line,
`hast-util-heading-rank@^3`, compatible with Node.js 16.

## Security

`hast-util-heading-rank` does not mutate.
There are no openings for [cross-site scripting (XSS)][xss] attacks.

## Related

*   [`hast-util-heading`](https://github.com/syntax-tree/hast-util-heading)
    — check if a node is heading content
*   [`hast-util-shift-heading`](https://github.com/syntax-tree/hast-util-heading)
    — change heading rank

## Contribute

See [`contributing.md`][contributing] in [`syntax-tree/.github`][health] for
ways to get started.
See [`support.md`][support] for ways to get help.

This project has a [code of conduct][coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## License

[MIT][license] © [Titus Wormer][author]

<!-- Definition -->

[build-badge]: https://github.com/syntax-tree/hast-util-heading-rank/workflows/main/badge.svg

[build]: https://github.com/syntax-tree/hast-util-heading-rank/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/syntax-tree/hast-util-heading-rank.svg

[coverage]: https://codecov.io/github/syntax-tree/hast-util-heading-rank

[downloads-badge]: https://img.shields.io/npm/dm/hast-util-heading-rank.svg

[downloads]: https://www.npmjs.com/package/hast-util-heading-rank

[size-badge]: https://img.shields.io/badge/dynamic/json?label=minzipped%20size&query=$.size.compressedSize&url=https://deno.bundlejs.com/?q=hast-util-heading-rank

[size]: https://bundlejs.com/?q=hast-util-heading-rank

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

[node]: https://github.com/syntax-tree/hast#nodes

[xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[hast-util-shift-heading]: https://github.com/syntax-tree/hast-util-shift-heading

[api-heading-rank]: #headingranknode
