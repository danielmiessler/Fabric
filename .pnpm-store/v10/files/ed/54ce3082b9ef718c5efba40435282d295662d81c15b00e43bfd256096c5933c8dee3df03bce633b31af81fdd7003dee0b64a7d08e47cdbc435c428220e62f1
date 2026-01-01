# svelte-eslint-parser

[Svelte] parser for [ESLint].  
You can check it on [Online DEMO](https://sveltejs.github.io/svelte-eslint-parser/playground).

**_Note that this parser has experimental support for Svelte v5, but may break with new versions of Svelte v5._**

[![NPM license](https://img.shields.io/npm/l/svelte-eslint-parser.svg)](https://www.npmjs.com/package/svelte-eslint-parser)
[![NPM version](https://img.shields.io/npm/v/svelte-eslint-parser.svg)](https://www.npmjs.com/package/svelte-eslint-parser)
[![NPM downloads](https://img.shields.io/badge/dynamic/json.svg?label=downloads&colorB=green&suffix=/day&query=$.downloads&uri=https://api.npmjs.org//downloads/point/last-day/svelte-eslint-parser&maxAge=3600)](http://www.npmtrends.com/svelte-eslint-parser)
[![NPM downloads](https://img.shields.io/npm/dw/svelte-eslint-parser.svg)](http://www.npmtrends.com/svelte-eslint-parser)
[![NPM downloads](https://img.shields.io/npm/dm/svelte-eslint-parser.svg)](http://www.npmtrends.com/svelte-eslint-parser)
[![NPM downloads](https://img.shields.io/npm/dy/svelte-eslint-parser.svg)](http://www.npmtrends.com/svelte-eslint-parser)
[![NPM downloads](https://img.shields.io/npm/dt/svelte-eslint-parser.svg)](http://www.npmtrends.com/svelte-eslint-parser)
[![Build Status](https://github.com/sveltejs/svelte-eslint-parser/workflows/CI/badge.svg?branch=main)](https://github.com/sveltejs/svelte-eslint-parser/actions?query=workflow%3ACI)
[![Coverage Status](https://coveralls.io/repos/github/sveltejs/svelte-eslint-parser/badge.svg?branch=main)](https://coveralls.io/github/sveltejs/svelte-eslint-parser?branch=main)

## ⤴️ Motivation

The [svelte-eslint-parser] aims to make it easy to create your own ESLint rules for [Svelte].

The [`eslint-plugin-svelte`] is an ESLint plugin that uses the [svelte-eslint-parser]. I have already [implemented some rules].

[`eslint-plugin-svelte`]: https://github.com/sveltejs/eslint-plugin-svelte
[implemented some rules]: https://sveltejs.github.io/eslint-plugin-svelte/rules/

### ESLint Plugins Using svelte-eslint-parser

#### [eslint-plugin-svelte](https://sveltejs.github.io/eslint-plugin-svelte/)

ESLint plugin for Svelte.  
It provides many unique check rules by using the template AST.

#### [@intlify/eslint-plugin-svelte](https://github.com/intlify/eslint-plugin-svelte)

ESLint plugin for internationalization (i18n) with Svelte.  
It provides rules to help internationalization your application created with Svelte.

## ❗ Attention

The [svelte-eslint-parser] can not be used with the [eslint-plugin-svelte3].

[svelte-eslint-parser]: https://github.com/sveltejs/svelte-eslint-parser

## 💿 Installation

```bash
npm install --save-dev eslint svelte-eslint-parser
```

## 📖 Usage

1. Write `parser` option into your ESLint Config file.
2. Use glob patterns or `--ext .svelte` CLI option.

### ESLint Config (`eslint.config.js`)

```js
import js from "@eslint/js";
import svelteParser from "svelte-eslint-parser";
export default [
  js.configs.recommended,
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
    },
  },
];
```

### ESLint Config (`.eslintrc.*`)

```json
{
  "extends": "eslint:recommended",
  "overrides": [
    {
      "files": ["*.svelte"],
      "parser": "svelte-eslint-parser"
    }
  ]
}
```

### CLI

```console
$ eslint "src/**/*.{js,svelte}"
# or
$ eslint src --ext .svelte
```

## 🔧 Options

[`parserOptions`] has the same properties as what [espree](https://github.com/eslint/espree#usage), the default parser of ESLint, is supporting.
For example:

```json
{
  "parserOptions": {
    "sourceType": "module",
    "ecmaVersion": 2021,
    "ecmaFeatures": {
      "globalReturn": false,
      "impliedStrict": false,
      "jsx": false
    }
  }
}
```

[`parserOptions`]: https://eslint.org/docs/latest/use/configure/parser#configure-parser-options

### parserOptions.parser

You can use `parserOptions.parser` property to specify a custom parser to parse `<script>` tags.
Other properties than parser would be given to the specified parser.

For example in `eslint.config.js`:

```js
import tsParser from "@typescript-eslint/parser";
export default [
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        parser: tsParser,
      },
    },
  },
];
```

For example in `.eslintrc.*`:

```json
{
  "parser": "svelte-eslint-parser",
  "parserOptions": {
    "parser": "@typescript-eslint/parser"
  }
}
```

If you are using the `"@typescript-eslint/parser"`, and if you want to use TypeScript in `<script>` of `.svelte`, you need to add more `parserOptions` configuration.

For example in `eslint.config.js`:

```js
import tsParser from "@typescript-eslint/parser";
export default [
  // ...
  {
    // ...
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        // ...
        project: "path/to/your/tsconfig.json",
        extraFileExtensions: [".svelte"], // This is a required setting in `@typescript-eslint/parser` v4.24.0.
      },
    },
  },
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
      // Parse the `<script>` in `.svelte` as TypeScript by adding the following configuration.
      parserOptions: {
        parser: tsParser,
      },
    },
    // ...
  },
];
```

For example in `.eslintrc.*`:

```js
module.exports = {
  // ...
  parser: "@typescript-eslint/parser",
  parserOptions: {
    // ...
    project: "path/to/your/tsconfig.json",
    extraFileExtensions: [".svelte"], // This is a required setting in `@typescript-eslint/parser` v4.24.0.
  },
  overrides: [
    {
      files: ["*.svelte"],
      parser: "svelte-eslint-parser",
      // Parse the `<script>` in `.svelte` as TypeScript by adding the following configuration.
      parserOptions: {
        parser: "@typescript-eslint/parser",
      },
    },
    // ...
  ],
  // ...
};
```

#### Multiple parsers

If you want to switch the parser for each lang, specify the object.

For example in `eslint.config.js`:

```js
import tsParser from "@typescript-eslint/parser";
import espree from "espree";
export default [
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        parser: {
          ts: tsParser,
          js: espree,
          typescript: tsParser,
        },
      },
    },
  },
];
```

For example in `.eslintrc.*`:

```json
{
  "parser": "svelte-eslint-parser",
  "parserOptions": {
    "parser": {
      "ts": "@typescript-eslint/parser",
      "js": "espree",
      "typescript": "@typescript-eslint/parser"
    }
  }
}
```

### parserOptions.svelteConfig

If you are using `eslint.config.js`, you can provide a `svelte.config.js` in the `parserOptions.svelteConfig` property.

For example:

```js
import svelteConfig from "./svelte.config.js";
export default [
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        svelteConfig: svelteConfig,
      },
    },
  },
];
```

If `parserOptions.svelteConfig` is not specified, some config will be statically parsed from the `svelte.config.js` file.

The `.eslintrc.*` style configuration cannot load `svelte.config.js` because it cannot use ESM. We recommend using the `eslint.config.js` style configuration.

### parserOptions.svelteFeatures

You can use `parserOptions.svelteFeatures` property to specify how to parse related to Svelte features.

For example in `eslint.config.js`:

```js
export default [
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        svelteFeatures: {
          /* -- Experimental Svelte Features -- */
          /* It may be changed or removed in minor versions without notice. */
          // This option is for Svelte 5. The default value is `true`.
          // If `false`, ESLint will not recognize rune symbols.
          // If not configured this option, The parser will try to read the option from `compilerOptions.runes` from `svelte.config.js`.
          // If `parserOptions.svelteConfig` is not specified and the file cannot be parsed by static analysis, it will behave as `true`.
          runes: true,
          /* -- Experimental Svelte Features -- */
          /* It may be changed or removed in minor versions without notice. */
          // Whether to parse the `generics` attribute.
          // See https://github.com/sveltejs/rfcs/pull/38
          experimentalGenerics: false,
        },
      },
    },
  },
];
```

For example in `.eslintrc.*`:

```jsonc
{
  "parser": "svelte-eslint-parser",
  "parserOptions": {
    "svelteFeatures": {
      /* -- Experimental Svelte Features -- */
      /* It may be changed or removed in minor versions without notice. */
      // This option is for Svelte 5. The default value is `true`.
      // If `false`, ESLint will not recognize rune symbols.
      // If not configured this option, The parser will try to read the option from `compilerOptions.runes` from `svelte.config.js`.
      // If `parserOptions.svelteConfig` is not specified and the file cannot be parsed by static analysis, it will behave as `true`.
      "runes": true,
      /* -- Experimental Svelte Features -- */
      /* It may be changed or removed in minor versions without notice. */
      // Whether to parse the `generics` attribute.
      // See https://github.com/sveltejs/rfcs/pull/38
      "experimentalGenerics": false,
    },
  },
}
```

### Runes support

**_This is an experimental feature. It may be changed or removed in minor versions without notice._**

If you install Svelte v5 the parser will be able to parse runes, and will also be able to parse `*.js` and `*.ts` files.
If you don't want to use Runes, you may need to configure. Please read [parserOptions.svelteFeatures](#parseroptionssveltefeatures) for more details.

When using this mode in an ESLint configuration, it is recommended to set it per file pattern as below.

For example in `eslint.config.js`:

```js
import svelteConfig from "./svelte.config.js";
export default [
  {
    files: ["**/*.svelte", "*.svelte"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        parser: "...",
        svelteConfig,
        /* ... */
      },
    },
  },
  {
    files: ["**/*.svelte.js", "*.svelte.js"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        svelteConfig,
        /* ... */
      },
    },
  },
  {
    files: ["**/*.svelte.ts", "*.svelte.ts"],
    languageOptions: {
      parser: svelteParser,
      parserOptions: {
        parser: "...(ts parser)...",
        svelteConfig,
        /* ... */
      },
    },
  },
];
```

For example in `.eslintrc.*`:

```jsonc
{
  "overrides": [
    {
      "files": ["*.svelte"],
      "parser": "svelte-eslint-parser",
      "parserOptions": {
        "parser": "...",
        /* ... */
      },
    },
    {
      "files": ["*.svelte.js"],
      "parser": "svelte-eslint-parser",
      "parserOptions": {
        /* ... */
      },
    },
    {
      "files": ["*.svelte.ts"],
      "parser": "svelte-eslint-parser",
      "parserOptions": {
        "parser": "...(ts parser)...",
        /* ... */
      },
    },
  ],
}
```

## :computer: Editor Integrations

### Visual Studio Code

Use the [dbaeumer.vscode-eslint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint) extension that Microsoft provides officially.

You have to configure the `eslint.validate` option of the extension to check `.svelte` files, because the extension targets only `*.js` or `*.jsx` files by default.

Example **.vscode/settings.json**:

```json
{
  "eslint.validate": ["javascript", "javascriptreact", "svelte"]
}
```

## Usage for Custom Rules / Plugins

- [AST.md](./docs/AST.md) is AST specification. You can check it on the [Online DEMO](https://sveltejs.github.io/svelte-eslint-parser/).
- The parser will generate its own [ScopeManager](https://eslint.org/docs/developer-guide/scope-manager-interface). You can check it on the [Online DEMO](https://sveltejs.github.io/svelte-eslint-parser/scope).
- I have already [implemented some rules] in the [`eslint-plugin-svelte`]. The source code for these rules will be helpful to you.

## :beers: Contributing

Welcome contributing!

Please use GitHub's Issues/PRs.

See also the documentation for the internal mechanism.

- [internal-mechanism.md](./docs/internal-mechanism.md)

## :lock: License

See the [LICENSE](LICENSE) file for license rights and limitations (MIT).

[Svelte]: https://svelte.dev/
[ESLint]: https://eslint.org/
[eslint-plugin-svelte3]: https://github.com/sveltejs/eslint-plugin-svelte3
