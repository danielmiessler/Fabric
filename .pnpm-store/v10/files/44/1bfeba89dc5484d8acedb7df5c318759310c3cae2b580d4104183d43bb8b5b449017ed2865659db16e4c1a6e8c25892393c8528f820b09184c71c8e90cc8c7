'use strict';

var test = require('tape');
var stringify = require('../');

test('custom comparison function', function (t) {
	t.plan(1);
	var obj = { c: 8, b: [{ z: 6, y: 5, x: 4 }, 7], a: 3 };
	var s = stringify(obj, function (a, b) {
		return a.key < b.key ? 1 : -1;
	});
	t.equal(s, '{"c":8,"b":[{"z":6,"y":5,"x":4},7],"a":3}');
});

test('custom comparison function with get', function (t) {
	t.plan(2);

	stringify({ a: 1, b: 2 }, /** @type {import('..').Comparator} */ function (_a, _b) { // eslint-disable-line no-unused-vars
		t.equal(arguments[2], undefined, 'comparator options not passed when not explicitly requested');
		return NaN;
	});

	var obj = { c: 8, b: [{ z: 7, y: 6, x: 4, v: 2, '!v': 3 }, 7], a: 3 };
	var s = stringify(obj, function (a, b, options) {
		var get = options.get;
		// @ts-expect-error implicit coercion here is fine
		var v1 = (get('!' + a.key) || 0) + a.value;
		// @ts-expect-error implicit coercion here is fine
		var v2 = (get('!' + b.key) || 0) + b.value;
		return v1 - v2;
	});
	t.equal(s, '{"c":8,"b":[{"!v":3,"x":4,"v":2,"y":6,"z":7},7],"a":3}');
});
