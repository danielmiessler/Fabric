# code-red

Experimental toolkit for writing x-to-JavaScript compilers. It is used in [Svelte](https://svelte.dev).


## API

The `code-red` package exposes three core functions — `b`, `x` and `print`.

`b` and `x` take a template literal and return an [ESTree](https://github.com/estree/estree) program body, or a single node:

```js
import { b, x } from 'code-red';

const expression = x`i + j`;

assert.equal(expression.type, 'AssignmentExpression');
assert.equal(expression.operator, '+');
assert.equal(expression.left.name, 'i');
assert.equal(expression.right.name, 'j');

const body = b`
	const i = 1;
	const j = 2;
	const k = i + j;
`;

assert.equal(body.length, 3);
assert.equal(body[0].type, 'VariableDeclaration');
```

Expressions in template literals correspond to replacement nodes — so you could express the above like so:

```js
const i = x`i`;
const j = x`j`;
const expression = x`${i} + ${j}`;

const body = b`
	const ${i} = 1;
	const ${j} = 2;
	const k = ${expression};
`;
```

The `print` function takes a node and turns it into a `{code, map}` object:

```js
const add = x`
	function add(${i}, ${j}) {
		return ${expression};
	}
`;

print(add).code;
/*
function add(i, j) {
	return i + j;
}
*/

i.name = 'foo';
j.name = 'bar';

print(add).code;
/*
function add(foo, bar) {
	return foo + bar;
}
*/
```

## Prefixes

### `@`-prefixed names (replaceable globals)

So that you can use globals in your code. In Svelte, we use this to insert utility functions.

```js
// input
import { x } from 'code-red';
x`@foo(bar)`

// output
FOO(bar)
```

### `#`-prefixed names (automatically deconflicted names)

So that you can insert variables in your code without worrying if they clash with existing variable names.


`bar` used in user code and in inserted code gets a `$1` suffix:

```js
// input
import { x } from 'code-red';
x`
function foo(#bar) {
	return #bar * bar;
}`;

// output
function foo(bar$1) {
	return bar$1 * bar;
}
```

Without conflicts, no `$1` suffix:

```js
// input
import { b } from 'code-red';
b`const foo = #bar => #bar * 2`;

// output
const foo = bar => bar * 2;
```

## Optimiser

TODO add an optimiser that e.g. collapses consecutive identical if blocks


## Compiler

TODO add a `code-red/compiler` module that replaces template literals with the nodes they evaluate to, so that there's nothing to parse at runtime.


## Sourcemaps

TODO support source mappings for inserted nodes with location information.


## License

[MIT](LICENSE)
