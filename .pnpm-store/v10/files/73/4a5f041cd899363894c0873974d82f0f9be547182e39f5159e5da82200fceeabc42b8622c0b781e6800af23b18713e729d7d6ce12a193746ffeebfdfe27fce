# periscopic

Utility for analyzing scopes belonging to an ESTree-compliant AST.


## API

```js
import { analyze } from 'periscopic';

const ast = acorn.parse(`
const a = b;
console.log(a);
`);

const { map, globals, scope } = analyze(ast);
```

* `map` is a `WeakMap<Node, Scope>`, where the keys are the nodes of your AST that create a scope
* `globals` is a `Map<string, Node>` of all the identifiers that are referenced without being declared anywhere in the program (in this case, `b` and `console`)
* `scope` is the top-level `Scope` belonging to the program


### Scope

Each `Scope` instance has the following properties:

* `scope.block` — true if the scope is created by a block statement (i.e. `let`, `const` and `class` are contained to it), false otherwise
* `scope.parent` — the parent scope object
* `scope.declarations` — a `Map<string, Node>` of all the variables declared in this scope, the node value referes to the declaration statement
* `scope.initialised_declarations` — a `Set<string>` of all the variables declared and initialised in this scope
* `scope.references` — a `Set<string>` of all the names referenced in this scope (or child scopes)

It also has two methods:

* `scope.has(name)` — returns `true` if `name` is declared in this scope or an ancestor scope
* `scope.find_owner(name)` — returns the scope object in which `name` is declared (or `null` if it is not declared)


### `extract_identifiers` and `extract_names`

This package also exposes utilities for extracting the identifiers contained in a declaration or a function parameter:

```js
import { extract_identifiers, extract_names } from 'periscopic';

const ast = acorn.parse(`
const { a, b: [c, d] = e } = opts;
`);

const lhs = ast.body[0].declarations[0].id;

extract_identifiers(lhs);
/*
[
	{ type: 'Identifier', name: 'a', start: 9, end: 10 },
	{ type: 'Identifier', name: 'c', start: 16, end: 17 },
	{ type: 'Identifier', name: 'd', start: 19, end: 20 }
]
*/

extract_names(lhs);
/*
['a', 'c', 'd']
*/
```


## License

[MIT](LICENSE)