# acorn-typescript

[![License](https://img.shields.io/npm/l/svelte.svg)](LICENSE.md) [![Chat](https://img.shields.io/discord/457912077277855764?label=chat&logo=discord)](https://svelte.dev/chat)

This is a plugin for [Acorn](http://marijnhaverbeke.nl/acorn/) - a tiny, fast JavaScript parser, written completely in JavaScript.

It was created as an experimental alternative, faster [TypeScript](https://www.typescriptlang.org/) parser. It will help you to parse
TypeScript using Acorn.

## Usage

To get started, import the plugin and use Acorn's extension mechanism to register it. You have to enable `options.locations` while using `acorn-typescript`.

```typescript
import * as acorn from 'acorn';
import tsPlugin from 'acorn-typescript';

const node = acorn.Parser.extend(tsPlugin()).parse(
	`
const a = 1
type A = number
export {
  a,
  type A as B
}
`,
	{
		sourceType: 'module',
		ecmaVersion: 'latest',
		locations: true
	}
);
```

If you want to enable parsing within a TypeScript ambient context, where certain syntax have different rules (like `.d.ts` files and inside [declare module blocks](https://www.typescriptlang.org/docs/handbook/declaration-files/introduction.html)):

```typescript
import * as acorn from 'acorn';
import tsPlugin from 'acorn-typescript';

const node = acorn.Parser.extend(tsPlugin({ dts: true })).parse(
	`
const a = 1
type A = number
export {
  a,
  type A as B
}
`,
	{
		sourceType: 'module',
		ecmaVersion: 'latest',
		locations: true
	}
);
```

## SUPPORTED

- Typescript normal syntax
- Support to parse TypeScript [Decorators](https://www.typescriptlang.org/docs/handbook/decorators.html)
- Support to parse JSX & TSX

## CHANGELOG

[click](./CHANGELOG.md)

## Acknowledgments

We want to thank [TyrealHu](https://github.com/TyrealHu) for his original work on this project. He maintained [`acorn-typescript`](https://github.com/TyrealHu/acorn-typescript) until early 2024.
