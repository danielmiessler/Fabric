'use strict';

var stringify = require('../');

var obj = { c: 8, b: [{ z: 6, y: 5, x: 4 }, 7], a: 3 };
var s = stringify(obj, function (a, b) {
	return b.key.localeCompare(a.key);
});
console.log(s);
