# jsonify <sup>[![Version Badge][npm-version-svg]][package-url]</sup>

[![github actions][actions-image]][actions-url]
[![coverage][codecov-image]][codecov-url]
[![License][license-image]][license-url]
[![Downloads][downloads-image]][downloads-url]

[![npm badge][npm-badge-png]][package-url]

This module provides Douglas Crockford's JSON implementation without modifying any globals.

`stringify` and `parse` are merely exported without respect to whether or not a global `JSON` object exists.

[![build status](https://secure.travis-ci.org/ljharb/jsonify.png)](https://travis-ci.org/ljharb/jsonify)

# methods

``` js
var json = require('jsonify');
```

## json.parse(source, reviver)

Return a new javascript object from a parse of the `source` string.

If a `reviver` function is specified, walk the structure passing each name/value pair to `reviver.call(parent, key, value)` to transform the `value` before parsing it.

## json.stringify(value, replacer, space)

Return a string representation for `value`.

If `replacer` is specified, walk the structure passing each name/value pair to `replacer.call(parent, key, value)` to transform the `value` before stringifying it.

If `space` is a number, indent the result by that many spaces.
If `space` is a string, use `space` as the indentation.

# install

With [npm](https://npmjs.org) do:

```
npm install jsonify
```

To use this module in the browser, check out
[browserify](https://github.com/browserify/browserify).

# license

public domain

[package-url]: https://npmjs.org/package/jsonify
[npm-version-svg]: https://versionbadg.es/ljharb/jsonify.svg
[deps-svg]: https://david-dm.org/ljharb/jsonify.svg
[deps-url]: https://david-dm.org/ljharb/jsonify
[dev-deps-svg]: https://david-dm.org/ljharb/jsonify/dev-status.svg
[dev-deps-url]: https://david-dm.org/ljharb/jsonify#info=devDependencies
[npm-badge-png]: https://nodei.co/npm/jsonify.png?downloads=true&stars=true
[license-image]: https://img.shields.io/npm/l/jsonify.svg
[license-url]: LICENSE
[downloads-image]: https://img.shields.io/npm/dm/jsonify.svg
[downloads-url]: https://npm-stat.com/charts.html?package=jsonify
[codecov-image]: https://codecov.io/gh/ljharb/jsonify/branch/main/graphs/badge.svg
[codecov-url]: https://app.codecov.io/gh/ljharb/jsonify/
[actions-image]: https://img.shields.io/endpoint?url=https://github-actions-badge-u3jn4tfpocch.runkit.sh/ljharb/jsonify
[actions-url]: https://github.com/ljharb/jsonify/actions
