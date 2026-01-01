import { re } from './id.js';

/** @typedef {import('estree').Comment} Comment */
/** @typedef {import('estree').Node} Node */

/**
 * @typedef {Node & {
 *   start: number;
 *   end: number;
 *   has_trailing_newline?: boolean
 * }} NodeWithLocation
 */

/**
 * @typedef {Comment & {
 *   start: number;
 *   end: number;
 *   has_trailing_newline?: boolean
 * }} CommentWithLocation
 */

/**
 * @param {CommentWithLocation[]} comments
 * @param {string} raw
 */
export const get_comment_handlers = (comments, raw) => ({
	// pass to acorn options
	/**
	 * @param {boolean} block
	 * @param {string} value
	 * @param {number} start
	 * @param {number} end
	 */
	onComment: (block, value, start, end) => {
		if (block && /\n/.test(value)) {
			let a = start;
			while (a > 0 && raw[a - 1] !== '\n') a -= 1;

			let b = a;
			while (/[ \t]/.test(raw[b])) b += 1;

			const indentation = raw.slice(a, b);
			value = value.replace(new RegExp(`^${indentation}`, 'gm'), '');
		}

		comments.push({ type: block ? 'Block' : 'Line', value, start, end });
	},

	// pass to estree-walker options
	/** @param {NodeWithLocation} node */
	enter(node) {
		let comment;

		while (comments[0] && comments[0].start < node.start) {
			comment = comments.shift();

			comment.value = comment.value.replace(
				re,
				(match, id, at, hash, value) => {
					if (hash) return `#${value}`;
					if (at) return `@${value}`;

					return match;
				}
			);

			const next = comments[0] || node;
			comment.has_trailing_newline =
				comment.type === 'Line' ||
				/\n/.test(raw.slice(comment.end, next.start));

			(node.leadingComments || (node.leadingComments = [])).push(comment);
		}
	},

	/** @param {NodeWithLocation} node */
	leave(node) {
		if (comments[0]) {
			const slice = raw.slice(node.end, comments[0].start);

			if (/^[,) \t]*$/.test(slice)) {
				node.trailingComments = [comments.shift()];
			}
		}
	}
});
