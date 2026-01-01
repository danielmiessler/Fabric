/**
 * Check if a match position is within a comment or a string
 * @param {string} content - The full content
 * @param {number} match_index - The index where the match starts
 * @returns {boolean} - True if the match is within a comment
 */
export function should_ignore(content, match_index) {
	// Track if we're inside different types of quotes and comments
	let in_single_quote = false;
	let in_double_quote = false;
	let in_template_literal = false;
	let in_single_line_comment = false;
	let in_multi_line_comment = false;
	let in_html_comment = false;

	for (let i = 0; i < match_index; i++) {
		const char = content[i];
		const next_two = content.slice(i, i + 2);
		const next_four = content.slice(i, i + 4);

		// Handle end of single line comment
		if (in_single_line_comment && char === '\n') {
			in_single_line_comment = false;
			continue;
		}

		// Handle end of multi-line comment
		if (in_multi_line_comment && next_two === '*/') {
			in_multi_line_comment = false;
			i++; // Skip the '/' part
			continue;
		}

		// Handle end of HTML comment
		if (in_html_comment && content.slice(i, i + 3) === '-->') {
			in_html_comment = false;
			i += 2; // Skip the '-->' part
			continue;
		}

		// If we're in any comment, skip processing
		if (in_single_line_comment || in_multi_line_comment || in_html_comment) {
			continue;
		}

		// Handle escape sequences in strings
		if ((in_single_quote || in_double_quote || in_template_literal) && char === '\\') {
			i++; // Skip the escaped character
			continue;
		}

		// Handle string boundaries
		if (!in_double_quote && !in_template_literal && char === "'") {
			in_single_quote = !in_single_quote;
			continue;
		}

		if (!in_single_quote && !in_template_literal && char === '"') {
			in_double_quote = !in_double_quote;
			continue;
		}

		if (!in_single_quote && !in_double_quote && char === '`') {
			in_template_literal = !in_template_literal;
			continue;
		}

		// If we're inside any string, don't process comment delimiters
		if (in_single_quote || in_double_quote || in_template_literal) {
			continue;
		}

		// Check for comment starts
		if (next_two === '//') {
			in_single_line_comment = true;
			i++; // Skip the second '/'
			continue;
		}

		if (next_two === '/*') {
			in_multi_line_comment = true;
			i++; // Skip the '*'
			continue;
		}

		if (next_four === '<!--') {
			in_html_comment = true;
			i += 3; // Skip the '<!--'
			continue;
		}
	}

	return (
		in_single_line_comment ||
		in_multi_line_comment ||
		in_html_comment ||
		in_single_quote ||
		in_double_quote ||
		in_template_literal
	);
}
