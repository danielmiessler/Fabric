# rehype-unwrap-images

[![Build][badge-build-image]][badge-build-url]
[![Coverage][badge-coverage-image]][badge-coverage-url]
[![Downloads][badge-downloads-image]][badge-downloads-url]
[![Size][badge-size-image]][badge-size-url]
[![Sponsors][badge-sponsors-image]][badge-collective-url]
[![Backers][badge-backers-image]][badge-collective-url]
[![Chat][badge-chat-image]][badge-chat-url]

**[rehype][github-rehype]** plugin to remove the wrapping paragraph (`<p>`) for
images (`<img>`).

## Contents

* [What is this?](#what-is-this)
* [When should I use this?](#when-should-i-use-this)
* [Install](#install)
* [Use](#use)
* [API](#api)
* [Types](#types)
* [Compatibility](#compatibility)
* [Security](#security)
* [Contribute](#contribute)
* [License](#license)

## What is this?

This package is a [unified][github-unified]
([rehype][github-rehype])
plugin that searches for paragraphs which contain only images
(possibly in links)
and nothing else,
and then removes those surrounding paragraphs.

## When should I use this?

When working with markdown,
this project can make it simpler to style images with CSS,
for example displaying them at the full available width,
because paragraph styles no longer interfere with them.

## Install

This package is [ESM only][github-gist-esm].
In Node.js (version 16+), install with [npm][npm-install]:

```sh
npm install rehype-unwrap-images
```

In Deno with [`esm.sh`][esmsh]:

```js
import rehypeUnwrapImages from 'https://esm.sh/rehype-unwrap-images@1'
```

In browsers with [`esm.sh`][esmsh]:

```html
<script type="module">
  import rehypeUnwrapImages from 'https://esm.sh/rehype-unwrap-images@1?bundle'
</script>
```

## Use

Say we have the following file `example.html`.

```html
<h1>Saturn</h1>
<p>Saturn is the sixth planet from the Sun and the second-largest in the Solar
System, after Jupiter.</p>
<p><img alt="Saturn" src="//upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Saturn_during_Equinox.jpg/300px-Saturn_during_Equinox.jpg"></p>
```

…and a module `example.js`:

```js
import rehypeParse from 'rehype-parse'
import rehypeStringify from 'rehype-stringify'
import rehypeUnwrapImages from 'rehype-unwrap-images'
import {read} from 'to-vfile'
import {unified} from 'unified'

const file = await read('example.html')

await unified()
  .use(rehypeParse, {fragment: true})
  .use(rehypeUnwrapImages)
  .use(rehypeStringify)
  .process(file)

console.log(String(file))
```

…then running `node example.js` yields:

```html
<h1>Saturn</h1>
<p>Saturn is the sixth planet from the Sun and the second-largest in the Solar
System, after Jupiter.</p>
<img alt="Saturn" src="//upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Saturn_during_Equinox.jpg/300px-Saturn_during_Equinox.jpg">
```

## API

This package exports no identifiers.
The default export is [`rehypeUnwrapImages`][api-rehype-unwrap-images].

#### `unified().use(rehypeUnwrapImages)`

Remove the wrapping paragraph for images.

###### Parameters

There are no parameters.

###### Returns

Transform ([`Transformer`][github-unified-transformer]).

## Types

This package is fully typed with [TypeScript][].
It exports no additional types.

## Compatibility

Projects maintained by the unified collective are compatible with maintained
versions of Node.js.

When we cut a new major release, we drop support for unmaintained versions of
Node.
This means we try to keep the current release line, `rehype-unwrap-images@1`,
compatible with Node.js 16.

This plugin works with `unified` version 6+ and `rehype` version 1+.

## Security

Use of `rehype-unwrap-images` does not change the tree other than sometimes
removing `<p>`.
Use of this plugin does not open you up for a
[cross-site scripting (XSS)][wiki-xss] attack.
When in doubt,
use [`rehype-sanitize`][github-rehype-sanitize].

## Contribute

See [`contributing.md`][health-contributing] in [`rehypejs/.github`][health]
for ways to get started.
See [`support.md`][health-support] for ways to get help.

This project has a [code of conduct][health-coc].
By interacting with this repository, organization, or community you agree to
abide by its terms.

## License

[MIT][file-license] © [Titus Wormer][wooorm]

<!-- Definitions -->

[api-rehype-unwrap-images]: #unifieduserehypeunwrapimages

[badge-backers-image]: https://opencollective.com/unified/backers/badge.svg

[badge-build-image]: https://github.com/rehypejs/rehype-unwrap-images/actions/workflows/main.yml/badge.svg

[badge-build-url]: https://github.com/rehypejs/rehype-unwrap-images/actions

[badge-collective-url]: https://opencollective.com/unified

[badge-coverage-image]: https://img.shields.io/codecov/c/github/rehypejs/rehype-unwrap-images.svg

[badge-coverage-url]: https://codecov.io/github/rehypejs/rehype-unwrap-images

[badge-downloads-image]: https://img.shields.io/npm/dm/rehype-unwrap-images.svg

[badge-downloads-url]: https://www.npmjs.com/package/rehype-unwrap-images

[badge-size-image]: https://img.shields.io/bundlejs/size/rehype-unwrap-images

[badge-size-url]: https://bundlejs.com/?q=rehype-unwrap-images

[badge-sponsors-image]: https://opencollective.com/unified/sponsors/badge.svg

[badge-chat-image]: https://img.shields.io/badge/chat-discussions-success.svg

[badge-chat-url]: https://github.com/rehypejs/rehype/discussions

[esmsh]: https://esm.sh

[file-license]: license

[github-gist-esm]: https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c

[github-rehype]: https://github.com/rehypejs/rehype

[github-rehype-sanitize]: https://github.com/rehypejs/rehype-sanitize

[github-unified]: https://github.com/unifiedjs/unified

[github-unified-transformer]: https://github.com/unifiedjs/unified#transformer

[health-coc]: https://github.com/rehypejs/.github/blob/main/code-of-conduct.md

[health-contributing]: https://github.com/rehypejs/.github/blob/main/contributing.md

[health-support]: https://github.com/rehypejs/.github/blob/main/support.md

[health]: https://github.com/rehypejs/.github

[npm-install]: https://docs.npmjs.com/cli/install

[typescript]: https://www.typescriptlang.org

[wooorm]: https://wooorm.com

[wiki-xss]: https://en.wikipedia.org/wiki/Cross-site_scripting
