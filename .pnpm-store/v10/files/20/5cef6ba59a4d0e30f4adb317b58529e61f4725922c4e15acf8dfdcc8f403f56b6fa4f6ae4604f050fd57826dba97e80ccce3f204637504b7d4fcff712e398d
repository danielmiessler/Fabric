# is-reference

Utility for determining whether an AST node is a reference.

`foo` is a reference in these cases:

```js
console.log(foo);
var foo;
function foo() {}
function bar(foo) {}
export { foo as x };
```

`foo` is *not* a reference in these cases:

```js
var obj = { foo: 1 };
console.log(obj.foo);
export { x as foo };
```

In all cases, `foo` is an `Identifier` node, but the two kinds must be treated differently for the purposes of scope analysis etc. (The examples are non-exhaustive.)


## Installation

```bash
npm install is-reference
```


## Usage

Example using [Acorn](https://github.com/ternjs/acorn) and [estree-walker](https://github.com/Rich-Harris/estree-walker):

```js
import { parse } from 'acorn';
import { walk } from 'estree-walker';
import is_reference from 'is-reference';

const identifiers = [];
const references = [];

const ast = parse(`var a = b.c;`);

walk(ast, {
	enter(node, parent) {
		if (node.type === 'Identifier') identifiers.push(node);
		if (is_reference(node, parent)) references.push(node);
	}
});

identifiers.forEach(node => console.log(node.name)); // a, b, c
references.forEach(node => console.log(node.name)); // a, b
```


## License

MIT
