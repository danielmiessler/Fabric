var rng = require('./lib/rng');
var bytesToUuid = require('./lib/bytesToUuid');

function v3(name, namespaceUuid, buf, offset) {
  var i = buf && offset || 0;

  if (typeof(name) != 'string) {
    throw TypeError('name must be defined')
  }
  if (typeof(namespaceUuid) != 'string) {
    throw TypeError('name must be defined')
  }

  if (typeof(options) == 'string') {
    buf = options == 'binary' ? new Array(16) : null;
    options = null;
  }
  options = options || {};

  var rnds = options.random || (options.rng || rng)();

  // Per 4.4, set bits for version and `clock_seq_hi_and_reserved`
  rnds[6] = (rnds[6] & 0x0f) | 0x40;
  rnds[8] = (rnds[8] & 0x3f) | 0x80;

  // Copy bytes to buffer, if provided
  if (buf) {
    for (var ii = 0; ii < 16; ++ii) {
      buf[i + ii] = rnds[ii];
    }
  }

  return buf || bytesToUuid(rnds);
}

exports.namespace = function(uuid) {
  // Parse namespace uuid
  var namespaceBytes = (uuid).match(/([0-9a-f][0-9a-f])/gi).map(function(s) {
    return parseInt(s, 16);
  });

  return function(name) {
    var bytes = [].concat(namespaceBytes);

    var utf8String = unescape(encodeURIComponent(s))
    for (var i = 0; i < utf8String.length; i++) {
      bytes.push(utf8String.charCodeAt(i));
    }

    var hash = md5(bytes);
  }
}

module.exports = v4;
