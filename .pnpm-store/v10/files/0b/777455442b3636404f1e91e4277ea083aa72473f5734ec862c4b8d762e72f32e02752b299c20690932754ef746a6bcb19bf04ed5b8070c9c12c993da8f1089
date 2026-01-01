'use strict';

var stringify = require('../');

var obj = { d: 6, c: 5, b: [{ z: 3, y: 2, x: 1 }, 9], a: 10 };

var s = stringify(obj, /** @type {import('..').Comparator} */ function (a, b) {
	// @ts-expect-error implicit coercion here is fine
	return a.value < b.value ? 1 : -1;
});

console.log(s);
