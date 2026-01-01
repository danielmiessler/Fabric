module.exports =
/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// identity function for calling harmony imports with the correct context
/******/ 	__webpack_require__.i = function(value) { return value; };
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 14);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports) {

module.exports = require("path");

/***/ }),
/* 1 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


exports.__esModule = true;

var _promise = __webpack_require__(173);

var _promise2 = _interopRequireDefault(_promise);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

exports.default = function (fn) {
  return function () {
    var gen = fn.apply(this, arguments);
    return new _promise2.default(function (resolve, reject) {
      function step(key, arg) {
        try {
          var info = gen[key](arg);
          var value = info.value;
        } catch (error) {
          reject(error);
          return;
        }

        if (info.done) {
          resolve(value);
        } else {
          return _promise2.default.resolve(value).then(function (value) {
            step("next", value);
          }, function (err) {
            step("throw", err);
          });
        }
      }

      return step("next");
    });
  };
};

/***/ }),
/* 2 */
/***/ (function(module, exports) {

module.exports = require("util");

/***/ }),
/* 3 */
/***/ (function(module, exports) {

module.exports = require("fs");

/***/ }),
/* 4 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
class MessageError extends Error {
  constructor(msg, code) {
    super(msg);
    this.code = code;
  }

}

exports.MessageError = MessageError;
class ProcessSpawnError extends MessageError {
  constructor(msg, code, process) {
    super(msg, code);
    this.process = process;
  }

}

exports.ProcessSpawnError = ProcessSpawnError;
class SecurityError extends MessageError {}

exports.SecurityError = SecurityError;
class ProcessTermError extends MessageError {}

exports.ProcessTermError = ProcessTermError;
class ResponseError extends Error {
  constructor(msg, responseCode) {
    super(msg);
    this.responseCode = responseCode;
  }

}
exports.ResponseError = ResponseError;

/***/ }),
/* 5 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getFirstSuitableFolder = exports.readFirstAvailableStream = exports.makeTempDir = exports.hardlinksWork = exports.writeFilePreservingEol = exports.getFileSizeOnDisk = exports.walk = exports.symlink = exports.find = exports.readJsonAndFile = exports.readJson = exports.readFileAny = exports.hardlinkBulk = exports.copyBulk = exports.unlink = exports.glob = exports.link = exports.chmod = exports.lstat = exports.exists = exports.mkdirp = exports.stat = exports.access = exports.rename = exports.readdir = exports.realpath = exports.readlink = exports.writeFile = exports.open = exports.readFileBuffer = exports.lockQueue = exports.constants = undefined;

var _asyncToGenerator2;

function _load_asyncToGenerator() {
  return _asyncToGenerator2 = _interopRequireDefault(__webpack_require__(1));
}

let buildActionsForCopy = (() => {
  var _ref = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (queue, events, possibleExtraneous, reporter) {

    //
    let build = (() => {
      var _ref5 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (data) {
        const src = data.src,
              dest = data.dest,
              type = data.type;

        const onFresh = data.onFresh || noop;
        const onDone = data.onDone || noop;

        // TODO https://github.com/yarnpkg/yarn/issues/3751
        // related to bundled dependencies handling
        if (files.has(dest.toLowerCase())) {
          reporter.verbose(`The case-insensitive file ${dest} shouldn't be copied twice in one bulk copy`);
        } else {
          files.add(dest.toLowerCase());
        }

        if (type === 'symlink') {
          yield mkdirp((_path || _load_path()).default.dirname(dest));
          onFresh();
          actions.symlink.push({
            dest,
            linkname: src
          });
          onDone();
          return;
        }

        if (events.ignoreBasenames.indexOf((_path || _load_path()).default.basename(src)) >= 0) {
          // ignored file
          return;
        }

        const srcStat = yield lstat(src);
        let srcFiles;

        if (srcStat.isDirectory()) {
          srcFiles = yield readdir(src);
        }

        let destStat;
        try {
          // try accessing the destination
          destStat = yield lstat(dest);
        } catch (e) {
          // proceed if destination doesn't exist, otherwise error
          if (e.code !== 'ENOENT') {
            throw e;
          }
        }

        // if destination exists
        if (destStat) {
          const bothSymlinks = srcStat.isSymbolicLink() && destStat.isSymbolicLink();
          const bothFolders = srcStat.isDirectory() && destStat.isDirectory();
          const bothFiles = srcStat.isFile() && destStat.isFile();

          // EINVAL access errors sometimes happen which shouldn't because node shouldn't be giving
          // us modes that aren't valid. investigate this, it's generally safe to proceed.

          /* if (srcStat.mode !== destStat.mode) {
            try {
              await access(dest, srcStat.mode);
            } catch (err) {}
          } */

          if (bothFiles && artifactFiles.has(dest)) {
            // this file gets changed during build, likely by a custom install script. Don't bother checking it.
            onDone();
            reporter.verbose(reporter.lang('verboseFileSkipArtifact', src));
            return;
          }

          if (bothFiles && srcStat.size === destStat.size && (0, (_fsNormalized || _load_fsNormalized()).fileDatesEqual)(srcStat.mtime, destStat.mtime)) {
            // we can safely assume this is the same file
            onDone();
            reporter.verbose(reporter.lang('verboseFileSkip', src, dest, srcStat.size, +srcStat.mtime));
            return;
          }

          if (bothSymlinks) {
            const srcReallink = yield readlink(src);
            if (srcReallink === (yield readlink(dest))) {
              // if both symlinks are the same then we can continue on
              onDone();
              reporter.verbose(reporter.lang('verboseFileSkipSymlink', src, dest, srcReallink));
              return;
            }
          }

          if (bothFolders) {
            // mark files that aren't in this folder as possibly extraneous
            const destFiles = yield readdir(dest);
            invariant(srcFiles, 'src files not initialised');

            for (var _iterator4 = destFiles, _isArray4 = Array.isArray(_iterator4), _i4 = 0, _iterator4 = _isArray4 ? _iterator4 : _iterator4[Symbol.iterator]();;) {
              var _ref6;

              if (_isArray4) {
                if (_i4 >= _iterator4.length) break;
                _ref6 = _iterator4[_i4++];
              } else {
                _i4 = _iterator4.next();
                if (_i4.done) break;
                _ref6 = _i4.value;
              }

              const file = _ref6;

              if (srcFiles.indexOf(file) < 0) {
                const loc = (_path || _load_path()).default.join(dest, file);
                possibleExtraneous.add(loc);

                if ((yield lstat(loc)).isDirectory()) {
                  for (var _iterator5 = yield readdir(loc), _isArray5 = Array.isArray(_iterator5), _i5 = 0, _iterator5 = _isArray5 ? _iterator5 : _iterator5[Symbol.iterator]();;) {
                    var _ref7;

                    if (_isArray5) {
                      if (_i5 >= _iterator5.length) break;
                      _ref7 = _iterator5[_i5++];
                    } else {
                      _i5 = _iterator5.next();
                      if (_i5.done) break;
                      _ref7 = _i5.value;
                    }

                    const file = _ref7;

                    possibleExtraneous.add((_path || _load_path()).default.join(loc, file));
                  }
                }
              }
            }
          }
        }

        if (destStat && destStat.isSymbolicLink()) {
          yield (0, (_fsNormalized || _load_fsNormalized()).unlink)(dest);
          destStat = null;
        }

        if (srcStat.isSymbolicLink()) {
          onFresh();
          const linkname = yield readlink(src);
          actions.symlink.push({
            dest,
            linkname
          });
          onDone();
        } else if (srcStat.isDirectory()) {
          if (!destStat) {
            reporter.verbose(reporter.lang('verboseFileFolder', dest));
            yield mkdirp(dest);
          }

          const destParts = dest.split((_path || _load_path()).default.sep);
          while (destParts.length) {
            files.add(destParts.join((_path || _load_path()).default.sep).toLowerCase());
            destParts.pop();
          }

          // push all files to queue
          invariant(srcFiles, 'src files not initialised');
          let remaining = srcFiles.length;
          if (!remaining) {
            onDone();
          }
          for (var _iterator6 = srcFiles, _isArray6 = Array.isArray(_iterator6), _i6 = 0, _iterator6 = _isArray6 ? _iterator6 : _iterator6[Symbol.iterator]();;) {
            var _ref8;

            if (_isArray6) {
              if (_i6 >= _iterator6.length) break;
              _ref8 = _iterator6[_i6++];
            } else {
              _i6 = _iterator6.next();
              if (_i6.done) break;
              _ref8 = _i6.value;
            }

            const file = _ref8;

            queue.push({
              dest: (_path || _load_path()).default.join(dest, file),
              onFresh,
              onDone: function (_onDone) {
                function onDone() {
                  return _onDone.apply(this, arguments);
                }

                onDone.toString = function () {
                  return _onDone.toString();
                };

                return onDone;
              }(function () {
                if (--remaining === 0) {
                  onDone();
                }
              }),
              src: (_path || _load_path()).default.join(src, file)
            });
          }
        } else if (srcStat.isFile()) {
          onFresh();
          actions.file.push({
            src,
            dest,
            atime: srcStat.atime,
            mtime: srcStat.mtime,
            mode: srcStat.mode
          });
          onDone();
        } else {
          throw new Error(`unsure how to copy this: ${src}`);
        }
      });

      return function build(_x5) {
        return _ref5.apply(this, arguments);
      };
    })();

    const artifactFiles = new Set(events.artifactFiles || []);
    const files = new Set();

    // initialise events
    for (var _iterator = queue, _isArray = Array.isArray(_iterator), _i = 0, _iterator = _isArray ? _iterator : _iterator[Symbol.iterator]();;) {
      var _ref2;

      if (_isArray) {
        if (_i >= _iterator.length) break;
        _ref2 = _iterator[_i++];
      } else {
        _i = _iterator.next();
        if (_i.done) break;
        _ref2 = _i.value;
      }

      const item = _ref2;

      const onDone = item.onDone;
      item.onDone = function () {
        events.onProgress(item.dest);
        if (onDone) {
          onDone();
        }
      };
    }
    events.onStart(queue.length);

    // start building actions
    const actions = {
      file: [],
      symlink: [],
      link: []
    };

    // custom concurrency logic as we're always executing stacks of CONCURRENT_QUEUE_ITEMS queue items
    // at a time due to the requirement to push items onto the queue
    while (queue.length) {
      const items = queue.splice(0, CONCURRENT_QUEUE_ITEMS);
      yield Promise.all(items.map(build));
    }

    // simulate the existence of some files to prevent considering them extraneous
    for (var _iterator2 = artifactFiles, _isArray2 = Array.isArray(_iterator2), _i2 = 0, _iterator2 = _isArray2 ? _iterator2 : _iterator2[Symbol.iterator]();;) {
      var _ref3;

      if (_isArray2) {
        if (_i2 >= _iterator2.length) break;
        _ref3 = _iterator2[_i2++];
      } else {
        _i2 = _iterator2.next();
        if (_i2.done) break;
        _ref3 = _i2.value;
      }

      const file = _ref3;

      if (possibleExtraneous.has(file)) {
        reporter.verbose(reporter.lang('verboseFilePhantomExtraneous', file));
        possibleExtraneous.delete(file);
      }
    }

    for (var _iterator3 = possibleExtraneous, _isArray3 = Array.isArray(_iterator3), _i3 = 0, _iterator3 = _isArray3 ? _iterator3 : _iterator3[Symbol.iterator]();;) {
      var _ref4;

      if (_isArray3) {
        if (_i3 >= _iterator3.length) break;
        _ref4 = _iterator3[_i3++];
      } else {
        _i3 = _iterator3.next();
        if (_i3.done) break;
        _ref4 = _i3.value;
      }

      const loc = _ref4;

      if (files.has(loc.toLowerCase())) {
        possibleExtraneous.delete(loc);
      }
    }

    return actions;
  });

  return function buildActionsForCopy(_x, _x2, _x3, _x4) {
    return _ref.apply(this, arguments);
  };
})();

let buildActionsForHardlink = (() => {
  var _ref9 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (queue, events, possibleExtraneous, reporter) {

    //
    let build = (() => {
      var _ref13 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (data) {
        const src = data.src,
              dest = data.dest;

        const onFresh = data.onFresh || noop;
        const onDone = data.onDone || noop;
        if (files.has(dest.toLowerCase())) {
          // Fixes issue https://github.com/yarnpkg/yarn/issues/2734
          // When bulk hardlinking we have A -> B structure that we want to hardlink to A1 -> B1,
          // package-linker passes that modules A1 and B1 need to be hardlinked,
          // the recursive linking algorithm of A1 ends up scheduling files in B1 to be linked twice which will case
          // an exception.
          onDone();
          return;
        }
        files.add(dest.toLowerCase());

        if (events.ignoreBasenames.indexOf((_path || _load_path()).default.basename(src)) >= 0) {
          // ignored file
          return;
        }

        const srcStat = yield lstat(src);
        let srcFiles;

        if (srcStat.isDirectory()) {
          srcFiles = yield readdir(src);
        }

        const destExists = yield exists(dest);
        if (destExists) {
          const destStat = yield lstat(dest);

          const bothSymlinks = srcStat.isSymbolicLink() && destStat.isSymbolicLink();
          const bothFolders = srcStat.isDirectory() && destStat.isDirectory();
          const bothFiles = srcStat.isFile() && destStat.isFile();

          if (srcStat.mode !== destStat.mode) {
            try {
              yield access(dest, srcStat.mode);
            } catch (err) {
              // EINVAL access errors sometimes happen which shouldn't because node shouldn't be giving
              // us modes that aren't valid. investigate this, it's generally safe to proceed.
              reporter.verbose(err);
            }
          }

          if (bothFiles && artifactFiles.has(dest)) {
            // this file gets changed during build, likely by a custom install script. Don't bother checking it.
            onDone();
            reporter.verbose(reporter.lang('verboseFileSkipArtifact', src));
            return;
          }

          // correct hardlink
          if (bothFiles && srcStat.ino !== null && srcStat.ino === destStat.ino) {
            onDone();
            reporter.verbose(reporter.lang('verboseFileSkip', src, dest, srcStat.ino));
            return;
          }

          if (bothSymlinks) {
            const srcReallink = yield readlink(src);
            if (srcReallink === (yield readlink(dest))) {
              // if both symlinks are the same then we can continue on
              onDone();
              reporter.verbose(reporter.lang('verboseFileSkipSymlink', src, dest, srcReallink));
              return;
            }
          }

          if (bothFolders) {
            // mark files that aren't in this folder as possibly extraneous
            const destFiles = yield readdir(dest);
            invariant(srcFiles, 'src files not initialised');

            for (var _iterator10 = destFiles, _isArray10 = Array.isArray(_iterator10), _i10 = 0, _iterator10 = _isArray10 ? _iterator10 : _iterator10[Symbol.iterator]();;) {
              var _ref14;

              if (_isArray10) {
                if (_i10 >= _iterator10.length) break;
                _ref14 = _iterator10[_i10++];
              } else {
                _i10 = _iterator10.next();
                if (_i10.done) break;
                _ref14 = _i10.value;
              }

              const file = _ref14;

              if (srcFiles.indexOf(file) < 0) {
                const loc = (_path || _load_path()).default.join(dest, file);
                possibleExtraneous.add(loc);

                if ((yield lstat(loc)).isDirectory()) {
                  for (var _iterator11 = yield readdir(loc), _isArray11 = Array.isArray(_iterator11), _i11 = 0, _iterator11 = _isArray11 ? _iterator11 : _iterator11[Symbol.iterator]();;) {
                    var _ref15;

                    if (_isArray11) {
                      if (_i11 >= _iterator11.length) break;
                      _ref15 = _iterator11[_i11++];
                    } else {
                      _i11 = _iterator11.next();
                      if (_i11.done) break;
                      _ref15 = _i11.value;
                    }

                    const file = _ref15;

                    possibleExtraneous.add((_path || _load_path()).default.join(loc, file));
                  }
                }
              }
            }
          }
        }

        if (srcStat.isSymbolicLink()) {
          onFresh();
          const linkname = yield readlink(src);
          actions.symlink.push({
            dest,
            linkname
          });
          onDone();
        } else if (srcStat.isDirectory()) {
          reporter.verbose(reporter.lang('verboseFileFolder', dest));
          yield mkdirp(dest);

          const destParts = dest.split((_path || _load_path()).default.sep);
          while (destParts.length) {
            files.add(destParts.join((_path || _load_path()).default.sep).toLowerCase());
            destParts.pop();
          }

          // push all files to queue
          invariant(srcFiles, 'src files not initialised');
          let remaining = srcFiles.length;
          if (!remaining) {
            onDone();
          }
          for (var _iterator12 = srcFiles, _isArray12 = Array.isArray(_iterator12), _i12 = 0, _iterator12 = _isArray12 ? _iterator12 : _iterator12[Symbol.iterator]();;) {
            var _ref16;

            if (_isArray12) {
              if (_i12 >= _iterator12.length) break;
              _ref16 = _iterator12[_i12++];
            } else {
              _i12 = _iterator12.next();
              if (_i12.done) break;
              _ref16 = _i12.value;
            }

            const file = _ref16;

            queue.push({
              onFresh,
              src: (_path || _load_path()).default.join(src, file),
              dest: (_path || _load_path()).default.join(dest, file),
              onDone: function (_onDone2) {
                function onDone() {
                  return _onDone2.apply(this, arguments);
                }

                onDone.toString = function () {
                  return _onDone2.toString();
                };

                return onDone;
              }(function () {
                if (--remaining === 0) {
                  onDone();
                }
              })
            });
          }
        } else if (srcStat.isFile()) {
          onFresh();
          actions.link.push({
            src,
            dest,
            removeDest: destExists
          });
          onDone();
        } else {
          throw new Error(`unsure how to copy this: ${src}`);
        }
      });

      return function build(_x10) {
        return _ref13.apply(this, arguments);
      };
    })();

    const artifactFiles = new Set(events.artifactFiles || []);
    const files = new Set();

    // initialise events
    for (var _iterator7 = queue, _isArray7 = Array.isArray(_iterator7), _i7 = 0, _iterator7 = _isArray7 ? _iterator7 : _iterator7[Symbol.iterator]();;) {
      var _ref10;

      if (_isArray7) {
        if (_i7 >= _iterator7.length) break;
        _ref10 = _iterator7[_i7++];
      } else {
        _i7 = _iterator7.next();
        if (_i7.done) break;
        _ref10 = _i7.value;
      }

      const item = _ref10;

      const onDone = item.onDone || noop;
      item.onDone = function () {
        events.onProgress(item.dest);
        onDone();
      };
    }
    events.onStart(queue.length);

    // start building actions
    const actions = {
      file: [],
      symlink: [],
      link: []
    };

    // custom concurrency logic as we're always executing stacks of CONCURRENT_QUEUE_ITEMS queue items
    // at a time due to the requirement to push items onto the queue
    while (queue.length) {
      const items = queue.splice(0, CONCURRENT_QUEUE_ITEMS);
      yield Promise.all(items.map(build));
    }

    // simulate the existence of some files to prevent considering them extraneous
    for (var _iterator8 = artifactFiles, _isArray8 = Array.isArray(_iterator8), _i8 = 0, _iterator8 = _isArray8 ? _iterator8 : _iterator8[Symbol.iterator]();;) {
      var _ref11;

      if (_isArray8) {
        if (_i8 >= _iterator8.length) break;
        _ref11 = _iterator8[_i8++];
      } else {
        _i8 = _iterator8.next();
        if (_i8.done) break;
        _ref11 = _i8.value;
      }

      const file = _ref11;

      if (possibleExtraneous.has(file)) {
        reporter.verbose(reporter.lang('verboseFilePhantomExtraneous', file));
        possibleExtraneous.delete(file);
      }
    }

    for (var _iterator9 = possibleExtraneous, _isArray9 = Array.isArray(_iterator9), _i9 = 0, _iterator9 = _isArray9 ? _iterator9 : _iterator9[Symbol.iterator]();;) {
      var _ref12;

      if (_isArray9) {
        if (_i9 >= _iterator9.length) break;
        _ref12 = _iterator9[_i9++];
      } else {
        _i9 = _iterator9.next();
        if (_i9.done) break;
        _ref12 = _i9.value;
      }

      const loc = _ref12;

      if (files.has(loc.toLowerCase())) {
        possibleExtraneous.delete(loc);
      }
    }

    return actions;
  });

  return function buildActionsForHardlink(_x6, _x7, _x8, _x9) {
    return _ref9.apply(this, arguments);
  };
})();

let copyBulk = exports.copyBulk = (() => {
  var _ref17 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (queue, reporter, _events) {
    const events = {
      onStart: _events && _events.onStart || noop,
      onProgress: _events && _events.onProgress || noop,
      possibleExtraneous: _events ? _events.possibleExtraneous : new Set(),
      ignoreBasenames: _events && _events.ignoreBasenames || [],
      artifactFiles: _events && _events.artifactFiles || []
    };

    const actions = yield buildActionsForCopy(queue, events, events.possibleExtraneous, reporter);
    events.onStart(actions.file.length + actions.symlink.length + actions.link.length);

    const fileActions = actions.file;

    const currentlyWriting = new Map();

    yield (_promise || _load_promise()).queue(fileActions, (() => {
      var _ref18 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (data) {
        let writePromise;
        while (writePromise = currentlyWriting.get(data.dest)) {
          yield writePromise;
        }

        reporter.verbose(reporter.lang('verboseFileCopy', data.src, data.dest));
        const copier = (0, (_fsNormalized || _load_fsNormalized()).copyFile)(data, function () {
          return currentlyWriting.delete(data.dest);
        });
        currentlyWriting.set(data.dest, copier);
        events.onProgress(data.dest);
        return copier;
      });

      return function (_x14) {
        return _ref18.apply(this, arguments);
      };
    })(), CONCURRENT_QUEUE_ITEMS);

    // we need to copy symlinks last as they could reference files we were copying
    const symlinkActions = actions.symlink;
    yield (_promise || _load_promise()).queue(symlinkActions, function (data) {
      const linkname = (_path || _load_path()).default.resolve((_path || _load_path()).default.dirname(data.dest), data.linkname);
      reporter.verbose(reporter.lang('verboseFileSymlink', data.dest, linkname));
      return symlink(linkname, data.dest);
    });
  });

  return function copyBulk(_x11, _x12, _x13) {
    return _ref17.apply(this, arguments);
  };
})();

let hardlinkBulk = exports.hardlinkBulk = (() => {
  var _ref19 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (queue, reporter, _events) {
    const events = {
      onStart: _events && _events.onStart || noop,
      onProgress: _events && _events.onProgress || noop,
      possibleExtraneous: _events ? _events.possibleExtraneous : new Set(),
      artifactFiles: _events && _events.artifactFiles || [],
      ignoreBasenames: []
    };

    const actions = yield buildActionsForHardlink(queue, events, events.possibleExtraneous, reporter);
    events.onStart(actions.file.length + actions.symlink.length + actions.link.length);

    const fileActions = actions.link;

    yield (_promise || _load_promise()).queue(fileActions, (() => {
      var _ref20 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (data) {
        reporter.verbose(reporter.lang('verboseFileLink', data.src, data.dest));
        if (data.removeDest) {
          yield (0, (_fsNormalized || _load_fsNormalized()).unlink)(data.dest);
        }
        yield link(data.src, data.dest);
      });

      return function (_x18) {
        return _ref20.apply(this, arguments);
      };
    })(), CONCURRENT_QUEUE_ITEMS);

    // we need to copy symlinks last as they could reference files we were copying
    const symlinkActions = actions.symlink;
    yield (_promise || _load_promise()).queue(symlinkActions, function (data) {
      const linkname = (_path || _load_path()).default.resolve((_path || _load_path()).default.dirname(data.dest), data.linkname);
      reporter.verbose(reporter.lang('verboseFileSymlink', data.dest, linkname));
      return symlink(linkname, data.dest);
    });
  });

  return function hardlinkBulk(_x15, _x16, _x17) {
    return _ref19.apply(this, arguments);
  };
})();

let readFileAny = exports.readFileAny = (() => {
  var _ref21 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (files) {
    for (var _iterator13 = files, _isArray13 = Array.isArray(_iterator13), _i13 = 0, _iterator13 = _isArray13 ? _iterator13 : _iterator13[Symbol.iterator]();;) {
      var _ref22;

      if (_isArray13) {
        if (_i13 >= _iterator13.length) break;
        _ref22 = _iterator13[_i13++];
      } else {
        _i13 = _iterator13.next();
        if (_i13.done) break;
        _ref22 = _i13.value;
      }

      const file = _ref22;

      if (yield exists(file)) {
        return readFile(file);
      }
    }
    return null;
  });

  return function readFileAny(_x19) {
    return _ref21.apply(this, arguments);
  };
})();

let readJson = exports.readJson = (() => {
  var _ref23 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (loc) {
    return (yield readJsonAndFile(loc)).object;
  });

  return function readJson(_x20) {
    return _ref23.apply(this, arguments);
  };
})();

let readJsonAndFile = exports.readJsonAndFile = (() => {
  var _ref24 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (loc) {
    const file = yield readFile(loc);
    try {
      return {
        object: (0, (_map || _load_map()).default)(JSON.parse(stripBOM(file))),
        content: file
      };
    } catch (err) {
      err.message = `${loc}: ${err.message}`;
      throw err;
    }
  });

  return function readJsonAndFile(_x21) {
    return _ref24.apply(this, arguments);
  };
})();

let find = exports.find = (() => {
  var _ref25 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (filename, dir) {
    const parts = dir.split((_path || _load_path()).default.sep);

    while (parts.length) {
      const loc = parts.concat(filename).join((_path || _load_path()).default.sep);

      if (yield exists(loc)) {
        return loc;
      } else {
        parts.pop();
      }
    }

    return false;
  });

  return function find(_x22, _x23) {
    return _ref25.apply(this, arguments);
  };
})();

let symlink = exports.symlink = (() => {
  var _ref26 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (src, dest) {
    try {
      const stats = yield lstat(dest);
      if (stats.isSymbolicLink()) {
        const resolved = yield realpath(dest);
        if (resolved === src) {
          return;
        }
      }
    } catch (err) {
      if (err.code !== 'ENOENT') {
        throw err;
      }
    }
    // We use rimraf for unlink which never throws an ENOENT on missing target
    yield (0, (_fsNormalized || _load_fsNormalized()).unlink)(dest);

    if (process.platform === 'win32') {
      // use directory junctions if possible on win32, this requires absolute paths
      yield fsSymlink(src, dest, 'junction');
    } else {
      // use relative paths otherwise which will be retained if the directory is moved
      let relative;
      try {
        relative = (_path || _load_path()).default.relative((_fs || _load_fs()).default.realpathSync((_path || _load_path()).default.dirname(dest)), (_fs || _load_fs()).default.realpathSync(src));
      } catch (err) {
        if (err.code !== 'ENOENT') {
          throw err;
        }
        relative = (_path || _load_path()).default.relative((_path || _load_path()).default.dirname(dest), src);
      }
      // When path.relative returns an empty string for the current directory, we should instead use
      // '.', which is a valid fs.symlink target.
      yield fsSymlink(relative || '.', dest);
    }
  });

  return function symlink(_x24, _x25) {
    return _ref26.apply(this, arguments);
  };
})();

let walk = exports.walk = (() => {
  var _ref27 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (dir, relativeDir, ignoreBasenames = new Set()) {
    let files = [];

    let filenames = yield readdir(dir);
    if (ignoreBasenames.size) {
      filenames = filenames.filter(function (name) {
        return !ignoreBasenames.has(name);
      });
    }

    for (var _iterator14 = filenames, _isArray14 = Array.isArray(_iterator14), _i14 = 0, _iterator14 = _isArray14 ? _iterator14 : _iterator14[Symbol.iterator]();;) {
      var _ref28;

      if (_isArray14) {
        if (_i14 >= _iterator14.length) break;
        _ref28 = _iterator14[_i14++];
      } else {
        _i14 = _iterator14.next();
        if (_i14.done) break;
        _ref28 = _i14.value;
      }

      const name = _ref28;

      const relative = relativeDir ? (_path || _load_path()).default.join(relativeDir, name) : name;
      const loc = (_path || _load_path()).default.join(dir, name);
      const stat = yield lstat(loc);

      files.push({
        relative,
        basename: name,
        absolute: loc,
        mtime: +stat.mtime
      });

      if (stat.isDirectory()) {
        files = files.concat((yield walk(loc, relative, ignoreBasenames)));
      }
    }

    return files;
  });

  return function walk(_x26, _x27) {
    return _ref27.apply(this, arguments);
  };
})();

let getFileSizeOnDisk = exports.getFileSizeOnDisk = (() => {
  var _ref29 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (loc) {
    const stat = yield lstat(loc);
    const size = stat.size,
          blockSize = stat.blksize;


    return Math.ceil(size / blockSize) * blockSize;
  });

  return function getFileSizeOnDisk(_x28) {
    return _ref29.apply(this, arguments);
  };
})();

let getEolFromFile = (() => {
  var _ref30 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (path) {
    if (!(yield exists(path))) {
      return undefined;
    }

    const buffer = yield readFileBuffer(path);

    for (let i = 0; i < buffer.length; ++i) {
      if (buffer[i] === cr) {
        return '\r\n';
      }
      if (buffer[i] === lf) {
        return '\n';
      }
    }
    return undefined;
  });

  return function getEolFromFile(_x29) {
    return _ref30.apply(this, arguments);
  };
})();

let writeFilePreservingEol = exports.writeFilePreservingEol = (() => {
  var _ref31 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (path, data) {
    const eol = (yield getEolFromFile(path)) || (_os || _load_os()).default.EOL;
    if (eol !== '\n') {
      data = data.replace(/\n/g, eol);
    }
    yield writeFile(path, data);
  });

  return function writeFilePreservingEol(_x30, _x31) {
    return _ref31.apply(this, arguments);
  };
})();

let hardlinksWork = exports.hardlinksWork = (() => {
  var _ref32 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (dir) {
    const filename = 'test-file' + Math.random();
    const file = (_path || _load_path()).default.join(dir, filename);
    const fileLink = (_path || _load_path()).default.join(dir, filename + '-link');
    try {
      yield writeFile(file, 'test');
      yield link(file, fileLink);
    } catch (err) {
      return false;
    } finally {
      yield (0, (_fsNormalized || _load_fsNormalized()).unlink)(file);
      yield (0, (_fsNormalized || _load_fsNormalized()).unlink)(fileLink);
    }
    return true;
  });

  return function hardlinksWork(_x32) {
    return _ref32.apply(this, arguments);
  };
})();

// not a strict polyfill for Node's fs.mkdtemp


let makeTempDir = exports.makeTempDir = (() => {
  var _ref33 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (prefix) {
    const dir = (_path || _load_path()).default.join((_os || _load_os()).default.tmpdir(), `yarn-${prefix || ''}-${Date.now()}-${Math.random()}`);
    yield (0, (_fsNormalized || _load_fsNormalized()).unlink)(dir);
    yield mkdirp(dir);
    return dir;
  });

  return function makeTempDir(_x33) {
    return _ref33.apply(this, arguments);
  };
})();

let readFirstAvailableStream = exports.readFirstAvailableStream = (() => {
  var _ref34 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (paths) {
    for (var _iterator15 = paths, _isArray15 = Array.isArray(_iterator15), _i15 = 0, _iterator15 = _isArray15 ? _iterator15 : _iterator15[Symbol.iterator]();;) {
      var _ref35;

      if (_isArray15) {
        if (_i15 >= _iterator15.length) break;
        _ref35 = _iterator15[_i15++];
      } else {
        _i15 = _iterator15.next();
        if (_i15.done) break;
        _ref35 = _i15.value;
      }

      const path = _ref35;

      try {
        const fd = yield open(path, 'r');
        return (_fs || _load_fs()).default.createReadStream(path, { fd });
      } catch (err) {
        // Try the next one
      }
    }
    return null;
  });

  return function readFirstAvailableStream(_x34) {
    return _ref34.apply(this, arguments);
  };
})();

let getFirstSuitableFolder = exports.getFirstSuitableFolder = (() => {
  var _ref36 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (paths, mode = constants.W_OK | constants.X_OK) {
    const result = {
      skipped: [],
      folder: null
    };

    for (var _iterator16 = paths, _isArray16 = Array.isArray(_iterator16), _i16 = 0, _iterator16 = _isArray16 ? _iterator16 : _iterator16[Symbol.iterator]();;) {
      var _ref37;

      if (_isArray16) {
        if (_i16 >= _iterator16.length) break;
        _ref37 = _iterator16[_i16++];
      } else {
        _i16 = _iterator16.next();
        if (_i16.done) break;
        _ref37 = _i16.value;
      }

      const folder = _ref37;

      try {
        yield mkdirp(folder);
        yield access(folder, mode);

        result.folder = folder;

        return result;
      } catch (error) {
        result.skipped.push({
          error,
          folder
        });
      }
    }
    return result;
  });

  return function getFirstSuitableFolder(_x35) {
    return _ref36.apply(this, arguments);
  };
})();

exports.copy = copy;
exports.readFile = readFile;
exports.readFileRaw = readFileRaw;
exports.normalizeOS = normalizeOS;

var _fs;

function _load_fs() {
  return _fs = _interopRequireDefault(__webpack_require__(3));
}

var _glob;

function _load_glob() {
  return _glob = _interopRequireDefault(__webpack_require__(75));
}

var _os;

function _load_os() {
  return _os = _interopRequireDefault(__webpack_require__(36));
}

var _path;

function _load_path() {
  return _path = _interopRequireDefault(__webpack_require__(0));
}

var _blockingQueue;

function _load_blockingQueue() {
  return _blockingQueue = _interopRequireDefault(__webpack_require__(84));
}

var _promise;

function _load_promise() {
  return _promise = _interopRequireWildcard(__webpack_require__(40));
}

var _promise2;

function _load_promise2() {
  return _promise2 = __webpack_require__(40);
}

var _map;

function _load_map() {
  return _map = _interopRequireDefault(__webpack_require__(20));
}

var _fsNormalized;

function _load_fsNormalized() {
  return _fsNormalized = __webpack_require__(164);
}

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const constants = exports.constants = typeof (_fs || _load_fs()).default.constants !== 'undefined' ? (_fs || _load_fs()).default.constants : {
  R_OK: (_fs || _load_fs()).default.R_OK,
  W_OK: (_fs || _load_fs()).default.W_OK,
  X_OK: (_fs || _load_fs()).default.X_OK
};

const lockQueue = exports.lockQueue = new (_blockingQueue || _load_blockingQueue()).default('fs lock');

const readFileBuffer = exports.readFileBuffer = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.readFile);
const open = exports.open = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.open);
const writeFile = exports.writeFile = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.writeFile);
const readlink = exports.readlink = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.readlink);
const realpath = exports.realpath = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.realpath);
const readdir = exports.readdir = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.readdir);
const rename = exports.rename = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.rename);
const access = exports.access = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.access);
const stat = exports.stat = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.stat);
const mkdirp = exports.mkdirp = (0, (_promise2 || _load_promise2()).promisify)(__webpack_require__(116));
const exists = exports.exists = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.exists, true);
const lstat = exports.lstat = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.lstat);
const chmod = exports.chmod = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.chmod);
const link = exports.link = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.link);
const glob = exports.glob = (0, (_promise2 || _load_promise2()).promisify)((_glob || _load_glob()).default);
exports.unlink = (_fsNormalized || _load_fsNormalized()).unlink;

// fs.copyFile uses the native file copying instructions on the system, performing much better
// than any JS-based solution and consumes fewer resources. Repeated testing to fine tune the
// concurrency level revealed 128 as the sweet spot on a quad-core, 16 CPU Intel system with SSD.

const CONCURRENT_QUEUE_ITEMS = (_fs || _load_fs()).default.copyFile ? 128 : 4;

const fsSymlink = (0, (_promise2 || _load_promise2()).promisify)((_fs || _load_fs()).default.symlink);
const invariant = __webpack_require__(7);
const stripBOM = __webpack_require__(122);

const noop = () => {};

function copy(src, dest, reporter) {
  return copyBulk([{ src, dest }], reporter);
}

function _readFile(loc, encoding) {
  return new Promise((resolve, reject) => {
    (_fs || _load_fs()).default.readFile(loc, encoding, function (err, content) {
      if (err) {
        reject(err);
      } else {
        resolve(content);
      }
    });
  });
}

function readFile(loc) {
  return _readFile(loc, 'utf8').then(normalizeOS);
}

function readFileRaw(loc) {
  return _readFile(loc, 'binary');
}

function normalizeOS(body) {
  return body.replace(/\r\n/g, '\n');
}

const cr = '\r'.charCodeAt(0);
const lf = '\n'.charCodeAt(0);

/***/ }),
/* 6 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getPathKey = getPathKey;
const os = __webpack_require__(36);
const path = __webpack_require__(0);
const userHome = __webpack_require__(45).default;

var _require = __webpack_require__(171);

const getCacheDir = _require.getCacheDir,
      getConfigDir = _require.getConfigDir,
      getDataDir = _require.getDataDir;

const isWebpackBundle = __webpack_require__(227);

const DEPENDENCY_TYPES = exports.DEPENDENCY_TYPES = ['devDependencies', 'dependencies', 'optionalDependencies', 'peerDependencies'];
const RESOLUTIONS = exports.RESOLUTIONS = 'resolutions';
const MANIFEST_FIELDS = exports.MANIFEST_FIELDS = [RESOLUTIONS, ...DEPENDENCY_TYPES];

const SUPPORTED_NODE_VERSIONS = exports.SUPPORTED_NODE_VERSIONS = '^4.8.0 || ^5.7.0 || ^6.2.2 || >=8.0.0';

const YARN_REGISTRY = exports.YARN_REGISTRY = 'https://registry.yarnpkg.com';

const YARN_DOCS = exports.YARN_DOCS = 'https://yarnpkg.com/en/docs/cli/';
const YARN_INSTALLER_SH = exports.YARN_INSTALLER_SH = 'https://yarnpkg.com/install.sh';
const YARN_INSTALLER_MSI = exports.YARN_INSTALLER_MSI = 'https://yarnpkg.com/latest.msi';

const SELF_UPDATE_VERSION_URL = exports.SELF_UPDATE_VERSION_URL = 'https://yarnpkg.com/latest-version';

// cache version, bump whenever we make backwards incompatible changes
const CACHE_VERSION = exports.CACHE_VERSION = 2;

// lockfile version, bump whenever we make backwards incompatible changes
const LOCKFILE_VERSION = exports.LOCKFILE_VERSION = 1;

// max amount of network requests to perform concurrently
const NETWORK_CONCURRENCY = exports.NETWORK_CONCURRENCY = 8;

// HTTP timeout used when downloading packages
const NETWORK_TIMEOUT = exports.NETWORK_TIMEOUT = 30 * 1000; // in milliseconds

// max amount of child processes to execute concurrently
const CHILD_CONCURRENCY = exports.CHILD_CONCURRENCY = 5;

const REQUIRED_PACKAGE_KEYS = exports.REQUIRED_PACKAGE_KEYS = ['name', 'version', '_uid'];

function getPreferredCacheDirectories() {
  const preferredCacheDirectories = [getCacheDir()];

  if (process.getuid) {
    // $FlowFixMe: process.getuid exists, dammit
    preferredCacheDirectories.push(path.join(os.tmpdir(), `.yarn-cache-${process.getuid()}`));
  }

  preferredCacheDirectories.push(path.join(os.tmpdir(), `.yarn-cache`));

  return preferredCacheDirectories;
}

const PREFERRED_MODULE_CACHE_DIRECTORIES = exports.PREFERRED_MODULE_CACHE_DIRECTORIES = getPreferredCacheDirectories();
const CONFIG_DIRECTORY = exports.CONFIG_DIRECTORY = getConfigDir();
const DATA_DIRECTORY = exports.DATA_DIRECTORY = getDataDir();
const LINK_REGISTRY_DIRECTORY = exports.LINK_REGISTRY_DIRECTORY = path.join(DATA_DIRECTORY, 'link');
const GLOBAL_MODULE_DIRECTORY = exports.GLOBAL_MODULE_DIRECTORY = path.join(DATA_DIRECTORY, 'global');

const NODE_BIN_PATH = exports.NODE_BIN_PATH = process.execPath;
const YARN_BIN_PATH = exports.YARN_BIN_PATH = getYarnBinPath();

// Webpack needs to be configured with node.__dirname/__filename = false
function getYarnBinPath() {
  if (isWebpackBundle) {
    return __filename;
  } else {
    return path.join(__dirname, '..', 'bin', 'yarn.js');
  }
}

const NODE_MODULES_FOLDER = exports.NODE_MODULES_FOLDER = 'node_modules';
const NODE_PACKAGE_JSON = exports.NODE_PACKAGE_JSON = 'package.json';

const POSIX_GLOBAL_PREFIX = exports.POSIX_GLOBAL_PREFIX = `${process.env.DESTDIR || ''}/usr/local`;
const FALLBACK_GLOBAL_PREFIX = exports.FALLBACK_GLOBAL_PREFIX = path.join(userHome, '.yarn');

const META_FOLDER = exports.META_FOLDER = '.yarn-meta';
const INTEGRITY_FILENAME = exports.INTEGRITY_FILENAME = '.yarn-integrity';
const LOCKFILE_FILENAME = exports.LOCKFILE_FILENAME = 'yarn.lock';
const METADATA_FILENAME = exports.METADATA_FILENAME = '.yarn-metadata.json';
const TARBALL_FILENAME = exports.TARBALL_FILENAME = '.yarn-tarball.tgz';
const CLEAN_FILENAME = exports.CLEAN_FILENAME = '.yarnclean';

const NPM_LOCK_FILENAME = exports.NPM_LOCK_FILENAME = 'package-lock.json';
const NPM_SHRINKWRAP_FILENAME = exports.NPM_SHRINKWRAP_FILENAME = 'npm-shrinkwrap.json';

const DEFAULT_INDENT = exports.DEFAULT_INDENT = '  ';
const SINGLE_INSTANCE_PORT = exports.SINGLE_INSTANCE_PORT = 31997;
const SINGLE_INSTANCE_FILENAME = exports.SINGLE_INSTANCE_FILENAME = '.yarn-single-instance';

const ENV_PATH_KEY = exports.ENV_PATH_KEY = getPathKey(process.platform, process.env);

function getPathKey(platform, env) {
  let pathKey = 'PATH';

  // windows calls its path "Path" usually, but this is not guaranteed.
  if (platform === 'win32') {
    pathKey = 'Path';

    for (const key in env) {
      if (key.toLowerCase() === 'path') {
        pathKey = key;
      }
    }
  }

  return pathKey;
}

const VERSION_COLOR_SCHEME = exports.VERSION_COLOR_SCHEME = {
  major: 'red',
  premajor: 'red',
  minor: 'yellow',
  preminor: 'yellow',
  patch: 'green',
  prepatch: 'green',
  prerelease: 'red',
  unchanged: 'white',
  unknown: 'red'
};

/***/ }),
/* 7 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
/**
 * Copyright (c) 2013-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */



/**
 * Use invariant() to assert state which your program assumes to be true.
 *
 * Provide sprintf-style format (only %s is supported) and arguments
 * to provide information about what broke and what you were
 * expecting.
 *
 * The invariant message will be stripped in production, but the invariant
 * will remain to ensure logic does not differ in production.
 */

var NODE_ENV = process.env.NODE_ENV;

var invariant = function(condition, format, a, b, c, d, e, f) {
  if (NODE_ENV !== 'production') {
    if (format === undefined) {
      throw new Error('invariant requires an error message argument');
    }
  }

  if (!condition) {
    var error;
    if (format === undefined) {
      error = new Error(
        'Minified exception occurred; use the non-minified dev environment ' +
        'for the full error message and additional helpful warnings.'
      );
    } else {
      var args = [a, b, c, d, e, f];
      var argIndex = 0;
      error = new Error(
        format.replace(/%s/g, function() { return args[argIndex++]; })
      );
      error.name = 'Invariant Violation';
    }

    error.framesToPop = 1; // we don't care about invariant's own frame
    throw error;
  }
};

module.exports = invariant;


/***/ }),
/* 8 */,
/* 9 */
/***/ (function(module, exports) {

module.exports = require("crypto");

/***/ }),
/* 10 */,
/* 11 */
/***/ (function(module, exports) {

// https://github.com/zloirock/core-js/issues/86#issuecomment-115759028
var global = module.exports = typeof window != 'undefined' && window.Math == Math
  ? window : typeof self != 'undefined' && self.Math == Math ? self
  // eslint-disable-next-line no-new-func
  : Function('return this')();
if (typeof __g == 'number') __g = global; // eslint-disable-line no-undef


/***/ }),
/* 12 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.sortAlpha = sortAlpha;
exports.entries = entries;
exports.removePrefix = removePrefix;
exports.removeSuffix = removeSuffix;
exports.addSuffix = addSuffix;
exports.hyphenate = hyphenate;
exports.camelCase = camelCase;
exports.compareSortedArrays = compareSortedArrays;
exports.sleep = sleep;
const _camelCase = __webpack_require__(176);

function sortAlpha(a, b) {
  // sort alphabetically in a deterministic way
  const shortLen = Math.min(a.length, b.length);
  for (let i = 0; i < shortLen; i++) {
    const aChar = a.charCodeAt(i);
    const bChar = b.charCodeAt(i);
    if (aChar !== bChar) {
      return aChar - bChar;
    }
  }
  return a.length - b.length;
}

function entries(obj) {
  const entries = [];
  if (obj) {
    for (const key in obj) {
      entries.push([key, obj[key]]);
    }
  }
  return entries;
}

function removePrefix(pattern, prefix) {
  if (pattern.startsWith(prefix)) {
    pattern = pattern.slice(prefix.length);
  }

  return pattern;
}

function removeSuffix(pattern, suffix) {
  if (pattern.endsWith(suffix)) {
    return pattern.slice(0, -suffix.length);
  }

  return pattern;
}

function addSuffix(pattern, suffix) {
  if (!pattern.endsWith(suffix)) {
    return pattern + suffix;
  }

  return pattern;
}

function hyphenate(str) {
  return str.replace(/[A-Z]/g, match => {
    return '-' + match.charAt(0).toLowerCase();
  });
}

function camelCase(str) {
  if (/[A-Z]/.test(str)) {
    return null;
  } else {
    return _camelCase(str);
  }
}

function compareSortedArrays(array1, array2) {
  if (array1.length !== array2.length) {
    return false;
  }
  for (let i = 0, len = array1.length; i < len; i++) {
    if (array1[i] !== array2[i]) {
      return false;
    }
  }
  return true;
}

function sleep(ms) {
  return new Promise(resolve => {
    setTimeout(resolve, ms);
  });
}

/***/ }),
/* 13 */
/***/ (function(module, exports, __webpack_require__) {

var store = __webpack_require__(107)('wks');
var uid = __webpack_require__(111);
var Symbol = __webpack_require__(11).Symbol;
var USE_SYMBOL = typeof Symbol == 'function';

var $exports = module.exports = function (name) {
  return store[name] || (store[name] =
    USE_SYMBOL && Symbol[name] || (USE_SYMBOL ? Symbol : uid)('Symbol.' + name));
};

$exports.store = store;


/***/ }),
/* 14 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.stringify = exports.parse = undefined;

var _asyncToGenerator2;

function _load_asyncToGenerator() {
  return _asyncToGenerator2 = _interopRequireDefault(__webpack_require__(1));
}

var _parse;

function _load_parse() {
  return _parse = __webpack_require__(81);
}

Object.defineProperty(exports, 'parse', {
  enumerable: true,
  get: function get() {
    return _interopRequireDefault(_parse || _load_parse()).default;
  }
});

var _stringify;

function _load_stringify() {
  return _stringify = __webpack_require__(150);
}

Object.defineProperty(exports, 'stringify', {
  enumerable: true,
  get: function get() {
    return _interopRequireDefault(_stringify || _load_stringify()).default;
  }
});
exports.implodeEntry = implodeEntry;
exports.explodeEntry = explodeEntry;

var _misc;

function _load_misc() {
  return _misc = __webpack_require__(12);
}

var _normalizePattern;

function _load_normalizePattern() {
  return _normalizePattern = __webpack_require__(29);
}

var _parse2;

function _load_parse2() {
  return _parse2 = _interopRequireDefault(__webpack_require__(81));
}

var _constants;

function _load_constants() {
  return _constants = __webpack_require__(6);
}

var _fs;

function _load_fs() {
  return _fs = _interopRequireWildcard(__webpack_require__(5));
}

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const invariant = __webpack_require__(7);

const path = __webpack_require__(0);
const ssri = __webpack_require__(55);

function getName(pattern) {
  return (0, (_normalizePattern || _load_normalizePattern()).normalizePattern)(pattern).name;
}

function blankObjectUndefined(obj) {
  return obj && Object.keys(obj).length ? obj : undefined;
}

function keyForRemote(remote) {
  return remote.resolved || (remote.reference && remote.hash ? `${remote.reference}#${remote.hash}` : null);
}

function serializeIntegrity(integrity) {
  // We need this because `Integrity.toString()` does not use sorting to ensure a stable string output
  // See https://git.io/vx2Hy
  return integrity.toString().split(' ').sort().join(' ');
}

function implodeEntry(pattern, obj) {
  const inferredName = getName(pattern);
  const integrity = obj.integrity ? serializeIntegrity(obj.integrity) : '';
  const imploded = {
    name: inferredName === obj.name ? undefined : obj.name,
    version: obj.version,
    uid: obj.uid === obj.version ? undefined : obj.uid,
    resolved: obj.resolved,
    registry: obj.registry === 'npm' ? undefined : obj.registry,
    dependencies: blankObjectUndefined(obj.dependencies),
    optionalDependencies: blankObjectUndefined(obj.optionalDependencies),
    permissions: blankObjectUndefined(obj.permissions),
    prebuiltVariants: blankObjectUndefined(obj.prebuiltVariants)
  };
  if (integrity) {
    imploded.integrity = integrity;
  }
  return imploded;
}

function explodeEntry(pattern, obj) {
  obj.optionalDependencies = obj.optionalDependencies || {};
  obj.dependencies = obj.dependencies || {};
  obj.uid = obj.uid || obj.version;
  obj.permissions = obj.permissions || {};
  obj.registry = obj.registry || 'npm';
  obj.name = obj.name || getName(pattern);
  const integrity = obj.integrity;
  if (integrity && integrity.isIntegrity) {
    obj.integrity = ssri.parse(integrity);
  }
  return obj;
}

class Lockfile {
  constructor({ cache, source, parseResultType } = {}) {
    this.source = source || '';
    this.cache = cache;
    this.parseResultType = parseResultType;
  }

  // source string if the `cache` was parsed


  // if true, we're parsing an old yarn file and need to update integrity fields
  hasEntriesExistWithoutIntegrity() {
    if (!this.cache) {
      return false;
    }

    for (const key in this.cache) {
      // $FlowFixMe - `this.cache` is clearly defined at this point
      if (!/^.*@(file:|http)/.test(key) && this.cache[key] && !this.cache[key].integrity) {
        return true;
      }
    }

    return false;
  }

  static fromDirectory(dir, reporter) {
    return (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* () {
      // read the manifest in this directory
      const lockfileLoc = path.join(dir, (_constants || _load_constants()).LOCKFILE_FILENAME);

      let lockfile;
      let rawLockfile = '';
      let parseResult;

      if (yield (_fs || _load_fs()).exists(lockfileLoc)) {
        rawLockfile = yield (_fs || _load_fs()).readFile(lockfileLoc);
        parseResult = (0, (_parse2 || _load_parse2()).default)(rawLockfile, lockfileLoc);

        if (reporter) {
          if (parseResult.type === 'merge') {
            reporter.info(reporter.lang('lockfileMerged'));
          } else if (parseResult.type === 'conflict') {
            reporter.warn(reporter.lang('lockfileConflict'));
          }
        }

        lockfile = parseResult.object;
      } else if (reporter) {
        reporter.info(reporter.lang('noLockfileFound'));
      }

      return new Lockfile({ cache: lockfile, source: rawLockfile, parseResultType: parseResult && parseResult.type });
    })();
  }

  getLocked(pattern) {
    const cache = this.cache;
    if (!cache) {
      return undefined;
    }

    const shrunk = pattern in cache && cache[pattern];

    if (typeof shrunk === 'string') {
      return this.getLocked(shrunk);
    } else if (shrunk) {
      explodeEntry(pattern, shrunk);
      return shrunk;
    }

    return undefined;
  }

  removePattern(pattern) {
    const cache = this.cache;
    if (!cache) {
      return;
    }
    delete cache[pattern];
  }

  getLockfile(patterns) {
    const lockfile = {};
    const seen = new Map();

    // order by name so that lockfile manifest is assigned to the first dependency with this manifest
    // the others that have the same remoteKey will just refer to the first
    // ordering allows for consistency in lockfile when it is serialized
    const sortedPatternsKeys = Object.keys(patterns).sort((_misc || _load_misc()).sortAlpha);

    for (var _iterator = sortedPatternsKeys, _isArray = Array.isArray(_iterator), _i = 0, _iterator = _isArray ? _iterator : _iterator[Symbol.iterator]();;) {
      var _ref;

      if (_isArray) {
        if (_i >= _iterator.length) break;
        _ref = _iterator[_i++];
      } else {
        _i = _iterator.next();
        if (_i.done) break;
        _ref = _i.value;
      }

      const pattern = _ref;

      const pkg = patterns[pattern];
      const remote = pkg._remote,
            ref = pkg._reference;

      invariant(ref, 'Package is missing a reference');
      invariant(remote, 'Package is missing a remote');

      const remoteKey = keyForRemote(remote);
      const seenPattern = remoteKey && seen.get(remoteKey);
      if (seenPattern) {
        // no point in duplicating it
        lockfile[pattern] = seenPattern;

        // if we're relying on our name being inferred and two of the patterns have
        // different inferred names then we need to set it
        if (!seenPattern.name && getName(pattern) !== pkg.name) {
          seenPattern.name = pkg.name;
        }
        continue;
      }
      const obj = implodeEntry(pattern, {
        name: pkg.name,
        version: pkg.version,
        uid: pkg._uid,
        resolved: remote.resolved,
        integrity: remote.integrity,
        registry: remote.registry,
        dependencies: pkg.dependencies,
        peerDependencies: pkg.peerDependencies,
        optionalDependencies: pkg.optionalDependencies,
        permissions: ref.permissions,
        prebuiltVariants: pkg.prebuiltVariants
      });

      lockfile[pattern] = obj;

      if (remoteKey) {
        seen.set(remoteKey, obj);
      }
    }

    return lockfile;
  }
}
exports.default = Lockfile;

/***/ }),
/* 15 */,
/* 16 */,
/* 17 */
/***/ (function(module, exports) {

module.exports = require("stream");

/***/ }),
/* 18 */,
/* 19 */,
/* 20 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = nullify;
function nullify(obj = {}) {
  if (Array.isArray(obj)) {
    for (var _iterator = obj, _isArray = Array.isArray(_iterator), _i = 0, _iterator = _isArray ? _iterator : _iterator[Symbol.iterator]();;) {
      var _ref;

      if (_isArray) {
        if (_i >= _iterator.length) break;
        _ref = _iterator[_i++];
      } else {
        _i = _iterator.next();
        if (_i.done) break;
        _ref = _i.value;
      }

      const item = _ref;

      nullify(item);
    }
  } else if (obj !== null && typeof obj === 'object' || typeof obj === 'function') {
    Object.setPrototypeOf(obj, null);

    // for..in can only be applied to 'object', not 'function'
    if (typeof obj === 'object') {
      for (const key in obj) {
        nullify(obj[key]);
      }
    }
  }

  return obj;
}

/***/ }),
/* 21 */,
/* 22 */
/***/ (function(module, exports) {

module.exports = require("assert");

/***/ }),
/* 23 */
/***/ (function(module, exports) {

var core = module.exports = { version: '2.5.7' };
if (typeof __e == 'number') __e = core; // eslint-disable-line no-undef


/***/ }),
/* 24 */,
/* 25 */,
/* 26 */,
/* 27 */
/***/ (function(module, exports, __webpack_require__) {

var isObject = __webpack_require__(34);
module.exports = function (it) {
  if (!isObject(it)) throw TypeError(it + ' is not an object!');
  return it;
};


/***/ }),
/* 28 */,
/* 29 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.normalizePattern = normalizePattern;

/**
 * Explode and normalize a pattern into its name and range.
 */

function normalizePattern(pattern) {
  let hasVersion = false;
  let range = 'latest';
  let name = pattern;

  // if we're a scope then remove the @ and add it back later
  let isScoped = false;
  if (name[0] === '@') {
    isScoped = true;
    name = name.slice(1);
  }

  // take first part as the name
  const parts = name.split('@');
  if (parts.length > 1) {
    name = parts.shift();
    range = parts.join('@');

    if (range) {
      hasVersion = true;
    } else {
      range = '*';
    }
  }

  // add back @ scope suffix
  if (isScoped) {
    name = `@${name}`;
  }

  return { name, range, hasVersion };
}

/***/ }),
/* 30 */,
/* 31 */
/***/ (function(module, exports, __webpack_require__) {

var dP = __webpack_require__(50);
var createDesc = __webpack_require__(106);
module.exports = __webpack_require__(33) ? function (object, key, value) {
  return dP.f(object, key, createDesc(1, value));
} : function (object, key, value) {
  object[key] = value;
  return object;
};


/***/ }),
/* 32 */
/***/ (function(module, exports, __webpack_require__) {

/* eslint-disable node/no-deprecated-api */
var buffer = __webpack_require__(63)
var Buffer = buffer.Buffer

// alternative to using Object.keys for old browsers
function copyProps (src, dst) {
  for (var key in src) {
    dst[key] = src[key]
  }
}
if (Buffer.from && Buffer.alloc && Buffer.allocUnsafe && Buffer.allocUnsafeSlow) {
  module.exports = buffer
} else {
  // Copy properties from require('buffer')
  copyProps(buffer, exports)
  exports.Buffer = SafeBuffer
}

function SafeBuffer (arg, encodingOrOffset, length) {
  return Buffer(arg, encodingOrOffset, length)
}

// Copy static methods from Buffer
copyProps(Buffer, SafeBuffer)

SafeBuffer.from = function (arg, encodingOrOffset, length) {
  if (typeof arg === 'number') {
    throw new TypeError('Argument must not be a number')
  }
  return Buffer(arg, encodingOrOffset, length)
}

SafeBuffer.alloc = function (size, fill, encoding) {
  if (typeof size !== 'number') {
    throw new TypeError('Argument must be a number')
  }
  var buf = Buffer(size)
  if (fill !== undefined) {
    if (typeof encoding === 'string') {
      buf.fill(fill, encoding)
    } else {
      buf.fill(fill)
    }
  } else {
    buf.fill(0)
  }
  return buf
}

SafeBuffer.allocUnsafe = function (size) {
  if (typeof size !== 'number') {
    throw new TypeError('Argument must be a number')
  }
  return Buffer(size)
}

SafeBuffer.allocUnsafeSlow = function (size) {
  if (typeof size !== 'number') {
    throw new TypeError('Argument must be a number')
  }
  return buffer.SlowBuffer(size)
}


/***/ }),
/* 33 */
/***/ (function(module, exports, __webpack_require__) {

// Thank's IE8 for his funny defineProperty
module.exports = !__webpack_require__(85)(function () {
  return Object.defineProperty({}, 'a', { get: function () { return 7; } }).a != 7;
});


/***/ }),
/* 34 */
/***/ (function(module, exports) {

module.exports = function (it) {
  return typeof it === 'object' ? it !== null : typeof it === 'function';
};


/***/ }),
/* 35 */
/***/ (function(module, exports) {

module.exports = {};


/***/ }),
/* 36 */
/***/ (function(module, exports) {

module.exports = require("os");

/***/ }),
/* 37 */,
/* 38 */,
/* 39 */,
/* 40 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.wait = wait;
exports.promisify = promisify;
exports.queue = queue;
function wait(delay) {
  return new Promise(resolve => {
    setTimeout(resolve, delay);
  });
}

function promisify(fn, firstData) {
  return function (...args) {
    return new Promise(function (resolve, reject) {
      args.push(function (err, ...result) {
        let res = result;

        if (result.length <= 1) {
          res = result[0];
        }

        if (firstData) {
          res = err;
          err = null;
        }

        if (err) {
          reject(err);
        } else {
          resolve(res);
        }
      });

      fn.apply(null, args);
    });
  };
}

function queue(arr, promiseProducer, concurrency = Infinity) {
  concurrency = Math.min(concurrency, arr.length);

  // clone
  arr = arr.slice();

  const results = [];
  let total = arr.length;
  if (!total) {
    return Promise.resolve(results);
  }

  return new Promise((resolve, reject) => {
    for (let i = 0; i < concurrency; i++) {
      next();
    }

    function next() {
      const item = arr.shift();
      const promise = promiseProducer(item);

      promise.then(function (result) {
        results.push(result);

        total--;
        if (total === 0) {
          resolve(results);
        } else {
          if (arr.length) {
            next();
          }
        }
      }, reject);
    }
  });
}

/***/ }),
/* 41 */
/***/ (function(module, exports, __webpack_require__) {

var global = __webpack_require__(11);
var core = __webpack_require__(23);
var ctx = __webpack_require__(48);
var hide = __webpack_require__(31);
var has = __webpack_require__(49);
var PROTOTYPE = 'prototype';

var $export = function (type, name, source) {
  var IS_FORCED = type & $export.F;
  var IS_GLOBAL = type & $export.G;
  var IS_STATIC = type & $export.S;
  var IS_PROTO = type & $export.P;
  var IS_BIND = type & $export.B;
  var IS_WRAP = type & $export.W;
  var exports = IS_GLOBAL ? core : core[name] || (core[name] = {});
  var expProto = exports[PROTOTYPE];
  var target = IS_GLOBAL ? global : IS_STATIC ? global[name] : (global[name] || {})[PROTOTYPE];
  var key, own, out;
  if (IS_GLOBAL) source = name;
  for (key in source) {
    // contains in native
    own = !IS_FORCED && target && target[key] !== undefined;
    if (own && has(exports, key)) continue;
    // export native or passed
    out = own ? target[key] : source[key];
    // prevent global pollution for namespaces
    exports[key] = IS_GLOBAL && typeof target[key] != 'function' ? source[key]
    // bind timers to global for call from export context
    : IS_BIND && own ? ctx(out, global)
    // wrap global constructors for prevent change them in library
    : IS_WRAP && target[key] == out ? (function (C) {
      var F = function (a, b, c) {
        if (this instanceof C) {
          switch (arguments.length) {
            case 0: return new C();
            case 1: return new C(a);
            case 2: return new C(a, b);
          } return new C(a, b, c);
        } return C.apply(this, arguments);
      };
      F[PROTOTYPE] = C[PROTOTYPE];
      return F;
    // make static versions for prototype methods
    })(out) : IS_PROTO && typeof out == 'function' ? ctx(Function.call, out) : out;
    // export proto methods to core.%CONSTRUCTOR%.methods.%NAME%
    if (IS_PROTO) {
      (exports.virtual || (exports.virtual = {}))[key] = out;
      // export proto methods to core.%CONSTRUCTOR%.prototype.%NAME%
      if (type & $export.R && expProto && !expProto[key]) hide(expProto, key, out);
    }
  }
};
// type bitmap
$export.F = 1;   // forced
$export.G = 2;   // global
$export.S = 4;   // static
$export.P = 8;   // proto
$export.B = 16;  // bind
$export.W = 32;  // wrap
$export.U = 64;  // safe
$export.R = 128; // real proto method for `library`
module.exports = $export;


/***/ }),
/* 42 */
/***/ (function(module, exports, __webpack_require__) {

try {
  var util = __webpack_require__(2);
  if (typeof util.inherits !== 'function') throw '';
  module.exports = util.inherits;
} catch (e) {
  module.exports = __webpack_require__(224);
}


/***/ }),
/* 43 */,
/* 44 */,
/* 45 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.home = undefined;

var _rootUser;

function _load_rootUser() {
  return _rootUser = _interopRequireDefault(__webpack_require__(169));
}

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const path = __webpack_require__(0);

const home = exports.home = __webpack_require__(36).homedir();

const userHomeDir = (_rootUser || _load_rootUser()).default ? path.resolve('/usr/local/share') : home;

exports.default = userHomeDir;

/***/ }),
/* 46 */
/***/ (function(module, exports) {

module.exports = function (it) {
  if (typeof it != 'function') throw TypeError(it + ' is not a function!');
  return it;
};


/***/ }),
/* 47 */
/***/ (function(module, exports) {

var toString = {}.toString;

module.exports = function (it) {
  return toString.call(it).slice(8, -1);
};


/***/ }),
/* 48 */
/***/ (function(module, exports, __webpack_require__) {

// optional / simple context binding
var aFunction = __webpack_require__(46);
module.exports = function (fn, that, length) {
  aFunction(fn);
  if (that === undefined) return fn;
  switch (length) {
    case 1: return function (a) {
      return fn.call(that, a);
    };
    case 2: return function (a, b) {
      return fn.call(that, a, b);
    };
    case 3: return function (a, b, c) {
      return fn.call(that, a, b, c);
    };
  }
  return function (/* ...args */) {
    return fn.apply(that, arguments);
  };
};


/***/ }),
/* 49 */
/***/ (function(module, exports) {

var hasOwnProperty = {}.hasOwnProperty;
module.exports = function (it, key) {
  return hasOwnProperty.call(it, key);
};


/***/ }),
/* 50 */
/***/ (function(module, exports, __webpack_require__) {

var anObject = __webpack_require__(27);
var IE8_DOM_DEFINE = __webpack_require__(184);
var toPrimitive = __webpack_require__(201);
var dP = Object.defineProperty;

exports.f = __webpack_require__(33) ? Object.defineProperty : function defineProperty(O, P, Attributes) {
  anObject(O);
  P = toPrimitive(P, true);
  anObject(Attributes);
  if (IE8_DOM_DEFINE) try {
    return dP(O, P, Attributes);
  } catch (e) { /* empty */ }
  if ('get' in Attributes || 'set' in Attributes) throw TypeError('Accessors not supported!');
  if ('value' in Attributes) O[P] = Attributes.value;
  return O;
};


/***/ }),
/* 51 */,
/* 52 */,
/* 53 */,
/* 54 */
/***/ (function(module, exports) {

module.exports = require("events");

/***/ }),
/* 55 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


const Buffer = __webpack_require__(32).Buffer

const crypto = __webpack_require__(9)
const Transform = __webpack_require__(17).Transform

const SPEC_ALGORITHMS = ['sha256', 'sha384', 'sha512']

const BASE64_REGEX = /^[a-z0-9+/]+(?:=?=?)$/i
const SRI_REGEX = /^([^-]+)-([^?]+)([?\S*]*)$/
const STRICT_SRI_REGEX = /^([^-]+)-([A-Za-z0-9+/=]{44,88})(\?[\x21-\x7E]*)*$/
const VCHAR_REGEX = /^[\x21-\x7E]+$/

class Hash {
  get isHash () { return true }
  constructor (hash, opts) {
    const strict = !!(opts && opts.strict)
    this.source = hash.trim()
    // 3.1. Integrity metadata (called "Hash" by ssri)
    // https://w3c.github.io/webappsec-subresource-integrity/#integrity-metadata-description
    const match = this.source.match(
      strict
      ? STRICT_SRI_REGEX
      : SRI_REGEX
    )
    if (!match) { return }
    if (strict && !SPEC_ALGORITHMS.some(a => a === match[1])) { return }
    this.algorithm = match[1]
    this.digest = match[2]

    const rawOpts = match[3]
    this.options = rawOpts ? rawOpts.slice(1).split('?') : []
  }
  hexDigest () {
    return this.digest && Buffer.from(this.digest, 'base64').toString('hex')
  }
  toJSON () {
    return this.toString()
  }
  toString (opts) {
    if (opts && opts.strict) {
      // Strict mode enforces the standard as close to the foot of the
      // letter as it can.
      if (!(
        // The spec has very restricted productions for algorithms.
        // https://www.w3.org/TR/CSP2/#source-list-syntax
        SPEC_ALGORITHMS.some(x => x === this.algorithm) &&
        // Usually, if someone insists on using a "different" base64, we
        // leave it as-is, since there's multiple standards, and the
        // specified is not a URL-safe variant.
        // https://www.w3.org/TR/CSP2/#base64_value
        this.digest.match(BASE64_REGEX) &&
        // Option syntax is strictly visual chars.
        // https://w3c.github.io/webappsec-subresource-integrity/#grammardef-option-expression
        // https://tools.ietf.org/html/rfc5234#appendix-B.1
        (this.options || []).every(opt => opt.match(VCHAR_REGEX))
      )) {
        return ''
      }
    }
    const options = this.options && this.options.length
    ? `?${this.options.join('?')}`
    : ''
    return `${this.algorithm}-${this.digest}${options}`
  }
}

class Integrity {
  get isIntegrity () { return true }
  toJSON () {
    return this.toString()
  }
  toString (opts) {
    opts = opts || {}
    let sep = opts.sep || ' '
    if (opts.strict) {
      // Entries must be separated by whitespace, according to spec.
      sep = sep.replace(/\S+/g, ' ')
    }
    return Object.keys(this).map(k => {
      return this[k].map(hash => {
        return Hash.prototype.toString.call(hash, opts)
      }).filter(x => x.length).join(sep)
    }).filter(x => x.length).join(sep)
  }
  concat (integrity, opts) {
    const other = typeof integrity === 'string'
    ? integrity
    : stringify(integrity, opts)
    return parse(`${this.toString(opts)} ${other}`, opts)
  }
  hexDigest () {
    return parse(this, {single: true}).hexDigest()
  }
  match (integrity, opts) {
    const other = parse(integrity, opts)
    const algo = other.pickAlgorithm(opts)
    return (
      this[algo] &&
      other[algo] &&
      this[algo].find(hash =>
        other[algo].find(otherhash =>
          hash.digest === otherhash.digest
        )
      )
    ) || false
  }
  pickAlgorithm (opts) {
    const pickAlgorithm = (opts && opts.pickAlgorithm) || getPrioritizedHash
    const keys = Object.keys(this)
    if (!keys.length) {
      throw new Error(`No algorithms available for ${
        JSON.stringify(this.toString())
      }`)
    }
    return keys.reduce((acc, algo) => {
      return pickAlgorithm(acc, algo) || acc
    })
  }
}

module.exports.parse = parse
function parse (sri, opts) {
  opts = opts || {}
  if (typeof sri === 'string') {
    return _parse(sri, opts)
  } else if (sri.algorithm && sri.digest) {
    const fullSri = new Integrity()
    fullSri[sri.algorithm] = [sri]
    return _parse(stringify(fullSri, opts), opts)
  } else {
    return _parse(stringify(sri, opts), opts)
  }
}

function _parse (integrity, opts) {
  // 3.4.3. Parse metadata
  // https://w3c.github.io/webappsec-subresource-integrity/#parse-metadata
  if (opts.single) {
    return new Hash(integrity, opts)
  }
  return integrity.trim().split(/\s+/).reduce((acc, string) => {
    const hash = new Hash(string, opts)
    if (hash.algorithm && hash.digest) {
      const algo = hash.algorithm
      if (!acc[algo]) { acc[algo] = [] }
      acc[algo].push(hash)
    }
    return acc
  }, new Integrity())
}

module.exports.stringify = stringify
function stringify (obj, opts) {
  if (obj.algorithm && obj.digest) {
    return Hash.prototype.toString.call(obj, opts)
  } else if (typeof obj === 'string') {
    return stringify(parse(obj, opts), opts)
  } else {
    return Integrity.prototype.toString.call(obj, opts)
  }
}

module.exports.fromHex = fromHex
function fromHex (hexDigest, algorithm, opts) {
  const optString = (opts && opts.options && opts.options.length)
  ? `?${opts.options.join('?')}`
  : ''
  return parse(
    `${algorithm}-${
      Buffer.from(hexDigest, 'hex').toString('base64')
    }${optString}`, opts
  )
}

module.exports.fromData = fromData
function fromData (data, opts) {
  opts = opts || {}
  const algorithms = opts.algorithms || ['sha512']
  const optString = opts.options && opts.options.length
  ? `?${opts.options.join('?')}`
  : ''
  return algorithms.reduce((acc, algo) => {
    const digest = crypto.createHash(algo).update(data).digest('base64')
    const hash = new Hash(
      `${algo}-${digest}${optString}`,
       opts
    )
    if (hash.algorithm && hash.digest) {
      const algo = hash.algorithm
      if (!acc[algo]) { acc[algo] = [] }
      acc[algo].push(hash)
    }
    return acc
  }, new Integrity())
}

module.exports.fromStream = fromStream
function fromStream (stream, opts) {
  opts = opts || {}
  const P = opts.Promise || Promise
  const istream = integrityStream(opts)
  return new P((resolve, reject) => {
    stream.pipe(istream)
    stream.on('error', reject)
    istream.on('error', reject)
    let sri
    istream.on('integrity', s => { sri = s })
    istream.on('end', () => resolve(sri))
    istream.on('data', () => {})
  })
}

module.exports.checkData = checkData
function checkData (data, sri, opts) {
  opts = opts || {}
  sri = parse(sri, opts)
  if (!Object.keys(sri).length) {
    if (opts.error) {
      throw Object.assign(
        new Error('No valid integrity hashes to check against'), {
          code: 'EINTEGRITY'
        }
      )
    } else {
      return false
    }
  }
  const algorithm = sri.pickAlgorithm(opts)
  const digest = crypto.createHash(algorithm).update(data).digest('base64')
  const newSri = parse({algorithm, digest})
  const match = newSri.match(sri, opts)
  if (match || !opts.error) {
    return match
  } else if (typeof opts.size === 'number' && (data.length !== opts.size)) {
    const err = new Error(`data size mismatch when checking ${sri}.\n  Wanted: ${opts.size}\n  Found: ${data.length}`)
    err.code = 'EBADSIZE'
    err.found = data.length
    err.expected = opts.size
    err.sri = sri
    throw err
  } else {
    const err = new Error(`Integrity checksum failed when using ${algorithm}: Wanted ${sri}, but got ${newSri}. (${data.length} bytes)`)
    err.code = 'EINTEGRITY'
    err.found = newSri
    err.expected = sri
    err.algorithm = algorithm
    err.sri = sri
    throw err
  }
}

module.exports.checkStream = checkStream
function checkStream (stream, sri, opts) {
  opts = opts || {}
  const P = opts.Promise || Promise
  const checker = integrityStream(Object.assign({}, opts, {
    integrity: sri
  }))
  return new P((resolve, reject) => {
    stream.pipe(checker)
    stream.on('error', reject)
    checker.on('error', reject)
    let sri
    checker.on('verified', s => { sri = s })
    checker.on('end', () => resolve(sri))
    checker.on('data', () => {})
  })
}

module.exports.integrityStream = integrityStream
function integrityStream (opts) {
  opts = opts || {}
  // For verification
  const sri = opts.integrity && parse(opts.integrity, opts)
  const goodSri = sri && Object.keys(sri).length
  const algorithm = goodSri && sri.pickAlgorithm(opts)
  const digests = goodSri && sri[algorithm]
  // Calculating stream
  const algorithms = Array.from(
    new Set(
      (opts.algorithms || ['sha512'])
      .concat(algorithm ? [algorithm] : [])
    )
  )
  const hashes = algorithms.map(crypto.createHash)
  let streamSize = 0
  const stream = new Transform({
    transform (chunk, enc, cb) {
      streamSize += chunk.length
      hashes.forEach(h => h.update(chunk, enc))
      cb(null, chunk, enc)
    }
  }).on('end', () => {
    const optString = (opts.options && opts.options.length)
    ? `?${opts.options.join('?')}`
    : ''
    const newSri = parse(hashes.map((h, i) => {
      return `${algorithms[i]}-${h.digest('base64')}${optString}`
    }).join(' '), opts)
    // Integrity verification mode
    const match = goodSri && newSri.match(sri, opts)
    if (typeof opts.size === 'number' && streamSize !== opts.size) {
      const err = new Error(`stream size mismatch when checking ${sri}.\n  Wanted: ${opts.size}\n  Found: ${streamSize}`)
      err.code = 'EBADSIZE'
      err.found = streamSize
      err.expected = opts.size
      err.sri = sri
      stream.emit('error', err)
    } else if (opts.integrity && !match) {
      const err = new Error(`${sri} integrity checksum failed when using ${algorithm}: wanted ${digests} but got ${newSri}. (${streamSize} bytes)`)
      err.code = 'EINTEGRITY'
      err.found = newSri
      err.expected = digests
      err.algorithm = algorithm
      err.sri = sri
      stream.emit('error', err)
    } else {
      stream.emit('size', streamSize)
      stream.emit('integrity', newSri)
      match && stream.emit('verified', match)
    }
  })
  return stream
}

module.exports.create = createIntegrity
function createIntegrity (opts) {
  opts = opts || {}
  const algorithms = opts.algorithms || ['sha512']
  const optString = opts.options && opts.options.length
  ? `?${opts.options.join('?')}`
  : ''

  const hashes = algorithms.map(crypto.createHash)

  return {
    update: function (chunk, enc) {
      hashes.forEach(h => h.update(chunk, enc))
      return this
    },
    digest: function (enc) {
      const integrity = algorithms.reduce((acc, algo) => {
        const digest = hashes.shift().digest('base64')
        const hash = new Hash(
          `${algo}-${digest}${optString}`,
          opts
        )
        if (hash.algorithm && hash.digest) {
          const algo = hash.algorithm
          if (!acc[algo]) { acc[algo] = [] }
          acc[algo].push(hash)
        }
        return acc
      }, new Integrity())

      return integrity
    }
  }
}

const NODE_HASHES = new Set(crypto.getHashes())

// This is a Best Effort™ at a reasonable priority for hash algos
const DEFAULT_PRIORITY = [
  'md5', 'whirlpool', 'sha1', 'sha224', 'sha256', 'sha384', 'sha512',
  // TODO - it's unclear _which_ of these Node will actually use as its name
  //        for the algorithm, so we guesswork it based on the OpenSSL names.
  'sha3',
  'sha3-256', 'sha3-384', 'sha3-512',
  'sha3_256', 'sha3_384', 'sha3_512'
].filter(algo => NODE_HASHES.has(algo))

function getPrioritizedHash (algo1, algo2) {
  return DEFAULT_PRIORITY.indexOf(algo1.toLowerCase()) >= DEFAULT_PRIORITY.indexOf(algo2.toLowerCase())
  ? algo1
  : algo2
}


/***/ }),
/* 56 */,
/* 57 */,
/* 58 */,
/* 59 */,
/* 60 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = minimatch
minimatch.Minimatch = Minimatch

var path = { sep: '/' }
try {
  path = __webpack_require__(0)
} catch (er) {}

var GLOBSTAR = minimatch.GLOBSTAR = Minimatch.GLOBSTAR = {}
var expand = __webpack_require__(175)

var plTypes = {
  '!': { open: '(?:(?!(?:', close: '))[^/]*?)'},
  '?': { open: '(?:', close: ')?' },
  '+': { open: '(?:', close: ')+' },
  '*': { open: '(?:', close: ')*' },
  '@': { open: '(?:', close: ')' }
}

// any single thing other than /
// don't need to escape / when using new RegExp()
var qmark = '[^/]'

// * => any number of characters
var star = qmark + '*?'

// ** when dots are allowed.  Anything goes, except .. and .
// not (^ or / followed by one or two dots followed by $ or /),
// followed by anything, any number of times.
var twoStarDot = '(?:(?!(?:\\\/|^)(?:\\.{1,2})($|\\\/)).)*?'

// not a ^ or / followed by a dot,
// followed by anything, any number of times.
var twoStarNoDot = '(?:(?!(?:\\\/|^)\\.).)*?'

// characters that need to be escaped in RegExp.
var reSpecials = charSet('().*{}+?[]^$\\!')

// "abc" -> { a:true, b:true, c:true }
function charSet (s) {
  return s.split('').reduce(function (set, c) {
    set[c] = true
    return set
  }, {})
}

// normalizes slashes.
var slashSplit = /\/+/

minimatch.filter = filter
function filter (pattern, options) {
  options = options || {}
  return function (p, i, list) {
    return minimatch(p, pattern, options)
  }
}

function ext (a, b) {
  a = a || {}
  b = b || {}
  var t = {}
  Object.keys(b).forEach(function (k) {
    t[k] = b[k]
  })
  Object.keys(a).forEach(function (k) {
    t[k] = a[k]
  })
  return t
}

minimatch.defaults = function (def) {
  if (!def || !Object.keys(def).length) return minimatch

  var orig = minimatch

  var m = function minimatch (p, pattern, options) {
    return orig.minimatch(p, pattern, ext(def, options))
  }

  m.Minimatch = function Minimatch (pattern, options) {
    return new orig.Minimatch(pattern, ext(def, options))
  }

  return m
}

Minimatch.defaults = function (def) {
  if (!def || !Object.keys(def).length) return Minimatch
  return minimatch.defaults(def).Minimatch
}

function minimatch (p, pattern, options) {
  if (typeof pattern !== 'string') {
    throw new TypeError('glob pattern string required')
  }

  if (!options) options = {}

  // shortcut: comments match nothing.
  if (!options.nocomment && pattern.charAt(0) === '#') {
    return false
  }

  // "" only matches ""
  if (pattern.trim() === '') return p === ''

  return new Minimatch(pattern, options).match(p)
}

function Minimatch (pattern, options) {
  if (!(this instanceof Minimatch)) {
    return new Minimatch(pattern, options)
  }

  if (typeof pattern !== 'string') {
    throw new TypeError('glob pattern string required')
  }

  if (!options) options = {}
  pattern = pattern.trim()

  // windows support: need to use /, not \
  if (path.sep !== '/') {
    pattern = pattern.split(path.sep).join('/')
  }

  this.options = options
  this.set = []
  this.pattern = pattern
  this.regexp = null
  this.negate = false
  this.comment = false
  this.empty = false

  // make the set of regexps etc.
  this.make()
}

Minimatch.prototype.debug = function () {}

Minimatch.prototype.make = make
function make () {
  // don't do it more than once.
  if (this._made) return

  var pattern = this.pattern
  var options = this.options

  // empty patterns and comments match nothing.
  if (!options.nocomment && pattern.charAt(0) === '#') {
    this.comment = true
    return
  }
  if (!pattern) {
    this.empty = true
    return
  }

  // step 1: figure out negation, etc.
  this.parseNegate()

  // step 2: expand braces
  var set = this.globSet = this.braceExpand()

  if (options.debug) this.debug = console.error

  this.debug(this.pattern, set)

  // step 3: now we have a set, so turn each one into a series of path-portion
  // matching patterns.
  // These will be regexps, except in the case of "**", which is
  // set to the GLOBSTAR object for globstar behavior,
  // and will not contain any / characters
  set = this.globParts = set.map(function (s) {
    return s.split(slashSplit)
  })

  this.debug(this.pattern, set)

  // glob --> regexps
  set = set.map(function (s, si, set) {
    return s.map(this.parse, this)
  }, this)

  this.debug(this.pattern, set)

  // filter out everything that didn't compile properly.
  set = set.filter(function (s) {
    return s.indexOf(false) === -1
  })

  this.debug(this.pattern, set)

  this.set = set
}

Minimatch.prototype.parseNegate = parseNegate
function parseNegate () {
  var pattern = this.pattern
  var negate = false
  var options = this.options
  var negateOffset = 0

  if (options.nonegate) return

  for (var i = 0, l = pattern.length
    ; i < l && pattern.charAt(i) === '!'
    ; i++) {
    negate = !negate
    negateOffset++
  }

  if (negateOffset) this.pattern = pattern.substr(negateOffset)
  this.negate = negate
}

// Brace expansion:
// a{b,c}d -> abd acd
// a{b,}c -> abc ac
// a{0..3}d -> a0d a1d a2d a3d
// a{b,c{d,e}f}g -> abg acdfg acefg
// a{b,c}d{e,f}g -> abdeg acdeg abdeg abdfg
//
// Invalid sets are not expanded.
// a{2..}b -> a{2..}b
// a{b}c -> a{b}c
minimatch.braceExpand = function (pattern, options) {
  return braceExpand(pattern, options)
}

Minimatch.prototype.braceExpand = braceExpand

function braceExpand (pattern, options) {
  if (!options) {
    if (this instanceof Minimatch) {
      options = this.options
    } else {
      options = {}
    }
  }

  pattern = typeof pattern === 'undefined'
    ? this.pattern : pattern

  if (typeof pattern === 'undefined') {
    throw new TypeError('undefined pattern')
  }

  if (options.nobrace ||
    !pattern.match(/\{.*\}/)) {
    // shortcut. no need to expand.
    return [pattern]
  }

  return expand(pattern)
}

// parse a component of the expanded set.
// At this point, no pattern may contain "/" in it
// so we're going to return a 2d array, where each entry is the full
// pattern, split on '/', and then turned into a regular expression.
// A regexp is made at the end which joins each array with an
// escaped /, and another full one which joins each regexp with |.
//
// Following the lead of Bash 4.1, note that "**" only has special meaning
// when it is the *only* thing in a path portion.  Otherwise, any series
// of * is equivalent to a single *.  Globstar behavior is enabled by
// default, and can be disabled by setting options.noglobstar.
Minimatch.prototype.parse = parse
var SUBPARSE = {}
function parse (pattern, isSub) {
  if (pattern.length > 1024 * 64) {
    throw new TypeError('pattern is too long')
  }

  var options = this.options

  // shortcuts
  if (!options.noglobstar && pattern === '**') return GLOBSTAR
  if (pattern === '') return ''

  var re = ''
  var hasMagic = !!options.nocase
  var escaping = false
  // ? => one single character
  var patternListStack = []
  var negativeLists = []
  var stateChar
  var inClass = false
  var reClassStart = -1
  var classStart = -1
  // . and .. never match anything that doesn't start with .,
  // even when options.dot is set.
  var patternStart = pattern.charAt(0) === '.' ? '' // anything
  // not (start or / followed by . or .. followed by / or end)
  : options.dot ? '(?!(?:^|\\\/)\\.{1,2}(?:$|\\\/))'
  : '(?!\\.)'
  var self = this

  function clearStateChar () {
    if (stateChar) {
      // we had some state-tracking character
      // that wasn't consumed by this pass.
      switch (stateChar) {
        case '*':
          re += star
          hasMagic = true
        break
        case '?':
          re += qmark
          hasMagic = true
        break
        default:
          re += '\\' + stateChar
        break
      }
      self.debug('clearStateChar %j %j', stateChar, re)
      stateChar = false
    }
  }

  for (var i = 0, len = pattern.length, c
    ; (i < len) && (c = pattern.charAt(i))
    ; i++) {
    this.debug('%s\t%s %s %j', pattern, i, re, c)

    // skip over any that are escaped.
    if (escaping && reSpecials[c]) {
      re += '\\' + c
      escaping = false
      continue
    }

    switch (c) {
      case '/':
        // completely not allowed, even escaped.
        // Should already be path-split by now.
        return false

      case '\\':
        clearStateChar()
        escaping = true
      continue

      // the various stateChar values
      // for the "extglob" stuff.
      case '?':
      case '*':
      case '+':
      case '@':
      case '!':
        this.debug('%s\t%s %s %j <-- stateChar', pattern, i, re, c)

        // all of those are literals inside a class, except that
        // the glob [!a] means [^a] in regexp
        if (inClass) {
          this.debug('  in class')
          if (c === '!' && i === classStart + 1) c = '^'
          re += c
          continue
        }

        // if we already have a stateChar, then it means
        // that there was something like ** or +? in there.
        // Handle the stateChar, then proceed with this one.
        self.debug('call clearStateChar %j', stateChar)
        clearStateChar()
        stateChar = c
        // if extglob is disabled, then +(asdf|foo) isn't a thing.
        // just clear the statechar *now*, rather than even diving into
        // the patternList stuff.
        if (options.noext) clearStateChar()
      continue

      case '(':
        if (inClass) {
          re += '('
          continue
        }

        if (!stateChar) {
          re += '\\('
          continue
        }

        patternListStack.push({
          type: stateChar,
          start: i - 1,
          reStart: re.length,
          open: plTypes[stateChar].open,
          close: plTypes[stateChar].close
        })
        // negation is (?:(?!js)[^/]*)
        re += stateChar === '!' ? '(?:(?!(?:' : '(?:'
        this.debug('plType %j %j', stateChar, re)
        stateChar = false
      continue

      case ')':
        if (inClass || !patternListStack.length) {
          re += '\\)'
          continue
        }

        clearStateChar()
        hasMagic = true
        var pl = patternListStack.pop()
        // negation is (?:(?!js)[^/]*)
        // The others are (?:<pattern>)<type>
        re += pl.close
        if (pl.type === '!') {
          negativeLists.push(pl)
        }
        pl.reEnd = re.length
      continue

      case '|':
        if (inClass || !patternListStack.length || escaping) {
          re += '\\|'
          escaping = false
          continue
        }

        clearStateChar()
        re += '|'
      continue

      // these are mostly the same in regexp and glob
      case '[':
        // swallow any state-tracking char before the [
        clearStateChar()

        if (inClass) {
          re += '\\' + c
          continue
        }

        inClass = true
        classStart = i
        reClassStart = re.length
        re += c
      continue

      case ']':
        //  a right bracket shall lose its special
        //  meaning and represent itself in
        //  a bracket expression if it occurs
        //  first in the list.  -- POSIX.2 2.8.3.2
        if (i === classStart + 1 || !inClass) {
          re += '\\' + c
          escaping = false
          continue
        }

        // handle the case where we left a class open.
        // "[z-a]" is valid, equivalent to "\[z-a\]"
        if (inClass) {
          // split where the last [ was, make sure we don't have
          // an invalid re. if so, re-walk the contents of the
          // would-be class to re-translate any characters that
          // were passed through as-is
          // TODO: It would probably be faster to determine this
          // without a try/catch and a new RegExp, but it's tricky
          // to do safely.  For now, this is safe and works.
          var cs = pattern.substring(classStart + 1, i)
          try {
            RegExp('[' + cs + ']')
          } catch (er) {
            // not a valid class!
            var sp = this.parse(cs, SUBPARSE)
            re = re.substr(0, reClassStart) + '\\[' + sp[0] + '\\]'
            hasMagic = hasMagic || sp[1]
            inClass = false
            continue
          }
        }

        // finish up the class.
        hasMagic = true
        inClass = false
        re += c
      continue

      default:
        // swallow any state char that wasn't consumed
        clearStateChar()

        if (escaping) {
          // no need
          escaping = false
        } else if (reSpecials[c]
          && !(c === '^' && inClass)) {
          re += '\\'
        }

        re += c

    } // switch
  } // for

  // handle the case where we left a class open.
  // "[abc" is valid, equivalent to "\[abc"
  if (inClass) {
    // split where the last [ was, and escape it
    // this is a huge pita.  We now have to re-walk
    // the contents of the would-be class to re-translate
    // any characters that were passed through as-is
    cs = pattern.substr(classStart + 1)
    sp = this.parse(cs, SUBPARSE)
    re = re.substr(0, reClassStart) + '\\[' + sp[0]
    hasMagic = hasMagic || sp[1]
  }

  // handle the case where we had a +( thing at the *end*
  // of the pattern.
  // each pattern list stack adds 3 chars, and we need to go through
  // and escape any | chars that were passed through as-is for the regexp.
  // Go through and escape them, taking care not to double-escape any
  // | chars that were already escaped.
  for (pl = patternListStack.pop(); pl; pl = patternListStack.pop()) {
    var tail = re.slice(pl.reStart + pl.open.length)
    this.debug('setting tail', re, pl)
    // maybe some even number of \, then maybe 1 \, followed by a |
    tail = tail.replace(/((?:\\{2}){0,64})(\\?)\|/g, function (_, $1, $2) {
      if (!$2) {
        // the | isn't already escaped, so escape it.
        $2 = '\\'
      }

      // need to escape all those slashes *again*, without escaping the
      // one that we need for escaping the | character.  As it works out,
      // escaping an even number of slashes can be done by simply repeating
      // it exactly after itself.  That's why this trick works.
      //
      // I am sorry that you have to see this.
      return $1 + $1 + $2 + '|'
    })

    this.debug('tail=%j\n   %s', tail, tail, pl, re)
    var t = pl.type === '*' ? star
      : pl.type === '?' ? qmark
      : '\\' + pl.type

    hasMagic = true
    re = re.slice(0, pl.reStart) + t + '\\(' + tail
  }

  // handle trailing things that only matter at the very end.
  clearStateChar()
  if (escaping) {
    // trailing \\
    re += '\\\\'
  }

  // only need to apply the nodot start if the re starts with
  // something that could conceivably capture a dot
  var addPatternStart = false
  switch (re.charAt(0)) {
    case '.':
    case '[':
    case '(': addPatternStart = true
  }

  // Hack to work around lack of negative lookbehind in JS
  // A pattern like: *.!(x).!(y|z) needs to ensure that a name
  // like 'a.xyz.yz' doesn't match.  So, the first negative
  // lookahead, has to look ALL the way ahead, to the end of
  // the pattern.
  for (var n = negativeLists.length - 1; n > -1; n--) {
    var nl = negativeLists[n]

    var nlBefore = re.slice(0, nl.reStart)
    var nlFirst = re.slice(nl.reStart, nl.reEnd - 8)
    var nlLast = re.slice(nl.reEnd - 8, nl.reEnd)
    var nlAfter = re.slice(nl.reEnd)

    nlLast += nlAfter

    // Handle nested stuff like *(*.js|!(*.json)), where open parens
    // mean that we should *not* include the ) in the bit that is considered
    // "after" the negated section.
    var openParensBefore = nlBefore.split('(').length - 1
    var cleanAfter = nlAfter
    for (i = 0; i < openParensBefore; i++) {
      cleanAfter = cleanAfter.replace(/\)[+*?]?/, '')
    }
    nlAfter = cleanAfter

    var dollar = ''
    if (nlAfter === '' && isSub !== SUBPARSE) {
      dollar = '$'
    }
    var newRe = nlBefore + nlFirst + nlAfter + dollar + nlLast
    re = newRe
  }

  // if the re is not "" at this point, then we need to make sure
  // it doesn't match against an empty path part.
  // Otherwise a/* will match a/, which it should not.
  if (re !== '' && hasMagic) {
    re = '(?=.)' + re
  }

  if (addPatternStart) {
    re = patternStart + re
  }

  // parsing just a piece of a larger pattern.
  if (isSub === SUBPARSE) {
    return [re, hasMagic]
  }

  // skip the regexp for non-magical patterns
  // unescape anything in it, though, so that it'll be
  // an exact match against a file etc.
  if (!hasMagic) {
    return globUnescape(pattern)
  }

  var flags = options.nocase ? 'i' : ''
  try {
    var regExp = new RegExp('^' + re + '$', flags)
  } catch (er) {
    // If it was an invalid regular expression, then it can't match
    // anything.  This trick looks for a character after the end of
    // the string, which is of course impossible, except in multi-line
    // mode, but it's not a /m regex.
    return new RegExp('$.')
  }

  regExp._glob = pattern
  regExp._src = re

  return regExp
}

minimatch.makeRe = function (pattern, options) {
  return new Minimatch(pattern, options || {}).makeRe()
}

Minimatch.prototype.makeRe = makeRe
function makeRe () {
  if (this.regexp || this.regexp === false) return this.regexp

  // at this point, this.set is a 2d array of partial
  // pattern strings, or "**".
  //
  // It's better to use .match().  This function shouldn't
  // be used, really, but it's pretty convenient sometimes,
  // when you just want to work with a regex.
  var set = this.set

  if (!set.length) {
    this.regexp = false
    return this.regexp
  }
  var options = this.options

  var twoStar = options.noglobstar ? star
    : options.dot ? twoStarDot
    : twoStarNoDot
  var flags = options.nocase ? 'i' : ''

  var re = set.map(function (pattern) {
    return pattern.map(function (p) {
      return (p === GLOBSTAR) ? twoStar
      : (typeof p === 'string') ? regExpEscape(p)
      : p._src
    }).join('\\\/')
  }).join('|')

  // must match entire pattern
  // ending in a * or ** will make it less strict.
  re = '^(?:' + re + ')$'

  // can match anything, as long as it's not this.
  if (this.negate) re = '^(?!' + re + ').*$'

  try {
    this.regexp = new RegExp(re, flags)
  } catch (ex) {
    this.regexp = false
  }
  return this.regexp
}

minimatch.match = function (list, pattern, options) {
  options = options || {}
  var mm = new Minimatch(pattern, options)
  list = list.filter(function (f) {
    return mm.match(f)
  })
  if (mm.options.nonull && !list.length) {
    list.push(pattern)
  }
  return list
}

Minimatch.prototype.match = match
function match (f, partial) {
  this.debug('match', f, this.pattern)
  // short-circuit in the case of busted things.
  // comments, etc.
  if (this.comment) return false
  if (this.empty) return f === ''

  if (f === '/' && partial) return true

  var options = this.options

  // windows: need to use /, not \
  if (path.sep !== '/') {
    f = f.split(path.sep).join('/')
  }

  // treat the test path as a set of pathparts.
  f = f.split(slashSplit)
  this.debug(this.pattern, 'split', f)

  // just ONE of the pattern sets in this.set needs to match
  // in order for it to be valid.  If negating, then just one
  // match means that we have failed.
  // Either way, return on the first hit.

  var set = this.set
  this.debug(this.pattern, 'set', set)

  // Find the basename of the path by looking for the last non-empty segment
  var filename
  var i
  for (i = f.length - 1; i >= 0; i--) {
    filename = f[i]
    if (filename) break
  }

  for (i = 0; i < set.length; i++) {
    var pattern = set[i]
    var file = f
    if (options.matchBase && pattern.length === 1) {
      file = [filename]
    }
    var hit = this.matchOne(file, pattern, partial)
    if (hit) {
      if (options.flipNegate) return true
      return !this.negate
    }
  }

  // didn't get any hits.  this is success if it's a negative
  // pattern, failure otherwise.
  if (options.flipNegate) return false
  return this.negate
}

// set partial to true to test if, for example,
// "/a/b" matches the start of "/*/b/*/d"
// Partial means, if you run out of file before you run
// out of pattern, then that's fine, as long as all
// the parts match.
Minimatch.prototype.matchOne = function (file, pattern, partial) {
  var options = this.options

  this.debug('matchOne',
    { 'this': this, file: file, pattern: pattern })

  this.debug('matchOne', file.length, pattern.length)

  for (var fi = 0,
      pi = 0,
      fl = file.length,
      pl = pattern.length
      ; (fi < fl) && (pi < pl)
      ; fi++, pi++) {
    this.debug('matchOne loop')
    var p = pattern[pi]
    var f = file[fi]

    this.debug(pattern, p, f)

    // should be impossible.
    // some invalid regexp stuff in the set.
    if (p === false) return false

    if (p === GLOBSTAR) {
      this.debug('GLOBSTAR', [pattern, p, f])

      // "**"
      // a/**/b/**/c would match the following:
      // a/b/x/y/z/c
      // a/x/y/z/b/c
      // a/b/x/b/x/c
      // a/b/c
      // To do this, take the rest of the pattern after
      // the **, and see if it would match the file remainder.
      // If so, return success.
      // If not, the ** "swallows" a segment, and try again.
      // This is recursively awful.
      //
      // a/**/b/**/c matching a/b/x/y/z/c
      // - a matches a
      // - doublestar
      //   - matchOne(b/x/y/z/c, b/**/c)
      //     - b matches b
      //     - doublestar
      //       - matchOne(x/y/z/c, c) -> no
      //       - matchOne(y/z/c, c) -> no
      //       - matchOne(z/c, c) -> no
      //       - matchOne(c, c) yes, hit
      var fr = fi
      var pr = pi + 1
      if (pr === pl) {
        this.debug('** at the end')
        // a ** at the end will just swallow the rest.
        // We have found a match.
        // however, it will not swallow /.x, unless
        // options.dot is set.
        // . and .. are *never* matched by **, for explosively
        // exponential reasons.
        for (; fi < fl; fi++) {
          if (file[fi] === '.' || file[fi] === '..' ||
            (!options.dot && file[fi].charAt(0) === '.')) return false
        }
        return true
      }

      // ok, let's see if we can swallow whatever we can.
      while (fr < fl) {
        var swallowee = file[fr]

        this.debug('\nglobstar while', file, fr, pattern, pr, swallowee)

        // XXX remove this slice.  Just pass the start index.
        if (this.matchOne(file.slice(fr), pattern.slice(pr), partial)) {
          this.debug('globstar found match!', fr, fl, swallowee)
          // found a match.
          return true
        } else {
          // can't swallow "." or ".." ever.
          // can only swallow ".foo" when explicitly asked.
          if (swallowee === '.' || swallowee === '..' ||
            (!options.dot && swallowee.charAt(0) === '.')) {
            this.debug('dot detected!', file, fr, pattern, pr)
            break
          }

          // ** swallows a segment, and continue.
          this.debug('globstar swallow a segment, and continue')
          fr++
        }
      }

      // no match was found.
      // However, in partial mode, we can't say this is necessarily over.
      // If there's more *pattern* left, then
      if (partial) {
        // ran out of file
        this.debug('\n>>> no match, partial?', file, fr, pattern, pr)
        if (fr === fl) return true
      }
      return false
    }

    // something other than **
    // non-magic patterns just have to match exactly
    // patterns with magic have been turned into regexps.
    var hit
    if (typeof p === 'string') {
      if (options.nocase) {
        hit = f.toLowerCase() === p.toLowerCase()
      } else {
        hit = f === p
      }
      this.debug('string match', p, f, hit)
    } else {
      hit = f.match(p)
      this.debug('pattern match', p, f, hit)
    }

    if (!hit) return false
  }

  // Note: ending in / means that we'll get a final ""
  // at the end of the pattern.  This can only match a
  // corresponding "" at the end of the file.
  // If the file ends in /, then it can only match a
  // a pattern that ends in /, unless the pattern just
  // doesn't have any more for it. But, a/b/ should *not*
  // match "a/b/*", even though "" matches against the
  // [^/]*? pattern, except in partial mode, where it might
  // simply not be reached yet.
  // However, a/b/ should still satisfy a/*

  // now either we fell off the end of the pattern, or we're done.
  if (fi === fl && pi === pl) {
    // ran out of pattern and filename at the same time.
    // an exact hit!
    return true
  } else if (fi === fl) {
    // ran out of file, but still had pattern left.
    // this is ok if we're doing the match as part of
    // a glob fs traversal.
    return partial
  } else if (pi === pl) {
    // ran out of pattern, still have file left.
    // this is only acceptable if we're on the very last
    // empty segment of a file with a trailing slash.
    // a/* should match a/b/
    var emptyFileEnd = (fi === fl - 1) && (file[fi] === '')
    return emptyFileEnd
  }

  // should be unreachable.
  throw new Error('wtf?')
}

// replace stuff like \* with *
function globUnescape (s) {
  return s.replace(/\\(.)/g, '$1')
}

function regExpEscape (s) {
  return s.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, '\\$&')
}


/***/ }),
/* 61 */
/***/ (function(module, exports, __webpack_require__) {

var wrappy = __webpack_require__(123)
module.exports = wrappy(once)
module.exports.strict = wrappy(onceStrict)

once.proto = once(function () {
  Object.defineProperty(Function.prototype, 'once', {
    value: function () {
      return once(this)
    },
    configurable: true
  })

  Object.defineProperty(Function.prototype, 'onceStrict', {
    value: function () {
      return onceStrict(this)
    },
    configurable: true
  })
})

function once (fn) {
  var f = function () {
    if (f.called) return f.value
    f.called = true
    return f.value = fn.apply(this, arguments)
  }
  f.called = false
  return f
}

function onceStrict (fn) {
  var f = function () {
    if (f.called)
      throw new Error(f.onceError)
    f.called = true
    return f.value = fn.apply(this, arguments)
  }
  var name = fn.name || 'Function wrapped with `once`'
  f.onceError = name + " shouldn't be called more than once"
  f.called = false
  return f
}


/***/ }),
/* 62 */,
/* 63 */
/***/ (function(module, exports) {

module.exports = require("buffer");

/***/ }),
/* 64 */,
/* 65 */,
/* 66 */,
/* 67 */
/***/ (function(module, exports) {

// 7.2.1 RequireObjectCoercible(argument)
module.exports = function (it) {
  if (it == undefined) throw TypeError("Can't call method on  " + it);
  return it;
};


/***/ }),
/* 68 */
/***/ (function(module, exports, __webpack_require__) {

var isObject = __webpack_require__(34);
var document = __webpack_require__(11).document;
// typeof document.createElement is 'object' in old IE
var is = isObject(document) && isObject(document.createElement);
module.exports = function (it) {
  return is ? document.createElement(it) : {};
};


/***/ }),
/* 69 */
/***/ (function(module, exports) {

module.exports = true;


/***/ }),
/* 70 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

// 25.4.1.5 NewPromiseCapability(C)
var aFunction = __webpack_require__(46);

function PromiseCapability(C) {
  var resolve, reject;
  this.promise = new C(function ($$resolve, $$reject) {
    if (resolve !== undefined || reject !== undefined) throw TypeError('Bad Promise constructor');
    resolve = $$resolve;
    reject = $$reject;
  });
  this.resolve = aFunction(resolve);
  this.reject = aFunction(reject);
}

module.exports.f = function (C) {
  return new PromiseCapability(C);
};


/***/ }),
/* 71 */
/***/ (function(module, exports, __webpack_require__) {

var def = __webpack_require__(50).f;
var has = __webpack_require__(49);
var TAG = __webpack_require__(13)('toStringTag');

module.exports = function (it, tag, stat) {
  if (it && !has(it = stat ? it : it.prototype, TAG)) def(it, TAG, { configurable: true, value: tag });
};


/***/ }),
/* 72 */
/***/ (function(module, exports, __webpack_require__) {

var shared = __webpack_require__(107)('keys');
var uid = __webpack_require__(111);
module.exports = function (key) {
  return shared[key] || (shared[key] = uid(key));
};


/***/ }),
/* 73 */
/***/ (function(module, exports) {

// 7.1.4 ToInteger
var ceil = Math.ceil;
var floor = Math.floor;
module.exports = function (it) {
  return isNaN(it = +it) ? 0 : (it > 0 ? floor : ceil)(it);
};


/***/ }),
/* 74 */
/***/ (function(module, exports, __webpack_require__) {

// to indexed object, toObject with fallback for non-array-like ES3 strings
var IObject = __webpack_require__(131);
var defined = __webpack_require__(67);
module.exports = function (it) {
  return IObject(defined(it));
};


/***/ }),
/* 75 */
/***/ (function(module, exports, __webpack_require__) {

// Approach:
//
// 1. Get the minimatch set
// 2. For each pattern in the set, PROCESS(pattern, false)
// 3. Store matches per-set, then uniq them
//
// PROCESS(pattern, inGlobStar)
// Get the first [n] items from pattern that are all strings
// Join these together.  This is PREFIX.
//   If there is no more remaining, then stat(PREFIX) and
//   add to matches if it succeeds.  END.
//
// If inGlobStar and PREFIX is symlink and points to dir
//   set ENTRIES = []
// else readdir(PREFIX) as ENTRIES
//   If fail, END
//
// with ENTRIES
//   If pattern[n] is GLOBSTAR
//     // handle the case where the globstar match is empty
//     // by pruning it out, and testing the resulting pattern
//     PROCESS(pattern[0..n] + pattern[n+1 .. $], false)
//     // handle other cases.
//     for ENTRY in ENTRIES (not dotfiles)
//       // attach globstar + tail onto the entry
//       // Mark that this entry is a globstar match
//       PROCESS(pattern[0..n] + ENTRY + pattern[n .. $], true)
//
//   else // not globstar
//     for ENTRY in ENTRIES (not dotfiles, unless pattern[n] is dot)
//       Test ENTRY against pattern[n]
//       If fails, continue
//       If passes, PROCESS(pattern[0..n] + item + pattern[n+1 .. $])
//
// Caveat:
//   Cache all stats and readdirs results to minimize syscall.  Since all
//   we ever care about is existence and directory-ness, we can just keep
//   `true` for files, and [children,...] for directories, or `false` for
//   things that don't exist.

module.exports = glob

var fs = __webpack_require__(3)
var rp = __webpack_require__(114)
var minimatch = __webpack_require__(60)
var Minimatch = minimatch.Minimatch
var inherits = __webpack_require__(42)
var EE = __webpack_require__(54).EventEmitter
var path = __webpack_require__(0)
var assert = __webpack_require__(22)
var isAbsolute = __webpack_require__(76)
var globSync = __webpack_require__(218)
var common = __webpack_require__(115)
var alphasort = common.alphasort
var alphasorti = common.alphasorti
var setopts = common.setopts
var ownProp = common.ownProp
var inflight = __webpack_require__(223)
var util = __webpack_require__(2)
var childrenIgnored = common.childrenIgnored
var isIgnored = common.isIgnored

var once = __webpack_require__(61)

function glob (pattern, options, cb) {
  if (typeof options === 'function') cb = options, options = {}
  if (!options) options = {}

  if (options.sync) {
    if (cb)
      throw new TypeError('callback provided to sync glob')
    return globSync(pattern, options)
  }

  return new Glob(pattern, options, cb)
}

glob.sync = globSync
var GlobSync = glob.GlobSync = globSync.GlobSync

// old api surface
glob.glob = glob

function extend (origin, add) {
  if (add === null || typeof add !== 'object') {
    return origin
  }

  var keys = Object.keys(add)
  var i = keys.length
  while (i--) {
    origin[keys[i]] = add[keys[i]]
  }
  return origin
}

glob.hasMagic = function (pattern, options_) {
  var options = extend({}, options_)
  options.noprocess = true

  var g = new Glob(pattern, options)
  var set = g.minimatch.set

  if (!pattern)
    return false

  if (set.length > 1)
    return true

  for (var j = 0; j < set[0].length; j++) {
    if (typeof set[0][j] !== 'string')
      return true
  }

  return false
}

glob.Glob = Glob
inherits(Glob, EE)
function Glob (pattern, options, cb) {
  if (typeof options === 'function') {
    cb = options
    options = null
  }

  if (options && options.sync) {
    if (cb)
      throw new TypeError('callback provided to sync glob')
    return new GlobSync(pattern, options)
  }

  if (!(this instanceof Glob))
    return new Glob(pattern, options, cb)

  setopts(this, pattern, options)
  this._didRealPath = false

  // process each pattern in the minimatch set
  var n = this.minimatch.set.length

  // The matches are stored as {<filename>: true,...} so that
  // duplicates are automagically pruned.
  // Later, we do an Object.keys() on these.
  // Keep them as a list so we can fill in when nonull is set.
  this.matches = new Array(n)

  if (typeof cb === 'function') {
    cb = once(cb)
    this.on('error', cb)
    this.on('end', function (matches) {
      cb(null, matches)
    })
  }

  var self = this
  this._processing = 0

  this._emitQueue = []
  this._processQueue = []
  this.paused = false

  if (this.noprocess)
    return this

  if (n === 0)
    return done()

  var sync = true
  for (var i = 0; i < n; i ++) {
    this._process(this.minimatch.set[i], i, false, done)
  }
  sync = false

  function done () {
    --self._processing
    if (self._processing <= 0) {
      if (sync) {
        process.nextTick(function () {
          self._finish()
        })
      } else {
        self._finish()
      }
    }
  }
}

Glob.prototype._finish = function () {
  assert(this instanceof Glob)
  if (this.aborted)
    return

  if (this.realpath && !this._didRealpath)
    return this._realpath()

  common.finish(this)
  this.emit('end', this.found)
}

Glob.prototype._realpath = function () {
  if (this._didRealpath)
    return

  this._didRealpath = true

  var n = this.matches.length
  if (n === 0)
    return this._finish()

  var self = this
  for (var i = 0; i < this.matches.length; i++)
    this._realpathSet(i, next)

  function next () {
    if (--n === 0)
      self._finish()
  }
}

Glob.prototype._realpathSet = function (index, cb) {
  var matchset = this.matches[index]
  if (!matchset)
    return cb()

  var found = Object.keys(matchset)
  var self = this
  var n = found.length

  if (n === 0)
    return cb()

  var set = this.matches[index] = Object.create(null)
  found.forEach(function (p, i) {
    // If there's a problem with the stat, then it means that
    // one or more of the links in the realpath couldn't be
    // resolved.  just return the abs value in that case.
    p = self._makeAbs(p)
    rp.realpath(p, self.realpathCache, function (er, real) {
      if (!er)
        set[real] = true
      else if (er.syscall === 'stat')
        set[p] = true
      else
        self.emit('error', er) // srsly wtf right here

      if (--n === 0) {
        self.matches[index] = set
        cb()
      }
    })
  })
}

Glob.prototype._mark = function (p) {
  return common.mark(this, p)
}

Glob.prototype._makeAbs = function (f) {
  return common.makeAbs(this, f)
}

Glob.prototype.abort = function () {
  this.aborted = true
  this.emit('abort')
}

Glob.prototype.pause = function () {
  if (!this.paused) {
    this.paused = true
    this.emit('pause')
  }
}

Glob.prototype.resume = function () {
  if (this.paused) {
    this.emit('resume')
    this.paused = false
    if (this._emitQueue.length) {
      var eq = this._emitQueue.slice(0)
      this._emitQueue.length = 0
      for (var i = 0; i < eq.length; i ++) {
        var e = eq[i]
        this._emitMatch(e[0], e[1])
      }
    }
    if (this._processQueue.length) {
      var pq = this._processQueue.slice(0)
      this._processQueue.length = 0
      for (var i = 0; i < pq.length; i ++) {
        var p = pq[i]
        this._processing--
        this._process(p[0], p[1], p[2], p[3])
      }
    }
  }
}

Glob.prototype._process = function (pattern, index, inGlobStar, cb) {
  assert(this instanceof Glob)
  assert(typeof cb === 'function')

  if (this.aborted)
    return

  this._processing++
  if (this.paused) {
    this._processQueue.push([pattern, index, inGlobStar, cb])
    return
  }

  //console.error('PROCESS %d', this._processing, pattern)

  // Get the first [n] parts of pattern that are all strings.
  var n = 0
  while (typeof pattern[n] === 'string') {
    n ++
  }
  // now n is the index of the first one that is *not* a string.

  // see if there's anything else
  var prefix
  switch (n) {
    // if not, then this is rather simple
    case pattern.length:
      this._processSimple(pattern.join('/'), index, cb)
      return

    case 0:
      // pattern *starts* with some non-trivial item.
      // going to readdir(cwd), but not include the prefix in matches.
      prefix = null
      break

    default:
      // pattern has some string bits in the front.
      // whatever it starts with, whether that's 'absolute' like /foo/bar,
      // or 'relative' like '../baz'
      prefix = pattern.slice(0, n).join('/')
      break
  }

  var remain = pattern.slice(n)

  // get the list of entries.
  var read
  if (prefix === null)
    read = '.'
  else if (isAbsolute(prefix) || isAbsolute(pattern.join('/'))) {
    if (!prefix || !isAbsolute(prefix))
      prefix = '/' + prefix
    read = prefix
  } else
    read = prefix

  var abs = this._makeAbs(read)

  //if ignored, skip _processing
  if (childrenIgnored(this, read))
    return cb()

  var isGlobStar = remain[0] === minimatch.GLOBSTAR
  if (isGlobStar)
    this._processGlobStar(prefix, read, abs, remain, index, inGlobStar, cb)
  else
    this._processReaddir(prefix, read, abs, remain, index, inGlobStar, cb)
}

Glob.prototype._processReaddir = function (prefix, read, abs, remain, index, inGlobStar, cb) {
  var self = this
  this._readdir(abs, inGlobStar, function (er, entries) {
    return self._processReaddir2(prefix, read, abs, remain, index, inGlobStar, entries, cb)
  })
}

Glob.prototype._processReaddir2 = function (prefix, read, abs, remain, index, inGlobStar, entries, cb) {

  // if the abs isn't a dir, then nothing can match!
  if (!entries)
    return cb()

  // It will only match dot entries if it starts with a dot, or if
  // dot is set.  Stuff like @(.foo|.bar) isn't allowed.
  var pn = remain[0]
  var negate = !!this.minimatch.negate
  var rawGlob = pn._glob
  var dotOk = this.dot || rawGlob.charAt(0) === '.'

  var matchedEntries = []
  for (var i = 0; i < entries.length; i++) {
    var e = entries[i]
    if (e.charAt(0) !== '.' || dotOk) {
      var m
      if (negate && !prefix) {
        m = !e.match(pn)
      } else {
        m = e.match(pn)
      }
      if (m)
        matchedEntries.push(e)
    }
  }

  //console.error('prd2', prefix, entries, remain[0]._glob, matchedEntries)

  var len = matchedEntries.length
  // If there are no matched entries, then nothing matches.
  if (len === 0)
    return cb()

  // if this is the last remaining pattern bit, then no need for
  // an additional stat *unless* the user has specified mark or
  // stat explicitly.  We know they exist, since readdir returned
  // them.

  if (remain.length === 1 && !this.mark && !this.stat) {
    if (!this.matches[index])
      this.matches[index] = Object.create(null)

    for (var i = 0; i < len; i ++) {
      var e = matchedEntries[i]
      if (prefix) {
        if (prefix !== '/')
          e = prefix + '/' + e
        else
          e = prefix + e
      }

      if (e.charAt(0) === '/' && !this.nomount) {
        e = path.join(this.root, e)
      }
      this._emitMatch(index, e)
    }
    // This was the last one, and no stats were needed
    return cb()
  }

  // now test all matched entries as stand-ins for that part
  // of the pattern.
  remain.shift()
  for (var i = 0; i < len; i ++) {
    var e = matchedEntries[i]
    var newPattern
    if (prefix) {
      if (prefix !== '/')
        e = prefix + '/' + e
      else
        e = prefix + e
    }
    this._process([e].concat(remain), index, inGlobStar, cb)
  }
  cb()
}

Glob.prototype._emitMatch = function (index, e) {
  if (this.aborted)
    return

  if (isIgnored(this, e))
    return

  if (this.paused) {
    this._emitQueue.push([index, e])
    return
  }

  var abs = isAbsolute(e) ? e : this._makeAbs(e)

  if (this.mark)
    e = this._mark(e)

  if (this.absolute)
    e = abs

  if (this.matches[index][e])
    return

  if (this.nodir) {
    var c = this.cache[abs]
    if (c === 'DIR' || Array.isArray(c))
      return
  }

  this.matches[index][e] = true

  var st = this.statCache[abs]
  if (st)
    this.emit('stat', e, st)

  this.emit('match', e)
}

Glob.prototype._readdirInGlobStar = function (abs, cb) {
  if (this.aborted)
    return

  // follow all symlinked directories forever
  // just proceed as if this is a non-globstar situation
  if (this.follow)
    return this._readdir(abs, false, cb)

  var lstatkey = 'lstat\0' + abs
  var self = this
  var lstatcb = inflight(lstatkey, lstatcb_)

  if (lstatcb)
    fs.lstat(abs, lstatcb)

  function lstatcb_ (er, lstat) {
    if (er && er.code === 'ENOENT')
      return cb()

    var isSym = lstat && lstat.isSymbolicLink()
    self.symlinks[abs] = isSym

    // If it's not a symlink or a dir, then it's definitely a regular file.
    // don't bother doing a readdir in that case.
    if (!isSym && lstat && !lstat.isDirectory()) {
      self.cache[abs] = 'FILE'
      cb()
    } else
      self._readdir(abs, false, cb)
  }
}

Glob.prototype._readdir = function (abs, inGlobStar, cb) {
  if (this.aborted)
    return

  cb = inflight('readdir\0'+abs+'\0'+inGlobStar, cb)
  if (!cb)
    return

  //console.error('RD %j %j', +inGlobStar, abs)
  if (inGlobStar && !ownProp(this.symlinks, abs))
    return this._readdirInGlobStar(abs, cb)

  if (ownProp(this.cache, abs)) {
    var c = this.cache[abs]
    if (!c || c === 'FILE')
      return cb()

    if (Array.isArray(c))
      return cb(null, c)
  }

  var self = this
  fs.readdir(abs, readdirCb(this, abs, cb))
}

function readdirCb (self, abs, cb) {
  return function (er, entries) {
    if (er)
      self._readdirError(abs, er, cb)
    else
      self._readdirEntries(abs, entries, cb)
  }
}

Glob.prototype._readdirEntries = function (abs, entries, cb) {
  if (this.aborted)
    return

  // if we haven't asked to stat everything, then just
  // assume that everything in there exists, so we can avoid
  // having to stat it a second time.
  if (!this.mark && !this.stat) {
    for (var i = 0; i < entries.length; i ++) {
      var e = entries[i]
      if (abs === '/')
        e = abs + e
      else
        e = abs + '/' + e
      this.cache[e] = true
    }
  }

  this.cache[abs] = entries
  return cb(null, entries)
}

Glob.prototype._readdirError = function (f, er, cb) {
  if (this.aborted)
    return

  // handle errors, and cache the information
  switch (er.code) {
    case 'ENOTSUP': // https://github.com/isaacs/node-glob/issues/205
    case 'ENOTDIR': // totally normal. means it *does* exist.
      var abs = this._makeAbs(f)
      this.cache[abs] = 'FILE'
      if (abs === this.cwdAbs) {
        var error = new Error(er.code + ' invalid cwd ' + this.cwd)
        error.path = this.cwd
        error.code = er.code
        this.emit('error', error)
        this.abort()
      }
      break

    case 'ENOENT': // not terribly unusual
    case 'ELOOP':
    case 'ENAMETOOLONG':
    case 'UNKNOWN':
      this.cache[this._makeAbs(f)] = false
      break

    default: // some unusual error.  Treat as failure.
      this.cache[this._makeAbs(f)] = false
      if (this.strict) {
        this.emit('error', er)
        // If the error is handled, then we abort
        // if not, we threw out of here
        this.abort()
      }
      if (!this.silent)
        console.error('glob error', er)
      break
  }

  return cb()
}

Glob.prototype._processGlobStar = function (prefix, read, abs, remain, index, inGlobStar, cb) {
  var self = this
  this._readdir(abs, inGlobStar, function (er, entries) {
    self._processGlobStar2(prefix, read, abs, remain, index, inGlobStar, entries, cb)
  })
}


Glob.prototype._processGlobStar2 = function (prefix, read, abs, remain, index, inGlobStar, entries, cb) {
  //console.error('pgs2', prefix, remain[0], entries)

  // no entries means not a dir, so it can never have matches
  // foo.txt/** doesn't match foo.txt
  if (!entries)
    return cb()

  // test without the globstar, and with every child both below
  // and replacing the globstar.
  var remainWithoutGlobStar = remain.slice(1)
  var gspref = prefix ? [ prefix ] : []
  var noGlobStar = gspref.concat(remainWithoutGlobStar)

  // the noGlobStar pattern exits the inGlobStar state
  this._process(noGlobStar, index, false, cb)

  var isSym = this.symlinks[abs]
  var len = entries.length

  // If it's a symlink, and we're in a globstar, then stop
  if (isSym && inGlobStar)
    return cb()

  for (var i = 0; i < len; i++) {
    var e = entries[i]
    if (e.charAt(0) === '.' && !this.dot)
      continue

    // these two cases enter the inGlobStar state
    var instead = gspref.concat(entries[i], remainWithoutGlobStar)
    this._process(instead, index, true, cb)

    var below = gspref.concat(entries[i], remain)
    this._process(below, index, true, cb)
  }

  cb()
}

Glob.prototype._processSimple = function (prefix, index, cb) {
  // XXX review this.  Shouldn't it be doing the mounting etc
  // before doing stat?  kinda weird?
  var self = this
  this._stat(prefix, function (er, exists) {
    self._processSimple2(prefix, index, er, exists, cb)
  })
}
Glob.prototype._processSimple2 = function (prefix, index, er, exists, cb) {

  //console.error('ps2', prefix, exists)

  if (!this.matches[index])
    this.matches[index] = Object.create(null)

  // If it doesn't exist, then just mark the lack of results
  if (!exists)
    return cb()

  if (prefix && isAbsolute(prefix) && !this.nomount) {
    var trail = /[\/\\]$/.test(prefix)
    if (prefix.charAt(0) === '/') {
      prefix = path.join(this.root, prefix)
    } else {
      prefix = path.resolve(this.root, prefix)
      if (trail)
        prefix += '/'
    }
  }

  if (process.platform === 'win32')
    prefix = prefix.replace(/\\/g, '/')

  // Mark this as a match
  this._emitMatch(index, prefix)
  cb()
}

// Returns either 'DIR', 'FILE', or false
Glob.prototype._stat = function (f, cb) {
  var abs = this._makeAbs(f)
  var needDir = f.slice(-1) === '/'

  if (f.length > this.maxLength)
    return cb()

  if (!this.stat && ownProp(this.cache, abs)) {
    var c = this.cache[abs]

    if (Array.isArray(c))
      c = 'DIR'

    // It exists, but maybe not how we need it
    if (!needDir || c === 'DIR')
      return cb(null, c)

    if (needDir && c === 'FILE')
      return cb()

    // otherwise we have to stat, because maybe c=true
    // if we know it exists, but not what it is.
  }

  var exists
  var stat = this.statCache[abs]
  if (stat !== undefined) {
    if (stat === false)
      return cb(null, stat)
    else {
      var type = stat.isDirectory() ? 'DIR' : 'FILE'
      if (needDir && type === 'FILE')
        return cb()
      else
        return cb(null, type, stat)
    }
  }

  var self = this
  var statcb = inflight('stat\0' + abs, lstatcb_)
  if (statcb)
    fs.lstat(abs, statcb)

  function lstatcb_ (er, lstat) {
    if (lstat && lstat.isSymbolicLink()) {
      // If it's a symlink, then treat it as the target, unless
      // the target does not exist, then treat it as a file.
      return fs.stat(abs, function (er, stat) {
        if (er)
          self._stat2(f, abs, null, lstat, cb)
        else
          self._stat2(f, abs, er, stat, cb)
      })
    } else {
      self._stat2(f, abs, er, lstat, cb)
    }
  }
}

Glob.prototype._stat2 = function (f, abs, er, stat, cb) {
  if (er && (er.code === 'ENOENT' || er.code === 'ENOTDIR')) {
    this.statCache[abs] = false
    return cb()
  }

  var needDir = f.slice(-1) === '/'
  this.statCache[abs] = stat

  if (abs.slice(-1) === '/' && stat && !stat.isDirectory())
    return cb(null, false, stat)

  var c = true
  if (stat)
    c = stat.isDirectory() ? 'DIR' : 'FILE'
  this.cache[abs] = this.cache[abs] || c

  if (needDir && c === 'FILE')
    return cb()

  return cb(null, c, stat)
}


/***/ }),
/* 76 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


function posix(path) {
	return path.charAt(0) === '/';
}

function win32(path) {
	// https://github.com/nodejs/node/blob/b3fcc245fb25539909ef1d5eaa01dbf92e168633/lib/path.js#L56
	var splitDeviceRe = /^([a-zA-Z]:|[\\\/]{2}[^\\\/]+[\\\/]+[^\\\/]+)?([\\\/])?([\s\S]*?)$/;
	var result = splitDeviceRe.exec(path);
	var device = result[1] || '';
	var isUnc = Boolean(device && device.charAt(1) !== ':');

	// UNC paths are always absolute
	return Boolean(result[2] || isUnc);
}

module.exports = process.platform === 'win32' ? win32 : posix;
module.exports.posix = posix;
module.exports.win32 = win32;


/***/ }),
/* 77 */,
/* 78 */,
/* 79 */
/***/ (function(module, exports) {

module.exports = require("tty");

/***/ }),
/* 80 */,
/* 81 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});

exports.default = function (str, fileLoc = 'lockfile') {
  str = (0, (_stripBom || _load_stripBom()).default)(str);
  return hasMergeConflicts(str) ? parseWithConflict(str, fileLoc) : { type: 'success', object: parse(str, fileLoc) };
};

var _util;

function _load_util() {
  return _util = _interopRequireDefault(__webpack_require__(2));
}

var _invariant;

function _load_invariant() {
  return _invariant = _interopRequireDefault(__webpack_require__(7));
}

var _stripBom;

function _load_stripBom() {
  return _stripBom = _interopRequireDefault(__webpack_require__(122));
}

var _constants;

function _load_constants() {
  return _constants = __webpack_require__(6);
}

var _errors;

function _load_errors() {
  return _errors = __webpack_require__(4);
}

var _map;

function _load_map() {
  return _map = _interopRequireDefault(__webpack_require__(20));
}

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

/* eslint quotes: 0 */

const VERSION_REGEX = /^yarn lockfile v(\d+)$/;

const TOKEN_TYPES = {
  boolean: 'BOOLEAN',
  string: 'STRING',
  identifier: 'IDENTIFIER',
  eof: 'EOF',
  colon: 'COLON',
  newline: 'NEWLINE',
  comment: 'COMMENT',
  indent: 'INDENT',
  invalid: 'INVALID',
  number: 'NUMBER',
  comma: 'COMMA'
};

const VALID_PROP_VALUE_TOKENS = [TOKEN_TYPES.boolean, TOKEN_TYPES.string, TOKEN_TYPES.number];

function isValidPropValueToken(token) {
  return VALID_PROP_VALUE_TOKENS.indexOf(token.type) >= 0;
}

function* tokenise(input) {
  let lastNewline = false;
  let line = 1;
  let col = 0;

  function buildToken(type, value) {
    return { line, col, type, value };
  }

  while (input.length) {
    let chop = 0;

    if (input[0] === '\n' || input[0] === '\r') {
      chop++;
      // If this is a \r\n line, ignore both chars but only add one new line
      if (input[1] === '\n') {
        chop++;
      }
      line++;
      col = 0;
      yield buildToken(TOKEN_TYPES.newline);
    } else if (input[0] === '#') {
      chop++;

      let val = '';
      while (input[chop] !== '\n') {
        val += input[chop];
        chop++;
      }
      yield buildToken(TOKEN_TYPES.comment, val);
    } else if (input[0] === ' ') {
      if (lastNewline) {
        let indent = '';
        for (let i = 0; input[i] === ' '; i++) {
          indent += input[i];
        }

        if (indent.length % 2) {
          throw new TypeError('Invalid number of spaces');
        } else {
          chop = indent.length;
          yield buildToken(TOKEN_TYPES.indent, indent.length / 2);
        }
      } else {
        chop++;
      }
    } else if (input[0] === '"') {
      let val = '';

      for (let i = 0;; i++) {
        const currentChar = input[i];
        val += currentChar;

        if (i > 0 && currentChar === '"') {
          const isEscaped = input[i - 1] === '\\' && input[i - 2] !== '\\';
          if (!isEscaped) {
            break;
          }
        }
      }

      chop = val.length;

      try {
        yield buildToken(TOKEN_TYPES.string, JSON.parse(val));
      } catch (err) {
        if (err instanceof SyntaxError) {
          yield buildToken(TOKEN_TYPES.invalid);
        } else {
          throw err;
        }
      }
    } else if (/^[0-9]/.test(input)) {
      let val = '';
      for (let i = 0; /^[0-9]$/.test(input[i]); i++) {
        val += input[i];
      }
      chop = val.length;

      yield buildToken(TOKEN_TYPES.number, +val);
    } else if (/^true/.test(input)) {
      yield buildToken(TOKEN_TYPES.boolean, true);
      chop = 4;
    } else if (/^false/.test(input)) {
      yield buildToken(TOKEN_TYPES.boolean, false);
      chop = 5;
    } else if (input[0] === ':') {
      yield buildToken(TOKEN_TYPES.colon);
      chop++;
    } else if (input[0] === ',') {
      yield buildToken(TOKEN_TYPES.comma);
      chop++;
    } else if (/^[a-zA-Z\/-]/g.test(input)) {
      let name = '';
      for (let i = 0; i < input.length; i++) {
        const char = input[i];
        if (char === ':' || char === ' ' || char === '\n' || char === '\r' || char === ',') {
          break;
        } else {
          name += char;
        }
      }
      chop = name.length;

      yield buildToken(TOKEN_TYPES.string, name);
    } else {
      yield buildToken(TOKEN_TYPES.invalid);
    }

    if (!chop) {
      // will trigger infinite recursion
      yield buildToken(TOKEN_TYPES.invalid);
    }

    col += chop;
    lastNewline = input[0] === '\n' || input[0] === '\r' && input[1] === '\n';
    input = input.slice(chop);
  }

  yield buildToken(TOKEN_TYPES.eof);
}

class Parser {
  constructor(input, fileLoc = 'lockfile') {
    this.comments = [];
    this.tokens = tokenise(input);
    this.fileLoc = fileLoc;
  }

  onComment(token) {
    const value = token.value;
    (0, (_invariant || _load_invariant()).default)(typeof value === 'string', 'expected token value to be a string');

    const comment = value.trim();

    const versionMatch = comment.match(VERSION_REGEX);
    if (versionMatch) {
      const version = +versionMatch[1];
      if (version > (_constants || _load_constants()).LOCKFILE_VERSION) {
        throw new (_errors || _load_errors()).MessageError(`Can't install from a lockfile of version ${version} as you're on an old yarn version that only supports ` + `versions up to ${(_constants || _load_constants()).LOCKFILE_VERSION}. Run \`$ yarn self-update\` to upgrade to the latest version.`);
      }
    }

    this.comments.push(comment);
  }

  next() {
    const item = this.tokens.next();
    (0, (_invariant || _load_invariant()).default)(item, 'expected a token');

    const done = item.done,
          value = item.value;

    if (done || !value) {
      throw new Error('No more tokens');
    } else if (value.type === TOKEN_TYPES.comment) {
      this.onComment(value);
      return this.next();
    } else {
      return this.token = value;
    }
  }

  unexpected(msg = 'Unexpected token') {
    throw new SyntaxError(`${msg} ${this.token.line}:${this.token.col} in ${this.fileLoc}`);
  }

  expect(tokType) {
    if (this.token.type === tokType) {
      this.next();
    } else {
      this.unexpected();
    }
  }

  eat(tokType) {
    if (this.token.type === tokType) {
      this.next();
      return true;
    } else {
      return false;
    }
  }

  parse(indent = 0) {
    const obj = (0, (_map || _load_map()).default)();

    while (true) {
      const propToken = this.token;

      if (propToken.type === TOKEN_TYPES.newline) {
        const nextToken = this.next();
        if (!indent) {
          // if we have 0 indentation then the next token doesn't matter
          continue;
        }

        if (nextToken.type !== TOKEN_TYPES.indent) {
          // if we have no indentation after a newline then we've gone down a level
          break;
        }

        if (nextToken.value === indent) {
          // all is good, the indent is on our level
          this.next();
        } else {
          // the indentation is less than our level
          break;
        }
      } else if (propToken.type === TOKEN_TYPES.indent) {
        if (propToken.value === indent) {
          this.next();
        } else {
          break;
        }
      } else if (propToken.type === TOKEN_TYPES.eof) {
        break;
      } else if (propToken.type === TOKEN_TYPES.string) {
        // property key
        const key = propToken.value;
        (0, (_invariant || _load_invariant()).default)(key, 'Expected a key');

        const keys = [key];
        this.next();

        // support multiple keys
        while (this.token.type === TOKEN_TYPES.comma) {
          this.next(); // skip comma

          const keyToken = this.token;
          if (keyToken.type !== TOKEN_TYPES.string) {
            this.unexpected('Expected string');
          }

          const key = keyToken.value;
          (0, (_invariant || _load_invariant()).default)(key, 'Expected a key');
          keys.push(key);
          this.next();
        }

        const valToken = this.token;

        if (valToken.type === TOKEN_TYPES.colon) {
          // object
          this.next();

          // parse object
          const val = this.parse(indent + 1);

          for (var _iterator = keys, _isArray = Array.isArray(_iterator), _i = 0, _iterator = _isArray ? _iterator : _iterator[Symbol.iterator]();;) {
            var _ref;

            if (_isArray) {
              if (_i >= _iterator.length) break;
              _ref = _iterator[_i++];
            } else {
              _i = _iterator.next();
              if (_i.done) break;
              _ref = _i.value;
            }

            const key = _ref;

            obj[key] = val;
          }

          if (indent && this.token.type !== TOKEN_TYPES.indent) {
            break;
          }
        } else if (isValidPropValueToken(valToken)) {
          // plain value
          for (var _iterator2 = keys, _isArray2 = Array.isArray(_iterator2), _i2 = 0, _iterator2 = _isArray2 ? _iterator2 : _iterator2[Symbol.iterator]();;) {
            var _ref2;

            if (_isArray2) {
              if (_i2 >= _iterator2.length) break;
              _ref2 = _iterator2[_i2++];
            } else {
              _i2 = _iterator2.next();
              if (_i2.done) break;
              _ref2 = _i2.value;
            }

            const key = _ref2;

            obj[key] = valToken.value;
          }

          this.next();
        } else {
          this.unexpected('Invalid value type');
        }
      } else {
        this.unexpected(`Unknown token: ${(_util || _load_util()).default.inspect(propToken)}`);
      }
    }

    return obj;
  }
}

const MERGE_CONFLICT_ANCESTOR = '|||||||';
const MERGE_CONFLICT_END = '>>>>>>>';
const MERGE_CONFLICT_SEP = '=======';
const MERGE_CONFLICT_START = '<<<<<<<';

/**
 * Extract the two versions of the lockfile from a merge conflict.
 */
function extractConflictVariants(str) {
  const variants = [[], []];
  const lines = str.split(/\r?\n/g);
  let skip = false;

  while (lines.length) {
    const line = lines.shift();
    if (line.startsWith(MERGE_CONFLICT_START)) {
      // get the first variant
      while (lines.length) {
        const conflictLine = lines.shift();
        if (conflictLine === MERGE_CONFLICT_SEP) {
          skip = false;
          break;
        } else if (skip || conflictLine.startsWith(MERGE_CONFLICT_ANCESTOR)) {
          skip = true;
          continue;
        } else {
          variants[0].push(conflictLine);
        }
      }

      // get the second variant
      while (lines.length) {
        const conflictLine = lines.shift();
        if (conflictLine.startsWith(MERGE_CONFLICT_END)) {
          break;
        } else {
          variants[1].push(conflictLine);
        }
      }
    } else {
      variants[0].push(line);
      variants[1].push(line);
    }
  }

  return [variants[0].join('\n'), variants[1].join('\n')];
}

/**
 * Check if a lockfile has merge conflicts.
 */
function hasMergeConflicts(str) {
  return str.includes(MERGE_CONFLICT_START) && str.includes(MERGE_CONFLICT_SEP) && str.includes(MERGE_CONFLICT_END);
}

/**
 * Parse the lockfile.
 */
function parse(str, fileLoc) {
  const parser = new Parser(str, fileLoc);
  parser.next();
  return parser.parse();
}

/**
 * Parse and merge the two variants in a conflicted lockfile.
 */
function parseWithConflict(str, fileLoc) {
  const variants = extractConflictVariants(str);
  try {
    return { type: 'merge', object: Object.assign({}, parse(variants[0], fileLoc), parse(variants[1], fileLoc)) };
  } catch (err) {
    if (err instanceof SyntaxError) {
      return { type: 'conflict', object: {} };
    } else {
      throw err;
    }
  }
}

/***/ }),
/* 82 */,
/* 83 */,
/* 84 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});

var _map;

function _load_map() {
  return _map = _interopRequireDefault(__webpack_require__(20));
}

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

const debug = __webpack_require__(212)('yarn');

class BlockingQueue {
  constructor(alias, maxConcurrency = Infinity) {
    this.concurrencyQueue = [];
    this.maxConcurrency = maxConcurrency;
    this.runningCount = 0;
    this.warnedStuck = false;
    this.alias = alias;
    this.first = true;

    this.running = (0, (_map || _load_map()).default)();
    this.queue = (0, (_map || _load_map()).default)();

    this.stuckTick = this.stuckTick.bind(this);
  }

  stillActive() {
    if (this.stuckTimer) {
      clearTimeout(this.stuckTimer);
    }

    this.stuckTimer = setTimeout(this.stuckTick, 5000);

    // We need to check the existence of unref because of https://github.com/facebook/jest/issues/4559
    // $FlowFixMe: Node's setInterval returns a Timeout, not a Number
    this.stuckTimer.unref && this.stuckTimer.unref();
  }

  stuckTick() {
    if (this.runningCount === 1) {
      this.warnedStuck = true;
      debug(`The ${JSON.stringify(this.alias)} blocking queue may be stuck. 5 seconds ` + `without any activity with 1 worker: ${Object.keys(this.running)[0]}`);
    }
  }

  push(key, factory) {
    if (this.first) {
      this.first = false;
    } else {
      this.stillActive();
    }

    return new Promise((resolve, reject) => {
      // we're already running so push ourselves to the queue
      const queue = this.queue[key] = this.queue[key] || [];
      queue.push({ factory, resolve, reject });

      if (!this.running[key]) {
        this.shift(key);
      }
    });
  }

  shift(key) {
    if (this.running[key]) {
      delete this.running[key];
      this.runningCount--;

      if (this.stuckTimer) {
        clearTimeout(this.stuckTimer);
        this.stuckTimer = null;
      }

      if (this.warnedStuck) {
        this.warnedStuck = false;
        debug(`${JSON.stringify(this.alias)} blocking queue finally resolved. Nothing to worry about.`);
      }
    }

    const queue = this.queue[key];
    if (!queue) {
      return;
    }

    var _queue$shift = queue.shift();

    const resolve = _queue$shift.resolve,
          reject = _queue$shift.reject,
          factory = _queue$shift.factory;

    if (!queue.length) {
      delete this.queue[key];
    }

    const next = () => {
      this.shift(key);
      this.shiftConcurrencyQueue();
    };

    const run = () => {
      this.running[key] = true;
      this.runningCount++;

      factory().then(function (val) {
        resolve(val);
        next();
        return null;
      }).catch(function (err) {
        reject(err);
        next();
      });
    };

    this.maybePushConcurrencyQueue(run);
  }

  maybePushConcurrencyQueue(run) {
    if (this.runningCount < this.maxConcurrency) {
      run();
    } else {
      this.concurrencyQueue.push(run);
    }
  }

  shiftConcurrencyQueue() {
    if (this.runningCount < this.maxConcurrency) {
      const fn = this.concurrencyQueue.shift();
      if (fn) {
        fn();
      }
    }
  }
}
exports.default = BlockingQueue;

/***/ }),
/* 85 */
/***/ (function(module, exports) {

module.exports = function (exec) {
  try {
    return !!exec();
  } catch (e) {
    return true;
  }
};


/***/ }),
/* 86 */,
/* 87 */,
/* 88 */,
/* 89 */,
/* 90 */,
/* 91 */,
/* 92 */,
/* 93 */,
/* 94 */,
/* 95 */,
/* 96 */,
/* 97 */,
/* 98 */,
/* 99 */,
/* 100 */
/***/ (function(module, exports, __webpack_require__) {

// getting tag from 19.1.3.6 Object.prototype.toString()
var cof = __webpack_require__(47);
var TAG = __webpack_require__(13)('toStringTag');
// ES3 wrong here
var ARG = cof(function () { return arguments; }()) == 'Arguments';

// fallback for IE11 Script Access Denied error
var tryGet = function (it, key) {
  try {
    return it[key];
  } catch (e) { /* empty */ }
};

module.exports = function (it) {
  var O, T, B;
  return it === undefined ? 'Undefined' : it === null ? 'Null'
    // @@toStringTag case
    : typeof (T = tryGet(O = Object(it), TAG)) == 'string' ? T
    // builtinTag case
    : ARG ? cof(O)
    // ES3 arguments fallback
    : (B = cof(O)) == 'Object' && typeof O.callee == 'function' ? 'Arguments' : B;
};


/***/ }),
/* 101 */
/***/ (function(module, exports) {

// IE 8- don't enum bug keys
module.exports = (
  'constructor,hasOwnProperty,isPrototypeOf,propertyIsEnumerable,toLocaleString,toString,valueOf'
).split(',');


/***/ }),
/* 102 */
/***/ (function(module, exports, __webpack_require__) {

var document = __webpack_require__(11).document;
module.exports = document && document.documentElement;


/***/ }),
/* 103 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var LIBRARY = __webpack_require__(69);
var $export = __webpack_require__(41);
var redefine = __webpack_require__(197);
var hide = __webpack_require__(31);
var Iterators = __webpack_require__(35);
var $iterCreate = __webpack_require__(188);
var setToStringTag = __webpack_require__(71);
var getPrototypeOf = __webpack_require__(194);
var ITERATOR = __webpack_require__(13)('iterator');
var BUGGY = !([].keys && 'next' in [].keys()); // Safari has buggy iterators w/o `next`
var FF_ITERATOR = '@@iterator';
var KEYS = 'keys';
var VALUES = 'values';

var returnThis = function () { return this; };

module.exports = function (Base, NAME, Constructor, next, DEFAULT, IS_SET, FORCED) {
  $iterCreate(Constructor, NAME, next);
  var getMethod = function (kind) {
    if (!BUGGY && kind in proto) return proto[kind];
    switch (kind) {
      case KEYS: return function keys() { return new Constructor(this, kind); };
      case VALUES: return function values() { return new Constructor(this, kind); };
    } return function entries() { return new Constructor(this, kind); };
  };
  var TAG = NAME + ' Iterator';
  var DEF_VALUES = DEFAULT == VALUES;
  var VALUES_BUG = false;
  var proto = Base.prototype;
  var $native = proto[ITERATOR] || proto[FF_ITERATOR] || DEFAULT && proto[DEFAULT];
  var $default = $native || getMethod(DEFAULT);
  var $entries = DEFAULT ? !DEF_VALUES ? $default : getMethod('entries') : undefined;
  var $anyNative = NAME == 'Array' ? proto.entries || $native : $native;
  var methods, key, IteratorPrototype;
  // Fix native
  if ($anyNative) {
    IteratorPrototype = getPrototypeOf($anyNative.call(new Base()));
    if (IteratorPrototype !== Object.prototype && IteratorPrototype.next) {
      // Set @@toStringTag to native iterators
      setToStringTag(IteratorPrototype, TAG, true);
      // fix for some old engines
      if (!LIBRARY && typeof IteratorPrototype[ITERATOR] != 'function') hide(IteratorPrototype, ITERATOR, returnThis);
    }
  }
  // fix Array#{values, @@iterator}.name in V8 / FF
  if (DEF_VALUES && $native && $native.name !== VALUES) {
    VALUES_BUG = true;
    $default = function values() { return $native.call(this); };
  }
  // Define iterator
  if ((!LIBRARY || FORCED) && (BUGGY || VALUES_BUG || !proto[ITERATOR])) {
    hide(proto, ITERATOR, $default);
  }
  // Plug for library
  Iterators[NAME] = $default;
  Iterators[TAG] = returnThis;
  if (DEFAULT) {
    methods = {
      values: DEF_VALUES ? $default : getMethod(VALUES),
      keys: IS_SET ? $default : getMethod(KEYS),
      entries: $entries
    };
    if (FORCED) for (key in methods) {
      if (!(key in proto)) redefine(proto, key, methods[key]);
    } else $export($export.P + $export.F * (BUGGY || VALUES_BUG), NAME, methods);
  }
  return methods;
};


/***/ }),
/* 104 */
/***/ (function(module, exports) {

module.exports = function (exec) {
  try {
    return { e: false, v: exec() };
  } catch (e) {
    return { e: true, v: e };
  }
};


/***/ }),
/* 105 */
/***/ (function(module, exports, __webpack_require__) {

var anObject = __webpack_require__(27);
var isObject = __webpack_require__(34);
var newPromiseCapability = __webpack_require__(70);

module.exports = function (C, x) {
  anObject(C);
  if (isObject(x) && x.constructor === C) return x;
  var promiseCapability = newPromiseCapability.f(C);
  var resolve = promiseCapability.resolve;
  resolve(x);
  return promiseCapability.promise;
};


/***/ }),
/* 106 */
/***/ (function(module, exports) {

module.exports = function (bitmap, value) {
  return {
    enumerable: !(bitmap & 1),
    configurable: !(bitmap & 2),
    writable: !(bitmap & 4),
    value: value
  };
};


/***/ }),
/* 107 */
/***/ (function(module, exports, __webpack_require__) {

var core = __webpack_require__(23);
var global = __webpack_require__(11);
var SHARED = '__core-js_shared__';
var store = global[SHARED] || (global[SHARED] = {});

(module.exports = function (key, value) {
  return store[key] || (store[key] = value !== undefined ? value : {});
})('versions', []).push({
  version: core.version,
  mode: __webpack_require__(69) ? 'pure' : 'global',
  copyright: '© 2018 Denis Pushkarev (zloirock.ru)'
});


/***/ }),
/* 108 */
/***/ (function(module, exports, __webpack_require__) {

// 7.3.20 SpeciesConstructor(O, defaultConstructor)
var anObject = __webpack_require__(27);
var aFunction = __webpack_require__(46);
var SPECIES = __webpack_require__(13)('species');
module.exports = function (O, D) {
  var C = anObject(O).constructor;
  var S;
  return C === undefined || (S = anObject(C)[SPECIES]) == undefined ? D : aFunction(S);
};


/***/ }),
/* 109 */
/***/ (function(module, exports, __webpack_require__) {

var ctx = __webpack_require__(48);
var invoke = __webpack_require__(185);
var html = __webpack_require__(102);
var cel = __webpack_require__(68);
var global = __webpack_require__(11);
var process = global.process;
var setTask = global.setImmediate;
var clearTask = global.clearImmediate;
var MessageChannel = global.MessageChannel;
var Dispatch = global.Dispatch;
var counter = 0;
var queue = {};
var ONREADYSTATECHANGE = 'onreadystatechange';
var defer, channel, port;
var run = function () {
  var id = +this;
  // eslint-disable-next-line no-prototype-builtins
  if (queue.hasOwnProperty(id)) {
    var fn = queue[id];
    delete queue[id];
    fn();
  }
};
var listener = function (event) {
  run.call(event.data);
};
// Node.js 0.9+ & IE10+ has setImmediate, otherwise:
if (!setTask || !clearTask) {
  setTask = function setImmediate(fn) {
    var args = [];
    var i = 1;
    while (arguments.length > i) args.push(arguments[i++]);
    queue[++counter] = function () {
      // eslint-disable-next-line no-new-func
      invoke(typeof fn == 'function' ? fn : Function(fn), args);
    };
    defer(counter);
    return counter;
  };
  clearTask = function clearImmediate(id) {
    delete queue[id];
  };
  // Node.js 0.8-
  if (__webpack_require__(47)(process) == 'process') {
    defer = function (id) {
      process.nextTick(ctx(run, id, 1));
    };
  // Sphere (JS game engine) Dispatch API
  } else if (Dispatch && Dispatch.now) {
    defer = function (id) {
      Dispatch.now(ctx(run, id, 1));
    };
  // Browsers with MessageChannel, includes WebWorkers
  } else if (MessageChannel) {
    channel = new MessageChannel();
    port = channel.port2;
    channel.port1.onmessage = listener;
    defer = ctx(port.postMessage, port, 1);
  // Browsers with postMessage, skip WebWorkers
  // IE8 has postMessage, but it's sync & typeof its postMessage is 'object'
  } else if (global.addEventListener && typeof postMessage == 'function' && !global.importScripts) {
    defer = function (id) {
      global.postMessage(id + '', '*');
    };
    global.addEventListener('message', listener, false);
  // IE8-
  } else if (ONREADYSTATECHANGE in cel('script')) {
    defer = function (id) {
      html.appendChild(cel('script'))[ONREADYSTATECHANGE] = function () {
        html.removeChild(this);
        run.call(id);
      };
    };
  // Rest old browsers
  } else {
    defer = function (id) {
      setTimeout(ctx(run, id, 1), 0);
    };
  }
}
module.exports = {
  set: setTask,
  clear: clearTask
};


/***/ }),
/* 110 */
/***/ (function(module, exports, __webpack_require__) {

// 7.1.15 ToLength
var toInteger = __webpack_require__(73);
var min = Math.min;
module.exports = function (it) {
  return it > 0 ? min(toInteger(it), 0x1fffffffffffff) : 0; // pow(2, 53) - 1 == 9007199254740991
};


/***/ }),
/* 111 */
/***/ (function(module, exports) {

var id = 0;
var px = Math.random();
module.exports = function (key) {
  return 'Symbol('.concat(key === undefined ? '' : key, ')_', (++id + px).toString(36));
};


/***/ }),
/* 112 */
/***/ (function(module, exports, __webpack_require__) {


/**
 * This is the common logic for both the Node.js and web browser
 * implementations of `debug()`.
 *
 * Expose `debug()` as the module.
 */

exports = module.exports = createDebug.debug = createDebug['default'] = createDebug;
exports.coerce = coerce;
exports.disable = disable;
exports.enable = enable;
exports.enabled = enabled;
exports.humanize = __webpack_require__(229);

/**
 * Active `debug` instances.
 */
exports.instances = [];

/**
 * The currently active debug mode names, and names to skip.
 */

exports.names = [];
exports.skips = [];

/**
 * Map of special "%n" handling functions, for the debug "format" argument.
 *
 * Valid key names are a single, lower or upper-case letter, i.e. "n" and "N".
 */

exports.formatters = {};

/**
 * Select a color.
 * @param {String} namespace
 * @return {Number}
 * @api private
 */

function selectColor(namespace) {
  var hash = 0, i;

  for (i in namespace) {
    hash  = ((hash << 5) - hash) + namespace.charCodeAt(i);
    hash |= 0; // Convert to 32bit integer
  }

  return exports.colors[Math.abs(hash) % exports.colors.length];
}

/**
 * Create a debugger with the given `namespace`.
 *
 * @param {String} namespace
 * @return {Function}
 * @api public
 */

function createDebug(namespace) {

  var prevTime;

  function debug() {
    // disabled?
    if (!debug.enabled) return;

    var self = debug;

    // set `diff` timestamp
    var curr = +new Date();
    var ms = curr - (prevTime || curr);
    self.diff = ms;
    self.prev = prevTime;
    self.curr = curr;
    prevTime = curr;

    // turn the `arguments` into a proper Array
    var args = new Array(arguments.length);
    for (var i = 0; i < args.length; i++) {
      args[i] = arguments[i];
    }

    args[0] = exports.coerce(args[0]);

    if ('string' !== typeof args[0]) {
      // anything else let's inspect with %O
      args.unshift('%O');
    }

    // apply any `formatters` transformations
    var index = 0;
    args[0] = args[0].replace(/%([a-zA-Z%])/g, function(match, format) {
      // if we encounter an escaped % then don't increase the array index
      if (match === '%%') return match;
      index++;
      var formatter = exports.formatters[format];
      if ('function' === typeof formatter) {
        var val = args[index];
        match = formatter.call(self, val);

        // now we need to remove `args[index]` since it's inlined in the `format`
        args.splice(index, 1);
        index--;
      }
      return match;
    });

    // apply env-specific formatting (colors, etc.)
    exports.formatArgs.call(self, args);

    var logFn = debug.log || exports.log || console.log.bind(console);
    logFn.apply(self, args);
  }

  debug.namespace = namespace;
  debug.enabled = exports.enabled(namespace);
  debug.useColors = exports.useColors();
  debug.color = selectColor(namespace);
  debug.destroy = destroy;

  // env-specific initialization logic for debug instances
  if ('function' === typeof exports.init) {
    exports.init(debug);
  }

  exports.instances.push(debug);

  return debug;
}

function destroy () {
  var index = exports.instances.indexOf(this);
  if (index !== -1) {
    exports.instances.splice(index, 1);
    return true;
  } else {
    return false;
  }
}

/**
 * Enables a debug mode by namespaces. This can include modes
 * separated by a colon and wildcards.
 *
 * @param {String} namespaces
 * @api public
 */

function enable(namespaces) {
  exports.save(namespaces);

  exports.names = [];
  exports.skips = [];

  var i;
  var split = (typeof namespaces === 'string' ? namespaces : '').split(/[\s,]+/);
  var len = split.length;

  for (i = 0; i < len; i++) {
    if (!split[i]) continue; // ignore empty strings
    namespaces = split[i].replace(/\*/g, '.*?');
    if (namespaces[0] === '-') {
      exports.skips.push(new RegExp('^' + namespaces.substr(1) + '$'));
    } else {
      exports.names.push(new RegExp('^' + namespaces + '$'));
    }
  }

  for (i = 0; i < exports.instances.length; i++) {
    var instance = exports.instances[i];
    instance.enabled = exports.enabled(instance.namespace);
  }
}

/**
 * Disable debug output.
 *
 * @api public
 */

function disable() {
  exports.enable('');
}

/**
 * Returns true if the given mode name is enabled, false otherwise.
 *
 * @param {String} name
 * @return {Boolean}
 * @api public
 */

function enabled(name) {
  if (name[name.length - 1] === '*') {
    return true;
  }
  var i, len;
  for (i = 0, len = exports.skips.length; i < len; i++) {
    if (exports.skips[i].test(name)) {
      return false;
    }
  }
  for (i = 0, len = exports.names.length; i < len; i++) {
    if (exports.names[i].test(name)) {
      return true;
    }
  }
  return false;
}

/**
 * Coerce `val`.
 *
 * @param {Mixed} val
 * @return {Mixed}
 * @api private
 */

function coerce(val) {
  if (val instanceof Error) return val.stack || val.message;
  return val;
}


/***/ }),
/* 113 */,
/* 114 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = realpath
realpath.realpath = realpath
realpath.sync = realpathSync
realpath.realpathSync = realpathSync
realpath.monkeypatch = monkeypatch
realpath.unmonkeypatch = unmonkeypatch

var fs = __webpack_require__(3)
var origRealpath = fs.realpath
var origRealpathSync = fs.realpathSync

var version = process.version
var ok = /^v[0-5]\./.test(version)
var old = __webpack_require__(217)

function newError (er) {
  return er && er.syscall === 'realpath' && (
    er.code === 'ELOOP' ||
    er.code === 'ENOMEM' ||
    er.code === 'ENAMETOOLONG'
  )
}

function realpath (p, cache, cb) {
  if (ok) {
    return origRealpath(p, cache, cb)
  }

  if (typeof cache === 'function') {
    cb = cache
    cache = null
  }
  origRealpath(p, cache, function (er, result) {
    if (newError(er)) {
      old.realpath(p, cache, cb)
    } else {
      cb(er, result)
    }
  })
}

function realpathSync (p, cache) {
  if (ok) {
    return origRealpathSync(p, cache)
  }

  try {
    return origRealpathSync(p, cache)
  } catch (er) {
    if (newError(er)) {
      return old.realpathSync(p, cache)
    } else {
      throw er
    }
  }
}

function monkeypatch () {
  fs.realpath = realpath
  fs.realpathSync = realpathSync
}

function unmonkeypatch () {
  fs.realpath = origRealpath
  fs.realpathSync = origRealpathSync
}


/***/ }),
/* 115 */
/***/ (function(module, exports, __webpack_require__) {

exports.alphasort = alphasort
exports.alphasorti = alphasorti
exports.setopts = setopts
exports.ownProp = ownProp
exports.makeAbs = makeAbs
exports.finish = finish
exports.mark = mark
exports.isIgnored = isIgnored
exports.childrenIgnored = childrenIgnored

function ownProp (obj, field) {
  return Object.prototype.hasOwnProperty.call(obj, field)
}

var path = __webpack_require__(0)
var minimatch = __webpack_require__(60)
var isAbsolute = __webpack_require__(76)
var Minimatch = minimatch.Minimatch

function alphasorti (a, b) {
  return a.toLowerCase().localeCompare(b.toLowerCase())
}

function alphasort (a, b) {
  return a.localeCompare(b)
}

function setupIgnores (self, options) {
  self.ignore = options.ignore || []

  if (!Array.isArray(self.ignore))
    self.ignore = [self.ignore]

  if (self.ignore.length) {
    self.ignore = self.ignore.map(ignoreMap)
  }
}

// ignore patterns are always in dot:true mode.
function ignoreMap (pattern) {
  var gmatcher = null
  if (pattern.slice(-3) === '/**') {
    var gpattern = pattern.replace(/(\/\*\*)+$/, '')
    gmatcher = new Minimatch(gpattern, { dot: true })
  }

  return {
    matcher: new Minimatch(pattern, { dot: true }),
    gmatcher: gmatcher
  }
}

function setopts (self, pattern, options) {
  if (!options)
    options = {}

  // base-matching: just use globstar for that.
  if (options.matchBase && -1 === pattern.indexOf("/")) {
    if (options.noglobstar) {
      throw new Error("base matching requires globstar")
    }
    pattern = "**/" + pattern
  }

  self.silent = !!options.silent
  self.pattern = pattern
  self.strict = options.strict !== false
  self.realpath = !!options.realpath
  self.realpathCache = options.realpathCache || Object.create(null)
  self.follow = !!options.follow
  self.dot = !!options.dot
  self.mark = !!options.mark
  self.nodir = !!options.nodir
  if (self.nodir)
    self.mark = true
  self.sync = !!options.sync
  self.nounique = !!options.nounique
  self.nonull = !!options.nonull
  self.nosort = !!options.nosort
  self.nocase = !!options.nocase
  self.stat = !!options.stat
  self.noprocess = !!options.noprocess
  self.absolute = !!options.absolute

  self.maxLength = options.maxLength || Infinity
  self.cache = options.cache || Object.create(null)
  self.statCache = options.statCache || Object.create(null)
  self.symlinks = options.symlinks || Object.create(null)

  setupIgnores(self, options)

  self.changedCwd = false
  var cwd = process.cwd()
  if (!ownProp(options, "cwd"))
    self.cwd = cwd
  else {
    self.cwd = path.resolve(options.cwd)
    self.changedCwd = self.cwd !== cwd
  }

  self.root = options.root || path.resolve(self.cwd, "/")
  self.root = path.resolve(self.root)
  if (process.platform === "win32")
    self.root = self.root.replace(/\\/g, "/")

  // TODO: is an absolute `cwd` supposed to be resolved against `root`?
  // e.g. { cwd: '/test', root: __dirname } === path.join(__dirname, '/test')
  self.cwdAbs = isAbsolute(self.cwd) ? self.cwd : makeAbs(self, self.cwd)
  if (process.platform === "win32")
    self.cwdAbs = self.cwdAbs.replace(/\\/g, "/")
  self.nomount = !!options.nomount

  // disable comments and negation in Minimatch.
  // Note that they are not supported in Glob itself anyway.
  options.nonegate = true
  options.nocomment = true

  self.minimatch = new Minimatch(pattern, options)
  self.options = self.minimatch.options
}

function finish (self) {
  var nou = self.nounique
  var all = nou ? [] : Object.create(null)

  for (var i = 0, l = self.matches.length; i < l; i ++) {
    var matches = self.matches[i]
    if (!matches || Object.keys(matches).length === 0) {
      if (self.nonull) {
        // do like the shell, and spit out the literal glob
        var literal = self.minimatch.globSet[i]
        if (nou)
          all.push(literal)
        else
          all[literal] = true
      }
    } else {
      // had matches
      var m = Object.keys(matches)
      if (nou)
        all.push.apply(all, m)
      else
        m.forEach(function (m) {
          all[m] = true
        })
    }
  }

  if (!nou)
    all = Object.keys(all)

  if (!self.nosort)
    all = all.sort(self.nocase ? alphasorti : alphasort)

  // at *some* point we statted all of these
  if (self.mark) {
    for (var i = 0; i < all.length; i++) {
      all[i] = self._mark(all[i])
    }
    if (self.nodir) {
      all = all.filter(function (e) {
        var notDir = !(/\/$/.test(e))
        var c = self.cache[e] || self.cache[makeAbs(self, e)]
        if (notDir && c)
          notDir = c !== 'DIR' && !Array.isArray(c)
        return notDir
      })
    }
  }

  if (self.ignore.length)
    all = all.filter(function(m) {
      return !isIgnored(self, m)
    })

  self.found = all
}

function mark (self, p) {
  var abs = makeAbs(self, p)
  var c = self.cache[abs]
  var m = p
  if (c) {
    var isDir = c === 'DIR' || Array.isArray(c)
    var slash = p.slice(-1) === '/'

    if (isDir && !slash)
      m += '/'
    else if (!isDir && slash)
      m = m.slice(0, -1)

    if (m !== p) {
      var mabs = makeAbs(self, m)
      self.statCache[mabs] = self.statCache[abs]
      self.cache[mabs] = self.cache[abs]
    }
  }

  return m
}

// lotta situps...
function makeAbs (self, f) {
  var abs = f
  if (f.charAt(0) === '/') {
    abs = path.join(self.root, f)
  } else if (isAbsolute(f) || f === '') {
    abs = f
  } else if (self.changedCwd) {
    abs = path.resolve(self.cwd, f)
  } else {
    abs = path.resolve(f)
  }

  if (process.platform === 'win32')
    abs = abs.replace(/\\/g, '/')

  return abs
}


// Return true, if pattern ends with globstar '**', for the accompanying parent directory.
// Ex:- If node_modules/** is the pattern, add 'node_modules' to ignore list along with it's contents
function isIgnored (self, path) {
  if (!self.ignore.length)
    return false

  return self.ignore.some(function(item) {
    return item.matcher.match(path) || !!(item.gmatcher && item.gmatcher.match(path))
  })
}

function childrenIgnored (self, path) {
  if (!self.ignore.length)
    return false

  return self.ignore.some(function(item) {
    return !!(item.gmatcher && item.gmatcher.match(path))
  })
}


/***/ }),
/* 116 */
/***/ (function(module, exports, __webpack_require__) {

var path = __webpack_require__(0);
var fs = __webpack_require__(3);
var _0777 = parseInt('0777', 8);

module.exports = mkdirP.mkdirp = mkdirP.mkdirP = mkdirP;

function mkdirP (p, opts, f, made) {
    if (typeof opts === 'function') {
        f = opts;
        opts = {};
    }
    else if (!opts || typeof opts !== 'object') {
        opts = { mode: opts };
    }
    
    var mode = opts.mode;
    var xfs = opts.fs || fs;
    
    if (mode === undefined) {
        mode = _0777 & (~process.umask());
    }
    if (!made) made = null;
    
    var cb = f || function () {};
    p = path.resolve(p);
    
    xfs.mkdir(p, mode, function (er) {
        if (!er) {
            made = made || p;
            return cb(null, made);
        }
        switch (er.code) {
            case 'ENOENT':
                mkdirP(path.dirname(p), opts, function (er, made) {
                    if (er) cb(er, made);
                    else mkdirP(p, opts, cb, made);
                });
                break;

            // In the case of any other error, just see if there's a dir
            // there already.  If so, then hooray!  If not, then something
            // is borked.
            default:
                xfs.stat(p, function (er2, stat) {
                    // if the stat fails, then that's super weird.
                    // let the original error be the failure reason.
                    if (er2 || !stat.isDirectory()) cb(er, made)
                    else cb(null, made);
                });
                break;
        }
    });
}

mkdirP.sync = function sync (p, opts, made) {
    if (!opts || typeof opts !== 'object') {
        opts = { mode: opts };
    }
    
    var mode = opts.mode;
    var xfs = opts.fs || fs;
    
    if (mode === undefined) {
        mode = _0777 & (~process.umask());
    }
    if (!made) made = null;

    p = path.resolve(p);

    try {
        xfs.mkdirSync(p, mode);
        made = made || p;
    }
    catch (err0) {
        switch (err0.code) {
            case 'ENOENT' :
                made = sync(path.dirname(p), opts, made);
                sync(p, opts, made);
                break;

            // In the case of any other error, just see if there's a dir
            // there already.  If so, then hooray!  If not, then something
            // is borked.
            default:
                var stat;
                try {
                    stat = xfs.statSync(p);
                }
                catch (err1) {
                    throw err0;
                }
                if (!stat.isDirectory()) throw err0;
                break;
        }
    }

    return made;
};


/***/ }),
/* 117 */,
/* 118 */,
/* 119 */,
/* 120 */,
/* 121 */,
/* 122 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

module.exports = x => {
	if (typeof x !== 'string') {
		throw new TypeError('Expected a string, got ' + typeof x);
	}

	// Catches EFBBBF (UTF-8 BOM) because the buffer-to-string
	// conversion translates it to FEFF (UTF-16 BOM)
	if (x.charCodeAt(0) === 0xFEFF) {
		return x.slice(1);
	}

	return x;
};


/***/ }),
/* 123 */
/***/ (function(module, exports) {

// Returns a wrapper function that returns a wrapped callback
// The wrapper function should do some stuff, and return a
// presumably different callback function.
// This makes sure that own properties are retained, so that
// decorations and such are not lost along the way.
module.exports = wrappy
function wrappy (fn, cb) {
  if (fn && cb) return wrappy(fn)(cb)

  if (typeof fn !== 'function')
    throw new TypeError('need wrapper function')

  Object.keys(fn).forEach(function (k) {
    wrapper[k] = fn[k]
  })

  return wrapper

  function wrapper() {
    var args = new Array(arguments.length)
    for (var i = 0; i < args.length; i++) {
      args[i] = arguments[i]
    }
    var ret = fn.apply(this, args)
    var cb = args[args.length-1]
    if (typeof ret === 'function' && ret !== cb) {
      Object.keys(cb).forEach(function (k) {
        ret[k] = cb[k]
      })
    }
    return ret
  }
}


/***/ }),
/* 124 */,
/* 125 */,
/* 126 */,
/* 127 */,
/* 128 */,
/* 129 */,
/* 130 */,
/* 131 */
/***/ (function(module, exports, __webpack_require__) {

// fallback for non-array-like ES3 and non-enumerable old V8 strings
var cof = __webpack_require__(47);
// eslint-disable-next-line no-prototype-builtins
module.exports = Object('z').propertyIsEnumerable(0) ? Object : function (it) {
  return cof(it) == 'String' ? it.split('') : Object(it);
};


/***/ }),
/* 132 */
/***/ (function(module, exports, __webpack_require__) {

// 19.1.2.14 / 15.2.3.14 Object.keys(O)
var $keys = __webpack_require__(195);
var enumBugKeys = __webpack_require__(101);

module.exports = Object.keys || function keys(O) {
  return $keys(O, enumBugKeys);
};


/***/ }),
/* 133 */
/***/ (function(module, exports, __webpack_require__) {

// 7.1.13 ToObject(argument)
var defined = __webpack_require__(67);
module.exports = function (it) {
  return Object(defined(it));
};


/***/ }),
/* 134 */,
/* 135 */,
/* 136 */,
/* 137 */,
/* 138 */,
/* 139 */,
/* 140 */,
/* 141 */,
/* 142 */,
/* 143 */,
/* 144 */,
/* 145 */
/***/ (function(module, exports) {

module.exports = {"name":"yarn","installationMethod":"unknown","version":"1.10.0-0","license":"BSD-2-Clause","preferGlobal":true,"description":"📦🐈 Fast, reliable, and secure dependency management.","dependencies":{"@zkochan/cmd-shim":"^2.2.4","babel-runtime":"^6.26.0","bytes":"^3.0.0","camelcase":"^4.0.0","chalk":"^2.1.0","commander":"^2.9.0","death":"^1.0.0","debug":"^3.0.0","deep-equal":"^1.0.1","detect-indent":"^5.0.0","dnscache":"^1.0.1","glob":"^7.1.1","gunzip-maybe":"^1.4.0","hash-for-dep":"^1.2.3","imports-loader":"^0.8.0","ini":"^1.3.4","inquirer":"^3.0.1","invariant":"^2.2.0","is-builtin-module":"^2.0.0","is-ci":"^1.0.10","is-webpack-bundle":"^1.0.0","leven":"^2.0.0","loud-rejection":"^1.2.0","micromatch":"^2.3.11","mkdirp":"^0.5.1","node-emoji":"^1.6.1","normalize-url":"^2.0.0","npm-logical-tree":"^1.2.1","object-path":"^0.11.2","proper-lockfile":"^2.0.0","puka":"^1.0.0","read":"^1.0.7","request":"^2.87.0","request-capture-har":"^1.2.2","rimraf":"^2.5.0","semver":"^5.1.0","ssri":"^5.3.0","strip-ansi":"^4.0.0","strip-bom":"^3.0.0","tar-fs":"^1.16.0","tar-stream":"^1.6.1","uuid":"^3.0.1","v8-compile-cache":"^2.0.0","validate-npm-package-license":"^3.0.3","yn":"^2.0.0"},"devDependencies":{"babel-core":"^6.26.0","babel-eslint":"^7.2.3","babel-loader":"^6.2.5","babel-plugin-array-includes":"^2.0.3","babel-plugin-transform-builtin-extend":"^1.1.2","babel-plugin-transform-inline-imports-commonjs":"^1.0.0","babel-plugin-transform-runtime":"^6.4.3","babel-preset-env":"^1.6.0","babel-preset-flow":"^6.23.0","babel-preset-stage-0":"^6.0.0","babylon":"^6.5.0","commitizen":"^2.9.6","cz-conventional-changelog":"^2.0.0","eslint":"^4.3.0","eslint-config-fb-strict":"^22.0.0","eslint-plugin-babel":"^5.0.0","eslint-plugin-flowtype":"^2.35.0","eslint-plugin-jasmine":"^2.6.2","eslint-plugin-jest":"^21.0.0","eslint-plugin-jsx-a11y":"^6.0.2","eslint-plugin-prefer-object-spread":"^1.2.1","eslint-plugin-prettier":"^2.1.2","eslint-plugin-react":"^7.1.0","eslint-plugin-relay":"^0.0.24","eslint-plugin-yarn-internal":"file:scripts/eslint-rules","execa":"^0.10.0","flow-bin":"^0.66.0","git-release-notes":"^3.0.0","gulp":"^3.9.0","gulp-babel":"^7.0.0","gulp-if":"^2.0.1","gulp-newer":"^1.0.0","gulp-plumber":"^1.0.1","gulp-sourcemaps":"^2.2.0","gulp-util":"^3.0.7","gulp-watch":"^5.0.0","jest":"^22.4.4","jsinspect":"^0.12.6","minimatch":"^3.0.4","mock-stdin":"^0.3.0","prettier":"^1.5.2","temp":"^0.8.3","webpack":"^2.1.0-beta.25","yargs":"^6.3.0"},"resolutions":{"sshpk":"^1.14.2"},"engines":{"node":">=4.0.0"},"repository":"yarnpkg/yarn","bin":{"yarn":"./bin/yarn.js","yarnpkg":"./bin/yarn.js"},"scripts":{"build":"gulp build","build-bundle":"node ./scripts/build-webpack.js","build-chocolatey":"powershell ./scripts/build-chocolatey.ps1","build-deb":"./scripts/build-deb.sh","build-dist":"bash ./scripts/build-dist.sh","build-win-installer":"scripts\\build-windows-installer.bat","changelog":"git-release-notes $(git describe --tags --abbrev=0 $(git describe --tags --abbrev=0)^)..$(git describe --tags --abbrev=0) scripts/changelog.md","dupe-check":"yarn jsinspect ./src","lint":"eslint . && flow check","pkg-tests":"yarn --cwd packages/pkg-tests jest yarn.test.js","prettier":"eslint src __tests__ --fix","release-branch":"./scripts/release-branch.sh","test":"yarn lint && yarn test-only","test-only":"node --max_old_space_size=4096 node_modules/jest/bin/jest.js --verbose","test-only-debug":"node --inspect-brk --max_old_space_size=4096 node_modules/jest/bin/jest.js --runInBand --verbose","test-coverage":"node --max_old_space_size=4096 node_modules/jest/bin/jest.js --coverage --verbose","watch":"gulp watch","commit":"git-cz"},"jest":{"collectCoverageFrom":["src/**/*.js"],"testEnvironment":"node","modulePathIgnorePatterns":["__tests__/fixtures/","packages/pkg-tests/pkg-tests-fixtures","dist/"],"testPathIgnorePatterns":["__tests__/(fixtures|__mocks__)/","updates/","_(temp|mock|install|init|helpers).js$","packages/pkg-tests"]},"config":{"commitizen":{"path":"./node_modules/cz-conventional-changelog"}}}

/***/ }),
/* 146 */,
/* 147 */,
/* 148 */,
/* 149 */,
/* 150 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = stringify;

var _misc;

function _load_misc() {
  return _misc = __webpack_require__(12);
}

var _constants;

function _load_constants() {
  return _constants = __webpack_require__(6);
}

var _package;

function _load_package() {
  return _package = __webpack_require__(145);
}

const NODE_VERSION = process.version;

function shouldWrapKey(str) {
  return str.indexOf('true') === 0 || str.indexOf('false') === 0 || /[:\s\n\\",\[\]]/g.test(str) || /^[0-9]/g.test(str) || !/^[a-zA-Z]/g.test(str);
}

function maybeWrap(str) {
  if (typeof str === 'boolean' || typeof str === 'number' || shouldWrapKey(str)) {
    return JSON.stringify(str);
  } else {
    return str;
  }
}

const priorities = {
  name: 1,
  version: 2,
  uid: 3,
  resolved: 4,
  integrity: 5,
  registry: 6,
  dependencies: 7
};

function priorityThenAlphaSort(a, b) {
  if (priorities[a] || priorities[b]) {
    return (priorities[a] || 100) > (priorities[b] || 100) ? 1 : -1;
  } else {
    return (0, (_misc || _load_misc()).sortAlpha)(a, b);
  }
}

function _stringify(obj, options) {
  if (typeof obj !== 'object') {
    throw new TypeError();
  }

  const indent = options.indent;
  const lines = [];

  // Sorting order needs to be consistent between runs, we run native sort by name because there are no
  // problems with it being unstable because there are no to keys the same
  // However priorities can be duplicated and native sort can shuffle things from run to run
  const keys = Object.keys(obj).sort(priorityThenAlphaSort);

  let addedKeys = [];

  for (let i = 0; i < keys.length; i++) {
    const key = keys[i];
    const val = obj[key];
    if (val == null || addedKeys.indexOf(key) >= 0) {
      continue;
    }

    const valKeys = [key];

    // get all keys that have the same value equality, we only want this for objects
    if (typeof val === 'object') {
      for (let j = i + 1; j < keys.length; j++) {
        const key = keys[j];
        if (val === obj[key]) {
          valKeys.push(key);
        }
      }
    }

    const keyLine = valKeys.sort((_misc || _load_misc()).sortAlpha).map(maybeWrap).join(', ');

    if (typeof val === 'string' || typeof val === 'boolean' || typeof val === 'number') {
      lines.push(`${keyLine} ${maybeWrap(val)}`);
    } else if (typeof val === 'object') {
      lines.push(`${keyLine}:\n${_stringify(val, { indent: indent + '  ' })}` + (options.topLevel ? '\n' : ''));
    } else {
      throw new TypeError();
    }

    addedKeys = addedKeys.concat(valKeys);
  }

  return indent + lines.join(`\n${indent}`);
}

function stringify(obj, noHeader, enableVersions) {
  const val = _stringify(obj, {
    indent: '',
    topLevel: true
  });
  if (noHeader) {
    return val;
  }

  const lines = [];
  lines.push('# THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.');
  lines.push(`# yarn lockfile v${(_constants || _load_constants()).LOCKFILE_VERSION}`);
  if (enableVersions) {
    lines.push(`# yarn v${(_package || _load_package()).version}`);
    lines.push(`# node ${NODE_VERSION}`);
  }
  lines.push('\n');
  lines.push(val);

  return lines.join('\n');
}

/***/ }),
/* 151 */,
/* 152 */,
/* 153 */,
/* 154 */,
/* 155 */,
/* 156 */,
/* 157 */,
/* 158 */,
/* 159 */,
/* 160 */,
/* 161 */,
/* 162 */,
/* 163 */,
/* 164 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.fileDatesEqual = exports.copyFile = exports.unlink = undefined;

var _asyncToGenerator2;

function _load_asyncToGenerator() {
  return _asyncToGenerator2 = _interopRequireDefault(__webpack_require__(1));
}

// We want to preserve file timestamps when copying a file, since yarn uses them to decide if a file has
// changed compared to the cache.
// There are some OS specific cases here:
// * On linux, fs.copyFile does not preserve timestamps, but does on OSX and Win.
// * On windows, you must open a file with write permissions to call `fs.futimes`.
// * On OSX you can open with read permissions and still call `fs.futimes`.
let fixTimes = (() => {
  var _ref3 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (fd, dest, data) {
    const doOpen = fd === undefined;
    let openfd = fd ? fd : -1;

    if (disableTimestampCorrection === undefined) {
      // if timestamps match already, no correction is needed.
      // the need to correct timestamps varies based on OS and node versions.
      const destStat = yield lstat(dest);
      disableTimestampCorrection = fileDatesEqual(destStat.mtime, data.mtime);
    }

    if (disableTimestampCorrection) {
      return;
    }

    if (doOpen) {
      try {
        openfd = yield open(dest, 'a', data.mode);
      } catch (er) {
        // file is likely read-only
        try {
          openfd = yield open(dest, 'r', data.mode);
        } catch (err) {
          // We can't even open this file for reading.
          return;
        }
      }
    }

    try {
      if (openfd) {
        yield futimes(openfd, data.atime, data.mtime);
      }
    } catch (er) {
      // If `futimes` throws an exception, we probably have a case of a read-only file on Windows.
      // In this case we can just return. The incorrect timestamp will just cause that file to be recopied
      // on subsequent installs, which will effect yarn performance but not break anything.
    } finally {
      if (doOpen && openfd) {
        yield close(openfd);
      }
    }
  });

  return function fixTimes(_x7, _x8, _x9) {
    return _ref3.apply(this, arguments);
  };
})();

// Compare file timestamps.
// Some versions of Node on windows zero the milliseconds when utime is used.


var _fs;

function _load_fs() {
  return _fs = _interopRequireDefault(__webpack_require__(3));
}

var _promise;

function _load_promise() {
  return _promise = __webpack_require__(40);
}

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

// This module serves as a wrapper for file operations that are inconsistant across node and OS versions.

let disableTimestampCorrection = undefined; // OS dependent. will be detected on first file copy.

const readFileBuffer = (0, (_promise || _load_promise()).promisify)((_fs || _load_fs()).default.readFile);
const close = (0, (_promise || _load_promise()).promisify)((_fs || _load_fs()).default.close);
const lstat = (0, (_promise || _load_promise()).promisify)((_fs || _load_fs()).default.lstat);
const open = (0, (_promise || _load_promise()).promisify)((_fs || _load_fs()).default.open);
const futimes = (0, (_promise || _load_promise()).promisify)((_fs || _load_fs()).default.futimes);

const write = (0, (_promise || _load_promise()).promisify)((_fs || _load_fs()).default.write);

const unlink = exports.unlink = (0, (_promise || _load_promise()).promisify)(__webpack_require__(233));

/**
 * Unlinks the destination to force a recreation. This is needed on case-insensitive file systems
 * to force the correct naming when the filename has changed only in character-casing. (Jest -> jest).
 */
const copyFile = exports.copyFile = (() => {
  var _ref = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (data, cleanup) {
    try {
      yield unlink(data.dest);
      yield copyFilePoly(data.src, data.dest, 0, data);
    } finally {
      if (cleanup) {
        cleanup();
      }
    }
  });

  return function copyFile(_x, _x2) {
    return _ref.apply(this, arguments);
  };
})();

// Node 8.5.0 introduced `fs.copyFile` which is much faster, so use that when available.
// Otherwise we fall back to reading and writing files as buffers.
const copyFilePoly = (src, dest, flags, data) => {
  if ((_fs || _load_fs()).default.copyFile) {
    return new Promise((resolve, reject) => (_fs || _load_fs()).default.copyFile(src, dest, flags, err => {
      if (err) {
        reject(err);
      } else {
        fixTimes(undefined, dest, data).then(() => resolve()).catch(ex => reject(ex));
      }
    }));
  } else {
    return copyWithBuffer(src, dest, flags, data);
  }
};

const copyWithBuffer = (() => {
  var _ref2 = (0, (_asyncToGenerator2 || _load_asyncToGenerator()).default)(function* (src, dest, flags, data) {
    // Use open -> write -> futimes -> close sequence to avoid opening the file twice:
    // one with writeFile and one with utimes
    const fd = yield open(dest, 'w', data.mode);
    try {
      const buffer = yield readFileBuffer(src);
      yield write(fd, buffer, 0, buffer.length);
      yield fixTimes(fd, dest, data);
    } finally {
      yield close(fd);
    }
  });

  return function copyWithBuffer(_x3, _x4, _x5, _x6) {
    return _ref2.apply(this, arguments);
  };
})();const fileDatesEqual = exports.fileDatesEqual = (a, b) => {
  const aTime = a.getTime();
  const bTime = b.getTime();

  if (process.platform !== 'win32') {
    return aTime === bTime;
  }

  // See https://github.com/nodejs/node/pull/12607
  // Submillisecond times from stat and utimes are truncated on Windows,
  // causing a file with mtime 8.0079998 and 8.0081144 to become 8.007 and 8.008
  // and making it impossible to update these files to their correct timestamps.
  if (Math.abs(aTime - bTime) <= 1) {
    return true;
  }

  const aTimeSec = Math.floor(aTime / 1000);
  const bTimeSec = Math.floor(bTime / 1000);

  // See https://github.com/nodejs/node/issues/2069
  // Some versions of Node on windows zero the milliseconds when utime is used
  // So if any of the time has a milliseconds part of zero we suspect that the
  // bug is present and compare only seconds.
  if (aTime - aTimeSec * 1000 === 0 || bTime - bTimeSec * 1000 === 0) {
    return aTimeSec === bTimeSec;
  }

  return aTime === bTime;
};

/***/ }),
/* 165 */,
/* 166 */,
/* 167 */,
/* 168 */,
/* 169 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.isFakeRoot = isFakeRoot;
exports.isRootUser = isRootUser;
function getUid() {
  if (process.platform !== 'win32' && process.getuid) {
    return process.getuid();
  }
  return null;
}

exports.default = isRootUser(getUid()) && !isFakeRoot();
function isFakeRoot() {
  return Boolean(process.env.FAKEROOTKEY);
}

function isRootUser(uid) {
  return uid === 0;
}

/***/ }),
/* 170 */,
/* 171 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getDataDir = getDataDir;
exports.getCacheDir = getCacheDir;
exports.getConfigDir = getConfigDir;
const path = __webpack_require__(0);
const userHome = __webpack_require__(45).default;

const FALLBACK_CONFIG_DIR = path.join(userHome, '.config', 'yarn');
const FALLBACK_CACHE_DIR = path.join(userHome, '.cache', 'yarn');

function getDataDir() {
  if (process.platform === 'win32') {
    const WIN32_APPDATA_DIR = getLocalAppDataDir();
    return WIN32_APPDATA_DIR == null ? FALLBACK_CONFIG_DIR : path.join(WIN32_APPDATA_DIR, 'Data');
  } else if (process.env.XDG_DATA_HOME) {
    return path.join(process.env.XDG_DATA_HOME, 'yarn');
  } else {
    // This could arguably be ~/Library/Application Support/Yarn on Macs,
    // but that feels unintuitive for a cli tool

    // Instead, use our prior fallback. Some day this could be
    // path.join(userHome, '.local', 'share', 'yarn')
    // or return path.join(WIN32_APPDATA_DIR, 'Data') on win32
    return FALLBACK_CONFIG_DIR;
  }
}

function getCacheDir() {
  if (process.platform === 'win32') {
    // process.env.TEMP also exists, but most apps put caches here
    return path.join(getLocalAppDataDir() || path.join(userHome, 'AppData', 'Local', 'Yarn'), 'Cache');
  } else if (process.env.XDG_CACHE_HOME) {
    return path.join(process.env.XDG_CACHE_HOME, 'yarn');
  } else if (process.platform === 'darwin') {
    return path.join(userHome, 'Library', 'Caches', 'Yarn');
  } else {
    return FALLBACK_CACHE_DIR;
  }
}

function getConfigDir() {
  if (process.platform === 'win32') {
    // Use our prior fallback. Some day this could be
    // return path.join(WIN32_APPDATA_DIR, 'Config')
    const WIN32_APPDATA_DIR = getLocalAppDataDir();
    return WIN32_APPDATA_DIR == null ? FALLBACK_CONFIG_DIR : path.join(WIN32_APPDATA_DIR, 'Config');
  } else if (process.env.XDG_CONFIG_HOME) {
    return path.join(process.env.XDG_CONFIG_HOME, 'yarn');
  } else {
    return FALLBACK_CONFIG_DIR;
  }
}

function getLocalAppDataDir() {
  return process.env.LOCALAPPDATA ? path.join(process.env.LOCALAPPDATA, 'Yarn') : null;
}

/***/ }),
/* 172 */,
/* 173 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = { "default": __webpack_require__(179), __esModule: true };

/***/ }),
/* 174 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

module.exports = balanced;
function balanced(a, b, str) {
  if (a instanceof RegExp) a = maybeMatch(a, str);
  if (b instanceof RegExp) b = maybeMatch(b, str);

  var r = range(a, b, str);

  return r && {
    start: r[0],
    end: r[1],
    pre: str.slice(0, r[0]),
    body: str.slice(r[0] + a.length, r[1]),
    post: str.slice(r[1] + b.length)
  };
}

function maybeMatch(reg, str) {
  var m = str.match(reg);
  return m ? m[0] : null;
}

balanced.range = range;
function range(a, b, str) {
  var begs, beg, left, right, result;
  var ai = str.indexOf(a);
  var bi = str.indexOf(b, ai + 1);
  var i = ai;

  if (ai >= 0 && bi > 0) {
    begs = [];
    left = str.length;

    while (i >= 0 && !result) {
      if (i == ai) {
        begs.push(i);
        ai = str.indexOf(a, i + 1);
      } else if (begs.length == 1) {
        result = [ begs.pop(), bi ];
      } else {
        beg = begs.pop();
        if (beg < left) {
          left = beg;
          right = bi;
        }

        bi = str.indexOf(b, i + 1);
      }

      i = ai < bi && ai >= 0 ? ai : bi;
    }

    if (begs.length) {
      result = [ left, right ];
    }
  }

  return result;
}


/***/ }),
/* 175 */
/***/ (function(module, exports, __webpack_require__) {

var concatMap = __webpack_require__(178);
var balanced = __webpack_require__(174);

module.exports = expandTop;

var escSlash = '\0SLASH'+Math.random()+'\0';
var escOpen = '\0OPEN'+Math.random()+'\0';
var escClose = '\0CLOSE'+Math.random()+'\0';
var escComma = '\0COMMA'+Math.random()+'\0';
var escPeriod = '\0PERIOD'+Math.random()+'\0';

function numeric(str) {
  return parseInt(str, 10) == str
    ? parseInt(str, 10)
    : str.charCodeAt(0);
}

function escapeBraces(str) {
  return str.split('\\\\').join(escSlash)
            .split('\\{').join(escOpen)
            .split('\\}').join(escClose)
            .split('\\,').join(escComma)
            .split('\\.').join(escPeriod);
}

function unescapeBraces(str) {
  return str.split(escSlash).join('\\')
            .split(escOpen).join('{')
            .split(escClose).join('}')
            .split(escComma).join(',')
            .split(escPeriod).join('.');
}


// Basically just str.split(","), but handling cases
// where we have nested braced sections, which should be
// treated as individual members, like {a,{b,c},d}
function parseCommaParts(str) {
  if (!str)
    return [''];

  var parts = [];
  var m = balanced('{', '}', str);

  if (!m)
    return str.split(',');

  var pre = m.pre;
  var body = m.body;
  var post = m.post;
  var p = pre.split(',');

  p[p.length-1] += '{' + body + '}';
  var postParts = parseCommaParts(post);
  if (post.length) {
    p[p.length-1] += postParts.shift();
    p.push.apply(p, postParts);
  }

  parts.push.apply(parts, p);

  return parts;
}

function expandTop(str) {
  if (!str)
    return [];

  // I don't know why Bash 4.3 does this, but it does.
  // Anything starting with {} will have the first two bytes preserved
  // but *only* at the top level, so {},a}b will not expand to anything,
  // but a{},b}c will be expanded to [a}c,abc].
  // One could argue that this is a bug in Bash, but since the goal of
  // this module is to match Bash's rules, we escape a leading {}
  if (str.substr(0, 2) === '{}') {
    str = '\\{\\}' + str.substr(2);
  }

  return expand(escapeBraces(str), true).map(unescapeBraces);
}

function identity(e) {
  return e;
}

function embrace(str) {
  return '{' + str + '}';
}
function isPadded(el) {
  return /^-?0\d/.test(el);
}

function lte(i, y) {
  return i <= y;
}
function gte(i, y) {
  return i >= y;
}

function expand(str, isTop) {
  var expansions = [];

  var m = balanced('{', '}', str);
  if (!m || /\$$/.test(m.pre)) return [str];

  var isNumericSequence = /^-?\d+\.\.-?\d+(?:\.\.-?\d+)?$/.test(m.body);
  var isAlphaSequence = /^[a-zA-Z]\.\.[a-zA-Z](?:\.\.-?\d+)?$/.test(m.body);
  var isSequence = isNumericSequence || isAlphaSequence;
  var isOptions = m.body.indexOf(',') >= 0;
  if (!isSequence && !isOptions) {
    // {a},b}
    if (m.post.match(/,.*\}/)) {
      str = m.pre + '{' + m.body + escClose + m.post;
      return expand(str);
    }
    return [str];
  }

  var n;
  if (isSequence) {
    n = m.body.split(/\.\./);
  } else {
    n = parseCommaParts(m.body);
    if (n.length === 1) {
      // x{{a,b}}y ==> x{a}y x{b}y
      n = expand(n[0], false).map(embrace);
      if (n.length === 1) {
        var post = m.post.length
          ? expand(m.post, false)
          : [''];
        return post.map(function(p) {
          return m.pre + n[0] + p;
        });
      }
    }
  }

  // at this point, n is the parts, and we know it's not a comma set
  // with a single entry.

  // no need to expand pre, since it is guaranteed to be free of brace-sets
  var pre = m.pre;
  var post = m.post.length
    ? expand(m.post, false)
    : [''];

  var N;

  if (isSequence) {
    var x = numeric(n[0]);
    var y = numeric(n[1]);
    var width = Math.max(n[0].length, n[1].length)
    var incr = n.length == 3
      ? Math.abs(numeric(n[2]))
      : 1;
    var test = lte;
    var reverse = y < x;
    if (reverse) {
      incr *= -1;
      test = gte;
    }
    var pad = n.some(isPadded);

    N = [];

    for (var i = x; test(i, y); i += incr) {
      var c;
      if (isAlphaSequence) {
        c = String.fromCharCode(i);
        if (c === '\\')
          c = '';
      } else {
        c = String(i);
        if (pad) {
          var need = width - c.length;
          if (need > 0) {
            var z = new Array(need + 1).join('0');
            if (i < 0)
              c = '-' + z + c.slice(1);
            else
              c = z + c;
          }
        }
      }
      N.push(c);
    }
  } else {
    N = concatMap(n, function(el) { return expand(el, false) });
  }

  for (var j = 0; j < N.length; j++) {
    for (var k = 0; k < post.length; k++) {
      var expansion = pre + N[j] + post[k];
      if (!isTop || isSequence || expansion)
        expansions.push(expansion);
    }
  }

  return expansions;
}



/***/ }),
/* 176 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


function preserveCamelCase(str) {
	let isLastCharLower = false;
	let isLastCharUpper = false;
	let isLastLastCharUpper = false;

	for (let i = 0; i < str.length; i++) {
		const c = str[i];

		if (isLastCharLower && /[a-zA-Z]/.test(c) && c.toUpperCase() === c) {
			str = str.substr(0, i) + '-' + str.substr(i);
			isLastCharLower = false;
			isLastLastCharUpper = isLastCharUpper;
			isLastCharUpper = true;
			i++;
		} else if (isLastCharUpper && isLastLastCharUpper && /[a-zA-Z]/.test(c) && c.toLowerCase() === c) {
			str = str.substr(0, i - 1) + '-' + str.substr(i - 1);
			isLastLastCharUpper = isLastCharUpper;
			isLastCharUpper = false;
			isLastCharLower = true;
		} else {
			isLastCharLower = c.toLowerCase() === c;
			isLastLastCharUpper = isLastCharUpper;
			isLastCharUpper = c.toUpperCase() === c;
		}
	}

	return str;
}

module.exports = function (str) {
	if (arguments.length > 1) {
		str = Array.from(arguments)
			.map(x => x.trim())
			.filter(x => x.length)
			.join('-');
	} else {
		str = str.trim();
	}

	if (str.length === 0) {
		return '';
	}

	if (str.length === 1) {
		return str.toLowerCase();
	}

	if (/^[a-z0-9]+$/.test(str)) {
		return str;
	}

	const hasUpperCase = str !== str.toLowerCase();

	if (hasUpperCase) {
		str = preserveCamelCase(str);
	}

	return str
		.replace(/^[_.\- ]+/, '')
		.toLowerCase()
		.replace(/[_.\- ]+(\w|$)/g, (m, p1) => p1.toUpperCase());
};


/***/ }),
/* 177 */,
/* 178 */
/***/ (function(module, exports) {

module.exports = function (xs, fn) {
    var res = [];
    for (var i = 0; i < xs.length; i++) {
        var x = fn(xs[i], i);
        if (isArray(x)) res.push.apply(res, x);
        else res.push(x);
    }
    return res;
};

var isArray = Array.isArray || function (xs) {
    return Object.prototype.toString.call(xs) === '[object Array]';
};


/***/ }),
/* 179 */
/***/ (function(module, exports, __webpack_require__) {

__webpack_require__(205);
__webpack_require__(207);
__webpack_require__(210);
__webpack_require__(206);
__webpack_require__(208);
__webpack_require__(209);
module.exports = __webpack_require__(23).Promise;


/***/ }),
/* 180 */
/***/ (function(module, exports) {

module.exports = function () { /* empty */ };


/***/ }),
/* 181 */
/***/ (function(module, exports) {

module.exports = function (it, Constructor, name, forbiddenField) {
  if (!(it instanceof Constructor) || (forbiddenField !== undefined && forbiddenField in it)) {
    throw TypeError(name + ': incorrect invocation!');
  } return it;
};


/***/ }),
/* 182 */
/***/ (function(module, exports, __webpack_require__) {

// false -> Array#indexOf
// true  -> Array#includes
var toIObject = __webpack_require__(74);
var toLength = __webpack_require__(110);
var toAbsoluteIndex = __webpack_require__(200);
module.exports = function (IS_INCLUDES) {
  return function ($this, el, fromIndex) {
    var O = toIObject($this);
    var length = toLength(O.length);
    var index = toAbsoluteIndex(fromIndex, length);
    var value;
    // Array#includes uses SameValueZero equality algorithm
    // eslint-disable-next-line no-self-compare
    if (IS_INCLUDES && el != el) while (length > index) {
      value = O[index++];
      // eslint-disable-next-line no-self-compare
      if (value != value) return true;
    // Array#indexOf ignores holes, Array#includes - not
    } else for (;length > index; index++) if (IS_INCLUDES || index in O) {
      if (O[index] === el) return IS_INCLUDES || index || 0;
    } return !IS_INCLUDES && -1;
  };
};


/***/ }),
/* 183 */
/***/ (function(module, exports, __webpack_require__) {

var ctx = __webpack_require__(48);
var call = __webpack_require__(187);
var isArrayIter = __webpack_require__(186);
var anObject = __webpack_require__(27);
var toLength = __webpack_require__(110);
var getIterFn = __webpack_require__(203);
var BREAK = {};
var RETURN = {};
var exports = module.exports = function (iterable, entries, fn, that, ITERATOR) {
  var iterFn = ITERATOR ? function () { return iterable; } : getIterFn(iterable);
  var f = ctx(fn, that, entries ? 2 : 1);
  var index = 0;
  var length, step, iterator, result;
  if (typeof iterFn != 'function') throw TypeError(iterable + ' is not iterable!');
  // fast case for arrays with default iterator
  if (isArrayIter(iterFn)) for (length = toLength(iterable.length); length > index; index++) {
    result = entries ? f(anObject(step = iterable[index])[0], step[1]) : f(iterable[index]);
    if (result === BREAK || result === RETURN) return result;
  } else for (iterator = iterFn.call(iterable); !(step = iterator.next()).done;) {
    result = call(iterator, f, step.value, entries);
    if (result === BREAK || result === RETURN) return result;
  }
};
exports.BREAK = BREAK;
exports.RETURN = RETURN;


/***/ }),
/* 184 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = !__webpack_require__(33) && !__webpack_require__(85)(function () {
  return Object.defineProperty(__webpack_require__(68)('div'), 'a', { get: function () { return 7; } }).a != 7;
});


/***/ }),
/* 185 */
/***/ (function(module, exports) {

// fast apply, http://jsperf.lnkit.com/fast-apply/5
module.exports = function (fn, args, that) {
  var un = that === undefined;
  switch (args.length) {
    case 0: return un ? fn()
                      : fn.call(that);
    case 1: return un ? fn(args[0])
                      : fn.call(that, args[0]);
    case 2: return un ? fn(args[0], args[1])
                      : fn.call(that, args[0], args[1]);
    case 3: return un ? fn(args[0], args[1], args[2])
                      : fn.call(that, args[0], args[1], args[2]);
    case 4: return un ? fn(args[0], args[1], args[2], args[3])
                      : fn.call(that, args[0], args[1], args[2], args[3]);
  } return fn.apply(that, args);
};


/***/ }),
/* 186 */
/***/ (function(module, exports, __webpack_require__) {

// check on default Array iterator
var Iterators = __webpack_require__(35);
var ITERATOR = __webpack_require__(13)('iterator');
var ArrayProto = Array.prototype;

module.exports = function (it) {
  return it !== undefined && (Iterators.Array === it || ArrayProto[ITERATOR] === it);
};


/***/ }),
/* 187 */
/***/ (function(module, exports, __webpack_require__) {

// call something on iterator step with safe closing on error
var anObject = __webpack_require__(27);
module.exports = function (iterator, fn, value, entries) {
  try {
    return entries ? fn(anObject(value)[0], value[1]) : fn(value);
  // 7.4.6 IteratorClose(iterator, completion)
  } catch (e) {
    var ret = iterator['return'];
    if (ret !== undefined) anObject(ret.call(iterator));
    throw e;
  }
};


/***/ }),
/* 188 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var create = __webpack_require__(192);
var descriptor = __webpack_require__(106);
var setToStringTag = __webpack_require__(71);
var IteratorPrototype = {};

// 25.1.2.1.1 %IteratorPrototype%[@@iterator]()
__webpack_require__(31)(IteratorPrototype, __webpack_require__(13)('iterator'), function () { return this; });

module.exports = function (Constructor, NAME, next) {
  Constructor.prototype = create(IteratorPrototype, { next: descriptor(1, next) });
  setToStringTag(Constructor, NAME + ' Iterator');
};


/***/ }),
/* 189 */
/***/ (function(module, exports, __webpack_require__) {

var ITERATOR = __webpack_require__(13)('iterator');
var SAFE_CLOSING = false;

try {
  var riter = [7][ITERATOR]();
  riter['return'] = function () { SAFE_CLOSING = true; };
  // eslint-disable-next-line no-throw-literal
  Array.from(riter, function () { throw 2; });
} catch (e) { /* empty */ }

module.exports = function (exec, skipClosing) {
  if (!skipClosing && !SAFE_CLOSING) return false;
  var safe = false;
  try {
    var arr = [7];
    var iter = arr[ITERATOR]();
    iter.next = function () { return { done: safe = true }; };
    arr[ITERATOR] = function () { return iter; };
    exec(arr);
  } catch (e) { /* empty */ }
  return safe;
};


/***/ }),
/* 190 */
/***/ (function(module, exports) {

module.exports = function (done, value) {
  return { value: value, done: !!done };
};


/***/ }),
/* 191 */
/***/ (function(module, exports, __webpack_require__) {

var global = __webpack_require__(11);
var macrotask = __webpack_require__(109).set;
var Observer = global.MutationObserver || global.WebKitMutationObserver;
var process = global.process;
var Promise = global.Promise;
var isNode = __webpack_require__(47)(process) == 'process';

module.exports = function () {
  var head, last, notify;

  var flush = function () {
    var parent, fn;
    if (isNode && (parent = process.domain)) parent.exit();
    while (head) {
      fn = head.fn;
      head = head.next;
      try {
        fn();
      } catch (e) {
        if (head) notify();
        else last = undefined;
        throw e;
      }
    } last = undefined;
    if (parent) parent.enter();
  };

  // Node.js
  if (isNode) {
    notify = function () {
      process.nextTick(flush);
    };
  // browsers with MutationObserver, except iOS Safari - https://github.com/zloirock/core-js/issues/339
  } else if (Observer && !(global.navigator && global.navigator.standalone)) {
    var toggle = true;
    var node = document.createTextNode('');
    new Observer(flush).observe(node, { characterData: true }); // eslint-disable-line no-new
    notify = function () {
      node.data = toggle = !toggle;
    };
  // environments with maybe non-completely correct, but existent Promise
  } else if (Promise && Promise.resolve) {
    // Promise.resolve without an argument throws an error in LG WebOS 2
    var promise = Promise.resolve(undefined);
    notify = function () {
      promise.then(flush);
    };
  // for other environments - macrotask based on:
  // - setImmediate
  // - MessageChannel
  // - window.postMessag
  // - onreadystatechange
  // - setTimeout
  } else {
    notify = function () {
      // strange IE + webpack dev server bug - use .call(global)
      macrotask.call(global, flush);
    };
  }

  return function (fn) {
    var task = { fn: fn, next: undefined };
    if (last) last.next = task;
    if (!head) {
      head = task;
      notify();
    } last = task;
  };
};


/***/ }),
/* 192 */
/***/ (function(module, exports, __webpack_require__) {

// 19.1.2.2 / 15.2.3.5 Object.create(O [, Properties])
var anObject = __webpack_require__(27);
var dPs = __webpack_require__(193);
var enumBugKeys = __webpack_require__(101);
var IE_PROTO = __webpack_require__(72)('IE_PROTO');
var Empty = function () { /* empty */ };
var PROTOTYPE = 'prototype';

// Create object with fake `null` prototype: use iframe Object with cleared prototype
var createDict = function () {
  // Thrash, waste and sodomy: IE GC bug
  var iframe = __webpack_require__(68)('iframe');
  var i = enumBugKeys.length;
  var lt = '<';
  var gt = '>';
  var iframeDocument;
  iframe.style.display = 'none';
  __webpack_require__(102).appendChild(iframe);
  iframe.src = 'javascript:'; // eslint-disable-line no-script-url
  // createDict = iframe.contentWindow.Object;
  // html.removeChild(iframe);
  iframeDocument = iframe.contentWindow.document;
  iframeDocument.open();
  iframeDocument.write(lt + 'script' + gt + 'document.F=Object' + lt + '/script' + gt);
  iframeDocument.close();
  createDict = iframeDocument.F;
  while (i--) delete createDict[PROTOTYPE][enumBugKeys[i]];
  return createDict();
};

module.exports = Object.create || function create(O, Properties) {
  var result;
  if (O !== null) {
    Empty[PROTOTYPE] = anObject(O);
    result = new Empty();
    Empty[PROTOTYPE] = null;
    // add "__proto__" for Object.getPrototypeOf polyfill
    result[IE_PROTO] = O;
  } else result = createDict();
  return Properties === undefined ? result : dPs(result, Properties);
};


/***/ }),
/* 193 */
/***/ (function(module, exports, __webpack_require__) {

var dP = __webpack_require__(50);
var anObject = __webpack_require__(27);
var getKeys = __webpack_require__(132);

module.exports = __webpack_require__(33) ? Object.defineProperties : function defineProperties(O, Properties) {
  anObject(O);
  var keys = getKeys(Properties);
  var length = keys.length;
  var i = 0;
  var P;
  while (length > i) dP.f(O, P = keys[i++], Properties[P]);
  return O;
};


/***/ }),
/* 194 */
/***/ (function(module, exports, __webpack_require__) {

// 19.1.2.9 / 15.2.3.2 Object.getPrototypeOf(O)
var has = __webpack_require__(49);
var toObject = __webpack_require__(133);
var IE_PROTO = __webpack_require__(72)('IE_PROTO');
var ObjectProto = Object.prototype;

module.exports = Object.getPrototypeOf || function (O) {
  O = toObject(O);
  if (has(O, IE_PROTO)) return O[IE_PROTO];
  if (typeof O.constructor == 'function' && O instanceof O.constructor) {
    return O.constructor.prototype;
  } return O instanceof Object ? ObjectProto : null;
};


/***/ }),
/* 195 */
/***/ (function(module, exports, __webpack_require__) {

var has = __webpack_require__(49);
var toIObject = __webpack_require__(74);
var arrayIndexOf = __webpack_require__(182)(false);
var IE_PROTO = __webpack_require__(72)('IE_PROTO');

module.exports = function (object, names) {
  var O = toIObject(object);
  var i = 0;
  var result = [];
  var key;
  for (key in O) if (key != IE_PROTO) has(O, key) && result.push(key);
  // Don't enum bug & hidden keys
  while (names.length > i) if (has(O, key = names[i++])) {
    ~arrayIndexOf(result, key) || result.push(key);
  }
  return result;
};


/***/ }),
/* 196 */
/***/ (function(module, exports, __webpack_require__) {

var hide = __webpack_require__(31);
module.exports = function (target, src, safe) {
  for (var key in src) {
    if (safe && target[key]) target[key] = src[key];
    else hide(target, key, src[key]);
  } return target;
};


/***/ }),
/* 197 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(31);


/***/ }),
/* 198 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var global = __webpack_require__(11);
var core = __webpack_require__(23);
var dP = __webpack_require__(50);
var DESCRIPTORS = __webpack_require__(33);
var SPECIES = __webpack_require__(13)('species');

module.exports = function (KEY) {
  var C = typeof core[KEY] == 'function' ? core[KEY] : global[KEY];
  if (DESCRIPTORS && C && !C[SPECIES]) dP.f(C, SPECIES, {
    configurable: true,
    get: function () { return this; }
  });
};


/***/ }),
/* 199 */
/***/ (function(module, exports, __webpack_require__) {

var toInteger = __webpack_require__(73);
var defined = __webpack_require__(67);
// true  -> String#at
// false -> String#codePointAt
module.exports = function (TO_STRING) {
  return function (that, pos) {
    var s = String(defined(that));
    var i = toInteger(pos);
    var l = s.length;
    var a, b;
    if (i < 0 || i >= l) return TO_STRING ? '' : undefined;
    a = s.charCodeAt(i);
    return a < 0xd800 || a > 0xdbff || i + 1 === l || (b = s.charCodeAt(i + 1)) < 0xdc00 || b > 0xdfff
      ? TO_STRING ? s.charAt(i) : a
      : TO_STRING ? s.slice(i, i + 2) : (a - 0xd800 << 10) + (b - 0xdc00) + 0x10000;
  };
};


/***/ }),
/* 200 */
/***/ (function(module, exports, __webpack_require__) {

var toInteger = __webpack_require__(73);
var max = Math.max;
var min = Math.min;
module.exports = function (index, length) {
  index = toInteger(index);
  return index < 0 ? max(index + length, 0) : min(index, length);
};


/***/ }),
/* 201 */
/***/ (function(module, exports, __webpack_require__) {

// 7.1.1 ToPrimitive(input [, PreferredType])
var isObject = __webpack_require__(34);
// instead of the ES6 spec version, we didn't implement @@toPrimitive case
// and the second argument - flag - preferred type is a string
module.exports = function (it, S) {
  if (!isObject(it)) return it;
  var fn, val;
  if (S && typeof (fn = it.toString) == 'function' && !isObject(val = fn.call(it))) return val;
  if (typeof (fn = it.valueOf) == 'function' && !isObject(val = fn.call(it))) return val;
  if (!S && typeof (fn = it.toString) == 'function' && !isObject(val = fn.call(it))) return val;
  throw TypeError("Can't convert object to primitive value");
};


/***/ }),
/* 202 */
/***/ (function(module, exports, __webpack_require__) {

var global = __webpack_require__(11);
var navigator = global.navigator;

module.exports = navigator && navigator.userAgent || '';


/***/ }),
/* 203 */
/***/ (function(module, exports, __webpack_require__) {

var classof = __webpack_require__(100);
var ITERATOR = __webpack_require__(13)('iterator');
var Iterators = __webpack_require__(35);
module.exports = __webpack_require__(23).getIteratorMethod = function (it) {
  if (it != undefined) return it[ITERATOR]
    || it['@@iterator']
    || Iterators[classof(it)];
};


/***/ }),
/* 204 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var addToUnscopables = __webpack_require__(180);
var step = __webpack_require__(190);
var Iterators = __webpack_require__(35);
var toIObject = __webpack_require__(74);

// 22.1.3.4 Array.prototype.entries()
// 22.1.3.13 Array.prototype.keys()
// 22.1.3.29 Array.prototype.values()
// 22.1.3.30 Array.prototype[@@iterator]()
module.exports = __webpack_require__(103)(Array, 'Array', function (iterated, kind) {
  this._t = toIObject(iterated); // target
  this._i = 0;                   // next index
  this._k = kind;                // kind
// 22.1.5.2.1 %ArrayIteratorPrototype%.next()
}, function () {
  var O = this._t;
  var kind = this._k;
  var index = this._i++;
  if (!O || index >= O.length) {
    this._t = undefined;
    return step(1);
  }
  if (kind == 'keys') return step(0, index);
  if (kind == 'values') return step(0, O[index]);
  return step(0, [index, O[index]]);
}, 'values');

// argumentsList[@@iterator] is %ArrayProto_values% (9.4.4.6, 9.4.4.7)
Iterators.Arguments = Iterators.Array;

addToUnscopables('keys');
addToUnscopables('values');
addToUnscopables('entries');


/***/ }),
/* 205 */
/***/ (function(module, exports) {



/***/ }),
/* 206 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var LIBRARY = __webpack_require__(69);
var global = __webpack_require__(11);
var ctx = __webpack_require__(48);
var classof = __webpack_require__(100);
var $export = __webpack_require__(41);
var isObject = __webpack_require__(34);
var aFunction = __webpack_require__(46);
var anInstance = __webpack_require__(181);
var forOf = __webpack_require__(183);
var speciesConstructor = __webpack_require__(108);
var task = __webpack_require__(109).set;
var microtask = __webpack_require__(191)();
var newPromiseCapabilityModule = __webpack_require__(70);
var perform = __webpack_require__(104);
var userAgent = __webpack_require__(202);
var promiseResolve = __webpack_require__(105);
var PROMISE = 'Promise';
var TypeError = global.TypeError;
var process = global.process;
var versions = process && process.versions;
var v8 = versions && versions.v8 || '';
var $Promise = global[PROMISE];
var isNode = classof(process) == 'process';
var empty = function () { /* empty */ };
var Internal, newGenericPromiseCapability, OwnPromiseCapability, Wrapper;
var newPromiseCapability = newGenericPromiseCapability = newPromiseCapabilityModule.f;

var USE_NATIVE = !!function () {
  try {
    // correct subclassing with @@species support
    var promise = $Promise.resolve(1);
    var FakePromise = (promise.constructor = {})[__webpack_require__(13)('species')] = function (exec) {
      exec(empty, empty);
    };
    // unhandled rejections tracking support, NodeJS Promise without it fails @@species test
    return (isNode || typeof PromiseRejectionEvent == 'function')
      && promise.then(empty) instanceof FakePromise
      // v8 6.6 (Node 10 and Chrome 66) have a bug with resolving custom thenables
      // https://bugs.chromium.org/p/chromium/issues/detail?id=830565
      // we can't detect it synchronously, so just check versions
      && v8.indexOf('6.6') !== 0
      && userAgent.indexOf('Chrome/66') === -1;
  } catch (e) { /* empty */ }
}();

// helpers
var isThenable = function (it) {
  var then;
  return isObject(it) && typeof (then = it.then) == 'function' ? then : false;
};
var notify = function (promise, isReject) {
  if (promise._n) return;
  promise._n = true;
  var chain = promise._c;
  microtask(function () {
    var value = promise._v;
    var ok = promise._s == 1;
    var i = 0;
    var run = function (reaction) {
      var handler = ok ? reaction.ok : reaction.fail;
      var resolve = reaction.resolve;
      var reject = reaction.reject;
      var domain = reaction.domain;
      var result, then, exited;
      try {
        if (handler) {
          if (!ok) {
            if (promise._h == 2) onHandleUnhandled(promise);
            promise._h = 1;
          }
          if (handler === true) result = value;
          else {
            if (domain) domain.enter();
            result = handler(value); // may throw
            if (domain) {
              domain.exit();
              exited = true;
            }
          }
          if (result === reaction.promise) {
            reject(TypeError('Promise-chain cycle'));
          } else if (then = isThenable(result)) {
            then.call(result, resolve, reject);
          } else resolve(result);
        } else reject(value);
      } catch (e) {
        if (domain && !exited) domain.exit();
        reject(e);
      }
    };
    while (chain.length > i) run(chain[i++]); // variable length - can't use forEach
    promise._c = [];
    promise._n = false;
    if (isReject && !promise._h) onUnhandled(promise);
  });
};
var onUnhandled = function (promise) {
  task.call(global, function () {
    var value = promise._v;
    var unhandled = isUnhandled(promise);
    var result, handler, console;
    if (unhandled) {
      result = perform(function () {
        if (isNode) {
          process.emit('unhandledRejection', value, promise);
        } else if (handler = global.onunhandledrejection) {
          handler({ promise: promise, reason: value });
        } else if ((console = global.console) && console.error) {
          console.error('Unhandled promise rejection', value);
        }
      });
      // Browsers should not trigger `rejectionHandled` event if it was handled here, NodeJS - should
      promise._h = isNode || isUnhandled(promise) ? 2 : 1;
    } promise._a = undefined;
    if (unhandled && result.e) throw result.v;
  });
};
var isUnhandled = function (promise) {
  return promise._h !== 1 && (promise._a || promise._c).length === 0;
};
var onHandleUnhandled = function (promise) {
  task.call(global, function () {
    var handler;
    if (isNode) {
      process.emit('rejectionHandled', promise);
    } else if (handler = global.onrejectionhandled) {
      handler({ promise: promise, reason: promise._v });
    }
  });
};
var $reject = function (value) {
  var promise = this;
  if (promise._d) return;
  promise._d = true;
  promise = promise._w || promise; // unwrap
  promise._v = value;
  promise._s = 2;
  if (!promise._a) promise._a = promise._c.slice();
  notify(promise, true);
};
var $resolve = function (value) {
  var promise = this;
  var then;
  if (promise._d) return;
  promise._d = true;
  promise = promise._w || promise; // unwrap
  try {
    if (promise === value) throw TypeError("Promise can't be resolved itself");
    if (then = isThenable(value)) {
      microtask(function () {
        var wrapper = { _w: promise, _d: false }; // wrap
        try {
          then.call(value, ctx($resolve, wrapper, 1), ctx($reject, wrapper, 1));
        } catch (e) {
          $reject.call(wrapper, e);
        }
      });
    } else {
      promise._v = value;
      promise._s = 1;
      notify(promise, false);
    }
  } catch (e) {
    $reject.call({ _w: promise, _d: false }, e); // wrap
  }
};

// constructor polyfill
if (!USE_NATIVE) {
  // 25.4.3.1 Promise(executor)
  $Promise = function Promise(executor) {
    anInstance(this, $Promise, PROMISE, '_h');
    aFunction(executor);
    Internal.call(this);
    try {
      executor(ctx($resolve, this, 1), ctx($reject, this, 1));
    } catch (err) {
      $reject.call(this, err);
    }
  };
  // eslint-disable-next-line no-unused-vars
  Internal = function Promise(executor) {
    this._c = [];             // <- awaiting reactions
    this._a = undefined;      // <- checked in isUnhandled reactions
    this._s = 0;              // <- state
    this._d = false;          // <- done
    this._v = undefined;      // <- value
    this._h = 0;              // <- rejection state, 0 - default, 1 - handled, 2 - unhandled
    this._n = false;          // <- notify
  };
  Internal.prototype = __webpack_require__(196)($Promise.prototype, {
    // 25.4.5.3 Promise.prototype.then(onFulfilled, onRejected)
    then: function then(onFulfilled, onRejected) {
      var reaction = newPromiseCapability(speciesConstructor(this, $Promise));
      reaction.ok = typeof onFulfilled == 'function' ? onFulfilled : true;
      reaction.fail = typeof onRejected == 'function' && onRejected;
      reaction.domain = isNode ? process.domain : undefined;
      this._c.push(reaction);
      if (this._a) this._a.push(reaction);
      if (this._s) notify(this, false);
      return reaction.promise;
    },
    // 25.4.5.1 Promise.prototype.catch(onRejected)
    'catch': function (onRejected) {
      return this.then(undefined, onRejected);
    }
  });
  OwnPromiseCapability = function () {
    var promise = new Internal();
    this.promise = promise;
    this.resolve = ctx($resolve, promise, 1);
    this.reject = ctx($reject, promise, 1);
  };
  newPromiseCapabilityModule.f = newPromiseCapability = function (C) {
    return C === $Promise || C === Wrapper
      ? new OwnPromiseCapability(C)
      : newGenericPromiseCapability(C);
  };
}

$export($export.G + $export.W + $export.F * !USE_NATIVE, { Promise: $Promise });
__webpack_require__(71)($Promise, PROMISE);
__webpack_require__(198)(PROMISE);
Wrapper = __webpack_require__(23)[PROMISE];

// statics
$export($export.S + $export.F * !USE_NATIVE, PROMISE, {
  // 25.4.4.5 Promise.reject(r)
  reject: function reject(r) {
    var capability = newPromiseCapability(this);
    var $$reject = capability.reject;
    $$reject(r);
    return capability.promise;
  }
});
$export($export.S + $export.F * (LIBRARY || !USE_NATIVE), PROMISE, {
  // 25.4.4.6 Promise.resolve(x)
  resolve: function resolve(x) {
    return promiseResolve(LIBRARY && this === Wrapper ? $Promise : this, x);
  }
});
$export($export.S + $export.F * !(USE_NATIVE && __webpack_require__(189)(function (iter) {
  $Promise.all(iter)['catch'](empty);
})), PROMISE, {
  // 25.4.4.1 Promise.all(iterable)
  all: function all(iterable) {
    var C = this;
    var capability = newPromiseCapability(C);
    var resolve = capability.resolve;
    var reject = capability.reject;
    var result = perform(function () {
      var values = [];
      var index = 0;
      var remaining = 1;
      forOf(iterable, false, function (promise) {
        var $index = index++;
        var alreadyCalled = false;
        values.push(undefined);
        remaining++;
        C.resolve(promise).then(function (value) {
          if (alreadyCalled) return;
          alreadyCalled = true;
          values[$index] = value;
          --remaining || resolve(values);
        }, reject);
      });
      --remaining || resolve(values);
    });
    if (result.e) reject(result.v);
    return capability.promise;
  },
  // 25.4.4.4 Promise.race(iterable)
  race: function race(iterable) {
    var C = this;
    var capability = newPromiseCapability(C);
    var reject = capability.reject;
    var result = perform(function () {
      forOf(iterable, false, function (promise) {
        C.resolve(promise).then(capability.resolve, reject);
      });
    });
    if (result.e) reject(result.v);
    return capability.promise;
  }
});


/***/ }),
/* 207 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var $at = __webpack_require__(199)(true);

// 21.1.3.27 String.prototype[@@iterator]()
__webpack_require__(103)(String, 'String', function (iterated) {
  this._t = String(iterated); // target
  this._i = 0;                // next index
// 21.1.5.2.1 %StringIteratorPrototype%.next()
}, function () {
  var O = this._t;
  var index = this._i;
  var point;
  if (index >= O.length) return { value: undefined, done: true };
  point = $at(O, index);
  this._i += point.length;
  return { value: point, done: false };
});


/***/ }),
/* 208 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
// https://github.com/tc39/proposal-promise-finally

var $export = __webpack_require__(41);
var core = __webpack_require__(23);
var global = __webpack_require__(11);
var speciesConstructor = __webpack_require__(108);
var promiseResolve = __webpack_require__(105);

$export($export.P + $export.R, 'Promise', { 'finally': function (onFinally) {
  var C = speciesConstructor(this, core.Promise || global.Promise);
  var isFunction = typeof onFinally == 'function';
  return this.then(
    isFunction ? function (x) {
      return promiseResolve(C, onFinally()).then(function () { return x; });
    } : onFinally,
    isFunction ? function (e) {
      return promiseResolve(C, onFinally()).then(function () { throw e; });
    } : onFinally
  );
} });


/***/ }),
/* 209 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

// https://github.com/tc39/proposal-promise-try
var $export = __webpack_require__(41);
var newPromiseCapability = __webpack_require__(70);
var perform = __webpack_require__(104);

$export($export.S, 'Promise', { 'try': function (callbackfn) {
  var promiseCapability = newPromiseCapability.f(this);
  var result = perform(callbackfn);
  (result.e ? promiseCapability.reject : promiseCapability.resolve)(result.v);
  return promiseCapability.promise;
} });


/***/ }),
/* 210 */
/***/ (function(module, exports, __webpack_require__) {

__webpack_require__(204);
var global = __webpack_require__(11);
var hide = __webpack_require__(31);
var Iterators = __webpack_require__(35);
var TO_STRING_TAG = __webpack_require__(13)('toStringTag');

var DOMIterables = ('CSSRuleList,CSSStyleDeclaration,CSSValueList,ClientRectList,DOMRectList,DOMStringList,' +
  'DOMTokenList,DataTransferItemList,FileList,HTMLAllCollection,HTMLCollection,HTMLFormElement,HTMLSelectElement,' +
  'MediaList,MimeTypeArray,NamedNodeMap,NodeList,PaintRequestList,Plugin,PluginArray,SVGLengthList,SVGNumberList,' +
  'SVGPathSegList,SVGPointList,SVGStringList,SVGTransformList,SourceBufferList,StyleSheetList,TextTrackCueList,' +
  'TextTrackList,TouchList').split(',');

for (var i = 0; i < DOMIterables.length; i++) {
  var NAME = DOMIterables[i];
  var Collection = global[NAME];
  var proto = Collection && Collection.prototype;
  if (proto && !proto[TO_STRING_TAG]) hide(proto, TO_STRING_TAG, NAME);
  Iterators[NAME] = Iterators.Array;
}


/***/ }),
/* 211 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * This is the web browser implementation of `debug()`.
 *
 * Expose `debug()` as the module.
 */

exports = module.exports = __webpack_require__(112);
exports.log = log;
exports.formatArgs = formatArgs;
exports.save = save;
exports.load = load;
exports.useColors = useColors;
exports.storage = 'undefined' != typeof chrome
               && 'undefined' != typeof chrome.storage
                  ? chrome.storage.local
                  : localstorage();

/**
 * Colors.
 */

exports.colors = [
  '#0000CC', '#0000FF', '#0033CC', '#0033FF', '#0066CC', '#0066FF', '#0099CC',
  '#0099FF', '#00CC00', '#00CC33', '#00CC66', '#00CC99', '#00CCCC', '#00CCFF',
  '#3300CC', '#3300FF', '#3333CC', '#3333FF', '#3366CC', '#3366FF', '#3399CC',
  '#3399FF', '#33CC00', '#33CC33', '#33CC66', '#33CC99', '#33CCCC', '#33CCFF',
  '#6600CC', '#6600FF', '#6633CC', '#6633FF', '#66CC00', '#66CC33', '#9900CC',
  '#9900FF', '#9933CC', '#9933FF', '#99CC00', '#99CC33', '#CC0000', '#CC0033',
  '#CC0066', '#CC0099', '#CC00CC', '#CC00FF', '#CC3300', '#CC3333', '#CC3366',
  '#CC3399', '#CC33CC', '#CC33FF', '#CC6600', '#CC6633', '#CC9900', '#CC9933',
  '#CCCC00', '#CCCC33', '#FF0000', '#FF0033', '#FF0066', '#FF0099', '#FF00CC',
  '#FF00FF', '#FF3300', '#FF3333', '#FF3366', '#FF3399', '#FF33CC', '#FF33FF',
  '#FF6600', '#FF6633', '#FF9900', '#FF9933', '#FFCC00', '#FFCC33'
];

/**
 * Currently only WebKit-based Web Inspectors, Firefox >= v31,
 * and the Firebug extension (any Firefox version) are known
 * to support "%c" CSS customizations.
 *
 * TODO: add a `localStorage` variable to explicitly enable/disable colors
 */

function useColors() {
  // NB: In an Electron preload script, document will be defined but not fully
  // initialized. Since we know we're in Chrome, we'll just detect this case
  // explicitly
  if (typeof window !== 'undefined' && window.process && window.process.type === 'renderer') {
    return true;
  }

  // Internet Explorer and Edge do not support colors.
  if (typeof navigator !== 'undefined' && navigator.userAgent && navigator.userAgent.toLowerCase().match(/(edge|trident)\/(\d+)/)) {
    return false;
  }

  // is webkit? http://stackoverflow.com/a/16459606/376773
  // document is undefined in react-native: https://github.com/facebook/react-native/pull/1632
  return (typeof document !== 'undefined' && document.documentElement && document.documentElement.style && document.documentElement.style.WebkitAppearance) ||
    // is firebug? http://stackoverflow.com/a/398120/376773
    (typeof window !== 'undefined' && window.console && (window.console.firebug || (window.console.exception && window.console.table))) ||
    // is firefox >= v31?
    // https://developer.mozilla.org/en-US/docs/Tools/Web_Console#Styling_messages
    (typeof navigator !== 'undefined' && navigator.userAgent && navigator.userAgent.toLowerCase().match(/firefox\/(\d+)/) && parseInt(RegExp.$1, 10) >= 31) ||
    // double check webkit in userAgent just in case we are in a worker
    (typeof navigator !== 'undefined' && navigator.userAgent && navigator.userAgent.toLowerCase().match(/applewebkit\/(\d+)/));
}

/**
 * Map %j to `JSON.stringify()`, since no Web Inspectors do that by default.
 */

exports.formatters.j = function(v) {
  try {
    return JSON.stringify(v);
  } catch (err) {
    return '[UnexpectedJSONParseError]: ' + err.message;
  }
};


/**
 * Colorize log arguments if enabled.
 *
 * @api public
 */

function formatArgs(args) {
  var useColors = this.useColors;

  args[0] = (useColors ? '%c' : '')
    + this.namespace
    + (useColors ? ' %c' : ' ')
    + args[0]
    + (useColors ? '%c ' : ' ')
    + '+' + exports.humanize(this.diff);

  if (!useColors) return;

  var c = 'color: ' + this.color;
  args.splice(1, 0, c, 'color: inherit')

  // the final "%c" is somewhat tricky, because there could be other
  // arguments passed either before or after the %c, so we need to
  // figure out the correct index to insert the CSS into
  var index = 0;
  var lastC = 0;
  args[0].replace(/%[a-zA-Z%]/g, function(match) {
    if ('%%' === match) return;
    index++;
    if ('%c' === match) {
      // we only are interested in the *last* %c
      // (the user may have provided their own)
      lastC = index;
    }
  });

  args.splice(lastC, 0, c);
}

/**
 * Invokes `console.log()` when available.
 * No-op when `console.log` is not a "function".
 *
 * @api public
 */

function log() {
  // this hackery is required for IE8/9, where
  // the `console.log` function doesn't have 'apply'
  return 'object' === typeof console
    && console.log
    && Function.prototype.apply.call(console.log, console, arguments);
}

/**
 * Save `namespaces`.
 *
 * @param {String} namespaces
 * @api private
 */

function save(namespaces) {
  try {
    if (null == namespaces) {
      exports.storage.removeItem('debug');
    } else {
      exports.storage.debug = namespaces;
    }
  } catch(e) {}
}

/**
 * Load `namespaces`.
 *
 * @return {String} returns the previously persisted debug modes
 * @api private
 */

function load() {
  var r;
  try {
    r = exports.storage.debug;
  } catch(e) {}

  // If debug isn't set in LS, and we're in Electron, try to load $DEBUG
  if (!r && typeof process !== 'undefined' && 'env' in process) {
    r = process.env.DEBUG;
  }

  return r;
}

/**
 * Enable namespaces listed in `localStorage.debug` initially.
 */

exports.enable(load());

/**
 * Localstorage attempts to return the localstorage.
 *
 * This is necessary because safari throws
 * when a user disables cookies/localstorage
 * and you attempt to access it.
 *
 * @return {LocalStorage}
 * @api private
 */

function localstorage() {
  try {
    return window.localStorage;
  } catch (e) {}
}


/***/ }),
/* 212 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * Detect Electron renderer process, which is node, but we should
 * treat as a browser.
 */

if (typeof process === 'undefined' || process.type === 'renderer') {
  module.exports = __webpack_require__(211);
} else {
  module.exports = __webpack_require__(213);
}


/***/ }),
/* 213 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * Module dependencies.
 */

var tty = __webpack_require__(79);
var util = __webpack_require__(2);

/**
 * This is the Node.js implementation of `debug()`.
 *
 * Expose `debug()` as the module.
 */

exports = module.exports = __webpack_require__(112);
exports.init = init;
exports.log = log;
exports.formatArgs = formatArgs;
exports.save = save;
exports.load = load;
exports.useColors = useColors;

/**
 * Colors.
 */

exports.colors = [ 6, 2, 3, 4, 5, 1 ];

try {
  var supportsColor = __webpack_require__(239);
  if (supportsColor && supportsColor.level >= 2) {
    exports.colors = [
      20, 21, 26, 27, 32, 33, 38, 39, 40, 41, 42, 43, 44, 45, 56, 57, 62, 63, 68,
      69, 74, 75, 76, 77, 78, 79, 80, 81, 92, 93, 98, 99, 112, 113, 128, 129, 134,
      135, 148, 149, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171,
      172, 173, 178, 179, 184, 185, 196, 197, 198, 199, 200, 201, 202, 203, 204,
      205, 206, 207, 208, 209, 214, 215, 220, 221
    ];
  }
} catch (err) {
  // swallow - we only care if `supports-color` is available; it doesn't have to be.
}

/**
 * Build up the default `inspectOpts` object from the environment variables.
 *
 *   $ DEBUG_COLORS=no DEBUG_DEPTH=10 DEBUG_SHOW_HIDDEN=enabled node script.js
 */

exports.inspectOpts = Object.keys(process.env).filter(function (key) {
  return /^debug_/i.test(key);
}).reduce(function (obj, key) {
  // camel-case
  var prop = key
    .substring(6)
    .toLowerCase()
    .replace(/_([a-z])/g, function (_, k) { return k.toUpperCase() });

  // coerce string value into JS value
  var val = process.env[key];
  if (/^(yes|on|true|enabled)$/i.test(val)) val = true;
  else if (/^(no|off|false|disabled)$/i.test(val)) val = false;
  else if (val === 'null') val = null;
  else val = Number(val);

  obj[prop] = val;
  return obj;
}, {});

/**
 * Is stdout a TTY? Colored output is enabled when `true`.
 */

function useColors() {
  return 'colors' in exports.inspectOpts
    ? Boolean(exports.inspectOpts.colors)
    : tty.isatty(process.stderr.fd);
}

/**
 * Map %o to `util.inspect()`, all on a single line.
 */

exports.formatters.o = function(v) {
  this.inspectOpts.colors = this.useColors;
  return util.inspect(v, this.inspectOpts)
    .split('\n').map(function(str) {
      return str.trim()
    }).join(' ');
};

/**
 * Map %o to `util.inspect()`, allowing multiple lines if needed.
 */

exports.formatters.O = function(v) {
  this.inspectOpts.colors = this.useColors;
  return util.inspect(v, this.inspectOpts);
};

/**
 * Adds ANSI color escape codes if enabled.
 *
 * @api public
 */

function formatArgs(args) {
  var name = this.namespace;
  var useColors = this.useColors;

  if (useColors) {
    var c = this.color;
    var colorCode = '\u001b[3' + (c < 8 ? c : '8;5;' + c);
    var prefix = '  ' + colorCode + ';1m' + name + ' ' + '\u001b[0m';

    args[0] = prefix + args[0].split('\n').join('\n' + prefix);
    args.push(colorCode + 'm+' + exports.humanize(this.diff) + '\u001b[0m');
  } else {
    args[0] = getDate() + name + ' ' + args[0];
  }
}

function getDate() {
  if (exports.inspectOpts.hideDate) {
    return '';
  } else {
    return new Date().toISOString() + ' ';
  }
}

/**
 * Invokes `util.format()` with the specified arguments and writes to stderr.
 */

function log() {
  return process.stderr.write(util.format.apply(util, arguments) + '\n');
}

/**
 * Save `namespaces`.
 *
 * @param {String} namespaces
 * @api private
 */

function save(namespaces) {
  if (null == namespaces) {
    // If you set a process.env field to null or undefined, it gets cast to the
    // string 'null' or 'undefined'. Just delete instead.
    delete process.env.DEBUG;
  } else {
    process.env.DEBUG = namespaces;
  }
}

/**
 * Load `namespaces`.
 *
 * @return {String} returns the previously persisted debug modes
 * @api private
 */

function load() {
  return process.env.DEBUG;
}

/**
 * Init logic for `debug` instances.
 *
 * Create a new `inspectOpts` object in case `useColors` is set
 * differently for a particular `debug` instance.
 */

function init (debug) {
  debug.inspectOpts = {};

  var keys = Object.keys(exports.inspectOpts);
  for (var i = 0; i < keys.length; i++) {
    debug.inspectOpts[keys[i]] = exports.inspectOpts[keys[i]];
  }
}

/**
 * Enable namespaces listed in `process.env.DEBUG` initially.
 */

exports.enable(load());


/***/ }),
/* 214 */,
/* 215 */,
/* 216 */,
/* 217 */
/***/ (function(module, exports, __webpack_require__) {

// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

var pathModule = __webpack_require__(0);
var isWindows = process.platform === 'win32';
var fs = __webpack_require__(3);

// JavaScript implementation of realpath, ported from node pre-v6

var DEBUG = process.env.NODE_DEBUG && /fs/.test(process.env.NODE_DEBUG);

function rethrow() {
  // Only enable in debug mode. A backtrace uses ~1000 bytes of heap space and
  // is fairly slow to generate.
  var callback;
  if (DEBUG) {
    var backtrace = new Error;
    callback = debugCallback;
  } else
    callback = missingCallback;

  return callback;

  function debugCallback(err) {
    if (err) {
      backtrace.message = err.message;
      err = backtrace;
      missingCallback(err);
    }
  }

  function missingCallback(err) {
    if (err) {
      if (process.throwDeprecation)
        throw err;  // Forgot a callback but don't know where? Use NODE_DEBUG=fs
      else if (!process.noDeprecation) {
        var msg = 'fs: missing callback ' + (err.stack || err.message);
        if (process.traceDeprecation)
          console.trace(msg);
        else
          console.error(msg);
      }
    }
  }
}

function maybeCallback(cb) {
  return typeof cb === 'function' ? cb : rethrow();
}

var normalize = pathModule.normalize;

// Regexp that finds the next partion of a (partial) path
// result is [base_with_slash, base], e.g. ['somedir/', 'somedir']
if (isWindows) {
  var nextPartRe = /(.*?)(?:[\/\\]+|$)/g;
} else {
  var nextPartRe = /(.*?)(?:[\/]+|$)/g;
}

// Regex to find the device root, including trailing slash. E.g. 'c:\\'.
if (isWindows) {
  var splitRootRe = /^(?:[a-zA-Z]:|[\\\/]{2}[^\\\/]+[\\\/][^\\\/]+)?[\\\/]*/;
} else {
  var splitRootRe = /^[\/]*/;
}

exports.realpathSync = function realpathSync(p, cache) {
  // make p is absolute
  p = pathModule.resolve(p);

  if (cache && Object.prototype.hasOwnProperty.call(cache, p)) {
    return cache[p];
  }

  var original = p,
      seenLinks = {},
      knownHard = {};

  // current character position in p
  var pos;
  // the partial path so far, including a trailing slash if any
  var current;
  // the partial path without a trailing slash (except when pointing at a root)
  var base;
  // the partial path scanned in the previous round, with slash
  var previous;

  start();

  function start() {
    // Skip over roots
    var m = splitRootRe.exec(p);
    pos = m[0].length;
    current = m[0];
    base = m[0];
    previous = '';

    // On windows, check that the root exists. On unix there is no need.
    if (isWindows && !knownHard[base]) {
      fs.lstatSync(base);
      knownHard[base] = true;
    }
  }

  // walk down the path, swapping out linked pathparts for their real
  // values
  // NB: p.length changes.
  while (pos < p.length) {
    // find the next part
    nextPartRe.lastIndex = pos;
    var result = nextPartRe.exec(p);
    previous = current;
    current += result[0];
    base = previous + result[1];
    pos = nextPartRe.lastIndex;

    // continue if not a symlink
    if (knownHard[base] || (cache && cache[base] === base)) {
      continue;
    }

    var resolvedLink;
    if (cache && Object.prototype.hasOwnProperty.call(cache, base)) {
      // some known symbolic link.  no need to stat again.
      resolvedLink = cache[base];
    } else {
      var stat = fs.lstatSync(base);
      if (!stat.isSymbolicLink()) {
        knownHard[base] = true;
        if (cache) cache[base] = base;
        continue;
      }

      // read the link if it wasn't read before
      // dev/ino always return 0 on windows, so skip the check.
      var linkTarget = null;
      if (!isWindows) {
        var id = stat.dev.toString(32) + ':' + stat.ino.toString(32);
        if (seenLinks.hasOwnProperty(id)) {
          linkTarget = seenLinks[id];
        }
      }
      if (linkTarget === null) {
        fs.statSync(base);
        linkTarget = fs.readlinkSync(base);
      }
      resolvedLink = pathModule.resolve(previous, linkTarget);
      // track this, if given a cache.
      if (cache) cache[base] = resolvedLink;
      if (!isWindows) seenLinks[id] = linkTarget;
    }

    // resolve the link, then start over
    p = pathModule.resolve(resolvedLink, p.slice(pos));
    start();
  }

  if (cache) cache[original] = p;

  return p;
};


exports.realpath = function realpath(p, cache, cb) {
  if (typeof cb !== 'function') {
    cb = maybeCallback(cache);
    cache = null;
  }

  // make p is absolute
  p = pathModule.resolve(p);

  if (cache && Object.prototype.hasOwnProperty.call(cache, p)) {
    return process.nextTick(cb.bind(null, null, cache[p]));
  }

  var original = p,
      seenLinks = {},
      knownHard = {};

  // current character position in p
  var pos;
  // the partial path so far, including a trailing slash if any
  var current;
  // the partial path without a trailing slash (except when pointing at a root)
  var base;
  // the partial path scanned in the previous round, with slash
  var previous;

  start();

  function start() {
    // Skip over roots
    var m = splitRootRe.exec(p);
    pos = m[0].length;
    current = m[0];
    base = m[0];
    previous = '';

    // On windows, check that the root exists. On unix there is no need.
    if (isWindows && !knownHard[base]) {
      fs.lstat(base, function(err) {
        if (err) return cb(err);
        knownHard[base] = true;
        LOOP();
      });
    } else {
      process.nextTick(LOOP);
    }
  }

  // walk down the path, swapping out linked pathparts for their real
  // values
  function LOOP() {
    // stop if scanned past end of path
    if (pos >= p.length) {
      if (cache) cache[original] = p;
      return cb(null, p);
    }

    // find the next part
    nextPartRe.lastIndex = pos;
    var result = nextPartRe.exec(p);
    previous = current;
    current += result[0];
    base = previous + result[1];
    pos = nextPartRe.lastIndex;

    // continue if not a symlink
    if (knownHard[base] || (cache && cache[base] === base)) {
      return process.nextTick(LOOP);
    }

    if (cache && Object.prototype.hasOwnProperty.call(cache, base)) {
      // known symbolic link.  no need to stat again.
      return gotResolvedLink(cache[base]);
    }

    return fs.lstat(base, gotStat);
  }

  function gotStat(err, stat) {
    if (err) return cb(err);

    // if not a symlink, skip to the next path part
    if (!stat.isSymbolicLink()) {
      knownHard[base] = true;
      if (cache) cache[base] = base;
      return process.nextTick(LOOP);
    }

    // stat & read the link if not read before
    // call gotTarget as soon as the link target is known
    // dev/ino always return 0 on windows, so skip the check.
    if (!isWindows) {
      var id = stat.dev.toString(32) + ':' + stat.ino.toString(32);
      if (seenLinks.hasOwnProperty(id)) {
        return gotTarget(null, seenLinks[id], base);
      }
    }
    fs.stat(base, function(err) {
      if (err) return cb(err);

      fs.readlink(base, function(err, target) {
        if (!isWindows) seenLinks[id] = target;
        gotTarget(err, target);
      });
    });
  }

  function gotTarget(err, target, base) {
    if (err) return cb(err);

    var resolvedLink = pathModule.resolve(previous, target);
    if (cache) cache[base] = resolvedLink;
    gotResolvedLink(resolvedLink);
  }

  function gotResolvedLink(resolvedLink) {
    // resolve the link, then start over
    p = pathModule.resolve(resolvedLink, p.slice(pos));
    start();
  }
};


/***/ }),
/* 218 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = globSync
globSync.GlobSync = GlobSync

var fs = __webpack_require__(3)
var rp = __webpack_require__(114)
var minimatch = __webpack_require__(60)
var Minimatch = minimatch.Minimatch
var Glob = __webpack_require__(75).Glob
var util = __webpack_require__(2)
var path = __webpack_require__(0)
var assert = __webpack_require__(22)
var isAbsolute = __webpack_require__(76)
var common = __webpack_require__(115)
var alphasort = common.alphasort
var alphasorti = common.alphasorti
var setopts = common.setopts
var ownProp = common.ownProp
var childrenIgnored = common.childrenIgnored
var isIgnored = common.isIgnored

function globSync (pattern, options) {
  if (typeof options === 'function' || arguments.length === 3)
    throw new TypeError('callback provided to sync glob\n'+
                        'See: https://github.com/isaacs/node-glob/issues/167')

  return new GlobSync(pattern, options).found
}

function GlobSync (pattern, options) {
  if (!pattern)
    throw new Error('must provide pattern')

  if (typeof options === 'function' || arguments.length === 3)
    throw new TypeError('callback provided to sync glob\n'+
                        'See: https://github.com/isaacs/node-glob/issues/167')

  if (!(this instanceof GlobSync))
    return new GlobSync(pattern, options)

  setopts(this, pattern, options)

  if (this.noprocess)
    return this

  var n = this.minimatch.set.length
  this.matches = new Array(n)
  for (var i = 0; i < n; i ++) {
    this._process(this.minimatch.set[i], i, false)
  }
  this._finish()
}

GlobSync.prototype._finish = function () {
  assert(this instanceof GlobSync)
  if (this.realpath) {
    var self = this
    this.matches.forEach(function (matchset, index) {
      var set = self.matches[index] = Object.create(null)
      for (var p in matchset) {
        try {
          p = self._makeAbs(p)
          var real = rp.realpathSync(p, self.realpathCache)
          set[real] = true
        } catch (er) {
          if (er.syscall === 'stat')
            set[self._makeAbs(p)] = true
          else
            throw er
        }
      }
    })
  }
  common.finish(this)
}


GlobSync.prototype._process = function (pattern, index, inGlobStar) {
  assert(this instanceof GlobSync)

  // Get the first [n] parts of pattern that are all strings.
  var n = 0
  while (typeof pattern[n] === 'string') {
    n ++
  }
  // now n is the index of the first one that is *not* a string.

  // See if there's anything else
  var prefix
  switch (n) {
    // if not, then this is rather simple
    case pattern.length:
      this._processSimple(pattern.join('/'), index)
      return

    case 0:
      // pattern *starts* with some non-trivial item.
      // going to readdir(cwd), but not include the prefix in matches.
      prefix = null
      break

    default:
      // pattern has some string bits in the front.
      // whatever it starts with, whether that's 'absolute' like /foo/bar,
      // or 'relative' like '../baz'
      prefix = pattern.slice(0, n).join('/')
      break
  }

  var remain = pattern.slice(n)

  // get the list of entries.
  var read
  if (prefix === null)
    read = '.'
  else if (isAbsolute(prefix) || isAbsolute(pattern.join('/'))) {
    if (!prefix || !isAbsolute(prefix))
      prefix = '/' + prefix
    read = prefix
  } else
    read = prefix

  var abs = this._makeAbs(read)

  //if ignored, skip processing
  if (childrenIgnored(this, read))
    return

  var isGlobStar = remain[0] === minimatch.GLOBSTAR
  if (isGlobStar)
    this._processGlobStar(prefix, read, abs, remain, index, inGlobStar)
  else
    this._processReaddir(prefix, read, abs, remain, index, inGlobStar)
}


GlobSync.prototype._processReaddir = function (prefix, read, abs, remain, index, inGlobStar) {
  var entries = this._readdir(abs, inGlobStar)

  // if the abs isn't a dir, then nothing can match!
  if (!entries)
    return

  // It will only match dot entries if it starts with a dot, or if
  // dot is set.  Stuff like @(.foo|.bar) isn't allowed.
  var pn = remain[0]
  var negate = !!this.minimatch.negate
  var rawGlob = pn._glob
  var dotOk = this.dot || rawGlob.charAt(0) === '.'

  var matchedEntries = []
  for (var i = 0; i < entries.length; i++) {
    var e = entries[i]
    if (e.charAt(0) !== '.' || dotOk) {
      var m
      if (negate && !prefix) {
        m = !e.match(pn)
      } else {
        m = e.match(pn)
      }
      if (m)
        matchedEntries.push(e)
    }
  }

  var len = matchedEntries.length
  // If there are no matched entries, then nothing matches.
  if (len === 0)
    return

  // if this is the last remaining pattern bit, then no need for
  // an additional stat *unless* the user has specified mark or
  // stat explicitly.  We know they exist, since readdir returned
  // them.

  if (remain.length === 1 && !this.mark && !this.stat) {
    if (!this.matches[index])
      this.matches[index] = Object.create(null)

    for (var i = 0; i < len; i ++) {
      var e = matchedEntries[i]
      if (prefix) {
        if (prefix.slice(-1) !== '/')
          e = prefix + '/' + e
        else
          e = prefix + e
      }

      if (e.charAt(0) === '/' && !this.nomount) {
        e = path.join(this.root, e)
      }
      this._emitMatch(index, e)
    }
    // This was the last one, and no stats were needed
    return
  }

  // now test all matched entries as stand-ins for that part
  // of the pattern.
  remain.shift()
  for (var i = 0; i < len; i ++) {
    var e = matchedEntries[i]
    var newPattern
    if (prefix)
      newPattern = [prefix, e]
    else
      newPattern = [e]
    this._process(newPattern.concat(remain), index, inGlobStar)
  }
}


GlobSync.prototype._emitMatch = function (index, e) {
  if (isIgnored(this, e))
    return

  var abs = this._makeAbs(e)

  if (this.mark)
    e = this._mark(e)

  if (this.absolute) {
    e = abs
  }

  if (this.matches[index][e])
    return

  if (this.nodir) {
    var c = this.cache[abs]
    if (c === 'DIR' || Array.isArray(c))
      return
  }

  this.matches[index][e] = true

  if (this.stat)
    this._stat(e)
}


GlobSync.prototype._readdirInGlobStar = function (abs) {
  // follow all symlinked directories forever
  // just proceed as if this is a non-globstar situation
  if (this.follow)
    return this._readdir(abs, false)

  var entries
  var lstat
  var stat
  try {
    lstat = fs.lstatSync(abs)
  } catch (er) {
    if (er.code === 'ENOENT') {
      // lstat failed, doesn't exist
      return null
    }
  }

  var isSym = lstat && lstat.isSymbolicLink()
  this.symlinks[abs] = isSym

  // If it's not a symlink or a dir, then it's definitely a regular file.
  // don't bother doing a readdir in that case.
  if (!isSym && lstat && !lstat.isDirectory())
    this.cache[abs] = 'FILE'
  else
    entries = this._readdir(abs, false)

  return entries
}

GlobSync.prototype._readdir = function (abs, inGlobStar) {
  var entries

  if (inGlobStar && !ownProp(this.symlinks, abs))
    return this._readdirInGlobStar(abs)

  if (ownProp(this.cache, abs)) {
    var c = this.cache[abs]
    if (!c || c === 'FILE')
      return null

    if (Array.isArray(c))
      return c
  }

  try {
    return this._readdirEntries(abs, fs.readdirSync(abs))
  } catch (er) {
    this._readdirError(abs, er)
    return null
  }
}

GlobSync.prototype._readdirEntries = function (abs, entries) {
  // if we haven't asked to stat everything, then just
  // assume that everything in there exists, so we can avoid
  // having to stat it a second time.
  if (!this.mark && !this.stat) {
    for (var i = 0; i < entries.length; i ++) {
      var e = entries[i]
      if (abs === '/')
        e = abs + e
      else
        e = abs + '/' + e
      this.cache[e] = true
    }
  }

  this.cache[abs] = entries

  // mark and cache dir-ness
  return entries
}

GlobSync.prototype._readdirError = function (f, er) {
  // handle errors, and cache the information
  switch (er.code) {
    case 'ENOTSUP': // https://github.com/isaacs/node-glob/issues/205
    case 'ENOTDIR': // totally normal. means it *does* exist.
      var abs = this._makeAbs(f)
      this.cache[abs] = 'FILE'
      if (abs === this.cwdAbs) {
        var error = new Error(er.code + ' invalid cwd ' + this.cwd)
        error.path = this.cwd
        error.code = er.code
        throw error
      }
      break

    case 'ENOENT': // not terribly unusual
    case 'ELOOP':
    case 'ENAMETOOLONG':
    case 'UNKNOWN':
      this.cache[this._makeAbs(f)] = false
      break

    default: // some unusual error.  Treat as failure.
      this.cache[this._makeAbs(f)] = false
      if (this.strict)
        throw er
      if (!this.silent)
        console.error('glob error', er)
      break
  }
}

GlobSync.prototype._processGlobStar = function (prefix, read, abs, remain, index, inGlobStar) {

  var entries = this._readdir(abs, inGlobStar)

  // no entries means not a dir, so it can never have matches
  // foo.txt/** doesn't match foo.txt
  if (!entries)
    return

  // test without the globstar, and with every child both below
  // and replacing the globstar.
  var remainWithoutGlobStar = remain.slice(1)
  var gspref = prefix ? [ prefix ] : []
  var noGlobStar = gspref.concat(remainWithoutGlobStar)

  // the noGlobStar pattern exits the inGlobStar state
  this._process(noGlobStar, index, false)

  var len = entries.length
  var isSym = this.symlinks[abs]

  // If it's a symlink, and we're in a globstar, then stop
  if (isSym && inGlobStar)
    return

  for (var i = 0; i < len; i++) {
    var e = entries[i]
    if (e.charAt(0) === '.' && !this.dot)
      continue

    // these two cases enter the inGlobStar state
    var instead = gspref.concat(entries[i], remainWithoutGlobStar)
    this._process(instead, index, true)

    var below = gspref.concat(entries[i], remain)
    this._process(below, index, true)
  }
}

GlobSync.prototype._processSimple = function (prefix, index) {
  // XXX review this.  Shouldn't it be doing the mounting etc
  // before doing stat?  kinda weird?
  var exists = this._stat(prefix)

  if (!this.matches[index])
    this.matches[index] = Object.create(null)

  // If it doesn't exist, then just mark the lack of results
  if (!exists)
    return

  if (prefix && isAbsolute(prefix) && !this.nomount) {
    var trail = /[\/\\]$/.test(prefix)
    if (prefix.charAt(0) === '/') {
      prefix = path.join(this.root, prefix)
    } else {
      prefix = path.resolve(this.root, prefix)
      if (trail)
        prefix += '/'
    }
  }

  if (process.platform === 'win32')
    prefix = prefix.replace(/\\/g, '/')

  // Mark this as a match
  this._emitMatch(index, prefix)
}

// Returns either 'DIR', 'FILE', or false
GlobSync.prototype._stat = function (f) {
  var abs = this._makeAbs(f)
  var needDir = f.slice(-1) === '/'

  if (f.length > this.maxLength)
    return false

  if (!this.stat && ownProp(this.cache, abs)) {
    var c = this.cache[abs]

    if (Array.isArray(c))
      c = 'DIR'

    // It exists, but maybe not how we need it
    if (!needDir || c === 'DIR')
      return c

    if (needDir && c === 'FILE')
      return false

    // otherwise we have to stat, because maybe c=true
    // if we know it exists, but not what it is.
  }

  var exists
  var stat = this.statCache[abs]
  if (!stat) {
    var lstat
    try {
      lstat = fs.lstatSync(abs)
    } catch (er) {
      if (er && (er.code === 'ENOENT' || er.code === 'ENOTDIR')) {
        this.statCache[abs] = false
        return false
      }
    }

    if (lstat && lstat.isSymbolicLink()) {
      try {
        stat = fs.statSync(abs)
      } catch (er) {
        stat = lstat
      }
    } else {
      stat = lstat
    }
  }

  this.statCache[abs] = stat

  var c = true
  if (stat)
    c = stat.isDirectory() ? 'DIR' : 'FILE'

  this.cache[abs] = this.cache[abs] || c

  if (needDir && c === 'FILE')
    return false

  return c
}

GlobSync.prototype._mark = function (p) {
  return common.mark(this, p)
}

GlobSync.prototype._makeAbs = function (f) {
  return common.makeAbs(this, f)
}


/***/ }),
/* 219 */,
/* 220 */,
/* 221 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

module.exports = function (flag, argv) {
	argv = argv || process.argv;

	var terminatorPos = argv.indexOf('--');
	var prefix = /^--/.test(flag) ? '' : '--';
	var pos = argv.indexOf(prefix + flag);

	return pos !== -1 && (terminatorPos !== -1 ? pos < terminatorPos : true);
};


/***/ }),
/* 222 */,
/* 223 */
/***/ (function(module, exports, __webpack_require__) {

var wrappy = __webpack_require__(123)
var reqs = Object.create(null)
var once = __webpack_require__(61)

module.exports = wrappy(inflight)

function inflight (key, cb) {
  if (reqs[key]) {
    reqs[key].push(cb)
    return null
  } else {
    reqs[key] = [cb]
    return makeres(key)
  }
}

function makeres (key) {
  return once(function RES () {
    var cbs = reqs[key]
    var len = cbs.length
    var args = slice(arguments)

    // XXX It's somewhat ambiguous whether a new callback added in this
    // pass should be queued for later execution if something in the
    // list of callbacks throws, or if it should just be discarded.
    // However, it's such an edge case that it hardly matters, and either
    // choice is likely as surprising as the other.
    // As it happens, we do go ahead and schedule it for later execution.
    try {
      for (var i = 0; i < len; i++) {
        cbs[i].apply(null, args)
      }
    } finally {
      if (cbs.length > len) {
        // added more in the interim.
        // de-zalgo, just in case, but don't call again.
        cbs.splice(0, len)
        process.nextTick(function () {
          RES.apply(null, args)
        })
      } else {
        delete reqs[key]
      }
    }
  })
}

function slice (args) {
  var length = args.length
  var array = []

  for (var i = 0; i < length; i++) array[i] = args[i]
  return array
}


/***/ }),
/* 224 */
/***/ (function(module, exports) {

if (typeof Object.create === 'function') {
  // implementation from standard node.js 'util' module
  module.exports = function inherits(ctor, superCtor) {
    ctor.super_ = superCtor
    ctor.prototype = Object.create(superCtor.prototype, {
      constructor: {
        value: ctor,
        enumerable: false,
        writable: true,
        configurable: true
      }
    });
  };
} else {
  // old school shim for old browsers
  module.exports = function inherits(ctor, superCtor) {
    ctor.super_ = superCtor
    var TempCtor = function () {}
    TempCtor.prototype = superCtor.prototype
    ctor.prototype = new TempCtor()
    ctor.prototype.constructor = ctor
  }
}


/***/ }),
/* 225 */,
/* 226 */,
/* 227 */
/***/ (function(module, exports, __webpack_require__) {

// @flow

/*::
declare var __webpack_require__: mixed;
*/

module.exports = typeof __webpack_require__ !== "undefined";


/***/ }),
/* 228 */,
/* 229 */
/***/ (function(module, exports) {

/**
 * Helpers.
 */

var s = 1000;
var m = s * 60;
var h = m * 60;
var d = h * 24;
var y = d * 365.25;

/**
 * Parse or format the given `val`.
 *
 * Options:
 *
 *  - `long` verbose formatting [false]
 *
 * @param {String|Number} val
 * @param {Object} [options]
 * @throws {Error} throw an error if val is not a non-empty string or a number
 * @return {String|Number}
 * @api public
 */

module.exports = function(val, options) {
  options = options || {};
  var type = typeof val;
  if (type === 'string' && val.length > 0) {
    return parse(val);
  } else if (type === 'number' && isNaN(val) === false) {
    return options.long ? fmtLong(val) : fmtShort(val);
  }
  throw new Error(
    'val is not a non-empty string or a valid number. val=' +
      JSON.stringify(val)
  );
};

/**
 * Parse the given `str` and return milliseconds.
 *
 * @param {String} str
 * @return {Number}
 * @api private
 */

function parse(str) {
  str = String(str);
  if (str.length > 100) {
    return;
  }
  var match = /^((?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|years?|yrs?|y)?$/i.exec(
    str
  );
  if (!match) {
    return;
  }
  var n = parseFloat(match[1]);
  var type = (match[2] || 'ms').toLowerCase();
  switch (type) {
    case 'years':
    case 'year':
    case 'yrs':
    case 'yr':
    case 'y':
      return n * y;
    case 'days':
    case 'day':
    case 'd':
      return n * d;
    case 'hours':
    case 'hour':
    case 'hrs':
    case 'hr':
    case 'h':
      return n * h;
    case 'minutes':
    case 'minute':
    case 'mins':
    case 'min':
    case 'm':
      return n * m;
    case 'seconds':
    case 'second':
    case 'secs':
    case 'sec':
    case 's':
      return n * s;
    case 'milliseconds':
    case 'millisecond':
    case 'msecs':
    case 'msec':
    case 'ms':
      return n;
    default:
      return undefined;
  }
}

/**
 * Short format for `ms`.
 *
 * @param {Number} ms
 * @return {String}
 * @api private
 */

function fmtShort(ms) {
  if (ms >= d) {
    return Math.round(ms / d) + 'd';
  }
  if (ms >= h) {
    return Math.round(ms / h) + 'h';
  }
  if (ms >= m) {
    return Math.round(ms / m) + 'm';
  }
  if (ms >= s) {
    return Math.round(ms / s) + 's';
  }
  return ms + 'ms';
}

/**
 * Long format for `ms`.
 *
 * @param {Number} ms
 * @return {String}
 * @api private
 */

function fmtLong(ms) {
  return plural(ms, d, 'day') ||
    plural(ms, h, 'hour') ||
    plural(ms, m, 'minute') ||
    plural(ms, s, 'second') ||
    ms + ' ms';
}

/**
 * Pluralization helper.
 */

function plural(ms, n, name) {
  if (ms < n) {
    return;
  }
  if (ms < n * 1.5) {
    return Math.floor(ms / n) + ' ' + name;
  }
  return Math.ceil(ms / n) + ' ' + name + 's';
}


/***/ }),
/* 230 */,
/* 231 */,
/* 232 */,
/* 233 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = rimraf
rimraf.sync = rimrafSync

var assert = __webpack_require__(22)
var path = __webpack_require__(0)
var fs = __webpack_require__(3)
var glob = __webpack_require__(75)
var _0666 = parseInt('666', 8)

var defaultGlobOpts = {
  nosort: true,
  silent: true
}

// for EMFILE handling
var timeout = 0

var isWindows = (process.platform === "win32")

function defaults (options) {
  var methods = [
    'unlink',
    'chmod',
    'stat',
    'lstat',
    'rmdir',
    'readdir'
  ]
  methods.forEach(function(m) {
    options[m] = options[m] || fs[m]
    m = m + 'Sync'
    options[m] = options[m] || fs[m]
  })

  options.maxBusyTries = options.maxBusyTries || 3
  options.emfileWait = options.emfileWait || 1000
  if (options.glob === false) {
    options.disableGlob = true
  }
  options.disableGlob = options.disableGlob || false
  options.glob = options.glob || defaultGlobOpts
}

function rimraf (p, options, cb) {
  if (typeof options === 'function') {
    cb = options
    options = {}
  }

  assert(p, 'rimraf: missing path')
  assert.equal(typeof p, 'string', 'rimraf: path should be a string')
  assert.equal(typeof cb, 'function', 'rimraf: callback function required')
  assert(options, 'rimraf: invalid options argument provided')
  assert.equal(typeof options, 'object', 'rimraf: options should be object')

  defaults(options)

  var busyTries = 0
  var errState = null
  var n = 0

  if (options.disableGlob || !glob.hasMagic(p))
    return afterGlob(null, [p])

  options.lstat(p, function (er, stat) {
    if (!er)
      return afterGlob(null, [p])

    glob(p, options.glob, afterGlob)
  })

  function next (er) {
    errState = errState || er
    if (--n === 0)
      cb(errState)
  }

  function afterGlob (er, results) {
    if (er)
      return cb(er)

    n = results.length
    if (n === 0)
      return cb()

    results.forEach(function (p) {
      rimraf_(p, options, function CB (er) {
        if (er) {
          if ((er.code === "EBUSY" || er.code === "ENOTEMPTY" || er.code === "EPERM") &&
              busyTries < options.maxBusyTries) {
            busyTries ++
            var time = busyTries * 100
            // try again, with the same exact callback as this one.
            return setTimeout(function () {
              rimraf_(p, options, CB)
            }, time)
          }

          // this one won't happen if graceful-fs is used.
          if (er.code === "EMFILE" && timeout < options.emfileWait) {
            return setTimeout(function () {
              rimraf_(p, options, CB)
            }, timeout ++)
          }

          // already gone
          if (er.code === "ENOENT") er = null
        }

        timeout = 0
        next(er)
      })
    })
  }
}

// Two possible strategies.
// 1. Assume it's a file.  unlink it, then do the dir stuff on EPERM or EISDIR
// 2. Assume it's a directory.  readdir, then do the file stuff on ENOTDIR
//
// Both result in an extra syscall when you guess wrong.  However, there
// are likely far more normal files in the world than directories.  This
// is based on the assumption that a the average number of files per
// directory is >= 1.
//
// If anyone ever complains about this, then I guess the strategy could
// be made configurable somehow.  But until then, YAGNI.
function rimraf_ (p, options, cb) {
  assert(p)
  assert(options)
  assert(typeof cb === 'function')

  // sunos lets the root user unlink directories, which is... weird.
  // so we have to lstat here and make sure it's not a dir.
  options.lstat(p, function (er, st) {
    if (er && er.code === "ENOENT")
      return cb(null)

    // Windows can EPERM on stat.  Life is suffering.
    if (er && er.code === "EPERM" && isWindows)
      fixWinEPERM(p, options, er, cb)

    if (st && st.isDirectory())
      return rmdir(p, options, er, cb)

    options.unlink(p, function (er) {
      if (er) {
        if (er.code === "ENOENT")
          return cb(null)
        if (er.code === "EPERM")
          return (isWindows)
            ? fixWinEPERM(p, options, er, cb)
            : rmdir(p, options, er, cb)
        if (er.code === "EISDIR")
          return rmdir(p, options, er, cb)
      }
      return cb(er)
    })
  })
}

function fixWinEPERM (p, options, er, cb) {
  assert(p)
  assert(options)
  assert(typeof cb === 'function')
  if (er)
    assert(er instanceof Error)

  options.chmod(p, _0666, function (er2) {
    if (er2)
      cb(er2.code === "ENOENT" ? null : er)
    else
      options.stat(p, function(er3, stats) {
        if (er3)
          cb(er3.code === "ENOENT" ? null : er)
        else if (stats.isDirectory())
          rmdir(p, options, er, cb)
        else
          options.unlink(p, cb)
      })
  })
}

function fixWinEPERMSync (p, options, er) {
  assert(p)
  assert(options)
  if (er)
    assert(er instanceof Error)

  try {
    options.chmodSync(p, _0666)
  } catch (er2) {
    if (er2.code === "ENOENT")
      return
    else
      throw er
  }

  try {
    var stats = options.statSync(p)
  } catch (er3) {
    if (er3.code === "ENOENT")
      return
    else
      throw er
  }

  if (stats.isDirectory())
    rmdirSync(p, options, er)
  else
    options.unlinkSync(p)
}

function rmdir (p, options, originalEr, cb) {
  assert(p)
  assert(options)
  if (originalEr)
    assert(originalEr instanceof Error)
  assert(typeof cb === 'function')

  // try to rmdir first, and only readdir on ENOTEMPTY or EEXIST (SunOS)
  // if we guessed wrong, and it's not a directory, then
  // raise the original error.
  options.rmdir(p, function (er) {
    if (er && (er.code === "ENOTEMPTY" || er.code === "EEXIST" || er.code === "EPERM"))
      rmkids(p, options, cb)
    else if (er && er.code === "ENOTDIR")
      cb(originalEr)
    else
      cb(er)
  })
}

function rmkids(p, options, cb) {
  assert(p)
  assert(options)
  assert(typeof cb === 'function')

  options.readdir(p, function (er, files) {
    if (er)
      return cb(er)
    var n = files.length
    if (n === 0)
      return options.rmdir(p, cb)
    var errState
    files.forEach(function (f) {
      rimraf(path.join(p, f), options, function (er) {
        if (errState)
          return
        if (er)
          return cb(errState = er)
        if (--n === 0)
          options.rmdir(p, cb)
      })
    })
  })
}

// this looks simpler, and is strictly *faster*, but will
// tie up the JavaScript thread and fail on excessively
// deep directory trees.
function rimrafSync (p, options) {
  options = options || {}
  defaults(options)

  assert(p, 'rimraf: missing path')
  assert.equal(typeof p, 'string', 'rimraf: path should be a string')
  assert(options, 'rimraf: missing options')
  assert.equal(typeof options, 'object', 'rimraf: options should be object')

  var results

  if (options.disableGlob || !glob.hasMagic(p)) {
    results = [p]
  } else {
    try {
      options.lstatSync(p)
      results = [p]
    } catch (er) {
      results = glob.sync(p, options.glob)
    }
  }

  if (!results.length)
    return

  for (var i = 0; i < results.length; i++) {
    var p = results[i]

    try {
      var st = options.lstatSync(p)
    } catch (er) {
      if (er.code === "ENOENT")
        return

      // Windows can EPERM on stat.  Life is suffering.
      if (er.code === "EPERM" && isWindows)
        fixWinEPERMSync(p, options, er)
    }

    try {
      // sunos lets the root user unlink directories, which is... weird.
      if (st && st.isDirectory())
        rmdirSync(p, options, null)
      else
        options.unlinkSync(p)
    } catch (er) {
      if (er.code === "ENOENT")
        return
      if (er.code === "EPERM")
        return isWindows ? fixWinEPERMSync(p, options, er) : rmdirSync(p, options, er)
      if (er.code !== "EISDIR")
        throw er

      rmdirSync(p, options, er)
    }
  }
}

function rmdirSync (p, options, originalEr) {
  assert(p)
  assert(options)
  if (originalEr)
    assert(originalEr instanceof Error)

  try {
    options.rmdirSync(p)
  } catch (er) {
    if (er.code === "ENOENT")
      return
    if (er.code === "ENOTDIR")
      throw originalEr
    if (er.code === "ENOTEMPTY" || er.code === "EEXIST" || er.code === "EPERM")
      rmkidsSync(p, options)
  }
}

function rmkidsSync (p, options) {
  assert(p)
  assert(options)
  options.readdirSync(p).forEach(function (f) {
    rimrafSync(path.join(p, f), options)
  })

  // We only end up here once we got ENOTEMPTY at least once, and
  // at this point, we are guaranteed to have removed all the kids.
  // So, we know that it won't be ENOENT or ENOTDIR or anything else.
  // try really hard to delete stuff on windows, because it has a
  // PROFOUNDLY annoying habit of not closing handles promptly when
  // files are deleted, resulting in spurious ENOTEMPTY errors.
  var retries = isWindows ? 100 : 1
  var i = 0
  do {
    var threw = true
    try {
      var ret = options.rmdirSync(p, options)
      threw = false
      return ret
    } finally {
      if (++i < retries && threw)
        continue
    }
  } while (true)
}


/***/ }),
/* 234 */,
/* 235 */,
/* 236 */,
/* 237 */,
/* 238 */,
/* 239 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";

var hasFlag = __webpack_require__(221);

var support = function (level) {
	if (level === 0) {
		return false;
	}

	return {
		level: level,
		hasBasic: true,
		has256: level >= 2,
		has16m: level >= 3
	};
};

var supportLevel = (function () {
	if (hasFlag('no-color') ||
		hasFlag('no-colors') ||
		hasFlag('color=false')) {
		return 0;
	}

	if (hasFlag('color=16m') ||
		hasFlag('color=full') ||
		hasFlag('color=truecolor')) {
		return 3;
	}

	if (hasFlag('color=256')) {
		return 2;
	}

	if (hasFlag('color') ||
		hasFlag('colors') ||
		hasFlag('color=true') ||
		hasFlag('color=always')) {
		return 1;
	}

	if (process.stdout && !process.stdout.isTTY) {
		return 0;
	}

	if (process.platform === 'win32') {
		return 1;
	}

	if ('CI' in process.env) {
		if ('TRAVIS' in process.env || process.env.CI === 'Travis') {
			return 1;
		}

		return 0;
	}

	if ('TEAMCITY_VERSION' in process.env) {
		return process.env.TEAMCITY_VERSION.match(/^(9\.(0*[1-9]\d*)\.|\d{2,}\.)/) === null ? 0 : 1;
	}

	if (/^(screen|xterm)-256(?:color)?/.test(process.env.TERM)) {
		return 2;
	}

	if (/^screen|^xterm|^vt100|color|ansi|cygwin|linux/i.test(process.env.TERM)) {
		return 1;
	}

	if ('COLORTERM' in process.env) {
		return 1;
	}

	if (process.env.TERM === 'dumb') {
		return 0;
	}

	return 0;
})();

if (supportLevel === 0 && 'FORCE_COLOR' in process.env) {
	supportLevel = 1;
}

module.exports = process && support(supportLevel);


/***/ })
/******/ ]);