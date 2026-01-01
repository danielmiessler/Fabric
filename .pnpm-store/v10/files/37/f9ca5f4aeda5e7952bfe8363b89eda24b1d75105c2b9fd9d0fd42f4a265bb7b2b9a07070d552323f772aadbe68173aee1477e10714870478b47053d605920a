import * as path$2 from 'path';
import path__default, { join } from 'path';
import fs from 'fs';
import { parse as parse$7 } from 'svelte/compiler';

const defineConfig = (config) => config;

function getDefaultExportFromCjs (x) {
	return x && x.__esModule && Object.prototype.hasOwnProperty.call(x, 'default') ? x['default'] : x;
}

var bail_1 = bail$2;

function bail$2(err) {
  if (err) {
    throw err
  }
}

/*!
 * Determine if an object is a Buffer
 *
 * @author   Feross Aboukhadijeh <https://feross.org>
 * @license  MIT
 */

var isBuffer = function isBuffer (obj) {
  return obj != null && obj.constructor != null &&
    typeof obj.constructor.isBuffer === 'function' && obj.constructor.isBuffer(obj)
};

var hasOwn = Object.prototype.hasOwnProperty;
var toStr = Object.prototype.toString;
var defineProperty = Object.defineProperty;
var gOPD = Object.getOwnPropertyDescriptor;

var isArray = function isArray(arr) {
	if (typeof Array.isArray === 'function') {
		return Array.isArray(arr);
	}

	return toStr.call(arr) === '[object Array]';
};

var isPlainObject = function isPlainObject(obj) {
	if (!obj || toStr.call(obj) !== '[object Object]') {
		return false;
	}

	var hasOwnConstructor = hasOwn.call(obj, 'constructor');
	var hasIsPrototypeOf = obj.constructor && obj.constructor.prototype && hasOwn.call(obj.constructor.prototype, 'isPrototypeOf');
	// Not own constructor property must be Object
	if (obj.constructor && !hasOwnConstructor && !hasIsPrototypeOf) {
		return false;
	}

	// Own properties are enumerated firstly, so to speed up,
	// if last one is own, then all properties are own.
	var key;
	for (key in obj) { /**/ }

	return typeof key === 'undefined' || hasOwn.call(obj, key);
};

// If name is '__proto__', and Object.defineProperty is available, define __proto__ as an own property on target
var setProperty = function setProperty(target, options) {
	if (defineProperty && options.name === '__proto__') {
		defineProperty(target, options.name, {
			enumerable: true,
			configurable: true,
			value: options.newValue,
			writable: true
		});
	} else {
		target[options.name] = options.newValue;
	}
};

// Return undefined instead of __proto__ if '__proto__' is not an own property
var getProperty = function getProperty(obj, name) {
	if (name === '__proto__') {
		if (!hasOwn.call(obj, name)) {
			return void 0;
		} else if (gOPD) {
			// In early versions of node, obj['__proto__'] is buggy when obj has
			// __proto__ as an own property. Object.getOwnPropertyDescriptor() works.
			return gOPD(obj, name).value;
		}
	}

	return obj[name];
};

var extend$5 = function extend() {
	var options, name, src, copy, copyIsArray, clone;
	var target = arguments[0];
	var i = 1;
	var length = arguments.length;
	var deep = false;

	// Handle a deep copy situation
	if (typeof target === 'boolean') {
		deep = target;
		target = arguments[1] || {};
		// skip the boolean and the target
		i = 2;
	}
	if (target == null || (typeof target !== 'object' && typeof target !== 'function')) {
		target = {};
	}

	for (; i < length; ++i) {
		options = arguments[i];
		// Only deal with non-null/undefined values
		if (options != null) {
			// Extend the base object
			for (name in options) {
				src = getProperty(target, name);
				copy = getProperty(options, name);

				// Prevent never-ending loop
				if (target !== copy) {
					// Recurse if we're merging plain objects or arrays
					if (deep && copy && (isPlainObject(copy) || (copyIsArray = isArray(copy)))) {
						if (copyIsArray) {
							copyIsArray = false;
							clone = src && isArray(src) ? src : [];
						} else {
							clone = src && isPlainObject(src) ? src : {};
						}

						// Never move original objects, clone them
						setProperty(target, { name: name, newValue: extend(deep, clone, copy) });

					// Don't bring in undefined values
					} else if (typeof copy !== 'undefined') {
						setProperty(target, { name: name, newValue: copy });
					}
				}
			}
		}
	}

	// Return the modified object
	return target;
};

var isPlainObj = value => {
	if (Object.prototype.toString.call(value) !== '[object Object]') {
		return false;
	}

	const prototype = Object.getPrototypeOf(value);
	return prototype === null || prototype === Object.prototype;
};

var slice$3 = [].slice;

var wrap_1$1 = wrap$8;

// Wrap `fn`.
// Can be sync or async; return a promise, receive a completion handler, return
// new values and errors.
function wrap$8(fn, callback) {
  var invoked;

  return wrapped

  function wrapped() {
    var params = slice$3.call(arguments, 0);
    var callback = fn.length > params.length;
    var result;

    if (callback) {
      params.push(done);
    }

    try {
      result = fn.apply(null, params);
    } catch (error) {
      // Well, this is quite the pickle.
      // `fn` received a callback and invoked it (thus continuing the pipeline),
      // but later also threw an error.
      // We’re not about to restart the pipeline again, so the only thing left
      // to do is to throw the thing instead.
      if (callback && invoked) {
        throw error
      }

      return done(error)
    }

    if (!callback) {
      if (result && typeof result.then === 'function') {
        result.then(then, done);
      } else if (result instanceof Error) {
        done(result);
      } else {
        then(result);
      }
    }
  }

  // Invoke `next`, only once.
  function done() {
    if (!invoked) {
      invoked = true;

      callback.apply(null, arguments);
    }
  }

  // Invoke `done` with one value.
  // Tracks if an error is passed, too.
  function then(value) {
    done(null, value);
  }
}

var wrap$7 = wrap_1$1;

var trough_1 = trough$2;

trough$2.wrap = wrap$7;

var slice$2 = [].slice;

// Create new middleware.
function trough$2() {
  var fns = [];
  var middleware = {};

  middleware.run = run;
  middleware.use = use;

  return middleware

  // Run `fns`.  Last argument must be a completion handler.
  function run() {
    var index = -1;
    var input = slice$2.call(arguments, 0, -1);
    var done = arguments[arguments.length - 1];

    if (typeof done !== 'function') {
      throw new Error('Expected function as last argument, not ' + done)
    }

    next.apply(null, [null].concat(input));

    // Run the next `fn`, if any.
    function next(err) {
      var fn = fns[++index];
      var params = slice$2.call(arguments, 0);
      var values = params.slice(1);
      var length = input.length;
      var pos = -1;

      if (err) {
        done(err);
        return
      }

      // Copy non-nully input into values.
      while (++pos < length) {
        if (values[pos] === null || values[pos] === undefined) {
          values[pos] = input[pos];
        }
      }

      input = values;

      // Next or done.
      if (fn) {
        wrap$7(fn, next).apply(null, input);
      } else {
        done.apply(null, [null].concat(input));
      }
    }
  }

  // Add `fn` to the list.
  function use(fn) {
    if (typeof fn !== 'function') {
      throw new Error('Expected `fn` to be a function, not ' + fn)
    }

    fns.push(fn);

    return middleware
  }
}

var own$c = {}.hasOwnProperty;

var unistUtilStringifyPosition = stringify$6;

function stringify$6(value) {
  // Nothing.
  if (!value || typeof value !== 'object') {
    return ''
  }

  // Node.
  if (own$c.call(value, 'position') || own$c.call(value, 'type')) {
    return position$3(value.position)
  }

  // Position.
  if (own$c.call(value, 'start') || own$c.call(value, 'end')) {
    return position$3(value)
  }

  // Point.
  if (own$c.call(value, 'line') || own$c.call(value, 'column')) {
    return point(value)
  }

  // ?
  return ''
}

function point(point) {
  if (!point || typeof point !== 'object') {
    point = {};
  }

  return index(point.line) + ':' + index(point.column)
}

function position$3(pos) {
  if (!pos || typeof pos !== 'object') {
    pos = {};
  }

  return point(pos.start) + '-' + point(pos.end)
}

function index(value) {
  return value && typeof value === 'number' ? value : 1
}

var stringify$5 = unistUtilStringifyPosition;

var vfileMessage = VMessage$1;

// Inherit from `Error#`.
function VMessagePrototype() {}
VMessagePrototype.prototype = Error.prototype;
VMessage$1.prototype = new VMessagePrototype();

// Message properties.
var proto$6 = VMessage$1.prototype;

proto$6.file = '';
proto$6.name = '';
proto$6.reason = '';
proto$6.message = '';
proto$6.stack = '';
proto$6.fatal = null;
proto$6.column = null;
proto$6.line = null;

// Construct a new VMessage.
//
// Note: We cannot invoke `Error` on the created context, as that adds readonly
// `line` and `column` attributes on Safari 9, thus throwing and failing the
// data.
function VMessage$1(reason, position, origin) {
  var parts;
  var range;
  var location;

  if (typeof position === 'string') {
    origin = position;
    position = null;
  }

  parts = parseOrigin(origin);
  range = stringify$5(position) || '1:1';

  location = {
    start: {line: null, column: null},
    end: {line: null, column: null}
  };

  // Node.
  if (position && position.position) {
    position = position.position;
  }

  if (position) {
    // Position.
    if (position.start) {
      location = position;
      position = position.start;
    } else {
      // Point.
      location.start = position;
    }
  }

  if (reason.stack) {
    this.stack = reason.stack;
    reason = reason.message;
  }

  this.message = reason;
  this.name = range;
  this.reason = reason;
  this.line = position ? position.line : null;
  this.column = position ? position.column : null;
  this.location = location;
  this.source = parts[0];
  this.ruleId = parts[1];
}

function parseOrigin(origin) {
  var result = [null, null];
  var index;

  if (typeof origin === 'string') {
    index = origin.indexOf(':');

    if (index === -1) {
      result[1] = origin;
    } else {
      result[0] = origin.slice(0, index);
      result[1] = origin.slice(index + 1);
    }
  }

  return result
}

var Message = /*@__PURE__*/getDefaultExportFromCjs(vfileMessage);

var path$1 = path__default;

function replaceExt(npath, ext) {
  if (typeof npath !== 'string') {
    return npath;
  }

  if (npath.length === 0) {
    return npath;
  }

  var nFileName = path$1.basename(npath, path$1.extname(npath)) + ext;
  return path$1.join(path$1.dirname(npath), nFileName);
}

var replaceExt_1 = replaceExt;

var path = path__default;
var replace = replaceExt_1;
var buffer$1 = isBuffer;

var core$1 = VFile$1;

var own$b = {}.hasOwnProperty;
var proto$5 = VFile$1.prototype;

// Order of setting (least specific to most), we need this because otherwise
// `{stem: 'a', path: '~/b.js'}` would throw, as a path is needed before a
// stem can be set.
var order$1 = ['history', 'path', 'basename', 'stem', 'extname', 'dirname'];

proto$5.toString = toString$3;

// Access full path (`~/index.min.js`).
Object.defineProperty(proto$5, 'path', {get: getPath, set: setPath});

// Access parent path (`~`).
Object.defineProperty(proto$5, 'dirname', {get: getDirname, set: setDirname});

// Access basename (`index.min.js`).
Object.defineProperty(proto$5, 'basename', {get: getBasename, set: setBasename});

// Access extname (`.js`).
Object.defineProperty(proto$5, 'extname', {get: getExtname, set: setExtname});

// Access stem (`index.min`).
Object.defineProperty(proto$5, 'stem', {get: getStem, set: setStem});

// Construct a new file.
function VFile$1(options) {
  var prop;
  var index;
  var length;

  if (!options) {
    options = {};
  } else if (typeof options === 'string' || buffer$1(options)) {
    options = {contents: options};
  } else if ('message' in options && 'messages' in options) {
    return options
  }

  if (!(this instanceof VFile$1)) {
    return new VFile$1(options)
  }

  this.data = {};
  this.messages = [];
  this.history = [];
  this.cwd = process.cwd();

  // Set path related properties in the correct order.
  index = -1;
  length = order$1.length;

  while (++index < length) {
    prop = order$1[index];

    if (own$b.call(options, prop)) {
      this[prop] = options[prop];
    }
  }

  // Set non-path related properties.
  for (prop in options) {
    if (order$1.indexOf(prop) === -1) {
      this[prop] = options[prop];
    }
  }
}

function getPath() {
  return this.history[this.history.length - 1]
}

function setPath(path) {
  assertNonEmpty(path, 'path');

  if (path !== this.path) {
    this.history.push(path);
  }
}

function getDirname() {
  return typeof this.path === 'string' ? path.dirname(this.path) : undefined
}

function setDirname(dirname) {
  assertPath(this.path, 'dirname');
  this.path = path.join(dirname || '', this.basename);
}

function getBasename() {
  return typeof this.path === 'string' ? path.basename(this.path) : undefined
}

function setBasename(basename) {
  assertNonEmpty(basename, 'basename');
  assertPart(basename, 'basename');
  this.path = path.join(this.dirname || '', basename);
}

function getExtname() {
  return typeof this.path === 'string' ? path.extname(this.path) : undefined
}

function setExtname(extname) {
  var ext = extname || '';

  assertPart(ext, 'extname');
  assertPath(this.path, 'extname');

  if (ext) {
    if (ext.charAt(0) !== '.') {
      throw new Error('`extname` must start with `.`')
    }

    if (ext.indexOf('.', 1) !== -1) {
      throw new Error('`extname` cannot contain multiple dots')
    }
  }

  this.path = replace(this.path, ext);
}

function getStem() {
  return typeof this.path === 'string'
    ? path.basename(this.path, this.extname)
    : undefined
}

function setStem(stem) {
  assertNonEmpty(stem, 'stem');
  assertPart(stem, 'stem');
  this.path = path.join(this.dirname || '', stem + (this.extname || ''));
}

// Get the value of the file.
function toString$3(encoding) {
  var value = this.contents || '';
  return buffer$1(value) ? value.toString(encoding) : String(value)
}

// Assert that `part` is not a path (i.e., does not contain `path.sep`).
function assertPart(part, name) {
  if (part.indexOf(path.sep) !== -1) {
    throw new Error(
      '`' + name + '` cannot be a path: did not expect `' + path.sep + '`'
    )
  }
}

// Assert that `part` is not empty.
function assertNonEmpty(part, name) {
  if (!part) {
    throw new Error('`' + name + '` cannot be empty')
  }
}

// Assert `path` exists.
function assertPath(path, name) {
  if (!path) {
    throw new Error('Setting `' + name + '` requires `path` to be set too')
  }
}

var VMessage = vfileMessage;
var VFile = core$1;

var vfile$2 = VFile;

var proto$4 = VFile.prototype;

proto$4.message = message;
proto$4.info = info$1;
proto$4.fail = fail;

// Create a message with `reason` at `position`.
// When an error is passed in as `reason`, copies the stack.
function message(reason, position, origin) {
  var filePath = this.path;
  var message = new VMessage(reason, position, origin);

  if (filePath) {
    message.name = filePath + ':' + message.name;
    message.file = filePath;
  }

  message.fatal = false;

  this.messages.push(message);

  return message
}

// Fail: creates a vmessage, associates it with the file, and throws it.
function fail() {
  var message = this.message.apply(this, arguments);

  message.fatal = true;

  throw message
}

// Info: creates a vmessage, associates it with the file, and marks the fatality
// as null.
function info$1() {
  var message = this.message.apply(this, arguments);

  message.fatal = null;

  return message
}

var bail$1 = bail_1;
var buffer = isBuffer;
var extend$4 = extend$5;
var plain$1 = isPlainObj;
var trough$1 = trough_1;
var vfile$1 = vfile$2;

// Expose a frozen processor.
var unified_1$1 = unified$2().freeze();

var slice$1 = [].slice;
var own$a = {}.hasOwnProperty;

// Process pipeline.
var pipeline$1 = trough$1()
  .use(pipelineParse$1)
  .use(pipelineRun$1)
  .use(pipelineStringify$1);

function pipelineParse$1(p, ctx) {
  ctx.tree = p.parse(ctx.file);
}

function pipelineRun$1(p, ctx, next) {
  p.run(ctx.tree, ctx.file, done);

  function done(err, tree, file) {
    if (err) {
      next(err);
    } else {
      ctx.tree = tree;
      ctx.file = file;
      next();
    }
  }
}

function pipelineStringify$1(p, ctx) {
  var result = p.stringify(ctx.tree, ctx.file);
  var file = ctx.file;

  if (result === undefined || result === null) ; else if (typeof result === 'string' || buffer(result)) {
    file.contents = result;
  } else {
    file.result = result;
  }
}

// Function to create the first processor.
function unified$2() {
  var attachers = [];
  var transformers = trough$1();
  var namespace = {};
  var frozen = false;
  var freezeIndex = -1;

  // Data management.
  processor.data = data;

  // Lock.
  processor.freeze = freeze;

  // Plugins.
  processor.attachers = attachers;
  processor.use = use;

  // API.
  processor.parse = parse;
  processor.stringify = stringify;
  processor.run = run;
  processor.runSync = runSync;
  processor.process = process;
  processor.processSync = processSync;

  // Expose.
  return processor

  // Create a new processor based on the processor in the current scope.
  function processor() {
    var destination = unified$2();
    var length = attachers.length;
    var index = -1;

    while (++index < length) {
      destination.use.apply(null, attachers[index]);
    }

    destination.data(extend$4(true, {}, namespace));

    return destination
  }

  // Freeze: used to signal a processor that has finished configuration.
  //
  // For example, take unified itself: it’s frozen.
  // Plugins should not be added to it.
  // Rather, it should be extended, by invoking it, before modifying it.
  //
  // In essence, always invoke this when exporting a processor.
  function freeze() {
    var values;
    var plugin;
    var options;
    var transformer;

    if (frozen) {
      return processor
    }

    while (++freezeIndex < attachers.length) {
      values = attachers[freezeIndex];
      plugin = values[0];
      options = values[1];
      transformer = null;

      if (options === false) {
        continue
      }

      if (options === true) {
        values[1] = undefined;
      }

      transformer = plugin.apply(processor, values.slice(1));

      if (typeof transformer === 'function') {
        transformers.use(transformer);
      }
    }

    frozen = true;
    freezeIndex = Infinity;

    return processor
  }

  // Data management.
  // Getter / setter for processor-specific informtion.
  function data(key, value) {
    if (typeof key === 'string') {
      // Set `key`.
      if (arguments.length === 2) {
        assertUnfrozen$1('data', frozen);

        namespace[key] = value;

        return processor
      }

      // Get `key`.
      return (own$a.call(namespace, key) && namespace[key]) || null
    }

    // Set space.
    if (key) {
      assertUnfrozen$1('data', frozen);
      namespace = key;
      return processor
    }

    // Get space.
    return namespace
  }

  // Plugin management.
  //
  // Pass it:
  // *   an attacher and options,
  // *   a preset,
  // *   a list of presets, attachers, and arguments (list of attachers and
  //     options).
  function use(value) {
    var settings;

    assertUnfrozen$1('use', frozen);

    if (value === null || value === undefined) ; else if (typeof value === 'function') {
      addPlugin.apply(null, arguments);
    } else if (typeof value === 'object') {
      if ('length' in value) {
        addList(value);
      } else {
        addPreset(value);
      }
    } else {
      throw new Error('Expected usable value, not `' + value + '`')
    }

    if (settings) {
      namespace.settings = extend$4(namespace.settings || {}, settings);
    }

    return processor

    function addPreset(result) {
      addList(result.plugins);

      if (result.settings) {
        settings = extend$4(settings || {}, result.settings);
      }
    }

    function add(value) {
      if (typeof value === 'function') {
        addPlugin(value);
      } else if (typeof value === 'object') {
        if ('length' in value) {
          addPlugin.apply(null, value);
        } else {
          addPreset(value);
        }
      } else {
        throw new Error('Expected usable value, not `' + value + '`')
      }
    }

    function addList(plugins) {
      var length;
      var index;

      if (plugins === null || plugins === undefined) ; else if (typeof plugins === 'object' && 'length' in plugins) {
        length = plugins.length;
        index = -1;

        while (++index < length) {
          add(plugins[index]);
        }
      } else {
        throw new Error('Expected a list of plugins, not `' + plugins + '`')
      }
    }

    function addPlugin(plugin, value) {
      var entry = find(plugin);

      if (entry) {
        if (plain$1(entry[1]) && plain$1(value)) {
          value = extend$4(entry[1], value);
        }

        entry[1] = value;
      } else {
        attachers.push(slice$1.call(arguments));
      }
    }
  }

  function find(plugin) {
    var length = attachers.length;
    var index = -1;
    var entry;

    while (++index < length) {
      entry = attachers[index];

      if (entry[0] === plugin) {
        return entry
      }
    }
  }

  // Parse a file (in string or vfile representation) into a unist node using
  // the `Parser` on the processor.
  function parse(doc) {
    var file = vfile$1(doc);
    var Parser;

    freeze();
    Parser = processor.Parser;
    assertParser$1('parse', Parser);

    if (newable$1(Parser, 'parse')) {
      return new Parser(String(file), file).parse()
    }

    return Parser(String(file), file) // eslint-disable-line new-cap
  }

  // Run transforms on a unist node representation of a file (in string or
  // vfile representation), async.
  function run(node, file, cb) {
    assertNode$1(node);
    freeze();

    if (!cb && typeof file === 'function') {
      cb = file;
      file = null;
    }

    if (!cb) {
      return new Promise(executor)
    }

    executor(null, cb);

    function executor(resolve, reject) {
      transformers.run(node, vfile$1(file), done);

      function done(err, tree, file) {
        tree = tree || node;
        if (err) {
          reject(err);
        } else if (resolve) {
          resolve(tree);
        } else {
          cb(null, tree, file);
        }
      }
    }
  }

  // Run transforms on a unist node representation of a file (in string or
  // vfile representation), sync.
  function runSync(node, file) {
    var complete = false;
    var result;

    run(node, file, done);

    assertDone$1('runSync', 'run', complete);

    return result

    function done(err, tree) {
      complete = true;
      bail$1(err);
      result = tree;
    }
  }

  // Stringify a unist node representation of a file (in string or vfile
  // representation) into a string using the `Compiler` on the processor.
  function stringify(node, doc) {
    var file = vfile$1(doc);
    var Compiler;

    freeze();
    Compiler = processor.Compiler;
    assertCompiler$1('stringify', Compiler);
    assertNode$1(node);

    if (newable$1(Compiler, 'compile')) {
      return new Compiler(node, file).compile()
    }

    return Compiler(node, file) // eslint-disable-line new-cap
  }

  // Parse a file (in string or vfile representation) into a unist node using
  // the `Parser` on the processor, then run transforms on that node, and
  // compile the resulting node using the `Compiler` on the processor, and
  // store that result on the vfile.
  function process(doc, cb) {
    freeze();
    assertParser$1('process', processor.Parser);
    assertCompiler$1('process', processor.Compiler);

    if (!cb) {
      return new Promise(executor)
    }

    executor(null, cb);

    function executor(resolve, reject) {
      var file = vfile$1(doc);

      pipeline$1.run(processor, {file: file}, done);

      function done(err) {
        if (err) {
          reject(err);
        } else if (resolve) {
          resolve(file);
        } else {
          cb(null, file);
        }
      }
    }
  }

  // Process the given document (in string or vfile representation), sync.
  function processSync(doc) {
    var complete = false;
    var file;

    freeze();
    assertParser$1('processSync', processor.Parser);
    assertCompiler$1('processSync', processor.Compiler);
    file = vfile$1(doc);

    process(file, done);

    assertDone$1('processSync', 'process', complete);

    return file

    function done(err) {
      complete = true;
      bail$1(err);
    }
  }
}

// Check if `value` is a constructor.
function newable$1(value, name) {
  return (
    typeof value === 'function' &&
    value.prototype &&
    // A function with keys in its prototype is probably a constructor.
    // Classes’ prototype methods are not enumerable, so we check if some value
    // exists in the prototype.
    (keys$2(value.prototype) || name in value.prototype)
  )
}

// Check if `value` is an object with keys.
function keys$2(value) {
  var key;
  for (key in value) {
    return true
  }

  return false
}

// Assert a parser is available.
function assertParser$1(name, Parser) {
  if (typeof Parser !== 'function') {
    throw new Error('Cannot `' + name + '` without `Parser`')
  }
}

// Assert a compiler is available.
function assertCompiler$1(name, Compiler) {
  if (typeof Compiler !== 'function') {
    throw new Error('Cannot `' + name + '` without `Compiler`')
  }
}

// Assert the processor is not frozen.
function assertUnfrozen$1(name, frozen) {
  if (frozen) {
    throw new Error(
      'Cannot invoke `' +
        name +
        '` on a frozen processor.\nCreate a new processor first, by invoking it: use `processor()` instead of `processor`.'
    )
  }
}

// Assert `node` is a unist node.
function assertNode$1(node) {
  if (!node || typeof node.type !== 'string') {
    throw new Error('Expected node, got `' + node + '`')
  }
}

// Assert that `complete` is `true`.
function assertDone$1(name, asyncName, complete) {
  if (!complete) {
    throw new Error(
      '`' + name + '` finished async. Use `' + asyncName + '` instead'
    )
  }
}

var unified$3 = /*@__PURE__*/getDefaultExportFromCjs(unified_1$1);

var immutable = extend$3;

var hasOwnProperty = Object.prototype.hasOwnProperty;

function extend$3() {
    var target = {};

    for (var i = 0; i < arguments.length; i++) {
        var source = arguments[i];

        for (var key in source) {
            if (hasOwnProperty.call(source, key)) {
                target[key] = source[key];
            }
        }
    }

    return target
}

var inherits$1 = {exports: {}};

var inherits_browser = {exports: {}};

var hasRequiredInherits_browser;

function requireInherits_browser () {
	if (hasRequiredInherits_browser) return inherits_browser.exports;
	hasRequiredInherits_browser = 1;
	if (typeof Object.create === 'function') {
	  // implementation from standard node.js 'util' module
	  inherits_browser.exports = function inherits(ctor, superCtor) {
	    if (superCtor) {
	      ctor.super_ = superCtor;
	      ctor.prototype = Object.create(superCtor.prototype, {
	        constructor: {
	          value: ctor,
	          enumerable: false,
	          writable: true,
	          configurable: true
	        }
	      });
	    }
	  };
	} else {
	  // old school shim for old browsers
	  inherits_browser.exports = function inherits(ctor, superCtor) {
	    if (superCtor) {
	      ctor.super_ = superCtor;
	      var TempCtor = function () {};
	      TempCtor.prototype = superCtor.prototype;
	      ctor.prototype = new TempCtor();
	      ctor.prototype.constructor = ctor;
	    }
	  };
	}
	return inherits_browser.exports;
}

try {
  var util = require('util');
  /* istanbul ignore next */
  if (typeof util.inherits !== 'function') throw '';
  inherits$1.exports = util.inherits;
} catch (e) {
  /* istanbul ignore next */
  inherits$1.exports = requireInherits_browser();
}

var inheritsExports = inherits$1.exports;

var xtend$4 = immutable;
var inherits = inheritsExports;

var unherit_1 = unherit$2;

// Create a custom constructor which can be modified without affecting the
// original class.
function unherit$2(Super) {
  var result;
  var key;
  var value;

  inherits(Of, Super);
  inherits(From, Of);

  // Clone values.
  result = Of.prototype;

  for (key in result) {
    value = result[key];

    if (value && typeof value === 'object') {
      result[key] = 'concat' in value ? value.concat() : xtend$4(value);
    }
  }

  return Of

  // Constructor accepting a single argument, which itself is an `arguments`
  // object.
  function From(parameters) {
    return Super.apply(this, parameters)
  }

  // Constructor accepting variadic arguments.
  function Of() {
    if (!(this instanceof Of)) {
      return new From(arguments)
    }

    return Super.apply(this, arguments)
  }
}

var stateToggle = factory$6;

// Construct a state `toggler`: a function which inverses `property` in context
// based on its current value.
// The by `toggler` returned function restores that value.
function factory$6(key, state, ctx) {
  return enter

  function enter() {
    var context = ctx || this;
    var current = context[key];

    context[key] = !state;

    return exit

    function exit() {
      context[key] = current;
    }
  }
}

var vfileLocation$1 = factory$5;

function factory$5(file) {
  var contents = indices(String(file));

  return {
    toPosition: offsetToPositionFactory(contents),
    toOffset: positionToOffsetFactory(contents)
  }
}

// Factory to get the line and column-based `position` for `offset` in the bound
// indices.
function offsetToPositionFactory(indices) {
  return offsetToPosition

  // Get the line and column-based `position` for `offset` in the bound indices.
  function offsetToPosition(offset) {
    var index = -1;
    var length = indices.length;

    if (offset < 0) {
      return {}
    }

    while (++index < length) {
      if (indices[index] > offset) {
        return {
          line: index + 1,
          column: offset - (indices[index - 1] || 0) + 1,
          offset: offset
        }
      }
    }

    return {}
  }
}

// Factory to get the `offset` for a line and column-based `position` in the
// bound indices.
function positionToOffsetFactory(indices) {
  return positionToOffset

  // Get the `offset` for a line and column-based `position` in the bound
  // indices.
  function positionToOffset(position) {
    var line = position && position.line;
    var column = position && position.column;

    if (!isNaN(line) && !isNaN(column) && line - 1 in indices) {
      return (indices[line - 2] || 0) + column - 1 || 0
    }

    return -1
  }
}

// Get indices of line-breaks in `value`.
function indices(value) {
  var result = [];
  var index = value.indexOf('\n');

  while (index !== -1) {
    result.push(index + 1);
    index = value.indexOf('\n', index + 1);
  }

  result.push(value.length + 1);

  return result
}

var _unescape = factory$4;

var backslash = '\\';

// Factory to de-escape a value, based on a list at `key` in `ctx`.
function factory$4(ctx, key) {
  return unescape

  // De-escape a string using the expression at `key` in `ctx`.
  function unescape(value) {
    var previous = 0;
    var index = value.indexOf(backslash);
    var escape = ctx[key];
    var queue = [];
    var character;

    while (index !== -1) {
      queue.push(value.slice(previous, index));
      previous = index + 1;
      character = value.charAt(previous);

      // If the following character is not a valid escape, add the slash.
      if (!character || escape.indexOf(character) === -1) {
        queue.push(backslash);
      }

      index = value.indexOf(backslash, previous + 1);
    }

    queue.push(value.slice(previous));

    return queue.join('')
  }
}

var AElig$2 = "Æ";
var AMP$2 = "&";
var Aacute$2 = "Á";
var Acirc$2 = "Â";
var Agrave$2 = "À";
var Aring$2 = "Å";
var Atilde$2 = "Ã";
var Auml$2 = "Ä";
var COPY$1 = "©";
var Ccedil$2 = "Ç";
var ETH$2 = "Ð";
var Eacute$2 = "É";
var Ecirc$2 = "Ê";
var Egrave$2 = "È";
var Euml$2 = "Ë";
var GT$2 = ">";
var Iacute$2 = "Í";
var Icirc$2 = "Î";
var Igrave$2 = "Ì";
var Iuml$2 = "Ï";
var LT$2 = "<";
var Ntilde$2 = "Ñ";
var Oacute$2 = "Ó";
var Ocirc$2 = "Ô";
var Ograve$2 = "Ò";
var Oslash$2 = "Ø";
var Otilde$2 = "Õ";
var Ouml$2 = "Ö";
var QUOT$1 = "\"";
var REG$1 = "®";
var THORN$2 = "Þ";
var Uacute$2 = "Ú";
var Ucirc$2 = "Û";
var Ugrave$2 = "Ù";
var Uuml$2 = "Ü";
var Yacute$2 = "Ý";
var aacute$2 = "á";
var acirc$2 = "â";
var acute$2 = "´";
var aelig$2 = "æ";
var agrave$2 = "à";
var amp$2 = "&";
var aring$2 = "å";
var atilde$2 = "ã";
var auml$2 = "ä";
var brvbar$2 = "¦";
var ccedil$2 = "ç";
var cedil$2 = "¸";
var cent$2 = "¢";
var copy$2 = "©";
var curren$2 = "¤";
var deg$2 = "°";
var divide$2 = "÷";
var eacute$2 = "é";
var ecirc$2 = "ê";
var egrave$2 = "è";
var eth$2 = "ð";
var euml$2 = "ë";
var frac12$2 = "½";
var frac14$2 = "¼";
var frac34$2 = "¾";
var gt$2 = ">";
var iacute$2 = "í";
var icirc$2 = "î";
var iexcl$2 = "¡";
var igrave$2 = "ì";
var iquest$2 = "¿";
var iuml$2 = "ï";
var laquo$2 = "«";
var lt$2 = "<";
var macr$2 = "¯";
var micro$2 = "µ";
var middot$2 = "·";
var nbsp$2 = " ";
var not$2 = "¬";
var ntilde$2 = "ñ";
var oacute$2 = "ó";
var ocirc$2 = "ô";
var ograve$2 = "ò";
var ordf$2 = "ª";
var ordm$2 = "º";
var oslash$2 = "ø";
var otilde$2 = "õ";
var ouml$2 = "ö";
var para$2 = "¶";
var plusmn$2 = "±";
var pound$2 = "£";
var quot$2 = "\"";
var raquo$2 = "»";
var reg$2 = "®";
var sect$2 = "§";
var shy$2 = "­";
var sup1$2 = "¹";
var sup2$2 = "²";
var sup3$2 = "³";
var szlig$2 = "ß";
var thorn$2 = "þ";
var times$2 = "×";
var uacute$2 = "ú";
var ucirc$2 = "û";
var ugrave$2 = "ù";
var uml$2 = "¨";
var uuml$2 = "ü";
var yacute$2 = "ý";
var yen$2 = "¥";
var yuml$2 = "ÿ";
var require$$1$2 = {
	AElig: AElig$2,
	AMP: AMP$2,
	Aacute: Aacute$2,
	Acirc: Acirc$2,
	Agrave: Agrave$2,
	Aring: Aring$2,
	Atilde: Atilde$2,
	Auml: Auml$2,
	COPY: COPY$1,
	Ccedil: Ccedil$2,
	ETH: ETH$2,
	Eacute: Eacute$2,
	Ecirc: Ecirc$2,
	Egrave: Egrave$2,
	Euml: Euml$2,
	GT: GT$2,
	Iacute: Iacute$2,
	Icirc: Icirc$2,
	Igrave: Igrave$2,
	Iuml: Iuml$2,
	LT: LT$2,
	Ntilde: Ntilde$2,
	Oacute: Oacute$2,
	Ocirc: Ocirc$2,
	Ograve: Ograve$2,
	Oslash: Oslash$2,
	Otilde: Otilde$2,
	Ouml: Ouml$2,
	QUOT: QUOT$1,
	REG: REG$1,
	THORN: THORN$2,
	Uacute: Uacute$2,
	Ucirc: Ucirc$2,
	Ugrave: Ugrave$2,
	Uuml: Uuml$2,
	Yacute: Yacute$2,
	aacute: aacute$2,
	acirc: acirc$2,
	acute: acute$2,
	aelig: aelig$2,
	agrave: agrave$2,
	amp: amp$2,
	aring: aring$2,
	atilde: atilde$2,
	auml: auml$2,
	brvbar: brvbar$2,
	ccedil: ccedil$2,
	cedil: cedil$2,
	cent: cent$2,
	copy: copy$2,
	curren: curren$2,
	deg: deg$2,
	divide: divide$2,
	eacute: eacute$2,
	ecirc: ecirc$2,
	egrave: egrave$2,
	eth: eth$2,
	euml: euml$2,
	frac12: frac12$2,
	frac14: frac14$2,
	frac34: frac34$2,
	gt: gt$2,
	iacute: iacute$2,
	icirc: icirc$2,
	iexcl: iexcl$2,
	igrave: igrave$2,
	iquest: iquest$2,
	iuml: iuml$2,
	laquo: laquo$2,
	lt: lt$2,
	macr: macr$2,
	micro: micro$2,
	middot: middot$2,
	nbsp: nbsp$2,
	not: not$2,
	ntilde: ntilde$2,
	oacute: oacute$2,
	ocirc: ocirc$2,
	ograve: ograve$2,
	ordf: ordf$2,
	ordm: ordm$2,
	oslash: oslash$2,
	otilde: otilde$2,
	ouml: ouml$2,
	para: para$2,
	plusmn: plusmn$2,
	pound: pound$2,
	quot: quot$2,
	raquo: raquo$2,
	reg: reg$2,
	sect: sect$2,
	shy: shy$2,
	sup1: sup1$2,
	sup2: sup2$2,
	sup3: sup3$2,
	szlig: szlig$2,
	thorn: thorn$2,
	times: times$2,
	uacute: uacute$2,
	ucirc: ucirc$2,
	ugrave: ugrave$2,
	uml: uml$2,
	uuml: uuml$2,
	yacute: yacute$2,
	yen: yen$2,
	yuml: yuml$2
};

var require$$1$1 = {
	"0": "�",
	"128": "€",
	"130": "‚",
	"131": "ƒ",
	"132": "„",
	"133": "…",
	"134": "†",
	"135": "‡",
	"136": "ˆ",
	"137": "‰",
	"138": "Š",
	"139": "‹",
	"140": "Œ",
	"142": "Ž",
	"145": "‘",
	"146": "’",
	"147": "“",
	"148": "”",
	"149": "•",
	"150": "–",
	"151": "—",
	"152": "˜",
	"153": "™",
	"154": "š",
	"155": "›",
	"156": "œ",
	"158": "ž",
	"159": "Ÿ"
};

var isDecimal;
var hasRequiredIsDecimal;

function requireIsDecimal () {
	if (hasRequiredIsDecimal) return isDecimal;
	hasRequiredIsDecimal = 1;

	isDecimal = decimal;

	// Check if the given character code, or the character code at the first
	// character, is decimal.
	function decimal(character) {
	  var code = typeof character === 'string' ? character.charCodeAt(0) : character;

	  return code >= 48 && code <= 57 /* 0-9 */
	}
	return isDecimal;
}

var isHexadecimal = hexadecimal$2;

// Check if the given character code, or the character code at the first
// character, is hexadecimal.
function hexadecimal$2(character) {
  var code = typeof character === 'string' ? character.charCodeAt(0) : character;

  return (
    (code >= 97 /* a */ && code <= 102) /* z */ ||
    (code >= 65 /* A */ && code <= 70) /* Z */ ||
    (code >= 48 /* A */ && code <= 57) /* Z */
  )
}

var isAlphabetical;
var hasRequiredIsAlphabetical;

function requireIsAlphabetical () {
	if (hasRequiredIsAlphabetical) return isAlphabetical;
	hasRequiredIsAlphabetical = 1;

	isAlphabetical = alphabetical;

	// Check if the given character code, or the character code at the first
	// character, is alphabetical.
	function alphabetical(character) {
	  var code = typeof character === 'string' ? character.charCodeAt(0) : character;

	  return (
	    (code >= 97 && code <= 122) /* a-z */ ||
	    (code >= 65 && code <= 90) /* A-Z */
	  )
	}
	return isAlphabetical;
}

var alphabetical = requireIsAlphabetical();
var decimal$2 = requireIsDecimal();

var isAlphanumerical = alphanumerical$2;

// Check if the given character code, or the character code at the first
// character, is alphanumerical.
function alphanumerical$2(character) {
  return alphabetical(character) || decimal$2(character)
}

var AEli = "Æ";
var AElig$1 = "Æ";
var AM = "&";
var AMP$1 = "&";
var Aacut = "Á";
var Aacute$1 = "Á";
var Abreve = "Ă";
var Acir = "Â";
var Acirc$1 = "Â";
var Acy = "А";
var Afr = "𝔄";
var Agrav = "À";
var Agrave$1 = "À";
var Alpha$1 = "Α";
var Amacr = "Ā";
var And = "⩓";
var Aogon = "Ą";
var Aopf = "𝔸";
var ApplyFunction = "⁡";
var Arin = "Å";
var Aring$1 = "Å";
var Ascr = "𝒜";
var Assign = "≔";
var Atild = "Ã";
var Atilde$1 = "Ã";
var Aum = "Ä";
var Auml$1 = "Ä";
var Backslash = "∖";
var Barv = "⫧";
var Barwed = "⌆";
var Bcy = "Б";
var Because = "∵";
var Bernoullis = "ℬ";
var Beta$1 = "Β";
var Bfr = "𝔅";
var Bopf = "𝔹";
var Breve = "˘";
var Bscr = "ℬ";
var Bumpeq = "≎";
var CHcy = "Ч";
var COP = "©";
var COPY = "©";
var Cacute = "Ć";
var Cap = "⋒";
var CapitalDifferentialD = "ⅅ";
var Cayleys = "ℭ";
var Ccaron = "Č";
var Ccedi = "Ç";
var Ccedil$1 = "Ç";
var Ccirc = "Ĉ";
var Cconint = "∰";
var Cdot = "Ċ";
var Cedilla = "¸";
var CenterDot = "·";
var Cfr = "ℭ";
var Chi$1 = "Χ";
var CircleDot = "⊙";
var CircleMinus = "⊖";
var CirclePlus = "⊕";
var CircleTimes = "⊗";
var ClockwiseContourIntegral = "∲";
var CloseCurlyDoubleQuote = "”";
var CloseCurlyQuote = "’";
var Colon = "∷";
var Colone = "⩴";
var Congruent = "≡";
var Conint = "∯";
var ContourIntegral = "∮";
var Copf = "ℂ";
var Coproduct = "∐";
var CounterClockwiseContourIntegral = "∳";
var Cross = "⨯";
var Cscr = "𝒞";
var Cup = "⋓";
var CupCap = "≍";
var DD = "ⅅ";
var DDotrahd = "⤑";
var DJcy = "Ђ";
var DScy = "Ѕ";
var DZcy = "Џ";
var Dagger$1 = "‡";
var Darr = "↡";
var Dashv = "⫤";
var Dcaron = "Ď";
var Dcy = "Д";
var Del = "∇";
var Delta$1 = "Δ";
var Dfr = "𝔇";
var DiacriticalAcute = "´";
var DiacriticalDot = "˙";
var DiacriticalDoubleAcute = "˝";
var DiacriticalGrave = "`";
var DiacriticalTilde = "˜";
var Diamond = "⋄";
var DifferentialD = "ⅆ";
var Dopf = "𝔻";
var Dot = "¨";
var DotDot = "⃜";
var DotEqual = "≐";
var DoubleContourIntegral = "∯";
var DoubleDot = "¨";
var DoubleDownArrow = "⇓";
var DoubleLeftArrow = "⇐";
var DoubleLeftRightArrow = "⇔";
var DoubleLeftTee = "⫤";
var DoubleLongLeftArrow = "⟸";
var DoubleLongLeftRightArrow = "⟺";
var DoubleLongRightArrow = "⟹";
var DoubleRightArrow = "⇒";
var DoubleRightTee = "⊨";
var DoubleUpArrow = "⇑";
var DoubleUpDownArrow = "⇕";
var DoubleVerticalBar = "∥";
var DownArrow = "↓";
var DownArrowBar = "⤓";
var DownArrowUpArrow = "⇵";
var DownBreve = "̑";
var DownLeftRightVector = "⥐";
var DownLeftTeeVector = "⥞";
var DownLeftVector = "↽";
var DownLeftVectorBar = "⥖";
var DownRightTeeVector = "⥟";
var DownRightVector = "⇁";
var DownRightVectorBar = "⥗";
var DownTee = "⊤";
var DownTeeArrow = "↧";
var Downarrow = "⇓";
var Dscr = "𝒟";
var Dstrok = "Đ";
var ENG = "Ŋ";
var ET = "Ð";
var ETH$1 = "Ð";
var Eacut = "É";
var Eacute$1 = "É";
var Ecaron = "Ě";
var Ecir = "Ê";
var Ecirc$1 = "Ê";
var Ecy = "Э";
var Edot = "Ė";
var Efr = "𝔈";
var Egrav = "È";
var Egrave$1 = "È";
var Element = "∈";
var Emacr = "Ē";
var EmptySmallSquare = "◻";
var EmptyVerySmallSquare = "▫";
var Eogon = "Ę";
var Eopf = "𝔼";
var Epsilon$1 = "Ε";
var Equal = "⩵";
var EqualTilde = "≂";
var Equilibrium = "⇌";
var Escr = "ℰ";
var Esim = "⩳";
var Eta$1 = "Η";
var Eum = "Ë";
var Euml$1 = "Ë";
var Exists = "∃";
var ExponentialE = "ⅇ";
var Fcy = "Ф";
var Ffr = "𝔉";
var FilledSmallSquare = "◼";
var FilledVerySmallSquare = "▪";
var Fopf = "𝔽";
var ForAll = "∀";
var Fouriertrf = "ℱ";
var Fscr = "ℱ";
var GJcy = "Ѓ";
var G = ">";
var GT$1 = ">";
var Gamma$1 = "Γ";
var Gammad = "Ϝ";
var Gbreve = "Ğ";
var Gcedil = "Ģ";
var Gcirc = "Ĝ";
var Gcy = "Г";
var Gdot = "Ġ";
var Gfr = "𝔊";
var Gg = "⋙";
var Gopf = "𝔾";
var GreaterEqual = "≥";
var GreaterEqualLess = "⋛";
var GreaterFullEqual = "≧";
var GreaterGreater = "⪢";
var GreaterLess = "≷";
var GreaterSlantEqual = "⩾";
var GreaterTilde = "≳";
var Gscr = "𝒢";
var Gt = "≫";
var HARDcy = "Ъ";
var Hacek = "ˇ";
var Hat = "^";
var Hcirc = "Ĥ";
var Hfr = "ℌ";
var HilbertSpace = "ℋ";
var Hopf = "ℍ";
var HorizontalLine = "─";
var Hscr = "ℋ";
var Hstrok = "Ħ";
var HumpDownHump = "≎";
var HumpEqual = "≏";
var IEcy = "Е";
var IJlig = "Ĳ";
var IOcy = "Ё";
var Iacut = "Í";
var Iacute$1 = "Í";
var Icir = "Î";
var Icirc$1 = "Î";
var Icy = "И";
var Idot = "İ";
var Ifr = "ℑ";
var Igrav = "Ì";
var Igrave$1 = "Ì";
var Im = "ℑ";
var Imacr = "Ī";
var ImaginaryI = "ⅈ";
var Implies = "⇒";
var Int = "∬";
var Integral = "∫";
var Intersection = "⋂";
var InvisibleComma = "⁣";
var InvisibleTimes = "⁢";
var Iogon = "Į";
var Iopf = "𝕀";
var Iota$1 = "Ι";
var Iscr = "ℐ";
var Itilde = "Ĩ";
var Iukcy = "І";
var Ium = "Ï";
var Iuml$1 = "Ï";
var Jcirc = "Ĵ";
var Jcy = "Й";
var Jfr = "𝔍";
var Jopf = "𝕁";
var Jscr = "𝒥";
var Jsercy = "Ј";
var Jukcy = "Є";
var KHcy = "Х";
var KJcy = "Ќ";
var Kappa$1 = "Κ";
var Kcedil = "Ķ";
var Kcy = "К";
var Kfr = "𝔎";
var Kopf = "𝕂";
var Kscr = "𝒦";
var LJcy = "Љ";
var L = "<";
var LT$1 = "<";
var Lacute = "Ĺ";
var Lambda$1 = "Λ";
var Lang = "⟪";
var Laplacetrf = "ℒ";
var Larr = "↞";
var Lcaron = "Ľ";
var Lcedil = "Ļ";
var Lcy = "Л";
var LeftAngleBracket = "⟨";
var LeftArrow = "←";
var LeftArrowBar = "⇤";
var LeftArrowRightArrow = "⇆";
var LeftCeiling = "⌈";
var LeftDoubleBracket = "⟦";
var LeftDownTeeVector = "⥡";
var LeftDownVector = "⇃";
var LeftDownVectorBar = "⥙";
var LeftFloor = "⌊";
var LeftRightArrow = "↔";
var LeftRightVector = "⥎";
var LeftTee = "⊣";
var LeftTeeArrow = "↤";
var LeftTeeVector = "⥚";
var LeftTriangle = "⊲";
var LeftTriangleBar = "⧏";
var LeftTriangleEqual = "⊴";
var LeftUpDownVector = "⥑";
var LeftUpTeeVector = "⥠";
var LeftUpVector = "↿";
var LeftUpVectorBar = "⥘";
var LeftVector = "↼";
var LeftVectorBar = "⥒";
var Leftarrow = "⇐";
var Leftrightarrow = "⇔";
var LessEqualGreater = "⋚";
var LessFullEqual = "≦";
var LessGreater = "≶";
var LessLess = "⪡";
var LessSlantEqual = "⩽";
var LessTilde = "≲";
var Lfr = "𝔏";
var Ll = "⋘";
var Lleftarrow = "⇚";
var Lmidot = "Ŀ";
var LongLeftArrow = "⟵";
var LongLeftRightArrow = "⟷";
var LongRightArrow = "⟶";
var Longleftarrow = "⟸";
var Longleftrightarrow = "⟺";
var Longrightarrow = "⟹";
var Lopf = "𝕃";
var LowerLeftArrow = "↙";
var LowerRightArrow = "↘";
var Lscr = "ℒ";
var Lsh = "↰";
var Lstrok = "Ł";
var Lt = "≪";
var Mcy = "М";
var MediumSpace = " ";
var Mellintrf = "ℳ";
var Mfr = "𝔐";
var MinusPlus = "∓";
var Mopf = "𝕄";
var Mscr = "ℳ";
var Mu$1 = "Μ";
var NJcy = "Њ";
var Nacute = "Ń";
var Ncaron = "Ň";
var Ncedil = "Ņ";
var Ncy = "Н";
var NegativeMediumSpace = "​";
var NegativeThickSpace = "​";
var NegativeThinSpace = "​";
var NegativeVeryThinSpace = "​";
var NestedGreaterGreater = "≫";
var NestedLessLess = "≪";
var NewLine = "\n";
var Nfr = "𝔑";
var NoBreak = "⁠";
var NonBreakingSpace = " ";
var Nopf = "ℕ";
var Not = "⫬";
var NotCongruent = "≢";
var NotCupCap = "≭";
var NotDoubleVerticalBar = "∦";
var NotElement = "∉";
var NotEqual = "≠";
var NotEqualTilde = "≂̸";
var NotExists = "∄";
var NotGreater = "≯";
var NotGreaterEqual = "≱";
var NotGreaterFullEqual = "≧̸";
var NotGreaterGreater = "≫̸";
var NotGreaterLess = "≹";
var NotGreaterSlantEqual = "⩾̸";
var NotGreaterTilde = "≵";
var NotHumpDownHump = "≎̸";
var NotHumpEqual = "≏̸";
var NotLeftTriangle = "⋪";
var NotLeftTriangleBar = "⧏̸";
var NotLeftTriangleEqual = "⋬";
var NotLess = "≮";
var NotLessEqual = "≰";
var NotLessGreater = "≸";
var NotLessLess = "≪̸";
var NotLessSlantEqual = "⩽̸";
var NotLessTilde = "≴";
var NotNestedGreaterGreater = "⪢̸";
var NotNestedLessLess = "⪡̸";
var NotPrecedes = "⊀";
var NotPrecedesEqual = "⪯̸";
var NotPrecedesSlantEqual = "⋠";
var NotReverseElement = "∌";
var NotRightTriangle = "⋫";
var NotRightTriangleBar = "⧐̸";
var NotRightTriangleEqual = "⋭";
var NotSquareSubset = "⊏̸";
var NotSquareSubsetEqual = "⋢";
var NotSquareSuperset = "⊐̸";
var NotSquareSupersetEqual = "⋣";
var NotSubset = "⊂⃒";
var NotSubsetEqual = "⊈";
var NotSucceeds = "⊁";
var NotSucceedsEqual = "⪰̸";
var NotSucceedsSlantEqual = "⋡";
var NotSucceedsTilde = "≿̸";
var NotSuperset = "⊃⃒";
var NotSupersetEqual = "⊉";
var NotTilde = "≁";
var NotTildeEqual = "≄";
var NotTildeFullEqual = "≇";
var NotTildeTilde = "≉";
var NotVerticalBar = "∤";
var Nscr = "𝒩";
var Ntild = "Ñ";
var Ntilde$1 = "Ñ";
var Nu$1 = "Ν";
var OElig$1 = "Œ";
var Oacut = "Ó";
var Oacute$1 = "Ó";
var Ocir = "Ô";
var Ocirc$1 = "Ô";
var Ocy = "О";
var Odblac = "Ő";
var Ofr = "𝔒";
var Ograv = "Ò";
var Ograve$1 = "Ò";
var Omacr = "Ō";
var Omega$1 = "Ω";
var Omicron$1 = "Ο";
var Oopf = "𝕆";
var OpenCurlyDoubleQuote = "“";
var OpenCurlyQuote = "‘";
var Or = "⩔";
var Oscr = "𝒪";
var Oslas = "Ø";
var Oslash$1 = "Ø";
var Otild = "Õ";
var Otilde$1 = "Õ";
var Otimes = "⨷";
var Oum = "Ö";
var Ouml$1 = "Ö";
var OverBar = "‾";
var OverBrace = "⏞";
var OverBracket = "⎴";
var OverParenthesis = "⏜";
var PartialD = "∂";
var Pcy = "П";
var Pfr = "𝔓";
var Phi$1 = "Φ";
var Pi$1 = "Π";
var PlusMinus = "±";
var Poincareplane = "ℌ";
var Popf = "ℙ";
var Pr = "⪻";
var Precedes = "≺";
var PrecedesEqual = "⪯";
var PrecedesSlantEqual = "≼";
var PrecedesTilde = "≾";
var Prime$1 = "″";
var Product = "∏";
var Proportion = "∷";
var Proportional = "∝";
var Pscr = "𝒫";
var Psi$1 = "Ψ";
var QUO = "\"";
var QUOT = "\"";
var Qfr = "𝔔";
var Qopf = "ℚ";
var Qscr = "𝒬";
var RBarr = "⤐";
var RE = "®";
var REG = "®";
var Racute = "Ŕ";
var Rang = "⟫";
var Rarr = "↠";
var Rarrtl = "⤖";
var Rcaron = "Ř";
var Rcedil = "Ŗ";
var Rcy = "Р";
var Re = "ℜ";
var ReverseElement = "∋";
var ReverseEquilibrium = "⇋";
var ReverseUpEquilibrium = "⥯";
var Rfr = "ℜ";
var Rho$1 = "Ρ";
var RightAngleBracket = "⟩";
var RightArrow = "→";
var RightArrowBar = "⇥";
var RightArrowLeftArrow = "⇄";
var RightCeiling = "⌉";
var RightDoubleBracket = "⟧";
var RightDownTeeVector = "⥝";
var RightDownVector = "⇂";
var RightDownVectorBar = "⥕";
var RightFloor = "⌋";
var RightTee = "⊢";
var RightTeeArrow = "↦";
var RightTeeVector = "⥛";
var RightTriangle = "⊳";
var RightTriangleBar = "⧐";
var RightTriangleEqual = "⊵";
var RightUpDownVector = "⥏";
var RightUpTeeVector = "⥜";
var RightUpVector = "↾";
var RightUpVectorBar = "⥔";
var RightVector = "⇀";
var RightVectorBar = "⥓";
var Rightarrow = "⇒";
var Ropf = "ℝ";
var RoundImplies = "⥰";
var Rrightarrow = "⇛";
var Rscr = "ℛ";
var Rsh = "↱";
var RuleDelayed = "⧴";
var SHCHcy = "Щ";
var SHcy = "Ш";
var SOFTcy = "Ь";
var Sacute = "Ś";
var Sc = "⪼";
var Scaron$1 = "Š";
var Scedil = "Ş";
var Scirc = "Ŝ";
var Scy = "С";
var Sfr = "𝔖";
var ShortDownArrow = "↓";
var ShortLeftArrow = "←";
var ShortRightArrow = "→";
var ShortUpArrow = "↑";
var Sigma$1 = "Σ";
var SmallCircle = "∘";
var Sopf = "𝕊";
var Sqrt = "√";
var Square = "□";
var SquareIntersection = "⊓";
var SquareSubset = "⊏";
var SquareSubsetEqual = "⊑";
var SquareSuperset = "⊐";
var SquareSupersetEqual = "⊒";
var SquareUnion = "⊔";
var Sscr = "𝒮";
var Star = "⋆";
var Sub = "⋐";
var Subset = "⋐";
var SubsetEqual = "⊆";
var Succeeds = "≻";
var SucceedsEqual = "⪰";
var SucceedsSlantEqual = "≽";
var SucceedsTilde = "≿";
var SuchThat = "∋";
var Sum = "∑";
var Sup = "⋑";
var Superset = "⊃";
var SupersetEqual = "⊇";
var Supset = "⋑";
var THOR = "Þ";
var THORN$1 = "Þ";
var TRADE = "™";
var TSHcy = "Ћ";
var TScy = "Ц";
var Tab = "\t";
var Tau$1 = "Τ";
var Tcaron = "Ť";
var Tcedil = "Ţ";
var Tcy = "Т";
var Tfr = "𝔗";
var Therefore = "∴";
var Theta$1 = "Θ";
var ThickSpace = "  ";
var ThinSpace = " ";
var Tilde = "∼";
var TildeEqual = "≃";
var TildeFullEqual = "≅";
var TildeTilde = "≈";
var Topf = "𝕋";
var TripleDot = "⃛";
var Tscr = "𝒯";
var Tstrok = "Ŧ";
var Uacut = "Ú";
var Uacute$1 = "Ú";
var Uarr = "↟";
var Uarrocir = "⥉";
var Ubrcy = "Ў";
var Ubreve = "Ŭ";
var Ucir = "Û";
var Ucirc$1 = "Û";
var Ucy = "У";
var Udblac = "Ű";
var Ufr = "𝔘";
var Ugrav = "Ù";
var Ugrave$1 = "Ù";
var Umacr = "Ū";
var UnderBar = "_";
var UnderBrace = "⏟";
var UnderBracket = "⎵";
var UnderParenthesis = "⏝";
var Union = "⋃";
var UnionPlus = "⊎";
var Uogon = "Ų";
var Uopf = "𝕌";
var UpArrow = "↑";
var UpArrowBar = "⤒";
var UpArrowDownArrow = "⇅";
var UpDownArrow = "↕";
var UpEquilibrium = "⥮";
var UpTee = "⊥";
var UpTeeArrow = "↥";
var Uparrow = "⇑";
var Updownarrow = "⇕";
var UpperLeftArrow = "↖";
var UpperRightArrow = "↗";
var Upsi = "ϒ";
var Upsilon$1 = "Υ";
var Uring = "Ů";
var Uscr = "𝒰";
var Utilde = "Ũ";
var Uum = "Ü";
var Uuml$1 = "Ü";
var VDash = "⊫";
var Vbar = "⫫";
var Vcy = "В";
var Vdash = "⊩";
var Vdashl = "⫦";
var Vee = "⋁";
var Verbar = "‖";
var Vert = "‖";
var VerticalBar = "∣";
var VerticalLine = "|";
var VerticalSeparator = "❘";
var VerticalTilde = "≀";
var VeryThinSpace = " ";
var Vfr = "𝔙";
var Vopf = "𝕍";
var Vscr = "𝒱";
var Vvdash = "⊪";
var Wcirc = "Ŵ";
var Wedge = "⋀";
var Wfr = "𝔚";
var Wopf = "𝕎";
var Wscr = "𝒲";
var Xfr = "𝔛";
var Xi$1 = "Ξ";
var Xopf = "𝕏";
var Xscr = "𝒳";
var YAcy = "Я";
var YIcy = "Ї";
var YUcy = "Ю";
var Yacut = "Ý";
var Yacute$1 = "Ý";
var Ycirc = "Ŷ";
var Ycy = "Ы";
var Yfr = "𝔜";
var Yopf = "𝕐";
var Yscr = "𝒴";
var Yuml$1 = "Ÿ";
var ZHcy = "Ж";
var Zacute = "Ź";
var Zcaron = "Ž";
var Zcy = "З";
var Zdot = "Ż";
var ZeroWidthSpace = "​";
var Zeta$1 = "Ζ";
var Zfr = "ℨ";
var Zopf = "ℤ";
var Zscr = "𝒵";
var aacut = "á";
var aacute$1 = "á";
var abreve = "ă";
var ac = "∾";
var acE = "∾̳";
var acd = "∿";
var acir = "â";
var acirc$1 = "â";
var acut = "´";
var acute$1 = "´";
var acy = "а";
var aeli = "æ";
var aelig$1 = "æ";
var af = "⁡";
var afr = "𝔞";
var agrav = "à";
var agrave$1 = "à";
var alefsym$1 = "ℵ";
var aleph = "ℵ";
var alpha$1 = "α";
var amacr = "ā";
var amalg = "⨿";
var am = "&";
var amp$1 = "&";
var and$1 = "∧";
var andand = "⩕";
var andd = "⩜";
var andslope = "⩘";
var andv = "⩚";
var ang$1 = "∠";
var ange = "⦤";
var angle = "∠";
var angmsd = "∡";
var angmsdaa = "⦨";
var angmsdab = "⦩";
var angmsdac = "⦪";
var angmsdad = "⦫";
var angmsdae = "⦬";
var angmsdaf = "⦭";
var angmsdag = "⦮";
var angmsdah = "⦯";
var angrt = "∟";
var angrtvb = "⊾";
var angrtvbd = "⦝";
var angsph = "∢";
var angst = "Å";
var angzarr = "⍼";
var aogon = "ą";
var aopf = "𝕒";
var ap = "≈";
var apE = "⩰";
var apacir = "⩯";
var ape = "≊";
var apid = "≋";
var apos = "'";
var approx = "≈";
var approxeq = "≊";
var arin = "å";
var aring$1 = "å";
var ascr = "𝒶";
var ast = "*";
var asymp$1 = "≈";
var asympeq = "≍";
var atild = "ã";
var atilde$1 = "ã";
var aum = "ä";
var auml$1 = "ä";
var awconint = "∳";
var awint = "⨑";
var bNot = "⫭";
var backcong = "≌";
var backepsilon = "϶";
var backprime = "‵";
var backsim = "∽";
var backsimeq = "⋍";
var barvee = "⊽";
var barwed = "⌅";
var barwedge = "⌅";
var bbrk = "⎵";
var bbrktbrk = "⎶";
var bcong = "≌";
var bcy = "б";
var bdquo$1 = "„";
var becaus = "∵";
var because = "∵";
var bemptyv = "⦰";
var bepsi = "϶";
var bernou = "ℬ";
var beta$1 = "β";
var beth = "ℶ";
var between = "≬";
var bfr = "𝔟";
var bigcap = "⋂";
var bigcirc = "◯";
var bigcup = "⋃";
var bigodot = "⨀";
var bigoplus = "⨁";
var bigotimes = "⨂";
var bigsqcup = "⨆";
var bigstar = "★";
var bigtriangledown = "▽";
var bigtriangleup = "△";
var biguplus = "⨄";
var bigvee = "⋁";
var bigwedge = "⋀";
var bkarow = "⤍";
var blacklozenge = "⧫";
var blacksquare = "▪";
var blacktriangle = "▴";
var blacktriangledown = "▾";
var blacktriangleleft = "◂";
var blacktriangleright = "▸";
var blank = "␣";
var blk12 = "▒";
var blk14 = "░";
var blk34 = "▓";
var block = "█";
var bne = "=⃥";
var bnequiv = "≡⃥";
var bnot = "⌐";
var bopf = "𝕓";
var bot = "⊥";
var bottom = "⊥";
var bowtie = "⋈";
var boxDL = "╗";
var boxDR = "╔";
var boxDl = "╖";
var boxDr = "╓";
var boxH = "═";
var boxHD = "╦";
var boxHU = "╩";
var boxHd = "╤";
var boxHu = "╧";
var boxUL = "╝";
var boxUR = "╚";
var boxUl = "╜";
var boxUr = "╙";
var boxV = "║";
var boxVH = "╬";
var boxVL = "╣";
var boxVR = "╠";
var boxVh = "╫";
var boxVl = "╢";
var boxVr = "╟";
var boxbox = "⧉";
var boxdL = "╕";
var boxdR = "╒";
var boxdl = "┐";
var boxdr = "┌";
var boxh = "─";
var boxhD = "╥";
var boxhU = "╨";
var boxhd = "┬";
var boxhu = "┴";
var boxminus = "⊟";
var boxplus = "⊞";
var boxtimes = "⊠";
var boxuL = "╛";
var boxuR = "╘";
var boxul = "┘";
var boxur = "└";
var boxv = "│";
var boxvH = "╪";
var boxvL = "╡";
var boxvR = "╞";
var boxvh = "┼";
var boxvl = "┤";
var boxvr = "├";
var bprime = "‵";
var breve = "˘";
var brvba = "¦";
var brvbar$1 = "¦";
var bscr = "𝒷";
var bsemi = "⁏";
var bsim = "∽";
var bsime = "⋍";
var bsol = "\\";
var bsolb = "⧅";
var bsolhsub = "⟈";
var bull$1 = "•";
var bullet = "•";
var bump = "≎";
var bumpE = "⪮";
var bumpe = "≏";
var bumpeq = "≏";
var cacute = "ć";
var cap$2 = "∩";
var capand = "⩄";
var capbrcup = "⩉";
var capcap = "⩋";
var capcup = "⩇";
var capdot = "⩀";
var caps = "∩︀";
var caret = "⁁";
var caron = "ˇ";
var ccaps = "⩍";
var ccaron = "č";
var ccedi = "ç";
var ccedil$1 = "ç";
var ccirc = "ĉ";
var ccups = "⩌";
var ccupssm = "⩐";
var cdot = "ċ";
var cedi = "¸";
var cedil$1 = "¸";
var cemptyv = "⦲";
var cen = "¢";
var cent$1 = "¢";
var centerdot = "·";
var cfr = "𝔠";
var chcy = "ч";
var check = "✓";
var checkmark = "✓";
var chi$1 = "χ";
var cir = "○";
var cirE = "⧃";
var circ$1 = "ˆ";
var circeq = "≗";
var circlearrowleft = "↺";
var circlearrowright = "↻";
var circledR = "®";
var circledS = "Ⓢ";
var circledast = "⊛";
var circledcirc = "⊚";
var circleddash = "⊝";
var cire = "≗";
var cirfnint = "⨐";
var cirmid = "⫯";
var cirscir = "⧂";
var clubs$1 = "♣";
var clubsuit = "♣";
var colon = ":";
var colone = "≔";
var coloneq = "≔";
var comma$1 = ",";
var commat = "@";
var comp = "∁";
var compfn = "∘";
var complement = "∁";
var complexes = "ℂ";
var cong$1 = "≅";
var congdot = "⩭";
var conint = "∮";
var copf = "𝕔";
var coprod = "∐";
var cop = "©";
var copy$1 = "©";
var copysr = "℗";
var crarr$1 = "↵";
var cross = "✗";
var cscr = "𝒸";
var csub = "⫏";
var csube = "⫑";
var csup = "⫐";
var csupe = "⫒";
var ctdot = "⋯";
var cudarrl = "⤸";
var cudarrr = "⤵";
var cuepr = "⋞";
var cuesc = "⋟";
var cularr = "↶";
var cularrp = "⤽";
var cup$1 = "∪";
var cupbrcap = "⩈";
var cupcap = "⩆";
var cupcup = "⩊";
var cupdot = "⊍";
var cupor = "⩅";
var cups = "∪︀";
var curarr = "↷";
var curarrm = "⤼";
var curlyeqprec = "⋞";
var curlyeqsucc = "⋟";
var curlyvee = "⋎";
var curlywedge = "⋏";
var curre = "¤";
var curren$1 = "¤";
var curvearrowleft = "↶";
var curvearrowright = "↷";
var cuvee = "⋎";
var cuwed = "⋏";
var cwconint = "∲";
var cwint = "∱";
var cylcty = "⌭";
var dArr$1 = "⇓";
var dHar = "⥥";
var dagger$1 = "†";
var daleth = "ℸ";
var darr$1 = "↓";
var dash$1 = "‐";
var dashv = "⊣";
var dbkarow = "⤏";
var dblac = "˝";
var dcaron = "ď";
var dcy = "д";
var dd$1 = "ⅆ";
var ddagger = "‡";
var ddarr = "⇊";
var ddotseq = "⩷";
var de = "°";
var deg$1 = "°";
var delta$1 = "δ";
var demptyv = "⦱";
var dfisht = "⥿";
var dfr = "𝔡";
var dharl = "⇃";
var dharr = "⇂";
var diam = "⋄";
var diamond = "⋄";
var diamondsuit = "♦";
var diams$1 = "♦";
var die = "¨";
var digamma = "ϝ";
var disin = "⋲";
var div = "÷";
var divid = "÷";
var divide$1 = "÷";
var divideontimes = "⋇";
var divonx = "⋇";
var djcy = "ђ";
var dlcorn = "⌞";
var dlcrop = "⌍";
var dollar = "$";
var dopf = "𝕕";
var dot = "˙";
var doteq = "≐";
var doteqdot = "≑";
var dotminus = "∸";
var dotplus = "∔";
var dotsquare = "⊡";
var doublebarwedge = "⌆";
var downarrow = "↓";
var downdownarrows = "⇊";
var downharpoonleft = "⇃";
var downharpoonright = "⇂";
var drbkarow = "⤐";
var drcorn = "⌟";
var drcrop = "⌌";
var dscr = "𝒹";
var dscy = "ѕ";
var dsol = "⧶";
var dstrok = "đ";
var dtdot = "⋱";
var dtri = "▿";
var dtrif = "▾";
var duarr = "⇵";
var duhar = "⥯";
var dwangle = "⦦";
var dzcy = "џ";
var dzigrarr = "⟿";
var eDDot = "⩷";
var eDot = "≑";
var eacut = "é";
var eacute$1 = "é";
var easter = "⩮";
var ecaron = "ě";
var ecir = "ê";
var ecirc$1 = "ê";
var ecolon = "≕";
var ecy = "э";
var edot = "ė";
var ee = "ⅇ";
var efDot = "≒";
var efr = "𝔢";
var eg = "⪚";
var egrav = "è";
var egrave$1 = "è";
var egs = "⪖";
var egsdot = "⪘";
var el = "⪙";
var elinters = "⏧";
var ell = "ℓ";
var els = "⪕";
var elsdot = "⪗";
var emacr = "ē";
var empty$3 = "∅";
var emptyset = "∅";
var emptyv = "∅";
var emsp13 = " ";
var emsp14 = " ";
var emsp$1 = " ";
var eng = "ŋ";
var ensp$1 = " ";
var eogon = "ę";
var eopf = "𝕖";
var epar = "⋕";
var eparsl = "⧣";
var eplus = "⩱";
var epsi = "ε";
var epsilon$1 = "ε";
var epsiv = "ϵ";
var eqcirc = "≖";
var eqcolon = "≕";
var eqsim = "≂";
var eqslantgtr = "⪖";
var eqslantless = "⪕";
var equals = "=";
var equest = "≟";
var equiv$1 = "≡";
var equivDD = "⩸";
var eqvparsl = "⧥";
var erDot = "≓";
var erarr = "⥱";
var escr = "ℯ";
var esdot = "≐";
var esim = "≂";
var eta$1 = "η";
var et = "ð";
var eth$1 = "ð";
var eum = "ë";
var euml$1 = "ë";
var euro$1 = "€";
var excl = "!";
var exist$1 = "∃";
var expectation = "ℰ";
var exponentiale = "ⅇ";
var fallingdotseq = "≒";
var fcy = "ф";
var female = "♀";
var ffilig = "ﬃ";
var fflig = "ﬀ";
var ffllig = "ﬄ";
var ffr = "𝔣";
var filig = "ﬁ";
var fjlig = "fj";
var flat = "♭";
var fllig = "ﬂ";
var fltns = "▱";
var fnof$1 = "ƒ";
var fopf = "𝕗";
var forall$1 = "∀";
var fork = "⋔";
var forkv = "⫙";
var fpartint = "⨍";
var frac1 = "¼";
var frac12$1 = "½";
var frac13 = "⅓";
var frac14$1 = "¼";
var frac15 = "⅕";
var frac16 = "⅙";
var frac18 = "⅛";
var frac23 = "⅔";
var frac25 = "⅖";
var frac3 = "¾";
var frac34$1 = "¾";
var frac35 = "⅗";
var frac38 = "⅜";
var frac45 = "⅘";
var frac56 = "⅚";
var frac58 = "⅝";
var frac78 = "⅞";
var frasl$1 = "⁄";
var frown = "⌢";
var fscr = "𝒻";
var gE = "≧";
var gEl = "⪌";
var gacute = "ǵ";
var gamma$1 = "γ";
var gammad = "ϝ";
var gap = "⪆";
var gbreve = "ğ";
var gcirc = "ĝ";
var gcy = "г";
var gdot = "ġ";
var ge$1 = "≥";
var gel = "⋛";
var geq = "≥";
var geqq = "≧";
var geqslant = "⩾";
var ges = "⩾";
var gescc = "⪩";
var gesdot = "⪀";
var gesdoto = "⪂";
var gesdotol = "⪄";
var gesl = "⋛︀";
var gesles = "⪔";
var gfr = "𝔤";
var gg = "≫";
var ggg = "⋙";
var gimel = "ℷ";
var gjcy = "ѓ";
var gl = "≷";
var glE = "⪒";
var gla = "⪥";
var glj = "⪤";
var gnE = "≩";
var gnap = "⪊";
var gnapprox = "⪊";
var gne = "⪈";
var gneq = "⪈";
var gneqq = "≩";
var gnsim = "⋧";
var gopf = "𝕘";
var grave = "`";
var gscr = "ℊ";
var gsim = "≳";
var gsime = "⪎";
var gsiml = "⪐";
var g = ">";
var gt$1 = ">";
var gtcc = "⪧";
var gtcir = "⩺";
var gtdot = "⋗";
var gtlPar = "⦕";
var gtquest = "⩼";
var gtrapprox = "⪆";
var gtrarr = "⥸";
var gtrdot = "⋗";
var gtreqless = "⋛";
var gtreqqless = "⪌";
var gtrless = "≷";
var gtrsim = "≳";
var gvertneqq = "≩︀";
var gvnE = "≩︀";
var hArr$1 = "⇔";
var hairsp = " ";
var half = "½";
var hamilt = "ℋ";
var hardcy = "ъ";
var harr$1 = "↔";
var harrcir = "⥈";
var harrw = "↭";
var hbar = "ℏ";
var hcirc = "ĥ";
var hearts$1 = "♥";
var heartsuit = "♥";
var hellip$1 = "…";
var hercon = "⊹";
var hfr = "𝔥";
var hksearow = "⤥";
var hkswarow = "⤦";
var hoarr = "⇿";
var homtht = "∻";
var hookleftarrow = "↩";
var hookrightarrow = "↪";
var hopf = "𝕙";
var horbar = "―";
var hscr = "𝒽";
var hslash = "ℏ";
var hstrok = "ħ";
var hybull = "⁃";
var hyphen = "‐";
var iacut = "í";
var iacute$1 = "í";
var ic = "⁣";
var icir = "î";
var icirc$1 = "î";
var icy = "и";
var iecy = "е";
var iexc = "¡";
var iexcl$1 = "¡";
var iff = "⇔";
var ifr = "𝔦";
var igrav = "ì";
var igrave$1 = "ì";
var ii = "ⅈ";
var iiiint = "⨌";
var iiint = "∭";
var iinfin = "⧜";
var iiota = "℩";
var ijlig = "ĳ";
var imacr = "ī";
var image$2 = "ℑ";
var imagline = "ℐ";
var imagpart = "ℑ";
var imath = "ı";
var imof = "⊷";
var imped = "Ƶ";
var incare = "℅";
var infin$1 = "∞";
var infintie = "⧝";
var inodot = "ı";
var int$2 = "∫";
var intcal = "⊺";
var integers = "ℤ";
var intercal = "⊺";
var intlarhk = "⨗";
var intprod = "⨼";
var iocy = "ё";
var iogon = "į";
var iopf = "𝕚";
var iota$1 = "ι";
var iprod = "⨼";
var iques = "¿";
var iquest$1 = "¿";
var iscr = "𝒾";
var isin$1 = "∈";
var isinE = "⋹";
var isindot = "⋵";
var isins = "⋴";
var isinsv = "⋳";
var isinv = "∈";
var it = "⁢";
var itilde = "ĩ";
var iukcy = "і";
var ium = "ï";
var iuml$1 = "ï";
var jcirc = "ĵ";
var jcy = "й";
var jfr = "𝔧";
var jmath = "ȷ";
var jopf = "𝕛";
var jscr = "𝒿";
var jsercy = "ј";
var jukcy = "є";
var kappa$1 = "κ";
var kappav = "ϰ";
var kcedil = "ķ";
var kcy = "к";
var kfr = "𝔨";
var kgreen = "ĸ";
var khcy = "х";
var kjcy = "ќ";
var kopf = "𝕜";
var kscr = "𝓀";
var lAarr = "⇚";
var lArr$1 = "⇐";
var lAtail = "⤛";
var lBarr = "⤎";
var lE = "≦";
var lEg = "⪋";
var lHar = "⥢";
var lacute = "ĺ";
var laemptyv = "⦴";
var lagran = "ℒ";
var lambda$1 = "λ";
var lang$1 = "⟨";
var langd = "⦑";
var langle = "⟨";
var lap = "⪅";
var laqu = "«";
var laquo$1 = "«";
var larr$1 = "←";
var larrb = "⇤";
var larrbfs = "⤟";
var larrfs = "⤝";
var larrhk = "↩";
var larrlp = "↫";
var larrpl = "⤹";
var larrsim = "⥳";
var larrtl = "↢";
var lat = "⪫";
var latail = "⤙";
var late = "⪭";
var lates = "⪭︀";
var lbarr = "⤌";
var lbbrk = "❲";
var lbrace = "{";
var lbrack = "[";
var lbrke = "⦋";
var lbrksld = "⦏";
var lbrkslu = "⦍";
var lcaron = "ľ";
var lcedil = "ļ";
var lceil$1 = "⌈";
var lcub = "{";
var lcy = "л";
var ldca = "⤶";
var ldquo$1 = "“";
var ldquor = "„";
var ldrdhar = "⥧";
var ldrushar = "⥋";
var ldsh = "↲";
var le$1 = "≤";
var leftarrow = "←";
var leftarrowtail = "↢";
var leftharpoondown = "↽";
var leftharpoonup = "↼";
var leftleftarrows = "⇇";
var leftrightarrow = "↔";
var leftrightarrows = "⇆";
var leftrightharpoons = "⇋";
var leftrightsquigarrow = "↭";
var leftthreetimes = "⋋";
var leg = "⋚";
var leq = "≤";
var leqq = "≦";
var leqslant = "⩽";
var les = "⩽";
var lescc = "⪨";
var lesdot = "⩿";
var lesdoto = "⪁";
var lesdotor = "⪃";
var lesg = "⋚︀";
var lesges = "⪓";
var lessapprox = "⪅";
var lessdot = "⋖";
var lesseqgtr = "⋚";
var lesseqqgtr = "⪋";
var lessgtr = "≶";
var lesssim = "≲";
var lfisht = "⥼";
var lfloor$1 = "⌊";
var lfr = "𝔩";
var lg = "≶";
var lgE = "⪑";
var lhard = "↽";
var lharu = "↼";
var lharul = "⥪";
var lhblk = "▄";
var ljcy = "љ";
var ll = "≪";
var llarr = "⇇";
var llcorner = "⌞";
var llhard = "⥫";
var lltri = "◺";
var lmidot = "ŀ";
var lmoust = "⎰";
var lmoustache = "⎰";
var lnE = "≨";
var lnap = "⪉";
var lnapprox = "⪉";
var lne = "⪇";
var lneq = "⪇";
var lneqq = "≨";
var lnsim = "⋦";
var loang = "⟬";
var loarr = "⇽";
var lobrk = "⟦";
var longleftarrow = "⟵";
var longleftrightarrow = "⟷";
var longmapsto = "⟼";
var longrightarrow = "⟶";
var looparrowleft = "↫";
var looparrowright = "↬";
var lopar = "⦅";
var lopf = "𝕝";
var loplus = "⨭";
var lotimes = "⨴";
var lowast$1 = "∗";
var lowbar = "_";
var loz$1 = "◊";
var lozenge = "◊";
var lozf = "⧫";
var lpar = "(";
var lparlt = "⦓";
var lrarr = "⇆";
var lrcorner = "⌟";
var lrhar = "⇋";
var lrhard = "⥭";
var lrm$1 = "‎";
var lrtri = "⊿";
var lsaquo$1 = "‹";
var lscr = "𝓁";
var lsh = "↰";
var lsim = "≲";
var lsime = "⪍";
var lsimg = "⪏";
var lsqb = "[";
var lsquo$1 = "‘";
var lsquor = "‚";
var lstrok = "ł";
var l = "<";
var lt$1 = "<";
var ltcc = "⪦";
var ltcir = "⩹";
var ltdot = "⋖";
var lthree = "⋋";
var ltimes = "⋉";
var ltlarr = "⥶";
var ltquest = "⩻";
var ltrPar = "⦖";
var ltri = "◃";
var ltrie = "⊴";
var ltrif = "◂";
var lurdshar = "⥊";
var luruhar = "⥦";
var lvertneqq = "≨︀";
var lvnE = "≨︀";
var mDDot = "∺";
var mac = "¯";
var macr$1 = "¯";
var male = "♂";
var malt = "✠";
var maltese = "✠";
var map$1 = "↦";
var mapsto = "↦";
var mapstodown = "↧";
var mapstoleft = "↤";
var mapstoup = "↥";
var marker = "▮";
var mcomma = "⨩";
var mcy = "м";
var mdash$1 = "—";
var measuredangle = "∡";
var mfr = "𝔪";
var mho = "℧";
var micr = "µ";
var micro$1 = "µ";
var mid = "∣";
var midast = "*";
var midcir = "⫰";
var middo = "·";
var middot$1 = "·";
var minus$1 = "−";
var minusb = "⊟";
var minusd = "∸";
var minusdu = "⨪";
var mlcp = "⫛";
var mldr = "…";
var mnplus = "∓";
var models = "⊧";
var mopf = "𝕞";
var mp = "∓";
var mscr = "𝓂";
var mstpos = "∾";
var mu$1 = "μ";
var multimap = "⊸";
var mumap = "⊸";
var nGg = "⋙̸";
var nGt = "≫⃒";
var nGtv = "≫̸";
var nLeftarrow = "⇍";
var nLeftrightarrow = "⇎";
var nLl = "⋘̸";
var nLt = "≪⃒";
var nLtv = "≪̸";
var nRightarrow = "⇏";
var nVDash = "⊯";
var nVdash = "⊮";
var nabla$1 = "∇";
var nacute = "ń";
var nang = "∠⃒";
var nap = "≉";
var napE = "⩰̸";
var napid = "≋̸";
var napos = "ŉ";
var napprox = "≉";
var natur = "♮";
var natural = "♮";
var naturals = "ℕ";
var nbs = " ";
var nbsp$1 = " ";
var nbump = "≎̸";
var nbumpe = "≏̸";
var ncap = "⩃";
var ncaron = "ň";
var ncedil = "ņ";
var ncong = "≇";
var ncongdot = "⩭̸";
var ncup = "⩂";
var ncy = "н";
var ndash$1 = "–";
var ne$1 = "≠";
var neArr = "⇗";
var nearhk = "⤤";
var nearr = "↗";
var nearrow = "↗";
var nedot = "≐̸";
var nequiv = "≢";
var nesear = "⤨";
var nesim = "≂̸";
var nexist = "∄";
var nexists = "∄";
var nfr = "𝔫";
var ngE = "≧̸";
var nge = "≱";
var ngeq = "≱";
var ngeqq = "≧̸";
var ngeqslant = "⩾̸";
var nges = "⩾̸";
var ngsim = "≵";
var ngt = "≯";
var ngtr = "≯";
var nhArr = "⇎";
var nharr = "↮";
var nhpar = "⫲";
var ni$1 = "∋";
var nis = "⋼";
var nisd = "⋺";
var niv = "∋";
var njcy = "њ";
var nlArr = "⇍";
var nlE = "≦̸";
var nlarr = "↚";
var nldr = "‥";
var nle = "≰";
var nleftarrow = "↚";
var nleftrightarrow = "↮";
var nleq = "≰";
var nleqq = "≦̸";
var nleqslant = "⩽̸";
var nles = "⩽̸";
var nless = "≮";
var nlsim = "≴";
var nlt = "≮";
var nltri = "⋪";
var nltrie = "⋬";
var nmid = "∤";
var nopf = "𝕟";
var no = "¬";
var not$1 = "¬";
var notin$1 = "∉";
var notinE = "⋹̸";
var notindot = "⋵̸";
var notinva = "∉";
var notinvb = "⋷";
var notinvc = "⋶";
var notni = "∌";
var notniva = "∌";
var notnivb = "⋾";
var notnivc = "⋽";
var npar = "∦";
var nparallel = "∦";
var nparsl = "⫽⃥";
var npart = "∂̸";
var npolint = "⨔";
var npr = "⊀";
var nprcue = "⋠";
var npre = "⪯̸";
var nprec = "⊀";
var npreceq = "⪯̸";
var nrArr = "⇏";
var nrarr = "↛";
var nrarrc = "⤳̸";
var nrarrw = "↝̸";
var nrightarrow = "↛";
var nrtri = "⋫";
var nrtrie = "⋭";
var nsc = "⊁";
var nsccue = "⋡";
var nsce = "⪰̸";
var nscr = "𝓃";
var nshortmid = "∤";
var nshortparallel = "∦";
var nsim = "≁";
var nsime = "≄";
var nsimeq = "≄";
var nsmid = "∤";
var nspar = "∦";
var nsqsube = "⋢";
var nsqsupe = "⋣";
var nsub$1 = "⊄";
var nsubE = "⫅̸";
var nsube = "⊈";
var nsubset = "⊂⃒";
var nsubseteq = "⊈";
var nsubseteqq = "⫅̸";
var nsucc = "⊁";
var nsucceq = "⪰̸";
var nsup = "⊅";
var nsupE = "⫆̸";
var nsupe = "⊉";
var nsupset = "⊃⃒";
var nsupseteq = "⊉";
var nsupseteqq = "⫆̸";
var ntgl = "≹";
var ntild = "ñ";
var ntilde$1 = "ñ";
var ntlg = "≸";
var ntriangleleft = "⋪";
var ntrianglelefteq = "⋬";
var ntriangleright = "⋫";
var ntrianglerighteq = "⋭";
var nu$1 = "ν";
var num = "#";
var numero = "№";
var numsp = " ";
var nvDash = "⊭";
var nvHarr = "⤄";
var nvap = "≍⃒";
var nvdash = "⊬";
var nvge = "≥⃒";
var nvgt = ">⃒";
var nvinfin = "⧞";
var nvlArr = "⤂";
var nvle = "≤⃒";
var nvlt = "<⃒";
var nvltrie = "⊴⃒";
var nvrArr = "⤃";
var nvrtrie = "⊵⃒";
var nvsim = "∼⃒";
var nwArr = "⇖";
var nwarhk = "⤣";
var nwarr = "↖";
var nwarrow = "↖";
var nwnear = "⤧";
var oS = "Ⓢ";
var oacut = "ó";
var oacute$1 = "ó";
var oast = "⊛";
var ocir = "ô";
var ocirc$1 = "ô";
var ocy = "о";
var odash = "⊝";
var odblac = "ő";
var odiv = "⨸";
var odot = "⊙";
var odsold = "⦼";
var oelig$1 = "œ";
var ofcir = "⦿";
var ofr = "𝔬";
var ogon = "˛";
var ograv = "ò";
var ograve$1 = "ò";
var ogt = "⧁";
var ohbar = "⦵";
var ohm = "Ω";
var oint = "∮";
var olarr = "↺";
var olcir = "⦾";
var olcross = "⦻";
var oline$1 = "‾";
var olt = "⧀";
var omacr = "ō";
var omega$1 = "ω";
var omicron$1 = "ο";
var omid = "⦶";
var ominus = "⊖";
var oopf = "𝕠";
var opar = "⦷";
var operp = "⦹";
var oplus$1 = "⊕";
var or$1 = "∨";
var orarr = "↻";
var ord = "º";
var order = "ℴ";
var orderof = "ℴ";
var ordf$1 = "ª";
var ordm$1 = "º";
var origof = "⊶";
var oror = "⩖";
var orslope = "⩗";
var orv = "⩛";
var oscr = "ℴ";
var oslas = "ø";
var oslash$1 = "ø";
var osol = "⊘";
var otild = "õ";
var otilde$1 = "õ";
var otimes$1 = "⊗";
var otimesas = "⨶";
var oum = "ö";
var ouml$1 = "ö";
var ovbar = "⌽";
var par = "¶";
var para$1 = "¶";
var parallel = "∥";
var parsim = "⫳";
var parsl = "⫽";
var part$1 = "∂";
var pcy = "п";
var percnt = "%";
var period = ".";
var permil$1 = "‰";
var perp$1 = "⊥";
var pertenk = "‱";
var pfr = "𝔭";
var phi$1 = "φ";
var phiv = "ϕ";
var phmmat = "ℳ";
var phone = "☎";
var pi$1 = "π";
var pitchfork = "⋔";
var piv$1 = "ϖ";
var planck = "ℏ";
var planckh = "ℎ";
var plankv = "ℏ";
var plus = "+";
var plusacir = "⨣";
var plusb = "⊞";
var pluscir = "⨢";
var plusdo = "∔";
var plusdu = "⨥";
var pluse = "⩲";
var plusm = "±";
var plusmn$1 = "±";
var plussim = "⨦";
var plustwo = "⨧";
var pm = "±";
var pointint = "⨕";
var popf = "𝕡";
var poun = "£";
var pound$1 = "£";
var pr = "≺";
var prE = "⪳";
var prap = "⪷";
var prcue = "≼";
var pre = "⪯";
var prec = "≺";
var precapprox = "⪷";
var preccurlyeq = "≼";
var preceq = "⪯";
var precnapprox = "⪹";
var precneqq = "⪵";
var precnsim = "⋨";
var precsim = "≾";
var prime$1 = "′";
var primes = "ℙ";
var prnE = "⪵";
var prnap = "⪹";
var prnsim = "⋨";
var prod$1 = "∏";
var profalar = "⌮";
var profline = "⌒";
var profsurf = "⌓";
var prop$1 = "∝";
var propto = "∝";
var prsim = "≾";
var prurel = "⊰";
var pscr = "𝓅";
var psi$1 = "ψ";
var puncsp = " ";
var qfr = "𝔮";
var qint = "⨌";
var qopf = "𝕢";
var qprime = "⁗";
var qscr = "𝓆";
var quaternions = "ℍ";
var quatint = "⨖";
var quest = "?";
var questeq = "≟";
var quo = "\"";
var quot$1 = "\"";
var rAarr = "⇛";
var rArr$1 = "⇒";
var rAtail = "⤜";
var rBarr = "⤏";
var rHar = "⥤";
var race = "∽̱";
var racute = "ŕ";
var radic$1 = "√";
var raemptyv = "⦳";
var rang$1 = "⟩";
var rangd = "⦒";
var range = "⦥";
var rangle = "⟩";
var raqu = "»";
var raquo$1 = "»";
var rarr$1 = "→";
var rarrap = "⥵";
var rarrb = "⇥";
var rarrbfs = "⤠";
var rarrc = "⤳";
var rarrfs = "⤞";
var rarrhk = "↪";
var rarrlp = "↬";
var rarrpl = "⥅";
var rarrsim = "⥴";
var rarrtl = "↣";
var rarrw = "↝";
var ratail = "⤚";
var ratio = "∶";
var rationals = "ℚ";
var rbarr = "⤍";
var rbbrk = "❳";
var rbrace = "}";
var rbrack = "]";
var rbrke = "⦌";
var rbrksld = "⦎";
var rbrkslu = "⦐";
var rcaron = "ř";
var rcedil = "ŗ";
var rceil$1 = "⌉";
var rcub = "}";
var rcy = "р";
var rdca = "⤷";
var rdldhar = "⥩";
var rdquo$1 = "”";
var rdquor = "”";
var rdsh = "↳";
var real$1 = "ℜ";
var realine = "ℛ";
var realpart = "ℜ";
var reals = "ℝ";
var rect = "▭";
var re$1 = "®";
var reg$1 = "®";
var rfisht = "⥽";
var rfloor$1 = "⌋";
var rfr = "𝔯";
var rhard = "⇁";
var rharu = "⇀";
var rharul = "⥬";
var rho$1 = "ρ";
var rhov = "ϱ";
var rightarrow = "→";
var rightarrowtail = "↣";
var rightharpoondown = "⇁";
var rightharpoonup = "⇀";
var rightleftarrows = "⇄";
var rightleftharpoons = "⇌";
var rightrightarrows = "⇉";
var rightsquigarrow = "↝";
var rightthreetimes = "⋌";
var ring = "˚";
var risingdotseq = "≓";
var rlarr = "⇄";
var rlhar = "⇌";
var rlm$1 = "‏";
var rmoust = "⎱";
var rmoustache = "⎱";
var rnmid = "⫮";
var roang = "⟭";
var roarr = "⇾";
var robrk = "⟧";
var ropar = "⦆";
var ropf = "𝕣";
var roplus = "⨮";
var rotimes = "⨵";
var rpar = ")";
var rpargt = "⦔";
var rppolint = "⨒";
var rrarr = "⇉";
var rsaquo$1 = "›";
var rscr = "𝓇";
var rsh = "↱";
var rsqb = "]";
var rsquo$1 = "’";
var rsquor = "’";
var rthree = "⋌";
var rtimes = "⋊";
var rtri = "▹";
var rtrie = "⊵";
var rtrif = "▸";
var rtriltri = "⧎";
var ruluhar = "⥨";
var rx = "℞";
var sacute = "ś";
var sbquo$1 = "‚";
var sc = "≻";
var scE = "⪴";
var scap = "⪸";
var scaron$1 = "š";
var sccue = "≽";
var sce = "⪰";
var scedil = "ş";
var scirc = "ŝ";
var scnE = "⪶";
var scnap = "⪺";
var scnsim = "⋩";
var scpolint = "⨓";
var scsim = "≿";
var scy = "с";
var sdot$1 = "⋅";
var sdotb = "⊡";
var sdote = "⩦";
var seArr = "⇘";
var searhk = "⤥";
var searr = "↘";
var searrow = "↘";
var sec = "§";
var sect$1 = "§";
var semi = ";";
var seswar = "⤩";
var setminus = "∖";
var setmn = "∖";
var sext = "✶";
var sfr = "𝔰";
var sfrown = "⌢";
var sharp = "♯";
var shchcy = "щ";
var shcy = "ш";
var shortmid = "∣";
var shortparallel = "∥";
var sh = "­";
var shy$1 = "­";
var sigma$1 = "σ";
var sigmaf$1 = "ς";
var sigmav = "ς";
var sim$1 = "∼";
var simdot = "⩪";
var sime = "≃";
var simeq = "≃";
var simg = "⪞";
var simgE = "⪠";
var siml = "⪝";
var simlE = "⪟";
var simne = "≆";
var simplus = "⨤";
var simrarr = "⥲";
var slarr = "←";
var smallsetminus = "∖";
var smashp = "⨳";
var smeparsl = "⧤";
var smid = "∣";
var smile = "⌣";
var smt = "⪪";
var smte = "⪬";
var smtes = "⪬︀";
var softcy = "ь";
var sol = "/";
var solb = "⧄";
var solbar = "⌿";
var sopf = "𝕤";
var spades$1 = "♠";
var spadesuit = "♠";
var spar = "∥";
var sqcap = "⊓";
var sqcaps = "⊓︀";
var sqcup = "⊔";
var sqcups = "⊔︀";
var sqsub = "⊏";
var sqsube = "⊑";
var sqsubset = "⊏";
var sqsubseteq = "⊑";
var sqsup = "⊐";
var sqsupe = "⊒";
var sqsupset = "⊐";
var sqsupseteq = "⊒";
var squ = "□";
var square = "□";
var squarf = "▪";
var squf = "▪";
var srarr = "→";
var sscr = "𝓈";
var ssetmn = "∖";
var ssmile = "⌣";
var sstarf = "⋆";
var star = "☆";
var starf = "★";
var straightepsilon = "ϵ";
var straightphi = "ϕ";
var strns = "¯";
var sub$1 = "⊂";
var subE = "⫅";
var subdot = "⪽";
var sube$1 = "⊆";
var subedot = "⫃";
var submult = "⫁";
var subnE = "⫋";
var subne = "⊊";
var subplus = "⪿";
var subrarr = "⥹";
var subset = "⊂";
var subseteq = "⊆";
var subseteqq = "⫅";
var subsetneq = "⊊";
var subsetneqq = "⫋";
var subsim = "⫇";
var subsub = "⫕";
var subsup = "⫓";
var succ = "≻";
var succapprox = "⪸";
var succcurlyeq = "≽";
var succeq = "⪰";
var succnapprox = "⪺";
var succneqq = "⪶";
var succnsim = "⋩";
var succsim = "≿";
var sum$1 = "∑";
var sung = "♪";
var sup$1 = "⊃";
var sup1$1 = "¹";
var sup2$1 = "²";
var sup3$1 = "³";
var supE = "⫆";
var supdot = "⪾";
var supdsub = "⫘";
var supe$1 = "⊇";
var supedot = "⫄";
var suphsol = "⟉";
var suphsub = "⫗";
var suplarr = "⥻";
var supmult = "⫂";
var supnE = "⫌";
var supne = "⊋";
var supplus = "⫀";
var supset = "⊃";
var supseteq = "⊇";
var supseteqq = "⫆";
var supsetneq = "⊋";
var supsetneqq = "⫌";
var supsim = "⫈";
var supsub = "⫔";
var supsup = "⫖";
var swArr = "⇙";
var swarhk = "⤦";
var swarr = "↙";
var swarrow = "↙";
var swnwar = "⤪";
var szli = "ß";
var szlig$1 = "ß";
var target = "⌖";
var tau$1 = "τ";
var tbrk = "⎴";
var tcaron = "ť";
var tcedil = "ţ";
var tcy = "т";
var tdot = "⃛";
var telrec = "⌕";
var tfr = "𝔱";
var there4$1 = "∴";
var therefore = "∴";
var theta$1 = "θ";
var thetasym$1 = "ϑ";
var thetav = "ϑ";
var thickapprox = "≈";
var thicksim = "∼";
var thinsp$1 = " ";
var thkap = "≈";
var thksim = "∼";
var thor = "þ";
var thorn$1 = "þ";
var tilde$1 = "˜";
var time = "×";
var times$1 = "×";
var timesb = "⊠";
var timesbar = "⨱";
var timesd = "⨰";
var tint = "∭";
var toea = "⤨";
var top = "⊤";
var topbot = "⌶";
var topcir = "⫱";
var topf = "𝕥";
var topfork = "⫚";
var tosa = "⤩";
var tprime = "‴";
var trade$1 = "™";
var triangle = "▵";
var triangledown = "▿";
var triangleleft = "◃";
var trianglelefteq = "⊴";
var triangleq = "≜";
var triangleright = "▹";
var trianglerighteq = "⊵";
var tridot = "◬";
var trie = "≜";
var triminus = "⨺";
var triplus = "⨹";
var trisb = "⧍";
var tritime = "⨻";
var trpezium = "⏢";
var tscr = "𝓉";
var tscy = "ц";
var tshcy = "ћ";
var tstrok = "ŧ";
var twixt = "≬";
var twoheadleftarrow = "↞";
var twoheadrightarrow = "↠";
var uArr$1 = "⇑";
var uHar = "⥣";
var uacut = "ú";
var uacute$1 = "ú";
var uarr$1 = "↑";
var ubrcy = "ў";
var ubreve = "ŭ";
var ucir = "û";
var ucirc$1 = "û";
var ucy = "у";
var udarr = "⇅";
var udblac = "ű";
var udhar = "⥮";
var ufisht = "⥾";
var ufr = "𝔲";
var ugrav = "ù";
var ugrave$1 = "ù";
var uharl = "↿";
var uharr = "↾";
var uhblk = "▀";
var ulcorn = "⌜";
var ulcorner = "⌜";
var ulcrop = "⌏";
var ultri = "◸";
var umacr = "ū";
var um = "¨";
var uml$1 = "¨";
var uogon = "ų";
var uopf = "𝕦";
var uparrow = "↑";
var updownarrow = "↕";
var upharpoonleft = "↿";
var upharpoonright = "↾";
var uplus = "⊎";
var upsi = "υ";
var upsih$1 = "ϒ";
var upsilon$1 = "υ";
var upuparrows = "⇈";
var urcorn = "⌝";
var urcorner = "⌝";
var urcrop = "⌎";
var uring = "ů";
var urtri = "◹";
var uscr = "𝓊";
var utdot = "⋰";
var utilde = "ũ";
var utri = "▵";
var utrif = "▴";
var uuarr = "⇈";
var uum = "ü";
var uuml$1 = "ü";
var uwangle = "⦧";
var vArr = "⇕";
var vBar = "⫨";
var vBarv = "⫩";
var vDash = "⊨";
var vangrt = "⦜";
var varepsilon = "ϵ";
var varkappa = "ϰ";
var varnothing = "∅";
var varphi = "ϕ";
var varpi = "ϖ";
var varpropto = "∝";
var varr = "↕";
var varrho = "ϱ";
var varsigma = "ς";
var varsubsetneq = "⊊︀";
var varsubsetneqq = "⫋︀";
var varsupsetneq = "⊋︀";
var varsupsetneqq = "⫌︀";
var vartheta = "ϑ";
var vartriangleleft = "⊲";
var vartriangleright = "⊳";
var vcy = "в";
var vdash = "⊢";
var vee = "∨";
var veebar = "⊻";
var veeeq = "≚";
var vellip = "⋮";
var verbar = "|";
var vert = "|";
var vfr = "𝔳";
var vltri = "⊲";
var vnsub = "⊂⃒";
var vnsup = "⊃⃒";
var vopf = "𝕧";
var vprop = "∝";
var vrtri = "⊳";
var vscr = "𝓋";
var vsubnE = "⫋︀";
var vsubne = "⊊︀";
var vsupnE = "⫌︀";
var vsupne = "⊋︀";
var vzigzag = "⦚";
var wcirc = "ŵ";
var wedbar = "⩟";
var wedge = "∧";
var wedgeq = "≙";
var weierp$1 = "℘";
var wfr = "𝔴";
var wopf = "𝕨";
var wp = "℘";
var wr = "≀";
var wreath = "≀";
var wscr = "𝓌";
var xcap = "⋂";
var xcirc = "◯";
var xcup = "⋃";
var xdtri = "▽";
var xfr = "𝔵";
var xhArr = "⟺";
var xharr = "⟷";
var xi$1 = "ξ";
var xlArr = "⟸";
var xlarr = "⟵";
var xmap = "⟼";
var xnis = "⋻";
var xodot = "⨀";
var xopf = "𝕩";
var xoplus = "⨁";
var xotime = "⨂";
var xrArr = "⟹";
var xrarr = "⟶";
var xscr = "𝓍";
var xsqcup = "⨆";
var xuplus = "⨄";
var xutri = "△";
var xvee = "⋁";
var xwedge = "⋀";
var yacut = "ý";
var yacute$1 = "ý";
var yacy = "я";
var ycirc = "ŷ";
var ycy = "ы";
var ye = "¥";
var yen$1 = "¥";
var yfr = "𝔶";
var yicy = "ї";
var yopf = "𝕪";
var yscr = "𝓎";
var yucy = "ю";
var yum = "ÿ";
var yuml$1 = "ÿ";
var zacute = "ź";
var zcaron = "ž";
var zcy = "з";
var zdot = "ż";
var zeetrf = "ℨ";
var zeta$1 = "ζ";
var zfr = "𝔷";
var zhcy = "ж";
var zigrarr = "⇝";
var zopf = "𝕫";
var zscr = "𝓏";
var zwj$1 = "‍";
var zwnj$1 = "‌";
var require$$0$1 = {
	AEli: AEli,
	AElig: AElig$1,
	AM: AM,
	AMP: AMP$1,
	Aacut: Aacut,
	Aacute: Aacute$1,
	Abreve: Abreve,
	Acir: Acir,
	Acirc: Acirc$1,
	Acy: Acy,
	Afr: Afr,
	Agrav: Agrav,
	Agrave: Agrave$1,
	Alpha: Alpha$1,
	Amacr: Amacr,
	And: And,
	Aogon: Aogon,
	Aopf: Aopf,
	ApplyFunction: ApplyFunction,
	Arin: Arin,
	Aring: Aring$1,
	Ascr: Ascr,
	Assign: Assign,
	Atild: Atild,
	Atilde: Atilde$1,
	Aum: Aum,
	Auml: Auml$1,
	Backslash: Backslash,
	Barv: Barv,
	Barwed: Barwed,
	Bcy: Bcy,
	Because: Because,
	Bernoullis: Bernoullis,
	Beta: Beta$1,
	Bfr: Bfr,
	Bopf: Bopf,
	Breve: Breve,
	Bscr: Bscr,
	Bumpeq: Bumpeq,
	CHcy: CHcy,
	COP: COP,
	COPY: COPY,
	Cacute: Cacute,
	Cap: Cap,
	CapitalDifferentialD: CapitalDifferentialD,
	Cayleys: Cayleys,
	Ccaron: Ccaron,
	Ccedi: Ccedi,
	Ccedil: Ccedil$1,
	Ccirc: Ccirc,
	Cconint: Cconint,
	Cdot: Cdot,
	Cedilla: Cedilla,
	CenterDot: CenterDot,
	Cfr: Cfr,
	Chi: Chi$1,
	CircleDot: CircleDot,
	CircleMinus: CircleMinus,
	CirclePlus: CirclePlus,
	CircleTimes: CircleTimes,
	ClockwiseContourIntegral: ClockwiseContourIntegral,
	CloseCurlyDoubleQuote: CloseCurlyDoubleQuote,
	CloseCurlyQuote: CloseCurlyQuote,
	Colon: Colon,
	Colone: Colone,
	Congruent: Congruent,
	Conint: Conint,
	ContourIntegral: ContourIntegral,
	Copf: Copf,
	Coproduct: Coproduct,
	CounterClockwiseContourIntegral: CounterClockwiseContourIntegral,
	Cross: Cross,
	Cscr: Cscr,
	Cup: Cup,
	CupCap: CupCap,
	DD: DD,
	DDotrahd: DDotrahd,
	DJcy: DJcy,
	DScy: DScy,
	DZcy: DZcy,
	Dagger: Dagger$1,
	Darr: Darr,
	Dashv: Dashv,
	Dcaron: Dcaron,
	Dcy: Dcy,
	Del: Del,
	Delta: Delta$1,
	Dfr: Dfr,
	DiacriticalAcute: DiacriticalAcute,
	DiacriticalDot: DiacriticalDot,
	DiacriticalDoubleAcute: DiacriticalDoubleAcute,
	DiacriticalGrave: DiacriticalGrave,
	DiacriticalTilde: DiacriticalTilde,
	Diamond: Diamond,
	DifferentialD: DifferentialD,
	Dopf: Dopf,
	Dot: Dot,
	DotDot: DotDot,
	DotEqual: DotEqual,
	DoubleContourIntegral: DoubleContourIntegral,
	DoubleDot: DoubleDot,
	DoubleDownArrow: DoubleDownArrow,
	DoubleLeftArrow: DoubleLeftArrow,
	DoubleLeftRightArrow: DoubleLeftRightArrow,
	DoubleLeftTee: DoubleLeftTee,
	DoubleLongLeftArrow: DoubleLongLeftArrow,
	DoubleLongLeftRightArrow: DoubleLongLeftRightArrow,
	DoubleLongRightArrow: DoubleLongRightArrow,
	DoubleRightArrow: DoubleRightArrow,
	DoubleRightTee: DoubleRightTee,
	DoubleUpArrow: DoubleUpArrow,
	DoubleUpDownArrow: DoubleUpDownArrow,
	DoubleVerticalBar: DoubleVerticalBar,
	DownArrow: DownArrow,
	DownArrowBar: DownArrowBar,
	DownArrowUpArrow: DownArrowUpArrow,
	DownBreve: DownBreve,
	DownLeftRightVector: DownLeftRightVector,
	DownLeftTeeVector: DownLeftTeeVector,
	DownLeftVector: DownLeftVector,
	DownLeftVectorBar: DownLeftVectorBar,
	DownRightTeeVector: DownRightTeeVector,
	DownRightVector: DownRightVector,
	DownRightVectorBar: DownRightVectorBar,
	DownTee: DownTee,
	DownTeeArrow: DownTeeArrow,
	Downarrow: Downarrow,
	Dscr: Dscr,
	Dstrok: Dstrok,
	ENG: ENG,
	ET: ET,
	ETH: ETH$1,
	Eacut: Eacut,
	Eacute: Eacute$1,
	Ecaron: Ecaron,
	Ecir: Ecir,
	Ecirc: Ecirc$1,
	Ecy: Ecy,
	Edot: Edot,
	Efr: Efr,
	Egrav: Egrav,
	Egrave: Egrave$1,
	Element: Element,
	Emacr: Emacr,
	EmptySmallSquare: EmptySmallSquare,
	EmptyVerySmallSquare: EmptyVerySmallSquare,
	Eogon: Eogon,
	Eopf: Eopf,
	Epsilon: Epsilon$1,
	Equal: Equal,
	EqualTilde: EqualTilde,
	Equilibrium: Equilibrium,
	Escr: Escr,
	Esim: Esim,
	Eta: Eta$1,
	Eum: Eum,
	Euml: Euml$1,
	Exists: Exists,
	ExponentialE: ExponentialE,
	Fcy: Fcy,
	Ffr: Ffr,
	FilledSmallSquare: FilledSmallSquare,
	FilledVerySmallSquare: FilledVerySmallSquare,
	Fopf: Fopf,
	ForAll: ForAll,
	Fouriertrf: Fouriertrf,
	Fscr: Fscr,
	GJcy: GJcy,
	G: G,
	GT: GT$1,
	Gamma: Gamma$1,
	Gammad: Gammad,
	Gbreve: Gbreve,
	Gcedil: Gcedil,
	Gcirc: Gcirc,
	Gcy: Gcy,
	Gdot: Gdot,
	Gfr: Gfr,
	Gg: Gg,
	Gopf: Gopf,
	GreaterEqual: GreaterEqual,
	GreaterEqualLess: GreaterEqualLess,
	GreaterFullEqual: GreaterFullEqual,
	GreaterGreater: GreaterGreater,
	GreaterLess: GreaterLess,
	GreaterSlantEqual: GreaterSlantEqual,
	GreaterTilde: GreaterTilde,
	Gscr: Gscr,
	Gt: Gt,
	HARDcy: HARDcy,
	Hacek: Hacek,
	Hat: Hat,
	Hcirc: Hcirc,
	Hfr: Hfr,
	HilbertSpace: HilbertSpace,
	Hopf: Hopf,
	HorizontalLine: HorizontalLine,
	Hscr: Hscr,
	Hstrok: Hstrok,
	HumpDownHump: HumpDownHump,
	HumpEqual: HumpEqual,
	IEcy: IEcy,
	IJlig: IJlig,
	IOcy: IOcy,
	Iacut: Iacut,
	Iacute: Iacute$1,
	Icir: Icir,
	Icirc: Icirc$1,
	Icy: Icy,
	Idot: Idot,
	Ifr: Ifr,
	Igrav: Igrav,
	Igrave: Igrave$1,
	Im: Im,
	Imacr: Imacr,
	ImaginaryI: ImaginaryI,
	Implies: Implies,
	Int: Int,
	Integral: Integral,
	Intersection: Intersection,
	InvisibleComma: InvisibleComma,
	InvisibleTimes: InvisibleTimes,
	Iogon: Iogon,
	Iopf: Iopf,
	Iota: Iota$1,
	Iscr: Iscr,
	Itilde: Itilde,
	Iukcy: Iukcy,
	Ium: Ium,
	Iuml: Iuml$1,
	Jcirc: Jcirc,
	Jcy: Jcy,
	Jfr: Jfr,
	Jopf: Jopf,
	Jscr: Jscr,
	Jsercy: Jsercy,
	Jukcy: Jukcy,
	KHcy: KHcy,
	KJcy: KJcy,
	Kappa: Kappa$1,
	Kcedil: Kcedil,
	Kcy: Kcy,
	Kfr: Kfr,
	Kopf: Kopf,
	Kscr: Kscr,
	LJcy: LJcy,
	L: L,
	LT: LT$1,
	Lacute: Lacute,
	Lambda: Lambda$1,
	Lang: Lang,
	Laplacetrf: Laplacetrf,
	Larr: Larr,
	Lcaron: Lcaron,
	Lcedil: Lcedil,
	Lcy: Lcy,
	LeftAngleBracket: LeftAngleBracket,
	LeftArrow: LeftArrow,
	LeftArrowBar: LeftArrowBar,
	LeftArrowRightArrow: LeftArrowRightArrow,
	LeftCeiling: LeftCeiling,
	LeftDoubleBracket: LeftDoubleBracket,
	LeftDownTeeVector: LeftDownTeeVector,
	LeftDownVector: LeftDownVector,
	LeftDownVectorBar: LeftDownVectorBar,
	LeftFloor: LeftFloor,
	LeftRightArrow: LeftRightArrow,
	LeftRightVector: LeftRightVector,
	LeftTee: LeftTee,
	LeftTeeArrow: LeftTeeArrow,
	LeftTeeVector: LeftTeeVector,
	LeftTriangle: LeftTriangle,
	LeftTriangleBar: LeftTriangleBar,
	LeftTriangleEqual: LeftTriangleEqual,
	LeftUpDownVector: LeftUpDownVector,
	LeftUpTeeVector: LeftUpTeeVector,
	LeftUpVector: LeftUpVector,
	LeftUpVectorBar: LeftUpVectorBar,
	LeftVector: LeftVector,
	LeftVectorBar: LeftVectorBar,
	Leftarrow: Leftarrow,
	Leftrightarrow: Leftrightarrow,
	LessEqualGreater: LessEqualGreater,
	LessFullEqual: LessFullEqual,
	LessGreater: LessGreater,
	LessLess: LessLess,
	LessSlantEqual: LessSlantEqual,
	LessTilde: LessTilde,
	Lfr: Lfr,
	Ll: Ll,
	Lleftarrow: Lleftarrow,
	Lmidot: Lmidot,
	LongLeftArrow: LongLeftArrow,
	LongLeftRightArrow: LongLeftRightArrow,
	LongRightArrow: LongRightArrow,
	Longleftarrow: Longleftarrow,
	Longleftrightarrow: Longleftrightarrow,
	Longrightarrow: Longrightarrow,
	Lopf: Lopf,
	LowerLeftArrow: LowerLeftArrow,
	LowerRightArrow: LowerRightArrow,
	Lscr: Lscr,
	Lsh: Lsh,
	Lstrok: Lstrok,
	Lt: Lt,
	"Map": "⤅",
	Mcy: Mcy,
	MediumSpace: MediumSpace,
	Mellintrf: Mellintrf,
	Mfr: Mfr,
	MinusPlus: MinusPlus,
	Mopf: Mopf,
	Mscr: Mscr,
	Mu: Mu$1,
	NJcy: NJcy,
	Nacute: Nacute,
	Ncaron: Ncaron,
	Ncedil: Ncedil,
	Ncy: Ncy,
	NegativeMediumSpace: NegativeMediumSpace,
	NegativeThickSpace: NegativeThickSpace,
	NegativeThinSpace: NegativeThinSpace,
	NegativeVeryThinSpace: NegativeVeryThinSpace,
	NestedGreaterGreater: NestedGreaterGreater,
	NestedLessLess: NestedLessLess,
	NewLine: NewLine,
	Nfr: Nfr,
	NoBreak: NoBreak,
	NonBreakingSpace: NonBreakingSpace,
	Nopf: Nopf,
	Not: Not,
	NotCongruent: NotCongruent,
	NotCupCap: NotCupCap,
	NotDoubleVerticalBar: NotDoubleVerticalBar,
	NotElement: NotElement,
	NotEqual: NotEqual,
	NotEqualTilde: NotEqualTilde,
	NotExists: NotExists,
	NotGreater: NotGreater,
	NotGreaterEqual: NotGreaterEqual,
	NotGreaterFullEqual: NotGreaterFullEqual,
	NotGreaterGreater: NotGreaterGreater,
	NotGreaterLess: NotGreaterLess,
	NotGreaterSlantEqual: NotGreaterSlantEqual,
	NotGreaterTilde: NotGreaterTilde,
	NotHumpDownHump: NotHumpDownHump,
	NotHumpEqual: NotHumpEqual,
	NotLeftTriangle: NotLeftTriangle,
	NotLeftTriangleBar: NotLeftTriangleBar,
	NotLeftTriangleEqual: NotLeftTriangleEqual,
	NotLess: NotLess,
	NotLessEqual: NotLessEqual,
	NotLessGreater: NotLessGreater,
	NotLessLess: NotLessLess,
	NotLessSlantEqual: NotLessSlantEqual,
	NotLessTilde: NotLessTilde,
	NotNestedGreaterGreater: NotNestedGreaterGreater,
	NotNestedLessLess: NotNestedLessLess,
	NotPrecedes: NotPrecedes,
	NotPrecedesEqual: NotPrecedesEqual,
	NotPrecedesSlantEqual: NotPrecedesSlantEqual,
	NotReverseElement: NotReverseElement,
	NotRightTriangle: NotRightTriangle,
	NotRightTriangleBar: NotRightTriangleBar,
	NotRightTriangleEqual: NotRightTriangleEqual,
	NotSquareSubset: NotSquareSubset,
	NotSquareSubsetEqual: NotSquareSubsetEqual,
	NotSquareSuperset: NotSquareSuperset,
	NotSquareSupersetEqual: NotSquareSupersetEqual,
	NotSubset: NotSubset,
	NotSubsetEqual: NotSubsetEqual,
	NotSucceeds: NotSucceeds,
	NotSucceedsEqual: NotSucceedsEqual,
	NotSucceedsSlantEqual: NotSucceedsSlantEqual,
	NotSucceedsTilde: NotSucceedsTilde,
	NotSuperset: NotSuperset,
	NotSupersetEqual: NotSupersetEqual,
	NotTilde: NotTilde,
	NotTildeEqual: NotTildeEqual,
	NotTildeFullEqual: NotTildeFullEqual,
	NotTildeTilde: NotTildeTilde,
	NotVerticalBar: NotVerticalBar,
	Nscr: Nscr,
	Ntild: Ntild,
	Ntilde: Ntilde$1,
	Nu: Nu$1,
	OElig: OElig$1,
	Oacut: Oacut,
	Oacute: Oacute$1,
	Ocir: Ocir,
	Ocirc: Ocirc$1,
	Ocy: Ocy,
	Odblac: Odblac,
	Ofr: Ofr,
	Ograv: Ograv,
	Ograve: Ograve$1,
	Omacr: Omacr,
	Omega: Omega$1,
	Omicron: Omicron$1,
	Oopf: Oopf,
	OpenCurlyDoubleQuote: OpenCurlyDoubleQuote,
	OpenCurlyQuote: OpenCurlyQuote,
	Or: Or,
	Oscr: Oscr,
	Oslas: Oslas,
	Oslash: Oslash$1,
	Otild: Otild,
	Otilde: Otilde$1,
	Otimes: Otimes,
	Oum: Oum,
	Ouml: Ouml$1,
	OverBar: OverBar,
	OverBrace: OverBrace,
	OverBracket: OverBracket,
	OverParenthesis: OverParenthesis,
	PartialD: PartialD,
	Pcy: Pcy,
	Pfr: Pfr,
	Phi: Phi$1,
	Pi: Pi$1,
	PlusMinus: PlusMinus,
	Poincareplane: Poincareplane,
	Popf: Popf,
	Pr: Pr,
	Precedes: Precedes,
	PrecedesEqual: PrecedesEqual,
	PrecedesSlantEqual: PrecedesSlantEqual,
	PrecedesTilde: PrecedesTilde,
	Prime: Prime$1,
	Product: Product,
	Proportion: Proportion,
	Proportional: Proportional,
	Pscr: Pscr,
	Psi: Psi$1,
	QUO: QUO,
	QUOT: QUOT,
	Qfr: Qfr,
	Qopf: Qopf,
	Qscr: Qscr,
	RBarr: RBarr,
	RE: RE,
	REG: REG,
	Racute: Racute,
	Rang: Rang,
	Rarr: Rarr,
	Rarrtl: Rarrtl,
	Rcaron: Rcaron,
	Rcedil: Rcedil,
	Rcy: Rcy,
	Re: Re,
	ReverseElement: ReverseElement,
	ReverseEquilibrium: ReverseEquilibrium,
	ReverseUpEquilibrium: ReverseUpEquilibrium,
	Rfr: Rfr,
	Rho: Rho$1,
	RightAngleBracket: RightAngleBracket,
	RightArrow: RightArrow,
	RightArrowBar: RightArrowBar,
	RightArrowLeftArrow: RightArrowLeftArrow,
	RightCeiling: RightCeiling,
	RightDoubleBracket: RightDoubleBracket,
	RightDownTeeVector: RightDownTeeVector,
	RightDownVector: RightDownVector,
	RightDownVectorBar: RightDownVectorBar,
	RightFloor: RightFloor,
	RightTee: RightTee,
	RightTeeArrow: RightTeeArrow,
	RightTeeVector: RightTeeVector,
	RightTriangle: RightTriangle,
	RightTriangleBar: RightTriangleBar,
	RightTriangleEqual: RightTriangleEqual,
	RightUpDownVector: RightUpDownVector,
	RightUpTeeVector: RightUpTeeVector,
	RightUpVector: RightUpVector,
	RightUpVectorBar: RightUpVectorBar,
	RightVector: RightVector,
	RightVectorBar: RightVectorBar,
	Rightarrow: Rightarrow,
	Ropf: Ropf,
	RoundImplies: RoundImplies,
	Rrightarrow: Rrightarrow,
	Rscr: Rscr,
	Rsh: Rsh,
	RuleDelayed: RuleDelayed,
	SHCHcy: SHCHcy,
	SHcy: SHcy,
	SOFTcy: SOFTcy,
	Sacute: Sacute,
	Sc: Sc,
	Scaron: Scaron$1,
	Scedil: Scedil,
	Scirc: Scirc,
	Scy: Scy,
	Sfr: Sfr,
	ShortDownArrow: ShortDownArrow,
	ShortLeftArrow: ShortLeftArrow,
	ShortRightArrow: ShortRightArrow,
	ShortUpArrow: ShortUpArrow,
	Sigma: Sigma$1,
	SmallCircle: SmallCircle,
	Sopf: Sopf,
	Sqrt: Sqrt,
	Square: Square,
	SquareIntersection: SquareIntersection,
	SquareSubset: SquareSubset,
	SquareSubsetEqual: SquareSubsetEqual,
	SquareSuperset: SquareSuperset,
	SquareSupersetEqual: SquareSupersetEqual,
	SquareUnion: SquareUnion,
	Sscr: Sscr,
	Star: Star,
	Sub: Sub,
	Subset: Subset,
	SubsetEqual: SubsetEqual,
	Succeeds: Succeeds,
	SucceedsEqual: SucceedsEqual,
	SucceedsSlantEqual: SucceedsSlantEqual,
	SucceedsTilde: SucceedsTilde,
	SuchThat: SuchThat,
	Sum: Sum,
	Sup: Sup,
	Superset: Superset,
	SupersetEqual: SupersetEqual,
	Supset: Supset,
	THOR: THOR,
	THORN: THORN$1,
	TRADE: TRADE,
	TSHcy: TSHcy,
	TScy: TScy,
	Tab: Tab,
	Tau: Tau$1,
	Tcaron: Tcaron,
	Tcedil: Tcedil,
	Tcy: Tcy,
	Tfr: Tfr,
	Therefore: Therefore,
	Theta: Theta$1,
	ThickSpace: ThickSpace,
	ThinSpace: ThinSpace,
	Tilde: Tilde,
	TildeEqual: TildeEqual,
	TildeFullEqual: TildeFullEqual,
	TildeTilde: TildeTilde,
	Topf: Topf,
	TripleDot: TripleDot,
	Tscr: Tscr,
	Tstrok: Tstrok,
	Uacut: Uacut,
	Uacute: Uacute$1,
	Uarr: Uarr,
	Uarrocir: Uarrocir,
	Ubrcy: Ubrcy,
	Ubreve: Ubreve,
	Ucir: Ucir,
	Ucirc: Ucirc$1,
	Ucy: Ucy,
	Udblac: Udblac,
	Ufr: Ufr,
	Ugrav: Ugrav,
	Ugrave: Ugrave$1,
	Umacr: Umacr,
	UnderBar: UnderBar,
	UnderBrace: UnderBrace,
	UnderBracket: UnderBracket,
	UnderParenthesis: UnderParenthesis,
	Union: Union,
	UnionPlus: UnionPlus,
	Uogon: Uogon,
	Uopf: Uopf,
	UpArrow: UpArrow,
	UpArrowBar: UpArrowBar,
	UpArrowDownArrow: UpArrowDownArrow,
	UpDownArrow: UpDownArrow,
	UpEquilibrium: UpEquilibrium,
	UpTee: UpTee,
	UpTeeArrow: UpTeeArrow,
	Uparrow: Uparrow,
	Updownarrow: Updownarrow,
	UpperLeftArrow: UpperLeftArrow,
	UpperRightArrow: UpperRightArrow,
	Upsi: Upsi,
	Upsilon: Upsilon$1,
	Uring: Uring,
	Uscr: Uscr,
	Utilde: Utilde,
	Uum: Uum,
	Uuml: Uuml$1,
	VDash: VDash,
	Vbar: Vbar,
	Vcy: Vcy,
	Vdash: Vdash,
	Vdashl: Vdashl,
	Vee: Vee,
	Verbar: Verbar,
	Vert: Vert,
	VerticalBar: VerticalBar,
	VerticalLine: VerticalLine,
	VerticalSeparator: VerticalSeparator,
	VerticalTilde: VerticalTilde,
	VeryThinSpace: VeryThinSpace,
	Vfr: Vfr,
	Vopf: Vopf,
	Vscr: Vscr,
	Vvdash: Vvdash,
	Wcirc: Wcirc,
	Wedge: Wedge,
	Wfr: Wfr,
	Wopf: Wopf,
	Wscr: Wscr,
	Xfr: Xfr,
	Xi: Xi$1,
	Xopf: Xopf,
	Xscr: Xscr,
	YAcy: YAcy,
	YIcy: YIcy,
	YUcy: YUcy,
	Yacut: Yacut,
	Yacute: Yacute$1,
	Ycirc: Ycirc,
	Ycy: Ycy,
	Yfr: Yfr,
	Yopf: Yopf,
	Yscr: Yscr,
	Yuml: Yuml$1,
	ZHcy: ZHcy,
	Zacute: Zacute,
	Zcaron: Zcaron,
	Zcy: Zcy,
	Zdot: Zdot,
	ZeroWidthSpace: ZeroWidthSpace,
	Zeta: Zeta$1,
	Zfr: Zfr,
	Zopf: Zopf,
	Zscr: Zscr,
	aacut: aacut,
	aacute: aacute$1,
	abreve: abreve,
	ac: ac,
	acE: acE,
	acd: acd,
	acir: acir,
	acirc: acirc$1,
	acut: acut,
	acute: acute$1,
	acy: acy,
	aeli: aeli,
	aelig: aelig$1,
	af: af,
	afr: afr,
	agrav: agrav,
	agrave: agrave$1,
	alefsym: alefsym$1,
	aleph: aleph,
	alpha: alpha$1,
	amacr: amacr,
	amalg: amalg,
	am: am,
	amp: amp$1,
	and: and$1,
	andand: andand,
	andd: andd,
	andslope: andslope,
	andv: andv,
	ang: ang$1,
	ange: ange,
	angle: angle,
	angmsd: angmsd,
	angmsdaa: angmsdaa,
	angmsdab: angmsdab,
	angmsdac: angmsdac,
	angmsdad: angmsdad,
	angmsdae: angmsdae,
	angmsdaf: angmsdaf,
	angmsdag: angmsdag,
	angmsdah: angmsdah,
	angrt: angrt,
	angrtvb: angrtvb,
	angrtvbd: angrtvbd,
	angsph: angsph,
	angst: angst,
	angzarr: angzarr,
	aogon: aogon,
	aopf: aopf,
	ap: ap,
	apE: apE,
	apacir: apacir,
	ape: ape,
	apid: apid,
	apos: apos,
	approx: approx,
	approxeq: approxeq,
	arin: arin,
	aring: aring$1,
	ascr: ascr,
	ast: ast,
	asymp: asymp$1,
	asympeq: asympeq,
	atild: atild,
	atilde: atilde$1,
	aum: aum,
	auml: auml$1,
	awconint: awconint,
	awint: awint,
	bNot: bNot,
	backcong: backcong,
	backepsilon: backepsilon,
	backprime: backprime,
	backsim: backsim,
	backsimeq: backsimeq,
	barvee: barvee,
	barwed: barwed,
	barwedge: barwedge,
	bbrk: bbrk,
	bbrktbrk: bbrktbrk,
	bcong: bcong,
	bcy: bcy,
	bdquo: bdquo$1,
	becaus: becaus,
	because: because,
	bemptyv: bemptyv,
	bepsi: bepsi,
	bernou: bernou,
	beta: beta$1,
	beth: beth,
	between: between,
	bfr: bfr,
	bigcap: bigcap,
	bigcirc: bigcirc,
	bigcup: bigcup,
	bigodot: bigodot,
	bigoplus: bigoplus,
	bigotimes: bigotimes,
	bigsqcup: bigsqcup,
	bigstar: bigstar,
	bigtriangledown: bigtriangledown,
	bigtriangleup: bigtriangleup,
	biguplus: biguplus,
	bigvee: bigvee,
	bigwedge: bigwedge,
	bkarow: bkarow,
	blacklozenge: blacklozenge,
	blacksquare: blacksquare,
	blacktriangle: blacktriangle,
	blacktriangledown: blacktriangledown,
	blacktriangleleft: blacktriangleleft,
	blacktriangleright: blacktriangleright,
	blank: blank,
	blk12: blk12,
	blk14: blk14,
	blk34: blk34,
	block: block,
	bne: bne,
	bnequiv: bnequiv,
	bnot: bnot,
	bopf: bopf,
	bot: bot,
	bottom: bottom,
	bowtie: bowtie,
	boxDL: boxDL,
	boxDR: boxDR,
	boxDl: boxDl,
	boxDr: boxDr,
	boxH: boxH,
	boxHD: boxHD,
	boxHU: boxHU,
	boxHd: boxHd,
	boxHu: boxHu,
	boxUL: boxUL,
	boxUR: boxUR,
	boxUl: boxUl,
	boxUr: boxUr,
	boxV: boxV,
	boxVH: boxVH,
	boxVL: boxVL,
	boxVR: boxVR,
	boxVh: boxVh,
	boxVl: boxVl,
	boxVr: boxVr,
	boxbox: boxbox,
	boxdL: boxdL,
	boxdR: boxdR,
	boxdl: boxdl,
	boxdr: boxdr,
	boxh: boxh,
	boxhD: boxhD,
	boxhU: boxhU,
	boxhd: boxhd,
	boxhu: boxhu,
	boxminus: boxminus,
	boxplus: boxplus,
	boxtimes: boxtimes,
	boxuL: boxuL,
	boxuR: boxuR,
	boxul: boxul,
	boxur: boxur,
	boxv: boxv,
	boxvH: boxvH,
	boxvL: boxvL,
	boxvR: boxvR,
	boxvh: boxvh,
	boxvl: boxvl,
	boxvr: boxvr,
	bprime: bprime,
	breve: breve,
	brvba: brvba,
	brvbar: brvbar$1,
	bscr: bscr,
	bsemi: bsemi,
	bsim: bsim,
	bsime: bsime,
	bsol: bsol,
	bsolb: bsolb,
	bsolhsub: bsolhsub,
	bull: bull$1,
	bullet: bullet,
	bump: bump,
	bumpE: bumpE,
	bumpe: bumpe,
	bumpeq: bumpeq,
	cacute: cacute,
	cap: cap$2,
	capand: capand,
	capbrcup: capbrcup,
	capcap: capcap,
	capcup: capcup,
	capdot: capdot,
	caps: caps,
	caret: caret,
	caron: caron,
	ccaps: ccaps,
	ccaron: ccaron,
	ccedi: ccedi,
	ccedil: ccedil$1,
	ccirc: ccirc,
	ccups: ccups,
	ccupssm: ccupssm,
	cdot: cdot,
	cedi: cedi,
	cedil: cedil$1,
	cemptyv: cemptyv,
	cen: cen,
	cent: cent$1,
	centerdot: centerdot,
	cfr: cfr,
	chcy: chcy,
	check: check,
	checkmark: checkmark,
	chi: chi$1,
	cir: cir,
	cirE: cirE,
	circ: circ$1,
	circeq: circeq,
	circlearrowleft: circlearrowleft,
	circlearrowright: circlearrowright,
	circledR: circledR,
	circledS: circledS,
	circledast: circledast,
	circledcirc: circledcirc,
	circleddash: circleddash,
	cire: cire,
	cirfnint: cirfnint,
	cirmid: cirmid,
	cirscir: cirscir,
	clubs: clubs$1,
	clubsuit: clubsuit,
	colon: colon,
	colone: colone,
	coloneq: coloneq,
	comma: comma$1,
	commat: commat,
	comp: comp,
	compfn: compfn,
	complement: complement,
	complexes: complexes,
	cong: cong$1,
	congdot: congdot,
	conint: conint,
	copf: copf,
	coprod: coprod,
	cop: cop,
	copy: copy$1,
	copysr: copysr,
	crarr: crarr$1,
	cross: cross,
	cscr: cscr,
	csub: csub,
	csube: csube,
	csup: csup,
	csupe: csupe,
	ctdot: ctdot,
	cudarrl: cudarrl,
	cudarrr: cudarrr,
	cuepr: cuepr,
	cuesc: cuesc,
	cularr: cularr,
	cularrp: cularrp,
	cup: cup$1,
	cupbrcap: cupbrcap,
	cupcap: cupcap,
	cupcup: cupcup,
	cupdot: cupdot,
	cupor: cupor,
	cups: cups,
	curarr: curarr,
	curarrm: curarrm,
	curlyeqprec: curlyeqprec,
	curlyeqsucc: curlyeqsucc,
	curlyvee: curlyvee,
	curlywedge: curlywedge,
	curre: curre,
	curren: curren$1,
	curvearrowleft: curvearrowleft,
	curvearrowright: curvearrowright,
	cuvee: cuvee,
	cuwed: cuwed,
	cwconint: cwconint,
	cwint: cwint,
	cylcty: cylcty,
	dArr: dArr$1,
	dHar: dHar,
	dagger: dagger$1,
	daleth: daleth,
	darr: darr$1,
	dash: dash$1,
	dashv: dashv,
	dbkarow: dbkarow,
	dblac: dblac,
	dcaron: dcaron,
	dcy: dcy,
	dd: dd$1,
	ddagger: ddagger,
	ddarr: ddarr,
	ddotseq: ddotseq,
	de: de,
	deg: deg$1,
	delta: delta$1,
	demptyv: demptyv,
	dfisht: dfisht,
	dfr: dfr,
	dharl: dharl,
	dharr: dharr,
	diam: diam,
	diamond: diamond,
	diamondsuit: diamondsuit,
	diams: diams$1,
	die: die,
	digamma: digamma,
	disin: disin,
	div: div,
	divid: divid,
	divide: divide$1,
	divideontimes: divideontimes,
	divonx: divonx,
	djcy: djcy,
	dlcorn: dlcorn,
	dlcrop: dlcrop,
	dollar: dollar,
	dopf: dopf,
	dot: dot,
	doteq: doteq,
	doteqdot: doteqdot,
	dotminus: dotminus,
	dotplus: dotplus,
	dotsquare: dotsquare,
	doublebarwedge: doublebarwedge,
	downarrow: downarrow,
	downdownarrows: downdownarrows,
	downharpoonleft: downharpoonleft,
	downharpoonright: downharpoonright,
	drbkarow: drbkarow,
	drcorn: drcorn,
	drcrop: drcrop,
	dscr: dscr,
	dscy: dscy,
	dsol: dsol,
	dstrok: dstrok,
	dtdot: dtdot,
	dtri: dtri,
	dtrif: dtrif,
	duarr: duarr,
	duhar: duhar,
	dwangle: dwangle,
	dzcy: dzcy,
	dzigrarr: dzigrarr,
	eDDot: eDDot,
	eDot: eDot,
	eacut: eacut,
	eacute: eacute$1,
	easter: easter,
	ecaron: ecaron,
	ecir: ecir,
	ecirc: ecirc$1,
	ecolon: ecolon,
	ecy: ecy,
	edot: edot,
	ee: ee,
	efDot: efDot,
	efr: efr,
	eg: eg,
	egrav: egrav,
	egrave: egrave$1,
	egs: egs,
	egsdot: egsdot,
	el: el,
	elinters: elinters,
	ell: ell,
	els: els,
	elsdot: elsdot,
	emacr: emacr,
	empty: empty$3,
	emptyset: emptyset,
	emptyv: emptyv,
	emsp13: emsp13,
	emsp14: emsp14,
	emsp: emsp$1,
	eng: eng,
	ensp: ensp$1,
	eogon: eogon,
	eopf: eopf,
	epar: epar,
	eparsl: eparsl,
	eplus: eplus,
	epsi: epsi,
	epsilon: epsilon$1,
	epsiv: epsiv,
	eqcirc: eqcirc,
	eqcolon: eqcolon,
	eqsim: eqsim,
	eqslantgtr: eqslantgtr,
	eqslantless: eqslantless,
	equals: equals,
	equest: equest,
	equiv: equiv$1,
	equivDD: equivDD,
	eqvparsl: eqvparsl,
	erDot: erDot,
	erarr: erarr,
	escr: escr,
	esdot: esdot,
	esim: esim,
	eta: eta$1,
	et: et,
	eth: eth$1,
	eum: eum,
	euml: euml$1,
	euro: euro$1,
	excl: excl,
	exist: exist$1,
	expectation: expectation,
	exponentiale: exponentiale,
	fallingdotseq: fallingdotseq,
	fcy: fcy,
	female: female,
	ffilig: ffilig,
	fflig: fflig,
	ffllig: ffllig,
	ffr: ffr,
	filig: filig,
	fjlig: fjlig,
	flat: flat,
	fllig: fllig,
	fltns: fltns,
	fnof: fnof$1,
	fopf: fopf,
	forall: forall$1,
	fork: fork,
	forkv: forkv,
	fpartint: fpartint,
	frac1: frac1,
	frac12: frac12$1,
	frac13: frac13,
	frac14: frac14$1,
	frac15: frac15,
	frac16: frac16,
	frac18: frac18,
	frac23: frac23,
	frac25: frac25,
	frac3: frac3,
	frac34: frac34$1,
	frac35: frac35,
	frac38: frac38,
	frac45: frac45,
	frac56: frac56,
	frac58: frac58,
	frac78: frac78,
	frasl: frasl$1,
	frown: frown,
	fscr: fscr,
	gE: gE,
	gEl: gEl,
	gacute: gacute,
	gamma: gamma$1,
	gammad: gammad,
	gap: gap,
	gbreve: gbreve,
	gcirc: gcirc,
	gcy: gcy,
	gdot: gdot,
	ge: ge$1,
	gel: gel,
	geq: geq,
	geqq: geqq,
	geqslant: geqslant,
	ges: ges,
	gescc: gescc,
	gesdot: gesdot,
	gesdoto: gesdoto,
	gesdotol: gesdotol,
	gesl: gesl,
	gesles: gesles,
	gfr: gfr,
	gg: gg,
	ggg: ggg,
	gimel: gimel,
	gjcy: gjcy,
	gl: gl,
	glE: glE,
	gla: gla,
	glj: glj,
	gnE: gnE,
	gnap: gnap,
	gnapprox: gnapprox,
	gne: gne,
	gneq: gneq,
	gneqq: gneqq,
	gnsim: gnsim,
	gopf: gopf,
	grave: grave,
	gscr: gscr,
	gsim: gsim,
	gsime: gsime,
	gsiml: gsiml,
	g: g,
	gt: gt$1,
	gtcc: gtcc,
	gtcir: gtcir,
	gtdot: gtdot,
	gtlPar: gtlPar,
	gtquest: gtquest,
	gtrapprox: gtrapprox,
	gtrarr: gtrarr,
	gtrdot: gtrdot,
	gtreqless: gtreqless,
	gtreqqless: gtreqqless,
	gtrless: gtrless,
	gtrsim: gtrsim,
	gvertneqq: gvertneqq,
	gvnE: gvnE,
	hArr: hArr$1,
	hairsp: hairsp,
	half: half,
	hamilt: hamilt,
	hardcy: hardcy,
	harr: harr$1,
	harrcir: harrcir,
	harrw: harrw,
	hbar: hbar,
	hcirc: hcirc,
	hearts: hearts$1,
	heartsuit: heartsuit,
	hellip: hellip$1,
	hercon: hercon,
	hfr: hfr,
	hksearow: hksearow,
	hkswarow: hkswarow,
	hoarr: hoarr,
	homtht: homtht,
	hookleftarrow: hookleftarrow,
	hookrightarrow: hookrightarrow,
	hopf: hopf,
	horbar: horbar,
	hscr: hscr,
	hslash: hslash,
	hstrok: hstrok,
	hybull: hybull,
	hyphen: hyphen,
	iacut: iacut,
	iacute: iacute$1,
	ic: ic,
	icir: icir,
	icirc: icirc$1,
	icy: icy,
	iecy: iecy,
	iexc: iexc,
	iexcl: iexcl$1,
	iff: iff,
	ifr: ifr,
	igrav: igrav,
	igrave: igrave$1,
	ii: ii,
	iiiint: iiiint,
	iiint: iiint,
	iinfin: iinfin,
	iiota: iiota,
	ijlig: ijlig,
	imacr: imacr,
	image: image$2,
	imagline: imagline,
	imagpart: imagpart,
	imath: imath,
	imof: imof,
	imped: imped,
	"in": "∈",
	incare: incare,
	infin: infin$1,
	infintie: infintie,
	inodot: inodot,
	int: int$2,
	intcal: intcal,
	integers: integers,
	intercal: intercal,
	intlarhk: intlarhk,
	intprod: intprod,
	iocy: iocy,
	iogon: iogon,
	iopf: iopf,
	iota: iota$1,
	iprod: iprod,
	iques: iques,
	iquest: iquest$1,
	iscr: iscr,
	isin: isin$1,
	isinE: isinE,
	isindot: isindot,
	isins: isins,
	isinsv: isinsv,
	isinv: isinv,
	it: it,
	itilde: itilde,
	iukcy: iukcy,
	ium: ium,
	iuml: iuml$1,
	jcirc: jcirc,
	jcy: jcy,
	jfr: jfr,
	jmath: jmath,
	jopf: jopf,
	jscr: jscr,
	jsercy: jsercy,
	jukcy: jukcy,
	kappa: kappa$1,
	kappav: kappav,
	kcedil: kcedil,
	kcy: kcy,
	kfr: kfr,
	kgreen: kgreen,
	khcy: khcy,
	kjcy: kjcy,
	kopf: kopf,
	kscr: kscr,
	lAarr: lAarr,
	lArr: lArr$1,
	lAtail: lAtail,
	lBarr: lBarr,
	lE: lE,
	lEg: lEg,
	lHar: lHar,
	lacute: lacute,
	laemptyv: laemptyv,
	lagran: lagran,
	lambda: lambda$1,
	lang: lang$1,
	langd: langd,
	langle: langle,
	lap: lap,
	laqu: laqu,
	laquo: laquo$1,
	larr: larr$1,
	larrb: larrb,
	larrbfs: larrbfs,
	larrfs: larrfs,
	larrhk: larrhk,
	larrlp: larrlp,
	larrpl: larrpl,
	larrsim: larrsim,
	larrtl: larrtl,
	lat: lat,
	latail: latail,
	late: late,
	lates: lates,
	lbarr: lbarr,
	lbbrk: lbbrk,
	lbrace: lbrace,
	lbrack: lbrack,
	lbrke: lbrke,
	lbrksld: lbrksld,
	lbrkslu: lbrkslu,
	lcaron: lcaron,
	lcedil: lcedil,
	lceil: lceil$1,
	lcub: lcub,
	lcy: lcy,
	ldca: ldca,
	ldquo: ldquo$1,
	ldquor: ldquor,
	ldrdhar: ldrdhar,
	ldrushar: ldrushar,
	ldsh: ldsh,
	le: le$1,
	leftarrow: leftarrow,
	leftarrowtail: leftarrowtail,
	leftharpoondown: leftharpoondown,
	leftharpoonup: leftharpoonup,
	leftleftarrows: leftleftarrows,
	leftrightarrow: leftrightarrow,
	leftrightarrows: leftrightarrows,
	leftrightharpoons: leftrightharpoons,
	leftrightsquigarrow: leftrightsquigarrow,
	leftthreetimes: leftthreetimes,
	leg: leg,
	leq: leq,
	leqq: leqq,
	leqslant: leqslant,
	les: les,
	lescc: lescc,
	lesdot: lesdot,
	lesdoto: lesdoto,
	lesdotor: lesdotor,
	lesg: lesg,
	lesges: lesges,
	lessapprox: lessapprox,
	lessdot: lessdot,
	lesseqgtr: lesseqgtr,
	lesseqqgtr: lesseqqgtr,
	lessgtr: lessgtr,
	lesssim: lesssim,
	lfisht: lfisht,
	lfloor: lfloor$1,
	lfr: lfr,
	lg: lg,
	lgE: lgE,
	lhard: lhard,
	lharu: lharu,
	lharul: lharul,
	lhblk: lhblk,
	ljcy: ljcy,
	ll: ll,
	llarr: llarr,
	llcorner: llcorner,
	llhard: llhard,
	lltri: lltri,
	lmidot: lmidot,
	lmoust: lmoust,
	lmoustache: lmoustache,
	lnE: lnE,
	lnap: lnap,
	lnapprox: lnapprox,
	lne: lne,
	lneq: lneq,
	lneqq: lneqq,
	lnsim: lnsim,
	loang: loang,
	loarr: loarr,
	lobrk: lobrk,
	longleftarrow: longleftarrow,
	longleftrightarrow: longleftrightarrow,
	longmapsto: longmapsto,
	longrightarrow: longrightarrow,
	looparrowleft: looparrowleft,
	looparrowright: looparrowright,
	lopar: lopar,
	lopf: lopf,
	loplus: loplus,
	lotimes: lotimes,
	lowast: lowast$1,
	lowbar: lowbar,
	loz: loz$1,
	lozenge: lozenge,
	lozf: lozf,
	lpar: lpar,
	lparlt: lparlt,
	lrarr: lrarr,
	lrcorner: lrcorner,
	lrhar: lrhar,
	lrhard: lrhard,
	lrm: lrm$1,
	lrtri: lrtri,
	lsaquo: lsaquo$1,
	lscr: lscr,
	lsh: lsh,
	lsim: lsim,
	lsime: lsime,
	lsimg: lsimg,
	lsqb: lsqb,
	lsquo: lsquo$1,
	lsquor: lsquor,
	lstrok: lstrok,
	l: l,
	lt: lt$1,
	ltcc: ltcc,
	ltcir: ltcir,
	ltdot: ltdot,
	lthree: lthree,
	ltimes: ltimes,
	ltlarr: ltlarr,
	ltquest: ltquest,
	ltrPar: ltrPar,
	ltri: ltri,
	ltrie: ltrie,
	ltrif: ltrif,
	lurdshar: lurdshar,
	luruhar: luruhar,
	lvertneqq: lvertneqq,
	lvnE: lvnE,
	mDDot: mDDot,
	mac: mac,
	macr: macr$1,
	male: male,
	malt: malt,
	maltese: maltese,
	map: map$1,
	mapsto: mapsto,
	mapstodown: mapstodown,
	mapstoleft: mapstoleft,
	mapstoup: mapstoup,
	marker: marker,
	mcomma: mcomma,
	mcy: mcy,
	mdash: mdash$1,
	measuredangle: measuredangle,
	mfr: mfr,
	mho: mho,
	micr: micr,
	micro: micro$1,
	mid: mid,
	midast: midast,
	midcir: midcir,
	middo: middo,
	middot: middot$1,
	minus: minus$1,
	minusb: minusb,
	minusd: minusd,
	minusdu: minusdu,
	mlcp: mlcp,
	mldr: mldr,
	mnplus: mnplus,
	models: models,
	mopf: mopf,
	mp: mp,
	mscr: mscr,
	mstpos: mstpos,
	mu: mu$1,
	multimap: multimap,
	mumap: mumap,
	nGg: nGg,
	nGt: nGt,
	nGtv: nGtv,
	nLeftarrow: nLeftarrow,
	nLeftrightarrow: nLeftrightarrow,
	nLl: nLl,
	nLt: nLt,
	nLtv: nLtv,
	nRightarrow: nRightarrow,
	nVDash: nVDash,
	nVdash: nVdash,
	nabla: nabla$1,
	nacute: nacute,
	nang: nang,
	nap: nap,
	napE: napE,
	napid: napid,
	napos: napos,
	napprox: napprox,
	natur: natur,
	natural: natural,
	naturals: naturals,
	nbs: nbs,
	nbsp: nbsp$1,
	nbump: nbump,
	nbumpe: nbumpe,
	ncap: ncap,
	ncaron: ncaron,
	ncedil: ncedil,
	ncong: ncong,
	ncongdot: ncongdot,
	ncup: ncup,
	ncy: ncy,
	ndash: ndash$1,
	ne: ne$1,
	neArr: neArr,
	nearhk: nearhk,
	nearr: nearr,
	nearrow: nearrow,
	nedot: nedot,
	nequiv: nequiv,
	nesear: nesear,
	nesim: nesim,
	nexist: nexist,
	nexists: nexists,
	nfr: nfr,
	ngE: ngE,
	nge: nge,
	ngeq: ngeq,
	ngeqq: ngeqq,
	ngeqslant: ngeqslant,
	nges: nges,
	ngsim: ngsim,
	ngt: ngt,
	ngtr: ngtr,
	nhArr: nhArr,
	nharr: nharr,
	nhpar: nhpar,
	ni: ni$1,
	nis: nis,
	nisd: nisd,
	niv: niv,
	njcy: njcy,
	nlArr: nlArr,
	nlE: nlE,
	nlarr: nlarr,
	nldr: nldr,
	nle: nle,
	nleftarrow: nleftarrow,
	nleftrightarrow: nleftrightarrow,
	nleq: nleq,
	nleqq: nleqq,
	nleqslant: nleqslant,
	nles: nles,
	nless: nless,
	nlsim: nlsim,
	nlt: nlt,
	nltri: nltri,
	nltrie: nltrie,
	nmid: nmid,
	nopf: nopf,
	no: no,
	not: not$1,
	notin: notin$1,
	notinE: notinE,
	notindot: notindot,
	notinva: notinva,
	notinvb: notinvb,
	notinvc: notinvc,
	notni: notni,
	notniva: notniva,
	notnivb: notnivb,
	notnivc: notnivc,
	npar: npar,
	nparallel: nparallel,
	nparsl: nparsl,
	npart: npart,
	npolint: npolint,
	npr: npr,
	nprcue: nprcue,
	npre: npre,
	nprec: nprec,
	npreceq: npreceq,
	nrArr: nrArr,
	nrarr: nrarr,
	nrarrc: nrarrc,
	nrarrw: nrarrw,
	nrightarrow: nrightarrow,
	nrtri: nrtri,
	nrtrie: nrtrie,
	nsc: nsc,
	nsccue: nsccue,
	nsce: nsce,
	nscr: nscr,
	nshortmid: nshortmid,
	nshortparallel: nshortparallel,
	nsim: nsim,
	nsime: nsime,
	nsimeq: nsimeq,
	nsmid: nsmid,
	nspar: nspar,
	nsqsube: nsqsube,
	nsqsupe: nsqsupe,
	nsub: nsub$1,
	nsubE: nsubE,
	nsube: nsube,
	nsubset: nsubset,
	nsubseteq: nsubseteq,
	nsubseteqq: nsubseteqq,
	nsucc: nsucc,
	nsucceq: nsucceq,
	nsup: nsup,
	nsupE: nsupE,
	nsupe: nsupe,
	nsupset: nsupset,
	nsupseteq: nsupseteq,
	nsupseteqq: nsupseteqq,
	ntgl: ntgl,
	ntild: ntild,
	ntilde: ntilde$1,
	ntlg: ntlg,
	ntriangleleft: ntriangleleft,
	ntrianglelefteq: ntrianglelefteq,
	ntriangleright: ntriangleright,
	ntrianglerighteq: ntrianglerighteq,
	nu: nu$1,
	num: num,
	numero: numero,
	numsp: numsp,
	nvDash: nvDash,
	nvHarr: nvHarr,
	nvap: nvap,
	nvdash: nvdash,
	nvge: nvge,
	nvgt: nvgt,
	nvinfin: nvinfin,
	nvlArr: nvlArr,
	nvle: nvle,
	nvlt: nvlt,
	nvltrie: nvltrie,
	nvrArr: nvrArr,
	nvrtrie: nvrtrie,
	nvsim: nvsim,
	nwArr: nwArr,
	nwarhk: nwarhk,
	nwarr: nwarr,
	nwarrow: nwarrow,
	nwnear: nwnear,
	oS: oS,
	oacut: oacut,
	oacute: oacute$1,
	oast: oast,
	ocir: ocir,
	ocirc: ocirc$1,
	ocy: ocy,
	odash: odash,
	odblac: odblac,
	odiv: odiv,
	odot: odot,
	odsold: odsold,
	oelig: oelig$1,
	ofcir: ofcir,
	ofr: ofr,
	ogon: ogon,
	ograv: ograv,
	ograve: ograve$1,
	ogt: ogt,
	ohbar: ohbar,
	ohm: ohm,
	oint: oint,
	olarr: olarr,
	olcir: olcir,
	olcross: olcross,
	oline: oline$1,
	olt: olt,
	omacr: omacr,
	omega: omega$1,
	omicron: omicron$1,
	omid: omid,
	ominus: ominus,
	oopf: oopf,
	opar: opar,
	operp: operp,
	oplus: oplus$1,
	or: or$1,
	orarr: orarr,
	ord: ord,
	order: order,
	orderof: orderof,
	ordf: ordf$1,
	ordm: ordm$1,
	origof: origof,
	oror: oror,
	orslope: orslope,
	orv: orv,
	oscr: oscr,
	oslas: oslas,
	oslash: oslash$1,
	osol: osol,
	otild: otild,
	otilde: otilde$1,
	otimes: otimes$1,
	otimesas: otimesas,
	oum: oum,
	ouml: ouml$1,
	ovbar: ovbar,
	par: par,
	para: para$1,
	parallel: parallel,
	parsim: parsim,
	parsl: parsl,
	part: part$1,
	pcy: pcy,
	percnt: percnt,
	period: period,
	permil: permil$1,
	perp: perp$1,
	pertenk: pertenk,
	pfr: pfr,
	phi: phi$1,
	phiv: phiv,
	phmmat: phmmat,
	phone: phone,
	pi: pi$1,
	pitchfork: pitchfork,
	piv: piv$1,
	planck: planck,
	planckh: planckh,
	plankv: plankv,
	plus: plus,
	plusacir: plusacir,
	plusb: plusb,
	pluscir: pluscir,
	plusdo: plusdo,
	plusdu: plusdu,
	pluse: pluse,
	plusm: plusm,
	plusmn: plusmn$1,
	plussim: plussim,
	plustwo: plustwo,
	pm: pm,
	pointint: pointint,
	popf: popf,
	poun: poun,
	pound: pound$1,
	pr: pr,
	prE: prE,
	prap: prap,
	prcue: prcue,
	pre: pre,
	prec: prec,
	precapprox: precapprox,
	preccurlyeq: preccurlyeq,
	preceq: preceq,
	precnapprox: precnapprox,
	precneqq: precneqq,
	precnsim: precnsim,
	precsim: precsim,
	prime: prime$1,
	primes: primes,
	prnE: prnE,
	prnap: prnap,
	prnsim: prnsim,
	prod: prod$1,
	profalar: profalar,
	profline: profline,
	profsurf: profsurf,
	prop: prop$1,
	propto: propto,
	prsim: prsim,
	prurel: prurel,
	pscr: pscr,
	psi: psi$1,
	puncsp: puncsp,
	qfr: qfr,
	qint: qint,
	qopf: qopf,
	qprime: qprime,
	qscr: qscr,
	quaternions: quaternions,
	quatint: quatint,
	quest: quest,
	questeq: questeq,
	quo: quo,
	quot: quot$1,
	rAarr: rAarr,
	rArr: rArr$1,
	rAtail: rAtail,
	rBarr: rBarr,
	rHar: rHar,
	race: race,
	racute: racute,
	radic: radic$1,
	raemptyv: raemptyv,
	rang: rang$1,
	rangd: rangd,
	range: range,
	rangle: rangle,
	raqu: raqu,
	raquo: raquo$1,
	rarr: rarr$1,
	rarrap: rarrap,
	rarrb: rarrb,
	rarrbfs: rarrbfs,
	rarrc: rarrc,
	rarrfs: rarrfs,
	rarrhk: rarrhk,
	rarrlp: rarrlp,
	rarrpl: rarrpl,
	rarrsim: rarrsim,
	rarrtl: rarrtl,
	rarrw: rarrw,
	ratail: ratail,
	ratio: ratio,
	rationals: rationals,
	rbarr: rbarr,
	rbbrk: rbbrk,
	rbrace: rbrace,
	rbrack: rbrack,
	rbrke: rbrke,
	rbrksld: rbrksld,
	rbrkslu: rbrkslu,
	rcaron: rcaron,
	rcedil: rcedil,
	rceil: rceil$1,
	rcub: rcub,
	rcy: rcy,
	rdca: rdca,
	rdldhar: rdldhar,
	rdquo: rdquo$1,
	rdquor: rdquor,
	rdsh: rdsh,
	real: real$1,
	realine: realine,
	realpart: realpart,
	reals: reals,
	rect: rect,
	re: re$1,
	reg: reg$1,
	rfisht: rfisht,
	rfloor: rfloor$1,
	rfr: rfr,
	rhard: rhard,
	rharu: rharu,
	rharul: rharul,
	rho: rho$1,
	rhov: rhov,
	rightarrow: rightarrow,
	rightarrowtail: rightarrowtail,
	rightharpoondown: rightharpoondown,
	rightharpoonup: rightharpoonup,
	rightleftarrows: rightleftarrows,
	rightleftharpoons: rightleftharpoons,
	rightrightarrows: rightrightarrows,
	rightsquigarrow: rightsquigarrow,
	rightthreetimes: rightthreetimes,
	ring: ring,
	risingdotseq: risingdotseq,
	rlarr: rlarr,
	rlhar: rlhar,
	rlm: rlm$1,
	rmoust: rmoust,
	rmoustache: rmoustache,
	rnmid: rnmid,
	roang: roang,
	roarr: roarr,
	robrk: robrk,
	ropar: ropar,
	ropf: ropf,
	roplus: roplus,
	rotimes: rotimes,
	rpar: rpar,
	rpargt: rpargt,
	rppolint: rppolint,
	rrarr: rrarr,
	rsaquo: rsaquo$1,
	rscr: rscr,
	rsh: rsh,
	rsqb: rsqb,
	rsquo: rsquo$1,
	rsquor: rsquor,
	rthree: rthree,
	rtimes: rtimes,
	rtri: rtri,
	rtrie: rtrie,
	rtrif: rtrif,
	rtriltri: rtriltri,
	ruluhar: ruluhar,
	rx: rx,
	sacute: sacute,
	sbquo: sbquo$1,
	sc: sc,
	scE: scE,
	scap: scap,
	scaron: scaron$1,
	sccue: sccue,
	sce: sce,
	scedil: scedil,
	scirc: scirc,
	scnE: scnE,
	scnap: scnap,
	scnsim: scnsim,
	scpolint: scpolint,
	scsim: scsim,
	scy: scy,
	sdot: sdot$1,
	sdotb: sdotb,
	sdote: sdote,
	seArr: seArr,
	searhk: searhk,
	searr: searr,
	searrow: searrow,
	sec: sec,
	sect: sect$1,
	semi: semi,
	seswar: seswar,
	setminus: setminus,
	setmn: setmn,
	sext: sext,
	sfr: sfr,
	sfrown: sfrown,
	sharp: sharp,
	shchcy: shchcy,
	shcy: shcy,
	shortmid: shortmid,
	shortparallel: shortparallel,
	sh: sh,
	shy: shy$1,
	sigma: sigma$1,
	sigmaf: sigmaf$1,
	sigmav: sigmav,
	sim: sim$1,
	simdot: simdot,
	sime: sime,
	simeq: simeq,
	simg: simg,
	simgE: simgE,
	siml: siml,
	simlE: simlE,
	simne: simne,
	simplus: simplus,
	simrarr: simrarr,
	slarr: slarr,
	smallsetminus: smallsetminus,
	smashp: smashp,
	smeparsl: smeparsl,
	smid: smid,
	smile: smile,
	smt: smt,
	smte: smte,
	smtes: smtes,
	softcy: softcy,
	sol: sol,
	solb: solb,
	solbar: solbar,
	sopf: sopf,
	spades: spades$1,
	spadesuit: spadesuit,
	spar: spar,
	sqcap: sqcap,
	sqcaps: sqcaps,
	sqcup: sqcup,
	sqcups: sqcups,
	sqsub: sqsub,
	sqsube: sqsube,
	sqsubset: sqsubset,
	sqsubseteq: sqsubseteq,
	sqsup: sqsup,
	sqsupe: sqsupe,
	sqsupset: sqsupset,
	sqsupseteq: sqsupseteq,
	squ: squ,
	square: square,
	squarf: squarf,
	squf: squf,
	srarr: srarr,
	sscr: sscr,
	ssetmn: ssetmn,
	ssmile: ssmile,
	sstarf: sstarf,
	star: star,
	starf: starf,
	straightepsilon: straightepsilon,
	straightphi: straightphi,
	strns: strns,
	sub: sub$1,
	subE: subE,
	subdot: subdot,
	sube: sube$1,
	subedot: subedot,
	submult: submult,
	subnE: subnE,
	subne: subne,
	subplus: subplus,
	subrarr: subrarr,
	subset: subset,
	subseteq: subseteq,
	subseteqq: subseteqq,
	subsetneq: subsetneq,
	subsetneqq: subsetneqq,
	subsim: subsim,
	subsub: subsub,
	subsup: subsup,
	succ: succ,
	succapprox: succapprox,
	succcurlyeq: succcurlyeq,
	succeq: succeq,
	succnapprox: succnapprox,
	succneqq: succneqq,
	succnsim: succnsim,
	succsim: succsim,
	sum: sum$1,
	sung: sung,
	sup: sup$1,
	sup1: sup1$1,
	sup2: sup2$1,
	sup3: sup3$1,
	supE: supE,
	supdot: supdot,
	supdsub: supdsub,
	supe: supe$1,
	supedot: supedot,
	suphsol: suphsol,
	suphsub: suphsub,
	suplarr: suplarr,
	supmult: supmult,
	supnE: supnE,
	supne: supne,
	supplus: supplus,
	supset: supset,
	supseteq: supseteq,
	supseteqq: supseteqq,
	supsetneq: supsetneq,
	supsetneqq: supsetneqq,
	supsim: supsim,
	supsub: supsub,
	supsup: supsup,
	swArr: swArr,
	swarhk: swarhk,
	swarr: swarr,
	swarrow: swarrow,
	swnwar: swnwar,
	szli: szli,
	szlig: szlig$1,
	target: target,
	tau: tau$1,
	tbrk: tbrk,
	tcaron: tcaron,
	tcedil: tcedil,
	tcy: tcy,
	tdot: tdot,
	telrec: telrec,
	tfr: tfr,
	there4: there4$1,
	therefore: therefore,
	theta: theta$1,
	thetasym: thetasym$1,
	thetav: thetav,
	thickapprox: thickapprox,
	thicksim: thicksim,
	thinsp: thinsp$1,
	thkap: thkap,
	thksim: thksim,
	thor: thor,
	thorn: thorn$1,
	tilde: tilde$1,
	time: time,
	times: times$1,
	timesb: timesb,
	timesbar: timesbar,
	timesd: timesd,
	tint: tint,
	toea: toea,
	top: top,
	topbot: topbot,
	topcir: topcir,
	topf: topf,
	topfork: topfork,
	tosa: tosa,
	tprime: tprime,
	trade: trade$1,
	triangle: triangle,
	triangledown: triangledown,
	triangleleft: triangleleft,
	trianglelefteq: trianglelefteq,
	triangleq: triangleq,
	triangleright: triangleright,
	trianglerighteq: trianglerighteq,
	tridot: tridot,
	trie: trie,
	triminus: triminus,
	triplus: triplus,
	trisb: trisb,
	tritime: tritime,
	trpezium: trpezium,
	tscr: tscr,
	tscy: tscy,
	tshcy: tshcy,
	tstrok: tstrok,
	twixt: twixt,
	twoheadleftarrow: twoheadleftarrow,
	twoheadrightarrow: twoheadrightarrow,
	uArr: uArr$1,
	uHar: uHar,
	uacut: uacut,
	uacute: uacute$1,
	uarr: uarr$1,
	ubrcy: ubrcy,
	ubreve: ubreve,
	ucir: ucir,
	ucirc: ucirc$1,
	ucy: ucy,
	udarr: udarr,
	udblac: udblac,
	udhar: udhar,
	ufisht: ufisht,
	ufr: ufr,
	ugrav: ugrav,
	ugrave: ugrave$1,
	uharl: uharl,
	uharr: uharr,
	uhblk: uhblk,
	ulcorn: ulcorn,
	ulcorner: ulcorner,
	ulcrop: ulcrop,
	ultri: ultri,
	umacr: umacr,
	um: um,
	uml: uml$1,
	uogon: uogon,
	uopf: uopf,
	uparrow: uparrow,
	updownarrow: updownarrow,
	upharpoonleft: upharpoonleft,
	upharpoonright: upharpoonright,
	uplus: uplus,
	upsi: upsi,
	upsih: upsih$1,
	upsilon: upsilon$1,
	upuparrows: upuparrows,
	urcorn: urcorn,
	urcorner: urcorner,
	urcrop: urcrop,
	uring: uring,
	urtri: urtri,
	uscr: uscr,
	utdot: utdot,
	utilde: utilde,
	utri: utri,
	utrif: utrif,
	uuarr: uuarr,
	uum: uum,
	uuml: uuml$1,
	uwangle: uwangle,
	vArr: vArr,
	vBar: vBar,
	vBarv: vBarv,
	vDash: vDash,
	vangrt: vangrt,
	varepsilon: varepsilon,
	varkappa: varkappa,
	varnothing: varnothing,
	varphi: varphi,
	varpi: varpi,
	varpropto: varpropto,
	varr: varr,
	varrho: varrho,
	varsigma: varsigma,
	varsubsetneq: varsubsetneq,
	varsubsetneqq: varsubsetneqq,
	varsupsetneq: varsupsetneq,
	varsupsetneqq: varsupsetneqq,
	vartheta: vartheta,
	vartriangleleft: vartriangleleft,
	vartriangleright: vartriangleright,
	vcy: vcy,
	vdash: vdash,
	vee: vee,
	veebar: veebar,
	veeeq: veeeq,
	vellip: vellip,
	verbar: verbar,
	vert: vert,
	vfr: vfr,
	vltri: vltri,
	vnsub: vnsub,
	vnsup: vnsup,
	vopf: vopf,
	vprop: vprop,
	vrtri: vrtri,
	vscr: vscr,
	vsubnE: vsubnE,
	vsubne: vsubne,
	vsupnE: vsupnE,
	vsupne: vsupne,
	vzigzag: vzigzag,
	wcirc: wcirc,
	wedbar: wedbar,
	wedge: wedge,
	wedgeq: wedgeq,
	weierp: weierp$1,
	wfr: wfr,
	wopf: wopf,
	wp: wp,
	wr: wr,
	wreath: wreath,
	wscr: wscr,
	xcap: xcap,
	xcirc: xcirc,
	xcup: xcup,
	xdtri: xdtri,
	xfr: xfr,
	xhArr: xhArr,
	xharr: xharr,
	xi: xi$1,
	xlArr: xlArr,
	xlarr: xlarr,
	xmap: xmap,
	xnis: xnis,
	xodot: xodot,
	xopf: xopf,
	xoplus: xoplus,
	xotime: xotime,
	xrArr: xrArr,
	xrarr: xrarr,
	xscr: xscr,
	xsqcup: xsqcup,
	xuplus: xuplus,
	xutri: xutri,
	xvee: xvee,
	xwedge: xwedge,
	yacut: yacut,
	yacute: yacute$1,
	yacy: yacy,
	ycirc: ycirc,
	ycy: ycy,
	ye: ye,
	yen: yen$1,
	yfr: yfr,
	yicy: yicy,
	yopf: yopf,
	yscr: yscr,
	yucy: yucy,
	yum: yum,
	yuml: yuml$1,
	zacute: zacute,
	zcaron: zcaron,
	zcy: zcy,
	zdot: zdot,
	zeetrf: zeetrf,
	zeta: zeta$1,
	zfr: zfr,
	zhcy: zhcy,
	zigrarr: zigrarr,
	zopf: zopf,
	zscr: zscr,
	zwj: zwj$1,
	zwnj: zwnj$1
};

var characterEntities = require$$0$1;

var decodeEntity_1 = decodeEntity$1;

var own$9 = {}.hasOwnProperty;

function decodeEntity$1(characters) {
  return own$9.call(characterEntities, characters)
    ? characterEntities[characters]
    : false
}

var legacy$1 = require$$1$2;
var invalid = require$$1$1;
var decimal$1 = requireIsDecimal();
var hexadecimal$1 = isHexadecimal;
var alphanumerical$1 = isAlphanumerical;
var decodeEntity = decodeEntity_1;

var parseEntities_1 = parseEntities;

var own$8 = {}.hasOwnProperty;
var fromCharCode = String.fromCharCode;
var noop = Function.prototype;

// Default settings.
var defaults$2 = {
  warning: null,
  reference: null,
  text: null,
  warningContext: null,
  referenceContext: null,
  textContext: null,
  position: {},
  additional: null,
  attribute: false,
  nonTerminated: true
};

// Characters.
var tab$1 = 9; // '\t'
var lineFeed$1 = 10; // '\n'
var formFeed = 12; // '\f'
var space$3 = 32; // ' '
var ampersand = 38; // '&'
var semicolon = 59; // ';'
var lessThan$1 = 60; // '<'
var equalsTo = 61; // '='
var numberSign = 35; // '#'
var uppercaseX = 88; // 'X'
var lowercaseX = 120; // 'x'
var replacementCharacter = 65533; // '�'

// Reference types.
var name$1 = 'named';
var hexa = 'hexadecimal';
var deci = 'decimal';

// Map of bases.
var bases = {};

bases[hexa] = 16;
bases[deci] = 10;

// Map of types to tests.
// Each type of character reference accepts different characters.
// This test is used to detect whether a reference has ended (as the semicolon
// is not strictly needed).
var tests = {};

tests[name$1] = alphanumerical$1;
tests[deci] = decimal$1;
tests[hexa] = hexadecimal$1;

// Warning types.
var namedNotTerminated = 1;
var numericNotTerminated = 2;
var namedEmpty = 3;
var numericEmpty = 4;
var namedUnknown = 5;
var numericDisallowed = 6;
var numericProhibited = 7;

// Warning messages.
var messages = {};

messages[namedNotTerminated] =
  'Named character references must be terminated by a semicolon';
messages[numericNotTerminated] =
  'Numeric character references must be terminated by a semicolon';
messages[namedEmpty] = 'Named character references cannot be empty';
messages[numericEmpty] = 'Numeric character references cannot be empty';
messages[namedUnknown] = 'Named character references must be known';
messages[numericDisallowed] =
  'Numeric character references cannot be disallowed';
messages[numericProhibited] =
  'Numeric character references cannot be outside the permissible Unicode range';

// Wrap to ensure clean parameters are given to `parse`.
function parseEntities(value, options) {
  var settings = {};
  var option;
  var key;

  if (!options) {
    options = {};
  }

  for (key in defaults$2) {
    option = options[key];
    settings[key] =
      option === null || option === undefined ? defaults$2[key] : option;
  }

  if (settings.position.indent || settings.position.start) {
    settings.indent = settings.position.indent || [];
    settings.position = settings.position.start;
  }

  return parse$6(value, settings)
}

// Parse entities.
// eslint-disable-next-line complexity
function parse$6(value, settings) {
  var additional = settings.additional;
  var nonTerminated = settings.nonTerminated;
  var handleText = settings.text;
  var handleReference = settings.reference;
  var handleWarning = settings.warning;
  var textContext = settings.textContext;
  var referenceContext = settings.referenceContext;
  var warningContext = settings.warningContext;
  var pos = settings.position;
  var indent = settings.indent || [];
  var length = value.length;
  var index = 0;
  var lines = -1;
  var column = pos.column || 1;
  var line = pos.line || 1;
  var queue = '';
  var result = [];
  var entityCharacters;
  var namedEntity;
  var terminated;
  var characters;
  var character;
  var reference;
  var following;
  var warning;
  var reason;
  var output;
  var entity;
  var begin;
  var start;
  var type;
  var test;
  var prev;
  var next;
  var diff;
  var end;

  if (typeof additional === 'string') {
    additional = additional.charCodeAt(0);
  }

  // Cache the current point.
  prev = now();

  // Wrap `handleWarning`.
  warning = handleWarning ? parseError : noop;

  // Ensure the algorithm walks over the first character and the end
  // (inclusive).
  index--;
  length++;

  while (++index < length) {
    // If the previous character was a newline.
    if (character === lineFeed$1) {
      column = indent[lines] || 1;
    }

    character = value.charCodeAt(index);

    if (character === ampersand) {
      following = value.charCodeAt(index + 1);

      // The behaviour depends on the identity of the next character.
      if (
        following === tab$1 ||
        following === lineFeed$1 ||
        following === formFeed ||
        following === space$3 ||
        following === ampersand ||
        following === lessThan$1 ||
        following !== following ||
        (additional && following === additional)
      ) {
        // Not a character reference.
        // No characters are consumed, and nothing is returned.
        // This is not an error, either.
        queue += fromCharCode(character);
        column++;

        continue
      }

      start = index + 1;
      begin = start;
      end = start;

      if (following === numberSign) {
        // Numerical entity.
        end = ++begin;

        // The behaviour further depends on the next character.
        following = value.charCodeAt(end);

        if (following === uppercaseX || following === lowercaseX) {
          // ASCII hex digits.
          type = hexa;
          end = ++begin;
        } else {
          // ASCII digits.
          type = deci;
        }
      } else {
        // Named entity.
        type = name$1;
      }

      entityCharacters = '';
      entity = '';
      characters = '';
      test = tests[type];
      end--;

      while (++end < length) {
        following = value.charCodeAt(end);

        if (!test(following)) {
          break
        }

        characters += fromCharCode(following);

        // Check if we can match a legacy named reference.
        // If so, we cache that as the last viable named reference.
        // This ensures we do not need to walk backwards later.
        if (type === name$1 && own$8.call(legacy$1, characters)) {
          entityCharacters = characters;
          entity = legacy$1[characters];
        }
      }

      terminated = value.charCodeAt(end) === semicolon;

      if (terminated) {
        end++;

        namedEntity = type === name$1 ? decodeEntity(characters) : false;

        if (namedEntity) {
          entityCharacters = characters;
          entity = namedEntity;
        }
      }

      diff = 1 + end - start;

      if (!terminated && !nonTerminated) ; else if (!characters) {
        // An empty (possible) entity is valid, unless it’s numeric (thus an
        // ampersand followed by an octothorp).
        if (type !== name$1) {
          warning(numericEmpty, diff);
        }
      } else if (type === name$1) {
        // An ampersand followed by anything unknown, and not terminated, is
        // invalid.
        if (terminated && !entity) {
          warning(namedUnknown, 1);
        } else {
          // If theres something after an entity name which is not known, cap
          // the reference.
          if (entityCharacters !== characters) {
            end = begin + entityCharacters.length;
            diff = 1 + end - begin;
            terminated = false;
          }

          // If the reference is not terminated, warn.
          if (!terminated) {
            reason = entityCharacters ? namedNotTerminated : namedEmpty;

            if (settings.attribute) {
              following = value.charCodeAt(end);

              if (following === equalsTo) {
                warning(reason, diff);
                entity = null;
              } else if (alphanumerical$1(following)) {
                entity = null;
              } else {
                warning(reason, diff);
              }
            } else {
              warning(reason, diff);
            }
          }
        }

        reference = entity;
      } else {
        if (!terminated) {
          // All non-terminated numeric entities are not rendered, and trigger a
          // warning.
          warning(numericNotTerminated, diff);
        }

        // When terminated and number, parse as either hexadecimal or decimal.
        reference = parseInt(characters, bases[type]);

        // Trigger a warning when the parsed number is prohibited, and replace
        // with replacement character.
        if (prohibited(reference)) {
          warning(numericProhibited, diff);
          reference = fromCharCode(replacementCharacter);
        } else if (reference in invalid) {
          // Trigger a warning when the parsed number is disallowed, and replace
          // by an alternative.
          warning(numericDisallowed, diff);
          reference = invalid[reference];
        } else {
          // Parse the number.
          output = '';

          // Trigger a warning when the parsed number should not be used.
          if (disallowed(reference)) {
            warning(numericDisallowed, diff);
          }

          // Stringify the number.
          if (reference > 0xffff) {
            reference -= 0x10000;
            output += fromCharCode((reference >>> (10 & 0x3ff)) | 0xd800);
            reference = 0xdc00 | (reference & 0x3ff);
          }

          reference = output + fromCharCode(reference);
        }
      }

      // Found it!
      // First eat the queued characters as normal text, then eat an entity.
      if (reference) {
        flush();

        prev = now();
        index = end - 1;
        column += end - start + 1;
        result.push(reference);
        next = now();
        next.offset++;

        if (handleReference) {
          handleReference.call(
            referenceContext,
            reference,
            {start: prev, end: next},
            value.slice(start - 1, end)
          );
        }

        prev = next;
      } else {
        // If we could not find a reference, queue the checked characters (as
        // normal characters), and move the pointer to their end.
        // This is possible because we can be certain neither newlines nor
        // ampersands are included.
        characters = value.slice(start - 1, end);
        queue += characters;
        column += characters.length;
        index = end - 1;
      }
    } else {
      // Handle anything other than an ampersand, including newlines and EOF.
      if (
        character === 10 // Line feed
      ) {
        line++;
        lines++;
        column = 0;
      }

      if (character === character) {
        queue += fromCharCode(character);
        column++;
      } else {
        flush();
      }
    }
  }

  // Return the reduced nodes.
  return result.join('')

  // Get current position.
  function now() {
    return {
      line: line,
      column: column,
      offset: index + (pos.offset || 0)
    }
  }

  // “Throw” a parse-error: a warning.
  function parseError(code, offset) {
    var position = now();

    position.column += offset;
    position.offset += offset;

    handleWarning.call(warningContext, messages[code], position, code);
  }

  // Flush `queue` (normal text).
  // Macro invoked before each entity and at the end of `value`.
  // Does nothing when `queue` is empty.
  function flush() {
    if (queue) {
      result.push(queue);

      if (handleText) {
        handleText.call(textContext, queue, {start: prev, end: now()});
      }

      queue = '';
    }
  }
}

// Check if `character` is outside the permissible unicode range.
function prohibited(code) {
  return (code >= 0xd800 && code <= 0xdfff) || code > 0x10ffff
}

// Check if `character` is disallowed.
function disallowed(code) {
  return (
    (code >= 0x0001 && code <= 0x0008) ||
    code === 0x000b ||
    (code >= 0x000d && code <= 0x001f) ||
    (code >= 0x007f && code <= 0x009f) ||
    (code >= 0xfdd0 && code <= 0xfdef) ||
    (code & 0xffff) === 0xffff ||
    (code & 0xffff) === 0xfffe
  )
}

var xtend$3 = immutable;
var entities$1 = parseEntities_1;

var decode$1 = factory$3;

// Factory to create an entity decoder.
function factory$3(ctx) {
  decoder.raw = decodeRaw;

  return decoder

  // Normalize `position` to add an `indent`.
  function normalize(position) {
    var offsets = ctx.offset;
    var line = position.line;
    var result = [];

    while (++line) {
      if (!(line in offsets)) {
        break
      }

      result.push((offsets[line] || 0) + 1);
    }

    return {start: position, indent: result}
  }

  // Decode `value` (at `position`) into text-nodes.
  function decoder(value, position, handler) {
    entities$1(value, {
      position: normalize(position),
      warning: handleWarning,
      text: handler,
      reference: handler,
      textContext: ctx,
      referenceContext: ctx
    });
  }

  // Decode `value` (at `position`) into a string.
  function decodeRaw(value, position, options) {
    return entities$1(
      value,
      xtend$3(options, {position: normalize(position), warning: handleWarning})
    )
  }

  // Handle a warning.
  // See <https://github.com/wooorm/parse-entities> for the warnings.
  function handleWarning(reason, position, code) {
    if (code !== 3) {
      ctx.file.message(reason, position);
    }
  }
}

var tokenizer$3 = factory$2;

// Construct a tokenizer.  This creates both `tokenizeInline` and `tokenizeBlock`.
function factory$2(type) {
  return tokenize

  // Tokenizer for a bound `type`.
  function tokenize(value, location) {
    var self = this;
    var offset = self.offset;
    var tokens = [];
    var methods = self[type + 'Methods'];
    var tokenizers = self[type + 'Tokenizers'];
    var line = location.line;
    var column = location.column;
    var index;
    var length;
    var method;
    var name;
    var matched;
    var valueLength;

    // Trim white space only lines.
    if (!value) {
      return tokens
    }

    // Expose on `eat`.
    eat.now = now;
    eat.file = self.file;

    // Sync initial offset.
    updatePosition('');

    // Iterate over `value`, and iterate over all tokenizers.  When one eats
    // something, re-iterate with the remaining value.  If no tokenizer eats,
    // something failed (should not happen) and an exception is thrown.
    while (value) {
      index = -1;
      length = methods.length;
      matched = false;

      while (++index < length) {
        name = methods[index];
        method = tokenizers[name];

        // Previously, we had constructs such as footnotes and YAML that used
        // these properties.
        // Those are now external (plus there are userland extensions), that may
        // still use them.
        if (
          method &&
          /* istanbul ignore next */ (!method.onlyAtStart || self.atStart) &&
          /* istanbul ignore next */ (!method.notInList || !self.inList) &&
          /* istanbul ignore next */ (!method.notInBlock || !self.inBlock) &&
          (!method.notInLink || !self.inLink)
        ) {
          valueLength = value.length;

          method.apply(self, [eat, value]);

          matched = valueLength !== value.length;

          if (matched) {
            break
          }
        }
      }

      /* istanbul ignore if */
      if (!matched) {
        self.file.fail(new Error('Infinite loop'), eat.now());
      }
    }

    self.eof = now();

    return tokens

    // Update line, column, and offset based on `value`.
    function updatePosition(subvalue) {
      var lastIndex = -1;
      var index = subvalue.indexOf('\n');

      while (index !== -1) {
        line++;
        lastIndex = index;
        index = subvalue.indexOf('\n', index + 1);
      }

      if (lastIndex === -1) {
        column += subvalue.length;
      } else {
        column = subvalue.length - lastIndex;
      }

      if (line in offset) {
        if (lastIndex !== -1) {
          column += offset[line];
        } else if (column <= offset[line]) {
          column = offset[line] + 1;
        }
      }
    }

    // Get offset.  Called before the first character is eaten to retrieve the
    // range’s offsets.
    function getOffset() {
      var indentation = [];
      var pos = line + 1;

      // Done.  Called when the last character is eaten to retrieve the range’s
      // offsets.
      return function () {
        var last = line + 1;

        while (pos < last) {
          indentation.push((offset[pos] || 0) + 1);

          pos++;
        }

        return indentation
      }
    }

    // Get the current position.
    function now() {
      var pos = {line: line, column: column};

      pos.offset = self.toOffset(pos);

      return pos
    }

    // Store position information for a node.
    function Position(start) {
      this.start = start;
      this.end = now();
    }

    // Throw when a value is incorrectly eaten.  This shouldn’t happen but will
    // throw on new, incorrect rules.
    function validateEat(subvalue) {
      /* istanbul ignore if */
      if (value.slice(0, subvalue.length) !== subvalue) {
        // Capture stack-trace.
        self.file.fail(
          new Error(
            'Incorrectly eaten value: please report this warning on https://git.io/vg5Ft'
          ),
          now()
        );
      }
    }

    // Mark position and patch `node.position`.
    function position() {
      var before = now();

      return update

      // Add the position to a node.
      function update(node, indent) {
        var previous = node.position;
        var start = previous ? previous.start : before;
        var combined = [];
        var n = previous && previous.end.line;
        var l = before.line;

        node.position = new Position(start);

        // If there was already a `position`, this node was merged.  Fixing
        // `start` wasn’t hard, but the indent is different.  Especially
        // because some information, the indent between `n` and `l` wasn’t
        // tracked.  Luckily, that space is (should be?) empty, so we can
        // safely check for it now.
        if (previous && indent && previous.indent) {
          combined = previous.indent;

          if (n < l) {
            while (++n < l) {
              combined.push((offset[n] || 0) + 1);
            }

            combined.push(before.column);
          }

          indent = combined.concat(indent);
        }

        node.position.indent = indent || [];

        return node
      }
    }

    // Add `node` to `parent`s children or to `tokens`.  Performs merges where
    // possible.
    function add(node, parent) {
      var children = parent ? parent.children : tokens;
      var previous = children[children.length - 1];
      var fn;

      if (
        previous &&
        node.type === previous.type &&
        (node.type === 'text' || node.type === 'blockquote') &&
        mergeable(previous) &&
        mergeable(node)
      ) {
        fn = node.type === 'text' ? mergeText : mergeBlockquote;
        node = fn.call(self, previous, node);
      }

      if (node !== previous) {
        children.push(node);
      }

      if (self.atStart && tokens.length !== 0) {
        self.exitStart();
      }

      return node
    }

    // Remove `subvalue` from `value`.  `subvalue` must be at the start of
    // `value`.
    function eat(subvalue) {
      var indent = getOffset();
      var pos = position();
      var current = now();

      validateEat(subvalue);

      apply.reset = reset;
      reset.test = test;
      apply.test = test;

      value = value.slice(subvalue.length);

      updatePosition(subvalue);

      indent = indent();

      return apply

      // Add the given arguments, add `position` to the returned node, and
      // return the node.
      function apply(node, parent) {
        return pos(add(pos(node), parent), indent)
      }

      // Functions just like apply, but resets the content: the line and
      // column are reversed, and the eaten value is re-added.   This is
      // useful for nodes with a single type of content, such as lists and
      // tables.  See `apply` above for what parameters are expected.
      function reset() {
        var node = apply.apply(null, arguments);

        line = current.line;
        column = current.column;
        value = subvalue + value;

        return node
      }

      // Test the position, after eating, and reverse to a not-eaten state.
      function test() {
        var result = pos({});

        line = current.line;
        column = current.column;
        value = subvalue + value;

        return result.position
      }
    }
  }
}

// Check whether a node is mergeable with adjacent nodes.
function mergeable(node) {
  var start;
  var end;

  if (node.type !== 'text' || !node.position) {
    return true
  }

  start = node.position.start;
  end = node.position.end;

  // Only merge nodes which occupy the same size as their `value`.
  return (
    start.line !== end.line || end.column - start.column === node.value.length
  )
}

// Merge two text nodes: `node` into `prev`.
function mergeText(previous, node) {
  previous.value += node.value;

  return previous
}

// Merge two blockquotes: `node` into `prev`, unless in CommonMark or gfm modes.
function mergeBlockquote(previous, node) {
  if (this.options.commonmark || this.options.gfm) {
    return node
  }

  previous.children = previous.children.concat(node.children);

  return previous
}

var markdownEscapes;
var hasRequiredMarkdownEscapes;

function requireMarkdownEscapes () {
	if (hasRequiredMarkdownEscapes) return markdownEscapes;
	hasRequiredMarkdownEscapes = 1;

	markdownEscapes = escapes;

	var defaults = [
	  '\\',
	  '`',
	  '*',
	  '{',
	  '}',
	  '[',
	  ']',
	  '(',
	  ')',
	  '#',
	  '+',
	  '-',
	  '.',
	  '!',
	  '_',
	  '>'
	];

	var gfm = defaults.concat(['~', '|']);

	var commonmark = gfm.concat([
	  '\n',
	  '"',
	  '$',
	  '%',
	  '&',
	  "'",
	  ',',
	  '/',
	  ':',
	  ';',
	  '<',
	  '=',
	  '?',
	  '@',
	  '^'
	]);

	escapes.default = defaults;
	escapes.gfm = gfm;
	escapes.commonmark = commonmark;

	// Get markdown escapes.
	function escapes(options) {
	  var settings = options || {};

	  if (settings.commonmark) {
	    return commonmark
	  }

	  return settings.gfm ? gfm : defaults
	}
	return markdownEscapes;
}

var blockElements;
var hasRequiredBlockElements;

function requireBlockElements () {
	if (hasRequiredBlockElements) return blockElements;
	hasRequiredBlockElements = 1;

	blockElements = [
	  'address',
	  'article',
	  'aside',
	  'base',
	  'basefont',
	  'blockquote',
	  'body',
	  'caption',
	  'center',
	  'col',
	  'colgroup',
	  'dd',
	  'details',
	  'dialog',
	  'dir',
	  'div',
	  'dl',
	  'dt',
	  'fieldset',
	  'figcaption',
	  'figure',
	  'footer',
	  'form',
	  'frame',
	  'frameset',
	  'h1',
	  'h2',
	  'h3',
	  'h4',
	  'h5',
	  'h6',
	  'head',
	  'header',
	  'hgroup',
	  'hr',
	  'html',
	  'iframe',
	  'legend',
	  'li',
	  'link',
	  'main',
	  'menu',
	  'menuitem',
	  'meta',
	  'nav',
	  'noframes',
	  'ol',
	  'optgroup',
	  'option',
	  'p',
	  'param',
	  'pre',
	  'section',
	  'source',
	  'title',
	  'summary',
	  'table',
	  'tbody',
	  'td',
	  'tfoot',
	  'th',
	  'thead',
	  'title',
	  'tr',
	  'track',
	  'ul'
	];
	return blockElements;
}

var defaults$1;
var hasRequiredDefaults;

function requireDefaults () {
	if (hasRequiredDefaults) return defaults$1;
	hasRequiredDefaults = 1;

	defaults$1 = {
	  position: true,
	  gfm: true,
	  commonmark: false,
	  pedantic: false,
	  blocks: requireBlockElements()
	};
	return defaults$1;
}

var setOptions_1;
var hasRequiredSetOptions;

function requireSetOptions () {
	if (hasRequiredSetOptions) return setOptions_1;
	hasRequiredSetOptions = 1;

	var xtend = immutable;
	var escapes = requireMarkdownEscapes();
	var defaults = requireDefaults();

	setOptions_1 = setOptions;

	function setOptions(options) {
	  var self = this;
	  var current = self.options;
	  var key;
	  var value;

	  if (options == null) {
	    options = {};
	  } else if (typeof options === 'object') {
	    options = xtend(options);
	  } else {
	    throw new Error('Invalid value `' + options + '` for setting `options`')
	  }

	  for (key in defaults) {
	    value = options[key];

	    if (value == null) {
	      value = current[key];
	    }

	    if (
	      (key !== 'blocks' && typeof value !== 'boolean') ||
	      (key === 'blocks' && typeof value !== 'object')
	    ) {
	      throw new Error(
	        'Invalid value `' + value + '` for setting `options.' + key + '`'
	      )
	    }

	    options[key] = value;
	  }

	  self.options = options;
	  self.escape = escapes(options);

	  return self
	}
	return setOptions_1;
}

var convert_1 = convert$2;

function convert$2(test) {
  if (test == null) {
    return ok$1
  }

  if (typeof test === 'string') {
    return typeFactory$1(test)
  }

  if (typeof test === 'object') {
    return 'length' in test ? anyFactory$1(test) : allFactory(test)
  }

  if (typeof test === 'function') {
    return test
  }

  throw new Error('Expected function, string, or object as test')
}

// Utility assert each property in `test` is represented in `node`, and each
// values are strictly equal.
function allFactory(test) {
  return all

  function all(node) {
    var key;

    for (key in test) {
      if (node[key] !== test[key]) return false
    }

    return true
  }
}

function anyFactory$1(tests) {
  var checks = [];
  var index = -1;

  while (++index < tests.length) {
    checks[index] = convert$2(tests[index]);
  }

  return any

  function any() {
    var index = -1;

    while (++index < checks.length) {
      if (checks[index].apply(this, arguments)) {
        return true
      }
    }

    return false
  }
}

// Utility to convert a string into a function which checks a given node’s type
// for said string.
function typeFactory$1(test) {
  return type

  function type(node) {
    return Boolean(node && node.type === test)
  }
}

// Utility to return true.
function ok$1() {
  return true
}

var color_1 = color$1;
function color$1(d) {
  return '\u001B[33m' + d + '\u001B[39m'
}

var unistUtilVisitParents = visitParents$1;

var convert$1 = convert_1;
var color = color_1;

var CONTINUE$1 = true;
var SKIP$1 = 'skip';
var EXIT$1 = false;

visitParents$1.CONTINUE = CONTINUE$1;
visitParents$1.SKIP = SKIP$1;
visitParents$1.EXIT = EXIT$1;

function visitParents$1(tree, test, visitor, reverse) {
  var step;
  var is;

  if (typeof test === 'function' && typeof visitor !== 'function') {
    reverse = visitor;
    visitor = test;
    test = null;
  }

  is = convert$1(test);
  step = reverse ? -1 : 1;

  factory(tree, null, [])();

  function factory(node, index, parents) {
    var value = typeof node === 'object' && node !== null ? node : {};
    var name;

    if (typeof value.type === 'string') {
      name =
        typeof value.tagName === 'string'
          ? value.tagName
          : typeof value.name === 'string'
          ? value.name
          : undefined;

      visit.displayName =
        'node (' + color(value.type + (name ? '<' + name + '>' : '')) + ')';
    }

    return visit

    function visit() {
      var grandparents = parents.concat(node);
      var result = [];
      var subresult;
      var offset;

      if (!test || is(node, index, parents[parents.length - 1] || null)) {
        result = toResult(visitor(node, parents));

        if (result[0] === EXIT$1) {
          return result
        }
      }

      if (node.children && result[0] !== SKIP$1) {
        offset = (reverse ? node.children.length : -1) + step;

        while (offset > -1 && offset < node.children.length) {
          subresult = factory(node.children[offset], offset, grandparents)();

          if (subresult[0] === EXIT$1) {
            return subresult
          }

          offset =
            typeof subresult[1] === 'number' ? subresult[1] : offset + step;
        }
      }

      return result
    }
  }
}

function toResult(value) {
  if (value !== null && typeof value === 'object' && 'length' in value) {
    return value
  }

  if (typeof value === 'number') {
    return [CONTINUE$1, value]
  }

  return [value]
}

var unistUtilVisit = visit$5;

var visitParents = unistUtilVisitParents;

var CONTINUE = visitParents.CONTINUE;
var SKIP = visitParents.SKIP;
var EXIT = visitParents.EXIT;

visit$5.CONTINUE = CONTINUE;
visit$5.SKIP = SKIP;
visit$5.EXIT = EXIT;

function visit$5(tree, test, visitor, reverse) {
  if (typeof test === 'function' && typeof visitor !== 'function') {
    reverse = visitor;
    visitor = test;
    test = null;
  }

  visitParents(tree, test, overload, reverse);

  function overload(node, parents) {
    var parent = parents[parents.length - 1];
    var index = parent ? parent.children.indexOf(node) : null;
    return visitor(node, index, parent)
  }
}

var visit$6 = /*@__PURE__*/getDefaultExportFromCjs(unistUtilVisit);

var unistUtilRemovePosition;
var hasRequiredUnistUtilRemovePosition;

function requireUnistUtilRemovePosition () {
	if (hasRequiredUnistUtilRemovePosition) return unistUtilRemovePosition;
	hasRequiredUnistUtilRemovePosition = 1;

	var visit = unistUtilVisit;

	unistUtilRemovePosition = removePosition;

	function removePosition(node, force) {
	  visit(node, force ? hard : soft);
	  return node
	}

	function hard(node) {
	  delete node.position;
	}

	function soft(node) {
	  node.position = undefined;
	}
	return unistUtilRemovePosition;
}

var parse_1;
var hasRequiredParse;

function requireParse () {
	if (hasRequiredParse) return parse_1;
	hasRequiredParse = 1;

	var xtend = immutable;
	var removePosition = requireUnistUtilRemovePosition();

	parse_1 = parse;

	var lineFeed = '\n';
	var lineBreaksExpression = /\r\n|\r/g;

	// Parse the bound file.
	function parse() {
	  var self = this;
	  var value = String(self.file);
	  var start = {line: 1, column: 1, offset: 0};
	  var content = xtend(start);
	  var node;

	  // Clean non-unix newlines: `\r\n` and `\r` are all changed to `\n`.
	  // This should not affect positional information.
	  value = value.replace(lineBreaksExpression, lineFeed);

	  // BOM.
	  if (value.charCodeAt(0) === 0xfeff) {
	    value = value.slice(1);

	    content.column++;
	    content.offset++;
	  }

	  node = {
	    type: 'root',
	    children: self.tokenizeBlock(value, content),
	    position: {start: start, end: self.eof || xtend(start)}
	  };

	  if (!self.options.position) {
	    removePosition(node, true);
	  }

	  return node
	}
	return parse_1;
}

var blankLine_1;
var hasRequiredBlankLine;

function requireBlankLine () {
	if (hasRequiredBlankLine) return blankLine_1;
	hasRequiredBlankLine = 1;

	// A line containing no characters, or a line containing only spaces (U+0020) or
	// tabs (U+0009), is called a blank line.
	// See <https://spec.commonmark.org/0.29/#blank-line>.
	var reBlankLine = /^[ \t]*(\n|$)/;

	// Note that though blank lines play a special role in lists to determine
	// whether the list is tight or loose
	// (<https://spec.commonmark.org/0.29/#blank-lines>), it’s done by the list
	// tokenizer and this blank line tokenizer does not have to be responsible for
	// that.
	// Therefore, configs such as `blankLine.notInList` do not have to be set here.
	blankLine_1 = blankLine;

	function blankLine(eat, value, silent) {
	  var match;
	  var subvalue = '';
	  var index = 0;
	  var length = value.length;

	  while (index < length) {
	    match = reBlankLine.exec(value.slice(index));

	    if (match == null) {
	      break
	    }

	    index += match[0].length;
	    subvalue += match[0];
	  }

	  if (subvalue === '') {
	    return
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  eat(subvalue);
	}
	return blankLine_1;
}

/*!
 * repeat-string <https://github.com/jonschlinkert/repeat-string>
 *
 * Copyright (c) 2014-2015, Jon Schlinkert.
 * Licensed under the MIT License.
 */

/**
 * Results cache
 */

var res = '';
var cache;

/**
 * Expose `repeat`
 */

var repeatString = repeat$2;

/**
 * Repeat the given `string` the specified `number`
 * of times.
 *
 * **Example:**
 *
 * ```js
 * var repeat = require('repeat-string');
 * repeat('A', 5);
 * //=> AAAAA
 * ```
 *
 * @param {String} `string` The string to repeat
 * @param {Number} `number` The number of times to repeat the string
 * @return {String} Repeated string
 * @api public
 */

function repeat$2(str, num) {
  if (typeof str !== 'string') {
    throw new TypeError('expected a string');
  }

  // cover common, quick use cases
  if (num === 1) return str;
  if (num === 2) return str + str;

  var max = str.length * num;
  if (cache !== str || typeof cache === 'undefined') {
    cache = str;
    res = '';
  } else if (res.length >= max) {
    return res.substr(0, max);
  }

  while (max > res.length && num > 1) {
    if (num & 1) {
      res += str;
    }

    num >>= 1;
    str += str;
  }

  res += str;
  res = res.substr(0, max);
  return res;
}

var trimTrailingLines_1;
var hasRequiredTrimTrailingLines;

function requireTrimTrailingLines () {
	if (hasRequiredTrimTrailingLines) return trimTrailingLines_1;
	hasRequiredTrimTrailingLines = 1;

	trimTrailingLines_1 = trimTrailingLines;

	var line = '\n';

	// Remove final newline characters from `value`.
	function trimTrailingLines(value) {
	  var val = String(value);
	  var index = val.length;

	  while (val.charAt(--index) === line) {
	    // Empty
	  }

	  return val.slice(0, index + 1)
	}
	return trimTrailingLines_1;
}

var codeIndented;
var hasRequiredCodeIndented;

function requireCodeIndented () {
	if (hasRequiredCodeIndented) return codeIndented;
	hasRequiredCodeIndented = 1;

	var repeat = repeatString;
	var trim = requireTrimTrailingLines();

	codeIndented = indentedCode;

	var lineFeed = '\n';
	var tab = '\t';
	var space = ' ';

	var tabSize = 4;
	var codeIndent = repeat(space, tabSize);

	function indentedCode(eat, value, silent) {
	  var index = -1;
	  var length = value.length;
	  var subvalue = '';
	  var content = '';
	  var subvalueQueue = '';
	  var contentQueue = '';
	  var character;
	  var blankQueue;
	  var indent;

	  while (++index < length) {
	    character = value.charAt(index);

	    if (indent) {
	      indent = false;

	      subvalue += subvalueQueue;
	      content += contentQueue;
	      subvalueQueue = '';
	      contentQueue = '';

	      if (character === lineFeed) {
	        subvalueQueue = character;
	        contentQueue = character;
	      } else {
	        subvalue += character;
	        content += character;

	        while (++index < length) {
	          character = value.charAt(index);

	          if (!character || character === lineFeed) {
	            contentQueue = character;
	            subvalueQueue = character;
	            break
	          }

	          subvalue += character;
	          content += character;
	        }
	      }
	    } else if (
	      character === space &&
	      value.charAt(index + 1) === character &&
	      value.charAt(index + 2) === character &&
	      value.charAt(index + 3) === character
	    ) {
	      subvalueQueue += codeIndent;
	      index += 3;
	      indent = true;
	    } else if (character === tab) {
	      subvalueQueue += character;
	      indent = true;
	    } else {
	      blankQueue = '';

	      while (character === tab || character === space) {
	        blankQueue += character;
	        character = value.charAt(++index);
	      }

	      if (character !== lineFeed) {
	        break
	      }

	      subvalueQueue += blankQueue + character;
	      contentQueue += character;
	    }
	  }

	  if (content) {
	    if (silent) {
	      return true
	    }

	    return eat(subvalue)({
	      type: 'code',
	      lang: null,
	      meta: null,
	      value: trim(content)
	    })
	  }
	}
	return codeIndented;
}

var codeFenced;
var hasRequiredCodeFenced;

function requireCodeFenced () {
	if (hasRequiredCodeFenced) return codeFenced;
	hasRequiredCodeFenced = 1;

	codeFenced = fencedCode;

	var lineFeed = '\n';
	var tab = '\t';
	var space = ' ';
	var tilde = '~';
	var graveAccent = '`';

	var minFenceCount = 3;
	var tabSize = 4;

	function fencedCode(eat, value, silent) {
	  var self = this;
	  var gfm = self.options.gfm;
	  var length = value.length + 1;
	  var index = 0;
	  var subvalue = '';
	  var fenceCount;
	  var marker;
	  var character;
	  var flag;
	  var lang;
	  var meta;
	  var queue;
	  var content;
	  var exdentedContent;
	  var closing;
	  var exdentedClosing;
	  var indent;
	  var now;

	  if (!gfm) {
	    return
	  }

	  // Eat initial spacing.
	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== space && character !== tab) {
	      break
	    }

	    subvalue += character;
	    index++;
	  }

	  indent = index;

	  // Eat the fence.
	  character = value.charAt(index);

	  if (character !== tilde && character !== graveAccent) {
	    return
	  }

	  index++;
	  marker = character;
	  fenceCount = 1;
	  subvalue += character;

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== marker) {
	      break
	    }

	    subvalue += character;
	    fenceCount++;
	    index++;
	  }

	  if (fenceCount < minFenceCount) {
	    return
	  }

	  // Eat spacing before flag.
	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== space && character !== tab) {
	      break
	    }

	    subvalue += character;
	    index++;
	  }

	  // Eat flag.
	  flag = '';
	  queue = '';

	  while (index < length) {
	    character = value.charAt(index);

	    if (
	      character === lineFeed ||
	      (marker === graveAccent && character === marker)
	    ) {
	      break
	    }

	    if (character === space || character === tab) {
	      queue += character;
	    } else {
	      flag += queue + character;
	      queue = '';
	    }

	    index++;
	  }

	  character = value.charAt(index);

	  if (character && character !== lineFeed) {
	    return
	  }

	  if (silent) {
	    return true
	  }

	  now = eat.now();
	  now.column += subvalue.length;
	  now.offset += subvalue.length;

	  subvalue += flag;
	  flag = self.decode.raw(self.unescape(flag), now);

	  if (queue) {
	    subvalue += queue;
	  }

	  queue = '';
	  closing = '';
	  exdentedClosing = '';
	  content = '';
	  exdentedContent = '';
	  var skip = true;

	  // Eat content.
	  while (index < length) {
	    character = value.charAt(index);
	    content += closing;
	    exdentedContent += exdentedClosing;
	    closing = '';
	    exdentedClosing = '';

	    if (character !== lineFeed) {
	      content += character;
	      exdentedClosing += character;
	      index++;
	      continue
	    }

	    // The first line feed is ignored. Others aren’t.
	    if (skip) {
	      subvalue += character;
	      skip = false;
	    } else {
	      closing += character;
	      exdentedClosing += character;
	    }

	    queue = '';
	    index++;

	    while (index < length) {
	      character = value.charAt(index);

	      if (character !== space) {
	        break
	      }

	      queue += character;
	      index++;
	    }

	    closing += queue;
	    exdentedClosing += queue.slice(indent);

	    if (queue.length >= tabSize) {
	      continue
	    }

	    queue = '';

	    while (index < length) {
	      character = value.charAt(index);

	      if (character !== marker) {
	        break
	      }

	      queue += character;
	      index++;
	    }

	    closing += queue;
	    exdentedClosing += queue;

	    if (queue.length < fenceCount) {
	      continue
	    }

	    queue = '';

	    while (index < length) {
	      character = value.charAt(index);

	      if (character !== space && character !== tab) {
	        break
	      }

	      closing += character;
	      exdentedClosing += character;
	      index++;
	    }

	    if (!character || character === lineFeed) {
	      break
	    }
	  }

	  subvalue += content + closing;

	  // Get lang and meta from the flag.
	  index = -1;
	  length = flag.length;

	  while (++index < length) {
	    character = flag.charAt(index);

	    if (character === space || character === tab) {
	      if (!lang) {
	        lang = flag.slice(0, index);
	      }
	    } else if (lang) {
	      meta = flag.slice(index);
	      break
	    }
	  }

	  return eat(subvalue)({
	    type: 'code',
	    lang: lang || flag || null,
	    meta: meta || null,
	    value: exdentedContent
	  })
	}
	return codeFenced;
}

var trim = {exports: {}};

trim.exports;

var hasRequiredTrim;

function requireTrim () {
	if (hasRequiredTrim) return trim.exports;
	hasRequiredTrim = 1;
	(function (module, exports) {
		exports = module.exports = trim;

		function trim(str){
		  return str.replace(/^\s*|\s*$/g, '');
		}

		exports.left = function(str){
		  return str.replace(/^\s*/, '');
		};

		exports.right = function(str){
		  return str.replace(/\s*$/, '');
		}; 
	} (trim, trim.exports));
	return trim.exports;
}

var interrupt_1;
var hasRequiredInterrupt;

function requireInterrupt () {
	if (hasRequiredInterrupt) return interrupt_1;
	hasRequiredInterrupt = 1;

	interrupt_1 = interrupt;

	function interrupt(interruptors, tokenizers, ctx, parameters) {
	  var length = interruptors.length;
	  var index = -1;
	  var interruptor;
	  var config;

	  while (++index < length) {
	    interruptor = interruptors[index];
	    config = interruptor[1] || {};

	    if (
	      config.pedantic !== undefined &&
	      config.pedantic !== ctx.options.pedantic
	    ) {
	      continue
	    }

	    if (
	      config.commonmark !== undefined &&
	      config.commonmark !== ctx.options.commonmark
	    ) {
	      continue
	    }

	    if (tokenizers[interruptor[0]].apply(ctx, parameters)) {
	      return true
	    }
	  }

	  return false
	}
	return interrupt_1;
}

var blockquote_1$1;
var hasRequiredBlockquote;

function requireBlockquote () {
	if (hasRequiredBlockquote) return blockquote_1$1;
	hasRequiredBlockquote = 1;

	var trim = requireTrim();
	var interrupt = requireInterrupt();

	blockquote_1$1 = blockquote;

	var lineFeed = '\n';
	var tab = '\t';
	var space = ' ';
	var greaterThan = '>';

	function blockquote(eat, value, silent) {
	  var self = this;
	  var offsets = self.offset;
	  var tokenizers = self.blockTokenizers;
	  var interruptors = self.interruptBlockquote;
	  var now = eat.now();
	  var currentLine = now.line;
	  var length = value.length;
	  var values = [];
	  var contents = [];
	  var indents = [];
	  var add;
	  var index = 0;
	  var character;
	  var rest;
	  var nextIndex;
	  var content;
	  var line;
	  var startIndex;
	  var prefixed;
	  var exit;

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== space && character !== tab) {
	      break
	    }

	    index++;
	  }

	  if (value.charAt(index) !== greaterThan) {
	    return
	  }

	  if (silent) {
	    return true
	  }

	  index = 0;

	  while (index < length) {
	    nextIndex = value.indexOf(lineFeed, index);
	    startIndex = index;
	    prefixed = false;

	    if (nextIndex === -1) {
	      nextIndex = length;
	    }

	    while (index < length) {
	      character = value.charAt(index);

	      if (character !== space && character !== tab) {
	        break
	      }

	      index++;
	    }

	    if (value.charAt(index) === greaterThan) {
	      index++;
	      prefixed = true;

	      if (value.charAt(index) === space) {
	        index++;
	      }
	    } else {
	      index = startIndex;
	    }

	    content = value.slice(index, nextIndex);

	    if (!prefixed && !trim(content)) {
	      index = startIndex;
	      break
	    }

	    if (!prefixed) {
	      rest = value.slice(index);

	      // Check if the following code contains a possible block.
	      if (interrupt(interruptors, tokenizers, self, [eat, rest, true])) {
	        break
	      }
	    }

	    line = startIndex === index ? content : value.slice(startIndex, nextIndex);

	    indents.push(index - startIndex);
	    values.push(line);
	    contents.push(content);

	    index = nextIndex + 1;
	  }

	  index = -1;
	  length = indents.length;
	  add = eat(values.join(lineFeed));

	  while (++index < length) {
	    offsets[currentLine] = (offsets[currentLine] || 0) + indents[index];
	    currentLine++;
	  }

	  exit = self.enterBlock();
	  contents = self.tokenizeBlock(contents.join(lineFeed), now);
	  exit();

	  return add({type: 'blockquote', children: contents})
	}
	return blockquote_1$1;
}

var headingAtx;
var hasRequiredHeadingAtx;

function requireHeadingAtx () {
	if (hasRequiredHeadingAtx) return headingAtx;
	hasRequiredHeadingAtx = 1;

	headingAtx = atxHeading;

	var lineFeed = '\n';
	var tab = '\t';
	var space = ' ';
	var numberSign = '#';

	var maxFenceCount = 6;

	function atxHeading(eat, value, silent) {
	  var self = this;
	  var pedantic = self.options.pedantic;
	  var length = value.length + 1;
	  var index = -1;
	  var now = eat.now();
	  var subvalue = '';
	  var content = '';
	  var character;
	  var queue;
	  var depth;

	  // Eat initial spacing.
	  while (++index < length) {
	    character = value.charAt(index);

	    if (character !== space && character !== tab) {
	      index--;
	      break
	    }

	    subvalue += character;
	  }

	  // Eat hashes.
	  depth = 0;

	  while (++index <= length) {
	    character = value.charAt(index);

	    if (character !== numberSign) {
	      index--;
	      break
	    }

	    subvalue += character;
	    depth++;
	  }

	  if (depth > maxFenceCount) {
	    return
	  }

	  if (!depth || (!pedantic && value.charAt(index + 1) === numberSign)) {
	    return
	  }

	  length = value.length + 1;

	  // Eat intermediate white-space.
	  queue = '';

	  while (++index < length) {
	    character = value.charAt(index);

	    if (character !== space && character !== tab) {
	      index--;
	      break
	    }

	    queue += character;
	  }

	  // Exit when not in pedantic mode without spacing.
	  if (!pedantic && queue.length === 0 && character && character !== lineFeed) {
	    return
	  }

	  if (silent) {
	    return true
	  }

	  // Eat content.
	  subvalue += queue;
	  queue = '';
	  content = '';

	  while (++index < length) {
	    character = value.charAt(index);

	    if (!character || character === lineFeed) {
	      break
	    }

	    if (character !== space && character !== tab && character !== numberSign) {
	      content += queue + character;
	      queue = '';
	      continue
	    }

	    while (character === space || character === tab) {
	      queue += character;
	      character = value.charAt(++index);
	    }

	    // `#` without a queue is part of the content.
	    if (!pedantic && content && !queue && character === numberSign) {
	      content += character;
	      continue
	    }

	    while (character === numberSign) {
	      queue += character;
	      character = value.charAt(++index);
	    }

	    while (character === space || character === tab) {
	      queue += character;
	      character = value.charAt(++index);
	    }

	    index--;
	  }

	  now.column += subvalue.length;
	  now.offset += subvalue.length;
	  subvalue += content + queue;

	  return eat(subvalue)({
	    type: 'heading',
	    depth: depth,
	    children: self.tokenizeInline(content, now)
	  })
	}
	return headingAtx;
}

var thematicBreak_1$1;
var hasRequiredThematicBreak;

function requireThematicBreak () {
	if (hasRequiredThematicBreak) return thematicBreak_1$1;
	hasRequiredThematicBreak = 1;

	thematicBreak_1$1 = thematicBreak;

	var tab = '\t';
	var lineFeed = '\n';
	var space = ' ';
	var asterisk = '*';
	var dash = '-';
	var underscore = '_';

	var maxCount = 3;

	function thematicBreak(eat, value, silent) {
	  var index = -1;
	  var length = value.length + 1;
	  var subvalue = '';
	  var character;
	  var marker;
	  var markerCount;
	  var queue;

	  while (++index < length) {
	    character = value.charAt(index);

	    if (character !== tab && character !== space) {
	      break
	    }

	    subvalue += character;
	  }

	  if (
	    character !== asterisk &&
	    character !== dash &&
	    character !== underscore
	  ) {
	    return
	  }

	  marker = character;
	  subvalue += character;
	  markerCount = 1;
	  queue = '';

	  while (++index < length) {
	    character = value.charAt(index);

	    if (character === marker) {
	      markerCount++;
	      subvalue += queue + marker;
	      queue = '';
	    } else if (character === space) {
	      queue += character;
	    } else if (
	      markerCount >= maxCount &&
	      (!character || character === lineFeed)
	    ) {
	      subvalue += queue;

	      if (silent) {
	        return true
	      }

	      return eat(subvalue)({type: 'thematicBreak'})
	    } else {
	      return
	    }
	  }
	}
	return thematicBreak_1$1;
}

var getIndentation;
var hasRequiredGetIndentation;

function requireGetIndentation () {
	if (hasRequiredGetIndentation) return getIndentation;
	hasRequiredGetIndentation = 1;

	getIndentation = indentation;

	var tab = '\t';
	var space = ' ';

	var spaceSize = 1;
	var tabSize = 4;

	// Gets indentation information for a line.
	function indentation(value) {
	  var index = 0;
	  var indent = 0;
	  var character = value.charAt(index);
	  var stops = {};
	  var size;
	  var lastIndent = 0;

	  while (character === tab || character === space) {
	    size = character === tab ? tabSize : spaceSize;

	    indent += size;

	    if (size > 1) {
	      indent = Math.floor(indent / size) * size;
	    }

	    while (lastIndent < indent) {
	      stops[++lastIndent] = index;
	    }

	    character = value.charAt(++index);
	  }

	  return {indent: indent, stops: stops}
	}
	return getIndentation;
}

var removeIndentation;
var hasRequiredRemoveIndentation;

function requireRemoveIndentation () {
	if (hasRequiredRemoveIndentation) return removeIndentation;
	hasRequiredRemoveIndentation = 1;

	var trim = requireTrim();
	var repeat = repeatString;
	var getIndent = requireGetIndentation();

	removeIndentation = indentation;

	var lineFeed = '\n';
	var space = ' ';
	var exclamationMark = '!';

	// Remove the minimum indent from every line in `value`.  Supports both tab,
	// spaced, and mixed indentation (as well as possible).
	function indentation(value, maximum) {
	  var values = value.split(lineFeed);
	  var position = values.length + 1;
	  var minIndent = Infinity;
	  var matrix = [];
	  var index;
	  var indentation;
	  var stops;

	  values.unshift(repeat(space, maximum) + exclamationMark);

	  while (position--) {
	    indentation = getIndent(values[position]);

	    matrix[position] = indentation.stops;

	    if (trim(values[position]).length === 0) {
	      continue
	    }

	    if (indentation.indent) {
	      if (indentation.indent > 0 && indentation.indent < minIndent) {
	        minIndent = indentation.indent;
	      }
	    } else {
	      minIndent = Infinity;

	      break
	    }
	  }

	  if (minIndent !== Infinity) {
	    position = values.length;

	    while (position--) {
	      stops = matrix[position];
	      index = minIndent;

	      while (index && !(index in stops)) {
	        index--;
	      }

	      values[position] = values[position].slice(stops[index] + 1);
	    }
	  }

	  values.shift();

	  return values.join(lineFeed)
	}
	return removeIndentation;
}

var list_1$1;
var hasRequiredList;

function requireList () {
	if (hasRequiredList) return list_1$1;
	hasRequiredList = 1;

	var trim = requireTrim();
	var repeat = repeatString;
	var decimal = requireIsDecimal();
	var getIndent = requireGetIndentation();
	var removeIndent = requireRemoveIndentation();
	var interrupt = requireInterrupt();

	list_1$1 = list;

	var asterisk = '*';
	var underscore = '_';
	var plusSign = '+';
	var dash = '-';
	var dot = '.';
	var space = ' ';
	var lineFeed = '\n';
	var tab = '\t';
	var rightParenthesis = ')';
	var lowercaseX = 'x';

	var tabSize = 4;
	var looseListItemExpression = /\n\n(?!\s*$)/;
	var taskItemExpression = /^\[([ X\tx])][ \t]/;
	var bulletExpression = /^([ \t]*)([*+-]|\d+[.)])( {1,4}(?! )| |\t|$|(?=\n))([^\n]*)/;
	var pedanticBulletExpression = /^([ \t]*)([*+-]|\d+[.)])([ \t]+)/;
	var initialIndentExpression = /^( {1,4}|\t)?/gm;

	function list(eat, value, silent) {
	  var self = this;
	  var commonmark = self.options.commonmark;
	  var pedantic = self.options.pedantic;
	  var tokenizers = self.blockTokenizers;
	  var interuptors = self.interruptList;
	  var index = 0;
	  var length = value.length;
	  var start = null;
	  var size;
	  var queue;
	  var ordered;
	  var character;
	  var marker;
	  var nextIndex;
	  var startIndex;
	  var prefixed;
	  var currentMarker;
	  var content;
	  var line;
	  var previousEmpty;
	  var empty;
	  var items;
	  var allLines;
	  var emptyLines;
	  var item;
	  var enterTop;
	  var exitBlockquote;
	  var spread = false;
	  var node;
	  var now;
	  var end;
	  var indented;

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== tab && character !== space) {
	      break
	    }

	    index++;
	  }

	  character = value.charAt(index);

	  if (character === asterisk || character === plusSign || character === dash) {
	    marker = character;
	    ordered = false;
	  } else {
	    ordered = true;
	    queue = '';

	    while (index < length) {
	      character = value.charAt(index);

	      if (!decimal(character)) {
	        break
	      }

	      queue += character;
	      index++;
	    }

	    character = value.charAt(index);

	    if (
	      !queue ||
	      !(character === dot || (commonmark && character === rightParenthesis))
	    ) {
	      return
	    }

	    /* Slightly abusing `silent` mode, whose goal is to make interrupting
	     * paragraphs work.
	     * Well, that’s exactly what we want to do here: don’t interrupt:
	     * 2. here, because the “list” doesn’t start with `1`. */
	    if (silent && queue !== '1') {
	      return
	    }

	    start = parseInt(queue, 10);
	    marker = character;
	  }

	  character = value.charAt(++index);

	  if (
	    character !== space &&
	    character !== tab &&
	    (pedantic || (character !== lineFeed && character !== ''))
	  ) {
	    return
	  }

	  if (silent) {
	    return true
	  }

	  index = 0;
	  items = [];
	  allLines = [];
	  emptyLines = [];

	  while (index < length) {
	    nextIndex = value.indexOf(lineFeed, index);
	    startIndex = index;
	    prefixed = false;
	    indented = false;

	    if (nextIndex === -1) {
	      nextIndex = length;
	    }

	    size = 0;

	    while (index < length) {
	      character = value.charAt(index);

	      if (character === tab) {
	        size += tabSize - (size % tabSize);
	      } else if (character === space) {
	        size++;
	      } else {
	        break
	      }

	      index++;
	    }

	    if (item && size >= item.indent) {
	      indented = true;
	    }

	    character = value.charAt(index);
	    currentMarker = null;

	    if (!indented) {
	      if (
	        character === asterisk ||
	        character === plusSign ||
	        character === dash
	      ) {
	        currentMarker = character;
	        index++;
	        size++;
	      } else {
	        queue = '';

	        while (index < length) {
	          character = value.charAt(index);

	          if (!decimal(character)) {
	            break
	          }

	          queue += character;
	          index++;
	        }

	        character = value.charAt(index);
	        index++;

	        if (
	          queue &&
	          (character === dot || (commonmark && character === rightParenthesis))
	        ) {
	          currentMarker = character;
	          size += queue.length + 1;
	        }
	      }

	      if (currentMarker) {
	        character = value.charAt(index);

	        if (character === tab) {
	          size += tabSize - (size % tabSize);
	          index++;
	        } else if (character === space) {
	          end = index + tabSize;

	          while (index < end) {
	            if (value.charAt(index) !== space) {
	              break
	            }

	            index++;
	            size++;
	          }

	          if (index === end && value.charAt(index) === space) {
	            index -= tabSize - 1;
	            size -= tabSize - 1;
	          }
	        } else if (character !== lineFeed && character !== '') {
	          currentMarker = null;
	        }
	      }
	    }

	    if (currentMarker) {
	      if (!pedantic && marker !== currentMarker) {
	        break
	      }

	      prefixed = true;
	    } else {
	      if (!commonmark && !indented && value.charAt(startIndex) === space) {
	        indented = true;
	      } else if (commonmark && item) {
	        indented = size >= item.indent || size > tabSize;
	      }

	      prefixed = false;
	      index = startIndex;
	    }

	    line = value.slice(startIndex, nextIndex);
	    content = startIndex === index ? line : value.slice(index, nextIndex);

	    if (
	      currentMarker === asterisk ||
	      currentMarker === underscore ||
	      currentMarker === dash
	    ) {
	      if (tokenizers.thematicBreak.call(self, eat, line, true)) {
	        break
	      }
	    }

	    previousEmpty = empty;
	    empty = !prefixed && !trim(content).length;

	    if (indented && item) {
	      item.value = item.value.concat(emptyLines, line);
	      allLines = allLines.concat(emptyLines, line);
	      emptyLines = [];
	    } else if (prefixed) {
	      if (emptyLines.length !== 0) {
	        spread = true;
	        item.value.push('');
	        item.trail = emptyLines.concat();
	      }

	      item = {
	        value: [line],
	        indent: size,
	        trail: []
	      };

	      items.push(item);
	      allLines = allLines.concat(emptyLines, line);
	      emptyLines = [];
	    } else if (empty) {
	      if (previousEmpty && !commonmark) {
	        break
	      }

	      emptyLines.push(line);
	    } else {
	      if (previousEmpty) {
	        break
	      }

	      if (interrupt(interuptors, tokenizers, self, [eat, line, true])) {
	        break
	      }

	      item.value = item.value.concat(emptyLines, line);
	      allLines = allLines.concat(emptyLines, line);
	      emptyLines = [];
	    }

	    index = nextIndex + 1;
	  }

	  node = eat(allLines.join(lineFeed)).reset({
	    type: 'list',
	    ordered: ordered,
	    start: start,
	    spread: spread,
	    children: []
	  });

	  enterTop = self.enterList();
	  exitBlockquote = self.enterBlock();
	  index = -1;
	  length = items.length;

	  while (++index < length) {
	    item = items[index].value.join(lineFeed);
	    now = eat.now();

	    eat(item)(listItem(self, item, now), node);

	    item = items[index].trail.join(lineFeed);

	    if (index !== length - 1) {
	      item += lineFeed;
	    }

	    eat(item);
	  }

	  enterTop();
	  exitBlockquote();

	  return node
	}

	function listItem(ctx, value, position) {
	  var offsets = ctx.offset;
	  var fn = ctx.options.pedantic ? pedanticListItem : normalListItem;
	  var checked = null;
	  var task;
	  var indent;

	  value = fn.apply(null, arguments);

	  if (ctx.options.gfm) {
	    task = value.match(taskItemExpression);

	    if (task) {
	      indent = task[0].length;
	      checked = task[1].toLowerCase() === lowercaseX;
	      offsets[position.line] += indent;
	      value = value.slice(indent);
	    }
	  }

	  return {
	    type: 'listItem',
	    spread: looseListItemExpression.test(value),
	    checked: checked,
	    children: ctx.tokenizeBlock(value, position)
	  }
	}

	// Create a list-item using overly simple mechanics.
	function pedanticListItem(ctx, value, position) {
	  var offsets = ctx.offset;
	  var line = position.line;

	  // Remove the list-item’s bullet.
	  value = value.replace(pedanticBulletExpression, replacer);

	  // The initial line was also matched by the below, so we reset the `line`.
	  line = position.line;

	  return value.replace(initialIndentExpression, replacer)

	  // A simple replacer which removed all matches, and adds their length to
	  // `offset`.
	  function replacer($0) {
	    offsets[line] = (offsets[line] || 0) + $0.length;
	    line++;

	    return ''
	  }
	}

	// Create a list-item using sane mechanics.
	function normalListItem(ctx, value, position) {
	  var offsets = ctx.offset;
	  var line = position.line;
	  var max;
	  var bullet;
	  var rest;
	  var lines;
	  var trimmedLines;
	  var index;
	  var length;

	  // Remove the list-item’s bullet.
	  value = value.replace(bulletExpression, replacer);

	  lines = value.split(lineFeed);

	  trimmedLines = removeIndent(value, getIndent(max).indent).split(lineFeed);

	  // We replaced the initial bullet with something else above, which was used
	  // to trick `removeIndentation` into removing some more characters when
	  // possible.  However, that could result in the initial line to be stripped
	  // more than it should be.
	  trimmedLines[0] = rest;

	  offsets[line] = (offsets[line] || 0) + bullet.length;
	  line++;

	  index = 0;
	  length = lines.length;

	  while (++index < length) {
	    offsets[line] =
	      (offsets[line] || 0) + lines[index].length - trimmedLines[index].length;
	    line++;
	  }

	  return trimmedLines.join(lineFeed)

	  /* eslint-disable-next-line max-params */
	  function replacer($0, $1, $2, $3, $4) {
	    bullet = $1 + $2 + $3;
	    rest = $4;

	    // Make sure that the first nine numbered list items can indent with an
	    // extra space.  That is, when the bullet did not receive an extra final
	    // space.
	    if (Number($2) < 10 && bullet.length % 2 === 1) {
	      $2 = space + $2;
	    }

	    max = $1 + repeat(space, $2.length) + $3;

	    return max + rest
	  }
	}
	return list_1$1;
}

var headingSetext;
var hasRequiredHeadingSetext;

function requireHeadingSetext () {
	if (hasRequiredHeadingSetext) return headingSetext;
	hasRequiredHeadingSetext = 1;

	headingSetext = setextHeading;

	var lineFeed = '\n';
	var tab = '\t';
	var space = ' ';
	var equalsTo = '=';
	var dash = '-';

	var maxIndent = 3;

	var equalsToDepth = 1;
	var dashDepth = 2;

	function setextHeading(eat, value, silent) {
	  var self = this;
	  var now = eat.now();
	  var length = value.length;
	  var index = -1;
	  var subvalue = '';
	  var content;
	  var queue;
	  var character;
	  var marker;
	  var depth;

	  // Eat initial indentation.
	  while (++index < length) {
	    character = value.charAt(index);

	    if (character !== space || index >= maxIndent) {
	      index--;
	      break
	    }

	    subvalue += character;
	  }

	  // Eat content.
	  content = '';
	  queue = '';

	  while (++index < length) {
	    character = value.charAt(index);

	    if (character === lineFeed) {
	      index--;
	      break
	    }

	    if (character === space || character === tab) {
	      queue += character;
	    } else {
	      content += queue + character;
	      queue = '';
	    }
	  }

	  now.column += subvalue.length;
	  now.offset += subvalue.length;
	  subvalue += content + queue;

	  // Ensure the content is followed by a newline and a valid marker.
	  character = value.charAt(++index);
	  marker = value.charAt(++index);

	  if (character !== lineFeed || (marker !== equalsTo && marker !== dash)) {
	    return
	  }

	  subvalue += character;

	  // Eat Setext-line.
	  queue = marker;
	  depth = marker === equalsTo ? equalsToDepth : dashDepth;

	  while (++index < length) {
	    character = value.charAt(index);

	    if (character !== marker) {
	      if (character !== lineFeed) {
	        return
	      }

	      index--;
	      break
	    }

	    queue += character;
	  }

	  if (silent) {
	    return true
	  }

	  return eat(subvalue + queue)({
	    type: 'heading',
	    depth: depth,
	    children: self.tokenizeInline(content, now)
	  })
	}
	return headingSetext;
}

var html$6 = {};

var hasRequiredHtml;

function requireHtml () {
	if (hasRequiredHtml) return html$6;
	hasRequiredHtml = 1;

	var attributeName = '[a-zA-Z_:][a-zA-Z0-9:._-]*';
	var unquoted = '[^"\'=<>`\\u0000-\\u0020]+';
	var singleQuoted = "'[^']*'";
	var doubleQuoted = '"[^"]*"';
	var attributeValue =
	  '(?:' + unquoted + '|' + singleQuoted + '|' + doubleQuoted + ')';
	var attribute =
	  '(?:\\s+' + attributeName + '(?:\\s*=\\s*' + attributeValue + ')?)';
	var openTag = '<[A-Za-z][A-Za-z0-9\\-]*' + attribute + '*\\s*\\/?>';
	var closeTag = '<\\/[A-Za-z][A-Za-z0-9\\-]*\\s*>';
	var comment = '<!---->|<!--(?:-?[^>-])(?:-?[^-])*-->';
	var processing = '<[?].*?[?]>';
	var declaration = '<![A-Za-z]+\\s+[^>]*>';
	var cdata = '<!\\[CDATA\\[[\\s\\S]*?\\]\\]>';

	html$6.openCloseTag = new RegExp('^(?:' + openTag + '|' + closeTag + ')');

	html$6.tag = new RegExp(
	  '^(?:' +
	    openTag +
	    '|' +
	    closeTag +
	    '|' +
	    comment +
	    '|' +
	    processing +
	    '|' +
	    declaration +
	    '|' +
	    cdata +
	    ')'
	);
	return html$6;
}

var htmlBlock;
var hasRequiredHtmlBlock;

function requireHtmlBlock () {
	if (hasRequiredHtmlBlock) return htmlBlock;
	hasRequiredHtmlBlock = 1;

	var openCloseTag = requireHtml().openCloseTag;

	htmlBlock = blockHtml;

	var tab = '\t';
	var space = ' ';
	var lineFeed = '\n';
	var lessThan = '<';

	var rawOpenExpression = /^<(script|pre|style)(?=(\s|>|$))/i;
	var rawCloseExpression = /<\/(script|pre|style)>/i;
	var commentOpenExpression = /^<!--/;
	var commentCloseExpression = /-->/;
	var instructionOpenExpression = /^<\?/;
	var instructionCloseExpression = /\?>/;
	var directiveOpenExpression = /^<![A-Za-z]/;
	var directiveCloseExpression = />/;
	var cdataOpenExpression = /^<!\[CDATA\[/;
	var cdataCloseExpression = /]]>/;
	var elementCloseExpression = /^$/;
	var otherElementOpenExpression = new RegExp(openCloseTag.source + '\\s*$');

	function blockHtml(eat, value, silent) {
	  var self = this;
	  var blocks = self.options.blocks.join('|');
	  var elementOpenExpression = new RegExp(
	    '^</?(' + blocks + ')(?=(\\s|/?>|$))',
	    'i'
	  );
	  var length = value.length;
	  var index = 0;
	  var next;
	  var line;
	  var offset;
	  var character;
	  var count;
	  var sequence;
	  var subvalue;

	  var sequences = [
	    [rawOpenExpression, rawCloseExpression, true],
	    [commentOpenExpression, commentCloseExpression, true],
	    [instructionOpenExpression, instructionCloseExpression, true],
	    [directiveOpenExpression, directiveCloseExpression, true],
	    [cdataOpenExpression, cdataCloseExpression, true],
	    [elementOpenExpression, elementCloseExpression, true],
	    [otherElementOpenExpression, elementCloseExpression, false]
	  ];

	  // Eat initial spacing.
	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== tab && character !== space) {
	      break
	    }

	    index++;
	  }

	  if (value.charAt(index) !== lessThan) {
	    return
	  }

	  next = value.indexOf(lineFeed, index + 1);
	  next = next === -1 ? length : next;
	  line = value.slice(index, next);
	  offset = -1;
	  count = sequences.length;

	  while (++offset < count) {
	    if (sequences[offset][0].test(line)) {
	      sequence = sequences[offset];
	      break
	    }
	  }

	  if (!sequence) {
	    return
	  }

	  if (silent) {
	    return sequence[2]
	  }

	  index = next;

	  if (!sequence[1].test(line)) {
	    while (index < length) {
	      next = value.indexOf(lineFeed, index + 1);
	      next = next === -1 ? length : next;
	      line = value.slice(index + 1, next);

	      if (sequence[1].test(line)) {
	        if (line) {
	          index = next;
	        }

	        break
	      }

	      index = next;
	    }
	  }

	  subvalue = value.slice(0, index);

	  return eat(subvalue)({type: 'html', value: subvalue})
	}
	return htmlBlock;
}

var isWhitespaceCharacter;
var hasRequiredIsWhitespaceCharacter;

function requireIsWhitespaceCharacter () {
	if (hasRequiredIsWhitespaceCharacter) return isWhitespaceCharacter;
	hasRequiredIsWhitespaceCharacter = 1;

	isWhitespaceCharacter = whitespace;

	var fromCode = String.fromCharCode;
	var re = /\s/;

	// Check if the given character code, or the character code at the first
	// character, is a whitespace character.
	function whitespace(character) {
	  return re.test(
	    typeof character === 'number' ? fromCode(character) : character.charAt(0)
	  )
	}
	return isWhitespaceCharacter;
}

var collapseWhiteSpace;
var hasRequiredCollapseWhiteSpace;

function requireCollapseWhiteSpace () {
	if (hasRequiredCollapseWhiteSpace) return collapseWhiteSpace;
	hasRequiredCollapseWhiteSpace = 1;

	collapseWhiteSpace = collapse;

	// `collapse(' \t\nbar \nbaz\t') // ' bar baz '`
	function collapse(value) {
	  return String(value).replace(/\s+/g, ' ')
	}
	return collapseWhiteSpace;
}

var normalize_1$1;
var hasRequiredNormalize;

function requireNormalize () {
	if (hasRequiredNormalize) return normalize_1$1;
	hasRequiredNormalize = 1;

	var collapseWhiteSpace = requireCollapseWhiteSpace();

	normalize_1$1 = normalize;

	// Normalize an identifier.  Collapses multiple white space characters into a
	// single space, and removes casing.
	function normalize(value) {
	  return collapseWhiteSpace(value).toLowerCase()
	}
	return normalize_1$1;
}

var definition_1;
var hasRequiredDefinition;

function requireDefinition () {
	if (hasRequiredDefinition) return definition_1;
	hasRequiredDefinition = 1;

	var whitespace = requireIsWhitespaceCharacter();
	var normalize = requireNormalize();

	definition_1 = definition;

	var quotationMark = '"';
	var apostrophe = "'";
	var backslash = '\\';
	var lineFeed = '\n';
	var tab = '\t';
	var space = ' ';
	var leftSquareBracket = '[';
	var rightSquareBracket = ']';
	var leftParenthesis = '(';
	var rightParenthesis = ')';
	var colon = ':';
	var lessThan = '<';
	var greaterThan = '>';

	function definition(eat, value, silent) {
	  var self = this;
	  var commonmark = self.options.commonmark;
	  var index = 0;
	  var length = value.length;
	  var subvalue = '';
	  var beforeURL;
	  var beforeTitle;
	  var queue;
	  var character;
	  var test;
	  var identifier;
	  var url;
	  var title;

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== space && character !== tab) {
	      break
	    }

	    subvalue += character;
	    index++;
	  }

	  character = value.charAt(index);

	  if (character !== leftSquareBracket) {
	    return
	  }

	  index++;
	  subvalue += character;
	  queue = '';

	  while (index < length) {
	    character = value.charAt(index);

	    if (character === rightSquareBracket) {
	      break
	    } else if (character === backslash) {
	      queue += character;
	      index++;
	      character = value.charAt(index);
	    }

	    queue += character;
	    index++;
	  }

	  if (
	    !queue ||
	    value.charAt(index) !== rightSquareBracket ||
	    value.charAt(index + 1) !== colon
	  ) {
	    return
	  }

	  identifier = queue;
	  subvalue += queue + rightSquareBracket + colon;
	  index = subvalue.length;
	  queue = '';

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== tab && character !== space && character !== lineFeed) {
	      break
	    }

	    subvalue += character;
	    index++;
	  }

	  character = value.charAt(index);
	  queue = '';
	  beforeURL = subvalue;

	  if (character === lessThan) {
	    index++;

	    while (index < length) {
	      character = value.charAt(index);

	      if (!isEnclosedURLCharacter(character)) {
	        break
	      }

	      queue += character;
	      index++;
	    }

	    character = value.charAt(index);

	    if (character === isEnclosedURLCharacter.delimiter) {
	      subvalue += lessThan + queue + character;
	      index++;
	    } else {
	      if (commonmark) {
	        return
	      }

	      index -= queue.length + 1;
	      queue = '';
	    }
	  }

	  if (!queue) {
	    while (index < length) {
	      character = value.charAt(index);

	      if (!isUnclosedURLCharacter(character)) {
	        break
	      }

	      queue += character;
	      index++;
	    }

	    subvalue += queue;
	  }

	  if (!queue) {
	    return
	  }

	  url = queue;
	  queue = '';

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== tab && character !== space && character !== lineFeed) {
	      break
	    }

	    queue += character;
	    index++;
	  }

	  character = value.charAt(index);
	  test = null;

	  if (character === quotationMark) {
	    test = quotationMark;
	  } else if (character === apostrophe) {
	    test = apostrophe;
	  } else if (character === leftParenthesis) {
	    test = rightParenthesis;
	  }

	  if (!test) {
	    queue = '';
	    index = subvalue.length;
	  } else if (queue) {
	    subvalue += queue + character;
	    index = subvalue.length;
	    queue = '';

	    while (index < length) {
	      character = value.charAt(index);

	      if (character === test) {
	        break
	      }

	      if (character === lineFeed) {
	        index++;
	        character = value.charAt(index);

	        if (character === lineFeed || character === test) {
	          return
	        }

	        queue += lineFeed;
	      }

	      queue += character;
	      index++;
	    }

	    character = value.charAt(index);

	    if (character !== test) {
	      return
	    }

	    beforeTitle = subvalue;
	    subvalue += queue + character;
	    index++;
	    title = queue;
	    queue = '';
	  } else {
	    return
	  }

	  while (index < length) {
	    character = value.charAt(index);

	    if (character !== tab && character !== space) {
	      break
	    }

	    subvalue += character;
	    index++;
	  }

	  character = value.charAt(index);

	  if (!character || character === lineFeed) {
	    if (silent) {
	      return true
	    }

	    beforeURL = eat(beforeURL).test().end;
	    url = self.decode.raw(self.unescape(url), beforeURL, {nonTerminated: false});

	    if (title) {
	      beforeTitle = eat(beforeTitle).test().end;
	      title = self.decode.raw(self.unescape(title), beforeTitle);
	    }

	    return eat(subvalue)({
	      type: 'definition',
	      identifier: normalize(identifier),
	      label: identifier,
	      title: title || null,
	      url: url
	    })
	  }
	}

	// Check if `character` can be inside an enclosed URI.
	function isEnclosedURLCharacter(character) {
	  return (
	    character !== greaterThan &&
	    character !== leftSquareBracket &&
	    character !== rightSquareBracket
	  )
	}

	isEnclosedURLCharacter.delimiter = greaterThan;

	// Check if `character` can be inside an unclosed URI.
	function isUnclosedURLCharacter(character) {
	  return (
	    character !== leftSquareBracket &&
	    character !== rightSquareBracket &&
	    !whitespace(character)
	  )
	}
	return definition_1;
}

var table_1$1;
var hasRequiredTable;

function requireTable () {
	if (hasRequiredTable) return table_1$1;
	hasRequiredTable = 1;

	var whitespace = requireIsWhitespaceCharacter();

	table_1$1 = table;

	var tab = '\t';
	var lineFeed = '\n';
	var space = ' ';
	var dash = '-';
	var colon = ':';
	var backslash = '\\';
	var verticalBar = '|';

	var minColumns = 1;
	var minRows = 2;

	var left = 'left';
	var center = 'center';
	var right = 'right';

	function table(eat, value, silent) {
	  var self = this;
	  var index;
	  var alignments;
	  var alignment;
	  var subvalue;
	  var row;
	  var length;
	  var lines;
	  var queue;
	  var character;
	  var hasDash;
	  var align;
	  var cell;
	  var preamble;
	  var now;
	  var position;
	  var lineCount;
	  var line;
	  var rows;
	  var table;
	  var lineIndex;
	  var pipeIndex;
	  var first;

	  // Exit when not in gfm-mode.
	  if (!self.options.gfm) {
	    return
	  }

	  // Get the rows.
	  // Detecting tables soon is hard, so there are some checks for performance
	  // here, such as the minimum number of rows, and allowed characters in the
	  // alignment row.
	  index = 0;
	  lineCount = 0;
	  length = value.length + 1;
	  lines = [];

	  while (index < length) {
	    lineIndex = value.indexOf(lineFeed, index);
	    pipeIndex = value.indexOf(verticalBar, index + 1);

	    if (lineIndex === -1) {
	      lineIndex = value.length;
	    }

	    if (pipeIndex === -1 || pipeIndex > lineIndex) {
	      if (lineCount < minRows) {
	        return
	      }

	      break
	    }

	    lines.push(value.slice(index, lineIndex));
	    lineCount++;
	    index = lineIndex + 1;
	  }

	  // Parse the alignment row.
	  subvalue = lines.join(lineFeed);
	  alignments = lines.splice(1, 1)[0] || [];
	  index = 0;
	  length = alignments.length;
	  lineCount--;
	  alignment = false;
	  align = [];

	  while (index < length) {
	    character = alignments.charAt(index);

	    if (character === verticalBar) {
	      hasDash = null;

	      if (alignment === false) {
	        if (first === false) {
	          return
	        }
	      } else {
	        align.push(alignment);
	        alignment = false;
	      }

	      first = false;
	    } else if (character === dash) {
	      hasDash = true;
	      alignment = alignment || null;
	    } else if (character === colon) {
	      if (alignment === left) {
	        alignment = center;
	      } else if (hasDash && alignment === null) {
	        alignment = right;
	      } else {
	        alignment = left;
	      }
	    } else if (!whitespace(character)) {
	      return
	    }

	    index++;
	  }

	  if (alignment !== false) {
	    align.push(alignment);
	  }

	  // Exit when without enough columns.
	  if (align.length < minColumns) {
	    return
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  // Parse the rows.
	  position = -1;
	  rows = [];

	  table = eat(subvalue).reset({type: 'table', align: align, children: rows});

	  while (++position < lineCount) {
	    line = lines[position];
	    row = {type: 'tableRow', children: []};

	    // Eat a newline character when this is not the first row.
	    if (position) {
	      eat(lineFeed);
	    }

	    // Eat the row.
	    eat(line).reset(row, table);

	    length = line.length + 1;
	    index = 0;
	    queue = '';
	    cell = '';
	    preamble = true;

	    while (index < length) {
	      character = line.charAt(index);

	      if (character === tab || character === space) {
	        if (cell) {
	          queue += character;
	        } else {
	          eat(character);
	        }

	        index++;
	        continue
	      }

	      if (character === '' || character === verticalBar) {
	        if (preamble) {
	          eat(character);
	        } else {
	          if ((cell || character) && !preamble) {
	            subvalue = cell;

	            if (queue.length > 1) {
	              if (character) {
	                subvalue += queue.slice(0, -1);
	                queue = queue.charAt(queue.length - 1);
	              } else {
	                subvalue += queue;
	                queue = '';
	              }
	            }

	            now = eat.now();

	            eat(subvalue)(
	              {type: 'tableCell', children: self.tokenizeInline(cell, now)},
	              row
	            );
	          }

	          eat(queue + character);

	          queue = '';
	          cell = '';
	        }
	      } else {
	        if (queue) {
	          cell += queue;
	          queue = '';
	        }

	        cell += character;

	        if (character === backslash && index !== length - 2) {
	          cell += line.charAt(index + 1);
	          index++;
	        }
	      }

	      preamble = false;
	      index++;
	    }

	    // Eat the alignment row.
	    if (!position) {
	      eat(lineFeed + alignments);
	    }
	  }

	  return table
	}
	return table_1$1;
}

var paragraph_1$1;
var hasRequiredParagraph;

function requireParagraph () {
	if (hasRequiredParagraph) return paragraph_1$1;
	hasRequiredParagraph = 1;

	var trim = requireTrim();
	var trimTrailingLines = requireTrimTrailingLines();
	var interrupt = requireInterrupt();

	paragraph_1$1 = paragraph;

	var tab = '\t';
	var lineFeed = '\n';
	var space = ' ';

	var tabSize = 4;

	// Tokenise paragraph.
	function paragraph(eat, value, silent) {
	  var self = this;
	  var settings = self.options;
	  var commonmark = settings.commonmark;
	  var tokenizers = self.blockTokenizers;
	  var interruptors = self.interruptParagraph;
	  var index = value.indexOf(lineFeed);
	  var length = value.length;
	  var position;
	  var subvalue;
	  var character;
	  var size;
	  var now;

	  while (index < length) {
	    // Eat everything if there’s no following newline.
	    if (index === -1) {
	      index = length;
	      break
	    }

	    // Stop if the next character is NEWLINE.
	    if (value.charAt(index + 1) === lineFeed) {
	      break
	    }

	    // In commonmark-mode, following indented lines are part of the paragraph.
	    if (commonmark) {
	      size = 0;
	      position = index + 1;

	      while (position < length) {
	        character = value.charAt(position);

	        if (character === tab) {
	          size = tabSize;
	          break
	        } else if (character === space) {
	          size++;
	        } else {
	          break
	        }

	        position++;
	      }

	      if (size >= tabSize && character !== lineFeed) {
	        index = value.indexOf(lineFeed, index + 1);
	        continue
	      }
	    }

	    subvalue = value.slice(index + 1);

	    // Check if the following code contains a possible block.
	    if (interrupt(interruptors, tokenizers, self, [eat, subvalue, true])) {
	      break
	    }

	    position = index;
	    index = value.indexOf(lineFeed, index + 1);

	    if (index !== -1 && trim(value.slice(position, index)) === '') {
	      index = position;
	      break
	    }
	  }

	  subvalue = value.slice(0, index);

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  now = eat.now();
	  subvalue = trimTrailingLines(subvalue);

	  return eat(subvalue)({
	    type: 'paragraph',
	    children: self.tokenizeInline(subvalue, now)
	  })
	}
	return paragraph_1$1;
}

var _escape$1;
var hasRequired_escape$1;

function require_escape$1 () {
	if (hasRequired_escape$1) return _escape$1;
	hasRequired_escape$1 = 1;

	_escape$1 = locate;

	function locate(value, fromIndex) {
	  return value.indexOf('\\', fromIndex)
	}
	return _escape$1;
}

var _escape;
var hasRequired_escape;

function require_escape () {
	if (hasRequired_escape) return _escape;
	hasRequired_escape = 1;

	var locate = require_escape$1();

	_escape = escape;
	escape.locator = locate;

	var lineFeed = '\n';
	var backslash = '\\';

	function escape(eat, value, silent) {
	  var self = this;
	  var character;
	  var node;

	  if (value.charAt(0) === backslash) {
	    character = value.charAt(1);

	    if (self.escape.indexOf(character) !== -1) {
	      /* istanbul ignore if - never used (yet) */
	      if (silent) {
	        return true
	      }

	      if (character === lineFeed) {
	        node = {type: 'break'};
	      } else {
	        node = {type: 'text', value: character};
	      }

	      return eat(backslash + character)(node)
	    }
	  }
	}
	return _escape;
}

var tag;
var hasRequiredTag;

function requireTag () {
	if (hasRequiredTag) return tag;
	hasRequiredTag = 1;

	tag = locate;

	function locate(value, fromIndex) {
	  return value.indexOf('<', fromIndex)
	}
	return tag;
}

var autoLink_1;
var hasRequiredAutoLink;

function requireAutoLink () {
	if (hasRequiredAutoLink) return autoLink_1;
	hasRequiredAutoLink = 1;

	var whitespace = requireIsWhitespaceCharacter();
	var decode = parseEntities_1;
	var locate = requireTag();

	autoLink_1 = autoLink;
	autoLink.locator = locate;
	autoLink.notInLink = true;

	var lessThan = '<';
	var greaterThan = '>';
	var atSign = '@';
	var slash = '/';
	var mailto = 'mailto:';
	var mailtoLength = mailto.length;

	function autoLink(eat, value, silent) {
	  var self = this;
	  var subvalue = '';
	  var length = value.length;
	  var index = 0;
	  var queue = '';
	  var hasAtCharacter = false;
	  var link = '';
	  var character;
	  var now;
	  var content;
	  var tokenizers;
	  var exit;

	  if (value.charAt(0) !== lessThan) {
	    return
	  }

	  index++;
	  subvalue = lessThan;

	  while (index < length) {
	    character = value.charAt(index);

	    if (
	      whitespace(character) ||
	      character === greaterThan ||
	      character === atSign ||
	      (character === ':' && value.charAt(index + 1) === slash)
	    ) {
	      break
	    }

	    queue += character;
	    index++;
	  }

	  if (!queue) {
	    return
	  }

	  link += queue;
	  queue = '';

	  character = value.charAt(index);
	  link += character;
	  index++;

	  if (character === atSign) {
	    hasAtCharacter = true;
	  } else {
	    if (character !== ':' || value.charAt(index + 1) !== slash) {
	      return
	    }

	    link += slash;
	    index++;
	  }

	  while (index < length) {
	    character = value.charAt(index);

	    if (whitespace(character) || character === greaterThan) {
	      break
	    }

	    queue += character;
	    index++;
	  }

	  character = value.charAt(index);

	  if (!queue || character !== greaterThan) {
	    return
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  link += queue;
	  content = link;
	  subvalue += link + character;
	  now = eat.now();
	  now.column++;
	  now.offset++;

	  if (hasAtCharacter) {
	    if (link.slice(0, mailtoLength).toLowerCase() === mailto) {
	      content = content.slice(mailtoLength);
	      now.column += mailtoLength;
	      now.offset += mailtoLength;
	    } else {
	      link = mailto + link;
	    }
	  }

	  // Temporarily remove all tokenizers except text in autolinks.
	  tokenizers = self.inlineTokenizers;
	  self.inlineTokenizers = {text: tokenizers.text};

	  exit = self.enterLink();

	  content = self.tokenizeInline(content, now);

	  self.inlineTokenizers = tokenizers;
	  exit();

	  return eat(subvalue)({
	    type: 'link',
	    title: null,
	    url: decode(link, {nonTerminated: false}),
	    children: content
	  })
	}
	return autoLink_1;
}

var ccount_1;
var hasRequiredCcount;

function requireCcount () {
	if (hasRequiredCcount) return ccount_1;
	hasRequiredCcount = 1;

	ccount_1 = ccount;

	function ccount(value, character) {
	  var val = String(value);
	  var count = 0;
	  var index;

	  if (typeof character !== 'string' || character.length !== 1) {
	    throw new Error('Expected character')
	  }

	  index = val.indexOf(character);

	  while (index !== -1) {
	    count++;
	    index = val.indexOf(character, index + 1);
	  }

	  return count
	}
	return ccount_1;
}

var url;
var hasRequiredUrl$1;

function requireUrl$1 () {
	if (hasRequiredUrl$1) return url;
	hasRequiredUrl$1 = 1;

	url = locate;

	var values = ['www.', 'http://', 'https://'];

	function locate(value, fromIndex) {
	  var min = -1;
	  var index;
	  var length;
	  var position;

	  if (!this.options.gfm) {
	    return min
	  }

	  length = values.length;
	  index = -1;

	  while (++index < length) {
	    position = value.indexOf(values[index], fromIndex);

	    if (position !== -1 && (min === -1 || position < min)) {
	      min = position;
	    }
	  }

	  return min
	}
	return url;
}

var url_1;
var hasRequiredUrl;

function requireUrl () {
	if (hasRequiredUrl) return url_1;
	hasRequiredUrl = 1;

	var ccount = requireCcount();
	var decode = parseEntities_1;
	var decimal = requireIsDecimal();
	var alphabetical = requireIsAlphabetical();
	var whitespace = requireIsWhitespaceCharacter();
	var locate = requireUrl$1();

	url_1 = url;
	url.locator = locate;
	url.notInLink = true;

	var exclamationMark = 33; // '!'
	var ampersand = 38; // '&'
	var rightParenthesis = 41; // ')'
	var asterisk = 42; // '*'
	var comma = 44; // ','
	var dash = 45; // '-'
	var dot = 46; // '.'
	var colon = 58; // ':'
	var semicolon = 59; // ';'
	var questionMark = 63; // '?'
	var lessThan = 60; // '<'
	var underscore = 95; // '_'
	var tilde = 126; // '~'

	var leftParenthesisCharacter = '(';
	var rightParenthesisCharacter = ')';

	function url(eat, value, silent) {
	  var self = this;
	  var gfm = self.options.gfm;
	  var tokenizers = self.inlineTokenizers;
	  var length = value.length;
	  var previousDot = -1;
	  var protocolless = false;
	  var dots;
	  var lastTwoPartsStart;
	  var start;
	  var index;
	  var pathStart;
	  var path;
	  var code;
	  var end;
	  var leftCount;
	  var rightCount;
	  var content;
	  var children;
	  var url;
	  var exit;

	  if (!gfm) {
	    return
	  }

	  // `WWW.` doesn’t work.
	  if (value.slice(0, 4) === 'www.') {
	    protocolless = true;
	    index = 4;
	  } else if (value.slice(0, 7).toLowerCase() === 'http://') {
	    index = 7;
	  } else if (value.slice(0, 8).toLowerCase() === 'https://') {
	    index = 8;
	  } else {
	    return
	  }

	  // Act as if the starting boundary is a dot.
	  previousDot = index - 1;

	  // Parse a valid domain.
	  start = index;
	  dots = [];

	  while (index < length) {
	    code = value.charCodeAt(index);

	    if (code === dot) {
	      // Dots may not appear after each other.
	      if (previousDot === index - 1) {
	        break
	      }

	      dots.push(index);
	      previousDot = index;
	      index++;
	      continue
	    }

	    if (
	      decimal(code) ||
	      alphabetical(code) ||
	      code === dash ||
	      code === underscore
	    ) {
	      index++;
	      continue
	    }

	    break
	  }

	  // Ignore a final dot:
	  if (code === dot) {
	    dots.pop();
	    index--;
	  }

	  // If there are not dots, exit.
	  if (dots[0] === undefined) {
	    return
	  }

	  // If there is an underscore in the last two domain parts, exit:
	  // `www.example.c_m` and `www.ex_ample.com` are not OK, but
	  // `www.sub_domain.example.com` is.
	  lastTwoPartsStart = dots.length < 2 ? start : dots[dots.length - 2] + 1;

	  if (value.slice(lastTwoPartsStart, index).indexOf('_') !== -1) {
	    return
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  end = index;
	  pathStart = index;

	  // Parse a path.
	  while (index < length) {
	    code = value.charCodeAt(index);

	    if (whitespace(code) || code === lessThan) {
	      break
	    }

	    index++;

	    if (
	      code === exclamationMark ||
	      code === asterisk ||
	      code === comma ||
	      code === dot ||
	      code === colon ||
	      code === questionMark ||
	      code === underscore ||
	      code === tilde
	    ) ; else {
	      end = index;
	    }
	  }

	  index = end;

	  // If the path ends in a closing paren, and the count of closing parens is
	  // higher than the opening count, then remove the supefluous closing parens.
	  if (value.charCodeAt(index - 1) === rightParenthesis) {
	    path = value.slice(pathStart, index);
	    leftCount = ccount(path, leftParenthesisCharacter);
	    rightCount = ccount(path, rightParenthesisCharacter);

	    while (rightCount > leftCount) {
	      index = pathStart + path.lastIndexOf(rightParenthesisCharacter);
	      path = value.slice(pathStart, index);
	      rightCount--;
	    }
	  }

	  if (value.charCodeAt(index - 1) === semicolon) {
	    // GitHub doesn’t document this, but final semicolons aren’t paret of the
	    // URL either.
	    index--;

	    // // If the path ends in what looks like an entity, it’s not part of the path.
	    if (alphabetical(value.charCodeAt(index - 1))) {
	      end = index - 2;

	      while (alphabetical(value.charCodeAt(end))) {
	        end--;
	      }

	      if (value.charCodeAt(end) === ampersand) {
	        index = end;
	      }
	    }
	  }

	  content = value.slice(0, index);
	  url = decode(content, {nonTerminated: false});

	  if (protocolless) {
	    url = 'http://' + url;
	  }

	  exit = self.enterLink();

	  // Temporarily remove all tokenizers except text in url.
	  self.inlineTokenizers = {text: tokenizers.text};
	  children = self.tokenizeInline(content, eat.now());
	  self.inlineTokenizers = tokenizers;

	  exit();

	  return eat(content)({type: 'link', title: null, url: url, children: children})
	}
	return url_1;
}

var email;
var hasRequiredEmail$1;

function requireEmail$1 () {
	if (hasRequiredEmail$1) return email;
	hasRequiredEmail$1 = 1;

	var decimal = requireIsDecimal();
	var alphabetical = requireIsAlphabetical();

	var plusSign = 43; // '+'
	var dash = 45; // '-'
	var dot = 46; // '.'
	var underscore = 95; // '_'

	email = locate;

	// See: <https://github.github.com/gfm/#extended-email-autolink>
	function locate(value, fromIndex) {
	  var self = this;
	  var at;
	  var position;

	  if (!this.options.gfm) {
	    return -1
	  }

	  at = value.indexOf('@', fromIndex);

	  if (at === -1) {
	    return -1
	  }

	  position = at;

	  if (position === fromIndex || !isGfmAtext(value.charCodeAt(position - 1))) {
	    return locate.call(self, value, at + 1)
	  }

	  while (position > fromIndex && isGfmAtext(value.charCodeAt(position - 1))) {
	    position--;
	  }

	  return position
	}

	function isGfmAtext(code) {
	  return (
	    decimal(code) ||
	    alphabetical(code) ||
	    code === plusSign ||
	    code === dash ||
	    code === dot ||
	    code === underscore
	  )
	}
	return email;
}

var email_1;
var hasRequiredEmail;

function requireEmail () {
	if (hasRequiredEmail) return email_1;
	hasRequiredEmail = 1;

	var decode = parseEntities_1;
	var decimal = requireIsDecimal();
	var alphabetical = requireIsAlphabetical();
	var locate = requireEmail$1();

	email_1 = email;
	email.locator = locate;
	email.notInLink = true;

	var plusSign = 43; // '+'
	var dash = 45; // '-'
	var dot = 46; // '.'
	var atSign = 64; // '@'
	var underscore = 95; // '_'

	function email(eat, value, silent) {
	  var self = this;
	  var gfm = self.options.gfm;
	  var tokenizers = self.inlineTokenizers;
	  var index = 0;
	  var length = value.length;
	  var firstDot = -1;
	  var code;
	  var content;
	  var children;
	  var exit;

	  if (!gfm) {
	    return
	  }

	  code = value.charCodeAt(index);

	  while (
	    decimal(code) ||
	    alphabetical(code) ||
	    code === plusSign ||
	    code === dash ||
	    code === dot ||
	    code === underscore
	  ) {
	    code = value.charCodeAt(++index);
	  }

	  if (index === 0) {
	    return
	  }

	  if (code !== atSign) {
	    return
	  }

	  index++;

	  while (index < length) {
	    code = value.charCodeAt(index);

	    if (
	      decimal(code) ||
	      alphabetical(code) ||
	      code === dash ||
	      code === dot ||
	      code === underscore
	    ) {
	      index++;

	      if (firstDot === -1 && code === dot) {
	        firstDot = index;
	      }

	      continue
	    }

	    break
	  }

	  if (
	    firstDot === -1 ||
	    firstDot === index ||
	    code === dash ||
	    code === underscore
	  ) {
	    return
	  }

	  if (code === dot) {
	    index--;
	  }

	  content = value.slice(0, index);

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  exit = self.enterLink();

	  // Temporarily remove all tokenizers except text in url.
	  self.inlineTokenizers = {text: tokenizers.text};
	  children = self.tokenizeInline(content, eat.now());
	  self.inlineTokenizers = tokenizers;

	  exit();

	  return eat(content)({
	    type: 'link',
	    title: null,
	    url: 'mailto:' + decode(content, {nonTerminated: false}),
	    children: children
	  })
	}
	return email_1;
}

var htmlInline;
var hasRequiredHtmlInline;

function requireHtmlInline () {
	if (hasRequiredHtmlInline) return htmlInline;
	hasRequiredHtmlInline = 1;

	var alphabetical = requireIsAlphabetical();
	var locate = requireTag();
	var tag = requireHtml().tag;

	htmlInline = inlineHTML;
	inlineHTML.locator = locate;

	var lessThan = '<';
	var questionMark = '?';
	var exclamationMark = '!';
	var slash = '/';

	var htmlLinkOpenExpression = /^<a /i;
	var htmlLinkCloseExpression = /^<\/a>/i;

	function inlineHTML(eat, value, silent) {
	  var self = this;
	  var length = value.length;
	  var character;
	  var subvalue;

	  if (value.charAt(0) !== lessThan || length < 3) {
	    return
	  }

	  character = value.charAt(1);

	  if (
	    !alphabetical(character) &&
	    character !== questionMark &&
	    character !== exclamationMark &&
	    character !== slash
	  ) {
	    return
	  }

	  subvalue = value.match(tag);

	  if (!subvalue) {
	    return
	  }

	  /* istanbul ignore if - not used yet. */
	  if (silent) {
	    return true
	  }

	  subvalue = subvalue[0];

	  if (!self.inLink && htmlLinkOpenExpression.test(subvalue)) {
	    self.inLink = true;
	  } else if (self.inLink && htmlLinkCloseExpression.test(subvalue)) {
	    self.inLink = false;
	  }

	  return eat(subvalue)({type: 'html', value: subvalue})
	}
	return htmlInline;
}

var link$1;
var hasRequiredLink$1;

function requireLink$1 () {
	if (hasRequiredLink$1) return link$1;
	hasRequiredLink$1 = 1;

	link$1 = locate;

	function locate(value, fromIndex) {
	  var link = value.indexOf('[', fromIndex);
	  var image = value.indexOf('![', fromIndex);

	  if (image === -1) {
	    return link
	  }

	  // Link can never be `-1` if an image is found, so we don’t need to check
	  // for that :)
	  return link < image ? link : image
	}
	return link$1;
}

var link_1$1;
var hasRequiredLink;

function requireLink () {
	if (hasRequiredLink) return link_1$1;
	hasRequiredLink = 1;

	var whitespace = requireIsWhitespaceCharacter();
	var locate = requireLink$1();

	link_1$1 = link;
	link.locator = locate;

	var lineFeed = '\n';
	var exclamationMark = '!';
	var quotationMark = '"';
	var apostrophe = "'";
	var leftParenthesis = '(';
	var rightParenthesis = ')';
	var lessThan = '<';
	var greaterThan = '>';
	var leftSquareBracket = '[';
	var backslash = '\\';
	var rightSquareBracket = ']';
	var graveAccent = '`';

	function link(eat, value, silent) {
	  var self = this;
	  var subvalue = '';
	  var index = 0;
	  var character = value.charAt(0);
	  var pedantic = self.options.pedantic;
	  var commonmark = self.options.commonmark;
	  var gfm = self.options.gfm;
	  var closed;
	  var count;
	  var opening;
	  var beforeURL;
	  var beforeTitle;
	  var subqueue;
	  var hasMarker;
	  var isImage;
	  var content;
	  var marker;
	  var length;
	  var title;
	  var depth;
	  var queue;
	  var url;
	  var now;
	  var exit;
	  var node;

	  // Detect whether this is an image.
	  if (character === exclamationMark) {
	    isImage = true;
	    subvalue = character;
	    character = value.charAt(++index);
	  }

	  // Eat the opening.
	  if (character !== leftSquareBracket) {
	    return
	  }

	  // Exit when this is a link and we’re already inside a link.
	  if (!isImage && self.inLink) {
	    return
	  }

	  subvalue += character;
	  queue = '';
	  index++;

	  // Eat the content.
	  length = value.length;
	  now = eat.now();
	  depth = 0;

	  now.column += index;
	  now.offset += index;

	  while (index < length) {
	    character = value.charAt(index);
	    subqueue = character;

	    if (character === graveAccent) {
	      // Inline-code in link content.
	      count = 1;

	      while (value.charAt(index + 1) === graveAccent) {
	        subqueue += character;
	        index++;
	        count++;
	      }

	      if (!opening) {
	        opening = count;
	      } else if (count >= opening) {
	        opening = 0;
	      }
	    } else if (character === backslash) {
	      // Allow brackets to be escaped.
	      index++;
	      subqueue += value.charAt(index);
	    } else if ((!opening || gfm) && character === leftSquareBracket) {
	      // In GFM mode, brackets in code still count.  In all other modes,
	      // they don’t.
	      depth++;
	    } else if ((!opening || gfm) && character === rightSquareBracket) {
	      if (depth) {
	        depth--;
	      } else {
	        if (value.charAt(index + 1) !== leftParenthesis) {
	          return
	        }

	        subqueue += leftParenthesis;
	        closed = true;
	        index++;

	        break
	      }
	    }

	    queue += subqueue;
	    subqueue = '';
	    index++;
	  }

	  // Eat the content closing.
	  if (!closed) {
	    return
	  }

	  content = queue;
	  subvalue += queue + subqueue;
	  index++;

	  // Eat white-space.
	  while (index < length) {
	    character = value.charAt(index);

	    if (!whitespace(character)) {
	      break
	    }

	    subvalue += character;
	    index++;
	  }

	  // Eat the URL.
	  character = value.charAt(index);
	  queue = '';
	  beforeURL = subvalue;

	  if (character === lessThan) {
	    index++;
	    beforeURL += lessThan;

	    while (index < length) {
	      character = value.charAt(index);

	      if (character === greaterThan) {
	        break
	      }

	      if (commonmark && character === lineFeed) {
	        return
	      }

	      queue += character;
	      index++;
	    }

	    if (value.charAt(index) !== greaterThan) {
	      return
	    }

	    subvalue += lessThan + queue + greaterThan;
	    url = queue;
	    index++;
	  } else {
	    character = null;
	    subqueue = '';

	    while (index < length) {
	      character = value.charAt(index);

	      if (
	        subqueue &&
	        (character === quotationMark ||
	          character === apostrophe ||
	          (commonmark && character === leftParenthesis))
	      ) {
	        break
	      }

	      if (whitespace(character)) {
	        if (!pedantic) {
	          break
	        }

	        subqueue += character;
	      } else {
	        if (character === leftParenthesis) {
	          depth++;
	        } else if (character === rightParenthesis) {
	          if (depth === 0) {
	            break
	          }

	          depth--;
	        }

	        queue += subqueue;
	        subqueue = '';

	        if (character === backslash) {
	          queue += backslash;
	          character = value.charAt(++index);
	        }

	        queue += character;
	      }

	      index++;
	    }

	    subvalue += queue;
	    url = queue;
	    index = subvalue.length;
	  }

	  // Eat white-space.
	  queue = '';

	  while (index < length) {
	    character = value.charAt(index);

	    if (!whitespace(character)) {
	      break
	    }

	    queue += character;
	    index++;
	  }

	  character = value.charAt(index);
	  subvalue += queue;

	  // Eat the title.
	  if (
	    queue &&
	    (character === quotationMark ||
	      character === apostrophe ||
	      (commonmark && character === leftParenthesis))
	  ) {
	    index++;
	    subvalue += character;
	    queue = '';
	    marker = character === leftParenthesis ? rightParenthesis : character;
	    beforeTitle = subvalue;

	    // In commonmark-mode, things are pretty easy: the marker cannot occur
	    // inside the title.  Non-commonmark does, however, support nested
	    // delimiters.
	    if (commonmark) {
	      while (index < length) {
	        character = value.charAt(index);

	        if (character === marker) {
	          break
	        }

	        if (character === backslash) {
	          queue += backslash;
	          character = value.charAt(++index);
	        }

	        index++;
	        queue += character;
	      }

	      character = value.charAt(index);

	      if (character !== marker) {
	        return
	      }

	      title = queue;
	      subvalue += queue + character;
	      index++;

	      while (index < length) {
	        character = value.charAt(index);

	        if (!whitespace(character)) {
	          break
	        }

	        subvalue += character;
	        index++;
	      }
	    } else {
	      subqueue = '';

	      while (index < length) {
	        character = value.charAt(index);

	        if (character === marker) {
	          if (hasMarker) {
	            queue += marker + subqueue;
	            subqueue = '';
	          }

	          hasMarker = true;
	        } else if (!hasMarker) {
	          queue += character;
	        } else if (character === rightParenthesis) {
	          subvalue += queue + marker + subqueue;
	          title = queue;
	          break
	        } else if (whitespace(character)) {
	          subqueue += character;
	        } else {
	          queue += marker + subqueue + character;
	          subqueue = '';
	          hasMarker = false;
	        }

	        index++;
	      }
	    }
	  }

	  if (value.charAt(index) !== rightParenthesis) {
	    return
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  subvalue += rightParenthesis;

	  url = self.decode.raw(self.unescape(url), eat(beforeURL).test().end, {
	    nonTerminated: false
	  });

	  if (title) {
	    beforeTitle = eat(beforeTitle).test().end;
	    title = self.decode.raw(self.unescape(title), beforeTitle);
	  }

	  node = {
	    type: isImage ? 'image' : 'link',
	    title: title || null,
	    url: url
	  };

	  if (isImage) {
	    node.alt = self.decode.raw(self.unescape(content), now) || null;
	  } else {
	    exit = self.enterLink();
	    node.children = self.tokenizeInline(content, now);
	    exit();
	  }

	  return eat(subvalue)(node)
	}
	return link_1$1;
}

var reference_1;
var hasRequiredReference;

function requireReference () {
	if (hasRequiredReference) return reference_1;
	hasRequiredReference = 1;

	var whitespace = requireIsWhitespaceCharacter();
	var locate = requireLink$1();
	var normalize = requireNormalize();

	reference_1 = reference;
	reference.locator = locate;

	var link = 'link';
	var image = 'image';
	var shortcut = 'shortcut';
	var collapsed = 'collapsed';
	var full = 'full';
	var exclamationMark = '!';
	var leftSquareBracket = '[';
	var backslash = '\\';
	var rightSquareBracket = ']';

	function reference(eat, value, silent) {
	  var self = this;
	  var commonmark = self.options.commonmark;
	  var character = value.charAt(0);
	  var index = 0;
	  var length = value.length;
	  var subvalue = '';
	  var intro = '';
	  var type = link;
	  var referenceType = shortcut;
	  var content;
	  var identifier;
	  var now;
	  var node;
	  var exit;
	  var queue;
	  var bracketed;
	  var depth;

	  // Check whether we’re eating an image.
	  if (character === exclamationMark) {
	    type = image;
	    intro = character;
	    character = value.charAt(++index);
	  }

	  if (character !== leftSquareBracket) {
	    return
	  }

	  index++;
	  intro += character;
	  queue = '';

	  // Eat the text.
	  depth = 0;

	  while (index < length) {
	    character = value.charAt(index);

	    if (character === leftSquareBracket) {
	      bracketed = true;
	      depth++;
	    } else if (character === rightSquareBracket) {
	      if (!depth) {
	        break
	      }

	      depth--;
	    }

	    if (character === backslash) {
	      queue += backslash;
	      character = value.charAt(++index);
	    }

	    queue += character;
	    index++;
	  }

	  subvalue = queue;
	  content = queue;
	  character = value.charAt(index);

	  if (character !== rightSquareBracket) {
	    return
	  }

	  index++;
	  subvalue += character;
	  queue = '';

	  if (!commonmark) {
	    // The original markdown syntax definition explicitly allows for whitespace
	    // between the link text and link label; commonmark departs from this, in
	    // part to improve support for shortcut reference links
	    while (index < length) {
	      character = value.charAt(index);

	      if (!whitespace(character)) {
	        break
	      }

	      queue += character;
	      index++;
	    }
	  }

	  character = value.charAt(index);

	  if (character === leftSquareBracket) {
	    identifier = '';
	    queue += character;
	    index++;

	    while (index < length) {
	      character = value.charAt(index);

	      if (character === leftSquareBracket || character === rightSquareBracket) {
	        break
	      }

	      if (character === backslash) {
	        identifier += backslash;
	        character = value.charAt(++index);
	      }

	      identifier += character;
	      index++;
	    }

	    character = value.charAt(index);

	    if (character === rightSquareBracket) {
	      referenceType = identifier ? full : collapsed;
	      queue += identifier + character;
	      index++;
	    } else {
	      identifier = '';
	    }

	    subvalue += queue;
	    queue = '';
	  } else {
	    if (!content) {
	      return
	    }

	    identifier = content;
	  }

	  // Brackets cannot be inside the identifier.
	  if (referenceType !== full && bracketed) {
	    return
	  }

	  subvalue = intro + subvalue;

	  if (type === link && self.inLink) {
	    return null
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  now = eat.now();
	  now.column += intro.length;
	  now.offset += intro.length;
	  identifier = referenceType === full ? identifier : content;

	  node = {
	    type: type + 'Reference',
	    identifier: normalize(identifier),
	    label: identifier,
	    referenceType: referenceType
	  };

	  if (type === link) {
	    exit = self.enterLink();
	    node.children = self.tokenizeInline(content, now);
	    exit();
	  } else {
	    node.alt = self.decode.raw(self.unescape(content), now) || null;
	  }

	  return eat(subvalue)(node)
	}
	return reference_1;
}

var strong$1;
var hasRequiredStrong$1;

function requireStrong$1 () {
	if (hasRequiredStrong$1) return strong$1;
	hasRequiredStrong$1 = 1;

	strong$1 = locate;

	function locate(value, fromIndex) {
	  var asterisk = value.indexOf('**', fromIndex);
	  var underscore = value.indexOf('__', fromIndex);

	  if (underscore === -1) {
	    return asterisk
	  }

	  if (asterisk === -1) {
	    return underscore
	  }

	  return underscore < asterisk ? underscore : asterisk
	}
	return strong$1;
}

var strong_1$1;
var hasRequiredStrong;

function requireStrong () {
	if (hasRequiredStrong) return strong_1$1;
	hasRequiredStrong = 1;

	var trim = requireTrim();
	var whitespace = requireIsWhitespaceCharacter();
	var locate = requireStrong$1();

	strong_1$1 = strong;
	strong.locator = locate;

	var backslash = '\\';
	var asterisk = '*';
	var underscore = '_';

	function strong(eat, value, silent) {
	  var self = this;
	  var index = 0;
	  var character = value.charAt(index);
	  var now;
	  var pedantic;
	  var marker;
	  var queue;
	  var subvalue;
	  var length;
	  var previous;

	  if (
	    (character !== asterisk && character !== underscore) ||
	    value.charAt(++index) !== character
	  ) {
	    return
	  }

	  pedantic = self.options.pedantic;
	  marker = character;
	  subvalue = marker + marker;
	  length = value.length;
	  index++;
	  queue = '';
	  character = '';

	  if (pedantic && whitespace(value.charAt(index))) {
	    return
	  }

	  while (index < length) {
	    previous = character;
	    character = value.charAt(index);

	    if (
	      character === marker &&
	      value.charAt(index + 1) === marker &&
	      (!pedantic || !whitespace(previous))
	    ) {
	      character = value.charAt(index + 2);

	      if (character !== marker) {
	        if (!trim(queue)) {
	          return
	        }

	        /* istanbul ignore if - never used (yet) */
	        if (silent) {
	          return true
	        }

	        now = eat.now();
	        now.column += 2;
	        now.offset += 2;

	        return eat(subvalue + queue + subvalue)({
	          type: 'strong',
	          children: self.tokenizeInline(queue, now)
	        })
	      }
	    }

	    if (!pedantic && character === backslash) {
	      queue += character;
	      character = value.charAt(++index);
	    }

	    queue += character;
	    index++;
	  }
	}
	return strong_1$1;
}

var isWordCharacter;
var hasRequiredIsWordCharacter;

function requireIsWordCharacter () {
	if (hasRequiredIsWordCharacter) return isWordCharacter;
	hasRequiredIsWordCharacter = 1;

	isWordCharacter = wordCharacter;

	var fromCode = String.fromCharCode;
	var re = /\w/;

	// Check if the given character code, or the character code at the first
	// character, is a word character.
	function wordCharacter(character) {
	  return re.test(
	    typeof character === 'number' ? fromCode(character) : character.charAt(0)
	  )
	}
	return isWordCharacter;
}

var emphasis$1;
var hasRequiredEmphasis$1;

function requireEmphasis$1 () {
	if (hasRequiredEmphasis$1) return emphasis$1;
	hasRequiredEmphasis$1 = 1;

	emphasis$1 = locate;

	function locate(value, fromIndex) {
	  var asterisk = value.indexOf('*', fromIndex);
	  var underscore = value.indexOf('_', fromIndex);

	  if (underscore === -1) {
	    return asterisk
	  }

	  if (asterisk === -1) {
	    return underscore
	  }

	  return underscore < asterisk ? underscore : asterisk
	}
	return emphasis$1;
}

var emphasis_1$1;
var hasRequiredEmphasis;

function requireEmphasis () {
	if (hasRequiredEmphasis) return emphasis_1$1;
	hasRequiredEmphasis = 1;

	var trim = requireTrim();
	var word = requireIsWordCharacter();
	var whitespace = requireIsWhitespaceCharacter();
	var locate = requireEmphasis$1();

	emphasis_1$1 = emphasis;
	emphasis.locator = locate;

	var asterisk = '*';
	var underscore = '_';
	var backslash = '\\';

	function emphasis(eat, value, silent) {
	  var self = this;
	  var index = 0;
	  var character = value.charAt(index);
	  var now;
	  var pedantic;
	  var marker;
	  var queue;
	  var subvalue;
	  var length;
	  var previous;

	  if (character !== asterisk && character !== underscore) {
	    return
	  }

	  pedantic = self.options.pedantic;
	  subvalue = character;
	  marker = character;
	  length = value.length;
	  index++;
	  queue = '';
	  character = '';

	  if (pedantic && whitespace(value.charAt(index))) {
	    return
	  }

	  while (index < length) {
	    previous = character;
	    character = value.charAt(index);

	    if (character === marker && (!pedantic || !whitespace(previous))) {
	      character = value.charAt(++index);

	      if (character !== marker) {
	        if (!trim(queue) || previous === marker) {
	          return
	        }

	        if (!pedantic && marker === underscore && word(character)) {
	          queue += marker;
	          continue
	        }

	        /* istanbul ignore if - never used (yet) */
	        if (silent) {
	          return true
	        }

	        now = eat.now();
	        now.column++;
	        now.offset++;

	        return eat(subvalue + queue + marker)({
	          type: 'emphasis',
	          children: self.tokenizeInline(queue, now)
	        })
	      }

	      queue += marker;
	    }

	    if (!pedantic && character === backslash) {
	      queue += character;
	      character = value.charAt(++index);
	    }

	    queue += character;
	    index++;
	  }
	}
	return emphasis_1$1;
}

var _delete$2;
var hasRequired_delete$1;

function require_delete$1 () {
	if (hasRequired_delete$1) return _delete$2;
	hasRequired_delete$1 = 1;

	_delete$2 = locate;

	function locate(value, fromIndex) {
	  return value.indexOf('~~', fromIndex)
	}
	return _delete$2;
}

var _delete$1;
var hasRequired_delete;

function require_delete () {
	if (hasRequired_delete) return _delete$1;
	hasRequired_delete = 1;

	var whitespace = requireIsWhitespaceCharacter();
	var locate = require_delete$1();

	_delete$1 = strikethrough;
	strikethrough.locator = locate;

	var tilde = '~';
	var fence = '~~';

	function strikethrough(eat, value, silent) {
	  var self = this;
	  var character = '';
	  var previous = '';
	  var preceding = '';
	  var subvalue = '';
	  var index;
	  var length;
	  var now;

	  if (
	    !self.options.gfm ||
	    value.charAt(0) !== tilde ||
	    value.charAt(1) !== tilde ||
	    whitespace(value.charAt(2))
	  ) {
	    return
	  }

	  index = 1;
	  length = value.length;
	  now = eat.now();
	  now.column += 2;
	  now.offset += 2;

	  while (++index < length) {
	    character = value.charAt(index);

	    if (
	      character === tilde &&
	      previous === tilde &&
	      (!preceding || !whitespace(preceding))
	    ) {
	      /* istanbul ignore if - never used (yet) */
	      if (silent) {
	        return true
	      }

	      return eat(fence + subvalue + fence)({
	        type: 'delete',
	        children: self.tokenizeInline(subvalue, now)
	      })
	    }

	    subvalue += previous;
	    preceding = previous;
	    previous = character;
	  }
	}
	return _delete$1;
}

var codeInline$1;
var hasRequiredCodeInline$1;

function requireCodeInline$1 () {
	if (hasRequiredCodeInline$1) return codeInline$1;
	hasRequiredCodeInline$1 = 1;

	codeInline$1 = locate;

	function locate(value, fromIndex) {
	  return value.indexOf('`', fromIndex)
	}
	return codeInline$1;
}

var codeInline;
var hasRequiredCodeInline;

function requireCodeInline () {
	if (hasRequiredCodeInline) return codeInline;
	hasRequiredCodeInline = 1;

	var locate = requireCodeInline$1();

	codeInline = inlineCode;
	inlineCode.locator = locate;

	var lineFeed = 10; //  '\n'
	var space = 32; // ' '
	var graveAccent = 96; //  '`'

	function inlineCode(eat, value, silent) {
	  var length = value.length;
	  var index = 0;
	  var openingFenceEnd;
	  var closingFenceStart;
	  var closingFenceEnd;
	  var code;
	  var next;
	  var found;

	  while (index < length) {
	    if (value.charCodeAt(index) !== graveAccent) {
	      break
	    }

	    index++;
	  }

	  if (index === 0 || index === length) {
	    return
	  }

	  openingFenceEnd = index;
	  next = value.charCodeAt(index);

	  while (index < length) {
	    code = next;
	    next = value.charCodeAt(index + 1);

	    if (code === graveAccent) {
	      if (closingFenceStart === undefined) {
	        closingFenceStart = index;
	      }

	      closingFenceEnd = index + 1;

	      if (
	        next !== graveAccent &&
	        closingFenceEnd - closingFenceStart === openingFenceEnd
	      ) {
	        found = true;
	        break
	      }
	    } else if (closingFenceStart !== undefined) {
	      closingFenceStart = undefined;
	      closingFenceEnd = undefined;
	    }

	    index++;
	  }

	  if (!found) {
	    return
	  }

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  // Remove the initial and final space (or line feed), iff they exist and there
	  // are non-space characters in the content.
	  index = openingFenceEnd;
	  length = closingFenceStart;
	  code = value.charCodeAt(index);
	  next = value.charCodeAt(length - 1);
	  found = false;

	  if (
	    length - index > 2 &&
	    (code === space || code === lineFeed) &&
	    (next === space || next === lineFeed)
	  ) {
	    index++;
	    length--;

	    while (index < length) {
	      code = value.charCodeAt(index);

	      if (code !== space && code !== lineFeed) {
	        found = true;
	        break
	      }

	      index++;
	    }

	    if (found === true) {
	      openingFenceEnd++;
	      closingFenceStart--;
	    }
	  }

	  return eat(value.slice(0, closingFenceEnd))({
	    type: 'inlineCode',
	    value: value.slice(openingFenceEnd, closingFenceStart)
	  })
	}
	return codeInline;
}

var _break$2;
var hasRequired_break$1;

function require_break$1 () {
	if (hasRequired_break$1) return _break$2;
	hasRequired_break$1 = 1;

	_break$2 = locate;

	function locate(value, fromIndex) {
	  var index = value.indexOf('\n', fromIndex);

	  while (index > fromIndex) {
	    if (value.charAt(index - 1) !== ' ') {
	      break
	    }

	    index--;
	  }

	  return index
	}
	return _break$2;
}

var _break$1;
var hasRequired_break;

function require_break () {
	if (hasRequired_break) return _break$1;
	hasRequired_break = 1;

	var locate = require_break$1();

	_break$1 = hardBreak;
	hardBreak.locator = locate;

	var space = ' ';
	var lineFeed = '\n';
	var minBreakLength = 2;

	function hardBreak(eat, value, silent) {
	  var length = value.length;
	  var index = -1;
	  var queue = '';
	  var character;

	  while (++index < length) {
	    character = value.charAt(index);

	    if (character === lineFeed) {
	      if (index < minBreakLength) {
	        return
	      }

	      /* istanbul ignore if - never used (yet) */
	      if (silent) {
	        return true
	      }

	      queue += character;

	      return eat(queue)({type: 'break'})
	    }

	    if (character !== space) {
	      return
	    }

	    queue += character;
	  }
	}
	return _break$1;
}

var text_1$2;
var hasRequiredText;

function requireText () {
	if (hasRequiredText) return text_1$2;
	hasRequiredText = 1;

	text_1$2 = text;

	function text(eat, value, silent) {
	  var self = this;
	  var methods;
	  var tokenizers;
	  var index;
	  var length;
	  var subvalue;
	  var position;
	  var tokenizer;
	  var name;
	  var min;
	  var now;

	  /* istanbul ignore if - never used (yet) */
	  if (silent) {
	    return true
	  }

	  methods = self.inlineMethods;
	  length = methods.length;
	  tokenizers = self.inlineTokenizers;
	  index = -1;
	  min = value.length;

	  while (++index < length) {
	    name = methods[index];

	    if (name === 'text' || !tokenizers[name]) {
	      continue
	    }

	    tokenizer = tokenizers[name].locator;

	    if (!tokenizer) {
	      eat.file.fail('Missing locator: `' + name + '`');
	    }

	    position = tokenizer.call(self, value, 1);

	    if (position !== -1 && position < min) {
	      min = position;
	    }
	  }

	  subvalue = value.slice(0, min);
	  now = eat.now();

	  self.decode(subvalue, now, handler);

	  function handler(content, position, source) {
	    eat(source || content)({type: 'text', value: content});
	  }
	}
	return text_1$2;
}

var xtend$2 = immutable;
var toggle = stateToggle;
var vfileLocation = vfileLocation$1;
var unescape = _unescape;
var decode = decode$1;
var tokenizer$2 = tokenizer$3;

var parser$1 = Parser$1;

function Parser$1(doc, file) {
  this.file = file;
  this.offset = {};
  this.options = xtend$2(this.options);
  this.setOptions({});

  this.inList = false;
  this.inBlock = false;
  this.inLink = false;
  this.atStart = true;

  this.toOffset = vfileLocation(file).toOffset;
  this.unescape = unescape(this, 'escape');
  this.decode = decode(this);
}

var proto$3 = Parser$1.prototype;

// Expose core.
proto$3.setOptions = requireSetOptions();
proto$3.parse = requireParse();

// Expose `defaults`.
proto$3.options = requireDefaults();

// Enter and exit helpers.
proto$3.exitStart = toggle('atStart', true);
proto$3.enterList = toggle('inList', false);
proto$3.enterLink = toggle('inLink', false);
proto$3.enterBlock = toggle('inBlock', false);

// Nodes that can interupt a paragraph:
//
// ```markdown
// A paragraph, followed by a thematic break.
// ___
// ```
//
// In the above example, the thematic break “interupts” the paragraph.
proto$3.interruptParagraph = [
  ['thematicBreak'],
  ['list'],
  ['atxHeading'],
  ['fencedCode'],
  ['blockquote'],
  ['html'],
  ['setextHeading', {commonmark: false}],
  ['definition', {commonmark: false}]
];

// Nodes that can interupt a list:
//
// ```markdown
// - One
// ___
// ```
//
// In the above example, the thematic break “interupts” the list.
proto$3.interruptList = [
  ['atxHeading', {pedantic: false}],
  ['fencedCode', {pedantic: false}],
  ['thematicBreak', {pedantic: false}],
  ['definition', {commonmark: false}]
];

// Nodes that can interupt a blockquote:
//
// ```markdown
// > A paragraph.
// ___
// ```
//
// In the above example, the thematic break “interupts” the blockquote.
proto$3.interruptBlockquote = [
  ['indentedCode', {commonmark: true}],
  ['fencedCode', {commonmark: true}],
  ['atxHeading', {commonmark: true}],
  ['setextHeading', {commonmark: true}],
  ['thematicBreak', {commonmark: true}],
  ['html', {commonmark: true}],
  ['list', {commonmark: true}],
  ['definition', {commonmark: false}]
];

// Handlers.
proto$3.blockTokenizers = {
  blankLine: requireBlankLine(),
  indentedCode: requireCodeIndented(),
  fencedCode: requireCodeFenced(),
  blockquote: requireBlockquote(),
  atxHeading: requireHeadingAtx(),
  thematicBreak: requireThematicBreak(),
  list: requireList(),
  setextHeading: requireHeadingSetext(),
  html: requireHtmlBlock(),
  definition: requireDefinition(),
  table: requireTable(),
  paragraph: requireParagraph()
};

proto$3.inlineTokenizers = {
  escape: require_escape(),
  autoLink: requireAutoLink(),
  url: requireUrl(),
  email: requireEmail(),
  html: requireHtmlInline(),
  link: requireLink(),
  reference: requireReference(),
  strong: requireStrong(),
  emphasis: requireEmphasis(),
  deletion: require_delete(),
  code: requireCodeInline(),
  break: require_break(),
  text: requireText()
};

// Expose precedence.
proto$3.blockMethods = keys$1(proto$3.blockTokenizers);
proto$3.inlineMethods = keys$1(proto$3.inlineTokenizers);

// Tokenizers.
proto$3.tokenizeBlock = tokenizer$2('block');
proto$3.tokenizeInline = tokenizer$2('inline');
proto$3.tokenizeFactory = tokenizer$2;

// Get all keys in `value`.
function keys$1(value) {
  var result = [];
  var key;

  for (key in value) {
    result.push(key);
  }

  return result
}

var unherit$1 = unherit_1;
var xtend$1 = immutable;
var Parser = parser$1;

var remarkParse = parse$5;
parse$5.Parser = Parser;

function parse$5(options) {
  var settings = this.data('settings');
  var Local = unherit$1(Parser);

  Local.prototype.options = xtend$1(Local.prototype.options, settings, options);

  this.Parser = Local;
}

var markdown = /*@__PURE__*/getDefaultExportFromCjs(remarkParse);

var visit$4 = unistUtilVisit;

var mdastUtilDefinitions$1 = getDefinitionFactory$1;

var own$7 = {}.hasOwnProperty;

// Get a definition in `node` by `identifier`.
function getDefinitionFactory$1(node, options) {
  return getterFactory$1(gather$1(node, options))
}

// Gather all definitions in `node`
function gather$1(node, options) {
  var cache = {};

  if (!node || !node.type) {
    throw new Error('mdast-util-definitions expected node')
  }

  visit$4(node, 'definition', options && options.commonmark ? commonmark : normal);

  return cache

  function commonmark(definition) {
    var id = normalise$1(definition.identifier);
    if (!own$7.call(cache, id)) {
      cache[id] = definition;
    }
  }

  function normal(definition) {
    cache[normalise$1(definition.identifier)] = definition;
  }
}

// Factory to get a node from the given definition-cache.
function getterFactory$1(cache) {
  return getter

  // Get a node from the bound definition-cache.
  function getter(identifier) {
    var id = identifier && normalise$1(identifier);
    return id && own$7.call(cache, id) ? cache[id] : null
  }
}

function normalise$1(identifier) {
  return identifier.toUpperCase()
}

var spaceSeparatedTokens = {};

spaceSeparatedTokens.parse = parse$4;
spaceSeparatedTokens.stringify = stringify$4;

var empty$2 = '';
var space$2 = ' ';
var whiteSpace$3 = /[ \t\n\r\f]+/g;

function parse$4(value) {
  var input = String(value || empty$2).trim();
  return input === empty$2 ? [] : input.split(whiteSpace$3)
}

function stringify$4(values) {
  return values.join(space$2).trim()
}

var isAbsoluteUrl = url => {
	if (typeof url !== 'string') {
		throw new TypeError(`Expected a \`string\`, got \`${typeof url}\``);
	}

	// Don't match Windows paths `c:\`
	if (/^[a-zA-Z]:\\/.test(url)) {
		return false;
	}

	// Scheme: https://tools.ietf.org/html/rfc3986#section-3.1
	// Absolute URL: https://tools.ietf.org/html/rfc3986#section-4.3
	return /^[a-zA-Z][a-zA-Z\d+\-.]*:/.test(url);
};

var visit$3 = unistUtilVisit;
var definitions$1 = mdastUtilDefinitions$1;
var spaceSeparated$3 = spaceSeparatedTokens.parse;
var absolute = isAbsoluteUrl;
var extend$2 = extend$5;

var remarkExternalLinks = externalLinks;

var defaultTarget = '_blank';
var defaultRel = ['nofollow', 'noopener', 'noreferrer'];
var defaultProtocols = ['http', 'https'];

function externalLinks(options) {
  var settings = options || {};
  var target = settings.target;
  var rel = settings.rel;
  var protocols = settings.protocols || defaultProtocols;
  var content = settings.content;
  var contentProperties = settings.contentProperties || {};

  if (typeof rel === 'string') {
    rel = spaceSeparated$3(rel);
  }

  if (content && typeof content === 'object' && !('length' in content)) {
    content = [content];
  }

  return transform

  function transform(tree) {
    var definition = definitions$1(tree);

    visit$3(tree, ['link', 'linkReference'], visitor);

    function visitor(node) {
      var ctx = node.type === 'link' ? node : definition(node.identifier);
      var protocol;
      var data;
      var props;

      if (!ctx) return

      protocol = ctx.url.slice(0, ctx.url.indexOf(':'));

      if (absolute(ctx.url) && protocols.indexOf(protocol) !== -1) {
        data = node.data || (node.data = {});
        props = data.hProperties || (data.hProperties = {});

        if (target !== false) {
          props.target = target || defaultTarget;
        }

        if (rel !== false) {
          props.rel = (rel || defaultRel).concat();
        }

        if (content) {
          // `fragment` is not a known mdast node, but unknown nodes with
          // children are handled as elements by `mdast-util-to-hast`:
          // See: <https://github.com/syntax-tree/mdast-util-to-hast#notes>.
          node.children.push({
            type: 'fragment',
            children: [],
            data: {
              hName: 'span',
              hProperties: extend$2(true, contentProperties),
              hChildren: extend$2(true, content)
            }
          });
        }
      }
    }
  }
}

var external = /*@__PURE__*/getDefaultExportFromCjs(remarkExternalLinks);

var format = {exports: {}};

format.exports;

(function (module) {
(function() {

	  //// Export the API
	  var namespace;

	  // CommonJS / Node module
	  {
	    namespace = module.exports = format;
	  }

	  namespace.format = format;
	  namespace.vsprintf = vsprintf;

	  if (typeof console !== 'undefined' && typeof console.log === 'function') {
	    namespace.printf = printf;
	  }

	  function printf(/* ... */) {
	    console.log(format.apply(null, arguments));
	  }

	  function vsprintf(fmt, replacements) {
	    return format.apply(null, [fmt].concat(replacements));
	  }

	  function format(fmt) {
	    var argIndex = 1 // skip initial format argument
	      , args = [].slice.call(arguments)
	      , i = 0
	      , n = fmt.length
	      , result = ''
	      , c
	      , escaped = false
	      , arg
	      , tmp
	      , leadingZero = false
	      , precision
	      , nextArg = function() { return args[argIndex++]; }
	      , slurpNumber = function() {
	          var digits = '';
	          while (/\d/.test(fmt[i])) {
	            digits += fmt[i++];
	            c = fmt[i];
	          }
	          return digits.length > 0 ? parseInt(digits) : null;
	        }
	      ;
	    for (; i < n; ++i) {
	      c = fmt[i];
	      if (escaped) {
	        escaped = false;
	        if (c == '.') {
	          leadingZero = false;
	          c = fmt[++i];
	        }
	        else if (c == '0' && fmt[i + 1] == '.') {
	          leadingZero = true;
	          i += 2;
	          c = fmt[i];
	        }
	        else {
	          leadingZero = true;
	        }
	        precision = slurpNumber();
	        switch (c) {
	        case 'b': // number in binary
	          result += parseInt(nextArg(), 10).toString(2);
	          break;
	        case 'c': // character
	          arg = nextArg();
	          if (typeof arg === 'string' || arg instanceof String)
	            result += arg;
	          else
	            result += String.fromCharCode(parseInt(arg, 10));
	          break;
	        case 'd': // number in decimal
	          result += parseInt(nextArg(), 10);
	          break;
	        case 'f': // floating point number
	          tmp = String(parseFloat(nextArg()).toFixed(precision || 6));
	          result += leadingZero ? tmp : tmp.replace(/^0/, '');
	          break;
	        case 'j': // JSON
	          result += JSON.stringify(nextArg());
	          break;
	        case 'o': // number in octal
	          result += '0' + parseInt(nextArg(), 10).toString(8);
	          break;
	        case 's': // string
	          result += nextArg();
	          break;
	        case 'x': // lowercase hexadecimal
	          result += '0x' + parseInt(nextArg(), 10).toString(16);
	          break;
	        case 'X': // uppercase hexadecimal
	          result += '0x' + parseInt(nextArg(), 10).toString(16).toUpperCase();
	          break;
	        default:
	          result += c;
	          break;
	        }
	      } else if (c === '%') {
	        escaped = true;
	      } else {
	        result += c;
	      }
	    }
	    return result;
	  }

	}()); 
} (format));

var formatExports = format.exports;

var formatter = formatExports;

var fault$1 = create$9(Error);

var fault_1 = fault$1;

fault$1.eval = create$9(EvalError);
fault$1.range = create$9(RangeError);
fault$1.reference = create$9(ReferenceError);
fault$1.syntax = create$9(SyntaxError);
fault$1.type = create$9(TypeError);
fault$1.uri = create$9(URIError);

fault$1.create = create$9;

// Create a new `EConstructor`, with the formatted `format` as a first argument.
function create$9(EConstructor) {
  FormattedError.displayName = EConstructor.displayName || EConstructor.name;

  return FormattedError

  function FormattedError(format) {
    if (format) {
      format = formatter.apply(null, arguments);
    }

    return new EConstructor(format)
  }
}

var fault = fault_1;

var matters_1 = matters$1;

var own$6 = {}.hasOwnProperty;

var markers = {
  yaml: '-',
  toml: '+'
};

function matters$1(options) {
  var results = [];
  var index = -1;
  var length;

  // One preset or matter.
  if (typeof options === 'string' || !('length' in options)) {
    options = [options];
  }

  length = options.length;

  while (++index < length) {
    results[index] = matter(options[index]);
  }

  return results
}

function matter(option) {
  var result = option;

  if (typeof result === 'string') {
    if (!own$6.call(markers, result)) {
      throw fault('Missing matter definition for `%s`', result)
    }

    result = {type: result, marker: markers[result]};
  } else if (typeof result !== 'object') {
    throw fault('Expected matter to be an object, not `%j`', result)
  }

  if (!own$6.call(result, 'type')) {
    throw fault('Missing `type` in matter `%j`', result)
  }

  if (!own$6.call(result, 'fence') && !own$6.call(result, 'marker')) {
    throw fault('Missing `marker` or `fence` in matter `%j`', result)
  }

  return result
}

var fence_1 = fence$2;

function fence$2(matter, prop) {
  var marker;

  if (matter.marker) {
    marker = pick(matter.marker, prop);
    return marker + marker + marker
  }

  return pick(matter.fence, prop)
}

function pick(schema, prop) {
  return typeof schema === 'string' ? schema : schema[prop]
}

var fence$1 = fence_1;

var parse$3 = create$8;

function create$8(matter) {
  var name = matter.type + 'FrontMatter';
  var open = fence$1(matter, 'open');
  var close = fence$1(matter, 'close');
  var newline = '\n';
  var anywhere = matter.anywhere;

  frontmatter.displayName = name;
  frontmatter.onlyAtStart = typeof anywhere === 'boolean' ? !anywhere : true;

  return [name, frontmatter]

  function frontmatter(eat, value, silent) {
    var index = open.length;
    var offset;

    if (value.slice(0, index) !== open || value.charAt(index) !== newline) {
      return
    }

    offset = value.indexOf(close, index);

    while (offset !== -1 && value.charAt(offset - 1) !== newline) {
      index = offset + close.length;
      offset = value.indexOf(close, index);
    }

    if (offset !== -1) {
      /* istanbul ignore if - never used (yet) */
      if (silent) {
        return true
      }

      return eat(value.slice(0, offset + close.length))({
        type: matter.type,
        value: value.slice(open.length + 1, offset - 1)
      })
    }
  }
}

var fence = fence_1;

var compile$1 = create$7;

function create$7(matter) {
  var type = matter.type;
  var open = fence(matter, 'open');
  var close = fence(matter, 'close');

  frontmatter.displayName = type + 'FrontMatter';

  return [type, frontmatter]

  function frontmatter(node) {
    return open + (node.value ? '\n' + node.value : '') + '\n' + close
  }
}

var matters = matters_1;
var parse$2 = parse$3;
var compile = compile$1;

var remarkFrontmatter = frontmatter;

function frontmatter(options) {
  var parser = this.Parser;
  var compiler = this.Compiler;
  var config = matters(options || ['yaml']);

  if (isRemarkParser(parser)) {
    attachParser(parser, config);
  }

  if (isRemarkCompiler(compiler)) {
    attachCompiler(compiler, config);
  }
}

function attachParser(parser, matters) {
  var proto = parser.prototype;
  var tokenizers = wrap$6(parse$2, matters);
  var names = [];
  var key;

  for (key in tokenizers) {
    names.push(key);
  }

  proto.blockMethods = names.concat(proto.blockMethods);
  proto.blockTokenizers = Object.assign({}, tokenizers, proto.blockTokenizers);
}

function attachCompiler(compiler, matters) {
  var proto = compiler.prototype;
  proto.visitors = Object.assign({}, wrap$6(compile, matters), proto.visitors);
}

function wrap$6(func, matters) {
  var result = {};
  var length = matters.length;
  var index = -1;
  var tuple;

  while (++index < length) {
    tuple = func(matters[index]);
    result[tuple[0]] = tuple[1];
  }

  return result
}

function isRemarkParser(parser) {
  return Boolean(parser && parser.prototype && parser.prototype.blockTokenizers)
}

function isRemarkCompiler(compiler) {
  return Boolean(compiler && compiler.prototype && compiler.prototype.visitors)
}

var extract_frontmatter = /*@__PURE__*/getDefaultExportFromCjs(remarkFrontmatter);

var unistBuilder = u$b;

function u$b(type, props, value) {
  var node;

  if (
    (value === null || value === undefined) &&
    (typeof props !== 'object' || Array.isArray(props))
  ) {
    value = props;
    props = {};
  }

  node = Object.assign({type: String(type)}, props);

  if (Array.isArray(value)) {
    node.children = value;
  } else if (value !== null && value !== undefined) {
    node.value = String(value);
  }

  return node
}

var start = factory$1('start');
var end = factory$1('end');

var unistUtilPosition = position$2;

position$2.start = start;
position$2.end = end;

function position$2(node) {
  return {start: start(node), end: end(node)}
}

function factory$1(type) {
  point.displayName = type;

  return point

  function point(node) {
    var point = (node && node.position && node.position[type]) || {};

    return {
      line: point.line || null,
      column: point.column || null,
      offset: isNaN(point.offset) ? null : point.offset
    }
  }
}

var unistUtilGenerated = generated$1;

function generated$1(node) {
  return (
    !node ||
    !node.position ||
    !node.position.start ||
    !node.position.start.line ||
    !node.position.start.column ||
    !node.position.end ||
    !node.position.end.line ||
    !node.position.end.column
  )
}

var visit$2 = unistUtilVisit;

var mdastUtilDefinitions = getDefinitionFactory;

var own$5 = {}.hasOwnProperty;

// Get a definition in `node` by `identifier`.
function getDefinitionFactory(node, options) {
  return getterFactory(gather(node, options))
}

// Gather all definitions in `node`
function gather(node, options) {
  var cache = {};

  if (!node || !node.type) {
    throw new Error('mdast-util-definitions expected node')
  }

  visit$2(node, 'definition', options && options.commonmark ? commonmark : normal);

  return cache

  function commonmark(definition) {
    var id = normalise(definition.identifier);
    if (!own$5.call(cache, id)) {
      cache[id] = definition;
    }
  }

  function normal(definition) {
    cache[normalise(definition.identifier)] = definition;
  }
}

// Factory to get a node from the given definition-cache.
function getterFactory(cache) {
  return getter

  // Get a node from the bound definition-cache.
  function getter(identifier) {
    var id = identifier && normalise(identifier);
    return id && own$5.call(cache, id) ? cache[id] : null
  }
}

function normalise(identifier) {
  return identifier.toUpperCase()
}

var all_1$1;
var hasRequiredAll$1;

function requireAll$1 () {
	if (hasRequiredAll$1) return all_1$1;
	hasRequiredAll$1 = 1;

	all_1$1 = all;

	var one = requireOne$1();

	function all(h, parent) {
	  var nodes = parent.children || [];
	  var length = nodes.length;
	  var values = [];
	  var index = -1;
	  var result;
	  var head;

	  while (++index < length) {
	    result = one(h, nodes[index], parent);

	    if (result) {
	      if (index && nodes[index - 1].type === 'break') {
	        if (result.value) {
	          result.value = result.value.replace(/^\s+/, '');
	        }

	        head = result.children && result.children[0];

	        if (head && head.value) {
	          head.value = head.value.replace(/^\s+/, '');
	        }
	      }

	      values = values.concat(result);
	    }
	  }

	  return values
	}
	return all_1$1;
}

var one_1$1;
var hasRequiredOne$1;

function requireOne$1 () {
	if (hasRequiredOne$1) return one_1$1;
	hasRequiredOne$1 = 1;

	one_1$1 = one;

	var u = unistBuilder;
	var all = requireAll$1();

	var own = {}.hasOwnProperty;

	// Transform an unknown node.
	function unknown(h, node) {
	  if (text(node)) {
	    return h.augment(node, u('text', node.value))
	  }

	  return h(node, 'div', all(h, node))
	}

	// Visit a node.
	function one(h, node, parent) {
	  var type = node && node.type;
	  var fn = own.call(h.handlers, type) ? h.handlers[type] : h.unknownHandler;

	  // Fail on non-nodes.
	  if (!type) {
	    throw new Error('Expected node, got `' + node + '`')
	  }

	  return (typeof fn === 'function' ? fn : unknown)(h, node, parent)
	}

	// Check if the node should be renderered as a text node.
	function text(node) {
	  var data = node.data || {};

	  if (
	    own.call(data, 'hName') ||
	    own.call(data, 'hProperties') ||
	    own.call(data, 'hChildren')
	  ) {
	    return false
	  }

	  return 'value' in node
	}
	return one_1$1;
}

var thematicBreak_1 = thematicBreak$1;

function thematicBreak$1(h, node) {
  return h(node, 'hr')
}

var wrap_1 = wrap$5;

var u$a = unistBuilder;

// Wrap `nodes` with line feeds between each entry.
// Optionally adds line feeds at the start and end.
function wrap$5(nodes, loose) {
  var result = [];
  var index = -1;
  var length = nodes.length;

  if (loose) {
    result.push(u$a('text', '\n'));
  }

  while (++index < length) {
    if (index) {
      result.push(u$a('text', '\n'));
    }

    result.push(nodes[index]);
  }

  if (loose && nodes.length !== 0) {
    result.push(u$a('text', '\n'));
  }

  return result
}

var list_1 = list$1;

var wrap$4 = wrap_1;
var all$d = requireAll$1();

function list$1(h, node) {
  var props = {};
  var name = node.ordered ? 'ol' : 'ul';
  var items;
  var index = -1;
  var length;

  if (typeof node.start === 'number' && node.start !== 1) {
    props.start = node.start;
  }

  items = all$d(h, node);
  length = items.length;

  // Like GitHub, add a class for custom styling.
  while (++index < length) {
    if (
      items[index].properties.className &&
      items[index].properties.className.indexOf('task-list-item') !== -1
    ) {
      props.className = ['contains-task-list'];
      break
    }
  }

  return h(node, name, props, wrap$4(items, true))
}

var footer$1 = generateFootnotes;

var thematicBreak = thematicBreak_1;
var list = list_1;
var wrap$3 = wrap_1;

function generateFootnotes(h) {
  var footnoteById = h.footnoteById;
  var footnoteOrder = h.footnoteOrder;
  var length = footnoteOrder.length;
  var index = -1;
  var listItems = [];
  var def;
  var backReference;
  var content;
  var tail;

  while (++index < length) {
    def = footnoteById[footnoteOrder[index].toUpperCase()];

    if (!def) {
      continue
    }

    content = def.children.concat();
    tail = content[content.length - 1];
    backReference = {
      type: 'link',
      url: '#fnref-' + def.identifier,
      data: {hProperties: {className: ['footnote-backref']}},
      children: [{type: 'text', value: '↩'}]
    };

    if (!tail || tail.type !== 'paragraph') {
      tail = {type: 'paragraph', children: []};
      content.push(tail);
    }

    tail.children.push(backReference);

    listItems.push({
      type: 'listItem',
      data: {hProperties: {id: 'fn-' + def.identifier}},
      children: content,
      position: def.position
    });
  }

  if (listItems.length === 0) {
    return null
  }

  return h(
    null,
    'div',
    {className: ['footnotes']},
    wrap$3(
      [
        thematicBreak(h),
        list(h, {type: 'list', ordered: true, children: listItems})
      ],
      true
    )
  )
}

var blockquote_1 = blockquote;

var wrap$2 = wrap_1;
var all$c = requireAll$1();

function blockquote(h, node) {
  return h(node, 'blockquote', wrap$2(all$c(h, node), true))
}

var _break = hardBreak;

var u$9 = unistBuilder;

function hardBreak(h, node) {
  return [h(node, 'br'), u$9('text', '\n')]
}

var code_1 = code;

var u$8 = unistBuilder;

function code(h, node) {
  var value = node.value ? node.value + '\n' : '';
  var lang = node.lang && node.lang.match(/^[^ \t]+(?=[ \t]|$)/);
  var props = {};

  if (lang) {
    props.className = ['language-' + lang];
  }

  return h(node.position, 'pre', [h(node, 'code', props, [u$8('text', value)])])
}

var _delete = strikethrough;

var all$b = requireAll$1();

function strikethrough(h, node) {
  return h(node, 'del', all$b(h, node))
}

var emphasis_1 = emphasis;

var all$a = requireAll$1();

function emphasis(h, node) {
  return h(node, 'em', all$a(h, node))
}

var footnoteReference_1 = footnoteReference$1;

var u$7 = unistBuilder;

function footnoteReference$1(h, node) {
  var footnoteOrder = h.footnoteOrder;
  var identifier = String(node.identifier);

  if (footnoteOrder.indexOf(identifier) === -1) {
    footnoteOrder.push(identifier);
  }

  return h(node.position, 'sup', {id: 'fnref-' + identifier}, [
    h(node, 'a', {href: '#fn-' + identifier, className: ['footnote-ref']}, [
      u$7('text', node.label || identifier)
    ])
  ])
}

var footnote_1 = footnote;

var footnoteReference = footnoteReference_1;

function footnote(h, node) {
  var footnoteById = h.footnoteById;
  var footnoteOrder = h.footnoteOrder;
  var identifier = 1;

  while (identifier in footnoteById) {
    identifier++;
  }

  identifier = String(identifier);

  // No need to check if `identifier` exists in `footnoteOrder`, it’s guaranteed
  // to not exist because we just generated it.
  footnoteOrder.push(identifier);

  footnoteById[identifier] = {
    type: 'footnoteDefinition',
    identifier: identifier,
    children: [{type: 'paragraph', children: node.children}],
    position: node.position
  };

  return footnoteReference(h, {
    type: 'footnoteReference',
    identifier: identifier,
    position: node.position
  })
}

var heading_1 = heading;

var all$9 = requireAll$1();

function heading(h, node) {
  return h(node, 'h' + node.depth, all$9(h, node))
}

var html_1$1 = html$5;

var u$6 = unistBuilder;

// Return either a `raw` node in dangerous mode, otherwise nothing.
function html$5(h, node) {
  return h.dangerous ? h.augment(node, u$6('raw', node.value)) : null
}

var encodeCache = {};


// Create a lookup array where anything but characters in `chars` string
// and alphanumeric chars is percent-encoded.
//
function getEncodeCache(exclude) {
  var i, ch, cache = encodeCache[exclude];
  if (cache) { return cache; }

  cache = encodeCache[exclude] = [];

  for (i = 0; i < 128; i++) {
    ch = String.fromCharCode(i);

    if (/^[0-9a-z]$/i.test(ch)) {
      // always allow unencoded alphanumeric characters
      cache.push(ch);
    } else {
      cache.push('%' + ('0' + i.toString(16).toUpperCase()).slice(-2));
    }
  }

  for (i = 0; i < exclude.length; i++) {
    cache[exclude.charCodeAt(i)] = exclude[i];
  }

  return cache;
}


// Encode unsafe characters with percent-encoding, skipping already
// encoded sequences.
//
//  - string       - string to encode
//  - exclude      - list of characters to ignore (in addition to a-zA-Z0-9)
//  - keepEscaped  - don't encode '%' in a correct escape sequence (default: true)
//
function encode$1(string, exclude, keepEscaped) {
  var i, l, code, nextCode, cache,
      result = '';

  if (typeof exclude !== 'string') {
    // encode(string, keepEscaped)
    keepEscaped  = exclude;
    exclude = encode$1.defaultChars;
  }

  if (typeof keepEscaped === 'undefined') {
    keepEscaped = true;
  }

  cache = getEncodeCache(exclude);

  for (i = 0, l = string.length; i < l; i++) {
    code = string.charCodeAt(i);

    if (keepEscaped && code === 0x25 /* % */ && i + 2 < l) {
      if (/^[0-9a-f]{2}$/i.test(string.slice(i + 1, i + 3))) {
        result += string.slice(i, i + 3);
        i += 2;
        continue;
      }
    }

    if (code < 128) {
      result += cache[code];
      continue;
    }

    if (code >= 0xD800 && code <= 0xDFFF) {
      if (code >= 0xD800 && code <= 0xDBFF && i + 1 < l) {
        nextCode = string.charCodeAt(i + 1);
        if (nextCode >= 0xDC00 && nextCode <= 0xDFFF) {
          result += encodeURIComponent(string[i] + string[i + 1]);
          i++;
          continue;
        }
      }
      result += '%EF%BF%BD';
      continue;
    }

    result += encodeURIComponent(string[i]);
  }

  return result;
}

encode$1.defaultChars   = ";/?:@&=+$,-_.!~*'()#";
encode$1.componentChars = "-_.!~*'()";


var encode_1 = encode$1;

var revert_1 = revert$2;

var u$5 = unistBuilder;
var all$8 = requireAll$1();

// Return the content of a reference without definition as Markdown.
function revert$2(h, node) {
  var subtype = node.referenceType;
  var suffix = ']';
  var contents;
  var head;
  var tail;

  if (subtype === 'collapsed') {
    suffix += '[]';
  } else if (subtype === 'full') {
    suffix += '[' + (node.label || node.identifier) + ']';
  }

  if (node.type === 'imageReference') {
    return u$5('text', '![' + node.alt + suffix)
  }

  contents = all$8(h, node);
  head = contents[0];

  if (head && head.type === 'text') {
    head.value = '[' + head.value;
  } else {
    contents.unshift(u$5('text', '['));
  }

  tail = contents[contents.length - 1];

  if (tail && tail.type === 'text') {
    tail.value += suffix;
  } else {
    contents.push(u$5('text', suffix));
  }

  return contents
}

var imageReference_1 = imageReference;

var normalize$6 = encode_1;
var revert$1 = revert_1;

function imageReference(h, node) {
  var def = h.definition(node.identifier);
  var props;

  if (!def) {
    return revert$1(h, node)
  }

  props = {src: normalize$6(def.url || ''), alt: node.alt};

  if (def.title !== null && def.title !== undefined) {
    props.title = def.title;
  }

  return h(node, 'img', props)
}

var normalize$5 = encode_1;

var image_1 = image$1;

function image$1(h, node) {
  var props = {src: normalize$5(node.url), alt: node.alt};

  if (node.title !== null && node.title !== undefined) {
    props.title = node.title;
  }

  return h(node, 'img', props)
}

var inlineCode_1 = inlineCode;

var u$4 = unistBuilder;

function inlineCode(h, node) {
  var value = node.value.replace(/\r?\n|\r/g, ' ');
  return h(node, 'code', [u$4('text', value)])
}

var linkReference_1 = linkReference;

var normalize$4 = encode_1;
var revert = revert_1;
var all$7 = requireAll$1();

function linkReference(h, node) {
  var def = h.definition(node.identifier);
  var props;

  if (!def) {
    return revert(h, node)
  }

  props = {href: normalize$4(def.url || '')};

  if (def.title !== null && def.title !== undefined) {
    props.title = def.title;
  }

  return h(node, 'a', props, all$7(h, node))
}

var normalize$3 = encode_1;
var all$6 = requireAll$1();

var link_1 = link;

function link(h, node) {
  var props = {href: normalize$3(node.url)};

  if (node.title !== null && node.title !== undefined) {
    props.title = node.title;
  }

  return h(node, 'a', props, all$6(h, node))
}

var listItem_1 = listItem$1;

var u$3 = unistBuilder;
var all$5 = requireAll$1();

function listItem$1(h, node, parent) {
  var result = all$5(h, node);
  var head = result[0];
  var loose = parent ? listLoose(parent) : listItemLoose(node);
  var props = {};
  var wrapped = [];
  var length;
  var index;
  var child;

  if (typeof node.checked === 'boolean') {
    if (!head || head.tagName !== 'p') {
      head = h(null, 'p', []);
      result.unshift(head);
    }

    if (head.children.length !== 0) {
      head.children.unshift(u$3('text', ' '));
    }

    head.children.unshift(
      h(null, 'input', {
        type: 'checkbox',
        checked: node.checked,
        disabled: true
      })
    );

    // According to github-markdown-css, this class hides bullet.
    // See: <https://github.com/sindresorhus/github-markdown-css>.
    props.className = ['task-list-item'];
  }

  length = result.length;
  index = -1;

  while (++index < length) {
    child = result[index];

    // Add eols before nodes, except if this is a loose, first paragraph.
    if (loose || index !== 0 || child.tagName !== 'p') {
      wrapped.push(u$3('text', '\n'));
    }

    if (child.tagName === 'p' && !loose) {
      wrapped = wrapped.concat(child.children);
    } else {
      wrapped.push(child);
    }
  }

  // Add a final eol.
  if (length && (loose || child.tagName !== 'p')) {
    wrapped.push(u$3('text', '\n'));
  }

  return h(node, 'li', props, wrapped)
}

function listLoose(node) {
  var loose = node.spread;
  var children = node.children;
  var length = children.length;
  var index = -1;

  while (!loose && ++index < length) {
    loose = listItemLoose(children[index]);
  }

  return loose
}

function listItemLoose(node) {
  var spread = node.spread;

  return spread === undefined || spread === null
    ? node.children.length > 1
    : spread
}

var paragraph_1 = paragraph;

var all$4 = requireAll$1();

function paragraph(h, node) {
  return h(node, 'p', all$4(h, node))
}

var root_1 = root;

var u$2 = unistBuilder;
var wrap$1 = wrap_1;
var all$3 = requireAll$1();

function root(h, node) {
  return h.augment(node, u$2('root', wrap$1(all$3(h, node))))
}

var strong_1 = strong;

var all$2 = requireAll$1();

function strong(h, node) {
  return h(node, 'strong', all$2(h, node))
}

var table_1 = table;

var position$1 = unistUtilPosition;
var wrap = wrap_1;
var all$1 = requireAll$1();

function table(h, node) {
  var rows = node.children;
  var index = rows.length;
  var align = node.align || [];
  var alignLength = align.length;
  var result = [];
  var pos;
  var row;
  var out;
  var name;
  var cell;

  while (index--) {
    row = rows[index].children;
    name = index === 0 ? 'th' : 'td';
    pos = alignLength || row.length;
    out = [];

    while (pos--) {
      cell = row[pos];
      out[pos] = h(cell, name, {align: align[pos]}, cell ? all$1(h, cell) : []);
    }

    result[index] = h(rows[index], 'tr', wrap(out, true));
  }

  return h(
    node,
    'table',
    wrap(
      [h(result[0].position, 'thead', wrap([result[0]], true))].concat(
        result[1]
          ? h(
              {
                start: position$1.start(result[1]),
                end: position$1.end(result[result.length - 1])
              },
              'tbody',
              wrap(result.slice(1), true)
            )
          : []
      ),
      true
    )
  )
}

var text_1$1 = text$1;

var u$1 = unistBuilder;

function text$1(h, node) {
  return h.augment(
    node,
    u$1('text', String(node.value).replace(/[ \t]*(\r?\n|\r)[ \t]*/g, '$1'))
  )
}

var handlers$1 = {
  blockquote: blockquote_1,
  break: _break,
  code: code_1,
  delete: _delete,
  emphasis: emphasis_1,
  footnoteReference: footnoteReference_1,
  footnote: footnote_1,
  heading: heading_1,
  html: html_1$1,
  imageReference: imageReference_1,
  image: image_1,
  inlineCode: inlineCode_1,
  linkReference: linkReference_1,
  link: link_1,
  listItem: listItem_1,
  list: list_1,
  paragraph: paragraph_1,
  root: root_1,
  strong: strong_1,
  table: table_1,
  text: text_1$1,
  thematicBreak: thematicBreak_1,
  toml: ignore,
  yaml: ignore,
  definition: ignore,
  footnoteDefinition: ignore
};

// Return nothing for nodes that are ignored.
function ignore() {
  return null
}

var lib$2 = toHast;

var u = unistBuilder;
var visit$1 = unistUtilVisit;
var position = unistUtilPosition;
var generated = unistUtilGenerated;
var definitions = mdastUtilDefinitions;
var one$2 = requireOne$1();
var footer = footer$1;
var handlers = handlers$1;

var own$4 = {}.hasOwnProperty;

var deprecationWarningIssued = false;

// Factory to transform.
function factory(tree, options) {
  var settings = options || {};

  // Issue a warning if the deprecated tag 'allowDangerousHTML' is used
  if (settings.allowDangerousHTML !== undefined && !deprecationWarningIssued) {
    deprecationWarningIssued = true;
    console.warn(
      'mdast-util-to-hast: deprecation: `allowDangerousHTML` is nonstandard, use `allowDangerousHtml` instead'
    );
  }

  var dangerous = settings.allowDangerousHtml || settings.allowDangerousHTML;
  var footnoteById = {};

  h.dangerous = dangerous;
  h.definition = definitions(tree, settings);
  h.footnoteById = footnoteById;
  h.footnoteOrder = [];
  h.augment = augment;
  h.handlers = Object.assign({}, handlers, settings.handlers);
  h.unknownHandler = settings.unknownHandler;

  visit$1(tree, 'footnoteDefinition', onfootnotedefinition);

  return h

  // Finalise the created `right`, a hast node, from `left`, an mdast node.
  function augment(left, right) {
    var data;
    var ctx;

    // Handle `data.hName`, `data.hProperties, `data.hChildren`.
    if (left && 'data' in left) {
      data = left.data;

      if (right.type === 'element' && data.hName) {
        right.tagName = data.hName;
      }

      if (right.type === 'element' && data.hProperties) {
        right.properties = Object.assign({}, right.properties, data.hProperties);
      }

      if (right.children && data.hChildren) {
        right.children = data.hChildren;
      }
    }

    ctx = left && left.position ? left : {position: left};

    if (!generated(ctx)) {
      right.position = {
        start: position.start(ctx),
        end: position.end(ctx)
      };
    }

    return right
  }

  // Create an element for `node`.
  function h(node, tagName, props, children) {
    if (
      (children === undefined || children === null) &&
      typeof props === 'object' &&
      'length' in props
    ) {
      children = props;
      props = {};
    }

    return augment(node, {
      type: 'element',
      tagName: tagName,
      properties: props || {},
      children: children || []
    })
  }

  function onfootnotedefinition(definition) {
    var id = String(definition.identifier).toUpperCase();

    // Mimick CM behavior of link definitions.
    // See: <https://github.com/syntax-tree/mdast-util-definitions/blob/8d48e57/index.js#L26>.
    if (!own$4.call(footnoteById, id)) {
      footnoteById[id] = definition;
    }
  }
}

// Transform `tree`, which is an mdast node, to a hast node.
function toHast(tree, options) {
  var h = factory(tree, options);
  var node = one$2(h, tree);
  var foot = footer(h);

  if (foot) {
    node.children = node.children.concat(u('text', '\n'), foot);
  }

  return node
}

var mdastUtilToHast = lib$2;

var mdast2hast = mdastUtilToHast;

var remarkRehype = remark2rehype;

// Attacher.
// If a destination is given, runs the destination with the new hast tree
// (bridge mode).
// Without destination, returns the tree: further plugins run on that tree
// (mutate mode).
function remark2rehype(destination, options) {
  if (destination && !destination.process) {
    options = destination;
    destination = null;
  }

  return destination ? bridge(destination, options) : mutate(options)
}

// Bridge mode.
// Runs the destination with the new hast tree.
function bridge(destination, options) {
  return transformer

  function transformer(node, file, next) {
    destination.run(mdast2hast(node, options), file, done);

    function done(err) {
      next(err);
    }
  }
}

// Mutate-mode.
// Further transformers run on the hast tree.
function mutate(options) {
  return transformer

  function transformer(node) {
    return mdast2hast(node, options)
  }
}

var remark2rehype$1 = /*@__PURE__*/getDefaultExportFromCjs(remarkRehype);

var schema$1 = Schema$8;

var proto$2 = Schema$8.prototype;

proto$2.space = null;
proto$2.normal = {};
proto$2.property = {};

function Schema$8(property, normal, space) {
  this.property = property;
  this.normal = normal;

  if (space) {
    this.space = space;
  }
}

var xtend = immutable;
var Schema$7 = schema$1;

var merge_1 = merge$3;

function merge$3(definitions) {
  var length = definitions.length;
  var property = [];
  var normal = [];
  var index = -1;
  var info;
  var space;

  while (++index < length) {
    info = definitions[index];
    property.push(info.property);
    normal.push(info.normal);
    space = info.space;
  }

  return new Schema$7(
    xtend.apply(null, property),
    xtend.apply(null, normal),
    space
  )
}

var normalize_1 = normalize$2;

function normalize$2(value) {
  return value.toLowerCase()
}

var info = Info$2;

var proto$1 = Info$2.prototype;

proto$1.space = null;
proto$1.attribute = null;
proto$1.property = null;
proto$1.boolean = false;
proto$1.booleanish = false;
proto$1.overloadedBoolean = false;
proto$1.number = false;
proto$1.commaSeparated = false;
proto$1.spaceSeparated = false;
proto$1.commaOrSpaceSeparated = false;
proto$1.mustUseProperty = false;
proto$1.defined = false;

function Info$2(property, attribute) {
  this.property = property;
  this.attribute = attribute;
}

var types$4 = {};

var powers = 0;

types$4.boolean = increment();
types$4.booleanish = increment();
types$4.overloadedBoolean = increment();
types$4.number = increment();
types$4.spaceSeparated = increment();
types$4.commaSeparated = increment();
types$4.commaOrSpaceSeparated = increment();

function increment() {
  return Math.pow(2, ++powers)
}

var Info$1 = info;
var types$3 = types$4;

var definedInfo = DefinedInfo$2;

DefinedInfo$2.prototype = new Info$1();
DefinedInfo$2.prototype.defined = true;

var checks = [
  'boolean',
  'booleanish',
  'overloadedBoolean',
  'number',
  'commaSeparated',
  'spaceSeparated',
  'commaOrSpaceSeparated'
];
var checksLength = checks.length;

function DefinedInfo$2(property, attribute, mask, space) {
  var index = -1;
  var check;

  mark$1(this, 'space', space);

  Info$1.call(this, property, attribute);

  while (++index < checksLength) {
    check = checks[index];
    mark$1(this, check, (mask & types$3[check]) === types$3[check]);
  }
}

function mark$1(values, key, value) {
  if (value) {
    values[key] = value;
  }
}

var normalize$1 = normalize_1;
var Schema$6 = schema$1;
var DefinedInfo$1 = definedInfo;

var create_1 = create$6;

function create$6(definition) {
  var space = definition.space;
  var mustUseProperty = definition.mustUseProperty || [];
  var attributes = definition.attributes || {};
  var props = definition.properties;
  var transform = definition.transform;
  var property = {};
  var normal = {};
  var prop;
  var info;

  for (prop in props) {
    info = new DefinedInfo$1(
      prop,
      transform(attributes, prop),
      props[prop],
      space
    );

    if (mustUseProperty.indexOf(prop) !== -1) {
      info.mustUseProperty = true;
    }

    property[prop] = info;

    normal[normalize$1(prop)] = prop;
    normal[normalize$1(info.attribute)] = prop;
  }

  return new Schema$6(property, normal, space)
}

var create$5 = create_1;

var xlink$2 = create$5({
  space: 'xlink',
  transform: xlinkTransform,
  properties: {
    xLinkActuate: null,
    xLinkArcRole: null,
    xLinkHref: null,
    xLinkRole: null,
    xLinkShow: null,
    xLinkTitle: null,
    xLinkType: null
  }
});

function xlinkTransform(_, prop) {
  return 'xlink:' + prop.slice(5).toLowerCase()
}

var create$4 = create_1;

var xml$2 = create$4({
  space: 'xml',
  transform: xmlTransform,
  properties: {
    xmlLang: null,
    xmlBase: null,
    xmlSpace: null
  }
});

function xmlTransform(_, prop) {
  return 'xml:' + prop.slice(3).toLowerCase()
}

var caseSensitiveTransform_1 = caseSensitiveTransform$2;

function caseSensitiveTransform$2(attributes, attribute) {
  return attribute in attributes ? attributes[attribute] : attribute
}

var caseSensitiveTransform$1 = caseSensitiveTransform_1;

var caseInsensitiveTransform_1 = caseInsensitiveTransform$2;

function caseInsensitiveTransform$2(attributes, property) {
  return caseSensitiveTransform$1(attributes, property.toLowerCase())
}

var create$3 = create_1;
var caseInsensitiveTransform$1 = caseInsensitiveTransform_1;

var xmlns$2 = create$3({
  space: 'xmlns',
  attributes: {
    xmlnsxlink: 'xmlns:xlink'
  },
  transform: caseInsensitiveTransform$1,
  properties: {
    xmlns: null,
    xmlnsXLink: null
  }
});

var types$2 = types$4;
var create$2 = create_1;

var booleanish$1 = types$2.booleanish;
var number$2 = types$2.number;
var spaceSeparated$2 = types$2.spaceSeparated;

var aria$2 = create$2({
  transform: ariaTransform,
  properties: {
    ariaActiveDescendant: null,
    ariaAtomic: booleanish$1,
    ariaAutoComplete: null,
    ariaBusy: booleanish$1,
    ariaChecked: booleanish$1,
    ariaColCount: number$2,
    ariaColIndex: number$2,
    ariaColSpan: number$2,
    ariaControls: spaceSeparated$2,
    ariaCurrent: null,
    ariaDescribedBy: spaceSeparated$2,
    ariaDetails: null,
    ariaDisabled: booleanish$1,
    ariaDropEffect: spaceSeparated$2,
    ariaErrorMessage: null,
    ariaExpanded: booleanish$1,
    ariaFlowTo: spaceSeparated$2,
    ariaGrabbed: booleanish$1,
    ariaHasPopup: null,
    ariaHidden: booleanish$1,
    ariaInvalid: null,
    ariaKeyShortcuts: null,
    ariaLabel: null,
    ariaLabelledBy: spaceSeparated$2,
    ariaLevel: number$2,
    ariaLive: null,
    ariaModal: booleanish$1,
    ariaMultiLine: booleanish$1,
    ariaMultiSelectable: booleanish$1,
    ariaOrientation: null,
    ariaOwns: spaceSeparated$2,
    ariaPlaceholder: null,
    ariaPosInSet: number$2,
    ariaPressed: booleanish$1,
    ariaReadOnly: booleanish$1,
    ariaRelevant: null,
    ariaRequired: booleanish$1,
    ariaRoleDescription: spaceSeparated$2,
    ariaRowCount: number$2,
    ariaRowIndex: number$2,
    ariaRowSpan: number$2,
    ariaSelected: booleanish$1,
    ariaSetSize: number$2,
    ariaSort: null,
    ariaValueMax: number$2,
    ariaValueMin: number$2,
    ariaValueNow: number$2,
    ariaValueText: null,
    role: null
  }
});

function ariaTransform(_, prop) {
  return prop === 'role' ? prop : 'aria-' + prop.slice(4).toLowerCase()
}

var types$1 = types$4;
var create$1 = create_1;
var caseInsensitiveTransform = caseInsensitiveTransform_1;

var boolean$1 = types$1.boolean;
var overloadedBoolean = types$1.overloadedBoolean;
var booleanish = types$1.booleanish;
var number$1 = types$1.number;
var spaceSeparated$1 = types$1.spaceSeparated;
var commaSeparated$1 = types$1.commaSeparated;

var html$4 = create$1({
  space: 'html',
  attributes: {
    acceptcharset: 'accept-charset',
    classname: 'class',
    htmlfor: 'for',
    httpequiv: 'http-equiv'
  },
  transform: caseInsensitiveTransform,
  mustUseProperty: ['checked', 'multiple', 'muted', 'selected'],
  properties: {
    // Standard Properties.
    abbr: null,
    accept: commaSeparated$1,
    acceptCharset: spaceSeparated$1,
    accessKey: spaceSeparated$1,
    action: null,
    allow: null,
    allowFullScreen: boolean$1,
    allowPaymentRequest: boolean$1,
    allowUserMedia: boolean$1,
    alt: null,
    as: null,
    async: boolean$1,
    autoCapitalize: null,
    autoComplete: spaceSeparated$1,
    autoFocus: boolean$1,
    autoPlay: boolean$1,
    capture: boolean$1,
    charSet: null,
    checked: boolean$1,
    cite: null,
    className: spaceSeparated$1,
    cols: number$1,
    colSpan: null,
    content: null,
    contentEditable: booleanish,
    controls: boolean$1,
    controlsList: spaceSeparated$1,
    coords: number$1 | commaSeparated$1,
    crossOrigin: null,
    data: null,
    dateTime: null,
    decoding: null,
    default: boolean$1,
    defer: boolean$1,
    dir: null,
    dirName: null,
    disabled: boolean$1,
    download: overloadedBoolean,
    draggable: booleanish,
    encType: null,
    enterKeyHint: null,
    form: null,
    formAction: null,
    formEncType: null,
    formMethod: null,
    formNoValidate: boolean$1,
    formTarget: null,
    headers: spaceSeparated$1,
    height: number$1,
    hidden: boolean$1,
    high: number$1,
    href: null,
    hrefLang: null,
    htmlFor: spaceSeparated$1,
    httpEquiv: spaceSeparated$1,
    id: null,
    imageSizes: null,
    imageSrcSet: commaSeparated$1,
    inputMode: null,
    integrity: null,
    is: null,
    isMap: boolean$1,
    itemId: null,
    itemProp: spaceSeparated$1,
    itemRef: spaceSeparated$1,
    itemScope: boolean$1,
    itemType: spaceSeparated$1,
    kind: null,
    label: null,
    lang: null,
    language: null,
    list: null,
    loop: boolean$1,
    low: number$1,
    manifest: null,
    max: null,
    maxLength: number$1,
    media: null,
    method: null,
    min: null,
    minLength: number$1,
    multiple: boolean$1,
    muted: boolean$1,
    name: null,
    nonce: null,
    noModule: boolean$1,
    noValidate: boolean$1,
    onAbort: null,
    onAfterPrint: null,
    onAuxClick: null,
    onBeforePrint: null,
    onBeforeUnload: null,
    onBlur: null,
    onCancel: null,
    onCanPlay: null,
    onCanPlayThrough: null,
    onChange: null,
    onClick: null,
    onClose: null,
    onContextMenu: null,
    onCopy: null,
    onCueChange: null,
    onCut: null,
    onDblClick: null,
    onDrag: null,
    onDragEnd: null,
    onDragEnter: null,
    onDragExit: null,
    onDragLeave: null,
    onDragOver: null,
    onDragStart: null,
    onDrop: null,
    onDurationChange: null,
    onEmptied: null,
    onEnded: null,
    onError: null,
    onFocus: null,
    onFormData: null,
    onHashChange: null,
    onInput: null,
    onInvalid: null,
    onKeyDown: null,
    onKeyPress: null,
    onKeyUp: null,
    onLanguageChange: null,
    onLoad: null,
    onLoadedData: null,
    onLoadedMetadata: null,
    onLoadEnd: null,
    onLoadStart: null,
    onMessage: null,
    onMessageError: null,
    onMouseDown: null,
    onMouseEnter: null,
    onMouseLeave: null,
    onMouseMove: null,
    onMouseOut: null,
    onMouseOver: null,
    onMouseUp: null,
    onOffline: null,
    onOnline: null,
    onPageHide: null,
    onPageShow: null,
    onPaste: null,
    onPause: null,
    onPlay: null,
    onPlaying: null,
    onPopState: null,
    onProgress: null,
    onRateChange: null,
    onRejectionHandled: null,
    onReset: null,
    onResize: null,
    onScroll: null,
    onSecurityPolicyViolation: null,
    onSeeked: null,
    onSeeking: null,
    onSelect: null,
    onSlotChange: null,
    onStalled: null,
    onStorage: null,
    onSubmit: null,
    onSuspend: null,
    onTimeUpdate: null,
    onToggle: null,
    onUnhandledRejection: null,
    onUnload: null,
    onVolumeChange: null,
    onWaiting: null,
    onWheel: null,
    open: boolean$1,
    optimum: number$1,
    pattern: null,
    ping: spaceSeparated$1,
    placeholder: null,
    playsInline: boolean$1,
    poster: null,
    preload: null,
    readOnly: boolean$1,
    referrerPolicy: null,
    rel: spaceSeparated$1,
    required: boolean$1,
    reversed: boolean$1,
    rows: number$1,
    rowSpan: number$1,
    sandbox: spaceSeparated$1,
    scope: null,
    scoped: boolean$1,
    seamless: boolean$1,
    selected: boolean$1,
    shape: null,
    size: number$1,
    sizes: null,
    slot: null,
    span: number$1,
    spellCheck: booleanish,
    src: null,
    srcDoc: null,
    srcLang: null,
    srcSet: commaSeparated$1,
    start: number$1,
    step: null,
    style: null,
    tabIndex: number$1,
    target: null,
    title: null,
    translate: null,
    type: null,
    typeMustMatch: boolean$1,
    useMap: null,
    value: booleanish,
    width: number$1,
    wrap: null,

    // Legacy.
    // See: https://html.spec.whatwg.org/#other-elements,-attributes-and-apis
    align: null, // Several. Use CSS `text-align` instead,
    aLink: null, // `<body>`. Use CSS `a:active {color}` instead
    archive: spaceSeparated$1, // `<object>`. List of URIs to archives
    axis: null, // `<td>` and `<th>`. Use `scope` on `<th>`
    background: null, // `<body>`. Use CSS `background-image` instead
    bgColor: null, // `<body>` and table elements. Use CSS `background-color` instead
    border: number$1, // `<table>`. Use CSS `border-width` instead,
    borderColor: null, // `<table>`. Use CSS `border-color` instead,
    bottomMargin: number$1, // `<body>`
    cellPadding: null, // `<table>`
    cellSpacing: null, // `<table>`
    char: null, // Several table elements. When `align=char`, sets the character to align on
    charOff: null, // Several table elements. When `char`, offsets the alignment
    classId: null, // `<object>`
    clear: null, // `<br>`. Use CSS `clear` instead
    code: null, // `<object>`
    codeBase: null, // `<object>`
    codeType: null, // `<object>`
    color: null, // `<font>` and `<hr>`. Use CSS instead
    compact: boolean$1, // Lists. Use CSS to reduce space between items instead
    declare: boolean$1, // `<object>`
    event: null, // `<script>`
    face: null, // `<font>`. Use CSS instead
    frame: null, // `<table>`
    frameBorder: null, // `<iframe>`. Use CSS `border` instead
    hSpace: number$1, // `<img>` and `<object>`
    leftMargin: number$1, // `<body>`
    link: null, // `<body>`. Use CSS `a:link {color: *}` instead
    longDesc: null, // `<frame>`, `<iframe>`, and `<img>`. Use an `<a>`
    lowSrc: null, // `<img>`. Use a `<picture>`
    marginHeight: number$1, // `<body>`
    marginWidth: number$1, // `<body>`
    noResize: boolean$1, // `<frame>`
    noHref: boolean$1, // `<area>`. Use no href instead of an explicit `nohref`
    noShade: boolean$1, // `<hr>`. Use background-color and height instead of borders
    noWrap: boolean$1, // `<td>` and `<th>`
    object: null, // `<applet>`
    profile: null, // `<head>`
    prompt: null, // `<isindex>`
    rev: null, // `<link>`
    rightMargin: number$1, // `<body>`
    rules: null, // `<table>`
    scheme: null, // `<meta>`
    scrolling: booleanish, // `<frame>`. Use overflow in the child context
    standby: null, // `<object>`
    summary: null, // `<table>`
    text: null, // `<body>`. Use CSS `color` instead
    topMargin: number$1, // `<body>`
    valueType: null, // `<param>`
    version: null, // `<html>`. Use a doctype.
    vAlign: null, // Several. Use CSS `vertical-align` instead
    vLink: null, // `<body>`. Use CSS `a:visited {color}` instead
    vSpace: number$1, // `<img>` and `<object>`

    // Non-standard Properties.
    allowTransparency: null,
    autoCorrect: null,
    autoSave: null,
    disablePictureInPicture: boolean$1,
    disableRemotePlayback: boolean$1,
    prefix: null,
    property: null,
    results: number$1,
    security: null,
    unselectable: null
  }
});

var merge$2 = merge_1;
var xlink$1 = xlink$2;
var xml$1 = xml$2;
var xmlns$1 = xmlns$2;
var aria$1 = aria$2;
var html$3 = html$4;

var html_1 = merge$2([xml$1, xlink$1, xmlns$1, aria$1, html$3]);

var types = types$4;
var create = create_1;
var caseSensitiveTransform = caseSensitiveTransform_1;

var boolean = types.boolean;
var number = types.number;
var spaceSeparated = types.spaceSeparated;
var commaSeparated = types.commaSeparated;
var commaOrSpaceSeparated = types.commaOrSpaceSeparated;

var svg$2 = create({
  space: 'svg',
  attributes: {
    accentHeight: 'accent-height',
    alignmentBaseline: 'alignment-baseline',
    arabicForm: 'arabic-form',
    baselineShift: 'baseline-shift',
    capHeight: 'cap-height',
    className: 'class',
    clipPath: 'clip-path',
    clipRule: 'clip-rule',
    colorInterpolation: 'color-interpolation',
    colorInterpolationFilters: 'color-interpolation-filters',
    colorProfile: 'color-profile',
    colorRendering: 'color-rendering',
    crossOrigin: 'crossorigin',
    dataType: 'datatype',
    dominantBaseline: 'dominant-baseline',
    enableBackground: 'enable-background',
    fillOpacity: 'fill-opacity',
    fillRule: 'fill-rule',
    floodColor: 'flood-color',
    floodOpacity: 'flood-opacity',
    fontFamily: 'font-family',
    fontSize: 'font-size',
    fontSizeAdjust: 'font-size-adjust',
    fontStretch: 'font-stretch',
    fontStyle: 'font-style',
    fontVariant: 'font-variant',
    fontWeight: 'font-weight',
    glyphName: 'glyph-name',
    glyphOrientationHorizontal: 'glyph-orientation-horizontal',
    glyphOrientationVertical: 'glyph-orientation-vertical',
    hrefLang: 'hreflang',
    horizAdvX: 'horiz-adv-x',
    horizOriginX: 'horiz-origin-x',
    horizOriginY: 'horiz-origin-y',
    imageRendering: 'image-rendering',
    letterSpacing: 'letter-spacing',
    lightingColor: 'lighting-color',
    markerEnd: 'marker-end',
    markerMid: 'marker-mid',
    markerStart: 'marker-start',
    navDown: 'nav-down',
    navDownLeft: 'nav-down-left',
    navDownRight: 'nav-down-right',
    navLeft: 'nav-left',
    navNext: 'nav-next',
    navPrev: 'nav-prev',
    navRight: 'nav-right',
    navUp: 'nav-up',
    navUpLeft: 'nav-up-left',
    navUpRight: 'nav-up-right',
    onAbort: 'onabort',
    onActivate: 'onactivate',
    onAfterPrint: 'onafterprint',
    onBeforePrint: 'onbeforeprint',
    onBegin: 'onbegin',
    onCancel: 'oncancel',
    onCanPlay: 'oncanplay',
    onCanPlayThrough: 'oncanplaythrough',
    onChange: 'onchange',
    onClick: 'onclick',
    onClose: 'onclose',
    onCopy: 'oncopy',
    onCueChange: 'oncuechange',
    onCut: 'oncut',
    onDblClick: 'ondblclick',
    onDrag: 'ondrag',
    onDragEnd: 'ondragend',
    onDragEnter: 'ondragenter',
    onDragExit: 'ondragexit',
    onDragLeave: 'ondragleave',
    onDragOver: 'ondragover',
    onDragStart: 'ondragstart',
    onDrop: 'ondrop',
    onDurationChange: 'ondurationchange',
    onEmptied: 'onemptied',
    onEnd: 'onend',
    onEnded: 'onended',
    onError: 'onerror',
    onFocus: 'onfocus',
    onFocusIn: 'onfocusin',
    onFocusOut: 'onfocusout',
    onHashChange: 'onhashchange',
    onInput: 'oninput',
    onInvalid: 'oninvalid',
    onKeyDown: 'onkeydown',
    onKeyPress: 'onkeypress',
    onKeyUp: 'onkeyup',
    onLoad: 'onload',
    onLoadedData: 'onloadeddata',
    onLoadedMetadata: 'onloadedmetadata',
    onLoadStart: 'onloadstart',
    onMessage: 'onmessage',
    onMouseDown: 'onmousedown',
    onMouseEnter: 'onmouseenter',
    onMouseLeave: 'onmouseleave',
    onMouseMove: 'onmousemove',
    onMouseOut: 'onmouseout',
    onMouseOver: 'onmouseover',
    onMouseUp: 'onmouseup',
    onMouseWheel: 'onmousewheel',
    onOffline: 'onoffline',
    onOnline: 'ononline',
    onPageHide: 'onpagehide',
    onPageShow: 'onpageshow',
    onPaste: 'onpaste',
    onPause: 'onpause',
    onPlay: 'onplay',
    onPlaying: 'onplaying',
    onPopState: 'onpopstate',
    onProgress: 'onprogress',
    onRateChange: 'onratechange',
    onRepeat: 'onrepeat',
    onReset: 'onreset',
    onResize: 'onresize',
    onScroll: 'onscroll',
    onSeeked: 'onseeked',
    onSeeking: 'onseeking',
    onSelect: 'onselect',
    onShow: 'onshow',
    onStalled: 'onstalled',
    onStorage: 'onstorage',
    onSubmit: 'onsubmit',
    onSuspend: 'onsuspend',
    onTimeUpdate: 'ontimeupdate',
    onToggle: 'ontoggle',
    onUnload: 'onunload',
    onVolumeChange: 'onvolumechange',
    onWaiting: 'onwaiting',
    onZoom: 'onzoom',
    overlinePosition: 'overline-position',
    overlineThickness: 'overline-thickness',
    paintOrder: 'paint-order',
    panose1: 'panose-1',
    pointerEvents: 'pointer-events',
    referrerPolicy: 'referrerpolicy',
    renderingIntent: 'rendering-intent',
    shapeRendering: 'shape-rendering',
    stopColor: 'stop-color',
    stopOpacity: 'stop-opacity',
    strikethroughPosition: 'strikethrough-position',
    strikethroughThickness: 'strikethrough-thickness',
    strokeDashArray: 'stroke-dasharray',
    strokeDashOffset: 'stroke-dashoffset',
    strokeLineCap: 'stroke-linecap',
    strokeLineJoin: 'stroke-linejoin',
    strokeMiterLimit: 'stroke-miterlimit',
    strokeOpacity: 'stroke-opacity',
    strokeWidth: 'stroke-width',
    tabIndex: 'tabindex',
    textAnchor: 'text-anchor',
    textDecoration: 'text-decoration',
    textRendering: 'text-rendering',
    typeOf: 'typeof',
    underlinePosition: 'underline-position',
    underlineThickness: 'underline-thickness',
    unicodeBidi: 'unicode-bidi',
    unicodeRange: 'unicode-range',
    unitsPerEm: 'units-per-em',
    vAlphabetic: 'v-alphabetic',
    vHanging: 'v-hanging',
    vIdeographic: 'v-ideographic',
    vMathematical: 'v-mathematical',
    vectorEffect: 'vector-effect',
    vertAdvY: 'vert-adv-y',
    vertOriginX: 'vert-origin-x',
    vertOriginY: 'vert-origin-y',
    wordSpacing: 'word-spacing',
    writingMode: 'writing-mode',
    xHeight: 'x-height',
    // These were camelcased in Tiny. Now lowercased in SVG 2
    playbackOrder: 'playbackorder',
    timelineBegin: 'timelinebegin'
  },
  transform: caseSensitiveTransform,
  properties: {
    about: commaOrSpaceSeparated,
    accentHeight: number,
    accumulate: null,
    additive: null,
    alignmentBaseline: null,
    alphabetic: number,
    amplitude: number,
    arabicForm: null,
    ascent: number,
    attributeName: null,
    attributeType: null,
    azimuth: number,
    bandwidth: null,
    baselineShift: null,
    baseFrequency: null,
    baseProfile: null,
    bbox: null,
    begin: null,
    bias: number,
    by: null,
    calcMode: null,
    capHeight: number,
    className: spaceSeparated,
    clip: null,
    clipPath: null,
    clipPathUnits: null,
    clipRule: null,
    color: null,
    colorInterpolation: null,
    colorInterpolationFilters: null,
    colorProfile: null,
    colorRendering: null,
    content: null,
    contentScriptType: null,
    contentStyleType: null,
    crossOrigin: null,
    cursor: null,
    cx: null,
    cy: null,
    d: null,
    dataType: null,
    defaultAction: null,
    descent: number,
    diffuseConstant: number,
    direction: null,
    display: null,
    dur: null,
    divisor: number,
    dominantBaseline: null,
    download: boolean,
    dx: null,
    dy: null,
    edgeMode: null,
    editable: null,
    elevation: number,
    enableBackground: null,
    end: null,
    event: null,
    exponent: number,
    externalResourcesRequired: null,
    fill: null,
    fillOpacity: number,
    fillRule: null,
    filter: null,
    filterRes: null,
    filterUnits: null,
    floodColor: null,
    floodOpacity: null,
    focusable: null,
    focusHighlight: null,
    fontFamily: null,
    fontSize: null,
    fontSizeAdjust: null,
    fontStretch: null,
    fontStyle: null,
    fontVariant: null,
    fontWeight: null,
    format: null,
    fr: null,
    from: null,
    fx: null,
    fy: null,
    g1: commaSeparated,
    g2: commaSeparated,
    glyphName: commaSeparated,
    glyphOrientationHorizontal: null,
    glyphOrientationVertical: null,
    glyphRef: null,
    gradientTransform: null,
    gradientUnits: null,
    handler: null,
    hanging: number,
    hatchContentUnits: null,
    hatchUnits: null,
    height: null,
    href: null,
    hrefLang: null,
    horizAdvX: number,
    horizOriginX: number,
    horizOriginY: number,
    id: null,
    ideographic: number,
    imageRendering: null,
    initialVisibility: null,
    in: null,
    in2: null,
    intercept: number,
    k: number,
    k1: number,
    k2: number,
    k3: number,
    k4: number,
    kernelMatrix: commaOrSpaceSeparated,
    kernelUnitLength: null,
    keyPoints: null, // SEMI_COLON_SEPARATED
    keySplines: null, // SEMI_COLON_SEPARATED
    keyTimes: null, // SEMI_COLON_SEPARATED
    kerning: null,
    lang: null,
    lengthAdjust: null,
    letterSpacing: null,
    lightingColor: null,
    limitingConeAngle: number,
    local: null,
    markerEnd: null,
    markerMid: null,
    markerStart: null,
    markerHeight: null,
    markerUnits: null,
    markerWidth: null,
    mask: null,
    maskContentUnits: null,
    maskUnits: null,
    mathematical: null,
    max: null,
    media: null,
    mediaCharacterEncoding: null,
    mediaContentEncodings: null,
    mediaSize: number,
    mediaTime: null,
    method: null,
    min: null,
    mode: null,
    name: null,
    navDown: null,
    navDownLeft: null,
    navDownRight: null,
    navLeft: null,
    navNext: null,
    navPrev: null,
    navRight: null,
    navUp: null,
    navUpLeft: null,
    navUpRight: null,
    numOctaves: null,
    observer: null,
    offset: null,
    onAbort: null,
    onActivate: null,
    onAfterPrint: null,
    onBeforePrint: null,
    onBegin: null,
    onCancel: null,
    onCanPlay: null,
    onCanPlayThrough: null,
    onChange: null,
    onClick: null,
    onClose: null,
    onCopy: null,
    onCueChange: null,
    onCut: null,
    onDblClick: null,
    onDrag: null,
    onDragEnd: null,
    onDragEnter: null,
    onDragExit: null,
    onDragLeave: null,
    onDragOver: null,
    onDragStart: null,
    onDrop: null,
    onDurationChange: null,
    onEmptied: null,
    onEnd: null,
    onEnded: null,
    onError: null,
    onFocus: null,
    onFocusIn: null,
    onFocusOut: null,
    onHashChange: null,
    onInput: null,
    onInvalid: null,
    onKeyDown: null,
    onKeyPress: null,
    onKeyUp: null,
    onLoad: null,
    onLoadedData: null,
    onLoadedMetadata: null,
    onLoadStart: null,
    onMessage: null,
    onMouseDown: null,
    onMouseEnter: null,
    onMouseLeave: null,
    onMouseMove: null,
    onMouseOut: null,
    onMouseOver: null,
    onMouseUp: null,
    onMouseWheel: null,
    onOffline: null,
    onOnline: null,
    onPageHide: null,
    onPageShow: null,
    onPaste: null,
    onPause: null,
    onPlay: null,
    onPlaying: null,
    onPopState: null,
    onProgress: null,
    onRateChange: null,
    onRepeat: null,
    onReset: null,
    onResize: null,
    onScroll: null,
    onSeeked: null,
    onSeeking: null,
    onSelect: null,
    onShow: null,
    onStalled: null,
    onStorage: null,
    onSubmit: null,
    onSuspend: null,
    onTimeUpdate: null,
    onToggle: null,
    onUnload: null,
    onVolumeChange: null,
    onWaiting: null,
    onZoom: null,
    opacity: null,
    operator: null,
    order: null,
    orient: null,
    orientation: null,
    origin: null,
    overflow: null,
    overlay: null,
    overlinePosition: number,
    overlineThickness: number,
    paintOrder: null,
    panose1: null,
    path: null,
    pathLength: number,
    patternContentUnits: null,
    patternTransform: null,
    patternUnits: null,
    phase: null,
    ping: spaceSeparated,
    pitch: null,
    playbackOrder: null,
    pointerEvents: null,
    points: null,
    pointsAtX: number,
    pointsAtY: number,
    pointsAtZ: number,
    preserveAlpha: null,
    preserveAspectRatio: null,
    primitiveUnits: null,
    propagate: null,
    property: commaOrSpaceSeparated,
    r: null,
    radius: null,
    referrerPolicy: null,
    refX: null,
    refY: null,
    rel: commaOrSpaceSeparated,
    rev: commaOrSpaceSeparated,
    renderingIntent: null,
    repeatCount: null,
    repeatDur: null,
    requiredExtensions: commaOrSpaceSeparated,
    requiredFeatures: commaOrSpaceSeparated,
    requiredFonts: commaOrSpaceSeparated,
    requiredFormats: commaOrSpaceSeparated,
    resource: null,
    restart: null,
    result: null,
    rotate: null,
    rx: null,
    ry: null,
    scale: null,
    seed: null,
    shapeRendering: null,
    side: null,
    slope: null,
    snapshotTime: null,
    specularConstant: number,
    specularExponent: number,
    spreadMethod: null,
    spacing: null,
    startOffset: null,
    stdDeviation: null,
    stemh: null,
    stemv: null,
    stitchTiles: null,
    stopColor: null,
    stopOpacity: null,
    strikethroughPosition: number,
    strikethroughThickness: number,
    string: null,
    stroke: null,
    strokeDashArray: commaOrSpaceSeparated,
    strokeDashOffset: null,
    strokeLineCap: null,
    strokeLineJoin: null,
    strokeMiterLimit: number,
    strokeOpacity: number,
    strokeWidth: null,
    style: null,
    surfaceScale: number,
    syncBehavior: null,
    syncBehaviorDefault: null,
    syncMaster: null,
    syncTolerance: null,
    syncToleranceDefault: null,
    systemLanguage: commaOrSpaceSeparated,
    tabIndex: number,
    tableValues: null,
    target: null,
    targetX: number,
    targetY: number,
    textAnchor: null,
    textDecoration: null,
    textRendering: null,
    textLength: null,
    timelineBegin: null,
    title: null,
    transformBehavior: null,
    type: null,
    typeOf: commaOrSpaceSeparated,
    to: null,
    transform: null,
    u1: null,
    u2: null,
    underlinePosition: number,
    underlineThickness: number,
    unicode: null,
    unicodeBidi: null,
    unicodeRange: null,
    unitsPerEm: number,
    values: null,
    vAlphabetic: number,
    vMathematical: number,
    vectorEffect: null,
    vHanging: number,
    vIdeographic: number,
    version: null,
    vertAdvY: number,
    vertOriginX: number,
    vertOriginY: number,
    viewBox: null,
    viewTarget: null,
    visibility: null,
    width: null,
    widths: null,
    wordSpacing: null,
    writingMode: null,
    x: null,
    x1: null,
    x2: null,
    xChannelSelector: null,
    xHeight: number,
    y: null,
    y1: null,
    y2: null,
    yChannelSelector: null,
    z: null,
    zoomAndPan: null
  }
});

var merge$1 = merge_1;
var xlink = xlink$2;
var xml = xml$2;
var xmlns = xmlns$2;
var aria = aria$2;
var svg$1 = svg$2;

var svg_1 = merge$1([xml, xlink, xmlns, aria, svg$1]);

var require$$2 = [
	"area",
	"base",
	"basefont",
	"bgsound",
	"br",
	"col",
	"command",
	"embed",
	"frame",
	"hr",
	"image",
	"img",
	"input",
	"isindex",
	"keygen",
	"link",
	"menuitem",
	"meta",
	"nextid",
	"param",
	"source",
	"track",
	"wbr"
];

var omission$4 = {};

var unistUtilIs = is$3;

// Assert if `test` passes for `node`.   When a `parent` node is known the
// `index` of node.
// eslint-disable-next-line max-params
function is$3(test, node, index, parent, context) {
  var hasParent = parent !== null && parent !== undefined;
  var hasIndex = index !== null && index !== undefined;
  var check = convert(test);

  if (
    hasIndex &&
    (typeof index !== 'number' || index < 0 || index === Infinity)
  ) {
    throw new Error('Expected positive finite index or child node')
  }

  if (hasParent && (!is$3(null, parent) || !parent.children)) {
    throw new Error('Expected parent node')
  }

  if (!node || !node.type || typeof node.type !== 'string') {
    return false
  }

  if (hasParent !== hasIndex) {
    throw new Error('Expected both parent and index')
  }

  return Boolean(check.call(context, node, index, parent))
}

function convert(test) {
  if (typeof test === 'string') {
    return typeFactory(test)
  }

  if (test === null || test === undefined) {
    return ok
  }

  if (typeof test === 'object') {
    return ('length' in test ? anyFactory : matchesFactory)(test)
  }

  if (typeof test === 'function') {
    return test
  }

  throw new Error('Expected function, string, or object as test')
}

function convertAll(tests) {
  var results = [];
  var length = tests.length;
  var index = -1;

  while (++index < length) {
    results[index] = convert(tests[index]);
  }

  return results
}

// Utility assert each property in `test` is represented in `node`, and each
// values are strictly equal.
function matchesFactory(test) {
  return matches

  function matches(node) {
    var key;

    for (key in test) {
      if (node[key] !== test[key]) {
        return false
      }
    }

    return true
  }
}

function anyFactory(tests) {
  var checks = convertAll(tests);
  var length = checks.length;

  return matches

  function matches() {
    var index = -1;

    while (++index < length) {
      if (checks[index].apply(this, arguments)) {
        return true
      }
    }

    return false
  }
}

// Utility to convert a string into a function which checks a given node’s type
// for said string.
function typeFactory(test) {
  return type

  function type(node) {
    return Boolean(node && node.type === test)
  }
}

// Utility to return true.
function ok() {
  return true
}

var hastUtilIsElement = isElement;

// Check if if `node` is an `element` and, if `tagNames` is given, `node`
// matches them `tagNames`.
function isElement(node, tagNames) {
  var name;

  if (
    !(
      tagNames === null ||
      tagNames === undefined ||
      typeof tagNames === 'string' ||
      (typeof tagNames === 'object' && tagNames.length !== 0)
    )
  ) {
    throw new Error(
      'Expected `string` or `Array.<string>` for `tagNames`, not `' +
        tagNames +
        '`'
    )
  }

  if (
    !node ||
    typeof node !== 'object' ||
    node.type !== 'element' ||
    typeof node.tagName !== 'string'
  ) {
    return false
  }

  if (tagNames === null || tagNames === undefined) {
    return true
  }

  name = node.tagName;

  if (typeof tagNames === 'string') {
    return name === tagNames
  }

  return tagNames.indexOf(name) !== -1
}

var siblings$1 = {};

var hastUtilWhitespace = interElementWhiteSpace;

// HTML white-space expression.
// See <https://html.spec.whatwg.org/#space-character>.
var re = /[ \t\n\f\r]/g;

function interElementWhiteSpace(node) {
  var value;

  if (node && typeof node === 'object' && node.type === 'text') {
    value = node.value || '';
  } else if (typeof node === 'string') {
    value = node;
  } else {
    return false
  }

  return value.replace(re, '') === ''
}

var whiteSpace$2 = hastUtilWhitespace;

siblings$1.before = siblings(-1);
siblings$1.after = siblings(1);

/* Factory to check siblings in a direction. */
function siblings(increment) {
  return sibling

  /* Find applicable siblings in a direction.   */
  function sibling(parent, index, includeWhiteSpace) {
    var siblings = parent && parent.children;
    var next;

    index += increment;
    next = siblings && siblings[index];

    if (!includeWhiteSpace) {
      while (next && whiteSpace$2(next)) {
        index += increment;
        next = siblings[index];
      }
    }

    return next
  }
}

var after$1 = siblings$1.after;

var first_1 = first$1;

/* Get the first child in `parent`. */
function first$1(parent, includeWhiteSpace) {
  return after$1(parent, -1, includeWhiteSpace)
}

var place_1 = place$1;

/* Get the position of `node` in `parent`. */
function place$1(parent, child) {
  return parent && parent.children && parent.children.indexOf(child)
}

var is$2 = unistUtilIs;
var whiteSpace$1 = hastUtilWhitespace;

var whiteSpaceLeft_1 = whiteSpaceLeft$2;

/* Check if `node` starts with white-space. */
function whiteSpaceLeft$2(node) {
  return is$2('text', node) && whiteSpace$1(node.value.charAt(0))
}

var omission_1 = omission$3;

var own$3 = {}.hasOwnProperty;

/* Factory to check if a given node can have a tag omitted. */
function omission$3(handlers) {
  return omit

  /* Check if a given node can have a tag omitted.   */
  function omit(node, index, parent) {
    var name = node.tagName;
    var fn = own$3.call(handlers, name) ? handlers[name] : false;

    return fn ? fn(node, index, parent) : false
  }
}

var is$1 = unistUtilIs;
var element$1 = hastUtilIsElement;
var whiteSpaceLeft$1 = whiteSpaceLeft_1;
var after = siblings$1.after;
var omission$2 = omission_1;

var optionGroup = 'optgroup';
var options = ['option'].concat(optionGroup);
var dataListItem = ['dt', 'dd'];
var listItem = 'li';
var menuContent = ['menuitem', 'hr', 'menu'];
var ruby = ['rp', 'rt'];
var tableContainer = ['tbody', 'tfoot'];
var tableRow$1 = 'tr';
var tableCell = ['td', 'th'];

var confusingParagraphParent = ['a', 'audio', 'del', 'ins', 'map', 'noscript', 'video'];

var clearParagraphSibling = [
  'address',
  'article',
  'aside',
  'blockquote',
  'details',
  'div',
  'dl',
  'fieldset',
  'figcaption',
  'figure',
  'footer',
  'form',
  'h1',
  'h2',
  'h3',
  'h4',
  'h5',
  'h6',
  'header',
  'hgroup',
  'hr',
  'main',
  'menu',
  'nav',
  'ol',
  'p',
  'pre',
  'section',
  'table',
  'ul'
];

var closing$1 = omission$2({
  html: html$2,
  head: headOrColgroupOrCaption,
  body: body$1,
  p: p,
  li: li,
  dt: dt,
  dd: dd,
  rt: rubyElement,
  rp: rubyElement,
  optgroup: optgroup,
  option: option,
  menuitem: menuitem,
  colgroup: headOrColgroupOrCaption,
  caption: headOrColgroupOrCaption,
  thead: thead,
  tbody: tbody$1,
  tfoot: tfoot,
  tr: tr,
  td: cells,
  th: cells
});

/* Macro for `</head>`, `</colgroup>`, and `</caption>`. */
function headOrColgroupOrCaption(node, index, parent) {
  var next = after(parent, index, true);
  return !next || (!is$1('comment', next) && !whiteSpaceLeft$1(next))
}

/* Whether to omit `</html>`. */
function html$2(node, index, parent) {
  var next = after(parent, index);
  return !next || !is$1('comment', next)
}

/* Whether to omit `</body>`. */
function body$1(node, index, parent) {
  var next = after(parent, index);
  return !next || !is$1('comment', next)
}

/* Whether to omit `</p>`. */
function p(node, index, parent) {
  var next = after(parent, index);
  return next ? element$1(next, clearParagraphSibling) : !parent || !element$1(parent, confusingParagraphParent)
}

/* Whether to omit `</li>`. */
function li(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, listItem)
}

/* Whether to omit `</dt>`. */
function dt(node, index, parent) {
  var next = after(parent, index);
  return next && element$1(next, dataListItem)
}

/* Whether to omit `</dd>`. */
function dd(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, dataListItem)
}

/* Whether to omit `</rt>` or `</rp>`. */
function rubyElement(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, ruby)
}

/* Whether to omit `</optgroup>`. */
function optgroup(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, optionGroup)
}

/* Whether to omit `</option>`. */
function option(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, options)
}

/* Whether to omit `</menuitem>`. */
function menuitem(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, menuContent)
}

/* Whether to omit `</thead>`. */
function thead(node, index, parent) {
  var next = after(parent, index);
  return next && element$1(next, tableContainer)
}

/* Whether to omit `</tbody>`. */
function tbody$1(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, tableContainer)
}

/* Whether to omit `</tfoot>`. */
function tfoot(node, index, parent) {
  return !after(parent, index)
}

/* Whether to omit `</tr>`. */
function tr(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, tableRow$1)
}

/* Whether to omit `</td>` or `</th>`. */
function cells(node, index, parent) {
  var next = after(parent, index);
  return !next || element$1(next, tableCell)
}

var is = unistUtilIs;
var element = hastUtilIsElement;
var before = siblings$1.before;
var first = first_1;
var place = place_1;
var whiteSpaceLeft = whiteSpaceLeft_1;
var closing = closing$1;
var omission$1 = omission_1;

var own$2 = {}.hasOwnProperty;

var uniqueHeadMetadata = ['title', 'base'];
var meta = ['meta', 'link', 'script', 'style', 'template'];
var tableContainers = ['thead', 'tbody'];
var tableRow = 'tr';

var opening = omission$1({
  html: html$1,
  head: head,
  body: body,
  colgroup: colgroup,
  tbody: tbody
});

/* Whether to omit `<html>`. */
function html$1(node) {
  var head = first(node);
  return !head || !is('comment', head)
}

/* Whether to omit `<head>`. */
function head(node) {
  var children = node.children;
  var length = children.length;
  var map = {};
  var index = -1;
  var child;
  var name;

  while (++index < length) {
    child = children[index];
    name = child.tagName;

    if (element(child, uniqueHeadMetadata)) {
      if (own$2.call(map, name)) {
        return false
      }

      map[name] = true;
    }
  }

  return Boolean(length)
}

/* Whether to omit `<body>`. */
function body(node) {
  var head = first(node, true);

  return !head || (!is('comment', head) && !whiteSpaceLeft(head) && !element(head, meta))
}

/* Whether to omit `<colgroup>`.
 * The spec describes some logic for the opening tag,
 * but it’s easier to implement in the closing tag, to
 * the same effect, so we handle it there instead. */
function colgroup(node, index, parent) {
  var prev = before(parent, index);
  var head = first(node, true);

  /* Previous colgroup was already omitted. */
  if (element(prev, 'colgroup') && closing(prev, place(parent, prev), parent)) {
    return false
  }

  return head && element(head, 'col')
}

/* Whether to omit `<tbody>`. */
function tbody(node, index, parent) {
  var prev = before(parent, index);
  var head = first(node);

  /* Previous table section was already omitted. */
  if (element(prev, tableContainers) && closing(prev, place(parent, prev), parent)) {
    return false
  }

  return head && element(head, tableRow)
}

omission$4.opening = opening;
omission$4.closing = closing$1;

var require$$1 = [
	"script",
	"style",
	"pre",
	"textarea"
];

var all_1;
var hasRequiredAll;

function requireAll () {
	if (hasRequiredAll) return all_1;
	hasRequiredAll = 1;

	var one = requireOne();
	var sensitive = require$$1;

	all_1 = all;

	/* Stringify all children of `parent`. */
	function all(ctx, parent) {
	  var children = parent && parent.children;
	  var length = children && children.length;
	  var index = -1;
	  var results = [];

	  let printWidthOffset = 0;
	  let innerTextLength = 0;
	  while (++index < length) {
	    innerTextLength = getInnerTextLength(children[index]);
	    results[index] = one(ctx, children[index], index, parent, printWidthOffset, innerTextLength);
	    printWidthOffset = results[index].replace(/\n+/g, '').length;
	  }

	  return results.join('')
	}

	/**
	 * Returns the text lenght of the first line of the first child.
	 * Whitespace sensitive elements are ignored.
	 * @param {*} node
	 */
	function getInnerTextLength(node) {
	  // ignore style, script, pre, textarea elements
	  if (sensitive.indexOf(node.tagName) !== -1) {
	    return 0
	  }

	  if (!node.children || !node.children.length) {
	    return 0
	  }

	  var child = node.children[0];

	  if (child.type === 'text' || child.type === 'comment') {
	    return child.value.split('\n')[0].length
	  }

	  return 0
	}
	return all_1;
}

var text_1 = text;

/* Stringify `text`. */
function text(ctx, node, index, parent) {
  var value = node.value;

  return value
}

var normalize = normalize_1;
var DefinedInfo = definedInfo;
var Info = info;

var data = 'data';

var find_1 = find;

var valid = /^data[-\w.:]+$/i;
var dash = /-[a-z]/g;
var cap$1 = /[A-Z]/g;

function find(schema, value) {
  var normal = normalize(value);
  var prop = value;
  var Type = Info;

  if (normal in schema.normal) {
    return schema.property[schema.normal[normal]]
  }

  if (normal.length > 4 && normal.slice(0, 4) === data && valid.test(value)) {
    // Attribute or property.
    if (value.charAt(4) === '-') {
      prop = datasetToProperty(value);
    } else {
      value = datasetToAttribute(value);
    }

    Type = DefinedInfo;
  }

  return new Type(prop, value)
}

function datasetToProperty(attribute) {
  var value = attribute.slice(5).replace(dash, camelcase);
  return data + value.charAt(0).toUpperCase() + value.slice(1)
}

function datasetToAttribute(property) {
  var value = property.slice(4);

  if (dash.test(value)) {
    return property
  }

  value = value.replace(cap$1, kebab);

  if (value.charAt(0) !== '-') {
    value = '-' + value;
  }

  return data + value
}

function kebab($0) {
  return '-' + $0.toLowerCase()
}

function camelcase($0) {
  return $0.charAt(1).toUpperCase()
}

var commaSeparatedTokens = {};

commaSeparatedTokens.parse = parse$1;
commaSeparatedTokens.stringify = stringify$3;

var comma = ',';
var space$1 = ' ';
var empty$1 = '';

// Parse comma-separated tokens to an array.
function parse$1(value) {
  var values = [];
  var input = String(value || empty$1);
  var index = input.indexOf(comma);
  var lastIndex = 0;
  var end = false;
  var val;

  while (!end) {
    if (index === -1) {
      index = input.length;
      end = true;
    }

    val = input.slice(lastIndex, index).trim();

    if (val || !end) {
      values.push(val);
    }

    lastIndex = index + 1;
    index = input.indexOf(comma, lastIndex);
  }

  return values
}

// Compile an array to comma-separated tokens.
// `options.padLeft` (default: `true`) pads a space left of each token, and
// `options.padRight` (default: `false`) pads a space to the right of each token.
function stringify$3(values, options) {
  var settings = options || {};
  var left = settings.padLeft === false ? empty$1 : space$1;
  var right = settings.padRight ? space$1 : empty$1;

  // Ensure the last empty entry is seen.
  if (values[values.length - 1] === empty$1) {
    values = values.concat(empty$1);
  }

  return values.join(right + comma + left).trim()
}

var nbsp = " ";
var iexcl = "¡";
var cent = "¢";
var pound = "£";
var curren = "¤";
var yen = "¥";
var brvbar = "¦";
var sect = "§";
var uml = "¨";
var copy = "©";
var ordf = "ª";
var laquo = "«";
var not = "¬";
var shy = "­";
var reg = "®";
var macr = "¯";
var deg = "°";
var plusmn = "±";
var sup2 = "²";
var sup3 = "³";
var acute = "´";
var micro = "µ";
var para = "¶";
var middot = "·";
var cedil = "¸";
var sup1 = "¹";
var ordm = "º";
var raquo = "»";
var frac14 = "¼";
var frac12 = "½";
var frac34 = "¾";
var iquest = "¿";
var Agrave = "À";
var Aacute = "Á";
var Acirc = "Â";
var Atilde = "Ã";
var Auml = "Ä";
var Aring = "Å";
var AElig = "Æ";
var Ccedil = "Ç";
var Egrave = "È";
var Eacute = "É";
var Ecirc = "Ê";
var Euml = "Ë";
var Igrave = "Ì";
var Iacute = "Í";
var Icirc = "Î";
var Iuml = "Ï";
var ETH = "Ð";
var Ntilde = "Ñ";
var Ograve = "Ò";
var Oacute = "Ó";
var Ocirc = "Ô";
var Otilde = "Õ";
var Ouml = "Ö";
var times = "×";
var Oslash = "Ø";
var Ugrave = "Ù";
var Uacute = "Ú";
var Ucirc = "Û";
var Uuml = "Ü";
var Yacute = "Ý";
var THORN = "Þ";
var szlig = "ß";
var agrave = "à";
var aacute = "á";
var acirc = "â";
var atilde = "ã";
var auml = "ä";
var aring = "å";
var aelig = "æ";
var ccedil = "ç";
var egrave = "è";
var eacute = "é";
var ecirc = "ê";
var euml = "ë";
var igrave = "ì";
var iacute = "í";
var icirc = "î";
var iuml = "ï";
var eth = "ð";
var ntilde = "ñ";
var ograve = "ò";
var oacute = "ó";
var ocirc = "ô";
var otilde = "õ";
var ouml = "ö";
var divide = "÷";
var oslash = "ø";
var ugrave = "ù";
var uacute = "ú";
var ucirc = "û";
var uuml = "ü";
var yacute = "ý";
var thorn = "þ";
var yuml = "ÿ";
var fnof = "ƒ";
var Alpha = "Α";
var Beta = "Β";
var Gamma = "Γ";
var Delta = "Δ";
var Epsilon = "Ε";
var Zeta = "Ζ";
var Eta = "Η";
var Theta = "Θ";
var Iota = "Ι";
var Kappa = "Κ";
var Lambda = "Λ";
var Mu = "Μ";
var Nu = "Ν";
var Xi = "Ξ";
var Omicron = "Ο";
var Pi = "Π";
var Rho = "Ρ";
var Sigma = "Σ";
var Tau = "Τ";
var Upsilon = "Υ";
var Phi = "Φ";
var Chi = "Χ";
var Psi = "Ψ";
var Omega = "Ω";
var alpha = "α";
var beta = "β";
var gamma = "γ";
var delta = "δ";
var epsilon = "ε";
var zeta = "ζ";
var eta = "η";
var theta = "θ";
var iota = "ι";
var kappa = "κ";
var lambda = "λ";
var mu = "μ";
var nu = "ν";
var xi = "ξ";
var omicron = "ο";
var pi = "π";
var rho = "ρ";
var sigmaf = "ς";
var sigma = "σ";
var tau = "τ";
var upsilon = "υ";
var phi = "φ";
var chi = "χ";
var psi = "ψ";
var omega = "ω";
var thetasym = "ϑ";
var upsih = "ϒ";
var piv = "ϖ";
var bull = "•";
var hellip = "…";
var prime = "′";
var Prime = "″";
var oline = "‾";
var frasl = "⁄";
var weierp = "℘";
var image = "ℑ";
var real = "ℜ";
var trade = "™";
var alefsym = "ℵ";
var larr = "←";
var uarr = "↑";
var rarr = "→";
var darr = "↓";
var harr = "↔";
var crarr = "↵";
var lArr = "⇐";
var uArr = "⇑";
var rArr = "⇒";
var dArr = "⇓";
var hArr = "⇔";
var forall = "∀";
var part = "∂";
var exist = "∃";
var empty = "∅";
var nabla = "∇";
var isin = "∈";
var notin = "∉";
var ni = "∋";
var prod = "∏";
var sum = "∑";
var minus = "−";
var lowast = "∗";
var radic = "√";
var prop = "∝";
var infin = "∞";
var ang = "∠";
var and = "∧";
var or = "∨";
var cap = "∩";
var cup = "∪";
var int$1 = "∫";
var there4 = "∴";
var sim = "∼";
var cong = "≅";
var asymp = "≈";
var ne = "≠";
var equiv = "≡";
var le = "≤";
var ge = "≥";
var sub = "⊂";
var sup = "⊃";
var nsub = "⊄";
var sube = "⊆";
var supe = "⊇";
var oplus = "⊕";
var otimes = "⊗";
var perp = "⊥";
var sdot = "⋅";
var lceil = "⌈";
var rceil = "⌉";
var lfloor = "⌊";
var rfloor = "⌋";
var lang = "〈";
var rang = "〉";
var loz = "◊";
var spades = "♠";
var clubs = "♣";
var hearts = "♥";
var diams = "♦";
var quot = "\"";
var amp = "&";
var lt = "<";
var gt = ">";
var OElig = "Œ";
var oelig = "œ";
var Scaron = "Š";
var scaron = "š";
var Yuml = "Ÿ";
var circ = "ˆ";
var tilde = "˜";
var ensp = " ";
var emsp = " ";
var thinsp = " ";
var zwnj = "‌";
var zwj = "‍";
var lrm = "‎";
var rlm = "‏";
var ndash = "–";
var mdash = "—";
var lsquo = "‘";
var rsquo = "’";
var sbquo = "‚";
var ldquo = "“";
var rdquo = "”";
var bdquo = "„";
var dagger = "†";
var Dagger = "‡";
var permil = "‰";
var lsaquo = "‹";
var rsaquo = "›";
var euro = "€";
var require$$0 = {
	nbsp: nbsp,
	iexcl: iexcl,
	cent: cent,
	pound: pound,
	curren: curren,
	yen: yen,
	brvbar: brvbar,
	sect: sect,
	uml: uml,
	copy: copy,
	ordf: ordf,
	laquo: laquo,
	not: not,
	shy: shy,
	reg: reg,
	macr: macr,
	deg: deg,
	plusmn: plusmn,
	sup2: sup2,
	sup3: sup3,
	acute: acute,
	micro: micro,
	para: para,
	middot: middot,
	cedil: cedil,
	sup1: sup1,
	ordm: ordm,
	raquo: raquo,
	frac14: frac14,
	frac12: frac12,
	frac34: frac34,
	iquest: iquest,
	Agrave: Agrave,
	Aacute: Aacute,
	Acirc: Acirc,
	Atilde: Atilde,
	Auml: Auml,
	Aring: Aring,
	AElig: AElig,
	Ccedil: Ccedil,
	Egrave: Egrave,
	Eacute: Eacute,
	Ecirc: Ecirc,
	Euml: Euml,
	Igrave: Igrave,
	Iacute: Iacute,
	Icirc: Icirc,
	Iuml: Iuml,
	ETH: ETH,
	Ntilde: Ntilde,
	Ograve: Ograve,
	Oacute: Oacute,
	Ocirc: Ocirc,
	Otilde: Otilde,
	Ouml: Ouml,
	times: times,
	Oslash: Oslash,
	Ugrave: Ugrave,
	Uacute: Uacute,
	Ucirc: Ucirc,
	Uuml: Uuml,
	Yacute: Yacute,
	THORN: THORN,
	szlig: szlig,
	agrave: agrave,
	aacute: aacute,
	acirc: acirc,
	atilde: atilde,
	auml: auml,
	aring: aring,
	aelig: aelig,
	ccedil: ccedil,
	egrave: egrave,
	eacute: eacute,
	ecirc: ecirc,
	euml: euml,
	igrave: igrave,
	iacute: iacute,
	icirc: icirc,
	iuml: iuml,
	eth: eth,
	ntilde: ntilde,
	ograve: ograve,
	oacute: oacute,
	ocirc: ocirc,
	otilde: otilde,
	ouml: ouml,
	divide: divide,
	oslash: oslash,
	ugrave: ugrave,
	uacute: uacute,
	ucirc: ucirc,
	uuml: uuml,
	yacute: yacute,
	thorn: thorn,
	yuml: yuml,
	fnof: fnof,
	Alpha: Alpha,
	Beta: Beta,
	Gamma: Gamma,
	Delta: Delta,
	Epsilon: Epsilon,
	Zeta: Zeta,
	Eta: Eta,
	Theta: Theta,
	Iota: Iota,
	Kappa: Kappa,
	Lambda: Lambda,
	Mu: Mu,
	Nu: Nu,
	Xi: Xi,
	Omicron: Omicron,
	Pi: Pi,
	Rho: Rho,
	Sigma: Sigma,
	Tau: Tau,
	Upsilon: Upsilon,
	Phi: Phi,
	Chi: Chi,
	Psi: Psi,
	Omega: Omega,
	alpha: alpha,
	beta: beta,
	gamma: gamma,
	delta: delta,
	epsilon: epsilon,
	zeta: zeta,
	eta: eta,
	theta: theta,
	iota: iota,
	kappa: kappa,
	lambda: lambda,
	mu: mu,
	nu: nu,
	xi: xi,
	omicron: omicron,
	pi: pi,
	rho: rho,
	sigmaf: sigmaf,
	sigma: sigma,
	tau: tau,
	upsilon: upsilon,
	phi: phi,
	chi: chi,
	psi: psi,
	omega: omega,
	thetasym: thetasym,
	upsih: upsih,
	piv: piv,
	bull: bull,
	hellip: hellip,
	prime: prime,
	Prime: Prime,
	oline: oline,
	frasl: frasl,
	weierp: weierp,
	image: image,
	real: real,
	trade: trade,
	alefsym: alefsym,
	larr: larr,
	uarr: uarr,
	rarr: rarr,
	darr: darr,
	harr: harr,
	crarr: crarr,
	lArr: lArr,
	uArr: uArr,
	rArr: rArr,
	dArr: dArr,
	hArr: hArr,
	forall: forall,
	part: part,
	exist: exist,
	empty: empty,
	nabla: nabla,
	isin: isin,
	notin: notin,
	ni: ni,
	prod: prod,
	sum: sum,
	minus: minus,
	lowast: lowast,
	radic: radic,
	prop: prop,
	infin: infin,
	ang: ang,
	and: and,
	or: or,
	cap: cap,
	cup: cup,
	int: int$1,
	there4: there4,
	sim: sim,
	cong: cong,
	asymp: asymp,
	ne: ne,
	equiv: equiv,
	le: le,
	ge: ge,
	sub: sub,
	sup: sup,
	nsub: nsub,
	sube: sube,
	supe: supe,
	oplus: oplus,
	otimes: otimes,
	perp: perp,
	sdot: sdot,
	lceil: lceil,
	rceil: rceil,
	lfloor: lfloor,
	rfloor: rfloor,
	lang: lang,
	rang: rang,
	loz: loz,
	spades: spades,
	clubs: clubs,
	hearts: hearts,
	diams: diams,
	quot: quot,
	amp: amp,
	lt: lt,
	gt: gt,
	OElig: OElig,
	oelig: oelig,
	Scaron: Scaron,
	scaron: scaron,
	Yuml: Yuml,
	circ: circ,
	tilde: tilde,
	ensp: ensp,
	emsp: emsp,
	thinsp: thinsp,
	zwnj: zwnj,
	zwj: zwj,
	lrm: lrm,
	rlm: rlm,
	ndash: ndash,
	mdash: mdash,
	lsquo: lsquo,
	rsquo: rsquo,
	sbquo: sbquo,
	ldquo: ldquo,
	rdquo: rdquo,
	bdquo: bdquo,
	dagger: dagger,
	Dagger: Dagger,
	permil: permil,
	lsaquo: lsaquo,
	rsaquo: rsaquo,
	euro: euro
};

var require$$5 = [
	"cent",
	"copy",
	"divide",
	"gt",
	"lt",
	"not",
	"para",
	"times"
];

var entities = require$$0;
var legacy = require$$1$2;
var hexadecimal = isHexadecimal;
var decimal = requireIsDecimal();
var alphanumerical = isAlphanumerical;
var dangerous = require$$5;

var stringifyEntities = encode;
encode.escape = escape$1;

var own$1 = {}.hasOwnProperty;

// List of enforced escapes.
var escapes = ['"', "'", '<', '>', '&', '`'];

// Map of characters to names.
var characters = construct();

// Default escapes.
var defaultEscapes = toExpression(escapes);

// Surrogate pairs.
var surrogatePair = /[\uD800-\uDBFF][\uDC00-\uDFFF]/g;

// Non-ASCII characters.
// eslint-disable-next-line no-control-regex, unicorn/no-hex-escape
var bmp = /[\x01-\t\x0B\f\x0E-\x1F\x7F\x81\x8D\x8F\x90\x9D\xA0-\uFFFF]/g;

// Encode special characters in `value`.
function encode(value, options) {
  var settings = options || {};
  var subset = settings.subset;
  var set = subset ? toExpression(subset) : defaultEscapes;
  var escapeOnly = settings.escapeOnly;
  var omit = settings.omitOptionalSemicolons;

  value = value.replace(set, replace);

  if (subset || escapeOnly) {
    return value
  }

  return value
    .replace(surrogatePair, replaceSurrogatePair)
    .replace(bmp, replace)

  function replaceSurrogatePair(pair, pos, val) {
    return toHexReference(
      (pair.charCodeAt(0) - 0xd800) * 0x400 +
        pair.charCodeAt(1) -
        0xdc00 +
        0x10000,
      val.charAt(pos + 2),
      omit
    )
  }

  function replace(char, pos, val) {
    return one$1(char, val.charAt(pos + 1), settings)
  }
}

// Shortcut to escape special characters in HTML.
function escape$1(value) {
  return encode(value, {escapeOnly: true, useNamedReferences: true})
}

// Encode `char` according to `options`.
function one$1(char, next, options) {
  var shortest = options.useShortestReferences;
  var omit = options.omitOptionalSemicolons;
  var named;
  var code;
  var numeric;
  var decimal;

  if ((shortest || options.useNamedReferences) && own$1.call(characters, char)) {
    named = toNamed(characters[char], next, omit, options.attribute);
  }

  if (shortest || !named) {
    code = char.charCodeAt(0);
    numeric = toHexReference(code, next, omit);

    // Use the shortest numeric reference when requested.
    // A simple algorithm would use decimal for all code points under 100, as
    // those are shorter than hexadecimal:
    //
    // * `&#99;` vs `&#x63;` (decimal shorter)
    // * `&#100;` vs `&#x64;` (equal)
    //
    // However, because we take `next` into consideration when `omit` is used,
    // And it would be possible that decimals are shorter on bigger values as
    // well if `next` is hexadecimal but not decimal, we instead compare both.
    if (shortest) {
      decimal = toDecimalReference(code, next, omit);

      if (decimal.length < numeric.length) {
        numeric = decimal;
      }
    }
  }

  if (named && (!shortest || named.length < numeric.length)) {
    return named
  }

  return numeric
}

// Transform `code` into an entity.
function toNamed(name, next, omit, attribute) {
  var value = '&' + name;

  if (
    omit &&
    own$1.call(legacy, name) &&
    dangerous.indexOf(name) === -1 &&
    (!attribute || (next && next !== '=' && !alphanumerical(next)))
  ) {
    return value
  }

  return value + ';'
}

// Transform `code` into a hexadecimal character reference.
function toHexReference(code, next, omit) {
  var value = '&#x' + code.toString(16).toUpperCase();
  return omit && next && !hexadecimal(next) ? value : value + ';'
}

// Transform `code` into a decimal character reference.
function toDecimalReference(code, next, omit) {
  var value = '&#' + String(code);
  return omit && next && !decimal(next) ? value : value + ';'
}

// Create an expression for `characters`.
function toExpression(characters) {
  return new RegExp('[' + characters.join('') + ']', 'g')
}

// Construct the map.
function construct() {
  var chars = {};
  var name;

  for (name in entities) {
    chars[entities[name]] = name;
  }

  return chars
}

// Characters.
var NULL = '\0';
var AMP = '&';
var SP = ' ';
var TB = '\t';
var GR = '`';
var DQ$1 = '"';
var SQ$1 = "'";
var EQ = '=';
var LT = '<';
var GT = '>';
var SO = '/';
var LF = '\n';
var CR = '\r';
var FF = '\f';

var whitespace = [SP, TB, LF, CR, FF];
// https://html.spec.whatwg.org/#attribute-name-state
var name = whitespace.concat(AMP, SO, GT, EQ);
// https://html.spec.whatwg.org/#attribute-value-(unquoted)-state
var unquoted$1 = whitespace.concat(AMP, GT);
var unquotedSafe = unquoted$1.concat(NULL, DQ$1, SQ$1, LT, EQ, GR);
// https://html.spec.whatwg.org/#attribute-value-(single-quoted)-state
var singleQuoted$1 = [AMP, SQ$1];
// https://html.spec.whatwg.org/#attribute-value-(double-quoted)-state
var doubleQuoted$1 = [AMP, DQ$1];

// Maps of subsets. Each value is a matrix of tuples.
// The first value causes parse errors, the second is valid.
// Of both values, the first value is unsafe, and the second is safe.
var constants = {
  name: [[name, name.concat(DQ$1, SQ$1, GR)], [name.concat(NULL, DQ$1, SQ$1, LT), name.concat(NULL, DQ$1, SQ$1, LT, GR)]],
  unquoted: [[unquoted$1, unquotedSafe], [unquotedSafe, unquotedSafe]],
  single: [
    [singleQuoted$1, singleQuoted$1.concat(DQ$1, GR)],
    [singleQuoted$1.concat(NULL), singleQuoted$1.concat(NULL, DQ$1, GR)]
  ],
  double: [
    [doubleQuoted$1, doubleQuoted$1.concat(SQ$1, GR)],
    [doubleQuoted$1.concat(NULL), doubleQuoted$1.concat(NULL, SQ$1, GR)]
  ]
};

var element_1;
var hasRequiredElement;

function requireElement () {
	if (hasRequiredElement) return element_1;
	hasRequiredElement = 1;

	var xtend = immutable;
	var svg = svg_1;
	var find = find_1;
	var spaces = spaceSeparatedTokens.stringify;
	var commas = commaSeparatedTokens.stringify;
	var entities = stringifyEntities;
	var all = requireAll();
	var constants$1 = constants;
	const repeat = repeatString;

	element_1 = element;

	/* Constants. */
	var emptyString = '';

	/* Characters. */
	var space = ' ';
	var quotationMark = '"';
	var apostrophe = "'";
	var equalsTo = '=';
	var lessThan = '<';
	var greaterThan = '>';
	var slash = '/';
	var newLine = '\n';

	/* Stringify an element `node`. */
	function element(ctx, node, index, parent, printWidthOffset, innerTextLength) {
	  var parentSchema = ctx.schema;
	  var name = node.tagName;
	  var value = '';
	  var selfClosing;
	  var close;
	  var omit;
	  var root = node;
	  var content;
	  var attrs;
	  var indentLevel = getNodeData(node, 'indentLevel', 0);
	  var printContext = {
	    offset: printWidthOffset,
	    wrapAttributes: false,
	    indentLevel
	  };
	  var isVoid = ctx.voids.indexOf(name) !== -1;
	  var ignoreAttrCollapsing =
	    getNodeData(node, 'ignore', false) || getNodeData(node, 'preserveAttrWrapping', false);

	  if (parentSchema.space === 'html' && name === 'svg') {
	    ctx.schema = svg;
	  }

	  if (ctx.schema.space === 'svg') {
	    omit = false;
	    close = true;
	    selfClosing = ctx.closeEmpty;
	  } else {
	    omit = ctx.omit;
	    close = ctx.close;
	    selfClosing = isVoid;
	  }

	  // check for 'selfClosing' property set by hast-util-from-webparser package
	  // in order to support custom self-closing elements
	  if (selfClosing === false) {
	    selfClosing = getNodeData(node, 'selfClosing', false);
	  }

	  // <
	  printContext.offset += lessThan.length;

	  // tagName length
	  printContext.offset += node.tagName.length;

	  // / closing tag
	  if (selfClosing && !isVoid) {
	    printContext.offset += slash.length;
	  }

	  // >
	  printContext.offset += greaterThan.length;

	  const propertyCount = Object.keys(node.properties).length;

	  // force to wrap attributes on multiple lines when the node contains
	  // more than one attribute
	  if (propertyCount > 1 && ctx.wrapAttributes) {
	    printContext.wrapAttributes = true;
	  }

	  // one space before each attribute
	  if (propertyCount) {
	    printContext.offset += propertyCount * space.length;
	  }

	  // represent the length of the inner text of the node
	  printContext.offset += innerTextLength;

	  attrs = attributes(ctx, node.properties, printContext, ignoreAttrCollapsing);

	  const shouldCollapse = ignoreAttrCollapsing === false && printContext.wrapAttributes;

	  content = all(ctx, root);

	  /* If the node is categorised as void, but it has
	   * children, remove the categorisation.  This
	   * enables for example `menuitem`s, which are
	   * void in W3C HTML but not void in WHATWG HTML, to
	   * be stringified properly. */
	  selfClosing = content ? false : selfClosing;

	  if (attrs || !omit || !omit.opening(node, index, parent)) {
	    value = lessThan + name;

	    if (attrs) {
	      // add no space after tagName when element is collapsed
	      if (shouldCollapse) {
	        value += attrs;
	      } else {
	        value += space + attrs;
	      }
	    }

	    let selfClosed = false;

	    // check if the should close self-closing elements
	    if (selfClosing && close) {
	      if ((!ctx.tightClose || attrs.charAt(attrs.length - 1) === slash) && !shouldCollapse) {
	        value += space;
	      }

	      if (shouldCollapse) {
	        value += newLine + repeat(ctx.tabWidth, printContext.indentLevel);
	      }

	      selfClosed = true;
	      value += slash;
	    }

	    // allow any element to self close itself except known HTML void elements
	    else if (selfClosing && !isVoid) {
	      if (shouldCollapse) {
	        value += newLine + repeat(ctx.tabWidth, printContext.indentLevel);
	      }

	      selfClosed = true;
	      value += slash;
	    }

	    // add newline when element should be wrappend on multiple lines and when
	    // it's no self-closing element because in that case the newline was already added before the slash (/)
	    if (shouldCollapse && !selfClosed) {
	      value += newLine + repeat(ctx.tabWidth, printContext.indentLevel);
	    }

	    value += greaterThan;
	  }

	  value += content;

	  if (!selfClosing && (!omit || !omit.closing(node, index, parent))) {
	    value += lessThan + slash + name + greaterThan;
	  }

	  ctx.schema = parentSchema;

	  return value
	}

	/* Stringify all attributes. */
	function attributes(ctx, props, printContext, ignoreIndent) {
	  var values = [];
	  var key;
	  var value;
	  var result;
	  var length;
	  var index;
	  var last;

	  for (key in props) {
	    value = props[key];

	    if (value == null) {
	      continue
	    }

	    result = attribute(ctx, key, value);

	    printContext.offset += result.length;

	    if (ignoreIndent === false && printContext.offset > ctx.printWidth) {
	      printContext.wrapAttributes = true;
	    }

	    if (result) {
	      values.push(result);
	    }
	  }

	  length = values.length;
	  index = -1;

	  while (++index < length) {
	    result = values[index];
	    last = null;

	    /* In tight mode, don’t add a space after quoted attributes. */
	    if (last !== quotationMark && last !== apostrophe) {
	      if (printContext.wrapAttributes) {
	        values[index] = newLine + repeat(ctx.tabWidth, printContext.indentLevel + 1) + result;
	      } else if (index !== length - 1) {
	        values[index] = result + space;
	      } else {
	        values[index] = result;
	      }
	    }
	  }

	  return values.join(emptyString)
	}

	/* Stringify one attribute. */
	function attribute(ctx, key, value) {
	  var schema = ctx.schema;
	  var info = find(schema, key);
	  var name = info.attribute;

	  if (value == null || (typeof value === 'number' && isNaN(value)) || (value === false && info.boolean)) {
	    return emptyString
	  }

	  name = attributeName(ctx, name);

	  if ((value === true && info.boolean) || (value === true && info.overloadedBoolean)) {
	    return name
	  }

	  return name + attributeValue(ctx, key, value, info)
	}

	/* Stringify the attribute name. */
	function attributeName(ctx, name) {
	  // Always encode without parse errors in non-HTML.
	  var valid = ctx.schema.space === 'html' ? ctx.valid : 1;
	  var subset = constants$1.name[valid][ctx.safe];

	  return entities(name, xtend(ctx.entities, { subset: subset }))
	}

	/* Stringify the attribute value. */
	function attributeValue(ctx, key, value, info) {
	  var quote = ctx.quote;

	  if (typeof value === 'object' && 'length' in value) {
	    /* `spaces` doesn’t accept a second argument, but it’s
	     * given here just to keep the code cleaner. */
	    value = (info.commaSeparated ? commas : spaces)(value, {
	      padLeft: !ctx.tightLists
	    });
	  }

	  value = String(value);

	  // When attr has no value we avoid quoting
	  if (value === '') {
	    return value
	  } else {
	    value = equalsTo + quote + value + quote;
	  }

	  return value
	}

	function getNodeData(node, key, defaultValue) {
	  let data = node.data || {};
	  return data[key] || defaultValue
	}
	return element_1;
}

var doctype_1 = doctype;

/* Stringify a doctype `node`. */
function doctype(ctx, node) {
  var sep = ctx.tightDoctype ? '' : ' ';
  var name = node.name;
  var pub = node.public;
  var sys = node.system;
  var val = ['<!doctype'];

  if (name) {
    val.push(sep, name);

    if (pub != null) {
      val.push(' public', sep, smart(pub));
    } else if (sys != null) {
      val.push(' system');
    }

    if (sys != null) {
      val.push(sep, smart(sys));
    }
  }

  return val.join('') + '>'
}

function smart(value) {
  var quote = value.indexOf('"') === -1 ? '"' : "'";
  return quote + value + quote
}

var comment_1 = comment;

/* Stringify a comment `node`. */
function comment(ctx, node) {
  return '<!--' + node.value + '-->'
}

var raw_1 = raw;

/* Stringify `raw`. */
function raw(ctx, node) {
  return node.value
}

var one_1;
var hasRequiredOne;

function requireOne () {
	if (hasRequiredOne) return one_1;
	hasRequiredOne = 1;

	one_1 = one;

	var own = {}.hasOwnProperty;

	var handlers = {};

	handlers.root = requireAll();
	handlers.text = text_1;
	handlers.element = requireElement();
	handlers.doctype = doctype_1;
	handlers.comment = comment_1;
	handlers.raw = raw_1;

	/* Stringify `node`. */
	function one(ctx, node, index, parent, printWidthOffset, innerTextLength) {
	  var type = node && node.type;

	  if (!type) {
	    throw new Error('Expected node, not `' + node + '`')
	  }

	  if (!own.call(handlers, type)) {
	    throw new Error('Cannot compile unknown node `' + type + '`')
	  }

	  return handlers[type](ctx, node, index, parent, printWidthOffset, innerTextLength)
	}
	return one_1;
}

var html = html_1;
var svg = svg_1;
var voids = require$$2;
var omission = omission$4;
var one = requireOne();
const repeat$1 = repeatString;

var lib$1 = toHTML;

/* Characters. */
var DQ = '"';
var SQ = "'";

/* Stringify the given HAST node. */
function toHTML(node, options) {
  var settings = options || {};
  var quote = settings.singleQuote ? SQ : DQ;
  var printWidth = settings.printWidth === undefined ? 80 : settings.printWidth;
  var useTabs = settings.useTabs;
  var tabWidth = settings.tabWidth || 2;
  var wrapAttributes = settings.wrapAttributes;

  if (useTabs) {
    tabWidth = '\t';
  } else if (typeof tabWidth === 'number') {
    tabWidth = repeat$1(' ', tabWidth);
  }

  return one(
    {
      valid: settings.allowParseErrors ? 0 : 1,
      safe: settings.allowDangerousCharacters ? 0 : 1,
      schema: settings.space === 'svg' ? svg : html,
      omit: settings.omitOptionalTags && omission,
      quote: quote,
      printWidth: printWidth,
      tabWidth: tabWidth,
      wrapAttributes: wrapAttributes,
      tightDoctype: Boolean(settings.tightDoctype),
      tightLists: settings.tightCommaSeparatedLists,
      voids: settings.voids || voids.concat(),
      entities: settings.entities || {},
      close: settings.closeSelfClosing,
      tightClose: settings.tightSelfClosing,
      closeEmpty: settings.closeEmptyElements
    },
    node
  )
}

var prettyhtmlHastToHtml = lib$1;

var hast_to_html = /*@__PURE__*/getDefaultExportFromCjs(prettyhtmlHastToHtml);

const void_els = [
	'area',
	'base',
	'br',
	'col',
	'embed',
	'hr',
	'img',
	'input',
	'link',
	'meta',
	'param',
	'source',
	'track',
	'wbr',
];

// these regex don't check if it is a valid svelte tag name
// i want to defer to svelte's compiler errors so i don't end up reimplementing the svelte parser

const RE_SVELTE_TAG = /^<svelte:([a-z]*)[\s\S]*(?:(?:svelte:[a-z]*)|(?:\/))>$/;
const RE_SVELTE_TAG_START = /(^\s*)<([\\/\s])*svelte:/;

function parse_svelte_tag(
	eat,
	value,
	silent
) {
	const is_svelte_tag = RE_SVELTE_TAG_START.exec(value);

	if (is_svelte_tag) {
		if (silent) return true;

		const trimmed_value = value.trim();
		let cbPos = 0;
		let pos = 1;
		let current_tag = '';
		let in_tag_name = false;

		while (cbPos > -1) {
			if (!trimmed_value[pos]) {
				break;
			}

			if (trimmed_value[pos].match(/</)) {
				cbPos++;
				current_tag = '';
				in_tag_name = true;
			}

			if (in_tag_name && trimmed_value[pos].match(/\s/)) {
				in_tag_name = false;
			}

			if (in_tag_name && !trimmed_value[pos].match(/</)) {
				current_tag += trimmed_value[pos];
			}

			const is_void = void_els.includes(current_tag);

			if (
				(is_void && trimmed_value[pos].match(/>/)) ||
				(trimmed_value[pos - 1] + trimmed_value[pos]).match(/\/>/)
			) {
				cbPos--;
			}

			if ((trimmed_value[pos - 1] + trimmed_value[pos]).match(/<\//)) {
				let inner_indent = 0;

				while (inner_indent > -1) {
					if (trimmed_value[pos].match(/>/)) {
						pos++;
						inner_indent -= 1;
						cbPos -= 2;
					} else {
						pos++;
					}
				}
			}

			pos++;
		}

		const match = RE_SVELTE_TAG.exec(trimmed_value.substring(0, pos).trim());

		if (!match) return;

		return eat(is_svelte_tag[1] + match[0])({
			type: 'svelteTag',
			value: match[0],
			name: match[1],
		});
	}
}

// these regex don't check if it is a valid block name
// i want to defer to svelte's compiler errors so i don't end up reimplementing the svelte parser
// 'else if' is a special case due to the annoying whitespace

const RE_SVELTE_BLOCK_START = /(^\s*){[#:/@]/;
const RE_SVELTE_BLOCK = /^{[#:/@](else if|[a-z]+).*}$/;

function parse_svelte_block(
	eat,
	value,
	silent
) {
	const is_svelte_block = RE_SVELTE_BLOCK_START.exec(value);

	if (is_svelte_block) {
		if (silent) return true;

		const trimmed_value = value.trim();
		let cbPos = 0;
		let pos = 1;

		while (cbPos > -1) {
			if (trimmed_value[pos].match(/{/)) cbPos++;
			if (trimmed_value[pos].match(/}/)) cbPos--;
			pos++;
		}

		const match = RE_SVELTE_BLOCK.exec(trimmed_value.substring(0, pos));

		if (!match) return;

		return eat(is_svelte_block[1] + match[0])({
			type: 'svelteBlock',
			value: `${is_svelte_block[1]}${match[0]}`,
			name: match[1],
		});
	}
}

const dotAllPolyfill = '[\0-\uFFFF]';

const attributeName = '[a-zA-Z_:][a-zA-Z0-9:._-]*';
const unquoted = '[^"\'=<>`\\u0000-\\u0020]+';
const singleQuoted = "'[^']*'";
const doubleQuoted = '"[^"]*"';
const jsProps = '{.*}'.replace('.', dotAllPolyfill);
const attributeValue =
	'(?:' +
	unquoted +
	'|' +
	singleQuoted +
	'|' +
	doubleQuoted +
	'|' +
	jsProps +
	')';
const attribute =
	'(?:\\s+' + attributeName + '(?:\\s*=\\s*' + attributeValue + ')?)';
const openTag = '<[A-Za-z]*[A-Za-z0-9\\.\\-]*' + attribute + '*\\s*\\/?>';
const closeTag = '<\\/[A-Za-z][A-Za-z0-9\\.\\-]*\\s*>';

const openCloseTag = new RegExp('^(?:' + openTag + '|' + closeTag + ')');

const tab = '\t';
const space = ' ';
const lineFeed = '\n';
const lessThan = '<';

const rawOpenExpression = /^<(script|pre|style)(?=(\s|>|$))/i;
const rawCloseExpression = /<\/(script|pre|style)>/i;
const commentOpenExpression = /^<!--/;
const commentCloseExpression = /-->/;
const instructionOpenExpression = /^<\?/;
const instructionCloseExpression = /\?>/;
const directiveOpenExpression = /^<![A-Za-z]/;
const directiveCloseExpression = />/;
const cdataOpenExpression = /^<!\[CDATA\[/;
const cdataCloseExpression = /\]\]>/;
const elementCloseExpression = /^$/;
const otherElementOpenExpression = new RegExp(openCloseTag.source + '\\s*$');
const fragmentOpenExpression = /^<>/;

//@ts-ignore
function blockHtml(
	eat,
	value,
	silent
) {
	const blocks = '[a-z\\.]*(\\.){0,1}[a-z][a-z0-9\\.]*';
	const elementOpenExpression = new RegExp(
		'^</?(' + blocks + ')(?=(\\s|/?>|$))',
		'i'
	);

	const length = value.length;
	let index = 0;
	let next;
	let line;
	let offset;
	let character;
	let sequence;

	const sequences = [
		[rawOpenExpression, rawCloseExpression, true],
		[commentOpenExpression, commentCloseExpression, true],
		[instructionOpenExpression, instructionCloseExpression, true],
		[directiveOpenExpression, directiveCloseExpression, true],
		[cdataOpenExpression, cdataCloseExpression, true],
		[elementOpenExpression, elementCloseExpression, true],
		[fragmentOpenExpression, elementCloseExpression, true],
		[otherElementOpenExpression, elementCloseExpression, false],
	];

	// Eat initial spacing.
	while (index < length) {
		character = value.charAt(index);

		if (character !== tab && character !== space) {
			break;
		}

		index++;
	}

	if (value.charAt(index) !== lessThan) {
		return;
	}

	next = value.indexOf(lineFeed, index + 1);
	next = next === -1 ? length : next;
	line = value.slice(index, next);
	offset = -1;
	const count = sequences.length;

	while (++offset < count) {
		if (sequences[offset][0].test(line)) {
			sequence = sequences[offset];
			break;
		}
	}

	if (!sequence) {
		return;
	}

	if (silent) {
		return sequence[2];
	}

	index = next;

	if (!sequence[1].test(line)) {
		while (index < length) {
			next = value.indexOf(lineFeed, index + 1);
			next = next === -1 ? length : next;
			line = value.slice(index + 1, next);

			if (sequence[1].test(line)) {
				if (line) {
					index = next;
				}

				break;
			}

			index = next;
		}
	}

	const subvalue = value.slice(0, index);

	return eat(subvalue)({ type: 'html', value: subvalue });
}

function mdsvex_parser() {
	const Parser = this.Parser;
	const block_tokenizers = Parser.prototype.blockTokenizers;
	const methods = Parser.prototype.blockMethods;

	block_tokenizers.svelteBlock = parse_svelte_block;
	block_tokenizers.svelteTag = parse_svelte_tag;
	block_tokenizers.html = blockHtml;
	block_tokenizers.indentedCode = indentedCode;

	methods.splice(methods.indexOf('html'), 0, 'svelteBlock');
	methods.splice(methods.indexOf('html'), 0, 'svelteTag');
}

function indentedCode() {
	return true;
}

var extend$1 = extend$5;
var bail = bail_1;
var vfile = vfile$2;
var trough = trough_1;
var plain = isPlainObj;

// Expose a frozen processor.
var unified_1 = unified$1().freeze();

var slice = [].slice;
var own = {}.hasOwnProperty;

// Process pipeline.
var pipeline = trough()
  .use(pipelineParse)
  .use(pipelineRun)
  .use(pipelineStringify);

function pipelineParse(p, ctx) {
  ctx.tree = p.parse(ctx.file);
}

function pipelineRun(p, ctx, next) {
  p.run(ctx.tree, ctx.file, done);

  function done(err, tree, file) {
    if (err) {
      next(err);
    } else {
      ctx.tree = tree;
      ctx.file = file;
      next();
    }
  }
}

function pipelineStringify(p, ctx) {
  ctx.file.contents = p.stringify(ctx.tree, ctx.file);
}

// Function to create the first processor.
function unified$1() {
  var attachers = [];
  var transformers = trough();
  var namespace = {};
  var frozen = false;
  var freezeIndex = -1;

  // Data management.
  processor.data = data;

  // Lock.
  processor.freeze = freeze;

  // Plugins.
  processor.attachers = attachers;
  processor.use = use;

  // API.
  processor.parse = parse;
  processor.stringify = stringify;
  processor.run = run;
  processor.runSync = runSync;
  processor.process = process;
  processor.processSync = processSync;

  // Expose.
  return processor

  // Create a new processor based on the processor in the current scope.
  function processor() {
    var destination = unified$1();
    var length = attachers.length;
    var index = -1;

    while (++index < length) {
      destination.use.apply(null, attachers[index]);
    }

    destination.data(extend$1(true, {}, namespace));

    return destination
  }

  // Freeze: used to signal a processor that has finished configuration.
  //
  // For example, take unified itself: it’s frozen.
  // Plugins should not be added to it.
  // Rather, it should be extended, by invoking it, before modifying it.
  //
  // In essence, always invoke this when exporting a processor.
  function freeze() {
    var values;
    var plugin;
    var options;
    var transformer;

    if (frozen) {
      return processor
    }

    while (++freezeIndex < attachers.length) {
      values = attachers[freezeIndex];
      plugin = values[0];
      options = values[1];
      transformer = null;

      if (options === false) {
        continue
      }

      if (options === true) {
        values[1] = undefined;
      }

      transformer = plugin.apply(processor, values.slice(1));

      if (typeof transformer === 'function') {
        transformers.use(transformer);
      }
    }

    frozen = true;
    freezeIndex = Infinity;

    return processor
  }

  // Data management.
  // Getter / setter for processor-specific informtion.
  function data(key, value) {
    if (typeof key === 'string') {
      // Set `key`.
      if (arguments.length === 2) {
        assertUnfrozen('data', frozen);

        namespace[key] = value;

        return processor
      }

      // Get `key`.
      return (own.call(namespace, key) && namespace[key]) || null
    }

    // Set space.
    if (key) {
      assertUnfrozen('data', frozen);
      namespace = key;
      return processor
    }

    // Get space.
    return namespace
  }

  // Plugin management.
  //
  // Pass it:
  // *   an attacher and options,
  // *   a preset,
  // *   a list of presets, attachers, and arguments (list of attachers and
  //     options).
  function use(value) {
    var settings;

    assertUnfrozen('use', frozen);

    if (value === null || value === undefined) ; else if (typeof value === 'function') {
      addPlugin.apply(null, arguments);
    } else if (typeof value === 'object') {
      if ('length' in value) {
        addList(value);
      } else {
        addPreset(value);
      }
    } else {
      throw new Error('Expected usable value, not `' + value + '`')
    }

    if (settings) {
      namespace.settings = extend$1(namespace.settings || {}, settings);
    }

    return processor

    function addPreset(result) {
      addList(result.plugins);

      if (result.settings) {
        settings = extend$1(settings || {}, result.settings);
      }
    }

    function add(value) {
      if (typeof value === 'function') {
        addPlugin(value);
      } else if (typeof value === 'object') {
        if ('length' in value) {
          addPlugin.apply(null, value);
        } else {
          addPreset(value);
        }
      } else {
        throw new Error('Expected usable value, not `' + value + '`')
      }
    }

    function addList(plugins) {
      var length;
      var index;

      if (plugins === null || plugins === undefined) ; else if (typeof plugins === 'object' && 'length' in plugins) {
        length = plugins.length;
        index = -1;

        while (++index < length) {
          add(plugins[index]);
        }
      } else {
        throw new Error('Expected a list of plugins, not `' + plugins + '`')
      }
    }

    function addPlugin(plugin, value) {
      var entry = find(plugin);

      if (entry) {
        if (plain(entry[1]) && plain(value)) {
          value = extend$1(entry[1], value);
        }

        entry[1] = value;
      } else {
        attachers.push(slice.call(arguments));
      }
    }
  }

  function find(plugin) {
    var length = attachers.length;
    var index = -1;
    var entry;

    while (++index < length) {
      entry = attachers[index];

      if (entry[0] === plugin) {
        return entry
      }
    }
  }

  // Parse a file (in string or vfile representation) into a unist node using
  // the `Parser` on the processor.
  function parse(doc) {
    var file = vfile(doc);
    var Parser;

    freeze();
    Parser = processor.Parser;
    assertParser('parse', Parser);

    if (newable(Parser, 'parse')) {
      return new Parser(String(file), file).parse()
    }

    return Parser(String(file), file) // eslint-disable-line new-cap
  }

  // Run transforms on a unist node representation of a file (in string or
  // vfile representation), async.
  function run(node, file, cb) {
    assertNode(node);
    freeze();

    if (!cb && typeof file === 'function') {
      cb = file;
      file = null;
    }

    if (!cb) {
      return new Promise(executor)
    }

    executor(null, cb);

    function executor(resolve, reject) {
      transformers.run(node, vfile(file), done);

      function done(err, tree, file) {
        tree = tree || node;
        if (err) {
          reject(err);
        } else if (resolve) {
          resolve(tree);
        } else {
          cb(null, tree, file);
        }
      }
    }
  }

  // Run transforms on a unist node representation of a file (in string or
  // vfile representation), sync.
  function runSync(node, file) {
    var complete = false;
    var result;

    run(node, file, done);

    assertDone('runSync', 'run', complete);

    return result

    function done(err, tree) {
      complete = true;
      bail(err);
      result = tree;
    }
  }

  // Stringify a unist node representation of a file (in string or vfile
  // representation) into a string using the `Compiler` on the processor.
  function stringify(node, doc) {
    var file = vfile(doc);
    var Compiler;

    freeze();
    Compiler = processor.Compiler;
    assertCompiler('stringify', Compiler);
    assertNode(node);

    if (newable(Compiler, 'compile')) {
      return new Compiler(node, file).compile()
    }

    return Compiler(node, file) // eslint-disable-line new-cap
  }

  // Parse a file (in string or vfile representation) into a unist node using
  // the `Parser` on the processor, then run transforms on that node, and
  // compile the resulting node using the `Compiler` on the processor, and
  // store that result on the vfile.
  function process(doc, cb) {
    freeze();
    assertParser('process', processor.Parser);
    assertCompiler('process', processor.Compiler);

    if (!cb) {
      return new Promise(executor)
    }

    executor(null, cb);

    function executor(resolve, reject) {
      var file = vfile(doc);

      pipeline.run(processor, {file: file}, done);

      function done(err) {
        if (err) {
          reject(err);
        } else if (resolve) {
          resolve(file);
        } else {
          cb(null, file);
        }
      }
    }
  }

  // Process the given document (in string or vfile representation), sync.
  function processSync(doc) {
    var complete = false;
    var file;

    freeze();
    assertParser('processSync', processor.Parser);
    assertCompiler('processSync', processor.Compiler);
    file = vfile(doc);

    process(file, done);

    assertDone('processSync', 'process', complete);

    return file

    function done(err) {
      complete = true;
      bail(err);
    }
  }
}

// Check if `value` is a constructor.
function newable(value, name) {
  return (
    typeof value === 'function' &&
    value.prototype &&
    // A function with keys in its prototype is probably a constructor.
    // Classes’ prototype methods are not enumerable, so we check if some value
    // exists in the prototype.
    (keys(value.prototype) || name in value.prototype)
  )
}

// Check if `value` is an object with keys.
function keys(value) {
  var key;
  for (key in value) {
    return true
  }

  return false
}

// Assert a parser is available.
function assertParser(name, Parser) {
  if (typeof Parser !== 'function') {
    throw new Error('Cannot `' + name + '` without `Parser`')
  }
}

// Assert a compiler is available.
function assertCompiler(name, Compiler) {
  if (typeof Compiler !== 'function') {
    throw new Error('Cannot `' + name + '` without `Compiler`')
  }
}

// Assert the processor is not frozen.
function assertUnfrozen(name, frozen) {
  if (frozen) {
    throw new Error(
      'Cannot invoke `' +
        name +
        '` on a frozen processor.\nCreate a new processor first, by invoking it: use `processor()` instead of `processor`.'
    )
  }
}

// Assert `node` is a unist node.
function assertNode(node) {
  if (!node || typeof node.type !== 'string') {
    throw new Error('Expected node, got `' + node + '`')
  }
}

// Assert that `complete` is `true`.
function assertDone(name, asyncName, complete) {
  if (!complete) {
    throw new Error(
      '`' + name + '` finished async. Use `' + asyncName + '` instead'
    )
  }
}

var nlcstToString_1 = nlcstToString;

// Stringify one nlcst node or list of nodes.
function nlcstToString(node, separator) {
  var sep = separator || '';
  var values;
  var length;
  var children;

  if (!node || (!('length' in node) && !node.type)) {
    throw new Error('Expected node, not `' + node + '`')
  }

  if (typeof node.value === 'string') {
    return node.value
  }

  children = 'length' in node ? node : node.children;
  length = children.length;

  // Shortcut: This is pretty common, and a small performance win.
  if (length === 1 && 'value' in children[0]) {
    return children[0].value
  }

  values = [];

  while (length--) {
    values[length] = nlcstToString(children[length], sep);
  }

  return values.join(sep)
}

var toString$2 = nlcstToString_1;

var tokenizer$1 = tokenizerFactory;

// Factory to create a tokenizer based on a given `expression`.
function tokenizerFactory(childType, expression) {
  return tokenizer

  // A function that splits.
  function tokenizer(node) {
    var children = [];
    var tokens = node.children;
    var type = node.type;
    var length = tokens.length;
    var index = -1;
    var lastIndex = length - 1;
    var start = 0;
    var first;
    var last;
    var parent;

    while (++index < length) {
      if (
        index === lastIndex ||
        (tokens[index].type === childType &&
          expression.test(toString$2(tokens[index])))
      ) {
        first = tokens[start];
        last = tokens[index];

        parent = {
          type: type,
          children: tokens.slice(start, index + 1)
        };

        if (first.position && last.position) {
          parent.position = {
            start: first.position.start,
            end: last.position.end
          };
        }

        children.push(parent);

        start = index + 1;
      }
    }

    return children
  }
}

var tokenizer = tokenizer$1;

var parser = parserFactory;

// Construct a parser based on `options`.
function parserFactory(options) {
  var type = options.type;
  var tokenizerProperty = options.tokenizer;
  var delimiter = options.delimiter;
  var tokenize = delimiter && tokenizer(options.delimiterType, delimiter);

  return parser

  function parser(value) {
    var children = this[tokenizerProperty](value);

    return {
      type: type,
      children: tokenize ? tokenize(children) : children
    }
  }
}

var expressions$1 = {
  affixSymbol: /^([\)\]\}\u0F3B\u0F3D\u169C\u2046\u207E\u208E\u2309\u230B\u232A\u2769\u276B\u276D\u276F\u2771\u2773\u2775\u27C6\u27E7\u27E9\u27EB\u27ED\u27EF\u2984\u2986\u2988\u298A\u298C\u298E\u2990\u2992\u2994\u2996\u2998\u29D9\u29DB\u29FD\u2E23\u2E25\u2E27\u2E29\u3009\u300B\u300D\u300F\u3011\u3015\u3017\u3019\u301B\u301E\u301F\uFD3E\uFE18\uFE36\uFE38\uFE3A\uFE3C\uFE3E\uFE40\uFE42\uFE44\uFE48\uFE5A\uFE5C\uFE5E\uFF09\uFF3D\uFF5D\uFF60\uFF63]|["'\xBB\u2019\u201D\u203A\u2E03\u2E05\u2E0A\u2E0D\u2E1D\u2E21]|[!\.\?\u2026\u203D])\1*$/,
  newLine: /^[ \t]*((\r?\n|\r)[\t ]*)+$/,
  newLineMulti: /^[ \t]*((\r?\n|\r)[\t ]*){2,}$/,
  terminalMarker: /^((?:[!\.\?\u2026\u203D])+)$/,
  wordSymbolInner: /^((?:[&'\x2D\.:=\?@\xAD\xB7\u2010\u2011\u2019\u2027])|(?:_)+)$/,
  numerical: /^(?:[0-9\xB2\xB3\xB9\xBC-\xBE\u0660-\u0669\u06F0-\u06F9\u07C0-\u07C9\u0966-\u096F\u09E6-\u09EF\u09F4-\u09F9\u0A66-\u0A6F\u0AE6-\u0AEF\u0B66-\u0B6F\u0B72-\u0B77\u0BE6-\u0BF2\u0C66-\u0C6F\u0C78-\u0C7E\u0CE6-\u0CEF\u0D58-\u0D5E\u0D66-\u0D78\u0DE6-\u0DEF\u0E50-\u0E59\u0ED0-\u0ED9\u0F20-\u0F33\u1040-\u1049\u1090-\u1099\u1369-\u137C\u16EE-\u16F0\u17E0-\u17E9\u17F0-\u17F9\u1810-\u1819\u1946-\u194F\u19D0-\u19DA\u1A80-\u1A89\u1A90-\u1A99\u1B50-\u1B59\u1BB0-\u1BB9\u1C40-\u1C49\u1C50-\u1C59\u2070\u2074-\u2079\u2080-\u2089\u2150-\u2182\u2185-\u2189\u2460-\u249B\u24EA-\u24FF\u2776-\u2793\u2CFD\u3007\u3021-\u3029\u3038-\u303A\u3192-\u3195\u3220-\u3229\u3248-\u324F\u3251-\u325F\u3280-\u3289\u32B1-\u32BF\uA620-\uA629\uA6E6-\uA6EF\uA830-\uA835\uA8D0-\uA8D9\uA900-\uA909\uA9D0-\uA9D9\uA9F0-\uA9F9\uAA50-\uAA59\uABF0-\uABF9\uFF10-\uFF19]|\uD800[\uDD07-\uDD33\uDD40-\uDD78\uDD8A\uDD8B\uDEE1-\uDEFB\uDF20-\uDF23\uDF41\uDF4A\uDFD1-\uDFD5]|\uD801[\uDCA0-\uDCA9]|\uD802[\uDC58-\uDC5F\uDC79-\uDC7F\uDCA7-\uDCAF\uDCFB-\uDCFF\uDD16-\uDD1B\uDDBC\uDDBD\uDDC0-\uDDCF\uDDD2-\uDDFF\uDE40-\uDE48\uDE7D\uDE7E\uDE9D-\uDE9F\uDEEB-\uDEEF\uDF58-\uDF5F\uDF78-\uDF7F\uDFA9-\uDFAF]|\uD803[\uDCFA-\uDCFF\uDD30-\uDD39\uDE60-\uDE7E\uDF1D-\uDF26\uDF51-\uDF54]|\uD804[\uDC52-\uDC6F\uDCF0-\uDCF9\uDD36-\uDD3F\uDDD0-\uDDD9\uDDE1-\uDDF4\uDEF0-\uDEF9]|\uD805[\uDC50-\uDC59\uDCD0-\uDCD9\uDE50-\uDE59\uDEC0-\uDEC9\uDF30-\uDF3B]|\uD806[\uDCE0-\uDCF2]|\uD807[\uDC50-\uDC6C\uDD50-\uDD59\uDDA0-\uDDA9\uDFC0-\uDFD4]|\uD809[\uDC00-\uDC6E]|\uD81A[\uDE60-\uDE69\uDF50-\uDF59\uDF5B-\uDF61]|\uD81B[\uDE80-\uDE96]|\uD834[\uDEE0-\uDEF3\uDF60-\uDF78]|\uD835[\uDFCE-\uDFFF]|\uD838[\uDD40-\uDD49\uDEF0-\uDEF9]|\uD83A[\uDCC7-\uDCCF\uDD50-\uDD59]|\uD83B[\uDC71-\uDCAB\uDCAD-\uDCAF\uDCB1-\uDCB4\uDD01-\uDD2D\uDD2F-\uDD3D]|\uD83C[\uDD00-\uDD0C])+$/,
  digitStart: /^\d/,
  lowerInitial: /^(?:[a-z\xB5\xDF-\xF6\xF8-\xFF\u0101\u0103\u0105\u0107\u0109\u010B\u010D\u010F\u0111\u0113\u0115\u0117\u0119\u011B\u011D\u011F\u0121\u0123\u0125\u0127\u0129\u012B\u012D\u012F\u0131\u0133\u0135\u0137\u0138\u013A\u013C\u013E\u0140\u0142\u0144\u0146\u0148\u0149\u014B\u014D\u014F\u0151\u0153\u0155\u0157\u0159\u015B\u015D\u015F\u0161\u0163\u0165\u0167\u0169\u016B\u016D\u016F\u0171\u0173\u0175\u0177\u017A\u017C\u017E-\u0180\u0183\u0185\u0188\u018C\u018D\u0192\u0195\u0199-\u019B\u019E\u01A1\u01A3\u01A5\u01A8\u01AA\u01AB\u01AD\u01B0\u01B4\u01B6\u01B9\u01BA\u01BD-\u01BF\u01C6\u01C9\u01CC\u01CE\u01D0\u01D2\u01D4\u01D6\u01D8\u01DA\u01DC\u01DD\u01DF\u01E1\u01E3\u01E5\u01E7\u01E9\u01EB\u01ED\u01EF\u01F0\u01F3\u01F5\u01F9\u01FB\u01FD\u01FF\u0201\u0203\u0205\u0207\u0209\u020B\u020D\u020F\u0211\u0213\u0215\u0217\u0219\u021B\u021D\u021F\u0221\u0223\u0225\u0227\u0229\u022B\u022D\u022F\u0231\u0233-\u0239\u023C\u023F\u0240\u0242\u0247\u0249\u024B\u024D\u024F-\u0293\u0295-\u02AF\u0371\u0373\u0377\u037B-\u037D\u0390\u03AC-\u03CE\u03D0\u03D1\u03D5-\u03D7\u03D9\u03DB\u03DD\u03DF\u03E1\u03E3\u03E5\u03E7\u03E9\u03EB\u03ED\u03EF-\u03F3\u03F5\u03F8\u03FB\u03FC\u0430-\u045F\u0461\u0463\u0465\u0467\u0469\u046B\u046D\u046F\u0471\u0473\u0475\u0477\u0479\u047B\u047D\u047F\u0481\u048B\u048D\u048F\u0491\u0493\u0495\u0497\u0499\u049B\u049D\u049F\u04A1\u04A3\u04A5\u04A7\u04A9\u04AB\u04AD\u04AF\u04B1\u04B3\u04B5\u04B7\u04B9\u04BB\u04BD\u04BF\u04C2\u04C4\u04C6\u04C8\u04CA\u04CC\u04CE\u04CF\u04D1\u04D3\u04D5\u04D7\u04D9\u04DB\u04DD\u04DF\u04E1\u04E3\u04E5\u04E7\u04E9\u04EB\u04ED\u04EF\u04F1\u04F3\u04F5\u04F7\u04F9\u04FB\u04FD\u04FF\u0501\u0503\u0505\u0507\u0509\u050B\u050D\u050F\u0511\u0513\u0515\u0517\u0519\u051B\u051D\u051F\u0521\u0523\u0525\u0527\u0529\u052B\u052D\u052F\u0560-\u0588\u10D0-\u10FA\u10FD-\u10FF\u13F8-\u13FD\u1C80-\u1C88\u1D00-\u1D2B\u1D6B-\u1D77\u1D79-\u1D9A\u1E01\u1E03\u1E05\u1E07\u1E09\u1E0B\u1E0D\u1E0F\u1E11\u1E13\u1E15\u1E17\u1E19\u1E1B\u1E1D\u1E1F\u1E21\u1E23\u1E25\u1E27\u1E29\u1E2B\u1E2D\u1E2F\u1E31\u1E33\u1E35\u1E37\u1E39\u1E3B\u1E3D\u1E3F\u1E41\u1E43\u1E45\u1E47\u1E49\u1E4B\u1E4D\u1E4F\u1E51\u1E53\u1E55\u1E57\u1E59\u1E5B\u1E5D\u1E5F\u1E61\u1E63\u1E65\u1E67\u1E69\u1E6B\u1E6D\u1E6F\u1E71\u1E73\u1E75\u1E77\u1E79\u1E7B\u1E7D\u1E7F\u1E81\u1E83\u1E85\u1E87\u1E89\u1E8B\u1E8D\u1E8F\u1E91\u1E93\u1E95-\u1E9D\u1E9F\u1EA1\u1EA3\u1EA5\u1EA7\u1EA9\u1EAB\u1EAD\u1EAF\u1EB1\u1EB3\u1EB5\u1EB7\u1EB9\u1EBB\u1EBD\u1EBF\u1EC1\u1EC3\u1EC5\u1EC7\u1EC9\u1ECB\u1ECD\u1ECF\u1ED1\u1ED3\u1ED5\u1ED7\u1ED9\u1EDB\u1EDD\u1EDF\u1EE1\u1EE3\u1EE5\u1EE7\u1EE9\u1EEB\u1EED\u1EEF\u1EF1\u1EF3\u1EF5\u1EF7\u1EF9\u1EFB\u1EFD\u1EFF-\u1F07\u1F10-\u1F15\u1F20-\u1F27\u1F30-\u1F37\u1F40-\u1F45\u1F50-\u1F57\u1F60-\u1F67\u1F70-\u1F7D\u1F80-\u1F87\u1F90-\u1F97\u1FA0-\u1FA7\u1FB0-\u1FB4\u1FB6\u1FB7\u1FBE\u1FC2-\u1FC4\u1FC6\u1FC7\u1FD0-\u1FD3\u1FD6\u1FD7\u1FE0-\u1FE7\u1FF2-\u1FF4\u1FF6\u1FF7\u210A\u210E\u210F\u2113\u212F\u2134\u2139\u213C\u213D\u2146-\u2149\u214E\u2184\u2C30-\u2C5E\u2C61\u2C65\u2C66\u2C68\u2C6A\u2C6C\u2C71\u2C73\u2C74\u2C76-\u2C7B\u2C81\u2C83\u2C85\u2C87\u2C89\u2C8B\u2C8D\u2C8F\u2C91\u2C93\u2C95\u2C97\u2C99\u2C9B\u2C9D\u2C9F\u2CA1\u2CA3\u2CA5\u2CA7\u2CA9\u2CAB\u2CAD\u2CAF\u2CB1\u2CB3\u2CB5\u2CB7\u2CB9\u2CBB\u2CBD\u2CBF\u2CC1\u2CC3\u2CC5\u2CC7\u2CC9\u2CCB\u2CCD\u2CCF\u2CD1\u2CD3\u2CD5\u2CD7\u2CD9\u2CDB\u2CDD\u2CDF\u2CE1\u2CE3\u2CE4\u2CEC\u2CEE\u2CF3\u2D00-\u2D25\u2D27\u2D2D\uA641\uA643\uA645\uA647\uA649\uA64B\uA64D\uA64F\uA651\uA653\uA655\uA657\uA659\uA65B\uA65D\uA65F\uA661\uA663\uA665\uA667\uA669\uA66B\uA66D\uA681\uA683\uA685\uA687\uA689\uA68B\uA68D\uA68F\uA691\uA693\uA695\uA697\uA699\uA69B\uA723\uA725\uA727\uA729\uA72B\uA72D\uA72F-\uA731\uA733\uA735\uA737\uA739\uA73B\uA73D\uA73F\uA741\uA743\uA745\uA747\uA749\uA74B\uA74D\uA74F\uA751\uA753\uA755\uA757\uA759\uA75B\uA75D\uA75F\uA761\uA763\uA765\uA767\uA769\uA76B\uA76D\uA76F\uA771-\uA778\uA77A\uA77C\uA77F\uA781\uA783\uA785\uA787\uA78C\uA78E\uA791\uA793-\uA795\uA797\uA799\uA79B\uA79D\uA79F\uA7A1\uA7A3\uA7A5\uA7A7\uA7A9\uA7AF\uA7B5\uA7B7\uA7B9\uA7BB\uA7BD\uA7BF\uA7C3\uA7FA\uAB30-\uAB5A\uAB60-\uAB67\uAB70-\uABBF\uFB00-\uFB06\uFB13-\uFB17\uFF41-\uFF5A]|\uD801[\uDC28-\uDC4F\uDCD8-\uDCFB]|\uD803[\uDCC0-\uDCF2]|\uD806[\uDCC0-\uDCDF]|\uD81B[\uDE60-\uDE7F]|\uD835[\uDC1A-\uDC33\uDC4E-\uDC54\uDC56-\uDC67\uDC82-\uDC9B\uDCB6-\uDCB9\uDCBB\uDCBD-\uDCC3\uDCC5-\uDCCF\uDCEA-\uDD03\uDD1E-\uDD37\uDD52-\uDD6B\uDD86-\uDD9F\uDDBA-\uDDD3\uDDEE-\uDE07\uDE22-\uDE3B\uDE56-\uDE6F\uDE8A-\uDEA5\uDEC2-\uDEDA\uDEDC-\uDEE1\uDEFC-\uDF14\uDF16-\uDF1B\uDF36-\uDF4E\uDF50-\uDF55\uDF70-\uDF88\uDF8A-\uDF8F\uDFAA-\uDFC2\uDFC4-\uDFC9\uDFCB]|\uD83A[\uDD22-\uDD43])/,
  surrogates: /[\uD800-\uDFFF]/,
  punctuation: /[!"'-\),-\/:;\?\[-\]_\{\}\xA1\xA7\xAB\xB6\xB7\xBB\xBF\u037E\u0387\u055A-\u055F\u0589\u058A\u05BE\u05C0\u05C3\u05C6\u05F3\u05F4\u0609\u060A\u060C\u060D\u061B\u061E\u061F\u066A-\u066D\u06D4\u0700-\u070D\u07F7-\u07F9\u0830-\u083E\u085E\u0964\u0965\u0970\u09FD\u0A76\u0AF0\u0C77\u0C84\u0DF4\u0E4F\u0E5A\u0E5B\u0F04-\u0F12\u0F14\u0F3A-\u0F3D\u0F85\u0FD0-\u0FD4\u0FD9\u0FDA\u104A-\u104F\u10FB\u1360-\u1368\u1400\u166E\u169B\u169C\u16EB-\u16ED\u1735\u1736\u17D4-\u17D6\u17D8-\u17DA\u1800-\u180A\u1944\u1945\u1A1E\u1A1F\u1AA0-\u1AA6\u1AA8-\u1AAD\u1B5A-\u1B60\u1BFC-\u1BFF\u1C3B-\u1C3F\u1C7E\u1C7F\u1CC0-\u1CC7\u1CD3\u2010-\u201F\u2022-\u2027\u2032-\u203A\u203C-\u2043\u2045-\u2051\u2053-\u205E\u207D\u207E\u208D\u208E\u2308-\u230B\u2329\u232A\u2768-\u2775\u27C5\u27C6\u27E6-\u27EF\u2983-\u2998\u29D8-\u29DB\u29FC\u29FD\u2CF9-\u2CFC\u2CFE\u2CFF\u2D70\u2E00-\u2E2E\u2E30-\u2E4F\u3001-\u3003\u3008-\u3011\u3014-\u301F\u3030\u303D\u30A0\u30FB\uA4FE\uA4FF\uA60D-\uA60F\uA673\uA67E\uA6F2-\uA6F7\uA874-\uA877\uA8CE\uA8CF\uA8F8-\uA8FA\uA8FC\uA92E\uA92F\uA95F\uA9C1-\uA9CD\uA9DE\uA9DF\uAA5C-\uAA5F\uAADE\uAADF\uAAF0\uAAF1\uABEB\uFD3E\uFD3F\uFE10-\uFE19\uFE30-\uFE52\uFE54-\uFE61\uFE63\uFE68\uFE6A\uFE6B\uFF01-\uFF03\uFF05-\uFF0A\uFF0C-\uFF0F\uFF1A\uFF1B\uFF1F\uFF20\uFF3B-\uFF3D\uFF3F\uFF5B\uFF5D\uFF5F-\uFF65]|\uD800[\uDD00-\uDD02\uDF9F\uDFD0]|\uD801\uDD6F|\uD802[\uDC57\uDD1F\uDD3F\uDE50-\uDE58\uDE7F\uDEF0-\uDEF6\uDF39-\uDF3F\uDF99-\uDF9C]|\uD803[\uDF55-\uDF59]|\uD804[\uDC47-\uDC4D\uDCBB\uDCBC\uDCBE-\uDCC1\uDD40-\uDD43\uDD74\uDD75\uDDC5-\uDDC8\uDDCD\uDDDB\uDDDD-\uDDDF\uDE38-\uDE3D\uDEA9]|\uD805[\uDC4B-\uDC4F\uDC5B\uDC5D\uDCC6\uDDC1-\uDDD7\uDE41-\uDE43\uDE60-\uDE6C\uDF3C-\uDF3E]|\uD806[\uDC3B\uDDE2\uDE3F-\uDE46\uDE9A-\uDE9C\uDE9E-\uDEA2]|\uD807[\uDC41-\uDC45\uDC70\uDC71\uDEF7\uDEF8\uDFFF]|\uD809[\uDC70-\uDC74]|\uD81A[\uDE6E\uDE6F\uDEF5\uDF37-\uDF3B\uDF44]|\uD81B[\uDE97-\uDE9A\uDFE2]|\uD82F\uDC9F|\uD836[\uDE87-\uDE8B]|\uD83A[\uDD5E\uDD5F]/,
  word: /[0-9A-Za-z\xAA\xB2\xB3\xB5\xB9\xBA\xBC-\xBE\xC0-\xD6\xD8-\xF6\xF8-\u02C1\u02C6-\u02D1\u02E0-\u02E4\u02EC\u02EE\u0300-\u0374\u0376\u0377\u037A-\u037D\u037F\u0386\u0388-\u038A\u038C\u038E-\u03A1\u03A3-\u03F5\u03F7-\u0481\u0483-\u052F\u0531-\u0556\u0559\u0560-\u0588\u0591-\u05BD\u05BF\u05C1\u05C2\u05C4\u05C5\u05C7\u05D0-\u05EA\u05EF-\u05F2\u0610-\u061A\u0620-\u0669\u066E-\u06D3\u06D5-\u06DC\u06DF-\u06E8\u06EA-\u06FC\u06FF\u0710-\u074A\u074D-\u07B1\u07C0-\u07F5\u07FA\u07FD\u0800-\u082D\u0840-\u085B\u0860-\u086A\u08A0-\u08B4\u08B6-\u08BD\u08D3-\u08E1\u08E3-\u0963\u0966-\u096F\u0971-\u0983\u0985-\u098C\u098F\u0990\u0993-\u09A8\u09AA-\u09B0\u09B2\u09B6-\u09B9\u09BC-\u09C4\u09C7\u09C8\u09CB-\u09CE\u09D7\u09DC\u09DD\u09DF-\u09E3\u09E6-\u09F1\u09F4-\u09F9\u09FC\u09FE\u0A01-\u0A03\u0A05-\u0A0A\u0A0F\u0A10\u0A13-\u0A28\u0A2A-\u0A30\u0A32\u0A33\u0A35\u0A36\u0A38\u0A39\u0A3C\u0A3E-\u0A42\u0A47\u0A48\u0A4B-\u0A4D\u0A51\u0A59-\u0A5C\u0A5E\u0A66-\u0A75\u0A81-\u0A83\u0A85-\u0A8D\u0A8F-\u0A91\u0A93-\u0AA8\u0AAA-\u0AB0\u0AB2\u0AB3\u0AB5-\u0AB9\u0ABC-\u0AC5\u0AC7-\u0AC9\u0ACB-\u0ACD\u0AD0\u0AE0-\u0AE3\u0AE6-\u0AEF\u0AF9-\u0AFF\u0B01-\u0B03\u0B05-\u0B0C\u0B0F\u0B10\u0B13-\u0B28\u0B2A-\u0B30\u0B32\u0B33\u0B35-\u0B39\u0B3C-\u0B44\u0B47\u0B48\u0B4B-\u0B4D\u0B56\u0B57\u0B5C\u0B5D\u0B5F-\u0B63\u0B66-\u0B6F\u0B71-\u0B77\u0B82\u0B83\u0B85-\u0B8A\u0B8E-\u0B90\u0B92-\u0B95\u0B99\u0B9A\u0B9C\u0B9E\u0B9F\u0BA3\u0BA4\u0BA8-\u0BAA\u0BAE-\u0BB9\u0BBE-\u0BC2\u0BC6-\u0BC8\u0BCA-\u0BCD\u0BD0\u0BD7\u0BE6-\u0BF2\u0C00-\u0C0C\u0C0E-\u0C10\u0C12-\u0C28\u0C2A-\u0C39\u0C3D-\u0C44\u0C46-\u0C48\u0C4A-\u0C4D\u0C55\u0C56\u0C58-\u0C5A\u0C60-\u0C63\u0C66-\u0C6F\u0C78-\u0C7E\u0C80-\u0C83\u0C85-\u0C8C\u0C8E-\u0C90\u0C92-\u0CA8\u0CAA-\u0CB3\u0CB5-\u0CB9\u0CBC-\u0CC4\u0CC6-\u0CC8\u0CCA-\u0CCD\u0CD5\u0CD6\u0CDE\u0CE0-\u0CE3\u0CE6-\u0CEF\u0CF1\u0CF2\u0D00-\u0D03\u0D05-\u0D0C\u0D0E-\u0D10\u0D12-\u0D44\u0D46-\u0D48\u0D4A-\u0D4E\u0D54-\u0D63\u0D66-\u0D78\u0D7A-\u0D7F\u0D82\u0D83\u0D85-\u0D96\u0D9A-\u0DB1\u0DB3-\u0DBB\u0DBD\u0DC0-\u0DC6\u0DCA\u0DCF-\u0DD4\u0DD6\u0DD8-\u0DDF\u0DE6-\u0DEF\u0DF2\u0DF3\u0E01-\u0E3A\u0E40-\u0E4E\u0E50-\u0E59\u0E81\u0E82\u0E84\u0E86-\u0E8A\u0E8C-\u0EA3\u0EA5\u0EA7-\u0EBD\u0EC0-\u0EC4\u0EC6\u0EC8-\u0ECD\u0ED0-\u0ED9\u0EDC-\u0EDF\u0F00\u0F18\u0F19\u0F20-\u0F33\u0F35\u0F37\u0F39\u0F3E-\u0F47\u0F49-\u0F6C\u0F71-\u0F84\u0F86-\u0F97\u0F99-\u0FBC\u0FC6\u1000-\u1049\u1050-\u109D\u10A0-\u10C5\u10C7\u10CD\u10D0-\u10FA\u10FC-\u1248\u124A-\u124D\u1250-\u1256\u1258\u125A-\u125D\u1260-\u1288\u128A-\u128D\u1290-\u12B0\u12B2-\u12B5\u12B8-\u12BE\u12C0\u12C2-\u12C5\u12C8-\u12D6\u12D8-\u1310\u1312-\u1315\u1318-\u135A\u135D-\u135F\u1369-\u137C\u1380-\u138F\u13A0-\u13F5\u13F8-\u13FD\u1401-\u166C\u166F-\u167F\u1681-\u169A\u16A0-\u16EA\u16EE-\u16F8\u1700-\u170C\u170E-\u1714\u1720-\u1734\u1740-\u1753\u1760-\u176C\u176E-\u1770\u1772\u1773\u1780-\u17D3\u17D7\u17DC\u17DD\u17E0-\u17E9\u17F0-\u17F9\u180B-\u180D\u1810-\u1819\u1820-\u1878\u1880-\u18AA\u18B0-\u18F5\u1900-\u191E\u1920-\u192B\u1930-\u193B\u1946-\u196D\u1970-\u1974\u1980-\u19AB\u19B0-\u19C9\u19D0-\u19DA\u1A00-\u1A1B\u1A20-\u1A5E\u1A60-\u1A7C\u1A7F-\u1A89\u1A90-\u1A99\u1AA7\u1AB0-\u1ABE\u1B00-\u1B4B\u1B50-\u1B59\u1B6B-\u1B73\u1B80-\u1BF3\u1C00-\u1C37\u1C40-\u1C49\u1C4D-\u1C7D\u1C80-\u1C88\u1C90-\u1CBA\u1CBD-\u1CBF\u1CD0-\u1CD2\u1CD4-\u1CFA\u1D00-\u1DF9\u1DFB-\u1F15\u1F18-\u1F1D\u1F20-\u1F45\u1F48-\u1F4D\u1F50-\u1F57\u1F59\u1F5B\u1F5D\u1F5F-\u1F7D\u1F80-\u1FB4\u1FB6-\u1FBC\u1FBE\u1FC2-\u1FC4\u1FC6-\u1FCC\u1FD0-\u1FD3\u1FD6-\u1FDB\u1FE0-\u1FEC\u1FF2-\u1FF4\u1FF6-\u1FFC\u2070\u2071\u2074-\u2079\u207F-\u2089\u2090-\u209C\u20D0-\u20F0\u2102\u2107\u210A-\u2113\u2115\u2119-\u211D\u2124\u2126\u2128\u212A-\u212D\u212F-\u2139\u213C-\u213F\u2145-\u2149\u214E\u2150-\u2189\u2460-\u249B\u24EA-\u24FF\u2776-\u2793\u2C00-\u2C2E\u2C30-\u2C5E\u2C60-\u2CE4\u2CEB-\u2CF3\u2CFD\u2D00-\u2D25\u2D27\u2D2D\u2D30-\u2D67\u2D6F\u2D7F-\u2D96\u2DA0-\u2DA6\u2DA8-\u2DAE\u2DB0-\u2DB6\u2DB8-\u2DBE\u2DC0-\u2DC6\u2DC8-\u2DCE\u2DD0-\u2DD6\u2DD8-\u2DDE\u2DE0-\u2DFF\u2E2F\u3005-\u3007\u3021-\u302F\u3031-\u3035\u3038-\u303C\u3041-\u3096\u3099\u309A\u309D-\u309F\u30A1-\u30FA\u30FC-\u30FF\u3105-\u312F\u3131-\u318E\u3192-\u3195\u31A0-\u31BA\u31F0-\u31FF\u3220-\u3229\u3248-\u324F\u3251-\u325F\u3280-\u3289\u32B1-\u32BF\u3400-\u4DB5\u4E00-\u9FEF\uA000-\uA48C\uA4D0-\uA4FD\uA500-\uA60C\uA610-\uA62B\uA640-\uA672\uA674-\uA67D\uA67F-\uA6F1\uA717-\uA71F\uA722-\uA788\uA78B-\uA7BF\uA7C2-\uA7C6\uA7F7-\uA827\uA830-\uA835\uA840-\uA873\uA880-\uA8C5\uA8D0-\uA8D9\uA8E0-\uA8F7\uA8FB\uA8FD-\uA92D\uA930-\uA953\uA960-\uA97C\uA980-\uA9C0\uA9CF-\uA9D9\uA9E0-\uA9FE\uAA00-\uAA36\uAA40-\uAA4D\uAA50-\uAA59\uAA60-\uAA76\uAA7A-\uAAC2\uAADB-\uAADD\uAAE0-\uAAEF\uAAF2-\uAAF6\uAB01-\uAB06\uAB09-\uAB0E\uAB11-\uAB16\uAB20-\uAB26\uAB28-\uAB2E\uAB30-\uAB5A\uAB5C-\uAB67\uAB70-\uABEA\uABEC\uABED\uABF0-\uABF9\uAC00-\uD7A3\uD7B0-\uD7C6\uD7CB-\uD7FB\uF900-\uFA6D\uFA70-\uFAD9\uFB00-\uFB06\uFB13-\uFB17\uFB1D-\uFB28\uFB2A-\uFB36\uFB38-\uFB3C\uFB3E\uFB40\uFB41\uFB43\uFB44\uFB46-\uFBB1\uFBD3-\uFD3D\uFD50-\uFD8F\uFD92-\uFDC7\uFDF0-\uFDFB\uFE00-\uFE0F\uFE20-\uFE2F\uFE70-\uFE74\uFE76-\uFEFC\uFF10-\uFF19\uFF21-\uFF3A\uFF41-\uFF5A\uFF66-\uFFBE\uFFC2-\uFFC7\uFFCA-\uFFCF\uFFD2-\uFFD7\uFFDA-\uFFDC]|\uD800[\uDC00-\uDC0B\uDC0D-\uDC26\uDC28-\uDC3A\uDC3C\uDC3D\uDC3F-\uDC4D\uDC50-\uDC5D\uDC80-\uDCFA\uDD07-\uDD33\uDD40-\uDD78\uDD8A\uDD8B\uDDFD\uDE80-\uDE9C\uDEA0-\uDED0\uDEE0-\uDEFB\uDF00-\uDF23\uDF2D-\uDF4A\uDF50-\uDF7A\uDF80-\uDF9D\uDFA0-\uDFC3\uDFC8-\uDFCF\uDFD1-\uDFD5]|\uD801[\uDC00-\uDC9D\uDCA0-\uDCA9\uDCB0-\uDCD3\uDCD8-\uDCFB\uDD00-\uDD27\uDD30-\uDD63\uDE00-\uDF36\uDF40-\uDF55\uDF60-\uDF67]|\uD802[\uDC00-\uDC05\uDC08\uDC0A-\uDC35\uDC37\uDC38\uDC3C\uDC3F-\uDC55\uDC58-\uDC76\uDC79-\uDC9E\uDCA7-\uDCAF\uDCE0-\uDCF2\uDCF4\uDCF5\uDCFB-\uDD1B\uDD20-\uDD39\uDD80-\uDDB7\uDDBC-\uDDCF\uDDD2-\uDE03\uDE05\uDE06\uDE0C-\uDE13\uDE15-\uDE17\uDE19-\uDE35\uDE38-\uDE3A\uDE3F-\uDE48\uDE60-\uDE7E\uDE80-\uDE9F\uDEC0-\uDEC7\uDEC9-\uDEE6\uDEEB-\uDEEF\uDF00-\uDF35\uDF40-\uDF55\uDF58-\uDF72\uDF78-\uDF91\uDFA9-\uDFAF]|\uD803[\uDC00-\uDC48\uDC80-\uDCB2\uDCC0-\uDCF2\uDCFA-\uDD27\uDD30-\uDD39\uDE60-\uDE7E\uDF00-\uDF27\uDF30-\uDF54\uDFE0-\uDFF6]|\uD804[\uDC00-\uDC46\uDC52-\uDC6F\uDC7F-\uDCBA\uDCD0-\uDCE8\uDCF0-\uDCF9\uDD00-\uDD34\uDD36-\uDD3F\uDD44-\uDD46\uDD50-\uDD73\uDD76\uDD80-\uDDC4\uDDC9-\uDDCC\uDDD0-\uDDDA\uDDDC\uDDE1-\uDDF4\uDE00-\uDE11\uDE13-\uDE37\uDE3E\uDE80-\uDE86\uDE88\uDE8A-\uDE8D\uDE8F-\uDE9D\uDE9F-\uDEA8\uDEB0-\uDEEA\uDEF0-\uDEF9\uDF00-\uDF03\uDF05-\uDF0C\uDF0F\uDF10\uDF13-\uDF28\uDF2A-\uDF30\uDF32\uDF33\uDF35-\uDF39\uDF3B-\uDF44\uDF47\uDF48\uDF4B-\uDF4D\uDF50\uDF57\uDF5D-\uDF63\uDF66-\uDF6C\uDF70-\uDF74]|\uD805[\uDC00-\uDC4A\uDC50-\uDC59\uDC5E\uDC5F\uDC80-\uDCC5\uDCC7\uDCD0-\uDCD9\uDD80-\uDDB5\uDDB8-\uDDC0\uDDD8-\uDDDD\uDE00-\uDE40\uDE44\uDE50-\uDE59\uDE80-\uDEB8\uDEC0-\uDEC9\uDF00-\uDF1A\uDF1D-\uDF2B\uDF30-\uDF3B]|\uD806[\uDC00-\uDC3A\uDCA0-\uDCF2\uDCFF\uDDA0-\uDDA7\uDDAA-\uDDD7\uDDDA-\uDDE1\uDDE3\uDDE4\uDE00-\uDE3E\uDE47\uDE50-\uDE99\uDE9D\uDEC0-\uDEF8]|\uD807[\uDC00-\uDC08\uDC0A-\uDC36\uDC38-\uDC40\uDC50-\uDC6C\uDC72-\uDC8F\uDC92-\uDCA7\uDCA9-\uDCB6\uDD00-\uDD06\uDD08\uDD09\uDD0B-\uDD36\uDD3A\uDD3C\uDD3D\uDD3F-\uDD47\uDD50-\uDD59\uDD60-\uDD65\uDD67\uDD68\uDD6A-\uDD8E\uDD90\uDD91\uDD93-\uDD98\uDDA0-\uDDA9\uDEE0-\uDEF6\uDFC0-\uDFD4]|\uD808[\uDC00-\uDF99]|\uD809[\uDC00-\uDC6E\uDC80-\uDD43]|[\uD80C\uD81C-\uD820\uD840-\uD868\uD86A-\uD86C\uD86F-\uD872\uD874-\uD879][\uDC00-\uDFFF]|\uD80D[\uDC00-\uDC2E]|\uD811[\uDC00-\uDE46]|\uD81A[\uDC00-\uDE38\uDE40-\uDE5E\uDE60-\uDE69\uDED0-\uDEED\uDEF0-\uDEF4\uDF00-\uDF36\uDF40-\uDF43\uDF50-\uDF59\uDF5B-\uDF61\uDF63-\uDF77\uDF7D-\uDF8F]|\uD81B[\uDE40-\uDE96\uDF00-\uDF4A\uDF4F-\uDF87\uDF8F-\uDF9F\uDFE0\uDFE1\uDFE3]|\uD821[\uDC00-\uDFF7]|\uD822[\uDC00-\uDEF2]|\uD82C[\uDC00-\uDD1E\uDD50-\uDD52\uDD64-\uDD67\uDD70-\uDEFB]|\uD82F[\uDC00-\uDC6A\uDC70-\uDC7C\uDC80-\uDC88\uDC90-\uDC99\uDC9D\uDC9E]|\uD834[\uDD65-\uDD69\uDD6D-\uDD72\uDD7B-\uDD82\uDD85-\uDD8B\uDDAA-\uDDAD\uDE42-\uDE44\uDEE0-\uDEF3\uDF60-\uDF78]|\uD835[\uDC00-\uDC54\uDC56-\uDC9C\uDC9E\uDC9F\uDCA2\uDCA5\uDCA6\uDCA9-\uDCAC\uDCAE-\uDCB9\uDCBB\uDCBD-\uDCC3\uDCC5-\uDD05\uDD07-\uDD0A\uDD0D-\uDD14\uDD16-\uDD1C\uDD1E-\uDD39\uDD3B-\uDD3E\uDD40-\uDD44\uDD46\uDD4A-\uDD50\uDD52-\uDEA5\uDEA8-\uDEC0\uDEC2-\uDEDA\uDEDC-\uDEFA\uDEFC-\uDF14\uDF16-\uDF34\uDF36-\uDF4E\uDF50-\uDF6E\uDF70-\uDF88\uDF8A-\uDFA8\uDFAA-\uDFC2\uDFC4-\uDFCB\uDFCE-\uDFFF]|\uD836[\uDE00-\uDE36\uDE3B-\uDE6C\uDE75\uDE84\uDE9B-\uDE9F\uDEA1-\uDEAF]|\uD838[\uDC00-\uDC06\uDC08-\uDC18\uDC1B-\uDC21\uDC23\uDC24\uDC26-\uDC2A\uDD00-\uDD2C\uDD30-\uDD3D\uDD40-\uDD49\uDD4E\uDEC0-\uDEF9]|\uD83A[\uDC00-\uDCC4\uDCC7-\uDCD6\uDD00-\uDD4B\uDD50-\uDD59]|\uD83B[\uDC71-\uDCAB\uDCAD-\uDCAF\uDCB1-\uDCB4\uDD01-\uDD2D\uDD2F-\uDD3D\uDE00-\uDE03\uDE05-\uDE1F\uDE21\uDE22\uDE24\uDE27\uDE29-\uDE32\uDE34-\uDE37\uDE39\uDE3B\uDE42\uDE47\uDE49\uDE4B\uDE4D-\uDE4F\uDE51\uDE52\uDE54\uDE57\uDE59\uDE5B\uDE5D\uDE5F\uDE61\uDE62\uDE64\uDE67-\uDE6A\uDE6C-\uDE72\uDE74-\uDE77\uDE79-\uDE7C\uDE7E\uDE80-\uDE89\uDE8B-\uDE9B\uDEA1-\uDEA3\uDEA5-\uDEA9\uDEAB-\uDEBB]|\uD83C[\uDD00-\uDD0C]|\uD869[\uDC00-\uDED6\uDF00-\uDFFF]|\uD86D[\uDC00-\uDF34\uDF40-\uDFFF]|\uD86E[\uDC00-\uDC1D\uDC20-\uDFFF]|\uD873[\uDC00-\uDEA1\uDEB0-\uDFFF]|\uD87A[\uDC00-\uDFE0]|\uD87E[\uDC00-\uDE1D]|\uDB40[\uDD00-\uDDEF]/,
  whiteSpace: /[\t-\r \x85\xA0\u1680\u2000-\u200A\u2028\u2029\u202F\u205F\u3000]/
};

var arrayIterate;
var hasRequiredArrayIterate;

function requireArrayIterate () {
	if (hasRequiredArrayIterate) return arrayIterate;
	hasRequiredArrayIterate = 1;

	arrayIterate = iterate;

	var own = {}.hasOwnProperty;

	function iterate(values, callback, context) {
	  var index = -1;
	  var result;

	  if (!values) {
	    throw new Error('Iterate requires that |this| not be ' + values)
	  }

	  if (!own.call(values, 'length')) {
	    throw new Error('Iterate requires that |this| has a `length`')
	  }

	  if (typeof callback !== 'function') {
	    throw new Error('`callback` must be a function')
	  }

	  // The length might change, so we do not cache it.
	  while (++index < values.length) {
	    // Skip missing values.
	    if (!(index in values)) {
	      continue
	    }

	    result = callback.call(context, values[index], index, values);

	    // If `callback` returns a `number`, move `index` over to `number`.
	    if (typeof result === 'number') {
	      // Make sure that negative numbers do not break the loop.
	      if (result < 0) {
	        index = 0;
	      }

	      index = result - 1;
	    }
	  }
	}
	return arrayIterate;
}

var unistUtilModifyChildren;
var hasRequiredUnistUtilModifyChildren;

function requireUnistUtilModifyChildren () {
	if (hasRequiredUnistUtilModifyChildren) return unistUtilModifyChildren;
	hasRequiredUnistUtilModifyChildren = 1;

	var iterate = requireArrayIterate();

	unistUtilModifyChildren = modifierFactory;

	// Turn `callback` into a child-modifier accepting a parent.  See
	// `array-iterate` for more info.
	function modifierFactory(callback) {
	  return iteratorFactory(wrapperFactory(callback))
	}

	// Turn `callback` into a `iterator' accepting a parent.
	function iteratorFactory(callback) {
	  return iterator

	  function iterator(parent) {
	    var children = parent && parent.children;

	    if (!children) {
	      throw new Error('Missing children in `parent` for `modifier`')
	    }

	    return iterate(children, callback, parent)
	  }
	}

	// Pass the context as the third argument to `callback`.
	function wrapperFactory(callback) {
	  return wrapper

	  function wrapper(value, index) {
	    return callback(value, index, this)
	  }
	}
	return unistUtilModifyChildren;
}

var mergeInitialWordSymbol_1;
var hasRequiredMergeInitialWordSymbol;

function requireMergeInitialWordSymbol () {
	if (hasRequiredMergeInitialWordSymbol) return mergeInitialWordSymbol_1;
	hasRequiredMergeInitialWordSymbol = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();

	mergeInitialWordSymbol_1 = modifyChildren(mergeInitialWordSymbol);

	// Merge certain punctuation marks into their following words.
	function mergeInitialWordSymbol(child, index, parent) {
	  var children;
	  var next;

	  if (
	    (child.type !== 'SymbolNode' && child.type !== 'PunctuationNode') ||
	    toString(child) !== '&'
	  ) {
	    return
	  }

	  children = parent.children;

	  next = children[index + 1];

	  // If either a previous word, or no following word, exists, exit early.
	  if (
	    (index !== 0 && children[index - 1].type === 'WordNode') ||
	    !(next && next.type === 'WordNode')
	  ) {
	    return
	  }

	  // Remove `child` from parent.
	  children.splice(index, 1);

	  // Add the punctuation mark at the start of the next node.
	  next.children.unshift(child);

	  // Update position.
	  if (next.position && child.position) {
	    next.position.start = child.position.start;
	  }

	  // Next, iterate over the node at the previous position, as it's now adjacent
	  // to a following word.
	  return index - 1
	}
	return mergeInitialWordSymbol_1;
}

var mergeFinalWordSymbol_1;
var hasRequiredMergeFinalWordSymbol;

function requireMergeFinalWordSymbol () {
	if (hasRequiredMergeFinalWordSymbol) return mergeFinalWordSymbol_1;
	hasRequiredMergeFinalWordSymbol = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();

	mergeFinalWordSymbol_1 = modifyChildren(mergeFinalWordSymbol);

	// Merge certain punctuation marks into their preceding words.
	function mergeFinalWordSymbol(child, index, parent) {
	  var children;
	  var prev;
	  var next;

	  if (
	    index !== 0 &&
	    (child.type === 'SymbolNode' || child.type === 'PunctuationNode') &&
	    toString(child) === '-'
	  ) {
	    children = parent.children;

	    prev = children[index - 1];
	    next = children[index + 1];

	    if (
	      (!next || next.type !== 'WordNode') &&
	      prev &&
	      prev.type === 'WordNode'
	    ) {
	      // Remove `child` from parent.
	      children.splice(index, 1);

	      // Add the punctuation mark at the end of the previous node.
	      prev.children.push(child);

	      // Update position.
	      if (prev.position && child.position) {
	        prev.position.end = child.position.end;
	      }

	      // Next, iterate over the node *now* at the current position (which was
	      // the next node).
	      return index
	    }
	  }
	}
	return mergeFinalWordSymbol_1;
}

var mergeInnerWordSymbol_1;
var hasRequiredMergeInnerWordSymbol;

function requireMergeInnerWordSymbol () {
	if (hasRequiredMergeInnerWordSymbol) return mergeInnerWordSymbol_1;
	hasRequiredMergeInnerWordSymbol = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();
	var expressions = expressions$1;

	mergeInnerWordSymbol_1 = modifyChildren(mergeInnerWordSymbol);

	// Symbols part of surrounding words.
	var wordSymbolInner = expressions.wordSymbolInner;

	// Merge words joined by certain punctuation marks.
	function mergeInnerWordSymbol(child, index, parent) {
	  var siblings;
	  var sibling;
	  var prev;
	  var last;
	  var position;
	  var tokens;
	  var queue;

	  if (
	    index !== 0 &&
	    (child.type === 'SymbolNode' || child.type === 'PunctuationNode')
	  ) {
	    siblings = parent.children;
	    prev = siblings[index - 1];

	    if (prev && prev.type === 'WordNode') {
	      position = index - 1;

	      tokens = [];
	      queue = [];

	      // -   If a token which is neither word nor inner word symbol is found,
	      //     the loop is broken
	      // -   If an inner word symbol is found,  it’s queued
	      // -   If a word is found, it’s queued (and the queue stored and emptied)
	      while (siblings[++position]) {
	        sibling = siblings[position];

	        if (sibling.type === 'WordNode') {
	          tokens = tokens.concat(queue, sibling.children);

	          queue = [];
	        } else if (
	          (sibling.type === 'SymbolNode' ||
	            sibling.type === 'PunctuationNode') &&
	          wordSymbolInner.test(toString(sibling))
	        ) {
	          queue.push(sibling);
	        } else {
	          break
	        }
	      }

	      if (tokens.length !== 0) {
	        // If there is a queue, remove its length from `position`.
	        if (queue.length !== 0) {
	          position -= queue.length;
	        }

	        // Remove every (one or more) inner-word punctuation marks and children
	        // of words.
	        siblings.splice(index, position - index);

	        // Add all found tokens to `prev`s children.
	        prev.children = prev.children.concat(tokens);

	        last = tokens[tokens.length - 1];

	        // Update position.
	        if (prev.position && last.position) {
	          prev.position.end = last.position.end;
	        }

	        // Next, iterate over the node *now* at the current position.
	        return index
	      }
	    }
	  }
	}
	return mergeInnerWordSymbol_1;
}

var mergeInnerWordSlash_1;
var hasRequiredMergeInnerWordSlash;

function requireMergeInnerWordSlash () {
	if (hasRequiredMergeInnerWordSlash) return mergeInnerWordSlash_1;
	hasRequiredMergeInnerWordSlash = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();

	mergeInnerWordSlash_1 = modifyChildren(mergeInnerWordSlash);

	var slash = '/';

	// Merge words joined by certain punctuation marks.
	function mergeInnerWordSlash(child, index, parent) {
	  var siblings = parent.children;
	  var prev;
	  var next;
	  var prevValue;
	  var nextValue;
	  var queue;
	  var tail;
	  var count;

	  prev = siblings[index - 1];
	  next = siblings[index + 1];

	  if (
	    prev &&
	    prev.type === 'WordNode' &&
	    (child.type === 'SymbolNode' || child.type === 'PunctuationNode') &&
	    toString(child) === slash
	  ) {
	    prevValue = toString(prev);
	    tail = child;
	    queue = [child];
	    count = 1;

	    if (next && next.type === 'WordNode') {
	      nextValue = toString(next);
	      tail = next;
	      queue = queue.concat(next.children);
	      count++;
	    }

	    if (prevValue.length < 3 && (!nextValue || nextValue.length < 3)) {
	      // Add all found tokens to `prev`s children.
	      prev.children = prev.children.concat(queue);

	      siblings.splice(index, count);

	      // Update position.
	      if (prev.position && tail.position) {
	        prev.position.end = tail.position.end;
	      }

	      // Next, iterate over the node *now* at the current position.
	      return index
	    }
	  }
	}
	return mergeInnerWordSlash_1;
}

var mergeInitialisms_1;
var hasRequiredMergeInitialisms;

function requireMergeInitialisms () {
	if (hasRequiredMergeInitialisms) return mergeInitialisms_1;
	hasRequiredMergeInitialisms = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();
	var expressions = expressions$1;

	mergeInitialisms_1 = modifyChildren(mergeInitialisms);

	var numerical = expressions.numerical;

	// Merge initialisms.
	function mergeInitialisms(child, index, parent) {
	  var siblings;
	  var prev;
	  var children;
	  var length;
	  var position;
	  var otherChild;
	  var isAllDigits;
	  var value;

	  if (index !== 0 && toString(child) === '.') {
	    siblings = parent.children;

	    prev = siblings[index - 1];
	    children = prev.children;

	    length = children && children.length;

	    if (prev.type === 'WordNode' && length !== 1 && length % 2 !== 0) {
	      position = length;

	      isAllDigits = true;

	      while (children[--position]) {
	        otherChild = children[position];

	        value = toString(otherChild);

	        if (position % 2 === 0) {
	          // Initialisms consist of one character values.
	          if (value.length > 1) {
	            return
	          }

	          if (!numerical.test(value)) {
	            isAllDigits = false;
	          }
	        } else if (value !== '.') {
	          if (position < length - 2) {
	            break
	          } else {
	            return
	          }
	        }
	      }

	      if (!isAllDigits) {
	        // Remove `child` from parent.
	        siblings.splice(index, 1);

	        // Add child to the previous children.
	        children.push(child);

	        // Update position.
	        if (prev.position && child.position) {
	          prev.position.end = child.position.end;
	        }

	        // Next, iterate over the node *now* at the current position.
	        return index
	      }
	    }
	  }
	}
	return mergeInitialisms_1;
}

var mergeWords;
var hasRequiredMergeWords;

function requireMergeWords () {
	if (hasRequiredMergeWords) return mergeWords;
	hasRequiredMergeWords = 1;

	var modifyChildren = requireUnistUtilModifyChildren();

	mergeWords = modifyChildren(mergeFinalWordSymbol);

	// Merge multiple words. This merges the children of adjacent words, something
	// which should not occur naturally by parse-latin, but might happen when custom
	// tokens were passed in.
	function mergeFinalWordSymbol(child, index, parent) {
	  var siblings = parent.children;
	  var next;

	  if (child.type === 'WordNode') {
	    next = siblings[index + 1];

	    if (next && next.type === 'WordNode') {
	      // Remove `next` from parent.
	      siblings.splice(index + 1, 1);

	      // Add the punctuation mark at the end of the previous node.
	      child.children = child.children.concat(next.children);

	      // Update position.
	      if (next.position && child.position) {
	        child.position.end = next.position.end;
	      }

	      // Next, re-iterate the current node.
	      return index
	    }
	  }
	}
	return mergeWords;
}

var unistUtilVisitChildren;
var hasRequiredUnistUtilVisitChildren;

function requireUnistUtilVisitChildren () {
	if (hasRequiredUnistUtilVisitChildren) return unistUtilVisitChildren;
	hasRequiredUnistUtilVisitChildren = 1;

	unistUtilVisitChildren = visitChildren;

	function visitChildren(callback) {
	  return visitor

	  // Visit `parent`, invoking `callback` for each child.
	  function visitor(parent) {
	    var index = -1;
	    var children = parent && parent.children;

	    if (!children) {
	      throw new Error('Missing children in `parent` for `visitor`')
	    }

	    while (++index in children) {
	      callback(children[index], index, parent);
	    }
	  }
	}
	return unistUtilVisitChildren;
}

var patchPosition_1;
var hasRequiredPatchPosition;

function requirePatchPosition () {
	if (hasRequiredPatchPosition) return patchPosition_1;
	hasRequiredPatchPosition = 1;

	var visitChildren = requireUnistUtilVisitChildren();

	patchPosition_1 = visitChildren(patchPosition);

	// Patch the position on a parent node based on its first and last child.
	function patchPosition(child, index, node) {
	  var siblings = node.children;

	  if (!child.position) {
	    return
	  }

	  if (
	    index === 0 &&
	    (!node.position || /* istanbul ignore next */ !node.position.start)
	  ) {
	    patch(node);
	    node.position.start = child.position.start;
	  }

	  if (index === siblings.length - 1 && (!node.position || !node.position.end)) {
	    patch(node);
	    node.position.end = child.position.end;
	  }
	}

	// Add a `position` object when it does not yet exist on `node`.
	function patch(node) {
	  if (!node.position) {
	    node.position = {};
	  }
	}
	return patchPosition_1;
}

var mergeNonWordSentences_1;
var hasRequiredMergeNonWordSentences;

function requireMergeNonWordSentences () {
	if (hasRequiredMergeNonWordSentences) return mergeNonWordSentences_1;
	hasRequiredMergeNonWordSentences = 1;

	var modifyChildren = requireUnistUtilModifyChildren();

	mergeNonWordSentences_1 = modifyChildren(mergeNonWordSentences);

	// Merge a sentence into the following sentence, when the sentence does not
	// contain word tokens.
	function mergeNonWordSentences(child, index, parent) {
	  var children = child.children;
	  var position = -1;
	  var prev;
	  var next;

	  while (children[++position]) {
	    if (children[position].type === 'WordNode') {
	      return
	    }
	  }

	  prev = parent.children[index - 1];

	  if (prev) {
	    prev.children = prev.children.concat(children);

	    // Remove the child.
	    parent.children.splice(index, 1);

	    // Patch position.
	    if (prev.position && child.position) {
	      prev.position.end = child.position.end;
	    }

	    // Next, iterate over the node *now* at the current position (which was the
	    // next node).
	    return index
	  }

	  next = parent.children[index + 1];

	  if (next) {
	    next.children = children.concat(next.children);

	    // Patch position.
	    if (next.position && child.position) {
	      next.position.start = child.position.start;
	    }

	    // Remove the child.
	    parent.children.splice(index, 1);
	  }
	}
	return mergeNonWordSentences_1;
}

var mergeAffixSymbol_1;
var hasRequiredMergeAffixSymbol;

function requireMergeAffixSymbol () {
	if (hasRequiredMergeAffixSymbol) return mergeAffixSymbol_1;
	hasRequiredMergeAffixSymbol = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();
	var expressions = expressions$1;

	mergeAffixSymbol_1 = modifyChildren(mergeAffixSymbol);

	// Closing or final punctuation, or terminal markers that should still be
	// included in the previous sentence, even though they follow the sentence’s
	// terminal marker.
	var affixSymbol = expressions.affixSymbol;

	// Move certain punctuation following a terminal marker (thus in the next
	// sentence) to the previous sentence.
	function mergeAffixSymbol(child, index, parent) {
	  var children = child.children;
	  var first;
	  var second;
	  var prev;

	  if (children && children.length !== 0 && index !== 0) {
	    first = children[0];
	    second = children[1];
	    prev = parent.children[index - 1];

	    if (
	      (first.type === 'SymbolNode' || first.type === 'PunctuationNode') &&
	      affixSymbol.test(toString(first))
	    ) {
	      prev.children.push(children.shift());

	      // Update position.
	      if (first.position && prev.position) {
	        prev.position.end = first.position.end;
	      }

	      if (second && second.position && child.position) {
	        child.position.start = second.position.start;
	      }

	      // Next, iterate over the previous node again.
	      return index - 1
	    }
	  }
	}
	return mergeAffixSymbol_1;
}

var mergeInitialLowerCaseLetterSentences_1;
var hasRequiredMergeInitialLowerCaseLetterSentences;

function requireMergeInitialLowerCaseLetterSentences () {
	if (hasRequiredMergeInitialLowerCaseLetterSentences) return mergeInitialLowerCaseLetterSentences_1;
	hasRequiredMergeInitialLowerCaseLetterSentences = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();
	var expressions = expressions$1;

	mergeInitialLowerCaseLetterSentences_1 = modifyChildren(mergeInitialLowerCaseLetterSentences);

	// Initial lowercase letter.
	var lowerInitial = expressions.lowerInitial;

	// Merge a sentence into its previous sentence, when the sentence starts with a
	// lower case letter.
	function mergeInitialLowerCaseLetterSentences(child, index, parent) {
	  var children = child.children;
	  var position;
	  var node;
	  var siblings;
	  var prev;

	  if (children && children.length !== 0 && index !== 0) {
	    position = -1;

	    while (children[++position]) {
	      node = children[position];

	      if (node.type === 'WordNode') {
	        if (!lowerInitial.test(toString(node))) {
	          return
	        }

	        siblings = parent.children;

	        prev = siblings[index - 1];

	        prev.children = prev.children.concat(children);

	        siblings.splice(index, 1);

	        // Update position.
	        if (prev.position && child.position) {
	          prev.position.end = child.position.end;
	        }

	        // Next, iterate over the node *now* at the current position.
	        return index
	      }

	      if (node.type === 'SymbolNode' || node.type === 'PunctuationNode') {
	        return
	      }
	    }
	  }
	}
	return mergeInitialLowerCaseLetterSentences_1;
}

var mergeInitialDigitSentences_1;
var hasRequiredMergeInitialDigitSentences;

function requireMergeInitialDigitSentences () {
	if (hasRequiredMergeInitialDigitSentences) return mergeInitialDigitSentences_1;
	hasRequiredMergeInitialDigitSentences = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();
	var expressions = expressions$1;

	mergeInitialDigitSentences_1 = modifyChildren(mergeInitialDigitSentences);

	// Initial lowercase letter.
	var digit = expressions.digitStart;

	// Merge a sentence into its previous sentence, when the sentence starts with a
	// lower case letter.
	function mergeInitialDigitSentences(child, index, parent) {
	  var children = child.children;
	  var siblings = parent.children;
	  var prev = siblings[index - 1];
	  var head = children[0];

	  if (prev && head && head.type === 'WordNode' && digit.test(toString(head))) {
	    prev.children = prev.children.concat(children);
	    siblings.splice(index, 1);

	    // Update position.
	    if (prev.position && child.position) {
	      prev.position.end = child.position.end;
	    }

	    // Next, iterate over the node *now* at the current position.
	    return index
	  }
	}
	return mergeInitialDigitSentences_1;
}

var mergePrefixExceptions_1;
var hasRequiredMergePrefixExceptions;

function requireMergePrefixExceptions () {
	if (hasRequiredMergePrefixExceptions) return mergePrefixExceptions_1;
	hasRequiredMergePrefixExceptions = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();

	mergePrefixExceptions_1 = modifyChildren(mergePrefixExceptions);

	// Blacklist of full stop characters that should not be treated as terminal
	// sentence markers: A case-insensitive abbreviation.
	var abbreviationPrefix = new RegExp(
	  '^(' +
	    '[0-9]{1,3}|' +
	    '[a-z]|' +
	    // Common Latin Abbreviations:
	    // Based on: <https://en.wikipedia.org/wiki/List_of_Latin_abbreviations>.
	    // Where only the abbreviations written without joining full stops,
	    // but with a final full stop, were extracted.
	    //
	    // circa, capitulus, confer, compare, centum weight, eadem, (et) alii,
	    // et cetera, floruit, foliis, ibidem, idem, nemine && contradicente,
	    // opere && citato, (per) cent, (per) procurationem, (pro) tempore,
	    // sic erat scriptum, (et) sequentia, statim, videlicet. */
	    'al|ca|cap|cca|cent|cf|cit|con|cp|cwt|ead|etc|ff|' +
	    'fl|ibid|id|nem|op|pro|seq|sic|stat|tem|viz' +
	    ')$'
	);

	// Merge a sentence into its next sentence, when the sentence ends with a
	// certain word.
	function mergePrefixExceptions(child, index, parent) {
	  var children = child.children;
	  var period;
	  var node;
	  var next;

	  if (children && children.length > 1) {
	    period = children[children.length - 1];

	    if (period && toString(period) === '.') {
	      node = children[children.length - 2];

	      if (
	        node &&
	        node.type === 'WordNode' &&
	        abbreviationPrefix.test(toString(node).toLowerCase())
	      ) {
	        // Merge period into abbreviation.
	        node.children.push(period);
	        children.pop();

	        // Update position.
	        if (period.position && node.position) {
	          node.position.end = period.position.end;
	        }

	        // Merge sentences.
	        next = parent.children[index + 1];

	        if (next) {
	          child.children = children.concat(next.children);

	          parent.children.splice(index + 1, 1);

	          // Update position.
	          if (next.position && child.position) {
	            child.position.end = next.position.end;
	          }

	          // Next, iterate over the current node again.
	          return index - 1
	        }
	      }
	    }
	  }
	}
	return mergePrefixExceptions_1;
}

var mergeAffixExceptions_1;
var hasRequiredMergeAffixExceptions;

function requireMergeAffixExceptions () {
	if (hasRequiredMergeAffixExceptions) return mergeAffixExceptions_1;
	hasRequiredMergeAffixExceptions = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();

	mergeAffixExceptions_1 = modifyChildren(mergeAffixExceptions);

	// Merge a sentence into its previous sentence, when the sentence starts with a
	// comma.
	function mergeAffixExceptions(child, index, parent) {
	  var children = child.children;
	  var node;
	  var position;
	  var value;
	  var previousChild;

	  if (!children || children.length === 0 || index === 0) {
	    return
	  }

	  position = -1;

	  while (children[++position]) {
	    node = children[position];

	    if (node.type === 'WordNode') {
	      return
	    }

	    if (node.type === 'SymbolNode' || node.type === 'PunctuationNode') {
	      value = toString(node);

	      if (value !== ',' && value !== ';') {
	        return
	      }

	      previousChild = parent.children[index - 1];

	      previousChild.children = previousChild.children.concat(children);

	      // Update position.
	      if (previousChild.position && child.position) {
	        previousChild.position.end = child.position.end;
	      }

	      parent.children.splice(index, 1);

	      // Next, iterate over the node *now* at the current position.
	      return index
	    }
	  }
	}
	return mergeAffixExceptions_1;
}

var mergeRemainingFullStops_1;
var hasRequiredMergeRemainingFullStops;

function requireMergeRemainingFullStops () {
	if (hasRequiredMergeRemainingFullStops) return mergeRemainingFullStops_1;
	hasRequiredMergeRemainingFullStops = 1;

	var toString = nlcstToString_1;
	var visitChildren = requireUnistUtilVisitChildren();
	var expressions = expressions$1;

	mergeRemainingFullStops_1 = visitChildren(mergeRemainingFullStops);

	// Blacklist of full stop characters that should not be treated as terminal
	// sentence markers: A case-insensitive abbreviation.
	var terminalMarker = expressions.terminalMarker;

	// Merge non-terminal-marker full stops into the previous word (if available),
	// or the next word (if available).
	function mergeRemainingFullStops(child) {
	  var children = child.children;
	  var position = children.length;
	  var hasFoundDelimiter = false;
	  var grandchild;
	  var prev;
	  var next;
	  var nextNext;

	  while (children[--position]) {
	    grandchild = children[position];

	    if (
	      grandchild.type !== 'SymbolNode' &&
	      grandchild.type !== 'PunctuationNode'
	    ) {
	      // This is a sentence without terminal marker, so we 'fool' the code to
	      // make it think we have found one.
	      if (grandchild.type === 'WordNode') {
	        hasFoundDelimiter = true;
	      }

	      continue
	    }

	    // Exit when this token is not a terminal marker.
	    if (!terminalMarker.test(toString(grandchild))) {
	      continue
	    }

	    // Ignore the first terminal marker found (starting at the end), as it
	    // should not be merged.
	    if (!hasFoundDelimiter) {
	      hasFoundDelimiter = true;

	      continue
	    }

	    // Only merge a single full stop.
	    if (toString(grandchild) !== '.') {
	      continue
	    }

	    prev = children[position - 1];
	    next = children[position + 1];

	    if (prev && prev.type === 'WordNode') {
	      nextNext = children[position + 2];

	      // Continue when the full stop is followed by a space and another full
	      // stop, such as: `{.} .`
	      if (
	        next &&
	        nextNext &&
	        next.type === 'WhiteSpaceNode' &&
	        toString(nextNext) === '.'
	      ) {
	        continue
	      }

	      // Remove `child` from parent.
	      children.splice(position, 1);

	      // Add the punctuation mark at the end of the previous node.
	      prev.children.push(grandchild);

	      // Update position.
	      if (grandchild.position && prev.position) {
	        prev.position.end = grandchild.position.end;
	      }

	      position--;
	    } else if (next && next.type === 'WordNode') {
	      // Remove `child` from parent.
	      children.splice(position, 1);

	      // Add the punctuation mark at the start of the next node.
	      next.children.unshift(grandchild);

	      if (grandchild.position && next.position) {
	        next.position.start = grandchild.position.start;
	      }
	    }
	  }
	}
	return mergeRemainingFullStops_1;
}

var makeInitialWhiteSpaceSiblings_1;
var hasRequiredMakeInitialWhiteSpaceSiblings;

function requireMakeInitialWhiteSpaceSiblings () {
	if (hasRequiredMakeInitialWhiteSpaceSiblings) return makeInitialWhiteSpaceSiblings_1;
	hasRequiredMakeInitialWhiteSpaceSiblings = 1;

	var visitChildren = requireUnistUtilVisitChildren();

	makeInitialWhiteSpaceSiblings_1 = visitChildren(makeInitialWhiteSpaceSiblings);

	// Move white space starting a sentence up, so they are the siblings of
	// sentences.
	function makeInitialWhiteSpaceSiblings(child, index, parent) {
	  var children = child.children;
	  var next;

	  if (
	    children &&
	    children.length !== 0 &&
	    children[0].type === 'WhiteSpaceNode'
	  ) {
	    parent.children.splice(index, 0, children.shift());
	    next = children[0];

	    if (next && next.position && child.position) {
	      child.position.start = next.position.start;
	    }
	  }
	}
	return makeInitialWhiteSpaceSiblings_1;
}

var makeFinalWhiteSpaceSiblings_1;
var hasRequiredMakeFinalWhiteSpaceSiblings;

function requireMakeFinalWhiteSpaceSiblings () {
	if (hasRequiredMakeFinalWhiteSpaceSiblings) return makeFinalWhiteSpaceSiblings_1;
	hasRequiredMakeFinalWhiteSpaceSiblings = 1;

	var modifyChildren = requireUnistUtilModifyChildren();

	makeFinalWhiteSpaceSiblings_1 = modifyChildren(makeFinalWhiteSpaceSiblings);

	// Move white space ending a paragraph up, so they are the siblings of
	// paragraphs.
	function makeFinalWhiteSpaceSiblings(child, index, parent) {
	  var children = child.children;
	  var prev;

	  if (
	    children &&
	    children.length !== 0 &&
	    children[children.length - 1].type === 'WhiteSpaceNode'
	  ) {
	    parent.children.splice(index + 1, 0, child.children.pop());
	    prev = children[children.length - 1];

	    if (prev && prev.position && child.position) {
	      child.position.end = prev.position.end;
	    }

	    // Next, iterate over the current node again.
	    return index
	  }
	}
	return makeFinalWhiteSpaceSiblings_1;
}

var breakImplicitSentences_1;
var hasRequiredBreakImplicitSentences;

function requireBreakImplicitSentences () {
	if (hasRequiredBreakImplicitSentences) return breakImplicitSentences_1;
	hasRequiredBreakImplicitSentences = 1;

	var toString = nlcstToString_1;
	var modifyChildren = requireUnistUtilModifyChildren();
	var expressions = expressions$1;

	breakImplicitSentences_1 = modifyChildren(breakImplicitSentences);

	// Two or more new line characters.
	var multiNewLine = expressions.newLineMulti;

	// Break a sentence if a white space with more than one new-line is found.
	function breakImplicitSentences(child, index, parent) {
	  var children;
	  var position;
	  var length;
	  var tail;
	  var head;
	  var end;
	  var insertion;
	  var node;

	  if (child.type !== 'SentenceNode') {
	    return
	  }

	  children = child.children;

	  // Ignore first and last child.
	  length = children.length - 1;
	  position = 0;

	  while (++position < length) {
	    node = children[position];

	    if (node.type !== 'WhiteSpaceNode' || !multiNewLine.test(toString(node))) {
	      continue
	    }

	    child.children = children.slice(0, position);

	    insertion = {
	      type: 'SentenceNode',
	      children: children.slice(position + 1)
	    };

	    tail = children[position - 1];
	    head = children[position + 1];

	    parent.children.splice(index + 1, 0, node, insertion);

	    if (child.position && tail.position && head.position) {
	      end = child.position.end;

	      child.position.end = tail.position.end;

	      insertion.position = {
	        start: head.position.start,
	        end: end
	      };
	    }

	    return index + 1
	  }
	}
	return breakImplicitSentences_1;
}

var removeEmptyNodes_1;
var hasRequiredRemoveEmptyNodes;

function requireRemoveEmptyNodes () {
	if (hasRequiredRemoveEmptyNodes) return removeEmptyNodes_1;
	hasRequiredRemoveEmptyNodes = 1;

	var modifyChildren = requireUnistUtilModifyChildren();

	removeEmptyNodes_1 = modifyChildren(removeEmptyNodes);

	// Remove empty children.
	function removeEmptyNodes(child, index, parent) {
	  if ('children' in child && child.children.length === 0) {
	    parent.children.splice(index, 1);

	    // Next, iterate over the node *now* at the current position (which was the
	    // next node).
	    return index
	  }
	}
	return removeEmptyNodes_1;
}

var createParser = parser;
var expressions = expressions$1;

var lib = ParseLatin;

// PARSE LATIN

// Transform Latin-script natural language into an NLCST-tree.
function ParseLatin(doc, file) {
  var value = file || doc;

  if (!(this instanceof ParseLatin)) {
    return new ParseLatin(doc, file)
  }

  this.doc = value ? String(value) : null;
}

// Quick access to the prototype.
var proto = ParseLatin.prototype;

// Default position.
proto.position = true;

// Create text nodes.
proto.tokenizeSymbol = createTextFactory('Symbol');
proto.tokenizeWhiteSpace = createTextFactory('WhiteSpace');
proto.tokenizePunctuation = createTextFactory('Punctuation');
proto.tokenizeSource = createTextFactory('Source');
proto.tokenizeText = createTextFactory('Text');

// Expose `run`.
proto.run = run;

// Inject `plugins` to modifiy the result of the method at `key` on the operated
// on context.
proto.use = useFactory(function(context, key, plugins) {
  context[key] = context[key].concat(plugins);
});

// Inject `plugins` to modifiy the result of the method at `key` on the operated
// on context, before any other.
proto.useFirst = useFactory(function(context, key, plugins) {
  context[key] = plugins.concat(context[key]);
});

// Easy access to the document parser. This additionally supports retext-style
// invocation: where an instance is created for each file, and the file is given
// on construction.
proto.parse = function(value) {
  return this.tokenizeRoot(value || this.doc)
};

// Transform a `value` into a list of `NLCSTNode`s.
proto.tokenize = function(value) {
  return tokenize(this, value)
};

// PARENT NODES
//
// All these nodes are `pluggable`: they come with a `use` method which accepts
// a plugin (`function(NLCSTNode)`).
// Every time one of these methods are called, the plugin is invoked with the
// node, allowing for easy modification.
//
// In fact, the internal transformation from `tokenize` (a list of words, white
// space, punctuation, and symbols) to `tokenizeRoot` (an NLCST tree), is also
// implemented through this mechanism.

// Create a `WordNode` with its children set to a single `TextNode`, its value
// set to the given `value`.
pluggable(ParseLatin, 'tokenizeWord', function(value, eat) {
  var add = (eat || noopEat)('');
  var parent = {type: 'WordNode', children: []};

  this.tokenizeText(value, eat, parent);

  return add(parent)
});

// Create a `SentenceNode` with its children set to `Node`s, their values set
// to the tokenized given `value`.
//
// Unless plugins add new nodes, the sentence is populated by `WordNode`s,
// `SymbolNode`s, `PunctuationNode`s, and `WhiteSpaceNode`s.
pluggable(
  ParseLatin,
  'tokenizeSentence',
  createParser({
    type: 'SentenceNode',
    tokenizer: 'tokenize'
  })
);

// Create a `ParagraphNode` with its children set to `Node`s, their values set
// to the tokenized given `value`.
//
// Unless plugins add new nodes, the paragraph is populated by `SentenceNode`s
// and `WhiteSpaceNode`s.
pluggable(
  ParseLatin,
  'tokenizeParagraph',
  createParser({
    type: 'ParagraphNode',
    delimiter: expressions.terminalMarker,
    delimiterType: 'PunctuationNode',
    tokenizer: 'tokenizeSentence'
  })
);

// Create a `RootNode` with its children set to `Node`s, their values set to the
// tokenized given `value`.
pluggable(
  ParseLatin,
  'tokenizeRoot',
  createParser({
    type: 'RootNode',
    delimiter: expressions.newLine,
    delimiterType: 'WhiteSpaceNode',
    tokenizer: 'tokenizeParagraph'
  })
);

// PLUGINS

proto.use('tokenizeSentence', [
  requireMergeInitialWordSymbol(),
  requireMergeFinalWordSymbol(),
  requireMergeInnerWordSymbol(),
  requireMergeInnerWordSlash(),
  requireMergeInitialisms(),
  requireMergeWords(),
  requirePatchPosition()
]);

proto.use('tokenizeParagraph', [
  requireMergeNonWordSentences(),
  requireMergeAffixSymbol(),
  requireMergeInitialLowerCaseLetterSentences(),
  requireMergeInitialDigitSentences(),
  requireMergePrefixExceptions(),
  requireMergeAffixExceptions(),
  requireMergeRemainingFullStops(),
  requireMakeInitialWhiteSpaceSiblings(),
  requireMakeFinalWhiteSpaceSiblings(),
  requireBreakImplicitSentences(),
  requireRemoveEmptyNodes(),
  requirePatchPosition()
]);

proto.use('tokenizeRoot', [
  requireMakeInitialWhiteSpaceSiblings(),
  requireMakeFinalWhiteSpaceSiblings(),
  requireRemoveEmptyNodes(),
  requirePatchPosition()
]);

// TEXT NODES

// Factory to create a `Text`.
function createTextFactory(type) {
  type += 'Node';

  return createText

  // Construct a `Text` from a bound `type`
  function createText(value, eat, parent) {
    if (value === null || value === undefined) {
      value = '';
    }

    return (eat || noopEat)(value)(
      {
        type: type,
        value: String(value)
      },
      parent
    )
  }
}

// Run transform plug-ins for `key` on `nodes`.
function run(key, nodes) {
  var wareKey = key + 'Plugins';
  var plugins = this[wareKey];
  var index = -1;

  if (plugins) {
    while (plugins[++index]) {
      plugins[index](nodes);
    }
  }

  return nodes
}

// Make a method “pluggable”.
function pluggable(Constructor, key, callback) {
  // Set a pluggable version of `callback` on `Constructor`.
  Constructor.prototype[key] = function() {
    return this.run(key, callback.apply(this, arguments))
  };
}

// Factory to inject `plugins`. Takes `callback` for the actual inserting.
function useFactory(callback) {
  return use

  // Validate if `plugins` can be inserted.
  // Invokes the bound `callback` to do the actual inserting.
  function use(key, plugins) {
    var self = this;
    var wareKey;

    // Throw if the method is not pluggable.
    if (!(key in self)) {
      throw new Error(
        'Illegal Invocation: Unsupported `key` for ' +
          '`use(key, plugins)`. Make sure `key` is a ' +
          'supported function'
      )
    }

    // Fail silently when no plugins are given.
    if (!plugins) {
      return
    }

    wareKey = key + 'Plugins';

    // Make sure `plugins` is a list.
    if (typeof plugins === 'function') {
      plugins = [plugins];
    } else {
      plugins = plugins.concat();
    }

    // Make sure `wareKey` exists.
    if (!self[wareKey]) {
      self[wareKey] = [];
    }

    // Invoke callback with the ware key and plugins.
    callback(self, wareKey, plugins);
  }
}

// CLASSIFY

// Match a word character.
var wordRe = expressions.word;

// Match a surrogate character.
var surrogatesRe = expressions.surrogates;

// Match a punctuation character.
var punctuationRe = expressions.punctuation;

// Match a white space character.
var whiteSpaceRe = expressions.whiteSpace;

// Transform a `value` into a list of `NLCSTNode`s.
function tokenize(parser, value) {
  var tokens;
  var offset;
  var line;
  var column;
  var index;
  var length;
  var character;
  var queue;
  var prev;
  var left;
  var right;
  var eater;

  if (value === null || value === undefined) {
    value = '';
  } else if (value instanceof String) {
    value = value.toString();
  }

  if (typeof value !== 'string') {
    // Return the given nodes if this is either an empty array, or an array with
    // a node as a first child.
    if ('length' in value && (!value[0] || value[0].type)) {
      return value
    }

    throw new Error(
      "Illegal invocation: '" +
        value +
        "' is not a valid argument for 'ParseLatin'"
    )
  }

  tokens = [];

  if (!value) {
    return tokens
  }

  index = 0;
  offset = 0;
  line = 1;
  column = 1;

  // Eat mechanism to use.
  eater = parser.position ? eat : noPositionEat;

  length = value.length;
  prev = '';
  queue = '';

  while (index < length) {
    character = value.charAt(index);

    if (whiteSpaceRe.test(character)) {
      right = 'WhiteSpace';
    } else if (punctuationRe.test(character)) {
      right = 'Punctuation';
    } else if (wordRe.test(character)) {
      right = 'Word';
    } else {
      right = 'Symbol';
    }

    tick();

    prev = character;
    character = '';
    left = right;
    right = null;

    index++;
  }

  tick();

  return tokens

  // Check one character.
  function tick() {
    if (
      left === right &&
      (left === 'Word' ||
        left === 'WhiteSpace' ||
        character === prev ||
        surrogatesRe.test(character))
    ) {
      queue += character;
    } else {
      // Flush the previous queue.
      if (queue) {
        parser['tokenize' + left](queue, eater);
      }

      queue = character;
    }
  }

  // Remove `subvalue` from `value`.
  // Expects `subvalue` to be at the start from `value`, and applies no
  // validation.
  function eat(subvalue) {
    var pos = position();

    update(subvalue);

    return apply

    // Add the given arguments, add `position` to the returned node, and return
    // the node.
    function apply() {
      return pos(add.apply(null, arguments))
    }
  }

  // Remove `subvalue` from `value`.
  // Does not patch positional information.
  function noPositionEat() {
    return apply

    // Add the given arguments and return the node.
    function apply() {
      return add.apply(null, arguments)
    }
  }

  // Add mechanism.
  function add(node, parent) {
    if (parent) {
      parent.children.push(node);
    } else {
      tokens.push(node);
    }

    return node
  }

  // Mark position and patch `node.position`.
  function position() {
    var before = now();

    // Add the position to a node.
    function patch(node) {
      node.position = new Position(before);

      return node
    }

    return patch
  }

  // Update line and column based on `value`.
  function update(subvalue) {
    var subvalueLength = subvalue.length;
    var character = -1;
    var lastIndex = -1;

    offset += subvalueLength;

    while (++character < subvalueLength) {
      if (subvalue.charAt(character) === '\n') {
        lastIndex = character;
        line++;
      }
    }

    if (lastIndex === -1) {
      column += subvalueLength;
    } else {
      column = subvalueLength - lastIndex;
    }
  }

  // Store position information for a node.
  function Position(start) {
    this.start = start;
    this.end = now();
  }

  // Get the current position.
  function now() {
    return {
      line: line,
      column: column,
      offset: offset
    }
  }
}

// Add mechanism used when text-tokenisers are called directly outside of the
// `tokenize` function.
function noopAdd(node, parent) {
  if (parent) {
    parent.children.push(node);
  }

  return node
}

// Eat and add mechanism without adding positional information, used when
// text-tokenisers are called directly outside of the `tokenize` function.
function noopEat() {
  return noopAdd
}

var parseLatin = lib;

var unherit = unherit_1;
var Latin = parseLatin;

var retextLatin = parse;
parse.Parser = Latin;

function parse() {
  this.Parser = unherit(Latin);
}

var toString$1 = nlcstToString_1;

var retextStringify = stringify$2;

function stringify$2() {
  this.Compiler = compiler;
}

function compiler(tree) {
  return toString$1(tree)
}

var unified = unified_1;
var latin = retextLatin;
var stringify$1 = retextStringify;

var retext = unified()
  .use(latin)
  .use(stringify$1)
  .freeze();

var retext$1 = /*@__PURE__*/getDefaultExportFromCjs(retext);

var visit = unistUtilVisit;
var toString = nlcstToString_1;

var retextSmartypants = smartypants;

var punctuation = 'PunctuationNode';
var symbol = 'SymbolNode';
var word = 'WordNode';
var whiteSpace = 'WhiteSpaceNode';

var decadeExpression = /^\d\ds$/;
var threeFullStopsExpression = /^\.{3,}$/;
var fullStopsExpression = /^\.+$/;
var threeDashes = '---';
var twoDashes = '--';
var emDash = '—';
var enDash = '–';
var ellipsis = '…';
var twoBackticks = '``';
var backtick = '`';
var twoSingleQuotes = "''";
var singleQuote = "'";
var apostrophe = '’';
var doubleQuote = '"';
var openingDoubleQuote = '“';
var closingDoubleQuote = '”';
var openingSingleQuote = '‘';
var closingSingleQuote = '’';
var closingQuotes = {};
var openingQuotes = {};

openingQuotes[doubleQuote] = openingDoubleQuote;
closingQuotes[doubleQuote] = closingDoubleQuote;
openingQuotes[singleQuote] = openingSingleQuote;
closingQuotes[singleQuote] = closingSingleQuote;

var educators = {};

// Expose educators.
educators.dashes = {
  true: dashes,
  oldschool: oldschool,
  inverted: inverted
};

educators.backticks = {
  true: backticks,
  all: all
};

educators.ellipses = {
  true: ellipses
};

educators.quotes = {
  true: quotes
};

// Attacher.
function smartypants(options) {
  var methods = [];
  var quotes;
  var ellipses;
  var backticks;
  var dashes;

  if (!options) {
    options = {};
  }

  if ('quotes' in options) {
    quotes = options.quotes;

    if (quotes !== Boolean(quotes)) {
      throw new TypeError(
        'Illegal invocation: `' +
          quotes +
          '` ' +
          'is not a valid value for `quotes` in ' +
          '`smartypants`'
      )
    }
  } else {
    quotes = true;
  }

  if ('ellipses' in options) {
    ellipses = options.ellipses;

    if (ellipses !== Boolean(ellipses)) {
      throw new TypeError(
        'Illegal invocation: `' +
          ellipses +
          '` ' +
          'is not a valid value for `ellipses` in ' +
          '`smartypants`'
      )
    }
  } else {
    ellipses = true;
  }

  if ('backticks' in options) {
    backticks = options.backticks;

    if (backticks !== Boolean(backticks) && backticks !== 'all') {
      throw new TypeError(
        'Illegal invocation: `' +
          backticks +
          '` ' +
          'is not a valid value for `backticks` in ' +
          '`smartypants`'
      )
    }

    if (backticks === 'all' && quotes === true) {
      throw new TypeError(
        'Illegal invocation: `backticks: ' +
          backticks +
          '` is not a valid value ' +
          'when `quotes: ' +
          quotes +
          '` in ' +
          '`smartypants`'
      )
    }
  } else {
    backticks = true;
  }

  if ('dashes' in options) {
    dashes = options.dashes;

    if (
      dashes !== Boolean(dashes) &&
      dashes !== 'oldschool' &&
      dashes !== 'inverted'
    ) {
      throw new TypeError(
        'Illegal invocation: `' +
          dashes +
          '` ' +
          'is not a valid value for `dahes` in ' +
          '`smartypants`'
      )
    }
  } else {
    dashes = true;
  }

  if (quotes !== false) {
    methods.push(educators.quotes[quotes]);
  }

  if (ellipses !== false) {
    methods.push(educators.ellipses[ellipses]);
  }

  if (backticks !== false) {
    methods.push(educators.backticks[backticks]);
  }

  if (dashes !== false) {
    methods.push(educators.dashes[dashes]);
  }

  return transformFactory(methods)
}

// Create a transformer for the bound methods.
function transformFactory(methods) {
  var length = methods.length;

  return transformer

  // Transformer.
  function transformer(tree) {
    visit(tree, visitor);
  }

  function visitor(node, position, parent) {
    var index = -1;

    if (node.type === punctuation || node.type === symbol) {
      while (++index < length) {
        methods[index](node, position, parent);
      }
    }
  }
}

// Transform three dahes into an em-dash, and two into an en-dash.
function oldschool(node) {
  if (node.value === threeDashes) {
    node.value = emDash;
  } else if (node.value === twoDashes) {
    node.value = enDash;
  }
}

// Transform two dahes into an em-dash.
function dashes(node) {
  if (node.value === twoDashes) {
    node.value = emDash;
  }
}

// Transform three dahes into an en-dash, and two into an em-dash.
function inverted(node) {
  if (node.value === threeDashes) {
    node.value = enDash;
  } else if (node.value === twoDashes) {
    node.value = emDash;
  }
}

// Transform double backticks and single quotes into smart quotes.
function backticks(node) {
  if (node.value === twoBackticks) {
    node.value = openingDoubleQuote;
  } else if (node.value === twoSingleQuotes) {
    node.value = closingDoubleQuote;
  }
}

// Transform single and double backticks and single quotes into smart quotes.
function all(node) {
  backticks(node);

  if (node.value === backtick) {
    node.value = openingSingleQuote;
  } else if (node.value === singleQuote) {
    node.value = closingSingleQuote;
  }
}

// Transform multiple dots into unicode ellipses.
function ellipses(node, index, parent) {
  var value = node.value;
  var siblings = parent.children;
  var position;
  var nodes;
  var sibling;
  var type;
  var count;
  var queue;

  // Simple node with three dots and without white-space.
  if (threeFullStopsExpression.test(node.value)) {
    node.value = ellipsis;
    return
  }

  if (!fullStopsExpression.test(value)) {
    return
  }

  // Search for dot-nodes with white-space between.
  nodes = [];
  position = index;
  count = 1;

  // It’s possible that the node is merged with an adjacent word-node.  In that
  // code, we cannot transform it because there’s no reference to the
  // grandparent.
  while (--position > 0) {
    sibling = siblings[position];

    if (sibling.type !== whiteSpace) {
      break
    }

    queue = sibling;
    sibling = siblings[--position];
    type = sibling && sibling.type;

    if (
      sibling &&
      (type === punctuation || type === symbol) &&
      fullStopsExpression.test(sibling.value)
    ) {
      nodes.push(queue, sibling);

      count++;

      continue
    }

    break
  }

  if (count < 3) {
    return
  }

  siblings.splice(index - nodes.length, nodes.length);

  node.value = ellipsis;
}

// Transform straight single- and double quotes into smart quotes.
// eslint-disable-next-line complexity
function quotes(node, index, parent) {
  var siblings = parent.children;
  var value = node.value;
  var next;
  var nextNext;
  var prev;
  var nextValue;

  if (value !== doubleQuote && value !== singleQuote) {
    return
  }

  prev = siblings[index - 1];
  next = siblings[index + 1];
  nextNext = siblings[index + 2];
  nextValue = next && toString(next);

  if (
    next &&
    nextNext &&
    (next.type === punctuation || next.type === symbol) &&
    nextNext.type !== word
  ) {
    // Special case if the very first character is a quote followed by
    // punctuation at a non-word-break. Close the quotes by brute force.
    node.value = closingQuotes[value];
  } else if (
    nextNext &&
    (nextValue === doubleQuote || nextValue === singleQuote) &&
    nextNext.type === word
  ) {
    // Special case for double sets of quotes:
    // `He said, "'Quoted' words in a larger quote."`
    node.value = openingQuotes[value];
    next.value = openingQuotes[nextValue];
  } else if (next && decadeExpression.test(nextValue)) {
    // Special case for decade abbreviations: `the '80s`
    node.value = closingQuotes[value];
  } else if (
    prev &&
    next &&
    (prev.type === whiteSpace ||
      prev.type === punctuation ||
      prev.type === symbol) &&
    next.type === word
  ) {
    // Get most opening single quotes.
    node.value = openingQuotes[value];
  } else if (
    prev &&
    prev.type !== whiteSpace &&
    prev.type !== symbol &&
    prev.type !== punctuation
  ) {
    // Closing quotes.
    node.value = closingQuotes[value];
  } else if (
    !next ||
    next.type === whiteSpace ||
    ((value === singleQuote || value === apostrophe) && nextValue === 's')
  ) {
    node.value = closingQuotes[value];
  } else {
    node.value = openingQuotes[value];
  }
}

var smartypants$1 = /*@__PURE__*/getDefaultExportFromCjs(retextSmartypants);

var jsYaml$1 = {};

var loader$1 = {};

var common$6 = {};

function isNothing(subject) {
  return (typeof subject === 'undefined') || (subject === null);
}


function isObject(subject) {
  return (typeof subject === 'object') && (subject !== null);
}


function toArray(sequence) {
  if (Array.isArray(sequence)) return sequence;
  else if (isNothing(sequence)) return [];

  return [ sequence ];
}


function extend(target, source) {
  var index, length, key, sourceKeys;

  if (source) {
    sourceKeys = Object.keys(source);

    for (index = 0, length = sourceKeys.length; index < length; index += 1) {
      key = sourceKeys[index];
      target[key] = source[key];
    }
  }

  return target;
}


function repeat(string, count) {
  var result = '', cycle;

  for (cycle = 0; cycle < count; cycle += 1) {
    result += string;
  }

  return result;
}


function isNegativeZero(number) {
  return (number === 0) && (Number.NEGATIVE_INFINITY === 1 / number);
}


common$6.isNothing      = isNothing;
common$6.isObject       = isObject;
common$6.toArray        = toArray;
common$6.repeat         = repeat;
common$6.isNegativeZero = isNegativeZero;
common$6.extend         = extend;

function YAMLException$4(reason, mark) {
  // Super constructor
  Error.call(this);

  this.name = 'YAMLException';
  this.reason = reason;
  this.mark = mark;
  this.message = (this.reason || '(unknown reason)') + (this.mark ? ' ' + this.mark.toString() : '');

  // Include stack trace in error object
  if (Error.captureStackTrace) {
    // Chrome and NodeJS
    Error.captureStackTrace(this, this.constructor);
  } else {
    // FF, IE 10+ and Safari 6+. Fallback for others
    this.stack = (new Error()).stack || '';
  }
}


// Inherit from Error
YAMLException$4.prototype = Object.create(Error.prototype);
YAMLException$4.prototype.constructor = YAMLException$4;


YAMLException$4.prototype.toString = function toString(compact) {
  var result = this.name + ': ';

  result += this.reason || '(unknown reason)';

  if (!compact && this.mark) {
    result += ' ' + this.mark.toString();
  }

  return result;
};


var exception = YAMLException$4;

var common$5 = common$6;


function Mark$1(name, buffer, position, line, column) {
  this.name     = name;
  this.buffer   = buffer;
  this.position = position;
  this.line     = line;
  this.column   = column;
}


Mark$1.prototype.getSnippet = function getSnippet(indent, maxLength) {
  var head, start, tail, end, snippet;

  if (!this.buffer) return null;

  indent = indent || 4;
  maxLength = maxLength || 75;

  head = '';
  start = this.position;

  while (start > 0 && '\x00\r\n\x85\u2028\u2029'.indexOf(this.buffer.charAt(start - 1)) === -1) {
    start -= 1;
    if (this.position - start > (maxLength / 2 - 1)) {
      head = ' ... ';
      start += 5;
      break;
    }
  }

  tail = '';
  end = this.position;

  while (end < this.buffer.length && '\x00\r\n\x85\u2028\u2029'.indexOf(this.buffer.charAt(end)) === -1) {
    end += 1;
    if (end - this.position > (maxLength / 2 - 1)) {
      tail = ' ... ';
      end -= 5;
      break;
    }
  }

  snippet = this.buffer.slice(start, end);

  return common$5.repeat(' ', indent) + head + snippet + tail + '\n' +
         common$5.repeat(' ', indent + this.position - start + head.length) + '^';
};


Mark$1.prototype.toString = function toString(compact) {
  var snippet, where = '';

  if (this.name) {
    where += 'in "' + this.name + '" ';
  }

  where += 'at line ' + (this.line + 1) + ', column ' + (this.column + 1);

  if (!compact) {
    snippet = this.getSnippet();

    if (snippet) {
      where += ':\n' + snippet;
    }
  }

  return where;
};


var mark = Mark$1;

var YAMLException$3 = exception;

var TYPE_CONSTRUCTOR_OPTIONS = [
  'kind',
  'resolve',
  'construct',
  'instanceOf',
  'predicate',
  'represent',
  'defaultStyle',
  'styleAliases'
];

var YAML_NODE_KINDS = [
  'scalar',
  'sequence',
  'mapping'
];

function compileStyleAliases(map) {
  var result = {};

  if (map !== null) {
    Object.keys(map).forEach(function (style) {
      map[style].forEach(function (alias) {
        result[String(alias)] = style;
      });
    });
  }

  return result;
}

function Type$h(tag, options) {
  options = options || {};

  Object.keys(options).forEach(function (name) {
    if (TYPE_CONSTRUCTOR_OPTIONS.indexOf(name) === -1) {
      throw new YAMLException$3('Unknown option "' + name + '" is met in definition of "' + tag + '" YAML type.');
    }
  });

  // TODO: Add tag format check.
  this.tag          = tag;
  this.kind         = options['kind']         || null;
  this.resolve      = options['resolve']      || function () { return true; };
  this.construct    = options['construct']    || function (data) { return data; };
  this.instanceOf   = options['instanceOf']   || null;
  this.predicate    = options['predicate']    || null;
  this.represent    = options['represent']    || null;
  this.defaultStyle = options['defaultStyle'] || null;
  this.styleAliases = compileStyleAliases(options['styleAliases'] || null);

  if (YAML_NODE_KINDS.indexOf(this.kind) === -1) {
    throw new YAMLException$3('Unknown kind "' + this.kind + '" is specified for "' + tag + '" YAML type.');
  }
}

var type = Type$h;

/*eslint-disable max-len*/

var common$4        = common$6;
var YAMLException$2 = exception;
var Type$g          = type;


function compileList(schema, name, result) {
  var exclude = [];

  schema.include.forEach(function (includedSchema) {
    result = compileList(includedSchema, name, result);
  });

  schema[name].forEach(function (currentType) {
    result.forEach(function (previousType, previousIndex) {
      if (previousType.tag === currentType.tag && previousType.kind === currentType.kind) {
        exclude.push(previousIndex);
      }
    });

    result.push(currentType);
  });

  return result.filter(function (type, index) {
    return exclude.indexOf(index) === -1;
  });
}


function compileMap(/* lists... */) {
  var result = {
        scalar: {},
        sequence: {},
        mapping: {},
        fallback: {}
      }, index, length;

  function collectType(type) {
    result[type.kind][type.tag] = result['fallback'][type.tag] = type;
  }

  for (index = 0, length = arguments.length; index < length; index += 1) {
    arguments[index].forEach(collectType);
  }
  return result;
}


function Schema$5(definition) {
  this.include  = definition.include  || [];
  this.implicit = definition.implicit || [];
  this.explicit = definition.explicit || [];

  this.implicit.forEach(function (type) {
    if (type.loadKind && type.loadKind !== 'scalar') {
      throw new YAMLException$2('There is a non-scalar type in the implicit list of a schema. Implicit resolving of such types is not supported.');
    }
  });

  this.compiledImplicit = compileList(this, 'implicit', []);
  this.compiledExplicit = compileList(this, 'explicit', []);
  this.compiledTypeMap  = compileMap(this.compiledImplicit, this.compiledExplicit);
}


Schema$5.DEFAULT = null;


Schema$5.create = function createSchema() {
  var schemas, types;

  switch (arguments.length) {
    case 1:
      schemas = Schema$5.DEFAULT;
      types = arguments[0];
      break;

    case 2:
      schemas = arguments[0];
      types = arguments[1];
      break;

    default:
      throw new YAMLException$2('Wrong number of arguments for Schema.create function');
  }

  schemas = common$4.toArray(schemas);
  types = common$4.toArray(types);

  if (!schemas.every(function (schema) { return schema instanceof Schema$5; })) {
    throw new YAMLException$2('Specified list of super schemas (or a single Schema object) contains a non-Schema object.');
  }

  if (!types.every(function (type) { return type instanceof Type$g; })) {
    throw new YAMLException$2('Specified list of YAML types (or a single Type object) contains a non-Type object.');
  }

  return new Schema$5({
    include: schemas,
    explicit: types
  });
};


var schema = Schema$5;

var Type$f = type;

var str = new Type$f('tag:yaml.org,2002:str', {
  kind: 'scalar',
  construct: function (data) { return data !== null ? data : ''; }
});

var Type$e = type;

var seq = new Type$e('tag:yaml.org,2002:seq', {
  kind: 'sequence',
  construct: function (data) { return data !== null ? data : []; }
});

var Type$d = type;

var map = new Type$d('tag:yaml.org,2002:map', {
  kind: 'mapping',
  construct: function (data) { return data !== null ? data : {}; }
});

var Schema$4 = schema;


var failsafe = new Schema$4({
  explicit: [
    str,
    seq,
    map
  ]
});

var Type$c = type;

function resolveYamlNull(data) {
  if (data === null) return true;

  var max = data.length;

  return (max === 1 && data === '~') ||
         (max === 4 && (data === 'null' || data === 'Null' || data === 'NULL'));
}

function constructYamlNull() {
  return null;
}

function isNull(object) {
  return object === null;
}

var _null = new Type$c('tag:yaml.org,2002:null', {
  kind: 'scalar',
  resolve: resolveYamlNull,
  construct: constructYamlNull,
  predicate: isNull,
  represent: {
    canonical: function () { return '~';    },
    lowercase: function () { return 'null'; },
    uppercase: function () { return 'NULL'; },
    camelcase: function () { return 'Null'; }
  },
  defaultStyle: 'lowercase'
});

var Type$b = type;

function resolveYamlBoolean(data) {
  if (data === null) return false;

  var max = data.length;

  return (max === 4 && (data === 'true' || data === 'True' || data === 'TRUE')) ||
         (max === 5 && (data === 'false' || data === 'False' || data === 'FALSE'));
}

function constructYamlBoolean(data) {
  return data === 'true' ||
         data === 'True' ||
         data === 'TRUE';
}

function isBoolean(object) {
  return Object.prototype.toString.call(object) === '[object Boolean]';
}

var bool = new Type$b('tag:yaml.org,2002:bool', {
  kind: 'scalar',
  resolve: resolveYamlBoolean,
  construct: constructYamlBoolean,
  predicate: isBoolean,
  represent: {
    lowercase: function (object) { return object ? 'true' : 'false'; },
    uppercase: function (object) { return object ? 'TRUE' : 'FALSE'; },
    camelcase: function (object) { return object ? 'True' : 'False'; }
  },
  defaultStyle: 'lowercase'
});

var common$3 = common$6;
var Type$a   = type;

function isHexCode(c) {
  return ((0x30/* 0 */ <= c) && (c <= 0x39/* 9 */)) ||
         ((0x41/* A */ <= c) && (c <= 0x46/* F */)) ||
         ((0x61/* a */ <= c) && (c <= 0x66/* f */));
}

function isOctCode(c) {
  return ((0x30/* 0 */ <= c) && (c <= 0x37/* 7 */));
}

function isDecCode(c) {
  return ((0x30/* 0 */ <= c) && (c <= 0x39/* 9 */));
}

function resolveYamlInteger(data) {
  if (data === null) return false;

  var max = data.length,
      index = 0,
      hasDigits = false,
      ch;

  if (!max) return false;

  ch = data[index];

  // sign
  if (ch === '-' || ch === '+') {
    ch = data[++index];
  }

  if (ch === '0') {
    // 0
    if (index + 1 === max) return true;
    ch = data[++index];

    // base 2, base 8, base 16

    if (ch === 'b') {
      // base 2
      index++;

      for (; index < max; index++) {
        ch = data[index];
        if (ch === '_') continue;
        if (ch !== '0' && ch !== '1') return false;
        hasDigits = true;
      }
      return hasDigits && ch !== '_';
    }


    if (ch === 'x') {
      // base 16
      index++;

      for (; index < max; index++) {
        ch = data[index];
        if (ch === '_') continue;
        if (!isHexCode(data.charCodeAt(index))) return false;
        hasDigits = true;
      }
      return hasDigits && ch !== '_';
    }

    // base 8
    for (; index < max; index++) {
      ch = data[index];
      if (ch === '_') continue;
      if (!isOctCode(data.charCodeAt(index))) return false;
      hasDigits = true;
    }
    return hasDigits && ch !== '_';
  }

  // base 10 (except 0) or base 60

  // value should not start with `_`;
  if (ch === '_') return false;

  for (; index < max; index++) {
    ch = data[index];
    if (ch === '_') continue;
    if (ch === ':') break;
    if (!isDecCode(data.charCodeAt(index))) {
      return false;
    }
    hasDigits = true;
  }

  // Should have digits and should not end with `_`
  if (!hasDigits || ch === '_') return false;

  // if !base60 - done;
  if (ch !== ':') return true;

  // base60 almost not used, no needs to optimize
  return /^(:[0-5]?[0-9])+$/.test(data.slice(index));
}

function constructYamlInteger(data) {
  var value = data, sign = 1, ch, base, digits = [];

  if (value.indexOf('_') !== -1) {
    value = value.replace(/_/g, '');
  }

  ch = value[0];

  if (ch === '-' || ch === '+') {
    if (ch === '-') sign = -1;
    value = value.slice(1);
    ch = value[0];
  }

  if (value === '0') return 0;

  if (ch === '0') {
    if (value[1] === 'b') return sign * parseInt(value.slice(2), 2);
    if (value[1] === 'x') return sign * parseInt(value, 16);
    return sign * parseInt(value, 8);
  }

  if (value.indexOf(':') !== -1) {
    value.split(':').forEach(function (v) {
      digits.unshift(parseInt(v, 10));
    });

    value = 0;
    base = 1;

    digits.forEach(function (d) {
      value += (d * base);
      base *= 60;
    });

    return sign * value;

  }

  return sign * parseInt(value, 10);
}

function isInteger(object) {
  return (Object.prototype.toString.call(object)) === '[object Number]' &&
         (object % 1 === 0 && !common$3.isNegativeZero(object));
}

var int = new Type$a('tag:yaml.org,2002:int', {
  kind: 'scalar',
  resolve: resolveYamlInteger,
  construct: constructYamlInteger,
  predicate: isInteger,
  represent: {
    binary:      function (obj) { return obj >= 0 ? '0b' + obj.toString(2) : '-0b' + obj.toString(2).slice(1); },
    octal:       function (obj) { return obj >= 0 ? '0'  + obj.toString(8) : '-0'  + obj.toString(8).slice(1); },
    decimal:     function (obj) { return obj.toString(10); },
    /* eslint-disable max-len */
    hexadecimal: function (obj) { return obj >= 0 ? '0x' + obj.toString(16).toUpperCase() :  '-0x' + obj.toString(16).toUpperCase().slice(1); }
  },
  defaultStyle: 'decimal',
  styleAliases: {
    binary:      [ 2,  'bin' ],
    octal:       [ 8,  'oct' ],
    decimal:     [ 10, 'dec' ],
    hexadecimal: [ 16, 'hex' ]
  }
});

var common$2 = common$6;
var Type$9   = type;

var YAML_FLOAT_PATTERN = new RegExp(
  // 2.5e4, 2.5 and integers
  '^(?:[-+]?(?:0|[1-9][0-9_]*)(?:\\.[0-9_]*)?(?:[eE][-+]?[0-9]+)?' +
  // .2e4, .2
  // special case, seems not from spec
  '|\\.[0-9_]+(?:[eE][-+]?[0-9]+)?' +
  // 20:59
  '|[-+]?[0-9][0-9_]*(?::[0-5]?[0-9])+\\.[0-9_]*' +
  // .inf
  '|[-+]?\\.(?:inf|Inf|INF)' +
  // .nan
  '|\\.(?:nan|NaN|NAN))$');

function resolveYamlFloat(data) {
  if (data === null) return false;

  if (!YAML_FLOAT_PATTERN.test(data) ||
      // Quick hack to not allow integers end with `_`
      // Probably should update regexp & check speed
      data[data.length - 1] === '_') {
    return false;
  }

  return true;
}

function constructYamlFloat(data) {
  var value, sign, base, digits;

  value  = data.replace(/_/g, '').toLowerCase();
  sign   = value[0] === '-' ? -1 : 1;
  digits = [];

  if ('+-'.indexOf(value[0]) >= 0) {
    value = value.slice(1);
  }

  if (value === '.inf') {
    return (sign === 1) ? Number.POSITIVE_INFINITY : Number.NEGATIVE_INFINITY;

  } else if (value === '.nan') {
    return NaN;

  } else if (value.indexOf(':') >= 0) {
    value.split(':').forEach(function (v) {
      digits.unshift(parseFloat(v, 10));
    });

    value = 0.0;
    base = 1;

    digits.forEach(function (d) {
      value += d * base;
      base *= 60;
    });

    return sign * value;

  }
  return sign * parseFloat(value, 10);
}


var SCIENTIFIC_WITHOUT_DOT = /^[-+]?[0-9]+e/;

function representYamlFloat(object, style) {
  var res;

  if (isNaN(object)) {
    switch (style) {
      case 'lowercase': return '.nan';
      case 'uppercase': return '.NAN';
      case 'camelcase': return '.NaN';
    }
  } else if (Number.POSITIVE_INFINITY === object) {
    switch (style) {
      case 'lowercase': return '.inf';
      case 'uppercase': return '.INF';
      case 'camelcase': return '.Inf';
    }
  } else if (Number.NEGATIVE_INFINITY === object) {
    switch (style) {
      case 'lowercase': return '-.inf';
      case 'uppercase': return '-.INF';
      case 'camelcase': return '-.Inf';
    }
  } else if (common$2.isNegativeZero(object)) {
    return '-0.0';
  }

  res = object.toString(10);

  // JS stringifier can build scientific format without dots: 5e-100,
  // while YAML requres dot: 5.e-100. Fix it with simple hack

  return SCIENTIFIC_WITHOUT_DOT.test(res) ? res.replace('e', '.e') : res;
}

function isFloat(object) {
  return (Object.prototype.toString.call(object) === '[object Number]') &&
         (object % 1 !== 0 || common$2.isNegativeZero(object));
}

var float = new Type$9('tag:yaml.org,2002:float', {
  kind: 'scalar',
  resolve: resolveYamlFloat,
  construct: constructYamlFloat,
  predicate: isFloat,
  represent: representYamlFloat,
  defaultStyle: 'lowercase'
});

var Schema$3 = schema;


var json = new Schema$3({
  include: [
    failsafe
  ],
  implicit: [
    _null,
    bool,
    int,
    float
  ]
});

var Schema$2 = schema;


var core = new Schema$2({
  include: [
    json
  ]
});

var Type$8 = type;

var YAML_DATE_REGEXP = new RegExp(
  '^([0-9][0-9][0-9][0-9])'          + // [1] year
  '-([0-9][0-9])'                    + // [2] month
  '-([0-9][0-9])$');                   // [3] day

var YAML_TIMESTAMP_REGEXP = new RegExp(
  '^([0-9][0-9][0-9][0-9])'          + // [1] year
  '-([0-9][0-9]?)'                   + // [2] month
  '-([0-9][0-9]?)'                   + // [3] day
  '(?:[Tt]|[ \\t]+)'                 + // ...
  '([0-9][0-9]?)'                    + // [4] hour
  ':([0-9][0-9])'                    + // [5] minute
  ':([0-9][0-9])'                    + // [6] second
  '(?:\\.([0-9]*))?'                 + // [7] fraction
  '(?:[ \\t]*(Z|([-+])([0-9][0-9]?)' + // [8] tz [9] tz_sign [10] tz_hour
  '(?::([0-9][0-9]))?))?$');           // [11] tz_minute

function resolveYamlTimestamp(data) {
  if (data === null) return false;
  if (YAML_DATE_REGEXP.exec(data) !== null) return true;
  if (YAML_TIMESTAMP_REGEXP.exec(data) !== null) return true;
  return false;
}

function constructYamlTimestamp(data) {
  var match, year, month, day, hour, minute, second, fraction = 0,
      delta = null, tz_hour, tz_minute, date;

  match = YAML_DATE_REGEXP.exec(data);
  if (match === null) match = YAML_TIMESTAMP_REGEXP.exec(data);

  if (match === null) throw new Error('Date resolve error');

  // match: [1] year [2] month [3] day

  year = +(match[1]);
  month = +(match[2]) - 1; // JS month starts with 0
  day = +(match[3]);

  if (!match[4]) { // no hour
    return new Date(Date.UTC(year, month, day));
  }

  // match: [4] hour [5] minute [6] second [7] fraction

  hour = +(match[4]);
  minute = +(match[5]);
  second = +(match[6]);

  if (match[7]) {
    fraction = match[7].slice(0, 3);
    while (fraction.length < 3) { // milli-seconds
      fraction += '0';
    }
    fraction = +fraction;
  }

  // match: [8] tz [9] tz_sign [10] tz_hour [11] tz_minute

  if (match[9]) {
    tz_hour = +(match[10]);
    tz_minute = +(match[11] || 0);
    delta = (tz_hour * 60 + tz_minute) * 60000; // delta in mili-seconds
    if (match[9] === '-') delta = -delta;
  }

  date = new Date(Date.UTC(year, month, day, hour, minute, second, fraction));

  if (delta) date.setTime(date.getTime() - delta);

  return date;
}

function representYamlTimestamp(object /*, style*/) {
  return object.toISOString();
}

var timestamp = new Type$8('tag:yaml.org,2002:timestamp', {
  kind: 'scalar',
  resolve: resolveYamlTimestamp,
  construct: constructYamlTimestamp,
  instanceOf: Date,
  represent: representYamlTimestamp
});

var Type$7 = type;

function resolveYamlMerge(data) {
  return data === '<<' || data === null;
}

var merge = new Type$7('tag:yaml.org,2002:merge', {
  kind: 'scalar',
  resolve: resolveYamlMerge
});

function commonjsRequire(path) {
	throw new Error('Could not dynamically require "' + path + '". Please configure the dynamicRequireTargets or/and ignoreDynamicRequires option of @rollup/plugin-commonjs appropriately for this require call to work.');
}

/*eslint-disable no-bitwise*/

var NodeBuffer;

try {
  // A trick for browserified version, to not include `Buffer` shim
  var _require$1 = commonjsRequire;
  NodeBuffer = _require$1('buffer').Buffer;
} catch (__) {}

var Type$6       = type;


// [ 64, 65, 66 ] -> [ padding, CR, LF ]
var BASE64_MAP = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=\n\r';


function resolveYamlBinary(data) {
  if (data === null) return false;

  var code, idx, bitlen = 0, max = data.length, map = BASE64_MAP;

  // Convert one by one.
  for (idx = 0; idx < max; idx++) {
    code = map.indexOf(data.charAt(idx));

    // Skip CR/LF
    if (code > 64) continue;

    // Fail on illegal characters
    if (code < 0) return false;

    bitlen += 6;
  }

  // If there are any bits left, source was corrupted
  return (bitlen % 8) === 0;
}

function constructYamlBinary(data) {
  var idx, tailbits,
      input = data.replace(/[\r\n=]/g, ''), // remove CR/LF & padding to simplify scan
      max = input.length,
      map = BASE64_MAP,
      bits = 0,
      result = [];

  // Collect by 6*4 bits (3 bytes)

  for (idx = 0; idx < max; idx++) {
    if ((idx % 4 === 0) && idx) {
      result.push((bits >> 16) & 0xFF);
      result.push((bits >> 8) & 0xFF);
      result.push(bits & 0xFF);
    }

    bits = (bits << 6) | map.indexOf(input.charAt(idx));
  }

  // Dump tail

  tailbits = (max % 4) * 6;

  if (tailbits === 0) {
    result.push((bits >> 16) & 0xFF);
    result.push((bits >> 8) & 0xFF);
    result.push(bits & 0xFF);
  } else if (tailbits === 18) {
    result.push((bits >> 10) & 0xFF);
    result.push((bits >> 2) & 0xFF);
  } else if (tailbits === 12) {
    result.push((bits >> 4) & 0xFF);
  }

  // Wrap into Buffer for NodeJS and leave Array for browser
  if (NodeBuffer) {
    // Support node 6.+ Buffer API when available
    return NodeBuffer.from ? NodeBuffer.from(result) : new NodeBuffer(result);
  }

  return result;
}

function representYamlBinary(object /*, style*/) {
  var result = '', bits = 0, idx, tail,
      max = object.length,
      map = BASE64_MAP;

  // Convert every three bytes to 4 ASCII characters.

  for (idx = 0; idx < max; idx++) {
    if ((idx % 3 === 0) && idx) {
      result += map[(bits >> 18) & 0x3F];
      result += map[(bits >> 12) & 0x3F];
      result += map[(bits >> 6) & 0x3F];
      result += map[bits & 0x3F];
    }

    bits = (bits << 8) + object[idx];
  }

  // Dump tail

  tail = max % 3;

  if (tail === 0) {
    result += map[(bits >> 18) & 0x3F];
    result += map[(bits >> 12) & 0x3F];
    result += map[(bits >> 6) & 0x3F];
    result += map[bits & 0x3F];
  } else if (tail === 2) {
    result += map[(bits >> 10) & 0x3F];
    result += map[(bits >> 4) & 0x3F];
    result += map[(bits << 2) & 0x3F];
    result += map[64];
  } else if (tail === 1) {
    result += map[(bits >> 2) & 0x3F];
    result += map[(bits << 4) & 0x3F];
    result += map[64];
    result += map[64];
  }

  return result;
}

function isBinary(object) {
  return NodeBuffer && NodeBuffer.isBuffer(object);
}

var binary = new Type$6('tag:yaml.org,2002:binary', {
  kind: 'scalar',
  resolve: resolveYamlBinary,
  construct: constructYamlBinary,
  predicate: isBinary,
  represent: representYamlBinary
});

var Type$5 = type;

var _hasOwnProperty$3 = Object.prototype.hasOwnProperty;
var _toString$2       = Object.prototype.toString;

function resolveYamlOmap(data) {
  if (data === null) return true;

  var objectKeys = [], index, length, pair, pairKey, pairHasKey,
      object = data;

  for (index = 0, length = object.length; index < length; index += 1) {
    pair = object[index];
    pairHasKey = false;

    if (_toString$2.call(pair) !== '[object Object]') return false;

    for (pairKey in pair) {
      if (_hasOwnProperty$3.call(pair, pairKey)) {
        if (!pairHasKey) pairHasKey = true;
        else return false;
      }
    }

    if (!pairHasKey) return false;

    if (objectKeys.indexOf(pairKey) === -1) objectKeys.push(pairKey);
    else return false;
  }

  return true;
}

function constructYamlOmap(data) {
  return data !== null ? data : [];
}

var omap = new Type$5('tag:yaml.org,2002:omap', {
  kind: 'sequence',
  resolve: resolveYamlOmap,
  construct: constructYamlOmap
});

var Type$4 = type;

var _toString$1 = Object.prototype.toString;

function resolveYamlPairs(data) {
  if (data === null) return true;

  var index, length, pair, keys, result,
      object = data;

  result = new Array(object.length);

  for (index = 0, length = object.length; index < length; index += 1) {
    pair = object[index];

    if (_toString$1.call(pair) !== '[object Object]') return false;

    keys = Object.keys(pair);

    if (keys.length !== 1) return false;

    result[index] = [ keys[0], pair[keys[0]] ];
  }

  return true;
}

function constructYamlPairs(data) {
  if (data === null) return [];

  var index, length, pair, keys, result,
      object = data;

  result = new Array(object.length);

  for (index = 0, length = object.length; index < length; index += 1) {
    pair = object[index];

    keys = Object.keys(pair);

    result[index] = [ keys[0], pair[keys[0]] ];
  }

  return result;
}

var pairs = new Type$4('tag:yaml.org,2002:pairs', {
  kind: 'sequence',
  resolve: resolveYamlPairs,
  construct: constructYamlPairs
});

var Type$3 = type;

var _hasOwnProperty$2 = Object.prototype.hasOwnProperty;

function resolveYamlSet(data) {
  if (data === null) return true;

  var key, object = data;

  for (key in object) {
    if (_hasOwnProperty$2.call(object, key)) {
      if (object[key] !== null) return false;
    }
  }

  return true;
}

function constructYamlSet(data) {
  return data !== null ? data : {};
}

var set = new Type$3('tag:yaml.org,2002:set', {
  kind: 'mapping',
  resolve: resolveYamlSet,
  construct: constructYamlSet
});

var Schema$1 = schema;


var default_safe = new Schema$1({
  include: [
    core
  ],
  implicit: [
    timestamp,
    merge
  ],
  explicit: [
    binary,
    omap,
    pairs,
    set
  ]
});

var Type$2 = type;

function resolveJavascriptUndefined() {
  return true;
}

function constructJavascriptUndefined() {
  /*eslint-disable no-undefined*/
  return undefined;
}

function representJavascriptUndefined() {
  return '';
}

function isUndefined(object) {
  return typeof object === 'undefined';
}

var _undefined = new Type$2('tag:yaml.org,2002:js/undefined', {
  kind: 'scalar',
  resolve: resolveJavascriptUndefined,
  construct: constructJavascriptUndefined,
  predicate: isUndefined,
  represent: representJavascriptUndefined
});

var Type$1 = type;

function resolveJavascriptRegExp(data) {
  if (data === null) return false;
  if (data.length === 0) return false;

  var regexp = data,
      tail   = /\/([gim]*)$/.exec(data),
      modifiers = '';

  // if regexp starts with '/' it can have modifiers and must be properly closed
  // `/foo/gim` - modifiers tail can be maximum 3 chars
  if (regexp[0] === '/') {
    if (tail) modifiers = tail[1];

    if (modifiers.length > 3) return false;
    // if expression starts with /, is should be properly terminated
    if (regexp[regexp.length - modifiers.length - 1] !== '/') return false;
  }

  return true;
}

function constructJavascriptRegExp(data) {
  var regexp = data,
      tail   = /\/([gim]*)$/.exec(data),
      modifiers = '';

  // `/foo/gim` - tail can be maximum 4 chars
  if (regexp[0] === '/') {
    if (tail) modifiers = tail[1];
    regexp = regexp.slice(1, regexp.length - modifiers.length - 1);
  }

  return new RegExp(regexp, modifiers);
}

function representJavascriptRegExp(object /*, style*/) {
  var result = '/' + object.source + '/';

  if (object.global) result += 'g';
  if (object.multiline) result += 'm';
  if (object.ignoreCase) result += 'i';

  return result;
}

function isRegExp(object) {
  return Object.prototype.toString.call(object) === '[object RegExp]';
}

var regexp = new Type$1('tag:yaml.org,2002:js/regexp', {
  kind: 'scalar',
  resolve: resolveJavascriptRegExp,
  construct: constructJavascriptRegExp,
  predicate: isRegExp,
  represent: representJavascriptRegExp
});

var esprima;

// Browserified version does not have esprima
//
// 1. For node.js just require module as deps
// 2. For browser try to require mudule via external AMD system.
//    If not found - try to fallback to window.esprima. If not
//    found too - then fail to parse.
//
try {
  // workaround to exclude package from browserify list.
  var _require = commonjsRequire;
  esprima = _require('esprima');
} catch (_) {
  /* eslint-disable no-redeclare */
  /* global window */
  if (typeof window !== 'undefined') esprima = window.esprima;
}

var Type = type;

function resolveJavascriptFunction(data) {
  if (data === null) return false;

  try {
    var source = '(' + data + ')',
        ast    = esprima.parse(source, { range: true });

    if (ast.type                    !== 'Program'             ||
        ast.body.length             !== 1                     ||
        ast.body[0].type            !== 'ExpressionStatement' ||
        (ast.body[0].expression.type !== 'ArrowFunctionExpression' &&
          ast.body[0].expression.type !== 'FunctionExpression')) {
      return false;
    }

    return true;
  } catch (err) {
    return false;
  }
}

function constructJavascriptFunction(data) {
  /*jslint evil:true*/

  var source = '(' + data + ')',
      ast    = esprima.parse(source, { range: true }),
      params = [],
      body;

  if (ast.type                    !== 'Program'             ||
      ast.body.length             !== 1                     ||
      ast.body[0].type            !== 'ExpressionStatement' ||
      (ast.body[0].expression.type !== 'ArrowFunctionExpression' &&
        ast.body[0].expression.type !== 'FunctionExpression')) {
    throw new Error('Failed to resolve function');
  }

  ast.body[0].expression.params.forEach(function (param) {
    params.push(param.name);
  });

  body = ast.body[0].expression.body.range;

  // Esprima's ranges include the first '{' and the last '}' characters on
  // function expressions. So cut them out.
  if (ast.body[0].expression.body.type === 'BlockStatement') {
    /*eslint-disable no-new-func*/
    return new Function(params, source.slice(body[0] + 1, body[1] - 1));
  }
  // ES6 arrow functions can omit the BlockStatement. In that case, just return
  // the body.
  /*eslint-disable no-new-func*/
  return new Function(params, 'return ' + source.slice(body[0], body[1]));
}

function representJavascriptFunction(object /*, style*/) {
  return object.toString();
}

function isFunction(object) {
  return Object.prototype.toString.call(object) === '[object Function]';
}

var _function = new Type('tag:yaml.org,2002:js/function', {
  kind: 'scalar',
  resolve: resolveJavascriptFunction,
  construct: constructJavascriptFunction,
  predicate: isFunction,
  represent: representJavascriptFunction
});

var Schema = schema;


var default_full = Schema.DEFAULT = new Schema({
  include: [
    default_safe
  ],
  explicit: [
    _undefined,
    regexp,
    _function
  ]
});

/*eslint-disable max-len,no-use-before-define*/

var common$1              = common$6;
var YAMLException$1       = exception;
var Mark                = mark;
var DEFAULT_SAFE_SCHEMA$1 = default_safe;
var DEFAULT_FULL_SCHEMA$1 = default_full;


var _hasOwnProperty$1 = Object.prototype.hasOwnProperty;


var CONTEXT_FLOW_IN   = 1;
var CONTEXT_FLOW_OUT  = 2;
var CONTEXT_BLOCK_IN  = 3;
var CONTEXT_BLOCK_OUT = 4;


var CHOMPING_CLIP  = 1;
var CHOMPING_STRIP = 2;
var CHOMPING_KEEP  = 3;


var PATTERN_NON_PRINTABLE         = /[\x00-\x08\x0B\x0C\x0E-\x1F\x7F-\x84\x86-\x9F\uFFFE\uFFFF]|[\uD800-\uDBFF](?![\uDC00-\uDFFF])|(?:[^\uD800-\uDBFF]|^)[\uDC00-\uDFFF]/;
var PATTERN_NON_ASCII_LINE_BREAKS = /[\x85\u2028\u2029]/;
var PATTERN_FLOW_INDICATORS       = /[,\[\]\{\}]/;
var PATTERN_TAG_HANDLE            = /^(?:!|!!|![a-z\-]+!)$/i;
var PATTERN_TAG_URI               = /^(?:!|[^,\[\]\{\}])(?:%[0-9a-f]{2}|[0-9a-z\-#;\/\?:@&=\+\$,_\.!~\*'\(\)\[\]])*$/i;


function _class(obj) { return Object.prototype.toString.call(obj); }

function is_EOL(c) {
  return (c === 0x0A/* LF */) || (c === 0x0D/* CR */);
}

function is_WHITE_SPACE(c) {
  return (c === 0x09/* Tab */) || (c === 0x20/* Space */);
}

function is_WS_OR_EOL(c) {
  return (c === 0x09/* Tab */) ||
         (c === 0x20/* Space */) ||
         (c === 0x0A/* LF */) ||
         (c === 0x0D/* CR */);
}

function is_FLOW_INDICATOR(c) {
  return c === 0x2C/* , */ ||
         c === 0x5B/* [ */ ||
         c === 0x5D/* ] */ ||
         c === 0x7B/* { */ ||
         c === 0x7D/* } */;
}

function fromHexCode(c) {
  var lc;

  if ((0x30/* 0 */ <= c) && (c <= 0x39/* 9 */)) {
    return c - 0x30;
  }

  /*eslint-disable no-bitwise*/
  lc = c | 0x20;

  if ((0x61/* a */ <= lc) && (lc <= 0x66/* f */)) {
    return lc - 0x61 + 10;
  }

  return -1;
}

function escapedHexLen(c) {
  if (c === 0x78/* x */) { return 2; }
  if (c === 0x75/* u */) { return 4; }
  if (c === 0x55/* U */) { return 8; }
  return 0;
}

function fromDecimalCode(c) {
  if ((0x30/* 0 */ <= c) && (c <= 0x39/* 9 */)) {
    return c - 0x30;
  }

  return -1;
}

function simpleEscapeSequence(c) {
  /* eslint-disable indent */
  return (c === 0x30/* 0 */) ? '\x00' :
        (c === 0x61/* a */) ? '\x07' :
        (c === 0x62/* b */) ? '\x08' :
        (c === 0x74/* t */) ? '\x09' :
        (c === 0x09/* Tab */) ? '\x09' :
        (c === 0x6E/* n */) ? '\x0A' :
        (c === 0x76/* v */) ? '\x0B' :
        (c === 0x66/* f */) ? '\x0C' :
        (c === 0x72/* r */) ? '\x0D' :
        (c === 0x65/* e */) ? '\x1B' :
        (c === 0x20/* Space */) ? ' ' :
        (c === 0x22/* " */) ? '\x22' :
        (c === 0x2F/* / */) ? '/' :
        (c === 0x5C/* \ */) ? '\x5C' :
        (c === 0x4E/* N */) ? '\x85' :
        (c === 0x5F/* _ */) ? '\xA0' :
        (c === 0x4C/* L */) ? '\u2028' :
        (c === 0x50/* P */) ? '\u2029' : '';
}

function charFromCodepoint(c) {
  if (c <= 0xFFFF) {
    return String.fromCharCode(c);
  }
  // Encode UTF-16 surrogate pair
  // https://en.wikipedia.org/wiki/UTF-16#Code_points_U.2B010000_to_U.2B10FFFF
  return String.fromCharCode(
    ((c - 0x010000) >> 10) + 0xD800,
    ((c - 0x010000) & 0x03FF) + 0xDC00
  );
}

var simpleEscapeCheck = new Array(256); // integer, for fast access
var simpleEscapeMap = new Array(256);
for (var i = 0; i < 256; i++) {
  simpleEscapeCheck[i] = simpleEscapeSequence(i) ? 1 : 0;
  simpleEscapeMap[i] = simpleEscapeSequence(i);
}


function State$1(input, options) {
  this.input = input;

  this.filename  = options['filename']  || null;
  this.schema    = options['schema']    || DEFAULT_FULL_SCHEMA$1;
  this.onWarning = options['onWarning'] || null;
  this.legacy    = options['legacy']    || false;
  this.json      = options['json']      || false;
  this.listener  = options['listener']  || null;

  this.implicitTypes = this.schema.compiledImplicit;
  this.typeMap       = this.schema.compiledTypeMap;

  this.length     = input.length;
  this.position   = 0;
  this.line       = 0;
  this.lineStart  = 0;
  this.lineIndent = 0;

  this.documents = [];

  /*
  this.version;
  this.checkLineBreaks;
  this.tagMap;
  this.anchorMap;
  this.tag;
  this.anchor;
  this.kind;
  this.result;*/

}


function generateError(state, message) {
  return new YAMLException$1(
    message,
    new Mark(state.filename, state.input, state.position, state.line, (state.position - state.lineStart)));
}

function throwError(state, message) {
  throw generateError(state, message);
}

function throwWarning(state, message) {
  if (state.onWarning) {
    state.onWarning.call(null, generateError(state, message));
  }
}


var directiveHandlers = {

  YAML: function handleYamlDirective(state, name, args) {

    var match, major, minor;

    if (state.version !== null) {
      throwError(state, 'duplication of %YAML directive');
    }

    if (args.length !== 1) {
      throwError(state, 'YAML directive accepts exactly one argument');
    }

    match = /^([0-9]+)\.([0-9]+)$/.exec(args[0]);

    if (match === null) {
      throwError(state, 'ill-formed argument of the YAML directive');
    }

    major = parseInt(match[1], 10);
    minor = parseInt(match[2], 10);

    if (major !== 1) {
      throwError(state, 'unacceptable YAML version of the document');
    }

    state.version = args[0];
    state.checkLineBreaks = (minor < 2);

    if (minor !== 1 && minor !== 2) {
      throwWarning(state, 'unsupported YAML version of the document');
    }
  },

  TAG: function handleTagDirective(state, name, args) {

    var handle, prefix;

    if (args.length !== 2) {
      throwError(state, 'TAG directive accepts exactly two arguments');
    }

    handle = args[0];
    prefix = args[1];

    if (!PATTERN_TAG_HANDLE.test(handle)) {
      throwError(state, 'ill-formed tag handle (first argument) of the TAG directive');
    }

    if (_hasOwnProperty$1.call(state.tagMap, handle)) {
      throwError(state, 'there is a previously declared suffix for "' + handle + '" tag handle');
    }

    if (!PATTERN_TAG_URI.test(prefix)) {
      throwError(state, 'ill-formed tag prefix (second argument) of the TAG directive');
    }

    state.tagMap[handle] = prefix;
  }
};


function captureSegment(state, start, end, checkJson) {
  var _position, _length, _character, _result;

  if (start < end) {
    _result = state.input.slice(start, end);

    if (checkJson) {
      for (_position = 0, _length = _result.length; _position < _length; _position += 1) {
        _character = _result.charCodeAt(_position);
        if (!(_character === 0x09 ||
              (0x20 <= _character && _character <= 0x10FFFF))) {
          throwError(state, 'expected valid JSON character');
        }
      }
    } else if (PATTERN_NON_PRINTABLE.test(_result)) {
      throwError(state, 'the stream contains non-printable characters');
    }

    state.result += _result;
  }
}

function mergeMappings(state, destination, source, overridableKeys) {
  var sourceKeys, key, index, quantity;

  if (!common$1.isObject(source)) {
    throwError(state, 'cannot merge mappings; the provided source object is unacceptable');
  }

  sourceKeys = Object.keys(source);

  for (index = 0, quantity = sourceKeys.length; index < quantity; index += 1) {
    key = sourceKeys[index];

    if (!_hasOwnProperty$1.call(destination, key)) {
      destination[key] = source[key];
      overridableKeys[key] = true;
    }
  }
}

function storeMappingPair(state, _result, overridableKeys, keyTag, keyNode, valueNode, startLine, startPos) {
  var index, quantity;

  // The output is a plain object here, so keys can only be strings.
  // We need to convert keyNode to a string, but doing so can hang the process
  // (deeply nested arrays that explode exponentially using aliases).
  if (Array.isArray(keyNode)) {
    keyNode = Array.prototype.slice.call(keyNode);

    for (index = 0, quantity = keyNode.length; index < quantity; index += 1) {
      if (Array.isArray(keyNode[index])) {
        throwError(state, 'nested arrays are not supported inside keys');
      }

      if (typeof keyNode === 'object' && _class(keyNode[index]) === '[object Object]') {
        keyNode[index] = '[object Object]';
      }
    }
  }

  // Avoid code execution in load() via toString property
  // (still use its own toString for arrays, timestamps,
  // and whatever user schema extensions happen to have @@toStringTag)
  if (typeof keyNode === 'object' && _class(keyNode) === '[object Object]') {
    keyNode = '[object Object]';
  }


  keyNode = String(keyNode);

  if (_result === null) {
    _result = {};
  }

  if (keyTag === 'tag:yaml.org,2002:merge') {
    if (Array.isArray(valueNode)) {
      for (index = 0, quantity = valueNode.length; index < quantity; index += 1) {
        mergeMappings(state, _result, valueNode[index], overridableKeys);
      }
    } else {
      mergeMappings(state, _result, valueNode, overridableKeys);
    }
  } else {
    if (!state.json &&
        !_hasOwnProperty$1.call(overridableKeys, keyNode) &&
        _hasOwnProperty$1.call(_result, keyNode)) {
      state.line = startLine || state.line;
      state.position = startPos || state.position;
      throwError(state, 'duplicated mapping key');
    }
    _result[keyNode] = valueNode;
    delete overridableKeys[keyNode];
  }

  return _result;
}

function readLineBreak(state) {
  var ch;

  ch = state.input.charCodeAt(state.position);

  if (ch === 0x0A/* LF */) {
    state.position++;
  } else if (ch === 0x0D/* CR */) {
    state.position++;
    if (state.input.charCodeAt(state.position) === 0x0A/* LF */) {
      state.position++;
    }
  } else {
    throwError(state, 'a line break is expected');
  }

  state.line += 1;
  state.lineStart = state.position;
}

function skipSeparationSpace(state, allowComments, checkIndent) {
  var lineBreaks = 0,
      ch = state.input.charCodeAt(state.position);

  while (ch !== 0) {
    while (is_WHITE_SPACE(ch)) {
      ch = state.input.charCodeAt(++state.position);
    }

    if (allowComments && ch === 0x23/* # */) {
      do {
        ch = state.input.charCodeAt(++state.position);
      } while (ch !== 0x0A/* LF */ && ch !== 0x0D/* CR */ && ch !== 0);
    }

    if (is_EOL(ch)) {
      readLineBreak(state);

      ch = state.input.charCodeAt(state.position);
      lineBreaks++;
      state.lineIndent = 0;

      while (ch === 0x20/* Space */) {
        state.lineIndent++;
        ch = state.input.charCodeAt(++state.position);
      }
    } else {
      break;
    }
  }

  if (checkIndent !== -1 && lineBreaks !== 0 && state.lineIndent < checkIndent) {
    throwWarning(state, 'deficient indentation');
  }

  return lineBreaks;
}

function testDocumentSeparator(state) {
  var _position = state.position,
      ch;

  ch = state.input.charCodeAt(_position);

  // Condition state.position === state.lineStart is tested
  // in parent on each call, for efficiency. No needs to test here again.
  if ((ch === 0x2D/* - */ || ch === 0x2E/* . */) &&
      ch === state.input.charCodeAt(_position + 1) &&
      ch === state.input.charCodeAt(_position + 2)) {

    _position += 3;

    ch = state.input.charCodeAt(_position);

    if (ch === 0 || is_WS_OR_EOL(ch)) {
      return true;
    }
  }

  return false;
}

function writeFoldedLines(state, count) {
  if (count === 1) {
    state.result += ' ';
  } else if (count > 1) {
    state.result += common$1.repeat('\n', count - 1);
  }
}


function readPlainScalar(state, nodeIndent, withinFlowCollection) {
  var preceding,
      following,
      captureStart,
      captureEnd,
      hasPendingContent,
      _line,
      _lineStart,
      _lineIndent,
      _kind = state.kind,
      _result = state.result,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (is_WS_OR_EOL(ch)      ||
      is_FLOW_INDICATOR(ch) ||
      ch === 0x23/* # */    ||
      ch === 0x26/* & */    ||
      ch === 0x2A/* * */    ||
      ch === 0x21/* ! */    ||
      ch === 0x7C/* | */    ||
      ch === 0x3E/* > */    ||
      ch === 0x27/* ' */    ||
      ch === 0x22/* " */    ||
      ch === 0x25/* % */    ||
      ch === 0x40/* @ */    ||
      ch === 0x60/* ` */) {
    return false;
  }

  if (ch === 0x3F/* ? */ || ch === 0x2D/* - */) {
    following = state.input.charCodeAt(state.position + 1);

    if (is_WS_OR_EOL(following) ||
        withinFlowCollection && is_FLOW_INDICATOR(following)) {
      return false;
    }
  }

  state.kind = 'scalar';
  state.result = '';
  captureStart = captureEnd = state.position;
  hasPendingContent = false;

  while (ch !== 0) {
    if (ch === 0x3A/* : */) {
      following = state.input.charCodeAt(state.position + 1);

      if (is_WS_OR_EOL(following) ||
          withinFlowCollection && is_FLOW_INDICATOR(following)) {
        break;
      }

    } else if (ch === 0x23/* # */) {
      preceding = state.input.charCodeAt(state.position - 1);

      if (is_WS_OR_EOL(preceding)) {
        break;
      }

    } else if ((state.position === state.lineStart && testDocumentSeparator(state)) ||
               withinFlowCollection && is_FLOW_INDICATOR(ch)) {
      break;

    } else if (is_EOL(ch)) {
      _line = state.line;
      _lineStart = state.lineStart;
      _lineIndent = state.lineIndent;
      skipSeparationSpace(state, false, -1);

      if (state.lineIndent >= nodeIndent) {
        hasPendingContent = true;
        ch = state.input.charCodeAt(state.position);
        continue;
      } else {
        state.position = captureEnd;
        state.line = _line;
        state.lineStart = _lineStart;
        state.lineIndent = _lineIndent;
        break;
      }
    }

    if (hasPendingContent) {
      captureSegment(state, captureStart, captureEnd, false);
      writeFoldedLines(state, state.line - _line);
      captureStart = captureEnd = state.position;
      hasPendingContent = false;
    }

    if (!is_WHITE_SPACE(ch)) {
      captureEnd = state.position + 1;
    }

    ch = state.input.charCodeAt(++state.position);
  }

  captureSegment(state, captureStart, captureEnd, false);

  if (state.result) {
    return true;
  }

  state.kind = _kind;
  state.result = _result;
  return false;
}

function readSingleQuotedScalar(state, nodeIndent) {
  var ch,
      captureStart, captureEnd;

  ch = state.input.charCodeAt(state.position);

  if (ch !== 0x27/* ' */) {
    return false;
  }

  state.kind = 'scalar';
  state.result = '';
  state.position++;
  captureStart = captureEnd = state.position;

  while ((ch = state.input.charCodeAt(state.position)) !== 0) {
    if (ch === 0x27/* ' */) {
      captureSegment(state, captureStart, state.position, true);
      ch = state.input.charCodeAt(++state.position);

      if (ch === 0x27/* ' */) {
        captureStart = state.position;
        state.position++;
        captureEnd = state.position;
      } else {
        return true;
      }

    } else if (is_EOL(ch)) {
      captureSegment(state, captureStart, captureEnd, true);
      writeFoldedLines(state, skipSeparationSpace(state, false, nodeIndent));
      captureStart = captureEnd = state.position;

    } else if (state.position === state.lineStart && testDocumentSeparator(state)) {
      throwError(state, 'unexpected end of the document within a single quoted scalar');

    } else {
      state.position++;
      captureEnd = state.position;
    }
  }

  throwError(state, 'unexpected end of the stream within a single quoted scalar');
}

function readDoubleQuotedScalar(state, nodeIndent) {
  var captureStart,
      captureEnd,
      hexLength,
      hexResult,
      tmp,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (ch !== 0x22/* " */) {
    return false;
  }

  state.kind = 'scalar';
  state.result = '';
  state.position++;
  captureStart = captureEnd = state.position;

  while ((ch = state.input.charCodeAt(state.position)) !== 0) {
    if (ch === 0x22/* " */) {
      captureSegment(state, captureStart, state.position, true);
      state.position++;
      return true;

    } else if (ch === 0x5C/* \ */) {
      captureSegment(state, captureStart, state.position, true);
      ch = state.input.charCodeAt(++state.position);

      if (is_EOL(ch)) {
        skipSeparationSpace(state, false, nodeIndent);

        // TODO: rework to inline fn with no type cast?
      } else if (ch < 256 && simpleEscapeCheck[ch]) {
        state.result += simpleEscapeMap[ch];
        state.position++;

      } else if ((tmp = escapedHexLen(ch)) > 0) {
        hexLength = tmp;
        hexResult = 0;

        for (; hexLength > 0; hexLength--) {
          ch = state.input.charCodeAt(++state.position);

          if ((tmp = fromHexCode(ch)) >= 0) {
            hexResult = (hexResult << 4) + tmp;

          } else {
            throwError(state, 'expected hexadecimal character');
          }
        }

        state.result += charFromCodepoint(hexResult);

        state.position++;

      } else {
        throwError(state, 'unknown escape sequence');
      }

      captureStart = captureEnd = state.position;

    } else if (is_EOL(ch)) {
      captureSegment(state, captureStart, captureEnd, true);
      writeFoldedLines(state, skipSeparationSpace(state, false, nodeIndent));
      captureStart = captureEnd = state.position;

    } else if (state.position === state.lineStart && testDocumentSeparator(state)) {
      throwError(state, 'unexpected end of the document within a double quoted scalar');

    } else {
      state.position++;
      captureEnd = state.position;
    }
  }

  throwError(state, 'unexpected end of the stream within a double quoted scalar');
}

function readFlowCollection(state, nodeIndent) {
  var readNext = true,
      _line,
      _tag     = state.tag,
      _result,
      _anchor  = state.anchor,
      following,
      terminator,
      isPair,
      isExplicitPair,
      isMapping,
      overridableKeys = {},
      keyNode,
      keyTag,
      valueNode,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (ch === 0x5B/* [ */) {
    terminator = 0x5D;/* ] */
    isMapping = false;
    _result = [];
  } else if (ch === 0x7B/* { */) {
    terminator = 0x7D;/* } */
    isMapping = true;
    _result = {};
  } else {
    return false;
  }

  if (state.anchor !== null) {
    state.anchorMap[state.anchor] = _result;
  }

  ch = state.input.charCodeAt(++state.position);

  while (ch !== 0) {
    skipSeparationSpace(state, true, nodeIndent);

    ch = state.input.charCodeAt(state.position);

    if (ch === terminator) {
      state.position++;
      state.tag = _tag;
      state.anchor = _anchor;
      state.kind = isMapping ? 'mapping' : 'sequence';
      state.result = _result;
      return true;
    } else if (!readNext) {
      throwError(state, 'missed comma between flow collection entries');
    }

    keyTag = keyNode = valueNode = null;
    isPair = isExplicitPair = false;

    if (ch === 0x3F/* ? */) {
      following = state.input.charCodeAt(state.position + 1);

      if (is_WS_OR_EOL(following)) {
        isPair = isExplicitPair = true;
        state.position++;
        skipSeparationSpace(state, true, nodeIndent);
      }
    }

    _line = state.line;
    composeNode(state, nodeIndent, CONTEXT_FLOW_IN, false, true);
    keyTag = state.tag;
    keyNode = state.result;
    skipSeparationSpace(state, true, nodeIndent);

    ch = state.input.charCodeAt(state.position);

    if ((isExplicitPair || state.line === _line) && ch === 0x3A/* : */) {
      isPair = true;
      ch = state.input.charCodeAt(++state.position);
      skipSeparationSpace(state, true, nodeIndent);
      composeNode(state, nodeIndent, CONTEXT_FLOW_IN, false, true);
      valueNode = state.result;
    }

    if (isMapping) {
      storeMappingPair(state, _result, overridableKeys, keyTag, keyNode, valueNode);
    } else if (isPair) {
      _result.push(storeMappingPair(state, null, overridableKeys, keyTag, keyNode, valueNode));
    } else {
      _result.push(keyNode);
    }

    skipSeparationSpace(state, true, nodeIndent);

    ch = state.input.charCodeAt(state.position);

    if (ch === 0x2C/* , */) {
      readNext = true;
      ch = state.input.charCodeAt(++state.position);
    } else {
      readNext = false;
    }
  }

  throwError(state, 'unexpected end of the stream within a flow collection');
}

function readBlockScalar(state, nodeIndent) {
  var captureStart,
      folding,
      chomping       = CHOMPING_CLIP,
      didReadContent = false,
      detectedIndent = false,
      textIndent     = nodeIndent,
      emptyLines     = 0,
      atMoreIndented = false,
      tmp,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (ch === 0x7C/* | */) {
    folding = false;
  } else if (ch === 0x3E/* > */) {
    folding = true;
  } else {
    return false;
  }

  state.kind = 'scalar';
  state.result = '';

  while (ch !== 0) {
    ch = state.input.charCodeAt(++state.position);

    if (ch === 0x2B/* + */ || ch === 0x2D/* - */) {
      if (CHOMPING_CLIP === chomping) {
        chomping = (ch === 0x2B/* + */) ? CHOMPING_KEEP : CHOMPING_STRIP;
      } else {
        throwError(state, 'repeat of a chomping mode identifier');
      }

    } else if ((tmp = fromDecimalCode(ch)) >= 0) {
      if (tmp === 0) {
        throwError(state, 'bad explicit indentation width of a block scalar; it cannot be less than one');
      } else if (!detectedIndent) {
        textIndent = nodeIndent + tmp - 1;
        detectedIndent = true;
      } else {
        throwError(state, 'repeat of an indentation width identifier');
      }

    } else {
      break;
    }
  }

  if (is_WHITE_SPACE(ch)) {
    do { ch = state.input.charCodeAt(++state.position); }
    while (is_WHITE_SPACE(ch));

    if (ch === 0x23/* # */) {
      do { ch = state.input.charCodeAt(++state.position); }
      while (!is_EOL(ch) && (ch !== 0));
    }
  }

  while (ch !== 0) {
    readLineBreak(state);
    state.lineIndent = 0;

    ch = state.input.charCodeAt(state.position);

    while ((!detectedIndent || state.lineIndent < textIndent) &&
           (ch === 0x20/* Space */)) {
      state.lineIndent++;
      ch = state.input.charCodeAt(++state.position);
    }

    if (!detectedIndent && state.lineIndent > textIndent) {
      textIndent = state.lineIndent;
    }

    if (is_EOL(ch)) {
      emptyLines++;
      continue;
    }

    // End of the scalar.
    if (state.lineIndent < textIndent) {

      // Perform the chomping.
      if (chomping === CHOMPING_KEEP) {
        state.result += common$1.repeat('\n', didReadContent ? 1 + emptyLines : emptyLines);
      } else if (chomping === CHOMPING_CLIP) {
        if (didReadContent) { // i.e. only if the scalar is not empty.
          state.result += '\n';
        }
      }

      // Break this `while` cycle and go to the funciton's epilogue.
      break;
    }

    // Folded style: use fancy rules to handle line breaks.
    if (folding) {

      // Lines starting with white space characters (more-indented lines) are not folded.
      if (is_WHITE_SPACE(ch)) {
        atMoreIndented = true;
        // except for the first content line (cf. Example 8.1)
        state.result += common$1.repeat('\n', didReadContent ? 1 + emptyLines : emptyLines);

      // End of more-indented block.
      } else if (atMoreIndented) {
        atMoreIndented = false;
        state.result += common$1.repeat('\n', emptyLines + 1);

      // Just one line break - perceive as the same line.
      } else if (emptyLines === 0) {
        if (didReadContent) { // i.e. only if we have already read some scalar content.
          state.result += ' ';
        }

      // Several line breaks - perceive as different lines.
      } else {
        state.result += common$1.repeat('\n', emptyLines);
      }

    // Literal style: just add exact number of line breaks between content lines.
    } else {
      // Keep all line breaks except the header line break.
      state.result += common$1.repeat('\n', didReadContent ? 1 + emptyLines : emptyLines);
    }

    didReadContent = true;
    detectedIndent = true;
    emptyLines = 0;
    captureStart = state.position;

    while (!is_EOL(ch) && (ch !== 0)) {
      ch = state.input.charCodeAt(++state.position);
    }

    captureSegment(state, captureStart, state.position, false);
  }

  return true;
}

function readBlockSequence(state, nodeIndent) {
  var _line,
      _tag      = state.tag,
      _anchor   = state.anchor,
      _result   = [],
      following,
      detected  = false,
      ch;

  if (state.anchor !== null) {
    state.anchorMap[state.anchor] = _result;
  }

  ch = state.input.charCodeAt(state.position);

  while (ch !== 0) {

    if (ch !== 0x2D/* - */) {
      break;
    }

    following = state.input.charCodeAt(state.position + 1);

    if (!is_WS_OR_EOL(following)) {
      break;
    }

    detected = true;
    state.position++;

    if (skipSeparationSpace(state, true, -1)) {
      if (state.lineIndent <= nodeIndent) {
        _result.push(null);
        ch = state.input.charCodeAt(state.position);
        continue;
      }
    }

    _line = state.line;
    composeNode(state, nodeIndent, CONTEXT_BLOCK_IN, false, true);
    _result.push(state.result);
    skipSeparationSpace(state, true, -1);

    ch = state.input.charCodeAt(state.position);

    if ((state.line === _line || state.lineIndent > nodeIndent) && (ch !== 0)) {
      throwError(state, 'bad indentation of a sequence entry');
    } else if (state.lineIndent < nodeIndent) {
      break;
    }
  }

  if (detected) {
    state.tag = _tag;
    state.anchor = _anchor;
    state.kind = 'sequence';
    state.result = _result;
    return true;
  }
  return false;
}

function readBlockMapping(state, nodeIndent, flowIndent) {
  var following,
      allowCompact,
      _line,
      _pos,
      _tag          = state.tag,
      _anchor       = state.anchor,
      _result       = {},
      overridableKeys = {},
      keyTag        = null,
      keyNode       = null,
      valueNode     = null,
      atExplicitKey = false,
      detected      = false,
      ch;

  if (state.anchor !== null) {
    state.anchorMap[state.anchor] = _result;
  }

  ch = state.input.charCodeAt(state.position);

  while (ch !== 0) {
    following = state.input.charCodeAt(state.position + 1);
    _line = state.line; // Save the current line.
    _pos = state.position;

    //
    // Explicit notation case. There are two separate blocks:
    // first for the key (denoted by "?") and second for the value (denoted by ":")
    //
    if ((ch === 0x3F/* ? */ || ch === 0x3A/* : */) && is_WS_OR_EOL(following)) {

      if (ch === 0x3F/* ? */) {
        if (atExplicitKey) {
          storeMappingPair(state, _result, overridableKeys, keyTag, keyNode, null);
          keyTag = keyNode = valueNode = null;
        }

        detected = true;
        atExplicitKey = true;
        allowCompact = true;

      } else if (atExplicitKey) {
        // i.e. 0x3A/* : */ === character after the explicit key.
        atExplicitKey = false;
        allowCompact = true;

      } else {
        throwError(state, 'incomplete explicit mapping pair; a key node is missed; or followed by a non-tabulated empty line');
      }

      state.position += 1;
      ch = following;

    //
    // Implicit notation case. Flow-style node as the key first, then ":", and the value.
    //
    } else if (composeNode(state, flowIndent, CONTEXT_FLOW_OUT, false, true)) {

      if (state.line === _line) {
        ch = state.input.charCodeAt(state.position);

        while (is_WHITE_SPACE(ch)) {
          ch = state.input.charCodeAt(++state.position);
        }

        if (ch === 0x3A/* : */) {
          ch = state.input.charCodeAt(++state.position);

          if (!is_WS_OR_EOL(ch)) {
            throwError(state, 'a whitespace character is expected after the key-value separator within a block mapping');
          }

          if (atExplicitKey) {
            storeMappingPair(state, _result, overridableKeys, keyTag, keyNode, null);
            keyTag = keyNode = valueNode = null;
          }

          detected = true;
          atExplicitKey = false;
          allowCompact = false;
          keyTag = state.tag;
          keyNode = state.result;

        } else if (detected) {
          throwError(state, 'can not read an implicit mapping pair; a colon is missed');

        } else {
          state.tag = _tag;
          state.anchor = _anchor;
          return true; // Keep the result of `composeNode`.
        }

      } else if (detected) {
        throwError(state, 'can not read a block mapping entry; a multiline key may not be an implicit key');

      } else {
        state.tag = _tag;
        state.anchor = _anchor;
        return true; // Keep the result of `composeNode`.
      }

    } else {
      break; // Reading is done. Go to the epilogue.
    }

    //
    // Common reading code for both explicit and implicit notations.
    //
    if (state.line === _line || state.lineIndent > nodeIndent) {
      if (composeNode(state, nodeIndent, CONTEXT_BLOCK_OUT, true, allowCompact)) {
        if (atExplicitKey) {
          keyNode = state.result;
        } else {
          valueNode = state.result;
        }
      }

      if (!atExplicitKey) {
        storeMappingPair(state, _result, overridableKeys, keyTag, keyNode, valueNode, _line, _pos);
        keyTag = keyNode = valueNode = null;
      }

      skipSeparationSpace(state, true, -1);
      ch = state.input.charCodeAt(state.position);
    }

    if (state.lineIndent > nodeIndent && (ch !== 0)) {
      throwError(state, 'bad indentation of a mapping entry');
    } else if (state.lineIndent < nodeIndent) {
      break;
    }
  }

  //
  // Epilogue.
  //

  // Special case: last mapping's node contains only the key in explicit notation.
  if (atExplicitKey) {
    storeMappingPair(state, _result, overridableKeys, keyTag, keyNode, null);
  }

  // Expose the resulting mapping.
  if (detected) {
    state.tag = _tag;
    state.anchor = _anchor;
    state.kind = 'mapping';
    state.result = _result;
  }

  return detected;
}

function readTagProperty(state) {
  var _position,
      isVerbatim = false,
      isNamed    = false,
      tagHandle,
      tagName,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (ch !== 0x21/* ! */) return false;

  if (state.tag !== null) {
    throwError(state, 'duplication of a tag property');
  }

  ch = state.input.charCodeAt(++state.position);

  if (ch === 0x3C/* < */) {
    isVerbatim = true;
    ch = state.input.charCodeAt(++state.position);

  } else if (ch === 0x21/* ! */) {
    isNamed = true;
    tagHandle = '!!';
    ch = state.input.charCodeAt(++state.position);

  } else {
    tagHandle = '!';
  }

  _position = state.position;

  if (isVerbatim) {
    do { ch = state.input.charCodeAt(++state.position); }
    while (ch !== 0 && ch !== 0x3E/* > */);

    if (state.position < state.length) {
      tagName = state.input.slice(_position, state.position);
      ch = state.input.charCodeAt(++state.position);
    } else {
      throwError(state, 'unexpected end of the stream within a verbatim tag');
    }
  } else {
    while (ch !== 0 && !is_WS_OR_EOL(ch)) {

      if (ch === 0x21/* ! */) {
        if (!isNamed) {
          tagHandle = state.input.slice(_position - 1, state.position + 1);

          if (!PATTERN_TAG_HANDLE.test(tagHandle)) {
            throwError(state, 'named tag handle cannot contain such characters');
          }

          isNamed = true;
          _position = state.position + 1;
        } else {
          throwError(state, 'tag suffix cannot contain exclamation marks');
        }
      }

      ch = state.input.charCodeAt(++state.position);
    }

    tagName = state.input.slice(_position, state.position);

    if (PATTERN_FLOW_INDICATORS.test(tagName)) {
      throwError(state, 'tag suffix cannot contain flow indicator characters');
    }
  }

  if (tagName && !PATTERN_TAG_URI.test(tagName)) {
    throwError(state, 'tag name cannot contain such characters: ' + tagName);
  }

  if (isVerbatim) {
    state.tag = tagName;

  } else if (_hasOwnProperty$1.call(state.tagMap, tagHandle)) {
    state.tag = state.tagMap[tagHandle] + tagName;

  } else if (tagHandle === '!') {
    state.tag = '!' + tagName;

  } else if (tagHandle === '!!') {
    state.tag = 'tag:yaml.org,2002:' + tagName;

  } else {
    throwError(state, 'undeclared tag handle "' + tagHandle + '"');
  }

  return true;
}

function readAnchorProperty(state) {
  var _position,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (ch !== 0x26/* & */) return false;

  if (state.anchor !== null) {
    throwError(state, 'duplication of an anchor property');
  }

  ch = state.input.charCodeAt(++state.position);
  _position = state.position;

  while (ch !== 0 && !is_WS_OR_EOL(ch) && !is_FLOW_INDICATOR(ch)) {
    ch = state.input.charCodeAt(++state.position);
  }

  if (state.position === _position) {
    throwError(state, 'name of an anchor node must contain at least one character');
  }

  state.anchor = state.input.slice(_position, state.position);
  return true;
}

function readAlias(state) {
  var _position, alias,
      ch;

  ch = state.input.charCodeAt(state.position);

  if (ch !== 0x2A/* * */) return false;

  ch = state.input.charCodeAt(++state.position);
  _position = state.position;

  while (ch !== 0 && !is_WS_OR_EOL(ch) && !is_FLOW_INDICATOR(ch)) {
    ch = state.input.charCodeAt(++state.position);
  }

  if (state.position === _position) {
    throwError(state, 'name of an alias node must contain at least one character');
  }

  alias = state.input.slice(_position, state.position);

  if (!state.anchorMap.hasOwnProperty(alias)) {
    throwError(state, 'unidentified alias "' + alias + '"');
  }

  state.result = state.anchorMap[alias];
  skipSeparationSpace(state, true, -1);
  return true;
}

function composeNode(state, parentIndent, nodeContext, allowToSeek, allowCompact) {
  var allowBlockStyles,
      allowBlockScalars,
      allowBlockCollections,
      indentStatus = 1, // 1: this>parent, 0: this=parent, -1: this<parent
      atNewLine  = false,
      hasContent = false,
      typeIndex,
      typeQuantity,
      type,
      flowIndent,
      blockIndent;

  if (state.listener !== null) {
    state.listener('open', state);
  }

  state.tag    = null;
  state.anchor = null;
  state.kind   = null;
  state.result = null;

  allowBlockStyles = allowBlockScalars = allowBlockCollections =
    CONTEXT_BLOCK_OUT === nodeContext ||
    CONTEXT_BLOCK_IN  === nodeContext;

  if (allowToSeek) {
    if (skipSeparationSpace(state, true, -1)) {
      atNewLine = true;

      if (state.lineIndent > parentIndent) {
        indentStatus = 1;
      } else if (state.lineIndent === parentIndent) {
        indentStatus = 0;
      } else if (state.lineIndent < parentIndent) {
        indentStatus = -1;
      }
    }
  }

  if (indentStatus === 1) {
    while (readTagProperty(state) || readAnchorProperty(state)) {
      if (skipSeparationSpace(state, true, -1)) {
        atNewLine = true;
        allowBlockCollections = allowBlockStyles;

        if (state.lineIndent > parentIndent) {
          indentStatus = 1;
        } else if (state.lineIndent === parentIndent) {
          indentStatus = 0;
        } else if (state.lineIndent < parentIndent) {
          indentStatus = -1;
        }
      } else {
        allowBlockCollections = false;
      }
    }
  }

  if (allowBlockCollections) {
    allowBlockCollections = atNewLine || allowCompact;
  }

  if (indentStatus === 1 || CONTEXT_BLOCK_OUT === nodeContext) {
    if (CONTEXT_FLOW_IN === nodeContext || CONTEXT_FLOW_OUT === nodeContext) {
      flowIndent = parentIndent;
    } else {
      flowIndent = parentIndent + 1;
    }

    blockIndent = state.position - state.lineStart;

    if (indentStatus === 1) {
      if (allowBlockCollections &&
          (readBlockSequence(state, blockIndent) ||
           readBlockMapping(state, blockIndent, flowIndent)) ||
          readFlowCollection(state, flowIndent)) {
        hasContent = true;
      } else {
        if ((allowBlockScalars && readBlockScalar(state, flowIndent)) ||
            readSingleQuotedScalar(state, flowIndent) ||
            readDoubleQuotedScalar(state, flowIndent)) {
          hasContent = true;

        } else if (readAlias(state)) {
          hasContent = true;

          if (state.tag !== null || state.anchor !== null) {
            throwError(state, 'alias node should not have any properties');
          }

        } else if (readPlainScalar(state, flowIndent, CONTEXT_FLOW_IN === nodeContext)) {
          hasContent = true;

          if (state.tag === null) {
            state.tag = '?';
          }
        }

        if (state.anchor !== null) {
          state.anchorMap[state.anchor] = state.result;
        }
      }
    } else if (indentStatus === 0) {
      // Special case: block sequences are allowed to have same indentation level as the parent.
      // http://www.yaml.org/spec/1.2/spec.html#id2799784
      hasContent = allowBlockCollections && readBlockSequence(state, blockIndent);
    }
  }

  if (state.tag !== null && state.tag !== '!') {
    if (state.tag === '?') {
      // Implicit resolving is not allowed for non-scalar types, and '?'
      // non-specific tag is only automatically assigned to plain scalars.
      //
      // We only need to check kind conformity in case user explicitly assigns '?'
      // tag, for example like this: "!<?> [0]"
      //
      if (state.result !== null && state.kind !== 'scalar') {
        throwError(state, 'unacceptable node kind for !<?> tag; it should be "scalar", not "' + state.kind + '"');
      }

      for (typeIndex = 0, typeQuantity = state.implicitTypes.length; typeIndex < typeQuantity; typeIndex += 1) {
        type = state.implicitTypes[typeIndex];

        if (type.resolve(state.result)) { // `state.result` updated in resolver if matched
          state.result = type.construct(state.result);
          state.tag = type.tag;
          if (state.anchor !== null) {
            state.anchorMap[state.anchor] = state.result;
          }
          break;
        }
      }
    } else if (_hasOwnProperty$1.call(state.typeMap[state.kind || 'fallback'], state.tag)) {
      type = state.typeMap[state.kind || 'fallback'][state.tag];

      if (state.result !== null && type.kind !== state.kind) {
        throwError(state, 'unacceptable node kind for !<' + state.tag + '> tag; it should be "' + type.kind + '", not "' + state.kind + '"');
      }

      if (!type.resolve(state.result)) { // `state.result` updated in resolver if matched
        throwError(state, 'cannot resolve a node with !<' + state.tag + '> explicit tag');
      } else {
        state.result = type.construct(state.result);
        if (state.anchor !== null) {
          state.anchorMap[state.anchor] = state.result;
        }
      }
    } else {
      throwError(state, 'unknown tag !<' + state.tag + '>');
    }
  }

  if (state.listener !== null) {
    state.listener('close', state);
  }
  return state.tag !== null ||  state.anchor !== null || hasContent;
}

function readDocument(state) {
  var documentStart = state.position,
      _position,
      directiveName,
      directiveArgs,
      hasDirectives = false,
      ch;

  state.version = null;
  state.checkLineBreaks = state.legacy;
  state.tagMap = {};
  state.anchorMap = {};

  while ((ch = state.input.charCodeAt(state.position)) !== 0) {
    skipSeparationSpace(state, true, -1);

    ch = state.input.charCodeAt(state.position);

    if (state.lineIndent > 0 || ch !== 0x25/* % */) {
      break;
    }

    hasDirectives = true;
    ch = state.input.charCodeAt(++state.position);
    _position = state.position;

    while (ch !== 0 && !is_WS_OR_EOL(ch)) {
      ch = state.input.charCodeAt(++state.position);
    }

    directiveName = state.input.slice(_position, state.position);
    directiveArgs = [];

    if (directiveName.length < 1) {
      throwError(state, 'directive name must not be less than one character in length');
    }

    while (ch !== 0) {
      while (is_WHITE_SPACE(ch)) {
        ch = state.input.charCodeAt(++state.position);
      }

      if (ch === 0x23/* # */) {
        do { ch = state.input.charCodeAt(++state.position); }
        while (ch !== 0 && !is_EOL(ch));
        break;
      }

      if (is_EOL(ch)) break;

      _position = state.position;

      while (ch !== 0 && !is_WS_OR_EOL(ch)) {
        ch = state.input.charCodeAt(++state.position);
      }

      directiveArgs.push(state.input.slice(_position, state.position));
    }

    if (ch !== 0) readLineBreak(state);

    if (_hasOwnProperty$1.call(directiveHandlers, directiveName)) {
      directiveHandlers[directiveName](state, directiveName, directiveArgs);
    } else {
      throwWarning(state, 'unknown document directive "' + directiveName + '"');
    }
  }

  skipSeparationSpace(state, true, -1);

  if (state.lineIndent === 0 &&
      state.input.charCodeAt(state.position)     === 0x2D/* - */ &&
      state.input.charCodeAt(state.position + 1) === 0x2D/* - */ &&
      state.input.charCodeAt(state.position + 2) === 0x2D/* - */) {
    state.position += 3;
    skipSeparationSpace(state, true, -1);

  } else if (hasDirectives) {
    throwError(state, 'directives end mark is expected');
  }

  composeNode(state, state.lineIndent - 1, CONTEXT_BLOCK_OUT, false, true);
  skipSeparationSpace(state, true, -1);

  if (state.checkLineBreaks &&
      PATTERN_NON_ASCII_LINE_BREAKS.test(state.input.slice(documentStart, state.position))) {
    throwWarning(state, 'non-ASCII line breaks are interpreted as content');
  }

  state.documents.push(state.result);

  if (state.position === state.lineStart && testDocumentSeparator(state)) {

    if (state.input.charCodeAt(state.position) === 0x2E/* . */) {
      state.position += 3;
      skipSeparationSpace(state, true, -1);
    }
    return;
  }

  if (state.position < (state.length - 1)) {
    throwError(state, 'end of the stream or a document separator is expected');
  } else {
    return;
  }
}


function loadDocuments(input, options) {
  input = String(input);
  options = options || {};

  if (input.length !== 0) {

    // Add tailing `\n` if not exists
    if (input.charCodeAt(input.length - 1) !== 0x0A/* LF */ &&
        input.charCodeAt(input.length - 1) !== 0x0D/* CR */) {
      input += '\n';
    }

    // Strip BOM
    if (input.charCodeAt(0) === 0xFEFF) {
      input = input.slice(1);
    }
  }

  var state = new State$1(input, options);

  var nullpos = input.indexOf('\0');

  if (nullpos !== -1) {
    state.position = nullpos;
    throwError(state, 'null byte is not allowed in input');
  }

  // Use 0 as string terminator. That significantly simplifies bounds check.
  state.input += '\0';

  while (state.input.charCodeAt(state.position) === 0x20/* Space */) {
    state.lineIndent += 1;
    state.position += 1;
  }

  while (state.position < (state.length - 1)) {
    readDocument(state);
  }

  return state.documents;
}


function loadAll(input, iterator, options) {
  if (iterator !== null && typeof iterator === 'object' && typeof options === 'undefined') {
    options = iterator;
    iterator = null;
  }

  var documents = loadDocuments(input, options);

  if (typeof iterator !== 'function') {
    return documents;
  }

  for (var index = 0, length = documents.length; index < length; index += 1) {
    iterator(documents[index]);
  }
}


function load(input, options) {
  var documents = loadDocuments(input, options);

  if (documents.length === 0) {
    /*eslint-disable no-undefined*/
    return undefined;
  } else if (documents.length === 1) {
    return documents[0];
  }
  throw new YAMLException$1('expected a single document in the stream, but found more');
}


function safeLoadAll(input, iterator, options) {
  if (typeof iterator === 'object' && iterator !== null && typeof options === 'undefined') {
    options = iterator;
    iterator = null;
  }

  return loadAll(input, iterator, common$1.extend({ schema: DEFAULT_SAFE_SCHEMA$1 }, options));
}


function safeLoad(input, options) {
  return load(input, common$1.extend({ schema: DEFAULT_SAFE_SCHEMA$1 }, options));
}


loader$1.loadAll     = loadAll;
loader$1.load        = load;
loader$1.safeLoadAll = safeLoadAll;
loader$1.safeLoad    = safeLoad;

var dumper$1 = {};

/*eslint-disable no-use-before-define*/

var common              = common$6;
var YAMLException       = exception;
var DEFAULT_FULL_SCHEMA = default_full;
var DEFAULT_SAFE_SCHEMA = default_safe;

var _toString       = Object.prototype.toString;
var _hasOwnProperty = Object.prototype.hasOwnProperty;

var CHAR_TAB                  = 0x09; /* Tab */
var CHAR_LINE_FEED            = 0x0A; /* LF */
var CHAR_CARRIAGE_RETURN      = 0x0D; /* CR */
var CHAR_SPACE                = 0x20; /* Space */
var CHAR_EXCLAMATION          = 0x21; /* ! */
var CHAR_DOUBLE_QUOTE         = 0x22; /* " */
var CHAR_SHARP                = 0x23; /* # */
var CHAR_PERCENT              = 0x25; /* % */
var CHAR_AMPERSAND            = 0x26; /* & */
var CHAR_SINGLE_QUOTE         = 0x27; /* ' */
var CHAR_ASTERISK             = 0x2A; /* * */
var CHAR_COMMA                = 0x2C; /* , */
var CHAR_MINUS                = 0x2D; /* - */
var CHAR_COLON                = 0x3A; /* : */
var CHAR_EQUALS               = 0x3D; /* = */
var CHAR_GREATER_THAN         = 0x3E; /* > */
var CHAR_QUESTION             = 0x3F; /* ? */
var CHAR_COMMERCIAL_AT        = 0x40; /* @ */
var CHAR_LEFT_SQUARE_BRACKET  = 0x5B; /* [ */
var CHAR_RIGHT_SQUARE_BRACKET = 0x5D; /* ] */
var CHAR_GRAVE_ACCENT         = 0x60; /* ` */
var CHAR_LEFT_CURLY_BRACKET   = 0x7B; /* { */
var CHAR_VERTICAL_LINE        = 0x7C; /* | */
var CHAR_RIGHT_CURLY_BRACKET  = 0x7D; /* } */

var ESCAPE_SEQUENCES = {};

ESCAPE_SEQUENCES[0x00]   = '\\0';
ESCAPE_SEQUENCES[0x07]   = '\\a';
ESCAPE_SEQUENCES[0x08]   = '\\b';
ESCAPE_SEQUENCES[0x09]   = '\\t';
ESCAPE_SEQUENCES[0x0A]   = '\\n';
ESCAPE_SEQUENCES[0x0B]   = '\\v';
ESCAPE_SEQUENCES[0x0C]   = '\\f';
ESCAPE_SEQUENCES[0x0D]   = '\\r';
ESCAPE_SEQUENCES[0x1B]   = '\\e';
ESCAPE_SEQUENCES[0x22]   = '\\"';
ESCAPE_SEQUENCES[0x5C]   = '\\\\';
ESCAPE_SEQUENCES[0x85]   = '\\N';
ESCAPE_SEQUENCES[0xA0]   = '\\_';
ESCAPE_SEQUENCES[0x2028] = '\\L';
ESCAPE_SEQUENCES[0x2029] = '\\P';

var DEPRECATED_BOOLEANS_SYNTAX = [
  'y', 'Y', 'yes', 'Yes', 'YES', 'on', 'On', 'ON',
  'n', 'N', 'no', 'No', 'NO', 'off', 'Off', 'OFF'
];

function compileStyleMap(schema, map) {
  var result, keys, index, length, tag, style, type;

  if (map === null) return {};

  result = {};
  keys = Object.keys(map);

  for (index = 0, length = keys.length; index < length; index += 1) {
    tag = keys[index];
    style = String(map[tag]);

    if (tag.slice(0, 2) === '!!') {
      tag = 'tag:yaml.org,2002:' + tag.slice(2);
    }
    type = schema.compiledTypeMap['fallback'][tag];

    if (type && _hasOwnProperty.call(type.styleAliases, style)) {
      style = type.styleAliases[style];
    }

    result[tag] = style;
  }

  return result;
}

function encodeHex(character) {
  var string, handle, length;

  string = character.toString(16).toUpperCase();

  if (character <= 0xFF) {
    handle = 'x';
    length = 2;
  } else if (character <= 0xFFFF) {
    handle = 'u';
    length = 4;
  } else if (character <= 0xFFFFFFFF) {
    handle = 'U';
    length = 8;
  } else {
    throw new YAMLException('code point within a string may not be greater than 0xFFFFFFFF');
  }

  return '\\' + handle + common.repeat('0', length - string.length) + string;
}

function State(options) {
  this.schema        = options['schema'] || DEFAULT_FULL_SCHEMA;
  this.indent        = Math.max(1, (options['indent'] || 2));
  this.noArrayIndent = options['noArrayIndent'] || false;
  this.skipInvalid   = options['skipInvalid'] || false;
  this.flowLevel     = (common.isNothing(options['flowLevel']) ? -1 : options['flowLevel']);
  this.styleMap      = compileStyleMap(this.schema, options['styles'] || null);
  this.sortKeys      = options['sortKeys'] || false;
  this.lineWidth     = options['lineWidth'] || 80;
  this.noRefs        = options['noRefs'] || false;
  this.noCompatMode  = options['noCompatMode'] || false;
  this.condenseFlow  = options['condenseFlow'] || false;

  this.implicitTypes = this.schema.compiledImplicit;
  this.explicitTypes = this.schema.compiledExplicit;

  this.tag = null;
  this.result = '';

  this.duplicates = [];
  this.usedDuplicates = null;
}

// Indents every line in a string. Empty lines (\n only) are not indented.
function indentString(string, spaces) {
  var ind = common.repeat(' ', spaces),
      position = 0,
      next = -1,
      result = '',
      line,
      length = string.length;

  while (position < length) {
    next = string.indexOf('\n', position);
    if (next === -1) {
      line = string.slice(position);
      position = length;
    } else {
      line = string.slice(position, next + 1);
      position = next + 1;
    }

    if (line.length && line !== '\n') result += ind;

    result += line;
  }

  return result;
}

function generateNextLine(state, level) {
  return '\n' + common.repeat(' ', state.indent * level);
}

function testImplicitResolving(state, str) {
  var index, length, type;

  for (index = 0, length = state.implicitTypes.length; index < length; index += 1) {
    type = state.implicitTypes[index];

    if (type.resolve(str)) {
      return true;
    }
  }

  return false;
}

// [33] s-white ::= s-space | s-tab
function isWhitespace(c) {
  return c === CHAR_SPACE || c === CHAR_TAB;
}

// Returns true if the character can be printed without escaping.
// From YAML 1.2: "any allowed characters known to be non-printable
// should also be escaped. [However,] This isn’t mandatory"
// Derived from nb-char - \t - #x85 - #xA0 - #x2028 - #x2029.
function isPrintable(c) {
  return  (0x00020 <= c && c <= 0x00007E)
      || ((0x000A1 <= c && c <= 0x00D7FF) && c !== 0x2028 && c !== 0x2029)
      || ((0x0E000 <= c && c <= 0x00FFFD) && c !== 0xFEFF /* BOM */)
      ||  (0x10000 <= c && c <= 0x10FFFF);
}

// [34] ns-char ::= nb-char - s-white
// [27] nb-char ::= c-printable - b-char - c-byte-order-mark
// [26] b-char  ::= b-line-feed | b-carriage-return
// [24] b-line-feed       ::=     #xA    /* LF */
// [25] b-carriage-return ::=     #xD    /* CR */
// [3]  c-byte-order-mark ::=     #xFEFF
function isNsChar(c) {
  return isPrintable(c) && !isWhitespace(c)
    // byte-order-mark
    && c !== 0xFEFF
    // b-char
    && c !== CHAR_CARRIAGE_RETURN
    && c !== CHAR_LINE_FEED;
}

// Simplified test for values allowed after the first character in plain style.
function isPlainSafe(c, prev) {
  // Uses a subset of nb-char - c-flow-indicator - ":" - "#"
  // where nb-char ::= c-printable - b-char - c-byte-order-mark.
  return isPrintable(c) && c !== 0xFEFF
    // - c-flow-indicator
    && c !== CHAR_COMMA
    && c !== CHAR_LEFT_SQUARE_BRACKET
    && c !== CHAR_RIGHT_SQUARE_BRACKET
    && c !== CHAR_LEFT_CURLY_BRACKET
    && c !== CHAR_RIGHT_CURLY_BRACKET
    // - ":" - "#"
    // /* An ns-char preceding */ "#"
    && c !== CHAR_COLON
    && ((c !== CHAR_SHARP) || (prev && isNsChar(prev)));
}

// Simplified test for values allowed as the first character in plain style.
function isPlainSafeFirst(c) {
  // Uses a subset of ns-char - c-indicator
  // where ns-char = nb-char - s-white.
  return isPrintable(c) && c !== 0xFEFF
    && !isWhitespace(c) // - s-white
    // - (c-indicator ::=
    // “-” | “?” | “:” | “,” | “[” | “]” | “{” | “}”
    && c !== CHAR_MINUS
    && c !== CHAR_QUESTION
    && c !== CHAR_COLON
    && c !== CHAR_COMMA
    && c !== CHAR_LEFT_SQUARE_BRACKET
    && c !== CHAR_RIGHT_SQUARE_BRACKET
    && c !== CHAR_LEFT_CURLY_BRACKET
    && c !== CHAR_RIGHT_CURLY_BRACKET
    // | “#” | “&” | “*” | “!” | “|” | “=” | “>” | “'” | “"”
    && c !== CHAR_SHARP
    && c !== CHAR_AMPERSAND
    && c !== CHAR_ASTERISK
    && c !== CHAR_EXCLAMATION
    && c !== CHAR_VERTICAL_LINE
    && c !== CHAR_EQUALS
    && c !== CHAR_GREATER_THAN
    && c !== CHAR_SINGLE_QUOTE
    && c !== CHAR_DOUBLE_QUOTE
    // | “%” | “@” | “`”)
    && c !== CHAR_PERCENT
    && c !== CHAR_COMMERCIAL_AT
    && c !== CHAR_GRAVE_ACCENT;
}

// Determines whether block indentation indicator is required.
function needIndentIndicator(string) {
  var leadingSpaceRe = /^\n* /;
  return leadingSpaceRe.test(string);
}

var STYLE_PLAIN   = 1,
    STYLE_SINGLE  = 2,
    STYLE_LITERAL = 3,
    STYLE_FOLDED  = 4,
    STYLE_DOUBLE  = 5;

// Determines which scalar styles are possible and returns the preferred style.
// lineWidth = -1 => no limit.
// Pre-conditions: str.length > 0.
// Post-conditions:
//    STYLE_PLAIN or STYLE_SINGLE => no \n are in the string.
//    STYLE_LITERAL => no lines are suitable for folding (or lineWidth is -1).
//    STYLE_FOLDED => a line > lineWidth and can be folded (and lineWidth != -1).
function chooseScalarStyle(string, singleLineOnly, indentPerLevel, lineWidth, testAmbiguousType) {
  var i;
  var char, prev_char;
  var hasLineBreak = false;
  var hasFoldableLine = false; // only checked if shouldTrackWidth
  var shouldTrackWidth = lineWidth !== -1;
  var previousLineBreak = -1; // count the first line correctly
  var plain = isPlainSafeFirst(string.charCodeAt(0))
          && !isWhitespace(string.charCodeAt(string.length - 1));

  if (singleLineOnly) {
    // Case: no block styles.
    // Check for disallowed characters to rule out plain and single.
    for (i = 0; i < string.length; i++) {
      char = string.charCodeAt(i);
      if (!isPrintable(char)) {
        return STYLE_DOUBLE;
      }
      prev_char = i > 0 ? string.charCodeAt(i - 1) : null;
      plain = plain && isPlainSafe(char, prev_char);
    }
  } else {
    // Case: block styles permitted.
    for (i = 0; i < string.length; i++) {
      char = string.charCodeAt(i);
      if (char === CHAR_LINE_FEED) {
        hasLineBreak = true;
        // Check if any line can be folded.
        if (shouldTrackWidth) {
          hasFoldableLine = hasFoldableLine ||
            // Foldable line = too long, and not more-indented.
            (i - previousLineBreak - 1 > lineWidth &&
             string[previousLineBreak + 1] !== ' ');
          previousLineBreak = i;
        }
      } else if (!isPrintable(char)) {
        return STYLE_DOUBLE;
      }
      prev_char = i > 0 ? string.charCodeAt(i - 1) : null;
      plain = plain && isPlainSafe(char, prev_char);
    }
    // in case the end is missing a \n
    hasFoldableLine = hasFoldableLine || (shouldTrackWidth &&
      (i - previousLineBreak - 1 > lineWidth &&
       string[previousLineBreak + 1] !== ' '));
  }
  // Although every style can represent \n without escaping, prefer block styles
  // for multiline, since they're more readable and they don't add empty lines.
  // Also prefer folding a super-long line.
  if (!hasLineBreak && !hasFoldableLine) {
    // Strings interpretable as another type have to be quoted;
    // e.g. the string 'true' vs. the boolean true.
    return plain && !testAmbiguousType(string)
      ? STYLE_PLAIN : STYLE_SINGLE;
  }
  // Edge case: block indentation indicator can only have one digit.
  if (indentPerLevel > 9 && needIndentIndicator(string)) {
    return STYLE_DOUBLE;
  }
  // At this point we know block styles are valid.
  // Prefer literal style unless we want to fold.
  return hasFoldableLine ? STYLE_FOLDED : STYLE_LITERAL;
}

// Note: line breaking/folding is implemented for only the folded style.
// NB. We drop the last trailing newline (if any) of a returned block scalar
//  since the dumper adds its own newline. This always works:
//    • No ending newline => unaffected; already using strip "-" chomping.
//    • Ending newline    => removed then restored.
//  Importantly, this keeps the "+" chomp indicator from gaining an extra line.
function writeScalar(state, string, level, iskey) {
  state.dump = (function () {
    if (string.length === 0) {
      return "''";
    }
    if (!state.noCompatMode &&
        DEPRECATED_BOOLEANS_SYNTAX.indexOf(string) !== -1) {
      return "'" + string + "'";
    }

    var indent = state.indent * Math.max(1, level); // no 0-indent scalars
    // As indentation gets deeper, let the width decrease monotonically
    // to the lower bound min(state.lineWidth, 40).
    // Note that this implies
    //  state.lineWidth ≤ 40 + state.indent: width is fixed at the lower bound.
    //  state.lineWidth > 40 + state.indent: width decreases until the lower bound.
    // This behaves better than a constant minimum width which disallows narrower options,
    // or an indent threshold which causes the width to suddenly increase.
    var lineWidth = state.lineWidth === -1
      ? -1 : Math.max(Math.min(state.lineWidth, 40), state.lineWidth - indent);

    // Without knowing if keys are implicit/explicit, assume implicit for safety.
    var singleLineOnly = iskey
      // No block styles in flow mode.
      || (state.flowLevel > -1 && level >= state.flowLevel);
    function testAmbiguity(string) {
      return testImplicitResolving(state, string);
    }

    switch (chooseScalarStyle(string, singleLineOnly, state.indent, lineWidth, testAmbiguity)) {
      case STYLE_PLAIN:
        return string;
      case STYLE_SINGLE:
        return "'" + string.replace(/'/g, "''") + "'";
      case STYLE_LITERAL:
        return '|' + blockHeader(string, state.indent)
          + dropEndingNewline(indentString(string, indent));
      case STYLE_FOLDED:
        return '>' + blockHeader(string, state.indent)
          + dropEndingNewline(indentString(foldString(string, lineWidth), indent));
      case STYLE_DOUBLE:
        return '"' + escapeString(string) + '"';
      default:
        throw new YAMLException('impossible error: invalid scalar style');
    }
  }());
}

// Pre-conditions: string is valid for a block scalar, 1 <= indentPerLevel <= 9.
function blockHeader(string, indentPerLevel) {
  var indentIndicator = needIndentIndicator(string) ? String(indentPerLevel) : '';

  // note the special case: the string '\n' counts as a "trailing" empty line.
  var clip =          string[string.length - 1] === '\n';
  var keep = clip && (string[string.length - 2] === '\n' || string === '\n');
  var chomp = keep ? '+' : (clip ? '' : '-');

  return indentIndicator + chomp + '\n';
}

// (See the note for writeScalar.)
function dropEndingNewline(string) {
  return string[string.length - 1] === '\n' ? string.slice(0, -1) : string;
}

// Note: a long line without a suitable break point will exceed the width limit.
// Pre-conditions: every char in str isPrintable, str.length > 0, width > 0.
function foldString(string, width) {
  // In folded style, $k$ consecutive newlines output as $k+1$ newlines—
  // unless they're before or after a more-indented line, or at the very
  // beginning or end, in which case $k$ maps to $k$.
  // Therefore, parse each chunk as newline(s) followed by a content line.
  var lineRe = /(\n+)([^\n]*)/g;

  // first line (possibly an empty line)
  var result = (function () {
    var nextLF = string.indexOf('\n');
    nextLF = nextLF !== -1 ? nextLF : string.length;
    lineRe.lastIndex = nextLF;
    return foldLine(string.slice(0, nextLF), width);
  }());
  // If we haven't reached the first content line yet, don't add an extra \n.
  var prevMoreIndented = string[0] === '\n' || string[0] === ' ';
  var moreIndented;

  // rest of the lines
  var match;
  while ((match = lineRe.exec(string))) {
    var prefix = match[1], line = match[2];
    moreIndented = (line[0] === ' ');
    result += prefix
      + (!prevMoreIndented && !moreIndented && line !== ''
        ? '\n' : '')
      + foldLine(line, width);
    prevMoreIndented = moreIndented;
  }

  return result;
}

// Greedy line breaking.
// Picks the longest line under the limit each time,
// otherwise settles for the shortest line over the limit.
// NB. More-indented lines *cannot* be folded, as that would add an extra \n.
function foldLine(line, width) {
  if (line === '' || line[0] === ' ') return line;

  // Since a more-indented line adds a \n, breaks can't be followed by a space.
  var breakRe = / [^ ]/g; // note: the match index will always be <= length-2.
  var match;
  // start is an inclusive index. end, curr, and next are exclusive.
  var start = 0, end, curr = 0, next = 0;
  var result = '';

  // Invariants: 0 <= start <= length-1.
  //   0 <= curr <= next <= max(0, length-2). curr - start <= width.
  // Inside the loop:
  //   A match implies length >= 2, so curr and next are <= length-2.
  while ((match = breakRe.exec(line))) {
    next = match.index;
    // maintain invariant: curr - start <= width
    if (next - start > width) {
      end = (curr > start) ? curr : next; // derive end <= length-2
      result += '\n' + line.slice(start, end);
      // skip the space that was output as \n
      start = end + 1;                    // derive start <= length-1
    }
    curr = next;
  }

  // By the invariants, start <= length-1, so there is something left over.
  // It is either the whole string or a part starting from non-whitespace.
  result += '\n';
  // Insert a break if the remainder is too long and there is a break available.
  if (line.length - start > width && curr > start) {
    result += line.slice(start, curr) + '\n' + line.slice(curr + 1);
  } else {
    result += line.slice(start);
  }

  return result.slice(1); // drop extra \n joiner
}

// Escapes a double-quoted string.
function escapeString(string) {
  var result = '';
  var char, nextChar;
  var escapeSeq;

  for (var i = 0; i < string.length; i++) {
    char = string.charCodeAt(i);
    // Check for surrogate pairs (reference Unicode 3.0 section "3.7 Surrogates").
    if (char >= 0xD800 && char <= 0xDBFF/* high surrogate */) {
      nextChar = string.charCodeAt(i + 1);
      if (nextChar >= 0xDC00 && nextChar <= 0xDFFF/* low surrogate */) {
        // Combine the surrogate pair and store it escaped.
        result += encodeHex((char - 0xD800) * 0x400 + nextChar - 0xDC00 + 0x10000);
        // Advance index one extra since we already used that char here.
        i++; continue;
      }
    }
    escapeSeq = ESCAPE_SEQUENCES[char];
    result += !escapeSeq && isPrintable(char)
      ? string[i]
      : escapeSeq || encodeHex(char);
  }

  return result;
}

function writeFlowSequence(state, level, object) {
  var _result = '',
      _tag    = state.tag,
      index,
      length;

  for (index = 0, length = object.length; index < length; index += 1) {
    // Write only valid elements.
    if (writeNode(state, level, object[index], false, false)) {
      if (index !== 0) _result += ',' + (!state.condenseFlow ? ' ' : '');
      _result += state.dump;
    }
  }

  state.tag = _tag;
  state.dump = '[' + _result + ']';
}

function writeBlockSequence(state, level, object, compact) {
  var _result = '',
      _tag    = state.tag,
      index,
      length;

  for (index = 0, length = object.length; index < length; index += 1) {
    // Write only valid elements.
    if (writeNode(state, level + 1, object[index], true, true)) {
      if (!compact || index !== 0) {
        _result += generateNextLine(state, level);
      }

      if (state.dump && CHAR_LINE_FEED === state.dump.charCodeAt(0)) {
        _result += '-';
      } else {
        _result += '- ';
      }

      _result += state.dump;
    }
  }

  state.tag = _tag;
  state.dump = _result || '[]'; // Empty sequence if no valid values.
}

function writeFlowMapping(state, level, object) {
  var _result       = '',
      _tag          = state.tag,
      objectKeyList = Object.keys(object),
      index,
      length,
      objectKey,
      objectValue,
      pairBuffer;

  for (index = 0, length = objectKeyList.length; index < length; index += 1) {

    pairBuffer = '';
    if (index !== 0) pairBuffer += ', ';

    if (state.condenseFlow) pairBuffer += '"';

    objectKey = objectKeyList[index];
    objectValue = object[objectKey];

    if (!writeNode(state, level, objectKey, false, false)) {
      continue; // Skip this pair because of invalid key;
    }

    if (state.dump.length > 1024) pairBuffer += '? ';

    pairBuffer += state.dump + (state.condenseFlow ? '"' : '') + ':' + (state.condenseFlow ? '' : ' ');

    if (!writeNode(state, level, objectValue, false, false)) {
      continue; // Skip this pair because of invalid value.
    }

    pairBuffer += state.dump;

    // Both key and value are valid.
    _result += pairBuffer;
  }

  state.tag = _tag;
  state.dump = '{' + _result + '}';
}

function writeBlockMapping(state, level, object, compact) {
  var _result       = '',
      _tag          = state.tag,
      objectKeyList = Object.keys(object),
      index,
      length,
      objectKey,
      objectValue,
      explicitPair,
      pairBuffer;

  // Allow sorting keys so that the output file is deterministic
  if (state.sortKeys === true) {
    // Default sorting
    objectKeyList.sort();
  } else if (typeof state.sortKeys === 'function') {
    // Custom sort function
    objectKeyList.sort(state.sortKeys);
  } else if (state.sortKeys) {
    // Something is wrong
    throw new YAMLException('sortKeys must be a boolean or a function');
  }

  for (index = 0, length = objectKeyList.length; index < length; index += 1) {
    pairBuffer = '';

    if (!compact || index !== 0) {
      pairBuffer += generateNextLine(state, level);
    }

    objectKey = objectKeyList[index];
    objectValue = object[objectKey];

    if (!writeNode(state, level + 1, objectKey, true, true, true)) {
      continue; // Skip this pair because of invalid key.
    }

    explicitPair = (state.tag !== null && state.tag !== '?') ||
                   (state.dump && state.dump.length > 1024);

    if (explicitPair) {
      if (state.dump && CHAR_LINE_FEED === state.dump.charCodeAt(0)) {
        pairBuffer += '?';
      } else {
        pairBuffer += '? ';
      }
    }

    pairBuffer += state.dump;

    if (explicitPair) {
      pairBuffer += generateNextLine(state, level);
    }

    if (!writeNode(state, level + 1, objectValue, true, explicitPair)) {
      continue; // Skip this pair because of invalid value.
    }

    if (state.dump && CHAR_LINE_FEED === state.dump.charCodeAt(0)) {
      pairBuffer += ':';
    } else {
      pairBuffer += ': ';
    }

    pairBuffer += state.dump;

    // Both key and value are valid.
    _result += pairBuffer;
  }

  state.tag = _tag;
  state.dump = _result || '{}'; // Empty mapping if no valid pairs.
}

function detectType(state, object, explicit) {
  var _result, typeList, index, length, type, style;

  typeList = explicit ? state.explicitTypes : state.implicitTypes;

  for (index = 0, length = typeList.length; index < length; index += 1) {
    type = typeList[index];

    if ((type.instanceOf  || type.predicate) &&
        (!type.instanceOf || ((typeof object === 'object') && (object instanceof type.instanceOf))) &&
        (!type.predicate  || type.predicate(object))) {

      state.tag = explicit ? type.tag : '?';

      if (type.represent) {
        style = state.styleMap[type.tag] || type.defaultStyle;

        if (_toString.call(type.represent) === '[object Function]') {
          _result = type.represent(object, style);
        } else if (_hasOwnProperty.call(type.represent, style)) {
          _result = type.represent[style](object, style);
        } else {
          throw new YAMLException('!<' + type.tag + '> tag resolver accepts not "' + style + '" style');
        }

        state.dump = _result;
      }

      return true;
    }
  }

  return false;
}

// Serializes `object` and writes it to global `result`.
// Returns true on success, or false on invalid object.
//
function writeNode(state, level, object, block, compact, iskey) {
  state.tag = null;
  state.dump = object;

  if (!detectType(state, object, false)) {
    detectType(state, object, true);
  }

  var type = _toString.call(state.dump);

  if (block) {
    block = (state.flowLevel < 0 || state.flowLevel > level);
  }

  var objectOrArray = type === '[object Object]' || type === '[object Array]',
      duplicateIndex,
      duplicate;

  if (objectOrArray) {
    duplicateIndex = state.duplicates.indexOf(object);
    duplicate = duplicateIndex !== -1;
  }

  if ((state.tag !== null && state.tag !== '?') || duplicate || (state.indent !== 2 && level > 0)) {
    compact = false;
  }

  if (duplicate && state.usedDuplicates[duplicateIndex]) {
    state.dump = '*ref_' + duplicateIndex;
  } else {
    if (objectOrArray && duplicate && !state.usedDuplicates[duplicateIndex]) {
      state.usedDuplicates[duplicateIndex] = true;
    }
    if (type === '[object Object]') {
      if (block && (Object.keys(state.dump).length !== 0)) {
        writeBlockMapping(state, level, state.dump, compact);
        if (duplicate) {
          state.dump = '&ref_' + duplicateIndex + state.dump;
        }
      } else {
        writeFlowMapping(state, level, state.dump);
        if (duplicate) {
          state.dump = '&ref_' + duplicateIndex + ' ' + state.dump;
        }
      }
    } else if (type === '[object Array]') {
      var arrayLevel = (state.noArrayIndent && (level > 0)) ? level - 1 : level;
      if (block && (state.dump.length !== 0)) {
        writeBlockSequence(state, arrayLevel, state.dump, compact);
        if (duplicate) {
          state.dump = '&ref_' + duplicateIndex + state.dump;
        }
      } else {
        writeFlowSequence(state, arrayLevel, state.dump);
        if (duplicate) {
          state.dump = '&ref_' + duplicateIndex + ' ' + state.dump;
        }
      }
    } else if (type === '[object String]') {
      if (state.tag !== '?') {
        writeScalar(state, state.dump, level, iskey);
      }
    } else {
      if (state.skipInvalid) return false;
      throw new YAMLException('unacceptable kind of an object to dump ' + type);
    }

    if (state.tag !== null && state.tag !== '?') {
      state.dump = '!<' + state.tag + '> ' + state.dump;
    }
  }

  return true;
}

function getDuplicateReferences(object, state) {
  var objects = [],
      duplicatesIndexes = [],
      index,
      length;

  inspectNode(object, objects, duplicatesIndexes);

  for (index = 0, length = duplicatesIndexes.length; index < length; index += 1) {
    state.duplicates.push(objects[duplicatesIndexes[index]]);
  }
  state.usedDuplicates = new Array(length);
}

function inspectNode(object, objects, duplicatesIndexes) {
  var objectKeyList,
      index,
      length;

  if (object !== null && typeof object === 'object') {
    index = objects.indexOf(object);
    if (index !== -1) {
      if (duplicatesIndexes.indexOf(index) === -1) {
        duplicatesIndexes.push(index);
      }
    } else {
      objects.push(object);

      if (Array.isArray(object)) {
        for (index = 0, length = object.length; index < length; index += 1) {
          inspectNode(object[index], objects, duplicatesIndexes);
        }
      } else {
        objectKeyList = Object.keys(object);

        for (index = 0, length = objectKeyList.length; index < length; index += 1) {
          inspectNode(object[objectKeyList[index]], objects, duplicatesIndexes);
        }
      }
    }
  }
}

function dump(input, options) {
  options = options || {};

  var state = new State(options);

  if (!state.noRefs) getDuplicateReferences(input, state);

  if (writeNode(state, 0, input, true, true)) return state.dump + '\n';

  return '';
}

function safeDump(input, options) {
  return dump(input, common.extend({ schema: DEFAULT_SAFE_SCHEMA }, options));
}

dumper$1.dump     = dump;
dumper$1.safeDump = safeDump;

var loader = loader$1;
var dumper = dumper$1;


function deprecated(name) {
  return function () {
    throw new Error('Function ' + name + ' is deprecated and cannot be used.');
  };
}


jsYaml$1.Type                = type;
jsYaml$1.Schema              = schema;
jsYaml$1.FAILSAFE_SCHEMA     = failsafe;
jsYaml$1.JSON_SCHEMA         = json;
jsYaml$1.CORE_SCHEMA         = core;
jsYaml$1.DEFAULT_SAFE_SCHEMA = default_safe;
jsYaml$1.DEFAULT_FULL_SCHEMA = default_full;
jsYaml$1.load                = loader.load;
jsYaml$1.loadAll             = loader.loadAll;
jsYaml$1.safeLoad            = loader.safeLoad;
jsYaml$1.safeLoadAll         = loader.safeLoadAll;
jsYaml$1.dump                = dumper.dump;
jsYaml$1.safeDump            = dumper.safeDump;
jsYaml$1.YAMLException       = exception;

// Deprecated schema names from JS-YAML 2.0.x
jsYaml$1.MINIMAL_SCHEMA = failsafe;
jsYaml$1.SAFE_SCHEMA    = default_safe;
jsYaml$1.DEFAULT_SCHEMA = default_full;

// Deprecated functions from JS-YAML 1.x.x
jsYaml$1.scan           = deprecated('scan');
jsYaml$1.parse          = deprecated('parse');
jsYaml$1.compose        = deprecated('compose');
jsYaml$1.addConstructor = deprecated('addConstructor');

var yaml = jsYaml$1;


var jsYaml = yaml;

var yaml$1 = /*@__PURE__*/getDefaultExportFromCjs(jsYaml);

/*!
 * escape-html
 * Copyright(c) 2012-2013 TJ Holowaychuk
 * Copyright(c) 2015 Andreas Lubbe
 * Copyright(c) 2015 Tiancheng "Timothy" Gu
 * MIT Licensed
 */

/**
 * Module variables.
 * @private
 */

var matchHtmlRegExp = /["'&<>]/;

/**
 * Module exports.
 * @public
 */

var escapeHtml_1 = escapeHtml;

/**
 * Escape special characters in the given string of html.
 *
 * @param  {string} string The string to escape for inserting into HTML
 * @return {string}
 * @public
 */

function escapeHtml(string) {
  var str = '' + string;
  var match = matchHtmlRegExp.exec(str);

  if (!match) {
    return str;
  }

  var escape;
  var html = '';
  var index = 0;
  var lastIndex = 0;

  for (index = match.index; index < str.length; index++) {
    switch (str.charCodeAt(index)) {
      case 34: // "
        escape = '&quot;';
        break;
      case 38: // &
        escape = '&amp;';
        break;
      case 39: // '
        escape = '&#39;';
        break;
      case 60: // <
        escape = '&lt;';
        break;
      case 62: // >
        escape = '&gt;';
        break;
      default:
        continue;
    }

    if (lastIndex !== index) {
      html += str.substring(lastIndex, index);
    }

    lastIndex = index + 1;
    html += escape;
  }

  return lastIndex !== index
    ? html + str.substring(lastIndex, index)
    : html;
}

var escape = /*@__PURE__*/getDefaultExportFromCjs(escapeHtml_1);

function _optionalChain(ops) { let lastAccessLHS = undefined; let value = ops[0]; let i = 1; while (i < ops.length) { const op = ops[i]; const fn = ops[i + 1]; i += 2; if ((op === 'optionalAccess' || op === 'optionalCall') && value == null) { return undefined; } if (op === 'access' || op === 'optionalAccess') { lastAccessLHS = value; value = fn(value); } else if (op === 'call' || op === 'optionalCall') { value = fn((...args) => value.call(lastAccessLHS, ...args)); lastAccessLHS = undefined; } } return value; }















// this needs a big old cleanup

const newline = '\n';
// extract the yaml from 'yaml' nodes and put them in the vfil for later use

function default_frontmatter(
	value,
	messages
) {
	try {
		return yaml$1.safeLoad(value) ;
	} catch (e) {
		messages.push(new Message('YAML failed to parse'));
	}
}

function parse_frontmatter({
	parse,
	type,
}) {
	const transformer = (tree, vFile) => {
		visit$6(tree, type, (node) => {
			const data = parse(node.value, vFile.messages);
			if (data) {
				// @ts-ignore
				vFile.data.fm = data;
			}
		});
	};

	return transformer;
}

// in code nodes replace the character witrh the html entities
// maybe I'll need more of these

const entites = [
	[/</g, '&lt;'],
	[/>/g, '&gt;'],
	[/{/g, '&#123;'],
	[/}/g, '&#125;'],
];

function escape_code({ blocks }) {
	return function (tree) {
		if (!blocks) {
			visit$6(tree, 'code', escape);
		}

		visit$6(tree, 'inlineCode', escape);

		function escape(node) {
			for (let i = 0; i < entites.length; i += 1) {
				node.value = node.value.replace(entites[i][0], entites[i][1]);
			}
		}
	};
}

// special case - process nodes with retext and smartypants
// retext plugins can't work generally due to the difficulties in converting between the two trees

function smartypants_transformer(options = {}) {
	const processor = retext$1().use(smartypants$1, options);

	return function (tree) {
		visit$6(tree, 'text', (node) => {
			node.value = String(processor.processSync(node.value));
		});
	};
}

// regex for scripts and attributes

const attrs = `(?:\\s{0,1}[a-zA-z]+=(?:"){0,1}[a-zA-Z0-9]+(?:"){0,1})*`;
const context = `(?:\\s{0,1}context)=(?:"){0,1}module(?:"){0,1}`;

const RE_BLANK = /^\n+$|^\s+$/;

const RE_SCRIPT = new RegExp(`^(<script` + attrs + `>)`);

const RE_MODULE_SCRIPT = new RegExp(
	`^(<script` + attrs + context + attrs + `>)`
);

function extract_parts(nodes) {
	// since we are wrapping and replacing we need to keep track of the different component 'parts'
	// many special tags cannot be wrapped nor can style or script tags
	const parts = {
		special: [],
		html: [],
		instance: [],
		module: [],
		css: [],
	};

	// iterate through all top level child nodes and assign them to the correct 'part'
	// anything that is a normal HAST node gets stored as HTML untouched
	// everything else gets parsed by the svelte parser

	children: for (let i = 0; i < nodes.length; i += 1) {
		const empty_node =
			nodes[i].type === 'text' && RE_BLANK.exec(nodes[i].value );

		// i no longer knwo why i did this

		if (empty_node || !nodes[i].value) {
			if (
				!parts.html.length ||
				!(
					RE_BLANK.exec(nodes[i].value ) &&
					RE_BLANK.exec(parts.html[parts.html.length - 1].value )
				)
			) {
				parts.html.push(nodes[i]);
			}

			continue children;
		}

		let result










;
		try {
			result = parse$7(nodes[i].value );
		} catch (e) {
			parts.html.push(nodes[i]);
			continue children;
		}

		// svelte special tags that have to be top level
		if (!result.html || !result.html.children) return parts;

		const _parts

 = result.html.children.map((v) => {
			if (
				v.type === 'Options' ||
				v.type === 'Head' ||
				v.type === 'Window' ||
				v.type === 'Body'
			) {
				return ['special', v.start, v.end];
			} else {
				return ['html', v.start, v.end];
			}
		});

		results: for (const key in result) {
			if (key === 'html' || !result[key ])
				continue results;
			_parts.push([
				key ,
				result[key ].start,
				result[key ].end,
			]);
		}

		// sort them to ensure the array is in the order they appear in the source, no gaps
		// this might not be necessary any more, i forget
		const sorted = _parts.sort((a, b) => a[1] - b[1]);

		// push the nodes into the correct 'part' since they are sorted everything should be in the correct order
		sorted.forEach((next) => {
			parts[next[0]].push({
				type: 'raw',
				value: (nodes[i].value ).substring(next[1], next[2]),
			});
		});
	}

	return parts;
}

function map_layout_to_path(
	filename,
	layout_map
) {
	const match = Object.keys(layout_map).find((l) =>
		new RegExp(`\\` + `${path$2.sep}${l}` + `\\` + `${path$2.sep}`).test(
			path$2.normalize(filename).replace(process.cwd(), '')
		)
	);

	if (match) {
		return layout_map[match];
	} else {
		return layout_map['_'] ? layout_map['_'] : undefined;
	}
}

function generate_layout_import(
	layout
) {
	if (!layout) return false;

	return `import Layout_MDSVEX_DEFAULT${
		layout.components.length ? `, * as Components` : ''
	} from '${layout.path}';`;
}

function generate_layout({
	frontmatter_layout,
	layout_options,
	layout_mode,
	filename,
}




) {
	let selected_layout;
	const error = { reason: '' };

	if (!layout_options || frontmatter_layout === false) {
		return [false, false, false];
	} else if (layout_mode === 'single') {
		selected_layout = layout_options.__mdsvex_default;
		if (frontmatter_layout)
			error.reason = `You attempted to apply a named layout in the front-matter of "${filename}", but did not provide any named layouts as options to the preprocessor. `;
	} else if (frontmatter_layout) {
		selected_layout = layout_options[frontmatter_layout];
		if (!selected_layout)
			error.reason = `Could not find a layout with the name "${frontmatter_layout}" and no fall back layout ("_") was provided.`;
	} else {
		selected_layout = map_layout_to_path(filename, layout_options);
	}

	return [
		generate_layout_import(selected_layout),
		selected_layout !== undefined &&
			selected_layout.components.length > 0 &&
			selected_layout.components,
		error.reason ? error : false,
	];
}

function transform_hast({
	layout,
	layout_mode,
}


) {
	return function transformer(tree, vFile) {
		// we need to keep { and } intact for svelte, so reverse the escaping in links and images
		// if anyone actually uses these characters for any other reason i'll probably just cry
		visit$6(tree, 'element', (node) => {
			if (
				node.tagName === 'a' &&
				node.properties &&
				typeof node.properties.href === 'string'
			) {
				node.properties.href = node.properties.href
					.replace(/%7B/g, '{')
					.replace(/%7D/g, '}');
			}

			if (
				node.tagName === 'img' &&
				node.properties &&
				typeof node.properties.src === 'string'
			) {
				node.properties.src = node.properties.src
					.replace(/%7B/g, '{')
					.replace(/%7D/g, '}');
			}
		});

		// the rest only applies to layouts and front matter
		// this  breaks position data
		// svelte preprocessors don't currently support sourcemaps
		// i'll fix this when they do

		//@ts-ignore
		if (!layout && !vFile.data.fm) return tree;

		visit$6(tree, 'root', (node) => {
			const { special, html, instance, module: _module, css } = extract_parts(
				node.children 
			);

			const { fm: metadata } = vFile.data ;

			// Workaround for script and style tags in strings
			// https://github.com/sveltejs/svelte/issues/5292
			const stringified =
				metadata &&
				JSON.stringify(metadata).replace(/<(\/?script|\/?style)/g, '<"+"$1');

			const fm =
				metadata &&
				`export const metadata = ${stringified};${newline}` +
					`\tconst { ${Object.keys(metadata).join(', ')} } = metadata;`;

			const frontmatter_layout =
				metadata && (metadata.layout );

			const [import_script, components, error] = generate_layout({
				frontmatter_layout,
				layout_options: layout,
				layout_mode,
				//@ts-ignore
				filename: vFile.filename,
			});

			if (error) vFile.messages.push(new Message(error.reason));

			if (components) {
				for (let i = 0; i < components.length; i++) {
					visit$6(tree, 'element', (node) => {
						if (node.tagName === components[i]) {
							node.tagName = `Components.${components[i]}`;
						}
					});
				}
			}

			// add the layout if we are using one, reusing the existing script if one exists
			if (import_script && !instance[0]) {
				instance.push({
					type: 'raw',
					value: `${newline}<script>${newline}\t${import_script}${newline}</script>${newline}`,
				});
			} else if (import_script) {
				instance[0].value = (instance[0].value ).replace(
					RE_SCRIPT,
					`$1${newline}\t${import_script}`
				);
			}

			// inject the frontmatter into the module script if there is any, reusing the existing module script if one exists
			if (!_module[0] && fm) {
				_module.push({
					type: 'raw',
					value: `<script context="module">${newline}\t${fm}${newline}</script>`,
				});
			} else if (fm) {
				// @ts-ignore
				_module[0].value = _module[0].value.replace(
					RE_MODULE_SCRIPT,
					(match) => `${match}${newline}\t${fm}`
				);
			}

			// smoosh it all together in an order that makes sense,
			// if using a layout we only wrap the html and nothing else
			//@ts-ignore
			node.children = [
				//@ts-ignore
				..._module,
				//@ts-ignore
				{ type: 'raw', value: _module[0] ? newline : '' },
				//@ts-ignore
				...instance,
				//@ts-ignore
				{ type: 'raw', value: instance[0] ? newline : '' },
				//@ts-ignore
				...css,
				//@ts-ignore
				{ type: 'raw', value: css[0] ? newline : '' },
				//@ts-ignore
				...special,
				//@ts-ignore
				{ type: 'raw', value: special[0] ? newline : '' },

				{
					//@ts-ignore
					type: 'raw',
					value: import_script
						? `<Layout_MDSVEX_DEFAULT {...$$props}${
								fm ? ' {...metadata}' : ''
						  }>`
						: '',
				},
				//@ts-ignore
				{ type: 'raw', value: newline },
				//@ts-ignore
				...html,
				//@ts-ignore
				{ type: 'raw', value: newline },
				//@ts-ignore
				{ type: 'raw', value: import_script ? '</Layout_MDSVEX_DEFAULT>' : '' },
			];
		});
	};
}

// highlighting stuff

// { [lang]: { path, deps: pointer to key } }
const langs = {};
let Prism;

const make_path = (base_path, id) =>
	base_path.replace('{id}', id);

// we need to get all language metadata
// also track if they depend on other languages so we can autoload without breaking
// i don't actually know what the require key means but it sounds important

function get_lang_info(
	name,
	lang_meta,
	base_path
) {
	const _lang_meta = {
		name,
		path: `prismjs/${make_path(base_path, name)}`,
		deps: new Set(),
	};

	const aliases = new Set();

	// TODO: DRY this up, it is literally identical

	if (lang_meta.require) {
		if (Array.isArray(lang_meta.require)) {
			lang_meta.require.forEach((id) => _lang_meta.deps.add(id));
		} else {
			_lang_meta.deps.add(lang_meta.require);
		}
	}

	if (lang_meta.peerDependencies) {
		if (Array.isArray(lang_meta.peerDependencies)) {
			lang_meta.peerDependencies.forEach((id) => _lang_meta.deps.add(id));
		} else {
			_lang_meta.deps.add(lang_meta.peerDependencies);
		}
	}

	if (lang_meta.alias) {
		if (Array.isArray(lang_meta.alias)) {
			lang_meta.alias.forEach((id) => aliases.add(id));
		} else {
			aliases.add(lang_meta.alias);
		}
	}

	return [{ ..._lang_meta, aliases }, aliases];
}

// workaround for ts weirdness - intersection types work better with interfaces vs object literals




function load_language_metadata() {
	if (!(process ).browser) {
		const {
			meta,
			...languages
		}
 = require('prismjs/components.json').languages;

		for (const lang in languages) {
			const [lang_info, aliases] = get_lang_info(
				lang,
				languages[lang],
				meta.path
			);

			langs[lang] = lang_info;
			aliases.forEach((_n) => {
				langs[_n] = langs[lang];
			});
		}

		const svelte_meta = {
			name: 'svelte',
			aliases: new Set(['sv']),
			path: 'prism-svelte',
			deps: new Set(['javscript', 'css']),
		};

		langs.svelte = svelte_meta;
		langs.sv = svelte_meta;
	}
}

function load_language(lang) {
	if (!(process ).browser) {
		if (!langs[lang]) return;

		langs[lang].deps.forEach((name) => load_language(name));

		require(langs[lang].path);
	}
}

function highlight_blocks({
	highlighter: highlight_fn,
	alias,
}


 = {}) {
	if (highlight_fn && !(process ).browser) {
		load_language_metadata();

		if (alias) {
			for (const lang in alias) {
				langs[lang] = langs[alias[lang]];
			}
		}
	}

	return async function (tree, vFile) {
		if (highlight_fn) {
			const nodes = [];
			visit$6(tree, 'code', (node) => {
				nodes.push(node);
			});

			await Promise.all(
				nodes.map(async (node) => {
					(node ).type = 'html';
					node.value = await highlight_fn(
						node.value,
						(node ).lang,
						(node ).meta,
						//@ts-ignore
						vFile.filename
					);
				})
			);
		}
	};
}
// escape curlies, backtick, \t, \r, \n to avoid breaking output of {@html `here`} in .svelte
const escape_svelty = (str) =>
	str
		.replace(
			/[{}`]/g,
			//@ts-ignore
			(c) => ({ '{': '&#123;', '}': '&#125;', '`': '&#96;' }[c])
		)
		.replace(/\\([trn])/g, '&#92;$1');

const code_highlight = (code, lang) => {
	const normalised_lang = _optionalChain([lang, 'optionalAccess', _ => _.toLowerCase, 'call', _2 => _2()]);
	if (!(process ).browser) {
		let _lang = !!normalised_lang && langs[normalised_lang];

		if (!Prism) Prism = require('prismjs');

		if (_lang && !Prism.languages[_lang.name]) {
			load_language(_lang.name);
		}

		if (!_lang && normalised_lang && Prism.languages[normalised_lang]) {
			langs[normalised_lang] = { name: lang } ;
			_lang = langs[normalised_lang];
		}
		const highlighted = escape_svelty(
			_lang
				? Prism.highlight(code, Prism.languages[_lang.name], _lang.name)
				: escape(code)
		);
		return `<pre class="language-${normalised_lang}">{@html \`<code class="language-${normalised_lang}">${highlighted}</code>\`}</pre>`;
	} else {
		const highlighted = escape_svelty(escape(code));
		return `<pre class="language-${normalised_lang}">{@html \`<code class="language-${normalised_lang}">${highlighted}</code>\`}</pre>`;
	}
};

function stringify( options = {}) {
	this.Compiler = compiler;

	function compiler(tree) {
		return hast_to_html(tree, options);
	}
}

const apply_plugins = (plugins, parser) => {
	(plugins ).forEach((plugin) => {
		if (Array.isArray(plugin)) {
			if (plugin[1] && plugin[1]) parser.use(plugin[0], plugin[1]);
			else parser.use(plugin[0]);
		} else {
			parser.use(plugin);
		}
	});

	return parser;
};

function transform(
	{
		remarkPlugins = [],
		rehypePlugins = [],
		frontmatter,
		smartypants,
		layout,
		layout_mode,
		highlight,
	} = { layout_mode: 'single' }
) {
	const fm_opts = frontmatter
		? frontmatter
		: { parse: default_frontmatter, type: 'yaml', marker: '-' };
	const toMDAST = unified$3()
		.use(markdown)
		.use(mdsvex_parser)
		.use(external, { target: false, rel: ['nofollow'] })
		.use(escape_code, { blocks: !!highlight })
		.use(extract_frontmatter, [{ type: fm_opts.type, marker: fm_opts.marker }])
		.use(parse_frontmatter, { parse: fm_opts.parse, type: fm_opts.type });

	if (smartypants) {
		toMDAST.use(
			smartypants_transformer,
			typeof smartypants === 'boolean' ? {} : smartypants
		);
	}

	apply_plugins(remarkPlugins, toMDAST).use(highlight_blocks, highlight || {});

	const toHAST = toMDAST
		.use(remark2rehype$1, {
			// @ts-ignore
			allowDangerousHtml: true,
			allowDangerousCharacters: true,
		})
		.use(transform_hast, { layout, layout_mode });

	apply_plugins(rehypePlugins, toHAST);

	const processor = toHAST.use(stringify, {
		allowDangerousHtml: true,
		allowDangerousCharacters: true,
	});

	return processor;
}

const defaults = {
	remarkPlugins: [],
	rehypePlugins: [],
	smartypants: true,
	extension: '.svx',
	highlight: { highlighter: code_highlight },
};

function to_posix(_path) {
	const isExtendedLengthPath = /^\\\\\?\\/.test(_path);
	const hasNonAscii = /[^\u0000-\u0080]+/.test(_path);

	if (isExtendedLengthPath || hasNonAscii) {
		return _path;
	}

	return _path.replace(/\\/g, '/');
}

function resolve_layout(layout_path) {
	try {
		return to_posix(require.resolve(layout_path));
	} catch (e) {
		try {
			const _path = join(process.cwd(), layout_path);
			return to_posix(require.resolve(_path));
		} catch (e) {
			throw new Error(
				`The layout path you provided couldn't be found at either ${layout_path} or ${join(
					process.cwd(),
					layout_path
				)}. Please double-check it and try again.`
			);
		}
	}
}

// handle custom components

function process_layouts(layouts) {
	const _layouts = layouts;

	for (const key in _layouts) {
		const layout = fs.readFileSync(_layouts[key].path, { encoding: 'utf8' });
		let ast;
		try {
			ast = parse$7(layout);
		} catch (e) {
			throw new Error(e.toString() + `\n	at ${_layouts[key].path}`);
		}

		if (ast.module) {
			const component_exports = ast.module.content.body.filter(
				(node) => node.type === 'ExportNamedDeclaration'
			) ;

			if (component_exports.length) {
				_layouts[key].components = [];

				for (let i = 0; i < component_exports.length; i++) {
					if (
						component_exports[i].specifiers &&
						component_exports[i].specifiers.length
					) {
						for (let j = 0; j < component_exports[i].specifiers.length; j++) {
							_layouts[key].components.push(
								component_exports[i].specifiers[j].exported.name
							);
						}
						//@ts-ignore
					} else if (component_exports[i].declaration.declarations) {
						//@ts-ignore
						const declarations = component_exports[i].declaration.declarations;

						for (let j = 0; j < declarations.length; j++) {
							_layouts[key].components.push(declarations[j].id.name);
						}
					} else if (component_exports[i].declaration) {
						_layouts[key].components.push(
							//@ts-ignore
							component_exports[i].declaration.id.name
						);
					}
				}
			}
		}
	}
	return _layouts;
}

/**
 * The svelte preprocessor for use with svelte.preprocess
 *
 * **options** - An options object with the following properties, all are optional.
 *
 * - `extension` - The extension to use for mdsvex files
 * - `extensions` - The extensions to use for mdsvex files
 * - `layout` - Layouts to apply to mdsvex documents
 * - `frontmatter` - frontmatter options for documents
 * - `highlight` - syntax highlighting options
 * - `smartypants` - smart typography options
 * - `remarkPlugins` - remark plugins to apply to the markdown
 * - `rehypePlugins` - rehype plugins to apply to the rendered html
 *
 */

const mdsvex = (options = defaults) => {
	const {
		remarkPlugins = [],
		rehypePlugins = [],
		smartypants = true,
		extension = '.svx',
		extensions,
		layout = false,
		highlight = { highlighter: code_highlight },
		frontmatter,
	} = options;

	//@ts-ignore
	if (options.layouts) {
		throw new Error(
			`mdsvex: "layouts" is not a valid option. Did you mean "layout"?`
		);
	}

	const unknown_opts = [];
	const known_opts = [
		'filename',
		'remarkPlugins',
		'rehypePlugins',
		'smartypants',
		'extension',
		'extensions',
		'layout',
		'highlight',
		'frontmatter',
	];

	for (const opt in options) {
		if (!known_opts.includes(opt)) unknown_opts.push(opt);
	}

	if (unknown_opts.length) {
		console.warn(
			`mdsvex: Received unknown options: ${unknown_opts.join(
				', '
			)}. Valid options are: ${known_opts.join(', ')}.`
		);
	}

	let _layout = {};
	let layout_mode = 'single';

	if (typeof layout === 'string') {
		_layout.__mdsvex_default = { path: resolve_layout(layout), components: [] };
	} else if (typeof layout === 'object') {
		layout_mode = 'named';
		for (const name in layout) {
			_layout[name] = { path: resolve_layout(layout[name]), components: [] };
		}
	}
	if (highlight && highlight.highlighter === undefined) {
		highlight.highlighter = code_highlight;
	}

	_layout = process_layouts(_layout);
	const parser = transform({
		remarkPlugins,
		rehypePlugins,
		smartypants,
		layout: _layout,
		layout_mode,
		highlight,
		frontmatter,
	});

	return {
		name: 'mdsvex',
		markup: async ({ content, filename }) => {
			const extensionsParts = (extensions || [extension]).map(ext => ext.startsWith('.') ? ext : '.' + ext);
			if (!extensionsParts.some(ext => filename.endsWith(ext))) return;

			const parsed = await parser.process({ contents: content, filename });
			return {
				code: parsed.contents ,
				data: parsed.data ,
				map: '',
			};
		},
	};
};

/**
 * The standalone compile function.
 *
 * - **source** - the source code to convert.
 * - **options** - An options object with the following properties, all are optional.
 *
 * - `filename` - The filename of the generated file
 * - `extension` - The extension to use for mdsvex files
 * - `extensions` - The extensions to use for mdsvex files
 * - `layout` - Layouts to apply to mdsvex documents
 * - `frontmatter` - frontmatter options for documents
 * - `highlight` - syntax highlighting options
 * - `smartypants` - smart typography options
 * - `remarkPlugins` - remark plugins to apply to the markdown
 * - `rehypePlugins` - rehype plugins to apply to the rendered html
 */

const _compile = (
	source,
	opts
) =>
	mdsvex(opts).markup({
		content: source,
		filename:
			(opts && opts.filename) ||
			`file${
				(opts && ((opts.extensions && opts.extensions[0]) || opts.extension)) ||
				'.svx'
			}`,
	});

export { code_highlight as code_highlighter, _compile as compile, defineConfig as defineMDSveXConfig, escape_svelty as escapeSvelte, mdsvex };
