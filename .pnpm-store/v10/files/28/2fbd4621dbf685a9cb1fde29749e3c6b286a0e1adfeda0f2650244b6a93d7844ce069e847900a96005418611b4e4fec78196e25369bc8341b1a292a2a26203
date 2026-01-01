# Svelte Preprocess

> A [Svelte](https://svelte.dev) preprocessor with sensible defaults and support for: PostCSS, SCSS, Less, Stylus, CoffeeScript, TypeScript, Pug and much more.

<p>
  <a href="https://www.npmjs.com/package/svelte-preprocess">
    <img src="https://img.shields.io/npm/v/svelte-preprocess.svg" alt="npm version">
  </a>

  <a href="https://github.com/sveltejs/svelte-preprocess/blob/master/LICENSE">
    <img src="https://img.shields.io/npm/l/svelte-preprocess.svg" alt="license">
  </a>

  <a href="https://github.com/sveltejs/svelte-preprocess/actions?query=workflow%3ACI">
    <img src="https://github.com/sveltejs/svelte-preprocess/workflows/CI/badge.svg" alt="action-CI">
  </a>
</p>

- [What is it?](#what-is-it)
- [Getting Started](docs/getting-started.md)
- [Usage](docs/usage.md)
- [Migration Guide](docs/migration-guide.md)
- [Preprocessing](docs/preprocessing.md)
  - [Preprocessors](docs/preprocessing.md#preprocessors)
- [Features](#features)
  - [Template tag](#template-tag)
  - [External files](#external-files)
  - [Global style](#global-style)
  - [Modern JavaScript syntax](#modern-javascript-syntax)
  - [Replace values](#replace-values)
  - [Built-in support for commonly used languages](#built-in-support-for-commonly-used-languages)

## What is it?

`Svelte`'s own parser understands only JavaScript, CSS and its HTML-like syntax. To make it possible to write components in other languages, such as TypeScript or SCSS, `Svelte` provides the [preprocess API](https://svelte.dev/docs#compile-time-svelte-preprocess), which allows to easily transform the content of your `markup` and your `style`/`script` tags.

Writing your own preprocessor for, i.e SCSS is easy enough, but it can be cumbersome to have to always configure multiple preprocessors for the languages you'll be using.

`svelte-preprocess` is a custom svelte preprocessor that acts as a facilitator to use other languages with Svelte, providing multiple features, sensible defaults and a less noisy development experience.

It is recommended to use with `svelte.config.js` file, located at the project root. For other usage, please refer to [usage documentation](#usage-documentation).

```js
import preprocess from 'svelte-preprocess';

const config = {
  preprocess: preprocess({ ... })
}

export default config;
```

## Features

### Template tag

_Vue-like_ support for defining your markup between a specific tag. The default tag is `template` but it can be [customized](/docs/preprocessing.md#auto-preprocessing-options).

```html
<template>
  <div>Hey</div>
</template>

<style></style>

<script></script>
```

### External files

```html
<template src="./template.html"></template>
<script src="./script.js"></script>
<style src="./style.css"></style>
```

> Note: using a relative path starting with `.` is important. Otherwise `svelte-preprocess` will ignore the `src` attribute.

### Global style

#### `global` attribute

Add a `global` attribute to your `style` tag and instead of scoping the CSS, all of its content will be interpreted as global style.

```html
<style global>
  div {
    color: red;
  }
</style>
```

#### `:global` rule

Use a `:global` rule to only expose parts of the stylesheet:

```html
<style lang="scss">
  .scoped-style {
  }

  :global {
    @import 'global-stylesheet.scss';

    .global-style {
      .global-child-style {
      }
    }
  }
</style>
```

Works best with nesting-enabled CSS preprocessors, but regular CSS selectors like `div :global .global1 .global2` are also supported.

_**Note**: needs PostCSS to be installed._

### Modern JavaScript syntax

`svelte-preprocess` allows you to run your component code through Babel before sending it to the compiler, allowing you to use new language features such as optional operators and nullish coalescing. However, note that Babel should transpile your component code to the javascript version supported by the Svelte compiler, so ES6+.

For example, with `@babel/preset-env` your config could be:

```js
import preprocess from 'svelte-preprocess'
  ...
  preprocess: preprocess({
    babel: {
      presets: [
        [
          '@babel/preset-env',
          {
            loose: true,
            // No need for babel to resolve modules
            modules: false,
            targets: {
              // ! Very important. Target es6+
              esmodules: true,
            },
          },
        ],
      ],
    },
  });
  ...
```

_**Note:** If you want to transpile your app to be supported in older browsers, you must run babel from the context of your bundler._

### Replace values

Replace a set of string patterns in your components markup by passing an array of `[RegExp, ReplaceFn | string]`, the same arguments received by the `String.prototype.replace` method.

In example, to replace inject the value of `process.env.NODE_ENV`:

```js
autoPreprocess({
  replace: [[/process\.env\.NODE_ENV/g, JSON.stringify(process.env.NODE_ENV)]],
});
```

Which, in a production environment, would turn

```svelte
{#if process.env.NODE_ENV !== 'development'}
  <h1>Production environment!</h1>
{/if}
```

into

```svelte
{#if "production" !== 'development'}
  <h1>Production environment!</h1>
{/if}
```

### Built-in support for commonly used languages

The current supported languages out-of-the-box are Sass, Stylus, Less, CoffeeScript, TypeScript, Pug, PostCSS, Babel.

```html
<template lang="pug">
  div Posts +each('posts as post') a(href="{post.url}") {post.title}
</template>

<script lang="ts">
  export const hello: string = 'world';
</script>

<style src="./style.scss"></style>
```

---

### [Getting started](/docs/getting-started.md)

### [Preprocessing documentation](/docs/preprocessing.md)

### [Usage documentation](/docs/usage.md)

### [Migration Guide](/docs/migration-guide.md)

---

## License

[MIT](LICENSE)
