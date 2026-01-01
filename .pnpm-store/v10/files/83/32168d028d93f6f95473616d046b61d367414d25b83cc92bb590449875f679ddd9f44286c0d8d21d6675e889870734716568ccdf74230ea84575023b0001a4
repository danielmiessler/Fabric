"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const shared_1 = require("../shared");
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------
const COMMENT_DIRECTIVE_B = /^\s*(eslint-(?:en|dis)able)(?:\s+|$)/;
const COMMENT_DIRECTIVE_L = /^\s*(eslint-disable(?:-next)?-line)(?:\s+|$)/;
// eslint-disable-next-line func-style -- ignore
const ALL_RULES = () => true;
/**
 * Remove the ignored part from a given directive comment and trim it.
 * @param {string} value The comment text to strip.
 * @returns {string} The stripped text.
 */
function stripDirectiveComment(value) {
    return value.split(/\s-{2,}\s/u)[0];
}
exports.default = (0, utils_1.createRule)('comment-directive', {
    meta: {
        docs: {
            description: 'support comment-directives in HTML template',
            category: 'System',
            recommended: 'base'
        },
        schema: [
            {
                type: 'object',
                properties: {
                    reportUnusedDisableDirectives: {
                        type: 'boolean'
                    }
                },
                additionalProperties: false
            }
        ],
        messages: {
            unused: 'Unused {{kind}} directive (no problems were reported).',
            unusedRule: "Unused {{kind}} directive (no problems were reported from '{{rule}}').",
            unusedEnable: 'Unused {{kind}} directive (reporting is not suppressed).',
            unusedEnableRule: "Unused {{kind}} directive (reporting from '{{rule}}' is not suppressed)."
        },
        type: 'problem'
    },
    create(context) {
        const shared = (0, shared_1.getShared)((0, compat_1.getFilename)(context));
        if (!shared)
            return {};
        const options = context.options[0] || {};
        const reportUnusedDisableDirectives = Boolean(options.reportUnusedDisableDirectives);
        const directives = shared.newCommentDirectives({
            ruleId: 'svelte/comment-directive',
            reportUnusedDisableDirectives
        });
        const sourceCode = (0, compat_1.getSourceCode)(context);
        /**
         * Parse a given comment.
         */
        function parse(pattern, comment) {
            const text = stripDirectiveComment(comment.value);
            const match = pattern.exec(text);
            if (match == null) {
                return null;
            }
            const type = match[1];
            const rules = [];
            const rulesRe = /([^\s,]+)[\s,]*/g;
            let startIndex = match[0].length;
            rulesRe.lastIndex = startIndex;
            let res;
            while ((res = rulesRe.exec(text))) {
                const ruleId = res[1].trim();
                const commentStart = comment.range[0] + 4; /* <!-- */
                const start = sourceCode.getLocFromIndex(commentStart + startIndex);
                const end = sourceCode.getLocFromIndex(commentStart + startIndex + ruleId.length);
                rules.push({
                    ruleId,
                    loc: {
                        start,
                        end
                    }
                });
                startIndex = rulesRe.lastIndex;
            }
            return { type, rules };
        }
        /**
         * Process a given comment token.
         * If the comment is `eslint-disable` or `eslint-enable` then it reports the comment.
         */
        function processBlock(directives, comment) {
            const parsed = parse(COMMENT_DIRECTIVE_B, comment);
            if (parsed != null) {
                if (parsed.type === 'eslint-disable') {
                    if (parsed.rules.length) {
                        for (const rule of parsed.rules) {
                            if (reportUnusedDisableDirectives) {
                                context.report({
                                    loc: rule.loc,
                                    messageId: 'unusedRule',
                                    data: { rule: rule.ruleId, kind: parsed.type }
                                });
                            }
                            directives.disableBlock(comment.loc.end, rule.ruleId, {
                                loc: rule.loc.start
                            });
                        }
                    }
                    else {
                        if (reportUnusedDisableDirectives) {
                            context.report({
                                loc: comment.loc,
                                messageId: 'unused',
                                data: { kind: parsed.type }
                            });
                        }
                        directives.disableBlock(comment.loc.end, ALL_RULES, {
                            loc: comment.loc.start
                        });
                    }
                }
                else if (parsed.rules.length) {
                    for (const rule of parsed.rules) {
                        if (reportUnusedDisableDirectives) {
                            context.report({
                                loc: rule.loc,
                                messageId: 'unusedEnableRule',
                                data: { rule: rule.ruleId, kind: parsed.type }
                            });
                        }
                        directives.enableBlock(comment.loc.start, rule.ruleId, {
                            loc: rule.loc.start
                        });
                    }
                }
                else {
                    if (reportUnusedDisableDirectives) {
                        context.report({
                            loc: comment.loc,
                            messageId: 'unusedEnable',
                            data: { kind: parsed.type }
                        });
                    }
                    directives.enableBlock(comment.loc.start, ALL_RULES, {
                        loc: comment.loc.start
                    });
                }
            }
        }
        /**
         * Process a given comment token.
         * If the comment is `eslint-disable-line` or `eslint-disable-next-line` then it reports the comment.
         */
        function processLine(directives, comment) {
            const parsed = parse(COMMENT_DIRECTIVE_L, comment);
            if (parsed != null && comment.loc.start.line === comment.loc.end.line) {
                const line = comment.loc.start.line + (parsed.type === 'eslint-disable-line' ? 0 : 1);
                if (parsed.rules.length) {
                    for (const rule of parsed.rules) {
                        if (reportUnusedDisableDirectives) {
                            context.report({
                                loc: rule.loc,
                                messageId: 'unusedRule',
                                data: { rule: rule.ruleId, kind: parsed.type }
                            });
                        }
                        directives.disableLine(line, rule.ruleId, {
                            loc: rule.loc.start
                        });
                    }
                }
                else {
                    if (reportUnusedDisableDirectives) {
                        context.report({
                            loc: comment.loc,
                            messageId: 'unused',
                            data: { kind: parsed.type }
                        });
                    }
                    directives.disableLine(line, ALL_RULES, {
                        loc: comment.loc.start
                    });
                }
            }
        }
        return {
            SvelteHTMLComment(node) {
                processBlock(directives, node);
                processLine(directives, node);
            },
            SvelteScriptElement(node) {
                directives.enableBlock(node.startTag.loc.end, ALL_RULES, {
                    loc: node.loc.start
                });
            }
        };
    }
});
