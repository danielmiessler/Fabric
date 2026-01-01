import { resolve, dirname, basename, relative, extname } from 'path';
import { Promise as Promise$1, readFileSync, readFile, writeFile, writeFileSync } from 'sander';

const comma = ','.charCodeAt(0);
const semicolon = ';'.charCodeAt(0);
const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';
const intToChar = new Uint8Array(64); // 64 possible chars.
const charToInt = new Uint8Array(128); // z is 122 in ASCII
for (let i = 0; i < chars.length; i++) {
    const c = chars.charCodeAt(i);
    intToChar[i] = c;
    charToInt[c] = i;
}
// Provide a fallback for older environments.
const td = typeof TextDecoder !== 'undefined'
    ? /* #__PURE__ */ new TextDecoder()
    : typeof Buffer !== 'undefined'
        ? {
            decode(buf) {
                const out = Buffer.from(buf.buffer, buf.byteOffset, buf.byteLength);
                return out.toString();
            },
        }
        : {
            decode(buf) {
                let out = '';
                for (let i = 0; i < buf.length; i++) {
                    out += String.fromCharCode(buf[i]);
                }
                return out;
            },
        };
function decode(mappings) {
    const state = new Int32Array(5);
    const decoded = [];
    let index = 0;
    do {
        const semi = indexOf(mappings, index);
        const line = [];
        let sorted = true;
        let lastCol = 0;
        state[0] = 0;
        for (let i = index; i < semi; i++) {
            let seg;
            i = decodeInteger(mappings, i, state, 0); // genColumn
            const col = state[0];
            if (col < lastCol)
                sorted = false;
            lastCol = col;
            if (hasMoreVlq(mappings, i, semi)) {
                i = decodeInteger(mappings, i, state, 1); // sourcesIndex
                i = decodeInteger(mappings, i, state, 2); // sourceLine
                i = decodeInteger(mappings, i, state, 3); // sourceColumn
                if (hasMoreVlq(mappings, i, semi)) {
                    i = decodeInteger(mappings, i, state, 4); // namesIndex
                    seg = [col, state[1], state[2], state[3], state[4]];
                }
                else {
                    seg = [col, state[1], state[2], state[3]];
                }
            }
            else {
                seg = [col];
            }
            line.push(seg);
        }
        if (!sorted)
            sort(line);
        decoded.push(line);
        index = semi + 1;
    } while (index <= mappings.length);
    return decoded;
}
function indexOf(mappings, index) {
    const idx = mappings.indexOf(';', index);
    return idx === -1 ? mappings.length : idx;
}
function decodeInteger(mappings, pos, state, j) {
    let value = 0;
    let shift = 0;
    let integer = 0;
    do {
        const c = mappings.charCodeAt(pos++);
        integer = charToInt[c];
        value |= (integer & 31) << shift;
        shift += 5;
    } while (integer & 32);
    const shouldNegate = value & 1;
    value >>>= 1;
    if (shouldNegate) {
        value = -0x80000000 | -value;
    }
    state[j] += value;
    return pos;
}
function hasMoreVlq(mappings, i, length) {
    if (i >= length)
        return false;
    return mappings.charCodeAt(i) !== comma;
}
function sort(line) {
    line.sort(sortComparator);
}
function sortComparator(a, b) {
    return a[0] - b[0];
}
function encode(decoded) {
    const state = new Int32Array(5);
    const bufLength = 1024 * 16;
    const subLength = bufLength - 36;
    const buf = new Uint8Array(bufLength);
    const sub = buf.subarray(0, subLength);
    let pos = 0;
    let out = '';
    for (let i = 0; i < decoded.length; i++) {
        const line = decoded[i];
        if (i > 0) {
            if (pos === bufLength) {
                out += td.decode(buf);
                pos = 0;
            }
            buf[pos++] = semicolon;
        }
        if (line.length === 0)
            continue;
        state[0] = 0;
        for (let j = 0; j < line.length; j++) {
            const segment = line[j];
            // We can push up to 5 ints, each int can take at most 7 chars, and we
            // may push a comma.
            if (pos > subLength) {
                out += td.decode(sub);
                buf.copyWithin(0, subLength, pos);
                pos -= subLength;
            }
            if (j > 0)
                buf[pos++] = comma;
            pos = encodeInteger(buf, pos, state, segment, 0); // genColumn
            if (segment.length === 1)
                continue;
            pos = encodeInteger(buf, pos, state, segment, 1); // sourcesIndex
            pos = encodeInteger(buf, pos, state, segment, 2); // sourceLine
            pos = encodeInteger(buf, pos, state, segment, 3); // sourceColumn
            if (segment.length === 4)
                continue;
            pos = encodeInteger(buf, pos, state, segment, 4); // namesIndex
        }
    }
    return out + td.decode(buf.subarray(0, pos));
}
function encodeInteger(buf, pos, state, segment, j) {
    const next = segment[j];
    let num = next - state[j];
    state[j] = next;
    num = num < 0 ? (-num << 1) | 1 : num << 1;
    do {
        let clamped = num & 0b011111;
        num >>>= 5;
        if (num > 0)
            clamped |= 0b100000;
        buf[pos++] = intToChar[clamped];
    } while (num > 0);
    return pos;
}

/**
 * Decodes a base64 string
 * @param {string} base64 - the string to decode
 * @returns {string}
 */
function atob ( base64 ) {
	return new Buffer( base64, 'base64' ).toString( 'utf8' );
}

// this looks ridiculous, but it prevents sourcemap tooling from mistaking
// this for an actual sourceMappingURL
var SOURCEMAPPING_URL = 'sourceMa';
SOURCEMAPPING_URL += 'ppingURL';

var SOURCEMAPPING_URL$1 = SOURCEMAPPING_URL;

function parseJSON ( json, url ) {
	try {
		return JSON.parse( json );
	} catch ( err ) {
		throw new Error( ("Could not parse sourcemap (" + url + "): " + (err.message)) );
	}
}

/**
 * Turns a sourceMappingURL into a sourcemap
 * @param {string} url - the sourceMappingURL. Can be a
   base64-encoded data URI
 * @param {string} base - the URL against which relative URLS
   should be resolved
 * @param {boolean} sync - if `true`, return a promise, otherwise
   return the sourcemap
 * @returns {object} - a version 3 sourcemap
 */
function getMapFromUrl ( url, base, sync ) {
	if ( /^data:/.test( url ) ) { // TODO beef this up
		var match = /base64,(.+)$/.exec( url );

		if ( !match ) {
			throw new Error( (SOURCEMAPPING_URL$1 + " is not base64-encoded") );
		}

		var json = atob( match[1] );
		var map = parseJSON( json, ("data URI in " + base) );
		return sync ? map : Promise$1.resolve( map );
	}

	url = resolve( dirname( base ), decodeURI( url ) );

	if ( sync ) {
		return parseJSON( readFileSync( url, { encoding: 'utf-8' }), url );
	} else {
		return readFile( url, { encoding: 'utf-8' }).then( function (json) { return parseJSON( json, url ); } );
	}
}

function getSourceMappingUrl ( str ) {
	var index, substring, url, match;

	// assume we want the last occurence
	index = str.lastIndexOf( (SOURCEMAPPING_URL$1 + "=") );

	if ( index === -1 ) {
		return null;
	}

	substring = str.substring( index + 17 );
	match = /^[^\r\n]+/.exec( substring );

	url = match ? match[0] : null;

	// possibly a better way to do this, but we don't want to exclude whitespace
	// from the sourceMappingURL because it might not have been correctly encoded
	if ( url && url.slice( -2 ) === '*/' ) {
		url = url.slice( 0, -2 ).trim();
	}

	return url;
}

function getMap ( node, sourceMapByPath, sync ) {
	if ( node.file in sourceMapByPath ) {
		var map = sourceMapByPath[ node.file ];
		return sync ? map : Promise$1.resolve( map );
	}

	else {
		var url = getSourceMappingUrl( node.content );

		if ( !url ) {
			node.isOriginalSource = true;
			return sync ? null : Promise$1.resolve( null );
		}

		return getMapFromUrl( url, node.file, sync );
	}
}

function Node (ref) {
	var file = ref.file;
	var content = ref.content;

	this.file = file ? resolve( file ) : null;
	this.content = content || null; // sometimes exists in sourcesContent, sometimes doesn't

	if ( !this.file && this.content === null ) {
		throw new Error( 'A source must specify either file or content' );
	}

	// these get filled in later
	this.map = null;
	this.mappings = null;
	this.sources = null;
	this.isOriginalSource = null;

	this._stats = {
		decodingTime: 0,
		encodingTime: 0,
		tracingTime: 0,

		untraceable: 0
	};
}

Node.prototype = {
	load: function load ( sourcesContentByPath, sourceMapByPath ) {
		var this$1$1 = this;

		return getContent( this, sourcesContentByPath ).then( function (content) {
			this$1$1.content = sourcesContentByPath[ this$1$1.file ] = content;

			return getMap( this$1$1, sourceMapByPath ).then( function (map) {
				if ( !map ) { return null; }

				this$1$1.map = map;

				var decodingStart = process.hrtime();
				this$1$1.mappings = decode( map.mappings );
				var decodingTime = process.hrtime( decodingStart );
				this$1$1._stats.decodingTime = 1e9 * decodingTime[0] + decodingTime[1];

				var sourcesContent = map.sourcesContent || [];

				var sourceRoot = resolve( dirname( this$1$1.file || '' ), map.sourceRoot || '' );

				this$1$1.sources = map.sources.map( function ( source, i ) {
					return new Node({
						file: source ? resolve( sourceRoot, source ) : null,
						content: sourcesContent[i]
					});
				});

				var promises = this$1$1.sources.map( function (node) { return node.load( sourcesContentByPath, sourceMapByPath ); } );
				return Promise$1.all( promises );
			});
		});
	},

	loadSync: function loadSync ( sourcesContentByPath, sourceMapByPath ) {
		if ( !this.content ) {
			if ( !sourcesContentByPath[ this.file ] ) {
				sourcesContentByPath[ this.file ] = readFileSync( this.file, { encoding: 'utf-8' });
			}

			this.content = sourcesContentByPath[ this.file ];
		}

		var map = getMap( this, sourceMapByPath, true );
		var sourcesContent;

		if ( !map ) {
			this.isOriginalSource = true;
		} else {
			this.map = map;
			this.mappings = decode( map.mappings );

			sourcesContent = map.sourcesContent || [];

			var sourceRoot = resolve( dirname( this.file || '' ), map.sourceRoot || '' );

			this.sources = map.sources.map( function ( source, i ) {
				var node = new Node({
					file: resolve( sourceRoot, source ),
					content: sourcesContent[i]
				});

				node.loadSync( sourcesContentByPath, sourceMapByPath );
				return node;
			});
		}
	},

	/**
	 * Traces a segment back to its origin
	 * @param {number} lineIndex - the zero-based line index of the
	   segment as found in `this`
	 * @param {number} columnIndex - the zero-based column index of the
	   segment as found in `this`
	 * @param {string || null} - if specified, the name that should be
	   (eventually) returned, as it is closest to the generated code
	 * @returns {object}
	     @property {string} source - the filepath of the source
	     @property {number} line - the one-based line index
	     @property {number} column - the zero-based column index
	     @property {string || null} name - the name corresponding
	     to the segment being traced
	 */
	trace: function trace ( lineIndex, columnIndex, name ) {
		// If this node doesn't have a source map, we have
		// to assume it is the original source
		if ( this.isOriginalSource ) {
			return {
				source: this.file,
				line: lineIndex + 1,
				column: columnIndex || 0,
				name: name
			};
		}

		// Otherwise, we need to figure out what this position in
		// the intermediate file corresponds to in *its* source
		var segments = this.mappings[ lineIndex ];

		if ( !segments || segments.length === 0 ) {
			return null;
		}

		if ( columnIndex != null ) {
			var len = segments.length;
			var i;

			for ( i = 0; i < len; i += 1 ) {
				var generatedCodeColumn = segments[i][0];

				if ( generatedCodeColumn > columnIndex ) {
					break;
				}

				if ( generatedCodeColumn === columnIndex ) {
					if ( segments[i].length < 4 ) { return null; }

					var sourceFileIndex$1 = segments[i][1];
					var sourceCodeLine$1 = segments[i][2];
					var sourceCodeColumn = segments[i][3];
					var nameIndex$1 = segments[i][4];

					var parent$1 = this.sources[ sourceFileIndex$1 ];
					return parent$1.trace( sourceCodeLine$1, sourceCodeColumn, this.map.names[ nameIndex$1 ] || name );
				}
			}
		}

		// fall back to a line mapping
		var sourceFileIndex = segments[0][1];
		var sourceCodeLine = segments[0][2];
		var nameIndex = segments[0][4];

		var parent = this.sources[ sourceFileIndex ];
		return parent.trace( sourceCodeLine, null, this.map.names[ nameIndex ] || name );
	}
};

function getContent ( node, sourcesContentByPath ) {
	if ( node.file in sourcesContentByPath ) {
		node.content = sourcesContentByPath[ node.file ];
	}

	if ( !node.content ) {
		return readFile( node.file, { encoding: 'utf-8' });
	}

	return Promise$1.resolve( node.content );
}

/**
 * Encodes a string as base64
 * @param {string} str - the string to encode
 * @returns {string}
 */
function btoa ( str ) {
	return new Buffer( str ).toString( 'base64' );
}

function SourceMap ( properties ) {
	this.version = 3;

	this.file           = properties.file;
	this.sources        = properties.sources;
	this.sourcesContent = properties.sourcesContent;
	this.names          = properties.names;
	this.mappings       = properties.mappings;
}

SourceMap.prototype = {
	toString: function toString () {
		return JSON.stringify( this );
	},

	toUrl: function toUrl () {
		return 'data:application/json;charset=utf-8;base64,' + btoa( this.toString() );
	}
};

function slash ( path ) {
	return typeof path === 'string' ?
		path.replace( /\\/g, '/' ) :
		path;
}

var SOURCEMAP_COMMENT = new RegExp( "\n*(?:" +
	"\\/\\/[@#]\\s*" + SOURCEMAPPING_URL$1 + "=([^\n]+)|" +      // js
	"\\/\\*#?\\s*" + SOURCEMAPPING_URL$1 + "=([^'\"]+)\\s\\*\\/)" + // css
'\\s*$', 'g' );

function Chain ( node, sourcesContentByPath ) {
	this.node = node;
	this.sourcesContentByPath = sourcesContentByPath;

	this._stats = {};
}

Chain.prototype = {
	stat: function stat () {
		return {
			selfDecodingTime: this._stats.decodingTime / 1e6,
			totalDecodingTime: ( this._stats.decodingTime + tally( this.node.sources, 'decodingTime' ) ) / 1e6,

			encodingTime: this._stats.encodingTime / 1e6,
			tracingTime: this._stats.tracingTime / 1e6,

			untraceable: this._stats.untraceable
		};
	},

	apply: function apply ( options ) {
		var this$1$1 = this;
		if ( options === void 0 ) options = {};

		var allNames = [];
		var allSources = [];

		var applySegment = function ( segment, result ) {
			if ( segment.length < 4 ) { return; }

			var traced = this$1$1.node.sources[ segment[1] ].trace( // source
				segment[2], // source code line
				segment[3], // source code column
				this$1$1.node.map.names[ segment[4] ]
			);

			if ( !traced ) {
				this$1$1._stats.untraceable += 1;
				return;
			}

			var sourceIndex = allSources.indexOf( traced.source );
			if ( !~sourceIndex ) {
				sourceIndex = allSources.length;
				allSources.push( traced.source );
			}

			var newSegment = [
				segment[0], // generated code column
				sourceIndex,
				traced.line - 1,
				traced.column
			];

			if ( traced.name ) {
				var nameIndex = allNames.indexOf( traced.name );
				if ( !~nameIndex ) {
					nameIndex = allNames.length;
					allNames.push( traced.name );
				}

				newSegment[4] = nameIndex;
			}

			result[ result.length ] = newSegment;
		};

		// Trace mappings
		var tracingStart = process.hrtime();

		var i = this.node.mappings.length;
		var resolved = new Array( i );

		var j, line, result;

		while ( i-- ) {
			line = this.node.mappings[i];
			resolved[i] = result = [];

			for ( j = 0; j < line.length; j += 1 ) {
				applySegment( line[j], result );
			}
		}

		var tracingTime = process.hrtime( tracingStart );
		this._stats.tracingTime = 1e9 * tracingTime[0] + tracingTime[1];

		// Encode mappings
		var encodingStart = process.hrtime();
		var mappings = encode( resolved );
		var encodingTime = process.hrtime( encodingStart );
		this._stats.encodingTime = 1e9 * encodingTime[0] + encodingTime[1];

		var includeContent = options.includeContent !== false;

		return new SourceMap({
			file: basename( this.node.file ),
			sources: allSources.map( function (source) { return slash( relative( options.base || dirname( this$1$1.node.file ), source ) ); } ),
			sourcesContent: allSources.map( function (source) { return includeContent ? this$1$1.sourcesContentByPath[ source ] : null; } ),
			names: allNames,
			mappings: mappings
		});
	},

	trace: function trace ( oneBasedLineIndex, zeroBasedColumnIndex ) {
		return this.node.trace( oneBasedLineIndex - 1, zeroBasedColumnIndex, null );
	},

	write: function write ( dest, options ) {
		if ( typeof dest !== 'string' ) {
			options = dest;
			dest = this.node.file;
		}

		options = options || {};

		var ref = processWriteOptions( dest, this, options );
		var resolved = ref.resolved;
		var content = ref.content;
		var map = ref.map;

		var promises = [ writeFile( resolved, content ) ];

		if ( !options.inline ) {
			promises.push( writeFile( resolved + '.map', map.toString() ) );
		}

		return Promise.all( promises );
	},

	writeSync: function writeSync ( dest, options ) {
		if ( typeof dest !== 'string' ) {
			options = dest;
			dest = this.node.file;
		}

		options = options || {};

		var ref = processWriteOptions( dest, this, options );
		var resolved = ref.resolved;
		var content = ref.content;
		var map = ref.map;

		writeFileSync( resolved, content );

		if ( !options.inline ) {
			writeFileSync( resolved + '.map', map.toString() );
		}
	}
};

function processWriteOptions ( dest, chain, options ) {
	var resolved = resolve( dest );

	var map = chain.apply({
		includeContent: options.includeContent,
		base: options.base ? resolve( options.base ) : dirname( resolved )
	});

	var url = options.inline ? map.toUrl() : ( options.absolutePath ? resolved : basename( resolved ) ) + '.map';

	// TODO shouldn't url be relative?
	var content = chain.node.content.replace( SOURCEMAP_COMMENT, '' ) + sourcemapComment( url, resolved );

	return { resolved: resolved, content: content, map: map };
}

function tally ( nodes, stat ) {
	return nodes.reduce( function ( total, node ) {
		return total + node._stats[ stat ];
	}, 0 );
}

function sourcemapComment ( url, dest ) {
	var ext = extname( dest );
	url = encodeURI( url );

	if ( ext === '.css' ) {
		return ("\n/*# " + SOURCEMAPPING_URL$1 + "=" + url + " */\n");
	}

	return ("\n//# " + SOURCEMAPPING_URL$1 + "=" + url + "\n");
}

function load ( file, options ) {
	var ref = init( file, options );
	var node = ref.node;
	var sourcesContentByPath = ref.sourcesContentByPath;
	var sourceMapByPath = ref.sourceMapByPath;

	return node.load( sourcesContentByPath, sourceMapByPath )
		.then( function () { return node.isOriginalSource ? null : new Chain( node, sourcesContentByPath ); } );
}

function loadSync ( file, options ) {
	if ( options === void 0 ) options = {};

	var ref = init( file, options );
	var node = ref.node;
	var sourcesContentByPath = ref.sourcesContentByPath;
	var sourceMapByPath = ref.sourceMapByPath;

	node.loadSync( sourcesContentByPath, sourceMapByPath );
	return node.isOriginalSource ? null : new Chain( node, sourcesContentByPath );
}

function init ( file, options ) {
	if ( options === void 0 ) options = {};

	var node = new Node({ file: file });

	var sourcesContentByPath = {};
	var sourceMapByPath = {};

	if ( options.content ) {
		Object.keys( options.content ).forEach( function (key) {
			sourcesContentByPath[ resolve( key ) ] = options.content[ key ];
		});
	}

	if ( options.sourcemaps ) {
		Object.keys( options.sourcemaps ).forEach( function (key) {
			sourceMapByPath[ resolve( key ) ] = options.sourcemaps[ key ];
		});
	}

	return { node: node, sourcesContentByPath: sourcesContentByPath, sourceMapByPath: sourceMapByPath };
}

export { load, loadSync };
