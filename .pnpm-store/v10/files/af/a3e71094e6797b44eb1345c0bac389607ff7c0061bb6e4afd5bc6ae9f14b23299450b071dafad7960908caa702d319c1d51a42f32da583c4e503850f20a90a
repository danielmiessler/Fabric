"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const svelte_compile_warns_1 = require("../shared/svelte-compile-warns");
const utils_1 = require("../utils");
const ignore_comment_1 = require("../shared/svelte-compile-warns/ignore-comment");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('no-unused-svelte-ignore', {
    meta: {
        docs: {
            description: 'disallow unused svelte-ignore comments',
            category: 'Best Practices',
            recommended: true
        },
        schema: [],
        messages: {
            unused: 'svelte-ignore comment is used, but not warned',
            missingCode: 'svelte-ignore comment must include the code'
        },
        type: 'suggestion'
    },
    create(context) {
        const sourceCode = (0, compat_1.getSourceCode)(context);
        if (!sourceCode.parserServices.isSvelte) {
            return {};
        }
        const ignoreComments = [];
        for (const item of (0, ignore_comment_1.getSvelteIgnoreItems)(context)) {
            if (item.code == null) {
                context.report({
                    node: item.token,
                    messageId: 'missingCode'
                });
            }
            else {
                ignoreComments.push(item);
            }
        }
        if (!ignoreComments.length) {
            return {};
        }
        const warnings = (0, svelte_compile_warns_1.getSvelteCompileWarnings)(context);
        if (warnings.kind === 'error') {
            return {};
        }
        for (const unused of warnings.unusedIgnores) {
            context.report({
                loc: {
                    start: sourceCode.getLocFromIndex(unused.range[0]),
                    end: sourceCode.getLocFromIndex(unused.range[1])
                },
                messageId: 'unused'
            });
        }
        return {};
    }
});
