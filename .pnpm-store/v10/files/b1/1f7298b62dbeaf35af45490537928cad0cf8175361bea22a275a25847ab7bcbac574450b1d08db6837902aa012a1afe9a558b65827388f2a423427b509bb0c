'use strict';

var test = require('tape');
var stringify = require('../');

// @ts-expect-error node ensures this will never fail
var isNode10OrLess = parseInt(process.version.match(/^v(\d+)\./)[1], 10) <= 10;

test('space parameter', function (t) {
	t.plan(1);
	var obj = { one: 1, two: 2 };
	t.equal(
		stringify(obj, { space: '  ' }),
		''
        + '{\n'
        + '  "one": 1,\n'
        + '  "two": 2\n'
        + '}'
	);
});

test('space parameter (with tabs)', function (t) {
	t.plan(1);
	var obj = { one: 1, two: 2 };
	t.equal(
		stringify(obj, { space: '\t' }),
		''
		+ '{\n'
		+ '\t"one": 1,\n'
		+ '\t"two": 2\n'
		+ '}'
	);
});

test('space parameter (with a number)', function (t) {
	t.plan(1);
	var obj = { one: 1, two: 2 };
	t.equal(
		stringify(obj, { space: 3 }),
		''
		+ '{\n'
		+ '   "one": 1,\n'
		+ '   "two": 2\n'
		+ '}'
	);
});

test('space parameter (nested objects)', function (t) {
	t.plan(1);
	var obj = { one: 1, two: { b: 4, a: [2, 3] } };
	t.equal(
		stringify(obj, { space: '  ' }),
		''
		+ '{\n'
		+ '  "one": 1,\n'
		+ '  "two": {\n'
		+ '    "a": [\n'
		+ '      2,\n'
		+ '      3\n'
		+ '    ],\n'
		+ '    "b": 4\n'
		+ '  }\n'
		+ '}'
	);
});

test('space parameter (same as native)', function (t) {
	t.plan(1);
	// for this test, properties need to be in alphabetical order
	var obj = { one: 1, two: { a: [2, 3], b: 4 } };
	t.equal(
		stringify(obj, { space: '  ' }),
		JSON.stringify(obj, null, '  ')
	);
});

test('space parameter base empty behavior: empty arrays and objects have added newline and space', function (t) {
	t.plan(1);
	var obj = { emptyArr: [], emptyObj: {} };
	t.equal(
		stringify(obj, { space: '  ' }),
		'{\n  "emptyArr": [\n  ],\n  "emptyObj": {\n  }\n}'
	);
});

test('space parameter, with collapseEmpty: true', function (t) {
	t.plan(2);
	var obj = { emptyArr: [], emptyObj: {} };

	t['throws'](
		// @ts-expect-error
		function () { stringify(obj, { collapseEmpty: 'not a boolean' }); },
		TypeError
	);

	t.equal(
		stringify(obj, { collapseEmpty: true, space: '  ' }),
		'{\n  "emptyArr": [],\n  "emptyObj": {}\n}'
	);
});

test('space parameter, on a cmp function', function (t) {
	t.plan(3);
	var obj = { one: 1, two: 2 };
	/** @type {import('..').Comparator & import('../').StableStringifyOptions} */
	var cmp = function (a, b) {
		return (a < b ? 1 : -1) * (isNode10OrLess ? -1 : 1);
	};

	t.equal(
		stringify(obj, { space: '\t' }),
		'{\n\t"one": 1,\n\t"two": 2\n}',
		'no cmp option (control)'
	);
	t.equal(
		stringify(obj, { cmp: cmp, space: '\t' }),
		'{\n\t"two": 2,\n\t"one": 1\n}',
		'cmp option in the object'
	);

	cmp.space = '\t';
	t.equal(
		stringify(obj, cmp),
		'{\n\t"two": 2,\n\t"one": 1\n}',
		'cmp passed directly, with space option on it'
	);
});
