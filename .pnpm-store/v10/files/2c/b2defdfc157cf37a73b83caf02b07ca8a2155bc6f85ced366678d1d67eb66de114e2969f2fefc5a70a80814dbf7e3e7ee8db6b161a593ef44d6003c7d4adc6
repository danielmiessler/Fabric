"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const utils_1 = require("../utils");
const svelte_store_1 = require("./reference-helpers/svelte-store");
exports.default = (0, utils_1.createRule)('derived-has-same-inputs-outputs', {
    meta: {
        docs: {
            description: 'derived store should use same variable names between values and callback',
            category: 'Stylistic Issues',
            recommended: false,
            conflictWithPrettier: false
        },
        schema: [],
        messages: {
            unexpected: "The argument name should be '{{name}}'."
        },
        type: 'suggestion'
    },
    create(context) {
        /** check node type */
        function isIdentifierOrArrayExpression(node) {
            return ['Identifier', 'ArrayExpression'].includes(node.type);
        }
        /** check node type */
        function isFunctionExpression(node) {
            return ['ArrowFunctionExpression', 'FunctionExpression'].includes(node.type);
        }
        /**
         * Check for identifier type.
         * e.g. derived(a, ($a) => {});
         */
        function checkIdentifier(context, args, fn) {
            const fnParam = fn.params[0];
            if (fnParam.type !== 'Identifier')
                return;
            const expectedName = `$${args.name}`;
            if (expectedName !== fnParam.name) {
                context.report({
                    node: fn,
                    loc: fnParam.loc,
                    messageId: 'unexpected',
                    data: { name: expectedName }
                });
            }
        }
        /**
         * Check for array type.
         * e.g. derived([ a, b ], ([ $a, $b ]) => {})
         */
        function checkArrayExpression(context, args, fn) {
            const fnParam = fn.params[0];
            if (fnParam.type !== 'ArrayPattern')
                return;
            const argNames = args.elements.map((element) => {
                return element && element.type === 'Identifier' ? element.name : null;
            });
            fnParam.elements.forEach((element, index) => {
                const argName = argNames[index];
                if (element && element.type === 'Identifier' && argName) {
                    const expectedName = `$${argName}`;
                    if (expectedName !== element.name) {
                        context.report({
                            node: fn,
                            loc: element.loc,
                            messageId: 'unexpected',
                            data: { name: expectedName }
                        });
                    }
                }
            });
        }
        return {
            Program() {
                for (const { node } of (0, svelte_store_1.extractStoreReferences)(context, ['derived'])) {
                    const [args, fn] = node.arguments;
                    if (!args || !isIdentifierOrArrayExpression(args))
                        continue;
                    if (!fn || !isFunctionExpression(fn))
                        continue;
                    if (!fn.params || fn.params.length === 0)
                        continue;
                    if (args.type === 'Identifier')
                        checkIdentifier(context, args, fn);
                    else
                        checkArrayExpression(context, args, fn);
                }
            }
        };
    }
});
