# json-stable-stringify <sup>[![Version Badge][npm-version-svg]][package-url]</sup>

[![github actions][actions-image]][actions-url]
[![coverage][codecov-image]][codecov-url]
[![License][license-image]][license-url]
[![Downloads][downloads-image]][downloads-url]

[![npm badge][npm-badge-png]][package-url]

deterministic version of `JSON.stringify()` so you can get a consistent hash from stringified results

You can also pass in a custom comparison function.

# example

``` js
const stringify = require('json-stable-stringify');

const obj = { c: 8, b: [{ z: 6, y: 5, x: 4 }, 7], a: 3 };

console.log(stringify(obj));
```

output:

```
{"a":3,"b":[{"x":4,"y":5,"z":6},7],"c":8}
```

# methods

``` js
const stringify = require('json-stable-stringify')
```

<a id="var-str--stringifyobj-opts"></a>
## const str = stringify(obj, opts)

Return a deterministic stringified string `str` from the object `obj`.

## options

### cmp

If `opts` is given, you can supply an `opts.cmp` to have a custom comparison function for object keys.
Your function `opts.cmp` is called with these parameters:

``` js
opts.cmp({ key: akey, value: avalue }, { key: bkey, value: bvalue }, { get(key): value })
```

For example, to sort on the object key names in reverse order you could write:

``` js
const stringify = require('json-stable-stringify');

const obj = { c: 8, b: [{ z: 6, y: 5, x: 4 },7], a: 3 };

const s = stringify(obj, function (a, b) {
	return b.key.localeCompare(a.key);
});

console.log(s);
```

which results in the output string:

``` js
{"c":8,"b":[{"z":6,"y":5,"x":4},7],"a":3}
```

Or if you wanted to sort on the object values in reverse order, you could write:

``` js
const stringify = require('json-stable-stringify');

const obj = { d: 6, c: 5, b: [{ z: 3, y: 2, x: 1 }, 9], a: 10 };

const s = stringify(obj, function (a, b) {
	return a.value < b.value ? 1 : -1;
});

console.log(s);
```

which outputs:

``` js
{"d":6,"c":5,"b":[{"z":3,"y":2,"x":1},9],"a":10}
```

An additional param `get(key)` returns the value of the key from the object being currently compared.

### space

If you specify `opts.space`, it will indent the output for pretty-printing.
Valid values are strings (e.g. `{space: \t}`) or a number of spaces
(`{space: 3}`).

For example:

```js
const obj = { b: 1, a: { foo: 'bar', and: [1, 2, 3] } };

const s = stringify(obj, { space: '  ' });

console.log(s);
```

which outputs:

```
{
  "a": {
    "and": [
      1,
      2,
      3
    ],
    "foo": "bar"
  },
  "b": 1
}
```

### replacer

The replacer parameter is a function `opts.replacer(key, value)` that behaves the same as the replacer
[from the core JSON object](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_native_JSON#The_replacer_parameter).

# install

With [npm](https://npmjs.org) do:

```
npm install json-stable-stringify
```

# license

MIT

[package-url]: https://npmjs.org/package/json-stable-stringify
[npm-version-svg]: https://versionbadg.es/ljharb/json-stable-stringify.svg
[deps-svg]: https://david-dm.org/ljharb/json-stable-stringify.svg
[deps-url]: https://david-dm.org/ljharb/json-stable-stringify
[dev-deps-svg]: https://david-dm.org/ljharb/json-stable-stringify/dev-status.svg
[dev-deps-url]: https://david-dm.org/ljharb/json-stable-stringify#info=devDependencies
[npm-badge-png]: https://nodei.co/npm/json-stable-stringify.png?downloads=true&stars=true
[license-image]: https://img.shields.io/npm/l/json-stable-stringify.svg
[license-url]: LICENSE
[downloads-image]: https://img.shields.io/npm/dm/json-stable-stringify.svg
[downloads-url]: https://npm-stat.com/charts.html?package=json-stable-stringify
[codecov-image]: https://codecov.io/gh/ljharb/json-stable-stringify/branch/main/graphs/badge.svg
[codecov-url]: https://app.codecov.io/gh/ljharb/json-stable-stringify/
[actions-image]: https://img.shields.io/endpoint?url=https://github-actions-badge-u3jn4tfpocch.runkit.sh/ljharb/json-stable-stringify
[actions-url]: https://github.com/ljharb/json-stable-stringify/actions
