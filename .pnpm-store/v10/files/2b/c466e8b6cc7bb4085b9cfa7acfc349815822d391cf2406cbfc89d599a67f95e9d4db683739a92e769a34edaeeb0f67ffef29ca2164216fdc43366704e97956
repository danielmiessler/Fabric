# hast-util-from-parse5

[![Build][badge-build-image]][badge-build-url]
[![Coverage][badge-coverage-image]][badge-coverage-url]
[![Downloads][badge-downloads-image]][badge-downloads-url]
[![Size][badge-size-image]][badge-size-url]

[hast][github-hast] utility to transform from the
[`parse5`][github-parse5] AST.

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
  * [`fromParse5(tree[, options])`](#fromparse5tree-options)
  * [`Options`](#options)
  * [`Space`](#space-1)
* [Types](#types)
* [Compatibility](#compatibility)
* [Security](#security)
* [Related](#related)
* [Contribute](#contribute)
* [License](#license)

## What is this?

This package is a utility that can turn a parse5 tree into a hast tree.

## When should I use this?

You can use this package when using `parse5` as an HTML parser and wanting to
work with hast.

The utility [`hast-util-to-parse5`][github-hast-util-to-parse5] does the
inverse of this utility.
It generates `parse5`s AST again.

The utility [`hast-util-from-html`][github-hast-util-from-html] wraps this
utility and `parse5` to both parse HTML and generate hast from it.

## Install

This package is [ESM only][github-gist-esm].
In Node.js (version 16+),
install with [npm][npmjs-install]:

```sh
npm install hast-util-from-parse5
```

In Deno with [`esm.sh`][esmsh]:

```js
import {fromParse5} from "https://esm.sh/hast-util-from-parse5@8"
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import {fromParse5} from "https://esm.sh/hast-util-from-parse5@8?bundle"
</script>
```

## Use

Say our document `example.html` contains:

```html
<!doctype html><title>Hello!</title><h1 id="world">World!<!--after-->
```

…and our module `example.js` looks as follows:

```js
import {fromParse5} from 'hast-util-from-parse5'
import {parse} from 'parse5'
import {read} from 'to-vfile'
import {inspect} from 'unist-util-inspect'

const file = await read('example.html')
const p5ast = parse(String(file), {sourceCodeLocationInfo: true})
const hast = fromParse5(p5ast, {file})

console.log(inspect(hast))
```

…now running `node example.js` yields:

```text
root[2] (1:1-2:1, 0-70)
│ data: {"quirksMode":false}
├─0 doctype (1:1-1:16, 0-15)
└─1 element<html>[2]
    │ properties: {}
    ├─0 element<head>[1]
    │   │ properties: {}
    │   └─0 element<title>[1] (1:16-1:37, 15-36)
    │       │ properties: {}
    │       └─0 text "Hello!" (1:23-1:29, 22-28)
    └─1 element<body>[1]
        │ properties: {}
        └─0 element<h1>[3] (1:37-2:1, 36-70)
            │ properties: {"id":"world"}
            ├─0 text "World!" (1:52-1:58, 51-57)
            ├─1 comment "after" (1:58-1:70, 57-69)
            └─2 text "\n" (1:70-2:1, 69-70)
```

## API

This package exports the identifier [`fromParse5`][api-from-parse5].
There is no default export.

### `fromParse5(tree[, options])`

Transform a `parse5` AST to hast.

###### Parameters

* `tree`
  ([`Parse5Node`][github-parse5-node])
  — `parse5` tree to transform
* `options`
  ([`Options`][api-options], optional)
  — configuration

###### Returns

hast tree ([`HastNode`][github-hast-nodes]).

### `Options`

Configuration (TypeScript type).

##### Fields

###### `file`

File used to add positional info to nodes
([`VFile`][github-vfile], optional).

If given,
the file should represent the original HTML source.

###### `space`

Which space the document is in
([`Space`][api-space], default: `'html'`).

When an `<svg>` element is found in the HTML space,
this package already automatically switches to and from the SVG space when
entering and exiting it.

###### `verbose`

Whether to add extra positional info about starting tags,
closing tags,
and attributes to elements
(`boolean`, default: `false`).

> 👉 **Note**:
> only used when `file` is given.

For the following HTML:

```html
<img src="http://example.com/fav.ico" alt="foo" title="bar">
```

The verbose info would looks as follows:

```js
{
  type: 'element',
  tagName: 'img',
  properties: {src: 'http://example.com/fav.ico', alt: 'foo', title: 'bar'},
  children: [],
  data: {
    position: {
      opening: {
        start: {line: 1, column: 1, offset: 0},
        end: {line: 1, column: 61, offset: 60}
      },
      closing: null,
      properties: {
        src: {
          start: {line: 1, column: 6, offset: 5},
          end: {line: 1, column: 38, offset: 37}
        },
        alt: {
          start: {line: 1, column: 39, offset: 38},
          end: {line: 1, column: 48, offset: 47}
        },
        title: {
          start: {line: 1, column: 49, offset: 48},
          end: {line: 1, column: 60, offset: 59}
        }
      }
    }
  },
  position: {
    start: {line: 1, column: 1, offset: 0},
    end: {line: 1, column: 61, offset: 60}
  }
}
```

### `Space`

Namespace (TypeScript type).

###### Type

```ts
type Space = 'html' | 'svg'
```

## Types

This package is fully typed with [TypeScript][].
It exports the additional types [`Options`][api-options] and
[`Space`][api-space].

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release,
we drop support for unmaintained versions of Node.
This means we try to keep the current release line,
`hast-util-from-parse5@8`,
compatible with Node.js 16.

## Security

Use of `hast-util-from-parse5` can open you up to a
[cross-site scripting (XSS)][wikipedia-xss] attack if Parse5’s AST is unsafe.

## Related

* [`hast-util-to-parse5`][github-hast-util-to-parse5]
  — transform hast to Parse5’s AST
* [`hast-util-to-nlcst`](https://github.com/syntax-tree/hast-util-to-nlcst)
  — transform hast to nlcst
* [`hast-util-to-mdast`](https://github.com/syntax-tree/hast-util-to-mdast)
  — transform hast to mdast
* [`hast-util-to-xast`](https://github.com/syntax-tree/hast-util-to-xast)
  — transform hast to xast
* [`mdast-util-to-hast`](https://github.com/syntax-tree/mdast-util-to-hast)
  — transform mdast to hast
* [`mdast-util-to-nlcst`](https://github.com/syntax-tree/mdast-util-to-nlcst)
  — transform mdast to nlcst

## Contribute

See [`contributing.md`][health-contributing]
in
[`syntax-tree/.github`][health]
for ways to get started.
See [`support.md`][health-support] for ways to get help.

This project has a [code of conduct][health-coc].
By interacting with this repository,
organization,
or community you agree to abide by its terms.

## License

[MIT][file-license] © [Titus Wormer][wooorm]

<!-- Definitions -->

[api-from-parse5]: #fromparse5tree-options

[api-options]: #options

[api-space]: #space-1

[badge-build-image]: https://github.com/syntax-tree/hast-util-from-parse5/workflows/main/badge.svg

[badge-build-url]: https://github.com/syntax-tree/hast-util-from-parse5/actions

[badge-coverage-image]: https://img.shields.io/codecov/c/github/syntax-tree/hast-util-from-parse5.svg

[badge-coverage-url]: https://codecov.io/github/syntax-tree/hast-util-from-parse5

[badge-downloads-image]: https://img.shields.io/npm/dm/hast-util-from-parse5.svg

[badge-downloads-url]: https://www.npmjs.com/package/hast-util-from-parse5

[badge-size-image]: https://img.shields.io/bundlejs/size/hast-util-from-parse5

[badge-size-url]: https://bundlejs.com/?q=hast-util-from-parse5

[esmsh]: https://esm.sh

[file-license]: license

[github-gist-esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[github-hast]: https://github.com/syntax-tree/hast

[github-hast-nodes]: https://github.com/syntax-tree/hast#nodes

[github-hast-util-from-html]: https://github.com/syntax-tree/hast-util-from-html

[github-hast-util-to-parse5]: https://github.com/syntax-tree/hast-util-to-parse5

[github-parse5]: https://github.com/inikulin/parse5

[github-parse5-node]: https://github.com/inikulin/parse5/blob/master/packages/parse5/lib/tree-adapters/default.ts

[github-vfile]: https://github.com/vfile/vfile

[health]: https://github.com/syntax-tree/.github

[health-coc]: https://github.com/syntax-tree/.github/blob/main/code-of-conduct.md

[health-contributing]: https://github.com/syntax-tree/.github/blob/main/contributing.md

[health-support]: https://github.com/syntax-tree/.github/blob/main/support.md

[npmjs-install]: https://docs.npmjs.com/cli/install

[typescript]: https://www.typescriptlang.org

[wikipedia-xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[wooorm]: https://wooorm.com
