let FORCE_COLOR, NODE_DISABLE_COLORS, NO_COLOR, TERM, isTTY=true;
if (typeof process !== 'undefined') {
	({ FORCE_COLOR, NODE_DISABLE_COLORS, NO_COLOR, TERM } = process.env || {});
	isTTY = process.stdout && process.stdout.isTTY;
}

const $ = exports.$ = {
	enabled: !NODE_DISABLE_COLORS && NO_COLOR == null && TERM !== 'dumb' && (
		FORCE_COLOR != null && FORCE_COLOR !== '0' || isTTY
	)
}

function init(x, y) {
	let rgx = new RegExp(`\\x1b\\[${y}m`, 'g');
	let open = `\x1b[${x}m`, close = `\x1b[${y}m`;

	return function (txt) {
		if (!$.enabled || txt == null) return txt;
		return open + (!!~(''+txt).indexOf(close) ? txt.replace(rgx, close + open) : txt) + close;
	};
}

// modifiers
exports.reset = init(0, 0);
exports.bold = init(1, 22);
exports.dim = init(2, 22);
exports.italic = init(3, 23);
exports.underline = init(4, 24);
exports.inverse = init(7, 27);
exports.hidden = init(8, 28);
exports.strikethrough = init(9, 29);

// colors
exports.black = init(30, 39);
exports.red = init(31, 39);
exports.green = init(32, 39);
exports.yellow = init(33, 39);
exports.blue = init(34, 39);
exports.magenta = init(35, 39);
exports.cyan = init(36, 39);
exports.white = init(37, 39);
exports.gray = init(90, 39);
exports.grey = init(90, 39);

// background colors
exports.bgBlack = init(40, 49);
exports.bgRed = init(41, 49);
exports.bgGreen = init(42, 49);
exports.bgYellow = init(43, 49);
exports.bgBlue = init(44, 49);
exports.bgMagenta = init(45, 49);
exports.bgCyan = init(46, 49);
exports.bgWhite = init(47, 49);
