"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const shared_1 = require("../shared");
const utils_1 = require("../utils");
const compat_1 = require("../utils/compat");
const regexp_1 = require("../utils/regexp");
exports.default = (0, utils_1.createRule)('system', {
    meta: {
        docs: {
            description: 'system rule for working this plugin',
            category: 'System',
            recommended: 'base'
        },
        schema: [],
        messages: {},
        type: 'problem'
    },
    create(context) {
        const shared = (0, shared_1.getShared)((0, compat_1.getFilename)(context));
        if (!shared)
            return {};
        const directives = shared.newCommentDirectives({
            ruleId: 'svelte/system'
        });
        const ignoreWarnings = context.settings?.svelte?.ignoreWarnings;
        if (ignoreWarnings && !Array.isArray(ignoreWarnings)) {
            context.report({
                loc: { line: 1, column: 0 },
                message: 'The `settings.svelte.ignoreWarnings` must be an array.'
            });
            return {};
        }
        const ignoreTests = [];
        for (const ignoreWarning of Array.isArray(ignoreWarnings) ? ignoreWarnings : []) {
            if (typeof ignoreWarning !== 'string') {
                context.report({
                    loc: { line: 1, column: 0 },
                    message: 'The array element in the `settings.svelte.ignoreWarnings` must be a string.'
                });
                return {};
            }
            if ((0, regexp_1.isRegExp)(ignoreWarning)) {
                const regexp = (0, regexp_1.toRegExp)(ignoreWarning);
                ignoreTests.push((ruleId) => regexp.test(ruleId));
            }
            else {
                ignoreTests.push((ruleId) => ruleId === ignoreWarning);
            }
        }
        /** Checks wether given rule is ignore rule or not */
        function isIgnoreRule(ruleId) {
            return ignoreTests.some((test) => test(ruleId));
        }
        directives.disableBlock({ line: 1, column: 0 }, isIgnoreRule, {
            loc: { line: 1, column: 0 }
        });
        return {
            SvelteScriptElement(node) {
                directives.enableBlock(node.startTag.loc.end, isIgnoreRule, {
                    loc: node.startTag.loc.end
                });
                if (node.endTag) {
                    directives.disableBlock(node.endTag.loc.start, isIgnoreRule, {
                        loc: node.endTag.loc.start
                    });
                }
            }
        };
    }
});
