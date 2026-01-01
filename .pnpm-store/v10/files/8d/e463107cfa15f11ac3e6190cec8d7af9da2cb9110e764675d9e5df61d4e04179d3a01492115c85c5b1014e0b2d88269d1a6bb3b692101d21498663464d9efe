'use strict';

var test = require('tape');
var json = require('../');
var garbage = require('garbage');

test('stringify', function (t) {
	for (var i = 0; i < 50; i++) {
		var obj = garbage(50);
		t.equal(
			json.stringify(obj),
			JSON.stringify(obj),
			'comparing JSON.stringify to jsonify.stringify: run ' + (i + 1)
		);
	}

	t.end();
});
