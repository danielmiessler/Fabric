# vfile-location

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]
[![Sponsors][sponsors-badge]][collective]
[![Backers][backers-badge]][collective]
[![Chat][chat-badge]][chat]

[vfile][] utility to convert between positional (line and column-based) and
offsets (range-based) locations.

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
  * [`Location`](#location)
  * [`location(file)`](#locationfile)
* [Types](#types)
* [Compatibility](#compatibility)
* [Contribute](#contribute)
* [License](#license)

## What is this?

This is a tiny but useful package to convert between arbitrary places in a
file.

## When should I use this?

This utility is useful when ASTs nodes don’t cut it.
For example, when you are making a lint rule that looks for dangerous
characters in a file, which you accomplish by searching the raw file value,
and still want to report it to users.

## Install

This package is [ESM only][esm].
In Node.js (version 16+), install with [npm][]:

```sh
npm install vfile-location
```

In Deno with [`esm.sh`][esmsh]:

```js
import {location} from 'https://esm.sh/vfile-location@5'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import {location} from 'https://esm.sh/vfile-location@5?bundle'
</script>
```

## Use

```js
import {VFile} from 'vfile'
import {location} from 'vfile-location'

const place = location(new VFile('foo\nbar\nbaz'))

const offset = place.toOffset({line: 3, column: 3}) // => 10
place.toPoint(offset) // => {line: 3, column: 3, offset: 10}
```

## API

### `Location`

Accessors for index.

###### Fields

* `toOffset` (`(point?: PointLike | null | undefined) => number | undefined`)
  — get the `offset` from a line/column based `Point` in the bound indices;
  returns `undefined` when given out of bounds input
* `toPoint` (`(offset?: number | null | undefined) => UnistPoint | undefined`)
  — get the line/column based `Point` for `offset` in the bound indices;
  returns `undefined` when given out of bounds input

### `location(file)`

Create an index of the given document to translate between line/column and
offset based positional info.

Also implemented in Rust in [`wooorm/markdown-rs`][markdown-rs].

[markdown-rs]: https://github.com/wooorm/markdown-rs/blob/main/src/util/location.rs

###### Parameters

* `file` (`VFile | Value`)
  — file to index

###### Returns

Accessors for index (`Location`).

## Types

This package is fully typed with [TypeScript][].
It exports the additional type [`Location`][api-location-map].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line, `vfile-location@^5`,
compatible with Node.js 16.

## Contribute

See [`contributing.md`][contributing] in [`vfile/.github`][health] for ways to
get started.
See [`support.md`][support] for ways to get help.

This project has a [code of conduct][coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## License

[MIT][license] © [Titus Wormer][author]

<!-- Definitions -->

[build-badge]: https://github.com/vfile/vfile-location/workflows/main/badge.svg

[build]: https://github.com/vfile/vfile-location/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/vfile/vfile-location.svg

[coverage]: https://codecov.io/github/vfile/vfile-location

[downloads-badge]: https://img.shields.io/npm/dm/vfile-location.svg

[downloads]: https://www.npmjs.com/package/vfile-location

[size-badge]: https://img.shields.io/badge/dynamic/json?label=minzipped%20size&query=$.size.compressedSize&url=https://deno.bundlejs.com/?q=vfile-location

[size]: https://bundlejs.com/?q=vfile-location

[sponsors-badge]: https://opencollective.com/unified/sponsors/badge.svg

[backers-badge]: https://opencollective.com/unified/backers/badge.svg

[collective]: https://opencollective.com/unified

[chat-badge]: https://img.shields.io/badge/chat-discussions-success.svg

[chat]: https://github.com/vfile/vfile/discussions

[npm]: https://docs.npmjs.com/cli/install

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[esmsh]: https://esm.sh

[typescript]: https://www.typescriptlang.org

[contributing]: https://github.com/vfile/.github/blob/main/contributing.md

[support]: https://github.com/vfile/.github/blob/main/support.md

[health]: https://github.com/vfile/.github

[coc]: https://github.com/vfile/.github/blob/main/code-of-conduct.md

[license]: license

[author]: https://wooorm.com

[vfile]: https://github.com/vfile/vfile

[api-location-map]: #location
