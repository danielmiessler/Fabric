# hastscript

[![Build][badge-build-image]][badge-build-url]
[![Coverage][badge-coverage-image]][badge-coverage-url]
[![Downloads][badge-downloads-image]][badge-downloads-url]
[![Size][badge-size-image]][badge-size-url]

[hast][github-hast] utility to create trees with ease.

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
  * [`h(selector?[, properties][, …children])`](#hselector-properties-children)
  * [`s(selector?[, properties][, …children])`](#sselector-properties-children)
  * [`Child`](#child)
  * [`Properties`](#properties-1)
  * [`Result`](#result)
* [Syntax tree](#syntax-tree)
* [JSX](#jsx)
* [Compatibility](#compatibility)
* [Security](#security)
* [Related](#related)
* [Contribute](#contribute)
* [License](#license)

## What is this?

This package is a hyperscript interface (like `createElement` from React and
`h` from Vue and such) to help with creating hast trees.

## When should I use this?

You can use this utility in your project when you generate hast syntax trees
with code.
It helps because it replaces most of the repetition otherwise needed in a syntax
tree with function calls.
It also helps as it improves the attributes you pass by turning them into the
form that is required by hast.

You can instead use [`unist-builder`][github-unist-builder]
when creating any unist nodes and
[`xastscript`][github-xastscript] when creating xast (XML) nodes.

## Install

This package is [ESM only][github-gist-esm].
In Node.js (version 16+),
install with [npm][npmjs-install]:

```sh
npm install hastscript
```

In Deno with [`esm.sh`][esmsh]:

```js
import {h} from 'https://esm.sh/hastscript@9'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import {h} from 'https://esm.sh/hastscript@9?bundle'
</script>
```

## Use

```js
import {h, s} from 'hastscript'

console.log(
  h('.foo#some-id', [
    h('span', 'some text'),
    h('input', {type: 'text', value: 'foo'}),
    h('a.alpha', {class: 'bravo charlie', download: 'download'}, [
      'delta',
      'echo'
    ])
  ])
)

console.log(
  s('svg', {viewbox: '0 0 500 500', xmlns: 'http://www.w3.org/2000/svg'}, [
    s('title', 'SVG `<circle>` element'),
    s('circle', {cx: 120, cy: 120, r: 100})
  ])
)
```

Yields:

```js
{
  type: 'element',
  tagName: 'div',
  properties: {className: ['foo'], id: 'some-id'},
  children: [
    {
      type: 'element',
      tagName: 'span',
      properties: {},
      children: [{type: 'text', value: 'some text'}]
    },
    {
      type: 'element',
      tagName: 'input',
      properties: {type: 'text', value: 'foo'},
      children: []
    },
    {
      type: 'element',
      tagName: 'a',
      properties: {className: ['alpha', 'bravo', 'charlie'], download: true},
      children: [{type: 'text', value: 'delta'}, {type: 'text', value: 'echo'}]
    }
  ]
}
{
  type: 'element',
  tagName: 'svg',
  properties: {viewBox: '0 0 500 500', xmlns: 'http://www.w3.org/2000/svg'},
  children: [
    {
      type: 'element',
      tagName: 'title',
      properties: {},
      children: [{type: 'text', value: 'SVG `<circle>` element'}]
    },
    {
      type: 'element',
      tagName: 'circle',
      properties: {cx: 120, cy: 120, r: 100},
      children: []
    }
  ]
}
```

## API

This package exports the identifiers [`h`][api-h] and [`s`][api-s].
There is no default export.
It exports the additional [TypeScript][] types
[`Child`][api-child],
[`Properties`][api-properties],
and
[`Result`][api-result].

The export map supports the automatic JSX runtime.
You can pass `hastscript` or `hastscript/svg` to your build tool
(TypeScript, Babel, SWC)
with an `importSource` option or similar.

### `h(selector?[, properties][, …children])`

Create virtual **[hast][github-hast]** trees for HTML.

##### Signatures

* `h(): root`
* `h(null[, …children]): root`
* `h(selector[, properties][, …children]): element`

##### Parameters

###### `selector`

Simple CSS selector
(`string`, optional).
When string, builds an [`Element`][github-hast-element].
When nullish, builds a [`Root`][github-hast-root] instead.
The selector can contain a tag name (`foo`),
IDs (`#bar`),
and classes (`.baz`).
If the selector is a string but there is no tag name in it then `h` defaults to
build a `div` element and `s` to a `g` element.
`selector` is parsed by
[`hast-util-parse-selector`][github-hast-util-parse-selector].

###### `properties`

Properties of the element
([`Properties`][api-properties], optional).

###### `children`

Children of the node ([`Child`][api-child] or `Array<Child>`, optional).

##### Returns

Created tree ([`Result`][api-result]).

[`Element`][github-hast-element] when a `selector` is passed,
otherwise [`Root`][github-hast-root].

### `s(selector?[, properties][, …children])`

Create virtual **[hast][github-hast]** trees for SVG.

Signatures, parameters, and return value are the same as `h` above.
Importantly,
the `selector` and `properties` parameters are interpreted as SVG.

### `Child`

(Lists of) children (TypeScript type).

When strings or numbers are encountered,
they are turned into [`Text`][github-hast-text]
nodes.
[`Root`][github-hast-root] nodes are treated as “fragments”,
meaning that their children are used instead.

###### Type

```ts
type Child =
  | Array<Node | number | string | null | undefined>
  | Node
  | number
  | string
  | null
  | undefined
```

### `Properties`

Map of properties (TypeScript type).
Keys should match either the HTML attribute name or the DOM property name,
but are case-insensitive.

###### Type

```ts
type Properties = Record<
  string,
  | boolean
  | number
  | string
  | null
  | undefined
  // For comma- and space-separated values such as `className`:
  | Array<number | string>
  // Accepts value for `style` prop as object.
  | Record<string, number | string>
>
```

### `Result`

Result from a `h` (or `s`) call (TypeScript type).

###### Type

```ts
type Result = Element | Root
```

## Syntax tree

The syntax tree is [hast][github-hast].

## JSX

This package can be used with JSX.
You should use the automatic JSX runtime set to `hastscript` or
`hastscript/svg`.

> 👉 **Note**
> while `h` supports dots (`.`) for classes or number signs (`#`)
> for IDs in `selector`,
> those are not supported in JSX.

> 🪦 **Legacy**:
> you can also use the classic JSX runtime,
> but this is not recommended.
> To do so,
> import `h` (or `s`) yourself and define it as the pragma
> (plus set the fragment to `null`).

The Use example above can then be written like so,
using inline pragmas,
so that SVG can be used too:

`example-html.jsx`:

```js
/** @jsxImportSource hastscript */
console.log(
  <div class="foo" id="some-id">
    <span>some text</span>
    <input type="text" value="foo" />
    <a class="alpha bravo charlie" download>
      deltaecho
    </a>
  </div>
)
```

`example-svg.jsx`:

```js
/** @jsxImportSource hastscript/svg */
console.log(
  <svg xmlns="http://www.w3.org/2000/svg" viewbox="0 0 500 500">
    <title>SVG `&lt;circle&gt;` element</title>
    <circle cx={120} cy={120} r={100} />
  </svg>
)
```

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release,
we drop support for unmaintained versions of Node.
This means we try to keep the current release line,
`hastscript@9`,
compatible with Node.js 16.

## Security

Use of `hastscript` can open you up to a
[cross-site scripting (XSS)][wikipedia-xss]
when you pass user-provided input to it because values are injected into the
syntax tree.

The following example shows how an image is injected that fails loading and
therefore runs code in a browser.

```js
const tree = h()

// Somehow someone injected these properties instead of an expected `src` and
// `alt`:
const otherProps = {onError: 'alert(1)', src: 'x'}

tree.children.push(h('img', {src: 'default.png', ...otherProps}))
```

Yields:

```html
<img onerror="alert(1)" src="x">
```

The following example shows how code can run in a browser because someone stored
an object in a database instead of the expected string.

```js
const tree = h()

// Somehow this isn’t the expected `'wooorm'`.
const username = {
  type: 'element',
  tagName: 'script',
  children: [{type: 'text', value: 'alert(2)'}]
}

tree.children.push(h('span.handle', username))
```

Yields:

```html
<span class="handle"><script>alert(2)</script></span>
```

Either do not use user-provided input in `hastscript` or use
[`hast-util-santize`][github-hast-util-sanitize].

## Related

* [`unist-builder`][github-unist-builder]
  — create unist trees
* [`xastscript`][github-xastscript]
  — create xast trees
* [`hast-to-hyperscript`](https://github.com/syntax-tree/hast-to-hyperscript)
  — turn hast into React, Preact, Vue, etc
* [`hast-util-to-html`](https://github.com/syntax-tree/hast-util-to-html)
  — turn hast into HTML
* [`hast-util-to-dom`](https://github.com/syntax-tree/hast-util-to-dom)
  — turn hast into DOM trees
* [`estree-util-build-jsx`](https://github.com/syntax-tree/estree-util-build-jsx)
  — compile JSX away

## Contribute

See
[`contributing.md`][health-contributing]
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

[api-child]: #child

[api-h]: #hselector-properties-children

[api-properties]: #properties-1

[api-result]: #result

[api-s]: #sselector-properties-children

[badge-build-image]: https://github.com/syntax-tree/hastscript/workflows/main/badge.svg

[badge-build-url]: https://github.com/syntax-tree/hastscript/actions

[badge-coverage-image]: https://img.shields.io/codecov/c/github/syntax-tree/hastscript.svg

[badge-coverage-url]: https://codecov.io/github/syntax-tree/hastscript

[badge-downloads-image]: https://img.shields.io/npm/dm/hastscript.svg

[badge-downloads-url]: https://www.npmjs.com/package/hastscript

[badge-size-image]: https://img.shields.io/bundlejs/size/hastscript

[badge-size-url]: https://bundlejs.com/?q=hastscript

[esmsh]: https://esm.sh

[file-license]: license

[github-gist-esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[github-hast]: https://github.com/syntax-tree/hast

[github-hast-element]: https://github.com/syntax-tree/hast#element

[github-hast-root]: https://github.com/syntax-tree/hast#root

[github-hast-text]: https://github.com/syntax-tree/hast#text

[github-hast-util-parse-selector]: https://github.com/syntax-tree/hast-util-parse-selector

[github-hast-util-sanitize]: https://github.com/syntax-tree/hast-util-sanitize

[github-unist-builder]: https://github.com/syntax-tree/unist-builder

[github-xastscript]: https://github.com/syntax-tree/xastscript

[health]: https://github.com/syntax-tree/.github

[health-coc]: https://github.com/syntax-tree/.github/blob/main/code-of-conduct.md

[health-contributing]: https://github.com/syntax-tree/.github/blob/main/contributing.md

[health-support]: https://github.com/syntax-tree/.github/blob/main/support.md

[npmjs-install]: https://docs.npmjs.com/cli/install

[typescript]: https://www.typescriptlang.org

[wikipedia-xss]: https://en.wikipedia.org/wiki/Cross-site_scripting

[wooorm]: https://wooorm.com
