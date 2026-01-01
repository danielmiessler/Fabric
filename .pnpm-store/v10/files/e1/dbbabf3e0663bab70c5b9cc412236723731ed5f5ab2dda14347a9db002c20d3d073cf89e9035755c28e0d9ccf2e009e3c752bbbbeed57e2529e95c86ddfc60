"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
exports.default = (0, utils_1.createRule)('no-spaces-around-equal-signs-in-attribute', {
    meta: {
        docs: {
            description: 'disallow spaces around equal signs in attribute',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: true
        },
        schema: [],
        fixable: 'whitespace',
        messages: {
            noSpaces: 'Unexpected spaces found around equal signs.'
        },
        type: 'layout'
    },
    create(ctx) {
        const source = ctx.getSourceCode();
        /**
         * Returns source text between attribute key and value, and range of that source
         */
        function getAttrEq(node) {
            const keyRange = node.key.range;
            const eqSource = /^[\s=]*/u.exec(source.text.slice(keyRange[1], node.range[1]))[0];
            const valueStart = keyRange[1] + eqSource.length;
            return [eqSource, [keyRange[1], valueStart]];
        }
        /**
         * Returns true if string contains whitespace characters
         */
        function containsWhitespace(string) {
            return /\s/u.test(string);
        }
        return {
            'SvelteAttribute, SvelteDirective, SvelteStyleDirective, SvelteSpecialDirective'(node) {
                const [eqSource, range] = getAttrEq(node);
                if (!containsWhitespace(eqSource))
                    return;
                const loc = {
                    start: source.getLocFromIndex(range[0]),
                    end: source.getLocFromIndex(range[1])
                };
                ctx.report({
                    loc,
                    messageId: 'noSpaces',
                    *fix(fixer) {
                        yield fixer.replaceTextRange(range, '=');
                    }
                });
            }
        };
    }
});
