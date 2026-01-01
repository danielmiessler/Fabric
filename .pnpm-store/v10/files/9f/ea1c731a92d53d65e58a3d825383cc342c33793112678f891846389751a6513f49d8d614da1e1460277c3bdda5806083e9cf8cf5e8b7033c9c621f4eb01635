'use strict';

/** @type {typeof JSON.stringify} */
var jsonStringify = (typeof JSON !== 'undefined' ? JSON : require('jsonify')).stringify;

var isArray = require('isarray');
var objectKeys = require('object-keys');
var callBind = require('call-bind');
var callBound = require('call-bound');

var $join = callBound('Array.prototype.join');
var $indexOf = callBound('Array.prototype.indexOf');
var $splice = callBound('Array.prototype.splice');
var $sort = callBound('Array.prototype.sort');

/** @type {(n: number, char: string) => string} */
var strRepeat = function repeat(n, char) {
	var str = '';
	for (var i = 0; i < n; i += 1) {
		str += char;
	}
	return str;
};

/** @type {(parent: import('.').Node, key: import('.').Key, value: unknown) => unknown} */
var defaultReplacer = function (_parent, _key, value) { return value; };

/** @type {import('.')} */
module.exports = function stableStringify(obj) {
	/** @type {Parameters<import('.')>[1]} */
	var opts = arguments.length > 1 ? arguments[1] : void undefined;
	var space = (opts && opts.space) || '';
	if (typeof space === 'number') { space = strRepeat(space, ' '); }
	var cycles = !!opts && typeof opts.cycles === 'boolean' && opts.cycles;
	/** @type {undefined | typeof defaultReplacer} */
	var replacer = opts && opts.replacer ? callBind(opts.replacer) : defaultReplacer;
	if (opts && typeof opts.collapseEmpty !== 'undefined' && typeof opts.collapseEmpty !== 'boolean') {
		throw new TypeError('`collapseEmpty` must be a boolean, if provided');
	}
	var collapseEmpty = !!opts && opts.collapseEmpty;

	var cmpOpt = typeof opts === 'function' ? opts : opts && opts.cmp;
	/** @type {undefined | (<T extends import('.').NonArrayNode>(node: T) => (a: Exclude<keyof T, symbol | number>, b: Exclude<keyof T, symbol | number>) => number)} */
	var cmp = cmpOpt && function (node) {
		// eslint-disable-next-line no-extra-parens
		var get = /** @type {NonNullable<typeof cmpOpt>} */ (cmpOpt).length > 2
			&& /** @type {import('.').Getter['get']} */ function get(k) { return node[k]; };
		return function (a, b) {
			// eslint-disable-next-line no-extra-parens
			return /** @type {NonNullable<typeof cmpOpt>} */ (cmpOpt)(
				{ key: a, value: node[a] },
				{ key: b, value: node[b] },
				// @ts-expect-error TS doesn't understand the optimization used here
				get ? /** @type {import('.').Getter} */ { __proto__: null, get: get } : void undefined
			);
		};
	};

	/** @type {import('.').Node[]} */
	var seen = [];
	return (/** @type {(parent: import('.').Node, key: string | number, node: unknown, level: number) => string | undefined} */
		function stringify(parent, key, node, level) {
			var indent = space ? '\n' + strRepeat(level, space) : '';
			var colonSeparator = space ? ': ' : ':';

			// eslint-disable-next-line no-extra-parens
			if (node && /** @type {{ toJSON?: unknown }} */ (node).toJSON && typeof /** @type {{ toJSON?: unknown }} */ (node).toJSON === 'function') {
				// eslint-disable-next-line no-extra-parens
				node = /** @type {{ toJSON: Function }} */ (node).toJSON();
			}

			node = replacer(parent, key, node);
			if (node === undefined) {
				return;
			}
			if (typeof node !== 'object' || node === null) {
				return jsonStringify(node);
			}

			/** @type {(out: string[], brackets: '[]' | '{}') => string} */
			var groupOutput = function (out, brackets) {
				return collapseEmpty && out.length === 0
					? brackets
					: (brackets === '[]' ? '[' : '{') + $join(out, ',') + indent + (brackets === '[]' ? ']' : '}');
			};

			if (isArray(node)) {
				var out = [];
				for (var i = 0; i < node.length; i++) {
					var item = stringify(node, i, node[i], level + 1) || jsonStringify(null);
					out[out.length] = indent + space + item;
				}
				return groupOutput(out, '[]');
			}

			if ($indexOf(seen, node) !== -1) {
				if (cycles) { return jsonStringify('__cycle__'); }
				throw new TypeError('Converting circular structure to JSON');
			} else {
				seen[seen.length] = /** @type {import('.').NonArrayNode} */ (node);
			}

			/** @type {import('.').Key[]} */
			// eslint-disable-next-line no-extra-parens
			var keys = $sort(objectKeys(node), cmp && cmp(/** @type {import('.').NonArrayNode} */ (node)));
			var out = [];
			for (var i = 0; i < keys.length; i++) {
				var key = keys[i];
				// eslint-disable-next-line no-extra-parens
				var value = stringify(/** @type {import('.').Node} */ (node), key, /** @type {import('.').NonArrayNode} */ (node)[key], level + 1);

				if (!value) { continue; }

				var keyValue = jsonStringify(key)
					+ colonSeparator
					+ value;

				out[out.length] = indent + space + keyValue;
			}
			$splice(seen, $indexOf(seen, node), 1);
			return groupOutput(out, '{}');
		}({ '': obj }, '', obj, 0)
	);
};
