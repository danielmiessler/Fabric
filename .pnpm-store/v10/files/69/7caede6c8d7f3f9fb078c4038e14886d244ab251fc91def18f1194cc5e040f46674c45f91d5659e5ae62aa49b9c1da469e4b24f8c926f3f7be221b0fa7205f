# bail

[![Build][build-badge]][build]
[![Coverage][coverage-badge]][coverage]
[![Downloads][downloads-badge]][downloads]
[![Size][size-badge]][size]

Throw if given an error.

## Contents

*   [What is this?](#what-is-this)
*   [When should I use this?](#when-should-i-use-this)
*   [Install](#install)
*   [Use](#use)
*   [API](#api)
    *   [`bail(err?)`](#bailerr)
*   [Types](#types)
*   [Compatibility](#compatibility)
*   [Security](#security)
*   [Related](#related)
*   [Contribute](#contribute)
*   [License](#license)

## What is this?

This package throws a given error.

## When should I use this?

Use this package if you’re building some scripts that might theoretically get
errors but frequently don’t and you keep writing `if (error) throw error` over
and over again and you’re just really done with that.

## Install

This package is [ESM only][esm].
In Node.js (version 12.20+, 14.14+, or 16.0+), install with [npm][]:

```sh
npm install bail
```

In Deno with [Skypack][]:

```js
import {bail} from 'https://cdn.skypack.dev/bail@2?dts'
```

In browsers with [Skypack][]:

```html
<script type="module">
  import {bail} from 'https://cdn.skypack.dev/bail@2?min'
</script>
```

## Use

```js
import {bail} from 'bail'

bail()

bail(new Error('failure'))
// Error: failure
//     at repl:1:6
//     at REPLServer.defaultEval (repl.js:154:27)
//     …
```

## API

This package exports the following identifier: `bail`.
There is no default export.

### `bail(err?)`

Throw a given error (`Error?`).

## Types

This package is fully typed with [TypeScript][].
There are no extra exported types.

## Compatibility

This package is at least compatible with all maintained versions of Node.js.
As of now, that is Node.js 12.20+, 14.14+, and 16.0+.
It also works in Deno and modern browsers.

## Security

This package is safe.

## Related

*   [`noop`][noop]
*   [`noop2`][noop2]
*   [`noop3`][noop3]

## Contribute

Yes please!
See [How to Contribute to Open Source][contribute].

## License

[MIT][license] © [Titus Wormer][author]

<!-- Definitions -->

[build-badge]: https://github.com/wooorm/bail/workflows/main/badge.svg

[build]: https://github.com/wooorm/bail/actions

[coverage-badge]: https://img.shields.io/codecov/c/github/wooorm/bail.svg

[coverage]: https://codecov.io/github/wooorm/bail

[downloads-badge]: https://img.shields.io/npm/dm/bail.svg

[downloads]: https://www.npmjs.com/package/bail

[size-badge]: https://img.shields.io/bundlephobia/minzip/bail.svg

[size]: https://bundlephobia.com/result?p=bail

[npm]: https://docs.npmjs.com/cli/install

[skypack]: https://www.skypack.dev

[license]: license

[author]: https://wooorm.com

[esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[typescript]: https://www.typescriptlang.org

[contribute]: https://opensource.guide/how-to-contribute/

[noop]: https://www.npmjs.com/package/noop

[noop2]: https://www.npmjs.com/package/noop2

[noop3]: https://www.npmjs.com/package/noop3
