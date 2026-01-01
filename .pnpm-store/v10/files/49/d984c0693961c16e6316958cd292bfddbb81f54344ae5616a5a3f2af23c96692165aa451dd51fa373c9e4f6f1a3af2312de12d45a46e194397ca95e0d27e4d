"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_compile_warns_1 = require("../shared/svelte-compile-warns");
const compat_1 = require("../utils/compat");
exports.default = (0, utils_1.createRule)('valid-compile', {
    meta: {
        docs: {
            description: 'disallow warnings when compiling.',
            category: 'Possible Errors',
            recommended: true
        },
        schema: [
            {
                type: 'object',
                properties: {
                    ignoreWarnings: { type: 'boolean' }
                },
                additionalProperties: false
            }
        ],
        messages: {},
        type: 'problem'
    },
    create(context) {
        const sourceCode = (0, compat_1.getSourceCode)(context);
        if (!sourceCode.parserServices.isSvelte) {
            return {};
        }
        const onwarn = sourceCode.parserServices.svelteParseContext?.svelteConfig?.onwarn;
        const transform = onwarn
            ? (warning) => {
                if (!warning.code)
                    return warning;
                let result = null;
                onwarn(warning, (reportWarn) => (result = reportWarn));
                return result;
            }
            : (warning) => warning;
        const ignoreWarnings = Boolean(context.options[0]?.ignoreWarnings);
        const ignores = [
            'missing-declaration',
            // Svelte v4
            'dynamic-slot-name',
            // Svelte v5
            'invalid-slot-name'
        ];
        /**
         * report
         */
        function report({ warnings, kind }) {
            for (const warn of warnings) {
                if (warn.code && ignores.includes(warn.code)) {
                    continue;
                }
                const reportWarn = kind === 'warn' ? transform(warn) : warn;
                if (!reportWarn) {
                    continue;
                }
                context.report({
                    loc: {
                        start: reportWarn.start || reportWarn.end || { line: 1, column: 0 },
                        end: reportWarn.end || reportWarn.start || { line: 1, column: 0 }
                    },
                    message: `${reportWarn.message}${reportWarn.code ? `(${reportWarn.code})` : ''}`
                });
            }
        }
        return {
            'Program:exit'() {
                const result = (0, svelte_compile_warns_1.getSvelteCompileWarnings)(context);
                if (ignoreWarnings && result.kind === 'warn') {
                    return;
                }
                report(result);
            }
        };
    }
});
